/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package eks_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &EksServicesK8SAwsClusterV1Alpha1Manifest{}
)

func NewEksServicesK8SAwsClusterV1Alpha1Manifest() datasource.DataSource {
	return &EksServicesK8SAwsClusterV1Alpha1Manifest{}
}

type EksServicesK8SAwsClusterV1Alpha1Manifest struct{}

type EksServicesK8SAwsClusterV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AccessConfig *struct {
			AuthenticationMode                      *string `tfsdk:"authentication_mode" json:"authenticationMode,omitempty"`
			BootstrapClusterCreatorAdminPermissions *bool   `tfsdk:"bootstrap_cluster_creator_admin_permissions" json:"bootstrapClusterCreatorAdminPermissions,omitempty"`
		} `tfsdk:"access_config" json:"accessConfig,omitempty"`
		ClientRequestToken *string `tfsdk:"client_request_token" json:"clientRequestToken,omitempty"`
		EncryptionConfig   *[]struct {
			Provider *struct {
				KeyARN *string `tfsdk:"key_arn" json:"keyARN,omitempty"`
				KeyRef *struct {
					From *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"from" json:"from,omitempty"`
				} `tfsdk:"key_ref" json:"keyRef,omitempty"`
			} `tfsdk:"provider" json:"provider,omitempty"`
			Resources *[]string `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"encryption_config" json:"encryptionConfig,omitempty"`
		KubernetesNetworkConfig *struct {
			IpFamily        *string `tfsdk:"ip_family" json:"ipFamily,omitempty"`
			ServiceIPv4CIDR *string `tfsdk:"service_i_pv4_cidr" json:"serviceIPv4CIDR,omitempty"`
		} `tfsdk:"kubernetes_network_config" json:"kubernetesNetworkConfig,omitempty"`
		Logging *struct {
			ClusterLogging *[]struct {
				Enabled *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				Types   *[]string `tfsdk:"types" json:"types,omitempty"`
			} `tfsdk:"cluster_logging" json:"clusterLogging,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		Name          *string `tfsdk:"name" json:"name,omitempty"`
		OutpostConfig *struct {
			ControlPlaneInstanceType *string `tfsdk:"control_plane_instance_type" json:"controlPlaneInstanceType,omitempty"`
			ControlPlanePlacement    *struct {
				GroupName *string `tfsdk:"group_name" json:"groupName,omitempty"`
			} `tfsdk:"control_plane_placement" json:"controlPlanePlacement,omitempty"`
			OutpostARNs *[]string `tfsdk:"outpost_ar_ns" json:"outpostARNs,omitempty"`
		} `tfsdk:"outpost_config" json:"outpostConfig,omitempty"`
		ResourcesVPCConfig *struct {
			EndpointPrivateAccess *bool     `tfsdk:"endpoint_private_access" json:"endpointPrivateAccess,omitempty"`
			EndpointPublicAccess  *bool     `tfsdk:"endpoint_public_access" json:"endpointPublicAccess,omitempty"`
			PublicAccessCIDRs     *[]string `tfsdk:"public_access_cid_rs" json:"publicAccessCIDRs,omitempty"`
			SecurityGroupIDs      *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
			SecurityGroupRefs     *[]struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"security_group_refs" json:"securityGroupRefs,omitempty"`
			SubnetIDs  *[]string `tfsdk:"subnet_i_ds" json:"subnetIDs,omitempty"`
			SubnetRefs *[]struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"subnet_refs" json:"subnetRefs,omitempty"`
		} `tfsdk:"resources_vpc_config" json:"resourcesVPCConfig,omitempty"`
		RoleARN *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		RoleRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"role_ref" json:"roleRef,omitempty"`
		Tags    *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		Version *string            `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EksServicesK8SAwsClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_eks_services_k8s_aws_cluster_v1alpha1_manifest"
}

func (r *EksServicesK8SAwsClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the Clusters API",
		MarkdownDescription: "Cluster is the Schema for the Clusters API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "ClusterSpec defines the desired state of Cluster.An object representing an Amazon EKS cluster.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.An object representing an Amazon EKS cluster.",
				Attributes: map[string]schema.Attribute{
					"access_config": schema.SingleNestedAttribute{
						Description:         "The access configuration for the cluster.",
						MarkdownDescription: "The access configuration for the cluster.",
						Attributes: map[string]schema.Attribute{
							"authentication_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bootstrap_cluster_creator_admin_permissions": schema.BoolAttribute{
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

					"client_request_token": schema.StringAttribute{
						Description:         "A unique, case-sensitive identifier that you provide to ensure the idempotencyof the request.",
						MarkdownDescription: "A unique, case-sensitive identifier that you provide to ensure the idempotencyof the request.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encryption_config": schema.ListNestedAttribute{
						Description:         "The encryption configuration for the cluster.",
						MarkdownDescription: "The encryption configuration for the cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"provider": schema.SingleNestedAttribute{
									Description:         "Identifies the Key Management Service (KMS) key used to encrypt the secrets.",
									MarkdownDescription: "Identifies the Key Management Service (KMS) key used to encrypt the secrets.",
									Attributes: map[string]schema.Attribute{
										"key_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_ref": schema.SingleNestedAttribute{
											Description:         "Reference field for KeyARN",
											MarkdownDescription: "Reference field for KeyARN",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"resources": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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

					"kubernetes_network_config": schema.SingleNestedAttribute{
						Description:         "The Kubernetes network configuration for the cluster.",
						MarkdownDescription: "The Kubernetes network configuration for the cluster.",
						Attributes: map[string]schema.Attribute{
							"ip_family": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_i_pv4_cidr": schema.StringAttribute{
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

					"logging": schema.SingleNestedAttribute{
						Description:         "Enable or disable exporting the Kubernetes control plane logs for your clusterto CloudWatch Logs. By default, cluster control plane logs aren't exportedto CloudWatch Logs. For more information, see Amazon EKS Cluster controlplane logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)in the Amazon EKS User Guide .CloudWatch Logs ingestion, archive storage, and data scanning rates applyto exported control plane logs. For more information, see CloudWatch Pricing(http://aws.amazon.com/cloudwatch/pricing/).",
						MarkdownDescription: "Enable or disable exporting the Kubernetes control plane logs for your clusterto CloudWatch Logs. By default, cluster control plane logs aren't exportedto CloudWatch Logs. For more information, see Amazon EKS Cluster controlplane logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)in the Amazon EKS User Guide .CloudWatch Logs ingestion, archive storage, and data scanning rates applyto exported control plane logs. For more information, see CloudWatch Pricing(http://aws.amazon.com/cloudwatch/pricing/).",
						Attributes: map[string]schema.Attribute{
							"cluster_logging": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"types": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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

					"name": schema.StringAttribute{
						Description:         "The unique name to give to your cluster.",
						MarkdownDescription: "The unique name to give to your cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"outpost_config": schema.SingleNestedAttribute{
						Description:         "An object representing the configuration of your local Amazon EKS clusteron an Amazon Web Services Outpost. Before creating a local cluster on anOutpost, review Local clusters for Amazon EKS on Amazon Web Services Outposts(https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-local-cluster-overview.html)in the Amazon EKS User Guide. This object isn't available for creating AmazonEKS clusters on the Amazon Web Services cloud.",
						MarkdownDescription: "An object representing the configuration of your local Amazon EKS clusteron an Amazon Web Services Outpost. Before creating a local cluster on anOutpost, review Local clusters for Amazon EKS on Amazon Web Services Outposts(https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-local-cluster-overview.html)in the Amazon EKS User Guide. This object isn't available for creating AmazonEKS clusters on the Amazon Web Services cloud.",
						Attributes: map[string]schema.Attribute{
							"control_plane_instance_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"control_plane_placement": schema.SingleNestedAttribute{
								Description:         "The placement configuration for all the control plane instances of your localAmazon EKS cluster on an Amazon Web Services Outpost. For more information,see Capacity considerations (https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html)in the Amazon EKS User Guide.",
								MarkdownDescription: "The placement configuration for all the control plane instances of your localAmazon EKS cluster on an Amazon Web Services Outpost. For more information,see Capacity considerations (https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-capacity-considerations.html)in the Amazon EKS User Guide.",
								Attributes: map[string]schema.Attribute{
									"group_name": schema.StringAttribute{
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

							"outpost_ar_ns": schema.ListAttribute{
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

					"resources_vpc_config": schema.SingleNestedAttribute{
						Description:         "The VPC configuration that's used by the cluster control plane. Amazon EKSVPC resources have specific requirements to work properly with Kubernetes.For more information, see Cluster VPC Considerations (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html)and Cluster Security Group Considerations (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html)in the Amazon EKS User Guide. You must specify at least two subnets. Youcan specify up to five security groups. However, we recommend that you usea dedicated security group for your cluster control plane.",
						MarkdownDescription: "The VPC configuration that's used by the cluster control plane. Amazon EKSVPC resources have specific requirements to work properly with Kubernetes.For more information, see Cluster VPC Considerations (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html)and Cluster Security Group Considerations (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html)in the Amazon EKS User Guide. You must specify at least two subnets. Youcan specify up to five security groups. However, we recommend that you usea dedicated security group for your cluster control plane.",
						Attributes: map[string]schema.Attribute{
							"endpoint_private_access": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint_public_access": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"public_access_cid_rs": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_group_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_group_refs": schema.ListNestedAttribute{
								Description:         "Reference field for SecurityGroupIDs",
								MarkdownDescription: "Reference field for SecurityGroupIDs",
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

							"subnet_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subnet_refs": schema.ListNestedAttribute{
								Description:         "Reference field for SubnetIDs",
								MarkdownDescription: "Reference field for SubnetIDs",
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the IAM role that provides permissionsfor the Kubernetes control plane to make calls to Amazon Web Services APIoperations on your behalf. For more information, see Amazon EKS Service IAMRole (https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html)in the Amazon EKS User Guide .",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM role that provides permissionsfor the Kubernetes control plane to make calls to Amazon Web Services APIoperations on your behalf. For more information, see Amazon EKS Service IAMRole (https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html)in the Amazon EKS User Guide .",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"role_ref": schema.SingleNestedAttribute{
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

					"tags": schema.MapAttribute{
						Description:         "Metadata that assists with categorization and organization. Each tag consistsof a key and an optional value. You define both. Tags don't propagate toany other cluster or Amazon Web Services resources.",
						MarkdownDescription: "Metadata that assists with categorization and organization. Each tag consistsof a key and an optional value. You define both. Tags don't propagate toany other cluster or Amazon Web Services resources.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "The desired Kubernetes version for your cluster. If you don't specify a valuehere, the default version available in Amazon EKS is used.The default version might not be the latest version available.",
						MarkdownDescription: "The desired Kubernetes version for your cluster. If you don't specify a valuehere, the default version available in Amazon EKS is used.The default version might not be the latest version available.",
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

func (r *EksServicesK8SAwsClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_eks_services_k8s_aws_cluster_v1alpha1_manifest")

	var model EksServicesK8SAwsClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("eks.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
