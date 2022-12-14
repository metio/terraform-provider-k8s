/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type LokiGrafanaComAlertingRuleV1Beta1Resource struct{}

var (
	_ resource.Resource = (*LokiGrafanaComAlertingRuleV1Beta1Resource)(nil)
)

type LokiGrafanaComAlertingRuleV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LokiGrafanaComAlertingRuleV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Groups *[]struct {
			Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

			Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Rules *[]struct {
				Alert *string `tfsdk:"alert" yaml:"alert,omitempty"`

				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Expr *string `tfsdk:"expr" yaml:"expr,omitempty"`

				For *string `tfsdk:"for" yaml:"for,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`
		} `tfsdk:"groups" yaml:"groups,omitempty"`

		TenantID *string `tfsdk:"tenant_id" yaml:"tenantID,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLokiGrafanaComAlertingRuleV1Beta1Resource() resource.Resource {
	return &LokiGrafanaComAlertingRuleV1Beta1Resource{}
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loki_grafana_com_alerting_rule_v1beta1"
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AlertingRule is the Schema for the alertingrules API",
		MarkdownDescription: "AlertingRule is the Schema for the alertingrules API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "AlertingRuleSpec defines the desired state of AlertingRule",
				MarkdownDescription: "AlertingRuleSpec defines the desired state of AlertingRule",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"groups": {
						Description:         "List of groups for alerting rules.",
						MarkdownDescription: "List of groups for alerting rules.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"interval": {
								Description:         "Interval defines the time interval between evaluation of the given alerting rule.",
								MarkdownDescription: "Interval defines the time interval between evaluation of the given alerting rule.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
								},
							},

							"limit": {
								Description:         "Limit defines the number of alerts an alerting rule can produce. 0 is no limit.",
								MarkdownDescription: "Limit defines the number of alerts an alerting rule can produce. 0 is no limit.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the alerting rule group. Must be unique within all alerting rules.",
								MarkdownDescription: "Name of the alerting rule group. Must be unique within all alerting rules.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"rules": {
								Description:         "Rules defines a list of alerting rules",
								MarkdownDescription: "Rules defines a list of alerting rules",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"alert": {
										Description:         "The name of the alert. Must be a valid label value.",
										MarkdownDescription: "The name of the alert. Must be a valid label value.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"annotations": {
										Description:         "Annotations to add to each alert.",
										MarkdownDescription: "Annotations to add to each alert.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expr": {
										Description:         "The LogQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.",
										MarkdownDescription: "The LogQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"for": {
										Description:         "Alerts are considered firing once they have been returned for this long. Alerts which have not yet fired for long enough are considered pending.",
										MarkdownDescription: "Alerts are considered firing once they have been returned for this long. Alerts which have not yet fired for long enough are considered pending.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"labels": {
										Description:         "Labels to add to each alert.",
										MarkdownDescription: "Labels to add to each alert.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tenant_id": {
						Description:         "TenantID of tenant where the alerting rules are evaluated in.",
						MarkdownDescription: "TenantID of tenant where the alerting rules are evaluated in.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_loki_grafana_com_alerting_rule_v1beta1")

	var state LokiGrafanaComAlertingRuleV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LokiGrafanaComAlertingRuleV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("loki.grafana.com/v1beta1")
	goModel.Kind = utilities.Ptr("AlertingRule")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_alerting_rule_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_loki_grafana_com_alerting_rule_v1beta1")

	var state LokiGrafanaComAlertingRuleV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LokiGrafanaComAlertingRuleV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("loki.grafana.com/v1beta1")
	goModel.Kind = utilities.Ptr("AlertingRule")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LokiGrafanaComAlertingRuleV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_loki_grafana_com_alerting_rule_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
