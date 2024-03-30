/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_karmada_io_v1alpha1

import (
	"context"
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
	_ datasource.DataSource = &AutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest{}
)

func NewAutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest() datasource.DataSource {
	return &AutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest{}
}

type AutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest struct{}

type AutoscalingKarmadaIoFederatedHpaV1Alpha1ManifestData struct {
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
		MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
		Metrics     *[]struct {
			ContainerResource *struct {
				Container *string `tfsdk:"container" json:"container,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Target    *struct {
					AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
					AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
					Value              *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"container_resource" json:"containerResource,omitempty"`
			External *struct {
				Metric *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"metric" json:"metric,omitempty"`
				Target *struct {
					AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
					AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
					Value              *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"external" json:"external,omitempty"`
			Object *struct {
				DescribedObject *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"described_object" json:"describedObject,omitempty"`
				Metric *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"metric" json:"metric,omitempty"`
				Target *struct {
					AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
					AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
					Value              *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"object" json:"object,omitempty"`
			Pods *struct {
				Metric *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"metric" json:"metric,omitempty"`
				Target *struct {
					AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
					AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
					Value              *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"pods" json:"pods,omitempty"`
			Resource *struct {
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Target *struct {
					AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
					AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
					Value              *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"resource" json:"resource,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		MinReplicas    *int64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		ScaleTargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"scale_target_ref" json:"scaleTargetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_karmada_io_federated_hpa_v1alpha1_manifest"
}

func (r *AutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FederatedHPA is centralized HPA that can aggregate the metrics in multiple clusters. When the system load increases, it will query the metrics from multiple clusters and scales up the replicas. When the system load decreases, it will query the metrics from multiple clusters and scales down the replicas. After the replicas are scaled up/down, karmada-scheduler will schedule the replicas based on the policy.",
		MarkdownDescription: "FederatedHPA is centralized HPA that can aggregate the metrics in multiple clusters. When the system load increases, it will query the metrics from multiple clusters and scales up the replicas. When the system load decreases, it will query the metrics from multiple clusters and scales down the replicas. After the replicas are scaled up/down, karmada-scheduler will schedule the replicas based on the policy.",
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
				Description:         "Spec is the specification of the FederatedHPA.",
				MarkdownDescription: "Spec is the specification of the FederatedHPA.",
				Attributes: map[string]schema.Attribute{
					"behavior": schema.SingleNestedAttribute{
						Description:         "Behavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively). If not set, the default HPAScalingRules for scale up and scale down are used.",
						MarkdownDescription: "Behavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively). If not set, the default HPAScalingRules for scale up and scale down are used.",
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

					"max_replicas": schema.Int64Attribute{
						Description:         "MaxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.",
						MarkdownDescription: "MaxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"metrics": schema.ListNestedAttribute{
						Description:         "Metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used). The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods. Ergo, metrics used must decrease as the pod count is increased, and vice-versa. See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.",
						MarkdownDescription: "Metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used). The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods. Ergo, metrics used must decrease as the pod count is increased, and vice-versa. See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"container_resource": schema.SingleNestedAttribute{
									Description:         "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.",
									MarkdownDescription: "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.",
									Attributes: map[string]schema.Attribute{
										"container": schema.StringAttribute{
											Description:         "container is the name of the container in the pods of the scaling target",
											MarkdownDescription: "container is the name of the container in the pods of the scaling target",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name is the name of the resource in question.",
											MarkdownDescription: "name is the name of the resource in question.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target": schema.SingleNestedAttribute{
											Description:         "target specifies the target value for the given metric",
											MarkdownDescription: "target specifies the target value for the given metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
													MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "value is the target value of the metric (as a quantity).",
													MarkdownDescription: "value is the target value of the metric (as a quantity).",
													Required:            false,
													Optional:            true,
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

								"external": schema.SingleNestedAttribute{
									Description:         "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
									MarkdownDescription: "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
									Attributes: map[string]schema.Attribute{
										"metric": schema.SingleNestedAttribute{
											Description:         "metric identifies the target metric by name and selector",
											MarkdownDescription: "metric identifies the target metric by name and selector",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the given metric",
													MarkdownDescription: "name is the name of the given metric",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
													MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

										"target": schema.SingleNestedAttribute{
											Description:         "target specifies the target value for the given metric",
											MarkdownDescription: "target specifies the target value for the given metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
													MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "value is the target value of the metric (as a quantity).",
													MarkdownDescription: "value is the target value of the metric (as a quantity).",
													Required:            false,
													Optional:            true,
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

								"object": schema.SingleNestedAttribute{
									Description:         "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
									MarkdownDescription: "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
									Attributes: map[string]schema.Attribute{
										"described_object": schema.SingleNestedAttribute{
											Description:         "describedObject specifies the descriptions of a object,such as kind,name apiVersion",
											MarkdownDescription: "describedObject specifies the descriptions of a object,such as kind,name apiVersion",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "apiVersion is the API version of the referent",
													MarkdownDescription: "apiVersion is the API version of the referent",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"metric": schema.SingleNestedAttribute{
											Description:         "metric identifies the target metric by name and selector",
											MarkdownDescription: "metric identifies the target metric by name and selector",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the given metric",
													MarkdownDescription: "name is the name of the given metric",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
													MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

										"target": schema.SingleNestedAttribute{
											Description:         "target specifies the target value for the given metric",
											MarkdownDescription: "target specifies the target value for the given metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
													MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "value is the target value of the metric (as a quantity).",
													MarkdownDescription: "value is the target value of the metric (as a quantity).",
													Required:            false,
													Optional:            true,
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

								"pods": schema.SingleNestedAttribute{
									Description:         "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
									MarkdownDescription: "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
									Attributes: map[string]schema.Attribute{
										"metric": schema.SingleNestedAttribute{
											Description:         "metric identifies the target metric by name and selector",
											MarkdownDescription: "metric identifies the target metric by name and selector",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the given metric",
													MarkdownDescription: "name is the name of the given metric",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
													MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

										"target": schema.SingleNestedAttribute{
											Description:         "target specifies the target value for the given metric",
											MarkdownDescription: "target specifies the target value for the given metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
													MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "value is the target value of the metric (as a quantity).",
													MarkdownDescription: "value is the target value of the metric (as a quantity).",
													Required:            false,
													Optional:            true,
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

								"resource": schema.SingleNestedAttribute{
									Description:         "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.",
									MarkdownDescription: "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is the name of the resource in question.",
											MarkdownDescription: "name is the name of the resource in question.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target": schema.SingleNestedAttribute{
											Description:         "target specifies the target value for the given metric",
											MarkdownDescription: "target specifies the target value for the given metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
													MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "value is the target value of the metric (as a quantity).",
													MarkdownDescription: "value is the target value of the metric (as a quantity).",
													Required:            false,
													Optional:            true,
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

								"type": schema.StringAttribute{
									Description:         "type is the type of metric source.  It should be one of 'ContainerResource', 'External', 'Object', 'Pods' or 'Resource', each mapping to a matching field in the object. Note: 'ContainerResource' type is available on when the feature-gate HPAContainerMetrics is enabled",
									MarkdownDescription: "type is the type of metric source.  It should be one of 'ContainerResource', 'External', 'Object', 'Pods' or 'Resource', each mapping to a matching field in the object. Note: 'ContainerResource' type is available on when the feature-gate HPAContainerMetrics is enabled",
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

					"min_replicas": schema.Int64Attribute{
						Description:         "MinReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod.",
						MarkdownDescription: "MinReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.",
						MarkdownDescription: "ScaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "apiVersion is the API version of the referent",
								MarkdownDescription: "apiVersion is the API version of the referent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *AutoscalingKarmadaIoFederatedHpaV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_karmada_io_federated_hpa_v1alpha1_manifest")

	var model AutoscalingKarmadaIoFederatedHpaV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("autoscaling.karmada.io/v1alpha1")
	model.Kind = pointer.String("FederatedHPA")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
