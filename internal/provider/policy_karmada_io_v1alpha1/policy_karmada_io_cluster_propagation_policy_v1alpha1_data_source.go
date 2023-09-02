/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policy_karmada_io_v1alpha1

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
	_ datasource.DataSource              = &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource{}
)

func NewPolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource() datasource.DataSource {
	return &PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource{}
}

type PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policy_karmada_io_cluster_propagation_policy_v1alpha1"
}

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec represents the desired behavior of ClusterPropagationPolicy.",
				MarkdownDescription: "Spec represents the desired behavior of ClusterPropagationPolicy.",
				Attributes: map[string]schema.Attribute{
					"association": schema.BoolAttribute{
						Description:         "Association tells if relevant resources should be selected automatically. e.g. a ConfigMap referred by a Deployment. default false. Deprecated: in favor of PropagateDeps.",
						MarkdownDescription: "Association tells if relevant resources should be selected automatically. e.g. a ConfigMap referred by a Deployment. default false. Deprecated: in favor of PropagateDeps.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"conflict_resolution": schema.StringAttribute{
						Description:         "ConflictResolution declares how potential conflict should be handled when a resource that is being propagated already exists in the target cluster.  It defaults to 'Abort' which means stop propagating to avoid unexpected overwrites. The 'Overwrite' might be useful when migrating legacy cluster resources to Karmada, in which case conflict is predictable and can be instructed to Karmada take over the resource by overwriting.",
						MarkdownDescription: "ConflictResolution declares how potential conflict should be handled when a resource that is being propagated already exists in the target cluster.  It defaults to 'Abort' which means stop propagating to avoid unexpected overwrites. The 'Overwrite' might be useful when migrating legacy cluster resources to Karmada, in which case conflict is predictable and can be instructed to Karmada take over the resource by overwriting.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"dependent_overrides": schema.ListAttribute{
						Description:         "DependentOverrides represents the list of overrides(OverridePolicy) which must present before the current PropagationPolicy takes effect.  It used to explicitly specify overrides which current PropagationPolicy rely on. A typical scenario is the users create OverridePolicy(ies) and resources at the same time, they want to ensure the new-created policies would be adopted.  Note: For the overrides, OverridePolicy(ies) in current namespace and ClusterOverridePolicy(ies), which not present in this list will still be applied if they matches the resources.",
						MarkdownDescription: "DependentOverrides represents the list of overrides(OverridePolicy) which must present before the current PropagationPolicy takes effect.  It used to explicitly specify overrides which current PropagationPolicy rely on. A typical scenario is the users create OverridePolicy(ies) and resources at the same time, they want to ensure the new-created policies would be adopted.  Note: For the overrides, OverridePolicy(ies) in current namespace and ClusterOverridePolicy(ies), which not present in this list will still be applied if they matches the resources.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"grace_period_seconds": schema.Int64Attribute{
										Description:         "GracePeriodSeconds is the maximum waiting duration in seconds before application on the migrated cluster should be deleted. Required only when PurgeMode is 'Graciously' and defaults to 600s. If the application on the new cluster cannot reach a Healthy state, Karmada will delete the application after GracePeriodSeconds is reached. Value must be positive integer.",
										MarkdownDescription: "GracePeriodSeconds is the maximum waiting duration in seconds before application on the migrated cluster should be deleted. Required only when PurgeMode is 'Graciously' and defaults to 600s. If the application on the new cluster cannot reach a Healthy state, Karmada will delete the application after GracePeriodSeconds is reached. Value must be positive integer.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"purge_mode": schema.StringAttribute{
										Description:         "PurgeMode represents how to deal with the legacy applications on the cluster from which the application is migrated. Valid options are 'Immediately', 'Graciously' and 'Never'. Defaults to 'Graciously'.",
										MarkdownDescription: "PurgeMode represents how to deal with the legacy applications on the cluster from which the application is migrated. Valid options are 'Immediately', 'Graciously' and 'Never'. Defaults to 'Graciously'.",
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
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"cluster_names": schema.ListAttribute{
											Description:         "ClusterNames is the list of clusters to be selected.",
											MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"exclude": schema.ListAttribute{
											Description:         "ExcludedClusters is the list of clusters to be ignored.",
											MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
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

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
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
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"exclude": schema.ListAttribute{
										Description:         "ExcludedClusters is the list of clusters to be ignored.",
										MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
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
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"operator": schema.StringAttribute{
															Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
															MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"values": schema.ListAttribute{
															Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
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
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
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

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
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

							"cluster_tolerations": schema.ListNestedAttribute{
								Description:         "ClusterTolerations represents the tolerations.",
								MarkdownDescription: "ClusterTolerations represents the tolerations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"replica_scheduling": schema.SingleNestedAttribute{
								Description:         "ReplicaScheduling represents the scheduling policy on dealing with the number of replicas when propagating resources that have replicas in spec (e.g. deployments, statefulsets) to member clusters.",
								MarkdownDescription: "ReplicaScheduling represents the scheduling policy on dealing with the number of replicas when propagating resources that have replicas in spec (e.g. deployments, statefulsets) to member clusters.",
								Attributes: map[string]schema.Attribute{
									"replica_division_preference": schema.StringAttribute{
										Description:         "ReplicaDivisionPreference determines how the replicas is divided when ReplicaSchedulingType is 'Divided'. Valid options are Aggregated and Weighted. 'Aggregated' divides replicas into clusters as few as possible, while respecting clusters' resource availabilities during the division. 'Weighted' divides replicas by weight according to WeightPreference.",
										MarkdownDescription: "ReplicaDivisionPreference determines how the replicas is divided when ReplicaSchedulingType is 'Divided'. Valid options are Aggregated and Weighted. 'Aggregated' divides replicas into clusters as few as possible, while respecting clusters' resource availabilities during the division. 'Weighted' divides replicas by weight according to WeightPreference.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replica_scheduling_type": schema.StringAttribute{
										Description:         "ReplicaSchedulingType determines how the replicas is scheduled when karmada propagating a resource. Valid options are Duplicated and Divided. 'Duplicated' duplicates the same replicas to each candidate member cluster from resource. 'Divided' divides replicas into parts according to number of valid candidate member clusters, and exact replicas for each cluster are determined by ReplicaDivisionPreference.",
										MarkdownDescription: "ReplicaSchedulingType determines how the replicas is scheduled when karmada propagating a resource. Valid options are Duplicated and Divided. 'Duplicated' duplicates the same replicas to each candidate member cluster from resource. 'Divided' divides replicas into parts according to number of valid candidate member clusters, and exact replicas for each cluster are determined by ReplicaDivisionPreference.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"weight_preference": schema.SingleNestedAttribute{
										Description:         "WeightPreference describes weight for each cluster or for each group of cluster If ReplicaDivisionPreference is set to 'Weighted', and WeightPreference is not set, scheduler will weight all clusters the same.",
										MarkdownDescription: "WeightPreference describes weight for each cluster or for each group of cluster If ReplicaDivisionPreference is set to 'Weighted', and WeightPreference is not set, scheduler will weight all clusters the same.",
										Attributes: map[string]schema.Attribute{
											"dynamic_weight": schema.StringAttribute{
												Description:         "DynamicWeight specifies the factor to generates dynamic weight list. If specified, StaticWeightList will be ignored.",
												MarkdownDescription: "DynamicWeight specifies the factor to generates dynamic weight list. If specified, StaticWeightList will be ignored.",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																	Optional:            false,
																	Computed:            true,
																},

																"exclude": schema.ListAttribute{
																	Description:         "ExcludedClusters is the list of clusters to be ignored.",
																	MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																						MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
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

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
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

														"weight": schema.Int64Attribute{
															Description:         "Weight expressing the preference to the cluster(s) specified by 'TargetCluster'.",
															MarkdownDescription: "Weight expressing the preference to the cluster(s) specified by 'TargetCluster'.",
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
								Required: false,
								Optional: false,
								Computed: true,
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
											Optional:            false,
											Computed:            true,
										},

										"min_groups": schema.Int64Attribute{
											Description:         "MinGroups restricts the minimum number of cluster groups to be selected. Defaults to 1.",
											MarkdownDescription: "MinGroups restricts the minimum number of cluster groups to be selected. Defaults to 1.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"spread_by_field": schema.StringAttribute{
											Description:         "SpreadByField represents the fields on Karmada cluster API used for dynamically grouping member clusters into different groups. Resources will be spread among different cluster groups. Available fields for spreading are: cluster, region, zone, and provider. SpreadByField should not co-exist with SpreadByLabel. If both SpreadByField and SpreadByLabel are empty, SpreadByField will be set to 'cluster' by system.",
											MarkdownDescription: "SpreadByField represents the fields on Karmada cluster API used for dynamically grouping member clusters into different groups. Resources will be spread among different cluster groups. Available fields for spreading are: cluster, region, zone, and provider. SpreadByField should not co-exist with SpreadByLabel. If both SpreadByField and SpreadByLabel are empty, SpreadByField will be set to 'cluster' by system.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"spread_by_label": schema.StringAttribute{
											Description:         "SpreadByLabel represents the label key used for grouping member clusters into different groups. Resources will be spread among different cluster groups. SpreadByLabel should not co-exist with SpreadByField.",
											MarkdownDescription: "SpreadByLabel represents the label key used for grouping member clusters into different groups. Resources will be spread among different cluster groups. SpreadByLabel should not co-exist with SpreadByField.",
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

					"preemption": schema.StringAttribute{
						Description:         "Preemption declares the behaviors for preempting. Valid options are 'Always' and 'Never'.",
						MarkdownDescription: "Preemption declares the behaviors for preempting. Valid options are 'Always' and 'Never'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority indicates the importance of a policy(PropagationPolicy or ClusterPropagationPolicy). A policy will be applied for the matched resource templates if there is no other policies with higher priority at the point of the resource template be processed. Once a resource template has been claimed by a policy, by default it will not be preempted by following policies even with a higher priority. See Preemption for more details.  In case of two policies have the same priority, the one with a more precise matching rules in ResourceSelectors wins: - matching by name(resourceSelector.name) has higher priority than by selector(resourceSelector.labelSelector) - matching by selector(resourceSelector.labelSelector) has higher priority than by APIVersion(resourceSelector.apiVersion) and Kind(resourceSelector.kind). If there is still no winner at this point, the one with the lower alphabetic order wins, e.g. policy 'bar' has higher priority than 'foo'.  The higher the value, the higher the priority. Defaults to zero.",
						MarkdownDescription: "Priority indicates the importance of a policy(PropagationPolicy or ClusterPropagationPolicy). A policy will be applied for the matched resource templates if there is no other policies with higher priority at the point of the resource template be processed. Once a resource template has been claimed by a policy, by default it will not be preempted by following policies even with a higher priority. See Preemption for more details.  In case of two policies have the same priority, the one with a more precise matching rules in ResourceSelectors wins: - matching by name(resourceSelector.name) has higher priority than by selector(resourceSelector.labelSelector) - matching by selector(resourceSelector.labelSelector) has higher priority than by APIVersion(resourceSelector.apiVersion) and Kind(resourceSelector.kind). If there is still no winner at this point, the one with the lower alphabetic order wins, e.g. policy 'bar' has higher priority than 'foo'.  The higher the value, the higher the priority. Defaults to zero.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"propagate_deps": schema.BoolAttribute{
						Description:         "PropagateDeps tells if relevant resources should be propagated automatically. Take 'Deployment' which referencing 'ConfigMap' and 'Secret' as an example, when 'propagateDeps' is 'true', the referencing resources could be omitted(for saving config effort) from 'resourceSelectors' as they will be propagated along with the Deployment. In addition to the propagating process, the referencing resources will be migrated along with the Deployment in the fail-over scenario.  Defaults to false.",
						MarkdownDescription: "PropagateDeps tells if relevant resources should be propagated automatically. Take 'Deployment' which referencing 'ConfigMap' and 'Secret' as an example, when 'propagateDeps' is 'true', the referencing resources could be omitted(for saving config effort) from 'resourceSelectors' as they will be propagated along with the Deployment. In addition to the propagating process, the referencing resources will be migrated along with the Deployment in the fail-over scenario.  Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_selectors": schema.ListNestedAttribute{
						Description:         "ResourceSelectors used to select resources. Nil or empty selector is not allowed and doesn't mean match all kinds of resources for security concerns that sensitive resources(like Secret) might be accidentally propagated.",
						MarkdownDescription: "ResourceSelectors used to select resources. Nil or empty selector is not allowed and doesn't mean match all kinds of resources for security concerns that sensitive resources(like Secret) might be accidentally propagated.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion represents the API version of the target resources.",
									MarkdownDescription: "APIVersion represents the API version of the target resources.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind represents the Kind of the target resources.",
									MarkdownDescription: "Kind represents the Kind of the target resources.",
									Required:            false,
									Optional:            false,
									Computed:            true,
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
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
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

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the target resource. Default is empty, which means selecting all resources.",
									MarkdownDescription: "Name of the target resource. Default is empty, which means selecting all resources.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the target resource. Default is empty, which means inherit from the parent object scope.",
									MarkdownDescription: "Namespace of the target resource. Default is empty, which means inherit from the parent object scope.",
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

					"scheduler_name": schema.StringAttribute{
						Description:         "SchedulerName represents which scheduler to proceed the scheduling. If specified, the policy will be dispatched by specified scheduler. If not specified, the policy will be dispatched by default scheduler.",
						MarkdownDescription: "SchedulerName represents which scheduler to proceed the scheduling. If specified, the policy will be dispatched by specified scheduler. If not specified, the policy will be dispatched by default scheduler.",
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
	}
}

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1")

	var data PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "ClusterPropagationPolicy"}).
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

	var readResponse PolicyKarmadaIoClusterPropagationPolicyV1Alpha1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("policy.karmada.io/v1alpha1")
	data.Kind = pointer.String("ClusterPropagationPolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
