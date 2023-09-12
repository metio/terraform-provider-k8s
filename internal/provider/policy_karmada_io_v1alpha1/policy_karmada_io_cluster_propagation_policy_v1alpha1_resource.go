/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policy_karmada_io_v1alpha1

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
	"time"
)

var (
	_ resource.Resource                = &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource{}
)

func NewPolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource() resource.Resource {
	return &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource{}
}

type PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData struct {
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
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Association        *bool     `tfsdk:"association" json:"association,omitempty"`
		ConflictResolution *string   `tfsdk:"conflict_resolution" json:"conflictResolution,omitempty"`
		DependentOverrides *[]string `tfsdk:"dependent_overrides" json:"dependentOverrides,omitempty"`
		Failover           *struct {
			Application *struct {
				DecisionConditions *struct {
					TolerationSeconds *int64 `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				} `tfsdk:"decision_conditions" json:"decisionConditions,omitempty"`
				GracePeriodSeconds *int64  `tfsdk:"grace_period_seconds" json:"gracePeriodSeconds,omitempty"`
				PurgeMode          *string `tfsdk:"purge_mode" json:"purgeMode,omitempty"`
			} `tfsdk:"application" json:"application,omitempty"`
		} `tfsdk:"failover" json:"failover,omitempty"`
		Placement *struct {
			ClusterAffinities *[]struct {
				AffinityName  *string   `tfsdk:"affinity_name" json:"affinityName,omitempty"`
				ClusterNames  *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
				Exclude       *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
				FieldSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				} `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			} `tfsdk:"cluster_affinities" json:"clusterAffinities,omitempty"`
			ClusterAffinity *struct {
				ClusterNames  *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
				Exclude       *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
				FieldSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				} `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			} `tfsdk:"cluster_affinity" json:"clusterAffinity,omitempty"`
			ClusterTolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"cluster_tolerations" json:"clusterTolerations,omitempty"`
			ReplicaScheduling *struct {
				ReplicaDivisionPreference *string `tfsdk:"replica_division_preference" json:"replicaDivisionPreference,omitempty"`
				ReplicaSchedulingType     *string `tfsdk:"replica_scheduling_type" json:"replicaSchedulingType,omitempty"`
				WeightPreference          *struct {
					DynamicWeight    *string `tfsdk:"dynamic_weight" json:"dynamicWeight,omitempty"`
					StaticWeightList *[]struct {
						TargetCluster *struct {
							ClusterNames  *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
							Exclude       *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
							FieldSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							} `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						} `tfsdk:"target_cluster" json:"targetCluster,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"static_weight_list" json:"staticWeightList,omitempty"`
				} `tfsdk:"weight_preference" json:"weightPreference,omitempty"`
			} `tfsdk:"replica_scheduling" json:"replicaScheduling,omitempty"`
			SpreadConstraints *[]struct {
				MaxGroups     *int64  `tfsdk:"max_groups" json:"maxGroups,omitempty"`
				MinGroups     *int64  `tfsdk:"min_groups" json:"minGroups,omitempty"`
				SpreadByField *string `tfsdk:"spread_by_field" json:"spreadByField,omitempty"`
				SpreadByLabel *string `tfsdk:"spread_by_label" json:"spreadByLabel,omitempty"`
			} `tfsdk:"spread_constraints" json:"spreadConstraints,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
		Preemption        *string `tfsdk:"preemption" json:"preemption,omitempty"`
		Priority          *int64  `tfsdk:"priority" json:"priority,omitempty"`
		PropagateDeps     *bool   `tfsdk:"propagate_deps" json:"propagateDeps,omitempty"`
		ResourceSelectors *[]struct {
			ApiVersion    *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind          *string `tfsdk:"kind" json:"kind,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"resource_selectors" json:"resourceSelectors,omitempty"`
		SchedulerName *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policy_karmada_io_cluster_propagation_policy_v1alpha1"
}

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterPropagationPolicy represents the cluster-wide policy that propagates a group of resources to one or more clusters. Different with PropagationPolicy that could only propagate resources in its own namespace, ClusterPropagationPolicy is able to propagate cluster level resources and resources in any namespace other than system reserved ones. System reserved namespaces are: karmada-system, karmada-cluster, karmada-es-*.",
		MarkdownDescription: "ClusterPropagationPolicy represents the cluster-wide policy that propagates a group of resources to one or more clusters. Different with PropagationPolicy that could only propagate resources in its own namespace, ClusterPropagationPolicy is able to propagate cluster level resources and resources in any namespace other than system reserved ones. System reserved namespaces are: karmada-system, karmada-cluster, karmada-es-*.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "Spec represents the desired behavior of ClusterPropagationPolicy.",
				MarkdownDescription: "Spec represents the desired behavior of ClusterPropagationPolicy.",
				Attributes: map[string]schema.Attribute{
					"association": schema.BoolAttribute{
						Description:         "Association tells if relevant resources should be selected automatically. e.g. a ConfigMap referred by a Deployment. default false. Deprecated: in favor of PropagateDeps.",
						MarkdownDescription: "Association tells if relevant resources should be selected automatically. e.g. a ConfigMap referred by a Deployment. default false. Deprecated: in favor of PropagateDeps.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"conflict_resolution": schema.StringAttribute{
						Description:         "ConflictResolution declares how potential conflict should be handled when a resource that is being propagated already exists in the target cluster.  It defaults to 'Abort' which means stop propagating to avoid unexpected overwrites. The 'Overwrite' might be useful when migrating legacy cluster resources to Karmada, in which case conflict is predictable and can be instructed to Karmada take over the resource by overwriting.",
						MarkdownDescription: "ConflictResolution declares how potential conflict should be handled when a resource that is being propagated already exists in the target cluster.  It defaults to 'Abort' which means stop propagating to avoid unexpected overwrites. The 'Overwrite' might be useful when migrating legacy cluster resources to Karmada, in which case conflict is predictable and can be instructed to Karmada take over the resource by overwriting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Abort", "Overwrite"),
						},
					},

					"dependent_overrides": schema.ListAttribute{
						Description:         "DependentOverrides represents the list of overrides(OverridePolicy) which must present before the current PropagationPolicy takes effect.  It used to explicitly specify overrides which current PropagationPolicy rely on. A typical scenario is the users create OverridePolicy(ies) and resources at the same time, they want to ensure the new-created policies would be adopted.  Note: For the overrides, OverridePolicy(ies) in current namespace and ClusterOverridePolicy(ies), which not present in this list will still be applied if they matches the resources.",
						MarkdownDescription: "DependentOverrides represents the list of overrides(OverridePolicy) which must present before the current PropagationPolicy takes effect.  It used to explicitly specify overrides which current PropagationPolicy rely on. A typical scenario is the users create OverridePolicy(ies) and resources at the same time, they want to ensure the new-created policies would be adopted.  Note: For the overrides, OverridePolicy(ies) in current namespace and ClusterOverridePolicy(ies), which not present in this list will still be applied if they matches the resources.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failover": schema.SingleNestedAttribute{
						Description:         "Failover indicates how Karmada migrates applications in case of failures. If this value is nil, failover is disabled.",
						MarkdownDescription: "Failover indicates how Karmada migrates applications in case of failures. If this value is nil, failover is disabled.",
						Attributes: map[string]schema.Attribute{
							"application": schema.SingleNestedAttribute{
								Description:         "Application indicates failover behaviors in case of application failure. If this value is nil, failover is disabled. If set, the PropagateDeps should be true so that the dependencies could be migrated along with the application.",
								MarkdownDescription: "Application indicates failover behaviors in case of application failure. If this value is nil, failover is disabled. If set, the PropagateDeps should be true so that the dependencies could be migrated along with the application.",
								Attributes: map[string]schema.Attribute{
									"decision_conditions": schema.SingleNestedAttribute{
										Description:         "DecisionConditions indicates the decision conditions of performing the failover process. Only when all conditions are met can the failover process be performed. Currently, DecisionConditions includes several conditions: - TolerationSeconds (optional)",
										MarkdownDescription: "DecisionConditions indicates the decision conditions of performing the failover process. Only when all conditions are met can the failover process be performed. Currently, DecisionConditions includes several conditions: - TolerationSeconds (optional)",
										Attributes: map[string]schema.Attribute{
											"toleration_seconds": schema.Int64Attribute{
												Description:         "TolerationSeconds represents the period of time Karmada should wait after reaching the desired state before performing failover process. If not specified, Karmada will immediately perform failover process. Defaults to 300s.",
												MarkdownDescription: "TolerationSeconds represents the period of time Karmada should wait after reaching the desired state before performing failover process. If not specified, Karmada will immediately perform failover process. Defaults to 300s.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"grace_period_seconds": schema.Int64Attribute{
										Description:         "GracePeriodSeconds is the maximum waiting duration in seconds before application on the migrated cluster should be deleted. Required only when PurgeMode is 'Graciously' and defaults to 600s. If the application on the new cluster cannot reach a Healthy state, Karmada will delete the application after GracePeriodSeconds is reached. Value must be positive integer.",
										MarkdownDescription: "GracePeriodSeconds is the maximum waiting duration in seconds before application on the migrated cluster should be deleted. Required only when PurgeMode is 'Graciously' and defaults to 600s. If the application on the new cluster cannot reach a Healthy state, Karmada will delete the application after GracePeriodSeconds is reached. Value must be positive integer.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"purge_mode": schema.StringAttribute{
										Description:         "PurgeMode represents how to deal with the legacy applications on the cluster from which the application is migrated. Valid options are 'Immediately', 'Graciously' and 'Never'. Defaults to 'Graciously'.",
										MarkdownDescription: "PurgeMode represents how to deal with the legacy applications on the cluster from which the application is migrated. Valid options are 'Immediately', 'Graciously' and 'Never'. Defaults to 'Graciously'.",
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

					"placement": schema.SingleNestedAttribute{
						Description:         "Placement represents the rule for select clusters to propagate resources.",
						MarkdownDescription: "Placement represents the rule for select clusters to propagate resources.",
						Attributes: map[string]schema.Attribute{
							"cluster_affinities": schema.ListNestedAttribute{
								Description:         "ClusterAffinities represents scheduling restrictions to multiple cluster groups that indicated by ClusterAffinityTerm.  The scheduler will evaluate these groups one by one in the order they appear in the spec, the group that does not satisfy scheduling restrictions will be ignored which means all clusters in this group will not be selected unless it also belongs to the next group(a cluster could belong to multiple groups).  If none of the groups satisfy the scheduling restrictions, then scheduling fails, which means no cluster will be selected.  Note: 1. ClusterAffinities can not co-exist with ClusterAffinity. 2. If both ClusterAffinity and ClusterAffinities are not set, any cluster can be scheduling candidates.  Potential use case 1: The private clusters in the local data center could be the main group, and the managed clusters provided by cluster providers could be the secondary group. So that the Karmada scheduler would prefer to schedule workloads to the main group and the second group will only be considered in case of the main group does not satisfy restrictions(like, lack of resources).  Potential use case 2: For the disaster recovery scenario, the clusters could be organized to primary and backup groups, the workloads would be scheduled to primary clusters firstly, and when primary cluster fails(like data center power off), Karmada scheduler could migrate workloads to the backup clusters.",
								MarkdownDescription: "ClusterAffinities represents scheduling restrictions to multiple cluster groups that indicated by ClusterAffinityTerm.  The scheduler will evaluate these groups one by one in the order they appear in the spec, the group that does not satisfy scheduling restrictions will be ignored which means all clusters in this group will not be selected unless it also belongs to the next group(a cluster could belong to multiple groups).  If none of the groups satisfy the scheduling restrictions, then scheduling fails, which means no cluster will be selected.  Note: 1. ClusterAffinities can not co-exist with ClusterAffinity. 2. If both ClusterAffinity and ClusterAffinities are not set, any cluster can be scheduling candidates.  Potential use case 1: The private clusters in the local data center could be the main group, and the managed clusters provided by cluster providers could be the secondary group. So that the Karmada scheduler would prefer to schedule workloads to the main group and the second group will only be considered in case of the main group does not satisfy restrictions(like, lack of resources).  Potential use case 2: For the disaster recovery scenario, the clusters could be organized to primary and backup groups, the workloads would be scheduled to primary clusters firstly, and when primary cluster fails(like data center power off), Karmada scheduler could migrate workloads to the backup clusters.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"affinity_name": schema.StringAttribute{
											Description:         "AffinityName is the name of the cluster group.",
											MarkdownDescription: "AffinityName is the name of the cluster group.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(32),
											},
										},

										"cluster_names": schema.ListAttribute{
											Description:         "ClusterNames is the list of clusters to be selected.",
											MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": schema.ListAttribute{
											Description:         "ExcludedClusters is the list of clusters to be ignored.",
											MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"field_selector": schema.SingleNestedAttribute{
											Description:         "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
											MarkdownDescription: "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of field selector requirements.",
													MarkdownDescription: "A list of field selector requirements.",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
											MarkdownDescription: "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_affinity": schema.SingleNestedAttribute{
								Description:         "ClusterAffinity represents scheduling restrictions to a certain set of clusters. Note: 1. ClusterAffinity can not co-exist with ClusterAffinities. 2. If both ClusterAffinity and ClusterAffinities are not set, any cluster can be scheduling candidates.",
								MarkdownDescription: "ClusterAffinity represents scheduling restrictions to a certain set of clusters. Note: 1. ClusterAffinity can not co-exist with ClusterAffinities. 2. If both ClusterAffinity and ClusterAffinities are not set, any cluster can be scheduling candidates.",
								Attributes: map[string]schema.Attribute{
									"cluster_names": schema.ListAttribute{
										Description:         "ClusterNames is the list of clusters to be selected.",
										MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"exclude": schema.ListAttribute{
										Description:         "ExcludedClusters is the list of clusters to be ignored.",
										MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_selector": schema.SingleNestedAttribute{
										Description:         "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
										MarkdownDescription: "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "A list of field selector requirements.",
												MarkdownDescription: "A list of field selector requirements.",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"label_selector": schema.SingleNestedAttribute{
										Description:         "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
										MarkdownDescription: "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_tolerations": schema.ListNestedAttribute{
								Description:         "ClusterTolerations represents the tolerations.",
								MarkdownDescription: "ClusterTolerations represents the tolerations.",
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

							"replica_scheduling": schema.SingleNestedAttribute{
								Description:         "ReplicaScheduling represents the scheduling policy on dealing with the number of replicas when propagating resources that have replicas in spec (e.g. deployments, statefulsets) to member clusters.",
								MarkdownDescription: "ReplicaScheduling represents the scheduling policy on dealing with the number of replicas when propagating resources that have replicas in spec (e.g. deployments, statefulsets) to member clusters.",
								Attributes: map[string]schema.Attribute{
									"replica_division_preference": schema.StringAttribute{
										Description:         "ReplicaDivisionPreference determines how the replicas is divided when ReplicaSchedulingType is 'Divided'. Valid options are Aggregated and Weighted. 'Aggregated' divides replicas into clusters as few as possible, while respecting clusters' resource availabilities during the division. 'Weighted' divides replicas by weight according to WeightPreference.",
										MarkdownDescription: "ReplicaDivisionPreference determines how the replicas is divided when ReplicaSchedulingType is 'Divided'. Valid options are Aggregated and Weighted. 'Aggregated' divides replicas into clusters as few as possible, while respecting clusters' resource availabilities during the division. 'Weighted' divides replicas by weight according to WeightPreference.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Aggregated", "Weighted"),
										},
									},

									"replica_scheduling_type": schema.StringAttribute{
										Description:         "ReplicaSchedulingType determines how the replicas is scheduled when karmada propagating a resource. Valid options are Duplicated and Divided. 'Duplicated' duplicates the same replicas to each candidate member cluster from resource. 'Divided' divides replicas into parts according to number of valid candidate member clusters, and exact replicas for each cluster are determined by ReplicaDivisionPreference.",
										MarkdownDescription: "ReplicaSchedulingType determines how the replicas is scheduled when karmada propagating a resource. Valid options are Duplicated and Divided. 'Duplicated' duplicates the same replicas to each candidate member cluster from resource. 'Divided' divides replicas into parts according to number of valid candidate member clusters, and exact replicas for each cluster are determined by ReplicaDivisionPreference.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Duplicated", "Divided"),
										},
									},

									"weight_preference": schema.SingleNestedAttribute{
										Description:         "WeightPreference describes weight for each cluster or for each group of cluster If ReplicaDivisionPreference is set to 'Weighted', and WeightPreference is not set, scheduler will weight all clusters the same.",
										MarkdownDescription: "WeightPreference describes weight for each cluster or for each group of cluster If ReplicaDivisionPreference is set to 'Weighted', and WeightPreference is not set, scheduler will weight all clusters the same.",
										Attributes: map[string]schema.Attribute{
											"dynamic_weight": schema.StringAttribute{
												Description:         "DynamicWeight specifies the factor to generates dynamic weight list. If specified, StaticWeightList will be ignored.",
												MarkdownDescription: "DynamicWeight specifies the factor to generates dynamic weight list. If specified, StaticWeightList will be ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("AvailableReplicas"),
												},
											},

											"static_weight_list": schema.ListNestedAttribute{
												Description:         "StaticWeightList defines the static cluster weight.",
												MarkdownDescription: "StaticWeightList defines the static cluster weight.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"target_cluster": schema.SingleNestedAttribute{
															Description:         "TargetCluster describes the filter to select clusters.",
															MarkdownDescription: "TargetCluster describes the filter to select clusters.",
															Attributes: map[string]schema.Attribute{
																"cluster_names": schema.ListAttribute{
																	Description:         "ClusterNames is the list of clusters to be selected.",
																	MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"exclude": schema.ListAttribute{
																	Description:         "ExcludedClusters is the list of clusters to be ignored.",
																	MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_selector": schema.SingleNestedAttribute{
																	Description:         "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
																	MarkdownDescription: "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "A list of field selector requirements.",
																			MarkdownDescription: "A list of field selector requirements.",
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
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"label_selector": schema.SingleNestedAttribute{
																	Description:         "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
																	MarkdownDescription: "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
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

														"weight": schema.Int64Attribute{
															Description:         "Weight expressing the preference to the cluster(s) specified by 'TargetCluster'.",
															MarkdownDescription: "Weight expressing the preference to the cluster(s) specified by 'TargetCluster'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spread_constraints": schema.ListNestedAttribute{
								Description:         "SpreadConstraints represents a list of the scheduling constraints.",
								MarkdownDescription: "SpreadConstraints represents a list of the scheduling constraints.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"max_groups": schema.Int64Attribute{
											Description:         "MaxGroups restricts the maximum number of cluster groups to be selected.",
											MarkdownDescription: "MaxGroups restricts the maximum number of cluster groups to be selected.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_groups": schema.Int64Attribute{
											Description:         "MinGroups restricts the minimum number of cluster groups to be selected. Defaults to 1.",
											MarkdownDescription: "MinGroups restricts the minimum number of cluster groups to be selected. Defaults to 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"spread_by_field": schema.StringAttribute{
											Description:         "SpreadByField represents the fields on Karmada cluster API used for dynamically grouping member clusters into different groups. Resources will be spread among different cluster groups. Available fields for spreading are: cluster, region, zone, and provider. SpreadByField should not co-exist with SpreadByLabel. If both SpreadByField and SpreadByLabel are empty, SpreadByField will be set to 'cluster' by system.",
											MarkdownDescription: "SpreadByField represents the fields on Karmada cluster API used for dynamically grouping member clusters into different groups. Resources will be spread among different cluster groups. Available fields for spreading are: cluster, region, zone, and provider. SpreadByField should not co-exist with SpreadByLabel. If both SpreadByField and SpreadByLabel are empty, SpreadByField will be set to 'cluster' by system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("cluster", "region", "zone", "provider"),
											},
										},

										"spread_by_label": schema.StringAttribute{
											Description:         "SpreadByLabel represents the label key used for grouping member clusters into different groups. Resources will be spread among different cluster groups. SpreadByLabel should not co-exist with SpreadByField.",
											MarkdownDescription: "SpreadByLabel represents the label key used for grouping member clusters into different groups. Resources will be spread among different cluster groups. SpreadByLabel should not co-exist with SpreadByField.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"preemption": schema.StringAttribute{
						Description:         "Preemption declares the behaviors for preempting. Valid options are 'Always' and 'Never'.",
						MarkdownDescription: "Preemption declares the behaviors for preempting. Valid options are 'Always' and 'Never'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Always", "Never"),
						},
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority indicates the importance of a policy(PropagationPolicy or ClusterPropagationPolicy). A policy will be applied for the matched resource templates if there is no other policies with higher priority at the point of the resource template be processed. Once a resource template has been claimed by a policy, by default it will not be preempted by following policies even with a higher priority. See Preemption for more details.  In case of two policies have the same priority, the one with a more precise matching rules in ResourceSelectors wins: - matching by name(resourceSelector.name) has higher priority than by selector(resourceSelector.labelSelector) - matching by selector(resourceSelector.labelSelector) has higher priority than by APIVersion(resourceSelector.apiVersion) and Kind(resourceSelector.kind). If there is still no winner at this point, the one with the lower alphabetic order wins, e.g. policy 'bar' has higher priority than 'foo'.  The higher the value, the higher the priority. Defaults to zero.",
						MarkdownDescription: "Priority indicates the importance of a policy(PropagationPolicy or ClusterPropagationPolicy). A policy will be applied for the matched resource templates if there is no other policies with higher priority at the point of the resource template be processed. Once a resource template has been claimed by a policy, by default it will not be preempted by following policies even with a higher priority. See Preemption for more details.  In case of two policies have the same priority, the one with a more precise matching rules in ResourceSelectors wins: - matching by name(resourceSelector.name) has higher priority than by selector(resourceSelector.labelSelector) - matching by selector(resourceSelector.labelSelector) has higher priority than by APIVersion(resourceSelector.apiVersion) and Kind(resourceSelector.kind). If there is still no winner at this point, the one with the lower alphabetic order wins, e.g. policy 'bar' has higher priority than 'foo'.  The higher the value, the higher the priority. Defaults to zero.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"propagate_deps": schema.BoolAttribute{
						Description:         "PropagateDeps tells if relevant resources should be propagated automatically. Take 'Deployment' which referencing 'ConfigMap' and 'Secret' as an example, when 'propagateDeps' is 'true', the referencing resources could be omitted(for saving config effort) from 'resourceSelectors' as they will be propagated along with the Deployment. In addition to the propagating process, the referencing resources will be migrated along with the Deployment in the fail-over scenario.  Defaults to false.",
						MarkdownDescription: "PropagateDeps tells if relevant resources should be propagated automatically. Take 'Deployment' which referencing 'ConfigMap' and 'Secret' as an example, when 'propagateDeps' is 'true', the referencing resources could be omitted(for saving config effort) from 'resourceSelectors' as they will be propagated along with the Deployment. In addition to the propagating process, the referencing resources will be migrated along with the Deployment in the fail-over scenario.  Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_selectors": schema.ListNestedAttribute{
						Description:         "ResourceSelectors used to select resources. Nil or empty selector is not allowed and doesn't mean match all kinds of resources for security concerns that sensitive resources(like Secret) might be accidentally propagated.",
						MarkdownDescription: "ResourceSelectors used to select resources. Nil or empty selector is not allowed and doesn't mean match all kinds of resources for security concerns that sensitive resources(like Secret) might be accidentally propagated.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion represents the API version of the target resources.",
									MarkdownDescription: "APIVersion represents the API version of the target resources.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind represents the Kind of the target resources.",
									MarkdownDescription: "Kind represents the Kind of the target resources.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"label_selector": schema.SingleNestedAttribute{
									Description:         "A label query over a set of resources. If name is not empty, labelSelector will be ignored.",
									MarkdownDescription: "A label query over a set of resources. If name is not empty, labelSelector will be ignored.",
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

								"name": schema.StringAttribute{
									Description:         "Name of the target resource. Default is empty, which means selecting all resources.",
									MarkdownDescription: "Name of the target resource. Default is empty, which means selecting all resources.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the target resource. Default is empty, which means inherit from the parent object scope.",
									MarkdownDescription: "Namespace of the target resource. Default is empty, which means inherit from the parent object scope.",
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

					"scheduler_name": schema.StringAttribute{
						Description:         "SchedulerName represents which scheduler to proceed the scheduling. If specified, the policy will be dispatched by specified scheduler. If not specified, the policy will be dispatched by default scheduler.",
						MarkdownDescription: "SchedulerName represents which scheduler to proceed the scheduling. If specified, the policy will be dispatched by specified scheduler. If not specified, the policy will be dispatched by default scheduler.",
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
	}
}

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1")

	var model PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("policy.karmada.io/v1alpha1")
	model.Kind = pointer.String("ClusterPropagationPolicy")

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
		Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "clusterpropagationpolicies"}).
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

	var readResponse PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
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

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1")

	var data PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "clusterpropagationpolicies"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
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

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1")

	var model PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("policy.karmada.io/v1alpha1")
	model.Kind = pointer.String("ClusterPropagationPolicy")

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
		Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "clusterpropagationpolicies"}).
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

	var readResponse PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1")

	var data PolicyKarmadaIoClusterPropagationPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "clusterpropagationpolicies"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "clusterpropagationpolicies"}).
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

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
