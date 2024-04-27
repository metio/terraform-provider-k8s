/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v2

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
	_ datasource.DataSource = &ResourcesTeleportDevTeleportProvisionTokenV2Manifest{}
)

func NewResourcesTeleportDevTeleportProvisionTokenV2Manifest() datasource.DataSource {
	return &ResourcesTeleportDevTeleportProvisionTokenV2Manifest{}
}

type ResourcesTeleportDevTeleportProvisionTokenV2Manifest struct{}

type ResourcesTeleportDevTeleportProvisionTokenV2ManifestData struct {
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
			Enterprise_slug        *string `tfsdk:"enterprise_slug" json:"enterprise_slug,omitempty"`
		} `tfsdk:"github" json:"github,omitempty"`
		Gitlab *struct {
			Allow *[]struct {
				Ci_config_ref_uri     *string `tfsdk:"ci_config_ref_uri" json:"ci_config_ref_uri,omitempty"`
				Ci_config_sha         *string `tfsdk:"ci_config_sha" json:"ci_config_sha,omitempty"`
				Deployment_tier       *string `tfsdk:"deployment_tier" json:"deployment_tier,omitempty"`
				Environment           *string `tfsdk:"environment" json:"environment,omitempty"`
				Environment_protected *bool   `tfsdk:"environment_protected" json:"environment_protected,omitempty"`
				Namespace_path        *string `tfsdk:"namespace_path" json:"namespace_path,omitempty"`
				Pipeline_source       *string `tfsdk:"pipeline_source" json:"pipeline_source,omitempty"`
				Project_path          *string `tfsdk:"project_path" json:"project_path,omitempty"`
				Project_visibility    *string `tfsdk:"project_visibility" json:"project_visibility,omitempty"`
				Ref                   *string `tfsdk:"ref" json:"ref,omitempty"`
				Ref_protected         *bool   `tfsdk:"ref_protected" json:"ref_protected,omitempty"`
				Ref_type              *string `tfsdk:"ref_type" json:"ref_type,omitempty"`
				Sub                   *string `tfsdk:"sub" json:"sub,omitempty"`
				User_email            *string `tfsdk:"user_email" json:"user_email,omitempty"`
				User_id               *string `tfsdk:"user_id" json:"user_id,omitempty"`
				User_login            *string `tfsdk:"user_login" json:"user_login,omitempty"`
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
		Roles     *[]string `tfsdk:"roles" json:"roles,omitempty"`
		Spacelift *struct {
			Allow *[]struct {
				Caller_id   *string `tfsdk:"caller_id" json:"caller_id,omitempty"`
				Caller_type *string `tfsdk:"caller_type" json:"caller_type,omitempty"`
				Scope       *string `tfsdk:"scope" json:"scope,omitempty"`
				Space_id    *string `tfsdk:"space_id" json:"space_id,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
			Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
		} `tfsdk:"spacelift" json:"spacelift,omitempty"`
		Suggested_agent_matcher_labels *map[string]string `tfsdk:"suggested_agent_matcher_labels" json:"suggested_agent_matcher_labels,omitempty"`
		Suggested_labels               *map[string]string `tfsdk:"suggested_labels" json:"suggested_labels,omitempty"`
		Tpm                            *struct {
			Allow *[]struct {
				Description           *string `tfsdk:"description" json:"description,omitempty"`
				Ek_certificate_serial *string `tfsdk:"ek_certificate_serial" json:"ek_certificate_serial,omitempty"`
				Ek_public_hash        *string `tfsdk:"ek_public_hash" json:"ek_public_hash,omitempty"`
			} `tfsdk:"allow" json:"allow,omitempty"`
			Ekcert_allowed_cas *[]string `tfsdk:"ekcert_allowed_cas" json:"ekcert_allowed_cas,omitempty"`
		} `tfsdk:"tpm" json:"tpm,omitempty"`
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
									Description:         "AWSRole is used for the EC2 join method and is the ARN of the AWS role that the auth server will assume in order to call the ec2 API.",
									MarkdownDescription: "AWSRole is used for the EC2 join method and is the ARN of the AWS role that the auth server will assume in order to call the ec2 API.",
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

							"enterprise_slug": schema.StringAttribute{
								Description:         "EnterpriseSlug allows the slug of a GitHub Enterprise organisation to be included in the expected issuer of the OIDC tokens. This is for compatibility with the 'include_enterprise_slug' option in GHE.  This field should be set to the slug of your enterprise if this is enabled. If this is not enabled, then this field must be left empty. This field cannot be specified if 'enterprise_server_host' is specified.  See https://docs.github.com/en/enterprise-cloud@latest/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#customizing-the-issuer-value-for-an-enterprise for more information about customized issuer values.",
								MarkdownDescription: "EnterpriseSlug allows the slug of a GitHub Enterprise organisation to be included in the expected issuer of the OIDC tokens. This is for compatibility with the 'include_enterprise_slug' option in GHE.  This field should be set to the slug of your enterprise if this is enabled. If this is not enabled, then this field must be left empty. This field cannot be specified if 'enterprise_server_host' is specified.  See https://docs.github.com/en/enterprise-cloud@latest/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#customizing-the-issuer-value-for-an-enterprise for more information about customized issuer values.",
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
										"ci_config_ref_uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ci_config_sha": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"deployment_tier": schema.StringAttribute{
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

										"environment_protected": schema.BoolAttribute{
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

										"project_visibility": schema.StringAttribute{
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

										"ref_protected": schema.BoolAttribute{
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

										"user_email": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_login": schema.StringAttribute{
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

					"spacelift": schema.SingleNestedAttribute{
						Description:         "Spacelift allows the configuration of options specific to the 'spacelift' join method.",
						MarkdownDescription: "Spacelift allows the configuration of options specific to the 'spacelift' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								MarkdownDescription: "Allow is a list of Rules, nodes using this token must match one allow rule to use this token.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"caller_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"caller_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scope": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"space_id": schema.StringAttribute{
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

							"hostname": schema.StringAttribute{
								Description:         "Hostname is the hostname of the Spacelift tenant that tokens will originate from. E.g 'example.app.spacelift.io'",
								MarkdownDescription: "Hostname is the hostname of the Spacelift tenant that tokens will originate from. E.g 'example.app.spacelift.io'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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

					"tpm": schema.SingleNestedAttribute{
						Description:         "TPM allows the configuration of options specific to the 'tpm' join method.",
						MarkdownDescription: "TPM allows the configuration of options specific to the 'tpm' join method.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListNestedAttribute{
								Description:         "Allow is a list of Rules, the presented delegated identity must match one allow rule to permit joining.",
								MarkdownDescription: "Allow is a list of Rules, the presented delegated identity must match one allow rule to permit joining.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ek_certificate_serial": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ek_public_hash": schema.StringAttribute{
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

							"ekcert_allowed_cas": schema.ListAttribute{
								Description:         "EKCertAllowedCAs is a list of CA certificates that will be used to validate TPM EKCerts. When specified, joining TPMs must present an EKCert signed by one of the specified CAs. TPMs that do not present an EKCert will be not permitted to join. When unspecified, TPMs will be allowed to join with either an EKCert or an EKPubHash.",
								MarkdownDescription: "EKCertAllowedCAs is a list of CA certificates that will be used to validate TPM EKCerts. When specified, joining TPMs must present an EKCert signed by one of the specified CAs. TPMs that do not present an EKCert will be not permitted to join. When unspecified, TPMs will be allowed to join with either an EKCert or an EKPubHash.",
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
