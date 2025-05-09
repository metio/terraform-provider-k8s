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
	_ datasource.DataSource = &RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest{}
}

type RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest struct{}

type RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1ManifestData struct {
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
		GCEalias       *string `tfsdk:"gc_ealias" json:"GCEalias,omitempty"`
		GCEmetadata    *string `tfsdk:"gc_emetadata" json:"GCEmetadata,omitempty"`
		GCPCredentials *struct {
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
		} `tfsdk:"gcp_credentials" json:"GCPCredentials,omitempty"`
		IAMalias       *string `tfsdk:"ia_malias" json:"IAMalias,omitempty"`
		IAMmetadata    *string `tfsdk:"ia_mmetadata" json:"IAMmetadata,omitempty"`
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
		CustomEndpoint *map[string]string `tfsdk:"custom_endpoint" json:"customEndpoint,omitempty"`
		Path           *string            `tfsdk:"path" json:"path,omitempty"`
		ServiceAccount *string            `tfsdk:"service_account" json:"serviceAccount,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_gcp_auth_engine_config_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GCPAuthEngineConfig is the Schema for the gcpauthengineconfigs API",
		MarkdownDescription: "GCPAuthEngineConfig is the Schema for the gcpauthengineconfigs API",
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
				Description:         "GCPAuthEngineConfigSpec defines the desired state of GCPAuthEngineConfig",
				MarkdownDescription: "GCPAuthEngineConfigSpec defines the desired state of GCPAuthEngineConfig",
				Attributes: map[string]schema.Attribute{
					"gc_ealias": schema.StringAttribute{
						Description:         "Must be either instance_id or role_id. If instance_id is specified, the GCE instance ID will be used for alias names during login. If role_id is specified, the ID of the Vault role will be used. Only used if role type is gce.",
						MarkdownDescription: "Must be either instance_id or role_id. If instance_id is specified, the GCE instance ID will be used for alias names during login. If role_id is specified, the ID of the Vault role will be used. Only used if role type is gce.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gc_emetadata": schema.StringAttribute{
						Description:         "The metadata to include on the token returned by the login endpoint. This metadata will be added to both audit logs, and on the gce_alias. By default, it includes instance_creation_timestamp, instance_id, instance_name, project_id, project_number, role, service_account_id, service_account_email, and zone. To include no metadata, set to '' via the CLI or [] via the API. To use only particular fields, select the explicit fields. To restore to defaults, send only a field of default. Only select fields that will have a low rate of change for your gce_alias because each change triggers a storage write and can have a performance impact at scale. Only used if role type is gce.",
						MarkdownDescription: "The metadata to include on the token returned by the login endpoint. This metadata will be added to both audit logs, and on the gce_alias. By default, it includes instance_creation_timestamp, instance_id, instance_name, project_id, project_number, role, service_account_id, service_account_email, and zone. To include no metadata, set to '' via the CLI or [] via the API. To use only particular fields, select the explicit fields. To restore to defaults, send only a field of default. Only select fields that will have a low rate of change for your gce_alias because each change triggers a storage write and can have a performance impact at scale. Only used if role type is gce.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gcp_credentials": schema.SingleNestedAttribute{
						Description:         "GCPCredentials in JSON string containing the contents of a GCP service account credentials file.",
						MarkdownDescription: "GCPCredentials in JSON string containing the contents of a GCP service account credentials file.",
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

					"ia_malias": schema.StringAttribute{
						Description:         "Must be either unique_id or role_id. If unique_id is specified, the service account's unique ID will be used for alias names during login. If role_id is specified, the ID of the Vault role will be used. Only used if role type is iam.",
						MarkdownDescription: "Must be either unique_id or role_id. If unique_id is specified, the service account's unique ID will be used for alias names during login. If role_id is specified, the ID of the Vault role will be used. Only used if role type is iam.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ia_mmetadata": schema.StringAttribute{
						Description:         "The metadata to include on the token returned by the login endpoint. This metadata will be added to both audit logs, and on the iam_alias. By default, it includes project_id, role, service_account_id, and service_account_email. To include no metadata, set to '' via the CLI or [] via the API. To use only particular fields, select the explicit fields. To restore to defaults, send only a field of default. Only select fields that will have a low rate of change for your iam_alias because each change triggers a storage write and can have a performance impact at scale. Only used if role type is iam.",
						MarkdownDescription: "The metadata to include on the token returned by the login endpoint. This metadata will be added to both audit logs, and on the iam_alias. By default, it includes project_id, role, service_account_id, and service_account_email. To include no metadata, set to '' via the CLI or [] via the API. To use only particular fields, select the explicit fields. To restore to defaults, send only a field of default. Only select fields that will have a low rate of change for your iam_alias because each change triggers a storage write and can have a performance impact at scale. Only used if role type is iam.",
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

					"custom_endpoint": schema.MapAttribute{
						Description:         "Specifies overrides to service endpoints used when making API requests. This allows specific requests made during authentication to target alternative service endpoints for use in Private Google Access environments. Overrides are set at the subdomain level using the following keys: api - Replaces the service endpoint used in API requests to https://www.googleapis.com. iam - Replaces the service endpoint used in API requests to https://iam.googleapis.com. crm - Replaces the service endpoint used in API requests to https://cloudresourcemanager.googleapis.com. compute - Replaces the service endpoint used in API requests to https://compute.googleapis.com. The endpoint value provided for a given key has the form of scheme://host:port. The scheme:// and :port portions of the endpoint value are optional.",
						MarkdownDescription: "Specifies overrides to service endpoints used when making API requests. This allows specific requests made during authentication to target alternative service endpoints for use in Private Google Access environments. Overrides are set at the subdomain level using the following keys: api - Replaces the service endpoint used in API requests to https://www.googleapis.com. iam - Replaces the service endpoint used in API requests to https://iam.googleapis.com. crm - Replaces the service endpoint used in API requests to https://cloudresourcemanager.googleapis.com. compute - Replaces the service endpoint used in API requests to https://compute.googleapis.com. The endpoint value provided for a given key has the form of scheme://host:port. The scheme:// and :port portions of the endpoint value are optional.",
						ElementType:         types.StringType,
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

					"service_account": schema.StringAttribute{
						Description:         "Service Account Name. A service account is a special kind of account typically used by an application or compute workload, such as a Compute Engine instance, rather than a person. A service account is identified by its email address, which is unique to the account. Applications use service accounts to make authorized API calls by authenticating as either the service account itself, or as Google Workspace or Cloud Identity users through domain-wide delegation. When an application authenticates as a service account, it has access to all resources that the service account has permission to access.",
						MarkdownDescription: "Service Account Name. A service account is a special kind of account typically used by an application or compute workload, such as a Compute Engine instance, rather than a person. A service account is identified by its email address, which is unique to the account. Applications use service accounts to make authorized API calls by authenticating as either the service account itself, or as Google Workspace or Cloud Identity users through domain-wide delegation. When an application authenticates as a service account, it has access to all resources that the service account has permission to access.",
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

func (r *RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_gcp_auth_engine_config_v1alpha1_manifest")

	var model RedhatcopRedhatIoGcpauthEngineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("GCPAuthEngineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
