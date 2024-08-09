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
	_ datasource.DataSource = &AppsKubeblocksIoComponentV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoComponentV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoComponentV1Alpha1Manifest{}
}

type AppsKubeblocksIoComponentV1Alpha1Manifest struct{}

type AppsKubeblocksIoComponentV1Alpha1ManifestData struct {
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
		Affinity *struct {
			NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
			PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
			Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
			TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
		CompDef     *string            `tfsdk:"comp_def" json:"compDef,omitempty"`
		Configs     *[]struct {
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
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"configs" json:"configs,omitempty"`
		DisableExporter *bool     `tfsdk:"disable_exporter" json:"disableExporter,omitempty"`
		EnabledLogs     *[]string `tfsdk:"enabled_logs" json:"enabledLogs,omitempty"`
		Env             *[]struct {
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
		Instances *[]struct {
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
		} `tfsdk:"instances" json:"instances,omitempty"`
		Labels                           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		OfflineInstances                 *[]string          `tfsdk:"offline_instances" json:"offlineInstances,omitempty"`
		ParallelPodManagementConcurrency *string            `tfsdk:"parallel_pod_management_concurrency" json:"parallelPodManagementConcurrency,omitempty"`
		PodUpdatePolicy                  *string            `tfsdk:"pod_update_policy" json:"podUpdatePolicy,omitempty"`
		Replicas                         *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources                        *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RuntimeClassName *string `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
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
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		ServiceRefs        *[]struct {
			Cluster                *string `tfsdk:"cluster" json:"cluster,omitempty"`
			ClusterServiceSelector *struct {
				Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Credential *struct {
					Component *string `tfsdk:"component" json:"component,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credential" json:"credential,omitempty"`
				Service *struct {
					Component *string `tfsdk:"component" json:"component,omitempty"`
					Port      *string `tfsdk:"port" json:"port,omitempty"`
					Service   *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
			} `tfsdk:"cluster_service_selector" json:"clusterServiceSelector,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			Namespace         *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ServiceDescriptor *string `tfsdk:"service_descriptor" json:"serviceDescriptor,omitempty"`
		} `tfsdk:"service_refs" json:"serviceRefs,omitempty"`
		ServiceVersion *string `tfsdk:"service_version" json:"serviceVersion,omitempty"`
		Services       *[]struct {
			Annotations          *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			DisableAutoProvision *bool              `tfsdk:"disable_auto_provision" json:"disableAutoProvision,omitempty"`
			Name                 *string            `tfsdk:"name" json:"name,omitempty"`
			PodService           *bool              `tfsdk:"pod_service" json:"podService,omitempty"`
			RoleSelector         *string            `tfsdk:"role_selector" json:"roleSelector,omitempty"`
			ServiceName          *string            `tfsdk:"service_name" json:"serviceName,omitempty"`
			Spec                 *struct {
				AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
				ClusterIP                     *string   `tfsdk:"cluster_ip" json:"clusterIP,omitempty"`
				ClusterIPs                    *[]string `tfsdk:"cluster_i_ps" json:"clusterIPs,omitempty"`
				ExternalIPs                   *[]string `tfsdk:"external_i_ps" json:"externalIPs,omitempty"`
				ExternalName                  *string   `tfsdk:"external_name" json:"externalName,omitempty"`
				ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
				HealthCheckNodePort           *int64    `tfsdk:"health_check_node_port" json:"healthCheckNodePort,omitempty"`
				InternalTrafficPolicy         *string   `tfsdk:"internal_traffic_policy" json:"internalTrafficPolicy,omitempty"`
				IpFamilies                    *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
				IpFamilyPolicy                *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
				LoadBalancerClass             *string   `tfsdk:"load_balancer_class" json:"loadBalancerClass,omitempty"`
				LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
				LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
				Ports                         *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
					TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				PublishNotReadyAddresses *bool              `tfsdk:"publish_not_ready_addresses" json:"publishNotReadyAddresses,omitempty"`
				Selector                 *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				SessionAffinity          *string            `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
				SessionAffinityConfig    *struct {
					ClientIP *struct {
						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					} `tfsdk:"client_ip" json:"clientIP,omitempty"`
				} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		Stop           *bool `tfsdk:"stop" json:"stop,omitempty"`
		SystemAccounts *[]struct {
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			PasswordConfig *struct {
				Length     *int64  `tfsdk:"length" json:"length,omitempty"`
				LetterCase *string `tfsdk:"letter_case" json:"letterCase,omitempty"`
				NumDigits  *int64  `tfsdk:"num_digits" json:"numDigits,omitempty"`
				NumSymbols *int64  `tfsdk:"num_symbols" json:"numSymbols,omitempty"`
				Seed       *string `tfsdk:"seed" json:"seed,omitempty"`
			} `tfsdk:"password_config" json:"passwordConfig,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"system_accounts" json:"systemAccounts,omitempty"`
		TlsConfig *struct {
			Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
			Issuer *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				SecretRef *struct {
					Ca   *string `tfsdk:"ca" json:"ca,omitempty"`
					Cert *string `tfsdk:"cert" json:"cert,omitempty"`
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"issuer" json:"issuer,omitempty"`
		} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
		Tolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoComponentV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_component_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoComponentV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Component is a fundamental building block of a Cluster object.For example, a Redis Cluster can include Components like 'redis', 'sentinel', and potentially a proxy like 'twemproxy'.The Component object is responsible for managing the lifecycle of all replicas within a Cluster component,It supports a wide range of operations including provisioning, stopping, restarting, termination, upgrading,configuration changes, vertical and horizontal scaling, failover, switchover, cross-node migration,scheduling configuration, exposing Services, managing system accounts, enabling/disabling exporter,and configuring log collection.Component is an internal sub-object derived from the user-submitted Cluster object.It is designed primarily to be used by the KubeBlocks controllers,users are discouraged from modifying Component objects directly and should use them only for monitoring Component statuses.",
		MarkdownDescription: "Component is a fundamental building block of a Cluster object.For example, a Redis Cluster can include Components like 'redis', 'sentinel', and potentially a proxy like 'twemproxy'.The Component object is responsible for managing the lifecycle of all replicas within a Cluster component,It supports a wide range of operations including provisioning, stopping, restarting, termination, upgrading,configuration changes, vertical and horizontal scaling, failover, switchover, cross-node migration,scheduling configuration, exposing Services, managing system accounts, enabling/disabling exporter,and configuring log collection.Component is an internal sub-object derived from the user-submitted Cluster object.It is designed primarily to be used by the KubeBlocks controllers,users are discouraged from modifying Component objects directly and should use them only for monitoring Component statuses.",
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
				Description:         "ComponentSpec defines the desired state of Component.",
				MarkdownDescription: "ComponentSpec defines the desired state of Component.",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Specifies a group of affinity scheduling rules for the Component.It allows users to control how the Component's Pods are scheduled onto nodes in the Cluster.Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
						MarkdownDescription: "Specifies a group of affinity scheduling rules for the Component.It allows users to control how the Component's Pods are scheduled onto nodes in the Cluster.Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
						Attributes: map[string]schema.Attribute{
							"node_labels": schema.MapAttribute{
								Description:         "Indicates the node labels that must be present on nodes for pods to be scheduled on them.It is a map where the keys are the label keys and the values are the corresponding label values.Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'},pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd'and the 'environment' label with value 'production'.This field allows users to control Pod placement based on specific node labels.It can be used to ensure that Pods are scheduled on nodes with certain characteristics,such as specific hardware (e.g., SSD), environment (e.g., production, staging),or any other custom labels assigned to nodes.",
								MarkdownDescription: "Indicates the node labels that must be present on nodes for pods to be scheduled on them.It is a map where the keys are the label keys and the values are the corresponding label values.Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'},pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd'and the 'environment' label with value 'production'.This field allows users to control Pod placement based on specific node labels.It can be used to ensure that Pods are scheduled on nodes with certain characteristics,such as specific hardware (e.g., SSD), environment (e.g., production, staging),or any other custom labels assigned to nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_anti_affinity": schema.StringAttribute{
								Description:         "Specifies the anti-affinity level of Pods within a Component.It determines how pods should be spread across nodes to improve availability and performance.It can have the following values: 'Preferred' and 'Required'.The default value is 'Preferred'.",
								MarkdownDescription: "Specifies the anti-affinity level of Pods within a Component.It determines how pods should be spread across nodes to improve availability and performance.It can have the following values: 'Preferred' and 'Required'.The default value is 'Preferred'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Preferred", "Required"),
								},
							},

							"tenancy": schema.StringAttribute{
								Description:         "Determines the level of resource isolation between Pods.It can have the following values: 'SharedNode' and 'DedicatedNode'.- SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s.- DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node.  In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node.  Which provides a higher level of isolation and resource guarantee for Pods. The default value is 'SharedNode'.",
								MarkdownDescription: "Determines the level of resource isolation between Pods.It can have the following values: 'SharedNode' and 'DedicatedNode'.- SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s.- DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node.  In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node.  Which provides a higher level of isolation and resource guarantee for Pods. The default value is 'SharedNode'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("SharedNode", "DedicatedNode"),
								},
							},

							"topology_keys": schema.ListAttribute{
								Description:         "Represents the key of node labels used to define the topology domain for Pod anti-affinityand Pod spread constraints.In K8s, a topology domain is a set of nodes that have the same value for a specific label key.Nodes with labels containing any of the specified TopologyKeys and identical values are consideredto be in the same topology domain.Note: The concept of topology in the context of K8s TopologyKeys is different from the concept oftopology in the ClusterDefinition.When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule thePod on nodes with different values for the specified TopologyKeys.This ensures that Pods are spread across different topology domains, promoting high availability andreducing the impact of node failures.Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone',are often used as TopologyKey.These keys represent the hostname and zone of a node, respectively.By including these keys in the TopologyKeys list, Pods will be spread across nodes withdifferent hostnames or zones.In addition to the well-known keys, users can also specify custom label keys as TopologyKeys.This allows for more flexible and custom topology definitions based on the specific needsof the application or environment.The TopologyKeys field is a slice of strings, where each string represents a label key.The order of the keys in the slice does not matter.",
								MarkdownDescription: "Represents the key of node labels used to define the topology domain for Pod anti-affinityand Pod spread constraints.In K8s, a topology domain is a set of nodes that have the same value for a specific label key.Nodes with labels containing any of the specified TopologyKeys and identical values are consideredto be in the same topology domain.Note: The concept of topology in the context of K8s TopologyKeys is different from the concept oftopology in the ClusterDefinition.When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule thePod on nodes with different values for the specified TopologyKeys.This ensures that Pods are spread across different topology domains, promoting high availability andreducing the impact of node failures.Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone',are often used as TopologyKey.These keys represent the hostname and zone of a node, respectively.By including these keys in the TopologyKeys list, Pods will be spread across nodes withdifferent hostnames or zones.In addition to the well-known keys, users can also specify custom label keys as TopologyKeys.This allows for more flexible and custom topology definitions based on the specific needsof the application or environment.The TopologyKeys field is a slice of strings, where each string represents a label key.The order of the keys in the slice does not matter.",
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

					"annotations": schema.MapAttribute{
						Description:         "Specifies Annotations to override or add for underlying Pods.",
						MarkdownDescription: "Specifies Annotations to override or add for underlying Pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"comp_def": schema.StringAttribute{
						Description:         "Specifies the name of the referenced ComponentDefinition.",
						MarkdownDescription: "Specifies the name of the referenced ComponentDefinition.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(64),
						},
					},

					"configs": schema.ListNestedAttribute{
						Description:         "Specifies the configuration content of a config template.",
						MarkdownDescription: "Specifies the configuration content of a config template.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_map": schema.SingleNestedAttribute{
									Description:         "ConfigMap source for the config.",
									MarkdownDescription: "ConfigMap source for the config.",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"items": schema.ListNestedAttribute{
											Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
											MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
														Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
														MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

								"name": schema.StringAttribute{
									Description:         "The name of the config.",
									MarkdownDescription: "The name of the config.",
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

					"disable_exporter": schema.BoolAttribute{
						Description:         "Determines whether metrics exporter information is annotated on the Component's headless Service.If set to true, the following annotations will not be patched into the Service:- 'monitor.kubeblocks.io/path'- 'monitor.kubeblocks.io/port'- 'monitor.kubeblocks.io/scheme'These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
						MarkdownDescription: "Determines whether metrics exporter information is annotated on the Component's headless Service.If set to true, the following annotations will not be patched into the Service:- 'monitor.kubeblocks.io/path'- 'monitor.kubeblocks.io/port'- 'monitor.kubeblocks.io/scheme'These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enabled_logs": schema.ListAttribute{
						Description:         "Specifies which types of logs should be collected for the Cluster.The log types are defined in the 'componentDefinition.spec.logConfigs' field with the LogConfig entries.The elements in the 'enabledLogs' array correspond to the names of the LogConfig entries.For example, if the 'componentDefinition.spec.logConfigs' defines LogConfig entries withnames 'slow_query_log' and 'error_log',you can enable the collection of these logs by including their names in the 'enabledLogs' array:'''yamlenabledLogs:- slow_query_log- error_log'''",
						MarkdownDescription: "Specifies which types of logs should be collected for the Cluster.The log types are defined in the 'componentDefinition.spec.logConfigs' field with the LogConfig entries.The elements in the 'enabledLogs' array correspond to the names of the LogConfig entries.For example, if the 'componentDefinition.spec.logConfigs' defines LogConfig entries withnames 'slow_query_log' and 'error_log',you can enable the collection of these logs by including their names in the 'enabledLogs' array:'''yamlenabledLogs:- slow_query_log- error_log'''",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.ListNestedAttribute{
						Description:         "List of environment variables to add.",
						MarkdownDescription: "List of environment variables to add.",
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
									Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
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
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
											Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
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
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

					"instances": schema.ListNestedAttribute{
						Description:         "Allows for the customization of configuration values for each instance within a Component.An Instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).While instances typically share a common configuration as defined in the ClusterComponentSpec,they can require unique settings in various scenarios:For example:- A database Component might require different resource allocations for primary and secondary instances,  with primaries needing more resources.- During a rolling upgrade, a Component may first update the image for one or a few instances,  and then update the remaining instances after verifying that the updated instances are functioning correctly.InstanceTemplate allows for specifying these unique configurations per instance.Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),starting with an ordinal of 0.It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the Component.Any remaining replicas will be generated using the default template and will follow the default naming rules.",
						MarkdownDescription: "Allows for the customization of configuration values for each instance within a Component.An Instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).While instances typically share a common configuration as defined in the ClusterComponentSpec,they can require unique settings in various scenarios:For example:- A database Component might require different resource allocations for primary and secondary instances,  with primaries needing more resources.- During a rolling upgrade, a Component may first update the image for one or a few instances,  and then update the remaining instances after verifying that the updated instances are functioning correctly.InstanceTemplate allows for specifying these unique configurations per instance.Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),starting with an ordinal of 0.It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the Component.Any remaining replicas will be generated using the default template and will follow the default naming rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotations": schema.MapAttribute{
									Description:         "Specifies a map of key-value pairs to be merged into the Pod's existing annotations.Existing keys will have their values overwritten, while new keys will be added to the annotations.",
									MarkdownDescription: "Specifies a map of key-value pairs to be merged into the Pod's existing annotations.Existing keys will have their values overwritten, while new keys will be added to the annotations.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"env": schema.ListNestedAttribute{
									Description:         "Defines Env to override.Add new or override existing envs.",
									MarkdownDescription: "Defines Env to override.Add new or override existing envs.",
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
												Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
												MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
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
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
														MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
														Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
														MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
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
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
									Description:         "Specifies a map of key-value pairs that will be merged into the Pod's existing labels.Values for existing keys will be overwritten, and new keys will be added.",
									MarkdownDescription: "Specifies a map of key-value pairs that will be merged into the Pod's existing labels.Values for existing keys will be overwritten, and new keys will be added.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name specifies the unique name of the instance Pod created using this InstanceTemplate.This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinalusing the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.The specified name overrides any default naming conventions or patterns.",
									MarkdownDescription: "Name specifies the unique name of the instance Pod created using this InstanceTemplate.This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinalusing the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.The specified name overrides any default naming conventions or patterns.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(54),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"replicas": schema.Int64Attribute{
									Description:         "Specifies the number of instances (Pods) to create from this InstanceTemplate.This field allows setting how many replicated instances of the Component,with the specific overrides in the InstanceTemplate, are created.The default value is 1. A value of 0 disables instance creation.",
									MarkdownDescription: "Specifies the number of instances (Pods) to create from this InstanceTemplate.This field allows setting how many replicated instances of the Component,with the specific overrides in the InstanceTemplate, are created.The default value is 1. A value of 0 disables instance creation.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Specifies an override for the resource requirements of the first container in the Pod.This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
									MarkdownDescription: "Specifies an override for the resource requirements of the first container in the Pod.This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
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
											Description:         "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,the scheduler simply schedules this Pod onto that node, assuming that it fits resourcerequirements.",
											MarkdownDescription: "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,the scheduler simply schedules this Pod onto that node, assuming that it fits resourcerequirements.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_selector": schema.MapAttribute{
											Description:         "NodeSelector is a selector which must be true for the Pod to fit on a node.Selector which must match a node's labels for the Pod to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
											MarkdownDescription: "NodeSelector is a selector which must be true for the Pod to fit on a node.Selector which must match a node's labels for the Pod to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scheduler_name": schema.StringAttribute{
											Description:         "If specified, the Pod will be dispatched by specified scheduler.If not specified, the Pod will be dispatched by default scheduler.",
											MarkdownDescription: "If specified, the Pod will be dispatched by specified scheduler.If not specified, the Pod will be dispatched by default scheduler.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerations": schema.ListNestedAttribute{
											Description:         "Allows Pods to be scheduled onto nodes with matching taints.Each toleration in the array allows the Pod to tolerate node taints based onspecified 'key', 'value', 'effect', and 'operator'.- The 'key', 'value', and 'effect' identify the taint that the toleration matches.- The 'operator' determines how the toleration matches the taint.Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
											MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints.Each toleration in the array allows the Pod to tolerate node taints based onspecified 'key', 'value', 'effect', and 'operator'.- The 'key', 'value', and 'effect' identify the taint that the toleration matches.- The 'operator' determines how the toleration matches the taint.Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
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
											Description:         "TopologySpreadConstraints describes how a group of Pods ought to spread across topologydomains. Scheduler will schedule Pods in a way which abides by the constraints.All topologySpreadConstraints are ANDed.",
											MarkdownDescription: "TopologySpreadConstraints describes how a group of Pods ought to spread across topologydomains. Scheduler will schedule Pods in a way which abides by the constraints.All topologySpreadConstraints are ANDed.",
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

								"volume_claim_templates": schema.ListNestedAttribute{
									Description:         "Defines VolumeClaimTemplates to override.Add new or override existing volume claim templates.",
									MarkdownDescription: "Defines VolumeClaimTemplates to override.Add new or override existing volume claim templates.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refers to the name of a volumeMount defined in either:- 'componentDefinition.spec.runtime.containers[*].volumeMounts'- 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
												MarkdownDescription: "Refers to the name of a volumeMount defined in either:- 'componentDefinition.spec.runtime.containers[*].volumeMounts'- 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"spec": schema.SingleNestedAttribute{
												Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volumewith the mount name specified in the 'name' field.When a Pod is created for this ClusterComponent, a new PVC will be created based on the specificationdefined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
												MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volumewith the mount name specified in the 'name' field.When a Pod is created for this ClusterComponent, a new PVC will be created based on the specificationdefined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
												Attributes: map[string]schema.Attribute{
													"access_modes": schema.MapAttribute{
														Description:         "Contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
														MarkdownDescription: "Contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Represents the minimum resources the volume should have.If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements thatare lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
														MarkdownDescription: "Represents the minimum resources the volume should have.If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements thatare lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
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

													"storage_class_name": schema.StringAttribute{
														Description:         "The name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
														MarkdownDescription: "The name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
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
									Description:         "Defines VolumeMounts to override.Add new or override existing volume mounts of the first container in the Pod.",
									MarkdownDescription: "Defines VolumeMounts to override.Add new or override existing volume mounts of the first container in the Pod.",
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

								"volumes": schema.ListNestedAttribute{
									Description:         "Defines Volumes to override.Add new or override existing volumes.",
									MarkdownDescription: "Defines Volumes to override.Add new or override existing volumes.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"aws_elastic_block_store": schema.SingleNestedAttribute{
												Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												Attributes: map[string]schema.Attribute{
													"fs_type": schema.StringAttribute{
														Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstoreTODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstoreTODO: how do we prevent errors in the filesystem from compromising the machine",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"partition": schema.Int64Attribute{
														Description:         "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
														MarkdownDescription: "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly value true will force the readOnly setting in VolumeMounts.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"volume_id": schema.StringAttribute{
														Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
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
														Description:         "fsType is Filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is Filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
														MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
														Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_name": schema.StringAttribute{
														Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
														MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",
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
														Description:         "monitors is Required: Monitors is a collection of Ceph monitorsMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitorsMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
														Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_file": schema.StringAttribute{
														Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secretMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secretMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "user is optional: User is the rados user name, default is adminMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "user is optional: User is the rados user name, default is adminMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
												Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												Attributes: map[string]schema.Attribute{
													"fs_type": schema.StringAttribute{
														Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef is optional: points to a secret object containing parameters used to connectto OpenStack.",
														MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connectto OpenStack.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "volumeID used to identify the volume in cinder.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "volumeID used to identify the volume in cinder.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
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
														Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"items": schema.ListNestedAttribute{
														Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
																	Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																	MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
														Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "driver is the name of the CSI driver that handles this volume.Consult with your admin for the correct name as registered in the cluster.",
														MarkdownDescription: "driver is the name of the CSI driver that handles this volume.Consult with your admin for the correct name as registered in the cluster.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"fs_type": schema.StringAttribute{
														Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'.If not provided, the empty value is passed to the associated CSI driverwhich will determine the default filesystem to apply.",
														MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'.If not provided, the empty value is passed to the associated CSI driverwhich will determine the default filesystem to apply.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_publish_secret_ref": schema.SingleNestedAttribute{
														Description:         "nodePublishSecretRef is a reference to the secret object containingsensitive information to pass to the CSI driver to complete the CSINodePublishVolume and NodeUnpublishVolume calls.This field is optional, and  may be empty if no secret is required. If thesecret object contains more than one secret, all secret references are passed.",
														MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containingsensitive information to pass to the CSI driver to complete the CSINodePublishVolume and NodeUnpublishVolume calls.This field is optional, and  may be empty if no secret is required. If thesecret object contains more than one secret, all secret references are passed.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "readOnly specifies a read-only configuration for the volume.Defaults to false (read/write).",
														MarkdownDescription: "readOnly specifies a read-only configuration for the volume.Defaults to false (read/write).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"volume_attributes": schema.MapAttribute{
														Description:         "volumeAttributes stores driver-specific properties that are passed to the CSIdriver. Consult your driver's documentation for supported values.",
														MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSIdriver. Consult your driver's documentation for supported values.",
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
														Description:         "Optional: mode bits to use on created files by default. Must be aOptional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits to use on created files by default. Must be aOptional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
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
																	Description:         "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																	MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"resource_field_ref": schema.SingleNestedAttribute{
																	Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																	MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
												Description:         "emptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												Attributes: map[string]schema.Attribute{
													"medium": schema.StringAttribute{
														Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"size_limit": schema.StringAttribute{
														Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
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
												Description:         "ephemeral represents a volume that is handled by a cluster storage driver.The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,and deleted when the pod is removed.Use this if:a) the volume is only needed while the pod runs,b) features of normal volumes like restoring from snapshot or capacity   tracking are needed,c) the storage driver is specified through a storage class, andd) the storage driver supports dynamic volume provisioning through   a PersistentVolumeClaim (see EphemeralVolumeSource for more   information on the connection between this volume type   and PersistentVolumeClaim).Use PersistentVolumeClaim or one of the vendor-specificAPIs for volumes that persist for longer than the lifecycleof an individual pod.Use CSI for light-weight local ephemeral volumes if the CSI driver is meant tobe used that way - see the documentation of the driver formore information.A pod can use both types of ephemeral volumes andpersistent volumes at the same time.",
												MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver.The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,and deleted when the pod is removed.Use this if:a) the volume is only needed while the pod runs,b) features of normal volumes like restoring from snapshot or capacity   tracking are needed,c) the storage driver is specified through a storage class, andd) the storage driver supports dynamic volume provisioning through   a PersistentVolumeClaim (see EphemeralVolumeSource for more   information on the connection between this volume type   and PersistentVolumeClaim).Use PersistentVolumeClaim or one of the vendor-specificAPIs for volumes that persist for longer than the lifecycleof an individual pod.Use CSI for light-weight local ephemeral volumes if the CSI driver is meant tobe used that way - see the documentation of the driver formore information.A pod can use both types of ephemeral volumes andpersistent volumes at the same time.",
												Attributes: map[string]schema.Attribute{
													"volume_claim_template": schema.SingleNestedAttribute{
														Description:         "Will be used to create a stand-alone PVC to provision the volume.The pod in which this EphemeralVolumeSource is embedded will be theowner of the PVC, i.e. the PVC will be deleted together with thepod.  The name of the PVC will be '<pod name>-<volume name>' where'<volume name>' is the name from the 'PodSpec.Volumes' arrayentry. Pod validation will reject the pod if the concatenated nameis not valid for a PVC (for example, too long).An existing PVC with that name that is not owned by the podwill *not* be used for the pod to avoid using an unrelatedvolume by mistake. Starting the pod is then blocked untilthe unrelated PVC is removed. If such a pre-created PVC ismeant to be used by the pod, the PVC has to updated with anowner reference to the pod once the pod exists. Normallythis should not be necessary, but it may be useful whenmanually reconstructing a broken cluster.This field is read-only and no changes will be made by Kubernetesto the PVC after it has been created.Required, must not be nil.",
														MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume.The pod in which this EphemeralVolumeSource is embedded will be theowner of the PVC, i.e. the PVC will be deleted together with thepod.  The name of the PVC will be '<pod name>-<volume name>' where'<volume name>' is the name from the 'PodSpec.Volumes' arrayentry. Pod validation will reject the pod if the concatenated nameis not valid for a PVC (for example, too long).An existing PVC with that name that is not owned by the podwill *not* be used for the pod to avoid using an unrelatedvolume by mistake. Starting the pod is then blocked untilthe unrelated PVC is removed. If such a pre-created PVC ismeant to be used by the pod, the PVC has to updated with anowner reference to the pod once the pod exists. Normallythis should not be necessary, but it may be useful whenmanually reconstructing a broken cluster.This field is read-only and no changes will be made by Kubernetesto the PVC after it has been created.Required, must not be nil.",
														Attributes: map[string]schema.Attribute{
															"metadata": schema.SingleNestedAttribute{
																Description:         "May contain labels and annotations that will be copied into the PVCwhen creating it. No other fields are allowed and will be rejected duringvalidation.",
																MarkdownDescription: "May contain labels and annotations that will be copied into the PVCwhen creating it. No other fields are allowed and will be rejected duringvalidation.",
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
																Description:         "The specification for the PersistentVolumeClaim. The entire content iscopied unchanged into the PVC that gets created from thistemplate. The same fields as in a PersistentVolumeClaimare also valid here.",
																MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content iscopied unchanged into the PVC that gets created from thistemplate. The same fields as in a PersistentVolumeClaimare also valid here.",
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
														Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.TODO: how do we prevent errors in the filesystem from compromising the machine",
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
														Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
														Description:         "wwids Optional: FC volume world wide identifiers (wwids)Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
														MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids)Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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
												Description:         "flexVolume represents a generic volume resource that isprovisioned/attached using an exec based plugin.",
												MarkdownDescription: "flexVolume represents a generic volume resource that isprovisioned/attached using an exec based plugin.",
												Attributes: map[string]schema.Attribute{
													"driver": schema.StringAttribute{
														Description:         "driver is the name of the driver to use for this volume.",
														MarkdownDescription: "driver is the name of the driver to use for this volume.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"fs_type": schema.StringAttribute{
														Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
														MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
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
														Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef is Optional: secretRef is reference to the secret object containingsensitive information to pass to the plugin scripts. This may beempty if no secret object is specified. If the secret objectcontains more than one secret, all secrets are passed to the pluginscripts.",
														MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containingsensitive information to pass to the plugin scripts. This may beempty if no secret object is specified. If the secret objectcontains more than one secret, all secrets are passed to the pluginscripts.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flockershould be considered as deprecated",
														MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flockershould be considered as deprecated",
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
												Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												Attributes: map[string]schema.Attribute{
													"fs_type": schema.StringAttribute{
														Description:         "fsType is filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdiskTODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdiskTODO: how do we prevent errors in the filesystem from compromising the machine",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"partition": schema.Int64Attribute{
														Description:         "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pd_name": schema.StringAttribute{
														Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
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
												Description:         "gitRepo represents a git repository at a particular revision.DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount anEmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDirinto the Pod's container.",
												MarkdownDescription: "gitRepo represents a git repository at a particular revision.DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount anEmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDirinto the Pod's container.",
												Attributes: map[string]schema.Attribute{
													"directory": schema.StringAttribute{
														Description:         "directory is the target directory name.Must not contain or start with '..'.  If '.' is supplied, the volume directory will be thegit repository.  Otherwise, if specified, the volume will contain the git repository inthe subdirectory with the given name.",
														MarkdownDescription: "directory is the target directory name.Must not contain or start with '..'.  If '.' is supplied, the volume directory will be thegit repository.  Otherwise, if specified, the volume will contain the git repository inthe subdirectory with the given name.",
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
												Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/glusterfs/README.md",
												MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/glusterfs/README.md",
												Attributes: map[string]schema.Attribute{
													"endpoints": schema.StringAttribute{
														Description:         "endpoints is the endpoint name that details Glusterfs topology.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "path is the Glusterfs volume path.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "path is the Glusterfs volume path.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions.Defaults to false.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions.Defaults to false.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
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
												Description:         "hostPath represents a pre-existing file or directory on the hostmachine that is directly exposed to the container. This is generallyused for system agents or other privileged things that are allowedto see the host machine. Most containers will NOT need this.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath---TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can notmount host directories as read/write.",
												MarkdownDescription: "hostPath represents a pre-existing file or directory on the hostmachine that is directly exposed to the container. This is generallyused for system agents or other privileged things that are allowedto see the host machine. Most containers will NOT need this.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath---TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can notmount host directories as read/write.",
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "path of the directory on the host.If the path is a symlink, it will follow the link to the real path.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "path of the directory on the host.If the path is a symlink, it will follow the link to the real path.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type for HostPath VolumeDefaults to ''More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "type for HostPath VolumeDefaults to ''More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
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
												Description:         "iscsi represents an ISCSI Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://examples.k8s.io/volumes/iscsi/README.md",
												MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://examples.k8s.io/volumes/iscsi/README.md",
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
														Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsiTODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsiTODO: how do we prevent errors in the filesystem from compromising the machine",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"initiator_name": schema.StringAttribute{
														Description:         "initiatorName is the custom iSCSI Initiator Name.If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<target portal>:<volume name> will be created for the connection.",
														MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name.If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<target portal>:<volume name> will be created for the connection.",
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
														Description:         "iscsiInterface is the interface Name that uses an iSCSI transport.Defaults to 'default' (tcp).",
														MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport.Defaults to 'default' (tcp).",
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
														Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
														MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
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
												Description:         "name of the volume.Must be a DNS_LABEL and unique within the pod.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "name of the volume.Must be a DNS_LABEL and unique within the pod.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"nfs": schema.SingleNestedAttribute{
												Description:         "nfs represents an NFS mount on the host that shares a pod's lifetimeMore info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetimeMore info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "path that is exported by the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "path that is exported by the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly here will force the NFS export to be mounted with read-only permissions.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"server": schema.StringAttribute{
														Description:         "server is the hostname or IP address of the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "server is the hostname or IP address of the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
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
												Description:         "persistentVolumeClaimVolumeSource represents a reference to aPersistentVolumeClaim in the same namespace.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to aPersistentVolumeClaim in the same namespace.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												Attributes: map[string]schema.Attribute{
													"claim_name": schema.StringAttribute{
														Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
														MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
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
														Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
														Description:         "fSType represents the filesystem type to mountMust be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fSType represents the filesystem type to mountMust be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
														Description:         "defaultMode are the mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
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
																	Description:         "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' fieldof ClusterTrustBundle objects in an auto-updating file.Alpha, gated by the ClusterTrustBundleProjection feature gate.ClusterTrustBundle objects can either be selected by name, or by thecombination of signer name and a label selector.Kubelet performs aggressive normalization of the PEM contents writteninto the pod filesystem.  Esoteric PEM features such as inter-blockcomments and block headers are stripped.  Certificates are deduplicated.The ordering of certificates within the file is arbitrary, and Kubeletmay change the order over time.",
																	MarkdownDescription: "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' fieldof ClusterTrustBundle objects in an auto-updating file.Alpha, gated by the ClusterTrustBundleProjection feature gate.ClusterTrustBundle objects can either be selected by name, or by thecombination of signer name and a label selector.Kubelet performs aggressive normalization of the PEM contents writteninto the pod filesystem.  Esoteric PEM features such as inter-blockcomments and block headers are stripped.  Certificates are deduplicated.The ordering of certificates within the file is arbitrary, and Kubeletmay change the order over time.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "Select all ClusterTrustBundles that match this label selector.  Only haseffect if signerName is set.  Mutually-exclusive with name.  If unset,interpreted as 'match nothing'.  If set but empty, interpreted as 'matcheverything'.",
																			MarkdownDescription: "Select all ClusterTrustBundles that match this label selector.  Only haseffect if signerName is set.  Mutually-exclusive with name.  If unset,interpreted as 'match nothing'.  If set but empty, interpreted as 'matcheverything'.",
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

																		"name": schema.StringAttribute{
																			Description:         "Select a single ClusterTrustBundle by object name.  Mutually-exclusivewith signerName and labelSelector.",
																			MarkdownDescription: "Select a single ClusterTrustBundle by object name.  Mutually-exclusivewith signerName and labelSelector.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
																			Description:         "If true, don't block pod startup if the referenced ClusterTrustBundle(s)aren't available.  If using name, then the named ClusterTrustBundle isallowed not to exist.  If using signerName, then the combination ofsignerName and labelSelector is allowed to match zeroClusterTrustBundles.",
																			MarkdownDescription: "If true, don't block pod startup if the referenced ClusterTrustBundle(s)aren't available.  If using name, then the named ClusterTrustBundle isallowed not to exist.  If using signerName, then the combination ofsignerName and labelSelector is allowed to match zeroClusterTrustBundles.",
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
																			Description:         "Select all ClusterTrustBundles that match this signer name.Mutually-exclusive with name.  The contents of all selectedClusterTrustBundles will be unified and deduplicated.",
																			MarkdownDescription: "Select all ClusterTrustBundles that match this signer name.Mutually-exclusive with name.  The contents of all selectedClusterTrustBundles will be unified and deduplicated.",
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
																			Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																			MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
																						Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																						MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																			Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																						Description:         "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																						MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																						MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
																			Description:         "items if unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																			MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
																						Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																						MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																			Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																			Description:         "audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
																			MarkdownDescription: "audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"expiration_seconds": schema.Int64Attribute{
																			Description:         "expirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours.Defaults to 1 hourand must be at least 10 minutes.",
																			MarkdownDescription: "expirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours.Defaults to 1 hourand must be at least 10 minutes.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the path relative to the mount point of the file to project thetoken into.",
																			MarkdownDescription: "path is the path relative to the mount point of the file to project thetoken into.",
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
														Description:         "group to map volume access toDefault is no group",
														MarkdownDescription: "group to map volume access toDefault is no group",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions.Defaults to false.",
														MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions.Defaults to false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"registry": schema.StringAttribute{
														Description:         "registry represents a single or multiple Quobyte Registry servicesspecified as a string as host:port pair (multiple entries are separated with commas)which acts as the central registry for volumes",
														MarkdownDescription: "registry represents a single or multiple Quobyte Registry servicesspecified as a string as host:port pair (multiple entries are separated with commas)which acts as the central registry for volumes",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"tenant": schema.StringAttribute{
														Description:         "tenant owning the given Quobyte volume in the BackendUsed with dynamically provisioned Quobyte volumes, value is set by the plugin",
														MarkdownDescription: "tenant owning the given Quobyte volume in the BackendUsed with dynamically provisioned Quobyte volumes, value is set by the plugin",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user": schema.StringAttribute{
														Description:         "user to map volume access toDefaults to serivceaccount user",
														MarkdownDescription: "user to map volume access toDefaults to serivceaccount user",
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
												Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/rbd/README.md",
												MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/rbd/README.md",
												Attributes: map[string]schema.Attribute{
													"fs_type": schema.StringAttribute{
														Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#rbdTODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#rbdTODO: how do we prevent errors in the filesystem from compromising the machine",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image": schema.StringAttribute{
														Description:         "image is the rados image name.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "image is the rados image name.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"keyring": schema.StringAttribute{
														Description:         "keyring is the path to key ring for RBDUser.Default is /etc/ceph/keyring.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "keyring is the path to key ring for RBDUser.Default is /etc/ceph/keyring.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"monitors": schema.ListAttribute{
														Description:         "monitors is a collection of Ceph monitors.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "monitors is a collection of Ceph monitors.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														ElementType:         types.StringType,
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"pool": schema.StringAttribute{
														Description:         "pool is the rados pool name.Default is rbd.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "pool is the rados pool name.Default is rbd.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef is name of the authentication secret for RBDUser. If providedoverrides keyring.Default is nil.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If providedoverrides keyring.Default is nil.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "user is the rados user name.Default is admin.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "user is the rados user name.Default is admin.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
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
														Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'.Default is 'xfs'.",
														MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'.Default is 'xfs'.",
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
														Description:         "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef references to the secret for ScaleIO user and othersensitive information. If this is not provided, Login operation will fail.",
														MarkdownDescription: "secretRef references to the secret for ScaleIO user and othersensitive information. If this is not provided, Login operation will fail.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.Default is ThinProvisioned.",
														MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.Default is ThinProvisioned.",
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
														Description:         "volumeName is the name of a volume already created in the ScaleIO systemthat is associated with this volume source.",
														MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO systemthat is associated with this volume source.",
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
												Description:         "secret represents a secret that should populate this volume.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
												MarkdownDescription: "secret represents a secret that should populate this volume.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
												Attributes: map[string]schema.Attribute{
													"default_mode": schema.Int64Attribute{
														Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"items": schema.ListNestedAttribute{
														Description:         "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
																	Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																	MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
														Description:         "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
														MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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
														Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "secretRef specifies the secret to use for obtaining the StorageOS APIcredentials.  If not specified, default values will be attempted.",
														MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS APIcredentials.  If not specified, default values will be attempted.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
														Description:         "volumeName is the human-readable name of the StorageOS volume.  Volumenames are only unique within a namespace.",
														MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volumenames are only unique within a namespace.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"volume_namespace": schema.StringAttribute{
														Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If nonamespace is specified then the Pod's namespace will be used.  This allows theKubernetes name scoping to be mirrored within StorageOS for tighter integration.Set VolumeName to any name to override the default behaviour.Set to 'default' if you are not using namespaces within StorageOS.Namespaces that do not pre-exist within StorageOS will be created.",
														MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If nonamespace is specified then the Pod's namespace will be used.  This allows theKubernetes name scoping to be mirrored within StorageOS for tighter integration.Set VolumeName to any name to override the default behaviour.Set to 'default' if you are not using namespaces within StorageOS.Namespaces that do not pre-exist within StorageOS will be created.",
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
														Description:         "fsType is filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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

					"labels": schema.MapAttribute{
						Description:         "Specifies Labels to override or add for underlying Pods.",
						MarkdownDescription: "Specifies Labels to override or add for underlying Pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"offline_instances": schema.ListAttribute{
						Description:         "Specifies the names of instances to be transitioned to offline status.Marking an instance as offline results in the following:1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential   future reuse or data recovery, but it is no longer actively used.2. The ordinal number assigned to this instance is preserved, ensuring it remains unique   and avoiding conflicts with new instances.Setting instances to offline allows for a controlled scale-in process, preserving their data and maintainingordinal consistency within the Cluster.Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.",
						MarkdownDescription: "Specifies the names of instances to be transitioned to offline status.Marking an instance as offline results in the following:1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential   future reuse or data recovery, but it is no longer actively used.2. The ordinal number assigned to this instance is preserved, ensuring it remains unique   and avoiding conflicts with new instances.Setting instances to offline allows for a controlled scale-in process, preserving their data and maintainingordinal consistency within the Cluster.Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parallel_pod_management_concurrency": schema.StringAttribute{
						Description:         "Controls the concurrency of pods during initial scale up, when replacing pods on nodes,or when scaling down. It only used when 'PodManagementPolicy' is set to 'Parallel'.The default Concurrency is 100%.",
						MarkdownDescription: "Controls the concurrency of pods during initial scale up, when replacing pods on nodes,or when scaling down. It only used when 'PodManagementPolicy' is set to 'Parallel'.The default Concurrency is 100%.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_update_policy": schema.StringAttribute{
						Description:         "PodUpdatePolicy indicates how pods should be updated- 'StrictInPlace' indicates that only allows in-place upgrades.Any attempt to modify other fields will be rejected.- 'PreferInPlace' indicates that we will first attempt an in-place upgrade of the Pod.If that fails, it will fall back to the ReCreate, where pod will be recreated.Default value is 'PreferInPlace'",
						MarkdownDescription: "PodUpdatePolicy indicates how pods should be updated- 'StrictInPlace' indicates that only allows in-place upgrades.Any attempt to modify other fields will be rejected.- 'PreferInPlace' indicates that we will first attempt an in-place upgrade of the Pod.If that fails, it will fall back to the ReCreate, where pod will be recreated.Default value is 'PreferInPlace'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.",
						MarkdownDescription: "Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Specifies the resources required by the Component.It allows defining the CPU, memory requirements and limits for the Component's containers.",
						MarkdownDescription: "Specifies the resources required by the Component.It allows defining the CPU, memory requirements and limits for the Component's containers.",
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

					"runtime_class_name": schema.StringAttribute{
						Description:         "Defines runtimeClassName for all Pods managed by this Component.",
						MarkdownDescription: "Defines runtimeClassName for all Pods managed by this Component.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Description:         "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,the scheduler simply schedules this Pod onto that node, assuming that it fits resourcerequirements.",
								MarkdownDescription: "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,the scheduler simply schedules this Pod onto that node, assuming that it fits resourcerequirements.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the Pod to fit on a node.Selector which must match a node's labels for the Pod to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
								MarkdownDescription: "NodeSelector is a selector which must be true for the Pod to fit on a node.Selector which must match a node's labels for the Pod to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "If specified, the Pod will be dispatched by specified scheduler.If not specified, the Pod will be dispatched by default scheduler.",
								MarkdownDescription: "If specified, the Pod will be dispatched by specified scheduler.If not specified, the Pod will be dispatched by default scheduler.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Allows Pods to be scheduled onto nodes with matching taints.Each toleration in the array allows the Pod to tolerate node taints based onspecified 'key', 'value', 'effect', and 'operator'.- The 'key', 'value', and 'effect' identify the taint that the toleration matches.- The 'operator' determines how the toleration matches the taint.Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
								MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints.Each toleration in the array allows the Pod to tolerate node taints based onspecified 'key', 'value', 'effect', and 'operator'.- The 'key', 'value', and 'effect' identify the taint that the toleration matches.- The 'operator' determines how the toleration matches the taint.Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
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
								Description:         "TopologySpreadConstraints describes how a group of Pods ought to spread across topologydomains. Scheduler will schedule Pods in a way which abides by the constraints.All topologySpreadConstraints are ANDed.",
								MarkdownDescription: "TopologySpreadConstraints describes how a group of Pods ought to spread across topologydomains. Scheduler will schedule Pods in a way which abides by the constraints.All topologySpreadConstraints are ANDed.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "Specifies the name of the ServiceAccount required by the running Component.This ServiceAccount is used to grant necessary permissions for the Component's Pods to interactwith other Kubernetes resources, such as modifying Pod labels or sending events.Defaults:If not specified, KubeBlocks automatically assigns a default ServiceAccount named 'kb-{cluster.name}',bound to a default role defined during KubeBlocks installation.Future Changes:Future versions might change the default ServiceAccount creation strategy to one per Component,potentially revising the naming to 'kb-{cluster.name}-{component.name}'.Users can override the automatic ServiceAccount assignment by explicitly setting the name ofan existed ServiceAccount in this field.",
						MarkdownDescription: "Specifies the name of the ServiceAccount required by the running Component.This ServiceAccount is used to grant necessary permissions for the Component's Pods to interactwith other Kubernetes resources, such as modifying Pod labels or sending events.Defaults:If not specified, KubeBlocks automatically assigns a default ServiceAccount named 'kb-{cluster.name}',bound to a default role defined during KubeBlocks installation.Future Changes:Future versions might change the default ServiceAccount creation strategy to one per Component,potentially revising the naming to 'kb-{cluster.name}-{component.name}'.Users can override the automatic ServiceAccount assignment by explicitly setting the name ofan existed ServiceAccount in this field.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_refs": schema.ListNestedAttribute{
						Description:         "Defines a list of ServiceRef for a Component, enabling access to both external services andServices provided by other Clusters.Types of services:- External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;  Require a ServiceDescriptor for connection details.- Services provided by a Cluster: Managed by the same KubeBlocks operator;  identified using Cluster, Component and Service names.ServiceRefs with identical 'serviceRef.name' in the same Cluster are considered the same.Example:'''yamlserviceRefs:  - name: 'redis-sentinel'    serviceDescriptor:      name: 'external-redis-sentinel'  - name: 'postgres-cluster'    clusterServiceSelector:      cluster: 'my-postgres-cluster'      service:        component: 'postgresql''''The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.",
						MarkdownDescription: "Defines a list of ServiceRef for a Component, enabling access to both external services andServices provided by other Clusters.Types of services:- External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;  Require a ServiceDescriptor for connection details.- Services provided by a Cluster: Managed by the same KubeBlocks operator;  identified using Cluster, Component and Service names.ServiceRefs with identical 'serviceRef.name' in the same Cluster are considered the same.Example:'''yamlserviceRefs:  - name: 'redis-sentinel'    serviceDescriptor:      name: 'external-redis-sentinel'  - name: 'postgres-cluster'    clusterServiceSelector:      cluster: 'my-postgres-cluster'      service:        component: 'postgresql''''The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster": schema.StringAttribute{
									Description:         "Specifies the name of the KubeBlocks Cluster being referenced.This is used when services from another KubeBlocks Cluster are consumed.By default, the referenced KubeBlocks Cluster's 'clusterDefinition.spec.connectionCredential'will be utilized to bind to the current Component. This credential should include:'endpoint', 'port', 'username', and 'password'.Note:- The 'ServiceKind' and 'ServiceVersion' specified in the service reference within the  ClusterDefinition are not validated when using this approach.- If both 'cluster' and 'serviceDescriptor' are present, 'cluster' will take precedence.Deprecated since v0.9 since 'clusterDefinition.spec.connectionCredential' is deprecated,use 'clusterServiceSelector' instead.This field is maintained for backward compatibility and its use is discouraged.Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
									MarkdownDescription: "Specifies the name of the KubeBlocks Cluster being referenced.This is used when services from another KubeBlocks Cluster are consumed.By default, the referenced KubeBlocks Cluster's 'clusterDefinition.spec.connectionCredential'will be utilized to bind to the current Component. This credential should include:'endpoint', 'port', 'username', and 'password'.Note:- The 'ServiceKind' and 'ServiceVersion' specified in the service reference within the  ClusterDefinition are not validated when using this approach.- If both 'cluster' and 'serviceDescriptor' are present, 'cluster' will take precedence.Deprecated since v0.9 since 'clusterDefinition.spec.connectionCredential' is deprecated,use 'clusterServiceSelector' instead.This field is maintained for backward compatibility and its use is discouraged.Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"cluster_service_selector": schema.SingleNestedAttribute{
									Description:         "References a service provided by another KubeBlocks Cluster.It specifies the ClusterService and the account credentials needed for access.",
									MarkdownDescription: "References a service provided by another KubeBlocks Cluster.It specifies the ClusterService and the account credentials needed for access.",
									Attributes: map[string]schema.Attribute{
										"cluster": schema.StringAttribute{
											Description:         "The name of the Cluster being referenced.",
											MarkdownDescription: "The name of the Cluster being referenced.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"credential": schema.SingleNestedAttribute{
											Description:         "Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster.The SystemAccount should be defined in 'componentDefinition.spec.systemAccounts'of the Component providing the service in the referenced Cluster.",
											MarkdownDescription: "Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster.The SystemAccount should be defined in 'componentDefinition.spec.systemAccounts'of the Component providing the service in the referenced Cluster.",
											Attributes: map[string]schema.Attribute{
												"component": schema.StringAttribute{
													Description:         "The name of the Component where the credential resides in.",
													MarkdownDescription: "The name of the Component where the credential resides in.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The name of the credential (SystemAccount) to reference.",
													MarkdownDescription: "The name of the credential (SystemAccount) to reference.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"service": schema.SingleNestedAttribute{
											Description:         "Identifies a ClusterService from the list of Services defined in 'cluster.spec.services' of the referenced Cluster.",
											MarkdownDescription: "Identifies a ClusterService from the list of Services defined in 'cluster.spec.services' of the referenced Cluster.",
											Attributes: map[string]schema.Attribute{
												"component": schema.StringAttribute{
													Description:         "The name of the Component where the Service resides in.It is required when referencing a Component's Service.",
													MarkdownDescription: "The name of the Component where the Service resides in.It is required when referencing a Component's Service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.StringAttribute{
													Description:         "The port name of the Service to be referenced.If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...",
													MarkdownDescription: "The port name of the Service to be referenced.If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service": schema.StringAttribute{
													Description:         "The name of the Service to be referenced.Leave it empty to reference the default Service. Set it to 'headless' to reference the default headless Service.If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,and the resolved value will be presented in the following format: service1.name,service2.name...",
													MarkdownDescription: "The name of the Service to be referenced.Leave it empty to reference the default Service. Set it to 'headless' to reference the default headless Service.If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,and the resolved value will be presented in the following format: service1.name,service2.name...",
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

								"name": schema.StringAttribute{
									Description:         "Specifies the identifier of the service reference declaration.It corresponds to the serviceRefDeclaration name defined in either:- 'componentDefinition.spec.serviceRefDeclarations[*].name'- 'clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name' (deprecated)",
									MarkdownDescription: "Specifies the identifier of the service reference declaration.It corresponds to the serviceRefDeclaration name defined in either:- 'componentDefinition.spec.serviceRefDeclarations[*].name'- 'clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name' (deprecated)",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the currentCluster by default.",
									MarkdownDescription: "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the currentCluster by default.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service_descriptor": schema.StringAttribute{
									Description:         "Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.When referencing a service provided by external sources, a ServiceDescriptor object is required to establishthe service binding.The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKindand serviceVersion declared in the definition.If both 'cluster' and 'serviceDescriptor' are specified, the 'cluster' takes precedence.",
									MarkdownDescription: "Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.When referencing a service provided by external sources, a ServiceDescriptor object is required to establishthe service binding.The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKindand serviceVersion declared in the definition.If both 'cluster' and 'serviceDescriptor' are specified, the 'cluster' takes precedence.",
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

					"service_version": schema.StringAttribute{
						Description:         "ServiceVersion specifies the version of the Service expected to be provisioned by this Component.The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/).",
						MarkdownDescription: "ServiceVersion specifies the version of the Service expected to be provisioned by this Component.The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/).",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(32),
						},
					},

					"services": schema.ListNestedAttribute{
						Description:         "Overrides Services defined in referenced ComponentDefinition and exposes endpoints that can be accessedby clients.",
						MarkdownDescription: "Overrides Services defined in referenced ComponentDefinition and exposes endpoints that can be accessedby clients.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotations": schema.MapAttribute{
									Description:         "If ServiceType is LoadBalancer, cloud provider related parameters can be put hereMore info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
									MarkdownDescription: "If ServiceType is LoadBalancer, cloud provider related parameters can be put hereMore info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"disable_auto_provision": schema.BoolAttribute{
									Description:         "Indicates whether the automatic provisioning of the service should be disabled.If set to true, the service will not be automatically created at the component provisioning.Instead, you can enable the creation of this service by specifying it explicitly in the cluster API.",
									MarkdownDescription: "Indicates whether the automatic provisioning of the service should be disabled.If set to true, the service will not be automatically created at the component provisioning.Instead, you can enable the creation of this service by specifying it explicitly in the cluster API.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name defines the name of the service.otherwise, it indicates the name of the service.Others can refer to this service by its name. (e.g., connection credential)Cannot be updated.",
									MarkdownDescription: "Name defines the name of the service.otherwise, it indicates the name of the service.Others can refer to this service by its name. (e.g., connection credential)Cannot be updated.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(25),
									},
								},

								"pod_service": schema.BoolAttribute{
									Description:         "Indicates whether to create a corresponding Service for each Pod of the selected Component.When set to true, a set of Services will be automatically generated for each Pod,and the 'roleSelector' field will be ignored.The names of the generated Services will follow the same suffix naming pattern: '$(serviceName)-$(podOrdinal)'.The total number of generated Services will be equal to the number of replicas specified for the Component.Example usage:'''yamlname: my-serviceserviceName: my-servicepodService: truedisableAutoProvision: truespec:  type: NodePort  ports:  - name: http    port: 80    targetPort: 8080'''In this example, if the Component has 3 replicas, three Services will be generated:- my-service-0: Points to the first Pod (podOrdinal: 0)- my-service-1: Points to the second Pod (podOrdinal: 1)- my-service-2: Points to the third Pod (podOrdinal: 2)Each generated Service will have the specified spec configuration and will target its respective Pod.This feature is useful when you need to expose each Pod of a Component individually, allowing external accessto specific instances of the Component.",
									MarkdownDescription: "Indicates whether to create a corresponding Service for each Pod of the selected Component.When set to true, a set of Services will be automatically generated for each Pod,and the 'roleSelector' field will be ignored.The names of the generated Services will follow the same suffix naming pattern: '$(serviceName)-$(podOrdinal)'.The total number of generated Services will be equal to the number of replicas specified for the Component.Example usage:'''yamlname: my-serviceserviceName: my-servicepodService: truedisableAutoProvision: truespec:  type: NodePort  ports:  - name: http    port: 80    targetPort: 8080'''In this example, if the Component has 3 replicas, three Services will be generated:- my-service-0: Points to the first Pod (podOrdinal: 0)- my-service-1: Points to the second Pod (podOrdinal: 1)- my-service-2: Points to the third Pod (podOrdinal: 2)Each generated Service will have the specified spec configuration and will target its respective Pod.This feature is useful when you need to expose each Pod of a Component individually, allowing external accessto specific instances of the Component.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"role_selector": schema.StringAttribute{
									Description:         "Extends the above 'serviceSpec.selector' by allowing you to specify defined role as selector for the service.When 'roleSelector' is set, it adds a label selector 'kubeblocks.io/role: {roleSelector}'to the 'serviceSpec.selector'.Example usage:	  roleSelector: 'leader'In this example, setting 'roleSelector' to 'leader' will add a label selector'kubeblocks.io/role: leader' to the 'serviceSpec.selector'.This means that the service will select and route traffic to Pods with the label'kubeblocks.io/role' set to 'leader'.Note that if 'podService' sets to true, RoleSelector will be ignored.The 'podService' flag takes precedence over 'roleSelector' and generates a service for each Pod.",
									MarkdownDescription: "Extends the above 'serviceSpec.selector' by allowing you to specify defined role as selector for the service.When 'roleSelector' is set, it adds a label selector 'kubeblocks.io/role: {roleSelector}'to the 'serviceSpec.selector'.Example usage:	  roleSelector: 'leader'In this example, setting 'roleSelector' to 'leader' will add a label selector'kubeblocks.io/role: leader' to the 'serviceSpec.selector'.This means that the service will select and route traffic to Pods with the label'kubeblocks.io/role' set to 'leader'.Note that if 'podService' sets to true, RoleSelector will be ignored.The 'podService' flag takes precedence over 'roleSelector' and generates a service for each Pod.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service_name": schema.StringAttribute{
									Description:         "ServiceName defines the name of the underlying service object.If not specified, the default service name with different patterns will be used:- CLUSTER_NAME: for cluster-level services- CLUSTER_NAME-COMPONENT_NAME: for component-level servicesOnly one default service name is allowed.Cannot be updated.",
									MarkdownDescription: "ServiceName defines the name of the underlying service object.If not specified, the default service name with different patterns will be used:- CLUSTER_NAME: for cluster-level services- CLUSTER_NAME-COMPONENT_NAME: for component-level servicesOnly one default service name is allowed.Cannot be updated.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(25),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"spec": schema.SingleNestedAttribute{
									Description:         "Spec defines the behavior of a service.https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
									MarkdownDescription: "Spec defines the behavior of a service.https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
									Attributes: map[string]schema.Attribute{
										"allocate_load_balancer_node_ports": schema.BoolAttribute{
											Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automaticallyallocated for services with type LoadBalancer.  Default is 'true'. Itmay be set to 'false' if the cluster load-balancer does not rely onNodePorts.  If the caller requests specific NodePorts (by specifying avalue), those requests will be respected, regardless of this field.This field may only be set for services with type LoadBalancer and willbe cleared if the type is changed to any other type.",
											MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automaticallyallocated for services with type LoadBalancer.  Default is 'true'. Itmay be set to 'false' if the cluster load-balancer does not rely onNodePorts.  If the caller requests specific NodePorts (by specifying avalue), those requests will be respected, regardless of this field.This field may only be set for services with type LoadBalancer and willbe cleared if the type is changed to any other type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cluster_ip": schema.StringAttribute{
											Description:         "clusterIP is the IP address of the service and is usually assignedrandomly. If an address is specified manually, is in-range (as persystem configuration), and is not in use, it will be allocated to theservice; otherwise creation of the service will fail. This field may notbe changed through updates unless the type field is also being changedto ExternalName (which requires this field to be blank) or the typefield is being changed from ExternalName (in which case this field mayoptionally be specified, as describe above).  Valid values are 'None',empty string (''), or a valid IP address. Setting this to 'None' makes a'headless service' (no virtual IP), which is useful when direct endpointconnections are preferred and proxying is not required.  Only applies totypes ClusterIP, NodePort, and LoadBalancer. If this field is specifiedwhen creating a Service of type ExternalName, creation will fail. Thisfield will be wiped when updating a Service to type ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											MarkdownDescription: "clusterIP is the IP address of the service and is usually assignedrandomly. If an address is specified manually, is in-range (as persystem configuration), and is not in use, it will be allocated to theservice; otherwise creation of the service will fail. This field may notbe changed through updates unless the type field is also being changedto ExternalName (which requires this field to be blank) or the typefield is being changed from ExternalName (in which case this field mayoptionally be specified, as describe above).  Valid values are 'None',empty string (''), or a valid IP address. Setting this to 'None' makes a'headless service' (no virtual IP), which is useful when direct endpointconnections are preferred and proxying is not required.  Only applies totypes ClusterIP, NodePort, and LoadBalancer. If this field is specifiedwhen creating a Service of type ExternalName, creation will fail. Thisfield will be wiped when updating a Service to type ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cluster_i_ps": schema.ListAttribute{
											Description:         "ClusterIPs is a list of IP addresses assigned to this service, and areusually assigned randomly.  If an address is specified manually, isin-range (as per system configuration), and is not in use, it will beallocated to the service; otherwise creation of the service will fail.This field may not be changed through updates unless the type field isalso being changed to ExternalName (which requires this field to beempty) or the type field is being changed from ExternalName (in whichcase this field may optionally be specified, as describe above).  Validvalues are 'None', empty string (''), or a valid IP address.  Settingthis to 'None' makes a 'headless service' (no virtual IP), which isuseful when direct endpoint connections are preferred and proxying isnot required.  Only applies to types ClusterIP, NodePort, andLoadBalancer. If this field is specified when creating a Service of typeExternalName, creation will fail. This field will be wiped when updatinga Service to type ExternalName.  If this field is not specified, it willbe initialized from the clusterIP field.  If this field is specified,clients must ensure that clusterIPs[0] and clusterIP have the samevalue.This field may hold a maximum of two entries (dual-stack IPs, in either order).These IPs must correspond to the values of the ipFamilies field. BothclusterIPs and ipFamilies are governed by the ipFamilyPolicy field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and areusually assigned randomly.  If an address is specified manually, isin-range (as per system configuration), and is not in use, it will beallocated to the service; otherwise creation of the service will fail.This field may not be changed through updates unless the type field isalso being changed to ExternalName (which requires this field to beempty) or the type field is being changed from ExternalName (in whichcase this field may optionally be specified, as describe above).  Validvalues are 'None', empty string (''), or a valid IP address.  Settingthis to 'None' makes a 'headless service' (no virtual IP), which isuseful when direct endpoint connections are preferred and proxying isnot required.  Only applies to types ClusterIP, NodePort, andLoadBalancer. If this field is specified when creating a Service of typeExternalName, creation will fail. This field will be wiped when updatinga Service to type ExternalName.  If this field is not specified, it willbe initialized from the clusterIP field.  If this field is specified,clients must ensure that clusterIPs[0] and clusterIP have the samevalue.This field may hold a maximum of two entries (dual-stack IPs, in either order).These IPs must correspond to the values of the ipFamilies field. BothclusterIPs and ipFamilies are governed by the ipFamilyPolicy field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"external_i_ps": schema.ListAttribute{
											Description:         "externalIPs is a list of IP addresses for which nodes in the clusterwill also accept traffic for this service.  These IPs are not managed byKubernetes.  The user is responsible for ensuring that traffic arrivesat a node with this IP.  A common example is external load-balancersthat are not part of the Kubernetes system.",
											MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the clusterwill also accept traffic for this service.  These IPs are not managed byKubernetes.  The user is responsible for ensuring that traffic arrivesat a node with this IP.  A common example is external load-balancersthat are not part of the Kubernetes system.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"external_name": schema.StringAttribute{
											Description:         "externalName is the external reference that discovery mechanisms willreturn as an alias for this service (e.g. a DNS CNAME record). Noproxying will be involved.  Must be a lowercase RFC-1123 hostname(https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
											MarkdownDescription: "externalName is the external reference that discovery mechanisms willreturn as an alias for this service (e.g. a DNS CNAME record). Noproxying will be involved.  Must be a lowercase RFC-1123 hostname(https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"external_traffic_policy": schema.StringAttribute{
											Description:         "externalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts,ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configurethe service in a way that assumes that external load balancers will take careof balancing the service traffic between nodes, and so each node will delivertraffic only to the node-local endpoints of the service, without masqueradingthe client source IP. (Traffic mistakenly sent to a node with no endpoints willbe dropped.) The default value, 'Cluster', uses the standard behavior ofrouting to all endpoints evenly (possibly modified by topology and otherfeatures). Note that traffic sent to an External IP or LoadBalancer IP fromwithin the cluster will always get 'Cluster' semantics, but clients sending toa NodePort from within the cluster may need to take traffic policy into accountwhen picking a node.",
											MarkdownDescription: "externalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts,ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configurethe service in a way that assumes that external load balancers will take careof balancing the service traffic between nodes, and so each node will delivertraffic only to the node-local endpoints of the service, without masqueradingthe client source IP. (Traffic mistakenly sent to a node with no endpoints willbe dropped.) The default value, 'Cluster', uses the standard behavior ofrouting to all endpoints evenly (possibly modified by topology and otherfeatures). Note that traffic sent to an External IP or LoadBalancer IP fromwithin the cluster will always get 'Cluster' semantics, but clients sending toa NodePort from within the cluster may need to take traffic policy into accountwhen picking a node.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"health_check_node_port": schema.Int64Attribute{
											Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service.This only applies when type is set to LoadBalancer andexternalTrafficPolicy is set to Local. If a value is specified, isin-range, and is not in use, it will be used.  If not specified, a valuewill be automatically allocated.  External systems (e.g. load-balancers)can use this port to determine if a given node holds endpoints for thisservice or not.  If this field is specified when creating a Servicewhich does not need it, creation will fail. This field will be wipedwhen updating a Service to no longer need it (e.g. changing type).This field cannot be updated once set.",
											MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service.This only applies when type is set to LoadBalancer andexternalTrafficPolicy is set to Local. If a value is specified, isin-range, and is not in use, it will be used.  If not specified, a valuewill be automatically allocated.  External systems (e.g. load-balancers)can use this port to determine if a given node holds endpoints for thisservice or not.  If this field is specified when creating a Servicewhich does not need it, creation will fail. This field will be wipedwhen updating a Service to no longer need it (e.g. changing type).This field cannot be updated once set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"internal_traffic_policy": schema.StringAttribute{
											Description:         "InternalTrafficPolicy describes how nodes distribute service traffic theyreceive on the ClusterIP. If set to 'Local', the proxy will assume that podsonly want to talk to endpoints of the service on the same node as the pod,dropping the traffic if there are no local endpoints. The default value,'Cluster', uses the standard behavior of routing to all endpoints evenly(possibly modified by topology and other features).",
											MarkdownDescription: "InternalTrafficPolicy describes how nodes distribute service traffic theyreceive on the ClusterIP. If set to 'Local', the proxy will assume that podsonly want to talk to endpoints of the service on the same node as the pod,dropping the traffic if there are no local endpoints. The default value,'Cluster', uses the standard behavior of routing to all endpoints evenly(possibly modified by topology and other features).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip_families": schema.ListAttribute{
											Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to thisservice. This field is usually assigned automatically based on clusterconfiguration and the ipFamilyPolicy field. If this field is specifiedmanually, the requested family is available in the cluster,and ipFamilyPolicy allows it, it will be used; otherwise creation ofthe service will fail. This field is conditionally mutable: it allowsfor adding or removing a secondary IP family, but it does not allowchanging the primary IP family of the Service. Valid values are 'IPv4'and 'IPv6'.  This field only applies to Services of types ClusterIP,NodePort, and LoadBalancer, and does apply to 'headless' services.This field will be wiped when updating a Service to type ExternalName.This field may hold a maximum of two entries (dual-stack families, ineither order).  These families must correspond to the values of theclusterIPs field, if specified. Both clusterIPs and ipFamilies aregoverned by the ipFamilyPolicy field.",
											MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to thisservice. This field is usually assigned automatically based on clusterconfiguration and the ipFamilyPolicy field. If this field is specifiedmanually, the requested family is available in the cluster,and ipFamilyPolicy allows it, it will be used; otherwise creation ofthe service will fail. This field is conditionally mutable: it allowsfor adding or removing a secondary IP family, but it does not allowchanging the primary IP family of the Service. Valid values are 'IPv4'and 'IPv6'.  This field only applies to Services of types ClusterIP,NodePort, and LoadBalancer, and does apply to 'headless' services.This field will be wiped when updating a Service to type ExternalName.This field may hold a maximum of two entries (dual-stack families, ineither order).  These families must correspond to the values of theclusterIPs field, if specified. Both clusterIPs and ipFamilies aregoverned by the ipFamilyPolicy field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip_family_policy": schema.StringAttribute{
											Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required bythis Service. If there is no value provided, then this field will be setto SingleStack. Services can be 'SingleStack' (a single IP family),'PreferDualStack' (two IP families on dual-stack configured clusters ora single IP family on single-stack clusters), or 'RequireDualStack'(two IP families on dual-stack configured clusters, otherwise fail). TheipFamilies and clusterIPs fields depend on the value of this field. Thisfield will be wiped when updating a service to type ExternalName.",
											MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required bythis Service. If there is no value provided, then this field will be setto SingleStack. Services can be 'SingleStack' (a single IP family),'PreferDualStack' (two IP families on dual-stack configured clusters ora single IP family on single-stack clusters), or 'RequireDualStack'(two IP families on dual-stack configured clusters, otherwise fail). TheipFamilies and clusterIPs fields depend on the value of this field. Thisfield will be wiped when updating a service to type ExternalName.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"load_balancer_class": schema.StringAttribute{
											Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to.If specified, the value of this field must be a label-style identifier, with an optional prefix,e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users.This field can only be set when the Service type is 'LoadBalancer'. If not set, the default loadbalancer implementation is used, today this is typically done through the cloud provider integration,but should apply for any default implementation. If set, it is assumed that a load balancerimplementation is watching for Services with a matching class. Any default load balancerimplementation (e.g. cloud providers) should ignore Services that set this field.This field can only be set when creating or updating a Service to type 'LoadBalancer'.Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
											MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to.If specified, the value of this field must be a label-style identifier, with an optional prefix,e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users.This field can only be set when the Service type is 'LoadBalancer'. If not set, the default loadbalancer implementation is used, today this is typically done through the cloud provider integration,but should apply for any default implementation. If set, it is assumed that a load balancerimplementation is watching for Services with a matching class. Any default load balancerimplementation (e.g. cloud providers) should ignore Services that set this field.This field can only be set when creating or updating a Service to type 'LoadBalancer'.Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"load_balancer_ip": schema.StringAttribute{
											Description:         "Only applies to Service Type: LoadBalancer.This feature depends on whether the underlying cloud-provider supports specifyingthe loadBalancerIP when a load balancer is created.This field will be ignored if the cloud-provider does not support the feature.Deprecated: This field was under-specified and its meaning varies across implementations.Using it is non-portable and it may not support dual-stack.Users are encouraged to use implementation-specific annotations when available.",
											MarkdownDescription: "Only applies to Service Type: LoadBalancer.This feature depends on whether the underlying cloud-provider supports specifyingthe loadBalancerIP when a load balancer is created.This field will be ignored if the cloud-provider does not support the feature.Deprecated: This field was under-specified and its meaning varies across implementations.Using it is non-portable and it may not support dual-stack.Users are encouraged to use implementation-specific annotations when available.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"load_balancer_source_ranges": schema.ListAttribute{
											Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-providerload-balancer will be restricted to the specified client IPs. This field will be ignored if thecloud-provider does not support the feature.'More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
											MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-providerload-balancer will be restricted to the specified client IPs. This field will be ignored if thecloud-provider does not support the feature.'More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "The list of ports that are exposed by this service.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											MarkdownDescription: "The list of ports that are exposed by this service.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"app_protocol": schema.StringAttribute{
														Description:         "The application protocol for this port.This is used as a hint for implementations to offer richer behavior for protocols that they understand.This field follows standard Kubernetes label syntax.Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as perRFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455* Other protocols should use implementation-defined prefixed names such asmycompany.com/my-custom-protocol.",
														MarkdownDescription: "The application protocol for this port.This is used as a hint for implementations to offer richer behavior for protocols that they understand.This field follows standard Kubernetes label syntax.Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as perRFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455* Other protocols should use implementation-defined prefixed names such asmycompany.com/my-custom-protocol.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of this port within the service. This must be a DNS_LABEL.All ports within a ServiceSpec must have unique names. When consideringthe endpoints for a Service, this must match the 'name' field in theEndpointPort.Optional if only one ServicePort is defined on this service.",
														MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL.All ports within a ServiceSpec must have unique names. When consideringthe endpoints for a Service, this must match the 'name' field in theEndpointPort.Optional if only one ServicePort is defined on this service.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_port": schema.Int64Attribute{
														Description:         "The port on each node on which this service is exposed when type isNodePort or LoadBalancer.  Usually assigned by the system. If a value isspecified, in-range, and not in use it will be used, otherwise theoperation will fail.  If not specified, a port will be allocated if thisService requires one.  If this field is specified when creating aService which does not need it, creation will fail. This field will bewiped when updating a Service to no longer need it (e.g. changing typefrom NodePort to ClusterIP).More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
														MarkdownDescription: "The port on each node on which this service is exposed when type isNodePort or LoadBalancer.  Usually assigned by the system. If a value isspecified, in-range, and not in use it will be used, otherwise theoperation will fail.  If not specified, a port will be allocated if thisService requires one.  If this field is specified when creating aService which does not need it, creation will fail. This field will bewiped when updating a Service to no longer need it (e.g. changing typefrom NodePort to ClusterIP).More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
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
														Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'.Default is TCP.",
														MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'.Default is TCP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_port": schema.StringAttribute{
														Description:         "Number or name of the port to access on the pods targeted by the service.Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.If this is a string, it will be looked up as a named port in thetarget Pod's container ports. If this is not specified, the valueof the 'port' field is used (an identity map).This field is ignored for services with clusterIP=None, and should beomitted or set equal to the 'port' field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
														MarkdownDescription: "Number or name of the port to access on the pods targeted by the service.Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.If this is a string, it will be looked up as a named port in thetarget Pod's container ports. If this is not specified, the valueof the 'port' field is used (an identity map).This field is ignored for services with clusterIP=None, and should beomitted or set equal to the 'port' field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

										"publish_not_ready_addresses": schema.BoolAttribute{
											Description:         "publishNotReadyAddresses indicates that any agent which deals with endpoints for thisService should disregard any indications of ready/not-ready.The primary use case for setting this field is for a StatefulSet's Headless Service topropagate SRV DNS records for its Pods for the purpose of peer discovery.The Kubernetes controllers that generate Endpoints and EndpointSlice resources forServices interpret this to mean that all endpoints are considered 'ready' even if thePods themselves are not. Agents which consume only Kubernetes generated endpointsthrough the Endpoints or EndpointSlice resources can safely assume this behavior.",
											MarkdownDescription: "publishNotReadyAddresses indicates that any agent which deals with endpoints for thisService should disregard any indications of ready/not-ready.The primary use case for setting this field is for a StatefulSet's Headless Service topropagate SRV DNS records for its Pods for the purpose of peer discovery.The Kubernetes controllers that generate Endpoints and EndpointSlice resources forServices interpret this to mean that all endpoints are considered 'ready' even if thePods themselves are not. Agents which consume only Kubernetes generated endpointsthrough the Endpoints or EndpointSlice resources can safely assume this behavior.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.MapAttribute{
											Description:         "Route service traffic to pods with label keys and values matching thisselector. If empty or not present, the service is assumed to have anexternal process managing its endpoints, which Kubernetes will notmodify. Only applies to types ClusterIP, NodePort, and LoadBalancer.Ignored if type is ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
											MarkdownDescription: "Route service traffic to pods with label keys and values matching thisselector. If empty or not present, the service is assumed to have anexternal process managing its endpoints, which Kubernetes will notmodify. Only applies to types ClusterIP, NodePort, and LoadBalancer.Ignored if type is ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"session_affinity": schema.StringAttribute{
											Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity.Enable client IP based session affinity.Must be ClientIP or None.Defaults to None.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity.Enable client IP based session affinity.Must be ClientIP or None.Defaults to None.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"session_affinity_config": schema.SingleNestedAttribute{
											Description:         "sessionAffinityConfig contains the configurations of session affinity.",
											MarkdownDescription: "sessionAffinityConfig contains the configurations of session affinity.",
											Attributes: map[string]schema.Attribute{
												"client_ip": schema.SingleNestedAttribute{
													Description:         "clientIP contains the configurations of Client IP based session affinity.",
													MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",
													Attributes: map[string]schema.Attribute{
														"timeout_seconds": schema.Int64Attribute{
															Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time.The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'.Default value is 10800(for 3 hours).",
															MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time.The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'.Default value is 10800(for 3 hours).",
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

										"type": schema.StringAttribute{
											Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Validoptions are ExternalName, ClusterIP, NodePort, and LoadBalancer.'ClusterIP' allocates a cluster-internal IP address for load-balancingto endpoints. Endpoints are determined by the selector or if that is notspecified, by manual construction of an Endpoints object orEndpointSlice objects. If clusterIP is 'None', no virtual IP isallocated and the endpoints are published as a set of endpoints ratherthan a virtual IP.'NodePort' builds on ClusterIP and allocates a port on every node whichroutes to the same endpoints as the clusterIP.'LoadBalancer' builds on NodePort and creates an external load-balancer(if supported in the current cloud) which routes to the same endpointsas the clusterIP.'ExternalName' aliases this service to the specified externalName.Several other fields do not apply to ExternalName services.More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
											MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Validoptions are ExternalName, ClusterIP, NodePort, and LoadBalancer.'ClusterIP' allocates a cluster-internal IP address for load-balancingto endpoints. Endpoints are determined by the selector or if that is notspecified, by manual construction of an Endpoints object orEndpointSlice objects. If clusterIP is 'None', no virtual IP isallocated and the endpoints are published as a set of endpoints ratherthan a virtual IP.'NodePort' builds on ClusterIP and allocates a port on every node whichroutes to the same endpoints as the clusterIP.'LoadBalancer' builds on NodePort and creates an external load-balancer(if supported in the current cloud) which routes to the same endpointsas the clusterIP.'ExternalName' aliases this service to the specified externalName.Several other fields do not apply to ExternalName services.More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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

					"stop": schema.BoolAttribute{
						Description:         "Stop the Component.If set, all the computing resources will be released.",
						MarkdownDescription: "Stop the Component.If set, all the computing resources will be released.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"system_accounts": schema.ListNestedAttribute{
						Description:         "Overrides system accounts defined in referenced ComponentDefinition.",
						MarkdownDescription: "Overrides system accounts defined in referenced ComponentDefinition.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of the system account.",
									MarkdownDescription: "The name of the system account.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"password_config": schema.SingleNestedAttribute{
									Description:         "Specifies the policy for generating the account's password.This field is immutable once set.",
									MarkdownDescription: "Specifies the policy for generating the account's password.This field is immutable once set.",
									Attributes: map[string]schema.Attribute{
										"length": schema.Int64Attribute{
											Description:         "The length of the password.",
											MarkdownDescription: "The length of the password.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(8),
												int64validator.AtMost(32),
											},
										},

										"letter_case": schema.StringAttribute{
											Description:         "The case of the letters in the password.",
											MarkdownDescription: "The case of the letters in the password.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("LowerCases", "UpperCases", "MixedCases"),
											},
										},

										"num_digits": schema.Int64Attribute{
											Description:         "The number of digits in the password.",
											MarkdownDescription: "The number of digits in the password.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(8),
											},
										},

										"num_symbols": schema.Int64Attribute{
											Description:         "The number of symbols in the password.",
											MarkdownDescription: "The number of symbols in the password.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(8),
											},
										},

										"seed": schema.StringAttribute{
											Description:         "Seed to generate the account's password.Cannot be updated.",
											MarkdownDescription: "Seed to generate the account's password.Cannot be updated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "Refers to the secret from which data will be copied to create the new account.This field is immutable once set.",
									MarkdownDescription: "Refers to the secret from which data will be copied to create the new account.This field is immutable once set.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "The unique identifier of the secret.",
											MarkdownDescription: "The unique identifier of the secret.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "The namespace where the secret is located.",
											MarkdownDescription: "The namespace where the secret is located.",
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

					"tls_config": schema.SingleNestedAttribute{
						Description:         "Specifies the TLS configuration for the Component, including:- A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication.- An optional field that specifies the configuration for the TLS certificates issuer when TLS is enabled.  It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.	 The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.",
						MarkdownDescription: "Specifies the TLS configuration for the Component, including:- A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication.- An optional field that specifies the configuration for the TLS certificates issuer when TLS is enabled.  It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.	 The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)for secure communication.When set to true, the Component will be configured to use TLS encryption for its network connections.This ensures that the data transmitted between the Component and its clients or other Components is encryptedand protected from unauthorized access.If TLS is enabled, the Component may require additional configuration,such as specifying TLS certificates and keys, to properly set up the secure communication channel.",
								MarkdownDescription: "A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)for secure communication.When set to true, the Component will be configured to use TLS encryption for its network connections.This ensures that the data transmitted between the Component and its clients or other Components is encryptedand protected from unauthorized access.If TLS is enabled, the Component may require additional configuration,such as specifying TLS certificates and keys, to properly set up the secure communication channel.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"issuer": schema.SingleNestedAttribute{
								Description:         "Specifies the configuration for the TLS certificates issuer.It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.Required when TLS is enabled.",
								MarkdownDescription: "Specifies the configuration for the TLS certificates issuer.It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.Required when TLS is enabled.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "The issuer for TLS certificates.It only allows two enum values: 'KubeBlocks' and 'UserProvided'.- 'KubeBlocks' indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used.- 'UserProvided' means that the user is responsible for providing their own CA, Cert, and Key.  In this case, the user-provided CA certificate, server certificate, and private key will be used  for TLS communication.",
										MarkdownDescription: "The issuer for TLS certificates.It only allows two enum values: 'KubeBlocks' and 'UserProvided'.- 'KubeBlocks' indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used.- 'UserProvided' means that the user is responsible for providing their own CA, Cert, and Key.  In this case, the user-provided CA certificate, server certificate, and private key will be used  for TLS communication.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"secret_ref": schema.SingleNestedAttribute{
										Description:         "SecretRef is the reference to the secret that contains user-provided certificates.It is required when the issuer is set to 'UserProvided'.",
										MarkdownDescription: "SecretRef is the reference to the secret that contains user-provided certificates.It is required when the issuer is set to 'UserProvided'.",
										Attributes: map[string]schema.Attribute{
											"ca": schema.StringAttribute{
												Description:         "Key of CA cert in Secret",
												MarkdownDescription: "Key of CA cert in Secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"cert": schema.StringAttribute{
												Description:         "Key of Cert in Secret",
												MarkdownDescription: "Key of Cert in Secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "Key of TLS private key in Secret",
												MarkdownDescription: "Key of TLS private key in Secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the Secret that contains user-provided certificates.",
												MarkdownDescription: "Name of the Secret that contains user-provided certificates.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Allows Pods to be scheduled onto nodes with matching taints.Each toleration in the array allows the Pod to tolerate node taints based onspecified 'key', 'value', 'effect', and 'operator'.- The 'key', 'value', and 'effect' identify the taint that the toleration matches.- The 'operator' determines how the toleration matches the taint.Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
						MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints.Each toleration in the array allows the Pod to tolerate node taints based onspecified 'key', 'value', 'effect', and 'operator'.- The 'key', 'value', and 'effect' identify the taint that the toleration matches.- The 'operator' determines how the toleration matches the taint.Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
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

					"volume_claim_templates": schema.ListNestedAttribute{
						Description:         "Specifies a list of PersistentVolumeClaim templates that define the storage requirements for the Component.Each template specifies the desired characteristics of a persistent volume, such as storage class,size, and access modes.These templates are used to dynamically provision persistent volumes for the Component.",
						MarkdownDescription: "Specifies a list of PersistentVolumeClaim templates that define the storage requirements for the Component.Each template specifies the desired characteristics of a persistent volume, such as storage class,size, and access modes.These templates are used to dynamically provision persistent volumes for the Component.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Refers to the name of a volumeMount defined in either:- 'componentDefinition.spec.runtime.containers[*].volumeMounts'- 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
									MarkdownDescription: "Refers to the name of a volumeMount defined in either:- 'componentDefinition.spec.runtime.containers[*].volumeMounts'- 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"spec": schema.SingleNestedAttribute{
									Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volumewith the mount name specified in the 'name' field.When a Pod is created for this ClusterComponent, a new PVC will be created based on the specificationdefined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
									MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volumewith the mount name specified in the 'name' field.When a Pod is created for this ClusterComponent, a new PVC will be created based on the specificationdefined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.MapAttribute{
											Description:         "Contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
											MarkdownDescription: "Contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "Represents the minimum resources the volume should have.If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements thatare lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
											MarkdownDescription: "Represents the minimum resources the volume should have.If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements thatare lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
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

										"storage_class_name": schema.StringAttribute{
											Description:         "The name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
											MarkdownDescription: "The name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
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

					"volumes": schema.ListNestedAttribute{
						Description:         "List of volumes to override.",
						MarkdownDescription: "List of volumes to override.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws_elastic_block_store": schema.SingleNestedAttribute{
									Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
									MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstoreTODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstoreTODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partition": schema.Int64Attribute{
											Description:         "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
											MarkdownDescription: "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly value true will force the readOnly setting in VolumeMounts.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_id": schema.StringAttribute{
											Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
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
											Description:         "fsType is Filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fsType is Filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
											MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
											Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
											MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",
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
											Description:         "monitors is Required: Monitors is a collection of Ceph monitorsMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitorsMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
											Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_file": schema.StringAttribute{
											Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secretMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secretMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "user is optional: User is the rados user name, default is adminMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "user is optional: User is the rados user name, default is adminMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
									Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
									MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef is optional: points to a secret object containing parameters used to connectto OpenStack.",
											MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connectto OpenStack.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "volumeID used to identify the volume in cinder.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "volumeID used to identify the volume in cinder.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
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
											Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"items": schema.ListNestedAttribute{
											Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
											MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
														Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
														MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "driver is the name of the CSI driver that handles this volume.Consult with your admin for the correct name as registered in the cluster.",
											MarkdownDescription: "driver is the name of the CSI driver that handles this volume.Consult with your admin for the correct name as registered in the cluster.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"fs_type": schema.StringAttribute{
											Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'.If not provided, the empty value is passed to the associated CSI driverwhich will determine the default filesystem to apply.",
											MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'.If not provided, the empty value is passed to the associated CSI driverwhich will determine the default filesystem to apply.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_publish_secret_ref": schema.SingleNestedAttribute{
											Description:         "nodePublishSecretRef is a reference to the secret object containingsensitive information to pass to the CSI driver to complete the CSINodePublishVolume and NodeUnpublishVolume calls.This field is optional, and  may be empty if no secret is required. If thesecret object contains more than one secret, all secret references are passed.",
											MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containingsensitive information to pass to the CSI driver to complete the CSINodePublishVolume and NodeUnpublishVolume calls.This field is optional, and  may be empty if no secret is required. If thesecret object contains more than one secret, all secret references are passed.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "readOnly specifies a read-only configuration for the volume.Defaults to false (read/write).",
											MarkdownDescription: "readOnly specifies a read-only configuration for the volume.Defaults to false (read/write).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_attributes": schema.MapAttribute{
											Description:         "volumeAttributes stores driver-specific properties that are passed to the CSIdriver. Consult your driver's documentation for supported values.",
											MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSIdriver. Consult your driver's documentation for supported values.",
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
											Description:         "Optional: mode bits to use on created files by default. Must be aOptional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "Optional: mode bits to use on created files by default. Must be aOptional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
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
														Description:         "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
														MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"resource_field_ref": schema.SingleNestedAttribute{
														Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
														MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
									Description:         "emptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
									MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
									Attributes: map[string]schema.Attribute{
										"medium": schema.StringAttribute{
											Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size_limit": schema.StringAttribute{
											Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
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
									Description:         "ephemeral represents a volume that is handled by a cluster storage driver.The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,and deleted when the pod is removed.Use this if:a) the volume is only needed while the pod runs,b) features of normal volumes like restoring from snapshot or capacity   tracking are needed,c) the storage driver is specified through a storage class, andd) the storage driver supports dynamic volume provisioning through   a PersistentVolumeClaim (see EphemeralVolumeSource for more   information on the connection between this volume type   and PersistentVolumeClaim).Use PersistentVolumeClaim or one of the vendor-specificAPIs for volumes that persist for longer than the lifecycleof an individual pod.Use CSI for light-weight local ephemeral volumes if the CSI driver is meant tobe used that way - see the documentation of the driver formore information.A pod can use both types of ephemeral volumes andpersistent volumes at the same time.",
									MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver.The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,and deleted when the pod is removed.Use this if:a) the volume is only needed while the pod runs,b) features of normal volumes like restoring from snapshot or capacity   tracking are needed,c) the storage driver is specified through a storage class, andd) the storage driver supports dynamic volume provisioning through   a PersistentVolumeClaim (see EphemeralVolumeSource for more   information on the connection between this volume type   and PersistentVolumeClaim).Use PersistentVolumeClaim or one of the vendor-specificAPIs for volumes that persist for longer than the lifecycleof an individual pod.Use CSI for light-weight local ephemeral volumes if the CSI driver is meant tobe used that way - see the documentation of the driver formore information.A pod can use both types of ephemeral volumes andpersistent volumes at the same time.",
									Attributes: map[string]schema.Attribute{
										"volume_claim_template": schema.SingleNestedAttribute{
											Description:         "Will be used to create a stand-alone PVC to provision the volume.The pod in which this EphemeralVolumeSource is embedded will be theowner of the PVC, i.e. the PVC will be deleted together with thepod.  The name of the PVC will be '<pod name>-<volume name>' where'<volume name>' is the name from the 'PodSpec.Volumes' arrayentry. Pod validation will reject the pod if the concatenated nameis not valid for a PVC (for example, too long).An existing PVC with that name that is not owned by the podwill *not* be used for the pod to avoid using an unrelatedvolume by mistake. Starting the pod is then blocked untilthe unrelated PVC is removed. If such a pre-created PVC ismeant to be used by the pod, the PVC has to updated with anowner reference to the pod once the pod exists. Normallythis should not be necessary, but it may be useful whenmanually reconstructing a broken cluster.This field is read-only and no changes will be made by Kubernetesto the PVC after it has been created.Required, must not be nil.",
											MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume.The pod in which this EphemeralVolumeSource is embedded will be theowner of the PVC, i.e. the PVC will be deleted together with thepod.  The name of the PVC will be '<pod name>-<volume name>' where'<volume name>' is the name from the 'PodSpec.Volumes' arrayentry. Pod validation will reject the pod if the concatenated nameis not valid for a PVC (for example, too long).An existing PVC with that name that is not owned by the podwill *not* be used for the pod to avoid using an unrelatedvolume by mistake. Starting the pod is then blocked untilthe unrelated PVC is removed. If such a pre-created PVC ismeant to be used by the pod, the PVC has to updated with anowner reference to the pod once the pod exists. Normallythis should not be necessary, but it may be useful whenmanually reconstructing a broken cluster.This field is read-only and no changes will be made by Kubernetesto the PVC after it has been created.Required, must not be nil.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.SingleNestedAttribute{
													Description:         "May contain labels and annotations that will be copied into the PVCwhen creating it. No other fields are allowed and will be rejected duringvalidation.",
													MarkdownDescription: "May contain labels and annotations that will be copied into the PVCwhen creating it. No other fields are allowed and will be rejected duringvalidation.",
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
													Description:         "The specification for the PersistentVolumeClaim. The entire content iscopied unchanged into the PVC that gets created from thistemplate. The same fields as in a PersistentVolumeClaimare also valid here.",
													MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content iscopied unchanged into the PVC that gets created from thistemplate. The same fields as in a PersistentVolumeClaimare also valid here.",
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
											Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.TODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.TODO: how do we prevent errors in the filesystem from compromising the machine",
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
											Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
											Description:         "wwids Optional: FC volume world wide identifiers (wwids)Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
											MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids)Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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
									Description:         "flexVolume represents a generic volume resource that isprovisioned/attached using an exec based plugin.",
									MarkdownDescription: "flexVolume represents a generic volume resource that isprovisioned/attached using an exec based plugin.",
									Attributes: map[string]schema.Attribute{
										"driver": schema.StringAttribute{
											Description:         "driver is the name of the driver to use for this volume.",
											MarkdownDescription: "driver is the name of the driver to use for this volume.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"fs_type": schema.StringAttribute{
											Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
											MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
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
											Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef is Optional: secretRef is reference to the secret object containingsensitive information to pass to the plugin scripts. This may beempty if no secret object is specified. If the secret objectcontains more than one secret, all secrets are passed to the pluginscripts.",
											MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containingsensitive information to pass to the plugin scripts. This may beempty if no secret object is specified. If the secret objectcontains more than one secret, all secrets are passed to the pluginscripts.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flockershould be considered as deprecated",
											MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flockershould be considered as deprecated",
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
									Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
									MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "fsType is filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdiskTODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "fsType is filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdiskTODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partition": schema.Int64Attribute{
											Description:         "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pd_name": schema.StringAttribute{
											Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
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
									Description:         "gitRepo represents a git repository at a particular revision.DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount anEmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDirinto the Pod's container.",
									MarkdownDescription: "gitRepo represents a git repository at a particular revision.DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount anEmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDirinto the Pod's container.",
									Attributes: map[string]schema.Attribute{
										"directory": schema.StringAttribute{
											Description:         "directory is the target directory name.Must not contain or start with '..'.  If '.' is supplied, the volume directory will be thegit repository.  Otherwise, if specified, the volume will contain the git repository inthe subdirectory with the given name.",
											MarkdownDescription: "directory is the target directory name.Must not contain or start with '..'.  If '.' is supplied, the volume directory will be thegit repository.  Otherwise, if specified, the volume will contain the git repository inthe subdirectory with the given name.",
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
									Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/glusterfs/README.md",
									MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/glusterfs/README.md",
									Attributes: map[string]schema.Attribute{
										"endpoints": schema.StringAttribute{
											Description:         "endpoints is the endpoint name that details Glusterfs topology.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "path is the Glusterfs volume path.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											MarkdownDescription: "path is the Glusterfs volume path.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions.Defaults to false.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions.Defaults to false.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
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
									Description:         "hostPath represents a pre-existing file or directory on the hostmachine that is directly exposed to the container. This is generallyused for system agents or other privileged things that are allowedto see the host machine. Most containers will NOT need this.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath---TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can notmount host directories as read/write.",
									MarkdownDescription: "hostPath represents a pre-existing file or directory on the hostmachine that is directly exposed to the container. This is generallyused for system agents or other privileged things that are allowedto see the host machine. Most containers will NOT need this.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath---TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can notmount host directories as read/write.",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
											Description:         "path of the directory on the host.If the path is a symlink, it will follow the link to the real path.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											MarkdownDescription: "path of the directory on the host.If the path is a symlink, it will follow the link to the real path.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "type for HostPath VolumeDefaults to ''More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											MarkdownDescription: "type for HostPath VolumeDefaults to ''More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
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
									Description:         "iscsi represents an ISCSI Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://examples.k8s.io/volumes/iscsi/README.md",
									MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://examples.k8s.io/volumes/iscsi/README.md",
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
											Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsiTODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsiTODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"initiator_name": schema.StringAttribute{
											Description:         "initiatorName is the custom iSCSI Initiator Name.If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<target portal>:<volume name> will be created for the connection.",
											MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name.If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<target portal>:<volume name> will be created for the connection.",
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
											Description:         "iscsiInterface is the interface Name that uses an iSCSI transport.Defaults to 'default' (tcp).",
											MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport.Defaults to 'default' (tcp).",
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
											Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
											MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.",
											MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
											MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
											MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
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
									Description:         "name of the volume.Must be a DNS_LABEL and unique within the pod.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									MarkdownDescription: "name of the volume.Must be a DNS_LABEL and unique within the pod.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"nfs": schema.SingleNestedAttribute{
									Description:         "nfs represents an NFS mount on the host that shares a pod's lifetimeMore info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
									MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetimeMore info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
											Description:         "path that is exported by the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "path that is exported by the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly here will force the NFS export to be mounted with read-only permissions.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"server": schema.StringAttribute{
											Description:         "server is the hostname or IP address of the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "server is the hostname or IP address of the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
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
									Description:         "persistentVolumeClaimVolumeSource represents a reference to aPersistentVolumeClaim in the same namespace.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to aPersistentVolumeClaim in the same namespace.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"claim_name": schema.StringAttribute{
											Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
											MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
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
											Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
											Description:         "fSType represents the filesystem type to mountMust be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fSType represents the filesystem type to mountMust be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
											Description:         "defaultMode are the mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
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
														Description:         "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' fieldof ClusterTrustBundle objects in an auto-updating file.Alpha, gated by the ClusterTrustBundleProjection feature gate.ClusterTrustBundle objects can either be selected by name, or by thecombination of signer name and a label selector.Kubelet performs aggressive normalization of the PEM contents writteninto the pod filesystem.  Esoteric PEM features such as inter-blockcomments and block headers are stripped.  Certificates are deduplicated.The ordering of certificates within the file is arbitrary, and Kubeletmay change the order over time.",
														MarkdownDescription: "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' fieldof ClusterTrustBundle objects in an auto-updating file.Alpha, gated by the ClusterTrustBundleProjection feature gate.ClusterTrustBundle objects can either be selected by name, or by thecombination of signer name and a label selector.Kubelet performs aggressive normalization of the PEM contents writteninto the pod filesystem.  Esoteric PEM features such as inter-blockcomments and block headers are stripped.  Certificates are deduplicated.The ordering of certificates within the file is arbitrary, and Kubeletmay change the order over time.",
														Attributes: map[string]schema.Attribute{
															"label_selector": schema.SingleNestedAttribute{
																Description:         "Select all ClusterTrustBundles that match this label selector.  Only haseffect if signerName is set.  Mutually-exclusive with name.  If unset,interpreted as 'match nothing'.  If set but empty, interpreted as 'matcheverything'.",
																MarkdownDescription: "Select all ClusterTrustBundles that match this label selector.  Only haseffect if signerName is set.  Mutually-exclusive with name.  If unset,interpreted as 'match nothing'.  If set but empty, interpreted as 'matcheverything'.",
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

															"name": schema.StringAttribute{
																Description:         "Select a single ClusterTrustBundle by object name.  Mutually-exclusivewith signerName and labelSelector.",
																MarkdownDescription: "Select a single ClusterTrustBundle by object name.  Mutually-exclusivewith signerName and labelSelector.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "If true, don't block pod startup if the referenced ClusterTrustBundle(s)aren't available.  If using name, then the named ClusterTrustBundle isallowed not to exist.  If using signerName, then the combination ofsignerName and labelSelector is allowed to match zeroClusterTrustBundles.",
																MarkdownDescription: "If true, don't block pod startup if the referenced ClusterTrustBundle(s)aren't available.  If using name, then the named ClusterTrustBundle isallowed not to exist.  If using signerName, then the combination ofsignerName and labelSelector is allowed to match zeroClusterTrustBundles.",
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
																Description:         "Select all ClusterTrustBundles that match this signer name.Mutually-exclusive with name.  The contents of all selectedClusterTrustBundles will be unified and deduplicated.",
																MarkdownDescription: "Select all ClusterTrustBundles that match this signer name.Mutually-exclusive with name.  The contents of all selectedClusterTrustBundles will be unified and deduplicated.",
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
																Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																			Description:         "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																			MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"resource_field_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																			MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
																Description:         "items if unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																Description:         "audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
																MarkdownDescription: "audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"expiration_seconds": schema.Int64Attribute{
																Description:         "expirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours.Defaults to 1 hourand must be at least 10 minutes.",
																MarkdownDescription: "expirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours.Defaults to 1 hourand must be at least 10 minutes.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "path is the path relative to the mount point of the file to project thetoken into.",
																MarkdownDescription: "path is the path relative to the mount point of the file to project thetoken into.",
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
											Description:         "group to map volume access toDefault is no group",
											MarkdownDescription: "group to map volume access toDefault is no group",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions.Defaults to false.",
											MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions.Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registry": schema.StringAttribute{
											Description:         "registry represents a single or multiple Quobyte Registry servicesspecified as a string as host:port pair (multiple entries are separated with commas)which acts as the central registry for volumes",
											MarkdownDescription: "registry represents a single or multiple Quobyte Registry servicesspecified as a string as host:port pair (multiple entries are separated with commas)which acts as the central registry for volumes",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"tenant": schema.StringAttribute{
											Description:         "tenant owning the given Quobyte volume in the BackendUsed with dynamically provisioned Quobyte volumes, value is set by the plugin",
											MarkdownDescription: "tenant owning the given Quobyte volume in the BackendUsed with dynamically provisioned Quobyte volumes, value is set by the plugin",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user": schema.StringAttribute{
											Description:         "user to map volume access toDefaults to serivceaccount user",
											MarkdownDescription: "user to map volume access toDefaults to serivceaccount user",
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
									Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/rbd/README.md",
									MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/rbd/README.md",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#rbdTODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#rbdTODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "image is the rados image name.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "image is the rados image name.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"keyring": schema.StringAttribute{
											Description:         "keyring is the path to key ring for RBDUser.Default is /etc/ceph/keyring.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "keyring is the path to key ring for RBDUser.Default is /etc/ceph/keyring.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"monitors": schema.ListAttribute{
											Description:         "monitors is a collection of Ceph monitors.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "monitors is a collection of Ceph monitors.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pool": schema.StringAttribute{
											Description:         "pool is the rados pool name.Default is rbd.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "pool is the rados pool name.Default is rbd.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef is name of the authentication secret for RBDUser. If providedoverrides keyring.Default is nil.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If providedoverrides keyring.Default is nil.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "user is the rados user name.Default is admin.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "user is the rados user name.Default is admin.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
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
											Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'.Default is 'xfs'.",
											MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'.Default is 'xfs'.",
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
											Description:         "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef references to the secret for ScaleIO user and othersensitive information. If this is not provided, Login operation will fail.",
											MarkdownDescription: "secretRef references to the secret for ScaleIO user and othersensitive information. If this is not provided, Login operation will fail.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.Default is ThinProvisioned.",
											MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.Default is ThinProvisioned.",
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
											Description:         "volumeName is the name of a volume already created in the ScaleIO systemthat is associated with this volume source.",
											MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO systemthat is associated with this volume source.",
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
									Description:         "secret represents a secret that should populate this volume.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
									MarkdownDescription: "secret represents a secret that should populate this volume.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"items": schema.ListNestedAttribute{
											Description:         "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
											MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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
														Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
														MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
											Description:         "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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
											Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "secretRef specifies the secret to use for obtaining the StorageOS APIcredentials.  If not specified, default values will be attempted.",
											MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS APIcredentials.  If not specified, default values will be attempted.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Description:         "volumeName is the human-readable name of the StorageOS volume.  Volumenames are only unique within a namespace.",
											MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volumenames are only unique within a namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_namespace": schema.StringAttribute{
											Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If nonamespace is specified then the Pod's namespace will be used.  This allows theKubernetes name scoping to be mirrored within StorageOS for tighter integration.Set VolumeName to any name to override the default behaviour.Set to 'default' if you are not using namespaces within StorageOS.Namespaces that do not pre-exist within StorageOS will be created.",
											MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If nonamespace is specified then the Pod's namespace will be used.  This allows theKubernetes name scoping to be mirrored within StorageOS for tighter integration.Set VolumeName to any name to override the default behaviour.Set to 'default' if you are not using namespaces within StorageOS.Namespaces that do not pre-exist within StorageOS will be created.",
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
											Description:         "fsType is filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fsType is filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoComponentV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_component_v1alpha1_manifest")

	var model AppsKubeblocksIoComponentV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Component")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
