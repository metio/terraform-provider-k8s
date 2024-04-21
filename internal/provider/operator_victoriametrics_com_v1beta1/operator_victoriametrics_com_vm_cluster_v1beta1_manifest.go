/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

import (
	"context"
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
	_ datasource.DataSource = &OperatorVictoriametricsComVmclusterV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmclusterV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmclusterV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmclusterV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmclusterV1Beta1ManifestData struct {
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
		ClusterVersion   *string `tfsdk:"cluster_version" json:"clusterVersion,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		License *struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			KeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_ref" json:"keyRef,omitempty"`
		} `tfsdk:"license" json:"license,omitempty"`
		ReplicationFactor  *int64  `tfsdk:"replication_factor" json:"replicationFactor,omitempty"`
		RetentionPeriod    *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		UseStrictSecurity  *bool   `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
		Vminsert           *struct {
			Affinity                *map[string]string   `tfsdk:"affinity" json:"affinity,omitempty"`
			ClusterNativeListenPort *string              `tfsdk:"cluster_native_listen_port" json:"clusterNativeListenPort,omitempty"`
			ConfigMaps              *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
			Containers              *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
			DnsConfig               *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy   *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			ExtraArgs   *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs   *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			HostNetwork *bool                `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Hpa         *map[string]string   `tfsdk:"hpa" json:"hpa,omitempty"`
			Image       *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			InsertPorts    *struct {
				GraphitePort     *string `tfsdk:"graphite_port" json:"graphitePort,omitempty"`
				InfluxPort       *string `tfsdk:"influx_port" json:"influxPort,omitempty"`
				OpenTSDBHTTPPort *string `tfsdk:"open_tsdbhttp_port" json:"openTSDBHTTPPort,omitempty"`
				OpenTSDBPort     *string `tfsdk:"open_tsdb_port" json:"openTSDBPort,omitempty"`
			} `tfsdk:"insert_ports" json:"insertPorts,omitempty"`
			LivenessProbe       *map[string]string `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogFormat           *string            `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogLevel            *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
			MinReadySeconds     *int64             `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			NodeSelector        *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodDisruptionBudget *struct {
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			PodMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Port              *string `tfsdk:"port" json:"port,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RevisionHistoryLimitCount *int64 `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
			RollingUpdate             *struct {
				MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
			RuntimeClassName  *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName     *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Secrets           *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext   *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
			ServiceScrapeSpec *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
			ServiceSpec       *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			StartupProbe                  *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
			TerminationGracePeriodSeconds *int64             `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			Tolerations                   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
			UpdateStrategy            *string              `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
			VolumeMounts              *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"vminsert" json:"vminsert,omitempty"`
		Vmselect *struct {
			Affinity       *map[string]string `tfsdk:"affinity" json:"affinity,omitempty"`
			CacheMountPath *string            `tfsdk:"cache_mount_path" json:"cacheMountPath,omitempty"`
			ClaimTemplates *[]struct {
				ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
				Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec       *struct {
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
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
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
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
				Status *struct {
					AccessModes        *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
					AllocatedResources *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
					Capacity           *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
					Conditions         *[]struct {
						LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
						LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
						Message            *string `tfsdk:"message" json:"message,omitempty"`
						Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
						Status             *string `tfsdk:"status" json:"status,omitempty"`
						Type               *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"conditions" json:"conditions,omitempty"`
					Phase        *string `tfsdk:"phase" json:"phase,omitempty"`
					ResizeStatus *string `tfsdk:"resize_status" json:"resizeStatus,omitempty"`
				} `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"claim_templates" json:"claimTemplates,omitempty"`
			ClusterNativeListenPort *string              `tfsdk:"cluster_native_listen_port" json:"clusterNativeListenPort,omitempty"`
			ConfigMaps              *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
			Containers              *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
			DnsConfig               *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy   *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			ExtraArgs   *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs   *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			HostNetwork *bool                `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Hpa         *map[string]string   `tfsdk:"hpa" json:"hpa,omitempty"`
			Image       *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			InitContainers   *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			LivenessProbe    *map[string]string   `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogFormat        *string              `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogLevel         *string              `tfsdk:"log_level" json:"logLevel,omitempty"`
			MinReadySeconds  *int64               `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			NodeSelector     *map[string]string   `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PersistentVolume *struct {
				DisableMountSubPath *bool `tfsdk:"disable_mount_sub_path" json:"disableMountSubPath,omitempty"`
				EmptyDir            *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				VolumeClaimTemplate *map[string]string `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"persistent_volume" json:"persistentVolume,omitempty"`
			PodDisruptionBudget *struct {
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			PodMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Port              *string `tfsdk:"port" json:"port,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RevisionHistoryLimitCount *int64             `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
			RollingUpdateStrategy     *string            `tfsdk:"rolling_update_strategy" json:"rollingUpdateStrategy,omitempty"`
			RuntimeClassName          *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName             *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Secrets                   *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext           *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
			ServiceScrapeSpec         *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
			ServiceSpec               *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			StartupProbe *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
			Storage      *struct {
				DisableMountSubPath *bool `tfsdk:"disable_mount_sub_path" json:"disableMountSubPath,omitempty"`
				EmptyDir            *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				VolumeClaimTemplate *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Metadata   *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name        *string            `tfsdk:"name" json:"name,omitempty"`
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
							Claims *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"claims" json:"claims,omitempty"`
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
						StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
						VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
						VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
					Status *struct {
						AccessModes        *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
						AllocatedResources *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
						Capacity           *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
						Conditions         *[]struct {
							LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
							LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
							Message            *string `tfsdk:"message" json:"message,omitempty"`
							Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
							Status             *string `tfsdk:"status" json:"status,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"conditions" json:"conditions,omitempty"`
						Phase        *string `tfsdk:"phase" json:"phase,omitempty"`
						ResizeStatus *string `tfsdk:"resize_status" json:"resizeStatus,omitempty"`
					} `tfsdk:"status" json:"status,omitempty"`
				} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			Tolerations                   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
			VolumeMounts              *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"vmselect" json:"vmselect,omitempty"`
		Vmstorage *struct {
			Affinity       *map[string]string `tfsdk:"affinity" json:"affinity,omitempty"`
			ClaimTemplates *[]struct {
				ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
				Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec       *struct {
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
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
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
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
				Status *struct {
					AccessModes        *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
					AllocatedResources *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
					Capacity           *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
					Conditions         *[]struct {
						LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
						LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
						Message            *string `tfsdk:"message" json:"message,omitempty"`
						Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
						Status             *string `tfsdk:"status" json:"status,omitempty"`
						Type               *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"conditions" json:"conditions,omitempty"`
					Phase        *string `tfsdk:"phase" json:"phase,omitempty"`
					ResizeStatus *string `tfsdk:"resize_status" json:"resizeStatus,omitempty"`
				} `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"claim_templates" json:"claimTemplates,omitempty"`
			ConfigMaps *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
			Containers *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
			DnsConfig  *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy   *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			ExtraArgs   *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs   *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			HostNetwork *bool                `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Image       *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			InitContainers           *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			LivenessProbe            *map[string]string   `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogFormat                *string              `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogLevel                 *string              `tfsdk:"log_level" json:"logLevel,omitempty"`
			MaintenanceInsertNodeIDs *[]string            `tfsdk:"maintenance_insert_node_i_ds" json:"maintenanceInsertNodeIDs,omitempty"`
			MaintenanceSelectNodeIDs *[]string            `tfsdk:"maintenance_select_node_i_ds" json:"maintenanceSelectNodeIDs,omitempty"`
			MinReadySeconds          *int64               `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			NodeSelector             *map[string]string   `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodDisruptionBudget      *struct {
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			PodMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Port              *string `tfsdk:"port" json:"port,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RevisionHistoryLimitCount *int64             `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
			RollingUpdateStrategy     *string            `tfsdk:"rolling_update_strategy" json:"rollingUpdateStrategy,omitempty"`
			RuntimeClassName          *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName             *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Secrets                   *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext           *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
			ServiceScrapeSpec         *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
			ServiceSpec               *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			StartupProbe *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
			Storage      *struct {
				DisableMountSubPath *bool `tfsdk:"disable_mount_sub_path" json:"disableMountSubPath,omitempty"`
				EmptyDir            *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				VolumeClaimTemplate *map[string]string `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			StorageDataPath               *string `tfsdk:"storage_data_path" json:"storageDataPath,omitempty"`
			TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			Tolerations                   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
			VmBackup                  *struct {
				AcceptEULA        *bool  `tfsdk:"accept_eula" json:"acceptEULA,omitempty"`
				Concurrency       *int64 `tfsdk:"concurrency" json:"concurrency,omitempty"`
				CredentialsSecret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"credentials_secret" json:"credentialsSecret,omitempty"`
				CustomS3Endpoint            *string            `tfsdk:"custom_s3_endpoint" json:"customS3Endpoint,omitempty"`
				Destination                 *string            `tfsdk:"destination" json:"destination,omitempty"`
				DestinationDisableSuffixAdd *bool              `tfsdk:"destination_disable_suffix_add" json:"destinationDisableSuffixAdd,omitempty"`
				DisableDaily                *bool              `tfsdk:"disable_daily" json:"disableDaily,omitempty"`
				DisableHourly               *bool              `tfsdk:"disable_hourly" json:"disableHourly,omitempty"`
				DisableMonthly              *bool              `tfsdk:"disable_monthly" json:"disableMonthly,omitempty"`
				DisableWeekly               *bool              `tfsdk:"disable_weekly" json:"disableWeekly,omitempty"`
				ExtraArgs                   *map[string]string `tfsdk:"extra_args" json:"extraArgs,omitempty"`
				ExtraEnvs                   *[]struct {
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
				} `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
				Image *struct {
					PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				LogFormat *string `tfsdk:"log_format" json:"logFormat,omitempty"`
				LogLevel  *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				Port      *string `tfsdk:"port" json:"port,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Restore *struct {
					OnStart *struct {
						Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					} `tfsdk:"on_start" json:"onStart,omitempty"`
				} `tfsdk:"restore" json:"restore,omitempty"`
				SnapshotCreateURL *string `tfsdk:"snapshot_create_url" json:"snapshotCreateURL,omitempty"`
				SnapshotDeleteURL *string `tfsdk:"snapshot_delete_url" json:"snapshotDeleteURL,omitempty"`
				VolumeMounts      *[]struct {
					MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
					Name             *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"vm_backup" json:"vmBackup,omitempty"`
			VmInsertPort *string `tfsdk:"vm_insert_port" json:"vmInsertPort,omitempty"`
			VmSelectPort *string `tfsdk:"vm_select_port" json:"vmSelectPort,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"vmstorage" json:"vmstorage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmclusterV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_cluster_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmclusterV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMCluster is fast, cost-effective and scalable time-series database.Cluster version with",
		MarkdownDescription: "VMCluster is fast, cost-effective and scalable time-series database.Cluster version with",
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
				Description:         "VMClusterSpec defines the desired state of VMCluster",
				MarkdownDescription: "VMClusterSpec defines the desired state of VMCluster",
				Attributes: map[string]schema.Attribute{
					"cluster_version": schema.StringAttribute{
						Description:         "ClusterVersion defines default images tag for all components.it can be overwritten with component specific image.tag value.",
						MarkdownDescription: "ClusterVersion defines default images tag for all components.it can be overwritten with component specific image.tag value.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets An optional list of references to secrets in the same namespaceto use for pulling images from registriessee https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespaceto use for pulling images from registriessee https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

					"license": schema.SingleNestedAttribute{
						Description:         "License allows to configure license key to be used for enterprise features.Using license key is supported starting from VictoriaMetrics v1.94.0.See: https://docs.victoriametrics.com/enterprise.html",
						MarkdownDescription: "License allows to configure license key to be used for enterprise features.Using license key is supported starting from VictoriaMetrics v1.94.0.See: https://docs.victoriametrics.com/enterprise.html",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Enterprise license key. This flag is available only in VictoriaMetrics enterprise.Documentation - https://docs.victoriametrics.com/enterprise.htmlfor more information, visit https://victoriametrics.com/products/enterprise/ .To request a trial license, go to https://victoriametrics.com/products/enterprise/trial/",
								MarkdownDescription: "Enterprise license key. This flag is available only in VictoriaMetrics enterprise.Documentation - https://docs.victoriametrics.com/enterprise.htmlfor more information, visit https://victoriametrics.com/products/enterprise/ .To request a trial license, go to https://victoriametrics.com/products/enterprise/trial/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_ref": schema.SingleNestedAttribute{
								Description:         "KeyRef is reference to secret with license key for enterprise features.",
								MarkdownDescription: "KeyRef is reference to secret with license key for enterprise features.",
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

					"replication_factor": schema.Int64Attribute{
						Description:         "ReplicationFactor defines how many copies of data make amongdistinct storage nodes",
						MarkdownDescription: "ReplicationFactor defines how many copies of data make amongdistinct storage nodes",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retention_period": schema.StringAttribute{
						Description:         "RetentionPeriod for the stored metricsNote VictoriaMetrics has data/ and indexdb/ foldersmetrics from data/ removed eventually as soon as partition leaves retention periodreverse index data at indexdb rotates once at the half of configured retention periodhttps://docs.victoriametrics.com/Single-server-VictoriaMetrics.html#retention",
						MarkdownDescription: "RetentionPeriod for the stored metricsNote VictoriaMetrics has data/ and indexdb/ foldersmetrics from data/ removed eventually as soon as partition leaves retention periodreverse index data at indexdb rotates once at the half of configured retention periodhttps://docs.victoriametrics.com/Single-server-VictoriaMetrics.html#retention",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run theVMSelect, VMStorage and VMInsert Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run theVMSelect, VMStorage and VMInsert Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_strict_security": schema.BoolAttribute{
						Description:         "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						MarkdownDescription: "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vminsert": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.MapAttribute{
								Description:         "Affinity If specified, the pod's scheduling constraints.",
								MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_native_listen_port": schema.StringAttribute{
								Description:         "ClusterNativePort for multi-level cluster setup.More details: https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#multi-level-cluster-setup",
								MarkdownDescription: "ClusterNativePort for multi-level cluster setup.More details: https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#multi-level-cluster-setup",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the VMInsertobject, which shall be mounted into the VMInsert Pods.The ConfigMaps are mounted into /etc/vm/configs/<configmap-name>.",
								MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the VMInsertobject, which shall be mounted into the VMInsert Pods.The ConfigMaps are mounted into /etc/vm/configs/<configmap-name>.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
								MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
								MarkdownDescription: "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Required.",
													MarkdownDescription: "Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"searches": schema.ListAttribute{
										Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
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

							"dns_policy": schema.StringAttribute{
								Description:         "DNSPolicy sets DNS policy for the pod",
								MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListAttribute{
								Description:         "ExtraEnvs that will be added to VMInsert pod",
								MarkdownDescription: "ExtraEnvs that will be added to VMInsert pod",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork controls whether the pod may use the node network namespace",
								MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hpa": schema.MapAttribute{
								Description:         "HPA defines kubernetes PodAutoScaling configuration version 2.",
								MarkdownDescription: "HPA defines kubernetes PodAutoScaling configuration version 2.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings for VMInsert",
								MarkdownDescription: "Image - docker image settings for VMInsert",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"init_containers": schema.ListAttribute{
								Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the VMInsert configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
								MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the VMInsert configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insert_ports": schema.SingleNestedAttribute{
								Description:         "InsertPorts - additional listen ports for data ingestion.",
								MarkdownDescription: "InsertPorts - additional listen ports for data ingestion.",
								Attributes: map[string]schema.Attribute{
									"graphite_port": schema.StringAttribute{
										Description:         "GraphitePort listen port",
										MarkdownDescription: "GraphitePort listen port",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"influx_port": schema.StringAttribute{
										Description:         "InfluxPort listen port",
										MarkdownDescription: "InfluxPort listen port",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"open_tsdbhttp_port": schema.StringAttribute{
										Description:         "OpenTSDBHTTPPort for http connections.",
										MarkdownDescription: "OpenTSDBHTTPPort for http connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"open_tsdb_port": schema.StringAttribute{
										Description:         "OpenTSDBPort for tcp and udp listen",
										MarkdownDescription: "OpenTSDBPort for tcp and udp listen",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"liveness_probe": schema.MapAttribute{
								Description:         "LivenessProbe that will be added CRD pod",
								MarkdownDescription: "LivenessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VMInsert to be configured with.default or json",
								MarkdownDescription: "LogFormat for VMInsert to be configured with.default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VMInsert to be configured with.",
								MarkdownDescription: "LogLevel for VMInsert to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
								MarkdownDescription: "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
								MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget created by operator",
								MarkdownDescription: "PodDisruptionBudget created by operator",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
										MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector_labels": schema.MapAttribute{
										Description:         "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
										MarkdownDescription: "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
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

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata configures Labels and Annotations which are propagated to the VMInsert pods.",
								MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VMInsert pods.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.StringAttribute{
								Description:         "Port listen port",
								MarkdownDescription: "Port listen port",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "Priority class assigned to the Pods",
								MarkdownDescription: "Priority class assigned to the Pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "ReadinessGates defines pod readiness gates",
								MarkdownDescription: "ReadinessGates defines pod readiness gates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
											Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
											MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

							"readiness_probe": schema.MapAttribute{
								Description:         "ReadinessProbe that will be added CRD pod",
								MarkdownDescription: "ReadinessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_count": schema.Int64Attribute{
								Description:         "ReplicaCount is the expected size of the VMInsert cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.",
								MarkdownDescription: "ReplicaCount is the expected size of the VMInsert cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"revision_history_limit_count": schema.Int64Attribute{
								Description:         "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
								MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update": schema.SingleNestedAttribute{
								Description:         "RollingUpdate - overrides deployment update params.",
								MarkdownDescription: "RollingUpdate - overrides deployment update params.",
								Attributes: map[string]schema.Attribute{
									"max_surge": schema.StringAttribute{
										Description:         "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_unavailable": schema.StringAttribute{
										Description:         "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
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
								Description:         "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
								MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "SchedulerName - defines kubernetes scheduler name",
								MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets is a list of Secrets in the same namespace as the VMInsertobject, which shall be mounted into the VMInsert Pods.The Secrets are mounted into /etc/vm/secrets/<secret-name>.",
								MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the VMInsertobject, which shall be mounted into the VMInsert Pods.The Secrets are mounted into /etc/vm/secrets/<secret-name>.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_scrape_spec": schema.MapAttribute{
								Description:         "ServiceScrapeSpec that will be added to vminsert VMServiceScrape spec",
								MarkdownDescription: "ServiceScrapeSpec that will be added to vminsert VMServiceScrape spec",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec that will be added to vminsert service spec",
								MarkdownDescription: "ServiceSpec that will be added to vminsert service spec",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
										MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.MapAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"use_as_default": schema.BoolAttribute{
										Description:         "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
										MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_probe": schema.MapAttribute{
								Description:         "StartupProbe that will be added to CRD pod",
								MarkdownDescription: "StartupProbe that will be added to CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "TerminationGracePeriodSeconds period for container graceful termination",
								MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations If specified, the pod's tolerations.",
								MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
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

							"topology_spread_constraints": schema.ListAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_strategy": schema.StringAttribute{
								Description:         "UpdateStrategy - overrides default update strategy.",
								MarkdownDescription: "UpdateStrategy - overrides default update strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Recreate", "RollingUpdate"),
								},
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMInsert container,that are generated as a result of StorageSpec objects.",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMInsert container,that are generated as a result of StorageSpec objects.",
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

							"volumes": schema.ListAttribute{
								Description:         "Volumes allows configuration of additional volumes on the output Deployment definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
								MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vmselect": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.MapAttribute{
								Description:         "Affinity If specified, the pod's scheduling constraints.",
								MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_mount_path": schema.StringAttribute{
								Description:         "CacheMountPath allows to add cache persistent for VMSelect,will use '/cache' as default if not specified.",
								MarkdownDescription: "CacheMountPath allows to add cache persistent for VMSelect,will use '/cache' as default if not specified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"claim_templates": schema.ListNestedAttribute{
								Description:         "ClaimTemplates allows adding additional VolumeClaimTemplates for StatefulSet",
								MarkdownDescription: "ClaimTemplates allows adding additional VolumeClaimTemplates for StatefulSet",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
											MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metadata": schema.MapAttribute{
											Description:         "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
											MarkdownDescription: "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"status": schema.SingleNestedAttribute{
											Description:         "status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											Attributes: map[string]schema.Attribute{
												"access_modes": schema.ListAttribute{
													Description:         "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"allocated_resources": schema.MapAttribute{
													Description:         "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
													MarkdownDescription: "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"capacity": schema.MapAttribute{
													Description:         "capacity represents the actual resources of the underlying volume.",
													MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conditions": schema.ListNestedAttribute{
													Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
													MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"last_probe_time": schema.StringAttribute{
																Description:         "lastProbeTime is the time we probed the condition.",
																MarkdownDescription: "lastProbeTime is the time we probed the condition.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	validators.DateTime64Validator(),
																},
															},

															"last_transition_time": schema.StringAttribute{
																Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
																MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	validators.DateTime64Validator(),
																},
															},

															"message": schema.StringAttribute{
																Description:         "message is the human-readable message indicating details about last transition.",
																MarkdownDescription: "message is the human-readable message indicating details about last transition.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"reason": schema.StringAttribute{
																Description:         "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
																MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"status": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
																MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

												"phase": schema.StringAttribute{
													Description:         "phase represents the current phase of PersistentVolumeClaim.",
													MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resize_status": schema.StringAttribute{
													Description:         "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
													MarkdownDescription: "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
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

							"cluster_native_listen_port": schema.StringAttribute{
								Description:         "ClusterNativePort for multi-level cluster setup.More details: https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#multi-level-cluster-setup",
								MarkdownDescription: "ClusterNativePort for multi-level cluster setup.More details: https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#multi-level-cluster-setup",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the VMSelectobject, which shall be mounted into the VMSelect Pods.The ConfigMaps are mounted into /etc/vm/configs/<configmap-name>.",
								MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the VMSelectobject, which shall be mounted into the VMSelect Pods.The ConfigMaps are mounted into /etc/vm/configs/<configmap-name>.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
								MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
								MarkdownDescription: "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Required.",
													MarkdownDescription: "Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"searches": schema.ListAttribute{
										Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
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

							"dns_policy": schema.StringAttribute{
								Description:         "DNSPolicy sets DNS policy for the pod",
								MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListAttribute{
								Description:         "ExtraEnvs that will be added to VMSelect pod",
								MarkdownDescription: "ExtraEnvs that will be added to VMSelect pod",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork controls whether the pod may use the node network namespace",
								MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hpa": schema.MapAttribute{
								Description:         "Configures horizontal pod autoscaling.Note, enabling this option disables vmselect to vmselect communication. In most cases it's not an issue.",
								MarkdownDescription: "Configures horizontal pod autoscaling.Note, enabling this option disables vmselect to vmselect communication. In most cases it's not an issue.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings for VMSelect",
								MarkdownDescription: "Image - docker image settings for VMSelect",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"init_containers": schema.ListAttribute{
								Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the VMSelect configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
								MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the VMSelect configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.MapAttribute{
								Description:         "LivenessProbe that will be added CRD pod",
								MarkdownDescription: "LivenessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VMSelect to be configured with.default or json",
								MarkdownDescription: "LogFormat for VMSelect to be configured with.default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VMSelect to be configured with.",
								MarkdownDescription: "LogLevel for VMSelect to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
								MarkdownDescription: "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
								MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"persistent_volume": schema.SingleNestedAttribute{
								Description:         "Storage - add persistent volume for cacheMounthPathits useful for persistent cacheuse storage instead of persistentVolume.",
								MarkdownDescription: "Storage - add persistent volume for cacheMounthPathits useful for persistent cacheuse storage instead of persistentVolume.",
								Attributes: map[string]schema.Attribute{
									"disable_mount_sub_path": schema.BoolAttribute{
										Description:         "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										MarkdownDescription: "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"empty_dir": schema.SingleNestedAttribute{
										Description:         "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
										MarkdownDescription: "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
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

									"volume_claim_template": schema.MapAttribute{
										Description:         "A PVC spec to be used by the VMAlertManager StatefulSets.",
										MarkdownDescription: "A PVC spec to be used by the VMAlertManager StatefulSets.",
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

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget created by operator",
								MarkdownDescription: "PodDisruptionBudget created by operator",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
										MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector_labels": schema.MapAttribute{
										Description:         "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
										MarkdownDescription: "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
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

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata configures Labels and Annotations which are propagated to the VMSelect pods.",
								MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VMSelect pods.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.StringAttribute{
								Description:         "Port listen port",
								MarkdownDescription: "Port listen port",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "Priority class assigned to the Pods",
								MarkdownDescription: "Priority class assigned to the Pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "ReadinessGates defines pod readiness gates",
								MarkdownDescription: "ReadinessGates defines pod readiness gates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
											Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
											MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

							"readiness_probe": schema.MapAttribute{
								Description:         "ReadinessProbe that will be added CRD pod",
								MarkdownDescription: "ReadinessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_count": schema.Int64Attribute{
								Description:         "ReplicaCount is the expected size of the VMSelect cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.",
								MarkdownDescription: "ReplicaCount is the expected size of the VMSelect cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"revision_history_limit_count": schema.Int64Attribute{
								Description:         "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
								MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update_strategy": schema.StringAttribute{
								Description:         "RollingUpdateStrategy defines strategy for application updatesDefault is OnDelete, in this case operator handles update processCan be changed for RollingUpdate",
								MarkdownDescription: "RollingUpdateStrategy defines strategy for application updatesDefault is OnDelete, in this case operator handles update processCan be changed for RollingUpdate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
								MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "SchedulerName - defines kubernetes scheduler name",
								MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets is a list of Secrets in the same namespace as the VMSelectobject, which shall be mounted into the VMSelect Pods.The Secrets are mounted into /etc/vm/secrets/<secret-name>.",
								MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the VMSelectobject, which shall be mounted into the VMSelect Pods.The Secrets are mounted into /etc/vm/secrets/<secret-name>.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_scrape_spec": schema.MapAttribute{
								Description:         "ServiceScrapeSpec that will be added to vmselect VMServiceScrape spec",
								MarkdownDescription: "ServiceScrapeSpec that will be added to vmselect VMServiceScrape spec",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec that will be added to vmselect service spec",
								MarkdownDescription: "ServiceSpec that will be added to vmselect service spec",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
										MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.MapAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"use_as_default": schema.BoolAttribute{
										Description:         "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
										MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_probe": schema.MapAttribute{
								Description:         "StartupProbe that will be added to CRD pod",
								MarkdownDescription: "StartupProbe that will be added to CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage": schema.SingleNestedAttribute{
								Description:         "StorageSpec - add persistent volume claim for cacheMountPathits needed for persistent cache",
								MarkdownDescription: "StorageSpec - add persistent volume claim for cacheMountPathits needed for persistent cache",
								Attributes: map[string]schema.Attribute{
									"disable_mount_sub_path": schema.BoolAttribute{
										Description:         "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										MarkdownDescription: "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"empty_dir": schema.SingleNestedAttribute{
										Description:         "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
										MarkdownDescription: "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
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

									"volume_claim_template": schema.SingleNestedAttribute{
										Description:         "A PVC spec to be used by the VMAlertManager StatefulSets.",
										MarkdownDescription: "A PVC spec to be used by the VMAlertManager StatefulSets.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
												MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "EmbeddedMetadata contains metadata relevant to an EmbeddedResource.",
												MarkdownDescription: "EmbeddedMetadata contains metadata relevant to an EmbeddedResource.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
														MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
														MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
														MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
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
												Description:         "Spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "Spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"status": schema.SingleNestedAttribute{
												Description:         "Status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "Status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												Attributes: map[string]schema.Attribute{
													"access_modes": schema.ListAttribute{
														Description:         "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
														MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"allocated_resources": schema.MapAttribute{
														Description:         "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
														MarkdownDescription: "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"capacity": schema.MapAttribute{
														Description:         "capacity represents the actual resources of the underlying volume.",
														MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"conditions": schema.ListNestedAttribute{
														Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
														MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"last_probe_time": schema.StringAttribute{
																	Description:         "lastProbeTime is the time we probed the condition.",
																	MarkdownDescription: "lastProbeTime is the time we probed the condition.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		validators.DateTime64Validator(),
																	},
																},

																"last_transition_time": schema.StringAttribute{
																	Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
																	MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		validators.DateTime64Validator(),
																	},
																},

																"message": schema.StringAttribute{
																	Description:         "message is the human-readable message indicating details about last transition.",
																	MarkdownDescription: "message is the human-readable message indicating details about last transition.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"reason": schema.StringAttribute{
																	Description:         "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
																	MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"status": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
																	MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

													"phase": schema.StringAttribute{
														Description:         "phase represents the current phase of PersistentVolumeClaim.",
														MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resize_status": schema.StringAttribute{
														Description:         "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
														MarkdownDescription: "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "TerminationGracePeriodSeconds period for container graceful termination",
								MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations If specified, the pod's tolerations.",
								MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
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

							"topology_spread_constraints": schema.ListAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMSelect container,that are generated as a result of StorageSpec objects.",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMSelect container,that are generated as a result of StorageSpec objects.",
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

							"volumes": schema.ListAttribute{
								Description:         "Volumes allows configuration of additional volumes on the output Deployment definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
								MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vmstorage": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.MapAttribute{
								Description:         "Affinity If specified, the pod's scheduling constraints.",
								MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"claim_templates": schema.ListNestedAttribute{
								Description:         "ClaimTemplates allows adding additional VolumeClaimTemplates for StatefulSet",
								MarkdownDescription: "ClaimTemplates allows adding additional VolumeClaimTemplates for StatefulSet",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
											MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metadata": schema.MapAttribute{
											Description:         "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
											MarkdownDescription: "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"status": schema.SingleNestedAttribute{
											Description:         "status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											Attributes: map[string]schema.Attribute{
												"access_modes": schema.ListAttribute{
													Description:         "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"allocated_resources": schema.MapAttribute{
													Description:         "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
													MarkdownDescription: "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"capacity": schema.MapAttribute{
													Description:         "capacity represents the actual resources of the underlying volume.",
													MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"conditions": schema.ListNestedAttribute{
													Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
													MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"last_probe_time": schema.StringAttribute{
																Description:         "lastProbeTime is the time we probed the condition.",
																MarkdownDescription: "lastProbeTime is the time we probed the condition.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	validators.DateTime64Validator(),
																},
															},

															"last_transition_time": schema.StringAttribute{
																Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
																MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	validators.DateTime64Validator(),
																},
															},

															"message": schema.StringAttribute{
																Description:         "message is the human-readable message indicating details about last transition.",
																MarkdownDescription: "message is the human-readable message indicating details about last transition.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"reason": schema.StringAttribute{
																Description:         "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
																MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"status": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
																MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

												"phase": schema.StringAttribute{
													Description:         "phase represents the current phase of PersistentVolumeClaim.",
													MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resize_status": schema.StringAttribute{
													Description:         "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
													MarkdownDescription: "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
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

							"config_maps": schema.ListAttribute{
								Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the VMStorageobject, which shall be mounted into the VMStorage Pods.The ConfigMaps are mounted into /etc/vm/configs/<configmap-name>.",
								MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the VMStorageobject, which shall be mounted into the VMStorage Pods.The ConfigMaps are mounted into /etc/vm/configs/<configmap-name>.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
								MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
								MarkdownDescription: "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Required.",
													MarkdownDescription: "Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"searches": schema.ListAttribute{
										Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
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

							"dns_policy": schema.StringAttribute{
								Description:         "DNSPolicy sets DNS policy for the pod",
								MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListAttribute{
								Description:         "ExtraEnvs that will be added to VMStorage pod",
								MarkdownDescription: "ExtraEnvs that will be added to VMStorage pod",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork controls whether the pod may use the node network namespace",
								MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings for VMStorage",
								MarkdownDescription: "Image - docker image settings for VMStorage",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"init_containers": schema.ListAttribute{
								Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the VMStorage configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
								MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the VMStorage configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.MapAttribute{
								Description:         "LivenessProbe that will be added CRD pod",
								MarkdownDescription: "LivenessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VMStorage to be configured with.default or json",
								MarkdownDescription: "LogFormat for VMStorage to be configured with.default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VMStorage to be configured with.",
								MarkdownDescription: "LogLevel for VMStorage to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"maintenance_insert_node_i_ds": schema.ListAttribute{
								Description:         "MaintenanceInsertNodeIDs - excludes given node ids from insert requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc.lets say, you have pod-0, pod-1, pod-2, pod-3. to exclude pod-0 and pod-3 from insert routing, define nodeIDs: [0,3].Useful at storage expanding, when you want to rebalance some data at cluster.",
								MarkdownDescription: "MaintenanceInsertNodeIDs - excludes given node ids from insert requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc.lets say, you have pod-0, pod-1, pod-2, pod-3. to exclude pod-0 and pod-3 from insert routing, define nodeIDs: [0,3].Useful at storage expanding, when you want to rebalance some data at cluster.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maintenance_select_node_i_ds": schema.ListAttribute{
								Description:         "MaintenanceInsertNodeIDs - excludes given node ids from select requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc.",
								MarkdownDescription: "MaintenanceInsertNodeIDs - excludes given node ids from select requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
								MarkdownDescription: "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
								MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget created by operator",
								MarkdownDescription: "PodDisruptionBudget created by operator",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
										MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector_labels": schema.MapAttribute{
										Description:         "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
										MarkdownDescription: "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
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

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata configures Labels and Annotations which are propagated to the VMStorage pods.",
								MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VMStorage pods.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.StringAttribute{
								Description:         "Port for health check connetions",
								MarkdownDescription: "Port for health check connetions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "Priority class assigned to the Pods",
								MarkdownDescription: "Priority class assigned to the Pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "ReadinessGates defines pod readiness gates",
								MarkdownDescription: "ReadinessGates defines pod readiness gates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
											Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
											MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

							"readiness_probe": schema.MapAttribute{
								Description:         "ReadinessProbe that will be added CRD pod",
								MarkdownDescription: "ReadinessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_count": schema.Int64Attribute{
								Description:         "ReplicaCount is the expected size of the VMStorage cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.",
								MarkdownDescription: "ReplicaCount is the expected size of the VMStorage cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"revision_history_limit_count": schema.Int64Attribute{
								Description:         "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
								MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update_strategy": schema.StringAttribute{
								Description:         "RollingUpdateStrategy defines strategy for application updatesDefault is OnDelete, in this case operator handles update processCan be changed for RollingUpdate",
								MarkdownDescription: "RollingUpdateStrategy defines strategy for application updatesDefault is OnDelete, in this case operator handles update processCan be changed for RollingUpdate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
								MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "SchedulerName - defines kubernetes scheduler name",
								MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets is a list of Secrets in the same namespace as the VMStorageobject, which shall be mounted into the VMStorage Pods.The Secrets are mounted into /etc/vm/secrets/<secret-name>.",
								MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the VMStorageobject, which shall be mounted into the VMStorage Pods.The Secrets are mounted into /etc/vm/secrets/<secret-name>.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_scrape_spec": schema.MapAttribute{
								Description:         "ServiceScrapeSpec that will be added to vmstorage VMServiceScrape spec",
								MarkdownDescription: "ServiceScrapeSpec that will be added to vmstorage VMServiceScrape spec",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec that will be create additional service for vmstorage",
								MarkdownDescription: "ServiceSpec that will be create additional service for vmstorage",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
										MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.MapAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"use_as_default": schema.BoolAttribute{
										Description:         "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
										MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_probe": schema.MapAttribute{
								Description:         "StartupProbe that will be added to CRD pod",
								MarkdownDescription: "StartupProbe that will be added to CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage": schema.SingleNestedAttribute{
								Description:         "Storage - add persistent volume for StorageDataPathits useful for persistent cache",
								MarkdownDescription: "Storage - add persistent volume for StorageDataPathits useful for persistent cache",
								Attributes: map[string]schema.Attribute{
									"disable_mount_sub_path": schema.BoolAttribute{
										Description:         "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										MarkdownDescription: "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"empty_dir": schema.SingleNestedAttribute{
										Description:         "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
										MarkdownDescription: "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
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

									"volume_claim_template": schema.MapAttribute{
										Description:         "A PVC spec to be used by the VMAlertManager StatefulSets.",
										MarkdownDescription: "A PVC spec to be used by the VMAlertManager StatefulSets.",
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

							"storage_data_path": schema.StringAttribute{
								Description:         "StorageDataPath - path to storage data",
								MarkdownDescription: "StorageDataPath - path to storage data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "TerminationGracePeriodSeconds period for container graceful termination",
								MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations If specified, the pod's tolerations.",
								MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
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

							"topology_spread_constraints": schema.ListAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vm_backup": schema.SingleNestedAttribute{
								Description:         "VMBackup configuration for backup",
								MarkdownDescription: "VMBackup configuration for backup",
								Attributes: map[string]schema.Attribute{
									"accept_eula": schema.BoolAttribute{
										Description:         "AcceptEULA accepts enterprise feature usage, must be set to true.otherwise backupmanager cannot be added to single/cluster version.https://victoriametrics.com/legal/esa/",
										MarkdownDescription: "AcceptEULA accepts enterprise feature usage, must be set to true.otherwise backupmanager cannot be added to single/cluster version.https://victoriametrics.com/legal/esa/",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"concurrency": schema.Int64Attribute{
										Description:         "Defines number of concurrent workers. Higher concurrency may reduce backup duration (default 10)",
										MarkdownDescription: "Defines number of concurrent workers. Higher concurrency may reduce backup duration (default 10)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials_secret": schema.SingleNestedAttribute{
										Description:         "CredentialsSecret is secret in the same namespace for access to remote storageThe secret is mounted into /etc/vm/creds.",
										MarkdownDescription: "CredentialsSecret is secret in the same namespace for access to remote storageThe secret is mounted into /etc/vm/creds.",
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

									"custom_s3_endpoint": schema.StringAttribute{
										Description:         "Custom S3 endpoint for use with S3-compatible storages (e.g. MinIO). S3 is used if not set",
										MarkdownDescription: "Custom S3 endpoint for use with S3-compatible storages (e.g. MinIO). S3 is used if not set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"destination": schema.StringAttribute{
										Description:         "Defines destination for backup",
										MarkdownDescription: "Defines destination for backup",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"destination_disable_suffix_add": schema.BoolAttribute{
										Description:         "DestinationDisableSuffixAdd - disables suffix adding for cluster version backupseach vmstorage backup must have unique backup folderso operator adds POD_NAME as suffix for backup destination folder.",
										MarkdownDescription: "DestinationDisableSuffixAdd - disables suffix adding for cluster version backupseach vmstorage backup must have unique backup folderso operator adds POD_NAME as suffix for backup destination folder.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_daily": schema.BoolAttribute{
										Description:         "Defines if daily backups disabled (default false)",
										MarkdownDescription: "Defines if daily backups disabled (default false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_hourly": schema.BoolAttribute{
										Description:         "Defines if hourly backups disabled (default false)",
										MarkdownDescription: "Defines if hourly backups disabled (default false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_monthly": schema.BoolAttribute{
										Description:         "Defines if monthly backups disabled (default false)",
										MarkdownDescription: "Defines if monthly backups disabled (default false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_weekly": schema.BoolAttribute{
										Description:         "Defines if weekly backups disabled (default false)",
										MarkdownDescription: "Defines if weekly backups disabled (default false)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"extra_args": schema.MapAttribute{
										Description:         "extra args like maxBytesPerSecond default 0",
										MarkdownDescription: "extra args like maxBytesPerSecond default 0",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"extra_envs": schema.ListNestedAttribute{
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

									"image": schema.SingleNestedAttribute{
										Description:         "Image - docker image settings for VMBackuper",
										MarkdownDescription: "Image - docker image settings for VMBackuper",
										Attributes: map[string]schema.Attribute{
											"pull_policy": schema.StringAttribute{
												Description:         "PullPolicy describes how to pull docker image",
												MarkdownDescription: "PullPolicy describes how to pull docker image",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"repository": schema.StringAttribute{
												Description:         "Repository contains name of docker image + it's repository if needed",
												MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tag": schema.StringAttribute{
												Description:         "Tag contains desired docker image version",
												MarkdownDescription: "Tag contains desired docker image version",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_format": schema.StringAttribute{
										Description:         "LogFormat for VMBackup to be configured with.default or json",
										MarkdownDescription: "LogFormat for VMBackup to be configured with.default or json",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("default", "json"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "LogLevel for VMBackup to be configured with.",
										MarkdownDescription: "LogLevel for VMBackup to be configured with.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
										},
									},

									"port": schema.StringAttribute{
										Description:         "Port for health check connections",
										MarkdownDescription: "Port for health check connections",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
										MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
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

									"restore": schema.SingleNestedAttribute{
										Description:         "Restore Allows to enable restore options for podRead more: https://docs.victoriametrics.com/vmbackupmanager.html#restore-commands",
										MarkdownDescription: "Restore Allows to enable restore options for podRead more: https://docs.victoriametrics.com/vmbackupmanager.html#restore-commands",
										Attributes: map[string]schema.Attribute{
											"on_start": schema.SingleNestedAttribute{
												Description:         "OnStart defines configuration for restore on pod start",
												MarkdownDescription: "OnStart defines configuration for restore on pod start",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "Enabled defines if restore on start enabled",
														MarkdownDescription: "Enabled defines if restore on start enabled",
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

									"snapshot_create_url": schema.StringAttribute{
										Description:         "SnapshotCreateURL overwrites url for snapshot create",
										MarkdownDescription: "SnapshotCreateURL overwrites url for snapshot create",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"snapshot_delete_url": schema.StringAttribute{
										Description:         "SnapShotDeleteURL overwrites url for snapshot delete",
										MarkdownDescription: "SnapShotDeleteURL overwrites url for snapshot delete",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_mounts": schema.ListNestedAttribute{
										Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the vmbackupmanager container,that are generated as a result of StorageSpec objects.",
										MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the vmbackupmanager container,that are generated as a result of StorageSpec objects.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"vm_insert_port": schema.StringAttribute{
								Description:         "VMInsertPort for VMInsert connections",
								MarkdownDescription: "VMInsertPort for VMInsert connections",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vm_select_port": schema.StringAttribute{
								Description:         "VMSelectPort for VMSelect connections",
								MarkdownDescription: "VMSelectPort for VMSelect connections",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMStorage container,that are generated as a result of StorageSpec objects.",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMStorage container,that are generated as a result of StorageSpec objects.",
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

							"volumes": schema.ListAttribute{
								Description:         "Volumes allows configuration of additional volumes on the output Deployment definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
								MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
								ElementType:         types.MapType{ElemType: types.StringType},
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

func (r *OperatorVictoriametricsComVmclusterV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_cluster_v1beta1_manifest")

	var model OperatorVictoriametricsComVmclusterV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
