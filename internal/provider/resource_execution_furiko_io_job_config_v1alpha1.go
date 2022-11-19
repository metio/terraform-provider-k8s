/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type ExecutionFurikoIoJobConfigV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ExecutionFurikoIoJobConfigV1Alpha1Resource)(nil)
)

type ExecutionFurikoIoJobConfigV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ExecutionFurikoIoJobConfigV1Alpha1GoModel struct {
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
		Concurrency *struct {
			MaxConcurrency *int64 `tfsdk:"max_concurrency" yaml:"maxConcurrency,omitempty"`

			Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`
		} `tfsdk:"concurrency" yaml:"concurrency,omitempty"`

		Option *struct {
			Options *[]struct {
				Bool *struct {
					Default *bool `tfsdk:"default" yaml:"default,omitempty"`

					FalseVal *string `tfsdk:"false_val" yaml:"falseVal,omitempty"`

					Format *string `tfsdk:"format" yaml:"format,omitempty"`

					TrueVal *string `tfsdk:"true_val" yaml:"trueVal,omitempty"`
				} `tfsdk:"bool" yaml:"bool,omitempty"`

				Date *struct {
					Format *string `tfsdk:"format" yaml:"format,omitempty"`
				} `tfsdk:"date" yaml:"date,omitempty"`

				Label *string `tfsdk:"label" yaml:"label,omitempty"`

				Multi *struct {
					AllowCustom *bool `tfsdk:"allow_custom" yaml:"allowCustom,omitempty"`

					Default *[]string `tfsdk:"default" yaml:"default,omitempty"`

					Delimiter *string `tfsdk:"delimiter" yaml:"delimiter,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"multi" yaml:"multi,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Required *bool `tfsdk:"required" yaml:"required,omitempty"`

				Select *struct {
					AllowCustom *bool `tfsdk:"allow_custom" yaml:"allowCustom,omitempty"`

					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"select" yaml:"select,omitempty"`

				String *struct {
					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					TrimSpaces *bool `tfsdk:"trim_spaces" yaml:"trimSpaces,omitempty"`
				} `tfsdk:"string" yaml:"string,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"options" yaml:"options,omitempty"`
		} `tfsdk:"option" yaml:"option,omitempty"`

		Schedule *struct {
			Constraints *struct {
				NotAfter *string `tfsdk:"not_after" yaml:"notAfter,omitempty"`

				NotBefore *string `tfsdk:"not_before" yaml:"notBefore,omitempty"`
			} `tfsdk:"constraints" yaml:"constraints,omitempty"`

			Cron *struct {
				Expression *string `tfsdk:"expression" yaml:"expression,omitempty"`

				Timezone *string `tfsdk:"timezone" yaml:"timezone,omitempty"`
			} `tfsdk:"cron" yaml:"cron,omitempty"`

			Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

			LastUpdated *string `tfsdk:"last_updated" yaml:"lastUpdated,omitempty"`
		} `tfsdk:"schedule" yaml:"schedule,omitempty"`

		Template *struct {
			Metadata utilities.Dynamic `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Spec *struct {
				ForbidTaskForceDeletion *bool `tfsdk:"forbid_task_force_deletion" yaml:"forbidTaskForceDeletion,omitempty"`

				MaxAttempts *int64 `tfsdk:"max_attempts" yaml:"maxAttempts,omitempty"`

				Parallelism *struct {
					CompletionStrategy *string `tfsdk:"completion_strategy" yaml:"completionStrategy,omitempty"`

					WithCount *int64 `tfsdk:"with_count" yaml:"withCount,omitempty"`

					WithKeys *[]string `tfsdk:"with_keys" yaml:"withKeys,omitempty"`

					WithMatrix *map[string][]string `tfsdk:"with_matrix" yaml:"withMatrix,omitempty"`
				} `tfsdk:"parallelism" yaml:"parallelism,omitempty"`

				RetryDelaySeconds *int64 `tfsdk:"retry_delay_seconds" yaml:"retryDelaySeconds,omitempty"`

				TaskPendingTimeoutSeconds *int64 `tfsdk:"task_pending_timeout_seconds" yaml:"taskPendingTimeoutSeconds,omitempty"`

				TaskTemplate *struct {
					Pod *struct {
						Metadata utilities.Dynamic `tfsdk:"metadata" yaml:"metadata,omitempty"`

						Spec utilities.Dynamic `tfsdk:"spec" yaml:"spec,omitempty"`
					} `tfsdk:"pod" yaml:"pod,omitempty"`
				} `tfsdk:"task_template" yaml:"taskTemplate,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewExecutionFurikoIoJobConfigV1Alpha1Resource() resource.Resource {
	return &ExecutionFurikoIoJobConfigV1Alpha1Resource{}
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_execution_furiko_io_job_config_v1alpha1"
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "JobConfig is the schema for a single job configuration. Multiple Job objects belong to a single JobConfig.",
		MarkdownDescription: "JobConfig is the schema for a single job configuration. Multiple Job objects belong to a single JobConfig.",
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
				Description:         "JobConfigSpec defines the desired state of the JobConfig.",
				MarkdownDescription: "JobConfigSpec defines the desired state of the JobConfig.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"concurrency": {
						Description:         "Concurrency defines the behaviour of multiple concurrent Jobs.",
						MarkdownDescription: "Concurrency defines the behaviour of multiple concurrent Jobs.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_concurrency": {
								Description:         "Maximum number of Jobs that can be running concurrently for the same JobConfig. Cannot be specified if Policy is set to Allow.  Defaults to 1.",
								MarkdownDescription: "Maximum number of Jobs that can be running concurrently for the same JobConfig. Cannot be specified if Policy is set to Allow.  Defaults to 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": {
								Description:         "Policy describes how to treat concurrent executions of the same JobConfig.",
								MarkdownDescription: "Policy describes how to treat concurrent executions of the same JobConfig.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"option": {
						Description:         "Option is an optional field that defines how the JobConfig is parameterized. Each option defined here can subsequently be used in the Template via context variable substitution.",
						MarkdownDescription: "Option is an optional field that defines how the JobConfig is parameterized. Each option defined here can subsequently be used in the Template via context variable substitution.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"options": {
								Description:         "Options is a list of job options.",
								MarkdownDescription: "Options is a list of job options.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"bool": {
										Description:         "Bool adds additional configuration for OptionTypeBool.",
										MarkdownDescription: "Bool adds additional configuration for OptionTypeBool.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default": {
												Description:         "Default value, will be used to populate the option if not specified.",
												MarkdownDescription: "Default value, will be used to populate the option if not specified.",

												Type: types.BoolType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"false_val": {
												Description:         "If Format is custom, will be substituted if value is false. Can also be an empty string.",
												MarkdownDescription: "If Format is custom, will be substituted if value is false. Can also be an empty string.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"format": {
												Description:         "Determines how to format the value as string. Can be one of: TrueFalse, OneZero, YesNo, Custom  Default: TrueFalse",
												MarkdownDescription: "Determines how to format the value as string. Can be one of: TrueFalse, OneZero, YesNo, Custom  Default: TrueFalse",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"true_val": {
												Description:         "If Format is custom, will be substituted if value is true. Can also be an empty string.",
												MarkdownDescription: "If Format is custom, will be substituted if value is true. Can also be an empty string.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"date": {
										Description:         "Date adds additional configuration for OptionTypeDate.",
										MarkdownDescription: "Date adds additional configuration for OptionTypeDate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"format": {
												Description:         "Date format in moment.js format. If not specified, will use RFC3339 format by default.  Date format reference: https://momentjs.com/docs/#/displaying/format/  Default:",
												MarkdownDescription: "Date format in moment.js format. If not specified, will use RFC3339 format by default.  Date format reference: https://momentjs.com/docs/#/displaying/format/  Default:",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"label": {
										Description:         "Label is an optional human-readable label for this option, which is purely used for display purposes.",
										MarkdownDescription: "Label is an optional human-readable label for this option, which is purely used for display purposes.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"multi": {
										Description:         "Multi adds additional configuration for OptionTypeMulti.",
										MarkdownDescription: "Multi adds additional configuration for OptionTypeMulti.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_custom": {
												Description:         "Whether to allow custom values instead of just the list of allowed values.  Default: false",
												MarkdownDescription: "Whether to allow custom values instead of just the list of allowed values.  Default: false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"default": {
												Description:         "Default values, will be used to populate the option if not specified.",
												MarkdownDescription: "Default values, will be used to populate the option if not specified.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delimiter": {
												Description:         "Delimiter to join values by.",
												MarkdownDescription: "Delimiter to join values by.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "List of values to be chosen from.",
												MarkdownDescription: "List of values to be chosen from.",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "The name of the job option. Will be substituted as '${option.NAME}'. Must match the following regex: ^[a-zA-Z_0-9.-]+$",
										MarkdownDescription: "The name of the job option. Will be substituted as '${option.NAME}'. Must match the following regex: ^[a-zA-Z_0-9.-]+$",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"required": {
										Description:         "Required defines whether this field is required.  Default: false",
										MarkdownDescription: "Required defines whether this field is required.  Default: false",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"select": {
										Description:         "Select adds additional configuration for OptionTypeSelect.",
										MarkdownDescription: "Select adds additional configuration for OptionTypeSelect.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_custom": {
												Description:         "Whether to allow custom values instead of just the list of allowed values.  Default: false",
												MarkdownDescription: "Whether to allow custom values instead of just the list of allowed values.  Default: false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"default": {
												Description:         "Default value, will be used to populate the option if not specified.",
												MarkdownDescription: "Default value, will be used to populate the option if not specified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"values": {
												Description:         "List of values to be chosen from.",
												MarkdownDescription: "List of values to be chosen from.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"string": {
										Description:         "String adds additional configuration for OptionTypeString.",
										MarkdownDescription: "String adds additional configuration for OptionTypeString.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default": {
												Description:         "Optional default value, will be used to populate the option if not specified.",
												MarkdownDescription: "Optional default value, will be used to populate the option if not specified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"trim_spaces": {
												Description:         "Whether to trim spaces before substitution.  Default: false",
												MarkdownDescription: "Whether to trim spaces before substitution.  Default: false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "The type of the job option. Can be one of: bool, string, select, multi, date",
										MarkdownDescription: "The type of the job option. Can be one of: bool, string, select, multi, date",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"schedule": {
						Description:         "Schedule is an optional field that defines automatic scheduling of the JobConfig.",
						MarkdownDescription: "Schedule is an optional field that defines automatic scheduling of the JobConfig.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"constraints": {
								Description:         "Specifies any constraints that should apply to this Schedule.",
								MarkdownDescription: "Specifies any constraints that should apply to this Schedule.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"not_after": {
										Description:         "Specifies the latest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules after this time.",
										MarkdownDescription: "Specifies the latest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules after this time.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.DateTime64Validator(),
										},
									},

									"not_before": {
										Description:         "Specifies the earliest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules before this time.",
										MarkdownDescription: "Specifies the earliest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules before this time.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.DateTime64Validator(),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cron": {
								Description:         "Specify a schedule using cron expressions.",
								MarkdownDescription: "Specify a schedule using cron expressions.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"expression": {
										Description:         "Cron expression to specify how the JobConfig will be periodically scheduled. Example: '0 0/5 * * *'.  Supports cron schedules with optional 'seconds' and 'years' fields, i.e. can parse between 5 to 7 tokens.  More information: https://github.com/furiko-io/cronexpr",
										MarkdownDescription: "Cron expression to specify how the JobConfig will be periodically scheduled. Example: '0 0/5 * * *'.  Supports cron schedules with optional 'seconds' and 'years' fields, i.e. can parse between 5 to 7 tokens.  More information: https://github.com/furiko-io/cronexpr",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"timezone": {
										Description:         "Timezone to interpret the cron schedule in. For example, a cron schedule of '0 10 * * *' with a timezone of 'Asia/Singapore' will be interpreted as running at 02:00:00 UTC time every day.  Timezone must be one of the following:  1. A valid tz string (e.g. 'Asia/Singapore', 'America/New_York'). 2. A UTC offset with minutes (e.g. UTC-10:00). 3. A GMT offset with minutes (e.g. GMT+05:30). The meaning is the same as its UTC counterpart.  This field merely is used for parsing the cron Expression, and has nothing to do with /etc/timezone inside the container (i.e. it will not set $TZ automatically).  Defaults to the controller's default configured timezone.",
										MarkdownDescription: "Timezone to interpret the cron schedule in. For example, a cron schedule of '0 10 * * *' with a timezone of 'Asia/Singapore' will be interpreted as running at 02:00:00 UTC time every day.  Timezone must be one of the following:  1. A valid tz string (e.g. 'Asia/Singapore', 'America/New_York'). 2. A UTC offset with minutes (e.g. UTC-10:00). 3. A GMT offset with minutes (e.g. GMT+05:30). The meaning is the same as its UTC counterpart.  This field merely is used for parsing the cron Expression, and has nothing to do with /etc/timezone inside the container (i.e. it will not set $TZ automatically).  Defaults to the controller's default configured timezone.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disabled": {
								Description:         "If true, then automatic scheduling will be disabled for the JobConfig.",
								MarkdownDescription: "If true, then automatic scheduling will be disabled for the JobConfig.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"last_updated": {
								Description:         "Specifies the time that the schedule was last updated. This prevents accidental back-scheduling.  For example, if a JobConfig that was previously disabled from automatic scheduling is now enabled, we do not want to perform back-scheduling for schedules after LastScheduled prior to updating of the JobConfig.",
								MarkdownDescription: "Specifies the time that the schedule was last updated. This prevents accidental back-scheduling.  For example, if a JobConfig that was previously disabled from automatic scheduling is now enabled, we do not want to perform back-scheduling for schedules after LastScheduled prior to updating of the JobConfig.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.DateTime64Validator(),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": {
						Description:         "Template for creating the Job.",
						MarkdownDescription: "Template for creating the Job.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "Standard object's metadata that will be added to Job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata that will be added to Job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": {
								Description:         "Specification of the desired behavior of the job.",
								MarkdownDescription: "Specification of the desired behavior of the job.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"forbid_task_force_deletion": {
										Description:         "Defines whether tasks are allowed to be force deleted or not. If the node is unresponsive, it may be possible that the task cannot be killed by normal graceful deletion. The controller may choose to force delete the task, which would ignore the final state of the task since the node is unable to return whether the task is actually still alive.  If not set to true, there may be some cases when there may actually be two concurrently running tasks when even when ConcurrencyPolicyForbid. Setting this to true would prevent this from happening, but the Job may remain stuck indefinitely until the node recovers.",
										MarkdownDescription: "Defines whether tasks are allowed to be force deleted or not. If the node is unresponsive, it may be possible that the task cannot be killed by normal graceful deletion. The controller may choose to force delete the task, which would ignore the final state of the task since the node is unable to return whether the task is actually still alive.  If not set to true, there may be some cases when there may actually be two concurrently running tasks when even when ConcurrencyPolicyForbid. Setting this to true would prevent this from happening, but the Job may remain stuck indefinitely until the node recovers.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_attempts": {
										Description:         "Specifies maximum number of attempts before the Job will terminate in failure. If defined, the controller will wait retryDelaySeconds before creating the next task. Once maxAttempts is reached, the Job terminates in RetryLimitExceeded.  If parallelism is also defined, this corresponds to the maximum attempts for each parallel task. That is, if there are 5 parallel task to be run at a time, with maxAttempts of 3, the Job may create up to a maximum of 15 tasks if each has failed.  Value must be a positive integer. Defaults to 1.",
										MarkdownDescription: "Specifies maximum number of attempts before the Job will terminate in failure. If defined, the controller will wait retryDelaySeconds before creating the next task. Once maxAttempts is reached, the Job terminates in RetryLimitExceeded.  If parallelism is also defined, this corresponds to the maximum attempts for each parallel task. That is, if there are 5 parallel task to be run at a time, with maxAttempts of 3, the Job may create up to a maximum of 15 tasks if each has failed.  Value must be a positive integer. Defaults to 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"parallelism": {
										Description:         "Describes how to run multiple tasks in parallel for the Job. If not set, then there will be at most a single task running at any time.",
										MarkdownDescription: "Describes how to run multiple tasks in parallel for the Job. If not set, then there will be at most a single task running at any time.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"completion_strategy": {
												Description:         "Defines when the Job will complete when there are multiple tasks running in parallel. For example, if using the AllSuccessful strategy, the Job will only terminate once all parallel tasks have terminated successfully, or once any task has exhausted its maxAttempts limit.",
												MarkdownDescription: "Defines when the Job will complete when there are multiple tasks running in parallel. For example, if using the AllSuccessful strategy, the Job will only terminate once all parallel tasks have terminated successfully, or once any task has exhausted its maxAttempts limit.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"with_count": {
												Description:         "Specifies an exact number of tasks to be run in parallel. The index number can be retrieved via the '${task.index_num}' context variable.",
												MarkdownDescription: "Specifies an exact number of tasks to be run in parallel. The index number can be retrieved via the '${task.index_num}' context variable.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"with_keys": {
												Description:         "Specifies a list of keys corresponding to each task that will be run in parallel. The index key can be retrieved via the '${task.index_key}' context variable.",
												MarkdownDescription: "Specifies a list of keys corresponding to each task that will be run in parallel. The index key can be retrieved via the '${task.index_key}' context variable.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"with_matrix": {
												Description:         "Specifies a matrix of key-value pairs, with each key mapped to a list of possible values, such that tasks will be started for each combination of key-value pairs. The matrix values can be retrieved via context variables in the following format: '${task.index_matrix.<key>}'.",
												MarkdownDescription: "Specifies a matrix of key-value pairs, with each key mapped to a list of possible values, such that tasks will be started for each combination of key-value pairs. The matrix values can be retrieved via context variables in the following format: '${task.index_matrix.<key>}'.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_delay_seconds": {
										Description:         "Optional duration in seconds to wait between retries. If left empty or zero, it means no delay (i.e. retry immediately).  If parallelism is also defined, the retry delay is from the time of the last failed task with the same index. That is, if there are two parallel tasks - index 0 and index 1 - which failed at t=0 and t=15, with retryDelaySeconds of 30, the controller will only create the next attempts at t=30 and t=45 respectively.  Value must be a non-negative integer.",
										MarkdownDescription: "Optional duration in seconds to wait between retries. If left empty or zero, it means no delay (i.e. retry immediately).  If parallelism is also defined, the retry delay is from the time of the last failed task with the same index. That is, if there are two parallel tasks - index 0 and index 1 - which failed at t=0 and t=15, with retryDelaySeconds of 30, the controller will only create the next attempts at t=30 and t=45 respectively.  Value must be a non-negative integer.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"task_pending_timeout_seconds": {
										Description:         "Optional duration in seconds to wait before terminating the task if it is still pending. This field is useful to prevent jobs from being stuck forever if the Job has a deadline to start running by. If not set, it will be set to the DefaultPendingTimeoutSeconds configuration value in the controller. To disable pending timeout, set this to 0.",
										MarkdownDescription: "Optional duration in seconds to wait before terminating the task if it is still pending. This field is useful to prevent jobs from being stuck forever if the Job has a deadline to start running by. If not set, it will be set to the DefaultPendingTimeoutSeconds configuration value in the controller. To disable pending timeout, set this to 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"task_template": {
										Description:         "Defines the template to create a single task in the Job.",
										MarkdownDescription: "Defines the template to create a single task in the Job.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"pod": {
												Description:         "Describes how to create tasks as Pods.",
												MarkdownDescription: "Describes how to create tasks as Pods.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"metadata": {
														Description:         "Standard object's metadata that will be added to Pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
														MarkdownDescription: "Standard object's metadata that will be added to Pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"spec": {
														Description:         "Specification of the desired behavior of the pod. API docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#podspec-v1-core  Supports context variable substitution in the following fields for containers and initContainers: image, command, args, env.value",
														MarkdownDescription: "Specification of the desired behavior of the pod. API docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#podspec-v1-core  Supports context variable substitution in the following fields for containers and initContainers: image, command, args, env.value",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

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

								Required: true,
								Optional: false,
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
		},
	}, nil
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_execution_furiko_io_job_config_v1alpha1")

	var state ExecutionFurikoIoJobConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExecutionFurikoIoJobConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("execution.furiko.io/v1alpha1")
	goModel.Kind = utilities.Ptr("JobConfig")

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

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_execution_furiko_io_job_config_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_execution_furiko_io_job_config_v1alpha1")

	var state ExecutionFurikoIoJobConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExecutionFurikoIoJobConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("execution.furiko.io/v1alpha1")
	goModel.Kind = utilities.Ptr("JobConfig")

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

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_execution_furiko_io_job_config_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
