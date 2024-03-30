/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_k8s_io_v1

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
	_ datasource.DataSource = &AutoscalingK8SIoVerticalPodAutoscalerV1Manifest{}
)

func NewAutoscalingK8SIoVerticalPodAutoscalerV1Manifest() datasource.DataSource {
	return &AutoscalingK8SIoVerticalPodAutoscalerV1Manifest{}
}

type AutoscalingK8SIoVerticalPodAutoscalerV1Manifest struct{}

type AutoscalingK8SIoVerticalPodAutoscalerV1ManifestData struct {
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
		Recommenders *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"recommenders" json:"recommenders,omitempty"`
		ResourcePolicy *struct {
			ContainerPolicies *[]struct {
				ContainerName       *string            `tfsdk:"container_name" json:"containerName,omitempty"`
				ControlledResources *[]string          `tfsdk:"controlled_resources" json:"controlledResources,omitempty"`
				ControlledValues    *string            `tfsdk:"controlled_values" json:"controlledValues,omitempty"`
				MaxAllowed          *map[string]string `tfsdk:"max_allowed" json:"maxAllowed,omitempty"`
				MinAllowed          *map[string]string `tfsdk:"min_allowed" json:"minAllowed,omitempty"`
				Mode                *string            `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"container_policies" json:"containerPolicies,omitempty"`
		} `tfsdk:"resource_policy" json:"resourcePolicy,omitempty"`
		TargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		UpdatePolicy *struct {
			EvictionRequirements *[]struct {
				ChangeRequirement *string   `tfsdk:"change_requirement" json:"changeRequirement,omitempty"`
				Resources         *[]string `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"eviction_requirements" json:"evictionRequirements,omitempty"`
			MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
			UpdateMode  *string `tfsdk:"update_mode" json:"updateMode,omitempty"`
		} `tfsdk:"update_policy" json:"updatePolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AutoscalingK8SIoVerticalPodAutoscalerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest"
}

func (r *AutoscalingK8SIoVerticalPodAutoscalerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VerticalPodAutoscaler is the configuration for a vertical pod autoscaler, which automatically manages pod resources based on historical and real time resource utilization.",
		MarkdownDescription: "VerticalPodAutoscaler is the configuration for a vertical pod autoscaler, which automatically manages pod resources based on historical and real time resource utilization.",
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
				Description:         "Specification of the behavior of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.",
				MarkdownDescription: "Specification of the behavior of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.",
				Attributes: map[string]schema.Attribute{
					"recommenders": schema.ListNestedAttribute{
						Description:         "Recommender responsible for generating recommendation for this object. List should be empty (then the default recommender will generate the recommendation) or contain exactly one recommender.",
						MarkdownDescription: "Recommender responsible for generating recommendation for this object. List should be empty (then the default recommender will generate the recommendation) or contain exactly one recommender.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the recommender responsible for generating recommendation for this object.",
									MarkdownDescription: "Name of the recommender responsible for generating recommendation for this object.",
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

					"resource_policy": schema.SingleNestedAttribute{
						Description:         "Controls how the autoscaler computes recommended resources. The resource policy may be used to set constraints on the recommendations for individual containers. If any individual containers need to be excluded from getting the VPA recommendations, then it must be disabled explicitly by setting mode to 'Off' under containerPolicies. If not specified, the autoscaler computes recommended resources for all containers in the pod, without additional constraints.",
						MarkdownDescription: "Controls how the autoscaler computes recommended resources. The resource policy may be used to set constraints on the recommendations for individual containers. If any individual containers need to be excluded from getting the VPA recommendations, then it must be disabled explicitly by setting mode to 'Off' under containerPolicies. If not specified, the autoscaler computes recommended resources for all containers in the pod, without additional constraints.",
						Attributes: map[string]schema.Attribute{
							"container_policies": schema.ListNestedAttribute{
								Description:         "Per-container resource policies.",
								MarkdownDescription: "Per-container resource policies.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"container_name": schema.StringAttribute{
											Description:         "Name of the container or DefaultContainerResourcePolicy, in which case the policy is used by the containers that don't have their own policy specified.",
											MarkdownDescription: "Name of the container or DefaultContainerResourcePolicy, in which case the policy is used by the containers that don't have their own policy specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"controlled_resources": schema.ListAttribute{
											Description:         "Specifies the type of recommendations that will be computed (and possibly applied) by VPA. If not specified, the default of [ResourceCPU, ResourceMemory] will be used.",
											MarkdownDescription: "Specifies the type of recommendations that will be computed (and possibly applied) by VPA. If not specified, the default of [ResourceCPU, ResourceMemory] will be used.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"controlled_values": schema.StringAttribute{
											Description:         "Specifies which resource values should be controlled. The default is 'RequestsAndLimits'.",
											MarkdownDescription: "Specifies which resource values should be controlled. The default is 'RequestsAndLimits'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("RequestsAndLimits", "RequestsOnly"),
											},
										},

										"max_allowed": schema.MapAttribute{
											Description:         "Specifies the maximum amount of resources that will be recommended for the container. The default is no maximum.",
											MarkdownDescription: "Specifies the maximum amount of resources that will be recommended for the container. The default is no maximum.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_allowed": schema.MapAttribute{
											Description:         "Specifies the minimal amount of resources that will be recommended for the container. The default is no minimum.",
											MarkdownDescription: "Specifies the minimal amount of resources that will be recommended for the container. The default is no minimum.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Whether autoscaler is enabled for the container. The default is 'Auto'.",
											MarkdownDescription: "Whether autoscaler is enabled for the container. The default is 'Auto'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Auto", "Off"),
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef points to the controller managing the set of pods for the autoscaler to control - e.g. Deployment, StatefulSet. VerticalPodAutoscaler can be targeted at controller implementing scale subresource (the pod set is retrieved from the controller's ScaleStatus) or some well known controllers (e.g. for DaemonSet the pod set is read from the controller's spec). If VerticalPodAutoscaler cannot use specified target it will report ConfigUnsupported condition. Note that VerticalPodAutoscaler does not require full implementation of scale subresource - it will not use it to modify the replica count. The only thing retrieved is a label selector matching pods grouped by the target resource.",
						MarkdownDescription: "TargetRef points to the controller managing the set of pods for the autoscaler to control - e.g. Deployment, StatefulSet. VerticalPodAutoscaler can be targeted at controller implementing scale subresource (the pod set is retrieved from the controller's ScaleStatus) or some well known controllers (e.g. for DaemonSet the pod set is read from the controller's spec). If VerticalPodAutoscaler cannot use specified target it will report ConfigUnsupported condition. Note that VerticalPodAutoscaler does not require full implementation of scale subresource - it will not use it to modify the replica count. The only thing retrieved is a label selector matching pods grouped by the target resource.",
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

					"update_policy": schema.SingleNestedAttribute{
						Description:         "Describes the rules on how changes are applied to the pods. If not specified, all fields in the 'PodUpdatePolicy' are set to their default values.",
						MarkdownDescription: "Describes the rules on how changes are applied to the pods. If not specified, all fields in the 'PodUpdatePolicy' are set to their default values.",
						Attributes: map[string]schema.Attribute{
							"eviction_requirements": schema.ListNestedAttribute{
								Description:         "EvictionRequirements is a list of EvictionRequirements that need to evaluate to true in order for a Pod to be evicted. If more than one EvictionRequirement is specified, all of them need to be fulfilled to allow eviction.",
								MarkdownDescription: "EvictionRequirements is a list of EvictionRequirements that need to evaluate to true in order for a Pod to be evicted. If more than one EvictionRequirement is specified, all of them need to be fulfilled to allow eviction.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"change_requirement": schema.StringAttribute{
											Description:         "EvictionChangeRequirement refers to the relationship between the new target recommendation for a Pod and its current requests, what kind of change is necessary for the Pod to be evicted",
											MarkdownDescription: "EvictionChangeRequirement refers to the relationship between the new target recommendation for a Pod and its current requests, what kind of change is necessary for the Pod to be evicted",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TargetHigherThanRequests", "TargetLowerThanRequests"),
											},
										},

										"resources": schema.ListAttribute{
											Description:         "Resources is a list of one or more resources that the condition applies to. If more than one resource is given, the EvictionRequirement is fulfilled if at least one resource meets 'changeRequirement'.",
											MarkdownDescription: "Resources is a list of one or more resources that the condition applies to. If more than one resource is given, the EvictionRequirement is fulfilled if at least one resource meets 'changeRequirement'.",
											ElementType:         types.StringType,
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
								Description:         "Minimal number of replicas which need to be alive for Updater to attempt pod eviction (pending other checks like PDB). Only positive values are allowed. Overrides global '--min-replicas' flag.",
								MarkdownDescription: "Minimal number of replicas which need to be alive for Updater to attempt pod eviction (pending other checks like PDB). Only positive values are allowed. Overrides global '--min-replicas' flag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_mode": schema.StringAttribute{
								Description:         "Controls when autoscaler applies changes to the pod resources. The default is 'Auto'.",
								MarkdownDescription: "Controls when autoscaler applies changes to the pod resources. The default is 'Auto'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Off", "Initial", "Recreate", "Auto"),
								},
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
	}
}

func (r *AutoscalingK8SIoVerticalPodAutoscalerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest")

	var model AutoscalingK8SIoVerticalPodAutoscalerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("autoscaling.k8s.io/v1")
	model.Kind = pointer.String("VerticalPodAutoscaler")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
