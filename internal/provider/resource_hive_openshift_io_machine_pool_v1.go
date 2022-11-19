/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type HiveOpenshiftIoMachinePoolV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoMachinePoolV1Resource)(nil)
)

type HiveOpenshiftIoMachinePoolV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoMachinePoolV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Autoscaling *struct {
			MaxReplicas *int64 `tfsdk:"max_replicas" yaml:"maxReplicas,omitempty"`

			MinReplicas *int64 `tfsdk:"min_replicas" yaml:"minReplicas,omitempty"`
		} `tfsdk:"autoscaling" yaml:"autoscaling,omitempty"`

		ClusterDeploymentRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"cluster_deployment_ref" yaml:"clusterDeploymentRef,omitempty"`

		Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Platform *struct {
			Alibabacloud *struct {
				ImageID *string `tfsdk:"image_id" yaml:"imageID,omitempty"`

				InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

				SystemDiskCategory *string `tfsdk:"system_disk_category" yaml:"systemDiskCategory,omitempty"`

				SystemDiskSize *int64 `tfsdk:"system_disk_size" yaml:"systemDiskSize,omitempty"`

				Zones *[]string `tfsdk:"zones" yaml:"zones,omitempty"`
			} `tfsdk:"alibabacloud" yaml:"alibabacloud,omitempty"`

			Aws *struct {
				RootVolume *struct {
					Iops *int64 `tfsdk:"iops" yaml:"iops,omitempty"`

					KmsKeyARN *string `tfsdk:"kms_key_arn" yaml:"kmsKeyARN,omitempty"`

					Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"root_volume" yaml:"rootVolume,omitempty"`

				SpotMarketOptions *struct {
					MaxPrice *string `tfsdk:"max_price" yaml:"maxPrice,omitempty"`
				} `tfsdk:"spot_market_options" yaml:"spotMarketOptions,omitempty"`

				Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Zones *[]string `tfsdk:"zones" yaml:"zones,omitempty"`
			} `tfsdk:"aws" yaml:"aws,omitempty"`

			Azure *struct {
				OsDisk *struct {
					DiskEncryptionSet *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						ResourceGroup *string `tfsdk:"resource_group" yaml:"resourceGroup,omitempty"`

						SubscriptionId *string `tfsdk:"subscription_id" yaml:"subscriptionId,omitempty"`
					} `tfsdk:"disk_encryption_set" yaml:"diskEncryptionSet,omitempty"`

					DiskSizeGB *int64 `tfsdk:"disk_size_gb" yaml:"diskSizeGB,omitempty"`

					DiskType *string `tfsdk:"disk_type" yaml:"diskType,omitempty"`
				} `tfsdk:"os_disk" yaml:"osDisk,omitempty"`

				OsImage *struct {
					Offer *string `tfsdk:"offer" yaml:"offer,omitempty"`

					Publisher *string `tfsdk:"publisher" yaml:"publisher,omitempty"`

					Sku *string `tfsdk:"sku" yaml:"sku,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"os_image" yaml:"osImage,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Zones *[]string `tfsdk:"zones" yaml:"zones,omitempty"`
			} `tfsdk:"azure" yaml:"azure,omitempty"`

			Gcp *struct {
				OsDisk *struct {
					DiskSizeGB *int64 `tfsdk:"disk_size_gb" yaml:"diskSizeGB,omitempty"`

					DiskType *string `tfsdk:"disk_type" yaml:"diskType,omitempty"`

					EncryptionKey *struct {
						KmsKey *struct {
							KeyRing *string `tfsdk:"key_ring" yaml:"keyRing,omitempty"`

							Location *string `tfsdk:"location" yaml:"location,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							ProjectID *string `tfsdk:"project_id" yaml:"projectID,omitempty"`
						} `tfsdk:"kms_key" yaml:"kmsKey,omitempty"`

						KmsKeyServiceAccount *string `tfsdk:"kms_key_service_account" yaml:"kmsKeyServiceAccount,omitempty"`
					} `tfsdk:"encryption_key" yaml:"encryptionKey,omitempty"`
				} `tfsdk:"os_disk" yaml:"osDisk,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Zones *[]string `tfsdk:"zones" yaml:"zones,omitempty"`
			} `tfsdk:"gcp" yaml:"gcp,omitempty"`

			Ibmcloud *struct {
				BootVolume *struct {
					EncryptionKey *string `tfsdk:"encryption_key" yaml:"encryptionKey,omitempty"`
				} `tfsdk:"boot_volume" yaml:"bootVolume,omitempty"`

				DedicatedHosts *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Profile *string `tfsdk:"profile" yaml:"profile,omitempty"`
				} `tfsdk:"dedicated_hosts" yaml:"dedicatedHosts,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Zones *[]string `tfsdk:"zones" yaml:"zones,omitempty"`
			} `tfsdk:"ibmcloud" yaml:"ibmcloud,omitempty"`

			Openstack *struct {
				Flavor *string `tfsdk:"flavor" yaml:"flavor,omitempty"`

				RootVolume *struct {
					Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"root_volume" yaml:"rootVolume,omitempty"`
			} `tfsdk:"openstack" yaml:"openstack,omitempty"`

			Ovirt *struct {
				Cpu *struct {
					Cores *int64 `tfsdk:"cores" yaml:"cores,omitempty"`

					Sockets *int64 `tfsdk:"sockets" yaml:"sockets,omitempty"`
				} `tfsdk:"cpu" yaml:"cpu,omitempty"`

				MemoryMB *int64 `tfsdk:"memory_mb" yaml:"memoryMB,omitempty"`

				OsDisk *struct {
					SizeGB *int64 `tfsdk:"size_gb" yaml:"sizeGB,omitempty"`
				} `tfsdk:"os_disk" yaml:"osDisk,omitempty"`

				VmType *string `tfsdk:"vm_type" yaml:"vmType,omitempty"`
			} `tfsdk:"ovirt" yaml:"ovirt,omitempty"`

			Vsphere *struct {
				CoresPerSocket *int64 `tfsdk:"cores_per_socket" yaml:"coresPerSocket,omitempty"`

				Cpus *int64 `tfsdk:"cpus" yaml:"cpus,omitempty"`

				MemoryMB *int64 `tfsdk:"memory_mb" yaml:"memoryMB,omitempty"`

				OsDisk *struct {
					DiskSizeGB *int64 `tfsdk:"disk_size_gb" yaml:"diskSizeGB,omitempty"`
				} `tfsdk:"os_disk" yaml:"osDisk,omitempty"`
			} `tfsdk:"vsphere" yaml:"vsphere,omitempty"`
		} `tfsdk:"platform" yaml:"platform,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Taints *[]struct {
			Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			TimeAdded *string `tfsdk:"time_added" yaml:"timeAdded,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"taints" yaml:"taints,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoMachinePoolV1Resource() resource.Resource {
	return &HiveOpenshiftIoMachinePoolV1Resource{}
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_machine_pool_v1"
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "MachinePool is the Schema for the machinepools API",
		MarkdownDescription: "MachinePool is the Schema for the machinepools API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "MachinePoolSpec defines the desired state of MachinePool",
				MarkdownDescription: "MachinePoolSpec defines the desired state of MachinePool",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"autoscaling": {
						Description:         "Autoscaling is the details for auto-scaling the machine pool. Replicas and autoscaling cannot be used together.",
						MarkdownDescription: "Autoscaling is the details for auto-scaling the machine pool. Replicas and autoscaling cannot be used together.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_replicas": {
								Description:         "MaxReplicas is the maximum number of replicas for the machine pool.",
								MarkdownDescription: "MaxReplicas is the maximum number of replicas for the machine pool.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"min_replicas": {
								Description:         "MinReplicas is the minimum number of replicas for the machine pool.",
								MarkdownDescription: "MinReplicas is the minimum number of replicas for the machine pool.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_deployment_ref": {
						Description:         "ClusterDeploymentRef references the cluster deployment to which this machine pool belongs.",
						MarkdownDescription: "ClusterDeploymentRef references the cluster deployment to which this machine pool belongs.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"labels": {
						Description:         "Map of label string keys and values that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node labels on an ongoing basis.",
						MarkdownDescription: "Map of label string keys and values that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node labels on an ongoing basis.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "Name is the name of the machine pool.",
						MarkdownDescription: "Name is the name of the machine pool.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"platform": {
						Description:         "Platform is configuration for machine pool specific to the platform.",
						MarkdownDescription: "Platform is configuration for machine pool specific to the platform.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alibabacloud": {
								Description:         "AlibabaCloud is the configuration used when installing on Alibaba Cloud.",
								MarkdownDescription: "AlibabaCloud is the configuration used when installing on Alibaba Cloud.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image_id": {
										Description:         "ImageID is the Image ID that should be used to create ECS instance. If set, the ImageID should belong to the same region as the cluster.",
										MarkdownDescription: "ImageID is the Image ID that should be used to create ECS instance. If set, the ImageID should belong to the same region as the cluster.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"instance_type": {
										Description:         "InstanceType defines the ECS instance type. eg. ecs.g6.large",
										MarkdownDescription: "InstanceType defines the ECS instance type. eg. ecs.g6.large",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"system_disk_category": {
										Description:         "SystemDiskCategory defines the category of the system disk.",
										MarkdownDescription: "SystemDiskCategory defines the category of the system disk.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("", "cloud_efficiency", "cloud_essd"),
										},
									},

									"system_disk_size": {
										Description:         "SystemDiskSize defines the size of the system disk in gibibytes (GiB).",
										MarkdownDescription: "SystemDiskSize defines the size of the system disk in gibibytes (GiB).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(120),
										},
									},

									"zones": {
										Description:         "Zones is list of availability zones that can be used. eg. ['cn-hangzhou-i', 'cn-hangzhou-h', 'cn-hangzhou-j']",
										MarkdownDescription: "Zones is list of availability zones that can be used. eg. ['cn-hangzhou-i', 'cn-hangzhou-h', 'cn-hangzhou-j']",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws": {
								Description:         "AWS is the configuration used when installing on AWS.",
								MarkdownDescription: "AWS is the configuration used when installing on AWS.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"root_volume": {
										Description:         "EC2RootVolume defines the storage for ec2 instance.",
										MarkdownDescription: "EC2RootVolume defines the storage for ec2 instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"iops": {
												Description:         "IOPS defines the iops for the storage.",
												MarkdownDescription: "IOPS defines the iops for the storage.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kms_key_arn": {
												Description:         "The KMS key that will be used to encrypt the EBS volume. If no key is provided the default KMS key for the account will be used. https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsDefaultKmsKeyId.html",
												MarkdownDescription: "The KMS key that will be used to encrypt the EBS volume. If no key is provided the default KMS key for the account will be used. https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsDefaultKmsKeyId.html",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "Size defines the size of the storage.",
												MarkdownDescription: "Size defines the size of the storage.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "Type defines the type of the storage.",
												MarkdownDescription: "Type defines the type of the storage.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"spot_market_options": {
										Description:         "SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.",
										MarkdownDescription: "SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_price": {
												Description:         "The maximum price the user is willing to pay for their instances Default: On-Demand price",
												MarkdownDescription: "The maximum price the user is willing to pay for their instances Default: On-Demand price",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subnets": {
										Description:         "Subnets is the list of subnets to which to attach the machines. There must be exactly one private subnet for each availability zone used. If public subnets are specified, there must be exactly one private and one public subnet specified for each availability zone.",
										MarkdownDescription: "Subnets is the list of subnets to which to attach the machines. There must be exactly one private subnet for each availability zone used. If public subnets are specified, there must be exactly one private and one public subnet specified for each availability zone.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "InstanceType defines the ec2 instance type. eg. m4-large",
										MarkdownDescription: "InstanceType defines the ec2 instance type. eg. m4-large",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"zones": {
										Description:         "Zones is list of availability zones that can be used.",
										MarkdownDescription: "Zones is list of availability zones that can be used.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": {
								Description:         "Azure is the configuration used when installing on Azure.",
								MarkdownDescription: "Azure is the configuration used when installing on Azure.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"os_disk": {
										Description:         "OSDisk defines the storage for instance.",
										MarkdownDescription: "OSDisk defines the storage for instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"disk_encryption_set": {
												Description:         "DiskEncryptionSet defines a disk encryption set.",
												MarkdownDescription: "DiskEncryptionSet defines a disk encryption set.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the disk encryption set.",
														MarkdownDescription: "Name is the name of the disk encryption set.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"resource_group": {
														Description:         "ResourceGroup defines the Azure resource group used by the disk encryption set.",
														MarkdownDescription: "ResourceGroup defines the Azure resource group used by the disk encryption set.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"subscription_id": {
														Description:         "SubscriptionID defines the Azure subscription the disk encryption set is in.",
														MarkdownDescription: "SubscriptionID defines the Azure subscription the disk encryption set is in.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disk_size_gb": {
												Description:         "DiskSizeGB defines the size of disk in GB.",
												MarkdownDescription: "DiskSizeGB defines the size of disk in GB.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"disk_type": {
												Description:         "DiskType defines the type of disk. For control plane nodes, the valid values are Premium_LRS and StandardSSD_LRS. Default is Premium_LRS.",
												MarkdownDescription: "DiskType defines the type of disk. For control plane nodes, the valid values are Premium_LRS and StandardSSD_LRS. Default is Premium_LRS.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Standard_LRS", "Premium_LRS", "StandardSSD_LRS"),
												},
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"os_image": {
										Description:         "OSImage defines the image to use for the OS.",
										MarkdownDescription: "OSImage defines the image to use for the OS.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"offer": {
												Description:         "Offer is the offer of the image.",
												MarkdownDescription: "Offer is the offer of the image.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"publisher": {
												Description:         "Publisher is the publisher of the image.",
												MarkdownDescription: "Publisher is the publisher of the image.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"sku": {
												Description:         "SKU is the SKU of the image.",
												MarkdownDescription: "SKU is the SKU of the image.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"version": {
												Description:         "Version is the version of the image.",
												MarkdownDescription: "Version is the version of the image.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "InstanceType defines the azure instance type. eg. Standard_DS_V2",
										MarkdownDescription: "InstanceType defines the azure instance type. eg. Standard_DS_V2",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"zones": {
										Description:         "Zones is list of availability zones that can be used. eg. ['1', '2', '3']",
										MarkdownDescription: "Zones is list of availability zones that can be used. eg. ['1', '2', '3']",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gcp": {
								Description:         "GCP is the configuration used when installing on GCP.",
								MarkdownDescription: "GCP is the configuration used when installing on GCP.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"os_disk": {
										Description:         "OSDisk defines the storage for instances.",
										MarkdownDescription: "OSDisk defines the storage for instances.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"disk_size_gb": {
												Description:         "DiskSizeGB defines the size of disk in GB. Defaulted internally to 128.",
												MarkdownDescription: "DiskSizeGB defines the size of disk in GB. Defaulted internally to 128.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(16),

													int64validator.AtMost(65536),
												},
											},

											"disk_type": {
												Description:         "DiskType defines the type of disk. The valid values are pd-standard and pd-ssd. Defaulted internally to pd-ssd.",
												MarkdownDescription: "DiskType defines the type of disk. The valid values are pd-standard and pd-ssd. Defaulted internally to pd-ssd.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("pd-ssd", "pd-standard"),
												},
											},

											"encryption_key": {
												Description:         "EncryptionKey defines the KMS key to be used to encrypt the disk.",
												MarkdownDescription: "EncryptionKey defines the KMS key to be used to encrypt the disk.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"kms_key": {
														Description:         "KMSKey is a reference to a KMS Key to use for the encryption.",
														MarkdownDescription: "KMSKey is a reference to a KMS Key to use for the encryption.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key_ring": {
																Description:         "KeyRing is the name of the KMS Key Ring which the KMS Key belongs to.",
																MarkdownDescription: "KeyRing is the name of the KMS Key Ring which the KMS Key belongs to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"location": {
																Description:         "Location is the GCP location in which the Key Ring exists.",
																MarkdownDescription: "Location is the GCP location in which the Key Ring exists.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
																Description:         "Name is the name of the customer managed encryption key to be used for the disk encryption.",
																MarkdownDescription: "Name is the name of the customer managed encryption key to be used for the disk encryption.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"project_id": {
																Description:         "ProjectID is the ID of the Project in which the KMS Key Ring exists. Defaults to the VM ProjectID if not set.",
																MarkdownDescription: "ProjectID is the ID of the Project in which the KMS Key Ring exists. Defaults to the VM ProjectID if not set.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kms_key_service_account": {
														Description:         "KMSKeyServiceAccount is the service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used. See https://cloud.google.com/compute/docs/access/service-accounts#compute_engine_service_account for details on the default service account.",
														MarkdownDescription: "KMSKeyServiceAccount is the service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used. See https://cloud.google.com/compute/docs/access/service-accounts#compute_engine_service_account for details on the default service account.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "InstanceType defines the GCP instance type. eg. n1-standard-4",
										MarkdownDescription: "InstanceType defines the GCP instance type. eg. n1-standard-4",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"zones": {
										Description:         "Zones is list of availability zones that can be used.",
										MarkdownDescription: "Zones is list of availability zones that can be used.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ibmcloud": {
								Description:         "IBMCloud is the configuration used when installing on IBM Cloud.",
								MarkdownDescription: "IBMCloud is the configuration used when installing on IBM Cloud.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"boot_volume": {
										Description:         "BootVolume is the configuration for the machine's boot volume.",
										MarkdownDescription: "BootVolume is the configuration for the machine's boot volume.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"encryption_key": {
												Description:         "EncryptionKey is the CRN referencing a Key Protect or Hyper Protect Crypto Services key to use for volume encryption. If not specified, a provider managed encryption key will be used.",
												MarkdownDescription: "EncryptionKey is the CRN referencing a Key Protect or Hyper Protect Crypto Services key to use for volume encryption. If not specified, a provider managed encryption key will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dedicated_hosts": {
										Description:         "DedicatedHosts is the configuration for the machine's dedicated host and profile.",
										MarkdownDescription: "DedicatedHosts is the configuration for the machine's dedicated host and profile.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of the dedicated host to provision the machine on. If specified, machines will be created on pre-existing dedicated host.",
												MarkdownDescription: "Name is the name of the dedicated host to provision the machine on. If specified, machines will be created on pre-existing dedicated host.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"profile": {
												Description:         "Profile is the profile ID for the dedicated host. If specified, new dedicated host will be created for machines.",
												MarkdownDescription: "Profile is the profile ID for the dedicated host. If specified, new dedicated host will be created for machines.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "InstanceType is the VSI machine profile.",
										MarkdownDescription: "InstanceType is the VSI machine profile.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"zones": {
										Description:         "Zones is the list of availability zones used for machines in the pool.",
										MarkdownDescription: "Zones is the list of availability zones used for machines in the pool.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"openstack": {
								Description:         "OpenStack is the configuration used when installing on OpenStack.",
								MarkdownDescription: "OpenStack is the configuration used when installing on OpenStack.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"flavor": {
										Description:         "Flavor defines the OpenStack Nova flavor. eg. m1.large The json key here differs from the installer which uses both 'computeFlavor' and type 'type' depending on which type you're looking at, and the resulting field on the MachineSet is 'flavor'. We are opting to stay consistent with the end result.",
										MarkdownDescription: "Flavor defines the OpenStack Nova flavor. eg. m1.large The json key here differs from the installer which uses both 'computeFlavor' and type 'type' depending on which type you're looking at, and the resulting field on the MachineSet is 'flavor'. We are opting to stay consistent with the end result.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"root_volume": {
										Description:         "RootVolume defines the root volume for instances in the machine pool. The instances use ephemeral disks if not set.",
										MarkdownDescription: "RootVolume defines the root volume for instances in the machine pool. The instances use ephemeral disks if not set.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"size": {
												Description:         "Size defines the size of the volume in gibibytes (GiB). Required",
												MarkdownDescription: "Size defines the size of the volume in gibibytes (GiB). Required",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "Type defines the type of the volume. Required",
												MarkdownDescription: "Type defines the type of the volume. Required",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ovirt": {
								Description:         "Ovirt is the configuration used when installing on oVirt.",
								MarkdownDescription: "Ovirt is the configuration used when installing on oVirt.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu": {
										Description:         "CPU defines the VM CPU.",
										MarkdownDescription: "CPU defines the VM CPU.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cores": {
												Description:         "Cores is the number of cores per socket. Total CPUs is (Sockets * Cores)",
												MarkdownDescription: "Cores is the number of cores per socket. Total CPUs is (Sockets * Cores)",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"sockets": {
												Description:         "Sockets is the number of sockets for a VM. Total CPUs is (Sockets * Cores)",
												MarkdownDescription: "Sockets is the number of sockets for a VM. Total CPUs is (Sockets * Cores)",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_mb": {
										Description:         "MemoryMB is the size of a VM's memory in MiBs.",
										MarkdownDescription: "MemoryMB is the size of a VM's memory in MiBs.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"os_disk": {
										Description:         "OSDisk is the the root disk of the node.",
										MarkdownDescription: "OSDisk is the the root disk of the node.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"size_gb": {
												Description:         "SizeGB size of the bootable disk in GiB.",
												MarkdownDescription: "SizeGB size of the bootable disk in GiB.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"vm_type": {
										Description:         "VMType defines the workload type of the VM.",
										MarkdownDescription: "VMType defines the workload type of the VM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("", "desktop", "server", "high_performance"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vsphere": {
								Description:         "VSphere is the configuration used when installing on vSphere",
								MarkdownDescription: "VSphere is the configuration used when installing on vSphere",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cores_per_socket": {
										Description:         "NumCoresPerSocket is the number of cores per socket in a vm. The number of vCPUs on the vm will be NumCPUs/NumCoresPerSocket.",
										MarkdownDescription: "NumCoresPerSocket is the number of cores per socket in a vm. The number of vCPUs on the vm will be NumCPUs/NumCoresPerSocket.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"cpus": {
										Description:         "NumCPUs is the total number of virtual processor cores to assign a vm.",
										MarkdownDescription: "NumCPUs is the total number of virtual processor cores to assign a vm.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"memory_mb": {
										Description:         "Memory is the size of a VM's memory in MB.",
										MarkdownDescription: "Memory is the size of a VM's memory in MB.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"os_disk": {
										Description:         "OSDisk defines the storage for instance.",
										MarkdownDescription: "OSDisk defines the storage for instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"disk_size_gb": {
												Description:         "DiskSizeGB defines the size of disk in GB.",
												MarkdownDescription: "DiskSizeGB defines the size of disk in GB.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"replicas": {
						Description:         "Replicas is the count of machines for this machine pool. Replicas and autoscaling cannot be used together. Default is 1, if autoscaling is not used.",
						MarkdownDescription: "Replicas is the count of machines for this machine pool. Replicas and autoscaling cannot be used together. Default is 1, if autoscaling is not used.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"taints": {
						Description:         "List of taints that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node taints on an ongoing basis.",
						MarkdownDescription: "List of taints that will be applied to the created MachineSet's MachineSpec. This list will overwrite any modifications made to Node taints on an ongoing basis.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"effect": {
								Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
								MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"key": {
								Description:         "Required. The taint key to be applied to a node.",
								MarkdownDescription: "Required. The taint key to be applied to a node.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"time_added": {
								Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
								MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.DateTime64Validator(),
								},
							},

							"value": {
								Description:         "The taint value corresponding to the taint key.",
								MarkdownDescription: "The taint value corresponding to the taint key.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_machine_pool_v1")

	var state HiveOpenshiftIoMachinePoolV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoMachinePoolV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("MachinePool")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_machine_pool_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_machine_pool_v1")

	var state HiveOpenshiftIoMachinePoolV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoMachinePoolV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("MachinePool")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *HiveOpenshiftIoMachinePoolV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_machine_pool_v1")
	// NO-OP: Terraform removes the state automatically for us
}
