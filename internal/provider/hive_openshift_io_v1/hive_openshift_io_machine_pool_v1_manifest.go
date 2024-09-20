/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &HiveOpenshiftIoMachinePoolV1Manifest{}
)

func NewHiveOpenshiftIoMachinePoolV1Manifest() datasource.DataSource {
	return &HiveOpenshiftIoMachinePoolV1Manifest{}
}

type HiveOpenshiftIoMachinePoolV1Manifest struct{}

type HiveOpenshiftIoMachinePoolV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

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
		Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		MachineLabels *map[string]string `tfsdk:"machine_labels" json:"machineLabels,omitempty"`
		Name          *string            `tfsdk:"name" json:"name,omitempty"`
		Platform      *struct {
			Aws *struct {
				AdditionalSecurityGroupIDs *[]string `tfsdk:"additional_security_group_i_ds" json:"additionalSecurityGroupIDs,omitempty"`
				MetadataService            *struct {
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
				Subnets  *[]string          `tfsdk:"subnets" json:"subnets,omitempty"`
				Type     *string            `tfsdk:"type" json:"type,omitempty"`
				UserTags *map[string]string `tfsdk:"user_tags" json:"userTags,omitempty"`
				Zones    *[]string          `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				ComputeSubnet            *string `tfsdk:"compute_subnet" json:"computeSubnet,omitempty"`
				NetworkResourceGroupName *string `tfsdk:"network_resource_group_name" json:"networkResourceGroupName,omitempty"`
				OsDisk                   *struct {
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
				Type             *string   `tfsdk:"type" json:"type,omitempty"`
				VirtualNetwork   *string   `tfsdk:"virtual_network" json:"virtualNetwork,omitempty"`
				VmNetworkingType *string   `tfsdk:"vm_networking_type" json:"vmNetworkingType,omitempty"`
				Zones            *[]string `tfsdk:"zones" json:"zones,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Gcp *struct {
				NetworkProjectID  *string `tfsdk:"network_project_id" json:"networkProjectID,omitempty"`
				OnHostMaintenance *string `tfsdk:"on_host_maintenance" json:"onHostMaintenance,omitempty"`
				OsDisk            *struct {
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
				SecureBoot     *string `tfsdk:"secure_boot" json:"secureBoot,omitempty"`
				ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				Type           *string `tfsdk:"type" json:"type,omitempty"`
				UserTags       *[]struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					ParentID *string `tfsdk:"parent_id" json:"parentID,omitempty"`
					Value    *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"user_tags" json:"userTags,omitempty"`
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
				ResourcePool *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
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

func (r *HiveOpenshiftIoMachinePoolV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_machine_pool_v1_manifest"
}

func (r *HiveOpenshiftIoMachinePoolV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachinePool is the Schema for the machinepools API",
		MarkdownDescription: "MachinePool is the Schema for the machinepools API",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
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
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
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
						Computed:            false,
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
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
						Description:         "Map of label string keys and values that will be applied to the created MachineSet's MachineSpec. This affects the labels that will end up on the *Nodes* (in contrast with the MachineLabels field). This list will overwrite any modifications made to Node labels on an ongoing basis.",
						MarkdownDescription: "Map of label string keys and values that will be applied to the created MachineSet's MachineSpec. This affects the labels that will end up on the *Nodes* (in contrast with the MachineLabels field). This list will overwrite any modifications made to Node labels on an ongoing basis.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"machine_labels": schema.MapAttribute{
						Description:         "Map of label string keys and values that will be applied to the created MachineSet's MachineTemplateSpec. This affects the labels that will end up on the *Machines* (in contrast with the Labels field). This list will overwrite any modifications made to Machine labels on an ongoing basis. Note: We ignore entries that conflict with generated labels.",
						MarkdownDescription: "Map of label string keys and values that will be applied to the created MachineSet's MachineTemplateSpec. This affects the labels that will end up on the *Machines* (in contrast with the Labels field). This list will overwrite any modifications made to Machine labels on an ongoing basis. Note: We ignore entries that conflict with generated labels.",
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
						Description:         "Platform is configuration for machine pool specific to the platform. When using a MachinePool to control the default worker machines created by installer, these must match the values provided in the install-config.",
						MarkdownDescription: "Platform is configuration for machine pool specific to the platform. When using a MachinePool to control the default worker machines created by installer, these must match the values provided in the install-config.",
						Attributes: map[string]schema.Attribute{
							"aws": schema.SingleNestedAttribute{
								Description:         "AWS is the configuration used when installing on AWS.",
								MarkdownDescription: "AWS is the configuration used when installing on AWS.",
								Attributes: map[string]schema.Attribute{
									"additional_security_group_i_ds": schema.ListAttribute{
										Description:         "AdditionalSecurityGroupIDs contains IDs of additional security groups for machines, where each ID is presented in the format sg-xxxx.",
										MarkdownDescription: "AdditionalSecurityGroupIDs contains IDs of additional security groups for machines, where each ID is presented in the format sg-xxxx.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

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
										Description:         "Subnets is the list of IDs of subnets to which to attach the machines. There must be exactly one subnet for each availability zone used. These subnets may be public or private. As a special case, for consistency with install-config, you may specify exactly one private and one public subnet for each availability zone. In this case, the public subnets will be filtered out and only the private subnets will be used. If empty/omitted, we will look for subnets in each availability zone tagged with Name=<clusterID>-private-<az> (legacy terraform) or <clusterID>-subnet-private-<az> (CAPA).",
										MarkdownDescription: "Subnets is the list of IDs of subnets to which to attach the machines. There must be exactly one subnet for each availability zone used. These subnets may be public or private. As a special case, for consistency with install-config, you may specify exactly one private and one public subnet for each availability zone. In this case, the public subnets will be filtered out and only the private subnets will be used. If empty/omitted, we will look for subnets in each availability zone tagged with Name=<clusterID>-private-<az> (legacy terraform) or <clusterID>-subnet-private-<az> (CAPA).",
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

									"user_tags": schema.MapAttribute{
										Description:         "UserTags contains the user defined tags to be supplied for the ec2 instance. Note that these will be merged with ClusterDeployment.Spec.Platform.AWS.UserTags, with this field taking precedence when keys collide.",
										MarkdownDescription: "UserTags contains the user defined tags to be supplied for the ec2 instance. Note that these will be merged with ClusterDeployment.Spec.Platform.AWS.UserTags, with this field taking precedence when keys collide.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
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
									"compute_subnet": schema.StringAttribute{
										Description:         "ComputeSubnet specifies an existing subnet for use by compute nodes. If omitted, the default (${infraID}-worker-subnet) will be used.",
										MarkdownDescription: "ComputeSubnet specifies an existing subnet for use by compute nodes. If omitted, the default (${infraID}-worker-subnet) will be used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"network_resource_group_name": schema.StringAttribute{
										Description:         "NetworkResourceGroupName specifies the network resource group that contains an existing VNet. Ignored unless VirtualNetwork is also specified.",
										MarkdownDescription: "NetworkResourceGroupName specifies the network resource group that contains an existing VNet. Ignored unless VirtualNetwork is also specified.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

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

									"virtual_network": schema.StringAttribute{
										Description:         "VirtualNetwork specifies the name of an existing VNet for the Machines to use If omitted, the default (${infraID}-vnet) will be used.",
										MarkdownDescription: "VirtualNetwork specifies the name of an existing VNet for the Machines to use If omitted, the default (${infraID}-vnet) will be used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vm_networking_type": schema.StringAttribute{
										Description:         "VMNetworkingType specifies whether to enable accelerated networking. Accelerated networking enables single root I/O virtualization (SR-IOV) to a VM, greatly improving its networking performance. eg. values: 'Accelerated', 'Basic'",
										MarkdownDescription: "VMNetworkingType specifies whether to enable accelerated networking. Accelerated networking enables single root I/O virtualization (SR-IOV) to a VM, greatly improving its networking performance. eg. values: 'Accelerated', 'Basic'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Accelerated", "Basic"),
										},
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
									"network_project_id": schema.StringAttribute{
										Description:         "NetworkProjectID specifies which project the network and subnets exist in when they are not in the main ProjectID.",
										MarkdownDescription: "NetworkProjectID specifies which project the network and subnets exist in when they are not in the main ProjectID.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"on_host_maintenance": schema.StringAttribute{
										Description:         "OnHostMaintenance determines the behavior when a maintenance event occurs that might cause the instance to reboot. This is required to be set to 'Terminate' if you want to provision machine with attached GPUs. Otherwise, allowed values are 'Migrate' and 'Terminate'. If omitted, the platform chooses a default, which is subject to change over time, currently that default is 'Migrate'.",
										MarkdownDescription: "OnHostMaintenance determines the behavior when a maintenance event occurs that might cause the instance to reboot. This is required to be set to 'Terminate' if you want to provision machine with attached GPUs. Otherwise, allowed values are 'Migrate' and 'Terminate'. If omitted, the platform chooses a default, which is subject to change over time, currently that default is 'Migrate'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Migrate", "Terminate"),
										},
									},

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

									"secure_boot": schema.StringAttribute{
										Description:         "SecureBoot Defines whether the instance should have secure boot enabled. Verifies the digital signature of all boot components, and halts the boot process if signature verification fails. If omitted, the platform chooses a default, which is subject to change over time. Currently that default is 'Disabled'.",
										MarkdownDescription: "SecureBoot Defines whether the instance should have secure boot enabled. Verifies the digital signature of all boot components, and halts the boot process if signature verification fails. If omitted, the platform chooses a default, which is subject to change over time. Currently that default is 'Disabled'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Enabled", "Disabled"),
										},
									},

									"service_account": schema.StringAttribute{
										Description:         "ServiceAccount is the email of a gcp service account to be attached to worker nodes in order to provide the permissions required by the cloud provider. For the default worker MachinePool, it is the user's responsibility to match this to the value provided in the install-config.",
										MarkdownDescription: "ServiceAccount is the email of a gcp service account to be attached to worker nodes in order to provide the permissions required by the cloud provider. For the default worker MachinePool, it is the user's responsibility to match this to the value provided in the install-config.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "InstanceType defines the GCP instance type. eg. n1-standard-4",
										MarkdownDescription: "InstanceType defines the GCP instance type. eg. n1-standard-4",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"user_tags": schema.ListNestedAttribute{
										Description:         "userTags has additional keys and values that we will add as tags to the providerSpec of MachineSets that we creates on GCP. Tag key and tag value should be the shortnames of the tag key and tag value resource. Consumer is responsible for using this only for spokes where custom tags are supported.",
										MarkdownDescription: "userTags has additional keys and values that we will add as tags to the providerSpec of MachineSets that we creates on GCP. Tag key and tag value should be the shortnames of the tag key and tag value resource. Consumer is responsible for using this only for spokes where custom tags are supported.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the key part of the tag. A tag key can have a maximum of 63 characters and cannot be empty. Tag key must begin and end with an alphanumeric character, and must contain only uppercase, lowercase alphanumeric characters, and the following special characters '._-'.",
													MarkdownDescription: "key is the key part of the tag. A tag key can have a maximum of 63 characters and cannot be empty. Tag key must begin and end with an alphanumeric character, and must contain only uppercase, lowercase alphanumeric characters, and the following special characters '._-'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"parent_id": schema.StringAttribute{
													Description:         "parentID is the ID of the hierarchical resource where the tags are defined, e.g. at the Organization or the Project level. To find the Organization ID or Project ID refer to the following pages: https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id, https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects. An OrganizationID must consist of decimal numbers, and cannot have leading zeroes. A ProjectID must be 6 to 30 characters in length, can only contain lowercase letters, numbers, and hyphens, and must start with a letter, and cannot end with a hyphen.",
													MarkdownDescription: "parentID is the ID of the hierarchical resource where the tags are defined, e.g. at the Organization or the Project level. To find the Organization ID or Project ID refer to the following pages: https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id, https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects. An OrganizationID must consist of decimal numbers, and cannot have leading zeroes. A ProjectID must be 6 to 30 characters in length, can only contain lowercase letters, numbers, and hyphens, and must start with a letter, and cannot end with a hyphen.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "value is the value part of the tag. A tag value can have a maximum of 63 characters and cannot be empty. Tag value must begin and end with an alphanumeric character, and must contain only uppercase, lowercase alphanumeric characters, and the following special characters '_-.@%=+:,*#&(){}[]' and spaces.",
													MarkdownDescription: "value is the value part of the tag. A tag value can have a maximum of 63 characters and cannot be empty. Tag value must begin and end with an alphanumeric character, and must contain only uppercase, lowercase alphanumeric characters, and the following special characters '_-.@%=+:,*#&(){}[]' and spaces.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
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

									"resource_pool": schema.StringAttribute{
										Description:         "ResourcePool is the name of the resource pool that will be used for virtual machines. If it is not present, a default value will be used.",
										MarkdownDescription: "ResourcePool is the name of the resource pool that will be used for virtual machines. If it is not present, a default value will be used.",
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

func (r *HiveOpenshiftIoMachinePoolV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_machine_pool_v1_manifest")

	var model HiveOpenshiftIoMachinePoolV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("MachinePool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
