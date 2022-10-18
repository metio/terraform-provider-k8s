/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type TraefikContainoUsMiddlewareV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TraefikContainoUsMiddlewareV1Alpha1Resource)(nil)
)

type TraefikContainoUsMiddlewareV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TraefikContainoUsMiddlewareV1Alpha1GoModel struct {
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
		AddPrefix *struct {
			Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
		} `tfsdk:"add_prefix" yaml:"addPrefix,omitempty"`

		BasicAuth *struct {
			HeaderField *string `tfsdk:"header_field" yaml:"headerField,omitempty"`

			Realm *string `tfsdk:"realm" yaml:"realm,omitempty"`

			RemoveHeader *bool `tfsdk:"remove_header" yaml:"removeHeader,omitempty"`

			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
		} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

		Buffering *struct {
			MaxRequestBodyBytes *int64 `tfsdk:"max_request_body_bytes" yaml:"maxRequestBodyBytes,omitempty"`

			MaxResponseBodyBytes *int64 `tfsdk:"max_response_body_bytes" yaml:"maxResponseBodyBytes,omitempty"`

			MemRequestBodyBytes *int64 `tfsdk:"mem_request_body_bytes" yaml:"memRequestBodyBytes,omitempty"`

			MemResponseBodyBytes *int64 `tfsdk:"mem_response_body_bytes" yaml:"memResponseBodyBytes,omitempty"`

			RetryExpression *string `tfsdk:"retry_expression" yaml:"retryExpression,omitempty"`
		} `tfsdk:"buffering" yaml:"buffering,omitempty"`

		Chain *struct {
			Middlewares *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"middlewares" yaml:"middlewares,omitempty"`
		} `tfsdk:"chain" yaml:"chain,omitempty"`

		CircuitBreaker *struct {
			CheckPeriod utilities.IntOrString `tfsdk:"check_period" yaml:"checkPeriod,omitempty"`

			Expression *string `tfsdk:"expression" yaml:"expression,omitempty"`

			FallbackDuration utilities.IntOrString `tfsdk:"fallback_duration" yaml:"fallbackDuration,omitempty"`

			RecoveryDuration utilities.IntOrString `tfsdk:"recovery_duration" yaml:"recoveryDuration,omitempty"`
		} `tfsdk:"circuit_breaker" yaml:"circuitBreaker,omitempty"`

		Compress *struct {
			ExcludedContentTypes *[]string `tfsdk:"excluded_content_types" yaml:"excludedContentTypes,omitempty"`

			MinResponseBodyBytes *int64 `tfsdk:"min_response_body_bytes" yaml:"minResponseBodyBytes,omitempty"`
		} `tfsdk:"compress" yaml:"compress,omitempty"`

		ContentType *struct {
			AutoDetect *bool `tfsdk:"auto_detect" yaml:"autoDetect,omitempty"`
		} `tfsdk:"content_type" yaml:"contentType,omitempty"`

		DigestAuth *struct {
			HeaderField *string `tfsdk:"header_field" yaml:"headerField,omitempty"`

			Realm *string `tfsdk:"realm" yaml:"realm,omitempty"`

			RemoveHeader *bool `tfsdk:"remove_header" yaml:"removeHeader,omitempty"`

			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
		} `tfsdk:"digest_auth" yaml:"digestAuth,omitempty"`

		Errors *struct {
			Query *string `tfsdk:"query" yaml:"query,omitempty"`

			Service *struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				PassHostHeader *bool `tfsdk:"pass_host_header" yaml:"passHostHeader,omitempty"`

				Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

				ResponseForwarding *struct {
					FlushInterval *string `tfsdk:"flush_interval" yaml:"flushInterval,omitempty"`
				} `tfsdk:"response_forwarding" yaml:"responseForwarding,omitempty"`

				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

				ServersTransport *string `tfsdk:"servers_transport" yaml:"serversTransport,omitempty"`

				Sticky *struct {
					Cookie *struct {
						HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						SameSite *string `tfsdk:"same_site" yaml:"sameSite,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`
					} `tfsdk:"cookie" yaml:"cookie,omitempty"`
				} `tfsdk:"sticky" yaml:"sticky,omitempty"`

				Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"service" yaml:"service,omitempty"`

			Status *[]string `tfsdk:"status" yaml:"status,omitempty"`
		} `tfsdk:"errors" yaml:"errors,omitempty"`

		ForwardAuth *struct {
			Address *string `tfsdk:"address" yaml:"address,omitempty"`

			AuthRequestHeaders *[]string `tfsdk:"auth_request_headers" yaml:"authRequestHeaders,omitempty"`

			AuthResponseHeaders *[]string `tfsdk:"auth_response_headers" yaml:"authResponseHeaders,omitempty"`

			AuthResponseHeadersRegex *string `tfsdk:"auth_response_headers_regex" yaml:"authResponseHeadersRegex,omitempty"`

			Tls *struct {
				CaOptional *bool `tfsdk:"ca_optional" yaml:"caOptional,omitempty"`

				CaSecret *string `tfsdk:"ca_secret" yaml:"caSecret,omitempty"`

				CertSecret *string `tfsdk:"cert_secret" yaml:"certSecret,omitempty"`

				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`

			TrustForwardHeader *bool `tfsdk:"trust_forward_header" yaml:"trustForwardHeader,omitempty"`
		} `tfsdk:"forward_auth" yaml:"forwardAuth,omitempty"`

		Headers *struct {
			AccessControlAllowCredentials *bool `tfsdk:"access_control_allow_credentials" yaml:"accessControlAllowCredentials,omitempty"`

			AccessControlAllowHeaders *[]string `tfsdk:"access_control_allow_headers" yaml:"accessControlAllowHeaders,omitempty"`

			AccessControlAllowMethods *[]string `tfsdk:"access_control_allow_methods" yaml:"accessControlAllowMethods,omitempty"`

			AccessControlAllowOriginList *[]string `tfsdk:"access_control_allow_origin_list" yaml:"accessControlAllowOriginList,omitempty"`

			AccessControlAllowOriginListRegex *[]string `tfsdk:"access_control_allow_origin_list_regex" yaml:"accessControlAllowOriginListRegex,omitempty"`

			AccessControlExposeHeaders *[]string `tfsdk:"access_control_expose_headers" yaml:"accessControlExposeHeaders,omitempty"`

			AccessControlMaxAge *int64 `tfsdk:"access_control_max_age" yaml:"accessControlMaxAge,omitempty"`

			AddVaryHeader *bool `tfsdk:"add_vary_header" yaml:"addVaryHeader,omitempty"`

			AllowedHosts *[]string `tfsdk:"allowed_hosts" yaml:"allowedHosts,omitempty"`

			BrowserXssFilter *bool `tfsdk:"browser_xss_filter" yaml:"browserXssFilter,omitempty"`

			ContentSecurityPolicy *string `tfsdk:"content_security_policy" yaml:"contentSecurityPolicy,omitempty"`

			ContentTypeNosniff *bool `tfsdk:"content_type_nosniff" yaml:"contentTypeNosniff,omitempty"`

			CustomBrowserXSSValue *string `tfsdk:"custom_browser_xss_value" yaml:"customBrowserXSSValue,omitempty"`

			CustomFrameOptionsValue *string `tfsdk:"custom_frame_options_value" yaml:"customFrameOptionsValue,omitempty"`

			CustomRequestHeaders *map[string]string `tfsdk:"custom_request_headers" yaml:"customRequestHeaders,omitempty"`

			CustomResponseHeaders *map[string]string `tfsdk:"custom_response_headers" yaml:"customResponseHeaders,omitempty"`

			FeaturePolicy *string `tfsdk:"feature_policy" yaml:"featurePolicy,omitempty"`

			ForceSTSHeader *bool `tfsdk:"force_sts_header" yaml:"forceSTSHeader,omitempty"`

			FrameDeny *bool `tfsdk:"frame_deny" yaml:"frameDeny,omitempty"`

			HostsProxyHeaders *[]string `tfsdk:"hosts_proxy_headers" yaml:"hostsProxyHeaders,omitempty"`

			IsDevelopment *bool `tfsdk:"is_development" yaml:"isDevelopment,omitempty"`

			PermissionsPolicy *string `tfsdk:"permissions_policy" yaml:"permissionsPolicy,omitempty"`

			PublicKey *string `tfsdk:"public_key" yaml:"publicKey,omitempty"`

			ReferrerPolicy *string `tfsdk:"referrer_policy" yaml:"referrerPolicy,omitempty"`

			SslForceHost *bool `tfsdk:"ssl_force_host" yaml:"sslForceHost,omitempty"`

			SslHost *string `tfsdk:"ssl_host" yaml:"sslHost,omitempty"`

			SslProxyHeaders *map[string]string `tfsdk:"ssl_proxy_headers" yaml:"sslProxyHeaders,omitempty"`

			SslRedirect *bool `tfsdk:"ssl_redirect" yaml:"sslRedirect,omitempty"`

			SslTemporaryRedirect *bool `tfsdk:"ssl_temporary_redirect" yaml:"sslTemporaryRedirect,omitempty"`

			StsIncludeSubdomains *bool `tfsdk:"sts_include_subdomains" yaml:"stsIncludeSubdomains,omitempty"`

			StsPreload *bool `tfsdk:"sts_preload" yaml:"stsPreload,omitempty"`

			StsSeconds *int64 `tfsdk:"sts_seconds" yaml:"stsSeconds,omitempty"`
		} `tfsdk:"headers" yaml:"headers,omitempty"`

		InFlightReq *struct {
			Amount *int64 `tfsdk:"amount" yaml:"amount,omitempty"`

			SourceCriterion *struct {
				IpStrategy *struct {
					Depth *int64 `tfsdk:"depth" yaml:"depth,omitempty"`

					ExcludedIPs *[]string `tfsdk:"excluded_i_ps" yaml:"excludedIPs,omitempty"`
				} `tfsdk:"ip_strategy" yaml:"ipStrategy,omitempty"`

				RequestHeaderName *string `tfsdk:"request_header_name" yaml:"requestHeaderName,omitempty"`

				RequestHost *bool `tfsdk:"request_host" yaml:"requestHost,omitempty"`
			} `tfsdk:"source_criterion" yaml:"sourceCriterion,omitempty"`
		} `tfsdk:"in_flight_req" yaml:"inFlightReq,omitempty"`

		IpWhiteList *struct {
			IpStrategy *struct {
				Depth *int64 `tfsdk:"depth" yaml:"depth,omitempty"`

				ExcludedIPs *[]string `tfsdk:"excluded_i_ps" yaml:"excludedIPs,omitempty"`
			} `tfsdk:"ip_strategy" yaml:"ipStrategy,omitempty"`

			SourceRange *[]string `tfsdk:"source_range" yaml:"sourceRange,omitempty"`
		} `tfsdk:"ip_white_list" yaml:"ipWhiteList,omitempty"`

		PassTLSClientCert *struct {
			Info *struct {
				Issuer *struct {
					CommonName *bool `tfsdk:"common_name" yaml:"commonName,omitempty"`

					Country *bool `tfsdk:"country" yaml:"country,omitempty"`

					DomainComponent *bool `tfsdk:"domain_component" yaml:"domainComponent,omitempty"`

					Locality *bool `tfsdk:"locality" yaml:"locality,omitempty"`

					Organization *bool `tfsdk:"organization" yaml:"organization,omitempty"`

					Province *bool `tfsdk:"province" yaml:"province,omitempty"`

					SerialNumber *bool `tfsdk:"serial_number" yaml:"serialNumber,omitempty"`
				} `tfsdk:"issuer" yaml:"issuer,omitempty"`

				NotAfter *bool `tfsdk:"not_after" yaml:"notAfter,omitempty"`

				NotBefore *bool `tfsdk:"not_before" yaml:"notBefore,omitempty"`

				Sans *bool `tfsdk:"sans" yaml:"sans,omitempty"`

				SerialNumber *bool `tfsdk:"serial_number" yaml:"serialNumber,omitempty"`

				Subject *struct {
					CommonName *bool `tfsdk:"common_name" yaml:"commonName,omitempty"`

					Country *bool `tfsdk:"country" yaml:"country,omitempty"`

					DomainComponent *bool `tfsdk:"domain_component" yaml:"domainComponent,omitempty"`

					Locality *bool `tfsdk:"locality" yaml:"locality,omitempty"`

					Organization *bool `tfsdk:"organization" yaml:"organization,omitempty"`

					OrganizationalUnit *bool `tfsdk:"organizational_unit" yaml:"organizationalUnit,omitempty"`

					Province *bool `tfsdk:"province" yaml:"province,omitempty"`

					SerialNumber *bool `tfsdk:"serial_number" yaml:"serialNumber,omitempty"`
				} `tfsdk:"subject" yaml:"subject,omitempty"`
			} `tfsdk:"info" yaml:"info,omitempty"`

			Pem *bool `tfsdk:"pem" yaml:"pem,omitempty"`
		} `tfsdk:"pass_tls_client_cert" yaml:"passTLSClientCert,omitempty"`

		Plugin *map[string]string `tfsdk:"plugin" yaml:"plugin,omitempty"`

		RateLimit *struct {
			Average *int64 `tfsdk:"average" yaml:"average,omitempty"`

			Burst *int64 `tfsdk:"burst" yaml:"burst,omitempty"`

			Period utilities.IntOrString `tfsdk:"period" yaml:"period,omitempty"`

			SourceCriterion *struct {
				IpStrategy *struct {
					Depth *int64 `tfsdk:"depth" yaml:"depth,omitempty"`

					ExcludedIPs *[]string `tfsdk:"excluded_i_ps" yaml:"excludedIPs,omitempty"`
				} `tfsdk:"ip_strategy" yaml:"ipStrategy,omitempty"`

				RequestHeaderName *string `tfsdk:"request_header_name" yaml:"requestHeaderName,omitempty"`

				RequestHost *bool `tfsdk:"request_host" yaml:"requestHost,omitempty"`
			} `tfsdk:"source_criterion" yaml:"sourceCriterion,omitempty"`
		} `tfsdk:"rate_limit" yaml:"rateLimit,omitempty"`

		RedirectRegex *struct {
			Permanent *bool `tfsdk:"permanent" yaml:"permanent,omitempty"`

			Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

			Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`
		} `tfsdk:"redirect_regex" yaml:"redirectRegex,omitempty"`

		RedirectScheme *struct {
			Permanent *bool `tfsdk:"permanent" yaml:"permanent,omitempty"`

			Port *string `tfsdk:"port" yaml:"port,omitempty"`

			Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
		} `tfsdk:"redirect_scheme" yaml:"redirectScheme,omitempty"`

		ReplacePath *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`
		} `tfsdk:"replace_path" yaml:"replacePath,omitempty"`

		ReplacePathRegex *struct {
			Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

			Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`
		} `tfsdk:"replace_path_regex" yaml:"replacePathRegex,omitempty"`

		Retry *struct {
			Attempts *int64 `tfsdk:"attempts" yaml:"attempts,omitempty"`

			InitialInterval utilities.IntOrString `tfsdk:"initial_interval" yaml:"initialInterval,omitempty"`
		} `tfsdk:"retry" yaml:"retry,omitempty"`

		StripPrefix *struct {
			ForceSlash *bool `tfsdk:"force_slash" yaml:"forceSlash,omitempty"`

			Prefixes *[]string `tfsdk:"prefixes" yaml:"prefixes,omitempty"`
		} `tfsdk:"strip_prefix" yaml:"stripPrefix,omitempty"`

		StripPrefixRegex *struct {
			Regex *[]string `tfsdk:"regex" yaml:"regex,omitempty"`
		} `tfsdk:"strip_prefix_regex" yaml:"stripPrefixRegex,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTraefikContainoUsMiddlewareV1Alpha1Resource() resource.Resource {
	return &TraefikContainoUsMiddlewareV1Alpha1Resource{}
}

func (r *TraefikContainoUsMiddlewareV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traefik_containo_us_middleware_v1alpha1"
}

func (r *TraefikContainoUsMiddlewareV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Middleware is the CRD implementation of a Traefik Middleware. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/overview/",
		MarkdownDescription: "Middleware is the CRD implementation of a Traefik Middleware. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/overview/",
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
				Description:         "MiddlewareSpec defines the desired state of a Middleware.",
				MarkdownDescription: "MiddlewareSpec defines the desired state of a Middleware.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"add_prefix": {
						Description:         "AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/addprefix/",
						MarkdownDescription: "AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/addprefix/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"prefix": {
								Description:         "Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).",
								MarkdownDescription: "Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).",

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

					"basic_auth": {
						Description:         "BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/",
						MarkdownDescription: "BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"header_field": {
								Description:         "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/#headerfield",
								MarkdownDescription: "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/#headerfield",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"realm": {
								Description:         "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",
								MarkdownDescription: "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"remove_header": {
								Description:         "RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service. Default: false.",
								MarkdownDescription: "RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service. Default: false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": {
								Description:         "Secret is the name of the referenced Kubernetes Secret containing user credentials.",
								MarkdownDescription: "Secret is the name of the referenced Kubernetes Secret containing user credentials.",

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

					"buffering": {
						Description:         "Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#maxrequestbodybytes",
						MarkdownDescription: "Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#maxrequestbodybytes",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_request_body_bytes": {
								Description:         "MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes). If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).",
								MarkdownDescription: "MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes). If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_response_body_bytes": {
								Description:         "MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes). If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).",
								MarkdownDescription: "MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes). If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mem_request_body_bytes": {
								Description:         "MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",
								MarkdownDescription: "MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mem_response_body_bytes": {
								Description:         "MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",
								MarkdownDescription: "MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory. Default: 1048576 (1Mi).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"retry_expression": {
								Description:         "RetryExpression defines the retry conditions. It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#retryexpression",
								MarkdownDescription: "RetryExpression defines the retry conditions. It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#retryexpression",

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

					"chain": {
						Description:         "Chain holds the configuration of the chain middleware. This middleware enables to define reusable combinations of other pieces of middleware. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/chain/",
						MarkdownDescription: "Chain holds the configuration of the chain middleware. This middleware enables to define reusable combinations of other pieces of middleware. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/chain/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"middlewares": {
								Description:         "Middlewares is the list of MiddlewareRef which composes the chain.",
								MarkdownDescription: "Middlewares is the list of MiddlewareRef which composes the chain.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name defines the name of the referenced Middleware resource.",
										MarkdownDescription: "Name defines the name of the referenced Middleware resource.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Middleware resource.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Middleware resource.",

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

					"circuit_breaker": {
						Description:         "CircuitBreaker holds the circuit breaker configuration.",
						MarkdownDescription: "CircuitBreaker holds the circuit breaker configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"check_period": {
								Description:         "CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).",
								MarkdownDescription: "CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"expression": {
								Description:         "Expression is the condition that triggers the tripped state.",
								MarkdownDescription: "Expression is the condition that triggers the tripped state.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fallback_duration": {
								Description:         "FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).",
								MarkdownDescription: "FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"recovery_duration": {
								Description:         "RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).",
								MarkdownDescription: "RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).",

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

					"compress": {
						Description:         "Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip compression. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/compress/",
						MarkdownDescription: "Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip compression. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/compress/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"excluded_content_types": {
								Description:         "ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing.",
								MarkdownDescription: "ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"min_response_body_bytes": {
								Description:         "MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed. Default: 1024.",
								MarkdownDescription: "MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed. Default: 1024.",

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

					"content_type": {
						Description:         "ContentType holds the content-type middleware configuration. This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.",
						MarkdownDescription: "ContentType holds the content-type middleware configuration. This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"auto_detect": {
								Description:         "AutoDetect specifies whether to let the 'Content-Type' header, if it has not been set by the backend, be automatically set to a value derived from the contents of the response. As a proxy, the default behavior should be to leave the header alone, regardless of what the backend did with it. However, the historic default was to always auto-detect and set the header if it was nil, and it is going to be kept that way in order to support users currently relying on it.",
								MarkdownDescription: "AutoDetect specifies whether to let the 'Content-Type' header, if it has not been set by the backend, be automatically set to a value derived from the contents of the response. As a proxy, the default behavior should be to leave the header alone, regardless of what the backend did with it. However, the historic default was to always auto-detect and set the header if it was nil, and it is going to be kept that way in order to support users currently relying on it.",

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

					"digest_auth": {
						Description:         "DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/digestauth/",
						MarkdownDescription: "DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/digestauth/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"header_field": {
								Description:         "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/#headerfield",
								MarkdownDescription: "HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/#headerfield",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"realm": {
								Description:         "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",
								MarkdownDescription: "Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"remove_header": {
								Description:         "RemoveHeader defines whether to remove the authorization header before forwarding the request to the backend.",
								MarkdownDescription: "RemoveHeader defines whether to remove the authorization header before forwarding the request to the backend.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": {
								Description:         "Secret is the name of the referenced Kubernetes Secret containing user credentials.",
								MarkdownDescription: "Secret is the name of the referenced Kubernetes Secret containing user credentials.",

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

					"errors": {
						Description:         "ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/errorpages/",
						MarkdownDescription: "ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/errorpages/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"query": {
								Description:         "Query defines the URL for the error page (hosted by service). The {status} variable can be used in order to insert the status code in the URL.",
								MarkdownDescription: "Query defines the URL for the error page (hosted by service). The {status} variable can be used in order to insert the status code in the URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service": {
								Description:         "Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/errorpages/#service",
								MarkdownDescription: "Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/errorpages/#service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "Kind defines the kind of the Service.",
										MarkdownDescription: "Kind defines the kind of the Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Service", "TraefikService"),
										},
									},

									"name": {
										Description:         "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",
										MarkdownDescription: "Name defines the name of the referenced Kubernetes Service or TraefikService. The differentiation between the two is specified in the Kind field.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pass_host_header": {
										Description:         "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",
										MarkdownDescription: "PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
										MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_forwarding": {
										Description:         "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",
										MarkdownDescription: "ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"flush_interval": {
												Description:         "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",
												MarkdownDescription: "FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms",

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

									"scheme": {
										Description:         "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",
										MarkdownDescription: "Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"servers_transport": {
										Description:         "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",
										MarkdownDescription: "ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sticky": {
										Description:         "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",
										MarkdownDescription: "Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cookie": {
												Description:         "Cookie defines the sticky cookie configuration.",
												MarkdownDescription: "Cookie defines the sticky cookie configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http_only": {
														Description:         "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",
														MarkdownDescription: "HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name defines the Cookie name.",
														MarkdownDescription: "Name defines the Cookie name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"same_site": {
														Description:         "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",
														MarkdownDescription: "SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secure": {
														Description:         "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",
														MarkdownDescription: "Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).",

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

									"strategy": {
										Description:         "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",
										MarkdownDescription: "Strategy defines the load balancing strategy between the servers. RoundRobin is the only supported value at the moment.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"weight": {
										Description:         "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",
										MarkdownDescription: "Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).",

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

							"status": {
								Description:         "Status defines which status or range of statuses should result in an error page. It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).",
								MarkdownDescription: "Status defines which status or range of statuses should result in an error page. It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).",

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

					"forward_auth": {
						Description:         "ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/",
						MarkdownDescription: "ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"address": {
								Description:         "Address defines the authentication server address.",
								MarkdownDescription: "Address defines the authentication server address.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auth_request_headers": {
								Description:         "AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server. If not set or empty then all request headers are passed.",
								MarkdownDescription: "AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server. If not set or empty then all request headers are passed.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auth_response_headers": {
								Description:         "AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.",
								MarkdownDescription: "AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auth_response_headers_regex": {
								Description:         "AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/#authresponseheadersregex",
								MarkdownDescription: "AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/#authresponseheadersregex",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "TLS defines the configuration used to secure the connection to the authentication server.",
								MarkdownDescription: "TLS defines the configuration used to secure the connection to the authentication server.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_optional": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ca_secret": {
										Description:         "CASecret is the name of the referenced Kubernetes Secret containing the CA to validate the server certificate. The CA certificate is extracted from key 'tls.ca' or 'ca.crt'.",
										MarkdownDescription: "CASecret is the name of the referenced Kubernetes Secret containing the CA to validate the server certificate. The CA certificate is extracted from key 'tls.ca' or 'ca.crt'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert_secret": {
										Description:         "CertSecret is the name of the referenced Kubernetes Secret containing the client certificate. The client certificate is extracted from the keys 'tls.crt' and 'tls.key'.",
										MarkdownDescription: "CertSecret is the name of the referenced Kubernetes Secret containing the client certificate. The client certificate is extracted from the keys 'tls.crt' and 'tls.key'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": {
										Description:         "InsecureSkipVerify defines whether the server certificates should be validated.",
										MarkdownDescription: "InsecureSkipVerify defines whether the server certificates should be validated.",

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

							"trust_forward_header": {
								Description:         "TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.",
								MarkdownDescription: "TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.",

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

					"headers": {
						Description:         "Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/headers/#customrequestheaders",
						MarkdownDescription: "Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/headers/#customrequestheaders",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"access_control_allow_credentials": {
								Description:         "AccessControlAllowCredentials defines whether the request can include user credentials.",
								MarkdownDescription: "AccessControlAllowCredentials defines whether the request can include user credentials.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_control_allow_headers": {
								Description:         "AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.",
								MarkdownDescription: "AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_control_allow_methods": {
								Description:         "AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.",
								MarkdownDescription: "AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_control_allow_origin_list": {
								Description:         "AccessControlAllowOriginList is a list of allowable origins. Can also be a wildcard origin '*'.",
								MarkdownDescription: "AccessControlAllowOriginList is a list of allowable origins. Can also be a wildcard origin '*'.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_control_allow_origin_list_regex": {
								Description:         "AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).",
								MarkdownDescription: "AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_control_expose_headers": {
								Description:         "AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.",
								MarkdownDescription: "AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_control_max_age": {
								Description:         "AccessControlMaxAge defines the time that a preflight request may be cached.",
								MarkdownDescription: "AccessControlMaxAge defines the time that a preflight request may be cached.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"add_vary_header": {
								Description:         "AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.",
								MarkdownDescription: "AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allowed_hosts": {
								Description:         "AllowedHosts defines the fully qualified list of allowed domain names.",
								MarkdownDescription: "AllowedHosts defines the fully qualified list of allowed domain names.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"browser_xss_filter": {
								Description:         "BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1; mode=block.",
								MarkdownDescription: "BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1; mode=block.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_security_policy": {
								Description:         "ContentSecurityPolicy defines the Content-Security-Policy header value.",
								MarkdownDescription: "ContentSecurityPolicy defines the Content-Security-Policy header value.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_type_nosniff": {
								Description:         "ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.",
								MarkdownDescription: "ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_browser_xss_value": {
								Description:         "CustomBrowserXSSValue defines the X-XSS-Protection header value. This overrides the BrowserXssFilter option.",
								MarkdownDescription: "CustomBrowserXSSValue defines the X-XSS-Protection header value. This overrides the BrowserXssFilter option.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_frame_options_value": {
								Description:         "CustomFrameOptionsValue defines the X-Frame-Options header value. This overrides the FrameDeny option.",
								MarkdownDescription: "CustomFrameOptionsValue defines the X-Frame-Options header value. This overrides the FrameDeny option.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_request_headers": {
								Description:         "CustomRequestHeaders defines the header names and values to apply to the request.",
								MarkdownDescription: "CustomRequestHeaders defines the header names and values to apply to the request.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_response_headers": {
								Description:         "CustomResponseHeaders defines the header names and values to apply to the response.",
								MarkdownDescription: "CustomResponseHeaders defines the header names and values to apply to the response.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"feature_policy": {
								Description:         "Deprecated: use PermissionsPolicy instead.",
								MarkdownDescription: "Deprecated: use PermissionsPolicy instead.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"force_sts_header": {
								Description:         "ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.",
								MarkdownDescription: "ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"frame_deny": {
								Description:         "FrameDeny defines whether to add the X-Frame-Options header with the DENY value.",
								MarkdownDescription: "FrameDeny defines whether to add the X-Frame-Options header with the DENY value.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hosts_proxy_headers": {
								Description:         "HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.",
								MarkdownDescription: "HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"is_development": {
								Description:         "IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing. Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.",
								MarkdownDescription: "IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing. Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"permissions_policy": {
								Description:         "PermissionsPolicy defines the Permissions-Policy header value. This allows sites to control browser features.",
								MarkdownDescription: "PermissionsPolicy defines the Permissions-Policy header value. This allows sites to control browser features.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"public_key": {
								Description:         "PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.",
								MarkdownDescription: "PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"referrer_policy": {
								Description:         "ReferrerPolicy defines the Referrer-Policy header value. This allows sites to control whether browsers forward the Referer header to other sites.",
								MarkdownDescription: "ReferrerPolicy defines the Referrer-Policy header value. This allows sites to control whether browsers forward the Referer header to other sites.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_force_host": {
								Description:         "Deprecated: use RedirectRegex instead.",
								MarkdownDescription: "Deprecated: use RedirectRegex instead.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_host": {
								Description:         "Deprecated: use RedirectRegex instead.",
								MarkdownDescription: "Deprecated: use RedirectRegex instead.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_proxy_headers": {
								Description:         "SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request. It can be useful when using other proxies (example: 'X-Forwarded-Proto': 'https').",
								MarkdownDescription: "SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request. It can be useful when using other proxies (example: 'X-Forwarded-Proto': 'https').",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_redirect": {
								Description:         "Deprecated: use EntryPoint redirection or RedirectScheme instead.",
								MarkdownDescription: "Deprecated: use EntryPoint redirection or RedirectScheme instead.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_temporary_redirect": {
								Description:         "Deprecated: use EntryPoint redirection or RedirectScheme instead.",
								MarkdownDescription: "Deprecated: use EntryPoint redirection or RedirectScheme instead.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sts_include_subdomains": {
								Description:         "STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.",
								MarkdownDescription: "STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sts_preload": {
								Description:         "STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.",
								MarkdownDescription: "STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sts_seconds": {
								Description:         "STSSeconds defines the max-age of the Strict-Transport-Security header. If set to 0, the header is not set.",
								MarkdownDescription: "STSSeconds defines the max-age of the Strict-Transport-Security header. If set to 0, the header is not set.",

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

					"in_flight_req": {
						Description:         "InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/inflightreq/",
						MarkdownDescription: "InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/inflightreq/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"amount": {
								Description:         "Amount defines the maximum amount of allowed simultaneous in-flight request. The middleware responds with HTTP 429 Too Many Requests if there are already amount requests in progress (based on the same sourceCriterion strategy).",
								MarkdownDescription: "Amount defines the maximum amount of allowed simultaneous in-flight request. The middleware responds with HTTP 429 Too Many Requests if there are already amount requests in progress (based on the same sourceCriterion strategy).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_criterion": {
								Description:         "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the requestHost. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/inflightreq/#sourcecriterion",
								MarkdownDescription: "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the requestHost. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/inflightreq/#sourcecriterion",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ip_strategy": {
										Description:         "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy",
										MarkdownDescription: "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"depth": {
												Description:         "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
												MarkdownDescription: "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"excluded_i_ps": {
												Description:         "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
												MarkdownDescription: "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",

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

									"request_header_name": {
										Description:         "RequestHeaderName defines the name of the header used to group incoming requests.",
										MarkdownDescription: "RequestHeaderName defines the name of the header used to group incoming requests.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_host": {
										Description:         "RequestHost defines whether to consider the request Host as the source.",
										MarkdownDescription: "RequestHost defines whether to consider the request Host as the source.",

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

					"ip_white_list": {
						Description:         "IPWhiteList holds the IP whitelist middleware configuration. This middleware accepts / refuses requests based on the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/",
						MarkdownDescription: "IPWhiteList holds the IP whitelist middleware configuration. This middleware accepts / refuses requests based on the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ip_strategy": {
								Description:         "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy",
								MarkdownDescription: "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"depth": {
										Description:         "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
										MarkdownDescription: "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"excluded_i_ps": {
										Description:         "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
										MarkdownDescription: "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",

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

							"source_range": {
								Description:         "SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).",
								MarkdownDescription: "SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).",

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

					"pass_tls_client_cert": {
						Description:         "PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/passtlsclientcert/",
						MarkdownDescription: "PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/passtlsclientcert/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"info": {
								Description:         "Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.",
								MarkdownDescription: "Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"issuer": {
										Description:         "Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.",
										MarkdownDescription: "Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"common_name": {
												Description:         "CommonName defines whether to add the organizationalUnit information into the issuer.",
												MarkdownDescription: "CommonName defines whether to add the organizationalUnit information into the issuer.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"country": {
												Description:         "Country defines whether to add the country information into the issuer.",
												MarkdownDescription: "Country defines whether to add the country information into the issuer.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"domain_component": {
												Description:         "DomainComponent defines whether to add the domainComponent information into the issuer.",
												MarkdownDescription: "DomainComponent defines whether to add the domainComponent information into the issuer.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"locality": {
												Description:         "Locality defines whether to add the locality information into the issuer.",
												MarkdownDescription: "Locality defines whether to add the locality information into the issuer.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "Organization defines whether to add the organization information into the issuer.",
												MarkdownDescription: "Organization defines whether to add the organization information into the issuer.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"province": {
												Description:         "Province defines whether to add the province information into the issuer.",
												MarkdownDescription: "Province defines whether to add the province information into the issuer.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"serial_number": {
												Description:         "SerialNumber defines whether to add the serialNumber information into the issuer.",
												MarkdownDescription: "SerialNumber defines whether to add the serialNumber information into the issuer.",

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

									"not_after": {
										Description:         "NotAfter defines whether to add the Not After information from the Validity part.",
										MarkdownDescription: "NotAfter defines whether to add the Not After information from the Validity part.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_before": {
										Description:         "NotBefore defines whether to add the Not Before information from the Validity part.",
										MarkdownDescription: "NotBefore defines whether to add the Not Before information from the Validity part.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sans": {
										Description:         "Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.",
										MarkdownDescription: "Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"serial_number": {
										Description:         "SerialNumber defines whether to add the client serialNumber information.",
										MarkdownDescription: "SerialNumber defines whether to add the client serialNumber information.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subject": {
										Description:         "Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.",
										MarkdownDescription: "Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"common_name": {
												Description:         "CommonName defines whether to add the organizationalUnit information into the subject.",
												MarkdownDescription: "CommonName defines whether to add the organizationalUnit information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"country": {
												Description:         "Country defines whether to add the country information into the subject.",
												MarkdownDescription: "Country defines whether to add the country information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"domain_component": {
												Description:         "DomainComponent defines whether to add the domainComponent information into the subject.",
												MarkdownDescription: "DomainComponent defines whether to add the domainComponent information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"locality": {
												Description:         "Locality defines whether to add the locality information into the subject.",
												MarkdownDescription: "Locality defines whether to add the locality information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organization": {
												Description:         "Organization defines whether to add the organization information into the subject.",
												MarkdownDescription: "Organization defines whether to add the organization information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"organizational_unit": {
												Description:         "OrganizationalUnit defines whether to add the organizationalUnit information into the subject.",
												MarkdownDescription: "OrganizationalUnit defines whether to add the organizationalUnit information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"province": {
												Description:         "Province defines whether to add the province information into the subject.",
												MarkdownDescription: "Province defines whether to add the province information into the subject.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"serial_number": {
												Description:         "SerialNumber defines whether to add the serialNumber information into the subject.",
												MarkdownDescription: "SerialNumber defines whether to add the serialNumber information into the subject.",

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

							"pem": {
								Description:         "PEM sets the X-Forwarded-Tls-Client-Cert header with the escaped certificate.",
								MarkdownDescription: "PEM sets the X-Forwarded-Tls-Client-Cert header with the escaped certificate.",

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

					"plugin": {
						Description:         "Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/plugins/",
						MarkdownDescription: "Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/plugins/",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rate_limit": {
						Description:         "RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ratelimit/",
						MarkdownDescription: "RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ratelimit/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"average": {
								Description:         "Average is the maximum rate, by default in requests/s, allowed for the given source. It defaults to 0, which means no rate limiting. The rate is actually defined by dividing Average by Period. So for a rate below 1req/s, one needs to define a Period larger than a second.",
								MarkdownDescription: "Average is the maximum rate, by default in requests/s, allowed for the given source. It defaults to 0, which means no rate limiting. The rate is actually defined by dividing Average by Period. So for a rate below 1req/s, one needs to define a Period larger than a second.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"burst": {
								Description:         "Burst is the maximum number of requests allowed to arrive in the same arbitrarily small period of time. It defaults to 1.",
								MarkdownDescription: "Burst is the maximum number of requests allowed to arrive in the same arbitrarily small period of time. It defaults to 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"period": {
								Description:         "Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period. It defaults to a second.",
								MarkdownDescription: "Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period. It defaults to a second.",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_criterion": {
								Description:         "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).",
								MarkdownDescription: "SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ip_strategy": {
										Description:         "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy",
										MarkdownDescription: "IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"depth": {
												Description:         "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",
												MarkdownDescription: "Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"excluded_i_ps": {
												Description:         "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",
												MarkdownDescription: "ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.",

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

									"request_header_name": {
										Description:         "RequestHeaderName defines the name of the header used to group incoming requests.",
										MarkdownDescription: "RequestHeaderName defines the name of the header used to group incoming requests.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_host": {
										Description:         "RequestHost defines whether to consider the request Host as the source.",
										MarkdownDescription: "RequestHost defines whether to consider the request Host as the source.",

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

					"redirect_regex": {
						Description:         "RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectregex/#regex",
						MarkdownDescription: "RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectregex/#regex",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"permanent": {
								Description:         "Permanent defines whether the redirection is permanent (301).",
								MarkdownDescription: "Permanent defines whether the redirection is permanent (301).",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"regex": {
								Description:         "Regex defines the regex used to match and capture elements from the request URL.",
								MarkdownDescription: "Regex defines the regex used to match and capture elements from the request URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replacement": {
								Description:         "Replacement defines how to modify the URL to have the new target URL.",
								MarkdownDescription: "Replacement defines how to modify the URL to have the new target URL.",

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

					"redirect_scheme": {
						Description:         "RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectscheme/",
						MarkdownDescription: "RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectscheme/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"permanent": {
								Description:         "Permanent defines whether the redirection is permanent (301).",
								MarkdownDescription: "Permanent defines whether the redirection is permanent (301).",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "Port defines the port of the new URL.",
								MarkdownDescription: "Port defines the port of the new URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scheme": {
								Description:         "Scheme defines the scheme of the new URL.",
								MarkdownDescription: "Scheme defines the scheme of the new URL.",

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

					"replace_path": {
						Description:         "ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/replacepath/",
						MarkdownDescription: "ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/replacepath/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "Path defines the path to use as replacement in the request URL.",
								MarkdownDescription: "Path defines the path to use as replacement in the request URL.",

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

					"replace_path_regex": {
						Description:         "ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/replacepathregex/",
						MarkdownDescription: "ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/replacepathregex/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"regex": {
								Description:         "Regex defines the regular expression used to match and capture the path from the request URL.",
								MarkdownDescription: "Regex defines the regular expression used to match and capture the path from the request URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replacement": {
								Description:         "Replacement defines the replacement path format, which can include captured variables.",
								MarkdownDescription: "Replacement defines the replacement path format, which can include captured variables.",

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

					"retry": {
						Description:         "Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/retry/",
						MarkdownDescription: "Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/retry/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"attempts": {
								Description:         "Attempts defines how many times the request should be retried.",
								MarkdownDescription: "Attempts defines how many times the request should be retried.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_interval": {
								Description:         "InitialInterval defines the first wait time in the exponential backoff series. The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.",
								MarkdownDescription: "InitialInterval defines the first wait time in the exponential backoff series. The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.",

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

					"strip_prefix": {
						Description:         "StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefix/",
						MarkdownDescription: "StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefix/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"force_slash": {
								Description:         "ForceSlash ensures that the resulting stripped path is not the empty string, by replacing it with / when necessary. Default: true.",
								MarkdownDescription: "ForceSlash ensures that the resulting stripped path is not the empty string, by replacing it with / when necessary. Default: true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prefixes": {
								Description:         "Prefixes defines the prefixes to strip from the request URL.",
								MarkdownDescription: "Prefixes defines the prefixes to strip from the request URL.",

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

					"strip_prefix_regex": {
						Description:         "StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefixregex/",
						MarkdownDescription: "StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefixregex/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"regex": {
								Description:         "Regex defines the regular expression to match the path prefix from the request URL.",
								MarkdownDescription: "Regex defines the regular expression to match the path prefix from the request URL.",

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

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *TraefikContainoUsMiddlewareV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_traefik_containo_us_middleware_v1alpha1")

	var state TraefikContainoUsMiddlewareV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsMiddlewareV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("Middleware")

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

func (r *TraefikContainoUsMiddlewareV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_containo_us_middleware_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TraefikContainoUsMiddlewareV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_traefik_containo_us_middleware_v1alpha1")

	var state TraefikContainoUsMiddlewareV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsMiddlewareV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("Middleware")

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

func (r *TraefikContainoUsMiddlewareV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_traefik_containo_us_middleware_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
