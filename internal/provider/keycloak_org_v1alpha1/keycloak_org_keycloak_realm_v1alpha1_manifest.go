/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keycloak_org_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KeycloakOrgKeycloakRealmV1Alpha1Manifest{}
)

func NewKeycloakOrgKeycloakRealmV1Alpha1Manifest() datasource.DataSource {
	return &KeycloakOrgKeycloakRealmV1Alpha1Manifest{}
}

type KeycloakOrgKeycloakRealmV1Alpha1Manifest struct{}

type KeycloakOrgKeycloakRealmV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		InstanceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"instance_selector" json:"instanceSelector,omitempty"`
		Realm *struct {
			AccessTokenLifespan                *int64  `tfsdk:"access_token_lifespan" json:"accessTokenLifespan,omitempty"`
			AccessTokenLifespanForImplicitFlow *int64  `tfsdk:"access_token_lifespan_for_implicit_flow" json:"accessTokenLifespanForImplicitFlow,omitempty"`
			AccountTheme                       *string `tfsdk:"account_theme" json:"accountTheme,omitempty"`
			AdminEventsDetailsEnabled          *bool   `tfsdk:"admin_events_details_enabled" json:"adminEventsDetailsEnabled,omitempty"`
			AdminEventsEnabled                 *bool   `tfsdk:"admin_events_enabled" json:"adminEventsEnabled,omitempty"`
			AdminTheme                         *string `tfsdk:"admin_theme" json:"adminTheme,omitempty"`
			AuthenticationFlows                *[]struct {
				Alias                    *string `tfsdk:"alias" json:"alias,omitempty"`
				AuthenticationExecutions *[]struct {
					Authenticator       *string `tfsdk:"authenticator" json:"authenticator,omitempty"`
					AuthenticatorConfig *string `tfsdk:"authenticator_config" json:"authenticatorConfig,omitempty"`
					AuthenticatorFlow   *bool   `tfsdk:"authenticator_flow" json:"authenticatorFlow,omitempty"`
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
			BrowserFlow              *string            `tfsdk:"browser_flow" json:"browserFlow,omitempty"`
			BruteForceProtected      *bool              `tfsdk:"brute_force_protected" json:"bruteForceProtected,omitempty"`
			ClientAuthenticationFlow *string            `tfsdk:"client_authentication_flow" json:"clientAuthenticationFlow,omitempty"`
			ClientScopeMappings      *map[string]string `tfsdk:"client_scope_mappings" json:"clientScopeMappings,omitempty"`
			ClientScopes             *[]struct {
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
			Clients *[]struct {
				Access                             *map[string]string `tfsdk:"access" json:"access,omitempty"`
				AdminUrl                           *string            `tfsdk:"admin_url" json:"adminUrl,omitempty"`
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
							_id                *string            `tfsdk:"_id" json:"_id,omitempty"`
							Attributes         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
							DisplayName        *string            `tfsdk:"display_name" json:"displayName,omitempty"`
							Icon_uri           *string            `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
							Name               *string            `tfsdk:"name" json:"name,omitempty"`
							OwnerManagedAccess *bool              `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
							Scopes             *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
							Type               *string            `tfsdk:"type" json:"type,omitempty"`
							Uris               *[]string          `tfsdk:"uris" json:"uris,omitempty"`
						} `tfsdk:"resources_data" json:"resourcesData,omitempty"`
						Scopes     *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
						ScopesData *[]string `tfsdk:"scopes_data" json:"scopesData,omitempty"`
						Type       *string   `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
					PolicyEnforcementMode *string `tfsdk:"policy_enforcement_mode" json:"policyEnforcementMode,omitempty"`
					Resources             *[]struct {
						_id                *string            `tfsdk:"_id" json:"_id,omitempty"`
						Attributes         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						DisplayName        *string            `tfsdk:"display_name" json:"displayName,omitempty"`
						Icon_uri           *string            `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
						Name               *string            `tfsdk:"name" json:"name,omitempty"`
						OwnerManagedAccess *bool              `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
						Scopes             *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						Type               *string            `tfsdk:"type" json:"type,omitempty"`
						Uris               *[]string          `tfsdk:"uris" json:"uris,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Scopes *[]struct {
						DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
						IconUri     *string `tfsdk:"icon_uri" json:"iconUri,omitempty"`
						Id          *string `tfsdk:"id" json:"id,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						Policies    *[]struct {
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
								_id                *string            `tfsdk:"_id" json:"_id,omitempty"`
								Attributes         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								DisplayName        *string            `tfsdk:"display_name" json:"displayName,omitempty"`
								Icon_uri           *string            `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
								Name               *string            `tfsdk:"name" json:"name,omitempty"`
								OwnerManagedAccess *bool              `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
								Scopes             *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
								Type               *string            `tfsdk:"type" json:"type,omitempty"`
								Uris               *[]string          `tfsdk:"uris" json:"uris,omitempty"`
							} `tfsdk:"resources_data" json:"resourcesData,omitempty"`
							Scopes     *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
							ScopesData *[]string `tfsdk:"scopes_data" json:"scopesData,omitempty"`
							Type       *string   `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"policies" json:"policies,omitempty"`
						Resources *[]struct {
							_id                *string            `tfsdk:"_id" json:"_id,omitempty"`
							Attributes         *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
							DisplayName        *string            `tfsdk:"display_name" json:"displayName,omitempty"`
							Icon_uri           *string            `tfsdk:"icon_uri" json:"icon_uri,omitempty"`
							Name               *string            `tfsdk:"name" json:"name,omitempty"`
							OwnerManagedAccess *bool              `tfsdk:"owner_managed_access" json:"ownerManagedAccess,omitempty"`
							Scopes             *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
							Type               *string            `tfsdk:"type" json:"type,omitempty"`
							Uris               *[]string          `tfsdk:"uris" json:"uris,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"scopes" json:"scopes,omitempty"`
				} `tfsdk:"authorization_settings" json:"authorizationSettings,omitempty"`
				BaseUrl                   *string   `tfsdk:"base_url" json:"baseUrl,omitempty"`
				BearerOnly                *bool     `tfsdk:"bearer_only" json:"bearerOnly,omitempty"`
				ClientAuthenticatorType   *string   `tfsdk:"client_authenticator_type" json:"clientAuthenticatorType,omitempty"`
				ClientId                  *string   `tfsdk:"client_id" json:"clientId,omitempty"`
				ConsentRequired           *bool     `tfsdk:"consent_required" json:"consentRequired,omitempty"`
				DefaultClientScopes       *[]string `tfsdk:"default_client_scopes" json:"defaultClientScopes,omitempty"`
				DefaultRoles              *[]string `tfsdk:"default_roles" json:"defaultRoles,omitempty"`
				Description               *string   `tfsdk:"description" json:"description,omitempty"`
				DirectAccessGrantsEnabled *bool     `tfsdk:"direct_access_grants_enabled" json:"directAccessGrantsEnabled,omitempty"`
				Enabled                   *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				FrontchannelLogout        *bool     `tfsdk:"frontchannel_logout" json:"frontchannelLogout,omitempty"`
				FullScopeAllowed          *bool     `tfsdk:"full_scope_allowed" json:"fullScopeAllowed,omitempty"`
				Id                        *string   `tfsdk:"id" json:"id,omitempty"`
				ImplicitFlowEnabled       *bool     `tfsdk:"implicit_flow_enabled" json:"implicitFlowEnabled,omitempty"`
				Name                      *string   `tfsdk:"name" json:"name,omitempty"`
				NodeReRegistrationTimeout *int64    `tfsdk:"node_re_registration_timeout" json:"nodeReRegistrationTimeout,omitempty"`
				NotBefore                 *int64    `tfsdk:"not_before" json:"notBefore,omitempty"`
				OptionalClientScopes      *[]string `tfsdk:"optional_client_scopes" json:"optionalClientScopes,omitempty"`
				Protocol                  *string   `tfsdk:"protocol" json:"protocol,omitempty"`
				ProtocolMappers           *[]struct {
					Config          *map[string]string `tfsdk:"config" json:"config,omitempty"`
					ConsentRequired *bool              `tfsdk:"consent_required" json:"consentRequired,omitempty"`
					ConsentText     *string            `tfsdk:"consent_text" json:"consentText,omitempty"`
					Id              *string            `tfsdk:"id" json:"id,omitempty"`
					Name            *string            `tfsdk:"name" json:"name,omitempty"`
					Protocol        *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					ProtocolMapper  *string            `tfsdk:"protocol_mapper" json:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" json:"protocolMappers,omitempty"`
				PublicClient           *bool     `tfsdk:"public_client" json:"publicClient,omitempty"`
				RedirectUris           *[]string `tfsdk:"redirect_uris" json:"redirectUris,omitempty"`
				RootUrl                *string   `tfsdk:"root_url" json:"rootUrl,omitempty"`
				Secret                 *string   `tfsdk:"secret" json:"secret,omitempty"`
				ServiceAccountsEnabled *bool     `tfsdk:"service_accounts_enabled" json:"serviceAccountsEnabled,omitempty"`
				StandardFlowEnabled    *bool     `tfsdk:"standard_flow_enabled" json:"standardFlowEnabled,omitempty"`
				SurrogateAuthRequired  *bool     `tfsdk:"surrogate_auth_required" json:"surrogateAuthRequired,omitempty"`
				UseTemplateConfig      *bool     `tfsdk:"use_template_config" json:"useTemplateConfig,omitempty"`
				UseTemplateMappers     *bool     `tfsdk:"use_template_mappers" json:"useTemplateMappers,omitempty"`
				UseTemplateScope       *bool     `tfsdk:"use_template_scope" json:"useTemplateScope,omitempty"`
				WebOrigins             *[]string `tfsdk:"web_origins" json:"webOrigins,omitempty"`
			} `tfsdk:"clients" json:"clients,omitempty"`
			DefaultDefaultClientScopes *[]string `tfsdk:"default_default_client_scopes" json:"defaultDefaultClientScopes,omitempty"`
			DefaultLocale              *string   `tfsdk:"default_locale" json:"defaultLocale,omitempty"`
			DefaultRole                *struct {
				Attributes *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClientRole *bool                `tfsdk:"client_role" json:"clientRole,omitempty"`
				Composite  *bool                `tfsdk:"composite" json:"composite,omitempty"`
				Composites *struct {
					Client *map[string][]string `tfsdk:"client" json:"client,omitempty"`
					Realm  *[]string            `tfsdk:"realm" json:"realm,omitempty"`
				} `tfsdk:"composites" json:"composites,omitempty"`
				ContainerId *string `tfsdk:"container_id" json:"containerId,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Id          *string `tfsdk:"id" json:"id,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"default_role" json:"defaultRole,omitempty"`
			DirectGrantFlow          *string   `tfsdk:"direct_grant_flow" json:"directGrantFlow,omitempty"`
			DisplayName              *string   `tfsdk:"display_name" json:"displayName,omitempty"`
			DisplayNameHtml          *string   `tfsdk:"display_name_html" json:"displayNameHtml,omitempty"`
			DockerAuthenticationFlow *string   `tfsdk:"docker_authentication_flow" json:"dockerAuthenticationFlow,omitempty"`
			DuplicateEmailsAllowed   *bool     `tfsdk:"duplicate_emails_allowed" json:"duplicateEmailsAllowed,omitempty"`
			EditUsernameAllowed      *bool     `tfsdk:"edit_username_allowed" json:"editUsernameAllowed,omitempty"`
			EmailTheme               *string   `tfsdk:"email_theme" json:"emailTheme,omitempty"`
			Enabled                  *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
			EnabledEventTypes        *[]string `tfsdk:"enabled_event_types" json:"enabledEventTypes,omitempty"`
			EventsEnabled            *bool     `tfsdk:"events_enabled" json:"eventsEnabled,omitempty"`
			EventsListeners          *[]string `tfsdk:"events_listeners" json:"eventsListeners,omitempty"`
			FailureFactor            *int64    `tfsdk:"failure_factor" json:"failureFactor,omitempty"`
			Id                       *string   `tfsdk:"id" json:"id,omitempty"`
			IdentityProviderMappers  *[]struct {
				Config                 *map[string]string `tfsdk:"config" json:"config,omitempty"`
				Id                     *string            `tfsdk:"id" json:"id,omitempty"`
				IdentityProviderAlias  *string            `tfsdk:"identity_provider_alias" json:"identityProviderAlias,omitempty"`
				IdentityProviderMapper *string            `tfsdk:"identity_provider_mapper" json:"identityProviderMapper,omitempty"`
				Name                   *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"identity_provider_mappers" json:"identityProviderMappers,omitempty"`
			IdentityProviders *[]struct {
				AddReadTokenRoleOnCreate  *bool              `tfsdk:"add_read_token_role_on_create" json:"addReadTokenRoleOnCreate,omitempty"`
				Alias                     *string            `tfsdk:"alias" json:"alias,omitempty"`
				Config                    *map[string]string `tfsdk:"config" json:"config,omitempty"`
				DisplayName               *string            `tfsdk:"display_name" json:"displayName,omitempty"`
				Enabled                   *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				FirstBrokerLoginFlowAlias *string            `tfsdk:"first_broker_login_flow_alias" json:"firstBrokerLoginFlowAlias,omitempty"`
				InternalId                *string            `tfsdk:"internal_id" json:"internalId,omitempty"`
				LinkOnly                  *bool              `tfsdk:"link_only" json:"linkOnly,omitempty"`
				PostBrokerLoginFlowAlias  *string            `tfsdk:"post_broker_login_flow_alias" json:"postBrokerLoginFlowAlias,omitempty"`
				ProviderId                *string            `tfsdk:"provider_id" json:"providerId,omitempty"`
				StoreToken                *bool              `tfsdk:"store_token" json:"storeToken,omitempty"`
				TrustEmail                *bool              `tfsdk:"trust_email" json:"trustEmail,omitempty"`
			} `tfsdk:"identity_providers" json:"identityProviders,omitempty"`
			InternationalizationEnabled  *bool     `tfsdk:"internationalization_enabled" json:"internationalizationEnabled,omitempty"`
			LoginTheme                   *string   `tfsdk:"login_theme" json:"loginTheme,omitempty"`
			LoginWithEmailAllowed        *bool     `tfsdk:"login_with_email_allowed" json:"loginWithEmailAllowed,omitempty"`
			MaxDeltaTimeSeconds          *int64    `tfsdk:"max_delta_time_seconds" json:"maxDeltaTimeSeconds,omitempty"`
			MaxFailureWaitSeconds        *int64    `tfsdk:"max_failure_wait_seconds" json:"maxFailureWaitSeconds,omitempty"`
			MinimumQuickLoginWaitSeconds *int64    `tfsdk:"minimum_quick_login_wait_seconds" json:"minimumQuickLoginWaitSeconds,omitempty"`
			OtpPolicyAlgorithm           *string   `tfsdk:"otp_policy_algorithm" json:"otpPolicyAlgorithm,omitempty"`
			OtpPolicyDigits              *int64    `tfsdk:"otp_policy_digits" json:"otpPolicyDigits,omitempty"`
			OtpPolicyInitialCounter      *int64    `tfsdk:"otp_policy_initial_counter" json:"otpPolicyInitialCounter,omitempty"`
			OtpPolicyLookAheadWindow     *int64    `tfsdk:"otp_policy_look_ahead_window" json:"otpPolicyLookAheadWindow,omitempty"`
			OtpPolicyPeriod              *int64    `tfsdk:"otp_policy_period" json:"otpPolicyPeriod,omitempty"`
			OtpPolicyType                *string   `tfsdk:"otp_policy_type" json:"otpPolicyType,omitempty"`
			OtpSupportedApplications     *[]string `tfsdk:"otp_supported_applications" json:"otpSupportedApplications,omitempty"`
			PasswordPolicy               *string   `tfsdk:"password_policy" json:"passwordPolicy,omitempty"`
			PermanentLockout             *bool     `tfsdk:"permanent_lockout" json:"permanentLockout,omitempty"`
			QuickLoginCheckMilliSeconds  *int64    `tfsdk:"quick_login_check_milli_seconds" json:"quickLoginCheckMilliSeconds,omitempty"`
			Realm                        *string   `tfsdk:"realm" json:"realm,omitempty"`
			RegistrationAllowed          *bool     `tfsdk:"registration_allowed" json:"registrationAllowed,omitempty"`
			RegistrationEmailAsUsername  *bool     `tfsdk:"registration_email_as_username" json:"registrationEmailAsUsername,omitempty"`
			RegistrationFlow             *string   `tfsdk:"registration_flow" json:"registrationFlow,omitempty"`
			RememberMe                   *bool     `tfsdk:"remember_me" json:"rememberMe,omitempty"`
			ResetCredentialsFlow         *string   `tfsdk:"reset_credentials_flow" json:"resetCredentialsFlow,omitempty"`
			ResetPasswordAllowed         *bool     `tfsdk:"reset_password_allowed" json:"resetPasswordAllowed,omitempty"`
			Roles                        *struct {
				Client *map[string]string `tfsdk:"client" json:"client,omitempty"`
				Realm  *[]struct {
					Attributes *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
					ClientRole *bool                `tfsdk:"client_role" json:"clientRole,omitempty"`
					Composite  *bool                `tfsdk:"composite" json:"composite,omitempty"`
					Composites *struct {
						Client *map[string][]string `tfsdk:"client" json:"client,omitempty"`
						Realm  *[]string            `tfsdk:"realm" json:"realm,omitempty"`
					} `tfsdk:"composites" json:"composites,omitempty"`
					ContainerId *string `tfsdk:"container_id" json:"containerId,omitempty"`
					Description *string `tfsdk:"description" json:"description,omitempty"`
					Id          *string `tfsdk:"id" json:"id,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"realm" json:"realm,omitempty"`
			} `tfsdk:"roles" json:"roles,omitempty"`
			ScopeMappings *[]struct {
				Client      *string   `tfsdk:"client" json:"client,omitempty"`
				ClientScope *string   `tfsdk:"client_scope" json:"clientScope,omitempty"`
				Roles       *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Self        *string   `tfsdk:"self" json:"self,omitempty"`
			} `tfsdk:"scope_mappings" json:"scopeMappings,omitempty"`
			SmtpServer            *map[string]string `tfsdk:"smtp_server" json:"smtpServer,omitempty"`
			SslRequired           *string            `tfsdk:"ssl_required" json:"sslRequired,omitempty"`
			SupportedLocales      *[]string          `tfsdk:"supported_locales" json:"supportedLocales,omitempty"`
			UserFederationMappers *[]struct {
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
				Priority          *int64             `tfsdk:"priority" json:"priority,omitempty"`
				ProviderName      *string            `tfsdk:"provider_name" json:"providerName,omitempty"`
			} `tfsdk:"user_federation_providers" json:"userFederationProviders,omitempty"`
			UserManagedAccessAllowed *bool `tfsdk:"user_managed_access_allowed" json:"userManagedAccessAllowed,omitempty"`
			Users                    *[]struct {
				Attributes  *map[string][]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClientRoles *map[string][]string `tfsdk:"client_roles" json:"clientRoles,omitempty"`
				Credentials *[]struct {
					Temporary *bool   `tfsdk:"temporary" json:"temporary,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Email               *string `tfsdk:"email" json:"email,omitempty"`
				EmailVerified       *bool   `tfsdk:"email_verified" json:"emailVerified,omitempty"`
				Enabled             *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				FederatedIdentities *[]struct {
					IdentityProvider *string `tfsdk:"identity_provider" json:"identityProvider,omitempty"`
					UserId           *string `tfsdk:"user_id" json:"userId,omitempty"`
					UserName         *string `tfsdk:"user_name" json:"userName,omitempty"`
				} `tfsdk:"federated_identities" json:"federatedIdentities,omitempty"`
				FirstName       *string   `tfsdk:"first_name" json:"firstName,omitempty"`
				Groups          *[]string `tfsdk:"groups" json:"groups,omitempty"`
				Id              *string   `tfsdk:"id" json:"id,omitempty"`
				LastName        *string   `tfsdk:"last_name" json:"lastName,omitempty"`
				RealmRoles      *[]string `tfsdk:"realm_roles" json:"realmRoles,omitempty"`
				RequiredActions *[]string `tfsdk:"required_actions" json:"requiredActions,omitempty"`
				Username        *string   `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"users" json:"users,omitempty"`
			VerifyEmail          *bool  `tfsdk:"verify_email" json:"verifyEmail,omitempty"`
			WaitIncrementSeconds *int64 `tfsdk:"wait_increment_seconds" json:"waitIncrementSeconds,omitempty"`
		} `tfsdk:"realm" json:"realm,omitempty"`
		RealmOverrides *[]struct {
			ForFlow          *string `tfsdk:"for_flow" json:"forFlow,omitempty"`
			IdentityProvider *string `tfsdk:"identity_provider" json:"identityProvider,omitempty"`
		} `tfsdk:"realm_overrides" json:"realmOverrides,omitempty"`
		Unmanaged *bool `tfsdk:"unmanaged" json:"unmanaged,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keycloak_org_keycloak_realm_v1alpha1_manifest"
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KeycloakRealm is the Schema for the keycloakrealms API",
		MarkdownDescription: "KeycloakRealm is the Schema for the keycloakrealms API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "KeycloakRealmSpec defines the desired state of KeycloakRealm.",
				MarkdownDescription: "KeycloakRealmSpec defines the desired state of KeycloakRealm.",
				Attributes: map[string]schema.Attribute{
					"instance_selector": schema.SingleNestedAttribute{
						Description:         "Selector for looking up Keycloak Custom Resources.",
						MarkdownDescription: "Selector for looking up Keycloak Custom Resources.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"realm": schema.SingleNestedAttribute{
						Description:         "Keycloak Realm REST object.",
						MarkdownDescription: "Keycloak Realm REST object.",
						Attributes: map[string]schema.Attribute{
							"access_token_lifespan": schema.Int64Attribute{
								Description:         "Access Token Lifespan",
								MarkdownDescription: "Access Token Lifespan",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_token_lifespan_for_implicit_flow": schema.Int64Attribute{
								Description:         "Access Token Lifespan For Implicit Flow",
								MarkdownDescription: "Access Token Lifespan For Implicit Flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"account_theme": schema.StringAttribute{
								Description:         "Account Theme",
								MarkdownDescription: "Account Theme",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"admin_events_details_enabled": schema.BoolAttribute{
								Description:         "Enable admin events details TODO: change to values and use kubebuilder default annotation once supported",
								MarkdownDescription: "Enable admin events details TODO: change to values and use kubebuilder default annotation once supported",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"admin_events_enabled": schema.BoolAttribute{
								Description:         "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",
								MarkdownDescription: "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"admin_theme": schema.StringAttribute{
								Description:         "Admin Console Theme",
								MarkdownDescription: "Admin Console Theme",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"authentication_flows": schema.ListNestedAttribute{
								Description:         "Authentication flows",
								MarkdownDescription: "Authentication flows",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias",
											MarkdownDescription: "Alias",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"authentication_executions": schema.ListNestedAttribute{
											Description:         "Authentication executions",
											MarkdownDescription: "Authentication executions",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"authenticator": schema.StringAttribute{
														Description:         "Authenticator",
														MarkdownDescription: "Authenticator",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"authenticator_config": schema.StringAttribute{
														Description:         "Authenticator Config",
														MarkdownDescription: "Authenticator Config",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"authenticator_flow": schema.BoolAttribute{
														Description:         "Authenticator flow",
														MarkdownDescription: "Authenticator flow",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"flow_alias": schema.StringAttribute{
														Description:         "Flow Alias",
														MarkdownDescription: "Flow Alias",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"priority": schema.Int64Attribute{
														Description:         "Priority",
														MarkdownDescription: "Priority",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requirement": schema.StringAttribute{
														Description:         "Requirement [REQUIRED, OPTIONAL, ALTERNATIVE, DISABLED]",
														MarkdownDescription: "Requirement [REQUIRED, OPTIONAL, ALTERNATIVE, DISABLED]",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_setup_allowed": schema.BoolAttribute{
														Description:         "User setup allowed",
														MarkdownDescription: "User setup allowed",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"built_in": schema.BoolAttribute{
											Description:         "Built in",
											MarkdownDescription: "Built in",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "Description",
											MarkdownDescription: "Description",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "ID",
											MarkdownDescription: "ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_id": schema.StringAttribute{
											Description:         "Provider ID",
											MarkdownDescription: "Provider ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"top_level": schema.BoolAttribute{
											Description:         "Top level",
											MarkdownDescription: "Top level",
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
								Description:         "Authenticator config",
								MarkdownDescription: "Authenticator config",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"alias": schema.StringAttribute{
											Description:         "Alias",
											MarkdownDescription: "Alias",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"config": schema.MapAttribute{
											Description:         "Config",
											MarkdownDescription: "Config",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "ID",
											MarkdownDescription: "ID",
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
								Description:         "Browser authentication flow",
								MarkdownDescription: "Browser authentication flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"brute_force_protected": schema.BoolAttribute{
								Description:         "Brute Force Detection",
								MarkdownDescription: "Brute Force Detection",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_authentication_flow": schema.StringAttribute{
								Description:         "Client authentication flow",
								MarkdownDescription: "Client authentication flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_scope_mappings": schema.MapAttribute{
								Description:         "Client Scope Mappings",
								MarkdownDescription: "Client Scope Mappings",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_scopes": schema.ListNestedAttribute{
								Description:         "Client scopes",
								MarkdownDescription: "Client scopes",
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
											Description:         "Protocol Mappers.",
											MarkdownDescription: "Protocol Mappers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "Config options.",
														MarkdownDescription: "Config options.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "True if Consent Screen is required.",
														MarkdownDescription: "True if Consent Screen is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "Text to use for displaying Consent Screen.",
														MarkdownDescription: "Text to use for displaying Consent Screen.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "Protocol Mapper ID.",
														MarkdownDescription: "Protocol Mapper ID.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Protocol Mapper Name.",
														MarkdownDescription: "Protocol Mapper Name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol to use.",
														MarkdownDescription: "Protocol to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
														Description:         "Protocol Mapper to use",
														MarkdownDescription: "Protocol Mapper to use",
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

							"clients": schema.ListNestedAttribute{
								Description:         "A set of Keycloak Clients.",
								MarkdownDescription: "A set of Keycloak Clients.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"access": schema.MapAttribute{
											Description:         "Access options.",
											MarkdownDescription: "Access options.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"admin_url": schema.StringAttribute{
											Description:         "Application Admin URL.",
											MarkdownDescription: "Application Admin URL.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"attributes": schema.MapAttribute{
											Description:         "Client Attributes.",
											MarkdownDescription: "Client Attributes.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authentication_flow_binding_overrides": schema.MapAttribute{
											Description:         "Authentication Flow Binding Overrides.",
											MarkdownDescription: "Authentication Flow Binding Overrides.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_services_enabled": schema.BoolAttribute{
											Description:         "True if fine-grained authorization support is enabled for this client.",
											MarkdownDescription: "True if fine-grained authorization support is enabled for this client.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"authorization_settings": schema.SingleNestedAttribute{
											Description:         "Authorization settings for this resource server.",
											MarkdownDescription: "Authorization settings for this resource server.",
											Attributes: map[string]schema.Attribute{
												"allow_remote_resource_management": schema.BoolAttribute{
													Description:         "True if resources should be managed remotely by the resource server.",
													MarkdownDescription: "True if resources should be managed remotely by the resource server.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_id": schema.StringAttribute{
													Description:         "Client ID.",
													MarkdownDescription: "Client ID.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"decision_strategy": schema.StringAttribute{
													Description:         "The decision strategy dictates how permissions are evaluated and how a final decision is obtained. 'Affirmative' means that at least one permission must evaluate to a positive decision in order to grant access to a resource and its scopes. 'Unanimous' means that all permissions must evaluate to a positive decision in order for the final decision to be also positive.",
													MarkdownDescription: "The decision strategy dictates how permissions are evaluated and how a final decision is obtained. 'Affirmative' means that at least one permission must evaluate to a positive decision in order to grant access to a resource and its scopes. 'Unanimous' means that all permissions must evaluate to a positive decision in order for the final decision to be also positive.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
													Description:         "ID.",
													MarkdownDescription: "ID.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name.",
													MarkdownDescription: "Name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"policies": schema.ListNestedAttribute{
													Description:         "Policies.",
													MarkdownDescription: "Policies.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config": schema.MapAttribute{
																Description:         "Config.",
																MarkdownDescription: "Config.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"decision_strategy": schema.StringAttribute{
																Description:         "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",
																MarkdownDescription: "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"description": schema.StringAttribute{
																Description:         "A description for this policy.",
																MarkdownDescription: "A description for this policy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
																Description:         "ID.",
																MarkdownDescription: "ID.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"logic": schema.StringAttribute{
																Description:         "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
																MarkdownDescription: "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of this policy.",
																MarkdownDescription: "The name of this policy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"owner": schema.StringAttribute{
																Description:         "Owner.",
																MarkdownDescription: "Owner.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"policies": schema.ListAttribute{
																Description:         "Policies.",
																MarkdownDescription: "Policies.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources": schema.ListAttribute{
																Description:         "Resources.",
																MarkdownDescription: "Resources.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources_data": schema.ListNestedAttribute{
																Description:         "Resources Data.",
																MarkdownDescription: "Resources Data.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"_id": schema.StringAttribute{
																			Description:         "ID.",
																			MarkdownDescription: "ID.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"attributes": schema.MapAttribute{
																			Description:         "The attributes associated with the resource.",
																			MarkdownDescription: "The attributes associated with the resource.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"display_name": schema.StringAttribute{
																			Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "An URI pointing to an icon.",
																			MarkdownDescription: "An URI pointing to an icon.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"owner_managed_access": schema.BoolAttribute{
																			Description:         "True if the access to this resource can be managed by the resource owner.",
																			MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"scopes": schema.ListAttribute{
																			Description:         "The scopes associated with this resource.",
																			MarkdownDescription: "The scopes associated with this resource.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																			MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uris": schema.ListAttribute{
																			Description:         "Set of URIs which are protected by resource.",
																			MarkdownDescription: "Set of URIs which are protected by resource.",
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
																Description:         "Scopes.",
																MarkdownDescription: "Scopes.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes_data": schema.ListAttribute{
																Description:         "Scopes Data.",
																MarkdownDescription: "Scopes Data.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type.",
																MarkdownDescription: "Type.",
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
													Description:         "The policy enforcement mode dictates how policies are enforced when evaluating authorization requests. 'Enforcing' means requests are denied by default even when there is no policy associated with a given resource. 'Permissive' means requests are allowed even when there is no policy associated with a given resource. 'Disabled' completely disables the evaluation of policies and allows access to any resource.",
													MarkdownDescription: "The policy enforcement mode dictates how policies are enforced when evaluating authorization requests. 'Enforcing' means requests are denied by default even when there is no policy associated with a given resource. 'Permissive' means requests are allowed even when there is no policy associated with a given resource. 'Disabled' completely disables the evaluation of policies and allows access to any resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resources": schema.ListNestedAttribute{
													Description:         "Resources.",
													MarkdownDescription: "Resources.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"_id": schema.StringAttribute{
																Description:         "ID.",
																MarkdownDescription: "ID.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "The attributes associated with the resource.",
																MarkdownDescription: "The attributes associated with the resource.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"display_name": schema.StringAttribute{
																Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
																Description:         "An URI pointing to an icon.",
																MarkdownDescription: "An URI pointing to an icon.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"owner_managed_access": schema.BoolAttribute{
																Description:         "True if the access to this resource can be managed by the resource owner.",
																MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "The scopes associated with this resource.",
																MarkdownDescription: "The scopes associated with this resource.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uris": schema.ListAttribute{
																Description:         "Set of URIs which are protected by resource.",
																MarkdownDescription: "Set of URIs which are protected by resource.",
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
													Description:         "Authorization Scopes.",
													MarkdownDescription: "Authorization Scopes.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"display_name": schema.StringAttribute{
																Description:         "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",
																MarkdownDescription: "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icon_uri": schema.StringAttribute{
																Description:         "An URI pointing to an icon.",
																MarkdownDescription: "An URI pointing to an icon.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
																Description:         "ID.",
																MarkdownDescription: "ID.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",
																MarkdownDescription: "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"policies": schema.ListNestedAttribute{
																Description:         "Policies.",
																MarkdownDescription: "Policies.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"config": schema.MapAttribute{
																			Description:         "Config.",
																			MarkdownDescription: "Config.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"decision_strategy": schema.StringAttribute{
																			Description:         "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",
																			MarkdownDescription: "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"description": schema.StringAttribute{
																			Description:         "A description for this policy.",
																			MarkdownDescription: "A description for this policy.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"id": schema.StringAttribute{
																			Description:         "ID.",
																			MarkdownDescription: "ID.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"logic": schema.StringAttribute{
																			Description:         "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
																			MarkdownDescription: "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "The name of this policy.",
																			MarkdownDescription: "The name of this policy.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"owner": schema.StringAttribute{
																			Description:         "Owner.",
																			MarkdownDescription: "Owner.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"policies": schema.ListAttribute{
																			Description:         "Policies.",
																			MarkdownDescription: "Policies.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"resources": schema.ListAttribute{
																			Description:         "Resources.",
																			MarkdownDescription: "Resources.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"resources_data": schema.ListNestedAttribute{
																			Description:         "Resources Data.",
																			MarkdownDescription: "Resources Data.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"_id": schema.StringAttribute{
																						Description:         "ID.",
																						MarkdownDescription: "ID.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"attributes": schema.MapAttribute{
																						Description:         "The attributes associated with the resource.",
																						MarkdownDescription: "The attributes associated with the resource.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"display_name": schema.StringAttribute{
																						Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																						MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"icon_uri": schema.StringAttribute{
																						Description:         "An URI pointing to an icon.",
																						MarkdownDescription: "An URI pointing to an icon.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																						MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"owner_managed_access": schema.BoolAttribute{
																						Description:         "True if the access to this resource can be managed by the resource owner.",
																						MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"scopes": schema.ListAttribute{
																						Description:         "The scopes associated with this resource.",
																						MarkdownDescription: "The scopes associated with this resource.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"type": schema.StringAttribute{
																						Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																						MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"uris": schema.ListAttribute{
																						Description:         "Set of URIs which are protected by resource.",
																						MarkdownDescription: "Set of URIs which are protected by resource.",
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
																			Description:         "Scopes.",
																			MarkdownDescription: "Scopes.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"scopes_data": schema.ListAttribute{
																			Description:         "Scopes Data.",
																			MarkdownDescription: "Scopes Data.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type.",
																			MarkdownDescription: "Type.",
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

															"resources": schema.ListNestedAttribute{
																Description:         "Resources.",
																MarkdownDescription: "Resources.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"_id": schema.StringAttribute{
																			Description:         "ID.",
																			MarkdownDescription: "ID.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"attributes": schema.MapAttribute{
																			Description:         "The attributes associated with the resource.",
																			MarkdownDescription: "The attributes associated with the resource.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"display_name": schema.StringAttribute{
																			Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"icon_uri": schema.StringAttribute{
																			Description:         "An URI pointing to an icon.",
																			MarkdownDescription: "An URI pointing to an icon.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"owner_managed_access": schema.BoolAttribute{
																			Description:         "True if the access to this resource can be managed by the resource owner.",
																			MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"scopes": schema.ListAttribute{
																			Description:         "The scopes associated with this resource.",
																			MarkdownDescription: "The scopes associated with this resource.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																			MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"uris": schema.ListAttribute{
																			Description:         "Set of URIs which are protected by resource.",
																			MarkdownDescription: "Set of URIs which are protected by resource.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"base_url": schema.StringAttribute{
											Description:         "Application base URL.",
											MarkdownDescription: "Application base URL.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bearer_only": schema.BoolAttribute{
											Description:         "True if a client supports only Bearer Tokens.",
											MarkdownDescription: "True if a client supports only Bearer Tokens.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_authenticator_type": schema.StringAttribute{
											Description:         "What Client authentication type to use.",
											MarkdownDescription: "What Client authentication type to use.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_id": schema.StringAttribute{
											Description:         "Client ID.",
											MarkdownDescription: "Client ID.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"consent_required": schema.BoolAttribute{
											Description:         "True if Consent Screen is required.",
											MarkdownDescription: "True if Consent Screen is required.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_client_scopes": schema.ListAttribute{
											Description:         "A list of default client scopes. Default client scopes are always applied when issuing OpenID Connect tokens or SAML assertions for this client.",
											MarkdownDescription: "A list of default client scopes. Default client scopes are always applied when issuing OpenID Connect tokens or SAML assertions for this client.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"default_roles": schema.ListAttribute{
											Description:         "Default Client roles.",
											MarkdownDescription: "Default Client roles.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"description": schema.StringAttribute{
											Description:         "Client description.",
											MarkdownDescription: "Client description.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"direct_access_grants_enabled": schema.BoolAttribute{
											Description:         "True if Direct Grant is enabled.",
											MarkdownDescription: "True if Direct Grant is enabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "Client enabled flag.",
											MarkdownDescription: "Client enabled flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"frontchannel_logout": schema.BoolAttribute{
											Description:         "True if this client supports Front Channel logout.",
											MarkdownDescription: "True if this client supports Front Channel logout.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"full_scope_allowed": schema.BoolAttribute{
											Description:         "True if Full Scope is allowed.",
											MarkdownDescription: "True if Full Scope is allowed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "Client ID. If not specified, automatically generated.",
											MarkdownDescription: "Client ID. If not specified, automatically generated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"implicit_flow_enabled": schema.BoolAttribute{
											Description:         "True if Implicit flow is enabled.",
											MarkdownDescription: "True if Implicit flow is enabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Client name.",
											MarkdownDescription: "Client name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_re_registration_timeout": schema.Int64Attribute{
											Description:         "Node registration timeout.",
											MarkdownDescription: "Node registration timeout.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_before": schema.Int64Attribute{
											Description:         "Not Before setting.",
											MarkdownDescription: "Not Before setting.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional_client_scopes": schema.ListAttribute{
											Description:         "A list of optional client scopes. Optional client scopes are applied when issuing tokens for this client, but only when they are requested by the scope parameter in the OpenID Connect authorization request.",
											MarkdownDescription: "A list of optional client scopes. Optional client scopes are applied when issuing tokens for this client, but only when they are requested by the scope parameter in the OpenID Connect authorization request.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol used for this Client.",
											MarkdownDescription: "Protocol used for this Client.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_mappers": schema.ListNestedAttribute{
											Description:         "Protocol Mappers.",
											MarkdownDescription: "Protocol Mappers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config": schema.MapAttribute{
														Description:         "Config options.",
														MarkdownDescription: "Config options.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_required": schema.BoolAttribute{
														Description:         "True if Consent Screen is required.",
														MarkdownDescription: "True if Consent Screen is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"consent_text": schema.StringAttribute{
														Description:         "Text to use for displaying Consent Screen.",
														MarkdownDescription: "Text to use for displaying Consent Screen.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "Protocol Mapper ID.",
														MarkdownDescription: "Protocol Mapper ID.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Protocol Mapper Name.",
														MarkdownDescription: "Protocol Mapper Name.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol to use.",
														MarkdownDescription: "Protocol to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol_mapper": schema.StringAttribute{
														Description:         "Protocol Mapper to use",
														MarkdownDescription: "Protocol Mapper to use",
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
											Description:         "True if this is a public Client.",
											MarkdownDescription: "True if this is a public Client.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redirect_uris": schema.ListAttribute{
											Description:         "A list of valid Redirection URLs.",
											MarkdownDescription: "A list of valid Redirection URLs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"root_url": schema.StringAttribute{
											Description:         "Application root URL.",
											MarkdownDescription: "Application root URL.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret": schema.StringAttribute{
											Description:         "Client Secret. The Operator will automatically create a Secret based on this value.",
											MarkdownDescription: "Client Secret. The Operator will automatically create a Secret based on this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts_enabled": schema.BoolAttribute{
											Description:         "True if Service Accounts are enabled.",
											MarkdownDescription: "True if Service Accounts are enabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"standard_flow_enabled": schema.BoolAttribute{
											Description:         "True if Standard flow is enabled.",
											MarkdownDescription: "True if Standard flow is enabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"surrogate_auth_required": schema.BoolAttribute{
											Description:         "Surrogate Authentication Required option.",
											MarkdownDescription: "Surrogate Authentication Required option.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_config": schema.BoolAttribute{
											Description:         "True to use a Template Config.",
											MarkdownDescription: "True to use a Template Config.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_mappers": schema.BoolAttribute{
											Description:         "True to use Template Mappers.",
											MarkdownDescription: "True to use Template Mappers.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_template_scope": schema.BoolAttribute{
											Description:         "True to use Template Scope.",
											MarkdownDescription: "True to use Template Scope.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"web_origins": schema.ListAttribute{
											Description:         "A list of valid Web Origins.",
											MarkdownDescription: "A list of valid Web Origins.",
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

							"default_default_client_scopes": schema.ListAttribute{
								Description:         "Default client scopes to add to all new clients",
								MarkdownDescription: "Default client scopes to add to all new clients",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_locale": schema.StringAttribute{
								Description:         "Default Locale",
								MarkdownDescription: "Default Locale",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_role": schema.SingleNestedAttribute{
								Description:         "Default role",
								MarkdownDescription: "Default role",
								Attributes: map[string]schema.Attribute{
									"attributes": schema.MapAttribute{
										Description:         "Role Attributes",
										MarkdownDescription: "Role Attributes",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_role": schema.BoolAttribute{
										Description:         "Client Role",
										MarkdownDescription: "Client Role",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"composite": schema.BoolAttribute{
										Description:         "Composite",
										MarkdownDescription: "Composite",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"composites": schema.SingleNestedAttribute{
										Description:         "Composites",
										MarkdownDescription: "Composites",
										Attributes: map[string]schema.Attribute{
											"client": schema.MapAttribute{
												Description:         "Map client => []role",
												MarkdownDescription: "Map client => []role",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"realm": schema.ListAttribute{
												Description:         "Realm roles",
												MarkdownDescription: "Realm roles",
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
										Description:         "Container Id",
										MarkdownDescription: "Container Id",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"description": schema.StringAttribute{
										Description:         "Description",
										MarkdownDescription: "Description",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"id": schema.StringAttribute{
										Description:         "Id",
										MarkdownDescription: "Id",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name",
										MarkdownDescription: "Name",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"direct_grant_flow": schema.StringAttribute{
								Description:         "Direct Grant authentication flow",
								MarkdownDescription: "Direct Grant authentication flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"display_name": schema.StringAttribute{
								Description:         "Realm display name.",
								MarkdownDescription: "Realm display name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"display_name_html": schema.StringAttribute{
								Description:         "Realm HTML display name.",
								MarkdownDescription: "Realm HTML display name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"docker_authentication_flow": schema.StringAttribute{
								Description:         "Docker Authentication flow",
								MarkdownDescription: "Docker Authentication flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"duplicate_emails_allowed": schema.BoolAttribute{
								Description:         "Duplicate emails",
								MarkdownDescription: "Duplicate emails",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"edit_username_allowed": schema.BoolAttribute{
								Description:         "Edit username",
								MarkdownDescription: "Edit username",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"email_theme": schema.StringAttribute{
								Description:         "Email Theme",
								MarkdownDescription: "Email Theme",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Realm enabled flag.",
								MarkdownDescription: "Realm enabled flag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled_event_types": schema.ListAttribute{
								Description:         "Enabled event types",
								MarkdownDescription: "Enabled event types",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"events_enabled": schema.BoolAttribute{
								Description:         "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",
								MarkdownDescription: "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"events_listeners": schema.ListAttribute{
								Description:         "A set of Event Listeners.",
								MarkdownDescription: "A set of Event Listeners.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"failure_factor": schema.Int64Attribute{
								Description:         "Max Login Failures",
								MarkdownDescription: "Max Login Failures",
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

							"identity_provider_mappers": schema.ListNestedAttribute{
								Description:         "A set of Identity Provider Mappers.",
								MarkdownDescription: "A set of Identity Provider Mappers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "Identity Provider Mapper config.",
											MarkdownDescription: "Identity Provider Mapper config.",
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
											Description:         "Identity Provider Alias.",
											MarkdownDescription: "Identity Provider Alias.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"identity_provider_mapper": schema.StringAttribute{
											Description:         "Identity Provider Mapper.",
											MarkdownDescription: "Identity Provider Mapper.",
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
								Description:         "A set of Identity Providers.",
								MarkdownDescription: "A set of Identity Providers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"add_read_token_role_on_create": schema.BoolAttribute{
											Description:         "Adds Read Token role when creating this Identity Provider.",
											MarkdownDescription: "Adds Read Token role when creating this Identity Provider.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"alias": schema.StringAttribute{
											Description:         "Identity Provider Alias.",
											MarkdownDescription: "Identity Provider Alias.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"config": schema.MapAttribute{
											Description:         "Identity Provider config.",
											MarkdownDescription: "Identity Provider config.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"display_name": schema.StringAttribute{
											Description:         "Identity Provider Display Name.",
											MarkdownDescription: "Identity Provider Display Name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "Identity Provider enabled flag.",
											MarkdownDescription: "Identity Provider enabled flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"first_broker_login_flow_alias": schema.StringAttribute{
											Description:         "Identity Provider First Broker Login Flow Alias.",
											MarkdownDescription: "Identity Provider First Broker Login Flow Alias.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"internal_id": schema.StringAttribute{
											Description:         "Identity Provider Internal ID.",
											MarkdownDescription: "Identity Provider Internal ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"link_only": schema.BoolAttribute{
											Description:         "Identity Provider Link Only setting.",
											MarkdownDescription: "Identity Provider Link Only setting.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"post_broker_login_flow_alias": schema.StringAttribute{
											Description:         "Identity Provider Post Broker Login Flow Alias.",
											MarkdownDescription: "Identity Provider Post Broker Login Flow Alias.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_id": schema.StringAttribute{
											Description:         "Identity Provider ID.",
											MarkdownDescription: "Identity Provider ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"store_token": schema.BoolAttribute{
											Description:         "Identity Provider Store to Token.",
											MarkdownDescription: "Identity Provider Store to Token.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"trust_email": schema.BoolAttribute{
											Description:         "Identity Provider Trust Email.",
											MarkdownDescription: "Identity Provider Trust Email.",
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
								Description:         "Internationalization Enabled",
								MarkdownDescription: "Internationalization Enabled",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"login_theme": schema.StringAttribute{
								Description:         "Login Theme",
								MarkdownDescription: "Login Theme",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"login_with_email_allowed": schema.BoolAttribute{
								Description:         "Login with email",
								MarkdownDescription: "Login with email",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_delta_time_seconds": schema.Int64Attribute{
								Description:         "Failure Reset Time",
								MarkdownDescription: "Failure Reset Time",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_failure_wait_seconds": schema.Int64Attribute{
								Description:         "Max Wait",
								MarkdownDescription: "Max Wait",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"minimum_quick_login_wait_seconds": schema.Int64Attribute{
								Description:         "Minimum Quick Login Wait",
								MarkdownDescription: "Minimum Quick Login Wait",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_algorithm": schema.StringAttribute{
								Description:         "OTP Policy Algorithm",
								MarkdownDescription: "OTP Policy Algorithm",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_digits": schema.Int64Attribute{
								Description:         "OTP Policy Digits",
								MarkdownDescription: "OTP Policy Digits",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_initial_counter": schema.Int64Attribute{
								Description:         "OTP Policy Initial Counter",
								MarkdownDescription: "OTP Policy Initial Counter",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_look_ahead_window": schema.Int64Attribute{
								Description:         "OTP Policy Look Ahead Window",
								MarkdownDescription: "OTP Policy Look Ahead Window",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_period": schema.Int64Attribute{
								Description:         "OTP Policy Period",
								MarkdownDescription: "OTP Policy Period",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_policy_type": schema.StringAttribute{
								Description:         "OTP Policy Type",
								MarkdownDescription: "OTP Policy Type",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"otp_supported_applications": schema.ListAttribute{
								Description:         "OTP Supported Applications",
								MarkdownDescription: "OTP Supported Applications",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password_policy": schema.StringAttribute{
								Description:         "Realm Password Policy",
								MarkdownDescription: "Realm Password Policy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"permanent_lockout": schema.BoolAttribute{
								Description:         "Permanent Lockout",
								MarkdownDescription: "Permanent Lockout",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"quick_login_check_milli_seconds": schema.Int64Attribute{
								Description:         "Quick Login Check Milli Seconds",
								MarkdownDescription: "Quick Login Check Milli Seconds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm": schema.StringAttribute{
								Description:         "Realm name.",
								MarkdownDescription: "Realm name.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"registration_allowed": schema.BoolAttribute{
								Description:         "User registration",
								MarkdownDescription: "User registration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_email_as_username": schema.BoolAttribute{
								Description:         "Email as username",
								MarkdownDescription: "Email as username",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_flow": schema.StringAttribute{
								Description:         "Registration flow",
								MarkdownDescription: "Registration flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remember_me": schema.BoolAttribute{
								Description:         "Remember me",
								MarkdownDescription: "Remember me",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reset_credentials_flow": schema.StringAttribute{
								Description:         "Reset Credentials authentication flow",
								MarkdownDescription: "Reset Credentials authentication flow",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reset_password_allowed": schema.BoolAttribute{
								Description:         "Forgot password",
								MarkdownDescription: "Forgot password",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"roles": schema.SingleNestedAttribute{
								Description:         "Roles",
								MarkdownDescription: "Roles",
								Attributes: map[string]schema.Attribute{
									"client": schema.MapAttribute{
										Description:         "Client Roles",
										MarkdownDescription: "Client Roles",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"realm": schema.ListNestedAttribute{
										Description:         "Realm Roles",
										MarkdownDescription: "Realm Roles",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.MapAttribute{
													Description:         "Role Attributes",
													MarkdownDescription: "Role Attributes",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_role": schema.BoolAttribute{
													Description:         "Client Role",
													MarkdownDescription: "Client Role",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"composite": schema.BoolAttribute{
													Description:         "Composite",
													MarkdownDescription: "Composite",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"composites": schema.SingleNestedAttribute{
													Description:         "Composites",
													MarkdownDescription: "Composites",
													Attributes: map[string]schema.Attribute{
														"client": schema.MapAttribute{
															Description:         "Map client => []role",
															MarkdownDescription: "Map client => []role",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"realm": schema.ListAttribute{
															Description:         "Realm roles",
															MarkdownDescription: "Realm roles",
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
													Description:         "Container Id",
													MarkdownDescription: "Container Id",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "Description",
													MarkdownDescription: "Description",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
													Description:         "Id",
													MarkdownDescription: "Id",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name",
													MarkdownDescription: "Name",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"scope_mappings": schema.ListNestedAttribute{
								Description:         "Scope Mappings",
								MarkdownDescription: "Scope Mappings",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"client": schema.StringAttribute{
											Description:         "Client",
											MarkdownDescription: "Client",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_scope": schema.StringAttribute{
											Description:         "Client Scope",
											MarkdownDescription: "Client Scope",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles",
											MarkdownDescription: "Roles",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"self": schema.StringAttribute{
											Description:         "Self",
											MarkdownDescription: "Self",
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
								Description:         "Email",
								MarkdownDescription: "Email",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_required": schema.StringAttribute{
								Description:         "Require SSL",
								MarkdownDescription: "Require SSL",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supported_locales": schema.ListAttribute{
								Description:         "Supported Locales",
								MarkdownDescription: "Supported Locales",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_federation_mappers": schema.ListNestedAttribute{
								Description:         "User federation mappers are extension points triggered by the user federation at various points.",
								MarkdownDescription: "User federation mappers are extension points triggered by the user federation at various points.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.MapAttribute{
											Description:         "User federation mapper config.",
											MarkdownDescription: "User federation mapper config.",
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
											Description:         "The displayName for the user federation provider this mapper applies to.",
											MarkdownDescription: "The displayName for the user federation provider this mapper applies to.",
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
								Description:         "Point keycloak to an external user provider to validate credentials or pull in identity information.",
								MarkdownDescription: "Point keycloak to an external user provider to validate credentials or pull in identity information.",
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
											Description:         "User federation provider config.",
											MarkdownDescription: "User federation provider config.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"display_name": schema.StringAttribute{
											Description:         "The display name of this provider instance.",
											MarkdownDescription: "The display name of this provider instance.",
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
											Description:         "The ID of this provider",
											MarkdownDescription: "The ID of this provider",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"priority": schema.Int64Attribute{
											Description:         "The priority of this provider when looking up users or adding a user.",
											MarkdownDescription: "The priority of this provider when looking up users or adding a user.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provider_name": schema.StringAttribute{
											Description:         "The name of the user provider, such as 'ldap', 'kerberos' or a custom SPI.",
											MarkdownDescription: "The name of the user provider, such as 'ldap', 'kerberos' or a custom SPI.",
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
								Description:         "User Managed Access Allowed",
								MarkdownDescription: "User Managed Access Allowed",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"users": schema.ListNestedAttribute{
								Description:         "A set of Keycloak Users.",
								MarkdownDescription: "A set of Keycloak Users.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attributes": schema.MapAttribute{
											Description:         "A set of Attributes.",
											MarkdownDescription: "A set of Attributes.",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_roles": schema.MapAttribute{
											Description:         "A set of Client Roles.",
											MarkdownDescription: "A set of Client Roles.",
											ElementType:         types.ListType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credentials": schema.ListNestedAttribute{
											Description:         "A set of Credentials.",
											MarkdownDescription: "A set of Credentials.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"temporary": schema.BoolAttribute{
														Description:         "True if this credential object is temporary.",
														MarkdownDescription: "True if this credential object is temporary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Credential Type.",
														MarkdownDescription: "Credential Type.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Credential Value.",
														MarkdownDescription: "Credential Value.",
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

										"email": schema.StringAttribute{
											Description:         "Email.",
											MarkdownDescription: "Email.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"email_verified": schema.BoolAttribute{
											Description:         "True if email has already been verified.",
											MarkdownDescription: "True if email has already been verified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "User enabled flag.",
											MarkdownDescription: "User enabled flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"federated_identities": schema.ListNestedAttribute{
											Description:         "A set of Federated Identities.",
											MarkdownDescription: "A set of Federated Identities.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"identity_provider": schema.StringAttribute{
														Description:         "Federated Identity Provider.",
														MarkdownDescription: "Federated Identity Provider.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_id": schema.StringAttribute{
														Description:         "Federated Identity User ID.",
														MarkdownDescription: "Federated Identity User ID.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_name": schema.StringAttribute{
														Description:         "Federated Identity User Name.",
														MarkdownDescription: "Federated Identity User Name.",
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

										"first_name": schema.StringAttribute{
											Description:         "First Name.",
											MarkdownDescription: "First Name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"groups": schema.ListAttribute{
											Description:         "A set of Groups.",
											MarkdownDescription: "A set of Groups.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "User ID.",
											MarkdownDescription: "User ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"last_name": schema.StringAttribute{
											Description:         "Last Name.",
											MarkdownDescription: "Last Name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"realm_roles": schema.ListAttribute{
											Description:         "A set of Realm Roles.",
											MarkdownDescription: "A set of Realm Roles.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"required_actions": schema.ListAttribute{
											Description:         "A set of Required Actions.",
											MarkdownDescription: "A set of Required Actions.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"username": schema.StringAttribute{
											Description:         "User Name.",
											MarkdownDescription: "User Name.",
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
								Description:         "Verify email",
								MarkdownDescription: "Verify email",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_increment_seconds": schema.Int64Attribute{
								Description:         "Wait Increment",
								MarkdownDescription: "Wait Increment",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"realm_overrides": schema.ListNestedAttribute{
						Description:         "A list of overrides to the default Realm behavior.",
						MarkdownDescription: "A list of overrides to the default Realm behavior.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"for_flow": schema.StringAttribute{
									Description:         "Flow to be overridden.",
									MarkdownDescription: "Flow to be overridden.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"identity_provider": schema.StringAttribute{
									Description:         "Identity Provider to be overridden.",
									MarkdownDescription: "Identity Provider to be overridden.",
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

					"unmanaged": schema.BoolAttribute{
						Description:         "When set to true, this KeycloakRealm will be marked as unmanaged and not be managed by this operator. It can then be used for targeting purposes.",
						MarkdownDescription: "When set to true, this KeycloakRealm will be marked as unmanaged and not be managed by this operator. It can then be used for targeting purposes.",
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
	}
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_realm_v1alpha1_manifest")

	var model KeycloakOrgKeycloakRealmV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("keycloak.org/v1alpha1")
	model.Kind = pointer.String("KeycloakRealm")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
