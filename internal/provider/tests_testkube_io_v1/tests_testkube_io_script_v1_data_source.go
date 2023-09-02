/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &TestsTestkubeIoScriptV1DataSource{}
	_ datasource.DataSourceWithConfigure = &TestsTestkubeIoScriptV1DataSource{}
)

func NewTestsTestkubeIoScriptV1DataSource() datasource.DataSource {
	return &TestsTestkubeIoScriptV1DataSource{}
}

type TestsTestkubeIoScriptV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TestsTestkubeIoScriptV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *TestsTestkubeIoScriptV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_script_v1"
}

func (r *TestsTestkubeIoScriptV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Script is the Schema for the scripts API",
		MarkdownDescription: "Script is the Schema for the scripts API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
						Optional:            false,
						Computed:            true,
					},

					"input_type": schema.StringAttribute{
						Description:         "script content type can be:  - direct content - created from file, - git repo directory checkout in case when test is some kind of project or have more than one file,",
						MarkdownDescription: "script content type can be:  - direct content - created from file, - git repo directory checkout in case when test is some kind of project or have more than one file,",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "script execution custom name",
						MarkdownDescription: "script execution custom name",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"params": schema.MapAttribute{
						Description:         "execution params passed to executor",
						MarkdownDescription: "execution params passed to executor",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"repository": schema.SingleNestedAttribute{
						Description:         "repository details if exists",
						MarkdownDescription: "repository details if exists",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "branch/tag name for checkout",
								MarkdownDescription: "branch/tag name for checkout",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"path": schema.StringAttribute{
								Description:         "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
								MarkdownDescription: "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"token": schema.StringAttribute{
								Description:         "git auth token for private repositories",
								MarkdownDescription: "git auth token for private repositories",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Type_ repository type",
								MarkdownDescription: "Type_ repository type",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uri": schema.StringAttribute{
								Description:         "Uri of content file or git directory",
								MarkdownDescription: "Uri of content file or git directory",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"username": schema.StringAttribute{
								Description:         "git auth username for private repositories",
								MarkdownDescription: "git auth username for private repositories",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tags": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"type": schema.StringAttribute{
						Description:         "script type",
						MarkdownDescription: "script type",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *TestsTestkubeIoScriptV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *TestsTestkubeIoScriptV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_tests_testkube_io_script_v1")

	var data TestsTestkubeIoScriptV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v1", Resource: "Script"}).
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

	var readResponse TestsTestkubeIoScriptV1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("tests.testkube.io/v1")
	data.Kind = pointer.String("Script")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
