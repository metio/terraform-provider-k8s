/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MonitoringCoreosComPrometheusRuleV1Manifest{}
)

func NewMonitoringCoreosComPrometheusRuleV1Manifest() datasource.DataSource {
	return &MonitoringCoreosComPrometheusRuleV1Manifest{}
}

type MonitoringCoreosComPrometheusRuleV1Manifest struct{}

type MonitoringCoreosComPrometheusRuleV1ManifestData struct {
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
		Groups *[]struct {
			Interval                  *string `tfsdk:"interval" json:"interval,omitempty"`
			Limit                     *int64  `tfsdk:"limit" json:"limit,omitempty"`
			Name                      *string `tfsdk:"name" json:"name,omitempty"`
			Partial_response_strategy *string `tfsdk:"partial_response_strategy" json:"partial_response_strategy,omitempty"`
			Rules                     *[]struct {
				Alert           *string            `tfsdk:"alert" json:"alert,omitempty"`
				Annotations     *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Expr            *string            `tfsdk:"expr" json:"expr,omitempty"`
				For             *string            `tfsdk:"for" json:"for,omitempty"`
				Keep_firing_for *string            `tfsdk:"keep_firing_for" json:"keep_firing_for,omitempty"`
				Labels          *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Record          *string            `tfsdk:"record" json:"record,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringCoreosComPrometheusRuleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_coreos_com_prometheus_rule_v1_manifest"
}

func (r *MonitoringCoreosComPrometheusRuleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The 'PrometheusRule' custom resource definition (CRD) defines [alerting](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/) and [recording](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/) rules to be evaluated by 'Prometheus' or 'ThanosRuler' objects.'Prometheus' and 'ThanosRuler' objects select 'PrometheusRule' objects using label and namespace selectors.",
		MarkdownDescription: "The 'PrometheusRule' custom resource definition (CRD) defines [alerting](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/) and [recording](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/) rules to be evaluated by 'Prometheus' or 'ThanosRuler' objects.'Prometheus' and 'ThanosRuler' objects select 'PrometheusRule' objects using label and namespace selectors.",
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
				Description:         "Specification of desired alerting rule definitions for Prometheus.",
				MarkdownDescription: "Specification of desired alerting rule definitions for Prometheus.",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "Content of Prometheus rule file",
						MarkdownDescription: "Content of Prometheus rule file",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{
									Description:         "Interval determines how often rules in the group are evaluated.",
									MarkdownDescription: "Interval determines how often rules in the group are evaluated.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
									},
								},

								"limit": schema.Int64Attribute{
									Description:         "Limit the number of alerts an alerting rule and series a recordingrule can produce.Limit is supported starting with Prometheus >= 2.31 and Thanos Ruler >= 0.24.",
									MarkdownDescription: "Limit the number of alerts an alerting rule and series a recordingrule can produce.Limit is supported starting with Prometheus >= 2.31 and Thanos Ruler >= 0.24.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the rule group.",
									MarkdownDescription: "Name of the rule group.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"partial_response_strategy": schema.StringAttribute{
									Description:         "PartialResponseStrategy is only used by ThanosRuler and willbe ignored by Prometheus instances.More info: https://github.com/thanos-io/thanos/blob/main/docs/components/rule.md#partial-response",
									MarkdownDescription: "PartialResponseStrategy is only used by ThanosRuler and willbe ignored by Prometheus instances.More info: https://github.com/thanos-io/thanos/blob/main/docs/components/rule.md#partial-response",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(abort|warn)?$`), ""),
									},
								},

								"rules": schema.ListNestedAttribute{
									Description:         "List of alerting and recording rules.",
									MarkdownDescription: "List of alerting and recording rules.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alert": schema.StringAttribute{
												Description:         "Name of the alert. Must be a valid label value.Only one of 'record' and 'alert' must be set.",
												MarkdownDescription: "Name of the alert. Must be a valid label value.Only one of 'record' and 'alert' must be set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"annotations": schema.MapAttribute{
												Description:         "Annotations to add to each alert.Only valid for alerting rules.",
												MarkdownDescription: "Annotations to add to each alert.Only valid for alerting rules.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"expr": schema.StringAttribute{
												Description:         "PromQL expression to evaluate.",
												MarkdownDescription: "PromQL expression to evaluate.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"for": schema.StringAttribute{
												Description:         "Alerts are considered firing once they have been returned for this long.",
												MarkdownDescription: "Alerts are considered firing once they have been returned for this long.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
												},
											},

											"keep_firing_for": schema.StringAttribute{
												Description:         "KeepFiringFor defines how long an alert will continue firing after the condition that triggered it has cleared.",
												MarkdownDescription: "KeepFiringFor defines how long an alert will continue firing after the condition that triggered it has cleared.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
												},
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to add or overwrite.",
												MarkdownDescription: "Labels to add or overwrite.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"record": schema.StringAttribute{
												Description:         "Name of the time series to output to. Must be a valid metric name.Only one of 'record' and 'alert' must be set.",
												MarkdownDescription: "Name of the time series to output to. Must be a valid metric name.Only one of 'record' and 'alert' must be set.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *MonitoringCoreosComPrometheusRuleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_prometheus_rule_v1_manifest")

	var model MonitoringCoreosComPrometheusRuleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("monitoring.coreos.com/v1")
	model.Kind = pointer.String("PrometheusRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
