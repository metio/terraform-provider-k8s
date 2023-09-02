/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
)

var (
	_ resource.Resource                = &ResourcesTeleportDevTeleportProvisionTokenV2Resource{}
	_ resource.ResourceWithConfigure   = &ResourcesTeleportDevTeleportProvisionTokenV2Resource{}
	_ resource.ResourceWithImportState = &ResourcesTeleportDevTeleportProvisionTokenV2Resource{}
)

func NewResourcesTeleportDevTeleportProvisionTokenV2Resource() resource.Resource {
	return &ResourcesTeleportDevTeleportProvisionTokenV2Resource{}
}

type ResourcesTeleportDevTeleportProvisionTokenV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ResourcesTeleportDevTeleportProvisionTokenV2ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Roles                          *[]string          `tfsdk:"roles" json:"roles,omitempty"`
		Suggested_agent_matcher_labels *map[string]string `tfsdk:"suggested_agent_matcher_labels" json:"suggested_agent_matcher_labels,omitempty"`
		Suggested_labels               *map[string]string `tfsdk:"suggested_labels" json:"suggested_labels,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_provision_token_v2"
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
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

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_resources_teleport_dev_teleport_provision_token_v2")

	var model ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportProvisionToken")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportProvisionToken"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_resources_teleport_dev_teleport_provision_token_v2")

	var data ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportProvisionToken"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_resources_teleport_dev_teleport_provision_token_v2")

	var model ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("resources.teleport.dev/v2")
	model.Kind = pointer.String("TeleportProvisionToken")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportProvisionToken"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_resources_teleport_dev_teleport_provision_token_v2")

	var data ResourcesTeleportDevTeleportProvisionTokenV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v2", Resource: "TeleportProvisionToken"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *ResourcesTeleportDevTeleportProvisionTokenV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
