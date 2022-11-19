/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type GatewaySoloIoMatchableHttpGatewayV1Resource struct{}

var (
	_ resource.Resource = (*GatewaySoloIoMatchableHttpGatewayV1Resource)(nil)
)

type GatewaySoloIoMatchableHttpGatewayV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GatewaySoloIoMatchableHttpGatewayV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		HttpGateway *struct {
			Options *struct {
				Buffer *struct {
					MaxRequestBytes *int64 `tfsdk:"max_request_bytes" yaml:"maxRequestBytes,omitempty"`
				} `tfsdk:"buffer" yaml:"buffer,omitempty"`

				Caching *struct {
					AllowedVaryHeaders *[]struct {
						Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

						IgnoreCase *bool `tfsdk:"ignore_case" yaml:"ignoreCase,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						SafeRegex *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"safe_regex" yaml:"safeRegex,omitempty"`

						Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
					} `tfsdk:"allowed_vary_headers" yaml:"allowedVaryHeaders,omitempty"`

					CachingServiceRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"caching_service_ref" yaml:"cachingServiceRef,omitempty"`

					MaxPayloadSize *struct {
						Value utilities.IntOrString `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"max_payload_size" yaml:"maxPayloadSize,omitempty"`

					Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
				} `tfsdk:"caching" yaml:"caching,omitempty"`

				Csrf *struct {
					AdditionalOrigins *[]struct {
						Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

						IgnoreCase *bool `tfsdk:"ignore_case" yaml:"ignoreCase,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						SafeRegex *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"safe_regex" yaml:"safeRegex,omitempty"`

						Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
					} `tfsdk:"additional_origins" yaml:"additionalOrigins,omitempty"`

					FilterEnabled *struct {
						DefaultValue *struct {
							Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

							Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
						} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

						RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
					} `tfsdk:"filter_enabled" yaml:"filterEnabled,omitempty"`

					ShadowEnabled *struct {
						DefaultValue *struct {
							Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

							Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
						} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

						RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
					} `tfsdk:"shadow_enabled" yaml:"shadowEnabled,omitempty"`
				} `tfsdk:"csrf" yaml:"csrf,omitempty"`

				Dlp *struct {
					DlpRules *[]struct {
						Actions *[]struct {
							ActionType utilities.IntOrString `tfsdk:"action_type" yaml:"actionType,omitempty"`

							CustomAction *struct {
								MaskChar *string `tfsdk:"mask_char" yaml:"maskChar,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Percent *struct {
									Value utilities.DynamicNumber `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"percent" yaml:"percent,omitempty"`

								Regex *[]string `tfsdk:"regex" yaml:"regex,omitempty"`

								RegexActions *[]struct {
									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"regex_actions" yaml:"regexActions,omitempty"`
							} `tfsdk:"custom_action" yaml:"customAction,omitempty"`

							KeyValueAction *struct {
								KeyToMask *string `tfsdk:"key_to_mask" yaml:"keyToMask,omitempty"`

								MaskChar *string `tfsdk:"mask_char" yaml:"maskChar,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Percent *struct {
									Value utilities.DynamicNumber `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"percent" yaml:"percent,omitempty"`
							} `tfsdk:"key_value_action" yaml:"keyValueAction,omitempty"`

							Shadow *bool `tfsdk:"shadow" yaml:"shadow,omitempty"`
						} `tfsdk:"actions" yaml:"actions,omitempty"`

						Matcher *struct {
							CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

							Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

							Headers *[]struct {
								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

							QueryParameters *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"matcher" yaml:"matcher,omitempty"`
					} `tfsdk:"dlp_rules" yaml:"dlpRules,omitempty"`

					EnabledFor utilities.IntOrString `tfsdk:"enabled_for" yaml:"enabledFor,omitempty"`
				} `tfsdk:"dlp" yaml:"dlp,omitempty"`

				DynamicForwardProxy *struct {
					DnsCacheConfig *struct {
						AppleDns *map[string]string `tfsdk:"apple_dns" yaml:"appleDns,omitempty"`

						CaresDns *struct {
							DnsResolverOptions *struct {
								NoDefaultSearchDomain *bool `tfsdk:"no_default_search_domain" yaml:"noDefaultSearchDomain,omitempty"`

								UseTcpForDnsLookups *bool `tfsdk:"use_tcp_for_dns_lookups" yaml:"useTcpForDnsLookups,omitempty"`
							} `tfsdk:"dns_resolver_options" yaml:"dnsResolverOptions,omitempty"`

							Resolvers *[]struct {
								Pipe *struct {
									Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"pipe" yaml:"pipe,omitempty"`

								SocketAddress *struct {
									Address *string `tfsdk:"address" yaml:"address,omitempty"`

									Ipv4Compat *bool `tfsdk:"ipv4_compat" yaml:"ipv4Compat,omitempty"`

									NamedPort *string `tfsdk:"named_port" yaml:"namedPort,omitempty"`

									PortValue *int64 `tfsdk:"port_value" yaml:"portValue,omitempty"`

									Protocol utilities.IntOrString `tfsdk:"protocol" yaml:"protocol,omitempty"`

									ResolverName *string `tfsdk:"resolver_name" yaml:"resolverName,omitempty"`
								} `tfsdk:"socket_address" yaml:"socketAddress,omitempty"`
							} `tfsdk:"resolvers" yaml:"resolvers,omitempty"`
						} `tfsdk:"cares_dns" yaml:"caresDns,omitempty"`

						DnsCacheCircuitBreaker *struct {
							MaxPendingRequests *int64 `tfsdk:"max_pending_requests" yaml:"maxPendingRequests,omitempty"`
						} `tfsdk:"dns_cache_circuit_breaker" yaml:"dnsCacheCircuitBreaker,omitempty"`

						DnsFailureRefreshRate *struct {
							BaseInterval *string `tfsdk:"base_interval" yaml:"baseInterval,omitempty"`

							MaxInterval *string `tfsdk:"max_interval" yaml:"maxInterval,omitempty"`
						} `tfsdk:"dns_failure_refresh_rate" yaml:"dnsFailureRefreshRate,omitempty"`

						DnsLookupFamily utilities.IntOrString `tfsdk:"dns_lookup_family" yaml:"dnsLookupFamily,omitempty"`

						DnsQueryTimeout *string `tfsdk:"dns_query_timeout" yaml:"dnsQueryTimeout,omitempty"`

						DnsRefreshRate *string `tfsdk:"dns_refresh_rate" yaml:"dnsRefreshRate,omitempty"`

						HostTtl *string `tfsdk:"host_ttl" yaml:"hostTtl,omitempty"`

						MaxHosts *int64 `tfsdk:"max_hosts" yaml:"maxHosts,omitempty"`

						PreresolveHostnames *[]struct {
							Address *string `tfsdk:"address" yaml:"address,omitempty"`

							Ipv4Compat *bool `tfsdk:"ipv4_compat" yaml:"ipv4Compat,omitempty"`

							NamedPort *string `tfsdk:"named_port" yaml:"namedPort,omitempty"`

							PortValue *int64 `tfsdk:"port_value" yaml:"portValue,omitempty"`

							Protocol utilities.IntOrString `tfsdk:"protocol" yaml:"protocol,omitempty"`

							ResolverName *string `tfsdk:"resolver_name" yaml:"resolverName,omitempty"`
						} `tfsdk:"preresolve_hostnames" yaml:"preresolveHostnames,omitempty"`
					} `tfsdk:"dns_cache_config" yaml:"dnsCacheConfig,omitempty"`

					SaveUpstreamAddress *bool `tfsdk:"save_upstream_address" yaml:"saveUpstreamAddress,omitempty"`
				} `tfsdk:"dynamic_forward_proxy" yaml:"dynamicForwardProxy,omitempty"`

				Extauth *struct {
					ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

					ExtauthzServerRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"extauthz_server_ref" yaml:"extauthzServerRef,omitempty"`

					FailureModeAllow *bool `tfsdk:"failure_mode_allow" yaml:"failureModeAllow,omitempty"`

					GrpcService *struct {
						Authority *string `tfsdk:"authority" yaml:"authority,omitempty"`
					} `tfsdk:"grpc_service" yaml:"grpcService,omitempty"`

					HttpService *struct {
						PathPrefix *string `tfsdk:"path_prefix" yaml:"pathPrefix,omitempty"`

						Request *struct {
							AllowedHeaders *[]string `tfsdk:"allowed_headers" yaml:"allowedHeaders,omitempty"`

							AllowedHeadersRegex *[]string `tfsdk:"allowed_headers_regex" yaml:"allowedHeadersRegex,omitempty"`

							HeadersToAdd *map[string]string `tfsdk:"headers_to_add" yaml:"headersToAdd,omitempty"`
						} `tfsdk:"request" yaml:"request,omitempty"`

						Response *struct {
							AllowedClientHeaders *[]string `tfsdk:"allowed_client_headers" yaml:"allowedClientHeaders,omitempty"`

							AllowedUpstreamHeaders *[]string `tfsdk:"allowed_upstream_headers" yaml:"allowedUpstreamHeaders,omitempty"`

							AllowedUpstreamHeadersToAppend *[]string `tfsdk:"allowed_upstream_headers_to_append" yaml:"allowedUpstreamHeadersToAppend,omitempty"`
						} `tfsdk:"response" yaml:"response,omitempty"`
					} `tfsdk:"http_service" yaml:"httpService,omitempty"`

					RequestBody *struct {
						AllowPartialMessage *bool `tfsdk:"allow_partial_message" yaml:"allowPartialMessage,omitempty"`

						MaxRequestBytes *int64 `tfsdk:"max_request_bytes" yaml:"maxRequestBytes,omitempty"`

						PackAsBytes *bool `tfsdk:"pack_as_bytes" yaml:"packAsBytes,omitempty"`
					} `tfsdk:"request_body" yaml:"requestBody,omitempty"`

					RequestTimeout *string `tfsdk:"request_timeout" yaml:"requestTimeout,omitempty"`

					StatPrefix *string `tfsdk:"stat_prefix" yaml:"statPrefix,omitempty"`

					StatusOnError *int64 `tfsdk:"status_on_error" yaml:"statusOnError,omitempty"`

					TransportApiVersion utilities.IntOrString `tfsdk:"transport_api_version" yaml:"transportApiVersion,omitempty"`

					UserIdHeader *string `tfsdk:"user_id_header" yaml:"userIdHeader,omitempty"`
				} `tfsdk:"extauth" yaml:"extauth,omitempty"`

				Extensions *struct {
					Configs utilities.Dynamic `tfsdk:"configs" yaml:"configs,omitempty"`
				} `tfsdk:"extensions" yaml:"extensions,omitempty"`

				GrpcJsonTranscoder *struct {
					AutoMapping *bool `tfsdk:"auto_mapping" yaml:"autoMapping,omitempty"`

					ConvertGrpcStatus *bool `tfsdk:"convert_grpc_status" yaml:"convertGrpcStatus,omitempty"`

					IgnoreUnknownQueryParameters *bool `tfsdk:"ignore_unknown_query_parameters" yaml:"ignoreUnknownQueryParameters,omitempty"`

					IgnoredQueryParameters *[]string `tfsdk:"ignored_query_parameters" yaml:"ignoredQueryParameters,omitempty"`

					MatchIncomingRequestRoute *bool `tfsdk:"match_incoming_request_route" yaml:"matchIncomingRequestRoute,omitempty"`

					PrintOptions *struct {
						AddWhitespace *bool `tfsdk:"add_whitespace" yaml:"addWhitespace,omitempty"`

						AlwaysPrintEnumsAsInts *bool `tfsdk:"always_print_enums_as_ints" yaml:"alwaysPrintEnumsAsInts,omitempty"`

						AlwaysPrintPrimitiveFields *bool `tfsdk:"always_print_primitive_fields" yaml:"alwaysPrintPrimitiveFields,omitempty"`

						PreserveProtoFieldNames *bool `tfsdk:"preserve_proto_field_names" yaml:"preserveProtoFieldNames,omitempty"`
					} `tfsdk:"print_options" yaml:"printOptions,omitempty"`

					ProtoDescriptor *string `tfsdk:"proto_descriptor" yaml:"protoDescriptor,omitempty"`

					ProtoDescriptorBin *string `tfsdk:"proto_descriptor_bin" yaml:"protoDescriptorBin,omitempty"`

					Services *[]string `tfsdk:"services" yaml:"services,omitempty"`
				} `tfsdk:"grpc_json_transcoder" yaml:"grpcJsonTranscoder,omitempty"`

				GrpcWeb *struct {
					Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
				} `tfsdk:"grpc_web" yaml:"grpcWeb,omitempty"`

				Gzip *struct {
					CompressionLevel utilities.IntOrString `tfsdk:"compression_level" yaml:"compressionLevel,omitempty"`

					CompressionStrategy utilities.IntOrString `tfsdk:"compression_strategy" yaml:"compressionStrategy,omitempty"`

					ContentLength *int64 `tfsdk:"content_length" yaml:"contentLength,omitempty"`

					ContentType *[]string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					DisableOnEtagHeader *bool `tfsdk:"disable_on_etag_header" yaml:"disableOnEtagHeader,omitempty"`

					MemoryLevel *int64 `tfsdk:"memory_level" yaml:"memoryLevel,omitempty"`

					RemoveAcceptEncodingHeader *bool `tfsdk:"remove_accept_encoding_header" yaml:"removeAcceptEncodingHeader,omitempty"`

					WindowBits *int64 `tfsdk:"window_bits" yaml:"windowBits,omitempty"`
				} `tfsdk:"gzip" yaml:"gzip,omitempty"`

				HealthCheck *struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"health_check" yaml:"healthCheck,omitempty"`

				HttpConnectionManagerSettings *struct {
					AcceptHttp10 *bool `tfsdk:"accept_http10" yaml:"acceptHttp10,omitempty"`

					AllowChunkedLength *bool `tfsdk:"allow_chunked_length" yaml:"allowChunkedLength,omitempty"`

					CodecType utilities.IntOrString `tfsdk:"codec_type" yaml:"codecType,omitempty"`

					DefaultHostForHttp10 *string `tfsdk:"default_host_for_http10" yaml:"defaultHostForHttp10,omitempty"`

					DelayedCloseTimeout *string `tfsdk:"delayed_close_timeout" yaml:"delayedCloseTimeout,omitempty"`

					DrainTimeout *string `tfsdk:"drain_timeout" yaml:"drainTimeout,omitempty"`

					EnableTrailers *bool `tfsdk:"enable_trailers" yaml:"enableTrailers,omitempty"`

					ForwardClientCertDetails utilities.IntOrString `tfsdk:"forward_client_cert_details" yaml:"forwardClientCertDetails,omitempty"`

					GenerateRequestId *bool `tfsdk:"generate_request_id" yaml:"generateRequestId,omitempty"`

					HeadersWithUnderscoresAction utilities.IntOrString `tfsdk:"headers_with_underscores_action" yaml:"headersWithUnderscoresAction,omitempty"`

					Http2ProtocolOptions *struct {
						InitialConnectionWindowSize *int64 `tfsdk:"initial_connection_window_size" yaml:"initialConnectionWindowSize,omitempty"`

						InitialStreamWindowSize *int64 `tfsdk:"initial_stream_window_size" yaml:"initialStreamWindowSize,omitempty"`

						MaxConcurrentStreams *int64 `tfsdk:"max_concurrent_streams" yaml:"maxConcurrentStreams,omitempty"`

						OverrideStreamErrorOnInvalidHttpMessage *bool `tfsdk:"override_stream_error_on_invalid_http_message" yaml:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`
					} `tfsdk:"http2_protocol_options" yaml:"http2ProtocolOptions,omitempty"`

					IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

					InternalAddressConfig *struct {
						CidrRanges *[]struct {
							AddressPrefix *string `tfsdk:"address_prefix" yaml:"addressPrefix,omitempty"`

							PrefixLen *int64 `tfsdk:"prefix_len" yaml:"prefixLen,omitempty"`
						} `tfsdk:"cidr_ranges" yaml:"cidrRanges,omitempty"`

						UnixSockets *bool `tfsdk:"unix_sockets" yaml:"unixSockets,omitempty"`
					} `tfsdk:"internal_address_config" yaml:"internalAddressConfig,omitempty"`

					MaxConnectionDuration *string `tfsdk:"max_connection_duration" yaml:"maxConnectionDuration,omitempty"`

					MaxHeadersCount *int64 `tfsdk:"max_headers_count" yaml:"maxHeadersCount,omitempty"`

					MaxRequestHeadersKb *int64 `tfsdk:"max_request_headers_kb" yaml:"maxRequestHeadersKb,omitempty"`

					MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

					MaxStreamDuration *string `tfsdk:"max_stream_duration" yaml:"maxStreamDuration,omitempty"`

					MergeSlashes *bool `tfsdk:"merge_slashes" yaml:"mergeSlashes,omitempty"`

					NormalizePath *bool `tfsdk:"normalize_path" yaml:"normalizePath,omitempty"`

					PathWithEscapedSlashesAction utilities.IntOrString `tfsdk:"path_with_escaped_slashes_action" yaml:"pathWithEscapedSlashesAction,omitempty"`

					PreserveCaseHeaderKeyFormat *bool `tfsdk:"preserve_case_header_key_format" yaml:"preserveCaseHeaderKeyFormat,omitempty"`

					PreserveExternalRequestId *bool `tfsdk:"preserve_external_request_id" yaml:"preserveExternalRequestId,omitempty"`

					ProperCaseHeaderKeyFormat *bool `tfsdk:"proper_case_header_key_format" yaml:"properCaseHeaderKeyFormat,omitempty"`

					Proxy100Continue *bool `tfsdk:"proxy100_continue" yaml:"proxy100Continue,omitempty"`

					RequestHeadersTimeout *string `tfsdk:"request_headers_timeout" yaml:"requestHeadersTimeout,omitempty"`

					RequestTimeout *string `tfsdk:"request_timeout" yaml:"requestTimeout,omitempty"`

					ServerHeaderTransformation utilities.IntOrString `tfsdk:"server_header_transformation" yaml:"serverHeaderTransformation,omitempty"`

					ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`

					SetCurrentClientCertDetails *struct {
						Cert *bool `tfsdk:"cert" yaml:"cert,omitempty"`

						Chain *bool `tfsdk:"chain" yaml:"chain,omitempty"`

						Dns *bool `tfsdk:"dns" yaml:"dns,omitempty"`

						Subject *bool `tfsdk:"subject" yaml:"subject,omitempty"`

						Uri *bool `tfsdk:"uri" yaml:"uri,omitempty"`
					} `tfsdk:"set_current_client_cert_details" yaml:"setCurrentClientCertDetails,omitempty"`

					SkipXffAppend *bool `tfsdk:"skip_xff_append" yaml:"skipXffAppend,omitempty"`

					StreamIdleTimeout *string `tfsdk:"stream_idle_timeout" yaml:"streamIdleTimeout,omitempty"`

					StripAnyHostPort *bool `tfsdk:"strip_any_host_port" yaml:"stripAnyHostPort,omitempty"`

					Tracing *struct {
						DatadogConfig *struct {
							ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

							CollectorUpstreamRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"collector_upstream_ref" yaml:"collectorUpstreamRef,omitempty"`

							ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
						} `tfsdk:"datadog_config" yaml:"datadogConfig,omitempty"`

						EnvironmentVariablesForTags *[]struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`
						} `tfsdk:"environment_variables_for_tags" yaml:"environmentVariablesForTags,omitempty"`

						LiteralsForTags *[]struct {
							Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"literals_for_tags" yaml:"literalsForTags,omitempty"`

						RequestHeadersForTags *[]string `tfsdk:"request_headers_for_tags" yaml:"requestHeadersForTags,omitempty"`

						TracePercentages *struct {
							ClientSamplePercentage utilities.DynamicNumber `tfsdk:"client_sample_percentage" yaml:"clientSamplePercentage,omitempty"`

							OverallSamplePercentage utilities.DynamicNumber `tfsdk:"overall_sample_percentage" yaml:"overallSamplePercentage,omitempty"`

							RandomSamplePercentage utilities.DynamicNumber `tfsdk:"random_sample_percentage" yaml:"randomSamplePercentage,omitempty"`
						} `tfsdk:"trace_percentages" yaml:"tracePercentages,omitempty"`

						Verbose *bool `tfsdk:"verbose" yaml:"verbose,omitempty"`

						ZipkinConfig *struct {
							ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

							CollectorEndpoint *string `tfsdk:"collector_endpoint" yaml:"collectorEndpoint,omitempty"`

							CollectorEndpointVersion utilities.IntOrString `tfsdk:"collector_endpoint_version" yaml:"collectorEndpointVersion,omitempty"`

							CollectorUpstreamRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"collector_upstream_ref" yaml:"collectorUpstreamRef,omitempty"`

							SharedSpanContext *bool `tfsdk:"shared_span_context" yaml:"sharedSpanContext,omitempty"`

							TraceId128bit *bool `tfsdk:"trace_id128bit" yaml:"traceId128bit,omitempty"`
						} `tfsdk:"zipkin_config" yaml:"zipkinConfig,omitempty"`
					} `tfsdk:"tracing" yaml:"tracing,omitempty"`

					Upgrades *[]struct {
						Websocket *struct {
							Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
						} `tfsdk:"websocket" yaml:"websocket,omitempty"`
					} `tfsdk:"upgrades" yaml:"upgrades,omitempty"`

					UseRemoteAddress *bool `tfsdk:"use_remote_address" yaml:"useRemoteAddress,omitempty"`

					UuidRequestIdConfig *struct {
						PackTraceReason *bool `tfsdk:"pack_trace_reason" yaml:"packTraceReason,omitempty"`

						UseRequestIdForTraceSampling *bool `tfsdk:"use_request_id_for_trace_sampling" yaml:"useRequestIdForTraceSampling,omitempty"`
					} `tfsdk:"uuid_request_id_config" yaml:"uuidRequestIdConfig,omitempty"`

					Via *string `tfsdk:"via" yaml:"via,omitempty"`

					XffNumTrustedHops *int64 `tfsdk:"xff_num_trusted_hops" yaml:"xffNumTrustedHops,omitempty"`
				} `tfsdk:"http_connection_manager_settings" yaml:"httpConnectionManagerSettings,omitempty"`

				LeftmostXffAddress *bool `tfsdk:"leftmost_xff_address" yaml:"leftmostXffAddress,omitempty"`

				ProxyLatency *struct {
					ChargeClusterStat *bool `tfsdk:"charge_cluster_stat" yaml:"chargeClusterStat,omitempty"`

					ChargeListenerStat *bool `tfsdk:"charge_listener_stat" yaml:"chargeListenerStat,omitempty"`

					EmitDynamicMetadata *bool `tfsdk:"emit_dynamic_metadata" yaml:"emitDynamicMetadata,omitempty"`

					MeasureRequestInternally *bool `tfsdk:"measure_request_internally" yaml:"measureRequestInternally,omitempty"`

					Request utilities.IntOrString `tfsdk:"request" yaml:"request,omitempty"`

					Response utilities.IntOrString `tfsdk:"response" yaml:"response,omitempty"`
				} `tfsdk:"proxy_latency" yaml:"proxyLatency,omitempty"`

				RatelimitServer *struct {
					DenyOnFail *bool `tfsdk:"deny_on_fail" yaml:"denyOnFail,omitempty"`

					EnableXRatelimitHeaders *bool `tfsdk:"enable_x_ratelimit_headers" yaml:"enableXRatelimitHeaders,omitempty"`

					RateLimitBeforeAuth *bool `tfsdk:"rate_limit_before_auth" yaml:"rateLimitBeforeAuth,omitempty"`

					RatelimitServerRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"ratelimit_server_ref" yaml:"ratelimitServerRef,omitempty"`

					RequestTimeout *string `tfsdk:"request_timeout" yaml:"requestTimeout,omitempty"`
				} `tfsdk:"ratelimit_server" yaml:"ratelimitServer,omitempty"`

				SanitizeClusterHeader *bool `tfsdk:"sanitize_cluster_header" yaml:"sanitizeClusterHeader,omitempty"`

				Waf *struct {
					AuditLogging *struct {
						Action utilities.IntOrString `tfsdk:"action" yaml:"action,omitempty"`

						Location utilities.IntOrString `tfsdk:"location" yaml:"location,omitempty"`
					} `tfsdk:"audit_logging" yaml:"auditLogging,omitempty"`

					ConfigMapRuleSets *[]struct {
						ConfigMapRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						DataMapKeys *[]string `tfsdk:"data_map_keys" yaml:"dataMapKeys,omitempty"`
					} `tfsdk:"config_map_rule_sets" yaml:"configMapRuleSets,omitempty"`

					CoreRuleSet *struct {
						CustomSettingsFile *string `tfsdk:"custom_settings_file" yaml:"customSettingsFile,omitempty"`

						CustomSettingsString *string `tfsdk:"custom_settings_string" yaml:"customSettingsString,omitempty"`
					} `tfsdk:"core_rule_set" yaml:"coreRuleSet,omitempty"`

					CustomInterventionMessage *string `tfsdk:"custom_intervention_message" yaml:"customInterventionMessage,omitempty"`

					Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

					RequestHeadersOnly *bool `tfsdk:"request_headers_only" yaml:"requestHeadersOnly,omitempty"`

					ResponseHeadersOnly *bool `tfsdk:"response_headers_only" yaml:"responseHeadersOnly,omitempty"`

					RuleSets *[]struct {
						Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

						Files *[]string `tfsdk:"files" yaml:"files,omitempty"`

						RuleStr *string `tfsdk:"rule_str" yaml:"ruleStr,omitempty"`
					} `tfsdk:"rule_sets" yaml:"ruleSets,omitempty"`
				} `tfsdk:"waf" yaml:"waf,omitempty"`

				Wasm *struct {
					Filters *[]struct {
						Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

						FailOpen *bool `tfsdk:"fail_open" yaml:"failOpen,omitempty"`

						FilePath *string `tfsdk:"file_path" yaml:"filePath,omitempty"`

						FilterStage *struct {
							Predicate utilities.IntOrString `tfsdk:"predicate" yaml:"predicate,omitempty"`

							Stage utilities.IntOrString `tfsdk:"stage" yaml:"stage,omitempty"`
						} `tfsdk:"filter_stage" yaml:"filterStage,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						RootId *string `tfsdk:"root_id" yaml:"rootId,omitempty"`

						VmType utilities.IntOrString `tfsdk:"vm_type" yaml:"vmType,omitempty"`
					} `tfsdk:"filters" yaml:"filters,omitempty"`
				} `tfsdk:"wasm" yaml:"wasm,omitempty"`
			} `tfsdk:"options" yaml:"options,omitempty"`

			VirtualServiceExpressions *struct {
				Expressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator utilities.IntOrString `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"expressions" yaml:"expressions,omitempty"`
			} `tfsdk:"virtual_service_expressions" yaml:"virtualServiceExpressions,omitempty"`

			VirtualServiceNamespaces *[]string `tfsdk:"virtual_service_namespaces" yaml:"virtualServiceNamespaces,omitempty"`

			VirtualServiceSelector *map[string]string `tfsdk:"virtual_service_selector" yaml:"virtualServiceSelector,omitempty"`

			VirtualServices *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"virtual_services" yaml:"virtualServices,omitempty"`
		} `tfsdk:"http_gateway" yaml:"httpGateway,omitempty"`

		Matcher *struct {
			SourcePrefixRanges *[]struct {
				AddressPrefix *string `tfsdk:"address_prefix" yaml:"addressPrefix,omitempty"`

				PrefixLen *int64 `tfsdk:"prefix_len" yaml:"prefixLen,omitempty"`
			} `tfsdk:"source_prefix_ranges" yaml:"sourcePrefixRanges,omitempty"`

			SslConfig *struct {
				AlpnProtocols *[]string `tfsdk:"alpn_protocols" yaml:"alpnProtocols,omitempty"`

				DisableTlsSessionResumption *bool `tfsdk:"disable_tls_session_resumption" yaml:"disableTlsSessionResumption,omitempty"`

				OneWayTls *bool `tfsdk:"one_way_tls" yaml:"oneWayTls,omitempty"`

				Parameters *struct {
					CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

					EcdhCurves *[]string `tfsdk:"ecdh_curves" yaml:"ecdhCurves,omitempty"`

					MaximumProtocolVersion utilities.IntOrString `tfsdk:"maximum_protocol_version" yaml:"maximumProtocolVersion,omitempty"`

					MinimumProtocolVersion utilities.IntOrString `tfsdk:"minimum_protocol_version" yaml:"minimumProtocolVersion,omitempty"`
				} `tfsdk:"parameters" yaml:"parameters,omitempty"`

				Sds *struct {
					CallCredentials *struct {
						FileCredentialSource *struct {
							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							TokenFileName *string `tfsdk:"token_file_name" yaml:"tokenFileName,omitempty"`
						} `tfsdk:"file_credential_source" yaml:"fileCredentialSource,omitempty"`
					} `tfsdk:"call_credentials" yaml:"callCredentials,omitempty"`

					CertificatesSecretName *string `tfsdk:"certificates_secret_name" yaml:"certificatesSecretName,omitempty"`

					ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

					TargetUri *string `tfsdk:"target_uri" yaml:"targetUri,omitempty"`

					ValidationContextName *string `tfsdk:"validation_context_name" yaml:"validationContextName,omitempty"`
				} `tfsdk:"sds" yaml:"sds,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				SniDomains *[]string `tfsdk:"sni_domains" yaml:"sniDomains,omitempty"`

				SslFiles *struct {
					RootCa *string `tfsdk:"root_ca" yaml:"rootCa,omitempty"`

					TlsCert *string `tfsdk:"tls_cert" yaml:"tlsCert,omitempty"`

					TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`
				} `tfsdk:"ssl_files" yaml:"sslFiles,omitempty"`

				TransportSocketConnectTimeout *string `tfsdk:"transport_socket_connect_timeout" yaml:"transportSocketConnectTimeout,omitempty"`

				VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" yaml:"verifySubjectAltName,omitempty"`
			} `tfsdk:"ssl_config" yaml:"sslConfig,omitempty"`
		} `tfsdk:"matcher" yaml:"matcher,omitempty"`

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGatewaySoloIoMatchableHttpGatewayV1Resource() resource.Resource {
	return &GatewaySoloIoMatchableHttpGatewayV1Resource{}
}

func (r *GatewaySoloIoMatchableHttpGatewayV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_solo_io_matchable_http_gateway_v1"
}

func (r *GatewaySoloIoMatchableHttpGatewayV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"http_gateway": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"buffer": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_request_bytes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"caching": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allowed_vary_headers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"exact": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_case": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"safe_regex": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"google_re2": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"max_program_size": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			int64validator.AtLeast(0),

																			int64validator.AtMost(4.294967295e+09),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
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

													"suffix": {
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

											"caching_service_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
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

											"max_payload_size": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"value": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout": {
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

									"csrf": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"additional_origins": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"exact": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_case": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"safe_regex": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"google_re2": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"max_program_size": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			int64validator.AtLeast(0),

																			int64validator.AtMost(4.294967295e+09),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
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

													"suffix": {
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

											"filter_enabled": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_value": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"denominator": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"numerator": {
																Description:         "",
																MarkdownDescription: "",

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

													"runtime_key": {
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

											"shadow_enabled": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_value": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"denominator": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"numerator": {
																Description:         "",
																MarkdownDescription: "",

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

													"runtime_key": {
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

									"dlp": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dlp_rules": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"action_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"custom_action": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mask_char": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"percent": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.DynamicNumberType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_actions": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"key_value_action": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key_to_mask": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"mask_char": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"percent": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.DynamicNumberType{},

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

															"shadow": {
																Description:         "",
																MarkdownDescription: "",

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

													"matcher": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"case_sensitive": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"exact": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

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

															"methods": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_parameters": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

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

															"regex": {
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

											"enabled_for": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dynamic_forward_proxy": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns_cache_config": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"apple_dns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cares_dns": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"dns_resolver_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"no_default_search_domain": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"use_tcp_for_dns_lookups": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"resolvers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pipe": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"mode": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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

																	"socket_address": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"address": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ipv4_compat": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"named_port": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"port_value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"protocol": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resolver_name": {
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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns_cache_circuit_breaker": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"max_pending_requests": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(0),

																	int64validator.AtMost(4.294967295e+09),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns_failure_refresh_rate": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"base_interval": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_interval": {
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

													"dns_lookup_family": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns_query_timeout": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns_refresh_rate": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_ttl": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_hosts": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
														},
													},

													"preresolve_hostnames": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ipv4_compat": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"named_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"protocol": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"resolver_name": {
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

											"save_upstream_address": {
												Description:         "",
												MarkdownDescription: "",

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

									"extauth": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"clear_route_cache": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"extauthz_server_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
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

											"failure_mode_allow": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"grpc_service": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"authority": {
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

											"http_service": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path_prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allowed_headers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"allowed_headers_regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers_to_add": {
																Description:         "",
																MarkdownDescription: "",

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

													"response": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allowed_client_headers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"allowed_upstream_headers": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"allowed_upstream_headers_to_append": {
																Description:         "",
																MarkdownDescription: "",

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

											"request_body": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_partial_message": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_request_bytes": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pack_as_bytes": {
														Description:         "",
														MarkdownDescription: "",

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

											"request_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stat_prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"status_on_error": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"transport_api_version": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user_id_header": {
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

									"extensions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"configs": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc_json_transcoder": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"auto_mapping": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"convert_grpc_status": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ignore_unknown_query_parameters": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ignored_query_parameters": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_incoming_request_route": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"print_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"add_whitespace": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"always_print_enums_as_ints": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"always_print_primitive_fields": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"preserve_proto_field_names": {
														Description:         "",
														MarkdownDescription: "",

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

											"proto_descriptor": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proto_descriptor_bin": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"services": {
												Description:         "",
												MarkdownDescription: "",

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

									"grpc_web": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"disable": {
												Description:         "",
												MarkdownDescription: "",

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

									"gzip": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"compression_level": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"compression_strategy": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_length": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disable_on_etag_header": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory_level": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},

											"remove_accept_encoding_header": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"window_bits": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"health_check": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
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

									"http_connection_manager_settings": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"accept_http10": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allow_chunked_length": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"codec_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"default_host_for_http10": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delayed_close_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"drain_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_trailers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"forward_client_cert_details": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"generate_request_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers_with_underscores_action": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http2_protocol_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"initial_connection_window_size": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
														},
													},

													"initial_stream_window_size": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
														},
													},

													"max_concurrent_streams": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
														},
													},

													"override_stream_error_on_invalid_http_message": {
														Description:         "",
														MarkdownDescription: "",

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

											"idle_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"internal_address_config": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cidr_ranges": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"address_prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix_len": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(0),

																	int64validator.AtMost(4.294967295e+09),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"unix_sockets": {
														Description:         "",
														MarkdownDescription: "",

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

											"max_connection_duration": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_headers_count": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},

											"max_request_headers_kb": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},

											"max_requests_per_connection": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},

											"max_stream_duration": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"merge_slashes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"normalize_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path_with_escaped_slashes_action": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preserve_case_header_key_format": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preserve_external_request_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proper_case_header_key_format": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy100_continue": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_headers_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_header_transformation": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set_current_client_cert_details": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cert": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"chain": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"subject": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"uri": {
														Description:         "",
														MarkdownDescription: "",

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

											"skip_xff_append": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stream_idle_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"strip_any_host_port": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tracing": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"datadog_config": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"cluster_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"collector_upstream_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace": {
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

															"service_name": {
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

													"environment_variables_for_tags": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tag": {
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

													"literals_for_tags": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"tag": {
																Description:         "",
																MarkdownDescription: "",

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

													"request_headers_for_tags": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"trace_percentages": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"client_sample_percentage": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.DynamicNumberType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"overall_sample_percentage": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.DynamicNumberType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"random_sample_percentage": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.DynamicNumberType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"verbose": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"zipkin_config": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"cluster_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"collector_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"collector_endpoint_version": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"collector_upstream_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace": {
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

															"shared_span_context": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"trace_id128bit": {
																Description:         "",
																MarkdownDescription: "",

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

											"upgrades": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"websocket": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"enabled": {
																Description:         "",
																MarkdownDescription: "",

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

											"use_remote_address": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"uuid_request_id_config": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"pack_trace_reason": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_request_id_for_trace_sampling": {
														Description:         "",
														MarkdownDescription: "",

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

											"via": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"xff_num_trusted_hops": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"leftmost_xff_address": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"proxy_latency": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"charge_cluster_stat": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"charge_listener_stat": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"emit_dynamic_metadata": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"measure_request_internally": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"response": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ratelimit_server": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"deny_on_fail": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_x_ratelimit_headers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rate_limit_before_auth": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ratelimit_server_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
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

											"request_timeout": {
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

									"sanitize_cluster_header": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"waf": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"audit_logging": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"action": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"location": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"config_map_rule_sets": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
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

													"data_map_keys": {
														Description:         "",
														MarkdownDescription: "",

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

											"core_rule_set": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"custom_settings_file": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom_settings_string": {
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

											"custom_intervention_message": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disabled": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_headers_only": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_headers_only": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rule_sets": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"directory": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"files": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"rule_str": {
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

									"wasm": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"filters": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"fail_open": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"file_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"filter_stage": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"predicate": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"stage": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

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
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"root_id": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"vm_type": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

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

							"virtual_service_expressions": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"expressions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"values": {
												Description:         "",
												MarkdownDescription: "",

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

							"virtual_service_namespaces": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"virtual_service_selector": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"virtual_services": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
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

					"matcher": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"source_prefix_ranges": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"address_prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix_len": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),

											int64validator.AtMost(4.294967295e+09),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_config": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"alpn_protocols": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_tls_session_resumption": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"one_way_tls": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"parameters": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cipher_suites": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ecdh_curves": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maximum_protocol_version": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum_protocol_version": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sds": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"call_credentials": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"file_credential_source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"token_file_name": {
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

											"certificates_secret_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_uri": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"validation_context_name": {
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

									"secret_ref": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
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

									"sni_domains": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ssl_files": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"root_ca": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_cert": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_key": {
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

									"transport_socket_connect_timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify_subject_alt_name": {
										Description:         "",
										MarkdownDescription: "",

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

					"namespaced_statuses": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"statuses": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

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
		},
	}, nil
}

func (r *GatewaySoloIoMatchableHttpGatewayV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gateway_solo_io_matchable_http_gateway_v1")

	var state GatewaySoloIoMatchableHttpGatewayV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewaySoloIoMatchableHttpGatewayV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.solo.io/v1")
	goModel.Kind = utilities.Ptr("MatchableHttpGateway")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GatewaySoloIoMatchableHttpGatewayV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_solo_io_matchable_http_gateway_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GatewaySoloIoMatchableHttpGatewayV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gateway_solo_io_matchable_http_gateway_v1")

	var state GatewaySoloIoMatchableHttpGatewayV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewaySoloIoMatchableHttpGatewayV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.solo.io/v1")
	goModel.Kind = utilities.Ptr("MatchableHttpGateway")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GatewaySoloIoMatchableHttpGatewayV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gateway_solo_io_matchable_http_gateway_v1")
	// NO-OP: Terraform removes the state automatically for us
}
