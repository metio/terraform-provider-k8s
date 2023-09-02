/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package executor_testkube_io_v1

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
	_ datasource.DataSource              = &ExecutorTestkubeIoExecutorV1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExecutorTestkubeIoExecutorV1DataSource{}
)

func NewExecutorTestkubeIoExecutorV1DataSource() datasource.DataSource {
	return &ExecutorTestkubeIoExecutorV1DataSource{}
}

type ExecutorTestkubeIoExecutorV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExecutorTestkubeIoExecutorV1DataSourceData struct {
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
		Args             *[]string `tfsdk:"args" json:"args,omitempty"`
		Command          *[]string `tfsdk:"command" json:"command,omitempty"`
		Content_types    *[]string `tfsdk:"content_types" json:"content_types,omitempty"`
		Executor_type    *string   `tfsdk:"executor_type" json:"executor_type,omitempty"`
		Features         *[]string `tfsdk:"features" json:"features,omitempty"`
		Image            *string   `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		Job_template *string `tfsdk:"job_template" json:"job_template,omitempty"`
		Meta         *struct {
			DocsURI  *string            `tfsdk:"docs_uri" json:"docsURI,omitempty"`
			IconURI  *string            `tfsdk:"icon_uri" json:"iconURI,omitempty"`
			Tooltips *map[string]string `tfsdk:"tooltips" json:"tooltips,omitempty"`
		} `tfsdk:"meta" json:"meta,omitempty"`
		Types *[]string `tfsdk:"types" json:"types,omitempty"`
		Uri   *string   `tfsdk:"uri" json:"uri,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExecutorTestkubeIoExecutorV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_executor_testkube_io_executor_v1"
}

func (r *ExecutorTestkubeIoExecutorV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Executor is the Schema for the executors API",
		MarkdownDescription: "Executor is the Schema for the executors API",
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
				Description:         "ExecutorSpec defines the desired state of Executor",
				MarkdownDescription: "ExecutorSpec defines the desired state of Executor",
				Attributes: map[string]schema.Attribute{
					"args": schema.ListAttribute{
						Description:         "executor binary arguments",
						MarkdownDescription: "executor binary arguments",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"command": schema.ListAttribute{
						Description:         "executor default binary command",
						MarkdownDescription: "executor default binary command",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"content_types": schema.ListAttribute{
						Description:         "ContentTypes list of handled content types",
						MarkdownDescription: "ContentTypes list of handled content types",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"executor_type": schema.StringAttribute{
						Description:         "ExecutorType one of 'rest' for rest openapi based executors or 'job' which will be default runners for testkube or 'container' for container executors",
						MarkdownDescription: "ExecutorType one of 'rest' for rest openapi based executors or 'job' which will be default runners for testkube or 'container' for container executors",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"features": schema.ListAttribute{
						Description:         "Features list of possible features which executor handles",
						MarkdownDescription: "Features list of possible features which executor handles",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image": schema.StringAttribute{
						Description:         "Image for kube-job",
						MarkdownDescription: "Image for kube-job",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "container executor default image pull secrets",
						MarkdownDescription: "container executor default image pull secrets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"job_template": schema.StringAttribute{
						Description:         "Job template to launch executor",
						MarkdownDescription: "Job template to launch executor",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"meta": schema.SingleNestedAttribute{
						Description:         "Meta data about executor",
						MarkdownDescription: "Meta data about executor",
						Attributes: map[string]schema.Attribute{
							"docs_uri": schema.StringAttribute{
								Description:         "URI for executor docs",
								MarkdownDescription: "URI for executor docs",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"icon_uri": schema.StringAttribute{
								Description:         "URI for executor icon",
								MarkdownDescription: "URI for executor icon",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tooltips": schema.MapAttribute{
								Description:         "executor tooltips",
								MarkdownDescription: "executor tooltips",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"types": schema.ListAttribute{
						Description:         "Types defines what types can be handled by executor e.g. 'postman/collection', ':curl/command' etc",
						MarkdownDescription: "Types defines what types can be handled by executor e.g. 'postman/collection', ':curl/command' etc",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"uri": schema.StringAttribute{
						Description:         "URI for rest based executors",
						MarkdownDescription: "URI for rest based executors",
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

func (r *ExecutorTestkubeIoExecutorV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ExecutorTestkubeIoExecutorV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_executor_testkube_io_executor_v1")

	var data ExecutorTestkubeIoExecutorV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "executor.testkube.io", Version: "v1", Resource: "Executor"}).
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

	var readResponse ExecutorTestkubeIoExecutorV1DataSourceData
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
	data.ApiVersion = pointer.String("executor.testkube.io/v1")
	data.Kind = pointer.String("Executor")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
