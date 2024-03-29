/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package executor_testkube_io_v1

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
	_ datasource.DataSource = &ExecutorTestkubeIoExecutorV1Manifest{}
)

func NewExecutorTestkubeIoExecutorV1Manifest() datasource.DataSource {
	return &ExecutorTestkubeIoExecutorV1Manifest{}
}

type ExecutorTestkubeIoExecutorV1Manifest struct{}

type ExecutorTestkubeIoExecutorV1ManifestData struct {
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
		Args             *[]string `tfsdk:"args" json:"args,omitempty"`
		Command          *[]string `tfsdk:"command" json:"command,omitempty"`
		Content_types    *[]string `tfsdk:"content_types" json:"content_types,omitempty"`
		Executor_type    *string   `tfsdk:"executor_type" json:"executor_type,omitempty"`
		Features         *[]string `tfsdk:"features" json:"features,omitempty"`
		Image            *string   `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		JobTemplateReference *string `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
		Job_template         *string `tfsdk:"job_template" json:"job_template,omitempty"`
		Meta                 *struct {
			DocsURI  *string            `tfsdk:"docs_uri" json:"docsURI,omitempty"`
			IconURI  *string            `tfsdk:"icon_uri" json:"iconURI,omitempty"`
			Tooltips *map[string]string `tfsdk:"tooltips" json:"tooltips,omitempty"`
		} `tfsdk:"meta" json:"meta,omitempty"`
		Slaves *struct {
			Image *string `tfsdk:"image" json:"image,omitempty"`
		} `tfsdk:"slaves" json:"slaves,omitempty"`
		Types                  *[]string `tfsdk:"types" json:"types,omitempty"`
		Uri                    *string   `tfsdk:"uri" json:"uri,omitempty"`
		UseDataDirAsWorkingDir *bool     `tfsdk:"use_data_dir_as_working_dir" json:"useDataDirAsWorkingDir,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExecutorTestkubeIoExecutorV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_executor_testkube_io_executor_v1_manifest"
}

func (r *ExecutorTestkubeIoExecutorV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "ExecutorSpec defines the desired state of Executor",
				MarkdownDescription: "ExecutorSpec defines the desired state of Executor",
				Attributes: map[string]schema.Attribute{
					"args": schema.ListAttribute{
						Description:         "executor binary arguments",
						MarkdownDescription: "executor binary arguments",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"command": schema.ListAttribute{
						Description:         "executor default binary command",
						MarkdownDescription: "executor default binary command",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"content_types": schema.ListAttribute{
						Description:         "ContentTypes list of handled content types",
						MarkdownDescription: "ContentTypes list of handled content types",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"executor_type": schema.StringAttribute{
						Description:         "ExecutorType one of 'rest' for rest openapi based executors or 'job' which will be default runners for testkube or 'container' for container executors",
						MarkdownDescription: "ExecutorType one of 'rest' for rest openapi based executors or 'job' which will be default runners for testkube or 'container' for container executors",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("job", "container"),
						},
					},

					"features": schema.ListAttribute{
						Description:         "Features list of possible features which executor handles",
						MarkdownDescription: "Features list of possible features which executor handles",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Image for kube-job",
						MarkdownDescription: "Image for kube-job",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"job_template_reference": schema.StringAttribute{
						Description:         "name of the template resource",
						MarkdownDescription: "name of the template resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"job_template": schema.StringAttribute{
						Description:         "Job template to launch executor",
						MarkdownDescription: "Job template to launch executor",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"meta": schema.SingleNestedAttribute{
						Description:         "Meta data about executor",
						MarkdownDescription: "Meta data about executor",
						Attributes: map[string]schema.Attribute{
							"docs_uri": schema.StringAttribute{
								Description:         "URI for executor docs",
								MarkdownDescription: "URI for executor docs",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"icon_uri": schema.StringAttribute{
								Description:         "URI for executor icon",
								MarkdownDescription: "URI for executor icon",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tooltips": schema.MapAttribute{
								Description:         "executor tooltips",
								MarkdownDescription: "executor tooltips",
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

					"slaves": schema.SingleNestedAttribute{
						Description:         "Slaves data to run test in distributed environment",
						MarkdownDescription: "Slaves data to run test in distributed environment",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
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

					"types": schema.ListAttribute{
						Description:         "Types defines what types can be handled by executor e.g. 'postman/collection', ':curl/command' etc",
						MarkdownDescription: "Types defines what types can be handled by executor e.g. 'postman/collection', ':curl/command' etc",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uri": schema.StringAttribute{
						Description:         "URI for rest based executors",
						MarkdownDescription: "URI for rest based executors",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_data_dir_as_working_dir": schema.BoolAttribute{
						Description:         "use data dir as working dir for executor",
						MarkdownDescription: "use data dir as working dir for executor",
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

func (r *ExecutorTestkubeIoExecutorV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_executor_testkube_io_executor_v1_manifest")

	var model ExecutorTestkubeIoExecutorV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("executor.testkube.io/v1")
	model.Kind = pointer.String("Executor")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
