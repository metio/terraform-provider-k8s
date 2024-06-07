/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_koordinator_sh_v1alpha1

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
	_ datasource.DataSource = &ConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest{}
)

func NewConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest() datasource.DataSource {
	return &ConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest{}
}

type ConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest struct{}

type ConfigKoordinatorShClusterColocationProfileV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AnnotationKeysMapping *map[string]string `tfsdk:"annotation_keys_mapping" json:"annotationKeysMapping,omitempty"`
		Annotations           *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
		KoordinatorPriority   *int64             `tfsdk:"koordinator_priority" json:"koordinatorPriority,omitempty"`
		LabelKeysMapping      *map[string]string `tfsdk:"label_keys_mapping" json:"labelKeysMapping,omitempty"`
		Labels                *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		NamespaceSelector     *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Patch             *map[string]string `tfsdk:"patch" json:"patch,omitempty"`
		PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		Probability       *string            `tfsdk:"probability" json:"probability,omitempty"`
		QosClass          *string            `tfsdk:"qos_class" json:"qosClass,omitempty"`
		SchedulerName     *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		Selector          *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest"
}

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterColocationProfile is the Schema for the ClusterColocationProfile API",
		MarkdownDescription: "ClusterColocationProfile is the Schema for the ClusterColocationProfile API",
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
				Description:         "ClusterColocationProfileSpec is a description of a ClusterColocationProfile.",
				MarkdownDescription: "ClusterColocationProfileSpec is a description of a ClusterColocationProfile.",
				Attributes: map[string]schema.Attribute{
					"annotation_keys_mapping": schema.MapAttribute{
						Description:         "AnnotationKeysMapping describes the annotations that needs to inject into Pod.Annotations with the same values.It sets the Pod.Annotations[AnnotationsToAnnotations[k]] = Pod.Annotations[k] for each key k.",
						MarkdownDescription: "AnnotationKeysMapping describes the annotations that needs to inject into Pod.Annotations with the same values.It sets the Pod.Annotations[AnnotationsToAnnotations[k]] = Pod.Annotations[k] for each key k.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"annotations": schema.MapAttribute{
						Description:         "Annotations describes the k/v pair that needs to inject into Pod.Annotations",
						MarkdownDescription: "Annotations describes the k/v pair that needs to inject into Pod.Annotations",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"koordinator_priority": schema.Int64Attribute{
						Description:         "KoordinatorPriority defines the Pod sub-priority in Koordinator.The priority value will be injected into Pod as label koordinator.sh/priority.Various Koordinator components determine the priority of the Podin the Koordinator through KoordinatorPriority and the priority value in PriorityClassName.The higher the value, the higher the priority.",
						MarkdownDescription: "KoordinatorPriority defines the Pod sub-priority in Koordinator.The priority value will be injected into Pod as label koordinator.sh/priority.Various Koordinator components determine the priority of the Podin the Koordinator through KoordinatorPriority and the priority value in PriorityClassName.The higher the value, the higher the priority.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_keys_mapping": schema.MapAttribute{
						Description:         "LabelKeysMapping describes the labels that needs to inject into Pod.Labels with the same values.It sets the Pod.Labels[LabelsToLabels[k]] = Pod.Labels[k] for each key k.",
						MarkdownDescription: "LabelKeysMapping describes the labels that needs to inject into Pod.Labels with the same values.It sets the Pod.Labels[LabelsToLabels[k]] = Pod.Labels[k] for each key k.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"labels": schema.MapAttribute{
						Description:         "Labels describes the k/v pair that needs to inject into Pod.Labels",
						MarkdownDescription: "Labels describes the k/v pair that needs to inject into Pod.Labels",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "NamespaceSelector decides whether to mutate/validate Pods if thenamespace matches the selector.Default to the empty LabelSelector, which matches everything.",
						MarkdownDescription: "NamespaceSelector decides whether to mutate/validate Pods if thenamespace matches the selector.Default to the empty LabelSelector, which matches everything.",
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
											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"patch": schema.MapAttribute{
						Description:         "Patch indicates patching podTemplate that will be injected to the Pod.",
						MarkdownDescription: "Patch indicates patching podTemplate that will be injected to the Pod.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "If specified, the priorityClassName and the priority value defined in PriorityClasswill be injected into the Pod.The PriorityClassName, priority value in PriorityClassName andKoordinatorPriority will affect the scheduling, preemption andother behaviors of Koordinator system.",
						MarkdownDescription: "If specified, the priorityClassName and the priority value defined in PriorityClasswill be injected into the Pod.The PriorityClassName, priority value in PriorityClassName andKoordinatorPriority will affect the scheduling, preemption andother behaviors of Koordinator system.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"probability": schema.StringAttribute{
						Description:         "Probability indicates profile will make effect with a probability.",
						MarkdownDescription: "Probability indicates profile will make effect with a probability.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"qos_class": schema.StringAttribute{
						Description:         "QoSClass describes the type of Koordinator QoS that the Pod is running.The value will be injected into Pod as label koordinator.sh/qosClass.Options are LSE/LSR/LS/BE/SYSTEM.",
						MarkdownDescription: "QoSClass describes the type of Koordinator QoS that the Pod is running.The value will be injected into Pod as label koordinator.sh/qosClass.Options are LSE/LSR/LS/BE/SYSTEM.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("LSE", "LSR", "LS", "BE", "SYSTEM"),
						},
					},

					"scheduler_name": schema.StringAttribute{
						Description:         "If specified, the pod will be dispatched by specified scheduler.",
						MarkdownDescription: "If specified, the pod will be dispatched by specified scheduler.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector decides whether to mutate/validate Pods if thePod matches the selector.Default to the empty LabelSelector, which matches everything.",
						MarkdownDescription: "Selector decides whether to mutate/validate Pods if thePod matches the selector.Default to the empty LabelSelector, which matches everything.",
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
											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ConfigKoordinatorShClusterColocationProfileV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest")

	var model ConfigKoordinatorShClusterColocationProfileV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.koordinator.sh/v1alpha1")
	model.Kind = pointer.String("ClusterColocationProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
