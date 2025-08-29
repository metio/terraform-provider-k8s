/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_nginx_org_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &K8SNginxOrgPolicyV1Manifest{}
)

func NewK8SNginxOrgPolicyV1Manifest() datasource.DataSource {
	return &K8SNginxOrgPolicyV1Manifest{}
}

type K8SNginxOrgPolicyV1Manifest struct{}

type K8SNginxOrgPolicyV1ManifestData struct {
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
		AccessControl *struct {
			Allow *[]string `tfsdk:"allow" json:"allow,omitempty"`
			Deny  *[]string `tfsdk:"deny" json:"deny,omitempty"`
		} `tfsdk:"access_control" json:"accessControl,omitempty"`
		ApiKey *struct {
			ClientSecret *string `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			SuppliedIn   *struct {
				Header *[]string `tfsdk:"header" json:"header,omitempty"`
				Query  *[]string `tfsdk:"query" json:"query,omitempty"`
			} `tfsdk:"supplied_in" json:"suppliedIn,omitempty"`
		} `tfsdk:"api_key" json:"apiKey,omitempty"`
		BasicAuth *struct {
			Realm  *string `tfsdk:"realm" json:"realm,omitempty"`
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
		Cache *struct {
			AllowedCodes          *[]string `tfsdk:"allowed_codes" json:"allowedCodes,omitempty"`
			AllowedMethods        *[]string `tfsdk:"allowed_methods" json:"allowedMethods,omitempty"`
			CachePurgeAllow       *[]string `tfsdk:"cache_purge_allow" json:"cachePurgeAllow,omitempty"`
			CacheZoneName         *string   `tfsdk:"cache_zone_name" json:"cacheZoneName,omitempty"`
			CacheZoneSize         *string   `tfsdk:"cache_zone_size" json:"cacheZoneSize,omitempty"`
			Levels                *string   `tfsdk:"levels" json:"levels,omitempty"`
			OverrideUpstreamCache *bool     `tfsdk:"override_upstream_cache" json:"overrideUpstreamCache,omitempty"`
			Time                  *string   `tfsdk:"time" json:"time,omitempty"`
		} `tfsdk:"cache" json:"cache,omitempty"`
		EgressMTLS *struct {
			Ciphers           *string `tfsdk:"ciphers" json:"ciphers,omitempty"`
			Protocols         *string `tfsdk:"protocols" json:"protocols,omitempty"`
			ServerName        *bool   `tfsdk:"server_name" json:"serverName,omitempty"`
			SessionReuse      *bool   `tfsdk:"session_reuse" json:"sessionReuse,omitempty"`
			SslName           *string `tfsdk:"ssl_name" json:"sslName,omitempty"`
			TlsSecret         *string `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
			TrustedCertSecret *string `tfsdk:"trusted_cert_secret" json:"trustedCertSecret,omitempty"`
			VerifyDepth       *int64  `tfsdk:"verify_depth" json:"verifyDepth,omitempty"`
			VerifyServer      *bool   `tfsdk:"verify_server" json:"verifyServer,omitempty"`
		} `tfsdk:"egress_mtls" json:"egressMTLS,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		IngressMTLS      *struct {
			ClientCertSecret *string `tfsdk:"client_cert_secret" json:"clientCertSecret,omitempty"`
			CrlFileName      *string `tfsdk:"crl_file_name" json:"crlFileName,omitempty"`
			VerifyClient     *string `tfsdk:"verify_client" json:"verifyClient,omitempty"`
			VerifyDepth      *int64  `tfsdk:"verify_depth" json:"verifyDepth,omitempty"`
		} `tfsdk:"ingress_mtls" json:"ingressMTLS,omitempty"`
		Jwt *struct {
			JwksURI    *string `tfsdk:"jwks_uri" json:"jwksURI,omitempty"`
			KeyCache   *string `tfsdk:"key_cache" json:"keyCache,omitempty"`
			Realm      *string `tfsdk:"realm" json:"realm,omitempty"`
			Secret     *string `tfsdk:"secret" json:"secret,omitempty"`
			SniEnabled *bool   `tfsdk:"sni_enabled" json:"sniEnabled,omitempty"`
			SniName    *string `tfsdk:"sni_name" json:"sniName,omitempty"`
			Token      *string `tfsdk:"token" json:"token,omitempty"`
		} `tfsdk:"jwt" json:"jwt,omitempty"`
		Oidc *struct {
			AccessTokenEnable     *bool     `tfsdk:"access_token_enable" json:"accessTokenEnable,omitempty"`
			AuthEndpoint          *string   `tfsdk:"auth_endpoint" json:"authEndpoint,omitempty"`
			AuthExtraArgs         *[]string `tfsdk:"auth_extra_args" json:"authExtraArgs,omitempty"`
			ClientID              *string   `tfsdk:"client_id" json:"clientID,omitempty"`
			ClientSecret          *string   `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			EndSessionEndpoint    *string   `tfsdk:"end_session_endpoint" json:"endSessionEndpoint,omitempty"`
			JwksURI               *string   `tfsdk:"jwks_uri" json:"jwksURI,omitempty"`
			PkceEnable            *bool     `tfsdk:"pkce_enable" json:"pkceEnable,omitempty"`
			PostLogoutRedirectURI *string   `tfsdk:"post_logout_redirect_uri" json:"postLogoutRedirectURI,omitempty"`
			RedirectURI           *string   `tfsdk:"redirect_uri" json:"redirectURI,omitempty"`
			Scope                 *string   `tfsdk:"scope" json:"scope,omitempty"`
			TokenEndpoint         *string   `tfsdk:"token_endpoint" json:"tokenEndpoint,omitempty"`
			ZoneSyncLeeway        *int64    `tfsdk:"zone_sync_leeway" json:"zoneSyncLeeway,omitempty"`
		} `tfsdk:"oidc" json:"oidc,omitempty"`
		RateLimit *struct {
			Burst     *int64 `tfsdk:"burst" json:"burst,omitempty"`
			Condition *struct {
				Default *bool `tfsdk:"default" json:"default,omitempty"`
				Jwt     *struct {
					Claim *string `tfsdk:"claim" json:"claim,omitempty"`
					Match *string `tfsdk:"match" json:"match,omitempty"`
				} `tfsdk:"jwt" json:"jwt,omitempty"`
				Variables *[]struct {
					Match *string `tfsdk:"match" json:"match,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"variables" json:"variables,omitempty"`
			} `tfsdk:"condition" json:"condition,omitempty"`
			Delay      *int64  `tfsdk:"delay" json:"delay,omitempty"`
			DryRun     *bool   `tfsdk:"dry_run" json:"dryRun,omitempty"`
			Key        *string `tfsdk:"key" json:"key,omitempty"`
			LogLevel   *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			NoDelay    *bool   `tfsdk:"no_delay" json:"noDelay,omitempty"`
			Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
			RejectCode *int64  `tfsdk:"reject_code" json:"rejectCode,omitempty"`
			Scale      *bool   `tfsdk:"scale" json:"scale,omitempty"`
			ZoneSize   *string `tfsdk:"zone_size" json:"zoneSize,omitempty"`
		} `tfsdk:"rate_limit" json:"rateLimit,omitempty"`
		Waf *struct {
			ApBundle    *string `tfsdk:"ap_bundle" json:"apBundle,omitempty"`
			ApPolicy    *string `tfsdk:"ap_policy" json:"apPolicy,omitempty"`
			Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
			SecurityLog *struct {
				ApLogBundle *string `tfsdk:"ap_log_bundle" json:"apLogBundle,omitempty"`
				ApLogConf   *string `tfsdk:"ap_log_conf" json:"apLogConf,omitempty"`
				Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
				LogDest     *string `tfsdk:"log_dest" json:"logDest,omitempty"`
			} `tfsdk:"security_log" json:"securityLog,omitempty"`
			SecurityLogs *[]struct {
				ApLogBundle *string `tfsdk:"ap_log_bundle" json:"apLogBundle,omitempty"`
				ApLogConf   *string `tfsdk:"ap_log_conf" json:"apLogConf,omitempty"`
				Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
				LogDest     *string `tfsdk:"log_dest" json:"logDest,omitempty"`
			} `tfsdk:"security_logs" json:"securityLogs,omitempty"`
		} `tfsdk:"waf" json:"waf,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SNginxOrgPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_nginx_org_policy_v1_manifest"
}

func (r *K8SNginxOrgPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
		MarkdownDescription: "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
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
				Description:         "PolicySpec is the spec of the Policy resource. The spec includes multiple fields, where each field represents a different policy. Only one policy (field) is allowed.",
				MarkdownDescription: "PolicySpec is the spec of the Policy resource. The spec includes multiple fields, where each field represents a different policy. Only one policy (field) is allowed.",
				Attributes: map[string]schema.Attribute{
					"access_control": schema.SingleNestedAttribute{
						Description:         "The access control policy based on the client IP address.",
						MarkdownDescription: "The access control policy based on the client IP address.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deny": schema.ListAttribute{
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

					"api_key": schema.SingleNestedAttribute{
						Description:         "The API Key policy configures NGINX to authorize requests which provide a valid API Key in a specified header or query param.",
						MarkdownDescription: "The API Key policy configures NGINX to authorize requests which provide a valid API Key in a specified header or query param.",
						Attributes: map[string]schema.Attribute{
							"client_secret": schema.StringAttribute{
								Description:         "The key to which the API key is applied. Can contain text, variables, or a combination of them. Accepted variables are $http_, $arg_, $cookie_.",
								MarkdownDescription: "The key to which the API key is applied. Can contain text, variables, or a combination of them. Accepted variables are $http_, $arg_, $cookie_.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supplied_in": schema.SingleNestedAttribute{
								Description:         "The location of the API Key. For example, $http_auth, $arg_apikey, $cookie_auth. Accepted variables are $http_, $arg_, $cookie_.",
								MarkdownDescription: "The location of the API Key. For example, $http_auth, $arg_apikey, $cookie_auth. Accepted variables are $http_, $arg_, $cookie_.",
								Attributes: map[string]schema.Attribute{
									"header": schema.ListAttribute{
										Description:         "The location of the API Key as a request header. For example, $http_auth. Accepted variables are $http_.",
										MarkdownDescription: "The location of the API Key as a request header. For example, $http_auth. Accepted variables are $http_.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query": schema.ListAttribute{
										Description:         "The location of the API Key as a query param. For example, $arg_apikey. Accepted variables are $arg_.",
										MarkdownDescription: "The location of the API Key as a query param. For example, $arg_apikey. Accepted variables are $arg_.",
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

					"basic_auth": schema.SingleNestedAttribute{
						Description:         "The basic auth policy configures NGINX to authenticate client requests using HTTP Basic authentication credentials.",
						MarkdownDescription: "The basic auth policy configures NGINX to authenticate client requests using HTTP Basic authentication credentials.",
						Attributes: map[string]schema.Attribute{
							"realm": schema.StringAttribute{
								Description:         "The realm for the basic authentication.",
								MarkdownDescription: "The realm for the basic authentication.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.StringAttribute{
								Description:         "The name of the Kubernetes secret that stores the Htpasswd configuration. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/htpasswd, and the config must be stored in the secret under the key htpasswd, otherwise the secret will be rejected as invalid.",
								MarkdownDescription: "The name of the Kubernetes secret that stores the Htpasswd configuration. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/htpasswd, and the config must be stored in the secret under the key htpasswd, otherwise the secret will be rejected as invalid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache": schema.SingleNestedAttribute{
						Description:         "The Cache Key defines a cache policy for proxy caching",
						MarkdownDescription: "The Cache Key defines a cache policy for proxy caching",
						Attributes: map[string]schema.Attribute{
							"allowed_codes": schema.ListAttribute{
								Description:         "AllowedCodes defines which HTTP response codes should be cached. Accepts either: - The string 'any' to cache all response codes (must be the only element) - A list of HTTP status codes as integers (100-599) Examples: ['any'], [200, 301, 404], [200]. Invalid: ['any', 200] (cannot mix 'any' with specific codes).",
								MarkdownDescription: "AllowedCodes defines which HTTP response codes should be cached. Accepts either: - The string 'any' to cache all response codes (must be the only element) - A list of HTTP status codes as integers (100-599) Examples: ['any'], [200, 301, 404], [200]. Invalid: ['any', 200] (cannot mix 'any' with specific codes).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allowed_methods": schema.ListAttribute{
								Description:         "AllowedMethods defines which HTTP methods should be cached. Only 'GET', 'HEAD', and 'POST' are supported by NGINX proxy_cache_methods directive. GET and HEAD are always cached by default even if not specified. Maximum of 3 items allowed. Examples: ['GET'], ['GET', 'HEAD', 'POST']. Invalid methods: PUT, DELETE, PATCH, etc.",
								MarkdownDescription: "AllowedMethods defines which HTTP methods should be cached. Only 'GET', 'HEAD', and 'POST' are supported by NGINX proxy_cache_methods directive. GET and HEAD are always cached by default even if not specified. Maximum of 3 items allowed. Examples: ['GET'], ['GET', 'HEAD', 'POST']. Invalid methods: PUT, DELETE, PATCH, etc.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_purge_allow": schema.ListAttribute{
								Description:         "CachePurgeAllow defines IP addresses or CIDR blocks allowed to purge cache. This feature is only available in NGINX Plus. Examples: ['192.168.1.100', '10.0.0.0/8', '::1']. Invalid in NGINX OSS (will be ignored).",
								MarkdownDescription: "CachePurgeAllow defines IP addresses or CIDR blocks allowed to purge cache. This feature is only available in NGINX Plus. Examples: ['192.168.1.100', '10.0.0.0/8', '::1']. Invalid in NGINX OSS (will be ignored).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_zone_name": schema.StringAttribute{
								Description:         "CacheZoneName defines the name of the cache zone. Must start with a lowercase letter, followed by alphanumeric characters or underscores, and end with an alphanumeric character. Single lowercase letters are also allowed. Examples: 'cache', 'my_cache', 'cache1'.",
								MarkdownDescription: "CacheZoneName defines the name of the cache zone. Must start with a lowercase letter, followed by alphanumeric characters or underscores, and end with an alphanumeric character. Single lowercase letters are also allowed. Examples: 'cache', 'my_cache', 'cache1'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z][a-zA-Z0-9_]*[a-zA-Z0-9]$|^[a-z]$`), ""),
								},
							},

							"cache_zone_size": schema.StringAttribute{
								Description:         "CacheZoneSize defines the size of the cache zone. Must be a number followed by a size unit: 'k' for kilobytes, 'm' for megabytes, or 'g' for gigabytes. Examples: '10m', '1g', '512k'.",
								MarkdownDescription: "CacheZoneSize defines the size of the cache zone. Must be a number followed by a size unit: 'k' for kilobytes, 'm' for megabytes, or 'g' for gigabytes. Examples: '10m', '1g', '512k'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[kmg]$`), ""),
								},
							},

							"levels": schema.StringAttribute{
								Description:         "Levels defines the cache directory hierarchy levels for storing cached files. Must be in format 'X:Y' or 'X:Y:Z' where X, Y, Z are either 1 or 2. This controls the number of subdirectory levels and their name lengths. Examples: '1:2', '2:2', '1:2:2'. Invalid: '3:1', '1:3', '1:2:3'.",
								MarkdownDescription: "Levels defines the cache directory hierarchy levels for storing cached files. Must be in format 'X:Y' or 'X:Y:Z' where X, Y, Z are either 1 or 2. This controls the number of subdirectory levels and their name lengths. Examples: '1:2', '2:2', '1:2:2'. Invalid: '3:1', '1:3', '1:2:3'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[12](?::[12]){0,2}$`), ""),
								},
							},

							"override_upstream_cache": schema.BoolAttribute{
								Description:         "OverrideUpstreamCache controls whether to override upstream cache headers (using proxy_ignore_headers directive). When true, NGINX will ignore cache-related headers from upstream servers like Cache-Control, Expires, etc. Default: false.",
								MarkdownDescription: "OverrideUpstreamCache controls whether to override upstream cache headers (using proxy_ignore_headers directive). When true, NGINX will ignore cache-related headers from upstream servers like Cache-Control, Expires, etc. Default: false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time": schema.StringAttribute{
								Description:         "Time defines the default cache time. Required when allowedCodes is specified. Must be a number followed by a time unit: 's' for seconds, 'm' for minutes, 'h' for hours, 'd' for days. Examples: '30s', '5m', '1h', '2d'.",
								MarkdownDescription: "Time defines the default cache time. Required when allowedCodes is specified. Must be a number followed by a time unit: 's' for seconds, 'm' for minutes, 'h' for hours, 'd' for days. Examples: '30s', '5m', '1h', '2d'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[smhd]$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress_mtls": schema.SingleNestedAttribute{
						Description:         "The EgressMTLS policy configures upstreams authentication and certificate verification.",
						MarkdownDescription: "The EgressMTLS policy configures upstreams authentication and certificate verification.",
						Attributes: map[string]schema.Attribute{
							"ciphers": schema.StringAttribute{
								Description:         "Specifies the enabled ciphers for requests to an upstream HTTPS server. The default is DEFAULT.",
								MarkdownDescription: "Specifies the enabled ciphers for requests to an upstream HTTPS server. The default is DEFAULT.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocols": schema.StringAttribute{
								Description:         "Specifies the protocols for requests to an upstream HTTPS server. The default is TLSv1 TLSv1.1 TLSv1.2.",
								MarkdownDescription: "Specifies the protocols for requests to an upstream HTTPS server. The default is TLSv1 TLSv1.1 TLSv1.2.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_name": schema.BoolAttribute{
								Description:         "Enables passing of the server name through Server Name Indication extension.",
								MarkdownDescription: "Enables passing of the server name through Server Name Indication extension.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"session_reuse": schema.BoolAttribute{
								Description:         "Enables reuse of SSL sessions to the upstreams. The default is true.",
								MarkdownDescription: "Enables reuse of SSL sessions to the upstreams. The default is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_name": schema.StringAttribute{
								Description:         "Allows overriding the server name used to verify the certificate of the upstream HTTPS server.",
								MarkdownDescription: "Allows overriding the server name used to verify the certificate of the upstream HTTPS server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_secret": schema.StringAttribute{
								Description:         "The name of the Kubernetes secret that stores the TLS certificate and key. It must be in the same namespace as the Policy resource. The secret must be of the type kubernetes.io/tls, the certificate must be stored in the secret under the key tls.crt, and the key must be stored under the key tls.key, otherwise the secret will be rejected as invalid.",
								MarkdownDescription: "The name of the Kubernetes secret that stores the TLS certificate and key. It must be in the same namespace as the Policy resource. The secret must be of the type kubernetes.io/tls, the certificate must be stored in the secret under the key tls.crt, and the key must be stored under the key tls.key, otherwise the secret will be rejected as invalid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trusted_cert_secret": schema.StringAttribute{
								Description:         "The name of the Kubernetes secret that stores the CA certificate. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt, otherwise the secret will be rejected as invalid.",
								MarkdownDescription: "The name of the Kubernetes secret that stores the CA certificate. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt, otherwise the secret will be rejected as invalid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_depth": schema.Int64Attribute{
								Description:         "Sets the verification depth in the proxied HTTPS server certificates chain. The default is 1.",
								MarkdownDescription: "Sets the verification depth in the proxied HTTPS server certificates chain. The default is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_server": schema.BoolAttribute{
								Description:         "Enables verification of the upstream HTTPS server certificate.",
								MarkdownDescription: "Enables verification of the upstream HTTPS server certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "Specifies which instance of NGINX Ingress Controller must handle the Policy resource.",
						MarkdownDescription: "Specifies which instance of NGINX Ingress Controller must handle the Policy resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_mtls": schema.SingleNestedAttribute{
						Description:         "The IngressMTLS policy configures client certificate verification.",
						MarkdownDescription: "The IngressMTLS policy configures client certificate verification.",
						Attributes: map[string]schema.Attribute{
							"client_cert_secret": schema.StringAttribute{
								Description:         "The name of the Kubernetes secret that stores the CA certificate. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt, otherwise the secret will be rejected as invalid.",
								MarkdownDescription: "The name of the Kubernetes secret that stores the CA certificate. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt, otherwise the secret will be rejected as invalid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"crl_file_name": schema.StringAttribute{
								Description:         "The file name of the Certificate Revocation List. NGINX Ingress Controller will look for this file in /etc/nginx/secrets",
								MarkdownDescription: "The file name of the Certificate Revocation List. NGINX Ingress Controller will look for this file in /etc/nginx/secrets",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_client": schema.StringAttribute{
								Description:         "Verification for the client. Possible values are 'on', 'off', 'optional', 'optional_no_ca'. The default is 'on'.",
								MarkdownDescription: "Verification for the client. Possible values are 'on', 'off', 'optional', 'optional_no_ca'. The default is 'on'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_depth": schema.Int64Attribute{
								Description:         "Sets the verification depth in the client certificates chain. The default is 1.",
								MarkdownDescription: "Sets the verification depth in the client certificates chain. The default is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jwt": schema.SingleNestedAttribute{
						Description:         "The JWT policy configures NGINX Plus to authenticate client requests using JSON Web Tokens.",
						MarkdownDescription: "The JWT policy configures NGINX Plus to authenticate client requests using JSON Web Tokens.",
						Attributes: map[string]schema.Attribute{
							"jwks_uri": schema.StringAttribute{
								Description:         "The remote URI where the request will be sent to retrieve JSON Web Key set",
								MarkdownDescription: "The remote URI where the request will be sent to retrieve JSON Web Key set",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_cache": schema.StringAttribute{
								Description:         "Enables in-memory caching of JWKS (JSON Web Key Sets) that are obtained from the jwksURI and sets a valid time for expiration.",
								MarkdownDescription: "Enables in-memory caching of JWKS (JSON Web Key Sets) that are obtained from the jwksURI and sets a valid time for expiration.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm": schema.StringAttribute{
								Description:         "The realm of the JWT.",
								MarkdownDescription: "The realm of the JWT.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.StringAttribute{
								Description:         "The name of the Kubernetes secret that stores the Htpasswd configuration. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/htpasswd, and the config must be stored in the secret under the key htpasswd, otherwise the secret will be rejected as invalid.",
								MarkdownDescription: "The name of the Kubernetes secret that stores the Htpasswd configuration. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/htpasswd, and the config must be stored in the secret under the key htpasswd, otherwise the secret will be rejected as invalid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sni_enabled": schema.BoolAttribute{
								Description:         "Enables SNI (Server Name Indication) for the JWT policy. This is useful when the remote server requires SNI to serve the correct certificate.",
								MarkdownDescription: "Enables SNI (Server Name Indication) for the JWT policy. This is useful when the remote server requires SNI to serve the correct certificate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sni_name": schema.StringAttribute{
								Description:         "The SNI name to use when connecting to the remote server. If not set, the hostname from the ''jwksURI'' will be used.",
								MarkdownDescription: "The SNI name to use when connecting to the remote server. If not set, the hostname from the ''jwksURI'' will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token": schema.StringAttribute{
								Description:         "The token specifies a variable that contains the JSON Web Token. By default the JWT is passed in the Authorization header as a Bearer Token. JWT may be also passed as a cookie or a part of a query string, for example: $cookie_auth_token. Accepted variables are $http_, $arg_, $cookie_.",
								MarkdownDescription: "The token specifies a variable that contains the JSON Web Token. By default the JWT is passed in the Authorization header as a Bearer Token. JWT may be also passed as a cookie or a part of a query string, for example: $cookie_auth_token. Accepted variables are $http_, $arg_, $cookie_.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"oidc": schema.SingleNestedAttribute{
						Description:         "The OpenID Connect policy configures NGINX to authenticate client requests by validating a JWT token against an OAuth2/OIDC token provider, such as Auth0 or Keycloak.",
						MarkdownDescription: "The OpenID Connect policy configures NGINX to authenticate client requests by validating a JWT token against an OAuth2/OIDC token provider, such as Auth0 or Keycloak.",
						Attributes: map[string]schema.Attribute{
							"access_token_enable": schema.BoolAttribute{
								Description:         "Option of whether Bearer token is used to authorize NGINX to access protected backend.",
								MarkdownDescription: "Option of whether Bearer token is used to authorize NGINX to access protected backend.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auth_endpoint": schema.StringAttribute{
								Description:         "URL for the authorization endpoint provided by your OpenID Connect provider.",
								MarkdownDescription: "URL for the authorization endpoint provided by your OpenID Connect provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auth_extra_args": schema.ListAttribute{
								Description:         "A list of extra URL arguments to pass to the authorization endpoint provided by your OpenID Connect provider. Arguments must be URL encoded, multiple arguments may be included in the list, for example [ arg1=value1, arg2=value2 ]",
								MarkdownDescription: "A list of extra URL arguments to pass to the authorization endpoint provided by your OpenID Connect provider. Arguments must be URL encoded, multiple arguments may be included in the list, for example [ arg1=value1, arg2=value2 ]",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_id": schema.StringAttribute{
								Description:         "The client ID provided by your OpenID Connect provider.",
								MarkdownDescription: "The client ID provided by your OpenID Connect provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_secret": schema.StringAttribute{
								Description:         "The name of the Kubernetes secret that stores the client secret provided by your OpenID Connect provider. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/oidc, and the secret under the key client-secret, otherwise the secret will be rejected as invalid. If PKCE is enabled, this should be not configured.",
								MarkdownDescription: "The name of the Kubernetes secret that stores the client secret provided by your OpenID Connect provider. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/oidc, and the secret under the key client-secret, otherwise the secret will be rejected as invalid. If PKCE is enabled, this should be not configured.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"end_session_endpoint": schema.StringAttribute{
								Description:         "URL provided by your OpenID Connect provider to request the end user be logged out.",
								MarkdownDescription: "URL provided by your OpenID Connect provider to request the end user be logged out.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jwks_uri": schema.StringAttribute{
								Description:         "URL for the JSON Web Key Set (JWK) document provided by your OpenID Connect provider.",
								MarkdownDescription: "URL for the JSON Web Key Set (JWK) document provided by your OpenID Connect provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pkce_enable": schema.BoolAttribute{
								Description:         "Switches Proof Key for Code Exchange on. The OpenID client needs to be in public mode. clientSecret is not used in this mode.",
								MarkdownDescription: "Switches Proof Key for Code Exchange on. The OpenID client needs to be in public mode. clientSecret is not used in this mode.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"post_logout_redirect_uri": schema.StringAttribute{
								Description:         "URI to redirect to after the logout has been performed. Requires endSessionEndpoint. The default is /_logout.",
								MarkdownDescription: "URI to redirect to after the logout has been performed. Requires endSessionEndpoint. The default is /_logout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redirect_uri": schema.StringAttribute{
								Description:         "Allows overriding the default redirect URI. The default is /_codexch.",
								MarkdownDescription: "Allows overriding the default redirect URI. The default is /_codexch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scope": schema.StringAttribute{
								Description:         "List of OpenID Connect scopes. The scope openid always needs to be present and others can be added concatenating them with a + sign, for example openid+profile+email, openid+email+userDefinedScope. The default is openid.",
								MarkdownDescription: "List of OpenID Connect scopes. The scope openid always needs to be present and others can be added concatenating them with a + sign, for example openid+profile+email, openid+email+userDefinedScope. The default is openid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_endpoint": schema.StringAttribute{
								Description:         "URL for the token endpoint provided by your OpenID Connect provider.",
								MarkdownDescription: "URL for the token endpoint provided by your OpenID Connect provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zone_sync_leeway": schema.Int64Attribute{
								Description:         "Specifies the maximum timeout in milliseconds for synchronizing ID/access tokens and shared values between Ingress Controller pods. The default is 200.",
								MarkdownDescription: "Specifies the maximum timeout in milliseconds for synchronizing ID/access tokens and shared values between Ingress Controller pods. The default is 200.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rate_limit": schema.SingleNestedAttribute{
						Description:         "The rate limit policy controls the rate of processing requests per a defined key.",
						MarkdownDescription: "The rate limit policy controls the rate of processing requests per a defined key.",
						Attributes: map[string]schema.Attribute{
							"burst": schema.Int64Attribute{
								Description:         "Excessive requests are delayed until their number exceeds the burst size, in which case the request is terminated with an error.",
								MarkdownDescription: "Excessive requests are delayed until their number exceeds the burst size, in which case the request is terminated with an error.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"condition": schema.SingleNestedAttribute{
								Description:         "Add a condition to a rate-limit policy.",
								MarkdownDescription: "Add a condition to a rate-limit policy.",
								Attributes: map[string]schema.Attribute{
									"default": schema.BoolAttribute{
										Description:         "sets the rate limit in this policy to be the default if no conditions are met. In a group of policies with the same condition, only one policy can be the default.",
										MarkdownDescription: "sets the rate limit in this policy to be the default if no conditions are met. In a group of policies with the same condition, only one policy can be the default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jwt": schema.SingleNestedAttribute{
										Description:         "defines a JWT condition to rate limit against.",
										MarkdownDescription: "defines a JWT condition to rate limit against.",
										Attributes: map[string]schema.Attribute{
											"claim": schema.StringAttribute{
												Description:         "the JWT claim to be rate limit by. Nested claims should be separated by '.'",
												MarkdownDescription: "the JWT claim to be rate limit by. Nested claims should be separated by '.'",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^([^$\s"'])*$`), ""),
												},
											},

											"match": schema.StringAttribute{
												Description:         "the value of the claim to match against.",
												MarkdownDescription: "the value of the claim to match against.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^([^$\s."'])*$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"variables": schema.ListNestedAttribute{
										Description:         "defines a Variables condition to rate limit against.",
										MarkdownDescription: "defines a Variables condition to rate limit against.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match": schema.StringAttribute{
													Description:         "the value of the variable to match against.",
													MarkdownDescription: "the value of the variable to match against.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([^\s"'])*$`), ""),
													},
												},

												"name": schema.StringAttribute{
													Description:         "the name of the variable to match against.",
													MarkdownDescription: "the name of the variable to match against.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([^\s"'])*$`), ""),
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

							"delay": schema.Int64Attribute{
								Description:         "The delay parameter specifies a limit at which excessive requests become delayed. If not set all excessive requests are delayed.",
								MarkdownDescription: "The delay parameter specifies a limit at which excessive requests become delayed. If not set all excessive requests are delayed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dry_run": schema.BoolAttribute{
								Description:         "Enables the dry run mode. In this mode, the rate limit is not actually applied, but the number of excessive requests is accounted as usual in the shared memory zone.",
								MarkdownDescription: "Enables the dry run mode. In this mode, the rate limit is not actually applied, but the number of excessive requests is accounted as usual in the shared memory zone.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "The key to which the rate limit is applied. Can contain text, variables, or a combination of them. Variables must be surrounded by ${}. For example: ${binary_remote_addr}. Accepted variables are $binary_remote_addr, $request_uri, $request_method, $url, $http_, $args, $arg_, $cookie_,$jwt_claim_ .",
								MarkdownDescription: "The key to which the rate limit is applied. Can contain text, variables, or a combination of them. Variables must be surrounded by ${}. For example: ${binary_remote_addr}. Accepted variables are $binary_remote_addr, $request_uri, $request_method, $url, $http_, $args, $arg_, $cookie_,$jwt_claim_ .",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Sets the desired logging level for cases when the server refuses to process requests due to rate exceeding, or delays request processing. Allowed values are info, notice, warn or error. Default is error.",
								MarkdownDescription: "Sets the desired logging level for cases when the server refuses to process requests due to rate exceeding, or delays request processing. Allowed values are info, notice, warn or error. Default is error.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"no_delay": schema.BoolAttribute{
								Description:         "Disables the delaying of excessive requests while requests are being limited. Overrides delay if both are set.",
								MarkdownDescription: "Disables the delaying of excessive requests while requests are being limited. Overrides delay if both are set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate": schema.StringAttribute{
								Description:         "The rate of requests permitted. The rate is specified in requests per second (r/s) or requests per minute (r/m).",
								MarkdownDescription: "The rate of requests permitted. The rate is specified in requests per second (r/s) or requests per minute (r/m).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reject_code": schema.Int64Attribute{
								Description:         "Sets the status code to return in response to rejected requests. Must fall into the range 400..599. Default is 503.",
								MarkdownDescription: "Sets the status code to return in response to rejected requests. Must fall into the range 400..599. Default is 503.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scale": schema.BoolAttribute{
								Description:         "Enables a constant rate-limit by dividing the configured rate by the number of nginx-ingress pods currently serving traffic. This adjustment ensures that the rate-limit remains consistent, even as the number of nginx-pods fluctuates due to autoscaling. This will not work properly if requests from a client are not evenly distributed across all ingress pods (Such as with sticky sessions, long lived TCP Connections with many requests, and so forth). In such cases using zone-sync instead would give better results. Enabling zone-sync will suppress this setting.",
								MarkdownDescription: "Enables a constant rate-limit by dividing the configured rate by the number of nginx-ingress pods currently serving traffic. This adjustment ensures that the rate-limit remains consistent, even as the number of nginx-pods fluctuates due to autoscaling. This will not work properly if requests from a client are not evenly distributed across all ingress pods (Such as with sticky sessions, long lived TCP Connections with many requests, and so forth). In such cases using zone-sync instead would give better results. Enabling zone-sync will suppress this setting.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zone_size": schema.StringAttribute{
								Description:         "Size of the shared memory zone. Only positive values are allowed. Allowed suffixes are k or m, if none are present k is assumed.",
								MarkdownDescription: "Size of the shared memory zone. Only positive values are allowed. Allowed suffixes are k or m, if none are present k is assumed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"waf": schema.SingleNestedAttribute{
						Description:         "The WAF policy configures WAF and log configuration policies for NGINX AppProtect",
						MarkdownDescription: "The WAF policy configures WAF and log configuration policies for NGINX AppProtect",
						Attributes: map[string]schema.Attribute{
							"ap_bundle": schema.StringAttribute{
								Description:         "The App Protect WAF policy bundle. Mutually exclusive with apPolicy.",
								MarkdownDescription: "The App Protect WAF policy bundle. Mutually exclusive with apPolicy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ap_policy": schema.StringAttribute{
								Description:         "The App Protect WAF policy of the WAF. Accepts an optional namespace. Mutually exclusive with apBundle.",
								MarkdownDescription: "The App Protect WAF policy of the WAF. Accepts an optional namespace. Mutually exclusive with apBundle.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Enables NGINX App Protect WAF.",
								MarkdownDescription: "Enables NGINX App Protect WAF.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_log": schema.SingleNestedAttribute{
								Description:         "SecurityLog defines the security log of a WAF policy.",
								MarkdownDescription: "SecurityLog defines the security log of a WAF policy.",
								Attributes: map[string]schema.Attribute{
									"ap_log_bundle": schema.StringAttribute{
										Description:         "The App Protect WAF log bundle resource. Only works with apBundle.",
										MarkdownDescription: "The App Protect WAF log bundle resource. Only works with apBundle.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ap_log_conf": schema.StringAttribute{
										Description:         "The App Protect WAF log conf resource. Accepts an optional namespace. Only works with apPolicy.",
										MarkdownDescription: "The App Protect WAF log conf resource. Accepts an optional namespace. Only works with apPolicy.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable": schema.BoolAttribute{
										Description:         "Enables security log.",
										MarkdownDescription: "Enables security log.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_dest": schema.StringAttribute{
										Description:         "The log destination for the security log. Only accepted variables are syslog:server=<ip-address>; localhost; fqdn>:<port>, stderr, <absolute path to file>.",
										MarkdownDescription: "The log destination for the security log. Only accepted variables are syslog:server=<ip-address>; localhost; fqdn>:<port>, stderr, <absolute path to file>.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_logs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ap_log_bundle": schema.StringAttribute{
											Description:         "The App Protect WAF log bundle resource. Only works with apBundle.",
											MarkdownDescription: "The App Protect WAF log bundle resource. Only works with apBundle.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ap_log_conf": schema.StringAttribute{
											Description:         "The App Protect WAF log conf resource. Accepts an optional namespace. Only works with apPolicy.",
											MarkdownDescription: "The App Protect WAF log conf resource. Accepts an optional namespace. Only works with apPolicy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable": schema.BoolAttribute{
											Description:         "Enables security log.",
											MarkdownDescription: "Enables security log.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_dest": schema.StringAttribute{
											Description:         "The log destination for the security log. Only accepted variables are syslog:server=<ip-address>; localhost; fqdn>:<port>, stderr, <absolute path to file>.",
											MarkdownDescription: "The log destination for the security log. Only accepted variables are syslog:server=<ip-address>; localhost; fqdn>:<port>, stderr, <absolute path to file>.",
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
		},
	}
}

func (r *K8SNginxOrgPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_nginx_org_policy_v1_manifest")

	var model K8SNginxOrgPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.nginx.org/v1")
	model.Kind = pointer.String("Policy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
