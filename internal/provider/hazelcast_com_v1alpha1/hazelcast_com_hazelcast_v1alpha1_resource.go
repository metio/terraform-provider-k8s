/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	"strings"
)

var (
	_ resource.Resource                = &HazelcastComHazelcastV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &HazelcastComHazelcastV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &HazelcastComHazelcastV1Alpha1Resource{}
)

func NewHazelcastComHazelcastV1Alpha1Resource() resource.Resource {
	return &HazelcastComHazelcastV1Alpha1Resource{}
}

type HazelcastComHazelcastV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type HazelcastComHazelcastV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdvancedNetwork *struct {
			MemberServerSocketEndpointConfig *struct {
				Interfaces *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
			} `tfsdk:"member_server_socket_endpoint_config" json:"memberServerSocketEndpointConfig,omitempty"`
			Wan *[]struct {
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				PortCount   *int64  `tfsdk:"port_count" json:"portCount,omitempty"`
				ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"wan" json:"wan,omitempty"`
		} `tfsdk:"advanced_network" json:"advancedNetwork,omitempty"`
		Agent *struct {
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			Version    *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"agent" json:"agent,omitempty"`
		ClusterName             *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterSize             *int64  `tfsdk:"cluster_size" json:"clusterSize,omitempty"`
		CustomConfigCmName      *string `tfsdk:"custom_config_cm_name" json:"customConfigCmName,omitempty"`
		DurableExecutorServices *[]struct {
			Capacity   *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			Durability *int64  `tfsdk:"durability" json:"durability,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			PoolSize   *int64  `tfsdk:"pool_size" json:"poolSize,omitempty"`
		} `tfsdk:"durable_executor_services" json:"durableExecutorServices,omitempty"`
		ExecutorServices *[]struct {
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			PoolSize      *int64  `tfsdk:"pool_size" json:"poolSize,omitempty"`
			QueueCapacity *int64  `tfsdk:"queue_capacity" json:"queueCapacity,omitempty"`
		} `tfsdk:"executor_services" json:"executorServices,omitempty"`
		ExposeExternally *struct {
			DiscoveryServiceType *string `tfsdk:"discovery_service_type" json:"discoveryServiceType,omitempty"`
			MemberAccess         *string `tfsdk:"member_access" json:"memberAccess,omitempty"`
			Type                 *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"expose_externally" json:"exposeExternally,omitempty"`
		HighAvailabilityMode *string `tfsdk:"high_availability_mode" json:"highAvailabilityMode,omitempty"`
		ImagePullPolicy      *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets     *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		Jet *struct {
			BucketConfig *struct {
				BucketURI  *string `tfsdk:"bucket_uri" json:"bucketURI,omitempty"`
				Secret     *string `tfsdk:"secret" json:"secret,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"bucket_config" json:"bucketConfig,omitempty"`
			ConfigMaps   *[]string `tfsdk:"config_maps" json:"configMaps,omitempty"`
			EdgeDefaults *struct {
				PacketSizeLimit         *int64 `tfsdk:"packet_size_limit" json:"packetSizeLimit,omitempty"`
				QueueSize               *int64 `tfsdk:"queue_size" json:"queueSize,omitempty"`
				ReceiveWindowMultiplier *int64 `tfsdk:"receive_window_multiplier" json:"receiveWindowMultiplier,omitempty"`
			} `tfsdk:"edge_defaults" json:"edgeDefaults,omitempty"`
			Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Instance *struct {
				BackupCount                    *int64 `tfsdk:"backup_count" json:"backupCount,omitempty"`
				CooperativeThreadCount         *int64 `tfsdk:"cooperative_thread_count" json:"cooperativeThreadCount,omitempty"`
				FlowControlPeriodMillis        *int64 `tfsdk:"flow_control_period_millis" json:"flowControlPeriodMillis,omitempty"`
				LosslessRestartEnabled         *bool  `tfsdk:"lossless_restart_enabled" json:"losslessRestartEnabled,omitempty"`
				MaxProcessorAccumulatedRecords *int64 `tfsdk:"max_processor_accumulated_records" json:"maxProcessorAccumulatedRecords,omitempty"`
				ScaleUpDelayMillis             *int64 `tfsdk:"scale_up_delay_millis" json:"scaleUpDelayMillis,omitempty"`
			} `tfsdk:"instance" json:"instance,omitempty"`
			RemoteURLs            *[]string `tfsdk:"remote_urls" json:"remoteURLs,omitempty"`
			ResourceUploadEnabled *bool     `tfsdk:"resource_upload_enabled" json:"resourceUploadEnabled,omitempty"`
		} `tfsdk:"jet" json:"jet,omitempty"`
		Jvm *struct {
			Args *[]string `tfsdk:"args" json:"args,omitempty"`
			Gc   *struct {
				Collector *string `tfsdk:"collector" json:"collector,omitempty"`
				Logging   *bool   `tfsdk:"logging" json:"logging,omitempty"`
			} `tfsdk:"gc" json:"gc,omitempty"`
			Memory *struct {
				InitialRAMPercentage *string `tfsdk:"initial_ram_percentage" json:"initialRAMPercentage,omitempty"`
				MaxRAMPercentage     *string `tfsdk:"max_ram_percentage" json:"maxRAMPercentage,omitempty"`
				MinRAMPercentage     *string `tfsdk:"min_ram_percentage" json:"minRAMPercentage,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
		} `tfsdk:"jvm" json:"jvm,omitempty"`
		LicenseKeySecret     *string `tfsdk:"license_key_secret" json:"licenseKeySecret,omitempty"`
		LicenseKeySecretName *string `tfsdk:"license_key_secret_name" json:"licenseKeySecretName,omitempty"`
		LoggingLevel         *string `tfsdk:"logging_level" json:"loggingLevel,omitempty"`
		ManagementCenter     *struct {
			ConsoleEnabled    *bool `tfsdk:"console_enabled" json:"consoleEnabled,omitempty"`
			DataAccessEnabled *bool `tfsdk:"data_access_enabled" json:"dataAccessEnabled,omitempty"`
			ScriptingEnabled  *bool `tfsdk:"scripting_enabled" json:"scriptingEnabled,omitempty"`
		} `tfsdk:"management_center" json:"managementCenter,omitempty"`
		NativeMemory *struct {
			AllocatorType           *string `tfsdk:"allocator_type" json:"allocatorType,omitempty"`
			MetadataSpacePercentage *int64  `tfsdk:"metadata_space_percentage" json:"metadataSpacePercentage,omitempty"`
			MinBlockSize            *int64  `tfsdk:"min_block_size" json:"minBlockSize,omitempty"`
			PageSize                *int64  `tfsdk:"page_size" json:"pageSize,omitempty"`
			Size                    *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"native_memory" json:"nativeMemory,omitempty"`
		Persistence *struct {
			BaseDir                   *string `tfsdk:"base_dir" json:"baseDir,omitempty"`
			ClusterDataRecoveryPolicy *string `tfsdk:"cluster_data_recovery_policy" json:"clusterDataRecoveryPolicy,omitempty"`
			DataRecoveryTimeout       *int64  `tfsdk:"data_recovery_timeout" json:"dataRecoveryTimeout,omitempty"`
			Pvc                       *struct {
				AccessModes      *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				RequestStorage   *string   `tfsdk:"request_storage" json:"requestStorage,omitempty"`
				StorageClassName *string   `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			} `tfsdk:"pvc" json:"pvc,omitempty"`
			Restore *struct {
				BucketConfig *struct {
					BucketURI  *string `tfsdk:"bucket_uri" json:"bucketURI,omitempty"`
					Secret     *string `tfsdk:"secret" json:"secret,omitempty"`
					SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"bucket_config" json:"bucketConfig,omitempty"`
				HotBackupResourceName *string `tfsdk:"hot_backup_resource_name" json:"hotBackupResourceName,omitempty"`
			} `tfsdk:"restore" json:"restore,omitempty"`
			StartupAction *string `tfsdk:"startup_action" json:"startupAction,omitempty"`
		} `tfsdk:"persistence" json:"persistence,omitempty"`
		Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
		Repository *string            `tfsdk:"repository" json:"repository,omitempty"`
		Resources  *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		ScheduledExecutorServices *[]struct {
			Capacity       *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			CapacityPolicy *string `tfsdk:"capacity_policy" json:"capacityPolicy,omitempty"`
			Durability     *int64  `tfsdk:"durability" json:"durability,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			PoolSize       *int64  `tfsdk:"pool_size" json:"poolSize,omitempty"`
		} `tfsdk:"scheduled_executor_services" json:"scheduledExecutorServices,omitempty"`
		Scheduling *struct {
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
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
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
		} `tfsdk:"scheduling" json:"scheduling,omitempty"`
		Serialization *struct {
			AllowUnsafe          *bool   `tfsdk:"allow_unsafe" json:"allowUnsafe,omitempty"`
			ByteOrder            *string `tfsdk:"byte_order" json:"byteOrder,omitempty"`
			CompactSerialization *struct {
				Classes     *[]string `tfsdk:"classes" json:"classes,omitempty"`
				Serializers *[]string `tfsdk:"serializers" json:"serializers,omitempty"`
			} `tfsdk:"compact_serialization" json:"compactSerialization,omitempty"`
			DataSerializableFactories *[]string `tfsdk:"data_serializable_factories" json:"dataSerializableFactories,omitempty"`
			EnableCompression         *bool     `tfsdk:"enable_compression" json:"enableCompression,omitempty"`
			EnableSharedObject        *bool     `tfsdk:"enable_shared_object" json:"enableSharedObject,omitempty"`
			GlobalSerializer          *struct {
				ClassName                 *string `tfsdk:"class_name" json:"className,omitempty"`
				OverrideJavaSerialization *bool   `tfsdk:"override_java_serialization" json:"overrideJavaSerialization,omitempty"`
			} `tfsdk:"global_serializer" json:"globalSerializer,omitempty"`
			JavaSerializationFilter *struct {
				Blacklist *struct {
					Classes  *[]string `tfsdk:"classes" json:"classes,omitempty"`
					Packages *[]string `tfsdk:"packages" json:"packages,omitempty"`
					Prefixes *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
				} `tfsdk:"blacklist" json:"blacklist,omitempty"`
				Whitelist *struct {
					Classes  *[]string `tfsdk:"classes" json:"classes,omitempty"`
					Packages *[]string `tfsdk:"packages" json:"packages,omitempty"`
					Prefixes *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
				} `tfsdk:"whitelist" json:"whitelist,omitempty"`
			} `tfsdk:"java_serialization_filter" json:"javaSerializationFilter,omitempty"`
			OverrideDefaultSerializers *bool     `tfsdk:"override_default_serializers" json:"overrideDefaultSerializers,omitempty"`
			PortableFactories          *[]string `tfsdk:"portable_factories" json:"portableFactories,omitempty"`
			Serializers                *[]struct {
				ClassName *string `tfsdk:"class_name" json:"className,omitempty"`
				TypeClass *string `tfsdk:"type_class" json:"typeClass,omitempty"`
			} `tfsdk:"serializers" json:"serializers,omitempty"`
		} `tfsdk:"serialization" json:"serialization,omitempty"`
		Sql *struct {
			CatalogPersistenceEnabled *bool  `tfsdk:"catalog_persistence_enabled" json:"catalogPersistenceEnabled,omitempty"`
			StatementTimeout          *int64 `tfsdk:"statement_timeout" json:"statementTimeout,omitempty"`
		} `tfsdk:"sql" json:"sql,omitempty"`
		Tls *struct {
			MutualAuthentication *string `tfsdk:"mutual_authentication" json:"mutualAuthentication,omitempty"`
			SecretName           *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		UserCodeDeployment *struct {
			BucketConfig *struct {
				BucketURI  *string `tfsdk:"bucket_uri" json:"bucketURI,omitempty"`
				Secret     *string `tfsdk:"secret" json:"secret,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"bucket_config" json:"bucketConfig,omitempty"`
			ClientEnabled   *bool     `tfsdk:"client_enabled" json:"clientEnabled,omitempty"`
			ConfigMaps      *[]string `tfsdk:"config_maps" json:"configMaps,omitempty"`
			RemoteURLs      *[]string `tfsdk:"remote_urls" json:"remoteURLs,omitempty"`
			TriggerSequence *string   `tfsdk:"trigger_sequence" json:"triggerSequence,omitempty"`
		} `tfsdk:"user_code_deployment" json:"userCodeDeployment,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HazelcastComHazelcastV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_hazelcast_v1alpha1"
}

func (r *HazelcastComHazelcastV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Hazelcast is the Schema for the hazelcasts API",
		MarkdownDescription: "Hazelcast is the Schema for the hazelcasts API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
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
				Description:         "Initial values will be filled with its fields' default values.",
				MarkdownDescription: "Initial values will be filled with its fields' default values.",
				Attributes: map[string]schema.Attribute{
					"advanced_network": schema.SingleNestedAttribute{
						Description:         "Hazelcast Advanced Network configuration",
						MarkdownDescription: "Hazelcast Advanced Network configuration",
						Attributes: map[string]schema.Attribute{
							"member_server_socket_endpoint_config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"interfaces": schema.ListAttribute{
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

							"wan": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port_count": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_type": schema.StringAttribute{
											Description:         "Service Type string describes ingress methods for a service",
											MarkdownDescription: "Service Type string describes ingress methods for a service",
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

					"agent": schema.SingleNestedAttribute{
						Description:         "B&R Agent configurations",
						MarkdownDescription: "B&R Agent configurations",
						Attributes: map[string]schema.Attribute{
							"repository": schema.StringAttribute{
								Description:         "Repository to pull Hazelcast Platform Operator Agent(https://github.com/hazelcast/platform-operator-agent)",
								MarkdownDescription: "Repository to pull Hazelcast Platform Operator Agent(https://github.com/hazelcast/platform-operator-agent)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version of Hazelcast Platform Operator Agent.",
								MarkdownDescription: "Version of Hazelcast Platform Operator Agent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "Name of the Hazelcast cluster.",
						MarkdownDescription: "Name of the Hazelcast cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_size": schema.Int64Attribute{
						Description:         "Number of Hazelcast members in the cluster.",
						MarkdownDescription: "Number of Hazelcast members in the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"custom_config_cm_name": schema.StringAttribute{
						Description:         "Name of the ConfigMap with the Hazelcast custom configuration. This configuration from the ConfigMap might be overridden by the Hazelcast CR configuration.",
						MarkdownDescription: "Name of the ConfigMap with the Hazelcast custom configuration. This configuration from the ConfigMap might be overridden by the Hazelcast CR configuration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"durable_executor_services": schema.ListNestedAttribute{
						Description:         "Durable Executor Service configurations, see https://docs.hazelcast.com/hazelcast/latest/computing/durable-executor-service",
						MarkdownDescription: "Durable Executor Service configurations, see https://docs.hazelcast.com/hazelcast/latest/computing/durable-executor-service",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"capacity": schema.Int64Attribute{
									Description:         "Capacity of the executor task per partition.",
									MarkdownDescription: "Capacity of the executor task per partition.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"durability": schema.Int64Attribute{
									Description:         "Durability of the executor.",
									MarkdownDescription: "Durability of the executor.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
									},
								},

								"name": schema.StringAttribute{
									Description:         "The name of the executor service",
									MarkdownDescription: "The name of the executor service",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pool_size": schema.Int64Attribute{
									Description:         "The number of executor threads per member.",
									MarkdownDescription: "The number of executor threads per member.",
									Required:            false,
									Optional:            true,
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

					"executor_services": schema.ListNestedAttribute{
						Description:         "Java Executor Service configurations, see https://docs.hazelcast.com/hazelcast/latest/computing/executor-service",
						MarkdownDescription: "Java Executor Service configurations, see https://docs.hazelcast.com/hazelcast/latest/computing/executor-service",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of the executor service",
									MarkdownDescription: "The name of the executor service",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pool_size": schema.Int64Attribute{
									Description:         "The number of executor threads per member.",
									MarkdownDescription: "The number of executor threads per member.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
									},
								},

								"queue_capacity": schema.Int64Attribute{
									Description:         "Task queue capacity of the executor.",
									MarkdownDescription: "Task queue capacity of the executor.",
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

					"expose_externally": schema.SingleNestedAttribute{
						Description:         "Configuration to expose Hazelcast cluster to external clients.",
						MarkdownDescription: "Configuration to expose Hazelcast cluster to external clients.",
						Attributes: map[string]schema.Attribute{
							"discovery_service_type": schema.StringAttribute{
								Description:         "Type of the service used to discover Hazelcast cluster.",
								MarkdownDescription: "Type of the service used to discover Hazelcast cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"member_access": schema.StringAttribute{
								Description:         "How each member is accessed from the external client. Only available for 'Smart' client and valid values are: - 'NodePortExternalIP' (default): each member is accessed by the NodePort service and the node external IP/hostname - 'NodePortNodeName': each member is accessed by the NodePort service and the node name - 'LoadBalancer': each member is accessed by the LoadBalancer service external address",
								MarkdownDescription: "How each member is accessed from the external client. Only available for 'Smart' client and valid values are: - 'NodePortExternalIP' (default): each member is accessed by the NodePort service and the node external IP/hostname - 'NodePortNodeName': each member is accessed by the NodePort service and the node name - 'LoadBalancer': each member is accessed by the LoadBalancer service external address",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("NodePortExternalIP", "NodePortNodeName", "LoadBalancer"),
								},
							},

							"type": schema.StringAttribute{
								Description:         "Specifies how members are exposed. Valid values are: - 'Smart' (default): each member pod is exposed with a separate external address - 'Unisocket': all member pods are exposed with one external address",
								MarkdownDescription: "Specifies how members are exposed. Valid values are: - 'Smart' (default): each member pod is exposed with a separate external address - 'Unisocket': all member pods are exposed with one external address",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Smart", "Unisocket"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"high_availability_mode": schema.StringAttribute{
						Description:         "Configuration to create clusters resilient to node and zone failures",
						MarkdownDescription: "Configuration to create clusters resilient to node and zone failures",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("NODE", "ZONE"),
						},
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "Pull policy for the Hazelcast Platform image",
						MarkdownDescription: "Pull policy for the Hazelcast Platform image",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "Image pull secrets for the Hazelcast Platform image",
						MarkdownDescription: "Image pull secrets for the Hazelcast Platform image",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

					"jet": schema.SingleNestedAttribute{
						Description:         "Jet Engine configuration",
						MarkdownDescription: "Jet Engine configuration",
						Attributes: map[string]schema.Attribute{
							"bucket_config": schema.SingleNestedAttribute{
								Description:         "Bucket config from where the JAR files will be downloaded.",
								MarkdownDescription: "Bucket config from where the JAR files will be downloaded.",
								Attributes: map[string]schema.Attribute{
									"bucket_uri": schema.StringAttribute{
										Description:         "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										MarkdownDescription: "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(6),
										},
									},

									"secret": schema.StringAttribute{
										Description:         "secret is a deprecated alias for secretName.",
										MarkdownDescription: "secret is a deprecated alias for secretName.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "Name of the secret with credentials for cloud providers.",
										MarkdownDescription: "Name of the secret with credentials for cloud providers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "Names of the list of ConfigMaps. Files in each ConfigMap will be downloaded.",
								MarkdownDescription: "Names of the list of ConfigMaps. Files in each ConfigMap will be downloaded.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"edge_defaults": schema.SingleNestedAttribute{
								Description:         "Jet Edge Defaults Configuration",
								MarkdownDescription: "Jet Edge Defaults Configuration",
								Attributes: map[string]schema.Attribute{
									"packet_size_limit": schema.Int64Attribute{
										Description:         "Limits the size of the packet in bytes.",
										MarkdownDescription: "Limits the size of the packet in bytes.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"queue_size": schema.Int64Attribute{
										Description:         "Sets the capacity of processor-to-processor concurrent queues.",
										MarkdownDescription: "Sets the capacity of processor-to-processor concurrent queues.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"receive_window_multiplier": schema.Int64Attribute{
										Description:         "Sets the scaling factor used by the adaptive receive window sizing function.",
										MarkdownDescription: "Sets the scaling factor used by the adaptive receive window sizing function.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "When false, disables Jet Engine.",
								MarkdownDescription: "When false, disables Jet Engine.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance": schema.SingleNestedAttribute{
								Description:         "Jet Instance Configuration",
								MarkdownDescription: "Jet Instance Configuration",
								Attributes: map[string]schema.Attribute{
									"backup_count": schema.Int64Attribute{
										Description:         "The number of synchronous backups to configure on the IMap that Jet needs internally to store job metadata and snapshots.",
										MarkdownDescription: "The number of synchronous backups to configure on the IMap that Jet needs internally to store job metadata and snapshots.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtMost(6),
										},
									},

									"cooperative_thread_count": schema.Int64Attribute{
										Description:         "The number of threads Jet creates in its cooperative multithreading pool. Its default value is the number of cores",
										MarkdownDescription: "The number of threads Jet creates in its cooperative multithreading pool. Its default value is the number of cores",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"flow_control_period_millis": schema.Int64Attribute{
										Description:         "The duration of the interval between flow-control packets.",
										MarkdownDescription: "The duration of the interval between flow-control packets.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"lossless_restart_enabled": schema.BoolAttribute{
										Description:         "Specifies whether the Lossless Cluster Restart feature is enabled.",
										MarkdownDescription: "Specifies whether the Lossless Cluster Restart feature is enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_processor_accumulated_records": schema.Int64Attribute{
										Description:         "Specifies the maximum number of records that can be accumulated by any single processor instance. Default value is Long.MAX_VALUE",
										MarkdownDescription: "Specifies the maximum number of records that can be accumulated by any single processor instance. Default value is Long.MAX_VALUE",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scale_up_delay_millis": schema.Int64Attribute{
										Description:         "The delay after which the auto-scaled jobs restart if a new member joins the cluster.",
										MarkdownDescription: "The delay after which the auto-scaled jobs restart if a new member joins the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"remote_urls": schema.ListAttribute{
								Description:         "List of URLs from where the files will be downloaded.",
								MarkdownDescription: "List of URLs from where the files will be downloaded.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_upload_enabled": schema.BoolAttribute{
								Description:         "When true, enables resource uploading for Jet jobs.",
								MarkdownDescription: "When true, enables resource uploading for Jet jobs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm": schema.SingleNestedAttribute{
						Description:         "Hazelcast JVM configuration",
						MarkdownDescription: "Hazelcast JVM configuration",
						Attributes: map[string]schema.Attribute{
							"args": schema.ListAttribute{
								Description:         "Args is for arbitrary JVM arguments",
								MarkdownDescription: "Args is for arbitrary JVM arguments",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gc": schema.SingleNestedAttribute{
								Description:         "GC is for configuring JVM Garbage Collector",
								MarkdownDescription: "GC is for configuring JVM Garbage Collector",
								Attributes: map[string]schema.Attribute{
									"collector": schema.StringAttribute{
										Description:         "Collector is the Garbage Collector type",
										MarkdownDescription: "Collector is the Garbage Collector type",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Serial", "Parallel", "G1"),
										},
									},

									"logging": schema.BoolAttribute{
										Description:         "Logging enables logging when set to true",
										MarkdownDescription: "Logging enables logging when set to true",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"memory": schema.SingleNestedAttribute{
								Description:         "Memory is a JVM memory configuration",
								MarkdownDescription: "Memory is a JVM memory configuration",
								Attributes: map[string]schema.Attribute{
									"initial_ram_percentage": schema.StringAttribute{
										Description:         "InitialRAMPercentage configures JVM initial heap size",
										MarkdownDescription: "InitialRAMPercentage configures JVM initial heap size",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_ram_percentage": schema.StringAttribute{
										Description:         "MaxRAMPercentage sets the maximum heap size for a JVM",
										MarkdownDescription: "MaxRAMPercentage sets the maximum heap size for a JVM",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_ram_percentage": schema.StringAttribute{
										Description:         "MinRAMPercentage sets the minimum heap size for a JVM",
										MarkdownDescription: "MinRAMPercentage sets the minimum heap size for a JVM",
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

					"license_key_secret": schema.StringAttribute{
						Description:         "licenseKeySecret is a deprecated alias for licenseKeySecretName.",
						MarkdownDescription: "licenseKeySecret is a deprecated alias for licenseKeySecretName.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"license_key_secret_name": schema.StringAttribute{
						Description:         "Name of the secret with Hazelcast Enterprise License Key.",
						MarkdownDescription: "Name of the secret with Hazelcast Enterprise License Key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"logging_level": schema.StringAttribute{
						Description:         "Logging level for Hazelcast members",
						MarkdownDescription: "Logging level for Hazelcast members",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("OFF", "FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE", "ALL"),
						},
					},

					"management_center": schema.SingleNestedAttribute{
						Description:         "Hazelcast Management Center Configuration",
						MarkdownDescription: "Hazelcast Management Center Configuration",
						Attributes: map[string]schema.Attribute{
							"console_enabled": schema.BoolAttribute{
								Description:         "Allows you to execute commands from a built-in console in the user interface.",
								MarkdownDescription: "Allows you to execute commands from a built-in console in the user interface.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data_access_enabled": schema.BoolAttribute{
								Description:         "Allows you to access contents of Hazelcast data structures via SQL Browser or Map Browser.",
								MarkdownDescription: "Allows you to access contents of Hazelcast data structures via SQL Browser or Map Browser.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scripting_enabled": schema.BoolAttribute{
								Description:         "Allows you to execute scripts that can automate interactions with the cluster.",
								MarkdownDescription: "Allows you to execute scripts that can automate interactions with the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"native_memory": schema.SingleNestedAttribute{
						Description:         "Hazelcast Native Memory (HD Memory) configuration",
						MarkdownDescription: "Hazelcast Native Memory (HD Memory) configuration",
						Attributes: map[string]schema.Attribute{
							"allocator_type": schema.StringAttribute{
								Description:         "AllocatorType specifies one of 2 types of mechanism for allocating memory to HD",
								MarkdownDescription: "AllocatorType specifies one of 2 types of mechanism for allocating memory to HD",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("STANDARD", "POOLED"),
								},
							},

							"metadata_space_percentage": schema.Int64Attribute{
								Description:         "MetadataSpacePercentage defines percentage of the allocated native memory that is used for the metadata of other map components such as index (for predicates), offset, etc.",
								MarkdownDescription: "MetadataSpacePercentage defines percentage of the allocated native memory that is used for the metadata of other map components such as index (for predicates), offset, etc.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_block_size": schema.Int64Attribute{
								Description:         "MinBlockSize is the size of smallest block that will be allocated. It is used only by the POOLED memory allocator.",
								MarkdownDescription: "MinBlockSize is the size of smallest block that will be allocated. It is used only by the POOLED memory allocator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"page_size": schema.Int64Attribute{
								Description:         "PageSize is the size of the page in bytes to allocate memory as a block. It is used only by the POOLED memory allocator.",
								MarkdownDescription: "PageSize is the size of the page in bytes to allocate memory as a block. It is used only by the POOLED memory allocator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size of the total native memory to allocate",
								MarkdownDescription: "Size of the total native memory to allocate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistence": schema.SingleNestedAttribute{
						Description:         "Persistence configuration",
						MarkdownDescription: "Persistence configuration",
						Attributes: map[string]schema.Attribute{
							"base_dir": schema.StringAttribute{
								Description:         "Persistence base directory.",
								MarkdownDescription: "Persistence base directory.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"cluster_data_recovery_policy": schema.StringAttribute{
								Description:         "Configuration of the cluster recovery strategy.",
								MarkdownDescription: "Configuration of the cluster recovery strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("FullRecoveryOnly", "PartialRecoveryMostRecent", "PartialRecoveryMostComplete"),
								},
							},

							"data_recovery_timeout": schema.Int64Attribute{
								Description:         "DataRecoveryTimeout is timeout for each step of data recovery in seconds. Maximum timeout is equal to DataRecoveryTimeout*2 (for each step: validation and data-load).",
								MarkdownDescription: "DataRecoveryTimeout is timeout for each step of data recovery in seconds. Maximum timeout is equal to DataRecoveryTimeout*2 (for each step: validation and data-load).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc": schema.SingleNestedAttribute{
								Description:         "Configuration of PersistenceVolumeClaim.",
								MarkdownDescription: "Configuration of PersistenceVolumeClaim.",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "AccessModes contains the actual access modes of the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "AccessModes contains the actual access modes of the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_storage": schema.StringAttribute{
										Description:         "A description of the PVC request capacity.",
										MarkdownDescription: "A description of the PVC request capacity.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_class_name": schema.StringAttribute{
										Description:         "Name of StorageClass which this persistent volume belongs to.",
										MarkdownDescription: "Name of StorageClass which this persistent volume belongs to.",
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
								Description:         "Restore configuration",
								MarkdownDescription: "Restore configuration",
								Attributes: map[string]schema.Attribute{
									"bucket_config": schema.SingleNestedAttribute{
										Description:         "Bucket Configuration from which the backup will be downloaded.",
										MarkdownDescription: "Bucket Configuration from which the backup will be downloaded.",
										Attributes: map[string]schema.Attribute{
											"bucket_uri": schema.StringAttribute{
												Description:         "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
												MarkdownDescription: "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(6),
												},
											},

											"secret": schema.StringAttribute{
												Description:         "secret is a deprecated alias for secretName.",
												MarkdownDescription: "secret is a deprecated alias for secretName.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_name": schema.StringAttribute{
												Description:         "Name of the secret with credentials for cloud providers.",
												MarkdownDescription: "Name of the secret with credentials for cloud providers.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"hot_backup_resource_name": schema.StringAttribute{
										Description:         "Name of the HotBackup resource from which backup will be fetched.",
										MarkdownDescription: "Name of the HotBackup resource from which backup will be fetched.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_action": schema.StringAttribute{
								Description:         "StartupAction represents the action triggered when the cluster starts to force the cluster startup.",
								MarkdownDescription: "StartupAction represents the action triggered when the cluster starts to force the cluster startup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ForceStart", "PartialStart"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"properties": schema.MapAttribute{
						Description:         "Hazelcast system properties, see https://docs.hazelcast.com/hazelcast/latest/system-properties",
						MarkdownDescription: "Hazelcast system properties, see https://docs.hazelcast.com/hazelcast/latest/system-properties",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repository": schema.StringAttribute{
						Description:         "Repository to pull the Hazelcast Platform image from.",
						MarkdownDescription: "Repository to pull the Hazelcast Platform image from.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Compute Resources required by the Hazelcast container.",
						MarkdownDescription: "Compute Resources required by the Hazelcast container.",
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
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"scheduled_executor_services": schema.ListNestedAttribute{
						Description:         "Scheduled Executor Service configurations, see https://docs.hazelcast.com/hazelcast/latest/computing/scheduled-executor-service",
						MarkdownDescription: "Scheduled Executor Service configurations, see https://docs.hazelcast.com/hazelcast/latest/computing/scheduled-executor-service",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"capacity": schema.Int64Attribute{
									Description:         "Capacity of the executor task per partition.",
									MarkdownDescription: "Capacity of the executor task per partition.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"capacity_policy": schema.StringAttribute{
									Description:         "The active policy for the capacity setting.",
									MarkdownDescription: "The active policy for the capacity setting.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"durability": schema.Int64Attribute{
									Description:         "Durability of the executor.",
									MarkdownDescription: "Durability of the executor.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
									},
								},

								"name": schema.StringAttribute{
									Description:         "The name of the executor service",
									MarkdownDescription: "The name of the executor service",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pool_size": schema.Int64Attribute{
									Description:         "The number of executor threads per member.",
									MarkdownDescription: "The number of executor threads per member.",
									Required:            false,
									Optional:            true,
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

					"scheduling": schema.SingleNestedAttribute{
						Description:         "Scheduling details",
						MarkdownDescription: "Scheduling details",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.SingleNestedAttribute{
								Description:         "Affinity",
								MarkdownDescription: "Affinity",
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
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector",
								MarkdownDescription: "NodeSelector",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations",
								MarkdownDescription: "Tolerations",
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
								Description:         "TopologySpreadConstraints",
								MarkdownDescription: "TopologySpreadConstraints",
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
											Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.",
											MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_skew": schema.Int64Attribute{
											Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
											MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"min_domains": schema.Int64Attribute{
											Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
											MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_affinity_policy": schema.StringAttribute{
											Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_taints_policy": schema.StringAttribute{
											Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
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
											Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
											MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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

					"serialization": schema.SingleNestedAttribute{
						Description:         "Hazelcast serialization configuration",
						MarkdownDescription: "Hazelcast serialization configuration",
						Attributes: map[string]schema.Attribute{
							"allow_unsafe": schema.BoolAttribute{
								Description:         "Allow the usage of unsafe.",
								MarkdownDescription: "Allow the usage of unsafe.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"byte_order": schema.StringAttribute{
								Description:         "Specifies the byte order that the serialization will use.",
								MarkdownDescription: "Specifies the byte order that the serialization will use.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Native", "BigEndian", "LittleEndian"),
								},
							},

							"compact_serialization": schema.SingleNestedAttribute{
								Description:         "Configuration attributes the compact serialization.",
								MarkdownDescription: "Configuration attributes the compact serialization.",
								Attributes: map[string]schema.Attribute{
									"classes": schema.ListAttribute{
										Description:         "Classes is the list of class names for which a zero-config serializer will be registered, without implementing an explicit serializer.",
										MarkdownDescription: "Classes is the list of class names for which a zero-config serializer will be registered, without implementing an explicit serializer.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.List{
											listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("serializers")),
										},
									},

									"serializers": schema.ListAttribute{
										Description:         "Serializers is the list of explicit serializers to be registered.",
										MarkdownDescription: "Serializers is the list of explicit serializers to be registered.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.List{
											listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("classes")),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_serializable_factories": schema.ListAttribute{
								Description:         "Lists class implementations of Hazelcast's DataSerializableFactory.",
								MarkdownDescription: "Lists class implementations of Hazelcast's DataSerializableFactory.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_compression": schema.BoolAttribute{
								Description:         "Enables compression when default Java serialization is used.",
								MarkdownDescription: "Enables compression when default Java serialization is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_shared_object": schema.BoolAttribute{
								Description:         "Enables shared object when default Java serialization is used.",
								MarkdownDescription: "Enables shared object when default Java serialization is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"global_serializer": schema.SingleNestedAttribute{
								Description:         "List of global serializers.",
								MarkdownDescription: "List of global serializers.",
								Attributes: map[string]schema.Attribute{
									"class_name": schema.StringAttribute{
										Description:         "Class name of the GlobalSerializer.",
										MarkdownDescription: "Class name of the GlobalSerializer.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"override_java_serialization": schema.BoolAttribute{
										Description:         "If set to true, will replace the internal Java serialization.",
										MarkdownDescription: "If set to true, will replace the internal Java serialization.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"java_serialization_filter": schema.SingleNestedAttribute{
								Description:         "Blacklist and whitelist for deserialized classes when Java serialization is used.",
								MarkdownDescription: "Blacklist and whitelist for deserialized classes when Java serialization is used.",
								Attributes: map[string]schema.Attribute{
									"blacklist": schema.SingleNestedAttribute{
										Description:         "Java deserialization protection Blacklist.",
										MarkdownDescription: "Java deserialization protection Blacklist.",
										Attributes: map[string]schema.Attribute{
											"classes": schema.ListAttribute{
												Description:         "List of class names to be filtered.",
												MarkdownDescription: "List of class names to be filtered.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.List{
													listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("packages"), path.MatchRelative().AtParent().AtName("prefixes")),
												},
											},

											"packages": schema.ListAttribute{
												Description:         "List of packages to be filtered",
												MarkdownDescription: "List of packages to be filtered",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.List{
													listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("classes"), path.MatchRelative().AtParent().AtName("prefixes")),
												},
											},

											"prefixes": schema.ListAttribute{
												Description:         "List of prefixes to be filtered.",
												MarkdownDescription: "List of prefixes to be filtered.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.List{
													listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("classes"), path.MatchRelative().AtParent().AtName("packages")),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.Object{
											objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("whitelist")),
										},
									},

									"whitelist": schema.SingleNestedAttribute{
										Description:         "Java deserialization protection Whitelist.",
										MarkdownDescription: "Java deserialization protection Whitelist.",
										Attributes: map[string]schema.Attribute{
											"classes": schema.ListAttribute{
												Description:         "List of class names to be filtered.",
												MarkdownDescription: "List of class names to be filtered.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.List{
													listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("packages"), path.MatchRelative().AtParent().AtName("prefixes")),
												},
											},

											"packages": schema.ListAttribute{
												Description:         "List of packages to be filtered",
												MarkdownDescription: "List of packages to be filtered",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.List{
													listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("classes"), path.MatchRelative().AtParent().AtName("prefixes")),
												},
											},

											"prefixes": schema.ListAttribute{
												Description:         "List of prefixes to be filtered.",
												MarkdownDescription: "List of prefixes to be filtered.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.List{
													listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("classes"), path.MatchRelative().AtParent().AtName("packages")),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
										Validators: []validator.Object{
											objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("blacklist")),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"override_default_serializers": schema.BoolAttribute{
								Description:         "Allows override of built-in default serializers.",
								MarkdownDescription: "Allows override of built-in default serializers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"portable_factories": schema.ListAttribute{
								Description:         "Lists class implementations of Hazelcast's PortableFactory.",
								MarkdownDescription: "Lists class implementations of Hazelcast's PortableFactory.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"serializers": schema.ListNestedAttribute{
								Description:         "List of serializers (classes) that implemented using Hazelcast's StreamSerializer, ByteArraySerializer etc.",
								MarkdownDescription: "List of serializers (classes) that implemented using Hazelcast's StreamSerializer, ByteArraySerializer etc.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"class_name": schema.StringAttribute{
											Description:         "Class name of the implementation of the serializer class.",
											MarkdownDescription: "Class name of the implementation of the serializer class.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type_class": schema.StringAttribute{
											Description:         "Name of the class that will be serialized via this implementation.",
											MarkdownDescription: "Name of the class that will be serialized via this implementation.",
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

					"sql": schema.SingleNestedAttribute{
						Description:         "Hazelcast SQL configuration",
						MarkdownDescription: "Hazelcast SQL configuration",
						Attributes: map[string]schema.Attribute{
							"catalog_persistence_enabled": schema.BoolAttribute{
								Description:         "CatalogPersistenceEnabled sets whether SQL Catalog persistence is enabled for the node. With SQL Catalog persistence enabled you can restart the whole cluster without losing schema definition objects (such as MAPPINGs, TYPEs, VIEWs and DATA CONNECTIONs). The feature is implemented on top of the Hot Restart feature of Hazelcast which persists the data to disk. If enabled, you have to also configure Hot Restart. Feature is disabled by default.",
								MarkdownDescription: "CatalogPersistenceEnabled sets whether SQL Catalog persistence is enabled for the node. With SQL Catalog persistence enabled you can restart the whole cluster without losing schema definition objects (such as MAPPINGs, TYPEs, VIEWs and DATA CONNECTIONs). The feature is implemented on top of the Hot Restart feature of Hazelcast which persists the data to disk. If enabled, you have to also configure Hot Restart. Feature is disabled by default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"statement_timeout": schema.Int64Attribute{
								Description:         "StatementTimeout defines the timeout in milliseconds that is applied to queries without an explicit timeout.",
								MarkdownDescription: "StatementTimeout defines the timeout in milliseconds that is applied to queries without an explicit timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "Hazelcast TLS configuration",
						MarkdownDescription: "Hazelcast TLS configuration",
						Attributes: map[string]schema.Attribute{
							"mutual_authentication": schema.StringAttribute{
								Description:         "Mutual authentication configuration. It’s None by default which means the client side of connection is not authenticated.",
								MarkdownDescription: "Mutual authentication configuration. It’s None by default which means the client side of connection is not authenticated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("None", "Required", "Optional"),
								},
							},

							"secret_name": schema.StringAttribute{
								Description:         "Name of the secret with TLS certificate and key.",
								MarkdownDescription: "Name of the secret with TLS certificate and key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"user_code_deployment": schema.SingleNestedAttribute{
						Description:         "User Codes to Download into CLASSPATH",
						MarkdownDescription: "User Codes to Download into CLASSPATH",
						Attributes: map[string]schema.Attribute{
							"bucket_config": schema.SingleNestedAttribute{
								Description:         "Bucket config from where the JAR files will be downloaded.",
								MarkdownDescription: "Bucket config from where the JAR files will be downloaded.",
								Attributes: map[string]schema.Attribute{
									"bucket_uri": schema.StringAttribute{
										Description:         "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										MarkdownDescription: "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(6),
										},
									},

									"secret": schema.StringAttribute{
										Description:         "secret is a deprecated alias for secretName.",
										MarkdownDescription: "secret is a deprecated alias for secretName.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "Name of the secret with credentials for cloud providers.",
										MarkdownDescription: "Name of the secret with credentials for cloud providers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_enabled": schema.BoolAttribute{
								Description:         "When true, allows user code deployment from clients.",
								MarkdownDescription: "When true, allows user code deployment from clients.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "Names of the list of ConfigMaps. Files in each ConfigMap will be downloaded.",
								MarkdownDescription: "Names of the list of ConfigMaps. Files in each ConfigMap will be downloaded.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remote_urls": schema.ListAttribute{
								Description:         "List of URLs from where the files will be downloaded.",
								MarkdownDescription: "List of URLs from where the files will be downloaded.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trigger_sequence": schema.StringAttribute{
								Description:         "A string for triggering a rolling restart for re-downloading the user code.",
								MarkdownDescription: "A string for triggering a rolling restart for re-downloading the user code.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": schema.StringAttribute{
						Description:         "Version of Hazelcast Platform.",
						MarkdownDescription: "Version of Hazelcast Platform.",
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

func (r *HazelcastComHazelcastV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *HazelcastComHazelcastV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hazelcast_com_hazelcast_v1alpha1")

	var model HazelcastComHazelcastV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	model.Kind = pointer.String("Hazelcast")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "hazelcasts"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse HazelcastComHazelcastV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *HazelcastComHazelcastV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_hazelcast_v1alpha1")

	var data HazelcastComHazelcastV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "hazelcasts"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse HazelcastComHazelcastV1Alpha1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *HazelcastComHazelcastV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hazelcast_com_hazelcast_v1alpha1")

	var model HazelcastComHazelcastV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	model.Kind = pointer.String("Hazelcast")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "hazelcasts"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse HazelcastComHazelcastV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *HazelcastComHazelcastV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hazelcast_com_hazelcast_v1alpha1")

	var data HazelcastComHazelcastV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "hazelcasts"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *HazelcastComHazelcastV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
