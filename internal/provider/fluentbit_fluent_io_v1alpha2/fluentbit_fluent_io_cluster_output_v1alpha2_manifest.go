/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentbit_fluent_io_v1alpha2

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
	_ datasource.DataSource = &FluentbitFluentIoClusterOutputV1Alpha2Manifest{}
)

func NewFluentbitFluentIoClusterOutputV1Alpha2Manifest() datasource.DataSource {
	return &FluentbitFluentIoClusterOutputV1Alpha2Manifest{}
}

type FluentbitFluentIoClusterOutputV1Alpha2Manifest struct{}

type FluentbitFluentIoClusterOutputV1Alpha2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Alias     *string `tfsdk:"alias" json:"alias,omitempty"`
		AzureBlob *struct {
			AccountName         *string `tfsdk:"account_name" json:"accountName,omitempty"`
			AutoCreateContainer *string `tfsdk:"auto_create_container" json:"autoCreateContainer,omitempty"`
			BlobType            *string `tfsdk:"blob_type" json:"blobType,omitempty"`
			ContainerName       *string `tfsdk:"container_name" json:"containerName,omitempty"`
			EmulatorMode        *string `tfsdk:"emulator_mode" json:"emulatorMode,omitempty"`
			Endpoint            *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Networking          *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Path      *string `tfsdk:"path" json:"path,omitempty"`
			SharedKey *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"shared_key" json:"sharedKey,omitempty"`
			Tls *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"azure_blob" json:"azureBlob,omitempty"`
		AzureLogAnalytics *struct {
			CustomerID *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"customer_id" json:"customerID,omitempty"`
			LogType   *string `tfsdk:"log_type" json:"logType,omitempty"`
			SharedKey *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"shared_key" json:"sharedKey,omitempty"`
			TimeGenerated *bool   `tfsdk:"time_generated" json:"timeGenerated,omitempty"`
			TimeKey       *string `tfsdk:"time_key" json:"timeKey,omitempty"`
		} `tfsdk:"azure_log_analytics" json:"azureLogAnalytics,omitempty"`
		CloudWatch *struct {
			AutoCreateGroup   *bool   `tfsdk:"auto_create_group" json:"autoCreateGroup,omitempty"`
			AutoRetryRequests *bool   `tfsdk:"auto_retry_requests" json:"autoRetryRequests,omitempty"`
			Endpoint          *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			ExternalID        *string `tfsdk:"external_id" json:"externalID,omitempty"`
			LogFormat         *string `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogGroupName      *string `tfsdk:"log_group_name" json:"logGroupName,omitempty"`
			LogGroupTemplate  *string `tfsdk:"log_group_template" json:"logGroupTemplate,omitempty"`
			LogKey            *string `tfsdk:"log_key" json:"logKey,omitempty"`
			LogRetentionDays  *int64  `tfsdk:"log_retention_days" json:"logRetentionDays,omitempty"`
			LogStreamName     *string `tfsdk:"log_stream_name" json:"logStreamName,omitempty"`
			LogStreamPrefix   *string `tfsdk:"log_stream_prefix" json:"logStreamPrefix,omitempty"`
			LogStreamTemplate *string `tfsdk:"log_stream_template" json:"logStreamTemplate,omitempty"`
			MetricDimensions  *string `tfsdk:"metric_dimensions" json:"metricDimensions,omitempty"`
			MetricNamespace   *string `tfsdk:"metric_namespace" json:"metricNamespace,omitempty"`
			Region            *string `tfsdk:"region" json:"region,omitempty"`
			RoleArn           *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
			StsEndpoint       *string `tfsdk:"sts_endpoint" json:"stsEndpoint,omitempty"`
		} `tfsdk:"cloud_watch" json:"cloudWatch,omitempty"`
		CustomPlugin *struct {
			Config     *string            `tfsdk:"config" json:"config,omitempty"`
			YamlConfig *map[string]string `tfsdk:"yaml_config" json:"yamlConfig,omitempty"`
		} `tfsdk:"custom_plugin" json:"customPlugin,omitempty"`
		Datadog *struct {
			Apikey *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"apikey" json:"apikey,omitempty"`
			Compress        *string `tfsdk:"compress" json:"compress,omitempty"`
			Dd_message_key  *string `tfsdk:"dd_message_key" json:"dd_message_key,omitempty"`
			Dd_service      *string `tfsdk:"dd_service" json:"dd_service,omitempty"`
			Dd_source       *string `tfsdk:"dd_source" json:"dd_source,omitempty"`
			Dd_tags         *string `tfsdk:"dd_tags" json:"dd_tags,omitempty"`
			Host            *string `tfsdk:"host" json:"host,omitempty"`
			Include_tag_key *bool   `tfsdk:"include_tag_key" json:"include_tag_key,omitempty"`
			Json_date_key   *string `tfsdk:"json_date_key" json:"json_date_key,omitempty"`
			Provider        *string `tfsdk:"provider" json:"provider,omitempty"`
			Proxy           *string `tfsdk:"proxy" json:"proxy,omitempty"`
			Tag_key         *string `tfsdk:"tag_key" json:"tag_key,omitempty"`
			Tls             *bool   `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"datadog" json:"datadog,omitempty"`
		Es *struct {
			AwsAuth          *string `tfsdk:"aws_auth" json:"awsAuth,omitempty"`
			AwsExternalID    *string `tfsdk:"aws_external_id" json:"awsExternalID,omitempty"`
			AwsRegion        *string `tfsdk:"aws_region" json:"awsRegion,omitempty"`
			AwsRoleARN       *string `tfsdk:"aws_role_arn" json:"awsRoleARN,omitempty"`
			AwsSTSEndpoint   *string `tfsdk:"aws_sts_endpoint" json:"awsSTSEndpoint,omitempty"`
			BufferSize       *string `tfsdk:"buffer_size" json:"bufferSize,omitempty"`
			CloudAuth        *string `tfsdk:"cloud_auth" json:"cloudAuth,omitempty"`
			CloudID          *string `tfsdk:"cloud_id" json:"cloudID,omitempty"`
			Compress         *string `tfsdk:"compress" json:"compress,omitempty"`
			CurrentTimeIndex *bool   `tfsdk:"current_time_index" json:"currentTimeIndex,omitempty"`
			GenerateID       *bool   `tfsdk:"generate_id" json:"generateID,omitempty"`
			Host             *string `tfsdk:"host" json:"host,omitempty"`
			HttpPassword     *struct {
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
			IdKey              *string `tfsdk:"id_key" json:"idKey,omitempty"`
			IncludeTagKey      *bool   `tfsdk:"include_tag_key" json:"includeTagKey,omitempty"`
			Index              *string `tfsdk:"index" json:"index,omitempty"`
			LogstashDateFormat *string `tfsdk:"logstash_date_format" json:"logstashDateFormat,omitempty"`
			LogstashFormat     *bool   `tfsdk:"logstash_format" json:"logstashFormat,omitempty"`
			LogstashPrefix     *string `tfsdk:"logstash_prefix" json:"logstashPrefix,omitempty"`
			LogstashPrefixKey  *string `tfsdk:"logstash_prefix_key" json:"logstashPrefixKey,omitempty"`
			Networking         *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Path             *string `tfsdk:"path" json:"path,omitempty"`
			Pipeline         *string `tfsdk:"pipeline" json:"pipeline,omitempty"`
			Port             *int64  `tfsdk:"port" json:"port,omitempty"`
			ReplaceDots      *bool   `tfsdk:"replace_dots" json:"replaceDots,omitempty"`
			SuppressTypeName *string `tfsdk:"suppress_type_name" json:"suppressTypeName,omitempty"`
			TagKey           *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
			TimeKey          *string `tfsdk:"time_key" json:"timeKey,omitempty"`
			TimeKeyFormat    *string `tfsdk:"time_key_format" json:"timeKeyFormat,omitempty"`
			TimeKeyNanos     *bool   `tfsdk:"time_key_nanos" json:"timeKeyNanos,omitempty"`
			Tls              *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			TotalLimitSize *string `tfsdk:"total_limit_size" json:"totalLimitSize,omitempty"`
			TraceError     *bool   `tfsdk:"trace_error" json:"traceError,omitempty"`
			TraceOutput    *bool   `tfsdk:"trace_output" json:"traceOutput,omitempty"`
			Type           *string `tfsdk:"type" json:"type,omitempty"`
			WriteOperation *string `tfsdk:"write_operation" json:"writeOperation,omitempty"`
		} `tfsdk:"es" json:"es,omitempty"`
		File *struct {
			Delimiter      *string `tfsdk:"delimiter" json:"delimiter,omitempty"`
			File           *string `tfsdk:"file" json:"file,omitempty"`
			Format         *string `tfsdk:"format" json:"format,omitempty"`
			LabelDelimiter *string `tfsdk:"label_delimiter" json:"labelDelimiter,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Template       *string `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"file" json:"file,omitempty"`
		Firehose *struct {
			AutoRetryRequests *bool   `tfsdk:"auto_retry_requests" json:"autoRetryRequests,omitempty"`
			DataKeys          *string `tfsdk:"data_keys" json:"dataKeys,omitempty"`
			DeliveryStream    *string `tfsdk:"delivery_stream" json:"deliveryStream,omitempty"`
			Endpoint          *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			LogKey            *string `tfsdk:"log_key" json:"logKey,omitempty"`
			Region            *string `tfsdk:"region" json:"region,omitempty"`
			RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
			StsEndpoint       *string `tfsdk:"sts_endpoint" json:"stsEndpoint,omitempty"`
			TimeKey           *string `tfsdk:"time_key" json:"timeKey,omitempty"`
			TimeKeyFormat     *string `tfsdk:"time_key_format" json:"timeKeyFormat,omitempty"`
		} `tfsdk:"firehose" json:"firehose,omitempty"`
		Forward *struct {
			EmptySharedKey *bool   `tfsdk:"empty_shared_key" json:"emptySharedKey,omitempty"`
			Host           *string `tfsdk:"host" json:"host,omitempty"`
			Networking     *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Password *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Port               *int64  `tfsdk:"port" json:"port,omitempty"`
			RequireAckResponse *bool   `tfsdk:"require_ack_response" json:"requireAckResponse,omitempty"`
			SelfHostname       *string `tfsdk:"self_hostname" json:"selfHostname,omitempty"`
			SendOptions        *bool   `tfsdk:"send_options" json:"sendOptions,omitempty"`
			SharedKey          *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
			Tag                *string `tfsdk:"tag" json:"tag,omitempty"`
			TimeAsInteger      *bool   `tfsdk:"time_as_integer" json:"timeAsInteger,omitempty"`
			Tls                *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Username *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"forward" json:"forward,omitempty"`
		Gelf *struct {
			Compress       *bool   `tfsdk:"compress" json:"compress,omitempty"`
			FullMessageKey *string `tfsdk:"full_message_key" json:"fullMessageKey,omitempty"`
			Host           *string `tfsdk:"host" json:"host,omitempty"`
			HostKey        *string `tfsdk:"host_key" json:"hostKey,omitempty"`
			LevelKey       *string `tfsdk:"level_key" json:"levelKey,omitempty"`
			Mode           *string `tfsdk:"mode" json:"mode,omitempty"`
			Networking     *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			PacketSize      *int64  `tfsdk:"packet_size" json:"packetSize,omitempty"`
			Port            *int64  `tfsdk:"port" json:"port,omitempty"`
			ShortMessageKey *string `tfsdk:"short_message_key" json:"shortMessageKey,omitempty"`
			TimestampKey    *string `tfsdk:"timestamp_key" json:"timestampKey,omitempty"`
			Tls             *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"gelf" json:"gelf,omitempty"`
		Http *struct {
			AllowDuplicatedHeaders *bool              `tfsdk:"allow_duplicated_headers" json:"allowDuplicatedHeaders,omitempty"`
			Compress               *string            `tfsdk:"compress" json:"compress,omitempty"`
			Format                 *string            `tfsdk:"format" json:"format,omitempty"`
			GelfFullMessageKey     *string            `tfsdk:"gelf_full_message_key" json:"gelfFullMessageKey,omitempty"`
			GelfHostKey            *string            `tfsdk:"gelf_host_key" json:"gelfHostKey,omitempty"`
			GelfLevelKey           *string            `tfsdk:"gelf_level_key" json:"gelfLevelKey,omitempty"`
			GelfShortMessageKey    *string            `tfsdk:"gelf_short_message_key" json:"gelfShortMessageKey,omitempty"`
			GelfTimestampKey       *string            `tfsdk:"gelf_timestamp_key" json:"gelfTimestampKey,omitempty"`
			HeaderTag              *string            `tfsdk:"header_tag" json:"headerTag,omitempty"`
			Headers                *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
			Host                   *string            `tfsdk:"host" json:"host,omitempty"`
			HttpPassword           *struct {
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
			JsonDateFormat *string `tfsdk:"json_date_format" json:"jsonDateFormat,omitempty"`
			JsonDateKey    *string `tfsdk:"json_date_key" json:"jsonDateKey,omitempty"`
			Networking     *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port  *int64  `tfsdk:"port" json:"port,omitempty"`
			Proxy *string `tfsdk:"proxy" json:"proxy,omitempty"`
			Tls   *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Uri *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
		InfluxDB *struct {
			AutoTags     *bool   `tfsdk:"auto_tags" json:"autoTags,omitempty"`
			Bucket       *string `tfsdk:"bucket" json:"bucket,omitempty"`
			Database     *string `tfsdk:"database" json:"database,omitempty"`
			Host         *string `tfsdk:"host" json:"host,omitempty"`
			HttpPassword *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"http_password" json:"httpPassword,omitempty"`
			HttpToken *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"http_token" json:"httpToken,omitempty"`
			HttpUser *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"http_user" json:"httpUser,omitempty"`
			Networking *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Org             *string   `tfsdk:"org" json:"org,omitempty"`
			Port            *int64    `tfsdk:"port" json:"port,omitempty"`
			SequenceTag     *string   `tfsdk:"sequence_tag" json:"sequenceTag,omitempty"`
			TagKeys         *[]string `tfsdk:"tag_keys" json:"tagKeys,omitempty"`
			TagListKey      *string   `tfsdk:"tag_list_key" json:"tagListKey,omitempty"`
			TagsListEnabled *bool     `tfsdk:"tags_list_enabled" json:"tagsListEnabled,omitempty"`
			Tls             *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"influx_db" json:"influxDB,omitempty"`
		Kafka *struct {
			Brokers          *string            `tfsdk:"brokers" json:"brokers,omitempty"`
			DynamicTopic     *bool              `tfsdk:"dynamic_topic" json:"dynamicTopic,omitempty"`
			Format           *string            `tfsdk:"format" json:"format,omitempty"`
			MessageKey       *string            `tfsdk:"message_key" json:"messageKey,omitempty"`
			MessageKeyField  *string            `tfsdk:"message_key_field" json:"messageKeyField,omitempty"`
			QueueFullRetries *int64             `tfsdk:"queue_full_retries" json:"queueFullRetries,omitempty"`
			Rdkafka          *map[string]string `tfsdk:"rdkafka" json:"rdkafka,omitempty"`
			TimestampFormat  *string            `tfsdk:"timestamp_format" json:"timestampFormat,omitempty"`
			TimestampKey     *string            `tfsdk:"timestamp_key" json:"timestampKey,omitempty"`
			TopicKey         *string            `tfsdk:"topic_key" json:"topicKey,omitempty"`
			Topics           *string            `tfsdk:"topics" json:"topics,omitempty"`
		} `tfsdk:"kafka" json:"kafka,omitempty"`
		Kinesis *struct {
			AutoRetryRequests *bool   `tfsdk:"auto_retry_requests" json:"autoRetryRequests,omitempty"`
			Endpoint          *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			ExternalID        *string `tfsdk:"external_id" json:"externalID,omitempty"`
			LogKey            *string `tfsdk:"log_key" json:"logKey,omitempty"`
			Region            *string `tfsdk:"region" json:"region,omitempty"`
			RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
			Stream            *string `tfsdk:"stream" json:"stream,omitempty"`
			StsEndpoint       *string `tfsdk:"sts_endpoint" json:"stsEndpoint,omitempty"`
			TimeKey           *string `tfsdk:"time_key" json:"timeKey,omitempty"`
			TimeKeyFormat     *string `tfsdk:"time_key_format" json:"timeKeyFormat,omitempty"`
		} `tfsdk:"kinesis" json:"kinesis,omitempty"`
		LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		Loki     *struct {
			AutoKubernetesLabels *string `tfsdk:"auto_kubernetes_labels" json:"autoKubernetesLabels,omitempty"`
			BearerToken          *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"bearer_token" json:"bearerToken,omitempty"`
			DropSingleKey *string `tfsdk:"drop_single_key" json:"dropSingleKey,omitempty"`
			Host          *string `tfsdk:"host" json:"host,omitempty"`
			HttpPassword  *struct {
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
			LabelKeys    *[]string `tfsdk:"label_keys" json:"labelKeys,omitempty"`
			LabelMapPath *string   `tfsdk:"label_map_path" json:"labelMapPath,omitempty"`
			Labels       *[]string `tfsdk:"labels" json:"labels,omitempty"`
			LineFormat   *string   `tfsdk:"line_format" json:"lineFormat,omitempty"`
			Networking   *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port       *int64    `tfsdk:"port" json:"port,omitempty"`
			RemoveKeys *[]string `tfsdk:"remove_keys" json:"removeKeys,omitempty"`
			TenantID   *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"tenant_id" json:"tenantID,omitempty"`
			TenantIDKey *string `tfsdk:"tenant_id_key" json:"tenantIDKey,omitempty"`
			Tls         *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Uri *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"loki" json:"loki,omitempty"`
		Match      *string            `tfsdk:"match" json:"match,omitempty"`
		MatchRegex *string            `tfsdk:"match_regex" json:"matchRegex,omitempty"`
		Null       *map[string]string `tfsdk:"null" json:"null,omitempty"`
		Opensearch *struct {
			Workers          *int64  `tfsdk:"workers" json:"Workers,omitempty"`
			AwsAuth          *string `tfsdk:"aws_auth" json:"awsAuth,omitempty"`
			AwsExternalID    *string `tfsdk:"aws_external_id" json:"awsExternalID,omitempty"`
			AwsRegion        *string `tfsdk:"aws_region" json:"awsRegion,omitempty"`
			AwsRoleARN       *string `tfsdk:"aws_role_arn" json:"awsRoleARN,omitempty"`
			AwsSTSEndpoint   *string `tfsdk:"aws_sts_endpoint" json:"awsSTSEndpoint,omitempty"`
			BufferSize       *string `tfsdk:"buffer_size" json:"bufferSize,omitempty"`
			Compress         *string `tfsdk:"compress" json:"compress,omitempty"`
			CurrentTimeIndex *bool   `tfsdk:"current_time_index" json:"currentTimeIndex,omitempty"`
			GenerateID       *bool   `tfsdk:"generate_id" json:"generateID,omitempty"`
			Host             *string `tfsdk:"host" json:"host,omitempty"`
			HttpPassword     *struct {
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
			IdKey              *string `tfsdk:"id_key" json:"idKey,omitempty"`
			IncludeTagKey      *bool   `tfsdk:"include_tag_key" json:"includeTagKey,omitempty"`
			Index              *string `tfsdk:"index" json:"index,omitempty"`
			LogstashDateFormat *string `tfsdk:"logstash_date_format" json:"logstashDateFormat,omitempty"`
			LogstashFormat     *bool   `tfsdk:"logstash_format" json:"logstashFormat,omitempty"`
			LogstashPrefix     *string `tfsdk:"logstash_prefix" json:"logstashPrefix,omitempty"`
			LogstashPrefixKey  *string `tfsdk:"logstash_prefix_key" json:"logstashPrefixKey,omitempty"`
			Networking         *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Path             *string `tfsdk:"path" json:"path,omitempty"`
			Pipeline         *string `tfsdk:"pipeline" json:"pipeline,omitempty"`
			Port             *int64  `tfsdk:"port" json:"port,omitempty"`
			ReplaceDots      *bool   `tfsdk:"replace_dots" json:"replaceDots,omitempty"`
			SuppressTypeName *bool   `tfsdk:"suppress_type_name" json:"suppressTypeName,omitempty"`
			TagKey           *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
			TimeKey          *string `tfsdk:"time_key" json:"timeKey,omitempty"`
			TimeKeyFormat    *string `tfsdk:"time_key_format" json:"timeKeyFormat,omitempty"`
			TimeKeyNanos     *bool   `tfsdk:"time_key_nanos" json:"timeKeyNanos,omitempty"`
			Tls              *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			TotalLimitSize *string `tfsdk:"total_limit_size" json:"totalLimitSize,omitempty"`
			TraceError     *bool   `tfsdk:"trace_error" json:"traceError,omitempty"`
			TraceOutput    *bool   `tfsdk:"trace_output" json:"traceOutput,omitempty"`
			Type           *string `tfsdk:"type" json:"type,omitempty"`
			WriteOperation *string `tfsdk:"write_operation" json:"writeOperation,omitempty"`
		} `tfsdk:"opensearch" json:"opensearch,omitempty"`
		Opentelemetry *struct {
			AddLabel     *map[string]string `tfsdk:"add_label" json:"addLabel,omitempty"`
			Header       *map[string]string `tfsdk:"header" json:"header,omitempty"`
			Host         *string            `tfsdk:"host" json:"host,omitempty"`
			HttpPassword *struct {
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
			LogResponsePayload *bool   `tfsdk:"log_response_payload" json:"logResponsePayload,omitempty"`
			LogsUri            *string `tfsdk:"logs_uri" json:"logsUri,omitempty"`
			MetricsUri         *string `tfsdk:"metrics_uri" json:"metricsUri,omitempty"`
			Networking         *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port  *int64  `tfsdk:"port" json:"port,omitempty"`
			Proxy *string `tfsdk:"proxy" json:"proxy,omitempty"`
			Tls   *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			TracesUri *string `tfsdk:"traces_uri" json:"tracesUri,omitempty"`
		} `tfsdk:"opentelemetry" json:"opentelemetry,omitempty"`
		Processors         *map[string]string `tfsdk:"processors" json:"processors,omitempty"`
		PrometheusExporter *struct {
			AddLabels *map[string]string `tfsdk:"add_labels" json:"addLabels,omitempty"`
			Host      *string            `tfsdk:"host" json:"host,omitempty"`
			Port      *int64             `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"prometheus_exporter" json:"prometheusExporter,omitempty"`
		PrometheusRemoteWrite *struct {
			AddLabels  *map[string]string `tfsdk:"add_labels" json:"addLabels,omitempty"`
			Headers    *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
			Host       *string            `tfsdk:"host" json:"host,omitempty"`
			HttpPasswd *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"http_passwd" json:"httpPasswd,omitempty"`
			HttpUser *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"http_user" json:"httpUser,omitempty"`
			LogResponsePayload *bool `tfsdk:"log_response_payload" json:"logResponsePayload,omitempty"`
			Networking         *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port  *int64  `tfsdk:"port" json:"port,omitempty"`
			Proxy *string `tfsdk:"proxy" json:"proxy,omitempty"`
			Tls   *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
			Workers *int64  `tfsdk:"workers" json:"workers,omitempty"`
		} `tfsdk:"prometheus_remote_write" json:"prometheusRemoteWrite,omitempty"`
		Retry_limit *string `tfsdk:"retry_limit" json:"retry_limit,omitempty"`
		S3          *struct {
			AutoRetryRequests        *bool   `tfsdk:"auto_retry_requests" json:"AutoRetryRequests,omitempty"`
			Bucket                   *string `tfsdk:"bucket" json:"Bucket,omitempty"`
			CannedAcl                *string `tfsdk:"canned_acl" json:"CannedAcl,omitempty"`
			Compression              *string `tfsdk:"compression" json:"Compression,omitempty"`
			ContentType              *string `tfsdk:"content_type" json:"ContentType,omitempty"`
			Endpoint                 *string `tfsdk:"endpoint" json:"Endpoint,omitempty"`
			ExternalId               *string `tfsdk:"external_id" json:"ExternalId,omitempty"`
			JsonDateFormat           *string `tfsdk:"json_date_format" json:"JsonDateFormat,omitempty"`
			JsonDateKey              *string `tfsdk:"json_date_key" json:"JsonDateKey,omitempty"`
			LogKey                   *string `tfsdk:"log_key" json:"LogKey,omitempty"`
			PreserveDataOrdering     *bool   `tfsdk:"preserve_data_ordering" json:"PreserveDataOrdering,omitempty"`
			Profile                  *string `tfsdk:"profile" json:"Profile,omitempty"`
			Region                   *string `tfsdk:"region" json:"Region,omitempty"`
			RetryLimit               *int64  `tfsdk:"retry_limit" json:"RetryLimit,omitempty"`
			RoleArn                  *string `tfsdk:"role_arn" json:"RoleArn,omitempty"`
			S3KeyFormat              *string `tfsdk:"s3_key_format" json:"S3KeyFormat,omitempty"`
			S3KeyFormatTagDelimiters *string `tfsdk:"s3_key_format_tag_delimiters" json:"S3KeyFormatTagDelimiters,omitempty"`
			SendContentMd5           *bool   `tfsdk:"send_content_md5" json:"SendContentMd5,omitempty"`
			StaticFilePath           *bool   `tfsdk:"static_file_path" json:"StaticFilePath,omitempty"`
			StorageClass             *string `tfsdk:"storage_class" json:"StorageClass,omitempty"`
			StoreDir                 *string `tfsdk:"store_dir" json:"StoreDir,omitempty"`
			StoreDirLimitSize        *string `tfsdk:"store_dir_limit_size" json:"StoreDirLimitSize,omitempty"`
			StsEndpoint              *string `tfsdk:"sts_endpoint" json:"StsEndpoint,omitempty"`
			TotalFileSize            *string `tfsdk:"total_file_size" json:"TotalFileSize,omitempty"`
			UploadChunkSize          *string `tfsdk:"upload_chunk_size" json:"UploadChunkSize,omitempty"`
			UploadTimeout            *string `tfsdk:"upload_timeout" json:"UploadTimeout,omitempty"`
			UsePutObject             *bool   `tfsdk:"use_put_object" json:"UsePutObject,omitempty"`
			Tls                      *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"s3" json:"s3,omitempty"`
		Splunk *struct {
			Workers             *int64    `tfsdk:"workers" json:"Workers,omitempty"`
			Channel             *string   `tfsdk:"channel" json:"channel,omitempty"`
			Compress            *string   `tfsdk:"compress" json:"compress,omitempty"`
			EventFields         *[]string `tfsdk:"event_fields" json:"eventFields,omitempty"`
			EventHost           *string   `tfsdk:"event_host" json:"eventHost,omitempty"`
			EventIndex          *string   `tfsdk:"event_index" json:"eventIndex,omitempty"`
			EventIndexKey       *string   `tfsdk:"event_index_key" json:"eventIndexKey,omitempty"`
			EventKey            *string   `tfsdk:"event_key" json:"eventKey,omitempty"`
			EventSource         *string   `tfsdk:"event_source" json:"eventSource,omitempty"`
			EventSourcetype     *string   `tfsdk:"event_sourcetype" json:"eventSourcetype,omitempty"`
			EventSourcetypeKey  *string   `tfsdk:"event_sourcetype_key" json:"eventSourcetypeKey,omitempty"`
			Host                *string   `tfsdk:"host" json:"host,omitempty"`
			HttpBufferSize      *string   `tfsdk:"http_buffer_size" json:"httpBufferSize,omitempty"`
			HttpDebugBadRequest *bool     `tfsdk:"http_debug_bad_request" json:"httpDebugBadRequest,omitempty"`
			HttpPassword        *struct {
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
			Networking *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port          *int64 `tfsdk:"port" json:"port,omitempty"`
			SplunkSendRaw *bool  `tfsdk:"splunk_send_raw" json:"splunkSendRaw,omitempty"`
			SplunkToken   *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"splunk_token" json:"splunkToken,omitempty"`
			Tls *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"splunk" json:"splunk,omitempty"`
		Stackdriver *struct {
			AutoformatStackdriverTrace *bool     `tfsdk:"autoformat_stackdriver_trace" json:"autoformatStackdriverTrace,omitempty"`
			CustomK8sRegex             *string   `tfsdk:"custom_k8s_regex" json:"customK8sRegex,omitempty"`
			ExportToProjectID          *string   `tfsdk:"export_to_project_id" json:"exportToProjectID,omitempty"`
			GoogleServiceCredentials   *string   `tfsdk:"google_service_credentials" json:"googleServiceCredentials,omitempty"`
			Job                        *string   `tfsdk:"job" json:"job,omitempty"`
			K8sClusterLocation         *string   `tfsdk:"k8s_cluster_location" json:"k8sClusterLocation,omitempty"`
			K8sClusterName             *string   `tfsdk:"k8s_cluster_name" json:"k8sClusterName,omitempty"`
			Labels                     *[]string `tfsdk:"labels" json:"labels,omitempty"`
			LabelsKey                  *string   `tfsdk:"labels_key" json:"labelsKey,omitempty"`
			Location                   *string   `tfsdk:"location" json:"location,omitempty"`
			LogNameKey                 *string   `tfsdk:"log_name_key" json:"logNameKey,omitempty"`
			MetadataServer             *string   `tfsdk:"metadata_server" json:"metadataServer,omitempty"`
			Namespace                  *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			NodeID                     *string   `tfsdk:"node_id" json:"nodeID,omitempty"`
			Resource                   *string   `tfsdk:"resource" json:"resource,omitempty"`
			ResourceLabels             *[]string `tfsdk:"resource_labels" json:"resourceLabels,omitempty"`
			ServiceAccountEmail        *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"service_account_email" json:"serviceAccountEmail,omitempty"`
			ServiceAccountSecret *struct {
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"service_account_secret" json:"serviceAccountSecret,omitempty"`
			SeverityKey *string `tfsdk:"severity_key" json:"severityKey,omitempty"`
			TagPrefix   *string `tfsdk:"tag_prefix" json:"tagPrefix,omitempty"`
			TaskID      *string `tfsdk:"task_id" json:"taskID,omitempty"`
			Workers     *int64  `tfsdk:"workers" json:"workers,omitempty"`
		} `tfsdk:"stackdriver" json:"stackdriver,omitempty"`
		Stdout *struct {
			Format         *string `tfsdk:"format" json:"format,omitempty"`
			JsonDateFormat *string `tfsdk:"json_date_format" json:"jsonDateFormat,omitempty"`
			JsonDateKey    *string `tfsdk:"json_date_key" json:"jsonDateKey,omitempty"`
		} `tfsdk:"stdout" json:"stdout,omitempty"`
		Syslog *struct {
			Host       *string `tfsdk:"host" json:"host,omitempty"`
			Mode       *string `tfsdk:"mode" json:"mode,omitempty"`
			Networking *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port               *int64  `tfsdk:"port" json:"port,omitempty"`
			SyslogAppnameKey   *string `tfsdk:"syslog_appname_key" json:"syslogAppnameKey,omitempty"`
			SyslogFacilityKey  *string `tfsdk:"syslog_facility_key" json:"syslogFacilityKey,omitempty"`
			SyslogFormat       *string `tfsdk:"syslog_format" json:"syslogFormat,omitempty"`
			SyslogHostnameKey  *string `tfsdk:"syslog_hostname_key" json:"syslogHostnameKey,omitempty"`
			SyslogMaxSize      *int64  `tfsdk:"syslog_max_size" json:"syslogMaxSize,omitempty"`
			SyslogMessageIDKey *string `tfsdk:"syslog_message_id_key" json:"syslogMessageIDKey,omitempty"`
			SyslogMessageKey   *string `tfsdk:"syslog_message_key" json:"syslogMessageKey,omitempty"`
			SyslogProcessIDKey *string `tfsdk:"syslog_process_id_key" json:"syslogProcessIDKey,omitempty"`
			SyslogSDKey        *string `tfsdk:"syslog_sd_key" json:"syslogSDKey,omitempty"`
			SyslogSeverityKey  *string `tfsdk:"syslog_severity_key" json:"syslogSeverityKey,omitempty"`
			Tls                *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"syslog" json:"syslog,omitempty"`
		Tcp *struct {
			Format         *string `tfsdk:"format" json:"format,omitempty"`
			Host           *string `tfsdk:"host" json:"host,omitempty"`
			JsonDateFormat *string `tfsdk:"json_date_format" json:"jsonDateFormat,omitempty"`
			JsonDateKey    *string `tfsdk:"json_date_key" json:"jsonDateKey,omitempty"`
			Networking     *struct {
				DNSMode                *string `tfsdk:"dns_mode" json:"DNSMode,omitempty"`
				DNSPreferIPv4          *bool   `tfsdk:"dns_prefer_i_pv4" json:"DNSPreferIPv4,omitempty"`
				DNSResolver            *string `tfsdk:"dns_resolver" json:"DNSResolver,omitempty"`
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				Keepalive              *string `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				MaxWorkerConnections   *int64  `tfsdk:"max_worker_connections" json:"maxWorkerConnections,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"networking" json:"networking,omitempty"`
			Port *int64 `tfsdk:"port" json:"port,omitempty"`
			Tls  *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"tcp" json:"tcp,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentbitFluentIoClusterOutputV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentbit_fluent_io_cluster_output_v1alpha2_manifest"
}

func (r *FluentbitFluentIoClusterOutputV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterOutput is the Schema for the cluster-level outputs API",
		MarkdownDescription: "ClusterOutput is the Schema for the cluster-level outputs API",
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
				Description:         "OutputSpec defines the desired state of ClusterOutput",
				MarkdownDescription: "OutputSpec defines the desired state of ClusterOutput",
				Attributes: map[string]schema.Attribute{
					"alias": schema.StringAttribute{
						Description:         "A user friendly alias name for this output plugin.Used in metrics for distinction of each configured output.",
						MarkdownDescription: "A user friendly alias name for this output plugin.Used in metrics for distinction of each configured output.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"azure_blob": schema.SingleNestedAttribute{
						Description:         "AzureBlob defines AzureBlob Output Configuration",
						MarkdownDescription: "AzureBlob defines AzureBlob Output Configuration",
						Attributes: map[string]schema.Attribute{
							"account_name": schema.StringAttribute{
								Description:         "Azure Storage account name",
								MarkdownDescription: "Azure Storage account name",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"auto_create_container": schema.StringAttribute{
								Description:         "Creates container if ContainerName is not set.",
								MarkdownDescription: "Creates container if ContainerName is not set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"blob_type": schema.StringAttribute{
								Description:         "Specify the desired blob type. Must be 'appendblob' or 'blockblob'",
								MarkdownDescription: "Specify the desired blob type. Must be 'appendblob' or 'blockblob'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("appendblob", "blockblob"),
								},
							},

							"container_name": schema.StringAttribute{
								Description:         "Name of the container that will contain the blobs",
								MarkdownDescription: "Name of the container that will contain the blobs",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"emulator_mode": schema.StringAttribute{
								Description:         "Optional toggle to use an Azure emulator",
								MarkdownDescription: "Optional toggle to use an Azure emulator",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"endpoint": schema.StringAttribute{
								Description:         "HTTP Service of the endpoint (if using EmulatorMode)",
								MarkdownDescription: "HTTP Service of the endpoint (if using EmulatorMode)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": schema.StringAttribute{
								Description:         "Optional path to store the blobs.",
								MarkdownDescription: "Optional path to store the blobs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"shared_key": schema.SingleNestedAttribute{
								Description:         "Specify the Azure Storage Shared Key to authenticate against the storage account",
								MarkdownDescription: "Specify the Azure Storage Shared Key to authenticate against the storage account",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Enable/Disable TLS Encryption. Azure services require TLS to be enabled.",
								MarkdownDescription: "Enable/Disable TLS Encryption. Azure services require TLS to be enabled.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"azure_log_analytics": schema.SingleNestedAttribute{
						Description:         "AzureLogAnalytics defines AzureLogAnalytics Output Configuration",
						MarkdownDescription: "AzureLogAnalytics defines AzureLogAnalytics Output Configuration",
						Attributes: map[string]schema.Attribute{
							"customer_id": schema.SingleNestedAttribute{
								Description:         "Customer ID or Workspace ID",
								MarkdownDescription: "Customer ID or Workspace ID",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"log_type": schema.StringAttribute{
								Description:         "Name of the event type.",
								MarkdownDescription: "Name of the event type.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"shared_key": schema.SingleNestedAttribute{
								Description:         "Specify the primary or the secondary client authentication key",
								MarkdownDescription: "Specify the primary or the secondary client authentication key",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"time_generated": schema.BoolAttribute{
								Description:         "If set, overrides the timeKey value with the 'time-generated-field' HTTP header value.",
								MarkdownDescription: "If set, overrides the timeKey value with the 'time-generated-field' HTTP header value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "Specify the name of the key where the timestamp is stored.",
								MarkdownDescription: "Specify the name of the key where the timestamp is stored.",
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
						Description:         "CloudWatch defines CloudWatch Output Configuration",
						MarkdownDescription: "CloudWatch defines CloudWatch Output Configuration",
						Attributes: map[string]schema.Attribute{
							"auto_create_group": schema.BoolAttribute{
								Description:         "Automatically create the log group. Defaults to False.",
								MarkdownDescription: "Automatically create the log group. Defaults to False.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auto_retry_requests": schema.BoolAttribute{
								Description:         "Automatically retry failed requests to CloudWatch once. Defaults to True.",
								MarkdownDescription: "Automatically retry failed requests to CloudWatch once. Defaults to True.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "Custom endpoint for CloudWatch logs API",
								MarkdownDescription: "Custom endpoint for CloudWatch logs API",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_id": schema.StringAttribute{
								Description:         "Specify an external ID for the STS API.",
								MarkdownDescription: "Specify an external ID for the STS API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "Optional parameter to tell CloudWatch the format of the data",
								MarkdownDescription: "Optional parameter to tell CloudWatch the format of the data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_group_name": schema.StringAttribute{
								Description:         "Name of Cloudwatch Log Group to send log records to",
								MarkdownDescription: "Name of Cloudwatch Log Group to send log records to",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_group_template": schema.StringAttribute{
								Description:         "Template for Log Group name, overrides LogGroupName if set.",
								MarkdownDescription: "Template for Log Group name, overrides LogGroupName if set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_key": schema.StringAttribute{
								Description:         "If set, only the value of the key will be sent to CloudWatch",
								MarkdownDescription: "If set, only the value of the key will be sent to CloudWatch",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_retention_days": schema.Int64Attribute{
								Description:         "Number of days logs are retained for",
								MarkdownDescription: "Number of days logs are retained for",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.OneOf(1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 3653),
								},
							},

							"log_stream_name": schema.StringAttribute{
								Description:         "The name of the CloudWatch Log Stream to send log records to",
								MarkdownDescription: "The name of the CloudWatch Log Stream to send log records to",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_stream_prefix": schema.StringAttribute{
								Description:         "Prefix for the Log Stream name. Not compatible with LogStreamName setting",
								MarkdownDescription: "Prefix for the Log Stream name. Not compatible with LogStreamName setting",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_stream_template": schema.StringAttribute{
								Description:         "Template for Log Stream name. Overrides LogStreamPrefix and LogStreamName if set.",
								MarkdownDescription: "Template for Log Stream name. Overrides LogStreamPrefix and LogStreamName if set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metric_dimensions": schema.StringAttribute{
								Description:         "Optional lists of lists for dimension keys to be added to all metrics. Use comma separated stringsfor one list of dimensions and semicolon separated strings for list of lists dimensions.",
								MarkdownDescription: "Optional lists of lists for dimension keys to be added to all metrics. Use comma separated stringsfor one list of dimensions and semicolon separated strings for list of lists dimensions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metric_namespace": schema.StringAttribute{
								Description:         "Optional string to represent the CloudWatch namespace.",
								MarkdownDescription: "Optional string to represent the CloudWatch namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "AWS Region",
								MarkdownDescription: "AWS Region",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"role_arn": schema.StringAttribute{
								Description:         "Role ARN to use for cross-account access",
								MarkdownDescription: "Role ARN to use for cross-account access",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sts_endpoint": schema.StringAttribute{
								Description:         "Specify a custom STS endpoint for the AWS STS API",
								MarkdownDescription: "Specify a custom STS endpoint for the AWS STS API",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_plugin": schema.SingleNestedAttribute{
						Description:         "CustomPlugin defines Custom Output configuration.",
						MarkdownDescription: "CustomPlugin defines Custom Output configuration.",
						Attributes: map[string]schema.Attribute{
							"config": schema.StringAttribute{
								Description:         "Config holds any unsupported plugins classic configurations,if ConfigFileFormat is set to yaml, this filed will be ignored",
								MarkdownDescription: "Config holds any unsupported plugins classic configurations,if ConfigFileFormat is set to yaml, this filed will be ignored",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"yaml_config": schema.MapAttribute{
								Description:         "YamlConfig holds the unsupported plugins yaml configurations, it only works when the ConfigFileFormat is yaml",
								MarkdownDescription: "YamlConfig holds the unsupported plugins yaml configurations, it only works when the ConfigFileFormat is yaml",
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

					"datadog": schema.SingleNestedAttribute{
						Description:         "DataDog defines DataDog Output configuration.",
						MarkdownDescription: "DataDog defines DataDog Output configuration.",
						Attributes: map[string]schema.Attribute{
							"apikey": schema.SingleNestedAttribute{
								Description:         "Your Datadog API key.",
								MarkdownDescription: "Your Datadog API key.",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"compress": schema.StringAttribute{
								Description:         "Compress  the payload in GZIP format.Datadog supports and recommends setting this to gzip.",
								MarkdownDescription: "Compress  the payload in GZIP format.Datadog supports and recommends setting this to gzip.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dd_message_key": schema.StringAttribute{
								Description:         "By default, the plugin searches for the key 'log' and remap the value to the key 'message'. If the property is set, the plugin will search the property name key.",
								MarkdownDescription: "By default, the plugin searches for the key 'log' and remap the value to the key 'message'. If the property is set, the plugin will search the property name key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dd_service": schema.StringAttribute{
								Description:         "The human readable name for your service generating the logs.",
								MarkdownDescription: "The human readable name for your service generating the logs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dd_source": schema.StringAttribute{
								Description:         "A human readable name for the underlying technology of your service.",
								MarkdownDescription: "A human readable name for the underlying technology of your service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dd_tags": schema.StringAttribute{
								Description:         "The tags you want to assign to your logs in Datadog.",
								MarkdownDescription: "The tags you want to assign to your logs in Datadog.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "Host is the Datadog server where you are sending your logs.",
								MarkdownDescription: "Host is the Datadog server where you are sending your logs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include_tag_key": schema.BoolAttribute{
								Description:         "If enabled, a tag is appended to output. The key name is used tag_key property.",
								MarkdownDescription: "If enabled, a tag is appended to output. The key name is used tag_key property.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"json_date_key": schema.StringAttribute{
								Description:         "Date key name for output.",
								MarkdownDescription: "Date key name for output.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provider": schema.StringAttribute{
								Description:         "To activate the remapping, specify configuration flag provider.",
								MarkdownDescription: "To activate the remapping, specify configuration flag provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy": schema.StringAttribute{
								Description:         "Specify an HTTP Proxy.",
								MarkdownDescription: "Specify an HTTP Proxy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_key": schema.StringAttribute{
								Description:         "The key name of tag. If include_tag_key is false, This property is ignored.",
								MarkdownDescription: "The key name of tag. If include_tag_key is false, This property is ignored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.BoolAttribute{
								Description:         "TLS controls whether to use end-to-end security communications security protocol.Datadog recommends setting this to on.",
								MarkdownDescription: "TLS controls whether to use end-to-end security communications security protocol.Datadog recommends setting this to on.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"es": schema.SingleNestedAttribute{
						Description:         "Elasticsearch defines Elasticsearch Output configuration.",
						MarkdownDescription: "Elasticsearch defines Elasticsearch Output configuration.",
						Attributes: map[string]schema.Attribute{
							"aws_auth": schema.StringAttribute{
								Description:         "Enable AWS Sigv4 Authentication for Amazon ElasticSearch Service.",
								MarkdownDescription: "Enable AWS Sigv4 Authentication for Amazon ElasticSearch Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_external_id": schema.StringAttribute{
								Description:         "External ID for the AWS IAM Role specified with aws_role_arn.",
								MarkdownDescription: "External ID for the AWS IAM Role specified with aws_role_arn.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_region": schema.StringAttribute{
								Description:         "Specify the AWS region for Amazon ElasticSearch Service.",
								MarkdownDescription: "Specify the AWS region for Amazon ElasticSearch Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_role_arn": schema.StringAttribute{
								Description:         "AWS IAM Role to assume to put records to your Amazon ES cluster.",
								MarkdownDescription: "AWS IAM Role to assume to put records to your Amazon ES cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_sts_endpoint": schema.StringAttribute{
								Description:         "Specify the custom sts endpoint to be used with STS API for Amazon ElasticSearch Service.",
								MarkdownDescription: "Specify the custom sts endpoint to be used with STS API for Amazon ElasticSearch Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"buffer_size": schema.StringAttribute{
								Description:         "Specify the buffer size used to read the response from the Elasticsearch HTTP service.This option is useful for debugging purposes where is required to read full responses,note that response size grows depending of the number of records inserted.To set an unlimited amount of memory set this value to False,otherwise the value must be according to the Unit Size specification.",
								MarkdownDescription: "Specify the buffer size used to read the response from the Elasticsearch HTTP service.This option is useful for debugging purposes where is required to read full responses,note that response size grows depending of the number of records inserted.To set an unlimited amount of memory set this value to False,otherwise the value must be according to the Unit Size specification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"cloud_auth": schema.StringAttribute{
								Description:         "Specify the credentials to use to connect to Elastic's Elasticsearch Service running on Elastic Cloud.",
								MarkdownDescription: "Specify the credentials to use to connect to Elastic's Elasticsearch Service running on Elastic Cloud.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cloud_id": schema.StringAttribute{
								Description:         "If you are using Elastic's Elasticsearch Service you can specify the cloud_id of the cluster running.",
								MarkdownDescription: "If you are using Elastic's Elasticsearch Service you can specify the cloud_id of the cluster running.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compress": schema.StringAttribute{
								Description:         "Set payload compression mechanism. Option available is 'gzip'",
								MarkdownDescription: "Set payload compression mechanism. Option available is 'gzip'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("gzip"),
								},
							},

							"current_time_index": schema.BoolAttribute{
								Description:         "Use current time for index generation instead of message record",
								MarkdownDescription: "Use current time for index generation instead of message record",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"generate_id": schema.BoolAttribute{
								Description:         "When enabled, generate _id for outgoing records.This prevents duplicate records when retrying ES.",
								MarkdownDescription: "When enabled, generate _id for outgoing records.This prevents duplicate records when retrying ES.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target Elasticsearch instance",
								MarkdownDescription: "IP address or hostname of the target Elasticsearch instance",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Password for user defined in HTTP_User",
								MarkdownDescription: "Password for user defined in HTTP_User",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Optional username credential for Elastic X-Pack access",
								MarkdownDescription: "Optional username credential for Elastic X-Pack access",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"id_key": schema.StringAttribute{
								Description:         "If set, _id will be the value of the key from incoming record and Generate_ID option is ignored.",
								MarkdownDescription: "If set, _id will be the value of the key from incoming record and Generate_ID option is ignored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include_tag_key": schema.BoolAttribute{
								Description:         "When enabled, it append the Tag name to the record.",
								MarkdownDescription: "When enabled, it append the Tag name to the record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"index": schema.StringAttribute{
								Description:         "Index name",
								MarkdownDescription: "Index name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_date_format": schema.StringAttribute{
								Description:         "Time format (based on strftime) to generate the second part of the Index name.",
								MarkdownDescription: "Time format (based on strftime) to generate the second part of the Index name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_format": schema.BoolAttribute{
								Description:         "Enable Logstash format compatibility.This option takes a boolean value: True/False, On/Off",
								MarkdownDescription: "Enable Logstash format compatibility.This option takes a boolean value: True/False, On/Off",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_prefix": schema.StringAttribute{
								Description:         "When Logstash_Format is enabled, the Index name is composed using a prefix and the date,e.g: If Logstash_Prefix is equals to 'mydata' your index will become 'mydata-YYYY.MM.DD'.The last string appended belongs to the date when the data is being generated.",
								MarkdownDescription: "When Logstash_Format is enabled, the Index name is composed using a prefix and the date,e.g: If Logstash_Prefix is equals to 'mydata' your index will become 'mydata-YYYY.MM.DD'.The last string appended belongs to the date when the data is being generated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_prefix_key": schema.StringAttribute{
								Description:         "Prefix keys with this string",
								MarkdownDescription: "Prefix keys with this string",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": schema.StringAttribute{
								Description:         "Elasticsearch accepts new data on HTTP query path '/_bulk'.But it is also possible to serve Elasticsearch behind a reverse proxy on a subpath.This option defines such path on the fluent-bit side.It simply adds a path prefix in the indexing HTTP POST URI.",
								MarkdownDescription: "Elasticsearch accepts new data on HTTP query path '/_bulk'.But it is also possible to serve Elasticsearch behind a reverse proxy on a subpath.This option defines such path on the fluent-bit side.It simply adds a path prefix in the indexing HTTP POST URI.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pipeline": schema.StringAttribute{
								Description:         "Newer versions of Elasticsearch allows setting up filters called pipelines.This option allows defining which pipeline the database should use.For performance reasons is strongly suggested parsingand filtering on Fluent Bit side, avoid pipelines.",
								MarkdownDescription: "Newer versions of Elasticsearch allows setting up filters called pipelines.This option allows defining which pipeline the database should use.For performance reasons is strongly suggested parsingand filtering on Fluent Bit side, avoid pipelines.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target Elasticsearch instance",
								MarkdownDescription: "TCP port of the target Elasticsearch instance",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"replace_dots": schema.BoolAttribute{
								Description:         "When enabled, replace field name dots with underscore, required by Elasticsearch 2.0-2.3.",
								MarkdownDescription: "When enabled, replace field name dots with underscore, required by Elasticsearch 2.0-2.3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"suppress_type_name": schema.StringAttribute{
								Description:         "When enabled, mapping types is removed and Type option is ignored. Types are deprecated in APIs in v7.0. This options is for v7.0 or later.",
								MarkdownDescription: "When enabled, mapping types is removed and Type option is ignored. Types are deprecated in APIs in v7.0. This options is for v7.0 or later.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_key": schema.StringAttribute{
								Description:         "When Include_Tag_Key is enabled, this property defines the key name for the tag.",
								MarkdownDescription: "When Include_Tag_Key is enabled, this property defines the key name for the tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "When Logstash_Format is enabled, each record will get a new timestamp field.The Time_Key property defines the name of that field.",
								MarkdownDescription: "When Logstash_Format is enabled, each record will get a new timestamp field.The Time_Key property defines the name of that field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key_format": schema.StringAttribute{
								Description:         "When Logstash_Format is enabled, this property defines the format of the timestamp.",
								MarkdownDescription: "When Logstash_Format is enabled, this property defines the format of the timestamp.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key_nanos": schema.BoolAttribute{
								Description:         "When Logstash_Format is enabled, enabling this property sends nanosecond precision timestamps.",
								MarkdownDescription: "When Logstash_Format is enabled, enabling this property sends nanosecond precision timestamps.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"total_limit_size": schema.StringAttribute{
								Description:         "Limit the maximum number of Chunks in the filesystem for the current output logical destination.",
								MarkdownDescription: "Limit the maximum number of Chunks in the filesystem for the current output logical destination.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trace_error": schema.BoolAttribute{
								Description:         "When enabled print the elasticsearch API calls to stdout when elasticsearch returns an error",
								MarkdownDescription: "When enabled print the elasticsearch API calls to stdout when elasticsearch returns an error",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trace_output": schema.BoolAttribute{
								Description:         "When enabled print the elasticsearch API calls to stdout (for diag only)",
								MarkdownDescription: "When enabled print the elasticsearch API calls to stdout (for diag only)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type name",
								MarkdownDescription: "Type name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_operation": schema.StringAttribute{
								Description:         "Operation to use to write in bulk requests.",
								MarkdownDescription: "Operation to use to write in bulk requests.",
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
						Description:         "File defines File Output configuration.",
						MarkdownDescription: "File defines File Output configuration.",
						Attributes: map[string]schema.Attribute{
							"delimiter": schema.StringAttribute{
								Description:         "The character to separate each pair. Applicable only if format is csv or ltsv.",
								MarkdownDescription: "The character to separate each pair. Applicable only if format is csv or ltsv.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"file": schema.StringAttribute{
								Description:         "Set file name to store the records. If not set, the file name will be the tag associated with the records.",
								MarkdownDescription: "Set file name to store the records. If not set, the file name will be the tag associated with the records.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "The format of the file content. See also Format section. Default: out_file.",
								MarkdownDescription: "The format of the file content. See also Format section. Default: out_file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("out_file", "plain", "csv", "ltsv", "template"),
								},
							},

							"label_delimiter": schema.StringAttribute{
								Description:         "The character to separate each pair. Applicable only if format is ltsv.",
								MarkdownDescription: "The character to separate each pair. Applicable only if format is ltsv.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Absolute directory path to store files. If not set, Fluent Bit will write the files on it's own positioned directory.",
								MarkdownDescription: "Absolute directory path to store files. If not set, Fluent Bit will write the files on it's own positioned directory.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"template": schema.StringAttribute{
								Description:         "The format string. Applicable only if format is template.",
								MarkdownDescription: "The format string. Applicable only if format is template.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"firehose": schema.SingleNestedAttribute{
						Description:         "Firehose defines Firehose Output configuration.",
						MarkdownDescription: "Firehose defines Firehose Output configuration.",
						Attributes: map[string]schema.Attribute{
							"auto_retry_requests": schema.BoolAttribute{
								Description:         "Immediately retry failed requests to AWS services once. This option does not affect the normal Fluent Bit retry mechanism with backoff. Instead, it enables an immediate retry with no delay for networking errors, which may help improve throughput when there are transient/random networking issues.",
								MarkdownDescription: "Immediately retry failed requests to AWS services once. This option does not affect the normal Fluent Bit retry mechanism with backoff. Instead, it enables an immediate retry with no delay for networking errors, which may help improve throughput when there are transient/random networking issues.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data_keys": schema.StringAttribute{
								Description:         "By default, the whole log record will be sent to Kinesis. If you specify a key name(s) with this option, then only those keys and values will be sent to Kinesis. For example, if you are using the Fluentd Docker log driver, you can specify data_keys log and only the log message will be sent to Kinesis. If you specify multiple keys, they should be comma delimited.",
								MarkdownDescription: "By default, the whole log record will be sent to Kinesis. If you specify a key name(s) with this option, then only those keys and values will be sent to Kinesis. For example, if you are using the Fluentd Docker log driver, you can specify data_keys log and only the log message will be sent to Kinesis. If you specify multiple keys, they should be comma delimited.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delivery_stream": schema.StringAttribute{
								Description:         "The name of the Kinesis Firehose Delivery stream that you want log records sent to.",
								MarkdownDescription: "The name of the Kinesis Firehose Delivery stream that you want log records sent to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "Specify a custom endpoint for the Kinesis Firehose API.",
								MarkdownDescription: "Specify a custom endpoint for the Kinesis Firehose API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_key": schema.StringAttribute{
								Description:         "By default, the whole log record will be sent to Firehose. If you specify a key name with this option, then only the value of that key will be sent to Firehose. For example, if you are using the Fluentd Docker log driver, you can specify log_key log and only the log message will be sent to Firehose.",
								MarkdownDescription: "By default, the whole log record will be sent to Firehose. If you specify a key name with this option, then only the value of that key will be sent to Firehose. For example, if you are using the Fluentd Docker log driver, you can specify log_key log and only the log message will be sent to Firehose.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "The AWS region.",
								MarkdownDescription: "The AWS region.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"role_arn": schema.StringAttribute{
								Description:         "ARN of an IAM role to assume (for cross account access).",
								MarkdownDescription: "ARN of an IAM role to assume (for cross account access).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sts_endpoint": schema.StringAttribute{
								Description:         "Specify a custom endpoint for the STS API; used to assume your custom role provided with role_arn.",
								MarkdownDescription: "Specify a custom endpoint for the STS API; used to assume your custom role provided with role_arn.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "Add the timestamp to the record under this key. By default, the timestamp from Fluent Bit will not be added to records sent to Kinesis.",
								MarkdownDescription: "Add the timestamp to the record under this key. By default, the timestamp from Fluent Bit will not be added to records sent to Kinesis.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key_format": schema.StringAttribute{
								Description:         "strftime compliant format string for the timestamp; for example, %Y-%m-%dT%H *string This option is used with time_key. You can also use %L for milliseconds and %f for microseconds. If you are using ECS FireLens, make sure you are running Amazon ECS Container Agent v1.42.0 or later, otherwise the timestamps associated with your container logs will only have second precision.",
								MarkdownDescription: "strftime compliant format string for the timestamp; for example, %Y-%m-%dT%H *string This option is used with time_key. You can also use %L for milliseconds and %f for microseconds. If you are using ECS FireLens, make sure you are running Amazon ECS Container Agent v1.42.0 or later, otherwise the timestamps associated with your container logs will only have second precision.",
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
						Description:         "Forward defines Forward Output configuration.",
						MarkdownDescription: "Forward defines Forward Output configuration.",
						Attributes: map[string]schema.Attribute{
							"empty_shared_key": schema.BoolAttribute{
								Description:         "Use this option to connect to Fluentd with a zero-length secret.",
								MarkdownDescription: "Use this option to connect to Fluentd with a zero-length secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "Target host where Fluent-Bit or Fluentd are listening for Forward messages.",
								MarkdownDescription: "Target host where Fluent-Bit or Fluentd are listening for Forward messages.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": schema.SingleNestedAttribute{
								Description:         "Specify the password corresponding to the username.",
								MarkdownDescription: "Specify the password corresponding to the username.",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"port": schema.Int64Attribute{
								Description:         "TCP Port of the target service.",
								MarkdownDescription: "TCP Port of the target service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"require_ack_response": schema.BoolAttribute{
								Description:         "Send 'chunk'-option and wait for 'ack' response from server.Enables at-least-once and receiving server can control rate of traffic.(Requires Fluentd v0.14.0+ server)",
								MarkdownDescription: "Send 'chunk'-option and wait for 'ack' response from server.Enables at-least-once and receiving server can control rate of traffic.(Requires Fluentd v0.14.0+ server)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"self_hostname": schema.StringAttribute{
								Description:         "Default value of the auto-generated certificate common name (CN).",
								MarkdownDescription: "Default value of the auto-generated certificate common name (CN).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"send_options": schema.BoolAttribute{
								Description:         "Always send options (with 'size'=count of messages)",
								MarkdownDescription: "Always send options (with 'size'=count of messages)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"shared_key": schema.StringAttribute{
								Description:         "A key string known by the remote Fluentd used for authorization.",
								MarkdownDescription: "A key string known by the remote Fluentd used for authorization.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Overwrite the tag as we transmit. This allows the receiving pipeline startfresh, or to attribute source.",
								MarkdownDescription: "Overwrite the tag as we transmit. This allows the receiving pipeline startfresh, or to attribute source.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_as_integer": schema.BoolAttribute{
								Description:         "Set timestamps in integer format, it enable compatibility mode for Fluentd v0.12 series.",
								MarkdownDescription: "Set timestamps in integer format, it enable compatibility mode for Fluentd v0.12 series.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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
								Description:         "Specify the username to present to a Fluentd server that enables user_auth.",
								MarkdownDescription: "Specify the username to present to a Fluentd server that enables user_auth.",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

					"gelf": schema.SingleNestedAttribute{
						Description:         "Gelf defines GELF Output configuration.",
						MarkdownDescription: "Gelf defines GELF Output configuration.",
						Attributes: map[string]schema.Attribute{
							"compress": schema.BoolAttribute{
								Description:         "If transport protocol is udp, it defines if UDP packets should be compressed.",
								MarkdownDescription: "If transport protocol is udp, it defines if UDP packets should be compressed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"full_message_key": schema.StringAttribute{
								Description:         "FullMessageKey is the key to use as the long message that can i.e. contain a backtrace.",
								MarkdownDescription: "FullMessageKey is the key to use as the long message that can i.e. contain a backtrace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target Graylog server.",
								MarkdownDescription: "IP address or hostname of the target Graylog server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_key": schema.StringAttribute{
								Description:         "HostKey is the key which its value is used as the name of the host, source or application that sent this message.",
								MarkdownDescription: "HostKey is the key which its value is used as the name of the host, source or application that sent this message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"level_key": schema.StringAttribute{
								Description:         "LevelKey is the key to be used as the log level.",
								MarkdownDescription: "LevelKey is the key to be used as the log level.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "The protocol to use (tls, tcp or udp).",
								MarkdownDescription: "The protocol to use (tls, tcp or udp).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("tls", "tcp", "udp"),
								},
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"packet_size": schema.Int64Attribute{
								Description:         "If transport protocol is udp, it sets the size of packets to be sent.",
								MarkdownDescription: "If transport protocol is udp, it sets the size of packets to be sent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port that the target Graylog server is listening on.",
								MarkdownDescription: "The port that the target Graylog server is listening on.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"short_message_key": schema.StringAttribute{
								Description:         "ShortMessageKey is the key to use as the short message.",
								MarkdownDescription: "ShortMessageKey is the key to use as the short message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timestamp_key": schema.StringAttribute{
								Description:         "TimestampKey is the key which its value is used as the timestamp of the message.",
								MarkdownDescription: "TimestampKey is the key which its value is used as the timestamp of the message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"http": schema.SingleNestedAttribute{
						Description:         "HTTP defines HTTP Output configuration.",
						MarkdownDescription: "HTTP defines HTTP Output configuration.",
						Attributes: map[string]schema.Attribute{
							"allow_duplicated_headers": schema.BoolAttribute{
								Description:         "Specify if duplicated headers are allowed.If a duplicated header is found, the latest key/value set is preserved.",
								MarkdownDescription: "Specify if duplicated headers are allowed.If a duplicated header is found, the latest key/value set is preserved.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compress": schema.StringAttribute{
								Description:         "Set payload compression mechanism. Option available is 'gzip'",
								MarkdownDescription: "Set payload compression mechanism. Option available is 'gzip'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "Specify the data format to be used in the HTTP request body, by default it uses msgpack.Other supported formats are json, json_stream and json_lines and gelf.",
								MarkdownDescription: "Specify the data format to be used in the HTTP request body, by default it uses msgpack.Other supported formats are json, json_stream and json_lines and gelf.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("msgpack", "json", "json_stream", "json_lines", "gelf"),
								},
							},

							"gelf_full_message_key": schema.StringAttribute{
								Description:         "Specify the key to use for the full message in gelf format",
								MarkdownDescription: "Specify the key to use for the full message in gelf format",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gelf_host_key": schema.StringAttribute{
								Description:         "Specify the key to use for the host in gelf format",
								MarkdownDescription: "Specify the key to use for the host in gelf format",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gelf_level_key": schema.StringAttribute{
								Description:         "Specify the key to use for the level in gelf format",
								MarkdownDescription: "Specify the key to use for the level in gelf format",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gelf_short_message_key": schema.StringAttribute{
								Description:         "Specify the key to use as the short message in gelf format",
								MarkdownDescription: "Specify the key to use as the short message in gelf format",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gelf_timestamp_key": schema.StringAttribute{
								Description:         "Specify the key to use for timestamp in gelf format",
								MarkdownDescription: "Specify the key to use for timestamp in gelf format",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"header_tag": schema.StringAttribute{
								Description:         "Specify an optional HTTP header field for the original message tag.",
								MarkdownDescription: "Specify an optional HTTP header field for the original message tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"headers": schema.MapAttribute{
								Description:         "Add a HTTP header key/value pair. Multiple headers can be set.",
								MarkdownDescription: "Add a HTTP header key/value pair. Multiple headers can be set.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target HTTP Server",
								MarkdownDescription: "IP address or hostname of the target HTTP Server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Basic Auth Password. Requires HTTP_User to be set",
								MarkdownDescription: "Basic Auth Password. Requires HTTP_User to be set",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Basic Auth Username",
								MarkdownDescription: "Basic Auth Username",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"json_date_format": schema.StringAttribute{
								Description:         "Specify the format of the date. Supported formats are double, epochand iso8601 (eg: 2018-05-30T09:39:52.000681Z)",
								MarkdownDescription: "Specify the format of the date. Supported formats are double, epochand iso8601 (eg: 2018-05-30T09:39:52.000681Z)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"json_date_key": schema.StringAttribute{
								Description:         "Specify the name of the time key in the output record.To disable the time key just set the value to false.",
								MarkdownDescription: "Specify the name of the time key in the output record.To disable the time key just set the value to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target HTTP Server",
								MarkdownDescription: "TCP port of the target HTTP Server",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"proxy": schema.StringAttribute{
								Description:         "Specify an HTTP Proxy. The expected format of this value is http://host:port.Note that https is not supported yet.",
								MarkdownDescription: "Specify an HTTP Proxy. The expected format of this value is http://host:port.Note that https is not supported yet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "HTTP output plugin supports TTL/SSL, for more details about the properties availableand general configuration, please refer to the TLS/SSL section.",
								MarkdownDescription: "HTTP output plugin supports TTL/SSL, for more details about the properties availableand general configuration, please refer to the TLS/SSL section.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"uri": schema.StringAttribute{
								Description:         "Specify an optional HTTP URI for the target web server, e.g: /something",
								MarkdownDescription: "Specify an optional HTTP URI for the target web server, e.g: /something",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"influx_db": schema.SingleNestedAttribute{
						Description:         "InfluxDB defines InfluxDB Output configuration.",
						MarkdownDescription: "InfluxDB defines InfluxDB Output configuration.",
						Attributes: map[string]schema.Attribute{
							"auto_tags": schema.BoolAttribute{
								Description:         "Automatically tag keys where value is string.",
								MarkdownDescription: "Automatically tag keys where value is string.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bucket": schema.StringAttribute{
								Description:         "InfluxDB bucket name where records will be inserted - if specified, database is ignored and v2 of API is used",
								MarkdownDescription: "InfluxDB bucket name where records will be inserted - if specified, database is ignored and v2 of API is used",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"database": schema.StringAttribute{
								Description:         "InfluxDB database name where records will be inserted.",
								MarkdownDescription: "InfluxDB database name where records will be inserted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target InfluxDB service.",
								MarkdownDescription: "IP address or hostname of the target InfluxDB service.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Password for user defined in HTTP_User",
								MarkdownDescription: "Password for user defined in HTTP_User",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"http_token": schema.SingleNestedAttribute{
								Description:         "Authentication token used with InfluxDB v2 - if specified, both HTTPUser and HTTPPasswd are ignored",
								MarkdownDescription: "Authentication token used with InfluxDB v2 - if specified, both HTTPUser and HTTPPasswd are ignored",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Optional username for HTTP Basic Authentication",
								MarkdownDescription: "Optional username for HTTP Basic Authentication",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"org": schema.StringAttribute{
								Description:         "InfluxDB organization name where the bucket is (v2 only)",
								MarkdownDescription: "InfluxDB organization name where the bucket is (v2 only)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target InfluxDB service.",
								MarkdownDescription: "TCP port of the target InfluxDB service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(65536),
								},
							},

							"sequence_tag": schema.StringAttribute{
								Description:         "The name of the tag whose value is incremented for the consecutive simultaneous events.",
								MarkdownDescription: "The name of the tag whose value is incremented for the consecutive simultaneous events.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_keys": schema.ListAttribute{
								Description:         "List of keys that needs to be tagged",
								MarkdownDescription: "List of keys that needs to be tagged",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_list_key": schema.StringAttribute{
								Description:         "Key of the string array optionally contained within each log record that contains tag keys for that record",
								MarkdownDescription: "Key of the string array optionally contained within each log record that contains tag keys for that record",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags_list_enabled": schema.BoolAttribute{
								Description:         "Dynamically tag keys which are in the string array at Tags_List_Key key.",
								MarkdownDescription: "Dynamically tag keys which are in the string array at Tags_List_Key key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"kafka": schema.SingleNestedAttribute{
						Description:         "Kafka defines Kafka Output configuration.",
						MarkdownDescription: "Kafka defines Kafka Output configuration.",
						Attributes: map[string]schema.Attribute{
							"brokers": schema.StringAttribute{
								Description:         "Single of multiple list of Kafka Brokers, e.g: 192.168.1.3:9092, 192.168.1.4:9092.",
								MarkdownDescription: "Single of multiple list of Kafka Brokers, e.g: 192.168.1.3:9092, 192.168.1.4:9092.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dynamic_topic": schema.BoolAttribute{
								Description:         "adds unknown topics (found in Topic_Key) to Topics. So in Topics only a default topic needs to be configured",
								MarkdownDescription: "adds unknown topics (found in Topic_Key) to Topics. So in Topics only a default topic needs to be configured",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "Specify data format, options available: json, msgpack.",
								MarkdownDescription: "Specify data format, options available: json, msgpack.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"message_key": schema.StringAttribute{
								Description:         "Optional key to store the message",
								MarkdownDescription: "Optional key to store the message",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"message_key_field": schema.StringAttribute{
								Description:         "If set, the value of Message_Key_Field in the record will indicate the message key.If not set nor found in the record, Message_Key will be used (if set).",
								MarkdownDescription: "If set, the value of Message_Key_Field in the record will indicate the message key.If not set nor found in the record, Message_Key will be used (if set).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"queue_full_retries": schema.Int64Attribute{
								Description:         "Fluent Bit queues data into rdkafka library,if for some reason the underlying library cannot flush the records the queue might fills up blocking new addition of records.The queue_full_retries option set the number of local retries to enqueue the data.The default value is 10 times, the interval between each retry is 1 second.Setting the queue_full_retries value to 0 set's an unlimited number of retries.",
								MarkdownDescription: "Fluent Bit queues data into rdkafka library,if for some reason the underlying library cannot flush the records the queue might fills up blocking new addition of records.The queue_full_retries option set the number of local retries to enqueue the data.The default value is 10 times, the interval between each retry is 1 second.Setting the queue_full_retries value to 0 set's an unlimited number of retries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rdkafka": schema.MapAttribute{
								Description:         "{property} can be any librdkafka properties",
								MarkdownDescription: "{property} can be any librdkafka properties",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timestamp_format": schema.StringAttribute{
								Description:         "iso8601 or double",
								MarkdownDescription: "iso8601 or double",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timestamp_key": schema.StringAttribute{
								Description:         "Set the key to store the record timestamp",
								MarkdownDescription: "Set the key to store the record timestamp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"topic_key": schema.StringAttribute{
								Description:         "If multiple Topics exists, the value of Topic_Key in the record will indicate the topic to use.E.g: if Topic_Key is router and the record is {'key1': 123, 'router': 'route_2'},Fluent Bit will use topic route_2. Note that if the value of Topic_Key is not present in Topics,then by default the first topic in the Topics list will indicate the topic to be used.",
								MarkdownDescription: "If multiple Topics exists, the value of Topic_Key in the record will indicate the topic to use.E.g: if Topic_Key is router and the record is {'key1': 123, 'router': 'route_2'},Fluent Bit will use topic route_2. Note that if the value of Topic_Key is not present in Topics,then by default the first topic in the Topics list will indicate the topic to be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"topics": schema.StringAttribute{
								Description:         "Single entry or list of topics separated by comma (,) that Fluent Bit will use to send messages to Kafka.If only one topic is set, that one will be used for all records.Instead if multiple topics exists, the one set in the record by Topic_Key will be used.",
								MarkdownDescription: "Single entry or list of topics separated by comma (,) that Fluent Bit will use to send messages to Kafka.If only one topic is set, that one will be used for all records.Instead if multiple topics exists, the one set in the record by Topic_Key will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kinesis": schema.SingleNestedAttribute{
						Description:         "Kinesis defines Kinesis Output configuration.",
						MarkdownDescription: "Kinesis defines Kinesis Output configuration.",
						Attributes: map[string]schema.Attribute{
							"auto_retry_requests": schema.BoolAttribute{
								Description:         "Immediately retry failed requests to AWS services once. This option does not affect the normal Fluent Bit retry mechanism with backoff. Instead, it enables an immediate retry with no delay for networking errors, which may help improve throughput when there are transient/random networking issues. This option defaults to true.",
								MarkdownDescription: "Immediately retry failed requests to AWS services once. This option does not affect the normal Fluent Bit retry mechanism with backoff. Instead, it enables an immediate retry with no delay for networking errors, which may help improve throughput when there are transient/random networking issues. This option defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "Specify a custom endpoint for the Kinesis API.",
								MarkdownDescription: "Specify a custom endpoint for the Kinesis API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_id": schema.StringAttribute{
								Description:         "Specify an external ID for the STS API, can be used with the role_arn parameter if your role requires an external ID.",
								MarkdownDescription: "Specify an external ID for the STS API, can be used with the role_arn parameter if your role requires an external ID.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_key": schema.StringAttribute{
								Description:         "By default, the whole log record will be sent to Kinesis. If you specify a key name with this option, then only the value of that key will be sent to Kinesis. For example, if you are using the Fluentd Docker log driver, you can specify log_key log and only the log message will be sent to Kinesis.",
								MarkdownDescription: "By default, the whole log record will be sent to Kinesis. If you specify a key name with this option, then only the value of that key will be sent to Kinesis. For example, if you are using the Fluentd Docker log driver, you can specify log_key log and only the log message will be sent to Kinesis.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "The AWS region.",
								MarkdownDescription: "The AWS region.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"role_arn": schema.StringAttribute{
								Description:         "ARN of an IAM role to assume (for cross account access).",
								MarkdownDescription: "ARN of an IAM role to assume (for cross account access).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stream": schema.StringAttribute{
								Description:         "The name of the Kinesis Streams Delivery stream that you want log records sent to.",
								MarkdownDescription: "The name of the Kinesis Streams Delivery stream that you want log records sent to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"sts_endpoint": schema.StringAttribute{
								Description:         "Custom endpoint for the STS API.",
								MarkdownDescription: "Custom endpoint for the STS API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "Add the timestamp to the record under this key. By default the timestamp from Fluent Bit will not be added to records sent to Kinesis.",
								MarkdownDescription: "Add the timestamp to the record under this key. By default the timestamp from Fluent Bit will not be added to records sent to Kinesis.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key_format": schema.StringAttribute{
								Description:         "strftime compliant format string for the timestamp; for example, the default is '%Y-%m-%dT%H:%M:%S'. Supports millisecond precision with '%3N' and supports nanosecond precision with '%9N' and '%L'; for example, adding '%3N' to support millisecond '%Y-%m-%dT%H:%M:%S.%3N'. This option is used with time_key.",
								MarkdownDescription: "strftime compliant format string for the timestamp; for example, the default is '%Y-%m-%dT%H:%M:%S'. Supports millisecond precision with '%3N' and supports nanosecond precision with '%9N' and '%L'; for example, adding '%3N' to support millisecond '%Y-%m-%dT%H:%M:%S.%3N'. This option is used with time_key.",
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
						Description:         "Set the plugin's logging verbosity level. Allowed values are: off, error, warn, info, debug and trace, Defaults to the SERVICE section's Log_Level",
						MarkdownDescription: "Set the plugin's logging verbosity level. Allowed values are: off, error, warn, info, debug and trace, Defaults to the SERVICE section's Log_Level",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("off", "error", "warning", "info", "debug", "trace"),
						},
					},

					"loki": schema.SingleNestedAttribute{
						Description:         "Loki defines Loki Output configuration.",
						MarkdownDescription: "Loki defines Loki Output configuration.",
						Attributes: map[string]schema.Attribute{
							"auto_kubernetes_labels": schema.StringAttribute{
								Description:         "If set to true, it will add all Kubernetes labels to the Stream labels.",
								MarkdownDescription: "If set to true, it will add all Kubernetes labels to the Stream labels.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"bearer_token": schema.SingleNestedAttribute{
								Description:         "Set bearer token authentication token value.Can be used as alterntative to HTTP basic authentication",
								MarkdownDescription: "Set bearer token authentication token value.Can be used as alterntative to HTTP basic authentication",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"drop_single_key": schema.StringAttribute{
								Description:         "If set to true and after extracting labels only a single key remains, the log line sent to Loki will be the value of that key in line_format.",
								MarkdownDescription: "If set to true and after extracting labels only a single key remains, the log line sent to Loki will be the value of that key in line_format.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"host": schema.StringAttribute{
								Description:         "Loki hostname or IP address.",
								MarkdownDescription: "Loki hostname or IP address.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Password for user defined in HTTP_UserSet HTTP basic authentication password",
								MarkdownDescription: "Password for user defined in HTTP_UserSet HTTP basic authentication password",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"label_keys": schema.ListAttribute{
								Description:         "Optional list of record keys that will be placed as stream labels.This configuration property is for records key only.",
								MarkdownDescription: "Optional list of record keys that will be placed as stream labels.This configuration property is for records key only.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label_map_path": schema.StringAttribute{
								Description:         "Specify the label map file path. The file defines how to extract labels from each record.",
								MarkdownDescription: "Specify the label map file path. The file defines how to extract labels from each record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.ListAttribute{
								Description:         "Stream labels for API request. It can be multiple comma separated of strings specifying  key=value pairs.In addition to fixed parameters, it also allows to add custom record keys (similar to label_keys property).",
								MarkdownDescription: "Stream labels for API request. It can be multiple comma separated of strings specifying  key=value pairs.In addition to fixed parameters, it also allows to add custom record keys (similar to label_keys property).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"line_format": schema.StringAttribute{
								Description:         "Format to use when flattening the record to a log line. Valid values are json or key_value.If set to json,  the log line sent to Loki will be the Fluent Bit record dumped as JSON.If set to key_value, the log line will be each item in the record concatenated together (separated by a single space) in the format.",
								MarkdownDescription: "Format to use when flattening the record to a log line. Valid values are json or key_value.If set to json,  the log line sent to Loki will be the Fluent Bit record dumped as JSON.If set to key_value, the log line will be each item in the record concatenated together (separated by a single space) in the format.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("json", "key_value"),
								},
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "Loki TCP port",
								MarkdownDescription: "Loki TCP port",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"remove_keys": schema.ListAttribute{
								Description:         "Optional list of keys to remove.",
								MarkdownDescription: "Optional list of keys to remove.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tenant_id": schema.SingleNestedAttribute{
								Description:         "Tenant ID used by default to push logs to Loki.If omitted or empty it assumes Loki is running in single-tenant mode and no X-Scope-OrgID header is sent.",
								MarkdownDescription: "Tenant ID used by default to push logs to Loki.If omitted or empty it assumes Loki is running in single-tenant mode and no X-Scope-OrgID header is sent.",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"tenant_id_key": schema.StringAttribute{
								Description:         "Specify the name of the key from the original record that contains the Tenant ID.The value of the key is set as X-Scope-OrgID of HTTP header. It is useful to set Tenant ID dynamically.",
								MarkdownDescription: "Specify the name of the key from the original record that contains the Tenant ID.The value of the key is set as X-Scope-OrgID of HTTP header. It is useful to set Tenant ID dynamically.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"uri": schema.StringAttribute{
								Description:         "Specify a custom HTTP URI. It must start with forward slash.",
								MarkdownDescription: "Specify a custom HTTP URI. It must start with forward slash.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"match": schema.StringAttribute{
						Description:         "A pattern to match against the tags of incoming records.It's case sensitive and support the star (*) character as a wildcard.",
						MarkdownDescription: "A pattern to match against the tags of incoming records.It's case sensitive and support the star (*) character as a wildcard.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"match_regex": schema.StringAttribute{
						Description:         "A regular expression to match against the tags of incoming records.Use this option if you want to use the full regex syntax.",
						MarkdownDescription: "A regular expression to match against the tags of incoming records.Use this option if you want to use the full regex syntax.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"null": schema.MapAttribute{
						Description:         "Null defines Null Output configuration.",
						MarkdownDescription: "Null defines Null Output configuration.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"opensearch": schema.SingleNestedAttribute{
						Description:         "OpenSearch defines OpenSearch Output configuration.",
						MarkdownDescription: "OpenSearch defines OpenSearch Output configuration.",
						Attributes: map[string]schema.Attribute{
							"workers": schema.Int64Attribute{
								Description:         "Enables dedicated thread(s) for this output. Default value is set since version 1.8.13. For previous versions is 0.",
								MarkdownDescription: "Enables dedicated thread(s) for this output. Default value is set since version 1.8.13. For previous versions is 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_auth": schema.StringAttribute{
								Description:         "Enable AWS Sigv4 Authentication for Amazon OpenSearch Service.",
								MarkdownDescription: "Enable AWS Sigv4 Authentication for Amazon OpenSearch Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_external_id": schema.StringAttribute{
								Description:         "External ID for the AWS IAM Role specified with aws_role_arn.",
								MarkdownDescription: "External ID for the AWS IAM Role specified with aws_role_arn.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_region": schema.StringAttribute{
								Description:         "Specify the AWS region for Amazon OpenSearch Service.",
								MarkdownDescription: "Specify the AWS region for Amazon OpenSearch Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_role_arn": schema.StringAttribute{
								Description:         "AWS IAM Role to assume to put records to your Amazon cluster.",
								MarkdownDescription: "AWS IAM Role to assume to put records to your Amazon cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_sts_endpoint": schema.StringAttribute{
								Description:         "Specify the custom sts endpoint to be used with STS API for Amazon OpenSearch Service.",
								MarkdownDescription: "Specify the custom sts endpoint to be used with STS API for Amazon OpenSearch Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"buffer_size": schema.StringAttribute{
								Description:         "Specify the buffer size used to read the response from the OpenSearch HTTP service.This option is useful for debugging purposes where is required to read full responses,note that response size grows depending of the number of records inserted.To set an unlimited amount of memory set this value to False,otherwise the value must be according to the Unit Size specification.",
								MarkdownDescription: "Specify the buffer size used to read the response from the OpenSearch HTTP service.This option is useful for debugging purposes where is required to read full responses,note that response size grows depending of the number of records inserted.To set an unlimited amount of memory set this value to False,otherwise the value must be according to the Unit Size specification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"compress": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("gzip"),
								},
							},

							"current_time_index": schema.BoolAttribute{
								Description:         "Use current time for index generation instead of message record",
								MarkdownDescription: "Use current time for index generation instead of message record",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"generate_id": schema.BoolAttribute{
								Description:         "When enabled, generate _id for outgoing records.This prevents duplicate records when retrying OpenSearch.",
								MarkdownDescription: "When enabled, generate _id for outgoing records.This prevents duplicate records when retrying OpenSearch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target OpenSearch instance, default '127.0.0.1'",
								MarkdownDescription: "IP address or hostname of the target OpenSearch instance, default '127.0.0.1'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Password for user defined in HTTP_User",
								MarkdownDescription: "Password for user defined in HTTP_User",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Optional username credential for access",
								MarkdownDescription: "Optional username credential for access",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"id_key": schema.StringAttribute{
								Description:         "If set, _id will be the value of the key from incoming record and Generate_ID option is ignored.",
								MarkdownDescription: "If set, _id will be the value of the key from incoming record and Generate_ID option is ignored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include_tag_key": schema.BoolAttribute{
								Description:         "When enabled, it append the Tag name to the record.",
								MarkdownDescription: "When enabled, it append the Tag name to the record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"index": schema.StringAttribute{
								Description:         "Index name",
								MarkdownDescription: "Index name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_date_format": schema.StringAttribute{
								Description:         "Time format (based on strftime) to generate the second part of the Index name.",
								MarkdownDescription: "Time format (based on strftime) to generate the second part of the Index name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_format": schema.BoolAttribute{
								Description:         "Enable Logstash format compatibility.This option takes a boolean value: True/False, On/Off",
								MarkdownDescription: "Enable Logstash format compatibility.This option takes a boolean value: True/False, On/Off",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_prefix": schema.StringAttribute{
								Description:         "When Logstash_Format is enabled, the Index name is composed using a prefix and the date,e.g: If Logstash_Prefix is equals to 'mydata' your index will become 'mydata-YYYY.MM.DD'.The last string appended belongs to the date when the data is being generated.",
								MarkdownDescription: "When Logstash_Format is enabled, the Index name is composed using a prefix and the date,e.g: If Logstash_Prefix is equals to 'mydata' your index will become 'mydata-YYYY.MM.DD'.The last string appended belongs to the date when the data is being generated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logstash_prefix_key": schema.StringAttribute{
								Description:         "Prefix keys with this string",
								MarkdownDescription: "Prefix keys with this string",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": schema.StringAttribute{
								Description:         "OpenSearch accepts new data on HTTP query path '/_bulk'.But it is also possible to serve OpenSearch behind a reverse proxy on a subpath.This option defines such path on the fluent-bit side.It simply adds a path prefix in the indexing HTTP POST URI.",
								MarkdownDescription: "OpenSearch accepts new data on HTTP query path '/_bulk'.But it is also possible to serve OpenSearch behind a reverse proxy on a subpath.This option defines such path on the fluent-bit side.It simply adds a path prefix in the indexing HTTP POST URI.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pipeline": schema.StringAttribute{
								Description:         "OpenSearch allows to setup filters called pipelines.This option allows to define which pipeline the database should use.For performance reasons is strongly suggested to do parsingand filtering on Fluent Bit side, avoid pipelines.",
								MarkdownDescription: "OpenSearch allows to setup filters called pipelines.This option allows to define which pipeline the database should use.For performance reasons is strongly suggested to do parsingand filtering on Fluent Bit side, avoid pipelines.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target OpenSearch instance, default '9200'",
								MarkdownDescription: "TCP port of the target OpenSearch instance, default '9200'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"replace_dots": schema.BoolAttribute{
								Description:         "When enabled, replace field name dots with underscore, required by Elasticsearch 2.0-2.3.",
								MarkdownDescription: "When enabled, replace field name dots with underscore, required by Elasticsearch 2.0-2.3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"suppress_type_name": schema.BoolAttribute{
								Description:         "When enabled, mapping types is removed and Type option is ignored. Types are deprecated in APIs in v7.0. This options is for v7.0 or later.",
								MarkdownDescription: "When enabled, mapping types is removed and Type option is ignored. Types are deprecated in APIs in v7.0. This options is for v7.0 or later.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_key": schema.StringAttribute{
								Description:         "When Include_Tag_Key is enabled, this property defines the key name for the tag.",
								MarkdownDescription: "When Include_Tag_Key is enabled, this property defines the key name for the tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key": schema.StringAttribute{
								Description:         "When Logstash_Format is enabled, each record will get a new timestamp field.The Time_Key property defines the name of that field.",
								MarkdownDescription: "When Logstash_Format is enabled, each record will get a new timestamp field.The Time_Key property defines the name of that field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key_format": schema.StringAttribute{
								Description:         "When Logstash_Format is enabled, this property defines the format of the timestamp.",
								MarkdownDescription: "When Logstash_Format is enabled, this property defines the format of the timestamp.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_key_nanos": schema.BoolAttribute{
								Description:         "When Logstash_Format is enabled, enabling this property sends nanosecond precision timestamps.",
								MarkdownDescription: "When Logstash_Format is enabled, enabling this property sends nanosecond precision timestamps.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"total_limit_size": schema.StringAttribute{
								Description:         "Limit the maximum number of Chunks in the filesystem for the current output logical destination.",
								MarkdownDescription: "Limit the maximum number of Chunks in the filesystem for the current output logical destination.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trace_error": schema.BoolAttribute{
								Description:         "When enabled print the elasticsearch API calls to stdout when elasticsearch returns an error",
								MarkdownDescription: "When enabled print the elasticsearch API calls to stdout when elasticsearch returns an error",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trace_output": schema.BoolAttribute{
								Description:         "When enabled print the elasticsearch API calls to stdout (for diag only)",
								MarkdownDescription: "When enabled print the elasticsearch API calls to stdout (for diag only)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type name",
								MarkdownDescription: "Type name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_operation": schema.StringAttribute{
								Description:         "Operation to use to write in bulk requests.",
								MarkdownDescription: "Operation to use to write in bulk requests.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"opentelemetry": schema.SingleNestedAttribute{
						Description:         "OpenTelemetry defines OpenTelemetry Output configuration.",
						MarkdownDescription: "OpenTelemetry defines OpenTelemetry Output configuration.",
						Attributes: map[string]schema.Attribute{
							"add_label": schema.MapAttribute{
								Description:         "This allows you to add custom labels to all metrics exposed through the OpenTelemetry exporter. You may have multiple of these fields.",
								MarkdownDescription: "This allows you to add custom labels to all metrics exposed through the OpenTelemetry exporter. You may have multiple of these fields.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"header": schema.MapAttribute{
								Description:         "Add a HTTP header key/value pair. Multiple headers can be set.",
								MarkdownDescription: "Add a HTTP header key/value pair. Multiple headers can be set.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target HTTP Server, default '127.0.0.1'",
								MarkdownDescription: "IP address or hostname of the target HTTP Server, default '127.0.0.1'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Password for user defined in HTTP_User",
								MarkdownDescription: "Password for user defined in HTTP_User",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Optional username credential for access",
								MarkdownDescription: "Optional username credential for access",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"log_response_payload": schema.BoolAttribute{
								Description:         "Log the response payload within the Fluent Bit log.",
								MarkdownDescription: "Log the response payload within the Fluent Bit log.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logs_uri": schema.StringAttribute{
								Description:         "Specify an optional HTTP URI for the target web server listening for logs, e.g: /v1/logs",
								MarkdownDescription: "Specify an optional HTTP URI for the target web server listening for logs, e.g: /v1/logs",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_uri": schema.StringAttribute{
								Description:         "Specify an optional HTTP URI for the target web server listening for metrics, e.g: /v1/metrics",
								MarkdownDescription: "Specify an optional HTTP URI for the target web server listening for metrics, e.g: /v1/metrics",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target OpenSearch instance, default '80'",
								MarkdownDescription: "TCP port of the target OpenSearch instance, default '80'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"proxy": schema.StringAttribute{
								Description:         "Specify an HTTP Proxy. The expected format of this value is http://HOST:PORT. Note that HTTPS is not currently supported.It is recommended not to set this and to configure the HTTP proxy environment variables instead as they support both HTTP and HTTPS.",
								MarkdownDescription: "Specify an HTTP Proxy. The expected format of this value is http://HOST:PORT. Note that HTTPS is not currently supported.It is recommended not to set this and to configure the HTTP proxy environment variables instead as they support both HTTP and HTTPS.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"traces_uri": schema.StringAttribute{
								Description:         "Specify an optional HTTP URI for the target web server listening for traces, e.g: /v1/traces",
								MarkdownDescription: "Specify an optional HTTP URI for the target web server listening for traces, e.g: /v1/traces",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"processors": schema.MapAttribute{
						Description:         "Processors defines the processors configuration",
						MarkdownDescription: "Processors defines the processors configuration",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_exporter": schema.SingleNestedAttribute{
						Description:         "PrometheusExporter_types defines Prometheus exporter configuration to expose metrics from Fluent Bit.",
						MarkdownDescription: "PrometheusExporter_types defines Prometheus exporter configuration to expose metrics from Fluent Bit.",
						Attributes: map[string]schema.Attribute{
							"add_labels": schema.MapAttribute{
								Description:         "This allows you to add custom labels to all metrics exposed through the prometheus exporter. You may have multiple of these fields",
								MarkdownDescription: "This allows you to add custom labels to all metrics exposed through the prometheus exporter. You may have multiple of these fields",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target HTTP Server, default: 0.0.0.0",
								MarkdownDescription: "IP address or hostname of the target HTTP Server, default: 0.0.0.0",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "This is the port Fluent Bit will bind to when hosting prometheus metrics.",
								MarkdownDescription: "This is the port Fluent Bit will bind to when hosting prometheus metrics.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_remote_write": schema.SingleNestedAttribute{
						Description:         "PrometheusRemoteWrite_types defines Prometheus Remote Write configuration.",
						MarkdownDescription: "PrometheusRemoteWrite_types defines Prometheus Remote Write configuration.",
						Attributes: map[string]schema.Attribute{
							"add_labels": schema.MapAttribute{
								Description:         "This allows you to add custom labels to all metrics exposed through the prometheus exporter. You may have multiple of these fields",
								MarkdownDescription: "This allows you to add custom labels to all metrics exposed through the prometheus exporter. You may have multiple of these fields",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"headers": schema.MapAttribute{
								Description:         "Add a HTTP header key/value pair. Multiple headers can be set.",
								MarkdownDescription: "Add a HTTP header key/value pair. Multiple headers can be set.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target HTTP Server, default: 127.0.0.1",
								MarkdownDescription: "IP address or hostname of the target HTTP Server, default: 127.0.0.1",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"http_passwd": schema.SingleNestedAttribute{
								Description:         "Basic Auth Password.Requires HTTP_user to be se",
								MarkdownDescription: "Basic Auth Password.Requires HTTP_user to be se",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Basic Auth Username",
								MarkdownDescription: "Basic Auth Username",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"log_response_payload": schema.BoolAttribute{
								Description:         "Log the response payload within the Fluent Bit log,default: false",
								MarkdownDescription: "Log the response payload within the Fluent Bit log,default: false",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target HTTP Serveri, default:80",
								MarkdownDescription: "TCP port of the target HTTP Serveri, default:80",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"proxy": schema.StringAttribute{
								Description:         "Specify an HTTP Proxy. The expected format of this value is http://HOST:PORT.",
								MarkdownDescription: "Specify an HTTP Proxy. The expected format of this value is http://HOST:PORT.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"uri": schema.StringAttribute{
								Description:         "Specify an optional HTTP URI for the target web server, e.g: /something ,default: /",
								MarkdownDescription: "Specify an optional HTTP URI for the target web server, e.g: /something ,default: /",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workers": schema.Int64Attribute{
								Description:         "Enables dedicated thread(s) for this output. Default value is set since version 1.8.13. For previous versions is 0,default : 2",
								MarkdownDescription: "Enables dedicated thread(s) for this output. Default value is set since version 1.8.13. For previous versions is 0,default : 2",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retry_limit": schema.StringAttribute{
						Description:         "RetryLimit represents configuration for the scheduler which can be set independently on each output section.This option allows to disable retries or impose a limit to try N times and then discard the data after reaching that limit.",
						MarkdownDescription: "RetryLimit represents configuration for the scheduler which can be set independently on each output section.This option allows to disable retries or impose a limit to try N times and then discard the data after reaching that limit.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"s3": schema.SingleNestedAttribute{
						Description:         "S3 defines S3 Output configuration.",
						MarkdownDescription: "S3 defines S3 Output configuration.",
						Attributes: map[string]schema.Attribute{
							"auto_retry_requests": schema.BoolAttribute{
								Description:         "Immediately retry failed requests to AWS services once.",
								MarkdownDescription: "Immediately retry failed requests to AWS services once.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bucket": schema.StringAttribute{
								Description:         "S3 Bucket name",
								MarkdownDescription: "S3 Bucket name",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"canned_acl": schema.StringAttribute{
								Description:         "Predefined Canned ACL Policy for S3 objects.",
								MarkdownDescription: "Predefined Canned ACL Policy for S3 objects.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compression": schema.StringAttribute{
								Description:         "Compression type for S3 objects.",
								MarkdownDescription: "Compression type for S3 objects.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"content_type": schema.StringAttribute{
								Description:         "A standard MIME type for the S3 object; this will be set as the Content-Type HTTP header.",
								MarkdownDescription: "A standard MIME type for the S3 object; this will be set as the Content-Type HTTP header.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "Custom endpoint for the S3 API.",
								MarkdownDescription: "Custom endpoint for the S3 API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_id": schema.StringAttribute{
								Description:         "Specify an external ID for the STS API, can be used with the role_arn parameter if your role requires an external ID.",
								MarkdownDescription: "Specify an external ID for the STS API, can be used with the role_arn parameter if your role requires an external ID.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"json_date_format": schema.StringAttribute{
								Description:         "Specify the format of the date. Supported formats are double, epoch, iso8601 (eg: 2018-05-30T09:39:52.000681Z) and java_sql_timestamp (eg: 2018-05-30 09:39:52.000681)",
								MarkdownDescription: "Specify the format of the date. Supported formats are double, epoch, iso8601 (eg: 2018-05-30T09:39:52.000681Z) and java_sql_timestamp (eg: 2018-05-30 09:39:52.000681)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"json_date_key": schema.StringAttribute{
								Description:         "Specify the name of the time key in the output record. To disable the time key just set the value to false.",
								MarkdownDescription: "Specify the name of the time key in the output record. To disable the time key just set the value to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_key": schema.StringAttribute{
								Description:         "By default, the whole log record will be sent to S3. If you specify a key name with this option, then only the value of that key will be sent to S3.",
								MarkdownDescription: "By default, the whole log record will be sent to S3. If you specify a key name with this option, then only the value of that key will be sent to S3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"preserve_data_ordering": schema.BoolAttribute{
								Description:         "Normally, when an upload request fails, there is a high chance for the last received chunk to be swapped with a later chunk, resulting in data shuffling. This feature prevents this shuffling by using a queue logic for uploads.",
								MarkdownDescription: "Normally, when an upload request fails, there is a high chance for the last received chunk to be swapped with a later chunk, resulting in data shuffling. This feature prevents this shuffling by using a queue logic for uploads.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"profile": schema.StringAttribute{
								Description:         "Option to specify an AWS Profile for credentials.",
								MarkdownDescription: "Option to specify an AWS Profile for credentials.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"region": schema.StringAttribute{
								Description:         "The AWS region of your S3 bucket",
								MarkdownDescription: "The AWS region of your S3 bucket",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"retry_limit": schema.Int64Attribute{
								Description:         "Integer value to set the maximum number of retries allowed.",
								MarkdownDescription: "Integer value to set the maximum number of retries allowed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_arn": schema.StringAttribute{
								Description:         "ARN of an IAM role to assume",
								MarkdownDescription: "ARN of an IAM role to assume",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_key_format": schema.StringAttribute{
								Description:         "Format string for keys in S3.",
								MarkdownDescription: "Format string for keys in S3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_key_format_tag_delimiters": schema.StringAttribute{
								Description:         "A series of characters which will be used to split the tag into 'parts' for use with the s3_key_format option.",
								MarkdownDescription: "A series of characters which will be used to split the tag into 'parts' for use with the s3_key_format option.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"send_content_md5": schema.BoolAttribute{
								Description:         "Send the Content-MD5 header with PutObject and UploadPart requests, as is required when Object Lock is enabled.",
								MarkdownDescription: "Send the Content-MD5 header with PutObject and UploadPart requests, as is required when Object Lock is enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"static_file_path": schema.BoolAttribute{
								Description:         "Disables behavior where UUID string is automatically appended to end of S3 key name when $UUID is not provided in s3_key_format. $UUID, time formatters, $TAG, and other dynamic key formatters all work as expected while this feature is set to true.",
								MarkdownDescription: "Disables behavior where UUID string is automatically appended to end of S3 key name when $UUID is not provided in s3_key_format. $UUID, time formatters, $TAG, and other dynamic key formatters all work as expected while this feature is set to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class": schema.StringAttribute{
								Description:         "Specify the storage class for S3 objects. If this option is not specified, objects will be stored with the default 'STANDARD' storage class.",
								MarkdownDescription: "Specify the storage class for S3 objects. If this option is not specified, objects will be stored with the default 'STANDARD' storage class.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"store_dir": schema.StringAttribute{
								Description:         "Directory to locally buffer data before sending.",
								MarkdownDescription: "Directory to locally buffer data before sending.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"store_dir_limit_size": schema.StringAttribute{
								Description:         "The size of the limitation for disk usage in S3.",
								MarkdownDescription: "The size of the limitation for disk usage in S3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sts_endpoint": schema.StringAttribute{
								Description:         "Custom endpoint for the STS API.",
								MarkdownDescription: "Custom endpoint for the STS API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"total_file_size": schema.StringAttribute{
								Description:         "Specifies the size of files in S3. Minimum size is 1M. With use_put_object On the maximum size is 1G. With multipart upload mode, the maximum size is 50G.",
								MarkdownDescription: "Specifies the size of files in S3. Minimum size is 1M. With use_put_object On the maximum size is 1G. With multipart upload mode, the maximum size is 50G.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"upload_chunk_size": schema.StringAttribute{
								Description:         "The size of each 'part' for multipart uploads. Max: 50M",
								MarkdownDescription: "The size of each 'part' for multipart uploads. Max: 50M",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"upload_timeout": schema.StringAttribute{
								Description:         "Whenever this amount of time has elapsed, Fluent Bit will complete an upload and create a new file in S3. For example, set this value to 60m and you will get a new file every hour.",
								MarkdownDescription: "Whenever this amount of time has elapsed, Fluent Bit will complete an upload and create a new file in S3. For example, set this value to 60m and you will get a new file every hour.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_put_object": schema.BoolAttribute{
								Description:         "Use the S3 PutObject API, instead of the multipart upload API.",
								MarkdownDescription: "Use the S3 PutObject API, instead of the multipart upload API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"splunk": schema.SingleNestedAttribute{
						Description:         "Splunk defines Splunk Output Configuration",
						MarkdownDescription: "Splunk defines Splunk Output Configuration",
						Attributes: map[string]schema.Attribute{
							"workers": schema.Int64Attribute{
								Description:         "Enables dedicated thread(s) for this output. Default value '2' is set since version 1.8.13. For previous versions is 0.",
								MarkdownDescription: "Enables dedicated thread(s) for this output. Default value '2' is set since version 1.8.13. For previous versions is 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"channel": schema.StringAttribute{
								Description:         "Specify X-Splunk-Request-Channel Header for the HTTP Event Collector interface.",
								MarkdownDescription: "Specify X-Splunk-Request-Channel Header for the HTTP Event Collector interface.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compress": schema.StringAttribute{
								Description:         "Set payload compression mechanism. The only available option is gzip.",
								MarkdownDescription: "Set payload compression mechanism. The only available option is gzip.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_fields": schema.ListAttribute{
								Description:         "Set event fields for the record. This option is an array and the format is 'key_namerecord_accessor_pattern'.",
								MarkdownDescription: "Set event fields for the record. This option is an array and the format is 'key_namerecord_accessor_pattern'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_host": schema.StringAttribute{
								Description:         "Specify the key name that contains the host value. This option allows a record accessors pattern.",
								MarkdownDescription: "Specify the key name that contains the host value. This option allows a record accessors pattern.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_index": schema.StringAttribute{
								Description:         "The name of the index by which the event data is to be indexed.",
								MarkdownDescription: "The name of the index by which the event data is to be indexed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_index_key": schema.StringAttribute{
								Description:         "Set a record key that will populate the index field. If the key is found, it will have precedenceover the value set in event_index.",
								MarkdownDescription: "Set a record key that will populate the index field. If the key is found, it will have precedenceover the value set in event_index.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_key": schema.StringAttribute{
								Description:         "Specify the key name that will be used to send a single value as part of the record.",
								MarkdownDescription: "Specify the key name that will be used to send a single value as part of the record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_source": schema.StringAttribute{
								Description:         "Set the source value to assign to the event data.",
								MarkdownDescription: "Set the source value to assign to the event data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_sourcetype": schema.StringAttribute{
								Description:         "Set the sourcetype value to assign to the event data.",
								MarkdownDescription: "Set the sourcetype value to assign to the event data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"event_sourcetype_key": schema.StringAttribute{
								Description:         "Set a record key that will populate 'sourcetype'. If the key is found, it will have precedenceover the value set in event_sourcetype.",
								MarkdownDescription: "Set a record key that will populate 'sourcetype'. If the key is found, it will have precedenceover the value set in event_sourcetype.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "IP address or hostname of the target OpenSearch instance, default '127.0.0.1'",
								MarkdownDescription: "IP address or hostname of the target OpenSearch instance, default '127.0.0.1'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_buffer_size": schema.StringAttribute{
								Description:         "Buffer size used to receive Splunk HTTP responses: Default '2M'",
								MarkdownDescription: "Buffer size used to receive Splunk HTTP responses: Default '2M'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"http_debug_bad_request": schema.BoolAttribute{
								Description:         "If the HTTP server response code is 400 (bad request) and this flag is enabled, it will print the full HTTP requestand response to the stdout interface. This feature is available for debugging purposes.",
								MarkdownDescription: "If the HTTP server response code is 400 (bad request) and this flag is enabled, it will print the full HTTP requestand response to the stdout interface. This feature is available for debugging purposes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_password": schema.SingleNestedAttribute{
								Description:         "Password for user defined in HTTP_User",
								MarkdownDescription: "Password for user defined in HTTP_User",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
								Description:         "Optional username credential for access",
								MarkdownDescription: "Optional username credential for access",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port of the target Splunk instance, default '8088'",
								MarkdownDescription: "TCP port of the target Splunk instance, default '8088'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"splunk_send_raw": schema.BoolAttribute{
								Description:         "When enabled, the record keys and values are set in the top level of the map instead of under the event key. Refer tothe Sending Raw Events section from the docs more details to make this option work properly.",
								MarkdownDescription: "When enabled, the record keys and values are set in the top level of the map instead of under the event key. Refer tothe Sending Raw Events section from the docs more details to make this option work properly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"splunk_token": schema.SingleNestedAttribute{
								Description:         "Specify the Authentication Token for the HTTP Event Collector interface.",
								MarkdownDescription: "Specify the Authentication Token for the HTTP Event Collector interface.",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"stackdriver": schema.SingleNestedAttribute{
						Description:         "Stackdriver defines Stackdriver Output Configuration",
						MarkdownDescription: "Stackdriver defines Stackdriver Output Configuration",
						Attributes: map[string]schema.Attribute{
							"autoformat_stackdriver_trace": schema.BoolAttribute{
								Description:         "Rewrite the trace field to be formatted for use with GCP Cloud Trace",
								MarkdownDescription: "Rewrite the trace field to be formatted for use with GCP Cloud Trace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_k8s_regex": schema.StringAttribute{
								Description:         "A custom regex to extract fields from the local_resource_id of the logs",
								MarkdownDescription: "A custom regex to extract fields from the local_resource_id of the logs",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"export_to_project_id": schema.StringAttribute{
								Description:         "The GCP Project that should receive the logs",
								MarkdownDescription: "The GCP Project that should receive the logs",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"google_service_credentials": schema.StringAttribute{
								Description:         "Path to GCP Credentials JSON file",
								MarkdownDescription: "Path to GCP Credentials JSON file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"job": schema.StringAttribute{
								Description:         "Identifier for a grouping of tasks. Required if Resource is generic_task",
								MarkdownDescription: "Identifier for a grouping of tasks. Required if Resource is generic_task",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"k8s_cluster_location": schema.StringAttribute{
								Description:         "Location of the cluster that contains the pods/nodes. Required if Resource is k8s_container, k8s_node, or k8s_pod",
								MarkdownDescription: "Location of the cluster that contains the pods/nodes. Required if Resource is k8s_container, k8s_node, or k8s_pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"k8s_cluster_name": schema.StringAttribute{
								Description:         "Name of the cluster that the pod is running in. Required if Resource is k8s_container, k8s_node, or k8s_pod",
								MarkdownDescription: "Name of the cluster that the pod is running in. Required if Resource is k8s_container, k8s_node, or k8s_pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.ListAttribute{
								Description:         "Optional list of comma separated of strings for key/value pairs",
								MarkdownDescription: "Optional list of comma separated of strings for key/value pairs",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels_key": schema.StringAttribute{
								Description:         "Used by Stackdriver to find related labels and extract them to LogEntry Labels",
								MarkdownDescription: "Used by Stackdriver to find related labels and extract them to LogEntry Labels",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"location": schema.StringAttribute{
								Description:         "GCP/AWS region to store data. Required if Resource is generic_node or generic_task",
								MarkdownDescription: "GCP/AWS region to store data. Required if Resource is generic_node or generic_task",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_name_key": schema.StringAttribute{
								Description:         "The value of this field is set as the logName field in Stackdriver",
								MarkdownDescription: "The value of this field is set as the logName field in Stackdriver",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_server": schema.StringAttribute{
								Description:         "Metadata Server Prefix",
								MarkdownDescription: "Metadata Server Prefix",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace identifier. Required if Resource is generic_node or generic_task",
								MarkdownDescription: "Namespace identifier. Required if Resource is generic_node or generic_task",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_id": schema.StringAttribute{
								Description:         "Node identifier within the namespace. Required if Resource is generic_node or generic_task",
								MarkdownDescription: "Node identifier within the namespace. Required if Resource is generic_node or generic_task",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource": schema.StringAttribute{
								Description:         "Set resource types of data",
								MarkdownDescription: "Set resource types of data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_labels": schema.ListAttribute{
								Description:         "Optional list of comma seperated strings. Setting these fields overrides the Stackdriver monitored resource API values",
								MarkdownDescription: "Optional list of comma seperated strings. Setting these fields overrides the Stackdriver monitored resource API values",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_email": schema.SingleNestedAttribute{
								Description:         "Email associated with the service",
								MarkdownDescription: "Email associated with the service",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"service_account_secret": schema.SingleNestedAttribute{
								Description:         "Private Key associated with the service",
								MarkdownDescription: "Private Key associated with the service",
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
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"severity_key": schema.StringAttribute{
								Description:         "Specify the key that contains the severity information for the logs",
								MarkdownDescription: "Specify the key that contains the severity information for the logs",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_prefix": schema.StringAttribute{
								Description:         "Used to validate the tags of logs that when the Resource is k8s_container, k8s_node, or k8s_pod",
								MarkdownDescription: "Used to validate the tags of logs that when the Resource is k8s_container, k8s_node, or k8s_pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"task_id": schema.StringAttribute{
								Description:         "Identifier for a task within a namespace. Required if Resource is generic_task",
								MarkdownDescription: "Identifier for a task within a namespace. Required if Resource is generic_task",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workers": schema.Int64Attribute{
								Description:         "Number of dedicated threads for the Stackdriver Output Plugin",
								MarkdownDescription: "Number of dedicated threads for the Stackdriver Output Plugin",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stdout": schema.SingleNestedAttribute{
						Description:         "Stdout defines Stdout Output configuration.",
						MarkdownDescription: "Stdout defines Stdout Output configuration.",
						Attributes: map[string]schema.Attribute{
							"format": schema.StringAttribute{
								Description:         "Specify the data format to be printed. Supported formats are msgpack json, json_lines and json_stream.",
								MarkdownDescription: "Specify the data format to be printed. Supported formats are msgpack json, json_lines and json_stream.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("msgpack", "json", "json_lines", "json_stream"),
								},
							},

							"json_date_format": schema.StringAttribute{
								Description:         "Specify the format of the date. Supported formats are double,  iso8601 (eg: 2018-05-30T09:39:52.000681Z) and epoch.",
								MarkdownDescription: "Specify the format of the date. Supported formats are double,  iso8601 (eg: 2018-05-30T09:39:52.000681Z) and epoch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("double", "iso8601", "epoch"),
								},
							},

							"json_date_key": schema.StringAttribute{
								Description:         "Specify the name of the date field in output.",
								MarkdownDescription: "Specify the name of the date field in output.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"syslog": schema.SingleNestedAttribute{
						Description:         "Syslog defines Syslog Output configuration.",
						MarkdownDescription: "Syslog defines Syslog Output configuration.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Host domain or IP address of the remote Syslog server.",
								MarkdownDescription: "Host domain or IP address of the remote Syslog server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "Mode of the desired transport type, the available options are tcp, tls and udp.",
								MarkdownDescription: "Mode of the desired transport type, the available options are tcp, tls and udp.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP or UDP port of the remote Syslog server.",
								MarkdownDescription: "TCP or UDP port of the remote Syslog server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"syslog_appname_key": schema.StringAttribute{
								Description:         "Key name from the original record that contains the application name that generated the message.",
								MarkdownDescription: "Key name from the original record that contains the application name that generated the message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_facility_key": schema.StringAttribute{
								Description:         "Key from the original record that contains the Syslog facility number.",
								MarkdownDescription: "Key from the original record that contains the Syslog facility number.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_format": schema.StringAttribute{
								Description:         "Syslog protocol format to use, the available options are rfc3164 and rfc5424.",
								MarkdownDescription: "Syslog protocol format to use, the available options are rfc3164 and rfc5424.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_hostname_key": schema.StringAttribute{
								Description:         "Key name from the original record that contains the hostname that generated the message.",
								MarkdownDescription: "Key name from the original record that contains the hostname that generated the message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_max_size": schema.Int64Attribute{
								Description:         "Maximum size allowed per message, in bytes.",
								MarkdownDescription: "Maximum size allowed per message, in bytes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_message_id_key": schema.StringAttribute{
								Description:         "Key name from the original record that contains the Message ID associated to the message.",
								MarkdownDescription: "Key name from the original record that contains the Message ID associated to the message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_message_key": schema.StringAttribute{
								Description:         "Key key name that contains the message to deliver.",
								MarkdownDescription: "Key key name that contains the message to deliver.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_process_id_key": schema.StringAttribute{
								Description:         "Key name from the original record that contains the Process ID that generated the message.",
								MarkdownDescription: "Key name from the original record that contains the Process ID that generated the message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_sd_key": schema.StringAttribute{
								Description:         "Key name from the original record that contains the Structured Data (SD) content.",
								MarkdownDescription: "Key name from the original record that contains the Structured Data (SD) content.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_severity_key": schema.StringAttribute{
								Description:         "Key from the original record that contains the Syslog severity number.",
								MarkdownDescription: "Key from the original record that contains the Syslog severity number.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Syslog output plugin supports TTL/SSL, for more details about the properties availableand general configuration, please refer to the TLS/SSL section.",
								MarkdownDescription: "Syslog output plugin supports TTL/SSL, for more details about the properties availableand general configuration, please refer to the TLS/SSL section.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"tcp": schema.SingleNestedAttribute{
						Description:         "TCP defines TCP Output configuration.",
						MarkdownDescription: "TCP defines TCP Output configuration.",
						Attributes: map[string]schema.Attribute{
							"format": schema.StringAttribute{
								Description:         "Specify the data format to be printed. Supported formats are msgpack json, json_lines and json_stream.",
								MarkdownDescription: "Specify the data format to be printed. Supported formats are msgpack json, json_lines and json_stream.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("msgpack", "json", "json_lines", "json_stream"),
								},
							},

							"host": schema.StringAttribute{
								Description:         "Target host where Fluent-Bit or Fluentd are listening for Forward messages.",
								MarkdownDescription: "Target host where Fluent-Bit or Fluentd are listening for Forward messages.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"json_date_format": schema.StringAttribute{
								Description:         "Specify the format of the date. Supported formats are double, epochand iso8601 (eg: 2018-05-30T09:39:52.000681Z)",
								MarkdownDescription: "Specify the format of the date. Supported formats are double, epochand iso8601 (eg: 2018-05-30T09:39:52.000681Z)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("double", "epoch", "iso8601"),
								},
							},

							"json_date_key": schema.StringAttribute{
								Description:         "TSpecify the name of the time key in the output record.To disable the time key just set the value to false.",
								MarkdownDescription: "TSpecify the name of the time key in the output record.To disable the time key just set the value to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"networking": schema.SingleNestedAttribute{
								Description:         "Include fluentbit networking options for this output-plugin",
								MarkdownDescription: "Include fluentbit networking options for this output-plugin",
								Attributes: map[string]schema.Attribute{
									"dns_mode": schema.StringAttribute{
										Description:         "Select the primary DNS connection type (TCP or UDP).",
										MarkdownDescription: "Select the primary DNS connection type (TCP or UDP).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP"),
										},
									},

									"dns_prefer_i_pv4": schema.BoolAttribute{
										Description:         "Prioritize IPv4 DNS results when trying to establish a connection.",
										MarkdownDescription: "Prioritize IPv4 DNS results when trying to establish a connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "Select the primary DNS resolver type (LEGACY or ASYNC).",
										MarkdownDescription: "Select the primary DNS resolver type (LEGACY or ASYNC).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("LEGACY", "ASYNC"),
										},
									},

									"connect_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										MarkdownDescription: "Set maximum time expressed in seconds to wait for a TCP connection to be established, this include the TLS handshake time.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										MarkdownDescription: "On connection timeout, specify if it should log an error. When disabled, the timeout is logged as a debug message.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.StringAttribute{
										Description:         "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										MarkdownDescription: "Enable or disable connection keepalive support. Accepts a boolean value: on / off.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "Set maximum time expressed in seconds for an idle keepalive connection.",
										MarkdownDescription: "Set maximum time expressed in seconds for an idle keepalive connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "Set maximum number of times a keepalive connection can be used before it is retired.",
										MarkdownDescription: "Set maximum number of times a keepalive connection can be used before it is retired.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_worker_connections": schema.Int64Attribute{
										Description:         "Set maximum number of TCP connections that can be established per worker.",
										MarkdownDescription: "Set maximum number of TCP connections that can be established per worker.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
										Description:         "Specify network address to bind for data traffic.",
										MarkdownDescription: "Specify network address to bind for data traffic.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP Port of the target service.",
								MarkdownDescription: "TCP Port of the target service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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
	}
}

func (r *FluentbitFluentIoClusterOutputV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentbit_fluent_io_cluster_output_v1alpha2_manifest")

	var model FluentbitFluentIoClusterOutputV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentbit.fluent.io/v1alpha2")
	model.Kind = pointer.String("ClusterOutput")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
