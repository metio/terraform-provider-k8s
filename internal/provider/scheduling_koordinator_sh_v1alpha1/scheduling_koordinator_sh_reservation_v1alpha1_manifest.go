/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scheduling_koordinator_sh_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SchedulingKoordinatorShReservationV1Alpha1Manifest{}
)

func NewSchedulingKoordinatorShReservationV1Alpha1Manifest() datasource.DataSource {
	return &SchedulingKoordinatorShReservationV1Alpha1Manifest{}
}

type SchedulingKoordinatorShReservationV1Alpha1Manifest struct{}

type SchedulingKoordinatorShReservationV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AllocateOnce   *bool   `tfsdk:"allocate_once" json:"allocateOnce,omitempty"`
		AllocatePolicy *string `tfsdk:"allocate_policy" json:"allocatePolicy,omitempty"`
		Expires        *string `tfsdk:"expires" json:"expires,omitempty"`
		Owners         *[]struct {
			Controller *struct {
				ApiVersion         *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				BlockOwnerDeletion *bool   `tfsdk:"block_owner_deletion" json:"blockOwnerDeletion,omitempty"`
				Controller         *bool   `tfsdk:"controller" json:"controller,omitempty"`
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Uid                *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"controller" json:"controller,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Object *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"object" json:"object,omitempty"`
		} `tfsdk:"owners" json:"owners,omitempty"`
		PreAllocation *bool `tfsdk:"pre_allocation" json:"preAllocation,omitempty"`
		Taints        *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
		Template      *map[string]string `tfsdk:"template" json:"template,omitempty"`
		Ttl           *string            `tfsdk:"ttl" json:"ttl,omitempty"`
		Unschedulable *bool              `tfsdk:"unschedulable" json:"unschedulable,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SchedulingKoordinatorShReservationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scheduling_koordinator_sh_reservation_v1alpha1_manifest"
}

func (r *SchedulingKoordinatorShReservationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Reservation is the Schema for the reservation API.A Reservation object is non-namespaced.Any namespaced affinity/anti-affinity of reservation scheduling can be specified in the spec.template.",
		MarkdownDescription: "Reservation is the Schema for the reservation API.A Reservation object is non-namespaced.Any namespaced affinity/anti-affinity of reservation scheduling can be specified in the spec.template.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"allocate_once": schema.BoolAttribute{
						Description:         "When 'AllocateOnce' is set, the reserved resources are only available for the first owner who allocates successfullyand are not allocatable to other owners anymore. Defaults to true.",
						MarkdownDescription: "When 'AllocateOnce' is set, the reserved resources are only available for the first owner who allocates successfullyand are not allocatable to other owners anymore. Defaults to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allocate_policy": schema.StringAttribute{
						Description:         "AllocatePolicy represents the allocation policy of reserved resources that Reservation expects.",
						MarkdownDescription: "AllocatePolicy represents the allocation policy of reserved resources that Reservation expects.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Aligned", "Restricted"),
						},
					},

					"expires": schema.StringAttribute{
						Description:         "Expired timestamp when the reservation is expected to expire.If both 'expires' and 'ttl' are set, 'expires' is checked first.'expires' and 'ttl' are mutually exclusive. Defaults to being set dynamically at runtime based on the 'ttl'.",
						MarkdownDescription: "Expired timestamp when the reservation is expected to expire.If both 'expires' and 'ttl' are set, 'expires' is checked first.'expires' and 'ttl' are mutually exclusive. Defaults to being set dynamically at runtime based on the 'ttl'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.DateTime64Validator(),
						},
					},

					"owners": schema.ListNestedAttribute{
						Description:         "Specify the owners who can allocate the reserved resources.Multiple owner selectors and ORed.",
						MarkdownDescription: "Specify the owners who can allocate the reserved resources.Multiple owner selectors and ORed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"controller": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"block_owner_deletion": schema.BoolAttribute{
											Description:         "If true, AND if the owner has the 'foregroundDeletion' finalizer, thenthe owner cannot be deleted from the key-value store until thisreference is removed.See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletionfor how the garbage collector interacts with this field and enforces the foreground deletion.Defaults to false.To set this field, a user needs 'delete' permission of the owner,otherwise 422 (Unprocessable Entity) will be returned.",
											MarkdownDescription: "If true, AND if the owner has the 'foregroundDeletion' finalizer, thenthe owner cannot be deleted from the key-value store until thisreference is removed.See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletionfor how the garbage collector interacts with this field and enforces the foreground deletion.Defaults to false.To set this field, a user needs 'delete' permission of the owner,otherwise 422 (Unprocessable Entity) will be returned.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"controller": schema.BoolAttribute{
											Description:         "If true, this reference points to the managing controller.",
											MarkdownDescription: "If true, this reference points to the managing controller.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uid": schema.StringAttribute{
											Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
											MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
									Validators: []validator.Object{
										objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("label_selector"), path.MatchRelative().AtParent().AtName("object")),
									},
								},

								"label_selector": schema.SingleNestedAttribute{
									Description:         "A label selector is a label query over a set of resources. The result of matchLabels andmatchExpressions are ANDed. An empty label selector matches all objects. A nulllabel selector matches no objects.",
									MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels andmatchExpressions are ANDed. An empty label selector matches all objects. A nulllabel selector matches no objects.",
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
									Validators: []validator.Object{
										objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("controller"), path.MatchRelative().AtParent().AtName("object")),
									},
								},

								"object": schema.SingleNestedAttribute{
									Description:         "Multiple field selectors are ANDed.",
									MarkdownDescription: "Multiple field selectors are ANDed.",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"field_path": schema.StringAttribute{
											Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
											MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
											MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_version": schema.StringAttribute{
											Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
											MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uid": schema.StringAttribute{
											Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
											MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
									Validators: []validator.Object{
										objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("controller"), path.MatchRelative().AtParent().AtName("label_selector")),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"pre_allocation": schema.BoolAttribute{
						Description:         "By default, the resources requirements of reservation (specified in 'template.spec') is filtered by whether thenode has sufficient free resources (i.e. Reservation Request <  Node Free).When 'preAllocation' is set, the scheduler will skip this validation and allow overcommitment. The scheduledreservation would be waiting to be available until free resources are sufficient.",
						MarkdownDescription: "By default, the resources requirements of reservation (specified in 'template.spec') is filtered by whether thenode has sufficient free resources (i.e. Reservation Request <  Node Free).When 'preAllocation' is set, the scheduler will skip this validation and allow overcommitment. The scheduledreservation would be waiting to be available until free resources are sufficient.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"taints": schema.ListNestedAttribute{
						Description:         "Specifies the reservation's taints. This can be toleranted by the reservation tolerance.Eviction is not supported for NoExecute taints",
						MarkdownDescription: "Specifies the reservation's taints. This can be toleranted by the reservation tolerance.Eviction is not supported for NoExecute taints",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"time_added": schema.StringAttribute{
									Description:         "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
									MarkdownDescription: "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"value": schema.StringAttribute{
									Description:         "The taint value corresponding to the taint key.",
									MarkdownDescription: "The taint value corresponding to the taint key.",
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

					"template": schema.MapAttribute{
						Description:         "Template defines the scheduling requirements (resources, affinities, images, ...) processed by the scheduler justlike a normal pod.If the 'template.spec.nodeName' is specified, the scheduler will not choose another node but reserve resources onthe specified node.",
						MarkdownDescription: "Template defines the scheduling requirements (resources, affinities, images, ...) processed by the scheduler justlike a normal pod.If the 'template.spec.nodeName' is specified, the scheduler will not choose another node but reserve resources onthe specified node.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ttl": schema.StringAttribute{
						Description:         "Time-to-Live period for the reservation.'expires' and 'ttl' are mutually exclusive. Defaults to 24h. Set 0 to disable expiration.",
						MarkdownDescription: "Time-to-Live period for the reservation.'expires' and 'ttl' are mutually exclusive. Defaults to 24h. Set 0 to disable expiration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unschedulable": schema.BoolAttribute{
						Description:         "Unschedulable controls reservation schedulability of new pods. By default, reservation is schedulable.",
						MarkdownDescription: "Unschedulable controls reservation schedulability of new pods. By default, reservation is schedulable.",
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

func (r *SchedulingKoordinatorShReservationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scheduling_koordinator_sh_reservation_v1alpha1_manifest")

	var model SchedulingKoordinatorShReservationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("scheduling.koordinator.sh/v1alpha1")
	model.Kind = pointer.String("Reservation")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
