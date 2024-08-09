/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dataprotection_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &DataprotectionKubeblocksIoRestoreV1Alpha1Manifest{}
)

func NewDataprotectionKubeblocksIoRestoreV1Alpha1Manifest() datasource.DataSource {
	return &DataprotectionKubeblocksIoRestoreV1Alpha1Manifest{}
}

type DataprotectionKubeblocksIoRestoreV1Alpha1Manifest struct{}

type DataprotectionKubeblocksIoRestoreV1Alpha1ManifestData struct {
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
		BackoffLimit *int64 `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
		Backup       *struct {
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			Namespace        *string `tfsdk:"namespace" json:"namespace,omitempty"`
			SourceTargetName *string `tfsdk:"source_target_name" json:"sourceTargetName,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		ContainerResources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"container_resources" json:"containerResources,omitempty"`
		Env               *map[string]string `tfsdk:"env" json:"env,omitempty"`
		PrepareDataConfig *struct {
			DataSourceRef *struct {
				MountPath    *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				VolumeSource *string `tfsdk:"volume_source" json:"volumeSource,omitempty"`
			} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
			RequiredPolicyForAllPodSelection *struct {
				DataRestorePolicy *string `tfsdk:"data_restore_policy" json:"dataRestorePolicy,omitempty"`
				SourceOfOneToMany *struct {
					TargetPodName *string `tfsdk:"target_pod_name" json:"targetPodName,omitempty"`
				} `tfsdk:"source_of_one_to_many" json:"sourceOfOneToMany,omitempty"`
			} `tfsdk:"required_policy_for_all_pod_selection" json:"requiredPolicyForAllPodSelection,omitempty"`
			SchedulingSpec *struct {
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
			} `tfsdk:"scheduling_spec" json:"schedulingSpec,omitempty"`
			VolumeClaimRestorePolicy *string `tfsdk:"volume_claim_restore_policy" json:"volumeClaimRestorePolicy,omitempty"`
			VolumeClaims             *[]struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
					Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				MountPath       *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				VolumeClaimSpec *struct {
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
				} `tfsdk:"volume_claim_spec" json:"volumeClaimSpec,omitempty"`
				VolumeSource *string `tfsdk:"volume_source" json:"volumeSource,omitempty"`
			} `tfsdk:"volume_claims" json:"volumeClaims,omitempty"`
			VolumeClaimsTemplate *struct {
				Replicas      *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				StartingIndex *int64 `tfsdk:"starting_index" json:"startingIndex,omitempty"`
				Templates     *[]struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name        *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					MountPath       *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					VolumeClaimSpec *struct {
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
					} `tfsdk:"volume_claim_spec" json:"volumeClaimSpec,omitempty"`
					VolumeSource *string `tfsdk:"volume_source" json:"volumeSource,omitempty"`
				} `tfsdk:"templates" json:"templates,omitempty"`
			} `tfsdk:"volume_claims_template" json:"volumeClaimsTemplate,omitempty"`
		} `tfsdk:"prepare_data_config" json:"prepareDataConfig,omitempty"`
		ReadyConfig *struct {
			ConnectionCredential *struct {
				HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"connection_credential" json:"connectionCredential,omitempty"`
			ExecAction *struct {
				Target *struct {
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"exec_action" json:"execAction,omitempty"`
			JobAction *struct {
				RequiredPolicyForAllPodSelection *struct {
					DataRestorePolicy *string `tfsdk:"data_restore_policy" json:"dataRestorePolicy,omitempty"`
					SourceOfOneToMany *struct {
						TargetPodName *string `tfsdk:"target_pod_name" json:"targetPodName,omitempty"`
					} `tfsdk:"source_of_one_to_many" json:"sourceOfOneToMany,omitempty"`
				} `tfsdk:"required_policy_for_all_pod_selection" json:"requiredPolicyForAllPodSelection,omitempty"`
				Target *struct {
					PodSelector *struct {
						FallbackLabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"fallback_label_selector" json:"fallbackLabelSelector,omitempty"`
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						Strategy    *string            `tfsdk:"strategy" json:"strategy,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					VolumeMounts *[]struct {
						MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						Name             *string `tfsdk:"name" json:"name,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"job_action" json:"jobAction,omitempty"`
			ReadinessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" json:"command,omitempty"`
					Image   *string   `tfsdk:"image" json:"image,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		} `tfsdk:"ready_config" json:"readyConfig,omitempty"`
		Resources *struct {
			Included *[]struct {
				GroupResource *string `tfsdk:"group_resource" json:"groupResource,omitempty"`
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			} `tfsdk:"included" json:"included,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RestoreTime        *string `tfsdk:"restore_time" json:"restoreTime,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataprotectionKubeblocksIoRestoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dataprotection_kubeblocks_io_restore_v1alpha1_manifest"
}

func (r *DataprotectionKubeblocksIoRestoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Restore is the Schema for the restores API",
		MarkdownDescription: "Restore is the Schema for the restores API",
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
				Description:         "RestoreSpec defines the desired state of Restore",
				MarkdownDescription: "RestoreSpec defines the desired state of Restore",
				Attributes: map[string]schema.Attribute{
					"backoff_limit": schema.Int64Attribute{
						Description:         "Specifies the number of retries before marking the restore failed.",
						MarkdownDescription: "Specifies the number of retries before marking the restore failed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(10),
						},
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "Specifies the backup to be restored. The restore behavior is based on the backup type:1. Full: will be restored the full backup directly.2. Incremental: will be restored sequentially from the most recent full backup of this incremental backup.3. Differential: will be restored sequentially from the parent backup of the differential backup.4. Continuous: will find the most recent full backup at this time point and the continuous backups after it to restore.",
						MarkdownDescription: "Specifies the backup to be restored. The restore behavior is based on the backup type:1. Full: will be restored the full backup directly.2. Incremental: will be restored sequentially from the most recent full backup of this incremental backup.3. Differential: will be restored sequentially from the parent backup of the differential backup.4. Continuous: will find the most recent full backup at this time point and the continuous backups after it to restore.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Specifies the backup name.",
								MarkdownDescription: "Specifies the backup name.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Specifies the backup namespace.",
								MarkdownDescription: "Specifies the backup namespace.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"source_target_name": schema.StringAttribute{
								Description:         "Specifies the source target for restoration, identified by its name.",
								MarkdownDescription: "Specifies the source target for restoration, identified by its name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"container_resources": schema.SingleNestedAttribute{
						Description:         "Specifies the required resources of restore job's container.",
						MarkdownDescription: "Specifies the required resources of restore job's container.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
								Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"env": schema.MapAttribute{
						Description:         "List of environment variables to set in the container for restore. These will bemerged with the env of Backup and ActionSet.The priority of merging is as follows: 'Restore env > Backup env > ActionSet env'.",
						MarkdownDescription: "List of environment variables to set in the container for restore. These will bemerged with the env of Backup and ActionSet.The priority of merging is as follows: 'Restore env > Backup env > ActionSet env'.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prepare_data_config": schema.SingleNestedAttribute{
						Description:         "Configuration for the action of 'prepareData' phase, including the persistent volume claimsthat need to be restored and scheduling strategy of temporary recovery pod.",
						MarkdownDescription: "Configuration for the action of 'prepareData' phase, including the persistent volume claimsthat need to be restored and scheduling strategy of temporary recovery pod.",
						Attributes: map[string]schema.Attribute{
							"data_source_ref": schema.SingleNestedAttribute{
								Description:         "Specifies the configuration when using 'persistentVolumeClaim.spec.dataSourceRef' method for restoring.Describes the source volume of the backup targetVolumes and the mount path in the restoring container.",
								MarkdownDescription: "Specifies the configuration when using 'persistentVolumeClaim.spec.dataSourceRef' method for restoring.Describes the source volume of the backup targetVolumes and the mount path in the restoring container.",
								Attributes: map[string]schema.Attribute{
									"mount_path": schema.StringAttribute{
										Description:         "Specifies the path within the restoring container at which the volume should be mounted.",
										MarkdownDescription: "Specifies the path within the restoring container at which the volume should be mounted.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_source": schema.StringAttribute{
										Description:         "Describes the volume that will be restored from the specified volume of the backup targetVolumes.This is required if the backup uses a volume snapshot.",
										MarkdownDescription: "Describes the volume that will be restored from the specified volume of the backup targetVolumes.This is required if the backup uses a volume snapshot.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"required_policy_for_all_pod_selection": schema.SingleNestedAttribute{
								Description:         "Specifies the restore policy, which is required when the pod selection strategy for the source target is 'All'.This field is ignored if the pod selection strategy is 'Any'.optional",
								MarkdownDescription: "Specifies the restore policy, which is required when the pod selection strategy for the source target is 'All'.This field is ignored if the pod selection strategy is 'Any'.optional",
								Attributes: map[string]schema.Attribute{
									"data_restore_policy": schema.StringAttribute{
										Description:         "Specifies the data restore policy. Options include:- OneToMany: Enables restoration of all volumes from a single data copy of the original target instance.The 'sourceOfOneToMany' field must be set when using this policy.- OneToOne: Restricts data restoration such that each data piece can only be restored to a single target instance.This is the default policy. When the number of target instances specified for restoration surpasses the count of original backup target instances.",
										MarkdownDescription: "Specifies the data restore policy. Options include:- OneToMany: Enables restoration of all volumes from a single data copy of the original target instance.The 'sourceOfOneToMany' field must be set when using this policy.- OneToOne: Restricts data restoration such that each data piece can only be restored to a single target instance.This is the default policy. When the number of target instances specified for restoration surpasses the count of original backup target instances.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"source_of_one_to_many": schema.SingleNestedAttribute{
										Description:         "Specifies the name of the source target pod. This field is mandatory when the DataRestorePolicy is configured to 'OneToMany'.",
										MarkdownDescription: "Specifies the name of the source target pod. This field is mandatory when the DataRestorePolicy is configured to 'OneToMany'.",
										Attributes: map[string]schema.Attribute{
											"target_pod_name": schema.StringAttribute{
												Description:         "Specifies the name of the source target pod.",
												MarkdownDescription: "Specifies the name of the source target pod.",
												Required:            true,
												Optional:            false,
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

							"scheduling_spec": schema.SingleNestedAttribute{
								Description:         "Specifies the scheduling spec for the restoring pod.",
								MarkdownDescription: "Specifies the scheduling spec for the restoring pod.",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "Contains a group of affinity scheduling rules.Refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
										MarkdownDescription: "Contains a group of affinity scheduling rules.Refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.SingleNestedAttribute{
												Description:         "Describes node affinity scheduling rules for the pod.",
												MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
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
														Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
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
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																	MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																		"match_label_keys": schema.ListAttribute{
																			Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"mismatch_label_keys": schema.ListAttribute{
																			Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																	Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																	MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
														Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																"match_label_keys": schema.ListAttribute{
																	Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																	MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																		"match_label_keys": schema.ListAttribute{
																			Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"mismatch_label_keys": schema.ListAttribute{
																			Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																	Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																	MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
														Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																"match_label_keys": schema.ListAttribute{
																	Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
										Description:         "Specifies a request to schedule this pod onto a specific node. If it is non-empty,the scheduler simply schedules this pod onto that node, assuming that it fits resourcerequirements.",
										MarkdownDescription: "Specifies a request to schedule this pod onto a specific node. If it is non-empty,the scheduler simply schedules this pod onto that node, assuming that it fits resourcerequirements.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_selector": schema.MapAttribute{
										Description:         "Defines a selector which must be true for the pod to fit on a node.The selector must match a node's labels for the pod to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
										MarkdownDescription: "Defines a selector which must be true for the pod to fit on a node.The selector must match a node's labels for the pod to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheduler_name": schema.StringAttribute{
										Description:         "Specifies the scheduler to dispatch the pod.If not specified, the pod will be dispatched by the default scheduler.",
										MarkdownDescription: "Specifies the scheduler to dispatch the pod.If not specified, the pod will be dispatched by the default scheduler.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Specifies the tolerations for the restoring pod.",
										MarkdownDescription: "Specifies the tolerations for the restoring pod.",
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

									"topology_spread_constraints": schema.ListNestedAttribute{
										Description:         "Describes how a group of pods ought to spread across topologydomains. The scheduler will schedule pods in a way which abides by the constraints.Refer to https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/",
										MarkdownDescription: "Describes how a group of pods ought to spread across topologydomains. The scheduler will schedule pods in a way which abides by the constraints.Refer to https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
													MarkdownDescription: "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_skew": schema.Int64Attribute{
													Description:         "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
													MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min_domains": schema.Int64Attribute{
													Description:         "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
													MarkdownDescription: "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_affinity_policy": schema.StringAttribute{
													Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_taints_policy": schema.StringAttribute{
													Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
													MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"when_unsatisfiable": schema.StringAttribute{
													Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
													MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
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

							"volume_claim_restore_policy": schema.StringAttribute{
								Description:         "Defines restore policy for persistent volume claim.Supported policies are as follows:- 'Parallel': parallel recovery of persistent volume claim.- 'Serial': restore the persistent volume claim in sequence, and wait until the previous persistent volume claim is restored before restoring a new one.",
								MarkdownDescription: "Defines restore policy for persistent volume claim.Supported policies are as follows:- 'Parallel': parallel recovery of persistent volume claim.- 'Serial': restore the persistent volume claim in sequence, and wait until the previous persistent volume claim is restored before restoring a new one.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Parallel", "Serial"),
								},
							},

							"volume_claims": schema.ListNestedAttribute{
								Description:         "Defines the persistent Volume claims that need to be restored and mounted together into the restore job.These persistent Volume claims will be created if they do not exist.",
								MarkdownDescription: "Defines the persistent Volume claims that need to be restored and mounted together into the restore job.These persistent Volume claims will be created if they do not exist.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "Specifies the standard metadata for the object.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
											MarkdownDescription: "Specifies the standard metadata for the object.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"finalizers": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"mount_path": schema.StringAttribute{
											Description:         "Specifies the path within the restoring container at which the volume should be mounted.",
											MarkdownDescription: "Specifies the path within the restoring container at which the volume should be mounted.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_claim_spec": schema.SingleNestedAttribute{
											Description:         "Defines the desired characteristics of a persistent volume claim.",
											MarkdownDescription: "Defines the desired characteristics of a persistent volume claim.",
											Attributes: map[string]schema.Attribute{
												"access_modes": schema.ListAttribute{
													Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"data_source": schema.SingleNestedAttribute{
													Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
													MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
													Attributes: map[string]schema.Attribute{
														"api_group": schema.StringAttribute{
															Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
															MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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
													Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													Attributes: map[string]schema.Attribute{
														"api_group": schema.StringAttribute{
															Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
															MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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
															Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
															MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
													Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
													MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
													Attributes: map[string]schema.Attribute{
														"limits": schema.MapAttribute{
															Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"requests": schema.MapAttribute{
															Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

												"storage_class_name": schema.StringAttribute{
													Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
													MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_attributes_class_name": schema.StringAttribute{
													Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
													MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_mode": schema.StringAttribute{
													Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
													MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"volume_source": schema.StringAttribute{
											Description:         "Describes the volume that will be restored from the specified volume of the backup targetVolumes.This is required if the backup uses a volume snapshot.",
											MarkdownDescription: "Describes the volume that will be restored from the specified volume of the backup targetVolumes.This is required if the backup uses a volume snapshot.",
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

							"volume_claims_template": schema.SingleNestedAttribute{
								Description:         "Defines a template to build persistent Volume claims that need to be restored.These claims will be created in an orderly manner based on the number of replicas or reused if they already exist.",
								MarkdownDescription: "Defines a template to build persistent Volume claims that need to be restored.These claims will be created in an orderly manner based on the number of replicas or reused if they already exist.",
								Attributes: map[string]schema.Attribute{
									"replicas": schema.Int64Attribute{
										Description:         "Specifies the replicas of persistent volume claim that need to be created and restored.The format of the created claim name is '$(template-name)-$(index)'.",
										MarkdownDescription: "Specifies the replicas of persistent volume claim that need to be created and restored.The format of the created claim name is '$(template-name)-$(index)'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"starting_index": schema.Int64Attribute{
										Description:         "Specifies the starting index for the created persistent volume claim according to the template.The minimum value is 0.",
										MarkdownDescription: "Specifies the starting index for the created persistent volume claim according to the template.The minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"templates": schema.ListNestedAttribute{
										Description:         "Contains a list of volume claims.",
										MarkdownDescription: "Contains a list of volume claims.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"metadata": schema.SingleNestedAttribute{
													Description:         "Specifies the standard metadata for the object.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
													MarkdownDescription: "Specifies the standard metadata for the object.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"finalizers": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"labels": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"mount_path": schema.StringAttribute{
													Description:         "Specifies the path within the restoring container at which the volume should be mounted.",
													MarkdownDescription: "Specifies the path within the restoring container at which the volume should be mounted.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_claim_spec": schema.SingleNestedAttribute{
													Description:         "Defines the desired characteristics of a persistent volume claim.",
													MarkdownDescription: "Defines the desired characteristics of a persistent volume claim.",
													Attributes: map[string]schema.Attribute{
														"access_modes": schema.ListAttribute{
															Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
															MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"data_source": schema.SingleNestedAttribute{
															Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
															MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																	MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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
															Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
															MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																	MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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
																	Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																	MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
															Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
															MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
															Attributes: map[string]schema.Attribute{
																"limits": schema.MapAttribute{
																	Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"requests": schema.MapAttribute{
																	Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

														"storage_class_name": schema.StringAttribute{
															Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
															MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_attributes_class_name": schema.StringAttribute{
															Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
															MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_mode": schema.StringAttribute{
															Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
															MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"volume_source": schema.StringAttribute{
													Description:         "Describes the volume that will be restored from the specified volume of the backup targetVolumes.This is required if the backup uses a volume snapshot.",
													MarkdownDescription: "Describes the volume that will be restored from the specified volume of the backup targetVolumes.This is required if the backup uses a volume snapshot.",
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

					"ready_config": schema.SingleNestedAttribute{
						Description:         "Configuration for the action of 'postReady' phase.",
						MarkdownDescription: "Configuration for the action of 'postReady' phase.",
						Attributes: map[string]schema.Attribute{
							"connection_credential": schema.SingleNestedAttribute{
								Description:         "Defines the credential template used to create a connection credential.",
								MarkdownDescription: "Defines the credential template used to create a connection credential.",
								Attributes: map[string]schema.Attribute{
									"host_key": schema.StringAttribute{
										Description:         "Specifies the map key of the host in the connection credential secret.",
										MarkdownDescription: "Specifies the map key of the host in the connection credential secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_key": schema.StringAttribute{
										Description:         "Specifies the map key of the password in the connection credential secret.This password will be saved in the backup annotation for full backup.You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
										MarkdownDescription: "Specifies the map key of the password in the connection credential secret.This password will be saved in the backup annotation for full backup.You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_key": schema.StringAttribute{
										Description:         "Specifies the map key of the port in the connection credential secret.",
										MarkdownDescription: "Specifies the map key of the port in the connection credential secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "Refers to the Secret object that contains the connection credential.",
										MarkdownDescription: "Refers to the Secret object that contains the connection credential.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
										},
									},

									"username_key": schema.StringAttribute{
										Description:         "Specifies the map key of the user in the connection credential secret.",
										MarkdownDescription: "Specifies the map key of the user in the connection credential secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"exec_action": schema.SingleNestedAttribute{
								Description:         "Specifies the configuration for an exec action.",
								MarkdownDescription: "Specifies the configuration for an exec action.",
								Attributes: map[string]schema.Attribute{
									"target": schema.SingleNestedAttribute{
										Description:         "Defines the pods that need to be executed for the exec action.Execution will occur on all pods that meet the conditions.",
										MarkdownDescription: "Defines the pods that need to be executed for the exec action.Execution will occur on all pods that meet the conditions.",
										Attributes: map[string]schema.Attribute{
											"pod_selector": schema.SingleNestedAttribute{
												Description:         "Executes kubectl in all selected pods.",
												MarkdownDescription: "Executes kubectl in all selected pods.",
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

							"job_action": schema.SingleNestedAttribute{
								Description:         "Specifies the configuration for a job action.",
								MarkdownDescription: "Specifies the configuration for a job action.",
								Attributes: map[string]schema.Attribute{
									"required_policy_for_all_pod_selection": schema.SingleNestedAttribute{
										Description:         "Specifies the restore policy, which is required when the pod selection strategy for the source target is 'All'.This field is ignored if the pod selection strategy is 'Any'.optional",
										MarkdownDescription: "Specifies the restore policy, which is required when the pod selection strategy for the source target is 'All'.This field is ignored if the pod selection strategy is 'Any'.optional",
										Attributes: map[string]schema.Attribute{
											"data_restore_policy": schema.StringAttribute{
												Description:         "Specifies the data restore policy. Options include:- OneToMany: Enables restoration of all volumes from a single data copy of the original target instance.The 'sourceOfOneToMany' field must be set when using this policy.- OneToOne: Restricts data restoration such that each data piece can only be restored to a single target instance.This is the default policy. When the number of target instances specified for restoration surpasses the count of original backup target instances.",
												MarkdownDescription: "Specifies the data restore policy. Options include:- OneToMany: Enables restoration of all volumes from a single data copy of the original target instance.The 'sourceOfOneToMany' field must be set when using this policy.- OneToOne: Restricts data restoration such that each data piece can only be restored to a single target instance.This is the default policy. When the number of target instances specified for restoration surpasses the count of original backup target instances.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"source_of_one_to_many": schema.SingleNestedAttribute{
												Description:         "Specifies the name of the source target pod. This field is mandatory when the DataRestorePolicy is configured to 'OneToMany'.",
												MarkdownDescription: "Specifies the name of the source target pod. This field is mandatory when the DataRestorePolicy is configured to 'OneToMany'.",
												Attributes: map[string]schema.Attribute{
													"target_pod_name": schema.StringAttribute{
														Description:         "Specifies the name of the source target pod.",
														MarkdownDescription: "Specifies the name of the source target pod.",
														Required:            true,
														Optional:            false,
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

									"target": schema.SingleNestedAttribute{
										Description:         "Defines the pods that needs to be executed for the job action.",
										MarkdownDescription: "Defines the pods that needs to be executed for the job action.",
										Attributes: map[string]schema.Attribute{
											"pod_selector": schema.SingleNestedAttribute{
												Description:         "Selects one of the pods, identified by labels, to build the job spec.This includes mounting required volumes and injecting built-in environment variables of the selected pod.",
												MarkdownDescription: "Selects one of the pods, identified by labels, to build the job spec.This includes mounting required volumes and injecting built-in environment variables of the selected pod.",
												Attributes: map[string]schema.Attribute{
													"fallback_label_selector": schema.SingleNestedAttribute{
														Description:         "fallbackLabelSelector is used to filter available pods when the labelSelector fails.This only takes effect when the 'strategy' field below is set to 'Any'.",
														MarkdownDescription: "fallbackLabelSelector is used to filter available pods when the labelSelector fails.This only takes effect when the 'strategy' field below is set to 'Any'.",
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

													"strategy": schema.StringAttribute{
														Description:         "Specifies the strategy to select the target pod when multiple pods are selected.Valid values are:- 'Any': select any one pod that match the labelsSelector.- 'All': select all pods that match the labelsSelector. The backup data for the current podwill be stored in a subdirectory named after the pod.",
														MarkdownDescription: "Specifies the strategy to select the target pod when multiple pods are selected.Valid values are:- 'Any': select any one pod that match the labelsSelector.- 'All': select all pods that match the labelsSelector. The backup data for the current podwill be stored in a subdirectory named after the pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Any", "All"),
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"volume_mounts": schema.ListNestedAttribute{
												Description:         "Defines which volumes of the selected pod need to be mounted on the restoring pod.",
												MarkdownDescription: "Defines which volumes of the selected pod need to be mounted on the restoring pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"mount_path": schema.StringAttribute{
															Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
															MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"mount_propagation": schema.StringAttribute{
															Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
															MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "This must match the Name of a Volume.",
															MarkdownDescription: "This must match the Name of a Volume.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
															MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sub_path": schema.StringAttribute{
															Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
															MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sub_path_expr": schema.StringAttribute{
															Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
															MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": schema.SingleNestedAttribute{
								Description:         "Defines a periodic probe of the service readiness.The controller will perform postReadyHooks of BackupScript.spec.restoreafter the service readiness when readinessProbe is configured.",
								MarkdownDescription: "Defines a periodic probe of the service readiness.The controller will perform postReadyHooks of BackupScript.spec.restoreafter the service readiness when readinessProbe is configured.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Specifies the action to take.",
										MarkdownDescription: "Specifies the action to take.",
										Attributes: map[string]schema.Attribute{
											"command": schema.ListAttribute{
												Description:         "Refers to the container command.",
												MarkdownDescription: "Refers to the container command.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"image": schema.StringAttribute{
												Description:         "Refers to the container image.",
												MarkdownDescription: "Refers to the container image.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "Specifies the number of seconds after the container has started before the probe is initiated.",
										MarkdownDescription: "Specifies the number of seconds after the container has started before the probe is initiated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "Specifies how often (in seconds) to perform the probe.The default value is 5 seconds, and the minimum value is 1.",
										MarkdownDescription: "Specifies how often (in seconds) to perform the probe.The default value is 5 seconds, and the minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "Specifies the number of seconds after which the probe times out.The default value is 30 seconds, and the minimum value is 1.",
										MarkdownDescription: "Specifies the number of seconds after which the probe times out.The default value is 30 seconds, and the minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
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

					"resources": schema.SingleNestedAttribute{
						Description:         "Restores the specified resources of Kubernetes.",
						MarkdownDescription: "Restores the specified resources of Kubernetes.",
						Attributes: map[string]schema.Attribute{
							"included": schema.ListNestedAttribute{
								Description:         "Restores the specified resources.",
								MarkdownDescription: "Restores the specified resources.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group_resource": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"label_selector": schema.SingleNestedAttribute{
											Description:         "Selects the specified resource for recovery by label.",
											MarkdownDescription: "Selects the specified resource for recovery by label.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"restore_time": schema.StringAttribute{
						Description:         "Specifies the point in time for restoring.",
						MarkdownDescription: "Specifies the point in time for restoring.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`), ""),
						},
					},

					"service_account_name": schema.StringAttribute{
						Description:         "Specifies the service account name needed for recovery pod.",
						MarkdownDescription: "Specifies the service account name needed for recovery pod.",
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

func (r *DataprotectionKubeblocksIoRestoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dataprotection_kubeblocks_io_restore_v1alpha1_manifest")

	var model DataprotectionKubeblocksIoRestoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("dataprotection.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Restore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
