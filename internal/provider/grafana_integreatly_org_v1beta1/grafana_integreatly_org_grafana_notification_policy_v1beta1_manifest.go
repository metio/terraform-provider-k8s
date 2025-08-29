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
	_ datasource.DataSource = &GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest{}
)

func NewGrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest() datasource.DataSource {
	return &GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest{}
}

type GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest struct{}

type GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1ManifestData struct {
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
		ResyncPeriod *string `tfsdk:"resync_period" json:"resyncPeriod,omitempty"`
		Route        *struct {
			Continue       *bool              `tfsdk:"continue" json:"continue,omitempty"`
			Group_by       *[]string          `tfsdk:"group_by" json:"group_by,omitempty"`
			Group_interval *string            `tfsdk:"group_interval" json:"group_interval,omitempty"`
			Group_wait     *string            `tfsdk:"group_wait" json:"group_wait,omitempty"`
			Match_re       *map[string]string `tfsdk:"match_re" json:"match_re,omitempty"`
			Matchers       *[]struct {
				IsEqual *bool   `tfsdk:"is_equal" json:"isEqual,omitempty"`
				IsRegex *bool   `tfsdk:"is_regex" json:"isRegex,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Value   *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"matchers" json:"matchers,omitempty"`
			Mute_time_intervals *[]string `tfsdk:"mute_time_intervals" json:"mute_time_intervals,omitempty"`
			Object_matchers     *[]string `tfsdk:"object_matchers" json:"object_matchers,omitempty"`
			Provenance          *string   `tfsdk:"provenance" json:"provenance,omitempty"`
			Receiver            *string   `tfsdk:"receiver" json:"receiver,omitempty"`
			Repeat_interval     *string   `tfsdk:"repeat_interval" json:"repeat_interval,omitempty"`
			RouteSelector       *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"route_selector" json:"routeSelector,omitempty"`
			Routes *map[string]string `tfsdk:"routes" json:"routes,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		Suspend *bool `tfsdk:"suspend" json:"suspend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_grafana_integreatly_org_grafana_notification_policy_v1beta1_manifest"
}

func (r *GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GrafanaNotificationPolicy is the Schema for the GrafanaNotificationPolicy API",
		MarkdownDescription: "GrafanaNotificationPolicy is the Schema for the GrafanaNotificationPolicy API",
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
				Description:         "GrafanaNotificationPolicySpec defines the desired state of GrafanaNotificationPolicy",
				MarkdownDescription: "GrafanaNotificationPolicySpec defines the desired state of GrafanaNotificationPolicy",
				Attributes: map[string]schema.Attribute{
					"allow_cross_namespace_import": schema.BoolAttribute{
						Description:         "Allow the Operator to match this resource with Grafanas outside the current namespace",
						MarkdownDescription: "Allow the Operator to match this resource with Grafanas outside the current namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"editable": schema.BoolAttribute{
						Description:         "Whether to enable or disable editing of the notification policy in Grafana UI",
						MarkdownDescription: "Whether to enable or disable editing of the notification policy in Grafana UI",
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

					"route": schema.SingleNestedAttribute{
						Description:         "Routes for alerts to match against",
						MarkdownDescription: "Routes for alerts to match against",
						Attributes: map[string]schema.Attribute{
							"continue": schema.BoolAttribute{
								Description:         "continue",
								MarkdownDescription: "continue",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_by": schema.ListAttribute{
								Description:         "group by",
								MarkdownDescription: "group by",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_interval": schema.StringAttribute{
								Description:         "group interval",
								MarkdownDescription: "group interval",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_wait": schema.StringAttribute{
								Description:         "group wait",
								MarkdownDescription: "group wait",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"match_re": schema.MapAttribute{
								Description:         "match re",
								MarkdownDescription: "match re",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"matchers": schema.ListNestedAttribute{
								Description:         "matchers",
								MarkdownDescription: "matchers",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"is_equal": schema.BoolAttribute{
											Description:         "is equal",
											MarkdownDescription: "is equal",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_regex": schema.BoolAttribute{
											Description:         "is regex",
											MarkdownDescription: "is regex",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name",
											MarkdownDescription: "name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "value",
											MarkdownDescription: "value",
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

							"mute_time_intervals": schema.ListAttribute{
								Description:         "mute time intervals",
								MarkdownDescription: "mute time intervals",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"object_matchers": schema.ListAttribute{
								Description:         "object matchers",
								MarkdownDescription: "object matchers",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provenance": schema.StringAttribute{
								Description:         "provenance",
								MarkdownDescription: "provenance",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"receiver": schema.StringAttribute{
								Description:         "receiver",
								MarkdownDescription: "receiver",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"repeat_interval": schema.StringAttribute{
								Description:         "repeat interval",
								MarkdownDescription: "repeat interval",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"route_selector": schema.SingleNestedAttribute{
								Description:         "selects GrafanaNotificationPolicyRoutes to merge in when specified mutually exclusive with Routes",
								MarkdownDescription: "selects GrafanaNotificationPolicyRoutes to merge in when specified mutually exclusive with Routes",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"routes": schema.MapAttribute{
								Description:         "routes, mutually exclusive with RouteSelector",
								MarkdownDescription: "routes, mutually exclusive with RouteSelector",
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

func (r *GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_grafana_integreatly_org_grafana_notification_policy_v1beta1_manifest")

	var model GrafanaIntegreatlyOrgGrafanaNotificationPolicyV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("grafana.integreatly.org/v1beta1")
	model.Kind = pointer.String("GrafanaNotificationPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
