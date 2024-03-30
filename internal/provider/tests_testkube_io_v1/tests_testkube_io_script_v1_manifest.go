/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v1

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
	_ datasource.DataSource = &TestsTestkubeIoScriptV1Manifest{}
)

func NewTestsTestkubeIoScriptV1Manifest() datasource.DataSource {
	return &TestsTestkubeIoScriptV1Manifest{}
}

type TestsTestkubeIoScriptV1Manifest struct{}

type TestsTestkubeIoScriptV1ManifestData struct {
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
		Content    *string            `tfsdk:"content" json:"content,omitempty"`
		Input_type *string            `tfsdk:"input_type" json:"input-type,omitempty"`
		Name       *string            `tfsdk:"name" json:"name,omitempty"`
		Params     *map[string]string `tfsdk:"params" json:"params,omitempty"`
		Repository *struct {
			Branch   *string `tfsdk:"branch" json:"branch,omitempty"`
			Path     *string `tfsdk:"path" json:"path,omitempty"`
			Token    *string `tfsdk:"token" json:"token,omitempty"`
			Type     *string `tfsdk:"type" json:"type,omitempty"`
			Uri      *string `tfsdk:"uri" json:"uri,omitempty"`
			Username *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"repository" json:"repository,omitempty"`
		Tags *[]string `tfsdk:"tags" json:"tags,omitempty"`
		Type *string   `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoScriptV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_script_v1_manifest"
}

func (r *TestsTestkubeIoScriptV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Script is the Schema for the scripts API",
		MarkdownDescription: "Script is the Schema for the scripts API",
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
				Description:         "ScriptSpec defines the desired state of Script",
				MarkdownDescription: "ScriptSpec defines the desired state of Script",
				Attributes: map[string]schema.Attribute{
					"content": schema.StringAttribute{
						Description:         "script content as string (content depends from executor)",
						MarkdownDescription: "script content as string (content depends from executor)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"input_type": schema.StringAttribute{
						Description:         "script content type can be:  - direct content - created from file, - git repo directory checkout in case when test is some kind of project or have more than one file,",
						MarkdownDescription: "script content type can be:  - direct content - created from file, - git repo directory checkout in case when test is some kind of project or have more than one file,",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "script execution custom name",
						MarkdownDescription: "script execution custom name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"params": schema.MapAttribute{
						Description:         "execution params passed to executor",
						MarkdownDescription: "execution params passed to executor",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repository": schema.SingleNestedAttribute{
						Description:         "repository details if exists",
						MarkdownDescription: "repository details if exists",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "branch/tag name for checkout",
								MarkdownDescription: "branch/tag name for checkout",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
								MarkdownDescription: "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token": schema.StringAttribute{
								Description:         "git auth token for private repositories",
								MarkdownDescription: "git auth token for private repositories",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type_ repository type",
								MarkdownDescription: "Type_ repository type",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"uri": schema.StringAttribute{
								Description:         "Uri of content file or git directory",
								MarkdownDescription: "Uri of content file or git directory",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"username": schema.StringAttribute{
								Description:         "git auth username for private repositories",
								MarkdownDescription: "git auth username for private repositories",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "script type",
						MarkdownDescription: "script type",
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

func (r *TestsTestkubeIoScriptV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tests_testkube_io_script_v1_manifest")

	var model TestsTestkubeIoScriptV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tests.testkube.io/v1")
	model.Kind = pointer.String("Script")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
