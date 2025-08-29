/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package v2_edp_epam_com_v1

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
	_ datasource.DataSource = &V2EdpEpamComCodebaseV1Manifest{}
)

func NewV2EdpEpamComCodebaseV1Manifest() datasource.DataSource {
	return &V2EdpEpamComCodebaseV1Manifest{}
}

type V2EdpEpamComCodebaseV1Manifest struct{}

type V2EdpEpamComCodebaseV1ManifestData struct {
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
		BranchToCopyInDefaultBranch *string `tfsdk:"branch_to_copy_in_default_branch" json:"branchToCopyInDefaultBranch,omitempty"`
		BuildTool                   *string `tfsdk:"build_tool" json:"buildTool,omitempty"`
		CiTool                      *string `tfsdk:"ci_tool" json:"ciTool,omitempty"`
		CommitMessagePattern        *string `tfsdk:"commit_message_pattern" json:"commitMessagePattern,omitempty"`
		DefaultBranch               *string `tfsdk:"default_branch" json:"defaultBranch,omitempty"`
		DeploymentScript            *string `tfsdk:"deployment_script" json:"deploymentScript,omitempty"`
		Description                 *string `tfsdk:"description" json:"description,omitempty"`
		DisablePutDeployTemplates   *bool   `tfsdk:"disable_put_deploy_templates" json:"disablePutDeployTemplates,omitempty"`
		EmptyProject                *bool   `tfsdk:"empty_project" json:"emptyProject,omitempty"`
		Framework                   *string `tfsdk:"framework" json:"framework,omitempty"`
		GitServer                   *string `tfsdk:"git_server" json:"gitServer,omitempty"`
		GitUrlPath                  *string `tfsdk:"git_url_path" json:"gitUrlPath,omitempty"`
		JiraIssueMetadataPayload    *string `tfsdk:"jira_issue_metadata_payload" json:"jiraIssueMetadataPayload,omitempty"`
		JiraServer                  *string `tfsdk:"jira_server" json:"jiraServer,omitempty"`
		Lang                        *string `tfsdk:"lang" json:"lang,omitempty"`
		Private                     *bool   `tfsdk:"private" json:"private,omitempty"`
		Repository                  *struct {
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"repository" json:"repository,omitempty"`
		Strategy            *string `tfsdk:"strategy" json:"strategy,omitempty"`
		TestReportFramework *string `tfsdk:"test_report_framework" json:"testReportFramework,omitempty"`
		TicketNamePattern   *string `tfsdk:"ticket_name_pattern" json:"ticketNamePattern,omitempty"`
		Type                *string `tfsdk:"type" json:"type,omitempty"`
		Versioning          *struct {
			StartFrom *string `tfsdk:"start_from" json:"startFrom,omitempty"`
			Type      *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"versioning" json:"versioning,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *V2EdpEpamComCodebaseV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_v2_edp_epam_com_codebase_v1_manifest"
}

func (r *V2EdpEpamComCodebaseV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Codebase is the Schema for the Codebases API.",
		MarkdownDescription: "Codebase is the Schema for the Codebases API.",
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
				Description:         "CodebaseSpec defines the desired state of Codebase.",
				MarkdownDescription: "CodebaseSpec defines the desired state of Codebase.",
				Attributes: map[string]schema.Attribute{
					"branch_to_copy_in_default_branch": schema.StringAttribute{
						Description:         "While we clone new codebase we can select specific branch to clone. Selected branch will become a default branch for a new codebase (e.g. master, main).",
						MarkdownDescription: "While we clone new codebase we can select specific branch to clone. Selected branch will become a default branch for a new codebase (e.g. master, main).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"build_tool": schema.StringAttribute{
						Description:         "A build tool which is used on codebase.",
						MarkdownDescription: "A build tool which is used on codebase.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ci_tool": schema.StringAttribute{
						Description:         "A name of tool which should be used as CI.",
						MarkdownDescription: "A name of tool which should be used as CI.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"commit_message_pattern": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_branch": schema.StringAttribute{
						Description:         "Name of default branch.",
						MarkdownDescription: "Name of default branch.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"deployment_script": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A short description of codebase.",
						MarkdownDescription: "A short description of codebase.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_put_deploy_templates": schema.BoolAttribute{
						Description:         "Controller must skip step 'put deploy templates' in action chain.",
						MarkdownDescription: "Controller must skip step 'put deploy templates' in action chain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"empty_project": schema.BoolAttribute{
						Description:         "A flag indicating how project should be provisioned. Default: false",
						MarkdownDescription: "A flag indicating how project should be provisioned. Default: false",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"framework": schema.StringAttribute{
						Description:         "A framework used in codebase.",
						MarkdownDescription: "A framework used in codebase.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"git_server": schema.StringAttribute{
						Description:         "A name of git server which will be used as VCS. Example: 'gerrit'.",
						MarkdownDescription: "A name of git server which will be used as VCS. Example: 'gerrit'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"git_url_path": schema.StringAttribute{
						Description:         "A relative path for git repository. Should start from /. Example: /company/api-app.",
						MarkdownDescription: "A relative path for git repository. Should start from /. Example: /company/api-app.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"jira_issue_metadata_payload": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"jira_server": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"lang": schema.StringAttribute{
						Description:         "Programming language used in codebase.",
						MarkdownDescription: "Programming language used in codebase.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"private": schema.BoolAttribute{
						Description:         "Private indicates if we need to create private repository.",
						MarkdownDescription: "Private indicates if we need to create private repository.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repository": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"strategy": schema.StringAttribute{
						Description:         "integration strategy for a codebase, e.g. clone, import, etc.",
						MarkdownDescription: "integration strategy for a codebase, e.g. clone, import, etc.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("create", "clone", "import"),
						},
					},

					"test_report_framework": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ticket_name_pattern": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type of codebase. E.g. application, autotest or library.",
						MarkdownDescription: "Type of codebase. E.g. application, autotest or library.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"versioning": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"start_from": schema.StringAttribute{
								Description:         "StartFrom is required when versioning type is not default.",
								MarkdownDescription: "StartFrom is required when versioning type is not default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
	}
}

func (r *V2EdpEpamComCodebaseV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_v2_edp_epam_com_codebase_v1_manifest")

	var model V2EdpEpamComCodebaseV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("v2.edp.epam.com/v1")
	model.Kind = pointer.String("Codebase")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
