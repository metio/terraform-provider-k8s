/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package redhatcop_redhat_io_v1alpha1

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
	_ datasource.DataSource = &RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest{}
}

type RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest struct{}

type RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1ManifestData struct {
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
		TLSMaxVersion        *string `tfsdk:"tls_max_version" json:"TLSMaxVersion,omitempty"`
		TLSMinVersion        *string `tfsdk:"tls_min_version" json:"TLSMinVersion,omitempty"`
		UPNDomain            *string `tfsdk:"upn_domain" json:"UPNDomain,omitempty"`
		AnonymousGroupSearch *bool   `tfsdk:"anonymous_group_search" json:"anonymousGroupSearch,omitempty"`
		Authentication       *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		BindCredentials *struct {
			PasswordKey  *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
			RandomSecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"random_secret" json:"randomSecret,omitempty"`
			Secret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			VaultSecret *struct {
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"vault_secret" json:"vaultSecret,omitempty"`
		} `tfsdk:"bind_credentials" json:"bindCredentials,omitempty"`
		BindDN             *string `tfsdk:"bind_dn" json:"bindDN,omitempty"`
		CaseSensitiveNames *bool   `tfsdk:"case_sensitive_names" json:"caseSensitiveNames,omitempty"`
		Certificate        *string `tfsdk:"certificate" json:"certificate,omitempty"`
		ClientTLSCert      *string `tfsdk:"client_tls_cert" json:"clientTLSCert,omitempty"`
		ClientTLSKey       *string `tfsdk:"client_tls_key" json:"clientTLSKey,omitempty"`
		Connection         *struct {
			Address    *string `tfsdk:"address" json:"address,omitempty"`
			MaxRetries *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			TLSConfig  *struct {
				Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
				SkipVerify *bool   `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
				TlsSecret  *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
				TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
			} `tfsdk:"t_ls_config" json:"tLSConfig,omitempty"`
			TimeOut *string `tfsdk:"time_out" json:"timeOut,omitempty"`
		} `tfsdk:"connection" json:"connection,omitempty"`
		DenyNullBind   *bool   `tfsdk:"deny_null_bind" json:"denyNullBind,omitempty"`
		DiscoverDN     *bool   `tfsdk:"discover_dn" json:"discoverDN,omitempty"`
		GroupAttr      *string `tfsdk:"group_attr" json:"groupAttr,omitempty"`
		GroupDN        *string `tfsdk:"group_dn" json:"groupDN,omitempty"`
		GroupFilter    *string `tfsdk:"group_filter" json:"groupFilter,omitempty"`
		InsecureTLS    *bool   `tfsdk:"insecure_tls" json:"insecureTLS,omitempty"`
		Path           *string `tfsdk:"path" json:"path,omitempty"`
		RequestTimeout *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
		StartTLS       *bool   `tfsdk:"start_tls" json:"startTLS,omitempty"`
		TLSConfig      *struct {
			Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
			SkipVerify *bool   `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
			TlsSecret  *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
			TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
		} `tfsdk:"t_ls_config" json:"tLSConfig,omitempty"`
		TokenBoundCIDRs      *string `tfsdk:"token_bound_cidrs" json:"tokenBoundCIDRs,omitempty"`
		TokenExplicitMaxTTL  *string `tfsdk:"token_explicit_max_ttl" json:"tokenExplicitMaxTTL,omitempty"`
		TokenMaxTTL          *string `tfsdk:"token_max_ttl" json:"tokenMaxTTL,omitempty"`
		TokenNoDefaultPolicy *bool   `tfsdk:"token_no_default_policy" json:"tokenNoDefaultPolicy,omitempty"`
		TokenNumUses         *int64  `tfsdk:"token_num_uses" json:"tokenNumUses,omitempty"`
		TokenPeriod          *int64  `tfsdk:"token_period" json:"tokenPeriod,omitempty"`
		TokenPolicies        *string `tfsdk:"token_policies" json:"tokenPolicies,omitempty"`
		TokenTTL             *string `tfsdk:"token_ttl" json:"tokenTTL,omitempty"`
		TokenType            *string `tfsdk:"token_type" json:"tokenType,omitempty"`
		Url                  *string `tfsdk:"url" json:"url,omitempty"`
		UserAttr             *string `tfsdk:"user_attr" json:"userAttr,omitempty"`
		UserDN               *string `tfsdk:"user_dn" json:"userDN,omitempty"`
		UserFilter           *string `tfsdk:"user_filter" json:"userFilter,omitempty"`
		UsernameAsAlias      *bool   `tfsdk:"username_as_alias" json:"usernameAsAlias,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_ldap_auth_engine_config_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LDAPAuthEngineConfig is the Schema for the ldapauthengineconfigs API",
		MarkdownDescription: "LDAPAuthEngineConfig is the Schema for the ldapauthengineconfigs API",
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
				Description:         "LDAPAuthEngineConfigSpec defines the desired state of LDAPAuthEngineConfig",
				MarkdownDescription: "LDAPAuthEngineConfigSpec defines the desired state of LDAPAuthEngineConfig",
				Attributes: map[string]schema.Attribute{
					"tls_max_version": schema.StringAttribute{
						Description:         "TLSMaxVersion Maximum TLS version to use. Accepted values are tls10, tls11, tls12 or tls13",
						MarkdownDescription: "TLSMaxVersion Maximum TLS version to use. Accepted values are tls10, tls11, tls12 or tls13",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_min_version": schema.StringAttribute{
						Description:         "TLSMinVersion Minimum TLS version to use. Accepted values are tls10, tls11, tls12 or tls13",
						MarkdownDescription: "TLSMinVersion Minimum TLS version to use. Accepted values are tls10, tls11, tls12 or tls13",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"upn_domain": schema.StringAttribute{
						Description:         "UPNDomain The userPrincipalDomain used to construct the UPN string for the authenticating user. The constructed UPN will appear as [username]@UPNDomain. Example: example.com, which will cause vault to bind as username@example.com",
						MarkdownDescription: "UPNDomain The userPrincipalDomain used to construct the UPN string for the authenticating user. The constructed UPN will appear as [username]@UPNDomain. Example: example.com, which will cause vault to bind as username@example.com",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"anonymous_group_search": schema.BoolAttribute{
						Description:         "AnonymousGroupSearch Use anonymous binds when performing LDAP group searches (note: even when true, the initial credentials will still be used for the initial connection test).",
						MarkdownDescription: "AnonymousGroupSearch Use anonymous binds when performing LDAP group searches (note: even when true, the initial credentials will still be used for the initial connection test).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"authentication": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								Description:         "Namespace is the Vault namespace to be used in all the operations withing this connection/authentication. Only available in Vault Enterprise.",
								MarkdownDescription: "Namespace is the Vault namespace to be used in all the operations withing this connection/authentication. Only available in Vault Enterprise.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the path of the role used for this kube auth authentication. The operator will try to authenticate at {[namespace/]}auth/{spec.path}",
								MarkdownDescription: "Path is the path of the role used for this kube auth authentication. The operator will try to authenticate at {[namespace/]}auth/{spec.path}",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
								},
							},

							"role": schema.StringAttribute{
								Description:         "Role the role to be used during authentication",
								MarkdownDescription: "Role the role to be used during authentication",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account": schema.SingleNestedAttribute{
								Description:         "ServiceAccount is the service account used for the kube auth authentication",
								MarkdownDescription: "ServiceAccount is the service account used for the kube auth authentication",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

					"bind_credentials": schema.SingleNestedAttribute{
						Description:         "BindCredentials is used to connect to the LDAP service on the specified LDAP Server. BindCredentials consists in bindDN and bindPass, which can be created as Kubernetes Secret, VaultSecret or RandomSecret.",
						MarkdownDescription: "BindCredentials is used to connect to the LDAP service on the specified LDAP Server. BindCredentials consists in bindDN and bindPass, which can be created as Kubernetes Secret, VaultSecret or RandomSecret.",
						Attributes: map[string]schema.Attribute{
							"password_key": schema.StringAttribute{
								Description:         "PasswordKey key to be used when retrieving the password, required with VaultSecrets and Kubernetes secrets, ignored with RandomSecret",
								MarkdownDescription: "PasswordKey key to be used when retrieving the password, required with VaultSecrets and Kubernetes secrets, ignored with RandomSecret",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"random_secret": schema.SingleNestedAttribute{
								Description:         "RandomSecret retrieves the credentials from the Vault secret corresponding to this RandomSecret. This will map the 'username' and 'password' keys of the secret to the username and password of this config. All other keys will be ignored. If the RandomSecret is refreshed the operator retrieves the new secret from Vault and updates this configuration. Only one of RootCredentialsFromVaultSecret or RootCredentialsFromSecret or RootCredentialsFromRandomSecret can be specified. When using randomSecret a username must be specified in the spec.username password: Specifies the password to use when connecting with the username. This value will not be returned by Vault when performing a read upon the configuration. This is typically used in the connection_url field via the templating directive '{{'password'}}''.",
								MarkdownDescription: "RandomSecret retrieves the credentials from the Vault secret corresponding to this RandomSecret. This will map the 'username' and 'password' keys of the secret to the username and password of this config. All other keys will be ignored. If the RandomSecret is refreshed the operator retrieves the new secret from Vault and updates this configuration. Only one of RootCredentialsFromVaultSecret or RootCredentialsFromSecret or RootCredentialsFromRandomSecret can be specified. When using randomSecret a username must be specified in the spec.username password: Specifies the password to use when connecting with the username. This value will not be returned by Vault when performing a read upon the configuration. This is typically used in the connection_url field via the templating directive '{{'password'}}''.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": schema.SingleNestedAttribute{
								Description:         "Secret retrieves the credentials from a Kubernetes secret. The secret must be of basicauth type (https://kubernetes.io/docs/concepts/configuration/secret/#basic-authentication-secret). This will map the 'username' and 'password' keys of the secret to the username and password of this config. If the kubernetes secret is updated, this configuration will also be updated. All other keys will be ignored. Only one of RootCredentialsFromVaultSecret or RootCredentialsFromSecret or RootCredentialsFromRandomSecret can be specified. username: Specifies the name of the user to use as the 'root' user when connecting to the database. This 'root' user is used to create/update/delete users managed by these plugins, so you will need to ensure that this user has permissions to manipulate users appropriate to the database. This is typically used in the connection_url field via the templating directive '{{'username'}}' or '{{'name'}}'. password: Specifies the password to use when connecting with the username. This value will not be returned by Vault when performing a read upon the configuration. This is typically used in the connection_url field via the templating directive '{{'password'}}'. If username is provided as spec.username, it takes precedence over the username retrieved from the referenced secret",
								MarkdownDescription: "Secret retrieves the credentials from a Kubernetes secret. The secret must be of basicauth type (https://kubernetes.io/docs/concepts/configuration/secret/#basic-authentication-secret). This will map the 'username' and 'password' keys of the secret to the username and password of this config. If the kubernetes secret is updated, this configuration will also be updated. All other keys will be ignored. Only one of RootCredentialsFromVaultSecret or RootCredentialsFromSecret or RootCredentialsFromRandomSecret can be specified. username: Specifies the name of the user to use as the 'root' user when connecting to the database. This 'root' user is used to create/update/delete users managed by these plugins, so you will need to ensure that this user has permissions to manipulate users appropriate to the database. This is typically used in the connection_url field via the templating directive '{{'username'}}' or '{{'name'}}'. password: Specifies the password to use when connecting with the username. This value will not be returned by Vault when performing a read upon the configuration. This is typically used in the connection_url field via the templating directive '{{'password'}}'. If username is provided as spec.username, it takes precedence over the username retrieved from the referenced secret",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"username_key": schema.StringAttribute{
								Description:         "UsernameKey key to be used when retrieving the username, optional with VaultSecrets and Kubernetes secrets, ignored with RandomSecret",
								MarkdownDescription: "UsernameKey key to be used when retrieving the username, optional with VaultSecrets and Kubernetes secrets, ignored with RandomSecret",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vault_secret": schema.SingleNestedAttribute{
								Description:         "VaultSecret retrieves the credentials from a Vault secret. This will map the 'username' and 'password' keys of the secret to the username and password of this config. All other keys will be ignored. Only one of RootCredentialsFromVaultSecret or RootCredentialsFromSecret or RootCredentialsFromRandomSecret can be specified. username: Specifies the name of the user to use as the 'root' user when connecting to the database. This 'root' user is used to create/update/delete users managed by these plugins, so you will need to ensure that this user has permissions to manipulate users appropriate to the database. This is typically used in the connection_url field via the templating directive '{{'username'}}' or '{{'name'}}'. password: Specifies the password to use when connecting with the username. This value will not be returned by Vault when performing a read upon the configuration. This is typically used in the connection_url field via the templating directive '{{'password'}}'. If username is provided as spec.username, it takes precedence over the username retrieved from the referenced secret",
								MarkdownDescription: "VaultSecret retrieves the credentials from a Vault secret. This will map the 'username' and 'password' keys of the secret to the username and password of this config. All other keys will be ignored. Only one of RootCredentialsFromVaultSecret or RootCredentialsFromSecret or RootCredentialsFromRandomSecret can be specified. username: Specifies the name of the user to use as the 'root' user when connecting to the database. This 'root' user is used to create/update/delete users managed by these plugins, so you will need to ensure that this user has permissions to manipulate users appropriate to the database. This is typically used in the connection_url field via the templating directive '{{'username'}}' or '{{'name'}}'. password: Specifies the password to use when connecting with the username. This value will not be returned by Vault when performing a read upon the configuration. This is typically used in the connection_url field via the templating directive '{{'password'}}'. If username is provided as spec.username, it takes precedence over the username retrieved from the referenced secret",
								Attributes: map[string]schema.Attribute{
									"path": schema.StringAttribute{
										Description:         "Path is the path to the secret",
										MarkdownDescription: "Path is the path to the secret",
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

					"bind_dn": schema.StringAttribute{
						Description:         "BindDN - Username used to connect to the LDAP service on the specified LDAP Server. If in the form accountname@domain.com, the username is transformed into a proper LDAP bind DN, for example, CN=accountname,CN=users,DC=domain,DC=com, when accessing the LDAP server. If username is provided it takes precedence over the username retrieved from the referenced secrets",
						MarkdownDescription: "BindDN - Username used to connect to the LDAP service on the specified LDAP Server. If in the form accountname@domain.com, the username is transformed into a proper LDAP bind DN, for example, CN=accountname,CN=users,DC=domain,DC=com, when accessing the LDAP server. If username is provided it takes precedence over the username retrieved from the referenced secrets",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"case_sensitive_names": schema.BoolAttribute{
						Description:         "CaseSensitiveNames If set, user and group names assigned to policies within the backend will be case sensitive. Otherwise, names will be normalized to lower case. Case will still be preserved when sending the username to the LDAP server at login time; this is only for matching local user/group definitions.",
						MarkdownDescription: "CaseSensitiveNames If set, user and group names assigned to policies within the backend will be case sensitive. Otherwise, names will be normalized to lower case. Case will still be preserved when sending the username to the LDAP server at login time; this is only for matching local user/group definitions.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"certificate": schema.StringAttribute{
						Description:         "Certificate CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.",
						MarkdownDescription: "Certificate CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_tls_cert": schema.StringAttribute{
						Description:         "ClientTLSCert Client certificate to provide to the LDAP server, must be x509 PEM encoded",
						MarkdownDescription: "ClientTLSCert Client certificate to provide to the LDAP server, must be x509 PEM encoded",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_tls_key": schema.StringAttribute{
						Description:         "ClientTLSKey Client certificate key to provide to the LDAP server, must be x509 PEM encoded",
						MarkdownDescription: "ClientTLSKey Client certificate key to provide to the LDAP server, must be x509 PEM encoded",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"connection": schema.SingleNestedAttribute{
						Description:         "Connection represents the information needed to connect to Vault. This operator uses the standard Vault environment variables to connect to Vault. If you need to override those settings and for example connect to a different Vault instance, you can do with this section of the CR.",
						MarkdownDescription: "Connection represents the information needed to connect to Vault. This operator uses the standard Vault environment variables to connect to Vault. If you need to override those settings and for example connect to a different Vault instance, you can do with this section of the CR.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Address Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/",
								MarkdownDescription: "Address Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_retries": schema.Int64Attribute{
								Description:         "MaxRetries Maximum number of retries when certain error codes are encountered. The default is 2, for three total attempts. Set this to 0 or less to disable retrying. Error codes that are retried are 412 (client consistency requirement not satisfied) and all 5xx except for 501 (not implemented).",
								MarkdownDescription: "MaxRetries Maximum number of retries when certain error codes are encountered. The default is 2, for three total attempts. Set this to 0 or less to disable retrying. Error codes that are retried are 412 (client consistency requirement not satisfied) and all 5xx except for 501 (not implemented).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"t_ls_config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cacert": schema.StringAttribute{
										Description:         "Cacert Path to a PEM-encoded CA certificate file on the local disk. This file is used to verify the Vault server's SSL certificate. This environment variable takes precedence over a cert passed via the secret.",
										MarkdownDescription: "Cacert Path to a PEM-encoded CA certificate file on the local disk. This file is used to verify the Vault server's SSL certificate. This environment variable takes precedence over a cert passed via the secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"skip_verify": schema.BoolAttribute{
										Description:         "SkipVerify Do not verify Vault's presented certificate before communicating with it. Setting this variable is not recommended and voids Vault's security model.",
										MarkdownDescription: "SkipVerify Do not verify Vault's presented certificate before communicating with it. Setting this variable is not recommended and voids Vault's security model.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_secret": schema.SingleNestedAttribute{
										Description:         "TLSSecret namespace-local secret containing the tls material for the connection. the expected keys for the secret are: ca bundle -> 'ca.crt', certificate -> 'tls.crt', key -> 'tls.key'",
										MarkdownDescription: "TLSSecret namespace-local secret containing the tls material for the connection. the expected keys for the secret are: ca bundle -> 'ca.crt', certificate -> 'tls.crt', key -> 'tls.key'",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_server_name": schema.StringAttribute{
										Description:         "TLSServerName Name to use as the SNI host when connecting via TLS.",
										MarkdownDescription: "TLSServerName Name to use as the SNI host when connecting via TLS.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"time_out": schema.StringAttribute{
								Description:         "Timeout Timeout variable. The default value is 60s.",
								MarkdownDescription: "Timeout Timeout variable. The default value is 60s.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"deny_null_bind": schema.BoolAttribute{
						Description:         "DenyNullBind This option prevents users from bypassing authentication when providing an empty password",
						MarkdownDescription: "DenyNullBind This option prevents users from bypassing authentication when providing an empty password",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"discover_dn": schema.BoolAttribute{
						Description:         "DiscoverDN Use anonymous bind to discover the bind DN of a user.",
						MarkdownDescription: "DiscoverDN Use anonymous bind to discover the bind DN of a user.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"group_attr": schema.StringAttribute{
						Description:         "GroupAttr LDAP attribute to follow on objects returned by groupfilter in order to enumerate user group membership. Examples: for groupfilter queries returning group objects, use: cn. For queries returning user objects, use: memberOf. The default is cn.",
						MarkdownDescription: "GroupAttr LDAP attribute to follow on objects returned by groupfilter in order to enumerate user group membership. Examples: for groupfilter queries returning group objects, use: cn. For queries returning user objects, use: memberOf. The default is cn.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"group_dn": schema.StringAttribute{
						Description:         "GroupDN LDAP search base to use for group membership search. This can be the root containing either groups or users. Example: ou=Groups,dc=example,dc=com",
						MarkdownDescription: "GroupDN LDAP search base to use for group membership search. This can be the root containing either groups or users. Example: ou=Groups,dc=example,dc=com",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"group_filter": schema.StringAttribute{
						Description:         "GroupFilter Go template used when constructing the group membership query. The template can access the following context variables: [UserDN, Username]. The default is (|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}})), which is compatible with several common directory schemas. To support nested group resolution for Active Directory, instead use the following query: (&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
						MarkdownDescription: "GroupFilter Go template used when constructing the group membership query. The template can access the following context variables: [UserDN, Username]. The default is (|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}})), which is compatible with several common directory schemas. To support nested group resolution for Active Directory, instead use the following query: (&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"insecure_tls": schema.BoolAttribute{
						Description:         "InsecureTLS If true, skips LDAP server SSL certificate verification - insecure, use with caution!",
						MarkdownDescription: "InsecureTLS If true, skips LDAP server SSL certificate verification - insecure, use with caution!",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/auth/{spec.path}/config/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						MarkdownDescription: "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/auth/{spec.path}/config/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"request_timeout": schema.StringAttribute{
						Description:         "RequestTimeout Timeout, in seconds, for the connection when making requests against the server before returning back an error.",
						MarkdownDescription: "RequestTimeout Timeout, in seconds, for the connection when making requests against the server before returning back an error.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"start_tls": schema.BoolAttribute{
						Description:         "StartTLS If true, issues a StartTLS command after establishing an unencrypted connection.",
						MarkdownDescription: "StartTLS If true, issues a StartTLS command after establishing an unencrypted connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"t_ls_config": schema.SingleNestedAttribute{
						Description:         "CertificateConfig represents the LDAP service certificate configuration. CertificateConfig consists in certificate, clientTLSCert and clientTLSKey which can be consumed from an Kubernetes Secret.",
						MarkdownDescription: "CertificateConfig represents the LDAP service certificate configuration. CertificateConfig consists in certificate, clientTLSCert and clientTLSKey which can be consumed from an Kubernetes Secret.",
						Attributes: map[string]schema.Attribute{
							"cacert": schema.StringAttribute{
								Description:         "Cacert Path to a PEM-encoded CA certificate file on the local disk. This file is used to verify the Vault server's SSL certificate. This environment variable takes precedence over a cert passed via the secret.",
								MarkdownDescription: "Cacert Path to a PEM-encoded CA certificate file on the local disk. This file is used to verify the Vault server's SSL certificate. This environment variable takes precedence over a cert passed via the secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"skip_verify": schema.BoolAttribute{
								Description:         "SkipVerify Do not verify Vault's presented certificate before communicating with it. Setting this variable is not recommended and voids Vault's security model.",
								MarkdownDescription: "SkipVerify Do not verify Vault's presented certificate before communicating with it. Setting this variable is not recommended and voids Vault's security model.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_secret": schema.SingleNestedAttribute{
								Description:         "TLSSecret namespace-local secret containing the tls material for the connection. the expected keys for the secret are: ca bundle -> 'ca.crt', certificate -> 'tls.crt', key -> 'tls.key'",
								MarkdownDescription: "TLSSecret namespace-local secret containing the tls material for the connection. the expected keys for the secret are: ca bundle -> 'ca.crt', certificate -> 'tls.crt', key -> 'tls.key'",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_server_name": schema.StringAttribute{
								Description:         "TLSServerName Name to use as the SNI host when connecting via TLS.",
								MarkdownDescription: "TLSServerName Name to use as the SNI host when connecting via TLS.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"token_bound_cidrs": schema.StringAttribute{
						Description:         "TokenBoundCIDRs List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.",
						MarkdownDescription: "TokenBoundCIDRs List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_explicit_max_ttl": schema.StringAttribute{
						Description:         "TonenExplicitMaxTTL If set, will encode an explicit max TTL onto the token. This is a hard cap even if token_ttl and token_max_ttl would otherwise allow a renewal.",
						MarkdownDescription: "TonenExplicitMaxTTL If set, will encode an explicit max TTL onto the token. This is a hard cap even if token_ttl and token_max_ttl would otherwise allow a renewal.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_max_ttl": schema.StringAttribute{
						Description:         "TokenMaxTTL The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time",
						MarkdownDescription: "TokenMaxTTL The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_no_default_policy": schema.BoolAttribute{
						Description:         "TokenNoDefaultPolicy If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.",
						MarkdownDescription: "TokenNoDefaultPolicy If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_num_uses": schema.Int64Attribute{
						Description:         "TokenNumUses The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0.",
						MarkdownDescription: "TokenNumUses The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_period": schema.Int64Attribute{
						Description:         "TokenPeriod The period, if any, to set on the token",
						MarkdownDescription: "TokenPeriod The period, if any, to set on the token",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_policies": schema.StringAttribute{
						Description:         "TokenPolicies List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.",
						MarkdownDescription: "TokenPolicies List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_ttl": schema.StringAttribute{
						Description:         "TokenTTL The incremental lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						MarkdownDescription: "TokenTTL The incremental lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_type": schema.StringAttribute{
						Description:         "The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time.",
						MarkdownDescription: "The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"url": schema.StringAttribute{
						Description:         "URL The LDAP server to connect to. Examples: ldap://ldap.myorg.com, ldaps://ldap.myorg.com:636. Multiple URLs can be specified with commas, e.g. ldap://ldap.myorg.com,ldap://ldap2.myorg.com; these will be tried in-order.",
						MarkdownDescription: "URL The LDAP server to connect to. Examples: ldap://ldap.myorg.com, ldaps://ldap.myorg.com:636. Multiple URLs can be specified with commas, e.g. ldap://ldap.myorg.com,ldap://ldap2.myorg.com; these will be tried in-order.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"user_attr": schema.StringAttribute{
						Description:         "UserAttr Attribute on user attribute object matching the username passed when authenticating. Examples: sAMAccountName, cn, uid",
						MarkdownDescription: "UserAttr Attribute on user attribute object matching the username passed when authenticating. Examples: sAMAccountName, cn, uid",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_dn": schema.StringAttribute{
						Description:         "UserDN Base DN under which to perform user search. Example: ou=Users,dc=example,dc=com",
						MarkdownDescription: "UserDN Base DN under which to perform user search. Example: ou=Users,dc=example,dc=com",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_filter": schema.StringAttribute{
						Description:         "UserFilter An optional LDAP user search filter. The template can access the following context variables: UserAttr, Username. The default is ({{.UserAttr}}={{.Username}}), or ({{.UserAttr}}={{.Username@.upndomain}}) if upndomain is set.",
						MarkdownDescription: "UserFilter An optional LDAP user search filter. The template can access the following context variables: UserAttr, Username. The default is ({{.UserAttr}}={{.Username}}), or ({{.UserAttr}}={{.Username@.upndomain}}) if upndomain is set.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username_as_alias": schema.BoolAttribute{
						Description:         "UsernameAsAlias If set to true, forces the auth method to use the username passed by the user as the alias name.",
						MarkdownDescription: "UsernameAsAlias If set to true, forces the auth method to use the username passed by the user as the alias name.",
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

func (r *RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_ldap_auth_engine_config_v1alpha1_manifest")

	var model RedhatcopRedhatIoLdapauthEngineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("LDAPAuthEngineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
