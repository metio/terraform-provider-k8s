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
	_ datasource.DataSource = &V2EdpEpamComGerritMergeRequestV1Manifest{}
)

func NewV2EdpEpamComGerritMergeRequestV1Manifest() datasource.DataSource {
	return &V2EdpEpamComGerritMergeRequestV1Manifest{}
}

type V2EdpEpamComGerritMergeRequestV1Manifest struct{}

type V2EdpEpamComGerritMergeRequestV1ManifestData struct {
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
		AdditionalArguments *[]string `tfsdk:"additional_arguments" json:"additionalArguments,omitempty"`
		AuthorEmail         *string   `tfsdk:"author_email" json:"authorEmail,omitempty"`
		AuthorName          *string   `tfsdk:"author_name" json:"authorName,omitempty"`
		ChangesConfigMap    *string   `tfsdk:"changes_config_map" json:"changesConfigMap,omitempty"`
		CommitMessage       *string   `tfsdk:"commit_message" json:"commitMessage,omitempty"`
		OwnerName           *string   `tfsdk:"owner_name" json:"ownerName,omitempty"`
		ProjectName         *string   `tfsdk:"project_name" json:"projectName,omitempty"`
		SourceBranch        *string   `tfsdk:"source_branch" json:"sourceBranch,omitempty"`
		TargetBranch        *string   `tfsdk:"target_branch" json:"targetBranch,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *V2EdpEpamComGerritMergeRequestV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_v2_edp_epam_com_gerrit_merge_request_v1_manifest"
}

func (r *V2EdpEpamComGerritMergeRequestV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GerritMergeRequest is the Schema for the gerrit merge request API.",
		MarkdownDescription: "GerritMergeRequest is the Schema for the gerrit merge request API.",
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
				Description:         "GerritMergeRequestSpec defines the desired state of GerritMergeRequest.",
				MarkdownDescription: "GerritMergeRequestSpec defines the desired state of GerritMergeRequest.",
				Attributes: map[string]schema.Attribute{
					"additional_arguments": schema.ListAttribute{
						Description:         "AdditionalArguments contains merge command additional command line arguments.",
						MarkdownDescription: "AdditionalArguments contains merge command additional command line arguments.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"author_email": schema.StringAttribute{
						Description:         "AuthorEmail is the email of the user who creates the merge request.",
						MarkdownDescription: "AuthorEmail is the email of the user who creates the merge request.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"author_name": schema.StringAttribute{
						Description:         "AuthorName is the name of the user who creates the merge request.",
						MarkdownDescription: "AuthorName is the name of the user who creates the merge request.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"changes_config_map": schema.StringAttribute{
						Description:         "ChangesConfigMap is the name of the ConfigMap, which contains files contents that should be merged. ConfigMap should contain eny data keys with content in the json format: {'path': '/controllers/user.go', 'contents': 'some code here'} - to add file or format: {'path': '/controllers/user.go'} - to remove file. If files already exist in the project, they will be overwritten. If empty, sourceBranch should be set.",
						MarkdownDescription: "ChangesConfigMap is the name of the ConfigMap, which contains files contents that should be merged. ConfigMap should contain eny data keys with content in the json format: {'path': '/controllers/user.go', 'contents': 'some code here'} - to add file or format: {'path': '/controllers/user.go'} - to remove file. If files already exist in the project, they will be overwritten. If empty, sourceBranch should be set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"commit_message": schema.StringAttribute{
						Description:         "CommitMessage is the commit message for the merge request. If empty, the operator will generate the commit message.",
						MarkdownDescription: "CommitMessage is the commit message for the merge request. If empty, the operator will generate the commit message.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"owner_name": schema.StringAttribute{
						Description:         "OwnerName is the name of Gerrit CR, which should be used to initialize the client. If empty, the operator will get first Gerrit CR from the namespace.",
						MarkdownDescription: "OwnerName is the name of Gerrit CR, which should be used to initialize the client. If empty, the operator will get first Gerrit CR from the namespace.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"project_name": schema.StringAttribute{
						Description:         "ProjectName is gerrit project name.",
						MarkdownDescription: "ProjectName is gerrit project name.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source_branch": schema.StringAttribute{
						Description:         "SourceBranch is the name of the branch from which the changes should be merged. If empty, changesConfigMap should be set.",
						MarkdownDescription: "SourceBranch is the name of the branch from which the changes should be merged. If empty, changesConfigMap should be set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_branch": schema.StringAttribute{
						Description:         "TargetBranch is the name of the branch to which the changes should be merged. If changesConfigMap is set, the targetBranch can be only the origin HEAD branch.",
						MarkdownDescription: "TargetBranch is the name of the branch to which the changes should be merged. If changesConfigMap is set, the targetBranch can be only the origin HEAD branch.",
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

func (r *V2EdpEpamComGerritMergeRequestV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_v2_edp_epam_com_gerrit_merge_request_v1_manifest")

	var model V2EdpEpamComGerritMergeRequestV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("v2.edp.epam.com/v1")
	model.Kind = pointer.String("GerritMergeRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
