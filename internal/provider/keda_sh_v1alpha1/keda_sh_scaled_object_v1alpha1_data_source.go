/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keda_sh_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &KedaShScaledObjectV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KedaShScaledObjectV1Alpha1DataSource{}
)

func NewKedaShScaledObjectV1Alpha1DataSource() datasource.DataSource {
	return &KedaShScaledObjectV1Alpha1DataSource{}
}

type KedaShScaledObjectV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KedaShScaledObjectV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *KedaShScaledObjectV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keda_sh_scaled_object_v1alpha1"
}

func (r *KedaShScaledObjectV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"type": schema.StringAttribute{
																	Description:         "type is used to specify the scaling policy.",
																	MarkdownDescription: "type is used to specify the scaling policy.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.Int64Attribute{
																	Description:         "value contains the amount of change which is permitted by the policy. It must be greater than zero",
																	MarkdownDescription: "value contains the amount of change which is permitted by the policy. It must be greater than zero",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"select_policy": schema.StringAttribute{
														Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"stabilization_window_seconds": schema.Int64Attribute{
														Description:         "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														MarkdownDescription: "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
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
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"type": schema.StringAttribute{
																	Description:         "type is used to specify the scaling policy.",
																	MarkdownDescription: "type is used to specify the scaling policy.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.Int64Attribute{
																	Description:         "value contains the amount of change which is permitted by the policy. It must be greater than zero",
																	MarkdownDescription: "value contains the amount of change which is permitted by the policy. It must be greater than zero",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"select_policy": schema.StringAttribute{
														Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"stabilization_window_seconds": schema.Int64Attribute{
														Description:         "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														MarkdownDescription: "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"restore_to_original_replica_count": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"cooldown_period": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"fallback": schema.SingleNestedAttribute{
						Description:         "Fallback is the spec for fallback options",
						MarkdownDescription: "Fallback is the spec for fallback options",
						Attributes: map[string]schema.Attribute{
							"failure_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"idle_replica_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_replica_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"min_replica_count": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"polling_interval": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTarget holds the reference to the scale target Object",
						MarkdownDescription: "ScaleTarget holds the reference to the scale target Object",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env_source_container_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"metadata": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"metric_type": schema.StringAttribute{
									Description:         "MetricTargetType specifies the type of metric being targeted, and should be either 'Value', 'AverageValue', or 'Utilization'",
									MarkdownDescription: "MetricTargetType specifies the type of metric being targeted, and should be either 'Value', 'AverageValue', or 'Utilization'",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"use_cached_metrics": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *KedaShScaledObjectV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KedaShScaledObjectV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_keda_sh_scaled_object_v1alpha1")

	var data KedaShScaledObjectV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "keda.sh", Version: "v1alpha1", Resource: "ScaledObject"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KedaShScaledObjectV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("keda.sh/v1alpha1")
	data.Kind = pointer.String("ScaledObject")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
