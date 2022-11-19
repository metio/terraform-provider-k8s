/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type EksServicesK8SAwsNodegroupV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*EksServicesK8SAwsNodegroupV1Alpha1Resource)(nil)
)

type EksServicesK8SAwsNodegroupV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type EksServicesK8SAwsNodegroupV1Alpha1GoModel struct {
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
		AmiType *string `tfsdk:"ami_type" yaml:"amiType,omitempty"`

		CapacityType *string `tfsdk:"capacity_type" yaml:"capacityType,omitempty"`

		ClientRequestToken *string `tfsdk:"client_request_token" yaml:"clientRequestToken,omitempty"`

		ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

		ClusterRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"cluster_ref" yaml:"clusterRef,omitempty"`

		DiskSize *int64 `tfsdk:"disk_size" yaml:"diskSize,omitempty"`

		InstanceTypes *[]string `tfsdk:"instance_types" yaml:"instanceTypes,omitempty"`

		Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

		LaunchTemplate *struct {
			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"launch_template" yaml:"launchTemplate,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		NodeRole *string `tfsdk:"node_role" yaml:"nodeRole,omitempty"`

		NodeRoleRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"node_role_ref" yaml:"nodeRoleRef,omitempty"`

		ReleaseVersion *string `tfsdk:"release_version" yaml:"releaseVersion,omitempty"`

		RemoteAccess *struct {
			Ec2SshKey *string `tfsdk:"ec2_ssh_key" yaml:"ec2SshKey,omitempty"`

			SourceSecurityGroupRefs *[]struct {
				From *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"from" yaml:"from,omitempty"`
			} `tfsdk:"source_security_group_refs" yaml:"sourceSecurityGroupRefs,omitempty"`

			SourceSecurityGroups *[]string `tfsdk:"source_security_groups" yaml:"sourceSecurityGroups,omitempty"`
		} `tfsdk:"remote_access" yaml:"remoteAccess,omitempty"`

		ScalingConfig *struct {
			DesiredSize *int64 `tfsdk:"desired_size" yaml:"desiredSize,omitempty"`

			MaxSize *int64 `tfsdk:"max_size" yaml:"maxSize,omitempty"`

			MinSize *int64 `tfsdk:"min_size" yaml:"minSize,omitempty"`
		} `tfsdk:"scaling_config" yaml:"scalingConfig,omitempty"`

		SubnetRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"subnet_refs" yaml:"subnetRefs,omitempty"`

		Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`

		Tags *map[string]string `tfsdk:"tags" yaml:"tags,omitempty"`

		Taints *[]struct {
			Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"taints" yaml:"taints,omitempty"`

		UpdateConfig *struct {
			MaxUnavailable *int64 `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

			MaxUnavailablePercentage *int64 `tfsdk:"max_unavailable_percentage" yaml:"maxUnavailablePercentage,omitempty"`
		} `tfsdk:"update_config" yaml:"updateConfig,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewEksServicesK8SAwsNodegroupV1Alpha1Resource() resource.Resource {
	return &EksServicesK8SAwsNodegroupV1Alpha1Resource{}
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_eks_services_k8s_aws_nodegroup_v1alpha1"
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Nodegroup is the Schema for the Nodegroups API",
		MarkdownDescription: "Nodegroup is the Schema for the Nodegroups API",
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
				Description:         "NodegroupSpec defines the desired state of Nodegroup.  An object representing an Amazon EKS managed node group.",
				MarkdownDescription: "NodegroupSpec defines the desired state of Nodegroup.  An object representing an Amazon EKS managed node group.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"ami_type": {
						Description:         "The AMI type for your node group. GPU instance types should use the AL2_x86_64_GPU AMI type. Non-GPU instances should use the AL2_x86_64 AMI type. Arm instances should use the AL2_ARM_64 AMI type. All types use the Amazon EKS optimized Amazon Linux 2 AMI. If you specify launchTemplate, and your launch template uses a custom AMI, then don't specify amiType, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The AMI type for your node group. GPU instance types should use the AL2_x86_64_GPU AMI type. Non-GPU instances should use the AL2_x86_64 AMI type. Arm instances should use the AL2_ARM_64 AMI type. All types use the Amazon EKS optimized Amazon Linux 2 AMI. If you specify launchTemplate, and your launch template uses a custom AMI, then don't specify amiType, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"capacity_type": {
						Description:         "The capacity type for your node group.",
						MarkdownDescription: "The capacity type for your node group.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"client_request_token": {
						Description:         "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",
						MarkdownDescription: "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_name": {
						Description:         "The name of the cluster to create the node group in.",
						MarkdownDescription: "The name of the cluster to create the node group in.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

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

					"disk_size": {
						Description:         "The root device disk size (in GiB) for your node group instances. The default disk size is 20 GiB. If you specify launchTemplate, then don't specify diskSize, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The root device disk size (in GiB) for your node group instances. The default disk size is 20 GiB. If you specify launchTemplate, then don't specify diskSize, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance_types": {
						Description:         "Specify the instance types for a node group. If you specify a GPU instance type, be sure to specify AL2_x86_64_GPU with the amiType parameter. If you specify launchTemplate, then you can specify zero or one instance type in your launch template or you can specify 0-20 instance types for instanceTypes. If however, you specify an instance type in your launch template and specify any instanceTypes, the node group deployment will fail. If you don't specify an instance type in a launch template or for instanceTypes, then t3.medium is used, by default. If you specify Spot for capacityType, then we recommend specifying multiple values for instanceTypes. For more information, see Managed node group capacity types (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "Specify the instance types for a node group. If you specify a GPU instance type, be sure to specify AL2_x86_64_GPU with the amiType parameter. If you specify launchTemplate, then you can specify zero or one instance type in your launch template or you can specify 0-20 instance types for instanceTypes. If however, you specify an instance type in your launch template and specify any instanceTypes, the node group deployment will fail. If you don't specify an instance type in a launch template or for instanceTypes, then t3.medium is used, by default. If you specify Spot for capacityType, then we recommend specifying multiple values for instanceTypes. For more information, see Managed node group capacity types (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types) and Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"labels": {
						Description:         "The Kubernetes labels to be applied to the nodes in the node group when they are created.",
						MarkdownDescription: "The Kubernetes labels to be applied to the nodes in the node group when they are created.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"launch_template": {
						Description:         "An object representing a node group's launch template specification. If specified, then do not specify instanceTypes, diskSize, or remoteAccess and make sure that the launch template meets the requirements in launchTemplateSpecification.",
						MarkdownDescription: "An object representing a node group's launch template specification. If specified, then do not specify instanceTypes, diskSize, or remoteAccess and make sure that the launch template meets the requirements in launchTemplateSpecification.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "",
								MarkdownDescription: "",

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

					"name": {
						Description:         "The unique name to give your node group.",
						MarkdownDescription: "The unique name to give your node group.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"node_role": {
						Description:         "The Amazon Resource Name (ARN) of the IAM role to associate with your node group. The Amazon EKS worker node kubelet daemon makes calls to Amazon Web Services APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see Amazon EKS node IAM role (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the Amazon EKS User Guide . If you specify launchTemplate, then don't specify IamInstanceProfile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM role to associate with your node group. The Amazon EKS worker node kubelet daemon makes calls to Amazon Web Services APIs on your behalf. Nodes receive permissions for these API calls through an IAM instance profile and associated policies. Before you can launch nodes and register them into a cluster, you must create an IAM role for those nodes to use when they are launched. For more information, see Amazon EKS node IAM role (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the Amazon EKS User Guide . If you specify launchTemplate, then don't specify IamInstanceProfile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_role_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

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

					"release_version": {
						Description:         "The AMI version of the Amazon EKS optimized AMI to use with your node group. By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see Amazon EKS optimized Amazon Linux 2 AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the Amazon EKS User Guide. If you specify launchTemplate, and your launch template uses a custom AMI, then don't specify releaseVersion, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The AMI version of the Amazon EKS optimized AMI to use with your node group. By default, the latest available AMI version for the node group's current Kubernetes version is used. For more information, see Amazon EKS optimized Amazon Linux 2 AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html) in the Amazon EKS User Guide. If you specify launchTemplate, and your launch template uses a custom AMI, then don't specify releaseVersion, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_access": {
						Description:         "The remote access (SSH) configuration to use with your node group. If you specify launchTemplate, then don't specify remoteAccess, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The remote access (SSH) configuration to use with your node group. If you specify launchTemplate, then don't specify remoteAccess, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ec2_ssh_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_security_group_refs": {
								Description:         "Reference field for SourceSecurityGroups",
								MarkdownDescription: "Reference field for SourceSecurityGroups",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"from": {
										Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
										MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

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

							"source_security_groups": {
								Description:         "",
								MarkdownDescription: "",

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

					"scaling_config": {
						Description:         "The scaling configuration details for the Auto Scaling group that is created for your node group.",
						MarkdownDescription: "The scaling configuration details for the Auto Scaling group that is created for your node group.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"desired_size": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_size": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"min_size": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnet_refs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

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

					"subnets": {
						Description:         "The subnets to use for the Auto Scaling group that is created for your node group. If you specify launchTemplate, then don't specify SubnetId (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The subnets to use for the Auto Scaling group that is created for your node group. If you specify launchTemplate, then don't specify SubnetId (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) in your launch template, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "The metadata to apply to the node group to assist with categorization and organization. Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.",
						MarkdownDescription: "The metadata to apply to the node group to assist with categorization and organization. Each tag consists of a key and an optional value. You define both. Node group tags do not propagate to any other resources associated with the node group, such as the Amazon EC2 instances or subnets.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"taints": {
						Description:         "The Kubernetes taints to be applied to the nodes in the node group. For more information, see Node taints on managed node groups (https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html).",
						MarkdownDescription: "The Kubernetes taints to be applied to the nodes in the node group. For more information, see Node taints on managed node groups (https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html).",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"effect": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

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

					"update_config": {
						Description:         "The node group update configuration.",
						MarkdownDescription: "The node group update configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_unavailable": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_unavailable_percentage": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "The Kubernetes version to use for your managed nodes. By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify launchTemplate, and your launch template uses a custom AMI, then don't specify version, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The Kubernetes version to use for your managed nodes. By default, the Kubernetes version of the cluster is used, and this is the only accepted specified value. If you specify launchTemplate, and your launch template uses a custom AMI, then don't specify version, or the node group deployment will fail. For more information about using launch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the Amazon EKS User Guide.",

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
		},
	}, nil
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_eks_services_k8s_aws_nodegroup_v1alpha1")

	var state EksServicesK8SAwsNodegroupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EksServicesK8SAwsNodegroupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("eks.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Nodegroup")

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

func (r *EksServicesK8SAwsNodegroupV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_eks_services_k8s_aws_nodegroup_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_eks_services_k8s_aws_nodegroup_v1alpha1")

	var state EksServicesK8SAwsNodegroupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EksServicesK8SAwsNodegroupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("eks.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Nodegroup")

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

func (r *EksServicesK8SAwsNodegroupV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_eks_services_k8s_aws_nodegroup_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
