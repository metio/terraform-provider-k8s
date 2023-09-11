/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_redhat_com_v1alpha1

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
	"time"
)

var (
	_ resource.Resource                = &AppsRedhatComClusterImpairmentV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &AppsRedhatComClusterImpairmentV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &AppsRedhatComClusterImpairmentV1Alpha1Resource{}
)

func NewAppsRedhatComClusterImpairmentV1Alpha1Resource() resource.Resource {
	return &AppsRedhatComClusterImpairmentV1Alpha1Resource{}
}

type AppsRedhatComClusterImpairmentV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type AppsRedhatComClusterImpairmentV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_redhat_com_cluster_impairment_v1alpha1"
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"wait_for_upsert": schema.ListNestedAttribute{
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
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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
				Description:         "Spec defines the desired state of ClusterImpairment",
				MarkdownDescription: "Spec defines the desired state of ClusterImpairment",
				Attributes: map[string]schema.Attribute{
					"duration": schema.Int64Attribute{
						Description:         "The duration of the impairment in seconds.",
						MarkdownDescription: "The duration of the impairment in seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"egress": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the egress impairments.",
						MarkdownDescription: "The configuration section that specifies the egress impairments.",
						Attributes: map[string]schema.Attribute{
							"bandwidth": schema.Int64Attribute{
								Description:         "The bandwidth limit in kbit/sec",
								MarkdownDescription: "The bandwidth limit in kbit/sec",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"corruption": schema.Float64Attribute{
								Description:         "The percent of packets that are corrupted",
								MarkdownDescription: "The percent of packets that are corrupted",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"corruption_options": schema.SingleNestedAttribute{
								Description:         "Advanced corruption options",
								MarkdownDescription: "Advanced corruption options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential corruption values",
										MarkdownDescription: "The correlation between sequential corruption values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplication": schema.Float64Attribute{
								Description:         "The percent of packets duplicated",
								MarkdownDescription: "The percent of packets duplicated",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"duplication_options": schema.SingleNestedAttribute{
								Description:         "Advanced duplication options",
								MarkdownDescription: "Advanced duplication options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential duplication values",
										MarkdownDescription: "The correlation between sequential duplication values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": schema.Float64Attribute{
								Description:         "The latency applied in ms",
								MarkdownDescription: "The latency applied in ms",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"latency_options": schema.SingleNestedAttribute{
								Description:         "Advanced latency options. Example: jitter",
								MarkdownDescription: "Advanced latency options. Example: jitter",
								Attributes: map[string]schema.Attribute{
									"distribution": schema.StringAttribute{
										Description:         "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										MarkdownDescription: "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jitter": schema.Float64Attribute{
										Description:         "Variation in the latency that follows the specified distribution.",
										MarkdownDescription: "Variation in the latency that follows the specified distribution.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jitter_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential jitter values",
										MarkdownDescription: "The correlation between sequential jitter values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reorder": schema.Float64Attribute{
										Description:         "The percentage of packets that are not delayed, causing reordering",
										MarkdownDescription: "The percentage of packets that are not delayed, causing reordering",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reorder_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential reorder values",
										MarkdownDescription: "The correlation between sequential reorder values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss": schema.Float64Attribute{
								Description:         "The packet loss in percent",
								MarkdownDescription: "The packet loss in percent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"loss_options": schema.SingleNestedAttribute{
								Description:         "Advanced packet loss options",
								MarkdownDescription: "Advanced packet loss options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential packet loss values",
										MarkdownDescription: "The correlation between sequential packet loss values",
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

					"ingress": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the ingress impairments.",
						MarkdownDescription: "The configuration section that specifies the ingress impairments.",
						Attributes: map[string]schema.Attribute{
							"bandwidth": schema.Int64Attribute{
								Description:         "The bandwidth limit in kbit/sec",
								MarkdownDescription: "The bandwidth limit in kbit/sec",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"corruption": schema.Float64Attribute{
								Description:         "The percent of packets that are corrupted",
								MarkdownDescription: "The percent of packets that are corrupted",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"corruption_options": schema.SingleNestedAttribute{
								Description:         "Advanced corruption options",
								MarkdownDescription: "Advanced corruption options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential corruption values",
										MarkdownDescription: "The correlation between sequential corruption values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplication": schema.Float64Attribute{
								Description:         "The percent of packets duplicated",
								MarkdownDescription: "The percent of packets duplicated",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"duplication_options": schema.SingleNestedAttribute{
								Description:         "Advanced duplication options",
								MarkdownDescription: "Advanced duplication options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential duplication values",
										MarkdownDescription: "The correlation between sequential duplication values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": schema.Float64Attribute{
								Description:         "The latency applied in ms",
								MarkdownDescription: "The latency applied in ms",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"latency_options": schema.SingleNestedAttribute{
								Description:         "Advanced latency options. Example: jitter",
								MarkdownDescription: "Advanced latency options. Example: jitter",
								Attributes: map[string]schema.Attribute{
									"distribution": schema.StringAttribute{
										Description:         "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										MarkdownDescription: "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jitter": schema.Float64Attribute{
										Description:         "Variation in the latency that follows the specified distribution.",
										MarkdownDescription: "Variation in the latency that follows the specified distribution.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jitter_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential jitter values",
										MarkdownDescription: "The correlation between sequential jitter values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reorder": schema.Float64Attribute{
										Description:         "The percentage of packets that are not delayed, causing reordering",
										MarkdownDescription: "The percentage of packets that are not delayed, causing reordering",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reorder_correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential reorder values",
										MarkdownDescription: "The correlation between sequential reorder values",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss": schema.Float64Attribute{
								Description:         "The packet loss in percent",
								MarkdownDescription: "The packet loss in percent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"loss_options": schema.SingleNestedAttribute{
								Description:         "Advanced packet loss options",
								MarkdownDescription: "Advanced packet loss options",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.Float64Attribute{
										Description:         "The correlation between sequential packet loss values",
										MarkdownDescription: "The correlation between sequential packet loss values",
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

					"interfaces": schema.ListAttribute{
						Description:         "All interfaces that the impairments should be applied to. Must be valid interfaces or the impairments will fail to apply.",
						MarkdownDescription: "All interfaces that the impairments should be applied to. Must be valid interfaces or the impairments will fail to apply.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"link_flapping": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the link flapping impairments.",
						MarkdownDescription: "The configuration section that specifies the link flapping impairments.",
						Attributes: map[string]schema.Attribute{
							"down_time": schema.Int64Attribute{
								Description:         "The duration that the link should be disabled.",
								MarkdownDescription: "The duration that the link should be disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Whether to enable link flapping.",
								MarkdownDescription: "Whether to enable link flapping.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"up_time": schema.Int64Attribute{
								Description:         "The duration that the link should be enabled.",
								MarkdownDescription: "The duration that the link should be enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": schema.SingleNestedAttribute{
						Description:         "The configuration section that specifies the node selector that should be applied to the daemonset. Default: worker nodes.",
						MarkdownDescription: "The configuration section that specifies the node selector that should be applied to the daemonset. Default: worker nodes.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key for the node selector",
								MarkdownDescription: "The key for the node selector",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"value": schema.StringAttribute{
								Description:         "The value for the node selector",
								MarkdownDescription: "The value for the node selector",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"start_delay": schema.Int64Attribute{
						Description:         "The delay (in seconds) before starting the impairments. At least 5 seconds recommended for Kubernetes and for synchronization of the impairments.",
						MarkdownDescription: "The delay (in seconds) before starting the impairments. At least 5 seconds recommended for Kubernetes and for synchronization of the impairments.",
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

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var model AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("apps.redhat.com/v1alpha1")
	model.Kind = pointer.String("ClusterImpairment")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.redhat.com", Version: "v1alpha1", Resource: "clusterimpairments"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var data AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
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

	var readResponse AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var model AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.redhat.com/v1alpha1")
	model.Kind = pointer.String("ClusterImpairment")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.redhat.com", Version: "v1alpha1", Resource: "clusterimpairments"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var data AppsRedhatComClusterImpairmentV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.redhat.com", Version: "v1alpha1", Resource: "clusterimpairments"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "apps.redhat.com", Version: "v1alpha1", Resource: "clusterimpairments"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
