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

type KeycloakOrgKeycloakClientV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*KeycloakOrgKeycloakClientV1Alpha1Resource)(nil)
)

type KeycloakOrgKeycloakClientV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KeycloakOrgKeycloakClientV1Alpha1GoModel struct {
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
		Client *struct {
			Access *map[string]string `tfsdk:"access" yaml:"access,omitempty"`

			AuthenticationFlowBindingOverrides *map[string]string `tfsdk:"authentication_flow_binding_overrides" yaml:"authenticationFlowBindingOverrides,omitempty"`

			DefaultClientScopes *[]string `tfsdk:"default_client_scopes" yaml:"defaultClientScopes,omitempty"`

			ImplicitFlowEnabled *bool `tfsdk:"implicit_flow_enabled" yaml:"implicitFlowEnabled,omitempty"`

			UseTemplateScope *bool `tfsdk:"use_template_scope" yaml:"useTemplateScope,omitempty"`

			BearerOnly *bool `tfsdk:"bearer_only" yaml:"bearerOnly,omitempty"`

			ConsentRequired *bool `tfsdk:"consent_required" yaml:"consentRequired,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			ProtocolMappers *[]struct {
				ProtocolMapper *string `tfsdk:"protocol_mapper" yaml:"protocolMapper,omitempty"`

				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				ConsentRequired *bool `tfsdk:"consent_required" yaml:"consentRequired,omitempty"`

				ConsentText *string `tfsdk:"consent_text" yaml:"consentText,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
			} `tfsdk:"protocol_mappers" yaml:"protocolMappers,omitempty"`

			NotBefore *int64 `tfsdk:"not_before" yaml:"notBefore,omitempty"`

			OptionalClientScopes *[]string `tfsdk:"optional_client_scopes" yaml:"optionalClientScopes,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

			AuthorizationServicesEnabled *bool `tfsdk:"authorization_services_enabled" yaml:"authorizationServicesEnabled,omitempty"`

			DefaultRoles *[]string `tfsdk:"default_roles" yaml:"defaultRoles,omitempty"`

			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			DirectAccessGrantsEnabled *bool `tfsdk:"direct_access_grants_enabled" yaml:"directAccessGrantsEnabled,omitempty"`

			FullScopeAllowed *bool `tfsdk:"full_scope_allowed" yaml:"fullScopeAllowed,omitempty"`

			NodeReRegistrationTimeout *int64 `tfsdk:"node_re_registration_timeout" yaml:"nodeReRegistrationTimeout,omitempty"`

			ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

			ServiceAccountsEnabled *bool `tfsdk:"service_accounts_enabled" yaml:"serviceAccountsEnabled,omitempty"`

			SurrogateAuthRequired *bool `tfsdk:"surrogate_auth_required" yaml:"surrogateAuthRequired,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			PublicClient *bool `tfsdk:"public_client" yaml:"publicClient,omitempty"`

			StandardFlowEnabled *bool `tfsdk:"standard_flow_enabled" yaml:"standardFlowEnabled,omitempty"`

			UseTemplateMappers *bool `tfsdk:"use_template_mappers" yaml:"useTemplateMappers,omitempty"`

			WebOrigins *[]string `tfsdk:"web_origins" yaml:"webOrigins,omitempty"`

			AdminUrl *string `tfsdk:"admin_url" yaml:"adminUrl,omitempty"`

			AuthorizationSettings *struct {
				Resources *[]struct {
					Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

					DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

					Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

					OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

					Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				AllowRemoteResourceManagement *bool `tfsdk:"allow_remote_resource_management" yaml:"allowRemoteResourceManagement,omitempty"`

				ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

				DecisionStrategy *string `tfsdk:"decision_strategy" yaml:"decisionStrategy,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				PolicyEnforcementMode *string `tfsdk:"policy_enforcement_mode" yaml:"policyEnforcementMode,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Policies *[]struct {
					DecisionStrategy *string `tfsdk:"decision_strategy" yaml:"decisionStrategy,omitempty"`

					Description *string `tfsdk:"description" yaml:"description,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Logic *string `tfsdk:"logic" yaml:"logic,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

					Policies *[]string `tfsdk:"policies" yaml:"policies,omitempty"`

					Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

					Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

					ResourcesData *[]struct {
						_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

						DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

						Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`
					} `tfsdk:"resources_data" yaml:"resourcesData,omitempty"`

					ScopesData *[]string `tfsdk:"scopes_data" yaml:"scopesData,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`
				} `tfsdk:"policies" yaml:"policies,omitempty"`

				Scopes *[]struct {
					DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

					IconUri *string `tfsdk:"icon_uri" yaml:"iconUri,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Policies *[]struct {
						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

						Description *string `tfsdk:"description" yaml:"description,omitempty"`

						Logic *string `tfsdk:"logic" yaml:"logic,omitempty"`

						Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

						Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

						ResourcesData *[]struct {
							_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

							DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

							OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`

							Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

							Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`
						} `tfsdk:"resources_data" yaml:"resourcesData,omitempty"`

						DecisionStrategy *string `tfsdk:"decision_strategy" yaml:"decisionStrategy,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Policies *[]string `tfsdk:"policies" yaml:"policies,omitempty"`

						ScopesData *[]string `tfsdk:"scopes_data" yaml:"scopesData,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"policies" yaml:"policies,omitempty"`

					Resources *[]struct {
						Uris *[]string `tfsdk:"uris" yaml:"uris,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

						Icon_uri *string `tfsdk:"icon_uri" yaml:"icon_uri,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						OwnerManagedAccess *bool `tfsdk:"owner_managed_access" yaml:"ownerManagedAccess,omitempty"`

						_id *string `tfsdk:"_id" yaml:"_id,omitempty"`

						Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`
				} `tfsdk:"scopes" yaml:"scopes,omitempty"`
			} `tfsdk:"authorization_settings" yaml:"authorizationSettings,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			RedirectUris *[]string `tfsdk:"redirect_uris" yaml:"redirectUris,omitempty"`

			RootUrl *string `tfsdk:"root_url" yaml:"rootUrl,omitempty"`

			UseTemplateConfig *bool `tfsdk:"use_template_config" yaml:"useTemplateConfig,omitempty"`

			Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

			BaseUrl *string `tfsdk:"base_url" yaml:"baseUrl,omitempty"`

			ClientAuthenticatorType *string `tfsdk:"client_authenticator_type" yaml:"clientAuthenticatorType,omitempty"`

			FrontchannelLogout *bool `tfsdk:"frontchannel_logout" yaml:"frontchannelLogout,omitempty"`
		} `tfsdk:"client" yaml:"client,omitempty"`

		RealmSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"realm_selector" yaml:"realmSelector,omitempty"`

		Roles *[]struct {
			ContainerId *string `tfsdk:"container_id" yaml:"containerId,omitempty"`

			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Attributes *map[string][]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

			ClientRole *bool `tfsdk:"client_role" yaml:"clientRole,omitempty"`

			Composite *bool `tfsdk:"composite" yaml:"composite,omitempty"`

			Composites *struct {
				Client *map[string][]string `tfsdk:"client" yaml:"client,omitempty"`

				Realm *[]string `tfsdk:"realm" yaml:"realm,omitempty"`
			} `tfsdk:"composites" yaml:"composites,omitempty"`
		} `tfsdk:"roles" yaml:"roles,omitempty"`

		ScopeMappings *struct {
			ClientMappings *map[string]string `tfsdk:"client_mappings" yaml:"clientMappings,omitempty"`

			RealmMappings *[]struct {
				ContainerId *string `tfsdk:"container_id" yaml:"containerId,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Attributes *map[string][]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

				ClientRole *bool `tfsdk:"client_role" yaml:"clientRole,omitempty"`

				Composite *bool `tfsdk:"composite" yaml:"composite,omitempty"`

				Composites *struct {
					Client *map[string][]string `tfsdk:"client" yaml:"client,omitempty"`

					Realm *[]string `tfsdk:"realm" yaml:"realm,omitempty"`
				} `tfsdk:"composites" yaml:"composites,omitempty"`
			} `tfsdk:"realm_mappings" yaml:"realmMappings,omitempty"`
		} `tfsdk:"scope_mappings" yaml:"scopeMappings,omitempty"`

		ServiceAccountClientRoles *map[string][]string `tfsdk:"service_account_client_roles" yaml:"serviceAccountClientRoles,omitempty"`

		ServiceAccountRealmRoles *[]string `tfsdk:"service_account_realm_roles" yaml:"serviceAccountRealmRoles,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKeycloakOrgKeycloakClientV1Alpha1Resource() resource.Resource {
	return &KeycloakOrgKeycloakClientV1Alpha1Resource{}
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keycloak_org_keycloak_client_v1alpha1"
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "KeycloakClient is the Schema for the keycloakclients API.",
		MarkdownDescription: "KeycloakClient is the Schema for the keycloakclients API.",
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
				Description:         "KeycloakClientSpec defines the desired state of KeycloakClient.",
				MarkdownDescription: "KeycloakClientSpec defines the desired state of KeycloakClient.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"client": {
						Description:         "Keycloak Client REST object.",
						MarkdownDescription: "Keycloak Client REST object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"access": {
								Description:         "Access options.",
								MarkdownDescription: "Access options.",

								Type: types.MapType{ElemType: types.StringType},

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

							"default_client_scopes": {
								Description:         "A list of default client scopes. Default client scopes are always applied when issuing OpenID Connect tokens or SAML assertions for this client.",
								MarkdownDescription: "A list of default client scopes. Default client scopes are always applied when issuing OpenID Connect tokens or SAML assertions for this client.",

								Type: types.ListType{ElemType: types.StringType},

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

							"use_template_scope": {
								Description:         "True to use Template Scope.",
								MarkdownDescription: "True to use Template Scope.",

								Type: types.BoolType,

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

							"consent_required": {
								Description:         "True if Consent Screen is required.",
								MarkdownDescription: "True if Consent Screen is required.",

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

							"protocol_mappers": {
								Description:         "Protocol Mappers.",
								MarkdownDescription: "Protocol Mappers.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
								}),

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

							"authorization_services_enabled": {
								Description:         "True if fine-grained authorization support is enabled for this client.",
								MarkdownDescription: "True if fine-grained authorization support is enabled for this client.",

								Type: types.BoolType,

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

							"full_scope_allowed": {
								Description:         "True if Full Scope is allowed.",
								MarkdownDescription: "True if Full Scope is allowed.",

								Type: types.BoolType,

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

							"client_id": {
								Description:         "Client ID.",
								MarkdownDescription: "Client ID.",

								Type: types.StringType,

								Required: true,
								Optional: false,
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

							"service_accounts_enabled": {
								Description:         "True if Service Accounts are enabled.",
								MarkdownDescription: "True if Service Accounts are enabled.",

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

							"standard_flow_enabled": {
								Description:         "True if Standard flow is enabled.",
								MarkdownDescription: "True if Standard flow is enabled.",

								Type: types.BoolType,

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

							"authorization_settings": {
								Description:         "Authorization settings for this resource server.",
								MarkdownDescription: "Authorization settings for this resource server.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"resources": {
										Description:         "Resources.",
										MarkdownDescription: "Resources.",

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

											"attributes": {
												Description:         "The attributes associated with the resource.",
												MarkdownDescription: "The attributes associated with the resource.",

												Type: types.MapType{ElemType: types.StringType},

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

									"allow_remote_resource_management": {
										Description:         "True if resources should be managed remotely by the resource server.",
										MarkdownDescription: "True if resources should be managed remotely by the resource server.",

										Type: types.BoolType,

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

									"decision_strategy": {
										Description:         "The decision strategy dictates how permissions are evaluated and how a final decision is obtained. 'Affirmative' means that at least one permission must evaluate to a positive decision in order to grant access to a resource and its scopes. 'Unanimous' means that all permissions must evaluate to a positive decision in order for the final decision to be also positive.",
										MarkdownDescription: "The decision strategy dictates how permissions are evaluated and how a final decision is obtained. 'Affirmative' means that at least one permission must evaluate to a positive decision in order to grant access to a resource and its scopes. 'Unanimous' means that all permissions must evaluate to a positive decision in order for the final decision to be also positive.",

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

									"policy_enforcement_mode": {
										Description:         "The policy enforcement mode dictates how policies are enforced when evaluating authorization requests. 'Enforcing' means requests are denied by default even when there is no policy associated with a given resource. 'Permissive' means requests are allowed even when there is no policy associated with a given resource. 'Disabled' completely disables the evaluation of policies and allows access to any resource.",
										MarkdownDescription: "The policy enforcement mode dictates how policies are enforced when evaluating authorization requests. 'Enforcing' means requests are denied by default even when there is no policy associated with a given resource. 'Permissive' means requests are allowed even when there is no policy associated with a given resource. 'Disabled' completely disables the evaluation of policies and allows access to any resource.",

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

											"logic": {
												Description:         "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",
												MarkdownDescription: "The logic dictates how the policy decision should be made. If 'Positive', the resulting effect (permit or deny) obtained during the evaluation of this policy will be used to perform a decision. If 'Negative', the resulting effect will be negated, in other words, a permit becomes a deny and vice-versa.",

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

											"policies": {
												Description:         "Policies.",
												MarkdownDescription: "Policies.",

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

											"scopes": {
												Description:         "Scopes.",
												MarkdownDescription: "Scopes.",

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

													"scopes": {
														Description:         "The scopes associated with this resource.",
														MarkdownDescription: "The scopes associated with this resource.",

														Type: types.ListType{ElemType: types.StringType},

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

											"resources": {
												Description:         "Resources.",
												MarkdownDescription: "Resources.",

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

													"description": {
														Description:         "A description for this policy.",
														MarkdownDescription: "A description for this policy.",

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

											"resources": {
												Description:         "Resources.",
												MarkdownDescription: "Resources.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"uris": {
														Description:         "Set of URIs which are protected by resource.",
														MarkdownDescription: "Set of URIs which are protected by resource.",

														Type: types.ListType{ElemType: types.StringType},

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

													"display_name": {
														Description:         "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",
														MarkdownDescription: "A unique name for this resource. The name can be used to uniquely identify a resource, useful when querying for a specific resource.",

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

													"_id": {
														Description:         "ID.",
														MarkdownDescription: "ID.",

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
								}),

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

							"use_template_config": {
								Description:         "True to use a Template Config.",
								MarkdownDescription: "True to use a Template Config.",

								Type: types.BoolType,

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

							"frontchannel_logout": {
								Description:         "True if this client supports Front Channel logout.",
								MarkdownDescription: "True if this client supports Front Channel logout.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"realm_selector": {
						Description:         "Selector for looking up KeycloakRealm Custom Resources.",
						MarkdownDescription: "Selector for looking up KeycloakRealm Custom Resources.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"roles": {
						Description:         "Client Roles",
						MarkdownDescription: "Client Roles",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"scope_mappings": {
						Description:         "Scope Mappings",
						MarkdownDescription: "Scope Mappings",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client_mappings": {
								Description:         "Client Mappings",
								MarkdownDescription: "Client Mappings",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"realm_mappings": {
								Description:         "Realm Mappings",
								MarkdownDescription: "Realm Mappings",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

					"service_account_client_roles": {
						Description:         "Service account client roles for this client.",
						MarkdownDescription: "Service account client roles for this client.",

						Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_realm_roles": {
						Description:         "Service account realm roles for this client.",
						MarkdownDescription: "Service account realm roles for this client.",

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
		},
	}, nil
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_keycloak_org_keycloak_client_v1alpha1")

	var state KeycloakOrgKeycloakClientV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakClientV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakClient")

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

func (r *KeycloakOrgKeycloakClientV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_client_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_keycloak_org_keycloak_client_v1alpha1")

	var state KeycloakOrgKeycloakClientV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KeycloakOrgKeycloakClientV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("keycloak.org/v1alpha1")
	goModel.Kind = utilities.Ptr("KeycloakClient")

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

func (r *KeycloakOrgKeycloakClientV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_keycloak_org_keycloak_client_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
