/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

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
	_ resource.Resource                = &CephRookIoCephRbdmirrorV1Resource{}
	_ resource.ResourceWithConfigure   = &CephRookIoCephRbdmirrorV1Resource{}
	_ resource.ResourceWithImportState = &CephRookIoCephRbdmirrorV1Resource{}
)

func NewCephRookIoCephRbdmirrorV1Resource() resource.Resource {
	return &CephRookIoCephRbdmirrorV1Resource{}
}

type CephRookIoCephRbdmirrorV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CephRookIoCephRbdmirrorV1ResourceData struct {
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
		Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
		Count       *int64             `tfsdk:"count" json:"count,omitempty"`
		Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Peers       *struct {
			SecretNames *[]string `tfsdk:"secret_names" json:"secretNames,omitempty"`
		} `tfsdk:"peers" json:"peers,omitempty"`
		Placement *struct {
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
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]struct {
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				MatchLabelKeys     *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
				MaxSkew            *int64    `tfsdk:"max_skew" json:"maxSkew,omitempty"`
				MinDomains         *int64    `tfsdk:"min_domains" json:"minDomains,omitempty"`
				NodeAffinityPolicy *string   `tfsdk:"node_affinity_policy" json:"nodeAffinityPolicy,omitempty"`
				NodeTaintsPolicy   *string   `tfsdk:"node_taints_policy" json:"nodeTaintsPolicy,omitempty"`
				TopologyKey        *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				WhenUnsatisfiable  *string   `tfsdk:"when_unsatisfiable" json:"whenUnsatisfiable,omitempty"`
			} `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		Resources         *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephRbdmirrorV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_rbd_mirror_v1"
}

func (r *CephRookIoCephRbdmirrorV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephRBDMirror represents a Ceph RBD Mirror",
		MarkdownDescription: "CephRBDMirror represents a Ceph RBD Mirror",
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
				Description:         "RBDMirroringSpec represents the specification of an RBD mirror daemon",
				MarkdownDescription: "RBDMirroringSpec represents the specification of an RBD mirror daemon",
				Attributes: map[string]schema.Attribute{
					"annotations": schema.MapAttribute{
						Description:         "The annotations-related configuration to add/set on each Pod related object.",
						MarkdownDescription: "The annotations-related configuration to add/set on each Pod related object.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"count": schema.Int64Attribute{
						Description:         "Count represents the number of rbd mirror instance to run",
						MarkdownDescription: "Count represents the number of rbd mirror instance to run",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "The labels-related configuration to add/set on each Pod related object.",
						MarkdownDescription: "The labels-related configuration to add/set on each Pod related object.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"peers": schema.SingleNestedAttribute{
						Description:         "Peers represents the peers spec",
						MarkdownDescription: "Peers represents the peers spec",
						Attributes: map[string]schema.Attribute{
							"secret_names": schema.ListAttribute{
								Description:         "SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers",
								MarkdownDescription: "SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers",
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

					"placement": schema.SingleNestedAttribute{
						Description:         "The affinity to place the rgw pods (default is to place on any available node)",
						MarkdownDescription: "The affinity to place the rgw pods (default is to place on any available node)",
						Attributes: map[string]schema.Attribute{
							"node_affinity": schema.SingleNestedAttribute{
								Description:         "NodeAffinity is a group of node affinity scheduling rules",
								MarkdownDescription: "NodeAffinity is a group of node affinity scheduling rules",
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
								Description:         "PodAffinity is a group of inter pod affinity scheduling rules",
								MarkdownDescription: "PodAffinity is a group of inter pod affinity scheduling rules",
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
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
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
													Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
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
								Description:         "PodAntiAffinity is a group of inter pod anti affinity scheduling rules",
								MarkdownDescription: "PodAntiAffinity is a group of inter pod anti affinity scheduling rules",
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
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
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
													Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
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

							"tolerations": schema.ListNestedAttribute{
								Description:         "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>",
								MarkdownDescription: "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"topology_spread_constraints": schema.ListNestedAttribute{
								Description:         "TopologySpreadConstraint specifies how to spread matching pods among the given topology",
								MarkdownDescription: "TopologySpreadConstraint specifies how to spread matching pods among the given topology",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
											MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
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

										"match_label_keys": schema.ListAttribute{
											Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
											MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_skew": schema.Int64Attribute{
											Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
											MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"min_domains": schema.Int64Attribute{
											Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
											MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_affinity_policy": schema.StringAttribute{
											Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_taints_policy": schema.StringAttribute{
											Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"topology_key": schema.StringAttribute{
											Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
											MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"when_unsatisfiable": schema.StringAttribute{
											Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
											MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName sets priority class on the rbd mirror pods",
						MarkdownDescription: "PriorityClassName sets priority class on the rbd mirror pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "The resource requirements for the rbd mirror pods",
						MarkdownDescription: "The resource requirements for the rbd mirror pods",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
	}
}

func (r *CephRookIoCephRbdmirrorV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *CephRookIoCephRbdmirrorV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ceph_rook_io_ceph_rbd_mirror_v1")

	var model CephRookIoCephRbdmirrorV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephRBDMirror")

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
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephrbdmirrors"}).
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

	var readResponse CephRookIoCephRbdmirrorV1ResourceData
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

func (r *CephRookIoCephRbdmirrorV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_rbd_mirror_v1")

	var data CephRookIoCephRbdmirrorV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephrbdmirrors"}).
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

	var readResponse CephRookIoCephRbdmirrorV1ResourceData
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

func (r *CephRookIoCephRbdmirrorV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ceph_rook_io_ceph_rbd_mirror_v1")

	var model CephRookIoCephRbdmirrorV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephRBDMirror")

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
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephrbdmirrors"}).
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

	var readResponse CephRookIoCephRbdmirrorV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CephRookIoCephRbdmirrorV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ceph_rook_io_ceph_rbd_mirror_v1")

	var data CephRookIoCephRbdmirrorV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephrbdmirrors"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephrbdmirrors"}).
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

func (r *CephRookIoCephRbdmirrorV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
