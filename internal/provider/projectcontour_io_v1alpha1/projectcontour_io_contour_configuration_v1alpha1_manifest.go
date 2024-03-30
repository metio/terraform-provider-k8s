/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package projectcontour_io_v1alpha1

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
	_ datasource.DataSource = &ProjectcontourIoContourConfigurationV1Alpha1Manifest{}
)

func NewProjectcontourIoContourConfigurationV1Alpha1Manifest() datasource.DataSource {
	return &ProjectcontourIoContourConfigurationV1Alpha1Manifest{}
}

type ProjectcontourIoContourConfigurationV1Alpha1Manifest struct{}

type ProjectcontourIoContourConfigurationV1Alpha1ManifestData struct {
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
		Debug *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"debug" json:"debug,omitempty"`
		EnableExternalNameService *bool `tfsdk:"enable_external_name_service" json:"enableExternalNameService,omitempty"`
		Envoy                     *struct {
			ClientCertificate *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
			Cluster *struct {
				CircuitBreakers *struct {
					MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
					MaxRequests        *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
					MaxRetries         *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
				} `tfsdk:"circuit_breakers" json:"circuitBreakers,omitempty"`
				DnsLookupFamily                   *string `tfsdk:"dns_lookup_family" json:"dnsLookupFamily,omitempty"`
				MaxRequestsPerConnection          *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
				Per_connection_buffer_limit_bytes *int64  `tfsdk:"per_connection_buffer_limit_bytes" json:"per-connection-buffer-limit-bytes,omitempty"`
				UpstreamTLS                       *struct {
					CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
					MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
					MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
				} `tfsdk:"upstream_tls" json:"upstreamTLS,omitempty"`
			} `tfsdk:"cluster" json:"cluster,omitempty"`
			DefaultHTTPVersions *[]string `tfsdk:"default_http_versions" json:"defaultHTTPVersions,omitempty"`
			Health              *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"health" json:"health,omitempty"`
			Http *struct {
				AccessLog *string `tfsdk:"access_log" json:"accessLog,omitempty"`
				Address   *string `tfsdk:"address" json:"address,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Https *struct {
				AccessLog *string `tfsdk:"access_log" json:"accessLog,omitempty"`
				Address   *string `tfsdk:"address" json:"address,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"https" json:"https,omitempty"`
			Listener *struct {
				ConnectionBalancer                *string `tfsdk:"connection_balancer" json:"connectionBalancer,omitempty"`
				DisableAllowChunkedLength         *bool   `tfsdk:"disable_allow_chunked_length" json:"disableAllowChunkedLength,omitempty"`
				DisableMergeSlashes               *bool   `tfsdk:"disable_merge_slashes" json:"disableMergeSlashes,omitempty"`
				HttpMaxConcurrentStreams          *int64  `tfsdk:"http_max_concurrent_streams" json:"httpMaxConcurrentStreams,omitempty"`
				MaxConnectionsPerListener         *int64  `tfsdk:"max_connections_per_listener" json:"maxConnectionsPerListener,omitempty"`
				MaxRequestsPerConnection          *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
				MaxRequestsPerIOCycle             *int64  `tfsdk:"max_requests_per_io_cycle" json:"maxRequestsPerIOCycle,omitempty"`
				Per_connection_buffer_limit_bytes *int64  `tfsdk:"per_connection_buffer_limit_bytes" json:"per-connection-buffer-limit-bytes,omitempty"`
				ServerHeaderTransformation        *string `tfsdk:"server_header_transformation" json:"serverHeaderTransformation,omitempty"`
				SocketOptions                     *struct {
					Tos          *int64 `tfsdk:"tos" json:"tos,omitempty"`
					TrafficClass *int64 `tfsdk:"traffic_class" json:"trafficClass,omitempty"`
				} `tfsdk:"socket_options" json:"socketOptions,omitempty"`
				Tls *struct {
					CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
					MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
					MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				UseProxyProtocol *bool `tfsdk:"use_proxy_protocol" json:"useProxyProtocol,omitempty"`
			} `tfsdk:"listener" json:"listener,omitempty"`
			Logging *struct {
				AccessLogFormat       *string   `tfsdk:"access_log_format" json:"accessLogFormat,omitempty"`
				AccessLogFormatString *string   `tfsdk:"access_log_format_string" json:"accessLogFormatString,omitempty"`
				AccessLogJSONFields   *[]string `tfsdk:"access_log_json_fields" json:"accessLogJSONFields,omitempty"`
				AccessLogLevel        *string   `tfsdk:"access_log_level" json:"accessLogLevel,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			Metrics *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Tls     *struct {
					CaFile   *string `tfsdk:"ca_file" json:"caFile,omitempty"`
					CertFile *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					KeyFile  *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Network *struct {
				AdminPort      *int64 `tfsdk:"admin_port" json:"adminPort,omitempty"`
				NumTrustedHops *int64 `tfsdk:"num_trusted_hops" json:"numTrustedHops,omitempty"`
			} `tfsdk:"network" json:"network,omitempty"`
			Service *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Timeouts *struct {
				ConnectTimeout                *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectionIdleTimeout         *string `tfsdk:"connection_idle_timeout" json:"connectionIdleTimeout,omitempty"`
				ConnectionShutdownGracePeriod *string `tfsdk:"connection_shutdown_grace_period" json:"connectionShutdownGracePeriod,omitempty"`
				DelayedCloseTimeout           *string `tfsdk:"delayed_close_timeout" json:"delayedCloseTimeout,omitempty"`
				MaxConnectionDuration         *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
				RequestTimeout                *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
				StreamIdleTimeout             *string `tfsdk:"stream_idle_timeout" json:"streamIdleTimeout,omitempty"`
			} `tfsdk:"timeouts" json:"timeouts,omitempty"`
		} `tfsdk:"envoy" json:"envoy,omitempty"`
		FeatureFlags *[]string `tfsdk:"feature_flags" json:"featureFlags,omitempty"`
		Gateway      *struct {
			GatewayRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"gateway_ref" json:"gatewayRef,omitempty"`
		} `tfsdk:"gateway" json:"gateway,omitempty"`
		GlobalExtAuth *struct {
			AuthPolicy *struct {
				Context  *map[string]string `tfsdk:"context" json:"context,omitempty"`
				Disabled *bool              `tfsdk:"disabled" json:"disabled,omitempty"`
			} `tfsdk:"auth_policy" json:"authPolicy,omitempty"`
			ExtensionRef *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"extension_ref" json:"extensionRef,omitempty"`
			FailOpen        *bool   `tfsdk:"fail_open" json:"failOpen,omitempty"`
			ResponseTimeout *string `tfsdk:"response_timeout" json:"responseTimeout,omitempty"`
			WithRequestBody *struct {
				AllowPartialMessage *bool  `tfsdk:"allow_partial_message" json:"allowPartialMessage,omitempty"`
				MaxRequestBytes     *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
				PackAsBytes         *bool  `tfsdk:"pack_as_bytes" json:"packAsBytes,omitempty"`
			} `tfsdk:"with_request_body" json:"withRequestBody,omitempty"`
		} `tfsdk:"global_ext_auth" json:"globalExtAuth,omitempty"`
		Health *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"health" json:"health,omitempty"`
		Httpproxy *struct {
			DisablePermitInsecure *bool `tfsdk:"disable_permit_insecure" json:"disablePermitInsecure,omitempty"`
			FallbackCertificate   *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"fallback_certificate" json:"fallbackCertificate,omitempty"`
			RootNamespaces *[]string `tfsdk:"root_namespaces" json:"rootNamespaces,omitempty"`
		} `tfsdk:"httpproxy" json:"httpproxy,omitempty"`
		Ingress *struct {
			ClassNames    *[]string `tfsdk:"class_names" json:"classNames,omitempty"`
			StatusAddress *string   `tfsdk:"status_address" json:"statusAddress,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Metrics *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			Tls     *struct {
				CaFile   *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CertFile *string `tfsdk:"cert_file" json:"certFile,omitempty"`
				KeyFile  *string `tfsdk:"key_file" json:"keyFile,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Policy *struct {
			ApplyToIngress *bool `tfsdk:"apply_to_ingress" json:"applyToIngress,omitempty"`
			RequestHeaders *struct {
				Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
				Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
			} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
			ResponseHeaders *struct {
				Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
				Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
			} `tfsdk:"response_headers" json:"responseHeaders,omitempty"`
		} `tfsdk:"policy" json:"policy,omitempty"`
		RateLimitService *struct {
			DefaultGlobalRateLimitPolicy *struct {
				Descriptors *[]struct {
					Entries *[]struct {
						GenericKey *struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						RemoteAddress *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeader *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_header" json:"requestHeader,omitempty"`
						RequestHeaderValueMatch *struct {
							ExpectMatch *bool `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers     *[]struct {
								Contains            *string `tfsdk:"contains" json:"contains,omitempty"`
								Exact               *string `tfsdk:"exact" json:"exact,omitempty"`
								IgnoreCase          *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
								Name                *string `tfsdk:"name" json:"name,omitempty"`
								Notcontains         *string `tfsdk:"notcontains" json:"notcontains,omitempty"`
								Notexact            *string `tfsdk:"notexact" json:"notexact,omitempty"`
								Notpresent          *bool   `tfsdk:"notpresent" json:"notpresent,omitempty"`
								Present             *bool   `tfsdk:"present" json:"present,omitempty"`
								Regex               *string `tfsdk:"regex" json:"regex,omitempty"`
								TreatMissingAsEmpty *bool   `tfsdk:"treat_missing_as_empty" json:"treatMissingAsEmpty,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"request_header_value_match" json:"requestHeaderValueMatch,omitempty"`
					} `tfsdk:"entries" json:"entries,omitempty"`
				} `tfsdk:"descriptors" json:"descriptors,omitempty"`
				Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
			} `tfsdk:"default_global_rate_limit_policy" json:"defaultGlobalRateLimitPolicy,omitempty"`
			Domain                      *string `tfsdk:"domain" json:"domain,omitempty"`
			EnableResourceExhaustedCode *bool   `tfsdk:"enable_resource_exhausted_code" json:"enableResourceExhaustedCode,omitempty"`
			EnableXRateLimitHeaders     *bool   `tfsdk:"enable_x_rate_limit_headers" json:"enableXRateLimitHeaders,omitempty"`
			ExtensionService            *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"extension_service" json:"extensionService,omitempty"`
			FailOpen *bool `tfsdk:"fail_open" json:"failOpen,omitempty"`
		} `tfsdk:"rate_limit_service" json:"rateLimitService,omitempty"`
		Tracing *struct {
			CustomTags *[]struct {
				Literal           *string `tfsdk:"literal" json:"literal,omitempty"`
				RequestHeaderName *string `tfsdk:"request_header_name" json:"requestHeaderName,omitempty"`
				TagName           *string `tfsdk:"tag_name" json:"tagName,omitempty"`
			} `tfsdk:"custom_tags" json:"customTags,omitempty"`
			ExtensionService *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"extension_service" json:"extensionService,omitempty"`
			IncludePodDetail *bool   `tfsdk:"include_pod_detail" json:"includePodDetail,omitempty"`
			MaxPathTagLength *int64  `tfsdk:"max_path_tag_length" json:"maxPathTagLength,omitempty"`
			OverallSampling  *string `tfsdk:"overall_sampling" json:"overallSampling,omitempty"`
			ServiceName      *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"tracing" json:"tracing,omitempty"`
		XdsServer *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			Tls     *struct {
				CaFile   *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CertFile *string `tfsdk:"cert_file" json:"certFile,omitempty"`
				Insecure *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
				KeyFile  *string `tfsdk:"key_file" json:"keyFile,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"xds_server" json:"xdsServer,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ProjectcontourIoContourConfigurationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_projectcontour_io_contour_configuration_v1alpha1_manifest"
}

func (r *ProjectcontourIoContourConfigurationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ContourConfiguration is the schema for a Contour instance.",
		MarkdownDescription: "ContourConfiguration is the schema for a Contour instance.",
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
				Description:         "ContourConfigurationSpec represents a configuration of a Contour controller.It contains most of all the options that can be customized, theother remaining options being command line flags.",
				MarkdownDescription: "ContourConfigurationSpec represents a configuration of a Contour controller.It contains most of all the options that can be customized, theother remaining options being command line flags.",
				Attributes: map[string]schema.Attribute{
					"debug": schema.SingleNestedAttribute{
						Description:         "Debug contains parameters to enable debug loggingand debug interfaces inside Contour.",
						MarkdownDescription: "Debug contains parameters to enable debug loggingand debug interfaces inside Contour.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Defines the Contour debug address interface.Contour's default is '127.0.0.1'.",
								MarkdownDescription: "Defines the Contour debug address interface.Contour's default is '127.0.0.1'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Defines the Contour debug address port.Contour's default is 6060.",
								MarkdownDescription: "Defines the Contour debug address port.Contour's default is 6060.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_external_name_service": schema.BoolAttribute{
						Description:         "EnableExternalNameService allows processing of ExternalNameServicesContour's default is false for security reasons.",
						MarkdownDescription: "EnableExternalNameService allows processing of ExternalNameServicesContour's default is false for security reasons.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"envoy": schema.SingleNestedAttribute{
						Description:         "Envoy contains parameters for Envoy as wellas how to optionally configure a managed Envoy fleet.",
						MarkdownDescription: "Envoy contains parameters for Envoy as wellas how to optionally configure a managed Envoy fleet.",
						Attributes: map[string]schema.Attribute{
							"client_certificate": schema.SingleNestedAttribute{
								Description:         "ClientCertificate defines the namespace/name of the Kubernetessecret containing the client certificate and private keyto be used when establishing TLS connection to upstreamcluster.",
								MarkdownDescription: "ClientCertificate defines the namespace/name of the Kubernetessecret containing the client certificate and private keyto be used when establishing TLS connection to upstreamcluster.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

							"cluster": schema.SingleNestedAttribute{
								Description:         "Cluster holds various configurable Envoy cluster values that canbe set in the config file.",
								MarkdownDescription: "Cluster holds various configurable Envoy cluster values that canbe set in the config file.",
								Attributes: map[string]schema.Attribute{
									"circuit_breakers": schema.SingleNestedAttribute{
										Description:         "GlobalCircuitBreakerDefaults specifies default circuit breaker budget across all services.If defined, this will be used as the default for all services.",
										MarkdownDescription: "GlobalCircuitBreakerDefaults specifies default circuit breaker budget across all services.If defined, this will be used as the default for all services.",
										Attributes: map[string]schema.Attribute{
											"max_connections": schema.Int64Attribute{
												Description:         "The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
												MarkdownDescription: "The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_pending_requests": schema.Int64Attribute{
												Description:         "The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
												MarkdownDescription: "The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_requests": schema.Int64Attribute{
												Description:         "The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024",
												MarkdownDescription: "The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_retries": schema.Int64Attribute{
												Description:         "The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3.",
												MarkdownDescription: "The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"dns_lookup_family": schema.StringAttribute{
										Description:         "DNSLookupFamily defines how external names are looked upWhen configured as V4, the DNS resolver will only perform a lookupfor addresses in the IPv4 family. If V6 is configured, the DNS resolverwill only perform a lookup for addresses in the IPv6 family.If AUTO is configured, the DNS resolver will first perform a lookupfor addresses in the IPv6 family and fallback to a lookup for addressesin the IPv4 family. If ALL is specified, the DNS resolver will perform a lookup forboth IPv4 and IPv6 families, and return all resolved addresses.When this is used, Happy Eyeballs will be enabled for upstream connections.Refer to Happy Eyeballs Support for more information.Note: This only applies to externalName clusters.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamilyfor more information.Values: 'auto' (default), 'v4', 'v6', 'all'.Other values will produce an error.",
										MarkdownDescription: "DNSLookupFamily defines how external names are looked upWhen configured as V4, the DNS resolver will only perform a lookupfor addresses in the IPv4 family. If V6 is configured, the DNS resolverwill only perform a lookup for addresses in the IPv6 family.If AUTO is configured, the DNS resolver will first perform a lookupfor addresses in the IPv6 family and fallback to a lookup for addressesin the IPv4 family. If ALL is specified, the DNS resolver will perform a lookup forboth IPv4 and IPv6 families, and return all resolved addresses.When this is used, Happy Eyeballs will be enabled for upstream connections.Refer to Happy Eyeballs Support for more information.Note: This only applies to externalName clusters.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamilyfor more information.Values: 'auto' (default), 'v4', 'v6', 'all'.Other values will produce an error.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_requests_per_connection": schema.Int64Attribute{
										Description:         "Defines the maximum requests for upstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
										MarkdownDescription: "Defines the maximum requests for upstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"per_connection_buffer_limit_bytes": schema.Int64Attribute{
										Description:         "Defines the soft limit on size of the cluster’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-per-connection-buffer-limit-bytesfor more information.",
										MarkdownDescription: "Defines the soft limit on size of the cluster’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-per-connection-buffer-limit-bytesfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"upstream_tls": schema.SingleNestedAttribute{
										Description:         "UpstreamTLS contains the TLS policy parameters for upstream connections",
										MarkdownDescription: "UpstreamTLS contains the TLS policy parameters for upstream connections",
										Attributes: map[string]schema.Attribute{
											"cipher_suites": schema.ListAttribute{
												Description:         "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
												MarkdownDescription: "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"maximum_protocol_version": schema.StringAttribute{
												Description:         "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
												MarkdownDescription: "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"minimum_protocol_version": schema.StringAttribute{
												Description:         "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
												MarkdownDescription: "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
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

							"default_http_versions": schema.ListAttribute{
								Description:         "DefaultHTTPVersions defines the default set of HTTPSversions the proxy should accept. HTTP versions arestrings of the form 'HTTP/xx'. Supported versions are'HTTP/1.1' and 'HTTP/2'.Values: 'HTTP/1.1', 'HTTP/2' (default: both).Other values will produce an error.",
								MarkdownDescription: "DefaultHTTPVersions defines the default set of HTTPSversions the proxy should accept. HTTP versions arestrings of the form 'HTTP/xx'. Supported versions are'HTTP/1.1' and 'HTTP/2'.Values: 'HTTP/1.1', 'HTTP/2' (default: both).Other values will produce an error.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"health": schema.SingleNestedAttribute{
								Description:         "Health defines the endpoint Envoy uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8002 }.",
								MarkdownDescription: "Health defines the endpoint Envoy uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8002 }.",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "Defines the health address interface.",
										MarkdownDescription: "Defines the health address interface.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines the health port.",
										MarkdownDescription: "Defines the health port.",
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
								Description:         "Defines the HTTP Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8080, accessLog: '/dev/stdout' }.",
								MarkdownDescription: "Defines the HTTP Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8080, accessLog: '/dev/stdout' }.",
								Attributes: map[string]schema.Attribute{
									"access_log": schema.StringAttribute{
										Description:         "AccessLog defines where Envoy logs are outputted for this listener.",
										MarkdownDescription: "AccessLog defines where Envoy logs are outputted for this listener.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"address": schema.StringAttribute{
										Description:         "Defines an Envoy Listener Address.",
										MarkdownDescription: "Defines an Envoy Listener Address.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines an Envoy listener Port.",
										MarkdownDescription: "Defines an Envoy listener Port.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"https": schema.SingleNestedAttribute{
								Description:         "Defines the HTTPS Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8443, accessLog: '/dev/stdout' }.",
								MarkdownDescription: "Defines the HTTPS Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8443, accessLog: '/dev/stdout' }.",
								Attributes: map[string]schema.Attribute{
									"access_log": schema.StringAttribute{
										Description:         "AccessLog defines where Envoy logs are outputted for this listener.",
										MarkdownDescription: "AccessLog defines where Envoy logs are outputted for this listener.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"address": schema.StringAttribute{
										Description:         "Defines an Envoy Listener Address.",
										MarkdownDescription: "Defines an Envoy Listener Address.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines an Envoy listener Port.",
										MarkdownDescription: "Defines an Envoy listener Port.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"listener": schema.SingleNestedAttribute{
								Description:         "Listener hold various configurable Envoy listener values.",
								MarkdownDescription: "Listener hold various configurable Envoy listener values.",
								Attributes: map[string]schema.Attribute{
									"connection_balancer": schema.StringAttribute{
										Description:         "ConnectionBalancer. If the value is exact, the listener will use the exact connection balancerSee https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/listener.proto#envoy-api-msg-listener-connectionbalanceconfigfor more information.Values: (empty string): use the default ConnectionBalancer, 'exact': use the Exact ConnectionBalancer.Other values will produce an error.",
										MarkdownDescription: "ConnectionBalancer. If the value is exact, the listener will use the exact connection balancerSee https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/listener.proto#envoy-api-msg-listener-connectionbalanceconfigfor more information.Values: (empty string): use the default ConnectionBalancer, 'exact': use the Exact ConnectionBalancer.Other values will produce an error.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_allow_chunked_length": schema.BoolAttribute{
										Description:         "DisableAllowChunkedLength disables the RFC-compliant Envoy behavior tostrip the 'Content-Length' header if 'Transfer-Encoding: chunked' isalso set. This is an emergency off-switch to revert back to Envoy'sdefault behavior in case of failures. Please file an issue if failuresare encountered.See: https://github.com/projectcontour/contour/issues/3221Contour's default is false.",
										MarkdownDescription: "DisableAllowChunkedLength disables the RFC-compliant Envoy behavior tostrip the 'Content-Length' header if 'Transfer-Encoding: chunked' isalso set. This is an emergency off-switch to revert back to Envoy'sdefault behavior in case of failures. Please file an issue if failuresare encountered.See: https://github.com/projectcontour/contour/issues/3221Contour's default is false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_merge_slashes": schema.BoolAttribute{
										Description:         "DisableMergeSlashes disables Envoy's non-standard merge_slashes path transformation optionwhich strips duplicate slashes from request URL paths.Contour's default is false.",
										MarkdownDescription: "DisableMergeSlashes disables Envoy's non-standard merge_slashes path transformation optionwhich strips duplicate slashes from request URL paths.Contour's default is false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http_max_concurrent_streams": schema.Int64Attribute{
										Description:         "Defines the value for SETTINGS_MAX_CONCURRENT_STREAMS Envoy will advertise in theSETTINGS frame in HTTP/2 connections and the limit for concurrent streams allowedfor a peer on a single HTTP/2 connection. It is recommended to not set this lowerthan 100 but this field can be used to bound resource usage by HTTP/2 connectionsand mitigate attacks like CVE-2023-44487. The default value when this is not set isunlimited.",
										MarkdownDescription: "Defines the value for SETTINGS_MAX_CONCURRENT_STREAMS Envoy will advertise in theSETTINGS frame in HTTP/2 connections and the limit for concurrent streams allowedfor a peer on a single HTTP/2 connection. It is recommended to not set this lowerthan 100 but this field can be used to bound resource usage by HTTP/2 connectionsand mitigate attacks like CVE-2023-44487. The default value when this is not set isunlimited.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"max_connections_per_listener": schema.Int64Attribute{
										Description:         "Defines the limit on number of active connections to a listener. The limit is appliedper listener. The default value when this is not set is unlimited.",
										MarkdownDescription: "Defines the limit on number of active connections to a listener. The limit is appliedper listener. The default value when this is not set is unlimited.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"max_requests_per_connection": schema.Int64Attribute{
										Description:         "Defines the maximum requests for downstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
										MarkdownDescription: "Defines the maximum requests for downstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"max_requests_per_io_cycle": schema.Int64Attribute{
										Description:         "Defines the limit on number of HTTP requests that Envoy will process from a singleconnection in a single I/O cycle. Requests over this limit are processed in subsequentI/O cycles. Can be used as a mitigation for CVE-2023-44487 when abusive traffic isdetected. Configures the http.max_requests_per_io_cycle Envoy runtime setting. The defaultvalue when this is not set is no limit.",
										MarkdownDescription: "Defines the limit on number of HTTP requests that Envoy will process from a singleconnection in a single I/O cycle. Requests over this limit are processed in subsequentI/O cycles. Can be used as a mitigation for CVE-2023-44487 when abusive traffic isdetected. Configures the http.max_requests_per_io_cycle Envoy runtime setting. The defaultvalue when this is not set is no limit.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"per_connection_buffer_limit_bytes": schema.Int64Attribute{
										Description:         "Defines the soft limit on size of the listener’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-per-connection-buffer-limit-bytesfor more information.",
										MarkdownDescription: "Defines the soft limit on size of the listener’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-per-connection-buffer-limit-bytesfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"server_header_transformation": schema.StringAttribute{
										Description:         "Defines the action to be applied to the Server header on the response path.When configured as overwrite, overwrites any Server header with 'envoy'.When configured as append_if_absent, if a Server header is present, pass it through, otherwise set it to 'envoy'.When configured as pass_through, pass through the value of the Server header, and do not append a header if none is present.Values: 'overwrite' (default), 'append_if_absent', 'pass_through'Other values will produce an error.Contour's default is overwrite.",
										MarkdownDescription: "Defines the action to be applied to the Server header on the response path.When configured as overwrite, overwrites any Server header with 'envoy'.When configured as append_if_absent, if a Server header is present, pass it through, otherwise set it to 'envoy'.When configured as pass_through, pass through the value of the Server header, and do not append a header if none is present.Values: 'overwrite' (default), 'append_if_absent', 'pass_through'Other values will produce an error.Contour's default is overwrite.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"socket_options": schema.SingleNestedAttribute{
										Description:         "SocketOptions defines configurable socket options for the listeners.Single set of options are applied to all listeners.",
										MarkdownDescription: "SocketOptions defines configurable socket options for the listeners.Single set of options are applied to all listeners.",
										Attributes: map[string]schema.Attribute{
											"tos": schema.Int64Attribute{
												Description:         "Defines the value for IPv4 TOS field (including 6 bit DSCP field) for IP packets originating from Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv6-only addresses, setting this option will cause an error.",
												MarkdownDescription: "Defines the value for IPv4 TOS field (including 6 bit DSCP field) for IP packets originating from Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv6-only addresses, setting this option will cause an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(255),
												},
											},

											"traffic_class": schema.Int64Attribute{
												Description:         "Defines the value for IPv6 Traffic Class field (including 6 bit DSCP field) for IP packets originating from the Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv4-only addresses, setting this option will cause an error.",
												MarkdownDescription: "Defines the value for IPv6 Traffic Class field (including 6 bit DSCP field) for IP packets originating from the Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv4-only addresses, setting this option will cause an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(255),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS holds various configurable Envoy TLS listener values.",
										MarkdownDescription: "TLS holds various configurable Envoy TLS listener values.",
										Attributes: map[string]schema.Attribute{
											"cipher_suites": schema.ListAttribute{
												Description:         "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
												MarkdownDescription: "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"maximum_protocol_version": schema.StringAttribute{
												Description:         "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
												MarkdownDescription: "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"minimum_protocol_version": schema.StringAttribute{
												Description:         "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
												MarkdownDescription: "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_proxy_protocol": schema.BoolAttribute{
										Description:         "Use PROXY protocol for all listeners.Contour's default is false.",
										MarkdownDescription: "Use PROXY protocol for all listeners.Contour's default is false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "Logging defines how Envoy's logs can be configured.",
								MarkdownDescription: "Logging defines how Envoy's logs can be configured.",
								Attributes: map[string]schema.Attribute{
									"access_log_format": schema.StringAttribute{
										Description:         "AccessLogFormat sets the global access log format.Values: 'envoy' (default), 'json'.Other values will produce an error.",
										MarkdownDescription: "AccessLogFormat sets the global access log format.Values: 'envoy' (default), 'json'.Other values will produce an error.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"access_log_format_string": schema.StringAttribute{
										Description:         "AccessLogFormatString sets the access log format when format is set to 'envoy'.When empty, Envoy's default format is used.",
										MarkdownDescription: "AccessLogFormatString sets the access log format when format is set to 'envoy'.When empty, Envoy's default format is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"access_log_json_fields": schema.ListAttribute{
										Description:         "AccessLogJSONFields sets the fields that JSON logging willoutput when AccessLogFormat is json.",
										MarkdownDescription: "AccessLogJSONFields sets the fields that JSON logging willoutput when AccessLogFormat is json.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"access_log_level": schema.StringAttribute{
										Description:         "AccessLogLevel sets the verbosity level of the access log.Values: 'info' (default, all requests are logged), 'error' (all non-success requests, i.e. 300+ response code, are logged), 'critical' (all 5xx requests are logged) and 'disabled'.Other values will produce an error.",
										MarkdownDescription: "AccessLogLevel sets the verbosity level of the access log.Values: 'info' (default, all requests are logged), 'error' (all non-success requests, i.e. 300+ response code, are logged), 'critical' (all 5xx requests are logged) and 'disabled'.Other values will produce an error.",
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
								Description:         "Metrics defines the endpoint Envoy uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8002 }.",
								MarkdownDescription: "Metrics defines the endpoint Envoy uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8002 }.",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "Defines the metrics address interface.",
										MarkdownDescription: "Defines the metrics address interface.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines the metrics port.",
										MarkdownDescription: "Defines the metrics port.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
										MarkdownDescription: "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
										Attributes: map[string]schema.Attribute{
											"ca_file": schema.StringAttribute{
												Description:         "CA filename.",
												MarkdownDescription: "CA filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_file": schema.StringAttribute{
												Description:         "Client certificate filename.",
												MarkdownDescription: "Client certificate filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_file": schema.StringAttribute{
												Description:         "Client key filename.",
												MarkdownDescription: "Client key filename.",
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

							"network": schema.SingleNestedAttribute{
								Description:         "Network holds various configurable Envoy network values.",
								MarkdownDescription: "Network holds various configurable Envoy network values.",
								Attributes: map[string]schema.Attribute{
									"admin_port": schema.Int64Attribute{
										Description:         "Configure the port used to access the Envoy Admin interface.If configured to port '0' then the admin interface is disabled.Contour's default is 9001.",
										MarkdownDescription: "Configure the port used to access the Envoy Admin interface.If configured to port '0' then the admin interface is disabled.Contour's default is 9001.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"num_trusted_hops": schema.Int64Attribute{
										Description:         "XffNumTrustedHops defines the number of additional ingress proxy hops from theright side of the x-forwarded-for HTTP header to trust when determining the originclient’s IP address.See https://www.envoyproxy.io/docs/envoy/v1.17.0/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto?highlight=xff_num_trusted_hopsfor more information.Contour's default is 0.",
										MarkdownDescription: "XffNumTrustedHops defines the number of additional ingress proxy hops from theright side of the x-forwarded-for HTTP header to trust when determining the originclient’s IP address.See https://www.envoyproxy.io/docs/envoy/v1.17.0/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto?highlight=xff_num_trusted_hopsfor more information.Contour's default is 0.",
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
								Description:         "Service holds Envoy service parameters for setting Ingress status.Contour's default is { namespace: 'projectcontour', name: 'envoy' }.",
								MarkdownDescription: "Service holds Envoy service parameters for setting Ingress status.Contour's default is { namespace: 'projectcontour', name: 'envoy' }.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

							"timeouts": schema.SingleNestedAttribute{
								Description:         "Timeouts holds various configurable timeouts that canbe set in the config file.",
								MarkdownDescription: "Timeouts holds various configurable timeouts that canbe set in the config file.",
								Attributes: map[string]schema.Attribute{
									"connect_timeout": schema.StringAttribute{
										Description:         "ConnectTimeout defines how long the proxy should wait when establishing connection to upstream service.If not set, a default value of 2 seconds will be used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-connect-timeoutfor more information.",
										MarkdownDescription: "ConnectTimeout defines how long the proxy should wait when establishing connection to upstream service.If not set, a default value of 2 seconds will be used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-connect-timeoutfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connection_idle_timeout": schema.StringAttribute{
										Description:         "ConnectionIdleTimeout defines how long the proxy should wait while there areno active requests (for HTTP/1.1) or streams (for HTTP/2) before terminatingan HTTP connection. Set to 'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-idle-timeoutfor more information.",
										MarkdownDescription: "ConnectionIdleTimeout defines how long the proxy should wait while there areno active requests (for HTTP/1.1) or streams (for HTTP/2) before terminatingan HTTP connection. Set to 'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-idle-timeoutfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connection_shutdown_grace_period": schema.StringAttribute{
										Description:         "ConnectionShutdownGracePeriod defines how long the proxy will wait between sending aninitial GOAWAY frame and a second, final GOAWAY frame when terminating an HTTP/2 connection.During this grace period, the proxy will continue to respond to new streams. After the finalGOAWAY frame has been sent, the proxy will refuse new streams.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-drain-timeoutfor more information.",
										MarkdownDescription: "ConnectionShutdownGracePeriod defines how long the proxy will wait between sending aninitial GOAWAY frame and a second, final GOAWAY frame when terminating an HTTP/2 connection.During this grace period, the proxy will continue to respond to new streams. After the finalGOAWAY frame has been sent, the proxy will refuse new streams.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-drain-timeoutfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"delayed_close_timeout": schema.StringAttribute{
										Description:         "DelayedCloseTimeout defines how long envoy will wait, once connectionclose processing has been initiated, for the downstream peer to closethe connection before Envoy closes the socket associated with the connection.Setting this timeout to 'infinity' will disable it, equivalent to setting it to '0'in Envoy. Leaving it unset will result in the Envoy default value being used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-delayed-close-timeoutfor more information.",
										MarkdownDescription: "DelayedCloseTimeout defines how long envoy will wait, once connectionclose processing has been initiated, for the downstream peer to closethe connection before Envoy closes the socket associated with the connection.Setting this timeout to 'infinity' will disable it, equivalent to setting it to '0'in Envoy. Leaving it unset will result in the Envoy default value being used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-delayed-close-timeoutfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_connection_duration": schema.StringAttribute{
										Description:         "MaxConnectionDuration defines the maximum period of time after an HTTP connectionhas been established from the client to the proxy before it is closed by the proxy,regardless of whether there has been activity or not. Omit or set to 'infinity' forno max duration.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-max-connection-durationfor more information.",
										MarkdownDescription: "MaxConnectionDuration defines the maximum period of time after an HTTP connectionhas been established from the client to the proxy before it is closed by the proxy,regardless of whether there has been activity or not. Omit or set to 'infinity' forno max duration.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-max-connection-durationfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_timeout": schema.StringAttribute{
										Description:         "RequestTimeout sets the client request timeout globally for Contour. Note thatthis is a timeout for the entire request, not an idle timeout. Omit or set to'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-request-timeoutfor more information.",
										MarkdownDescription: "RequestTimeout sets the client request timeout globally for Contour. Note thatthis is a timeout for the entire request, not an idle timeout. Omit or set to'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-request-timeoutfor more information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stream_idle_timeout": schema.StringAttribute{
										Description:         "StreamIdleTimeout defines how long the proxy should wait while there is norequest activity (for HTTP/1.1) or stream activity (for HTTP/2) beforeterminating the HTTP request or stream. Set to 'infinity' to disable thetimeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-stream-idle-timeoutfor more information.",
										MarkdownDescription: "StreamIdleTimeout defines how long the proxy should wait while there is norequest activity (for HTTP/1.1) or stream activity (for HTTP/2) beforeterminating the HTTP request or stream. Set to 'infinity' to disable thetimeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-stream-idle-timeoutfor more information.",
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

					"feature_flags": schema.ListAttribute{
						Description:         "FeatureFlags defines toggle to enable new contour features.Available toggles are:useEndpointSlices - configures contour to fetch endpoint datafrom k8s endpoint slices. defaults to false and reading endpointdata from the k8s endpoints.",
						MarkdownDescription: "FeatureFlags defines toggle to enable new contour features.Available toggles are:useEndpointSlices - configures contour to fetch endpoint datafrom k8s endpoint slices. defaults to false and reading endpointdata from the k8s endpoints.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gateway": schema.SingleNestedAttribute{
						Description:         "Gateway contains parameters for the gateway-api Gateway that Contouris configured to serve traffic.",
						MarkdownDescription: "Gateway contains parameters for the gateway-api Gateway that Contouris configured to serve traffic.",
						Attributes: map[string]schema.Attribute{
							"gateway_ref": schema.SingleNestedAttribute{
								Description:         "GatewayRef defines the specific Gateway that this Contourinstance corresponds to.",
								MarkdownDescription: "GatewayRef defines the specific Gateway that this Contourinstance corresponds to.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
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

					"global_ext_auth": schema.SingleNestedAttribute{
						Description:         "GlobalExternalAuthorization allows envoys external authorization filterto be enabled for all virtual hosts.",
						MarkdownDescription: "GlobalExternalAuthorization allows envoys external authorization filterto be enabled for all virtual hosts.",
						Attributes: map[string]schema.Attribute{
							"auth_policy": schema.SingleNestedAttribute{
								Description:         "AuthPolicy sets a default authorization policy for client requests.This policy will be used unless overridden by individual routes.",
								MarkdownDescription: "AuthPolicy sets a default authorization policy for client requests.This policy will be used unless overridden by individual routes.",
								Attributes: map[string]schema.Attribute{
									"context": schema.MapAttribute{
										Description:         "Context is a set of key/value pairs that are sent to theauthentication server in the check request. If a contextis provided at an enclosing scope, the entries are mergedsuch that the inner scope overrides matching keys from theouter scope.",
										MarkdownDescription: "Context is a set of key/value pairs that are sent to theauthentication server in the check request. If a contextis provided at an enclosing scope, the entries are mergedsuch that the inner scope overrides matching keys from theouter scope.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disabled": schema.BoolAttribute{
										Description:         "When true, this field disables client request authenticationfor the scope of the policy.",
										MarkdownDescription: "When true, this field disables client request authenticationfor the scope of the policy.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"extension_ref": schema.SingleNestedAttribute{
								Description:         "ExtensionServiceRef specifies the extension resource that will authorize client requests.",
								MarkdownDescription: "ExtensionServiceRef specifies the extension resource that will authorize client requests.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.If this field is not specified, the default 'projectcontour.io/v1alpha1' will be used",
										MarkdownDescription: "API version of the referent.If this field is not specified, the default 'projectcontour.io/v1alpha1' will be used",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the referent.If this field is not specifies, the namespace of the resource that targets the referent will be used.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent.If this field is not specifies, the namespace of the resource that targets the referent will be used.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										Required:            false,
										Optional:            true,
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

							"fail_open": schema.BoolAttribute{
								Description:         "If FailOpen is true, the client request is forwarded to the upstream serviceeven if the authorization server fails to respond. This field should not beset in most cases. It is intended for use only while migrating applicationsfrom internal authorization to Contour external authorization.",
								MarkdownDescription: "If FailOpen is true, the client request is forwarded to the upstream serviceeven if the authorization server fails to respond. This field should not beset in most cases. It is intended for use only while migrating applicationsfrom internal authorization to Contour external authorization.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"response_timeout": schema.StringAttribute{
								Description:         "ResponseTimeout configures maximum time to wait for a check response from the authorization server.Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.The string 'infinity' is also a valid input and specifies no timeout.",
								MarkdownDescription: "ResponseTimeout configures maximum time to wait for a check response from the authorization server.Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.The string 'infinity' is also a valid input and specifies no timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
								},
							},

							"with_request_body": schema.SingleNestedAttribute{
								Description:         "WithRequestBody specifies configuration for sending the client request's body to authorization server.",
								MarkdownDescription: "WithRequestBody specifies configuration for sending the client request's body to authorization server.",
								Attributes: map[string]schema.Attribute{
									"allow_partial_message": schema.BoolAttribute{
										Description:         "If AllowPartialMessage is true, then Envoy will buffer the body until MaxRequestBytes are reached.",
										MarkdownDescription: "If AllowPartialMessage is true, then Envoy will buffer the body until MaxRequestBytes are reached.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_request_bytes": schema.Int64Attribute{
										Description:         "MaxRequestBytes sets the maximum size of message body ExtAuthz filter will hold in-memory.",
										MarkdownDescription: "MaxRequestBytes sets the maximum size of message body ExtAuthz filter will hold in-memory.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"pack_as_bytes": schema.BoolAttribute{
										Description:         "If PackAsBytes is true, the body sent to Authorization Server is in raw bytes.",
										MarkdownDescription: "If PackAsBytes is true, the body sent to Authorization Server is in raw bytes.",
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

					"health": schema.SingleNestedAttribute{
						Description:         "Health defines the endpoints Contour uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8000 }.",
						MarkdownDescription: "Health defines the endpoints Contour uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8000 }.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Defines the health address interface.",
								MarkdownDescription: "Defines the health address interface.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "Defines the health port.",
								MarkdownDescription: "Defines the health port.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"httpproxy": schema.SingleNestedAttribute{
						Description:         "HTTPProxy defines parameters on HTTPProxy.",
						MarkdownDescription: "HTTPProxy defines parameters on HTTPProxy.",
						Attributes: map[string]schema.Attribute{
							"disable_permit_insecure": schema.BoolAttribute{
								Description:         "DisablePermitInsecure disables the use of thepermitInsecure field in HTTPProxy.Contour's default is false.",
								MarkdownDescription: "DisablePermitInsecure disables the use of thepermitInsecure field in HTTPProxy.Contour's default is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fallback_certificate": schema.SingleNestedAttribute{
								Description:         "FallbackCertificate defines the namespace/name of the Kubernetes secret touse as fallback when a non-SNI request is received.",
								MarkdownDescription: "FallbackCertificate defines the namespace/name of the Kubernetes secret touse as fallback when a non-SNI request is received.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
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

							"root_namespaces": schema.ListAttribute{
								Description:         "Restrict Contour to searching these namespaces for root ingress routes.",
								MarkdownDescription: "Restrict Contour to searching these namespaces for root ingress routes.",
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

					"ingress": schema.SingleNestedAttribute{
						Description:         "Ingress contains parameters for ingress options.",
						MarkdownDescription: "Ingress contains parameters for ingress options.",
						Attributes: map[string]schema.Attribute{
							"class_names": schema.ListAttribute{
								Description:         "Ingress Class Names Contour should use.",
								MarkdownDescription: "Ingress Class Names Contour should use.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"status_address": schema.StringAttribute{
								Description:         "Address to set in Ingress object status.",
								MarkdownDescription: "Address to set in Ingress object status.",
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
						Description:         "Metrics defines the endpoint Contour uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8000 }.",
						MarkdownDescription: "Metrics defines the endpoint Contour uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8000 }.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Defines the metrics address interface.",
								MarkdownDescription: "Defines the metrics address interface.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "Defines the metrics port.",
								MarkdownDescription: "Defines the metrics port.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
								MarkdownDescription: "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "CA filename.",
										MarkdownDescription: "CA filename.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cert_file": schema.StringAttribute{
										Description:         "Client certificate filename.",
										MarkdownDescription: "Client certificate filename.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_file": schema.StringAttribute{
										Description:         "Client key filename.",
										MarkdownDescription: "Client key filename.",
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

					"policy": schema.SingleNestedAttribute{
						Description:         "Policy specifies default policy applied if not overridden by the user",
						MarkdownDescription: "Policy specifies default policy applied if not overridden by the user",
						Attributes: map[string]schema.Attribute{
							"apply_to_ingress": schema.BoolAttribute{
								Description:         "ApplyToIngress determines if the Policies will apply to ingress objectsContour's default is false.",
								MarkdownDescription: "ApplyToIngress determines if the Policies will apply to ingress objectsContour's default is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"request_headers": schema.SingleNestedAttribute{
								Description:         "RequestHeadersPolicy defines the request headers set/removed on all routes",
								MarkdownDescription: "RequestHeadersPolicy defines the request headers set/removed on all routes",
								Attributes: map[string]schema.Attribute{
									"remove": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"set": schema.MapAttribute{
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

							"response_headers": schema.SingleNestedAttribute{
								Description:         "ResponseHeadersPolicy defines the response headers set/removed on all routes",
								MarkdownDescription: "ResponseHeadersPolicy defines the response headers set/removed on all routes",
								Attributes: map[string]schema.Attribute{
									"remove": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"set": schema.MapAttribute{
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

					"rate_limit_service": schema.SingleNestedAttribute{
						Description:         "RateLimitService optionally holds properties of the Rate Limit Serviceto be used for global rate limiting.",
						MarkdownDescription: "RateLimitService optionally holds properties of the Rate Limit Serviceto be used for global rate limiting.",
						Attributes: map[string]schema.Attribute{
							"default_global_rate_limit_policy": schema.SingleNestedAttribute{
								Description:         "DefaultGlobalRateLimitPolicy allows setting a default global rate limit policy for every HTTPProxy.HTTPProxy can overwrite this configuration.",
								MarkdownDescription: "DefaultGlobalRateLimitPolicy allows setting a default global rate limit policy for every HTTPProxy.HTTPProxy can overwrite this configuration.",
								Attributes: map[string]schema.Attribute{
									"descriptors": schema.ListNestedAttribute{
										Description:         "Descriptors defines the list of descriptors that willbe generated and sent to the rate limit service. Eachdescriptor contains 1+ key-value pair entries.",
										MarkdownDescription: "Descriptors defines the list of descriptors that willbe generated and sent to the rate limit service. Eachdescriptor contains 1+ key-value pair entries.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"entries": schema.ListNestedAttribute{
													Description:         "Entries is the list of key-value pair generators.",
													MarkdownDescription: "Entries is the list of key-value pair generators.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"generic_key": schema.SingleNestedAttribute{
																Description:         "GenericKey defines a descriptor entry with a static key and value.",
																MarkdownDescription: "GenericKey defines a descriptor entry with a static key and value.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "Key defines the key of the descriptor entry. If not set, thekey is set to 'generic_key'.",
																		MarkdownDescription: "Key defines the key of the descriptor entry. If not set, thekey is set to 'generic_key'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value defines the value of the descriptor entry.",
																		MarkdownDescription: "Value defines the value of the descriptor entry.",
																		Required:            false,
																		Optional:            true,
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

															"remote_address": schema.MapAttribute{
																Description:         "RemoteAddress defines a descriptor entry with a key of 'remote_address'and a value equal to the client's IP address (from x-forwarded-for).",
																MarkdownDescription: "RemoteAddress defines a descriptor entry with a key of 'remote_address'and a value equal to the client's IP address (from x-forwarded-for).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_header": schema.SingleNestedAttribute{
																Description:         "RequestHeader defines a descriptor entry that's populated only ifa given header is present on the request. The descriptor key is static,and the descriptor value is equal to the value of the header.",
																MarkdownDescription: "RequestHeader defines a descriptor entry that's populated only ifa given header is present on the request. The descriptor key is static,and the descriptor value is equal to the value of the header.",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "DescriptorKey defines the key to use on the descriptor entry.",
																		MarkdownDescription: "DescriptorKey defines the key to use on the descriptor entry.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																		},
																	},

																	"header_name": schema.StringAttribute{
																		Description:         "HeaderName defines the name of the header to look for on the request.",
																		MarkdownDescription: "HeaderName defines the name of the header to look for on the request.",
																		Required:            false,
																		Optional:            true,
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

															"request_header_value_match": schema.SingleNestedAttribute{
																Description:         "RequestHeaderValueMatch defines a descriptor entry that's populatedif the request's headers match a set of 1+ match criteria. Thedescriptor key is 'header_match', and the descriptor value is static.",
																MarkdownDescription: "RequestHeaderValueMatch defines a descriptor entry that's populatedif the request's headers match a set of 1+ match criteria. Thedescriptor key is 'header_match', and the descriptor value is static.",
																Attributes: map[string]schema.Attribute{
																	"expect_match": schema.BoolAttribute{
																		Description:         "ExpectMatch defines whether the request must positively match the matchcriteria in order to generate a descriptor entry (i.e. true), or notmatch the match criteria in order to generate a descriptor entry (i.e. false).The default is true.",
																		MarkdownDescription: "ExpectMatch defines whether the request must positively match the matchcriteria in order to generate a descriptor entry (i.e. true), or notmatch the match criteria in order to generate a descriptor entry (i.e. false).The default is true.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "Headers is a list of 1+ match criteria to apply against the requestto determine whether to populate the descriptor entry or not.",
																		MarkdownDescription: "Headers is a list of 1+ match criteria to apply against the requestto determine whether to populate the descriptor entry or not.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"contains": schema.StringAttribute{
																					Description:         "Contains specifies a substring that must be present inthe header value.",
																					MarkdownDescription: "Contains specifies a substring that must be present inthe header value.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exact": schema.StringAttribute{
																					Description:         "Exact specifies a string that the header value must be equal to.",
																					MarkdownDescription: "Exact specifies a string that the header value must be equal to.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ignore_case": schema.BoolAttribute{
																					Description:         "IgnoreCase specifies that string matching should be case insensitive.Note that this has no effect on the Regex parameter.",
																					MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive.Note that this has no effect on the Regex parameter.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name is the name of the header to match against. Name is required.Header names are case insensitive.",
																					MarkdownDescription: "Name is the name of the header to match against. Name is required.Header names are case insensitive.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"notcontains": schema.StringAttribute{
																					Description:         "NotContains specifies a substring that must not be presentin the header value.",
																					MarkdownDescription: "NotContains specifies a substring that must not be presentin the header value.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"notexact": schema.StringAttribute{
																					Description:         "NoExact specifies a string that the header value must not beequal to. The condition is true if the header has any other value.",
																					MarkdownDescription: "NoExact specifies a string that the header value must not beequal to. The condition is true if the header has any other value.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"notpresent": schema.BoolAttribute{
																					Description:         "NotPresent specifies that condition is true when the named headeris not present. Note that setting NotPresent to false does notmake the condition true if the named header is present.",
																					MarkdownDescription: "NotPresent specifies that condition is true when the named headeris not present. Note that setting NotPresent to false does notmake the condition true if the named header is present.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present": schema.BoolAttribute{
																					Description:         "Present specifies that condition is true when the named headeris present, regardless of its value. Note that setting Presentto false does not make the condition true if the named headeris absent.",
																					MarkdownDescription: "Present specifies that condition is true when the named headeris present, regardless of its value. Note that setting Presentto false does not make the condition true if the named headeris absent.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "Regex specifies a regular expression pattern that must match the headervalue.",
																					MarkdownDescription: "Regex specifies a regular expression pattern that must match the headervalue.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"treat_missing_as_empty": schema.BoolAttribute{
																					Description:         "TreatMissingAsEmpty specifies if the header match rule specified headerdoes not exist, this header value will be treated as empty. Defaults to false.Unlike the underlying Envoy implementation this is **only** supported fornegative matches (e.g. NotContains, NotExact).",
																					MarkdownDescription: "TreatMissingAsEmpty specifies if the header match rule specified headerdoes not exist, this header value will be treated as empty. Defaults to false.Unlike the underlying Envoy implementation this is **only** supported fornegative matches (e.g. NotContains, NotExact).",
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

																	"value": schema.StringAttribute{
																		Description:         "Value defines the value of the descriptor entry.",
																		MarkdownDescription: "Value defines the value of the descriptor entry.",
																		Required:            false,
																		Optional:            true,
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

									"disabled": schema.BoolAttribute{
										Description:         "Disabled configures the HTTPProxy to not usethe default global rate limit policy defined by the Contour configuration.",
										MarkdownDescription: "Disabled configures the HTTPProxy to not usethe default global rate limit policy defined by the Contour configuration.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"domain": schema.StringAttribute{
								Description:         "Domain is passed to the Rate Limit Service.",
								MarkdownDescription: "Domain is passed to the Rate Limit Service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_resource_exhausted_code": schema.BoolAttribute{
								Description:         "EnableResourceExhaustedCode enables translating error code 429 togrpc code RESOURCE_EXHAUSTED. When disabled it's translated to UNAVAILABLE",
								MarkdownDescription: "EnableResourceExhaustedCode enables translating error code 429 togrpc code RESOURCE_EXHAUSTED. When disabled it's translated to UNAVAILABLE",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_x_rate_limit_headers": schema.BoolAttribute{
								Description:         "EnableXRateLimitHeaders defines whether to include the X-RateLimitheaders X-RateLimit-Limit, X-RateLimit-Remaining, and X-RateLimit-Reset(as defined by the IETF Internet-Draft linked below), on responsesto clients when the Rate Limit Service is consulted for a request.ref. https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html",
								MarkdownDescription: "EnableXRateLimitHeaders defines whether to include the X-RateLimitheaders X-RateLimit-Limit, X-RateLimit-Remaining, and X-RateLimit-Reset(as defined by the IETF Internet-Draft linked below), on responsesto clients when the Rate Limit Service is consulted for a request.ref. https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extension_service": schema.SingleNestedAttribute{
								Description:         "ExtensionService identifies the extension service defining the RLS.",
								MarkdownDescription: "ExtensionService identifies the extension service defining the RLS.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"fail_open": schema.BoolAttribute{
								Description:         "FailOpen defines whether to allow requests to proceed when theRate Limit Service fails to respond with a valid rate limitdecision within the timeout defined on the extension service.",
								MarkdownDescription: "FailOpen defines whether to allow requests to proceed when theRate Limit Service fails to respond with a valid rate limitdecision within the timeout defined on the extension service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tracing": schema.SingleNestedAttribute{
						Description:         "Tracing defines properties for exporting trace data to OpenTelemetry.",
						MarkdownDescription: "Tracing defines properties for exporting trace data to OpenTelemetry.",
						Attributes: map[string]schema.Attribute{
							"custom_tags": schema.ListNestedAttribute{
								Description:         "CustomTags defines a list of custom tags with unique tag name.",
								MarkdownDescription: "CustomTags defines a list of custom tags with unique tag name.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"literal": schema.StringAttribute{
											Description:         "Literal is a static custom tag value.Precisely one of Literal, RequestHeaderName must be set.",
											MarkdownDescription: "Literal is a static custom tag value.Precisely one of Literal, RequestHeaderName must be set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"request_header_name": schema.StringAttribute{
											Description:         "RequestHeaderName indicates which request headerthe label value is obtained from.Precisely one of Literal, RequestHeaderName must be set.",
											MarkdownDescription: "RequestHeaderName indicates which request headerthe label value is obtained from.Precisely one of Literal, RequestHeaderName must be set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag_name": schema.StringAttribute{
											Description:         "TagName is the unique name of the custom tag.",
											MarkdownDescription: "TagName is the unique name of the custom tag.",
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

							"extension_service": schema.SingleNestedAttribute{
								Description:         "ExtensionService identifies the extension service defining the otel-collector.",
								MarkdownDescription: "ExtensionService identifies the extension service defining the otel-collector.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"include_pod_detail": schema.BoolAttribute{
								Description:         "IncludePodDetail defines a flag.If it is true, contour will add the pod name and namespace to the span of the trace.the default is true.Note: The Envoy pods MUST have the HOSTNAME and CONTOUR_NAMESPACE environment variables set for this to work properly.",
								MarkdownDescription: "IncludePodDetail defines a flag.If it is true, contour will add the pod name and namespace to the span of the trace.the default is true.Note: The Envoy pods MUST have the HOSTNAME and CONTOUR_NAMESPACE environment variables set for this to work properly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_path_tag_length": schema.Int64Attribute{
								Description:         "MaxPathTagLength defines maximum length of the request pathto extract and include in the HttpUrl tag.contour's default is 256.",
								MarkdownDescription: "MaxPathTagLength defines maximum length of the request pathto extract and include in the HttpUrl tag.contour's default is 256.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"overall_sampling": schema.StringAttribute{
								Description:         "OverallSampling defines the sampling rate of trace data.contour's default is 100.",
								MarkdownDescription: "OverallSampling defines the sampling rate of trace data.contour's default is 100.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName defines the name for the service.contour's default is contour.",
								MarkdownDescription: "ServiceName defines the name for the service.contour's default is contour.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"xds_server": schema.SingleNestedAttribute{
						Description:         "XDSServer contains parameters for the xDS server.",
						MarkdownDescription: "XDSServer contains parameters for the xDS server.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Defines the xDS gRPC API address which Contour will serve.Contour's default is '0.0.0.0'.",
								MarkdownDescription: "Defines the xDS gRPC API address which Contour will serve.Contour's default is '0.0.0.0'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "Defines the xDS gRPC API port which Contour will serve.Contour's default is 8001.",
								MarkdownDescription: "Defines the xDS gRPC API port which Contour will serve.Contour's default is 8001.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS holds TLS file config details.Contour's default is { caFile: '/certs/ca.crt', certFile: '/certs/tls.cert', keyFile: '/certs/tls.key', insecure: false }.",
								MarkdownDescription: "TLS holds TLS file config details.Contour's default is { caFile: '/certs/ca.crt', certFile: '/certs/tls.cert', keyFile: '/certs/tls.key', insecure: false }.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "CA filename.",
										MarkdownDescription: "CA filename.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cert_file": schema.StringAttribute{
										Description:         "Client certificate filename.",
										MarkdownDescription: "Client certificate filename.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure": schema.BoolAttribute{
										Description:         "Allow serving the xDS gRPC API without TLS.",
										MarkdownDescription: "Allow serving the xDS gRPC API without TLS.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_file": schema.StringAttribute{
										Description:         "Client key filename.",
										MarkdownDescription: "Client key filename.",
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
								Description:         "Defines the XDSServer to use for 'contour serve'.Values: 'envoy' (default), 'contour (deprecated)'.Other values will produce an error.",
								MarkdownDescription: "Defines the XDSServer to use for 'contour serve'.Values: 'envoy' (default), 'contour (deprecated)'.Other values will produce an error.",
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
	}
}

func (r *ProjectcontourIoContourConfigurationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_projectcontour_io_contour_configuration_v1alpha1_manifest")

	var model ProjectcontourIoContourConfigurationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("projectcontour.io/v1alpha1")
	model.Kind = pointer.String("ContourConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
