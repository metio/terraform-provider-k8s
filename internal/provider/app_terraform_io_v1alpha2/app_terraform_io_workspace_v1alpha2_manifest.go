/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_terraform_io_v1alpha2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppTerraformIoWorkspaceV1Alpha2Manifest{}
)

func NewAppTerraformIoWorkspaceV1Alpha2Manifest() datasource.DataSource {
	return &AppTerraformIoWorkspaceV1Alpha2Manifest{}
}

type AppTerraformIoWorkspaceV1Alpha2Manifest struct{}

type AppTerraformIoWorkspaceV1Alpha2ManifestData struct {
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
		AgentPool *struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"agent_pool" json:"agentPool,omitempty"`
		AllowDestroyPlan     *bool   `tfsdk:"allow_destroy_plan" json:"allowDestroyPlan,omitempty"`
		ApplyMethod          *string `tfsdk:"apply_method" json:"applyMethod,omitempty"`
		Description          *string `tfsdk:"description" json:"description,omitempty"`
		EnvironmentVariables *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Hcl         *bool   `tfsdk:"hcl" json:"hcl,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Sensitive   *bool   `tfsdk:"sensitive" json:"sensitive,omitempty"`
			Value       *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom   *struct {
				ConfigMapKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				SecretKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"environment_variables" json:"environmentVariables,omitempty"`
		ExecutionMode *string `tfsdk:"execution_mode" json:"executionMode,omitempty"`
		Name          *string `tfsdk:"name" json:"name,omitempty"`
		Notifications *[]struct {
			EmailAddresses *[]string `tfsdk:"email_addresses" json:"emailAddresses,omitempty"`
			EmailUsers     *[]string `tfsdk:"email_users" json:"emailUsers,omitempty"`
			Enabled        *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
			Name           *string   `tfsdk:"name" json:"name,omitempty"`
			Token          *string   `tfsdk:"token" json:"token,omitempty"`
			Triggers       *[]string `tfsdk:"triggers" json:"triggers,omitempty"`
			Type           *string   `tfsdk:"type" json:"type,omitempty"`
			Url            *string   `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"notifications" json:"notifications,omitempty"`
		Organization *string `tfsdk:"organization" json:"organization,omitempty"`
		Project      *struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"project" json:"project,omitempty"`
		RemoteStateSharing *struct {
			AllWorkspaces *bool `tfsdk:"all_workspaces" json:"allWorkspaces,omitempty"`
			Workspaces    *[]struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"workspaces" json:"workspaces,omitempty"`
		} `tfsdk:"remote_state_sharing" json:"remoteStateSharing,omitempty"`
		RunTasks *[]struct {
			EnforcementLevel *string `tfsdk:"enforcement_level" json:"enforcementLevel,omitempty"`
			Id               *string `tfsdk:"id" json:"id,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			Stage            *string `tfsdk:"stage" json:"stage,omitempty"`
		} `tfsdk:"run_tasks" json:"runTasks,omitempty"`
		RunTriggers *[]struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"run_triggers" json:"runTriggers,omitempty"`
		SshKey *struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"ssh_key" json:"sshKey,omitempty"`
		Tags       *[]string `tfsdk:"tags" json:"tags,omitempty"`
		TeamAccess *[]struct {
			Access *string `tfsdk:"access" json:"access,omitempty"`
			Custom *struct {
				RunTasks         *bool   `tfsdk:"run_tasks" json:"runTasks,omitempty"`
				Runs             *string `tfsdk:"runs" json:"runs,omitempty"`
				Sentinel         *string `tfsdk:"sentinel" json:"sentinel,omitempty"`
				StateVersions    *string `tfsdk:"state_versions" json:"stateVersions,omitempty"`
				Variables        *string `tfsdk:"variables" json:"variables,omitempty"`
				WorkspaceLocking *bool   `tfsdk:"workspace_locking" json:"workspaceLocking,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Team *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"team" json:"team,omitempty"`
		} `tfsdk:"team_access" json:"teamAccess,omitempty"`
		TerraformVariables *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Hcl         *bool   `tfsdk:"hcl" json:"hcl,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Sensitive   *bool   `tfsdk:"sensitive" json:"sensitive,omitempty"`
			Value       *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom   *struct {
				ConfigMapKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				SecretKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"terraform_variables" json:"terraformVariables,omitempty"`
		TerraformVersion *string `tfsdk:"terraform_version" json:"terraformVersion,omitempty"`
		Token            *struct {
			SecretKeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
		} `tfsdk:"token" json:"token,omitempty"`
		VersionControl *struct {
			Branch       *string `tfsdk:"branch" json:"branch,omitempty"`
			OAuthTokenID *string `tfsdk:"o_auth_token_id" json:"oAuthTokenID,omitempty"`
			Repository   *string `tfsdk:"repository" json:"repository,omitempty"`
		} `tfsdk:"version_control" json:"versionControl,omitempty"`
		WorkingDirectory *string `tfsdk:"working_directory" json:"workingDirectory,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppTerraformIoWorkspaceV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_terraform_io_workspace_v1alpha2_manifest"
}

func (r *AppTerraformIoWorkspaceV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Workspace is the Schema for the workspaces API",
		MarkdownDescription: "Workspace is the Schema for the workspaces API",
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
				Description:         "WorkspaceSpec defines the desired state of Workspace.",
				MarkdownDescription: "WorkspaceSpec defines the desired state of Workspace.",
				Attributes: map[string]schema.Attribute{
					"agent_pool": schema.SingleNestedAttribute{
						Description:         "Terraform Cloud Agents allow Terraform Cloud to communicate with isolated, private, or on-premises infrastructure. More information: - https://developer.hashicorp.com/terraform/cloud-docs/agents",
						MarkdownDescription: "Terraform Cloud Agents allow Terraform Cloud to communicate with isolated, private, or on-premises infrastructure. More information: - https://developer.hashicorp.com/terraform/cloud-docs/agents",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "Agent Pool ID. Must match pattern: '^apool-[a-zA-Z0-9]+$'",
								MarkdownDescription: "Agent Pool ID. Must match pattern: '^apool-[a-zA-Z0-9]+$'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^apool-[a-zA-Z0-9]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Agent Pool name.",
								MarkdownDescription: "Agent Pool name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"allow_destroy_plan": schema.BoolAttribute{
						Description:         "Allows a destroy plan to be created and applied. Default: 'true'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#destruction-and-deletion",
						MarkdownDescription: "Allows a destroy plan to be created and applied. Default: 'true'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#destruction-and-deletion",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"apply_method": schema.StringAttribute{
						Description:         "Define either change will be applied automatically(auto) or require an operator to confirm(manual). Must be one of the following values: 'auto', 'manual'. Default: 'manual'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#auto-apply-and-manual-apply",
						MarkdownDescription: "Define either change will be applied automatically(auto) or require an operator to confirm(manual). Must be one of the following values: 'auto', 'manual'. Default: 'manual'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#auto-apply-and-manual-apply",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(auto|manual)$`), ""),
						},
					},

					"description": schema.StringAttribute{
						Description:         "Workspace description.",
						MarkdownDescription: "Workspace description.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"environment_variables": schema.ListNestedAttribute{
						Description:         "Terraform Environment variables for all plans and applies in this workspace. Variables defined within a workspace always overwrite variables from variable sets that have the same type and the same key. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables#environment-variables",
						MarkdownDescription: "Terraform Environment variables for all plans and applies in this workspace. Variables defined within a workspace always overwrite variables from variable sets that have the same type and the same key. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables#environment-variables",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description of the variable.",
									MarkdownDescription: "Description of the variable.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"hcl": schema.BoolAttribute{
									Description:         "Parse this field as HashiCorp Configuration Language (HCL). This allows you to interpolate values at runtime. Default: 'false'.",
									MarkdownDescription: "Parse this field as HashiCorp Configuration Language (HCL). This allows you to interpolate values at runtime. Default: 'false'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the variable.",
									MarkdownDescription: "Name of the variable.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"sensitive": schema.BoolAttribute{
									Description:         "Sensitive variables are never shown in the UI or API. They may appear in Terraform logs if your configuration is designed to output them. Default: 'false'.",
									MarkdownDescription: "Sensitive variables are never shown in the UI or API. They may appear in Terraform logs if your configuration is designed to output them. Default: 'false'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value of the variable.",
									MarkdownDescription: "Value of the variable.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "Source for the variable's value. Cannot be used if value is not empty.",
									MarkdownDescription: "Source for the variable's value. Cannot be used if value is not empty.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a ConfigMap.",
											MarkdownDescription: "Selects a key of a ConfigMap.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap or its key must be defined",
													MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a Secret.",
											MarkdownDescription: "Selects a key of a Secret.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"execution_mode": schema.StringAttribute{
						Description:         "Define where the Terraform code will be executed. Must be one of the following values: 'agent', 'local', 'remote'. Default: 'remote'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#execution-mode",
						MarkdownDescription: "Define where the Terraform code will be executed. Must be one of the following values: 'agent', 'local', 'remote'. Default: 'remote'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#execution-mode",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(agent|local|remote)$`), ""),
						},
					},

					"name": schema.StringAttribute{
						Description:         "Workspace name.",
						MarkdownDescription: "Workspace name.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"notifications": schema.ListNestedAttribute{
						Description:         "Notifications allow you to send messages to other applications based on run and workspace events. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/notifications",
						MarkdownDescription: "Notifications allow you to send messages to other applications based on run and workspace events. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/notifications",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"email_addresses": schema.ListAttribute{
									Description:         "The list of email addresses that will receive notification emails. It is only available for Terraform Enterprise users. It is not available in Terraform Cloud.",
									MarkdownDescription: "The list of email addresses that will receive notification emails. It is only available for Terraform Enterprise users. It is not available in Terraform Cloud.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"email_users": schema.ListAttribute{
									Description:         "The list of users belonging to the organization that will receive notification emails.",
									MarkdownDescription: "The list of users belonging to the organization that will receive notification emails.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enabled": schema.BoolAttribute{
									Description:         "Whether the notification configuration should be enabled or not. Default: 'true'.",
									MarkdownDescription: "Whether the notification configuration should be enabled or not. Default: 'true'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Notification name.",
									MarkdownDescription: "Notification name.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"token": schema.StringAttribute{
									Description:         "The token of the notification.",
									MarkdownDescription: "The token of the notification.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"triggers": schema.ListAttribute{
									Description:         "The list of run events that will trigger notifications. Trigger represents the different TFC notifications that can be sent as a run's progress transitions between different states. There are two categories of triggers: - Health Events: 'assessment:check_failure', 'assessment:drifted', 'assessment:failed'. - Run Events: 'run:applying', 'run:completed', 'run:created', 'run:errored', 'run:needs_attention', 'run:planning'.",
									MarkdownDescription: "The list of run events that will trigger notifications. Trigger represents the different TFC notifications that can be sent as a run's progress transitions between different states. There are two categories of triggers: - Health Events: 'assessment:check_failure', 'assessment:drifted', 'assessment:failed'. - Run Events: 'run:applying', 'run:completed', 'run:created', 'run:errored', 'run:needs_attention', 'run:planning'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "The type of the notification. Must be one of the following values: 'email', 'generic', 'microsoft-teams', 'slack'.",
									MarkdownDescription: "The type of the notification. Must be one of the following values: 'email', 'generic', 'microsoft-teams', 'slack'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("email", "generic", "microsoft-teams", "slack"),
									},
								},

								"url": schema.StringAttribute{
									Description:         "The URL of the notification. Must match pattern: '^https?://.*'",
									MarkdownDescription: "The URL of the notification. Must match pattern: '^https?://.*'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^https?://.*`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"organization": schema.StringAttribute{
						Description:         "Organization name where the Workspace will be created. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/organizations",
						MarkdownDescription: "Organization name where the Workspace will be created. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/organizations",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"project": schema.SingleNestedAttribute{
						Description:         "Projects let you organize your workspaces into groups. Default: default organization project. More information: - https://developer.hashicorp.com/terraform/tutorials/cloud/projects",
						MarkdownDescription: "Projects let you organize your workspaces into groups. Default: default organization project. More information: - https://developer.hashicorp.com/terraform/tutorials/cloud/projects",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "Project ID. Must match pattern: '^prj-[a-zA-Z0-9]+$'",
								MarkdownDescription: "Project ID. Must match pattern: '^prj-[a-zA-Z0-9]+$'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^prj-[a-zA-Z0-9]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Project name.",
								MarkdownDescription: "Project name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_state_sharing": schema.SingleNestedAttribute{
						Description:         "Remote state access between workspaces. By default, new workspaces in Terraform Cloud do not allow other workspaces to access their state. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/state#accessing-state-from-other-workspaces",
						MarkdownDescription: "Remote state access between workspaces. By default, new workspaces in Terraform Cloud do not allow other workspaces to access their state. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/state#accessing-state-from-other-workspaces",
						Attributes: map[string]schema.Attribute{
							"all_workspaces": schema.BoolAttribute{
								Description:         "Allow access to the state for all workspaces within the same organization. Default: 'false'.",
								MarkdownDescription: "Allow access to the state for all workspaces within the same organization. Default: 'false'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workspaces": schema.ListNestedAttribute{
								Description:         "Allow access to the state for specific workspaces within the same organization.",
								MarkdownDescription: "Allow access to the state for specific workspaces within the same organization.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "Consumer Workspace ID. Must match pattern: '^ws-[a-zA-Z0-9]+$'",
											MarkdownDescription: "Consumer Workspace ID. Must match pattern: '^ws-[a-zA-Z0-9]+$'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^ws-[a-zA-Z0-9]+$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Consumer Workspace name.",
											MarkdownDescription: "Consumer Workspace name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
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

					"run_tasks": schema.ListNestedAttribute{
						Description:         "Run tasks allow Terraform Cloud to interact with external systems at specific points in the Terraform Cloud run lifecycle. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/run-tasks",
						MarkdownDescription: "Run tasks allow Terraform Cloud to interact with external systems at specific points in the Terraform Cloud run lifecycle. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/run-tasks",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"enforcement_level": schema.StringAttribute{
									Description:         "Run Task Enforcement Level. Can be one of 'advisory' or 'mandatory'. Default: 'advisory'. Must be one of the following values: 'advisory', 'mandatory' Default: 'advisory'.",
									MarkdownDescription: "Run Task Enforcement Level. Can be one of 'advisory' or 'mandatory'. Default: 'advisory'. Must be one of the following values: 'advisory', 'mandatory' Default: 'advisory'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(advisory|mandatory)$`), ""),
									},
								},

								"id": schema.StringAttribute{
									Description:         "Run Task ID. Must match pattern: '^task-[a-zA-Z0-9]+$'",
									MarkdownDescription: "Run Task ID. Must match pattern: '^task-[a-zA-Z0-9]+$'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^task-[a-zA-Z0-9]+$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Run Task Name.",
									MarkdownDescription: "Run Task Name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"stage": schema.StringAttribute{
									Description:         "Run Task Stage. Must be one of the following values: 'pre_apply', 'pre_plan', 'post_plan'. Default: 'post_plan'.",
									MarkdownDescription: "Run Task Stage. Must be one of the following values: 'pre_apply', 'pre_plan', 'post_plan'. Default: 'post_plan'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(pre_apply|pre_plan|post_plan)$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"run_triggers": schema.ListNestedAttribute{
						Description:         "Run triggers allow you to connect this workspace to one or more source workspaces. These connections allow runs to queue automatically in this workspace on successful apply of runs in any of the source workspaces. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/run-triggers",
						MarkdownDescription: "Run triggers allow you to connect this workspace to one or more source workspaces. These connections allow runs to queue automatically in this workspace on successful apply of runs in any of the source workspaces. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/run-triggers",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "Source Workspace ID. Must match pattern: '^ws-[a-zA-Z0-9]+$'",
									MarkdownDescription: "Source Workspace ID. Must match pattern: '^ws-[a-zA-Z0-9]+$'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^ws-[a-zA-Z0-9]+$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Source Workspace Name.",
									MarkdownDescription: "Source Workspace Name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ssh_key": schema.SingleNestedAttribute{
						Description:         "SSH key used to clone Terraform modules. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/ssh-keys",
						MarkdownDescription: "SSH key used to clone Terraform modules. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/ssh-keys",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "SSH key ID. Must match pattern: '^sshkey-[a-zA-Z0-9]+$'",
								MarkdownDescription: "SSH key ID. Must match pattern: '^sshkey-[a-zA-Z0-9]+$'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^sshkey-[a-zA-Z0-9]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "SSH key name.",
								MarkdownDescription: "SSH key name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.ListAttribute{
						Description:         "Workspace tags are used to help identify and group together workspaces. Tags must be one or more characters; can include letters, numbers, colons, hyphens, and underscores; and must begin and end with a letter or number.",
						MarkdownDescription: "Workspace tags are used to help identify and group together workspaces. Tags must be one or more characters; can include letters, numbers, colons, hyphens, and underscores; and must begin and end with a letter or number.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"team_access": schema.ListNestedAttribute{
						Description:         "Terraform Cloud workspaces can only be accessed by users with the correct permissions. You can manage permissions for a workspace on a per-team basis. When a workspace is created, only the owners team and teams with the 'manage workspaces' permission can access it, with full admin permissions. These teams' access can't be removed from a workspace. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/access",
						MarkdownDescription: "Terraform Cloud workspaces can only be accessed by users with the correct permissions. You can manage permissions for a workspace on a per-team basis. When a workspace is created, only the owners team and teams with the 'manage workspaces' permission can access it, with full admin permissions. These teams' access can't be removed from a workspace. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings/access",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access": schema.StringAttribute{
									Description:         "There are two ways to choose which permissions a given team has on a workspace: fixed permission sets, and custom permissions. Must be one of the following values: 'admin', 'custom', 'plan', 'read', 'write'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/permissions#workspace-permissions",
									MarkdownDescription: "There are two ways to choose which permissions a given team has on a workspace: fixed permission sets, and custom permissions. Must be one of the following values: 'admin', 'custom', 'plan', 'read', 'write'. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/permissions#workspace-permissions",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(admin|custom|plan|read|write)$`), ""),
									},
								},

								"custom": schema.SingleNestedAttribute{
									Description:         "Custom permissions let you assign specific, finer-grained permissions to a team than the broader fixed permission sets provide. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/permissions#custom-workspace-permissions",
									MarkdownDescription: "Custom permissions let you assign specific, finer-grained permissions to a team than the broader fixed permission sets provide. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/permissions#custom-workspace-permissions",
									Attributes: map[string]schema.Attribute{
										"run_tasks": schema.BoolAttribute{
											Description:         "Manage Workspace Run Tasks. Default: 'false'.",
											MarkdownDescription: "Manage Workspace Run Tasks. Default: 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"runs": schema.StringAttribute{
											Description:         "Run access. Must be one of the following values: 'apply', 'plan', 'read'. Default: 'read'.",
											MarkdownDescription: "Run access. Must be one of the following values: 'apply', 'plan', 'read'. Default: 'read'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(apply|plan|read)$`), ""),
											},
										},

										"sentinel": schema.StringAttribute{
											Description:         "Download Sentinel mocks. Must be one of the following values: 'none', 'read'. Default: 'none'.",
											MarkdownDescription: "Download Sentinel mocks. Must be one of the following values: 'none', 'read'. Default: 'none'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(none|read)$`), ""),
											},
										},

										"state_versions": schema.StringAttribute{
											Description:         "State access. Must be one of the following values: 'none', 'read', 'read-outputs', 'write'. Default: 'none'.",
											MarkdownDescription: "State access. Must be one of the following values: 'none', 'read', 'read-outputs', 'write'. Default: 'none'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(none|read|read-outputs|write)$`), ""),
											},
										},

										"variables": schema.StringAttribute{
											Description:         "Variable access. Must be one of the following values: 'none', 'read', 'write'. Default: 'none'.",
											MarkdownDescription: "Variable access. Must be one of the following values: 'none', 'read', 'write'. Default: 'none'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(none|read|write)$`), ""),
											},
										},

										"workspace_locking": schema.BoolAttribute{
											Description:         "Lock/unlock workspace. Default: 'false'.",
											MarkdownDescription: "Lock/unlock workspace. Default: 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"team": schema.SingleNestedAttribute{
									Description:         "Team to grant access. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/teams",
									MarkdownDescription: "Team to grant access. More information: - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/teams",
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "Team ID. Must match pattern: '^team-[a-zA-Z0-9]+$'",
											MarkdownDescription: "Team ID. Must match pattern: '^team-[a-zA-Z0-9]+$'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^team-[a-zA-Z0-9]+$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Team name.",
											MarkdownDescription: "Team name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"terraform_variables": schema.ListNestedAttribute{
						Description:         "Terraform variables for all plans and applies in this workspace. Variables defined within a workspace always overwrite variables from variable sets that have the same type and the same key. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables#terraform-variables",
						MarkdownDescription: "Terraform variables for all plans and applies in this workspace. Variables defined within a workspace always overwrite variables from variable sets that have the same type and the same key. More information: - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables - https://developer.hashicorp.com/terraform/cloud-docs/workspaces/variables#terraform-variables",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description of the variable.",
									MarkdownDescription: "Description of the variable.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"hcl": schema.BoolAttribute{
									Description:         "Parse this field as HashiCorp Configuration Language (HCL). This allows you to interpolate values at runtime. Default: 'false'.",
									MarkdownDescription: "Parse this field as HashiCorp Configuration Language (HCL). This allows you to interpolate values at runtime. Default: 'false'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the variable.",
									MarkdownDescription: "Name of the variable.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"sensitive": schema.BoolAttribute{
									Description:         "Sensitive variables are never shown in the UI or API. They may appear in Terraform logs if your configuration is designed to output them. Default: 'false'.",
									MarkdownDescription: "Sensitive variables are never shown in the UI or API. They may appear in Terraform logs if your configuration is designed to output them. Default: 'false'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value of the variable.",
									MarkdownDescription: "Value of the variable.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "Source for the variable's value. Cannot be used if value is not empty.",
									MarkdownDescription: "Source for the variable's value. Cannot be used if value is not empty.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a ConfigMap.",
											MarkdownDescription: "Selects a key of a ConfigMap.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap or its key must be defined",
													MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a Secret.",
											MarkdownDescription: "Selects a key of a Secret.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"terraform_version": schema.StringAttribute{
						Description:         "The version of Terraform to use for this workspace. If not specified, the latest available version will be used. Must match pattern: '^d{1}.d{1,2}.d{1,2}$' More information: - https://www.terraform.io/cloud-docs/workspaces/settings#terraform-version",
						MarkdownDescription: "The version of Terraform to use for this workspace. If not specified, the latest available version will be used. Must match pattern: '^d{1}.d{1,2}.d{1,2}$' More information: - https://www.terraform.io/cloud-docs/workspaces/settings#terraform-version",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1}\.\d{1,2}\.\d{1,2}$`), ""),
						},
					},

					"token": schema.SingleNestedAttribute{
						Description:         "API Token to be used for API calls.",
						MarkdownDescription: "API Token to be used for API calls.",
						Attributes: map[string]schema.Attribute{
							"secret_key_ref": schema.SingleNestedAttribute{
								Description:         "Selects a key of a secret in the workspace's namespace",
								MarkdownDescription: "Selects a key of a secret in the workspace's namespace",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"version_control": schema.SingleNestedAttribute{
						Description:         "Settings for the workspace's VCS repository, enabling the UI/VCS-driven run workflow. Omit this argument to utilize the CLI-driven and API-driven workflows, where runs are not driven by webhooks on your VCS provider. More information: - https://www.terraform.io/cloud-docs/run/ui - https://www.terraform.io/cloud-docs/vcs",
						MarkdownDescription: "Settings for the workspace's VCS repository, enabling the UI/VCS-driven run workflow. Omit this argument to utilize the CLI-driven and API-driven workflows, where runs are not driven by webhooks on your VCS provider. More information: - https://www.terraform.io/cloud-docs/run/ui - https://www.terraform.io/cloud-docs/vcs",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "The repository branch that Run will execute from. This defaults to the repository's default branch (e.g. main).",
								MarkdownDescription: "The repository branch that Run will execute from. This defaults to the repository's default branch (e.g. main).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"o_auth_token_id": schema.StringAttribute{
								Description:         "The VCS Connection (OAuth Connection + Token) to use. Must match pattern: '^ot-[a-zA-Z0-9]+$'",
								MarkdownDescription: "The VCS Connection (OAuth Connection + Token) to use. Must match pattern: '^ot-[a-zA-Z0-9]+$'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^ot-[a-zA-Z0-9]+$`), ""),
								},
							},

							"repository": schema.StringAttribute{
								Description:         "A reference to your VCS repository in the format '<organization>/<repository>' where '<organization>' and '<repository>' refer to the organization and repository in your VCS provider.",
								MarkdownDescription: "A reference to your VCS repository in the format '<organization>/<repository>' where '<organization>' and '<repository>' refer to the organization and repository in your VCS provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"working_directory": schema.StringAttribute{
						Description:         "The directory where Terraform will execute, specified as a relative path from the root of the configuration directory. More information: - https://www.terraform.io/cloud-docs/workspaces/settings#terraform-working-directory",
						MarkdownDescription: "The directory where Terraform will execute, specified as a relative path from the root of the configuration directory. More information: - https://www.terraform.io/cloud-docs/workspaces/settings#terraform-working-directory",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *AppTerraformIoWorkspaceV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_terraform_io_workspace_v1alpha2_manifest")

	var model AppTerraformIoWorkspaceV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("app.terraform.io/v1alpha2")
	model.Kind = pointer.String("Workspace")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
