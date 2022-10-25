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

type GlooSoloIoSettingsV1Resource struct{}

var (
	_ resource.Resource = (*GlooSoloIoSettingsV1Resource)(nil)
)

type GlooSoloIoSettingsV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GlooSoloIoSettingsV1GoModel struct {
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
		CachingServer *struct {
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
		} `tfsdk:"caching_server" yaml:"cachingServer,omitempty"`

		ConsoleOptions *struct {
			ApiExplorerEnabled *bool `tfsdk:"api_explorer_enabled" yaml:"apiExplorerEnabled,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
		} `tfsdk:"console_options" yaml:"consoleOptions,omitempty"`

		Consul *struct {
			Address *string `tfsdk:"address" yaml:"address,omitempty"`

			CaFile *string `tfsdk:"ca_file" yaml:"caFile,omitempty"`

			CaPath *string `tfsdk:"ca_path" yaml:"caPath,omitempty"`

			CertFile *string `tfsdk:"cert_file" yaml:"certFile,omitempty"`

			Datacenter *string `tfsdk:"datacenter" yaml:"datacenter,omitempty"`

			DnsAddress *string `tfsdk:"dns_address" yaml:"dnsAddress,omitempty"`

			DnsPollingInterval *string `tfsdk:"dns_polling_interval" yaml:"dnsPollingInterval,omitempty"`

			HttpAddress *string `tfsdk:"http_address" yaml:"httpAddress,omitempty"`

			InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

			KeyFile *string `tfsdk:"key_file" yaml:"keyFile,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			ServiceDiscovery *struct {
				DataCenters *[]string `tfsdk:"data_centers" yaml:"dataCenters,omitempty"`
			} `tfsdk:"service_discovery" yaml:"serviceDiscovery,omitempty"`

			Token *string `tfsdk:"token" yaml:"token,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`

			WaitTime *string `tfsdk:"wait_time" yaml:"waitTime,omitempty"`
		} `tfsdk:"consul" yaml:"consul,omitempty"`

		ConsulDiscovery *struct {
			ConsistencyMode utilities.IntOrString `tfsdk:"consistency_mode" yaml:"consistencyMode,omitempty"`

			EdsBlockingQueries *bool `tfsdk:"eds_blocking_queries" yaml:"edsBlockingQueries,omitempty"`

			QueryOptions *struct {
				UseCache *bool `tfsdk:"use_cache" yaml:"useCache,omitempty"`
			} `tfsdk:"query_options" yaml:"queryOptions,omitempty"`

			RootCa *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"root_ca" yaml:"rootCa,omitempty"`

			ServiceTagsAllowlist *[]string `tfsdk:"service_tags_allowlist" yaml:"serviceTagsAllowlist,omitempty"`

			SplitTlsServices *bool `tfsdk:"split_tls_services" yaml:"splitTlsServices,omitempty"`

			TlsTagName *string `tfsdk:"tls_tag_name" yaml:"tlsTagName,omitempty"`

			UseTlsTagging *bool `tfsdk:"use_tls_tagging" yaml:"useTlsTagging,omitempty"`
		} `tfsdk:"consul_discovery" yaml:"consulDiscovery,omitempty"`

		ConsulKvArtifactSource *struct {
			RootKey *string `tfsdk:"root_key" yaml:"rootKey,omitempty"`
		} `tfsdk:"consul_kv_artifact_source" yaml:"consulKvArtifactSource,omitempty"`

		ConsulKvSource *struct {
			RootKey *string `tfsdk:"root_key" yaml:"rootKey,omitempty"`
		} `tfsdk:"consul_kv_source" yaml:"consulKvSource,omitempty"`

		DevMode *bool `tfsdk:"dev_mode" yaml:"devMode,omitempty"`

		DirectoryArtifactSource *struct {
			Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`
		} `tfsdk:"directory_artifact_source" yaml:"directoryArtifactSource,omitempty"`

		DirectoryConfigSource *struct {
			Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`
		} `tfsdk:"directory_config_source" yaml:"directoryConfigSource,omitempty"`

		DirectorySecretSource *struct {
			Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`
		} `tfsdk:"directory_secret_source" yaml:"directorySecretSource,omitempty"`

		Discovery *struct {
			FdsMode utilities.IntOrString `tfsdk:"fds_mode" yaml:"fdsMode,omitempty"`

			FdsOptions *struct {
				GraphqlEnabled *bool `tfsdk:"graphql_enabled" yaml:"graphqlEnabled,omitempty"`
			} `tfsdk:"fds_options" yaml:"fdsOptions,omitempty"`

			UdsOptions *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				WatchLabels *map[string]string `tfsdk:"watch_labels" yaml:"watchLabels,omitempty"`
			} `tfsdk:"uds_options" yaml:"udsOptions,omitempty"`
		} `tfsdk:"discovery" yaml:"discovery,omitempty"`

		DiscoveryNamespace *string `tfsdk:"discovery_namespace" yaml:"discoveryNamespace,omitempty"`

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

		Gateway *struct {
			AlwaysSortRouteTableRoutes *bool `tfsdk:"always_sort_route_table_routes" yaml:"alwaysSortRouteTableRoutes,omitempty"`

			CompressedProxySpec *bool `tfsdk:"compressed_proxy_spec" yaml:"compressedProxySpec,omitempty"`

			EnableGatewayController *bool `tfsdk:"enable_gateway_controller" yaml:"enableGatewayController,omitempty"`

			IsolateVirtualHostsBySslConfig *bool `tfsdk:"isolate_virtual_hosts_by_ssl_config" yaml:"isolateVirtualHostsBySslConfig,omitempty"`

			PersistProxySpec *bool `tfsdk:"persist_proxy_spec" yaml:"persistProxySpec,omitempty"`

			ReadGatewaysFromAllNamespaces *bool `tfsdk:"read_gateways_from_all_namespaces" yaml:"readGatewaysFromAllNamespaces,omitempty"`

			Validation *struct {
				AllowWarnings *bool `tfsdk:"allow_warnings" yaml:"allowWarnings,omitempty"`

				AlwaysAccept *bool `tfsdk:"always_accept" yaml:"alwaysAccept,omitempty"`

				DisableTransformationValidation *bool `tfsdk:"disable_transformation_validation" yaml:"disableTransformationValidation,omitempty"`

				IgnoreGlooValidationFailure *bool `tfsdk:"ignore_gloo_validation_failure" yaml:"ignoreGlooValidationFailure,omitempty"`

				ProxyValidationServerAddr *string `tfsdk:"proxy_validation_server_addr" yaml:"proxyValidationServerAddr,omitempty"`

				ServerEnabled *bool `tfsdk:"server_enabled" yaml:"serverEnabled,omitempty"`

				ValidationServerGrpcMaxSizeBytes *int64 `tfsdk:"validation_server_grpc_max_size_bytes" yaml:"validationServerGrpcMaxSizeBytes,omitempty"`

				ValidationWebhookTlsCert *string `tfsdk:"validation_webhook_tls_cert" yaml:"validationWebhookTlsCert,omitempty"`

				ValidationWebhookTlsKey *string `tfsdk:"validation_webhook_tls_key" yaml:"validationWebhookTlsKey,omitempty"`

				WarnRouteShortCircuiting *bool `tfsdk:"warn_route_short_circuiting" yaml:"warnRouteShortCircuiting,omitempty"`
			} `tfsdk:"validation" yaml:"validation,omitempty"`

			ValidationServerAddr *string `tfsdk:"validation_server_addr" yaml:"validationServerAddr,omitempty"`

			VirtualServiceOptions *struct {
				OneWayTls *bool `tfsdk:"one_way_tls" yaml:"oneWayTls,omitempty"`
			} `tfsdk:"virtual_service_options" yaml:"virtualServiceOptions,omitempty"`
		} `tfsdk:"gateway" yaml:"gateway,omitempty"`

		Gloo *struct {
			AwsOptions *struct {
				CredentialRefreshDelay *string `tfsdk:"credential_refresh_delay" yaml:"credentialRefreshDelay,omitempty"`

				EnableCredentialsDiscovey *bool `tfsdk:"enable_credentials_discovey" yaml:"enableCredentialsDiscovey,omitempty"`

				PropagateOriginalRouting *bool `tfsdk:"propagate_original_routing" yaml:"propagateOriginalRouting,omitempty"`

				ServiceAccountCredentials *struct {
					Cluster *string `tfsdk:"cluster" yaml:"cluster,omitempty"`

					Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
				} `tfsdk:"service_account_credentials" yaml:"serviceAccountCredentials,omitempty"`
			} `tfsdk:"aws_options" yaml:"awsOptions,omitempty"`

			CircuitBreakers *struct {
				MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

				MaxPendingRequests *int64 `tfsdk:"max_pending_requests" yaml:"maxPendingRequests,omitempty"`

				MaxRequests *int64 `tfsdk:"max_requests" yaml:"maxRequests,omitempty"`

				MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`
			} `tfsdk:"circuit_breakers" yaml:"circuitBreakers,omitempty"`

			DisableGrpcWeb *bool `tfsdk:"disable_grpc_web" yaml:"disableGrpcWeb,omitempty"`

			DisableKubernetesDestinations *bool `tfsdk:"disable_kubernetes_destinations" yaml:"disableKubernetesDestinations,omitempty"`

			DisableProxyGarbageCollection *bool `tfsdk:"disable_proxy_garbage_collection" yaml:"disableProxyGarbageCollection,omitempty"`

			EnableRestEds *bool `tfsdk:"enable_rest_eds" yaml:"enableRestEds,omitempty"`

			EndpointsWarmingTimeout *string `tfsdk:"endpoints_warming_timeout" yaml:"endpointsWarmingTimeout,omitempty"`

			FailoverUpstreamDnsPollingInterval *string `tfsdk:"failover_upstream_dns_polling_interval" yaml:"failoverUpstreamDnsPollingInterval,omitempty"`

			InvalidConfigPolicy *struct {
				InvalidRouteResponseBody *string `tfsdk:"invalid_route_response_body" yaml:"invalidRouteResponseBody,omitempty"`

				InvalidRouteResponseCode *int64 `tfsdk:"invalid_route_response_code" yaml:"invalidRouteResponseCode,omitempty"`

				ReplaceInvalidRoutes *bool `tfsdk:"replace_invalid_routes" yaml:"replaceInvalidRoutes,omitempty"`
			} `tfsdk:"invalid_config_policy" yaml:"invalidConfigPolicy,omitempty"`

			ProxyDebugBindAddr *string `tfsdk:"proxy_debug_bind_addr" yaml:"proxyDebugBindAddr,omitempty"`

			RegexMaxProgramSize *int64 `tfsdk:"regex_max_program_size" yaml:"regexMaxProgramSize,omitempty"`

			RemoveUnusedFilters *bool `tfsdk:"remove_unused_filters" yaml:"removeUnusedFilters,omitempty"`

			RestXdsBindAddr *string `tfsdk:"rest_xds_bind_addr" yaml:"restXdsBindAddr,omitempty"`

			ValidationBindAddr *string `tfsdk:"validation_bind_addr" yaml:"validationBindAddr,omitempty"`

			XdsBindAddr *string `tfsdk:"xds_bind_addr" yaml:"xdsBindAddr,omitempty"`
		} `tfsdk:"gloo" yaml:"gloo,omitempty"`

		GraphqlOptions *struct {
			SchemaChangeValidationOptions *struct {
				ProcessingRules *[]string `tfsdk:"processing_rules" yaml:"processingRules,omitempty"`

				RejectBreakingChanges *bool `tfsdk:"reject_breaking_changes" yaml:"rejectBreakingChanges,omitempty"`
			} `tfsdk:"schema_change_validation_options" yaml:"schemaChangeValidationOptions,omitempty"`
		} `tfsdk:"graphql_options" yaml:"graphqlOptions,omitempty"`

		Knative *struct {
			ClusterIngressProxyAddress *string `tfsdk:"cluster_ingress_proxy_address" yaml:"clusterIngressProxyAddress,omitempty"`

			KnativeExternalProxyAddress *string `tfsdk:"knative_external_proxy_address" yaml:"knativeExternalProxyAddress,omitempty"`

			KnativeInternalProxyAddress *string `tfsdk:"knative_internal_proxy_address" yaml:"knativeInternalProxyAddress,omitempty"`
		} `tfsdk:"knative" yaml:"knative,omitempty"`

		Kubernetes *struct {
			RateLimits *struct {
				QPS utilities.DynamicNumber `tfsdk:"qps" yaml:"QPS,omitempty"`

				Burst *int64 `tfsdk:"burst" yaml:"burst,omitempty"`
			} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
		} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

		KubernetesArtifactSource *map[string]string `tfsdk:"kubernetes_artifact_source" yaml:"kubernetesArtifactSource,omitempty"`

		KubernetesConfigSource *map[string]string `tfsdk:"kubernetes_config_source" yaml:"kubernetesConfigSource,omitempty"`

		KubernetesSecretSource *map[string]string `tfsdk:"kubernetes_secret_source" yaml:"kubernetesSecretSource,omitempty"`

		Linkerd *bool `tfsdk:"linkerd" yaml:"linkerd,omitempty"`

		NamedExtauth *struct {
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
		} `tfsdk:"named_extauth" yaml:"namedExtauth,omitempty"`

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`

		ObservabilityOptions *struct {
			ConfigStatusMetricLabels *struct {
				LabelToPath *map[string]string `tfsdk:"label_to_path" yaml:"labelToPath,omitempty"`
			} `tfsdk:"config_status_metric_labels" yaml:"configStatusMetricLabels,omitempty"`

			GrafanaIntegration *struct {
				DefaultDashboardFolderId *int64 `tfsdk:"default_dashboard_folder_id" yaml:"defaultDashboardFolderId,omitempty"`
			} `tfsdk:"grafana_integration" yaml:"grafanaIntegration,omitempty"`
		} `tfsdk:"observability_options" yaml:"observabilityOptions,omitempty"`

		Ratelimit *struct {
			Descriptors *[]map[string]string `tfsdk:"descriptors" yaml:"descriptors,omitempty"`

			SetDescriptors *[]struct {
				AlwaysApply *bool `tfsdk:"always_apply" yaml:"alwaysApply,omitempty"`

				RateLimit *struct {
					RequestsPerUnit *int64 `tfsdk:"requests_per_unit" yaml:"requestsPerUnit,omitempty"`

					Unit utilities.IntOrString `tfsdk:"unit" yaml:"unit,omitempty"`
				} `tfsdk:"rate_limit" yaml:"rateLimit,omitempty"`

				SimpleDescriptors *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"simple_descriptors" yaml:"simpleDescriptors,omitempty"`
			} `tfsdk:"set_descriptors" yaml:"setDescriptors,omitempty"`
		} `tfsdk:"ratelimit" yaml:"ratelimit,omitempty"`

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

		Rbac *struct {
			RequireRbac *bool `tfsdk:"require_rbac" yaml:"requireRbac,omitempty"`
		} `tfsdk:"rbac" yaml:"rbac,omitempty"`

		RefreshRate *string `tfsdk:"refresh_rate" yaml:"refreshRate,omitempty"`

		UpstreamOptions *struct {
			GlobalAnnotations *map[string]string `tfsdk:"global_annotations" yaml:"globalAnnotations,omitempty"`

			SslParameters *struct {
				CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

				EcdhCurves *[]string `tfsdk:"ecdh_curves" yaml:"ecdhCurves,omitempty"`

				MaximumProtocolVersion utilities.IntOrString `tfsdk:"maximum_protocol_version" yaml:"maximumProtocolVersion,omitempty"`

				MinimumProtocolVersion utilities.IntOrString `tfsdk:"minimum_protocol_version" yaml:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"ssl_parameters" yaml:"sslParameters,omitempty"`
		} `tfsdk:"upstream_options" yaml:"upstreamOptions,omitempty"`

		VaultSecretSource *struct {
			Address *string `tfsdk:"address" yaml:"address,omitempty"`

			CaCert *string `tfsdk:"ca_cert" yaml:"caCert,omitempty"`

			CaPath *string `tfsdk:"ca_path" yaml:"caPath,omitempty"`

			ClientCert *string `tfsdk:"client_cert" yaml:"clientCert,omitempty"`

			ClientKey *string `tfsdk:"client_key" yaml:"clientKey,omitempty"`

			Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

			PathPrefix *string `tfsdk:"path_prefix" yaml:"pathPrefix,omitempty"`

			RootKey *string `tfsdk:"root_key" yaml:"rootKey,omitempty"`

			TlsServerName *string `tfsdk:"tls_server_name" yaml:"tlsServerName,omitempty"`

			Token *string `tfsdk:"token" yaml:"token,omitempty"`
		} `tfsdk:"vault_secret_source" yaml:"vaultSecretSource,omitempty"`

		WatchNamespaces *[]string `tfsdk:"watch_namespaces" yaml:"watchNamespaces,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGlooSoloIoSettingsV1Resource() resource.Resource {
	return &GlooSoloIoSettingsV1Resource{}
}

func (r *GlooSoloIoSettingsV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gloo_solo_io_settings_v1"
}

func (r *GlooSoloIoSettingsV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

					"caching_server": {
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

					"console_options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_explorer_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_only": {
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

					"consul": {
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

							"ca_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ca_path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cert_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"datacenter": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns_polling_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"insecure_skip_verify": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_discovery": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"data_centers": {
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

							"token": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wait_time": {
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

					"consul_discovery": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"consistency_mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"eds_blocking_queries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"query_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"use_cache": {
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

							"root_ca": {
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

							"service_tags_allowlist": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"split_tls_services": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_tag_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_tls_tagging": {
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

					"consul_kv_artifact_source": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"root_key": {
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

					"consul_kv_source": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"root_key": {
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

					"dev_mode": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"directory_artifact_source": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"directory": {
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

					"directory_config_source": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"directory": {
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

					"directory_secret_source": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"directory": {
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

					"discovery": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fds_mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fds_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"graphql_enabled": {
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

							"uds_options": {
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

									"watch_labels": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"discovery_namespace": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

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

					"gateway": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"always_sort_route_table_routes": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"compressed_proxy_spec": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_gateway_controller": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"isolate_virtual_hosts_by_ssl_config": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"persist_proxy_spec": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_gateways_from_all_namespaces": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"validation": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow_warnings": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"always_accept": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_transformation_validation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ignore_gloo_validation_failure": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"proxy_validation_server_addr": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"server_enabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"validation_server_grpc_max_size_bytes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(-2.147483648e+09),

											int64validator.AtMost(2.147483647e+09),
										},
									},

									"validation_webhook_tls_cert": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"validation_webhook_tls_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"warn_route_short_circuiting": {
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

							"validation_server_addr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"virtual_service_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"one_way_tls": {
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

					"gloo": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"aws_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credential_refresh_delay": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable_credentials_discovey": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"propagate_original_routing": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_account_credentials": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cluster": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"uri": {
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

							"circuit_breakers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_connections": {
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

									"max_requests": {
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

									"max_retries": {
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

							"disable_grpc_web": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_kubernetes_destinations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_proxy_garbage_collection": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_rest_eds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"endpoints_warming_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"failover_upstream_dns_polling_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"invalid_config_policy": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"invalid_route_response_body": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"invalid_route_response_code": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replace_invalid_routes": {
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

							"proxy_debug_bind_addr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"regex_max_program_size": {
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

							"remove_unused_filters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rest_xds_bind_addr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"validation_bind_addr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"xds_bind_addr": {
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

					"graphql_options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"schema_change_validation_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"processing_rules": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reject_breaking_changes": {
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

					"knative": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cluster_ingress_proxy_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"knative_external_proxy_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"knative_internal_proxy_address": {
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

					"kubernetes": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rate_limits": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"qps": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"burst": {
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

					"kubernetes_artifact_source": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes_config_source": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes_secret_source": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"linkerd": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"named_extauth": {
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

					"observability_options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"config_status_metric_labels": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"label_to_path": {
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

							"grafana_integration": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_dashboard_folder_id": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ratelimit": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"descriptors": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"set_descriptors": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"always_apply": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rate_limit": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"requests_per_unit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unit": {
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

									"simple_descriptors": {
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

					"rbac": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"require_rbac": {
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

					"refresh_rate": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"upstream_options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"global_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_parameters": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vault_secret_source": {
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

							"ca_cert": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ca_path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_cert": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"insecure": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path_prefix": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"root_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_server_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"token": {
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

					"watch_namespaces": {
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
		},
	}, nil
}

func (r *GlooSoloIoSettingsV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gloo_solo_io_settings_v1")

	var state GlooSoloIoSettingsV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoSettingsV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("Settings")

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

func (r *GlooSoloIoSettingsV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_settings_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GlooSoloIoSettingsV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gloo_solo_io_settings_v1")

	var state GlooSoloIoSettingsV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoSettingsV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("Settings")

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

func (r *GlooSoloIoSettingsV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gloo_solo_io_settings_v1")
	// NO-OP: Terraform removes the state automatically for us
}
