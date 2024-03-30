/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_kubeadmiral_io_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest{}
}

type CoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest struct{}

type CoreKubeadmiralIoPropagationPolicyV1Alpha1ManifestData struct {
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
		AutoMigration *struct {
			KeepUnschedulableReplicas *bool `tfsdk:"keep_unschedulable_replicas" json:"keepUnschedulableReplicas,omitempty"`
			When                      *struct {
				PodUnschedulableFor *string `tfsdk:"pod_unschedulable_for" json:"podUnschedulableFor,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"auto_migration" json:"autoMigration,omitempty"`
		ClusterAffinity *[]struct {
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
		} `tfsdk:"cluster_affinity" json:"clusterAffinity,omitempty"`
		ClusterSelector           *map[string]string `tfsdk:"cluster_selector" json:"clusterSelector,omitempty"`
		DisableFollowerScheduling *bool              `tfsdk:"disable_follower_scheduling" json:"disableFollowerScheduling,omitempty"`
		MaxClusters               *int64             `tfsdk:"max_clusters" json:"maxClusters,omitempty"`
		Placement                 *[]struct {
			Cluster     *string `tfsdk:"cluster" json:"cluster,omitempty"`
			Preferences *struct {
				MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
				MinReplicas *int64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
				Priority    *int64 `tfsdk:"priority" json:"priority,omitempty"`
				Weight      *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"preferences" json:"preferences,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
		ReplicasStrategy *string `tfsdk:"replicas_strategy" json:"replicasStrategy,omitempty"`
		ReschedulePolicy *struct {
			DisableRescheduling *bool `tfsdk:"disable_rescheduling" json:"disableRescheduling,omitempty"`
			ReplicaRescheduling *struct {
				AvoidDisruption *bool `tfsdk:"avoid_disruption" json:"avoidDisruption,omitempty"`
			} `tfsdk:"replica_rescheduling" json:"replicaRescheduling,omitempty"`
			RescheduleWhen *struct {
				ClusterAPIResourcesChanged *bool `tfsdk:"cluster_api_resources_changed" json:"clusterAPIResourcesChanged,omitempty"`
				ClusterJoined              *bool `tfsdk:"cluster_joined" json:"clusterJoined,omitempty"`
				ClusterLabelsChanged       *bool `tfsdk:"cluster_labels_changed" json:"clusterLabelsChanged,omitempty"`
				PolicyContentChanged       *bool `tfsdk:"policy_content_changed" json:"policyContentChanged,omitempty"`
			} `tfsdk:"reschedule_when" json:"rescheduleWhen,omitempty"`
		} `tfsdk:"reschedule_policy" json:"reschedulePolicy,omitempty"`
		SchedulingMode    *string `tfsdk:"scheduling_mode" json:"schedulingMode,omitempty"`
		SchedulingProfile *string `tfsdk:"scheduling_profile" json:"schedulingProfile,omitempty"`
		Tolerations       *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_propagation_policy_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PropagationPolicy describes the scheduling rules for a resource.",
		MarkdownDescription: "PropagationPolicy describes the scheduling rules for a resource.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"auto_migration": schema.SingleNestedAttribute{
						Description:         "Configures behaviors related to auto migration. If absent, auto migration will be disabled.",
						MarkdownDescription: "Configures behaviors related to auto migration. If absent, auto migration will be disabled.",
						Attributes: map[string]schema.Attribute{
							"keep_unschedulable_replicas": schema.BoolAttribute{
								Description:         "Besides starting new replicas in other cluster(s), whether to keep the unschedulable replicas in the original cluster so we can go back to the desired state when the cluster recovers.",
								MarkdownDescription: "Besides starting new replicas in other cluster(s), whether to keep the unschedulable replicas in the original cluster so we can go back to the desired state when the cluster recovers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"when": schema.SingleNestedAttribute{
								Description:         "When a replica should be subject to auto migration.",
								MarkdownDescription: "When a replica should be subject to auto migration.",
								Attributes: map[string]schema.Attribute{
									"pod_unschedulable_for": schema.StringAttribute{
										Description:         "A pod will be subject to auto migration if it remains unschedulable beyond this duration. Duration should be specified in a format that can be parsed by Go's time.ParseDuration.",
										MarkdownDescription: "A pod will be subject to auto migration if it remains unschedulable beyond this duration. Duration should be specified in a format that can be parsed by Go's time.ParseDuration.",
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

					"cluster_affinity": schema.ListNestedAttribute{
						Description:         "ClusterAffinity is a list of cluster selector terms, the terms are ORed. A empty or nil ClusterAffinity selects everything.",
						MarkdownDescription: "ClusterAffinity is a list of cluster selector terms, the terms are ORed. A empty or nil ClusterAffinity selects everything.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match_expressions": schema.ListNestedAttribute{
									Description:         "A list of cluster selector requirements by cluster labels.",
									MarkdownDescription: "A list of cluster selector requirements by cluster labels.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operator": schema.StringAttribute{
												Description:         "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
												MarkdownDescription: "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
												},
											},

											"values": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"match_fields": schema.ListNestedAttribute{
									Description:         "A list of cluster selector requirements by cluster fields.",
									MarkdownDescription: "A list of cluster selector requirements by cluster fields.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operator": schema.StringAttribute{
												Description:         "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
												MarkdownDescription: "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
												},
											},

											"values": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_selector": schema.MapAttribute{
						Description:         "ClusterSelector is a label query over clusters to consider for scheduling. An empty or nil ClusterSelector selects everything.",
						MarkdownDescription: "ClusterSelector is a label query over clusters to consider for scheduling. An empty or nil ClusterSelector selects everything.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_follower_scheduling": schema.BoolAttribute{
						Description:         "DisableFollowerScheduling is a boolean that determines if follower scheduling is disabled. Resources that depend on other resources (e.g. deployments) are called leaders, and resources that are depended on (e.g. configmaps and secrets) are called followers. If a leader enables follower scheduling, its followers will additionally be scheduled to clusters where the leader is scheduled.",
						MarkdownDescription: "DisableFollowerScheduling is a boolean that determines if follower scheduling is disabled. Resources that depend on other resources (e.g. deployments) are called leaders, and resources that are depended on (e.g. configmaps and secrets) are called followers. If a leader enables follower scheduling, its followers will additionally be scheduled to clusters where the leader is scheduled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_clusters": schema.Int64Attribute{
						Description:         "MaxClusters is the maximum number of replicas that the federated object can be propagated to. The maximum number of clusters is unbounded if no value is provided.",
						MarkdownDescription: "MaxClusters is the maximum number of replicas that the federated object can be propagated to. The maximum number of clusters is unbounded if no value is provided.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"placement": schema.ListNestedAttribute{
						Description:         "Placement is an explicit list of clusters used to select member clusters to propagate resources to.",
						MarkdownDescription: "Placement is an explicit list of clusters used to select member clusters to propagate resources to.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster": schema.StringAttribute{
									Description:         "Cluster is the name of the FederatedCluster to propagate to.",
									MarkdownDescription: "Cluster is the name of the FederatedCluster to propagate to.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"preferences": schema.SingleNestedAttribute{
									Description:         "Preferences contains the cluster's propagation preferences.",
									MarkdownDescription: "Preferences contains the cluster's propagation preferences.",
									Attributes: map[string]schema.Attribute{
										"max_replicas": schema.Int64Attribute{
											Description:         "Maximum number of replicas that should be assigned to this cluster workload object. Unbounded if no value provided (default).",
											MarkdownDescription: "Maximum number of replicas that should be assigned to this cluster workload object. Unbounded if no value provided (default).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"min_replicas": schema.Int64Attribute{
											Description:         "Minimum number of replicas that should be assigned to this cluster workload object. 0 by default.",
											MarkdownDescription: "Minimum number of replicas that should be assigned to this cluster workload object. 0 by default.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"priority": schema.Int64Attribute{
											Description:         "A number expressing the priority of the cluster. The higher the value, the higher the priority. When selecting clusters for propagation, clusters with higher priority are preferred. When the Binpack ReplicasStrategy is selected, replicas will be scheduled to clusters with higher priority first.",
											MarkdownDescription: "A number expressing the priority of the cluster. The higher the value, the higher the priority. When selecting clusters for propagation, clusters with higher priority are preferred. When the Binpack ReplicasStrategy is selected, replicas will be scheduled to clusters with higher priority first.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"weight": schema.Int64Attribute{
											Description:         "A number expressing the preference to put an additional replica to this cluster workload object. It will not take effect when ReplicasStrategy is Binpack.",
											MarkdownDescription: "A number expressing the preference to put an additional replica to this cluster workload object. It will not take effect when ReplicasStrategy is Binpack.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
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

					"replicas_strategy": schema.StringAttribute{
						Description:         "ReplicasStrategy is the strategy used for scheduling replicas.",
						MarkdownDescription: "ReplicasStrategy is the strategy used for scheduling replicas.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Binpack", "Spread"),
						},
					},

					"reschedule_policy": schema.SingleNestedAttribute{
						Description:         "Configures behaviors related to rescheduling.",
						MarkdownDescription: "Configures behaviors related to rescheduling.",
						Attributes: map[string]schema.Attribute{
							"disable_rescheduling": schema.BoolAttribute{
								Description:         "DisableRescheduling determines if a federated object can be rescheduled.",
								MarkdownDescription: "DisableRescheduling determines if a federated object can be rescheduled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_rescheduling": schema.SingleNestedAttribute{
								Description:         "Configures behaviors related to replica rescheduling. Default set via a post-generation patch. See patch file for details.",
								MarkdownDescription: "Configures behaviors related to replica rescheduling. Default set via a post-generation patch. See patch file for details.",
								Attributes: map[string]schema.Attribute{
									"avoid_disruption": schema.BoolAttribute{
										Description:         "If set to true, the scheduler will attempt to prevent migrating existing replicas during rescheduling. In order to do so, replica scheduling preferences might not be fully respected. If set to false, the scheduler will always rebalance the replicas based on the specified preferences, which might cause temporary service disruption.",
										MarkdownDescription: "If set to true, the scheduler will attempt to prevent migrating existing replicas during rescheduling. In order to do so, replica scheduling preferences might not be fully respected. If set to false, the scheduler will always rebalance the replicas based on the specified preferences, which might cause temporary service disruption.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"reschedule_when": schema.SingleNestedAttribute{
								Description:         "When the related objects should be subject to reschedule.",
								MarkdownDescription: "When the related objects should be subject to reschedule.",
								Attributes: map[string]schema.Attribute{
									"cluster_api_resources_changed": schema.BoolAttribute{
										Description:         "If set to true, changes to clusters' enabled list of api resources will trigger rescheduling. It set to false, the scheduler will reschedule only when other options are triggered or the replicas or the requested resources of the template changed.",
										MarkdownDescription: "If set to true, changes to clusters' enabled list of api resources will trigger rescheduling. It set to false, the scheduler will reschedule only when other options are triggered or the replicas or the requested resources of the template changed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_joined": schema.BoolAttribute{
										Description:         "If set to true, clusters joining the federation will trigger rescheduling. It set to false, the scheduler will reschedule only when other options are triggered or the replicas or the requested resources of the template changed.",
										MarkdownDescription: "If set to true, clusters joining the federation will trigger rescheduling. It set to false, the scheduler will reschedule only when other options are triggered or the replicas or the requested resources of the template changed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_labels_changed": schema.BoolAttribute{
										Description:         "If set to true, changes to cluster labels will trigger rescheduling. It set to false, the scheduler will reschedule only when other options are triggered or the replicas or the requested resources of the template changed.",
										MarkdownDescription: "If set to true, changes to cluster labels will trigger rescheduling. It set to false, the scheduler will reschedule only when other options are triggered or the replicas or the requested resources of the template changed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"policy_content_changed": schema.BoolAttribute{
										Description:         "If set to true, the scheduler will trigger rescheduling when the semantics of the policy changes. For example, modifying placement, schedulingMode, maxClusters, clusterSelector, and other configurations related to scheduling (includes reschedulePolicy itself) will immediately trigger rescheduling. Modifying the labels, annotations, autoMigration configuration will not trigger rescheduling. It set to false, the scheduler will not reschedule when the policy content changes.",
										MarkdownDescription: "If set to true, the scheduler will trigger rescheduling when the semantics of the policy changes. For example, modifying placement, schedulingMode, maxClusters, clusterSelector, and other configurations related to scheduling (includes reschedulePolicy itself) will immediately trigger rescheduling. Modifying the labels, annotations, autoMigration configuration will not trigger rescheduling. It set to false, the scheduler will not reschedule when the policy content changes.",
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

					"scheduling_mode": schema.StringAttribute{
						Description:         "SchedulingMode determines the mode used for scheduling.",
						MarkdownDescription: "SchedulingMode determines the mode used for scheduling.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Duplicate", "Divide"),
						},
					},

					"scheduling_profile": schema.StringAttribute{
						Description:         "Profile determines the scheduling profile to be used for scheduling",
						MarkdownDescription: "Profile determines the scheduling profile to be used for scheduling",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations describe a set of cluster taints that the policy tolerates.",
						MarkdownDescription: "Tolerations describe a set of cluster taints that the policy tolerates.",
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CoreKubeadmiralIoPropagationPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_propagation_policy_v1alpha1_manifest")

	var model CoreKubeadmiralIoPropagationPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("PropagationPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
