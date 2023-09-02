/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package slo_koordinator_sh_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SloKoordinatorShNodeMetricV1Alpha1Manifest{}
)

func NewSloKoordinatorShNodeMetricV1Alpha1Manifest() datasource.DataSource {
	return &SloKoordinatorShNodeMetricV1Alpha1Manifest{}
}

type SloKoordinatorShNodeMetricV1Alpha1Manifest struct{}

type SloKoordinatorShNodeMetricV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		MetricCollectPolicy *struct {
			AggregateDurationSeconds *int64 `tfsdk:"aggregate_duration_seconds" json:"aggregateDurationSeconds,omitempty"`
			NodeAggregatePolicy      *struct {
				Durations *[]string `tfsdk:"durations" json:"durations,omitempty"`
			} `tfsdk:"node_aggregate_policy" json:"nodeAggregatePolicy,omitempty"`
			NodeMemoryCollectPolicy *string `tfsdk:"node_memory_collect_policy" json:"nodeMemoryCollectPolicy,omitempty"`
			ReportIntervalSeconds   *int64  `tfsdk:"report_interval_seconds" json:"reportIntervalSeconds,omitempty"`
		} `tfsdk:"metric_collect_policy" json:"metricCollectPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SloKoordinatorShNodeMetricV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_slo_koordinator_sh_node_metric_v1alpha1_manifest"
}

func (r *SloKoordinatorShNodeMetricV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeMetric is the Schema for the nodemetrics API",
		MarkdownDescription: "NodeMetric is the Schema for the nodemetrics API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "NodeMetricSpec defines the desired state of NodeMetric",
				MarkdownDescription: "NodeMetricSpec defines the desired state of NodeMetric",
				Attributes: map[string]schema.Attribute{
					"metric_collect_policy": schema.SingleNestedAttribute{
						Description:         "CollectPolicy defines the Metric collection policy",
						MarkdownDescription: "CollectPolicy defines the Metric collection policy",
						Attributes: map[string]schema.Attribute{
							"aggregate_duration_seconds": schema.Int64Attribute{
								Description:         "AggregateDurationSeconds represents the aggregation period in seconds",
								MarkdownDescription: "AggregateDurationSeconds represents the aggregation period in seconds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_aggregate_policy": schema.SingleNestedAttribute{
								Description:         "NodeAggregatePolicy represents the target grain of node aggregated usage",
								MarkdownDescription: "NodeAggregatePolicy represents the target grain of node aggregated usage",
								Attributes: map[string]schema.Attribute{
									"durations": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"node_memory_collect_policy": schema.StringAttribute{
								Description:         "NodeMemoryPolicy represents apply which method collect memory info",
								MarkdownDescription: "NodeMemoryPolicy represents apply which method collect memory info",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("usageWithHotPageCache", "usageWithoutPageCache", "usageWithPageCache"),
								},
							},

							"report_interval_seconds": schema.Int64Attribute{
								Description:         "ReportIntervalSeconds represents the report period in seconds",
								MarkdownDescription: "ReportIntervalSeconds represents the report period in seconds",
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

func (r *SloKoordinatorShNodeMetricV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_slo_koordinator_sh_node_metric_v1alpha1_manifest")

	var model SloKoordinatorShNodeMetricV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("slo.koordinator.sh/v1alpha1")
	model.Kind = pointer.String("NodeMetric")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
