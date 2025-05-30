/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operations_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &OperationsKubeblocksIoOpsRequestV1Alpha1Manifest{}
)

func NewOperationsKubeblocksIoOpsRequestV1Alpha1Manifest() datasource.DataSource {
	return &OperationsKubeblocksIoOpsRequestV1Alpha1Manifest{}
}

type OperationsKubeblocksIoOpsRequestV1Alpha1Manifest struct{}

type OperationsKubeblocksIoOpsRequestV1Alpha1ManifestData struct {
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
		Backup *struct {
			BackupMethod     *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
			BackupName       *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			BackupPolicyName *string `tfsdk:"backup_policy_name" json:"backupPolicyName,omitempty"`
			DeletionPolicy   *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
			Parameters       *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			ParentBackupName *string `tfsdk:"parent_backup_name" json:"parentBackupName,omitempty"`
			RetentionPeriod  *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		Cancel      *bool   `tfsdk:"cancel" json:"cancel,omitempty"`
		ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		Custom      *struct {
			Components *[]struct {
				ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
				Parameters    *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
			MaxConcurrentComponents *string `tfsdk:"max_concurrent_components" json:"maxConcurrentComponents,omitempty"`
			OpsDefinitionName       *string `tfsdk:"ops_definition_name" json:"opsDefinitionName,omitempty"`
			ServiceAccountName      *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		} `tfsdk:"custom" json:"custom,omitempty"`
		EnqueueOnForce *bool `tfsdk:"enqueue_on_force" json:"enqueueOnForce,omitempty"`
		Expose         *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Services      *[]struct {
				Annotations    *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				IpFamilies     *[]string          `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
				IpFamilyPolicy *string            `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
				Name           *string            `tfsdk:"name" json:"name,omitempty"`
				PodSelector    *map[string]string `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				Ports          *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
					TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				RoleSelector *string `tfsdk:"role_selector" json:"roleSelector,omitempty"`
				ServiceType  *string `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			Switch *string `tfsdk:"switch" json:"switch,omitempty"`
		} `tfsdk:"expose" json:"expose,omitempty"`
		Force             *bool `tfsdk:"force" json:"force,omitempty"`
		HorizontalScaling *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			ScaleIn       *struct {
				Instances *[]struct {
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					ReplicaChanges *int64  `tfsdk:"replica_changes" json:"replicaChanges,omitempty"`
				} `tfsdk:"instances" json:"instances,omitempty"`
				OnlineInstancesToOffline *[]string `tfsdk:"online_instances_to_offline" json:"onlineInstancesToOffline,omitempty"`
				ReplicaChanges           *int64    `tfsdk:"replica_changes" json:"replicaChanges,omitempty"`
			} `tfsdk:"scale_in" json:"scaleIn,omitempty"`
			ScaleOut *struct {
				Instances *[]struct {
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					ReplicaChanges *int64  `tfsdk:"replica_changes" json:"replicaChanges,omitempty"`
				} `tfsdk:"instances" json:"instances,omitempty"`
				NewInstances *[]struct {
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
					Labels   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name     *string            `tfsdk:"name" json:"name,omitempty"`
					Ordinals *struct {
						Discrete *[]string `tfsdk:"discrete" json:"discrete,omitempty"`
						Ranges   *[]struct {
							End   *int64 `tfsdk:"end" json:"end,omitempty"`
							Start *int64 `tfsdk:"start" json:"start,omitempty"`
						} `tfsdk:"ranges" json:"ranges,omitempty"`
					} `tfsdk:"ordinals" json:"ordinals,omitempty"`
					Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
					Resources *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					SchedulingPolicy *struct {
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
										MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
										MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
										MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
										MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
						NodeName      *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
						NodeSelector  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						SchedulerName *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
						Tolerations   *[]struct {
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
					} `tfsdk:"scheduling_policy" json:"schedulingPolicy,omitempty"`
					VolumeClaimTemplates *[]struct {
						Annotations               *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels                    *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name                      *string            `tfsdk:"name" json:"name,omitempty"`
						PersistentVolumeClaimName *string            `tfsdk:"persistent_volume_claim_name" json:"persistentVolumeClaimName,omitempty"`
						Spec                      *struct {
							AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
							DataSource  *struct {
								ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
								Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"data_source" json:"dataSource,omitempty"`
							DataSourceRef *struct {
								ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
							Resources *struct {
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
							StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
							VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
							VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
							VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
				} `tfsdk:"new_instances" json:"newInstances,omitempty"`
				OfflineInstancesToOnline *[]string `tfsdk:"offline_instances_to_online" json:"offlineInstancesToOnline,omitempty"`
				ReplicaChanges           *int64    `tfsdk:"replica_changes" json:"replicaChanges,omitempty"`
			} `tfsdk:"scale_out" json:"scaleOut,omitempty"`
			Shards *int64 `tfsdk:"shards" json:"shards,omitempty"`
		} `tfsdk:"horizontal_scaling" json:"horizontalScaling,omitempty"`
		PreConditionDeadlineSeconds *int64 `tfsdk:"pre_condition_deadline_seconds" json:"preConditionDeadlineSeconds,omitempty"`
		RebuildFrom                 *[]struct {
			BackupName    *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			InPlace       *bool   `tfsdk:"in_place" json:"inPlace,omitempty"`
			Instances     *[]struct {
				Name           *string `tfsdk:"name" json:"name,omitempty"`
				TargetNodeName *string `tfsdk:"target_node_name" json:"targetNodeName,omitempty"`
			} `tfsdk:"instances" json:"instances,omitempty"`
			RestoreEnv             *map[string]string `tfsdk:"restore_env" json:"restoreEnv,omitempty"`
			SourceBackupTargetName *string            `tfsdk:"source_backup_target_name" json:"sourceBackupTargetName,omitempty"`
		} `tfsdk:"rebuild_from" json:"rebuildFrom,omitempty"`
		Reconfigures *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Parameters    *[]struct {
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
		} `tfsdk:"reconfigures" json:"reconfigures,omitempty"`
		Restart *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
		} `tfsdk:"restart" json:"restart,omitempty"`
		Restore *struct {
			BackupName                        *string            `tfsdk:"backup_name" json:"backupName,omitempty"`
			BackupNamespace                   *string            `tfsdk:"backup_namespace" json:"backupNamespace,omitempty"`
			DeferPostReadyUntilClusterRunning *bool              `tfsdk:"defer_post_ready_until_cluster_running" json:"deferPostReadyUntilClusterRunning,omitempty"`
			Env                               *map[string]string `tfsdk:"env" json:"env,omitempty"`
			Parameters                        *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			RestorePointInTime  *string `tfsdk:"restore_point_in_time" json:"restorePointInTime,omitempty"`
			VolumeRestorePolicy *string `tfsdk:"volume_restore_policy" json:"volumeRestorePolicy,omitempty"`
		} `tfsdk:"restore" json:"restore,omitempty"`
		Start *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
		} `tfsdk:"start" json:"start,omitempty"`
		Stop *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
		} `tfsdk:"stop" json:"stop,omitempty"`
		Switchover *[]struct {
			CandidateName       *string `tfsdk:"candidate_name" json:"candidateName,omitempty"`
			ComponentName       *string `tfsdk:"component_name" json:"componentName,omitempty"`
			ComponentObjectName *string `tfsdk:"component_object_name" json:"componentObjectName,omitempty"`
			InstanceName        *string `tfsdk:"instance_name" json:"instanceName,omitempty"`
		} `tfsdk:"switchover" json:"switchover,omitempty"`
		TimeoutSeconds                        *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		TtlSecondsAfterSucceed                *int64  `tfsdk:"ttl_seconds_after_succeed" json:"ttlSecondsAfterSucceed,omitempty"`
		TtlSecondsAfterUnsuccessfulCompletion *int64  `tfsdk:"ttl_seconds_after_unsuccessful_completion" json:"ttlSecondsAfterUnsuccessfulCompletion,omitempty"`
		Type                                  *string `tfsdk:"type" json:"type,omitempty"`
		Upgrade                               *struct {
			Components *[]struct {
				ComponentDefinitionName *string `tfsdk:"component_definition_name" json:"componentDefinitionName,omitempty"`
				ComponentName           *string `tfsdk:"component_name" json:"componentName,omitempty"`
				ServiceVersion          *string `tfsdk:"service_version" json:"serviceVersion,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
		} `tfsdk:"upgrade" json:"upgrade,omitempty"`
		VerticalScaling *[]map[string]string `tfsdk:"vertical_scaling" json:"verticalScaling,omitempty"`
		VolumeExpansion *[]struct {
			ComponentName        *string `tfsdk:"component_name" json:"componentName,omitempty"`
			VolumeClaimTemplates *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"volume_expansion" json:"volumeExpansion,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperationsKubeblocksIoOpsRequestV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operations_kubeblocks_io_ops_request_v1alpha1_manifest"
}

func (r *OperationsKubeblocksIoOpsRequestV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OpsRequest is the Schema for the opsrequests API",
		MarkdownDescription: "OpsRequest is the Schema for the opsrequests API",
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
				Description:         "OpsRequestSpec defines the desired state of OpsRequest",
				MarkdownDescription: "OpsRequestSpec defines the desired state of OpsRequest",
				Attributes: map[string]schema.Attribute{
					"backup": schema.SingleNestedAttribute{
						Description:         "Specifies the parameters to back up a Cluster.",
						MarkdownDescription: "Specifies the parameters to back up a Cluster.",
						Attributes: map[string]schema.Attribute{
							"backup_method": schema.StringAttribute{
								Description:         "Specifies the name of BackupMethod. The specified BackupMethod must be defined in the BackupPolicy.",
								MarkdownDescription: "Specifies the name of BackupMethod. The specified BackupMethod must be defined in the BackupPolicy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the Backup custom resource.",
								MarkdownDescription: "Specifies the name of the Backup custom resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backup_policy_name": schema.StringAttribute{
								Description:         "Indicates the name of the BackupPolicy applied to perform this Backup.",
								MarkdownDescription: "Indicates the name of the BackupPolicy applied to perform this Backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deletion_policy": schema.StringAttribute{
								Description:         "Determines whether the backup contents stored in backup repository should be deleted when the Backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.",
								MarkdownDescription: "Determines whether the backup contents stored in backup repository should be deleted when the Backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Delete", "Retain"),
								},
							},

							"parameters": schema.ListNestedAttribute{
								Description:         "Specifies a list of name-value pairs representing parameters and their corresponding values. Parameters match the schema specified in the 'actionset.spec.parametersSchema'",
								MarkdownDescription: "Specifies a list of name-value pairs representing parameters and their corresponding values. Parameters match the schema specified in the 'actionset.spec.parametersSchema'",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Represents the name of the parameter.",
											MarkdownDescription: "Represents the name of the parameter.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Represents the parameter values.",
											MarkdownDescription: "Represents the parameter values.",
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

							"parent_backup_name": schema.StringAttribute{
								Description:         "If the specified BackupMethod is incremental, 'parentBackupName' is required.",
								MarkdownDescription: "If the specified BackupMethod is incremental, 'parentBackupName' is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retention_period": schema.StringAttribute{
								Description:         "Determines the duration for which the Backup custom resources should be retained. The controller will automatically remove all Backup objects that are older than the specified RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the Backup objects of last 30 days. Sample duration format: - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m You can also combine the above durations. For example: 30d12h30m. If not set, the Backup objects will be kept forever. If the 'deletionPolicy' is set to 'Delete', then the associated backup data will also be deleted along with the Backup object. Otherwise, only the Backup custom resource will be deleted.",
								MarkdownDescription: "Determines the duration for which the Backup custom resources should be retained. The controller will automatically remove all Backup objects that are older than the specified RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the Backup objects of last 30 days. Sample duration format: - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m You can also combine the above durations. For example: 30d12h30m. If not set, the Backup objects will be kept forever. If the 'deletionPolicy' is set to 'Delete', then the associated backup data will also be deleted along with the Backup object. Otherwise, only the Backup custom resource will be deleted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cancel": schema.BoolAttribute{
						Description:         "Indicates whether the current operation should be canceled and terminated gracefully if it's in the 'Pending', 'Creating', or 'Running' state. This field applies only to 'VerticalScaling' and 'HorizontalScaling' opsRequests. Note: Setting 'cancel' to true is irreversible; further modifications to this field are ineffective.",
						MarkdownDescription: "Indicates whether the current operation should be canceled and terminated gracefully if it's in the 'Pending', 'Creating', or 'Running' state. This field applies only to 'VerticalScaling' and 'HorizontalScaling' opsRequests. Note: Setting 'cancel' to true is irreversible; further modifications to this field are ineffective.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "Specifies the name of the Cluster resource that this operation is targeting.",
						MarkdownDescription: "Specifies the name of the Cluster resource that this operation is targeting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom": schema.SingleNestedAttribute{
						Description:         "Specifies a custom operation defined by OpsDefinition.",
						MarkdownDescription: "Specifies a custom operation defined by OpsDefinition.",
						Attributes: map[string]schema.Attribute{
							"components": schema.ListNestedAttribute{
								Description:         "Specifies the components and their parameters for executing custom actions as defined in OpsDefinition. Requires at least one component.",
								MarkdownDescription: "Specifies the components and their parameters for executing custom actions as defined in OpsDefinition. Requires at least one component.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"component_name": schema.StringAttribute{
											Description:         "Specifies the name of the Component as defined in the cluster.spec",
											MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameters": schema.ListNestedAttribute{
											Description:         "Specifies the parameters that match the schema specified in the 'opsDefinition.spec.parametersSchema'.",
											MarkdownDescription: "Specifies the parameters that match the schema specified in the 'opsDefinition.spec.parametersSchema'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Specifies the identifier of the parameter as defined in the OpsDefinition.",
														MarkdownDescription: "Specifies the identifier of the parameter as defined in the OpsDefinition.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.",
														MarkdownDescription: "Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value_from": schema.SingleNestedAttribute{
														Description:         "Source for the parameter's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the parameter's value. Cannot be used if value is not empty.",
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

															"secret_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret.",
																MarkdownDescription: "Selects a key of a Secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from. Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
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
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"max_concurrent_components": schema.StringAttribute{
								Description:         "Specifies the maximum number of components to be operated on concurrently to mitigate performance impact on clusters with multiple components. It accepts an absolute number (e.g., 5) or a percentage of components to execute in parallel (e.g., '10%'). Percentages are rounded up to the nearest whole number of components. For example, if '10%' results in less than one, it rounds up to 1. When unspecified, all components are processed simultaneously by default. Note: This feature is not implemented yet.",
								MarkdownDescription: "Specifies the maximum number of components to be operated on concurrently to mitigate performance impact on clusters with multiple components. It accepts an absolute number (e.g., 5) or a percentage of components to execute in parallel (e.g., '10%'). Percentages are rounded up to the nearest whole number of components. For example, if '10%' results in less than one, it rounds up to 1. When unspecified, all components are processed simultaneously by default. Note: This feature is not implemented yet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ops_definition_name": schema.StringAttribute{
								Description:         "Specifies the name of the OpsDefinition.",
								MarkdownDescription: "Specifies the name of the OpsDefinition.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"service_account_name": schema.StringAttribute{
								Description:         "Specifies the name of the ServiceAccount to be used for executing the custom operation.",
								MarkdownDescription: "Specifies the name of the ServiceAccount to be used for executing the custom operation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"enqueue_on_force": schema.BoolAttribute{
						Description:         "Indicates whether opsRequest should continue to queue when 'force' is set to true.",
						MarkdownDescription: "Indicates whether opsRequest should continue to queue when 'force' is set to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"expose": schema.ListNestedAttribute{
						Description:         "Lists Expose objects, each specifying a Component and its services to be exposed.",
						MarkdownDescription: "Lists Expose objects, each specifying a Component and its services to be exposed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"services": schema.ListNestedAttribute{
									Description:         "Specifies a list of OpsService. When an OpsService is exposed, a corresponding ClusterService will be added to 'cluster.spec.services'. On the other hand, when an OpsService is unexposed, the corresponding ClusterService will be removed from 'cluster.spec.services'. Note: If 'componentName' is not specified, the 'ports' and 'selector' fields must be provided in each OpsService definition.",
									MarkdownDescription: "Specifies a list of OpsService. When an OpsService is exposed, a corresponding ClusterService will be added to 'cluster.spec.services'. On the other hand, when an OpsService is unexposed, the corresponding ClusterService will be removed from 'cluster.spec.services'. Note: If 'componentName' is not specified, the 'ports' and 'selector' fields must be provided in each OpsService definition.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Contains cloud provider related parameters if ServiceType is LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												MarkdownDescription: "Contains cloud provider related parameters if ServiceType is LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_families": schema.ListAttribute{
												Description:         "A list of IP families (e.g., IPv4, IPv6) assigned to this Service. Usually assigned automatically based on the cluster configuration and the 'ipFamilyPolicy' field. If specified manually, the requested IP family must be available in the cluster and allowed by the 'ipFamilyPolicy'. If the requested IP family is not available or not allowed, the Service creation will fail. Valid values: - 'IPv4' - 'IPv6' This field may hold a maximum of two entries (dual-stack families, in either order). Common combinations of 'ipFamilies' and 'ipFamilyPolicy' are: - ipFamilies=[] + ipFamilyPolicy='PreferDualStack' : The Service prefers dual-stack but can fall back to single-stack if the cluster does not support dual-stack. The IP family is automatically assigned based on the cluster configuration. - ipFamilies=['IPV4','IPV6'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV4. - ipFamilies=['IPV6','IPV4'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV6. - ipFamilies=['IPV4'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv4 only. - ipFamilies=['IPV6'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv6 only.",
												MarkdownDescription: "A list of IP families (e.g., IPv4, IPv6) assigned to this Service. Usually assigned automatically based on the cluster configuration and the 'ipFamilyPolicy' field. If specified manually, the requested IP family must be available in the cluster and allowed by the 'ipFamilyPolicy'. If the requested IP family is not available or not allowed, the Service creation will fail. Valid values: - 'IPv4' - 'IPv6' This field may hold a maximum of two entries (dual-stack families, in either order). Common combinations of 'ipFamilies' and 'ipFamilyPolicy' are: - ipFamilies=[] + ipFamilyPolicy='PreferDualStack' : The Service prefers dual-stack but can fall back to single-stack if the cluster does not support dual-stack. The IP family is automatically assigned based on the cluster configuration. - ipFamilies=['IPV4','IPV6'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV4. - ipFamilies=['IPV6','IPV4'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV6. - ipFamilies=['IPV4'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv4 only. - ipFamilies=['IPV6'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv6 only.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies whether the Service should use a single IP family (SingleStack) or two IP families (DualStack). Possible values: - 'SingleStack' (default) : The Service uses a single IP family. If no value is provided, IPFamilyPolicy defaults to SingleStack. - 'PreferDualStack' : The Service prefers to use two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. - 'RequiredDualStack' : The Service requires two IP families on dual-stack configured clusters. If the cluster is not configured for dual-stack, the Service creation fails.",
												MarkdownDescription: "Specifies whether the Service should use a single IP family (SingleStack) or two IP families (DualStack). Possible values: - 'SingleStack' (default) : The Service uses a single IP family. If no value is provided, IPFamilyPolicy defaults to SingleStack. - 'PreferDualStack' : The Service prefers to use two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. - 'RequiredDualStack' : The Service requires two IP families on dual-stack configured clusters. If the cluster is not configured for dual-stack, the Service creation fails.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Specifies the name of the Service. This name is used to set 'clusterService.name'. Note: This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the Service. This name is used to set 'clusterService.name'. Note: This field cannot be updated.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"pod_selector": schema.MapAttribute{
												Description:         "Routes service traffic to pods with matching label keys and values. If specified, the service will only be exposed to pods matching the selector. Note: If the component has roles, at least one of 'roleSelector' or 'podSelector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												MarkdownDescription: "Routes service traffic to pods with matching label keys and values. If specified, the service will only be exposed to pods matching the selector. Note: If the component has roles, at least one of 'roleSelector' or 'podSelector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ports": schema.ListNestedAttribute{
												Description:         "Specifies Port definitions that are to be exposed by a ClusterService. If not specified, the Port definitions from non-NodePort and non-LoadBalancer type ComponentService defined in the ComponentDefinition ('componentDefinition.spec.services') will be used. If no matching ComponentService is found, the expose operation will fail. More info: https://kubernetes.io/docs/concepts/services-networking/service/#field-spec-ports",
												MarkdownDescription: "Specifies Port definitions that are to be exposed by a ClusterService. If not specified, the Port definitions from non-NodePort and non-LoadBalancer type ComponentService defined in the ComponentDefinition ('componentDefinition.spec.services') will be used. If no matching ComponentService is found, the expose operation will fail. More info: https://kubernetes.io/docs/concepts/services-networking/service/#field-spec-ports",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"app_protocol": schema.StringAttribute{
															Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_port": schema.Int64Attribute{
															Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The port that will be exposed by this service.",
															MarkdownDescription: "The port that will be exposed by this service.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"protocol": schema.StringAttribute{
															Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
															MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

											"role_selector": schema.StringAttribute{
												Description:         "Specifies a role to target with the service. If specified, the service will only be exposed to pods with the matching role. Note: If the component has roles, at least one of 'roleSelector' or 'podSelector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												MarkdownDescription: "Specifies a role to target with the service. If specified, the service will only be exposed to pods with the matching role. Note: If the component has roles, at least one of 'roleSelector' or 'podSelector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_type": schema.StringAttribute{
												Description:         "Determines how the Service is exposed. Defaults to 'ClusterIP'. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'. - 'ClusterIP': allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort': builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer': builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for the expose operation. For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												MarkdownDescription: "Determines how the Service is exposed. Defaults to 'ClusterIP'. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'. - 'ClusterIP': allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort': builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer': builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for the expose operation. For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
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

								"switch": schema.StringAttribute{
									Description:         "Indicates whether the services will be exposed. 'Enable' exposes the services. while 'Disable' removes the exposed Service.",
									MarkdownDescription: "Indicates whether the services will be exposed. 'Enable' exposes the services. while 'Disable' removes the exposed Service.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Enable", "Disable"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"force": schema.BoolAttribute{
						Description:         "Instructs the system to bypass pre-checks (including cluster state checks and customized pre-conditions hooks) and immediately execute the opsRequest, except for the opsRequest of 'Start' type, which will still undergo pre-checks even if 'force' is true. This is useful for concurrent execution of 'VerticalScaling' and 'HorizontalScaling' opsRequests. By setting 'force' to true, you can bypass the default checks and demand these opsRequests to run simultaneously. Note: Once set, the 'force' field is immutable and cannot be updated.",
						MarkdownDescription: "Instructs the system to bypass pre-checks (including cluster state checks and customized pre-conditions hooks) and immediately execute the opsRequest, except for the opsRequest of 'Start' type, which will still undergo pre-checks even if 'force' is true. This is useful for concurrent execution of 'VerticalScaling' and 'HorizontalScaling' opsRequests. By setting 'force' to true, you can bypass the default checks and demand these opsRequests to run simultaneously. Note: Once set, the 'force' field is immutable and cannot be updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"horizontal_scaling": schema.ListNestedAttribute{
						Description:         "Lists HorizontalScaling objects, each specifying scaling requirements for a Component, including desired replica changes, configurations for new instances, modifications for existing instances, and take offline/online the specified instances.",
						MarkdownDescription: "Lists HorizontalScaling objects, each specifying scaling requirements for a Component, including desired replica changes, configurations for new instances, modifications for existing instances, and take offline/online the specified instances.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"scale_in": schema.SingleNestedAttribute{
									Description:         "Specifies the replica changes for scaling in components and instance templates, and takes specified instances offline. Can be used in conjunction with the 'scaleOut' operation. Note: Any configuration that creates instances is considered invalid.",
									MarkdownDescription: "Specifies the replica changes for scaling in components and instance templates, and takes specified instances offline. Can be used in conjunction with the 'scaleOut' operation. Note: Any configuration that creates instances is considered invalid.",
									Attributes: map[string]schema.Attribute{
										"instances": schema.ListNestedAttribute{
											Description:         "Modifies the desired replicas count for existing InstanceTemplate. if the inst",
											MarkdownDescription: "Modifies the desired replicas count for existing InstanceTemplate. if the inst",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Specifies the name of the instance template.",
														MarkdownDescription: "Specifies the name of the instance template.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"replica_changes": schema.Int64Attribute{
														Description:         "Specifies the replica changes for the instance template.",
														MarkdownDescription: "Specifies the replica changes for the instance template.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"online_instances_to_offline": schema.ListAttribute{
											Description:         "Specifies the instance names that need to be taken offline.",
											MarkdownDescription: "Specifies the instance names that need to be taken offline.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replica_changes": schema.Int64Attribute{
											Description:         "Specifies the replica changes for the component.",
											MarkdownDescription: "Specifies the replica changes for the component.",
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

								"scale_out": schema.SingleNestedAttribute{
									Description:         "Specifies the replica changes for scaling out components and instance templates, and brings offline instances back online. Can be used in conjunction with the 'scaleIn' operation. Note: Any configuration that deletes instances is considered invalid.",
									MarkdownDescription: "Specifies the replica changes for scaling out components and instance templates, and brings offline instances back online. Can be used in conjunction with the 'scaleIn' operation. Note: Any configuration that deletes instances is considered invalid.",
									Attributes: map[string]schema.Attribute{
										"instances": schema.ListNestedAttribute{
											Description:         "Modifies the desired replicas count for existing InstanceTemplate. if the inst",
											MarkdownDescription: "Modifies the desired replicas count for existing InstanceTemplate. if the inst",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Specifies the name of the instance template.",
														MarkdownDescription: "Specifies the name of the instance template.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"replica_changes": schema.Int64Attribute{
														Description:         "Specifies the replica changes for the instance template.",
														MarkdownDescription: "Specifies the replica changes for the instance template.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"new_instances": schema.ListNestedAttribute{
											Description:         "Defines the configuration for new instances added during scaling, including resource requirements, labels, annotations, etc. New instances are created based on the provided instance templates.",
											MarkdownDescription: "Defines the configuration for new instances added during scaling, including resource requirements, labels, annotations, etc. New instances are created based on the provided instance templates.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Specifies a map of key-value pairs to be merged into the Pod's existing annotations. Existing keys will have their values overwritten, while new keys will be added to the annotations.",
														MarkdownDescription: "Specifies a map of key-value pairs to be merged into the Pod's existing annotations. Existing keys will have their values overwritten, while new keys will be added to the annotations.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Defines Env to override. Add new or override existing envs.",
														MarkdownDescription: "Defines Env to override. Add new or override existing envs.",
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
																					Description:         "The key of the secret to select from. Must be a valid secret key.",
																					MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
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
														Description:         "Specifies a map of key-value pairs that will be merged into the Pod's existing labels. Values for existing keys will be overwritten, and new keys will be added.",
														MarkdownDescription: "Specifies a map of key-value pairs that will be merged into the Pod's existing labels. Values for existing keys will be overwritten, and new keys will be added.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name specifies the unique name of the instance Pod created using this InstanceTemplate. This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0. The specified name overrides any default naming conventions or patterns.",
														MarkdownDescription: "Name specifies the unique name of the instance Pod created using this InstanceTemplate. This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0. The specified name overrides any default naming conventions or patterns.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(54),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
														},
													},

													"ordinals": schema.SingleNestedAttribute{
														Description:         "Specifies the desired Ordinals of this InstanceTemplate. The Ordinals used to specify the ordinal of the instance (pod) names to be generated under this InstanceTemplate. For example, if Ordinals is {ranges: [{start: 0, end: 1}], discrete: [7]}, then the instance names generated under this InstanceTemplate would be $(cluster.name)-$(component.name)-$(template.name)-0、$(cluster.name)-$(component.name)-$(template.name)-1 and $(cluster.name)-$(component.name)-$(template.name)-7",
														MarkdownDescription: "Specifies the desired Ordinals of this InstanceTemplate. The Ordinals used to specify the ordinal of the instance (pod) names to be generated under this InstanceTemplate. For example, if Ordinals is {ranges: [{start: 0, end: 1}], discrete: [7]}, then the instance names generated under this InstanceTemplate would be $(cluster.name)-$(component.name)-$(template.name)-0、$(cluster.name)-$(component.name)-$(template.name)-1 and $(cluster.name)-$(component.name)-$(template.name)-7",
														Attributes: map[string]schema.Attribute{
															"discrete": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"ranges": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"end": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"start": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
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

													"replicas": schema.Int64Attribute{
														Description:         "Specifies the number of instances (Pods) to create from this InstanceTemplate. This field allows setting how many replicated instances of the Component, with the specific overrides in the InstanceTemplate, are created. The default value is 1. A value of 0 disables instance creation.",
														MarkdownDescription: "Specifies the number of instances (Pods) to create from this InstanceTemplate. This field allows setting how many replicated instances of the Component, with the specific overrides in the InstanceTemplate, are created. The default value is 1. A value of 0 disables instance creation.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Specifies an override for the resource requirements of the first container in the Pod. This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
														MarkdownDescription: "Specifies an override for the resource requirements of the first container in the Pod. This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
														Attributes: map[string]schema.Attribute{
															"claims": schema.ListNestedAttribute{
																Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
																MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
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

													"scheduling_policy": schema.SingleNestedAttribute{
														Description:         "Specifies the scheduling policy for the instance. If defined, it will overwrite the scheduling policy defined in ClusterSpec and/or ClusterComponentSpec.",
														MarkdownDescription: "Specifies the scheduling policy for the instance. If defined, it will overwrite the scheduling policy defined in ClusterSpec and/or ClusterComponentSpec.",
														Attributes: map[string]schema.Attribute{
															"affinity": schema.SingleNestedAttribute{
																Description:         "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
																MarkdownDescription: "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
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
																									Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																									MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
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
																									Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"mismatch_label_keys": schema.ListAttribute{
																									Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
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
																							Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																							MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
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
																							Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"mismatch_label_keys": schema.ListAttribute{
																							Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
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
																									Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																									MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
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
																									Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"mismatch_label_keys": schema.ListAttribute{
																									Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
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
																							Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																							MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
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
																							Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. Also, MatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"mismatch_label_keys": schema.ListAttribute{
																							Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector. Also, MismatchLabelKeys cannot be set when LabelSelector isn't set. This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"node_name": schema.StringAttribute{
																Description:         "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
																MarkdownDescription: "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"node_selector": schema.MapAttribute{
																Description:         "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																MarkdownDescription: "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scheduler_name": schema.StringAttribute{
																Description:         "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
																MarkdownDescription: "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tolerations": schema.ListNestedAttribute{
																Description:         "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'. - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint. Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
																MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'. - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint. Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
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
																Description:         "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
																MarkdownDescription: "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
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
																			Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector. This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector. This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"max_skew": schema.Int64Attribute{
																			Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | | P P | P P | P | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
																			MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | | P P | P P | P | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"min_domains": schema.Int64Attribute{
																			Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule. For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | | P P | P P | P P | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew. This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
																			MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule. For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | | P P | P P | P P | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew. This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"node_affinity_policy": schema.StringAttribute{
																			Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations. If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																			MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations. If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"node_taints_policy": schema.StringAttribute{
																			Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included. If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																			MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included. If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
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
																			Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P | P | P | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
																			MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P | P | P | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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

													"volume_claim_templates": schema.ListNestedAttribute{
														Description:         "Specifies an override for the storage requirements of the instances.",
														MarkdownDescription: "Specifies an override for the storage requirements of the instances.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"annotations": schema.MapAttribute{
																	Description:         "Specifies the annotations for the PVC of the volume.",
																	MarkdownDescription: "Specifies the annotations for the PVC of the volume.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"labels": schema.MapAttribute{
																	Description:         "Specifies the labels for the PVC of the volume.",
																	MarkdownDescription: "Specifies the labels for the PVC of the volume.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Refers to the name of a volumeMount defined in either: - 'componentDefinition.spec.runtime.containers[*].volumeMounts' The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
																	MarkdownDescription: "Refers to the name of a volumeMount defined in either: - 'componentDefinition.spec.runtime.containers[*].volumeMounts' The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"persistent_volume_claim_name": schema.StringAttribute{
																	Description:         "Specifies the prefix of the PVC name for the volume. For each replica, the final name of the PVC will be in format: <persistentVolumeClaimName>-<ordinal>",
																	MarkdownDescription: "Specifies the prefix of the PVC name for the volume. For each replica, the final name of the PVC will be in format: <persistentVolumeClaimName>-<ordinal>",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"spec": schema.SingleNestedAttribute{
																	Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.",
																	MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.",
																	Attributes: map[string]schema.Attribute{
																		"access_modes": schema.ListAttribute{
																			Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																			MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"data_source": schema.SingleNestedAttribute{
																			Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																			MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																			Attributes: map[string]schema.Attribute{
																				"api_group": schema.StringAttribute{
																					Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"kind": schema.StringAttribute{
																					Description:         "Kind is the type of resource being referenced",
																					MarkdownDescription: "Kind is the type of resource being referenced",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name is the name of resource being referenced",
																					MarkdownDescription: "Name is the name of resource being referenced",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"data_source_ref": schema.SingleNestedAttribute{
																			Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																			MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																			Attributes: map[string]schema.Attribute{
																				"api_group": schema.StringAttribute{
																					Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"kind": schema.StringAttribute{
																					Description:         "Kind is the type of resource being referenced",
																					MarkdownDescription: "Kind is the type of resource being referenced",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name is the name of resource being referenced",
																					MarkdownDescription: "Name is the name of resource being referenced",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"namespace": schema.StringAttribute{
																					Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"resources": schema.SingleNestedAttribute{
																			Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																			MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																			Attributes: map[string]schema.Attribute{
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

																		"selector": schema.SingleNestedAttribute{
																			Description:         "selector is a label query over volumes to consider for binding.",
																			MarkdownDescription: "selector is a label query over volumes to consider for binding.",
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

																		"storage_class_name": schema.StringAttribute{
																			Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																			MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_attributes_class_name": schema.StringAttribute{
																			Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass (Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
																			MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass (Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_mode": schema.StringAttribute{
																			Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																			MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_name": schema.StringAttribute{
																			Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																			MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"offline_instances_to_online": schema.ListAttribute{
											Description:         "Specifies the instances in the offline list to bring back online.",
											MarkdownDescription: "Specifies the instances in the offline list to bring back online.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replica_changes": schema.Int64Attribute{
											Description:         "Specifies the replica changes for the component.",
											MarkdownDescription: "Specifies the replica changes for the component.",
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

								"shards": schema.Int64Attribute{
									Description:         "Specifies the desired number of shards for the component. This parameter is mutually exclusive with other parameters.",
									MarkdownDescription: "Specifies the desired number of shards for the component. This parameter is mutually exclusive with other parameters.",
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

					"pre_condition_deadline_seconds": schema.Int64Attribute{
						Description:         "Specifies the maximum time in seconds that the OpsRequest will wait for its pre-conditions to be met before it aborts the operation. If set to 0 (default), pre-conditions must be satisfied immediately for the OpsRequest to proceed.",
						MarkdownDescription: "Specifies the maximum time in seconds that the OpsRequest will wait for its pre-conditions to be met before it aborts the operation. If set to 0 (default), pre-conditions must be satisfied immediately for the OpsRequest to proceed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rebuild_from": schema.ListNestedAttribute{
						Description:         "Specifies the parameters to rebuild some instances. Rebuilding an instance involves restoring its data from a backup or another database replica. The instances being rebuilt usually serve as standby in the cluster. Hence, rebuilding instances is often also referred to as 'standby reconstruction'.",
						MarkdownDescription: "Specifies the parameters to rebuild some instances. Rebuilding an instance involves restoring its data from a backup or another database replica. The instances being rebuilt usually serve as standby in the cluster. Hence, rebuilding instances is often also referred to as 'standby reconstruction'.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup_name": schema.StringAttribute{
									Description:         "Indicates the name of the Backup custom resource from which to recover the instance. Defaults to an empty PersistentVolume if unspecified. Note: - Only full physical backups are supported for multi-replica Components (e.g., 'xtrabackup' for MySQL). - Logical backups (e.g., 'mysqldump' for MySQL) are unsupported in the current version.",
									MarkdownDescription: "Indicates the name of the Backup custom resource from which to recover the instance. Defaults to an empty PersistentVolume if unspecified. Note: - Only full physical backups are supported for multi-replica Components (e.g., 'xtrabackup' for MySQL). - Logical backups (e.g., 'mysqldump' for MySQL) are unsupported in the current version.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"in_place": schema.BoolAttribute{
									Description:         "When it is set to true, the instance will be rebuilt in-place. If false, a new pod will be created. Once the new pod is ready to serve, the instance that require rebuilding will be taken offline.",
									MarkdownDescription: "When it is set to true, the instance will be rebuilt in-place. If false, a new pod will be created. Once the new pod is ready to serve, the instance that require rebuilding will be taken offline.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instances": schema.ListNestedAttribute{
									Description:         "Specifies the instances (Pods) that need to be rebuilt, typically operating as standbys.",
									MarkdownDescription: "Specifies the instances (Pods) that need to be rebuilt, typically operating as standbys.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Pod name of the instance.",
												MarkdownDescription: "Pod name of the instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"target_node_name": schema.StringAttribute{
												Description:         "The instance will rebuild on the specified node. If not set, it will rebuild on a random node.",
												MarkdownDescription: "The instance will rebuild on the specified node. If not set, it will rebuild on a random node.",
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

								"restore_env": schema.MapAttribute{
									Description:         "Defines container environment variables for the restore process. merged with the ones specified in the Backup and ActionSet resources. Merge priority: Restore env > Backup env > ActionSet env. Purpose: Some databases require different configurations when being restored as a standby compared to being restored as a primary. For example, when restoring MySQL as a replica, you need to set 'skip_slave_start='ON'' for 5.7 or 'skip_replica_start='ON'' for 8.0. Allowing environment variables to be passed in makes it more convenient to control these behavioral differences during the restore process.",
									MarkdownDescription: "Defines container environment variables for the restore process. merged with the ones specified in the Backup and ActionSet resources. Merge priority: Restore env > Backup env > ActionSet env. Purpose: Some databases require different configurations when being restored as a standby compared to being restored as a primary. For example, when restoring MySQL as a replica, you need to set 'skip_slave_start='ON'' for 5.7 or 'skip_replica_start='ON'' for 8.0. Allowing environment variables to be passed in makes it more convenient to control these behavioral differences during the restore process.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_backup_target_name": schema.StringAttribute{
									Description:         "When multiple source targets exist of the backup, you must specify the source target to restore.",
									MarkdownDescription: "When multiple source targets exist of the backup, you must specify the source target to restore.",
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

					"reconfigures": schema.ListNestedAttribute{
						Description:         "Lists Reconfigure objects, each specifying a Component and its configuration updates.",
						MarkdownDescription: "Lists Reconfigure objects, each specifying a Component and its configuration updates.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"parameters": schema.ListNestedAttribute{
									Description:         "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file.",
									MarkdownDescription: "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Represents the name of the parameter that is to be updated.",
												MarkdownDescription: "Represents the name of the parameter that is to be updated.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"value": schema.StringAttribute{
												Description:         "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
												MarkdownDescription: "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"restart": schema.ListNestedAttribute{
						Description:         "Lists Components to be restarted.",
						MarkdownDescription: "Lists Components to be restarted.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
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

					"restore": schema.SingleNestedAttribute{
						Description:         "Specifies the parameters to restore a Cluster. Note that this restore operation will roll back cluster services.",
						MarkdownDescription: "Specifies the parameters to restore a Cluster. Note that this restore operation will roll back cluster services.",
						Attributes: map[string]schema.Attribute{
							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the Backup custom resource.",
								MarkdownDescription: "Specifies the name of the Backup custom resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"backup_namespace": schema.StringAttribute{
								Description:         "Specifies the namespace of the backup custom resource. If not specified, the namespace of the opsRequest will be used.",
								MarkdownDescription: "Specifies the namespace of the backup custom resource. If not specified, the namespace of the opsRequest will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"defer_post_ready_until_cluster_running": schema.BoolAttribute{
								Description:         "Controls the timing of PostReady actions during the recovery process. If false (default), PostReady actions execute when the Component reaches the 'Running' state. If true, PostReady actions are delayed until the entire Cluster is 'Running,' ensuring the cluster's overall stability before proceeding. This setting is useful for coordinating PostReady operations across the Cluster for optimal cluster conditions.",
								MarkdownDescription: "Controls the timing of PostReady actions during the recovery process. If false (default), PostReady actions execute when the Component reaches the 'Running' state. If true, PostReady actions are delayed until the entire Cluster is 'Running,' ensuring the cluster's overall stability before proceeding. This setting is useful for coordinating PostReady operations across the Cluster for optimal cluster conditions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env": schema.MapAttribute{
								Description:         "Specifies a list of environment variables to be set in the container.",
								MarkdownDescription: "Specifies a list of environment variables to be set in the container.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parameters": schema.ListNestedAttribute{
								Description:         "Specifies a list of name-value pairs representing parameters and their corresponding values. Parameters match the schema specified in the 'actionset.spec.parametersSchema'",
								MarkdownDescription: "Specifies a list of name-value pairs representing parameters and their corresponding values. Parameters match the schema specified in the 'actionset.spec.parametersSchema'",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Represents the name of the parameter.",
											MarkdownDescription: "Represents the name of the parameter.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Represents the parameter values.",
											MarkdownDescription: "Represents the parameter values.",
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

							"restore_point_in_time": schema.StringAttribute{
								Description:         "Specifies the point in time to which the restore should be performed. Supported time formats: - RFC3339 format, e.g. '2023-11-25T18:52:53Z' - A human-readable date-time format, e.g. 'Jul 25,2023 18:52:53 UTC+0800'",
								MarkdownDescription: "Specifies the point in time to which the restore should be performed. Supported time formats: - RFC3339 format, e.g. '2023-11-25T18:52:53Z' - A human-readable date-time format, e.g. 'Jul 25,2023 18:52:53 UTC+0800'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_restore_policy": schema.StringAttribute{
								Description:         "Specifies the policy for restoring volume claims of a Component's Pods. It determines whether the volume claims should be restored sequentially (one by one) or in parallel (all at once). Support values: - 'Serial' - 'Parallel'",
								MarkdownDescription: "Specifies the policy for restoring volume claims of a Component's Pods. It determines whether the volume claims should be restored sequentially (one by one) or in parallel (all at once). Support values: - 'Serial' - 'Parallel'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Serial", "Parallel"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"start": schema.ListNestedAttribute{
						Description:         "Lists Components to be started. If empty, all components will be started.",
						MarkdownDescription: "Lists Components to be started. If empty, all components will be started.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
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

					"stop": schema.ListNestedAttribute{
						Description:         "Lists Components to be stopped. If empty, all components will be stopped.",
						MarkdownDescription: "Lists Components to be stopped. If empty, all components will be stopped.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
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

					"switchover": schema.ListNestedAttribute{
						Description:         "Lists Switchover objects, each specifying a Component to perform the switchover operation.",
						MarkdownDescription: "Lists Switchover objects, each specifying a Component to perform the switchover operation.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"candidate_name": schema.StringAttribute{
									Description:         "If CandidateName is specified, the role will be transferred to this instance. The name must match one of the pods in the component. Refer to ComponentDefinition's Swtichover lifecycle action for more details.",
									MarkdownDescription: "If CandidateName is specified, the role will be transferred to this instance. The name must match one of the pods in the component. Refer to ComponentDefinition's Swtichover lifecycle action for more details.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec.",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"component_object_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component object.",
									MarkdownDescription: "Specifies the name of the Component object.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"instance_name": schema.StringAttribute{
									Description:         "Specifies the instance whose role will be transferred. A typical usage is to transfer the leader role in a consensus system.",
									MarkdownDescription: "Specifies the instance whose role will be transferred. A typical usage is to transfer the leader role in a consensus system.",
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

					"timeout_seconds": schema.Int64Attribute{
						Description:         "Specifies the maximum duration (in seconds) that an opsRequest is allowed to run. If the opsRequest runs longer than this duration, its phase will be marked as Aborted. If this value is not set or set to 0, the timeout will be ignored and the opsRequest will run indefinitely.",
						MarkdownDescription: "Specifies the maximum duration (in seconds) that an opsRequest is allowed to run. If the opsRequest runs longer than this duration, its phase will be marked as Aborted. If this value is not set or set to 0, the timeout will be ignored and the opsRequest will run indefinitely.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl_seconds_after_succeed": schema.Int64Attribute{
						Description:         "Specifies the duration in seconds that an OpsRequest will remain in the system after successfully completing (when 'opsRequest.status.phase' is 'Succeed') before automatic deletion.",
						MarkdownDescription: "Specifies the duration in seconds that an OpsRequest will remain in the system after successfully completing (when 'opsRequest.status.phase' is 'Succeed') before automatic deletion.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl_seconds_after_unsuccessful_completion": schema.Int64Attribute{
						Description:         "Specifies the duration in seconds that an OpsRequest will remain in the system after completion for any phase other than 'Succeed' (e.g., 'Failed', 'Cancelled', 'Aborted') before automatic deletion.",
						MarkdownDescription: "Specifies the duration in seconds that an OpsRequest will remain in the system after completion for any phase other than 'Succeed' (e.g., 'Failed', 'Cancelled', 'Aborted') before automatic deletion.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Specifies the type of this operation. Supported types include 'Start', 'Stop', 'Restart', 'Switchover', 'VerticalScaling', 'HorizontalScaling', 'VolumeExpansion', 'Reconfiguring', 'Upgrade', 'Backup', 'Restore', 'Expose', 'RebuildInstance', 'Custom'. Note: This field is immutable once set.",
						MarkdownDescription: "Specifies the type of this operation. Supported types include 'Start', 'Stop', 'Restart', 'Switchover', 'VerticalScaling', 'HorizontalScaling', 'VolumeExpansion', 'Reconfiguring', 'Upgrade', 'Backup', 'Restore', 'Expose', 'RebuildInstance', 'Custom'. Note: This field is immutable once set.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Upgrade", "VerticalScaling", "VolumeExpansion", "HorizontalScaling", "Restart", "Reconfiguring", "Start", "Stop", "Expose", "Switchover", "Backup", "Restore", "RebuildInstance", "Custom"),
						},
					},

					"upgrade": schema.SingleNestedAttribute{
						Description:         "Specifies the desired new version of the Cluster. Note: This field is immutable once set.",
						MarkdownDescription: "Specifies the desired new version of the Cluster. Note: This field is immutable once set.",
						Attributes: map[string]schema.Attribute{
							"components": schema.ListNestedAttribute{
								Description:         "Lists components to be upgrade based on desired ComponentDefinition and ServiceVersion. From the perspective of cluster API, the reasonable combinations should be: 1. (comp-def, service-ver) - upgrade to the specified service version and component definition, the user takes the responsibility to ensure that they are compatible. 2. ('', service-ver) - upgrade to the specified service version, let the operator choose the latest compatible component definition. 3. (comp-def, '') - upgrade to the specified component definition, let the operator choose the latest compatible service version. 4. ('', '') - upgrade to the latest service version and component definition, the operator will ensure the compatibility between the selected versions.",
								MarkdownDescription: "Lists components to be upgrade based on desired ComponentDefinition and ServiceVersion. From the perspective of cluster API, the reasonable combinations should be: 1. (comp-def, service-ver) - upgrade to the specified service version and component definition, the user takes the responsibility to ensure that they are compatible. 2. ('', service-ver) - upgrade to the specified service version, let the operator choose the latest compatible component definition. 3. (comp-def, '') - upgrade to the specified component definition, let the operator choose the latest compatible service version. 4. ('', '') - upgrade to the latest service version and component definition, the operator will ensure the compatibility between the selected versions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"component_definition_name": schema.StringAttribute{
											Description:         "Specifies the name of the ComponentDefinition, only exact matches are supported.",
											MarkdownDescription: "Specifies the name of the ComponentDefinition, only exact matches are supported.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(64),
											},
										},

										"component_name": schema.StringAttribute{
											Description:         "Specifies the name of the Component as defined in the cluster.spec",
											MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"service_version": schema.StringAttribute{
											Description:         "Specifies the version of the Service expected to be provisioned by this Component. Referring to the ServiceVersion defined by the ComponentDefinition and ComponentVersion. And ServiceVersion in ClusterComponentSpec is optional, when no version is specified, use the latest available version in ComponentVersion.",
											MarkdownDescription: "Specifies the version of the Service expected to be provisioned by this Component. Referring to the ServiceVersion defined by the ComponentDefinition and ComponentVersion. And ServiceVersion in ClusterComponentSpec is optional, when no version is specified, use the latest available version in ComponentVersion.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(32),
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

					"vertical_scaling": schema.ListAttribute{
						Description:         "Lists VerticalScaling objects, each specifying a component and its desired compute resources for vertical scaling.",
						MarkdownDescription: "Lists VerticalScaling objects, each specifying a component and its desired compute resources for vertical scaling.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_expansion": schema.ListNestedAttribute{
						Description:         "Lists VolumeExpansion objects, each specifying a component and its corresponding volumeClaimTemplates that requires storage expansion.",
						MarkdownDescription: "Lists VolumeExpansion objects, each specifying a component and its corresponding volumeClaimTemplates that requires storage expansion.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component as defined in the cluster.spec",
									MarkdownDescription: "Specifies the name of the Component as defined in the cluster.spec",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"volume_claim_templates": schema.ListNestedAttribute{
									Description:         "Specifies a list of OpsRequestVolumeClaimTemplate objects, defining the volumeClaimTemplates that are used to expand the storage and the desired storage size for each one.",
									MarkdownDescription: "Specifies a list of OpsRequestVolumeClaimTemplate objects, defining the volumeClaimTemplates that are used to expand the storage and the desired storage size for each one.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Specify the name of the volumeClaimTemplate in the Component. The specified name must match one of the volumeClaimTemplates defined in the 'clusterComponentSpec.volumeClaimTemplates' field.",
												MarkdownDescription: "Specify the name of the volumeClaimTemplate in the Component. The specified name must match one of the volumeClaimTemplates defined in the 'clusterComponentSpec.volumeClaimTemplates' field.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"storage": schema.StringAttribute{
												Description:         "Specifies the desired storage size for the volume.",
												MarkdownDescription: "Specifies the desired storage size for the volume.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
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
	}
}

func (r *OperationsKubeblocksIoOpsRequestV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operations_kubeblocks_io_ops_request_v1alpha1_manifest")

	var model OperationsKubeblocksIoOpsRequestV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operations.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("OpsRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
