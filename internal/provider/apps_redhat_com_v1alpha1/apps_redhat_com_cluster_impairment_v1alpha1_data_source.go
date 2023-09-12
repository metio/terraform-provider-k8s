/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_redhat_com_v1alpha1

import (
	"context"
	"encoding/json"
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
	_ datasource.DataSource              = &AppsRedhatComClusterImpairmentV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &AppsRedhatComClusterImpairmentV1Alpha1DataSource{}
)

func NewAppsRedhatComClusterImpairmentV1Alpha1DataSource() datasource.DataSource {
	return &AppsRedhatComClusterImpairmentV1Alpha1DataSource{}
}

type AppsRedhatComClusterImpairmentV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppsRedhatComClusterImpairmentV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Duration *int64 `tfsdk:"duration" json:"duration,omitempty"`
		Egress   *struct {
			Bandwidth         *int64   `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
			Corruption        *float64 `tfsdk:"corruption" json:"corruption,omitempty"`
			CorruptionOptions *struct {
				Correlation *float64 `tfsdk:"correlation" json:"correlation,omitempty"`
			} `tfsdk:"corruption_options" json:"corruptionOptions,omitempty"`
			Duplication        *float64 `tfsdk:"duplication" json:"duplication,omitempty"`
			DuplicationOptions *struct {
				Correlation *float64 `tfsdk:"correlation" json:"correlation,omitempty"`
			} `tfsdk:"duplication_options" json:"duplicationOptions,omitempty"`
			Latency        *float64 `tfsdk:"latency" json:"latency,omitempty"`
			LatencyOptions *struct {
				Distribution       *string  `tfsdk:"distribution" json:"distribution,omitempty"`
				Jitter             *float64 `tfsdk:"jitter" json:"jitter,omitempty"`
				JitterCorrelation  *float64 `tfsdk:"jitter_correlation" json:"jitterCorrelation,omitempty"`
				Reorder            *float64 `tfsdk:"reorder" json:"reorder,omitempty"`
				ReorderCorrelation *float64 `tfsdk:"reorder_correlation" json:"reorderCorrelation,omitempty"`
			} `tfsdk:"latency_options" json:"latencyOptions,omitempty"`
			Loss        *float64 `tfsdk:"loss" json:"loss,omitempty"`
			LossOptions *struct {
				Correlation *float64 `tfsdk:"correlation" json:"correlation,omitempty"`
			} `tfsdk:"loss_options" json:"lossOptions,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		Ingress *struct {
			Bandwidth         *int64   `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
			Corruption        *float64 `tfsdk:"corruption" json:"corruption,omitempty"`
			CorruptionOptions *struct {
				Correlation *float64 `tfsdk:"correlation" json:"correlation,omitempty"`
			} `tfsdk:"corruption_options" json:"corruptionOptions,omitempty"`
			Duplication        *float64 `tfsdk:"duplication" json:"duplication,omitempty"`
			DuplicationOptions *struct {
				Correlation *float64 `tfsdk:"correlation" json:"correlation,omitempty"`
			} `tfsdk:"duplication_options" json:"duplicationOptions,omitempty"`
			Latency        *float64 `tfsdk:"latency" json:"latency,omitempty"`
			LatencyOptions *struct {
				Distribution       *string  `tfsdk:"distribution" json:"distribution,omitempty"`
				Jitter             *float64 `tfsdk:"jitter" json:"jitter,omitempty"`
				JitterCorrelation  *float64 `tfsdk:"jitter_correlation" json:"jitterCorrelation,omitempty"`
				Reorder            *float64 `tfsdk:"reorder" json:"reorder,omitempty"`
				ReorderCorrelation *float64 `tfsdk:"reorder_correlation" json:"reorderCorrelation,omitempty"`
			} `tfsdk:"latency_options" json:"latencyOptions,omitempty"`
			Loss        *float64 `tfsdk:"loss" json:"loss,omitempty"`
			LossOptions *struct {
				Correlation *float64 `tfsdk:"correlation" json:"correlation,omitempty"`
			} `tfsdk:"loss_options" json:"lossOptions,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Interfaces   *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
		LinkFlapping *struct {
			DownTime *int64 `tfsdk:"down_time" json:"downTime,omitempty"`
			Enable   *bool  `tfsdk:"enable" json:"enable,omitempty"`
			UpTime   *int64 `tfsdk:"up_time" json:"upTime,omitempty"`
		} `tfsdk:"link_flapping" json:"linkFlapping,omitempty"`
		NodeSelector *struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		StartDelay *int64 `tfsdk:"start_delay" json:"startDelay,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_redhat_com_cluster_impairment_v1alpha1"
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterImpairment is the Schema for the clusterimpairments API",
		MarkdownDescription: "ClusterImpairment is the Schema for the clusterimpairments API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "Spec defines the desired state of ClusterImpairment",
				MarkdownDescription: "Spec defines the desired state of ClusterImpairment",
				Attributes: map[string]schema.Attribute{
					"duration": schema.Int64Attribute{
						Description:         "The duration of the impairment in seconds.",
						MarkdownDescription: "The duration of the impairment in seconds.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"egress": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the egress impairments.",
						MarkdownDescription: "The configuration section that specifies the egress impairments.",
						Attributes: map[string]schema.Attribute{
							"bandwidth": schema.Int64Attribute{
								Description:         "The bandwidth limit in kbit/sec",
								MarkdownDescription: "The bandwidth limit in kbit/sec",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"corruption": schema.Float64Attribute{
								Description:         "The percent of packets that are corrupted",
								MarkdownDescription: "The percent of packets that are corrupted",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"corruption_options": schema.SingleNestedAttribute{
								Description:         "Advanced corruption options",
								MarkdownDescription: "Advanced corruption options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential corruption values",
										MarkdownDescription: "The correlation between sequential corruption values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"duplication": schema.Float64Attribute{
								Description:         "The percent of packets duplicated",
								MarkdownDescription: "The percent of packets duplicated",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"duplication_options": schema.SingleNestedAttribute{
								Description:         "Advanced duplication options",
								MarkdownDescription: "Advanced duplication options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential duplication values",
										MarkdownDescription: "The correlation between sequential duplication values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"latency": schema.Float64Attribute{
								Description:         "The latency applied in ms",
								MarkdownDescription: "The latency applied in ms",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"latency_options": schema.SingleNestedAttribute{
								Description:         "Advanced latency options. Example: jitter",
								MarkdownDescription: "Advanced latency options. Example: jitter",
								Attributes: map[string]schema.Attribute{
									"distribution": schema.StringAttribute{
										Description:         "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										MarkdownDescription: "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"jitter": schema.Float64Attribute{
										Description:         "Variation in the latency that follows the specified distribution.",
										MarkdownDescription: "Variation in the latency that follows the specified distribution.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"jitter_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential jitter values",
										MarkdownDescription: "The correlation between sequential jitter values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"reorder": schema.Float64Attribute{
										Description:         "The percentage of packets that are not delayed, causing reordering",
										MarkdownDescription: "The percentage of packets that are not delayed, causing reordering",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"reorder_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential reorder values",
										MarkdownDescription: "The correlation between sequential reorder values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"loss": schema.Float64Attribute{
								Description:         "The packet loss in percent",
								MarkdownDescription: "The packet loss in percent",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"loss_options": schema.SingleNestedAttribute{
								Description:         "Advanced packet loss options",
								MarkdownDescription: "Advanced packet loss options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential packet loss values",
										MarkdownDescription: "The correlation between sequential packet loss values",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the ingress impairments.",
						MarkdownDescription: "The configuration section that specifies the ingress impairments.",
						Attributes: map[string]schema.Attribute{
							"bandwidth": schema.Int64Attribute{
								Description:         "The bandwidth limit in kbit/sec",
								MarkdownDescription: "The bandwidth limit in kbit/sec",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"corruption": schema.Float64Attribute{
								Description:         "The percent of packets that are corrupted",
								MarkdownDescription: "The percent of packets that are corrupted",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"corruption_options": schema.SingleNestedAttribute{
								Description:         "Advanced corruption options",
								MarkdownDescription: "Advanced corruption options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential corruption values",
										MarkdownDescription: "The correlation between sequential corruption values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"duplication": schema.Float64Attribute{
								Description:         "The percent of packets duplicated",
								MarkdownDescription: "The percent of packets duplicated",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"duplication_options": schema.SingleNestedAttribute{
								Description:         "Advanced duplication options",
								MarkdownDescription: "Advanced duplication options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential duplication values",
										MarkdownDescription: "The correlation between sequential duplication values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"latency": schema.Float64Attribute{
								Description:         "The latency applied in ms",
								MarkdownDescription: "The latency applied in ms",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"latency_options": schema.SingleNestedAttribute{
								Description:         "Advanced latency options. Example: jitter",
								MarkdownDescription: "Advanced latency options. Example: jitter",
								Attributes: map[string]schema.Attribute{
									"distribution": schema.StringAttribute{
										Description:         "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										MarkdownDescription: "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"jitter": schema.Float64Attribute{
										Description:         "Variation in the latency that follows the specified distribution.",
										MarkdownDescription: "Variation in the latency that follows the specified distribution.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"jitter_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential jitter values",
										MarkdownDescription: "The correlation between sequential jitter values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"reorder": schema.Float64Attribute{
										Description:         "The percentage of packets that are not delayed, causing reordering",
										MarkdownDescription: "The percentage of packets that are not delayed, causing reordering",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"reorder_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential reorder values",
										MarkdownDescription: "The correlation between sequential reorder values",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"loss": schema.Float64Attribute{
								Description:         "The packet loss in percent",
								MarkdownDescription: "The packet loss in percent",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"loss_options": schema.SingleNestedAttribute{
								Description:         "Advanced packet loss options",
								MarkdownDescription: "Advanced packet loss options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential packet loss values",
										MarkdownDescription: "The correlation between sequential packet loss values",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"interfaces": schema.ListAttribute{
						Description:         "All interfaces that the impairments should be applied to. Must be valid interfaces or the impairments will fail to apply.",
						MarkdownDescription: "All interfaces that the impairments should be applied to. Must be valid interfaces or the impairments will fail to apply.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"link_flapping": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the link flapping impairments.",
						MarkdownDescription: "The configuration section that specifies the link flapping impairments.",
						Attributes: map[string]schema.Attribute{
							"down_time": schema.Int64Attribute{
								Description:         "The duration that the link should be disabled.",
								MarkdownDescription: "The duration that the link should be disabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable": schema.BoolAttribute{
								Description:         "Whether to enable link flapping.",
								MarkdownDescription: "Whether to enable link flapping.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"up_time": schema.Int64Attribute{
								Description:         "The duration that the link should be enabled.",
								MarkdownDescription: "The duration that the link should be enabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"node_selector": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the node selector that should be applied to the daemonset. Default: worker nodes.",
						MarkdownDescription: "The configuration section that specifies the node selector that should be applied to the daemonset. Default: worker nodes.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key for the node selector",
								MarkdownDescription: "The key for the node selector",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"value": schema.StringAttribute{
								Description:         "The value for the node selector",
								MarkdownDescription: "The value for the node selector",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"start_delay": schema.Int64Attribute{
						Description:         "The delay (in seconds) before starting the impairments. At least 5 seconds recommended for Kubernetes and for synchronization of the impairments.",
						MarkdownDescription: "The delay (in seconds) before starting the impairments. At least 5 seconds recommended for Kubernetes and for synchronization of the impairments.",
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

func (r *AppsRedhatComClusterImpairmentV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var data AppsRedhatComClusterImpairmentV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.redhat.com", Version: "v1alpha1", Resource: "clusterimpairments"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AppsRedhatComClusterImpairmentV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("apps.redhat.com/v1alpha1")
	data.Kind = pointer.String("ClusterImpairment")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
