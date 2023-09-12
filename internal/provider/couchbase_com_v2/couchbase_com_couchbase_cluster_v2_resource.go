/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	"regexp"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &CouchbaseComCouchbaseClusterV2Resource{}
	_ resource.ResourceWithConfigure   = &CouchbaseComCouchbaseClusterV2Resource{}
	_ resource.ResourceWithImportState = &CouchbaseComCouchbaseClusterV2Resource{}
)

func NewCouchbaseComCouchbaseClusterV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseClusterV2Resource{}
}

type CouchbaseComCouchbaseClusterV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CouchbaseComCouchbaseClusterV2ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AntiAffinity           *bool `tfsdk:"anti_affinity" json:"antiAffinity,omitempty"`
		AutoResourceAllocation *struct {
			CpuLimits       *string `tfsdk:"cpu_limits" json:"cpuLimits,omitempty"`
			CpuRequests     *string `tfsdk:"cpu_requests" json:"cpuRequests,omitempty"`
			Enabled         *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			OverheadPercent *int64  `tfsdk:"overhead_percent" json:"overheadPercent,omitempty"`
		} `tfsdk:"auto_resource_allocation" json:"autoResourceAllocation,omitempty"`
		AutoscaleStabilizationPeriod *string `tfsdk:"autoscale_stabilization_period" json:"autoscaleStabilizationPeriod,omitempty"`
		Backup                       *struct {
			Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Image            *string            `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Managed        *bool              `tfsdk:"managed" json:"managed,omitempty"`
			NodeSelector   *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			ObjectEndpoint *struct {
				Secret         *string `tfsdk:"secret" json:"secret,omitempty"`
				Url            *string `tfsdk:"url" json:"url,omitempty"`
				UseVirtualPath *bool   `tfsdk:"use_virtual_path" json:"useVirtualPath,omitempty"`
			} `tfsdk:"object_endpoint" json:"objectEndpoint,omitempty"`
			Resources *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			S3Secret *string `tfsdk:"s3_secret" json:"s3Secret,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			Tolerations        *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			UseIAMRole *bool `tfsdk:"use_iam_role" json:"useIAMRole,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		Buckets *struct {
			Managed  *bool `tfsdk:"managed" json:"managed,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Synchronize *bool `tfsdk:"synchronize" json:"synchronize,omitempty"`
		} `tfsdk:"buckets" json:"buckets,omitempty"`
		Cluster *struct {
			AnalyticsServiceMemoryQuota *string `tfsdk:"analytics_service_memory_quota" json:"analyticsServiceMemoryQuota,omitempty"`
			AutoCompaction              *struct {
				DatabaseFragmentationThreshold *struct {
					Percent *int64  `tfsdk:"percent" json:"percent,omitempty"`
					Size    *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"database_fragmentation_threshold" json:"databaseFragmentationThreshold,omitempty"`
				ParallelCompaction *bool `tfsdk:"parallel_compaction" json:"parallelCompaction,omitempty"`
				TimeWindow         *struct {
					AbortCompactionOutsideWindow *bool   `tfsdk:"abort_compaction_outside_window" json:"abortCompactionOutsideWindow,omitempty"`
					End                          *string `tfsdk:"end" json:"end,omitempty"`
					Start                        *string `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"time_window" json:"timeWindow,omitempty"`
				TombstonePurgeInterval     *string `tfsdk:"tombstone_purge_interval" json:"tombstonePurgeInterval,omitempty"`
				ViewFragmentationThreshold *struct {
					Percent *int64  `tfsdk:"percent" json:"percent,omitempty"`
					Size    *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"view_fragmentation_threshold" json:"viewFragmentationThreshold,omitempty"`
			} `tfsdk:"auto_compaction" json:"autoCompaction,omitempty"`
			AutoFailoverMaxCount                   *int64  `tfsdk:"auto_failover_max_count" json:"autoFailoverMaxCount,omitempty"`
			AutoFailoverOnDataDiskIssues           *bool   `tfsdk:"auto_failover_on_data_disk_issues" json:"autoFailoverOnDataDiskIssues,omitempty"`
			AutoFailoverOnDataDiskIssuesTimePeriod *string `tfsdk:"auto_failover_on_data_disk_issues_time_period" json:"autoFailoverOnDataDiskIssuesTimePeriod,omitempty"`
			AutoFailoverServerGroup                *bool   `tfsdk:"auto_failover_server_group" json:"autoFailoverServerGroup,omitempty"`
			AutoFailoverTimeout                    *string `tfsdk:"auto_failover_timeout" json:"autoFailoverTimeout,omitempty"`
			ClusterName                            *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
			Data                                   *struct {
				ReaderThreads *int64 `tfsdk:"reader_threads" json:"readerThreads,omitempty"`
				WriterThreads *int64 `tfsdk:"writer_threads" json:"writerThreads,omitempty"`
			} `tfsdk:"data" json:"data,omitempty"`
			DataServiceMemoryQuota     *string `tfsdk:"data_service_memory_quota" json:"dataServiceMemoryQuota,omitempty"`
			EventingServiceMemoryQuota *string `tfsdk:"eventing_service_memory_quota" json:"eventingServiceMemoryQuota,omitempty"`
			IndexServiceMemoryQuota    *string `tfsdk:"index_service_memory_quota" json:"indexServiceMemoryQuota,omitempty"`
			IndexStorageSetting        *string `tfsdk:"index_storage_setting" json:"indexStorageSetting,omitempty"`
			Indexer                    *struct {
				LogLevel               *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				MaxRollbackPoints      *int64  `tfsdk:"max_rollback_points" json:"maxRollbackPoints,omitempty"`
				MemorySnapshotInterval *string `tfsdk:"memory_snapshot_interval" json:"memorySnapshotInterval,omitempty"`
				NumReplica             *int64  `tfsdk:"num_replica" json:"numReplica,omitempty"`
				RedistributeIndexes    *bool   `tfsdk:"redistribute_indexes" json:"redistributeIndexes,omitempty"`
				StableSnapshotInterval *string `tfsdk:"stable_snapshot_interval" json:"stableSnapshotInterval,omitempty"`
				StorageMode            *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
				Threads                *int64  `tfsdk:"threads" json:"threads,omitempty"`
			} `tfsdk:"indexer" json:"indexer,omitempty"`
			Query *struct {
				BackfillEnabled         *bool   `tfsdk:"backfill_enabled" json:"backfillEnabled,omitempty"`
				TemporarySpace          *string `tfsdk:"temporary_space" json:"temporarySpace,omitempty"`
				TemporarySpaceUnlimited *bool   `tfsdk:"temporary_space_unlimited" json:"temporarySpaceUnlimited,omitempty"`
			} `tfsdk:"query" json:"query,omitempty"`
			QueryServiceMemoryQuota  *string `tfsdk:"query_service_memory_quota" json:"queryServiceMemoryQuota,omitempty"`
			SearchServiceMemoryQuota *string `tfsdk:"search_service_memory_quota" json:"searchServiceMemoryQuota,omitempty"`
		} `tfsdk:"cluster" json:"cluster,omitempty"`
		EnableOnlineVolumeExpansion *bool   `tfsdk:"enable_online_volume_expansion" json:"enableOnlineVolumeExpansion,omitempty"`
		EnablePreviewScaling        *bool   `tfsdk:"enable_preview_scaling" json:"enablePreviewScaling,omitempty"`
		EnvImagePrecedence          *bool   `tfsdk:"env_image_precedence" json:"envImagePrecedence,omitempty"`
		Hibernate                   *bool   `tfsdk:"hibernate" json:"hibernate,omitempty"`
		HibernationStrategy         *string `tfsdk:"hibernation_strategy" json:"hibernationStrategy,omitempty"`
		Image                       *string `tfsdk:"image" json:"image,omitempty"`
		Logging                     *struct {
			Audit *struct {
				DisabledEvents    *[]string `tfsdk:"disabled_events" json:"disabledEvents,omitempty"`
				DisabledUsers     *[]string `tfsdk:"disabled_users" json:"disabledUsers,omitempty"`
				Enabled           *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				GarbageCollection *struct {
					Sidecar *struct {
						Age       *string `tfsdk:"age" json:"age,omitempty"`
						Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
						Image     *string `tfsdk:"image" json:"image,omitempty"`
						Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
						Resources *struct {
							Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"sidecar" json:"sidecar,omitempty"`
				} `tfsdk:"garbage_collection" json:"garbageCollection,omitempty"`
				Rotation *struct {
					Interval *string `tfsdk:"interval" json:"interval,omitempty"`
					Size     *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"rotation" json:"rotation,omitempty"`
			} `tfsdk:"audit" json:"audit,omitempty"`
			LogRetentionCount *int64  `tfsdk:"log_retention_count" json:"logRetentionCount,omitempty"`
			LogRetentionTime  *string `tfsdk:"log_retention_time" json:"logRetentionTime,omitempty"`
			Server            *struct {
				ConfigurationName   *string `tfsdk:"configuration_name" json:"configurationName,omitempty"`
				Enabled             *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				ManageConfiguration *bool   `tfsdk:"manage_configuration" json:"manageConfiguration,omitempty"`
				Sidecar             *struct {
					ConfigurationMountPath *string `tfsdk:"configuration_mount_path" json:"configurationMountPath,omitempty"`
					Image                  *string `tfsdk:"image" json:"image,omitempty"`
					Resources              *struct {
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"sidecar" json:"sidecar,omitempty"`
			} `tfsdk:"server" json:"server,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		Monitoring *struct {
			Prometheus *struct {
				AuthorizationSecret *string `tfsdk:"authorization_secret" json:"authorizationSecret,omitempty"`
				Enabled             *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Image               *string `tfsdk:"image" json:"image,omitempty"`
				RefreshRate         *int64  `tfsdk:"refresh_rate" json:"refreshRate,omitempty"`
				Resources           *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		Networking *struct {
			AddressFamily               *string `tfsdk:"address_family" json:"addressFamily,omitempty"`
			AdminConsoleServiceTemplate *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
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
					SessionAffinity               *string   `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
					SessionAffinityConfig         *struct {
						ClientIP *struct {
							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
						} `tfsdk:"client_ip" json:"clientIP,omitempty"`
					} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"admin_console_service_template" json:"adminConsoleServiceTemplate,omitempty"`
			AdminConsoleServiceType *string   `tfsdk:"admin_console_service_type" json:"adminConsoleServiceType,omitempty"`
			AdminConsoleServices    *[]string `tfsdk:"admin_console_services" json:"adminConsoleServices,omitempty"`
			DisableUIOverHTTP       *bool     `tfsdk:"disable_ui_over_http" json:"disableUIOverHTTP,omitempty"`
			DisableUIOverHTTPS      *bool     `tfsdk:"disable_ui_over_https" json:"disableUIOverHTTPS,omitempty"`
			Dns                     *struct {
				Domain *string `tfsdk:"domain" json:"domain,omitempty"`
			} `tfsdk:"dns" json:"dns,omitempty"`
			ExposeAdminConsole            *bool `tfsdk:"expose_admin_console" json:"exposeAdminConsole,omitempty"`
			ExposedFeatureServiceTemplate *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
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
					SessionAffinity               *string   `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
					SessionAffinityConfig         *struct {
						ClientIP *struct {
							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
						} `tfsdk:"client_ip" json:"clientIP,omitempty"`
					} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"exposed_feature_service_template" json:"exposedFeatureServiceTemplate,omitempty"`
			ExposedFeatureServiceType   *string            `tfsdk:"exposed_feature_service_type" json:"exposedFeatureServiceType,omitempty"`
			ExposedFeatureTrafficPolicy *string            `tfsdk:"exposed_feature_traffic_policy" json:"exposedFeatureTrafficPolicy,omitempty"`
			ExposedFeatures             *[]string          `tfsdk:"exposed_features" json:"exposedFeatures,omitempty"`
			LoadBalancerSourceRanges    *[]string          `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			NetworkPlatform             *string            `tfsdk:"network_platform" json:"networkPlatform,omitempty"`
			ServiceAnnotations          *map[string]string `tfsdk:"service_annotations" json:"serviceAnnotations,omitempty"`
			Tls                         *struct {
				AllowPlainTextCertReload *bool     `tfsdk:"allow_plain_text_cert_reload" json:"allowPlainTextCertReload,omitempty"`
				CipherSuites             *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				ClientCertificatePaths   *[]struct {
					Delimiter *string `tfsdk:"delimiter" json:"delimiter,omitempty"`
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"client_certificate_paths" json:"clientCertificatePaths,omitempty"`
				ClientCertificatePolicy *string `tfsdk:"client_certificate_policy" json:"clientCertificatePolicy,omitempty"`
				NodeToNodeEncryption    *string `tfsdk:"node_to_node_encryption" json:"nodeToNodeEncryption,omitempty"`
				Passphrase              *struct {
					Rest *struct {
						AddressFamily *string            `tfsdk:"address_family" json:"addressFamily,omitempty"`
						Headers       *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
						Timeout       *int64             `tfsdk:"timeout" json:"timeout,omitempty"`
						Url           *string            `tfsdk:"url" json:"url,omitempty"`
						VerifyPeer    *bool              `tfsdk:"verify_peer" json:"verifyPeer,omitempty"`
					} `tfsdk:"rest" json:"rest,omitempty"`
					Script *struct {
						Secret *string `tfsdk:"secret" json:"secret,omitempty"`
					} `tfsdk:"script" json:"script,omitempty"`
				} `tfsdk:"passphrase" json:"passphrase,omitempty"`
				RootCAs      *[]string `tfsdk:"root_c_as" json:"rootCAs,omitempty"`
				SecretSource *struct {
					ClientSecretName *string `tfsdk:"client_secret_name" json:"clientSecretName,omitempty"`
					ServerSecretName *string `tfsdk:"server_secret_name" json:"serverSecretName,omitempty"`
				} `tfsdk:"secret_source" json:"secretSource,omitempty"`
				Static *struct {
					OperatorSecret *string `tfsdk:"operator_secret" json:"operatorSecret,omitempty"`
					ServerSecret   *string `tfsdk:"server_secret" json:"serverSecret,omitempty"`
				} `tfsdk:"static" json:"static,omitempty"`
				TlsMinimumVersion *string `tfsdk:"tls_minimum_version" json:"tlsMinimumVersion,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			WaitForAddressReachable      *string `tfsdk:"wait_for_address_reachable" json:"waitForAddressReachable,omitempty"`
			WaitForAddressReachableDelay *string `tfsdk:"wait_for_address_reachable_delay" json:"waitForAddressReachableDelay,omitempty"`
		} `tfsdk:"networking" json:"networking,omitempty"`
		Paused         *bool   `tfsdk:"paused" json:"paused,omitempty"`
		Platform       *string `tfsdk:"platform" json:"platform,omitempty"`
		RecoveryPolicy *string `tfsdk:"recovery_policy" json:"recoveryPolicy,omitempty"`
		RollingUpgrade *struct {
			MaxUpgradable        *int64  `tfsdk:"max_upgradable" json:"maxUpgradable,omitempty"`
			MaxUpgradablePercent *string `tfsdk:"max_upgradable_percent" json:"maxUpgradablePercent,omitempty"`
		} `tfsdk:"rolling_upgrade" json:"rollingUpgrade,omitempty"`
		Security *struct {
			AdminSecret *string `tfsdk:"admin_secret" json:"adminSecret,omitempty"`
			Ldap        *struct {
				AuthenticationEnabled *bool     `tfsdk:"authentication_enabled" json:"authenticationEnabled,omitempty"`
				AuthorizationEnabled  *bool     `tfsdk:"authorization_enabled" json:"authorizationEnabled,omitempty"`
				BindDN                *string   `tfsdk:"bind_dn" json:"bindDN,omitempty"`
				BindSecret            *string   `tfsdk:"bind_secret" json:"bindSecret,omitempty"`
				Cacert                *string   `tfsdk:"cacert" json:"cacert,omitempty"`
				CacheValueLifetime    *int64    `tfsdk:"cache_value_lifetime" json:"cacheValueLifetime,omitempty"`
				Encryption            *string   `tfsdk:"encryption" json:"encryption,omitempty"`
				GroupsQuery           *string   `tfsdk:"groups_query" json:"groupsQuery,omitempty"`
				Hosts                 *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				NestedGroupsEnabled   *bool     `tfsdk:"nested_groups_enabled" json:"nestedGroupsEnabled,omitempty"`
				NestedGroupsMaxDepth  *int64    `tfsdk:"nested_groups_max_depth" json:"nestedGroupsMaxDepth,omitempty"`
				Port                  *int64    `tfsdk:"port" json:"port,omitempty"`
				ServerCertValidation  *bool     `tfsdk:"server_cert_validation" json:"serverCertValidation,omitempty"`
				TlsSecret             *string   `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
				UserDNMapping         *struct {
					Query    *string `tfsdk:"query" json:"query,omitempty"`
					Template *string `tfsdk:"template" json:"template,omitempty"`
				} `tfsdk:"user_dn_mapping" json:"userDNMapping,omitempty"`
			} `tfsdk:"ldap" json:"ldap,omitempty"`
			Rbac *struct {
				Managed  *bool `tfsdk:"managed" json:"managed,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
			} `tfsdk:"rbac" json:"rbac,omitempty"`
			UiSessionTimeout *int64 `tfsdk:"ui_session_timeout" json:"uiSessionTimeout,omitempty"`
		} `tfsdk:"security" json:"security,omitempty"`
		SecurityContext *struct {
			FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
			FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
			RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
			RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
			RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
			SeLinuxOptions      *struct {
				Level *string `tfsdk:"level" json:"level,omitempty"`
				Role  *string `tfsdk:"role" json:"role,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				User  *string `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
			SeccompProfile *struct {
				LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
			SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
			Sysctls            *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"sysctls" json:"sysctls,omitempty"`
			WindowsOptions *struct {
				GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
				GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
				HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
				RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
			} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
		} `tfsdk:"security_context" json:"securityContext,omitempty"`
		ServerGroups *[]string `tfsdk:"server_groups" json:"serverGroups,omitempty"`
		Servers      *[]struct {
			AutoscaleEnabled *bool `tfsdk:"autoscale_enabled" json:"autoscaleEnabled,omitempty"`
			Env              *[]struct {
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
			EnvFrom *[]struct {
				ConfigMapRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
				SecretRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"env_from" json:"envFrom,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Pod  *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
					Affinity              *struct {
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
					AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
					DnsConfig                    *struct {
						Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
						Options     *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"options" json:"options,omitempty"`
						Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
					} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
					DnsPolicy          *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
					EnableServiceLinks *bool   `tfsdk:"enable_service_links" json:"enableServiceLinks,omitempty"`
					HostIPC            *bool   `tfsdk:"host_ipc" json:"hostIPC,omitempty"`
					HostNetwork        *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
					HostPID            *bool   `tfsdk:"host_pid" json:"hostPID,omitempty"`
					HostUsers          *bool   `tfsdk:"host_users" json:"hostUsers,omitempty"`
					ImagePullSecrets   *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
					NodeName     *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Os           *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"os" json:"os,omitempty"`
					Overhead                      *map[string]string `tfsdk:"overhead" json:"overhead,omitempty"`
					PreemptionPolicy              *string            `tfsdk:"preemption_policy" json:"preemptionPolicy,omitempty"`
					Priority                      *int64             `tfsdk:"priority" json:"priority,omitempty"`
					PriorityClassName             *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
					RuntimeClassName              *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
					SchedulerName                 *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
					ServiceAccount                *string            `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					ServiceAccountName            *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
					SetHostnameAsFQDN             *bool              `tfsdk:"set_hostname_as_fqdn" json:"setHostnameAsFQDN,omitempty"`
					ShareProcessNamespace         *bool              `tfsdk:"share_process_namespace" json:"shareProcessNamespace,omitempty"`
					TerminationGracePeriodSeconds *int64             `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					Tolerations                   *[]struct {
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
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"pod" json:"pod,omitempty"`
			Resources *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			ServerGroups *[]string `tfsdk:"server_groups" json:"serverGroups,omitempty"`
			Services     *[]string `tfsdk:"services" json:"services,omitempty"`
			Size         *int64    `tfsdk:"size" json:"size,omitempty"`
			VolumeMounts *struct {
				Analytics *[]string `tfsdk:"analytics" json:"analytics,omitempty"`
				Data      *string   `tfsdk:"data" json:"data,omitempty"`
				Default   *string   `tfsdk:"default" json:"default,omitempty"`
				Index     *string   `tfsdk:"index" json:"index,omitempty"`
				Logs      *string   `tfsdk:"logs" json:"logs,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"servers" json:"servers,omitempty"`
		SoftwareUpdateNotifications *bool   `tfsdk:"software_update_notifications" json:"softwareUpdateNotifications,omitempty"`
		UpgradeStrategy             *string `tfsdk:"upgrade_strategy" json:"upgradeStrategy,omitempty"`
		VolumeClaimTemplates        *[]struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				AccessModes   *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				DataSourceRef *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
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
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		Xdcr *struct {
			Managed        *bool `tfsdk:"managed" json:"managed,omitempty"`
			RemoteClusters *[]struct {
				AuthenticationSecret *string `tfsdk:"authentication_secret" json:"authenticationSecret,omitempty"`
				Hostname             *string `tfsdk:"hostname" json:"hostname,omitempty"`
				Name                 *string `tfsdk:"name" json:"name,omitempty"`
				Replications         *struct {
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"replications" json:"replications,omitempty"`
				Tls *struct {
					Secret *string `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
			} `tfsdk:"remote_clusters" json:"remoteClusters,omitempty"`
		} `tfsdk:"xdcr" json:"xdcr,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_cluster_v2"
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The CouchbaseCluster resource represents a Couchbase cluster.  It allows configuration of cluster topology, networking, storage and security options.",
		MarkdownDescription: "The CouchbaseCluster resource represents a Couchbase cluster.  It allows configuration of cluster topology, networking, storage and security options.",
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

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
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
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
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
				Description:         "ClusterSpec is the specification for a CouchbaseCluster resources, and allows the cluster to be customized.",
				MarkdownDescription: "ClusterSpec is the specification for a CouchbaseCluster resources, and allows the cluster to be customized.",
				Attributes: map[string]schema.Attribute{
					"anti_affinity": schema.BoolAttribute{
						Description:         "AntiAffinity forces the Operator to schedule different Couchbase server pods on different Kubernetes nodes.  Anti-affinity reduces the likelihood of unrecoverable failure in the event of a node issue.  Use of anti-affinity is highly recommended for production clusters.",
						MarkdownDescription: "AntiAffinity forces the Operator to schedule different Couchbase server pods on different Kubernetes nodes.  Anti-affinity reduces the likelihood of unrecoverable failure in the event of a node issue.  Use of anti-affinity is highly recommended for production clusters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_resource_allocation": schema.SingleNestedAttribute{
						Description:         "AutoResourceAllocation populates pod resource requests based on the services running on that pod.  When enabled, this feature will calculate the memory request as the total of service allocations defined in 'spec.cluster', plus an overhead defined by 'spec.autoResourceAllocation.overheadPercent'.Changing individual allocations for a service will cause a cluster upgrade as allocations are modified in the underlying pods.  This field also allows default pod CPU requests and limits to be applied. All resource allocations can be overridden by explicitly configuring them in the 'spec.servers.resources' field.",
						MarkdownDescription: "AutoResourceAllocation populates pod resource requests based on the services running on that pod.  When enabled, this feature will calculate the memory request as the total of service allocations defined in 'spec.cluster', plus an overhead defined by 'spec.autoResourceAllocation.overheadPercent'.Changing individual allocations for a service will cause a cluster upgrade as allocations are modified in the underlying pods.  This field also allows default pod CPU requests and limits to be applied. All resource allocations can be overridden by explicitly configuring them in the 'spec.servers.resources' field.",
						Attributes: map[string]schema.Attribute{
							"cpu_limits": schema.StringAttribute{
								Description:         "CPULimits automatically populates the CPU limits across all Couchbase server pods.  This field defaults to '4' CPUs.  Explicitly specifying the CPU limit for a particular server class will override this value.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "CPULimits automatically populates the CPU limits across all Couchbase server pods.  This field defaults to '4' CPUs.  Explicitly specifying the CPU limit for a particular server class will override this value.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"cpu_requests": schema.StringAttribute{
								Description:         "CPURequests automatically populates the CPU requests across all Couchbase server pods.  The default value of '2', is the minimum recommended number of CPUs required to run Couchbase Server.  Explicitly specifying the CPU request for a particular server class will override this value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "CPURequests automatically populates the CPU requests across all Couchbase server pods.  The default value of '2', is the minimum recommended number of CPUs required to run Couchbase Server.  Explicitly specifying the CPU request for a particular server class will override this value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines whether auto-resource allocation is enabled.",
								MarkdownDescription: "Enabled defines whether auto-resource allocation is enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"overhead_percent": schema.Int64Attribute{
								Description:         "OverheadPercent defines the amount of memory above that required for individual services on a pod.  For Couchbase Server this should be approximately 25%.",
								MarkdownDescription: "OverheadPercent defines the amount of memory above that required for individual services on a pod.  For Couchbase Server this should be approximately 25%.",
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

					"autoscale_stabilization_period": schema.StringAttribute{
						Description:         "AutoscaleStabilizationPeriod defines how long after a rebalance the corresponding HorizontalPodAutoscaler should remain in maintenance mode. During maintenance mode all autoscaling is disabled since every HorizontalPodAutoscaler associated with the cluster becomes inactive. Since certain metrics can be unpredictable when Couchbase is rebalancing or upgrading, setting a stabilization period helps to prevent scaling recommendations from the HorizontalPodAutoscaler for a provided period of time.  Values must be a valid Kubernetes duration of 0s or higher: https://golang.org/pkg/time/#ParseDuration A value of 0, puts the cluster in maintenance mode during rebalance but immediately exits this mode once the rebalance has completed. When undefined, the HPA is never put into maintenance mode during rebalance.",
						MarkdownDescription: "AutoscaleStabilizationPeriod defines how long after a rebalance the corresponding HorizontalPodAutoscaler should remain in maintenance mode. During maintenance mode all autoscaling is disabled since every HorizontalPodAutoscaler associated with the cluster becomes inactive. Since certain metrics can be unpredictable when Couchbase is rebalancing or upgrading, setting a stabilization period helps to prevent scaling recommendations from the HorizontalPodAutoscaler for a provided period of time.  Values must be a valid Kubernetes duration of 0s or higher: https://golang.org/pkg/time/#ParseDuration A value of 0, puts the cluster in maintenance mode during rebalance but immediately exits this mode once the rebalance has completed. When undefined, the HPA is never put into maintenance mode during rebalance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "Backup defines whether the Operator should manage automated backups, and how to lookup backup resources.",
						MarkdownDescription: "Backup defines whether the Operator should manage automated backups, and how to lookup backup resources.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations defines additional annotations to appear on the backup/restore pods.",
								MarkdownDescription: "Annotations defines additional annotations to appear on the backup/restore pods.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "The Backup Image to run on backup pods.",
								MarkdownDescription: "The Backup Image to run on backup pods.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "ImagePullSecrets allow you to use an image from private repositories and non-dockerhub ones.",
								MarkdownDescription: "ImagePullSecrets allow you to use an image from private repositories and non-dockerhub ones.",
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

							"labels": schema.MapAttribute{
								Description:         "Labels defines additional labels to appear on the backup/restore pods.",
								MarkdownDescription: "Labels defines additional labels to appear on the backup/restore pods.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"managed": schema.BoolAttribute{
								Description:         "Managed defines whether backups are managed by us or the clients.",
								MarkdownDescription: "Managed defines whether backups are managed by us or the clients.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector defines which nodes to constrain the pods that run any backup and restore operations to.",
								MarkdownDescription: "NodeSelector defines which nodes to constrain the pods that run any backup and restore operations to.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"object_endpoint": schema.SingleNestedAttribute{
								Description:         "Deprecated: by CouchbaseBackup.spec.objectStore.Endpoint ObjectEndpoint contains the configuration for connecting to a custom S3 compliant object store.",
								MarkdownDescription: "Deprecated: by CouchbaseBackup.spec.objectStore.Endpoint ObjectEndpoint contains the configuration for connecting to a custom S3 compliant object store.",
								Attributes: map[string]schema.Attribute{
									"secret": schema.StringAttribute{
										Description:         "The name of the secret, in this namespace, that contains the CA certificate for verification of a TLS endpoint The secret must have the key with the name 'tls.crt'",
										MarkdownDescription: "The name of the secret, in this namespace, that contains the CA certificate for verification of a TLS endpoint The secret must have the key with the name 'tls.crt'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "The host/address of the custom object endpoint.",
										MarkdownDescription: "The host/address of the custom object endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"use_virtual_path": schema.BoolAttribute{
										Description:         "UseVirtualPath will force the AWS SDK to use the new virtual style paths which are often required by S3 compatible object stores.",
										MarkdownDescription: "UseVirtualPath will force the AWS SDK to use the new virtual style paths which are often required by S3 compatible object stores.",
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
								Description:         "Resources is the resource requirements for the backup and restore containers.  Will be populated by defaults if not specified.",
								MarkdownDescription: "Resources is the resource requirements for the backup and restore containers.  Will be populated by defaults if not specified.",
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

							"s3_secret": schema.StringAttribute{
								Description:         "Deprecated: by CouchbaseBackup.spec.objectStore.secret S3Secret contains the key region and optionally access-key-id and secret-access-key for operating backups in S3. This field must be popluated when the 'spec.s3bucket' field is specified for a backup or restore resource.",
								MarkdownDescription: "Deprecated: by CouchbaseBackup.spec.objectStore.secret S3Secret contains the key region and optionally access-key-id and secret-access-key for operating backups in S3. This field must be popluated when the 'spec.s3bucket' field is specified for a backup or restore resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector allows CouchbaseBackup and CouchbaseBackupRestore resources to be filtered based on labels.",
								MarkdownDescription: "Selector allows CouchbaseBackup and CouchbaseBackupRestore resources to be filtered based on labels.",
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

							"service_account_name": schema.StringAttribute{
								Description:         "The Service Account to run backup (and restore) pods under. Without this backup pods will not be able to update status.",
								MarkdownDescription: "The Service Account to run backup (and restore) pods under. Without this backup pods will not be able to update status.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations specifies all backup and restore pod tolerations.",
								MarkdownDescription: "Tolerations specifies all backup and restore pod tolerations.",
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

							"use_iam_role": schema.BoolAttribute{
								Description:         "Deprecated: by CouchbaseBackup.spec.objectStore.useIAM UseIAMRole enables backup to fetch EC2 instance metadata. This allows the AWS SDK to use the EC2's IAM Role for S3 access. UseIAMRole will ignore credentials in s3Secret.",
								MarkdownDescription: "Deprecated: by CouchbaseBackup.spec.objectStore.useIAM UseIAMRole enables backup to fetch EC2 instance metadata. This allows the AWS SDK to use the EC2's IAM Role for S3 access. UseIAMRole will ignore credentials in s3Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"buckets": schema.SingleNestedAttribute{
						Description:         "Buckets defines whether the Operator should manage buckets, and how to lookup bucket resources.",
						MarkdownDescription: "Buckets defines whether the Operator should manage buckets, and how to lookup bucket resources.",
						Attributes: map[string]schema.Attribute{
							"managed": schema.BoolAttribute{
								Description:         "Managed defines whether buckets are managed by the Operator (true), or user managed (false). When Operator managed, all buckets must be defined with either CouchbaseBucket, CouchbaseEphemeralBucket or CouchbaseMemcachedBucket resources.  Manual addition of buckets will be reverted by the Operator.  When user managed, the Operator will not interrogate buckets at all.  This field defaults to false.",
								MarkdownDescription: "Managed defines whether buckets are managed by the Operator (true), or user managed (false). When Operator managed, all buckets must be defined with either CouchbaseBucket, CouchbaseEphemeralBucket or CouchbaseMemcachedBucket resources.  Manual addition of buckets will be reverted by the Operator.  When user managed, the Operator will not interrogate buckets at all.  This field defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector is a label selector used to list buckets in the namespace that are managed by the Operator.",
								MarkdownDescription: "Selector is a label selector used to list buckets in the namespace that are managed by the Operator.",
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

							"synchronize": schema.BoolAttribute{
								Description:         "Synchronize allows unmanaged buckets, scopes, and collections to be synchronized as Kubernetes resources by the Operator.  This feature is intended for development only and should not be used for production workloads.  The synchronization workflow starts with 'spec.buckets.managed' being set to false, the user can manually create buckets, scopes, and collections using the Couchbase UI, or other tooling.  When you wish to commit to Kubernetes resources, you must specify a unique label selector in the 'spec.buckets.selector' field, and this field is set to true.  The Operator will create Kubernetes resources for you, and upon completion set the cluster's 'Synchronized' status condition.  You may then safely set 'spec.buckets.managed' to true and the Operator will manage these resources as per usual.  To update an already managed data topology, you must first set it to unmanaged, make any changes, and delete any old resources, then follow the standard synchronization workflow.  The Operator can not, and will not, ever delete, or make modifications to resource specifications that are intended to be user managed, or managed by a life cycle management tool. These actions must be instigated by an end user.  For a more complete experience, refer to the documentation for the 'cao save' and 'cao restore' CLI commands.",
								MarkdownDescription: "Synchronize allows unmanaged buckets, scopes, and collections to be synchronized as Kubernetes resources by the Operator.  This feature is intended for development only and should not be used for production workloads.  The synchronization workflow starts with 'spec.buckets.managed' being set to false, the user can manually create buckets, scopes, and collections using the Couchbase UI, or other tooling.  When you wish to commit to Kubernetes resources, you must specify a unique label selector in the 'spec.buckets.selector' field, and this field is set to true.  The Operator will create Kubernetes resources for you, and upon completion set the cluster's 'Synchronized' status condition.  You may then safely set 'spec.buckets.managed' to true and the Operator will manage these resources as per usual.  To update an already managed data topology, you must first set it to unmanaged, make any changes, and delete any old resources, then follow the standard synchronization workflow.  The Operator can not, and will not, ever delete, or make modifications to resource specifications that are intended to be user managed, or managed by a life cycle management tool. These actions must be instigated by an end user.  For a more complete experience, refer to the documentation for the 'cao save' and 'cao restore' CLI commands.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster": schema.SingleNestedAttribute{
						Description:         "ClusterSettings define Couchbase cluster-wide settings such as memory allocation, failover characteristics and index settings.",
						MarkdownDescription: "ClusterSettings define Couchbase cluster-wide settings such as memory allocation, failover characteristics and index settings.",
						Attributes: map[string]schema.Attribute{
							"analytics_service_memory_quota": schema.StringAttribute{
								Description:         "AnalyticsServiceMemQuota is the amount of memory that should be allocated to the analytics service. This value is per-pod, and only applicable to pods belonging to server classes running the analytics service.  This field must be a quantity greater than or equal to 1Gi.  This field defaults to 1Gi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "AnalyticsServiceMemQuota is the amount of memory that should be allocated to the analytics service. This value is per-pod, and only applicable to pods belonging to server classes running the analytics service.  This field must be a quantity greater than or equal to 1Gi.  This field defaults to 1Gi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"auto_compaction": schema.SingleNestedAttribute{
								Description:         "AutoCompaction allows the configuration of auto-compaction, including on what conditions disk space is reclaimed and when it is allowed to run.",
								MarkdownDescription: "AutoCompaction allows the configuration of auto-compaction, including on what conditions disk space is reclaimed and when it is allowed to run.",
								Attributes: map[string]schema.Attribute{
									"database_fragmentation_threshold": schema.SingleNestedAttribute{
										Description:         "DatabaseFragmentationThreshold defines triggers for when database compaction should start.",
										MarkdownDescription: "DatabaseFragmentationThreshold defines triggers for when database compaction should start.",
										Attributes: map[string]schema.Attribute{
											"percent": schema.Int64Attribute{
												Description:         "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",
												MarkdownDescription: "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(2),
													int64validator.AtMost(100),
												},
											},

											"size": schema.StringAttribute{
												Description:         "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												MarkdownDescription: "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"parallel_compaction": schema.BoolAttribute{
										Description:         "ParallelCompaction controls whether database and view compactions can happen in parallel.",
										MarkdownDescription: "ParallelCompaction controls whether database and view compactions can happen in parallel.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"time_window": schema.SingleNestedAttribute{
										Description:         "TimeWindow allows restriction of when compaction can occur.",
										MarkdownDescription: "TimeWindow allows restriction of when compaction can occur.",
										Attributes: map[string]schema.Attribute{
											"abort_compaction_outside_window": schema.BoolAttribute{
												Description:         "AbortCompactionOutsideWindow stops compaction processes when the process moves outside the window.",
												MarkdownDescription: "AbortCompactionOutsideWindow stops compaction processes when the process moves outside the window.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"end": schema.StringAttribute{
												Description:         "End is a wallclock time, in the form HH:MM, when a compaction should stop.",
												MarkdownDescription: "End is a wallclock time, in the form HH:MM, when a compaction should stop.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(2[0-3]|[01]?[0-9]):([0-5]?[0-9])$`), ""),
												},
											},

											"start": schema.StringAttribute{
												Description:         "Start is a wallclock time, in the form HH:MM, when a compaction is permitted to start.",
												MarkdownDescription: "Start is a wallclock time, in the form HH:MM, when a compaction is permitted to start.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(2[0-3]|[01]?[0-9]):([0-5]?[0-9])$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tombstone_purge_interval": schema.StringAttribute{
										Description:         "TombstonePurgeInterval controls how long to wait before purging tombstones. This field must be in the range 1h-1440h, defaulting to 72h. More info:  https://golang.org/pkg/time/#ParseDuration",
										MarkdownDescription: "TombstonePurgeInterval controls how long to wait before purging tombstones. This field must be in the range 1h-1440h, defaulting to 72h. More info:  https://golang.org/pkg/time/#ParseDuration",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"view_fragmentation_threshold": schema.SingleNestedAttribute{
										Description:         "ViewFragmentationThreshold defines triggers for when view compaction should start.",
										MarkdownDescription: "ViewFragmentationThreshold defines triggers for when view compaction should start.",
										Attributes: map[string]schema.Attribute{
											"percent": schema.Int64Attribute{
												Description:         "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",
												MarkdownDescription: "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(2),
													int64validator.AtMost(100),
												},
											},

											"size": schema.StringAttribute{
												Description:         "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												MarkdownDescription: "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
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

							"auto_failover_max_count": schema.Int64Attribute{
								Description:         "AutoFailoverMaxCount is the maximum number of automatic failovers Couchbase server will allow before not allowing any more.  This field must be between 1-3 for server versions prior to 7.1.0 default is 3.",
								MarkdownDescription: "AutoFailoverMaxCount is the maximum number of automatic failovers Couchbase server will allow before not allowing any more.  This field must be between 1-3 for server versions prior to 7.1.0 default is 3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"auto_failover_on_data_disk_issues": schema.BoolAttribute{
								Description:         "AutoFailoverOnDataDiskIssues defines whether Couchbase server should failover a pod if a disk issue was detected.",
								MarkdownDescription: "AutoFailoverOnDataDiskIssues defines whether Couchbase server should failover a pod if a disk issue was detected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auto_failover_on_data_disk_issues_time_period": schema.StringAttribute{
								Description:         "AutoFailoverOnDataDiskIssuesTimePeriod defines how long to wait for transient errors before failing over a faulty disk.  This field must be in the range 5-3600s, defaulting to 120s.  More info:  https://golang.org/pkg/time/#ParseDuration",
								MarkdownDescription: "AutoFailoverOnDataDiskIssuesTimePeriod defines how long to wait for transient errors before failing over a faulty disk.  This field must be in the range 5-3600s, defaulting to 120s.  More info:  https://golang.org/pkg/time/#ParseDuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auto_failover_server_group": schema.BoolAttribute{
								Description:         "AutoFailoverServerGroup whether to enable failing over a server group. This field is ignored in server versions 7.1+ as it has been removed from the Couchbase API",
								MarkdownDescription: "AutoFailoverServerGroup whether to enable failing over a server group. This field is ignored in server versions 7.1+ as it has been removed from the Couchbase API",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auto_failover_timeout": schema.StringAttribute{
								Description:         "AutoFailoverTimeout defines how long Couchbase server will wait between a pod being witnessed as down, until when it will failover the pod.  Couchbase server will only failover pods if it deems it safe to do so, and not result in data loss.  This field must be in the range 5-3600s, defaulting to 120s. More info:  https://golang.org/pkg/time/#ParseDuration",
								MarkdownDescription: "AutoFailoverTimeout defines how long Couchbase server will wait between a pod being witnessed as down, until when it will failover the pod.  Couchbase server will only failover pods if it deems it safe to do so, and not result in data loss.  This field must be in the range 5-3600s, defaulting to 120s. More info:  https://golang.org/pkg/time/#ParseDuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_name": schema.StringAttribute{
								Description:         "ClusterName defines the name of the cluster, as displayed in the Couchbase UI. By default, the cluster name is that specified in the CouchbaseCluster resource's metadata.",
								MarkdownDescription: "ClusterName defines the name of the cluster, as displayed in the Couchbase UI. By default, the cluster name is that specified in the CouchbaseCluster resource's metadata.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data": schema.SingleNestedAttribute{
								Description:         "Data allows the data service to be configured.",
								MarkdownDescription: "Data allows the data service to be configured.",
								Attributes: map[string]schema.Attribute{
									"reader_threads": schema.Int64Attribute{
										Description:         "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use.  If not specified, this defaults to the default value set by Couchbase Server.",
										MarkdownDescription: "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use.  If not specified, this defaults to the default value set by Couchbase Server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(4),
											int64validator.AtMost(64),
										},
									},

									"writer_threads": schema.Int64Attribute{
										Description:         "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This setting is especially relevant when using 'durable writes', increasing this field will have a large impact on performance.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use. If not specified, this defaults to the default value set by Couchbase Server.",
										MarkdownDescription: "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This setting is especially relevant when using 'durable writes', increasing this field will have a large impact on performance.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use. If not specified, this defaults to the default value set by Couchbase Server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(4),
											int64validator.AtMost(64),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_service_memory_quota": schema.StringAttribute{
								Description:         "DataServiceMemQuota is the amount of memory that should be allocated to the data service. This value is per-pod, and only applicable to pods belonging to server classes running the data service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "DataServiceMemQuota is the amount of memory that should be allocated to the data service. This value is per-pod, and only applicable to pods belonging to server classes running the data service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"eventing_service_memory_quota": schema.StringAttribute{
								Description:         "EventingServiceMemQuota is the amount of memory that should be allocated to the eventing service. This value is per-pod, and only applicable to pods belonging to server classes running the eventing service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "EventingServiceMemQuota is the amount of memory that should be allocated to the eventing service. This value is per-pod, and only applicable to pods belonging to server classes running the eventing service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"index_service_memory_quota": schema.StringAttribute{
								Description:         "IndexServiceMemQuota is the amount of memory that should be allocated to the index service. This value is per-pod, and only applicable to pods belonging to server classes running the index service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "IndexServiceMemQuota is the amount of memory that should be allocated to the index service. This value is per-pod, and only applicable to pods belonging to server classes running the index service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"index_storage_setting": schema.StringAttribute{
								Description:         "DEPRECATED - by indexer. The index storage mode to use for secondary indexing.  This field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.  This field is immutable and cannot be changed unless there are no server classes running the index service in the cluster.",
								MarkdownDescription: "DEPRECATED - by indexer. The index storage mode to use for secondary indexing.  This field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.  This field is immutable and cannot be changed unless there are no server classes running the index service in the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("memory_optimized", "plasma"),
								},
							},

							"indexer": schema.SingleNestedAttribute{
								Description:         "Indexer allows the indexer to be configured.",
								MarkdownDescription: "Indexer allows the indexer to be configured.",
								Attributes: map[string]schema.Attribute{
									"log_level": schema.StringAttribute{
										Description:         "LogLevel controls the verbosity of indexer logs.  This field must be one of 'silent', 'fatal', 'error', 'warn', 'info', 'verbose', 'timing', 'debug' or 'trace', defaulting to 'info'.",
										MarkdownDescription: "LogLevel controls the verbosity of indexer logs.  This field must be one of 'silent', 'fatal', 'error', 'warn', 'info', 'verbose', 'timing', 'debug' or 'trace', defaulting to 'info'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("silent", "fatal", "error", "warn", "info", "verbose", "timing", "debug", "trace"),
										},
									},

									"max_rollback_points": schema.Int64Attribute{
										Description:         "MaxRollbackPoints controls the number of checkpoints that can be rolled back to.  The default is 2, with a minimum of 1.",
										MarkdownDescription: "MaxRollbackPoints controls the number of checkpoints that can be rolled back to.  The default is 2, with a minimum of 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"memory_snapshot_interval": schema.StringAttribute{
										Description:         "MemorySnapshotInterval controls when memory indexes should be snapshotted. This defaults to 200ms, and must be greater than or equal to 1ms.",
										MarkdownDescription: "MemorySnapshotInterval controls when memory indexes should be snapshotted. This defaults to 200ms, and must be greater than or equal to 1ms.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"num_replica": schema.Int64Attribute{
										Description:         "NumberOfReplica specifies number of secondary index replicas to be created by the Index Service whenever CREATE INDEX is invoked, which ensures high availability and high performance. Note, if nodes and num_replica are both specified in the WITH clause, the specified number of nodes must be one greater than num_replica This defaults to 0, which means no index replicas to be created by default. Minimum must be 0.",
										MarkdownDescription: "NumberOfReplica specifies number of secondary index replicas to be created by the Index Service whenever CREATE INDEX is invoked, which ensures high availability and high performance. Note, if nodes and num_replica are both specified in the WITH clause, the specified number of nodes must be one greater than num_replica This defaults to 0, which means no index replicas to be created by default. Minimum must be 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"redistribute_indexes": schema.BoolAttribute{
										Description:         "RedistributeIndexes when true, Couchbase Server redistributes indexes when rebalance occurs, in order to optimize performance. If false (the default), such redistribution does not occur.",
										MarkdownDescription: "RedistributeIndexes when true, Couchbase Server redistributes indexes when rebalance occurs, in order to optimize performance. If false (the default), such redistribution does not occur.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stable_snapshot_interval": schema.StringAttribute{
										Description:         "StableSnapshotInterval controls when disk indexes should be snapshotted. This defaults to 5s, and must be greater than or equal to 1ms.",
										MarkdownDescription: "StableSnapshotInterval controls when disk indexes should be snapshotted. This defaults to 5s, and must be greater than or equal to 1ms.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_mode": schema.StringAttribute{
										Description:         "StorageMode controls the underlying storage engine for indexes.  Once set it can only be modified if there are no nodes in the cluster running the index service.  The field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.",
										MarkdownDescription: "StorageMode controls the underlying storage engine for indexes.  Once set it can only be modified if there are no nodes in the cluster running the index service.  The field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("memory_optimized", "plasma"),
										},
									},

									"threads": schema.Int64Attribute{
										Description:         "Threads controls the number of processor threads to use for indexing. A value of 0 means 1 per CPU.  This attribute must be greater than or equal to 0, defaulting to 0.",
										MarkdownDescription: "Threads controls the number of processor threads to use for indexing. A value of 0 means 1 per CPU.  This attribute must be greater than or equal to 0, defaulting to 0.",
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

							"query": schema.SingleNestedAttribute{
								Description:         "Query allows the query service to be configured.",
								MarkdownDescription: "Query allows the query service to be configured.",
								Attributes: map[string]schema.Attribute{
									"backfill_enabled": schema.BoolAttribute{
										Description:         "BackfillEnabled allows the query service to backfill.",
										MarkdownDescription: "BackfillEnabled allows the query service to backfill.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"temporary_space": schema.StringAttribute{
										Description:         "TemporarySpace allows the temporary storage used by the query service backfill, per-pod, to be modified.  This field requires 'backfillEnabled' to be set to true in order to have any effect. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
										MarkdownDescription: "TemporarySpace allows the temporary storage used by the query service backfill, per-pod, to be modified.  This field requires 'backfillEnabled' to be set to true in order to have any effect. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
										},
									},

									"temporary_space_unlimited": schema.BoolAttribute{
										Description:         "TemporarySpaceUnlimited allows the temporary storage used by the query service backfill, per-pod, to be unconstrained.  This field requires 'backfillEnabled' to be set to true in order to have any effect. This field overrides 'temporarySpace'.",
										MarkdownDescription: "TemporarySpaceUnlimited allows the temporary storage used by the query service backfill, per-pod, to be unconstrained.  This field requires 'backfillEnabled' to be set to true in order to have any effect. This field overrides 'temporarySpace'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"query_service_memory_quota": schema.StringAttribute{
								Description:         "QueryServiceMemQuota is a dummy field.  By default, Couchbase server provides no memory resource constraints for the query service, so this has no effect on Couchbase server.  It is, however, used when the spec.autoResourceAllocation feature is enabled, and is used to define the amount of memory reserved by the query service for use with Kubernetes resource scheduling. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "QueryServiceMemQuota is a dummy field.  By default, Couchbase server provides no memory resource constraints for the query service, so this has no effect on Couchbase server.  It is, however, used when the spec.autoResourceAllocation feature is enabled, and is used to define the amount of memory reserved by the query service for use with Kubernetes resource scheduling. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"search_service_memory_quota": schema.StringAttribute{
								Description:         "SearchServiceMemQuota is the amount of memory that should be allocated to the search service. This value is per-pod, and only applicable to pods belonging to server classes running the search service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "SearchServiceMemQuota is the amount of memory that should be allocated to the search service. This value is per-pod, and only applicable to pods belonging to server classes running the search service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_online_volume_expansion": schema.BoolAttribute{
						Description:         "EnableOnlineVolumeExpansion enables online expansion of Persistent Volumes. You can only expand a PVC if its storage class's 'allowVolumeExpansion' field is set to true. Additionally, Kubernetes feature 'ExpandInUsePersistentVolumes' must be enabled in order to expand the volumes which are actively bound to Pods. Volumes can only be expanded and not reduced to a smaller size. See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#resizing-an-in-use-persistentvolumeclaim  If 'EnableOnlineVolumeExpansion' is enabled for use within an environment that does not actually support online volume and file system expansion then the cluster will fallback to rolling upgrade procedure to create a new set of Pods for use with resized Volumes. More info:  https://kubernetes.io/docs/concepts/storage/persistent-volumes/#expanding-persistent-volumes-claims",
						MarkdownDescription: "EnableOnlineVolumeExpansion enables online expansion of Persistent Volumes. You can only expand a PVC if its storage class's 'allowVolumeExpansion' field is set to true. Additionally, Kubernetes feature 'ExpandInUsePersistentVolumes' must be enabled in order to expand the volumes which are actively bound to Pods. Volumes can only be expanded and not reduced to a smaller size. See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#resizing-an-in-use-persistentvolumeclaim  If 'EnableOnlineVolumeExpansion' is enabled for use within an environment that does not actually support online volume and file system expansion then the cluster will fallback to rolling upgrade procedure to create a new set of Pods for use with resized Volumes. More info:  https://kubernetes.io/docs/concepts/storage/persistent-volumes/#expanding-persistent-volumes-claims",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_preview_scaling": schema.BoolAttribute{
						Description:         "DEPRECATED - This option only exists for backwards compatibility and no longer restricts autoscaling to ephemeral services. EnablePreviewScaling enables autoscaling for stateful services and buckets.",
						MarkdownDescription: "DEPRECATED - This option only exists for backwards compatibility and no longer restricts autoscaling to ephemeral services. EnablePreviewScaling enables autoscaling for stateful services and buckets.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env_image_precedence": schema.BoolAttribute{
						Description:         "EnvImagePrecedence gives precedence over the default container image name in 'spec.Image' to an image name provided through Operator environment variables. For more info on using Operator environment variables: https://docs.couchbase.com/operator/current/reference-operator-configuration.html",
						MarkdownDescription: "EnvImagePrecedence gives precedence over the default container image name in 'spec.Image' to an image name provided through Operator environment variables. For more info on using Operator environment variables: https://docs.couchbase.com/operator/current/reference-operator-configuration.html",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hibernate": schema.BoolAttribute{
						Description:         "Hibernate is whether to hibernate the cluster.",
						MarkdownDescription: "Hibernate is whether to hibernate the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hibernation_strategy": schema.StringAttribute{
						Description:         "HibernationStrategy defines how to hibernate the cluster.  When Immediate the Operator will immediately delete all pods and take no further action until the hibernate field is set to false.",
						MarkdownDescription: "HibernationStrategy defines how to hibernate the cluster.  When Immediate the Operator will immediately delete all pods and take no further action until the hibernate field is set to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Immediate"),
						},
					},

					"image": schema.StringAttribute{
						Description:         "Image is the container image name that will be used to launch Couchbase server instances.  Updating this field will cause an automatic upgrade of the cluster.",
						MarkdownDescription: "Image is the container image name that will be used to launch Couchbase server instances.  Updating this field will cause an automatic upgrade of the cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(.*?(:\d+)?/)?.*?/.*?(:.*?\d+\.\d+\.\d+.*|@sha256:[0-9a-f]{64})$`), ""),
						},
					},

					"logging": schema.SingleNestedAttribute{
						Description:         "Logging defines Operator logging options.",
						MarkdownDescription: "Logging defines Operator logging options.",
						Attributes: map[string]schema.Attribute{
							"audit": schema.SingleNestedAttribute{
								Description:         "Used to manage the audit configuration directly",
								MarkdownDescription: "Used to manage the audit configuration directly",
								Attributes: map[string]schema.Attribute{
									"disabled_events": schema.ListAttribute{
										Description:         "The list of event ids to disable for auditing purposes. This is passed to the REST API with no verification by the operator. Refer to the documentation for details: https://docs.couchbase.com/server/current/audit-event-reference/audit-event-reference.html",
										MarkdownDescription: "The list of event ids to disable for auditing purposes. This is passed to the REST API with no verification by the operator. Refer to the documentation for details: https://docs.couchbase.com/server/current/audit-event-reference/audit-event-reference.html",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disabled_users": schema.ListAttribute{
										Description:         "The list of users to ignore for auditing purposes. This is passed to the REST API with minimal validation it meets an acceptable regex pattern. Refer to the documentation for full details on how to configure this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html#ignoring-events-by-user",
										MarkdownDescription: "The list of users to ignore for auditing purposes. This is passed to the REST API with minimal validation it meets an acceptable regex pattern. Refer to the documentation for full details on how to configure this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html#ignoring-events-by-user",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled is a boolean that enables the audit capabilities.",
										MarkdownDescription: "Enabled is a boolean that enables the audit capabilities.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"garbage_collection": schema.SingleNestedAttribute{
										Description:         "Handle all optional garbage collection (GC) configuration for the audit functionality. This is not part of the audit REST API, it is intended to handle GC automatically for the audit logs. By default the Couchbase Server rotates the audit logs but does not clean up the rotated logs. This is left as an operation for the cluster administrator to manage, the operator allows for us to automate this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",
										MarkdownDescription: "Handle all optional garbage collection (GC) configuration for the audit functionality. This is not part of the audit REST API, it is intended to handle GC automatically for the audit logs. By default the Couchbase Server rotates the audit logs but does not clean up the rotated logs. This is left as an operation for the cluster administrator to manage, the operator allows for us to automate this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",
										Attributes: map[string]schema.Attribute{
											"sidecar": schema.SingleNestedAttribute{
												Description:         "Provide the sidecar configuration required (if so desired) to automatically clean up audit logs.",
												MarkdownDescription: "Provide the sidecar configuration required (if so desired) to automatically clean up audit logs.",
												Attributes: map[string]schema.Attribute{
													"age": schema.StringAttribute{
														Description:         "The minimum age of rotated log files to remove, defaults to one hour.",
														MarkdownDescription: "The minimum age of rotated log files to remove, defaults to one hour.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enabled": schema.BoolAttribute{
														Description:         "Enable this sidecar by setting to true, defaults to being disabled.",
														MarkdownDescription: "Enable this sidecar by setting to true, defaults to being disabled.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image": schema.StringAttribute{
														Description:         "Image is the image to be used to run the audit sidecar helper. No validation is carried out as this can be any arbitrary repo and tag.",
														MarkdownDescription: "Image is the image to be used to run the audit sidecar helper. No validation is carried out as this can be any arbitrary repo and tag.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"interval": schema.StringAttribute{
														Description:         "The interval at which to check for rotated log files to remove, defaults to 20 minutes.",
														MarkdownDescription: "The interval at which to check for rotated log files to remove, defaults to 20 minutes.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Resources is the resource requirements for the cleanup container. Will be populated by Kubernetes defaults if not specified.",
														MarkdownDescription: "Resources is the resource requirements for the cleanup container. Will be populated by Kubernetes defaults if not specified.",
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

									"rotation": schema.SingleNestedAttribute{
										Description:         "The interval to optionally rotate the audit log. This is passed to the REST API, see here for details: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",
										MarkdownDescription: "The interval to optionally rotate the audit log. This is passed to the REST API, see here for details: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",
										Attributes: map[string]schema.Attribute{
											"interval": schema.StringAttribute{
												Description:         "The interval at which to rotate log files, defaults to 15 minutes.",
												MarkdownDescription: "The interval at which to rotate log files, defaults to 15 minutes.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"size": schema.StringAttribute{
												Description:         "Size allows the specification of a rotation size for the log, defaults to 20Mi. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												MarkdownDescription: "Size allows the specification of a rotation size for the log, defaults to 20Mi. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
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

							"log_retention_count": schema.Int64Attribute{
								Description:         "LogRetentionCount gives the number of persistent log PVCs to keep.",
								MarkdownDescription: "LogRetentionCount gives the number of persistent log PVCs to keep.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"log_retention_time": schema.StringAttribute{
								Description:         "LogRetentionTime gives the time to keep persistent log PVCs alive for.",
								MarkdownDescription: "LogRetentionTime gives the time to keep persistent log PVCs alive for.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|us|ms|s|m|h)$`), ""),
								},
							},

							"server": schema.SingleNestedAttribute{
								Description:         "Specification of all logging configuration required to manage the sidecar containers in each pod.",
								MarkdownDescription: "Specification of all logging configuration required to manage the sidecar containers in each pod.",
								Attributes: map[string]schema.Attribute{
									"configuration_name": schema.StringAttribute{
										Description:         "ConfigurationName is the name of the Secret to use holding the logging configuration in the namespace. A Secret is used to ensure we can safely store credentials but this can be populated from plaintext if acceptable too. If it does not exist then one will be created with defaults in the namespace so it can be easily updated whilst running. Note that if running multiple clusters in the same kubernetes namespace then you should use a separate Secret for each, otherwise the first cluster will take ownership (if created) and the Secret will be cleaned up when that cluster is removed. If running clusters in separate namespaces then they will be separate Secrets anyway.",
										MarkdownDescription: "ConfigurationName is the name of the Secret to use holding the logging configuration in the namespace. A Secret is used to ensure we can safely store credentials but this can be populated from plaintext if acceptable too. If it does not exist then one will be created with defaults in the namespace so it can be easily updated whilst running. Note that if running multiple clusters in the same kubernetes namespace then you should use a separate Secret for each, otherwise the first cluster will take ownership (if created) and the Secret will be cleaned up when that cluster is removed. If running clusters in separate namespaces then they will be separate Secrets anyway.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled is a boolean that enables the logging sidecar container.",
										MarkdownDescription: "Enabled is a boolean that enables the logging sidecar container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"manage_configuration": schema.BoolAttribute{
										Description:         "A boolean which indicates whether the operator should manage the configuration or not. If omitted then this defaults to true which means the operator will attempt to reconcile it to default values. To use a custom configuration make sure to set this to false. Note that the ownership of any Secret is not changed so if a Secret is created externally it can be updated by the operator but it's ownership stays the same so it will be cleaned up when it's owner is.",
										MarkdownDescription: "A boolean which indicates whether the operator should manage the configuration or not. If omitted then this defaults to true which means the operator will attempt to reconcile it to default values. To use a custom configuration make sure to set this to false. Note that the ownership of any Secret is not changed so if a Secret is created externally it can be updated by the operator but it's ownership stays the same so it will be cleaned up when it's owner is.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sidecar": schema.SingleNestedAttribute{
										Description:         "Any specific logging sidecar container configuration.",
										MarkdownDescription: "Any specific logging sidecar container configuration.",
										Attributes: map[string]schema.Attribute{
											"configuration_mount_path": schema.StringAttribute{
												Description:         "ConfigurationMountPath is the location to mount the ConfigurationName Secret into the image. If another log shipping image is used that needs a different mount then modify this. Note that the configuration file must be called 'fluent-bit.conf' at the root of this path, there is no provision for overriding the name of the config file passed as the COUCHBASE_LOGS_CONFIG_FILE environment variable.",
												MarkdownDescription: "ConfigurationMountPath is the location to mount the ConfigurationName Secret into the image. If another log shipping image is used that needs a different mount then modify this. Note that the configuration file must be called 'fluent-bit.conf' at the root of this path, there is no provision for overriding the name of the config file passed as the COUCHBASE_LOGS_CONFIG_FILE environment variable.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image": schema.StringAttribute{
												Description:         "Image is the image to be used to deal with logging as a sidecar. No validation is carried out as this can be any arbitrary repo and tag. It will default to the latest supported version of Fluent Bit.",
												MarkdownDescription: "Image is the image to be used to deal with logging as a sidecar. No validation is carried out as this can be any arbitrary repo and tag. It will default to the latest supported version of Fluent Bit.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resources is the resource requirements for the sidecar container. Will be populated by Kubernetes defaults if not specified.",
												MarkdownDescription: "Resources is the resource requirements for the sidecar container. Will be populated by Kubernetes defaults if not specified.",
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

					"monitoring": schema.SingleNestedAttribute{
						Description:         "Monitoring defines any Operator managed integration into 3rd party monitoring infrastructure.",
						MarkdownDescription: "Monitoring defines any Operator managed integration into 3rd party monitoring infrastructure.",
						Attributes: map[string]schema.Attribute{
							"prometheus": schema.SingleNestedAttribute{
								Description:         "Prometheus provides integration with Prometheus monitoring.",
								MarkdownDescription: "Prometheus provides integration with Prometheus monitoring.",
								Attributes: map[string]schema.Attribute{
									"authorization_secret": schema.StringAttribute{
										Description:         "AuthorizationSecret is the name of a Kubernetes secret that contains a bearer token to authorize GET requests to the metrics endpoint",
										MarkdownDescription: "AuthorizationSecret is the name of a Kubernetes secret that contains a bearer token to authorize GET requests to the metrics endpoint",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled is a boolean that enables/disables the metrics sidecar container.",
										MarkdownDescription: "Enabled is a boolean that enables/disables the metrics sidecar container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "Image is the metrics image to be used to collect metrics. No validation is carried out as this can be any arbitrary repo and tag.",
										MarkdownDescription: "Image is the metrics image to be used to collect metrics. No validation is carried out as this can be any arbitrary repo and tag.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"refresh_rate": schema.Int64Attribute{
										Description:         "RefreshRate is the frequency in which cached statistics are updated in seconds. Shorter intervals will add additional resource overhead to clusters running Couchbase Server 7.0+ Default is 60 seconds, Maximum value is 600 seconds, and minimum value is 1 second.",
										MarkdownDescription: "RefreshRate is the frequency in which cached statistics are updated in seconds. Shorter intervals will add additional resource overhead to clusters running Couchbase Server 7.0+ Default is 60 seconds, Maximum value is 600 seconds, and minimum value is 1 second.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(600),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources is the resource requirements for the metrics container. Will be populated by Kubernetes defaults if not specified.",
										MarkdownDescription: "Resources is the resource requirements for the metrics container. Will be populated by Kubernetes defaults if not specified.",
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

					"networking": schema.SingleNestedAttribute{
						Description:         "Networking defines Couchbase cluster networking options such as network topology, TLS and DDNS settings.",
						MarkdownDescription: "Networking defines Couchbase cluster networking options such as network topology, TLS and DDNS settings.",
						Attributes: map[string]schema.Attribute{
							"address_family": schema.StringAttribute{
								Description:         "AddressFamily allows the manual selection of the address family to use. When this field is not set, Couchbase server will default to using IPv4 for internal communication and also support IPv6 on dual stack systems. Setting this field to either IPv4 or IPv6 will force Couchbase to use the selected protocol for internal communication, and also disable all other protocols to provide added security and simplicty when defining firewall rules.  Disabling of address families is only supported in Couchbase Server 7.0.2+.",
								MarkdownDescription: "AddressFamily allows the manual selection of the address family to use. When this field is not set, Couchbase server will default to using IPv4 for internal communication and also support IPv6 on dual stack systems. Setting this field to either IPv4 or IPv6 will force Couchbase to use the selected protocol for internal communication, and also disable all other protocols to provide added security and simplicty when defining firewall rules.  Disabling of address families is only supported in Couchbase Server 7.0.2+.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IPv4", "IPv6"),
								},
							},

							"admin_console_service_template": schema.SingleNestedAttribute{
								Description:         "AdminConsoleServiceTemplate provides a template used by the Operator to create and manage the admin console service.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",
								MarkdownDescription: "AdminConsoleServiceTemplate provides a template used by the Operator to create and manage the admin console service.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

									"spec": schema.SingleNestedAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service.",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.",
										Attributes: map[string]schema.Attribute{
											"allocate_load_balancer_node_ports": schema.BoolAttribute{
												Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
												MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_ip": schema.StringAttribute{
												Description:         "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_i_ps": schema.ListAttribute{
												Description:         "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_i_ps": schema.ListAttribute{
												Description:         "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
												MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_name": schema.StringAttribute{
												Description:         "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												MarkdownDescription: "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_traffic_policy": schema.StringAttribute{
												Description:         "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
												MarkdownDescription: "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"health_check_node_port": schema.Int64Attribute{
												Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
												MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"internal_traffic_policy": schema.StringAttribute{
												Description:         "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
												MarkdownDescription: "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_families": schema.ListAttribute{
												Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_class": schema.StringAttribute{
												Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_ip": schema.StringAttribute{
												Description:         "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version.",
												MarkdownDescription: "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_source_ranges": schema.ListAttribute{
												Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"session_affinity": schema.StringAttribute{
												Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
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
																Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
																MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
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
												Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
												MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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

							"admin_console_service_type": schema.StringAttribute{
								Description:         "DEPRECATED - by adminConsoleServiceTemplate. AdminConsoleServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",
								MarkdownDescription: "DEPRECATED - by adminConsoleServiceTemplate. AdminConsoleServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("NodePort", "LoadBalancer"),
								},
							},

							"admin_console_services": schema.ListAttribute{
								Description:         "DEPRECATED - not required by Couchbase Server. AdminConsoleServices is a selector to choose specific services to expose via the admin console. This field may contain any of 'data', 'index', 'query', 'search', 'eventing' and 'analytics'.  Each service may only be included once.",
								MarkdownDescription: "DEPRECATED - not required by Couchbase Server. AdminConsoleServices is a selector to choose specific services to expose via the admin console. This field may contain any of 'data', 'index', 'query', 'search', 'eventing' and 'analytics'.  Each service may only be included once.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_ui_over_http": schema.BoolAttribute{
								Description:         "DisableUIOverHTTP is used to explicitly enable and disable UI access over the HTTP protocol.  If not specified, this field defaults to false.",
								MarkdownDescription: "DisableUIOverHTTP is used to explicitly enable and disable UI access over the HTTP protocol.  If not specified, this field defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_ui_over_https": schema.BoolAttribute{
								Description:         "DisableUIOverHTTPS is used to explicitly enable and disable UI access over the HTTPS protocol.  If not specified, this field defaults to false.",
								MarkdownDescription: "DisableUIOverHTTPS is used to explicitly enable and disable UI access over the HTTPS protocol.  If not specified, this field defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns": schema.SingleNestedAttribute{
								Description:         "DNS defines information required for Dynamic DNS support.",
								MarkdownDescription: "DNS defines information required for Dynamic DNS support.",
								Attributes: map[string]schema.Attribute{
									"domain": schema.StringAttribute{
										Description:         "Domain is the domain to create pods in.  When populated the Operator will annotate the admin console and per-pod services with the key 'external-dns.alpha.kubernetes.io/hostname'.  These annotations can be used directly by a Kubernetes External-DNS controller to replicate load balancer service IP addresses into a public DNS server.",
										MarkdownDescription: "Domain is the domain to create pods in.  When populated the Operator will annotate the admin console and per-pod services with the key 'external-dns.alpha.kubernetes.io/hostname'.  These annotations can be used directly by a Kubernetes External-DNS controller to replicate load balancer service IP addresses into a public DNS server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"expose_admin_console": schema.BoolAttribute{
								Description:         "ExposeAdminConsole creates a service referencing the admin console. The service is configured by the adminConsoleServiceTemplate field.",
								MarkdownDescription: "ExposeAdminConsole creates a service referencing the admin console. The service is configured by the adminConsoleServiceTemplate field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exposed_feature_service_template": schema.SingleNestedAttribute{
								Description:         "ExposedFeatureServiceTemplate provides a template used by the Operator to create and manage per-pod services.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",
								MarkdownDescription: "ExposedFeatureServiceTemplate provides a template used by the Operator to create and manage per-pod services.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

									"spec": schema.SingleNestedAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service.",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.",
										Attributes: map[string]schema.Attribute{
											"allocate_load_balancer_node_ports": schema.BoolAttribute{
												Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
												MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_ip": schema.StringAttribute{
												Description:         "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_i_ps": schema.ListAttribute{
												Description:         "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_i_ps": schema.ListAttribute{
												Description:         "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
												MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_name": schema.StringAttribute{
												Description:         "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												MarkdownDescription: "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_traffic_policy": schema.StringAttribute{
												Description:         "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
												MarkdownDescription: "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"health_check_node_port": schema.Int64Attribute{
												Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
												MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"internal_traffic_policy": schema.StringAttribute{
												Description:         "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
												MarkdownDescription: "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_families": schema.ListAttribute{
												Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_class": schema.StringAttribute{
												Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_ip": schema.StringAttribute{
												Description:         "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version.",
												MarkdownDescription: "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_source_ranges": schema.ListAttribute{
												Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"session_affinity": schema.StringAttribute{
												Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
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
																Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
																MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
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
												Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
												MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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

							"exposed_feature_service_type": schema.StringAttribute{
								Description:         "DEPRECATED - by exposedFeatureServiceTemplate. ExposedFeatureServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",
								MarkdownDescription: "DEPRECATED - by exposedFeatureServiceTemplate. ExposedFeatureServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("NodePort", "LoadBalancer"),
								},
							},

							"exposed_feature_traffic_policy": schema.StringAttribute{
								Description:         "DEPRECATED  - by exposedFeatureServiceTemplate. ExposedFeatureTrafficPolicy defines how packets should be routed from a load balancer service to a Couchbase pod.  When local, traffic is routed directly to the pod.  When cluster, traffic is routed to any node, then forwarded on.  While cluster routing may be slower, there are some situations where it is required for connectivity.  This field must be either 'Cluster' or 'Local', defaulting to 'Local',",
								MarkdownDescription: "DEPRECATED  - by exposedFeatureServiceTemplate. ExposedFeatureTrafficPolicy defines how packets should be routed from a load balancer service to a Couchbase pod.  When local, traffic is routed directly to the pod.  When cluster, traffic is routed to any node, then forwarded on.  While cluster routing may be slower, there are some situations where it is required for connectivity.  This field must be either 'Cluster' or 'Local', defaulting to 'Local',",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Cluster", "Local"),
								},
							},

							"exposed_features": schema.ListAttribute{
								Description:         "ExposedFeatures is a list of Couchbase features to expose when using a networking model that exposes the Couchbase cluster externally to Kubernetes.  This field also triggers the creation of per-pod services used by clients to connect to the Couchbase cluster.  When admin, only the administrator port is exposed, allowing remote administration.  When xdcr, only the services required for remote replication are exposed. The xdcr feature is only required when the cluster is the destination of an XDCR replication.  When client, all services are exposed as required for client SDK operation. This field may contain any of 'admin', 'xdcr' and 'client'.  Each feature may only be included once.",
								MarkdownDescription: "ExposedFeatures is a list of Couchbase features to expose when using a networking model that exposes the Couchbase cluster externally to Kubernetes.  This field also triggers the creation of per-pod services used by clients to connect to the Couchbase cluster.  When admin, only the administrator port is exposed, allowing remote administration.  When xdcr, only the services required for remote replication are exposed. The xdcr feature is only required when the cluster is the destination of an XDCR replication.  When client, all services are exposed as required for client SDK operation. This field may contain any of 'admin', 'xdcr' and 'client'.  Each feature may only be included once.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. LoadBalancerSourceRanges applies only when an exposed service is of type LoadBalancer and limits the source IP ranges that are allowed to use the service.  Items must use IPv4 class-less interdomain routing (CIDR) notation e.g. 10.0.0.0/16.",
								MarkdownDescription: "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. LoadBalancerSourceRanges applies only when an exposed service is of type LoadBalancer and limits the source IP ranges that are allowed to use the service.  Items must use IPv4 class-less interdomain routing (CIDR) notation e.g. 10.0.0.0/16.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"network_platform": schema.StringAttribute{
								Description:         "NetworkPlatform is used to enable support for various networking technologies.  This field must be one of 'Istio'.",
								MarkdownDescription: "NetworkPlatform is used to enable support for various networking technologies.  This field must be one of 'Istio'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Istio"),
								},
							},

							"service_annotations": schema.MapAttribute{
								Description:         "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. ServiceAnnotations allows services to be annotated with custom labels. Operator annotations are merged on top of these so have precedence as they are required for correct operation.",
								MarkdownDescription: "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. ServiceAnnotations allows services to be annotated with custom labels. Operator annotations are merged on top of these so have precedence as they are required for correct operation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS defines the TLS configuration for the cluster including server and client certificate configuration, and TLS security policies.",
								MarkdownDescription: "TLS defines the TLS configuration for the cluster including server and client certificate configuration, and TLS security policies.",
								Attributes: map[string]schema.Attribute{
									"allow_plain_text_cert_reload": schema.BoolAttribute{
										Description:         "AllowPlainTextCertReload allows the reload of TLS certificates in plain text. This option should only be enabled as a means to recover connectivity with server in the event that any of the server certificates expire. When enabled the Operator only attempts plain text cert reloading when expired certificates are detected.",
										MarkdownDescription: "AllowPlainTextCertReload allows the reload of TLS certificates in plain text. This option should only be enabled as a means to recover connectivity with server in the event that any of the server certificates expire. When enabled the Operator only attempts plain text cert reloading when expired certificates are detected.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cipher_suites": schema.ListAttribute{
										Description:         "CipherSuites specifies a list of cipher suites for Couchbase server to select from when negotiating TLS handshakes with a client.  Suites are not validated by the Operator.  Run 'openssl ciphers -v' in a Couchbase server pod to interrogate supported values.",
										MarkdownDescription: "CipherSuites specifies a list of cipher suites for Couchbase server to select from when negotiating TLS handshakes with a client.  Suites are not validated by the Operator.  Run 'openssl ciphers -v' in a Couchbase server pod to interrogate supported values.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_certificate_paths": schema.ListNestedAttribute{
										Description:         "ClientCertificatePaths defines where to look in client certificates in order to extract the user name.",
										MarkdownDescription: "ClientCertificatePaths defines where to look in client certificates in order to extract the user name.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"delimiter": schema.StringAttribute{
													Description:         "Delimiter if specified allows a suffix to be stripped from the username, once extracted from the certificate path.",
													MarkdownDescription: "Delimiter if specified allows a suffix to be stripped from the username, once extracted from the certificate path.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path defines where in the X.509 specification to extract the username from. This field must be either 'subject.cn', 'san.uri', 'san.dnsname' or  'san.email'.",
													MarkdownDescription: "Path defines where in the X.509 specification to extract the username from. This field must be either 'subject.cn', 'san.uri', 'san.dnsname' or  'san.email'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^subject\.cn|san\.uri|san\.dnsname|san\.email$`), ""),
													},
												},

												"prefix": schema.StringAttribute{
													Description:         "Prefix allows a prefix to be stripped from the username, once extracted from the certificate path.",
													MarkdownDescription: "Prefix allows a prefix to be stripped from the username, once extracted from the certificate path.",
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

									"client_certificate_policy": schema.StringAttribute{
										Description:         "ClientCertificatePolicy defines the client authentication policy to use. If set, the Operator expects TLS configuration to contain a valid certificate/key pair for the Administrator account.",
										MarkdownDescription: "ClientCertificatePolicy defines the client authentication policy to use. If set, the Operator expects TLS configuration to contain a valid certificate/key pair for the Administrator account.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enable", "mandatory"),
										},
									},

									"node_to_node_encryption": schema.StringAttribute{
										Description:         "NodeToNodeEncryption specifies whether to encrypt data between Couchbase nodes within the same cluster.  This may come at the expense of performance.  When control plane only encryption is used, only cluster management traffic is encrypted between nodes.  When all, all traffic is encrypted, including database documents. When strict mode is used, it is the same as all, but also disables all plaintext ports.  Strict mode is only available on Couchbase Server versions 7.1 and greater. Node to node encryption can only be used when TLS certificates are managed by the Operator.  This field must be either 'ControlPlaneOnly', 'All', or 'Strict'.",
										MarkdownDescription: "NodeToNodeEncryption specifies whether to encrypt data between Couchbase nodes within the same cluster.  This may come at the expense of performance.  When control plane only encryption is used, only cluster management traffic is encrypted between nodes.  When all, all traffic is encrypted, including database documents. When strict mode is used, it is the same as all, but also disables all plaintext ports.  Strict mode is only available on Couchbase Server versions 7.1 and greater. Node to node encryption can only be used when TLS certificates are managed by the Operator.  This field must be either 'ControlPlaneOnly', 'All', or 'Strict'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ControlPlaneOnly", "All", "Strict"),
										},
									},

									"passphrase": schema.SingleNestedAttribute{
										Description:         "PassphraseConfig configures the passphrase key to use with encrypted certificates. The passphrase may be registered with Couchbase Server using a local script or a rest endpoint. Private key encryption is only available on Couchbase Server versions 7.1 and greater.",
										MarkdownDescription: "PassphraseConfig configures the passphrase key to use with encrypted certificates. The passphrase may be registered with Couchbase Server using a local script or a rest endpoint. Private key encryption is only available on Couchbase Server versions 7.1 and greater.",
										Attributes: map[string]schema.Attribute{
											"rest": schema.SingleNestedAttribute{
												Description:         "PassphraseRestConfig is the configuration to register a private key passphrase with a rest endpoint. When the private key is accessed, Couchbase Server attempts to extract the password by means of the specified endpoint. The response status must be 200 and the response text must be the exact passphrase excluding newlines and extraneous spaces.",
												MarkdownDescription: "PassphraseRestConfig is the configuration to register a private key passphrase with a rest endpoint. When the private key is accessed, Couchbase Server attempts to extract the password by means of the specified endpoint. The response status must be 200 and the response text must be the exact passphrase excluding newlines and extraneous spaces.",
												Attributes: map[string]schema.Attribute{
													"address_family": schema.StringAttribute{
														Description:         "AddressFamily is the address family to use. By default inet (meaning IPV4) is used.",
														MarkdownDescription: "AddressFamily is the address family to use. By default inet (meaning IPV4) is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("inet", "inet6"),
														},
													},

													"headers": schema.MapAttribute{
														Description:         "Headers is a map of one or more key-value pairs to pass alongside the Get request.",
														MarkdownDescription: "Headers is a map of one or more key-value pairs to pass alongside the Get request.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Timeout is  the number of milliseconds that must elapse before the call is timed out.",
														MarkdownDescription: "Timeout is  the number of milliseconds that must elapse before the call is timed out.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
														Description:         "URL is the endpoint to be called to retrieve the passphrase. URL will be called using the GET method and may use http/https protocol.",
														MarkdownDescription: "URL is the endpoint to be called to retrieve the passphrase. URL will be called using the GET method and may use http/https protocol.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"verify_peer": schema.BoolAttribute{
														Description:         "VerifyPeer ensures peer verification is performed when Https is used.",
														MarkdownDescription: "VerifyPeer ensures peer verification is performed when Https is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"script": schema.SingleNestedAttribute{
												Description:         "PassphraseScriptConfig is the configuration to register a private key passphrase with a script. The Operator auto-provisions the underlying script so this config simply provides a mechanism to perform the decryption of the Couchbase Private Key using a local script.",
												MarkdownDescription: "PassphraseScriptConfig is the configuration to register a private key passphrase with a script. The Operator auto-provisions the underlying script so this config simply provides a mechanism to perform the decryption of the Couchbase Private Key using a local script.",
												Attributes: map[string]schema.Attribute{
													"secret": schema.StringAttribute{
														Description:         "Secret is the secret containing the passphrase string. The secret is expected to contain 'passphrase' key with the passphrase string as a value.",
														MarkdownDescription: "Secret is the secret containing the passphrase string. The secret is expected to contain 'passphrase' key with the passphrase string as a value.",
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

									"root_c_as": schema.ListAttribute{
										Description:         "RootCAs defines a set of secrets that reside in this namespace that contain additional CA certificates that should be installed in Couchbase.  The CA certificates that are defined here are in addition to those defined for the cluster, optionally by couchbaseclusters.spec.networking.tls.secretSource, and thus should not be duplicated.  Each Secret referred to must be of well-known type 'kubernetes.io/tls' and must contain one or more CA certificates under the key 'tls.crt'. Multiple root CA certificates are only supported on Couchbase Server 7.1 and greater, and not with legacy couchbaseclusters.spec.networking.tls.static configuration.",
										MarkdownDescription: "RootCAs defines a set of secrets that reside in this namespace that contain additional CA certificates that should be installed in Couchbase.  The CA certificates that are defined here are in addition to those defined for the cluster, optionally by couchbaseclusters.spec.networking.tls.secretSource, and thus should not be duplicated.  Each Secret referred to must be of well-known type 'kubernetes.io/tls' and must contain one or more CA certificates under the key 'tls.crt'. Multiple root CA certificates are only supported on Couchbase Server 7.1 and greater, and not with legacy couchbaseclusters.spec.networking.tls.static configuration.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_source": schema.SingleNestedAttribute{
										Description:         "SecretSource enables the user to specify a secret conforming to the Kubernetes TLS secret specification that is used for the Couchbase server certificate, and optionally the Operator's client certificate, providing cert-manager compatibility without having to specify a separate root CA.  A server CA certificate must be supplied by one of the provided methods. Certificates referred to must conform to the keys of well-known type 'kubernetes.io/tls' with 'tls.crt' and 'tls.key'. If the 'tls.key' is an encrypted private key then the secret type can be the generic Opaque type since 'kubernetes.io/tls' type secrets cannot verify encrypted keys.",
										MarkdownDescription: "SecretSource enables the user to specify a secret conforming to the Kubernetes TLS secret specification that is used for the Couchbase server certificate, and optionally the Operator's client certificate, providing cert-manager compatibility without having to specify a separate root CA.  A server CA certificate must be supplied by one of the provided methods. Certificates referred to must conform to the keys of well-known type 'kubernetes.io/tls' with 'tls.crt' and 'tls.key'. If the 'tls.key' is an encrypted private key then the secret type can be the generic Opaque type since 'kubernetes.io/tls' type secrets cannot verify encrypted keys.",
										Attributes: map[string]schema.Attribute{
											"client_secret_name": schema.StringAttribute{
												Description:         "ClientSecretName specifies the secret name, in the same namespace as the cluster, the contains client TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the Kubernetes.io/tls secret type.",
												MarkdownDescription: "ClientSecretName specifies the secret name, in the same namespace as the cluster, the contains client TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the Kubernetes.io/tls secret type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"server_secret_name": schema.StringAttribute{
												Description:         "ServerSecretName specifies the secret name, in the same namespace as the cluster, that contains server TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the kubernetes.io/tls secret type.  It may also contain 'ca.crt'. Only a single PEM formated x509 certificate can be provided to 'ca.crt'. The single certificate may also bundle together multiple root CA certificates. Multiple root CA certificates are only supported on Couchbase Server 7.1 and greater.",
												MarkdownDescription: "ServerSecretName specifies the secret name, in the same namespace as the cluster, that contains server TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the kubernetes.io/tls secret type.  It may also contain 'ca.crt'. Only a single PEM formated x509 certificate can be provided to 'ca.crt'. The single certificate may also bundle together multiple root CA certificates. Multiple root CA certificates are only supported on Couchbase Server 7.1 and greater.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"static": schema.SingleNestedAttribute{
										Description:         "DEPRECATED - by couchbaseclusters.spec.networking.tls.secretSource. Static enables user to generate static x509 certificates and keys, put them into Kubernetes secrets, and specify them here.  Static secrets are Couchbase specific, and follow no well-known standards.",
										MarkdownDescription: "DEPRECATED - by couchbaseclusters.spec.networking.tls.secretSource. Static enables user to generate static x509 certificates and keys, put them into Kubernetes secrets, and specify them here.  Static secrets are Couchbase specific, and follow no well-known standards.",
										Attributes: map[string]schema.Attribute{
											"operator_secret": schema.StringAttribute{
												Description:         "OperatorSecret is a secret name containing TLS certs used by operator to talk securely to this cluster.  The secret must contain a CA certificate (data key ca.crt).  If client authentication is enabled, then the secret must also contain a client certificate chain (data key 'couchbase-operator.crt') and private key (data key 'couchbase-operator.key').",
												MarkdownDescription: "OperatorSecret is a secret name containing TLS certs used by operator to talk securely to this cluster.  The secret must contain a CA certificate (data key ca.crt).  If client authentication is enabled, then the secret must also contain a client certificate chain (data key 'couchbase-operator.crt') and private key (data key 'couchbase-operator.key').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"server_secret": schema.StringAttribute{
												Description:         "ServerSecret is a secret name containing TLS certs used by each Couchbase member pod for the communication between Couchbase server and its clients.  The secret must contain a certificate chain (data key 'chain.pem') and a private key (data key 'pkey.key').  The private key must be in the PKCS#1 RSA format.  The certificate chain must have a required set of X.509v3 subject alternative names for all cluster addressing modes.  See the Operator TLS documentation for more information.",
												MarkdownDescription: "ServerSecret is a secret name containing TLS certs used by each Couchbase member pod for the communication between Couchbase server and its clients.  The secret must contain a certificate chain (data key 'chain.pem') and a private key (data key 'pkey.key').  The private key must be in the PKCS#1 RSA format.  The certificate chain must have a required set of X.509v3 subject alternative names for all cluster addressing modes.  See the Operator TLS documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_minimum_version": schema.StringAttribute{
										Description:         "TLSMinimumVersion specifies the minimum TLS version the Couchbase server can negotiate with a client.  Must be one of TLS1.0, TLS1.1 TLS1.2 or TLS1.3, defaulting to TLS1.2.  TLS1.3 is only valid for Couchbase Server 7.1.0 onward.",
										MarkdownDescription: "TLSMinimumVersion specifies the minimum TLS version the Couchbase server can negotiate with a client.  Must be one of TLS1.0, TLS1.1 TLS1.2 or TLS1.3, defaulting to TLS1.2.  TLS1.3 is only valid for Couchbase Server 7.1.0 onward.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TLS1.0", "TLS1.1", "TLS1.2", "TLS1.3"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"wait_for_address_reachable": schema.StringAttribute{
								Description:         "WaitForAddressReachable is used to set the timeout between when polling of external addresses is started, and when it is deemed a failure.  Polling of DNS name availability inherently dangerous due to negative caching, so prefer the use of an initial 'waitForAddressReachableDelay' to allow propagation.",
								MarkdownDescription: "WaitForAddressReachable is used to set the timeout between when polling of external addresses is started, and when it is deemed a failure.  Polling of DNS name availability inherently dangerous due to negative caching, so prefer the use of an initial 'waitForAddressReachableDelay' to allow propagation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_for_address_reachable_delay": schema.StringAttribute{
								Description:         "WaitForAddressReachableDelay is used to defer operator checks that ensure external addresses are reachable before new nodes are balanced in to the cluster.  This prevents negative DNS caching while waiting for external-DDNS controllers to propagate addresses.",
								MarkdownDescription: "WaitForAddressReachableDelay is used to defer operator checks that ensure external addresses are reachable before new nodes are balanced in to the cluster.  This prevents negative DNS caching while waiting for external-DDNS controllers to propagate addresses.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": schema.BoolAttribute{
						Description:         "Paused is to pause the control of the operator for the Couchbase cluster. This does not pause the cluster itself, instead stopping the operator from taking any action.",
						MarkdownDescription: "Paused is to pause the control of the operator for the Couchbase cluster. This does not pause the cluster itself, instead stopping the operator from taking any action.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"platform": schema.StringAttribute{
						Description:         "Platform gives a hint as to what platform we are running on and how to configure services.  This field must be one of 'aws', 'gke' or 'azure'.",
						MarkdownDescription: "Platform gives a hint as to what platform we are running on and how to configure services.  This field must be one of 'aws', 'gke' or 'azure'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("aws", "gce", "azure"),
						},
					},

					"recovery_policy": schema.StringAttribute{
						Description:         "RecoveryPolicy controls how aggressive the Operator is when recovering cluster topology.  When PrioritizeDataIntegrity, the Operator will delegate failover exclusively to Couchbase server, relying on it to only allow recovery when safe to do so.  When PrioritizeUptime, the Operator will wait for a period after the expected auto-failover of the cluster, before forcefully failing-over the pods. This may cause data loss, and is only expected to be used on clusters with ephemeral data, where the loss of the pod means that the data is known to be unrecoverable. This field must be either 'PrioritizeDataIntegrity' or 'PrioritizeUptime', defaulting to 'PrioritizeDataIntegrity'.",
						MarkdownDescription: "RecoveryPolicy controls how aggressive the Operator is when recovering cluster topology.  When PrioritizeDataIntegrity, the Operator will delegate failover exclusively to Couchbase server, relying on it to only allow recovery when safe to do so.  When PrioritizeUptime, the Operator will wait for a period after the expected auto-failover of the cluster, before forcefully failing-over the pods. This may cause data loss, and is only expected to be used on clusters with ephemeral data, where the loss of the pod means that the data is known to be unrecoverable. This field must be either 'PrioritizeDataIntegrity' or 'PrioritizeUptime', defaulting to 'PrioritizeDataIntegrity'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("PrioritizeDataIntegrity", "PrioritizeUptime"),
						},
					},

					"rolling_upgrade": schema.SingleNestedAttribute{
						Description:         "When 'spec.upgradeStrategy' is set to 'RollingUpgrade' it will, by default, upgrade one pod at a time.  If this field is specified then that number can be increased.",
						MarkdownDescription: "When 'spec.upgradeStrategy' is set to 'RollingUpgrade' it will, by default, upgrade one pod at a time.  If this field is specified then that number can be increased.",
						Attributes: map[string]schema.Attribute{
							"max_upgradable": schema.Int64Attribute{
								Description:         "MaxUpgradable allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be greater than zero. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",
								MarkdownDescription: "MaxUpgradable allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be greater than zero. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"max_upgradable_percent": schema.StringAttribute{
								Description:         "MaxUpgradablePercent allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be an integer percentage, e.g. '10%', in the range 1% to 100%. Percentages are relative to the total cluster size, and rounded down to the nearest whole number, with a minimum of 1.  For example, a 10 pod cluster, and 25% allowed to upgrade, would yield 2.5 pods per iteration, rounded down to 2. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",
								MarkdownDescription: "MaxUpgradablePercent allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be an integer percentage, e.g. '10%', in the range 1% to 100%. Percentages are relative to the total cluster size, and rounded down to the nearest whole number, with a minimum of 1.  For example, a 10 pod cluster, and 25% allowed to upgrade, would yield 2.5 pods per iteration, rounded down to 2. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(100|[1-9][0-9]|[1-9])%$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"security": schema.SingleNestedAttribute{
						Description:         "Security defines Couchbase cluster security options such as the administrator account username and password, and user RBAC settings.",
						MarkdownDescription: "Security defines Couchbase cluster security options such as the administrator account username and password, and user RBAC settings.",
						Attributes: map[string]schema.Attribute{
							"admin_secret": schema.StringAttribute{
								Description:         "AdminSecret is the name of a Kubernetes secret to use for administrator authentication. The admin secret must contain the keys 'username' and 'password'.  The password data must be at least 6 characters in length, and not contain the any of the characters '()<>,;:'/[]?={}'.",
								MarkdownDescription: "AdminSecret is the name of a Kubernetes secret to use for administrator authentication. The admin secret must contain the keys 'username' and 'password'.  The password data must be at least 6 characters in length, and not contain the any of the characters '()<>,;:'/[]?={}'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"ldap": schema.SingleNestedAttribute{
								Description:         "LDAP provides settings to authenticate and authorize LDAP users with Couchbase Server. When specified, the Operator keeps these settings in sync with Cocuhbase Server's LDAP configuration. Leave empty to manually manage LDAP configuration.",
								MarkdownDescription: "LDAP provides settings to authenticate and authorize LDAP users with Couchbase Server. When specified, the Operator keeps these settings in sync with Cocuhbase Server's LDAP configuration. Leave empty to manually manage LDAP configuration.",
								Attributes: map[string]schema.Attribute{
									"authentication_enabled": schema.BoolAttribute{
										Description:         "AuthenticationEnabled allows users who attempt to access Couchbase Server without having been added as local users to be authenticated against the specified LDAP Host(s).",
										MarkdownDescription: "AuthenticationEnabled allows users who attempt to access Couchbase Server without having been added as local users to be authenticated against the specified LDAP Host(s).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"authorization_enabled": schema.BoolAttribute{
										Description:         "AuthorizationEnabled allows authenticated LDAP users to be authorized with RBAC roles granted to any Couchbase Server group associated with the user.",
										MarkdownDescription: "AuthorizationEnabled allows authenticated LDAP users to be authorized with RBAC roles granted to any Couchbase Server group associated with the user.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"bind_dn": schema.StringAttribute{
										Description:         "DN to use for searching users and groups synchronization. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "DN to use for searching users and groups synchronization. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"bind_secret": schema.StringAttribute{
										Description:         "BindSecret is the name of a Kubernetes secret to use containing password for LDAP user binding. The bindSecret must have a key with the name 'password' and a value which corresponds to the password of the binding LDAP user.",
										MarkdownDescription: "BindSecret is the name of a Kubernetes secret to use containing password for LDAP user binding. The bindSecret must have a key with the name 'password' and a value which corresponds to the password of the binding LDAP user.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"cacert": schema.StringAttribute{
										Description:         "DEPRECATED - Field is ignored, use tlsSecret. CA Certificate in PEM format to be used in LDAP server certificate validation. This cert is the string form of the secret provided to 'spec.tls.tlsSecret'.",
										MarkdownDescription: "DEPRECATED - Field is ignored, use tlsSecret. CA Certificate in PEM format to be used in LDAP server certificate validation. This cert is the string form of the secret provided to 'spec.tls.tlsSecret'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cache_value_lifetime": schema.Int64Attribute{
										Description:         "Lifetime of values in cache in milliseconds. Default 300000 ms.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "Lifetime of values in cache in milliseconds. Default 300000 ms.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"encryption": schema.StringAttribute{
										Description:         "Encryption determines how the connection with the LDAP server should be encrypted. Encryption may set as either StartTLSExtension, TLS, or false. When set to 'false' then no verification of the LDAP hostname is performed. When Encryption is StartTLSExtension, or TLS is set then the default behavior is to use the certificate already loaded into the Couchbase Cluster for certificate validation, otherwise 'ldap.tlsSecret' may be set to override The Couchbase certificate.",
										MarkdownDescription: "Encryption determines how the connection with the LDAP server should be encrypted. Encryption may set as either StartTLSExtension, TLS, or false. When set to 'false' then no verification of the LDAP hostname is performed. When Encryption is StartTLSExtension, or TLS is set then the default behavior is to use the certificate already loaded into the Couchbase Cluster for certificate validation, otherwise 'ldap.tlsSecret' may be set to override The Couchbase certificate.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("None", "StartTLSExtension", "TLS"),
										},
									},

									"groups_query": schema.StringAttribute{
										Description:         "LDAP query, to get the users' groups by username in RFC4516 format.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "LDAP query, to get the users' groups by username in RFC4516 format.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hosts": schema.ListAttribute{
										Description:         "List of LDAP hosts to provide authentication-support for Couchbase Server. Host name must be a valid IP address or DNS Name e.g openldap.default.svc, 10.0.92.147.",
										MarkdownDescription: "List of LDAP hosts to provide authentication-support for Couchbase Server. Host name must be a valid IP address or DNS Name e.g openldap.default.svc, 10.0.92.147.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"nested_groups_enabled": schema.BoolAttribute{
										Description:         "If enabled Couchbase server will try to recursively search for groups for every discovered ldap group. groups_query will be user for the search. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "If enabled Couchbase server will try to recursively search for groups for every discovered ldap group. groups_query will be user for the search. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"nested_groups_max_depth": schema.Int64Attribute{
										Description:         "Maximum number of recursive groups requests the server is allowed to perform. Requires NestedGroupsEnabled.  Values between 1 and 100: the default is 10. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "Maximum number of recursive groups requests the server is allowed to perform. Requires NestedGroupsEnabled.  Values between 1 and 100: the default is 10. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(100),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "LDAP port. This is typically 389 for LDAP, and 636 for LDAPS.",
										MarkdownDescription: "LDAP port. This is typically 389 for LDAP, and 636 for LDAPS.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"server_cert_validation": schema.BoolAttribute{
										Description:         "Whether server certificate validation be enabled.",
										MarkdownDescription: "Whether server certificate validation be enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_secret": schema.StringAttribute{
										Description:         "TLSSecret is the name of a Kubernetes secret to use explcitly for LDAP ca cert. If TLSSecret is not provided, certificates found in 'couchbaseclusters.spec.networking.tls.rootCAs' will be used instead. If provided, the secret must contain the ca to be used under the name 'ca.crt'.",
										MarkdownDescription: "TLSSecret is the name of a Kubernetes secret to use explcitly for LDAP ca cert. If TLSSecret is not provided, certificates found in 'couchbaseclusters.spec.networking.tls.rootCAs' will be used instead. If provided, the secret must contain the ca to be used under the name 'ca.crt'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_dn_mapping": schema.SingleNestedAttribute{
										Description:         "User to distinguished name (DN) mapping. If none is specified, the username is used as the user’s distinguished name.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "User to distinguished name (DN) mapping. If none is specified, the username is used as the user’s distinguished name.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										Attributes: map[string]schema.Attribute{
											"query": schema.StringAttribute{
												Description:         "Query is the LDAP query to run to map from Couchbase user to LDAP distinguished name.",
												MarkdownDescription: "Query is the LDAP query to run to map from Couchbase user to LDAP distinguished name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"template": schema.StringAttribute{
												Description:         "This field specifies list of templates to use for providing username to DN mapping. The template may contain a placeholder specified as '%u' to represent the Couchbase user who is attempting to gain access.",
												MarkdownDescription: "This field specifies list of templates to use for providing username to DN mapping. The template may contain a placeholder specified as '%u' to represent the Couchbase user who is attempting to gain access.",
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

							"rbac": schema.SingleNestedAttribute{
								Description:         "RBAC is the options provided for enabling and selecting RBAC User resources to manage.",
								MarkdownDescription: "RBAC is the options provided for enabling and selecting RBAC User resources to manage.",
								Attributes: map[string]schema.Attribute{
									"managed": schema.BoolAttribute{
										Description:         "Managed defines whether RBAC is managed by us or the clients.",
										MarkdownDescription: "Managed defines whether RBAC is managed by us or the clients.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector": schema.SingleNestedAttribute{
										Description:         "Selector is a label selector used to list RBAC resources in the namespace that are managed by the Operator.",
										MarkdownDescription: "Selector is a label selector used to list RBAC resources in the namespace that are managed by the Operator.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ui_session_timeout": schema.Int64Attribute{
								Description:         "UISessionTimeout sets how long, in minutes, before a user is declared inactive and signed out from the Couchbase Server UI. 0 represents no time out.",
								MarkdownDescription: "UISessionTimeout sets how long, in minutes, before a user is declared inactive and signed out from the Couchbase Server UI. 0 represents no time out.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(16666),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext allows the configuration of the security context for all Couchbase server pods.  When using persistent volumes you may need to set the fsGroup field in order to write to the volume.  For non-root clusters you must also set runAsUser to 1000, corresponding to the Couchbase user in official container images.  More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
						MarkdownDescription: "SecurityContext allows the configuration of the security context for all Couchbase server pods.  When using persistent volumes you may need to set the fsGroup field in order to write to the volume.  For non-root clusters you must also set runAsUser to 1000, corresponding to the Couchbase user in official container images.  More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
						Attributes: map[string]schema.Attribute{
							"fs_group": schema.Int64Attribute{
								Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fs_group_change_policy": schema.StringAttribute{
								Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_group": schema.Int64Attribute{
								Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
								MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_user": schema.Int64Attribute{
								Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"se_linux_options": schema.SingleNestedAttribute{
								Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"level": schema.StringAttribute{
										Description:         "Level is SELinux level label that applies to the container.",
										MarkdownDescription: "Level is SELinux level label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"role": schema.StringAttribute{
										Description:         "Role is a SELinux role label that applies to the container.",
										MarkdownDescription: "Role is a SELinux role label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is a SELinux type label that applies to the container.",
										MarkdownDescription: "Type is a SELinux type label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user": schema.StringAttribute{
										Description:         "User is a SELinux user label that applies to the container.",
										MarkdownDescription: "User is a SELinux user label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"seccomp_profile": schema.SingleNestedAttribute{
								Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"supplemental_groups": schema.ListAttribute{
								Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sysctls": schema.ListNestedAttribute{
								Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of a property to set",
											MarkdownDescription: "Name of a property to set",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value of a property to set",
											MarkdownDescription: "Value of a property to set",
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

							"windows_options": schema.SingleNestedAttribute{
								Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
								MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
								Attributes: map[string]schema.Attribute{
									"gmsa_credential_spec": schema.StringAttribute{
										Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
										MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"gmsa_credential_spec_name": schema.StringAttribute{
										Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_process": schema.BoolAttribute{
										Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
										MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user_name": schema.StringAttribute{
										Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

					"server_groups": schema.ListAttribute{
						Description:         "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",
						MarkdownDescription: "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"servers": schema.ListNestedAttribute{
						Description:         "Servers defines server classes for the Operator to provision and manage. A server class defines what services are running and how many members make up that class.  Specifying multiple server classes allows the Operator to provision clusters with Multi-Dimensional Scaling (MDS).  At least one server class must be defined, and at least one server class must be running the data service.",
						MarkdownDescription: "Servers defines server classes for the Operator to provision and manage. A server class defines what services are running and how many members make up that class.  Specifying multiple server classes allows the Operator to provision clusters with Multi-Dimensional Scaling (MDS).  At least one server class must be defined, and at least one server class must be running the data service.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"autoscale_enabled": schema.BoolAttribute{
									Description:         "AutoscaledEnabled defines whether the autoscaling feature is enabled for this class. When true, the Operator will create a CouchbaseAutoscaler resource for this server class.  The CouchbaseAutoscaler implements the Kubernetes scale API and can be controlled by the Kubernetes horizontal pod autoscaler (HPA).",
									MarkdownDescription: "AutoscaledEnabled defines whether the autoscaling feature is enabled for this class. When true, the Operator will create a CouchbaseAutoscaler resource for this server class.  The CouchbaseAutoscaler implements the Kubernetes scale API and can be controlled by the Kubernetes horizontal pod autoscaler (HPA).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"env": schema.ListNestedAttribute{
									Description:         "Env allows the setting of environment variables in the Couchbase server container.",
									MarkdownDescription: "Env allows the setting of environment variables in the Couchbase server container.",
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
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
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

								"env_from": schema.ListNestedAttribute{
									Description:         "EnvFrom allows the setting of environment variables in the Couchbase server container.",
									MarkdownDescription: "EnvFrom allows the setting of environment variables in the Couchbase server container.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"config_map_ref": schema.SingleNestedAttribute{
												Description:         "The ConfigMap to select from",
												MarkdownDescription: "The ConfigMap to select from",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap must be defined",
														MarkdownDescription: "Specify whether the ConfigMap must be defined",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": schema.StringAttribute{
												Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
												MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "The Secret to select from",
												MarkdownDescription: "The Secret to select from",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret must be defined",
														MarkdownDescription: "Specify whether the Secret must be defined",
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

								"name": schema.StringAttribute{
									Description:         "Name is a textual name for the server configuration and must be unique. The name is used by the operator to uniquely identify a server class, and map pods back to an intended configuration.",
									MarkdownDescription: "Name is a textual name for the server configuration and must be unique. The name is used by the operator to uniquely identify a server class, and map pods back to an intended configuration.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"pod": schema.SingleNestedAttribute{
									Description:         "Pod defines a template used to create pod for each Couchbase server instance.  Modifying pod metadata such as labels and annotations will update the pod in-place.  Any other modification will result in a cluster upgrade in order to fulfill the request. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#pod-v1-core",
									MarkdownDescription: "Pod defines a template used to create pod for each Couchbase server instance.  Modifying pod metadata such as labels and annotations will update the pod in-place.  Any other modification will result in a cluster upgrade in order to fulfill the request. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#pod-v1-core",
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
											MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
													MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
													MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

										"spec": schema.SingleNestedAttribute{
											Description:         "PodSpec is a description of a pod.",
											MarkdownDescription: "PodSpec is a description of a pod.",
											Attributes: map[string]schema.Attribute{
												"active_deadline_seconds": schema.Int64Attribute{
													Description:         "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
													MarkdownDescription: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"affinity": schema.SingleNestedAttribute{
													Description:         "If specified, the pod's scheduling constraints",
													MarkdownDescription: "If specified, the pod's scheduling constraints",
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

												"automount_service_account_token": schema.BoolAttribute{
													Description:         "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",
													MarkdownDescription: "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dns_config": schema.SingleNestedAttribute{
													Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
													MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
													Attributes: map[string]schema.Attribute{
														"nameservers": schema.ListAttribute{
															Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
															MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"options": schema.ListNestedAttribute{
															Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
															MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
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
															Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
															MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
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
													Description:         "Set DNS policy for the pod. Defaults to 'ClusterFirst'. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",
													MarkdownDescription: "Set DNS policy for the pod. Defaults to 'ClusterFirst'. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enable_service_links": schema.BoolAttribute{
													Description:         "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",
													MarkdownDescription: "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host_ipc": schema.BoolAttribute{
													Description:         "Use the host's ipc namespace. Optional: Default to false.",
													MarkdownDescription: "Use the host's ipc namespace. Optional: Default to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host_network": schema.BoolAttribute{
													Description:         "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",
													MarkdownDescription: "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host_pid": schema.BoolAttribute{
													Description:         "Use the host's pid namespace. Optional: Default to false.",
													MarkdownDescription: "Use the host's pid namespace. Optional: Default to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host_users": schema.BoolAttribute{
													Description:         "Use the host's user namespace. Optional: Default to true. If set to true or not present, the pod will be run in the host user namespace, useful for when the pod needs a feature only available to the host user namespace, such as loading a kernel module with CAP_SYS_MODULE. When set to false, a new userns is created for the pod. Setting false is useful for mitigating container breakout vulnerabilities even allowing users to run their containers as root without actually having root privileges on the host. This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.",
													MarkdownDescription: "Use the host's user namespace. Optional: Default to true. If set to true or not present, the pod will be run in the host user namespace, useful for when the pod needs a feature only available to the host user namespace, such as loading a kernel module with CAP_SYS_MODULE. When set to false, a new userns is created for the pod. Setting false is useful for mitigating container breakout vulnerabilities even allowing users to run their containers as root without actually having root privileges on the host. This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_pull_secrets": schema.ListNestedAttribute{
													Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
													MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
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

												"node_name": schema.StringAttribute{
													Description:         "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
													MarkdownDescription: "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selector": schema.MapAttribute{
													Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
													MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"os": schema.SingleNestedAttribute{
													Description:         "Specifies the OS of the containers in the pod. Some pod and container fields are restricted if this is set.  If the OS field is set to linux, the following fields must be unset: -securityContext.windowsOptions  If the OS field is set to windows, following fields must be unset: - spec.hostPID - spec.hostIPC - spec.hostUsers - spec.securityContext.seLinuxOptions - spec.securityContext.seccompProfile - spec.securityContext.fsGroup - spec.securityContext.fsGroupChangePolicy - spec.securityContext.sysctls - spec.shareProcessNamespace - spec.securityContext.runAsUser - spec.securityContext.runAsGroup - spec.securityContext.supplementalGroups - spec.containers[*].securityContext.seLinuxOptions - spec.containers[*].securityContext.seccompProfile - spec.containers[*].securityContext.capabilities - spec.containers[*].securityContext.readOnlyRootFilesystem - spec.containers[*].securityContext.privileged - spec.containers[*].securityContext.allowPrivilegeEscalation - spec.containers[*].securityContext.procMount - spec.containers[*].securityContext.runAsUser - spec.containers[*].securityContext.runAsGroup",
													MarkdownDescription: "Specifies the OS of the containers in the pod. Some pod and container fields are restricted if this is set.  If the OS field is set to linux, the following fields must be unset: -securityContext.windowsOptions  If the OS field is set to windows, following fields must be unset: - spec.hostPID - spec.hostIPC - spec.hostUsers - spec.securityContext.seLinuxOptions - spec.securityContext.seccompProfile - spec.securityContext.fsGroup - spec.securityContext.fsGroupChangePolicy - spec.securityContext.sysctls - spec.shareProcessNamespace - spec.securityContext.runAsUser - spec.securityContext.runAsGroup - spec.securityContext.supplementalGroups - spec.containers[*].securityContext.seLinuxOptions - spec.containers[*].securityContext.seccompProfile - spec.containers[*].securityContext.capabilities - spec.containers[*].securityContext.readOnlyRootFilesystem - spec.containers[*].securityContext.privileged - spec.containers[*].securityContext.allowPrivilegeEscalation - spec.containers[*].securityContext.procMount - spec.containers[*].securityContext.runAsUser - spec.containers[*].securityContext.runAsGroup",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null",
															MarkdownDescription: "Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"overhead": schema.MapAttribute{
													Description:         "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md",
													MarkdownDescription: "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"preemption_policy": schema.StringAttribute{
													Description:         "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
													MarkdownDescription: "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"priority": schema.Int64Attribute{
													Description:         "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",
													MarkdownDescription: "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"priority_class_name": schema.StringAttribute{
													Description:         "If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
													MarkdownDescription: "If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"runtime_class_name": schema.StringAttribute{
													Description:         "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the 'legacy' RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class",
													MarkdownDescription: "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the 'legacy' RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scheduler_name": schema.StringAttribute{
													Description:         "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
													MarkdownDescription: "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_account": schema.StringAttribute{
													Description:         "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.",
													MarkdownDescription: "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_account_name": schema.StringAttribute{
													Description:         "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
													MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"set_hostname_as_fqdn": schema.BoolAttribute{
													Description:         "If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINESYSTEMCurrentControlSetServicesTcpipParameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.",
													MarkdownDescription: "If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINESYSTEMCurrentControlSetServicesTcpipParameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"share_process_namespace": schema.BoolAttribute{
													Description:         "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.",
													MarkdownDescription: "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
													Description:         "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",
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
													Description:         "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
													MarkdownDescription: "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Resources are the resource requirements for the Couchbase server container. This field overrides any automatic allocation as defined by 'spec.autoResourceAllocation'.",
									MarkdownDescription: "Resources are the resource requirements for the Couchbase server container. This field overrides any automatic allocation as defined by 'spec.autoResourceAllocation'.",
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

								"server_groups": schema.ListAttribute{
									Description:         "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",
									MarkdownDescription: "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"services": schema.ListAttribute{
									Description:         "Services is the set of Couchbase services to run on this server class. At least one class must contain the data service.  The field may contain any of 'data', 'index', 'query', 'search', 'eventing' or 'analytics'. Each service may only be specified once.",
									MarkdownDescription: "Services is the set of Couchbase services to run on this server class. At least one class must contain the data service.  The field may contain any of 'data', 'index', 'query', 'search', 'eventing' or 'analytics'. Each service may only be specified once.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"size": schema.Int64Attribute{
									Description:         "Size is the expected requested of the server class.  This field must be greater than or equal to 1.",
									MarkdownDescription: "Size is the expected requested of the server class.  This field must be greater than or equal to 1.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
									},
								},

								"volume_mounts": schema.SingleNestedAttribute{
									Description:         "VolumeMounts define persistent volume claims to attach to pod.",
									MarkdownDescription: "VolumeMounts define persistent volume claims to attach to pod.",
									Attributes: map[string]schema.Attribute{
										"analytics": schema.ListAttribute{
											Description:         "AnalyticsClaims are persistent volumes that encompass analytics storage associated with the analytics service.  Analytics claims can only be used on server classes running the analytics service, and must be used in conjunction with the default claim. This field allows the analytics service to use different storage media (e.g. SSD), and scale horizontally, to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
											MarkdownDescription: "AnalyticsClaims are persistent volumes that encompass analytics storage associated with the analytics service.  Analytics claims can only be used on server classes running the analytics service, and must be used in conjunction with the default claim. This field allows the analytics service to use different storage media (e.g. SSD), and scale horizontally, to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"data": schema.StringAttribute{
											Description:         "DataClaim is a persistent volume that encompasses key/value storage associated with the data service.  The data claim can only be used on server classes running the data service, and must be used in conjunction with the default claim.  This field allows the data service to use different storage media (e.g. SSD) to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
											MarkdownDescription: "DataClaim is a persistent volume that encompasses key/value storage associated with the data service.  The data claim can only be used on server classes running the data service, and must be used in conjunction with the default claim.  This field allows the data service to use different storage media (e.g. SSD) to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default": schema.StringAttribute{
											Description:         "DefaultClaim is a persistent volume that encompasses all Couchbase persistent data, including document storage, indexes and logs.  The default volume can be used with any server class.  Use of the default claim allows the Operator to recover failed pods from the persistent volume far quicker than if the pod were using ephemeral storage.  The default claim cannot be used at the same time as the logs claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
											MarkdownDescription: "DefaultClaim is a persistent volume that encompasses all Couchbase persistent data, including document storage, indexes and logs.  The default volume can be used with any server class.  Use of the default claim allows the Operator to recover failed pods from the persistent volume far quicker than if the pod were using ephemeral storage.  The default claim cannot be used at the same time as the logs claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"index": schema.StringAttribute{
											Description:         "IndexClaim s a persistent volume that encompasses index storage associated with the index and search services.  The index claim can only be used on server classes running the index or search services, and must be used in conjunction with the default claim.  This field allows the index and/or search service to use different storage media (e.g. SSD) to improve performance of this service. This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst this references index primarily, note that the full text search (FTS) service also uses this same mount.",
											MarkdownDescription: "IndexClaim s a persistent volume that encompasses index storage associated with the index and search services.  The index claim can only be used on server classes running the index or search services, and must be used in conjunction with the default claim.  This field allows the index and/or search service to use different storage media (e.g. SSD) to improve performance of this service. This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst this references index primarily, note that the full text search (FTS) service also uses this same mount.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logs": schema.StringAttribute{
											Description:         "LogsClaim is a persistent volume that encompasses only Couchbase server logs to aid with supporting the product.  The logs claim can only be used on server classes running the following services: query, search & eventing.  The logs claim cannot be used at the same time as the default claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst the logs claim can be used with the search service, the recommendation is to use the default claim for these. The reason for this is that a failure of these nodes will require indexes to be rebuilt and subsequent performance impact.",
											MarkdownDescription: "LogsClaim is a persistent volume that encompasses only Couchbase server logs to aid with supporting the product.  The logs claim can only be used on server classes running the following services: query, search & eventing.  The logs claim cannot be used at the same time as the default claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst the logs claim can be used with the search service, the recommendation is to use the default claim for these. The reason for this is that a failure of these nodes will require indexes to be rebuilt and subsequent performance impact.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"software_update_notifications": schema.BoolAttribute{
						Description:         "SoftwareUpdateNotifications enables software update notifications in the UI. When enabled, the UI will alert when a Couchbase server upgrade is available.",
						MarkdownDescription: "SoftwareUpdateNotifications enables software update notifications in the UI. When enabled, the UI will alert when a Couchbase server upgrade is available.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upgrade_strategy": schema.StringAttribute{
						Description:         "UpgradeStrategy controls how aggressive the Operator is when performing a cluster upgrade.  When a rolling upgrade is requested, pods are upgraded one at a time.  This strategy is slower, however less disruptive.  When an immediate upgrade strategy is requested, all pods are upgraded at the same time.  This strategy is faster, but more disruptive.  This field must be either 'RollingUpgrade' or 'ImmediateUpgrade', defaulting to 'RollingUpgrade'.",
						MarkdownDescription: "UpgradeStrategy controls how aggressive the Operator is when performing a cluster upgrade.  When a rolling upgrade is requested, pods are upgraded one at a time.  This strategy is slower, however less disruptive.  When an immediate upgrade strategy is requested, all pods are upgraded at the same time.  This strategy is faster, but more disruptive.  This field must be either 'RollingUpgrade' or 'ImmediateUpgrade', defaulting to 'RollingUpgrade'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("RollingUpgrade", "ImmediateUpgrade"),
						},
					},

					"volume_claim_templates": schema.ListNestedAttribute{
						Description:         "VolumeClaimTemplates define the desired characteristics of a volume that can be requested/claimed by a pod, for example the storage class to use and the volume size.  Volume claim templates are referred to by name by server class volume mount configuration.",
						MarkdownDescription: "VolumeClaimTemplates define the desired characteristics of a volume that can be requested/claimed by a pod, for example the storage class to use and the volume size.  Volume claim templates are referred to by name by server class volume mount configuration.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"metadata": schema.SingleNestedAttribute{
									Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
									MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
											MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
											MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
											MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"spec": schema.SingleNestedAttribute{
									Description:         "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",
									MarkdownDescription: "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.ListAttribute{
											Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"data_source_ref": schema.SingleNestedAttribute{
											Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
											MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"xdcr": schema.SingleNestedAttribute{
						Description:         "XDCR defines whether the Operator should manage XDCR, remote clusters and how to lookup replication resources.",
						MarkdownDescription: "XDCR defines whether the Operator should manage XDCR, remote clusters and how to lookup replication resources.",
						Attributes: map[string]schema.Attribute{
							"managed": schema.BoolAttribute{
								Description:         "Managed defines whether XDCR is managed by the operator or not.",
								MarkdownDescription: "Managed defines whether XDCR is managed by the operator or not.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remote_clusters": schema.ListNestedAttribute{
								Description:         "RemoteClusters is a set of named remote clusters to establish replications to.",
								MarkdownDescription: "RemoteClusters is a set of named remote clusters to establish replications to.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication_secret": schema.StringAttribute{
											Description:         "AuthenticationSecret is a secret used to authenticate when establishing a remote connection.  It is only required when not using mTLS.  The secret must contain a username (secret key 'username') and password (secret key 'password').",
											MarkdownDescription: "AuthenticationSecret is a secret used to authenticate when establishing a remote connection.  It is only required when not using mTLS.  The secret must contain a username (secret key 'username') and password (secret key 'password').",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hostname": schema.StringAttribute{
											Description:         "Hostname is the connection string to use to connect the remote cluster.  To use IPv6, place brackets ('[', ']') around the IPv6 value.",
											MarkdownDescription: "Hostname is the connection string to use to connect the remote cluster.  To use IPv6, place brackets ('[', ']') around the IPv6 value.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^((couchbase|http)(s)?(://))?((\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b)|((([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]))|\[(\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*\]))(:[0-9]{0,5})?(\\{0,1}\?network=[^&]+)?$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name of the remote cluster.",
											MarkdownDescription: "Name of the remote cluster.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"replications": schema.SingleNestedAttribute{
											Description:         "Replications are replication streams from this cluster to the remote one. This field defines how to look up CouchbaseReplication resources.  By default any CouchbaseReplication resources in the namespace will be considered.",
											MarkdownDescription: "Replications are replication streams from this cluster to the remote one. This field defines how to look up CouchbaseReplication resources.  By default any CouchbaseReplication resources in the namespace will be considered.",
											Attributes: map[string]schema.Attribute{
												"selector": schema.SingleNestedAttribute{
													Description:         "Selector allows CouchbaseReplication resources to be filtered based on labels.",
													MarkdownDescription: "Selector allows CouchbaseReplication resources to be filtered based on labels.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS if specified references a resource containing the necessary certificate data for an encrypted connection.",
											MarkdownDescription: "TLS if specified references a resource containing the necessary certificate data for an encrypted connection.",
											Attributes: map[string]schema.Attribute{
												"secret": schema.StringAttribute{
													Description:         "Secret references a secret containing the CA certificate (data key 'ca'), and optionally a client certificate (data key 'certificate') and key (data key 'key').",
													MarkdownDescription: "Secret references a secret containing the CA certificate (data key 'ca'), and optionally a client certificate (data key 'certificate') and key (data key 'key').",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"uuid": schema.StringAttribute{
											Description:         "UUID of the remote cluster.  The UUID of a CouchbaseCluster resource is advertised in the status.clusterId field of the resource.",
											MarkdownDescription: "UUID of the remote cluster.  The UUID of a CouchbaseCluster resource is advertised in the status.clusterId field of the resource.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9a-f]{32}$`), ""),
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_cluster_v2")

	var model CouchbaseComCouchbaseClusterV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseCluster")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbaseclusters"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CouchbaseComCouchbaseClusterV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_cluster_v2")

	var data CouchbaseComCouchbaseClusterV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbaseclusters"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CouchbaseComCouchbaseClusterV2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_cluster_v2")

	var model CouchbaseComCouchbaseClusterV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseCluster")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbaseclusters"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CouchbaseComCouchbaseClusterV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_cluster_v2")

	var data CouchbaseComCouchbaseClusterV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbaseclusters"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbaseclusters"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *CouchbaseComCouchbaseClusterV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
