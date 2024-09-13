/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tf_tungsten_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TfTungstenIoManagerV1Alpha1Manifest{}
)

func NewTfTungstenIoManagerV1Alpha1Manifest() datasource.DataSource {
	return &TfTungstenIoManagerV1Alpha1Manifest{}
}

type TfTungstenIoManagerV1Alpha1Manifest struct{}

type TfTungstenIoManagerV1Alpha1ManifestData struct {
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
		CommonConfiguration *struct {
			AuthParameters *struct {
				AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
				KeystoneAuthParameters *struct {
					Address           *string `tfsdk:"address" json:"address,omitempty"`
					AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
					AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
					AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
					AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
					AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
					Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
					Region            *string `tfsdk:"region" json:"region,omitempty"`
					UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
				} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
				KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
			} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
			CertKeyLength    *int64             `tfsdk:"cert_key_length" json:"certKeyLength,omitempty"`
			CertSigner       *string            `tfsdk:"cert_signer" json:"certSigner,omitempty"`
			Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
			ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
			NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations      *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
		Services *struct {
			Analytics *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AaaMode                 *string `tfsdk:"aaa_mode" json:"aaaMode,omitempty"`
						AnalyticsConfigAuditTTL *int64  `tfsdk:"analytics_config_audit_ttl" json:"analyticsConfigAuditTTL,omitempty"`
						AnalyticsDataTTL        *int64  `tfsdk:"analytics_data_ttl" json:"analyticsDataTTL,omitempty"`
						AnalyticsFlowTTL        *int64  `tfsdk:"analytics_flow_ttl" json:"analyticsFlowTTL,omitempty"`
						AnalyticsIntrospectPort *int64  `tfsdk:"analytics_introspect_port" json:"analyticsIntrospectPort,omitempty"`
						AnalyticsPort           *int64  `tfsdk:"analytics_port" json:"analyticsPort,omitempty"`
						AnalyticsStatisticsTTL  *int64  `tfsdk:"analytics_statistics_ttl" json:"analyticsStatisticsTTL,omitempty"`
						CollectorIntrospectPort *int64  `tfsdk:"collector_introspect_port" json:"collectorIntrospectPort,omitempty"`
						CollectorPort           *int64  `tfsdk:"collector_port" json:"collectorPort,omitempty"`
						Containers              *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"analytics" json:"analytics,omitempty"`
			AnalyticsAlarm *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AlarmgenIntrospectListenPort   *int64  `tfsdk:"alarmgen_introspect_listen_port" json:"alarmgenIntrospectListenPort,omitempty"`
						AlarmgenLogFileName            *string `tfsdk:"alarmgen_log_file_name" json:"alarmgenLogFileName,omitempty"`
						AlarmgenPartitions             *int64  `tfsdk:"alarmgen_partitions" json:"alarmgenPartitions,omitempty"`
						AlarmgenRedisAggregateDbOffset *int64  `tfsdk:"alarmgen_redis_aggregate_db_offset" json:"alarmgenRedisAggregateDbOffset,omitempty"`
						Containers                     *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						LogFilePath *string `tfsdk:"log_file_path" json:"logFilePath,omitempty"`
						LogLocal    *string `tfsdk:"log_local" json:"logLocal,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"analytics_alarm" json:"analyticsAlarm,omitempty"`
			AnalyticsSnmp *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						Containers *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						LogFilePath                       *string `tfsdk:"log_file_path" json:"logFilePath,omitempty"`
						LogLocal                          *string `tfsdk:"log_local" json:"logLocal,omitempty"`
						SnmpCollectorFastScanFrequency    *int64  `tfsdk:"snmp_collector_fast_scan_frequency" json:"snmpCollectorFastScanFrequency,omitempty"`
						SnmpCollectorIntrospectListenPort *int64  `tfsdk:"snmp_collector_introspect_listen_port" json:"snmpCollectorIntrospectListenPort,omitempty"`
						SnmpCollectorLogFileName          *string `tfsdk:"snmp_collector_log_file_name" json:"snmpCollectorLogFileName,omitempty"`
						SnmpCollectorScanFrequency        *int64  `tfsdk:"snmp_collector_scan_frequency" json:"snmpCollectorScanFrequency,omitempty"`
						TopologyIntrospectListenPort      *int64  `tfsdk:"topology_introspect_listen_port" json:"topologyIntrospectListenPort,omitempty"`
						TopologyLogFileName               *string `tfsdk:"topology_log_file_name" json:"topologyLogFileName,omitempty"`
						TopologySnmpFrequency             *int64  `tfsdk:"topology_snmp_frequency" json:"topologySnmpFrequency,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"analytics_snmp" json:"analyticsSnmp,omitempty"`
			Cassandras *[]struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						CassandraParameters *struct {
							CompactionThroughputMbPerSec     *int64  `tfsdk:"compaction_throughput_mb_per_sec" json:"compactionThroughputMbPerSec,omitempty"`
							ConcurrentCompactors             *int64  `tfsdk:"concurrent_compactors" json:"concurrentCompactors,omitempty"`
							ConcurrentCounterWrites          *int64  `tfsdk:"concurrent_counter_writes" json:"concurrentCounterWrites,omitempty"`
							ConcurrentMaterializedViewWrites *int64  `tfsdk:"concurrent_materialized_view_writes" json:"concurrentMaterializedViewWrites,omitempty"`
							ConcurrentReads                  *int64  `tfsdk:"concurrent_reads" json:"concurrentReads,omitempty"`
							ConcurrentWrites                 *int64  `tfsdk:"concurrent_writes" json:"concurrentWrites,omitempty"`
							MemtableAllocationType           *string `tfsdk:"memtable_allocation_type" json:"memtableAllocationType,omitempty"`
							MemtableFlushWriters             *int64  `tfsdk:"memtable_flush_writers" json:"memtableFlushWriters,omitempty"`
						} `tfsdk:"cassandra_parameters" json:"cassandraParameters,omitempty"`
						Containers *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						CqlPort        *int64  `tfsdk:"cql_port" json:"cqlPort,omitempty"`
						JmxLocalPort   *int64  `tfsdk:"jmx_local_port" json:"jmxLocalPort,omitempty"`
						ListenAddress  *string `tfsdk:"listen_address" json:"listenAddress,omitempty"`
						MaxHeapSize    *string `tfsdk:"max_heap_size" json:"maxHeapSize,omitempty"`
						MinHeapSize    *string `tfsdk:"min_heap_size" json:"minHeapSize,omitempty"`
						MinimumDiskGB  *int64  `tfsdk:"minimum_disk_gb" json:"minimumDiskGB,omitempty"`
						Port           *int64  `tfsdk:"port" json:"port,omitempty"`
						ReaperAdmPort  *int64  `tfsdk:"reaper_adm_port" json:"reaperAdmPort,omitempty"`
						ReaperAppPort  *int64  `tfsdk:"reaper_app_port" json:"reaperAppPort,omitempty"`
						ReaperEnabled  *bool   `tfsdk:"reaper_enabled" json:"reaperEnabled,omitempty"`
						SslStoragePort *int64  `tfsdk:"ssl_storage_port" json:"sslStoragePort,omitempty"`
						StartRPC       *bool   `tfsdk:"start_rpc" json:"startRPC,omitempty"`
						StoragePort    *int64  `tfsdk:"storage_port" json:"storagePort,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"cassandras" json:"cassandras,omitempty"`
			Config *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AaaMode           *string `tfsdk:"aaa_mode" json:"aaaMode,omitempty"`
						ApiAdminPort      *int64  `tfsdk:"api_admin_port" json:"apiAdminPort,omitempty"`
						ApiIntrospectPort *int64  `tfsdk:"api_introspect_port" json:"apiIntrospectPort,omitempty"`
						ApiPort           *int64  `tfsdk:"api_port" json:"apiPort,omitempty"`
						ApiWorkerCount    *int64  `tfsdk:"api_worker_count" json:"apiWorkerCount,omitempty"`
						BgpAutoMesh       *bool   `tfsdk:"bgp_auto_mesh" json:"bgpAutoMesh,omitempty"`
						BgpEnable4Byte    *bool   `tfsdk:"bgp_enable4_byte" json:"bgpEnable4Byte,omitempty"`
						Containers        *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						DeviceManagerIntrospectPort *int64  `tfsdk:"device_manager_introspect_port" json:"deviceManagerIntrospectPort,omitempty"`
						FabricMgmtIP                *string `tfsdk:"fabric_mgmt_ip" json:"fabricMgmtIP,omitempty"`
						GlobalASNNumber             *int64  `tfsdk:"global_asn_number" json:"globalASNNumber,omitempty"`
						LinklocalServiceConfig      *struct {
							Ip                  *string `tfsdk:"ip" json:"ip,omitempty"`
							IpFabricServiceHost *string `tfsdk:"ip_fabric_service_host" json:"ipFabricServiceHost,omitempty"`
							IpFabricServicePort *int64  `tfsdk:"ip_fabric_service_port" json:"ipFabricServicePort,omitempty"`
							Name                *string `tfsdk:"name" json:"name,omitempty"`
							Port                *int64  `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"linklocal_service_config" json:"linklocalServiceConfig,omitempty"`
						SchemaIntrospectPort     *int64 `tfsdk:"schema_introspect_port" json:"schemaIntrospectPort,omitempty"`
						SvcMonitorIntrospectPort *int64 `tfsdk:"svc_monitor_introspect_port" json:"svcMonitorIntrospectPort,omitempty"`
						UseExternalTFTP          *bool  `tfsdk:"use_external_tftp" json:"useExternalTFTP,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			Controls *[]struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AsnNumber  *int64 `tfsdk:"asn_number" json:"asnNumber,omitempty"`
						BgpPort    *int64 `tfsdk:"bgp_port" json:"bgpPort,omitempty"`
						Containers *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						DataSubnet        *string `tfsdk:"data_subnet" json:"dataSubnet,omitempty"`
						DnsIntrospectPort *int64  `tfsdk:"dns_introspect_port" json:"dnsIntrospectPort,omitempty"`
						DnsPort           *int64  `tfsdk:"dns_port" json:"dnsPort,omitempty"`
						Rndckey           *string `tfsdk:"rndckey" json:"rndckey,omitempty"`
						Subcluster        *string `tfsdk:"subcluster" json:"subcluster,omitempty"`
						XmppPort          *int64  `tfsdk:"xmpp_port" json:"xmppPort,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"controls" json:"controls,omitempty"`
			Kubemanager *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						CloudOrchestrator *string `tfsdk:"cloud_orchestrator" json:"cloudOrchestrator,omitempty"`
						Containers        *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						HostNetworkService   *bool   `tfsdk:"host_network_service" json:"hostNetworkService,omitempty"`
						IpFabricForwarding   *bool   `tfsdk:"ip_fabric_forwarding" json:"ipFabricForwarding,omitempty"`
						IpFabricSnat         *bool   `tfsdk:"ip_fabric_snat" json:"ipFabricSnat,omitempty"`
						IpFabricSubnets      *string `tfsdk:"ip_fabric_subnets" json:"ipFabricSubnets,omitempty"`
						KubernetesAPIPort    *int64  `tfsdk:"kubernetes_api_port" json:"kubernetesAPIPort,omitempty"`
						KubernetesAPISSLPort *int64  `tfsdk:"kubernetes_apissl_port" json:"kubernetesAPISSLPort,omitempty"`
						KubernetesAPIServer  *string `tfsdk:"kubernetes_api_server" json:"kubernetesAPIServer,omitempty"`
						KubernetesTokenFile  *string `tfsdk:"kubernetes_token_file" json:"kubernetesTokenFile,omitempty"`
						PodSubnet            *string `tfsdk:"pod_subnet" json:"podSubnet,omitempty"`
						PublicFIPPool        *string `tfsdk:"public_fip_pool" json:"publicFIPPool,omitempty"`
						ServiceSubnet        *string `tfsdk:"service_subnet" json:"serviceSubnet,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"kubemanager" json:"kubemanager,omitempty"`
			Queryengine *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AnalyticsdbIntrospectPort *int64 `tfsdk:"analyticsdb_introspect_port" json:"analyticsdbIntrospectPort,omitempty"`
						AnalyticsdbPort           *int64 `tfsdk:"analyticsdb_port" json:"analyticsdbPort,omitempty"`
						Containers                *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"queryengine" json:"queryengine,omitempty"`
			Rabbitmq *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						ClusterPartitionHandling *string `tfsdk:"cluster_partition_handling" json:"clusterPartitionHandling,omitempty"`
						Containers               *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						ErlEpmdPort       *int64  `tfsdk:"erl_epmd_port" json:"erlEpmdPort,omitempty"`
						ErlangCookie      *string `tfsdk:"erlang_cookie" json:"erlangCookie,omitempty"`
						MirroredQueueMode *string `tfsdk:"mirrored_queue_mode" json:"mirroredQueueMode,omitempty"`
						Password          *string `tfsdk:"password" json:"password,omitempty"`
						Port              *int64  `tfsdk:"port" json:"port,omitempty"`
						Secret            *string `tfsdk:"secret" json:"secret,omitempty"`
						TcpListenOptions  *struct {
							Backlog       *int64 `tfsdk:"backlog" json:"backlog,omitempty"`
							ExitOnClose   *bool  `tfsdk:"exit_on_close" json:"exitOnClose,omitempty"`
							LingerOn      *bool  `tfsdk:"linger_on" json:"lingerOn,omitempty"`
							LingerTimeout *int64 `tfsdk:"linger_timeout" json:"lingerTimeout,omitempty"`
							Nodelay       *bool  `tfsdk:"nodelay" json:"nodelay,omitempty"`
						} `tfsdk:"tcp_listen_options" json:"tcpListenOptions,omitempty"`
						User  *string `tfsdk:"user" json:"user,omitempty"`
						Vhost *string `tfsdk:"vhost" json:"vhost,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"rabbitmq" json:"rabbitmq,omitempty"`
			Redis *[]struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
						Containers  *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						RedisPort *int64 `tfsdk:"redis_port" json:"redisPort,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"redis" json:"redis,omitempty"`
			Vrouters *[]struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AgentMode                 *string `tfsdk:"agent_mode" json:"agentMode,omitempty"`
						BarbicanPassword          *string `tfsdk:"barbican_password" json:"barbicanPassword,omitempty"`
						BarbicanTenantName        *string `tfsdk:"barbican_tenant_name" json:"barbicanTenantName,omitempty"`
						BarbicanUser              *string `tfsdk:"barbican_user" json:"barbicanUser,omitempty"`
						CloudOrchestrator         *string `tfsdk:"cloud_orchestrator" json:"cloudOrchestrator,omitempty"`
						CniMTU                    *int64  `tfsdk:"cni_mtu" json:"cniMTU,omitempty"`
						CollectorPort             *string `tfsdk:"collector_port" json:"collectorPort,omitempty"`
						ConfigApiPort             *string `tfsdk:"config_api_port" json:"configApiPort,omitempty"`
						ConfigApiServerCaCertfile *string `tfsdk:"config_api_server_ca_certfile" json:"configApiServerCaCertfile,omitempty"`
						ConfigApiSslEnable        *bool   `tfsdk:"config_api_ssl_enable" json:"configApiSslEnable,omitempty"`
						Containers                *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						ControlInstance                 *string            `tfsdk:"control_instance" json:"controlInstance,omitempty"`
						DataSubnet                      *string            `tfsdk:"data_subnet" json:"dataSubnet,omitempty"`
						DnsServerPort                   *string            `tfsdk:"dns_server_port" json:"dnsServerPort,omitempty"`
						DpdkUioDriver                   *string            `tfsdk:"dpdk_uio_driver" json:"dpdkUioDriver,omitempty"`
						EnvVariablesConfig              *map[string]string `tfsdk:"env_variables_config" json:"envVariablesConfig,omitempty"`
						FabricSntHashTableSize          *string            `tfsdk:"fabric_snt_hash_table_size" json:"fabricSntHashTableSize,omitempty"`
						HugePages1G                     *int64             `tfsdk:"huge_pages1_g" json:"hugePages1G,omitempty"`
						HugePages2M                     *int64             `tfsdk:"huge_pages2_m" json:"hugePages2M,omitempty"`
						HypervisorType                  *string            `tfsdk:"hypervisor_type" json:"hypervisorType,omitempty"`
						IntrospectSslEnable             *bool              `tfsdk:"introspect_ssl_enable" json:"introspectSslEnable,omitempty"`
						K8sToken                        *string            `tfsdk:"k8s_token" json:"k8sToken,omitempty"`
						K8sTokenFile                    *string            `tfsdk:"k8s_token_file" json:"k8sTokenFile,omitempty"`
						KeystoneAuthAdminPassword       *string            `tfsdk:"keystone_auth_admin_password" json:"keystoneAuthAdminPassword,omitempty"`
						KeystoneAuthAdminPort           *string            `tfsdk:"keystone_auth_admin_port" json:"keystoneAuthAdminPort,omitempty"`
						KeystoneAuthCaCertfile          *string            `tfsdk:"keystone_auth_ca_certfile" json:"keystoneAuthCaCertfile,omitempty"`
						KeystoneAuthCertfile            *string            `tfsdk:"keystone_auth_certfile" json:"keystoneAuthCertfile,omitempty"`
						KeystoneAuthHost                *string            `tfsdk:"keystone_auth_host" json:"keystoneAuthHost,omitempty"`
						KeystoneAuthInsecure            *bool              `tfsdk:"keystone_auth_insecure" json:"keystoneAuthInsecure,omitempty"`
						KeystoneAuthKeyfile             *string            `tfsdk:"keystone_auth_keyfile" json:"keystoneAuthKeyfile,omitempty"`
						KeystoneAuthProjectDomainName   *string            `tfsdk:"keystone_auth_project_domain_name" json:"keystoneAuthProjectDomainName,omitempty"`
						KeystoneAuthProto               *string            `tfsdk:"keystone_auth_proto" json:"keystoneAuthProto,omitempty"`
						KeystoneAuthRegionName          *string            `tfsdk:"keystone_auth_region_name" json:"keystoneAuthRegionName,omitempty"`
						KeystoneAuthUrlTokens           *string            `tfsdk:"keystone_auth_url_tokens" json:"keystoneAuthUrlTokens,omitempty"`
						KeystoneAuthUrlVersion          *string            `tfsdk:"keystone_auth_url_version" json:"keystoneAuthUrlVersion,omitempty"`
						KeystoneAuthUserDomainName      *string            `tfsdk:"keystone_auth_user_domain_name" json:"keystoneAuthUserDomainName,omitempty"`
						KubernetesApiPort               *string            `tfsdk:"kubernetes_api_port" json:"kubernetesApiPort,omitempty"`
						KubernetesApiSecurePort         *string            `tfsdk:"kubernetes_api_secure_port" json:"kubernetesApiSecurePort,omitempty"`
						KubernetesPodSubnet             *string            `tfsdk:"kubernetes_pod_subnet" json:"kubernetesPodSubnet,omitempty"`
						L3mhCidr                        *string            `tfsdk:"l3mh_cidr" json:"l3mhCidr,omitempty"`
						LogDir                          *string            `tfsdk:"log_dir" json:"logDir,omitempty"`
						LogLocal                        *int64             `tfsdk:"log_local" json:"logLocal,omitempty"`
						MetadataProxySecret             *string            `tfsdk:"metadata_proxy_secret" json:"metadataProxySecret,omitempty"`
						MetadataSslCaCertfile           *string            `tfsdk:"metadata_ssl_ca_certfile" json:"metadataSslCaCertfile,omitempty"`
						MetadataSslCertType             *string            `tfsdk:"metadata_ssl_cert_type" json:"metadataSslCertType,omitempty"`
						MetadataSslCertfile             *string            `tfsdk:"metadata_ssl_certfile" json:"metadataSslCertfile,omitempty"`
						MetadataSslEnable               *string            `tfsdk:"metadata_ssl_enable" json:"metadataSslEnable,omitempty"`
						MetadataSslKeyfile              *string            `tfsdk:"metadata_ssl_keyfile" json:"metadataSslKeyfile,omitempty"`
						PhysicalInterface               *string            `tfsdk:"physical_interface" json:"physicalInterface,omitempty"`
						PriorityBandwidth               *string            `tfsdk:"priority_bandwidth" json:"priorityBandwidth,omitempty"`
						PriorityId                      *string            `tfsdk:"priority_id" json:"priorityId,omitempty"`
						PriorityScheduling              *string            `tfsdk:"priority_scheduling" json:"priorityScheduling,omitempty"`
						PriorityTagging                 *bool              `tfsdk:"priority_tagging" json:"priorityTagging,omitempty"`
						QosDefHwQueue                   *bool              `tfsdk:"qos_def_hw_queue" json:"qosDefHwQueue,omitempty"`
						QosLogicalQueues                *string            `tfsdk:"qos_logical_queues" json:"qosLogicalQueues,omitempty"`
						QosQueueId                      *string            `tfsdk:"qos_queue_id" json:"qosQueueId,omitempty"`
						RequiredKernelVrouterEncryption *string            `tfsdk:"required_kernel_vrouter_encryption" json:"requiredKernelVrouterEncryption,omitempty"`
						SampleDestination               *string            `tfsdk:"sample_destination" json:"sampleDestination,omitempty"`
						SandeshCaCertfile               *string            `tfsdk:"sandesh_ca_certfile" json:"sandeshCaCertfile,omitempty"`
						SandeshCertfile                 *string            `tfsdk:"sandesh_certfile" json:"sandeshCertfile,omitempty"`
						SandeshKeyfile                  *string            `tfsdk:"sandesh_keyfile" json:"sandeshKeyfile,omitempty"`
						SandeshServerCertfile           *string            `tfsdk:"sandesh_server_certfile" json:"sandeshServerCertfile,omitempty"`
						SandeshServerKeyfile            *string            `tfsdk:"sandesh_server_keyfile" json:"sandeshServerKeyfile,omitempty"`
						SandeshSslEnable                *bool              `tfsdk:"sandesh_ssl_enable" json:"sandeshSslEnable,omitempty"`
						ServerCaCertfile                *string            `tfsdk:"server_ca_certfile" json:"serverCaCertfile,omitempty"`
						ServerCertfile                  *string            `tfsdk:"server_certfile" json:"serverCertfile,omitempty"`
						ServerKeyfile                   *string            `tfsdk:"server_keyfile" json:"serverKeyfile,omitempty"`
						SloDestination                  *string            `tfsdk:"slo_destination" json:"sloDestination,omitempty"`
						SriovPhysicalInterface          *string            `tfsdk:"sriov_physical_interface" json:"sriovPhysicalInterface,omitempty"`
						SriovPhysicalNetwork            *string            `tfsdk:"sriov_physical_network" json:"sriovPhysicalNetwork,omitempty"`
						SriovVf                         *string            `tfsdk:"sriov_vf" json:"sriovVf,omitempty"`
						SslEnable                       *bool              `tfsdk:"ssl_enable" json:"sslEnable,omitempty"`
						SslInsecure                     *bool              `tfsdk:"ssl_insecure" json:"sslInsecure,omitempty"`
						StatsCollectorDestinationPath   *string            `tfsdk:"stats_collector_destination_path" json:"statsCollectorDestinationPath,omitempty"`
						Subcluster                      *string            `tfsdk:"subcluster" json:"subcluster,omitempty"`
						TsnAgentMode                    *string            `tfsdk:"tsn_agent_mode" json:"tsnAgentMode,omitempty"`
						VrouterCryptInterface           *string            `tfsdk:"vrouter_crypt_interface" json:"vrouterCryptInterface,omitempty"`
						VrouterDecryptInterface         *string            `tfsdk:"vrouter_decrypt_interface" json:"vrouterDecryptInterface,omitempty"`
						VrouterDecryptKey               *string            `tfsdk:"vrouter_decrypt_key" json:"vrouterDecryptKey,omitempty"`
						VrouterEncryption               *bool              `tfsdk:"vrouter_encryption" json:"vrouterEncryption,omitempty"`
						VrouterGateway                  *string            `tfsdk:"vrouter_gateway" json:"vrouterGateway,omitempty"`
						XmmpSslEnable                   *bool              `tfsdk:"xmmp_ssl_enable" json:"xmmpSslEnable,omitempty"`
						XmppServerCaCertfile            *string            `tfsdk:"xmpp_server_ca_certfile" json:"xmppServerCaCertfile,omitempty"`
						XmppServerCertfile              *string            `tfsdk:"xmpp_server_certfile" json:"xmppServerCertfile,omitempty"`
						XmppServerKeyfile               *string            `tfsdk:"xmpp_server_keyfile" json:"xmppServerKeyfile,omitempty"`
						XmppServerPort                  *string            `tfsdk:"xmpp_server_port" json:"xmppServerPort,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"vrouters" json:"vrouters,omitempty"`
			Webui *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						Containers *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						ControlInstance *string `tfsdk:"control_instance" json:"controlInstance,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"webui" json:"webui,omitempty"`
			Zookeeper *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name   *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					CommonConfiguration *struct {
						AuthParameters *struct {
							AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
							KeystoneAuthParameters *struct {
								Address           *string `tfsdk:"address" json:"address,omitempty"`
								AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
								AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
								AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
								AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
								AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
								Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
								Port              *int64  `tfsdk:"port" json:"port,omitempty"`
								ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
								Region            *string `tfsdk:"region" json:"region,omitempty"`
								UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
							} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
							KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
						} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
						Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
						ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
						NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Tolerations      *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
					ServiceConfiguration *struct {
						AdminEnabled *bool  `tfsdk:"admin_enabled" json:"adminEnabled,omitempty"`
						AdminPort    *int64 `tfsdk:"admin_port" json:"adminPort,omitempty"`
						ClientPort   *int64 `tfsdk:"client_port" json:"clientPort,omitempty"`
						Containers   *[]struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
							Image   *string   `tfsdk:"image" json:"image,omitempty"`
							Name    *string   `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						ElectionPort *int64 `tfsdk:"election_port" json:"electionPort,omitempty"`
						ServerPort   *int64 `tfsdk:"server_port" json:"serverPort,omitempty"`
					} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"zookeeper" json:"zookeeper,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TfTungstenIoManagerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tf_tungsten_io_manager_v1alpha1_manifest"
}

func (r *TfTungstenIoManagerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Manager is the Schema for the managers API.",
		MarkdownDescription: "Manager is the Schema for the managers API.",
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
				Description:         "ManagerSpec defines the desired state of Manager.",
				MarkdownDescription: "ManagerSpec defines the desired state of Manager.",
				Attributes: map[string]schema.Attribute{
					"common_configuration": schema.SingleNestedAttribute{
						Description:         "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'operator-sdk generate k8s' to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
						MarkdownDescription: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'operator-sdk generate k8s' to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
						Attributes: map[string]schema.Attribute{
							"auth_parameters": schema.SingleNestedAttribute{
								Description:         "AuthParameters auth parameters",
								MarkdownDescription: "AuthParameters auth parameters",
								Attributes: map[string]schema.Attribute{
									"auth_mode": schema.StringAttribute{
										Description:         "AuthenticationMode auth mode",
										MarkdownDescription: "AuthenticationMode auth mode",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("noauth", "keystone"),
										},
									},

									"keystone_auth_parameters": schema.SingleNestedAttribute{
										Description:         "KeystoneAuthParameters keystone parameters",
										MarkdownDescription: "KeystoneAuthParameters keystone parameters",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_password": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_tenant": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_username": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auth_protocol": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure": schema.BoolAttribute{
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

											"project_domain_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"region": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_domain_name": schema.StringAttribute{
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

									"keystone_secret_name": schema.StringAttribute{
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

							"cert_key_length": schema.Int64Attribute{
								Description:         "Certificate private key length",
								MarkdownDescription: "Certificate private key length",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_signer": schema.StringAttribute{
								Description:         "Certificate signer",
								MarkdownDescription: "Certificate signer",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"distribution": schema.StringAttribute{
								Description:         "OS family",
								MarkdownDescription: "OS family",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListAttribute{
								Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Kubernetes Cluster Configuration",
								MarkdownDescription: "Kubernetes Cluster Configuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
								},
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
								MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
								ElementType:         types.StringType,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"services": schema.SingleNestedAttribute{
						Description:         "Services defines the desired state of Services.",
						MarkdownDescription: "Services defines the desired state of Services.",
						Attributes: map[string]schema.Attribute{
							"analytics": schema.SingleNestedAttribute{
								Description:         "AnalyticsInput is the Schema for the analytics API.",
								MarkdownDescription: "AnalyticsInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "AnalyticsSpec is the Spec for the Analytics API.",
										MarkdownDescription: "AnalyticsSpec is the Spec for the Analytics API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "AnalyticsConfiguration is the Spec for the Analytics API.",
												MarkdownDescription: "AnalyticsConfiguration is the Spec for the Analytics API.",
												Attributes: map[string]schema.Attribute{
													"aaa_mode": schema.StringAttribute{
														Description:         "AAAMode aaa mode",
														MarkdownDescription: "AAAMode aaa mode",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("noauth", "rbac"),
														},
													},

													"analytics_config_audit_ttl": schema.Int64Attribute{
														Description:         "Time (in hours) the analytics config data entering the collector stays in the Cassandra database. Defaults to 2160 hours.",
														MarkdownDescription: "Time (in hours) the analytics config data entering the collector stays in the Cassandra database. Defaults to 2160 hours.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"analytics_data_ttl": schema.Int64Attribute{
														Description:         "Time (in hours) that the analytics object and log data stays in the Cassandra database. Defaults to 48 hours.",
														MarkdownDescription: "Time (in hours) that the analytics object and log data stays in the Cassandra database. Defaults to 48 hours.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"analytics_flow_ttl": schema.Int64Attribute{
														Description:         "Time to live (TTL) for flow data in hours. Defaults to 2 hours.",
														MarkdownDescription: "Time to live (TTL) for flow data in hours. Defaults to 2 hours.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"analytics_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"analytics_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"analytics_statistics_ttl": schema.Int64Attribute{
														Description:         "Time to live (TTL) for statistics data in hours. Defaults to 4 hours.",
														MarkdownDescription: "Time to live (TTL) for statistics data in hours. Defaults to 4 hours.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"collector_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"collector_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"analytics_alarm": schema.SingleNestedAttribute{
								Description:         "AnalyticsAlarmInput is the Schema for the analytics API.",
								MarkdownDescription: "AnalyticsAlarmInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "AnalyticsAlarmSpec is the Spec for the Analytics Alarm API.",
										MarkdownDescription: "AnalyticsAlarmSpec is the Spec for the Analytics Alarm API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "AnalyticsAlarmConfiguration is the Spec for the Analytics Alarm API.",
												MarkdownDescription: "AnalyticsAlarmConfiguration is the Spec for the Analytics Alarm API.",
												Attributes: map[string]schema.Attribute{
													"alarmgen_introspect_listen_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"alarmgen_log_file_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"alarmgen_partitions": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"alarmgen_redis_aggregate_db_offset": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"log_file_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_local": schema.StringAttribute{
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

							"analytics_snmp": schema.SingleNestedAttribute{
								Description:         "AnalyticsSnmpInput is the Schema for the analytics API.",
								MarkdownDescription: "AnalyticsSnmpInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "AnalyticsSnmpSpec is the Spec for the Analytics SNMP API.",
										MarkdownDescription: "AnalyticsSnmpSpec is the Spec for the Analytics SNMP API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "AnalyticsSnmpConfiguration is the Spec for the Analytics SNMP API.",
												MarkdownDescription: "AnalyticsSnmpConfiguration is the Spec for the Analytics SNMP API.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"log_file_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_local": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"snmp_collector_fast_scan_frequency": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"snmp_collector_introspect_listen_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"snmp_collector_log_file_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"snmp_collector_scan_frequency": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"topology_introspect_listen_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"topology_log_file_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"topology_snmp_frequency": schema.Int64Attribute{
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

							"cassandras": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "Input data is the Schema for the analytics API.",
											MarkdownDescription: "Input data is the Schema for the analytics API.",
											Attributes: map[string]schema.Attribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "CassandraSpec is the Spec for the cassandras API.",
											MarkdownDescription: "CassandraSpec is the Spec for the cassandras API.",
											Attributes: map[string]schema.Attribute{
												"common_configuration": schema.SingleNestedAttribute{
													Description:         "PodConfiguration is the common services struct.",
													MarkdownDescription: "PodConfiguration is the common services struct.",
													Attributes: map[string]schema.Attribute{
														"auth_parameters": schema.SingleNestedAttribute{
															Description:         "AuthParameters auth parameters",
															MarkdownDescription: "AuthParameters auth parameters",
															Attributes: map[string]schema.Attribute{
																"auth_mode": schema.StringAttribute{
																	Description:         "AuthenticationMode auth mode",
																	MarkdownDescription: "AuthenticationMode auth mode",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("noauth", "keystone"),
																	},
																},

																"keystone_auth_parameters": schema.SingleNestedAttribute{
																	Description:         "KeystoneAuthParameters keystone parameters",
																	MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_password": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_tenant": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_username": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"auth_protocol": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"insecure": schema.BoolAttribute{
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

																		"project_domain_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"region": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user_domain_name": schema.StringAttribute{
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

																"keystone_secret_name": schema.StringAttribute{
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

														"distribution": schema.StringAttribute{
															Description:         "OS family",
															MarkdownDescription: "OS family",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_secrets": schema.ListAttribute{
															Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_level": schema.StringAttribute{
															Description:         "Kubernetes Cluster Configuration",
															MarkdownDescription: "Kubernetes Cluster Configuration",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
															},
														},

														"node_selector": schema.MapAttribute{
															Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															ElementType:         types.StringType,
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_configuration": schema.SingleNestedAttribute{
													Description:         "CassandraConfiguration is the Spec for the cassandras API.",
													MarkdownDescription: "CassandraConfiguration is the Spec for the cassandras API.",
													Attributes: map[string]schema.Attribute{
														"cassandra_parameters": schema.SingleNestedAttribute{
															Description:         "CassandraConfigParameters defines additional parameters for Cassandra confgiuration",
															MarkdownDescription: "CassandraConfigParameters defines additional parameters for Cassandra confgiuration",
															Attributes: map[string]schema.Attribute{
																"compaction_throughput_mb_per_sec": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"concurrent_compactors": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"concurrent_counter_writes": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"concurrent_materialized_view_writes": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"concurrent_reads": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"concurrent_writes": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"memtable_allocation_type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("heap_buffers", "offheap_buffers", "offheap_objects"),
																	},
																},

																"memtable_flush_writers": schema.Int64Attribute{
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

														"containers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"command": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"cql_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"jmx_local_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"listen_address": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_heap_size": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"min_heap_size": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"minimum_disk_gb": schema.Int64Attribute{
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

														"reaper_adm_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"reaper_app_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"reaper_enabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ssl_storage_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"start_rpc": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"storage_port": schema.Int64Attribute{
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

							"config": schema.SingleNestedAttribute{
								Description:         "ConfigInput is the Schema for the analytics API.",
								MarkdownDescription: "ConfigInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "ConfigSpec is the Spec for the Config API.",
										MarkdownDescription: "ConfigSpec is the Spec for the Config API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "ConfigConfiguration is the Spec for the Config API.",
												MarkdownDescription: "ConfigConfiguration is the Spec for the Config API.",
												Attributes: map[string]schema.Attribute{
													"aaa_mode": schema.StringAttribute{
														Description:         "AAAMode aaa mode",
														MarkdownDescription: "AAAMode aaa mode",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("noauth", "rbac"),
														},
													},

													"api_admin_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"api_worker_count": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bgp_auto_mesh": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bgp_enable4_byte": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"device_manager_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"fabric_mgmt_ip": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"global_asn_number": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"linklocal_service_config": schema.SingleNestedAttribute{
														Description:         "LinklocalServiceConfig is the Spec for link local coniguration",
														MarkdownDescription: "LinklocalServiceConfig is the Spec for link local coniguration",
														Attributes: map[string]schema.Attribute{
															"ip": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"ip_fabric_service_host": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"ip_fabric_service_port": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
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

															"port": schema.Int64Attribute{
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

													"schema_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"svc_monitor_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_external_tftp": schema.BoolAttribute{
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

							"controls": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "Input data is the Schema for the analytics API.",
											MarkdownDescription: "Input data is the Schema for the analytics API.",
											Attributes: map[string]schema.Attribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "ControlSpec is the Spec for the controls API.",
											MarkdownDescription: "ControlSpec is the Spec for the controls API.",
											Attributes: map[string]schema.Attribute{
												"common_configuration": schema.SingleNestedAttribute{
													Description:         "PodConfiguration is the common services struct.",
													MarkdownDescription: "PodConfiguration is the common services struct.",
													Attributes: map[string]schema.Attribute{
														"auth_parameters": schema.SingleNestedAttribute{
															Description:         "AuthParameters auth parameters",
															MarkdownDescription: "AuthParameters auth parameters",
															Attributes: map[string]schema.Attribute{
																"auth_mode": schema.StringAttribute{
																	Description:         "AuthenticationMode auth mode",
																	MarkdownDescription: "AuthenticationMode auth mode",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("noauth", "keystone"),
																	},
																},

																"keystone_auth_parameters": schema.SingleNestedAttribute{
																	Description:         "KeystoneAuthParameters keystone parameters",
																	MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_password": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_tenant": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_username": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"auth_protocol": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"insecure": schema.BoolAttribute{
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

																		"project_domain_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"region": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user_domain_name": schema.StringAttribute{
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

																"keystone_secret_name": schema.StringAttribute{
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

														"distribution": schema.StringAttribute{
															Description:         "OS family",
															MarkdownDescription: "OS family",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_secrets": schema.ListAttribute{
															Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_level": schema.StringAttribute{
															Description:         "Kubernetes Cluster Configuration",
															MarkdownDescription: "Kubernetes Cluster Configuration",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
															},
														},

														"node_selector": schema.MapAttribute{
															Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															ElementType:         types.StringType,
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_configuration": schema.SingleNestedAttribute{
													Description:         "ControlConfiguration is the Spec for the controls API.",
													MarkdownDescription: "ControlConfiguration is the Spec for the controls API.",
													Attributes: map[string]schema.Attribute{
														"asn_number": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"bgp_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"containers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"command": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"data_subnet": schema.StringAttribute{
															Description:         "DataSubnet allow to set alternative network in which control, nodemanager and dns services will listen. Local pod address from this subnet will be discovered and used both in configuration for hostip directive and provision script.",
															MarkdownDescription: "DataSubnet allow to set alternative network in which control, nodemanager and dns services will listen. Local pod address from this subnet will be discovered and used both in configuration for hostip directive and provision script.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\/(3[0-2]|2[0-9]|1[0-9]|[0-9]))$`), ""),
															},
														},

														"dns_introspect_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dns_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"rndckey": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"subcluster": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"xmpp_port": schema.Int64Attribute{
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

							"kubemanager": schema.SingleNestedAttribute{
								Description:         "KubemanagerInput is the Schema for the analytics API.",
								MarkdownDescription: "KubemanagerInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "KubemanagerSpec is the Spec for the kubemanager API.",
										MarkdownDescription: "KubemanagerSpec is the Spec for the kubemanager API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "KubemanagerConfiguration is the configuration for the kubemanager API.",
												MarkdownDescription: "KubemanagerConfiguration is the configuration for the kubemanager API.",
												Attributes: map[string]schema.Attribute{
													"cloud_orchestrator": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_network_service": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_fabric_forwarding": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_fabric_snat": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_fabric_subnets": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kubernetes_api_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kubernetes_apissl_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kubernetes_api_server": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kubernetes_token_file": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_subnet": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"public_fip_pool": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_subnet": schema.StringAttribute{
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

							"queryengine": schema.SingleNestedAttribute{
								Description:         "QueryEngineInput is the Schema for the analytics API.",
								MarkdownDescription: "QueryEngineInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "QueryEngineSpec is the Spec for the AnalyticsDB query engine.",
										MarkdownDescription: "QueryEngineSpec is the Spec for the AnalyticsDB query engine.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "QueryEngineConfiguration is the Spec for the AnalyticsDB query engine.",
												MarkdownDescription: "QueryEngineConfiguration is the Spec for the AnalyticsDB query engine.",
												Attributes: map[string]schema.Attribute{
													"analyticsdb_introspect_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"analyticsdb_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rabbitmq": schema.SingleNestedAttribute{
								Description:         "RabbitmqInput is the Schema for the analytics API.",
								MarkdownDescription: "RabbitmqInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "RabbitmqSpec is the Spec for the cassandras API.",
										MarkdownDescription: "RabbitmqSpec is the Spec for the cassandras API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "RabbitmqConfiguration is the Spec for the cassandras API.",
												MarkdownDescription: "RabbitmqConfiguration is the Spec for the cassandras API.",
												Attributes: map[string]schema.Attribute{
													"cluster_partition_handling": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"erl_epmd_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"erlang_cookie": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"mirrored_queue_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("exactly", "all", "nodes"),
														},
													},

													"password": schema.StringAttribute{
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

													"secret": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tcp_listen_options": schema.SingleNestedAttribute{
														Description:         "TCPListenOptionsConfig is configuration for RabbitMQ TCP listen",
														MarkdownDescription: "TCPListenOptionsConfig is configuration for RabbitMQ TCP listen",
														Attributes: map[string]schema.Attribute{
															"backlog": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exit_on_close": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"linger_on": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"linger_timeout": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"nodelay": schema.BoolAttribute{
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

													"user": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"vhost": schema.StringAttribute{
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

							"redis": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "Input data is the Schema for the analytics API.",
											MarkdownDescription: "Input data is the Schema for the analytics API.",
											Attributes: map[string]schema.Attribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "RedisSpec is the Spec for the redis API.",
											MarkdownDescription: "RedisSpec is the Spec for the redis API.",
											Attributes: map[string]schema.Attribute{
												"common_configuration": schema.SingleNestedAttribute{
													Description:         "PodConfiguration is the common services struct.",
													MarkdownDescription: "PodConfiguration is the common services struct.",
													Attributes: map[string]schema.Attribute{
														"auth_parameters": schema.SingleNestedAttribute{
															Description:         "AuthParameters auth parameters",
															MarkdownDescription: "AuthParameters auth parameters",
															Attributes: map[string]schema.Attribute{
																"auth_mode": schema.StringAttribute{
																	Description:         "AuthenticationMode auth mode",
																	MarkdownDescription: "AuthenticationMode auth mode",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("noauth", "keystone"),
																	},
																},

																"keystone_auth_parameters": schema.SingleNestedAttribute{
																	Description:         "KeystoneAuthParameters keystone parameters",
																	MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_password": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_tenant": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_username": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"auth_protocol": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"insecure": schema.BoolAttribute{
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

																		"project_domain_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"region": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user_domain_name": schema.StringAttribute{
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

																"keystone_secret_name": schema.StringAttribute{
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

														"distribution": schema.StringAttribute{
															Description:         "OS family",
															MarkdownDescription: "OS family",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_secrets": schema.ListAttribute{
															Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_level": schema.StringAttribute{
															Description:         "Kubernetes Cluster Configuration",
															MarkdownDescription: "Kubernetes Cluster Configuration",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
															},
														},

														"node_selector": schema.MapAttribute{
															Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															ElementType:         types.StringType,
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_configuration": schema.SingleNestedAttribute{
													Description:         "RedisConfiguration is the Spec for the redis API.",
													MarkdownDescription: "RedisConfiguration is the Spec for the redis API.",
													Attributes: map[string]schema.Attribute{
														"cluster_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"containers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"command": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"redis_port": schema.Int64Attribute{
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

							"vrouters": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "Input data is the Schema for the analytics API.",
											MarkdownDescription: "Input data is the Schema for the analytics API.",
											Attributes: map[string]schema.Attribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "VrouterSpec is the Spec for the vrouter API.",
											MarkdownDescription: "VrouterSpec is the Spec for the vrouter API.",
											Attributes: map[string]schema.Attribute{
												"common_configuration": schema.SingleNestedAttribute{
													Description:         "PodConfiguration is the common services struct.",
													MarkdownDescription: "PodConfiguration is the common services struct.",
													Attributes: map[string]schema.Attribute{
														"auth_parameters": schema.SingleNestedAttribute{
															Description:         "AuthParameters auth parameters",
															MarkdownDescription: "AuthParameters auth parameters",
															Attributes: map[string]schema.Attribute{
																"auth_mode": schema.StringAttribute{
																	Description:         "AuthenticationMode auth mode",
																	MarkdownDescription: "AuthenticationMode auth mode",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("noauth", "keystone"),
																	},
																},

																"keystone_auth_parameters": schema.SingleNestedAttribute{
																	Description:         "KeystoneAuthParameters keystone parameters",
																	MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_password": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_tenant": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"admin_username": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"auth_protocol": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"insecure": schema.BoolAttribute{
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

																		"project_domain_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"region": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user_domain_name": schema.StringAttribute{
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

																"keystone_secret_name": schema.StringAttribute{
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

														"distribution": schema.StringAttribute{
															Description:         "OS family",
															MarkdownDescription: "OS family",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_secrets": schema.ListAttribute{
															Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_level": schema.StringAttribute{
															Description:         "Kubernetes Cluster Configuration",
															MarkdownDescription: "Kubernetes Cluster Configuration",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
															},
														},

														"node_selector": schema.MapAttribute{
															Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
															ElementType:         types.StringType,
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_configuration": schema.SingleNestedAttribute{
													Description:         "VrouterConfiguration is the Spec for the vrouter API.",
													MarkdownDescription: "VrouterConfiguration is the Spec for the vrouter API.",
													Attributes: map[string]schema.Attribute{
														"agent_mode": schema.StringAttribute{
															Description:         "vRouter",
															MarkdownDescription: "vRouter",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"barbican_password": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"barbican_tenant_name": schema.StringAttribute{
															Description:         "Openstack",
															MarkdownDescription: "Openstack",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"barbican_user": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cloud_orchestrator": schema.StringAttribute{
															Description:         "New params for vrouter configuration",
															MarkdownDescription: "New params for vrouter configuration",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cni_mtu": schema.Int64Attribute{
															Description:         "CniMTU - mtu for virtual tap devices",
															MarkdownDescription: "CniMTU - mtu for virtual tap devices",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"collector_port": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"config_api_port": schema.StringAttribute{
															Description:         "Config",
															MarkdownDescription: "Config",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"config_api_server_ca_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"config_api_ssl_enable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"containers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"command": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"control_instance": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"data_subnet": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dns_server_port": schema.StringAttribute{
															Description:         "DNS",
															MarkdownDescription: "DNS",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dpdk_uio_driver": schema.StringAttribute{
															Description:         "Host",
															MarkdownDescription: "Host",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"env_variables_config": schema.MapAttribute{
															Description:         "What is it doing? VrouterEncryption bool 'json:'vrouterEncryption,omitempty'' What is it doing? What is it doing?",
															MarkdownDescription: "What is it doing? VrouterEncryption bool 'json:'vrouterEncryption,omitempty'' What is it doing? What is it doing?",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"fabric_snt_hash_table_size": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"huge_pages1_g": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"huge_pages2_m": schema.Int64Attribute{
															Description:         "HugePages",
															MarkdownDescription: "HugePages",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"hypervisor_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"introspect_ssl_enable": schema.BoolAttribute{
															Description:         "Introspect",
															MarkdownDescription: "Introspect",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"k8s_token": schema.StringAttribute{
															Description:         "Kubernetes",
															MarkdownDescription: "Kubernetes",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"k8s_token_file": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_admin_password": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_admin_port": schema.StringAttribute{
															Description:         "Keystone authentication",
															MarkdownDescription: "Keystone authentication",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_ca_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_insecure": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_keyfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_project_domain_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_proto": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_region_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_url_tokens": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_url_version": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keystone_auth_user_domain_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kubernetes_api_port": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kubernetes_api_secure_port": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kubernetes_pod_subnet": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"l3mh_cidr": schema.StringAttribute{
															Description:         "L3MH",
															MarkdownDescription: "L3MH",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_dir": schema.StringAttribute{
															Description:         "Logging",
															MarkdownDescription: "Logging",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_local": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata_proxy_secret": schema.StringAttribute{
															Description:         "Metadata",
															MarkdownDescription: "Metadata",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata_ssl_ca_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata_ssl_cert_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata_ssl_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata_ssl_enable": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata_ssl_keyfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"physical_interface": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"priority_bandwidth": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"priority_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"priority_scheduling": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"priority_tagging": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"qos_def_hw_queue": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"qos_logical_queues": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"qos_queue_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"required_kernel_vrouter_encryption": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sample_destination": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sandesh_ca_certfile": schema.StringAttribute{
															Description:         "Sandesh",
															MarkdownDescription: "Sandesh",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sandesh_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sandesh_keyfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sandesh_server_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sandesh_server_keyfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sandesh_ssl_enable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"server_ca_certfile": schema.StringAttribute{
															Description:         "Server SSL",
															MarkdownDescription: "Server SSL",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"server_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"server_keyfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"slo_destination": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sriov_physical_interface": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sriov_physical_network": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sriov_vf": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ssl_enable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ssl_insecure": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"stats_collector_destination_path": schema.StringAttribute{
															Description:         "Collector",
															MarkdownDescription: "Collector",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"subcluster": schema.StringAttribute{
															Description:         "XMPP",
															MarkdownDescription: "XMPP",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tsn_agent_mode": schema.StringAttribute{
															Description:         "TSN",
															MarkdownDescription: "TSN",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"vrouter_crypt_interface": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"vrouter_decrypt_interface": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"vrouter_decrypt_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"vrouter_encryption": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"vrouter_gateway": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"xmmp_ssl_enable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"xmpp_server_ca_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"xmpp_server_certfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"xmpp_server_keyfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"xmpp_server_port": schema.StringAttribute{
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

							"webui": schema.SingleNestedAttribute{
								Description:         "WebuiInput is the Schema for the analytics API.",
								MarkdownDescription: "WebuiInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "WebuiSpec is the Spec for the cassandras API.",
										MarkdownDescription: "WebuiSpec is the Spec for the cassandras API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "WebuiConfiguration is the Spec for the cassandras API.",
												MarkdownDescription: "WebuiConfiguration is the Spec for the cassandras API.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"control_instance": schema.StringAttribute{
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

							"zookeeper": schema.SingleNestedAttribute{
								Description:         "ZookeeperInput is the Schema for the analytics API.",
								MarkdownDescription: "ZookeeperInput is the Schema for the analytics API.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Input data is the Schema for the analytics API.",
										MarkdownDescription: "Input data is the Schema for the analytics API.",
										Attributes: map[string]schema.Attribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "ZookeeperSpec is the Spec for the zookeeper API.",
										MarkdownDescription: "ZookeeperSpec is the Spec for the zookeeper API.",
										Attributes: map[string]schema.Attribute{
											"common_configuration": schema.SingleNestedAttribute{
												Description:         "PodConfiguration is the common services struct.",
												MarkdownDescription: "PodConfiguration is the common services struct.",
												Attributes: map[string]schema.Attribute{
													"auth_parameters": schema.SingleNestedAttribute{
														Description:         "AuthParameters auth parameters",
														MarkdownDescription: "AuthParameters auth parameters",
														Attributes: map[string]schema.Attribute{
															"auth_mode": schema.StringAttribute{
																Description:         "AuthenticationMode auth mode",
																MarkdownDescription: "AuthenticationMode auth mode",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("noauth", "keystone"),
																},
															},

															"keystone_auth_parameters": schema.SingleNestedAttribute{
																Description:         "KeystoneAuthParameters keystone parameters",
																MarkdownDescription: "KeystoneAuthParameters keystone parameters",
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_password": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_port": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_tenant": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"admin_username": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"auth_protocol": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"insecure": schema.BoolAttribute{
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

																	"project_domain_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"region": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"user_domain_name": schema.StringAttribute{
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

															"keystone_secret_name": schema.StringAttribute{
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

													"distribution": schema.StringAttribute{
														Description:         "OS family",
														MarkdownDescription: "OS family",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListAttribute{
														Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "Kubernetes Cluster Configuration",
														MarkdownDescription: "Kubernetes Cluster Configuration",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
														},
													},

													"node_selector": schema.MapAttribute{
														Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_configuration": schema.SingleNestedAttribute{
												Description:         "ZookeeperConfiguration is the Spec for the zookeeper API.",
												MarkdownDescription: "ZookeeperConfiguration is the Spec for the zookeeper API.",
												Attributes: map[string]schema.Attribute{
													"admin_enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"admin_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"election_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"server_port": schema.Int64Attribute{
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
		},
	}
}

func (r *TfTungstenIoManagerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tf_tungsten_io_manager_v1alpha1_manifest")

	var model TfTungstenIoManagerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tf.tungsten.io/v1alpha1")
	model.Kind = pointer.String("Manager")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
