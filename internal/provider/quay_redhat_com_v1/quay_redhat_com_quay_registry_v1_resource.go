/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package quay_redhat_com_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	_ resource.Resource                = &QuayRedhatComQuayRegistryV1Resource{}
	_ resource.ResourceWithConfigure   = &QuayRedhatComQuayRegistryV1Resource{}
	_ resource.ResourceWithImportState = &QuayRedhatComQuayRegistryV1Resource{}
)

func NewQuayRedhatComQuayRegistryV1Resource() resource.Resource {
	return &QuayRedhatComQuayRegistryV1Resource{}
}

type QuayRedhatComQuayRegistryV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type QuayRedhatComQuayRegistryV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Components *[]struct {
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Managed   *bool   `tfsdk:"managed" json:"managed,omitempty"`
			Overrides *struct {
				Affinity *struct {
					NodeAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							Preference *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchFields *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_fields" json:"matchFields,omitempty"`
							} `tfsdk:"preference" json:"preference,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
						RequiredDuringSchedulingIgnoredDuringExecution *struct {
							NodeSelectorTerms *[]struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchFields *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_fields" json:"matchFields,omitempty"`
							} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					PodAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Env         *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				VolumeSize *string            `tfsdk:"volume_size" json:"volumeSize,omitempty"`
			} `tfsdk:"overrides" json:"overrides,omitempty"`
		} `tfsdk:"components" json:"components,omitempty"`
		ConfigBundleSecret *string `tfsdk:"config_bundle_secret" json:"configBundleSecret,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *QuayRedhatComQuayRegistryV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_quay_redhat_com_quay_registry_v1"
}

func (r *QuayRedhatComQuayRegistryV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "QuayRegistry is the Schema for the quayregistries API.",
		MarkdownDescription: "QuayRegistry is the Schema for the quayregistries API.",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
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
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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
				Description:         "QuayRegistrySpec defines the desired state of QuayRegistry.",
				MarkdownDescription: "QuayRegistrySpec defines the desired state of QuayRegistry.",
				Attributes: map[string]schema.Attribute{
					"components": schema.ListNestedAttribute{
						Description:         "Components declare how the Operator should handle backing Quay services.",
						MarkdownDescription: "Components declare how the Operator should handle backing Quay services.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind is the unique name of this type of component.",
									MarkdownDescription: "Kind is the unique name of this type of component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"managed": schema.BoolAttribute{
									Description:         "Managed indicates whether or not the Operator is responsible for the lifecycle of this component. Default is true.",
									MarkdownDescription: "Managed indicates whether or not the Operator is responsible for the lifecycle of this component. Default is true.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"overrides": schema.SingleNestedAttribute{
									Description:         "Overrides holds information regarding component specific configurations.",
									MarkdownDescription: "Overrides holds information regarding component specific configurations.",
									Attributes: map[string]schema.Attribute{
										"affinity": schema.SingleNestedAttribute{
											Description:         "Affinity is a group of affinity scheduling rules.",
											MarkdownDescription: "Affinity is a group of affinity scheduling rules.",
											Attributes: map[string]schema.Attribute{
												"node_affinity": schema.SingleNestedAttribute{
													Description:         "Describes node affinity scheduling rules for the pod.",
													MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"preference": schema.SingleNestedAttribute{
																		Description:         "A node selector term, associated with the corresponding weight.",
																		MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"match_expressions": schema.ListNestedAttribute{
																				Description:         "A list of node selector requirements by node's labels.",
																				MarkdownDescription: "A list of node selector requirements by node's labels.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "The label key that the selector applies to.",
																							MarkdownDescription: "The label key that the selector applies to.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"operator": schema.StringAttribute{
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																			"match_fields": schema.ListNestedAttribute{
																				Description:         "A list of node selector requirements by node's fields.",
																				MarkdownDescription: "A list of node selector requirements by node's fields.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "The label key that the selector applies to.",
																							MarkdownDescription: "The label key that the selector applies to.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"operator": schema.StringAttribute{
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																		},
																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"weight": schema.Int64Attribute{
																		Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																		MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

														"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
															Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
															MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
															Attributes: map[string]schema.Attribute{
																"node_selector_terms": schema.ListNestedAttribute{
																	Description:         "Required. A list of node selector terms. The terms are ORed.",
																	MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"match_expressions": schema.ListNestedAttribute{
																				Description:         "A list of node selector requirements by node's labels.",
																				MarkdownDescription: "A list of node selector requirements by node's labels.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "The label key that the selector applies to.",
																							MarkdownDescription: "The label key that the selector applies to.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"operator": schema.StringAttribute{
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																			"match_fields": schema.ListNestedAttribute{
																				Description:         "A list of node selector requirements by node's fields.",
																				MarkdownDescription: "A list of node selector requirements by node's fields.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "The label key that the selector applies to.",
																							MarkdownDescription: "The label key that the selector applies to.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"operator": schema.StringAttribute{
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pod_affinity": schema.SingleNestedAttribute{
													Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
													MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"pod_affinity_term": schema.SingleNestedAttribute{
																		Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																		MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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

																			"namespaces": schema.ListAttribute{
																				Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"weight": schema.Int64Attribute{
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

														"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																	"namespace_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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

																	"namespaces": schema.ListAttribute{
																		Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"topology_key": schema.StringAttribute{
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

												"pod_anti_affinity": schema.SingleNestedAttribute{
													Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
													MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"pod_affinity_term": schema.SingleNestedAttribute{
																		Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																		MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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

																			"namespaces": schema.ListAttribute{
																				Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"weight": schema.Int64Attribute{
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

														"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																	"namespace_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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

																	"namespaces": schema.ListAttribute{
																		Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"topology_key": schema.StringAttribute{
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value_from": schema.SingleNestedAttribute{
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
														Attributes: map[string]schema.Attribute{
															"config_map_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"field_ref": schema.SingleNestedAttribute{
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																Attributes: map[string]schema.Attribute{
																	"api_version": schema.StringAttribute{
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"field_path": schema.StringAttribute{
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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

										"volume_size": schema.StringAttribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_bundle_secret": schema.StringAttribute{
						Description:         "ConfigBundleSecret is the name of the Kubernetes 'Secret' in the same namespace which contains the base Quay config and extra certs.",
						MarkdownDescription: "ConfigBundleSecret is the name of the Kubernetes 'Secret' in the same namespace which contains the base Quay config and extra certs.",
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

func (r *QuayRedhatComQuayRegistryV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *QuayRedhatComQuayRegistryV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_quay_redhat_com_quay_registry_v1")

	var model QuayRedhatComQuayRegistryV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("quay.redhat.com/v1")
	model.Kind = pointer.String("QuayRegistry")

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
		Resource(k8sSchema.GroupVersionResource{Group: "quay.redhat.com", Version: "v1", Resource: "quayregistries"}).
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

	var readResponse QuayRedhatComQuayRegistryV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *QuayRedhatComQuayRegistryV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_quay_redhat_com_quay_registry_v1")

	var data QuayRedhatComQuayRegistryV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "quay.redhat.com", Version: "v1", Resource: "quayregistries"}).
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

	var readResponse QuayRedhatComQuayRegistryV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *QuayRedhatComQuayRegistryV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_quay_redhat_com_quay_registry_v1")

	var model QuayRedhatComQuayRegistryV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("quay.redhat.com/v1")
	model.Kind = pointer.String("QuayRegistry")

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
		Resource(k8sSchema.GroupVersionResource{Group: "quay.redhat.com", Version: "v1", Resource: "quayregistries"}).
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

	var readResponse QuayRedhatComQuayRegistryV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *QuayRedhatComQuayRegistryV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_quay_redhat_com_quay_registry_v1")

	var data QuayRedhatComQuayRegistryV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "quay.redhat.com", Version: "v1", Resource: "quayregistries"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "quay.redhat.com", Version: "v1", Resource: "quayregistries"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
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

func (r *QuayRedhatComQuayRegistryV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
