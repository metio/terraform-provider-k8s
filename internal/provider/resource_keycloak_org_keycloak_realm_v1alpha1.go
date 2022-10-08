/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
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

type KeycloakOrgKeycloakRealmV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*KeycloakOrgKeycloakRealmV1Alpha1Resource)(nil)
)

type KeycloakOrgKeycloakRealmV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KeycloakOrgKeycloakRealmV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		InstanceSelector *struct {
			MatchExpressions *[]struct {
				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"instance_selector" yaml:"instanceSelector,omitempty"`

		Realm *struct {
			Realm *string `tfsdk:"realm" yaml:"realm,omitempty"`

			AccessTokenLifespan *int64 `tfsdk:"access_token_lifespan" yaml:"accessTokenLifespan,omitempty"`

			AuthenticationFlows *[]struct {
				Alias *string `tfsdk:"alias" yaml:"alias,omitempty"`

				AuthenticationExecutions *[]struct {
					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					Requirement *string `tfsdk:"requirement" yaml:"requirement,omitempty"`

					UserSetupAllowed *bool `tfsdk:"user_setup_allowed" yaml:"userSetupAllowed,omitempty"`

					Authenticator *string `tfsdk:"authenticator" yaml:"authenticator,omitempty"`

					AuthenticatorConfig *string `tfsdk:"authenticator_config" yaml:"authenticatorConfig,omitempty"`

					AuthenticatorFlow *bool `tfsdk:"authenticator_flow" yaml:"authenticatorFlow,omitempty"`

					FlowAlias *string `tfsdk:"flow_alias" yaml:"flowAlias,omitempty"`
				} `tfsdk:"authentication_executions" yaml:"authenticationExecutions,omitempty"`

				BuiltIn *bool `tfsdk:"built_in" yaml:"builtIn,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				ProviderId *string `tfsdk:"provider_id" yaml:"providerId,omitempty"`

				TopLevel *bool `tfsdk:"top_level" yaml:"topLevel,omitempty"`
			} `tfsdk:"authentication_flows" yaml:"authenticationFlows,omitempty"`

			DefaultRole *struct {
				ClientRole *bool `tfsdk:"client_role" yaml:"clientRole,omitempty"`

				Composite *bool `tfsdk:"composite" yaml:"composite,omitempty"`

				Composites *struct {
					Client *map[string][]string `tfsdk:"client" yaml:"client,omitempty"`

					Realm *[]string `tfsdk:"realm" yaml:"realm,omitempty"`
				} `tfsdk:"composites" yaml:"composites,omitempty"`

				ContainerId *string `tfsdk:"container_id" yaml:"containerId,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Attributes *map[string][]string `tfsdk:"attributes" yaml:"attributes,omitempty"`
			} `tfsdk:"default_role" yaml:"defaultRole,omitempty"`

			EventsEnabled *bool `tfsdk:"events_enabled" yaml:"eventsEnabled,omitempty"`

			FailureFactor *int64 `tfsdk:"failure_factor" yaml:"failureFactor,omitempty"`

			MaxDeltaTimeSeconds *int64 `tfsdk:"max_delta_time_seconds" yaml:"maxDeltaTimeSeconds,omitempty"`

			QuickLoginCheckMilliSeconds *int64 `tfsdk:"quick_login_check_milli_seconds" yaml:"quickLoginCheckMilliSeconds,omitempty"`

			RegistrationAllowed *bool `tfsdk:"registration_allowed" yaml:"registrationAllowed,omitempty"`

			UserManagedAccessAllowed *bool `tfsdk:"user_managed_access_allowed" yaml:"userManagedAccessAllowed,omitempty"`

			DisplayNameHtml *string `tfsdk:"display_name_html" yaml:"displayNameHtml,omitempty"`

			PermanentLockout *bool `tfsdk:"permanent_lockout" yaml:"permanentLockout,omitempty"`

			WaitIncrementSeconds *int64 `tfsdk:"wait_increment_seconds" yaml:"waitIncrementSeconds,omitempty"`

			AdminEventsEnabled *bool `tfsdk:"admin_events_enabled" yaml:"adminEventsEnabled,omitempty"`

			DuplicateEmailsAllowed *bool `tfsdk:"duplicate_emails_allowed" yaml:"duplicateEmailsAllowed,omitempty"`

			MinimumQuickLoginWaitSeconds *int64 `tfsdk:"minimum_quick_login_wait_seconds" yaml:"minimumQuickLoginWaitSeconds,omitempty"`

			OtpPolicyPeriod *int64 `tfsdk:"otp_policy_period" yaml:"otpPolicyPeriod,omitempty"`

			OtpPolicyType *string `tfsdk:"otp_policy_type" yaml:"otpPolicyType,omitempty"`

			AuthenticatorConfig *[]struct {
				Alias *string `tfsdk:"alias" yaml:"alias,omitempty"`

				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`
			} `tfsdk:"authenticator_config" yaml:"authenticatorConfig,omitempty"`

			ResetPasswordAllowed *bool `tfsdk:"reset_password_allowed" yaml:"resetPasswordAllowed,omitempty"`

			Users *[]struct {
				Credentials *[]struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					Temporary *bool `tfsdk:"temporary" yaml:"temporary,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"credentials" yaml:"credentials,omitempty"`

				Email *string `tfsdk:"email" yaml:"email,omitempty"`

				FederatedIdentities *[]struct {
					IdentityProvider *string `tfsdk:"identity_provider" yaml:"identityProvider,omitempty"`

					UserId *string `tfsdk:"user_id" yaml:"userId,omitempty"`

					UserName *string `tfsdk:"user_name" yaml:"userName,omitempty"`
				} `tfsdk:"federated_identities" yaml:"federatedIdentities,omitempty"`

				FirstName *string `tfsdk:"first_name" yaml:"firstName,omitempty"`

				Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

				RequiredActions *[]string `tfsdk:"required_actions" yaml:"requiredActions,omitempty"`

				Attributes *map[string][]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

				ClientRoles *map[string][]string `tfsdk:"client_roles" yaml:"clientRoles,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				LastName *string `tfsdk:"last_name" yaml:"lastName,omitempty"`

				RealmRoles *[]string `tfsdk:"realm_roles" yaml:"realmRoles,omitempty"`

				EmailVerified *bool `tfsdk:"email_verified" yaml:"emailVerified,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Username *string `tfsdk:"username" yaml:"username,omitempty"`
			} `tfsdk:"users" yaml:"users,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			LoginTheme *string `tfsdk:"login_theme" yaml:"loginTheme,omitempty"`

			MaxFailureWaitSeconds *int64 `tfsdk:"max_failure_wait_seconds" yaml:"maxFailureWaitSeconds,omitempty"`

			OtpPolicyDigits *int64 `tfsdk:"otp_policy_digits" yaml:"otpPolicyDigits,omitempty"`

			Roles *struct {
				Client *map[string]string `tfsdk:"client" yaml:"client,omitempty"`

				Realm *[]struct {
					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Attributes *map[string][]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

					ClientRole *bool `tfsdk:"client_role" yaml:"clientRole,omitempty"`

					Composite *bool `tfsdk:"composite" yaml:"composite,omitempty"`

					Composites *struct {
						Client *map[string][]string `tfsdk:"client" yaml:"client,omitempty"`

						Realm *[]string `tfsdk:"realm" yaml:"realm,omitempty"`
					} `tfsdk:"composites" yaml:"composites,omitempty"`

					ContainerId *string `tfsdk:"container_id" yaml:"containerId,omitempty"`

					Description *string `tfsdk:"description" yaml:"description,omitempty"`
				} `tfsdk:"realm" yaml:"realm,omitempty"`
			} `tfsdk:"roles" yaml:"roles,omitempty"`

			SslRequired *string `tfsdk:"ssl_required" yaml:"sslRequired,omitempty"`

			UserFederationProviders *[]struct {
				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

				ProviderName *string `tfsdk:"provider_name" yaml:"providerName,omitempty"`

				ChangedSyncPeriod *int64 `tfsdk:"changed_sync_period" yaml:"changedSyncPeriod,omitempty"`

				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

				FullSyncPeriod *int64 `tfsdk:"full_sync_period" yaml:"fullSyncPeriod,omitempty"`
			} `tfsdk:"user_federation_providers" yaml:"userFederationProviders,omitempty"`

			BruteForceProtected *bool `tfsdk:"brute_force_protected" yaml:"bruteForceProtected,omitempty"`

			OtpSupportedApplications *[]string `tfsdk:"otp_supported_applications" yaml:"otpSupportedApplications,omitempty"`

			AccessTokenLifespanForImplicitFlow *int64 `tfsdk:"access_token_lifespan_for_implicit_flow" yaml:"accessTokenLifespanForImplicitFlow,omitempty"`

			AdminEventsDetailsEnabled *bool `tfsdk:"admin_events_details_enabled" yaml:"adminEventsDetailsEnabled,omitempty"`

			OtpPolicyLookAheadWindow *int64 `tfsdk:"otp_policy_look_ahead_window" yaml:"otpPolicyLookAheadWindow,omitempty"`

			RegistrationEmailAsUsername *bool `tfsdk:"registration_email_as_username" yaml:"registrationEmailAsUsername,omitempty"`

			VerifyEmail *bool `tfsdk:"verify_email" yaml:"verifyEmail,omitempty"`

			PasswordPolicy *string `tfsdk:"password_policy" yaml:"passwordPolicy,omitempty"`

			ClientScopeMappings *map[string]string `tfsdk:"client_scope_mappings" yaml:"clientScopeMappings,omitempty"`

			ClientScopes *[]struct {
				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

				ProtocolMappers *[]struct {
					ConsentText *string `tfsdk:"consent_text" yaml:"consentText,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

					ProtocolMapper *string `tfsdk:"protocol_mapper" yaml:"protocolMapper,omitempty"`

					Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

					ConsentRequired *bool `tfsdk:"consent_required" yaml:"consentRequired,omitempty"`
				} `tfsdk:"protocol_mappers" yaml:"protocolMappers,omitempty"`

				Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"client_scopes" yaml:"clientScopes,omitempty"`

			DefaultDefaultClientScopes *[]string `tfsdk:"default_default_client_scopes" yaml:"defaultDefaultClientScopes,omitempty"`

			EmailTheme *string `tfsdk:"email_theme" yaml:"emailTheme,omitempty"`

			LoginWithEmailAllowed *bool `tfsdk:"login_with_email_allowed" yaml:"loginWithEmailAllowed,omitempty"`

			ScopeMappings *[]struct {
				Client *string `tfsdk:"client" yaml:"client,omitempty"`

				ClientScope *string `tfsdk:"client_scope" yaml:"clientScope,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				Self *string `tfsdk:"self" yaml:"self,omitempty"`
			} `tfsdk:"scope_mappings" yaml:"scopeMappings,omitempty"`

			UserFederationMappers *[]struct {
				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				FederationMapperType *string `tfsdk:"federation_mapper_type" yaml:"federationMapperType,omitempty"`

				FederationProviderDisplayName *string `tfsdk:"federation_provider_display_name" yaml:"federationProviderDisplayName,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"user_federation_mappers" yaml:"userFederationMappers,omitempty"`

			DefaultLocale *string `tfsdk:"default_locale" yaml:"defaultLocale,omitempty"`

			DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

			EnabledEventTypes *[]string `tfsdk:"enabled_event_types" yaml:"enabledEventTypes,omitempty"`

			OtpPolicyAlgorithm *string `tfsdk:"otp_policy_algorithm" yaml:"otpPolicyAlgorithm,omitempty"`

			RememberMe *bool `tfsdk:"remember_me" yaml:"rememberMe,omitempty"`

			AdminTheme *string `tfsdk:"admin_theme" yaml:"adminTheme,omitempty"`

			EditUsernameAllowed *bool `tfsdk:"edit_username_allowed" yaml:"editUsernameAllowed,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			SmtpServer *map[string]string `tfsdk:"smtp_server" yaml:"smtpServer,omitempty"`

			AccountTheme *string `tfsdk:"account_theme" yaml:"accountTheme,omitempty"`

			Clients *[]struct {
				ConsentRequired *bool `tfsdk:"consent_required" yaml:"consentRequired,omitempty"`

				FrontchannelLogout *bool `tfsdk:"frontchannel_logout" yaml:"frontchannelLogout,omitempty"`

				ServiceAccountsEnabled *bool `tfsdk:"service_accounts_enabled" yaml:"serviceAccountsEnabled,omitempty"`

				AuthenticationFlowBindingOverrides *map[string]string `tfsdk:"authentication_flow_binding_overrides" yaml:"authenticationFlowBindingOverrides,omitempty"`

				BaseUrl *string `tfsdk:"base_url" yaml:"baseUrl,omitempty"`

				ClientAuthenticatorType *string `tfsdk:"client_authenticator_type" yaml:"clientAuthenticatorType,omitempty"`

				OptionalClientScopes *[]string `tfsdk:"optional_client_scopes" yaml:"optionalClientScopes,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

				BearerOnly *bool `tfsdk:"bearer_only" yaml:"bearerOnly,omitempty"`

				RedirectUris *[]string `tfsdk:"redirect_uris" yaml:"redirectUris,omitempty"`

				RootUrl *string `tfsdk:"root_url" yaml:"rootUrl,omitempty"`

				StandardFlowEnabled *bool `tfsdk:"standard_flow_enabled" yaml:"standardFlowEnabled,omitempty"`

				SurrogateAuthRequired *bool `tfsdk:"surrogate_auth_required" yaml:"surrogateAuthRequired,omitempty"`

				UseTemplateScope *bool `tfsdk:"use_template_scope" yaml:"useTemplateScope,omitempty"`

				WebOrigins *[]string `tfsdk:"web_origins" yaml:"webOrigins,omitempty"`

				AdminUrl *string `tfsdk:"admin_url" yaml:"adminUrl,omitempty"`

				AuthorizationServicesEnabled *bool `tfsdk:"authorization_services_enabled" yaml:"authorizationServicesEnabled,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ImplicitFlowEnabled *bool `tfsdk:"implicit_flow_enabled" yaml:"implicitFlowEnabled,omitempty"`

				UseTemplateConfig *bool `tfsdk:"use_template_config" yaml:"useTemplateConfig,omitempty"`

				ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				DirectAccessGrantsEnabled *bool `tfsdk:"direct_access_grants_enabled" yaml:"directAccessGrantsEnabled,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				NodeReRegistrationTimeout *int64 `tfsdk:"node_re_registration_timeout" yaml:"nodeReRegistrationTimeout,omitempty"`

				AuthorizationSettings *struct {
					AllowRemoteResourceManagement *bool `tfsdk:"allow_remote_resource_management" yaml:"allowRemoteResourceManagement,omitempty"`

					DecisionStrategy *string `tfsdk:"decision_strategy" yaml:"decisionStrategy,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Resources *[]struct {
						_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

						Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

						Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

						DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					Scopes *[]struct {
						DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

						IconUri *string `tfsdk:"icon_uri" yaml:"iconUri,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Policies *[]struct {
							DecisionStrategy *string `tfsdk:"decision_strategy" yaml:"decisionStrategy,omitempty"`

							Id *string `tfsdk:"id" yaml:"id,omitempty"`

							Logic *string `tfsdk:"logic" yaml:"logic,omitempty"`

							Policies *[]string `tfsdk:"policies" yaml:"policies,omitempty"`

							ResourcesData *[]struct {
								Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`

								Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

								DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`

								_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

								Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

								OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`
							} `tfsdk:"resources_data" yaml:"resourcesData,omitempty"`

							Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

							ScopesData *[]string `tfsdk:"scopes_data" yaml:"scopesData,omitempty"`

							Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

							Description *string `tfsdk:"description" yaml:"description,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

							Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`
						} `tfsdk:"policies" yaml:"policies,omitempty"`

						Resources *[]struct {
							_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

							DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

							OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`

							Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

							Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`
					} `tfsdk:"scopes" yaml:"scopes,omitempty"`

					ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

					Policies *[]struct {
						DecisionStrategy *string `tfsdk:"decision_strategy" yaml:"decisionStrategy,omitempty"`

						Description *string `tfsdk:"description" yaml:"description,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Policies *[]string `tfsdk:"policies" yaml:"policies,omitempty"`

						ResourcesData *[]struct {
							_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

							Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

							OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`

							Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

							DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`
						} `tfsdk:"resources_data" yaml:"resourcesData,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

						Logic *string `tfsdk:"logic" yaml:"logic,omitempty"`

						Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

						Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

						ScopesData *[]string `tfsdk:"scopes_data" yaml:"scopesData,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"policies" yaml:"policies,omitempty"`

					PolicyEnforcementMode *string `tfsdk:"policy_enforcement_mode" yaml:"policyEnforcementMode,omitempty"`
				} `tfsdk:"authorization_settings" yaml:"authorizationSettings,omitempty"`

				Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

				UseTemplateMappers *bool `tfsdk:"use_template_mappers" yaml:"useTemplateMappers,omitempty"`

				DefaultClientScopes *[]string `tfsdk:"default_client_scopes" yaml:"defaultClientScopes,omitempty"`

				DefaultRoles *[]string `tfsdk:"default_roles" yaml:"defaultRoles,omitempty"`

				FullScopeAllowed *bool `tfsdk:"full_scope_allowed" yaml:"fullScopeAllowed,omitempty"`

				NotBefore *int64 `tfsdk:"not_before" yaml:"notBefore,omitempty"`

				ProtocolMappers *[]struct {
					Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

					ConsentRequired *bool `tfsdk:"consent_required" yaml:"consentRequired,omitempty"`

					ConsentText *string `tfsdk:"consent_text" yaml:"consentText,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

					ProtocolMapper *string `tfsdk:"protocol_mapper" yaml:"protocolMapper,omitempty"`
				} `tfsdk:"protocol_mappers" yaml:"protocolMappers,omitempty"`

				Access *map[string]string `tfsdk:"access" yaml:"access,omitempty"`

				Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				PublicClient *bool `tfsdk:"public_client" yaml:"publicClient,omitempty"`
			} `tfsdk:"clients" yaml:"clients,omitempty"`

			EventsListeners *[]string `tfsdk:"events_listeners" yaml:"eventsListeners,omitempty"`

			IdentityProviders *[]struct {
				DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				FirstBrokerLoginFlowAlias *string `tfsdk:"first_broker_login_flow_alias" yaml:"firstBrokerLoginFlowAlias,omitempty"`

				InternalId *string `tfsdk:"internal_id" yaml:"internalId,omitempty"`

				LinkOnly *bool `tfsdk:"link_only" yaml:"linkOnly,omitempty"`

				StoreToken *bool `tfsdk:"store_token" yaml:"storeToken,omitempty"`

				Alias *string `tfsdk:"alias" yaml:"alias,omitempty"`

				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				PostBrokerLoginFlowAlias *string `tfsdk:"post_broker_login_flow_alias" yaml:"postBrokerLoginFlowAlias,omitempty"`

				ProviderId *string `tfsdk:"provider_id" yaml:"providerId,omitempty"`

				TrustEmail *bool `tfsdk:"trust_email" yaml:"trustEmail,omitempty"`

				AddReadTokenRoleOnCreate *bool `tfsdk:"add_read_token_role_on_create" yaml:"addReadTokenRoleOnCreate,omitempty"`
			} `tfsdk:"identity_providers" yaml:"identityProviders,omitempty"`

			InternationalizationEnabled *bool `tfsdk:"internationalization_enabled" yaml:"internationalizationEnabled,omitempty"`

			OtpPolicyInitialCounter *int64 `tfsdk:"otp_policy_initial_counter" yaml:"otpPolicyInitialCounter,omitempty"`

			SupportedLocales *[]string `tfsdk:"supported_locales" yaml:"supportedLocales,omitempty"`
		} `tfsdk:"realm" yaml:"realm,omitempty"`

		RealmOverrides *[]struct {
			ForFlow *string `tfsdk:"for_flow" yaml:"forFlow,omitempty"`

			IdentityProvider *string `tfsdk:"identity_provider" yaml:"identityProvider,omitempty"`
		} `tfsdk:"realm_overrides" yaml:"realmOverrides,omitempty"`

		Unmanaged *bool `tfsdk:"unmanaged" yaml:"unmanaged,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKeycloakOrgKeycloakRealmV1Alpha1Resource() resource.Resource {
	return &KeycloakOrgKeycloakRealmV1Alpha1Resource{}
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keycloak_org_keycloak_realm_v1alpha1"
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "KeycloakRealm is the Schema for the keycloakrealms API",
		MarkdownDescription: "KeycloakRealm is the Schema for the keycloakrealms API",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "KeycloakRealmSpec defines the desired state of KeycloakRealm.",
				MarkdownDescription: "KeycloakRealmSpec defines the desired state of KeycloakRealm.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"instance_selector": {
						Description:         "Selector for looking up Keycloak Custom Resources.",
						MarkdownDescription: "Selector for looking up Keycloak Custom Resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

					"realm": {
						Description:         "Keycloak Realm REST object.",
						MarkdownDescription: "Keycloak Realm REST object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"realm": {
								Description:         "Realm name.",
								MarkdownDescription: "Realm name.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"access_token_lifespan": {
								Description:         "Access Token Lifespan",
								MarkdownDescription: "Access Token Lifespan",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"authentication_flows": {
								Description:         "Authentication flows",
								MarkdownDescription: "Authentication flows",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"alias": {
										Description:         "Alias",
										MarkdownDescription: "Alias",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"authentication_executions": {
										Description:         "Authentication executions",
										MarkdownDescription: "Authentication executions",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"priority": {
												Description:         "Priority",
												MarkdownDescription: "Priority",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requirement": {
												Description:         "Requirement [REQUIRED, OPTIONAL, ALTERNATIVE, DISABLED]",
												MarkdownDescription: "Requirement [REQUIRED, OPTIONAL, ALTERNATIVE, DISABLED]",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user_setup_allowed": {
												Description:         "User setup allowed",
												MarkdownDescription: "User setup allowed",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"authenticator": {
												Description:         "Authenticator",
												MarkdownDescription: "Authenticator",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"authenticator_config": {
												Description:         "Authenticator Config",
												MarkdownDescription: "Authenticator Config",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"authenticator_flow": {
												Description:         "Authenticator flow",
												MarkdownDescription: "Authenticator flow",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"flow_alias": {
												Description:         "Flow Alias",
												MarkdownDescription: "Flow Alias",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"built_in": {
										Description:         "Built in",
										MarkdownDescription: "Built in",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "Description",
										MarkdownDescription: "Description",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "ID",
										MarkdownDescription: "ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"provider_id": {
										Description:         "Provider ID",
										MarkdownDescription: "Provider ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"top_level": {
										Description:         "Top level",
										MarkdownDescription: "Top level",

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

							"default_role": {
								Description:         "Default role",
								MarkdownDescription: "Default role",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_role": {
										Description:         "Client Role",
										MarkdownDescription: "Client Role",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"composite": {
										Description:         "Composite",
										MarkdownDescription: "Composite",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"composites": {
										Description:         "Composites",
										MarkdownDescription: "Composites",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"client": {
												Description:         "Map client => []role",
												MarkdownDescription: "Map client => []role",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"realm": {
												Description:         "Realm roles",
												MarkdownDescription: "Realm roles",

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

									"container_id": {
										Description:         "Container Id",
										MarkdownDescription: "Container Id",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "Description",
										MarkdownDescription: "Description",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "Id",
										MarkdownDescription: "Id",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name",
										MarkdownDescription: "Name",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"attributes": {
										Description:         "Role Attributes",
										MarkdownDescription: "Role Attributes",

										Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"events_enabled": {
								Description:         "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",
								MarkdownDescription: "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"failure_factor": {
								Description:         "Max Login Failures",
								MarkdownDescription: "Max Login Failures",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_delta_time_seconds": {
								Description:         "Failure Reset Time",
								MarkdownDescription: "Failure Reset Time",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"quick_login_check_milli_seconds": {
								Description:         "Quick Login Check Milli Seconds",
								MarkdownDescription: "Quick Login Check Milli Seconds",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"registration_allowed": {
								Description:         "User registration",
								MarkdownDescription: "User registration",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_managed_access_allowed": {
								Description:         "User Managed Access Allowed",
								MarkdownDescription: "User Managed Access Allowed",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"display_name_html": {
								Description:         "Realm HTML display name.",
								MarkdownDescription: "Realm HTML display name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"permanent_lockout": {
								Description:         "Permanent Lockout",
								MarkdownDescription: "Permanent Lockout",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wait_increment_seconds": {
								Description:         "Wait Increment",
								MarkdownDescription: "Wait Increment",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_events_enabled": {
								Description:         "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",
								MarkdownDescription: "Enable events recording TODO: change to values and use kubebuilder default annotation once supported",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplicate_emails_allowed": {
								Description:         "Duplicate emails",
								MarkdownDescription: "Duplicate emails",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"minimum_quick_login_wait_seconds": {
								Description:         "Minimum Quick Login Wait",
								MarkdownDescription: "Minimum Quick Login Wait",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_policy_period": {
								Description:         "OTP Policy Period",
								MarkdownDescription: "OTP Policy Period",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_policy_type": {
								Description:         "OTP Policy Type",
								MarkdownDescription: "OTP Policy Type",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"authenticator_config": {
								Description:         "Authenticator config",
								MarkdownDescription: "Authenticator config",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"alias": {
										Description:         "Alias",
										MarkdownDescription: "Alias",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"config": {
										Description:         "Config",
										MarkdownDescription: "Config",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "ID",
										MarkdownDescription: "ID",

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

							"reset_password_allowed": {
								Description:         "Forgot password",
								MarkdownDescription: "Forgot password",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"users": {
								Description:         "A set of Keycloak Users.",
								MarkdownDescription: "A set of Keycloak Users.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"credentials": {
										Description:         "A set of Credentials.",
										MarkdownDescription: "A set of Credentials.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "Credential Value.",
												MarkdownDescription: "Credential Value.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"temporary": {
												Description:         "True if this credential object is temporary.",
												MarkdownDescription: "True if this credential object is temporary.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Credential Type.",
												MarkdownDescription: "Credential Type.",

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

									"email": {
										Description:         "Email.",
										MarkdownDescription: "Email.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"federated_identities": {
										Description:         "A set of Federated Identities.",
										MarkdownDescription: "A set of Federated Identities.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"identity_provider": {
												Description:         "Federated Identity Provider.",
												MarkdownDescription: "Federated Identity Provider.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user_id": {
												Description:         "Federated Identity User ID.",
												MarkdownDescription: "Federated Identity User ID.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user_name": {
												Description:         "Federated Identity User Name.",
												MarkdownDescription: "Federated Identity User Name.",

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

									"first_name": {
										Description:         "First Name.",
										MarkdownDescription: "First Name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"groups": {
										Description:         "A set of Groups.",
										MarkdownDescription: "A set of Groups.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_actions": {
										Description:         "A set of Required Actions.",
										MarkdownDescription: "A set of Required Actions.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"attributes": {
										Description:         "A set of Attributes.",
										MarkdownDescription: "A set of Attributes.",

										Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_roles": {
										Description:         "A set of Client Roles.",
										MarkdownDescription: "A set of Client Roles.",

										Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "User enabled flag.",
										MarkdownDescription: "User enabled flag.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"last_name": {
										Description:         "Last Name.",
										MarkdownDescription: "Last Name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"realm_roles": {
										Description:         "A set of Realm Roles.",
										MarkdownDescription: "A set of Realm Roles.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"email_verified": {
										Description:         "True if email has already been verified.",
										MarkdownDescription: "True if email has already been verified.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "User ID.",
										MarkdownDescription: "User ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"username": {
										Description:         "User Name.",
										MarkdownDescription: "User Name.",

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

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"login_theme": {
								Description:         "Login Theme",
								MarkdownDescription: "Login Theme",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_failure_wait_seconds": {
								Description:         "Max Wait",
								MarkdownDescription: "Max Wait",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_policy_digits": {
								Description:         "OTP Policy Digits",
								MarkdownDescription: "OTP Policy Digits",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"roles": {
								Description:         "Roles",
								MarkdownDescription: "Roles",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client": {
										Description:         "Client Roles",
										MarkdownDescription: "Client Roles",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"realm": {
										Description:         "Realm Roles",
										MarkdownDescription: "Realm Roles",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"id": {
												Description:         "Id",
												MarkdownDescription: "Id",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name",
												MarkdownDescription: "Name",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"attributes": {
												Description:         "Role Attributes",
												MarkdownDescription: "Role Attributes",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_role": {
												Description:         "Client Role",
												MarkdownDescription: "Client Role",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"composite": {
												Description:         "Composite",
												MarkdownDescription: "Composite",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"composites": {
												Description:         "Composites",
												MarkdownDescription: "Composites",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client": {
														Description:         "Map client => []role",
														MarkdownDescription: "Map client => []role",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"realm": {
														Description:         "Realm roles",
														MarkdownDescription: "Realm roles",

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

											"container_id": {
												Description:         "Container Id",
												MarkdownDescription: "Container Id",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": {
												Description:         "Description",
												MarkdownDescription: "Description",

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

							"ssl_required": {
								Description:         "Require SSL",
								MarkdownDescription: "Require SSL",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_federation_providers": {
								Description:         "Point keycloak to an external user provider to validate credentials or pull in identity information.",
								MarkdownDescription: "Point keycloak to an external user provider to validate credentials or pull in identity information.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"id": {
										Description:         "The ID of this provider",
										MarkdownDescription: "The ID of this provider",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": {
										Description:         "The priority of this provider when looking up users or adding a user.",
										MarkdownDescription: "The priority of this provider when looking up users or adding a user.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"provider_name": {
										Description:         "The name of the user provider, such as 'ldap', 'kerberos' or a custom SPI.",
										MarkdownDescription: "The name of the user provider, such as 'ldap', 'kerberos' or a custom SPI.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"changed_sync_period": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config": {
										Description:         "User federation provider config.",
										MarkdownDescription: "User federation provider config.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"display_name": {
										Description:         "The display name of this provider instance.",
										MarkdownDescription: "The display name of this provider instance.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"full_sync_period": {
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

							"brute_force_protected": {
								Description:         "Brute Force Detection",
								MarkdownDescription: "Brute Force Detection",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_supported_applications": {
								Description:         "OTP Supported Applications",
								MarkdownDescription: "OTP Supported Applications",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"access_token_lifespan_for_implicit_flow": {
								Description:         "Access Token Lifespan For Implicit Flow",
								MarkdownDescription: "Access Token Lifespan For Implicit Flow",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_events_details_enabled": {
								Description:         "Enable admin events details TODO: change to values and use kubebuilder default annotation once supported",
								MarkdownDescription: "Enable admin events details TODO: change to values and use kubebuilder default annotation once supported",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_policy_look_ahead_window": {
								Description:         "OTP Policy Look Ahead Window",
								MarkdownDescription: "OTP Policy Look Ahead Window",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"registration_email_as_username": {
								Description:         "Email as username",
								MarkdownDescription: "Email as username",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"verify_email": {
								Description:         "Verify email",
								MarkdownDescription: "Verify email",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password_policy": {
								Description:         "Realm Password Policy",
								MarkdownDescription: "Realm Password Policy",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_scope_mappings": {
								Description:         "Client Scope Mappings",
								MarkdownDescription: "Client Scope Mappings",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_scopes": {
								Description:         "Client scopes",
								MarkdownDescription: "Client scopes",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"protocol": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol_mappers": {
										Description:         "Protocol Mappers.",
										MarkdownDescription: "Protocol Mappers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"consent_text": {
												Description:         "Text to use for displaying Consent Screen.",
												MarkdownDescription: "Text to use for displaying Consent Screen.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "Protocol Mapper ID.",
												MarkdownDescription: "Protocol Mapper ID.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Protocol Mapper Name.",
												MarkdownDescription: "Protocol Mapper Name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol": {
												Description:         "Protocol to use.",
												MarkdownDescription: "Protocol to use.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol_mapper": {
												Description:         "Protocol Mapper to use",
												MarkdownDescription: "Protocol Mapper to use",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"config": {
												Description:         "Config options.",
												MarkdownDescription: "Config options.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consent_required": {
												Description:         "True if Consent Screen is required.",
												MarkdownDescription: "True if Consent Screen is required.",

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

									"attributes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_default_client_scopes": {
								Description:         "Default client scopes to add to all new clients",
								MarkdownDescription: "Default client scopes to add to all new clients",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"email_theme": {
								Description:         "Email Theme",
								MarkdownDescription: "Email Theme",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"login_with_email_allowed": {
								Description:         "Login with email",
								MarkdownDescription: "Login with email",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scope_mappings": {
								Description:         "Scope Mappings",
								MarkdownDescription: "Scope Mappings",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"client": {
										Description:         "Client",
										MarkdownDescription: "Client",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_scope": {
										Description:         "Client Scope",
										MarkdownDescription: "Client Scope",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles",
										MarkdownDescription: "Roles",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"self": {
										Description:         "Self",
										MarkdownDescription: "Self",

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

							"user_federation_mappers": {
								Description:         "User federation mappers are extension points triggered by the user federation at various points.",
								MarkdownDescription: "User federation mappers are extension points triggered by the user federation at various points.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "User federation mapper config.",
										MarkdownDescription: "User federation mapper config.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"federation_mapper_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"federation_provider_display_name": {
										Description:         "The displayName for the user federation provider this mapper applies to.",
										MarkdownDescription: "The displayName for the user federation provider this mapper applies to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_locale": {
								Description:         "Default Locale",
								MarkdownDescription: "Default Locale",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"display_name": {
								Description:         "Realm display name.",
								MarkdownDescription: "Realm display name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled_event_types": {
								Description:         "Enabled event types",
								MarkdownDescription: "Enabled event types",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_policy_algorithm": {
								Description:         "OTP Policy Algorithm",
								MarkdownDescription: "OTP Policy Algorithm",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"remember_me": {
								Description:         "Remember me",
								MarkdownDescription: "Remember me",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"admin_theme": {
								Description:         "Admin Console Theme",
								MarkdownDescription: "Admin Console Theme",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"edit_username_allowed": {
								Description:         "Edit username",
								MarkdownDescription: "Edit username",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "Realm enabled flag.",
								MarkdownDescription: "Realm enabled flag.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"smtp_server": {
								Description:         "Email",
								MarkdownDescription: "Email",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"account_theme": {
								Description:         "Account Theme",
								MarkdownDescription: "Account Theme",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"clients": {
								Description:         "A set of Keycloak Clients.",
								MarkdownDescription: "A set of Keycloak Clients.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"consent_required": {
										Description:         "True if Consent Screen is required.",
										MarkdownDescription: "True if Consent Screen is required.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"frontchannel_logout": {
										Description:         "True if this client supports Front Channel logout.",
										MarkdownDescription: "True if this client supports Front Channel logout.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_accounts_enabled": {
										Description:         "True if Service Accounts are enabled.",
										MarkdownDescription: "True if Service Accounts are enabled.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authentication_flow_binding_overrides": {
										Description:         "Authentication Flow Binding Overrides.",
										MarkdownDescription: "Authentication Flow Binding Overrides.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"base_url": {
										Description:         "Application base URL.",
										MarkdownDescription: "Application base URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_authenticator_type": {
										Description:         "What Client authentication type to use.",
										MarkdownDescription: "What Client authentication type to use.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional_client_scopes": {
										Description:         "A list of optional client scopes. Optional client scopes are applied when issuing tokens for this client, but only when they are requested by the scope parameter in the OpenID Connect authorization request.",
										MarkdownDescription: "A list of optional client scopes. Optional client scopes are applied when issuing tokens for this client, but only when they are requested by the scope parameter in the OpenID Connect authorization request.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "Protocol used for this Client.",
										MarkdownDescription: "Protocol used for this Client.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"bearer_only": {
										Description:         "True if a client supports only Bearer Tokens.",
										MarkdownDescription: "True if a client supports only Bearer Tokens.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"redirect_uris": {
										Description:         "A list of valid Redirection URLs.",
										MarkdownDescription: "A list of valid Redirection URLs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_url": {
										Description:         "Application root URL.",
										MarkdownDescription: "Application root URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"standard_flow_enabled": {
										Description:         "True if Standard flow is enabled.",
										MarkdownDescription: "True if Standard flow is enabled.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"surrogate_auth_required": {
										Description:         "Surrogate Authentication Required option.",
										MarkdownDescription: "Surrogate Authentication Required option.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_template_scope": {
										Description:         "True to use Template Scope.",
										MarkdownDescription: "True to use Template Scope.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"web_origins": {
										Description:         "A list of valid Web Origins.",
										MarkdownDescription: "A list of valid Web Origins.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"admin_url": {
										Description:         "Application Admin URL.",
										MarkdownDescription: "Application Admin URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization_services_enabled": {
										Description:         "True if fine-grained authorization support is enabled for this client.",
										MarkdownDescription: "True if fine-grained authorization support is enabled for this client.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Client enabled flag.",
										MarkdownDescription: "Client enabled flag.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"implicit_flow_enabled": {
										Description:         "True if Implicit flow is enabled.",
										MarkdownDescription: "True if Implicit flow is enabled.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_template_config": {
										Description:         "True to use a Template Config.",
										MarkdownDescription: "True to use a Template Config.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_id": {
										Description:         "Client ID.",
										MarkdownDescription: "Client ID.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"description": {
										Description:         "Client description.",
										MarkdownDescription: "Client description.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"direct_access_grants_enabled": {
										Description:         "True if Direct Grant is enabled.",
										MarkdownDescription: "True if Direct Grant is enabled.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "Client ID. If not specified, automatically generated.",
										MarkdownDescription: "Client ID. If not specified, automatically generated.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_re_registration_timeout": {
										Description:         "Node registration timeout.",
										MarkdownDescription: "Node registration timeout.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization_settings": {
										Description:         "Authorization settings for this resource server.",
										MarkdownDescription: "Authorization settings for this resource server.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_remote_resource_management": {
												Description:         "True if resources should be managed remotely by the resource server.",
												MarkdownDescription: "True if resources should be managed remotely by the resource server.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"decision_strategy": {
												Description:         "The decision strategy dictates how permissions are evaluated and how a final decision is obtained. 'Affirmative' means that at least one permission must evaluate to a positive decision in order to grant access to a resource and its scopes. 'Unanimous' means that all permissions must evaluate to a positive decision in order for the final decision to be also positive.",
												MarkdownDescription: "The decision strategy dictates how permissions are evaluated and how a final decision is obtained. 'Affirmative' means that at least one permission must evaluate to a positive decision in order to grant access to a resource and its scopes. 'Unanimous' means that all permissions must evaluate to a positive decision in order for the final decision to be also positive.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "ID.",
												MarkdownDescription: "ID.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name.",
												MarkdownDescription: "Name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "Resources.",
												MarkdownDescription: "Resources.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"_id": {
														Description:         "ID.",
														MarkdownDescription: "ID.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"icon_uri": {
														Description:         "An URI pointing to an icon.",
														MarkdownDescription: "An URI pointing to an icon.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
														MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"owner_managed_access": {
														Description:         "True if the access to this resource can be managed by the resource owner.",
														MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "The attributes associated with the resource.",
														MarkdownDescription: "The attributes associated with the resource.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"display_name": {
														Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
														MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "The scopes associated with this resource.",
														MarkdownDescription: "The scopes associated with this resource.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
														MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"uris": {
														Description:         "Set of URIs which are protected by resource.",
														MarkdownDescription: "Set of URIs which are protected by resource.",

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

											"scopes": {
												Description:         "Authorization Scopes.",
												MarkdownDescription: "Authorization Scopes.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"display_name": {
														Description:         "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",
														MarkdownDescription: "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"icon_uri": {
														Description:         "An URI pointing to an icon.",
														MarkdownDescription: "An URI pointing to an icon.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id": {
														Description:         "ID.",
														MarkdownDescription: "ID.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",
														MarkdownDescription: "A unique name for this scope. The name can be used to uniquely identify a scope, useful when querying for a specific scope.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"policies": {
														Description:         "Policies.",
														MarkdownDescription: "Policies.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"decision_strategy": {
																Description:         "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",
																MarkdownDescription: "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"id": {
																Description:         "ID.",
																MarkdownDescription: "ID.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"logic": {
																Description:         "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
																MarkdownDescription: "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"policies": {
																Description:         "Policies.",
																MarkdownDescription: "Policies.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"resources_data": {
																Description:         "Resources Data.",
																MarkdownDescription: "Resources Data.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"scopes": {
																		Description:         "The scopes associated with this resource.",
																		MarkdownDescription: "The scopes associated with this resource.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
																		Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																		MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "The attributes associated with the resource.",
																		MarkdownDescription: "The attributes associated with the resource.",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"display_name": {
																		Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																		MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																		MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"uris": {
																		Description:         "Set of URIs which are protected by resource.",
																		MarkdownDescription: "Set of URIs which are protected by resource.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"_id": {
																		Description:         "ID.",
																		MarkdownDescription: "ID.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"icon_uri": {
																		Description:         "An URI pointing to an icon.",
																		MarkdownDescription: "An URI pointing to an icon.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"owner_managed_access": {
																		Description:         "True if the access to this resource can be managed by the resource owner.",
																		MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",

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

															"scopes": {
																Description:         "Scopes.",
																MarkdownDescription: "Scopes.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"scopes_data": {
																Description:         "Scopes Data.",
																MarkdownDescription: "Scopes Data.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"config": {
																Description:         "Config.",
																MarkdownDescription: "Config.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"description": {
																Description:         "A description for this policy.",
																MarkdownDescription: "A description for this policy.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of this policy.",
																MarkdownDescription: "The name of this policy.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner": {
																Description:         "Owner.",
																MarkdownDescription: "Owner.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"resources": {
																Description:         "Resources.",
																MarkdownDescription: "Resources.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "Type.",
																MarkdownDescription: "Type.",

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

													"resources": {
														Description:         "Resources.",
														MarkdownDescription: "Resources.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"_id": {
																Description:         "ID.",
																MarkdownDescription: "ID.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"display_name": {
																Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner_managed_access": {
																Description:         "True if the access to this resource can be managed by the resource owner.",
																MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"uris": {
																Description:         "Set of URIs which are protected by resource.",
																MarkdownDescription: "Set of URIs which are protected by resource.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"attributes": {
																Description:         "The attributes associated with the resource.",
																MarkdownDescription: "The attributes associated with the resource.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"icon_uri": {
																Description:         "An URI pointing to an icon.",
																MarkdownDescription: "An URI pointing to an icon.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"scopes": {
																Description:         "The scopes associated with this resource.",
																MarkdownDescription: "The scopes associated with this resource.",

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

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_id": {
												Description:         "Client ID.",
												MarkdownDescription: "Client ID.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"policies": {
												Description:         "Policies.",
												MarkdownDescription: "Policies.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"decision_strategy": {
														Description:         "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",
														MarkdownDescription: "The decision strategy dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. 'Affirmative' means that at least one policy must evaluate to a positive decision in order for the final decision to be also positive. 'Unanimous' means that all policies must evaluate to a positive decision in order for the final decision to be also positive. 'Consensus' means that the number of positive decisions must be greater than the number of negative decisions. If the number of positive and negative is the same, the final decision will be negative.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"description": {
														Description:         "A description for this policy.",
														MarkdownDescription: "A description for this policy.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id": {
														Description:         "ID.",
														MarkdownDescription: "ID.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of this policy.",
														MarkdownDescription: "The name of this policy.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"policies": {
														Description:         "Policies.",
														MarkdownDescription: "Policies.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources_data": {
														Description:         "Resources Data.",
														MarkdownDescription: "Resources Data.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"_id": {
																Description:         "ID.",
																MarkdownDescription: "ID.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"icon_uri": {
																Description:         "An URI pointing to an icon.",
																MarkdownDescription: "An URI pointing to an icon.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"owner_managed_access": {
																Description:         "True if the access to this resource can be managed by the resource owner.",
																MarkdownDescription: "True if the access to this resource can be managed by the resource owner.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "The type of this resource. It can be used to group different resource instances with the same type.",
																MarkdownDescription: "The type of this resource. It can be used to group different resource instances with the same type.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"uris": {
																Description:         "Set of URIs which are protected by resource.",
																MarkdownDescription: "Set of URIs which are protected by resource.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"attributes": {
																Description:         "The attributes associated with the resource.",
																MarkdownDescription: "The attributes associated with the resource.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"display_name": {
																Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
																MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"scopes": {
																Description:         "The scopes associated with this resource.",
																MarkdownDescription: "The scopes associated with this resource.",

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

													"scopes": {
														Description:         "Scopes.",
														MarkdownDescription: "Scopes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"config": {
														Description:         "Config.",
														MarkdownDescription: "Config.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"logic": {
														Description:         "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
														MarkdownDescription: "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"owner": {
														Description:         "Owner.",
														MarkdownDescription: "Owner.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Resources.",
														MarkdownDescription: "Resources.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes_data": {
														Description:         "Scopes Data.",
														MarkdownDescription: "Scopes Data.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "Type.",
														MarkdownDescription: "Type.",

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

											"policy_enforcement_mode": {
												Description:         "The policy enforcement mode dictates how policies are enforced when evaluating authorization requests. 'Enforcing' means requests are denied by default even when there is no policy associated with a given resource. 'Permissive' means requests are allowed even when there is no policy associated with a given resource. 'Disabled' completely disables the evaluation of policies and allows access to any resource.",
												MarkdownDescription: "The policy enforcement mode dictates how policies are enforced when evaluating authorization requests. 'Enforcing' means requests are denied by default even when there is no policy associated with a given resource. 'Permissive' means requests are allowed even when there is no policy associated with a given resource. 'Disabled' completely disables the evaluation of policies and allows access to any resource.",

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

									"secret": {
										Description:         "Client Secret. The Operator will automatically create a Secret based on this value.",
										MarkdownDescription: "Client Secret. The Operator will automatically create a Secret based on this value.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_template_mappers": {
										Description:         "True to use Template Mappers.",
										MarkdownDescription: "True to use Template Mappers.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_client_scopes": {
										Description:         "A list of default client scopes. Default client scopes are always applied when issuing OpenID Connect tokens or SAML assertions for this client.",
										MarkdownDescription: "A list of default client scopes. Default client scopes are always applied when issuing OpenID Connect tokens or SAML assertions for this client.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_roles": {
										Description:         "Default Client roles.",
										MarkdownDescription: "Default Client roles.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"full_scope_allowed": {
										Description:         "True if Full Scope is allowed.",
										MarkdownDescription: "True if Full Scope is allowed.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_before": {
										Description:         "Not Before setting.",
										MarkdownDescription: "Not Before setting.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol_mappers": {
										Description:         "Protocol Mappers.",
										MarkdownDescription: "Protocol Mappers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"config": {
												Description:         "Config options.",
												MarkdownDescription: "Config options.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consent_required": {
												Description:         "True if Consent Screen is required.",
												MarkdownDescription: "True if Consent Screen is required.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consent_text": {
												Description:         "Text to use for displaying Consent Screen.",
												MarkdownDescription: "Text to use for displaying Consent Screen.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "Protocol Mapper ID.",
												MarkdownDescription: "Protocol Mapper ID.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Protocol Mapper Name.",
												MarkdownDescription: "Protocol Mapper Name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol": {
												Description:         "Protocol to use.",
												MarkdownDescription: "Protocol to use.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol_mapper": {
												Description:         "Protocol Mapper to use",
												MarkdownDescription: "Protocol Mapper to use",

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

									"access": {
										Description:         "Access options.",
										MarkdownDescription: "Access options.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"attributes": {
										Description:         "Client Attributes.",
										MarkdownDescription: "Client Attributes.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Client name.",
										MarkdownDescription: "Client name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"public_client": {
										Description:         "True if this is a public Client.",
										MarkdownDescription: "True if this is a public Client.",

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

							"events_listeners": {
								Description:         "A set of Event Listeners.",
								MarkdownDescription: "A set of Event Listeners.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_providers": {
								Description:         "A set of Identity Providers.",
								MarkdownDescription: "A set of Identity Providers.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"display_name": {
										Description:         "Identity Provider Display Name.",
										MarkdownDescription: "Identity Provider Display Name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Identity Provider enabled flag.",
										MarkdownDescription: "Identity Provider enabled flag.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"first_broker_login_flow_alias": {
										Description:         "Identity Provider First Broker Login Flow Alias.",
										MarkdownDescription: "Identity Provider First Broker Login Flow Alias.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"internal_id": {
										Description:         "Identity Provider Internal ID.",
										MarkdownDescription: "Identity Provider Internal ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"link_only": {
										Description:         "Identity Provider Link Only setting.",
										MarkdownDescription: "Identity Provider Link Only setting.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"store_token": {
										Description:         "Identity Provider Store to Token.",
										MarkdownDescription: "Identity Provider Store to Token.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"alias": {
										Description:         "Identity Provider Alias.",
										MarkdownDescription: "Identity Provider Alias.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config": {
										Description:         "Identity Provider config.",
										MarkdownDescription: "Identity Provider config.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"post_broker_login_flow_alias": {
										Description:         "Identity Provider Post Broker Login Flow Alias.",
										MarkdownDescription: "Identity Provider Post Broker Login Flow Alias.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"provider_id": {
										Description:         "Identity Provider ID.",
										MarkdownDescription: "Identity Provider ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"trust_email": {
										Description:         "Identity Provider Trust Email.",
										MarkdownDescription: "Identity Provider Trust Email.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"add_read_token_role_on_create": {
										Description:         "Adds Read Token role when creating this Identity Provider.",
										MarkdownDescription: "Adds Read Token role when creating this Identity Provider.",

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

							"internationalization_enabled": {
								Description:         "Internationalization Enabled",
								MarkdownDescription: "Internationalization Enabled",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"otp_policy_initial_counter": {
								Description:         "OTP Policy Initial Counter",
								MarkdownDescription: "OTP Policy Initial Counter",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_locales": {
								Description:         "Supported Locales",
								MarkdownDescription: "Supported Locales",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"realm_overrides": {
						Description:         "A list of overrides to the default Realm behavior.",
						MarkdownDescription: "A list of overrides to the default Realm behavior.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"for_flow": {
								Description:         "Flow to be overridden.",
								MarkdownDescription: "Flow to be overridden.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider": {
								Description:         "Identity Provider to be overridden.",
								MarkdownDescription: "Identity Provider to be overridden.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"unmanaged": {
						Description:         "When set to true, this KeycloakRealm will be marked as unmanaged and not be managed by this operator. It can then be used for targeting purposes.",
						MarkdownDescription: "When set to true, this KeycloakRealm will be marked as unmanaged and not be managed by this operator. It can then be used for targeting purposes.",

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
		},
	}, nil
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_keycloak_org_keycloak_realm_v1alpha1")

	var state KeycloakOrgKeycloakRealmV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakRealmV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakRealm")

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

func (r *KeycloakOrgKeycloakRealmV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_realm_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *KeycloakOrgKeycloakRealmV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_keycloak_org_keycloak_realm_v1alpha1")

	var state KeycloakOrgKeycloakRealmV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakRealmV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakRealm")

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

func (r *KeycloakOrgKeycloakRealmV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_keycloak_org_keycloak_realm_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
