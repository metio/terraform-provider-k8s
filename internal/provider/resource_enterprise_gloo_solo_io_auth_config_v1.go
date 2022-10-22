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

type EnterpriseGlooSoloIoAuthConfigV1Resource struct{}

var (
	_ resource.Resource = (*EnterpriseGlooSoloIoAuthConfigV1Resource)(nil)
)

type EnterpriseGlooSoloIoAuthConfigV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type EnterpriseGlooSoloIoAuthConfigV1GoModel struct {
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
		BooleanExpr *string `tfsdk:"boolean_expr" yaml:"booleanExpr,omitempty"`

		Configs *[]struct {
			ApiKeyAuth *struct {
				AerospikeApikeyStorage *struct {
					AllowInsecure *bool `tfsdk:"allow_insecure" yaml:"allowInsecure,omitempty"`

					BatchSize *int64 `tfsdk:"batch_size" yaml:"batchSize,omitempty"`

					CertPath *string `tfsdk:"cert_path" yaml:"certPath,omitempty"`

					CommitAll *int64 `tfsdk:"commit_all" yaml:"commitAll,omitempty"`

					CommitMaster *int64 `tfsdk:"commit_master" yaml:"commitMaster,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					KeyPath *string `tfsdk:"key_path" yaml:"keyPath,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					NodeTlsName *string `tfsdk:"node_tls_name" yaml:"nodeTlsName,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					ReadModeAp *struct {
						ReadModeApAll *int64 `tfsdk:"read_mode_ap_all" yaml:"readModeApAll,omitempty"`

						ReadModeApOne *int64 `tfsdk:"read_mode_ap_one" yaml:"readModeApOne,omitempty"`
					} `tfsdk:"read_mode_ap" yaml:"readModeAp,omitempty"`

					ReadModeSc *struct {
						ReadModeScAllowUnavailable *int64 `tfsdk:"read_mode_sc_allow_unavailable" yaml:"readModeScAllowUnavailable,omitempty"`

						ReadModeScLinearize *int64 `tfsdk:"read_mode_sc_linearize" yaml:"readModeScLinearize,omitempty"`

						ReadModeScReplica *int64 `tfsdk:"read_mode_sc_replica" yaml:"readModeScReplica,omitempty"`

						ReadModeScSession *int64 `tfsdk:"read_mode_sc_session" yaml:"readModeScSession,omitempty"`
					} `tfsdk:"read_mode_sc" yaml:"readModeSc,omitempty"`

					RootCaPath *string `tfsdk:"root_ca_path" yaml:"rootCaPath,omitempty"`

					Set *string `tfsdk:"set" yaml:"set,omitempty"`

					TlsCurveGroups *[]struct {
						CurveP256 *int64 `tfsdk:"curve_p256" yaml:"curveP256,omitempty"`

						CurveP384 *int64 `tfsdk:"curve_p384" yaml:"curveP384,omitempty"`

						CurveP521 *int64 `tfsdk:"curve_p521" yaml:"curveP521,omitempty"`

						X25519 *int64 `tfsdk:"x25519" yaml:"x25519,omitempty"`
					} `tfsdk:"tls_curve_groups" yaml:"tlsCurveGroups,omitempty"`

					TlsVersion *string `tfsdk:"tls_version" yaml:"tlsVersion,omitempty"`
				} `tfsdk:"aerospike_apikey_storage" yaml:"aerospikeApikeyStorage,omitempty"`

				ApiKeySecretRefs *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"api_key_secret_refs" yaml:"apiKeySecretRefs,omitempty"`

				HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`

				HeadersFromMetadata *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Required *bool `tfsdk:"required" yaml:"required,omitempty"`
				} `tfsdk:"headers_from_metadata" yaml:"headersFromMetadata,omitempty"`

				HeadersFromMetadataEntry *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Required *bool `tfsdk:"required" yaml:"required,omitempty"`
				} `tfsdk:"headers_from_metadata_entry" yaml:"headersFromMetadataEntry,omitempty"`

				K8sSecretApikeyStorage *struct {
					ApiKeySecretRefs *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"api_key_secret_refs" yaml:"apiKeySecretRefs,omitempty"`

					LabelSelector *map[string]string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`
				} `tfsdk:"k8s_secret_apikey_storage" yaml:"k8sSecretApikeyStorage,omitempty"`

				LabelSelector *map[string]string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`
			} `tfsdk:"api_key_auth" yaml:"apiKeyAuth,omitempty"`

			BasicAuth *struct {
				Apr *struct {
					Users *struct {
						HashedPassword *string `tfsdk:"hashed_password" yaml:"hashedPassword,omitempty"`

						Salt *string `tfsdk:"salt" yaml:"salt,omitempty"`
					} `tfsdk:"users" yaml:"users,omitempty"`
				} `tfsdk:"apr" yaml:"apr,omitempty"`

				Realm *string `tfsdk:"realm" yaml:"realm,omitempty"`
			} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

			Jwt *map[string]string `tfsdk:"jwt" yaml:"jwt,omitempty"`

			Ldap *struct {
				Address *string `tfsdk:"address" yaml:"address,omitempty"`

				AllowedGroups *[]string `tfsdk:"allowed_groups" yaml:"allowedGroups,omitempty"`

				DisableGroupChecking *bool `tfsdk:"disable_group_checking" yaml:"disableGroupChecking,omitempty"`

				MembershipAttributeName *string `tfsdk:"membership_attribute_name" yaml:"membershipAttributeName,omitempty"`

				Pool *struct {
					InitialSize *int64 `tfsdk:"initial_size" yaml:"initialSize,omitempty"`

					MaxSize *int64 `tfsdk:"max_size" yaml:"maxSize,omitempty"`
				} `tfsdk:"pool" yaml:"pool,omitempty"`

				SearchFilter *string `tfsdk:"search_filter" yaml:"searchFilter,omitempty"`

				UserDnTemplate *string `tfsdk:"user_dn_template" yaml:"userDnTemplate,omitempty"`
			} `tfsdk:"ldap" yaml:"ldap,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Oauth *struct {
				AppUrl *string `tfsdk:"app_url" yaml:"appUrl,omitempty"`

				AuthEndpointQueryParams *map[string]string `tfsdk:"auth_endpoint_query_params" yaml:"authEndpointQueryParams,omitempty"`

				CallbackPath *string `tfsdk:"callback_path" yaml:"callbackPath,omitempty"`

				ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

				ClientSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"client_secret_ref" yaml:"clientSecretRef,omitempty"`

				IssuerUrl *string `tfsdk:"issuer_url" yaml:"issuerUrl,omitempty"`

				Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`
			} `tfsdk:"oauth" yaml:"oauth,omitempty"`

			Oauth2 *struct {
				AccessTokenValidation *struct {
					CacheTimeout *string `tfsdk:"cache_timeout" yaml:"cacheTimeout,omitempty"`

					Introspection *struct {
						ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"client_secret_ref" yaml:"clientSecretRef,omitempty"`

						IntrospectionUrl *string `tfsdk:"introspection_url" yaml:"introspectionUrl,omitempty"`

						UserIdAttributeName *string `tfsdk:"user_id_attribute_name" yaml:"userIdAttributeName,omitempty"`
					} `tfsdk:"introspection" yaml:"introspection,omitempty"`

					IntrospectionUrl *string `tfsdk:"introspection_url" yaml:"introspectionUrl,omitempty"`

					Jwt *struct {
						Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

						LocalJwks *struct {
							InlineString *string `tfsdk:"inline_string" yaml:"inlineString,omitempty"`
						} `tfsdk:"local_jwks" yaml:"localJwks,omitempty"`

						RemoteJwks *struct {
							RefreshInterval *string `tfsdk:"refresh_interval" yaml:"refreshInterval,omitempty"`

							Url *string `tfsdk:"url" yaml:"url,omitempty"`
						} `tfsdk:"remote_jwks" yaml:"remoteJwks,omitempty"`
					} `tfsdk:"jwt" yaml:"jwt,omitempty"`

					RequiredScopes *struct {
						Scope *[]string `tfsdk:"scope" yaml:"scope,omitempty"`
					} `tfsdk:"required_scopes" yaml:"requiredScopes,omitempty"`

					UserinfoUrl *string `tfsdk:"userinfo_url" yaml:"userinfoUrl,omitempty"`
				} `tfsdk:"access_token_validation" yaml:"accessTokenValidation,omitempty"`

				Oauth2 *struct {
					AfterLogoutUrl *string `tfsdk:"after_logout_url" yaml:"afterLogoutUrl,omitempty"`

					AppUrl *string `tfsdk:"app_url" yaml:"appUrl,omitempty"`

					AuthEndpoint *string `tfsdk:"auth_endpoint" yaml:"authEndpoint,omitempty"`

					AuthEndpointQueryParams *map[string]string `tfsdk:"auth_endpoint_query_params" yaml:"authEndpointQueryParams,omitempty"`

					CallbackPath *string `tfsdk:"callback_path" yaml:"callbackPath,omitempty"`

					ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

					ClientSecretRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"client_secret_ref" yaml:"clientSecretRef,omitempty"`

					LogoutPath *string `tfsdk:"logout_path" yaml:"logoutPath,omitempty"`

					RevocationEndpoint *string `tfsdk:"revocation_endpoint" yaml:"revocationEndpoint,omitempty"`

					Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

					Session *struct {
						Cookie *struct {
							AllowRefreshing *bool `tfsdk:"allow_refreshing" yaml:"allowRefreshing,omitempty"`

							KeyPrefix *string `tfsdk:"key_prefix" yaml:"keyPrefix,omitempty"`

							TargetDomain *string `tfsdk:"target_domain" yaml:"targetDomain,omitempty"`
						} `tfsdk:"cookie" yaml:"cookie,omitempty"`

						CookieOptions *struct {
							Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

							HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`

							MaxAge *int64 `tfsdk:"max_age" yaml:"maxAge,omitempty"`

							NotSecure *bool `tfsdk:"not_secure" yaml:"notSecure,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							SameSite utilities.IntOrString `tfsdk:"same_site" yaml:"sameSite,omitempty"`
						} `tfsdk:"cookie_options" yaml:"cookieOptions,omitempty"`

						FailOnFetchFailure *bool `tfsdk:"fail_on_fetch_failure" yaml:"failOnFetchFailure,omitempty"`

						Redis *struct {
							AllowRefreshing *bool `tfsdk:"allow_refreshing" yaml:"allowRefreshing,omitempty"`

							CookieName *string `tfsdk:"cookie_name" yaml:"cookieName,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`

							KeyPrefix *string `tfsdk:"key_prefix" yaml:"keyPrefix,omitempty"`

							Options *struct {
								Db *int64 `tfsdk:"db" yaml:"db,omitempty"`

								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								PoolSize *int64 `tfsdk:"pool_size" yaml:"poolSize,omitempty"`

								SocketType utilities.IntOrString `tfsdk:"socket_type" yaml:"socketType,omitempty"`

								TlsCertMountPath *string `tfsdk:"tls_cert_mount_path" yaml:"tlsCertMountPath,omitempty"`
							} `tfsdk:"options" yaml:"options,omitempty"`

							PreExpiryBuffer *string `tfsdk:"pre_expiry_buffer" yaml:"preExpiryBuffer,omitempty"`

							TargetDomain *string `tfsdk:"target_domain" yaml:"targetDomain,omitempty"`
						} `tfsdk:"redis" yaml:"redis,omitempty"`
					} `tfsdk:"session" yaml:"session,omitempty"`

					TokenEndpoint *string `tfsdk:"token_endpoint" yaml:"tokenEndpoint,omitempty"`

					TokenEndpointQueryParams *map[string]string `tfsdk:"token_endpoint_query_params" yaml:"tokenEndpointQueryParams,omitempty"`
				} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

				OidcAuthorizationCode *struct {
					AfterLogoutUrl *string `tfsdk:"after_logout_url" yaml:"afterLogoutUrl,omitempty"`

					AppUrl *string `tfsdk:"app_url" yaml:"appUrl,omitempty"`

					AuthEndpointQueryParams *map[string]string `tfsdk:"auth_endpoint_query_params" yaml:"authEndpointQueryParams,omitempty"`

					AutoMapFromMetadata *struct {
						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"auto_map_from_metadata" yaml:"autoMapFromMetadata,omitempty"`

					CallbackPath *string `tfsdk:"callback_path" yaml:"callbackPath,omitempty"`

					ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

					ClientSecretRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"client_secret_ref" yaml:"clientSecretRef,omitempty"`

					DiscoveryOverride *struct {
						AuthEndpoint *string `tfsdk:"auth_endpoint" yaml:"authEndpoint,omitempty"`

						AuthMethods *[]string `tfsdk:"auth_methods" yaml:"authMethods,omitempty"`

						Claims *[]string `tfsdk:"claims" yaml:"claims,omitempty"`

						EndSessionEndpoint *string `tfsdk:"end_session_endpoint" yaml:"endSessionEndpoint,omitempty"`

						IdTokenAlgs *[]string `tfsdk:"id_token_algs" yaml:"idTokenAlgs,omitempty"`

						JwksUri *string `tfsdk:"jwks_uri" yaml:"jwksUri,omitempty"`

						ResponseTypes *[]string `tfsdk:"response_types" yaml:"responseTypes,omitempty"`

						RevocationEndpoint *string `tfsdk:"revocation_endpoint" yaml:"revocationEndpoint,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						Subjects *[]string `tfsdk:"subjects" yaml:"subjects,omitempty"`

						TokenEndpoint *string `tfsdk:"token_endpoint" yaml:"tokenEndpoint,omitempty"`
					} `tfsdk:"discovery_override" yaml:"discoveryOverride,omitempty"`

					DiscoveryPollInterval *string `tfsdk:"discovery_poll_interval" yaml:"discoveryPollInterval,omitempty"`

					EndSessionProperties *struct {
						MethodType utilities.IntOrString `tfsdk:"method_type" yaml:"methodType,omitempty"`
					} `tfsdk:"end_session_properties" yaml:"endSessionProperties,omitempty"`

					Headers *struct {
						AccessTokenHeader *string `tfsdk:"access_token_header" yaml:"accessTokenHeader,omitempty"`

						IdTokenHeader *string `tfsdk:"id_token_header" yaml:"idTokenHeader,omitempty"`

						UseBearerSchemaForAuthorization *bool `tfsdk:"use_bearer_schema_for_authorization" yaml:"useBearerSchemaForAuthorization,omitempty"`
					} `tfsdk:"headers" yaml:"headers,omitempty"`

					IssuerUrl *string `tfsdk:"issuer_url" yaml:"issuerUrl,omitempty"`

					JwksCacheRefreshPolicy *struct {
						Always *map[string]string `tfsdk:"always" yaml:"always,omitempty"`

						MaxIdpReqPerPollingInterval *int64 `tfsdk:"max_idp_req_per_polling_interval" yaml:"maxIdpReqPerPollingInterval,omitempty"`

						Never *map[string]string `tfsdk:"never" yaml:"never,omitempty"`
					} `tfsdk:"jwks_cache_refresh_policy" yaml:"jwksCacheRefreshPolicy,omitempty"`

					LogoutPath *string `tfsdk:"logout_path" yaml:"logoutPath,omitempty"`

					ParseCallbackPathAsRegex *bool `tfsdk:"parse_callback_path_as_regex" yaml:"parseCallbackPathAsRegex,omitempty"`

					Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

					Session *struct {
						Cookie *struct {
							AllowRefreshing *bool `tfsdk:"allow_refreshing" yaml:"allowRefreshing,omitempty"`

							KeyPrefix *string `tfsdk:"key_prefix" yaml:"keyPrefix,omitempty"`

							TargetDomain *string `tfsdk:"target_domain" yaml:"targetDomain,omitempty"`
						} `tfsdk:"cookie" yaml:"cookie,omitempty"`

						CookieOptions *struct {
							Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

							HttpOnly *bool `tfsdk:"http_only" yaml:"httpOnly,omitempty"`

							MaxAge *int64 `tfsdk:"max_age" yaml:"maxAge,omitempty"`

							NotSecure *bool `tfsdk:"not_secure" yaml:"notSecure,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							SameSite utilities.IntOrString `tfsdk:"same_site" yaml:"sameSite,omitempty"`
						} `tfsdk:"cookie_options" yaml:"cookieOptions,omitempty"`

						FailOnFetchFailure *bool `tfsdk:"fail_on_fetch_failure" yaml:"failOnFetchFailure,omitempty"`

						Redis *struct {
							AllowRefreshing *bool `tfsdk:"allow_refreshing" yaml:"allowRefreshing,omitempty"`

							CookieName *string `tfsdk:"cookie_name" yaml:"cookieName,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`

							KeyPrefix *string `tfsdk:"key_prefix" yaml:"keyPrefix,omitempty"`

							Options *struct {
								Db *int64 `tfsdk:"db" yaml:"db,omitempty"`

								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								PoolSize *int64 `tfsdk:"pool_size" yaml:"poolSize,omitempty"`

								SocketType utilities.IntOrString `tfsdk:"socket_type" yaml:"socketType,omitempty"`

								TlsCertMountPath *string `tfsdk:"tls_cert_mount_path" yaml:"tlsCertMountPath,omitempty"`
							} `tfsdk:"options" yaml:"options,omitempty"`

							PreExpiryBuffer *string `tfsdk:"pre_expiry_buffer" yaml:"preExpiryBuffer,omitempty"`

							TargetDomain *string `tfsdk:"target_domain" yaml:"targetDomain,omitempty"`
						} `tfsdk:"redis" yaml:"redis,omitempty"`
					} `tfsdk:"session" yaml:"session,omitempty"`

					SessionIdHeaderName *string `tfsdk:"session_id_header_name" yaml:"sessionIdHeaderName,omitempty"`

					TokenEndpointQueryParams *map[string]string `tfsdk:"token_endpoint_query_params" yaml:"tokenEndpointQueryParams,omitempty"`
				} `tfsdk:"oidc_authorization_code" yaml:"oidcAuthorizationCode,omitempty"`
			} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

			OpaAuth *struct {
				Modules *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"modules" yaml:"modules,omitempty"`

				Options *struct {
					FastInputConversion *bool `tfsdk:"fast_input_conversion" yaml:"fastInputConversion,omitempty"`
				} `tfsdk:"options" yaml:"options,omitempty"`

				Query *string `tfsdk:"query" yaml:"query,omitempty"`
			} `tfsdk:"opa_auth" yaml:"opaAuth,omitempty"`

			PassThroughAuth *struct {
				Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

				FailureModeAllow *bool `tfsdk:"failure_mode_allow" yaml:"failureModeAllow,omitempty"`

				Grpc *struct {
					Address *string `tfsdk:"address" yaml:"address,omitempty"`

					ConnectionTimeout *string `tfsdk:"connection_timeout" yaml:"connectionTimeout,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				Http *struct {
					ConnectionTimeout *string `tfsdk:"connection_timeout" yaml:"connectionTimeout,omitempty"`

					Request *struct {
						AllowedHeaders *[]string `tfsdk:"allowed_headers" yaml:"allowedHeaders,omitempty"`

						HeadersToAdd *map[string]string `tfsdk:"headers_to_add" yaml:"headersToAdd,omitempty"`

						PassThroughBody *bool `tfsdk:"pass_through_body" yaml:"passThroughBody,omitempty"`

						PassThroughFilterMetadata *bool `tfsdk:"pass_through_filter_metadata" yaml:"passThroughFilterMetadata,omitempty"`

						PassThroughState *bool `tfsdk:"pass_through_state" yaml:"passThroughState,omitempty"`
					} `tfsdk:"request" yaml:"request,omitempty"`

					Response *struct {
						AllowedClientHeadersOnDenied *[]string `tfsdk:"allowed_client_headers_on_denied" yaml:"allowedClientHeadersOnDenied,omitempty"`

						AllowedUpstreamHeaders *[]string `tfsdk:"allowed_upstream_headers" yaml:"allowedUpstreamHeaders,omitempty"`

						ReadStateFromResponse *bool `tfsdk:"read_state_from_response" yaml:"readStateFromResponse,omitempty"`
					} `tfsdk:"response" yaml:"response,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"http" yaml:"http,omitempty"`
			} `tfsdk:"pass_through_auth" yaml:"passThroughAuth,omitempty"`

			PluginAuth *struct {
				Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

				ExportedSymbolName *string `tfsdk:"exported_symbol_name" yaml:"exportedSymbolName,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				PluginFileName *string `tfsdk:"plugin_file_name" yaml:"pluginFileName,omitempty"`
			} `tfsdk:"plugin_auth" yaml:"pluginAuth,omitempty"`
		} `tfsdk:"configs" yaml:"configs,omitempty"`

		FailOnRedirect *bool `tfsdk:"fail_on_redirect" yaml:"failOnRedirect,omitempty"`

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewEnterpriseGlooSoloIoAuthConfigV1Resource() resource.Resource {
	return &EnterpriseGlooSoloIoAuthConfigV1Resource{}
}

func (r *EnterpriseGlooSoloIoAuthConfigV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_enterprise_gloo_solo_io_auth_config_v1"
}

func (r *EnterpriseGlooSoloIoAuthConfigV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

					"boolean_expr": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"configs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"api_key_auth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"aerospike_apikey_storage": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_insecure": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"batch_size": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"commit_all": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"commit_master": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key_path": {
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

											"node_tls_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_mode_ap": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"read_mode_ap_all": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_mode_ap_one": {
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

											"read_mode_sc": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"read_mode_sc_allow_unavailable": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_mode_sc_linearize": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_mode_sc_replica": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_mode_sc_session": {
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

											"root_ca_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_curve_groups": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"curve_p256": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"curve_p384": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"curve_p521": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"x25519": {
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

											"tls_version": {
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

									"api_key_secret_refs": {
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

									"header_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers_from_metadata": {
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

											"required": {
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

									"headers_from_metadata_entry": {
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

											"required": {
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

									"k8s_secret_apikey_storage": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"api_key_secret_refs": {
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

											"label_selector": {
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

									"label_selector": {
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

							"basic_auth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"apr": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"users": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"hashed_password": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"salt": {
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

									"realm": {
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

							"jwt": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ldap": {
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

									"allowed_groups": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_group_checking": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"membership_attribute_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pool": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"initial_size": {
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

											"max_size": {
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

									"search_filter": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_dn_template": {
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

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"oauth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"app_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth_endpoint_query_params": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"callback_path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret_ref": {
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

									"issuer_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scopes": {
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

							"oauth2": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_token_validation": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cache_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"introspection": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_secret_ref": {
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

													"introspection_url": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"user_id_attribute_name": {
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

											"introspection_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwt": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"issuer": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"local_jwks": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"inline_string": {
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

													"remote_jwks": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"refresh_interval": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"url": {
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

											"required_scopes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"scope": {
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

											"userinfo_url": {
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

									"oauth2": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"after_logout_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"app_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth_endpoint": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth_endpoint_query_params": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"callback_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_secret_ref": {
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

											"logout_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"revocation_endpoint": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scopes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"session": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cookie": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_refreshing": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"key_prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"target_domain": {
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

													"cookie_options": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"domain": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_age": {
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

															"not_secure": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

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

															"same_site": {
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

													"fail_on_fetch_failure": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"redis": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_refreshing": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cookie_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"key_prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"db": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"pool_size": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"socket_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"tls_cert_mount_path": {
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

															"pre_expiry_buffer": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"target_domain": {
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

											"token_endpoint": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_endpoint_query_params": {
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

									"oidc_authorization_code": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"after_logout_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"app_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth_endpoint_query_params": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_map_from_metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"callback_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_secret_ref": {
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

											"discovery_override": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"auth_endpoint": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"auth_methods": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"claims": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"end_session_endpoint": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id_token_algs": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jwks_uri": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"response_types": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"revocation_endpoint": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"subjects": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_endpoint": {
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

											"discovery_poll_interval": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"end_session_properties": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"method_type": {
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

											"headers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_token_header": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id_token_header": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_bearer_schema_for_authorization": {
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

											"issuer_url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwks_cache_refresh_policy": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"always": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_idp_req_per_polling_interval": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"never": {
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

											"logout_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"parse_callback_path_as_regex": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scopes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"session": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cookie": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_refreshing": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"key_prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"target_domain": {
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

													"cookie_options": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"domain": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_age": {
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

															"not_secure": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

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

															"same_site": {
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

													"fail_on_fetch_failure": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"redis": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_refreshing": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cookie_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"key_prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"db": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"pool_size": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"socket_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"tls_cert_mount_path": {
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

															"pre_expiry_buffer": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"target_domain": {
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

											"session_id_header_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_endpoint_query_params": {
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

							"opa_auth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"modules": {
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

									"options": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fast_input_conversion": {
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

									"query": {
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

							"pass_through_auth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

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

									"grpc": {
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

											"connection_timeout": {
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

									"http": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"connection_timeout": {
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

													"headers_to_add": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass_through_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass_through_filter_metadata": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass_through_state": {
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

											"response": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allowed_client_headers_on_denied": {
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

													"read_state_from_response": {
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

											"url": {
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

							"plugin_auth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"exported_symbol_name": {
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

									"plugin_file_name": {
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

					"fail_on_redirect": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

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

func (r *EnterpriseGlooSoloIoAuthConfigV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_enterprise_gloo_solo_io_auth_config_v1")

	var state EnterpriseGlooSoloIoAuthConfigV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EnterpriseGlooSoloIoAuthConfigV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("enterprise.gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("AuthConfig")

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

func (r *EnterpriseGlooSoloIoAuthConfigV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_enterprise_gloo_solo_io_auth_config_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *EnterpriseGlooSoloIoAuthConfigV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_enterprise_gloo_solo_io_auth_config_v1")

	var state EnterpriseGlooSoloIoAuthConfigV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EnterpriseGlooSoloIoAuthConfigV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("enterprise.gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("AuthConfig")

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

func (r *EnterpriseGlooSoloIoAuthConfigV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_enterprise_gloo_solo_io_auth_config_v1")
	// NO-OP: Terraform removes the state automatically for us
}
