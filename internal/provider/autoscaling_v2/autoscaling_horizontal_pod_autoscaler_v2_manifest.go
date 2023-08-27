/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_v2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AutoscalingHorizontalPodAutoscalerV2Manifest{}
)

func NewAutoscalingHorizontalPodAutoscalerV2Manifest() datasource.DataSource {
	return &AutoscalingHorizontalPodAutoscalerV2Manifest{}
}

type AutoscalingHorizontalPodAutoscalerV2Manifest struct{}

type AutoscalingHorizontalPodAutoscalerV2ManifestData struct {
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

func (r *AutoscalingHorizontalPodAutoscalerV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_horizontal_pod_autoscaler_v2_manifest"
}

func (r *AutoscalingHorizontalPodAutoscalerV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.",
		MarkdownDescription: "HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.",
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
				Description:         "HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.",
				MarkdownDescription: "HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.",
				Attributes: map[string]schema.Attribute{
					"behavior": schema.SingleNestedAttribute{
						Description:         "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
						MarkdownDescription: "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
						Attributes: map[string]schema.Attribute{
							"scale_down": schema.SingleNestedAttribute{
								Description:         "HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.",
								MarkdownDescription: "HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.",
								Attributes: map[string]schema.Attribute{
									"policies": schema.ListNestedAttribute{
										Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
										MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"period_seconds": schema.Int64Attribute{
													Description:         "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
													MarkdownDescription: "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type is used to specify the scaling policy.",
													MarkdownDescription: "Type is used to specify the scaling policy.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.Int64Attribute{
													Description:         "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
													MarkdownDescription: "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
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
										Description:         "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
										MarkdownDescription: "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
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
								Description:         "HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.",
								MarkdownDescription: "HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.",
								Attributes: map[string]schema.Attribute{
									"policies": schema.ListNestedAttribute{
										Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
										MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"period_seconds": schema.Int64Attribute{
													Description:         "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
													MarkdownDescription: "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type is used to specify the scaling policy.",
													MarkdownDescription: "Type is used to specify the scaling policy.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.Int64Attribute{
													Description:         "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
													MarkdownDescription: "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
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
										Description:         "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
										MarkdownDescription: "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
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
						Description:         "maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.",
						MarkdownDescription: "maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"metrics": schema.ListNestedAttribute{
						Description:         "metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.",
						MarkdownDescription: "metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"container_resource": schema.SingleNestedAttribute{
									Description:         "ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.  Only one 'target' type should be set.",
									MarkdownDescription: "ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.  Only one 'target' type should be set.",
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
											Description:         "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											MarkdownDescription: "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
									Description:         "ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
									MarkdownDescription: "ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
									Attributes: map[string]schema.Attribute{
										"metric": schema.SingleNestedAttribute{
											Description:         "MetricIdentifier defines the name and optionally selector for a metric",
											MarkdownDescription: "MetricIdentifier defines the name and optionally selector for a metric",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the given metric",
													MarkdownDescription: "name is the name of the given metric",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
													MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
											Description:         "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											MarkdownDescription: "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
									Description:         "ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).",
									MarkdownDescription: "ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).",
									Attributes: map[string]schema.Attribute{
										"described_object": schema.SingleNestedAttribute{
											Description:         "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
											MarkdownDescription: "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent",
													MarkdownDescription: "API version of the referent",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
													MarkdownDescription: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
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
											Description:         "MetricIdentifier defines the name and optionally selector for a metric",
											MarkdownDescription: "MetricIdentifier defines the name and optionally selector for a metric",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the given metric",
													MarkdownDescription: "name is the name of the given metric",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
													MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
											Description:         "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											MarkdownDescription: "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
									Description:         "PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.",
									MarkdownDescription: "PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.",
									Attributes: map[string]schema.Attribute{
										"metric": schema.SingleNestedAttribute{
											Description:         "MetricIdentifier defines the name and optionally selector for a metric",
											MarkdownDescription: "MetricIdentifier defines the name and optionally selector for a metric",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the name of the given metric",
													MarkdownDescription: "name is the name of the given metric",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
													MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
											Description:         "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											MarkdownDescription: "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
									Description:         "ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.  Only one 'target' type should be set.",
									MarkdownDescription: "ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the 'pods' source.  Only one 'target' type should be set.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is the name of the resource in question.",
											MarkdownDescription: "name is the name of the resource in question.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target": schema.SingleNestedAttribute{
											Description:         "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											MarkdownDescription: "MetricTarget defines the target value, average value, or average utilization of a specific metric",
											Attributes: map[string]schema.Attribute{
												"average_utilization": schema.Int64Attribute{
													Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"average_value": schema.StringAttribute{
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
													Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
													MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.The serialization format is:''' <quantity>        ::= <signedNumber><suffix>	(Note that <suffix> may be empty, from the '' case in <decimalSI>.)<digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= '+' | '-' <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei	(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)<decimalSI>       ::= m | '' | k | M | G | T | P | E	(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)<decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> '''No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:- No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.The sign will be omitted unless the number is negative.Examples:- 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi'Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
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
						Description:         "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
						MarkdownDescription: "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
						MarkdownDescription: "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent",
								MarkdownDescription: "API version of the referent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
								MarkdownDescription: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
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
		},
	}
}

func (r *AutoscalingHorizontalPodAutoscalerV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_horizontal_pod_autoscaler_v2_manifest")

	var model AutoscalingHorizontalPodAutoscalerV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("autoscaling/v2")
	model.Kind = pointer.String("HorizontalPodAutoscaler")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
