/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package trino_stackable_tech_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TrinoStackableTechTrinoClusterV1Alpha1Manifest{}
)

func NewTrinoStackableTechTrinoClusterV1Alpha1Manifest() datasource.DataSource {
	return &TrinoStackableTechTrinoClusterV1Alpha1Manifest{}
}

type TrinoStackableTechTrinoClusterV1Alpha1Manifest struct{}

type TrinoStackableTechTrinoClusterV1Alpha1ManifestData struct {
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
		ClusterConfig *struct {
			Authentication *[]struct {
				AuthenticationClass *string `tfsdk:"authentication_class" json:"authenticationClass,omitempty"`
				Oidc                *struct {
					ClientCredentialsSecret *string   `tfsdk:"client_credentials_secret" json:"clientCredentialsSecret,omitempty"`
					ExtraScopes             *[]string `tfsdk:"extra_scopes" json:"extraScopes,omitempty"`
				} `tfsdk:"oidc" json:"oidc,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			Authorization *struct {
				Opa *struct {
					ConfigMapName *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
					Package       *string `tfsdk:"package" json:"package,omitempty"`
				} `tfsdk:"opa" json:"opa,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			CatalogLabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"catalog_label_selector" json:"catalogLabelSelector,omitempty"`
			FaultTolerantExecution *struct {
				Query *struct {
					ExchangeDeduplicationBufferSize *string `tfsdk:"exchange_deduplication_buffer_size" json:"exchangeDeduplicationBufferSize,omitempty"`
					ExchangeManager                 *struct {
						ConfigOverrides   *map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
						EncryptionEnabled *bool              `tfsdk:"encryption_enabled" json:"encryptionEnabled,omitempty"`
						Hdfs              *struct {
							BaseDirectories *[]string `tfsdk:"base_directories" json:"baseDirectories,omitempty"`
							BlockSize       *string   `tfsdk:"block_size" json:"blockSize,omitempty"`
							Hdfs            *struct {
								ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
							} `tfsdk:"hdfs" json:"hdfs,omitempty"`
							SkipDirectorySchemeValidation *bool `tfsdk:"skip_directory_scheme_validation" json:"skipDirectorySchemeValidation,omitempty"`
						} `tfsdk:"hdfs" json:"hdfs,omitempty"`
						Local *struct {
							BaseDirectories *[]string `tfsdk:"base_directories" json:"baseDirectories,omitempty"`
						} `tfsdk:"local" json:"local,omitempty"`
						S3 *struct {
							BaseDirectories *[]string `tfsdk:"base_directories" json:"baseDirectories,omitempty"`
							Connection      *struct {
								Inline *struct {
									AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
									Credentials *struct {
										Scope *struct {
											ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
											Node            *bool     `tfsdk:"node" json:"node,omitempty"`
											Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
											Services        *[]string `tfsdk:"services" json:"services,omitempty"`
										} `tfsdk:"scope" json:"scope,omitempty"`
										SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
									} `tfsdk:"credentials" json:"credentials,omitempty"`
									Host   *string `tfsdk:"host" json:"host,omitempty"`
									Port   *int64  `tfsdk:"port" json:"port,omitempty"`
									Region *struct {
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"region" json:"region,omitempty"`
									Tls *struct {
										Verification *struct {
											None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
											Server *struct {
												CaCert *struct {
													SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
													WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
												} `tfsdk:"ca_cert" json:"caCert,omitempty"`
											} `tfsdk:"server" json:"server,omitempty"`
										} `tfsdk:"verification" json:"verification,omitempty"`
									} `tfsdk:"tls" json:"tls,omitempty"`
								} `tfsdk:"inline" json:"inline,omitempty"`
								Reference *string `tfsdk:"reference" json:"reference,omitempty"`
							} `tfsdk:"connection" json:"connection,omitempty"`
							ExternalId      *string `tfsdk:"external_id" json:"externalId,omitempty"`
							IamRole         *string `tfsdk:"iam_role" json:"iamRole,omitempty"`
							MaxErrorRetries *int64  `tfsdk:"max_error_retries" json:"maxErrorRetries,omitempty"`
							UploadPartSize  *string `tfsdk:"upload_part_size" json:"uploadPartSize,omitempty"`
						} `tfsdk:"s3" json:"s3,omitempty"`
						SinkBufferPoolMinSize   *int64  `tfsdk:"sink_buffer_pool_min_size" json:"sinkBufferPoolMinSize,omitempty"`
						SinkBuffersPerPartition *int64  `tfsdk:"sink_buffers_per_partition" json:"sinkBuffersPerPartition,omitempty"`
						SinkMaxFileSize         *string `tfsdk:"sink_max_file_size" json:"sinkMaxFileSize,omitempty"`
						SourceConcurrentReaders *int64  `tfsdk:"source_concurrent_readers" json:"sourceConcurrentReaders,omitempty"`
					} `tfsdk:"exchange_manager" json:"exchangeManager,omitempty"`
					RetryAttempts         *int64   `tfsdk:"retry_attempts" json:"retryAttempts,omitempty"`
					RetryDelayScaleFactor *float64 `tfsdk:"retry_delay_scale_factor" json:"retryDelayScaleFactor,omitempty"`
					RetryInitialDelay     *string  `tfsdk:"retry_initial_delay" json:"retryInitialDelay,omitempty"`
					RetryMaxDelay         *string  `tfsdk:"retry_max_delay" json:"retryMaxDelay,omitempty"`
				} `tfsdk:"query" json:"query,omitempty"`
				Task *struct {
					ExchangeDeduplicationBufferSize *string `tfsdk:"exchange_deduplication_buffer_size" json:"exchangeDeduplicationBufferSize,omitempty"`
					ExchangeManager                 *struct {
						ConfigOverrides   *map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
						EncryptionEnabled *bool              `tfsdk:"encryption_enabled" json:"encryptionEnabled,omitempty"`
						Hdfs              *struct {
							BaseDirectories *[]string `tfsdk:"base_directories" json:"baseDirectories,omitempty"`
							BlockSize       *string   `tfsdk:"block_size" json:"blockSize,omitempty"`
							Hdfs            *struct {
								ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
							} `tfsdk:"hdfs" json:"hdfs,omitempty"`
							SkipDirectorySchemeValidation *bool `tfsdk:"skip_directory_scheme_validation" json:"skipDirectorySchemeValidation,omitempty"`
						} `tfsdk:"hdfs" json:"hdfs,omitempty"`
						Local *struct {
							BaseDirectories *[]string `tfsdk:"base_directories" json:"baseDirectories,omitempty"`
						} `tfsdk:"local" json:"local,omitempty"`
						S3 *struct {
							BaseDirectories *[]string `tfsdk:"base_directories" json:"baseDirectories,omitempty"`
							Connection      *struct {
								Inline *struct {
									AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
									Credentials *struct {
										Scope *struct {
											ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
											Node            *bool     `tfsdk:"node" json:"node,omitempty"`
											Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
											Services        *[]string `tfsdk:"services" json:"services,omitempty"`
										} `tfsdk:"scope" json:"scope,omitempty"`
										SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
									} `tfsdk:"credentials" json:"credentials,omitempty"`
									Host   *string `tfsdk:"host" json:"host,omitempty"`
									Port   *int64  `tfsdk:"port" json:"port,omitempty"`
									Region *struct {
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"region" json:"region,omitempty"`
									Tls *struct {
										Verification *struct {
											None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
											Server *struct {
												CaCert *struct {
													SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
													WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
												} `tfsdk:"ca_cert" json:"caCert,omitempty"`
											} `tfsdk:"server" json:"server,omitempty"`
										} `tfsdk:"verification" json:"verification,omitempty"`
									} `tfsdk:"tls" json:"tls,omitempty"`
								} `tfsdk:"inline" json:"inline,omitempty"`
								Reference *string `tfsdk:"reference" json:"reference,omitempty"`
							} `tfsdk:"connection" json:"connection,omitempty"`
							ExternalId      *string `tfsdk:"external_id" json:"externalId,omitempty"`
							IamRole         *string `tfsdk:"iam_role" json:"iamRole,omitempty"`
							MaxErrorRetries *int64  `tfsdk:"max_error_retries" json:"maxErrorRetries,omitempty"`
							UploadPartSize  *string `tfsdk:"upload_part_size" json:"uploadPartSize,omitempty"`
						} `tfsdk:"s3" json:"s3,omitempty"`
						SinkBufferPoolMinSize   *int64  `tfsdk:"sink_buffer_pool_min_size" json:"sinkBufferPoolMinSize,omitempty"`
						SinkBuffersPerPartition *int64  `tfsdk:"sink_buffers_per_partition" json:"sinkBuffersPerPartition,omitempty"`
						SinkMaxFileSize         *string `tfsdk:"sink_max_file_size" json:"sinkMaxFileSize,omitempty"`
						SourceConcurrentReaders *int64  `tfsdk:"source_concurrent_readers" json:"sourceConcurrentReaders,omitempty"`
					} `tfsdk:"exchange_manager" json:"exchangeManager,omitempty"`
					RetryAttemptsPerTask  *int64   `tfsdk:"retry_attempts_per_task" json:"retryAttemptsPerTask,omitempty"`
					RetryDelayScaleFactor *float64 `tfsdk:"retry_delay_scale_factor" json:"retryDelayScaleFactor,omitempty"`
					RetryInitialDelay     *string  `tfsdk:"retry_initial_delay" json:"retryInitialDelay,omitempty"`
					RetryMaxDelay         *string  `tfsdk:"retry_max_delay" json:"retryMaxDelay,omitempty"`
				} `tfsdk:"task" json:"task,omitempty"`
			} `tfsdk:"fault_tolerant_execution" json:"faultTolerantExecution,omitempty"`
			Tls *struct {
				InternalSecretClass *string `tfsdk:"internal_secret_class" json:"internalSecretClass,omitempty"`
				ServerSecretClass   *string `tfsdk:"server_secret_class" json:"serverSecretClass,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			VectorAggregatorConfigMapName *string `tfsdk:"vector_aggregator_config_map_name" json:"vectorAggregatorConfigMapName,omitempty"`
		} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		ClusterOperation *struct {
			ReconciliationPaused *bool `tfsdk:"reconciliation_paused" json:"reconciliationPaused,omitempty"`
			Stopped              *bool `tfsdk:"stopped" json:"stopped,omitempty"`
		} `tfsdk:"cluster_operation" json:"clusterOperation,omitempty"`
		Coordinators *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
				Logging                 *struct {
					Containers *struct {
						Console *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"console" json:"console,omitempty"`
						Custom *struct {
							ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
						} `tfsdk:"custom" json:"custom,omitempty"`
						File *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Loggers *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"loggers" json:"loggers,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				QueryMaxMemory          *string `tfsdk:"query_max_memory" json:"queryMaxMemory,omitempty"`
				QueryMaxMemoryPerNode   *string `tfsdk:"query_max_memory_per_node" json:"queryMaxMemoryPerNode,omitempty"`
				RequestedSecretLifetime *string `tfsdk:"requested_secret_lifetime" json:"requestedSecretLifetime,omitempty"`
				Resources               *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides      *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides         *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			JvmArgumentOverrides *struct {
				Add         *[]string `tfsdk:"add" json:"add,omitempty"`
				Remove      *[]string `tfsdk:"remove" json:"remove,omitempty"`
				RemoveRegex *[]string `tfsdk:"remove_regex" json:"removeRegex,omitempty"`
			} `tfsdk:"jvm_argument_overrides" json:"jvmArgumentOverrides,omitempty"`
			PodOverrides *map[string]string `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
			RoleConfig   *struct {
				ListenerClass       *string `tfsdk:"listener_class" json:"listenerClass,omitempty"`
				PodDisruptionBudget *struct {
					Enabled        *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			} `tfsdk:"role_config" json:"roleConfig,omitempty"`
			RoleGroups *struct {
				CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
				Config       *struct {
					Affinity *struct {
						NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
						NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
						PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
					Logging                 *struct {
						Containers *struct {
							Console *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"console" json:"console,omitempty"`
							Custom *struct {
								ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
							} `tfsdk:"custom" json:"custom,omitempty"`
							File *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"file" json:"file,omitempty"`
							Loggers *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"loggers" json:"loggers,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
					} `tfsdk:"logging" json:"logging,omitempty"`
					QueryMaxMemory          *string `tfsdk:"query_max_memory" json:"queryMaxMemory,omitempty"`
					QueryMaxMemoryPerNode   *string `tfsdk:"query_max_memory_per_node" json:"queryMaxMemoryPerNode,omitempty"`
					RequestedSecretLifetime *string `tfsdk:"requested_secret_lifetime" json:"requestedSecretLifetime,omitempty"`
					Resources               *struct {
						Cpu *struct {
							Max *string `tfsdk:"max" json:"max,omitempty"`
							Min *string `tfsdk:"min" json:"min,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
							RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides      *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides         *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				JvmArgumentOverrides *struct {
					Add         *[]string `tfsdk:"add" json:"add,omitempty"`
					Remove      *[]string `tfsdk:"remove" json:"remove,omitempty"`
					RemoveRegex *[]string `tfsdk:"remove_regex" json:"removeRegex,omitempty"`
				} `tfsdk:"jvm_argument_overrides" json:"jvmArgumentOverrides,omitempty"`
				PodOverrides *map[string]string `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"coordinators" json:"coordinators,omitempty"`
		Image *struct {
			Custom         *string `tfsdk:"custom" json:"custom,omitempty"`
			ProductVersion *string `tfsdk:"product_version" json:"productVersion,omitempty"`
			PullPolicy     *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			PullSecrets    *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
			Repo             *string `tfsdk:"repo" json:"repo,omitempty"`
			StackableVersion *string `tfsdk:"stackable_version" json:"stackableVersion,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		Workers *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
				Logging                 *struct {
					Containers *struct {
						Console *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"console" json:"console,omitempty"`
						Custom *struct {
							ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
						} `tfsdk:"custom" json:"custom,omitempty"`
						File *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Loggers *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"loggers" json:"loggers,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				QueryMaxMemory          *string `tfsdk:"query_max_memory" json:"queryMaxMemory,omitempty"`
				QueryMaxMemoryPerNode   *string `tfsdk:"query_max_memory_per_node" json:"queryMaxMemoryPerNode,omitempty"`
				RequestedSecretLifetime *string `tfsdk:"requested_secret_lifetime" json:"requestedSecretLifetime,omitempty"`
				Resources               *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides      *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides         *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			JvmArgumentOverrides *struct {
				Add         *[]string `tfsdk:"add" json:"add,omitempty"`
				Remove      *[]string `tfsdk:"remove" json:"remove,omitempty"`
				RemoveRegex *[]string `tfsdk:"remove_regex" json:"removeRegex,omitempty"`
			} `tfsdk:"jvm_argument_overrides" json:"jvmArgumentOverrides,omitempty"`
			PodOverrides *map[string]string `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
			RoleConfig   *struct {
				PodDisruptionBudget *struct {
					Enabled        *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			} `tfsdk:"role_config" json:"roleConfig,omitempty"`
			RoleGroups *struct {
				CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
				Config       *struct {
					Affinity *struct {
						NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
						NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
						PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
					Logging                 *struct {
						Containers *struct {
							Console *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"console" json:"console,omitempty"`
							Custom *struct {
								ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
							} `tfsdk:"custom" json:"custom,omitempty"`
							File *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"file" json:"file,omitempty"`
							Loggers *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"loggers" json:"loggers,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
					} `tfsdk:"logging" json:"logging,omitempty"`
					QueryMaxMemory          *string `tfsdk:"query_max_memory" json:"queryMaxMemory,omitempty"`
					QueryMaxMemoryPerNode   *string `tfsdk:"query_max_memory_per_node" json:"queryMaxMemoryPerNode,omitempty"`
					RequestedSecretLifetime *string `tfsdk:"requested_secret_lifetime" json:"requestedSecretLifetime,omitempty"`
					Resources               *struct {
						Cpu *struct {
							Max *string `tfsdk:"max" json:"max,omitempty"`
							Min *string `tfsdk:"min" json:"min,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
							RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides      *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides         *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				JvmArgumentOverrides *struct {
					Add         *[]string `tfsdk:"add" json:"add,omitempty"`
					Remove      *[]string `tfsdk:"remove" json:"remove,omitempty"`
					RemoveRegex *[]string `tfsdk:"remove_regex" json:"removeRegex,omitempty"`
				} `tfsdk:"jvm_argument_overrides" json:"jvmArgumentOverrides,omitempty"`
				PodOverrides *map[string]string `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"workers" json:"workers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TrinoStackableTechTrinoClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_trino_stackable_tech_trino_cluster_v1alpha1_manifest"
}

func (r *TrinoStackableTechTrinoClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for TrinoClusterSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for TrinoClusterSpec via 'CustomResource'",
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
				Description:         "A Trino cluster stacklet. This resource is managed by the Stackable operator for Trino. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/trino/).",
				MarkdownDescription: "A Trino cluster stacklet. This resource is managed by the Stackable operator for Trino. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/trino/).",
				Attributes: map[string]schema.Attribute{
					"cluster_config": schema.SingleNestedAttribute{
						Description:         "Settings that affect all roles and role groups. The settings in the 'clusterConfig' are cluster wide settings that do not need to be configurable at role or role group level.",
						MarkdownDescription: "Settings that affect all roles and role groups. The settings in the 'clusterConfig' are cluster wide settings that do not need to be configurable at role or role group level.",
						Attributes: map[string]schema.Attribute{
							"authentication": schema.ListNestedAttribute{
								Description:         "Authentication options for Trino. Learn more in the [Trino authentication usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/security#authentication).",
								MarkdownDescription: "Authentication options for Trino. Learn more in the [Trino authentication usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/security#authentication).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication_class": schema.StringAttribute{
											Description:         "Name of the [AuthenticationClass](https://docs.stackable.tech/home/nightly/concepts/authentication) used to authenticate users",
											MarkdownDescription: "Name of the [AuthenticationClass](https://docs.stackable.tech/home/nightly/concepts/authentication) used to authenticate users",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"oidc": schema.SingleNestedAttribute{
											Description:         "This field contains OIDC-specific configuration. It is only required in case OIDC is used.",
											MarkdownDescription: "This field contains OIDC-specific configuration. It is only required in case OIDC is used.",
											Attributes: map[string]schema.Attribute{
												"client_credentials_secret": schema.StringAttribute{
													Description:         "A reference to the OIDC client credentials secret. The secret contains the client id and secret.",
													MarkdownDescription: "A reference to the OIDC client credentials secret. The secret contains the client id and secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"extra_scopes": schema.ListAttribute{
													Description:         "An optional list of extra scopes which get merged with the scopes defined in the AuthenticationClass",
													MarkdownDescription: "An optional list of extra scopes which get merged with the scopes defined in the AuthenticationClass",
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

							"authorization": schema.SingleNestedAttribute{
								Description:         "Authorization options for Trino. Learn more in the [Trino authorization usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/security#authorization).",
								MarkdownDescription: "Authorization options for Trino. Learn more in the [Trino authorization usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/security#authorization).",
								Attributes: map[string]schema.Attribute{
									"opa": schema.SingleNestedAttribute{
										Description:         "Configure the OPA stacklet [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) and the name of the Rego package containing your authorization rules. Consult the [OPA authorization documentation](https://docs.stackable.tech/home/nightly/concepts/opa) to learn how to deploy Rego authorization rules with OPA.",
										MarkdownDescription: "Configure the OPA stacklet [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) and the name of the Rego package containing your authorization rules. Consult the [OPA authorization documentation](https://docs.stackable.tech/home/nightly/concepts/opa) to learn how to deploy Rego authorization rules with OPA.",
										Attributes: map[string]schema.Attribute{
											"config_map_name": schema.StringAttribute{
												Description:         "The [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) for the OPA stacklet that should be used for authorization requests.",
												MarkdownDescription: "The [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) for the OPA stacklet that should be used for authorization requests.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"package": schema.StringAttribute{
												Description:         "The name of the Rego package containing the Rego rules for the product.",
												MarkdownDescription: "The name of the Rego package containing the Rego rules for the product.",
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

							"catalog_label_selector": schema.SingleNestedAttribute{
								Description:         "[LabelSelector](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors) selecting the Catalogs to include in the Trino instance.",
								MarkdownDescription: "[LabelSelector](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors) selecting the Catalogs to include in the Trino instance.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"fault_tolerant_execution": schema.SingleNestedAttribute{
								Description:         "Fault tolerant execution configuration. When enabled, Trino can automatically retry queries or tasks in case of failures.",
								MarkdownDescription: "Fault tolerant execution configuration. When enabled, Trino can automatically retry queries or tasks in case of failures.",
								Attributes: map[string]schema.Attribute{
									"query": schema.SingleNestedAttribute{
										Description:         "Query-level fault tolerant execution. Retries entire queries on failure.",
										MarkdownDescription: "Query-level fault tolerant execution. Retries entire queries on failure.",
										Attributes: map[string]schema.Attribute{
											"exchange_deduplication_buffer_size": schema.StringAttribute{
												Description:         "Data size of the coordinator's in-memory buffer used to store output of query stages.",
												MarkdownDescription: "Data size of the coordinator's in-memory buffer used to store output of query stages.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"exchange_manager": schema.SingleNestedAttribute{
												Description:         "Exchange manager configuration for spooling intermediate data during fault tolerant execution. Optional for Query retry policy, recommended for large result sets.",
												MarkdownDescription: "Exchange manager configuration for spooling intermediate data during fault tolerant execution. Optional for Query retry policy, recommended for large result sets.",
												Attributes: map[string]schema.Attribute{
													"config_overrides": schema.MapAttribute{
														Description:         "The 'configOverrides' allow overriding arbitrary exchange manager properties.",
														MarkdownDescription: "The 'configOverrides' allow overriding arbitrary exchange manager properties.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"encryption_enabled": schema.BoolAttribute{
														Description:         "Whether to enable encryption of spooling data.",
														MarkdownDescription: "Whether to enable encryption of spooling data.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hdfs": schema.SingleNestedAttribute{
														Description:         "HDFS-based exchange manager.",
														MarkdownDescription: "HDFS-based exchange manager.",
														Attributes: map[string]schema.Attribute{
															"base_directories": schema.ListAttribute{
																Description:         "HDFS URIs for spooling data.",
																MarkdownDescription: "HDFS URIs for spooling data.",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"block_size": schema.StringAttribute{
																Description:         "Block data size for HDFS storage.",
																MarkdownDescription: "Block data size for HDFS storage.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"hdfs": schema.SingleNestedAttribute{
																Description:         "HDFS connection configuration.",
																MarkdownDescription: "HDFS connection configuration.",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.StringAttribute{
																		Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
																		MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"skip_directory_scheme_validation": schema.BoolAttribute{
																Description:         "Skip directory scheme validation to support Hadoop-compatible file systems.",
																MarkdownDescription: "Skip directory scheme validation to support Hadoop-compatible file systems.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"local": schema.SingleNestedAttribute{
														Description:         "Local filesystem storage (not recommended for production).",
														MarkdownDescription: "Local filesystem storage (not recommended for production).",
														Attributes: map[string]schema.Attribute{
															"base_directories": schema.ListAttribute{
																Description:         "Local filesystem paths for exchange storage.",
																MarkdownDescription: "Local filesystem paths for exchange storage.",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"s3": schema.SingleNestedAttribute{
														Description:         "S3-compatible storage configuration.",
														MarkdownDescription: "S3-compatible storage configuration.",
														Attributes: map[string]schema.Attribute{
															"base_directories": schema.ListAttribute{
																Description:         "S3 bucket URIs for spooling data (e.g., s3://bucket1,s3://bucket2).",
																MarkdownDescription: "S3 bucket URIs for spooling data (e.g., s3://bucket1,s3://bucket2).",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"connection": schema.SingleNestedAttribute{
																Description:         "S3 connection configuration. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
																MarkdownDescription: "S3 connection configuration. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
																Attributes: map[string]schema.Attribute{
																	"inline": schema.SingleNestedAttribute{
																		Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
																		MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
																		Attributes: map[string]schema.Attribute{
																			"access_style": schema.StringAttribute{
																				Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
																				MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("Path", "VirtualHosted"),
																				},
																			},

																			"credentials": schema.SingleNestedAttribute{
																				Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
																				MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
																				Attributes: map[string]schema.Attribute{
																					"scope": schema.SingleNestedAttribute{
																						Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																						MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																						Attributes: map[string]schema.Attribute{
																							"listener_volumes": schema.ListAttribute{
																								Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																								MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"node": schema.BoolAttribute{
																								Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																								MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"pod": schema.BoolAttribute{
																								Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																								MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"services": schema.ListAttribute{
																								Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
																								MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

																					"secret_class": schema.StringAttribute{
																						Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																						MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"host": schema.StringAttribute{
																				Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
																				MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"port": schema.Int64Attribute{
																				Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
																				MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.Int64{
																					int64validator.AtLeast(0),
																				},
																			},

																			"region": schema.SingleNestedAttribute{
																				Description:         "Bucket region used for signing headers (sigv4). This defaults to 'us-east-1' which is compatible with other implementations such as Minio. WARNING: Some products use the Hadoop S3 implementation which falls back to us-east-2.",
																				MarkdownDescription: "Bucket region used for signing headers (sigv4). This defaults to 'us-east-1' which is compatible with other implementations such as Minio. WARNING: Some products use the Hadoop S3 implementation which falls back to us-east-2.",
																				Attributes: map[string]schema.Attribute{
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

																			"tls": schema.SingleNestedAttribute{
																				Description:         "Use a TLS connection. If not specified no TLS will be used.",
																				MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
																				Attributes: map[string]schema.Attribute{
																					"verification": schema.SingleNestedAttribute{
																						Description:         "The verification method used to verify the certificates of the server and/or the client.",
																						MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
																						Attributes: map[string]schema.Attribute{
																							"none": schema.MapAttribute{
																								Description:         "Use TLS but don't verify certificates.",
																								MarkdownDescription: "Use TLS but don't verify certificates.",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"server": schema.SingleNestedAttribute{
																								Description:         "Use TLS and a CA certificate to verify the server.",
																								MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
																								Attributes: map[string]schema.Attribute{
																									"ca_cert": schema.SingleNestedAttribute{
																										Description:         "CA cert to verify the server.",
																										MarkdownDescription: "CA cert to verify the server.",
																										Attributes: map[string]schema.Attribute{
																											"secret_class": schema.StringAttribute{
																												Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																												MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"web_pki": schema.MapAttribute{
																												Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																												MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
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

																	"reference": schema.StringAttribute{
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

															"external_id": schema.StringAttribute{
																Description:         "External ID for the IAM role trust policy.",
																MarkdownDescription: "External ID for the IAM role trust policy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"iam_role": schema.StringAttribute{
																Description:         "IAM role to assume for S3 access.",
																MarkdownDescription: "IAM role to assume for S3 access.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"max_error_retries": schema.Int64Attribute{
																Description:         "Maximum number of times the S3 client should retry a request.",
																MarkdownDescription: "Maximum number of times the S3 client should retry a request.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},

															"upload_part_size": schema.StringAttribute{
																Description:         "Part data size for S3 multi-part upload.",
																MarkdownDescription: "Part data size for S3 multi-part upload.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"sink_buffer_pool_min_size": schema.Int64Attribute{
														Description:         "The minimum buffer pool size for an exchange sink. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														MarkdownDescription: "The minimum buffer pool size for an exchange sink. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"sink_buffers_per_partition": schema.Int64Attribute{
														Description:         "The number of buffers per partition in the buffer pool. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														MarkdownDescription: "The number of buffers per partition in the buffer pool. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"sink_max_file_size": schema.StringAttribute{
														Description:         "Max data size of files written by exchange sinks.",
														MarkdownDescription: "Max data size of files written by exchange sinks.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_concurrent_readers": schema.Int64Attribute{
														Description:         "Number of concurrent readers to read from spooling storage. The larger the number of concurrent readers, the larger the read parallelism and memory usage.",
														MarkdownDescription: "Number of concurrent readers to read from spooling storage. The larger the number of concurrent readers, the larger the read parallelism and memory usage.",
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

											"retry_attempts": schema.Int64Attribute{
												Description:         "Maximum number of times Trino may attempt to retry a query before declaring it failed.",
												MarkdownDescription: "Maximum number of times Trino may attempt to retry a query before declaring it failed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"retry_delay_scale_factor": schema.Float64Attribute{
												Description:         "Factor by which retry delay is increased on each query failure.",
												MarkdownDescription: "Factor by which retry delay is increased on each query failure.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_initial_delay": schema.StringAttribute{
												Description:         "Minimum time that a failed query must wait before it is retried.",
												MarkdownDescription: "Minimum time that a failed query must wait before it is retried.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_max_delay": schema.StringAttribute{
												Description:         "Maximum time that a failed query must wait before it is retried.",
												MarkdownDescription: "Maximum time that a failed query must wait before it is retried.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"task": schema.SingleNestedAttribute{
										Description:         "Task-level fault tolerant execution. Retries individual tasks on failure (requires exchange manager).",
										MarkdownDescription: "Task-level fault tolerant execution. Retries individual tasks on failure (requires exchange manager).",
										Attributes: map[string]schema.Attribute{
											"exchange_deduplication_buffer_size": schema.StringAttribute{
												Description:         "Data size of the coordinator's in-memory buffer used to store output of query stages.",
												MarkdownDescription: "Data size of the coordinator's in-memory buffer used to store output of query stages.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"exchange_manager": schema.SingleNestedAttribute{
												Description:         "Exchange manager configuration for spooling intermediate data during fault tolerant execution. Required for Task retry policy.",
												MarkdownDescription: "Exchange manager configuration for spooling intermediate data during fault tolerant execution. Required for Task retry policy.",
												Attributes: map[string]schema.Attribute{
													"config_overrides": schema.MapAttribute{
														Description:         "The 'configOverrides' allow overriding arbitrary exchange manager properties.",
														MarkdownDescription: "The 'configOverrides' allow overriding arbitrary exchange manager properties.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"encryption_enabled": schema.BoolAttribute{
														Description:         "Whether to enable encryption of spooling data.",
														MarkdownDescription: "Whether to enable encryption of spooling data.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hdfs": schema.SingleNestedAttribute{
														Description:         "HDFS-based exchange manager.",
														MarkdownDescription: "HDFS-based exchange manager.",
														Attributes: map[string]schema.Attribute{
															"base_directories": schema.ListAttribute{
																Description:         "HDFS URIs for spooling data.",
																MarkdownDescription: "HDFS URIs for spooling data.",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"block_size": schema.StringAttribute{
																Description:         "Block data size for HDFS storage.",
																MarkdownDescription: "Block data size for HDFS storage.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"hdfs": schema.SingleNestedAttribute{
																Description:         "HDFS connection configuration.",
																MarkdownDescription: "HDFS connection configuration.",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.StringAttribute{
																		Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
																		MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"skip_directory_scheme_validation": schema.BoolAttribute{
																Description:         "Skip directory scheme validation to support Hadoop-compatible file systems.",
																MarkdownDescription: "Skip directory scheme validation to support Hadoop-compatible file systems.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"local": schema.SingleNestedAttribute{
														Description:         "Local filesystem storage (not recommended for production).",
														MarkdownDescription: "Local filesystem storage (not recommended for production).",
														Attributes: map[string]schema.Attribute{
															"base_directories": schema.ListAttribute{
																Description:         "Local filesystem paths for exchange storage.",
																MarkdownDescription: "Local filesystem paths for exchange storage.",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"s3": schema.SingleNestedAttribute{
														Description:         "S3-compatible storage configuration.",
														MarkdownDescription: "S3-compatible storage configuration.",
														Attributes: map[string]schema.Attribute{
															"base_directories": schema.ListAttribute{
																Description:         "S3 bucket URIs for spooling data (e.g., s3://bucket1,s3://bucket2).",
																MarkdownDescription: "S3 bucket URIs for spooling data (e.g., s3://bucket1,s3://bucket2).",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"connection": schema.SingleNestedAttribute{
																Description:         "S3 connection configuration. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
																MarkdownDescription: "S3 connection configuration. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
																Attributes: map[string]schema.Attribute{
																	"inline": schema.SingleNestedAttribute{
																		Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
																		MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
																		Attributes: map[string]schema.Attribute{
																			"access_style": schema.StringAttribute{
																				Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
																				MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("Path", "VirtualHosted"),
																				},
																			},

																			"credentials": schema.SingleNestedAttribute{
																				Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
																				MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
																				Attributes: map[string]schema.Attribute{
																					"scope": schema.SingleNestedAttribute{
																						Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																						MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																						Attributes: map[string]schema.Attribute{
																							"listener_volumes": schema.ListAttribute{
																								Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																								MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"node": schema.BoolAttribute{
																								Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																								MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"pod": schema.BoolAttribute{
																								Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																								MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"services": schema.ListAttribute{
																								Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
																								MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

																					"secret_class": schema.StringAttribute{
																						Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																						MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"host": schema.StringAttribute{
																				Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
																				MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"port": schema.Int64Attribute{
																				Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
																				MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.Int64{
																					int64validator.AtLeast(0),
																				},
																			},

																			"region": schema.SingleNestedAttribute{
																				Description:         "Bucket region used for signing headers (sigv4). This defaults to 'us-east-1' which is compatible with other implementations such as Minio. WARNING: Some products use the Hadoop S3 implementation which falls back to us-east-2.",
																				MarkdownDescription: "Bucket region used for signing headers (sigv4). This defaults to 'us-east-1' which is compatible with other implementations such as Minio. WARNING: Some products use the Hadoop S3 implementation which falls back to us-east-2.",
																				Attributes: map[string]schema.Attribute{
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

																			"tls": schema.SingleNestedAttribute{
																				Description:         "Use a TLS connection. If not specified no TLS will be used.",
																				MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
																				Attributes: map[string]schema.Attribute{
																					"verification": schema.SingleNestedAttribute{
																						Description:         "The verification method used to verify the certificates of the server and/or the client.",
																						MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
																						Attributes: map[string]schema.Attribute{
																							"none": schema.MapAttribute{
																								Description:         "Use TLS but don't verify certificates.",
																								MarkdownDescription: "Use TLS but don't verify certificates.",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"server": schema.SingleNestedAttribute{
																								Description:         "Use TLS and a CA certificate to verify the server.",
																								MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
																								Attributes: map[string]schema.Attribute{
																									"ca_cert": schema.SingleNestedAttribute{
																										Description:         "CA cert to verify the server.",
																										MarkdownDescription: "CA cert to verify the server.",
																										Attributes: map[string]schema.Attribute{
																											"secret_class": schema.StringAttribute{
																												Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																												MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"web_pki": schema.MapAttribute{
																												Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																												MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
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

																	"reference": schema.StringAttribute{
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

															"external_id": schema.StringAttribute{
																Description:         "External ID for the IAM role trust policy.",
																MarkdownDescription: "External ID for the IAM role trust policy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"iam_role": schema.StringAttribute{
																Description:         "IAM role to assume for S3 access.",
																MarkdownDescription: "IAM role to assume for S3 access.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"max_error_retries": schema.Int64Attribute{
																Description:         "Maximum number of times the S3 client should retry a request.",
																MarkdownDescription: "Maximum number of times the S3 client should retry a request.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},

															"upload_part_size": schema.StringAttribute{
																Description:         "Part data size for S3 multi-part upload.",
																MarkdownDescription: "Part data size for S3 multi-part upload.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"sink_buffer_pool_min_size": schema.Int64Attribute{
														Description:         "The minimum buffer pool size for an exchange sink. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														MarkdownDescription: "The minimum buffer pool size for an exchange sink. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"sink_buffers_per_partition": schema.Int64Attribute{
														Description:         "The number of buffers per partition in the buffer pool. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														MarkdownDescription: "The number of buffers per partition in the buffer pool. The larger the buffer pool size, the larger the write parallelism and memory usage.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"sink_max_file_size": schema.StringAttribute{
														Description:         "Max data size of files written by exchange sinks.",
														MarkdownDescription: "Max data size of files written by exchange sinks.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_concurrent_readers": schema.Int64Attribute{
														Description:         "Number of concurrent readers to read from spooling storage. The larger the number of concurrent readers, the larger the read parallelism and memory usage.",
														MarkdownDescription: "Number of concurrent readers to read from spooling storage. The larger the number of concurrent readers, the larger the read parallelism and memory usage.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"retry_attempts_per_task": schema.Int64Attribute{
												Description:         "Maximum number of times Trino may attempt to retry a single task before declaring the query failed.",
												MarkdownDescription: "Maximum number of times Trino may attempt to retry a single task before declaring the query failed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"retry_delay_scale_factor": schema.Float64Attribute{
												Description:         "Factor by which retry delay is increased on each task failure.",
												MarkdownDescription: "Factor by which retry delay is increased on each task failure.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_initial_delay": schema.StringAttribute{
												Description:         "Minimum time that a failed task must wait before it is retried.",
												MarkdownDescription: "Minimum time that a failed task must wait before it is retried.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_max_delay": schema.StringAttribute{
												Description:         "Maximum time that a failed task must wait before it is retried.",
												MarkdownDescription: "Maximum time that a failed task must wait before it is retried.",
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
								Description:         "TLS configuration options for server and internal communication.",
								MarkdownDescription: "TLS configuration options for server and internal communication.",
								Attributes: map[string]schema.Attribute{
									"internal_secret_class": schema.StringAttribute{
										Description:         "Only affects internal communication. Use mutual verification between Trino nodes This setting controls: - Which cert the servers should use to authenticate themselves against other servers - Which ca.crt to use when validating the other server",
										MarkdownDescription: "Only affects internal communication. Use mutual verification between Trino nodes This setting controls: - Which cert the servers should use to authenticate themselves against other servers - Which ca.crt to use when validating the other server",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server_secret_class": schema.StringAttribute{
										Description:         "Only affects client connections. This setting controls: - If TLS encryption is used at all - Which cert the servers should use to authenticate themselves against the client",
										MarkdownDescription: "Only affects client connections. This setting controls: - If TLS encryption is used at all - Which cert the servers should use to authenticate themselves against the client",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"vector_aggregator_config_map_name": schema.StringAttribute{
								Description:         "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
								MarkdownDescription: "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"cluster_operation": schema.SingleNestedAttribute{
						Description:         "[Cluster operations](https://docs.stackable.tech/home/nightly/concepts/operations/cluster_operations) properties, allow stopping the product instance as well as pausing reconciliation.",
						MarkdownDescription: "[Cluster operations](https://docs.stackable.tech/home/nightly/concepts/operations/cluster_operations) properties, allow stopping the product instance as well as pausing reconciliation.",
						Attributes: map[string]schema.Attribute{
							"reconciliation_paused": schema.BoolAttribute{
								Description:         "Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.",
								MarkdownDescription: "Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stopped": schema.BoolAttribute{
								Description:         "Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.",
								MarkdownDescription: "Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"coordinators": schema.SingleNestedAttribute{
						Description:         "This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups).",
						MarkdownDescription: "This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups).",
						Attributes: map[string]schema.Attribute{
							"cli_overrides": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

									"graceful_shutdown_timeout": schema.StringAttribute{
										Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										Attributes: map[string]schema.Attribute{
											"containers": schema.SingleNestedAttribute{
												Description:         "Log configuration per container.",
												MarkdownDescription: "Log configuration per container.",
												Attributes: map[string]schema.Attribute{
													"console": schema.SingleNestedAttribute{
														Description:         "Configuration for the console appender",
														MarkdownDescription: "Configuration for the console appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom": schema.SingleNestedAttribute{
														Description:         "Custom log configuration provided in a ConfigMap",
														MarkdownDescription: "Custom log configuration provided in a ConfigMap",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.StringAttribute{
																Description:         "ConfigMap containing the log configuration files",
																MarkdownDescription: "ConfigMap containing the log configuration files",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.SingleNestedAttribute{
														Description:         "Configuration for the file appender",
														MarkdownDescription: "Configuration for the file appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"loggers": schema.SingleNestedAttribute{
														Description:         "Configuration per logger",
														MarkdownDescription: "Configuration per logger",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

											"enable_vector_agent": schema.BoolAttribute{
												Description:         "Wether or not to deploy a container with the Vector log agent.",
												MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_max_memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_max_memory_per_node": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requested_secret_lifetime": schema.StringAttribute{
										Description:         "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
										MarkdownDescription: "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max": schema.StringAttribute{
														Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min": schema.StringAttribute{
														Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"limit": schema.StringAttribute{
														Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_limits": schema.MapAttribute{
														Description:         "Additional options that can be specified.",
														MarkdownDescription: "Additional options that can be specified.",
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

											"storage": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_overrides": schema.MapAttribute{
								Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_overrides": schema.MapAttribute{
								Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jvm_argument_overrides": schema.SingleNestedAttribute{
								Description:         "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
								MarkdownDescription: "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
								Attributes: map[string]schema.Attribute{
									"add": schema.ListAttribute{
										Description:         "JVM arguments to be added",
										MarkdownDescription: "JVM arguments to be added",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remove": schema.ListAttribute{
										Description:         "JVM arguments to be removed by exact match",
										MarkdownDescription: "JVM arguments to be removed by exact match",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remove_regex": schema.ListAttribute{
										Description:         "JVM arguments matching any of this regexes will be removed",
										MarkdownDescription: "JVM arguments matching any of this regexes will be removed",
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

							"pod_overrides": schema.MapAttribute{
								Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_config": schema.SingleNestedAttribute{
								Description:         "This is a product-agnostic RoleConfig, which is sufficient for most of the products.",
								MarkdownDescription: "This is a product-agnostic RoleConfig, which is sufficient for most of the products.",
								Attributes: map[string]schema.Attribute{
									"listener_class": schema.StringAttribute{
										Description:         "This field controls which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) is used to expose the coordinator.",
										MarkdownDescription: "This field controls which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) is used to expose the coordinator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_disruption_budget": schema.SingleNestedAttribute{
										Description:         "This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions).",
										MarkdownDescription: "This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions).",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.",
												MarkdownDescription: "Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.Int64Attribute{
												Description:         "The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.",
												MarkdownDescription: "The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_groups": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cli_overrides": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.SingleNestedAttribute{
												Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
												MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
												Attributes: map[string]schema.Attribute{
													"node_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_selector": schema.MapAttribute{
														Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_anti_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

											"graceful_shutdown_timeout": schema.StringAttribute{
												Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"logging": schema.SingleNestedAttribute{
												Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
												MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
												Attributes: map[string]schema.Attribute{
													"containers": schema.SingleNestedAttribute{
														Description:         "Log configuration per container.",
														MarkdownDescription: "Log configuration per container.",
														Attributes: map[string]schema.Attribute{
															"console": schema.SingleNestedAttribute{
																Description:         "Configuration for the console appender",
																MarkdownDescription: "Configuration for the console appender",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"custom": schema.SingleNestedAttribute{
																Description:         "Custom log configuration provided in a ConfigMap",
																MarkdownDescription: "Custom log configuration provided in a ConfigMap",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.StringAttribute{
																		Description:         "ConfigMap containing the log configuration files",
																		MarkdownDescription: "ConfigMap containing the log configuration files",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"file": schema.SingleNestedAttribute{
																Description:         "Configuration for the file appender",
																MarkdownDescription: "Configuration for the file appender",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"loggers": schema.SingleNestedAttribute{
																Description:         "Configuration per logger",
																MarkdownDescription: "Configuration per logger",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

													"enable_vector_agent": schema.BoolAttribute{
														Description:         "Wether or not to deploy a container with the Vector log agent.",
														MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_max_memory": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_max_memory_per_node": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requested_secret_lifetime": schema.StringAttribute{
												Description:         "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
												MarkdownDescription: "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
												MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"max": schema.StringAttribute{
																Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"min": schema.StringAttribute{
																Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"limit": schema.StringAttribute{
																Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
																MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"runtime_limits": schema.MapAttribute{
																Description:         "Additional options that can be specified.",
																MarkdownDescription: "Additional options that can be specified.",
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

													"storage": schema.MapAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_overrides": schema.MapAttribute{
										Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
										MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env_overrides": schema.MapAttribute{
										Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
										MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jvm_argument_overrides": schema.SingleNestedAttribute{
										Description:         "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
										MarkdownDescription: "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
										Attributes: map[string]schema.Attribute{
											"add": schema.ListAttribute{
												Description:         "JVM arguments to be added",
												MarkdownDescription: "JVM arguments to be added",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remove": schema.ListAttribute{
												Description:         "JVM arguments to be removed by exact match",
												MarkdownDescription: "JVM arguments to be removed by exact match",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remove_regex": schema.ListAttribute{
												Description:         "JVM arguments matching any of this regexes will be removed",
												MarkdownDescription: "JVM arguments matching any of this regexes will be removed",
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

									"pod_overrides": schema.MapAttribute{
										Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
										MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
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

					"image": schema.SingleNestedAttribute{
						Description:         "Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details.",
						MarkdownDescription: "Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.StringAttribute{
								Description:         "Overwrite the docker image. Specify the full docker image name, e.g. 'oci.stackable.tech/sdp/superset:1.4.1-stackable2.1.0'",
								MarkdownDescription: "Overwrite the docker image. Specify the full docker image name, e.g. 'oci.stackable.tech/sdp/superset:1.4.1-stackable2.1.0'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"product_version": schema.StringAttribute{
								Description:         "Version of the product, e.g. '1.4.1'.",
								MarkdownDescription: "Version of the product, e.g. '1.4.1'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pull_policy": schema.StringAttribute{
								Description:         "[Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.",
								MarkdownDescription: "[Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"pull_secrets": schema.ListNestedAttribute{
								Description:         "[Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry.",
								MarkdownDescription: "[Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"repo": schema.StringAttribute{
								Description:         "Name of the docker repo, e.g. 'oci.stackable.tech/sdp'",
								MarkdownDescription: "Name of the docker repo, e.g. 'oci.stackable.tech/sdp'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stackable_version": schema.StringAttribute{
								Description:         "Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.",
								MarkdownDescription: "Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"workers": schema.SingleNestedAttribute{
						Description:         "This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups).",
						MarkdownDescription: "This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups).",
						Attributes: map[string]schema.Attribute{
							"cli_overrides": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

									"graceful_shutdown_timeout": schema.StringAttribute{
										Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										Attributes: map[string]schema.Attribute{
											"containers": schema.SingleNestedAttribute{
												Description:         "Log configuration per container.",
												MarkdownDescription: "Log configuration per container.",
												Attributes: map[string]schema.Attribute{
													"console": schema.SingleNestedAttribute{
														Description:         "Configuration for the console appender",
														MarkdownDescription: "Configuration for the console appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom": schema.SingleNestedAttribute{
														Description:         "Custom log configuration provided in a ConfigMap",
														MarkdownDescription: "Custom log configuration provided in a ConfigMap",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.StringAttribute{
																Description:         "ConfigMap containing the log configuration files",
																MarkdownDescription: "ConfigMap containing the log configuration files",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.SingleNestedAttribute{
														Description:         "Configuration for the file appender",
														MarkdownDescription: "Configuration for the file appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"loggers": schema.SingleNestedAttribute{
														Description:         "Configuration per logger",
														MarkdownDescription: "Configuration per logger",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

											"enable_vector_agent": schema.BoolAttribute{
												Description:         "Wether or not to deploy a container with the Vector log agent.",
												MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_max_memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_max_memory_per_node": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requested_secret_lifetime": schema.StringAttribute{
										Description:         "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
										MarkdownDescription: "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max": schema.StringAttribute{
														Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min": schema.StringAttribute{
														Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"limit": schema.StringAttribute{
														Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_limits": schema.MapAttribute{
														Description:         "Additional options that can be specified.",
														MarkdownDescription: "Additional options that can be specified.",
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

											"storage": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_overrides": schema.MapAttribute{
								Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_overrides": schema.MapAttribute{
								Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jvm_argument_overrides": schema.SingleNestedAttribute{
								Description:         "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
								MarkdownDescription: "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
								Attributes: map[string]schema.Attribute{
									"add": schema.ListAttribute{
										Description:         "JVM arguments to be added",
										MarkdownDescription: "JVM arguments to be added",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remove": schema.ListAttribute{
										Description:         "JVM arguments to be removed by exact match",
										MarkdownDescription: "JVM arguments to be removed by exact match",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remove_regex": schema.ListAttribute{
										Description:         "JVM arguments matching any of this regexes will be removed",
										MarkdownDescription: "JVM arguments matching any of this regexes will be removed",
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

							"pod_overrides": schema.MapAttribute{
								Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_config": schema.SingleNestedAttribute{
								Description:         "This is a product-agnostic RoleConfig, which is sufficient for most of the products.",
								MarkdownDescription: "This is a product-agnostic RoleConfig, which is sufficient for most of the products.",
								Attributes: map[string]schema.Attribute{
									"pod_disruption_budget": schema.SingleNestedAttribute{
										Description:         "This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions).",
										MarkdownDescription: "This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions).",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.",
												MarkdownDescription: "Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.Int64Attribute{
												Description:         "The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.",
												MarkdownDescription: "The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_groups": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cli_overrides": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.SingleNestedAttribute{
												Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
												MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
												Attributes: map[string]schema.Attribute{
													"node_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_selector": schema.MapAttribute{
														Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_anti_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

											"graceful_shutdown_timeout": schema.StringAttribute{
												Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"logging": schema.SingleNestedAttribute{
												Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
												MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
												Attributes: map[string]schema.Attribute{
													"containers": schema.SingleNestedAttribute{
														Description:         "Log configuration per container.",
														MarkdownDescription: "Log configuration per container.",
														Attributes: map[string]schema.Attribute{
															"console": schema.SingleNestedAttribute{
																Description:         "Configuration for the console appender",
																MarkdownDescription: "Configuration for the console appender",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"custom": schema.SingleNestedAttribute{
																Description:         "Custom log configuration provided in a ConfigMap",
																MarkdownDescription: "Custom log configuration provided in a ConfigMap",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.StringAttribute{
																		Description:         "ConfigMap containing the log configuration files",
																		MarkdownDescription: "ConfigMap containing the log configuration files",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"file": schema.SingleNestedAttribute{
																Description:         "Configuration for the file appender",
																MarkdownDescription: "Configuration for the file appender",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"loggers": schema.SingleNestedAttribute{
																Description:         "Configuration per logger",
																MarkdownDescription: "Configuration per logger",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

													"enable_vector_agent": schema.BoolAttribute{
														Description:         "Wether or not to deploy a container with the Vector log agent.",
														MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_max_memory": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_max_memory_per_node": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requested_secret_lifetime": schema.StringAttribute{
												Description:         "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
												MarkdownDescription: "Request secret (currently only autoTls certificates) lifetime from the secret operator, e.g. '7d', or '30d'. This can be shortened by the 'maxCertificateLifetime' setting on the SecretClass issuing the TLS certificate. Defaults to '15d' for coordinators (as currently a restart kills all running queries) and '1d' for workers.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
												MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"max": schema.StringAttribute{
																Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"min": schema.StringAttribute{
																Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"limit": schema.StringAttribute{
																Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
																MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"runtime_limits": schema.MapAttribute{
																Description:         "Additional options that can be specified.",
																MarkdownDescription: "Additional options that can be specified.",
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

													"storage": schema.MapAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_overrides": schema.MapAttribute{
										Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
										MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env_overrides": schema.MapAttribute{
										Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
										MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jvm_argument_overrides": schema.SingleNestedAttribute{
										Description:         "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
										MarkdownDescription: "Allows overriding JVM arguments. Please read on the [JVM argument overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#jvm-argument-overrides) for details on the usage.",
										Attributes: map[string]schema.Attribute{
											"add": schema.ListAttribute{
												Description:         "JVM arguments to be added",
												MarkdownDescription: "JVM arguments to be added",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remove": schema.ListAttribute{
												Description:         "JVM arguments to be removed by exact match",
												MarkdownDescription: "JVM arguments to be removed by exact match",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remove_regex": schema.ListAttribute{
												Description:         "JVM arguments matching any of this regexes will be removed",
												MarkdownDescription: "JVM arguments matching any of this regexes will be removed",
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

									"pod_overrides": schema.MapAttribute{
										Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
										MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TrinoStackableTechTrinoClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_trino_stackable_tech_trino_cluster_v1alpha1_manifest")

	var model TrinoStackableTechTrinoClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("trino.stackable.tech/v1alpha1")
	model.Kind = pointer.String("TrinoCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
