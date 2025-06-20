/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package temporal_io_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &TemporalIoTemporalScheduleV1Beta1Manifest{}
)

func NewTemporalIoTemporalScheduleV1Beta1Manifest() datasource.DataSource {
	return &TemporalIoTemporalScheduleV1Beta1Manifest{}
}

type TemporalIoTemporalScheduleV1Beta1Manifest struct{}

type TemporalIoTemporalScheduleV1Beta1ManifestData struct {
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
		AllowDeletion *bool              `tfsdk:"allow_deletion" json:"allowDeletion,omitempty"`
		Memo          *map[string]string `tfsdk:"memo" json:"memo,omitempty"`
		NamespaceRef  *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"namespace_ref" json:"namespaceRef,omitempty"`
		Schedule *struct {
			Action *struct {
				Workflow *struct {
					ExecutionTimeout *string            `tfsdk:"execution_timeout" json:"executionTimeout,omitempty"`
					Id               *string            `tfsdk:"id" json:"id,omitempty"`
					Inputs           *map[string]string `tfsdk:"inputs" json:"inputs,omitempty"`
					Memo             *map[string]string `tfsdk:"memo" json:"memo,omitempty"`
					RetryPolicy      *struct {
						BackoffCoefficient     *string   `tfsdk:"backoff_coefficient" json:"backoffCoefficient,omitempty"`
						InitialInterval        *string   `tfsdk:"initial_interval" json:"initialInterval,omitempty"`
						MaximumAttempts        *int64    `tfsdk:"maximum_attempts" json:"maximumAttempts,omitempty"`
						MaximumInterval        *string   `tfsdk:"maximum_interval" json:"maximumInterval,omitempty"`
						NonRetryableErrorTypes *[]string `tfsdk:"non_retryable_error_types" json:"nonRetryableErrorTypes,omitempty"`
					} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
					RunTimeout       *string            `tfsdk:"run_timeout" json:"runTimeout,omitempty"`
					SearchAttributes *map[string]string `tfsdk:"search_attributes" json:"searchAttributes,omitempty"`
					TaskQueue        *string            `tfsdk:"task_queue" json:"taskQueue,omitempty"`
					TaskTimeout      *string            `tfsdk:"task_timeout" json:"taskTimeout,omitempty"`
					Type             *string            `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"workflow" json:"workflow,omitempty"`
			} `tfsdk:"action" json:"action,omitempty"`
			Policy *struct {
				CatchupWindow  *string `tfsdk:"catchup_window" json:"catchupWindow,omitempty"`
				Overlap        *string `tfsdk:"overlap" json:"overlap,omitempty"`
				PauseOnFailure *bool   `tfsdk:"pause_on_failure" json:"pauseOnFailure,omitempty"`
			} `tfsdk:"policy" json:"policy,omitempty"`
			Spec *struct {
				Calendars *[]struct {
					Comment    *string `tfsdk:"comment" json:"comment,omitempty"`
					DayOfMonth *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"day_of_month" json:"dayOfMonth,omitempty"`
					DayOfWeek *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"day_of_week" json:"dayOfWeek,omitempty"`
					Hour *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"hour" json:"hour,omitempty"`
					Minute *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"minute" json:"minute,omitempty"`
					Month *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"month" json:"month,omitempty"`
					Second *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"second" json:"second,omitempty"`
					Year *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"year" json:"year,omitempty"`
				} `tfsdk:"calendars" json:"calendars,omitempty"`
				Crons            *[]string `tfsdk:"crons" json:"crons,omitempty"`
				EndAt            *string   `tfsdk:"end_at" json:"endAt,omitempty"`
				ExcludeCalendars *[]struct {
					Comment    *string `tfsdk:"comment" json:"comment,omitempty"`
					DayOfMonth *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"day_of_month" json:"dayOfMonth,omitempty"`
					DayOfWeek *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"day_of_week" json:"dayOfWeek,omitempty"`
					Hour *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"hour" json:"hour,omitempty"`
					Minute *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"minute" json:"minute,omitempty"`
					Month *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"month" json:"month,omitempty"`
					Second *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"second" json:"second,omitempty"`
					Year *[]struct {
						End   *int64 `tfsdk:"end" json:"end,omitempty"`
						Start *int64 `tfsdk:"start" json:"start,omitempty"`
						Step  *int64 `tfsdk:"step" json:"step,omitempty"`
					} `tfsdk:"year" json:"year,omitempty"`
				} `tfsdk:"exclude_calendars" json:"excludeCalendars,omitempty"`
				Intervals *[]struct {
					Every  *string `tfsdk:"every" json:"every,omitempty"`
					Offset *string `tfsdk:"offset" json:"offset,omitempty"`
				} `tfsdk:"intervals" json:"intervals,omitempty"`
				Jitter       *string `tfsdk:"jitter" json:"jitter,omitempty"`
				StartAt      *string `tfsdk:"start_at" json:"startAt,omitempty"`
				TimezoneName *string `tfsdk:"timezone_name" json:"timezoneName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
			State *struct {
				LimitedActions   *bool   `tfsdk:"limited_actions" json:"limitedActions,omitempty"`
				Notes            *string `tfsdk:"notes" json:"notes,omitempty"`
				Paused           *bool   `tfsdk:"paused" json:"paused,omitempty"`
				RemainingActions *int64  `tfsdk:"remaining_actions" json:"remainingActions,omitempty"`
			} `tfsdk:"state" json:"state,omitempty"`
		} `tfsdk:"schedule" json:"schedule,omitempty"`
		SearchAttributes *map[string]string `tfsdk:"search_attributes" json:"searchAttributes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TemporalIoTemporalScheduleV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_temporal_io_temporal_schedule_v1beta1_manifest"
}

func (r *TemporalIoTemporalScheduleV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A TemporalSchedule creates a schedule in the targeted temporal cluster.",
		MarkdownDescription: "A TemporalSchedule creates a schedule in the targeted temporal cluster.",
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
				Description:         "TemporalScheduleSpec defines the desired state of Schedule.",
				MarkdownDescription: "TemporalScheduleSpec defines the desired state of Schedule.",
				Attributes: map[string]schema.Attribute{
					"allow_deletion": schema.BoolAttribute{
						Description:         "AllowDeletion makes the controller delete the Temporal schedule if the CRD is deleted.",
						MarkdownDescription: "AllowDeletion makes the controller delete the Temporal schedule if the CRD is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"memo": schema.MapAttribute{
						Description:         "Memo is optional non-indexed info that will be shown in list workflow.",
						MarkdownDescription: "Memo is optional non-indexed info that will be shown in list workflow.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace_ref": schema.SingleNestedAttribute{
						Description:         "Reference to the temporal namespace the schedule will be created in.",
						MarkdownDescription: "Reference to the temporal namespace the schedule will be created in.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of the temporal object to reference.",
								MarkdownDescription: "The name of the temporal object to reference.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "The namespace of the temporal object to reference. Defaults to the namespace of the requested resource if omitted.",
								MarkdownDescription: "The namespace of the temporal object to reference. Defaults to the namespace of the requested resource if omitted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"schedule": schema.SingleNestedAttribute{
						Description:         "Schedule contains all fields related to a schedule.",
						MarkdownDescription: "Schedule contains all fields related to a schedule.",
						Attributes: map[string]schema.Attribute{
							"action": schema.SingleNestedAttribute{
								Description:         "ScheduleAction contains the actions that the schedule should perform.",
								MarkdownDescription: "ScheduleAction contains the actions that the schedule should perform.",
								Attributes: map[string]schema.Attribute{
									"workflow": schema.SingleNestedAttribute{
										Description:         "ScheduleWorkflowAction describes a workflow to launch.",
										MarkdownDescription: "ScheduleWorkflowAction describes a workflow to launch.",
										Attributes: map[string]schema.Attribute{
											"execution_timeout": schema.StringAttribute{
												Description:         "WorkflowExecutionTimeout is the timeout for duration of workflow execution.",
												MarkdownDescription: "WorkflowExecutionTimeout is the timeout for duration of workflow execution.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"id": schema.StringAttribute{
												Description:         "WorkflowID represents the business identifier of the workflow execution. The WorkflowID of the started workflow may not match this exactly, it may have a timestamp appended for uniqueness. Defaults to a uuid.",
												MarkdownDescription: "WorkflowID represents the business identifier of the workflow execution. The WorkflowID of the started workflow may not match this exactly, it may have a timestamp appended for uniqueness. Defaults to a uuid.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"inputs": schema.MapAttribute{
												Description:         "Inputs contains arguments to pass to the workflow.",
												MarkdownDescription: "Inputs contains arguments to pass to the workflow.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memo": schema.MapAttribute{
												Description:         "Memo is optional non-indexed info that will be shown in list workflow.",
												MarkdownDescription: "Memo is optional non-indexed info that will be shown in list workflow.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_policy": schema.SingleNestedAttribute{
												Description:         "RetryPolicy is the retry policy for the workflow. If a retry policy is specified, in case of workflow failure server will start new workflow execution if needed based on the retry policy.",
												MarkdownDescription: "RetryPolicy is the retry policy for the workflow. If a retry policy is specified, in case of workflow failure server will start new workflow execution if needed based on the retry policy.",
												Attributes: map[string]schema.Attribute{
													"backoff_coefficient": schema.StringAttribute{
														Description:         "Coefficient used to calculate the next retry interval. The next retry interval is previous interval multiplied by the coefficient. Must be 1 or larger.",
														MarkdownDescription: "Coefficient used to calculate the next retry interval. The next retry interval is previous interval multiplied by the coefficient. Must be 1 or larger.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"initial_interval": schema.StringAttribute{
														Description:         "Interval of the first retry. If retryBackoffCoefficient is 1.0 then it is used for all retries.",
														MarkdownDescription: "Interval of the first retry. If retryBackoffCoefficient is 1.0 then it is used for all retries.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum_attempts": schema.Int64Attribute{
														Description:         "Maximum number of attempts. When exceeded the retries stop even if not expired yet. 1 disables retries. 0 means unlimited (up to the timeouts).",
														MarkdownDescription: "Maximum number of attempts. When exceeded the retries stop even if not expired yet. 1 disables retries. 0 means unlimited (up to the timeouts).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum_interval": schema.StringAttribute{
														Description:         "Maximum interval between retries. Exponential backoff leads to interval increase. This value is the cap of the increase. Default is 100x of the initial interval.",
														MarkdownDescription: "Maximum interval between retries. Exponential backoff leads to interval increase. This value is the cap of the increase. Default is 100x of the initial interval.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"non_retryable_error_types": schema.ListAttribute{
														Description:         "Non-Retryable errors types. Will stop retrying if the error type matches this list. Note that this is not a substring match, the error *type* (not message) must match exactly.",
														MarkdownDescription: "Non-Retryable errors types. Will stop retrying if the error type matches this list. Note that this is not a substring match, the error *type* (not message) must match exactly.",
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

											"run_timeout": schema.StringAttribute{
												Description:         "WorkflowRunTimeout is the timeout for duration of a single workflow run.",
												MarkdownDescription: "WorkflowRunTimeout is the timeout for duration of a single workflow run.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"search_attributes": schema.MapAttribute{
												Description:         "SearchAttributes is optional indexed info that can be used in query of List/Scan/Count workflow APIs. The key and value type must be registered on Temporal server side. For supported operations on different server versions see [Visibility]. [Visibility]: https://docs.temporal.io/visibility",
												MarkdownDescription: "SearchAttributes is optional indexed info that can be used in query of List/Scan/Count workflow APIs. The key and value type must be registered on Temporal server side. For supported operations on different server versions see [Visibility]. [Visibility]: https://docs.temporal.io/visibility",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"task_queue": schema.StringAttribute{
												Description:         "TaskQueue represents a workflow task queue. This is also the name of the activity task queue on which activities are scheduled.",
												MarkdownDescription: "TaskQueue represents a workflow task queue. This is also the name of the activity task queue on which activities are scheduled.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"task_timeout": schema.StringAttribute{
												Description:         "WorkflowTaskTimeout is The timeout for processing workflow task from the time the worker pulled this task.",
												MarkdownDescription: "WorkflowTaskTimeout is The timeout for processing workflow task from the time the worker pulled this task.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "WorkflowType represents the identifier used by a workflow author to define the workflow Workflow type name.",
												MarkdownDescription: "WorkflowType represents the identifier used by a workflow author to define the workflow Workflow type name.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": schema.SingleNestedAttribute{
								Description:         "SchedulePolicies represent policies for overlaps, catchups, pause on failure, and workflow ID.",
								MarkdownDescription: "SchedulePolicies represent policies for overlaps, catchups, pause on failure, and workflow ID.",
								Attributes: map[string]schema.Attribute{
									"catchup_window": schema.StringAttribute{
										Description:         "CatchupWindow The Temporal Server might be down or unavailable at the time when a Schedule should take an Action. When the Server comes back up, CatchupWindow controls which missed Actions should be taken at that point.",
										MarkdownDescription: "CatchupWindow The Temporal Server might be down or unavailable at the time when a Schedule should take an Action. When the Server comes back up, CatchupWindow controls which missed Actions should be taken at that point.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overlap": schema.StringAttribute{
										Description:         "Overlap controls what happens when an Action would be started by a Schedule at the same time that an older Action is still running. Supported values: 'skip' - Default. Nothing happens; the Workflow Execution is not started. 'bufferOne' - Starts the Workflow Execution as soon as the current one completes. The buffer is limited to one. If another Workflow Execution is supposed to start, but one is already in the buffer, only the one in the buffer eventually starts. 'bufferAll' - Allows an unlimited number of Workflows to buffer. They are started sequentially. 'cancelOther' - Cancels the running Workflow Execution, and then starts the new one after the old one completes cancellation. 'terminateOther' - Terminates the running Workflow Execution and starts the new one immediately. 'allowAll' - Starts any number of concurrent Workflow Executions. With this policy (and only this policy), more than one Workflow Execution, started by the Schedule, can run simultaneously.",
										MarkdownDescription: "Overlap controls what happens when an Action would be started by a Schedule at the same time that an older Action is still running. Supported values: 'skip' - Default. Nothing happens; the Workflow Execution is not started. 'bufferOne' - Starts the Workflow Execution as soon as the current one completes. The buffer is limited to one. If another Workflow Execution is supposed to start, but one is already in the buffer, only the one in the buffer eventually starts. 'bufferAll' - Allows an unlimited number of Workflows to buffer. They are started sequentially. 'cancelOther' - Cancels the running Workflow Execution, and then starts the new one after the old one completes cancellation. 'terminateOther' - Terminates the running Workflow Execution and starts the new one immediately. 'allowAll' - Starts any number of concurrent Workflow Executions. With this policy (and only this policy), more than one Workflow Execution, started by the Schedule, can run simultaneously.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("skip", "bufferOne", "bufferAll", "cancelOther", "terminateOther", "allowAll"),
										},
									},

									"pause_on_failure": schema.BoolAttribute{
										Description:         "PauseOnFailure if true, and a workflow run fails or times out, turn on 'paused'. This applies after retry policies: the full chain of retries must fail to trigger a pause here.",
										MarkdownDescription: "PauseOnFailure if true, and a workflow run fails or times out, turn on 'paused'. This applies after retry policies: the full chain of retries must fail to trigger a pause here.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "ScheduleSpec is a complete description of a set of absolute timestamps.",
								MarkdownDescription: "ScheduleSpec is a complete description of a set of absolute timestamps.",
								Attributes: map[string]schema.Attribute{
									"calendars": schema.ListNestedAttribute{
										Description:         "Calendars represents calendar-based specifications of times.",
										MarkdownDescription: "Calendars represents calendar-based specifications of times.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"comment": schema.StringAttribute{
													Description:         "Comment describes the intention of this schedule.",
													MarkdownDescription: "Comment describes the intention of this schedule.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"day_of_month": schema.ListNestedAttribute{
													Description:         "DayOfMonth range to match (1-31) Defaults to match all days.",
													MarkdownDescription: "DayOfMonth range to match (1-31) Defaults to match all days.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(31),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 1.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(31),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(31),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"day_of_week": schema.ListNestedAttribute{
													Description:         "DayOfWeek range to match (0-6; 0 is Sunday) Defaults to match all days of the week.",
													MarkdownDescription: "DayOfWeek range to match (0-6; 0 is Sunday) Defaults to match all days of the week.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(6),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(6),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(6),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"hour": schema.ListNestedAttribute{
													Description:         "Hour range to match (0-23). Defaults to 0.",
													MarkdownDescription: "Hour range to match (0-23). Defaults to 0.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(23),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(23),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(23),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"minute": schema.ListNestedAttribute{
													Description:         "Minute range to match (0-59). Defaults to 0.",
													MarkdownDescription: "Minute range to match (0-59). Defaults to 0.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(59),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"month": schema.ListNestedAttribute{
													Description:         "Month range to match (1-12). Defaults to match all months.",
													MarkdownDescription: "Month range to match (1-12). Defaults to match all months.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(12),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 1.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(12),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(12),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"second": schema.ListNestedAttribute{
													Description:         "Second range to match (0-59). Defaults to 0.",
													MarkdownDescription: "Second range to match (0-59). Defaults to 0.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(59),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"year": schema.ListNestedAttribute{
													Description:         "Year range to match. Defaults to match all years.",
													MarkdownDescription: "Year range to match. Defaults to match all years.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1970),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive).",
																MarkdownDescription: "Start of the range (inclusive).",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1970),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																},
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

									"crons": schema.ListAttribute{
										Description:         "Crons are cron based specifications of times. Crons is provided for easy migration from legacy Cron Workflows. For new use cases, we recommend using ScheduleSpec.Calendars or ScheduleSpec. Intervals for readability and maintainability. Once a schedule is created all expressions in Crons will be translated to ScheduleSpec.Calendars on the server. For example, '0 12 * * MON-WED,FRI' is every M/Tu/W/F at noon The string can have 5, 6, or 7 fields, separated by spaces, and they are interpreted in the same way as a ScheduleCalendarSpec: - 5 fields: Minute, Hour, DayOfMonth, Month, DayOfWeek - 6 fields: Minute, Hour, DayOfMonth, Month, DayOfWeek, Year - 7 fields: Second, Minute, Hour, DayOfMonth, Month, DayOfWeek, Year Notes: - If Year is not given, it defaults to *. - If Second is not given, it defaults to 0. - Shorthands @yearly, @monthly, @weekly, @daily, and @hourly are also accepted instead of the 5-7 time fields. - @every <interval>[/<phase>] is accepted and gets compiled into an IntervalSpec instead. <interval> and <phase> should be a decimal integer with a unit suffix s, m, h, or d. - Optionally, the string can be preceded by CRON_TZ=<time zone name> or TZ=<time zone name>, which will get copied to ScheduleSpec.TimeZoneName. (In which case the ScheduleSpec.TimeZone field should be left empty.) - Optionally, '#' followed by a comment can appear at the end of the string. - Note that the special case that some cron implementations have for treating DayOfMonth and DayOfWeek as 'or' instead of 'and' when both are set is not implemented.",
										MarkdownDescription: "Crons are cron based specifications of times. Crons is provided for easy migration from legacy Cron Workflows. For new use cases, we recommend using ScheduleSpec.Calendars or ScheduleSpec. Intervals for readability and maintainability. Once a schedule is created all expressions in Crons will be translated to ScheduleSpec.Calendars on the server. For example, '0 12 * * MON-WED,FRI' is every M/Tu/W/F at noon The string can have 5, 6, or 7 fields, separated by spaces, and they are interpreted in the same way as a ScheduleCalendarSpec: - 5 fields: Minute, Hour, DayOfMonth, Month, DayOfWeek - 6 fields: Minute, Hour, DayOfMonth, Month, DayOfWeek, Year - 7 fields: Second, Minute, Hour, DayOfMonth, Month, DayOfWeek, Year Notes: - If Year is not given, it defaults to *. - If Second is not given, it defaults to 0. - Shorthands @yearly, @monthly, @weekly, @daily, and @hourly are also accepted instead of the 5-7 time fields. - @every <interval>[/<phase>] is accepted and gets compiled into an IntervalSpec instead. <interval> and <phase> should be a decimal integer with a unit suffix s, m, h, or d. - Optionally, the string can be preceded by CRON_TZ=<time zone name> or TZ=<time zone name>, which will get copied to ScheduleSpec.TimeZoneName. (In which case the ScheduleSpec.TimeZone field should be left empty.) - Optionally, '#' followed by a comment can appear at the end of the string. - Note that the special case that some cron implementations have for treating DayOfMonth and DayOfWeek as 'or' instead of 'and' when both are set is not implemented.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"end_at": schema.StringAttribute{
										Description:         "EndAt represents the end of the schedule. Any times after 'endAt' will be skipped. Defaults to the end of time. For example: 2024-05-13T00:00:00Z",
										MarkdownDescription: "EndAt represents the end of the schedule. Any times after 'endAt' will be skipped. Defaults to the end of time. For example: 2024-05-13T00:00:00Z",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
										},
									},

									"exclude_calendars": schema.ListNestedAttribute{
										Description:         "ExcludeCalendars defines any matching times that will be skipped. All fields of the ScheduleCalendarSpec including seconds must match a time for the time to be skipped.",
										MarkdownDescription: "ExcludeCalendars defines any matching times that will be skipped. All fields of the ScheduleCalendarSpec including seconds must match a time for the time to be skipped.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"comment": schema.StringAttribute{
													Description:         "Comment describes the intention of this schedule.",
													MarkdownDescription: "Comment describes the intention of this schedule.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"day_of_month": schema.ListNestedAttribute{
													Description:         "DayOfMonth range to match (1-31) Defaults to match all days.",
													MarkdownDescription: "DayOfMonth range to match (1-31) Defaults to match all days.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(31),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 1.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(31),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(31),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"day_of_week": schema.ListNestedAttribute{
													Description:         "DayOfWeek range to match (0-6; 0 is Sunday) Defaults to match all days of the week.",
													MarkdownDescription: "DayOfWeek range to match (0-6; 0 is Sunday) Defaults to match all days of the week.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(6),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(6),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(6),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"hour": schema.ListNestedAttribute{
													Description:         "Hour range to match (0-23). Defaults to 0.",
													MarkdownDescription: "Hour range to match (0-23). Defaults to 0.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(23),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(23),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(23),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"minute": schema.ListNestedAttribute{
													Description:         "Minute range to match (0-59). Defaults to 0.",
													MarkdownDescription: "Minute range to match (0-59). Defaults to 0.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(59),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"month": schema.ListNestedAttribute{
													Description:         "Month range to match (1-12). Defaults to match all months.",
													MarkdownDescription: "Month range to match (1-12). Defaults to match all months.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(12),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 1.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(12),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(12),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"second": schema.ListNestedAttribute{
													Description:         "Second range to match (0-59). Defaults to 0.",
													MarkdownDescription: "Second range to match (0-59). Defaults to 0.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive). Defaults to 0.",
																MarkdownDescription: "Start of the range (inclusive). Defaults to 0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(59),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																	int64validator.AtMost(59),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"year": schema.ListNestedAttribute{
													Description:         "Year range to match. Defaults to match all years.",
													MarkdownDescription: "Year range to match. Defaults to match all years.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"end": schema.Int64Attribute{
																Description:         "End of the range (inclusive). Defaults to start.",
																MarkdownDescription: "End of the range (inclusive). Defaults to start.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1970),
																},
															},

															"start": schema.Int64Attribute{
																Description:         "Start of the range (inclusive).",
																MarkdownDescription: "Start of the range (inclusive).",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1970),
																},
															},

															"step": schema.Int64Attribute{
																Description:         "Step to be take between each value. Defaults to 1.",
																MarkdownDescription: "Step to be take between each value. Defaults to 1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																},
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

									"intervals": schema.ListNestedAttribute{
										Description:         "Intervals represents interval-based specifications of times.",
										MarkdownDescription: "Intervals represents interval-based specifications of times.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"every": schema.StringAttribute{
													Description:         "Every describes the period to repeat the interval.",
													MarkdownDescription: "Every describes the period to repeat the interval.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"offset": schema.StringAttribute{
													Description:         "Offset is a fixed offset added to the intervals period. Defaults to 0.",
													MarkdownDescription: "Offset is a fixed offset added to the intervals period. Defaults to 0.",
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

									"jitter": schema.StringAttribute{
										Description:         "Jitter represents a duration that is used to apply a jitter to scheduled times. All times will be incremented by a random value from 0 to this amount of jitter, capped by the time until the next schedule. Defaults to 0.",
										MarkdownDescription: "Jitter represents a duration that is used to apply a jitter to scheduled times. All times will be incremented by a random value from 0 to this amount of jitter, capped by the time until the next schedule. Defaults to 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"start_at": schema.StringAttribute{
										Description:         "StartAt represents the start of the schedule. Any times before 'startAt' will be skipped. Together, 'startAt' and 'endAt' make an inclusive interval. Defaults to the beginning of time. For example: 2024-05-13T00:00:00Z",
										MarkdownDescription: "StartAt represents the start of the schedule. Any times before 'startAt' will be skipped. Together, 'startAt' and 'endAt' make an inclusive interval. Defaults to the beginning of time. For example: 2024-05-13T00:00:00Z",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
										},
									},

									"timezone_name": schema.StringAttribute{
										Description:         "TimeZoneName represents the IANA time zone name, for example 'US/Pacific'. The definition will be loaded by Temporal Server from the environment it runs in. Calendar spec matching is based on literal matching of the clock time with no special handling of DST: if you write a calendar spec that fires at 2:30am and specify a time zone that follows DST, that action will not be triggered on the day that has no 2:30am. Similarly, an action that fires at 1:30am will be triggered twice on the day that has two 1:30s. Note: No actions are taken on leap-seconds (e.g. 23:59:60 UTC). Defaults to UTC.",
										MarkdownDescription: "TimeZoneName represents the IANA time zone name, for example 'US/Pacific'. The definition will be loaded by Temporal Server from the environment it runs in. Calendar spec matching is based on literal matching of the clock time with no special handling of DST: if you write a calendar spec that fires at 2:30am and specify a time zone that follows DST, that action will not be triggered on the day that has no 2:30am. Similarly, an action that fires at 1:30am will be triggered twice on the day that has two 1:30s. Note: No actions are taken on leap-seconds (e.g. 23:59:60 UTC). Defaults to UTC.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"state": schema.SingleNestedAttribute{
								Description:         "ScheduleState describes the current state of a schedule.",
								MarkdownDescription: "ScheduleState describes the current state of a schedule.",
								Attributes: map[string]schema.Attribute{
									"limited_actions": schema.BoolAttribute{
										Description:         "LimitedActions limits actions. While true RemainingActions will be decremented for each action taken. Skipped actions (due to overlap policy) do not count against remaining actions.",
										MarkdownDescription: "LimitedActions limits actions. While true RemainingActions will be decremented for each action taken. Skipped actions (due to overlap policy) do not count against remaining actions.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"notes": schema.StringAttribute{
										Description:         "Note is an informative human-readable message with contextual notes, e.g. the reason a Schedule is paused. The system may overwrite this message on certain conditions, e.g. when pause-on-failure happens.",
										MarkdownDescription: "Note is an informative human-readable message with contextual notes, e.g. the reason a Schedule is paused. The system may overwrite this message on certain conditions, e.g. when pause-on-failure happens.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"paused": schema.BoolAttribute{
										Description:         "Paused is true if the schedule is paused.",
										MarkdownDescription: "Paused is true if the schedule is paused.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remaining_actions": schema.Int64Attribute{
										Description:         "RemainingActions represents the Actions remaining in this Schedule. Once this number hits 0, no further Actions are taken. manual actions through backfill or ScheduleHandle.Trigger still run.",
										MarkdownDescription: "RemainingActions represents the Actions remaining in this Schedule. Once this number hits 0, no further Actions are taken. manual actions through backfill or ScheduleHandle.Trigger still run.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"search_attributes": schema.MapAttribute{
						Description:         "SearchAttributes is optional indexed info that can be used in query of List/Scan/Count workflow APIs. The key and value type must be registered on Temporal server side. For supported operations on different server versions see [Visibility]. [Visibility]: https://docs.temporal.io/visibility",
						MarkdownDescription: "SearchAttributes is optional indexed info that can be used in query of List/Scan/Count workflow APIs. The key and value type must be registered on Temporal server side. For supported operations on different server versions see [Visibility]. [Visibility]: https://docs.temporal.io/visibility",
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
		},
	}
}

func (r *TemporalIoTemporalScheduleV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_temporal_io_temporal_schedule_v1beta1_manifest")

	var model TemporalIoTemporalScheduleV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("temporal.io/v1beta1")
	model.Kind = pointer.String("TemporalSchedule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
