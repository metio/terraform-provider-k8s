/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_openshift_io_v1

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
	_ datasource.DataSource = &MonitoringOpenshiftIoAlertingRuleV1Manifest{}
)

func NewMonitoringOpenshiftIoAlertingRuleV1Manifest() datasource.DataSource {
	return &MonitoringOpenshiftIoAlertingRuleV1Manifest{}
}

type MonitoringOpenshiftIoAlertingRuleV1Manifest struct{}

type MonitoringOpenshiftIoAlertingRuleV1ManifestData struct {
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
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Rules    *[]struct {
				Alert       *string            `tfsdk:"alert" json:"alert,omitempty"`
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Expr        *string            `tfsdk:"expr" json:"expr,omitempty"`
				For         *string            `tfsdk:"for" json:"for,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringOpenshiftIoAlertingRuleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_openshift_io_alerting_rule_v1_manifest"
}

func (r *MonitoringOpenshiftIoAlertingRuleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AlertingRule represents a set of user-defined Prometheus rule groups containing alerting rules.  This resource is the supported method for cluster admins to create alerts based on metrics recorded by the platform monitoring stack in OpenShift, i.e. the Prometheus instance deployed to the openshift-monitoring namespace.  You might use this to create custom alerting rules not shipped with OpenShift based on metrics from components such as the node_exporter, which provides machine-level metrics such as CPU usage, or kube-state-metrics, which provides metrics on Kubernetes usage.  The API is mostly compatible with the upstream PrometheusRule type from the prometheus-operator.  The primary difference being that recording rules are not allowed here -- only alerting rules.  For each AlertingRule resource created, a corresponding PrometheusRule will be created in the openshift-monitoring namespace.  OpenShift requires admins to use the AlertingRule resource rather than the upstream type in order to allow better OpenShift specific defaulting and validation, while not modifying the upstream APIs directly.  You can find upstream API documentation for PrometheusRule resources here:  https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "AlertingRule represents a set of user-defined Prometheus rule groups containing alerting rules.  This resource is the supported method for cluster admins to create alerts based on metrics recorded by the platform monitoring stack in OpenShift, i.e. the Prometheus instance deployed to the openshift-monitoring namespace.  You might use this to create custom alerting rules not shipped with OpenShift based on metrics from components such as the node_exporter, which provides machine-level metrics such as CPU usage, or kube-state-metrics, which provides metrics on Kubernetes usage.  The API is mostly compatible with the upstream PrometheusRule type from the prometheus-operator.  The primary difference being that recording rules are not allowed here -- only alerting rules.  For each AlertingRule resource created, a corresponding PrometheusRule will be created in the openshift-monitoring namespace.  OpenShift requires admins to use the AlertingRule resource rather than the upstream type in order to allow better OpenShift specific defaulting and validation, while not modifying the upstream APIs directly.  You can find upstream API documentation for PrometheusRule resources here:  https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec describes the desired state of this AlertingRule object.",
				MarkdownDescription: "spec describes the desired state of this AlertingRule object.",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "groups is a list of grouped alerting rules.  Rule groups are the unit at which Prometheus parallelizes rule processing.  All rules in a single group share a configured evaluation interval.  All rules in the group will be processed together on this interval, sequentially, and all rules will be processed.  It's common to group related alerting rules into a single AlertingRule resources, and within that resource, closely related alerts, or simply alerts with the same interval, into individual groups.  You are also free to create AlertingRule resources with only a single rule group, but be aware that this can have a performance impact on Prometheus if the group is extremely large or has very complex query expressions to evaluate. Spreading very complex rules across multiple groups to allow them to be processed in parallel is also a common use-case.",
						MarkdownDescription: "groups is a list of grouped alerting rules.  Rule groups are the unit at which Prometheus parallelizes rule processing.  All rules in a single group share a configured evaluation interval.  All rules in the group will be processed together on this interval, sequentially, and all rules will be processed.  It's common to group related alerting rules into a single AlertingRule resources, and within that resource, closely related alerts, or simply alerts with the same interval, into individual groups.  You are also free to create AlertingRule resources with only a single rule group, but be aware that this can have a performance impact on Prometheus if the group is extremely large or has very complex query expressions to evaluate. Spreading very complex rules across multiple groups to allow them to be processed in parallel is also a common use-case.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{
									Description:         "interval is how often rules in the group are evaluated.  If not specified, it defaults to the global.evaluation_interval configured in Prometheus, which itself defaults to 30 seconds.  You can check if this value has been modified from the default on your cluster by inspecting the platform Prometheus configuration: The relevant field in that resource is: spec.evaluationInterval",
									MarkdownDescription: "interval is how often rules in the group are evaluated.  If not specified, it defaults to the global.evaluation_interval configured in Prometheus, which itself defaults to 30 seconds.  You can check if this value has been modified from the default on your cluster by inspecting the platform Prometheus configuration: The relevant field in that resource is: spec.evaluationInterval",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(2048),
										stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "name is the name of the group.",
									MarkdownDescription: "name is the name of the group.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(2048),
									},
								},

								"rules": schema.ListNestedAttribute{
									Description:         "rules is a list of sequentially evaluated alerting rules.  Prometheus may process rule groups in parallel, but rules within a single group are always processed sequentially, and all rules are processed.",
									MarkdownDescription: "rules is a list of sequentially evaluated alerting rules.  Prometheus may process rule groups in parallel, but rules within a single group are always processed sequentially, and all rules are processed.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alert": schema.StringAttribute{
												Description:         "alert is the name of the alert. Must be a valid label value, i.e. may contain any Unicode character.",
												MarkdownDescription: "alert is the name of the alert. Must be a valid label value, i.e. may contain any Unicode character.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(2048),
												},
											},

											"annotations": schema.MapAttribute{
												Description:         "annotations to add to each alert.  These are values that can be used to store longer additional information that you won't query on, such as alert descriptions or runbook links.",
												MarkdownDescription: "annotations to add to each alert.  These are values that can be used to store longer additional information that you won't query on, such as alert descriptions or runbook links.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"expr": schema.StringAttribute{
												Description:         "expr is the PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending or firing alerts.  This is most often a string representing a PromQL expression, e.g.: mapi_current_pending_csr > mapi_max_pending_csr In rare cases this could be a simple integer, e.g. a simple '1' if the intent is to create an alert that is always firing.  This is sometimes used to create an always-firing 'Watchdog' alert in order to ensure the alerting pipeline is functional.",
												MarkdownDescription: "expr is the PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending or firing alerts.  This is most often a string representing a PromQL expression, e.g.: mapi_current_pending_csr > mapi_max_pending_csr In rare cases this could be a simple integer, e.g. a simple '1' if the intent is to create an alert that is always firing.  This is sometimes used to create an always-firing 'Watchdog' alert in order to ensure the alerting pipeline is functional.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"for": schema.StringAttribute{
												Description:         "for is the time period after which alerts are considered firing after first returning results.  Alerts which have not yet fired for long enough are considered pending.",
												MarkdownDescription: "for is the time period after which alerts are considered firing after first returning results.  Alerts which have not yet fired for long enough are considered pending.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(2048),
													stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
												},
											},

											"labels": schema.MapAttribute{
												Description:         "labels to add or overwrite for each alert.  The results of the PromQL expression for the alert will result in an existing set of labels for the alert, after evaluating the expression, for any label specified here with the same name as a label in that set, the label here wins and overwrites the previous value.  These should typically be short identifying values that may be useful to query against.  A common example is the alert severity, where one sets 'severity: warning' under the 'labels' key:",
												MarkdownDescription: "labels to add or overwrite for each alert.  The results of the PromQL expression for the alert will result in an existing set of labels for the alert, after evaluating the expression, for any label specified here with the same name as a label in that set, the label here wins and overwrites the previous value.  These should typically be short identifying values that may be useful to query against.  A common example is the alert severity, where one sets 'severity: warning' under the 'labels' key:",
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

func (r *MonitoringOpenshiftIoAlertingRuleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_openshift_io_alerting_rule_v1_manifest")

	var model MonitoringOpenshiftIoAlertingRuleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("monitoring.openshift.io/v1")
	model.Kind = pointer.String("AlertingRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
