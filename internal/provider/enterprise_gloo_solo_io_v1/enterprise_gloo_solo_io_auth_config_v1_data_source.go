/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package enterprise_gloo_solo_io_v1

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
	_ datasource.DataSource              = &EnterpriseGlooSoloIoAuthConfigV1DataSource{}
	_ datasource.DataSourceWithConfigure = &EnterpriseGlooSoloIoAuthConfigV1DataSource{}
)

func NewEnterpriseGlooSoloIoAuthConfigV1DataSource() datasource.DataSource {
	return &EnterpriseGlooSoloIoAuthConfigV1DataSource{}
}

type EnterpriseGlooSoloIoAuthConfigV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type EnterpriseGlooSoloIoAuthConfigV1DataSourceData struct {
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
		BooleanExpr *string `tfsdk:"boolean_expr" json:"booleanExpr,omitempty"`
		Configs     *[]struct {
			ApiKeyAuth *struct {
				AerospikeApikeyStorage *struct {
					AllowInsecure *bool   `tfsdk:"allow_insecure" json:"allowInsecure,omitempty"`
					BatchSize     *int64  `tfsdk:"batch_size" json:"batchSize,omitempty"`
					CertPath      *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					CommitAll     *int64  `tfsdk:"commit_all" json:"commitAll,omitempty"`
					CommitMaster  *int64  `tfsdk:"commit_master" json:"commitMaster,omitempty"`
					Hostname      *string `tfsdk:"hostname" json:"hostname,omitempty"`
					KeyPath       *string `tfsdk:"key_path" json:"keyPath,omitempty"`
					Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
					NodeTlsName   *string `tfsdk:"node_tls_name" json:"nodeTlsName,omitempty"`
					Port          *int64  `tfsdk:"port" json:"port,omitempty"`
					ReadModeAp    *struct {
						ReadModeApAll *int64 `tfsdk:"read_mode_ap_all" json:"readModeApAll,omitempty"`
						ReadModeApOne *int64 `tfsdk:"read_mode_ap_one" json:"readModeApOne,omitempty"`
					} `tfsdk:"read_mode_ap" json:"readModeAp,omitempty"`
					ReadModeSc *struct {
						ReadModeScAllowUnavailable *int64 `tfsdk:"read_mode_sc_allow_unavailable" json:"readModeScAllowUnavailable,omitempty"`
						ReadModeScLinearize        *int64 `tfsdk:"read_mode_sc_linearize" json:"readModeScLinearize,omitempty"`
						ReadModeScReplica          *int64 `tfsdk:"read_mode_sc_replica" json:"readModeScReplica,omitempty"`
						ReadModeScSession          *int64 `tfsdk:"read_mode_sc_session" json:"readModeScSession,omitempty"`
					} `tfsdk:"read_mode_sc" json:"readModeSc,omitempty"`
					RootCaPath     *string `tfsdk:"root_ca_path" json:"rootCaPath,omitempty"`
					Set            *string `tfsdk:"set" json:"set,omitempty"`
					TlsCurveGroups *[]struct {
						CurveP256 *int64 `tfsdk:"curve_p256" json:"curveP256,omitempty"`
						CurveP384 *int64 `tfsdk:"curve_p384" json:"curveP384,omitempty"`
						CurveP521 *int64 `tfsdk:"curve_p521" json:"curveP521,omitempty"`
						X25519    *int64 `tfsdk:"x25519" json:"x25519,omitempty"`
					} `tfsdk:"tls_curve_groups" json:"tlsCurveGroups,omitempty"`
					TlsVersion *string `tfsdk:"tls_version" json:"tlsVersion,omitempty"`
				} `tfsdk:"aerospike_apikey_storage" json:"aerospikeApikeyStorage,omitempty"`
				ApiKeySecretRefs *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"api_key_secret_refs" json:"apiKeySecretRefs,omitempty"`
				HeaderName          *string `tfsdk:"header_name" json:"headerName,omitempty"`
				HeadersFromMetadata *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Required *bool   `tfsdk:"required" json:"required,omitempty"`
				} `tfsdk:"headers_from_metadata" json:"headersFromMetadata,omitempty"`
				HeadersFromMetadataEntry *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Required *bool   `tfsdk:"required" json:"required,omitempty"`
				} `tfsdk:"headers_from_metadata_entry" json:"headersFromMetadataEntry,omitempty"`
				K8sSecretApikeyStorage *struct {
					ApiKeySecretRefs *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"api_key_secret_refs" json:"apiKeySecretRefs,omitempty"`
					LabelSelector *map[string]string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				} `tfsdk:"k8s_secret_apikey_storage" json:"k8sSecretApikeyStorage,omitempty"`
				LabelSelector *map[string]string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			} `tfsdk:"api_key_auth" json:"apiKeyAuth,omitempty"`
			BasicAuth *struct {
				Apr *struct {
					Users *struct {
						HashedPassword *string `tfsdk:"hashed_password" json:"hashedPassword,omitempty"`
						Salt           *string `tfsdk:"salt" json:"salt,omitempty"`
					} `tfsdk:"users" json:"users,omitempty"`
				} `tfsdk:"apr" json:"apr,omitempty"`
				Realm *string `tfsdk:"realm" json:"realm,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			HmacAuth *struct {
				ParametersInHeaders *map[string]string `tfsdk:"parameters_in_headers" json:"parametersInHeaders,omitempty"`
				SecretRefs          *struct {
					SecretRefs *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
			} `tfsdk:"hmac_auth" json:"hmacAuth,omitempty"`
			Jwt  *map[string]string `tfsdk:"jwt" json:"jwt,omitempty"`
			Ldap *struct {
				Address              *string   `tfsdk:"address" json:"address,omitempty"`
				AllowedGroups        *[]string `tfsdk:"allowed_groups" json:"allowedGroups,omitempty"`
				DisableGroupChecking *bool     `tfsdk:"disable_group_checking" json:"disableGroupChecking,omitempty"`
				GroupLookupSettings  *struct {
					CheckGroupsWithServiceAccount *bool `tfsdk:"check_groups_with_service_account" json:"checkGroupsWithServiceAccount,omitempty"`
					CredentialsSecretRef          *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				} `tfsdk:"group_lookup_settings" json:"groupLookupSettings,omitempty"`
				MembershipAttributeName *string `tfsdk:"membership_attribute_name" json:"membershipAttributeName,omitempty"`
				Pool                    *struct {
					InitialSize *int64 `tfsdk:"initial_size" json:"initialSize,omitempty"`
					MaxSize     *int64 `tfsdk:"max_size" json:"maxSize,omitempty"`
				} `tfsdk:"pool" json:"pool,omitempty"`
				SearchFilter   *string `tfsdk:"search_filter" json:"searchFilter,omitempty"`
				UserDnTemplate *string `tfsdk:"user_dn_template" json:"userDnTemplate,omitempty"`
			} `tfsdk:"ldap" json:"ldap,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Oauth *struct {
				AppUrl                  *string            `tfsdk:"app_url" json:"appUrl,omitempty"`
				AuthEndpointQueryParams *map[string]string `tfsdk:"auth_endpoint_query_params" json:"authEndpointQueryParams,omitempty"`
				CallbackPath            *string            `tfsdk:"callback_path" json:"callbackPath,omitempty"`
				ClientId                *string            `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientSecretRef         *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
				IssuerUrl *string   `tfsdk:"issuer_url" json:"issuerUrl,omitempty"`
				Scopes    *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
			} `tfsdk:"oauth" json:"oauth,omitempty"`
			Oauth2 *struct {
				AccessTokenValidation *struct {
					CacheTimeout              *string            `tfsdk:"cache_timeout" json:"cacheTimeout,omitempty"`
					DynamicMetadataFromClaims *map[string]string `tfsdk:"dynamic_metadata_from_claims" json:"dynamicMetadataFromClaims,omitempty"`
					Introspection             *struct {
						ClientId        *string `tfsdk:"client_id" json:"clientId,omitempty"`
						ClientSecretRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
						DisableClientSecret *bool   `tfsdk:"disable_client_secret" json:"disableClientSecret,omitempty"`
						IntrospectionUrl    *string `tfsdk:"introspection_url" json:"introspectionUrl,omitempty"`
						UserIdAttributeName *string `tfsdk:"user_id_attribute_name" json:"userIdAttributeName,omitempty"`
					} `tfsdk:"introspection" json:"introspection,omitempty"`
					IntrospectionUrl *string `tfsdk:"introspection_url" json:"introspectionUrl,omitempty"`
					Jwt              *struct {
						Issuer    *string `tfsdk:"issuer" json:"issuer,omitempty"`
						LocalJwks *struct {
							InlineString *string `tfsdk:"inline_string" json:"inlineString,omitempty"`
						} `tfsdk:"local_jwks" json:"localJwks,omitempty"`
						RemoteJwks *struct {
							RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
							Url             *string `tfsdk:"url" json:"url,omitempty"`
						} `tfsdk:"remote_jwks" json:"remoteJwks,omitempty"`
					} `tfsdk:"jwt" json:"jwt,omitempty"`
					RequiredScopes *struct {
						Scope *[]string `tfsdk:"scope" json:"scope,omitempty"`
					} `tfsdk:"required_scopes" json:"requiredScopes,omitempty"`
					UserinfoUrl *string `tfsdk:"userinfo_url" json:"userinfoUrl,omitempty"`
				} `tfsdk:"access_token_validation" json:"accessTokenValidation,omitempty"`
				Oauth2 *struct {
					AfterLogoutUrl          *string            `tfsdk:"after_logout_url" json:"afterLogoutUrl,omitempty"`
					AppUrl                  *string            `tfsdk:"app_url" json:"appUrl,omitempty"`
					AuthEndpoint            *string            `tfsdk:"auth_endpoint" json:"authEndpoint,omitempty"`
					AuthEndpointQueryParams *map[string]string `tfsdk:"auth_endpoint_query_params" json:"authEndpointQueryParams,omitempty"`
					CallbackPath            *string            `tfsdk:"callback_path" json:"callbackPath,omitempty"`
					ClientId                *string            `tfsdk:"client_id" json:"clientId,omitempty"`
					ClientSecretRef         *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
					DisableClientSecret *bool     `tfsdk:"disable_client_secret" json:"disableClientSecret,omitempty"`
					LogoutPath          *string   `tfsdk:"logout_path" json:"logoutPath,omitempty"`
					RevocationEndpoint  *string   `tfsdk:"revocation_endpoint" json:"revocationEndpoint,omitempty"`
					Scopes              *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
					Session             *struct {
						CipherConfig *struct {
							KeyRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"key_ref" json:"keyRef,omitempty"`
						} `tfsdk:"cipher_config" json:"cipherConfig,omitempty"`
						Cookie *struct {
							AllowRefreshing *bool   `tfsdk:"allow_refreshing" json:"allowRefreshing,omitempty"`
							KeyPrefix       *string `tfsdk:"key_prefix" json:"keyPrefix,omitempty"`
							TargetDomain    *string `tfsdk:"target_domain" json:"targetDomain,omitempty"`
						} `tfsdk:"cookie" json:"cookie,omitempty"`
						CookieOptions *struct {
							Domain    *string `tfsdk:"domain" json:"domain,omitempty"`
							HttpOnly  *bool   `tfsdk:"http_only" json:"httpOnly,omitempty"`
							MaxAge    *int64  `tfsdk:"max_age" json:"maxAge,omitempty"`
							NotSecure *bool   `tfsdk:"not_secure" json:"notSecure,omitempty"`
							Path      *string `tfsdk:"path" json:"path,omitempty"`
							SameSite  *string `tfsdk:"same_site" json:"sameSite,omitempty"`
						} `tfsdk:"cookie_options" json:"cookieOptions,omitempty"`
						FailOnFetchFailure *bool `tfsdk:"fail_on_fetch_failure" json:"failOnFetchFailure,omitempty"`
						Redis              *struct {
							AllowRefreshing *bool   `tfsdk:"allow_refreshing" json:"allowRefreshing,omitempty"`
							CookieName      *string `tfsdk:"cookie_name" json:"cookieName,omitempty"`
							HeaderName      *string `tfsdk:"header_name" json:"headerName,omitempty"`
							KeyPrefix       *string `tfsdk:"key_prefix" json:"keyPrefix,omitempty"`
							Options         *struct {
								Db               *int64  `tfsdk:"db" json:"db,omitempty"`
								Host             *string `tfsdk:"host" json:"host,omitempty"`
								PoolSize         *int64  `tfsdk:"pool_size" json:"poolSize,omitempty"`
								SocketType       *string `tfsdk:"socket_type" json:"socketType,omitempty"`
								TlsCertMountPath *string `tfsdk:"tls_cert_mount_path" json:"tlsCertMountPath,omitempty"`
							} `tfsdk:"options" json:"options,omitempty"`
							PreExpiryBuffer *string `tfsdk:"pre_expiry_buffer" json:"preExpiryBuffer,omitempty"`
							TargetDomain    *string `tfsdk:"target_domain" json:"targetDomain,omitempty"`
						} `tfsdk:"redis" json:"redis,omitempty"`
					} `tfsdk:"session" json:"session,omitempty"`
					TokenEndpoint            *string            `tfsdk:"token_endpoint" json:"tokenEndpoint,omitempty"`
					TokenEndpointQueryParams *map[string]string `tfsdk:"token_endpoint_query_params" json:"tokenEndpointQueryParams,omitempty"`
				} `tfsdk:"oauth2" json:"oauth2,omitempty"`
				OidcAuthorizationCode *struct {
					AccessToken *struct {
						ClaimsToHeaders *[]struct {
							Append *bool   `tfsdk:"append" json:"append,omitempty"`
							Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
							Header *string `tfsdk:"header" json:"header,omitempty"`
						} `tfsdk:"claims_to_headers" json:"claimsToHeaders,omitempty"`
					} `tfsdk:"access_token" json:"accessToken,omitempty"`
					AfterLogoutUrl          *string            `tfsdk:"after_logout_url" json:"afterLogoutUrl,omitempty"`
					AppUrl                  *string            `tfsdk:"app_url" json:"appUrl,omitempty"`
					AuthEndpointQueryParams *map[string]string `tfsdk:"auth_endpoint_query_params" json:"authEndpointQueryParams,omitempty"`
					AutoMapFromMetadata     *struct {
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"auto_map_from_metadata" json:"autoMapFromMetadata,omitempty"`
					CallbackPath    *string `tfsdk:"callback_path" json:"callbackPath,omitempty"`
					ClientId        *string `tfsdk:"client_id" json:"clientId,omitempty"`
					ClientSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
					DisableClientSecret *bool `tfsdk:"disable_client_secret" json:"disableClientSecret,omitempty"`
					DiscoveryOverride   *struct {
						AuthEndpoint       *string   `tfsdk:"auth_endpoint" json:"authEndpoint,omitempty"`
						AuthMethods        *[]string `tfsdk:"auth_methods" json:"authMethods,omitempty"`
						Claims             *[]string `tfsdk:"claims" json:"claims,omitempty"`
						EndSessionEndpoint *string   `tfsdk:"end_session_endpoint" json:"endSessionEndpoint,omitempty"`
						IdTokenAlgs        *[]string `tfsdk:"id_token_algs" json:"idTokenAlgs,omitempty"`
						JwksUri            *string   `tfsdk:"jwks_uri" json:"jwksUri,omitempty"`
						ResponseTypes      *[]string `tfsdk:"response_types" json:"responseTypes,omitempty"`
						RevocationEndpoint *string   `tfsdk:"revocation_endpoint" json:"revocationEndpoint,omitempty"`
						Scopes             *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
						Subjects           *[]string `tfsdk:"subjects" json:"subjects,omitempty"`
						TokenEndpoint      *string   `tfsdk:"token_endpoint" json:"tokenEndpoint,omitempty"`
					} `tfsdk:"discovery_override" json:"discoveryOverride,omitempty"`
					DiscoveryPollInterval     *string            `tfsdk:"discovery_poll_interval" json:"discoveryPollInterval,omitempty"`
					DynamicMetadataFromClaims *map[string]string `tfsdk:"dynamic_metadata_from_claims" json:"dynamicMetadataFromClaims,omitempty"`
					EndSessionProperties      *struct {
						MethodType *string `tfsdk:"method_type" json:"methodType,omitempty"`
					} `tfsdk:"end_session_properties" json:"endSessionProperties,omitempty"`
					Headers *struct {
						AccessTokenHeader               *string `tfsdk:"access_token_header" json:"accessTokenHeader,omitempty"`
						IdTokenHeader                   *string `tfsdk:"id_token_header" json:"idTokenHeader,omitempty"`
						UseBearerSchemaForAuthorization *bool   `tfsdk:"use_bearer_schema_for_authorization" json:"useBearerSchemaForAuthorization,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					IdentityToken *struct {
						ClaimsToHeaders *[]struct {
							Append *bool   `tfsdk:"append" json:"append,omitempty"`
							Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
							Header *string `tfsdk:"header" json:"header,omitempty"`
						} `tfsdk:"claims_to_headers" json:"claimsToHeaders,omitempty"`
					} `tfsdk:"identity_token" json:"identityToken,omitempty"`
					IssuerUrl              *string `tfsdk:"issuer_url" json:"issuerUrl,omitempty"`
					JwksCacheRefreshPolicy *struct {
						Always                      *map[string]string `tfsdk:"always" json:"always,omitempty"`
						MaxIdpReqPerPollingInterval *int64             `tfsdk:"max_idp_req_per_polling_interval" json:"maxIdpReqPerPollingInterval,omitempty"`
						Never                       *map[string]string `tfsdk:"never" json:"never,omitempty"`
					} `tfsdk:"jwks_cache_refresh_policy" json:"jwksCacheRefreshPolicy,omitempty"`
					LogoutPath               *string   `tfsdk:"logout_path" json:"logoutPath,omitempty"`
					ParseCallbackPathAsRegex *bool     `tfsdk:"parse_callback_path_as_regex" json:"parseCallbackPathAsRegex,omitempty"`
					Scopes                   *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
					Session                  *struct {
						CipherConfig *struct {
							KeyRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"key_ref" json:"keyRef,omitempty"`
						} `tfsdk:"cipher_config" json:"cipherConfig,omitempty"`
						Cookie *struct {
							AllowRefreshing *bool   `tfsdk:"allow_refreshing" json:"allowRefreshing,omitempty"`
							KeyPrefix       *string `tfsdk:"key_prefix" json:"keyPrefix,omitempty"`
							TargetDomain    *string `tfsdk:"target_domain" json:"targetDomain,omitempty"`
						} `tfsdk:"cookie" json:"cookie,omitempty"`
						CookieOptions *struct {
							Domain    *string `tfsdk:"domain" json:"domain,omitempty"`
							HttpOnly  *bool   `tfsdk:"http_only" json:"httpOnly,omitempty"`
							MaxAge    *int64  `tfsdk:"max_age" json:"maxAge,omitempty"`
							NotSecure *bool   `tfsdk:"not_secure" json:"notSecure,omitempty"`
							Path      *string `tfsdk:"path" json:"path,omitempty"`
							SameSite  *string `tfsdk:"same_site" json:"sameSite,omitempty"`
						} `tfsdk:"cookie_options" json:"cookieOptions,omitempty"`
						FailOnFetchFailure *bool `tfsdk:"fail_on_fetch_failure" json:"failOnFetchFailure,omitempty"`
						Redis              *struct {
							AllowRefreshing *bool   `tfsdk:"allow_refreshing" json:"allowRefreshing,omitempty"`
							CookieName      *string `tfsdk:"cookie_name" json:"cookieName,omitempty"`
							HeaderName      *string `tfsdk:"header_name" json:"headerName,omitempty"`
							KeyPrefix       *string `tfsdk:"key_prefix" json:"keyPrefix,omitempty"`
							Options         *struct {
								Db               *int64  `tfsdk:"db" json:"db,omitempty"`
								Host             *string `tfsdk:"host" json:"host,omitempty"`
								PoolSize         *int64  `tfsdk:"pool_size" json:"poolSize,omitempty"`
								SocketType       *string `tfsdk:"socket_type" json:"socketType,omitempty"`
								TlsCertMountPath *string `tfsdk:"tls_cert_mount_path" json:"tlsCertMountPath,omitempty"`
							} `tfsdk:"options" json:"options,omitempty"`
							PreExpiryBuffer *string `tfsdk:"pre_expiry_buffer" json:"preExpiryBuffer,omitempty"`
							TargetDomain    *string `tfsdk:"target_domain" json:"targetDomain,omitempty"`
						} `tfsdk:"redis" json:"redis,omitempty"`
					} `tfsdk:"session" json:"session,omitempty"`
					SessionIdHeaderName      *string            `tfsdk:"session_id_header_name" json:"sessionIdHeaderName,omitempty"`
					TokenEndpointQueryParams *map[string]string `tfsdk:"token_endpoint_query_params" json:"tokenEndpointQueryParams,omitempty"`
				} `tfsdk:"oidc_authorization_code" json:"oidcAuthorizationCode,omitempty"`
			} `tfsdk:"oauth2" json:"oauth2,omitempty"`
			OpaAuth *struct {
				Modules *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"modules" json:"modules,omitempty"`
				Options *struct {
					FastInputConversion  *bool `tfsdk:"fast_input_conversion" json:"fastInputConversion,omitempty"`
					ReturnDecisionReason *bool `tfsdk:"return_decision_reason" json:"returnDecisionReason,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Query *string `tfsdk:"query" json:"query,omitempty"`
			} `tfsdk:"opa_auth" json:"opaAuth,omitempty"`
			OpaServerAuth *struct {
				Options *struct {
					FastInputConversion  *bool `tfsdk:"fast_input_conversion" json:"fastInputConversion,omitempty"`
					ReturnDecisionReason *bool `tfsdk:"return_decision_reason" json:"returnDecisionReason,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Package    *string `tfsdk:"package" json:"package,omitempty"`
				RuleName   *string `tfsdk:"rule_name" json:"ruleName,omitempty"`
				ServerAddr *string `tfsdk:"server_addr" json:"serverAddr,omitempty"`
			} `tfsdk:"opa_server_auth" json:"opaServerAuth,omitempty"`
			PassThroughAuth *struct {
				Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
				FailureModeAllow *bool              `tfsdk:"failure_mode_allow" json:"failureModeAllow,omitempty"`
				Grpc             *struct {
					Address           *string            `tfsdk:"address" json:"address,omitempty"`
					ConnectionTimeout *string            `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
					TlsConfig         *map[string]string `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"grpc" json:"grpc,omitempty"`
				Http *struct {
					ConnectionTimeout *string `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
					Request           *struct {
						AllowedHeaders            *[]string          `tfsdk:"allowed_headers" json:"allowedHeaders,omitempty"`
						HeadersToAdd              *map[string]string `tfsdk:"headers_to_add" json:"headersToAdd,omitempty"`
						PassThroughBody           *bool              `tfsdk:"pass_through_body" json:"passThroughBody,omitempty"`
						PassThroughFilterMetadata *bool              `tfsdk:"pass_through_filter_metadata" json:"passThroughFilterMetadata,omitempty"`
						PassThroughState          *bool              `tfsdk:"pass_through_state" json:"passThroughState,omitempty"`
					} `tfsdk:"request" json:"request,omitempty"`
					Response *struct {
						AllowedClientHeadersOnDenied      *[]string `tfsdk:"allowed_client_headers_on_denied" json:"allowedClientHeadersOnDenied,omitempty"`
						AllowedUpstreamHeaders            *[]string `tfsdk:"allowed_upstream_headers" json:"allowedUpstreamHeaders,omitempty"`
						AllowedUpstreamHeadersToOverwrite *[]string `tfsdk:"allowed_upstream_headers_to_overwrite" json:"allowedUpstreamHeadersToOverwrite,omitempty"`
						ReadStateFromResponse             *bool     `tfsdk:"read_state_from_response" json:"readStateFromResponse,omitempty"`
					} `tfsdk:"response" json:"response,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
			} `tfsdk:"pass_through_auth" json:"passThroughAuth,omitempty"`
			PluginAuth *struct {
				Config             *map[string]string `tfsdk:"config" json:"config,omitempty"`
				ExportedSymbolName *string            `tfsdk:"exported_symbol_name" json:"exportedSymbolName,omitempty"`
				Name               *string            `tfsdk:"name" json:"name,omitempty"`
				PluginFileName     *string            `tfsdk:"plugin_file_name" json:"pluginFileName,omitempty"`
			} `tfsdk:"plugin_auth" json:"pluginAuth,omitempty"`
		} `tfsdk:"configs" json:"configs,omitempty"`
		FailOnRedirect     *bool `tfsdk:"fail_on_redirect" json:"failOnRedirect,omitempty"`
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EnterpriseGlooSoloIoAuthConfigV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_enterprise_gloo_solo_io_auth_config_v1"
}

func (r *EnterpriseGlooSoloIoAuthConfigV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"boolean_expr": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"configs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_key_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"aerospike_apikey_storage": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_insecure": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"batch_size": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"cert_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"commit_all": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"commit_master": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"hostname": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"key_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_tls_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_mode_ap": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"read_mode_ap_all": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_mode_ap_one": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"read_mode_sc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"read_mode_sc_allow_unavailable": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_mode_sc_linearize": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_mode_sc_replica": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_mode_sc_session": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"root_ca_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"set": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tls_curve_groups": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"curve_p256": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"curve_p384": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"curve_p521": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"x25519": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
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

												"tls_version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"api_key_secret_refs": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

										"header_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"headers_from_metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"required": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"headers_from_metadata_entry": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"required": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"k8s_secret_apikey_storage": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"api_key_secret_refs": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"label_selector": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"label_selector": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"basic_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"apr": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"users": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"hashed_password": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"salt": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

										"realm": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"hmac_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"parameters_in_headers": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_refs": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"secret_refs": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"jwt": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ldap": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"allowed_groups": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"disable_group_checking": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"group_lookup_settings": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"check_groups_with_service_account": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"credentials_secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

										"membership_attribute_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"pool": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"initial_size": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_size": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"search_filter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"user_dn_template": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"oauth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"app_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"auth_endpoint_query_params": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"callback_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"issuer_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"oauth2": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"access_token_validation": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"cache_timeout": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"dynamic_metadata_from_claims": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"introspection": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"client_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"client_secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"disable_client_secret": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"introspection_url": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"user_id_attribute_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"introspection_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"jwt": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"issuer": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"local_jwks": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"inline_string": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"remote_jwks": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"refresh_interval": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"url": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"required_scopes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"scope": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
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

												"userinfo_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"oauth2": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"after_logout_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"app_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"auth_endpoint": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"auth_endpoint_query_params": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"callback_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"client_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"client_secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"disable_client_secret": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"logout_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"revocation_endpoint": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"scopes": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"session": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cipher_config": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

														"cookie": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"allow_refreshing": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"key_prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"target_domain": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"cookie_options": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"domain": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"http_only": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_age": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"not_secure": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"same_site": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"fail_on_fetch_failure": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"redis": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"allow_refreshing": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"cookie_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"header_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"key_prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"options": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"db": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"host": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"pool_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"socket_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tls_cert_mount_path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"pre_expiry_buffer": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"target_domain": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"token_endpoint": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"token_endpoint_query_params": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"oidc_authorization_code": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_token": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"claims_to_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"append": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"claim": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

												"after_logout_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"app_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"auth_endpoint_query_params": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"auto_map_from_metadata": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"callback_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"client_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"client_secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"disable_client_secret": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"discovery_override": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auth_endpoint": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"auth_methods": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"claims": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"end_session_endpoint": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"id_token_algs": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"jwks_uri": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"response_types": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"revocation_endpoint": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"scopes": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"subjects": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"token_endpoint": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"discovery_poll_interval": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"dynamic_metadata_from_claims": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"end_session_properties": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"method_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"access_token_header": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"id_token_header": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"use_bearer_schema_for_authorization": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"identity_token": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"claims_to_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"append": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"claim": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

												"issuer_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"jwks_cache_refresh_policy": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"always": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"max_idp_req_per_polling_interval": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"never": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
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

												"logout_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"parse_callback_path_as_regex": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"scopes": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"session": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cipher_config": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

														"cookie": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"allow_refreshing": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"key_prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"target_domain": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"cookie_options": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"domain": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"http_only": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_age": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"not_secure": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"same_site": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"fail_on_fetch_failure": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"redis": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"allow_refreshing": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"cookie_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"header_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"key_prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"options": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"db": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"host": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"pool_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"socket_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tls_cert_mount_path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"pre_expiry_buffer": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"target_domain": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"session_id_header_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"token_endpoint_query_params": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"opa_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"modules": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

										"options": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fast_input_conversion": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"return_decision_reason": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"query": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"opa_server_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"options": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fast_input_conversion": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"return_decision_reason": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"package": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"rule_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"server_addr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"pass_through_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"failure_mode_allow": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"grpc": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"connection_timeout": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tls_config": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"http": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"connection_timeout": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
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
															Optional:            false,
															Computed:            true,
														},

														"headers_to_add": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pass_through_body": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pass_through_filter_metadata": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pass_through_state": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"response": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"allowed_client_headers_on_denied": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"allowed_upstream_headers": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"allowed_upstream_headers_to_overwrite": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_state_from_response": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"plugin_auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"exported_symbol_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"plugin_file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"fail_on_redirect": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
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

func (r *EnterpriseGlooSoloIoAuthConfigV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *EnterpriseGlooSoloIoAuthConfigV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_enterprise_gloo_solo_io_auth_config_v1")

	var data EnterpriseGlooSoloIoAuthConfigV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "enterprise.gloo.solo.io", Version: "v1", Resource: "AuthConfig"}).
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

	var readResponse EnterpriseGlooSoloIoAuthConfigV1DataSourceData
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
	data.ApiVersion = pointer.String("enterprise.gloo.solo.io/v1")
	data.Kind = pointer.String("AuthConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
