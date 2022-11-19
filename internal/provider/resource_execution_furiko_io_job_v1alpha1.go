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

type ExecutionFurikoIoJobV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ExecutionFurikoIoJobV1Alpha1Resource)(nil)
)

type ExecutionFurikoIoJobV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ExecutionFurikoIoJobV1Alpha1GoModel struct {
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
		ConfigName *string `tfsdk:"config_name" yaml:"configName,omitempty"`

		KillTimestamp *string `tfsdk:"kill_timestamp" yaml:"killTimestamp,omitempty"`

		OptionValues *string `tfsdk:"option_values" yaml:"optionValues,omitempty"`

		StartPolicy *struct {
			ConcurrencyPolicy *string `tfsdk:"concurrency_policy" yaml:"concurrencyPolicy,omitempty"`

			StartAfter *string `tfsdk:"start_after" yaml:"startAfter,omitempty"`
		} `tfsdk:"start_policy" yaml:"startPolicy,omitempty"`

		Substitutions *map[string]string `tfsdk:"substitutions" yaml:"substitutions,omitempty"`

		Template *struct {
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
		} `tfsdk:"template" yaml:"template,omitempty"`

		TtlSecondsAfterFinished *int64 `tfsdk:"ttl_seconds_after_finished" yaml:"ttlSecondsAfterFinished,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewExecutionFurikoIoJobV1Alpha1Resource() resource.Resource {
	return &ExecutionFurikoIoJobV1Alpha1Resource{}
}

func (r *ExecutionFurikoIoJobV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_execution_furiko_io_job_v1alpha1"
}

func (r *ExecutionFurikoIoJobV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Job is the schema for a single job execution, which may consist of multiple tasks.",
		MarkdownDescription: "Job is the schema for a single job execution, which may consist of multiple tasks.",
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
				Description:         "JobSpec defines the desired state of a Job.",
				MarkdownDescription: "JobSpec defines the desired state of a Job.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"config_name": {
						Description:         "ConfigName allows specifying the name of the JobConfig to create the Job from. The JobConfig must be in the same namespace as the Job.  It is provided as a write-only input field for convenience, and will override the template, labels and annotations from the JobConfig's template.  This field will never be returned from the API. To look up the parent JobConfig, use ownerReferences.",
						MarkdownDescription: "ConfigName allows specifying the name of the JobConfig to create the Job from. The JobConfig must be in the same namespace as the Job.  It is provided as a write-only input field for convenience, and will override the template, labels and annotations from the JobConfig's template.  This field will never be returned from the API. To look up the parent JobConfig, use ownerReferences.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kill_timestamp": {
						Description:         "Specifies the time to start killing the job. When the time passes this timestamp, the controller will start attempting to kill all tasks.",
						MarkdownDescription: "Specifies the time to start killing the job. When the time passes this timestamp, the controller will start attempting to kill all tasks.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							validators.DateTime64Validator(),
						},
					},

					"option_values": {
						Description:         "Specifies key-values pairs of values for Options, in JSON or YAML format.  Example specification:  spec: optionValues: |- myStringOption: 'value' myBoolOption: true mySelectOption: - option1 - option3  Each entry in the optionValues struct should consist of the option's name, and the value could be an arbitrary type that corresponds to the option's type itself. Each option value specified will be evaluated to a string based on the JobConfig's OptionsSpec and added to Substitutions. If the key also exists in Substitutions, that one takes priority.  Cannot be updated after creation.",
						MarkdownDescription: "Specifies key-values pairs of values for Options, in JSON or YAML format.  Example specification:  spec: optionValues: |- myStringOption: 'value' myBoolOption: true mySelectOption: - option1 - option3  Each entry in the optionValues struct should consist of the option's name, and the value could be an arbitrary type that corresponds to the option's type itself. Each option value specified will be evaluated to a string based on the JobConfig's OptionsSpec and added to Substitutions. If the key also exists in Substitutions, that one takes priority.  Cannot be updated after creation.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"start_policy": {
						Description:         "Specifies optional start policy for a Job, which specifies certain conditions which have to be met before a Job is started.",
						MarkdownDescription: "Specifies optional start policy for a Job, which specifies certain conditions which have to be met before a Job is started.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"concurrency_policy": {
								Description:         "Specifies the behaviour when there are other concurrent jobs for the JobConfig.",
								MarkdownDescription: "Specifies the behaviour when there are other concurrent jobs for the JobConfig.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"start_after": {
								Description:         "Specifies the earliest time that the Job can be started after. Can be specified together with other fields.",
								MarkdownDescription: "Specifies the earliest time that the Job can be started after. Can be specified together with other fields.",

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

					"substitutions": {
						Description:         "Defines key-value pairs of context variables to be substituted into the TaskTemplate. Each entry should consist of the full context variable name (i.e. 'ctx.name'), and the values must be a string. Substitutions defined here take highest precedence over both predefined context variables and evaluated OptionValues.  Most users should be using OptionValues to specify custom Job Option values for running the Job instead of using Subsitutions directly.  Cannot be updated after creation.",
						MarkdownDescription: "Defines key-value pairs of context variables to be substituted into the TaskTemplate. Each entry should consist of the full context variable name (i.e. 'ctx.name'), and the values must be a string. Substitutions defined here take highest precedence over both predefined context variables and evaluated OptionValues.  Most users should be using OptionValues to specify custom Job Option values for running the Job instead of using Subsitutions directly.  Cannot be updated after creation.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": {
						Description:         "Template specifies how to create the Job.",
						MarkdownDescription: "Template specifies how to create the Job.",

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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ttl_seconds_after_finished": {
						Description:         "Specifies the maximum lifetime of a Job that is finished. If not set, it will be set to the DefaultTTLSecondsAfterFinished configuration value in the controller.",
						MarkdownDescription: "Specifies the maximum lifetime of a Job that is finished. If not set, it will be set to the DefaultTTLSecondsAfterFinished configuration value in the controller.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": {
						Description:         "Specifies the type of Job. Can be one of: Adhoc, Scheduled  Default: Adhoc",
						MarkdownDescription: "Specifies the type of Job. Can be one of: Adhoc, Scheduled  Default: Adhoc",

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
		},
	}, nil
}

func (r *ExecutionFurikoIoJobV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_execution_furiko_io_job_v1alpha1")

	var state ExecutionFurikoIoJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExecutionFurikoIoJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("execution.furiko.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Job")

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

func (r *ExecutionFurikoIoJobV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_execution_furiko_io_job_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ExecutionFurikoIoJobV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_execution_furiko_io_job_v1alpha1")

	var state ExecutionFurikoIoJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExecutionFurikoIoJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("execution.furiko.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Job")

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

func (r *ExecutionFurikoIoJobV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_execution_furiko_io_job_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
