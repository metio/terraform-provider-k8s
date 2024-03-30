/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package temporal_io_v1beta1

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
	_ datasource.DataSource = &TemporalIoTemporalClusterV1Beta1Manifest{}
)

func NewTemporalIoTemporalClusterV1Beta1Manifest() datasource.DataSource {
	return &TemporalIoTemporalClusterV1Beta1Manifest{}
}

type TemporalIoTemporalClusterV1Beta1Manifest struct{}

type TemporalIoTemporalClusterV1Beta1ManifestData struct {
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
		Admintools *struct {
			Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Image     *string `tfsdk:"image" json:"image,omitempty"`
			Overrides *struct {
				Deployment *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec *struct {
						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
								Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"template" json:"template,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
			} `tfsdk:"overrides" json:"overrides,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"admintools" json:"admintools,omitempty"`
		Archival *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			History *struct {
				EnableRead *bool   `tfsdk:"enable_read" json:"enableRead,omitempty"`
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Path       *string `tfsdk:"path" json:"path,omitempty"`
				Paused     *bool   `tfsdk:"paused" json:"paused,omitempty"`
			} `tfsdk:"history" json:"history,omitempty"`
			Provider *struct {
				Filestore *struct {
					DirPermissions  *string `tfsdk:"dir_permissions" json:"dirPermissions,omitempty"`
					FilePermissions *string `tfsdk:"file_permissions" json:"filePermissions,omitempty"`
				} `tfsdk:"filestore" json:"filestore,omitempty"`
				Gcs *struct {
					CredentialsRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"credentials_ref" json:"credentialsRef,omitempty"`
				} `tfsdk:"gcs" json:"gcs,omitempty"`
				S3 *struct {
					Credentials *struct {
						AccessKeyIdRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"access_key_id_ref" json:"accessKeyIdRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"credentials" json:"credentials,omitempty"`
					Endpoint         *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
					Region           *string `tfsdk:"region" json:"region,omitempty"`
					RoleName         *string `tfsdk:"role_name" json:"roleName,omitempty"`
					S3ForcePathStyle *bool   `tfsdk:"s3_force_path_style" json:"s3ForcePathStyle,omitempty"`
				} `tfsdk:"s3" json:"s3,omitempty"`
			} `tfsdk:"provider" json:"provider,omitempty"`
			Visibility *struct {
				EnableRead *bool   `tfsdk:"enable_read" json:"enableRead,omitempty"`
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Path       *string `tfsdk:"path" json:"path,omitempty"`
				Paused     *bool   `tfsdk:"paused" json:"paused,omitempty"`
			} `tfsdk:"visibility" json:"visibility,omitempty"`
		} `tfsdk:"archival" json:"archival,omitempty"`
		Authorization *struct {
			Authorizer     *string `tfsdk:"authorizer" json:"authorizer,omitempty"`
			ClaimMapper    *string `tfsdk:"claim_mapper" json:"claimMapper,omitempty"`
			JwtKeyProvider *struct {
				KeySourceURIs   *[]string `tfsdk:"key_source_ur_is" json:"keySourceURIs,omitempty"`
				RefreshInterval *string   `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
			} `tfsdk:"jwt_key_provider" json:"jwtKeyProvider,omitempty"`
			PermissionsClaimName *string `tfsdk:"permissions_claim_name" json:"permissionsClaimName,omitempty"`
		} `tfsdk:"authorization" json:"authorization,omitempty"`
		DynamicConfig *struct {
			PollInterval *string            `tfsdk:"poll_interval" json:"pollInterval,omitempty"`
			Values       *map[string]string `tfsdk:"values" json:"values,omitempty"`
		} `tfsdk:"dynamic_config" json:"dynamicConfig,omitempty"`
		Image            *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		JobInitContainers *[]map[string]string `tfsdk:"job_init_containers" json:"jobInitContainers,omitempty"`
		JobResources      *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"job_resources" json:"jobResources,omitempty"`
		JobTtlSecondsAfterFinished *int64 `tfsdk:"job_ttl_seconds_after_finished" json:"jobTtlSecondsAfterFinished,omitempty"`
		Log                        *struct {
			Development *bool   `tfsdk:"development" json:"development,omitempty"`
			Format      *string `tfsdk:"format" json:"format,omitempty"`
			Level       *string `tfsdk:"level" json:"level,omitempty"`
			OutputFile  *string `tfsdk:"output_file" json:"outputFile,omitempty"`
			Stdout      *bool   `tfsdk:"stdout" json:"stdout,omitempty"`
		} `tfsdk:"log" json:"log,omitempty"`
		MTLS *struct {
			CertificatesDuration *struct {
				ClientCertificates          *string `tfsdk:"client_certificates" json:"clientCertificates,omitempty"`
				FrontendCertificate         *string `tfsdk:"frontend_certificate" json:"frontendCertificate,omitempty"`
				IntermediateCAsCertificates *string `tfsdk:"intermediate_c_as_certificates" json:"intermediateCAsCertificates,omitempty"`
				InternodeCertificate        *string `tfsdk:"internode_certificate" json:"internodeCertificate,omitempty"`
				RootCACertificate           *string `tfsdk:"root_ca_certificate" json:"rootCACertificate,omitempty"`
			} `tfsdk:"certificates_duration" json:"certificatesDuration,omitempty"`
			Frontend *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"frontend" json:"frontend,omitempty"`
			Internode *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"internode" json:"internode,omitempty"`
			Provider        *string `tfsdk:"provider" json:"provider,omitempty"`
			RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
			RenewBefore     *string `tfsdk:"renew_before" json:"renewBefore,omitempty"`
		} `tfsdk:"m_tls" json:"mTLS,omitempty"`
		Metrics *struct {
			Enabled                    *bool                `tfsdk:"enabled" json:"enabled,omitempty"`
			ExcludeTags                *map[string][]string `tfsdk:"exclude_tags" json:"excludeTags,omitempty"`
			PerUnitHistogramBoundaries *map[string][]string `tfsdk:"per_unit_histogram_boundaries" json:"perUnitHistogramBoundaries,omitempty"`
			Prefix                     *string              `tfsdk:"prefix" json:"prefix,omitempty"`
			Prometheus                 *struct {
				ListenAddress *string `tfsdk:"listen_address" json:"listenAddress,omitempty"`
				ListenPort    *int64  `tfsdk:"listen_port" json:"listenPort,omitempty"`
				ScrapeConfig  *struct {
					Annotations    *bool `tfsdk:"annotations" json:"annotations,omitempty"`
					ServiceMonitor *struct {
						Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
						Labels            *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						MetricRelabelings *[]struct {
							Action       *string   `tfsdk:"action" json:"action,omitempty"`
							Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
							Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
							Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
							Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
							SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
							TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
						} `tfsdk:"metric_relabelings" json:"metricRelabelings,omitempty"`
						Override *struct {
							AttachMetadata *struct {
								Node *bool `tfsdk:"node" json:"node,omitempty"`
							} `tfsdk:"attach_metadata" json:"attachMetadata,omitempty"`
							Endpoints *[]struct {
								Authorization *struct {
									Credentials *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"credentials" json:"credentials,omitempty"`
									Type *string `tfsdk:"type" json:"type,omitempty"`
								} `tfsdk:"authorization" json:"authorization,omitempty"`
								BasicAuth *struct {
									Password *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"password" json:"password,omitempty"`
									Username *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"username" json:"username,omitempty"`
								} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
								BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
								BearerTokenSecret *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
								EnableHttp2       *bool   `tfsdk:"enable_http2" json:"enableHttp2,omitempty"`
								FilterRunning     *bool   `tfsdk:"filter_running" json:"filterRunning,omitempty"`
								FollowRedirects   *bool   `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
								HonorLabels       *bool   `tfsdk:"honor_labels" json:"honorLabels,omitempty"`
								HonorTimestamps   *bool   `tfsdk:"honor_timestamps" json:"honorTimestamps,omitempty"`
								Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
								MetricRelabelings *[]struct {
									Action       *string   `tfsdk:"action" json:"action,omitempty"`
									Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
									Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
									Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
									Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
									SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
									TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
								} `tfsdk:"metric_relabelings" json:"metricRelabelings,omitempty"`
								Oauth2 *struct {
									ClientId *struct {
										ConfigMap *struct {
											Key      *string `tfsdk:"key" json:"key,omitempty"`
											Name     *string `tfsdk:"name" json:"name,omitempty"`
											Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
										} `tfsdk:"config_map" json:"configMap,omitempty"`
										Secret *struct {
											Key      *string `tfsdk:"key" json:"key,omitempty"`
											Name     *string `tfsdk:"name" json:"name,omitempty"`
											Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
										} `tfsdk:"secret" json:"secret,omitempty"`
									} `tfsdk:"client_id" json:"clientId,omitempty"`
									ClientSecret *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
									EndpointParams *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
									Scopes         *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
									TokenUrl       *string            `tfsdk:"token_url" json:"tokenUrl,omitempty"`
								} `tfsdk:"oauth2" json:"oauth2,omitempty"`
								Params      *map[string][]string `tfsdk:"params" json:"params,omitempty"`
								Path        *string              `tfsdk:"path" json:"path,omitempty"`
								Port        *string              `tfsdk:"port" json:"port,omitempty"`
								ProxyUrl    *string              `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
								Relabelings *[]struct {
									Action       *string   `tfsdk:"action" json:"action,omitempty"`
									Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
									Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
									Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
									Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
									SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
									TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
								} `tfsdk:"relabelings" json:"relabelings,omitempty"`
								Scheme        *string `tfsdk:"scheme" json:"scheme,omitempty"`
								ScrapeTimeout *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
								TargetPort    *string `tfsdk:"target_port" json:"targetPort,omitempty"`
								TlsConfig     *struct {
									Ca *struct {
										ConfigMap *struct {
											Key      *string `tfsdk:"key" json:"key,omitempty"`
											Name     *string `tfsdk:"name" json:"name,omitempty"`
											Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
										} `tfsdk:"config_map" json:"configMap,omitempty"`
										Secret *struct {
											Key      *string `tfsdk:"key" json:"key,omitempty"`
											Name     *string `tfsdk:"name" json:"name,omitempty"`
											Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
										} `tfsdk:"secret" json:"secret,omitempty"`
									} `tfsdk:"ca" json:"ca,omitempty"`
									CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
									Cert   *struct {
										ConfigMap *struct {
											Key      *string `tfsdk:"key" json:"key,omitempty"`
											Name     *string `tfsdk:"name" json:"name,omitempty"`
											Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
										} `tfsdk:"config_map" json:"configMap,omitempty"`
										Secret *struct {
											Key      *string `tfsdk:"key" json:"key,omitempty"`
											Name     *string `tfsdk:"name" json:"name,omitempty"`
											Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
										} `tfsdk:"secret" json:"secret,omitempty"`
									} `tfsdk:"cert" json:"cert,omitempty"`
									CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
									InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
									KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
									KeySecret          *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"key_secret" json:"keySecret,omitempty"`
									ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
								} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
								TrackTimestampsStaleness *bool `tfsdk:"track_timestamps_staleness" json:"trackTimestampsStaleness,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							JobLabel              *string `tfsdk:"job_label" json:"jobLabel,omitempty"`
							KeepDroppedTargets    *int64  `tfsdk:"keep_dropped_targets" json:"keepDroppedTargets,omitempty"`
							LabelLimit            *int64  `tfsdk:"label_limit" json:"labelLimit,omitempty"`
							LabelNameLengthLimit  *int64  `tfsdk:"label_name_length_limit" json:"labelNameLengthLimit,omitempty"`
							LabelValueLengthLimit *int64  `tfsdk:"label_value_length_limit" json:"labelValueLengthLimit,omitempty"`
							NamespaceSelector     *struct {
								Any        *bool     `tfsdk:"any" json:"any,omitempty"`
								MatchNames *[]string `tfsdk:"match_names" json:"matchNames,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							PodTargetLabels *[]string `tfsdk:"pod_target_labels" json:"podTargetLabels,omitempty"`
							SampleLimit     *int64    `tfsdk:"sample_limit" json:"sampleLimit,omitempty"`
							Selector        *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
							TargetLabels *[]string `tfsdk:"target_labels" json:"targetLabels,omitempty"`
							TargetLimit  *int64    `tfsdk:"target_limit" json:"targetLimit,omitempty"`
						} `tfsdk:"override" json:"override,omitempty"`
					} `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
				} `tfsdk:"scrape_config" json:"scrapeConfig,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		NumHistoryShards *int64 `tfsdk:"num_history_shards" json:"numHistoryShards,omitempty"`
		Persistence      *struct {
			AdvancedVisibilityStore *struct {
				Cassandra *struct {
					ConnectTimeout *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					Consistency    *struct {
						Consistency       *int64 `tfsdk:"consistency" json:"consistency,omitempty"`
						SerialConsistency *int64 `tfsdk:"serial_consistency" json:"serialConsistency,omitempty"`
					} `tfsdk:"consistency" json:"consistency,omitempty"`
					Datacenter               *string   `tfsdk:"datacenter" json:"datacenter,omitempty"`
					DisableInitialHostLookup *bool     `tfsdk:"disable_initial_host_lookup" json:"disableInitialHostLookup,omitempty"`
					Hosts                    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					Keyspace                 *string   `tfsdk:"keyspace" json:"keyspace,omitempty"`
					MaxConns                 *int64    `tfsdk:"max_conns" json:"maxConns,omitempty"`
					Port                     *int64    `tfsdk:"port" json:"port,omitempty"`
					User                     *string   `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"cassandra" json:"cassandra,omitempty"`
				Elasticsearch *struct {
					CloseIdleConnectionsInterval *string `tfsdk:"close_idle_connections_interval" json:"closeIdleConnectionsInterval,omitempty"`
					EnableHealthcheck            *bool   `tfsdk:"enable_healthcheck" json:"enableHealthcheck,omitempty"`
					EnableSniff                  *bool   `tfsdk:"enable_sniff" json:"enableSniff,omitempty"`
					Indices                      *struct {
						SecondaryVisibility *string `tfsdk:"secondary_visibility" json:"secondaryVisibility,omitempty"`
						Visibility          *string `tfsdk:"visibility" json:"visibility,omitempty"`
					} `tfsdk:"indices" json:"indices,omitempty"`
					LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Url      *string `tfsdk:"url" json:"url,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
					Version  *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
				SkipCreate *bool `tfsdk:"skip_create" json:"skipCreate,omitempty"`
				Sql        *struct {
					ConnectAddr        *string            `tfsdk:"connect_addr" json:"connectAddr,omitempty"`
					ConnectAttributes  *map[string]string `tfsdk:"connect_attributes" json:"connectAttributes,omitempty"`
					ConnectProtocol    *string            `tfsdk:"connect_protocol" json:"connectProtocol,omitempty"`
					DatabaseName       *string            `tfsdk:"database_name" json:"databaseName,omitempty"`
					GcpServiceAccount  *string            `tfsdk:"gcp_service_account" json:"gcpServiceAccount,omitempty"`
					MaxConnLifetime    *string            `tfsdk:"max_conn_lifetime" json:"maxConnLifetime,omitempty"`
					MaxConns           *int64             `tfsdk:"max_conns" json:"maxConns,omitempty"`
					MaxIdleConns       *int64             `tfsdk:"max_idle_conns" json:"maxIdleConns,omitempty"`
					PluginName         *string            `tfsdk:"plugin_name" json:"pluginName,omitempty"`
					TaskScanPartitions *int64             `tfsdk:"task_scan_partitions" json:"taskScanPartitions,omitempty"`
					User               *string            `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"sql" json:"sql,omitempty"`
				Tls *struct {
					CaFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_file_ref" json:"caFileRef,omitempty"`
					CertFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"cert_file_ref" json:"certFileRef,omitempty"`
					EnableHostVerification *bool `tfsdk:"enable_host_verification" json:"enableHostVerification,omitempty"`
					Enabled                *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					KeyFileRef             *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"key_file_ref" json:"keyFileRef,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"advanced_visibility_store" json:"advancedVisibilityStore,omitempty"`
			DefaultStore *struct {
				Cassandra *struct {
					ConnectTimeout *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					Consistency    *struct {
						Consistency       *int64 `tfsdk:"consistency" json:"consistency,omitempty"`
						SerialConsistency *int64 `tfsdk:"serial_consistency" json:"serialConsistency,omitempty"`
					} `tfsdk:"consistency" json:"consistency,omitempty"`
					Datacenter               *string   `tfsdk:"datacenter" json:"datacenter,omitempty"`
					DisableInitialHostLookup *bool     `tfsdk:"disable_initial_host_lookup" json:"disableInitialHostLookup,omitempty"`
					Hosts                    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					Keyspace                 *string   `tfsdk:"keyspace" json:"keyspace,omitempty"`
					MaxConns                 *int64    `tfsdk:"max_conns" json:"maxConns,omitempty"`
					Port                     *int64    `tfsdk:"port" json:"port,omitempty"`
					User                     *string   `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"cassandra" json:"cassandra,omitempty"`
				Elasticsearch *struct {
					CloseIdleConnectionsInterval *string `tfsdk:"close_idle_connections_interval" json:"closeIdleConnectionsInterval,omitempty"`
					EnableHealthcheck            *bool   `tfsdk:"enable_healthcheck" json:"enableHealthcheck,omitempty"`
					EnableSniff                  *bool   `tfsdk:"enable_sniff" json:"enableSniff,omitempty"`
					Indices                      *struct {
						SecondaryVisibility *string `tfsdk:"secondary_visibility" json:"secondaryVisibility,omitempty"`
						Visibility          *string `tfsdk:"visibility" json:"visibility,omitempty"`
					} `tfsdk:"indices" json:"indices,omitempty"`
					LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Url      *string `tfsdk:"url" json:"url,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
					Version  *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
				SkipCreate *bool `tfsdk:"skip_create" json:"skipCreate,omitempty"`
				Sql        *struct {
					ConnectAddr        *string            `tfsdk:"connect_addr" json:"connectAddr,omitempty"`
					ConnectAttributes  *map[string]string `tfsdk:"connect_attributes" json:"connectAttributes,omitempty"`
					ConnectProtocol    *string            `tfsdk:"connect_protocol" json:"connectProtocol,omitempty"`
					DatabaseName       *string            `tfsdk:"database_name" json:"databaseName,omitempty"`
					GcpServiceAccount  *string            `tfsdk:"gcp_service_account" json:"gcpServiceAccount,omitempty"`
					MaxConnLifetime    *string            `tfsdk:"max_conn_lifetime" json:"maxConnLifetime,omitempty"`
					MaxConns           *int64             `tfsdk:"max_conns" json:"maxConns,omitempty"`
					MaxIdleConns       *int64             `tfsdk:"max_idle_conns" json:"maxIdleConns,omitempty"`
					PluginName         *string            `tfsdk:"plugin_name" json:"pluginName,omitempty"`
					TaskScanPartitions *int64             `tfsdk:"task_scan_partitions" json:"taskScanPartitions,omitempty"`
					User               *string            `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"sql" json:"sql,omitempty"`
				Tls *struct {
					CaFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_file_ref" json:"caFileRef,omitempty"`
					CertFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"cert_file_ref" json:"certFileRef,omitempty"`
					EnableHostVerification *bool `tfsdk:"enable_host_verification" json:"enableHostVerification,omitempty"`
					Enabled                *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					KeyFileRef             *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"key_file_ref" json:"keyFileRef,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"default_store" json:"defaultStore,omitempty"`
			SecondaryVisibilityStore *struct {
				Cassandra *struct {
					ConnectTimeout *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					Consistency    *struct {
						Consistency       *int64 `tfsdk:"consistency" json:"consistency,omitempty"`
						SerialConsistency *int64 `tfsdk:"serial_consistency" json:"serialConsistency,omitempty"`
					} `tfsdk:"consistency" json:"consistency,omitempty"`
					Datacenter               *string   `tfsdk:"datacenter" json:"datacenter,omitempty"`
					DisableInitialHostLookup *bool     `tfsdk:"disable_initial_host_lookup" json:"disableInitialHostLookup,omitempty"`
					Hosts                    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					Keyspace                 *string   `tfsdk:"keyspace" json:"keyspace,omitempty"`
					MaxConns                 *int64    `tfsdk:"max_conns" json:"maxConns,omitempty"`
					Port                     *int64    `tfsdk:"port" json:"port,omitempty"`
					User                     *string   `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"cassandra" json:"cassandra,omitempty"`
				Elasticsearch *struct {
					CloseIdleConnectionsInterval *string `tfsdk:"close_idle_connections_interval" json:"closeIdleConnectionsInterval,omitempty"`
					EnableHealthcheck            *bool   `tfsdk:"enable_healthcheck" json:"enableHealthcheck,omitempty"`
					EnableSniff                  *bool   `tfsdk:"enable_sniff" json:"enableSniff,omitempty"`
					Indices                      *struct {
						SecondaryVisibility *string `tfsdk:"secondary_visibility" json:"secondaryVisibility,omitempty"`
						Visibility          *string `tfsdk:"visibility" json:"visibility,omitempty"`
					} `tfsdk:"indices" json:"indices,omitempty"`
					LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Url      *string `tfsdk:"url" json:"url,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
					Version  *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
				SkipCreate *bool `tfsdk:"skip_create" json:"skipCreate,omitempty"`
				Sql        *struct {
					ConnectAddr        *string            `tfsdk:"connect_addr" json:"connectAddr,omitempty"`
					ConnectAttributes  *map[string]string `tfsdk:"connect_attributes" json:"connectAttributes,omitempty"`
					ConnectProtocol    *string            `tfsdk:"connect_protocol" json:"connectProtocol,omitempty"`
					DatabaseName       *string            `tfsdk:"database_name" json:"databaseName,omitempty"`
					GcpServiceAccount  *string            `tfsdk:"gcp_service_account" json:"gcpServiceAccount,omitempty"`
					MaxConnLifetime    *string            `tfsdk:"max_conn_lifetime" json:"maxConnLifetime,omitempty"`
					MaxConns           *int64             `tfsdk:"max_conns" json:"maxConns,omitempty"`
					MaxIdleConns       *int64             `tfsdk:"max_idle_conns" json:"maxIdleConns,omitempty"`
					PluginName         *string            `tfsdk:"plugin_name" json:"pluginName,omitempty"`
					TaskScanPartitions *int64             `tfsdk:"task_scan_partitions" json:"taskScanPartitions,omitempty"`
					User               *string            `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"sql" json:"sql,omitempty"`
				Tls *struct {
					CaFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_file_ref" json:"caFileRef,omitempty"`
					CertFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"cert_file_ref" json:"certFileRef,omitempty"`
					EnableHostVerification *bool `tfsdk:"enable_host_verification" json:"enableHostVerification,omitempty"`
					Enabled                *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					KeyFileRef             *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"key_file_ref" json:"keyFileRef,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"secondary_visibility_store" json:"secondaryVisibilityStore,omitempty"`
			VisibilityStore *struct {
				Cassandra *struct {
					ConnectTimeout *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					Consistency    *struct {
						Consistency       *int64 `tfsdk:"consistency" json:"consistency,omitempty"`
						SerialConsistency *int64 `tfsdk:"serial_consistency" json:"serialConsistency,omitempty"`
					} `tfsdk:"consistency" json:"consistency,omitempty"`
					Datacenter               *string   `tfsdk:"datacenter" json:"datacenter,omitempty"`
					DisableInitialHostLookup *bool     `tfsdk:"disable_initial_host_lookup" json:"disableInitialHostLookup,omitempty"`
					Hosts                    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					Keyspace                 *string   `tfsdk:"keyspace" json:"keyspace,omitempty"`
					MaxConns                 *int64    `tfsdk:"max_conns" json:"maxConns,omitempty"`
					Port                     *int64    `tfsdk:"port" json:"port,omitempty"`
					User                     *string   `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"cassandra" json:"cassandra,omitempty"`
				Elasticsearch *struct {
					CloseIdleConnectionsInterval *string `tfsdk:"close_idle_connections_interval" json:"closeIdleConnectionsInterval,omitempty"`
					EnableHealthcheck            *bool   `tfsdk:"enable_healthcheck" json:"enableHealthcheck,omitempty"`
					EnableSniff                  *bool   `tfsdk:"enable_sniff" json:"enableSniff,omitempty"`
					Indices                      *struct {
						SecondaryVisibility *string `tfsdk:"secondary_visibility" json:"secondaryVisibility,omitempty"`
						Visibility          *string `tfsdk:"visibility" json:"visibility,omitempty"`
					} `tfsdk:"indices" json:"indices,omitempty"`
					LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Url      *string `tfsdk:"url" json:"url,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
					Version  *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
				SkipCreate *bool `tfsdk:"skip_create" json:"skipCreate,omitempty"`
				Sql        *struct {
					ConnectAddr        *string            `tfsdk:"connect_addr" json:"connectAddr,omitempty"`
					ConnectAttributes  *map[string]string `tfsdk:"connect_attributes" json:"connectAttributes,omitempty"`
					ConnectProtocol    *string            `tfsdk:"connect_protocol" json:"connectProtocol,omitempty"`
					DatabaseName       *string            `tfsdk:"database_name" json:"databaseName,omitempty"`
					GcpServiceAccount  *string            `tfsdk:"gcp_service_account" json:"gcpServiceAccount,omitempty"`
					MaxConnLifetime    *string            `tfsdk:"max_conn_lifetime" json:"maxConnLifetime,omitempty"`
					MaxConns           *int64             `tfsdk:"max_conns" json:"maxConns,omitempty"`
					MaxIdleConns       *int64             `tfsdk:"max_idle_conns" json:"maxIdleConns,omitempty"`
					PluginName         *string            `tfsdk:"plugin_name" json:"pluginName,omitempty"`
					TaskScanPartitions *int64             `tfsdk:"task_scan_partitions" json:"taskScanPartitions,omitempty"`
					User               *string            `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"sql" json:"sql,omitempty"`
				Tls *struct {
					CaFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_file_ref" json:"caFileRef,omitempty"`
					CertFileRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"cert_file_ref" json:"certFileRef,omitempty"`
					EnableHostVerification *bool `tfsdk:"enable_host_verification" json:"enableHostVerification,omitempty"`
					Enabled                *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					KeyFileRef             *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"key_file_ref" json:"keyFileRef,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"visibility_store" json:"visibilityStore,omitempty"`
		} `tfsdk:"persistence" json:"persistence,omitempty"`
		Services *struct {
			Frontend *struct {
				HttpPort       *int64               `tfsdk:"http_port" json:"httpPort,omitempty"`
				InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
				MembershipPort *int64               `tfsdk:"membership_port" json:"membershipPort,omitempty"`
				Overrides      *struct {
					Deployment *struct {
						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec *struct {
							Template *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
							} `tfsdk:"template" json:"template,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"deployment" json:"deployment,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Port      *int64 `tfsdk:"port" json:"port,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"frontend" json:"frontend,omitempty"`
			History *struct {
				HttpPort       *int64               `tfsdk:"http_port" json:"httpPort,omitempty"`
				InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
				MembershipPort *int64               `tfsdk:"membership_port" json:"membershipPort,omitempty"`
				Overrides      *struct {
					Deployment *struct {
						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec *struct {
							Template *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
							} `tfsdk:"template" json:"template,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"deployment" json:"deployment,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Port      *int64 `tfsdk:"port" json:"port,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"history" json:"history,omitempty"`
			InternalFrontend *struct {
				Enabled        *bool                `tfsdk:"enabled" json:"enabled,omitempty"`
				HttpPort       *int64               `tfsdk:"http_port" json:"httpPort,omitempty"`
				InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
				MembershipPort *int64               `tfsdk:"membership_port" json:"membershipPort,omitempty"`
				Overrides      *struct {
					Deployment *struct {
						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec *struct {
							Template *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
							} `tfsdk:"template" json:"template,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"deployment" json:"deployment,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Port      *int64 `tfsdk:"port" json:"port,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"internal_frontend" json:"internalFrontend,omitempty"`
			Matching *struct {
				HttpPort       *int64               `tfsdk:"http_port" json:"httpPort,omitempty"`
				InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
				MembershipPort *int64               `tfsdk:"membership_port" json:"membershipPort,omitempty"`
				Overrides      *struct {
					Deployment *struct {
						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec *struct {
							Template *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
							} `tfsdk:"template" json:"template,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"deployment" json:"deployment,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Port      *int64 `tfsdk:"port" json:"port,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"matching" json:"matching,omitempty"`
			Overrides *struct {
				Deployment *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec *struct {
						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
								Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"template" json:"template,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
			} `tfsdk:"overrides" json:"overrides,omitempty"`
			Worker *struct {
				HttpPort       *int64               `tfsdk:"http_port" json:"httpPort,omitempty"`
				InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
				MembershipPort *int64               `tfsdk:"membership_port" json:"membershipPort,omitempty"`
				Overrides      *struct {
					Deployment *struct {
						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec *struct {
							Template *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
							} `tfsdk:"template" json:"template,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"deployment" json:"deployment,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Port      *int64 `tfsdk:"port" json:"port,omitempty"`
				Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"worker" json:"worker,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		Ui *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Image   *string `tfsdk:"image" json:"image,omitempty"`
			Ingress *struct {
				Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Hosts            *[]string          `tfsdk:"hosts" json:"hosts,omitempty"`
				IngressClassName *string            `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
				Tls              *[]struct {
					Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			Overrides *struct {
				Deployment *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec *struct {
						Template *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
								Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"template" json:"template,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
			} `tfsdk:"overrides" json:"overrides,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Service *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"ui" json:"ui,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TemporalIoTemporalClusterV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_temporal_io_temporal_cluster_v1beta1_manifest"
}

func (r *TemporalIoTemporalClusterV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TemporalCluster defines a temporal cluster deployment.",
		MarkdownDescription: "TemporalCluster defines a temporal cluster deployment.",
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
				Description:         "Specification of the desired behavior of the Temporal cluster.",
				MarkdownDescription: "Specification of the desired behavior of the Temporal cluster.",
				Attributes: map[string]schema.Attribute{
					"admintools": schema.SingleNestedAttribute{
						Description:         "AdminTools allows configuration of the optional admin tool pod deployed alongside the cluster.",
						MarkdownDescription: "AdminTools allows configuration of the optional admin tool pod deployed alongside the cluster.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines if the operator should deploy the admin tools alongside the cluster.",
								MarkdownDescription: "Enabled defines if the operator should deploy the admin tools alongside the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Image defines the temporal admin tools docker image the instance should run.",
								MarkdownDescription: "Image defines the temporal admin tools docker image the instance should run.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"overrides": schema.SingleNestedAttribute{
								Description:         "Overrides adds some overrides to the resources deployed for the ui.",
								MarkdownDescription: "Overrides adds some overrides to the resources deployed for the ui.",
								Attributes: map[string]schema.Attribute{
									"deployment": schema.SingleNestedAttribute{
										Description:         "Override configuration for the temporal service Deployment.",
										MarkdownDescription: "Override configuration for the temporal service Deployment.",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.SingleNestedAttribute{
												Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
												MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
														MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
														MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
												Description:         "Specification of the desired behavior of the Deployment.",
												MarkdownDescription: "Specification of the desired behavior of the Deployment.",
												Attributes: map[string]schema.Attribute{
													"template": schema.SingleNestedAttribute{
														Description:         "Template describes the pods that will be created.",
														MarkdownDescription: "Template describes the pods that will be created.",
														Attributes: map[string]schema.Attribute{
															"metadata": schema.SingleNestedAttribute{
																Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																Attributes: map[string]schema.Attribute{
																	"annotations": schema.MapAttribute{
																		Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																		MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"labels": schema.MapAttribute{
																		Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																		MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

															"spec": schema.MapAttribute{
																Description:         "Specification of the desired behavior of the pod.",
																MarkdownDescription: "Specification of the desired behavior of the pod.",
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

							"resources": schema.SingleNestedAttribute{
								Description:         "Compute Resources required by the ui. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Compute Resources required by the ui. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"archival": schema.SingleNestedAttribute{
						Description:         "Archival allows Workflow Execution Event Histories and Visibility data backups for the temporal cluster.",
						MarkdownDescription: "Archival allows Workflow Execution Event Histories and Visibility data backups for the temporal cluster.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines if the archival is enabled for the cluster.",
								MarkdownDescription: "Enabled defines if the archival is enabled for the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"history": schema.SingleNestedAttribute{
								Description:         "History is the default config for the history archival.",
								MarkdownDescription: "History is the default config for the history archival.",
								Attributes: map[string]schema.Attribute{
									"enable_read": schema.BoolAttribute{
										Description:         "EnableRead allows temporal to read from the archived Event History.",
										MarkdownDescription: "EnableRead allows temporal to read from the archived Event History.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if the archival is enabled by default for all namespaces or for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										MarkdownDescription: "Enabled defines if the archival is enabled by default for all namespaces or for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is ...",
										MarkdownDescription: "Path is ...",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"paused": schema.BoolAttribute{
										Description:         "Paused defines if the archival is paused.",
										MarkdownDescription: "Paused defines if the archival is paused.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider": schema.SingleNestedAttribute{
								Description:         "Provider defines the archival provider for the cluster. The same provider is used for both history and visibility, but some config can be changed using spec.archival.[history|visibility].config.",
								MarkdownDescription: "Provider defines the archival provider for the cluster. The same provider is used for both history and visibility, but some config can be changed using spec.archival.[history|visibility].config.",
								Attributes: map[string]schema.Attribute{
									"filestore": schema.SingleNestedAttribute{
										Description:         "FilestoreArchiver is the file store archival provider configuration.",
										MarkdownDescription: "FilestoreArchiver is the file store archival provider configuration.",
										Attributes: map[string]schema.Attribute{
											"dir_permissions": schema.StringAttribute{
												Description:         "DirPermissions sets the directory permissions of the archive directory. It's recommend to leave it empty and use the default value of '0766' to avoid read/write issues.",
												MarkdownDescription: "DirPermissions sets the directory permissions of the archive directory. It's recommend to leave it empty and use the default value of '0766' to avoid read/write issues.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"file_permissions": schema.StringAttribute{
												Description:         "FilePermissions sets the file permissions of the archived files. It's recommend to leave it empty and use the default value of '0666' to avoid read/write issues.",
												MarkdownDescription: "FilePermissions sets the file permissions of the archived files. It's recommend to leave it empty and use the default value of '0666' to avoid read/write issues.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"gcs": schema.SingleNestedAttribute{
										Description:         "GCSArchiver is the GCS archival provider configuration.",
										MarkdownDescription: "GCSArchiver is the GCS archival provider configuration.",
										Attributes: map[string]schema.Attribute{
											"credentials_ref": schema.SingleNestedAttribute{
												Description:         "SecretAccessKeyRef is the secret key selector containing Google Cloud Storage credentials file.",
												MarkdownDescription: "SecretAccessKeyRef is the secret key selector containing Google Cloud Storage credentials file.",
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
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3": schema.SingleNestedAttribute{
										Description:         "S3Archiver is the S3 archival provider configuration.",
										MarkdownDescription: "S3Archiver is the S3 archival provider configuration.",
										Attributes: map[string]schema.Attribute{
											"credentials": schema.SingleNestedAttribute{
												Description:         "Use credentials if you want to use aws credentials from secret.",
												MarkdownDescription: "Use credentials if you want to use aws credentials from secret.",
												Attributes: map[string]schema.Attribute{
													"access_key_id_ref": schema.SingleNestedAttribute{
														Description:         "AccessKeyIDRef is the secret key selector containing AWS access key ID.",
														MarkdownDescription: "AccessKeyIDRef is the secret key selector containing AWS access key ID.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "SecretAccessKeyRef is the secret key selector containing AWS secret access key.",
														MarkdownDescription: "SecretAccessKeyRef is the secret key selector containing AWS secret access key.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoint": schema.StringAttribute{
												Description:         "Use Endpoint if you want to use s3-compatible object storage.",
												MarkdownDescription: "Use Endpoint if you want to use s3-compatible object storage.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"region": schema.StringAttribute{
												Description:         "Region is the aws s3 region.",
												MarkdownDescription: "Region is the aws s3 region.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"role_name": schema.StringAttribute{
												Description:         "Use RoleName if you want the temporal service account to assume an AWS Identity and Access Management (IAM) role.",
												MarkdownDescription: "Use RoleName if you want the temporal service account to assume an AWS Identity and Access Management (IAM) role.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_force_path_style": schema.BoolAttribute{
												Description:         "Use s3ForcePathStyle if you want to use s3 path style.",
												MarkdownDescription: "Use s3ForcePathStyle if you want to use s3 path style.",
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

							"visibility": schema.SingleNestedAttribute{
								Description:         "Visibility is the default config for visibility archival.",
								MarkdownDescription: "Visibility is the default config for visibility archival.",
								Attributes: map[string]schema.Attribute{
									"enable_read": schema.BoolAttribute{
										Description:         "EnableRead allows temporal to read from the archived Event History.",
										MarkdownDescription: "EnableRead allows temporal to read from the archived Event History.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if the archival is enabled by default for all namespaces or for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										MarkdownDescription: "Enabled defines if the archival is enabled by default for all namespaces or for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is ...",
										MarkdownDescription: "Path is ...",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"paused": schema.BoolAttribute{
										Description:         "Paused defines if the archival is paused.",
										MarkdownDescription: "Paused defines if the archival is paused.",
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

					"authorization": schema.SingleNestedAttribute{
						Description:         "Authorization allows authorization configuration for the temporal cluster.",
						MarkdownDescription: "Authorization allows authorization configuration for the temporal cluster.",
						Attributes: map[string]schema.Attribute{
							"authorizer": schema.StringAttribute{
								Description:         "Authorizer defines the authorization mechanism to be used. It can be left as an empty string to use a no-operation authorizer (noopAuthorizer), or set to 'default' to use the temporal's default authorizer (defaultAuthorizer).",
								MarkdownDescription: "Authorizer defines the authorization mechanism to be used. It can be left as an empty string to use a no-operation authorizer (noopAuthorizer), or set to 'default' to use the temporal's default authorizer (defaultAuthorizer).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"claim_mapper": schema.StringAttribute{
								Description:         "ClaimMapper specifies the claim mapping mechanism used for handling JWT claims. Similar to the Authorizer, it can be left as an empty string to use a no-operation claim mapper (noopClaimMapper), or set to 'default' to use the default JWT claim mapper (defaultJWTClaimMapper).",
								MarkdownDescription: "ClaimMapper specifies the claim mapping mechanism used for handling JWT claims. Similar to the Authorizer, it can be left as an empty string to use a no-operation claim mapper (noopClaimMapper), or set to 'default' to use the default JWT claim mapper (defaultJWTClaimMapper).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jwt_key_provider": schema.SingleNestedAttribute{
								Description:         "JWTKeyProvider specifies the signing key provider used for validating JWT tokens.",
								MarkdownDescription: "JWTKeyProvider specifies the signing key provider used for validating JWT tokens.",
								Attributes: map[string]schema.Attribute{
									"key_source_ur_is": schema.ListAttribute{
										Description:         "KeySourceURIs is a list of URIs where the JWT signing keys can be obtained. These URIs are used by the authorization system to fetch the public keys necessary for validating JWT tokens.",
										MarkdownDescription: "KeySourceURIs is a list of URIs where the JWT signing keys can be obtained. These URIs are used by the authorization system to fetch the public keys necessary for validating JWT tokens.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"refresh_interval": schema.StringAttribute{
										Description:         "RefreshInterval defines the time interval at which temporal should refresh the JWT signing keys from the specified URIs.",
										MarkdownDescription: "RefreshInterval defines the time interval at which temporal should refresh the JWT signing keys from the specified URIs.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"permissions_claim_name": schema.StringAttribute{
								Description:         "PermissionsClaimName is the name of the claim within the JWT token that contains the user's permissions.",
								MarkdownDescription: "PermissionsClaimName is the name of the claim within the JWT token that contains the user's permissions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"dynamic_config": schema.SingleNestedAttribute{
						Description:         "DynamicConfig allows advanced configuration for the temporal cluster.",
						MarkdownDescription: "DynamicConfig allows advanced configuration for the temporal cluster.",
						Attributes: map[string]schema.Attribute{
							"poll_interval": schema.StringAttribute{
								Description:         "PollInterval defines how often the config should be updated by checking provided values. Defaults to 10s.",
								MarkdownDescription: "PollInterval defines how often the config should be updated by checking provided values. Defaults to 10s.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"values": schema.MapAttribute{
								Description:         "Values contains all dynamic config keys and their constrained values.",
								MarkdownDescription: "Values contains all dynamic config keys and their constrained values.",
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

					"image": schema.StringAttribute{
						Description:         "Image defines the temporal server docker image the cluster should use for each services.",
						MarkdownDescription: "Image defines the temporal server docker image the cluster should use for each services.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "An optional list of references to secrets in the same namespace to use for pulling temporal images from registries.",
						MarkdownDescription: "An optional list of references to secrets in the same namespace to use for pulling temporal images from registries.",
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

					"job_init_containers": schema.ListAttribute{
						Description:         "JobInitContainers adds a list of init containers to the setup's jobs.",
						MarkdownDescription: "JobInitContainers adds a list of init containers to the setup's jobs.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"job_resources": schema.SingleNestedAttribute{
						Description:         "JobResources allows set resources for setup/update jobs.",
						MarkdownDescription: "JobResources allows set resources for setup/update jobs.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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

					"job_ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "JobTTLSecondsAfterFinished is amount of time to keep job pods after jobs are completed. Defaults to 300 seconds.",
						MarkdownDescription: "JobTTLSecondsAfterFinished is amount of time to keep job pods after jobs are completed. Defaults to 300 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"log": schema.SingleNestedAttribute{
						Description:         "Log defines temporal cluster's logger configuration.",
						MarkdownDescription: "Log defines temporal cluster's logger configuration.",
						Attributes: map[string]schema.Attribute{
							"development": schema.BoolAttribute{
								Description:         "Development determines whether the logger is run in Development (== Test) or in Production mode.  Default is Production.  Production-stage disables panics from DPanic logging.",
								MarkdownDescription: "Development determines whether the logger is run in Development (== Test) or in Production mode.  Default is Production.  Production-stage disables panics from DPanic logging.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "Format determines the format of each log file printed to the output. Use 'console' if you want stack traces to appear on multiple lines.",
								MarkdownDescription: "Format determines the format of each log file printed to the output. Use 'console' if you want stack traces to appear on multiple lines.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("json", "console"),
								},
							},

							"level": schema.StringAttribute{
								Description:         "Level is the desired log level; see colocated zap_logger.go::parseZapLevel()",
								MarkdownDescription: "Level is the desired log level; see colocated zap_logger.go::parseZapLevel()",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("debug", "info", "warn", "error", "dpanic", "panic", "fatal"),
								},
							},

							"output_file": schema.StringAttribute{
								Description:         "OutputFile is the path to the log output file.",
								MarkdownDescription: "OutputFile is the path to the log output file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stdout": schema.BoolAttribute{
								Description:         "Stdout is true if the output needs to goto standard out; default is stderr.",
								MarkdownDescription: "Stdout is true if the output needs to goto standard out; default is stderr.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"m_tls": schema.SingleNestedAttribute{
						Description:         "MTLS allows configuration of the network traffic encryption for the cluster.",
						MarkdownDescription: "MTLS allows configuration of the network traffic encryption for the cluster.",
						Attributes: map[string]schema.Attribute{
							"certificates_duration": schema.SingleNestedAttribute{
								Description:         "CertificatesDuration allows configuration of maximum certificates lifetime. Useless if mTLS provider is not cert-manager.",
								MarkdownDescription: "CertificatesDuration allows configuration of maximum certificates lifetime. Useless if mTLS provider is not cert-manager.",
								Attributes: map[string]schema.Attribute{
									"client_certificates": schema.StringAttribute{
										Description:         "ClientCertificates is the 'duration' (i.e. lifetime) of the client certificates. It defaults to 1 year.",
										MarkdownDescription: "ClientCertificates is the 'duration' (i.e. lifetime) of the client certificates. It defaults to 1 year.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"frontend_certificate": schema.StringAttribute{
										Description:         "FrontendCertificate is the 'duration' (i.e. lifetime) of the frontend certificate. It defaults to 1 year.",
										MarkdownDescription: "FrontendCertificate is the 'duration' (i.e. lifetime) of the frontend certificate. It defaults to 1 year.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"intermediate_c_as_certificates": schema.StringAttribute{
										Description:         "IntermediateCACertificates is the 'duration' (i.e. lifetime) of the intermediate CAs Certificates. It defaults to 5 years.",
										MarkdownDescription: "IntermediateCACertificates is the 'duration' (i.e. lifetime) of the intermediate CAs Certificates. It defaults to 5 years.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"internode_certificate": schema.StringAttribute{
										Description:         "InternodeCertificate is the 'duration' (i.e. lifetime) of the internode certificate. It defaults to 1 year.",
										MarkdownDescription: "InternodeCertificate is the 'duration' (i.e. lifetime) of the internode certificate. It defaults to 1 year.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"root_ca_certificate": schema.StringAttribute{
										Description:         "RootCACertificate is the 'duration' (i.e. lifetime) of the Root CA Certificate. It defaults to 10 years.",
										MarkdownDescription: "RootCACertificate is the 'duration' (i.e. lifetime) of the Root CA Certificate. It defaults to 10 years.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"frontend": schema.SingleNestedAttribute{
								Description:         "Frontend allows configuration of the frontend's public endpoint traffic encryption. Useless if mTLS provider is not cert-manager.",
								MarkdownDescription: "Frontend allows configuration of the frontend's public endpoint traffic encryption. Useless if mTLS provider is not cert-manager.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if the operator should enable mTLS for cluster's public endpoints.",
										MarkdownDescription: "Enabled defines if the operator should enable mTLS for cluster's public endpoints.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"internode": schema.SingleNestedAttribute{
								Description:         "Internode allows configuration of the internode traffic encryption. Useless if mTLS provider is not cert-manager.",
								MarkdownDescription: "Internode allows configuration of the internode traffic encryption. Useless if mTLS provider is not cert-manager.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if the operator should enable mTLS for network between cluster nodes.",
										MarkdownDescription: "Enabled defines if the operator should enable mTLS for network between cluster nodes.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider": schema.StringAttribute{
								Description:         "Provider defines the tool used to manage mTLS certificates.",
								MarkdownDescription: "Provider defines the tool used to manage mTLS certificates.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cert-manager", "linkerd", "istio"),
								},
							},

							"refresh_interval": schema.StringAttribute{
								Description:         "RefreshInterval defines interval between refreshes of certificates in the cluster components. Defaults to 1 hour. Useless if mTLS provider is not cert-manager.",
								MarkdownDescription: "RefreshInterval defines interval between refreshes of certificates in the cluster components. Defaults to 1 hour. Useless if mTLS provider is not cert-manager.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"renew_before": schema.StringAttribute{
								Description:         "RenewBefore is defines how long before the currently issued certificate's expiry cert-manager should renew the certificate. The default is 2/3 of the issued certificate's duration. Minimum accepted value is 5 minutes. Useless if mTLS provider is not cert-manager.",
								MarkdownDescription: "RenewBefore is defines how long before the currently issued certificate's expiry cert-manager should renew the certificate. The default is 2/3 of the issued certificate's duration. Minimum accepted value is 5 minutes. Useless if mTLS provider is not cert-manager.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics allows configuration of scraping endpoints for stats. prometheus or m3.",
						MarkdownDescription: "Metrics allows configuration of scraping endpoints for stats. prometheus or m3.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines if the operator should enable metrics exposition on temporal components.",
								MarkdownDescription: "Enabled defines if the operator should enable metrics exposition on temporal components.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"exclude_tags": schema.MapAttribute{
								Description:         "ExcludeTags is a map from tag name string to tag values string list. Each value present in keys will have relevant tag value replaced with '_tag_excluded_' Each value in values list will white-list tag values to be reported as usual.",
								MarkdownDescription: "ExcludeTags is a map from tag name string to tag values string list. Each value present in keys will have relevant tag value replaced with '_tag_excluded_' Each value in values list will white-list tag values to be reported as usual.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"per_unit_histogram_boundaries": schema.MapAttribute{
								Description:         "PerUnitHistogramBoundaries defines the default histogram bucket boundaries. Configuration of histogram boundaries for given metric unit.  Supported values: - 'dimensionless' - 'milliseconds' - 'bytes'",
								MarkdownDescription: "PerUnitHistogramBoundaries defines the default histogram bucket boundaries. Configuration of histogram boundaries for given metric unit.  Supported values: - 'dimensionless' - 'milliseconds' - 'bytes'",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prefix": schema.StringAttribute{
								Description:         "Prefix sets the prefix to all outgoing metrics",
								MarkdownDescription: "Prefix sets the prefix to all outgoing metrics",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prometheus": schema.SingleNestedAttribute{
								Description:         "Prometheus reporter configuration.",
								MarkdownDescription: "Prometheus reporter configuration.",
								Attributes: map[string]schema.Attribute{
									"listen_address": schema.StringAttribute{
										Description:         "Deprecated. Address for prometheus to serve metrics from.",
										MarkdownDescription: "Deprecated. Address for prometheus to serve metrics from.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"listen_port": schema.Int64Attribute{
										Description:         "ListenPort for prometheus to serve metrics from.",
										MarkdownDescription: "ListenPort for prometheus to serve metrics from.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scrape_config": schema.SingleNestedAttribute{
										Description:         "ScrapeConfig is the prometheus scrape configuration.",
										MarkdownDescription: "ScrapeConfig is the prometheus scrape configuration.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.BoolAttribute{
												Description:         "Annotations defines if the operator should add prometheus scrape annotations to the services pods.",
												MarkdownDescription: "Annotations defines if the operator should add prometheus scrape annotations to the services pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_monitor": schema.SingleNestedAttribute{
												Description:         "PrometheusScrapeConfigServiceMonitor is the configuration for prometheus operator ServiceMonitor.",
												MarkdownDescription: "PrometheusScrapeConfigServiceMonitor is the configuration for prometheus operator ServiceMonitor.",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "Enabled defines if the operator should create a ServiceMonitor for each services.",
														MarkdownDescription: "Enabled defines if the operator should create a ServiceMonitor for each services.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels adds extra labels to the ServiceMonitor.",
														MarkdownDescription: "Labels adds extra labels to the ServiceMonitor.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metric_relabelings": schema.ListNestedAttribute{
														Description:         "MetricRelabelConfigs to apply to samples before ingestion.",
														MarkdownDescription: "MetricRelabelConfigs to apply to samples before ingestion.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
																	MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
																	},
																},

																"modulus": schema.Int64Attribute{
																	Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
																	MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"regex": schema.StringAttribute{
																	Description:         "Regular expression against which the extracted value is matched.",
																	MarkdownDescription: "Regular expression against which the extracted value is matched.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"replacement": schema.StringAttribute{
																	Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
																	MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"separator": schema.StringAttribute{
																	Description:         "Separator is the string between concatenated SourceLabels.",
																	MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"source_labels": schema.ListAttribute{
																	Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
																	MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"target_label": schema.StringAttribute{
																	Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
																	MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

													"override": schema.SingleNestedAttribute{
														Description:         "Override allows customization of the created ServiceMonitor. All fields can be overwritten except 'endpoints', 'selector' and 'namespaceSelector'.",
														MarkdownDescription: "Override allows customization of the created ServiceMonitor. All fields can be overwritten except 'endpoints', 'selector' and 'namespaceSelector'.",
														Attributes: map[string]schema.Attribute{
															"attach_metadata": schema.SingleNestedAttribute{
																Description:         "'attachMetadata' defines additional metadata which is added to the discovered targets.  It requires Prometheus >= v2.37.0.",
																MarkdownDescription: "'attachMetadata' defines additional metadata which is added to the discovered targets.  It requires Prometheus >= v2.37.0.",
																Attributes: map[string]schema.Attribute{
																	"node": schema.BoolAttribute{
																		Description:         "When set to true, Prometheus must have the 'get' permission on the 'Nodes' objects.",
																		MarkdownDescription: "When set to true, Prometheus must have the 'get' permission on the 'Nodes' objects.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": schema.ListNestedAttribute{
																Description:         "List of endpoints part of this ServiceMonitor.",
																MarkdownDescription: "List of endpoints part of this ServiceMonitor.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"authorization": schema.SingleNestedAttribute{
																			Description:         "'authorization' configures the Authorization header credentials to use when scraping the target.  Cannot be set at the same time as 'basicAuth', or 'oauth2'.",
																			MarkdownDescription: "'authorization' configures the Authorization header credentials to use when scraping the target.  Cannot be set at the same time as 'basicAuth', or 'oauth2'.",
																			Attributes: map[string]schema.Attribute{
																				"credentials": schema.SingleNestedAttribute{
																					Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																					MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
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

																				"type": schema.StringAttribute{
																					Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																					MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"basic_auth": schema.SingleNestedAttribute{
																			Description:         "'basicAuth' configures the Basic Authentication credentials to use when scraping the target.  Cannot be set at the same time as 'authorization', or 'oauth2'.",
																			MarkdownDescription: "'basicAuth' configures the Basic Authentication credentials to use when scraping the target.  Cannot be set at the same time as 'authorization', or 'oauth2'.",
																			Attributes: map[string]schema.Attribute{
																				"password": schema.SingleNestedAttribute{
																					Description:         "'password' specifies a key of a Secret containing the password for authentication.",
																					MarkdownDescription: "'password' specifies a key of a Secret containing the password for authentication.",
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

																				"username": schema.SingleNestedAttribute{
																					Description:         "'username' specifies a key of a Secret containing the username for authentication.",
																					MarkdownDescription: "'username' specifies a key of a Secret containing the username for authentication.",
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

																		"bearer_token_file": schema.StringAttribute{
																			Description:         "File to read bearer token for scraping the target.  Deprecated: use 'authorization' instead.",
																			MarkdownDescription: "File to read bearer token for scraping the target.  Deprecated: use 'authorization' instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"bearer_token_secret": schema.SingleNestedAttribute{
																			Description:         "'bearerTokenSecret' specifies a key of a Secret containing the bearer token for scraping targets. The secret needs to be in the same namespace as the ServiceMonitor object and readable by the Prometheus Operator.  Deprecated: use 'authorization' instead.",
																			MarkdownDescription: "'bearerTokenSecret' specifies a key of a Secret containing the bearer token for scraping targets. The secret needs to be in the same namespace as the ServiceMonitor object and readable by the Prometheus Operator.  Deprecated: use 'authorization' instead.",
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

																		"enable_http2": schema.BoolAttribute{
																			Description:         "'enableHttp2' can be used to disable HTTP2 when scraping the target.",
																			MarkdownDescription: "'enableHttp2' can be used to disable HTTP2 when scraping the target.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"filter_running": schema.BoolAttribute{
																			Description:         "When true, the pods which are not running (e.g. either in Failed or Succeeded state) are dropped during the target discovery.  If unset, the filtering is enabled.  More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase",
																			MarkdownDescription: "When true, the pods which are not running (e.g. either in Failed or Succeeded state) are dropped during the target discovery.  If unset, the filtering is enabled.  More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"follow_redirects": schema.BoolAttribute{
																			Description:         "'followRedirects' defines whether the scrape requests should follow HTTP 3xx redirects.",
																			MarkdownDescription: "'followRedirects' defines whether the scrape requests should follow HTTP 3xx redirects.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"honor_labels": schema.BoolAttribute{
																			Description:         "When true, 'honorLabels' preserves the metric's labels when they collide with the target's labels.",
																			MarkdownDescription: "When true, 'honorLabels' preserves the metric's labels when they collide with the target's labels.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"honor_timestamps": schema.BoolAttribute{
																			Description:         "'honorTimestamps' controls whether Prometheus preserves the timestamps when exposed by the target.",
																			MarkdownDescription: "'honorTimestamps' controls whether Prometheus preserves the timestamps when exposed by the target.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"interval": schema.StringAttribute{
																			Description:         "Interval at which Prometheus scrapes the metrics from the target.  If empty, Prometheus uses the global scrape interval.",
																			MarkdownDescription: "Interval at which Prometheus scrapes the metrics from the target.  If empty, Prometheus uses the global scrape interval.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
																			},
																		},

																		"metric_relabelings": schema.ListNestedAttribute{
																			Description:         "'metricRelabelings' configures the relabeling rules to apply to the samples before ingestion.",
																			MarkdownDescription: "'metricRelabelings' configures the relabeling rules to apply to the samples before ingestion.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"action": schema.StringAttribute{
																						Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
																						MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
																						},
																					},

																					"modulus": schema.Int64Attribute{
																						Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
																						MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"regex": schema.StringAttribute{
																						Description:         "Regular expression against which the extracted value is matched.",
																						MarkdownDescription: "Regular expression against which the extracted value is matched.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"replacement": schema.StringAttribute{
																						Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
																						MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"separator": schema.StringAttribute{
																						Description:         "Separator is the string between concatenated SourceLabels.",
																						MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"source_labels": schema.ListAttribute{
																						Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
																						MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"target_label": schema.StringAttribute{
																						Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
																						MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

																		"oauth2": schema.SingleNestedAttribute{
																			Description:         "'oauth2' configures the OAuth2 settings to use when scraping the target.  It requires Prometheus >= 2.27.0.  Cannot be set at the same time as 'authorization', or 'basicAuth'.",
																			MarkdownDescription: "'oauth2' configures the OAuth2 settings to use when scraping the target.  It requires Prometheus >= 2.27.0.  Cannot be set at the same time as 'authorization', or 'basicAuth'.",
																			Attributes: map[string]schema.Attribute{
																				"client_id": schema.SingleNestedAttribute{
																					Description:         "'clientId' specifies a key of a Secret or ConfigMap containing the OAuth2 client's ID.",
																					MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing the OAuth2 client's ID.",
																					Attributes: map[string]schema.Attribute{
																						"config_map": schema.SingleNestedAttribute{
																							Description:         "ConfigMap containing data to use for the targets.",
																							MarkdownDescription: "ConfigMap containing data to use for the targets.",
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

																						"secret": schema.SingleNestedAttribute{
																							Description:         "Secret containing data to use for the targets.",
																							MarkdownDescription: "Secret containing data to use for the targets.",
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
																					Required: true,
																					Optional: false,
																					Computed: false,
																				},

																				"client_secret": schema.SingleNestedAttribute{
																					Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2 client's secret.",
																					MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2 client's secret.",
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
																					Required: true,
																					Optional: false,
																					Computed: false,
																				},

																				"endpoint_params": schema.MapAttribute{
																					Description:         "'endpointParams' configures the HTTP parameters to append to the token URL.",
																					MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the token URL.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"scopes": schema.ListAttribute{
																					Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																					MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"token_url": schema.StringAttribute{
																					Description:         "'tokenURL' configures the URL to fetch the token from.",
																					MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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

																		"params": schema.MapAttribute{
																			Description:         "params define optional HTTP URL parameters.",
																			MarkdownDescription: "params define optional HTTP URL parameters.",
																			ElementType:         types.ListType{ElemType: types.StringType},
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "HTTP path from which to scrape for metrics.  If empty, Prometheus uses the default value (e.g. '/metrics').",
																			MarkdownDescription: "HTTP path from which to scrape for metrics.  If empty, Prometheus uses the default value (e.g. '/metrics').",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name of the Service port which this endpoint refers to.  It takes precedence over 'targetPort'.",
																			MarkdownDescription: "Name of the Service port which this endpoint refers to.  It takes precedence over 'targetPort'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"proxy_url": schema.StringAttribute{
																			Description:         "'proxyURL' configures the HTTP Proxy URL (e.g. 'http://proxyserver:2195') to go through when scraping the target.",
																			MarkdownDescription: "'proxyURL' configures the HTTP Proxy URL (e.g. 'http://proxyserver:2195') to go through when scraping the target.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"relabelings": schema.ListNestedAttribute{
																			Description:         "'relabelings' configures the relabeling rules to apply the target's metadata labels.  The Operator automatically adds relabelings for a few standard Kubernetes fields.  The original scrape job's name is available via the '__tmp_prometheus_job_name' label.  More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
																			MarkdownDescription: "'relabelings' configures the relabeling rules to apply the target's metadata labels.  The Operator automatically adds relabelings for a few standard Kubernetes fields.  The original scrape job's name is available via the '__tmp_prometheus_job_name' label.  More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"action": schema.StringAttribute{
																						Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
																						MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
																						},
																					},

																					"modulus": schema.Int64Attribute{
																						Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
																						MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"regex": schema.StringAttribute{
																						Description:         "Regular expression against which the extracted value is matched.",
																						MarkdownDescription: "Regular expression against which the extracted value is matched.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"replacement": schema.StringAttribute{
																						Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
																						MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"separator": schema.StringAttribute{
																						Description:         "Separator is the string between concatenated SourceLabels.",
																						MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"source_labels": schema.ListAttribute{
																						Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
																						MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"target_label": schema.StringAttribute{
																						Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
																						MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

																		"scheme": schema.StringAttribute{
																			Description:         "HTTP scheme to use for scraping.  'http' and 'https' are the expected values unless you rewrite the '__scheme__' label via relabeling.  If empty, Prometheus uses the default value 'http'.",
																			MarkdownDescription: "HTTP scheme to use for scraping.  'http' and 'https' are the expected values unless you rewrite the '__scheme__' label via relabeling.  If empty, Prometheus uses the default value 'http'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("http", "https"),
																			},
																		},

																		"scrape_timeout": schema.StringAttribute{
																			Description:         "Timeout after which Prometheus considers the scrape to be failed.  If empty, Prometheus uses the global scrape timeout unless it is less than the target's scrape interval value in which the latter is used.",
																			MarkdownDescription: "Timeout after which Prometheus considers the scrape to be failed.  If empty, Prometheus uses the global scrape timeout unless it is less than the target's scrape interval value in which the latter is used.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
																			},
																		},

																		"target_port": schema.StringAttribute{
																			Description:         "Name or number of the target port of the 'Pod' object behind the Service, the port must be specified with container port property.  Deprecated: use 'port' instead.",
																			MarkdownDescription: "Name or number of the target port of the 'Pod' object behind the Service, the port must be specified with container port property.  Deprecated: use 'port' instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"tls_config": schema.SingleNestedAttribute{
																			Description:         "TLS configuration to use when scraping the target.",
																			MarkdownDescription: "TLS configuration to use when scraping the target.",
																			Attributes: map[string]schema.Attribute{
																				"ca": schema.SingleNestedAttribute{
																					Description:         "Certificate authority used when verifying server certificates.",
																					MarkdownDescription: "Certificate authority used when verifying server certificates.",
																					Attributes: map[string]schema.Attribute{
																						"config_map": schema.SingleNestedAttribute{
																							Description:         "ConfigMap containing data to use for the targets.",
																							MarkdownDescription: "ConfigMap containing data to use for the targets.",
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

																						"secret": schema.SingleNestedAttribute{
																							Description:         "Secret containing data to use for the targets.",
																							MarkdownDescription: "Secret containing data to use for the targets.",
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

																				"ca_file": schema.StringAttribute{
																					Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
																					MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cert": schema.SingleNestedAttribute{
																					Description:         "Client certificate to present when doing client-authentication.",
																					MarkdownDescription: "Client certificate to present when doing client-authentication.",
																					Attributes: map[string]schema.Attribute{
																						"config_map": schema.SingleNestedAttribute{
																							Description:         "ConfigMap containing data to use for the targets.",
																							MarkdownDescription: "ConfigMap containing data to use for the targets.",
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

																						"secret": schema.SingleNestedAttribute{
																							Description:         "Secret containing data to use for the targets.",
																							MarkdownDescription: "Secret containing data to use for the targets.",
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

																				"cert_file": schema.StringAttribute{
																					Description:         "Path to the client cert file in the Prometheus container for the targets.",
																					MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"insecure_skip_verify": schema.BoolAttribute{
																					Description:         "Disable target certificate validation.",
																					MarkdownDescription: "Disable target certificate validation.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"key_file": schema.StringAttribute{
																					Description:         "Path to the client key file in the Prometheus container for the targets.",
																					MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"key_secret": schema.SingleNestedAttribute{
																					Description:         "Secret containing the client key file for the targets.",
																					MarkdownDescription: "Secret containing the client key file for the targets.",
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

																				"server_name": schema.StringAttribute{
																					Description:         "Used to verify the hostname for the targets.",
																					MarkdownDescription: "Used to verify the hostname for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"track_timestamps_staleness": schema.BoolAttribute{
																			Description:         "'trackTimestampsStaleness' defines whether Prometheus tracks staleness of the metrics that have an explicit timestamp present in scraped data. Has no effect if 'honorTimestamps' is false.  It requires Prometheus >= v2.48.0.",
																			MarkdownDescription: "'trackTimestampsStaleness' defines whether Prometheus tracks staleness of the metrics that have an explicit timestamp present in scraped data. Has no effect if 'honorTimestamps' is false.  It requires Prometheus >= v2.48.0.",
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

															"job_label": schema.StringAttribute{
																Description:         "'jobLabel' selects the label from the associated Kubernetes 'Service' object which will be used as the 'job' label for all metrics.  For example if 'jobLabel' is set to 'foo' and the Kubernetes 'Service' object is labeled with 'foo: bar', then Prometheus adds the 'job='bar'' label to all ingested metrics.  If the value of this field is empty or if the label doesn't exist for the given Service, the 'job' label of the metrics defaults to the name of the associated Kubernetes 'Service'.",
																MarkdownDescription: "'jobLabel' selects the label from the associated Kubernetes 'Service' object which will be used as the 'job' label for all metrics.  For example if 'jobLabel' is set to 'foo' and the Kubernetes 'Service' object is labeled with 'foo: bar', then Prometheus adds the 'job='bar'' label to all ingested metrics.  If the value of this field is empty or if the label doesn't exist for the given Service, the 'job' label of the metrics defaults to the name of the associated Kubernetes 'Service'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"keep_dropped_targets": schema.Int64Attribute{
																Description:         "Per-scrape limit on the number of targets dropped by relabeling that will be kept in memory. 0 means no limit.  It requires Prometheus >= v2.47.0.",
																MarkdownDescription: "Per-scrape limit on the number of targets dropped by relabeling that will be kept in memory. 0 means no limit.  It requires Prometheus >= v2.47.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label_limit": schema.Int64Attribute{
																Description:         "Per-scrape limit on number of labels that will be accepted for a sample.  It requires Prometheus >= v2.27.0.",
																MarkdownDescription: "Per-scrape limit on number of labels that will be accepted for a sample.  It requires Prometheus >= v2.27.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label_name_length_limit": schema.Int64Attribute{
																Description:         "Per-scrape limit on length of labels name that will be accepted for a sample.  It requires Prometheus >= v2.27.0.",
																MarkdownDescription: "Per-scrape limit on length of labels name that will be accepted for a sample.  It requires Prometheus >= v2.27.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label_value_length_limit": schema.Int64Attribute{
																Description:         "Per-scrape limit on length of labels value that will be accepted for a sample.  It requires Prometheus >= v2.27.0.",
																MarkdownDescription: "Per-scrape limit on length of labels value that will be accepted for a sample.  It requires Prometheus >= v2.27.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "Selector to select which namespaces the Kubernetes 'Endpoints' objects are discovered from.",
																MarkdownDescription: "Selector to select which namespaces the Kubernetes 'Endpoints' objects are discovered from.",
																Attributes: map[string]schema.Attribute{
																	"any": schema.BoolAttribute{
																		Description:         "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
																		MarkdownDescription: "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"match_names": schema.ListAttribute{
																		Description:         "List of namespace names to select from.",
																		MarkdownDescription: "List of namespace names to select from.",
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

															"pod_target_labels": schema.ListAttribute{
																Description:         "'podTargetLabels' defines the labels which are transferred from the associated Kubernetes 'Pod' object onto the ingested metrics.",
																MarkdownDescription: "'podTargetLabels' defines the labels which are transferred from the associated Kubernetes 'Pod' object onto the ingested metrics.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sample_limit": schema.Int64Attribute{
																Description:         "'sampleLimit' defines a per-scrape limit on the number of scraped samples that will be accepted.",
																MarkdownDescription: "'sampleLimit' defines a per-scrape limit on the number of scraped samples that will be accepted.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "Label selector to select the Kubernetes 'Endpoints' objects.",
																MarkdownDescription: "Label selector to select the Kubernetes 'Endpoints' objects.",
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

															"target_labels": schema.ListAttribute{
																Description:         "'targetLabels' defines the labels which are transferred from the associated Kubernetes 'Service' object onto the ingested metrics.",
																MarkdownDescription: "'targetLabels' defines the labels which are transferred from the associated Kubernetes 'Service' object onto the ingested metrics.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_limit": schema.Int64Attribute{
																Description:         "'targetLimit' defines a limit on the number of scraped targets that will be accepted.",
																MarkdownDescription: "'targetLimit' defines a limit on the number of scraped targets that will be accepted.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"num_history_shards": schema.Int64Attribute{
						Description:         "NumHistoryShards is the desired number of history shards. This field is immutable.",
						MarkdownDescription: "NumHistoryShards is the desired number of history shards. This field is immutable.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"persistence": schema.SingleNestedAttribute{
						Description:         "Persistence defines temporal persistence configuration.",
						MarkdownDescription: "Persistence defines temporal persistence configuration.",
						Attributes: map[string]schema.Attribute{
							"advanced_visibility_store": schema.SingleNestedAttribute{
								Description:         "AdvancedVisibilityStore holds the advanced visibility datastore specs.",
								MarkdownDescription: "AdvancedVisibilityStore holds the advanced visibility datastore specs.",
								Attributes: map[string]schema.Attribute{
									"cassandra": schema.SingleNestedAttribute{
										Description:         "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										MarkdownDescription: "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										Attributes: map[string]schema.Attribute{
											"connect_timeout": schema.StringAttribute{
												Description:         "ConnectTimeout is a timeout for initial dial to cassandra server.",
												MarkdownDescription: "ConnectTimeout is a timeout for initial dial to cassandra server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"consistency": schema.SingleNestedAttribute{
												Description:         "Consistency configuration.",
												MarkdownDescription: "Consistency configuration.",
												Attributes: map[string]schema.Attribute{
													"consistency": schema.Int64Attribute{
														Description:         "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														MarkdownDescription: "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},

													"serial_consistency": schema.Int64Attribute{
														Description:         "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														MarkdownDescription: "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"datacenter": schema.StringAttribute{
												Description:         "Datacenter is the data center filter arg for cassandra.",
												MarkdownDescription: "Datacenter is the data center filter arg for cassandra.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_initial_host_lookup": schema.BoolAttribute{
												Description:         "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												MarkdownDescription: "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"hosts": schema.ListAttribute{
												Description:         "Hosts is a list of cassandra endpoints.",
												MarkdownDescription: "Hosts is a list of cassandra endpoints.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"keyspace": schema.StringAttribute{
												Description:         "Keyspace is the cassandra keyspace.",
												MarkdownDescription: "Keyspace is the cassandra keyspace.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns is the max number of connections to this datastore for a single keyspace.",
												MarkdownDescription: "MaxConns is the max number of connections to this datastore for a single keyspace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the cassandra port used for connection by gocql client.",
												MarkdownDescription: "Port is the cassandra port used for connection by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the cassandra user used for authentication by gocql client.",
												MarkdownDescription: "User is the cassandra user used for authentication by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"elasticsearch": schema.SingleNestedAttribute{
										Description:         "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										MarkdownDescription: "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										Attributes: map[string]schema.Attribute{
											"close_idle_connections_interval": schema.StringAttribute{
												Description:         "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												MarkdownDescription: "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_healthcheck": schema.BoolAttribute{
												Description:         "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												MarkdownDescription: "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_sniff": schema.BoolAttribute{
												Description:         "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												MarkdownDescription: "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"indices": schema.SingleNestedAttribute{
												Description:         "Indices holds visibility index names.",
												MarkdownDescription: "Indices holds visibility index names.",
												Attributes: map[string]schema.Attribute{
													"secondary_visibility": schema.StringAttribute{
														Description:         "SecondaryVisibility defines secondary visibility's index name.",
														MarkdownDescription: "SecondaryVisibility defines secondary visibility's index name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"visibility": schema.StringAttribute{
														Description:         "Visibility defines visibility's index name.",
														MarkdownDescription: "Visibility defines visibility's index name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"log_level": schema.StringAttribute{
												Description:         "LogLevel defines the temporal cluster's es client logger level.",
												MarkdownDescription: "LogLevel defines the temporal cluster's es client logger level.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "URL is the connection url to connect to the instance.",
												MarkdownDescription: "URL is the connection url to connect to the instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.+$`), ""),
												},
											},

											"username": schema.StringAttribute{
												Description:         "Username is the username to be used for the connection.",
												MarkdownDescription: "Username is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "Version defines the elasticsearch version.",
												MarkdownDescription: "Version defines the elasticsearch version.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^v(6|7|8)$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										MarkdownDescription: "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecret is the reference to the secret holding the password.",
										MarkdownDescription: "PasswordSecret is the reference to the secret holding the password.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key in the Secret.",
												MarkdownDescription: "Key in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the Secret.",
												MarkdownDescription: "Name of the Secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"skip_create": schema.BoolAttribute{
										Description:         "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										MarkdownDescription: "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sql": schema.SingleNestedAttribute{
										Description:         "SQL holds all connection parameters for SQL datastores.",
										MarkdownDescription: "SQL holds all connection parameters for SQL datastores.",
										Attributes: map[string]schema.Attribute{
											"connect_addr": schema.StringAttribute{
												Description:         "ConnectAddr is the remote addr of the database.",
												MarkdownDescription: "ConnectAddr is the remote addr of the database.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"connect_attributes": schema.MapAttribute{
												Description:         "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												MarkdownDescription: "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connect_protocol": schema.StringAttribute{
												Description:         "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												MarkdownDescription: "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"database_name": schema.StringAttribute{
												Description:         "DatabaseName is the name of SQL database to connect to.",
												MarkdownDescription: "DatabaseName is the name of SQL database to connect to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"gcp_service_account": schema.StringAttribute{
												Description:         "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												MarkdownDescription: "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conn_lifetime": schema.StringAttribute{
												Description:         "MaxConnLifetime is the maximum time a connection can be alive",
												MarkdownDescription: "MaxConnLifetime is the maximum time a connection can be alive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns the max number of connections to this datastore.",
												MarkdownDescription: "MaxConns the max number of connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_idle_conns": schema.Int64Attribute{
												Description:         "MaxIdleConns is the max number of idle connections to this datastore.",
												MarkdownDescription: "MaxIdleConns is the max number of idle connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"plugin_name": schema.StringAttribute{
												Description:         "PluginName is the name of SQL plugin.",
												MarkdownDescription: "PluginName is the name of SQL plugin.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("postgres", "postgres12", "mysql", "mysql8"),
												},
											},

											"task_scan_partitions": schema.Int64Attribute{
												Description:         "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												MarkdownDescription: "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the username to be used for the connection.",
												MarkdownDescription: "User is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS is an optional option to connect to the datastore using TLS.",
										MarkdownDescription: "TLS is an optional option to connect to the datastore using TLS.",
										Attributes: map[string]schema.Attribute{
											"ca_file_ref": schema.SingleNestedAttribute{
												Description:         "CaFileRef is a reference to a secret containing the ca file.",
												MarkdownDescription: "CaFileRef is a reference to a secret containing the ca file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert_file_ref": schema.SingleNestedAttribute{
												Description:         "CertFileRef is a reference to a secret containing the cert file.",
												MarkdownDescription: "CertFileRef is a reference to a secret containing the cert file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_host_verification": schema.BoolAttribute{
												Description:         "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												MarkdownDescription: "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												MarkdownDescription: "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"key_file_ref": schema.SingleNestedAttribute{
												Description:         "KeyFileRef is a reference to a secret containing the key file.",
												MarkdownDescription: "KeyFileRef is a reference to a secret containing the key file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_name": schema.StringAttribute{
												Description:         "ServerName the datastore should present.",
												MarkdownDescription: "ServerName the datastore should present.",
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

							"default_store": schema.SingleNestedAttribute{
								Description:         "DefaultStore holds the default datastore specs.",
								MarkdownDescription: "DefaultStore holds the default datastore specs.",
								Attributes: map[string]schema.Attribute{
									"cassandra": schema.SingleNestedAttribute{
										Description:         "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										MarkdownDescription: "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										Attributes: map[string]schema.Attribute{
											"connect_timeout": schema.StringAttribute{
												Description:         "ConnectTimeout is a timeout for initial dial to cassandra server.",
												MarkdownDescription: "ConnectTimeout is a timeout for initial dial to cassandra server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"consistency": schema.SingleNestedAttribute{
												Description:         "Consistency configuration.",
												MarkdownDescription: "Consistency configuration.",
												Attributes: map[string]schema.Attribute{
													"consistency": schema.Int64Attribute{
														Description:         "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														MarkdownDescription: "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},

													"serial_consistency": schema.Int64Attribute{
														Description:         "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														MarkdownDescription: "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"datacenter": schema.StringAttribute{
												Description:         "Datacenter is the data center filter arg for cassandra.",
												MarkdownDescription: "Datacenter is the data center filter arg for cassandra.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_initial_host_lookup": schema.BoolAttribute{
												Description:         "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												MarkdownDescription: "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"hosts": schema.ListAttribute{
												Description:         "Hosts is a list of cassandra endpoints.",
												MarkdownDescription: "Hosts is a list of cassandra endpoints.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"keyspace": schema.StringAttribute{
												Description:         "Keyspace is the cassandra keyspace.",
												MarkdownDescription: "Keyspace is the cassandra keyspace.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns is the max number of connections to this datastore for a single keyspace.",
												MarkdownDescription: "MaxConns is the max number of connections to this datastore for a single keyspace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the cassandra port used for connection by gocql client.",
												MarkdownDescription: "Port is the cassandra port used for connection by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the cassandra user used for authentication by gocql client.",
												MarkdownDescription: "User is the cassandra user used for authentication by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"elasticsearch": schema.SingleNestedAttribute{
										Description:         "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										MarkdownDescription: "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										Attributes: map[string]schema.Attribute{
											"close_idle_connections_interval": schema.StringAttribute{
												Description:         "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												MarkdownDescription: "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_healthcheck": schema.BoolAttribute{
												Description:         "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												MarkdownDescription: "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_sniff": schema.BoolAttribute{
												Description:         "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												MarkdownDescription: "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"indices": schema.SingleNestedAttribute{
												Description:         "Indices holds visibility index names.",
												MarkdownDescription: "Indices holds visibility index names.",
												Attributes: map[string]schema.Attribute{
													"secondary_visibility": schema.StringAttribute{
														Description:         "SecondaryVisibility defines secondary visibility's index name.",
														MarkdownDescription: "SecondaryVisibility defines secondary visibility's index name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"visibility": schema.StringAttribute{
														Description:         "Visibility defines visibility's index name.",
														MarkdownDescription: "Visibility defines visibility's index name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"log_level": schema.StringAttribute{
												Description:         "LogLevel defines the temporal cluster's es client logger level.",
												MarkdownDescription: "LogLevel defines the temporal cluster's es client logger level.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "URL is the connection url to connect to the instance.",
												MarkdownDescription: "URL is the connection url to connect to the instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.+$`), ""),
												},
											},

											"username": schema.StringAttribute{
												Description:         "Username is the username to be used for the connection.",
												MarkdownDescription: "Username is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "Version defines the elasticsearch version.",
												MarkdownDescription: "Version defines the elasticsearch version.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^v(6|7|8)$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										MarkdownDescription: "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecret is the reference to the secret holding the password.",
										MarkdownDescription: "PasswordSecret is the reference to the secret holding the password.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key in the Secret.",
												MarkdownDescription: "Key in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the Secret.",
												MarkdownDescription: "Name of the Secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"skip_create": schema.BoolAttribute{
										Description:         "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										MarkdownDescription: "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sql": schema.SingleNestedAttribute{
										Description:         "SQL holds all connection parameters for SQL datastores.",
										MarkdownDescription: "SQL holds all connection parameters for SQL datastores.",
										Attributes: map[string]schema.Attribute{
											"connect_addr": schema.StringAttribute{
												Description:         "ConnectAddr is the remote addr of the database.",
												MarkdownDescription: "ConnectAddr is the remote addr of the database.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"connect_attributes": schema.MapAttribute{
												Description:         "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												MarkdownDescription: "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connect_protocol": schema.StringAttribute{
												Description:         "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												MarkdownDescription: "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"database_name": schema.StringAttribute{
												Description:         "DatabaseName is the name of SQL database to connect to.",
												MarkdownDescription: "DatabaseName is the name of SQL database to connect to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"gcp_service_account": schema.StringAttribute{
												Description:         "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												MarkdownDescription: "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conn_lifetime": schema.StringAttribute{
												Description:         "MaxConnLifetime is the maximum time a connection can be alive",
												MarkdownDescription: "MaxConnLifetime is the maximum time a connection can be alive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns the max number of connections to this datastore.",
												MarkdownDescription: "MaxConns the max number of connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_idle_conns": schema.Int64Attribute{
												Description:         "MaxIdleConns is the max number of idle connections to this datastore.",
												MarkdownDescription: "MaxIdleConns is the max number of idle connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"plugin_name": schema.StringAttribute{
												Description:         "PluginName is the name of SQL plugin.",
												MarkdownDescription: "PluginName is the name of SQL plugin.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("postgres", "postgres12", "mysql", "mysql8"),
												},
											},

											"task_scan_partitions": schema.Int64Attribute{
												Description:         "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												MarkdownDescription: "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the username to be used for the connection.",
												MarkdownDescription: "User is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS is an optional option to connect to the datastore using TLS.",
										MarkdownDescription: "TLS is an optional option to connect to the datastore using TLS.",
										Attributes: map[string]schema.Attribute{
											"ca_file_ref": schema.SingleNestedAttribute{
												Description:         "CaFileRef is a reference to a secret containing the ca file.",
												MarkdownDescription: "CaFileRef is a reference to a secret containing the ca file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert_file_ref": schema.SingleNestedAttribute{
												Description:         "CertFileRef is a reference to a secret containing the cert file.",
												MarkdownDescription: "CertFileRef is a reference to a secret containing the cert file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_host_verification": schema.BoolAttribute{
												Description:         "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												MarkdownDescription: "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												MarkdownDescription: "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"key_file_ref": schema.SingleNestedAttribute{
												Description:         "KeyFileRef is a reference to a secret containing the key file.",
												MarkdownDescription: "KeyFileRef is a reference to a secret containing the key file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_name": schema.StringAttribute{
												Description:         "ServerName the datastore should present.",
												MarkdownDescription: "ServerName the datastore should present.",
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

							"secondary_visibility_store": schema.SingleNestedAttribute{
								Description:         "SecondaryVisibilityStore holds the secondary visibility datastore specs. Feature only available for clusters >= 1.21.0.",
								MarkdownDescription: "SecondaryVisibilityStore holds the secondary visibility datastore specs. Feature only available for clusters >= 1.21.0.",
								Attributes: map[string]schema.Attribute{
									"cassandra": schema.SingleNestedAttribute{
										Description:         "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										MarkdownDescription: "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										Attributes: map[string]schema.Attribute{
											"connect_timeout": schema.StringAttribute{
												Description:         "ConnectTimeout is a timeout for initial dial to cassandra server.",
												MarkdownDescription: "ConnectTimeout is a timeout for initial dial to cassandra server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"consistency": schema.SingleNestedAttribute{
												Description:         "Consistency configuration.",
												MarkdownDescription: "Consistency configuration.",
												Attributes: map[string]schema.Attribute{
													"consistency": schema.Int64Attribute{
														Description:         "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														MarkdownDescription: "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},

													"serial_consistency": schema.Int64Attribute{
														Description:         "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														MarkdownDescription: "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"datacenter": schema.StringAttribute{
												Description:         "Datacenter is the data center filter arg for cassandra.",
												MarkdownDescription: "Datacenter is the data center filter arg for cassandra.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_initial_host_lookup": schema.BoolAttribute{
												Description:         "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												MarkdownDescription: "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"hosts": schema.ListAttribute{
												Description:         "Hosts is a list of cassandra endpoints.",
												MarkdownDescription: "Hosts is a list of cassandra endpoints.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"keyspace": schema.StringAttribute{
												Description:         "Keyspace is the cassandra keyspace.",
												MarkdownDescription: "Keyspace is the cassandra keyspace.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns is the max number of connections to this datastore for a single keyspace.",
												MarkdownDescription: "MaxConns is the max number of connections to this datastore for a single keyspace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the cassandra port used for connection by gocql client.",
												MarkdownDescription: "Port is the cassandra port used for connection by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the cassandra user used for authentication by gocql client.",
												MarkdownDescription: "User is the cassandra user used for authentication by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"elasticsearch": schema.SingleNestedAttribute{
										Description:         "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										MarkdownDescription: "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										Attributes: map[string]schema.Attribute{
											"close_idle_connections_interval": schema.StringAttribute{
												Description:         "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												MarkdownDescription: "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_healthcheck": schema.BoolAttribute{
												Description:         "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												MarkdownDescription: "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_sniff": schema.BoolAttribute{
												Description:         "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												MarkdownDescription: "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"indices": schema.SingleNestedAttribute{
												Description:         "Indices holds visibility index names.",
												MarkdownDescription: "Indices holds visibility index names.",
												Attributes: map[string]schema.Attribute{
													"secondary_visibility": schema.StringAttribute{
														Description:         "SecondaryVisibility defines secondary visibility's index name.",
														MarkdownDescription: "SecondaryVisibility defines secondary visibility's index name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"visibility": schema.StringAttribute{
														Description:         "Visibility defines visibility's index name.",
														MarkdownDescription: "Visibility defines visibility's index name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"log_level": schema.StringAttribute{
												Description:         "LogLevel defines the temporal cluster's es client logger level.",
												MarkdownDescription: "LogLevel defines the temporal cluster's es client logger level.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "URL is the connection url to connect to the instance.",
												MarkdownDescription: "URL is the connection url to connect to the instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.+$`), ""),
												},
											},

											"username": schema.StringAttribute{
												Description:         "Username is the username to be used for the connection.",
												MarkdownDescription: "Username is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "Version defines the elasticsearch version.",
												MarkdownDescription: "Version defines the elasticsearch version.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^v(6|7|8)$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										MarkdownDescription: "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecret is the reference to the secret holding the password.",
										MarkdownDescription: "PasswordSecret is the reference to the secret holding the password.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key in the Secret.",
												MarkdownDescription: "Key in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the Secret.",
												MarkdownDescription: "Name of the Secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"skip_create": schema.BoolAttribute{
										Description:         "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										MarkdownDescription: "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sql": schema.SingleNestedAttribute{
										Description:         "SQL holds all connection parameters for SQL datastores.",
										MarkdownDescription: "SQL holds all connection parameters for SQL datastores.",
										Attributes: map[string]schema.Attribute{
											"connect_addr": schema.StringAttribute{
												Description:         "ConnectAddr is the remote addr of the database.",
												MarkdownDescription: "ConnectAddr is the remote addr of the database.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"connect_attributes": schema.MapAttribute{
												Description:         "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												MarkdownDescription: "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connect_protocol": schema.StringAttribute{
												Description:         "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												MarkdownDescription: "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"database_name": schema.StringAttribute{
												Description:         "DatabaseName is the name of SQL database to connect to.",
												MarkdownDescription: "DatabaseName is the name of SQL database to connect to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"gcp_service_account": schema.StringAttribute{
												Description:         "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												MarkdownDescription: "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conn_lifetime": schema.StringAttribute{
												Description:         "MaxConnLifetime is the maximum time a connection can be alive",
												MarkdownDescription: "MaxConnLifetime is the maximum time a connection can be alive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns the max number of connections to this datastore.",
												MarkdownDescription: "MaxConns the max number of connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_idle_conns": schema.Int64Attribute{
												Description:         "MaxIdleConns is the max number of idle connections to this datastore.",
												MarkdownDescription: "MaxIdleConns is the max number of idle connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"plugin_name": schema.StringAttribute{
												Description:         "PluginName is the name of SQL plugin.",
												MarkdownDescription: "PluginName is the name of SQL plugin.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("postgres", "postgres12", "mysql", "mysql8"),
												},
											},

											"task_scan_partitions": schema.Int64Attribute{
												Description:         "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												MarkdownDescription: "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the username to be used for the connection.",
												MarkdownDescription: "User is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS is an optional option to connect to the datastore using TLS.",
										MarkdownDescription: "TLS is an optional option to connect to the datastore using TLS.",
										Attributes: map[string]schema.Attribute{
											"ca_file_ref": schema.SingleNestedAttribute{
												Description:         "CaFileRef is a reference to a secret containing the ca file.",
												MarkdownDescription: "CaFileRef is a reference to a secret containing the ca file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert_file_ref": schema.SingleNestedAttribute{
												Description:         "CertFileRef is a reference to a secret containing the cert file.",
												MarkdownDescription: "CertFileRef is a reference to a secret containing the cert file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_host_verification": schema.BoolAttribute{
												Description:         "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												MarkdownDescription: "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												MarkdownDescription: "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"key_file_ref": schema.SingleNestedAttribute{
												Description:         "KeyFileRef is a reference to a secret containing the key file.",
												MarkdownDescription: "KeyFileRef is a reference to a secret containing the key file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_name": schema.StringAttribute{
												Description:         "ServerName the datastore should present.",
												MarkdownDescription: "ServerName the datastore should present.",
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

							"visibility_store": schema.SingleNestedAttribute{
								Description:         "VisibilityStore holds the visibility datastore specs.",
								MarkdownDescription: "VisibilityStore holds the visibility datastore specs.",
								Attributes: map[string]schema.Attribute{
									"cassandra": schema.SingleNestedAttribute{
										Description:         "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										MarkdownDescription: "Cassandra holds all connection parameters for Cassandra datastore. Note that cassandra is now deprecated for visibility store.",
										Attributes: map[string]schema.Attribute{
											"connect_timeout": schema.StringAttribute{
												Description:         "ConnectTimeout is a timeout for initial dial to cassandra server.",
												MarkdownDescription: "ConnectTimeout is a timeout for initial dial to cassandra server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"consistency": schema.SingleNestedAttribute{
												Description:         "Consistency configuration.",
												MarkdownDescription: "Consistency configuration.",
												Attributes: map[string]schema.Attribute{
													"consistency": schema.Int64Attribute{
														Description:         "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														MarkdownDescription: "Consistency sets the default consistency level. Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},

													"serial_consistency": schema.Int64Attribute{
														Description:         "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														MarkdownDescription: "SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values. (defaults to LOCAL_SERIAL if not set)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.OneOf(),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"datacenter": schema.StringAttribute{
												Description:         "Datacenter is the data center filter arg for cassandra.",
												MarkdownDescription: "Datacenter is the data center filter arg for cassandra.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_initial_host_lookup": schema.BoolAttribute{
												Description:         "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												MarkdownDescription: "DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"hosts": schema.ListAttribute{
												Description:         "Hosts is a list of cassandra endpoints.",
												MarkdownDescription: "Hosts is a list of cassandra endpoints.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"keyspace": schema.StringAttribute{
												Description:         "Keyspace is the cassandra keyspace.",
												MarkdownDescription: "Keyspace is the cassandra keyspace.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns is the max number of connections to this datastore for a single keyspace.",
												MarkdownDescription: "MaxConns is the max number of connections to this datastore for a single keyspace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port is the cassandra port used for connection by gocql client.",
												MarkdownDescription: "Port is the cassandra port used for connection by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the cassandra user used for authentication by gocql client.",
												MarkdownDescription: "User is the cassandra user used for authentication by gocql client.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"elasticsearch": schema.SingleNestedAttribute{
										Description:         "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										MarkdownDescription: "Elasticsearch holds all connection parameters for Elasticsearch datastores.",
										Attributes: map[string]schema.Attribute{
											"close_idle_connections_interval": schema.StringAttribute{
												Description:         "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												MarkdownDescription: "CloseIdleConnectionsInterval is the max duration a connection stay open while idle.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_healthcheck": schema.BoolAttribute{
												Description:         "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												MarkdownDescription: "EnableHealthcheck enables or disables healthcheck on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable_sniff": schema.BoolAttribute{
												Description:         "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												MarkdownDescription: "EnableSniff enables or disables sniffer on the temporal cluster's es client.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"indices": schema.SingleNestedAttribute{
												Description:         "Indices holds visibility index names.",
												MarkdownDescription: "Indices holds visibility index names.",
												Attributes: map[string]schema.Attribute{
													"secondary_visibility": schema.StringAttribute{
														Description:         "SecondaryVisibility defines secondary visibility's index name.",
														MarkdownDescription: "SecondaryVisibility defines secondary visibility's index name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"visibility": schema.StringAttribute{
														Description:         "Visibility defines visibility's index name.",
														MarkdownDescription: "Visibility defines visibility's index name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"log_level": schema.StringAttribute{
												Description:         "LogLevel defines the temporal cluster's es client logger level.",
												MarkdownDescription: "LogLevel defines the temporal cluster's es client logger level.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "URL is the connection url to connect to the instance.",
												MarkdownDescription: "URL is the connection url to connect to the instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^https?:\/\/.+$`), ""),
												},
											},

											"username": schema.StringAttribute{
												Description:         "Username is the username to be used for the connection.",
												MarkdownDescription: "Username is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "Version defines the elasticsearch version.",
												MarkdownDescription: "Version defines the elasticsearch version.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^v(6|7|8)$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										MarkdownDescription: "Name is the name of the datastore. It should be unique and will be referenced within the persistence spec. Defaults to 'default' for default sore, 'visibility' for visibility store, 'secondaryVisibility' for secondary visibility store and 'advancedVisibility' for advanced visibility store.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecret is the reference to the secret holding the password.",
										MarkdownDescription: "PasswordSecret is the reference to the secret holding the password.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key in the Secret.",
												MarkdownDescription: "Key in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the Secret.",
												MarkdownDescription: "Name of the Secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"skip_create": schema.BoolAttribute{
										Description:         "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										MarkdownDescription: "SkipCreate instructs the operator to skip creating the database for SQL datastores or to skip creating keyspace for Cassandra. Use this option if your database or keyspace has already been provisioned by an administrator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sql": schema.SingleNestedAttribute{
										Description:         "SQL holds all connection parameters for SQL datastores.",
										MarkdownDescription: "SQL holds all connection parameters for SQL datastores.",
										Attributes: map[string]schema.Attribute{
											"connect_addr": schema.StringAttribute{
												Description:         "ConnectAddr is the remote addr of the database.",
												MarkdownDescription: "ConnectAddr is the remote addr of the database.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"connect_attributes": schema.MapAttribute{
												Description:         "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												MarkdownDescription: "ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connect_protocol": schema.StringAttribute{
												Description:         "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												MarkdownDescription: "ConnectProtocol is the protocol that goes with the ConnectAddr.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"database_name": schema.StringAttribute{
												Description:         "DatabaseName is the name of SQL database to connect to.",
												MarkdownDescription: "DatabaseName is the name of SQL database to connect to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"gcp_service_account": schema.StringAttribute{
												Description:         "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												MarkdownDescription: "GCPServiceAccount is the service account to use to authenticate with GCP CloudSQL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conn_lifetime": schema.StringAttribute{
												Description:         "MaxConnLifetime is the maximum time a connection can be alive",
												MarkdownDescription: "MaxConnLifetime is the maximum time a connection can be alive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_conns": schema.Int64Attribute{
												Description:         "MaxConns the max number of connections to this datastore.",
												MarkdownDescription: "MaxConns the max number of connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_idle_conns": schema.Int64Attribute{
												Description:         "MaxIdleConns is the max number of idle connections to this datastore.",
												MarkdownDescription: "MaxIdleConns is the max number of idle connections to this datastore.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"plugin_name": schema.StringAttribute{
												Description:         "PluginName is the name of SQL plugin.",
												MarkdownDescription: "PluginName is the name of SQL plugin.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("postgres", "postgres12", "mysql", "mysql8"),
												},
											},

											"task_scan_partitions": schema.Int64Attribute{
												Description:         "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												MarkdownDescription: "TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is the username to be used for the connection.",
												MarkdownDescription: "User is the username to be used for the connection.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS is an optional option to connect to the datastore using TLS.",
										MarkdownDescription: "TLS is an optional option to connect to the datastore using TLS.",
										Attributes: map[string]schema.Attribute{
											"ca_file_ref": schema.SingleNestedAttribute{
												Description:         "CaFileRef is a reference to a secret containing the ca file.",
												MarkdownDescription: "CaFileRef is a reference to a secret containing the ca file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert_file_ref": schema.SingleNestedAttribute{
												Description:         "CertFileRef is a reference to a secret containing the cert file.",
												MarkdownDescription: "CertFileRef is a reference to a secret containing the cert file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_host_verification": schema.BoolAttribute{
												Description:         "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												MarkdownDescription: "EnableHostVerification defines if the hostname should be verified when connecting to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												MarkdownDescription: "Enabled defines if the cluster should use a TLS connection to connect to the datastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"key_file_ref": schema.SingleNestedAttribute{
												Description:         "KeyFileRef is a reference to a secret containing the key file.",
												MarkdownDescription: "KeyFileRef is a reference to a secret containing the key file.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key in the Secret.",
														MarkdownDescription: "Key in the Secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the Secret.",
														MarkdownDescription: "Name of the Secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_name": schema.StringAttribute{
												Description:         "ServerName the datastore should present.",
												MarkdownDescription: "ServerName the datastore should present.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"services": schema.SingleNestedAttribute{
						Description:         "Services allows customizations for each temporal services deployment.",
						MarkdownDescription: "Services allows customizations for each temporal services deployment.",
						Attributes: map[string]schema.Attribute{
							"frontend": schema.SingleNestedAttribute{
								Description:         "Frontend service custom specifications.",
								MarkdownDescription: "Frontend service custom specifications.",
								Attributes: map[string]schema.Attribute{
									"http_port": schema.Int64Attribute{
										Description:         "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										MarkdownDescription: "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"init_containers": schema.ListAttribute{
										Description:         "InitContainers adds a list of init containers to the service's deployment.",
										MarkdownDescription: "InitContainers adds a list of init containers to the service's deployment.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"membership_port": schema.Int64Attribute{
										Description:         "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										MarkdownDescription: "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										MarkdownDescription: "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										Attributes: map[string]schema.Attribute{
											"deployment": schema.SingleNestedAttribute{
												Description:         "Override configuration for the temporal service Deployment.",
												MarkdownDescription: "Override configuration for the temporal service Deployment.",
												Attributes: map[string]schema.Attribute{
													"metadata": schema.SingleNestedAttribute{
														Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
														Description:         "Specification of the desired behavior of the Deployment.",
														MarkdownDescription: "Specification of the desired behavior of the Deployment.",
														Attributes: map[string]schema.Attribute{
															"template": schema.SingleNestedAttribute{
																Description:         "Template describes the pods that will be created.",
																MarkdownDescription: "Template describes the pods that will be created.",
																Attributes: map[string]schema.Attribute{
																	"metadata": schema.SingleNestedAttribute{
																		Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		Attributes: map[string]schema.Attribute{
																			"annotations": schema.MapAttribute{
																				Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"labels": schema.MapAttribute{
																				Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																				MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

																	"spec": schema.MapAttribute{
																		Description:         "Specification of the desired behavior of the pod.",
																		MarkdownDescription: "Specification of the desired behavior of the pod.",
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

									"port": schema.Int64Attribute{
										Description:         "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										MarkdownDescription: "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of desired replicas for the service. Default to 1.",
										MarkdownDescription: "Number of desired replicas for the service. Default to 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"history": schema.SingleNestedAttribute{
								Description:         "History service custom specifications.",
								MarkdownDescription: "History service custom specifications.",
								Attributes: map[string]schema.Attribute{
									"http_port": schema.Int64Attribute{
										Description:         "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										MarkdownDescription: "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"init_containers": schema.ListAttribute{
										Description:         "InitContainers adds a list of init containers to the service's deployment.",
										MarkdownDescription: "InitContainers adds a list of init containers to the service's deployment.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"membership_port": schema.Int64Attribute{
										Description:         "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										MarkdownDescription: "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										MarkdownDescription: "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										Attributes: map[string]schema.Attribute{
											"deployment": schema.SingleNestedAttribute{
												Description:         "Override configuration for the temporal service Deployment.",
												MarkdownDescription: "Override configuration for the temporal service Deployment.",
												Attributes: map[string]schema.Attribute{
													"metadata": schema.SingleNestedAttribute{
														Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
														Description:         "Specification of the desired behavior of the Deployment.",
														MarkdownDescription: "Specification of the desired behavior of the Deployment.",
														Attributes: map[string]schema.Attribute{
															"template": schema.SingleNestedAttribute{
																Description:         "Template describes the pods that will be created.",
																MarkdownDescription: "Template describes the pods that will be created.",
																Attributes: map[string]schema.Attribute{
																	"metadata": schema.SingleNestedAttribute{
																		Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		Attributes: map[string]schema.Attribute{
																			"annotations": schema.MapAttribute{
																				Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"labels": schema.MapAttribute{
																				Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																				MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

																	"spec": schema.MapAttribute{
																		Description:         "Specification of the desired behavior of the pod.",
																		MarkdownDescription: "Specification of the desired behavior of the pod.",
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

									"port": schema.Int64Attribute{
										Description:         "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										MarkdownDescription: "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of desired replicas for the service. Default to 1.",
										MarkdownDescription: "Number of desired replicas for the service. Default to 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"internal_frontend": schema.SingleNestedAttribute{
								Description:         "Internal Frontend service custom specifications. Only compatible with temporal >= 1.20.0",
								MarkdownDescription: "Internal Frontend service custom specifications. Only compatible with temporal >= 1.20.0",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if we want to spawn the internal frontend service.",
										MarkdownDescription: "Enabled defines if we want to spawn the internal frontend service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http_port": schema.Int64Attribute{
										Description:         "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										MarkdownDescription: "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"init_containers": schema.ListAttribute{
										Description:         "InitContainers adds a list of init containers to the service's deployment.",
										MarkdownDescription: "InitContainers adds a list of init containers to the service's deployment.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"membership_port": schema.Int64Attribute{
										Description:         "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										MarkdownDescription: "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										MarkdownDescription: "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										Attributes: map[string]schema.Attribute{
											"deployment": schema.SingleNestedAttribute{
												Description:         "Override configuration for the temporal service Deployment.",
												MarkdownDescription: "Override configuration for the temporal service Deployment.",
												Attributes: map[string]schema.Attribute{
													"metadata": schema.SingleNestedAttribute{
														Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
														Description:         "Specification of the desired behavior of the Deployment.",
														MarkdownDescription: "Specification of the desired behavior of the Deployment.",
														Attributes: map[string]schema.Attribute{
															"template": schema.SingleNestedAttribute{
																Description:         "Template describes the pods that will be created.",
																MarkdownDescription: "Template describes the pods that will be created.",
																Attributes: map[string]schema.Attribute{
																	"metadata": schema.SingleNestedAttribute{
																		Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		Attributes: map[string]schema.Attribute{
																			"annotations": schema.MapAttribute{
																				Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"labels": schema.MapAttribute{
																				Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																				MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

																	"spec": schema.MapAttribute{
																		Description:         "Specification of the desired behavior of the pod.",
																		MarkdownDescription: "Specification of the desired behavior of the pod.",
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

									"port": schema.Int64Attribute{
										Description:         "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										MarkdownDescription: "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of desired replicas for the service. Default to 1.",
										MarkdownDescription: "Number of desired replicas for the service. Default to 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"matching": schema.SingleNestedAttribute{
								Description:         "Matching service custom specifications.",
								MarkdownDescription: "Matching service custom specifications.",
								Attributes: map[string]schema.Attribute{
									"http_port": schema.Int64Attribute{
										Description:         "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										MarkdownDescription: "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"init_containers": schema.ListAttribute{
										Description:         "InitContainers adds a list of init containers to the service's deployment.",
										MarkdownDescription: "InitContainers adds a list of init containers to the service's deployment.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"membership_port": schema.Int64Attribute{
										Description:         "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										MarkdownDescription: "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										MarkdownDescription: "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										Attributes: map[string]schema.Attribute{
											"deployment": schema.SingleNestedAttribute{
												Description:         "Override configuration for the temporal service Deployment.",
												MarkdownDescription: "Override configuration for the temporal service Deployment.",
												Attributes: map[string]schema.Attribute{
													"metadata": schema.SingleNestedAttribute{
														Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
														Description:         "Specification of the desired behavior of the Deployment.",
														MarkdownDescription: "Specification of the desired behavior of the Deployment.",
														Attributes: map[string]schema.Attribute{
															"template": schema.SingleNestedAttribute{
																Description:         "Template describes the pods that will be created.",
																MarkdownDescription: "Template describes the pods that will be created.",
																Attributes: map[string]schema.Attribute{
																	"metadata": schema.SingleNestedAttribute{
																		Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		Attributes: map[string]schema.Attribute{
																			"annotations": schema.MapAttribute{
																				Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"labels": schema.MapAttribute{
																				Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																				MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

																	"spec": schema.MapAttribute{
																		Description:         "Specification of the desired behavior of the pod.",
																		MarkdownDescription: "Specification of the desired behavior of the pod.",
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

									"port": schema.Int64Attribute{
										Description:         "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										MarkdownDescription: "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of desired replicas for the service. Default to 1.",
										MarkdownDescription: "Number of desired replicas for the service. Default to 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"overrides": schema.SingleNestedAttribute{
								Description:         "Overrides adds some overrides to the resources deployed for all temporal services services. Those overrides can be customized per service using spec.services.<serviceName>.overrides.",
								MarkdownDescription: "Overrides adds some overrides to the resources deployed for all temporal services services. Those overrides can be customized per service using spec.services.<serviceName>.overrides.",
								Attributes: map[string]schema.Attribute{
									"deployment": schema.SingleNestedAttribute{
										Description:         "Override configuration for the temporal service Deployment.",
										MarkdownDescription: "Override configuration for the temporal service Deployment.",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.SingleNestedAttribute{
												Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
												MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
														MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
														MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
												Description:         "Specification of the desired behavior of the Deployment.",
												MarkdownDescription: "Specification of the desired behavior of the Deployment.",
												Attributes: map[string]schema.Attribute{
													"template": schema.SingleNestedAttribute{
														Description:         "Template describes the pods that will be created.",
														MarkdownDescription: "Template describes the pods that will be created.",
														Attributes: map[string]schema.Attribute{
															"metadata": schema.SingleNestedAttribute{
																Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																Attributes: map[string]schema.Attribute{
																	"annotations": schema.MapAttribute{
																		Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																		MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"labels": schema.MapAttribute{
																		Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																		MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

															"spec": schema.MapAttribute{
																Description:         "Specification of the desired behavior of the pod.",
																MarkdownDescription: "Specification of the desired behavior of the pod.",
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

							"worker": schema.SingleNestedAttribute{
								Description:         "Worker service custom specifications.",
								MarkdownDescription: "Worker service custom specifications.",
								Attributes: map[string]schema.Attribute{
									"http_port": schema.Int64Attribute{
										Description:         "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										MarkdownDescription: "HTTPPort defines a custom http port for the service. Default values are: 7243 for Frontend service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"init_containers": schema.ListAttribute{
										Description:         "InitContainers adds a list of init containers to the service's deployment.",
										MarkdownDescription: "InitContainers adds a list of init containers to the service's deployment.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"membership_port": schema.Int64Attribute{
										Description:         "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										MarkdownDescription: "MembershipPort defines a custom membership port for the service. Default values are: 6933 for Frontend service 6934 for History service 6935 for Matching service 6939 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										MarkdownDescription: "Overrides adds some overrides to the resources deployed for the service. Those overrides takes precedence over spec.services.overrides.",
										Attributes: map[string]schema.Attribute{
											"deployment": schema.SingleNestedAttribute{
												Description:         "Override configuration for the temporal service Deployment.",
												MarkdownDescription: "Override configuration for the temporal service Deployment.",
												Attributes: map[string]schema.Attribute{
													"metadata": schema.SingleNestedAttribute{
														Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
														Attributes: map[string]schema.Attribute{
															"annotations": schema.MapAttribute{
																Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
														Description:         "Specification of the desired behavior of the Deployment.",
														MarkdownDescription: "Specification of the desired behavior of the Deployment.",
														Attributes: map[string]schema.Attribute{
															"template": schema.SingleNestedAttribute{
																Description:         "Template describes the pods that will be created.",
																MarkdownDescription: "Template describes the pods that will be created.",
																Attributes: map[string]schema.Attribute{
																	"metadata": schema.SingleNestedAttribute{
																		Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																		Attributes: map[string]schema.Attribute{
																			"annotations": schema.MapAttribute{
																				Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"labels": schema.MapAttribute{
																				Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																				MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

																	"spec": schema.MapAttribute{
																		Description:         "Specification of the desired behavior of the pod.",
																		MarkdownDescription: "Specification of the desired behavior of the pod.",
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

									"port": schema.Int64Attribute{
										Description:         "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										MarkdownDescription: "Port defines a custom gRPC port for the service. Default values are: 7233 for Frontend service 7234 for History service 7235 for Matching service 7239 for Worker service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of desired replicas for the service. Default to 1.",
										MarkdownDescription: "Number of desired replicas for the service. Default to 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Compute Resources required by this service. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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

					"ui": schema.SingleNestedAttribute{
						Description:         "UI allows configuration of the optional temporal web ui deployed alongside the cluster.",
						MarkdownDescription: "UI allows configuration of the optional temporal web ui deployed alongside the cluster.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines if the operator should deploy the web ui alongside the cluster.",
								MarkdownDescription: "Enabled defines if the operator should deploy the web ui alongside the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Image defines the temporal ui docker image the instance should run.",
								MarkdownDescription: "Image defines the temporal ui docker image the instance should run.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress": schema.SingleNestedAttribute{
								Description:         "Ingress is an optional ingress configuration for the UI. If lived empty, no ingress configuration will be created and the UI will only by available trough ClusterIP service.",
								MarkdownDescription: "Ingress is an optional ingress configuration for the UI. If lived empty, no ingress configuration will be created and the UI will only by available trough ClusterIP service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations allows custom annotations on the ingress resource.",
										MarkdownDescription: "Annotations allows custom annotations on the ingress resource.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hosts": schema.ListAttribute{
										Description:         "Host is the list of host the ingress should use.",
										MarkdownDescription: "Host is the list of host the ingress should use.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"ingress_class_name": schema.StringAttribute{
										Description:         "IngressClassName is the name of the IngressClass the deployed ingress resource should use.",
										MarkdownDescription: "IngressClassName is the name of the IngressClass the deployed ingress resource should use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.ListNestedAttribute{
										Description:         "TLS configuration.",
										MarkdownDescription: "TLS configuration.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hosts": schema.ListAttribute{
													Description:         "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
													MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
													MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
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

							"overrides": schema.SingleNestedAttribute{
								Description:         "Overrides adds some overrides to the resources deployed for the ui.",
								MarkdownDescription: "Overrides adds some overrides to the resources deployed for the ui.",
								Attributes: map[string]schema.Attribute{
									"deployment": schema.SingleNestedAttribute{
										Description:         "Override configuration for the temporal service Deployment.",
										MarkdownDescription: "Override configuration for the temporal service Deployment.",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.SingleNestedAttribute{
												Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
												MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
														MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
														MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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
												Description:         "Specification of the desired behavior of the Deployment.",
												MarkdownDescription: "Specification of the desired behavior of the Deployment.",
												Attributes: map[string]schema.Attribute{
													"template": schema.SingleNestedAttribute{
														Description:         "Template describes the pods that will be created.",
														MarkdownDescription: "Template describes the pods that will be created.",
														Attributes: map[string]schema.Attribute{
															"metadata": schema.SingleNestedAttribute{
																Description:         "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																MarkdownDescription: "ObjectMetaOverride provides the ability to override an object metadata. It's a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta.",
																Attributes: map[string]schema.Attribute{
																	"annotations": schema.MapAttribute{
																		Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																		MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"labels": schema.MapAttribute{
																		Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
																		MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

															"spec": schema.MapAttribute{
																Description:         "Specification of the desired behavior of the pod.",
																MarkdownDescription: "Specification of the desired behavior of the pod.",
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

							"replicas": schema.Int64Attribute{
								Description:         "Number of desired replicas for the ui. Default to 1.",
								MarkdownDescription: "Number of desired replicas for the ui. Default to 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Compute Resources required by the ui. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Compute Resources required by the ui. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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

							"service": schema.SingleNestedAttribute{
								Description:         "Service is an optional service resource configuration for the UI.",
								MarkdownDescription: "Service is an optional service resource configuration for the UI.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects.",
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

							"version": schema.StringAttribute{
								Description:         "Version defines the temporal ui version the instance should run.",
								MarkdownDescription: "Version defines the temporal ui version the instance should run.",
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
						Description:         "Version defines the temporal version the cluster to be deployed. This version impacts the underlying persistence schemas versions.",
						MarkdownDescription: "Version defines the temporal version the cluster to be deployed. This version impacts the underlying persistence schemas versions.",
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

func (r *TemporalIoTemporalClusterV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_temporal_io_temporal_cluster_v1beta1_manifest")

	var model TemporalIoTemporalClusterV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("temporal.io/v1beta1")
	model.Kind = pointer.String("TemporalCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
