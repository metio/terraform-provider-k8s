/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1beta1

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
	_ datasource.DataSource = &LokiGrafanaComRecordingRuleV1Beta1Manifest{}
)

func NewLokiGrafanaComRecordingRuleV1Beta1Manifest() datasource.DataSource {
	return &LokiGrafanaComRecordingRuleV1Beta1Manifest{}
}

type LokiGrafanaComRecordingRuleV1Beta1Manifest struct{}

type LokiGrafanaComRecordingRuleV1Beta1ManifestData struct {
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
				Expr   *string `tfsdk:"expr" json:"expr,omitempty"`
				Record *string `tfsdk:"record" json:"record,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
		TenantID *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LokiGrafanaComRecordingRuleV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_loki_grafana_com_recording_rule_v1beta1_manifest"
}

func (r *LokiGrafanaComRecordingRuleV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RecordingRule is the Schema for the recordingrules API",
		MarkdownDescription: "RecordingRule is the Schema for the recordingrules API",
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
				Description:         "RecordingRuleSpec defines the desired state of RecordingRule",
				MarkdownDescription: "RecordingRuleSpec defines the desired state of RecordingRule",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "List of groups for recording rules.",
						MarkdownDescription: "List of groups for recording rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{
									Description:         "Interval defines the time interval between evaluation of the given recoding rule.",
									MarkdownDescription: "Interval defines the time interval between evaluation of the given recoding rule.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
									},
								},

								"limit": schema.Int64Attribute{
									Description:         "Limit defines the number of series a recording rule can produce. 0 is no limit.",
									MarkdownDescription: "Limit defines the number of series a recording rule can produce. 0 is no limit.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the recording rule group. Must be unique within all recording rules.",
									MarkdownDescription: "Name of the recording rule group. Must be unique within all recording rules.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"rules": schema.ListNestedAttribute{
									Description:         "Rules defines a list of recording rules",
									MarkdownDescription: "Rules defines a list of recording rules",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"expr": schema.StringAttribute{
												Description:         "The LogQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.",
												MarkdownDescription: "The LogQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"record": schema.StringAttribute{
												Description:         "The name of the time series to output to. Must be a valid metric name.",
												MarkdownDescription: "The name of the time series to output to. Must be a valid metric name.",
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
						Description:         "TenantID of tenant where the recording rules are evaluated in.",
						MarkdownDescription: "TenantID of tenant where the recording rules are evaluated in.",
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

func (r *LokiGrafanaComRecordingRuleV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_recording_rule_v1beta1_manifest")

	var model LokiGrafanaComRecordingRuleV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("loki.grafana.com/v1beta1")
	model.Kind = pointer.String("RecordingRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
