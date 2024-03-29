/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoOauthV1Manifest{}
)

func NewConfigOpenshiftIoOauthV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoOauthV1Manifest{}
}

type ConfigOpenshiftIoOauthV1Manifest struct{}

type ConfigOpenshiftIoOauthV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		IdentityProviders *[]struct {
			BasicAuth *struct {
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				TlsClientCert *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"tls_client_cert" json:"tlsClientCert,omitempty"`
				TlsClientKey *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"tls_client_key" json:"tlsClientKey,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			Github *struct {
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				ClientID     *string `tfsdk:"client_id" json:"clientID,omitempty"`
				ClientSecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				Hostname      *string   `tfsdk:"hostname" json:"hostname,omitempty"`
				Organizations *[]string `tfsdk:"organizations" json:"organizations,omitempty"`
				Teams         *[]string `tfsdk:"teams" json:"teams,omitempty"`
			} `tfsdk:"github" json:"github,omitempty"`
			Gitlab *struct {
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				ClientID     *string `tfsdk:"client_id" json:"clientID,omitempty"`
				ClientSecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"gitlab" json:"gitlab,omitempty"`
			Google *struct {
				ClientID     *string `tfsdk:"client_id" json:"clientID,omitempty"`
				ClientSecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				HostedDomain *string `tfsdk:"hosted_domain" json:"hostedDomain,omitempty"`
			} `tfsdk:"google" json:"google,omitempty"`
			Htpasswd *struct {
				FileData *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"file_data" json:"fileData,omitempty"`
			} `tfsdk:"htpasswd" json:"htpasswd,omitempty"`
			Keystone *struct {
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				DomainName    *string `tfsdk:"domain_name" json:"domainName,omitempty"`
				TlsClientCert *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"tls_client_cert" json:"tlsClientCert,omitempty"`
				TlsClientKey *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"tls_client_key" json:"tlsClientKey,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"keystone" json:"keystone,omitempty"`
			Ldap *struct {
				Attributes *struct {
					Email             *[]string `tfsdk:"email" json:"email,omitempty"`
					Id                *[]string `tfsdk:"id" json:"id,omitempty"`
					Name              *[]string `tfsdk:"name" json:"name,omitempty"`
					PreferredUsername *[]string `tfsdk:"preferred_username" json:"preferredUsername,omitempty"`
				} `tfsdk:"attributes" json:"attributes,omitempty"`
				BindDN       *string `tfsdk:"bind_dn" json:"bindDN,omitempty"`
				BindPassword *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"bind_password" json:"bindPassword,omitempty"`
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				Insecure *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"ldap" json:"ldap,omitempty"`
			MappingMethod *string `tfsdk:"mapping_method" json:"mappingMethod,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			OpenID        *struct {
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				Claims *struct {
					Email             *[]string `tfsdk:"email" json:"email,omitempty"`
					Groups            *[]string `tfsdk:"groups" json:"groups,omitempty"`
					Name              *[]string `tfsdk:"name" json:"name,omitempty"`
					PreferredUsername *[]string `tfsdk:"preferred_username" json:"preferredUsername,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				ClientID     *string `tfsdk:"client_id" json:"clientID,omitempty"`
				ClientSecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				ExtraAuthorizeParameters *map[string]string `tfsdk:"extra_authorize_parameters" json:"extraAuthorizeParameters,omitempty"`
				ExtraScopes              *[]string          `tfsdk:"extra_scopes" json:"extraScopes,omitempty"`
				Issuer                   *string            `tfsdk:"issuer" json:"issuer,omitempty"`
			} `tfsdk:"open_id" json:"openID,omitempty"`
			RequestHeader *struct {
				Ca *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				ChallengeURL             *string   `tfsdk:"challenge_url" json:"challengeURL,omitempty"`
				ClientCommonNames        *[]string `tfsdk:"client_common_names" json:"clientCommonNames,omitempty"`
				EmailHeaders             *[]string `tfsdk:"email_headers" json:"emailHeaders,omitempty"`
				Headers                  *[]string `tfsdk:"headers" json:"headers,omitempty"`
				LoginURL                 *string   `tfsdk:"login_url" json:"loginURL,omitempty"`
				NameHeaders              *[]string `tfsdk:"name_headers" json:"nameHeaders,omitempty"`
				PreferredUsernameHeaders *[]string `tfsdk:"preferred_username_headers" json:"preferredUsernameHeaders,omitempty"`
			} `tfsdk:"request_header" json:"requestHeader,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"identity_providers" json:"identityProviders,omitempty"`
		Templates *struct {
			Error *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"error" json:"error,omitempty"`
			Login *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"login" json:"login,omitempty"`
			ProviderSelection *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"provider_selection" json:"providerSelection,omitempty"`
		} `tfsdk:"templates" json:"templates,omitempty"`
		TokenConfig *struct {
			AccessTokenInactivityTimeout        *string `tfsdk:"access_token_inactivity_timeout" json:"accessTokenInactivityTimeout,omitempty"`
			AccessTokenInactivityTimeoutSeconds *int64  `tfsdk:"access_token_inactivity_timeout_seconds" json:"accessTokenInactivityTimeoutSeconds,omitempty"`
			AccessTokenMaxAgeSeconds            *int64  `tfsdk:"access_token_max_age_seconds" json:"accessTokenMaxAgeSeconds,omitempty"`
		} `tfsdk:"token_config" json:"tokenConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoOauthV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_o_auth_v1_manifest"
}

func (r *ConfigOpenshiftIoOauthV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OAuth holds cluster-wide information about OAuth.  The canonical name is 'cluster'. It is used to configure the integrated OAuth server. This configuration is only honored when the top level Authentication config has type set to IntegratedOAuth.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "OAuth holds cluster-wide information about OAuth.  The canonical name is 'cluster'. It is used to configure the integrated OAuth server. This configuration is only honored when the top level Authentication config has type set to IntegratedOAuth.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"identity_providers": schema.ListNestedAttribute{
						Description:         "identityProviders is an ordered list of ways for a user to identify themselves. When this list is empty, no identities are provisioned for users.",
						MarkdownDescription: "identityProviders is an ordered list of ways for a user to identify themselves. When this list is empty, no identities are provisioned for users.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"basic_auth": schema.SingleNestedAttribute{
									Description:         "basicAuth contains configuration options for the BasicAuth IdP",
									MarkdownDescription: "basicAuth contains configuration options for the BasicAuth IdP",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls_client_cert": schema.SingleNestedAttribute{
											Description:         "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls_client_key": schema.SingleNestedAttribute{
											Description:         "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"url": schema.StringAttribute{
											Description:         "url is the remote URL to connect to",
											MarkdownDescription: "url is the remote URL to connect to",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"github": schema.SingleNestedAttribute{
									Description:         "github enables user authentication using GitHub credentials",
									MarkdownDescription: "github enables user authentication using GitHub credentials",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. This can only be configured when hostname is set to a non-empty value. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. This can only be configured when hostname is set to a non-empty value. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"hostname": schema.StringAttribute{
											Description:         "hostname is the optional domain (e.g. 'mycompany.com') for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value configured at /setup/settings#hostname.",
											MarkdownDescription: "hostname is the optional domain (e.g. 'mycompany.com') for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value configured at /setup/settings#hostname.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"organizations": schema.ListAttribute{
											Description:         "organizations optionally restricts which organizations are allowed to log in",
											MarkdownDescription: "organizations optionally restricts which organizations are allowed to log in",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"teams": schema.ListAttribute{
											Description:         "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",
											MarkdownDescription: "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",
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

								"gitlab": schema.SingleNestedAttribute{
									Description:         "gitlab enables user authentication using GitLab credentials",
									MarkdownDescription: "gitlab enables user authentication using GitLab credentials",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"url": schema.StringAttribute{
											Description:         "url is the oauth server base URL",
											MarkdownDescription: "url is the oauth server base URL",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"google": schema.SingleNestedAttribute{
									Description:         "google enables user authentication using Google credentials",
									MarkdownDescription: "google enables user authentication using Google credentials",
									Attributes: map[string]schema.Attribute{
										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"hosted_domain": schema.StringAttribute{
											Description:         "hostedDomain is the optional Google App domain (e.g. 'mycompany.com') to restrict logins to",
											MarkdownDescription: "hostedDomain is the optional Google App domain (e.g. 'mycompany.com') to restrict logins to",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"htpasswd": schema.SingleNestedAttribute{
									Description:         "htpasswd enables user authentication using an HTPasswd file to validate credentials",
									MarkdownDescription: "htpasswd enables user authentication using an HTPasswd file to validate credentials",
									Attributes: map[string]schema.Attribute{
										"file_data": schema.SingleNestedAttribute{
											Description:         "fileData is a required reference to a secret by name containing the data to use as the htpasswd file. The key 'htpasswd' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. If the specified htpasswd data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "fileData is a required reference to a secret by name containing the data to use as the htpasswd file. The key 'htpasswd' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. If the specified htpasswd data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"keystone": schema.SingleNestedAttribute{
									Description:         "keystone enables user authentication using keystone password credentials",
									MarkdownDescription: "keystone enables user authentication using keystone password credentials",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"domain_name": schema.StringAttribute{
											Description:         "domainName is required for keystone v3",
											MarkdownDescription: "domainName is required for keystone v3",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_client_cert": schema.SingleNestedAttribute{
											Description:         "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls_client_key": schema.SingleNestedAttribute{
											Description:         "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"url": schema.StringAttribute{
											Description:         "url is the remote URL to connect to",
											MarkdownDescription: "url is the remote URL to connect to",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ldap": schema.SingleNestedAttribute{
									Description:         "ldap enables user authentication using LDAP credentials",
									MarkdownDescription: "ldap enables user authentication using LDAP credentials",
									Attributes: map[string]schema.Attribute{
										"attributes": schema.SingleNestedAttribute{
											Description:         "attributes maps LDAP attributes to identities",
											MarkdownDescription: "attributes maps LDAP attributes to identities",
											Attributes: map[string]schema.Attribute{
												"email": schema.ListAttribute{
													Description:         "email is the list of attributes whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
													MarkdownDescription: "email is the list of attributes whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.ListAttribute{
													Description:         "id is the list of attributes whose values should be used as the user ID. Required. First non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. LDAP standard identity attribute is 'dn'",
													MarkdownDescription: "id is the list of attributes whose values should be used as the user ID. Required. First non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. LDAP standard identity attribute is 'dn'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.ListAttribute{
													Description:         "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is 'cn'",
													MarkdownDescription: "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is 'cn'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"preferred_username": schema.ListAttribute{
													Description:         "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is 'uid'",
													MarkdownDescription: "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is 'uid'",
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

										"bind_dn": schema.StringAttribute{
											Description:         "bindDN is an optional DN to bind with during the search phase.",
											MarkdownDescription: "bindDN is an optional DN to bind with during the search phase.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bind_password": schema.SingleNestedAttribute{
											Description:         "bindPassword is an optional reference to a secret by name containing a password to bind with during the search phase. The key 'bindPassword' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "bindPassword is an optional reference to a secret by name containing a password to bind with during the search phase. The key 'bindPassword' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"insecure": schema.BoolAttribute{
											Description:         "insecure, if true, indicates the connection should not use TLS WARNING: Should not be set to 'true' with the URL scheme 'ldaps://' as 'ldaps://' URLs always attempt to connect using TLS, even when 'insecure' is set to 'true' When 'true', 'ldap://' URLS connect insecurely. When 'false', 'ldap://' URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830.",
											MarkdownDescription: "insecure, if true, indicates the connection should not use TLS WARNING: Should not be set to 'true' with the URL scheme 'ldaps://' as 'ldaps://' URLs always attempt to connect using TLS, even when 'insecure' is set to 'true' When 'true', 'ldap://' URLS connect insecurely. When 'false', 'ldap://' URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"url": schema.StringAttribute{
											Description:         "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is: ldap://host:port/basedn?attribute?scope?filter",
											MarkdownDescription: "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is: ldap://host:port/basedn?attribute?scope?filter",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"mapping_method": schema.StringAttribute{
									Description:         "mappingMethod determines how identities from this provider are mapped to users Defaults to 'claim'",
									MarkdownDescription: "mappingMethod determines how identities from this provider are mapped to users Defaults to 'claim'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "name is used to qualify the identities returned by this provider. - It MUST be unique and not shared by any other identity provider used - It MUST be a valid path segment: name cannot equal '.' or '..' or contain '/' or '%' or ':' Ref: https://godoc.org/github.com/openshift/origin/pkg/user/apis/user/validation#ValidateIdentityProviderName",
									MarkdownDescription: "name is used to qualify the identities returned by this provider. - It MUST be unique and not shared by any other identity provider used - It MUST be a valid path segment: name cannot equal '.' or '..' or contain '/' or '%' or ':' Ref: https://godoc.org/github.com/openshift/origin/pkg/user/apis/user/validation#ValidateIdentityProviderName",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"open_id": schema.SingleNestedAttribute{
									Description:         "openID enables user authentication using OpenID credentials",
									MarkdownDescription: "openID enables user authentication using OpenID credentials",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"claims": schema.SingleNestedAttribute{
											Description:         "claims mappings",
											MarkdownDescription: "claims mappings",
											Attributes: map[string]schema.Attribute{
												"email": schema.ListAttribute{
													Description:         "email is the list of claims whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
													MarkdownDescription: "email is the list of claims whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"groups": schema.ListAttribute{
													Description:         "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user. If multiple claims are specified, the first one with a non-empty value is used.",
													MarkdownDescription: "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user. If multiple claims are specified, the first one with a non-empty value is used.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.ListAttribute{
													Description:         "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",
													MarkdownDescription: "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"preferred_username": schema.ListAttribute{
													Description:         "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the sub claim",
													MarkdownDescription: "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the sub claim",
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

										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"extra_authorize_parameters": schema.MapAttribute{
											Description:         "extraAuthorizeParameters are any custom parameters to add to the authorize request.",
											MarkdownDescription: "extraAuthorizeParameters are any custom parameters to add to the authorize request.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"extra_scopes": schema.ListAttribute{
											Description:         "extraScopes are any scopes to request in addition to the standard 'openid' scope.",
											MarkdownDescription: "extraScopes are any scopes to request in addition to the standard 'openid' scope.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"issuer": schema.StringAttribute{
											Description:         "issuer is the URL that the OpenID Provider asserts as its Issuer Identifier. It must use the https scheme with no query or fragment component.",
											MarkdownDescription: "issuer is the URL that the OpenID Provider asserts as its Issuer Identifier. It must use the https scheme with no query or fragment component.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"request_header": schema.SingleNestedAttribute{
									Description:         "requestHeader enables user authentication using request header credentials",
									MarkdownDescription: "requestHeader enables user authentication using request header credentials",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "ca is a required reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. Specifically, it allows verification of incoming requests to prevent header spoofing. The key 'ca.crt' is used to locate the data. If the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is a required reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. Specifically, it allows verification of incoming requests to prevent header spoofing. The key 'ca.crt' is used to locate the data. If the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"challenge_url": schema.StringAttribute{
											Description:         "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here. ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when challenge is set to true.",
											MarkdownDescription: "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here. ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when challenge is set to true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"client_common_names": schema.ListAttribute{
											Description:         "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",
											MarkdownDescription: "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"email_headers": schema.ListAttribute{
											Description:         "emailHeaders is the set of headers to check for the email address",
											MarkdownDescription: "emailHeaders is the set of headers to check for the email address",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"headers": schema.ListAttribute{
											Description:         "headers is the set of headers to check for identity information",
											MarkdownDescription: "headers is the set of headers to check for identity information",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"login_url": schema.StringAttribute{
											Description:         "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when login is set to true.",
											MarkdownDescription: "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when login is set to true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name_headers": schema.ListAttribute{
											Description:         "nameHeaders is the set of headers to check for the display name",
											MarkdownDescription: "nameHeaders is the set of headers to check for the display name",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"preferred_username_headers": schema.ListAttribute{
											Description:         "preferredUsernameHeaders is the set of headers to check for the preferred username",
											MarkdownDescription: "preferredUsernameHeaders is the set of headers to check for the preferred username",
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

								"type": schema.StringAttribute{
									Description:         "type identifies the identity provider type for this entry.",
									MarkdownDescription: "type identifies the identity provider type for this entry.",
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

					"templates": schema.SingleNestedAttribute{
						Description:         "templates allow you to customize pages like the login page.",
						MarkdownDescription: "templates allow you to customize pages like the login page.",
						Attributes: map[string]schema.Attribute{
							"error": schema.SingleNestedAttribute{
								Description:         "error is the name of a secret that specifies a go template to use to render error pages during the authentication or grant flow. The key 'errors.html' is used to locate the template data. If specified and the secret or expected key is not found, the default error page is used. If the specified template is not valid, the default error page is used. If unspecified, the default error page is used. The namespace for this secret is openshift-config.",
								MarkdownDescription: "error is the name of a secret that specifies a go template to use to render error pages during the authentication or grant flow. The key 'errors.html' is used to locate the template data. If specified and the secret or expected key is not found, the default error page is used. If the specified template is not valid, the default error page is used. If unspecified, the default error page is used. The namespace for this secret is openshift-config.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the metadata.name of the referenced secret",
										MarkdownDescription: "name is the metadata.name of the referenced secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"login": schema.SingleNestedAttribute{
								Description:         "login is the name of a secret that specifies a go template to use to render the login page. The key 'login.html' is used to locate the template data. If specified and the secret or expected key is not found, the default login page is used. If the specified template is not valid, the default login page is used. If unspecified, the default login page is used. The namespace for this secret is openshift-config.",
								MarkdownDescription: "login is the name of a secret that specifies a go template to use to render the login page. The key 'login.html' is used to locate the template data. If specified and the secret or expected key is not found, the default login page is used. If the specified template is not valid, the default login page is used. If unspecified, the default login page is used. The namespace for this secret is openshift-config.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the metadata.name of the referenced secret",
										MarkdownDescription: "name is the metadata.name of the referenced secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider_selection": schema.SingleNestedAttribute{
								Description:         "providerSelection is the name of a secret that specifies a go template to use to render the provider selection page. The key 'providers.html' is used to locate the template data. If specified and the secret or expected key is not found, the default provider selection page is used. If the specified template is not valid, the default provider selection page is used. If unspecified, the default provider selection page is used. The namespace for this secret is openshift-config.",
								MarkdownDescription: "providerSelection is the name of a secret that specifies a go template to use to render the provider selection page. The key 'providers.html' is used to locate the template data. If specified and the secret or expected key is not found, the default provider selection page is used. If the specified template is not valid, the default provider selection page is used. If unspecified, the default provider selection page is used. The namespace for this secret is openshift-config.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the metadata.name of the referenced secret",
										MarkdownDescription: "name is the metadata.name of the referenced secret",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"token_config": schema.SingleNestedAttribute{
						Description:         "tokenConfig contains options for authorization and access tokens",
						MarkdownDescription: "tokenConfig contains options for authorization and access tokens",
						Attributes: map[string]schema.Attribute{
							"access_token_inactivity_timeout": schema.StringAttribute{
								Description:         "accessTokenInactivityTimeout defines the token inactivity timeout for tokens granted by any client. The value represents the maximum amount of time that can occur between consecutive uses of the token. Tokens become invalid if they are not used within this temporal window. The user will need to acquire a new token to regain access once a token times out. Takes valid time duration string such as '5m', '1.5h' or '2h45m'. The minimum allowed value for duration is 300s (5 minutes). If the timeout is configured per client, then that value takes precedence. If the timeout value is not specified and the client does not override the value, then tokens are valid until their lifetime.  WARNING: existing tokens' timeout will not be affected (lowered) by changing this value",
								MarkdownDescription: "accessTokenInactivityTimeout defines the token inactivity timeout for tokens granted by any client. The value represents the maximum amount of time that can occur between consecutive uses of the token. Tokens become invalid if they are not used within this temporal window. The user will need to acquire a new token to regain access once a token times out. Takes valid time duration string such as '5m', '1.5h' or '2h45m'. The minimum allowed value for duration is 300s (5 minutes). If the timeout is configured per client, then that value takes precedence. If the timeout value is not specified and the client does not override the value, then tokens are valid until their lifetime.  WARNING: existing tokens' timeout will not be affected (lowered) by changing this value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_token_inactivity_timeout_seconds": schema.Int64Attribute{
								Description:         "accessTokenInactivityTimeoutSeconds - DEPRECATED: setting this field has no effect.",
								MarkdownDescription: "accessTokenInactivityTimeoutSeconds - DEPRECATED: setting this field has no effect.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"access_token_max_age_seconds": schema.Int64Attribute{
								Description:         "accessTokenMaxAgeSeconds defines the maximum age of access tokens",
								MarkdownDescription: "accessTokenMaxAgeSeconds defines the maximum age of access tokens",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigOpenshiftIoOauthV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_o_auth_v1_manifest")

	var model ConfigOpenshiftIoOauthV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("OAuth")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
