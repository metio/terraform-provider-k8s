/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package execution_furiko_io_v1alpha1

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
	_ datasource.DataSource = &ExecutionFurikoIoJobV1Alpha1Manifest{}
)

func NewExecutionFurikoIoJobV1Alpha1Manifest() datasource.DataSource {
	return &ExecutionFurikoIoJobV1Alpha1Manifest{}
}

type ExecutionFurikoIoJobV1Alpha1Manifest struct{}

type ExecutionFurikoIoJobV1Alpha1ManifestData struct {
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
		ConfigName    *string `tfsdk:"config_name" json:"configName,omitempty"`
		KillTimestamp *string `tfsdk:"kill_timestamp" json:"killTimestamp,omitempty"`
		OptionValues  *string `tfsdk:"option_values" json:"optionValues,omitempty"`
		StartPolicy   *struct {
			ConcurrencyPolicy *string `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
			StartAfter        *string `tfsdk:"start_after" json:"startAfter,omitempty"`
		} `tfsdk:"start_policy" json:"startPolicy,omitempty"`
		Substitutions *map[string]string `tfsdk:"substitutions" json:"substitutions,omitempty"`
		Template      *struct {
			ForbidTaskForceDeletion *bool  `tfsdk:"forbid_task_force_deletion" json:"forbidTaskForceDeletion,omitempty"`
			MaxAttempts             *int64 `tfsdk:"max_attempts" json:"maxAttempts,omitempty"`
			Parallelism             *struct {
				CompletionStrategy *string              `tfsdk:"completion_strategy" json:"completionStrategy,omitempty"`
				WithCount          *int64               `tfsdk:"with_count" json:"withCount,omitempty"`
				WithKeys           *[]string            `tfsdk:"with_keys" json:"withKeys,omitempty"`
				WithMatrix         *map[string][]string `tfsdk:"with_matrix" json:"withMatrix,omitempty"`
			} `tfsdk:"parallelism" json:"parallelism,omitempty"`
			RetryDelaySeconds         *int64 `tfsdk:"retry_delay_seconds" json:"retryDelaySeconds,omitempty"`
			TaskPendingTimeoutSeconds *int64 `tfsdk:"task_pending_timeout_seconds" json:"taskPendingTimeoutSeconds,omitempty"`
			TaskTemplate              *struct {
				Pod *struct {
					Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec     *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"pod" json:"pod,omitempty"`
			} `tfsdk:"task_template" json:"taskTemplate,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		TtlSecondsAfterFinished *int64  `tfsdk:"ttl_seconds_after_finished" json:"ttlSecondsAfterFinished,omitempty"`
		Type                    *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExecutionFurikoIoJobV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_execution_furiko_io_job_v1alpha1_manifest"
}

func (r *ExecutionFurikoIoJobV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Job is the schema for a single job execution, which may consist of multiple tasks.",
		MarkdownDescription: "Job is the schema for a single job execution, which may consist of multiple tasks.",
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
				Description:         "JobSpec defines the desired state of a Job.",
				MarkdownDescription: "JobSpec defines the desired state of a Job.",
				Attributes: map[string]schema.Attribute{
					"config_name": schema.StringAttribute{
						Description:         "ConfigName allows specifying the name of the JobConfig to create the Job from. The JobConfig must be in the same namespace as the Job.  It is provided as a write-only input field for convenience, and will override the template, labels and annotations from the JobConfig's template.  This field will never be returned from the API. To look up the parent JobConfig, use ownerReferences.",
						MarkdownDescription: "ConfigName allows specifying the name of the JobConfig to create the Job from. The JobConfig must be in the same namespace as the Job.  It is provided as a write-only input field for convenience, and will override the template, labels and annotations from the JobConfig's template.  This field will never be returned from the API. To look up the parent JobConfig, use ownerReferences.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kill_timestamp": schema.StringAttribute{
						Description:         "Specifies the time to start killing the job. When the time passes this timestamp, the controller will start attempting to kill all tasks.",
						MarkdownDescription: "Specifies the time to start killing the job. When the time passes this timestamp, the controller will start attempting to kill all tasks.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.DateTime64Validator(),
						},
					},

					"option_values": schema.StringAttribute{
						Description:         "Specifies key-values pairs of values for Options, in JSON or YAML format.  Example specification:  spec: optionValues: |- myStringOption: 'value' myBoolOption: true mySelectOption: - option1 - option3  Each entry in the optionValues struct should consist of the option's name, and the value could be an arbitrary type that corresponds to the option's type itself. Each option value specified will be evaluated to a string based on the JobConfig's OptionsSpec and added to Substitutions. If the key also exists in Substitutions, that one takes priority.  Cannot be updated after creation.",
						MarkdownDescription: "Specifies key-values pairs of values for Options, in JSON or YAML format.  Example specification:  spec: optionValues: |- myStringOption: 'value' myBoolOption: true mySelectOption: - option1 - option3  Each entry in the optionValues struct should consist of the option's name, and the value could be an arbitrary type that corresponds to the option's type itself. Each option value specified will be evaluated to a string based on the JobConfig's OptionsSpec and added to Substitutions. If the key also exists in Substitutions, that one takes priority.  Cannot be updated after creation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"start_policy": schema.SingleNestedAttribute{
						Description:         "Specifies optional start policy for a Job, which specifies certain conditions which have to be met before a Job is started.",
						MarkdownDescription: "Specifies optional start policy for a Job, which specifies certain conditions which have to be met before a Job is started.",
						Attributes: map[string]schema.Attribute{
							"concurrency_policy": schema.StringAttribute{
								Description:         "Specifies the behaviour when there are other concurrent jobs for the JobConfig.",
								MarkdownDescription: "Specifies the behaviour when there are other concurrent jobs for the JobConfig.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"start_after": schema.StringAttribute{
								Description:         "Specifies the earliest time that the Job can be started after. Can be specified together with other fields.",
								MarkdownDescription: "Specifies the earliest time that the Job can be started after. Can be specified together with other fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.DateTime64Validator(),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"substitutions": schema.MapAttribute{
						Description:         "Defines key-value pairs of context variables to be substituted into the TaskTemplate. Each entry should consist of the full context variable name (i.e. 'ctx.name'), and the values must be a string. Substitutions defined here take highest precedence over both predefined context variables and evaluated OptionValues.  Most users should be using OptionValues to specify custom Job Option values for running the Job instead of using Subsitutions directly.  Cannot be updated after creation.",
						MarkdownDescription: "Defines key-value pairs of context variables to be substituted into the TaskTemplate. Each entry should consist of the full context variable name (i.e. 'ctx.name'), and the values must be a string. Substitutions defined here take highest precedence over both predefined context variables and evaluated OptionValues.  Most users should be using OptionValues to specify custom Job Option values for running the Job instead of using Subsitutions directly.  Cannot be updated after creation.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template specifies how to create the Job.",
						MarkdownDescription: "Template specifies how to create the Job.",
						Attributes: map[string]schema.Attribute{
							"forbid_task_force_deletion": schema.BoolAttribute{
								Description:         "Defines whether tasks are allowed to be force deleted or not. If the node is unresponsive, it may be possible that the task cannot be killed by normal graceful deletion. The controller may choose to force delete the task, which would ignore the final state of the task since the node is unable to return whether the task is actually still alive.  If not set to true, there may be some cases when there may actually be two concurrently running tasks when even when ConcurrencyPolicyForbid. Setting this to true would prevent this from happening, but the Job may remain stuck indefinitely until the node recovers.",
								MarkdownDescription: "Defines whether tasks are allowed to be force deleted or not. If the node is unresponsive, it may be possible that the task cannot be killed by normal graceful deletion. The controller may choose to force delete the task, which would ignore the final state of the task since the node is unable to return whether the task is actually still alive.  If not set to true, there may be some cases when there may actually be two concurrently running tasks when even when ConcurrencyPolicyForbid. Setting this to true would prevent this from happening, but the Job may remain stuck indefinitely until the node recovers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_attempts": schema.Int64Attribute{
								Description:         "Specifies maximum number of attempts before the Job will terminate in failure. If defined, the controller will wait retryDelaySeconds before creating the next task. Once maxAttempts is reached, the Job terminates in RetryLimitExceeded.  If parallelism is also defined, this corresponds to the maximum attempts for each parallel task. That is, if there are 5 parallel task to be run at a time, with maxAttempts of 3, the Job may create up to a maximum of 15 tasks if each has failed.  Value must be a positive integer. Defaults to 1.",
								MarkdownDescription: "Specifies maximum number of attempts before the Job will terminate in failure. If defined, the controller will wait retryDelaySeconds before creating the next task. Once maxAttempts is reached, the Job terminates in RetryLimitExceeded.  If parallelism is also defined, this corresponds to the maximum attempts for each parallel task. That is, if there are 5 parallel task to be run at a time, with maxAttempts of 3, the Job may create up to a maximum of 15 tasks if each has failed.  Value must be a positive integer. Defaults to 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parallelism": schema.SingleNestedAttribute{
								Description:         "Describes how to run multiple tasks in parallel for the Job. If not set, then there will be at most a single task running at any time.",
								MarkdownDescription: "Describes how to run multiple tasks in parallel for the Job. If not set, then there will be at most a single task running at any time.",
								Attributes: map[string]schema.Attribute{
									"completion_strategy": schema.StringAttribute{
										Description:         "Defines when the Job will complete when there are multiple tasks running in parallel. For example, if using the AllSuccessful strategy, the Job will only terminate once all parallel tasks have terminated successfully, or once any task has exhausted its maxAttempts limit.",
										MarkdownDescription: "Defines when the Job will complete when there are multiple tasks running in parallel. For example, if using the AllSuccessful strategy, the Job will only terminate once all parallel tasks have terminated successfully, or once any task has exhausted its maxAttempts limit.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"with_count": schema.Int64Attribute{
										Description:         "Specifies an exact number of tasks to be run in parallel. The index number can be retrieved via the '${task.index_num}' context variable.",
										MarkdownDescription: "Specifies an exact number of tasks to be run in parallel. The index number can be retrieved via the '${task.index_num}' context variable.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"with_keys": schema.ListAttribute{
										Description:         "Specifies a list of keys corresponding to each task that will be run in parallel. The index key can be retrieved via the '${task.index_key}' context variable.",
										MarkdownDescription: "Specifies a list of keys corresponding to each task that will be run in parallel. The index key can be retrieved via the '${task.index_key}' context variable.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"with_matrix": schema.MapAttribute{
										Description:         "Specifies a matrix of key-value pairs, with each key mapped to a list of possible values, such that tasks will be started for each combination of key-value pairs. The matrix values can be retrieved via context variables in the following format: '${task.index_matrix.<key>}'.",
										MarkdownDescription: "Specifies a matrix of key-value pairs, with each key mapped to a list of possible values, such that tasks will be started for each combination of key-value pairs. The matrix values can be retrieved via context variables in the following format: '${task.index_matrix.<key>}'.",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"retry_delay_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds to wait between retries. If left empty or zero, it means no delay (i.e. retry immediately).  If parallelism is also defined, the retry delay is from the time of the last failed task with the same index. That is, if there are two parallel tasks - index 0 and index 1 - which failed at t=0 and t=15, with retryDelaySeconds of 30, the controller will only create the next attempts at t=30 and t=45 respectively.  Value must be a non-negative integer.",
								MarkdownDescription: "Optional duration in seconds to wait between retries. If left empty or zero, it means no delay (i.e. retry immediately).  If parallelism is also defined, the retry delay is from the time of the last failed task with the same index. That is, if there are two parallel tasks - index 0 and index 1 - which failed at t=0 and t=15, with retryDelaySeconds of 30, the controller will only create the next attempts at t=30 and t=45 respectively.  Value must be a non-negative integer.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"task_pending_timeout_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds to wait before terminating the task if it is still pending. This field is useful to prevent jobs from being stuck forever if the Job has a deadline to start running by. If not set, it will be set to the DefaultPendingTimeoutSeconds configuration value in the controller. To disable pending timeout, set this to 0.",
								MarkdownDescription: "Optional duration in seconds to wait before terminating the task if it is still pending. This field is useful to prevent jobs from being stuck forever if the Job has a deadline to start running by. If not set, it will be set to the DefaultPendingTimeoutSeconds configuration value in the controller. To disable pending timeout, set this to 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"task_template": schema.SingleNestedAttribute{
								Description:         "Defines the template to create a single task in the Job.",
								MarkdownDescription: "Defines the template to create a single task in the Job.",
								Attributes: map[string]schema.Attribute{
									"pod": schema.SingleNestedAttribute{
										Description:         "Describes how to create tasks as Pods.",
										MarkdownDescription: "Describes how to create tasks as Pods.",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.MapAttribute{
												Description:         "Standard object's metadata that will be added to Pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
												MarkdownDescription: "Standard object's metadata that will be added to Pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"spec": schema.MapAttribute{
												Description:         "Specification of the desired behavior of the pod. API docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#podspec-v1-core  Supports context variable substitution in the following fields for containers and initContainers: image, command, args, env.value",
												MarkdownDescription: "Specification of the desired behavior of the pod. API docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#podspec-v1-core  Supports context variable substitution in the following fields for containers and initContainers: image, command, args, env.value",
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
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "Specifies the maximum lifetime of a Job that is finished. If not set, it will be set to the DefaultTTLSecondsAfterFinished configuration value in the controller.",
						MarkdownDescription: "Specifies the maximum lifetime of a Job that is finished. If not set, it will be set to the DefaultTTLSecondsAfterFinished configuration value in the controller.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Specifies the type of Job. Can be one of: Adhoc, Scheduled  Default: Adhoc",
						MarkdownDescription: "Specifies the type of Job. Can be one of: Adhoc, Scheduled  Default: Adhoc",
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

func (r *ExecutionFurikoIoJobV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_execution_furiko_io_job_v1alpha1_manifest")

	var model ExecutionFurikoIoJobV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("execution.furiko.io/v1alpha1")
	model.Kind = pointer.String("Job")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
