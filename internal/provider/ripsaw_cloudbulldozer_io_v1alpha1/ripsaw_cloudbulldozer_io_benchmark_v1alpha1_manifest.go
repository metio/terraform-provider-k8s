/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ripsaw_cloudbulldozer_io_v1alpha1

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
	_ datasource.DataSource = &RipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest{}
)

func NewRipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest() datasource.DataSource {
	return &RipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest{}
}

type RipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest struct{}

type RipsawCloudbulldozerIoBenchmarkV1Alpha1ManifestData struct {
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
		Cleanup              *bool   `tfsdk:"cleanup" json:"cleanup,omitempty"`
		Clustername          *string `tfsdk:"clustername" json:"clustername,omitempty"`
		Drop_cache_kernel    *bool   `tfsdk:"drop_cache_kernel" json:"drop_cache_kernel,omitempty"`
		Drop_cache_rook_ceph *bool   `tfsdk:"drop_cache_rook_ceph" json:"drop_cache_rook_ceph,omitempty"`
		Elasticsearch        *struct {
			Index_name  *string `tfsdk:"index_name" json:"index_name,omitempty"`
			Parallel    *bool   `tfsdk:"parallel" json:"parallel,omitempty"`
			Url         *string `tfsdk:"url" json:"url,omitempty"`
			Verify_cert *bool   `tfsdk:"verify_cert" json:"verify_cert,omitempty"`
		} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
		Global_overrides *[]string            `tfsdk:"global_overrides" json:"global_overrides,omitempty"`
		Hostpath         *string              `tfsdk:"hostpath" json:"hostpath,omitempty"`
		Job_params       *[]map[string]string `tfsdk:"job_params" json:"job_params,omitempty"`
		Metadata         *struct {
			Collection        *bool              `tfsdk:"collection" json:"collection,omitempty"`
			Force             *bool              `tfsdk:"force" json:"force,omitempty"`
			Image             *string            `tfsdk:"image" json:"image,omitempty"`
			Label             *map[string]string `tfsdk:"label" json:"label,omitempty"`
			Privileged        *bool              `tfsdk:"privileged" json:"privileged,omitempty"`
			Ssl               *bool              `tfsdk:"ssl" json:"ssl,omitempty"`
			StockpileSkipTags *[]string          `tfsdk:"stockpile_skip_tags" json:"stockpileSkipTags,omitempty"`
			StockpileTags     *[]string          `tfsdk:"stockpile_tags" json:"stockpileTags,omitempty"`
			Targeted          *bool              `tfsdk:"targeted" json:"targeted,omitempty"`
		} `tfsdk:"metadata" json:"metadata,omitempty"`
		Prometheus *struct {
			Es_parallel *bool   `tfsdk:"es_parallel" json:"es_parallel,omitempty"`
			Es_url      *string `tfsdk:"es_url" json:"es_url,omitempty"`
			Prom_token  *string `tfsdk:"prom_token" json:"prom_token,omitempty"`
			Prom_url    *string `tfsdk:"prom_url" json:"prom_url,omitempty"`
		} `tfsdk:"prometheus" json:"prometheus,omitempty"`
		Snappy *struct {
			Password *string `tfsdk:"password" json:"password,omitempty"`
			Url      *string `tfsdk:"url" json:"url,omitempty"`
			User     *string `tfsdk:"user" json:"user,omitempty"`
		} `tfsdk:"snappy" json:"snappy,omitempty"`
		System_metrics *struct {
			Collection      *bool   `tfsdk:"collection" json:"collection,omitempty"`
			Es_url          *string `tfsdk:"es_url" json:"es_url,omitempty"`
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			Index_name      *string `tfsdk:"index_name" json:"index_name,omitempty"`
			Metrics_profile *string `tfsdk:"metrics_profile" json:"metrics_profile,omitempty"`
			Prom_token      *string `tfsdk:"prom_token" json:"prom_token,omitempty"`
			Prom_url        *string `tfsdk:"prom_url" json:"prom_url,omitempty"`
			Step            *string `tfsdk:"step" json:"step,omitempty"`
		} `tfsdk:"system_metrics" json:"system_metrics,omitempty"`
		Test_user *string `tfsdk:"test_user" json:"test_user,omitempty"`
		Uuid      *string `tfsdk:"uuid" json:"uuid,omitempty"`
		Workload  *struct {
			Args *struct {
				Debug       *bool  `tfsdk:"debug" json:"debug,omitempty"`
				Job_timeout *int64 `tfsdk:"job_timeout" json:"job_timeout,omitempty"`
			} `tfsdk:"args" json:"args,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"workload" json:"workload,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest"
}

func (r *RipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"cleanup": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clustername": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"drop_cache_kernel": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"drop_cache_rook_ceph": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"elasticsearch": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"index_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parallel": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_cert": schema.BoolAttribute{
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

					"global_overrides": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hostpath": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"job_params": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"collection": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"force": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"privileged": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stockpile_skip_tags": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stockpile_tags": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"targeted": schema.BoolAttribute{
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

					"prometheus": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"es_parallel": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"es_url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prom_token": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prom_url": schema.StringAttribute{
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

					"snappy": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"password": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user": schema.StringAttribute{
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

					"system_metrics": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"collection": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"es_url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"index_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_profile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prom_token": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prom_url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"step": schema.StringAttribute{
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

					"test_user": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uuid": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"args": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"debug": schema.BoolAttribute{
										Description:         "Set debug log level in the workload",
										MarkdownDescription: "Set debug log level in the workload",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"job_timeout": schema.Int64Attribute{
										Description:         "Timeout used for jobs deployed by this benchmark",
										MarkdownDescription: "Timeout used for jobs deployed by this benchmark",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": schema.StringAttribute{
								Description:         "Workload name",
								MarkdownDescription: "Workload name",
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

func (r *RipsawCloudbulldozerIoBenchmarkV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest")

	var model RipsawCloudbulldozerIoBenchmarkV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("ripsaw.cloudbulldozer.io/v1alpha1")
	model.Kind = pointer.String("Benchmark")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
