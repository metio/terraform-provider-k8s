/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v2

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
	_ datasource.DataSource = &ResourcesTeleportDevTeleportProvisionTokenV2Manifest{}
)

func NewResourcesTeleportDevTeleportProvisionTokenV2Manifest() datasource.DataSource {
	return &ResourcesTeleportDevTeleportProvisionTokenV2Manifest{}
}

type ResourcesTeleportDevTeleportProvisionTokenV2Manifest struct{}

type ResourcesTeleportDevTeleportProvisionTokenV2ManifestData struct {
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
		Allow *[]struct {
			Aws_account *string   `tfsdk:"aws_account" json:"aws_account,omitempty"`
			Aws_arn     *string   `tfsdk:"aws_arn" json:"aws_arn,omitempty"`
			Aws_regions *[]string `tfsdk:"aws_regions" json:"aws_regions,omitempty"`
			Aws_role    *string   `tfsdk:"aws_role" json:"aws_role,omitempty"`
		} `tfsdk:"allow" json:"allow,omitempty"`
		Aws_iid_ttl *string `tfsdk:"aws_iid_ttl" json:"aws_iid_ttl,omitempty"`
		Azure       *struct {
			Allow *[]struct {
				Resource_groups *[]string `tfsdk:"resource_groups" json:"resource_groups,omitempty"`
				Subscription    *string   `tfsdk:"subscription" json:"subscription,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
		} `tfsdk:"azure" json:"azure,omitempty"`
		Bot_name *string `tfsdk:"bot_name" json:"bot_name,omitempty"`
		Circleci *struct {
			Allow *[]struct {
				Context_id *string `tfsdk:"context_id" json:"context_id,omitempty"`
				Project_id *string `tfsdk:"project_id" json:"project_id,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
			Organization_id *string `tfsdk:"organization_id" json:"organization_id,omitempty"`
		} `tfsdk:"circleci" json:"circleci,omitempty"`
		Gcp *struct {
			Allow *[]struct {
				Locations        *[]string `tfsdk:"locations" json:"locations,omitempty"`
				Project_ids      *[]string `tfsdk:"project_ids" json:"project_ids,omitempty"`
				Service_accounts *[]string `tfsdk:"service_accounts" json:"service_accounts,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
		} `tfsdk:"gcp" json:"gcp,omitempty"`
		Github *struct {
			Allow *[]struct {
				Actor            *string `tfsdk:"actor" json:"actor,omitempty"`
				Environment      *string `tfsdk:"environment" json:"environment,omitempty"`
				Ref              *string `tfsdk:"ref" json:"ref,omitempty"`
				Ref_type         *string `tfsdk:"ref_type" json:"ref_type,omitempty"`
				Repository       *string `tfsdk:"repository" json:"repository,omitempty"`
				Repository_owner *string `tfsdk:"repository_owner" json:"repository_owner,omitempty"`
				Sub              *string `tfsdk:"sub" json:"sub,omitempty"`
				Workflow         *string `tfsdk:"workflow" json:"workflow,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
			Enterprise_server_host *string `tfsdk:"enterprise_server_host" json:"enterprise_server_host,omitempty"`
		} `tfsdk:"github" json:"github,omitempty"`
		Gitlab *struct {
			Allow *[]struct {
				Environment     *string `tfsdk:"environment" json:"environment,omitempty"`
				Namespace_path  *string `tfsdk:"namespace_path" json:"namespace_path,omitempty"`
				Pipeline_source *string `tfsdk:"pipeline_source" json:"pipeline_source,omitempty"`
				Project_path    *string `tfsdk:"project_path" json:"project_path,omitempty"`
				Ref             *string `tfsdk:"ref" json:"ref,omitempty"`
				Ref_type        *string `tfsdk:"ref_type" json:"ref_type,omitempty"`
				Sub             *string `tfsdk:"sub" json:"sub,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
			Domain *string `tfsdk:"domain" json:"domain,omitempty"`
		} `tfsdk:"gitlab" json:"gitlab,omitempty"`
		Join_method *string `tfsdk:"join_method" json:"join_method,omitempty"`
		Kubernetes  *struct {
			Allow *[]struct {
				Service_account *string `tfsdk:"service_account" json:"service_account,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
			Static_jwks *struct {
				Jwks *string `tfsdk:"jwks" json:"jwks,omitempty"`
			} `tfsdk:"static_jwks" json:"static_jwks,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Roles                          *[]string          `tfsdk:"roles" json:"roles,omitempty"`
		Suggested_agent_matcher_labels *map[string]string `tfsdk:"suggested_agent_matcher_labels" json:"suggested_agent_matcher_labels,omitempty"`
		Suggested_labels               *map[string]string `tfsdk:"suggested_labels" json:"suggested_labels,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_provision_token_v2_manifest"
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ProvisionToken is the Schema for the provisiontokens API",
		MarkdownDescription: "ProvisionToken is the Schema for the provisiontokens API",
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
				Description:         "ProvisionToken resource definition v2 from Teleport",
				MarkdownDescription: "ProvisionToken resource definition v2 from Teleport",
				Attributes: map[string]schema.Attribute{
					"allow": schema.ListNestedAttribute{
						Description:         "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
						MarkdownDescription: "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws_account": schema.StringAttribute{
									Description:         "AWSAccount is the AWS account ID.",
									MarkdownDescription: "AWSAccount is the AWS account ID.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"aws_arn": schema.StringAttribute{
									Description:         "AWSARN is used for the IAM join method, the AWS identity of joining nodes must match this ARN. Supports wildcards '*' and '?'.",
									MarkdownDescription: "AWSARN is used for the IAM join method, the AWS identity of joining nodes must match this ARN. Supports wildcards '*' and '?'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"aws_regions": schema.ListAttribute{
									Description:         "AWSRegions is used for the EC2 join method and is a list of AWS regions a node is allowed to join from.",
									MarkdownDescription: "AWSRegions is used for the EC2 join method and is a list of AWS regions a node is allowed to join from.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"aws_role": schema.StringAttribute{
									Description:         "AWSRole is used for the EC2 join method and is the the ARN of the AWS role that the auth server will assume in order to call the ec2 API.",
									MarkdownDescription: "AWSRole is used for the EC2 join method and is the the ARN of the AWS role that the auth server will assume in order to call the ec2 API.",
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

					"aws_iid_ttl": schema.StringAttribute{
						Description:         "AWSIIDTTL is the TTL to use for AWS EC2 Instance Identity Documents used to join the cluster with this token.",
						MarkdownDescription: "AWSIIDTTL is the TTL to use for AWS EC2 Instance Identity Documents used to join the cluster with this token.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"azure": schema.SingleNestedAttribute{
						Description:         "Azure allows the configuration of options specific to the 'azure' join method.",
						MarkdownDescription: "Azure allows the configuration of options specific to the 'azure' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"resource_groups": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"subscription": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"bot_name": schema.StringAttribute{
						Description:         "BotName is the name of the bot this token grants access to, if any",
						MarkdownDescription: "BotName is the name of the bot this token grants access to, if any",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"circleci": schema.SingleNestedAttribute{
						Description:         "CircleCI allows the configuration of options specific to the 'circleci' join method.",
						MarkdownDescription: "CircleCI allows the configuration of options specific to the 'circleci' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"context_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"project_id": schema.StringAttribute{
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

							"organization_id": schema.StringAttribute{
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

					"gcp": schema.SingleNestedAttribute{
						Description:         "GCP allows the configuration of options specific to the 'gcp' join method.",
						MarkdownDescription: "GCP allows the configuration of options specific to the 'gcp' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"locations": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"project_ids": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts": schema.ListAttribute{
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

					"github": schema.SingleNestedAttribute{
						Description:         "GitHub allows the configuration of options specific to the 'github' join method.",
						MarkdownDescription: "GitHub allows the configuration of options specific to the 'github' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"actor": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"environment": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ref": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ref_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"repository": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"repository_owner": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"workflow": schema.StringAttribute{
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

							"enterprise_server_host": schema.StringAttribute{
								Description:         "EnterpriseServerHost allows joining from runners associated with a GitHub Enterprise Server instance. When unconfigured, tokens will be validated against github.com, but when configured to the host of a GHES instance, then the tokens will be validated against host.  This value should be the hostname of the GHES instance, and should not include the scheme or a path. The instance must be accessible over HTTPS at this hostname and the certificate must be trusted by the Auth Server.",
								MarkdownDescription: "EnterpriseServerHost allows joining from runners associated with a GitHub Enterprise Server instance. When unconfigured, tokens will be validated against github.com, but when configured to the host of a GHES instance, then the tokens will be validated against host.  This value should be the hostname of the GHES instance, and should not include the scheme or a path. The instance must be accessible over HTTPS at this hostname and the certificate must be trusted by the Auth Server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gitlab": schema.SingleNestedAttribute{
						Description:         "GitLab allows the configuration of options specific to the 'gitlab' join method.",
						MarkdownDescription: "GitLab allows the configuration of options specific to the 'gitlab' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"environment": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pipeline_source": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"project_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ref": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ref_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub": schema.StringAttribute{
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

							"domain": schema.StringAttribute{
								Description:         "Domain is the domain of your GitLab instance. This will default to 'gitlab.com' - but can be set to the domain of your self-hosted GitLab e.g 'gitlab.example.com'.",
								MarkdownDescription: "Domain is the domain of your GitLab instance. This will default to 'gitlab.com' - but can be set to the domain of your self-hosted GitLab e.g 'gitlab.example.com'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"join_method": schema.StringAttribute{
						Description:         "JoinMethod is the joining method required in order to use this token. Supported joining methods include 'token', 'ec2', and 'iam'.",
						MarkdownDescription: "JoinMethod is the joining method required in order to use this token. Supported joining methods include 'token', 'ec2', and 'iam'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "Kubernetes allows the configuration of options specific to the 'kubernetes' join method.",
						MarkdownDescription: "Kubernetes allows the configuration of options specific to the 'kubernetes' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"service_account": schema.StringAttribute{
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

							"static_jwks": schema.SingleNestedAttribute{
								Description:         "StaticJWKS is the configuration specific to the 'static_jwks' type.",
								MarkdownDescription: "StaticJWKS is the configuration specific to the 'static_jwks' type.",
								Attributes: map[string]schema.Attribute{
									"jwks": schema.StringAttribute{
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

							"type": schema.StringAttribute{
								Description:         "Type controls which behavior should be used for validating the Kubernetes Service Account token. Support values: - 'in_cluster' - 'static_jwks' If unset, this defaults to 'in_cluster'.",
								MarkdownDescription: "Type controls which behavior should be used for validating the Kubernetes Service Account token. Support values: - 'in_cluster' - 'static_jwks' If unset, this defaults to 'in_cluster'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"roles": schema.ListAttribute{
						Description:         "Roles is a list of roles associated with the token, that will be converted to metadata in the SSH and X509 certificates issued to the user of the token",
						MarkdownDescription: "Roles is a list of roles associated with the token, that will be converted to metadata in the SSH and X509 certificates issued to the user of the token",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suggested_agent_matcher_labels": schema.MapAttribute{
						Description:         "SuggestedAgentMatcherLabels is a set of labels to be used by agents to match on resources. When an agent uses this token, the agent should monitor resources that match those labels. For databases, this means adding the labels to 'db_service.resources.labels'. Currently, only node-join scripts create a configuration according to the suggestion.",
						MarkdownDescription: "SuggestedAgentMatcherLabels is a set of labels to be used by agents to match on resources. When an agent uses this token, the agent should monitor resources that match those labels. For databases, this means adding the labels to 'db_service.resources.labels'. Currently, only node-join scripts create a configuration according to the suggestion.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suggested_labels": schema.MapAttribute{
						Description:         "SuggestedLabels is a set of labels that resources should set when using this token to enroll themselves in the cluster. Currently, only node-join scripts create a configuration according to the suggestion.",
						MarkdownDescription: "SuggestedLabels is a set of labels that resources should set when using this token to enroll themselves in the cluster. Currently, only node-join scripts create a configuration according to the suggestion.",
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
		},
	}
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_provision_token_v2_manifest")

	var model ResourcesTeleportDevTeleportProvisionTokenV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportProvisionToken")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
