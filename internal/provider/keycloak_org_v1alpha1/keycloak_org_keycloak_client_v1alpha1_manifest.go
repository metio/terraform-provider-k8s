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
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KeycloakOrgKeycloakClientV1Alpha1Manifest{}
)

func NewKeycloakOrgKeycloakClientV1Alpha1Manifest() datasource.DataSource {
	return &KeycloakOrgKeycloakClientV1Alpha1Manifest{}
}

type KeycloakOrgKeycloakClientV1Alpha1Manifest struct{}

type KeycloakOrgKeycloakClientV1Alpha1ManifestData struct {
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
		Client *struct {
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
		} `tfsdk:"client" json:"client,omitempty"`
		RealmSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"realm_selector" json:"realmSelector,omitempty"`
		Roles *[]struct {
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
		} `tfsdk:"roles" json:"roles,omitempty"`
		ScopeMappings *struct {
			ClientMappings *struct {
				Client   *string `tfsdk:"client" json:"client,omitempty"`
				Id       *string `tfsdk:"id" json:"id,omitempty"`
				Mappings *[]struct {
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
				} `tfsdk:"mappings" json:"mappings,omitempty"`
			} `tfsdk:"client_mappings" json:"clientMappings,omitempty"`
			RealmMappings *[]struct {
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
			} `tfsdk:"realm_mappings" json:"realmMappings,omitempty"`
		} `tfsdk:"scope_mappings" json:"scopeMappings,omitempty"`
		ServiceAccountClientRoles *map[string][]string `tfsdk:"service_account_client_roles" json:"serviceAccountClientRoles,omitempty"`
		ServiceAccountRealmRoles  *[]string            `tfsdk:"service_account_realm_roles" json:"serviceAccountRealmRoles,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keycloak_org_keycloak_client_v1alpha1_manifest"
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KeycloakClient is the Schema for the keycloakclients API.",
		MarkdownDescription: "KeycloakClient is the Schema for the keycloakclients API.",
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
				Description:         "KeycloakClientSpec defines the desired state of KeycloakClient.",
				MarkdownDescription: "KeycloakClientSpec defines the desired state of KeycloakClient.",
				Attributes: map[string]schema.Attribute{
					"client": schema.SingleNestedAttribute{
						Description:         "Keycloak Client REST object.",
						MarkdownDescription: "Keycloak Client REST object.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"realm_selector": schema.SingleNestedAttribute{
						Description:         "Selector for looking up KeycloakRealm Custom Resources.",
						MarkdownDescription: "Selector for looking up KeycloakRealm Custom Resources.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"roles": schema.ListNestedAttribute{
						Description:         "Client Roles",
						MarkdownDescription: "Client Roles",
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

					"scope_mappings": schema.SingleNestedAttribute{
						Description:         "Scope Mappings",
						MarkdownDescription: "Scope Mappings",
						Attributes: map[string]schema.Attribute{
							"client_mappings": schema.SingleNestedAttribute{
								Description:         "Client Mappings",
								MarkdownDescription: "Client Mappings",
								Attributes: map[string]schema.Attribute{
									"client": schema.StringAttribute{
										Description:         "Client",
										MarkdownDescription: "Client",
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

									"mappings": schema.ListNestedAttribute{
										Description:         "Mappings",
										MarkdownDescription: "Mappings",
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

							"realm_mappings": schema.ListNestedAttribute{
								Description:         "Realm Mappings",
								MarkdownDescription: "Realm Mappings",
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

					"service_account_client_roles": schema.MapAttribute{
						Description:         "Service account client roles for this client.",
						MarkdownDescription: "Service account client roles for this client.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_realm_roles": schema.ListAttribute{
						Description:         "Service account realm roles for this client.",
						MarkdownDescription: "Service account realm roles for this client.",
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
	}
}

func (r *KeycloakOrgKeycloakClientV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keycloak_org_keycloak_client_v1alpha1_manifest")

	var model KeycloakOrgKeycloakClientV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("keycloak.org/v1alpha1")
	model.Kind = pointer.String("KeycloakClient")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
