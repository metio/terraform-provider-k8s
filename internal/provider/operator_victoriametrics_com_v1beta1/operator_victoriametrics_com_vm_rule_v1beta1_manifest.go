/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OperatorVictoriametricsComVmruleV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmruleV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmruleV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmruleV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmruleV1Beta1ManifestData struct {
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
			Concurrency         *int64               `tfsdk:"concurrency" json:"concurrency,omitempty"`
			Extra_filter_labels *map[string]string   `tfsdk:"extra_filter_labels" json:"extra_filter_labels,omitempty"`
			Headers             *[]string            `tfsdk:"headers" json:"headers,omitempty"`
			Interval            *string              `tfsdk:"interval" json:"interval,omitempty"`
			Labels              *map[string]string   `tfsdk:"labels" json:"labels,omitempty"`
			Limit               *int64               `tfsdk:"limit" json:"limit,omitempty"`
			Name                *string              `tfsdk:"name" json:"name,omitempty"`
			Notifier_headers    *[]string            `tfsdk:"notifier_headers" json:"notifier_headers,omitempty"`
			Params              *map[string][]string `tfsdk:"params" json:"params,omitempty"`
			Rules               *[]struct {
				Alert                *string            `tfsdk:"alert" json:"alert,omitempty"`
				Annotations          *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Debug                *bool              `tfsdk:"debug" json:"debug,omitempty"`
				Expr                 *string            `tfsdk:"expr" json:"expr,omitempty"`
				For                  *string            `tfsdk:"for" json:"for,omitempty"`
				Keep_firing_for      *string            `tfsdk:"keep_firing_for" json:"keep_firing_for,omitempty"`
				Labels               *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Record               *string            `tfsdk:"record" json:"record,omitempty"`
				Update_entries_limit *int64             `tfsdk:"update_entries_limit" json:"update_entries_limit,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
			Tenant *string `tfsdk:"tenant" json:"tenant,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmruleV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_rule_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmruleV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMRule defines rule records for vmalert application",
		MarkdownDescription: "VMRule defines rule records for vmalert application",
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
				Description:         "VMRuleSpec defines the desired state of VMRule",
				MarkdownDescription: "VMRuleSpec defines the desired state of VMRule",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "Groups list of group rules",
						MarkdownDescription: "Groups list of group rules",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"concurrency": schema.Int64Attribute{
									Description:         "Concurrency defines how many rules execute at once.",
									MarkdownDescription: "Concurrency defines how many rules execute at once.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"extra_filter_labels": schema.MapAttribute{
									Description:         "ExtraFilterLabels optional list of label filters applied to every rule'srequest withing a group. Is compatible only with VM datasource.See more details at https://docs.victoriametrics.com#prometheus-querying-api-enhancementsDeprecated, use params instead",
									MarkdownDescription: "ExtraFilterLabels optional list of label filters applied to every rule'srequest withing a group. Is compatible only with VM datasource.See more details at https://docs.victoriametrics.com#prometheus-querying-api-enhancementsDeprecated, use params instead",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"headers": schema.ListAttribute{
									Description:         "Headers contains optional HTTP headers added to each rule requestMust be in form 'header-name: value'For example: headers:   - 'CustomHeader: foo'   - 'CustomHeader2: bar'",
									MarkdownDescription: "Headers contains optional HTTP headers added to each rule requestMust be in form 'header-name: value'For example: headers:   - 'CustomHeader: foo'   - 'CustomHeader2: bar'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval": schema.StringAttribute{
									Description:         "evaluation interval for group",
									MarkdownDescription: "evaluation interval for group",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels optional list of labels added to every rule within a group.It has priority over the external labels.Labels are commonly used for adding environmentor tenant-specific tag.",
									MarkdownDescription: "Labels optional list of labels added to every rule within a group.It has priority over the external labels.Labels are commonly used for adding environmentor tenant-specific tag.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"limit": schema.Int64Attribute{
									Description:         "Limit the number of alerts an alerting rule and series a recordingrule can produce",
									MarkdownDescription: "Limit the number of alerts an alerting rule and series a recordingrule can produce",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of group",
									MarkdownDescription: "Name of group",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"notifier_headers": schema.ListAttribute{
									Description:         "NotifierHeaders contains optional HTTP headers added to each alert request which will send to notifierMust be in form 'header-name: value'For example: headers:   - 'CustomHeader: foo'   - 'CustomHeader2: bar'",
									MarkdownDescription: "NotifierHeaders contains optional HTTP headers added to each alert request which will send to notifierMust be in form 'header-name: value'For example: headers:   - 'CustomHeader: foo'   - 'CustomHeader2: bar'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"params": schema.MapAttribute{
									Description:         "Params optional HTTP URL parameters added to each rule request",
									MarkdownDescription: "Params optional HTTP URL parameters added to each rule request",
									ElementType:         types.ListType{ElemType: types.StringType},
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rules": schema.ListNestedAttribute{
									Description:         "Rules list of alert rules",
									MarkdownDescription: "Rules list of alert rules",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alert": schema.StringAttribute{
												Description:         "Alert is a name for alert",
												MarkdownDescription: "Alert is a name for alert",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"annotations": schema.MapAttribute{
												Description:         "Annotations will be added to rule configuration",
												MarkdownDescription: "Annotations will be added to rule configuration",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"debug": schema.BoolAttribute{
												Description:         "Debug enables logging for ruleit useful for tracking",
												MarkdownDescription: "Debug enables logging for ruleit useful for tracking",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"expr": schema.StringAttribute{
												Description:         "Expr is query, that will be evaluated at dataSource",
												MarkdownDescription: "Expr is query, that will be evaluated at dataSource",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"for": schema.StringAttribute{
												Description:         "For evaluation interval in time.Duration format30s, 1m, 1h  or nanoseconds",
												MarkdownDescription: "For evaluation interval in time.Duration format30s, 1m, 1h  or nanoseconds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"keep_firing_for": schema.StringAttribute{
												Description:         "KeepFiringFor will make alert continue firing for this longeven when the alerting expression no longer has results.Use time.Duration format, 30s, 1m, 1h  or nanoseconds",
												MarkdownDescription: "KeepFiringFor will make alert continue firing for this longeven when the alerting expression no longer has results.Use time.Duration format, 30s, 1m, 1h  or nanoseconds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels will be added to rule configuration",
												MarkdownDescription: "Labels will be added to rule configuration",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"record": schema.StringAttribute{
												Description:         "Record represents a query, that will be recorded to dataSource",
												MarkdownDescription: "Record represents a query, that will be recorded to dataSource",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"update_entries_limit": schema.Int64Attribute{
												Description:         "UpdateEntriesLimit defines max number of rule's state updates stored in memory.Overrides '-rule.updateEntriesLimit' in vmalert.",
												MarkdownDescription: "UpdateEntriesLimit defines max number of rule's state updates stored in memory.Overrides '-rule.updateEntriesLimit' in vmalert.",
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

								"tenant": schema.StringAttribute{
									Description:         "Tenant id for group, can be used only with enterprise version of vmalertSee more details at https://docs.victoriametrics.com/vmalert.html#multitenancy",
									MarkdownDescription: "Tenant id for group, can be used only with enterprise version of vmalertSee more details at https://docs.victoriametrics.com/vmalert.html#multitenancy",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type defines datasource type for enterprise version of vmalertpossible values - prometheus,graphite",
									MarkdownDescription: "Type defines datasource type for enterprise version of vmalertpossible values - prometheus,graphite",
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

func (r *OperatorVictoriametricsComVmruleV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_rule_v1beta1_manifest")

	var model OperatorVictoriametricsComVmruleV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
