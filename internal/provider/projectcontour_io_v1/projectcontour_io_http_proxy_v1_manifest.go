/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package projectcontour_io_v1

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
	_ datasource.DataSource = &ProjectcontourIoHttpproxyV1Manifest{}
)

func NewProjectcontourIoHttpproxyV1Manifest() datasource.DataSource {
	return &ProjectcontourIoHttpproxyV1Manifest{}
}

type ProjectcontourIoHttpproxyV1Manifest struct{}

type ProjectcontourIoHttpproxyV1ManifestData struct {
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
		Includes *[]struct {
			Conditions *[]struct {
				Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
				Header *struct {
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
				} `tfsdk:"header" json:"header,omitempty"`
				Prefix         *string `tfsdk:"prefix" json:"prefix,omitempty"`
				QueryParameter *struct {
					Contains   *string `tfsdk:"contains" json:"contains,omitempty"`
					Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
					IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Present    *bool   `tfsdk:"present" json:"present,omitempty"`
					Regex      *string `tfsdk:"regex" json:"regex,omitempty"`
					Suffix     *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"query_parameter" json:"queryParameter,omitempty"`
				Regex *string `tfsdk:"regex" json:"regex,omitempty"`
			} `tfsdk:"conditions" json:"conditions,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"includes" json:"includes,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		Routes           *[]struct {
			AuthPolicy *struct {
				Context  *map[string]string `tfsdk:"context" json:"context,omitempty"`
				Disabled *bool              `tfsdk:"disabled" json:"disabled,omitempty"`
			} `tfsdk:"auth_policy" json:"authPolicy,omitempty"`
			Conditions *[]struct {
				Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
				Header *struct {
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
				} `tfsdk:"header" json:"header,omitempty"`
				Prefix         *string `tfsdk:"prefix" json:"prefix,omitempty"`
				QueryParameter *struct {
					Contains   *string `tfsdk:"contains" json:"contains,omitempty"`
					Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
					IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Present    *bool   `tfsdk:"present" json:"present,omitempty"`
					Regex      *string `tfsdk:"regex" json:"regex,omitempty"`
					Suffix     *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"query_parameter" json:"queryParameter,omitempty"`
				Regex *string `tfsdk:"regex" json:"regex,omitempty"`
			} `tfsdk:"conditions" json:"conditions,omitempty"`
			CookieRewritePolicies *[]struct {
				DomainRewrite *struct {
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"domain_rewrite" json:"domainRewrite,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				PathRewrite *struct {
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"path_rewrite" json:"pathRewrite,omitempty"`
				SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
				Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
			} `tfsdk:"cookie_rewrite_policies" json:"cookieRewritePolicies,omitempty"`
			DirectResponsePolicy *struct {
				Body       *string `tfsdk:"body" json:"body,omitempty"`
				StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
			} `tfsdk:"direct_response_policy" json:"directResponsePolicy,omitempty"`
			EnableWebsockets  *bool `tfsdk:"enable_websockets" json:"enableWebsockets,omitempty"`
			HealthCheckPolicy *struct {
				ExpectedStatuses *[]struct {
					End   *int64 `tfsdk:"end" json:"end,omitempty"`
					Start *int64 `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"expected_statuses" json:"expectedStatuses,omitempty"`
				HealthyThresholdCount   *int64  `tfsdk:"healthy_threshold_count" json:"healthyThresholdCount,omitempty"`
				Host                    *string `tfsdk:"host" json:"host,omitempty"`
				IntervalSeconds         *int64  `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
				Path                    *string `tfsdk:"path" json:"path,omitempty"`
				TimeoutSeconds          *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				UnhealthyThresholdCount *int64  `tfsdk:"unhealthy_threshold_count" json:"unhealthyThresholdCount,omitempty"`
			} `tfsdk:"health_check_policy" json:"healthCheckPolicy,omitempty"`
			InternalRedirectPolicy *struct {
				AllowCrossSchemeRedirect  *string   `tfsdk:"allow_cross_scheme_redirect" json:"allowCrossSchemeRedirect,omitempty"`
				DenyRepeatedRouteRedirect *bool     `tfsdk:"deny_repeated_route_redirect" json:"denyRepeatedRouteRedirect,omitempty"`
				MaxInternalRedirects      *int64    `tfsdk:"max_internal_redirects" json:"maxInternalRedirects,omitempty"`
				RedirectResponseCodes     *[]string `tfsdk:"redirect_response_codes" json:"redirectResponseCodes,omitempty"`
			} `tfsdk:"internal_redirect_policy" json:"internalRedirectPolicy,omitempty"`
			IpAllowPolicy *[]struct {
				Cidr   *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"ip_allow_policy" json:"ipAllowPolicy,omitempty"`
			IpDenyPolicy *[]struct {
				Cidr   *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"ip_deny_policy" json:"ipDenyPolicy,omitempty"`
			JwtVerificationPolicy *struct {
				Disabled *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
				Require  *string `tfsdk:"require" json:"require,omitempty"`
			} `tfsdk:"jwt_verification_policy" json:"jwtVerificationPolicy,omitempty"`
			LoadBalancerPolicy *struct {
				RequestHashPolicies *[]struct {
					HashSourceIP      *bool `tfsdk:"hash_source_ip" json:"hashSourceIP,omitempty"`
					HeaderHashOptions *struct {
						HeaderName *string `tfsdk:"header_name" json:"headerName,omitempty"`
					} `tfsdk:"header_hash_options" json:"headerHashOptions,omitempty"`
					QueryParameterHashOptions *struct {
						ParameterName *string `tfsdk:"parameter_name" json:"parameterName,omitempty"`
					} `tfsdk:"query_parameter_hash_options" json:"queryParameterHashOptions,omitempty"`
					Terminal *bool `tfsdk:"terminal" json:"terminal,omitempty"`
				} `tfsdk:"request_hash_policies" json:"requestHashPolicies,omitempty"`
				Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"load_balancer_policy" json:"loadBalancerPolicy,omitempty"`
			PathRewritePolicy *struct {
				ReplacePrefix *[]struct {
					Prefix      *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Replacement *string `tfsdk:"replacement" json:"replacement,omitempty"`
				} `tfsdk:"replace_prefix" json:"replacePrefix,omitempty"`
			} `tfsdk:"path_rewrite_policy" json:"pathRewritePolicy,omitempty"`
			PermitInsecure  *bool `tfsdk:"permit_insecure" json:"permitInsecure,omitempty"`
			RateLimitPolicy *struct {
				Global *struct {
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
				} `tfsdk:"global" json:"global,omitempty"`
				Local *struct {
					Burst                *int64 `tfsdk:"burst" json:"burst,omitempty"`
					Requests             *int64 `tfsdk:"requests" json:"requests,omitempty"`
					ResponseHeadersToAdd *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"response_headers_to_add" json:"responseHeadersToAdd,omitempty"`
					ResponseStatusCode *int64  `tfsdk:"response_status_code" json:"responseStatusCode,omitempty"`
					Unit               *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"local" json:"local,omitempty"`
			} `tfsdk:"rate_limit_policy" json:"rateLimitPolicy,omitempty"`
			RequestHeadersPolicy *struct {
				Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
				Set    *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"set" json:"set,omitempty"`
			} `tfsdk:"request_headers_policy" json:"requestHeadersPolicy,omitempty"`
			RequestRedirectPolicy *struct {
				Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
				Path       *string `tfsdk:"path" json:"path,omitempty"`
				Port       *int64  `tfsdk:"port" json:"port,omitempty"`
				Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Scheme     *string `tfsdk:"scheme" json:"scheme,omitempty"`
				StatusCode *int64  `tfsdk:"status_code" json:"statusCode,omitempty"`
			} `tfsdk:"request_redirect_policy" json:"requestRedirectPolicy,omitempty"`
			ResponseHeadersPolicy *struct {
				Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
				Set    *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"set" json:"set,omitempty"`
			} `tfsdk:"response_headers_policy" json:"responseHeadersPolicy,omitempty"`
			RetryPolicy *struct {
				Count                *int64    `tfsdk:"count" json:"count,omitempty"`
				PerTryTimeout        *string   `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
				RetriableStatusCodes *[]string `tfsdk:"retriable_status_codes" json:"retriableStatusCodes,omitempty"`
				RetryOn              *[]string `tfsdk:"retry_on" json:"retryOn,omitempty"`
			} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
			Services *[]struct {
				CookieRewritePolicies *[]struct {
					DomainRewrite *struct {
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"domain_rewrite" json:"domainRewrite,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					PathRewrite *struct {
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"path_rewrite" json:"pathRewrite,omitempty"`
					SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
					Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
				} `tfsdk:"cookie_rewrite_policies" json:"cookieRewritePolicies,omitempty"`
				HealthPort           *int64  `tfsdk:"health_port" json:"healthPort,omitempty"`
				Mirror               *bool   `tfsdk:"mirror" json:"mirror,omitempty"`
				Name                 *string `tfsdk:"name" json:"name,omitempty"`
				Port                 *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol             *string `tfsdk:"protocol" json:"protocol,omitempty"`
				RequestHeadersPolicy *struct {
					Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
					Set    *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"request_headers_policy" json:"requestHeadersPolicy,omitempty"`
				ResponseHeadersPolicy *struct {
					Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
					Set    *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"response_headers_policy" json:"responseHeadersPolicy,omitempty"`
				SlowStartPolicy *struct {
					Aggression       *string `tfsdk:"aggression" json:"aggression,omitempty"`
					MinWeightPercent *int64  `tfsdk:"min_weight_percent" json:"minWeightPercent,omitempty"`
					Window           *string `tfsdk:"window" json:"window,omitempty"`
				} `tfsdk:"slow_start_policy" json:"slowStartPolicy,omitempty"`
				Validation *struct {
					CaSecret     *string   `tfsdk:"ca_secret" json:"caSecret,omitempty"`
					SubjectName  *string   `tfsdk:"subject_name" json:"subjectName,omitempty"`
					SubjectNames *[]string `tfsdk:"subject_names" json:"subjectNames,omitempty"`
				} `tfsdk:"validation" json:"validation,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			TimeoutPolicy *struct {
				Idle           *string `tfsdk:"idle" json:"idle,omitempty"`
				IdleConnection *string `tfsdk:"idle_connection" json:"idleConnection,omitempty"`
				Response       *string `tfsdk:"response" json:"response,omitempty"`
			} `tfsdk:"timeout_policy" json:"timeoutPolicy,omitempty"`
		} `tfsdk:"routes" json:"routes,omitempty"`
		Tcpproxy *struct {
			HealthCheckPolicy *struct {
				HealthyThresholdCount   *int64 `tfsdk:"healthy_threshold_count" json:"healthyThresholdCount,omitempty"`
				IntervalSeconds         *int64 `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
				TimeoutSeconds          *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				UnhealthyThresholdCount *int64 `tfsdk:"unhealthy_threshold_count" json:"unhealthyThresholdCount,omitempty"`
			} `tfsdk:"health_check_policy" json:"healthCheckPolicy,omitempty"`
			Include *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"include" json:"include,omitempty"`
			Includes *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"includes" json:"includes,omitempty"`
			LoadBalancerPolicy *struct {
				RequestHashPolicies *[]struct {
					HashSourceIP      *bool `tfsdk:"hash_source_ip" json:"hashSourceIP,omitempty"`
					HeaderHashOptions *struct {
						HeaderName *string `tfsdk:"header_name" json:"headerName,omitempty"`
					} `tfsdk:"header_hash_options" json:"headerHashOptions,omitempty"`
					QueryParameterHashOptions *struct {
						ParameterName *string `tfsdk:"parameter_name" json:"parameterName,omitempty"`
					} `tfsdk:"query_parameter_hash_options" json:"queryParameterHashOptions,omitempty"`
					Terminal *bool `tfsdk:"terminal" json:"terminal,omitempty"`
				} `tfsdk:"request_hash_policies" json:"requestHashPolicies,omitempty"`
				Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"load_balancer_policy" json:"loadBalancerPolicy,omitempty"`
			Services *[]struct {
				CookieRewritePolicies *[]struct {
					DomainRewrite *struct {
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"domain_rewrite" json:"domainRewrite,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					PathRewrite *struct {
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"path_rewrite" json:"pathRewrite,omitempty"`
					SameSite *string `tfsdk:"same_site" json:"sameSite,omitempty"`
					Secure   *bool   `tfsdk:"secure" json:"secure,omitempty"`
				} `tfsdk:"cookie_rewrite_policies" json:"cookieRewritePolicies,omitempty"`
				HealthPort           *int64  `tfsdk:"health_port" json:"healthPort,omitempty"`
				Mirror               *bool   `tfsdk:"mirror" json:"mirror,omitempty"`
				Name                 *string `tfsdk:"name" json:"name,omitempty"`
				Port                 *int64  `tfsdk:"port" json:"port,omitempty"`
				Protocol             *string `tfsdk:"protocol" json:"protocol,omitempty"`
				RequestHeadersPolicy *struct {
					Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
					Set    *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"request_headers_policy" json:"requestHeadersPolicy,omitempty"`
				ResponseHeadersPolicy *struct {
					Remove *[]string `tfsdk:"remove" json:"remove,omitempty"`
					Set    *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"response_headers_policy" json:"responseHeadersPolicy,omitempty"`
				SlowStartPolicy *struct {
					Aggression       *string `tfsdk:"aggression" json:"aggression,omitempty"`
					MinWeightPercent *int64  `tfsdk:"min_weight_percent" json:"minWeightPercent,omitempty"`
					Window           *string `tfsdk:"window" json:"window,omitempty"`
				} `tfsdk:"slow_start_policy" json:"slowStartPolicy,omitempty"`
				Validation *struct {
					CaSecret     *string   `tfsdk:"ca_secret" json:"caSecret,omitempty"`
					SubjectName  *string   `tfsdk:"subject_name" json:"subjectName,omitempty"`
					SubjectNames *[]string `tfsdk:"subject_names" json:"subjectNames,omitempty"`
				} `tfsdk:"validation" json:"validation,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
		} `tfsdk:"tcpproxy" json:"tcpproxy,omitempty"`
		Virtualhost *struct {
			Authorization *struct {
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
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			CorsPolicy *struct {
				AllowCredentials    *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
				AllowHeaders        *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
				AllowMethods        *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
				AllowOrigin         *[]string `tfsdk:"allow_origin" json:"allowOrigin,omitempty"`
				AllowPrivateNetwork *bool     `tfsdk:"allow_private_network" json:"allowPrivateNetwork,omitempty"`
				ExposeHeaders       *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
				MaxAge              *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
			} `tfsdk:"cors_policy" json:"corsPolicy,omitempty"`
			Fqdn          *string `tfsdk:"fqdn" json:"fqdn,omitempty"`
			IpAllowPolicy *[]struct {
				Cidr   *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"ip_allow_policy" json:"ipAllowPolicy,omitempty"`
			IpDenyPolicy *[]struct {
				Cidr   *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"ip_deny_policy" json:"ipDenyPolicy,omitempty"`
			JwtProviders *[]struct {
				Audiences  *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
				Default    *bool     `tfsdk:"default" json:"default,omitempty"`
				ForwardJWT *bool     `tfsdk:"forward_jwt" json:"forwardJWT,omitempty"`
				Issuer     *string   `tfsdk:"issuer" json:"issuer,omitempty"`
				Name       *string   `tfsdk:"name" json:"name,omitempty"`
				RemoteJWKS *struct {
					CacheDuration   *string `tfsdk:"cache_duration" json:"cacheDuration,omitempty"`
					DnsLookupFamily *string `tfsdk:"dns_lookup_family" json:"dnsLookupFamily,omitempty"`
					Timeout         *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Uri             *string `tfsdk:"uri" json:"uri,omitempty"`
					Validation      *struct {
						CaSecret     *string   `tfsdk:"ca_secret" json:"caSecret,omitempty"`
						SubjectName  *string   `tfsdk:"subject_name" json:"subjectName,omitempty"`
						SubjectNames *[]string `tfsdk:"subject_names" json:"subjectNames,omitempty"`
					} `tfsdk:"validation" json:"validation,omitempty"`
				} `tfsdk:"remote_jwks" json:"remoteJWKS,omitempty"`
			} `tfsdk:"jwt_providers" json:"jwtProviders,omitempty"`
			RateLimitPolicy *struct {
				Global *struct {
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
				} `tfsdk:"global" json:"global,omitempty"`
				Local *struct {
					Burst                *int64 `tfsdk:"burst" json:"burst,omitempty"`
					Requests             *int64 `tfsdk:"requests" json:"requests,omitempty"`
					ResponseHeadersToAdd *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"response_headers_to_add" json:"responseHeadersToAdd,omitempty"`
					ResponseStatusCode *int64  `tfsdk:"response_status_code" json:"responseStatusCode,omitempty"`
					Unit               *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"local" json:"local,omitempty"`
			} `tfsdk:"rate_limit_policy" json:"rateLimitPolicy,omitempty"`
			Tls *struct {
				ClientValidation *struct {
					CaSecret                 *string `tfsdk:"ca_secret" json:"caSecret,omitempty"`
					CrlOnlyVerifyLeafCert    *bool   `tfsdk:"crl_only_verify_leaf_cert" json:"crlOnlyVerifyLeafCert,omitempty"`
					CrlSecret                *string `tfsdk:"crl_secret" json:"crlSecret,omitempty"`
					ForwardClientCertificate *struct {
						Cert    *bool `tfsdk:"cert" json:"cert,omitempty"`
						Chain   *bool `tfsdk:"chain" json:"chain,omitempty"`
						Dns     *bool `tfsdk:"dns" json:"dns,omitempty"`
						Subject *bool `tfsdk:"subject" json:"subject,omitempty"`
						Uri     *bool `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"forward_client_certificate" json:"forwardClientCertificate,omitempty"`
					OptionalClientCertificate *bool `tfsdk:"optional_client_certificate" json:"optionalClientCertificate,omitempty"`
					SkipClientCertValidation  *bool `tfsdk:"skip_client_cert_validation" json:"skipClientCertValidation,omitempty"`
				} `tfsdk:"client_validation" json:"clientValidation,omitempty"`
				EnableFallbackCertificate *bool   `tfsdk:"enable_fallback_certificate" json:"enableFallbackCertificate,omitempty"`
				MaximumProtocolVersion    *string `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
				MinimumProtocolVersion    *string `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
				Passthrough               *bool   `tfsdk:"passthrough" json:"passthrough,omitempty"`
				SecretName                *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"virtualhost" json:"virtualhost,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ProjectcontourIoHttpproxyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_projectcontour_io_http_proxy_v1_manifest"
}

func (r *ProjectcontourIoHttpproxyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HTTPProxy is an Ingress CRD specification.",
		MarkdownDescription: "HTTPProxy is an Ingress CRD specification.",
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
				Description:         "HTTPProxySpec defines the spec of the CRD.",
				MarkdownDescription: "HTTPProxySpec defines the spec of the CRD.",
				Attributes: map[string]schema.Attribute{
					"includes": schema.ListNestedAttribute{
						Description:         "Includes allow for specific routing configuration to be included from another HTTPProxy, possibly in another namespace.",
						MarkdownDescription: "Includes allow for specific routing configuration to be included from another HTTPProxy, possibly in another namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conditions": schema.ListNestedAttribute{
									Description:         "Conditions are a set of rules that are applied to included HTTPProxies. In effect, they are added onto the Conditions of included HTTPProxy Route structs. When applied, they are merged using AND, with one exception: There can be only one Prefix MatchCondition per Conditions slice. More than one Prefix, or contradictory Conditions, will make the include invalid. Exact and Regex match conditions are not allowed on includes.",
									MarkdownDescription: "Conditions are a set of rules that are applied to included HTTPProxies. In effect, they are added onto the Conditions of included HTTPProxy Route structs. When applied, they are merged using AND, with one exception: There can be only one Prefix MatchCondition per Conditions slice. More than one Prefix, or contradictory Conditions, will make the include invalid. Exact and Regex match conditions are not allowed on includes.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exact": schema.StringAttribute{
												Description:         "Exact defines a exact match for a request. This field is not allowed in include match conditions.",
												MarkdownDescription: "Exact defines a exact match for a request. This field is not allowed in include match conditions.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"header": schema.SingleNestedAttribute{
												Description:         "Header specifies the header condition to match.",
												MarkdownDescription: "Header specifies the header condition to match.",
												Attributes: map[string]schema.Attribute{
													"contains": schema.StringAttribute{
														Description:         "Contains specifies a substring that must be present in the header value.",
														MarkdownDescription: "Contains specifies a substring that must be present in the header value.",
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
														Description:         "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
														MarkdownDescription: "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"notcontains": schema.StringAttribute{
														Description:         "NotContains specifies a substring that must not be present in the header value.",
														MarkdownDescription: "NotContains specifies a substring that must not be present in the header value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"notexact": schema.StringAttribute{
														Description:         "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
														MarkdownDescription: "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"notpresent": schema.BoolAttribute{
														Description:         "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
														MarkdownDescription: "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"present": schema.BoolAttribute{
														Description:         "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
														MarkdownDescription: "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"regex": schema.StringAttribute{
														Description:         "Regex specifies a regular expression pattern that must match the header value.",
														MarkdownDescription: "Regex specifies a regular expression pattern that must match the header value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"treat_missing_as_empty": schema.BoolAttribute{
														Description:         "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
														MarkdownDescription: "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": schema.StringAttribute{
												Description:         "Prefix defines a prefix match for a request.",
												MarkdownDescription: "Prefix defines a prefix match for a request.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_parameter": schema.SingleNestedAttribute{
												Description:         "QueryParameter specifies the query parameter condition to match.",
												MarkdownDescription: "QueryParameter specifies the query parameter condition to match.",
												Attributes: map[string]schema.Attribute{
													"contains": schema.StringAttribute{
														Description:         "Contains specifies a substring that must be present in the query parameter value.",
														MarkdownDescription: "Contains specifies a substring that must be present in the query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exact": schema.StringAttribute{
														Description:         "Exact specifies a string that the query parameter value must be equal to.",
														MarkdownDescription: "Exact specifies a string that the query parameter value must be equal to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_case": schema.BoolAttribute{
														Description:         "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the query parameter to match against. Name is required. Query parameter names are case insensitive.",
														MarkdownDescription: "Name is the name of the query parameter to match against. Name is required. Query parameter names are case insensitive.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"prefix": schema.StringAttribute{
														Description:         "Prefix defines a prefix match for the query parameter value.",
														MarkdownDescription: "Prefix defines a prefix match for the query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"present": schema.BoolAttribute{
														Description:         "Present specifies that condition is true when the named query parameter is present, regardless of its value. Note that setting Present to false does not make the condition true if the named query parameter is absent.",
														MarkdownDescription: "Present specifies that condition is true when the named query parameter is present, regardless of its value. Note that setting Present to false does not make the condition true if the named query parameter is absent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"regex": schema.StringAttribute{
														Description:         "Regex specifies a regular expression pattern that must match the query parameter value.",
														MarkdownDescription: "Regex specifies a regular expression pattern that must match the query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"suffix": schema.StringAttribute{
														Description:         "Suffix defines a suffix match for a query parameter value.",
														MarkdownDescription: "Suffix defines a suffix match for a query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": schema.StringAttribute{
												Description:         "Regex defines a regex match for a request. This field is not allowed in include match conditions.",
												MarkdownDescription: "Regex defines a regex match for a request. This field is not allowed in include match conditions.",
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

								"name": schema.StringAttribute{
									Description:         "Name of the HTTPProxy",
									MarkdownDescription: "Name of the HTTPProxy",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.",
									MarkdownDescription: "Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.",
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

					"ingress_class_name": schema.StringAttribute{
						Description:         "IngressClassName optionally specifies the ingress class to use for this HTTPProxy. This replaces the deprecated 'kubernetes.io/ingress.class' annotation. For backwards compatibility, when that annotation is set, it is given precedence over this field.",
						MarkdownDescription: "IngressClassName optionally specifies the ingress class to use for this HTTPProxy. This replaces the deprecated 'kubernetes.io/ingress.class' annotation. For backwards compatibility, when that annotation is set, it is given precedence over this field.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"routes": schema.ListNestedAttribute{
						Description:         "Routes are the ingress routes. If TCPProxy is present, Routes is ignored.",
						MarkdownDescription: "Routes are the ingress routes. If TCPProxy is present, Routes is ignored.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"auth_policy": schema.SingleNestedAttribute{
									Description:         "AuthPolicy updates the authorization policy that was set on the root HTTPProxy object for client requests that match this route.",
									MarkdownDescription: "AuthPolicy updates the authorization policy that was set on the root HTTPProxy object for client requests that match this route.",
									Attributes: map[string]schema.Attribute{
										"context": schema.MapAttribute{
											Description:         "Context is a set of key/value pairs that are sent to the authentication server in the check request. If a context is provided at an enclosing scope, the entries are merged such that the inner scope overrides matching keys from the outer scope.",
											MarkdownDescription: "Context is a set of key/value pairs that are sent to the authentication server in the check request. If a context is provided at an enclosing scope, the entries are merged such that the inner scope overrides matching keys from the outer scope.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"disabled": schema.BoolAttribute{
											Description:         "When true, this field disables client request authentication for the scope of the policy.",
											MarkdownDescription: "When true, this field disables client request authentication for the scope of the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"conditions": schema.ListNestedAttribute{
									Description:         "Conditions are a set of rules that are applied to a Route. When applied, they are merged using AND, with one exception: There can be only one Prefix, Exact or Regex MatchCondition per Conditions slice. More than one of these condition types, or contradictory Conditions, will make the route invalid.",
									MarkdownDescription: "Conditions are a set of rules that are applied to a Route. When applied, they are merged using AND, with one exception: There can be only one Prefix, Exact or Regex MatchCondition per Conditions slice. More than one of these condition types, or contradictory Conditions, will make the route invalid.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exact": schema.StringAttribute{
												Description:         "Exact defines a exact match for a request. This field is not allowed in include match conditions.",
												MarkdownDescription: "Exact defines a exact match for a request. This field is not allowed in include match conditions.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"header": schema.SingleNestedAttribute{
												Description:         "Header specifies the header condition to match.",
												MarkdownDescription: "Header specifies the header condition to match.",
												Attributes: map[string]schema.Attribute{
													"contains": schema.StringAttribute{
														Description:         "Contains specifies a substring that must be present in the header value.",
														MarkdownDescription: "Contains specifies a substring that must be present in the header value.",
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
														Description:         "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
														MarkdownDescription: "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"notcontains": schema.StringAttribute{
														Description:         "NotContains specifies a substring that must not be present in the header value.",
														MarkdownDescription: "NotContains specifies a substring that must not be present in the header value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"notexact": schema.StringAttribute{
														Description:         "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
														MarkdownDescription: "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"notpresent": schema.BoolAttribute{
														Description:         "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
														MarkdownDescription: "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"present": schema.BoolAttribute{
														Description:         "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
														MarkdownDescription: "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"regex": schema.StringAttribute{
														Description:         "Regex specifies a regular expression pattern that must match the header value.",
														MarkdownDescription: "Regex specifies a regular expression pattern that must match the header value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"treat_missing_as_empty": schema.BoolAttribute{
														Description:         "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
														MarkdownDescription: "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": schema.StringAttribute{
												Description:         "Prefix defines a prefix match for a request.",
												MarkdownDescription: "Prefix defines a prefix match for a request.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_parameter": schema.SingleNestedAttribute{
												Description:         "QueryParameter specifies the query parameter condition to match.",
												MarkdownDescription: "QueryParameter specifies the query parameter condition to match.",
												Attributes: map[string]schema.Attribute{
													"contains": schema.StringAttribute{
														Description:         "Contains specifies a substring that must be present in the query parameter value.",
														MarkdownDescription: "Contains specifies a substring that must be present in the query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exact": schema.StringAttribute{
														Description:         "Exact specifies a string that the query parameter value must be equal to.",
														MarkdownDescription: "Exact specifies a string that the query parameter value must be equal to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_case": schema.BoolAttribute{
														Description:         "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the query parameter to match against. Name is required. Query parameter names are case insensitive.",
														MarkdownDescription: "Name is the name of the query parameter to match against. Name is required. Query parameter names are case insensitive.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"prefix": schema.StringAttribute{
														Description:         "Prefix defines a prefix match for the query parameter value.",
														MarkdownDescription: "Prefix defines a prefix match for the query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"present": schema.BoolAttribute{
														Description:         "Present specifies that condition is true when the named query parameter is present, regardless of its value. Note that setting Present to false does not make the condition true if the named query parameter is absent.",
														MarkdownDescription: "Present specifies that condition is true when the named query parameter is present, regardless of its value. Note that setting Present to false does not make the condition true if the named query parameter is absent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"regex": schema.StringAttribute{
														Description:         "Regex specifies a regular expression pattern that must match the query parameter value.",
														MarkdownDescription: "Regex specifies a regular expression pattern that must match the query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"suffix": schema.StringAttribute{
														Description:         "Suffix defines a suffix match for a query parameter value.",
														MarkdownDescription: "Suffix defines a suffix match for a query parameter value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": schema.StringAttribute{
												Description:         "Regex defines a regex match for a request. This field is not allowed in include match conditions.",
												MarkdownDescription: "Regex defines a regex match for a request. This field is not allowed in include match conditions.",
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

								"cookie_rewrite_policies": schema.ListNestedAttribute{
									Description:         "The policies for rewriting Set-Cookie header attributes. Note that rewritten cookie names must be unique in this list. Order rewrite policies are specified in does not matter.",
									MarkdownDescription: "The policies for rewriting Set-Cookie header attributes. Note that rewritten cookie names must be unique in this list. Order rewrite policies are specified in does not matter.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"domain_rewrite": schema.SingleNestedAttribute{
												Description:         "DomainRewrite enables rewriting the Set-Cookie Domain element. If not set, Domain will not be rewritten.",
												MarkdownDescription: "DomainRewrite enables rewriting the Set-Cookie Domain element. If not set, Domain will not be rewritten.",
												Attributes: map[string]schema.Attribute{
													"value": schema.StringAttribute{
														Description:         "Value is the value to rewrite the Domain attribute to. For now this is required.",
														MarkdownDescription: "Value is the value to rewrite the Domain attribute to. For now this is required.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(4096),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of the cookie for which attributes will be rewritten.",
												MarkdownDescription: "Name is the name of the cookie for which attributes will be rewritten.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(4096),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[^()<>@,;:\\"\/[\]?={} \t\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`), ""),
												},
											},

											"path_rewrite": schema.SingleNestedAttribute{
												Description:         "PathRewrite enables rewriting the Set-Cookie Path element. If not set, Path will not be rewritten.",
												MarkdownDescription: "PathRewrite enables rewriting the Set-Cookie Path element. If not set, Path will not be rewritten.",
												Attributes: map[string]schema.Attribute{
													"value": schema.StringAttribute{
														Description:         "Value is the value to rewrite the Path attribute to. For now this is required.",
														MarkdownDescription: "Value is the value to rewrite the Path attribute to. For now this is required.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(4096),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[^;\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"same_site": schema.StringAttribute{
												Description:         "SameSite enables rewriting the Set-Cookie SameSite element. If not set, SameSite attribute will not be rewritten.",
												MarkdownDescription: "SameSite enables rewriting the Set-Cookie SameSite element. If not set, SameSite attribute will not be rewritten.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Strict", "Lax", "None"),
												},
											},

											"secure": schema.BoolAttribute{
												Description:         "Secure enables rewriting the Set-Cookie Secure element. If not set, Secure attribute will not be rewritten.",
												MarkdownDescription: "Secure enables rewriting the Set-Cookie Secure element. If not set, Secure attribute will not be rewritten.",
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

								"direct_response_policy": schema.SingleNestedAttribute{
									Description:         "DirectResponsePolicy returns an arbitrary HTTP response directly.",
									MarkdownDescription: "DirectResponsePolicy returns an arbitrary HTTP response directly.",
									Attributes: map[string]schema.Attribute{
										"body": schema.StringAttribute{
											Description:         "Body is the content of the response body. If this setting is omitted, no body is included in the generated response. Note: Body is not recommended to set too long otherwise it can have significant resource usage impacts.",
											MarkdownDescription: "Body is the content of the response body. If this setting is omitted, no body is included in the generated response. Note: Body is not recommended to set too long otherwise it can have significant resource usage impacts.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"status_code": schema.Int64Attribute{
											Description:         "StatusCode is the HTTP response status to be returned.",
											MarkdownDescription: "StatusCode is the HTTP response status to be returned.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(200),
												int64validator.AtMost(599),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"enable_websockets": schema.BoolAttribute{
									Description:         "Enables websocket support for the route.",
									MarkdownDescription: "Enables websocket support for the route.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"health_check_policy": schema.SingleNestedAttribute{
									Description:         "The health check policy for this route.",
									MarkdownDescription: "The health check policy for this route.",
									Attributes: map[string]schema.Attribute{
										"expected_statuses": schema.ListNestedAttribute{
											Description:         "The ranges of HTTP response statuses considered healthy. Follow half-open semantics, i.e. for each range the start is inclusive and the end is exclusive. Must be within the range [100,600). If not specified, only a 200 response status is considered healthy.",
											MarkdownDescription: "The ranges of HTTP response statuses considered healthy. Follow half-open semantics, i.e. for each range the start is inclusive and the end is exclusive. Must be within the range [100,600). If not specified, only a 200 response status is considered healthy.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "The end (exclusive) of a range of HTTP status codes.",
														MarkdownDescription: "The end (exclusive) of a range of HTTP status codes.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(101),
															int64validator.AtMost(600),
														},
													},

													"start": schema.Int64Attribute{
														Description:         "The start (inclusive) of a range of HTTP status codes.",
														MarkdownDescription: "The start (inclusive) of a range of HTTP status codes.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(100),
															int64validator.AtMost(599),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"healthy_threshold_count": schema.Int64Attribute{
											Description:         "The number of healthy health checks required before a host is marked healthy",
											MarkdownDescription: "The number of healthy health checks required before a host is marked healthy",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"host": schema.StringAttribute{
											Description:         "The value of the host header in the HTTP health check request. If left empty (default value), the name 'contour-envoy-healthcheck' will be used.",
											MarkdownDescription: "The value of the host header in the HTTP health check request. If left empty (default value), the name 'contour-envoy-healthcheck' will be used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interval_seconds": schema.Int64Attribute{
											Description:         "The interval (seconds) between health checks",
											MarkdownDescription: "The interval (seconds) between health checks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "HTTP endpoint used to perform health checks on upstream service",
											MarkdownDescription: "HTTP endpoint used to perform health checks on upstream service",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"timeout_seconds": schema.Int64Attribute{
											Description:         "The time to wait (seconds) for a health check response",
											MarkdownDescription: "The time to wait (seconds) for a health check response",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"unhealthy_threshold_count": schema.Int64Attribute{
											Description:         "The number of unhealthy health checks required before a host is marked unhealthy",
											MarkdownDescription: "The number of unhealthy health checks required before a host is marked unhealthy",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"internal_redirect_policy": schema.SingleNestedAttribute{
									Description:         "The policy to define when to handle redirects responses internally.",
									MarkdownDescription: "The policy to define when to handle redirects responses internally.",
									Attributes: map[string]schema.Attribute{
										"allow_cross_scheme_redirect": schema.StringAttribute{
											Description:         "AllowCrossSchemeRedirect Allow internal redirect to follow a target URI with a different scheme than the value of x-forwarded-proto. SafeOnly allows same scheme redirect and safe cross scheme redirect, which means if the downstream scheme is HTTPS, both HTTPS and HTTP redirect targets are allowed, but if the downstream scheme is HTTP, only HTTP redirect targets are allowed.",
											MarkdownDescription: "AllowCrossSchemeRedirect Allow internal redirect to follow a target URI with a different scheme than the value of x-forwarded-proto. SafeOnly allows same scheme redirect and safe cross scheme redirect, which means if the downstream scheme is HTTPS, both HTTPS and HTTP redirect targets are allowed, but if the downstream scheme is HTTP, only HTTP redirect targets are allowed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Always", "Never", "SafeOnly"),
											},
										},

										"deny_repeated_route_redirect": schema.BoolAttribute{
											Description:         "If DenyRepeatedRouteRedirect is true, rejects redirect targets that are pointing to a route that has been followed by a previous redirect from the current route.",
											MarkdownDescription: "If DenyRepeatedRouteRedirect is true, rejects redirect targets that are pointing to a route that has been followed by a previous redirect from the current route.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_internal_redirects": schema.Int64Attribute{
											Description:         "MaxInternalRedirects An internal redirect is not handled, unless the number of previous internal redirects that a downstream request has encountered is lower than this value.",
											MarkdownDescription: "MaxInternalRedirects An internal redirect is not handled, unless the number of previous internal redirects that a downstream request has encountered is lower than this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redirect_response_codes": schema.ListAttribute{
											Description:         "RedirectResponseCodes If unspecified, only 302 will be treated as internal redirect. Only 301, 302, 303, 307 and 308 are valid values.",
											MarkdownDescription: "RedirectResponseCodes If unspecified, only 302 will be treated as internal redirect. Only 301, 302, 303, 307 and 308 are valid values.",
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

								"ip_allow_policy": schema.ListNestedAttribute{
									Description:         "IPAllowFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be allowed. All other requests will be denied. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here override any rules set on the root HTTPProxy.",
									MarkdownDescription: "IPAllowFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be allowed. All other requests will be denied. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here override any rules set on the root HTTPProxy.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
												Description:         "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
												MarkdownDescription: "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"source": schema.StringAttribute{
												Description:         "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
												MarkdownDescription: "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Peer", "Remote"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ip_deny_policy": schema.ListNestedAttribute{
									Description:         "IPDenyFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be denied. All other requests will be allowed. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here override any rules set on the root HTTPProxy.",
									MarkdownDescription: "IPDenyFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be denied. All other requests will be allowed. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here override any rules set on the root HTTPProxy.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
												Description:         "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
												MarkdownDescription: "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"source": schema.StringAttribute{
												Description:         "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
												MarkdownDescription: "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Peer", "Remote"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"jwt_verification_policy": schema.SingleNestedAttribute{
									Description:         "The policy for verifying JWTs for requests to this route.",
									MarkdownDescription: "The policy for verifying JWTs for requests to this route.",
									Attributes: map[string]schema.Attribute{
										"disabled": schema.BoolAttribute{
											Description:         "Disabled defines whether to disable all JWT verification for this route. This can be used to opt specific routes out of the default JWT provider for the HTTPProxy. At most one of this field or the 'require' field can be specified.",
											MarkdownDescription: "Disabled defines whether to disable all JWT verification for this route. This can be used to opt specific routes out of the default JWT provider for the HTTPProxy. At most one of this field or the 'require' field can be specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"require": schema.StringAttribute{
											Description:         "Require names a specific JWT provider (defined in the virtual host) to require for the route. If specified, this field overrides the default provider if one exists. If this field is not specified, the default provider will be required if one exists. At most one of this field or the 'disabled' field can be specified.",
											MarkdownDescription: "Require names a specific JWT provider (defined in the virtual host) to require for the route. If specified, this field overrides the default provider if one exists. If this field is not specified, the default provider will be required if one exists. At most one of this field or the 'disabled' field can be specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"load_balancer_policy": schema.SingleNestedAttribute{
									Description:         "The load balancing policy for this route.",
									MarkdownDescription: "The load balancing policy for this route.",
									Attributes: map[string]schema.Attribute{
										"request_hash_policies": schema.ListNestedAttribute{
											Description:         "RequestHashPolicies contains a list of hash policies to apply when the 'RequestHash' load balancing strategy is chosen. If an element of the supplied list of hash policies is invalid, it will be ignored. If the list of hash policies is empty after validation, the load balancing strategy will fall back to the default 'RoundRobin'.",
											MarkdownDescription: "RequestHashPolicies contains a list of hash policies to apply when the 'RequestHash' load balancing strategy is chosen. If an element of the supplied list of hash policies is invalid, it will be ignored. If the list of hash policies is empty after validation, the load balancing strategy will fall back to the default 'RoundRobin'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"hash_source_ip": schema.BoolAttribute{
														Description:         "HashSourceIP should be set to true when request source IP hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
														MarkdownDescription: "HashSourceIP should be set to true when request source IP hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"header_hash_options": schema.SingleNestedAttribute{
														Description:         "HeaderHashOptions should be set when request header hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
														MarkdownDescription: "HeaderHashOptions should be set when request header hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
														Attributes: map[string]schema.Attribute{
															"header_name": schema.StringAttribute{
																Description:         "HeaderName is the name of the HTTP request header that will be used to calculate the hash key. If the header specified is not present on a request, no hash will be produced.",
																MarkdownDescription: "HeaderName is the name of the HTTP request header that will be used to calculate the hash key. If the header specified is not present on a request, no hash will be produced.",
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

													"query_parameter_hash_options": schema.SingleNestedAttribute{
														Description:         "QueryParameterHashOptions should be set when request query parameter hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
														MarkdownDescription: "QueryParameterHashOptions should be set when request query parameter hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
														Attributes: map[string]schema.Attribute{
															"parameter_name": schema.StringAttribute{
																Description:         "ParameterName is the name of the HTTP request query parameter that will be used to calculate the hash key. If the query parameter specified is not present on a request, no hash will be produced.",
																MarkdownDescription: "ParameterName is the name of the HTTP request query parameter that will be used to calculate the hash key. If the query parameter specified is not present on a request, no hash will be produced.",
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

													"terminal": schema.BoolAttribute{
														Description:         "Terminal is a flag that allows for short-circuiting computing of a hash for a given request. If set to true, and the request attribute specified in the attribute hash options is present, no further hash policies will be used to calculate a hash for the request.",
														MarkdownDescription: "Terminal is a flag that allows for short-circuiting computing of a hash for a given request. If set to true, and the request attribute specified in the attribute hash options is present, no further hash policies will be used to calculate a hash for the request.",
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

										"strategy": schema.StringAttribute{
											Description:         "Strategy specifies the policy used to balance requests across the pool of backend pods. Valid policy names are 'Random', 'RoundRobin', 'WeightedLeastRequest', 'Cookie', and 'RequestHash'. If an unknown strategy name is specified or no policy is supplied, the default 'RoundRobin' policy is used.",
											MarkdownDescription: "Strategy specifies the policy used to balance requests across the pool of backend pods. Valid policy names are 'Random', 'RoundRobin', 'WeightedLeastRequest', 'Cookie', and 'RequestHash'. If an unknown strategy name is specified or no policy is supplied, the default 'RoundRobin' policy is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"path_rewrite_policy": schema.SingleNestedAttribute{
									Description:         "The policy for rewriting the path of the request URL after the request has been routed to a Service.",
									MarkdownDescription: "The policy for rewriting the path of the request URL after the request has been routed to a Service.",
									Attributes: map[string]schema.Attribute{
										"replace_prefix": schema.ListNestedAttribute{
											Description:         "ReplacePrefix describes how the path prefix should be replaced.",
											MarkdownDescription: "ReplacePrefix describes how the path prefix should be replaced.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"prefix": schema.StringAttribute{
														Description:         "Prefix specifies the URL path prefix to be replaced. If Prefix is specified, it must exactly match the MatchCondition prefix that is rendered by the chain of including HTTPProxies and only that path prefix will be replaced by Replacement. This allows HTTPProxies that are included through multiple roots to only replace specific path prefixes, leaving others unmodified. If Prefix is not specified, all routing prefixes rendered by the include chain will be replaced.",
														MarkdownDescription: "Prefix specifies the URL path prefix to be replaced. If Prefix is specified, it must exactly match the MatchCondition prefix that is rendered by the chain of including HTTPProxies and only that path prefix will be replaced by Replacement. This allows HTTPProxies that are included through multiple roots to only replace specific path prefixes, leaving others unmodified. If Prefix is not specified, all routing prefixes rendered by the include chain will be replaced.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"replacement": schema.StringAttribute{
														Description:         "Replacement is the string that the routing path prefix will be replaced with. This must not be empty.",
														MarkdownDescription: "Replacement is the string that the routing path prefix will be replaced with. This must not be empty.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
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

								"permit_insecure": schema.BoolAttribute{
									Description:         "Allow this path to respond to insecure requests over HTTP which are normally not permitted when a 'virtualhost.tls' block is present.",
									MarkdownDescription: "Allow this path to respond to insecure requests over HTTP which are normally not permitted when a 'virtualhost.tls' block is present.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rate_limit_policy": schema.SingleNestedAttribute{
									Description:         "The policy for rate limiting on the route.",
									MarkdownDescription: "The policy for rate limiting on the route.",
									Attributes: map[string]schema.Attribute{
										"global": schema.SingleNestedAttribute{
											Description:         "Global defines global rate limiting parameters, i.e. parameters defining descriptors that are sent to an external rate limit service (RLS) for a rate limit decision on each request.",
											MarkdownDescription: "Global defines global rate limiting parameters, i.e. parameters defining descriptors that are sent to an external rate limit service (RLS) for a rate limit decision on each request.",
											Attributes: map[string]schema.Attribute{
												"descriptors": schema.ListNestedAttribute{
													Description:         "Descriptors defines the list of descriptors that will be generated and sent to the rate limit service. Each descriptor contains 1+ key-value pair entries.",
													MarkdownDescription: "Descriptors defines the list of descriptors that will be generated and sent to the rate limit service. Each descriptor contains 1+ key-value pair entries.",
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
																					Description:         "Key defines the key of the descriptor entry. If not set, the key is set to 'generic_key'.",
																					MarkdownDescription: "Key defines the key of the descriptor entry. If not set, the key is set to 'generic_key'.",
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
																			Description:         "RemoteAddress defines a descriptor entry with a key of 'remote_address' and a value equal to the client's IP address (from x-forwarded-for).",
																			MarkdownDescription: "RemoteAddress defines a descriptor entry with a key of 'remote_address' and a value equal to the client's IP address (from x-forwarded-for).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"request_header": schema.SingleNestedAttribute{
																			Description:         "RequestHeader defines a descriptor entry that's populated only if a given header is present on the request. The descriptor key is static, and the descriptor value is equal to the value of the header.",
																			MarkdownDescription: "RequestHeader defines a descriptor entry that's populated only if a given header is present on the request. The descriptor key is static, and the descriptor value is equal to the value of the header.",
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
																			Description:         "RequestHeaderValueMatch defines a descriptor entry that's populated if the request's headers match a set of 1+ match criteria. The descriptor key is 'header_match', and the descriptor value is static.",
																			MarkdownDescription: "RequestHeaderValueMatch defines a descriptor entry that's populated if the request's headers match a set of 1+ match criteria. The descriptor key is 'header_match', and the descriptor value is static.",
																			Attributes: map[string]schema.Attribute{
																				"expect_match": schema.BoolAttribute{
																					Description:         "ExpectMatch defines whether the request must positively match the match criteria in order to generate a descriptor entry (i.e. true), or not match the match criteria in order to generate a descriptor entry (i.e. false). The default is true.",
																					MarkdownDescription: "ExpectMatch defines whether the request must positively match the match criteria in order to generate a descriptor entry (i.e. true), or not match the match criteria in order to generate a descriptor entry (i.e. false). The default is true.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"headers": schema.ListNestedAttribute{
																					Description:         "Headers is a list of 1+ match criteria to apply against the request to determine whether to populate the descriptor entry or not.",
																					MarkdownDescription: "Headers is a list of 1+ match criteria to apply against the request to determine whether to populate the descriptor entry or not.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"contains": schema.StringAttribute{
																								Description:         "Contains specifies a substring that must be present in the header value.",
																								MarkdownDescription: "Contains specifies a substring that must be present in the header value.",
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
																								Description:         "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
																								MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"name": schema.StringAttribute{
																								Description:         "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
																								MarkdownDescription: "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"notcontains": schema.StringAttribute{
																								Description:         "NotContains specifies a substring that must not be present in the header value.",
																								MarkdownDescription: "NotContains specifies a substring that must not be present in the header value.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"notexact": schema.StringAttribute{
																								Description:         "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
																								MarkdownDescription: "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"notpresent": schema.BoolAttribute{
																								Description:         "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
																								MarkdownDescription: "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"present": schema.BoolAttribute{
																								Description:         "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
																								MarkdownDescription: "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"regex": schema.StringAttribute{
																								Description:         "Regex specifies a regular expression pattern that must match the header value.",
																								MarkdownDescription: "Regex specifies a regular expression pattern that must match the header value.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"treat_missing_as_empty": schema.BoolAttribute{
																								Description:         "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
																								MarkdownDescription: "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
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
													Description:         "Disabled configures the HTTPProxy to not use the default global rate limit policy defined by the Contour configuration.",
													MarkdownDescription: "Disabled configures the HTTPProxy to not use the default global rate limit policy defined by the Contour configuration.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"local": schema.SingleNestedAttribute{
											Description:         "Local defines local rate limiting parameters, i.e. parameters for rate limiting that occurs within each Envoy pod as requests are handled.",
											MarkdownDescription: "Local defines local rate limiting parameters, i.e. parameters for rate limiting that occurs within each Envoy pod as requests are handled.",
											Attributes: map[string]schema.Attribute{
												"burst": schema.Int64Attribute{
													Description:         "Burst defines the number of requests above the requests per unit that should be allowed within a short period of time.",
													MarkdownDescription: "Burst defines the number of requests above the requests per unit that should be allowed within a short period of time.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.Int64Attribute{
													Description:         "Requests defines how many requests per unit of time should be allowed before rate limiting occurs.",
													MarkdownDescription: "Requests defines how many requests per unit of time should be allowed before rate limiting occurs.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"response_headers_to_add": schema.ListNestedAttribute{
													Description:         "ResponseHeadersToAdd is an optional list of response headers to set when a request is rate-limited.",
													MarkdownDescription: "ResponseHeadersToAdd is an optional list of response headers to set when a request is rate-limited.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name represents a key of a header",
																MarkdownDescription: "Name represents a key of a header",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"value": schema.StringAttribute{
																Description:         "Value represents the value of a header specified by a key",
																MarkdownDescription: "Value represents the value of a header specified by a key",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"response_status_code": schema.Int64Attribute{
													Description:         "ResponseStatusCode is the HTTP status code to use for responses to rate-limited requests. Codes must be in the 400-599 range (inclusive). If not specified, the Envoy default of 429 (Too Many Requests) is used.",
													MarkdownDescription: "ResponseStatusCode is the HTTP status code to use for responses to rate-limited requests. Codes must be in the 400-599 range (inclusive). If not specified, the Envoy default of 429 (Too Many Requests) is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(400),
														int64validator.AtMost(599),
													},
												},

												"unit": schema.StringAttribute{
													Description:         "Unit defines the period of time within which requests over the limit will be rate limited. Valid values are 'second', 'minute' and 'hour'.",
													MarkdownDescription: "Unit defines the period of time within which requests over the limit will be rate limited. Valid values are 'second', 'minute' and 'hour'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("second", "minute", "hour"),
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

								"request_headers_policy": schema.SingleNestedAttribute{
									Description:         "The policy for managing request headers during proxying. You may dynamically rewrite the Host header to be forwarded upstream to the content of a request header using the below format '%REQ(X-Header-Name)%'. If the value of the header is empty, it is ignored. *NOTE: Pay attention to the potential security implications of using this option. Provided header must come from trusted source. **NOTE: The header rewrite is only done while forwarding and has no bearing on the routing decision.",
									MarkdownDescription: "The policy for managing request headers during proxying. You may dynamically rewrite the Host header to be forwarded upstream to the content of a request header using the below format '%REQ(X-Header-Name)%'. If the value of the header is empty, it is ignored. *NOTE: Pay attention to the potential security implications of using this option. Provided header must come from trusted source. **NOTE: The header rewrite is only done while forwarding and has no bearing on the routing decision.",
									Attributes: map[string]schema.Attribute{
										"remove": schema.ListAttribute{
											Description:         "Remove specifies a list of HTTP header names to remove.",
											MarkdownDescription: "Remove specifies a list of HTTP header names to remove.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"set": schema.ListNestedAttribute{
											Description:         "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
											MarkdownDescription: "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name represents a key of a header",
														MarkdownDescription: "Name represents a key of a header",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"value": schema.StringAttribute{
														Description:         "Value represents the value of a header specified by a key",
														MarkdownDescription: "Value represents the value of a header specified by a key",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
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

								"request_redirect_policy": schema.SingleNestedAttribute{
									Description:         "RequestRedirectPolicy defines an HTTP redirection.",
									MarkdownDescription: "RequestRedirectPolicy defines an HTTP redirection.",
									Attributes: map[string]schema.Attribute{
										"hostname": schema.StringAttribute{
											Description:         "Hostname is the precise hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used. No wildcards are allowed.",
											MarkdownDescription: "Hostname is the precise hostname to be used in the value of the 'Location' header in the response. When empty, the hostname of the request is used. No wildcards are allowed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
											},
										},

										"path": schema.StringAttribute{
											Description:         "Path allows for redirection to a different path from the original on the request. The path must start with a leading slash. Note: Only one of Path or Prefix can be defined.",
											MarkdownDescription: "Path allows for redirection to a different path from the original on the request. The path must start with a leading slash. Note: Only one of Path or Prefix can be defined.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\/.*$`), ""),
											},
										},

										"port": schema.Int64Attribute{
											Description:         "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.",
											MarkdownDescription: "Port is the port to be used in the value of the 'Location' header in the response. When empty, port (if specified) of the request is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"prefix": schema.StringAttribute{
											Description:         "Prefix defines the value to swap the matched prefix or path with. The prefix must start with a leading slash. Note: Only one of Path or Prefix can be defined.",
											MarkdownDescription: "Prefix defines the value to swap the matched prefix or path with. The prefix must start with a leading slash. Note: Only one of Path or Prefix can be defined.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\/.*$`), ""),
											},
										},

										"scheme": schema.StringAttribute{
											Description:         "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.",
											MarkdownDescription: "Scheme is the scheme to be used in the value of the 'Location' header in the response. When empty, the scheme of the request is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("http", "https"),
											},
										},

										"status_code": schema.Int64Attribute{
											Description:         "StatusCode is the HTTP status code to be used in response.",
											MarkdownDescription: "StatusCode is the HTTP status code to be used in response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.OneOf(301, 302),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"response_headers_policy": schema.SingleNestedAttribute{
									Description:         "The policy for managing response headers during proxying. Rewriting the 'Host' header is not supported.",
									MarkdownDescription: "The policy for managing response headers during proxying. Rewriting the 'Host' header is not supported.",
									Attributes: map[string]schema.Attribute{
										"remove": schema.ListAttribute{
											Description:         "Remove specifies a list of HTTP header names to remove.",
											MarkdownDescription: "Remove specifies a list of HTTP header names to remove.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"set": schema.ListNestedAttribute{
											Description:         "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
											MarkdownDescription: "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name represents a key of a header",
														MarkdownDescription: "Name represents a key of a header",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"value": schema.StringAttribute{
														Description:         "Value represents the value of a header specified by a key",
														MarkdownDescription: "Value represents the value of a header specified by a key",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
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

								"retry_policy": schema.SingleNestedAttribute{
									Description:         "The retry policy for this route.",
									MarkdownDescription: "The retry policy for this route.",
									Attributes: map[string]schema.Attribute{
										"count": schema.Int64Attribute{
											Description:         "NumRetries is maximum allowed number of retries. If set to -1, then retries are disabled. If set to 0 or not supplied, the value is set to the Envoy default of 1.",
											MarkdownDescription: "NumRetries is maximum allowed number of retries. If set to -1, then retries are disabled. If set to 0 or not supplied, the value is set to the Envoy default of 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(-1),
											},
										},

										"per_try_timeout": schema.StringAttribute{
											Description:         "PerTryTimeout specifies the timeout per retry attempt. Ignored if NumRetries is not supplied.",
											MarkdownDescription: "PerTryTimeout specifies the timeout per retry attempt. Ignored if NumRetries is not supplied.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
											},
										},

										"retriable_status_codes": schema.ListAttribute{
											Description:         "RetriableStatusCodes specifies the HTTP status codes that should be retried. This field is only respected when you include 'retriable-status-codes' in the 'RetryOn' field.",
											MarkdownDescription: "RetriableStatusCodes specifies the HTTP status codes that should be retried. This field is only respected when you include 'retriable-status-codes' in the 'RetryOn' field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_on": schema.ListAttribute{
											Description:         "RetryOn specifies the conditions on which to retry a request. Supported [HTTP conditions](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on): - '5xx' - 'gateway-error' - 'reset' - 'connect-failure' - 'retriable-4xx' - 'refused-stream' - 'retriable-status-codes' - 'retriable-headers' Supported [gRPC conditions](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-grpc-on): - 'cancelled' - 'deadline-exceeded' - 'internal' - 'resource-exhausted' - 'unavailable'",
											MarkdownDescription: "RetryOn specifies the conditions on which to retry a request. Supported [HTTP conditions](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on): - '5xx' - 'gateway-error' - 'reset' - 'connect-failure' - 'retriable-4xx' - 'refused-stream' - 'retriable-status-codes' - 'retriable-headers' Supported [gRPC conditions](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-grpc-on): - 'cancelled' - 'deadline-exceeded' - 'internal' - 'resource-exhausted' - 'unavailable'",
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

								"services": schema.ListNestedAttribute{
									Description:         "Services are the services to proxy traffic.",
									MarkdownDescription: "Services are the services to proxy traffic.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cookie_rewrite_policies": schema.ListNestedAttribute{
												Description:         "The policies for rewriting Set-Cookie header attributes.",
												MarkdownDescription: "The policies for rewriting Set-Cookie header attributes.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"domain_rewrite": schema.SingleNestedAttribute{
															Description:         "DomainRewrite enables rewriting the Set-Cookie Domain element. If not set, Domain will not be rewritten.",
															MarkdownDescription: "DomainRewrite enables rewriting the Set-Cookie Domain element. If not set, Domain will not be rewritten.",
															Attributes: map[string]schema.Attribute{
																"value": schema.StringAttribute{
																	Description:         "Value is the value to rewrite the Domain attribute to. For now this is required.",
																	MarkdownDescription: "Value is the value to rewrite the Domain attribute to. For now this is required.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(4096),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of the cookie for which attributes will be rewritten.",
															MarkdownDescription: "Name is the name of the cookie for which attributes will be rewritten.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(4096),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[^()<>@,;:\\"\/[\]?={} \t\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`), ""),
															},
														},

														"path_rewrite": schema.SingleNestedAttribute{
															Description:         "PathRewrite enables rewriting the Set-Cookie Path element. If not set, Path will not be rewritten.",
															MarkdownDescription: "PathRewrite enables rewriting the Set-Cookie Path element. If not set, Path will not be rewritten.",
															Attributes: map[string]schema.Attribute{
																"value": schema.StringAttribute{
																	Description:         "Value is the value to rewrite the Path attribute to. For now this is required.",
																	MarkdownDescription: "Value is the value to rewrite the Path attribute to. For now this is required.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(4096),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[^;\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`), ""),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"same_site": schema.StringAttribute{
															Description:         "SameSite enables rewriting the Set-Cookie SameSite element. If not set, SameSite attribute will not be rewritten.",
															MarkdownDescription: "SameSite enables rewriting the Set-Cookie SameSite element. If not set, SameSite attribute will not be rewritten.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Strict", "Lax", "None"),
															},
														},

														"secure": schema.BoolAttribute{
															Description:         "Secure enables rewriting the Set-Cookie Secure element. If not set, Secure attribute will not be rewritten.",
															MarkdownDescription: "Secure enables rewriting the Set-Cookie Secure element. If not set, Secure attribute will not be rewritten.",
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

											"health_port": schema.Int64Attribute{
												Description:         "HealthPort is the port for this service healthcheck. If not specified, Port is used for service healthchecks.",
												MarkdownDescription: "HealthPort is the port for this service healthcheck. If not specified, Port is used for service healthchecks.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"mirror": schema.BoolAttribute{
												Description:         "If Mirror is true the Service will receive a read only mirror of the traffic for this route. If Mirror is true, then fractional mirroring can be enabled by optionally setting the Weight field. Legal values for Weight are 1-100. Omitting the Weight field will result in 100% mirroring. NOTE: Setting Weight explicitly to 0 will unexpectedly result in 100% traffic mirroring. This occurs since we cannot distinguish omitted fields from those explicitly set to their default values",
												MarkdownDescription: "If Mirror is true the Service will receive a read only mirror of the traffic for this route. If Mirror is true, then fractional mirroring can be enabled by optionally setting the Weight field. Legal values for Weight are 1-100. Omitting the Weight field will result in 100% mirroring. NOTE: Setting Weight explicitly to 0 will unexpectedly result in 100% traffic mirroring. This occurs since we cannot distinguish omitted fields from those explicitly set to their default values",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of Kubernetes service to proxy traffic. Names defined here will be used to look up corresponding endpoints which contain the ips to route.",
												MarkdownDescription: "Name is the name of Kubernetes service to proxy traffic. Names defined here will be used to look up corresponding endpoints which contain the ips to route.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port (defined as Integer) to proxy traffic to since a service can have multiple defined.",
												MarkdownDescription: "Port (defined as Integer) to proxy traffic to since a service can have multiple defined.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"protocol": schema.StringAttribute{
												Description:         "Protocol may be used to specify (or override) the protocol used to reach this Service. Values may be tls, h2, h2c. If omitted, protocol-selection falls back on Service annotations.",
												MarkdownDescription: "Protocol may be used to specify (or override) the protocol used to reach this Service. Values may be tls, h2, h2c. If omitted, protocol-selection falls back on Service annotations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("h2", "h2c", "tls"),
												},
											},

											"request_headers_policy": schema.SingleNestedAttribute{
												Description:         "The policy for managing request headers during proxying.",
												MarkdownDescription: "The policy for managing request headers during proxying.",
												Attributes: map[string]schema.Attribute{
													"remove": schema.ListAttribute{
														Description:         "Remove specifies a list of HTTP header names to remove.",
														MarkdownDescription: "Remove specifies a list of HTTP header names to remove.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
														MarkdownDescription: "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name represents a key of a header",
																	MarkdownDescription: "Name represents a key of a header",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value represents the value of a header specified by a key",
																	MarkdownDescription: "Value represents the value of a header specified by a key",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
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

											"response_headers_policy": schema.SingleNestedAttribute{
												Description:         "The policy for managing response headers during proxying. Rewriting the 'Host' header is not supported.",
												MarkdownDescription: "The policy for managing response headers during proxying. Rewriting the 'Host' header is not supported.",
												Attributes: map[string]schema.Attribute{
													"remove": schema.ListAttribute{
														Description:         "Remove specifies a list of HTTP header names to remove.",
														MarkdownDescription: "Remove specifies a list of HTTP header names to remove.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set": schema.ListNestedAttribute{
														Description:         "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
														MarkdownDescription: "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name represents a key of a header",
																	MarkdownDescription: "Name represents a key of a header",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value represents the value of a header specified by a key",
																	MarkdownDescription: "Value represents the value of a header specified by a key",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
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

											"slow_start_policy": schema.SingleNestedAttribute{
												Description:         "Slow start will gradually increase amount of traffic to a newly added endpoint.",
												MarkdownDescription: "Slow start will gradually increase amount of traffic to a newly added endpoint.",
												Attributes: map[string]schema.Attribute{
													"aggression": schema.StringAttribute{
														Description:         "The speed of traffic increase over the slow start window. Defaults to 1.0, so that endpoint would get linearly increasing amount of traffic. When increasing the value for this parameter, the speed of traffic ramp-up increases non-linearly. The value of aggression parameter should be greater than 0.0. More info: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/slow_start",
														MarkdownDescription: "The speed of traffic increase over the slow start window. Defaults to 1.0, so that endpoint would get linearly increasing amount of traffic. When increasing the value for this parameter, the speed of traffic ramp-up increases non-linearly. The value of aggression parameter should be greater than 0.0. More info: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/slow_start",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+([.][0-9]+)?|[.][0-9]+)$`), ""),
														},
													},

													"min_weight_percent": schema.Int64Attribute{
														Description:         "The minimum or starting percentage of traffic to send to new endpoints. A non-zero value helps avoid a too small initial weight, which may cause endpoints in slow start mode to receive no traffic in the beginning of the slow start window. If not specified, the default is 10%.",
														MarkdownDescription: "The minimum or starting percentage of traffic to send to new endpoints. A non-zero value helps avoid a too small initial weight, which may cause endpoints in slow start mode to receive no traffic in the beginning of the slow start window. If not specified, the default is 10%.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(100),
														},
													},

													"window": schema.StringAttribute{
														Description:         "The duration of slow start window. Duration is expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
														MarkdownDescription: "The duration of slow start window. Duration is expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"validation": schema.SingleNestedAttribute{
												Description:         "UpstreamValidation defines how to verify the backend service's certificate",
												MarkdownDescription: "UpstreamValidation defines how to verify the backend service's certificate",
												Attributes: map[string]schema.Attribute{
													"ca_secret": schema.StringAttribute{
														Description:         "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
														MarkdownDescription: "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(317),
														},
													},

													"subject_name": schema.StringAttribute{
														Description:         "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
														MarkdownDescription: "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(250),
														},
													},

													"subject_names": schema.ListAttribute{
														Description:         "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
														MarkdownDescription: "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
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

											"weight": schema.Int64Attribute{
												Description:         "Weight defines percentage of traffic to balance traffic",
												MarkdownDescription: "Weight defines percentage of traffic to balance traffic",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"timeout_policy": schema.SingleNestedAttribute{
									Description:         "The timeout policy for this route.",
									MarkdownDescription: "The timeout policy for this route.",
									Attributes: map[string]schema.Attribute{
										"idle": schema.StringAttribute{
											Description:         "Timeout for how long the proxy should wait while there is no activity during single request/response (for HTTP/1.1) or stream (for HTTP/2). Timeout will not trigger while HTTP/1.1 connection is idle between two consecutive requests. If not specified, there is no per-route idle timeout, though a connection manager-wide stream_idle_timeout default of 5m still applies.",
											MarkdownDescription: "Timeout for how long the proxy should wait while there is no activity during single request/response (for HTTP/1.1) or stream (for HTTP/2). Timeout will not trigger while HTTP/1.1 connection is idle between two consecutive requests. If not specified, there is no per-route idle timeout, though a connection manager-wide stream_idle_timeout default of 5m still applies.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
											},
										},

										"idle_connection": schema.StringAttribute{
											Description:         "Timeout for how long connection from the proxy to the upstream service is kept when there are no active requests. If not supplied, Envoy's default value of 1h applies.",
											MarkdownDescription: "Timeout for how long connection from the proxy to the upstream service is kept when there are no active requests. If not supplied, Envoy's default value of 1h applies.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
											},
										},

										"response": schema.StringAttribute{
											Description:         "Timeout for receiving a response from the server after processing a request from client. If not supplied, Envoy's default value of 15s applies.",
											MarkdownDescription: "Timeout for receiving a response from the server after processing a request from client. If not supplied, Envoy's default value of 15s applies.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
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

					"tcpproxy": schema.SingleNestedAttribute{
						Description:         "TCPProxy holds TCP proxy information.",
						MarkdownDescription: "TCPProxy holds TCP proxy information.",
						Attributes: map[string]schema.Attribute{
							"health_check_policy": schema.SingleNestedAttribute{
								Description:         "The health check policy for this tcp proxy",
								MarkdownDescription: "The health check policy for this tcp proxy",
								Attributes: map[string]schema.Attribute{
									"healthy_threshold_count": schema.Int64Attribute{
										Description:         "The number of healthy health checks required before a host is marked healthy",
										MarkdownDescription: "The number of healthy health checks required before a host is marked healthy",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval_seconds": schema.Int64Attribute{
										Description:         "The interval (seconds) between health checks",
										MarkdownDescription: "The interval (seconds) between health checks",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The time to wait (seconds) for a health check response",
										MarkdownDescription: "The time to wait (seconds) for a health check response",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"unhealthy_threshold_count": schema.Int64Attribute{
										Description:         "The number of unhealthy health checks required before a host is marked unhealthy",
										MarkdownDescription: "The number of unhealthy health checks required before a host is marked unhealthy",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"include": schema.SingleNestedAttribute{
								Description:         "Include specifies that this tcpproxy should be delegated to another HTTPProxy.",
								MarkdownDescription: "Include specifies that this tcpproxy should be delegated to another HTTPProxy.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the child HTTPProxy",
										MarkdownDescription: "Name of the child HTTPProxy",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.",
										MarkdownDescription: "Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"includes": schema.SingleNestedAttribute{
								Description:         "IncludesDeprecated allow for specific routing configuration to be appended to another HTTPProxy in another namespace. Exists due to a mistake when developing HTTPProxy and the field was marked plural when it should have been singular. This field should stay to not break backwards compatibility to v1 users.",
								MarkdownDescription: "IncludesDeprecated allow for specific routing configuration to be appended to another HTTPProxy in another namespace. Exists due to a mistake when developing HTTPProxy and the field was marked plural when it should have been singular. This field should stay to not break backwards compatibility to v1 users.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the child HTTPProxy",
										MarkdownDescription: "Name of the child HTTPProxy",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.",
										MarkdownDescription: "Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"load_balancer_policy": schema.SingleNestedAttribute{
								Description:         "The load balancing policy for the backend services. Note that the 'Cookie' and 'RequestHash' load balancing strategies cannot be used here.",
								MarkdownDescription: "The load balancing policy for the backend services. Note that the 'Cookie' and 'RequestHash' load balancing strategies cannot be used here.",
								Attributes: map[string]schema.Attribute{
									"request_hash_policies": schema.ListNestedAttribute{
										Description:         "RequestHashPolicies contains a list of hash policies to apply when the 'RequestHash' load balancing strategy is chosen. If an element of the supplied list of hash policies is invalid, it will be ignored. If the list of hash policies is empty after validation, the load balancing strategy will fall back to the default 'RoundRobin'.",
										MarkdownDescription: "RequestHashPolicies contains a list of hash policies to apply when the 'RequestHash' load balancing strategy is chosen. If an element of the supplied list of hash policies is invalid, it will be ignored. If the list of hash policies is empty after validation, the load balancing strategy will fall back to the default 'RoundRobin'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hash_source_ip": schema.BoolAttribute{
													Description:         "HashSourceIP should be set to true when request source IP hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
													MarkdownDescription: "HashSourceIP should be set to true when request source IP hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"header_hash_options": schema.SingleNestedAttribute{
													Description:         "HeaderHashOptions should be set when request header hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
													MarkdownDescription: "HeaderHashOptions should be set when request header hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
													Attributes: map[string]schema.Attribute{
														"header_name": schema.StringAttribute{
															Description:         "HeaderName is the name of the HTTP request header that will be used to calculate the hash key. If the header specified is not present on a request, no hash will be produced.",
															MarkdownDescription: "HeaderName is the name of the HTTP request header that will be used to calculate the hash key. If the header specified is not present on a request, no hash will be produced.",
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

												"query_parameter_hash_options": schema.SingleNestedAttribute{
													Description:         "QueryParameterHashOptions should be set when request query parameter hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
													MarkdownDescription: "QueryParameterHashOptions should be set when request query parameter hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
													Attributes: map[string]schema.Attribute{
														"parameter_name": schema.StringAttribute{
															Description:         "ParameterName is the name of the HTTP request query parameter that will be used to calculate the hash key. If the query parameter specified is not present on a request, no hash will be produced.",
															MarkdownDescription: "ParameterName is the name of the HTTP request query parameter that will be used to calculate the hash key. If the query parameter specified is not present on a request, no hash will be produced.",
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

												"terminal": schema.BoolAttribute{
													Description:         "Terminal is a flag that allows for short-circuiting computing of a hash for a given request. If set to true, and the request attribute specified in the attribute hash options is present, no further hash policies will be used to calculate a hash for the request.",
													MarkdownDescription: "Terminal is a flag that allows for short-circuiting computing of a hash for a given request. If set to true, and the request attribute specified in the attribute hash options is present, no further hash policies will be used to calculate a hash for the request.",
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

									"strategy": schema.StringAttribute{
										Description:         "Strategy specifies the policy used to balance requests across the pool of backend pods. Valid policy names are 'Random', 'RoundRobin', 'WeightedLeastRequest', 'Cookie', and 'RequestHash'. If an unknown strategy name is specified or no policy is supplied, the default 'RoundRobin' policy is used.",
										MarkdownDescription: "Strategy specifies the policy used to balance requests across the pool of backend pods. Valid policy names are 'Random', 'RoundRobin', 'WeightedLeastRequest', 'Cookie', and 'RequestHash'. If an unknown strategy name is specified or no policy is supplied, the default 'RoundRobin' policy is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"services": schema.ListNestedAttribute{
								Description:         "Services are the services to proxy traffic",
								MarkdownDescription: "Services are the services to proxy traffic",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cookie_rewrite_policies": schema.ListNestedAttribute{
											Description:         "The policies for rewriting Set-Cookie header attributes.",
											MarkdownDescription: "The policies for rewriting Set-Cookie header attributes.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"domain_rewrite": schema.SingleNestedAttribute{
														Description:         "DomainRewrite enables rewriting the Set-Cookie Domain element. If not set, Domain will not be rewritten.",
														MarkdownDescription: "DomainRewrite enables rewriting the Set-Cookie Domain element. If not set, Domain will not be rewritten.",
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Description:         "Value is the value to rewrite the Domain attribute to. For now this is required.",
																MarkdownDescription: "Value is the value to rewrite the Domain attribute to. For now this is required.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(4096),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the cookie for which attributes will be rewritten.",
														MarkdownDescription: "Name is the name of the cookie for which attributes will be rewritten.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(4096),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[^()<>@,;:\\"\/[\]?={} \t\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`), ""),
														},
													},

													"path_rewrite": schema.SingleNestedAttribute{
														Description:         "PathRewrite enables rewriting the Set-Cookie Path element. If not set, Path will not be rewritten.",
														MarkdownDescription: "PathRewrite enables rewriting the Set-Cookie Path element. If not set, Path will not be rewritten.",
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Description:         "Value is the value to rewrite the Path attribute to. For now this is required.",
																MarkdownDescription: "Value is the value to rewrite the Path attribute to. For now this is required.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(4096),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[^;\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`), ""),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"same_site": schema.StringAttribute{
														Description:         "SameSite enables rewriting the Set-Cookie SameSite element. If not set, SameSite attribute will not be rewritten.",
														MarkdownDescription: "SameSite enables rewriting the Set-Cookie SameSite element. If not set, SameSite attribute will not be rewritten.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Strict", "Lax", "None"),
														},
													},

													"secure": schema.BoolAttribute{
														Description:         "Secure enables rewriting the Set-Cookie Secure element. If not set, Secure attribute will not be rewritten.",
														MarkdownDescription: "Secure enables rewriting the Set-Cookie Secure element. If not set, Secure attribute will not be rewritten.",
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

										"health_port": schema.Int64Attribute{
											Description:         "HealthPort is the port for this service healthcheck. If not specified, Port is used for service healthchecks.",
											MarkdownDescription: "HealthPort is the port for this service healthcheck. If not specified, Port is used for service healthchecks.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"mirror": schema.BoolAttribute{
											Description:         "If Mirror is true the Service will receive a read only mirror of the traffic for this route. If Mirror is true, then fractional mirroring can be enabled by optionally setting the Weight field. Legal values for Weight are 1-100. Omitting the Weight field will result in 100% mirroring. NOTE: Setting Weight explicitly to 0 will unexpectedly result in 100% traffic mirroring. This occurs since we cannot distinguish omitted fields from those explicitly set to their default values",
											MarkdownDescription: "If Mirror is true the Service will receive a read only mirror of the traffic for this route. If Mirror is true, then fractional mirroring can be enabled by optionally setting the Weight field. Legal values for Weight are 1-100. Omitting the Weight field will result in 100% mirroring. NOTE: Setting Weight explicitly to 0 will unexpectedly result in 100% traffic mirroring. This occurs since we cannot distinguish omitted fields from those explicitly set to their default values",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of Kubernetes service to proxy traffic. Names defined here will be used to look up corresponding endpoints which contain the ips to route.",
											MarkdownDescription: "Name is the name of Kubernetes service to proxy traffic. Names defined here will be used to look up corresponding endpoints which contain the ips to route.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port (defined as Integer) to proxy traffic to since a service can have multiple defined.",
											MarkdownDescription: "Port (defined as Integer) to proxy traffic to since a service can have multiple defined.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol may be used to specify (or override) the protocol used to reach this Service. Values may be tls, h2, h2c. If omitted, protocol-selection falls back on Service annotations.",
											MarkdownDescription: "Protocol may be used to specify (or override) the protocol used to reach this Service. Values may be tls, h2, h2c. If omitted, protocol-selection falls back on Service annotations.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("h2", "h2c", "tls"),
											},
										},

										"request_headers_policy": schema.SingleNestedAttribute{
											Description:         "The policy for managing request headers during proxying.",
											MarkdownDescription: "The policy for managing request headers during proxying.",
											Attributes: map[string]schema.Attribute{
												"remove": schema.ListAttribute{
													Description:         "Remove specifies a list of HTTP header names to remove.",
													MarkdownDescription: "Remove specifies a list of HTTP header names to remove.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"set": schema.ListNestedAttribute{
													Description:         "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
													MarkdownDescription: "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name represents a key of a header",
																MarkdownDescription: "Name represents a key of a header",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"value": schema.StringAttribute{
																Description:         "Value represents the value of a header specified by a key",
																MarkdownDescription: "Value represents the value of a header specified by a key",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
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

										"response_headers_policy": schema.SingleNestedAttribute{
											Description:         "The policy for managing response headers during proxying. Rewriting the 'Host' header is not supported.",
											MarkdownDescription: "The policy for managing response headers during proxying. Rewriting the 'Host' header is not supported.",
											Attributes: map[string]schema.Attribute{
												"remove": schema.ListAttribute{
													Description:         "Remove specifies a list of HTTP header names to remove.",
													MarkdownDescription: "Remove specifies a list of HTTP header names to remove.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"set": schema.ListNestedAttribute{
													Description:         "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
													MarkdownDescription: "Set specifies a list of HTTP header values that will be set in the HTTP header. If the header does not exist it will be added, otherwise it will be overwritten with the new value.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name represents a key of a header",
																MarkdownDescription: "Name represents a key of a header",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"value": schema.StringAttribute{
																Description:         "Value represents the value of a header specified by a key",
																MarkdownDescription: "Value represents the value of a header specified by a key",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
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

										"slow_start_policy": schema.SingleNestedAttribute{
											Description:         "Slow start will gradually increase amount of traffic to a newly added endpoint.",
											MarkdownDescription: "Slow start will gradually increase amount of traffic to a newly added endpoint.",
											Attributes: map[string]schema.Attribute{
												"aggression": schema.StringAttribute{
													Description:         "The speed of traffic increase over the slow start window. Defaults to 1.0, so that endpoint would get linearly increasing amount of traffic. When increasing the value for this parameter, the speed of traffic ramp-up increases non-linearly. The value of aggression parameter should be greater than 0.0. More info: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/slow_start",
													MarkdownDescription: "The speed of traffic increase over the slow start window. Defaults to 1.0, so that endpoint would get linearly increasing amount of traffic. When increasing the value for this parameter, the speed of traffic ramp-up increases non-linearly. The value of aggression parameter should be greater than 0.0. More info: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/slow_start",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+([.][0-9]+)?|[.][0-9]+)$`), ""),
													},
												},

												"min_weight_percent": schema.Int64Attribute{
													Description:         "The minimum or starting percentage of traffic to send to new endpoints. A non-zero value helps avoid a too small initial weight, which may cause endpoints in slow start mode to receive no traffic in the beginning of the slow start window. If not specified, the default is 10%.",
													MarkdownDescription: "The minimum or starting percentage of traffic to send to new endpoints. A non-zero value helps avoid a too small initial weight, which may cause endpoints in slow start mode to receive no traffic in the beginning of the slow start window. If not specified, the default is 10%.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(100),
													},
												},

												"window": schema.StringAttribute{
													Description:         "The duration of slow start window. Duration is expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													MarkdownDescription: "The duration of slow start window. Duration is expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"validation": schema.SingleNestedAttribute{
											Description:         "UpstreamValidation defines how to verify the backend service's certificate",
											MarkdownDescription: "UpstreamValidation defines how to verify the backend service's certificate",
											Attributes: map[string]schema.Attribute{
												"ca_secret": schema.StringAttribute{
													Description:         "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
													MarkdownDescription: "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(317),
													},
												},

												"subject_name": schema.StringAttribute{
													Description:         "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
													MarkdownDescription: "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(250),
													},
												},

												"subject_names": schema.ListAttribute{
													Description:         "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
													MarkdownDescription: "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
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

										"weight": schema.Int64Attribute{
											Description:         "Weight defines percentage of traffic to balance traffic",
											MarkdownDescription: "Weight defines percentage of traffic to balance traffic",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
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

					"virtualhost": schema.SingleNestedAttribute{
						Description:         "Virtualhost appears at most once. If it is present, the object is considered to be a 'root' HTTPProxy.",
						MarkdownDescription: "Virtualhost appears at most once. If it is present, the object is considered to be a 'root' HTTPProxy.",
						Attributes: map[string]schema.Attribute{
							"authorization": schema.SingleNestedAttribute{
								Description:         "This field configures an extension service to perform authorization for this virtual host. Authorization can only be configured on virtual hosts that have TLS enabled. If the TLS configuration requires client certificate validation, the client certificate is always included in the authentication check request.",
								MarkdownDescription: "This field configures an extension service to perform authorization for this virtual host. Authorization can only be configured on virtual hosts that have TLS enabled. If the TLS configuration requires client certificate validation, the client certificate is always included in the authentication check request.",
								Attributes: map[string]schema.Attribute{
									"auth_policy": schema.SingleNestedAttribute{
										Description:         "AuthPolicy sets a default authorization policy for client requests. This policy will be used unless overridden by individual routes.",
										MarkdownDescription: "AuthPolicy sets a default authorization policy for client requests. This policy will be used unless overridden by individual routes.",
										Attributes: map[string]schema.Attribute{
											"context": schema.MapAttribute{
												Description:         "Context is a set of key/value pairs that are sent to the authentication server in the check request. If a context is provided at an enclosing scope, the entries are merged such that the inner scope overrides matching keys from the outer scope.",
												MarkdownDescription: "Context is a set of key/value pairs that are sent to the authentication server in the check request. If a context is provided at an enclosing scope, the entries are merged such that the inner scope overrides matching keys from the outer scope.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disabled": schema.BoolAttribute{
												Description:         "When true, this field disables client request authentication for the scope of the policy.",
												MarkdownDescription: "When true, this field disables client request authentication for the scope of the policy.",
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
												Description:         "API version of the referent. If this field is not specified, the default 'projectcontour.io/v1alpha1' will be used",
												MarkdownDescription: "API version of the referent. If this field is not specified, the default 'projectcontour.io/v1alpha1' will be used",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent. If this field is not specifies, the namespace of the resource that targets the referent will be used. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent. If this field is not specifies, the namespace of the resource that targets the referent will be used. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
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
										Description:         "If FailOpen is true, the client request is forwarded to the upstream service even if the authorization server fails to respond. This field should not be set in most cases. It is intended for use only while migrating applications from internal authorization to Contour external authorization.",
										MarkdownDescription: "If FailOpen is true, the client request is forwarded to the upstream service even if the authorization server fails to respond. This field should not be set in most cases. It is intended for use only while migrating applications from internal authorization to Contour external authorization.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_timeout": schema.StringAttribute{
										Description:         "ResponseTimeout configures maximum time to wait for a check response from the authorization server. Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'. The string 'infinity' is also a valid input and specifies no timeout.",
										MarkdownDescription: "ResponseTimeout configures maximum time to wait for a check response from the authorization server. Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'. The string 'infinity' is also a valid input and specifies no timeout.",
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

							"cors_policy": schema.SingleNestedAttribute{
								Description:         "Specifies the cross-origin policy to apply to the VirtualHost.",
								MarkdownDescription: "Specifies the cross-origin policy to apply to the VirtualHost.",
								Attributes: map[string]schema.Attribute{
									"allow_credentials": schema.BoolAttribute{
										Description:         "Specifies whether the resource allows credentials.",
										MarkdownDescription: "Specifies whether the resource allows credentials.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_headers": schema.ListAttribute{
										Description:         "AllowHeaders specifies the content for the *access-control-allow-headers* header.",
										MarkdownDescription: "AllowHeaders specifies the content for the *access-control-allow-headers* header.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_methods": schema.ListAttribute{
										Description:         "AllowMethods specifies the content for the *access-control-allow-methods* header.",
										MarkdownDescription: "AllowMethods specifies the content for the *access-control-allow-methods* header.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"allow_origin": schema.ListAttribute{
										Description:         "AllowOrigin specifies the origins that will be allowed to do CORS requests. Allowed values include '*' which signifies any origin is allowed, an exact origin of the form 'scheme://host[:port]' (where port is optional), or a valid regex pattern. Note that regex patterns are validated and a simple 'glob' pattern (e.g. *.foo.com) will be rejected or produce unexpected matches when applied as a regex.",
										MarkdownDescription: "AllowOrigin specifies the origins that will be allowed to do CORS requests. Allowed values include '*' which signifies any origin is allowed, an exact origin of the form 'scheme://host[:port]' (where port is optional), or a valid regex pattern. Note that regex patterns are validated and a simple 'glob' pattern (e.g. *.foo.com) will be rejected or produce unexpected matches when applied as a regex.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"allow_private_network": schema.BoolAttribute{
										Description:         "AllowPrivateNetwork specifies whether to allow private network requests. See https://developer.chrome.com/blog/private-network-access-preflight.",
										MarkdownDescription: "AllowPrivateNetwork specifies whether to allow private network requests. See https://developer.chrome.com/blog/private-network-access-preflight.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"expose_headers": schema.ListAttribute{
										Description:         "ExposeHeaders Specifies the content for the *access-control-expose-headers* header.",
										MarkdownDescription: "ExposeHeaders Specifies the content for the *access-control-expose-headers* header.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_age": schema.StringAttribute{
										Description:         "MaxAge indicates for how long the results of a preflight request can be cached. MaxAge durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'. Only positive values are allowed while 0 disables the cache requiring a preflight OPTIONS check for all cross-origin requests.",
										MarkdownDescription: "MaxAge indicates for how long the results of a preflight request can be cached. MaxAge durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration). Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'. Only positive values are allowed while 0 disables the cache requiring a preflight OPTIONS check for all cross-origin requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|0)$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"fqdn": schema.StringAttribute{
								Description:         "The fully qualified domain name of the root of the ingress tree all leaves of the DAG rooted at this object relate to the fqdn.",
								MarkdownDescription: "The fully qualified domain name of the root of the ingress tree all leaves of the DAG rooted at this object relate to the fqdn.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\*\.)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"ip_allow_policy": schema.ListNestedAttribute{
								Description:         "IPAllowFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be allowed. All other requests will be denied. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here may be overridden in a Route.",
								MarkdownDescription: "IPAllowFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be allowed. All other requests will be denied. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here may be overridden in a Route.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cidr": schema.StringAttribute{
											Description:         "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
											MarkdownDescription: "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"source": schema.StringAttribute{
											Description:         "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
											MarkdownDescription: "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Peer", "Remote"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_deny_policy": schema.ListNestedAttribute{
								Description:         "IPDenyFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be denied. All other requests will be allowed. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here may be overridden in a Route.",
								MarkdownDescription: "IPDenyFilterPolicy is a list of ipv4/6 filter rules for which matching requests should be denied. All other requests will be allowed. Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined. The rules defined here may be overridden in a Route.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cidr": schema.StringAttribute{
											Description:         "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
											MarkdownDescription: "CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be a bare IP address (without a mask) to filter on exactly one address.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"source": schema.StringAttribute{
											Description:         "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
											MarkdownDescription: "Source indicates how to determine the ip address to filter on, and can be one of two values: - 'Remote' filters on the ip address of the client, accounting for PROXY and X-Forwarded-For as needed. - 'Peer' filters on the ip of the network request, ignoring PROXY and X-Forwarded-For.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Peer", "Remote"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"jwt_providers": schema.ListNestedAttribute{
								Description:         "Providers to use for verifying JSON Web Tokens (JWTs) on the virtual host.",
								MarkdownDescription: "Providers to use for verifying JSON Web Tokens (JWTs) on the virtual host.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"audiences": schema.ListAttribute{
											Description:         "Audiences that JWTs are allowed to have in the 'aud' field. If not provided, JWT audiences are not checked.",
											MarkdownDescription: "Audiences that JWTs are allowed to have in the 'aud' field. If not provided, JWT audiences are not checked.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default": schema.BoolAttribute{
											Description:         "Whether the provider should apply to all routes in the HTTPProxy/its includes by default. At most one provider can be marked as the default. If no provider is marked as the default, individual routes must explicitly identify the provider they require.",
											MarkdownDescription: "Whether the provider should apply to all routes in the HTTPProxy/its includes by default. At most one provider can be marked as the default. If no provider is marked as the default, individual routes must explicitly identify the provider they require.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"forward_jwt": schema.BoolAttribute{
											Description:         "Whether the JWT should be forwarded to the backend service after successful verification. By default, the JWT is not forwarded.",
											MarkdownDescription: "Whether the JWT should be forwarded to the backend service after successful verification. By default, the JWT is not forwarded.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"issuer": schema.StringAttribute{
											Description:         "Issuer that JWTs are required to have in the 'iss' field. If not provided, JWT issuers are not checked.",
											MarkdownDescription: "Issuer that JWTs are required to have in the 'iss' field. If not provided, JWT issuers are not checked.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Unique name for the provider.",
											MarkdownDescription: "Unique name for the provider.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},

										"remote_jwks": schema.SingleNestedAttribute{
											Description:         "Remote JWKS to use for verifying JWT signatures.",
											MarkdownDescription: "Remote JWKS to use for verifying JWT signatures.",
											Attributes: map[string]schema.Attribute{
												"cache_duration": schema.StringAttribute{
													Description:         "How long to cache the JWKS locally. If not specified, Envoy's default of 5m applies.",
													MarkdownDescription: "How long to cache the JWKS locally. If not specified, Envoy's default of 5m applies.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`), ""),
													},
												},

												"dns_lookup_family": schema.StringAttribute{
													Description:         "The DNS IP address resolution policy for the JWKS URI. When configured as 'v4', the DNS resolver will only perform a lookup for addresses in the IPv4 family. If 'v6' is configured, the DNS resolver will only perform a lookup for addresses in the IPv6 family. If 'all' is configured, the DNS resolver will perform a lookup for addresses in both the IPv4 and IPv6 family. If 'auto' is configured, the DNS resolver will first perform a lookup for addresses in the IPv6 family and fallback to a lookup for addresses in the IPv4 family. If not specified, the Contour-wide setting defined in the config file or ContourConfiguration applies (defaults to 'auto'). See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamily for more information.",
													MarkdownDescription: "The DNS IP address resolution policy for the JWKS URI. When configured as 'v4', the DNS resolver will only perform a lookup for addresses in the IPv4 family. If 'v6' is configured, the DNS resolver will only perform a lookup for addresses in the IPv6 family. If 'all' is configured, the DNS resolver will perform a lookup for addresses in both the IPv4 and IPv6 family. If 'auto' is configured, the DNS resolver will first perform a lookup for addresses in the IPv6 family and fallback to a lookup for addresses in the IPv4 family. If not specified, the Contour-wide setting defined in the config file or ContourConfiguration applies (defaults to 'auto'). See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamily for more information.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("auto", "v4", "v6"),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "How long to wait for a response from the URI. If not specified, a default of 1s applies.",
													MarkdownDescription: "How long to wait for a response from the URI. If not specified, a default of 1s applies.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`), ""),
													},
												},

												"uri": schema.StringAttribute{
													Description:         "The URI for the JWKS.",
													MarkdownDescription: "The URI for the JWKS.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},

												"validation": schema.SingleNestedAttribute{
													Description:         "UpstreamValidation defines how to verify the JWKS's TLS certificate.",
													MarkdownDescription: "UpstreamValidation defines how to verify the JWKS's TLS certificate.",
													Attributes: map[string]schema.Attribute{
														"ca_secret": schema.StringAttribute{
															Description:         "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
															MarkdownDescription: "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(317),
															},
														},

														"subject_name": schema.StringAttribute{
															Description:         "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
															MarkdownDescription: "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(250),
															},
														},

														"subject_names": schema.ListAttribute{
															Description:         "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
															MarkdownDescription: "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rate_limit_policy": schema.SingleNestedAttribute{
								Description:         "The policy for rate limiting on the virtual host.",
								MarkdownDescription: "The policy for rate limiting on the virtual host.",
								Attributes: map[string]schema.Attribute{
									"global": schema.SingleNestedAttribute{
										Description:         "Global defines global rate limiting parameters, i.e. parameters defining descriptors that are sent to an external rate limit service (RLS) for a rate limit decision on each request.",
										MarkdownDescription: "Global defines global rate limiting parameters, i.e. parameters defining descriptors that are sent to an external rate limit service (RLS) for a rate limit decision on each request.",
										Attributes: map[string]schema.Attribute{
											"descriptors": schema.ListNestedAttribute{
												Description:         "Descriptors defines the list of descriptors that will be generated and sent to the rate limit service. Each descriptor contains 1+ key-value pair entries.",
												MarkdownDescription: "Descriptors defines the list of descriptors that will be generated and sent to the rate limit service. Each descriptor contains 1+ key-value pair entries.",
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
																				Description:         "Key defines the key of the descriptor entry. If not set, the key is set to 'generic_key'.",
																				MarkdownDescription: "Key defines the key of the descriptor entry. If not set, the key is set to 'generic_key'.",
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
																		Description:         "RemoteAddress defines a descriptor entry with a key of 'remote_address' and a value equal to the client's IP address (from x-forwarded-for).",
																		MarkdownDescription: "RemoteAddress defines a descriptor entry with a key of 'remote_address' and a value equal to the client's IP address (from x-forwarded-for).",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"request_header": schema.SingleNestedAttribute{
																		Description:         "RequestHeader defines a descriptor entry that's populated only if a given header is present on the request. The descriptor key is static, and the descriptor value is equal to the value of the header.",
																		MarkdownDescription: "RequestHeader defines a descriptor entry that's populated only if a given header is present on the request. The descriptor key is static, and the descriptor value is equal to the value of the header.",
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
																		Description:         "RequestHeaderValueMatch defines a descriptor entry that's populated if the request's headers match a set of 1+ match criteria. The descriptor key is 'header_match', and the descriptor value is static.",
																		MarkdownDescription: "RequestHeaderValueMatch defines a descriptor entry that's populated if the request's headers match a set of 1+ match criteria. The descriptor key is 'header_match', and the descriptor value is static.",
																		Attributes: map[string]schema.Attribute{
																			"expect_match": schema.BoolAttribute{
																				Description:         "ExpectMatch defines whether the request must positively match the match criteria in order to generate a descriptor entry (i.e. true), or not match the match criteria in order to generate a descriptor entry (i.e. false). The default is true.",
																				MarkdownDescription: "ExpectMatch defines whether the request must positively match the match criteria in order to generate a descriptor entry (i.e. true), or not match the match criteria in order to generate a descriptor entry (i.e. false). The default is true.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "Headers is a list of 1+ match criteria to apply against the request to determine whether to populate the descriptor entry or not.",
																				MarkdownDescription: "Headers is a list of 1+ match criteria to apply against the request to determine whether to populate the descriptor entry or not.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"contains": schema.StringAttribute{
																							Description:         "Contains specifies a substring that must be present in the header value.",
																							MarkdownDescription: "Contains specifies a substring that must be present in the header value.",
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
																							Description:         "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
																							MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive. Note that this has no effect on the Regex parameter.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
																							MarkdownDescription: "Name is the name of the header to match against. Name is required. Header names are case insensitive.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"notcontains": schema.StringAttribute{
																							Description:         "NotContains specifies a substring that must not be present in the header value.",
																							MarkdownDescription: "NotContains specifies a substring that must not be present in the header value.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"notexact": schema.StringAttribute{
																							Description:         "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
																							MarkdownDescription: "NoExact specifies a string that the header value must not be equal to. The condition is true if the header has any other value.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"notpresent": schema.BoolAttribute{
																							Description:         "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
																							MarkdownDescription: "NotPresent specifies that condition is true when the named header is not present. Note that setting NotPresent to false does not make the condition true if the named header is present.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"present": schema.BoolAttribute{
																							Description:         "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
																							MarkdownDescription: "Present specifies that condition is true when the named header is present, regardless of its value. Note that setting Present to false does not make the condition true if the named header is absent.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "Regex specifies a regular expression pattern that must match the header value.",
																							MarkdownDescription: "Regex specifies a regular expression pattern that must match the header value.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"treat_missing_as_empty": schema.BoolAttribute{
																							Description:         "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
																							MarkdownDescription: "TreatMissingAsEmpty specifies if the header match rule specified header does not exist, this header value will be treated as empty. Defaults to false. Unlike the underlying Envoy implementation this is **only** supported for negative matches (e.g. NotContains, NotExact).",
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
												Description:         "Disabled configures the HTTPProxy to not use the default global rate limit policy defined by the Contour configuration.",
												MarkdownDescription: "Disabled configures the HTTPProxy to not use the default global rate limit policy defined by the Contour configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"local": schema.SingleNestedAttribute{
										Description:         "Local defines local rate limiting parameters, i.e. parameters for rate limiting that occurs within each Envoy pod as requests are handled.",
										MarkdownDescription: "Local defines local rate limiting parameters, i.e. parameters for rate limiting that occurs within each Envoy pod as requests are handled.",
										Attributes: map[string]schema.Attribute{
											"burst": schema.Int64Attribute{
												Description:         "Burst defines the number of requests above the requests per unit that should be allowed within a short period of time.",
												MarkdownDescription: "Burst defines the number of requests above the requests per unit that should be allowed within a short period of time.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.Int64Attribute{
												Description:         "Requests defines how many requests per unit of time should be allowed before rate limiting occurs.",
												MarkdownDescription: "Requests defines how many requests per unit of time should be allowed before rate limiting occurs.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"response_headers_to_add": schema.ListNestedAttribute{
												Description:         "ResponseHeadersToAdd is an optional list of response headers to set when a request is rate-limited.",
												MarkdownDescription: "ResponseHeadersToAdd is an optional list of response headers to set when a request is rate-limited.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name represents a key of a header",
															MarkdownDescription: "Name represents a key of a header",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value represents the value of a header specified by a key",
															MarkdownDescription: "Value represents the value of a header specified by a key",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_status_code": schema.Int64Attribute{
												Description:         "ResponseStatusCode is the HTTP status code to use for responses to rate-limited requests. Codes must be in the 400-599 range (inclusive). If not specified, the Envoy default of 429 (Too Many Requests) is used.",
												MarkdownDescription: "ResponseStatusCode is the HTTP status code to use for responses to rate-limited requests. Codes must be in the 400-599 range (inclusive). If not specified, the Envoy default of 429 (Too Many Requests) is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(400),
													int64validator.AtMost(599),
												},
											},

											"unit": schema.StringAttribute{
												Description:         "Unit defines the period of time within which requests over the limit will be rate limited. Valid values are 'second', 'minute' and 'hour'.",
												MarkdownDescription: "Unit defines the period of time within which requests over the limit will be rate limited. Valid values are 'second', 'minute' and 'hour'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("second", "minute", "hour"),
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

							"tls": schema.SingleNestedAttribute{
								Description:         "If present the fields describes TLS properties of the virtual host. The SNI names that will be matched on are described in fqdn, the tls.secretName secret must contain a certificate that itself contains a name that matches the FQDN.",
								MarkdownDescription: "If present the fields describes TLS properties of the virtual host. The SNI names that will be matched on are described in fqdn, the tls.secretName secret must contain a certificate that itself contains a name that matches the FQDN.",
								Attributes: map[string]schema.Attribute{
									"client_validation": schema.SingleNestedAttribute{
										Description:         "ClientValidation defines how to verify the client certificate when an external client establishes a TLS connection to Envoy. This setting: 1. Enables TLS client certificate validation. 2. Specifies how the client certificate will be validated (i.e. validation required or skipped). Note: Setting client certificate validation to be skipped should be only used in conjunction with an external authorization server that performs client validation as Contour will ensure client certificates are passed along.",
										MarkdownDescription: "ClientValidation defines how to verify the client certificate when an external client establishes a TLS connection to Envoy. This setting: 1. Enables TLS client certificate validation. 2. Specifies how the client certificate will be validated (i.e. validation required or skipped). Note: Setting client certificate validation to be skipped should be only used in conjunction with an external authorization server that performs client validation as Contour will ensure client certificates are passed along.",
										Attributes: map[string]schema.Attribute{
											"ca_secret": schema.StringAttribute{
												Description:         "Name of a Kubernetes secret that contains a CA certificate bundle. The secret must contain key named ca.crt. The client certificate must validate against the certificates in the bundle. If specified and SkipClientCertValidation is true, client certificates will be required on requests. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.",
												MarkdownDescription: "Name of a Kubernetes secret that contains a CA certificate bundle. The secret must contain key named ca.crt. The client certificate must validate against the certificates in the bundle. If specified and SkipClientCertValidation is true, client certificates will be required on requests. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"crl_only_verify_leaf_cert": schema.BoolAttribute{
												Description:         "If this option is set to true, only the certificate at the end of the certificate chain will be subject to validation by CRL.",
												MarkdownDescription: "If this option is set to true, only the certificate at the end of the certificate chain will be subject to validation by CRL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"crl_secret": schema.StringAttribute{
												Description:         "Name of a Kubernetes opaque secret that contains a concatenated list of PEM encoded CRLs. The secret must contain key named crl.pem. This field will be used to verify that a client certificate has not been revoked. CRLs must be available from all CAs, unless crlOnlyVerifyLeafCert is true. Large CRL lists are not supported since individual secrets are limited to 1MiB in size. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.",
												MarkdownDescription: "Name of a Kubernetes opaque secret that contains a concatenated list of PEM encoded CRLs. The secret must contain key named crl.pem. This field will be used to verify that a client certificate has not been revoked. CRLs must be available from all CAs, unless crlOnlyVerifyLeafCert is true. Large CRL lists are not supported since individual secrets are limited to 1MiB in size. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"forward_client_certificate": schema.SingleNestedAttribute{
												Description:         "ForwardClientCertificate adds the selected data from the passed client TLS certificate to the x-forwarded-client-cert header.",
												MarkdownDescription: "ForwardClientCertificate adds the selected data from the passed client TLS certificate to the x-forwarded-client-cert header.",
												Attributes: map[string]schema.Attribute{
													"cert": schema.BoolAttribute{
														Description:         "Client cert in URL encoded PEM format.",
														MarkdownDescription: "Client cert in URL encoded PEM format.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"chain": schema.BoolAttribute{
														Description:         "Client cert chain (including the leaf cert) in URL encoded PEM format.",
														MarkdownDescription: "Client cert chain (including the leaf cert) in URL encoded PEM format.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"dns": schema.BoolAttribute{
														Description:         "DNS type Subject Alternative Names of the client cert.",
														MarkdownDescription: "DNS type Subject Alternative Names of the client cert.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subject": schema.BoolAttribute{
														Description:         "Subject of the client cert.",
														MarkdownDescription: "Subject of the client cert.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"uri": schema.BoolAttribute{
														Description:         "URI type Subject Alternative Name of the client cert.",
														MarkdownDescription: "URI type Subject Alternative Name of the client cert.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional_client_certificate": schema.BoolAttribute{
												Description:         "OptionalClientCertificate when set to true will request a client certificate but allow the connection to continue if the client does not provide one. If a client certificate is sent, it will be verified according to the other properties, which includes disabling validation if SkipClientCertValidation is set. Defaults to false.",
												MarkdownDescription: "OptionalClientCertificate when set to true will request a client certificate but allow the connection to continue if the client does not provide one. If a client certificate is sent, it will be verified according to the other properties, which includes disabling validation if SkipClientCertValidation is set. Defaults to false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"skip_client_cert_validation": schema.BoolAttribute{
												Description:         "SkipClientCertValidation disables downstream client certificate validation. Defaults to false. This field is intended to be used in conjunction with external authorization in order to enable the external authorization server to validate client certificates. When this field is set to true, client certificates are requested but not verified by Envoy. If CACertificate is specified, client certificates are required on requests, but not verified. If external authorization is in use, they are presented to the external authorization server.",
												MarkdownDescription: "SkipClientCertValidation disables downstream client certificate validation. Defaults to false. This field is intended to be used in conjunction with external authorization in order to enable the external authorization server to validate client certificates. When this field is set to true, client certificates are requested but not verified by Envoy. If CACertificate is specified, client certificates are required on requests, but not verified. If external authorization is in use, they are presented to the external authorization server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable_fallback_certificate": schema.BoolAttribute{
										Description:         "EnableFallbackCertificate defines if the vhost should allow a default certificate to be applied which handles all requests which don't match the SNI defined in this vhost.",
										MarkdownDescription: "EnableFallbackCertificate defines if the vhost should allow a default certificate to be applied which handles all requests which don't match the SNI defined in this vhost.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_protocol_version": schema.StringAttribute{
										Description:         "MaximumProtocolVersion is the maximum TLS version this vhost should negotiate. Valid options are '1.2' and '1.3' (default). Any other value defaults to TLS 1.3.",
										MarkdownDescription: "MaximumProtocolVersion is the maximum TLS version this vhost should negotiate. Valid options are '1.2' and '1.3' (default). Any other value defaults to TLS 1.3.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"minimum_protocol_version": schema.StringAttribute{
										Description:         "MinimumProtocolVersion is the minimum TLS version this vhost should negotiate. Valid options are '1.2' (default) and '1.3'. Any other value defaults to TLS 1.2.",
										MarkdownDescription: "MinimumProtocolVersion is the minimum TLS version this vhost should negotiate. Valid options are '1.2' (default) and '1.3'. Any other value defaults to TLS 1.2.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"passthrough": schema.BoolAttribute{
										Description:         "Passthrough defines whether the encrypted TLS handshake will be passed through to the backing cluster. Either Passthrough or SecretName must be specified, but not both.",
										MarkdownDescription: "Passthrough defines whether the encrypted TLS handshake will be passed through to the backing cluster. Either Passthrough or SecretName must be specified, but not both.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "SecretName is the name of a TLS secret. Either SecretName or Passthrough must be specified, but not both. If specified, the named secret must contain a matching certificate for the virtual host's FQDN. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.",
										MarkdownDescription: "SecretName is the name of a TLS secret. Either SecretName or Passthrough must be specified, but not both. If specified, the named secret must contain a matching certificate for the virtual host's FQDN. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.",
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
		},
	}
}

func (r *ProjectcontourIoHttpproxyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_projectcontour_io_http_proxy_v1_manifest")

	var model ProjectcontourIoHttpproxyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("projectcontour.io/v1")
	model.Kind = pointer.String("HTTPProxy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
