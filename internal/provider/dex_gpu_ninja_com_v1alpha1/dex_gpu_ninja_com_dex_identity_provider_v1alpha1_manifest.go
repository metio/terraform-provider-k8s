/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dex_gpu_ninja_com_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &DexGpuNinjaComDexIdentityProviderV1Alpha1Manifest{}
)

func NewDexGpuNinjaComDexIdentityProviderV1Alpha1Manifest() datasource.DataSource {
	return &DexGpuNinjaComDexIdentityProviderV1Alpha1Manifest{}
}

type DexGpuNinjaComDexIdentityProviderV1Alpha1Manifest struct{}

type DexGpuNinjaComDexIdentityProviderV1Alpha1ManifestData struct {
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
		ClientCertificateSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"client_certificate_secret_ref" json:"clientCertificateSecretRef,omitempty"`
		Connectors *[]struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Ldap *struct {
				BindPasswordSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"bind_password_secret_ref" json:"bindPasswordSecretRef,omitempty"`
				BindUsername *string `tfsdk:"bind_username" json:"bindUsername,omitempty"`
				CaSecretRef  *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca_secret_ref" json:"caSecretRef,omitempty"`
				ClientCertificateSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_certificate_secret_ref" json:"clientCertificateSecretRef,omitempty"`
				GroupSearch *struct {
					BaseDN       *string `tfsdk:"base_dn" json:"baseDN,omitempty"`
					Filter       *string `tfsdk:"filter" json:"filter,omitempty"`
					NameAttr     *string `tfsdk:"name_attr" json:"nameAttr,omitempty"`
					Scope        *string `tfsdk:"scope" json:"scope,omitempty"`
					UserMatchers *[]struct {
						GroupAttr *string `tfsdk:"group_attr" json:"groupAttr,omitempty"`
						UserAttr  *string `tfsdk:"user_attr" json:"userAttr,omitempty"`
					} `tfsdk:"user_matchers" json:"userMatchers,omitempty"`
				} `tfsdk:"group_search" json:"groupSearch,omitempty"`
				Host               *string `tfsdk:"host" json:"host,omitempty"`
				InsecureNoSSL      *bool   `tfsdk:"insecure_no_ssl" json:"insecureNoSSL,omitempty"`
				InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				StartTLS           *bool   `tfsdk:"start_tls" json:"startTLS,omitempty"`
				UserSearch         *struct {
					BaseDN                *string `tfsdk:"base_dn" json:"baseDN,omitempty"`
					EmailAttr             *string `tfsdk:"email_attr" json:"emailAttr,omitempty"`
					EmailSuffix           *string `tfsdk:"email_suffix" json:"emailSuffix,omitempty"`
					Filter                *string `tfsdk:"filter" json:"filter,omitempty"`
					IdAttr                *string `tfsdk:"id_attr" json:"idAttr,omitempty"`
					NameAttr              *string `tfsdk:"name_attr" json:"nameAttr,omitempty"`
					PreferredUsernameAttr *string `tfsdk:"preferred_username_attr" json:"preferredUsernameAttr,omitempty"`
					Scope                 *string `tfsdk:"scope" json:"scope,omitempty"`
					Username              *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"user_search" json:"userSearch,omitempty"`
				UsernamePrompt *string `tfsdk:"username_prompt" json:"usernamePrompt,omitempty"`
			} `tfsdk:"ldap" json:"ldap,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Oidc *struct {
				AcrValues            *[]string `tfsdk:"acr_values" json:"acrValues,omitempty"`
				BasicAuthUnsupported *bool     `tfsdk:"basic_auth_unsupported" json:"basicAuthUnsupported,omitempty"`
				CaSecretRef          *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca_secret_ref" json:"caSecretRef,omitempty"`
				ClaimMapping *struct {
					Email              *string `tfsdk:"email" json:"email,omitempty"`
					Groups             *string `tfsdk:"groups" json:"groups,omitempty"`
					Preferred_username *string `tfsdk:"preferred_username" json:"preferred_username,omitempty"`
				} `tfsdk:"claim_mapping" json:"claimMapping,omitempty"`
				ClientSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
				GetUserInfo               *bool     `tfsdk:"get_user_info" json:"getUserInfo,omitempty"`
				InsecureEnableGroups      *bool     `tfsdk:"insecure_enable_groups" json:"insecureEnableGroups,omitempty"`
				InsecureSkipEmailVerified *bool     `tfsdk:"insecure_skip_email_verified" json:"insecureSkipEmailVerified,omitempty"`
				InsecureSkipVerify        *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				Issuer                    *string   `tfsdk:"issuer" json:"issuer,omitempty"`
				OverrideClaimMapping      *bool     `tfsdk:"override_claim_mapping" json:"overrideClaimMapping,omitempty"`
				PromptType                *string   `tfsdk:"prompt_type" json:"promptType,omitempty"`
				RedirectURI               *string   `tfsdk:"redirect_uri" json:"redirectURI,omitempty"`
				Scopes                    *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
				UserIDKey                 *string   `tfsdk:"user_id_key" json:"userIDKey,omitempty"`
				UserNameKey               *string   `tfsdk:"user_name_key" json:"userNameKey,omitempty"`
			} `tfsdk:"oidc" json:"oidc,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"connectors" json:"connectors,omitempty"`
		Expiry *struct {
			AuthRequests   *string `tfsdk:"auth_requests" json:"authRequests,omitempty"`
			DeviceRequests *string `tfsdk:"device_requests" json:"deviceRequests,omitempty"`
			IdTokens       *string `tfsdk:"id_tokens" json:"idTokens,omitempty"`
			RefreshTokens  *struct {
				AbsoluteLifetime  *string `tfsdk:"absolute_lifetime" json:"absoluteLifetime,omitempty"`
				DisableRotation   *bool   `tfsdk:"disable_rotation" json:"disableRotation,omitempty"`
				ReuseInterval     *string `tfsdk:"reuse_interval" json:"reuseInterval,omitempty"`
				ValidIfNotUsedFor *string `tfsdk:"valid_if_not_used_for" json:"validIfNotUsedFor,omitempty"`
			} `tfsdk:"refresh_tokens" json:"refreshTokens,omitempty"`
			SigningKeys *string `tfsdk:"signing_keys" json:"signingKeys,omitempty"`
		} `tfsdk:"expiry" json:"expiry,omitempty"`
		Frontend *struct {
			Dir     *string `tfsdk:"dir" json:"dir,omitempty"`
			Issuer  *string `tfsdk:"issuer" json:"issuer,omitempty"`
			LogoURL *string `tfsdk:"logo_url" json:"logoURL,omitempty"`
			Theme   *string `tfsdk:"theme" json:"theme,omitempty"`
		} `tfsdk:"frontend" json:"frontend,omitempty"`
		Grpc *struct {
			Annotations          *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			CertificateSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"certificate_secret_ref" json:"certificateSecretRef,omitempty"`
			ClientCASecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"client_ca_secret_ref" json:"clientCASecretRef,omitempty"`
			Reflection *bool `tfsdk:"reflection" json:"reflection,omitempty"`
		} `tfsdk:"grpc" json:"grpc,omitempty"`
		Image   *string `tfsdk:"image" json:"image,omitempty"`
		Ingress *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Enabled     *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Hosts       *[]struct {
				Host  *string `tfsdk:"host" json:"host,omitempty"`
				Paths *[]struct {
					Path     *string `tfsdk:"path" json:"path,omitempty"`
					PathType *string `tfsdk:"path_type" json:"pathType,omitempty"`
				} `tfsdk:"paths" json:"paths,omitempty"`
			} `tfsdk:"hosts" json:"hosts,omitempty"`
			IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
			Tls              *[]struct {
				Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Issuer *string `tfsdk:"issuer" json:"issuer,omitempty"`
		Logger *struct {
			Format *string `tfsdk:"format" json:"format,omitempty"`
			Level  *string `tfsdk:"level" json:"level,omitempty"`
		} `tfsdk:"logger" json:"logger,omitempty"`
		Metrics *struct {
			Enabled  *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Interval *string `tfsdk:"interval" json:"interval,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Oauth2 *struct {
			AlwaysShowLoginScreen *bool     `tfsdk:"always_show_login_screen" json:"alwaysShowLoginScreen,omitempty"`
			GrantTypes            *[]string `tfsdk:"grant_types" json:"grantTypes,omitempty"`
			PasswordConnector     *string   `tfsdk:"password_connector" json:"passwordConnector,omitempty"`
			ResponseTypes         *[]string `tfsdk:"response_types" json:"responseTypes,omitempty"`
			SkipApprovalScreen    *bool     `tfsdk:"skip_approval_screen" json:"skipApprovalScreen,omitempty"`
		} `tfsdk:"oauth2" json:"oauth2,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Storage *struct {
			Postgres *struct {
				ConnMaxLifetime      *string `tfsdk:"conn_max_lifetime" json:"connMaxLifetime,omitempty"`
				ConnectionTimeout    *string `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
				CredentialsSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
				Database     *string `tfsdk:"database" json:"database,omitempty"`
				Host         *string `tfsdk:"host" json:"host,omitempty"`
				MaxIdleConns *int64  `tfsdk:"max_idle_conns" json:"maxIdleConns,omitempty"`
				MaxOpenConns *int64  `tfsdk:"max_open_conns" json:"maxOpenConns,omitempty"`
				Port         *int64  `tfsdk:"port" json:"port,omitempty"`
				Ssl          *struct {
					CaSecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_secret_ref" json:"caSecretRef,omitempty"`
					ClientCertificateSecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"client_certificate_secret_ref" json:"clientCertificateSecretRef,omitempty"`
					Mode       *string `tfsdk:"mode" json:"mode,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"ssl" json:"ssl,omitempty"`
			} `tfsdk:"postgres" json:"postgres,omitempty"`
			Sqlite3 *struct {
				File *string `tfsdk:"file" json:"file,omitempty"`
			} `tfsdk:"sqlite3" json:"sqlite3,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		VolumeClaimTemplates *[]struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				DataSource  *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"data_source" json:"dataSource,omitempty"`
				DataSourceRef *struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
			Status *struct {
				AccessModes               *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
				AllocatedResourceStatuses *map[string]string `tfsdk:"allocated_resource_statuses" json:"allocatedResourceStatuses,omitempty"`
				AllocatedResources        *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
				Capacity                  *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
				Conditions                *[]struct {
					LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
					LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
					Message            *string `tfsdk:"message" json:"message,omitempty"`
					Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
					Status             *string `tfsdk:"status" json:"status,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				Phase *string `tfsdk:"phase" json:"phase,omitempty"`
			} `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		VolumeMounts *[]struct {
			MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		Web *struct {
			AllowedOrigins       *[]string          `tfsdk:"allowed_origins" json:"allowedOrigins,omitempty"`
			Annotations          *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			CertificateSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"certificate_secret_ref" json:"certificateSecretRef,omitempty"`
		} `tfsdk:"web" json:"web,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DexGpuNinjaComDexIdentityProviderV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest"
}

func (r *DexGpuNinjaComDexIdentityProviderV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DexIdentityProvider is a Dex identity provider instance.",
		MarkdownDescription: "DexIdentityProvider is a Dex identity provider instance.",
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
				Description:         "DexIdentityProviderSpec defines the desired state of the Dex identity provider.",
				MarkdownDescription: "DexIdentityProviderSpec defines the desired state of the Dex identity provider.",
				Attributes: map[string]schema.Attribute{
					"client_certificate_secret_ref": schema.SingleNestedAttribute{
						Description:         "ClientCertificateSecretRef is an optional reference to a secret containing a client certificate that the operator can use for connecting to the Dex API gRPC server.",
						MarkdownDescription: "ClientCertificateSecretRef is an optional reference to a secret containing a client certificate that the operator can use for connecting to the Dex API gRPC server.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the secret.",
								MarkdownDescription: "Name is the name of the secret.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"connectors": schema.ListNestedAttribute{
						Description:         "Connectors holds configuration for connectors.",
						MarkdownDescription: "Connectors holds configuration for connectors.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "ID is the connector ID.",
									MarkdownDescription: "ID is the connector ID.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"ldap": schema.SingleNestedAttribute{
									Description:         "LDAP holds configuration for the LDAP connector.",
									MarkdownDescription: "LDAP holds configuration for the LDAP connector.",
									Attributes: map[string]schema.Attribute{
										"bind_password_secret_ref": schema.SingleNestedAttribute{
											Description:         "BindPasswordSecretRef is a reference to a secret containing the bind password. The connector uses these credentials to search for users and groups.",
											MarkdownDescription: "BindPasswordSecretRef is a reference to a secret containing the bind password. The connector uses these credentials to search for users and groups.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the secret.",
													MarkdownDescription: "Name is the name of the secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"bind_username": schema.StringAttribute{
											Description:         "BindUsername is the DN of the user to bind with. The connector uses these credentials to search for users and groups.",
											MarkdownDescription: "BindUsername is the DN of the user to bind with. The connector uses these credentials to search for users and groups.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"ca_secret_ref": schema.SingleNestedAttribute{
											Description:         "CASecretRef is an optional reference to a secret containing the CA certificate.",
											MarkdownDescription: "CASecretRef is an optional reference to a secret containing the CA certificate.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the secret.",
													MarkdownDescription: "Name is the name of the secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"client_certificate_secret_ref": schema.SingleNestedAttribute{
											Description:         "ClientCertificateSecretRef is an optional reference to a secret containing the client certificate and key.",
											MarkdownDescription: "ClientCertificateSecretRef is an optional reference to a secret containing the client certificate and key.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the secret.",
													MarkdownDescription: "Name is the name of the secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"group_search": schema.SingleNestedAttribute{
											Description:         "GroupSearch contains configuration for searching LDAP groups.",
											MarkdownDescription: "GroupSearch contains configuration for searching LDAP groups.",
											Attributes: map[string]schema.Attribute{
												"base_dn": schema.StringAttribute{
													Description:         "BaseDN to start the search from. For example 'cn=groups,dc=example,dc=com'",
													MarkdownDescription: "BaseDN to start the search from. For example 'cn=groups,dc=example,dc=com'",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"filter": schema.StringAttribute{
													Description:         "Filter is an optional filter to apply when searching the directory. For example '(objectClass=posixGroup)'",
													MarkdownDescription: "Filter is an optional filter to apply when searching the directory. For example '(objectClass=posixGroup)'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name_attr": schema.StringAttribute{
													Description:         "NameAttr is the attribute of the group that represents its name.",
													MarkdownDescription: "NameAttr is the attribute of the group that represents its name.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"scope": schema.StringAttribute{
													Description:         "Scope is the optional scope of the search (default 'sub'). Can either be: * 'sub' - search the whole sub tree * 'one' - only search one level",
													MarkdownDescription: "Scope is the optional scope of the search (default 'sub'). Can either be: * 'sub' - search the whole sub tree * 'one' - only search one level",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("sub", "one"),
													},
												},

												"user_matchers": schema.ListNestedAttribute{
													Description:         "UserMatchers is an array of the field pairs used to match a user to a group. See the 'DexIdentityProviderConnectorLDAPGroupSearchUserMatcher' struct for the exact field names  Each pair adds an additional requirement to the filter that an attribute in the group match the user's attribute value. For example that the 'members' attribute of a group matches the 'uid' of the user. The exact filter being added is:  (userMatchers[n].<groupAttr>=userMatchers[n].<userAttr value>)",
													MarkdownDescription: "UserMatchers is an array of the field pairs used to match a user to a group. See the 'DexIdentityProviderConnectorLDAPGroupSearchUserMatcher' struct for the exact field names  Each pair adds an additional requirement to the filter that an attribute in the group match the user's attribute value. For example that the 'members' attribute of a group matches the 'uid' of the user. The exact filter being added is:  (userMatchers[n].<groupAttr>=userMatchers[n].<userAttr value>)",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"group_attr": schema.StringAttribute{
																Description:         "GroupAttr is the attribute to match against the group ID.",
																MarkdownDescription: "GroupAttr is the attribute to match against the group ID.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"user_attr": schema.StringAttribute{
																Description:         "UserAttr is the attribute to match against the user ID.",
																MarkdownDescription: "UserAttr is the attribute to match against the user ID.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"host": schema.StringAttribute{
											Description:         "Host is the host and optional port of the LDAP server. If port isn't supplied, it will be guessed based on the TLS configuration.",
											MarkdownDescription: "Host is the host and optional port of the LDAP server. If port isn't supplied, it will be guessed based on the TLS configuration.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"insecure_no_ssl": schema.BoolAttribute{
											Description:         "InsecureNoSSL is required to connect to a server without TLS.",
											MarkdownDescription: "InsecureNoSSL is required to connect to a server without TLS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure_skip_verify": schema.BoolAttribute{
											Description:         "InsecureSkipVerify allows connecting to a server without verifying the TLS certificate.",
											MarkdownDescription: "InsecureSkipVerify allows connecting to a server without verifying the TLS certificate.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"start_tls": schema.BoolAttribute{
											Description:         "StartTLS allows connecting to a server that supports the StartTLS command. If unsupplied secure connections will use the LDAPS protocol.",
											MarkdownDescription: "StartTLS allows connecting to a server that supports the StartTLS command. If unsupplied secure connections will use the LDAPS protocol.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_search": schema.SingleNestedAttribute{
											Description:         "UserSearch contains configuration for searching LDAP users.",
											MarkdownDescription: "UserSearch contains configuration for searching LDAP users.",
											Attributes: map[string]schema.Attribute{
												"base_dn": schema.StringAttribute{
													Description:         "BaseDN to start the search from. For example 'cn=users,dc=example,dc=com'",
													MarkdownDescription: "BaseDN to start the search from. For example 'cn=users,dc=example,dc=com'",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"email_attr": schema.StringAttribute{
													Description:         "EmailAttr is the attribute to use as the user email (default 'mail').",
													MarkdownDescription: "EmailAttr is the attribute to use as the user email (default 'mail').",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"email_suffix": schema.StringAttribute{
													Description:         "EmailSuffix if set, will be appended to the idAttr to construct the email claim. This should not include the @ character.",
													MarkdownDescription: "EmailSuffix if set, will be appended to the idAttr to construct the email claim. This should not include the @ character.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"filter": schema.StringAttribute{
													Description:         "Filter is an optional filter to apply when searching the directory. For example '(objectClass=person)'",
													MarkdownDescription: "Filter is an optional filter to apply when searching the directory. For example '(objectClass=person)'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id_attr": schema.StringAttribute{
													Description:         "IDAttr is the attribute to use as the user ID (default 'uid').",
													MarkdownDescription: "IDAttr is the attribute to use as the user ID (default 'uid').",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name_attr": schema.StringAttribute{
													Description:         "NameAttr is the attribute to use as the display name for the user.",
													MarkdownDescription: "NameAttr is the attribute to use as the display name for the user.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"preferred_username_attr": schema.StringAttribute{
													Description:         "PreferredUsernameAttr is the attribute to use as the preferred username for the user.",
													MarkdownDescription: "PreferredUsernameAttr is the attribute to use as the preferred username for the user.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scope": schema.StringAttribute{
													Description:         "Scope is the optional scope of the search (default 'sub'). Can either be: * 'sub' - search the whole sub tree * 'one' - only search one level",
													MarkdownDescription: "Scope is the optional scope of the search (default 'sub'). Can either be: * 'sub' - search the whole sub tree * 'one' - only search one level",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("sub", "one"),
													},
												},

												"username": schema.StringAttribute{
													Description:         "Username is the attribute to match against the inputted username. This will be translated and combined with the other filter as '(<attr>=<username>)'.",
													MarkdownDescription: "Username is the attribute to match against the inputted username. This will be translated and combined with the other filter as '(<attr>=<username>)'.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"username_prompt": schema.StringAttribute{
											Description:         "UsernamePrompt allows users to override the username attribute (displayed in the username/password prompt). If unset, the handler will use 'Username'.",
											MarkdownDescription: "UsernamePrompt allows users to override the username attribute (displayed in the username/password prompt). If unset, the handler will use 'Username'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the connector name.",
									MarkdownDescription: "Name is the connector name.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"oidc": schema.SingleNestedAttribute{
									Description:         "OIDC holds configuration for the OIDC connector.",
									MarkdownDescription: "OIDC holds configuration for the OIDC connector.",
									Attributes: map[string]schema.Attribute{
										"acr_values": schema.ListAttribute{
											Description:         "AcrValues (Authentication Context Class Reference Values) that specifies the Authentication Context Class Values within the Authentication Request that the Authorization Server is being requested to use for processing requests from this Client, with the values appearing in order of preference.",
											MarkdownDescription: "AcrValues (Authentication Context Class Reference Values) that specifies the Authentication Context Class Values within the Authentication Request that the Authorization Server is being requested to use for processing requests from this Client, with the values appearing in order of preference.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"basic_auth_unsupported": schema.BoolAttribute{
											Description:         "BasicAuthUnsupported causes client_secret to be passed as POST parameters instead of basic auth. This is specifically 'NOT RECOMMENDED' by the OAuth2 RFC, but some providers require it.  https://tools.ietf.org/html/rfc6749#section-2.3.1",
											MarkdownDescription: "BasicAuthUnsupported causes client_secret to be passed as POST parameters instead of basic auth. This is specifically 'NOT RECOMMENDED' by the OAuth2 RFC, but some providers require it.  https://tools.ietf.org/html/rfc6749#section-2.3.1",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ca_secret_ref": schema.SingleNestedAttribute{
											Description:         "CASecretRef is an optional reference to a secret containing the CA certificate. Only required if your provider uses a self-signed certificate.",
											MarkdownDescription: "CASecretRef is an optional reference to a secret containing the CA certificate. Only required if your provider uses a self-signed certificate.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the secret.",
													MarkdownDescription: "Name is the name of the secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"claim_mapping": schema.SingleNestedAttribute{
											Description:         "ClaimMapping is used to map non-standard claims to standard claims. Some providers return non-standard claims (eg. mail). https://openid.net/specs/openid-connect-core-1_0.html#Claims",
											MarkdownDescription: "ClaimMapping is used to map non-standard claims to standard claims. Some providers return non-standard claims (eg. mail). https://openid.net/specs/openid-connect-core-1_0.html#Claims",
											Attributes: map[string]schema.Attribute{
												"email": schema.StringAttribute{
													Description:         "EmailKey is the key which contains the email claims, defaults to 'email'.",
													MarkdownDescription: "EmailKey is the key which contains the email claims, defaults to 'email'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"groups": schema.StringAttribute{
													Description:         "GroupsKey is the key which contains the groups claims, defaults to 'groups'.",
													MarkdownDescription: "GroupsKey is the key which contains the groups claims, defaults to 'groups'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"preferred_username": schema.StringAttribute{
													Description:         "PreferredUsernameKey is the key which contains the preferred username claims, defaults to 'preferred_username'.",
													MarkdownDescription: "PreferredUsernameKey is the key which contains the preferred username claims, defaults to 'preferred_username'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"client_secret_ref": schema.SingleNestedAttribute{
											Description:         "ClientSecretRef is a reference to a secret containing the OAuth client id and secret.",
											MarkdownDescription: "ClientSecretRef is a reference to a secret containing the OAuth client id and secret.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the name of the secret.",
													MarkdownDescription: "Name is the name of the secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"get_user_info": schema.BoolAttribute{
											Description:         "GetUserInfo uses the userinfo endpoint to get additional claims for the token. This is especially useful where upstreams return 'thin' id tokens",
											MarkdownDescription: "GetUserInfo uses the userinfo endpoint to get additional claims for the token. This is especially useful where upstreams return 'thin' id tokens",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure_enable_groups": schema.BoolAttribute{
											Description:         "InsecureEnableGroups enables groups claims.",
											MarkdownDescription: "InsecureEnableGroups enables groups claims.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure_skip_email_verified": schema.BoolAttribute{
											Description:         "InsecureSkipEmailVerified if set will override the value of email_verified to true in the returned claims.",
											MarkdownDescription: "InsecureSkipEmailVerified if set will override the value of email_verified to true in the returned claims.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure_skip_verify": schema.BoolAttribute{
											Description:         "InsecureSkipVerify disables TLS certificate verification.",
											MarkdownDescription: "InsecureSkipVerify disables TLS certificate verification.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"issuer": schema.StringAttribute{
											Description:         "Issuer is the URL of the OIDC issuer.",
											MarkdownDescription: "Issuer is the URL of the OIDC issuer.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"override_claim_mapping": schema.BoolAttribute{
											Description:         "OverrideClaimMapping will be used to override the options defined in claimMappings. i.e. if there are 'email' and 'preferred_email' claims available, by default Dex will always use the 'email' claim independent of the ClaimMapping.EmailKey. This setting allows you to override the default behavior of Dex and enforce the mappings defined in 'claimMapping'. Defaults to false.",
											MarkdownDescription: "OverrideClaimMapping will be used to override the options defined in claimMappings. i.e. if there are 'email' and 'preferred_email' claims available, by default Dex will always use the 'email' claim independent of the ClaimMapping.EmailKey. This setting allows you to override the default behavior of Dex and enforce the mappings defined in 'claimMapping'. Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"prompt_type": schema.StringAttribute{
											Description:         "PromptType will be used fot the prompt parameter (when offline_access, by default prompt=consent).",
											MarkdownDescription: "PromptType will be used fot the prompt parameter (when offline_access, by default prompt=consent).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redirect_uri": schema.StringAttribute{
											Description:         "RedirectURI is the OAuth redirect URI.",
											MarkdownDescription: "RedirectURI is the OAuth redirect URI.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"scopes": schema.ListAttribute{
											Description:         "Scopes is an optional list of scopes to request. If omitted, defaults to 'profile' and 'email'.",
											MarkdownDescription: "Scopes is an optional list of scopes to request. If omitted, defaults to 'profile' and 'email'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_id_key": schema.StringAttribute{
											Description:         "UserIDKey is the claim key to use for the user ID (default sub).",
											MarkdownDescription: "UserIDKey is the claim key to use for the user ID (default sub).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_name_key": schema.StringAttribute{
											Description:         "UserNameKey is the claim key to use for the username (default name).",
											MarkdownDescription: "UserNameKey is the claim key to use for the username (default name).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "Type is the connector type to use.",
									MarkdownDescription: "Type is the connector type to use.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("ldap", "oidc"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"expiry": schema.SingleNestedAttribute{
						Description:         "Expiry holds configuration for tokens, signing keys, etc.",
						MarkdownDescription: "Expiry holds configuration for tokens, signing keys, etc.",
						Attributes: map[string]schema.Attribute{
							"auth_requests": schema.StringAttribute{
								Description:         "AuthRequests defines the duration of time for which the AuthRequests will be valid.",
								MarkdownDescription: "AuthRequests defines the duration of time for which the AuthRequests will be valid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device_requests": schema.StringAttribute{
								Description:         "DeviceRequests defines the duration of time for which the DeviceRequests will be valid.",
								MarkdownDescription: "DeviceRequests defines the duration of time for which the DeviceRequests will be valid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"id_tokens": schema.StringAttribute{
								Description:         "IDTokens defines the duration of time for which the IdTokens will be valid.",
								MarkdownDescription: "IDTokens defines the duration of time for which the IdTokens will be valid.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"refresh_tokens": schema.SingleNestedAttribute{
								Description:         "RefreshTokens defines refresh tokens expiry policy.",
								MarkdownDescription: "RefreshTokens defines refresh tokens expiry policy.",
								Attributes: map[string]schema.Attribute{
									"absolute_lifetime": schema.StringAttribute{
										Description:         "AbsoluteLifetime defines the duration of time after which a refresh token will expire.",
										MarkdownDescription: "AbsoluteLifetime defines the duration of time after which a refresh token will expire.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_rotation": schema.BoolAttribute{
										Description:         "DisableRotation disables refresh token rotation.",
										MarkdownDescription: "DisableRotation disables refresh token rotation.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reuse_interval": schema.StringAttribute{
										Description:         "ReuseInterval defines the duration of time after which a refresh token can be reused.",
										MarkdownDescription: "ReuseInterval defines the duration of time after which a refresh token can be reused.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"valid_if_not_used_for": schema.StringAttribute{
										Description:         "ValidIfNotUsedFor defines the duration of time after which a refresh token will expire if not used.",
										MarkdownDescription: "ValidIfNotUsedFor defines the duration of time after which a refresh token will expire if not used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"signing_keys": schema.StringAttribute{
								Description:         "SigningKeys defines the duration of time after which the SigningKeys will be rotated.",
								MarkdownDescription: "SigningKeys defines the duration of time after which the SigningKeys will be rotated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"frontend": schema.SingleNestedAttribute{
						Description:         "Frontend holds the web server's frontend templates and asset configuration.",
						MarkdownDescription: "Frontend holds the web server's frontend templates and asset configuration.",
						Attributes: map[string]schema.Attribute{
							"dir": schema.StringAttribute{
								Description:         "Dir is a file path to static web assets.  It is expected to contain the following directories: * static - Static static served at '( issuer URL )/static'. * templates - HTML templates controlled by dex. * themes/(theme) - Static static served at '( issuer URL )/theme'.",
								MarkdownDescription: "Dir is a file path to static web assets.  It is expected to contain the following directories: * static - Static static served at '( issuer URL )/static'. * templates - HTML templates controlled by dex. * themes/(theme) - Static static served at '( issuer URL )/theme'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"issuer": schema.StringAttribute{
								Description:         "Issuer is the name of the issuer, used in the HTML templates. Defaults to 'dex'.",
								MarkdownDescription: "Issuer is the name of the issuer, used in the HTML templates. Defaults to 'dex'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logo_url": schema.StringAttribute{
								Description:         "LogoURL is the URL of the logo to use in the HTML templates. Defaults to '( issuer URL )/theme/logo.png'",
								MarkdownDescription: "LogoURL is the URL of the logo to use in the HTML templates. Defaults to '( issuer URL )/theme/logo.png'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"theme": schema.StringAttribute{
								Description:         "Theme is the name of the theme to use. Defaults to 'light'.",
								MarkdownDescription: "Theme is the name of the theme to use. Defaults to 'light'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"grpc": schema.SingleNestedAttribute{
						Description:         "GRPC holds configuration for the Dex API gRPC server.",
						MarkdownDescription: "GRPC holds configuration for the Dex API gRPC server.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an optional map of additional annotations to add to the Dex API gRPC service.",
								MarkdownDescription: "Annotations is an optional map of additional annotations to add to the Dex API gRPC service.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"certificate_secret_ref": schema.SingleNestedAttribute{
								Description:         "CertificateSecretRef is an optional reference to a secret containing the TLS certificate and key to use for the Dex API gRPC server.",
								MarkdownDescription: "CertificateSecretRef is an optional reference to a secret containing the TLS certificate and key to use for the Dex API gRPC server.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is the name of the secret.",
										MarkdownDescription: "Name is the name of the secret.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_ca_secret_ref": schema.SingleNestedAttribute{
								Description:         "ClientCASecretRef is an optional reference to a secret containing the client CA.",
								MarkdownDescription: "ClientCASecretRef is an optional reference to a secret containing the client CA.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is the name of the secret.",
										MarkdownDescription: "Name is the name of the secret.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"reflection": schema.BoolAttribute{
								Description:         "Reflection enables gRPC server reflection.",
								MarkdownDescription: "Reflection enables gRPC server reflection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image is the Dex image to use.",
						MarkdownDescription: "Image is the Dex image to use.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "Ingress is the optional ingress configuration for the Dex identity provider.",
						MarkdownDescription: "Ingress is the optional ingress configuration for the Dex identity provider.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an optional map of additional annotations to add to the ingress.",
								MarkdownDescription: "Annotations is an optional map of additional annotations to add to the ingress.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled enables ingress for the Dex identity provider.",
								MarkdownDescription: "Enabled enables ingress for the Dex identity provider.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"hosts": schema.ListNestedAttribute{
								Description:         "Hosts is a list of hosts and paths to route traffic to the Dex identity provider.",
								MarkdownDescription: "Hosts is a list of hosts and paths to route traffic to the Dex identity provider.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "Host is the host to route traffic to the Dex identity provider.",
											MarkdownDescription: "Host is the host to route traffic to the Dex identity provider.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"paths": schema.ListNestedAttribute{
											Description:         "Paths is a list of paths to route traffic to the Dex identity provider.",
											MarkdownDescription: "Paths is a list of paths to route traffic to the Dex identity provider.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "Path is matched against the path of an incoming request.",
														MarkdownDescription: "Path is matched against the path of an incoming request.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"path_type": schema.StringAttribute{
														Description:         "PathType determines the interpretation of the Path matching.",
														MarkdownDescription: "PathType determines the interpretation of the Path matching.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Exact", "Prefix", "ImplementationSpecific"),
														},
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"ingress_class_name": schema.StringAttribute{
								Description:         "IngressClassName is the optional ingress class to use for the Dex identity provider.",
								MarkdownDescription: "IngressClassName is the optional ingress class to use for the Dex identity provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.ListNestedAttribute{
								Description:         "TLS is an optional list of TLS configurations for the ingress.",
								MarkdownDescription: "TLS is an optional list of TLS configurations for the ingress.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hosts": schema.ListAttribute{
											Description:         "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
											MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
											MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
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

					"issuer": schema.StringAttribute{
						Description:         "Issuer is the base path of Dex and the external name of the OpenID Connect service. This is the canonical URL that all clients MUST use to refer to Dex.",
						MarkdownDescription: "Issuer is the base path of Dex and the external name of the OpenID Connect service. This is the canonical URL that all clients MUST use to refer to Dex.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"logger": schema.SingleNestedAttribute{
						Description:         "Logger holds configuration required to customize logging for dex.",
						MarkdownDescription: "Logger holds configuration required to customize logging for dex.",
						Attributes: map[string]schema.Attribute{
							"format": schema.StringAttribute{
								Description:         "Format specifies the format to be used for logging.",
								MarkdownDescription: "Format specifies the format to be used for logging.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"level": schema.StringAttribute{
								Description:         "Level sets logging level severity.",
								MarkdownDescription: "Level sets logging level severity.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics holds configuration for metrics.",
						MarkdownDescription: "Metrics holds configuration for metrics.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled enables Prometheus metric scraping.",
								MarkdownDescription: "Enabled enables Prometheus metric scraping.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"interval": schema.StringAttribute{
								Description:         "Interval at which metrics should be scraped If not specified Prometheus' global scrape interval is used.",
								MarkdownDescription: "Interval at which metrics should be scraped If not specified Prometheus' global scrape interval is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"oauth2": schema.SingleNestedAttribute{
						Description:         "OAuth2 holds configuration for OAuth2.",
						MarkdownDescription: "OAuth2 holds configuration for OAuth2.",
						Attributes: map[string]schema.Attribute{
							"always_show_login_screen": schema.BoolAttribute{
								Description:         "AlwaysShowLoginScreen, if specified, show the connector selection screen even if there's only one.",
								MarkdownDescription: "AlwaysShowLoginScreen, if specified, show the connector selection screen even if there's only one.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grant_types": schema.ListAttribute{
								Description:         "GrantTypes is a list of allowed grant types, defaults to all supported types.",
								MarkdownDescription: "GrantTypes is a list of allowed grant types, defaults to all supported types.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password_connector": schema.StringAttribute{
								Description:         "PasswordConnector is a specific connector to user for password grants.",
								MarkdownDescription: "PasswordConnector is a specific connector to user for password grants.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"response_types": schema.ListAttribute{
								Description:         "ResponseTypes is a list of allowed response types, defaults to all supported types.",
								MarkdownDescription: "ResponseTypes is a list of allowed response types, defaults to all supported types.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"skip_approval_screen": schema.BoolAttribute{
								Description:         "SkipApprovalScreen, if specified, do not prompt the user to approve client authorization. The act of logging in implies authorization.",
								MarkdownDescription: "SkipApprovalScreen, if specified, do not prompt the user to approve client authorization. The act of logging in implies authorization.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the optional number of replicas of the Dex identity provider pod to run. Only supported if using postgresql storage.",
						MarkdownDescription: "Replicas is the optional number of replicas of the Dex identity provider pod to run. Only supported if using postgresql storage.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources allows specifying the resource requirements for the Dex identity provider container.",
						MarkdownDescription: "Resources allows specifying the resource requirements for the Dex identity provider container.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage configures the storage for Dex.",
						MarkdownDescription: "Storage configures the storage for Dex.",
						Attributes: map[string]schema.Attribute{
							"postgres": schema.SingleNestedAttribute{
								Description:         "Postgres holds the configuration for the postgres storage type.",
								MarkdownDescription: "Postgres holds the configuration for the postgres storage type.",
								Attributes: map[string]schema.Attribute{
									"conn_max_lifetime": schema.StringAttribute{
										Description:         "ConnMaxLifetime is the maximum amount of time a connection may be reused (default forever).",
										MarkdownDescription: "ConnMaxLifetime is the maximum amount of time a connection may be reused (default forever).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connection_timeout": schema.StringAttribute{
										Description:         "ConnectionTimeout is the maximum amount of time to wait for a connection to become available.",
										MarkdownDescription: "ConnectionTimeout is the maximum amount of time to wait for a connection to become available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials_secret_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsSecretRef is a reference to a secret containing the username and password to use for authentication.",
										MarkdownDescription: "CredentialsSecretRef is a reference to a secret containing the username and password to use for authentication.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the name of the secret.",
												MarkdownDescription: "Name is the name of the secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"database": schema.StringAttribute{
										Description:         "Database is the name of the database to connect to.",
										MarkdownDescription: "Database is the name of the database to connect to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"host": schema.StringAttribute{
										Description:         "Host is the host to connect to.",
										MarkdownDescription: "Host is the host to connect to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"max_idle_conns": schema.Int64Attribute{
										Description:         "MaxIdleConns is the maximum number of connections in the idle connection pool (default 5).",
										MarkdownDescription: "MaxIdleConns is the maximum number of connections in the idle connection pool (default 5).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_open_conns": schema.Int64Attribute{
										Description:         "MaxOpenConns is the maximum number of open connections to the database (default 5).",
										MarkdownDescription: "MaxOpenConns is the maximum number of open connections to the database (default 5).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port is the port to connect to.",
										MarkdownDescription: "Port is the port to connect to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"ssl": schema.SingleNestedAttribute{
										Description:         "SSL holds optional TLS configuration for postgres.",
										MarkdownDescription: "SSL holds optional TLS configuration for postgres.",
										Attributes: map[string]schema.Attribute{
											"ca_secret_ref": schema.SingleNestedAttribute{
												Description:         "CASecretRef is an optional reference to a secret containing the CA certificate.",
												MarkdownDescription: "CASecretRef is an optional reference to a secret containing the CA certificate.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_certificate_secret_ref": schema.SingleNestedAttribute{
												Description:         "ClientCertificateSecretRef is an optional reference to a secret containing the client certificate and key.",
												MarkdownDescription: "ClientCertificateSecretRef is an optional reference to a secret containing the client certificate and key.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": schema.StringAttribute{
												Description:         "Mode is the SSL mode to use.",
												MarkdownDescription: "Mode is the SSL mode to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"server_name": schema.StringAttribute{
												Description:         "ServerName ensures that the certificate matches the given hostname the client is connecting to.",
												MarkdownDescription: "ServerName ensures that the certificate matches the given hostname the client is connecting to.",
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

							"sqlite3": schema.SingleNestedAttribute{
								Description:         "Sqlite3 holds the configuration for the sqlite3 storage type.",
								MarkdownDescription: "Sqlite3 holds the configuration for the sqlite3 storage type.",
								Attributes: map[string]schema.Attribute{
									"file": schema.StringAttribute{
										Description:         "File is the path to the sqlite3 database file.",
										MarkdownDescription: "File is the path to the sqlite3 database file.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the storage type to use.",
								MarkdownDescription: "Type is the storage type to use.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("memory", "sqlite3", "postgres"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"volume_claim_templates": schema.ListNestedAttribute{
						Description:         "VolumeClaimTemplates are volume claim templates for the Dex identity provider pod.",
						MarkdownDescription: "VolumeClaimTemplates are volume claim templates for the Dex identity provider pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
									MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
									MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metadata": schema.SingleNestedAttribute{
									Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
									MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"finalizers": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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

										"namespace": schema.StringAttribute{
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

								"spec": schema.SingleNestedAttribute{
									Description:         "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.ListAttribute{
											Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"data_source": schema.SingleNestedAttribute{
											Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
											MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
											Attributes: map[string]schema.Attribute{
												"api_group": schema.StringAttribute{
													Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
													MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind is the type of resource being referenced",
													MarkdownDescription: "Kind is the type of resource being referenced",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of resource being referenced",
													MarkdownDescription: "Name is the name of resource being referenced",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"data_source_ref": schema.SingleNestedAttribute{
											Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
											MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
											Attributes: map[string]schema.Attribute{
												"api_group": schema.StringAttribute{
													Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
													MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind is the type of resource being referenced",
													MarkdownDescription: "Kind is the type of resource being referenced",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of resource being referenced",
													MarkdownDescription: "Name is the name of resource being referenced",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
											MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

												"limits": schema.MapAttribute{
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

										"selector": schema.SingleNestedAttribute{
											Description:         "selector is a label query over volumes to consider for binding.",
											MarkdownDescription: "selector is a label query over volumes to consider for binding.",
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

										"storage_class_name": schema.StringAttribute{
											Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
											MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_mode": schema.StringAttribute{
											Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
											MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_name": schema.StringAttribute{
											Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
											MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"status": schema.SingleNestedAttribute{
									Description:         "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.ListAttribute{
											Description:         "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allocated_resource_statuses": schema.MapAttribute{
											Description:         "allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used.  ClaimResourceStatus can be in any of following states: - ControllerResizeInProgress: State set when resize controller starts resizing the volume in control-plane. - ControllerResizeFailed: State set when resize has failed in resize controller with a terminal error. - NodeResizePending: State set when resize controller has finished resizing the volume but further resizing of volume is needed on the node. - NodeResizeInProgress: State set when kubelet starts resizing the volume. - NodeResizeFailed: State set when resizing has failed in kubelet with a terminal error. Transient errors don't set NodeResizeFailed. For example: if expanding a PVC for more capacity - this field can be one of the following states: - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeFailed' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizePending' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeFailed' When this field is not set, it means that no resize operation is in progress for the given PVC.  A controller that receives PVC update with previously unknown resourceName or ClaimResourceStatus should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC.  This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											MarkdownDescription: "allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used.  ClaimResourceStatus can be in any of following states: - ControllerResizeInProgress: State set when resize controller starts resizing the volume in control-plane. - ControllerResizeFailed: State set when resize has failed in resize controller with a terminal error. - NodeResizePending: State set when resize controller has finished resizing the volume but further resizing of volume is needed on the node. - NodeResizeInProgress: State set when kubelet starts resizing the volume. - NodeResizeFailed: State set when resizing has failed in kubelet with a terminal error. Transient errors don't set NodeResizeFailed. For example: if expanding a PVC for more capacity - this field can be one of the following states: - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'ControllerResizeFailed' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizePending' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeInProgress' - pvc.status.allocatedResourceStatus['storage'] = 'NodeResizeFailed' When this field is not set, it means that no resize operation is in progress for the given PVC.  A controller that receives PVC update with previously unknown resourceName or ClaimResourceStatus should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC.  This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allocated_resources": schema.MapAttribute{
											Description:         "allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used.  Capacity reported here may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity.  A controller that receives PVC update with previously unknown resourceName should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC.  This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											MarkdownDescription: "allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax. Valid values are either: * Un-prefixed keys: - storage - the capacity of the volume. * Custom resources must use implementation-defined prefixed names such as 'example.com/my-custom-resource' Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used.  Capacity reported here may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity.  A controller that receives PVC update with previously unknown resourceName should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC.  This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"capacity": schema.MapAttribute{
											Description:         "capacity represents the actual resources of the underlying volume.",
											MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"conditions": schema.ListNestedAttribute{
											Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.",
											MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"last_probe_time": schema.StringAttribute{
														Description:         "lastProbeTime is the time we probed the condition.",
														MarkdownDescription: "lastProbeTime is the time we probed the condition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"last_transition_time": schema.StringAttribute{
														Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
														MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"message": schema.StringAttribute{
														Description:         "message is the human-readable message indicating details about last transition.",
														MarkdownDescription: "message is the human-readable message indicating details about last transition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"reason": schema.StringAttribute{
														Description:         "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized.",
														MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"status": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
														MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

										"phase": schema.StringAttribute{
											Description:         "phase represents the current phase of PersistentVolumeClaim.",
											MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts are volume mounts for the Dex identity provider container.",
						MarkdownDescription: "VolumeMounts are volume mounts for the Dex identity provider container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
									MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"mount_propagation": schema.StringAttribute{
									Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
									MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "This must match the Name of a Volume.",
									MarkdownDescription: "This must match the Name of a Volume.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"read_only": schema.BoolAttribute{
									Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
									MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path": schema.StringAttribute{
									Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
									MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path_expr": schema.StringAttribute{
									Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
									MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

					"web": schema.SingleNestedAttribute{
						Description:         "Web holds configuration for the web server.",
						MarkdownDescription: "Web holds configuration for the web server.",
						Attributes: map[string]schema.Attribute{
							"allowed_origins": schema.ListAttribute{
								Description:         "AllowedOrigins is a list of allowed origins for CORS requests.",
								MarkdownDescription: "AllowedOrigins is a list of allowed origins for CORS requests.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"annotations": schema.MapAttribute{
								Description:         "Annotations is an optional map of additional annotations to add to the web service.",
								MarkdownDescription: "Annotations is an optional map of additional annotations to add to the web service.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"certificate_secret_ref": schema.SingleNestedAttribute{
								Description:         "CertificateSecretRef is an optional reference to a secret containing the TLS certificate and key to use for HTTPS.",
								MarkdownDescription: "CertificateSecretRef is an optional reference to a secret containing the TLS certificate and key to use for HTTPS.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name is the name of the secret.",
										MarkdownDescription: "Name is the name of the secret.",
										Required:            true,
										Optional:            false,
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *DexGpuNinjaComDexIdentityProviderV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest")

	var model DexGpuNinjaComDexIdentityProviderV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("dex.gpu-ninja.com/v1alpha1")
	model.Kind = pointer.String("DexIdentityProvider")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
