/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

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
	_ datasource.DataSource = &CephRookIoCephFilesystemV1Manifest{}
)

func NewCephRookIoCephFilesystemV1Manifest() datasource.DataSource {
	return &CephRookIoCephFilesystemV1Manifest{}
}

type CephRookIoCephFilesystemV1Manifest struct{}

type CephRookIoCephFilesystemV1ManifestData struct {
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
		DataPools *[]struct {
			Application        *string `tfsdk:"application" json:"application,omitempty"`
			CompressionMode    *string `tfsdk:"compression_mode" json:"compressionMode,omitempty"`
			CrushRoot          *string `tfsdk:"crush_root" json:"crushRoot,omitempty"`
			DeviceClass        *string `tfsdk:"device_class" json:"deviceClass,omitempty"`
			EnableCrushUpdates *bool   `tfsdk:"enable_crush_updates" json:"enableCrushUpdates,omitempty"`
			EnableRBDStats     *bool   `tfsdk:"enable_rbd_stats" json:"enableRBDStats,omitempty"`
			ErasureCoded       *struct {
				Algorithm    *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
				CodingChunks *int64  `tfsdk:"coding_chunks" json:"codingChunks,omitempty"`
				DataChunks   *int64  `tfsdk:"data_chunks" json:"dataChunks,omitempty"`
			} `tfsdk:"erasure_coded" json:"erasureCoded,omitempty"`
			FailureDomain *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
			Mirroring     *struct {
				Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
				Peers   *struct {
					SecretNames *[]string `tfsdk:"secret_names" json:"secretNames,omitempty"`
				} `tfsdk:"peers" json:"peers,omitempty"`
				SnapshotSchedules *[]struct {
					Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					StartTime *string `tfsdk:"start_time" json:"startTime,omitempty"`
				} `tfsdk:"snapshot_schedules" json:"snapshotSchedules,omitempty"`
			} `tfsdk:"mirroring" json:"mirroring,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			Quotas     *struct {
				MaxBytes   *int64  `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
				MaxObjects *int64  `tfsdk:"max_objects" json:"maxObjects,omitempty"`
				MaxSize    *string `tfsdk:"max_size" json:"maxSize,omitempty"`
			} `tfsdk:"quotas" json:"quotas,omitempty"`
			Replicated *struct {
				HybridStorage *struct {
					PrimaryDeviceClass   *string `tfsdk:"primary_device_class" json:"primaryDeviceClass,omitempty"`
					SecondaryDeviceClass *string `tfsdk:"secondary_device_class" json:"secondaryDeviceClass,omitempty"`
				} `tfsdk:"hybrid_storage" json:"hybridStorage,omitempty"`
				ReplicasPerFailureDomain *int64   `tfsdk:"replicas_per_failure_domain" json:"replicasPerFailureDomain,omitempty"`
				RequireSafeReplicaSize   *bool    `tfsdk:"require_safe_replica_size" json:"requireSafeReplicaSize,omitempty"`
				Size                     *int64   `tfsdk:"size" json:"size,omitempty"`
				SubFailureDomain         *string  `tfsdk:"sub_failure_domain" json:"subFailureDomain,omitempty"`
				TargetSizeRatio          *float64 `tfsdk:"target_size_ratio" json:"targetSizeRatio,omitempty"`
			} `tfsdk:"replicated" json:"replicated,omitempty"`
			StatusCheck *struct {
				Mirror *struct {
					Disabled *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					Interval *string `tfsdk:"interval" json:"interval,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"mirror" json:"mirror,omitempty"`
			} `tfsdk:"status_check" json:"statusCheck,omitempty"`
		} `tfsdk:"data_pools" json:"dataPools,omitempty"`
		MetadataPool *struct {
			Application        *string `tfsdk:"application" json:"application,omitempty"`
			CompressionMode    *string `tfsdk:"compression_mode" json:"compressionMode,omitempty"`
			CrushRoot          *string `tfsdk:"crush_root" json:"crushRoot,omitempty"`
			DeviceClass        *string `tfsdk:"device_class" json:"deviceClass,omitempty"`
			EnableCrushUpdates *bool   `tfsdk:"enable_crush_updates" json:"enableCrushUpdates,omitempty"`
			EnableRBDStats     *bool   `tfsdk:"enable_rbd_stats" json:"enableRBDStats,omitempty"`
			ErasureCoded       *struct {
				Algorithm    *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
				CodingChunks *int64  `tfsdk:"coding_chunks" json:"codingChunks,omitempty"`
				DataChunks   *int64  `tfsdk:"data_chunks" json:"dataChunks,omitempty"`
			} `tfsdk:"erasure_coded" json:"erasureCoded,omitempty"`
			FailureDomain *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
			Mirroring     *struct {
				Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
				Peers   *struct {
					SecretNames *[]string `tfsdk:"secret_names" json:"secretNames,omitempty"`
				} `tfsdk:"peers" json:"peers,omitempty"`
				SnapshotSchedules *[]struct {
					Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					StartTime *string `tfsdk:"start_time" json:"startTime,omitempty"`
				} `tfsdk:"snapshot_schedules" json:"snapshotSchedules,omitempty"`
			} `tfsdk:"mirroring" json:"mirroring,omitempty"`
			Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			Quotas     *struct {
				MaxBytes   *int64  `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
				MaxObjects *int64  `tfsdk:"max_objects" json:"maxObjects,omitempty"`
				MaxSize    *string `tfsdk:"max_size" json:"maxSize,omitempty"`
			} `tfsdk:"quotas" json:"quotas,omitempty"`
			Replicated *struct {
				HybridStorage *struct {
					PrimaryDeviceClass   *string `tfsdk:"primary_device_class" json:"primaryDeviceClass,omitempty"`
					SecondaryDeviceClass *string `tfsdk:"secondary_device_class" json:"secondaryDeviceClass,omitempty"`
				} `tfsdk:"hybrid_storage" json:"hybridStorage,omitempty"`
				ReplicasPerFailureDomain *int64   `tfsdk:"replicas_per_failure_domain" json:"replicasPerFailureDomain,omitempty"`
				RequireSafeReplicaSize   *bool    `tfsdk:"require_safe_replica_size" json:"requireSafeReplicaSize,omitempty"`
				Size                     *int64   `tfsdk:"size" json:"size,omitempty"`
				SubFailureDomain         *string  `tfsdk:"sub_failure_domain" json:"subFailureDomain,omitempty"`
				TargetSizeRatio          *float64 `tfsdk:"target_size_ratio" json:"targetSizeRatio,omitempty"`
			} `tfsdk:"replicated" json:"replicated,omitempty"`
			StatusCheck *struct {
				Mirror *struct {
					Disabled *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					Interval *string `tfsdk:"interval" json:"interval,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"mirror" json:"mirror,omitempty"`
			} `tfsdk:"status_check" json:"statusCheck,omitempty"`
		} `tfsdk:"metadata_pool" json:"metadataPool,omitempty"`
		MetadataServer *struct {
			ActiveCount   *int64             `tfsdk:"active_count" json:"activeCount,omitempty"`
			ActiveStandby *bool              `tfsdk:"active_standby" json:"activeStandby,omitempty"`
			Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			LivenessProbe *struct {
				Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				Probe    *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					Grpc             *struct {
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					HttpGet *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						HttpHeaders *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TcpSocket           *struct {
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *string `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"probe" json:"probe,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
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
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			StartupProbe *struct {
				Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				Probe    *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					Grpc             *struct {
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					HttpGet *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						HttpHeaders *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TcpSocket           *struct {
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *string `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"probe" json:"probe,omitempty"`
			} `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
		} `tfsdk:"metadata_server" json:"metadataServer,omitempty"`
		Mirroring *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Peers   *struct {
				SecretNames *[]string `tfsdk:"secret_names" json:"secretNames,omitempty"`
			} `tfsdk:"peers" json:"peers,omitempty"`
			SnapshotRetention *[]struct {
				Duration *string `tfsdk:"duration" json:"duration,omitempty"`
				Path     *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"snapshot_retention" json:"snapshotRetention,omitempty"`
			SnapshotSchedules *[]struct {
				Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				StartTime *string `tfsdk:"start_time" json:"startTime,omitempty"`
			} `tfsdk:"snapshot_schedules" json:"snapshotSchedules,omitempty"`
		} `tfsdk:"mirroring" json:"mirroring,omitempty"`
		PreserveFilesystemOnDelete *bool `tfsdk:"preserve_filesystem_on_delete" json:"preserveFilesystemOnDelete,omitempty"`
		PreservePoolsOnDelete      *bool `tfsdk:"preserve_pools_on_delete" json:"preservePoolsOnDelete,omitempty"`
		StatusCheck                *struct {
			Mirror *struct {
				Disabled *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"mirror" json:"mirror,omitempty"`
		} `tfsdk:"status_check" json:"statusCheck,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephFilesystemV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_filesystem_v1_manifest"
}

func (r *CephRookIoCephFilesystemV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephFilesystem represents a Ceph Filesystem",
		MarkdownDescription: "CephFilesystem represents a Ceph Filesystem",
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
				Description:         "FilesystemSpec represents the spec of a file system",
				MarkdownDescription: "FilesystemSpec represents the spec of a file system",
				Attributes: map[string]schema.Attribute{
					"data_pools": schema.ListNestedAttribute{
						Description:         "The data pool settings, with optional predefined pool name.",
						MarkdownDescription: "The data pool settings, with optional predefined pool name.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"application": schema.StringAttribute{
									Description:         "The application name to set on the pool. Only expected to be set for rgw pools.",
									MarkdownDescription: "The application name to set on the pool. Only expected to be set for rgw pools.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"compression_mode": schema.StringAttribute{
									Description:         "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force'The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)Do NOT set a default value for kubebuilder as this will override the Parameters",
									MarkdownDescription: "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force'The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)Do NOT set a default value for kubebuilder as this will override the Parameters",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("none", "passive", "aggressive", "force", ""),
									},
								},

								"crush_root": schema.StringAttribute{
									Description:         "The root of the crush hierarchy utilized by the pool",
									MarkdownDescription: "The root of the crush hierarchy utilized by the pool",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"device_class": schema.StringAttribute{
									Description:         "The device class the OSD should set to for use in the pool",
									MarkdownDescription: "The device class the OSD should set to for use in the pool",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enable_crush_updates": schema.BoolAttribute{
									Description:         "Allow rook operator to change the pool CRUSH tunables once the pool is created",
									MarkdownDescription: "Allow rook operator to change the pool CRUSH tunables once the pool is created",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enable_rbd_stats": schema.BoolAttribute{
									Description:         "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
									MarkdownDescription: "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"erasure_coded": schema.SingleNestedAttribute{
									Description:         "The erasure code settings",
									MarkdownDescription: "The erasure code settings",
									Attributes: map[string]schema.Attribute{
										"algorithm": schema.StringAttribute{
											Description:         "The algorithm for erasure coding",
											MarkdownDescription: "The algorithm for erasure coding",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"coding_chunks": schema.Int64Attribute{
											Description:         "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type).This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
											MarkdownDescription: "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type).This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"data_chunks": schema.Int64Attribute{
											Description:         "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type).The number of chunks required to recover an object when any single OSD is lost is the sameas dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
											MarkdownDescription: "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type).The number of chunks required to recover an object when any single OSD is lost is the sameas dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
											Required:            true,
											Optional:            false,
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

								"failure_domain": schema.StringAttribute{
									Description:         "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
									MarkdownDescription: "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"mirroring": schema.SingleNestedAttribute{
									Description:         "The mirroring settings",
									MarkdownDescription: "The mirroring settings",
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description:         "Enabled whether this pool is mirrored or not",
											MarkdownDescription: "Enabled whether this pool is mirrored or not",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is the mirroring mode: either pool or image",
											MarkdownDescription: "Mode is the mirroring mode: either pool or image",
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

										"snapshot_schedules": schema.ListNestedAttribute{
											Description:         "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
											MarkdownDescription: "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"interval": schema.StringAttribute{
														Description:         "Interval represent the periodicity of the snapshot.",
														MarkdownDescription: "Interval represent the periodicity of the snapshot.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "Path is the path to snapshot, only valid for CephFS",
														MarkdownDescription: "Path is the path to snapshot, only valid for CephFS",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"start_time": schema.StringAttribute{
														Description:         "StartTime indicates when to start the snapshot",
														MarkdownDescription: "StartTime indicates when to start the snapshot",
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

								"name": schema.StringAttribute{
									Description:         "Name of the pool",
									MarkdownDescription: "Name of the pool",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"parameters": schema.MapAttribute{
									Description:         "Parameters is a list of properties to enable on a given pool",
									MarkdownDescription: "Parameters is a list of properties to enable on a given pool",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"quotas": schema.SingleNestedAttribute{
									Description:         "The quota settings",
									MarkdownDescription: "The quota settings",
									Attributes: map[string]schema.Attribute{
										"max_bytes": schema.Int64Attribute{
											Description:         "MaxBytes represents the quota in bytesDeprecated in favor of MaxSize",
											MarkdownDescription: "MaxBytes represents the quota in bytesDeprecated in favor of MaxSize",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_objects": schema.Int64Attribute{
											Description:         "MaxObjects represents the quota in objects",
											MarkdownDescription: "MaxObjects represents the quota in objects",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_size": schema.StringAttribute{
											Description:         "MaxSize represents the quota in bytes as a string",
											MarkdownDescription: "MaxSize represents the quota in bytes as a string",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[\.]?[0-9]*([KMGTPE]i|[kMGTPE])?$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"replicated": schema.SingleNestedAttribute{
									Description:         "The replication settings",
									MarkdownDescription: "The replication settings",
									Attributes: map[string]schema.Attribute{
										"hybrid_storage": schema.SingleNestedAttribute{
											Description:         "HybridStorage represents hybrid storage tier settings",
											MarkdownDescription: "HybridStorage represents hybrid storage tier settings",
											Attributes: map[string]schema.Attribute{
												"primary_device_class": schema.StringAttribute{
													Description:         "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
													MarkdownDescription: "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},

												"secondary_device_class": schema.StringAttribute{
													Description:         "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
													MarkdownDescription: "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"replicas_per_failure_domain": schema.Int64Attribute{
											Description:         "ReplicasPerFailureDomain the number of replica in the specified failure domain",
											MarkdownDescription: "ReplicasPerFailureDomain the number of replica in the specified failure domain",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"require_safe_replica_size": schema.BoolAttribute{
											Description:         "RequireSafeReplicaSize if false allows you to set replica 1",
											MarkdownDescription: "RequireSafeReplicaSize if false allows you to set replica 1",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size": schema.Int64Attribute{
											Description:         "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
											MarkdownDescription: "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"sub_failure_domain": schema.StringAttribute{
											Description:         "SubFailureDomain the name of the sub-failure domain",
											MarkdownDescription: "SubFailureDomain the name of the sub-failure domain",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_size_ratio": schema.Float64Attribute{
											Description:         "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
											MarkdownDescription: "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"status_check": schema.SingleNestedAttribute{
									Description:         "The mirroring statusCheck",
									MarkdownDescription: "The mirroring statusCheck",
									Attributes: map[string]schema.Attribute{
										"mirror": schema.SingleNestedAttribute{
											Description:         "HealthCheckSpec represents the health check of an object store bucket",
											MarkdownDescription: "HealthCheckSpec represents the health check of an object store bucket",
											Attributes: map[string]schema.Attribute{
												"disabled": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
													MarkdownDescription: "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
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

					"metadata_pool": schema.SingleNestedAttribute{
						Description:         "The metadata pool settings",
						MarkdownDescription: "The metadata pool settings",
						Attributes: map[string]schema.Attribute{
							"application": schema.StringAttribute{
								Description:         "The application name to set on the pool. Only expected to be set for rgw pools.",
								MarkdownDescription: "The application name to set on the pool. Only expected to be set for rgw pools.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compression_mode": schema.StringAttribute{
								Description:         "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force'The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)Do NOT set a default value for kubebuilder as this will override the Parameters",
								MarkdownDescription: "DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force'The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force)Do NOT set a default value for kubebuilder as this will override the Parameters",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("none", "passive", "aggressive", "force", ""),
								},
							},

							"crush_root": schema.StringAttribute{
								Description:         "The root of the crush hierarchy utilized by the pool",
								MarkdownDescription: "The root of the crush hierarchy utilized by the pool",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device_class": schema.StringAttribute{
								Description:         "The device class the OSD should set to for use in the pool",
								MarkdownDescription: "The device class the OSD should set to for use in the pool",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_crush_updates": schema.BoolAttribute{
								Description:         "Allow rook operator to change the pool CRUSH tunables once the pool is created",
								MarkdownDescription: "Allow rook operator to change the pool CRUSH tunables once the pool is created",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_rbd_stats": schema.BoolAttribute{
								Description:         "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
								MarkdownDescription: "EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"erasure_coded": schema.SingleNestedAttribute{
								Description:         "The erasure code settings",
								MarkdownDescription: "The erasure code settings",
								Attributes: map[string]schema.Attribute{
									"algorithm": schema.StringAttribute{
										Description:         "The algorithm for erasure coding",
										MarkdownDescription: "The algorithm for erasure coding",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"coding_chunks": schema.Int64Attribute{
										Description:         "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type).This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
										MarkdownDescription: "Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type).This is the number of OSDs that can be lost simultaneously before data cannot be recovered.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"data_chunks": schema.Int64Attribute{
										Description:         "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type).The number of chunks required to recover an object when any single OSD is lost is the sameas dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
										MarkdownDescription: "Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type).The number of chunks required to recover an object when any single OSD is lost is the sameas dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.",
										Required:            true,
										Optional:            false,
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

							"failure_domain": schema.StringAttribute{
								Description:         "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
								MarkdownDescription: "The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mirroring": schema.SingleNestedAttribute{
								Description:         "The mirroring settings",
								MarkdownDescription: "The mirroring settings",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled whether this pool is mirrored or not",
										MarkdownDescription: "Enabled whether this pool is mirrored or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mode": schema.StringAttribute{
										Description:         "Mode is the mirroring mode: either pool or image",
										MarkdownDescription: "Mode is the mirroring mode: either pool or image",
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

									"snapshot_schedules": schema.ListNestedAttribute{
										Description:         "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
										MarkdownDescription: "SnapshotSchedules is the scheduling of snapshot for mirrored images/pools",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"interval": schema.StringAttribute{
													Description:         "Interval represent the periodicity of the snapshot.",
													MarkdownDescription: "Interval represent the periodicity of the snapshot.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path is the path to snapshot, only valid for CephFS",
													MarkdownDescription: "Path is the path to snapshot, only valid for CephFS",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"start_time": schema.StringAttribute{
													Description:         "StartTime indicates when to start the snapshot",
													MarkdownDescription: "StartTime indicates when to start the snapshot",
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

							"parameters": schema.MapAttribute{
								Description:         "Parameters is a list of properties to enable on a given pool",
								MarkdownDescription: "Parameters is a list of properties to enable on a given pool",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"quotas": schema.SingleNestedAttribute{
								Description:         "The quota settings",
								MarkdownDescription: "The quota settings",
								Attributes: map[string]schema.Attribute{
									"max_bytes": schema.Int64Attribute{
										Description:         "MaxBytes represents the quota in bytesDeprecated in favor of MaxSize",
										MarkdownDescription: "MaxBytes represents the quota in bytesDeprecated in favor of MaxSize",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_objects": schema.Int64Attribute{
										Description:         "MaxObjects represents the quota in objects",
										MarkdownDescription: "MaxObjects represents the quota in objects",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_size": schema.StringAttribute{
										Description:         "MaxSize represents the quota in bytes as a string",
										MarkdownDescription: "MaxSize represents the quota in bytes as a string",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[\.]?[0-9]*([KMGTPE]i|[kMGTPE])?$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicated": schema.SingleNestedAttribute{
								Description:         "The replication settings",
								MarkdownDescription: "The replication settings",
								Attributes: map[string]schema.Attribute{
									"hybrid_storage": schema.SingleNestedAttribute{
										Description:         "HybridStorage represents hybrid storage tier settings",
										MarkdownDescription: "HybridStorage represents hybrid storage tier settings",
										Attributes: map[string]schema.Attribute{
											"primary_device_class": schema.StringAttribute{
												Description:         "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
												MarkdownDescription: "PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"secondary_device_class": schema.StringAttribute{
												Description:         "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
												MarkdownDescription: "SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas_per_failure_domain": schema.Int64Attribute{
										Description:         "ReplicasPerFailureDomain the number of replica in the specified failure domain",
										MarkdownDescription: "ReplicasPerFailureDomain the number of replica in the specified failure domain",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"require_safe_replica_size": schema.BoolAttribute{
										Description:         "RequireSafeReplicaSize if false allows you to set replica 1",
										MarkdownDescription: "RequireSafeReplicaSize if false allows you to set replica 1",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"size": schema.Int64Attribute{
										Description:         "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
										MarkdownDescription: "Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"sub_failure_domain": schema.StringAttribute{
										Description:         "SubFailureDomain the name of the sub-failure domain",
										MarkdownDescription: "SubFailureDomain the name of the sub-failure domain",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_size_ratio": schema.Float64Attribute{
										Description:         "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
										MarkdownDescription: "TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"status_check": schema.SingleNestedAttribute{
								Description:         "The mirroring statusCheck",
								MarkdownDescription: "The mirroring statusCheck",
								Attributes: map[string]schema.Attribute{
									"mirror": schema.SingleNestedAttribute{
										Description:         "HealthCheckSpec represents the health check of an object store bucket",
										MarkdownDescription: "HealthCheckSpec represents the health check of an object store bucket",
										Attributes: map[string]schema.Attribute{
											"disabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"interval": schema.StringAttribute{
												Description:         "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
												MarkdownDescription: "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout": schema.StringAttribute{
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

					"metadata_server": schema.SingleNestedAttribute{
						Description:         "The mds pod info",
						MarkdownDescription: "The mds pod info",
						Attributes: map[string]schema.Attribute{
							"active_count": schema.Int64Attribute{
								Description:         "The number of metadata servers that are active. The remaining servers in the cluster will be in standby mode.",
								MarkdownDescription: "The number of metadata servers that are active. The remaining servers in the cluster will be in standby mode.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(50),
								},
							},

							"active_standby": schema.BoolAttribute{
								Description:         "Whether each active MDS instance will have an active standby with a warm metadata cache for faster failover.If false, standbys will still be available, but will not have a warm metadata cache.",
								MarkdownDescription: "Whether each active MDS instance will have an active standby with a warm metadata cache for faster failover.If false, standbys will still be available, but will not have a warm metadata cache.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"annotations": schema.MapAttribute{
								Description:         "The annotations-related configuration to add/set on each Pod related object.",
								MarkdownDescription: "The annotations-related configuration to add/set on each Pod related object.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "The labels-related configuration to add/set on each Pod related object.",
								MarkdownDescription: "The labels-related configuration to add/set on each Pod related object.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "ProbeSpec is a wrapper around Probe so it can be enabled or disabled for a Ceph daemon",
								MarkdownDescription: "ProbeSpec is a wrapper around Probe so it can be enabled or disabled for a Ceph daemon",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.BoolAttribute{
										Description:         "Disabled determines whether probe is disable or not",
										MarkdownDescription: "Disabled determines whether probe is disable or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"probe": schema.SingleNestedAttribute{
										Description:         "Probe describes a health check to be performed against a container to determine whether it isalive or ready to receive traffic.",
										MarkdownDescription: "Probe describes a health check to be performed against a container to determine whether it isalive or ready to receive traffic.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
														Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"grpc": schema.SingleNestedAttribute{
												Description:         "GRPC specifies an action involving a GRPC port.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
														MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"service": schema.StringAttribute{
														Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
														MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "HTTPGet specifies the http request to perform.",
												MarkdownDescription: "HTTPGet specifies the http request to perform.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_headers": schema.ListNestedAttribute{
														Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
														MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																	MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "The header field value",
																	MarkdownDescription: "The header field value",
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

													"path": schema.StringAttribute{
														Description:         "Path to access on the HTTP server.",
														MarkdownDescription: "Path to access on the HTTP server.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
														MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_socket": schema.SingleNestedAttribute{
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Optional: Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_grace_period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"node_affinity": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"preference": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

											"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"node_selector_terms": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

											"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
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

									"pod_anti_affinity": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

											"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
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

									"topology_spread_constraints": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_skew": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min_domains": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_affinity_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_taints_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"when_unsatisfiable": schema.StringAttribute{
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

							"priority_class_name": schema.StringAttribute{
								Description:         "PriorityClassName sets priority classes on components",
								MarkdownDescription: "PriorityClassName sets priority classes on components",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "The resource requirements for the mds pods",
								MarkdownDescription: "The resource requirements for the mds pods",
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

												"request": schema.StringAttribute{
													Description:         "Request is the name chosen for a request in the referenced claim.If empty, everything from the claim is made available, otherwiseonly the result of this request.",
													MarkdownDescription: "Request is the name chosen for a request in the referenced claim.If empty, everything from the claim is made available, otherwiseonly the result of this request.",
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

							"startup_probe": schema.SingleNestedAttribute{
								Description:         "ProbeSpec is a wrapper around Probe so it can be enabled or disabled for a Ceph daemon",
								MarkdownDescription: "ProbeSpec is a wrapper around Probe so it can be enabled or disabled for a Ceph daemon",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.BoolAttribute{
										Description:         "Disabled determines whether probe is disable or not",
										MarkdownDescription: "Disabled determines whether probe is disable or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"probe": schema.SingleNestedAttribute{
										Description:         "Probe describes a health check to be performed against a container to determine whether it isalive or ready to receive traffic.",
										MarkdownDescription: "Probe describes a health check to be performed against a container to determine whether it isalive or ready to receive traffic.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
														Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"grpc": schema.SingleNestedAttribute{
												Description:         "GRPC specifies an action involving a GRPC port.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
														MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"service": schema.StringAttribute{
														Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
														MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "HTTPGet specifies the http request to perform.",
												MarkdownDescription: "HTTPGet specifies the http request to perform.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_headers": schema.ListNestedAttribute{
														Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
														MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																	MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "The header field value",
																	MarkdownDescription: "The header field value",
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

													"path": schema.StringAttribute{
														Description:         "Path to access on the HTTP server.",
														MarkdownDescription: "Path to access on the HTTP server.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
														MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_socket": schema.SingleNestedAttribute{
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Optional: Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_grace_period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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

					"mirroring": schema.SingleNestedAttribute{
						Description:         "The mirroring settings",
						MarkdownDescription: "The mirroring settings",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled whether this filesystem is mirrored or not",
								MarkdownDescription: "Enabled whether this filesystem is mirrored or not",
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

							"snapshot_retention": schema.ListNestedAttribute{
								Description:         "Retention is the retention policy for a snapshot scheduleOne path has exactly one retention policy.A policy can however contain multiple count-time period pairs in order to specify complex retention policies",
								MarkdownDescription: "Retention is the retention policy for a snapshot scheduleOne path has exactly one retention policy.A policy can however contain multiple count-time period pairs in order to specify complex retention policies",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description:         "Duration represents the retention duration for a snapshot",
											MarkdownDescription: "Duration represents the retention duration for a snapshot",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path is the path to snapshot",
											MarkdownDescription: "Path is the path to snapshot",
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

							"snapshot_schedules": schema.ListNestedAttribute{
								Description:         "SnapshotSchedules is the scheduling of snapshot for mirrored filesystems",
								MarkdownDescription: "SnapshotSchedules is the scheduling of snapshot for mirrored filesystems",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"interval": schema.StringAttribute{
											Description:         "Interval represent the periodicity of the snapshot.",
											MarkdownDescription: "Interval represent the periodicity of the snapshot.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path is the path to snapshot, only valid for CephFS",
											MarkdownDescription: "Path is the path to snapshot, only valid for CephFS",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"start_time": schema.StringAttribute{
											Description:         "StartTime indicates when to start the snapshot",
											MarkdownDescription: "StartTime indicates when to start the snapshot",
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

					"preserve_filesystem_on_delete": schema.BoolAttribute{
						Description:         "Preserve the fs in the cluster on CephFilesystem CR deletion. Setting this to true automatically implies PreservePoolsOnDelete is true.",
						MarkdownDescription: "Preserve the fs in the cluster on CephFilesystem CR deletion. Setting this to true automatically implies PreservePoolsOnDelete is true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preserve_pools_on_delete": schema.BoolAttribute{
						Description:         "Preserve pools on filesystem deletion",
						MarkdownDescription: "Preserve pools on filesystem deletion",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"status_check": schema.SingleNestedAttribute{
						Description:         "The mirroring statusCheck",
						MarkdownDescription: "The mirroring statusCheck",
						Attributes: map[string]schema.Attribute{
							"mirror": schema.SingleNestedAttribute{
								Description:         "HealthCheckSpec represents the health check of an object store bucket",
								MarkdownDescription: "HealthCheckSpec represents the health check of an object store bucket",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
										MarkdownDescription: "Interval is the internal in second or minute for the health check to run like 60s for 60 seconds",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout": schema.StringAttribute{
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

func (r *CephRookIoCephFilesystemV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_filesystem_v1_manifest")

	var model CephRookIoCephFilesystemV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephFilesystem")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
