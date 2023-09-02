/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ExternalSecretsIoSecretStoreV1Alpha1Manifest{}
)

func NewExternalSecretsIoSecretStoreV1Alpha1Manifest() datasource.DataSource {
	return &ExternalSecretsIoSecretStoreV1Alpha1Manifest{}
}

type ExternalSecretsIoSecretStoreV1Alpha1Manifest struct{}

type ExternalSecretsIoSecretStoreV1Alpha1ManifestData struct {
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
		Controller *string `tfsdk:"controller" json:"controller,omitempty"`
		Provider   *struct {
			Akeyless *struct {
				AkeylessGWApiURL *string `tfsdk:"akeyless_gw_api_url" json:"akeylessGWApiURL,omitempty"`
				AuthSecretRef    *struct {
					KubernetesAuth *struct {
						AccessID    *string `tfsdk:"access_id" json:"accessID,omitempty"`
						K8sConfName *string `tfsdk:"k8s_conf_name" json:"k8sConfName,omitempty"`
						SecretRef   *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
					} `tfsdk:"kubernetes_auth" json:"kubernetesAuth,omitempty"`
					SecretRef *struct {
						AccessID *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_id" json:"accessID,omitempty"`
						AccessType *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_type" json:"accessType,omitempty"`
						AccessTypeParam *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_type_param" json:"accessTypeParam,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth_secret_ref" json:"authSecretRef,omitempty"`
				CaBundle   *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
				CaProvider *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
			} `tfsdk:"akeyless" json:"akeyless,omitempty"`
			Alibaba *struct {
				Auth *struct {
					Rrsa *struct {
						OidcProviderArn   *string `tfsdk:"oidc_provider_arn" json:"oidcProviderArn,omitempty"`
						OidcTokenFilePath *string `tfsdk:"oidc_token_file_path" json:"oidcTokenFilePath,omitempty"`
						RoleArn           *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
						SessionName       *string `tfsdk:"session_name" json:"sessionName,omitempty"`
					} `tfsdk:"rrsa" json:"rrsa,omitempty"`
					SecretRef *struct {
						AccessKeyIDSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" json:"accessKeyIDSecretRef,omitempty"`
						AccessKeySecretSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_key_secret_secret_ref" json:"accessKeySecretSecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				RegionID *string `tfsdk:"region_id" json:"regionID,omitempty"`
			} `tfsdk:"alibaba" json:"alibaba,omitempty"`
			Aws *struct {
				Auth *struct {
					Jwt *struct {
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
					} `tfsdk:"jwt" json:"jwt,omitempty"`
					SecretRef *struct {
						AccessKeyIDSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" json:"accessKeyIDSecretRef,omitempty"`
						SecretAccessKeySecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" json:"secretAccessKeySecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Region  *string `tfsdk:"region" json:"region,omitempty"`
				Role    *string `tfsdk:"role" json:"role,omitempty"`
				Service *string `tfsdk:"service" json:"service,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azurekv *struct {
				AuthSecretRef *struct {
					ClientId *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"client_id" json:"clientId,omitempty"`
					ClientSecret *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				} `tfsdk:"auth_secret_ref" json:"authSecretRef,omitempty"`
				AuthType          *string `tfsdk:"auth_type" json:"authType,omitempty"`
				IdentityId        *string `tfsdk:"identity_id" json:"identityId,omitempty"`
				ServiceAccountRef *struct {
					Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
					Name      *string   `tfsdk:"name" json:"name,omitempty"`
					Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
				TenantId *string `tfsdk:"tenant_id" json:"tenantId,omitempty"`
				VaultUrl *string `tfsdk:"vault_url" json:"vaultUrl,omitempty"`
			} `tfsdk:"azurekv" json:"azurekv,omitempty"`
			Fake *struct {
				Data *[]struct {
					Key      *string            `tfsdk:"key" json:"key,omitempty"`
					Value    *string            `tfsdk:"value" json:"value,omitempty"`
					ValueMap *map[string]string `tfsdk:"value_map" json:"valueMap,omitempty"`
					Version  *string            `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"data" json:"data,omitempty"`
			} `tfsdk:"fake" json:"fake,omitempty"`
			Gcpsm *struct {
				Auth *struct {
					SecretRef *struct {
						SecretAccessKeySecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" json:"secretAccessKeySecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					WorkloadIdentity *struct {
						ClusterLocation   *string `tfsdk:"cluster_location" json:"clusterLocation,omitempty"`
						ClusterName       *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
						ClusterProjectID  *string `tfsdk:"cluster_project_id" json:"clusterProjectID,omitempty"`
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
					} `tfsdk:"workload_identity" json:"workloadIdentity,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				ProjectID *string `tfsdk:"project_id" json:"projectID,omitempty"`
			} `tfsdk:"gcpsm" json:"gcpsm,omitempty"`
			Gitlab *struct {
				Auth *struct {
					SecretRef *struct {
						AccessToken *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_token" json:"accessToken,omitempty"`
					} `tfsdk:"secret_ref" json:"SecretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				ProjectID *string `tfsdk:"project_id" json:"projectID,omitempty"`
				Url       *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"gitlab" json:"gitlab,omitempty"`
			Ibm *struct {
				Auth *struct {
					SecretRef *struct {
						SecretApiKeySecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_api_key_secret_ref" json:"secretApiKeySecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				ServiceUrl *string `tfsdk:"service_url" json:"serviceUrl,omitempty"`
			} `tfsdk:"ibm" json:"ibm,omitempty"`
			Kubernetes *struct {
				Auth *struct {
					Cert *struct {
						ClientCert *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"client_key" json:"clientKey,omitempty"`
					} `tfsdk:"cert" json:"cert,omitempty"`
					ServiceAccount *struct {
						ServiceAccount *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					Token *struct {
						BearerToken *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"bearer_token" json:"bearerToken,omitempty"`
					} `tfsdk:"token" json:"token,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				RemoteNamespace *string `tfsdk:"remote_namespace" json:"remoteNamespace,omitempty"`
				Server          *struct {
					CaBundle   *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
					CaProvider *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			Oracle *struct {
				Auth *struct {
					SecretRef *struct {
						Fingerprint *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"fingerprint" json:"fingerprint,omitempty"`
						Privatekey *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"privatekey" json:"privatekey,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Tenancy *string `tfsdk:"tenancy" json:"tenancy,omitempty"`
					User    *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Region *string `tfsdk:"region" json:"region,omitempty"`
				Vault  *string `tfsdk:"vault" json:"vault,omitempty"`
			} `tfsdk:"oracle" json:"oracle,omitempty"`
			Vault *struct {
				Auth *struct {
					AppRole *struct {
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						RoleId    *string `tfsdk:"role_id" json:"roleId,omitempty"`
						SecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"app_role" json:"appRole,omitempty"`
					Cert *struct {
						ClientCert *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"client_cert" json:"clientCert,omitempty"`
						SecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"cert" json:"cert,omitempty"`
					Jwt *struct {
						KubernetesServiceAccountToken *struct {
							Audiences         *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							ExpirationSeconds *int64    `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
							ServiceAccountRef *struct {
								Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
								Name      *string   `tfsdk:"name" json:"name,omitempty"`
								Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
						} `tfsdk:"kubernetes_service_account_token" json:"kubernetesServiceAccountToken,omitempty"`
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						Role      *string `tfsdk:"role" json:"role,omitempty"`
						SecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"jwt" json:"jwt,omitempty"`
					Kubernetes *struct {
						MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						Role      *string `tfsdk:"role" json:"role,omitempty"`
						SecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
					} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
					Ldap *struct {
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						SecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						Username *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"ldap" json:"ldap,omitempty"`
					TokenSecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"token_secret_ref" json:"tokenSecretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				CaBundle   *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
				CaProvider *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
				ForwardInconsistent *bool   `tfsdk:"forward_inconsistent" json:"forwardInconsistent,omitempty"`
				Namespace           *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Path                *string `tfsdk:"path" json:"path,omitempty"`
				ReadYourWrites      *bool   `tfsdk:"read_your_writes" json:"readYourWrites,omitempty"`
				Server              *string `tfsdk:"server" json:"server,omitempty"`
				Version             *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"vault" json:"vault,omitempty"`
			Webhook *struct {
				Body       *string `tfsdk:"body" json:"body,omitempty"`
				CaBundle   *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
				CaProvider *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
				Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
				Method  *string            `tfsdk:"method" json:"method,omitempty"`
				Result  *struct {
					JsonPath *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				} `tfsdk:"result" json:"result,omitempty"`
				Secrets *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"secrets" json:"secrets,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
				Url     *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"webhook" json:"webhook,omitempty"`
			Yandexlockbox *struct {
				ApiEndpoint *string `tfsdk:"api_endpoint" json:"apiEndpoint,omitempty"`
				Auth        *struct {
					AuthorizedKeySecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"authorized_key_secret_ref" json:"authorizedKeySecretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				CaProvider *struct {
					CertSecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
				} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
			} `tfsdk:"yandexlockbox" json:"yandexlockbox,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
		RetrySettings *struct {
			MaxRetries    *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
		} `tfsdk:"retry_settings" json:"retrySettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternalSecretsIoSecretStoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_secret_store_v1alpha1_manifest"
}

func (r *ExternalSecretsIoSecretStoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SecretStore represents a secure external location for storing secrets, which can be referenced as part of 'storeRef' fields.",
		MarkdownDescription: "SecretStore represents a secure external location for storing secrets, which can be referenced as part of 'storeRef' fields.",
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
				Description:         "SecretStoreSpec defines the desired state of SecretStore.",
				MarkdownDescription: "SecretStoreSpec defines the desired state of SecretStore.",
				Attributes: map[string]schema.Attribute{
					"controller": schema.StringAttribute{
						Description:         "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
						MarkdownDescription: "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Used to configure the provider. Only one provider may be set",
						MarkdownDescription: "Used to configure the provider. Only one provider may be set",
						Attributes: map[string]schema.Attribute{
							"akeyless": schema.SingleNestedAttribute{
								Description:         "Akeyless configures this store to sync secrets using Akeyless Vault provider",
								MarkdownDescription: "Akeyless configures this store to sync secrets using Akeyless Vault provider",
								Attributes: map[string]schema.Attribute{
									"akeyless_gw_api_url": schema.StringAttribute{
										Description:         "Akeyless GW API Url from which the secrets to be fetched from.",
										MarkdownDescription: "Akeyless GW API Url from which the secrets to be fetched from.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"auth_secret_ref": schema.SingleNestedAttribute{
										Description:         "Auth configures how the operator authenticates with Akeyless.",
										MarkdownDescription: "Auth configures how the operator authenticates with Akeyless.",
										Attributes: map[string]schema.Attribute{
											"kubernetes_auth": schema.SingleNestedAttribute{
												Description:         "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
												MarkdownDescription: "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
												Attributes: map[string]schema.Attribute{
													"access_id": schema.StringAttribute{
														Description:         "the Akeyless Kubernetes auth-method access-id",
														MarkdownDescription: "the Akeyless Kubernetes auth-method access-id",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"k8s_conf_name": schema.StringAttribute{
														Description:         "Kubernetes-auth configuration name in Akeyless-Gateway",
														MarkdownDescription: "Kubernetes-auth configuration name in Akeyless-Gateway",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_account_ref": schema.SingleNestedAttribute{
														Description:         "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
														MarkdownDescription: "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a Secret that contains the details to authenticate with Akeyless.",
												MarkdownDescription: "Reference to a Secret that contains the details to authenticate with Akeyless.",
												Attributes: map[string]schema.Attribute{
													"access_id": schema.SingleNestedAttribute{
														Description:         "The SecretAccessID is used for authentication",
														MarkdownDescription: "The SecretAccessID is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"access_type": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"access_type_param": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"ca_bundle": schema.StringAttribute{
										Description:         "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.Base64Validator(),
										},
									},

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
												MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace the Provider type is in.",
												MarkdownDescription: "The namespace the Provider type is in.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Secret", "ConfigMap"),
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
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"alibaba": schema.SingleNestedAttribute{
								Description:         "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
								MarkdownDescription: "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "AlibabaAuth contains a secretRef for credentials.",
										MarkdownDescription: "AlibabaAuth contains a secretRef for credentials.",
										Attributes: map[string]schema.Attribute{
											"rrsa": schema.SingleNestedAttribute{
												Description:         "Authenticate against Alibaba using RRSA.",
												MarkdownDescription: "Authenticate against Alibaba using RRSA.",
												Attributes: map[string]schema.Attribute{
													"oidc_provider_arn": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"oidc_token_file_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"role_arn": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"session_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
												MarkdownDescription: "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
												Attributes: map[string]schema.Attribute{
													"access_key_id_secret_ref": schema.SingleNestedAttribute{
														Description:         "The AccessKeyID is used for authentication",
														MarkdownDescription: "The AccessKeyID is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"access_key_secret_secret_ref": schema.SingleNestedAttribute{
														Description:         "The AccessKeySecret is used for authentication",
														MarkdownDescription: "The AccessKeySecret is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"region_id": schema.StringAttribute{
										Description:         "Alibaba Region to be used for the provider",
										MarkdownDescription: "Alibaba Region to be used for the provider",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"aws": schema.SingleNestedAttribute{
								Description:         "AWS configures this store to sync secrets using AWS Secret Manager provider",
								MarkdownDescription: "AWS configures this store to sync secrets using AWS Secret Manager provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
										MarkdownDescription: "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
										Attributes: map[string]schema.Attribute{
											"jwt": schema.SingleNestedAttribute{
												Description:         "Authenticate against AWS using service account tokens.",
												MarkdownDescription: "Authenticate against AWS using service account tokens.",
												Attributes: map[string]schema.Attribute{
													"service_account_ref": schema.SingleNestedAttribute{
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
												MarkdownDescription: "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
												Attributes: map[string]schema.Attribute{
													"access_key_id_secret_ref": schema.SingleNestedAttribute{
														Description:         "The AccessKeyID is used for authentication",
														MarkdownDescription: "The AccessKeyID is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_access_key_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": schema.StringAttribute{
										Description:         "AWS Region to be used for the provider",
										MarkdownDescription: "AWS Region to be used for the provider",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"role": schema.StringAttribute{
										Description:         "Role is a Role ARN which the SecretManager provider will assume",
										MarkdownDescription: "Role is a Role ARN which the SecretManager provider will assume",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service": schema.StringAttribute{
										Description:         "Service defines which service should be used to fetch the secrets",
										MarkdownDescription: "Service defines which service should be used to fetch the secrets",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("SecretsManager", "ParameterStore"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"azurekv": schema.SingleNestedAttribute{
								Description:         "AzureKV configures this store to sync secrets using Azure Key Vault provider",
								MarkdownDescription: "AzureKV configures this store to sync secrets using Azure Key Vault provider",
								Attributes: map[string]schema.Attribute{
									"auth_secret_ref": schema.SingleNestedAttribute{
										Description:         "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
										MarkdownDescription: "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
										Attributes: map[string]schema.Attribute{
											"client_id": schema.SingleNestedAttribute{
												Description:         "The Azure clientId of the service principle used for authentication.",
												MarkdownDescription: "The Azure clientId of the service principle used for authentication.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_secret": schema.SingleNestedAttribute{
												Description:         "The Azure ClientSecret of the service principle used for authentication.",
												MarkdownDescription: "The Azure ClientSecret of the service principle used for authentication.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

									"auth_type": schema.StringAttribute{
										Description:         "Auth type defines how to authenticate to the keyvault service. Valid values are: - 'ServicePrincipal' (default): Using a service principal (tenantId, clientId, clientSecret) - 'ManagedIdentity': Using Managed Identity assigned to the pod (see aad-pod-identity)",
										MarkdownDescription: "Auth type defines how to authenticate to the keyvault service. Valid values are: - 'ServicePrincipal' (default): Using a service principal (tenantId, clientId, clientSecret) - 'ManagedIdentity': Using Managed Identity assigned to the pod (see aad-pod-identity)",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ServicePrincipal", "ManagedIdentity", "WorkloadIdentity"),
										},
									},

									"identity_id": schema.StringAttribute{
										Description:         "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
										MarkdownDescription: "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account_ref": schema.SingleNestedAttribute{
										Description:         "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
										MarkdownDescription: "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
										Attributes: map[string]schema.Attribute{
											"audiences": schema.ListAttribute{
												Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the ServiceAccount resource being referred to.",
												MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
										MarkdownDescription: "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vault_url": schema.StringAttribute{
										Description:         "Vault Url from which the secrets to be fetched from.",
										MarkdownDescription: "Vault Url from which the secrets to be fetched from.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"fake": schema.SingleNestedAttribute{
								Description:         "Fake configures a store with static key/value pairs",
								MarkdownDescription: "Fake configures a store with static key/value pairs",
								Attributes: map[string]schema.Attribute{
									"data": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_map": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
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
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"gcpsm": schema.SingleNestedAttribute{
								Description:         "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
								MarkdownDescription: "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth defines the information necessary to authenticate against GCP",
										MarkdownDescription: "Auth defines the information necessary to authenticate against GCP",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_access_key_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"workload_identity": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cluster_location": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"cluster_project_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account_ref": schema.SingleNestedAttribute{
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"project_id": schema.StringAttribute{
										Description:         "ProjectID project where secret is located",
										MarkdownDescription: "ProjectID project where secret is located",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"gitlab": schema.SingleNestedAttribute{
								Description:         "GitLab configures this store to sync secrets using GitLab Variables provider",
								MarkdownDescription: "GitLab configures this store to sync secrets using GitLab Variables provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how secret-manager authenticates with a GitLab instance.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with a GitLab instance.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"access_token": schema.SingleNestedAttribute{
														Description:         "AccessToken is used for authentication.",
														MarkdownDescription: "AccessToken is used for authentication.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"project_id": schema.StringAttribute{
										Description:         "ProjectID specifies a project where secrets are located.",
										MarkdownDescription: "ProjectID specifies a project where secrets are located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
										MarkdownDescription: "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"ibm": schema.SingleNestedAttribute{
								Description:         "IBM configures this store to sync secrets using IBM Cloud provider",
								MarkdownDescription: "IBM configures this store to sync secrets using IBM Cloud provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how secret-manager authenticates with the IBM secrets manager.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the IBM secrets manager.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_api_key_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"service_url": schema.StringAttribute{
										Description:         "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
										MarkdownDescription: "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"kubernetes": schema.SingleNestedAttribute{
								Description:         "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
								MarkdownDescription: "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how secret-manager authenticates with a Kubernetes instance.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with a Kubernetes instance.",
										Attributes: map[string]schema.Attribute{
											"cert": schema.SingleNestedAttribute{
												Description:         "has both clientCert and clientKey as secretKeySelector",
												MarkdownDescription: "has both clientCert and clientKey as secretKeySelector",
												Attributes: map[string]schema.Attribute{
													"client_cert": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_key": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
												Validators: []UNKNOWN{
													schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("service_account"), path.MatchRelative().AtParent().AtName("token")),
												},
											},

											"service_account": schema.SingleNestedAttribute{
												Description:         "points to a service account that should be used for authentication",
												MarkdownDescription: "points to a service account that should be used for authentication",
												Attributes: map[string]schema.Attribute{
													"service_account": schema.SingleNestedAttribute{
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
												Validators: []UNKNOWN{
													schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("cert"), path.MatchRelative().AtParent().AtName("token")),
												},
											},

											"token": schema.SingleNestedAttribute{
												Description:         "use static token to authenticate with",
												MarkdownDescription: "use static token to authenticate with",
												Attributes: map[string]schema.Attribute{
													"bearer_token": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
												Validators: []UNKNOWN{
													schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("cert"), path.MatchRelative().AtParent().AtName("service_account")),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"remote_namespace": schema.StringAttribute{
										Description:         "Remote namespace to fetch the secrets from",
										MarkdownDescription: "Remote namespace to fetch the secrets from",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.SingleNestedAttribute{
										Description:         "configures the Kubernetes server Address.",
										MarkdownDescription: "configures the Kubernetes server Address.",
										Attributes: map[string]schema.Attribute{
											"ca_bundle": schema.StringAttribute{
												Description:         "CABundle is a base64-encoded CA certificate",
												MarkdownDescription: "CABundle is a base64-encoded CA certificate",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"ca_provider": schema.SingleNestedAttribute{
												Description:         "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
												MarkdownDescription: "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
														MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the object located at the provider type.",
														MarkdownDescription: "The name of the object located at the provider type.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace the Provider type is in.",
														MarkdownDescription: "The namespace the Provider type is in.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
														MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Secret", "ConfigMap"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": schema.StringAttribute{
												Description:         "configures the Kubernetes server Address.",
												MarkdownDescription: "configures the Kubernetes server Address.",
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
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"oracle": schema.SingleNestedAttribute{
								Description:         "Oracle configures this store to sync secrets using Oracle Vault provider",
								MarkdownDescription: "Oracle configures this store to sync secrets using Oracle Vault provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, use the instance principal, otherwise the user credentials specified in Auth.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, use the instance principal, otherwise the user credentials specified in Auth.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef to pass through sensitive information.",
												MarkdownDescription: "SecretRef to pass through sensitive information.",
												Attributes: map[string]schema.Attribute{
													"fingerprint": schema.SingleNestedAttribute{
														Description:         "Fingerprint is the fingerprint of the API private key.",
														MarkdownDescription: "Fingerprint is the fingerprint of the API private key.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"privatekey": schema.SingleNestedAttribute{
														Description:         "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
														MarkdownDescription: "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
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

											"tenancy": schema.StringAttribute{
												Description:         "Tenancy is the tenancy OCID where user is located.",
												MarkdownDescription: "Tenancy is the tenancy OCID where user is located.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is an access OCID specific to the account.",
												MarkdownDescription: "User is an access OCID specific to the account.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the region where vault is located.",
										MarkdownDescription: "Region is the region where vault is located.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"vault": schema.StringAttribute{
										Description:         "Vault is the vault's OCID of the specific vault where secret is located.",
										MarkdownDescription: "Vault is the vault's OCID of the specific vault where secret is located.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"vault": schema.SingleNestedAttribute{
								Description:         "Vault configures this store to sync secrets using Hashi provider",
								MarkdownDescription: "Vault configures this store to sync secrets using Hashi provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how secret-manager authenticates with the Vault server.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the Vault server.",
										Attributes: map[string]schema.Attribute{
											"app_role": schema.SingleNestedAttribute{
												Description:         "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
												MarkdownDescription: "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",
														MarkdownDescription: "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"role_id": schema.StringAttribute{
														Description:         "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
														MarkdownDescription: "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
														MarkdownDescription: "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
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

											"cert": schema.SingleNestedAttribute{
												Description:         "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
												MarkdownDescription: "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
												Attributes: map[string]schema.Attribute{
													"client_cert": schema.SingleNestedAttribute{
														Description:         "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
														MarkdownDescription: "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"jwt": schema.SingleNestedAttribute{
												Description:         "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
												MarkdownDescription: "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
												Attributes: map[string]schema.Attribute{
													"kubernetes_service_account_token": schema.SingleNestedAttribute{
														Description:         "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the 'TokenRequest' API.",
														MarkdownDescription: "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the 'TokenRequest' API.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified.",
																MarkdownDescription: "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"expiration_seconds": schema.Int64Attribute{
																Description:         "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to 10 minutes.",
																MarkdownDescription: "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to 10 minutes.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_ref": schema.SingleNestedAttribute{
																Description:         "Service account field containing the name of a kubernetes ServiceAccount.",
																MarkdownDescription: "Service account field containing the name of a kubernetes ServiceAccount.",
																Attributes: map[string]schema.Attribute{
																	"audiences": schema.ListAttribute{
																		Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the ServiceAccount resource being referred to.",
																		MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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

													"path": schema.StringAttribute{
														Description:         "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
														MarkdownDescription: "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"role": schema.StringAttribute{
														Description:         "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
														MarkdownDescription: "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
														MarkdownDescription: "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"kubernetes": schema.SingleNestedAttribute{
												Description:         "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
												MarkdownDescription: "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
														MarkdownDescription: "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"role": schema.StringAttribute{
														Description:         "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
														MarkdownDescription: "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_account_ref": schema.SingleNestedAttribute{
														Description:         "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
														MarkdownDescription: "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"ldap": schema.SingleNestedAttribute{
												Description:         "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
												MarkdownDescription: "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
														MarkdownDescription: "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"username": schema.StringAttribute{
														Description:         "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
														MarkdownDescription: "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_secret_ref": schema.SingleNestedAttribute{
												Description:         "TokenSecretRef authenticates with Vault by presenting a token.",
												MarkdownDescription: "TokenSecretRef authenticates with Vault by presenting a token.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

									"ca_bundle": schema.StringAttribute{
										Description:         "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.Base64Validator(),
										},
									},

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate Vault server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Vault server certificate.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
												MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace the Provider type is in.",
												MarkdownDescription: "The namespace the Provider type is in.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Secret", "ConfigMap"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"forward_inconsistent": schema.BoolAttribute{
										Description:         "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
										MarkdownDescription: "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
										MarkdownDescription: "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
										MarkdownDescription: "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"read_your_writes": schema.BoolAttribute{
										Description:         "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
										MarkdownDescription: "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.StringAttribute{
										Description:         "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
										MarkdownDescription: "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"version": schema.StringAttribute{
										Description:         "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
										MarkdownDescription: "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("v1", "v2"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"webhook": schema.SingleNestedAttribute{
								Description:         "Webhook configures this store to sync secrets using a generic templated webhook",
								MarkdownDescription: "Webhook configures this store to sync secrets using a generic templated webhook",
								Attributes: map[string]schema.Attribute{
									"body": schema.StringAttribute{
										Description:         "Body",
										MarkdownDescription: "Body",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_bundle": schema.StringAttribute{
										Description:         "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.Base64Validator(),
										},
									},

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate webhook server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate webhook server certificate.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
												MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace the Provider type is in.",
												MarkdownDescription: "The namespace the Provider type is in.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Secret", "ConfigMap"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers": schema.MapAttribute{
										Description:         "Headers",
										MarkdownDescription: "Headers",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"method": schema.StringAttribute{
										Description:         "Webhook Method",
										MarkdownDescription: "Webhook Method",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"result": schema.SingleNestedAttribute{
										Description:         "Result formatting",
										MarkdownDescription: "Result formatting",
										Attributes: map[string]schema.Attribute{
											"json_path": schema.StringAttribute{
												Description:         "Json path of return value",
												MarkdownDescription: "Json path of return value",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"secrets": schema.ListNestedAttribute{
										Description:         "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
										MarkdownDescription: "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of this secret in templates",
													MarkdownDescription: "Name of this secret in templates",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "Secret ref to fill in credentials",
													MarkdownDescription: "Secret ref to fill in credentials",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
															MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The name of the Secret resource being referred to.",
															MarkdownDescription: "The name of the Secret resource being referred to.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
															MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout": schema.StringAttribute{
										Description:         "Timeout",
										MarkdownDescription: "Timeout",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "Webhook url to call",
										MarkdownDescription: "Webhook url to call",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"yandexlockbox": schema.SingleNestedAttribute{
								Description:         "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
								MarkdownDescription: "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
								Attributes: map[string]schema.Attribute{
									"api_endpoint": schema.StringAttribute{
										Description:         "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										MarkdownDescription: "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"auth": schema.SingleNestedAttribute{
										Description:         "Auth defines the information necessary to authenticate against Yandex Lockbox",
										MarkdownDescription: "Auth defines the information necessary to authenticate against Yandex Lockbox",
										Attributes: map[string]schema.Attribute{
											"authorized_key_secret_ref": schema.SingleNestedAttribute{
												Description:         "The authorized key used for authentication",
												MarkdownDescription: "The authorized key used for authentication",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
										Attributes: map[string]schema.Attribute{
											"cert_secret_ref": schema.SingleNestedAttribute{
												Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
												MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []UNKNOWN{
									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook")),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"retry_settings": schema.SingleNestedAttribute{
						Description:         "Used to configure http retries if failed",
						MarkdownDescription: "Used to configure http retries if failed",
						Attributes: map[string]schema.Attribute{
							"max_retries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retry_interval": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ExternalSecretsIoSecretStoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_secret_store_v1alpha1_manifest")

	var model ExternalSecretsIoSecretStoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("external-secrets.io/v1alpha1")
	model.Kind = pointer.String("SecretStore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
