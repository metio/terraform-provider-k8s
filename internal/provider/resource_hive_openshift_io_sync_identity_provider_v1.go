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

type HiveOpenshiftIoSyncIdentityProviderV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoSyncIdentityProviderV1Resource)(nil)
)

type HiveOpenshiftIoSyncIdentityProviderV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoSyncIdentityProviderV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ClusterDeploymentRefs *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"cluster_deployment_refs" yaml:"clusterDeploymentRefs,omitempty"`

		IdentityProviders *[]struct {
			BasicAuth *struct {
				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				TlsClientCert *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"tls_client_cert" yaml:"tlsClientCert,omitempty"`

				TlsClientKey *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"tls_client_key" yaml:"tlsClientKey,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

			Github *struct {
				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

				ClientSecret *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

				Organizations *[]string `tfsdk:"organizations" yaml:"organizations,omitempty"`

				Teams *[]string `tfsdk:"teams" yaml:"teams,omitempty"`
			} `tfsdk:"github" yaml:"github,omitempty"`

			Gitlab *struct {
				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

				ClientSecret *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

			Google *struct {
				ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

				ClientSecret *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				HostedDomain *string `tfsdk:"hosted_domain" yaml:"hostedDomain,omitempty"`
			} `tfsdk:"google" yaml:"google,omitempty"`

			Htpasswd *struct {
				FileData *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"file_data" yaml:"fileData,omitempty"`
			} `tfsdk:"htpasswd" yaml:"htpasswd,omitempty"`

			Keystone *struct {
				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				DomainName *string `tfsdk:"domain_name" yaml:"domainName,omitempty"`

				TlsClientCert *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"tls_client_cert" yaml:"tlsClientCert,omitempty"`

				TlsClientKey *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"tls_client_key" yaml:"tlsClientKey,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"keystone" yaml:"keystone,omitempty"`

			Ldap *struct {
				Attributes *struct {
					Email *[]string `tfsdk:"email" yaml:"email,omitempty"`

					Id *[]string `tfsdk:"id" yaml:"id,omitempty"`

					Name *[]string `tfsdk:"name" yaml:"name,omitempty"`

					PreferredUsername *[]string `tfsdk:"preferred_username" yaml:"preferredUsername,omitempty"`
				} `tfsdk:"attributes" yaml:"attributes,omitempty"`

				BindDN *string `tfsdk:"bind_dn" yaml:"bindDN,omitempty"`

				BindPassword *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"bind_password" yaml:"bindPassword,omitempty"`

				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"ldap" yaml:"ldap,omitempty"`

			MappingMethod *string `tfsdk:"mapping_method" yaml:"mappingMethod,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			OpenID *struct {
				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				Claims *struct {
					Email *[]string `tfsdk:"email" yaml:"email,omitempty"`

					Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

					Name *[]string `tfsdk:"name" yaml:"name,omitempty"`

					PreferredUsername *[]string `tfsdk:"preferred_username" yaml:"preferredUsername,omitempty"`
				} `tfsdk:"claims" yaml:"claims,omitempty"`

				ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

				ClientSecret *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				ExtraAuthorizeParameters *map[string]string `tfsdk:"extra_authorize_parameters" yaml:"extraAuthorizeParameters,omitempty"`

				ExtraScopes *[]string `tfsdk:"extra_scopes" yaml:"extraScopes,omitempty"`

				Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`
			} `tfsdk:"open_id" yaml:"openID,omitempty"`

			RequestHeader *struct {
				Ca *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				ChallengeURL *string `tfsdk:"challenge_url" yaml:"challengeURL,omitempty"`

				ClientCommonNames *[]string `tfsdk:"client_common_names" yaml:"clientCommonNames,omitempty"`

				EmailHeaders *[]string `tfsdk:"email_headers" yaml:"emailHeaders,omitempty"`

				Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

				LoginURL *string `tfsdk:"login_url" yaml:"loginURL,omitempty"`

				NameHeaders *[]string `tfsdk:"name_headers" yaml:"nameHeaders,omitempty"`

				PreferredUsernameHeaders *[]string `tfsdk:"preferred_username_headers" yaml:"preferredUsernameHeaders,omitempty"`
			} `tfsdk:"request_header" yaml:"requestHeader,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"identity_providers" yaml:"identityProviders,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoSyncIdentityProviderV1Resource() resource.Resource {
	return &HiveOpenshiftIoSyncIdentityProviderV1Resource{}
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_sync_identity_provider_v1"
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "SyncIdentityProvider is the Schema for the SyncIdentityProvider API",
		MarkdownDescription: "SyncIdentityProvider is the Schema for the SyncIdentityProvider API",
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
				Description:         "SyncIdentityProviderSpec defines the SyncIdentityProviderCommonSpec identity providers to sync along with ClusterDeploymentRefs indicating which clusters the SyncIdentityProvider applies to in the SyncIdentityProvider's namespace.",
				MarkdownDescription: "SyncIdentityProviderSpec defines the SyncIdentityProviderCommonSpec identity providers to sync along with ClusterDeploymentRefs indicating which clusters the SyncIdentityProvider applies to in the SyncIdentityProvider's namespace.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cluster_deployment_refs": {
						Description:         "ClusterDeploymentRefs is the list of LocalObjectReference indicating which clusters the SyncSet applies to in the SyncSet's namespace.",
						MarkdownDescription: "ClusterDeploymentRefs is the list of LocalObjectReference indicating which clusters the SyncSet applies to in the SyncSet's namespace.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"identity_providers": {
						Description:         "IdentityProviders is an ordered list of ways for a user to identify themselves",
						MarkdownDescription: "IdentityProviders is an ordered list of ways for a user to identify themselves",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"basic_auth": {
								Description:         "basicAuth contains configuration options for the BasicAuth IdP",
								MarkdownDescription: "basicAuth contains configuration options for the BasicAuth IdP",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"tls_client_cert": {
										Description:         "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"tls_client_key": {
										Description:         "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"url": {
										Description:         "url is the remote URL to connect to",
										MarkdownDescription: "url is the remote URL to connect to",

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

							"github": {
								Description:         "github enables user authentication using GitHub credentials",
								MarkdownDescription: "github enables user authentication using GitHub credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. This can only be configured when hostname is set to a non-empty value. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. This can only be configured when hostname is set to a non-empty value. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"client_id": {
										Description:         "clientID is the oauth client ID",
										MarkdownDescription: "clientID is the oauth client ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret": {
										Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"hostname": {
										Description:         "hostname is the optional domain (e.g. 'mycompany.com') for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value configured at /setup/settings#hostname.",
										MarkdownDescription: "hostname is the optional domain (e.g. 'mycompany.com') for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value configured at /setup/settings#hostname.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"organizations": {
										Description:         "organizations optionally restricts which organizations are allowed to log in",
										MarkdownDescription: "organizations optionally restricts which organizations are allowed to log in",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"teams": {
										Description:         "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",
										MarkdownDescription: "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",

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

							"gitlab": {
								Description:         "gitlab enables user authentication using GitLab credentials",
								MarkdownDescription: "gitlab enables user authentication using GitLab credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"client_id": {
										Description:         "clientID is the oauth client ID",
										MarkdownDescription: "clientID is the oauth client ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret": {
										Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"url": {
										Description:         "url is the oauth server base URL",
										MarkdownDescription: "url is the oauth server base URL",

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

							"google": {
								Description:         "google enables user authentication using Google credentials",
								MarkdownDescription: "google enables user authentication using Google credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_id": {
										Description:         "clientID is the oauth client ID",
										MarkdownDescription: "clientID is the oauth client ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret": {
										Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"hosted_domain": {
										Description:         "hostedDomain is the optional Google App domain (e.g. 'mycompany.com') to restrict logins to",
										MarkdownDescription: "hostedDomain is the optional Google App domain (e.g. 'mycompany.com') to restrict logins to",

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

							"htpasswd": {
								Description:         "htpasswd enables user authentication using an HTPasswd file to validate credentials",
								MarkdownDescription: "htpasswd enables user authentication using an HTPasswd file to validate credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"file_data": {
										Description:         "fileData is a required reference to a secret by name containing the data to use as the htpasswd file. The key 'htpasswd' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. If the specified htpasswd data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "fileData is a required reference to a secret by name containing the data to use as the htpasswd file. The key 'htpasswd' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. If the specified htpasswd data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keystone": {
								Description:         "keystone enables user authentication using keystone password credentials",
								MarkdownDescription: "keystone enables user authentication using keystone password credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"domain_name": {
										Description:         "domainName is required for keystone v3",
										MarkdownDescription: "domainName is required for keystone v3",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_client_cert": {
										Description:         "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "tlsClientCert is an optional reference to a secret by name that contains the PEM-encoded TLS client certificate to present when connecting to the server. The key 'tls.crt' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"tls_client_key": {
										Description:         "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "tlsClientKey is an optional reference to a secret by name that contains the PEM-encoded TLS private key for the client certificate referenced in tlsClientCert. The key 'tls.key' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. If the specified certificate data is not valid, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"url": {
										Description:         "url is the remote URL to connect to",
										MarkdownDescription: "url is the remote URL to connect to",

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

							"ldap": {
								Description:         "ldap enables user authentication using LDAP credentials",
								MarkdownDescription: "ldap enables user authentication using LDAP credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"attributes": {
										Description:         "attributes maps LDAP attributes to identities",
										MarkdownDescription: "attributes maps LDAP attributes to identities",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"email": {
												Description:         "email is the list of attributes whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
												MarkdownDescription: "email is the list of attributes whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "id is the list of attributes whose values should be used as the user ID. Required. First non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. LDAP standard identity attribute is 'dn'",
												MarkdownDescription: "id is the list of attributes whose values should be used as the user ID. Required. First non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. LDAP standard identity attribute is 'dn'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is 'cn'",
												MarkdownDescription: "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is 'cn'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preferred_username": {
												Description:         "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is 'uid'",
												MarkdownDescription: "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is 'uid'",

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

									"bind_dn": {
										Description:         "bindDN is an optional DN to bind with during the search phase.",
										MarkdownDescription: "bindDN is an optional DN to bind with during the search phase.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"bind_password": {
										Description:         "bindPassword is an optional reference to a secret by name containing a password to bind with during the search phase. The key 'bindPassword' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "bindPassword is an optional reference to a secret by name containing a password to bind with during the search phase. The key 'bindPassword' is used to locate the data. If specified and the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"ca": {
										Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"insecure": {
										Description:         "insecure, if true, indicates the connection should not use TLS WARNING: Should not be set to 'true' with the URL scheme 'ldaps://' as 'ldaps://' URLs always          attempt to connect using TLS, even when 'insecure' is set to 'true' When 'true', 'ldap://' URLS connect insecurely. When 'false', 'ldap://' URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830.",
										MarkdownDescription: "insecure, if true, indicates the connection should not use TLS WARNING: Should not be set to 'true' with the URL scheme 'ldaps://' as 'ldaps://' URLs always          attempt to connect using TLS, even when 'insecure' is set to 'true' When 'true', 'ldap://' URLS connect insecurely. When 'false', 'ldap://' URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is: ldap://host:port/basedn?attribute?scope?filter",
										MarkdownDescription: "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is: ldap://host:port/basedn?attribute?scope?filter",

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

							"mapping_method": {
								Description:         "mappingMethod determines how identities from this provider are mapped to users Defaults to 'claim'",
								MarkdownDescription: "mappingMethod determines how identities from this provider are mapped to users Defaults to 'claim'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "name is used to qualify the identities returned by this provider. - It MUST be unique and not shared by any other identity provider used - It MUST be a valid path segment: name cannot equal '.' or '..' or contain '/' or '%' or ':'   Ref: https://godoc.org/github.com/openshift/origin/pkg/user/apis/user/validation#ValidateIdentityProviderName",
								MarkdownDescription: "name is used to qualify the identities returned by this provider. - It MUST be unique and not shared by any other identity provider used - It MUST be a valid path segment: name cannot equal '.' or '..' or contain '/' or '%' or ':'   Ref: https://godoc.org/github.com/openshift/origin/pkg/user/apis/user/validation#ValidateIdentityProviderName",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"open_id": {
								Description:         "openID enables user authentication using OpenID credentials",
								MarkdownDescription: "openID enables user authentication using OpenID credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is an optional reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. The key 'ca.crt' is used to locate the data. If specified and the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. If empty, the default system roots are used. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"claims": {
										Description:         "claims mappings",
										MarkdownDescription: "claims mappings",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"email": {
												Description:         "email is the list of claims whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
												MarkdownDescription: "email is the list of claims whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"groups": {
												Description:         "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user. If multiple claims are specified, the first one with a non-empty value is used.",
												MarkdownDescription: "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user. If multiple claims are specified, the first one with a non-empty value is used.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",
												MarkdownDescription: "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preferred_username": {
												Description:         "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the sub claim",
												MarkdownDescription: "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the sub claim",

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

									"client_id": {
										Description:         "clientID is the oauth client ID",
										MarkdownDescription: "clientID is the oauth client ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret": {
										Description:         "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",
										MarkdownDescription: "clientSecret is a required reference to the secret by name containing the oauth client secret. The key 'clientSecret' is used to locate the data. If the secret or expected key is not found, the identity provider is not honored. The namespace for this secret is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced secret",
												MarkdownDescription: "name is the metadata.name of the referenced secret",

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

									"extra_authorize_parameters": {
										Description:         "extraAuthorizeParameters are any custom parameters to add to the authorize request.",
										MarkdownDescription: "extraAuthorizeParameters are any custom parameters to add to the authorize request.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"extra_scopes": {
										Description:         "extraScopes are any scopes to request in addition to the standard 'openid' scope.",
										MarkdownDescription: "extraScopes are any scopes to request in addition to the standard 'openid' scope.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"issuer": {
										Description:         "issuer is the URL that the OpenID Provider asserts as its Issuer Identifier. It must use the https scheme with no query or fragment component.",
										MarkdownDescription: "issuer is the URL that the OpenID Provider asserts as its Issuer Identifier. It must use the https scheme with no query or fragment component.",

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

							"request_header": {
								Description:         "requestHeader enables user authentication using request header credentials",
								MarkdownDescription: "requestHeader enables user authentication using request header credentials",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "ca is a required reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. Specifically, it allows verification of incoming requests to prevent header spoofing. The key 'ca.crt' is used to locate the data. If the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. The namespace for this config map is openshift-config.",
										MarkdownDescription: "ca is a required reference to a config map by name containing the PEM-encoded CA bundle. It is used as a trust anchor to validate the TLS certificate presented by the remote server. Specifically, it allows verification of incoming requests to prevent header spoofing. The key 'ca.crt' is used to locate the data. If the config map or expected key is not found, the identity provider is not honored. If the specified ca data is not valid, the identity provider is not honored. The namespace for this config map is openshift-config.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",

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

									"challenge_url": {
										Description:         "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here. ${url} is replaced with the current URL, escaped to be safe in a query parameter   https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string   https://www.example.com/auth-proxy/oauth/authorize?${query} Required when challenge is set to true.",
										MarkdownDescription: "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here. ${url} is replaced with the current URL, escaped to be safe in a query parameter   https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string   https://www.example.com/auth-proxy/oauth/authorize?${query} Required when challenge is set to true.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_common_names": {
										Description:         "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",
										MarkdownDescription: "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"email_headers": {
										Description:         "emailHeaders is the set of headers to check for the email address",
										MarkdownDescription: "emailHeaders is the set of headers to check for the email address",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers": {
										Description:         "headers is the set of headers to check for identity information",
										MarkdownDescription: "headers is the set of headers to check for identity information",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"login_url": {
										Description:         "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter   https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string   https://www.example.com/auth-proxy/oauth/authorize?${query} Required when login is set to true.",
										MarkdownDescription: "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter   https://www.example.com/sso-login?then=${url} ${query} is replaced with the current query string   https://www.example.com/auth-proxy/oauth/authorize?${query} Required when login is set to true.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name_headers": {
										Description:         "nameHeaders is the set of headers to check for the display name",
										MarkdownDescription: "nameHeaders is the set of headers to check for the display name",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"preferred_username_headers": {
										Description:         "preferredUsernameHeaders is the set of headers to check for the preferred username",
										MarkdownDescription: "preferredUsernameHeaders is the set of headers to check for the preferred username",

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

							"type": {
								Description:         "type identifies the identity provider type for this entry.",
								MarkdownDescription: "type identifies the identity provider type for this entry.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_sync_identity_provider_v1")

	var state HiveOpenshiftIoSyncIdentityProviderV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoSyncIdentityProviderV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("SyncIdentityProvider")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_sync_identity_provider_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_sync_identity_provider_v1")

	var state HiveOpenshiftIoSyncIdentityProviderV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoSyncIdentityProviderV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("SyncIdentityProvider")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *HiveOpenshiftIoSyncIdentityProviderV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_sync_identity_provider_v1")
	// NO-OP: Terraform removes the state automatically for us
}
