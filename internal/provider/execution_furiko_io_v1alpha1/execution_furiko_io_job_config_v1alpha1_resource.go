/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package execution_furiko_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &ExecutionFurikoIoJobConfigV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &ExecutionFurikoIoJobConfigV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &ExecutionFurikoIoJobConfigV1Alpha1Resource{}
)

func NewExecutionFurikoIoJobConfigV1Alpha1Resource() resource.Resource {
	return &ExecutionFurikoIoJobConfigV1Alpha1Resource{}
}

type ExecutionFurikoIoJobConfigV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ExecutionFurikoIoJobConfigV1Alpha1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Concurrency *struct {
			MaxConcurrency *int64  `tfsdk:"max_concurrency" json:"maxConcurrency,omitempty"`
			Policy         *string `tfsdk:"policy" json:"policy,omitempty"`
		} `tfsdk:"concurrency" json:"concurrency,omitempty"`
		Option *struct {
			Options *[]struct {
				Bool *struct {
					Default  *bool   `tfsdk:"default" json:"default,omitempty"`
					FalseVal *string `tfsdk:"false_val" json:"falseVal,omitempty"`
					Format   *string `tfsdk:"format" json:"format,omitempty"`
					TrueVal  *string `tfsdk:"true_val" json:"trueVal,omitempty"`
				} `tfsdk:"bool" json:"bool,omitempty"`
				Date *struct {
					Format *string `tfsdk:"format" json:"format,omitempty"`
				} `tfsdk:"date" json:"date,omitempty"`
				Label *string `tfsdk:"label" json:"label,omitempty"`
				Multi *struct {
					AllowCustom *bool     `tfsdk:"allow_custom" json:"allowCustom,omitempty"`
					Default     *[]string `tfsdk:"default" json:"default,omitempty"`
					Delimiter   *string   `tfsdk:"delimiter" json:"delimiter,omitempty"`
					Values      *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"multi" json:"multi,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Required *bool   `tfsdk:"required" json:"required,omitempty"`
				Select   *struct {
					AllowCustom *bool     `tfsdk:"allow_custom" json:"allowCustom,omitempty"`
					Default     *string   `tfsdk:"default" json:"default,omitempty"`
					Values      *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"select" json:"select,omitempty"`
				String *struct {
					Default    *string `tfsdk:"default" json:"default,omitempty"`
					TrimSpaces *bool   `tfsdk:"trim_spaces" json:"trimSpaces,omitempty"`
				} `tfsdk:"string" json:"string,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
		} `tfsdk:"option" json:"option,omitempty"`
		Schedule *struct {
			Constraints *struct {
				NotAfter  *string `tfsdk:"not_after" json:"notAfter,omitempty"`
				NotBefore *string `tfsdk:"not_before" json:"notBefore,omitempty"`
			} `tfsdk:"constraints" json:"constraints,omitempty"`
			Cron *struct {
				Expression  *string   `tfsdk:"expression" json:"expression,omitempty"`
				Expressions *[]string `tfsdk:"expressions" json:"expressions,omitempty"`
				Timezone    *string   `tfsdk:"timezone" json:"timezone,omitempty"`
			} `tfsdk:"cron" json:"cron,omitempty"`
			Disabled    *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
			LastUpdated *string `tfsdk:"last_updated" json:"lastUpdated,omitempty"`
		} `tfsdk:"schedule" json:"schedule,omitempty"`
		Template *struct {
			Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec     *struct {
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
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_execution_furiko_io_job_config_v1alpha1"
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "JobConfig is the schema for a single job configuration. Multiple Job objects belong to a single JobConfig.",
		MarkdownDescription: "JobConfig is the schema for a single job configuration. Multiple Job objects belong to a single JobConfig.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "JobConfigSpec defines the desired state of the JobConfig.",
				MarkdownDescription: "JobConfigSpec defines the desired state of the JobConfig.",
				Attributes: map[string]schema.Attribute{
					"concurrency": schema.SingleNestedAttribute{
						Description:         "Concurrency defines the behaviour of multiple concurrent Jobs.",
						MarkdownDescription: "Concurrency defines the behaviour of multiple concurrent Jobs.",
						Attributes: map[string]schema.Attribute{
							"max_concurrency": schema.Int64Attribute{
								Description:         "Maximum number of Jobs that can be running concurrently for the same JobConfig. Cannot be specified if Policy is set to Allow.  Defaults to 1.",
								MarkdownDescription: "Maximum number of Jobs that can be running concurrently for the same JobConfig. Cannot be specified if Policy is set to Allow.  Defaults to 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"policy": schema.StringAttribute{
								Description:         "Policy describes how to treat concurrent executions of the same JobConfig.",
								MarkdownDescription: "Policy describes how to treat concurrent executions of the same JobConfig.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"option": schema.SingleNestedAttribute{
						Description:         "Option is an optional field that defines how the JobConfig is parameterized. Each option defined here can subsequently be used in the Template via context variable substitution.",
						MarkdownDescription: "Option is an optional field that defines how the JobConfig is parameterized. Each option defined here can subsequently be used in the Template via context variable substitution.",
						Attributes: map[string]schema.Attribute{
							"options": schema.ListNestedAttribute{
								Description:         "Options is a list of job options.",
								MarkdownDescription: "Options is a list of job options.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"bool": schema.SingleNestedAttribute{
											Description:         "Bool adds additional configuration for OptionTypeBool.",
											MarkdownDescription: "Bool adds additional configuration for OptionTypeBool.",
											Attributes: map[string]schema.Attribute{
												"default": schema.BoolAttribute{
													Description:         "Default value, will be used to populate the option if not specified.",
													MarkdownDescription: "Default value, will be used to populate the option if not specified.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"false_val": schema.StringAttribute{
													Description:         "If Format is custom, will be substituted if value is false. Can also be an empty string.",
													MarkdownDescription: "If Format is custom, will be substituted if value is false. Can also be an empty string.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"format": schema.StringAttribute{
													Description:         "Determines how to format the value as string. Can be one of: TrueFalse, OneZero, YesNo, Custom  Default: TrueFalse",
													MarkdownDescription: "Determines how to format the value as string. Can be one of: TrueFalse, OneZero, YesNo, Custom  Default: TrueFalse",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"true_val": schema.StringAttribute{
													Description:         "If Format is custom, will be substituted if value is true. Can also be an empty string.",
													MarkdownDescription: "If Format is custom, will be substituted if value is true. Can also be an empty string.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"date": schema.SingleNestedAttribute{
											Description:         "Date adds additional configuration for OptionTypeDate.",
											MarkdownDescription: "Date adds additional configuration for OptionTypeDate.",
											Attributes: map[string]schema.Attribute{
												"format": schema.StringAttribute{
													Description:         "Date format in moment.js format. If not specified, will use RFC3339 format by default.  Date format reference: https://momentjs.com/docs/#/displaying/format/  Default:",
													MarkdownDescription: "Date format in moment.js format. If not specified, will use RFC3339 format by default.  Date format reference: https://momentjs.com/docs/#/displaying/format/  Default:",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"label": schema.StringAttribute{
											Description:         "Label is an optional human-readable label for this option, which is purely used for display purposes.",
											MarkdownDescription: "Label is an optional human-readable label for this option, which is purely used for display purposes.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"multi": schema.SingleNestedAttribute{
											Description:         "Multi adds additional configuration for OptionTypeMulti.",
											MarkdownDescription: "Multi adds additional configuration for OptionTypeMulti.",
											Attributes: map[string]schema.Attribute{
												"allow_custom": schema.BoolAttribute{
													Description:         "Whether to allow custom values instead of just the list of allowed values.  Default: false",
													MarkdownDescription: "Whether to allow custom values instead of just the list of allowed values.  Default: false",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"default": schema.ListAttribute{
													Description:         "Default values, will be used to populate the option if not specified.",
													MarkdownDescription: "Default values, will be used to populate the option if not specified.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"delimiter": schema.StringAttribute{
													Description:         "Delimiter to join values by.",
													MarkdownDescription: "Delimiter to join values by.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "List of values to be chosen from.",
													MarkdownDescription: "List of values to be chosen from.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "The name of the job option. Will be substituted as '${option.NAME}'. Must match the following regex: ^[a-zA-Z_0-9.-]+$",
											MarkdownDescription: "The name of the job option. Will be substituted as '${option.NAME}'. Must match the following regex: ^[a-zA-Z_0-9.-]+$",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"required": schema.BoolAttribute{
											Description:         "Required defines whether this field is required.  Default: false",
											MarkdownDescription: "Required defines whether this field is required.  Default: false",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"select": schema.SingleNestedAttribute{
											Description:         "Select adds additional configuration for OptionTypeSelect.",
											MarkdownDescription: "Select adds additional configuration for OptionTypeSelect.",
											Attributes: map[string]schema.Attribute{
												"allow_custom": schema.BoolAttribute{
													Description:         "Whether to allow custom values instead of just the list of allowed values.  Default: false",
													MarkdownDescription: "Whether to allow custom values instead of just the list of allowed values.  Default: false",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"default": schema.StringAttribute{
													Description:         "Default value, will be used to populate the option if not specified.",
													MarkdownDescription: "Default value, will be used to populate the option if not specified.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "List of values to be chosen from.",
													MarkdownDescription: "List of values to be chosen from.",
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

										"string": schema.SingleNestedAttribute{
											Description:         "String adds additional configuration for OptionTypeString.",
											MarkdownDescription: "String adds additional configuration for OptionTypeString.",
											Attributes: map[string]schema.Attribute{
												"default": schema.StringAttribute{
													Description:         "Optional default value, will be used to populate the option if not specified.",
													MarkdownDescription: "Optional default value, will be used to populate the option if not specified.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"trim_spaces": schema.BoolAttribute{
													Description:         "Whether to trim spaces before substitution.  Default: false",
													MarkdownDescription: "Whether to trim spaces before substitution.  Default: false",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"type": schema.StringAttribute{
											Description:         "The type of the job option. Can be one of: bool, string, select, multi, date",
											MarkdownDescription: "The type of the job option. Can be one of: bool, string, select, multi, date",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"schedule": schema.SingleNestedAttribute{
						Description:         "Schedule is an optional field that defines automatic scheduling of the JobConfig.",
						MarkdownDescription: "Schedule is an optional field that defines automatic scheduling of the JobConfig.",
						Attributes: map[string]schema.Attribute{
							"constraints": schema.SingleNestedAttribute{
								Description:         "Specifies any constraints that should apply to this Schedule.",
								MarkdownDescription: "Specifies any constraints that should apply to this Schedule.",
								Attributes: map[string]schema.Attribute{
									"not_after": schema.StringAttribute{
										Description:         "Specifies the latest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules after this time.",
										MarkdownDescription: "Specifies the latest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules after this time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
										},
									},

									"not_before": schema.StringAttribute{
										Description:         "Specifies the earliest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules before this time.",
										MarkdownDescription: "Specifies the earliest possible time that is allowed to be scheduled. If set, the scheduler should not create schedules before this time.",
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

							"cron": schema.SingleNestedAttribute{
								Description:         "Specify a schedule using cron expressions.",
								MarkdownDescription: "Specify a schedule using cron expressions.",
								Attributes: map[string]schema.Attribute{
									"expression": schema.StringAttribute{
										Description:         "Cron expression to specify how the JobConfig will be periodically scheduled. Example: '0 0/5 * * *'.  Supports cron schedules with optional 'seconds' and 'years' fields, i.e. can parse between 5 to 7 tokens.  More information: https://github.com/furiko-io/cronexpr",
										MarkdownDescription: "Cron expression to specify how the JobConfig will be periodically scheduled. Example: '0 0/5 * * *'.  Supports cron schedules with optional 'seconds' and 'years' fields, i.e. can parse between 5 to 7 tokens.  More information: https://github.com/furiko-io/cronexpr",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"expressions": schema.ListAttribute{
										Description:         "List of cron expressions to specify how the JobConfig will be periodically scheduled.  Take for example a requirement to schedule a job every day at 3AM, 3:30AM and 4AM. There is no way to represent this with a single cron expression, but we could do so with two cron expressions: '0/30 3 * * *', and '0 4 * * *'.  Exactly one of Expression or Expressions must be specified. If two expressions fall on the same time, only one of them will take effect.",
										MarkdownDescription: "List of cron expressions to specify how the JobConfig will be periodically scheduled.  Take for example a requirement to schedule a job every day at 3AM, 3:30AM and 4AM. There is no way to represent this with a single cron expression, but we could do so with two cron expressions: '0/30 3 * * *', and '0 4 * * *'.  Exactly one of Expression or Expressions must be specified. If two expressions fall on the same time, only one of them will take effect.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timezone": schema.StringAttribute{
										Description:         "Timezone to interpret the cron schedule in. For example, a cron schedule of '0 10 * * *' with a timezone of 'Asia/Singapore' will be interpreted as running at 02:00:00 UTC time every day.  Timezone must be one of the following:  1. A valid tz string (e.g. 'Asia/Singapore', 'America/New_York'). 2. A UTC offset with minutes (e.g. UTC-10:00). 3. A GMT offset with minutes (e.g. GMT+05:30). The meaning is the same as its UTC counterpart.  This field merely is used for parsing the cron Expression, and has nothing to do with /etc/timezone inside the container (i.e. it will not set $TZ automatically).  Defaults to the controller's default configured timezone.",
										MarkdownDescription: "Timezone to interpret the cron schedule in. For example, a cron schedule of '0 10 * * *' with a timezone of 'Asia/Singapore' will be interpreted as running at 02:00:00 UTC time every day.  Timezone must be one of the following:  1. A valid tz string (e.g. 'Asia/Singapore', 'America/New_York'). 2. A UTC offset with minutes (e.g. UTC-10:00). 3. A GMT offset with minutes (e.g. GMT+05:30). The meaning is the same as its UTC counterpart.  This field merely is used for parsing the cron Expression, and has nothing to do with /etc/timezone inside the container (i.e. it will not set $TZ automatically).  Defaults to the controller's default configured timezone.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"disabled": schema.BoolAttribute{
								Description:         "If true, then automatic scheduling will be disabled for the JobConfig.",
								MarkdownDescription: "If true, then automatic scheduling will be disabled for the JobConfig.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"last_updated": schema.StringAttribute{
								Description:         "Specifies the time that the schedule was last updated. This prevents accidental back-scheduling.  For example, if a JobConfig that was previously disabled from automatic scheduling is now enabled, we do not want to perform back-scheduling for schedules after LastScheduled prior to updating of the JobConfig.",
								MarkdownDescription: "Specifies the time that the schedule was last updated. This prevents accidental back-scheduling.  For example, if a JobConfig that was previously disabled from automatic scheduling is now enabled, we do not want to perform back-scheduling for schedules after LastScheduled prior to updating of the JobConfig.",
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

					"template": schema.SingleNestedAttribute{
						Description:         "Template for creating the Job.",
						MarkdownDescription: "Template for creating the Job.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.MapAttribute{
								Description:         "Standard object's metadata that will be added to Job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata that will be added to Job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Specification of the desired behavior of the job.",
								MarkdownDescription: "Specification of the desired behavior of the job.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_execution_furiko_io_job_config_v1alpha1")

	var model ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("execution.furiko.io/v1alpha1")
	model.Kind = pointer.String("JobConfig")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "execution.furiko.io", Version: "v1alpha1", Resource: "jobconfigs"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_execution_furiko_io_job_config_v1alpha1")

	var data ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "execution.furiko.io", Version: "v1alpha1", Resource: "jobconfigs"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_execution_furiko_io_job_config_v1alpha1")

	var model ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("execution.furiko.io/v1alpha1")
	model.Kind = pointer.String("JobConfig")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "execution.furiko.io", Version: "v1alpha1", Resource: "jobconfigs"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_execution_furiko_io_job_config_v1alpha1")

	var data ExecutionFurikoIoJobConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "execution.furiko.io", Version: "v1alpha1", Resource: "jobconfigs"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "execution.furiko.io", Version: "v1alpha1", Resource: "jobconfigs"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *ExecutionFurikoIoJobConfigV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
