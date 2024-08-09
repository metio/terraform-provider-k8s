/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package work_karmada_io_v1alpha2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &WorkKarmadaIoClusterResourceBindingV1Alpha2Manifest{}
)

func NewWorkKarmadaIoClusterResourceBindingV1Alpha2Manifest() datasource.DataSource {
	return &WorkKarmadaIoClusterResourceBindingV1Alpha2Manifest{}
}

type WorkKarmadaIoClusterResourceBindingV1Alpha2Manifest struct{}

type WorkKarmadaIoClusterResourceBindingV1Alpha2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Clusters *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Replicas *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"clusters" json:"clusters,omitempty"`
		ConflictResolution *string `tfsdk:"conflict_resolution" json:"conflictResolution,omitempty"`
		Failover           *struct {
			Application *struct {
				DecisionConditions *struct {
					TolerationSeconds *int64 `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				} `tfsdk:"decision_conditions" json:"decisionConditions,omitempty"`
				GracePeriodSeconds *int64  `tfsdk:"grace_period_seconds" json:"gracePeriodSeconds,omitempty"`
				PurgeMode          *string `tfsdk:"purge_mode" json:"purgeMode,omitempty"`
			} `tfsdk:"application" json:"application,omitempty"`
		} `tfsdk:"failover" json:"failover,omitempty"`
		GracefulEvictionTasks *[]struct {
			CreationTimestamp  *string `tfsdk:"creation_timestamp" json:"creationTimestamp,omitempty"`
			FromCluster        *string `tfsdk:"from_cluster" json:"fromCluster,omitempty"`
			GracePeriodSeconds *int64  `tfsdk:"grace_period_seconds" json:"gracePeriodSeconds,omitempty"`
			Message            *string `tfsdk:"message" json:"message,omitempty"`
			Producer           *string `tfsdk:"producer" json:"producer,omitempty"`
			Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
			Replicas           *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
			SuppressDeletion   *bool   `tfsdk:"suppress_deletion" json:"suppressDeletion,omitempty"`
		} `tfsdk:"graceful_eviction_tasks" json:"gracefulEvictionTasks,omitempty"`
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
		PropagateDeps       *bool `tfsdk:"propagate_deps" json:"propagateDeps,omitempty"`
		ReplicaRequirements *struct {
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			NodeClaim *struct {
				HardNodeAffinity *struct {
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
				} `tfsdk:"hard_node_affinity" json:"hardNodeAffinity,omitempty"`
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"node_claim" json:"nodeClaim,omitempty"`
			PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ResourceRequest   *map[string]string `tfsdk:"resource_request" json:"resourceRequest,omitempty"`
		} `tfsdk:"replica_requirements" json:"replicaRequirements,omitempty"`
		Replicas   *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		RequiredBy *[]struct {
			Clusters *[]struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Replicas *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"required_by" json:"requiredBy,omitempty"`
		RescheduleTriggeredAt *string `tfsdk:"reschedule_triggered_at" json:"rescheduleTriggeredAt,omitempty"`
		Resource              *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"resource" json:"resource,omitempty"`
		SchedulerName *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		Suspension    *struct {
			Dispatching           *bool `tfsdk:"dispatching" json:"dispatching,omitempty"`
			DispatchingOnClusters *struct {
				ClusterNames *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
			} `tfsdk:"dispatching_on_clusters" json:"dispatchingOnClusters,omitempty"`
		} `tfsdk:"suspension" json:"suspension,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *WorkKarmadaIoClusterResourceBindingV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_work_karmada_io_cluster_resource_binding_v1alpha2_manifest"
}

func (r *WorkKarmadaIoClusterResourceBindingV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterResourceBinding represents a binding of a kubernetes resource with a ClusterPropagationPolicy.",
		MarkdownDescription: "ClusterResourceBinding represents a binding of a kubernetes resource with a ClusterPropagationPolicy.",
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
				Description:         "Spec represents the desired behavior.",
				MarkdownDescription: "Spec represents the desired behavior.",
				Attributes: map[string]schema.Attribute{
					"clusters": schema.ListNestedAttribute{
						Description:         "Clusters represents target member clusters where the resource to be deployed.",
						MarkdownDescription: "Clusters represents target member clusters where the resource to be deployed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of target cluster.",
									MarkdownDescription: "Name of target cluster.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Replicas in target cluster",
									MarkdownDescription: "Replicas in target cluster",
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

					"conflict_resolution": schema.StringAttribute{
						Description:         "ConflictResolution declares how potential conflict should be handled whena resource that is being propagated already exists in the target cluster.It defaults to 'Abort' which means stop propagating to avoid unexpectedoverwrites. The 'Overwrite' might be useful when migrating legacy clusterresources to Karmada, in which case conflict is predictable and can beinstructed to Karmada take over the resource by overwriting.",
						MarkdownDescription: "ConflictResolution declares how potential conflict should be handled whena resource that is being propagated already exists in the target cluster.It defaults to 'Abort' which means stop propagating to avoid unexpectedoverwrites. The 'Overwrite' might be useful when migrating legacy clusterresources to Karmada, in which case conflict is predictable and can beinstructed to Karmada take over the resource by overwriting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Abort", "Overwrite"),
						},
					},

					"failover": schema.SingleNestedAttribute{
						Description:         "Failover indicates how Karmada migrates applications in case of failures.It inherits directly from the associated PropagationPolicy(or ClusterPropagationPolicy).",
						MarkdownDescription: "Failover indicates how Karmada migrates applications in case of failures.It inherits directly from the associated PropagationPolicy(or ClusterPropagationPolicy).",
						Attributes: map[string]schema.Attribute{
							"application": schema.SingleNestedAttribute{
								Description:         "Application indicates failover behaviors in case of application failure.If this value is nil, failover is disabled.If set, the PropagateDeps should be true so that the dependencies couldbe migrated along with the application.",
								MarkdownDescription: "Application indicates failover behaviors in case of application failure.If this value is nil, failover is disabled.If set, the PropagateDeps should be true so that the dependencies couldbe migrated along with the application.",
								Attributes: map[string]schema.Attribute{
									"decision_conditions": schema.SingleNestedAttribute{
										Description:         "DecisionConditions indicates the decision conditions of performing the failover process.Only when all conditions are met can the failover process be performed.Currently, DecisionConditions includes several conditions:- TolerationSeconds (optional)",
										MarkdownDescription: "DecisionConditions indicates the decision conditions of performing the failover process.Only when all conditions are met can the failover process be performed.Currently, DecisionConditions includes several conditions:- TolerationSeconds (optional)",
										Attributes: map[string]schema.Attribute{
											"toleration_seconds": schema.Int64Attribute{
												Description:         "TolerationSeconds represents the period of time Karmada should waitafter reaching the desired state before performing failover process.If not specified, Karmada will immediately perform failover process.Defaults to 300s.",
												MarkdownDescription: "TolerationSeconds represents the period of time Karmada should waitafter reaching the desired state before performing failover process.If not specified, Karmada will immediately perform failover process.Defaults to 300s.",
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
										Description:         "GracePeriodSeconds is the maximum waiting duration in seconds beforeapplication on the migrated cluster should be deleted.Required only when PurgeMode is 'Graciously' and defaults to 600s.If the application on the new cluster cannot reach a Healthy state,Karmada will delete the application after GracePeriodSeconds is reached.Value must be positive integer.",
										MarkdownDescription: "GracePeriodSeconds is the maximum waiting duration in seconds beforeapplication on the migrated cluster should be deleted.Required only when PurgeMode is 'Graciously' and defaults to 600s.If the application on the new cluster cannot reach a Healthy state,Karmada will delete the application after GracePeriodSeconds is reached.Value must be positive integer.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"purge_mode": schema.StringAttribute{
										Description:         "PurgeMode represents how to deal with the legacy applications on thecluster from which the application is migrated.Valid options are 'Immediately', 'Graciously' and 'Never'.Defaults to 'Graciously'.",
										MarkdownDescription: "PurgeMode represents how to deal with the legacy applications on thecluster from which the application is migrated.Valid options are 'Immediately', 'Graciously' and 'Never'.Defaults to 'Graciously'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Immediately", "Graciously", "Never"),
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

					"graceful_eviction_tasks": schema.ListNestedAttribute{
						Description:         "GracefulEvictionTasks holds the eviction tasks that are expected to performthe eviction in a graceful way.The intended workflow is:1. Once the controller(such as 'taint-manager') decided to evict the resource that   is referenced by current ResourceBinding or ClusterResourceBinding from a target   cluster, it removes(or scale down the replicas) the target from Clusters(.spec.Clusters)   and builds a graceful eviction task.2. The scheduler may perform a re-scheduler and probably select a substitute cluster   to take over the evicting workload(resource).3. The graceful eviction controller takes care of the graceful eviction tasks and   performs the final removal after the workload(resource) is available on the substitute   cluster or exceed the grace termination period(defaults to 10 minutes).",
						MarkdownDescription: "GracefulEvictionTasks holds the eviction tasks that are expected to performthe eviction in a graceful way.The intended workflow is:1. Once the controller(such as 'taint-manager') decided to evict the resource that   is referenced by current ResourceBinding or ClusterResourceBinding from a target   cluster, it removes(or scale down the replicas) the target from Clusters(.spec.Clusters)   and builds a graceful eviction task.2. The scheduler may perform a re-scheduler and probably select a substitute cluster   to take over the evicting workload(resource).3. The graceful eviction controller takes care of the graceful eviction tasks and   performs the final removal after the workload(resource) is available on the substitute   cluster or exceed the grace termination period(defaults to 10 minutes).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"creation_timestamp": schema.StringAttribute{
									Description:         "CreationTimestamp is a timestamp representing the server time when this object wascreated.Clients should not set this value to avoid the time inconsistency issue.It is represented in RFC3339 form(like '2021-04-25T10:02:10Z') and is in UTC.Populated by the system. Read-only.",
									MarkdownDescription: "CreationTimestamp is a timestamp representing the server time when this object wascreated.Clients should not set this value to avoid the time inconsistency issue.It is represented in RFC3339 form(like '2021-04-25T10:02:10Z') and is in UTC.Populated by the system. Read-only.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"from_cluster": schema.StringAttribute{
									Description:         "FromCluster which cluster the eviction perform from.",
									MarkdownDescription: "FromCluster which cluster the eviction perform from.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"grace_period_seconds": schema.Int64Attribute{
									Description:         "GracePeriodSeconds is the maximum waiting duration in seconds before the itemshould be deleted. If the application on the new cluster cannot reach a Healthy state,Karmada will delete the item after GracePeriodSeconds is reached.Value must be positive integer.It can not co-exist with SuppressDeletion.",
									MarkdownDescription: "GracePeriodSeconds is the maximum waiting duration in seconds before the itemshould be deleted. If the application on the new cluster cannot reach a Healthy state,Karmada will delete the item after GracePeriodSeconds is reached.Value must be positive integer.It can not co-exist with SuppressDeletion.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"message": schema.StringAttribute{
									Description:         "Message is a human-readable message indicating details about the eviction.This may be an empty string.",
									MarkdownDescription: "Message is a human-readable message indicating details about the eviction.This may be an empty string.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(1024),
									},
								},

								"producer": schema.StringAttribute{
									Description:         "Producer indicates the controller who triggered the eviction.",
									MarkdownDescription: "Producer indicates the controller who triggered the eviction.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"reason": schema.StringAttribute{
									Description:         "Reason contains a programmatic identifier indicating the reason for the eviction.Producers may define expected values and meanings for this field,and whether the values are considered a guaranteed API.The value should be a CamelCase string.This field may not be empty.",
									MarkdownDescription: "Reason contains a programmatic identifier indicating the reason for the eviction.Producers may define expected values and meanings for this field,and whether the values are considered a guaranteed API.The value should be a CamelCase string.This field may not be empty.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(32),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`), ""),
									},
								},

								"replicas": schema.Int64Attribute{
									Description:         "Replicas indicates the number of replicas should be evicted.Should be ignored for resource type that doesn't have replica.",
									MarkdownDescription: "Replicas indicates the number of replicas should be evicted.Should be ignored for resource type that doesn't have replica.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"suppress_deletion": schema.BoolAttribute{
									Description:         "SuppressDeletion represents the grace period will be persistent untilthe tools or human intervention stops it.It can not co-exist with GracePeriodSeconds.",
									MarkdownDescription: "SuppressDeletion represents the grace period will be persistent untilthe tools or human intervention stops it.It can not co-exist with GracePeriodSeconds.",
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

					"placement": schema.SingleNestedAttribute{
						Description:         "Placement represents the rule for select clusters to propagate resources.",
						MarkdownDescription: "Placement represents the rule for select clusters to propagate resources.",
						Attributes: map[string]schema.Attribute{
							"cluster_affinities": schema.ListNestedAttribute{
								Description:         "ClusterAffinities represents scheduling restrictions to multiple clustergroups that indicated by ClusterAffinityTerm.The scheduler will evaluate these groups one by one in the order theyappear in the spec, the group that does not satisfy scheduling restrictionswill be ignored which means all clusters in this group will not be selectedunless it also belongs to the next group(a cluster could belong to multiplegroups).If none of the groups satisfy the scheduling restrictions, then schedulingfails, which means no cluster will be selected.Note:  1. ClusterAffinities can not co-exist with ClusterAffinity.  2. If both ClusterAffinity and ClusterAffinities are not set, any cluster     can be scheduling candidates.Potential use case 1:The private clusters in the local data center could be the main group, andthe managed clusters provided by cluster providers could be the secondarygroup. So that the Karmada scheduler would prefer to schedule workloadsto the main group and the second group will only be considered in case ofthe main group does not satisfy restrictions(like, lack of resources).Potential use case 2:For the disaster recovery scenario, the clusters could be organized toprimary and backup groups, the workloads would be scheduled to primaryclusters firstly, and when primary cluster fails(like data center power off),Karmada scheduler could migrate workloads to the backup clusters.",
								MarkdownDescription: "ClusterAffinities represents scheduling restrictions to multiple clustergroups that indicated by ClusterAffinityTerm.The scheduler will evaluate these groups one by one in the order theyappear in the spec, the group that does not satisfy scheduling restrictionswill be ignored which means all clusters in this group will not be selectedunless it also belongs to the next group(a cluster could belong to multiplegroups).If none of the groups satisfy the scheduling restrictions, then schedulingfails, which means no cluster will be selected.Note:  1. ClusterAffinities can not co-exist with ClusterAffinity.  2. If both ClusterAffinity and ClusterAffinities are not set, any cluster     can be scheduling candidates.Potential use case 1:The private clusters in the local data center could be the main group, andthe managed clusters provided by cluster providers could be the secondarygroup. So that the Karmada scheduler would prefer to schedule workloadsto the main group and the second group will only be considered in case ofthe main group does not satisfy restrictions(like, lack of resources).Potential use case 2:For the disaster recovery scenario, the clusters could be organized toprimary and backup groups, the workloads would be scheduled to primaryclusters firstly, and when primary cluster fails(like data center power off),Karmada scheduler could migrate workloads to the backup clusters.",
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
											Description:         "FieldSelector is a filter to select member clusters by fields.The key(field) of the match expression should be 'provider', 'region', or 'zone',and the operator of the match expression should be 'In' or 'NotIn'.If non-nil and non-empty, only the clusters match this filter will be selected.",
											MarkdownDescription: "FieldSelector is a filter to select member clusters by fields.The key(field) of the match expression should be 'provider', 'region', or 'zone',and the operator of the match expression should be 'In' or 'NotIn'.If non-nil and non-empty, only the clusters match this filter will be selected.",
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
																Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
											Description:         "LabelSelector is a filter to select member clusters by labels.If non-nil and non-empty, only the clusters match this filter will be selected.",
											MarkdownDescription: "LabelSelector is a filter to select member clusters by labels.If non-nil and non-empty, only the clusters match this filter will be selected.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_affinity": schema.SingleNestedAttribute{
								Description:         "ClusterAffinity represents scheduling restrictions to a certain set of clusters.Note:  1. ClusterAffinity can not co-exist with ClusterAffinities.  2. If both ClusterAffinity and ClusterAffinities are not set, any cluster     can be scheduling candidates.",
								MarkdownDescription: "ClusterAffinity represents scheduling restrictions to a certain set of clusters.Note:  1. ClusterAffinity can not co-exist with ClusterAffinities.  2. If both ClusterAffinity and ClusterAffinities are not set, any cluster     can be scheduling candidates.",
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
										Description:         "FieldSelector is a filter to select member clusters by fields.The key(field) of the match expression should be 'provider', 'region', or 'zone',and the operator of the match expression should be 'In' or 'NotIn'.If non-nil and non-empty, only the clusters match this filter will be selected.",
										MarkdownDescription: "FieldSelector is a filter to select member clusters by fields.The key(field) of the match expression should be 'provider', 'region', or 'zone',and the operator of the match expression should be 'In' or 'NotIn'.If non-nil and non-empty, only the clusters match this filter will be selected.",
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
															Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
															MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"values": schema.ListAttribute{
															Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
															MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
										Description:         "LabelSelector is a filter to select member clusters by labels.If non-nil and non-empty, only the clusters match this filter will be selected.",
										MarkdownDescription: "LabelSelector is a filter to select member clusters by labels.If non-nil and non-empty, only the clusters match this filter will be selected.",
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

							"cluster_tolerations": schema.ListNestedAttribute{
								Description:         "ClusterTolerations represents the tolerations.",
								MarkdownDescription: "ClusterTolerations represents the tolerations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
								Description:         "ReplicaScheduling represents the scheduling policy on dealing with the number of replicaswhen propagating resources that have replicas in spec (e.g. deployments, statefulsets) to member clusters.",
								MarkdownDescription: "ReplicaScheduling represents the scheduling policy on dealing with the number of replicaswhen propagating resources that have replicas in spec (e.g. deployments, statefulsets) to member clusters.",
								Attributes: map[string]schema.Attribute{
									"replica_division_preference": schema.StringAttribute{
										Description:         "ReplicaDivisionPreference determines how the replicas is dividedwhen ReplicaSchedulingType is 'Divided'. Valid options are Aggregated and Weighted.'Aggregated' divides replicas into clusters as few as possible,while respecting clusters' resource availabilities during the division.'Weighted' divides replicas by weight according to WeightPreference.",
										MarkdownDescription: "ReplicaDivisionPreference determines how the replicas is dividedwhen ReplicaSchedulingType is 'Divided'. Valid options are Aggregated and Weighted.'Aggregated' divides replicas into clusters as few as possible,while respecting clusters' resource availabilities during the division.'Weighted' divides replicas by weight according to WeightPreference.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Aggregated", "Weighted"),
										},
									},

									"replica_scheduling_type": schema.StringAttribute{
										Description:         "ReplicaSchedulingType determines how the replicas is scheduled when karmada propagatinga resource. Valid options are Duplicated and Divided.'Duplicated' duplicates the same replicas to each candidate member cluster from resource.'Divided' divides replicas into parts according to number of valid candidate memberclusters, and exact replicas for each cluster are determined by ReplicaDivisionPreference.",
										MarkdownDescription: "ReplicaSchedulingType determines how the replicas is scheduled when karmada propagatinga resource. Valid options are Duplicated and Divided.'Duplicated' duplicates the same replicas to each candidate member cluster from resource.'Divided' divides replicas into parts according to number of valid candidate memberclusters, and exact replicas for each cluster are determined by ReplicaDivisionPreference.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Duplicated", "Divided"),
										},
									},

									"weight_preference": schema.SingleNestedAttribute{
										Description:         "WeightPreference describes weight for each cluster or for each group of clusterIf ReplicaDivisionPreference is set to 'Weighted', and WeightPreference is not set, scheduler will weight all clusters the same.",
										MarkdownDescription: "WeightPreference describes weight for each cluster or for each group of clusterIf ReplicaDivisionPreference is set to 'Weighted', and WeightPreference is not set, scheduler will weight all clusters the same.",
										Attributes: map[string]schema.Attribute{
											"dynamic_weight": schema.StringAttribute{
												Description:         "DynamicWeight specifies the factor to generates dynamic weight list.If specified, StaticWeightList will be ignored.",
												MarkdownDescription: "DynamicWeight specifies the factor to generates dynamic weight list.If specified, StaticWeightList will be ignored.",
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
																	Description:         "FieldSelector is a filter to select member clusters by fields.The key(field) of the match expression should be 'provider', 'region', or 'zone',and the operator of the match expression should be 'In' or 'NotIn'.If non-nil and non-empty, only the clusters match this filter will be selected.",
																	MarkdownDescription: "FieldSelector is a filter to select member clusters by fields.The key(field) of the match expression should be 'provider', 'region', or 'zone',and the operator of the match expression should be 'In' or 'NotIn'.If non-nil and non-empty, only the clusters match this filter will be selected.",
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
																						Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																	Description:         "LabelSelector is a filter to select member clusters by labels.If non-nil and non-empty, only the clusters match this filter will be selected.",
																	MarkdownDescription: "LabelSelector is a filter to select member clusters by labels.If non-nil and non-empty, only the clusters match this filter will be selected.",
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
											Description:         "MinGroups restricts the minimum number of cluster groups to be selected.Defaults to 1.",
											MarkdownDescription: "MinGroups restricts the minimum number of cluster groups to be selected.Defaults to 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"spread_by_field": schema.StringAttribute{
											Description:         "SpreadByField represents the fields on Karmada cluster API used fordynamically grouping member clusters into different groups.Resources will be spread among different cluster groups.Available fields for spreading are: cluster, region, zone, and provider.SpreadByField should not co-exist with SpreadByLabel.If both SpreadByField and SpreadByLabel are empty, SpreadByField will be set to 'cluster' by system.",
											MarkdownDescription: "SpreadByField represents the fields on Karmada cluster API used fordynamically grouping member clusters into different groups.Resources will be spread among different cluster groups.Available fields for spreading are: cluster, region, zone, and provider.SpreadByField should not co-exist with SpreadByLabel.If both SpreadByField and SpreadByLabel are empty, SpreadByField will be set to 'cluster' by system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("cluster", "region", "zone", "provider"),
											},
										},

										"spread_by_label": schema.StringAttribute{
											Description:         "SpreadByLabel represents the label key used forgrouping member clusters into different groups.Resources will be spread among different cluster groups.SpreadByLabel should not co-exist with SpreadByField.",
											MarkdownDescription: "SpreadByLabel represents the label key used forgrouping member clusters into different groups.Resources will be spread among different cluster groups.SpreadByLabel should not co-exist with SpreadByField.",
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

					"propagate_deps": schema.BoolAttribute{
						Description:         "PropagateDeps tells if relevant resources should be propagated automatically.It is inherited from PropagationPolicy or ClusterPropagationPolicy.default false.",
						MarkdownDescription: "PropagateDeps tells if relevant resources should be propagated automatically.It is inherited from PropagationPolicy or ClusterPropagationPolicy.default false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_requirements": schema.SingleNestedAttribute{
						Description:         "ReplicaRequirements represents the requirements required by each replica.",
						MarkdownDescription: "ReplicaRequirements represents the requirements required by each replica.",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								Description:         "Namespace represents the resources namespaces",
								MarkdownDescription: "Namespace represents the resources namespaces",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_claim": schema.SingleNestedAttribute{
								Description:         "NodeClaim represents the node claim HardNodeAffinity, NodeSelector and Tolerations required by each replica.",
								MarkdownDescription: "NodeClaim represents the node claim HardNodeAffinity, NodeSelector and Tolerations required by each replica.",
								Attributes: map[string]schema.Attribute{
									"hard_node_affinity": schema.SingleNestedAttribute{
										Description:         "A node selector represents the union of the results of one or more label queries over a set ofnodes; that is, it represents the OR of the selectors represented by the node selector terms.Note that only PodSpec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecutionis included here because it has a hard limit on pod scheduling.",
										MarkdownDescription: "A node selector represents the union of the results of one or more label queries over a set ofnodes; that is, it represents the OR of the selectors represented by the node selector terms.Note that only PodSpec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecutionis included here because it has a hard limit on pod scheduling.",
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
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector is a selector which must be true for the pod to fit on a node.Selector which must match a node's labels for the pod to be scheduled on that node.",
										MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node.Selector which must match a node's labels for the pod to be scheduled on that node.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "If specified, the pod's tolerations.",
										MarkdownDescription: "If specified, the pod's tolerations.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"priority_class_name": schema.StringAttribute{
								Description:         "PriorityClassName represents the resources priorityClassName",
								MarkdownDescription: "PriorityClassName represents the resources priorityClassName",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_request": schema.MapAttribute{
								Description:         "ResourceRequest represents the resources required by each replica.",
								MarkdownDescription: "ResourceRequest represents the resources required by each replica.",
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

					"replicas": schema.Int64Attribute{
						Description:         "Replicas represents the replica number of the referencing resource.",
						MarkdownDescription: "Replicas represents the replica number of the referencing resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"required_by": schema.ListNestedAttribute{
						Description:         "RequiredBy represents the list of Bindings that depend on the referencing resource.",
						MarkdownDescription: "RequiredBy represents the list of Bindings that depend on the referencing resource.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"clusters": schema.ListNestedAttribute{
									Description:         "Clusters represents the scheduled result.",
									MarkdownDescription: "Clusters represents the scheduled result.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of target cluster.",
												MarkdownDescription: "Name of target cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"replicas": schema.Int64Attribute{
												Description:         "Replicas in target cluster",
												MarkdownDescription: "Replicas in target cluster",
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

								"name": schema.StringAttribute{
									Description:         "Name represents the name of the Binding.",
									MarkdownDescription: "Name represents the name of the Binding.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace represents the namespace of the Binding.It is required for ResourceBinding.If Namespace is not specified, means the referencing is ClusterResourceBinding.",
									MarkdownDescription: "Namespace represents the namespace of the Binding.It is required for ResourceBinding.If Namespace is not specified, means the referencing is ClusterResourceBinding.",
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

					"reschedule_triggered_at": schema.StringAttribute{
						Description:         "RescheduleTriggeredAt is a timestamp representing when the referenced resource is triggered rescheduling.When this field is updated, it means a rescheduling is manually triggered by user, and the expected behaviorof this action is to do a complete recalculation without referring to last scheduling results.It works with the status.lastScheduledTime field, and only when this timestamp is later than timestamp instatus.lastScheduledTime will the rescheduling actually execute, otherwise, ignored.It is represented in RFC3339 form (like '2006-01-02T15:04:05Z') and is in UTC.",
						MarkdownDescription: "RescheduleTriggeredAt is a timestamp representing when the referenced resource is triggered rescheduling.When this field is updated, it means a rescheduling is manually triggered by user, and the expected behaviorof this action is to do a complete recalculation without referring to last scheduling results.It works with the status.lastScheduledTime field, and only when this timestamp is later than timestamp instatus.lastScheduledTime will the rescheduling actually execute, otherwise, ignored.It is represented in RFC3339 form (like '2006-01-02T15:04:05Z') and is in UTC.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.DateTime64Validator(),
						},
					},

					"resource": schema.SingleNestedAttribute{
						Description:         "Resource represents the Kubernetes resource to be propagated.",
						MarkdownDescription: "Resource represents the Kubernetes resource to be propagated.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion represents the API version of the referent.",
								MarkdownDescription: "APIVersion represents the API version of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind represents the Kind of the referent.",
								MarkdownDescription: "Kind represents the Kind of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name represents the name of the referent.",
								MarkdownDescription: "Name represents the name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace represents the namespace for the referent.For non-namespace scoped resources(e.g. 'ClusterRole')，do not need specify Namespace,and for namespace scoped resources, Namespace is required.If Namespace is not specified, means the resource is non-namespace scoped.",
								MarkdownDescription: "Namespace represents the namespace for the referent.For non-namespace scoped resources(e.g. 'ClusterRole')，do not need specify Namespace,and for namespace scoped resources, Namespace is required.If Namespace is not specified, means the resource is non-namespace scoped.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "ResourceVersion represents the internal version of the referenced object, that can be used by clients todetermine when object has changed.",
								MarkdownDescription: "ResourceVersion represents the internal version of the referenced object, that can be used by clients todetermine when object has changed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.",
								MarkdownDescription: "UID of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"scheduler_name": schema.StringAttribute{
						Description:         "SchedulerName represents which scheduler to proceed the scheduling.It inherits directly from the associated PropagationPolicy(or ClusterPropagationPolicy).",
						MarkdownDescription: "SchedulerName represents which scheduler to proceed the scheduling.It inherits directly from the associated PropagationPolicy(or ClusterPropagationPolicy).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suspension": schema.SingleNestedAttribute{
						Description:         "Suspension declares the policy for suspending different aspects of propagation.nil means no suspension. no default values.",
						MarkdownDescription: "Suspension declares the policy for suspending different aspects of propagation.nil means no suspension. no default values.",
						Attributes: map[string]schema.Attribute{
							"dispatching": schema.BoolAttribute{
								Description:         "Dispatching controls whether dispatching should be suspended.nil means not suspend, no default value, only accepts 'true'.Note: true means stop propagating to all clusters. Can not co-existwith DispatchingOnClusters which is used to suspend particular clusters.",
								MarkdownDescription: "Dispatching controls whether dispatching should be suspended.nil means not suspend, no default value, only accepts 'true'.Note: true means stop propagating to all clusters. Can not co-existwith DispatchingOnClusters which is used to suspend particular clusters.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dispatching_on_clusters": schema.SingleNestedAttribute{
								Description:         "DispatchingOnClusters declares a list of clusters to which the dispatchingshould be suspended.Note: Can not co-exist with Dispatching which is used to suspend all.",
								MarkdownDescription: "DispatchingOnClusters declares a list of clusters to which the dispatchingshould be suspended.Note: Can not co-exist with Dispatching which is used to suspend all.",
								Attributes: map[string]schema.Attribute{
									"cluster_names": schema.ListAttribute{
										Description:         "ClusterNames is the list of clusters to be selected.",
										MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *WorkKarmadaIoClusterResourceBindingV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_work_karmada_io_cluster_resource_binding_v1alpha2_manifest")

	var model WorkKarmadaIoClusterResourceBindingV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("work.karmada.io/v1alpha2")
	model.Kind = pointer.String("ClusterResourceBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
