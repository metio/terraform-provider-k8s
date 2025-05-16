/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package grafana_integreatly_org_v1beta1

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
	_ datasource.DataSource = &GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest{}
)

func NewGrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest() datasource.DataSource {
	return &GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest{}
}

type GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest struct{}

type GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1ManifestData struct {
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
		AllowCrossNamespaceImport *bool `tfsdk:"allow_cross_namespace_import" json:"allowCrossNamespaceImport,omitempty"`
		Editable                  *bool `tfsdk:"editable" json:"editable,omitempty"`
		InstanceSelector          *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"instance_selector" json:"instanceSelector,omitempty"`
		Name           *string `tfsdk:"name" json:"name,omitempty"`
		ResyncPeriod   *string `tfsdk:"resync_period" json:"resyncPeriod,omitempty"`
		Time_intervals *[]struct {
			Days_of_month *[]string `tfsdk:"days_of_month" json:"days_of_month,omitempty"`
			Location      *string   `tfsdk:"location" json:"location,omitempty"`
			Months        *[]string `tfsdk:"months" json:"months,omitempty"`
			Times         *[]struct {
				End_time   *string `tfsdk:"end_time" json:"end_time,omitempty"`
				Start_time *string `tfsdk:"start_time" json:"start_time,omitempty"`
			} `tfsdk:"times" json:"times,omitempty"`
			Weekdays *[]string `tfsdk:"weekdays" json:"weekdays,omitempty"`
			Years    *[]string `tfsdk:"years" json:"years,omitempty"`
		} `tfsdk:"time_intervals" json:"time_intervals,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_grafana_integreatly_org_grafana_mute_timing_v1beta1_manifest"
}

func (r *GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GrafanaMuteTiming is the Schema for the GrafanaMuteTiming API",
		MarkdownDescription: "GrafanaMuteTiming is the Schema for the GrafanaMuteTiming API",
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
				Description:         "GrafanaMuteTimingSpec defines the desired state of GrafanaMuteTiming",
				MarkdownDescription: "GrafanaMuteTimingSpec defines the desired state of GrafanaMuteTiming",
				Attributes: map[string]schema.Attribute{
					"allow_cross_namespace_import": schema.BoolAttribute{
						Description:         "Allow the Operator to match this resource with Grafanas outside the current namespace",
						MarkdownDescription: "Allow the Operator to match this resource with Grafanas outside the current namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"editable": schema.BoolAttribute{
						Description:         "Whether to enable or disable editing of the mute timing in Grafana UI",
						MarkdownDescription: "Whether to enable or disable editing of the mute timing in Grafana UI",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_selector": schema.SingleNestedAttribute{
						Description:         "Selects Grafana instances for import",
						MarkdownDescription: "Selects Grafana instances for import",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "A unique name for the mute timing",
						MarkdownDescription: "A unique name for the mute timing",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"resync_period": schema.StringAttribute{
						Description:         "How often the resource is synced, defaults to 10m0s if not set",
						MarkdownDescription: "How often the resource is synced, defaults to 10m0s if not set",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"time_intervals": schema.ListNestedAttribute{
						Description:         "Time intervals for muting",
						MarkdownDescription: "Time intervals for muting",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"days_of_month": schema.ListAttribute{
									Description:         "The date 1-31 of a month. Negative values can also be used to represent days that begin at the end of the month. For example: -1 for the last day of the month.",
									MarkdownDescription: "The date 1-31 of a month. Negative values can also be used to represent days that begin at the end of the month. For example: -1 for the last day of the month.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"location": schema.StringAttribute{
									Description:         "Depending on the location, the time range is displayed in local time.",
									MarkdownDescription: "Depending on the location, the time range is displayed in local time.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"months": schema.ListAttribute{
									Description:         "The months of the year in either numerical or the full calendar month. For example: 1, may.",
									MarkdownDescription: "The months of the year in either numerical or the full calendar month. For example: 1, may.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"times": schema.ListNestedAttribute{
									Description:         "The time inclusive of the start and exclusive of the end time (in UTC if no location has been selected, otherwise local time).",
									MarkdownDescription: "The time inclusive of the start and exclusive of the end time (in UTC if no location has been selected, otherwise local time).",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"end_time": schema.StringAttribute{
												Description:         "end time",
												MarkdownDescription: "end time",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"start_time": schema.StringAttribute{
												Description:         "start time",
												MarkdownDescription: "start time",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"weekdays": schema.ListAttribute{
									Description:         "The day or range of days of the week. For example: monday, thursday",
									MarkdownDescription: "The day or range of days of the week. For example: monday, thursday",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"years": schema.ListAttribute{
									Description:         "The year or years for the interval. For example: 2021",
									MarkdownDescription: "The year or years for the interval. For example: 2021",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
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

func (r *GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_grafana_integreatly_org_grafana_mute_timing_v1beta1_manifest")

	var model GrafanaIntegreatlyOrgGrafanaMuteTimingV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("grafana.integreatly.org/v1beta1")
	model.Kind = pointer.String("GrafanaMuteTiming")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
