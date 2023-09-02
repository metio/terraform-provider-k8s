/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gloo_solo_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ resource.Resource                = &GlooSoloIoUpstreamV1Resource{}
	_ resource.ResourceWithConfigure   = &GlooSoloIoUpstreamV1Resource{}
	_ resource.ResourceWithImportState = &GlooSoloIoUpstreamV1Resource{}
)

func NewGlooSoloIoUpstreamV1Resource() resource.Resource {
	return &GlooSoloIoUpstreamV1Resource{}
}

type GlooSoloIoUpstreamV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type GlooSoloIoUpstreamV1ResourceData struct {
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
		Aws *struct {
			AwsAccountId         *string `tfsdk:"aws_account_id" json:"awsAccountId,omitempty"`
			DestinationOverrides *struct {
				InvocationStyle        *string `tfsdk:"invocation_style" json:"invocationStyle,omitempty"`
				LogicalName            *string `tfsdk:"logical_name" json:"logicalName,omitempty"`
				RequestTransformation  *bool   `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
				ResponseTransformation *bool   `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
				UnwrapAsAlb            *bool   `tfsdk:"unwrap_as_alb" json:"unwrapAsAlb,omitempty"`
				UnwrapAsApiGateway     *bool   `tfsdk:"unwrap_as_api_gateway" json:"unwrapAsApiGateway,omitempty"`
				WrapAsApiGateway       *bool   `tfsdk:"wrap_as_api_gateway" json:"wrapAsApiGateway,omitempty"`
			} `tfsdk:"destination_overrides" json:"destinationOverrides,omitempty"`
			DisableRoleChaining *bool `tfsdk:"disable_role_chaining" json:"disableRoleChaining,omitempty"`
			LambdaFunctions     *[]struct {
				LambdaFunctionName *string `tfsdk:"lambda_function_name" json:"lambdaFunctionName,omitempty"`
				LogicalName        *string `tfsdk:"logical_name" json:"logicalName,omitempty"`
				Qualifier          *string `tfsdk:"qualifier" json:"qualifier,omitempty"`
			} `tfsdk:"lambda_functions" json:"lambdaFunctions,omitempty"`
			Region    *string `tfsdk:"region" json:"region,omitempty"`
			RoleArn   *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"aws" json:"aws,omitempty"`
		AwsEc2 *struct {
			Filters *[]struct {
				Key    *string `tfsdk:"key" json:"key,omitempty"`
				KvPair *struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"kv_pair" json:"kvPair,omitempty"`
			} `tfsdk:"filters" json:"filters,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			PublicIp  *bool   `tfsdk:"public_ip" json:"publicIp,omitempty"`
			Region    *string `tfsdk:"region" json:"region,omitempty"`
			RoleArn   *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"aws_ec2" json:"awsEc2,omitempty"`
		Azure *struct {
			FunctionAppName *string `tfsdk:"function_app_name" json:"functionAppName,omitempty"`
			Functions       *[]struct {
				AuthLevel    *string `tfsdk:"auth_level" json:"authLevel,omitempty"`
				FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
			} `tfsdk:"functions" json:"functions,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"azure" json:"azure,omitempty"`
		CircuitBreakers *struct {
			MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
			MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
			MaxRequests        *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
			MaxRetries         *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
		} `tfsdk:"circuit_breakers" json:"circuitBreakers,omitempty"`
		ConnectionConfig *struct {
			CommonHttpProtocolOptions *struct {
				HeadersWithUnderscoresAction *string `tfsdk:"headers_with_underscores_action" json:"headersWithUnderscoresAction,omitempty"`
				IdleTimeout                  *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
				MaxHeadersCount              *int64  `tfsdk:"max_headers_count" json:"maxHeadersCount,omitempty"`
				MaxStreamDuration            *string `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
			} `tfsdk:"common_http_protocol_options" json:"commonHttpProtocolOptions,omitempty"`
			ConnectTimeout       *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
			Http1ProtocolOptions *struct {
				EnableTrailers                          *bool `tfsdk:"enable_trailers" json:"enableTrailers,omitempty"`
				OverrideStreamErrorOnInvalidHttpMessage *bool `tfsdk:"override_stream_error_on_invalid_http_message" json:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`
				PreserveCaseHeaderKeyFormat             *bool `tfsdk:"preserve_case_header_key_format" json:"preserveCaseHeaderKeyFormat,omitempty"`
				ProperCaseHeaderKeyFormat               *bool `tfsdk:"proper_case_header_key_format" json:"properCaseHeaderKeyFormat,omitempty"`
			} `tfsdk:"http1_protocol_options" json:"http1ProtocolOptions,omitempty"`
			MaxRequestsPerConnection      *int64 `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
			PerConnectionBufferLimitBytes *int64 `tfsdk:"per_connection_buffer_limit_bytes" json:"perConnectionBufferLimitBytes,omitempty"`
			TcpKeepalive                  *struct {
				KeepaliveInterval *string `tfsdk:"keepalive_interval" json:"keepaliveInterval,omitempty"`
				KeepaliveProbes   *int64  `tfsdk:"keepalive_probes" json:"keepaliveProbes,omitempty"`
				KeepaliveTime     *string `tfsdk:"keepalive_time" json:"keepaliveTime,omitempty"`
			} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
		} `tfsdk:"connection_config" json:"connectionConfig,omitempty"`
		Consul *struct {
			ConnectEnabled        *bool     `tfsdk:"connect_enabled" json:"connectEnabled,omitempty"`
			ConsistencyMode       *string   `tfsdk:"consistency_mode" json:"consistencyMode,omitempty"`
			DataCenters           *[]string `tfsdk:"data_centers" json:"dataCenters,omitempty"`
			InstanceBlacklistTags *[]string `tfsdk:"instance_blacklist_tags" json:"instanceBlacklistTags,omitempty"`
			InstanceTags          *[]string `tfsdk:"instance_tags" json:"instanceTags,omitempty"`
			QueryOptions          *struct {
				UseCache *bool `tfsdk:"use_cache" json:"useCache,omitempty"`
			} `tfsdk:"query_options" json:"queryOptions,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
			ServiceSpec *struct {
				Graphql *struct {
					Endpoint *struct {
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"graphql" json:"graphql,omitempty"`
				Grpc *struct {
					Descriptors  *string `tfsdk:"descriptors" json:"descriptors,omitempty"`
					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" json:"functionNames,omitempty"`
						PackageName   *string   `tfsdk:"package_name" json:"packageName,omitempty"`
						ServiceName   *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" json:"grpcServices,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				GrpcJsonTranscoder *struct {
					AutoMapping                  *bool     `tfsdk:"auto_mapping" json:"autoMapping,omitempty"`
					ConvertGrpcStatus            *bool     `tfsdk:"convert_grpc_status" json:"convertGrpcStatus,omitempty"`
					IgnoreUnknownQueryParameters *bool     `tfsdk:"ignore_unknown_query_parameters" json:"ignoreUnknownQueryParameters,omitempty"`
					IgnoredQueryParameters       *[]string `tfsdk:"ignored_query_parameters" json:"ignoredQueryParameters,omitempty"`
					MatchIncomingRequestRoute    *bool     `tfsdk:"match_incoming_request_route" json:"matchIncomingRequestRoute,omitempty"`
					PrintOptions                 *struct {
						AddWhitespace              *bool `tfsdk:"add_whitespace" json:"addWhitespace,omitempty"`
						AlwaysPrintEnumsAsInts     *bool `tfsdk:"always_print_enums_as_ints" json:"alwaysPrintEnumsAsInts,omitempty"`
						AlwaysPrintPrimitiveFields *bool `tfsdk:"always_print_primitive_fields" json:"alwaysPrintPrimitiveFields,omitempty"`
						PreserveProtoFieldNames    *bool `tfsdk:"preserve_proto_field_names" json:"preserveProtoFieldNames,omitempty"`
					} `tfsdk:"print_options" json:"printOptions,omitempty"`
					ProtoDescriptor          *string `tfsdk:"proto_descriptor" json:"protoDescriptor,omitempty"`
					ProtoDescriptorBin       *string `tfsdk:"proto_descriptor_bin" json:"protoDescriptorBin,omitempty"`
					ProtoDescriptorConfigMap *struct {
						ConfigMapRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						Key *string `tfsdk:"key" json:"key,omitempty"`
					} `tfsdk:"proto_descriptor_config_map" json:"protoDescriptorConfigMap,omitempty"`
					Services *[]string `tfsdk:"services" json:"services,omitempty"`
				} `tfsdk:"grpc_json_transcoder" json:"grpcJsonTranscoder,omitempty"`
				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" json:"inline,omitempty"`
						Url    *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"swagger_info" json:"swaggerInfo,omitempty"`
					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
						Body              *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						DynamicMetadataValues *[]struct {
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
							Value             *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
						EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						Extractors       *struct {
							Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header   *string            `tfsdk:"header" json:"header,omitempty"`
							Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
							Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"extractors" json:"extractors,omitempty"`
						Headers *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						HeadersToAppend *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
						HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
						IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
						ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
						Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
					} `tfsdk:"transformations" json:"transformations,omitempty"`
				} `tfsdk:"rest" json:"rest,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			ServiceTags *[]string `tfsdk:"service_tags" json:"serviceTags,omitempty"`
			SubsetTags  *[]string `tfsdk:"subset_tags" json:"subsetTags,omitempty"`
		} `tfsdk:"consul" json:"consul,omitempty"`
		DiscoveryMetadata *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"discovery_metadata" json:"discoveryMetadata,omitempty"`
		DnsRefreshRate *string `tfsdk:"dns_refresh_rate" json:"dnsRefreshRate,omitempty"`
		Failover       *struct {
			Policy *struct {
				OverprovisioningFactor *int64 `tfsdk:"overprovisioning_factor" json:"overprovisioningFactor,omitempty"`
			} `tfsdk:"policy" json:"policy,omitempty"`
			PrioritizedLocalities *[]struct {
				LocalityEndpoints *[]struct {
					LbEndpoints *[]struct {
						Address           *string `tfsdk:"address" json:"address,omitempty"`
						HealthCheckConfig *struct {
							Hostname  *string `tfsdk:"hostname" json:"hostname,omitempty"`
							Method    *string `tfsdk:"method" json:"method,omitempty"`
							Path      *string `tfsdk:"path" json:"path,omitempty"`
							PortValue *int64  `tfsdk:"port_value" json:"portValue,omitempty"`
						} `tfsdk:"health_check_config" json:"healthCheckConfig,omitempty"`
						LoadBalancingWeight *int64 `tfsdk:"load_balancing_weight" json:"loadBalancingWeight,omitempty"`
						Port                *int64 `tfsdk:"port" json:"port,omitempty"`
						UpstreamSslConfig   *struct {
							AllowRenegotiation *bool     `tfsdk:"allow_renegotiation" json:"allowRenegotiation,omitempty"`
							AlpnProtocols      *[]string `tfsdk:"alpn_protocols" json:"alpnProtocols,omitempty"`
							Parameters         *struct {
								CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
								EcdhCurves             *[]string `tfsdk:"ecdh_curves" json:"ecdhCurves,omitempty"`
								MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
								MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
							} `tfsdk:"parameters" json:"parameters,omitempty"`
							Sds *struct {
								CallCredentials *struct {
									FileCredentialSource *struct {
										Header        *string `tfsdk:"header" json:"header,omitempty"`
										TokenFileName *string `tfsdk:"token_file_name" json:"tokenFileName,omitempty"`
									} `tfsdk:"file_credential_source" json:"fileCredentialSource,omitempty"`
								} `tfsdk:"call_credentials" json:"callCredentials,omitempty"`
								CertificatesSecretName *string `tfsdk:"certificates_secret_name" json:"certificatesSecretName,omitempty"`
								ClusterName            *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
								TargetUri              *string `tfsdk:"target_uri" json:"targetUri,omitempty"`
								ValidationContextName  *string `tfsdk:"validation_context_name" json:"validationContextName,omitempty"`
							} `tfsdk:"sds" json:"sds,omitempty"`
							SecretRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							Sni      *string `tfsdk:"sni" json:"sni,omitempty"`
							SslFiles *struct {
								OcspStaple *string `tfsdk:"ocsp_staple" json:"ocspStaple,omitempty"`
								RootCa     *string `tfsdk:"root_ca" json:"rootCa,omitempty"`
								TlsCert    *string `tfsdk:"tls_cert" json:"tlsCert,omitempty"`
								TlsKey     *string `tfsdk:"tls_key" json:"tlsKey,omitempty"`
							} `tfsdk:"ssl_files" json:"sslFiles,omitempty"`
							VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" json:"verifySubjectAltName,omitempty"`
						} `tfsdk:"upstream_ssl_config" json:"upstreamSslConfig,omitempty"`
					} `tfsdk:"lb_endpoints" json:"lbEndpoints,omitempty"`
					LoadBalancingWeight *int64 `tfsdk:"load_balancing_weight" json:"loadBalancingWeight,omitempty"`
					Locality            *struct {
						Region  *string `tfsdk:"region" json:"region,omitempty"`
						SubZone *string `tfsdk:"sub_zone" json:"subZone,omitempty"`
						Zone    *string `tfsdk:"zone" json:"zone,omitempty"`
					} `tfsdk:"locality" json:"locality,omitempty"`
				} `tfsdk:"locality_endpoints" json:"localityEndpoints,omitempty"`
			} `tfsdk:"prioritized_localities" json:"prioritizedLocalities,omitempty"`
		} `tfsdk:"failover" json:"failover,omitempty"`
		HealthChecks *[]struct {
			AlwaysLogHealthCheckFailures *bool `tfsdk:"always_log_health_check_failures" json:"alwaysLogHealthCheckFailures,omitempty"`
			CustomHealthCheck            *struct {
				Config      *map[string]string `tfsdk:"config" json:"config,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				TypedConfig *map[string]string `tfsdk:"typed_config" json:"typedConfig,omitempty"`
			} `tfsdk:"custom_health_check" json:"customHealthCheck,omitempty"`
			EventLogPath    *string `tfsdk:"event_log_path" json:"eventLogPath,omitempty"`
			GrpcHealthCheck *struct {
				Authority       *string `tfsdk:"authority" json:"authority,omitempty"`
				InitialMetadata *[]struct {
					Append *bool `tfsdk:"append" json:"append,omitempty"`
					Header *struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"header" json:"header,omitempty"`
					HeaderSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"header_secret_ref" json:"headerSecretRef,omitempty"`
				} `tfsdk:"initial_metadata" json:"initialMetadata,omitempty"`
				ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
			} `tfsdk:"grpc_health_check" json:"grpcHealthCheck,omitempty"`
			HealthyEdgeInterval *string `tfsdk:"healthy_edge_interval" json:"healthyEdgeInterval,omitempty"`
			HealthyThreshold    *int64  `tfsdk:"healthy_threshold" json:"healthyThreshold,omitempty"`
			HttpHealthCheck     *struct {
				ExpectedStatuses *[]struct {
					End   *int64 `tfsdk:"end" json:"end,omitempty"`
					Start *int64 `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"expected_statuses" json:"expectedStatuses,omitempty"`
				Host                *string `tfsdk:"host" json:"host,omitempty"`
				Method              *string `tfsdk:"method" json:"method,omitempty"`
				Path                *string `tfsdk:"path" json:"path,omitempty"`
				RequestHeadersToAdd *[]struct {
					Append *bool `tfsdk:"append" json:"append,omitempty"`
					Header *struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"header" json:"header,omitempty"`
					HeaderSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"header_secret_ref" json:"headerSecretRef,omitempty"`
				} `tfsdk:"request_headers_to_add" json:"requestHeadersToAdd,omitempty"`
				RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" json:"requestHeadersToRemove,omitempty"`
				ResponseAssertions     *struct {
					NoMatchHealth    *string `tfsdk:"no_match_health" json:"noMatchHealth,omitempty"`
					ResponseMatchers *[]struct {
						MatchHealth   *string `tfsdk:"match_health" json:"matchHealth,omitempty"`
						ResponseMatch *struct {
							Body               *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header             *string            `tfsdk:"header" json:"header,omitempty"`
							IgnoreErrorOnParse *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
							JsonKey            *struct {
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"json_key" json:"jsonKey,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"response_match" json:"responseMatch,omitempty"`
					} `tfsdk:"response_matchers" json:"responseMatchers,omitempty"`
				} `tfsdk:"response_assertions" json:"responseAssertions,omitempty"`
				ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				UseHttp2    *bool   `tfsdk:"use_http2" json:"useHttp2,omitempty"`
			} `tfsdk:"http_health_check" json:"httpHealthCheck,omitempty"`
			InitialJitter         *string `tfsdk:"initial_jitter" json:"initialJitter,omitempty"`
			Interval              *string `tfsdk:"interval" json:"interval,omitempty"`
			IntervalJitter        *string `tfsdk:"interval_jitter" json:"intervalJitter,omitempty"`
			IntervalJitterPercent *int64  `tfsdk:"interval_jitter_percent" json:"intervalJitterPercent,omitempty"`
			NoTrafficInterval     *string `tfsdk:"no_traffic_interval" json:"noTrafficInterval,omitempty"`
			ReuseConnection       *bool   `tfsdk:"reuse_connection" json:"reuseConnection,omitempty"`
			TcpHealthCheck        *struct {
				Receive *[]struct {
					Text *string `tfsdk:"text" json:"text,omitempty"`
				} `tfsdk:"receive" json:"receive,omitempty"`
				Send *struct {
					Text *string `tfsdk:"text" json:"text,omitempty"`
				} `tfsdk:"send" json:"send,omitempty"`
			} `tfsdk:"tcp_health_check" json:"tcpHealthCheck,omitempty"`
			Timeout               *string `tfsdk:"timeout" json:"timeout,omitempty"`
			UnhealthyEdgeInterval *string `tfsdk:"unhealthy_edge_interval" json:"unhealthyEdgeInterval,omitempty"`
			UnhealthyInterval     *string `tfsdk:"unhealthy_interval" json:"unhealthyInterval,omitempty"`
			UnhealthyThreshold    *int64  `tfsdk:"unhealthy_threshold" json:"unhealthyThreshold,omitempty"`
		} `tfsdk:"health_checks" json:"healthChecks,omitempty"`
		HttpConnectHeaders *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"http_connect_headers" json:"httpConnectHeaders,omitempty"`
		HttpConnectSslConfig *struct {
			AllowRenegotiation *bool     `tfsdk:"allow_renegotiation" json:"allowRenegotiation,omitempty"`
			AlpnProtocols      *[]string `tfsdk:"alpn_protocols" json:"alpnProtocols,omitempty"`
			Parameters         *struct {
				CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				EcdhCurves             *[]string `tfsdk:"ecdh_curves" json:"ecdhCurves,omitempty"`
				MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
				MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			Sds *struct {
				CallCredentials *struct {
					FileCredentialSource *struct {
						Header        *string `tfsdk:"header" json:"header,omitempty"`
						TokenFileName *string `tfsdk:"token_file_name" json:"tokenFileName,omitempty"`
					} `tfsdk:"file_credential_source" json:"fileCredentialSource,omitempty"`
				} `tfsdk:"call_credentials" json:"callCredentials,omitempty"`
				CertificatesSecretName *string `tfsdk:"certificates_secret_name" json:"certificatesSecretName,omitempty"`
				ClusterName            *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
				TargetUri              *string `tfsdk:"target_uri" json:"targetUri,omitempty"`
				ValidationContextName  *string `tfsdk:"validation_context_name" json:"validationContextName,omitempty"`
			} `tfsdk:"sds" json:"sds,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			Sni      *string `tfsdk:"sni" json:"sni,omitempty"`
			SslFiles *struct {
				OcspStaple *string `tfsdk:"ocsp_staple" json:"ocspStaple,omitempty"`
				RootCa     *string `tfsdk:"root_ca" json:"rootCa,omitempty"`
				TlsCert    *string `tfsdk:"tls_cert" json:"tlsCert,omitempty"`
				TlsKey     *string `tfsdk:"tls_key" json:"tlsKey,omitempty"`
			} `tfsdk:"ssl_files" json:"sslFiles,omitempty"`
			VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" json:"verifySubjectAltName,omitempty"`
		} `tfsdk:"http_connect_ssl_config" json:"httpConnectSslConfig,omitempty"`
		HttpProxyHostname           *string `tfsdk:"http_proxy_hostname" json:"httpProxyHostname,omitempty"`
		IgnoreHealthOnHostRemoval   *bool   `tfsdk:"ignore_health_on_host_removal" json:"ignoreHealthOnHostRemoval,omitempty"`
		InitialConnectionWindowSize *int64  `tfsdk:"initial_connection_window_size" json:"initialConnectionWindowSize,omitempty"`
		InitialStreamWindowSize     *int64  `tfsdk:"initial_stream_window_size" json:"initialStreamWindowSize,omitempty"`
		Kube                        *struct {
			Selector         *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
			ServiceName      *string            `tfsdk:"service_name" json:"serviceName,omitempty"`
			ServiceNamespace *string            `tfsdk:"service_namespace" json:"serviceNamespace,omitempty"`
			ServicePort      *int64             `tfsdk:"service_port" json:"servicePort,omitempty"`
			ServiceSpec      *struct {
				Graphql *struct {
					Endpoint *struct {
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"graphql" json:"graphql,omitempty"`
				Grpc *struct {
					Descriptors  *string `tfsdk:"descriptors" json:"descriptors,omitempty"`
					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" json:"functionNames,omitempty"`
						PackageName   *string   `tfsdk:"package_name" json:"packageName,omitempty"`
						ServiceName   *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" json:"grpcServices,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				GrpcJsonTranscoder *struct {
					AutoMapping                  *bool     `tfsdk:"auto_mapping" json:"autoMapping,omitempty"`
					ConvertGrpcStatus            *bool     `tfsdk:"convert_grpc_status" json:"convertGrpcStatus,omitempty"`
					IgnoreUnknownQueryParameters *bool     `tfsdk:"ignore_unknown_query_parameters" json:"ignoreUnknownQueryParameters,omitempty"`
					IgnoredQueryParameters       *[]string `tfsdk:"ignored_query_parameters" json:"ignoredQueryParameters,omitempty"`
					MatchIncomingRequestRoute    *bool     `tfsdk:"match_incoming_request_route" json:"matchIncomingRequestRoute,omitempty"`
					PrintOptions                 *struct {
						AddWhitespace              *bool `tfsdk:"add_whitespace" json:"addWhitespace,omitempty"`
						AlwaysPrintEnumsAsInts     *bool `tfsdk:"always_print_enums_as_ints" json:"alwaysPrintEnumsAsInts,omitempty"`
						AlwaysPrintPrimitiveFields *bool `tfsdk:"always_print_primitive_fields" json:"alwaysPrintPrimitiveFields,omitempty"`
						PreserveProtoFieldNames    *bool `tfsdk:"preserve_proto_field_names" json:"preserveProtoFieldNames,omitempty"`
					} `tfsdk:"print_options" json:"printOptions,omitempty"`
					ProtoDescriptor          *string `tfsdk:"proto_descriptor" json:"protoDescriptor,omitempty"`
					ProtoDescriptorBin       *string `tfsdk:"proto_descriptor_bin" json:"protoDescriptorBin,omitempty"`
					ProtoDescriptorConfigMap *struct {
						ConfigMapRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						Key *string `tfsdk:"key" json:"key,omitempty"`
					} `tfsdk:"proto_descriptor_config_map" json:"protoDescriptorConfigMap,omitempty"`
					Services *[]string `tfsdk:"services" json:"services,omitempty"`
				} `tfsdk:"grpc_json_transcoder" json:"grpcJsonTranscoder,omitempty"`
				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" json:"inline,omitempty"`
						Url    *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"swagger_info" json:"swaggerInfo,omitempty"`
					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
						Body              *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						DynamicMetadataValues *[]struct {
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
							Value             *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
						EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						Extractors       *struct {
							Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header   *string            `tfsdk:"header" json:"header,omitempty"`
							Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
							Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"extractors" json:"extractors,omitempty"`
						Headers *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						HeadersToAppend *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
						HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
						IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
						ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
						Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
					} `tfsdk:"transformations" json:"transformations,omitempty"`
				} `tfsdk:"rest" json:"rest,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			SubsetSpec *struct {
				DefaultSubset *struct {
					Values *map[string]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"default_subset" json:"defaultSubset,omitempty"`
				FallbackPolicy *string `tfsdk:"fallback_policy" json:"fallbackPolicy,omitempty"`
				Selectors      *[]struct {
					Keys                *[]string `tfsdk:"keys" json:"keys,omitempty"`
					SingleHostPerSubset *bool     `tfsdk:"single_host_per_subset" json:"singleHostPerSubset,omitempty"`
				} `tfsdk:"selectors" json:"selectors,omitempty"`
			} `tfsdk:"subset_spec" json:"subsetSpec,omitempty"`
		} `tfsdk:"kube" json:"kube,omitempty"`
		LoadBalancerConfig *struct {
			HealthyPanicThreshold *big.Float `tfsdk:"healthy_panic_threshold" json:"healthyPanicThreshold,omitempty"`
			LeastRequest          *struct {
				ChoiceCount     *int64 `tfsdk:"choice_count" json:"choiceCount,omitempty"`
				SlowStartConfig *struct {
					Aggression       *big.Float `tfsdk:"aggression" json:"aggression,omitempty"`
					MinWeightPercent *big.Float `tfsdk:"min_weight_percent" json:"minWeightPercent,omitempty"`
					SlowStartWindow  *string    `tfsdk:"slow_start_window" json:"slowStartWindow,omitempty"`
				} `tfsdk:"slow_start_config" json:"slowStartConfig,omitempty"`
			} `tfsdk:"least_request" json:"leastRequest,omitempty"`
			LocalityWeightedLbConfig *map[string]string `tfsdk:"locality_weighted_lb_config" json:"localityWeightedLbConfig,omitempty"`
			Maglev                   *map[string]string `tfsdk:"maglev" json:"maglev,omitempty"`
			Random                   *map[string]string `tfsdk:"random" json:"random,omitempty"`
			RingHash                 *struct {
				RingHashConfig *struct {
					MaximumRingSize *int64 `tfsdk:"maximum_ring_size" json:"maximumRingSize,omitempty"`
					MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
				} `tfsdk:"ring_hash_config" json:"ringHashConfig,omitempty"`
			} `tfsdk:"ring_hash" json:"ringHash,omitempty"`
			RoundRobin *struct {
				SlowStartConfig *struct {
					Aggression       *big.Float `tfsdk:"aggression" json:"aggression,omitempty"`
					MinWeightPercent *big.Float `tfsdk:"min_weight_percent" json:"minWeightPercent,omitempty"`
					SlowStartWindow  *string    `tfsdk:"slow_start_window" json:"slowStartWindow,omitempty"`
				} `tfsdk:"slow_start_config" json:"slowStartConfig,omitempty"`
			} `tfsdk:"round_robin" json:"roundRobin,omitempty"`
			UpdateMergeWindow *string `tfsdk:"update_merge_window" json:"updateMergeWindow,omitempty"`
		} `tfsdk:"load_balancer_config" json:"loadBalancerConfig,omitempty"`
		MaxConcurrentStreams *int64 `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
		NamespacedStatuses   *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
		OutlierDetection *struct {
			BaseEjectionTime                       *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
			Consecutive5xx                         *int64  `tfsdk:"consecutive5xx" json:"consecutive5xx,omitempty"`
			ConsecutiveGatewayFailure              *int64  `tfsdk:"consecutive_gateway_failure" json:"consecutiveGatewayFailure,omitempty"`
			ConsecutiveLocalOriginFailure          *int64  `tfsdk:"consecutive_local_origin_failure" json:"consecutiveLocalOriginFailure,omitempty"`
			EnforcingConsecutive5xx                *int64  `tfsdk:"enforcing_consecutive5xx" json:"enforcingConsecutive5xx,omitempty"`
			EnforcingConsecutiveGatewayFailure     *int64  `tfsdk:"enforcing_consecutive_gateway_failure" json:"enforcingConsecutiveGatewayFailure,omitempty"`
			EnforcingConsecutiveLocalOriginFailure *int64  `tfsdk:"enforcing_consecutive_local_origin_failure" json:"enforcingConsecutiveLocalOriginFailure,omitempty"`
			EnforcingLocalOriginSuccessRate        *int64  `tfsdk:"enforcing_local_origin_success_rate" json:"enforcingLocalOriginSuccessRate,omitempty"`
			EnforcingSuccessRate                   *int64  `tfsdk:"enforcing_success_rate" json:"enforcingSuccessRate,omitempty"`
			Interval                               *string `tfsdk:"interval" json:"interval,omitempty"`
			MaxEjectionPercent                     *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
			SplitExternalLocalOriginErrors         *bool   `tfsdk:"split_external_local_origin_errors" json:"splitExternalLocalOriginErrors,omitempty"`
			SuccessRateMinimumHosts                *int64  `tfsdk:"success_rate_minimum_hosts" json:"successRateMinimumHosts,omitempty"`
			SuccessRateRequestVolume               *int64  `tfsdk:"success_rate_request_volume" json:"successRateRequestVolume,omitempty"`
			SuccessRateStdevFactor                 *int64  `tfsdk:"success_rate_stdev_factor" json:"successRateStdevFactor,omitempty"`
		} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
		OverrideStreamErrorOnInvalidHttpMessage *bool `tfsdk:"override_stream_error_on_invalid_http_message" json:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`
		Pipe                                    *struct {
			Path        *string `tfsdk:"path" json:"path,omitempty"`
			ServiceSpec *struct {
				Graphql *struct {
					Endpoint *struct {
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"graphql" json:"graphql,omitempty"`
				Grpc *struct {
					Descriptors  *string `tfsdk:"descriptors" json:"descriptors,omitempty"`
					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" json:"functionNames,omitempty"`
						PackageName   *string   `tfsdk:"package_name" json:"packageName,omitempty"`
						ServiceName   *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" json:"grpcServices,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				GrpcJsonTranscoder *struct {
					AutoMapping                  *bool     `tfsdk:"auto_mapping" json:"autoMapping,omitempty"`
					ConvertGrpcStatus            *bool     `tfsdk:"convert_grpc_status" json:"convertGrpcStatus,omitempty"`
					IgnoreUnknownQueryParameters *bool     `tfsdk:"ignore_unknown_query_parameters" json:"ignoreUnknownQueryParameters,omitempty"`
					IgnoredQueryParameters       *[]string `tfsdk:"ignored_query_parameters" json:"ignoredQueryParameters,omitempty"`
					MatchIncomingRequestRoute    *bool     `tfsdk:"match_incoming_request_route" json:"matchIncomingRequestRoute,omitempty"`
					PrintOptions                 *struct {
						AddWhitespace              *bool `tfsdk:"add_whitespace" json:"addWhitespace,omitempty"`
						AlwaysPrintEnumsAsInts     *bool `tfsdk:"always_print_enums_as_ints" json:"alwaysPrintEnumsAsInts,omitempty"`
						AlwaysPrintPrimitiveFields *bool `tfsdk:"always_print_primitive_fields" json:"alwaysPrintPrimitiveFields,omitempty"`
						PreserveProtoFieldNames    *bool `tfsdk:"preserve_proto_field_names" json:"preserveProtoFieldNames,omitempty"`
					} `tfsdk:"print_options" json:"printOptions,omitempty"`
					ProtoDescriptor          *string `tfsdk:"proto_descriptor" json:"protoDescriptor,omitempty"`
					ProtoDescriptorBin       *string `tfsdk:"proto_descriptor_bin" json:"protoDescriptorBin,omitempty"`
					ProtoDescriptorConfigMap *struct {
						ConfigMapRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						Key *string `tfsdk:"key" json:"key,omitempty"`
					} `tfsdk:"proto_descriptor_config_map" json:"protoDescriptorConfigMap,omitempty"`
					Services *[]string `tfsdk:"services" json:"services,omitempty"`
				} `tfsdk:"grpc_json_transcoder" json:"grpcJsonTranscoder,omitempty"`
				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" json:"inline,omitempty"`
						Url    *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"swagger_info" json:"swaggerInfo,omitempty"`
					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
						Body              *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						DynamicMetadataValues *[]struct {
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
							Value             *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
						EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						Extractors       *struct {
							Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header   *string            `tfsdk:"header" json:"header,omitempty"`
							Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
							Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"extractors" json:"extractors,omitempty"`
						Headers *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						HeadersToAppend *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
						HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
						IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
						ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
						Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
					} `tfsdk:"transformations" json:"transformations,omitempty"`
				} `tfsdk:"rest" json:"rest,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
		} `tfsdk:"pipe" json:"pipe,omitempty"`
		ProtocolSelection    *string `tfsdk:"protocol_selection" json:"protocolSelection,omitempty"`
		ProxyProtocolVersion *string `tfsdk:"proxy_protocol_version" json:"proxyProtocolVersion,omitempty"`
		RespectDnsTtl        *bool   `tfsdk:"respect_dns_ttl" json:"respectDnsTtl,omitempty"`
		SslConfig            *struct {
			AllowRenegotiation *bool     `tfsdk:"allow_renegotiation" json:"allowRenegotiation,omitempty"`
			AlpnProtocols      *[]string `tfsdk:"alpn_protocols" json:"alpnProtocols,omitempty"`
			Parameters         *struct {
				CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				EcdhCurves             *[]string `tfsdk:"ecdh_curves" json:"ecdhCurves,omitempty"`
				MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
				MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			Sds *struct {
				CallCredentials *struct {
					FileCredentialSource *struct {
						Header        *string `tfsdk:"header" json:"header,omitempty"`
						TokenFileName *string `tfsdk:"token_file_name" json:"tokenFileName,omitempty"`
					} `tfsdk:"file_credential_source" json:"fileCredentialSource,omitempty"`
				} `tfsdk:"call_credentials" json:"callCredentials,omitempty"`
				CertificatesSecretName *string `tfsdk:"certificates_secret_name" json:"certificatesSecretName,omitempty"`
				ClusterName            *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
				TargetUri              *string `tfsdk:"target_uri" json:"targetUri,omitempty"`
				ValidationContextName  *string `tfsdk:"validation_context_name" json:"validationContextName,omitempty"`
			} `tfsdk:"sds" json:"sds,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			Sni      *string `tfsdk:"sni" json:"sni,omitempty"`
			SslFiles *struct {
				OcspStaple *string `tfsdk:"ocsp_staple" json:"ocspStaple,omitempty"`
				RootCa     *string `tfsdk:"root_ca" json:"rootCa,omitempty"`
				TlsCert    *string `tfsdk:"tls_cert" json:"tlsCert,omitempty"`
				TlsKey     *string `tfsdk:"tls_key" json:"tlsKey,omitempty"`
			} `tfsdk:"ssl_files" json:"sslFiles,omitempty"`
			VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" json:"verifySubjectAltName,omitempty"`
		} `tfsdk:"ssl_config" json:"sslConfig,omitempty"`
		Static *struct {
			AutoSniRewrite *bool `tfsdk:"auto_sni_rewrite" json:"autoSniRewrite,omitempty"`
			Hosts          *[]struct {
				Addr              *string `tfsdk:"addr" json:"addr,omitempty"`
				HealthCheckConfig *struct {
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Path   *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"health_check_config" json:"healthCheckConfig,omitempty"`
				LoadBalancingWeight *int64  `tfsdk:"load_balancing_weight" json:"loadBalancingWeight,omitempty"`
				Port                *int64  `tfsdk:"port" json:"port,omitempty"`
				SniAddr             *string `tfsdk:"sni_addr" json:"sniAddr,omitempty"`
			} `tfsdk:"hosts" json:"hosts,omitempty"`
			ServiceSpec *struct {
				Graphql *struct {
					Endpoint *struct {
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"graphql" json:"graphql,omitempty"`
				Grpc *struct {
					Descriptors  *string `tfsdk:"descriptors" json:"descriptors,omitempty"`
					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" json:"functionNames,omitempty"`
						PackageName   *string   `tfsdk:"package_name" json:"packageName,omitempty"`
						ServiceName   *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" json:"grpcServices,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				GrpcJsonTranscoder *struct {
					AutoMapping                  *bool     `tfsdk:"auto_mapping" json:"autoMapping,omitempty"`
					ConvertGrpcStatus            *bool     `tfsdk:"convert_grpc_status" json:"convertGrpcStatus,omitempty"`
					IgnoreUnknownQueryParameters *bool     `tfsdk:"ignore_unknown_query_parameters" json:"ignoreUnknownQueryParameters,omitempty"`
					IgnoredQueryParameters       *[]string `tfsdk:"ignored_query_parameters" json:"ignoredQueryParameters,omitempty"`
					MatchIncomingRequestRoute    *bool     `tfsdk:"match_incoming_request_route" json:"matchIncomingRequestRoute,omitempty"`
					PrintOptions                 *struct {
						AddWhitespace              *bool `tfsdk:"add_whitespace" json:"addWhitespace,omitempty"`
						AlwaysPrintEnumsAsInts     *bool `tfsdk:"always_print_enums_as_ints" json:"alwaysPrintEnumsAsInts,omitempty"`
						AlwaysPrintPrimitiveFields *bool `tfsdk:"always_print_primitive_fields" json:"alwaysPrintPrimitiveFields,omitempty"`
						PreserveProtoFieldNames    *bool `tfsdk:"preserve_proto_field_names" json:"preserveProtoFieldNames,omitempty"`
					} `tfsdk:"print_options" json:"printOptions,omitempty"`
					ProtoDescriptor          *string `tfsdk:"proto_descriptor" json:"protoDescriptor,omitempty"`
					ProtoDescriptorBin       *string `tfsdk:"proto_descriptor_bin" json:"protoDescriptorBin,omitempty"`
					ProtoDescriptorConfigMap *struct {
						ConfigMapRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						Key *string `tfsdk:"key" json:"key,omitempty"`
					} `tfsdk:"proto_descriptor_config_map" json:"protoDescriptorConfigMap,omitempty"`
					Services *[]string `tfsdk:"services" json:"services,omitempty"`
				} `tfsdk:"grpc_json_transcoder" json:"grpcJsonTranscoder,omitempty"`
				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" json:"inline,omitempty"`
						Url    *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"swagger_info" json:"swaggerInfo,omitempty"`
					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
						Body              *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						DynamicMetadataValues *[]struct {
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
							Value             *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
						EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						Extractors       *struct {
							Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header   *string            `tfsdk:"header" json:"header,omitempty"`
							Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
							Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"extractors" json:"extractors,omitempty"`
						Headers *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						HeadersToAppend *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
						HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
						IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
						ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
						Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
					} `tfsdk:"transformations" json:"transformations,omitempty"`
				} `tfsdk:"rest" json:"rest,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			UseTls *bool `tfsdk:"use_tls" json:"useTls,omitempty"`
		} `tfsdk:"static" json:"static,omitempty"`
		UseHttp2 *bool `tfsdk:"use_http2" json:"useHttp2,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GlooSoloIoUpstreamV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gloo_solo_io_upstream_v1"
}

func (r *GlooSoloIoUpstreamV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"aws": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"aws_account_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"destination_overrides": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"invocation_style": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"logical_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_transformation": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_transformation": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"unwrap_as_alb": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"unwrap_as_api_gateway": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"wrap_as_api_gateway": schema.BoolAttribute{
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

							"disable_role_chaining": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"lambda_functions": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"lambda_function_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logical_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"qualifier": schema.StringAttribute{
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

							"region": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

					"aws_ec2": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"filters": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kv_pair": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
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

							"port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"public_ip": schema.BoolAttribute{
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

							"role_arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

					"azure": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"function_app_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"functions": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"auth_level": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"function_name": schema.StringAttribute{
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

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

					"circuit_breakers": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"max_connections": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"max_pending_requests": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"max_requests": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"max_retries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"connection_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"common_http_protocol_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"headers_with_underscores_action": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"idle_timeout": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_headers_count": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_stream_duration": schema.StringAttribute{
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

							"connect_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http1_protocol_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enable_trailers": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"override_stream_error_on_invalid_http_message": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"preserve_case_header_key_format": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proper_case_header_key_format": schema.BoolAttribute{
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

							"max_requests_per_connection": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"per_connection_buffer_limit_bytes": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"tcp_keepalive": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"keepalive_interval": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_probes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_time": schema.StringAttribute{
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

					"consul": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"connect_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"consistency_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data_centers": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_blacklist_tags": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_tags": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"query_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"use_cache": schema.BoolAttribute{
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

							"service_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"graphql": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"endpoint": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"url": schema.StringAttribute{
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

									"grpc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"descriptors": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"grpc_services": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"function_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"package_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_name": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc_json_transcoder": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"auto_mapping": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"convert_grpc_status": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignore_unknown_query_parameters": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignored_query_parameters": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_incoming_request_route": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"print_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_whitespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_enums_as_ints": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_primitive_fields": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preserve_proto_field_names": schema.BoolAttribute{
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

											"proto_descriptor": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proto_descriptor_bin": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"proto_descriptor_config_map": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"key": schema.StringAttribute{
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

											"services": schema.ListAttribute{
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

									"rest": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"swagger_info": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"inline": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"transformations": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"advanced_templates": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"body": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"dynamic_metadata_values": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"metadata_namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"escape_characters": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"extractors": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"header": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subgroup": schema.Int64Attribute{
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

													"headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"headers_to_append": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers_to_remove": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_error_on_parse": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"merge_extractors_to_body": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"parse_body_behavior": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"passthrough": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_tags": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subset_tags": schema.ListAttribute{
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

					"discovery_metadata": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
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

					"dns_refresh_rate": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failover": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"policy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"overprovisioning_factor": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(4.294967295e+09),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"prioritized_localities": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"locality_endpoints": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"lb_endpoints": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"address": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"health_check_config": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"hostname": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"method": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port_value": schema.Int64Attribute{
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

																"load_balancing_weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(4.294967295e+09),
																	},
																},

																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"upstream_ssl_config": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_renegotiation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"alpn_protocols": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parameters": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"cipher_suites": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ecdh_curves": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"maximum_protocol_version": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"minimum_protocol_version": schema.StringAttribute{
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

																		"sds": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"call_credentials": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"file_credential_source": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"header": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"token_file_name": schema.StringAttribute{
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

																				"certificates_secret_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cluster_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_uri": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"validation_context_name": schema.StringAttribute{
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

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace": schema.StringAttribute{
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

																		"sni": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ssl_files": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"ocsp_staple": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"root_ca": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"tls_cert": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"tls_key": schema.StringAttribute{
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

																		"verify_subject_alt_name": schema.ListAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"load_balancing_weight": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4.294967295e+09),
														},
													},

													"locality": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_zone": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"zone": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_checks": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"always_log_health_check_failures": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"custom_health_check": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
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

										"typed_config": schema.MapAttribute{
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

								"event_log_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"grpc_health_check": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"authority": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"initial_metadata": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"append": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"header_secret_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"service_name": schema.StringAttribute{
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

								"healthy_edge_interval": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"healthy_threshold": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(4.294967295e+09),
									},
								},

								"http_health_check": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expected_statuses": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"start": schema.Int64Attribute{
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

										"host": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"method": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"request_headers_to_add": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"append": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"header_secret_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"request_headers_to_remove": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_assertions": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"no_match_health": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"response_matchers": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"match_health": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"response_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"body": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"ignore_error_on_parse": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"json_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"path": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"service_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_http2": schema.BoolAttribute{
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

								"initial_jitter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval_jitter": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval_jitter_percent": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"no_traffic_interval": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"reuse_connection": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tcp_health_check": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"receive": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"text": schema.StringAttribute{
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

										"send": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"text": schema.StringAttribute{
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

								"timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"unhealthy_edge_interval": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"unhealthy_interval": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"unhealthy_threshold": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(4.294967295e+09),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_connect_headers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
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

					"http_connect_ssl_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"allow_renegotiation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"alpn_protocols": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parameters": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cipher_suites": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ecdh_curves": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_protocol_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"minimum_protocol_version": schema.StringAttribute{
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

							"sds": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"call_credentials": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"file_credential_source": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"header": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"token_file_name": schema.StringAttribute{
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

									"certificates_secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_uri": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"validation_context_name": schema.StringAttribute{
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

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

							"sni": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_files": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"ocsp_staple": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"root_ca": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_cert": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_key": schema.StringAttribute{
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

							"verify_subject_alt_name": schema.ListAttribute{
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

					"http_proxy_hostname": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ignore_health_on_host_removal": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"initial_connection_window_size": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4.294967295e+09),
						},
					},

					"initial_stream_window_size": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4.294967295e+09),
						},
					},

					"kube": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"selector": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"graphql": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"endpoint": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"url": schema.StringAttribute{
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

									"grpc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"descriptors": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"grpc_services": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"function_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"package_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_name": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc_json_transcoder": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"auto_mapping": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"convert_grpc_status": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignore_unknown_query_parameters": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignored_query_parameters": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_incoming_request_route": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"print_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_whitespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_enums_as_ints": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_primitive_fields": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preserve_proto_field_names": schema.BoolAttribute{
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

											"proto_descriptor": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proto_descriptor_bin": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"proto_descriptor_config_map": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"key": schema.StringAttribute{
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

											"services": schema.ListAttribute{
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

									"rest": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"swagger_info": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"inline": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"transformations": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"advanced_templates": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"body": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"dynamic_metadata_values": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"metadata_namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"escape_characters": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"extractors": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"header": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subgroup": schema.Int64Attribute{
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

													"headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"headers_to_append": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers_to_remove": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_error_on_parse": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"merge_extractors_to_body": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"parse_body_behavior": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"passthrough": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"subset_spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"default_subset": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"values": schema.MapAttribute{
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

									"fallback_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selectors": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"keys": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"single_host_per_subset": schema.BoolAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_balancer_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"healthy_panic_threshold": types.NumberType{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"least_request": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"choice_count": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"slow_start_config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"aggression": types.NumberType{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"min_weight_percent": types.NumberType{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"slow_start_window": schema.StringAttribute{
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

							"locality_weighted_lb_config": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maglev": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"random": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ring_hash": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"ring_hash_config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"maximum_ring_size": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"minimum_ring_size": schema.Int64Attribute{
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

							"round_robin": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"slow_start_config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"aggression": types.NumberType{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"min_weight_percent": types.NumberType{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"slow_start_window": schema.StringAttribute{
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

							"update_merge_window": schema.StringAttribute{
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

					"max_concurrent_streams": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4.294967295e+09),
						},
					},

					"namespaced_statuses": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"statuses": schema.MapAttribute{
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

					"outlier_detection": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"base_ejection_time": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"consecutive5xx": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"consecutive_gateway_failure": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"consecutive_local_origin_failure": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_consecutive5xx": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_consecutive_gateway_failure": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_consecutive_local_origin_failure": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_local_origin_success_rate": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_success_rate": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_ejection_percent": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"split_external_local_origin_errors": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_rate_minimum_hosts": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"success_rate_request_volume": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},

							"success_rate_stdev_factor": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(4.294967295e+09),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"override_stream_error_on_invalid_http_message": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pipe": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"graphql": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"endpoint": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"url": schema.StringAttribute{
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

									"grpc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"descriptors": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"grpc_services": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"function_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"package_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_name": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc_json_transcoder": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"auto_mapping": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"convert_grpc_status": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignore_unknown_query_parameters": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignored_query_parameters": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_incoming_request_route": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"print_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_whitespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_enums_as_ints": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_primitive_fields": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preserve_proto_field_names": schema.BoolAttribute{
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

											"proto_descriptor": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proto_descriptor_bin": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"proto_descriptor_config_map": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"key": schema.StringAttribute{
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

											"services": schema.ListAttribute{
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

									"rest": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"swagger_info": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"inline": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"transformations": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"advanced_templates": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"body": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"dynamic_metadata_values": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"metadata_namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"escape_characters": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"extractors": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"header": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subgroup": schema.Int64Attribute{
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

													"headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"headers_to_append": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers_to_remove": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_error_on_parse": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"merge_extractors_to_body": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"parse_body_behavior": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"passthrough": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol_selection": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy_protocol_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"respect_dns_ttl": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ssl_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"allow_renegotiation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"alpn_protocols": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parameters": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cipher_suites": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ecdh_curves": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_protocol_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"minimum_protocol_version": schema.StringAttribute{
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

							"sds": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"call_credentials": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"file_credential_source": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"header": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"token_file_name": schema.StringAttribute{
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

									"certificates_secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_uri": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"validation_context_name": schema.StringAttribute{
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

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

							"sni": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_files": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"ocsp_staple": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"root_ca": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_cert": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_key": schema.StringAttribute{
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

							"verify_subject_alt_name": schema.ListAttribute{
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

					"static": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"auto_sni_rewrite": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hosts": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"addr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"health_check_config": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"method": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
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

										"load_balancing_weight": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(4.294967295e+09),
											},
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sni_addr": schema.StringAttribute{
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

							"service_spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"graphql": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"endpoint": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"url": schema.StringAttribute{
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

									"grpc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"descriptors": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"grpc_services": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"function_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"package_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_name": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc_json_transcoder": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"auto_mapping": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"convert_grpc_status": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignore_unknown_query_parameters": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ignored_query_parameters": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_incoming_request_route": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"print_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_whitespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_enums_as_ints": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"always_print_primitive_fields": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preserve_proto_field_names": schema.BoolAttribute{
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

											"proto_descriptor": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proto_descriptor_bin": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"proto_descriptor_config_map": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"key": schema.StringAttribute{
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

											"services": schema.ListAttribute{
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

									"rest": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"swagger_info": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"inline": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"transformations": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"advanced_templates": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"body": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"dynamic_metadata_values": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"metadata_namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"escape_characters": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"extractors": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"header": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subgroup": schema.Int64Attribute{
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

													"headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"headers_to_append": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers_to_remove": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_error_on_parse": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"merge_extractors_to_body": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"parse_body_behavior": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"passthrough": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_tls": schema.BoolAttribute{
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

					"use_http2": schema.BoolAttribute{
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
	}
}

func (r *GlooSoloIoUpstreamV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *GlooSoloIoUpstreamV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gloo_solo_io_upstream_v1")

	var model GlooSoloIoUpstreamV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("gloo.solo.io/v1")
	model.Kind = pointer.String("Upstream")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "Upstream"}).
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

	var readResponse GlooSoloIoUpstreamV1ResourceData
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

func (r *GlooSoloIoUpstreamV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_upstream_v1")

	var data GlooSoloIoUpstreamV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "Upstream"}).
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

	var readResponse GlooSoloIoUpstreamV1ResourceData
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

func (r *GlooSoloIoUpstreamV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gloo_solo_io_upstream_v1")

	var model GlooSoloIoUpstreamV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gloo.solo.io/v1")
	model.Kind = pointer.String("Upstream")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "Upstream"}).
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

	var readResponse GlooSoloIoUpstreamV1ResourceData
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

func (r *GlooSoloIoUpstreamV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gloo_solo_io_upstream_v1")

	var data GlooSoloIoUpstreamV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "Upstream"}).
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

func (r *GlooSoloIoUpstreamV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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