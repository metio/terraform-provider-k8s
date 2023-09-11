/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &HiveOpenshiftIoSyncIdentityProviderV1DataSource{}
	_ datasource.DataSourceWithConfigure = &HiveOpenshiftIoSyncIdentityProviderV1DataSource{}
)

func NewHiveOpenshiftIoSyncIdentityProviderV1DataSource() datasource.DataSource {
	return &HiveOpenshiftIoSyncIdentityProviderV1DataSource{}
}

type HiveOpenshiftIoSyncIdentityProviderV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HiveOpenshiftIoSyncIdentityProviderV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterDeploymentRefs *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cluster_deployment_refs" json:"clusterDeploymentRefs,omitempty"`
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_sync_identity_provider_v1"
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SyncIdentityProvider is the Schema for the SyncIdentityProvider API",
		MarkdownDescription: "SyncIdentityProvider is the Schema for the SyncIdentityProvider API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "SyncIdentityProviderSpec defines the SyncIdentityProviderCommonSpec identity providers to sync along with ClusterDeploymentRefs indicating which clusters the SyncIdentityProvider applies to in the SyncIdentityProvider's namespace.",
				MarkdownDescription: "SyncIdentityProviderSpec defines the SyncIdentityProviderCommonSpec identity providers to sync along with ClusterDeploymentRefs indicating which clusters the SyncIdentityProvider applies to in the SyncIdentityProvider's namespace.",
				Attributes: map[string]schema.Attribute{
					"cluster_deployment_refs": schema.ListNestedAttribute{
						Description:         "ClusterDeploymentRefs is the list of LocalObjectReference indicating which clusters the SyncSet applies to in the SyncSet's namespace.",
						MarkdownDescription: "ClusterDeploymentRefs is the list of LocalObjectReference indicating which clusters the SyncSet applies to in the SyncSet's namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"identity_providers": schema.ListNestedAttribute{
						Description:         "IdentityProviders is an ordered list of ways for a user to identify themselves",
						MarkdownDescription: "IdentityProviders is an ordered list of ways for a user to identify themselves",
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"tls_client_cert": schema.SingleNestedAttribute{
											Description:         "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"tls_client_key": schema.SingleNestedAttribute{
											Description:         "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"url": schema.StringAttribute{
											Description:         "url is the remote URL to connect to",
											MarkdownDescription: "url is the remote URL to connect to",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"hostname": schema.StringAttribute{
											Description:         "hostname is the optional domain (e.g. 'mycompany.com') for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value configured at /setup/settings#hostname.",
											MarkdownDescription: "hostname is the optional domain (e.g. 'mycompany.com') for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value configured at /setup/settings#hostname.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"organizations": schema.ListAttribute{
											Description:         "organizations optionally restricts which organizations are allowed to log in",
											MarkdownDescription: "organizations optionally restricts which organizations are allowed to log in",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"teams": schema.ListAttribute{
											Description:         "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",
											MarkdownDescription: "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"url": schema.StringAttribute{
											Description:         "url is the oauth server base URL",
											MarkdownDescription: "url is the oauth server base URL",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"google": schema.SingleNestedAttribute{
									Description:         "google enables user authentication using Google credentials",
									MarkdownDescription: "google enables user authentication using Google credentials",
									Attributes: map[string]schema.Attribute{
										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"hosted_domain": schema.StringAttribute{
											Description:         "hostedDomain is the optional Google App domain (e.g. 'mycompany.com') to restrict logins to",
											MarkdownDescription: "hostedDomain is the optional Google App domain (e.g. 'mycompany.com') to restrict logins to",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"domain_name": schema.StringAttribute{
											Description:         "domainName is required for keystone v3",
											MarkdownDescription: "domainName is required for keystone v3",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tls_client_cert": schema.SingleNestedAttribute{
											Description:         "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"tls_client_key": schema.SingleNestedAttribute{
											Description:         "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"url": schema.StringAttribute{
											Description:         "url is the remote URL to connect to",
											MarkdownDescription: "url is the remote URL to connect to",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
													Optional:            false,
													Computed:            true,
												},

												"id": schema.ListAttribute{
													Description:         "id is the list of attributes whose values should be used as the user ID. Required. First non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. LDAP standard identity attribute is 'dn'",
													MarkdownDescription: "id is the list of attributes whose values should be used as the user ID. Required. First non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. LDAP standard identity attribute is 'dn'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.ListAttribute{
													Description:         "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is 'cn'",
													MarkdownDescription: "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is 'cn'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"preferred_username": schema.ListAttribute{
													Description:         "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is 'uid'",
													MarkdownDescription: "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is 'uid'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"bind_dn": schema.StringAttribute{
											Description:         "bindDN is an optional DN to bind with during the search phase.",
											MarkdownDescription: "bindDN is an optional DN to bind with during the search phase.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"bind_password": schema.SingleNestedAttribute{
											Description:         "bindPassword is an optional reference to a secret by name containing a password to bind with during the search phase. The key 'bindPassword' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "bindPassword is an optional reference to a secret by name containing a password to bind with during the search phase. The key 'bindPassword' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"ca": schema.SingleNestedAttribute{
											Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced config map",
													MarkdownDescription: "name is the metadata.name of the referenced config map",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"insecure": schema.BoolAttribute{
											Description:         "insecure, if true, indicates the connection should not use TLS WARNING: Should not be set to 'true' with the URL scheme 'ldaps://' as 'ldaps://' URLs always attempt to connect using TLS, even when 'insecure' is set to 'true' When 'true', 'ldap://' URLS connect insecurely. When 'false', 'ldap://' URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830.",
											MarkdownDescription: "insecure, if true, indicates the connection should not use TLS WARNING: Should not be set to 'true' with the URL scheme 'ldaps://' as 'ldaps://' URLs always attempt to connect using TLS, even when 'insecure' is set to 'true' When 'true', 'ldap://' URLS connect insecurely. When 'false', 'ldap://' URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"url": schema.StringAttribute{
											Description:         "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is: ldap://host:port/basedn?attribute?scope?filter",
											MarkdownDescription: "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is: ldap://host:port/basedn?attribute?scope?filter",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"mapping_method": schema.StringAttribute{
									Description:         "mappingMethod determines how identities from this provider are mapped to users Defaults to 'claim'",
									MarkdownDescription: "mappingMethod determines how identities from this provider are mapped to users Defaults to 'claim'",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "name is used to qualify the identities returned by this provider. - It MUST be unique and not shared by any other identity provider used - It MUST be a valid path segment: name cannot equal '.' or '..' or contain '/' or '%' or ':' Ref: https://godoc.org/github.com/openshift/origin/pkg/user/apis/user/validation#ValidateIdentityProviderName",
									MarkdownDescription: "name is used to qualify the identities returned by this provider. - It MUST be unique and not shared by any other identity provider used - It MUST be a valid path segment: name cannot equal '.' or '..' or contain '/' or '%' or ':' Ref: https://godoc.org/github.com/openshift/origin/pkg/user/apis/user/validation#ValidateIdentityProviderName",
									Required:            false,
									Optional:            false,
									Computed:            true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
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
													Optional:            false,
													Computed:            true,
												},

												"groups": schema.ListAttribute{
													Description:         "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user. If multiple claims are specified, the first one with a non-empty value is used.",
													MarkdownDescription: "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user. If multiple claims are specified, the first one with a non-empty value is used.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.ListAttribute{
													Description:         "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",
													MarkdownDescription: "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"preferred_username": schema.ListAttribute{
													Description:         "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the sub claim",
													MarkdownDescription: "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the sub claim",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"client_id": schema.StringAttribute{
											Description:         "clientID is the oauth client ID",
											MarkdownDescription: "clientID is the oauth client ID",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "name is the metadata.name of the referenced secret",
													MarkdownDescription: "name is the metadata.name of the referenced secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"extra_authorize_parameters": schema.MapAttribute{
											Description:         "extraAuthorizeParameters are any custom parameters to add to the authorize request.",
											MarkdownDescription: "extraAuthorizeParameters are any custom parameters to add to the authorize request.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"extra_scopes": schema.ListAttribute{
											Description:         "extraScopes are any scopes to request in addition to the standard 'openid' scope.",
											MarkdownDescription: "extraScopes are any scopes to request in addition to the standard 'openid' scope.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"issuer": schema.StringAttribute{
											Description:         "issuer is the URL that the OpenID Provider asserts as its Issuer Identifier. It must use the https scheme with no query or fragment component.",
											MarkdownDescription: "issuer is the URL that the OpenID Provider asserts as its Issuer Identifier. It must use the https scheme with no query or fragment component.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"challenge_url": schema.StringAttribute{
											Description:         "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here. ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when challenge is set to true.",
											MarkdownDescription: "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here. ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when challenge is set to true.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"client_common_names": schema.ListAttribute{
											Description:         "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",
											MarkdownDescription: "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"email_headers": schema.ListAttribute{
											Description:         "emailHeaders is the set of headers to check for the email address",
											MarkdownDescription: "emailHeaders is the set of headers to check for the email address",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"headers": schema.ListAttribute{
											Description:         "headers is the set of headers to check for identity information",
											MarkdownDescription: "headers is the set of headers to check for identity information",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"login_url": schema.StringAttribute{
											Description:         "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when login is set to true.",
											MarkdownDescription: "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string https://www.example.com/auth-proxy/oauth/authorize?${query} Required when login is set to true.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name_headers": schema.ListAttribute{
											Description:         "nameHeaders is the set of headers to check for the display name",
											MarkdownDescription: "nameHeaders is the set of headers to check for the display name",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"preferred_username_headers": schema.ListAttribute{
											Description:         "preferredUsernameHeaders is the set of headers to check for the preferred username",
											MarkdownDescription: "preferredUsernameHeaders is the set of headers to check for the preferred username",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"type": schema.StringAttribute{
									Description:         "type identifies the identity provider type for this entry.",
									MarkdownDescription: "type identifies the identity provider type for this entry.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hive_openshift_io_sync_identity_provider_v1")

	var data HiveOpenshiftIoSyncIdentityProviderV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "syncidentityproviders"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse HiveOpenshiftIoSyncIdentityProviderV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("hive.openshift.io/v1")
	data.Kind = pointer.String("SyncIdentityProvider")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
