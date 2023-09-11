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
	"time"
)

var (
	_ resource.Resource                = &GlooSoloIoSettingsV1Resource{}
	_ resource.ResourceWithConfigure   = &GlooSoloIoSettingsV1Resource{}
	_ resource.ResourceWithImportState = &GlooSoloIoSettingsV1Resource{}
)

func NewGlooSoloIoSettingsV1Resource() resource.Resource {
	return &GlooSoloIoSettingsV1Resource{}
}

type GlooSoloIoSettingsV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type GlooSoloIoSettingsV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CachingServer *struct {
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
		} `tfsdk:"caching_server" json:"cachingServer,omitempty"`
		ConsoleOptions *struct {
			ApiExplorerEnabled *bool `tfsdk:"api_explorer_enabled" json:"apiExplorerEnabled,omitempty"`
			ReadOnly           *bool `tfsdk:"read_only" json:"readOnly,omitempty"`
		} `tfsdk:"console_options" json:"consoleOptions,omitempty"`
		Consul *struct {
			Address            *string `tfsdk:"address" json:"address,omitempty"`
			CaFile             *string `tfsdk:"ca_file" json:"caFile,omitempty"`
			CaPath             *string `tfsdk:"ca_path" json:"caPath,omitempty"`
			CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
			Datacenter         *string `tfsdk:"datacenter" json:"datacenter,omitempty"`
			DnsAddress         *string `tfsdk:"dns_address" json:"dnsAddress,omitempty"`
			DnsPollingInterval *string `tfsdk:"dns_polling_interval" json:"dnsPollingInterval,omitempty"`
			HttpAddress        *string `tfsdk:"http_address" json:"httpAddress,omitempty"`
			InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
			Password           *string `tfsdk:"password" json:"password,omitempty"`
			ServiceDiscovery   *struct {
				DataCenters *[]string `tfsdk:"data_centers" json:"dataCenters,omitempty"`
			} `tfsdk:"service_discovery" json:"serviceDiscovery,omitempty"`
			Token    *string `tfsdk:"token" json:"token,omitempty"`
			Username *string `tfsdk:"username" json:"username,omitempty"`
			WaitTime *string `tfsdk:"wait_time" json:"waitTime,omitempty"`
		} `tfsdk:"consul" json:"consul,omitempty"`
		ConsulDiscovery *struct {
			ConsistencyMode    *string `tfsdk:"consistency_mode" json:"consistencyMode,omitempty"`
			EdsBlockingQueries *bool   `tfsdk:"eds_blocking_queries" json:"edsBlockingQueries,omitempty"`
			QueryOptions       *struct {
				UseCache *bool `tfsdk:"use_cache" json:"useCache,omitempty"`
			} `tfsdk:"query_options" json:"queryOptions,omitempty"`
			RootCa *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"root_ca" json:"rootCa,omitempty"`
			ServiceTagsAllowlist *[]string `tfsdk:"service_tags_allowlist" json:"serviceTagsAllowlist,omitempty"`
			SplitTlsServices     *bool     `tfsdk:"split_tls_services" json:"splitTlsServices,omitempty"`
			TlsTagName           *string   `tfsdk:"tls_tag_name" json:"tlsTagName,omitempty"`
			UseTlsTagging        *bool     `tfsdk:"use_tls_tagging" json:"useTlsTagging,omitempty"`
		} `tfsdk:"consul_discovery" json:"consulDiscovery,omitempty"`
		ConsulKvArtifactSource *struct {
			RootKey *string `tfsdk:"root_key" json:"rootKey,omitempty"`
		} `tfsdk:"consul_kv_artifact_source" json:"consulKvArtifactSource,omitempty"`
		ConsulKvSource *struct {
			RootKey *string `tfsdk:"root_key" json:"rootKey,omitempty"`
		} `tfsdk:"consul_kv_source" json:"consulKvSource,omitempty"`
		DevMode                 *bool `tfsdk:"dev_mode" json:"devMode,omitempty"`
		DirectoryArtifactSource *struct {
			Directory *string `tfsdk:"directory" json:"directory,omitempty"`
		} `tfsdk:"directory_artifact_source" json:"directoryArtifactSource,omitempty"`
		DirectoryConfigSource *struct {
			Directory *string `tfsdk:"directory" json:"directory,omitempty"`
		} `tfsdk:"directory_config_source" json:"directoryConfigSource,omitempty"`
		DirectorySecretSource *struct {
			Directory *string `tfsdk:"directory" json:"directory,omitempty"`
		} `tfsdk:"directory_secret_source" json:"directorySecretSource,omitempty"`
		Discovery *struct {
			FdsMode    *string `tfsdk:"fds_mode" json:"fdsMode,omitempty"`
			FdsOptions *struct {
				GraphqlEnabled *bool `tfsdk:"graphql_enabled" json:"graphqlEnabled,omitempty"`
			} `tfsdk:"fds_options" json:"fdsOptions,omitempty"`
			UdsOptions *struct {
				Enabled     *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				WatchLabels *map[string]string `tfsdk:"watch_labels" json:"watchLabels,omitempty"`
			} `tfsdk:"uds_options" json:"udsOptions,omitempty"`
		} `tfsdk:"discovery" json:"discovery,omitempty"`
		DiscoveryNamespace *string `tfsdk:"discovery_namespace" json:"discoveryNamespace,omitempty"`
		ExtProc            *struct {
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
		Gateway *struct {
			AlwaysSortRouteTableRoutes     *bool `tfsdk:"always_sort_route_table_routes" json:"alwaysSortRouteTableRoutes,omitempty"`
			CompressedProxySpec            *bool `tfsdk:"compressed_proxy_spec" json:"compressedProxySpec,omitempty"`
			EnableGatewayController        *bool `tfsdk:"enable_gateway_controller" json:"enableGatewayController,omitempty"`
			IsolateVirtualHostsBySslConfig *bool `tfsdk:"isolate_virtual_hosts_by_ssl_config" json:"isolateVirtualHostsBySslConfig,omitempty"`
			PersistProxySpec               *bool `tfsdk:"persist_proxy_spec" json:"persistProxySpec,omitempty"`
			ReadGatewaysFromAllNamespaces  *bool `tfsdk:"read_gateways_from_all_namespaces" json:"readGatewaysFromAllNamespaces,omitempty"`
			Validation                     *struct {
				AllowWarnings                    *bool   `tfsdk:"allow_warnings" json:"allowWarnings,omitempty"`
				AlwaysAccept                     *bool   `tfsdk:"always_accept" json:"alwaysAccept,omitempty"`
				DisableTransformationValidation  *bool   `tfsdk:"disable_transformation_validation" json:"disableTransformationValidation,omitempty"`
				IgnoreGlooValidationFailure      *bool   `tfsdk:"ignore_gloo_validation_failure" json:"ignoreGlooValidationFailure,omitempty"`
				ProxyValidationServerAddr        *string `tfsdk:"proxy_validation_server_addr" json:"proxyValidationServerAddr,omitempty"`
				ServerEnabled                    *bool   `tfsdk:"server_enabled" json:"serverEnabled,omitempty"`
				ValidationServerGrpcMaxSizeBytes *int64  `tfsdk:"validation_server_grpc_max_size_bytes" json:"validationServerGrpcMaxSizeBytes,omitempty"`
				ValidationWebhookTlsCert         *string `tfsdk:"validation_webhook_tls_cert" json:"validationWebhookTlsCert,omitempty"`
				ValidationWebhookTlsKey          *string `tfsdk:"validation_webhook_tls_key" json:"validationWebhookTlsKey,omitempty"`
				WarnRouteShortCircuiting         *bool   `tfsdk:"warn_route_short_circuiting" json:"warnRouteShortCircuiting,omitempty"`
			} `tfsdk:"validation" json:"validation,omitempty"`
			ValidationServerAddr  *string `tfsdk:"validation_server_addr" json:"validationServerAddr,omitempty"`
			VirtualServiceOptions *struct {
				OneWayTls *bool `tfsdk:"one_way_tls" json:"oneWayTls,omitempty"`
			} `tfsdk:"virtual_service_options" json:"virtualServiceOptions,omitempty"`
		} `tfsdk:"gateway" json:"gateway,omitempty"`
		Gloo *struct {
			AwsOptions *struct {
				CredentialRefreshDelay    *string `tfsdk:"credential_refresh_delay" json:"credentialRefreshDelay,omitempty"`
				EnableCredentialsDiscovey *bool   `tfsdk:"enable_credentials_discovey" json:"enableCredentialsDiscovey,omitempty"`
				FallbackToFirstFunction   *bool   `tfsdk:"fallback_to_first_function" json:"fallbackToFirstFunction,omitempty"`
				PropagateOriginalRouting  *bool   `tfsdk:"propagate_original_routing" json:"propagateOriginalRouting,omitempty"`
				ServiceAccountCredentials *struct {
					Cluster *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"service_account_credentials" json:"serviceAccountCredentials,omitempty"`
			} `tfsdk:"aws_options" json:"awsOptions,omitempty"`
			CircuitBreakers *struct {
				MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
				MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
				MaxRequests        *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
				MaxRetries         *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			} `tfsdk:"circuit_breakers" json:"circuitBreakers,omitempty"`
			DisableGrpcWeb                     *bool   `tfsdk:"disable_grpc_web" json:"disableGrpcWeb,omitempty"`
			DisableKubernetesDestinations      *bool   `tfsdk:"disable_kubernetes_destinations" json:"disableKubernetesDestinations,omitempty"`
			DisableProxyGarbageCollection      *bool   `tfsdk:"disable_proxy_garbage_collection" json:"disableProxyGarbageCollection,omitempty"`
			EnableRestEds                      *bool   `tfsdk:"enable_rest_eds" json:"enableRestEds,omitempty"`
			EndpointsWarmingTimeout            *string `tfsdk:"endpoints_warming_timeout" json:"endpointsWarmingTimeout,omitempty"`
			FailoverUpstreamDnsPollingInterval *string `tfsdk:"failover_upstream_dns_polling_interval" json:"failoverUpstreamDnsPollingInterval,omitempty"`
			InvalidConfigPolicy                *struct {
				InvalidRouteResponseBody *string `tfsdk:"invalid_route_response_body" json:"invalidRouteResponseBody,omitempty"`
				InvalidRouteResponseCode *int64  `tfsdk:"invalid_route_response_code" json:"invalidRouteResponseCode,omitempty"`
				ReplaceInvalidRoutes     *bool   `tfsdk:"replace_invalid_routes" json:"replaceInvalidRoutes,omitempty"`
			} `tfsdk:"invalid_config_policy" json:"invalidConfigPolicy,omitempty"`
			LogTransformationRequestResponseInfo *bool   `tfsdk:"log_transformation_request_response_info" json:"logTransformationRequestResponseInfo,omitempty"`
			ProxyDebugBindAddr                   *string `tfsdk:"proxy_debug_bind_addr" json:"proxyDebugBindAddr,omitempty"`
			RegexMaxProgramSize                  *int64  `tfsdk:"regex_max_program_size" json:"regexMaxProgramSize,omitempty"`
			RemoveUnusedFilters                  *bool   `tfsdk:"remove_unused_filters" json:"removeUnusedFilters,omitempty"`
			RestXdsBindAddr                      *string `tfsdk:"rest_xds_bind_addr" json:"restXdsBindAddr,omitempty"`
			TransformationEscapeCharacters       *bool   `tfsdk:"transformation_escape_characters" json:"transformationEscapeCharacters,omitempty"`
			ValidationBindAddr                   *string `tfsdk:"validation_bind_addr" json:"validationBindAddr,omitempty"`
			XdsBindAddr                          *string `tfsdk:"xds_bind_addr" json:"xdsBindAddr,omitempty"`
		} `tfsdk:"gloo" json:"gloo,omitempty"`
		GraphqlOptions *struct {
			SchemaChangeValidationOptions *struct {
				ProcessingRules       *[]string `tfsdk:"processing_rules" json:"processingRules,omitempty"`
				RejectBreakingChanges *bool     `tfsdk:"reject_breaking_changes" json:"rejectBreakingChanges,omitempty"`
			} `tfsdk:"schema_change_validation_options" json:"schemaChangeValidationOptions,omitempty"`
		} `tfsdk:"graphql_options" json:"graphqlOptions,omitempty"`
		Knative *struct {
			ClusterIngressProxyAddress  *string `tfsdk:"cluster_ingress_proxy_address" json:"clusterIngressProxyAddress,omitempty"`
			KnativeExternalProxyAddress *string `tfsdk:"knative_external_proxy_address" json:"knativeExternalProxyAddress,omitempty"`
			KnativeInternalProxyAddress *string `tfsdk:"knative_internal_proxy_address" json:"knativeInternalProxyAddress,omitempty"`
		} `tfsdk:"knative" json:"knative,omitempty"`
		Kubernetes *struct {
			RateLimits *struct {
				QPS   *float64 `tfsdk:"qps" json:"QPS,omitempty"`
				Burst *int64   `tfsdk:"burst" json:"burst,omitempty"`
			} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		KubernetesArtifactSource *map[string]string `tfsdk:"kubernetes_artifact_source" json:"kubernetesArtifactSource,omitempty"`
		KubernetesConfigSource   *map[string]string `tfsdk:"kubernetes_config_source" json:"kubernetesConfigSource,omitempty"`
		KubernetesSecretSource   *map[string]string `tfsdk:"kubernetes_secret_source" json:"kubernetesSecretSource,omitempty"`
		Linkerd                  *bool              `tfsdk:"linkerd" json:"linkerd,omitempty"`
		NamedExtauth             *struct {
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
		} `tfsdk:"named_extauth" json:"namedExtauth,omitempty"`
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
		ObservabilityOptions *struct {
			ConfigStatusMetricLabels *struct {
				LabelToPath *map[string]string `tfsdk:"label_to_path" json:"labelToPath,omitempty"`
			} `tfsdk:"config_status_metric_labels" json:"configStatusMetricLabels,omitempty"`
			GrafanaIntegration *struct {
				DefaultDashboardFolderId *int64 `tfsdk:"default_dashboard_folder_id" json:"defaultDashboardFolderId,omitempty"`
			} `tfsdk:"grafana_integration" json:"grafanaIntegration,omitempty"`
		} `tfsdk:"observability_options" json:"observabilityOptions,omitempty"`
		Ratelimit *struct {
			Descriptors    *[]map[string]string `tfsdk:"descriptors" json:"descriptors,omitempty"`
			SetDescriptors *[]struct {
				AlwaysApply *bool `tfsdk:"always_apply" json:"alwaysApply,omitempty"`
				RateLimit   *struct {
					RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
					Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"rate_limit" json:"rateLimit,omitempty"`
				SimpleDescriptors *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"simple_descriptors" json:"simpleDescriptors,omitempty"`
			} `tfsdk:"set_descriptors" json:"setDescriptors,omitempty"`
		} `tfsdk:"ratelimit" json:"ratelimit,omitempty"`
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
		Rbac *struct {
			RequireRbac *bool `tfsdk:"require_rbac" json:"requireRbac,omitempty"`
		} `tfsdk:"rbac" json:"rbac,omitempty"`
		RefreshRate   *string `tfsdk:"refresh_rate" json:"refreshRate,omitempty"`
		SecretOptions *struct {
			Sources *[]struct {
				Directory *struct {
					Directory *string `tfsdk:"directory" json:"directory,omitempty"`
				} `tfsdk:"directory" json:"directory,omitempty"`
				Kubernetes *map[string]string `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
				Vault      *struct {
					AccessToken *string `tfsdk:"access_token" json:"accessToken,omitempty"`
					Address     *string `tfsdk:"address" json:"address,omitempty"`
					Aws         *struct {
						AccessKeyId       *string `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
						IamServerIdHeader *string `tfsdk:"iam_server_id_header" json:"iamServerIdHeader,omitempty"`
						MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						Region            *string `tfsdk:"region" json:"region,omitempty"`
						SecretAccessKey   *string `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
						SessionToken      *string `tfsdk:"session_token" json:"sessionToken,omitempty"`
						VaultRole         *string `tfsdk:"vault_role" json:"vaultRole,omitempty"`
					} `tfsdk:"aws" json:"aws,omitempty"`
					CaCert     *string `tfsdk:"ca_cert" json:"caCert,omitempty"`
					CaPath     *string `tfsdk:"ca_path" json:"caPath,omitempty"`
					ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
					ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
					Insecure   *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					PathPrefix *string `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
					RootKey    *string `tfsdk:"root_key" json:"rootKey,omitempty"`
					TlsConfig  *struct {
						CaCert        *string `tfsdk:"ca_cert" json:"caCert,omitempty"`
						CaPath        *string `tfsdk:"ca_path" json:"caPath,omitempty"`
						ClientCert    *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey     *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Insecure      *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
						TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
					TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
					Token         *string `tfsdk:"token" json:"token,omitempty"`
				} `tfsdk:"vault" json:"vault,omitempty"`
			} `tfsdk:"sources" json:"sources,omitempty"`
		} `tfsdk:"secret_options" json:"secretOptions,omitempty"`
		UpstreamOptions *struct {
			GlobalAnnotations *map[string]string `tfsdk:"global_annotations" json:"globalAnnotations,omitempty"`
			SslParameters     *struct {
				CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				EcdhCurves             *[]string `tfsdk:"ecdh_curves" json:"ecdhCurves,omitempty"`
				MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
				MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"ssl_parameters" json:"sslParameters,omitempty"`
		} `tfsdk:"upstream_options" json:"upstreamOptions,omitempty"`
		VaultSecretSource *struct {
			AccessToken *string `tfsdk:"access_token" json:"accessToken,omitempty"`
			Address     *string `tfsdk:"address" json:"address,omitempty"`
			Aws         *struct {
				AccessKeyId       *string `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
				IamServerIdHeader *string `tfsdk:"iam_server_id_header" json:"iamServerIdHeader,omitempty"`
				MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Region            *string `tfsdk:"region" json:"region,omitempty"`
				SecretAccessKey   *string `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
				SessionToken      *string `tfsdk:"session_token" json:"sessionToken,omitempty"`
				VaultRole         *string `tfsdk:"vault_role" json:"vaultRole,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			CaCert     *string `tfsdk:"ca_cert" json:"caCert,omitempty"`
			CaPath     *string `tfsdk:"ca_path" json:"caPath,omitempty"`
			ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
			ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
			Insecure   *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
			PathPrefix *string `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
			RootKey    *string `tfsdk:"root_key" json:"rootKey,omitempty"`
			TlsConfig  *struct {
				CaCert        *string `tfsdk:"ca_cert" json:"caCert,omitempty"`
				CaPath        *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				ClientCert    *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
				ClientKey     *string `tfsdk:"client_key" json:"clientKey,omitempty"`
				Insecure      *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
				TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
			Token         *string `tfsdk:"token" json:"token,omitempty"`
		} `tfsdk:"vault_secret_source" json:"vaultSecretSource,omitempty"`
		WatchNamespaces *[]string `tfsdk:"watch_namespaces" json:"watchNamespaces,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GlooSoloIoSettingsV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gloo_solo_io_settings_v1"
}

func (r *GlooSoloIoSettingsV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
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
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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
					"caching_server": schema.SingleNestedAttribute{
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
											Optional:            true,
											Computed:            false,
										},

										"ignore_case": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

										"suffix": schema.StringAttribute{
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

							"caching_service_ref": schema.SingleNestedAttribute{
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

							"max_payload_size": schema.Int64Attribute{
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

							"timeout": schema.StringAttribute{
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

					"console_options": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_explorer_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_only": schema.BoolAttribute{
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

					"consul": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"datacenter": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_polling_interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_discovery": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"data_centers": schema.ListAttribute{
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

							"token": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"username": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_time": schema.StringAttribute{
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

					"consul_discovery": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"consistency_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eds_blocking_queries": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
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

							"root_ca": schema.SingleNestedAttribute{
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

							"service_tags_allowlist": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"split_tls_services": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_tag_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_tls_tagging": schema.BoolAttribute{
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

					"consul_kv_artifact_source": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"root_key": schema.StringAttribute{
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

					"consul_kv_source": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"root_key": schema.StringAttribute{
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

					"dev_mode": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"directory_artifact_source": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"directory": schema.StringAttribute{
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

					"directory_config_source": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"directory": schema.StringAttribute{
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

					"directory_secret_source": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"directory": schema.StringAttribute{
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

					"discovery": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"fds_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fds_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"graphql_enabled": schema.BoolAttribute{
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

							"uds_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"watch_labels": schema.MapAttribute{
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

					"discovery_namespace": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ext_proc": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"allow_mode_override": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"async_mode": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_clear_route_cache": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"failure_mode_allow": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"filter_metadata": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"filter_stage": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"predicate": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stage": schema.StringAttribute{
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
															Optional:            true,
															Computed:            false,
														},

														"ignore_case": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
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

														"suffix": schema.StringAttribute{
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
															Optional:            true,
															Computed:            false,
														},

														"ignore_case": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
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

														"suffix": schema.StringAttribute{
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

							"grpc_service": schema.SingleNestedAttribute{
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

									"ext_proc_server_ref": schema.SingleNestedAttribute{
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

									"initial_metadata": schema.ListNestedAttribute{
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

									"retry_policy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"num_retries": schema.Int64Attribute{
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

											"retry_back_off": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"base_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_interval": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_message_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"message_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_context_namespaces": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mutation_rules": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allow_all_routing": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_envoy": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
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

									"disallow_all": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
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

									"disallow_is_error": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disallow_system": schema.BoolAttribute{
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

							"processing_mode": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"request_body_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_header_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_trailer_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_body_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_header_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_trailer_mode": schema.StringAttribute{
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

							"request_attributes": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"response_attributes": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stat_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"typed_metadata_context_namespaces": schema.ListAttribute{
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

					"extauth": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"clear_route_cache": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extauthz_server_ref": schema.SingleNestedAttribute{
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

							"failure_mode_allow": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grpc_service": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_service": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"path_prefix": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
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
												Optional:            true,
												Computed:            false,
											},

											"allowed_headers_regex": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"headers_to_add": schema.MapAttribute{
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

									"response": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"allowed_client_headers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allowed_upstream_headers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allowed_upstream_headers_to_append": schema.ListAttribute{
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

							"request_body": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allow_partial_message": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_request_bytes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pack_as_bytes": schema.BoolAttribute{
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

							"request_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stat_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"status_on_error": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"transport_api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_id_header": schema.StringAttribute{
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

					"extensions": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"configs": schema.MapAttribute{
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

					"gateway": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"always_sort_route_table_routes": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compressed_proxy_spec": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_gateway_controller": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"isolate_virtual_hosts_by_ssl_config": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"persist_proxy_spec": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_gateways_from_all_namespaces": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"validation": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allow_warnings": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"always_accept": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_transformation_validation": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ignore_gloo_validation_failure": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy_validation_server_addr": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server_enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"validation_server_grpc_max_size_bytes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(-2.147483648e+09),
											int64validator.AtMost(2.147483647e+09),
										},
									},

									"validation_webhook_tls_cert": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"validation_webhook_tls_key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"warn_route_short_circuiting": schema.BoolAttribute{
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

							"validation_server_addr": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"virtual_service_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"one_way_tls": schema.BoolAttribute{
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

					"gloo": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"aws_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"credential_refresh_delay": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_credentials_discovey": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fallback_to_first_function": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"propagate_original_routing": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account_credentials": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"cluster": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uri": schema.StringAttribute{
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

							"disable_grpc_web": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_kubernetes_destinations": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_proxy_garbage_collection": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_rest_eds": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoints_warming_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"failover_upstream_dns_polling_interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"invalid_config_policy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"invalid_route_response_body": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"invalid_route_response_code": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replace_invalid_routes": schema.BoolAttribute{
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

							"log_transformation_request_response_info": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_debug_bind_addr": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"regex_max_program_size": schema.Int64Attribute{
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

							"remove_unused_filters": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rest_xds_bind_addr": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"transformation_escape_characters": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"validation_bind_addr": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xds_bind_addr": schema.StringAttribute{
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

					"graphql_options": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"schema_change_validation_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"processing_rules": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reject_breaking_changes": schema.BoolAttribute{
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

					"knative": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cluster_ingress_proxy_address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"knative_external_proxy_address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"knative_internal_proxy_address": schema.StringAttribute{
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

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"rate_limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"qps": schema.Float64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"burst": schema.Int64Attribute{
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

					"kubernetes_artifact_source": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubernetes_config_source": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubernetes_secret_source": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"linkerd": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"named_extauth": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"clear_route_cache": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extauthz_server_ref": schema.SingleNestedAttribute{
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

							"failure_mode_allow": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grpc_service": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_service": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"path_prefix": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
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
												Optional:            true,
												Computed:            false,
											},

											"allowed_headers_regex": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"headers_to_add": schema.MapAttribute{
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

									"response": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"allowed_client_headers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allowed_upstream_headers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allowed_upstream_headers_to_append": schema.ListAttribute{
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

							"request_body": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allow_partial_message": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_request_bytes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pack_as_bytes": schema.BoolAttribute{
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

							"request_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stat_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"status_on_error": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"transport_api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_id_header": schema.StringAttribute{
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

					"observability_options": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"config_status_metric_labels": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"label_to_path": schema.MapAttribute{
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

							"grafana_integration": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"default_dashboard_folder_id": schema.Int64Attribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ratelimit": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"descriptors": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"set_descriptors": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"always_apply": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rate_limit": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"requests_per_unit": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"unit": schema.StringAttribute{
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

										"simple_descriptors": schema.ListNestedAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ratelimit_server": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"deny_on_fail": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_x_ratelimit_headers": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate_limit_before_auth": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ratelimit_server_ref": schema.SingleNestedAttribute{
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

							"request_timeout": schema.StringAttribute{
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

					"rbac": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"require_rbac": schema.BoolAttribute{
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

					"refresh_rate": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_options": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"sources": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"directory": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"directory": schema.StringAttribute{
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

										"kubernetes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"vault": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_token": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"address": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"aws": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"access_key_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"iam_server_id_header": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mount_path": schema.StringAttribute{
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

														"secret_access_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"session_token": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"vault_role": schema.StringAttribute{
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

												"ca_cert": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_cert": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_key": schema.StringAttribute{
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

												"path_prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"root_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls_config": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"ca_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ca_path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
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

														"tls_server_name": schema.StringAttribute{
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

												"tls_server_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"token": schema.StringAttribute{
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

					"upstream_options": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"global_annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_parameters": schema.SingleNestedAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vault_secret_source": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"access_token": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"iam_server_id_header": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mount_path": schema.StringAttribute{
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

									"secret_access_key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"session_token": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vault_role": schema.StringAttribute{
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

							"ca_cert": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_cert": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_key": schema.StringAttribute{
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

							"path_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"root_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_cert": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_key": schema.StringAttribute{
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

									"tls_server_name": schema.StringAttribute{
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

							"tls_server_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token": schema.StringAttribute{
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

					"watch_namespaces": schema.ListAttribute{
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
	}
}

func (r *GlooSoloIoSettingsV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *GlooSoloIoSettingsV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gloo_solo_io_settings_v1")

	var model GlooSoloIoSettingsV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("gloo.solo.io/v1")
	model.Kind = pointer.String("Settings")

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
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "settings"}).
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

	var readResponse GlooSoloIoSettingsV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *GlooSoloIoSettingsV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_settings_v1")

	var data GlooSoloIoSettingsV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "settings"}).
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

	var readResponse GlooSoloIoSettingsV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *GlooSoloIoSettingsV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gloo_solo_io_settings_v1")

	var model GlooSoloIoSettingsV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gloo.solo.io/v1")
	model.Kind = pointer.String("Settings")

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
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "settings"}).
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

	var readResponse GlooSoloIoSettingsV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *GlooSoloIoSettingsV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gloo_solo_io_settings_v1")

	var data GlooSoloIoSettingsV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "settings"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "settings"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
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

func (r *GlooSoloIoSettingsV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
