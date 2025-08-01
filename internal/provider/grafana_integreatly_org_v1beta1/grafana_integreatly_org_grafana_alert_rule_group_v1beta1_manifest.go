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
	_ datasource.DataSource = &GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest{}
)

func NewGrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest() datasource.DataSource {
	return &GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest{}
}

type GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest struct{}

type GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1ManifestData struct {
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
		AllowCrossNamespaceImport *bool   `tfsdk:"allow_cross_namespace_import" json:"allowCrossNamespaceImport,omitempty"`
		Editable                  *bool   `tfsdk:"editable" json:"editable,omitempty"`
		FolderRef                 *string `tfsdk:"folder_ref" json:"folderRef,omitempty"`
		FolderUID                 *string `tfsdk:"folder_uid" json:"folderUID,omitempty"`
		InstanceSelector          *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"instance_selector" json:"instanceSelector,omitempty"`
		Interval     *string `tfsdk:"interval" json:"interval,omitempty"`
		Name         *string `tfsdk:"name" json:"name,omitempty"`
		ResyncPeriod *string `tfsdk:"resync_period" json:"resyncPeriod,omitempty"`
		Rules        *[]struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Condition   *string            `tfsdk:"condition" json:"condition,omitempty"`
			Data        *[]struct {
				DatasourceUid     *string            `tfsdk:"datasource_uid" json:"datasourceUid,omitempty"`
				Model             *map[string]string `tfsdk:"model" json:"model,omitempty"`
				QueryType         *string            `tfsdk:"query_type" json:"queryType,omitempty"`
				RefId             *string            `tfsdk:"ref_id" json:"refId,omitempty"`
				RelativeTimeRange *struct {
					From *int64 `tfsdk:"from" json:"from,omitempty"`
					To   *int64 `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"relative_time_range" json:"relativeTimeRange,omitempty"`
			} `tfsdk:"data" json:"data,omitempty"`
			ExecErrState                *string            `tfsdk:"exec_err_state" json:"execErrState,omitempty"`
			For                         *string            `tfsdk:"for" json:"for,omitempty"`
			IsPaused                    *bool              `tfsdk:"is_paused" json:"isPaused,omitempty"`
			KeepFiringFor               *string            `tfsdk:"keep_firing_for" json:"keepFiringFor,omitempty"`
			Labels                      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MissingSeriesEvalsToResolve *int64             `tfsdk:"missing_series_evals_to_resolve" json:"missingSeriesEvalsToResolve,omitempty"`
			NoDataState                 *string            `tfsdk:"no_data_state" json:"noDataState,omitempty"`
			NotificationSettings        *struct {
				Group_by            *[]string `tfsdk:"group_by" json:"group_by,omitempty"`
				Group_interval      *string   `tfsdk:"group_interval" json:"group_interval,omitempty"`
				Group_wait          *string   `tfsdk:"group_wait" json:"group_wait,omitempty"`
				Mute_time_intervals *[]string `tfsdk:"mute_time_intervals" json:"mute_time_intervals,omitempty"`
				Receiver            *string   `tfsdk:"receiver" json:"receiver,omitempty"`
				Repeat_interval     *string   `tfsdk:"repeat_interval" json:"repeat_interval,omitempty"`
			} `tfsdk:"notification_settings" json:"notificationSettings,omitempty"`
			Record *struct {
				From   *string `tfsdk:"from" json:"from,omitempty"`
				Metric *string `tfsdk:"metric" json:"metric,omitempty"`
			} `tfsdk:"record" json:"record,omitempty"`
			Title *string `tfsdk:"title" json:"title,omitempty"`
			Uid   *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		Suspend *bool `tfsdk:"suspend" json:"suspend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_grafana_integreatly_org_grafana_alert_rule_group_v1beta1_manifest"
}

func (r *GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GrafanaAlertRuleGroup is the Schema for the grafanaalertrulegroups API",
		MarkdownDescription: "GrafanaAlertRuleGroup is the Schema for the grafanaalertrulegroups API",
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
				Description:         "GrafanaAlertRuleGroupSpec defines the desired state of GrafanaAlertRuleGroup",
				MarkdownDescription: "GrafanaAlertRuleGroupSpec defines the desired state of GrafanaAlertRuleGroup",
				Attributes: map[string]schema.Attribute{
					"allow_cross_namespace_import": schema.BoolAttribute{
						Description:         "Allow the Operator to match this resource with Grafanas outside the current namespace",
						MarkdownDescription: "Allow the Operator to match this resource with Grafanas outside the current namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"editable": schema.BoolAttribute{
						Description:         "Whether to enable or disable editing of the alert rule group in Grafana UI",
						MarkdownDescription: "Whether to enable or disable editing of the alert rule group in Grafana UI",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"folder_ref": schema.StringAttribute{
						Description:         "Match GrafanaFolders CRs to infer the uid",
						MarkdownDescription: "Match GrafanaFolders CRs to infer the uid",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"folder_uid": schema.StringAttribute{
						Description:         "UID of the folder containing this rule group Overrides the FolderSelector",
						MarkdownDescription: "UID of the folder containing this rule group Overrides the FolderSelector",
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

					"interval": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"name": schema.StringAttribute{
						Description:         "Name of the alert rule group. If not specified, the resource name will be used.",
						MarkdownDescription: "Name of the alert rule group. If not specified, the resource name will be used.",
						Required:            false,
						Optional:            true,
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

					"rules": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotations": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"condition": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"data": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"datasource_uid": schema.StringAttribute{
												Description:         "Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.",
												MarkdownDescription: "Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"model": schema.MapAttribute{
												Description:         "JSON is the raw JSON query and includes the above properties as well as custom properties.",
												MarkdownDescription: "JSON is the raw JSON query and includes the above properties as well as custom properties.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_type": schema.StringAttribute{
												Description:         "QueryType is an optional identifier for the type of query. It can be used to distinguish different types of queries.",
												MarkdownDescription: "QueryType is an optional identifier for the type of query. It can be used to distinguish different types of queries.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ref_id": schema.StringAttribute{
												Description:         "RefID is the unique identifier of the query, set by the frontend call.",
												MarkdownDescription: "RefID is the unique identifier of the query, set by the frontend call.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"relative_time_range": schema.SingleNestedAttribute{
												Description:         "relative time range",
												MarkdownDescription: "relative time range",
												Attributes: map[string]schema.Attribute{
													"from": schema.Int64Attribute{
														Description:         "from",
														MarkdownDescription: "from",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"to": schema.Int64Attribute{
														Description:         "to",
														MarkdownDescription: "to",
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
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"exec_err_state": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("OK", "Alerting", "Error", "KeepLast"),
									},
								},

								"for": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
									},
								},

								"is_paused": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"keep_firing_for": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
									},
								},

								"labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"missing_series_evals_to_resolve": schema.Int64Attribute{
									Description:         "The number of missing series evaluations that must occur before the rule is considered to be resolved.",
									MarkdownDescription: "The number of missing series evaluations that must occur before the rule is considered to be resolved.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"no_data_state": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Alerting", "NoData", "OK", "KeepLast"),
									},
								},

								"notification_settings": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"group_by": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_interval": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_wait": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mute_time_intervals": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"receiver": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"repeat_interval": schema.StringAttribute{
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

								"record": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"from": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"metric": schema.StringAttribute{
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

								"title": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(190),
									},
								},

								"uid": schema.StringAttribute{
									Description:         "UID of the alert rule. Can be any string consisting of alphanumeric characters, - and _ with a maximum length of 40",
									MarkdownDescription: "UID of the alert rule. Can be any string consisting of alphanumeric characters, - and _ with a maximum length of 40",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(40),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_]+$`), ""),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend pauses synchronizing attempts and tells the operator to ignore changes",
						MarkdownDescription: "Suspend pauses synchronizing attempts and tells the operator to ignore changes",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_grafana_integreatly_org_grafana_alert_rule_group_v1beta1_manifest")

	var model GrafanaIntegreatlyOrgGrafanaAlertRuleGroupV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("grafana.integreatly.org/v1beta1")
	model.Kind = pointer.String("GrafanaAlertRuleGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
