/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_solo_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &GatewaySoloIoMatchableHttpGatewayV1DataSource{}
	_ datasource.DataSourceWithConfigure = &GatewaySoloIoMatchableHttpGatewayV1DataSource{}
)

func NewGatewaySoloIoMatchableHttpGatewayV1DataSource() datasource.DataSource {
	return &GatewaySoloIoMatchableHttpGatewayV1DataSource{}
}

type GatewaySoloIoMatchableHttpGatewayV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GatewaySoloIoMatchableHttpGatewayV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		HttpGateway *struct {
			Options *struct {
				Buffer *struct {
					MaxRequestBytes *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
				} `tfsdk:"buffer" json:"buffer,omitempty"`
				Caching *struct {
					AllowedVaryHeaders *[]struct {
						Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
						IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
						Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
						SafeRegex  *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
						Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
					} `tfsdk:"allowed_vary_headers" json:"allowedVaryHeaders,omitempty"`
					CachingServiceRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"caching_service_ref" json:"cachingServiceRef,omitempty"`
					MaxPayloadSize *int64  `tfsdk:"max_payload_size" json:"maxPayloadSize,omitempty"`
					Timeout        *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"caching" json:"caching,omitempty"`
				ConnectionLimit *struct {
					DelayBeforeClose     *string `tfsdk:"delay_before_close" json:"delayBeforeClose,omitempty"`
					MaxActiveConnections *int64  `tfsdk:"max_active_connections" json:"maxActiveConnections,omitempty"`
				} `tfsdk:"connection_limit" json:"connectionLimit,omitempty"`
				Csrf *struct {
					AdditionalOrigins *[]struct {
						Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
						IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
						Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
						SafeRegex  *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
						Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
					} `tfsdk:"additional_origins" json:"additionalOrigins,omitempty"`
					FilterEnabled *struct {
						DefaultValue *struct {
							Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
							Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
						} `tfsdk:"default_value" json:"defaultValue,omitempty"`
						RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
					} `tfsdk:"filter_enabled" json:"filterEnabled,omitempty"`
					ShadowEnabled *struct {
						DefaultValue *struct {
							Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
							Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
						} `tfsdk:"default_value" json:"defaultValue,omitempty"`
						RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
					} `tfsdk:"shadow_enabled" json:"shadowEnabled,omitempty"`
				} `tfsdk:"csrf" json:"csrf,omitempty"`
				DisableExtProc *bool `tfsdk:"disable_ext_proc" json:"disableExtProc,omitempty"`
				Dlp            *struct {
					DlpRules *[]struct {
						Actions *[]struct {
							ActionType   *string `tfsdk:"action_type" json:"actionType,omitempty"`
							CustomAction *struct {
								MaskChar *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Percent  *struct {
									Value *float64 `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"percent" json:"percent,omitempty"`
								Regex        *[]string `tfsdk:"regex" json:"regex,omitempty"`
								RegexActions *[]struct {
									Regex    *string `tfsdk:"regex" json:"regex,omitempty"`
									Subgroup *int64  `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"regex_actions" json:"regexActions,omitempty"`
							} `tfsdk:"custom_action" json:"customAction,omitempty"`
							KeyValueAction *struct {
								KeyToMask *string `tfsdk:"key_to_mask" json:"keyToMask,omitempty"`
								MaskChar  *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Percent   *struct {
									Value *float64 `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"percent" json:"percent,omitempty"`
							} `tfsdk:"key_value_action" json:"keyValueAction,omitempty"`
							Shadow *bool `tfsdk:"shadow" json:"shadow,omitempty"`
						} `tfsdk:"actions" json:"actions,omitempty"`
						Matcher *struct {
							CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
							ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
							Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
							Headers        *[]struct {
								InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value       *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
							Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
							QueryParameters *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"matcher" json:"matcher,omitempty"`
					} `tfsdk:"dlp_rules" json:"dlpRules,omitempty"`
					EnabledFor *string `tfsdk:"enabled_for" json:"enabledFor,omitempty"`
				} `tfsdk:"dlp" json:"dlp,omitempty"`
				DynamicForwardProxy *struct {
					DnsCacheConfig *struct {
						AppleDns *map[string]string `tfsdk:"apple_dns" json:"appleDns,omitempty"`
						CaresDns *struct {
							DnsResolverOptions *struct {
								NoDefaultSearchDomain *bool `tfsdk:"no_default_search_domain" json:"noDefaultSearchDomain,omitempty"`
								UseTcpForDnsLookups   *bool `tfsdk:"use_tcp_for_dns_lookups" json:"useTcpForDnsLookups,omitempty"`
							} `tfsdk:"dns_resolver_options" json:"dnsResolverOptions,omitempty"`
							Resolvers *[]struct {
								Pipe *struct {
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"pipe" json:"pipe,omitempty"`
								SocketAddress *struct {
									Address      *string `tfsdk:"address" json:"address,omitempty"`
									Ipv4Compat   *bool   `tfsdk:"ipv4_compat" json:"ipv4Compat,omitempty"`
									NamedPort    *string `tfsdk:"named_port" json:"namedPort,omitempty"`
									PortValue    *int64  `tfsdk:"port_value" json:"portValue,omitempty"`
									Protocol     *string `tfsdk:"protocol" json:"protocol,omitempty"`
									ResolverName *string `tfsdk:"resolver_name" json:"resolverName,omitempty"`
								} `tfsdk:"socket_address" json:"socketAddress,omitempty"`
							} `tfsdk:"resolvers" json:"resolvers,omitempty"`
						} `tfsdk:"cares_dns" json:"caresDns,omitempty"`
						DnsCacheCircuitBreaker *struct {
							MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
						} `tfsdk:"dns_cache_circuit_breaker" json:"dnsCacheCircuitBreaker,omitempty"`
						DnsFailureRefreshRate *struct {
							BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
							MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
						} `tfsdk:"dns_failure_refresh_rate" json:"dnsFailureRefreshRate,omitempty"`
						DnsLookupFamily     *string `tfsdk:"dns_lookup_family" json:"dnsLookupFamily,omitempty"`
						DnsQueryTimeout     *string `tfsdk:"dns_query_timeout" json:"dnsQueryTimeout,omitempty"`
						DnsRefreshRate      *string `tfsdk:"dns_refresh_rate" json:"dnsRefreshRate,omitempty"`
						HostTtl             *string `tfsdk:"host_ttl" json:"hostTtl,omitempty"`
						MaxHosts            *int64  `tfsdk:"max_hosts" json:"maxHosts,omitempty"`
						PreresolveHostnames *[]struct {
							Address      *string `tfsdk:"address" json:"address,omitempty"`
							Ipv4Compat   *bool   `tfsdk:"ipv4_compat" json:"ipv4Compat,omitempty"`
							NamedPort    *string `tfsdk:"named_port" json:"namedPort,omitempty"`
							PortValue    *int64  `tfsdk:"port_value" json:"portValue,omitempty"`
							Protocol     *string `tfsdk:"protocol" json:"protocol,omitempty"`
							ResolverName *string `tfsdk:"resolver_name" json:"resolverName,omitempty"`
						} `tfsdk:"preresolve_hostnames" json:"preresolveHostnames,omitempty"`
					} `tfsdk:"dns_cache_config" json:"dnsCacheConfig,omitempty"`
					SaveUpstreamAddress *bool `tfsdk:"save_upstream_address" json:"saveUpstreamAddress,omitempty"`
					SslConfig           *struct {
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
				} `tfsdk:"dynamic_forward_proxy" json:"dynamicForwardProxy,omitempty"`
				ExtProc *struct {
					AllowModeOverride      *bool              `tfsdk:"allow_mode_override" json:"allowModeOverride,omitempty"`
					AsyncMode              *bool              `tfsdk:"async_mode" json:"asyncMode,omitempty"`
					DisableClearRouteCache *bool              `tfsdk:"disable_clear_route_cache" json:"disableClearRouteCache,omitempty"`
					FailureModeAllow       *bool              `tfsdk:"failure_mode_allow" json:"failureModeAllow,omitempty"`
					FilterMetadata         *map[string]string `tfsdk:"filter_metadata" json:"filterMetadata,omitempty"`
					FilterStage            *struct {
						Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
						Stage     *string `tfsdk:"stage" json:"stage,omitempty"`
					} `tfsdk:"filter_stage" json:"filterStage,omitempty"`
					ForwardRules *struct {
						AllowedHeaders *struct {
							Patterns *[]struct {
								Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
								IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
								Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
								SafeRegex  *struct {
									GoogleRe2 *struct {
										MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
									} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
									Regex *string `tfsdk:"regex" json:"regex,omitempty"`
								} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
								Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
							} `tfsdk:"patterns" json:"patterns,omitempty"`
						} `tfsdk:"allowed_headers" json:"allowedHeaders,omitempty"`
						DisallowedHeaders *struct {
							Patterns *[]struct {
								Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
								IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
								Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
								SafeRegex  *struct {
									GoogleRe2 *struct {
										MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
									} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
									Regex *string `tfsdk:"regex" json:"regex,omitempty"`
								} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
								Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
							} `tfsdk:"patterns" json:"patterns,omitempty"`
						} `tfsdk:"disallowed_headers" json:"disallowedHeaders,omitempty"`
					} `tfsdk:"forward_rules" json:"forwardRules,omitempty"`
					GrpcService *struct {
						Authority        *string `tfsdk:"authority" json:"authority,omitempty"`
						ExtProcServerRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"ext_proc_server_ref" json:"extProcServerRef,omitempty"`
						InitialMetadata *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"initial_metadata" json:"initialMetadata,omitempty"`
						RetryPolicy *struct {
							NumRetries   *int64 `tfsdk:"num_retries" json:"numRetries,omitempty"`
							RetryBackOff *struct {
								BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
								MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
							} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
						} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
						Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					} `tfsdk:"grpc_service" json:"grpcService,omitempty"`
					MaxMessageTimeout         *string   `tfsdk:"max_message_timeout" json:"maxMessageTimeout,omitempty"`
					MessageTimeout            *string   `tfsdk:"message_timeout" json:"messageTimeout,omitempty"`
					MetadataContextNamespaces *[]string `tfsdk:"metadata_context_namespaces" json:"metadataContextNamespaces,omitempty"`
					MutationRules             *struct {
						AllowAllRouting *bool `tfsdk:"allow_all_routing" json:"allowAllRouting,omitempty"`
						AllowEnvoy      *bool `tfsdk:"allow_envoy" json:"allowEnvoy,omitempty"`
						AllowExpression *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"allow_expression" json:"allowExpression,omitempty"`
						DisallowAll        *bool `tfsdk:"disallow_all" json:"disallowAll,omitempty"`
						DisallowExpression *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"disallow_expression" json:"disallowExpression,omitempty"`
						DisallowIsError *bool `tfsdk:"disallow_is_error" json:"disallowIsError,omitempty"`
						DisallowSystem  *bool `tfsdk:"disallow_system" json:"disallowSystem,omitempty"`
					} `tfsdk:"mutation_rules" json:"mutationRules,omitempty"`
					ProcessingMode *struct {
						RequestBodyMode     *string `tfsdk:"request_body_mode" json:"requestBodyMode,omitempty"`
						RequestHeaderMode   *string `tfsdk:"request_header_mode" json:"requestHeaderMode,omitempty"`
						RequestTrailerMode  *string `tfsdk:"request_trailer_mode" json:"requestTrailerMode,omitempty"`
						ResponseBodyMode    *string `tfsdk:"response_body_mode" json:"responseBodyMode,omitempty"`
						ResponseHeaderMode  *string `tfsdk:"response_header_mode" json:"responseHeaderMode,omitempty"`
						ResponseTrailerMode *string `tfsdk:"response_trailer_mode" json:"responseTrailerMode,omitempty"`
					} `tfsdk:"processing_mode" json:"processingMode,omitempty"`
					RequestAttributes              *[]string `tfsdk:"request_attributes" json:"requestAttributes,omitempty"`
					ResponseAttributes             *[]string `tfsdk:"response_attributes" json:"responseAttributes,omitempty"`
					StatPrefix                     *string   `tfsdk:"stat_prefix" json:"statPrefix,omitempty"`
					TypedMetadataContextNamespaces *[]string `tfsdk:"typed_metadata_context_namespaces" json:"typedMetadataContextNamespaces,omitempty"`
				} `tfsdk:"ext_proc" json:"extProc,omitempty"`
				Extauth *struct {
					ClearRouteCache   *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
					ExtauthzServerRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"extauthz_server_ref" json:"extauthzServerRef,omitempty"`
					FailureModeAllow *bool `tfsdk:"failure_mode_allow" json:"failureModeAllow,omitempty"`
					GrpcService      *struct {
						Authority *string `tfsdk:"authority" json:"authority,omitempty"`
					} `tfsdk:"grpc_service" json:"grpcService,omitempty"`
					HttpService *struct {
						PathPrefix *string `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
						Request    *struct {
							AllowedHeaders      *[]string          `tfsdk:"allowed_headers" json:"allowedHeaders,omitempty"`
							AllowedHeadersRegex *[]string          `tfsdk:"allowed_headers_regex" json:"allowedHeadersRegex,omitempty"`
							HeadersToAdd        *map[string]string `tfsdk:"headers_to_add" json:"headersToAdd,omitempty"`
						} `tfsdk:"request" json:"request,omitempty"`
						Response *struct {
							AllowedClientHeaders           *[]string `tfsdk:"allowed_client_headers" json:"allowedClientHeaders,omitempty"`
							AllowedUpstreamHeaders         *[]string `tfsdk:"allowed_upstream_headers" json:"allowedUpstreamHeaders,omitempty"`
							AllowedUpstreamHeadersToAppend *[]string `tfsdk:"allowed_upstream_headers_to_append" json:"allowedUpstreamHeadersToAppend,omitempty"`
						} `tfsdk:"response" json:"response,omitempty"`
					} `tfsdk:"http_service" json:"httpService,omitempty"`
					RequestBody *struct {
						AllowPartialMessage *bool  `tfsdk:"allow_partial_message" json:"allowPartialMessage,omitempty"`
						MaxRequestBytes     *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
						PackAsBytes         *bool  `tfsdk:"pack_as_bytes" json:"packAsBytes,omitempty"`
					} `tfsdk:"request_body" json:"requestBody,omitempty"`
					RequestTimeout      *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
					StatPrefix          *string `tfsdk:"stat_prefix" json:"statPrefix,omitempty"`
					StatusOnError       *int64  `tfsdk:"status_on_error" json:"statusOnError,omitempty"`
					TransportApiVersion *string `tfsdk:"transport_api_version" json:"transportApiVersion,omitempty"`
					UserIdHeader        *string `tfsdk:"user_id_header" json:"userIdHeader,omitempty"`
				} `tfsdk:"extauth" json:"extauth,omitempty"`
				Extensions *struct {
					Configs *map[string]string `tfsdk:"configs" json:"configs,omitempty"`
				} `tfsdk:"extensions" json:"extensions,omitempty"`
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
				GrpcWeb *struct {
					Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
				} `tfsdk:"grpc_web" json:"grpcWeb,omitempty"`
				Gzip *struct {
					CompressionLevel           *string   `tfsdk:"compression_level" json:"compressionLevel,omitempty"`
					CompressionStrategy        *string   `tfsdk:"compression_strategy" json:"compressionStrategy,omitempty"`
					ContentLength              *int64    `tfsdk:"content_length" json:"contentLength,omitempty"`
					ContentType                *[]string `tfsdk:"content_type" json:"contentType,omitempty"`
					DisableOnEtagHeader        *bool     `tfsdk:"disable_on_etag_header" json:"disableOnEtagHeader,omitempty"`
					MemoryLevel                *int64    `tfsdk:"memory_level" json:"memoryLevel,omitempty"`
					RemoveAcceptEncodingHeader *bool     `tfsdk:"remove_accept_encoding_header" json:"removeAcceptEncodingHeader,omitempty"`
					WindowBits                 *int64    `tfsdk:"window_bits" json:"windowBits,omitempty"`
				} `tfsdk:"gzip" json:"gzip,omitempty"`
				HealthCheck *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"health_check" json:"healthCheck,omitempty"`
				HttpConnectionManagerSettings *struct {
					AcceptHttp10                 *bool   `tfsdk:"accept_http10" json:"acceptHttp10,omitempty"`
					AllowChunkedLength           *bool   `tfsdk:"allow_chunked_length" json:"allowChunkedLength,omitempty"`
					AppendXForwardedPort         *bool   `tfsdk:"append_x_forwarded_port" json:"appendXForwardedPort,omitempty"`
					CodecType                    *string `tfsdk:"codec_type" json:"codecType,omitempty"`
					DefaultHostForHttp10         *string `tfsdk:"default_host_for_http10" json:"defaultHostForHttp10,omitempty"`
					DelayedCloseTimeout          *string `tfsdk:"delayed_close_timeout" json:"delayedCloseTimeout,omitempty"`
					DrainTimeout                 *string `tfsdk:"drain_timeout" json:"drainTimeout,omitempty"`
					EnableTrailers               *bool   `tfsdk:"enable_trailers" json:"enableTrailers,omitempty"`
					ForwardClientCertDetails     *string `tfsdk:"forward_client_cert_details" json:"forwardClientCertDetails,omitempty"`
					GenerateRequestId            *bool   `tfsdk:"generate_request_id" json:"generateRequestId,omitempty"`
					HeadersWithUnderscoresAction *string `tfsdk:"headers_with_underscores_action" json:"headersWithUnderscoresAction,omitempty"`
					Http2ProtocolOptions         *struct {
						InitialConnectionWindowSize             *int64 `tfsdk:"initial_connection_window_size" json:"initialConnectionWindowSize,omitempty"`
						InitialStreamWindowSize                 *int64 `tfsdk:"initial_stream_window_size" json:"initialStreamWindowSize,omitempty"`
						MaxConcurrentStreams                    *int64 `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
						OverrideStreamErrorOnInvalidHttpMessage *bool  `tfsdk:"override_stream_error_on_invalid_http_message" json:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`
					} `tfsdk:"http2_protocol_options" json:"http2ProtocolOptions,omitempty"`
					IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
					InternalAddressConfig *struct {
						CidrRanges *[]struct {
							AddressPrefix *string `tfsdk:"address_prefix" json:"addressPrefix,omitempty"`
							PrefixLen     *int64  `tfsdk:"prefix_len" json:"prefixLen,omitempty"`
						} `tfsdk:"cidr_ranges" json:"cidrRanges,omitempty"`
						UnixSockets *bool `tfsdk:"unix_sockets" json:"unixSockets,omitempty"`
					} `tfsdk:"internal_address_config" json:"internalAddressConfig,omitempty"`
					MaxConnectionDuration        *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
					MaxHeadersCount              *int64  `tfsdk:"max_headers_count" json:"maxHeadersCount,omitempty"`
					MaxRequestHeadersKb          *int64  `tfsdk:"max_request_headers_kb" json:"maxRequestHeadersKb,omitempty"`
					MaxRequestsPerConnection     *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
					MaxStreamDuration            *string `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
					MergeSlashes                 *bool   `tfsdk:"merge_slashes" json:"mergeSlashes,omitempty"`
					NormalizePath                *bool   `tfsdk:"normalize_path" json:"normalizePath,omitempty"`
					PathWithEscapedSlashesAction *string `tfsdk:"path_with_escaped_slashes_action" json:"pathWithEscapedSlashesAction,omitempty"`
					PreserveCaseHeaderKeyFormat  *bool   `tfsdk:"preserve_case_header_key_format" json:"preserveCaseHeaderKeyFormat,omitempty"`
					PreserveExternalRequestId    *bool   `tfsdk:"preserve_external_request_id" json:"preserveExternalRequestId,omitempty"`
					ProperCaseHeaderKeyFormat    *bool   `tfsdk:"proper_case_header_key_format" json:"properCaseHeaderKeyFormat,omitempty"`
					Proxy100Continue             *bool   `tfsdk:"proxy100_continue" json:"proxy100Continue,omitempty"`
					RequestHeadersTimeout        *string `tfsdk:"request_headers_timeout" json:"requestHeadersTimeout,omitempty"`
					RequestTimeout               *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
					ServerHeaderTransformation   *string `tfsdk:"server_header_transformation" json:"serverHeaderTransformation,omitempty"`
					ServerName                   *string `tfsdk:"server_name" json:"serverName,omitempty"`
					SetCurrentClientCertDetails  *struct {
						Cert    *bool `tfsdk:"cert" json:"cert,omitempty"`
						Chain   *bool `tfsdk:"chain" json:"chain,omitempty"`
						Dns     *bool `tfsdk:"dns" json:"dns,omitempty"`
						Subject *bool `tfsdk:"subject" json:"subject,omitempty"`
						Uri     *bool `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"set_current_client_cert_details" json:"setCurrentClientCertDetails,omitempty"`
					SkipXffAppend     *bool   `tfsdk:"skip_xff_append" json:"skipXffAppend,omitempty"`
					StreamIdleTimeout *string `tfsdk:"stream_idle_timeout" json:"streamIdleTimeout,omitempty"`
					StripAnyHostPort  *bool   `tfsdk:"strip_any_host_port" json:"stripAnyHostPort,omitempty"`
					Tracing           *struct {
						DatadogConfig *struct {
							ClusterName          *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
							CollectorUpstreamRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"collector_upstream_ref" json:"collectorUpstreamRef,omitempty"`
							ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
						} `tfsdk:"datadog_config" json:"datadogConfig,omitempty"`
						EnvironmentVariablesForTags *[]struct {
							DefaultValue *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							Name         *string `tfsdk:"name" json:"name,omitempty"`
							Tag          *string `tfsdk:"tag" json:"tag,omitempty"`
						} `tfsdk:"environment_variables_for_tags" json:"environmentVariablesForTags,omitempty"`
						LiteralsForTags *[]struct {
							Tag   *string `tfsdk:"tag" json:"tag,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"literals_for_tags" json:"literalsForTags,omitempty"`
						OpenCensusConfig *struct {
							GrpcAddress *struct {
								StatPrefix *string `tfsdk:"stat_prefix" json:"statPrefix,omitempty"`
								TargetUri  *string `tfsdk:"target_uri" json:"targetUri,omitempty"`
							} `tfsdk:"grpc_address" json:"grpcAddress,omitempty"`
							HttpAddress            *string   `tfsdk:"http_address" json:"httpAddress,omitempty"`
							IncomingTraceContext   *[]string `tfsdk:"incoming_trace_context" json:"incomingTraceContext,omitempty"`
							OcagentExporterEnabled *bool     `tfsdk:"ocagent_exporter_enabled" json:"ocagentExporterEnabled,omitempty"`
							OutgoingTraceContext   *[]string `tfsdk:"outgoing_trace_context" json:"outgoingTraceContext,omitempty"`
							TraceConfig            *struct {
								ConstantSampler *struct {
									Decision *string `tfsdk:"decision" json:"decision,omitempty"`
								} `tfsdk:"constant_sampler" json:"constantSampler,omitempty"`
								MaxNumberOfAnnotations   *int64 `tfsdk:"max_number_of_annotations" json:"maxNumberOfAnnotations,omitempty"`
								MaxNumberOfAttributes    *int64 `tfsdk:"max_number_of_attributes" json:"maxNumberOfAttributes,omitempty"`
								MaxNumberOfLinks         *int64 `tfsdk:"max_number_of_links" json:"maxNumberOfLinks,omitempty"`
								MaxNumberOfMessageEvents *int64 `tfsdk:"max_number_of_message_events" json:"maxNumberOfMessageEvents,omitempty"`
								ProbabilitySampler       *struct {
									SamplingProbability *float64 `tfsdk:"sampling_probability" json:"samplingProbability,omitempty"`
								} `tfsdk:"probability_sampler" json:"probabilitySampler,omitempty"`
								RateLimitingSampler *struct {
									Qps *int64 `tfsdk:"qps" json:"qps,omitempty"`
								} `tfsdk:"rate_limiting_sampler" json:"rateLimitingSampler,omitempty"`
							} `tfsdk:"trace_config" json:"traceConfig,omitempty"`
						} `tfsdk:"open_census_config" json:"openCensusConfig,omitempty"`
						OpenTelemetryConfig *struct {
							ClusterName          *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
							CollectorUpstreamRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"collector_upstream_ref" json:"collectorUpstreamRef,omitempty"`
						} `tfsdk:"open_telemetry_config" json:"openTelemetryConfig,omitempty"`
						RequestHeadersForTags *[]string `tfsdk:"request_headers_for_tags" json:"requestHeadersForTags,omitempty"`
						TracePercentages      *struct {
							ClientSamplePercentage  *float64 `tfsdk:"client_sample_percentage" json:"clientSamplePercentage,omitempty"`
							OverallSamplePercentage *float64 `tfsdk:"overall_sample_percentage" json:"overallSamplePercentage,omitempty"`
							RandomSamplePercentage  *float64 `tfsdk:"random_sample_percentage" json:"randomSamplePercentage,omitempty"`
						} `tfsdk:"trace_percentages" json:"tracePercentages,omitempty"`
						Verbose      *bool `tfsdk:"verbose" json:"verbose,omitempty"`
						ZipkinConfig *struct {
							ClusterName              *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
							CollectorEndpoint        *string `tfsdk:"collector_endpoint" json:"collectorEndpoint,omitempty"`
							CollectorEndpointVersion *string `tfsdk:"collector_endpoint_version" json:"collectorEndpointVersion,omitempty"`
							CollectorUpstreamRef     *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"collector_upstream_ref" json:"collectorUpstreamRef,omitempty"`
							SharedSpanContext *bool `tfsdk:"shared_span_context" json:"sharedSpanContext,omitempty"`
							TraceId128bit     *bool `tfsdk:"trace_id128bit" json:"traceId128bit,omitempty"`
						} `tfsdk:"zipkin_config" json:"zipkinConfig,omitempty"`
					} `tfsdk:"tracing" json:"tracing,omitempty"`
					Upgrades *[]struct {
						Connect *struct {
							Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						} `tfsdk:"connect" json:"connect,omitempty"`
						Websocket *struct {
							Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						} `tfsdk:"websocket" json:"websocket,omitempty"`
					} `tfsdk:"upgrades" json:"upgrades,omitempty"`
					UseRemoteAddress    *bool `tfsdk:"use_remote_address" json:"useRemoteAddress,omitempty"`
					UuidRequestIdConfig *struct {
						PackTraceReason              *bool `tfsdk:"pack_trace_reason" json:"packTraceReason,omitempty"`
						UseRequestIdForTraceSampling *bool `tfsdk:"use_request_id_for_trace_sampling" json:"useRequestIdForTraceSampling,omitempty"`
					} `tfsdk:"uuid_request_id_config" json:"uuidRequestIdConfig,omitempty"`
					Via               *string `tfsdk:"via" json:"via,omitempty"`
					XffNumTrustedHops *int64  `tfsdk:"xff_num_trusted_hops" json:"xffNumTrustedHops,omitempty"`
				} `tfsdk:"http_connection_manager_settings" json:"httpConnectionManagerSettings,omitempty"`
				HttpLocalRatelimit *struct {
					DefaultLimit *struct {
						FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
						MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
						TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
					} `tfsdk:"default_limit" json:"defaultLimit,omitempty"`
					EnableXRatelimitHeaders               *bool `tfsdk:"enable_x_ratelimit_headers" json:"enableXRatelimitHeaders,omitempty"`
					LocalRateLimitPerDownstreamConnection *bool `tfsdk:"local_rate_limit_per_downstream_connection" json:"localRateLimitPerDownstreamConnection,omitempty"`
				} `tfsdk:"http_local_ratelimit" json:"httpLocalRatelimit,omitempty"`
				LeftmostXffAddress    *bool `tfsdk:"leftmost_xff_address" json:"leftmostXffAddress,omitempty"`
				NetworkLocalRatelimit *struct {
					FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
					MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
					TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
				} `tfsdk:"network_local_ratelimit" json:"networkLocalRatelimit,omitempty"`
				ProxyLatency *struct {
					ChargeClusterStat        *bool   `tfsdk:"charge_cluster_stat" json:"chargeClusterStat,omitempty"`
					ChargeListenerStat       *bool   `tfsdk:"charge_listener_stat" json:"chargeListenerStat,omitempty"`
					EmitDynamicMetadata      *bool   `tfsdk:"emit_dynamic_metadata" json:"emitDynamicMetadata,omitempty"`
					MeasureRequestInternally *bool   `tfsdk:"measure_request_internally" json:"measureRequestInternally,omitempty"`
					Request                  *string `tfsdk:"request" json:"request,omitempty"`
					Response                 *string `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"proxy_latency" json:"proxyLatency,omitempty"`
				RatelimitServer *struct {
					DenyOnFail              *bool `tfsdk:"deny_on_fail" json:"denyOnFail,omitempty"`
					EnableXRatelimitHeaders *bool `tfsdk:"enable_x_ratelimit_headers" json:"enableXRatelimitHeaders,omitempty"`
					RateLimitBeforeAuth     *bool `tfsdk:"rate_limit_before_auth" json:"rateLimitBeforeAuth,omitempty"`
					RatelimitServerRef      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ratelimit_server_ref" json:"ratelimitServerRef,omitempty"`
					RequestTimeout *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
				} `tfsdk:"ratelimit_server" json:"ratelimitServer,omitempty"`
				Router *struct {
					SuppressEnvoyHeaders *bool `tfsdk:"suppress_envoy_headers" json:"suppressEnvoyHeaders,omitempty"`
				} `tfsdk:"router" json:"router,omitempty"`
				SanitizeClusterHeader *bool `tfsdk:"sanitize_cluster_header" json:"sanitizeClusterHeader,omitempty"`
				Waf                   *struct {
					AuditLogging *struct {
						Action   *string `tfsdk:"action" json:"action,omitempty"`
						Location *string `tfsdk:"location" json:"location,omitempty"`
					} `tfsdk:"audit_logging" json:"auditLogging,omitempty"`
					ConfigMapRuleSets *[]struct {
						ConfigMapRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						DataMapKeys *[]string `tfsdk:"data_map_keys" json:"dataMapKeys,omitempty"`
					} `tfsdk:"config_map_rule_sets" json:"configMapRuleSets,omitempty"`
					CoreRuleSet *struct {
						CustomSettingsFile   *string `tfsdk:"custom_settings_file" json:"customSettingsFile,omitempty"`
						CustomSettingsString *string `tfsdk:"custom_settings_string" json:"customSettingsString,omitempty"`
					} `tfsdk:"core_rule_set" json:"coreRuleSet,omitempty"`
					CustomInterventionMessage *string `tfsdk:"custom_intervention_message" json:"customInterventionMessage,omitempty"`
					Disabled                  *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					RequestHeadersOnly        *bool   `tfsdk:"request_headers_only" json:"requestHeadersOnly,omitempty"`
					ResponseHeadersOnly       *bool   `tfsdk:"response_headers_only" json:"responseHeadersOnly,omitempty"`
					RuleSets                  *[]struct {
						Directory *string   `tfsdk:"directory" json:"directory,omitempty"`
						Files     *[]string `tfsdk:"files" json:"files,omitempty"`
						RuleStr   *string   `tfsdk:"rule_str" json:"ruleStr,omitempty"`
					} `tfsdk:"rule_sets" json:"ruleSets,omitempty"`
				} `tfsdk:"waf" json:"waf,omitempty"`
				Wasm *struct {
					Filters *[]struct {
						Config      *map[string]string `tfsdk:"config" json:"config,omitempty"`
						FailOpen    *bool              `tfsdk:"fail_open" json:"failOpen,omitempty"`
						FilePath    *string            `tfsdk:"file_path" json:"filePath,omitempty"`
						FilterStage *struct {
							Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
							Stage     *string `tfsdk:"stage" json:"stage,omitempty"`
						} `tfsdk:"filter_stage" json:"filterStage,omitempty"`
						Image  *string `tfsdk:"image" json:"image,omitempty"`
						Name   *string `tfsdk:"name" json:"name,omitempty"`
						RootId *string `tfsdk:"root_id" json:"rootId,omitempty"`
						VmType *string `tfsdk:"vm_type" json:"vmType,omitempty"`
					} `tfsdk:"filters" json:"filters,omitempty"`
				} `tfsdk:"wasm" json:"wasm,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			VirtualServiceExpressions *struct {
				Expressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"expressions" json:"expressions,omitempty"`
			} `tfsdk:"virtual_service_expressions" json:"virtualServiceExpressions,omitempty"`
			VirtualServiceNamespaces *[]string          `tfsdk:"virtual_service_namespaces" json:"virtualServiceNamespaces,omitempty"`
			VirtualServiceSelector   *map[string]string `tfsdk:"virtual_service_selector" json:"virtualServiceSelector,omitempty"`
			VirtualServices          *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"virtual_services" json:"virtualServices,omitempty"`
		} `tfsdk:"http_gateway" json:"httpGateway,omitempty"`
		Matcher *struct {
			SourcePrefixRanges *[]struct {
				AddressPrefix *string `tfsdk:"address_prefix" json:"addressPrefix,omitempty"`
				PrefixLen     *int64  `tfsdk:"prefix_len" json:"prefixLen,omitempty"`
			} `tfsdk:"source_prefix_ranges" json:"sourcePrefixRanges,omitempty"`
			SslConfig *struct {
				AlpnProtocols               *[]string `tfsdk:"alpn_protocols" json:"alpnProtocols,omitempty"`
				DisableTlsSessionResumption *bool     `tfsdk:"disable_tls_session_resumption" json:"disableTlsSessionResumption,omitempty"`
				OcspStaplePolicy            *string   `tfsdk:"ocsp_staple_policy" json:"ocspStaplePolicy,omitempty"`
				OneWayTls                   *bool     `tfsdk:"one_way_tls" json:"oneWayTls,omitempty"`
				Parameters                  *struct {
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
				SniDomains *[]string `tfsdk:"sni_domains" json:"sniDomains,omitempty"`
				SslFiles   *struct {
					OcspStaple *string `tfsdk:"ocsp_staple" json:"ocspStaple,omitempty"`
					RootCa     *string `tfsdk:"root_ca" json:"rootCa,omitempty"`
					TlsCert    *string `tfsdk:"tls_cert" json:"tlsCert,omitempty"`
					TlsKey     *string `tfsdk:"tls_key" json:"tlsKey,omitempty"`
				} `tfsdk:"ssl_files" json:"sslFiles,omitempty"`
				TransportSocketConnectTimeout *string   `tfsdk:"transport_socket_connect_timeout" json:"transportSocketConnectTimeout,omitempty"`
				VerifySubjectAltName          *[]string `tfsdk:"verify_subject_alt_name" json:"verifySubjectAltName,omitempty"`
			} `tfsdk:"ssl_config" json:"sslConfig,omitempty"`
		} `tfsdk:"matcher" json:"matcher,omitempty"`
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewaySoloIoMatchableHttpGatewayV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_solo_io_matchable_http_gateway_v1"
}

func (r *GatewaySoloIoMatchableHttpGatewayV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"http_gateway": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"buffer": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"max_request_bytes": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"caching": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"allowed_vary_headers": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"exact": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"ignore_case": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"safe_regex": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"google_re2": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"max_program_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"suffix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"caching_service_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"max_payload_size": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"connection_limit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"delay_before_close": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_active_connections": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"csrf": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"additional_origins": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"exact": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"ignore_case": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"safe_regex": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"google_re2": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"max_program_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"suffix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"filter_enabled": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default_value": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"denominator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"numerator": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"runtime_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"shadow_enabled": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default_value": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"denominator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"numerator": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"runtime_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"disable_ext_proc": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"dlp": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"dlp_rules": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"action_type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"custom_action": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"mask_char": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"percent": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"value": schema.Float64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"regex": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"regex_actions": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"key_value_action": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key_to_mask": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"mask_char": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"percent": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"value": schema.Float64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"shadow": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"matcher": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"case_sensitive": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"connect_matcher": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"exact": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"invert_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"methods": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"query_parameters": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"enabled_for": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"dynamic_forward_proxy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"dns_cache_config": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"apple_dns": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cares_dns": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"dns_resolver_options": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"no_default_search_domain": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"use_tcp_for_dns_lookups": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"resolvers": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"pipe": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"mode": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"socket_address": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"address": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ipv4_compat": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"named_port": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"port_value": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"resolver_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"dns_cache_circuit_breaker": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"max_pending_requests": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"dns_failure_refresh_rate": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"base_interval": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"max_interval": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"dns_lookup_family": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"dns_query_timeout": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"dns_refresh_rate": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_ttl": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_hosts": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"preresolve_hostnames": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"address": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"ipv4_compat": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"named_port": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"port_value": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"protocol": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"resolver_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"save_upstream_address": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ssl_config": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"allow_renegotiation": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"alpn_protocols": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																Optional:            false,
																Computed:            true,
															},

															"ecdh_curves": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"maximum_protocol_version": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"minimum_protocol_version": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																				Optional:            false,
																				Computed:            true,
																			},

																			"token_file_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"certificates_secret_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"cluster_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"target_uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"validation_context_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"sni": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"ssl_files": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"ocsp_staple": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"root_ca": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"tls_cert": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"tls_key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"verify_subject_alt_name": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ext_proc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"allow_mode_override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"async_mode": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"disable_clear_route_cache": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"failure_mode_allow": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"filter_metadata": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"filter_stage": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"predicate": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"stage": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"forward_rules": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"allowed_headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"patterns": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"exact": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"ignore_case": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"prefix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"safe_regex": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"google_re2": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"max_program_size": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"suffix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"disallowed_headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"patterns": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"exact": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"ignore_case": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"prefix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"safe_regex": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"google_re2": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"max_program_size": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"suffix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"grpc_service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"authority": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"ext_proc_server_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"initial_metadata": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"retry_policy": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"num_retries": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"retry_back_off": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"base_interval": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"max_interval": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"timeout": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"max_message_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"message_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"metadata_context_namespaces": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"mutation_rules": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"allow_all_routing": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"allow_envoy": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"allow_expression": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"google_re2": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"max_program_size": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"disallow_all": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"disallow_expression": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"google_re2": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"max_program_size": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"disallow_is_error": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"disallow_system": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"processing_mode": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"request_body_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"request_header_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"request_trailer_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"response_body_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"response_header_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"response_trailer_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_attributes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_attributes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"stat_prefix": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"typed_metadata_context_namespaces": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"extauth": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"clear_route_cache": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"extauthz_server_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"failure_mode_allow": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"grpc_service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"authority": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"http_service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"path_prefix": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"request": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"allowed_headers": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"allowed_headers_regex": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"headers_to_add": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"response": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"allowed_client_headers": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"allowed_upstream_headers": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"allowed_upstream_headers_to_append": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_body": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"allow_partial_message": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_request_bytes": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"pack_as_bytes": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"stat_prefix": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"status_on_error": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"transport_api_version": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"user_id_header": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"extensions": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"configs": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"grpc_json_transcoder": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"auto_mapping": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"convert_grpc_status": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ignore_unknown_query_parameters": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ignored_query_parameters": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"match_incoming_request_route": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"print_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_whitespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"always_print_enums_as_ints": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"always_print_primitive_fields": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"preserve_proto_field_names": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"proto_descriptor": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"proto_descriptor_bin": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"services": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"grpc_web": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"disable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"gzip": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"compression_level": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"compression_strategy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_length": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"disable_on_etag_header": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"memory_level": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"remove_accept_encoding_header": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"window_bits": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"health_check": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"http_connection_manager_settings": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"accept_http10": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"allow_chunked_length": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"append_x_forwarded_port": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"codec_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"default_host_for_http10": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"delayed_close_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"drain_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"enable_trailers": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"forward_client_cert_details": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"generate_request_id": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"headers_with_underscores_action": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"http2_protocol_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"initial_connection_window_size": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"initial_stream_window_size": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_concurrent_streams": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"override_stream_error_on_invalid_http_message": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"idle_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"internal_address_config": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cidr_ranges": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"address_prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"prefix_len": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"unix_sockets": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"max_connection_duration": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_headers_count": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_request_headers_kb": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_requests_per_connection": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_stream_duration": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"merge_slashes": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"normalize_path": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"path_with_escaped_slashes_action": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"preserve_case_header_key_format": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"preserve_external_request_id": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"proper_case_header_key_format": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"proxy100_continue": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request_headers_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"server_header_transformation": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"server_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"set_current_client_cert_details": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cert": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"chain": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"dns": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"subject": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"uri": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"skip_xff_append": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"stream_idle_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"strip_any_host_port": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tracing": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"datadog_config": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"cluster_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"collector_upstream_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"service_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"environment_variables_for_tags": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"default_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tag": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"literals_for_tags": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"tag": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"open_census_config": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"grpc_address": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"stat_prefix": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"target_uri": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"http_address": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"incoming_trace_context": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"ocagent_exporter_enabled": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"outgoing_trace_context": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"trace_config": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"constant_sampler": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"decision": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"max_number_of_annotations": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"max_number_of_attributes": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"max_number_of_links": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"max_number_of_message_events": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"probability_sampler": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"sampling_probability": schema.Float64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"rate_limiting_sampler": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"qps": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"open_telemetry_config": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"cluster_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"collector_upstream_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"request_headers_for_tags": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"trace_percentages": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"client_sample_percentage": schema.Float64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"overall_sample_percentage": schema.Float64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"random_sample_percentage": schema.Float64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"verbose": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"zipkin_config": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"cluster_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"collector_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"collector_endpoint_version": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"collector_upstream_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"shared_span_context": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"trace_id128bit": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"upgrades": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"connect": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enabled": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"websocket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enabled": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"use_remote_address": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"uuid_request_id_config": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"pack_trace_reason": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"use_request_id_for_trace_sampling": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"via": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"xff_num_trusted_hops": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"http_local_ratelimit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"default_limit": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"fill_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_tokens": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"tokens_per_fill": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"enable_x_ratelimit_headers": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"local_rate_limit_per_downstream_connection": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"leftmost_xff_address": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"network_local_ratelimit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"fill_interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_tokens": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tokens_per_fill": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"proxy_latency": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"charge_cluster_stat": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"charge_listener_stat": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"emit_dynamic_metadata": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"measure_request_internally": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ratelimit_server": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"deny_on_fail": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"enable_x_ratelimit_headers": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"rate_limit_before_auth": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ratelimit_server_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"request_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"router": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"suppress_envoy_headers": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"sanitize_cluster_header": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"waf": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"audit_logging": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"location": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"config_map_rule_sets": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"config_map_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"data_map_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"core_rule_set": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"custom_settings_file": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"custom_settings_string": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"custom_intervention_message": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"disabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request_headers_only": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_headers_only": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"rule_sets": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"directory": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"files": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"rule_str": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"wasm": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"filters": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"config": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"fail_open": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"file_path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"filter_stage": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"predicate": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"stage": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"image": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"root_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"vm_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"virtual_service_expressions": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"expressions": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"virtual_service_namespaces": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"virtual_service_selector": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"virtual_services": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"matcher": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"source_prefix_ranges": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"address_prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"prefix_len": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ssl_config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"alpn_protocols": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"disable_tls_session_resumption": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ocsp_staple_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"one_way_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"ecdh_curves": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"maximum_protocol_version": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"minimum_protocol_version": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"token_file_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"certificates_secret_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"cluster_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"target_uri": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"validation_context_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret_ref": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"sni_domains": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ssl_files": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"ocsp_staple": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"root_ca": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tls_cert": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tls_key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"transport_socket_connect_timeout": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"verify_subject_alt_name": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *GatewaySoloIoMatchableHttpGatewayV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *GatewaySoloIoMatchableHttpGatewayV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_gateway_solo_io_matchable_http_gateway_v1")

	var data GatewaySoloIoMatchableHttpGatewayV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gateway.solo.io", Version: "v1", Resource: "httpgateways"}).
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

	var readResponse GatewaySoloIoMatchableHttpGatewayV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("gateway.solo.io/v1")
	data.Kind = pointer.String("MatchableHttpGateway")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
