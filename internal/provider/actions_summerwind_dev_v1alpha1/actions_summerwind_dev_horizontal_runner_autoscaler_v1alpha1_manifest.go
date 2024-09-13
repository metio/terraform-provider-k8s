/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package actions_summerwind_dev_v1alpha1

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
	_ datasource.DataSource = &ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest{}
)

func NewActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest() datasource.DataSource {
	return &ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest{}
}

type ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest struct{}

type ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1ManifestData struct {
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
		CapacityReservations *[]struct {
			EffectiveTime  *string `tfsdk:"effective_time" json:"effectiveTime,omitempty"`
			ExpirationTime *string `tfsdk:"expiration_time" json:"expirationTime,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Replicas       *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"capacity_reservations" json:"capacityReservations,omitempty"`
		GithubAPICredentialsFrom *struct {
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"github_api_credentials_from" json:"githubAPICredentialsFrom,omitempty"`
		MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
		Metrics     *[]struct {
			RepositoryNames     *[]string `tfsdk:"repository_names" json:"repositoryNames,omitempty"`
			ScaleDownAdjustment *int64    `tfsdk:"scale_down_adjustment" json:"scaleDownAdjustment,omitempty"`
			ScaleDownFactor     *string   `tfsdk:"scale_down_factor" json:"scaleDownFactor,omitempty"`
			ScaleDownThreshold  *string   `tfsdk:"scale_down_threshold" json:"scaleDownThreshold,omitempty"`
			ScaleUpAdjustment   *int64    `tfsdk:"scale_up_adjustment" json:"scaleUpAdjustment,omitempty"`
			ScaleUpFactor       *string   `tfsdk:"scale_up_factor" json:"scaleUpFactor,omitempty"`
			ScaleUpThreshold    *string   `tfsdk:"scale_up_threshold" json:"scaleUpThreshold,omitempty"`
			Type                *string   `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		MinReplicas                        *int64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		ScaleDownDelaySecondsAfterScaleOut *int64 `tfsdk:"scale_down_delay_seconds_after_scale_out" json:"scaleDownDelaySecondsAfterScaleOut,omitempty"`
		ScaleTargetRef                     *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"scale_target_ref" json:"scaleTargetRef,omitempty"`
		ScaleUpTriggers *[]struct {
			Amount      *int64  `tfsdk:"amount" json:"amount,omitempty"`
			Duration    *string `tfsdk:"duration" json:"duration,omitempty"`
			GithubEvent *struct {
				CheckRun *struct {
					Names        *[]string `tfsdk:"names" json:"names,omitempty"`
					Repositories *[]string `tfsdk:"repositories" json:"repositories,omitempty"`
					Status       *string   `tfsdk:"status" json:"status,omitempty"`
					Types        *[]string `tfsdk:"types" json:"types,omitempty"`
				} `tfsdk:"check_run" json:"checkRun,omitempty"`
				PullRequest *struct {
					Branches *[]string `tfsdk:"branches" json:"branches,omitempty"`
					Types    *[]string `tfsdk:"types" json:"types,omitempty"`
				} `tfsdk:"pull_request" json:"pullRequest,omitempty"`
				Push        *map[string]string `tfsdk:"push" json:"push,omitempty"`
				WorkflowJob *map[string]string `tfsdk:"workflow_job" json:"workflowJob,omitempty"`
			} `tfsdk:"github_event" json:"githubEvent,omitempty"`
		} `tfsdk:"scale_up_triggers" json:"scaleUpTriggers,omitempty"`
		ScheduledOverrides *[]struct {
			EndTime        *string `tfsdk:"end_time" json:"endTime,omitempty"`
			MinReplicas    *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
			RecurrenceRule *struct {
				Frequency *string `tfsdk:"frequency" json:"frequency,omitempty"`
				UntilTime *string `tfsdk:"until_time" json:"untilTime,omitempty"`
			} `tfsdk:"recurrence_rule" json:"recurrenceRule,omitempty"`
			StartTime *string `tfsdk:"start_time" json:"startTime,omitempty"`
		} `tfsdk:"scheduled_overrides" json:"scheduledOverrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest"
}

func (r *ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HorizontalRunnerAutoscaler is the Schema for the horizontalrunnerautoscaler API",
		MarkdownDescription: "HorizontalRunnerAutoscaler is the Schema for the horizontalrunnerautoscaler API",
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
				Description:         "HorizontalRunnerAutoscalerSpec defines the desired state of HorizontalRunnerAutoscaler",
				MarkdownDescription: "HorizontalRunnerAutoscalerSpec defines the desired state of HorizontalRunnerAutoscaler",
				Attributes: map[string]schema.Attribute{
					"capacity_reservations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effective_time": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"expiration_time": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
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

					"github_api_credentials_from": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_replicas": schema.Int64Attribute{
						Description:         "MaxReplicas is the maximum number of replicas the deployment is allowed to scale",
						MarkdownDescription: "MaxReplicas is the maximum number of replicas the deployment is allowed to scale",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics": schema.ListNestedAttribute{
						Description:         "Metrics is the collection of various metric targets to calculate desired number of runners",
						MarkdownDescription: "Metrics is the collection of various metric targets to calculate desired number of runners",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"repository_names": schema.ListAttribute{
									Description:         "RepositoryNames is the list of repository names to be used for calculating the metric. For example, a repository name is the REPO part of 'github.com/USER/REPO'.",
									MarkdownDescription: "RepositoryNames is the list of repository names to be used for calculating the metric. For example, a repository name is the REPO part of 'github.com/USER/REPO'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scale_down_adjustment": schema.Int64Attribute{
									Description:         "ScaleDownAdjustment is the number of runners removed on scale-down. You can only specify either ScaleDownFactor or ScaleDownAdjustment.",
									MarkdownDescription: "ScaleDownAdjustment is the number of runners removed on scale-down. You can only specify either ScaleDownFactor or ScaleDownAdjustment.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scale_down_factor": schema.StringAttribute{
									Description:         "ScaleDownFactor is the multiplicative factor applied to the current number of runners used to determine how many pods should be removed.",
									MarkdownDescription: "ScaleDownFactor is the multiplicative factor applied to the current number of runners used to determine how many pods should be removed.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scale_down_threshold": schema.StringAttribute{
									Description:         "ScaleDownThreshold is the percentage of busy runners less than which will trigger the hpa to scale the runners down.",
									MarkdownDescription: "ScaleDownThreshold is the percentage of busy runners less than which will trigger the hpa to scale the runners down.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scale_up_adjustment": schema.Int64Attribute{
									Description:         "ScaleUpAdjustment is the number of runners added on scale-up. You can only specify either ScaleUpFactor or ScaleUpAdjustment.",
									MarkdownDescription: "ScaleUpAdjustment is the number of runners added on scale-up. You can only specify either ScaleUpFactor or ScaleUpAdjustment.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scale_up_factor": schema.StringAttribute{
									Description:         "ScaleUpFactor is the multiplicative factor applied to the current number of runners used to determine how many pods should be added.",
									MarkdownDescription: "ScaleUpFactor is the multiplicative factor applied to the current number of runners used to determine how many pods should be added.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scale_up_threshold": schema.StringAttribute{
									Description:         "ScaleUpThreshold is the percentage of busy runners greater than which will trigger the hpa to scale runners up.",
									MarkdownDescription: "ScaleUpThreshold is the percentage of busy runners greater than which will trigger the hpa to scale runners up.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type is the type of metric to be used for autoscaling. It can be TotalNumberOfQueuedAndInProgressWorkflowRuns or PercentageRunnersBusy.",
									MarkdownDescription: "Type is the type of metric to be used for autoscaling. It can be TotalNumberOfQueuedAndInProgressWorkflowRuns or PercentageRunnersBusy.",
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

					"min_replicas": schema.Int64Attribute{
						Description:         "MinReplicas is the minimum number of replicas the deployment is allowed to scale",
						MarkdownDescription: "MinReplicas is the minimum number of replicas the deployment is allowed to scale",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scale_down_delay_seconds_after_scale_out": schema.Int64Attribute{
						Description:         "ScaleDownDelaySecondsAfterScaleUp is the approximate delay for a scale down followed by a scale up Used to prevent flapping (down->up->down->... loop)",
						MarkdownDescription: "ScaleDownDelaySecondsAfterScaleUp is the approximate delay for a scale down followed by a scale up Used to prevent flapping (down->up->down->... loop)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTargetRef is the reference to scaled resource like RunnerDeployment",
						MarkdownDescription: "ScaleTargetRef is the reference to scaled resource like RunnerDeployment",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind is the type of resource being referenced",
								MarkdownDescription: "Kind is the type of resource being referenced",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("RunnerDeployment", "RunnerSet"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of resource being referenced",
								MarkdownDescription: "Name is the name of resource being referenced",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"scale_up_triggers": schema.ListNestedAttribute{
						Description:         "ScaleUpTriggers is an experimental feature to increase the desired replicas by 1 on each webhook requested received by the webhookBasedAutoscaler. This feature requires you to also enable and deploy the webhookBasedAutoscaler onto your cluster. Note that the added runners remain until the next sync period at least, and they may or may not be used by GitHub Actions depending on the timing. They are intended to be used to gain 'resource slack' immediately after you receive a webhook from GitHub, so that you can loosely expect MinReplicas runners to be always available.",
						MarkdownDescription: "ScaleUpTriggers is an experimental feature to increase the desired replicas by 1 on each webhook requested received by the webhookBasedAutoscaler. This feature requires you to also enable and deploy the webhookBasedAutoscaler onto your cluster. Note that the added runners remain until the next sync period at least, and they may or may not be used by GitHub Actions depending on the timing. They are intended to be used to gain 'resource slack' immediately after you receive a webhook from GitHub, so that you can loosely expect MinReplicas runners to be always available.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"duration": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"github_event": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"check_run": schema.SingleNestedAttribute{
											Description:         "https://docs.github.com/en/actions/reference/events-that-trigger-workflows#check_run",
											MarkdownDescription: "https://docs.github.com/en/actions/reference/events-that-trigger-workflows#check_run",
											Attributes: map[string]schema.Attribute{
												"names": schema.ListAttribute{
													Description:         "Names is a list of GitHub Actions glob patterns. Any check_run event whose name matches one of patterns in the list can trigger autoscaling. Note that check_run name seem to equal to the job name you've defined in your actions workflow yaml file. So it is very likely that you can utilize this to trigger depending on the job.",
													MarkdownDescription: "Names is a list of GitHub Actions glob patterns. Any check_run event whose name matches one of patterns in the list can trigger autoscaling. Note that check_run name seem to equal to the job name you've defined in your actions workflow yaml file. So it is very likely that you can utilize this to trigger depending on the job.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"repositories": schema.ListAttribute{
													Description:         "Repositories is a list of GitHub repositories. Any check_run event whose repository matches one of repositories in the list can trigger autoscaling.",
													MarkdownDescription: "Repositories is a list of GitHub repositories. Any check_run event whose repository matches one of repositories in the list can trigger autoscaling.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"status": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"types": schema.ListAttribute{
													Description:         "One of: created, rerequested, or completed",
													MarkdownDescription: "One of: created, rerequested, or completed",
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

										"pull_request": schema.SingleNestedAttribute{
											Description:         "https://docs.github.com/en/actions/reference/events-that-trigger-workflows#pull_request",
											MarkdownDescription: "https://docs.github.com/en/actions/reference/events-that-trigger-workflows#pull_request",
											Attributes: map[string]schema.Attribute{
												"branches": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"types": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"push": schema.MapAttribute{
											Description:         "PushSpec is the condition for triggering scale-up on push event Also see https://docs.github.com/en/actions/reference/events-that-trigger-workflows#push",
											MarkdownDescription: "PushSpec is the condition for triggering scale-up on push event Also see https://docs.github.com/en/actions/reference/events-that-trigger-workflows#push",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"workflow_job": schema.MapAttribute{
											Description:         "https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job",
											MarkdownDescription: "https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#workflow_job",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"scheduled_overrides": schema.ListNestedAttribute{
						Description:         "ScheduledOverrides is the list of ScheduledOverride. It can be used to override a few fields of HorizontalRunnerAutoscalerSpec on schedule. The earlier a scheduled override is, the higher it is prioritized.",
						MarkdownDescription: "ScheduledOverrides is the list of ScheduledOverride. It can be used to override a few fields of HorizontalRunnerAutoscalerSpec on schedule. The earlier a scheduled override is, the higher it is prioritized.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"end_time": schema.StringAttribute{
									Description:         "EndTime is the time at which the first override ends.",
									MarkdownDescription: "EndTime is the time at which the first override ends.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"min_replicas": schema.Int64Attribute{
									Description:         "MinReplicas is the number of runners while overriding. If omitted, it doesn't override minReplicas.",
									MarkdownDescription: "MinReplicas is the number of runners while overriding. If omitted, it doesn't override minReplicas.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},

								"recurrence_rule": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"frequency": schema.StringAttribute{
											Description:         "Frequency is the name of a predefined interval of each recurrence. The valid values are 'Daily', 'Weekly', 'Monthly', and 'Yearly'. If empty, the corresponding override happens only once.",
											MarkdownDescription: "Frequency is the name of a predefined interval of each recurrence. The valid values are 'Daily', 'Weekly', 'Monthly', and 'Yearly'. If empty, the corresponding override happens only once.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Daily", "Weekly", "Monthly", "Yearly"),
											},
										},

										"until_time": schema.StringAttribute{
											Description:         "UntilTime is the time of the final recurrence. If empty, the schedule recurs forever.",
											MarkdownDescription: "UntilTime is the time of the final recurrence. If empty, the schedule recurs forever.",
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

								"start_time": schema.StringAttribute{
									Description:         "StartTime is the time at which the first override starts.",
									MarkdownDescription: "StartTime is the time at which the first override starts.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
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
		},
	}
}

func (r *ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest")

	var model ActionsSummerwindDevHorizontalRunnerAutoscalerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("actions.summerwind.dev/v1alpha1")
	model.Kind = pointer.String("HorizontalRunnerAutoscaler")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
