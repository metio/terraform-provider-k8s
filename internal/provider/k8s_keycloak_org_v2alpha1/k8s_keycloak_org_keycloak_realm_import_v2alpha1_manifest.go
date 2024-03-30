/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_keycloak_org_v2alpha1

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
	_ datasource.DataSource = &K8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest{}
)

func NewK8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest() datasource.DataSource {
	return &K8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest{}
}

type K8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest struct{}

type K8SKeycloakOrgKeycloakRealmImportV2Alpha1ManifestData struct {
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
		KeycloakCRName *string `tfsdk:"keycloak_cr_name" json:"keycloakCRName,omitempty"`
		Realm          *struct {
			AccessCodeLifespan                  *int64             `tfsdk:"access_code_lifespan" json:"accessCodeLifespan,omitempty"`
			AccessCodeLifespanLogin             *int64             `tfsdk:"access_code_lifespan_login" json:"accessCodeLifespanLogin,omitempty"`
			AccessCodeLifespanUserAction        *int64             `tfsdk:"access_code_lifespan_user_action" json:"accessCodeLifespanUserAction,omitempty"`
			AccessTokenLifespan                 *int64             `tfsdk:"access_token_lifespan" json:"accessTokenLifespan,omitempty"`
			AccessTokenLifespanForImplicitFlow  *int64             `tfsdk:"access_token_lifespan_for_implicit_flow" json:"accessTokenLifespanForImplicitFlow,omitempty"`
			AccountTheme                        *string            `tfsdk:"account_theme" json:"accountTheme,omitempty"`
			ActionTokenGeneratedByAdminLifespan *int64             `tfsdk:"action_token_generated_by_admin_lifespan" json:"actionTokenGeneratedByAdminLifespan,omitempty"`
			ActionTokenGeneratedByUserLifespan  *int64             `tfsdk:"action_token_generated_by_user_lifespan" json:"actionTokenGeneratedByUserLifespan,omitempty"`
			AdminEventsDetailsEnabled           *bool              `tfsdk:"admin_events_details_enabled" json:"adminEventsDetailsEnabled,omitempty"`
			AdminEventsEnabled                  *bool              `tfsdk:"admin_events_enabled" json:"adminEventsEnabled,omitempty"`
			AdminTheme                          *string            `tfsdk:"admin_theme" json:"adminTheme,omitempty"`
			ApplicationScopeMappings            *map[string]string `tfsdk:"application_scope_mappings" json:"applicationScopeMappings,omitempty"`
			Applications                        *[]struct {
				Access                             *map[string]string `tfsdk:"access" json:"access,omitempty"`
				AdminUrl                           *string            `tfsdk:"admin_url" json:"adminUrl,omitempty"`
				AlwaysDisplayInConsole             *bool              `tfsdk:"always_display_in_console" json:"alwaysDisplayInConsole,omitempty"`
				Attributes                         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				AuthenticationFlowBindingOverrides *map[string]string `tfsdk:"authentication_flow_binding_overrides" json:"authenticationFlowBindingOverrides,omitempty"`
				AuthorizationServicesEnabled       *bool              `tfsdk:"authorization_services_enabled" json:"authorizationServicesEnabled,omitempty"`
				AuthorizationSettings              *struct {
					AllowRemoteResourceManagement *bool   `tfsdk:"allow_remote_resource_management" json:"allowRemoteResourceManagement,omitempty"`
					ClientId                      *string `tfsdk:"client_id" json:"clientId,omitempty"`
					DecisionStrategy              *string `tfsdk:"decision_strategy" json:"decisionStrategy,omitempty"`
					Id                            *string `tfsdk:"id" json:"id,omitempty"`
					Name                          *string `tfsdk:"name" json:"name,omitempty"`
					Policies                      *[]struct {
						Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
						DecisionStrategy *string            `tfsdk:"decision_strategy" json:"decisionStrategy,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						Logic            *string            `tfsdk:"logic" json:"logic,omitempty"`
						Name             *string            `tfsdk:"name" json:"name,omitempty"`
						Owner            *string            `tfsdk:"owner" json:"owner,omitempty"`
						Policies         *[]string          `tfsdk:"policies" json:"policies,omitempty"`
						Resources        *[]string          `tfsdk:"resources" json:"resources,omitempty"`
						ResourcesData    *[]struct {
							_id         *string              `tfsdk:"_id" json:"_id,omitempty"`
							Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
							DisplayName *string              `tfsdk:"display_name" json:"displayName,omitempty"`
							Icon_uri    *string              `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
							Name        *string              `tfsdk:"name" json:"name,omitempty"`
							Owner       *struct {
								Id   *string `tfsdk:"id" json:"id,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"owner" json:"owner,omitempty"`
							OwnerManagedAccess *bool `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
							Scopes             *[]struct {
								DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
								IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
								Id          *string `tfsdk:"id" json:"id,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"scopes" json:"scopes,omitempty"`
							Type *string   `tfsdk:"type" json:"type,omitempty"`
							Uris *[]string `tfsdk:"uris" json:"uris,omitempty"`
						} `tfsdk:"resources_data" json:"resourcesData,omitempty"`
						Scopes     *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
						ScopesData *[]struct {
							DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
							IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"scopes_data" json:"scopesData,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
					PolicyEnforcementMode *string `tfsdk:"policy_enforcement_mode" json:"policyEnforcementMode,omitempty"`
					Resources             *[]struct {
						_id         *string              `tfsdk:"_id" json:"_id,omitempty"`
						Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
						DisplayName *string              `tfsdk:"display_name" json:"displayName,omitempty"`
						Icon_uri    *string              `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
						Name        *string              `tfsdk:"name" json:"name,omitempty"`
						Owner       *struct {
							Id   *string `tfsdk:"id" json:"id,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"owner" json:"owner,omitempty"`
						OwnerManagedAccess *bool `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
						Scopes             *[]struct {
							DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
							IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"scopes" json:"scopes,omitempty"`
						Type *string   `tfsdk:"type" json:"type,omitempty"`
						Uris *[]string `tfsdk:"uris" json:"uris,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Scopes *[]struct {
						DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
						IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
						Id          *string `tfsdk:"id" json:"id,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"scopes" json:"scopes,omitempty"`
				} `tfsdk:"authorization_settings" json:"authorizationSettings,omitempty"`
				BaseUrl    *string `tfsdk:"base_url" json:"baseUrl,omitempty"`
				BearerOnly *bool   `tfsdk:"bearer_only" json:"bearerOnly,omitempty"`
				Claims     *struct {
					Address  *bool `tfsdk:"address" json:"address,omitempty"`
					Email    *bool `tfsdk:"email" json:"email,omitempty"`
					Gender   *bool `tfsdk:"gender" json:"gender,omitempty"`
					Locale   *bool `tfsdk:"locale" json:"locale,omitempty"`
					Name     *bool `tfsdk:"name" json:"name,omitempty"`
					Phone    *bool `tfsdk:"phone" json:"phone,omitempty"`
					Picture  *bool `tfsdk:"picture" json:"picture,omitempty"`
					Profile  *bool `tfsdk:"profile" json:"profile,omitempty"`
					Username *bool `tfsdk:"username" json:"username,omitempty"`
					Website  *bool `tfsdk:"website" json:"website,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				ClientAuthenticatorType               *string   `tfsdk:"client_authenticator_type" json:"clientAuthenticatorType,omitempty"`
				ClientId                              *string   `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientTemplate                        *string   `tfsdk:"client_template" json:"clientTemplate,omitempty"`
				ConsentRequired                       *bool     `tfsdk:"consent_required" json:"consentRequired,omitempty"`
				DefaultClientScopes                   *[]string `tfsdk:"default_client_scopes" json:"defaultClientScopes,omitempty"`
				DefaultRoles                          *[]string `tfsdk:"default_roles" json:"defaultRoles,omitempty"`
				Description                           *string   `tfsdk:"description" json:"description,omitempty"`
				DirectAccessGrantsEnabled             *bool     `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
				DirectGrantsOnly                      *bool     `tfsdk:"direct_grants_only" json:"directGrantsOnly,omitempty"`
				Enabled                               *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				FrontchannelLogout                    *bool     `tfsdk:"frontchannel_logout" json:"frontchannelLogout,omitempty"`
				FullScopeAllowed                      *bool     `tfsdk:"full_scope_allowed" json:"fullScopeAllowed,omitempty"`
				Id                                    *string   `tfsdk:"id" json:"id,omitempty"`
				ImplicitFlowEnabled                   *bool     `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
				Name                                  *string   `tfsdk:"name" json:"name,omitempty"`
				NodeReRegistrationTimeout             *int64    `tfsdk:"node_re_registration_timeout" json:"nodeReRegistrationTimeout,omitempty"`
				NotBefore                             *int64    `tfsdk:"not_before" json:"notBefore,omitempty"`
				Oauth2DeviceAuthorizationGrantEnabled *bool     `tfsdk:"oauth2_device_authorization_grant_enabled" json:"oauth2DeviceAuthorizationGrantEnabled,omitempty"`
				OptionalClientScopes                  *[]string `tfsdk:"optional_client_scopes" json:"optionalClientScopes,omitempty"`
				Origin                                *string   `tfsdk:"origin" json:"origin,omitempty"`
				Protocol                              *string   `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMappers                       *[]struct {
					Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
					ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
					ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
					Id              *string            `tfsdk:"id" json:"id,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
				PublicClient            *bool              `tfsdk:"public_client" json:"publicClient,omitempty"`
				RedirectUris            *[]string          `tfsdk:"redirect_uris" json:"redirectUris,omitempty"`
				RegisteredNodes         *map[string]string `tfsdk:"registered_nodes" json:"registeredNodes,omitempty"`
				RegistrationAccessToken *string            `tfsdk:"registration_access_token" json:"registrationAccessToken,omitempty"`
				RootUrl                 *string            `tfsdk:"root_url" json:"rootUrl,omitempty"`
				Secret                  *string            `tfsdk:"secret" json:"secret,omitempty"`
				ServiceAccountsEnabled  *bool              `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
				StandardFlowEnabled     *bool              `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
				SurrogateAuthRequired   *bool              `tfsdk:"surrogate_auth_required" json:"surrogateAuthRequired,omitempty"`
				UseTemplateConfig       *bool              `tfsdk:"use_template_config" json:"useTemplateConfig,omitempty"`
				UseTemplateMappers      *bool              `tfsdk:"use_template_mappers" json:"useTemplateMappers,omitempty"`
				UseTemplateScope        *bool              `tfsdk:"use_template_scope" json:"useTemplateScope,omitempty"`
				WebOrigins              *[]string          `tfsdk:"web_origins" json:"webOrigins,omitempty"`
			} `tfsdk:"applications" json:"applications,omitempty"`
			Attributes          *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
			AuthenticationFlows *[]struct {
				Alias                    *string `tfsdk:"alias" json:"alias,omitempty"`
				AuthenticationExecutions *[]struct {
					Authenticator       *string `tfsdk:"authenticator" json:"authenticator,omitempty"`
					AuthenticatorConfig *string `tfsdk:"authenticator_config" json:"authenticatorConfig,omitempty"`
					AuthenticatorFlow   *bool   `tfsdk:"authenticator_flow" json:"authenticatorFlow,omitempty"`
					AutheticatorFlow    *bool   `tfsdk:"autheticator_flow" json:"autheticatorFlow,omitempty"`
					FlowAlias           *string `tfsdk:"flow_alias" json:"flowAlias,omitempty"`
					Priority            *int64  `tfsdk:"priority" json:"priority,omitempty"`
					Requirement         *string `tfsdk:"requirement" json:"requirement,omitempty"`
					UserSetupAllowed    *bool   `tfsdk:"user_setup_allowed" json:"userSetupAllowed,omitempty"`
				} `tfsdk:"authentication_executions" json:"authenticationExecutions,omitempty"`
				BuiltIn     *bool   `tfsdk:"built_in" json:"builtIn,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Id          *string `tfsdk:"id" json:"id,omitempty"`
				ProviderId  *string `tfsdk:"provider_id" json:"providerId,omitempty"`
				TopLevel    *bool   `tfsdk:"top_level" json:"topLevel,omitempty"`
			} `tfsdk:"authentication_flows" json:"authenticationFlows,omitempty"`
			AuthenticatorConfig *[]struct {
				Alias  *string            `tfsdk:"alias" json:"alias,omitempty"`
				Config *map[string]string `tfsdk:"config" json:"config,omitempty"`
				Id     *string            `tfsdk:"id" json:"id,omitempty"`
			} `tfsdk:"authenticator_config" json:"authenticatorConfig,omitempty"`
			BrowserFlow                     *string            `tfsdk:"browser_flow" json:"browserFlow,omitempty"`
			BrowserSecurityHeaders          *map[string]string `tfsdk:"browser_security_headers" json:"browserSecurityHeaders,omitempty"`
			BruteForceProtected             *bool              `tfsdk:"brute_force_protected" json:"bruteForceProtected,omitempty"`
			Certificate                     *string            `tfsdk:"certificate" json:"certificate,omitempty"`
			ClientAuthenticationFlow        *string            `tfsdk:"client_authentication_flow" json:"clientAuthenticationFlow,omitempty"`
			ClientOfflineSessionIdleTimeout *int64             `tfsdk:"client_offline_session_idle_timeout" json:"clientOfflineSessionIdleTimeout,omitempty"`
			ClientOfflineSessionMaxLifespan *int64             `tfsdk:"client_offline_session_max_lifespan" json:"clientOfflineSessionMaxLifespan,omitempty"`
			ClientPolicies                  *map[string]string `tfsdk:"client_policies" json:"clientPolicies,omitempty"`
			ClientProfiles                  *map[string]string `tfsdk:"client_profiles" json:"clientProfiles,omitempty"`
			ClientScopeMappings             *map[string]string `tfsdk:"client_scope_mappings" json:"clientScopeMappings,omitempty"`
			ClientScopes                    *[]struct {
				Attributes      *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				Description     *string            `tfsdk:"description" json:"description,omitempty"`
				Id              *string            `tfsdk:"id" json:"id,omitempty"`
				Name            *string            `tfsdk:"name" json:"name,omitempty"`
				Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMappers *[]struct {
					Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
					ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
					ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
					Id              *string            `tfsdk:"id" json:"id,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
			} `tfsdk:"client_scopes" json:"clientScopes,omitempty"`
			ClientSessionIdleTimeout *int64 `tfsdk:"client_session_idle_timeout" json:"clientSessionIdleTimeout,omitempty"`
			ClientSessionMaxLifespan *int64 `tfsdk:"client_session_max_lifespan" json:"clientSessionMaxLifespan,omitempty"`
			ClientTemplates          *[]struct {
				Attributes                *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				BearerOnly                *bool              `tfsdk:"bearer_only" json:"bearerOnly,omitempty"`
				ConsentRequired           *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
				Description               *string            `tfsdk:"description" json:"description,omitempty"`
				DirectAccessGrantsEnabled *bool              `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
				FrontchannelLogout        *bool              `tfsdk:"frontchannel_logout" json:"frontchannelLogout,omitempty"`
				FullScopeAllowed          *bool              `tfsdk:"full_scope_allowed" json:"fullScopeAllowed,omitempty"`
				Id                        *string            `tfsdk:"id" json:"id,omitempty"`
				ImplicitFlowEnabled       *bool              `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
				Name                      *string            `tfsdk:"name" json:"name,omitempty"`
				Protocol                  *string            `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMappers           *[]struct {
					Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
					ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
					ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
					Id              *string            `tfsdk:"id" json:"id,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
				PublicClient           *bool `tfsdk:"public_client" json:"publicClient,omitempty"`
				ServiceAccountsEnabled *bool `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
				StandardFlowEnabled    *bool `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
			} `tfsdk:"client_templates" json:"clientTemplates,omitempty"`
			Clients *[]struct {
				Access                             *map[string]string `tfsdk:"access" json:"access,omitempty"`
				AdminUrl                           *string            `tfsdk:"admin_url" json:"adminUrl,omitempty"`
				AlwaysDisplayInConsole             *bool              `tfsdk:"always_display_in_console" json:"alwaysDisplayInConsole,omitempty"`
				Attributes                         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				AuthenticationFlowBindingOverrides *map[string]string `tfsdk:"authentication_flow_binding_overrides" json:"authenticationFlowBindingOverrides,omitempty"`
				AuthorizationServicesEnabled       *bool              `tfsdk:"authorization_services_enabled" json:"authorizationServicesEnabled,omitempty"`
				AuthorizationSettings              *struct {
					AllowRemoteResourceManagement *bool   `tfsdk:"allow_remote_resource_management" json:"allowRemoteResourceManagement,omitempty"`
					ClientId                      *string `tfsdk:"client_id" json:"clientId,omitempty"`
					DecisionStrategy              *string `tfsdk:"decision_strategy" json:"decisionStrategy,omitempty"`
					Id                            *string `tfsdk:"id" json:"id,omitempty"`
					Name                          *string `tfsdk:"name" json:"name,omitempty"`
					Policies                      *[]struct {
						Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
						DecisionStrategy *string            `tfsdk:"decision_strategy" json:"decisionStrategy,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						Logic            *string            `tfsdk:"logic" json:"logic,omitempty"`
						Name             *string            `tfsdk:"name" json:"name,omitempty"`
						Owner            *string            `tfsdk:"owner" json:"owner,omitempty"`
						Policies         *[]string          `tfsdk:"policies" json:"policies,omitempty"`
						Resources        *[]string          `tfsdk:"resources" json:"resources,omitempty"`
						ResourcesData    *[]struct {
							_id         *string              `tfsdk:"_id" json:"_id,omitempty"`
							Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
							DisplayName *string              `tfsdk:"display_name" json:"displayName,omitempty"`
							Icon_uri    *string              `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
							Name        *string              `tfsdk:"name" json:"name,omitempty"`
							Owner       *struct {
								Id   *string `tfsdk:"id" json:"id,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"owner" json:"owner,omitempty"`
							OwnerManagedAccess *bool `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
							Scopes             *[]struct {
								DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
								IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
								Id          *string `tfsdk:"id" json:"id,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"scopes" json:"scopes,omitempty"`
							Type *string   `tfsdk:"type" json:"type,omitempty"`
							Uris *[]string `tfsdk:"uris" json:"uris,omitempty"`
						} `tfsdk:"resources_data" json:"resourcesData,omitempty"`
						Scopes     *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
						ScopesData *[]struct {
							DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
							IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"scopes_data" json:"scopesData,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
					PolicyEnforcementMode *string `tfsdk:"policy_enforcement_mode" json:"policyEnforcementMode,omitempty"`
					Resources             *[]struct {
						_id         *string              `tfsdk:"_id" json:"_id,omitempty"`
						Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
						DisplayName *string              `tfsdk:"display_name" json:"displayName,omitempty"`
						Icon_uri    *string              `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
						Name        *string              `tfsdk:"name" json:"name,omitempty"`
						Owner       *struct {
							Id   *string `tfsdk:"id" json:"id,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"owner" json:"owner,omitempty"`
						OwnerManagedAccess *bool `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
						Scopes             *[]struct {
							DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
							IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"scopes" json:"scopes,omitempty"`
						Type *string   `tfsdk:"type" json:"type,omitempty"`
						Uris *[]string `tfsdk:"uris" json:"uris,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Scopes *[]struct {
						DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
						IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
						Id          *string `tfsdk:"id" json:"id,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"scopes" json:"scopes,omitempty"`
				} `tfsdk:"authorization_settings" json:"authorizationSettings,omitempty"`
				BaseUrl                               *string   `tfsdk:"base_url" json:"baseUrl,omitempty"`
				BearerOnly                            *bool     `tfsdk:"bearer_only" json:"bearerOnly,omitempty"`
				ClientAuthenticatorType               *string   `tfsdk:"client_authenticator_type" json:"clientAuthenticatorType,omitempty"`
				ClientId                              *string   `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientTemplate                        *string   `tfsdk:"client_template" json:"clientTemplate,omitempty"`
				ConsentRequired                       *bool     `tfsdk:"consent_required" json:"consentRequired,omitempty"`
				DefaultClientScopes                   *[]string `tfsdk:"default_client_scopes" json:"defaultClientScopes,omitempty"`
				DefaultRoles                          *[]string `tfsdk:"default_roles" json:"defaultRoles,omitempty"`
				Description                           *string   `tfsdk:"description" json:"description,omitempty"`
				DirectAccessGrantsEnabled             *bool     `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
				DirectGrantsOnly                      *bool     `tfsdk:"direct_grants_only" json:"directGrantsOnly,omitempty"`
				Enabled                               *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				FrontchannelLogout                    *bool     `tfsdk:"frontchannel_logout" json:"frontchannelLogout,omitempty"`
				FullScopeAllowed                      *bool     `tfsdk:"full_scope_allowed" json:"fullScopeAllowed,omitempty"`
				Id                                    *string   `tfsdk:"id" json:"id,omitempty"`
				ImplicitFlowEnabled                   *bool     `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
				Name                                  *string   `tfsdk:"name" json:"name,omitempty"`
				NodeReRegistrationTimeout             *int64    `tfsdk:"node_re_registration_timeout" json:"nodeReRegistrationTimeout,omitempty"`
				NotBefore                             *int64    `tfsdk:"not_before" json:"notBefore,omitempty"`
				Oauth2DeviceAuthorizationGrantEnabled *bool     `tfsdk:"oauth2_device_authorization_grant_enabled" json:"oauth2DeviceAuthorizationGrantEnabled,omitempty"`
				OptionalClientScopes                  *[]string `tfsdk:"optional_client_scopes" json:"optionalClientScopes,omitempty"`
				Origin                                *string   `tfsdk:"origin" json:"origin,omitempty"`
				Protocol                              *string   `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMappers                       *[]struct {
					Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
					ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
					ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
					Id              *string            `tfsdk:"id" json:"id,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
				PublicClient            *bool              `tfsdk:"public_client" json:"publicClient,omitempty"`
				RedirectUris            *[]string          `tfsdk:"redirect_uris" json:"redirectUris,omitempty"`
				RegisteredNodes         *map[string]string `tfsdk:"registered_nodes" json:"registeredNodes,omitempty"`
				RegistrationAccessToken *string            `tfsdk:"registration_access_token" json:"registrationAccessToken,omitempty"`
				RootUrl                 *string            `tfsdk:"root_url" json:"rootUrl,omitempty"`
				Secret                  *string            `tfsdk:"secret" json:"secret,omitempty"`
				ServiceAccountsEnabled  *bool              `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
				StandardFlowEnabled     *bool              `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
				SurrogateAuthRequired   *bool              `tfsdk:"surrogate_auth_required" json:"surrogateAuthRequired,omitempty"`
				UseTemplateConfig       *bool              `tfsdk:"use_template_config" json:"useTemplateConfig,omitempty"`
				UseTemplateMappers      *bool              `tfsdk:"use_template_mappers" json:"useTemplateMappers,omitempty"`
				UseTemplateScope        *bool              `tfsdk:"use_template_scope" json:"useTemplateScope,omitempty"`
				WebOrigins              *[]string          `tfsdk:"web_origins" json:"webOrigins,omitempty"`
			} `tfsdk:"clients" json:"clients,omitempty"`
			CodeSecret                  *string            `tfsdk:"code_secret" json:"codeSecret,omitempty"`
			Components                  *map[string]string `tfsdk:"components" json:"components,omitempty"`
			DefaultDefaultClientScopes  *[]string          `tfsdk:"default_default_client_scopes" json:"defaultDefaultClientScopes,omitempty"`
			DefaultGroups               *[]string          `tfsdk:"default_groups" json:"defaultGroups,omitempty"`
			DefaultLocale               *string            `tfsdk:"default_locale" json:"defaultLocale,omitempty"`
			DefaultOptionalClientScopes *[]string          `tfsdk:"default_optional_client_scopes" json:"defaultOptionalClientScopes,omitempty"`
			DefaultRole                 *struct {
				Attributes *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClientRole *bool                `tfsdk:"client_role" json:"clientRole,omitempty"`
				Composite  *bool                `tfsdk:"composite" json:"composite,omitempty"`
				Composites *struct {
					Application *map[string][]string `tfsdk:"application" json:"application,omitempty"`
					Client      *map[string][]string `tfsdk:"client" json:"client,omitempty"`
					Realm       *[]string            `tfsdk:"realm" json:"realm,omitempty"`
				} `tfsdk:"composites" json:"composites,omitempty"`
				ContainerId        *string `tfsdk:"container_id" json:"containerId,omitempty"`
				Description        *string `tfsdk:"description" json:"description,omitempty"`
				Id                 *string `tfsdk:"id" json:"id,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				ScopeParamRequired *bool   `tfsdk:"scope_param_required" json:"scopeParamRequired,omitempty"`
			} `tfsdk:"default_role" json:"defaultRole,omitempty"`
			DefaultRoles              *[]string `tfsdk:"default_roles" json:"defaultRoles,omitempty"`
			DefaultSignatureAlgorithm *string   `tfsdk:"default_signature_algorithm" json:"defaultSignatureAlgorithm,omitempty"`
			DirectGrantFlow           *string   `tfsdk:"direct_grant_flow" json:"directGrantFlow,omitempty"`
			DisplayName               *string   `tfsdk:"display_name" json:"displayName,omitempty"`
			DisplayNameHtml           *string   `tfsdk:"display_name_html" json:"displayNameHtml,omitempty"`
			DockerAuthenticationFlow  *string   `tfsdk:"docker_authentication_flow" json:"dockerAuthenticationFlow,omitempty"`
			DuplicateEmailsAllowed    *bool     `tfsdk:"duplicate_emails_allowed" json:"duplicateEmailsAllowed,omitempty"`
			EditUsernameAllowed       *bool     `tfsdk:"edit_username_allowed" json:"editUsernameAllowed,omitempty"`
			EmailTheme                *string   `tfsdk:"email_theme" json:"emailTheme,omitempty"`
			Enabled                   *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
			EnabledEventTypes         *[]string `tfsdk:"enabled_event_types" json:"enabledEventTypes,omitempty"`
			EventsEnabled             *bool     `tfsdk:"events_enabled" json:"eventsEnabled,omitempty"`
			EventsExpiration          *int64    `tfsdk:"events_expiration" json:"eventsExpiration,omitempty"`
			EventsListeners           *[]string `tfsdk:"events_listeners" json:"eventsListeners,omitempty"`
			FailureFactor             *int64    `tfsdk:"failure_factor" json:"failureFactor,omitempty"`
			FederatedUsers            *[]struct {
				Access           *map[string]string   `tfsdk:"access" json:"access,omitempty"`
				ApplicationRoles *map[string][]string `tfsdk:"application_roles" json:"applicationRoles,omitempty"`
				Attributes       *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClientConsents   *[]struct {
					ClientId            *string   `tfsdk:"client_id" json:"clientId,omitempty"`
					CreatedDate         *int64    `tfsdk:"created_date" json:"createdDate,omitempty"`
					GrantedClientScopes *[]string `tfsdk:"granted_client_scopes" json:"grantedClientScopes,omitempty"`
					GrantedRealmRoles   *[]string `tfsdk:"granted_realm_roles" json:"grantedRealmRoles,omitempty"`
					LastUpdatedDate     *int64    `tfsdk:"last_updated_date" json:"lastUpdatedDate,omitempty"`
				} `tfsdk:"client_consents" json:"clientConsents,omitempty"`
				ClientRoles      *map[string][]string `tfsdk:"client_roles" json:"clientRoles,omitempty"`
				CreatedTimestamp *int64               `tfsdk:"created_timestamp" json:"createdTimestamp,omitempty"`
				Credentials      *[]struct {
					Algorithm         *string              `tfsdk:"algorithm" json:"algorithm,omitempty"`
					Config            *map[string][]string `tfsdk:"config" json:"config,omitempty"`
					Counter           *int64               `tfsdk:"counter" json:"counter,omitempty"`
					CreatedDate       *int64               `tfsdk:"created_date" json:"createdDate,omitempty"`
					CredentialData    *string              `tfsdk:"credential_data" json:"credentialData,omitempty"`
					Device            *string              `tfsdk:"device" json:"device,omitempty"`
					Digits            *int64               `tfsdk:"digits" json:"digits,omitempty"`
					HashIterations    *int64               `tfsdk:"hash_iterations" json:"hashIterations,omitempty"`
					HashedSaltedValue *string              `tfsdk:"hashed_salted_value" json:"hashedSaltedValue,omitempty"`
					Id                *string              `tfsdk:"id" json:"id,omitempty"`
					Period            *int64               `tfsdk:"period" json:"period,omitempty"`
					Priority          *int64               `tfsdk:"priority" json:"priority,omitempty"`
					Salt              *string              `tfsdk:"salt" json:"salt,omitempty"`
					SecretData        *string              `tfsdk:"secret_data" json:"secretData,omitempty"`
					Temporary         *bool                `tfsdk:"temporary" json:"temporary,omitempty"`
					Type              *string              `tfsdk:"type" json:"type,omitempty"`
					UserLabel         *string              `tfsdk:"user_label" json:"userLabel,omitempty"`
					Value             *string              `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				DisableableCredentialTypes *[]string `tfsdk:"disableable_credential_types" json:"disableableCredentialTypes,omitempty"`
				Email                      *string   `tfsdk:"email" json:"email,omitempty"`
				EmailVerified              *bool     `tfsdk:"email_verified" json:"emailVerified,omitempty"`
				Enabled                    *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				FederatedIdentities        *[]struct {
					IdentityProvider *string `tfsdk:"identity_provider" json:"identityProvider,omitempty"`
					UserId           *string `tfsdk:"user_id" json:"userId,omitempty"`
					UserName         *string `tfsdk:"user_name" json:"userName,omitempty"`
				} `tfsdk:"federated_identities" json:"federatedIdentities,omitempty"`
				FederationLink         *string   `tfsdk:"federation_link" json:"federationLink,omitempty"`
				FirstName              *string   `tfsdk:"first_name" json:"firstName,omitempty"`
				Groups                 *[]string `tfsdk:"groups" json:"groups,omitempty"`
				Id                     *string   `tfsdk:"id" json:"id,omitempty"`
				LastName               *string   `tfsdk:"last_name" json:"lastName,omitempty"`
				NotBefore              *int64    `tfsdk:"not_before" json:"notBefore,omitempty"`
				Origin                 *string   `tfsdk:"origin" json:"origin,omitempty"`
				RealmRoles             *[]string `tfsdk:"realm_roles" json:"realmRoles,omitempty"`
				RequiredActions        *[]string `tfsdk:"required_actions" json:"requiredActions,omitempty"`
				Self                   *string   `tfsdk:"self" json:"self,omitempty"`
				ServiceAccountClientId *string   `tfsdk:"service_account_client_id" json:"serviceAccountClientId,omitempty"`
				SocialLinks            *[]struct {
					SocialProvider *string `tfsdk:"social_provider" json:"socialProvider,omitempty"`
					SocialUserId   *string `tfsdk:"social_user_id" json:"socialUserId,omitempty"`
					SocialUsername *string `tfsdk:"social_username" json:"socialUsername,omitempty"`
				} `tfsdk:"social_links" json:"socialLinks,omitempty"`
				Totp     *bool   `tfsdk:"totp" json:"totp,omitempty"`
				Username *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"federated_users" json:"federatedUsers,omitempty"`
			Groups *[]struct {
				Access      *map[string]string   `tfsdk:"access" json:"access,omitempty"`
				Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClientRoles *map[string][]string `tfsdk:"client_roles" json:"clientRoles,omitempty"`
				Id          *string              `tfsdk:"id" json:"id,omitempty"`
				Name        *string              `tfsdk:"name" json:"name,omitempty"`
				Path        *string              `tfsdk:"path" json:"path,omitempty"`
				RealmRoles  *[]string            `tfsdk:"realm_roles" json:"realmRoles,omitempty"`
				SubGroups   *[]struct {
					Access      *map[string]string   `tfsdk:"access" json:"access,omitempty"`
					Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
					ClientRoles *map[string][]string `tfsdk:"client_roles" json:"clientRoles,omitempty"`
					Id          *string              `tfsdk:"id" json:"id,omitempty"`
					Name        *string              `tfsdk:"name" json:"name,omitempty"`
					Path        *string              `tfsdk:"path" json:"path,omitempty"`
					RealmRoles  *[]string            `tfsdk:"realm_roles" json:"realmRoles,omitempty"`
				} `tfsdk:"sub_groups" json:"subGroups,omitempty"`
			} `tfsdk:"groups" json:"groups,omitempty"`
			Id                      *string `tfsdk:"id" json:"id,omitempty"`
			IdentityProviderMappers *[]struct {
				Config                 *map[string]string `tfsdk:"config" json:"config,omitempty"`
				Id                     *string            `tfsdk:"id" json:"id,omitempty"`
				IdentityProviderAlias  *string            `tfsdk:"identity_provider_alias" json:"identityProviderAlias,omitempty"`
				IdentityProviderMapper *string            `tfsdk:"identity_provider_mapper" json:"identityProviderMapper,omitempty"`
				Name                   *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"identity_provider_mappers" json:"identityProviderMappers,omitempty"`
			IdentityProviders *[]struct {
				AddReadTokenRoleOnCreate    *bool              `tfsdk:"add_read_token_role_on_create" json:"addReadTokenRoleOnCreate,omitempty"`
				Alias                       *string            `tfsdk:"alias" json:"alias,omitempty"`
				AuthenticateByDefault       *bool              `tfsdk:"authenticate_by_default" json:"authenticateByDefault,omitempty"`
				Config                      *map[string]string `tfsdk:"config" json:"config,omitempty"`
				DisplayName                 *string            `tfsdk:"display_name" json:"displayName,omitempty"`
				Enabled                     *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				FirstBrokerLoginFlowAlias   *string            `tfsdk:"first_broker_login_flow_alias" json:"firstBrokerLoginFlowAlias,omitempty"`
				InternalId                  *string            `tfsdk:"internal_id" json:"internalId,omitempty"`
				LinkOnly                    *bool              `tfsdk:"link_only" json:"linkOnly,omitempty"`
				PostBrokerLoginFlowAlias    *string            `tfsdk:"post_broker_login_flow_alias" json:"postBrokerLoginFlowAlias,omitempty"`
				ProviderId                  *string            `tfsdk:"provider_id" json:"providerId,omitempty"`
				StoreToken                  *bool              `tfsdk:"store_token" json:"storeToken,omitempty"`
				TrustEmail                  *bool              `tfsdk:"trust_email" json:"trustEmail,omitempty"`
				UpdateProfileFirstLoginMode *string            `tfsdk:"update_profile_first_login_mode" json:"updateProfileFirstLoginMode,omitempty"`
			} `tfsdk:"identity_providers" json:"identityProviders,omitempty"`
			InternationalizationEnabled  *bool   `tfsdk:"internationalization_enabled" json:"internationalizationEnabled,omitempty"`
			KeycloakVersion              *string `tfsdk:"keycloak_version" json:"keycloakVersion,omitempty"`
			LoginTheme                   *string `tfsdk:"login_theme" json:"loginTheme,omitempty"`
			LoginWithEmailAllowed        *bool   `tfsdk:"login_with_email_allowed" json:"loginWithEmailAllowed,omitempty"`
			MaxDeltaTimeSeconds          *int64  `tfsdk:"max_delta_time_seconds" json:"maxDeltaTimeSeconds,omitempty"`
			MaxFailureWaitSeconds        *int64  `tfsdk:"max_failure_wait_seconds" json:"maxFailureWaitSeconds,omitempty"`
			MinimumQuickLoginWaitSeconds *int64  `tfsdk:"minimum_quick_login_wait_seconds" json:"minimumQuickLoginWaitSeconds,omitempty"`
			NotBefore                    *int64  `tfsdk:"not_before" json:"notBefore,omitempty"`
			Oauth2DeviceCodeLifespan     *int64  `tfsdk:"oauth2_device_code_lifespan" json:"oauth2DeviceCodeLifespan,omitempty"`
			Oauth2DevicePollingInterval  *int64  `tfsdk:"oauth2_device_polling_interval" json:"oauth2DevicePollingInterval,omitempty"`
			OauthClients                 *[]struct {
				Access                             *map[string]string `tfsdk:"access" json:"access,omitempty"`
				AdminUrl                           *string            `tfsdk:"admin_url" json:"adminUrl,omitempty"`
				AlwaysDisplayInConsole             *bool              `tfsdk:"always_display_in_console" json:"alwaysDisplayInConsole,omitempty"`
				Attributes                         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				AuthenticationFlowBindingOverrides *map[string]string `tfsdk:"authentication_flow_binding_overrides" json:"authenticationFlowBindingOverrides,omitempty"`
				AuthorizationServicesEnabled       *bool              `tfsdk:"authorization_services_enabled" json:"authorizationServicesEnabled,omitempty"`
				AuthorizationSettings              *struct {
					AllowRemoteResourceManagement *bool   `tfsdk:"allow_remote_resource_management" json:"allowRemoteResourceManagement,omitempty"`
					ClientId                      *string `tfsdk:"client_id" json:"clientId,omitempty"`
					DecisionStrategy              *string `tfsdk:"decision_strategy" json:"decisionStrategy,omitempty"`
					Id                            *string `tfsdk:"id" json:"id,omitempty"`
					Name                          *string `tfsdk:"name" json:"name,omitempty"`
					Policies                      *[]struct {
						Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
						DecisionStrategy *string            `tfsdk:"decision_strategy" json:"decisionStrategy,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						Logic            *string            `tfsdk:"logic" json:"logic,omitempty"`
						Name             *string            `tfsdk:"name" json:"name,omitempty"`
						Owner            *string            `tfsdk:"owner" json:"owner,omitempty"`
						Policies         *[]string          `tfsdk:"policies" json:"policies,omitempty"`
						Resources        *[]string          `tfsdk:"resources" json:"resources,omitempty"`
						ResourcesData    *[]struct {
							_id         *string              `tfsdk:"_id" json:"_id,omitempty"`
							Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
							DisplayName *string              `tfsdk:"display_name" json:"displayName,omitempty"`
							Icon_uri    *string              `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
							Name        *string              `tfsdk:"name" json:"name,omitempty"`
							Owner       *struct {
								Id   *string `tfsdk:"id" json:"id,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"owner" json:"owner,omitempty"`
							OwnerManagedAccess *bool `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
							Scopes             *[]struct {
								DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
								IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
								Id          *string `tfsdk:"id" json:"id,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"scopes" json:"scopes,omitempty"`
							Type *string   `tfsdk:"type" json:"type,omitempty"`
							Uris *[]string `tfsdk:"uris" json:"uris,omitempty"`
						} `tfsdk:"resources_data" json:"resourcesData,omitempty"`
						Scopes     *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
						ScopesData *[]struct {
							DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
							IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"scopes_data" json:"scopesData,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
					PolicyEnforcementMode *string `tfsdk:"policy_enforcement_mode" json:"policyEnforcementMode,omitempty"`
					Resources             *[]struct {
						_id         *string              `tfsdk:"_id" json:"_id,omitempty"`
						Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
						DisplayName *string              `tfsdk:"display_name" json:"displayName,omitempty"`
						Icon_uri    *string              `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
						Name        *string              `tfsdk:"name" json:"name,omitempty"`
						Owner       *struct {
							Id   *string `tfsdk:"id" json:"id,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"owner" json:"owner,omitempty"`
						OwnerManagedAccess *bool `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
						Scopes             *[]struct {
							DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
							IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"scopes" json:"scopes,omitempty"`
						Type *string   `tfsdk:"type" json:"type,omitempty"`
						Uris *[]string `tfsdk:"uris" json:"uris,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Scopes *[]struct {
						DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
						IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
						Id          *string `tfsdk:"id" json:"id,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"scopes" json:"scopes,omitempty"`
				} `tfsdk:"authorization_settings" json:"authorizationSettings,omitempty"`
				BaseUrl    *string `tfsdk:"base_url" json:"baseUrl,omitempty"`
				BearerOnly *bool   `tfsdk:"bearer_only" json:"bearerOnly,omitempty"`
				Claims     *struct {
					Address  *bool `tfsdk:"address" json:"address,omitempty"`
					Email    *bool `tfsdk:"email" json:"email,omitempty"`
					Gender   *bool `tfsdk:"gender" json:"gender,omitempty"`
					Locale   *bool `tfsdk:"locale" json:"locale,omitempty"`
					Name     *bool `tfsdk:"name" json:"name,omitempty"`
					Phone    *bool `tfsdk:"phone" json:"phone,omitempty"`
					Picture  *bool `tfsdk:"picture" json:"picture,omitempty"`
					Profile  *bool `tfsdk:"profile" json:"profile,omitempty"`
					Username *bool `tfsdk:"username" json:"username,omitempty"`
					Website  *bool `tfsdk:"website" json:"website,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				ClientAuthenticatorType               *string   `tfsdk:"client_authenticator_type" json:"clientAuthenticatorType,omitempty"`
				ClientId                              *string   `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientTemplate                        *string   `tfsdk:"client_template" json:"clientTemplate,omitempty"`
				ConsentRequired                       *bool     `tfsdk:"consent_required" json:"consentRequired,omitempty"`
				DefaultClientScopes                   *[]string `tfsdk:"default_client_scopes" json:"defaultClientScopes,omitempty"`
				DefaultRoles                          *[]string `tfsdk:"default_roles" json:"defaultRoles,omitempty"`
				Description                           *string   `tfsdk:"description" json:"description,omitempty"`
				DirectAccessGrantsEnabled             *bool     `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
				DirectGrantsOnly                      *bool     `tfsdk:"direct_grants_only" json:"directGrantsOnly,omitempty"`
				Enabled                               *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				FrontchannelLogout                    *bool     `tfsdk:"frontchannel_logout" json:"frontchannelLogout,omitempty"`
				FullScopeAllowed                      *bool     `tfsdk:"full_scope_allowed" json:"fullScopeAllowed,omitempty"`
				Id                                    *string   `tfsdk:"id" json:"id,omitempty"`
				ImplicitFlowEnabled                   *bool     `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
				Name                                  *string   `tfsdk:"name" json:"name,omitempty"`
				NodeReRegistrationTimeout             *int64    `tfsdk:"node_re_registration_timeout" json:"nodeReRegistrationTimeout,omitempty"`
				NotBefore                             *int64    `tfsdk:"not_before" json:"notBefore,omitempty"`
				Oauth2DeviceAuthorizationGrantEnabled *bool     `tfsdk:"oauth2_device_authorization_grant_enabled" json:"oauth2DeviceAuthorizationGrantEnabled,omitempty"`
				OptionalClientScopes                  *[]string `tfsdk:"optional_client_scopes" json:"optionalClientScopes,omitempty"`
				Origin                                *string   `tfsdk:"origin" json:"origin,omitempty"`
				Protocol                              *string   `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMappers                       *[]struct {
					Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
					ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
					ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
					Id              *string            `tfsdk:"id" json:"id,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
				PublicClient            *bool              `tfsdk:"public_client" json:"publicClient,omitempty"`
				RedirectUris            *[]string          `tfsdk:"redirect_uris" json:"redirectUris,omitempty"`
				RegisteredNodes         *map[string]string `tfsdk:"registered_nodes" json:"registeredNodes,omitempty"`
				RegistrationAccessToken *string            `tfsdk:"registration_access_token" json:"registrationAccessToken,omitempty"`
				RootUrl                 *string            `tfsdk:"root_url" json:"rootUrl,omitempty"`
				Secret                  *string            `tfsdk:"secret" json:"secret,omitempty"`
				ServiceAccountsEnabled  *bool              `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
				StandardFlowEnabled     *bool              `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
				SurrogateAuthRequired   *bool              `tfsdk:"surrogate_auth_required" json:"surrogateAuthRequired,omitempty"`
				UseTemplateConfig       *bool              `tfsdk:"use_template_config" json:"useTemplateConfig,omitempty"`
				UseTemplateMappers      *bool              `tfsdk:"use_template_mappers" json:"useTemplateMappers,omitempty"`
				UseTemplateScope        *bool              `tfsdk:"use_template_scope" json:"useTemplateScope,omitempty"`
				WebOrigins              *[]string          `tfsdk:"web_origins" json:"webOrigins,omitempty"`
			} `tfsdk:"oauth_clients" json:"oauthClients,omitempty"`
			OfflineSessionIdleTimeout        *int64    `tfsdk:"offline_session_idle_timeout" json:"offlineSessionIdleTimeout,omitempty"`
			OfflineSessionMaxLifespan        *int64    `tfsdk:"offline_session_max_lifespan" json:"offlineSessionMaxLifespan,omitempty"`
			OfflineSessionMaxLifespanEnabled *bool     `tfsdk:"offline_session_max_lifespan_enabled" json:"offlineSessionMaxLifespanEnabled,omitempty"`
			OtpPolicyAlgorithm               *string   `tfsdk:"otp_policy_algorithm" json:"otpPolicyAlgorithm,omitempty"`
			OtpPolicyDigits                  *int64    `tfsdk:"otp_policy_digits" json:"otpPolicyDigits,omitempty"`
			OtpPolicyInitialCounter          *int64    `tfsdk:"otp_policy_initial_counter" json:"otpPolicyInitialCounter,omitempty"`
			OtpPolicyLookAheadWindow         *int64    `tfsdk:"otp_policy_look_ahead_window" json:"otpPolicyLookAheadWindow,omitempty"`
			OtpPolicyPeriod                  *int64    `tfsdk:"otp_policy_period" json:"otpPolicyPeriod,omitempty"`
			OtpPolicyType                    *string   `tfsdk:"otp_policy_type" json:"otpPolicyType,omitempty"`
			OtpSupportedApplications         *[]string `tfsdk:"otp_supported_applications" json:"otpSupportedApplications,omitempty"`
			PasswordCredentialGrantAllowed   *bool     `tfsdk:"password_credential_grant_allowed" json:"passwordCredentialGrantAllowed,omitempty"`
			PasswordPolicy                   *string   `tfsdk:"password_policy" json:"passwordPolicy,omitempty"`
			PermanentLockout                 *bool     `tfsdk:"permanent_lockout" json:"permanentLockout,omitempty"`
			PrivateKey                       *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
			ProtocolMappers                  *[]struct {
				Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
				ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
				ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
				Id              *string            `tfsdk:"id" json:"id,omitempty"`
				Name            *string            `tfsdk:"name" json:"name,omitempty"`
				Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
			} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
			PublicKey                   *string `tfsdk:"public_key" json:"publicKey,omitempty"`
			QuickLoginCheckMilliSeconds *int64  `tfsdk:"quick_login_check_milli_seconds" json:"quickLoginCheckMilliSeconds,omitempty"`
			Realm                       *string `tfsdk:"realm" json:"realm,omitempty"`
			RealmCacheEnabled           *bool   `tfsdk:"realm_cache_enabled" json:"realmCacheEnabled,omitempty"`
			RefreshTokenMaxReuse        *int64  `tfsdk:"refresh_token_max_reuse" json:"refreshTokenMaxReuse,omitempty"`
			RegistrationAllowed         *bool   `tfsdk:"registration_allowed" json:"registrationAllowed,omitempty"`
			RegistrationEmailAsUsername *bool   `tfsdk:"registration_email_as_username" json:"registrationEmailAsUsername,omitempty"`
			RegistrationFlow            *string `tfsdk:"registration_flow" json:"registrationFlow,omitempty"`
			RememberMe                  *bool   `tfsdk:"remember_me" json:"rememberMe,omitempty"`
			RequiredActions             *[]struct {
				Alias         *string            `tfsdk:"alias" json:"alias,omitempty"`
				Config        *map[string]string `tfsdk:"config" json:"config,omitempty"`
				DefaultAction *bool              `tfsdk:"default_action" json:"defaultAction,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Name          *string            `tfsdk:"name" json:"name,omitempty"`
				Priority      *int64             `tfsdk:"priority" json:"priority,omitempty"`
				ProviderId    *string            `tfsdk:"provider_id" json:"providerId,omitempty"`
			} `tfsdk:"required_actions" json:"requiredActions,omitempty"`
			RequiredCredentials  *[]string `tfsdk:"required_credentials" json:"requiredCredentials,omitempty"`
			ResetCredentialsFlow *string   `tfsdk:"reset_credentials_flow" json:"resetCredentialsFlow,omitempty"`
			ResetPasswordAllowed *bool     `tfsdk:"reset_password_allowed" json:"resetPasswordAllowed,omitempty"`
			RevokeRefreshToken   *bool     `tfsdk:"revoke_refresh_token" json:"revokeRefreshToken,omitempty"`
			Roles                *struct {
				Application *map[string]string `tfsdk:"application" json:"application,omitempty"`
				Client      *map[string]string `tfsdk:"client" json:"client,omitempty"`
				Realm       *[]struct {
					Attributes *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
					ClientRole *bool                `tfsdk:"client_role" json:"clientRole,omitempty"`
					Composite  *bool                `tfsdk:"composite" json:"composite,omitempty"`
					Composites *struct {
						Application *map[string][]string `tfsdk:"application" json:"application,omitempty"`
						Client      *map[string][]string `tfsdk:"client" json:"client,omitempty"`
						Realm       *[]string            `tfsdk:"realm" json:"realm,omitempty"`
					} `tfsdk:"composites" json:"composites,omitempty"`
					ContainerId        *string `tfsdk:"container_id" json:"containerId,omitempty"`
					Description        *string `tfsdk:"description" json:"description,omitempty"`
					Id                 *string `tfsdk:"id" json:"id,omitempty"`
					Name               *string `tfsdk:"name" json:"name,omitempty"`
					ScopeParamRequired *bool   `tfsdk:"scope_param_required" json:"scopeParamRequired,omitempty"`
				} `tfsdk:"realm" json:"realm,omitempty"`
			} `tfsdk:"roles" json:"roles,omitempty"`
			ScopeMappings *[]struct {
				Client         *string   `tfsdk:"client" json:"client,omitempty"`
				ClientScope    *string   `tfsdk:"client_scope" json:"clientScope,omitempty"`
				ClientTemplate *string   `tfsdk:"client_template" json:"clientTemplate,omitempty"`
				Roles          *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Self           *string   `tfsdk:"self" json:"self,omitempty"`
			} `tfsdk:"scope_mappings" json:"scopeMappings,omitempty"`
			SmtpServer                        *map[string]string `tfsdk:"smtp_server" json:"smtpServer,omitempty"`
			Social                            *bool              `tfsdk:"social" json:"social,omitempty"`
			SocialProviders                   *map[string]string `tfsdk:"social_providers" json:"socialProviders,omitempty"`
			SslRequired                       *string            `tfsdk:"ssl_required" json:"sslRequired,omitempty"`
			SsoSessionIdleTimeout             *int64             `tfsdk:"sso_session_idle_timeout" json:"ssoSessionIdleTimeout,omitempty"`
			SsoSessionIdleTimeoutRememberMe   *int64             `tfsdk:"sso_session_idle_timeout_remember_me" json:"ssoSessionIdleTimeoutRememberMe,omitempty"`
			SsoSessionMaxLifespan             *int64             `tfsdk:"sso_session_max_lifespan" json:"ssoSessionMaxLifespan,omitempty"`
			SsoSessionMaxLifespanRememberMe   *int64             `tfsdk:"sso_session_max_lifespan_remember_me" json:"ssoSessionMaxLifespanRememberMe,omitempty"`
			SupportedLocales                  *[]string          `tfsdk:"supported_locales" json:"supportedLocales,omitempty"`
			UpdateProfileOnInitialSocialLogin *bool              `tfsdk:"update_profile_on_initial_social_login" json:"updateProfileOnInitialSocialLogin,omitempty"`
			UserCacheEnabled                  *bool              `tfsdk:"user_cache_enabled" json:"userCacheEnabled,omitempty"`
			UserFederationMappers             *[]struct {
				Config                        *map[string]string `tfsdk:"config" json:"config,omitempty"`
				FederationMapperType          *string            `tfsdk:"federation_mapper_type" json:"federationMapperType,omitempty"`
				FederationProviderDisplayName *string            `tfsdk:"federation_provider_display_name" json:"federationProviderDisplayName,omitempty"`
				Id                            *string            `tfsdk:"id" json:"id,omitempty"`
				Name                          *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"user_federation_mappers" json:"userFederationMappers,omitempty"`
			UserFederationProviders *[]struct {
				ChangedSyncPeriod *int64             `tfsdk:"changed_sync_period" json:"changedSyncPeriod,omitempty"`
				Config            *map[string]string `tfsdk:"config" json:"config,omitempty"`
				DisplayName       *string            `tfsdk:"display_name" json:"displayName,omitempty"`
				FullSyncPeriod    *int64             `tfsdk:"full_sync_period" json:"fullSyncPeriod,omitempty"`
				Id                *string            `tfsdk:"id" json:"id,omitempty"`
				LastSync          *int64             `tfsdk:"last_sync" json:"lastSync,omitempty"`
				Priority          *int64             `tfsdk:"priority" json:"priority,omitempty"`
				ProviderName      *string            `tfsdk:"provider_name" json:"providerName,omitempty"`
			} `tfsdk:"user_federation_providers" json:"userFederationProviders,omitempty"`
			UserManagedAccessAllowed *bool `tfsdk:"user_managed_access_allowed" json:"userManagedAccessAllowed,omitempty"`
			Users                    *[]struct {
				Access           *map[string]string   `tfsdk:"access" json:"access,omitempty"`
				ApplicationRoles *map[string][]string `tfsdk:"application_roles" json:"applicationRoles,omitempty"`
				Attributes       *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClientConsents   *[]struct {
					ClientId            *string   `tfsdk:"client_id" json:"clientId,omitempty"`
					CreatedDate         *int64    `tfsdk:"created_date" json:"createdDate,omitempty"`
					GrantedClientScopes *[]string `tfsdk:"granted_client_scopes" json:"grantedClientScopes,omitempty"`
					GrantedRealmRoles   *[]string `tfsdk:"granted_realm_roles" json:"grantedRealmRoles,omitempty"`
					LastUpdatedDate     *int64    `tfsdk:"last_updated_date" json:"lastUpdatedDate,omitempty"`
				} `tfsdk:"client_consents" json:"clientConsents,omitempty"`
				ClientRoles      *map[string][]string `tfsdk:"client_roles" json:"clientRoles,omitempty"`
				CreatedTimestamp *int64               `tfsdk:"created_timestamp" json:"createdTimestamp,omitempty"`
				Credentials      *[]struct {
					Algorithm         *string              `tfsdk:"algorithm" json:"algorithm,omitempty"`
					Config            *map[string][]string `tfsdk:"config" json:"config,omitempty"`
					Counter           *int64               `tfsdk:"counter" json:"counter,omitempty"`
					CreatedDate       *int64               `tfsdk:"created_date" json:"createdDate,omitempty"`
					CredentialData    *string              `tfsdk:"credential_data" json:"credentialData,omitempty"`
					Device            *string              `tfsdk:"device" json:"device,omitempty"`
					Digits            *int64               `tfsdk:"digits" json:"digits,omitempty"`
					HashIterations    *int64               `tfsdk:"hash_iterations" json:"hashIterations,omitempty"`
					HashedSaltedValue *string              `tfsdk:"hashed_salted_value" json:"hashedSaltedValue,omitempty"`
					Id                *string              `tfsdk:"id" json:"id,omitempty"`
					Period            *int64               `tfsdk:"period" json:"period,omitempty"`
					Priority          *int64               `tfsdk:"priority" json:"priority,omitempty"`
					Salt              *string              `tfsdk:"salt" json:"salt,omitempty"`
					SecretData        *string              `tfsdk:"secret_data" json:"secretData,omitempty"`
					Temporary         *bool                `tfsdk:"temporary" json:"temporary,omitempty"`
					Type              *string              `tfsdk:"type" json:"type,omitempty"`
					UserLabel         *string              `tfsdk:"user_label" json:"userLabel,omitempty"`
					Value             *string              `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				DisableableCredentialTypes *[]string `tfsdk:"disableable_credential_types" json:"disableableCredentialTypes,omitempty"`
				Email                      *string   `tfsdk:"email" json:"email,omitempty"`
				EmailVerified              *bool     `tfsdk:"email_verified" json:"emailVerified,omitempty"`
				Enabled                    *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				FederatedIdentities        *[]struct {
					IdentityProvider *string `tfsdk:"identity_provider" json:"identityProvider,omitempty"`
					UserId           *string `tfsdk:"user_id" json:"userId,omitempty"`
					UserName         *string `tfsdk:"user_name" json:"userName,omitempty"`
				} `tfsdk:"federated_identities" json:"federatedIdentities,omitempty"`
				FederationLink         *string   `tfsdk:"federation_link" json:"federationLink,omitempty"`
				FirstName              *string   `tfsdk:"first_name" json:"firstName,omitempty"`
				Groups                 *[]string `tfsdk:"groups" json:"groups,omitempty"`
				Id                     *string   `tfsdk:"id" json:"id,omitempty"`
				LastName               *string   `tfsdk:"last_name" json:"lastName,omitempty"`
				NotBefore              *int64    `tfsdk:"not_before" json:"notBefore,omitempty"`
				Origin                 *string   `tfsdk:"origin" json:"origin,omitempty"`
				RealmRoles             *[]string `tfsdk:"realm_roles" json:"realmRoles,omitempty"`
				RequiredActions        *[]string `tfsdk:"required_actions" json:"requiredActions,omitempty"`
				Self                   *string   `tfsdk:"self" json:"self,omitempty"`
				ServiceAccountClientId *string   `tfsdk:"service_account_client_id" json:"serviceAccountClientId,omitempty"`
				SocialLinks            *[]struct {
					SocialProvider *string `tfsdk:"social_provider" json:"socialProvider,omitempty"`
					SocialUserId   *string `tfsdk:"social_user_id" json:"socialUserId,omitempty"`
					SocialUsername *string `tfsdk:"social_username" json:"socialUsername,omitempty"`
				} `tfsdk:"social_links" json:"socialLinks,omitempty"`
				Totp     *bool   `tfsdk:"totp" json:"totp,omitempty"`
				Username *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"users" json:"users,omitempty"`
			VerifyEmail                                               *bool     `tfsdk:"verify_email" json:"verifyEmail,omitempty"`
			WaitIncrementSeconds                                      *int64    `tfsdk:"wait_increment_seconds" json:"waitIncrementSeconds,omitempty"`
			WebAuthnPolicyAcceptableAaguids                           *[]string `tfsdk:"web_authn_policy_acceptable_aaguids" json:"webAuthnPolicyAcceptableAaguids,omitempty"`
			WebAuthnPolicyAttestationConveyancePreference             *string   `tfsdk:"web_authn_policy_attestation_conveyance_preference" json:"webAuthnPolicyAttestationConveyancePreference,omitempty"`
			WebAuthnPolicyAuthenticatorAttachment                     *string   `tfsdk:"web_authn_policy_authenticator_attachment" json:"webAuthnPolicyAuthenticatorAttachment,omitempty"`
			WebAuthnPolicyAvoidSameAuthenticatorRegister              *bool     `tfsdk:"web_authn_policy_avoid_same_authenticator_register" json:"webAuthnPolicyAvoidSameAuthenticatorRegister,omitempty"`
			WebAuthnPolicyCreateTimeout                               *int64    `tfsdk:"web_authn_policy_create_timeout" json:"webAuthnPolicyCreateTimeout,omitempty"`
			WebAuthnPolicyPasswordlessAcceptableAaguids               *[]string `tfsdk:"web_authn_policy_passwordless_acceptable_aaguids" json:"webAuthnPolicyPasswordlessAcceptableAaguids,omitempty"`
			WebAuthnPolicyPasswordlessAttestationConveyancePreference *string   `tfsdk:"web_authn_policy_passwordless_attestation_conveyance_preference" json:"webAuthnPolicyPasswordlessAttestationConveyancePreference,omitempty"`
			WebAuthnPolicyPasswordlessAuthenticatorAttachment         *string   `tfsdk:"web_authn_policy_passwordless_authenticator_attachment" json:"webAuthnPolicyPasswordlessAuthenticatorAttachment,omitempty"`
			WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister  *bool     `tfsdk:"web_authn_policy_passwordless_avoid_same_authenticator_register" json:"webAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister,omitempty"`
			WebAuthnPolicyPasswordlessCreateTimeout                   *int64    `tfsdk:"web_authn_policy_passwordless_create_timeout" json:"webAuthnPolicyPasswordlessCreateTimeout,omitempty"`
			WebAuthnPolicyPasswordlessRequireResidentKey              *string   `tfsdk:"web_authn_policy_passwordless_require_resident_key" json:"webAuthnPolicyPasswordlessRequireResidentKey,omitempty"`
			WebAuthnPolicyPasswordlessRpEntityName                    *string   `tfsdk:"web_authn_policy_passwordless_rp_entity_name" json:"webAuthnPolicyPasswordlessRpEntityName,omitempty"`
			WebAuthnPolicyPasswordlessRpId                            *string   `tfsdk:"web_authn_policy_passwordless_rp_id" json:"webAuthnPolicyPasswordlessRpId,omitempty"`
			WebAuthnPolicyPasswordlessSignatureAlgorithms             *[]string `tfsdk:"web_authn_policy_passwordless_signature_algorithms" json:"webAuthnPolicyPasswordlessSignatureAlgorithms,omitempty"`
			WebAuthnPolicyPasswordlessUserVerificationRequirement     *string   `tfsdk:"web_authn_policy_passwordless_user_verification_requirement" json:"webAuthnPolicyPasswordlessUserVerificationRequirement,omitempty"`
			WebAuthnPolicyRequireResidentKey                          *string   `tfsdk:"web_authn_policy_require_resident_key" json:"webAuthnPolicyRequireResidentKey,omitempty"`
			WebAuthnPolicyRpEntityName                                *string   `tfsdk:"web_authn_policy_rp_entity_name" json:"webAuthnPolicyRpEntityName,omitempty"`
			WebAuthnPolicyRpId                                        *string   `tfsdk:"web_authn_policy_rp_id" json:"webAuthnPolicyRpId,omitempty"`
			WebAuthnPolicySignatureAlgorithms                         *[]string `tfsdk:"web_authn_policy_signature_algorithms" json:"webAuthnPolicySignatureAlgorithms,omitempty"`
			WebAuthnPolicyUserVerificationRequirement                 *string   `tfsdk:"web_authn_policy_user_verification_requirement" json:"webAuthnPolicyUserVerificationRequirement,omitempty"`
		} `tfsdk:"realm" json:"realm,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest"
}

func (r *K8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
					"keycloak_cr_name": schema.StringAttribute{
						Description:         "The name of the Keycloak CR to reference, in the same namespace.",
						MarkdownDescription: "The name of the Keycloak CR to reference, in the same namespace.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"realm": schema.SingleNestedAttribute{
						Description:         "The RealmRepresentation to import into Keycloak.",
						MarkdownDescription: "The RealmRepresentation to import into Keycloak.",
						Attributes: map[string]schema.Attribute{
							"access_code_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_code_lifespan_login": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_code_lifespan_user_action": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_token_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_token_lifespan_for_implicit_flow": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"account_theme": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"action_token_generated_by_admin_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"action_token_generated_by_user_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"admin_events_details_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"admin_events_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"admin_theme": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"application_scope_mappings": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"applications": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"admin_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"always_display_in_console": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authentication_flow_binding_overrides": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_services_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_settings": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_remote_resource_management": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"decision_strategy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("AFFIRMATIVE", "CONSENSUS", "UNANIMOUS"),
													},
												},

												"id": schema.StringAttribute{
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

												"policies": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"decision_strategy": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("AFFIRMATIVE", "CONSENSUS", "UNANIMOUS"),
																},
															},

															"description": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"logic": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("POSITIVE", "NEGATIVE"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"owner": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"policies": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources_data": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"_id": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"attributes": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.ListType{ElemType: types.StringType},
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
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

																		"owner": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"id": schema.StringAttribute{
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

																		"owner_managed_access": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"scopes": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"display_name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"icon_uri": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"id": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uris": schema.ListAttribute{
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

															"scopes": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes_data": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
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

												"policy_enforcement_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("PERMISSIVE", "ENFORCING", "DISABLED"),
													},
												},

												"resources": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"_id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.ListType{ElemType: types.StringType},
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"display_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
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

															"owner": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"id": schema.StringAttribute{
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

															"owner_managed_access": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uris": schema.ListAttribute{
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

												"scopes": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"display_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
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

										"base_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bearer_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"claims": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"address": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"email": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"gender": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"locale": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"phone": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"picture": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"profile": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"username": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"website": schema.BoolAttribute{
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

										"client_authenticator_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_template": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"consent_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_client_scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_access_grants_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_grants_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"frontchannel_logout": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"full_scope_allowed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"implicit_flow_enabled": schema.BoolAttribute{
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

										"node_re_registration_timeout": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_before": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"oauth2_device_authorization_grant_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional_client_scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"origin": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mappers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
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

													"protocol": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
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

										"public_client": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redirect_uris": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registered_nodes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registration_access_token": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"root_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"standard_flow_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"surrogate_auth_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_config": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_mappers": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_scope": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"web_origins": schema.ListAttribute{
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

							"attributes": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"authentication_flows": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authentication_executions": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"authenticator": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"authenticator_config": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"authenticator_flow": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"autheticator_flow": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"flow_alias": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"priority": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requirement": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_setup_allowed": schema.BoolAttribute{
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

										"built_in": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"top_level": schema.BoolAttribute{
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

							"authenticator_config": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
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

							"browser_flow": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"browser_security_headers": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"brute_force_protected": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"certificate": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_authentication_flow": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_offline_session_idle_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_offline_session_max_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_policies": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_profiles": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_scope_mappings": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_scopes": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
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

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mappers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
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

													"protocol": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
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

							"client_session_idle_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_session_max_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_templates": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bearer_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"consent_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_access_grants_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"frontchannel_logout": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"full_scope_allowed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"implicit_flow_enabled": schema.BoolAttribute{
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

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mappers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
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

													"protocol": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
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

										"public_client": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"standard_flow_enabled": schema.BoolAttribute{
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

							"clients": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"admin_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"always_display_in_console": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authentication_flow_binding_overrides": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_services_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_settings": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_remote_resource_management": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"decision_strategy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("AFFIRMATIVE", "CONSENSUS", "UNANIMOUS"),
													},
												},

												"id": schema.StringAttribute{
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

												"policies": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"decision_strategy": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("AFFIRMATIVE", "CONSENSUS", "UNANIMOUS"),
																},
															},

															"description": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"logic": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("POSITIVE", "NEGATIVE"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"owner": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"policies": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources_data": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"_id": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"attributes": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.ListType{ElemType: types.StringType},
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
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

																		"owner": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"id": schema.StringAttribute{
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

																		"owner_managed_access": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"scopes": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"display_name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"icon_uri": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"id": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uris": schema.ListAttribute{
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

															"scopes": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes_data": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
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

												"policy_enforcement_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("PERMISSIVE", "ENFORCING", "DISABLED"),
													},
												},

												"resources": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"_id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.ListType{ElemType: types.StringType},
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"display_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
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

															"owner": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"id": schema.StringAttribute{
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

															"owner_managed_access": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uris": schema.ListAttribute{
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

												"scopes": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"display_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
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

										"base_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bearer_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_authenticator_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_template": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"consent_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_client_scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_access_grants_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_grants_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"frontchannel_logout": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"full_scope_allowed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"implicit_flow_enabled": schema.BoolAttribute{
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

										"node_re_registration_timeout": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_before": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"oauth2_device_authorization_grant_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional_client_scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"origin": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mappers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
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

													"protocol": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
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

										"public_client": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redirect_uris": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registered_nodes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registration_access_token": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"root_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"standard_flow_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"surrogate_auth_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_config": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_mappers": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_scope": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"web_origins": schema.ListAttribute{
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

							"code_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"components": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_default_client_scopes": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_locale": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_optional_client_scopes": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_role": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"attributes": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_role": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"composite": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"composites": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"application": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"client": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"realm": schema.ListAttribute{
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

									"container_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"description": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"id": schema.StringAttribute{
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

									"scope_param_required": schema.BoolAttribute{
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

							"default_roles": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_signature_algorithm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"direct_grant_flow": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"display_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"display_name_html": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"docker_authentication_flow": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"duplicate_emails_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"edit_username_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"email_theme": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled_event_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"events_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"events_expiration": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"events_listeners": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"failure_factor": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"federated_users": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"application_roles": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_consents": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"client_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"created_date": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"granted_client_scopes": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"granted_realm_roles": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"last_updated_date": schema.Int64Attribute{
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

										"client_roles": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"created_timestamp": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credentials": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"algorithm": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.ListType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"counter": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"created_date": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credential_data": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"device": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"digits": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hash_iterations": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hashed_salted_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"period": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"priority": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"salt": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_data": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"temporary": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_label": schema.StringAttribute{
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

										"disableable_credential_types": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"email": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"email_verified": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"federated_identities": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"identity_provider": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_name": schema.StringAttribute{
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

										"federation_link": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"first_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"groups": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"last_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_before": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"origin": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"realm_roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"required_actions": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"self": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_account_client_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"social_links": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"social_provider": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"social_user_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"social_username": schema.StringAttribute{
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

										"totp": schema.BoolAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_roles": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
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

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"realm_roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_groups": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"access": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"attributes": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.ListType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_roles": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.ListType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
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

													"path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"realm_roles": schema.ListAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_mappers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"identity_provider_alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"identity_provider_mapper": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_providers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"add_read_token_role_on_create": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authenticate_by_default": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"display_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"first_broker_login_flow_alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"internal_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"link_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"post_broker_login_flow_alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"store_token": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"trust_email": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"update_profile_first_login_mode": schema.StringAttribute{
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

							"internationalization_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keycloak_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"login_theme": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"login_with_email_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_delta_time_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_failure_wait_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"minimum_quick_login_wait_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"not_before": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth2_device_code_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth2_device_polling_interval": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth_clients": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"admin_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"always_display_in_console": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authentication_flow_binding_overrides": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_services_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_settings": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_remote_resource_management": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"decision_strategy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("AFFIRMATIVE", "CONSENSUS", "UNANIMOUS"),
													},
												},

												"id": schema.StringAttribute{
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

												"policies": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"decision_strategy": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("AFFIRMATIVE", "CONSENSUS", "UNANIMOUS"),
																},
															},

															"description": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"logic": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("POSITIVE", "NEGATIVE"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"owner": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"policies": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources_data": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"_id": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"attributes": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.ListType{ElemType: types.StringType},
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
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

																		"owner": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"id": schema.StringAttribute{
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

																		"owner_managed_access": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"scopes": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"display_name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"icon_uri": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"id": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uris": schema.ListAttribute{
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

															"scopes": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes_data": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
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

												"policy_enforcement_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("PERMISSIVE", "ENFORCING", "DISABLED"),
													},
												},

												"resources": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"_id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.ListType{ElemType: types.StringType},
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"display_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
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

															"owner": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"id": schema.StringAttribute{
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

															"owner_managed_access": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"display_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uris": schema.ListAttribute{
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

												"scopes": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"display_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
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

										"base_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bearer_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"claims": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"address": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"email": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"gender": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"locale": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"phone": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"picture": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"profile": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"username": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"website": schema.BoolAttribute{
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

										"client_authenticator_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_template": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"consent_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_client_scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_access_grants_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_grants_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"frontchannel_logout": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"full_scope_allowed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"implicit_flow_enabled": schema.BoolAttribute{
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

										"node_re_registration_timeout": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_before": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"oauth2_device_authorization_grant_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional_client_scopes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"origin": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mappers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
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

													"protocol": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
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

										"public_client": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redirect_uris": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registered_nodes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registration_access_token": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"root_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"standard_flow_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"surrogate_auth_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_config": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_mappers": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_scope": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"web_origins": schema.ListAttribute{
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

							"offline_session_idle_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"offline_session_max_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"offline_session_max_lifespan_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_algorithm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_digits": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_initial_counter": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_look_ahead_window": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_period": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_supported_applications": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password_credential_grant_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"permanent_lockout": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocol_mappers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"consent_required": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"consent_text": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
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

										"protocol": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mapper": schema.StringAttribute{
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

							"public_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"quick_login_check_milli_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm_cache_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"refresh_token_max_reuse": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_email_as_username": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_flow": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remember_me": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"required_actions": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_action": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
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

										"priority": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_id": schema.StringAttribute{
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

							"required_credentials": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reset_credentials_flow": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reset_password_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"revoke_refresh_token": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"roles": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"application": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"realm": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_role": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"composite": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"composites": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"application": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"realm": schema.ListAttribute{
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

												"container_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
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

												"scope_param_required": schema.BoolAttribute{
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

							"scope_mappings": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"client": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_scope": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_template": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"self": schema.StringAttribute{
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

							"smtp_server": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"social": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"social_providers": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_required": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sso_session_idle_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sso_session_idle_timeout_remember_me": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sso_session_max_lifespan": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sso_session_max_lifespan_remember_me": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supported_locales": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_profile_on_initial_social_login": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_cache_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_federation_mappers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"federation_mapper_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"federation_provider_display_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_federation_providers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"changed_sync_period": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"display_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"full_sync_period": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"last_sync": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"priority": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_name": schema.StringAttribute{
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

							"user_managed_access_allowed": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"users": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"application_roles": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_consents": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"client_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"created_date": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"granted_client_scopes": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"granted_realm_roles": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"last_updated_date": schema.Int64Attribute{
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

										"client_roles": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"created_timestamp": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credentials": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"algorithm": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"config": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.ListType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"counter": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"created_date": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credential_data": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"device": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"digits": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hash_iterations": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hashed_salted_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"period": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"priority": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"salt": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_data": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"temporary": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_label": schema.StringAttribute{
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

										"disableable_credential_types": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"email": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"email_verified": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"federated_identities": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"identity_provider": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_name": schema.StringAttribute{
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

										"federation_link": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"first_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"groups": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"last_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_before": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"origin": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"realm_roles": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"required_actions": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"self": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_account_client_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"social_links": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"social_provider": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"social_user_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"social_username": schema.StringAttribute{
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

										"totp": schema.BoolAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"verify_email": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_increment_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_acceptable_aaguids": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_attestation_conveyance_preference": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_authenticator_attachment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_avoid_same_authenticator_register": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_create_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_acceptable_aaguids": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_attestation_conveyance_preference": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_authenticator_attachment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_avoid_same_authenticator_register": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_create_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_require_resident_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_rp_entity_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_rp_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_signature_algorithms": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_passwordless_user_verification_requirement": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_require_resident_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_rp_entity_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_rp_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_signature_algorithms": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_authn_policy_user_verification_requirement": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
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
		},
	}
}

func (r *K8SKeycloakOrgKeycloakRealmImportV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest")

	var model K8SKeycloakOrgKeycloakRealmImportV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.keycloak.org/v2alpha1")
	model.Kind = pointer.String("KeycloakRealmImport")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
