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
	_ datasource.DataSource = &RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest{}
}

type RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest struct{}

type RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1ManifestData struct {
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
		AllowedRoles   *[]string `tfsdk:"allowed_roles" json:"allowedRoles,omitempty"`
		Authentication *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		Connection *struct {
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
		ConnectionURL          *string            `tfsdk:"connection_url" json:"connectionURL,omitempty"`
		DatabaseSpecificConfig *map[string]string `tfsdk:"database_specific_config" json:"databaseSpecificConfig,omitempty"`
		DisableEscaping        *bool              `tfsdk:"disable_escaping" json:"disableEscaping,omitempty"`
		Name                   *string            `tfsdk:"name" json:"name,omitempty"`
		PasswordAuthentication *string            `tfsdk:"password_authentication" json:"passwordAuthentication,omitempty"`
		PasswordPolicy         *string            `tfsdk:"password_policy" json:"passwordPolicy,omitempty"`
		Path                   *string            `tfsdk:"path" json:"path,omitempty"`
		PluginName             *string            `tfsdk:"plugin_name" json:"pluginName,omitempty"`
		PluginVersion          *string            `tfsdk:"plugin_version" json:"pluginVersion,omitempty"`
		RootCredentials        *struct {
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
		} `tfsdk:"root_credentials" json:"rootCredentials,omitempty"`
		RootPasswordRotation *struct {
			Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
			RotationPeriod *string `tfsdk:"rotation_period" json:"rotationPeriod,omitempty"`
		} `tfsdk:"root_password_rotation" json:"rootPasswordRotation,omitempty"`
		RootRotationStatements *[]string `tfsdk:"root_rotation_statements" json:"rootRotationStatements,omitempty"`
		Username               *string   `tfsdk:"username" json:"username,omitempty"`
		VerifyConnection       *bool     `tfsdk:"verify_connection" json:"verifyConnection,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_database_secret_engine_config_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DatabaseSecretEngineConfig is the Schema for the databasesecretengineconfigs API",
		MarkdownDescription: "DatabaseSecretEngineConfig is the Schema for the databasesecretengineconfigs API",
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
				Description:         "DatabaseSecretEngineConfigSpec defines the desired state of DatabaseSecretEngineConfig",
				MarkdownDescription: "DatabaseSecretEngineConfigSpec defines the desired state of DatabaseSecretEngineConfig",
				Attributes: map[string]schema.Attribute{
					"allowed_roles": schema.ListAttribute{
						Description:         "AllowedRoles List of the roles allowed to use this connection. Defaults to empty (no roles), if contains a '*' any role can use this connection. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "AllowedRoles List of the roles allowed to use this connection. Defaults to empty (no roles), if contains a '*' any role can use this connection. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"authentication": schema.SingleNestedAttribute{
						Description:         "Authentication is the kube auth configuration to be used to execute this request",
						MarkdownDescription: "Authentication is the kube auth configuration to be used to execute this request",
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

					"connection_url": schema.StringAttribute{
						Description:         "ConnectionURL Specifies the connection string used to connect to the database. Some plugins use url rather than connection_url. This allows for simple templating of the username and password of the root user. Typically, this is done by including a '{{'username'}}', '{{'name'}}', and/or '{{'password'}}' field within the string. These fields are typically be replaced with the values in the username and password fields.",
						MarkdownDescription: "ConnectionURL Specifies the connection string used to connect to the database. Some plugins use url rather than connection_url. This allows for simple templating of the username and password of the root user. Typically, this is done by including a '{{'username'}}', '{{'name'}}', and/or '{{'password'}}' field within the string. These fields are typically be replaced with the values in the username and password fields.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"database_specific_config": schema.MapAttribute{
						Description:         "DatabaseSpecificConfig this are the configuration specific to each database type",
						MarkdownDescription: "DatabaseSpecificConfig this are the configuration specific to each database type",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_escaping": schema.BoolAttribute{
						Description:         "DisableEscaping Determines whether special characters in the username and password fields will be escaped. Useful for alternate connection string formats like ADO. More information regarding this parameter can be found on the databases secrets engine docs. Defaults to false",
						MarkdownDescription: "DisableEscaping Determines whether special characters in the username and password fields will be escaped. Useful for alternate connection string formats like ADO. More information regarding this parameter can be found on the databases secrets engine docs. Defaults to false",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the obejct created in Vault. If this is specified it takes precedence over {metatada.name}",
						MarkdownDescription: "The name of the obejct created in Vault. If this is specified it takes precedence over {metatada.name}",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
						},
					},

					"password_authentication": schema.StringAttribute{
						Description:         "PasswordAuthentication When set to 'scram-sha-256', passwords will be hashed by Vault and stored as-is by PostgreSQL. Using 'scram-sha-256' requires a minimum version of PostgreSQL 10. Available options are 'scram-sha-256' and 'password'. The default is 'password'. When set to 'password', passwords will be sent to PostgreSQL in plaintext format and may appear in PostgreSQL logs as-is.",
						MarkdownDescription: "PasswordAuthentication When set to 'scram-sha-256', passwords will be hashed by Vault and stored as-is by PostgreSQL. Using 'scram-sha-256' requires a minimum version of PostgreSQL 10. Available options are 'scram-sha-256' and 'password'. The default is 'password'. When set to 'password', passwords will be sent to PostgreSQL in plaintext format and may appear in PostgreSQL logs as-is.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("password", "scram-sha-256"),
						},
					},

					"password_policy": schema.StringAttribute{
						Description:         "PasswordPolicy The name of the password policy to use when generating passwords for this database. If not specified, this will use a default policy defined as: 20 characters with at least 1 uppercase, 1 lowercase, 1 number, and 1 dash character.",
						MarkdownDescription: "PasswordPolicy The name of the password policy to use when generating passwords for this database. If not specified, this will use a default policy defined as: 20 characters with at least 1 uppercase, 1 lowercase, 1 number, and 1 dash character.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/config/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						MarkdownDescription: "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/config/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"plugin_name": schema.StringAttribute{
						Description:         "PluginName Specifies the name of the plugin to use for this connection.",
						MarkdownDescription: "PluginName Specifies the name of the plugin to use for this connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"plugin_version": schema.StringAttribute{
						Description:         "PluginVersion Specifies the semantic version of the plugin to use for this connection.",
						MarkdownDescription: "PluginVersion Specifies the semantic version of the plugin to use for this connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"root_credentials": schema.SingleNestedAttribute{
						Description:         "RootCredentials specifies how to retrieve the credentials for this DatabaseEngine connection.",
						MarkdownDescription: "RootCredentials specifies how to retrieve the credentials for this DatabaseEngine connection.",
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

					"root_password_rotation": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "Enabled whether the toot password should be rotated with the rotation statement. If set to true the root password will be rotated immediately.",
								MarkdownDescription: "Enabled whether the toot password should be rotated with the rotation statement. If set to true the root password will be rotated immediately.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rotation_period": schema.StringAttribute{
								Description:         "RotationPeriod if this value is set, the root password will be rotated approximately with teh requested frequency.",
								MarkdownDescription: "RotationPeriod if this value is set, the root password will be rotated approximately with teh requested frequency.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"root_rotation_statements": schema.ListAttribute{
						Description:         "RootRotationStatements Specifies the database statements to be executed to rotate the root user's credentials. See the plugin's API page for more information on support and formatting for this parameter. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "RootRotationStatements Specifies the database statements to be executed to rotate the root user's credentials. See the plugin's API page for more information on support and formatting for this parameter. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username": schema.StringAttribute{
						Description:         "Username Specifies the name of the user to use as the 'root' user when connecting to the database. This 'root' user is used to create/update/delete users managed by these plugins, so you will need to ensure that this user has permissions to manipulate users appropriate to the database. This is typically used in the connection_url field via the templating directive '{{'username'}}' or '{{'name'}}' If username is provided it takes precedence over the username retrieved from the referenced secrets",
						MarkdownDescription: "Username Specifies the name of the user to use as the 'root' user when connecting to the database. This 'root' user is used to create/update/delete users managed by these plugins, so you will need to ensure that this user has permissions to manipulate users appropriate to the database. This is typically used in the connection_url field via the templating directive '{{'username'}}' or '{{'name'}}' If username is provided it takes precedence over the username retrieved from the referenced secrets",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"verify_connection": schema.BoolAttribute{
						Description:         "VerifyConnection Specifies if the connection is verified during initial configuration. Defaults to true.",
						MarkdownDescription: "VerifyConnection Specifies if the connection is verified during initial configuration. Defaults to true.",
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

func (r *RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_database_secret_engine_config_v1alpha1_manifest")

	var model RedhatcopRedhatIoDatabaseSecretEngineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("DatabaseSecretEngineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
