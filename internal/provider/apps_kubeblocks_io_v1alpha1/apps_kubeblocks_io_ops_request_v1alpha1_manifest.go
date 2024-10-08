/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoOpsRequestV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoOpsRequestV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoOpsRequestV1Alpha1Manifest{}
}

type AppsKubeblocksIoOpsRequestV1Alpha1Manifest struct{}

type AppsKubeblocksIoOpsRequestV1Alpha1ManifestData struct {
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
			ParentBackupName *string `tfsdk:"parent_backup_name" json:"parentBackupName,omitempty"`
			RetentionPeriod  *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		BackupSpec *struct {
			BackupMethod     *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
			BackupName       *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			BackupPolicyName *string `tfsdk:"backup_policy_name" json:"backupPolicyName,omitempty"`
			DeletionPolicy   *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
			ParentBackupName *string `tfsdk:"parent_backup_name" json:"parentBackupName,omitempty"`
			RetentionPeriod  *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"backup_spec" json:"backupSpec,omitempty"`
		Cancel      *bool   `tfsdk:"cancel" json:"cancel,omitempty"`
		ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterRef  *string `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		Custom      *struct {
			Components *[]struct {
				ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
				Parameters    *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
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
			Replicas      *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
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
					Image     *string            `tfsdk:"image" json:"image,omitempty"`
					Labels    *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name      *string            `tfsdk:"name" json:"name,omitempty"`
					Replicas  *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
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
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Spec *struct {
							AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
							Resources   *struct {
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
							VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
					VolumeMounts *[]struct {
						MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						Name             *string `tfsdk:"name" json:"name,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
					Volumes *[]struct {
						AwsElasticBlockStore *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							VolumeID  *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
						} `tfsdk:"aws_elastic_block_store" json:"awsElasticBlockStore,omitempty"`
						AzureDisk *struct {
							CachingMode *string `tfsdk:"caching_mode" json:"cachingMode,omitempty"`
							DiskName    *string `tfsdk:"disk_name" json:"diskName,omitempty"`
							DiskURI     *string `tfsdk:"disk_uri" json:"diskURI,omitempty"`
							FsType      *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
							ReadOnly    *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"azure_disk" json:"azureDisk,omitempty"`
						AzureFile *struct {
							ReadOnly   *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
							ShareName  *string `tfsdk:"share_name" json:"shareName,omitempty"`
						} `tfsdk:"azure_file" json:"azureFile,omitempty"`
						Cephfs *struct {
							Monitors   *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
							Path       *string   `tfsdk:"path" json:"path,omitempty"`
							ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretFile *string   `tfsdk:"secret_file" json:"secretFile,omitempty"`
							SecretRef  *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							User *string `tfsdk:"user" json:"user,omitempty"`
						} `tfsdk:"cephfs" json:"cephfs,omitempty"`
						Cinder *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
						} `tfsdk:"cinder" json:"cinder,omitempty"`
						ConfigMap *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						Csi *struct {
							Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
							FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							NodePublishSecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
							ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
							VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
						} `tfsdk:"csi" json:"csi,omitempty"`
						DownwardAPI *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
									FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
								Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path             *string `tfsdk:"path" json:"path,omitempty"`
								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
									Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
									Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
						EmptyDir *struct {
							Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
							SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
						} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
						Ephemeral *struct {
							VolumeClaimTemplate *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
									Name        *string            `tfsdk:"name" json:"name,omitempty"`
									Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *struct {
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
							} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
						} `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
						Fc *struct {
							FsType     *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
							Lun        *int64    `tfsdk:"lun" json:"lun,omitempty"`
							ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							TargetWWNs *[]string `tfsdk:"target_ww_ns" json:"targetWWNs,omitempty"`
							Wwids      *[]string `tfsdk:"wwids" json:"wwids,omitempty"`
						} `tfsdk:"fc" json:"fc,omitempty"`
						FlexVolume *struct {
							Driver    *string            `tfsdk:"driver" json:"driver,omitempty"`
							FsType    *string            `tfsdk:"fs_type" json:"fsType,omitempty"`
							Options   *map[string]string `tfsdk:"options" json:"options,omitempty"`
							ReadOnly  *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						} `tfsdk:"flex_volume" json:"flexVolume,omitempty"`
						Flocker *struct {
							DatasetName *string `tfsdk:"dataset_name" json:"datasetName,omitempty"`
							DatasetUUID *string `tfsdk:"dataset_uuid" json:"datasetUUID,omitempty"`
						} `tfsdk:"flocker" json:"flocker,omitempty"`
						GcePersistentDisk *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
							PdName    *string `tfsdk:"pd_name" json:"pdName,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"gce_persistent_disk" json:"gcePersistentDisk,omitempty"`
						GitRepo *struct {
							Directory  *string `tfsdk:"directory" json:"directory,omitempty"`
							Repository *string `tfsdk:"repository" json:"repository,omitempty"`
							Revision   *string `tfsdk:"revision" json:"revision,omitempty"`
						} `tfsdk:"git_repo" json:"gitRepo,omitempty"`
						Glusterfs *struct {
							Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Path      *string `tfsdk:"path" json:"path,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"glusterfs" json:"glusterfs,omitempty"`
						HostPath *struct {
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"host_path" json:"hostPath,omitempty"`
						Iscsi *struct {
							ChapAuthDiscovery *bool     `tfsdk:"chap_auth_discovery" json:"chapAuthDiscovery,omitempty"`
							ChapAuthSession   *bool     `tfsdk:"chap_auth_session" json:"chapAuthSession,omitempty"`
							FsType            *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
							InitiatorName     *string   `tfsdk:"initiator_name" json:"initiatorName,omitempty"`
							Iqn               *string   `tfsdk:"iqn" json:"iqn,omitempty"`
							IscsiInterface    *string   `tfsdk:"iscsi_interface" json:"iscsiInterface,omitempty"`
							Lun               *int64    `tfsdk:"lun" json:"lun,omitempty"`
							Portals           *[]string `tfsdk:"portals" json:"portals,omitempty"`
							ReadOnly          *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef         *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							TargetPortal *string `tfsdk:"target_portal" json:"targetPortal,omitempty"`
						} `tfsdk:"iscsi" json:"iscsi,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Nfs  *struct {
							Path     *string `tfsdk:"path" json:"path,omitempty"`
							ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							Server   *string `tfsdk:"server" json:"server,omitempty"`
						} `tfsdk:"nfs" json:"nfs,omitempty"`
						PersistentVolumeClaim *struct {
							ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
						PhotonPersistentDisk *struct {
							FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							PdID   *string `tfsdk:"pd_id" json:"pdID,omitempty"`
						} `tfsdk:"photon_persistent_disk" json:"photonPersistentDisk,omitempty"`
						PortworxVolume *struct {
							FsType   *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
						} `tfsdk:"portworx_volume" json:"portworxVolume,omitempty"`
						Projected *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Sources     *[]struct {
								ClusterTrustBundle *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									Name       *string `tfsdk:"name" json:"name,omitempty"`
									Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
									Path       *string `tfsdk:"path" json:"path,omitempty"`
									SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
								} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
								ConfigMap *struct {
									Items *[]struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map" json:"configMap,omitempty"`
								DownwardAPI *struct {
									Items *[]struct {
										FieldRef *struct {
											ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
											FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
										} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
										Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path             *string `tfsdk:"path" json:"path,omitempty"`
										ResourceFieldRef *struct {
											ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
											Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
											Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
										} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
								} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
								Secret *struct {
									Items *[]struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret" json:"secret,omitempty"`
								ServiceAccountToken *struct {
									Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
									ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
									Path              *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
							} `tfsdk:"sources" json:"sources,omitempty"`
						} `tfsdk:"projected" json:"projected,omitempty"`
						Quobyte *struct {
							Group    *string `tfsdk:"group" json:"group,omitempty"`
							ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							Registry *string `tfsdk:"registry" json:"registry,omitempty"`
							Tenant   *string `tfsdk:"tenant" json:"tenant,omitempty"`
							User     *string `tfsdk:"user" json:"user,omitempty"`
							Volume   *string `tfsdk:"volume" json:"volume,omitempty"`
						} `tfsdk:"quobyte" json:"quobyte,omitempty"`
						Rbd *struct {
							FsType    *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
							Image     *string   `tfsdk:"image" json:"image,omitempty"`
							Keyring   *string   `tfsdk:"keyring" json:"keyring,omitempty"`
							Monitors  *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
							Pool      *string   `tfsdk:"pool" json:"pool,omitempty"`
							ReadOnly  *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							User *string `tfsdk:"user" json:"user,omitempty"`
						} `tfsdk:"rbd" json:"rbd,omitempty"`
						ScaleIO *struct {
							FsType           *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Gateway          *string `tfsdk:"gateway" json:"gateway,omitempty"`
							ProtectionDomain *string `tfsdk:"protection_domain" json:"protectionDomain,omitempty"`
							ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef        *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							SslEnabled  *bool   `tfsdk:"ssl_enabled" json:"sslEnabled,omitempty"`
							StorageMode *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
							StoragePool *string `tfsdk:"storage_pool" json:"storagePool,omitempty"`
							System      *string `tfsdk:"system" json:"system,omitempty"`
							VolumeName  *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
						} `tfsdk:"scale_io" json:"scaleIO,omitempty"`
						Secret *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						Storageos *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							VolumeName      *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
							VolumeNamespace *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
						} `tfsdk:"storageos" json:"storageos,omitempty"`
						VsphereVolume *struct {
							FsType            *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							StoragePolicyID   *string `tfsdk:"storage_policy_id" json:"storagePolicyID,omitempty"`
							StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
							VolumePath        *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
						} `tfsdk:"vsphere_volume" json:"vsphereVolume,omitempty"`
					} `tfsdk:"volumes" json:"volumes,omitempty"`
				} `tfsdk:"new_instances" json:"newInstances,omitempty"`
				OfflineInstancesToOnline *[]string `tfsdk:"offline_instances_to_online" json:"offlineInstancesToOnline,omitempty"`
				ReplicaChanges           *int64    `tfsdk:"replica_changes" json:"replicaChanges,omitempty"`
			} `tfsdk:"scale_out" json:"scaleOut,omitempty"`
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
			RestoreEnv *map[string]string `tfsdk:"restore_env" json:"restoreEnv,omitempty"`
		} `tfsdk:"rebuild_from" json:"rebuildFrom,omitempty"`
		Reconfigure *struct {
			ComponentName  *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Configurations *[]struct {
				Keys *[]struct {
					FileContent *string `tfsdk:"file_content" json:"fileContent,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Parameters  *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"keys" json:"keys,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"configurations" json:"configurations,omitempty"`
		} `tfsdk:"reconfigure" json:"reconfigure,omitempty"`
		Reconfigures *[]struct {
			ComponentName  *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Configurations *[]struct {
				Keys *[]struct {
					FileContent *string `tfsdk:"file_content" json:"fileContent,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Parameters  *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"keys" json:"keys,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"configurations" json:"configurations,omitempty"`
		} `tfsdk:"reconfigures" json:"reconfigures,omitempty"`
		Restart *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
		} `tfsdk:"restart" json:"restart,omitempty"`
		Restore *struct {
			BackupName                        *string            `tfsdk:"backup_name" json:"backupName,omitempty"`
			DeferPostReadyUntilClusterRunning *bool              `tfsdk:"defer_post_ready_until_cluster_running" json:"deferPostReadyUntilClusterRunning,omitempty"`
			Env                               *map[string]string `tfsdk:"env" json:"env,omitempty"`
			RestorePointInTime                *string            `tfsdk:"restore_point_in_time" json:"restorePointInTime,omitempty"`
			VolumeRestorePolicy               *string            `tfsdk:"volume_restore_policy" json:"volumeRestorePolicy,omitempty"`
		} `tfsdk:"restore" json:"restore,omitempty"`
		RestoreSpec *struct {
			BackupName                        *string            `tfsdk:"backup_name" json:"backupName,omitempty"`
			DeferPostReadyUntilClusterRunning *bool              `tfsdk:"defer_post_ready_until_cluster_running" json:"deferPostReadyUntilClusterRunning,omitempty"`
			Env                               *map[string]string `tfsdk:"env" json:"env,omitempty"`
			RestorePointInTime                *string            `tfsdk:"restore_point_in_time" json:"restorePointInTime,omitempty"`
			VolumeRestorePolicy               *string            `tfsdk:"volume_restore_policy" json:"volumeRestorePolicy,omitempty"`
		} `tfsdk:"restore_spec" json:"restoreSpec,omitempty"`
		Switchover *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			InstanceName  *string `tfsdk:"instance_name" json:"instanceName,omitempty"`
		} `tfsdk:"switchover" json:"switchover,omitempty"`
		TimeoutSeconds                        *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		TtlSecondsAfterSucceed                *int64  `tfsdk:"ttl_seconds_after_succeed" json:"ttlSecondsAfterSucceed,omitempty"`
		TtlSecondsAfterUnsuccessfulCompletion *int64  `tfsdk:"ttl_seconds_after_unsuccessful_completion" json:"ttlSecondsAfterUnsuccessfulCompletion,omitempty"`
		Type                                  *string `tfsdk:"type" json:"type,omitempty"`
		Upgrade                               *struct {
			ClusterVersionRef *string `tfsdk:"cluster_version_ref" json:"clusterVersionRef,omitempty"`
			Components        *[]struct {
				ComponentDefinitionName *string `tfsdk:"component_definition_name" json:"componentDefinitionName,omitempty"`
				ComponentName           *string `tfsdk:"component_name" json:"componentName,omitempty"`
				ServiceVersion          *string `tfsdk:"service_version" json:"serviceVersion,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
		} `tfsdk:"upgrade" json:"upgrade,omitempty"`
		VerticalScaling *[]map[string]string `tfsdk:"vertical_scaling" json:"verticalScaling,omitempty"`
		VolumeExpansion *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Instances     *[]struct {
				Name                 *string `tfsdk:"name" json:"name,omitempty"`
				VolumeClaimTemplates *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Storage *string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
			} `tfsdk:"instances" json:"instances,omitempty"`
			VolumeClaimTemplates *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"volume_expansion" json:"volumeExpansion,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_ops_request_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Description:         "Specifies the parameters to backup a Cluster.",
						MarkdownDescription: "Specifies the parameters to backup a Cluster.",
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

					"backup_spec": schema.SingleNestedAttribute{
						Description:         "Deprecated: since v0.9, use backup instead. Specifies the parameters to backup a Cluster.",
						MarkdownDescription: "Deprecated: since v0.9, use backup instead. Specifies the parameters to backup a Cluster.",
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

					"cluster_ref": schema.StringAttribute{
						Description:         "Deprecated: since v0.9, use clusterName instead. Specifies the name of the Cluster resource that this operation is targeting.",
						MarkdownDescription: "Deprecated: since v0.9, use clusterName instead. Specifies the name of the Cluster resource that this operation is targeting.",
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
											Description:         "Specifies the name of the Component.",
											MarkdownDescription: "Specifies the name of the Component.",
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
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Deprecated: since v0.9, use scaleOut and scaleIn instead. Specifies the number of replicas for the component. Cannot be used with 'scaleIn' and 'scaleOut'.",
									MarkdownDescription: "Deprecated: since v0.9, use scaleOut and scaleIn instead. Specifies the number of replicas for the component. Cannot be used with 'scaleIn' and 'scaleOut'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
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

													"image": schema.StringAttribute{
														Description:         "Specifies an override for the first container's image in the Pod.",
														MarkdownDescription: "Specifies an override for the first container's image in the Pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
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
														Description:         "Specifies the scheduling policy for the Component.",
														MarkdownDescription: "Specifies the scheduling policy for the Component.",
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
														Description:         "Defines VolumeClaimTemplates to override. Add new or override existing volume claim templates.",
														MarkdownDescription: "Defines VolumeClaimTemplates to override. Add new or override existing volume claim templates.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Refers to the name of a volumeMount defined in either: - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated) The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
																	MarkdownDescription: "Refers to the name of a volumeMount defined in either: - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated) The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"spec": schema.SingleNestedAttribute{
																	Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field. When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
																	MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field. When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
																	Attributes: map[string]schema.Attribute{
																		"access_modes": schema.MapAttribute{
																			Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																			MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"resources": schema.SingleNestedAttribute{
																			Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																			MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
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

																		"storage_class_name": schema.StringAttribute{
																			Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																			MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_mode": schema.StringAttribute{
																			Description:         "Defines what type of volume is required by the claim, either Block or Filesystem.",
																			MarkdownDescription: "Defines what type of volume is required by the claim, either Block or Filesystem.",
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

													"volume_mounts": schema.ListNestedAttribute{
														Description:         "Defines VolumeMounts to override. Add new or override existing volume mounts of the first container in the Pod.",
														MarkdownDescription: "Defines VolumeMounts to override. Add new or override existing volume mounts of the first container in the Pod.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"mount_path": schema.StringAttribute{
																	Description:         "Path within the container at which the volume should be mounted. Must not contain ':'.",
																	MarkdownDescription: "Path within the container at which the volume should be mounted. Must not contain ':'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"mount_propagation": schema.StringAttribute{
																	Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																	MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
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
																	Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
																	MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_path": schema.StringAttribute{
																	Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
																	MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_path_expr": schema.StringAttribute{
																	Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
																	MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

													"volumes": schema.ListNestedAttribute{
														Description:         "Defines Volumes to override. Add new or override existing volumes.",
														MarkdownDescription: "Defines Volumes to override. Add new or override existing volumes.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"aws_elastic_block_store": schema.SingleNestedAttribute{
																	Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																			MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_id": schema.StringAttribute{
																			Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"azure_disk": schema.SingleNestedAttribute{
																	Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
																	MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
																	Attributes: map[string]schema.Attribute{
																		"caching_mode": schema.StringAttribute{
																			Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																			MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"disk_name": schema.StringAttribute{
																			Description:         "diskName is the Name of the data disk in the blob storage",
																			MarkdownDescription: "diskName is the Name of the data disk in the blob storage",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"disk_uri": schema.StringAttribute{
																			Description:         "diskURI is the URI of data disk in the blob storage",
																			MarkdownDescription: "diskURI is the URI of data disk in the blob storage",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "kind expected values are Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared",
																			MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"azure_file": schema.SingleNestedAttribute{
																	Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
																	MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
																	Attributes: map[string]schema.Attribute{
																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
																			Description:         "secretName is the name of secret that contains Azure Storage Account Name and Key",
																			MarkdownDescription: "secretName is the name of secret that contains Azure Storage Account Name and Key",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"share_name": schema.StringAttribute{
																			Description:         "shareName is the azure share Name",
																			MarkdownDescription: "shareName is the azure share Name",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"cephfs": schema.SingleNestedAttribute{
																	Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																	MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																	Attributes: map[string]schema.Attribute{
																		"monitors": schema.ListAttribute{
																			Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																			MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_file": schema.StringAttribute{
																			Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"cinder": schema.SingleNestedAttribute{
																	Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																			MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"volume_id": schema.StringAttribute{
																			Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"config_map": schema.SingleNestedAttribute{
																	Description:         "configMap represents a configMap that should populate this volume",
																	MarkdownDescription: "configMap represents a configMap that should populate this volume",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																			MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the key to project.",
																						MarkdownDescription: "key is the key to project.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"mode": schema.Int64Attribute{
																						Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																						MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
																			Description:         "optional specify whether the ConfigMap or its keys must be defined",
																			MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"csi": schema.SingleNestedAttribute{
																	Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
																	MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																			MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																			MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"node_publish_secret_ref": schema.SingleNestedAttribute{
																			Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																			MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																			MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_attributes": schema.MapAttribute{
																			Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																			MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
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

																"downward_api": schema.SingleNestedAttribute{
																	Description:         "downwardAPI represents downward API about the pod that should populate this volume",
																	MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "Items is a list of downward API volume file",
																			MarkdownDescription: "Items is a list of downward API volume file",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"field_ref": schema.SingleNestedAttribute{
																						Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																						MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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

																					"mode": schema.Int64Attribute{
																						Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																						MarkdownDescription: "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																						MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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

																"empty_dir": schema.SingleNestedAttribute{
																	Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	Attributes: map[string]schema.Attribute{
																		"medium": schema.StringAttribute{
																			Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"size_limit": schema.StringAttribute{
																			Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"ephemeral": schema.SingleNestedAttribute{
																	Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed. Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim). Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod. Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information. A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
																	MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed. Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim). Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod. Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information. A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
																	Attributes: map[string]schema.Attribute{
																		"volume_claim_template": schema.SingleNestedAttribute{
																			Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod. The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long). An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster. This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created. Required, must not be nil.",
																			MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod. The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long). An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster. This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created. Required, must not be nil.",
																			Attributes: map[string]schema.Attribute{
																				"metadata": schema.SingleNestedAttribute{
																					Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																					MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
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
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"spec": schema.SingleNestedAttribute{
																					Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																					MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
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

																"fc": schema.SingleNestedAttribute{
																	Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																	MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "lun is Optional: FC target lun number",
																			MarkdownDescription: "lun is Optional: FC target lun number",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"target_ww_ns": schema.ListAttribute{
																			Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
																			MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"wwids": schema.ListAttribute{
																			Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																			MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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

																"flex_volume": schema.SingleNestedAttribute{
																	Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																	MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "driver is the name of the driver to use for this volume.",
																			MarkdownDescription: "driver is the name of the driver to use for this volume.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"options": schema.MapAttribute{
																			Description:         "options is Optional: this field holds extra command options if any.",
																			MarkdownDescription: "options is Optional: this field holds extra command options if any.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																			MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

																"flocker": schema.SingleNestedAttribute{
																	Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																	MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																	Attributes: map[string]schema.Attribute{
																		"dataset_name": schema.StringAttribute{
																			Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																			MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"dataset_uuid": schema.StringAttribute{
																			Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																			MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"gce_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pd_name": schema.StringAttribute{
																			Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"git_repo": schema.SingleNestedAttribute{
																	Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																	MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																	Attributes: map[string]schema.Attribute{
																		"directory": schema.StringAttribute{
																			Description:         "directory is the target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																			MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"repository": schema.StringAttribute{
																			Description:         "repository is the URL",
																			MarkdownDescription: "repository is the URL",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"revision": schema.StringAttribute{
																			Description:         "revision is the commit hash for the specified revision.",
																			MarkdownDescription: "revision is the commit hash for the specified revision.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"glusterfs": schema.SingleNestedAttribute{
																	Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																	MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																	Attributes: map[string]schema.Attribute{
																		"endpoints": schema.StringAttribute{
																			Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"host_path": schema.SingleNestedAttribute{
																	Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																	MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"iscsi": schema.SingleNestedAttribute{
																	Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																	MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																	Attributes: map[string]schema.Attribute{
																		"chap_auth_discovery": schema.BoolAttribute{
																			Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																			MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"chap_auth_session": schema.BoolAttribute{
																			Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																			MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"initiator_name": schema.StringAttribute{
																			Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																			MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"iqn": schema.StringAttribute{
																			Description:         "iqn is the target iSCSI Qualified Name.",
																			MarkdownDescription: "iqn is the target iSCSI Qualified Name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"iscsi_interface": schema.StringAttribute{
																			Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																			MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "lun represents iSCSI Target Lun number.",
																			MarkdownDescription: "lun represents iSCSI Target Lun number.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"portals": schema.ListAttribute{
																			Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																			MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																			MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"target_portal": schema.StringAttribute{
																			Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"name": schema.StringAttribute{
																	Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"nfs": schema.SingleNestedAttribute{
																	Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"server": schema.StringAttribute{
																			Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"persistent_volume_claim": schema.SingleNestedAttribute{
																	Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	Attributes: map[string]schema.Attribute{
																		"claim_name": schema.StringAttribute{
																			Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																			MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																			MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"photon_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																	MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pd_id": schema.StringAttribute{
																			Description:         "pdID is the ID that identifies Photon Controller persistent disk",
																			MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"portworx_volume": schema.SingleNestedAttribute{
																	Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																	MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_id": schema.StringAttribute{
																			Description:         "volumeID uniquely identifies a Portworx volume",
																			MarkdownDescription: "volumeID uniquely identifies a Portworx volume",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"projected": schema.SingleNestedAttribute{
																	Description:         "projected items for all in one resources secrets, configmaps, and downward API",
																	MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sources": schema.ListNestedAttribute{
																			Description:         "sources is the list of volume projections",
																			MarkdownDescription: "sources is the list of volume projections",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"cluster_trust_bundle": schema.SingleNestedAttribute{
																						Description:         "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' field of ClusterTrustBundle objects in an auto-updating file. Alpha, gated by the ClusterTrustBundleProjection feature gate. ClusterTrustBundle objects can either be selected by name, or by the combination of signer name and a label selector. Kubelet performs aggressive normalization of the PEM contents written into the pod filesystem. Esoteric PEM features such as inter-block comments and block headers are stripped. Certificates are deduplicated. The ordering of certificates within the file is arbitrary, and Kubelet may change the order over time.",
																						MarkdownDescription: "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' field of ClusterTrustBundle objects in an auto-updating file. Alpha, gated by the ClusterTrustBundleProjection feature gate. ClusterTrustBundle objects can either be selected by name, or by the combination of signer name and a label selector. Kubelet performs aggressive normalization of the PEM contents written into the pod filesystem. Esoteric PEM features such as inter-block comments and block headers are stripped. Certificates are deduplicated. The ordering of certificates within the file is arbitrary, and Kubelet may change the order over time.",
																						Attributes: map[string]schema.Attribute{
																							"label_selector": schema.SingleNestedAttribute{
																								Description:         "Select all ClusterTrustBundles that match this label selector. Only has effect if signerName is set. Mutually-exclusive with name. If unset, interpreted as 'match nothing'. If set but empty, interpreted as 'match everything'.",
																								MarkdownDescription: "Select all ClusterTrustBundles that match this label selector. Only has effect if signerName is set. Mutually-exclusive with name. If unset, interpreted as 'match nothing'. If set but empty, interpreted as 'match everything'.",
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
																								Description:         "Select a single ClusterTrustBundle by object name. Mutually-exclusive with signerName and labelSelector.",
																								MarkdownDescription: "Select a single ClusterTrustBundle by object name. Mutually-exclusive with signerName and labelSelector.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
																								Description:         "If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available. If using name, then the named ClusterTrustBundle is allowed not to exist. If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.",
																								MarkdownDescription: "If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available. If using name, then the named ClusterTrustBundle is allowed not to exist. If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "Relative path from the volume root to write the bundle.",
																								MarkdownDescription: "Relative path from the volume root to write the bundle.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"signer_name": schema.StringAttribute{
																								Description:         "Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name. The contents of all selected ClusterTrustBundles will be unified and deduplicated.",
																								MarkdownDescription: "Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name. The contents of all selected ClusterTrustBundles will be unified and deduplicated.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"config_map": schema.SingleNestedAttribute{
																						Description:         "configMap information about the configMap data to project",
																						MarkdownDescription: "configMap information about the configMap data to project",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "key is the key to project.",
																											MarkdownDescription: "key is the key to project.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"mode": schema.Int64Attribute{
																											Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																											MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																							"name": schema.StringAttribute{
																								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
																								Description:         "optional specify whether the ConfigMap or its keys must be defined",
																								MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"downward_api": schema.SingleNestedAttribute{
																						Description:         "downwardAPI information about the downwardAPI data to project",
																						MarkdownDescription: "downwardAPI information about the downwardAPI data to project",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "Items is a list of DownwardAPIVolume file",
																								MarkdownDescription: "Items is a list of DownwardAPIVolume file",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"field_ref": schema.SingleNestedAttribute{
																											Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																											MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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

																										"mode": schema.Int64Attribute{
																											Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																											MarkdownDescription: "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"resource_field_ref": schema.SingleNestedAttribute{
																											Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																											MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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

																					"secret": schema.SingleNestedAttribute{
																						Description:         "secret information about the secret data to project",
																						MarkdownDescription: "secret information about the secret data to project",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "key is the key to project.",
																											MarkdownDescription: "key is the key to project.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"mode": schema.Int64Attribute{
																											Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																											MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																							"name": schema.StringAttribute{
																								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
																								Description:         "optional field specify whether the Secret or its key must be defined",
																								MarkdownDescription: "optional field specify whether the Secret or its key must be defined",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"service_account_token": schema.SingleNestedAttribute{
																						Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
																						MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",
																						Attributes: map[string]schema.Attribute{
																							"audience": schema.StringAttribute{
																								Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																								MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"expiration_seconds": schema.Int64Attribute{
																								Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																								MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "path is the path relative to the mount point of the file to project the token into.",
																								MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",
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

																"quobyte": schema.SingleNestedAttribute{
																	Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																	MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "group to map volume access to Default is no group",
																			MarkdownDescription: "group to map volume access to Default is no group",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																			MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"registry": schema.StringAttribute{
																			Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																			MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"tenant": schema.StringAttribute{
																			Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																			MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "user to map volume access to Defaults to serivceaccount user",
																			MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume": schema.StringAttribute{
																			Description:         "volume is a string that references an already created Quobyte volume by name.",
																			MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"rbd": schema.SingleNestedAttribute{
																	Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																	MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"image": schema.StringAttribute{
																			Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"keyring": schema.StringAttribute{
																			Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"monitors": schema.ListAttribute{
																			Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"pool": schema.StringAttribute{
																			Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"scale_io": schema.SingleNestedAttribute{
																	Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																	MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"gateway": schema.StringAttribute{
																			Description:         "gateway is the host address of the ScaleIO API Gateway.",
																			MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"protection_domain": schema.StringAttribute{
																			Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																			MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																			MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"ssl_enabled": schema.BoolAttribute{
																			Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																			MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_mode": schema.StringAttribute{
																			Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																			MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_pool": schema.StringAttribute{
																			Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																			MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"system": schema.StringAttribute{
																			Description:         "system is the name of the storage system as configured in ScaleIO.",
																			MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"volume_name": schema.StringAttribute{
																			Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
																			MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"secret": schema.SingleNestedAttribute{
																	Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																			MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the key to project.",
																						MarkdownDescription: "key is the key to project.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"mode": schema.Int64Attribute{
																						Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																						MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																		"optional": schema.BoolAttribute{
																			Description:         "optional field specify whether the Secret or its keys must be defined",
																			MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
																			Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																			MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"storageos": schema.SingleNestedAttribute{
																	Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																	MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials. If not specified, default values will be attempted.",
																			MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials. If not specified, default values will be attempted.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"volume_name": schema.StringAttribute{
																			Description:         "volumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.",
																			MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_namespace": schema.StringAttribute{
																			Description:         "volumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																			MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"vsphere_volume": schema.SingleNestedAttribute{
																	Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																	MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_policy_id": schema.StringAttribute{
																			Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																			MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_policy_name": schema.StringAttribute{
																			Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																			MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_path": schema.StringAttribute{
																			Description:         "volumePath is the path that identifies vSphere volume vmdk",
																			MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",
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
						Description:         "Specifies the parameters to rebuild some instances. Rebuilding an instance involves restoring its data from a backup or another database replica. The instances being rebuilt usually serve as standby in the cluster. Hence rebuilding instances is often also referred to as 'standby reconstruction'.",
						MarkdownDescription: "Specifies the parameters to rebuild some instances. Rebuilding an instance involves restoring its data from a backup or another database replica. The instances being rebuilt usually serve as standby in the cluster. Hence rebuilding instances is often also referred to as 'standby reconstruction'.",
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
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"in_place": schema.BoolAttribute{
									Description:         "When it is set to true, the instance will be rebuilt in-place. By default, a new pod will be created. Once the new pod is ready to serve, the instance that require rebuilding will be taken offline.",
									MarkdownDescription: "When it is set to true, the instance will be rebuilt in-place. By default, a new pod will be created. Once the new pod is ready to serve, the instance that require rebuilding will be taken offline.",
									Required:            false,
									Optional:            true,
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"reconfigure": schema.SingleNestedAttribute{
						Description:         "Specifies a component and its configuration updates. This field is deprecated and replaced by 'reconfigures'.",
						MarkdownDescription: "Specifies a component and its configuration updates. This field is deprecated and replaced by 'reconfigures'.",
						Attributes: map[string]schema.Attribute{
							"component_name": schema.StringAttribute{
								Description:         "Specifies the name of the Component.",
								MarkdownDescription: "Specifies the name of the Component.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"configurations": schema.ListNestedAttribute{
								Description:         "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
								MarkdownDescription: "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"keys": schema.ListNestedAttribute{
											Description:         "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
											MarkdownDescription: "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"file_content": schema.StringAttribute{
														Description:         "Specifies the content of the entire configuration file. This field is used to update the complete configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														MarkdownDescription: "Specifies the content of the entire configuration file. This field is used to update the complete configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
														MarkdownDescription: "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"parameters": schema.ListNestedAttribute{
														Description:         "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														MarkdownDescription: "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the configuration template.",
											MarkdownDescription: "Specifies the name of the configuration template.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"policy": schema.StringAttribute{
											Description:         "Defines the upgrade policy for the configuration.",
											MarkdownDescription: "Defines the upgrade policy for the configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("simple", "parallel", "rolling", "autoReload", "operatorSyncUpdate", "dynamicReloadBeginRestart"),
											},
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

					"reconfigures": schema.ListNestedAttribute{
						Description:         "Lists Reconfigure objects, each specifying a Component and its configuration updates.",
						MarkdownDescription: "Lists Reconfigure objects, each specifying a Component and its configuration updates.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"configurations": schema.ListNestedAttribute{
									Description:         "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
									MarkdownDescription: "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"keys": schema.ListNestedAttribute{
												Description:         "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
												MarkdownDescription: "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"file_content": schema.StringAttribute{
															Description:         "Specifies the content of the entire configuration file. This field is used to update the complete configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															MarkdownDescription: "Specifies the content of the entire configuration file. This field is used to update the complete configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
															MarkdownDescription: "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"parameters": schema.ListNestedAttribute{
															Description:         "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															MarkdownDescription: "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file. Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
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
												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "Specifies the name of the configuration template.",
												MarkdownDescription: "Specifies the name of the configuration template.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
												},
											},

											"policy": schema.StringAttribute{
												Description:         "Defines the upgrade policy for the configuration.",
												MarkdownDescription: "Defines the upgrade policy for the configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("simple", "parallel", "rolling", "autoReload", "operatorSyncUpdate", "dynamicReloadBeginRestart"),
												},
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

					"restart": schema.ListNestedAttribute{
						Description:         "Lists Components to be restarted.",
						MarkdownDescription: "Lists Components to be restarted.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
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

					"restore_spec": schema.SingleNestedAttribute{
						Description:         "Deprecated: since v0.9, use restore instead. Specifies the parameters to restore a Cluster. Note that this restore operation will roll back cluster services.",
						MarkdownDescription: "Deprecated: since v0.9, use restore instead. Specifies the parameters to restore a Cluster. Note that this restore operation will roll back cluster services.",
						Attributes: map[string]schema.Attribute{
							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the Backup custom resource.",
								MarkdownDescription: "Specifies the name of the Backup custom resource.",
								Required:            true,
								Optional:            false,
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

					"switchover": schema.ListNestedAttribute{
						Description:         "Lists Switchover objects, each specifying a Component to perform the switchover operation.",
						MarkdownDescription: "Lists Switchover objects, each specifying a Component to perform the switchover operation.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instance_name": schema.StringAttribute{
									Description:         "Specifies the instance to become the primary or leader during a switchover operation. The value of 'instanceName' can be either: 1. '*' (wildcard value): - Indicates no specific instance is designated as the primary or leader. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withoutCandidate'. - 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' must be defined when using '*'. 2. A valid instance name (pod name): - Designates a specific instance (pod) as the primary or leader. - The name must match one of the pods in the component. Any non-valid pod name is considered invalid. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate'. - 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate' must be defined when specifying a valid instance name.",
									MarkdownDescription: "Specifies the instance to become the primary or leader during a switchover operation. The value of 'instanceName' can be either: 1. '*' (wildcard value): - Indicates no specific instance is designated as the primary or leader. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withoutCandidate'. - 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' must be defined when using '*'. 2. A valid instance name (pod name): - Designates a specific instance (pod) as the primary or leader. - The name must match one of the pods in the component. Any non-valid pod name is considered invalid. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate'. - 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate' must be defined when specifying a valid instance name.",
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
							"cluster_version_ref": schema.StringAttribute{
								Description:         "Deprecated: since v0.9 because ClusterVersion is deprecated. Specifies the name of the target ClusterVersion for the upgrade.",
								MarkdownDescription: "Deprecated: since v0.9 because ClusterVersion is deprecated. Specifies the name of the target ClusterVersion for the upgrade.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"components": schema.ListNestedAttribute{
								Description:         "Lists components to be upgrade based on desired ComponentDefinition and ServiceVersion. From the perspective of cluster API, the reasonable combinations should be: 1. (comp-def, service-ver) - upgrade to the specified service version and component definition, the user takes the responsibility to ensure that they are compatible. 2. ('', service-ver) - upgrade to the specified service version, let the operator choose the latest compatible component definition. 3. (comp-def, '') - upgrade to the specified component definition, let the operator choose the latest compatible service version. 4. ('', '') - upgrade to the latest service version and component definition, the operator will ensure the compatibility between the selected versions.",
								MarkdownDescription: "Lists components to be upgrade based on desired ComponentDefinition and ServiceVersion. From the perspective of cluster API, the reasonable combinations should be: 1. (comp-def, service-ver) - upgrade to the specified service version and component definition, the user takes the responsibility to ensure that they are compatible. 2. ('', service-ver) - upgrade to the specified service version, let the operator choose the latest compatible component definition. 3. (comp-def, '') - upgrade to the specified component definition, let the operator choose the latest compatible service version. 4. ('', '') - upgrade to the latest service version and component definition, the operator will ensure the compatibility between the selected versions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"component_definition_name": schema.StringAttribute{
											Description:         "Specifies the name of the ComponentDefinition.",
											MarkdownDescription: "Specifies the name of the ComponentDefinition.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(64),
											},
										},

										"component_name": schema.StringAttribute{
											Description:         "Specifies the name of the Component.",
											MarkdownDescription: "Specifies the name of the Component.",
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
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instances": schema.ListNestedAttribute{
									Description:         "Specifies the desired storage size of the instance template that need to volume expand.",
									MarkdownDescription: "Specifies the desired storage size of the instance template that need to volume expand.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refer to the instance template name of the component or sharding.",
												MarkdownDescription: "Refer to the instance template name of the component or sharding.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"volume_claim_templates": schema.ListNestedAttribute{
												Description:         "volumeClaimTemplates specifies the storage size and volumeClaimTemplate name.",
												MarkdownDescription: "volumeClaimTemplates specifies the storage size and volumeClaimTemplate name.",
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

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest")

	var model AppsKubeblocksIoOpsRequestV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("OpsRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
