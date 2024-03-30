/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentd_fluent_io_v1alpha1

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
	_ datasource.DataSource = &FluentdFluentIoClusterOutputV1Alpha1Manifest{}
)

func NewFluentdFluentIoClusterOutputV1Alpha1Manifest() datasource.DataSource {
	return &FluentdFluentIoClusterOutputV1Alpha1Manifest{}
}

type FluentdFluentIoClusterOutputV1Alpha1Manifest struct{}

type FluentdFluentIoClusterOutputV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Outputs *[]struct {
			Buffer *struct {
				CalcNumRecords              *string `tfsdk:"calc_num_records" json:"calcNumRecords,omitempty"`
				ChunkFormat                 *string `tfsdk:"chunk_format" json:"chunkFormat,omitempty"`
				ChunkLimitRecords           *string `tfsdk:"chunk_limit_records" json:"chunkLimitRecords,omitempty"`
				ChunkLimitSize              *string `tfsdk:"chunk_limit_size" json:"chunkLimitSize,omitempty"`
				Compress                    *string `tfsdk:"compress" json:"compress,omitempty"`
				DelayedCommitTimeout        *string `tfsdk:"delayed_commit_timeout" json:"delayedCommitTimeout,omitempty"`
				DisableChunkBackup          *bool   `tfsdk:"disable_chunk_backup" json:"disableChunkBackup,omitempty"`
				FlushAtShutdown             *bool   `tfsdk:"flush_at_shutdown" json:"flushAtShutdown,omitempty"`
				FlushInterval               *string `tfsdk:"flush_interval" json:"flushInterval,omitempty"`
				FlushMode                   *string `tfsdk:"flush_mode" json:"flushMode,omitempty"`
				FlushThreadCount            *string `tfsdk:"flush_thread_count" json:"flushThreadCount,omitempty"`
				Id                          *string `tfsdk:"id" json:"id,omitempty"`
				Localtime                   *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
				LogLevel                    *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				OverflowAction              *string `tfsdk:"overflow_action" json:"overflowAction,omitempty"`
				Path                        *string `tfsdk:"path" json:"path,omitempty"`
				PathSuffix                  *string `tfsdk:"path_suffix" json:"pathSuffix,omitempty"`
				QueueLimitLength            *string `tfsdk:"queue_limit_length" json:"queueLimitLength,omitempty"`
				QueuedChunksLimitSize       *int64  `tfsdk:"queued_chunks_limit_size" json:"queuedChunksLimitSize,omitempty"`
				RetryExponentialBackoffBase *string `tfsdk:"retry_exponential_backoff_base" json:"retryExponentialBackoffBase,omitempty"`
				RetryForever                *bool   `tfsdk:"retry_forever" json:"retryForever,omitempty"`
				RetryMaxInterval            *string `tfsdk:"retry_max_interval" json:"retryMaxInterval,omitempty"`
				RetryMaxTimes               *int64  `tfsdk:"retry_max_times" json:"retryMaxTimes,omitempty"`
				RetryRandomize              *bool   `tfsdk:"retry_randomize" json:"retryRandomize,omitempty"`
				RetrySecondaryThreshold     *string `tfsdk:"retry_secondary_threshold" json:"retrySecondaryThreshold,omitempty"`
				RetryTimeout                *string `tfsdk:"retry_timeout" json:"retryTimeout,omitempty"`
				RetryType                   *string `tfsdk:"retry_type" json:"retryType,omitempty"`
				RetryWait                   *string `tfsdk:"retry_wait" json:"retryWait,omitempty"`
				Tag                         *string `tfsdk:"tag" json:"tag,omitempty"`
				TimeFormat                  *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
				TimeFormatFallbacks         *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
				TimeType                    *string `tfsdk:"time_type" json:"timeType,omitempty"`
				Timekey                     *string `tfsdk:"timekey" json:"timekey,omitempty"`
				TimekeyWait                 *string `tfsdk:"timekey_wait" json:"timekeyWait,omitempty"`
				Timezone                    *string `tfsdk:"timezone" json:"timezone,omitempty"`
				TotalLimitSize              *string `tfsdk:"total_limit_size" json:"totalLimitSize,omitempty"`
				Type                        *string `tfsdk:"type" json:"type,omitempty"`
				Utc                         *bool   `tfsdk:"utc" json:"utc,omitempty"`
			} `tfsdk:"buffer" json:"buffer,omitempty"`
			CloudWatch *struct {
				AutoCreateStream     *bool `tfsdk:"auto_create_stream" json:"autoCreateStream,omitempty"`
				AwsEcsAuthentication *bool `tfsdk:"aws_ecs_authentication" json:"awsEcsAuthentication,omitempty"`
				AwsKeyId             *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"aws_key_id" json:"awsKeyId,omitempty"`
				AwsSecKey *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"aws_sec_key" json:"awsSecKey,omitempty"`
				AwsStsDurationSeconds         *string `tfsdk:"aws_sts_duration_seconds" json:"awsStsDurationSeconds,omitempty"`
				AwsStsEndpointUrl             *string `tfsdk:"aws_sts_endpoint_url" json:"awsStsEndpointUrl,omitempty"`
				AwsStsExternalId              *string `tfsdk:"aws_sts_external_id" json:"awsStsExternalId,omitempty"`
				AwsStsPolicy                  *string `tfsdk:"aws_sts_policy" json:"awsStsPolicy,omitempty"`
				AwsStsRoleArn                 *string `tfsdk:"aws_sts_role_arn" json:"awsStsRoleArn,omitempty"`
				AwsStsSessionName             *string `tfsdk:"aws_sts_session_name" json:"awsStsSessionName,omitempty"`
				AwsUseSts                     *bool   `tfsdk:"aws_use_sts" json:"awsUseSts,omitempty"`
				Concurrency                   *int64  `tfsdk:"concurrency" json:"concurrency,omitempty"`
				DurationSeconds               *string `tfsdk:"duration_seconds" json:"durationSeconds,omitempty"`
				Endpoint                      *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				HttpProxy                     *string `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
				IncludeTimeKey                *bool   `tfsdk:"include_time_key" json:"includeTimeKey,omitempty"`
				JsonHandler                   *string `tfsdk:"json_handler" json:"jsonHandler,omitempty"`
				Localtime                     *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
				LogGroupAwsTags               *string `tfsdk:"log_group_aws_tags" json:"logGroupAwsTags,omitempty"`
				LogGroupAwsTagsKey            *string `tfsdk:"log_group_aws_tags_key" json:"logGroupAwsTagsKey,omitempty"`
				LogGroupName                  *string `tfsdk:"log_group_name" json:"logGroupName,omitempty"`
				LogGroupNameKey               *string `tfsdk:"log_group_name_key" json:"logGroupNameKey,omitempty"`
				LogRejectedRequest            *string `tfsdk:"log_rejected_request" json:"logRejectedRequest,omitempty"`
				LogStreamName                 *string `tfsdk:"log_stream_name" json:"logStreamName,omitempty"`
				LogStreamNameKey              *string `tfsdk:"log_stream_name_key" json:"logStreamNameKey,omitempty"`
				MaxEventsPerBatch             *string `tfsdk:"max_events_per_batch" json:"maxEventsPerBatch,omitempty"`
				MaxMessageLength              *string `tfsdk:"max_message_length" json:"maxMessageLength,omitempty"`
				MessageKeys                   *string `tfsdk:"message_keys" json:"messageKeys,omitempty"`
				Policy                        *string `tfsdk:"policy" json:"policy,omitempty"`
				PutLogEventsDisableRetryLimit *bool   `tfsdk:"put_log_events_disable_retry_limit" json:"putLogEventsDisableRetryLimit,omitempty"`
				PutLogEventsRetryLimit        *string `tfsdk:"put_log_events_retry_limit" json:"putLogEventsRetryLimit,omitempty"`
				PutLogEventsRetryWait         *string `tfsdk:"put_log_events_retry_wait" json:"putLogEventsRetryWait,omitempty"`
				Region                        *string `tfsdk:"region" json:"region,omitempty"`
				RemoveLogGroupAwsTagsKey      *bool   `tfsdk:"remove_log_group_aws_tags_key" json:"removeLogGroupAwsTagsKey,omitempty"`
				RemoveLogGroupNameKey         *bool   `tfsdk:"remove_log_group_name_key" json:"removeLogGroupNameKey,omitempty"`
				RemoveLogStreamNameKey        *bool   `tfsdk:"remove_log_stream_name_key" json:"removeLogStreamNameKey,omitempty"`
				RemoveRetentionInDaysKey      *bool   `tfsdk:"remove_retention_in_days_key" json:"removeRetentionInDaysKey,omitempty"`
				RetentionInDays               *string `tfsdk:"retention_in_days" json:"retentionInDays,omitempty"`
				RetentionInDaysKey            *string `tfsdk:"retention_in_days_key" json:"retentionInDaysKey,omitempty"`
				RoleArn                       *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
				RoleSessionName               *string `tfsdk:"role_session_name" json:"roleSessionName,omitempty"`
				SslVerifyPeer                 *bool   `tfsdk:"ssl_verify_peer" json:"sslVerifyPeer,omitempty"`
				UseTagAsGroup                 *string `tfsdk:"use_tag_as_group" json:"useTagAsGroup,omitempty"`
				UseTagAsStream                *string `tfsdk:"use_tag_as_stream" json:"useTagAsStream,omitempty"`
				WebIdentityTokenFile          *string `tfsdk:"web_identity_token_file" json:"webIdentityTokenFile,omitempty"`
			} `tfsdk:"cloud_watch" json:"cloudWatch,omitempty"`
			Copy *struct {
				CopyMode *string `tfsdk:"copy_mode" json:"copyMode,omitempty"`
			} `tfsdk:"copy" json:"copy,omitempty"`
			CustomPlugin *struct {
				Config *string `tfsdk:"config" json:"config,omitempty"`
			} `tfsdk:"custom_plugin" json:"customPlugin,omitempty"`
			Datadog *struct {
				ApiKey *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"api_key" json:"apiKey,omitempty"`
				CompressionLevel *int64  `tfsdk:"compression_level" json:"compressionLevel,omitempty"`
				DdHostname       *string `tfsdk:"dd_hostname" json:"ddHostname,omitempty"`
				DdSource         *string `tfsdk:"dd_source" json:"ddSource,omitempty"`
				DdSourcecategory *string `tfsdk:"dd_sourcecategory" json:"ddSourcecategory,omitempty"`
				DdTags           *string `tfsdk:"dd_tags" json:"ddTags,omitempty"`
				Host             *string `tfsdk:"host" json:"host,omitempty"`
				HttpProxy        *string `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
				IncludeTagKey    *bool   `tfsdk:"include_tag_key" json:"includeTagKey,omitempty"`
				MaxBackoff       *int64  `tfsdk:"max_backoff" json:"maxBackoff,omitempty"`
				MaxRetries       *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
				NoSSLValidation  *bool   `tfsdk:"no_ssl_validation" json:"noSSLValidation,omitempty"`
				Port             *int64  `tfsdk:"port" json:"port,omitempty"`
				Service          *string `tfsdk:"service" json:"service,omitempty"`
				SslPort          *int64  `tfsdk:"ssl_port" json:"sslPort,omitempty"`
				TagKey           *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
				TimestampKey     *string `tfsdk:"timestamp_key" json:"timestampKey,omitempty"`
				UseCompression   *bool   `tfsdk:"use_compression" json:"useCompression,omitempty"`
				UseHTTP          *bool   `tfsdk:"use_http" json:"useHTTP,omitempty"`
				UseJson          *bool   `tfsdk:"use_json" json:"useJson,omitempty"`
				UseSSL           *bool   `tfsdk:"use_ssl" json:"useSSL,omitempty"`
			} `tfsdk:"datadog" json:"datadog,omitempty"`
			Elasticsearch *struct {
				CaFile            *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				ClientCert        *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
				ClientKey         *string `tfsdk:"client_key" json:"clientKey,omitempty"`
				ClientKeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"client_key_password" json:"clientKeyPassword,omitempty"`
				Host           *string `tfsdk:"host" json:"host,omitempty"`
				Hosts          *string `tfsdk:"hosts" json:"hosts,omitempty"`
				IndexName      *string `tfsdk:"index_name" json:"indexName,omitempty"`
				LogstashFormat *bool   `tfsdk:"logstash_format" json:"logstashFormat,omitempty"`
				LogstashPrefix *string `tfsdk:"logstash_prefix" json:"logstashPrefix,omitempty"`
				Password       *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				Scheme    *string `tfsdk:"scheme" json:"scheme,omitempty"`
				SslVerify *bool   `tfsdk:"ssl_verify" json:"sslVerify,omitempty"`
				User      *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
			Format *struct {
				Delimiter           *string `tfsdk:"delimiter" json:"delimiter,omitempty"`
				Id                  *string `tfsdk:"id" json:"id,omitempty"`
				Localtime           *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
				LogLevel            *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				Newline             *string `tfsdk:"newline" json:"newline,omitempty"`
				OutputTag           *bool   `tfsdk:"output_tag" json:"outputTag,omitempty"`
				OutputTime          *bool   `tfsdk:"output_time" json:"outputTime,omitempty"`
				TimeFormat          *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
				TimeFormatFallbacks *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
				TimeType            *string `tfsdk:"time_type" json:"timeType,omitempty"`
				Timezone            *string `tfsdk:"timezone" json:"timezone,omitempty"`
				Type                *string `tfsdk:"type" json:"type,omitempty"`
				Utc                 *bool   `tfsdk:"utc" json:"utc,omitempty"`
			} `tfsdk:"format" json:"format,omitempty"`
			Forward *struct {
				AckResponseTimeout           *string `tfsdk:"ack_response_timeout" json:"ackResponseTimeout,omitempty"`
				ConnectTimeout               *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				DnsRoundRobin                *bool   `tfsdk:"dns_round_robin" json:"dnsRoundRobin,omitempty"`
				ExpireDnsCache               *string `tfsdk:"expire_dns_cache" json:"expireDnsCache,omitempty"`
				HardTimeout                  *string `tfsdk:"hard_timeout" json:"hardTimeout,omitempty"`
				HeartbeatInterval            *string `tfsdk:"heartbeat_interval" json:"heartbeatInterval,omitempty"`
				HeartbeatType                *string `tfsdk:"heartbeat_type" json:"heartbeatType,omitempty"`
				IgnoreNetworkErrorsAtStartup *bool   `tfsdk:"ignore_network_errors_at_startup" json:"ignoreNetworkErrorsAtStartup,omitempty"`
				Keepalive                    *bool   `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveTimeout             *string `tfsdk:"keepalive_timeout" json:"keepaliveTimeout,omitempty"`
				PhiFailureDetector           *bool   `tfsdk:"phi_failure_detector" json:"phiFailureDetector,omitempty"`
				PhiThreshold                 *int64  `tfsdk:"phi_threshold" json:"phiThreshold,omitempty"`
				RecoverWait                  *string `tfsdk:"recover_wait" json:"recoverWait,omitempty"`
				RequireAckResponse           *bool   `tfsdk:"require_ack_response" json:"requireAckResponse,omitempty"`
				Security                     *struct {
					AllowAnonymousSource *string `tfsdk:"allow_anonymous_source" json:"allowAnonymousSource,omitempty"`
					SelfHostname         *string `tfsdk:"self_hostname" json:"selfHostname,omitempty"`
					SharedKey            *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
					User                 *struct {
						Password *struct {
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Username *struct {
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"user" json:"user,omitempty"`
					UserAuth *string `tfsdk:"user_auth" json:"userAuth,omitempty"`
				} `tfsdk:"security" json:"security,omitempty"`
				SendTimeout *string `tfsdk:"send_timeout" json:"sendTimeout,omitempty"`
				Servers     *[]struct {
					Host     *string `tfsdk:"host" json:"host,omitempty"`
					Id       *string `tfsdk:"id" json:"id,omitempty"`
					LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Password *struct {
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"password" json:"password,omitempty"`
					Port      *string `tfsdk:"port" json:"port,omitempty"`
					SharedKey *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
					Standby   *string `tfsdk:"standby" json:"standby,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
					Username  *struct {
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"username" json:"username,omitempty"`
					Weight *string `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"servers" json:"servers,omitempty"`
				ServiceDiscovery *struct {
					ConfEncoding  *string `tfsdk:"conf_encoding" json:"confEncoding,omitempty"`
					DnsLookup     *string `tfsdk:"dns_lookup" json:"dnsLookup,omitempty"`
					DnsServerHost *string `tfsdk:"dns_server_host" json:"dnsServerHost,omitempty"`
					Hostname      *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Id            *string `tfsdk:"id" json:"id,omitempty"`
					Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
					LogLevel      *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Path          *string `tfsdk:"path" json:"path,omitempty"`
					Proto         *string `tfsdk:"proto" json:"proto,omitempty"`
					Server        *struct {
						Host     *string `tfsdk:"host" json:"host,omitempty"`
						Id       *string `tfsdk:"id" json:"id,omitempty"`
						LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Password *struct {
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Port      *string `tfsdk:"port" json:"port,omitempty"`
						SharedKey *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
						Standby   *string `tfsdk:"standby" json:"standby,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
						Username  *struct {
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
						Weight *string `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"server" json:"server,omitempty"`
					Service *string `tfsdk:"service" json:"service,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"service_discovery" json:"serviceDiscovery,omitempty"`
				TlsAllowSelfSignedCert        *bool   `tfsdk:"tls_allow_self_signed_cert" json:"tlsAllowSelfSignedCert,omitempty"`
				TlsCertLogicalStoreName       *string `tfsdk:"tls_cert_logical_store_name" json:"tlsCertLogicalStoreName,omitempty"`
				TlsCertPath                   *string `tfsdk:"tls_cert_path" json:"tlsCertPath,omitempty"`
				TlsCertThumbprint             *string `tfsdk:"tls_cert_thumbprint" json:"tlsCertThumbprint,omitempty"`
				TlsCertUseEnterpriseStore     *bool   `tfsdk:"tls_cert_use_enterprise_store" json:"tlsCertUseEnterpriseStore,omitempty"`
				TlsCiphers                    *string `tfsdk:"tls_ciphers" json:"tlsCiphers,omitempty"`
				TlsClientCertPath             *string `tfsdk:"tls_client_cert_path" json:"tlsClientCertPath,omitempty"`
				TlsClientPrivateKeyPassphrase *string `tfsdk:"tls_client_private_key_passphrase" json:"tlsClientPrivateKeyPassphrase,omitempty"`
				TlsClientPrivateKeyPath       *string `tfsdk:"tls_client_private_key_path" json:"tlsClientPrivateKeyPath,omitempty"`
				TlsInsecureMode               *bool   `tfsdk:"tls_insecure_mode" json:"tlsInsecureMode,omitempty"`
				TlsVerifyHostname             *bool   `tfsdk:"tls_verify_hostname" json:"tlsVerifyHostname,omitempty"`
				TlsVersion                    *string `tfsdk:"tls_version" json:"tlsVersion,omitempty"`
				VerifyConnectionAtStartup     *bool   `tfsdk:"verify_connection_at_startup" json:"verifyConnectionAtStartup,omitempty"`
			} `tfsdk:"forward" json:"forward,omitempty"`
			Http *struct {
				Auth *struct {
					Auth     *string `tfsdk:"auth" json:"auth,omitempty"`
					Password *struct {
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"password" json:"password,omitempty"`
					Username *struct {
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				ContentType                  *string `tfsdk:"content_type" json:"contentType,omitempty"`
				Endpoint                     *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				ErrorResponseAsUnrecoverable *bool   `tfsdk:"error_response_as_unrecoverable" json:"errorResponseAsUnrecoverable,omitempty"`
				Headers                      *string `tfsdk:"headers" json:"headers,omitempty"`
				HeadersFromPlaceholders      *string `tfsdk:"headers_from_placeholders" json:"headersFromPlaceholders,omitempty"`
				HttpMethod                   *string `tfsdk:"http_method" json:"httpMethod,omitempty"`
				JsonArray                    *bool   `tfsdk:"json_array" json:"jsonArray,omitempty"`
				OpenTimeout                  *int64  `tfsdk:"open_timeout" json:"openTimeout,omitempty"`
				Proxy                        *string `tfsdk:"proxy" json:"proxy,omitempty"`
				ReadTimeout                  *int64  `tfsdk:"read_timeout" json:"readTimeout,omitempty"`
				RetryableResponseCodes       *string `tfsdk:"retryable_response_codes" json:"retryableResponseCodes,omitempty"`
				SslTimeout                   *int64  `tfsdk:"ssl_timeout" json:"sslTimeout,omitempty"`
				TlsCaCertPath                *string `tfsdk:"tls_ca_cert_path" json:"tlsCaCertPath,omitempty"`
				TlsCiphers                   *string `tfsdk:"tls_ciphers" json:"tlsCiphers,omitempty"`
				TlsClientCertPath            *string `tfsdk:"tls_client_cert_path" json:"tlsClientCertPath,omitempty"`
				TlsPrivateKeyPassphrase      *string `tfsdk:"tls_private_key_passphrase" json:"tlsPrivateKeyPassphrase,omitempty"`
				TlsPrivateKeyPath            *string `tfsdk:"tls_private_key_path" json:"tlsPrivateKeyPath,omitempty"`
				TlsVerifyMode                *string `tfsdk:"tls_verify_mode" json:"tlsVerifyMode,omitempty"`
				TlsVersion                   *string `tfsdk:"tls_version" json:"tlsVersion,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Inject *struct {
				Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
				HostnameKey *string `tfsdk:"hostname_key" json:"hostnameKey,omitempty"`
				Inline      *struct {
					Localtime           *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
					TimeFormat          *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
					TimeFormatFallbacks *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
					TimeType            *string `tfsdk:"time_type" json:"timeType,omitempty"`
					Timezone            *string `tfsdk:"timezone" json:"timezone,omitempty"`
					Utc                 *bool   `tfsdk:"utc" json:"utc,omitempty"`
				} `tfsdk:"inline" json:"inline,omitempty"`
				TagKey      *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
				TimeKey     *string `tfsdk:"time_key" json:"timeKey,omitempty"`
				WorkerIdKey *string `tfsdk:"worker_id_key" json:"workerIdKey,omitempty"`
			} `tfsdk:"inject" json:"inject,omitempty"`
			Kafka *struct {
				Brokers          *string `tfsdk:"brokers" json:"brokers,omitempty"`
				CompressionCodec *string `tfsdk:"compression_codec" json:"compressionCodec,omitempty"`
				DefaultTopic     *string `tfsdk:"default_topic" json:"defaultTopic,omitempty"`
				RequiredAcks     *int64  `tfsdk:"required_acks" json:"requiredAcks,omitempty"`
				TopicKey         *string `tfsdk:"topic_key" json:"topicKey,omitempty"`
				UseEventTime     *bool   `tfsdk:"use_event_time" json:"useEventTime,omitempty"`
			} `tfsdk:"kafka" json:"kafka,omitempty"`
			LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Loki     *struct {
				DropSingleKey           *bool `tfsdk:"drop_single_key" json:"dropSingleKey,omitempty"`
				ExtractKubernetesLabels *bool `tfsdk:"extract_kubernetes_labels" json:"extractKubernetesLabels,omitempty"`
				HttpPassword            *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"http_password" json:"httpPassword,omitempty"`
				HttpUser *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"http_user" json:"httpUser,omitempty"`
				IncludeThreadLabel *bool     `tfsdk:"include_thread_label" json:"includeThreadLabel,omitempty"`
				Insecure           *bool     `tfsdk:"insecure" json:"insecure,omitempty"`
				LabelKeys          *[]string `tfsdk:"label_keys" json:"labelKeys,omitempty"`
				Labels             *[]string `tfsdk:"labels" json:"labels,omitempty"`
				LineFormat         *string   `tfsdk:"line_format" json:"lineFormat,omitempty"`
				RemoveKeys         *[]string `tfsdk:"remove_keys" json:"removeKeys,omitempty"`
				TenantID           *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				TlsCaCertFile     *string `tfsdk:"tls_ca_cert_file" json:"tlsCaCertFile,omitempty"`
				TlsClientCertFile *string `tfsdk:"tls_client_cert_file" json:"tlsClientCertFile,omitempty"`
				TlsPrivateKeyFile *string `tfsdk:"tls_private_key_file" json:"tlsPrivateKeyFile,omitempty"`
				Url               *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"loki" json:"loki,omitempty"`
			Opensearch *struct {
				Host           *string `tfsdk:"host" json:"host,omitempty"`
				Hosts          *string `tfsdk:"hosts" json:"hosts,omitempty"`
				IndexName      *string `tfsdk:"index_name" json:"indexName,omitempty"`
				LogstashFormat *bool   `tfsdk:"logstash_format" json:"logstashFormat,omitempty"`
				LogstashPrefix *string `tfsdk:"logstash_prefix" json:"logstashPrefix,omitempty"`
				Password       *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *int64  `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
				User   *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"opensearch" json:"opensearch,omitempty"`
			S3 *struct {
				AwsKeyId                *string `tfsdk:"aws_key_id" json:"awsKeyId,omitempty"`
				AwsSecKey               *string `tfsdk:"aws_sec_key" json:"awsSecKey,omitempty"`
				ForcePathStyle          *bool   `tfsdk:"force_path_style" json:"forcePathStyle,omitempty"`
				Path                    *string `tfsdk:"path" json:"path,omitempty"`
				ProxyUri                *string `tfsdk:"proxy_uri" json:"proxyUri,omitempty"`
				S3Bucket                *string `tfsdk:"s3_bucket" json:"s3Bucket,omitempty"`
				S3Endpoint              *string `tfsdk:"s3_endpoint" json:"s3Endpoint,omitempty"`
				S3ObjectKeyFormat       *string `tfsdk:"s3_object_key_format" json:"s3ObjectKeyFormat,omitempty"`
				S3Region                *string `tfsdk:"s3_region" json:"s3Region,omitempty"`
				SseCustomerAlgorithm    *string `tfsdk:"sse_customer_algorithm" json:"sseCustomerAlgorithm,omitempty"`
				SseCustomerKey          *string `tfsdk:"sse_customer_key" json:"sseCustomerKey,omitempty"`
				SseCustomerKeyMd5       *string `tfsdk:"sse_customer_key_md5" json:"sseCustomerKeyMd5,omitempty"`
				SsekmsKeyId             *string `tfsdk:"ssekms_key_id" json:"ssekmsKeyId,omitempty"`
				SslVerifyPeer           *bool   `tfsdk:"ssl_verify_peer" json:"sslVerifyPeer,omitempty"`
				StoreAs                 *string `tfsdk:"store_as" json:"storeAs,omitempty"`
				TimeSliceFormat         *string `tfsdk:"time_slice_format" json:"timeSliceFormat,omitempty"`
				UseServerSideEncryption *string `tfsdk:"use_server_side_encryption" json:"useServerSideEncryption,omitempty"`
			} `tfsdk:"s3" json:"s3,omitempty"`
			Stdout *map[string]string `tfsdk:"stdout" json:"stdout,omitempty"`
			Tag    *string            `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"outputs" json:"outputs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentdFluentIoClusterOutputV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentd_fluent_io_cluster_output_v1alpha1_manifest"
}

func (r *FluentdFluentIoClusterOutputV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterOutput is the Schema for the clusteroutputs API",
		MarkdownDescription: "ClusterOutput is the Schema for the clusteroutputs API",
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
				Description:         "ClusterOutputSpec defines the desired state of ClusterOutput",
				MarkdownDescription: "ClusterOutputSpec defines the desired state of ClusterOutput",
				Attributes: map[string]schema.Attribute{
					"outputs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"buffer": schema.SingleNestedAttribute{
									Description:         "buffer section",
									MarkdownDescription: "buffer section",
									Attributes: map[string]schema.Attribute{
										"calc_num_records": schema.StringAttribute{
											Description:         "Calculates the number of records, chunk size, during chunk resume.",
											MarkdownDescription: "Calculates the number of records, chunk size, during chunk resume.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"chunk_format": schema.StringAttribute{
											Description:         "ChunkFormat specifies the chunk format for calc_num_records.",
											MarkdownDescription: "ChunkFormat specifies the chunk format for calc_num_records.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("msgpack", "text", "auto"),
											},
										},

										"chunk_limit_records": schema.StringAttribute{
											Description:         "The max number of events that each chunks can store in it.",
											MarkdownDescription: "The max number of events that each chunks can store in it.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(KB|MB|GB|TB)$`), ""),
											},
										},

										"chunk_limit_size": schema.StringAttribute{
											Description:         "Buffer parameters The max size of each chunks: events will be written into chunks until the size of chunks become this size Default: 8MB (memory) / 256MB (file)",
											MarkdownDescription: "Buffer parameters The max size of each chunks: events will be written into chunks until the size of chunks become this size Default: 8MB (memory) / 256MB (file)",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(KB|MB|GB|TB)$`), ""),
											},
										},

										"compress": schema.StringAttribute{
											Description:         "Fluentd will decompress these compressed chunks automatically before passing them to the output plugin If gzip is set, Fluentd compresses data records before writing to buffer chunks. Default:text.",
											MarkdownDescription: "Fluentd will decompress these compressed chunks automatically before passing them to the output plugin If gzip is set, Fluentd compresses data records before writing to buffer chunks. Default:text.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("text", "gzip"),
											},
										},

										"delayed_commit_timeout": schema.StringAttribute{
											Description:         "The timeout (seconds) until output plugin decides if the async write operation has failed. Default is 60s",
											MarkdownDescription: "The timeout (seconds) until output plugin decides if the async write operation has failed. Default is 60s",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"disable_chunk_backup": schema.BoolAttribute{
											Description:         "Instead of storing unrecoverable chunks in the backup directory, just discard them. This option is new in Fluentd v1.2.6.",
											MarkdownDescription: "Instead of storing unrecoverable chunks in the backup directory, just discard them. This option is new in Fluentd v1.2.6.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"flush_at_shutdown": schema.BoolAttribute{
											Description:         "Flush parameters This specifies whether to flush/write all buffer chunks on shutdown or not.",
											MarkdownDescription: "Flush parameters This specifies whether to flush/write all buffer chunks on shutdown or not.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"flush_interval": schema.StringAttribute{
											Description:         "FlushInterval defines the flush interval",
											MarkdownDescription: "FlushInterval defines the flush interval",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"flush_mode": schema.StringAttribute{
											Description:         "FlushMode defines the flush mode: lazy: flushes/writes chunks once per timekey interval: flushes/writes chunks per specified time via flush_interval immediate: flushes/writes chunks immediately after events are appended into chunks default: equals to lazy if time is specified as chunk key, interval otherwise",
											MarkdownDescription: "FlushMode defines the flush mode: lazy: flushes/writes chunks once per timekey interval: flushes/writes chunks per specified time via flush_interval immediate: flushes/writes chunks immediately after events are appended into chunks default: equals to lazy if time is specified as chunk key, interval otherwise",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("default", "lazy", "interval", "immediate"),
											},
										},

										"flush_thread_count": schema.StringAttribute{
											Description:         "The number of threads to flush/write chunks in parallel",
											MarkdownDescription: "The number of threads to flush/write chunks in parallel",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+$`), ""),
											},
										},

										"id": schema.StringAttribute{
											Description:         "The @id parameter specifies a unique name for the configuration.",
											MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"localtime": schema.BoolAttribute{
											Description:         "If true, uses local time.",
											MarkdownDescription: "If true, uses local time.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_level": schema.StringAttribute{
											Description:         "The @log_level parameter specifies the plugin-specific logging level",
											MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"overflow_action": schema.StringAttribute{
											Description:         "OverflowAtction defines the output plugin behave when its buffer queue is full. Default: throw_exception",
											MarkdownDescription: "OverflowAtction defines the output plugin behave when its buffer queue is full. Default: throw_exception",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "The path where buffer chunks are stored. This field would make no effect in memory buffer plugin.",
											MarkdownDescription: "The path where buffer chunks are stored. This field would make no effect in memory buffer plugin.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path_suffix": schema.StringAttribute{
											Description:         "Changes the suffix of the buffer file.",
											MarkdownDescription: "Changes the suffix of the buffer file.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"queue_limit_length": schema.StringAttribute{
											Description:         "The queue length limitation of this buffer plugin instance. Default: 0.95",
											MarkdownDescription: "The queue length limitation of this buffer plugin instance. Default: 0.95",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+.?\d+$`), ""),
											},
										},

										"queued_chunks_limit_size": schema.Int64Attribute{
											Description:         "Limit the number of queued chunks. Default: 1 If a smaller flush_interval is set, e.g. 1s, there are lots of small queued chunks in the buffer. With file buffer, it may consume a lot of fd resources when output destination has a problem. This parameter mitigates such situations.",
											MarkdownDescription: "Limit the number of queued chunks. Default: 1 If a smaller flush_interval is set, e.g. 1s, there are lots of small queued chunks in the buffer. With file buffer, it may consume a lot of fd resources when output destination has a problem. This parameter mitigates such situations.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"retry_exponential_backoff_base": schema.StringAttribute{
											Description:         "The base number of exponential backoff for retries.",
											MarkdownDescription: "The base number of exponential backoff for retries.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?$`), ""),
											},
										},

										"retry_forever": schema.BoolAttribute{
											Description:         "If true, plugin will ignore retry_timeout and retry_max_times options and retry flushing forever.",
											MarkdownDescription: "If true, plugin will ignore retry_timeout and retry_max_times options and retry flushing forever.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_max_interval": schema.StringAttribute{
											Description:         "The maximum interval (seconds) for exponential backoff between retries while failing",
											MarkdownDescription: "The maximum interval (seconds) for exponential backoff between retries while failing",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"retry_max_times": schema.Int64Attribute{
											Description:         "The maximum number of times to retry to flush the failed chunks. Default: none",
											MarkdownDescription: "The maximum number of times to retry to flush the failed chunks. Default: none",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_randomize": schema.BoolAttribute{
											Description:         "If true, the output plugin will retry after randomized interval not to do burst retries",
											MarkdownDescription: "If true, the output plugin will retry after randomized interval not to do burst retries",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_secondary_threshold": schema.StringAttribute{
											Description:         "The ratio of retry_timeout to switch to use the secondary while failing.",
											MarkdownDescription: "The ratio of retry_timeout to switch to use the secondary while failing.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+.?\d+$`), ""),
											},
										},

										"retry_timeout": schema.StringAttribute{
											Description:         "Retry parameters The maximum time (seconds) to retry to flush again the failed chunks, until the plugin discards the buffer chunks",
											MarkdownDescription: "Retry parameters The maximum time (seconds) to retry to flush again the failed chunks, until the plugin discards the buffer chunks",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"retry_type": schema.StringAttribute{
											Description:         "Output plugin will retry periodically with fixed intervals.",
											MarkdownDescription: "Output plugin will retry periodically with fixed intervals.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_wait": schema.StringAttribute{
											Description:         "Wait in seconds before the next retry to flush or constant factor of exponential backoff",
											MarkdownDescription: "Wait in seconds before the next retry to flush or constant factor of exponential backoff",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"tag": schema.StringAttribute{
											Description:         "The output plugins group events into chunks. Chunk keys, specified as the argument of <buffer> section, control how to group events into chunks. If tag is empty, which means blank Chunk Keys. Tag also supports Nested Field, combination of Chunk Keys, placeholders, etc. See https://docs.fluentd.org/configuration/buffer-section.",
											MarkdownDescription: "The output plugins group events into chunks. Chunk keys, specified as the argument of <buffer> section, control how to group events into chunks. If tag is empty, which means blank Chunk Keys. Tag also supports Nested Field, combination of Chunk Keys, placeholders, etc. See https://docs.fluentd.org/configuration/buffer-section.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_format": schema.StringAttribute{
											Description:         "Process value according to the specified format. This is available only when time_type is string",
											MarkdownDescription: "Process value according to the specified format. This is available only when time_type is string",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_format_fallbacks": schema.StringAttribute{
											Description:         "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
											MarkdownDescription: "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_type": schema.StringAttribute{
											Description:         "parses/formats value according to this type, default is string",
											MarkdownDescription: "parses/formats value according to this type, default is string",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("float", "unixtime", "string", "mixed"),
											},
										},

										"timekey": schema.StringAttribute{
											Description:         "Output plugin will flush chunks per specified time (enabled when time is specified in chunk keys)",
											MarkdownDescription: "Output plugin will flush chunks per specified time (enabled when time is specified in chunk keys)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timekey_wait": schema.StringAttribute{
											Description:         "Output plugin will write chunks after timekey_wait seconds later after timekey expiration",
											MarkdownDescription: "Output plugin will write chunks after timekey_wait seconds later after timekey expiration",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timezone": schema.StringAttribute{
											Description:         "Uses the specified timezone.",
											MarkdownDescription: "Uses the specified timezone.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"total_limit_size": schema.StringAttribute{
											Description:         "The size limitation of this buffer plugin instance Default: 512MB (memory) / 64GB (file)",
											MarkdownDescription: "The size limitation of this buffer plugin instance Default: 512MB (memory) / 64GB (file)",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(KB|MB|GB|TB)$`), ""),
											},
										},

										"type": schema.StringAttribute{
											Description:         "The @type parameter specifies the type of the plugin.",
											MarkdownDescription: "The @type parameter specifies the type of the plugin.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("file", "memory", "file_single"),
											},
										},

										"utc": schema.BoolAttribute{
											Description:         "If true, uses UTC.",
											MarkdownDescription: "If true, uses UTC.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cloud_watch": schema.SingleNestedAttribute{
									Description:         "out_cloudwatch plugin",
									MarkdownDescription: "out_cloudwatch plugin",
									Attributes: map[string]schema.Attribute{
										"auto_create_stream": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_ecs_authentication": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_key_id": schema.SingleNestedAttribute{
											Description:         "Secret defines the key of a value.",
											MarkdownDescription: "Secret defines the key of a value.",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"aws_sec_key": schema.SingleNestedAttribute{
											Description:         "Secret defines the key of a value.",
											MarkdownDescription: "Secret defines the key of a value.",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"aws_sts_duration_seconds": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_sts_endpoint_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_sts_external_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_sts_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_sts_role_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_sts_session_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_use_sts": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"concurrency": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"duration_seconds": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"endpoint": schema.StringAttribute{
											Description:         "Specify an AWS endpoint to send data to.",
											MarkdownDescription: "Specify an AWS endpoint to send data to.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http_proxy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_time_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_handler": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"localtime": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_group_aws_tags": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_group_aws_tags_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_group_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_group_name_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_rejected_request": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_stream_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_stream_name_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_events_per_batch": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_message_length": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"message_keys": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"put_log_events_disable_retry_limit": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"put_log_events_retry_limit": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"put_log_events_retry_wait": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"region": schema.StringAttribute{
											Description:         "The AWS region.",
											MarkdownDescription: "The AWS region.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_log_group_aws_tags_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_log_group_name_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_log_stream_name_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_retention_in_days_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retention_in_days": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retention_in_days_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"role_arn": schema.StringAttribute{
											Description:         "ARN of an IAM role to assume (for cross account access).",
											MarkdownDescription: "ARN of an IAM role to assume (for cross account access).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"role_session_name": schema.StringAttribute{
											Description:         "Role Session name",
											MarkdownDescription: "Role Session name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ssl_verify_peer": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_tag_as_group": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_tag_as_stream": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"web_identity_token_file": schema.StringAttribute{
											Description:         "Web identity token file",
											MarkdownDescription: "Web identity token file",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"copy": schema.SingleNestedAttribute{
									Description:         "copy plugin",
									MarkdownDescription: "copy plugin",
									Attributes: map[string]schema.Attribute{
										"copy_mode": schema.StringAttribute{
											Description:         "CopyMode defines how to pass the events to <store> plugins.",
											MarkdownDescription: "CopyMode defines how to pass the events to <store> plugins.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("no_copy", "shallow", "deep", "marshal"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"custom_plugin": schema.SingleNestedAttribute{
									Description:         "Custom plugin type",
									MarkdownDescription: "Custom plugin type",
									Attributes: map[string]schema.Attribute{
										"config": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"datadog": schema.SingleNestedAttribute{
									Description:         "datadog plugin",
									MarkdownDescription: "datadog plugin",
									Attributes: map[string]schema.Attribute{
										"api_key": schema.SingleNestedAttribute{
											Description:         "This parameter is required in order to authenticate your fluent agent.",
											MarkdownDescription: "This parameter is required in order to authenticate your fluent agent.",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"compression_level": schema.Int64Attribute{
											Description:         "Set the log compression level for HTTP (1 to 9, 9 being the best ratio)",
											MarkdownDescription: "Set the log compression level for HTTP (1 to 9, 9 being the best ratio)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dd_hostname": schema.StringAttribute{
											Description:         "Used by Datadog to identify the host submitting the logs.",
											MarkdownDescription: "Used by Datadog to identify the host submitting the logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dd_source": schema.StringAttribute{
											Description:         "This tells Datadog what integration it is",
											MarkdownDescription: "This tells Datadog what integration it is",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dd_sourcecategory": schema.StringAttribute{
											Description:         "Multiple value attribute. Can be used to refine the source attribute",
											MarkdownDescription: "Multiple value attribute. Can be used to refine the source attribute",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dd_tags": schema.StringAttribute{
											Description:         "Custom tags with the following format 'key1:value1, key2:value2'",
											MarkdownDescription: "Custom tags with the following format 'key1:value1, key2:value2'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"host": schema.StringAttribute{
											Description:         "Proxy endpoint when logs are not directly forwarded to Datadog",
											MarkdownDescription: "Proxy endpoint when logs are not directly forwarded to Datadog",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http_proxy": schema.StringAttribute{
											Description:         "HTTP proxy, only takes effect if HTTP forwarding is enabled (use_http). Defaults to HTTP_PROXY/http_proxy env vars.",
											MarkdownDescription: "HTTP proxy, only takes effect if HTTP forwarding is enabled (use_http). Defaults to HTTP_PROXY/http_proxy env vars.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_tag_key": schema.BoolAttribute{
											Description:         "Automatically include the Fluentd tag in the record.",
											MarkdownDescription: "Automatically include the Fluentd tag in the record.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_backoff": schema.Int64Attribute{
											Description:         "The maximum time waited between each retry in seconds",
											MarkdownDescription: "The maximum time waited between each retry in seconds",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_retries": schema.Int64Attribute{
											Description:         "The number of retries before the output plugin stops. Set to -1 for unlimited retries",
											MarkdownDescription: "The number of retries before the output plugin stops. Set to -1 for unlimited retries",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"no_ssl_validation": schema.BoolAttribute{
											Description:         "Disable SSL validation (useful for proxy forwarding)",
											MarkdownDescription: "Disable SSL validation (useful for proxy forwarding)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Proxy port when logs are not directly forwarded to Datadog and ssl is not used",
											MarkdownDescription: "Proxy port when logs are not directly forwarded to Datadog and ssl is not used",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"service": schema.StringAttribute{
											Description:         "Used by Datadog to correlate between logs, traces and metrics.",
											MarkdownDescription: "Used by Datadog to correlate between logs, traces and metrics.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ssl_port": schema.Int64Attribute{
											Description:         "Port used to send logs over a SSL encrypted connection to Datadog. If use_http is disabled, use 10516 for the US region and 443 for the EU region.",
											MarkdownDescription: "Port used to send logs over a SSL encrypted connection to Datadog. If use_http is disabled, use 10516 for the US region and 443 for the EU region.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"tag_key": schema.StringAttribute{
											Description:         "Where to store the Fluentd tag.",
											MarkdownDescription: "Where to store the Fluentd tag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timestamp_key": schema.StringAttribute{
											Description:         "Name of the attribute which will contain timestamp of the log event. If nil, timestamp attribute is not added.",
											MarkdownDescription: "Name of the attribute which will contain timestamp of the log event. If nil, timestamp attribute is not added.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_compression": schema.BoolAttribute{
											Description:         "Enable log compression for HTTP",
											MarkdownDescription: "Enable log compression for HTTP",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_http": schema.BoolAttribute{
											Description:         "Enable HTTP forwarding. If you disable it, make sure to change the port to 10514 or ssl_port to 10516",
											MarkdownDescription: "Enable HTTP forwarding. If you disable it, make sure to change the port to 10514 or ssl_port to 10516",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_json": schema.BoolAttribute{
											Description:         "Event format, if true, the event is sent in json format. Othwerwise, in plain text.",
											MarkdownDescription: "Event format, if true, the event is sent in json format. Othwerwise, in plain text.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_ssl": schema.BoolAttribute{
											Description:         "If true, the agent initializes a secure connection to Datadog. In clear TCP otherwise.",
											MarkdownDescription: "If true, the agent initializes a secure connection to Datadog. In clear TCP otherwise.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"elasticsearch": schema.SingleNestedAttribute{
									Description:         "out_es plugin",
									MarkdownDescription: "out_es plugin",
									Attributes: map[string]schema.Attribute{
										"ca_file": schema.StringAttribute{
											Description:         "Optional, Absolute path to CA certificate file",
											MarkdownDescription: "Optional, Absolute path to CA certificate file",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_cert": schema.StringAttribute{
											Description:         "Optional, Absolute path to client Certificate file",
											MarkdownDescription: "Optional, Absolute path to client Certificate file",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_key": schema.StringAttribute{
											Description:         "Optional, Absolute path to client private Key file",
											MarkdownDescription: "Optional, Absolute path to client private Key file",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_key_password": schema.SingleNestedAttribute{
											Description:         "Optional, password for ClientKey file",
											MarkdownDescription: "Optional, password for ClientKey file",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"host": schema.StringAttribute{
											Description:         "The hostname of your Elasticsearch node (default: localhost).",
											MarkdownDescription: "The hostname of your Elasticsearch node (default: localhost).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hosts": schema.StringAttribute{
											Description:         "Hosts defines a list of hosts if you want to connect to more than one Elasticsearch nodes",
											MarkdownDescription: "Hosts defines a list of hosts if you want to connect to more than one Elasticsearch nodes",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"index_name": schema.StringAttribute{
											Description:         "IndexName defines the placeholder syntax of Fluentd plugin API. See https://docs.fluentd.org/configuration/buffer-section.",
											MarkdownDescription: "IndexName defines the placeholder syntax of Fluentd plugin API. See https://docs.fluentd.org/configuration/buffer-section.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logstash_format": schema.BoolAttribute{
											Description:         "If true, Fluentd uses the conventional index name format logstash-%Y.%m.%d (default: false). This option supersedes the index_name option.",
											MarkdownDescription: "If true, Fluentd uses the conventional index name format logstash-%Y.%m.%d (default: false). This option supersedes the index_name option.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logstash_prefix": schema.StringAttribute{
											Description:         "LogstashPrefix defines the logstash prefix index name to write events when logstash_format is true (default: logstash).",
											MarkdownDescription: "LogstashPrefix defines the logstash prefix index name to write events when logstash_format is true (default: logstash).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"password": schema.SingleNestedAttribute{
											Description:         "Optional, The login credentials to connect to Elasticsearch",
											MarkdownDescription: "Optional, The login credentials to connect to Elasticsearch",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "Path defines the REST API endpoint of Elasticsearch to post write requests (default: nil).",
											MarkdownDescription: "Path defines the REST API endpoint of Elasticsearch to post write requests (default: nil).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port number of your Elasticsearch node (default: 9200).",
											MarkdownDescription: "The port number of your Elasticsearch node (default: 9200).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"scheme": schema.StringAttribute{
											Description:         "Specify https if your Elasticsearch endpoint supports SSL (default: http).",
											MarkdownDescription: "Specify https if your Elasticsearch endpoint supports SSL (default: http).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ssl_verify": schema.BoolAttribute{
											Description:         "Optional, Force certificate validation",
											MarkdownDescription: "Optional, Force certificate validation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user": schema.SingleNestedAttribute{
											Description:         "Optional, The login credentials to connect to Elasticsearch",
											MarkdownDescription: "Optional, The login credentials to connect to Elasticsearch",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"format": schema.SingleNestedAttribute{
									Description:         "format section",
									MarkdownDescription: "format section",
									Attributes: map[string]schema.Attribute{
										"delimiter": schema.StringAttribute{
											Description:         "Delimiter for each field.",
											MarkdownDescription: "Delimiter for each field.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "The @id parameter specifies a unique name for the configuration.",
											MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"localtime": schema.BoolAttribute{
											Description:         "If true, uses local time.",
											MarkdownDescription: "If true, uses local time.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_level": schema.StringAttribute{
											Description:         "The @log_level parameter specifies the plugin-specific logging level",
											MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"newline": schema.StringAttribute{
											Description:         "Specify newline characters.",
											MarkdownDescription: "Specify newline characters.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("lf", "crlf"),
											},
										},

										"output_tag": schema.BoolAttribute{
											Description:         "Output tag field if true.",
											MarkdownDescription: "Output tag field if true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"output_time": schema.BoolAttribute{
											Description:         "Output time field if true.",
											MarkdownDescription: "Output time field if true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_format": schema.StringAttribute{
											Description:         "Process value according to the specified format. This is available only when time_type is string",
											MarkdownDescription: "Process value according to the specified format. This is available only when time_type is string",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_format_fallbacks": schema.StringAttribute{
											Description:         "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
											MarkdownDescription: "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_type": schema.StringAttribute{
											Description:         "parses/formats value according to this type, default is string",
											MarkdownDescription: "parses/formats value according to this type, default is string",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("float", "unixtime", "string", "mixed"),
											},
										},

										"timezone": schema.StringAttribute{
											Description:         "Uses the specified timezone.",
											MarkdownDescription: "Uses the specified timezone.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "The @type parameter specifies the type of the plugin.",
											MarkdownDescription: "The @type parameter specifies the type of the plugin.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("out_file", "json", "ltsv", "csv", "msgpack", "hash", "single_value"),
											},
										},

										"utc": schema.BoolAttribute{
											Description:         "If true, uses UTC.",
											MarkdownDescription: "If true, uses UTC.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"forward": schema.SingleNestedAttribute{
									Description:         "out_forward plugin",
									MarkdownDescription: "out_forward plugin",
									Attributes: map[string]schema.Attribute{
										"ack_response_timeout": schema.StringAttribute{
											Description:         "This option is used when require_ack_response is true. This default value is based on popular tcp_syn_retries.",
											MarkdownDescription: "This option is used when require_ack_response is true. This default value is based on popular tcp_syn_retries.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"connect_timeout": schema.StringAttribute{
											Description:         "The connection timeout for the socket. When the connection is timed out during the connection establishment, Errno::ETIMEDOUT error is raised.",
											MarkdownDescription: "The connection timeout for the socket. When the connection is timed out during the connection establishment, Errno::ETIMEDOUT error is raised.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"dns_round_robin": schema.BoolAttribute{
											Description:         "Enable client-side DNS round robin. Uniform randomly pick an IP address to send data when a hostname has several IP addresses. heartbeat_type udp is not available with dns_round_robintrue. Use heartbeat_type tcp or heartbeat_type none.",
											MarkdownDescription: "Enable client-side DNS round robin. Uniform randomly pick an IP address to send data when a hostname has several IP addresses. heartbeat_type udp is not available with dns_round_robintrue. Use heartbeat_type tcp or heartbeat_type none.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"expire_dns_cache": schema.StringAttribute{
											Description:         "Sets TTL to expire DNS cache in seconds. Set 0 not to use DNS Cache.",
											MarkdownDescription: "Sets TTL to expire DNS cache in seconds. Set 0 not to use DNS Cache.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"hard_timeout": schema.StringAttribute{
											Description:         "The hard timeout used to detect server failure. The default value is equal to the send_timeout parameter.",
											MarkdownDescription: "The hard timeout used to detect server failure. The default value is equal to the send_timeout parameter.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"heartbeat_interval": schema.StringAttribute{
											Description:         "The interval of the heartbeat packer.",
											MarkdownDescription: "The interval of the heartbeat packer.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"heartbeat_type": schema.StringAttribute{
											Description:         "Specifies the transport protocol for heartbeats. Set none to disable.",
											MarkdownDescription: "Specifies the transport protocol for heartbeats. Set none to disable.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("transport", "tcp", "udp", "none"),
											},
										},

										"ignore_network_errors_at_startup": schema.BoolAttribute{
											Description:         "Ignores DNS resolution and errors at startup time.",
											MarkdownDescription: "Ignores DNS resolution and errors at startup time.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keepalive": schema.BoolAttribute{
											Description:         "Enables the keepalive connection.",
											MarkdownDescription: "Enables the keepalive connection.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keepalive_timeout": schema.StringAttribute{
											Description:         "Timeout for keepalive. Default value is nil which means to keep the connection alive as long as possible.",
											MarkdownDescription: "Timeout for keepalive. Default value is nil which means to keep the connection alive as long as possible.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"phi_failure_detector": schema.BoolAttribute{
											Description:         "Use the 'Phi accrual failure detector' to detect server failure.",
											MarkdownDescription: "Use the 'Phi accrual failure detector' to detect server failure.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"phi_threshold": schema.Int64Attribute{
											Description:         "The threshold parameter used to detect server faults.",
											MarkdownDescription: "The threshold parameter used to detect server faults.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"recover_wait": schema.StringAttribute{
											Description:         "The wait time before accepting a server fault recovery.",
											MarkdownDescription: "The wait time before accepting a server fault recovery.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"require_ack_response": schema.BoolAttribute{
											Description:         "Changes the protocol to at-least-once. The plugin waits the ack from destination's in_forward plugin.",
											MarkdownDescription: "Changes the protocol to at-least-once. The plugin waits the ack from destination's in_forward plugin.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"security": schema.SingleNestedAttribute{
											Description:         "ServiceDiscovery defines the security section",
											MarkdownDescription: "ServiceDiscovery defines the security section",
											Attributes: map[string]schema.Attribute{
												"allow_anonymous_source": schema.StringAttribute{
													Description:         "Allows the anonymous source. <client> sections are required, if disabled.",
													MarkdownDescription: "Allows the anonymous source. <client> sections are required, if disabled.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"self_hostname": schema.StringAttribute{
													Description:         "The hostname.",
													MarkdownDescription: "The hostname.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"shared_key": schema.StringAttribute{
													Description:         "The shared key for authentication.",
													MarkdownDescription: "The shared key for authentication.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user": schema.SingleNestedAttribute{
													Description:         "Defines user section directly.",
													MarkdownDescription: "Defines user section directly.",
													Attributes: map[string]schema.Attribute{
														"password": schema.SingleNestedAttribute{
															Description:         "Secret defines the key of a value.",
															MarkdownDescription: "Secret defines the key of a value.",
															Attributes: map[string]schema.Attribute{
																"value_from": schema.SingleNestedAttribute{
																	Description:         "ValueSource defines how to find a value's key.",
																	MarkdownDescription: "ValueSource defines how to find a value's key.",
																	Attributes: map[string]schema.Attribute{
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"username": schema.SingleNestedAttribute{
															Description:         "Secret defines the key of a value.",
															MarkdownDescription: "Secret defines the key of a value.",
															Attributes: map[string]schema.Attribute{
																"value_from": schema.SingleNestedAttribute{
																	Description:         "ValueSource defines how to find a value's key.",
																	MarkdownDescription: "ValueSource defines how to find a value's key.",
																	Attributes: map[string]schema.Attribute{
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
															Required: false,
															Optional: true,
															Computed: false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"user_auth": schema.StringAttribute{
													Description:         "If true, user-based authentication is used.",
													MarkdownDescription: "If true, user-based authentication is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"send_timeout": schema.StringAttribute{
											Description:         "The timeout time when sending event logs.",
											MarkdownDescription: "The timeout time when sending event logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"servers": schema.ListNestedAttribute{
											Description:         "Servers defines the servers section, at least one is required",
											MarkdownDescription: "Servers defines the servers section, at least one is required",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Host defines the IP address or host name of the server.",
														MarkdownDescription: "Host defines the IP address or host name of the server.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "The @id parameter specifies a unique name for the configuration.",
														MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"log_level": schema.StringAttribute{
														Description:         "The @log_level parameter specifies the plugin-specific logging level",
														MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name defines the name of the server. Used for logging and certificate verification in TLS transport (when the host is the address).",
														MarkdownDescription: "Name defines the name of the server. Used for logging and certificate verification in TLS transport (when the host is the address).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.SingleNestedAttribute{
														Description:         "Password defines the password for authentication.",
														MarkdownDescription: "Password defines the password for authentication.",
														Attributes: map[string]schema.Attribute{
															"value_from": schema.SingleNestedAttribute{
																Description:         "ValueSource defines how to find a value's key.",
																MarkdownDescription: "ValueSource defines how to find a value's key.",
																Attributes: map[string]schema.Attribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": schema.StringAttribute{
														Description:         "Port defines the port number of the host. Note that both TCP packets (event stream) and UDP packets (heartbeat messages) are sent to this port.",
														MarkdownDescription: "Port defines the port number of the host. Note that both TCP packets (event stream) and UDP packets (heartbeat messages) are sent to this port.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"shared_key": schema.StringAttribute{
														Description:         "SharedKey defines the shared key per server.",
														MarkdownDescription: "SharedKey defines the shared key per server.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"standby": schema.StringAttribute{
														Description:         "Standby marks a node as the standby node for an Active-Standby model between Fluentd nodes.",
														MarkdownDescription: "Standby marks a node as the standby node for an Active-Standby model between Fluentd nodes.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "The @type parameter specifies the type of the plugin.",
														MarkdownDescription: "The @type parameter specifies the type of the plugin.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"username": schema.SingleNestedAttribute{
														Description:         "Username defines the username for authentication.",
														MarkdownDescription: "Username defines the username for authentication.",
														Attributes: map[string]schema.Attribute{
															"value_from": schema.SingleNestedAttribute{
																Description:         "ValueSource defines how to find a value's key.",
																MarkdownDescription: "ValueSource defines how to find a value's key.",
																Attributes: map[string]schema.Attribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"weight": schema.StringAttribute{
														Description:         "Weight defines the load balancing weight",
														MarkdownDescription: "Weight defines the load balancing weight",
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

										"service_discovery": schema.SingleNestedAttribute{
											Description:         "ServiceDiscovery defines the service_discovery section",
											MarkdownDescription: "ServiceDiscovery defines the service_discovery section",
											Attributes: map[string]schema.Attribute{
												"conf_encoding": schema.StringAttribute{
													Description:         "The encoding of the configuration file.",
													MarkdownDescription: "The encoding of the configuration file.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dns_lookup": schema.StringAttribute{
													Description:         "DnsLookup resolves the hostname to IP address of the SRV's Target.",
													MarkdownDescription: "DnsLookup resolves the hostname to IP address of the SRV's Target.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dns_server_host": schema.StringAttribute{
													Description:         "DnsServerHost defines the hostname of the DNS server to request the SRV record.",
													MarkdownDescription: "DnsServerHost defines the hostname of the DNS server to request the SRV record.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname": schema.StringAttribute{
													Description:         "The name in RFC2782.",
													MarkdownDescription: "The name in RFC2782.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
													Description:         "The @id parameter specifies a unique name for the configuration.",
													MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Interval defines the interval of sending requests to DNS server.",
													MarkdownDescription: "Interval defines the interval of sending requests to DNS server.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"log_level": schema.StringAttribute{
													Description:         "The @log_level parameter specifies the plugin-specific logging level",
													MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "The path of the target list. Default is '/etc/fluent/sd.yaml'",
													MarkdownDescription: "The path of the target list. Default is '/etc/fluent/sd.yaml'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proto": schema.StringAttribute{
													Description:         "Proto without the underscore in RFC2782.",
													MarkdownDescription: "Proto without the underscore in RFC2782.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"server": schema.SingleNestedAttribute{
													Description:         "The server section of this plugin",
													MarkdownDescription: "The server section of this plugin",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host defines the IP address or host name of the server.",
															MarkdownDescription: "Host defines the IP address or host name of the server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"id": schema.StringAttribute{
															Description:         "The @id parameter specifies a unique name for the configuration.",
															MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_level": schema.StringAttribute{
															Description:         "The @log_level parameter specifies the plugin-specific logging level",
															MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name defines the name of the server. Used for logging and certificate verification in TLS transport (when the host is the address).",
															MarkdownDescription: "Name defines the name of the server. Used for logging and certificate verification in TLS transport (when the host is the address).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password": schema.SingleNestedAttribute{
															Description:         "Password defines the password for authentication.",
															MarkdownDescription: "Password defines the password for authentication.",
															Attributes: map[string]schema.Attribute{
																"value_from": schema.SingleNestedAttribute{
																	Description:         "ValueSource defines how to find a value's key.",
																	MarkdownDescription: "ValueSource defines how to find a value's key.",
																	Attributes: map[string]schema.Attribute{
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"port": schema.StringAttribute{
															Description:         "Port defines the port number of the host. Note that both TCP packets (event stream) and UDP packets (heartbeat messages) are sent to this port.",
															MarkdownDescription: "Port defines the port number of the host. Note that both TCP packets (event stream) and UDP packets (heartbeat messages) are sent to this port.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"shared_key": schema.StringAttribute{
															Description:         "SharedKey defines the shared key per server.",
															MarkdownDescription: "SharedKey defines the shared key per server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"standby": schema.StringAttribute{
															Description:         "Standby marks a node as the standby node for an Active-Standby model between Fluentd nodes.",
															MarkdownDescription: "Standby marks a node as the standby node for an Active-Standby model between Fluentd nodes.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "The @type parameter specifies the type of the plugin.",
															MarkdownDescription: "The @type parameter specifies the type of the plugin.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"username": schema.SingleNestedAttribute{
															Description:         "Username defines the username for authentication.",
															MarkdownDescription: "Username defines the username for authentication.",
															Attributes: map[string]schema.Attribute{
																"value_from": schema.SingleNestedAttribute{
																	Description:         "ValueSource defines how to find a value's key.",
																	MarkdownDescription: "ValueSource defines how to find a value's key.",
																	Attributes: map[string]schema.Attribute{
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"weight": schema.StringAttribute{
															Description:         "Weight defines the load balancing weight",
															MarkdownDescription: "Weight defines the load balancing weight",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service": schema.StringAttribute{
													Description:         "Service without the underscore in RFC2782.",
													MarkdownDescription: "Service without the underscore in RFC2782.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "The @type parameter specifies the type of the plugin.",
													MarkdownDescription: "The @type parameter specifies the type of the plugin.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("static", "file", "srv"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls_allow_self_signed_cert": schema.BoolAttribute{
											Description:         "Allows self-signed certificates or not.",
											MarkdownDescription: "Allows self-signed certificates or not.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_cert_logical_store_name": schema.StringAttribute{
											Description:         "The certificate logical store name on Windows system certstore. This parameter is for Windows only.",
											MarkdownDescription: "The certificate logical store name on Windows system certstore. This parameter is for Windows only.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_cert_path": schema.StringAttribute{
											Description:         "The additional CA certificate path for TLS.",
											MarkdownDescription: "The additional CA certificate path for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_cert_thumbprint": schema.StringAttribute{
											Description:         "The certificate thumbprint for searching from Windows system certstore. This parameter is for Windows only.",
											MarkdownDescription: "The certificate thumbprint for searching from Windows system certstore. This parameter is for Windows only.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_cert_use_enterprise_store": schema.BoolAttribute{
											Description:         "Enables the certificate enterprise store on Windows system certstore. This parameter is for Windows only.",
											MarkdownDescription: "Enables the certificate enterprise store on Windows system certstore. This parameter is for Windows only.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_ciphers": schema.StringAttribute{
											Description:         "The cipher configuration of TLS transport.",
											MarkdownDescription: "The cipher configuration of TLS transport.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_client_cert_path": schema.StringAttribute{
											Description:         "The client certificate path for TLS.",
											MarkdownDescription: "The client certificate path for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_client_private_key_passphrase": schema.StringAttribute{
											Description:         "The TLS private key passphrase for the client.",
											MarkdownDescription: "The TLS private key passphrase for the client.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_client_private_key_path": schema.StringAttribute{
											Description:         "The client private key path for TLS.",
											MarkdownDescription: "The client private key path for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_insecure_mode": schema.BoolAttribute{
											Description:         "Skips all verification of certificates or not.",
											MarkdownDescription: "Skips all verification of certificates or not.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_verify_hostname": schema.BoolAttribute{
											Description:         "Verifies hostname of servers and certificates or not in TLS transport.",
											MarkdownDescription: "Verifies hostname of servers and certificates or not in TLS transport.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_version": schema.StringAttribute{
											Description:         "The default version of TLS transport.",
											MarkdownDescription: "The default version of TLS transport.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLSv1_1", "TLSv1_2"),
											},
										},

										"verify_connection_at_startup": schema.BoolAttribute{
											Description:         "Verify that a connection can be made with one of out_forward nodes at the time of startup.",
											MarkdownDescription: "Verify that a connection can be made with one of out_forward nodes at the time of startup.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "out_http plugin",
									MarkdownDescription: "out_http plugin",
									Attributes: map[string]schema.Attribute{
										"auth": schema.SingleNestedAttribute{
											Description:         "Auth section for this plugin",
											MarkdownDescription: "Auth section for this plugin",
											Attributes: map[string]schema.Attribute{
												"auth": schema.StringAttribute{
													Description:         "The method for HTTP authentication. Now only basic.",
													MarkdownDescription: "The method for HTTP authentication. Now only basic.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password": schema.SingleNestedAttribute{
													Description:         "The password for basic authentication.",
													MarkdownDescription: "The password for basic authentication.",
													Attributes: map[string]schema.Attribute{
														"value_from": schema.SingleNestedAttribute{
															Description:         "ValueSource defines how to find a value's key.",
															MarkdownDescription: "ValueSource defines how to find a value's key.",
															Attributes: map[string]schema.Attribute{
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"username": schema.SingleNestedAttribute{
													Description:         "The username for basic authentication.",
													MarkdownDescription: "The username for basic authentication.",
													Attributes: map[string]schema.Attribute{
														"value_from": schema.SingleNestedAttribute{
															Description:         "ValueSource defines how to find a value's key.",
															MarkdownDescription: "ValueSource defines how to find a value's key.",
															Attributes: map[string]schema.Attribute{
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
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"content_type": schema.StringAttribute{
											Description:         "ContentType defines Content-Type for HTTP request. out_http automatically set Content-Type for built-in formatters when this parameter is not specified.",
											MarkdownDescription: "ContentType defines Content-Type for HTTP request. out_http automatically set Content-Type for built-in formatters when this parameter is not specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"endpoint": schema.StringAttribute{
											Description:         "Endpoint defines the endpoint for HTTP request. If you want to use HTTPS, use https prefix.",
											MarkdownDescription: "Endpoint defines the endpoint for HTTP request. If you want to use HTTPS, use https prefix.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"error_response_as_unrecoverable": schema.BoolAttribute{
											Description:         "Raise UnrecoverableError when the response code is not SUCCESS.",
											MarkdownDescription: "Raise UnrecoverableError when the response code is not SUCCESS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"headers": schema.StringAttribute{
											Description:         "Headers defines the additional headers for HTTP request.",
											MarkdownDescription: "Headers defines the additional headers for HTTP request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"headers_from_placeholders": schema.StringAttribute{
											Description:         "Additional placeholder based headers for HTTP request. If you want to use tag or record field, use this parameter instead of headers.",
											MarkdownDescription: "Additional placeholder based headers for HTTP request. If you want to use tag or record field, use this parameter instead of headers.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http_method": schema.StringAttribute{
											Description:         "HttpMethod defines the method for HTTP request.",
											MarkdownDescription: "HttpMethod defines the method for HTTP request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("post", "put"),
											},
										},

										"json_array": schema.BoolAttribute{
											Description:         "JsonArray defines whether to use the array format of JSON or not",
											MarkdownDescription: "JsonArray defines whether to use the array format of JSON or not",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"open_timeout": schema.Int64Attribute{
											Description:         "OpenTimeout defines the connection open timeout in seconds.",
											MarkdownDescription: "OpenTimeout defines the connection open timeout in seconds.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy": schema.StringAttribute{
											Description:         "Proxy defines the proxy for HTTP request.",
											MarkdownDescription: "Proxy defines the proxy for HTTP request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_timeout": schema.Int64Attribute{
											Description:         "ReadTimeout defines the read timeout in seconds.",
											MarkdownDescription: "ReadTimeout defines the read timeout in seconds.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retryable_response_codes": schema.StringAttribute{
											Description:         "The list of retryable response codes. If the response code is included in this list, out_http retries the buffer flush.",
											MarkdownDescription: "The list of retryable response codes. If the response code is included in this list, out_http retries the buffer flush.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ssl_timeout": schema.Int64Attribute{
											Description:         "SslTimeout defines the TLS timeout in seconds.",
											MarkdownDescription: "SslTimeout defines the TLS timeout in seconds.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_ca_cert_path": schema.StringAttribute{
											Description:         "TlsCaCertPath defines the CA certificate path for TLS.",
											MarkdownDescription: "TlsCaCertPath defines the CA certificate path for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_ciphers": schema.StringAttribute{
											Description:         "TlsCiphers defines the cipher suites configuration of TLS.",
											MarkdownDescription: "TlsCiphers defines the cipher suites configuration of TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_client_cert_path": schema.StringAttribute{
											Description:         "TlsClientCertPath defines the client certificate path for TLS.",
											MarkdownDescription: "TlsClientCertPath defines the client certificate path for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_private_key_passphrase": schema.StringAttribute{
											Description:         "TlsPrivateKeyPassphrase defines the client private key passphrase for TLS.",
											MarkdownDescription: "TlsPrivateKeyPassphrase defines the client private key passphrase for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_private_key_path": schema.StringAttribute{
											Description:         "TlsPrivateKeyPath defines the client private key path for TLS.",
											MarkdownDescription: "TlsPrivateKeyPath defines the client private key path for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_verify_mode": schema.StringAttribute{
											Description:         "TlsVerifyMode defines the verify mode of TLS.",
											MarkdownDescription: "TlsVerifyMode defines the verify mode of TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("peer", "none"),
											},
										},

										"tls_version": schema.StringAttribute{
											Description:         "TlsVersion defines the default version of TLS transport.",
											MarkdownDescription: "TlsVersion defines the default version of TLS transport.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLSv1_1", "TLSv1_2"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"inject": schema.SingleNestedAttribute{
									Description:         "inject section",
									MarkdownDescription: "inject section",
									Attributes: map[string]schema.Attribute{
										"hostname": schema.StringAttribute{
											Description:         "Hostname value",
											MarkdownDescription: "Hostname value",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hostname_key": schema.StringAttribute{
											Description:         "The field name to inject hostname",
											MarkdownDescription: "The field name to inject hostname",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"inline": schema.SingleNestedAttribute{
											Description:         "Time section",
											MarkdownDescription: "Time section",
											Attributes: map[string]schema.Attribute{
												"localtime": schema.BoolAttribute{
													Description:         "If true, uses local time.",
													MarkdownDescription: "If true, uses local time.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_format": schema.StringAttribute{
													Description:         "Process value according to the specified format. This is available only when time_type is string",
													MarkdownDescription: "Process value according to the specified format. This is available only when time_type is string",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_format_fallbacks": schema.StringAttribute{
													Description:         "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
													MarkdownDescription: "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_type": schema.StringAttribute{
													Description:         "parses/formats value according to this type, default is string",
													MarkdownDescription: "parses/formats value according to this type, default is string",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("float", "unixtime", "string", "mixed"),
													},
												},

												"timezone": schema.StringAttribute{
													Description:         "Uses the specified timezone.",
													MarkdownDescription: "Uses the specified timezone.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"utc": schema.BoolAttribute{
													Description:         "If true, uses UTC.",
													MarkdownDescription: "If true, uses UTC.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tag_key": schema.StringAttribute{
											Description:         "The field name to inject tag",
											MarkdownDescription: "The field name to inject tag",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_key": schema.StringAttribute{
											Description:         "The field name to inject time",
											MarkdownDescription: "The field name to inject time",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"worker_id_key": schema.StringAttribute{
											Description:         "The field name to inject worker_id",
											MarkdownDescription: "The field name to inject worker_id",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kafka": schema.SingleNestedAttribute{
									Description:         "out_kafka plugin",
									MarkdownDescription: "out_kafka plugin",
									Attributes: map[string]schema.Attribute{
										"brokers": schema.StringAttribute{
											Description:         "The list of all seed brokers, with their host and port information. Default: localhost:9092",
											MarkdownDescription: "The list of all seed brokers, with their host and port information. Default: localhost:9092",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"compression_codec": schema.StringAttribute{
											Description:         "The codec the producer uses to compress messages (default: nil).",
											MarkdownDescription: "The codec the producer uses to compress messages (default: nil).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("gzip", "snappy"),
											},
										},

										"default_topic": schema.StringAttribute{
											Description:         "The name of the default topic. (default: nil)",
											MarkdownDescription: "The name of the default topic. (default: nil)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"required_acks": schema.Int64Attribute{
											Description:         "The number of acks required per request.",
											MarkdownDescription: "The number of acks required per request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"topic_key": schema.StringAttribute{
											Description:         "The field name for the target topic. If the field value is app, this plugin writes events to the app topic.",
											MarkdownDescription: "The field name for the target topic. If the field value is app, this plugin writes events to the app topic.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_event_time": schema.BoolAttribute{
											Description:         "Set fluentd event time to Kafka's CreateTime.",
											MarkdownDescription: "Set fluentd event time to Kafka's CreateTime.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"log_level": schema.StringAttribute{
									Description:         "The @log_level parameter specifies the plugin-specific logging level",
									MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"loki": schema.SingleNestedAttribute{
									Description:         "out_loki plugin",
									MarkdownDescription: "out_loki plugin",
									Attributes: map[string]schema.Attribute{
										"drop_single_key": schema.BoolAttribute{
											Description:         "If a record only has 1 key, then just set the log line to the value and discard the key.",
											MarkdownDescription: "If a record only has 1 key, then just set the log line to the value and discard the key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"extract_kubernetes_labels": schema.BoolAttribute{
											Description:         "If set to true, it will add all Kubernetes labels to the Stream labels.",
											MarkdownDescription: "If set to true, it will add all Kubernetes labels to the Stream labels.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http_password": schema.SingleNestedAttribute{
											Description:         "Password for user defined in HTTP_User Set HTTP basic authentication password",
											MarkdownDescription: "Password for user defined in HTTP_User Set HTTP basic authentication password",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http_user": schema.SingleNestedAttribute{
											Description:         "Set HTTP basic authentication user name.",
											MarkdownDescription: "Set HTTP basic authentication user name.",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"include_thread_label": schema.BoolAttribute{
											Description:         "Whether or not to include the fluentd_thread label when multiple threads are used for flushing",
											MarkdownDescription: "Whether or not to include the fluentd_thread label when multiple threads are used for flushing",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure": schema.BoolAttribute{
											Description:         "Disable certificate validation",
											MarkdownDescription: "Disable certificate validation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"label_keys": schema.ListAttribute{
											Description:         "Optional list of record keys that will be placed as stream labels. This configuration property is for records key only.",
											MarkdownDescription: "Optional list of record keys that will be placed as stream labels. This configuration property is for records key only.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.ListAttribute{
											Description:         "Stream labels for API request. It can be multiple comma separated of strings specifying  key=value pairs. In addition to fixed parameters, it also allows to add custom record keys (similar to label_keys property).",
											MarkdownDescription: "Stream labels for API request. It can be multiple comma separated of strings specifying  key=value pairs. In addition to fixed parameters, it also allows to add custom record keys (similar to label_keys property).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"line_format": schema.StringAttribute{
											Description:         "Format to use when flattening the record to a log line. Valid values are json or key_value. If set to json,  the log line sent to Loki will be the Fluentd record dumped as JSON. If set to key_value, the log line will be each item in the record concatenated together (separated by a single space) in the format.",
											MarkdownDescription: "Format to use when flattening the record to a log line. Valid values are json or key_value. If set to json,  the log line sent to Loki will be the Fluentd record dumped as JSON. If set to key_value, the log line will be each item in the record concatenated together (separated by a single space) in the format.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("json", "key_value"),
											},
										},

										"remove_keys": schema.ListAttribute{
											Description:         "Optional list of record keys that will be removed from stream labels. This configuration property is for records key only.",
											MarkdownDescription: "Optional list of record keys that will be removed from stream labels. This configuration property is for records key only.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tenant_id": schema.SingleNestedAttribute{
											Description:         "Tenant ID used by default to push logs to Loki. If omitted or empty it assumes Loki is running in single-tenant mode and no X-Scope-OrgID header is sent.",
											MarkdownDescription: "Tenant ID used by default to push logs to Loki. If omitted or empty it assumes Loki is running in single-tenant mode and no X-Scope-OrgID header is sent.",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls_ca_cert_file": schema.StringAttribute{
											Description:         "TlsCaCert defines the CA certificate file for TLS.",
											MarkdownDescription: "TlsCaCert defines the CA certificate file for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_client_cert_file": schema.StringAttribute{
											Description:         "TlsClientCert defines the client certificate file for TLS.",
											MarkdownDescription: "TlsClientCert defines the client certificate file for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_private_key_file": schema.StringAttribute{
											Description:         "TlsPrivateKey defines the client private key file for TLS.",
											MarkdownDescription: "TlsPrivateKey defines the client private key file for TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"url": schema.StringAttribute{
											Description:         "Loki URL.",
											MarkdownDescription: "Loki URL.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"opensearch": schema.SingleNestedAttribute{
									Description:         "out_opensearch plugin",
									MarkdownDescription: "out_opensearch plugin",
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "The hostname of your Opensearch node (default: localhost).",
											MarkdownDescription: "The hostname of your Opensearch node (default: localhost).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hosts": schema.StringAttribute{
											Description:         "Hosts defines a list of hosts if you want to connect to more than one Openearch nodes",
											MarkdownDescription: "Hosts defines a list of hosts if you want to connect to more than one Openearch nodes",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"index_name": schema.StringAttribute{
											Description:         "IndexName defines the placeholder syntax of Fluentd plugin API. See https://docs.fluentd.org/configuration/buffer-section.",
											MarkdownDescription: "IndexName defines the placeholder syntax of Fluentd plugin API. See https://docs.fluentd.org/configuration/buffer-section.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logstash_format": schema.BoolAttribute{
											Description:         "If true, Fluentd uses the conventional index name format logstash-%Y.%m.%d (default: false). This option supersedes the index_name option.",
											MarkdownDescription: "If true, Fluentd uses the conventional index name format logstash-%Y.%m.%d (default: false). This option supersedes the index_name option.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logstash_prefix": schema.StringAttribute{
											Description:         "LogstashPrefix defines the logstash prefix index name to write events when logstash_format is true (default: logstash).",
											MarkdownDescription: "LogstashPrefix defines the logstash prefix index name to write events when logstash_format is true (default: logstash).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"password": schema.SingleNestedAttribute{
											Description:         "Optional, The login credentials to connect to Opensearch",
											MarkdownDescription: "Optional, The login credentials to connect to Opensearch",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "Path defines the REST API endpoint of Opensearch to post write requests (default: nil).",
											MarkdownDescription: "Path defines the REST API endpoint of Opensearch to post write requests (default: nil).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port number of your Opensearch node (default: 9200).",
											MarkdownDescription: "The port number of your Opensearch node (default: 9200).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"scheme": schema.StringAttribute{
											Description:         "Specify https if your Opensearch endpoint supports SSL (default: http).",
											MarkdownDescription: "Specify https if your Opensearch endpoint supports SSL (default: http).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user": schema.SingleNestedAttribute{
											Description:         "Optional, The login credentials to connect to Opensearch",
											MarkdownDescription: "Optional, The login credentials to connect to Opensearch",
											Attributes: map[string]schema.Attribute{
												"value_from": schema.SingleNestedAttribute{
													Description:         "ValueSource defines how to find a value's key.",
													MarkdownDescription: "ValueSource defines how to find a value's key.",
													Attributes: map[string]schema.Attribute{
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
											Required: false,
											Optional: true,
											Computed: false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"s3": schema.SingleNestedAttribute{
									Description:         "out_s3 plugin",
									MarkdownDescription: "out_s3 plugin",
									Attributes: map[string]schema.Attribute{
										"aws_key_id": schema.StringAttribute{
											Description:         "The AWS access key id.",
											MarkdownDescription: "The AWS access key id.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"aws_sec_key": schema.StringAttribute{
											Description:         "The AWS secret key.",
											MarkdownDescription: "The AWS secret key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"force_path_style": schema.BoolAttribute{
											Description:         "This prevents AWS SDK from breaking endpoint URL",
											MarkdownDescription: "This prevents AWS SDK from breaking endpoint URL",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "The path prefix of the files on S3.",
											MarkdownDescription: "The path prefix of the files on S3.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_uri": schema.StringAttribute{
											Description:         "The proxy URL.",
											MarkdownDescription: "The proxy URL.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_bucket": schema.StringAttribute{
											Description:         "The Amazon S3 bucket name.",
											MarkdownDescription: "The Amazon S3 bucket name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_endpoint": schema.StringAttribute{
											Description:         "The endpoint URL (like 'http://localhost:9000/')",
											MarkdownDescription: "The endpoint URL (like 'http://localhost:9000/')",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_object_key_format": schema.StringAttribute{
											Description:         "The actual S3 path. This is interpolated to the actual path.",
											MarkdownDescription: "The actual S3 path. This is interpolated to the actual path.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_region": schema.StringAttribute{
											Description:         "The Amazon S3 region name",
											MarkdownDescription: "The Amazon S3 region name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sse_customer_algorithm": schema.StringAttribute{
											Description:         "The AWS KMS enctyption algorithm.",
											MarkdownDescription: "The AWS KMS enctyption algorithm.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sse_customer_key": schema.StringAttribute{
											Description:         "The AWS KMS key.",
											MarkdownDescription: "The AWS KMS key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sse_customer_key_md5": schema.StringAttribute{
											Description:         "The AWS KMS key MD5.",
											MarkdownDescription: "The AWS KMS key MD5.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ssekms_key_id": schema.StringAttribute{
											Description:         "The AWS KMS key ID.",
											MarkdownDescription: "The AWS KMS key ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ssl_verify_peer": schema.BoolAttribute{
											Description:         "Verify the SSL certificate of the endpoint.",
											MarkdownDescription: "Verify the SSL certificate of the endpoint.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"store_as": schema.StringAttribute{
											Description:         "The compression type.",
											MarkdownDescription: "The compression type.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("gzip", "lzo", "json", "txt"),
											},
										},

										"time_slice_format": schema.StringAttribute{
											Description:         "This timestamp is added to each file name",
											MarkdownDescription: "This timestamp is added to each file name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_server_side_encryption": schema.StringAttribute{
											Description:         "the following parameters are for S3 kms https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html",
											MarkdownDescription: "the following parameters are for S3 kms https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"stdout": schema.MapAttribute{
									Description:         "out_stdout plugin",
									MarkdownDescription: "out_stdout plugin",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tag": schema.StringAttribute{
									Description:         "Which tag to be matched.",
									MarkdownDescription: "Which tag to be matched.",
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
		},
	}
}

func (r *FluentdFluentIoClusterOutputV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentd_fluent_io_cluster_output_v1alpha1_manifest")

	var model FluentdFluentIoClusterOutputV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentd.fluent.io/v1alpha1")
	model.Kind = pointer.String("ClusterOutput")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
