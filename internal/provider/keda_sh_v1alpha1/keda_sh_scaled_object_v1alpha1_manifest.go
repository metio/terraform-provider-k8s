/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keda_sh_v1alpha1

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
	_ datasource.DataSource = &KedaShScaledObjectV1Alpha1Manifest{}
)

func NewKedaShScaledObjectV1Alpha1Manifest() datasource.DataSource {
	return &KedaShScaledObjectV1Alpha1Manifest{}
}

type KedaShScaledObjectV1Alpha1Manifest struct{}

type KedaShScaledObjectV1Alpha1ManifestData struct {
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
		Advanced *struct {
			HorizontalPodAutoscalerConfig *struct {
				Behavior *struct {
					ScaleDown *struct {
						Policies *[]struct {
							PeriodSeconds *int64  `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
							Type          *string `tfsdk:"type" json:"type,omitempty"`
							Value         *int64  `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"policies" json:"policies,omitempty"`
						SelectPolicy               *string `tfsdk:"select_policy" json:"selectPolicy,omitempty"`
						StabilizationWindowSeconds *int64  `tfsdk:"stabilization_window_seconds" json:"stabilizationWindowSeconds,omitempty"`
					} `tfsdk:"scale_down" json:"scaleDown,omitempty"`
					ScaleUp *struct {
						Policies *[]struct {
							PeriodSeconds *int64  `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
							Type          *string `tfsdk:"type" json:"type,omitempty"`
							Value         *int64  `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"policies" json:"policies,omitempty"`
						SelectPolicy               *string `tfsdk:"select_policy" json:"selectPolicy,omitempty"`
						StabilizationWindowSeconds *int64  `tfsdk:"stabilization_window_seconds" json:"stabilizationWindowSeconds,omitempty"`
					} `tfsdk:"scale_up" json:"scaleUp,omitempty"`
				} `tfsdk:"behavior" json:"behavior,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"horizontal_pod_autoscaler_config" json:"horizontalPodAutoscalerConfig,omitempty"`
			RestoreToOriginalReplicaCount *bool `tfsdk:"restore_to_original_replica_count" json:"restoreToOriginalReplicaCount,omitempty"`
		} `tfsdk:"advanced" json:"advanced,omitempty"`
		CooldownPeriod *int64 `tfsdk:"cooldown_period" json:"cooldownPeriod,omitempty"`
		Fallback       *struct {
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			Replicas         *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"fallback" json:"fallback,omitempty"`
		IdleReplicaCount *int64 `tfsdk:"idle_replica_count" json:"idleReplicaCount,omitempty"`
		MaxReplicaCount  *int64 `tfsdk:"max_replica_count" json:"maxReplicaCount,omitempty"`
		MinReplicaCount  *int64 `tfsdk:"min_replica_count" json:"minReplicaCount,omitempty"`
		PollingInterval  *int64 `tfsdk:"polling_interval" json:"pollingInterval,omitempty"`
		ScaleTargetRef   *struct {
			ApiVersion             *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			EnvSourceContainerName *string `tfsdk:"env_source_container_name" json:"envSourceContainerName,omitempty"`
			Kind                   *string `tfsdk:"kind" json:"kind,omitempty"`
			Name                   *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"scale_target_ref" json:"scaleTargetRef,omitempty"`
		Triggers *[]struct {
			AuthenticationRef *struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"authentication_ref" json:"authenticationRef,omitempty"`
			Metadata         *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			MetricType       *string            `tfsdk:"metric_type" json:"metricType,omitempty"`
			Name             *string            `tfsdk:"name" json:"name,omitempty"`
			Type             *string            `tfsdk:"type" json:"type,omitempty"`
			UseCachedMetrics *bool              `tfsdk:"use_cached_metrics" json:"useCachedMetrics,omitempty"`
		} `tfsdk:"triggers" json:"triggers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KedaShScaledObjectV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keda_sh_scaled_object_v1alpha1_manifest"
}

func (r *KedaShScaledObjectV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ScaledObject is a specification for a ScaledObject resource",
		MarkdownDescription: "ScaledObject is a specification for a ScaledObject resource",
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
				Description:         "ScaledObjectSpec is the spec for a ScaledObject resource",
				MarkdownDescription: "ScaledObjectSpec is the spec for a ScaledObject resource",
				Attributes: map[string]schema.Attribute{
					"advanced": schema.SingleNestedAttribute{
						Description:         "AdvancedConfig specifies advance scaling options",
						MarkdownDescription: "AdvancedConfig specifies advance scaling options",
						Attributes: map[string]schema.Attribute{
							"horizontal_pod_autoscaler_config": schema.SingleNestedAttribute{
								Description:         "HorizontalPodAutoscalerConfig specifies horizontal scale config",
								MarkdownDescription: "HorizontalPodAutoscalerConfig specifies horizontal scale config",
								Attributes: map[string]schema.Attribute{
									"behavior": schema.SingleNestedAttribute{
										Description:         "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
										MarkdownDescription: "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
										Attributes: map[string]schema.Attribute{
											"scale_down": schema.SingleNestedAttribute{
												Description:         "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).",
												MarkdownDescription: "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).",
												Attributes: map[string]schema.Attribute{
													"policies": schema.ListNestedAttribute{
														Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
														MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"period_seconds": schema.Int64Attribute{
																	Description:         "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
																	MarkdownDescription: "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type is used to specify the scaling policy.",
																	MarkdownDescription: "type is used to specify the scaling policy.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.Int64Attribute{
																	Description:         "value contains the amount of change which is permitted by the policy. It must be greater than zero",
																	MarkdownDescription: "value contains the amount of change which is permitted by the policy. It must be greater than zero",
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

													"select_policy": schema.StringAttribute{
														Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"stabilization_window_seconds": schema.Int64Attribute{
														Description:         "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														MarkdownDescription: "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"scale_up": schema.SingleNestedAttribute{
												Description:         "scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of: * increase no more than 4 pods per 60 seconds * double the number of pods per 60 seconds No stabilization is used.",
												MarkdownDescription: "scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of: * increase no more than 4 pods per 60 seconds * double the number of pods per 60 seconds No stabilization is used.",
												Attributes: map[string]schema.Attribute{
													"policies": schema.ListNestedAttribute{
														Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
														MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"period_seconds": schema.Int64Attribute{
																	Description:         "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
																	MarkdownDescription: "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type is used to specify the scaling policy.",
																	MarkdownDescription: "type is used to specify the scaling policy.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.Int64Attribute{
																	Description:         "value contains the amount of change which is permitted by the policy. It must be greater than zero",
																	MarkdownDescription: "value contains the amount of change which is permitted by the policy. It must be greater than zero",
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

													"select_policy": schema.StringAttribute{
														Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"stabilization_window_seconds": schema.Int64Attribute{
														Description:         "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														MarkdownDescription: "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"restore_to_original_replica_count": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cooldown_period": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"fallback": schema.SingleNestedAttribute{
						Description:         "Fallback is the spec for fallback options",
						MarkdownDescription: "Fallback is the spec for fallback options",
						Attributes: map[string]schema.Attribute{
							"failure_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
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

					"idle_replica_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_replica_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_replica_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"polling_interval": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTarget holds the reference to the scale target Object",
						MarkdownDescription: "ScaleTarget holds the reference to the scale target Object",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_source_container_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"triggers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authentication_ref": schema.SingleNestedAttribute{
									Description:         "AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that is used to authenticate the scaler with the environment",
									MarkdownDescription: "AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that is used to authenticate the scaler with the environment",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the resource being referred to. Defaults to TriggerAuthentication.",
											MarkdownDescription: "Kind of the resource being referred to. Defaults to TriggerAuthentication.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

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

								"metadata": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"metric_type": schema.StringAttribute{
									Description:         "MetricTargetType specifies the type of metric being targeted, and should be either 'Value', 'AverageValue', or 'Utilization'",
									MarkdownDescription: "MetricTargetType specifies the type of metric being targeted, and should be either 'Value', 'AverageValue', or 'Utilization'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"use_cached_metrics": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
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

func (r *KedaShScaledObjectV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keda_sh_scaled_object_v1alpha1_manifest")

	var model KedaShScaledObjectV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("keda.sh/v1alpha1")
	model.Kind = pointer.String("ScaledObject")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
