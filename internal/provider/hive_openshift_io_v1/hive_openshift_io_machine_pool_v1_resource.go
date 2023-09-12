/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &HiveOpenshiftIoMachinePoolV1Resource{}
	_ resource.ResourceWithConfigure   = &HiveOpenshiftIoMachinePoolV1Resource{}
	_ resource.ResourceWithImportState = &HiveOpenshiftIoMachinePoolV1Resource{}
)

func NewHiveOpenshiftIoMachinePoolV1Resource() resource.Resource {
	return &HiveOpenshiftIoMachinePoolV1Resource{}
}

type HiveOpenshiftIoMachinePoolV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type HiveOpenshiftIoMachinePoolV1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Autoscaling *struct {
			MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
			MinReplicas *int64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		} `tfsdk:"autoscaling" json:"autoscaling,omitempty"`
		ClusterDeploymentRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cluster_deployment_ref" json:"clusterDeploymentRef,omitempty"`
		Labels   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Name     *string            `tfsdk:"name" json:"name,omitempty"`
		Platform *struct {
			Alibabacloud *struct {
				ImageID            *string   `tfsdk:"image_id" json:"imageID,omitempty"`
				InstanceType       *string   `tfsdk:"instance_type" json:"instanceType,omitempty"`
				SystemDiskCategory *string   `tfsdk:"system_disk_category" json:"systemDiskCategory,omitempty"`
				SystemDiskSize     *int64    `tfsdk:"system_disk_size" json:"systemDiskSize,omitempty"`
				Zones              *[]string `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"alibabacloud" json:"alibabacloud,omitempty"`
			Aws *struct {
				MetadataService *struct {
					Authentication *string `tfsdk:"authentication" json:"authentication,omitempty"`
				} `tfsdk:"metadata_service" json:"metadataService,omitempty"`
				RootVolume *struct {
					Iops      *int64  `tfsdk:"iops" json:"iops,omitempty"`
					KmsKeyARN *string `tfsdk:"kms_key_arn" json:"kmsKeyARN,omitempty"`
					Size      *int64  `tfsdk:"size" json:"size,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"root_volume" json:"rootVolume,omitempty"`
				SpotMarketOptions *struct {
					MaxPrice *string `tfsdk:"max_price" json:"maxPrice,omitempty"`
				} `tfsdk:"spot_market_options" json:"spotMarketOptions,omitempty"`
				Subnets *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
				Type    *string   `tfsdk:"type" json:"type,omitempty"`
				Zones   *[]string `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				OsDisk *struct {
					DiskEncryptionSet *struct {
						Name           *string `tfsdk:"name" json:"name,omitempty"`
						ResourceGroup  *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
						SubscriptionId *string `tfsdk:"subscription_id" json:"subscriptionId,omitempty"`
					} `tfsdk:"disk_encryption_set" json:"diskEncryptionSet,omitempty"`
					DiskSizeGB *int64  `tfsdk:"disk_size_gb" json:"diskSizeGB,omitempty"`
					DiskType   *string `tfsdk:"disk_type" json:"diskType,omitempty"`
				} `tfsdk:"os_disk" json:"osDisk,omitempty"`
				OsImage *struct {
					Offer     *string `tfsdk:"offer" json:"offer,omitempty"`
					Publisher *string `tfsdk:"publisher" json:"publisher,omitempty"`
					Sku       *string `tfsdk:"sku" json:"sku,omitempty"`
					Version   *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"os_image" json:"osImage,omitempty"`
				Type  *string   `tfsdk:"type" json:"type,omitempty"`
				Zones *[]string `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Gcp *struct {
				OsDisk *struct {
					DiskSizeGB    *int64  `tfsdk:"disk_size_gb" json:"diskSizeGB,omitempty"`
					DiskType      *string `tfsdk:"disk_type" json:"diskType,omitempty"`
					EncryptionKey *struct {
						KmsKey *struct {
							KeyRing   *string `tfsdk:"key_ring" json:"keyRing,omitempty"`
							Location  *string `tfsdk:"location" json:"location,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							ProjectID *string `tfsdk:"project_id" json:"projectID,omitempty"`
						} `tfsdk:"kms_key" json:"kmsKey,omitempty"`
						KmsKeyServiceAccount *string `tfsdk:"kms_key_service_account" json:"kmsKeyServiceAccount,omitempty"`
					} `tfsdk:"encryption_key" json:"encryptionKey,omitempty"`
				} `tfsdk:"os_disk" json:"osDisk,omitempty"`
				Type  *string   `tfsdk:"type" json:"type,omitempty"`
				Zones *[]string `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
			Ibmcloud *struct {
				BootVolume *struct {
					EncryptionKey *string `tfsdk:"encryption_key" json:"encryptionKey,omitempty"`
				} `tfsdk:"boot_volume" json:"bootVolume,omitempty"`
				DedicatedHosts *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Profile *string `tfsdk:"profile" json:"profile,omitempty"`
				} `tfsdk:"dedicated_hosts" json:"dedicatedHosts,omitempty"`
				Type  *string   `tfsdk:"type" json:"type,omitempty"`
				Zones *[]string `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"ibmcloud" json:"ibmcloud,omitempty"`
			Openstack *struct {
				Flavor     *string `tfsdk:"flavor" json:"flavor,omitempty"`
				RootVolume *struct {
					Size *int64  `tfsdk:"size" json:"size,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"root_volume" json:"rootVolume,omitempty"`
			} `tfsdk:"openstack" json:"openstack,omitempty"`
			Ovirt *struct {
				Cpu *struct {
					Cores   *int64 `tfsdk:"cores" json:"cores,omitempty"`
					Sockets *int64 `tfsdk:"sockets" json:"sockets,omitempty"`
				} `tfsdk:"cpu" json:"cpu,omitempty"`
				MemoryMB *int64 `tfsdk:"memory_mb" json:"memoryMB,omitempty"`
				OsDisk   *struct {
					SizeGB *int64 `tfsdk:"size_gb" json:"sizeGB,omitempty"`
				} `tfsdk:"os_disk" json:"osDisk,omitempty"`
				VmType *string `tfsdk:"vm_type" json:"vmType,omitempty"`
			} `tfsdk:"ovirt" json:"ovirt,omitempty"`
			Vsphere *struct {
				CoresPerSocket *int64 `tfsdk:"cores_per_socket" json:"coresPerSocket,omitempty"`
				Cpus           *int64 `tfsdk:"cpus" json:"cpus,omitempty"`
				MemoryMB       *int64 `tfsdk:"memory_mb" json:"memoryMB,omitempty"`
				OsDisk         *struct {
					DiskSizeGB *int64 `tfsdk:"disk_size_gb" json:"diskSizeGB,omitempty"`
				} `tfsdk:"os_disk" json:"osDisk,omitempty"`
			} `tfsdk:"vsphere" json:"vsphere,omitempty"`
		} `tfsdk:"platform" json:"platform,omitempty"`
		Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Taints   *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_machine_pool_v1"
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachinePool is the Schema for the machinepools API",
		MarkdownDescription: "MachinePool is the Schema for the machinepools API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "MachinePoolSpec defines the desired state of MachinePool",
				MarkdownDescription: "MachinePoolSpec defines the desired state of MachinePool",
				Attributes: map[string]schema.Attribute{
					"autoscaling": schema.SingleNestedAttribute{
						Description:         "Autoscaling is the details for auto-scaling the machine pool. Replicas and autoscaling cannot be used together.",
						MarkdownDescription: "Autoscaling is the details for auto-scaling the machine pool. Replicas and autoscaling cannot be used together.",
						Attributes: map[string]schema.Attribute{
							"max_replicas": schema.Int64Attribute{
								Description:         "MaxReplicas is the maximum number of replicas for the machine pool.",
								MarkdownDescription: "MaxReplicas is the maximum number of replicas for the machine pool.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"min_replicas": schema.Int64Attribute{
								Description:         "MinReplicas is the minimum number of replicas for the machine pool.",
								MarkdownDescription: "MinReplicas is the minimum number of replicas for the machine pool.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_deployment_ref": schema.SingleNestedAttribute{
						Description:         "ClusterDeploymentRef references the cluster deployment to which this machine pool belongs.",
						MarkdownDescription: "ClusterDeploymentRef references the cluster deployment to which this machine pool belongs.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"labels": schema.MapAttribute{
						Description:         "Map of label string keys and values that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node labels on an ongoing basis.",
						MarkdownDescription: "Map of label string keys and values that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node labels on an ongoing basis.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is the name of the machine pool.",
						MarkdownDescription: "Name is the name of the machine pool.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"platform": schema.SingleNestedAttribute{
						Description:         "Platform is configuration for machine pool specific to the platform.",
						MarkdownDescription: "Platform is configuration for machine pool specific to the platform.",
						Attributes: map[string]schema.Attribute{
							"alibabacloud": schema.SingleNestedAttribute{
								Description:         "AlibabaCloud is the configuration used when installing on Alibaba Cloud.",
								MarkdownDescription: "AlibabaCloud is the configuration used when installing on Alibaba Cloud.",
								Attributes: map[string]schema.Attribute{
									"image_id": schema.StringAttribute{
										Description:         "ImageID is the Image ID that should be used to create ECS instance. If set, the ImageID should belong to the same region as the cluster.",
										MarkdownDescription: "ImageID is the Image ID that should be used to create ECS instance. If set, the ImageID should belong to the same region as the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"instance_type": schema.StringAttribute{
										Description:         "InstanceType defines the ECS instance type. eg. ecs.g6.large",
										MarkdownDescription: "InstanceType defines the ECS instance type. eg. ecs.g6.large",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"system_disk_category": schema.StringAttribute{
										Description:         "SystemDiskCategory defines the category of the system disk.",
										MarkdownDescription: "SystemDiskCategory defines the category of the system disk.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "cloud_efficiency", "cloud_essd"),
										},
									},

									"system_disk_size": schema.Int64Attribute{
										Description:         "SystemDiskSize defines the size of the system disk in gibibytes (GiB).",
										MarkdownDescription: "SystemDiskSize defines the size of the system disk in gibibytes (GiB).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(120),
										},
									},

									"zones": schema.ListAttribute{
										Description:         "Zones is list of availability zones that can be used. eg. ['cn-hangzhou-i', 'cn-hangzhou-h', 'cn-hangzhou-j']",
										MarkdownDescription: "Zones is list of availability zones that can be used. eg. ['cn-hangzhou-i', 'cn-hangzhou-h', 'cn-hangzhou-j']",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws": schema.SingleNestedAttribute{
								Description:         "AWS is the configuration used when installing on AWS.",
								MarkdownDescription: "AWS is the configuration used when installing on AWS.",
								Attributes: map[string]schema.Attribute{
									"metadata_service": schema.SingleNestedAttribute{
										Description:         "EC2MetadataOptions defines metadata service interaction options for EC2 instances in the machine pool.",
										MarkdownDescription: "EC2MetadataOptions defines metadata service interaction options for EC2 instances in the machine pool.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.StringAttribute{
												Description:         "Authentication determines whether or not the host requires the use of authentication when interacting with the metadata service. When using authentication, this enforces v2 interaction method (IMDSv2) with the metadata service. When omitted, this means the user has no opinion and the value is left to the platform to choose a good default, which is subject to change over time. The current default is optional. At this point this field represents 'HttpTokens' parameter from 'InstanceMetadataOptionsRequest' structure in AWS EC2 API https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceMetadataOptionsRequest.html",
												MarkdownDescription: "Authentication determines whether or not the host requires the use of authentication when interacting with the metadata service. When using authentication, this enforces v2 interaction method (IMDSv2) with the metadata service. When omitted, this means the user has no opinion and the value is left to the platform to choose a good default, which is subject to change over time. The current default is optional. At this point this field represents 'HttpTokens' parameter from 'InstanceMetadataOptionsRequest' structure in AWS EC2 API https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceMetadataOptionsRequest.html",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_volume": schema.SingleNestedAttribute{
										Description:         "EC2RootVolume defines the storage for ec2 instance.",
										MarkdownDescription: "EC2RootVolume defines the storage for ec2 instance.",
										Attributes: map[string]schema.Attribute{
											"iops": schema.Int64Attribute{
												Description:         "IOPS defines the iops for the storage.",
												MarkdownDescription: "IOPS defines the iops for the storage.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kms_key_arn": schema.StringAttribute{
												Description:         "The KMS key that will be used to encrypt the EBS volume. If no key is provided the default KMS key for the account will be used. https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsDefaultKmsKeyId.html",
												MarkdownDescription: "The KMS key that will be used to encrypt the EBS volume. If no key is provided the default KMS key for the account will be used. https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsDefaultKmsKeyId.html",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"size": schema.Int64Attribute{
												Description:         "Size defines the size of the storage.",
												MarkdownDescription: "Size defines the size of the storage.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type defines the type of the storage.",
												MarkdownDescription: "Type defines the type of the storage.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"spot_market_options": schema.SingleNestedAttribute{
										Description:         "SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.",
										MarkdownDescription: "SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.",
										Attributes: map[string]schema.Attribute{
											"max_price": schema.StringAttribute{
												Description:         "The maximum price the user is willing to pay for their instances Default: On-Demand price",
												MarkdownDescription: "The maximum price the user is willing to pay for their instances Default: On-Demand price",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"subnets": schema.ListAttribute{
										Description:         "Subnets is the list of IDs of subnets to which to attach the machines. There must be exactly one subnet for each availability zone used. These subnets may be public or private. As a special case, for consistency with install-config, you may specify exactly one private and one public subnet for each availability zone. In this case, the public subnets will be filtered out and only the private subnets will be used. If empty/omitted, we will look for subnets in each availability zone tagged with Name=<clusterID>-private-<az>.",
										MarkdownDescription: "Subnets is the list of IDs of subnets to which to attach the machines. There must be exactly one subnet for each availability zone used. These subnets may be public or private. As a special case, for consistency with install-config, you may specify exactly one private and one public subnet for each availability zone. In this case, the public subnets will be filtered out and only the private subnets will be used. If empty/omitted, we will look for subnets in each availability zone tagged with Name=<clusterID>-private-<az>.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "InstanceType defines the ec2 instance type. eg. m4-large",
										MarkdownDescription: "InstanceType defines the ec2 instance type. eg. m4-large",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"zones": schema.ListAttribute{
										Description:         "Zones is list of availability zones that can be used.",
										MarkdownDescription: "Zones is list of availability zones that can be used.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": schema.SingleNestedAttribute{
								Description:         "Azure is the configuration used when installing on Azure.",
								MarkdownDescription: "Azure is the configuration used when installing on Azure.",
								Attributes: map[string]schema.Attribute{
									"os_disk": schema.SingleNestedAttribute{
										Description:         "OSDisk defines the storage for instance.",
										MarkdownDescription: "OSDisk defines the storage for instance.",
										Attributes: map[string]schema.Attribute{
											"disk_encryption_set": schema.SingleNestedAttribute{
												Description:         "DiskEncryptionSet defines a disk encryption set.",
												MarkdownDescription: "DiskEncryptionSet defines a disk encryption set.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name is the name of the disk encryption set.",
														MarkdownDescription: "Name is the name of the disk encryption set.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"resource_group": schema.StringAttribute{
														Description:         "ResourceGroup defines the Azure resource group used by the disk encryption set.",
														MarkdownDescription: "ResourceGroup defines the Azure resource group used by the disk encryption set.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"subscription_id": schema.StringAttribute{
														Description:         "SubscriptionID defines the Azure subscription the disk encryption set is in.",
														MarkdownDescription: "SubscriptionID defines the Azure subscription the disk encryption set is in.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"disk_size_gb": schema.Int64Attribute{
												Description:         "DiskSizeGB defines the size of disk in GB.",
												MarkdownDescription: "DiskSizeGB defines the size of disk in GB.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"disk_type": schema.StringAttribute{
												Description:         "DiskType defines the type of disk. For control plane nodes, the valid values are Premium_LRS and StandardSSD_LRS. Default is Premium_LRS.",
												MarkdownDescription: "DiskType defines the type of disk. For control plane nodes, the valid values are Premium_LRS and StandardSSD_LRS. Default is Premium_LRS.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Standard_LRS", "Premium_LRS", "StandardSSD_LRS"),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"os_image": schema.SingleNestedAttribute{
										Description:         "OSImage defines the image to use for the OS.",
										MarkdownDescription: "OSImage defines the image to use for the OS.",
										Attributes: map[string]schema.Attribute{
											"offer": schema.StringAttribute{
												Description:         "Offer is the offer of the image.",
												MarkdownDescription: "Offer is the offer of the image.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"publisher": schema.StringAttribute{
												Description:         "Publisher is the publisher of the image.",
												MarkdownDescription: "Publisher is the publisher of the image.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"sku": schema.StringAttribute{
												Description:         "SKU is the SKU of the image.",
												MarkdownDescription: "SKU is the SKU of the image.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "Version is the version of the image.",
												MarkdownDescription: "Version is the version of the image.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "InstanceType defines the azure instance type. eg. Standard_DS_V2",
										MarkdownDescription: "InstanceType defines the azure instance type. eg. Standard_DS_V2",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"zones": schema.ListAttribute{
										Description:         "Zones is list of availability zones that can be used. eg. ['1', '2', '3']",
										MarkdownDescription: "Zones is list of availability zones that can be used. eg. ['1', '2', '3']",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"gcp": schema.SingleNestedAttribute{
								Description:         "GCP is the configuration used when installing on GCP.",
								MarkdownDescription: "GCP is the configuration used when installing on GCP.",
								Attributes: map[string]schema.Attribute{
									"os_disk": schema.SingleNestedAttribute{
										Description:         "OSDisk defines the storage for instances.",
										MarkdownDescription: "OSDisk defines the storage for instances.",
										Attributes: map[string]schema.Attribute{
											"disk_size_gb": schema.Int64Attribute{
												Description:         "DiskSizeGB defines the size of disk in GB. Defaulted internally to 128.",
												MarkdownDescription: "DiskSizeGB defines the size of disk in GB. Defaulted internally to 128.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(16),
													int64validator.AtMost(65536),
												},
											},

											"disk_type": schema.StringAttribute{
												Description:         "DiskType defines the type of disk. The valid values are pd-standard and pd-ssd. Defaulted internally to pd-ssd.",
												MarkdownDescription: "DiskType defines the type of disk. The valid values are pd-standard and pd-ssd. Defaulted internally to pd-ssd.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("pd-ssd", "pd-standard"),
												},
											},

											"encryption_key": schema.SingleNestedAttribute{
												Description:         "EncryptionKey defines the KMS key to be used to encrypt the disk.",
												MarkdownDescription: "EncryptionKey defines the KMS key to be used to encrypt the disk.",
												Attributes: map[string]schema.Attribute{
													"kms_key": schema.SingleNestedAttribute{
														Description:         "KMSKey is a reference to a KMS Key to use for the encryption.",
														MarkdownDescription: "KMSKey is a reference to a KMS Key to use for the encryption.",
														Attributes: map[string]schema.Attribute{
															"key_ring": schema.StringAttribute{
																Description:         "KeyRing is the name of the KMS Key Ring which the KMS Key belongs to.",
																MarkdownDescription: "KeyRing is the name of the KMS Key Ring which the KMS Key belongs to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"location": schema.StringAttribute{
																Description:         "Location is the GCP location in which the Key Ring exists.",
																MarkdownDescription: "Location is the GCP location in which the Key Ring exists.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the customer managed encryption key to be used for the disk encryption.",
																MarkdownDescription: "Name is the name of the customer managed encryption key to be used for the disk encryption.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"project_id": schema.StringAttribute{
																Description:         "ProjectID is the ID of the Project in which the KMS Key Ring exists. Defaults to the VM ProjectID if not set.",
																MarkdownDescription: "ProjectID is the ID of the Project in which the KMS Key Ring exists. Defaults to the VM ProjectID if not set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kms_key_service_account": schema.StringAttribute{
														Description:         "KMSKeyServiceAccount is the service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used. See https://cloud.google.com/compute/docs/access/service-accounts#compute_engine_service_account for details on the default service account.",
														MarkdownDescription: "KMSKeyServiceAccount is the service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used. See https://cloud.google.com/compute/docs/access/service-accounts#compute_engine_service_account for details on the default service account.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "InstanceType defines the GCP instance type. eg. n1-standard-4",
										MarkdownDescription: "InstanceType defines the GCP instance type. eg. n1-standard-4",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"zones": schema.ListAttribute{
										Description:         "Zones is list of availability zones that can be used.",
										MarkdownDescription: "Zones is list of availability zones that can be used.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ibmcloud": schema.SingleNestedAttribute{
								Description:         "IBMCloud is the configuration used when installing on IBM Cloud.",
								MarkdownDescription: "IBMCloud is the configuration used when installing on IBM Cloud.",
								Attributes: map[string]schema.Attribute{
									"boot_volume": schema.SingleNestedAttribute{
										Description:         "BootVolume is the configuration for the machine's boot volume.",
										MarkdownDescription: "BootVolume is the configuration for the machine's boot volume.",
										Attributes: map[string]schema.Attribute{
											"encryption_key": schema.StringAttribute{
												Description:         "EncryptionKey is the CRN referencing a Key Protect or Hyper Protect Crypto Services key to use for volume encryption. If not specified, a provider managed encryption key will be used.",
												MarkdownDescription: "EncryptionKey is the CRN referencing a Key Protect or Hyper Protect Crypto Services key to use for volume encryption. If not specified, a provider managed encryption key will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"dedicated_hosts": schema.ListNestedAttribute{
										Description:         "DedicatedHosts is the configuration for the machine's dedicated host and profile.",
										MarkdownDescription: "DedicatedHosts is the configuration for the machine's dedicated host and profile.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the dedicated host to provision the machine on. If specified, machines will be created on pre-existing dedicated host.",
													MarkdownDescription: "Name is the name of the dedicated host to provision the machine on. If specified, machines will be created on pre-existing dedicated host.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"profile": schema.StringAttribute{
													Description:         "Profile is the profile ID for the dedicated host. If specified, new dedicated host will be created for machines.",
													MarkdownDescription: "Profile is the profile ID for the dedicated host. If specified, new dedicated host will be created for machines.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "InstanceType is the VSI machine profile.",
										MarkdownDescription: "InstanceType is the VSI machine profile.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"zones": schema.ListAttribute{
										Description:         "Zones is the list of availability zones used for machines in the pool.",
										MarkdownDescription: "Zones is the list of availability zones used for machines in the pool.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"openstack": schema.SingleNestedAttribute{
								Description:         "OpenStack is the configuration used when installing on OpenStack.",
								MarkdownDescription: "OpenStack is the configuration used when installing on OpenStack.",
								Attributes: map[string]schema.Attribute{
									"flavor": schema.StringAttribute{
										Description:         "Flavor defines the OpenStack Nova flavor. eg. m1.large The json key here differs from the installer which uses both 'computeFlavor' and type 'type' depending on which type you're looking at, and the resulting field on the MachineSet is 'flavor'. We are opting to stay consistent with the end result.",
										MarkdownDescription: "Flavor defines the OpenStack Nova flavor. eg. m1.large The json key here differs from the installer which uses both 'computeFlavor' and type 'type' depending on which type you're looking at, and the resulting field on the MachineSet is 'flavor'. We are opting to stay consistent with the end result.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"root_volume": schema.SingleNestedAttribute{
										Description:         "RootVolume defines the root volume for instances in the machine pool. The instances use ephemeral disks if not set.",
										MarkdownDescription: "RootVolume defines the root volume for instances in the machine pool. The instances use ephemeral disks if not set.",
										Attributes: map[string]schema.Attribute{
											"size": schema.Int64Attribute{
												Description:         "Size defines the size of the volume in gibibytes (GiB). Required",
												MarkdownDescription: "Size defines the size of the volume in gibibytes (GiB). Required",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type defines the type of the volume. Required",
												MarkdownDescription: "Type defines the type of the volume. Required",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ovirt": schema.SingleNestedAttribute{
								Description:         "Ovirt is the configuration used when installing on oVirt.",
								MarkdownDescription: "Ovirt is the configuration used when installing on oVirt.",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.SingleNestedAttribute{
										Description:         "CPU defines the VM CPU.",
										MarkdownDescription: "CPU defines the VM CPU.",
										Attributes: map[string]schema.Attribute{
											"cores": schema.Int64Attribute{
												Description:         "Cores is the number of cores per socket. Total CPUs is (Sockets * Cores)",
												MarkdownDescription: "Cores is the number of cores per socket. Total CPUs is (Sockets * Cores)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"sockets": schema.Int64Attribute{
												Description:         "Sockets is the number of sockets for a VM. Total CPUs is (Sockets * Cores)",
												MarkdownDescription: "Sockets is the number of sockets for a VM. Total CPUs is (Sockets * Cores)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_mb": schema.Int64Attribute{
										Description:         "MemoryMB is the size of a VM's memory in MiBs.",
										MarkdownDescription: "MemoryMB is the size of a VM's memory in MiBs.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"os_disk": schema.SingleNestedAttribute{
										Description:         "OSDisk is the the root disk of the node.",
										MarkdownDescription: "OSDisk is the the root disk of the node.",
										Attributes: map[string]schema.Attribute{
											"size_gb": schema.Int64Attribute{
												Description:         "SizeGB size of the bootable disk in GiB.",
												MarkdownDescription: "SizeGB size of the bootable disk in GiB.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"vm_type": schema.StringAttribute{
										Description:         "VMType defines the workload type of the VM.",
										MarkdownDescription: "VMType defines the workload type of the VM.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "desktop", "server", "high_performance"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"vsphere": schema.SingleNestedAttribute{
								Description:         "VSphere is the configuration used when installing on vSphere",
								MarkdownDescription: "VSphere is the configuration used when installing on vSphere",
								Attributes: map[string]schema.Attribute{
									"cores_per_socket": schema.Int64Attribute{
										Description:         "NumCoresPerSocket is the number of cores per socket in a vm. The number of vCPUs on the vm will be NumCPUs/NumCoresPerSocket.",
										MarkdownDescription: "NumCoresPerSocket is the number of cores per socket in a vm. The number of vCPUs on the vm will be NumCPUs/NumCoresPerSocket.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"cpus": schema.Int64Attribute{
										Description:         "NumCPUs is the total number of virtual processor cores to assign a vm.",
										MarkdownDescription: "NumCPUs is the total number of virtual processor cores to assign a vm.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"memory_mb": schema.Int64Attribute{
										Description:         "Memory is the size of a VM's memory in MB.",
										MarkdownDescription: "Memory is the size of a VM's memory in MB.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"os_disk": schema.SingleNestedAttribute{
										Description:         "OSDisk defines the storage for instance.",
										MarkdownDescription: "OSDisk defines the storage for instance.",
										Attributes: map[string]schema.Attribute{
											"disk_size_gb": schema.Int64Attribute{
												Description:         "DiskSizeGB defines the size of disk in GB.",
												MarkdownDescription: "DiskSizeGB defines the size of disk in GB.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the count of machines for this machine pool. Replicas and autoscaling cannot be used together. Default is 1, if autoscaling is not used.",
						MarkdownDescription: "Replicas is the count of machines for this machine pool. Replicas and autoscaling cannot be used together. Default is 1, if autoscaling is not used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"taints": schema.ListNestedAttribute{
						Description:         "List of taints that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node taints on an ongoing basis. In case of duplicate entries, first encountered taint Value will be preserved, and the rest collapsed on the corresponding MachineSets. Note that taints are uniquely identified based on key+effect, not just key.",
						MarkdownDescription: "List of taints that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node taints on an ongoing basis. In case of duplicate entries, first encountered taint Value will be preserved, and the rest collapsed on the corresponding MachineSets. Note that taints are uniquely identified based on key+effect, not just key.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"time_added": schema.StringAttribute{
									Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
									MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"value": schema.StringAttribute{
									Description:         "The taint value corresponding to the taint key.",
									MarkdownDescription: "The taint value corresponding to the taint key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_machine_pool_v1")

	var model HiveOpenshiftIoMachinePoolV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("MachinePool")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "machinepools"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HiveOpenshiftIoMachinePoolV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_machine_pool_v1")

	var data HiveOpenshiftIoMachinePoolV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "machinepools"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HiveOpenshiftIoMachinePoolV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_machine_pool_v1")

	var model HiveOpenshiftIoMachinePoolV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("MachinePool")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "machinepools"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HiveOpenshiftIoMachinePoolV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_machine_pool_v1")

	var data HiveOpenshiftIoMachinePoolV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "machinepools"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "machinepools"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
