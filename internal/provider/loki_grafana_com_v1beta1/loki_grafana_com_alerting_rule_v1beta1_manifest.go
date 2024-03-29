/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &LokiGrafanaComAlertingRuleV1Beta1Manifest{}
)

func NewLokiGrafanaComAlertingRuleV1Beta1Manifest() datasource.DataSource {
	return &LokiGrafanaComAlertingRuleV1Beta1Manifest{}
}

type LokiGrafanaComAlertingRuleV1Beta1Manifest struct{}

type LokiGrafanaComAlertingRuleV1Beta1ManifestData struct {
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
		Groups *[]struct {
			Interval *string `tfsdk:"interval" json:"interval,omitempty"`
			Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Rules    *[]struct {
				Alert       *string            `tfsdk:"alert" json:"alert,omitempty"`
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Expr        *string            `tfsdk:"expr" json:"expr,omitempty"`
				For         *string            `tfsdk:"for" json:"for,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
		TenantID *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_loki_grafana_com_alerting_rule_v1beta1_manifest"
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AlertingRule is the Schema for the alertingrules API",
		MarkdownDescription: "AlertingRule is the Schema for the alertingrules API",
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
				Description:         "AlertingRuleSpec defines the desired state of AlertingRule",
				MarkdownDescription: "AlertingRuleSpec defines the desired state of AlertingRule",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "List of groups for alerting rules.",
						MarkdownDescription: "List of groups for alerting rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{
									Description:         "Interval defines the time interval between evaluation of the givenalerting rule.",
									MarkdownDescription: "Interval defines the time interval between evaluation of the givenalerting rule.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
									},
								},

								"limit": schema.Int64Attribute{
									Description:         "Limit defines the number of alerts an alerting rule can produce. 0 is no limit.",
									MarkdownDescription: "Limit defines the number of alerts an alerting rule can produce. 0 is no limit.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the alerting rule group. Must be unique within all alerting rules.",
									MarkdownDescription: "Name of the alerting rule group. Must be unique within all alerting rules.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"rules": schema.ListNestedAttribute{
									Description:         "Rules defines a list of alerting rules",
									MarkdownDescription: "Rules defines a list of alerting rules",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alert": schema.StringAttribute{
												Description:         "The name of the alert. Must be a valid label value.",
												MarkdownDescription: "The name of the alert. Must be a valid label value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"annotations": schema.MapAttribute{
												Description:         "Annotations to add to each alert.",
												MarkdownDescription: "Annotations to add to each alert.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"expr": schema.StringAttribute{
												Description:         "The LogQL expression to evaluate. Every evaluation cycle this isevaluated at the current time, and all resultant time series becomepending/firing alerts.",
												MarkdownDescription: "The LogQL expression to evaluate. Every evaluation cycle this isevaluated at the current time, and all resultant time series becomepending/firing alerts.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"for": schema.StringAttribute{
												Description:         "Alerts are considered firing once they have been returned for this long.Alerts which have not yet fired for long enough are considered pending.",
												MarkdownDescription: "Alerts are considered firing once they have been returned for this long.Alerts which have not yet fired for long enough are considered pending.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
												},
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to add to each alert.",
												MarkdownDescription: "Labels to add to each alert.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tenant_id": schema.StringAttribute{
						Description:         "TenantID of tenant where the alerting rules are evaluated in.",
						MarkdownDescription: "TenantID of tenant where the alerting rules are evaluated in.",
						Required:            true,
						Optional:            false,
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

func (r *LokiGrafanaComAlertingRuleV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_alerting_rule_v1beta1_manifest")

	var model LokiGrafanaComAlertingRuleV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("loki.grafana.com/v1beta1")
	model.Kind = pointer.String("AlertingRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
