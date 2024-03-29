/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package datadoghq_com_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &DatadoghqComDatadogMonitorV1Alpha1Manifest{}
)

func NewDatadoghqComDatadogMonitorV1Alpha1Manifest() datasource.DataSource {
	return &DatadoghqComDatadogMonitorV1Alpha1Manifest{}
}

type DatadoghqComDatadogMonitorV1Alpha1Manifest struct{}

type DatadoghqComDatadogMonitorV1Alpha1ManifestData struct {
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
		ControllerOptions *struct {
			DisableRequiredTags *bool `tfsdk:"disable_required_tags" json:"disableRequiredTags,omitempty"`
		} `tfsdk:"controller_options" json:"controllerOptions,omitempty"`
		Message *string `tfsdk:"message" json:"message,omitempty"`
		Name    *string `tfsdk:"name" json:"name,omitempty"`
		Options *struct {
			EnableLogsSample       *bool   `tfsdk:"enable_logs_sample" json:"enableLogsSample,omitempty"`
			EscalationMessage      *string `tfsdk:"escalation_message" json:"escalationMessage,omitempty"`
			EvaluationDelay        *int64  `tfsdk:"evaluation_delay" json:"evaluationDelay,omitempty"`
			IncludeTags            *bool   `tfsdk:"include_tags" json:"includeTags,omitempty"`
			Locked                 *bool   `tfsdk:"locked" json:"locked,omitempty"`
			NewGroupDelay          *int64  `tfsdk:"new_group_delay" json:"newGroupDelay,omitempty"`
			NoDataTimeframe        *int64  `tfsdk:"no_data_timeframe" json:"noDataTimeframe,omitempty"`
			NotificationPresetName *string `tfsdk:"notification_preset_name" json:"notificationPresetName,omitempty"`
			NotifyAudit            *bool   `tfsdk:"notify_audit" json:"notifyAudit,omitempty"`
			NotifyNoData           *bool   `tfsdk:"notify_no_data" json:"notifyNoData,omitempty"`
			RenotifyInterval       *int64  `tfsdk:"renotify_interval" json:"renotifyInterval,omitempty"`
			RequireFullWindow      *bool   `tfsdk:"require_full_window" json:"requireFullWindow,omitempty"`
			ThresholdWindows       *struct {
				RecoveryWindow *string `tfsdk:"recovery_window" json:"recoveryWindow,omitempty"`
				TriggerWindow  *string `tfsdk:"trigger_window" json:"triggerWindow,omitempty"`
			} `tfsdk:"threshold_windows" json:"thresholdWindows,omitempty"`
			Thresholds *struct {
				Critical         *string `tfsdk:"critical" json:"critical,omitempty"`
				CriticalRecovery *string `tfsdk:"critical_recovery" json:"criticalRecovery,omitempty"`
				Ok               *string `tfsdk:"ok" json:"ok,omitempty"`
				Unknown          *string `tfsdk:"unknown" json:"unknown,omitempty"`
				Warning          *string `tfsdk:"warning" json:"warning,omitempty"`
				WarningRecovery  *string `tfsdk:"warning_recovery" json:"warningRecovery,omitempty"`
			} `tfsdk:"thresholds" json:"thresholds,omitempty"`
			TimeoutH *int64 `tfsdk:"timeout_h" json:"timeoutH,omitempty"`
		} `tfsdk:"options" json:"options,omitempty"`
		Priority        *int64    `tfsdk:"priority" json:"priority,omitempty"`
		Query           *string   `tfsdk:"query" json:"query,omitempty"`
		RestrictedRoles *[]string `tfsdk:"restricted_roles" json:"restrictedRoles,omitempty"`
		Tags            *[]string `tfsdk:"tags" json:"tags,omitempty"`
		Type            *string   `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DatadoghqComDatadogMonitorV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_datadoghq_com_datadog_monitor_v1alpha1_manifest"
}

func (r *DatadoghqComDatadogMonitorV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DatadogMonitor allows to define and manage Monitors from your Kubernetes Cluster",
		MarkdownDescription: "DatadogMonitor allows to define and manage Monitors from your Kubernetes Cluster",
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
				Description:         "DatadogMonitorSpec defines the desired state of DatadogMonitor",
				MarkdownDescription: "DatadogMonitorSpec defines the desired state of DatadogMonitor",
				Attributes: map[string]schema.Attribute{
					"controller_options": schema.SingleNestedAttribute{
						Description:         "ControllerOptions are the optional parameters in the DatadogMonitor controller",
						MarkdownDescription: "ControllerOptions are the optional parameters in the DatadogMonitor controller",
						Attributes: map[string]schema.Attribute{
							"disable_required_tags": schema.BoolAttribute{
								Description:         "DisableRequiredTags disables the automatic addition of required tags to monitors.",
								MarkdownDescription: "DisableRequiredTags disables the automatic addition of required tags to monitors.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"message": schema.StringAttribute{
						Description:         "Message is a message to include with notifications for this monitor",
						MarkdownDescription: "Message is a message to include with notifications for this monitor",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is the monitor name",
						MarkdownDescription: "Name is the monitor name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"options": schema.SingleNestedAttribute{
						Description:         "Options are the optional parameters associated with your monitor",
						MarkdownDescription: "Options are the optional parameters associated with your monitor",
						Attributes: map[string]schema.Attribute{
							"enable_logs_sample": schema.BoolAttribute{
								Description:         "A Boolean indicating whether to send a log sample when the log monitor triggers.",
								MarkdownDescription: "A Boolean indicating whether to send a log sample when the log monitor triggers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"escalation_message": schema.StringAttribute{
								Description:         "A message to include with a re-notification.",
								MarkdownDescription: "A message to include with a re-notification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"evaluation_delay": schema.Int64Attribute{
								Description:         "Time (in seconds) to delay evaluation, as a non-negative integer. For example, if the value is set to 300 (5min), the timeframe is set to last_5m and the time is 7:00, the monitor evaluates data from 6:50 to 6:55. This is useful for AWS CloudWatch and other backfilled metrics to ensure the monitor always has data during evaluation.",
								MarkdownDescription: "Time (in seconds) to delay evaluation, as a non-negative integer. For example, if the value is set to 300 (5min), the timeframe is set to last_5m and the time is 7:00, the monitor evaluates data from 6:50 to 6:55. This is useful for AWS CloudWatch and other backfilled metrics to ensure the monitor always has data during evaluation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include_tags": schema.BoolAttribute{
								Description:         "A Boolean indicating whether notifications from this monitor automatically inserts its triggering tags into the title.",
								MarkdownDescription: "A Boolean indicating whether notifications from this monitor automatically inserts its triggering tags into the title.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"locked": schema.BoolAttribute{
								Description:         "Whether or not the monitor is locked (only editable by creator and admins).",
								MarkdownDescription: "Whether or not the monitor is locked (only editable by creator and admins).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"new_group_delay": schema.Int64Attribute{
								Description:         "Time (in seconds) to allow a host to boot and applications to fully start before starting the evaluation of monitor results. Should be a non negative integer.",
								MarkdownDescription: "Time (in seconds) to allow a host to boot and applications to fully start before starting the evaluation of monitor results. Should be a non negative integer.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"no_data_timeframe": schema.Int64Attribute{
								Description:         "The number of minutes before a monitor notifies after data stops reporting. Datadog recommends at least 2x the monitor timeframe for metric alerts or 2 minutes for service checks. If omitted, 2x the evaluation timeframe is used for metric alerts, and 24 hours is used for service checks.",
								MarkdownDescription: "The number of minutes before a monitor notifies after data stops reporting. Datadog recommends at least 2x the monitor timeframe for metric alerts or 2 minutes for service checks. If omitted, 2x the evaluation timeframe is used for metric alerts, and 24 hours is used for service checks.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"notification_preset_name": schema.StringAttribute{
								Description:         "An enum that toggles the display of additional content sent in the monitor notification.",
								MarkdownDescription: "An enum that toggles the display of additional content sent in the monitor notification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"notify_audit": schema.BoolAttribute{
								Description:         "A Boolean indicating whether tagged users are notified on changes to this monitor.",
								MarkdownDescription: "A Boolean indicating whether tagged users are notified on changes to this monitor.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"notify_no_data": schema.BoolAttribute{
								Description:         "A Boolean indicating whether this monitor notifies when data stops reporting.",
								MarkdownDescription: "A Boolean indicating whether this monitor notifies when data stops reporting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"renotify_interval": schema.Int64Attribute{
								Description:         "The number of minutes after the last notification before a monitor re-notifies on the current status. It only re-notifies if it’s not resolved.",
								MarkdownDescription: "The number of minutes after the last notification before a monitor re-notifies on the current status. It only re-notifies if it’s not resolved.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"require_full_window": schema.BoolAttribute{
								Description:         "A Boolean indicating whether this monitor needs a full window of data before it’s evaluated. We highly recommend you set this to false for sparse metrics, otherwise some evaluations are skipped. Default is false.",
								MarkdownDescription: "A Boolean indicating whether this monitor needs a full window of data before it’s evaluated. We highly recommend you set this to false for sparse metrics, otherwise some evaluations are skipped. Default is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"threshold_windows": schema.SingleNestedAttribute{
								Description:         "A struct of the alerting time window options.",
								MarkdownDescription: "A struct of the alerting time window options.",
								Attributes: map[string]schema.Attribute{
									"recovery_window": schema.StringAttribute{
										Description:         "Describes how long an anomalous metric must be normal before the alert recovers.",
										MarkdownDescription: "Describes how long an anomalous metric must be normal before the alert recovers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"trigger_window": schema.StringAttribute{
										Description:         "Describes how long a metric must be anomalous before an alert triggers.",
										MarkdownDescription: "Describes how long a metric must be anomalous before an alert triggers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"thresholds": schema.SingleNestedAttribute{
								Description:         "A struct of the different monitor threshold values.",
								MarkdownDescription: "A struct of the different monitor threshold values.",
								Attributes: map[string]schema.Attribute{
									"critical": schema.StringAttribute{
										Description:         "The monitor CRITICAL threshold.",
										MarkdownDescription: "The monitor CRITICAL threshold.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"critical_recovery": schema.StringAttribute{
										Description:         "The monitor CRITICAL recovery threshold.",
										MarkdownDescription: "The monitor CRITICAL recovery threshold.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ok": schema.StringAttribute{
										Description:         "The monitor OK threshold.",
										MarkdownDescription: "The monitor OK threshold.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"unknown": schema.StringAttribute{
										Description:         "The monitor UNKNOWN threshold.",
										MarkdownDescription: "The monitor UNKNOWN threshold.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"warning": schema.StringAttribute{
										Description:         "The monitor WARNING threshold.",
										MarkdownDescription: "The monitor WARNING threshold.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"warning_recovery": schema.StringAttribute{
										Description:         "The monitor WARNING recovery threshold.",
										MarkdownDescription: "The monitor WARNING recovery threshold.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_h": schema.Int64Attribute{
								Description:         "The number of hours of the monitor not reporting data before it automatically resolves from a triggered state.",
								MarkdownDescription: "The number of hours of the monitor not reporting data before it automatically resolves from a triggered state.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority is an integer from 1 (high) to 5 (low) indicating alert severity",
						MarkdownDescription: "Priority is an integer from 1 (high) to 5 (low) indicating alert severity",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"query": schema.StringAttribute{
						Description:         "Query is the Datadog monitor query",
						MarkdownDescription: "Query is the Datadog monitor query",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"restricted_roles": schema.ListAttribute{
						Description:         "RestrictedRoles is a list of unique role identifiers to define which roles are allowed to edit the monitor. 'restricted_roles' is the successor of 'locked'. For more information about 'locked' and 'restricted_roles', see the [monitor options docs](https://docs.datadoghq.com/monitors/guide/monitor_api_options/#permissions-options).",
						MarkdownDescription: "RestrictedRoles is a list of unique role identifiers to define which roles are allowed to edit the monitor. 'restricted_roles' is the successor of 'locked'. For more information about 'locked' and 'restricted_roles', see the [monitor options docs](https://docs.datadoghq.com/monitors/guide/monitor_api_options/#permissions-options).",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListAttribute{
						Description:         "Tags is the monitor tags associated with your monitor",
						MarkdownDescription: "Tags is the monitor tags associated with your monitor",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the monitor type",
						MarkdownDescription: "Type is the monitor type",
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
	}
}

func (r *DatadoghqComDatadogMonitorV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_datadoghq_com_datadog_monitor_v1alpha1_manifest")

	var model DatadoghqComDatadogMonitorV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("datadoghq.com/v1alpha1")
	model.Kind = pointer.String("DatadogMonitor")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
