/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package eks_services_k8s_aws_v1alpha1

import (
	"context"
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
	_ datasource.DataSource = &EksServicesK8SAwsNodegroupV1Alpha1Manifest{}
)

func NewEksServicesK8SAwsNodegroupV1Alpha1Manifest() datasource.DataSource {
	return &EksServicesK8SAwsNodegroupV1Alpha1Manifest{}
}

type EksServicesK8SAwsNodegroupV1Alpha1Manifest struct{}

type EksServicesK8SAwsNodegroupV1Alpha1ManifestData struct {
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
		AmiType            *string `tfsdk:"ami_type" json:"amiType,omitempty"`
		CapacityType       *string `tfsdk:"capacity_type" json:"capacityType,omitempty"`
		ClientRequestToken *string `tfsdk:"client_request_token" json:"clientRequestToken,omitempty"`
		ClusterName        *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterRef         *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		DiskSize       *int64             `tfsdk:"disk_size" json:"diskSize,omitempty"`
		InstanceTypes  *[]string          `tfsdk:"instance_types" json:"instanceTypes,omitempty"`
		Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		LaunchTemplate *struct {
			Id      *string `tfsdk:"id" json:"id,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"launch_template" json:"launchTemplate,omitempty"`
		Name        *string `tfsdk:"name" json:"name,omitempty"`
		NodeRole    *string `tfsdk:"node_role" json:"nodeRole,omitempty"`
		NodeRoleRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"node_role_ref" json:"nodeRoleRef,omitempty"`
		ReleaseVersion *string `tfsdk:"release_version" json:"releaseVersion,omitempty"`
		RemoteAccess   *struct {
			Ec2SshKey               *string `tfsdk:"ec2_ssh_key" json:"ec2SshKey,omitempty"`
			SourceSecurityGroupRefs *[]struct {
				From *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"source_security_group_refs" json:"sourceSecurityGroupRefs,omitempty"`
			SourceSecurityGroups *[]string `tfsdk:"source_security_groups" json:"sourceSecurityGroups,omitempty"`
		} `tfsdk:"remote_access" json:"remoteAccess,omitempty"`
		ScalingConfig *struct {
			DesiredSize *int64 `tfsdk:"desired_size" json:"desiredSize,omitempty"`
			MaxSize     *int64 `tfsdk:"max_size" json:"maxSize,omitempty"`
			MinSize     *int64 `tfsdk:"min_size" json:"minSize,omitempty"`
		} `tfsdk:"scaling_config" json:"scalingConfig,omitempty"`
		SubnetRefs *[]struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"subnet_refs" json:"subnetRefs,omitempty"`
		Subnets *[]string          `tfsdk:"subnets" json:"subnets,omitempty"`
		Tags    *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		Taints  *[]struct {
			Effect *string `tfsdk:"effect" json:"effect,omitempty"`
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			Value  *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
		UpdateConfig *struct {
			MaxUnavailable           *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			MaxUnavailablePercentage *int64 `tfsdk:"max_unavailable_percentage" json:"maxUnavailablePercentage,omitempty"`
		} `tfsdk:"update_config" json:"updateConfig,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_eks_services_k8s_aws_nodegroup_v1alpha1_manifest"
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Nodegroup is the Schema for the Nodegroups API",
		MarkdownDescription: "Nodegroup is the Schema for the Nodegroups API",
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
				Description:         "NodegroupSpec defines the desired state of Nodegroup.An object representing an Amazon EKS managed node group.",
				MarkdownDescription: "NodegroupSpec defines the desired state of Nodegroup.An object representing an Amazon EKS managed node group.",
				Attributes: map[string]schema.Attribute{
					"ami_type": schema.StringAttribute{
						Description:         "The AMI type for your node group. If you specify launchTemplate, and yourlaunch template uses a custom AMI, then don't specify amiType, or the nodegroup deployment will fail. If your launch template uses a Windows customAMI, then add eks:kube-proxy-windows to your Windows nodes rolearn in theaws-auth ConfigMap. For more information about using launch templates withAmazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The AMI type for your node group. If you specify launchTemplate, and yourlaunch template uses a custom AMI, then don't specify amiType, or the nodegroup deployment will fail. If your launch template uses a Windows customAMI, then add eks:kube-proxy-windows to your Windows nodes rolearn in theaws-auth ConfigMap. For more information about using launch templates withAmazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"capacity_type": schema.StringAttribute{
						Description:         "The capacity type for your node group.",
						MarkdownDescription: "The capacity type for your node group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_request_token": schema.StringAttribute{
						Description:         "A unique, case-sensitive identifier that you provide to ensure the idempotencyof the request.",
						MarkdownDescription: "A unique, case-sensitive identifier that you provide to ensure the idempotencyof the request.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "The name of your cluster.",
						MarkdownDescription: "The name of your cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"disk_size": schema.Int64Attribute{
						Description:         "The root device disk size (in GiB) for your node group instances. The defaultdisk size is 20 GiB for Linux and Bottlerocket. The default disk size is50 GiB for Windows. If you specify launchTemplate, then don't specify diskSize,or the node group deployment will fail. For more information about usinglaunch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The root device disk size (in GiB) for your node group instances. The defaultdisk size is 20 GiB for Linux and Bottlerocket. The default disk size is50 GiB for Windows. If you specify launchTemplate, then don't specify diskSize,or the node group deployment will fail. For more information about usinglaunch templates with Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_types": schema.ListAttribute{
						Description:         "Specify the instance types for a node group. If you specify a GPU instancetype, make sure to also specify an applicable GPU AMI type with the amiTypeparameter. If you specify launchTemplate, then you can specify zero or oneinstance type in your launch template or you can specify 0-20 instance typesfor instanceTypes. If however, you specify an instance type in your launchtemplate and specify any instanceTypes, the node group deployment will fail.If you don't specify an instance type in a launch template or for instanceTypes,then t3.medium is used, by default. If you specify Spot for capacityType,then we recommend specifying multiple values for instanceTypes. For moreinformation, see Managed node group capacity types (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types)and Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "Specify the instance types for a node group. If you specify a GPU instancetype, make sure to also specify an applicable GPU AMI type with the amiTypeparameter. If you specify launchTemplate, then you can specify zero or oneinstance type in your launch template or you can specify 0-20 instance typesfor instanceTypes. If however, you specify an instance type in your launchtemplate and specify any instanceTypes, the node group deployment will fail.If you don't specify an instance type in a launch template or for instanceTypes,then t3.medium is used, by default. If you specify Spot for capacityType,then we recommend specifying multiple values for instanceTypes. For moreinformation, see Managed node group capacity types (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html#managed-node-group-capacity-types)and Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"labels": schema.MapAttribute{
						Description:         "The Kubernetes labels to apply to the nodes in the node group when they arecreated.",
						MarkdownDescription: "The Kubernetes labels to apply to the nodes in the node group when they arecreated.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"launch_template": schema.SingleNestedAttribute{
						Description:         "An object representing a node group's launch template specification. If specified,then do not specify instanceTypes, diskSize, or remoteAccess and make surethat the launch template meets the requirements in launchTemplateSpecification.",
						MarkdownDescription: "An object representing a node group's launch template specification. If specified,then do not specify instanceTypes, diskSize, or remoteAccess and make surethat the launch template meets the requirements in launchTemplateSpecification.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "The unique name to give your node group.",
						MarkdownDescription: "The unique name to give your node group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"node_role": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the IAM role to associate with your nodegroup. The Amazon EKS worker node kubelet daemon makes calls to Amazon WebServices APIs on your behalf. Nodes receive permissions for these API callsthrough an IAM instance profile and associated policies. Before you can launchnodes and register them into a cluster, you must create an IAM role for thosenodes to use when they are launched. For more information, see Amazon EKSnode IAM role (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)in the Amazon EKS User Guide . If you specify launchTemplate, then don'tspecify IamInstanceProfile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html)in your launch template, or the node group deployment will fail. For moreinformation about using launch templates with Amazon EKS, see Launch templatesupport (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM role to associate with your nodegroup. The Amazon EKS worker node kubelet daemon makes calls to Amazon WebServices APIs on your behalf. Nodes receive permissions for these API callsthrough an IAM instance profile and associated policies. Before you can launchnodes and register them into a cluster, you must create an IAM role for thosenodes to use when they are launched. For more information, see Amazon EKSnode IAM role (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)in the Amazon EKS User Guide . If you specify launchTemplate, then don'tspecify IamInstanceProfile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html)in your launch template, or the node group deployment will fail. For moreinformation about using launch templates with Amazon EKS, see Launch templatesupport (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_role_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"release_version": schema.StringAttribute{
						Description:         "The AMI version of the Amazon EKS optimized AMI to use with your node group.By default, the latest available AMI version for the node group's currentKubernetes version is used. For information about Linux versions, see AmazonEKS optimized Amazon Linux AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)in the Amazon EKS User Guide. Amazon EKS managed node groups support theNovember 2022 and later releases of the Windows AMIs. For information aboutWindows versions, see Amazon EKS optimized Windows AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-versions-windows.html)in the Amazon EKS User Guide.If you specify launchTemplate, and your launch template uses a custom AMI,then don't specify releaseVersion, or the node group deployment will fail.For more information about using launch templates with Amazon EKS, see Launchtemplate support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The AMI version of the Amazon EKS optimized AMI to use with your node group.By default, the latest available AMI version for the node group's currentKubernetes version is used. For information about Linux versions, see AmazonEKS optimized Amazon Linux AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)in the Amazon EKS User Guide. Amazon EKS managed node groups support theNovember 2022 and later releases of the Windows AMIs. For information aboutWindows versions, see Amazon EKS optimized Windows AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-versions-windows.html)in the Amazon EKS User Guide.If you specify launchTemplate, and your launch template uses a custom AMI,then don't specify releaseVersion, or the node group deployment will fail.For more information about using launch templates with Amazon EKS, see Launchtemplate support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remote_access": schema.SingleNestedAttribute{
						Description:         "The remote access configuration to use with your node group. For Linux, theprotocol is SSH. For Windows, the protocol is RDP. If you specify launchTemplate,then don't specify remoteAccess, or the node group deployment will fail.For more information about using launch templates with Amazon EKS, see Launchtemplate support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The remote access configuration to use with your node group. For Linux, theprotocol is SSH. For Windows, the protocol is RDP. If you specify launchTemplate,then don't specify remoteAccess, or the node group deployment will fail.For more information about using launch templates with Amazon EKS, see Launchtemplate support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						Attributes: map[string]schema.Attribute{
							"ec2_ssh_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_security_group_refs": schema.ListNestedAttribute{
								Description:         "Reference field for SourceSecurityGroups",
								MarkdownDescription: "Reference field for SourceSecurityGroups",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from": schema.SingleNestedAttribute{
											Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
											MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_security_groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"scaling_config": schema.SingleNestedAttribute{
						Description:         "The scaling configuration details for the Auto Scaling group that is createdfor your node group.",
						MarkdownDescription: "The scaling configuration details for the Auto Scaling group that is createdfor your node group.",
						Attributes: map[string]schema.Attribute{
							"desired_size": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_size": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_size": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnet_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnets": schema.ListAttribute{
						Description:         "The subnets to use for the Auto Scaling group that is created for your nodegroup. If you specify launchTemplate, then don't specify SubnetId (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html)in your launch template, or the node group deployment will fail. For moreinformation about using launch templates with Amazon EKS, see Launch templatesupport (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The subnets to use for the Auto Scaling group that is created for your nodegroup. If you specify launchTemplate, then don't specify SubnetId (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html)in your launch template, or the node group deployment will fail. For moreinformation about using launch templates with Amazon EKS, see Launch templatesupport (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.MapAttribute{
						Description:         "Metadata that assists with categorization and organization. Each tag consistsof a key and an optional value. You define both. Tags don't propagate toany other cluster or Amazon Web Services resources.",
						MarkdownDescription: "Metadata that assists with categorization and organization. Each tag consistsof a key and an optional value. You define both. Tags don't propagate toany other cluster or Amazon Web Services resources.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"taints": schema.ListNestedAttribute{
						Description:         "The Kubernetes taints to be applied to the nodes in the node group. For moreinformation, see Node taints on managed node groups (https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html).",
						MarkdownDescription: "The Kubernetes taints to be applied to the nodes in the node group. For moreinformation, see Node taints on managed node groups (https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"update_config": schema.SingleNestedAttribute{
						Description:         "The node group update configuration.",
						MarkdownDescription: "The node group update configuration.",
						Attributes: map[string]schema.Attribute{
							"max_unavailable": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_unavailable_percentage": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": schema.StringAttribute{
						Description:         "The Kubernetes version to use for your managed nodes. By default, the Kubernetesversion of the cluster is used, and this is the only accepted specified value.If you specify launchTemplate, and your launch template uses a custom AMI,then don't specify version, or the node group deployment will fail. For moreinformation about using launch templates with Amazon EKS, see Launch templatesupport (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
						MarkdownDescription: "The Kubernetes version to use for your managed nodes. By default, the Kubernetesversion of the cluster is used, and this is the only accepted specified value.If you specify launchTemplate, and your launch template uses a custom AMI,then don't specify version, or the node group deployment will fail. For moreinformation about using launch templates with Amazon EKS, see Launch templatesupport (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)in the Amazon EKS User Guide.",
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
	}
}

func (r *EksServicesK8SAwsNodegroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_eks_services_k8s_aws_nodegroup_v1alpha1_manifest")

	var model EksServicesK8SAwsNodegroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("eks.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Nodegroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
