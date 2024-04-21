/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package servicemesh_cisco_com_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ServicemeshCiscoComIstioMeshV1Alpha1Manifest{}
)

func NewServicemeshCiscoComIstioMeshV1Alpha1Manifest() datasource.DataSource {
	return &ServicemeshCiscoComIstioMeshV1Alpha1Manifest{}
}

type ServicemeshCiscoComIstioMeshV1Alpha1Manifest struct{}

type ServicemeshCiscoComIstioMeshV1Alpha1ManifestData struct {
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
		Config *struct {
			AccessLogEncoding *string `tfsdk:"access_log_encoding" json:"accessLogEncoding,omitempty"`
			AccessLogFile     *string `tfsdk:"access_log_file" json:"accessLogFile,omitempty"`
			AccessLogFormat   *string `tfsdk:"access_log_format" json:"accessLogFormat,omitempty"`
			Ca                *struct {
				Address        *string `tfsdk:"address" json:"address,omitempty"`
				IstiodSide     *bool   `tfsdk:"istiod_side" json:"istiodSide,omitempty"`
				RequestTimeout *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
				TlsSettings    *struct {
					CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
					ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
					CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
					InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
					PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
					Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
					SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"tls_settings" json:"tlsSettings,omitempty"`
			} `tfsdk:"ca" json:"ca,omitempty"`
			CaCertificates *[]struct {
				CertSigners     *[]string `tfsdk:"cert_signers" json:"certSigners,omitempty"`
				Pem             *string   `tfsdk:"pem" json:"pem,omitempty"`
				SpiffeBundleUrl *string   `tfsdk:"spiffe_bundle_url" json:"spiffeBundleUrl,omitempty"`
				TrustDomains    *[]string `tfsdk:"trust_domains" json:"trustDomains,omitempty"`
			} `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
			Certificates *[]struct {
				DnsNames   *[]string `tfsdk:"dns_names" json:"dnsNames,omitempty"`
				SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"certificates" json:"certificates,omitempty"`
			ConfigSources *[]struct {
				Address             *string   `tfsdk:"address" json:"address,omitempty"`
				SubscribedResources *[]string `tfsdk:"subscribed_resources" json:"subscribedResources,omitempty"`
				TlsSettings         *struct {
					CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
					ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
					CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
					InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
					PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
					Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
					SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"tls_settings" json:"tlsSettings,omitempty"`
			} `tfsdk:"config_sources" json:"configSources,omitempty"`
			ConnectTimeout *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
			DefaultConfig  *struct {
				AvailabilityZone       *string   `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
				BinaryPath             *string   `tfsdk:"binary_path" json:"binaryPath,omitempty"`
				CaCertificatesPem      *[]string `tfsdk:"ca_certificates_pem" json:"caCertificatesPem,omitempty"`
				Concurrency            *int64    `tfsdk:"concurrency" json:"concurrency,omitempty"`
				ConfigPath             *string   `tfsdk:"config_path" json:"configPath,omitempty"`
				ControlPlaneAuthPolicy *string   `tfsdk:"control_plane_auth_policy" json:"controlPlaneAuthPolicy,omitempty"`
				CustomConfigFile       *string   `tfsdk:"custom_config_file" json:"customConfigFile,omitempty"`
				DiscoveryAddress       *string   `tfsdk:"discovery_address" json:"discoveryAddress,omitempty"`
				DiscoveryRefreshDelay  *string   `tfsdk:"discovery_refresh_delay" json:"discoveryRefreshDelay,omitempty"`
				DrainDuration          *string   `tfsdk:"drain_duration" json:"drainDuration,omitempty"`
				EnvoyAccessLogService  *struct {
					Address      *string `tfsdk:"address" json:"address,omitempty"`
					TcpKeepalive *struct {
						Interval *string `tfsdk:"interval" json:"interval,omitempty"`
						Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
						Time     *string `tfsdk:"time" json:"time,omitempty"`
					} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
					TlsSettings *struct {
						CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
						ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
						CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
						InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
						PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
						Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
						SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
					} `tfsdk:"tls_settings" json:"tlsSettings,omitempty"`
				} `tfsdk:"envoy_access_log_service" json:"envoyAccessLogService,omitempty"`
				EnvoyMetricsService *struct {
					Address      *string `tfsdk:"address" json:"address,omitempty"`
					TcpKeepalive *struct {
						Interval *string `tfsdk:"interval" json:"interval,omitempty"`
						Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
						Time     *string `tfsdk:"time" json:"time,omitempty"`
					} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
					TlsSettings *struct {
						CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
						ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
						CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
						InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
						PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
						Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
						SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
					} `tfsdk:"tls_settings" json:"tlsSettings,omitempty"`
				} `tfsdk:"envoy_metrics_service" json:"envoyMetricsService,omitempty"`
				EnvoyMetricsServiceAddress *string   `tfsdk:"envoy_metrics_service_address" json:"envoyMetricsServiceAddress,omitempty"`
				ExtraStatTags              *[]string `tfsdk:"extra_stat_tags" json:"extraStatTags,omitempty"`
				GatewayTopology            *struct {
					ForwardClientCertDetails *string `tfsdk:"forward_client_cert_details" json:"forwardClientCertDetails,omitempty"`
					NumTrustedProxies        *int64  `tfsdk:"num_trusted_proxies" json:"numTrustedProxies,omitempty"`
				} `tfsdk:"gateway_topology" json:"gatewayTopology,omitempty"`
				HoldApplicationUntilProxyStarts *bool `tfsdk:"hold_application_until_proxy_starts" json:"holdApplicationUntilProxyStarts,omitempty"`
				Image                           *struct {
					ImageType *string `tfsdk:"image_type" json:"imageType,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				InterceptionMode   *string `tfsdk:"interception_mode" json:"interceptionMode,omitempty"`
				MeshId             *string `tfsdk:"mesh_id" json:"meshId,omitempty"`
				PrivateKeyProvider *struct {
					Cryptomb *struct {
						PollDelay *string `tfsdk:"poll_delay" json:"pollDelay,omitempty"`
					} `tfsdk:"cryptomb" json:"cryptomb,omitempty"`
					Qat *struct {
						PollDelay *string `tfsdk:"poll_delay" json:"pollDelay,omitempty"`
					} `tfsdk:"qat" json:"qat,omitempty"`
				} `tfsdk:"private_key_provider" json:"privateKeyProvider,omitempty"`
				ProxyAdminPort             *int64             `tfsdk:"proxy_admin_port" json:"proxyAdminPort,omitempty"`
				ProxyBootstrapTemplatePath *string            `tfsdk:"proxy_bootstrap_template_path" json:"proxyBootstrapTemplatePath,omitempty"`
				ProxyMetadata              *map[string]string `tfsdk:"proxy_metadata" json:"proxyMetadata,omitempty"`
				ProxyStatsMatcher          *struct {
					InclusionPrefixes *[]string `tfsdk:"inclusion_prefixes" json:"inclusionPrefixes,omitempty"`
					InclusionRegexps  *[]string `tfsdk:"inclusion_regexps" json:"inclusionRegexps,omitempty"`
					InclusionSuffixes *[]string `tfsdk:"inclusion_suffixes" json:"inclusionSuffixes,omitempty"`
				} `tfsdk:"proxy_stats_matcher" json:"proxyStatsMatcher,omitempty"`
				ReadinessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					HttpGet          *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						HttpHeaders *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *int64  `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TcpSocket           *struct {
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *int64  `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				RuntimeValues *map[string]string `tfsdk:"runtime_values" json:"runtimeValues,omitempty"`
				Sds           *struct {
					Enabled      *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					K8sSaJwtPath *string `tfsdk:"k8s_sa_jwt_path" json:"k8sSaJwtPath,omitempty"`
				} `tfsdk:"sds" json:"sds,omitempty"`
				ServiceCluster           *string `tfsdk:"service_cluster" json:"serviceCluster,omitempty"`
				StatNameLength           *int64  `tfsdk:"stat_name_length" json:"statNameLength,omitempty"`
				StatsdUdpAddress         *string `tfsdk:"statsd_udp_address" json:"statsdUdpAddress,omitempty"`
				StatusPort               *int64  `tfsdk:"status_port" json:"statusPort,omitempty"`
				TerminationDrainDuration *string `tfsdk:"termination_drain_duration" json:"terminationDrainDuration,omitempty"`
				Tracing                  *struct {
					CustomTags *struct {
						Environment *struct {
							DefaultValue *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							Name         *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"environment" json:"environment,omitempty"`
						Header *struct {
							DefaultValue *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							Name         *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"header" json:"header,omitempty"`
						Literal *struct {
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"literal" json:"literal,omitempty"`
					} `tfsdk:"custom_tags" json:"customTags,omitempty"`
					Datadog *struct {
						Address *string `tfsdk:"address" json:"address,omitempty"`
					} `tfsdk:"datadog" json:"datadog,omitempty"`
					Lightstep *struct {
						AccessToken *string `tfsdk:"access_token" json:"accessToken,omitempty"`
						Address     *string `tfsdk:"address" json:"address,omitempty"`
					} `tfsdk:"lightstep" json:"lightstep,omitempty"`
					MaxPathTagLength *int64 `tfsdk:"max_path_tag_length" json:"maxPathTagLength,omitempty"`
					OpenCensusAgent  *struct {
						Address *string   `tfsdk:"address" json:"address,omitempty"`
						Context *[]string `tfsdk:"context" json:"context,omitempty"`
					} `tfsdk:"open_census_agent" json:"openCensusAgent,omitempty"`
					Sampling    *float64 `tfsdk:"sampling" json:"sampling,omitempty"`
					Stackdriver *struct {
						Debug                    *bool  `tfsdk:"debug" json:"debug,omitempty"`
						MaxNumberOfAnnotations   *int64 `tfsdk:"max_number_of_annotations" json:"maxNumberOfAnnotations,omitempty"`
						MaxNumberOfAttributes    *int64 `tfsdk:"max_number_of_attributes" json:"maxNumberOfAttributes,omitempty"`
						MaxNumberOfMessageEvents *int64 `tfsdk:"max_number_of_message_events" json:"maxNumberOfMessageEvents,omitempty"`
					} `tfsdk:"stackdriver" json:"stackdriver,omitempty"`
					TlsSettings *struct {
						CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
						ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
						CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
						InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
						PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
						Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
						SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
					} `tfsdk:"tls_settings" json:"tlsSettings,omitempty"`
					Zipkin *struct {
						Address *string `tfsdk:"address" json:"address,omitempty"`
					} `tfsdk:"zipkin" json:"zipkin,omitempty"`
				} `tfsdk:"tracing" json:"tracing,omitempty"`
				TracingServiceName *string `tfsdk:"tracing_service_name" json:"tracingServiceName,omitempty"`
				ZipkinAddress      *string `tfsdk:"zipkin_address" json:"zipkinAddress,omitempty"`
			} `tfsdk:"default_config" json:"defaultConfig,omitempty"`
			DefaultDestinationRuleExportTo *[]string `tfsdk:"default_destination_rule_export_to" json:"defaultDestinationRuleExportTo,omitempty"`
			DefaultHttpRetryPolicy         *struct {
				Attempts              *int64  `tfsdk:"attempts" json:"attempts,omitempty"`
				PerTryTimeout         *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
				RetryOn               *string `tfsdk:"retry_on" json:"retryOn,omitempty"`
				RetryRemoteLocalities *bool   `tfsdk:"retry_remote_localities" json:"retryRemoteLocalities,omitempty"`
			} `tfsdk:"default_http_retry_policy" json:"defaultHttpRetryPolicy,omitempty"`
			DefaultProviders *struct {
				AccessLogging *[]string `tfsdk:"access_logging" json:"accessLogging,omitempty"`
				Metrics       *[]string `tfsdk:"metrics" json:"metrics,omitempty"`
				Tracing       *[]string `tfsdk:"tracing" json:"tracing,omitempty"`
			} `tfsdk:"default_providers" json:"defaultProviders,omitempty"`
			DefaultServiceExportTo        *[]string `tfsdk:"default_service_export_to" json:"defaultServiceExportTo,omitempty"`
			DefaultVirtualServiceExportTo *[]string `tfsdk:"default_virtual_service_export_to" json:"defaultVirtualServiceExportTo,omitempty"`
			DisableEnvoyListenerLog       *bool     `tfsdk:"disable_envoy_listener_log" json:"disableEnvoyListenerLog,omitempty"`
			DiscoverySelectors            *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"discovery_selectors" json:"discoverySelectors,omitempty"`
			DnsRefreshRate              *string `tfsdk:"dns_refresh_rate" json:"dnsRefreshRate,omitempty"`
			EnableAutoMtls              *bool   `tfsdk:"enable_auto_mtls" json:"enableAutoMtls,omitempty"`
			EnableEnvoyAccessLogService *bool   `tfsdk:"enable_envoy_access_log_service" json:"enableEnvoyAccessLogService,omitempty"`
			EnablePrometheusMerge       *bool   `tfsdk:"enable_prometheus_merge" json:"enablePrometheusMerge,omitempty"`
			EnableTracing               *bool   `tfsdk:"enable_tracing" json:"enableTracing,omitempty"`
			ExtensionProviders          *[]struct {
				Datadog *struct {
					MaxTagLength *int64  `tfsdk:"max_tag_length" json:"maxTagLength,omitempty"`
					Port         *int64  `tfsdk:"port" json:"port,omitempty"`
					Service      *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"datadog" json:"datadog,omitempty"`
				EnvoyExtAuthzGrpc *struct {
					FailOpen                  *bool `tfsdk:"fail_open" json:"failOpen,omitempty"`
					IncludeRequestBodyInCheck *struct {
						AllowPartialMessage *bool  `tfsdk:"allow_partial_message" json:"allowPartialMessage,omitempty"`
						MaxRequestBytes     *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
						PackAsBytes         *bool  `tfsdk:"pack_as_bytes" json:"packAsBytes,omitempty"`
					} `tfsdk:"include_request_body_in_check" json:"includeRequestBodyInCheck,omitempty"`
					Port          *int64  `tfsdk:"port" json:"port,omitempty"`
					Service       *string `tfsdk:"service" json:"service,omitempty"`
					StatusOnError *string `tfsdk:"status_on_error" json:"statusOnError,omitempty"`
					Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"envoy_ext_authz_grpc" json:"envoyExtAuthzGrpc,omitempty"`
				EnvoyExtAuthzHttp *struct {
					FailOpen                        *bool              `tfsdk:"fail_open" json:"failOpen,omitempty"`
					HeadersToDownstreamOnAllow      *[]string          `tfsdk:"headers_to_downstream_on_allow" json:"headersToDownstreamOnAllow,omitempty"`
					HeadersToDownstreamOnDeny       *[]string          `tfsdk:"headers_to_downstream_on_deny" json:"headersToDownstreamOnDeny,omitempty"`
					HeadersToUpstreamOnAllow        *[]string          `tfsdk:"headers_to_upstream_on_allow" json:"headersToUpstreamOnAllow,omitempty"`
					IncludeAdditionalHeadersInCheck *map[string]string `tfsdk:"include_additional_headers_in_check" json:"includeAdditionalHeadersInCheck,omitempty"`
					IncludeHeadersInCheck           *[]string          `tfsdk:"include_headers_in_check" json:"includeHeadersInCheck,omitempty"`
					IncludeRequestBodyInCheck       *struct {
						AllowPartialMessage *bool  `tfsdk:"allow_partial_message" json:"allowPartialMessage,omitempty"`
						MaxRequestBytes     *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
						PackAsBytes         *bool  `tfsdk:"pack_as_bytes" json:"packAsBytes,omitempty"`
					} `tfsdk:"include_request_body_in_check" json:"includeRequestBodyInCheck,omitempty"`
					IncludeRequestHeadersInCheck *[]string `tfsdk:"include_request_headers_in_check" json:"includeRequestHeadersInCheck,omitempty"`
					PathPrefix                   *string   `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
					Port                         *int64    `tfsdk:"port" json:"port,omitempty"`
					Service                      *string   `tfsdk:"service" json:"service,omitempty"`
					StatusOnError                *string   `tfsdk:"status_on_error" json:"statusOnError,omitempty"`
					Timeout                      *string   `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"envoy_ext_authz_http" json:"envoyExtAuthzHttp,omitempty"`
				EnvoyFileAccessLog *struct {
					LogFormat *struct {
						Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Text   *string            `tfsdk:"text" json:"text,omitempty"`
					} `tfsdk:"log_format" json:"logFormat,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"envoy_file_access_log" json:"envoyFileAccessLog,omitempty"`
				EnvoyHttpAls *struct {
					AdditionalRequestHeadersToLog   *[]string `tfsdk:"additional_request_headers_to_log" json:"additionalRequestHeadersToLog,omitempty"`
					AdditionalResponseHeadersToLog  *[]string `tfsdk:"additional_response_headers_to_log" json:"additionalResponseHeadersToLog,omitempty"`
					AdditionalResponseTrailersToLog *[]string `tfsdk:"additional_response_trailers_to_log" json:"additionalResponseTrailersToLog,omitempty"`
					FilterStateObjectsToLog         *[]string `tfsdk:"filter_state_objects_to_log" json:"filterStateObjectsToLog,omitempty"`
					LogName                         *string   `tfsdk:"log_name" json:"logName,omitempty"`
					Port                            *int64    `tfsdk:"port" json:"port,omitempty"`
					Service                         *string   `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"envoy_http_als" json:"envoyHttpAls,omitempty"`
				EnvoyOtelAls *struct {
					LogFormat *struct {
						Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Text   *string            `tfsdk:"text" json:"text,omitempty"`
					} `tfsdk:"log_format" json:"logFormat,omitempty"`
					LogName *string `tfsdk:"log_name" json:"logName,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
					Service *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"envoy_otel_als" json:"envoyOtelAls,omitempty"`
				EnvoyTcpAls *struct {
					FilterStateObjectsToLog *[]string `tfsdk:"filter_state_objects_to_log" json:"filterStateObjectsToLog,omitempty"`
					LogName                 *string   `tfsdk:"log_name" json:"logName,omitempty"`
					Port                    *int64    `tfsdk:"port" json:"port,omitempty"`
					Service                 *string   `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"envoy_tcp_als" json:"envoyTcpAls,omitempty"`
				Lightstep *struct {
					AccessToken  *string `tfsdk:"access_token" json:"accessToken,omitempty"`
					MaxTagLength *int64  `tfsdk:"max_tag_length" json:"maxTagLength,omitempty"`
					Port         *int64  `tfsdk:"port" json:"port,omitempty"`
					Service      *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"lightstep" json:"lightstep,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Opencensus *struct {
					Context      *[]string `tfsdk:"context" json:"context,omitempty"`
					MaxTagLength *int64    `tfsdk:"max_tag_length" json:"maxTagLength,omitempty"`
					Port         *int64    `tfsdk:"port" json:"port,omitempty"`
					Service      *string   `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"opencensus" json:"opencensus,omitempty"`
				Opentelemetry *struct {
					MaxTagLength *int64  `tfsdk:"max_tag_length" json:"maxTagLength,omitempty"`
					Port         *int64  `tfsdk:"port" json:"port,omitempty"`
					Service      *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"opentelemetry" json:"opentelemetry,omitempty"`
				Prometheus *map[string]string `tfsdk:"prometheus" json:"prometheus,omitempty"`
				Skywalking *struct {
					AccessToken *string `tfsdk:"access_token" json:"accessToken,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Service     *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"skywalking" json:"skywalking,omitempty"`
				Stackdriver *struct {
					Debug   *bool `tfsdk:"debug" json:"debug,omitempty"`
					Logging *struct {
						Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"logging" json:"logging,omitempty"`
					MaxNumberOfAnnotations   *int64 `tfsdk:"max_number_of_annotations" json:"maxNumberOfAnnotations,omitempty"`
					MaxNumberOfAttributes    *int64 `tfsdk:"max_number_of_attributes" json:"maxNumberOfAttributes,omitempty"`
					MaxNumberOfMessageEvents *int64 `tfsdk:"max_number_of_message_events" json:"maxNumberOfMessageEvents,omitempty"`
					MaxTagLength             *int64 `tfsdk:"max_tag_length" json:"maxTagLength,omitempty"`
				} `tfsdk:"stackdriver" json:"stackdriver,omitempty"`
				Zipkin *struct {
					MaxTagLength *int64  `tfsdk:"max_tag_length" json:"maxTagLength,omitempty"`
					Port         *int64  `tfsdk:"port" json:"port,omitempty"`
					Service      *string `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"zipkin" json:"zipkin,omitempty"`
			} `tfsdk:"extension_providers" json:"extensionProviders,omitempty"`
			H2UpgradePolicy        *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
			InboundClusterStatName *string `tfsdk:"inbound_cluster_stat_name" json:"inboundClusterStatName,omitempty"`
			IngressClass           *string `tfsdk:"ingress_class" json:"ingressClass,omitempty"`
			IngressControllerMode  *string `tfsdk:"ingress_controller_mode" json:"ingressControllerMode,omitempty"`
			IngressSelector        *string `tfsdk:"ingress_selector" json:"ingressSelector,omitempty"`
			IngressService         *string `tfsdk:"ingress_service" json:"ingressService,omitempty"`
			LocalityLbSetting      *struct {
				Distribute *[]struct {
					From *string            `tfsdk:"from" json:"from,omitempty"`
					To   *map[string]string `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"distribute" json:"distribute,omitempty"`
				Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				Failover *[]struct {
					From *string `tfsdk:"from" json:"from,omitempty"`
					To   *string `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"failover" json:"failover,omitempty"`
				FailoverPriority *[]string `tfsdk:"failover_priority" json:"failoverPriority,omitempty"`
			} `tfsdk:"locality_lb_setting" json:"localityLbSetting,omitempty"`
			MeshMTLS *struct {
				MinProtocolVersion *string `tfsdk:"min_protocol_version" json:"minProtocolVersion,omitempty"`
			} `tfsdk:"mesh_mtls" json:"meshMTLS,omitempty"`
			OutboundClusterStatName *string `tfsdk:"outbound_cluster_stat_name" json:"outboundClusterStatName,omitempty"`
			OutboundTrafficPolicy   *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"outbound_traffic_policy" json:"outboundTrafficPolicy,omitempty"`
			PathNormalization *struct {
				Normalization *string `tfsdk:"normalization" json:"normalization,omitempty"`
			} `tfsdk:"path_normalization" json:"pathNormalization,omitempty"`
			ProtocolDetectionTimeout *string `tfsdk:"protocol_detection_timeout" json:"protocolDetectionTimeout,omitempty"`
			ProxyHttpPort            *int64  `tfsdk:"proxy_http_port" json:"proxyHttpPort,omitempty"`
			ProxyListenPort          *int64  `tfsdk:"proxy_listen_port" json:"proxyListenPort,omitempty"`
			RootNamespace            *string `tfsdk:"root_namespace" json:"rootNamespace,omitempty"`
			ServiceSettings          *[]struct {
				Hosts    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				Settings *struct {
					ClusterLocal *bool `tfsdk:"cluster_local" json:"clusterLocal,omitempty"`
				} `tfsdk:"settings" json:"settings,omitempty"`
			} `tfsdk:"service_settings" json:"serviceSettings,omitempty"`
			TcpKeepalive *struct {
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
				Time     *string `tfsdk:"time" json:"time,omitempty"`
			} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
			TrustDomain               *string   `tfsdk:"trust_domain" json:"trustDomain,omitempty"`
			TrustDomainAliases        *[]string `tfsdk:"trust_domain_aliases" json:"trustDomainAliases,omitempty"`
			VerifyCertificateAtClient *bool     `tfsdk:"verify_certificate_at_client" json:"verifyCertificateAtClient,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ServicemeshCiscoComIstioMeshV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_servicemesh_cisco_com_istio_mesh_v1alpha1_manifest"
}

func (r *ServicemeshCiscoComIstioMeshV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"access_log_encoding": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("TEXT", "JSON"),
								},
							},

							"access_log_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_log_format": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca": schema.SingleNestedAttribute{
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

									"istiod_side": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_timeout": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_settings": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"ca_certificates": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"client_certificate": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"credential_name": schema.StringAttribute{
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

											"mode": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
												},
											},

											"private_key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sni": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"subject_alt_names": schema.ListAttribute{
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

							"ca_certificates": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cert_signers": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pem": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"spiffe_bundle_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"trust_domains": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"certificates": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dns_names": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
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

							"config_sources": schema.ListNestedAttribute{
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

										"subscribed_resources": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_settings": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"ca_certificates": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_certificate": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"credential_name": schema.StringAttribute{
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

												"mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
													},
												},

												"private_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sni": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"subject_alt_names": schema.ListAttribute{
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

							"connect_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"availability_zone": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"binary_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_certificates_pem": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
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

									"config_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"control_plane_auth_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("NONE", "MUTUAL_TLS", "INHERIT"),
										},
									},

									"custom_config_file": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"discovery_address": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"discovery_refresh_delay": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"drain_duration": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"envoy_access_log_service": schema.SingleNestedAttribute{
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

											"tcp_keepalive": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"probes": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"time": schema.StringAttribute{
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

											"tls_settings": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"ca_certificates": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_certificate": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credential_name": schema.StringAttribute{
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

													"mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
														},
													},

													"private_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sni": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subject_alt_names": schema.ListAttribute{
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

									"envoy_metrics_service": schema.SingleNestedAttribute{
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

											"tcp_keepalive": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"probes": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"time": schema.StringAttribute{
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

											"tls_settings": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"ca_certificates": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_certificate": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credential_name": schema.StringAttribute{
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

													"mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
														},
													},

													"private_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sni": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subject_alt_names": schema.ListAttribute{
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

									"envoy_metrics_service_address": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"extra_stat_tags": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"gateway_topology": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"forward_client_cert_details": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("UNDEFINED", "SANITIZE", "FORWARD_ONLY", "APPEND_FORWARD", "SANITIZE_SET", "ALWAYS_FORWARD_ONLY"),
												},
											},

											"num_trusted_proxies": schema.Int64Attribute{
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

									"hold_application_until_proxy_starts": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"image_type": schema.StringAttribute{
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

									"interception_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("REDIRECT", "TPROXY", "NONE"),
										},
									},

									"mesh_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"private_key_provider": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"cryptomb": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"poll_delay": schema.StringAttribute{
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

											"qat": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"poll_delay": schema.StringAttribute{
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

									"proxy_admin_port": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy_bootstrap_template_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy_metadata": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy_stats_matcher": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"inclusion_prefixes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"inclusion_regexps": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"inclusion_suffixes": schema.ListAttribute{
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

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_headers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
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

													"path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
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

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_socket": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
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

											"timeout_seconds": schema.Int64Attribute{
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

									"runtime_values": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sds": schema.SingleNestedAttribute{
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

											"k8s_sa_jwt_path": schema.StringAttribute{
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

									"service_cluster": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stat_name_length": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"statsd_udp_address": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status_port": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"termination_drain_duration": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tracing": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"custom_tags": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"environment": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_value": schema.StringAttribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_value": schema.StringAttribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"literal": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"datadog": schema.SingleNestedAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"lightstep": schema.SingleNestedAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_path_tag_length": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"open_census_agent": schema.SingleNestedAttribute{
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

													"context": schema.ListAttribute{
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

											"sampling": schema.Float64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"stackdriver": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"debug": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_number_of_annotations": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_number_of_attributes": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_number_of_message_events": schema.Int64Attribute{
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

											"tls_settings": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"ca_certificates": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_certificate": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credential_name": schema.StringAttribute{
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

													"mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
														},
													},

													"private_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sni": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subject_alt_names": schema.ListAttribute{
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

											"zipkin": schema.SingleNestedAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tracing_service_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("APP_LABEL_AND_NAMESPACE", "CANONICAL_NAME_ONLY", "CANONICAL_NAME_AND_NAMESPACE"),
										},
									},

									"zipkin_address": schema.StringAttribute{
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

							"default_destination_rule_export_to": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_http_retry_policy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"attempts": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"per_try_timeout": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_on": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_remote_localities": schema.BoolAttribute{
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

							"default_providers": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"access_logging": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tracing": schema.ListAttribute{
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

							"default_service_export_to": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_virtual_service_export_to": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_envoy_listener_log": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"discovery_selectors": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
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

													"operator": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
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
											Description:         "",
											MarkdownDescription: "",
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

							"dns_refresh_rate": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_auto_mtls": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_envoy_access_log_service": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_prometheus_merge": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_tracing": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extension_providers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"datadog": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"max_tag_length": schema.Int64Attribute{
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

												"service": schema.StringAttribute{
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

										"envoy_ext_authz_grpc": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fail_open": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"include_request_body_in_check": schema.SingleNestedAttribute{
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

												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"status_on_error": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"envoy_ext_authz_http": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fail_open": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers_to_downstream_on_allow": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers_to_downstream_on_deny": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers_to_upstream_on_allow": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"include_additional_headers_in_check": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"include_headers_in_check": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"include_request_body_in_check": schema.SingleNestedAttribute{
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

												"include_request_headers_in_check": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
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

												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"status_on_error": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"envoy_file_access_log": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"log_format": schema.SingleNestedAttribute{
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

										"envoy_http_als": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"additional_request_headers_to_log": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"additional_response_headers_to_log": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"additional_response_trailers_to_log": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"filter_state_objects_to_log": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"log_name": schema.StringAttribute{
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

												"service": schema.StringAttribute{
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

										"envoy_otel_als": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"log_format": schema.SingleNestedAttribute{
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

												"log_name": schema.StringAttribute{
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

												"service": schema.StringAttribute{
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

										"envoy_tcp_als": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"filter_state_objects_to_log": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"log_name": schema.StringAttribute{
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

												"service": schema.StringAttribute{
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

										"lightstep": schema.SingleNestedAttribute{
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

												"max_tag_length": schema.Int64Attribute{
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

												"service": schema.StringAttribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"opencensus": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"context": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_tag_length": schema.Int64Attribute{
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

												"service": schema.StringAttribute{
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

										"opentelemetry": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"max_tag_length": schema.Int64Attribute{
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

												"service": schema.StringAttribute{
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

										"prometheus": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"skywalking": schema.SingleNestedAttribute{
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

												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service": schema.StringAttribute{
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

										"stackdriver": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"debug": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"logging": schema.SingleNestedAttribute{
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

												"max_number_of_annotations": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_number_of_attributes": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_number_of_message_events": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_tag_length": schema.Int64Attribute{
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

										"zipkin": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"max_tag_length": schema.Int64Attribute{
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

												"service": schema.StringAttribute{
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

							"h2_upgrade_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("DO_NOT_UPGRADE", "UPGRADE"),
								},
							},

							"inbound_cluster_stat_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_class": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_controller_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("UNSPECIFIED", "OFF", "DEFAULT", "STRICT"),
								},
							},

							"ingress_selector": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_service": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"locality_lb_setting": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"distribute": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"from": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"to": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"failover": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"from": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"to": schema.StringAttribute{
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

									"failover_priority": schema.ListAttribute{
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

							"mesh_mtls": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"min_protocol_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TLS_AUTO", "TLSV1_2", "TLSV1_3"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"outbound_cluster_stat_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"outbound_traffic_policy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("REGISTRY_ONLY", "ALLOW_ANY"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"path_normalization": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"normalization": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DEFAULT", "NONE", "BASE", "MERGE_SLASHES", "DECODE_AND_MERGE_SLASHES"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"protocol_detection_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_http_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_listen_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"root_namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_settings": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hosts": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"settings": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"cluster_local": schema.BoolAttribute{
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

							"tcp_keepalive": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"probes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"time": schema.StringAttribute{
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

							"trust_domain": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trust_domain_aliases": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_certificate_at_client": schema.BoolAttribute{
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
		},
	}
}

func (r *ServicemeshCiscoComIstioMeshV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_servicemesh_cisco_com_istio_mesh_v1alpha1_manifest")

	var model ServicemeshCiscoComIstioMeshV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("servicemesh.cisco.com/v1alpha1")
	model.Kind = pointer.String("IstioMesh")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
