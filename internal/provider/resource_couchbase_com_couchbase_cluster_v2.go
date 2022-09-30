/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CouchbaseComCouchbaseClusterV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseClusterV2Resource)(nil)
)

type CouchbaseComCouchbaseClusterV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CouchbaseComCouchbaseClusterV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Cluster *struct {
			EventingServiceMemoryQuota *string `tfsdk:"eventing_service_memory_quota" yaml:"eventingServiceMemoryQuota,omitempty"`

			SearchServiceMemoryQuota *string `tfsdk:"search_service_memory_quota" yaml:"searchServiceMemoryQuota,omitempty"`

			AutoFailoverOnDataDiskIssues *bool `tfsdk:"auto_failover_on_data_disk_issues" yaml:"autoFailoverOnDataDiskIssues,omitempty"`

			ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

			IndexServiceMemoryQuota *string `tfsdk:"index_service_memory_quota" yaml:"indexServiceMemoryQuota,omitempty"`

			IndexStorageSetting *string `tfsdk:"index_storage_setting" yaml:"indexStorageSetting,omitempty"`

			AnalyticsServiceMemoryQuota *string `tfsdk:"analytics_service_memory_quota" yaml:"analyticsServiceMemoryQuota,omitempty"`

			AutoFailoverOnDataDiskIssuesTimePeriod *string `tfsdk:"auto_failover_on_data_disk_issues_time_period" yaml:"autoFailoverOnDataDiskIssuesTimePeriod,omitempty"`

			AutoFailoverServerGroup *bool `tfsdk:"auto_failover_server_group" yaml:"autoFailoverServerGroup,omitempty"`

			AutoFailoverTimeout *string `tfsdk:"auto_failover_timeout" yaml:"autoFailoverTimeout,omitempty"`

			Data *struct {
				ReaderThreads *int64 `tfsdk:"reader_threads" yaml:"readerThreads,omitempty"`

				WriterThreads *int64 `tfsdk:"writer_threads" yaml:"writerThreads,omitempty"`
			} `tfsdk:"data" yaml:"data,omitempty"`

			DataServiceMemoryQuota *string `tfsdk:"data_service_memory_quota" yaml:"dataServiceMemoryQuota,omitempty"`

			Indexer *struct {
				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				MaxRollbackPoints *int64 `tfsdk:"max_rollback_points" yaml:"maxRollbackPoints,omitempty"`

				MemorySnapshotInterval *string `tfsdk:"memory_snapshot_interval" yaml:"memorySnapshotInterval,omitempty"`

				StableSnapshotInterval *string `tfsdk:"stable_snapshot_interval" yaml:"stableSnapshotInterval,omitempty"`

				StorageMode *string `tfsdk:"storage_mode" yaml:"storageMode,omitempty"`

				Threads *int64 `tfsdk:"threads" yaml:"threads,omitempty"`
			} `tfsdk:"indexer" yaml:"indexer,omitempty"`

			AutoCompaction *struct {
				TimeWindow *struct {
					AbortCompactionOutsideWindow *bool `tfsdk:"abort_compaction_outside_window" yaml:"abortCompactionOutsideWindow,omitempty"`

					End *string `tfsdk:"end" yaml:"end,omitempty"`

					Start *string `tfsdk:"start" yaml:"start,omitempty"`
				} `tfsdk:"time_window" yaml:"timeWindow,omitempty"`

				TombstonePurgeInterval *string `tfsdk:"tombstone_purge_interval" yaml:"tombstonePurgeInterval,omitempty"`

				ViewFragmentationThreshold *struct {
					Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"view_fragmentation_threshold" yaml:"viewFragmentationThreshold,omitempty"`

				DatabaseFragmentationThreshold *struct {
					Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"database_fragmentation_threshold" yaml:"databaseFragmentationThreshold,omitempty"`

				ParallelCompaction *bool `tfsdk:"parallel_compaction" yaml:"parallelCompaction,omitempty"`
			} `tfsdk:"auto_compaction" yaml:"autoCompaction,omitempty"`

			AutoFailoverMaxCount *int64 `tfsdk:"auto_failover_max_count" yaml:"autoFailoverMaxCount,omitempty"`

			Query *struct {
				BackfillEnabled *bool `tfsdk:"backfill_enabled" yaml:"backfillEnabled,omitempty"`

				TemporarySpace *string `tfsdk:"temporary_space" yaml:"temporarySpace,omitempty"`

				TemporarySpaceUnlimited *bool `tfsdk:"temporary_space_unlimited" yaml:"temporarySpaceUnlimited,omitempty"`
			} `tfsdk:"query" yaml:"query,omitempty"`

			QueryServiceMemoryQuota *string `tfsdk:"query_service_memory_quota" yaml:"queryServiceMemoryQuota,omitempty"`
		} `tfsdk:"cluster" yaml:"cluster,omitempty"`

		Hibernate *bool `tfsdk:"hibernate" yaml:"hibernate,omitempty"`

		HibernationStrategy *string `tfsdk:"hibernation_strategy" yaml:"hibernationStrategy,omitempty"`

		Logging *struct {
			Server *struct {
				ConfigurationName *string `tfsdk:"configuration_name" yaml:"configurationName,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ManageConfiguration *bool `tfsdk:"manage_configuration" yaml:"manageConfiguration,omitempty"`

				Sidecar *struct {
					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					Resources *struct {
						Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

						Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					ConfigurationMountPath *string `tfsdk:"configuration_mount_path" yaml:"configurationMountPath,omitempty"`
				} `tfsdk:"sidecar" yaml:"sidecar,omitempty"`
			} `tfsdk:"server" yaml:"server,omitempty"`

			Audit *struct {
				DisabledEvents *[]string `tfsdk:"disabled_events" yaml:"disabledEvents,omitempty"`

				DisabledUsers *[]string `tfsdk:"disabled_users" yaml:"disabledUsers,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				GarbageCollection *struct {
					Sidecar *struct {
						Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

						Resources *struct {
							Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

							Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`

						Age *string `tfsdk:"age" yaml:"age,omitempty"`

						Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`
					} `tfsdk:"sidecar" yaml:"sidecar,omitempty"`
				} `tfsdk:"garbage_collection" yaml:"garbageCollection,omitempty"`

				Rotation *struct {
					Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"rotation" yaml:"rotation,omitempty"`
			} `tfsdk:"audit" yaml:"audit,omitempty"`

			LogRetentionCount *int64 `tfsdk:"log_retention_count" yaml:"logRetentionCount,omitempty"`

			LogRetentionTime *string `tfsdk:"log_retention_time" yaml:"logRetentionTime,omitempty"`
		} `tfsdk:"logging" yaml:"logging,omitempty"`

		Monitoring *struct {
			Prometheus *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Resources *struct {
					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`

					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				AuthorizationSecret *string `tfsdk:"authorization_secret" yaml:"authorizationSecret,omitempty"`
			} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`
		} `tfsdk:"monitoring" yaml:"monitoring,omitempty"`

		RollingUpgrade *struct {
			MaxUpgradable *int64 `tfsdk:"max_upgradable" yaml:"maxUpgradable,omitempty"`

			MaxUpgradablePercent *string `tfsdk:"max_upgradable_percent" yaml:"maxUpgradablePercent,omitempty"`
		} `tfsdk:"rolling_upgrade" yaml:"rollingUpgrade,omitempty"`

		Servers *[]struct {
			Pod *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				Spec *struct {
					AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" yaml:"automountServiceAccountToken,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostIPC *bool `tfsdk:"host_ipc" yaml:"hostIPC,omitempty"`

					HostPID *bool `tfsdk:"host_pid" yaml:"hostPID,omitempty"`

					Os *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"os" yaml:"os,omitempty"`

					Overhead *map[string]string `tfsdk:"overhead" yaml:"overhead,omitempty"`

					RuntimeClassName *string `tfsdk:"runtime_class_name" yaml:"runtimeClassName,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

					ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

					ShareProcessNamespace *bool `tfsdk:"share_process_namespace" yaml:"shareProcessNamespace,omitempty"`

					ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

					DnsConfig *struct {
						Searches *[]string `tfsdk:"searches" yaml:"searches,omitempty"`

						Nameservers *[]string `tfsdk:"nameservers" yaml:"nameservers,omitempty"`

						Options *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"options" yaml:"options,omitempty"`
					} `tfsdk:"dns_config" yaml:"dnsConfig,omitempty"`

					DnsPolicy *string `tfsdk:"dns_policy" yaml:"dnsPolicy,omitempty"`

					HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					PreemptionPolicy *string `tfsdk:"preemption_policy" yaml:"preemptionPolicy,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					SetHostnameAsFQDN *bool `tfsdk:"set_hostname_as_fqdn" yaml:"setHostnameAsFQDN,omitempty"`

					Tolerations *[]struct {
						Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

					NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

					NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

					SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

					ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

					TopologySpreadConstraints *[]struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"spec" yaml:"spec,omitempty"`
			} `tfsdk:"pod" yaml:"pod,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

			VolumeMounts *struct {
				Logs *string `tfsdk:"logs" yaml:"logs,omitempty"`

				Analytics *[]string `tfsdk:"analytics" yaml:"analytics,omitempty"`

				Data *string `tfsdk:"data" yaml:"data,omitempty"`

				Default *string `tfsdk:"default" yaml:"default,omitempty"`

				Index *string `tfsdk:"index" yaml:"index,omitempty"`
			} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

			AutoscaleEnabled *bool `tfsdk:"autoscale_enabled" yaml:"autoscaleEnabled,omitempty"`

			EnvFrom *[]struct {
				ConfigMapRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
			} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

			ServerGroups *[]string `tfsdk:"server_groups" yaml:"serverGroups,omitempty"`

			Services *[]string `tfsdk:"services" yaml:"services,omitempty"`

			Env *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"env" yaml:"env,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"servers" yaml:"servers,omitempty"`

		Backup *struct {
			UseIAMRole *bool `tfsdk:"use_iam_role" yaml:"useIAMRole,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			ObjectEndpoint *struct {
				Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				UseVirtualPath *bool `tfsdk:"use_virtual_path" yaml:"useVirtualPath,omitempty"`
			} `tfsdk:"object_endpoint" yaml:"objectEndpoint,omitempty"`

			S3Secret *string `tfsdk:"s3_secret" yaml:"s3Secret,omitempty"`

			Selector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`

			Tolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

			Managed *bool `tfsdk:"managed" yaml:"managed,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`
		} `tfsdk:"backup" yaml:"backup,omitempty"`

		UpgradeStrategy *string `tfsdk:"upgrade_strategy" yaml:"upgradeStrategy,omitempty"`

		SoftwareUpdateNotifications *bool `tfsdk:"software_update_notifications" yaml:"softwareUpdateNotifications,omitempty"`

		AutoscaleStabilizationPeriod *string `tfsdk:"autoscale_stabilization_period" yaml:"autoscaleStabilizationPeriod,omitempty"`

		EnablePreviewScaling *bool `tfsdk:"enable_preview_scaling" yaml:"enablePreviewScaling,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		RecoveryPolicy *string `tfsdk:"recovery_policy" yaml:"recoveryPolicy,omitempty"`

		Security *struct {
			Ldap *struct {
				Cacert *string `tfsdk:"cacert" yaml:"cacert,omitempty"`

				Encryption *string `tfsdk:"encryption" yaml:"encryption,omitempty"`

				Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

				UserDNMapping *struct {
					Query *string `tfsdk:"query" yaml:"query,omitempty"`

					Template *string `tfsdk:"template" yaml:"template,omitempty"`
				} `tfsdk:"user_dn_mapping" yaml:"userDNMapping,omitempty"`

				AuthenticationEnabled *bool `tfsdk:"authentication_enabled" yaml:"authenticationEnabled,omitempty"`

				BindSecret *string `tfsdk:"bind_secret" yaml:"bindSecret,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				TlsSecret *string `tfsdk:"tls_secret" yaml:"tlsSecret,omitempty"`

				GroupsQuery *string `tfsdk:"groups_query" yaml:"groupsQuery,omitempty"`

				NestedGroupsMaxDepth *int64 `tfsdk:"nested_groups_max_depth" yaml:"nestedGroupsMaxDepth,omitempty"`

				BindDN *string `tfsdk:"bind_dn" yaml:"bindDN,omitempty"`

				NestedGroupsEnabled *bool `tfsdk:"nested_groups_enabled" yaml:"nestedGroupsEnabled,omitempty"`

				ServerCertValidation *bool `tfsdk:"server_cert_validation" yaml:"serverCertValidation,omitempty"`

				AuthorizationEnabled *bool `tfsdk:"authorization_enabled" yaml:"authorizationEnabled,omitempty"`

				CacheValueLifetime *int64 `tfsdk:"cache_value_lifetime" yaml:"cacheValueLifetime,omitempty"`
			} `tfsdk:"ldap" yaml:"ldap,omitempty"`

			Rbac *struct {
				Selector *struct {
					MatchExpressions *[]struct {
						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Managed *bool `tfsdk:"managed" yaml:"managed,omitempty"`
			} `tfsdk:"rbac" yaml:"rbac,omitempty"`

			AdminSecret *string `tfsdk:"admin_secret" yaml:"adminSecret,omitempty"`
		} `tfsdk:"security" yaml:"security,omitempty"`

		AntiAffinity *bool `tfsdk:"anti_affinity" yaml:"antiAffinity,omitempty"`

		ServerGroups *[]string `tfsdk:"server_groups" yaml:"serverGroups,omitempty"`

		SecurityContext *struct {
			SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`

			Sysctls *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

			FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

			RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

			RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

			SeccompProfile *struct {
				LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

			FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" yaml:"fsGroupChangePolicy,omitempty"`

			RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

			SeLinuxOptions *struct {
				User *string `tfsdk:"user" yaml:"user,omitempty"`

				Level *string `tfsdk:"level" yaml:"level,omitempty"`

				Role *string `tfsdk:"role" yaml:"role,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

			WindowsOptions *struct {
				HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

				RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

				GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

				GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`
			} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
		} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

		Buckets *struct {
			Managed *bool `tfsdk:"managed" yaml:"managed,omitempty"`

			Selector *struct {
				MatchExpressions *[]struct {
					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`

			Synchronize *bool `tfsdk:"synchronize" yaml:"synchronize,omitempty"`
		} `tfsdk:"buckets" yaml:"buckets,omitempty"`

		EnableOnlineVolumeExpansion *bool `tfsdk:"enable_online_volume_expansion" yaml:"enableOnlineVolumeExpansion,omitempty"`

		Networking *struct {
			DisableUIOverHTTP *bool `tfsdk:"disable_ui_over_http" yaml:"disableUIOverHTTP,omitempty"`

			ExposedFeatureServiceTemplate *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				Spec *struct {
					SessionAffinity *string `tfsdk:"session_affinity" yaml:"sessionAffinity,omitempty"`

					HealthCheckNodePort *int64 `tfsdk:"health_check_node_port" yaml:"healthCheckNodePort,omitempty"`

					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					LoadBalancerClass *string `tfsdk:"load_balancer_class" yaml:"loadBalancerClass,omitempty"`

					ExternalTrafficPolicy *string `tfsdk:"external_traffic_policy" yaml:"externalTrafficPolicy,omitempty"`

					ClusterIP *string `tfsdk:"cluster_ip" yaml:"clusterIP,omitempty"`

					ExternalIPs *[]string `tfsdk:"external_i_ps" yaml:"externalIPs,omitempty"`

					ExternalName *string `tfsdk:"external_name" yaml:"externalName,omitempty"`

					LoadBalancerSourceRanges *[]string `tfsdk:"load_balancer_source_ranges" yaml:"loadBalancerSourceRanges,omitempty"`

					SessionAffinityConfig *struct {
						ClientIP *struct {
							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"client_ip" yaml:"clientIP,omitempty"`
					} `tfsdk:"session_affinity_config" yaml:"sessionAffinityConfig,omitempty"`

					ClusterIPs *[]string `tfsdk:"cluster_i_ps" yaml:"clusterIPs,omitempty"`

					InternalTrafficPolicy *string `tfsdk:"internal_traffic_policy" yaml:"internalTrafficPolicy,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					AllocateLoadBalancerNodePorts *bool `tfsdk:"allocate_load_balancer_node_ports" yaml:"allocateLoadBalancerNodePorts,omitempty"`

					LoadBalancerIP *string `tfsdk:"load_balancer_ip" yaml:"loadBalancerIP,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"spec" yaml:"spec,omitempty"`
			} `tfsdk:"exposed_feature_service_template" yaml:"exposedFeatureServiceTemplate,omitempty"`

			NetworkPlatform *string `tfsdk:"network_platform" yaml:"networkPlatform,omitempty"`

			AdminConsoleServiceType *string `tfsdk:"admin_console_service_type" yaml:"adminConsoleServiceType,omitempty"`

			ExposedFeatureServiceType *string `tfsdk:"exposed_feature_service_type" yaml:"exposedFeatureServiceType,omitempty"`

			ExposedFeatureTrafficPolicy *string `tfsdk:"exposed_feature_traffic_policy" yaml:"exposedFeatureTrafficPolicy,omitempty"`

			ExposedFeatures *[]string `tfsdk:"exposed_features" yaml:"exposedFeatures,omitempty"`

			LoadBalancerSourceRanges *[]string `tfsdk:"load_balancer_source_ranges" yaml:"loadBalancerSourceRanges,omitempty"`

			Tls *struct {
				ClientCertificatePaths *[]struct {
					Delimiter *string `tfsdk:"delimiter" yaml:"delimiter,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
				} `tfsdk:"client_certificate_paths" yaml:"clientCertificatePaths,omitempty"`

				ClientCertificatePolicy *string `tfsdk:"client_certificate_policy" yaml:"clientCertificatePolicy,omitempty"`

				NodeToNodeEncryption *string `tfsdk:"node_to_node_encryption" yaml:"nodeToNodeEncryption,omitempty"`

				RootCAs *[]string `tfsdk:"root_c_as" yaml:"rootCAs,omitempty"`

				SecretSource *struct {
					ClientSecretName *string `tfsdk:"client_secret_name" yaml:"clientSecretName,omitempty"`

					ServerSecretName *string `tfsdk:"server_secret_name" yaml:"serverSecretName,omitempty"`
				} `tfsdk:"secret_source" yaml:"secretSource,omitempty"`

				Static *struct {
					OperatorSecret *string `tfsdk:"operator_secret" yaml:"operatorSecret,omitempty"`

					ServerSecret *string `tfsdk:"server_secret" yaml:"serverSecret,omitempty"`
				} `tfsdk:"static" yaml:"static,omitempty"`

				TlsMinimumVersion *string `tfsdk:"tls_minimum_version" yaml:"tlsMinimumVersion,omitempty"`

				CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`

			WaitForAddressReachableDelay *string `tfsdk:"wait_for_address_reachable_delay" yaml:"waitForAddressReachableDelay,omitempty"`

			DisableUIOverHTTPS *bool `tfsdk:"disable_ui_over_https" yaml:"disableUIOverHTTPS,omitempty"`

			ServiceAnnotations *map[string]string `tfsdk:"service_annotations" yaml:"serviceAnnotations,omitempty"`

			WaitForAddressReachable *string `tfsdk:"wait_for_address_reachable" yaml:"waitForAddressReachable,omitempty"`

			AdminConsoleServiceTemplate *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				Spec *struct {
					AllocateLoadBalancerNodePorts *bool `tfsdk:"allocate_load_balancer_node_ports" yaml:"allocateLoadBalancerNodePorts,omitempty"`

					ExternalIPs *[]string `tfsdk:"external_i_ps" yaml:"externalIPs,omitempty"`

					ExternalTrafficPolicy *string `tfsdk:"external_traffic_policy" yaml:"externalTrafficPolicy,omitempty"`

					LoadBalancerIP *string `tfsdk:"load_balancer_ip" yaml:"loadBalancerIP,omitempty"`

					SessionAffinityConfig *struct {
						ClientIP *struct {
							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"client_ip" yaml:"clientIP,omitempty"`
					} `tfsdk:"session_affinity_config" yaml:"sessionAffinityConfig,omitempty"`

					ClusterIP *string `tfsdk:"cluster_ip" yaml:"clusterIP,omitempty"`

					ExternalName *string `tfsdk:"external_name" yaml:"externalName,omitempty"`

					InternalTrafficPolicy *string `tfsdk:"internal_traffic_policy" yaml:"internalTrafficPolicy,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					LoadBalancerClass *string `tfsdk:"load_balancer_class" yaml:"loadBalancerClass,omitempty"`

					LoadBalancerSourceRanges *[]string `tfsdk:"load_balancer_source_ranges" yaml:"loadBalancerSourceRanges,omitempty"`

					ClusterIPs *[]string `tfsdk:"cluster_i_ps" yaml:"clusterIPs,omitempty"`

					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					SessionAffinity *string `tfsdk:"session_affinity" yaml:"sessionAffinity,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					HealthCheckNodePort *int64 `tfsdk:"health_check_node_port" yaml:"healthCheckNodePort,omitempty"`
				} `tfsdk:"spec" yaml:"spec,omitempty"`
			} `tfsdk:"admin_console_service_template" yaml:"adminConsoleServiceTemplate,omitempty"`

			AdminConsoleServices *[]string `tfsdk:"admin_console_services" yaml:"adminConsoleServices,omitempty"`

			Dns *struct {
				Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`
			} `tfsdk:"dns" yaml:"dns,omitempty"`

			ExposeAdminConsole *bool `tfsdk:"expose_admin_console" yaml:"exposeAdminConsole,omitempty"`

			AddressFamily *string `tfsdk:"address_family" yaml:"addressFamily,omitempty"`
		} `tfsdk:"networking" yaml:"networking,omitempty"`

		Paused *bool `tfsdk:"paused" yaml:"paused,omitempty"`

		Platform *string `tfsdk:"platform" yaml:"platform,omitempty"`

		VolumeClaimTemplates *[]struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Spec *struct {
				DataSourceRef *struct {
					ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"data_source_ref" yaml:"dataSourceRef,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				Selector *struct {
					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

				VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`

				VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

				AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"volume_claim_templates" yaml:"volumeClaimTemplates,omitempty"`

		Xdcr *struct {
			RemoteClusters *[]struct {
				AuthenticationSecret *string `tfsdk:"authentication_secret" yaml:"authenticationSecret,omitempty"`

				Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Replications *struct {
					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"replications" yaml:"replications,omitempty"`

				Tls *struct {
					Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`

				Uuid *string `tfsdk:"uuid" yaml:"uuid,omitempty"`
			} `tfsdk:"remote_clusters" yaml:"remoteClusters,omitempty"`

			Managed *bool `tfsdk:"managed" yaml:"managed,omitempty"`
		} `tfsdk:"xdcr" yaml:"xdcr,omitempty"`

		AutoResourceAllocation *struct {
			CpuLimits *string `tfsdk:"cpu_limits" yaml:"cpuLimits,omitempty"`

			CpuRequests *string `tfsdk:"cpu_requests" yaml:"cpuRequests,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			OverheadPercent *int64 `tfsdk:"overhead_percent" yaml:"overheadPercent,omitempty"`
		} `tfsdk:"auto_resource_allocation" yaml:"autoResourceAllocation,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCouchbaseComCouchbaseClusterV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseClusterV2Resource{}
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_cluster_v2"
}

func (r *CouchbaseComCouchbaseClusterV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "The CouchbaseCluster resource represents a Couchbase cluster.  It allows configuration of cluster topology, networking, storage and security options.",
		MarkdownDescription: "The CouchbaseCluster resource represents a Couchbase cluster.  It allows configuration of cluster topology, networking, storage and security options.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "ClusterSpec is the specification for a CouchbaseCluster resources, and allows the cluster to be customized.",
				MarkdownDescription: "ClusterSpec is the specification for a CouchbaseCluster resources, and allows the cluster to be customized.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cluster": {
						Description:         "ClusterSettings define Couchbase cluster-wide settings such as memory allocation, failover characteristics and index settings.",
						MarkdownDescription: "ClusterSettings define Couchbase cluster-wide settings such as memory allocation, failover characteristics and index settings.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"eventing_service_memory_quota": {
								Description:         "EventingServiceMemQuota is the amount of memory that should be allocated to the eventing service. This value is per-pod, and only applicable to pods belonging to server classes running the eventing service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "EventingServiceMemQuota is the amount of memory that should be allocated to the eventing service. This value is per-pod, and only applicable to pods belonging to server classes running the eventing service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"search_service_memory_quota": {
								Description:         "SearchServiceMemQuota is the amount of memory that should be allocated to the search service. This value is per-pod, and only applicable to pods belonging to server classes running the search service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "SearchServiceMemQuota is the amount of memory that should be allocated to the search service. This value is per-pod, and only applicable to pods belonging to server classes running the search service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auto_failover_on_data_disk_issues": {
								Description:         "AutoFailoverOnDataDiskIssues defines whether Couchbase server should failover a pod if a disk issue was detected.",
								MarkdownDescription: "AutoFailoverOnDataDiskIssues defines whether Couchbase server should failover a pod if a disk issue was detected.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_name": {
								Description:         "ClusterName defines the name of the cluster, as displayed in the Couchbase UI. By default, the cluster name is that specified in the CouchbaseCluster resource's metadata.",
								MarkdownDescription: "ClusterName defines the name of the cluster, as displayed in the Couchbase UI. By default, the cluster name is that specified in the CouchbaseCluster resource's metadata.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"index_service_memory_quota": {
								Description:         "IndexServiceMemQuota is the amount of memory that should be allocated to the index service. This value is per-pod, and only applicable to pods belonging to server classes running the index service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "IndexServiceMemQuota is the amount of memory that should be allocated to the index service. This value is per-pod, and only applicable to pods belonging to server classes running the index service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"index_storage_setting": {
								Description:         "DEPRECATED - by indexer. The index storage mode to use for secondary indexing.  This field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.  This field is immutable and cannot be changed unless there are no server classes running the index service in the cluster.",
								MarkdownDescription: "DEPRECATED - by indexer. The index storage mode to use for secondary indexing.  This field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.  This field is immutable and cannot be changed unless there are no server classes running the index service in the cluster.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"analytics_service_memory_quota": {
								Description:         "AnalyticsServiceMemQuota is the amount of memory that should be allocated to the analytics service. This value is per-pod, and only applicable to pods belonging to server classes running the analytics service.  This field must be a quantity greater than or equal to 1Gi.  This field defaults to 1Gi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "AnalyticsServiceMemQuota is the amount of memory that should be allocated to the analytics service. This value is per-pod, and only applicable to pods belonging to server classes running the analytics service.  This field must be a quantity greater than or equal to 1Gi.  This field defaults to 1Gi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auto_failover_on_data_disk_issues_time_period": {
								Description:         "AutoFailoverOnDataDiskIssuesTimePeriod defines how long to wait for transient errors before failing over a faulty disk.  This field must be in the range 5-3600s, defaulting to 120s.  More info:  https://golang.org/pkg/time/#ParseDuration",
								MarkdownDescription: "AutoFailoverOnDataDiskIssuesTimePeriod defines how long to wait for transient errors before failing over a faulty disk.  This field must be in the range 5-3600s, defaulting to 120s.  More info:  https://golang.org/pkg/time/#ParseDuration",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auto_failover_server_group": {
								Description:         "AutoFailoverServerGroup whether to enable failing over a server group.",
								MarkdownDescription: "AutoFailoverServerGroup whether to enable failing over a server group.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auto_failover_timeout": {
								Description:         "AutoFailoverTimeout defines how long Couchbase server will wait between a pod being witnessed as down, until when it will failover the pod.  Couchbase server will only failover pods if it deems it safe to do so, and not result in data loss.  This field must be in the range 5-3600s, defaulting to 120s. More info:  https://golang.org/pkg/time/#ParseDuration",
								MarkdownDescription: "AutoFailoverTimeout defines how long Couchbase server will wait between a pod being witnessed as down, until when it will failover the pod.  Couchbase server will only failover pods if it deems it safe to do so, and not result in data loss.  This field must be in the range 5-3600s, defaulting to 120s. More info:  https://golang.org/pkg/time/#ParseDuration",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data": {
								Description:         "Data allows the data service to be configured.",
								MarkdownDescription: "Data allows the data service to be configured.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"reader_threads": {
										Description:         "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use.  If not specified, this defaults to the default value set by Couchbase Server.",
										MarkdownDescription: "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use.  If not specified, this defaults to the default value set by Couchbase Server.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"writer_threads": {
										Description:         "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This setting is especially relevant when using 'durable writes', increasing this field will have a large impact on performance.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use. If not specified, this defaults to the default value set by Couchbase Server.",
										MarkdownDescription: "ReaderThreads allows the number of threads used by the data service, per pod, to be altered.  This setting is especially relevant when using 'durable writes', increasing this field will have a large impact on performance.  This value must be between 4 and 64 threads, and should only be increased where there are sufficient CPU resources allocated for their use. If not specified, this defaults to the default value set by Couchbase Server.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_service_memory_quota": {
								Description:         "DataServiceMemQuota is the amount of memory that should be allocated to the data service. This value is per-pod, and only applicable to pods belonging to server classes running the data service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "DataServiceMemQuota is the amount of memory that should be allocated to the data service. This value is per-pod, and only applicable to pods belonging to server classes running the data service.  This field must be a quantity greater than or equal to 256Mi.  This field defaults to 256Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"indexer": {
								Description:         "Indexer allows the indexer to be configured.",
								MarkdownDescription: "Indexer allows the indexer to be configured.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"log_level": {
										Description:         "LogLevel controls the verbosity of indexer logs.  This field must be one of 'silent', 'fatal', 'error', 'warn', 'info', 'verbose', 'timing', 'debug' or 'trace', defaulting to 'info'.",
										MarkdownDescription: "LogLevel controls the verbosity of indexer logs.  This field must be one of 'silent', 'fatal', 'error', 'warn', 'info', 'verbose', 'timing', 'debug' or 'trace', defaulting to 'info'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_rollback_points": {
										Description:         "MaxRollbackPoints controls the number of checkpoints that can be rolled back to.  The default is 2, with a minimum of 1.",
										MarkdownDescription: "MaxRollbackPoints controls the number of checkpoints that can be rolled back to.  The default is 2, with a minimum of 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_snapshot_interval": {
										Description:         "MemorySnapshotInterval controls when memory indexes should be snapshotted. This defaults to 200ms, and must be greater than or equal to 1ms.",
										MarkdownDescription: "MemorySnapshotInterval controls when memory indexes should be snapshotted. This defaults to 200ms, and must be greater than or equal to 1ms.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stable_snapshot_interval": {
										Description:         "StableSnapshotInterval controls when disk indexes should be snapshotted. This defaults to 5s, and must be greater than or equal to 1ms.",
										MarkdownDescription: "StableSnapshotInterval controls when disk indexes should be snapshotted. This defaults to 5s, and must be greater than or equal to 1ms.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_mode": {
										Description:         "StorageMode controls the underlying storage engine for indexes.  Once set it can only be modified if there are no nodes in the cluster running the index service.  The field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.",
										MarkdownDescription: "StorageMode controls the underlying storage engine for indexes.  Once set it can only be modified if there are no nodes in the cluster running the index service.  The field must be one of 'memory_optimized' or 'plasma', defaulting to 'memory_optimized'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"threads": {
										Description:         "Threads controls the number of processor threads to use for indexing. A value of 0 means 1 per CPU.  This attribute must be greater than or equal to 0, defaulting to 0.",
										MarkdownDescription: "Threads controls the number of processor threads to use for indexing. A value of 0 means 1 per CPU.  This attribute must be greater than or equal to 0, defaulting to 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auto_compaction": {
								Description:         "AutoCompaction allows the configuration of auto-compaction, including on what conditions disk space is reclaimed and when it is allowed to run.",
								MarkdownDescription: "AutoCompaction allows the configuration of auto-compaction, including on what conditions disk space is reclaimed and when it is allowed to run.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"time_window": {
										Description:         "TimeWindow allows restriction of when compaction can occur.",
										MarkdownDescription: "TimeWindow allows restriction of when compaction can occur.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"abort_compaction_outside_window": {
												Description:         "AbortCompactionOutsideWindow stops compaction processes when the process moves outside the window.",
												MarkdownDescription: "AbortCompactionOutsideWindow stops compaction processes when the process moves outside the window.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"end": {
												Description:         "End is a wallclock time, in the form HH:MM, when a compaction should stop.",
												MarkdownDescription: "End is a wallclock time, in the form HH:MM, when a compaction should stop.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"start": {
												Description:         "Start is a wallclock time, in the form HH:MM, when a compaction is permitted to start.",
												MarkdownDescription: "Start is a wallclock time, in the form HH:MM, when a compaction is permitted to start.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tombstone_purge_interval": {
										Description:         "TombstonePurgeInterval controls how long to wait before purging tombstones. This field must be in the range 1h-1440h, defaulting to 72h. More info:  https://golang.org/pkg/time/#ParseDuration",
										MarkdownDescription: "TombstonePurgeInterval controls how long to wait before purging tombstones. This field must be in the range 1h-1440h, defaulting to 72h. More info:  https://golang.org/pkg/time/#ParseDuration",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"view_fragmentation_threshold": {
										Description:         "ViewFragmentationThreshold defines triggers for when view compaction should start.",
										MarkdownDescription: "ViewFragmentationThreshold defines triggers for when view compaction should start.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"percent": {
												Description:         "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",
												MarkdownDescription: "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												MarkdownDescription: "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"database_fragmentation_threshold": {
										Description:         "DatabaseFragmentationThreshold defines triggers for when database compaction should start.",
										MarkdownDescription: "DatabaseFragmentationThreshold defines triggers for when database compaction should start.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"percent": {
												Description:         "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",
												MarkdownDescription: "Percent is the percentage of disk fragmentation after which to decompaction will be triggered. This field must be in the range 2-100, defaulting to 30.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												MarkdownDescription: "Size is the amount of disk framentation, that once exceeded, will trigger decompaction. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"parallel_compaction": {
										Description:         "ParallelCompaction controls whether database and view compactions can happen in parallel.",
										MarkdownDescription: "ParallelCompaction controls whether database and view compactions can happen in parallel.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auto_failover_max_count": {
								Description:         "AutoFailoverMaxCount is the maximum number of automatic failovers Couchbase server will allow before not allowing any more.  This field must be between 1-3, default 3.",
								MarkdownDescription: "AutoFailoverMaxCount is the maximum number of automatic failovers Couchbase server will allow before not allowing any more.  This field must be between 1-3, default 3.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"query": {
								Description:         "Query allows the query service to be configured.",
								MarkdownDescription: "Query allows the query service to be configured.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"backfill_enabled": {
										Description:         "BackfillEnabled allows the query service to backfill.",
										MarkdownDescription: "BackfillEnabled allows the query service to backfill.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"temporary_space": {
										Description:         "TemporarySpace allows the temporary storage used by the query service backfill, per-pod, to be modified.  This field requires 'backfillEnabled' to be set to true in order to have any effect. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
										MarkdownDescription: "TemporarySpace allows the temporary storage used by the query service backfill, per-pod, to be modified.  This field requires 'backfillEnabled' to be set to true in order to have any effect. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"temporary_space_unlimited": {
										Description:         "TemporarySpaceUnlimited allows the temporary storage used by the query service backfill, per-pod, to be unconstrained.  This field requires 'backfillEnabled' to be set to true in order to have any effect. This field overrides 'temporarySpace'.",
										MarkdownDescription: "TemporarySpaceUnlimited allows the temporary storage used by the query service backfill, per-pod, to be unconstrained.  This field requires 'backfillEnabled' to be set to true in order to have any effect. This field overrides 'temporarySpace'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"query_service_memory_quota": {
								Description:         "QueryServiceMemQuota is a dummy field.  By default, Couchbase server provides no memory resource constraints for the query service, so this has no effect on Couchbase server.  It is, however, used when the spec.autoResourceAllocation feature is enabled, and is used to define the amount of memory reserved by the query service for use with Kubernetes resource scheduling. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "QueryServiceMemQuota is a dummy field.  By default, Couchbase server provides no memory resource constraints for the query service, so this has no effect on Couchbase server.  It is, however, used when the spec.autoResourceAllocation feature is enabled, and is used to define the amount of memory reserved by the query service for use with Kubernetes resource scheduling. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hibernate": {
						Description:         "Hibernate is whether to hibernate the cluster.",
						MarkdownDescription: "Hibernate is whether to hibernate the cluster.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hibernation_strategy": {
						Description:         "HibernationStrategy defines how to hibernate the cluster.  When Immediate the Operator will immediately delete all pods and take no further action until the hibernate field is set to false.",
						MarkdownDescription: "HibernationStrategy defines how to hibernate the cluster.  When Immediate the Operator will immediately delete all pods and take no further action until the hibernate field is set to false.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"logging": {
						Description:         "Logging defines Operator logging options.",
						MarkdownDescription: "Logging defines Operator logging options.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"server": {
								Description:         "Specification of all logging configuration required to manage the sidecar containers in each pod.",
								MarkdownDescription: "Specification of all logging configuration required to manage the sidecar containers in each pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration_name": {
										Description:         "ConfigurationName is the name of the Secret to use holding the logging configuration in the namespace. A Secret is used to ensure we can safely store credentials but this can be populated from plaintext if acceptable too. If it does not exist then one will be created with defaults in the namespace so it can be easily updated whilst running. Note that if running multiple clusters in the same kubernetes namespace then you should use a separate Secret for each, otherwise the first cluster will take ownership (if created) and the Secret will be cleaned up when that cluster is removed. If running clusters in separate namespaces then they will be separate Secrets anyway.",
										MarkdownDescription: "ConfigurationName is the name of the Secret to use holding the logging configuration in the namespace. A Secret is used to ensure we can safely store credentials but this can be populated from plaintext if acceptable too. If it does not exist then one will be created with defaults in the namespace so it can be easily updated whilst running. Note that if running multiple clusters in the same kubernetes namespace then you should use a separate Secret for each, otherwise the first cluster will take ownership (if created) and the Secret will be cleaned up when that cluster is removed. If running clusters in separate namespaces then they will be separate Secrets anyway.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled is a boolean that enables the logging sidecar container.",
										MarkdownDescription: "Enabled is a boolean that enables the logging sidecar container.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"manage_configuration": {
										Description:         "A boolean which indicates whether the operator should manage the configuration or not. If omitted then this defaults to true which means the operator will attempt to reconcile it to default values. To use a custom configuration make sure to set this to false. Note that the ownership of any Secret is not changed so if a Secret is created externally it can be updated by the operator but it's ownership stays the same so it will be cleaned up when it's owner is.",
										MarkdownDescription: "A boolean which indicates whether the operator should manage the configuration or not. If omitted then this defaults to true which means the operator will attempt to reconcile it to default values. To use a custom configuration make sure to set this to false. Note that the ownership of any Secret is not changed so if a Secret is created externally it can be updated by the operator but it's ownership stays the same so it will be cleaned up when it's owner is.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sidecar": {
										Description:         "Any specific logging sidecar container configuration.",
										MarkdownDescription: "Any specific logging sidecar container configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"image": {
												Description:         "Image is the image to be used to deal with logging as a sidecar. No validation is carried out as this can be any arbitrary repo and tag. It will default to the latest supported version of Fluent Bit.",
												MarkdownDescription: "Image is the image to be used to deal with logging as a sidecar. No validation is carried out as this can be any arbitrary repo and tag. It will default to the latest supported version of Fluent Bit.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "Resources is the resource requirements for the sidecar container. Will be populated by Kubernetes defaults if not specified.",
												MarkdownDescription: "Resources is the resource requirements for the sidecar container. Will be populated by Kubernetes defaults if not specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"configuration_mount_path": {
												Description:         "ConfigurationMountPath is the location to mount the ConfigurationName Secret into the image. If another log shipping image is used that needs a different mount then modify this. Note that the configuration file must be called 'fluent-bit.conf' at the root of this path, there is no provision for overriding the name of the config file passed as the COUCHBASE_LOGS_CONFIG_FILE environment variable.",
												MarkdownDescription: "ConfigurationMountPath is the location to mount the ConfigurationName Secret into the image. If another log shipping image is used that needs a different mount then modify this. Note that the configuration file must be called 'fluent-bit.conf' at the root of this path, there is no provision for overriding the name of the config file passed as the COUCHBASE_LOGS_CONFIG_FILE environment variable.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"audit": {
								Description:         "Used to manage the audit configuration directly",
								MarkdownDescription: "Used to manage the audit configuration directly",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"disabled_events": {
										Description:         "The list of event ids to disable for auditing purposes. This is passed to the REST API with no verification by the operator. Refer to the documentation for details: https://docs.couchbase.com/server/current/audit-event-reference/audit-event-reference.html",
										MarkdownDescription: "The list of event ids to disable for auditing purposes. This is passed to the REST API with no verification by the operator. Refer to the documentation for details: https://docs.couchbase.com/server/current/audit-event-reference/audit-event-reference.html",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disabled_users": {
										Description:         "The list of users to ignore for auditing purposes. This is passed to the REST API with minimal validation it meets an acceptable regex pattern. Refer to the documentation for full details on how to configure this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html#ignoring-events-by-user",
										MarkdownDescription: "The list of users to ignore for auditing purposes. This is passed to the REST API with minimal validation it meets an acceptable regex pattern. Refer to the documentation for full details on how to configure this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html#ignoring-events-by-user",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled is a boolean that enables the audit capabilities.",
										MarkdownDescription: "Enabled is a boolean that enables the audit capabilities.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"garbage_collection": {
										Description:         "Handle all optional garbage collection (GC) configuration for the audit functionality. This is not part of the audit REST API, it is intended to handle GC automatically for the audit logs. By default the Couchbase Server rotates the audit logs but does not clean up the rotated logs. This is left as an operation for the cluster administrator to manage, the operator allows for us to automate this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",
										MarkdownDescription: "Handle all optional garbage collection (GC) configuration for the audit functionality. This is not part of the audit REST API, it is intended to handle GC automatically for the audit logs. By default the Couchbase Server rotates the audit logs but does not clean up the rotated logs. This is left as an operation for the cluster administrator to manage, the operator allows for us to automate this: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"sidecar": {
												Description:         "Provide the sidecar configuration required (if so desired) to automatically clean up audit logs.",
												MarkdownDescription: "Provide the sidecar configuration required (if so desired) to automatically clean up audit logs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"interval": {
														Description:         "The interval at which to check for rotated log files to remove, defaults to 20 minutes.",
														MarkdownDescription: "The interval at which to check for rotated log files to remove, defaults to 20 minutes.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Resources is the resource requirements for the cleanup container. Will be populated by Kubernetes defaults if not specified.",
														MarkdownDescription: "Resources is the resource requirements for the cleanup container. Will be populated by Kubernetes defaults if not specified.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"requests": {
																Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"age": {
														Description:         "The minimum age of rotated log files to remove, defaults to one hour.",
														MarkdownDescription: "The minimum age of rotated log files to remove, defaults to one hour.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"enabled": {
														Description:         "Enable this sidecar by setting to true, defaults to being disabled.",
														MarkdownDescription: "Enable this sidecar by setting to true, defaults to being disabled.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "Image is the image to be used to run the audit sidecar helper. No validation is carried out as this can be any arbitrary repo and tag.",
														MarkdownDescription: "Image is the image to be used to run the audit sidecar helper. No validation is carried out as this can be any arbitrary repo and tag.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rotation": {
										Description:         "The interval to optionally rotate the audit log. This is passed to the REST API, see here for details: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",
										MarkdownDescription: "The interval to optionally rotate the audit log. This is passed to the REST API, see here for details: https://docs.couchbase.com/server/current/manage/manage-security/manage-auditing.html",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"interval": {
												Description:         "The interval at which to rotate log files, defaults to 15 minutes.",
												MarkdownDescription: "The interval at which to rotate log files, defaults to 15 minutes.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "Size allows the specification of a rotation size for the log, defaults to 20Mi. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
												MarkdownDescription: "Size allows the specification of a rotation size for the log, defaults to 20Mi. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_retention_count": {
								Description:         "LogRetentionCount gives the number of persistent log PVCs to keep.",
								MarkdownDescription: "LogRetentionCount gives the number of persistent log PVCs to keep.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_retention_time": {
								Description:         "LogRetentionTime gives the time to keep persistent log PVCs alive for.",
								MarkdownDescription: "LogRetentionTime gives the time to keep persistent log PVCs alive for.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring": {
						Description:         "Monitoring defines any Operator managed integration into 3rd party monitoring infrastructure.",
						MarkdownDescription: "Monitoring defines any Operator managed integration into 3rd party monitoring infrastructure.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"prometheus": {
								Description:         "Prometheus provides integration with Prometheus monitoring.",
								MarkdownDescription: "Prometheus provides integration with Prometheus monitoring.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Enabled is a boolean that enables/disables the metrics sidecar container.",
										MarkdownDescription: "Enabled is a boolean that enables/disables the metrics sidecar container.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "Image is the metrics image to be used to collect metrics. No validation is carried out as this can be any arbitrary repo and tag.",
										MarkdownDescription: "Image is the metrics image to be used to collect metrics. No validation is carried out as this can be any arbitrary repo and tag.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"resources": {
										Description:         "Resources is the resource requirements for the metrics container. Will be populated by Kubernetes defaults if not specified.",
										MarkdownDescription: "Resources is the resource requirements for the metrics container. Will be populated by Kubernetes defaults if not specified.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization_secret": {
										Description:         "AuthorizationSecret is the name of a Kubernetes secret that contains a bearer token to authorize GET requests to the metrics endpoint",
										MarkdownDescription: "AuthorizationSecret is the name of a Kubernetes secret that contains a bearer token to authorize GET requests to the metrics endpoint",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rolling_upgrade": {
						Description:         "When 'spec.upgradeStrategy' is set to 'RollingUpgrade' it will, by default, upgrade one pod at a time.  If this field is specified then that number can be increased.",
						MarkdownDescription: "When 'spec.upgradeStrategy' is set to 'RollingUpgrade' it will, by default, upgrade one pod at a time.  If this field is specified then that number can be increased.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_upgradable": {
								Description:         "MaxUpgradable allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be greater than zero. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",
								MarkdownDescription: "MaxUpgradable allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be greater than zero. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_upgradable_percent": {
								Description:         "MaxUpgradablePercent allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be an integer percentage, e.g. '10%', in the range 1% to 100%. Percentages are relative to the total cluster size, and rounded down to the nearest whole number, with a minimum of 1.  For example, a 10 pod cluster, and 25% allowed to upgrade, would yield 2.5 pods per iteration, rounded down to 2. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",
								MarkdownDescription: "MaxUpgradablePercent allows the number of pods affected by an upgrade at any one time to be increased.  By default a rolling upgrade will upgrade one pod at a time.  This field allows that limit to be removed. This field must be an integer percentage, e.g. '10%', in the range 1% to 100%. Percentages are relative to the total cluster size, and rounded down to the nearest whole number, with a minimum of 1.  For example, a 10 pod cluster, and 25% allowed to upgrade, would yield 2.5 pods per iteration, rounded down to 2. The smallest of 'maxUpgradable' and 'maxUpgradablePercent' takes precedence if both are defined.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"servers": {
						Description:         "Servers defines server classes for the Operator to provision and manage. A server class defines what services are running and how many members make up that class.  Specifying multiple server classes allows the Operator to provision clusters with Multi-Dimensional Scaling (MDS).  At least one server class must be defined, and at least one server class must be running the data service.",
						MarkdownDescription: "Servers defines server classes for the Operator to provision and manage. A server class defines what services are running and how many members make up that class.  Specifying multiple server classes allows the Operator to provision clusters with Multi-Dimensional Scaling (MDS).  At least one server class must be defined, and at least one server class must be running the data service.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"pod": {
								Description:         "Pod defines a template used to create pod for each Couchbase server instance.  Modifying pod metadata such as labels and annotations will update the pod in-place.  Any other modification will result in a cluster upgrade in order to fulfill the request. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#pod-v1-core",
								MarkdownDescription: "Pod defines a template used to create pod for each Couchbase server instance.  Modifying pod metadata such as labels and annotations will update the pod in-place.  Any other modification will result in a cluster upgrade in order to fulfill the request. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#pod-v1-core",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": {
										Description:         "PodSpec is a description of a pod.",
										MarkdownDescription: "PodSpec is a description of a pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"automount_service_account_token": {
												Description:         "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",
												MarkdownDescription: "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_service_links": {
												Description:         "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",
												MarkdownDescription: "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_ipc": {
												Description:         "Use the host's ipc namespace. Optional: Default to false.",
												MarkdownDescription: "Use the host's ipc namespace. Optional: Default to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_pid": {
												Description:         "Use the host's pid namespace. Optional: Default to false.",
												MarkdownDescription: "Use the host's pid namespace. Optional: Default to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"os": {
												Description:         "Specifies the OS of the containers in the pod. Some pod and container fields are restricted if this is set.  If the OS field is set to linux, the following fields must be unset: -securityContext.windowsOptions  If the OS field is set to windows, following fields must be unset: - spec.hostPID - spec.hostIPC - spec.securityContext.seLinuxOptions - spec.securityContext.seccompProfile - spec.securityContext.fsGroup - spec.securityContext.fsGroupChangePolicy - spec.securityContext.sysctls - spec.shareProcessNamespace - spec.securityContext.runAsUser - spec.securityContext.runAsGroup - spec.securityContext.supplementalGroups - spec.containers[*].securityContext.seLinuxOptions - spec.containers[*].securityContext.seccompProfile - spec.containers[*].securityContext.capabilities - spec.containers[*].securityContext.readOnlyRootFilesystem - spec.containers[*].securityContext.privileged - spec.containers[*].securityContext.allowPrivilegeEscalation - spec.containers[*].securityContext.procMount - spec.containers[*].securityContext.runAsUser - spec.containers[*].securityContext.runAsGroup This is an alpha field and requires the IdentifyPodOS feature",
												MarkdownDescription: "Specifies the OS of the containers in the pod. Some pod and container fields are restricted if this is set.  If the OS field is set to linux, the following fields must be unset: -securityContext.windowsOptions  If the OS field is set to windows, following fields must be unset: - spec.hostPID - spec.hostIPC - spec.securityContext.seLinuxOptions - spec.securityContext.seccompProfile - spec.securityContext.fsGroup - spec.securityContext.fsGroupChangePolicy - spec.securityContext.sysctls - spec.shareProcessNamespace - spec.securityContext.runAsUser - spec.securityContext.runAsGroup - spec.securityContext.supplementalGroups - spec.containers[*].securityContext.seLinuxOptions - spec.containers[*].securityContext.seccompProfile - spec.containers[*].securityContext.capabilities - spec.containers[*].securityContext.readOnlyRootFilesystem - spec.containers[*].securityContext.privileged - spec.containers[*].securityContext.allowPrivilegeEscalation - spec.containers[*].securityContext.procMount - spec.containers[*].securityContext.runAsUser - spec.containers[*].securityContext.runAsGroup This is an alpha field and requires the IdentifyPodOS feature",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null",
														MarkdownDescription: "Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"overhead": {
												Description:         "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md This field is beta-level as of Kubernetes v1.18, and is only honored by servers that enable the PodOverhead feature.",
												MarkdownDescription: "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md This field is beta-level as of Kubernetes v1.18, and is only honored by servers that enable the PodOverhead feature.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"runtime_class_name": {
												Description:         "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the 'legacy' RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class This is a beta feature as of Kubernetes v1.14.",
												MarkdownDescription: "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the 'legacy' RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class This is a beta feature as of Kubernetes v1.14.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_grace_period_seconds": {
												Description:         "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_class_name": {
												Description:         "If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
												MarkdownDescription: "If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account": {
												Description:         "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.",
												MarkdownDescription: "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"share_process_namespace": {
												Description:         "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.",
												MarkdownDescription: "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"active_deadline_seconds": {
												Description:         "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
												MarkdownDescription: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_config": {
												Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
												MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"searches": {
														Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
														MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nameservers": {
														Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
														MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"options": {
														Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
														MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Required.",
																MarkdownDescription: "Required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_policy": {
												Description:         "Set DNS policy for the pod. Defaults to 'ClusterFirst'. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",
												MarkdownDescription: "Set DNS policy for the pod. Defaults to 'ClusterFirst'. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_network": {
												Description:         "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",
												MarkdownDescription: "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_secrets": {
												Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
												MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preemption_policy": {
												Description:         "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.",
												MarkdownDescription: "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",
												MarkdownDescription: "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set_hostname_as_fqdn": {
												Description:         "If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINESYSTEMCurrentControlSetServicesTcpipParameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.",
												MarkdownDescription: "If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINESYSTEMCurrentControlSetServicesTcpipParameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tolerations": {
												Description:         "If specified, the pod's tolerations.",
												MarkdownDescription: "If specified, the pod's tolerations.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"effect": {
														Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
														MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
														MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operator": {
														Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
														MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"toleration_seconds": {
														Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
														MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
														MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_name": {
												Description:         "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
												MarkdownDescription: "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selector": {
												Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
												MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
												MarkdownDescription: "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account_name": {
												Description:         "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
												MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topology_spread_constraints": {
												Description:         "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
												MarkdownDescription: "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
														MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0: | zone1 | zone2 | zone3 | |   P   |   P   |       | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1; scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
														MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0: | zone1 | zone2 | zone3 | |   P   |   P   |       | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1; scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"topology_key": {
														Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. It's a required field.",
														MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. It's a required field.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"when_unsatisfiable": {
														Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
														MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources are the resource requirements for the Couchbase server container. This field overrides any automatic allocation as defined by 'spec.autoResourceAllocation'.",
								MarkdownDescription: "Resources are the resource requirements for the Couchbase server container. This field overrides any automatic allocation as defined by 'spec.autoResourceAllocation'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "Size is the expected requested of the server class.  This field must be greater than or equal to 1.",
								MarkdownDescription: "Size is the expected requested of the server class.  This field must be greater than or equal to 1.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"volume_mounts": {
								Description:         "VolumeMounts define persistent volume claims to attach to pod.",
								MarkdownDescription: "VolumeMounts define persistent volume claims to attach to pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"logs": {
										Description:         "LogsClaim is a persistent volume that encompasses only Couchbase server logs to aid with supporting the product.  The logs claim can only be used on server classes running the following services: query, search & eventing.  The logs claim cannot be used at the same time as the default claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst the logs claim can be used with the search service, the recommendation is to use the default claim for these. The reason for this is that a failure of these nodes will require indexes to be rebuilt and subsequent performance impact.",
										MarkdownDescription: "LogsClaim is a persistent volume that encompasses only Couchbase server logs to aid with supporting the product.  The logs claim can only be used on server classes running the following services: query, search & eventing.  The logs claim cannot be used at the same time as the default claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst the logs claim can be used with the search service, the recommendation is to use the default claim for these. The reason for this is that a failure of these nodes will require indexes to be rebuilt and subsequent performance impact.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"analytics": {
										Description:         "AnalyticsClaims are persistent volumes that encompass analytics storage associated with the analytics service.  Analytics claims can only be used on server classes running the analytics service, and must be used in conjunction with the default claim. This field allows the analytics service to use different storage media (e.g. SSD), and scale horizontally, to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
										MarkdownDescription: "AnalyticsClaims are persistent volumes that encompass analytics storage associated with the analytics service.  Analytics claims can only be used on server classes running the analytics service, and must be used in conjunction with the default claim. This field allows the analytics service to use different storage media (e.g. SSD), and scale horizontally, to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"data": {
										Description:         "DataClaim is a persistent volume that encompasses key/value storage associated with the data service.  The data claim can only be used on server classes running the data service, and must be used in conjunction with the default claim.  This field allows the data service to use different storage media (e.g. SSD) to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
										MarkdownDescription: "DataClaim is a persistent volume that encompasses key/value storage associated with the data service.  The data claim can only be used on server classes running the data service, and must be used in conjunction with the default claim.  This field allows the data service to use different storage media (e.g. SSD) to improve performance of this service.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default": {
										Description:         "DefaultClaim is a persistent volume that encompasses all Couchbase persistent data, including document storage, indexes and logs.  The default volume can be used with any server class.  Use of the default claim allows the Operator to recover failed pods from the persistent volume far quicker than if the pod were using ephemeral storage.  The default claim cannot be used at the same time as the logs claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",
										MarkdownDescription: "DefaultClaim is a persistent volume that encompasses all Couchbase persistent data, including document storage, indexes and logs.  The default volume can be used with any server class.  Use of the default claim allows the Operator to recover failed pods from the persistent volume far quicker than if the pod were using ephemeral storage.  The default claim cannot be used at the same time as the logs claim within the same server class.  This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"index": {
										Description:         "IndexClaim s a persistent volume that encompasses index storage associated with the index and search services.  The index claim can only be used on server classes running the index or search services, and must be used in conjunction with the default claim.  This field allows the index and/or search service to use different storage media (e.g. SSD) to improve performance of this service. This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst this references index primarily, note that the full text search (FTS) service also uses this same mount.",
										MarkdownDescription: "IndexClaim s a persistent volume that encompasses index storage associated with the index and search services.  The index claim can only be used on server classes running the index or search services, and must be used in conjunction with the default claim.  This field allows the index and/or search service to use different storage media (e.g. SSD) to improve performance of this service. This field references a volume claim template name as defined in 'spec.volumeClaimTemplates'. Whilst this references index primarily, note that the full text search (FTS) service also uses this same mount.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"autoscale_enabled": {
								Description:         "AutoscaledEnabled defines whether the autoscaling feature is enabled for this class. When true, the Operator will create a CouchbaseAutoscaler resource for this server class.  The CouchbaseAutoscaler implements the Kubernetes scale API and can be controlled by the Kubernetes horizontal pod autoscaler (HPA).",
								MarkdownDescription: "AutoscaledEnabled defines whether the autoscaling feature is enabled for this class. When true, the Operator will create a CouchbaseAutoscaler resource for this server class.  The CouchbaseAutoscaler implements the Kubernetes scale API and can be controlled by the Kubernetes horizontal pod autoscaler (HPA).",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env_from": {
								Description:         "EnvFrom allows the setting of environment variables in the Couchbase server container.",
								MarkdownDescription: "EnvFrom allows the setting of environment variables in the Couchbase server container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"config_map_ref": {
										Description:         "The ConfigMap to select from",
										MarkdownDescription: "The ConfigMap to select from",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap must be defined",
												MarkdownDescription: "Specify whether the ConfigMap must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix": {
										Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
										MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "The Secret to select from",
										MarkdownDescription: "The Secret to select from",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret must be defined",
												MarkdownDescription: "Specify whether the Secret must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_groups": {
								Description:         "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",
								MarkdownDescription: "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"services": {
								Description:         "Services is the set of Couchbase services to run on this server class. At least one class must contain the data service.  The field may contain any of 'data', 'index', 'query', 'search', 'eventing' or 'analytics'. Each service may only be specified once.",
								MarkdownDescription: "Services is the set of Couchbase services to run on this server class. At least one class must contain the data service.  The field may contain any of 'data', 'index', 'query', 'search', 'eventing' or 'analytics'. Each service may only be specified once.",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"env": {
								Description:         "Env allows the setting of environment variables in the Couchbase server container.",
								MarkdownDescription: "Env allows the setting of environment variables in the Couchbase server container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name is a textual name for the server configuration and must be unique. The name is used by the operator to uniquely identify a server class, and map pods back to an intended configuration.",
								MarkdownDescription: "Name is a textual name for the server configuration and must be unique. The name is used by the operator to uniquely identify a server class, and map pods back to an intended configuration.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"backup": {
						Description:         "Backup defines whether the Operator should manage automated backups, and how to lookup backup resources.",
						MarkdownDescription: "Backup defines whether the Operator should manage automated backups, and how to lookup backup resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"use_iam_role": {
								Description:         "UseIAMRole enables backup to fetch EC2 instance metadata. This allows the AWS SDK to use the EC2's IAM Role for S3 access. UseIAMRole will ignore credentials in s3Secret.",
								MarkdownDescription: "UseIAMRole enables backup to fetch EC2 instance metadata. This allows the AWS SDK to use the EC2's IAM Role for S3 access. UseIAMRole will ignore credentials in s3Secret.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_selector": {
								Description:         "NodeSelector defines which nodes to constrain the pods that run any backup and restore operations to.",
								MarkdownDescription: "NodeSelector defines which nodes to constrain the pods that run any backup and restore operations to.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"object_endpoint": {
								Description:         "ObjectEndpoint contains the configuration for connecting to a custom S3 compliant object store.",
								MarkdownDescription: "ObjectEndpoint contains the configuration for connecting to a custom S3 compliant object store.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"secret": {
										Description:         "The name of the secret, in this namespace, that contains the CA certificate for verification of a TLS endpoint (when required, e.g. not signed by a public CA). The secret must have the key with the name 'tls.crt'",
										MarkdownDescription: "The name of the secret, in this namespace, that contains the CA certificate for verification of a TLS endpoint (when required, e.g. not signed by a public CA). The secret must have the key with the name 'tls.crt'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "The host/address of the custom object endpoint.",
										MarkdownDescription: "The host/address of the custom object endpoint.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_virtual_path": {
										Description:         "UseVirtualPath will force the AWS SDK to use the new virtual style paths. by default alternative path style URLs which are often required by S3 compatible object stores.",
										MarkdownDescription: "UseVirtualPath will force the AWS SDK to use the new virtual style paths. by default alternative path style URLs which are often required by S3 compatible object stores.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_secret": {
								Description:         "S3Secret contains the region and credentials for operating backups in S3. This field must be popluated when the 'spec.s3bucket' field is specified for a backup or restore resource.",
								MarkdownDescription: "S3Secret contains the region and credentials for operating backups in S3. This field must be popluated when the 'spec.s3bucket' field is specified for a backup or restore resource.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": {
								Description:         "Selector allows CouchbaseBackup and CouchbaseBackupRestore resources to be filtered based on labels.",
								MarkdownDescription: "Selector allows CouchbaseBackup and CouchbaseBackupRestore resources to be filtered based on labels.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": {
								Description:         "Tolerations specifies all backup and restore pod tolerations.",
								MarkdownDescription: "Tolerations specifies all backup and restore pod tolerations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"effect": {
										Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
										MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
										MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operator": {
										Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
										MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"toleration_seconds": {
										Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
										MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
										MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "The Backup Image to run on backup pods.",
								MarkdownDescription: "The Backup Image to run on backup pods.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"image_pull_secrets": {
								Description:         "ImagePullSecrets allow you to use an image from private repositories and non-dockerhub ones.",
								MarkdownDescription: "ImagePullSecrets allow you to use an image from private repositories and non-dockerhub ones.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"managed": {
								Description:         "Managed defines whether backups are managed by us or the clients.",
								MarkdownDescription: "Managed defines whether backups are managed by us or the clients.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources is the resource requirements for the backup and restore containers.  Will be populated by defaults if not specified.",
								MarkdownDescription: "Resources is the resource requirements for the backup and restore containers.  Will be populated by defaults if not specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account_name": {
								Description:         "The Service Account to run backup (and restore) pods under. Without this backup pods will not be able to update status.",
								MarkdownDescription: "The Service Account to run backup (and restore) pods under. Without this backup pods will not be able to update status.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"upgrade_strategy": {
						Description:         "UpgradeStrategy controls how aggressive the Operator is when performing a cluster upgrade.  When a rolling upgrade is requested, pods are upgraded one at a time.  This strategy is slower, however less disruptive.  When an immediate upgrade strategy is requested, all pods are upgraded at the same time.  This strategy is faster, but more disruptive.  This field must be either 'RollingUpgrade' or 'ImmediateUpgrade', defaulting to 'RollingUpgrade'.",
						MarkdownDescription: "UpgradeStrategy controls how aggressive the Operator is when performing a cluster upgrade.  When a rolling upgrade is requested, pods are upgraded one at a time.  This strategy is slower, however less disruptive.  When an immediate upgrade strategy is requested, all pods are upgraded at the same time.  This strategy is faster, but more disruptive.  This field must be either 'RollingUpgrade' or 'ImmediateUpgrade', defaulting to 'RollingUpgrade'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"software_update_notifications": {
						Description:         "SoftwareUpdateNotifications enables software update notifications in the UI. When enabled, the UI will alert when a Couchbase server upgrade is available.",
						MarkdownDescription: "SoftwareUpdateNotifications enables software update notifications in the UI. When enabled, the UI will alert when a Couchbase server upgrade is available.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"autoscale_stabilization_period": {
						Description:         "AutoscaleStabilizationPeriod defines how long after a rebalance the corresponding HorizontalPodAutoscaler should remain in maintenance mode. During maintenance mode all autoscaling is disabled since every HorizontalPodAutoscaler associated with the cluster becomes inactive. Since certain metrics can be unpredictable when Couchbase is rebalancing or upgrading, setting a stabilization period helps to prevent scaling recommendations from the HorizontalPodAutoscaler for a provided period of time.  Values must be a valid Kubernetes duration of 0s or higher: https://golang.org/pkg/time/#ParseDuration A value of 0, puts the cluster in maintenance mode during rebalance but immediately exits this mode once the rebalance has completed. When undefined, the HPA is never put into maintenance mode during rebalance.",
						MarkdownDescription: "AutoscaleStabilizationPeriod defines how long after a rebalance the corresponding HorizontalPodAutoscaler should remain in maintenance mode. During maintenance mode all autoscaling is disabled since every HorizontalPodAutoscaler associated with the cluster becomes inactive. Since certain metrics can be unpredictable when Couchbase is rebalancing or upgrading, setting a stabilization period helps to prevent scaling recommendations from the HorizontalPodAutoscaler for a provided period of time.  Values must be a valid Kubernetes duration of 0s or higher: https://golang.org/pkg/time/#ParseDuration A value of 0, puts the cluster in maintenance mode during rebalance but immediately exits this mode once the rebalance has completed. When undefined, the HPA is never put into maintenance mode during rebalance.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_preview_scaling": {
						Description:         "DEPRECATED - This option only exists for backwards compatibility and no longer restricts autoscaling to ephemeral services. EnablePreviewScaling enables autoscaling for stateful services and buckets.",
						MarkdownDescription: "DEPRECATED - This option only exists for backwards compatibility and no longer restricts autoscaling to ephemeral services. EnablePreviewScaling enables autoscaling for stateful services and buckets.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Image is the container image name that will be used to launch Couchbase server instances.  Updating this field will cause an automatic upgrade of the cluster.",
						MarkdownDescription: "Image is the container image name that will be used to launch Couchbase server instances.  Updating this field will cause an automatic upgrade of the cluster.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"recovery_policy": {
						Description:         "RecoveryPolicy controls how aggressive the Operator is when recovering cluster topology.  When PrioritizeDataIntegrity, the Operator will delegate failover exclusively to Couchbase server, relying on it to only allow recovery when safe to do so.  When PrioritizeUptime, the Operator will wait for a period after the expected auto-failover of the cluster, before forcefully failing-over the pods. This may cause data loss, and is only expected to be used on clusters with ephemeral data, where the loss of the pod means that the data is known to be unrecoverable. This field must be either 'PrioritizeDataIntegrity' or 'PrioritizeUptime', defaulting to 'PrioritizeDataIntegrity'.",
						MarkdownDescription: "RecoveryPolicy controls how aggressive the Operator is when recovering cluster topology.  When PrioritizeDataIntegrity, the Operator will delegate failover exclusively to Couchbase server, relying on it to only allow recovery when safe to do so.  When PrioritizeUptime, the Operator will wait for a period after the expected auto-failover of the cluster, before forcefully failing-over the pods. This may cause data loss, and is only expected to be used on clusters with ephemeral data, where the loss of the pod means that the data is known to be unrecoverable. This field must be either 'PrioritizeDataIntegrity' or 'PrioritizeUptime', defaulting to 'PrioritizeDataIntegrity'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"security": {
						Description:         "Security defines Couchbase cluster security options such as the administrator account username and password, and user RBAC settings.",
						MarkdownDescription: "Security defines Couchbase cluster security options such as the administrator account username and password, and user RBAC settings.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ldap": {
								Description:         "LDAP Settings",
								MarkdownDescription: "LDAP Settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cacert": {
										Description:         "CA Certificate in PEM format to be used in LDAP server certificate validation",
										MarkdownDescription: "CA Certificate in PEM format to be used in LDAP server certificate validation",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"encryption": {
										Description:         "Encryption method to communicate with LDAP servers. Can be StartTLSExtension, TLS, or false.",
										MarkdownDescription: "Encryption method to communicate with LDAP servers. Can be StartTLSExtension, TLS, or false.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"hosts": {
										Description:         "List of LDAP hosts.",
										MarkdownDescription: "List of LDAP hosts.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"user_dn_mapping": {
										Description:         "User to distinguished name (DN) mapping. If none is specified, the username is used as the user’s distinguished name.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "User to distinguished name (DN) mapping. If none is specified, the username is used as the user’s distinguished name.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"query": {
												Description:         "Query is the LDAP query to run to map from Couchbase user to LDAP distinguished name.",
												MarkdownDescription: "Query is the LDAP query to run to map from Couchbase user to LDAP distinguished name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"template": {
												Description:         "This field specifies list of templates to use for providing username to DN mapping. The template may contain a placeholder specified as '%u' to represent the Couchbase user who is attempting to gain access.",
												MarkdownDescription: "This field specifies list of templates to use for providing username to DN mapping. The template may contain a placeholder specified as '%u' to represent the Couchbase user who is attempting to gain access.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authentication_enabled": {
										Description:         "Enables using LDAP to authenticate users.",
										MarkdownDescription: "Enables using LDAP to authenticate users.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"bind_secret": {
										Description:         "BindSecret is the name of a Kubernetes secret to use containing password for LDAP user binding",
										MarkdownDescription: "BindSecret is the name of a Kubernetes secret to use containing password for LDAP user binding",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"port": {
										Description:         "LDAP port. This is typically 389 for LDAP, and 636 for LDAPS.",
										MarkdownDescription: "LDAP port. This is typically 389 for LDAP, and 636 for LDAPS.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"tls_secret": {
										Description:         "TLSSecret is the name of a Kubernetes secret to use for LDAP ca cert.",
										MarkdownDescription: "TLSSecret is the name of a Kubernetes secret to use for LDAP ca cert.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"groups_query": {
										Description:         "LDAP query, to get the users' groups by username in RFC4516 format.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "LDAP query, to get the users' groups by username in RFC4516 format.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nested_groups_max_depth": {
										Description:         "Maximum number of recursive groups requests the server is allowed to perform. Requires NestedGroupsEnabled.  Values between 1 and 100: the default is 10. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "Maximum number of recursive groups requests the server is allowed to perform. Requires NestedGroupsEnabled.  Values between 1 and 100: the default is 10. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"bind_dn": {
										Description:         "DN to use for searching users and groups synchronization. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "DN to use for searching users and groups synchronization. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nested_groups_enabled": {
										Description:         "If enabled Couchbase server will try to recursively search for groups for every discovered ldap group. groups_query will be user for the search. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "If enabled Couchbase server will try to recursively search for groups for every discovered ldap group. groups_query will be user for the search. More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"server_cert_validation": {
										Description:         "Whether server certificate validation be enabled",
										MarkdownDescription: "Whether server certificate validation be enabled",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization_enabled": {
										Description:         "Enables use of LDAP groups for authorization.",
										MarkdownDescription: "Enables use of LDAP groups for authorization.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache_value_lifetime": {
										Description:         "Lifetime of values in cache in milliseconds. Default 300000 ms.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",
										MarkdownDescription: "Lifetime of values in cache in milliseconds. Default 300000 ms.  More info: https://docs.couchbase.com/server/current/manage/manage-security/configure-ldap.html",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rbac": {
								Description:         "Couchbase RBAC Users",
								MarkdownDescription: "Couchbase RBAC Users",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"selector": {
										Description:         "Selector is a label selector used to list RBAC resources in the namespace that are managed by the Operator.",
										MarkdownDescription: "Selector is a label selector used to list RBAC resources in the namespace that are managed by the Operator.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": {
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"managed": {
										Description:         "Managed defines whether RBAC is managed by us or the clients.",
										MarkdownDescription: "Managed defines whether RBAC is managed by us or the clients.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_secret": {
								Description:         "AdminSecret is the name of a Kubernetes secret to use for administrator authentication. The admin secret must contain the keys 'username' and 'password'.  The password data must be at least 6 characters in length, and not contain the any of the characters '()<>,;:'/[]?={}'.",
								MarkdownDescription: "AdminSecret is the name of a Kubernetes secret to use for administrator authentication. The admin secret must contain the keys 'username' and 'password'.  The password data must be at least 6 characters in length, and not contain the any of the characters '()<>,;:'/[]?={}'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"anti_affinity": {
						Description:         "AntiAffinity forces the Operator to schedule different Couchbase server pods on different Kubernetes nodes.  Anti-affinity reduces the likelihood of unrecoverable failure in the event of a node issue.  Use of anti-affinity is highly recommended for production clusters.",
						MarkdownDescription: "AntiAffinity forces the Operator to schedule different Couchbase server pods on different Kubernetes nodes.  Anti-affinity reduces the likelihood of unrecoverable failure in the event of a node issue.  Use of anti-affinity is highly recommended for production clusters.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"server_groups": {
						Description:         "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",
						MarkdownDescription: "ServerGroups define the set of availability zones you want to distribute pods over, and construct Couchbase server groups for.  By default, most cloud providers will label nodes with the key 'topology.kubernetes.io/zone', the values associated with that key are used here to provide explicit scheduling by the Operator.  You may manually label nodes using the 'topology.kubernetes.io/zone' key, to provide failure-domain aware scheduling when none is provided for you.  Global server groups are applied to all server classes, and may be overridden on a per-server class basis to give more control over scheduling and server groups.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"security_context": {
						Description:         "SecurityContext allows the configuration of the security context for all Couchbase server pods.  When using persistent volumes you may need to set the fsGroup field in order to write to the volume.  For non-root clusters you must also set runAsUser to 1000, corresponding to the Couchbase user in official container images.  More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
						MarkdownDescription: "SecurityContext allows the configuration of the security context for all Couchbase server pods.  When using persistent volumes you may need to set the fsGroup field in order to write to the volume.  For non-root clusters you must also set runAsUser to 1000, corresponding to the Couchbase user in official container images.  More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"supplemental_groups": {
								Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sysctls": {
								Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of a property to set",
										MarkdownDescription: "Name of a property to set",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value of a property to set",
										MarkdownDescription: "Value of a property to set",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fs_group": {
								Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"run_as_group": {
								Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"run_as_non_root": {
								Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
								MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"seccomp_profile": {
								Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"localhost_profile": {
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fs_group_change_policy": {
								Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"run_as_user": {
								Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"se_linux_options": {
								Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"user": {
										Description:         "User is a SELinux user label that applies to the container.",
										MarkdownDescription: "User is a SELinux user label that applies to the container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"level": {
										Description:         "Level is SELinux level label that applies to the container.",
										MarkdownDescription: "Level is SELinux level label that applies to the container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"role": {
										Description:         "Role is a SELinux role label that applies to the container.",
										MarkdownDescription: "Role is a SELinux role label that applies to the container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type is a SELinux type label that applies to the container.",
										MarkdownDescription: "Type is a SELinux type label that applies to the container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"windows_options": {
								Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
								MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host_process": {
										Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
										MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"run_as_user_name": {
										Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gmsa_credential_spec": {
										Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
										MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gmsa_credential_spec_name": {
										Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"buckets": {
						Description:         "Buckets defines whether the Operator should manage buckets, and how to lookup bucket resources.",
						MarkdownDescription: "Buckets defines whether the Operator should manage buckets, and how to lookup bucket resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"managed": {
								Description:         "Managed defines whether buckets are managed by the Operator (true), or user managed (false). When Operator managed, all buckets must be defined with either CouchbaseBucket, CouchbaseEphemeralBucket or CouchbaseMemcachedBucket resources.  Manual addition of buckets will be reverted by the Operator.  When user managed, the Operator will not interrogate buckets at all.  This field defaults to false.",
								MarkdownDescription: "Managed defines whether buckets are managed by the Operator (true), or user managed (false). When Operator managed, all buckets must be defined with either CouchbaseBucket, CouchbaseEphemeralBucket or CouchbaseMemcachedBucket resources.  Manual addition of buckets will be reverted by the Operator.  When user managed, the Operator will not interrogate buckets at all.  This field defaults to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": {
								Description:         "Selector is a label selector used to list buckets in the namespace that are managed by the Operator.",
								MarkdownDescription: "Selector is a label selector used to list buckets in the namespace that are managed by the Operator.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"synchronize": {
								Description:         "Synchronize allows unmanaged buckets, scopes, and collections to be synchronized as Kubernetes resources by the Operator.  This feature is intended for development only and should not be used for production workloads.  The synchronization workflow starts with 'spec.buckets.managed' being set to false, the user can manually create buckets, scopes, and collections using the Couchbase UI, or other tooling.  When you wish to commit to Kubernetes resources, you must specify a unique label selector in the 'spec.buckets.selector' field, and this field is set to true.  The Operator will create Kubernetes resources for you, and upon completion set the cluster's 'Synchronized' status condition.  You may then safely set 'spec.buckets.managed' to true and the Operator will manage these resources as per usual.  To update an already managed data topology, you must first set it to unmanaged, make any changes, and delete any old resources, then follow the standard synchronization workflow.  The Operator can not, and will not, ever delete, or make modifications to resource specifications that are intended to be user managed, or managed by a life cycle management tool. These actions must be instigated by an end user.  For a more complete experience, refer to the documentation for the 'cao save' and 'cao restore' CLI commands.",
								MarkdownDescription: "Synchronize allows unmanaged buckets, scopes, and collections to be synchronized as Kubernetes resources by the Operator.  This feature is intended for development only and should not be used for production workloads.  The synchronization workflow starts with 'spec.buckets.managed' being set to false, the user can manually create buckets, scopes, and collections using the Couchbase UI, or other tooling.  When you wish to commit to Kubernetes resources, you must specify a unique label selector in the 'spec.buckets.selector' field, and this field is set to true.  The Operator will create Kubernetes resources for you, and upon completion set the cluster's 'Synchronized' status condition.  You may then safely set 'spec.buckets.managed' to true and the Operator will manage these resources as per usual.  To update an already managed data topology, you must first set it to unmanaged, make any changes, and delete any old resources, then follow the standard synchronization workflow.  The Operator can not, and will not, ever delete, or make modifications to resource specifications that are intended to be user managed, or managed by a life cycle management tool. These actions must be instigated by an end user.  For a more complete experience, refer to the documentation for the 'cao save' and 'cao restore' CLI commands.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_online_volume_expansion": {
						Description:         "EnableOnlineVolumeExpansion enables online expansion of Persistent Volumes. You can only expand a PVC if its storage class's 'allowVolumeExpansion' field is set to true. Additionally, Kubernetes feature 'ExpandInUsePersistentVolumes' must be enabled in order to expand the volumes which are actively bound to Pods. Volumes can only be expanded and not reduced to a smaller size. See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#resizing-an-in-use-persistentvolumeclaim  If 'EnableOnlineVolumeExpansion' is enabled for use within an environment that does not actually support online volume and file system expansion then the cluster will fallback to rolling upgrade procedure to create a new set of Pods for use with resized Volumes. More info:  https://kubernetes.io/docs/concepts/storage/persistent-volumes/#expanding-persistent-volumes-claims",
						MarkdownDescription: "EnableOnlineVolumeExpansion enables online expansion of Persistent Volumes. You can only expand a PVC if its storage class's 'allowVolumeExpansion' field is set to true. Additionally, Kubernetes feature 'ExpandInUsePersistentVolumes' must be enabled in order to expand the volumes which are actively bound to Pods. Volumes can only be expanded and not reduced to a smaller size. See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#resizing-an-in-use-persistentvolumeclaim  If 'EnableOnlineVolumeExpansion' is enabled for use within an environment that does not actually support online volume and file system expansion then the cluster will fallback to rolling upgrade procedure to create a new set of Pods for use with resized Volumes. More info:  https://kubernetes.io/docs/concepts/storage/persistent-volumes/#expanding-persistent-volumes-claims",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"networking": {
						Description:         "Networking defines Couchbase cluster networking options such as network topology, TLS and DDNS settings.",
						MarkdownDescription: "Networking defines Couchbase cluster networking options such as network topology, TLS and DDNS settings.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"disable_ui_over_http": {
								Description:         "DisableUIOverHTTP is used to explicitly enable and disable UI access over the HTTP protocol.  If not specified, this field defaults to false.",
								MarkdownDescription: "DisableUIOverHTTP is used to explicitly enable and disable UI access over the HTTP protocol.  If not specified, this field defaults to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exposed_feature_service_template": {
								Description:         "ExposedFeatureServiceTemplate provides a template used by the Operator to create and manage per-pod services.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",
								MarkdownDescription: "ExposedFeatureServiceTemplate provides a template used by the Operator to create and manage per-pod services.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"labels": {
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"annotations": {
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": {
										Description:         "ServiceSpec describes the attributes that a user creates on a service.",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"session_affinity": {
												Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"health_check_node_port": {
												Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type).",
												MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_families": {
												Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer_class": {
												Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_traffic_policy": {
												Description:         "externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. 'Local' preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. 'Cluster' obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.",
												MarkdownDescription: "externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. 'Local' preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. 'Cluster' obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_ip": {
												Description:         "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_i_ps": {
												Description:         "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
												MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_name": {
												Description:         "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												MarkdownDescription: "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer_source_ranges": {
												Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"session_affinity_config": {
												Description:         "sessionAffinityConfig contains the configurations of session affinity.",
												MarkdownDescription: "sessionAffinityConfig contains the configurations of session affinity.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_ip": {
														Description:         "clientIP contains the configurations of Client IP based session affinity.",
														MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"timeout_seconds": {
																Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
																MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_i_ps": {
												Description:         "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"internal_traffic_policy": {
												Description:         "InternalTrafficPolicy specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. 'Cluster' routes internal traffic to a Service to all endpoints. 'Local' routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is 'Cluster'.",
												MarkdownDescription: "InternalTrafficPolicy specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. 'Cluster' routes internal traffic to a Service to all endpoints. 'Local' routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is 'Cluster'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allocate_load_balancer_node_ports": {
												Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type. This field is beta-level and is only honored by servers that enable the ServiceLBNodePortControl feature.",
												MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type. This field is beta-level and is only honored by servers that enable the ServiceLBNodePortControl feature.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer_ip": {
												Description:         "Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.",
												MarkdownDescription: "Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
												MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"network_platform": {
								Description:         "NetworkPlatform is used to enable support for various networking technologies.  This field must be one of 'Istio'.",
								MarkdownDescription: "NetworkPlatform is used to enable support for various networking technologies.  This field must be one of 'Istio'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_console_service_type": {
								Description:         "DEPRECATED - by adminConsoleServiceTemplate. AdminConsoleServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",
								MarkdownDescription: "DEPRECATED - by adminConsoleServiceTemplate. AdminConsoleServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exposed_feature_service_type": {
								Description:         "DEPRECATED - by exposedFeatureServiceTemplate. ExposedFeatureServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",
								MarkdownDescription: "DEPRECATED - by exposedFeatureServiceTemplate. ExposedFeatureServiceType defines whether to create a node port or load balancer service. When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. This field must be one of 'NodePort' or 'LoadBalancer', defaulting to 'NodePort'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exposed_feature_traffic_policy": {
								Description:         "DEPRECATED  - by exposedFeatureServiceTemplate. ExposedFeatureTrafficPolicy defines how packets should be routed from a load balancer service to a Couchbase pod.  When local, traffic is routed directly to the pod.  When cluster, traffic is routed to any node, then forwarded on.  While cluster routing may be slower, there are some situations where it is required for connectivity.  This field must be either 'Cluster' or 'Local', defaulting to 'Local',",
								MarkdownDescription: "DEPRECATED  - by exposedFeatureServiceTemplate. ExposedFeatureTrafficPolicy defines how packets should be routed from a load balancer service to a Couchbase pod.  When local, traffic is routed directly to the pod.  When cluster, traffic is routed to any node, then forwarded on.  While cluster routing may be slower, there are some situations where it is required for connectivity.  This field must be either 'Cluster' or 'Local', defaulting to 'Local',",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exposed_features": {
								Description:         "ExposedFeatures is a list of Couchbase features to expose when using a networking model that exposes the Couchbase cluster externally to Kubernetes.  This field also triggers the creation of per-pod services used by clients to connect to the Couchbase cluster.  When admin, only the administrator port is exposed, allowing remote administration.  When xdcr, only the services required for remote replication are exposed. The xdcr feature is only required when the cluster is the destination of an XDCR replication.  When client, all services are exposed as required for client SDK operation. This field may contain any of 'admin', 'xdcr' and 'client'.  Each feature may only be included once.",
								MarkdownDescription: "ExposedFeatures is a list of Couchbase features to expose when using a networking model that exposes the Couchbase cluster externally to Kubernetes.  This field also triggers the creation of per-pod services used by clients to connect to the Couchbase cluster.  When admin, only the administrator port is exposed, allowing remote administration.  When xdcr, only the services required for remote replication are exposed. The xdcr feature is only required when the cluster is the destination of an XDCR replication.  When client, all services are exposed as required for client SDK operation. This field may contain any of 'admin', 'xdcr' and 'client'.  Each feature may only be included once.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"load_balancer_source_ranges": {
								Description:         "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. LoadBalancerSourceRanges applies only when an exposed service is of type LoadBalancer and limits the source IP ranges that are allowed to use the service.  Items must use IPv4 class-less interdomain routing (CIDR) notation e.g. 10.0.0.0/16.",
								MarkdownDescription: "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. LoadBalancerSourceRanges applies only when an exposed service is of type LoadBalancer and limits the source IP ranges that are allowed to use the service.  Items must use IPv4 class-less interdomain routing (CIDR) notation e.g. 10.0.0.0/16.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "TLS defines the TLS configuration for the cluster including server and client certificate configuration, and TLS security policies.",
								MarkdownDescription: "TLS defines the TLS configuration for the cluster including server and client certificate configuration, and TLS security policies.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_certificate_paths": {
										Description:         "ClientCertificatePaths defines where to look in client certificates in order to extract the user name.",
										MarkdownDescription: "ClientCertificatePaths defines where to look in client certificates in order to extract the user name.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"delimiter": {
												Description:         "Delimiter if specified allows a suffix to be stripped from the username, once extracted from the certificate path.",
												MarkdownDescription: "Delimiter if specified allows a suffix to be stripped from the username, once extracted from the certificate path.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path defines where in the X.509 specification to extract the username from. This field must be either 'subject.cn', 'san.uri', 'san.dnsname' or  'san.email'.",
												MarkdownDescription: "Path defines where in the X.509 specification to extract the username from. This field must be either 'subject.cn', 'san.uri', 'san.dnsname' or  'san.email'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"prefix": {
												Description:         "Prefix allows a prefix to be stripped from the username, once extracted from the certificate path.",
												MarkdownDescription: "Prefix allows a prefix to be stripped from the username, once extracted from the certificate path.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_certificate_policy": {
										Description:         "ClientCertificatePolicy defines the client authentication policy to use. If set, the Operator expects TLS configuration to contain a valid certificate/key pair for the Administrator account.",
										MarkdownDescription: "ClientCertificatePolicy defines the client authentication policy to use. If set, the Operator expects TLS configuration to contain a valid certificate/key pair for the Administrator account.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_to_node_encryption": {
										Description:         "NodeToNodeEncryption specifies whether to encrypt data between Couchbase nodes within the same cluster.  This may come at the expense of performance.  When control plane only encryption is used, only cluster management traffic is encrypted between nodes.  When all, all traffic is encrypted, including database documents. When strict mode is used, it is the same as all, but also disables all plaintext ports.  Strict mode is only available on Couchbase Server versions 7.1 and greater. Node to node encryption can only be used when TLS certificates are managed by the Operator.  This field must be either 'ControlPlaneOnly', 'All', or 'Strict'.",
										MarkdownDescription: "NodeToNodeEncryption specifies whether to encrypt data between Couchbase nodes within the same cluster.  This may come at the expense of performance.  When control plane only encryption is used, only cluster management traffic is encrypted between nodes.  When all, all traffic is encrypted, including database documents. When strict mode is used, it is the same as all, but also disables all plaintext ports.  Strict mode is only available on Couchbase Server versions 7.1 and greater. Node to node encryption can only be used when TLS certificates are managed by the Operator.  This field must be either 'ControlPlaneOnly', 'All', or 'Strict'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_c_as": {
										Description:         "RootCAs defines a set of secrets that reside in this namespace that contain additional CA certificates that should be installed in Couchbase.  The CA certificates that are defined here are in addition to those defined for the cluster, optionally by couchbaseclusters.spec.networking.tls.secretSource, and thus should not be duplicated.  Secrets referred to must be of well-know type 'kubernetes.io/tls' and must contain the CA certificate under the key 'tls.crt'. Multiple root CA certificates are only supported on Couchbase Server 7.1 and greater, and not with legacy couchbaseclusters.spec.networking.tls.static configuration.",
										MarkdownDescription: "RootCAs defines a set of secrets that reside in this namespace that contain additional CA certificates that should be installed in Couchbase.  The CA certificates that are defined here are in addition to those defined for the cluster, optionally by couchbaseclusters.spec.networking.tls.secretSource, and thus should not be duplicated.  Secrets referred to must be of well-know type 'kubernetes.io/tls' and must contain the CA certificate under the key 'tls.crt'. Multiple root CA certificates are only supported on Couchbase Server 7.1 and greater, and not with legacy couchbaseclusters.spec.networking.tls.static configuration.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_source": {
										Description:         "SecretSource enables the user to specify a secret conforming to the Kubernetes TLS secret specification that is used for the Couchbase server certificate, and optionally the Operator's client certificate, providing cert-manager compatibility without having to specify a separate root CA.  A server CA certificate must be supplied by one of the provided methods. Certificates referred to must be of well-known type 'kubernetes.io/tls'.",
										MarkdownDescription: "SecretSource enables the user to specify a secret conforming to the Kubernetes TLS secret specification that is used for the Couchbase server certificate, and optionally the Operator's client certificate, providing cert-manager compatibility without having to specify a separate root CA.  A server CA certificate must be supplied by one of the provided methods. Certificates referred to must be of well-known type 'kubernetes.io/tls'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"client_secret_name": {
												Description:         "ClientSecretName specifies the secret name, in the same namespace as the cluster, the contains client TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the Kubernetes.io/tls secret type.",
												MarkdownDescription: "ClientSecretName specifies the secret name, in the same namespace as the cluster, the contains client TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the Kubernetes.io/tls secret type.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_secret_name": {
												Description:         "ServerSecretName specifies the secret name, in the same namespace as the cluster, that contains server TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the kubernetes.io/tls secret type.  It may also contain 'ca.crt'.",
												MarkdownDescription: "ServerSecretName specifies the secret name, in the same namespace as the cluster, that contains server TLS data.  The secret is expected to contain 'tls.crt' and 'tls.key' as per the kubernetes.io/tls secret type.  It may also contain 'ca.crt'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"static": {
										Description:         "DEPRECATED - by couchbaseclusters.spec.networking.tls.secretSource. Static enables user to generate static x509 certificates and keys, put them into Kubernetes secrets, and specify them here.  Static secrets are Couchbase specific, and follow no well-known standards.",
										MarkdownDescription: "DEPRECATED - by couchbaseclusters.spec.networking.tls.secretSource. Static enables user to generate static x509 certificates and keys, put them into Kubernetes secrets, and specify them here.  Static secrets are Couchbase specific, and follow no well-known standards.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"operator_secret": {
												Description:         "OperatorSecret is a secret name containing TLS certs used by operator to talk securely to this cluster.  The secret must contain a CA certificate (data key ca.crt).  If client authentication is enabled, then the secret must also contain a client certificate chain (data key 'couchbase-operator.crt') and private key (data key 'couchbase-operator.key').",
												MarkdownDescription: "OperatorSecret is a secret name containing TLS certs used by operator to talk securely to this cluster.  The secret must contain a CA certificate (data key ca.crt).  If client authentication is enabled, then the secret must also contain a client certificate chain (data key 'couchbase-operator.crt') and private key (data key 'couchbase-operator.key').",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_secret": {
												Description:         "ServerSecret is a secret name containing TLS certs used by each Couchbase member pod for the communication between Couchbase server and its clients.  The secret must contain a certificate chain (data key 'couchbase-operator.crt') and a private key (data key 'couchbase-operator.key').  The private key must be in the PKCS#1 RSA format.  The certificate chain must have a required set of X.509v3 subject alternative names for all cluster addressing modes.  See the Operator TLS documentation for more information.",
												MarkdownDescription: "ServerSecret is a secret name containing TLS certs used by each Couchbase member pod for the communication between Couchbase server and its clients.  The secret must contain a certificate chain (data key 'couchbase-operator.crt') and a private key (data key 'couchbase-operator.key').  The private key must be in the PKCS#1 RSA format.  The certificate chain must have a required set of X.509v3 subject alternative names for all cluster addressing modes.  See the Operator TLS documentation for more information.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_minimum_version": {
										Description:         "TLSMinimumVersion specifies the minimum TLS version the Couchbase server can negotiate with a client.  Must be one of TLS1.0, TLS1.1 TLS1.2 or TLS1.3, defaulting to TLS1.2.  TLS1.3 is only valid for Couchbase Server 7.1.0 onward.",
										MarkdownDescription: "TLSMinimumVersion specifies the minimum TLS version the Couchbase server can negotiate with a client.  Must be one of TLS1.0, TLS1.1 TLS1.2 or TLS1.3, defaulting to TLS1.2.  TLS1.3 is only valid for Couchbase Server 7.1.0 onward.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cipher_suites": {
										Description:         "CipherSuites specifies a list of cipher suites for Couchbase server to select from when negotiating TLS handshakes with a client.  Suites are not validated by the Operator.  Run 'openssl ciphers -v' in a Couchbase server pod to interrogate supported values.",
										MarkdownDescription: "CipherSuites specifies a list of cipher suites for Couchbase server to select from when negotiating TLS handshakes with a client.  Suites are not validated by the Operator.  Run 'openssl ciphers -v' in a Couchbase server pod to interrogate supported values.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wait_for_address_reachable_delay": {
								Description:         "WaitForAddressReachableDelay is used to defer operator checks that ensure external addresses are reachable before new nodes are balanced in to the cluster.  This prevents negative DNS caching while waiting for external-DDNS controllers to propagate addresses.",
								MarkdownDescription: "WaitForAddressReachableDelay is used to defer operator checks that ensure external addresses are reachable before new nodes are balanced in to the cluster.  This prevents negative DNS caching while waiting for external-DDNS controllers to propagate addresses.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_ui_over_https": {
								Description:         "DisableUIOverHTTPS is used to explicitly enable and disable UI access over the HTTPS protocol.  If not specified, this field defaults to false.",
								MarkdownDescription: "DisableUIOverHTTPS is used to explicitly enable and disable UI access over the HTTPS protocol.  If not specified, this field defaults to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_annotations": {
								Description:         "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. ServiceAnnotations allows services to be annotated with custom labels. Operator annotations are merged on top of these so have precedence as they are required for correct operation.",
								MarkdownDescription: "DEPRECATED - by adminConsoleServiceTemplate and exposedFeatureServiceTemplate. ServiceAnnotations allows services to be annotated with custom labels. Operator annotations are merged on top of these so have precedence as they are required for correct operation.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wait_for_address_reachable": {
								Description:         "WaitForAddressReachable is used to set the timeout between when polling of external addresses is started, and when it is deemed a failure.  Polling of DNS name availability inherently dangerous due to negative caching, so prefer the use of an initial 'waitForAddressReachableDelay' to allow propagation.",
								MarkdownDescription: "WaitForAddressReachable is used to set the timeout between when polling of external addresses is started, and when it is deemed a failure.  Polling of DNS name availability inherently dangerous due to negative caching, so prefer the use of an initial 'waitForAddressReachableDelay' to allow propagation.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_console_service_template": {
								Description:         "AdminConsoleServiceTemplate provides a template used by the Operator to create and manage the admin console service.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",
								MarkdownDescription: "AdminConsoleServiceTemplate provides a template used by the Operator to create and manage the admin console service.  This allows services to be annotated, the service type defined and any other options that Kubernetes provides.  When using a LoadBalancer service type, TLS and dynamic DNS must also be enabled. The Operator reserves the right to modify or replace any field.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#service-v1-core",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
										MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": {
										Description:         "ServiceSpec describes the attributes that a user creates on a service.",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allocate_load_balancer_node_ports": {
												Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type. This field is beta-level and is only honored by servers that enable the ServiceLBNodePortControl feature.",
												MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer.  Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts.  If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type. This field is beta-level and is only honored by servers that enable the ServiceLBNodePortControl feature.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_i_ps": {
												Description:         "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
												MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_traffic_policy": {
												Description:         "externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. 'Local' preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. 'Cluster' obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.",
												MarkdownDescription: "externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. 'Local' preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. 'Cluster' obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer_ip": {
												Description:         "Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.",
												MarkdownDescription: "Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"session_affinity_config": {
												Description:         "sessionAffinityConfig contains the configurations of session affinity.",
												MarkdownDescription: "sessionAffinityConfig contains the configurations of session affinity.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_ip": {
														Description:         "clientIP contains the configurations of Client IP based session affinity.",
														MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"timeout_seconds": {
																Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
																MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_ip": {
												Description:         "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_name": {
												Description:         "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												MarkdownDescription: "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved.  Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"internal_traffic_policy": {
												Description:         "InternalTrafficPolicy specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. 'Cluster' routes internal traffic to a Service to all endpoints. 'Local' routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is 'Cluster'.",
												MarkdownDescription: "InternalTrafficPolicy specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only. 'Cluster' routes internal traffic to a Service to all endpoints. 'Local' routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is 'Cluster'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer_class": {
												Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer_source_ranges": {
												Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_i_ps": {
												Description:         "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly.  If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above).  Valid values are 'None', empty string (''), or a valid IP address.  Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required.  Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName.  If this field is not specified, it will be initialized from the clusterIP field.  If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.  This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_families": {
												Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'.  This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName.  This field may hold a maximum of two entries (dual-stack families, in either order).  These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"session_affinity": {
												Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
												MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"health_check_node_port": {
												Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type).",
												MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used.  If not specified, a value will be automatically allocated.  External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_console_services": {
								Description:         "DEPRECATED - not required by Couchbase Server 6.5.0 onward. AdminConsoleServices is a selector to choose specific services to expose via the admin console. This field may contain any of 'data', 'index', 'query', 'search', 'eventing' and 'analytics'.  Each service may only be included once.",
								MarkdownDescription: "DEPRECATED - not required by Couchbase Server 6.5.0 onward. AdminConsoleServices is a selector to choose specific services to expose via the admin console. This field may contain any of 'data', 'index', 'query', 'search', 'eventing' and 'analytics'.  Each service may only be included once.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns": {
								Description:         "DNS defines information required for Dynamic DNS support.",
								MarkdownDescription: "DNS defines information required for Dynamic DNS support.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"domain": {
										Description:         "Domain is the domain to create pods in.  When populated the Operator will annotate the admin console and per-pod services with the key 'external-dns.alpha.kubernetes.io/hostname'.  These annotations can be used directly by a Kubernetes External-DNS controller to replicate load balancer service IP addresses into a public DNS server.",
										MarkdownDescription: "Domain is the domain to create pods in.  When populated the Operator will annotate the admin console and per-pod services with the key 'external-dns.alpha.kubernetes.io/hostname'.  These annotations can be used directly by a Kubernetes External-DNS controller to replicate load balancer service IP addresses into a public DNS server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"expose_admin_console": {
								Description:         "ExposeAdminConsole creates a service referencing the admin console. The service is configured by the adminConsoleServiceTemplate field.",
								MarkdownDescription: "ExposeAdminConsole creates a service referencing the admin console. The service is configured by the adminConsoleServiceTemplate field.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"address_family": {
								Description:         "AddressFamily allows the manual selection of the address family to use. When this field is not set, Couchbase server will default to using IPv4 for internal communication and also support IPv6 on dual stack systems. Setting this field to either IPv4 or IPv6 will force Couchbase to use the selected protocol for internal communication, and also disable all other protocols to provide added security and simplicty when defining firewall rules.  Disabling of address families is only supported in Couchbase Server 7.0.2+.",
								MarkdownDescription: "AddressFamily allows the manual selection of the address family to use. When this field is not set, Couchbase server will default to using IPv4 for internal communication and also support IPv6 on dual stack systems. Setting this field to either IPv4 or IPv6 will force Couchbase to use the selected protocol for internal communication, and also disable all other protocols to provide added security and simplicty when defining firewall rules.  Disabling of address families is only supported in Couchbase Server 7.0.2+.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": {
						Description:         "Paused is to pause the control of the operator for the Couchbase cluster. This does not pause the cluster itself, instead stopping the operator from taking any action.",
						MarkdownDescription: "Paused is to pause the control of the operator for the Couchbase cluster. This does not pause the cluster itself, instead stopping the operator from taking any action.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"platform": {
						Description:         "Platform gives a hint as to what platform we are running on and how to configure services.  This field must be one of 'aws', 'gke' or 'azure'.",
						MarkdownDescription: "Platform gives a hint as to what platform we are running on and how to configure services.  This field must be one of 'aws', 'gke' or 'azure'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_claim_templates": {
						Description:         "VolumeClaimTemplates define the desired characteristics of a volume that can be requested/claimed by a pod, for example the storage class to use and the volume size.  Volume claim templates are referred to by name by server class volume mount configuration.",
						MarkdownDescription: "VolumeClaimTemplates define the desired characteristics of a volume that can be requested/claimed by a pod, for example the storage class to use and the volume size.  Volume claim templates are referred to by name by server class volume mount configuration.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",
								MarkdownDescription: "Standard objects metadata.  This is a curated version for use with Couchbase resource templates.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"spec": {
								Description:         "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",
								MarkdownDescription: "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"data_source_ref": {
										Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api_group": {
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "A label query over volumes to consider for binding.",
										MarkdownDescription: "A label query over volumes to consider for binding.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_labels": {
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_expressions": {
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_class_name": {
										Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_mode": {
										Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
										MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_name": {
										Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
										MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"access_modes": {
										Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"xdcr": {
						Description:         "XDCR defines whether the Operator should manage XDCR, remote clusters and how to lookup replication resources.",
						MarkdownDescription: "XDCR defines whether the Operator should manage XDCR, remote clusters and how to lookup replication resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"remote_clusters": {
								Description:         "RemoteClusters is a set of named remote clusters to establish replications to.",
								MarkdownDescription: "RemoteClusters is a set of named remote clusters to establish replications to.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"authentication_secret": {
										Description:         "AuthenticationSecret is a secret used to authenticate when establishing a remote connection.  It is only required when not using mTLS.  The secret must contain a username (secret key 'username') and password (secret key 'password').",
										MarkdownDescription: "AuthenticationSecret is a secret used to authenticate when establishing a remote connection.  It is only required when not using mTLS.  The secret must contain a username (secret key 'username') and password (secret key 'password').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"hostname": {
										Description:         "Hostname is the connection string to use to connect the remote cluster.",
										MarkdownDescription: "Hostname is the connection string to use to connect the remote cluster.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the remote cluster.",
										MarkdownDescription: "Name of the remote cluster.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"replications": {
										Description:         "Replications are replication streams from this cluster to the remote one. This field defines how to look up CouchbaseReplication resources.  By default any CouchbaseReplication resources in the namespace will be considered.",
										MarkdownDescription: "Replications are replication streams from this cluster to the remote one. This field defines how to look up CouchbaseReplication resources.  By default any CouchbaseReplication resources in the namespace will be considered.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"selector": {
												Description:         "Selector allows CouchbaseReplication resources to be filtered based on labels.",
												MarkdownDescription: "Selector allows CouchbaseReplication resources to be filtered based on labels.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS if specified references a resource containing the necessary certificate data for an encrypted connection.",
										MarkdownDescription: "TLS if specified references a resource containing the necessary certificate data for an encrypted connection.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret": {
												Description:         "Secret references a secret containing the CA certificate (data key 'ca'), and optionally a client certificate (data key 'certificate') and key (data key 'key').",
												MarkdownDescription: "Secret references a secret containing the CA certificate (data key 'ca'), and optionally a client certificate (data key 'certificate') and key (data key 'key').",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uuid": {
										Description:         "UUID of the remote cluster.  The UUID of a CouchbaseCluster resource is advertised in the status.clusterId field of the resource.",
										MarkdownDescription: "UUID of the remote cluster.  The UUID of a CouchbaseCluster resource is advertised in the status.clusterId field of the resource.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"managed": {
								Description:         "Managed defines whether XDCR is managed by the operator or not.",
								MarkdownDescription: "Managed defines whether XDCR is managed by the operator or not.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auto_resource_allocation": {
						Description:         "AutoResourceAllocation populates pod resource requests based on the services running on that pod.  When enabled, this feature will calculate the memory request as the total of service allocations defined in 'spec.cluster', plus an overhead defined by 'spec.autoResourceAllocation.overheadPercent'.Changing individual allocations for a service will cause a cluster upgrade as allocations are modified in the underlying pods.  This field also allows default pod CPU requests and limits to be applied. All resource allocations can be overridden by explicitly configuring them in the 'spec.servers.resources' field.",
						MarkdownDescription: "AutoResourceAllocation populates pod resource requests based on the services running on that pod.  When enabled, this feature will calculate the memory request as the total of service allocations defined in 'spec.cluster', plus an overhead defined by 'spec.autoResourceAllocation.overheadPercent'.Changing individual allocations for a service will cause a cluster upgrade as allocations are modified in the underlying pods.  This field also allows default pod CPU requests and limits to be applied. All resource allocations can be overridden by explicitly configuring them in the 'spec.servers.resources' field.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cpu_limits": {
								Description:         "CPULimits automatically populates the CPU limits across all Couchbase server pods.  This field defaults to '4' CPUs.  Explicitly specifying the CPU limit for a particular server class will override this value.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "CPULimits automatically populates the CPU limits across all Couchbase server pods.  This field defaults to '4' CPUs.  Explicitly specifying the CPU limit for a particular server class will override this value.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu_requests": {
								Description:         "CPURequests automatically populates the CPU requests across all Couchbase server pods.  The default value of '2', is the minimum recommended number of CPUs required to run Couchbase Server.  Explicitly specifying the CPU request for a particular server class will override this value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "CPURequests automatically populates the CPU requests across all Couchbase server pods.  The default value of '2', is the minimum recommended number of CPUs required to run Couchbase Server.  Explicitly specifying the CPU request for a particular server class will override this value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "Enabled defines whether auto-resource allocation is enabled.",
								MarkdownDescription: "Enabled defines whether auto-resource allocation is enabled.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"overhead_percent": {
								Description:         "OverheadPercent defines the amount of memory above that required for individual services on a pod.  For Couchbase Server this should be approximately 25%.",
								MarkdownDescription: "OverheadPercent defines the amount of memory above that required for individual services on a pod.  For Couchbase Server this should be approximately 25%.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_cluster_v2")

	var state CouchbaseComCouchbaseClusterV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseClusterV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseCluster")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_cluster_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_cluster_v2")

	var state CouchbaseComCouchbaseClusterV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseClusterV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseCluster")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CouchbaseComCouchbaseClusterV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_cluster_v2")
	// NO-OP: Terraform removes the state automatically for us
}
