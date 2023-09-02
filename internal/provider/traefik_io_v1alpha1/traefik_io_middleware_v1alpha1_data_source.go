/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

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
	_ datasource.DataSource              = &TraefikIoMiddlewareV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &TraefikIoMiddlewareV1Alpha1DataSource{}
)

func NewTraefikIoMiddlewareV1Alpha1DataSource() datasource.DataSource {
	return &TraefikIoMiddlewareV1Alpha1DataSource{}
}

type TraefikIoMiddlewareV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TraefikIoMiddlewareV1Alpha1DataSourceData struct {
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
		AddPrefix *struct {
			Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
		} `tfsdk:"add_prefix" json:"addPrefix,omitempty"`
		BasicAuth *struct {
			HeaderField  *string `tfsdk:"header_field" json:"headerField,omitempty"`
			Realm        *string `tfsdk:"realm" json:"realm,omitempty"`
			RemoveHeader *bool   `tfsdk:"remove_header" json:"removeHeader,omitempty"`
			Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
		Buffering *struct {
			MaxRequestBodyBytes  *int64  `tfsdk:"max_request_body_bytes" json:"maxRequestBodyBytes,omitempty"`
			MaxResponseBodyBytes *int64  `tfsdk:"max_response_body_bytes" json:"maxResponseBodyBytes,omitempty"`
			MemRequestBodyBytes  *int64  `tfsdk:"mem_request_body_bytes" json:"memRequestBodyBytes,omitempty"`
			MemResponseBodyBytes *int64  `tfsdk:"mem_response_body_bytes" json:"memResponseBodyBytes,omitempty"`
			RetryExpression      *string `tfsdk:"retry_expression" json:"retryExpression,omitempty"`
		} `tfsdk:"buffering" json:"buffering,omitempty"`
		Chain *struct {
			Middlewares *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"middlewares" json:"middlewares,omitempty"`
		} `tfsdk:"chain" json:"chain,omitempty"`
		CircuitBreaker *struct {
			CheckPeriod      *string `tfsdk:"check_period" json:"checkPeriod,omitempty"`
			Expression       *string `tfsdk:"expression" json:"expression,omitempty"`
			FallbackDuration *string `tfsdk:"fallback_duration" json:"fallbackDuration,omitempty"`
			RecoveryDuration *string `tfsdk:"recovery_duration" json:"recoveryDuration,omitempty"`
		} `tfsdk:"circuit_breaker" json:"circuitBreaker,omitempty"`
		Compress *struct {
			ExcludedContentTypes *[]string `tfsdk:"excluded_content_types" json:"excludedContentTypes,omitempty"`
			MinResponseBodyBytes *int64    `tfsdk:"min_response_body_bytes" json:"minResponseBodyBytes,omitempty"`
		} `tfsdk:"compress" json:"compress,omitempty"`
		ContentType *map[string]string `tfsdk:"content_type" json:"contentType,omitempty"`
		DigestAuth  *struct {
			HeaderField  *string `tfsdk:"header_field" json:"headerField,omitempty"`
			Realm        *string `tfsdk:"realm" json:"realm,omitempty"`
			RemoveHeader *bool   `tfsdk:"remove_header" json:"removeHeader,omitempty"`
			Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"digest_auth" json:"digestAuth,omitempty"`
		Errors *struct {
			Query   *string `tfsdk:"query" json:"query,omitempty"`
			Service *struct {
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				NativeLB           *bool   `tfsdk:"native_lb" json:"nativeLB,omitempty"`
				PassHostHeader     *bool   `tfsdk:"pass_host_header" json:"passHostHeader,omitempty"`
				Port               *string `tfsdk:"port" json:"port,omitempty"`
				ResponseForwarding *struct {
					FlushInterval *string `tfsdk:"flush_interval" json:"flushInterval,omitempty"`
				} `tfsdk:"response_forwarding" json:"responseForwarding,omitempty"`
				Scheme           *string `tfsdk:"scheme" json:"scheme,omitempty"`
				ServersTransport *string `tfsdk:"servers_transport" json:"serversTransport,omitempty"`
				Sticky           *struct {
					Cookie *struct {
						HttpOnly *bool   `tfsdk:"http_only" json:"httpOnly,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
						Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
					} `tfsdk:"cookie" json:"cookie,omitempty"`
				} `tfsdk:"sticky" json:"sticky,omitempty"`
				Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
				Weight   *int64  `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Status *[]string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"errors" json:"errors,omitempty"`
		ForwardAuth *struct {
			Address                  *string   `tfsdk:"address" json:"address,omitempty"`
			AuthRequestHeaders       *[]string `tfsdk:"auth_request_headers" json:"authRequestHeaders,omitempty"`
			AuthResponseHeaders      *[]string `tfsdk:"auth_response_headers" json:"authResponseHeaders,omitempty"`
			AuthResponseHeadersRegex *string   `tfsdk:"auth_response_headers_regex" json:"authResponseHeadersRegex,omitempty"`
			Tls                      *struct {
				CaSecret           *string `tfsdk:"ca_secret" json:"caSecret,omitempty"`
				CertSecret         *string `tfsdk:"cert_secret" json:"certSecret,omitempty"`
				InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			TrustForwardHeader *bool `tfsdk:"trust_forward_header" json:"trustForwardHeader,omitempty"`
		} `tfsdk:"forward_auth" json:"forwardAuth,omitempty"`
		GrpcWeb *struct {
			AllowOrigins *[]string `tfsdk:"allow_origins" json:"allowOrigins,omitempty"`
		} `tfsdk:"grpc_web" json:"grpcWeb,omitempty"`
		Headers *struct {
			AccessControlAllowCredentials     *bool              `tfsdk:"access_control_allow_credentials" json:"accessControlAllowCredentials,omitempty"`
			AccessControlAllowHeaders         *[]string          `tfsdk:"access_control_allow_headers" json:"accessControlAllowHeaders,omitempty"`
			AccessControlAllowMethods         *[]string          `tfsdk:"access_control_allow_methods" json:"accessControlAllowMethods,omitempty"`
			AccessControlAllowOriginList      *[]string          `tfsdk:"access_control_allow_origin_list" json:"accessControlAllowOriginList,omitempty"`
			AccessControlAllowOriginListRegex *[]string          `tfsdk:"access_control_allow_origin_list_regex" json:"accessControlAllowOriginListRegex,omitempty"`
			AccessControlExposeHeaders        *[]string          `tfsdk:"access_control_expose_headers" json:"accessControlExposeHeaders,omitempty"`
			AccessControlMaxAge               *int64             `tfsdk:"access_control_max_age" json:"accessControlMaxAge,omitempty"`
			AddVaryHeader                     *bool              `tfsdk:"add_vary_header" json:"addVaryHeader,omitempty"`
			AllowedHosts                      *[]string          `tfsdk:"allowed_hosts" json:"allowedHosts,omitempty"`
			BrowserXssFilter                  *bool              `tfsdk:"browser_xss_filter" json:"browserXssFilter,omitempty"`
			ContentSecurityPolicy             *string            `tfsdk:"content_security_policy" json:"contentSecurityPolicy,omitempty"`
			ContentTypeNosniff                *bool              `tfsdk:"content_type_nosniff" json:"contentTypeNosniff,omitempty"`
			CustomBrowserXSSValue             *string            `tfsdk:"custom_browser_xss_value" json:"customBrowserXSSValue,omitempty"`
			CustomFrameOptionsValue           *string            `tfsdk:"custom_frame_options_value" json:"customFrameOptionsValue,omitempty"`
			CustomRequestHeaders              *map[string]string `tfsdk:"custom_request_headers" json:"customRequestHeaders,omitempty"`
			CustomResponseHeaders             *map[string]string `tfsdk:"custom_response_headers" json:"customResponseHeaders,omitempty"`
			ForceSTSHeader                    *bool              `tfsdk:"force_sts_header" json:"forceSTSHeader,omitempty"`
			FrameDeny                         *bool              `tfsdk:"frame_deny" json:"frameDeny,omitempty"`
			HostsProxyHeaders                 *[]string          `tfsdk:"hosts_proxy_headers" json:"hostsProxyHeaders,omitempty"`
			IsDevelopment                     *bool              `tfsdk:"is_development" json:"isDevelopment,omitempty"`
			PermissionsPolicy                 *string            `tfsdk:"permissions_policy" json:"permissionsPolicy,omitempty"`
			PublicKey                         *string            `tfsdk:"public_key" json:"publicKey,omitempty"`
			ReferrerPolicy                    *string            `tfsdk:"referrer_policy" json:"referrerPolicy,omitempty"`
			SslProxyHeaders                   *map[string]string `tfsdk:"ssl_proxy_headers" json:"sslProxyHeaders,omitempty"`
			StsIncludeSubdomains              *bool              `tfsdk:"sts_include_subdomains" json:"stsIncludeSubdomains,omitempty"`
			StsPreload                        *bool              `tfsdk:"sts_preload" json:"stsPreload,omitempty"`
			StsSeconds                        *int64             `tfsdk:"sts_seconds" json:"stsSeconds,omitempty"`
		} `tfsdk:"headers" json:"headers,omitempty"`
		InFlightReq *struct {
			Amount          *int64 `tfsdk:"amount" json:"amount,omitempty"`
			SourceCriterion *struct {
				IpStrategy *struct {
					Depth       *int64    `tfsdk:"depth" json:"depth,omitempty"`
					ExcludedIPs *[]string `tfsdk:"excluded_i_ps" json:"excludedIPs,omitempty"`
				} `tfsdk:"ip_strategy" json:"ipStrategy,omitempty"`
				RequestHeaderName *string `tfsdk:"request_header_name" json:"requestHeaderName,omitempty"`
				RequestHost       *bool   `tfsdk:"request_host" json:"requestHost,omitempty"`
			} `tfsdk:"source_criterion" json:"sourceCriterion,omitempty"`
		} `tfsdk:"in_flight_req" json:"inFlightReq,omitempty"`
		IpAllowList *struct {
			IpStrategy *struct {
				Depth       *int64    `tfsdk:"depth" json:"depth,omitempty"`
				ExcludedIPs *[]string `tfsdk:"excluded_i_ps" json:"excludedIPs,omitempty"`
			} `tfsdk:"ip_strategy" json:"ipStrategy,omitempty"`
			SourceRange *[]string `tfsdk:"source_range" json:"sourceRange,omitempty"`
		} `tfsdk:"ip_allow_list" json:"ipAllowList,omitempty"`
		PassTLSClientCert *struct {
			Info *struct {
				Issuer *struct {
					CommonName      *bool `tfsdk:"common_name" json:"commonName,omitempty"`
					Country         *bool `tfsdk:"country" json:"country,omitempty"`
					DomainComponent *bool `tfsdk:"domain_component" json:"domainComponent,omitempty"`
					Locality        *bool `tfsdk:"locality" json:"locality,omitempty"`
					Organization    *bool `tfsdk:"organization" json:"organization,omitempty"`
					Province        *bool `tfsdk:"province" json:"province,omitempty"`
					SerialNumber    *bool `tfsdk:"serial_number" json:"serialNumber,omitempty"`
				} `tfsdk:"issuer" json:"issuer,omitempty"`
				NotAfter     *bool `tfsdk:"not_after" json:"notAfter,omitempty"`
				NotBefore    *bool `tfsdk:"not_before" json:"notBefore,omitempty"`
				Sans         *bool `tfsdk:"sans" json:"sans,omitempty"`
				SerialNumber *bool `tfsdk:"serial_number" json:"serialNumber,omitempty"`
				Subject      *struct {
					CommonName         *bool `tfsdk:"common_name" json:"commonName,omitempty"`
					Country            *bool `tfsdk:"country" json:"country,omitempty"`
					DomainComponent    *bool `tfsdk:"domain_component" json:"domainComponent,omitempty"`
					Locality           *bool `tfsdk:"locality" json:"locality,omitempty"`
					Organization       *bool `tfsdk:"organization" json:"organization,omitempty"`
					OrganizationalUnit *bool `tfsdk:"organizational_unit" json:"organizationalUnit,omitempty"`
					Province           *bool `tfsdk:"province" json:"province,omitempty"`
					SerialNumber       *bool `tfsdk:"serial_number" json:"serialNumber,omitempty"`
				} `tfsdk:"subject" json:"subject,omitempty"`
			} `tfsdk:"info" json:"info,omitempty"`
			Pem *bool `tfsdk:"pem" json:"pem,omitempty"`
		} `tfsdk:"pass_tls_client_cert" json:"passTLSClientCert,omitempty"`
		Plugin    *map[string]string `tfsdk:"plugin" json:"plugin,omitempty"`
		RateLimit *struct {
			Average         *int64  `tfsdk:"average" json:"average,omitempty"`
			Burst           *int64  `tfsdk:"burst" json:"burst,omitempty"`
			Period          *string `tfsdk:"period" json:"period,omitempty"`
			SourceCriterion *struct {
				IpStrategy *struct {
					Depth       *int64    `tfsdk:"depth" json:"depth,omitempty"`
					ExcludedIPs *[]string `tfsdk:"excluded_i_ps" json:"excludedIPs,omitempty"`
				} `tfsdk:"ip_strategy" json:"ipStrategy,omitempty"`
				RequestHeaderName *string `tfsdk:"request_header_name" json:"requestHeaderName,omitempty"`
				RequestHost       *bool   `tfsdk:"request_host" json:"requestHost,omitempty"`
			} `tfsdk:"source_criterion" json:"sourceCriterion,omitempty"`
		} `tfsdk:"rate_limit" json:"rateLimit,omitempty"`
		RedirectRegex *struct {
			Permanent   *bool   `tfsdk:"permanent" json:"permanent,omitempty"`
			Regex       *string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement *string `tfsdk:"replacement" json:"replacement,omitempty"`
		} `tfsdk:"redirect_regex" json:"redirectRegex,omitempty"`
		RedirectScheme *struct {
			Permanent *bool   `tfsdk:"permanent" json:"permanent,omitempty"`
			Port      *string `tfsdk:"port" json:"port,omitempty"`
			Scheme    *string `tfsdk:"scheme" json:"scheme,omitempty"`
		} `tfsdk:"redirect_scheme" json:"redirectScheme,omitempty"`
		ReplacePath *struct {
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"replace_path" json:"replacePath,omitempty"`
		ReplacePathRegex *struct {
			Regex       *string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement *string `tfsdk:"replacement" json:"replacement,omitempty"`
		} `tfsdk:"replace_path_regex" json:"replacePathRegex,omitempty"`
		Retry *struct {
			Attempts        *int64  `tfsdk:"attempts" json:"attempts,omitempty"`
			InitialInterval *string `tfsdk:"initial_interval" json:"initialInterval,omitempty"`
		} `tfsdk:"retry" json:"retry,omitempty"`
		StripPrefix *struct {
			Prefixes *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
		} `tfsdk:"strip_prefix" json:"stripPrefix,omitempty"`
		StripPrefixRegex *struct {
			Regex *[]string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"strip_prefix_regex" json:"stripPrefixRegex,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoMiddlewareV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_middleware_v1alpha1"
}

func (r *TraefikIoMiddlewareV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Middleware is the CRD implementation of a Traefik Middleware. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/overview/",
		MarkdownDescription: "Middleware is the CRD implementation of a Traefik Middleware. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/overview/",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "MiddlewareSpec defines the desired state of a Middleware.",
				MarkdownDescription: "MiddlewareSpec defines the desired state of a Middleware.",
				Attributes: map[string]schema.Attribute{
					"add_prefix": schema.SingleNestedAttribute{
						Description:         "AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/addprefix/",
						MarkdownDescription: "AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/addprefix/",
						Attributes: map[string]schema.Attribute{
							"prefix": schema.StringAttribute{
								Description:         "Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).",
								MarkdownDescription: "Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"basic_auth": schema.SingleNestedAttribute{
						Description:         "BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/basicauth/",
						MarkdownDescription: "BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/basicauth/",
						Attributes: map[string]schema.Attribute{
							"header_field": schema.StringAttribute{
								Description:         "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/basicauth/#headerfield",
								MarkdownDescription: "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/basicauth/#headerfield",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"realm": schema.StringAttribute{
								Description:         "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",
								MarkdownDescription: "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"remove_header": schema.BoolAttribute{
								Description:         "RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service. Default: false.",
								MarkdownDescription: "RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service. Default: false.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret": schema.StringAttribute{
								Description:         "Secret is the name of the referenced Kubernetes Secret containing user credentials.",
								MarkdownDescription: "Secret is the name of the referenced Kubernetes Secret containing user credentials.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"buffering": schema.SingleNestedAttribute{
						Description:         "Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/buffering/#maxrequestbodybytes",
						MarkdownDescription: "Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/buffering/#maxrequestbodybytes",
						Attributes: map[string]schema.Attribute{
							"max_request_body_bytes": schema.Int64Attribute{
								Description:         "MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes). If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).",
								MarkdownDescription: "MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes). If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_response_body_bytes": schema.Int64Attribute{
								Description:         "MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes). If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).",
								MarkdownDescription: "MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes). If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mem_request_body_bytes": schema.Int64Attribute{
								Description:         "MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",
								MarkdownDescription: "MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mem_response_body_bytes": schema.Int64Attribute{
								Description:         "MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",
								MarkdownDescription: "MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"retry_expression": schema.StringAttribute{
								Description:         "RetryExpression defines the retry conditions. It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/buffering/#retryexpression",
								MarkdownDescription: "RetryExpression defines the retry conditions. It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/buffering/#retryexpression",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"chain": schema.SingleNestedAttribute{
						Description:         "Chain holds the configuration of the chain middleware. This middleware enables to define reusable combinations of other pieces of middleware. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/chain/",
						MarkdownDescription: "Chain holds the configuration of the chain middleware. This middleware enables to define reusable combinations of other pieces of middleware. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/chain/",
						Attributes: map[string]schema.Attribute{
							"middlewares": schema.ListNestedAttribute{
								Description:         "Middlewares is the list of MiddlewareRef which composes the chain.",
								MarkdownDescription: "Middlewares is the list of MiddlewareRef which composes the chain.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name defines the name of the referenced Middleware resource.",
											MarkdownDescription: "Name defines the name of the referenced Middleware resource.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace defines the namespace of the referenced Middleware resource.",
											MarkdownDescription: "Namespace defines the namespace of the referenced Middleware resource.",
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

					"circuit_breaker": schema.SingleNestedAttribute{
						Description:         "CircuitBreaker holds the circuit breaker configuration.",
						MarkdownDescription: "CircuitBreaker holds the circuit breaker configuration.",
						Attributes: map[string]schema.Attribute{
							"check_period": schema.StringAttribute{
								Description:         "CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).",
								MarkdownDescription: "CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"expression": schema.StringAttribute{
								Description:         "Expression is the condition that triggers the tripped state.",
								MarkdownDescription: "Expression is the condition that triggers the tripped state.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"fallback_duration": schema.StringAttribute{
								Description:         "FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).",
								MarkdownDescription: "FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"recovery_duration": schema.StringAttribute{
								Description:         "RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).",
								MarkdownDescription: "RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"compress": schema.SingleNestedAttribute{
						Description:         "Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip compression. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/compress/",
						MarkdownDescription: "Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip compression. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/compress/",
						Attributes: map[string]schema.Attribute{
							"excluded_content_types": schema.ListAttribute{
								Description:         "ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing. 'application/grpc' is always excluded.",
								MarkdownDescription: "ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing. 'application/grpc' is always excluded.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_response_body_bytes": schema.Int64Attribute{
								Description:         "MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed. Default: 1024.",
								MarkdownDescription: "MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed. Default: 1024.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"content_type": schema.MapAttribute{
						Description:         "ContentType holds the content-type middleware configuration. This middleware sets the 'Content-Type' header value to the media type detected from the response content, when it is not set by the backend.",
						MarkdownDescription: "ContentType holds the content-type middleware configuration. This middleware sets the 'Content-Type' header value to the media type detected from the response content, when it is not set by the backend.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"digest_auth": schema.SingleNestedAttribute{
						Description:         "DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/digestauth/",
						MarkdownDescription: "DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/digestauth/",
						Attributes: map[string]schema.Attribute{
							"header_field": schema.StringAttribute{
								Description:         "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/basicauth/#headerfield",
								MarkdownDescription: "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/basicauth/#headerfield",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"realm": schema.StringAttribute{
								Description:         "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",
								MarkdownDescription: "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"remove_header": schema.BoolAttribute{
								Description:         "RemoveHeader defines whether to remove the authorization header before forwarding the request to the backend.",
								MarkdownDescription: "RemoveHeader defines whether to remove the authorization header before forwarding the request to the backend.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret": schema.StringAttribute{
								Description:         "Secret is the name of the referenced Kubernetes Secret containing user credentials.",
								MarkdownDescription: "Secret is the name of the referenced Kubernetes Secret containing user credentials.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"errors": schema.SingleNestedAttribute{
						Description:         "ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/errorpages/",
						MarkdownDescription: "ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/errorpages/",
						Attributes: map[string]schema.Attribute{
							"query": schema.StringAttribute{
								Description:         "Query defines the URL for the error page (hosted by service). The {status} variable can be used in order to insert the status code in the URL.",
								MarkdownDescription: "Query defines the URL for the error page (hosted by service). The {status} variable can be used in order to insert the status code in the URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"service": schema.SingleNestedAttribute{
								Description:         "Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/errorpages/#service",
								MarkdownDescription: "Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/errorpages/#service",
								Attributes: map[string]schema.Attribute{
									"kind": schema.StringAttribute{
										Description:         "Kind defines the kind of the Service.",
										MarkdownDescription: "Kind defines the kind of the Service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
										MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"native_lb": schema.BoolAttribute{
										Description:         "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
										MarkdownDescription: "NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pass_host_header": schema.BoolAttribute{
										Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
										MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.StringAttribute{
										Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
										MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"response_forwarding": schema.SingleNestedAttribute{
										Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
										MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
										Attributes: map[string]schema.Attribute{
											"flush_interval": schema.StringAttribute{
												Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
												MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"scheme": schema.StringAttribute{
										Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
										MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"servers_transport": schema.StringAttribute{
										Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
										MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"sticky": schema.SingleNestedAttribute{
										Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/services/#sticky-sessions",
										MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/services/#sticky-sessions",
										Attributes: map[string]schema.Attribute{
											"cookie": schema.SingleNestedAttribute{
												Description:         "Cookie defines the sticky cookie configuration.",
												MarkdownDescription: "Cookie defines the sticky cookie configuration.",
												Attributes: map[string]schema.Attribute{
													"http_only": schema.BoolAttribute{
														Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
														MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name defines the Cookie name.",
														MarkdownDescription: "Name defines the Cookie name.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"same_site": schema.StringAttribute{
														Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
														MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secure": schema.BoolAttribute{
														Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
														MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
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

									"strategy": schema.StringAttribute{
										Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
										MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"weight": schema.Int64Attribute{
										Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
										MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"status": schema.ListAttribute{
								Description:         "Status defines which status or range of statuses should result in an error page. It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).",
								MarkdownDescription: "Status defines which status or range of statuses should result in an error page. It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).",
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

					"forward_auth": schema.SingleNestedAttribute{
						Description:         "ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/forwardauth/",
						MarkdownDescription: "ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/forwardauth/",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Address defines the authentication server address.",
								MarkdownDescription: "Address defines the authentication server address.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"auth_request_headers": schema.ListAttribute{
								Description:         "AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server. If not set or empty then all request headers are passed.",
								MarkdownDescription: "AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server. If not set or empty then all request headers are passed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"auth_response_headers": schema.ListAttribute{
								Description:         "AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.",
								MarkdownDescription: "AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"auth_response_headers_regex": schema.StringAttribute{
								Description:         "AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/forwardauth/#authresponseheadersregex",
								MarkdownDescription: "AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/forwardauth/#authresponseheadersregex",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS defines the configuration used to secure the connection to the authentication server.",
								MarkdownDescription: "TLS defines the configuration used to secure the connection to the authentication server.",
								Attributes: map[string]schema.Attribute{
									"ca_secret": schema.StringAttribute{
										Description:         "CASecret is the name of the referenced Kubernetes Secret containing the CA to validate the server certificate. The CA certificate is extracted from key 'tls.ca' or 'ca.crt'.",
										MarkdownDescription: "CASecret is the name of the referenced Kubernetes Secret containing the CA to validate the server certificate. The CA certificate is extracted from key 'tls.ca' or 'ca.crt'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cert_secret": schema.StringAttribute{
										Description:         "CertSecret is the name of the referenced Kubernetes Secret containing the client certificate. The client certificate is extracted from the keys 'tls.crt' and 'tls.key'.",
										MarkdownDescription: "CertSecret is the name of the referenced Kubernetes Secret containing the client certificate. The client certificate is extracted from the keys 'tls.crt' and 'tls.key'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "InsecureSkipVerify defines whether the server certificates should be validated.",
										MarkdownDescription: "InsecureSkipVerify defines whether the server certificates should be validated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"trust_forward_header": schema.BoolAttribute{
								Description:         "TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.",
								MarkdownDescription: "TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.",
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
						Description:         "GrpcWeb holds the gRPC web middleware configuration. This middleware converts a gRPC web request to an HTTP/2 gRPC request.",
						MarkdownDescription: "GrpcWeb holds the gRPC web middleware configuration. This middleware converts a gRPC web request to an HTTP/2 gRPC request.",
						Attributes: map[string]schema.Attribute{
							"allow_origins": schema.ListAttribute{
								Description:         "AllowOrigins is a list of allowable origins. Can also be a wildcard origin '*'.",
								MarkdownDescription: "AllowOrigins is a list of allowable origins. Can also be a wildcard origin '*'.",
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

					"headers": schema.SingleNestedAttribute{
						Description:         "Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/headers/#customrequestheaders",
						MarkdownDescription: "Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/headers/#customrequestheaders",
						Attributes: map[string]schema.Attribute{
							"access_control_allow_credentials": schema.BoolAttribute{
								Description:         "AccessControlAllowCredentials defines whether the request can include user credentials.",
								MarkdownDescription: "AccessControlAllowCredentials defines whether the request can include user credentials.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"access_control_allow_headers": schema.ListAttribute{
								Description:         "AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.",
								MarkdownDescription: "AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"access_control_allow_methods": schema.ListAttribute{
								Description:         "AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.",
								MarkdownDescription: "AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"access_control_allow_origin_list": schema.ListAttribute{
								Description:         "AccessControlAllowOriginList is a list of allowable origins. Can also be a wildcard origin '*'.",
								MarkdownDescription: "AccessControlAllowOriginList is a list of allowable origins. Can also be a wildcard origin '*'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"access_control_allow_origin_list_regex": schema.ListAttribute{
								Description:         "AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).",
								MarkdownDescription: "AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"access_control_expose_headers": schema.ListAttribute{
								Description:         "AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.",
								MarkdownDescription: "AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"access_control_max_age": schema.Int64Attribute{
								Description:         "AccessControlMaxAge defines the time that a preflight request may be cached.",
								MarkdownDescription: "AccessControlMaxAge defines the time that a preflight request may be cached.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"add_vary_header": schema.BoolAttribute{
								Description:         "AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.",
								MarkdownDescription: "AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"allowed_hosts": schema.ListAttribute{
								Description:         "AllowedHosts defines the fully qualified list of allowed domain names.",
								MarkdownDescription: "AllowedHosts defines the fully qualified list of allowed domain names.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"browser_xss_filter": schema.BoolAttribute{
								Description:         "BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1; mode=block.",
								MarkdownDescription: "BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1; mode=block.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"content_security_policy": schema.StringAttribute{
								Description:         "ContentSecurityPolicy defines the Content-Security-Policy header value.",
								MarkdownDescription: "ContentSecurityPolicy defines the Content-Security-Policy header value.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"content_type_nosniff": schema.BoolAttribute{
								Description:         "ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.",
								MarkdownDescription: "ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"custom_browser_xss_value": schema.StringAttribute{
								Description:         "CustomBrowserXSSValue defines the X-XSS-Protection header value. This overrides the BrowserXssFilter option.",
								MarkdownDescription: "CustomBrowserXSSValue defines the X-XSS-Protection header value. This overrides the BrowserXssFilter option.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"custom_frame_options_value": schema.StringAttribute{
								Description:         "CustomFrameOptionsValue defines the X-Frame-Options header value. This overrides the FrameDeny option.",
								MarkdownDescription: "CustomFrameOptionsValue defines the X-Frame-Options header value. This overrides the FrameDeny option.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"custom_request_headers": schema.MapAttribute{
								Description:         "CustomRequestHeaders defines the header names and values to apply to the request.",
								MarkdownDescription: "CustomRequestHeaders defines the header names and values to apply to the request.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"custom_response_headers": schema.MapAttribute{
								Description:         "CustomResponseHeaders defines the header names and values to apply to the response.",
								MarkdownDescription: "CustomResponseHeaders defines the header names and values to apply to the response.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"force_sts_header": schema.BoolAttribute{
								Description:         "ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.",
								MarkdownDescription: "ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"frame_deny": schema.BoolAttribute{
								Description:         "FrameDeny defines whether to add the X-Frame-Options header with the DENY value.",
								MarkdownDescription: "FrameDeny defines whether to add the X-Frame-Options header with the DENY value.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"hosts_proxy_headers": schema.ListAttribute{
								Description:         "HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.",
								MarkdownDescription: "HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"is_development": schema.BoolAttribute{
								Description:         "IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing. Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.",
								MarkdownDescription: "IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing. Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"permissions_policy": schema.StringAttribute{
								Description:         "PermissionsPolicy defines the Permissions-Policy header value. This allows sites to control browser features.",
								MarkdownDescription: "PermissionsPolicy defines the Permissions-Policy header value. This allows sites to control browser features.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"public_key": schema.StringAttribute{
								Description:         "PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.",
								MarkdownDescription: "PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"referrer_policy": schema.StringAttribute{
								Description:         "ReferrerPolicy defines the Referrer-Policy header value. This allows sites to control whether browsers forward the Referer header to other sites.",
								MarkdownDescription: "ReferrerPolicy defines the Referrer-Policy header value. This allows sites to control whether browsers forward the Referer header to other sites.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ssl_proxy_headers": schema.MapAttribute{
								Description:         "SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request. It can be useful when using other proxies (example: 'X-Forwarded-Proto': 'https').",
								MarkdownDescription: "SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request. It can be useful when using other proxies (example: 'X-Forwarded-Proto': 'https').",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sts_include_subdomains": schema.BoolAttribute{
								Description:         "STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.",
								MarkdownDescription: "STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sts_preload": schema.BoolAttribute{
								Description:         "STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.",
								MarkdownDescription: "STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sts_seconds": schema.Int64Attribute{
								Description:         "STSSeconds defines the max-age of the Strict-Transport-Security header. If set to 0, the header is not set.",
								MarkdownDescription: "STSSeconds defines the max-age of the Strict-Transport-Security header. If set to 0, the header is not set.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"in_flight_req": schema.SingleNestedAttribute{
						Description:         "InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/inflightreq/",
						MarkdownDescription: "InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/inflightreq/",
						Attributes: map[string]schema.Attribute{
							"amount": schema.Int64Attribute{
								Description:         "Amount defines the maximum amount of allowed simultaneous in-flight request. The middleware responds with HTTP 429 Too Many Requests if there are already amount requests in progress (based on the same sourceCriterion strategy).",
								MarkdownDescription: "Amount defines the maximum amount of allowed simultaneous in-flight request. The middleware responds with HTTP 429 Too Many Requests if there are already amount requests in progress (based on the same sourceCriterion strategy).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"source_criterion": schema.SingleNestedAttribute{
								Description:         "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the requestHost. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/inflightreq/#sourcecriterion",
								MarkdownDescription: "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the requestHost. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/inflightreq/#sourcecriterion",
								Attributes: map[string]schema.Attribute{
									"ip_strategy": schema.SingleNestedAttribute{
										Description:         "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/#ipstrategy",
										MarkdownDescription: "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/#ipstrategy",
										Attributes: map[string]schema.Attribute{
											"depth": schema.Int64Attribute{
												Description:         "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
												MarkdownDescription: "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"excluded_i_ps": schema.ListAttribute{
												Description:         "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
												MarkdownDescription: "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
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

									"request_header_name": schema.StringAttribute{
										Description:         "RequestHeaderName defines the name of the header used to group incoming requests.",
										MarkdownDescription: "RequestHeaderName defines the name of the header used to group incoming requests.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_host": schema.BoolAttribute{
										Description:         "RequestHost defines whether to consider the request Host as the source.",
										MarkdownDescription: "RequestHost defines whether to consider the request Host as the source.",
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

					"ip_allow_list": schema.SingleNestedAttribute{
						Description:         "IPAllowList holds the IP allowlist middleware configuration. This middleware accepts / refuses requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/",
						MarkdownDescription: "IPAllowList holds the IP allowlist middleware configuration. This middleware accepts / refuses requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/",
						Attributes: map[string]schema.Attribute{
							"ip_strategy": schema.SingleNestedAttribute{
								Description:         "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/#ipstrategy",
								MarkdownDescription: "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/#ipstrategy",
								Attributes: map[string]schema.Attribute{
									"depth": schema.Int64Attribute{
										Description:         "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
										MarkdownDescription: "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"excluded_i_ps": schema.ListAttribute{
										Description:         "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
										MarkdownDescription: "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
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

							"source_range": schema.ListAttribute{
								Description:         "SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).",
								MarkdownDescription: "SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).",
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

					"pass_tls_client_cert": schema.SingleNestedAttribute{
						Description:         "PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/passtlsclientcert/",
						MarkdownDescription: "PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/passtlsclientcert/",
						Attributes: map[string]schema.Attribute{
							"info": schema.SingleNestedAttribute{
								Description:         "Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.",
								MarkdownDescription: "Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.",
								Attributes: map[string]schema.Attribute{
									"issuer": schema.SingleNestedAttribute{
										Description:         "Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.",
										MarkdownDescription: "Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.",
										Attributes: map[string]schema.Attribute{
											"common_name": schema.BoolAttribute{
												Description:         "CommonName defines whether to add the organizationalUnit information into the issuer.",
												MarkdownDescription: "CommonName defines whether to add the organizationalUnit information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"country": schema.BoolAttribute{
												Description:         "Country defines whether to add the country information into the issuer.",
												MarkdownDescription: "Country defines whether to add the country information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"domain_component": schema.BoolAttribute{
												Description:         "DomainComponent defines whether to add the domainComponent information into the issuer.",
												MarkdownDescription: "DomainComponent defines whether to add the domainComponent information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"locality": schema.BoolAttribute{
												Description:         "Locality defines whether to add the locality information into the issuer.",
												MarkdownDescription: "Locality defines whether to add the locality information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"organization": schema.BoolAttribute{
												Description:         "Organization defines whether to add the organization information into the issuer.",
												MarkdownDescription: "Organization defines whether to add the organization information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"province": schema.BoolAttribute{
												Description:         "Province defines whether to add the province information into the issuer.",
												MarkdownDescription: "Province defines whether to add the province information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"serial_number": schema.BoolAttribute{
												Description:         "SerialNumber defines whether to add the serialNumber information into the issuer.",
												MarkdownDescription: "SerialNumber defines whether to add the serialNumber information into the issuer.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"not_after": schema.BoolAttribute{
										Description:         "NotAfter defines whether to add the Not After information from the Validity part.",
										MarkdownDescription: "NotAfter defines whether to add the Not After information from the Validity part.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"not_before": schema.BoolAttribute{
										Description:         "NotBefore defines whether to add the Not Before information from the Validity part.",
										MarkdownDescription: "NotBefore defines whether to add the Not Before information from the Validity part.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"sans": schema.BoolAttribute{
										Description:         "Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.",
										MarkdownDescription: "Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"serial_number": schema.BoolAttribute{
										Description:         "SerialNumber defines whether to add the client serialNumber information.",
										MarkdownDescription: "SerialNumber defines whether to add the client serialNumber information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"subject": schema.SingleNestedAttribute{
										Description:         "Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.",
										MarkdownDescription: "Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.",
										Attributes: map[string]schema.Attribute{
											"common_name": schema.BoolAttribute{
												Description:         "CommonName defines whether to add the organizationalUnit information into the subject.",
												MarkdownDescription: "CommonName defines whether to add the organizationalUnit information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"country": schema.BoolAttribute{
												Description:         "Country defines whether to add the country information into the subject.",
												MarkdownDescription: "Country defines whether to add the country information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"domain_component": schema.BoolAttribute{
												Description:         "DomainComponent defines whether to add the domainComponent information into the subject.",
												MarkdownDescription: "DomainComponent defines whether to add the domainComponent information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"locality": schema.BoolAttribute{
												Description:         "Locality defines whether to add the locality information into the subject.",
												MarkdownDescription: "Locality defines whether to add the locality information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"organization": schema.BoolAttribute{
												Description:         "Organization defines whether to add the organization information into the subject.",
												MarkdownDescription: "Organization defines whether to add the organization information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"organizational_unit": schema.BoolAttribute{
												Description:         "OrganizationalUnit defines whether to add the organizationalUnit information into the subject.",
												MarkdownDescription: "OrganizationalUnit defines whether to add the organizationalUnit information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"province": schema.BoolAttribute{
												Description:         "Province defines whether to add the province information into the subject.",
												MarkdownDescription: "Province defines whether to add the province information into the subject.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"serial_number": schema.BoolAttribute{
												Description:         "SerialNumber defines whether to add the serialNumber information into the subject.",
												MarkdownDescription: "SerialNumber defines whether to add the serialNumber information into the subject.",
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

							"pem": schema.BoolAttribute{
								Description:         "PEM sets the X-Forwarded-Tls-Client-Cert header with the certificate.",
								MarkdownDescription: "PEM sets the X-Forwarded-Tls-Client-Cert header with the certificate.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"plugin": schema.MapAttribute{
						Description:         "Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/plugins/",
						MarkdownDescription: "Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/plugins/",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"rate_limit": schema.SingleNestedAttribute{
						Description:         "RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ratelimit/",
						MarkdownDescription: "RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ratelimit/",
						Attributes: map[string]schema.Attribute{
							"average": schema.Int64Attribute{
								Description:         "Average is the maximum rate, by default in requests/s, allowed for the given source. It defaults to 0, which means no rate limiting. The rate is actually defined by dividing Average by Period. So for a rate below 1req/s, one needs to define a Period larger than a second.",
								MarkdownDescription: "Average is the maximum rate, by default in requests/s, allowed for the given source. It defaults to 0, which means no rate limiting. The rate is actually defined by dividing Average by Period. So for a rate below 1req/s, one needs to define a Period larger than a second.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"burst": schema.Int64Attribute{
								Description:         "Burst is the maximum number of requests allowed to arrive in the same arbitrarily small period of time. It defaults to 1.",
								MarkdownDescription: "Burst is the maximum number of requests allowed to arrive in the same arbitrarily small period of time. It defaults to 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"period": schema.StringAttribute{
								Description:         "Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period. It defaults to a second.",
								MarkdownDescription: "Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period. It defaults to a second.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"source_criterion": schema.SingleNestedAttribute{
								Description:         "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).",
								MarkdownDescription: "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).",
								Attributes: map[string]schema.Attribute{
									"ip_strategy": schema.SingleNestedAttribute{
										Description:         "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/#ipstrategy",
										MarkdownDescription: "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/ipallowlist/#ipstrategy",
										Attributes: map[string]schema.Attribute{
											"depth": schema.Int64Attribute{
												Description:         "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
												MarkdownDescription: "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"excluded_i_ps": schema.ListAttribute{
												Description:         "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
												MarkdownDescription: "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
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

									"request_header_name": schema.StringAttribute{
										Description:         "RequestHeaderName defines the name of the header used to group incoming requests.",
										MarkdownDescription: "RequestHeaderName defines the name of the header used to group incoming requests.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_host": schema.BoolAttribute{
										Description:         "RequestHost defines whether to consider the request Host as the source.",
										MarkdownDescription: "RequestHost defines whether to consider the request Host as the source.",
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

					"redirect_regex": schema.SingleNestedAttribute{
						Description:         "RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/redirectregex/#regex",
						MarkdownDescription: "RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/redirectregex/#regex",
						Attributes: map[string]schema.Attribute{
							"permanent": schema.BoolAttribute{
								Description:         "Permanent defines whether the redirection is permanent (301).",
								MarkdownDescription: "Permanent defines whether the redirection is permanent (301).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"regex": schema.StringAttribute{
								Description:         "Regex defines the regex used to match and capture elements from the request URL.",
								MarkdownDescription: "Regex defines the regex used to match and capture elements from the request URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replacement": schema.StringAttribute{
								Description:         "Replacement defines how to modify the URL to have the new target URL.",
								MarkdownDescription: "Replacement defines how to modify the URL to have the new target URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"redirect_scheme": schema.SingleNestedAttribute{
						Description:         "RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/redirectscheme/",
						MarkdownDescription: "RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/redirectscheme/",
						Attributes: map[string]schema.Attribute{
							"permanent": schema.BoolAttribute{
								Description:         "Permanent defines whether the redirection is permanent (301).",
								MarkdownDescription: "Permanent defines whether the redirection is permanent (301).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.StringAttribute{
								Description:         "Port defines the port of the new URL.",
								MarkdownDescription: "Port defines the port of the new URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"scheme": schema.StringAttribute{
								Description:         "Scheme defines the scheme of the new URL.",
								MarkdownDescription: "Scheme defines the scheme of the new URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"replace_path": schema.SingleNestedAttribute{
						Description:         "ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/replacepath/",
						MarkdownDescription: "ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/replacepath/",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "Path defines the path to use as replacement in the request URL.",
								MarkdownDescription: "Path defines the path to use as replacement in the request URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"replace_path_regex": schema.SingleNestedAttribute{
						Description:         "ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/replacepathregex/",
						MarkdownDescription: "ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/replacepathregex/",
						Attributes: map[string]schema.Attribute{
							"regex": schema.StringAttribute{
								Description:         "Regex defines the regular expression used to match and capture the path from the request URL.",
								MarkdownDescription: "Regex defines the regular expression used to match and capture the path from the request URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replacement": schema.StringAttribute{
								Description:         "Replacement defines the replacement path format, which can include captured variables.",
								MarkdownDescription: "Replacement defines the replacement path format, which can include captured variables.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"retry": schema.SingleNestedAttribute{
						Description:         "Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/retry/",
						MarkdownDescription: "Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/retry/",
						Attributes: map[string]schema.Attribute{
							"attempts": schema.Int64Attribute{
								Description:         "Attempts defines how many times the request should be retried.",
								MarkdownDescription: "Attempts defines how many times the request should be retried.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"initial_interval": schema.StringAttribute{
								Description:         "InitialInterval defines the first wait time in the exponential backoff series. The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.",
								MarkdownDescription: "InitialInterval defines the first wait time in the exponential backoff series. The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"strip_prefix": schema.SingleNestedAttribute{
						Description:         "StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/stripprefix/",
						MarkdownDescription: "StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/stripprefix/",
						Attributes: map[string]schema.Attribute{
							"prefixes": schema.ListAttribute{
								Description:         "Prefixes defines the prefixes to strip from the request URL.",
								MarkdownDescription: "Prefixes defines the prefixes to strip from the request URL.",
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

					"strip_prefix_regex": schema.SingleNestedAttribute{
						Description:         "StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/stripprefixregex/",
						MarkdownDescription: "StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.0/middlewares/http/stripprefixregex/",
						Attributes: map[string]schema.Attribute{
							"regex": schema.ListAttribute{
								Description:         "Regex defines the regular expression to match the path prefix from the request URL.",
								MarkdownDescription: "Regex defines the regular expression to match the path prefix from the request URL.",
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

func (r *TraefikIoMiddlewareV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *TraefikIoMiddlewareV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_traefik_io_middleware_v1alpha1")

	var data TraefikIoMiddlewareV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "traefik.io", Version: "v1alpha1", Resource: "Middleware"}).
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

	var readResponse TraefikIoMiddlewareV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("traefik.io/v1alpha1")
	data.Kind = pointer.String("Middleware")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
