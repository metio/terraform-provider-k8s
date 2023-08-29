/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"

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

type ExternalSecretsIoClusterSecretStoreV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ExternalSecretsIoClusterSecretStoreV1Alpha1Resource)(nil)
)

type ExternalSecretsIoClusterSecretStoreV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ExternalSecretsIoClusterSecretStoreV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Controller *string `tfsdk:"controller" yaml:"controller,omitempty"`

		Provider *struct {
			Akeyless *struct {
				AkeylessGWApiURL *string `tfsdk:"akeyless_gw_api_url" yaml:"akeylessGWApiURL,omitempty"`

				AuthSecretRef *struct {
					KubernetesAuth *struct {
						AccessID *string `tfsdk:"access_id" yaml:"accessID,omitempty"`

						K8sConfName *string `tfsdk:"k8s_conf_name" yaml:"k8sConfName,omitempty"`

						SecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" yaml:"serviceAccountRef,omitempty"`
					} `tfsdk:"kubernetes_auth" yaml:"kubernetesAuth,omitempty"`

					SecretRef *struct {
						AccessID *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_id" yaml:"accessID,omitempty"`

						AccessType *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_type" yaml:"accessType,omitempty"`

						AccessTypeParam *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_type_param" yaml:"accessTypeParam,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"auth_secret_ref" yaml:"authSecretRef,omitempty"`
			} `tfsdk:"akeyless" yaml:"akeyless,omitempty"`

			Alibaba *struct {
				Auth *struct {
					SecretRef *struct {
						AccessKeyIDSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" yaml:"accessKeyIDSecretRef,omitempty"`

						AccessKeySecretSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_key_secret_secret_ref" yaml:"accessKeySecretSecretRef,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				RegionID *string `tfsdk:"region_id" yaml:"regionID,omitempty"`
			} `tfsdk:"alibaba" yaml:"alibaba,omitempty"`

			Aws *struct {
				Auth *struct {
					Jwt *struct {
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" yaml:"serviceAccountRef,omitempty"`
					} `tfsdk:"jwt" yaml:"jwt,omitempty"`

					SecretRef *struct {
						AccessKeyIDSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" yaml:"accessKeyIDSecretRef,omitempty"`

						SecretAccessKeySecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" yaml:"secretAccessKeySecretRef,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				Role *string `tfsdk:"role" yaml:"role,omitempty"`

				Service *string `tfsdk:"service" yaml:"service,omitempty"`
			} `tfsdk:"aws" yaml:"aws,omitempty"`

			Azurekv *struct {
				AuthSecretRef *struct {
					ClientId *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"client_id" yaml:"clientId,omitempty"`

					ClientSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`
				} `tfsdk:"auth_secret_ref" yaml:"authSecretRef,omitempty"`

				AuthType *string `tfsdk:"auth_type" yaml:"authType,omitempty"`

				IdentityId *string `tfsdk:"identity_id" yaml:"identityId,omitempty"`

				ServiceAccountRef *struct {
					Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"service_account_ref" yaml:"serviceAccountRef,omitempty"`

				TenantId *string `tfsdk:"tenant_id" yaml:"tenantId,omitempty"`

				VaultUrl *string `tfsdk:"vault_url" yaml:"vaultUrl,omitempty"`
			} `tfsdk:"azurekv" yaml:"azurekv,omitempty"`

			Fake *struct {
				Data *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueMap *map[string]string `tfsdk:"value_map" yaml:"valueMap,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"data" yaml:"data,omitempty"`
			} `tfsdk:"fake" yaml:"fake,omitempty"`

			Gcpsm *struct {
				Auth *struct {
					SecretRef *struct {
						SecretAccessKeySecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" yaml:"secretAccessKeySecretRef,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

					WorkloadIdentity *struct {
						ClusterLocation *string `tfsdk:"cluster_location" yaml:"clusterLocation,omitempty"`

						ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

						ClusterProjectID *string `tfsdk:"cluster_project_id" yaml:"clusterProjectID,omitempty"`

						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" yaml:"serviceAccountRef,omitempty"`
					} `tfsdk:"workload_identity" yaml:"workloadIdentity,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				ProjectID *string `tfsdk:"project_id" yaml:"projectID,omitempty"`
			} `tfsdk:"gcpsm" yaml:"gcpsm,omitempty"`

			Gitlab *struct {
				Auth *struct {
					SecretRef *struct {
						AccessToken *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"access_token" yaml:"accessToken,omitempty"`
					} `tfsdk:"secret_ref" yaml:"SecretRef,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				ProjectID *string `tfsdk:"project_id" yaml:"projectID,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`

			Ibm *struct {
				Auth *struct {
					SecretRef *struct {
						SecretApiKeySecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_api_key_secret_ref" yaml:"secretApiKeySecretRef,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				ServiceUrl *string `tfsdk:"service_url" yaml:"serviceUrl,omitempty"`
			} `tfsdk:"ibm" yaml:"ibm,omitempty"`

			Kubernetes *struct {
				Auth *struct {
					Cert *struct {
						ClientCert *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"client_cert" yaml:"clientCert,omitempty"`

						ClientKey *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"client_key" yaml:"clientKey,omitempty"`
					} `tfsdk:"cert" yaml:"cert,omitempty"`

					ServiceAccount *struct {
						ServiceAccount *struct {
							Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`
					} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

					Token *struct {
						BearerToken *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"bearer_token" yaml:"bearerToken,omitempty"`
					} `tfsdk:"token" yaml:"token,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				RemoteNamespace *string `tfsdk:"remote_namespace" yaml:"remoteNamespace,omitempty"`

				Server *struct {
					CaBundle *string `tfsdk:"ca_bundle" yaml:"caBundle,omitempty"`

					CaProvider *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"ca_provider" yaml:"caProvider,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"server" yaml:"server,omitempty"`
			} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

			Oracle *struct {
				Auth *struct {
					SecretRef *struct {
						Fingerprint *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"fingerprint" yaml:"fingerprint,omitempty"`

						Privatekey *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"privatekey" yaml:"privatekey,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

					Tenancy *string `tfsdk:"tenancy" yaml:"tenancy,omitempty"`

					User *string `tfsdk:"user" yaml:"user,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				Vault *string `tfsdk:"vault" yaml:"vault,omitempty"`
			} `tfsdk:"oracle" yaml:"oracle,omitempty"`

			Vault *struct {
				Auth *struct {
					AppRole *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						RoleId *string `tfsdk:"role_id" yaml:"roleId,omitempty"`

						SecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"app_role" yaml:"appRole,omitempty"`

					Cert *struct {
						ClientCert *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"client_cert" yaml:"clientCert,omitempty"`

						SecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"cert" yaml:"cert,omitempty"`

					Jwt *struct {
						KubernetesServiceAccountToken *struct {
							Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

							ExpirationSeconds *int64 `tfsdk:"expiration_seconds" yaml:"expirationSeconds,omitempty"`

							ServiceAccountRef *struct {
								Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"service_account_ref" yaml:"serviceAccountRef,omitempty"`
						} `tfsdk:"kubernetes_service_account_token" yaml:"kubernetesServiceAccountToken,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						SecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"jwt" yaml:"jwt,omitempty"`

					Kubernetes *struct {
						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						SecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" yaml:"serviceAccountRef,omitempty"`
					} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

					Ldap *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						SecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"ldap" yaml:"ldap,omitempty"`

					TokenSecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"token_secret_ref" yaml:"tokenSecretRef,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				CaBundle *string `tfsdk:"ca_bundle" yaml:"caBundle,omitempty"`

				CaProvider *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"ca_provider" yaml:"caProvider,omitempty"`

				ForwardInconsistent *bool `tfsdk:"forward_inconsistent" yaml:"forwardInconsistent,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				ReadYourWrites *bool `tfsdk:"read_your_writes" yaml:"readYourWrites,omitempty"`

				Server *string `tfsdk:"server" yaml:"server,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"vault" yaml:"vault,omitempty"`

			Webhook *struct {
				Body *string `tfsdk:"body" yaml:"body,omitempty"`

				CaBundle *string `tfsdk:"ca_bundle" yaml:"caBundle,omitempty"`

				CaProvider *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"ca_provider" yaml:"caProvider,omitempty"`

				Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Result *struct {
					JsonPath *string `tfsdk:"json_path" yaml:"jsonPath,omitempty"`
				} `tfsdk:"result" yaml:"result,omitempty"`

				Secrets *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					SecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"secrets" yaml:"secrets,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"webhook" yaml:"webhook,omitempty"`

			Yandexlockbox *struct {
				ApiEndpoint *string `tfsdk:"api_endpoint" yaml:"apiEndpoint,omitempty"`

				Auth *struct {
					AuthorizedKeySecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"authorized_key_secret_ref" yaml:"authorizedKeySecretRef,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				CaProvider *struct {
					CertSecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"cert_secret_ref" yaml:"certSecretRef,omitempty"`
				} `tfsdk:"ca_provider" yaml:"caProvider,omitempty"`
			} `tfsdk:"yandexlockbox" yaml:"yandexlockbox,omitempty"`
		} `tfsdk:"provider" yaml:"provider,omitempty"`

		RetrySettings *struct {
			MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`

			RetryInterval *string `tfsdk:"retry_interval" yaml:"retryInterval,omitempty"`
		} `tfsdk:"retry_settings" yaml:"retrySettings,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewExternalSecretsIoClusterSecretStoreV1Alpha1Resource() resource.Resource {
	return &ExternalSecretsIoClusterSecretStoreV1Alpha1Resource{}
}

func (r *ExternalSecretsIoClusterSecretStoreV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_secrets_io_cluster_secret_store_v1alpha1"
}

func (r *ExternalSecretsIoClusterSecretStoreV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterSecretStore represents a secure external location for storing secrets, which can be referenced as part of 'storeRef' fields.",
		MarkdownDescription: "ClusterSecretStore represents a secure external location for storing secrets, which can be referenced as part of 'storeRef' fields.",
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
				Description:         "SecretStoreSpec defines the desired state of SecretStore.",
				MarkdownDescription: "SecretStoreSpec defines the desired state of SecretStore.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"controller": {
						Description:         "Used to select the correct KES controller (think: ingress.ingressClassName) The KES controller is instantiated with a specific controller name and filters ES based on this property",
						MarkdownDescription: "Used to select the correct KES controller (think: ingress.ingressClassName) The KES controller is instantiated with a specific controller name and filters ES based on this property",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"provider": {
						Description:         "Used to configure the provider. Only one provider may be set",
						MarkdownDescription: "Used to configure the provider. Only one provider may be set",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"akeyless": {
								Description:         "Akeyless configures this store to sync secrets using Akeyless Vault provider",
								MarkdownDescription: "Akeyless configures this store to sync secrets using Akeyless Vault provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"akeyless_gw_api_url": {
										Description:         "Akeyless GW API Url from which the secrets to be fetched from.",
										MarkdownDescription: "Akeyless GW API Url from which the secrets to be fetched from.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"auth_secret_ref": {
										Description:         "Auth configures how the operator authenticates with Akeyless.",
										MarkdownDescription: "Auth configures how the operator authenticates with Akeyless.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"kubernetes_auth": {
												Description:         "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
												MarkdownDescription: "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_id": {
														Description:         "the Akeyless Kubernetes auth-method access-id",
														MarkdownDescription: "the Akeyless Kubernetes auth-method access-id",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"k8s_conf_name": {
														Description:         "Kubernetes-auth configuration name in Akeyless-Gateway",
														MarkdownDescription: "Kubernetes-auth configuration name in Akeyless-Gateway",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"service_account_ref": {
														Description:         "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
														MarkdownDescription: "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audiences": {
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "Reference to a Secret that contains the details to authenticate with Akeyless.",
												MarkdownDescription: "Reference to a Secret that contains the details to authenticate with Akeyless.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_id": {
														Description:         "The SecretAccessID is used for authentication",
														MarkdownDescription: "The SecretAccessID is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"access_type": {
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"access_type_param": {
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

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

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"alibaba": {
								Description:         "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
								MarkdownDescription: "Alibaba configures this store to sync secrets using Alibaba Cloud provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "AlibabaAuth contains a secretRef for credentials.",
										MarkdownDescription: "AlibabaAuth contains a secretRef for credentials.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_ref": {
												Description:         "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
												MarkdownDescription: "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_key_id_secret_ref": {
														Description:         "The AccessKeyID is used for authentication",
														MarkdownDescription: "The AccessKeyID is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"access_key_secret_secret_ref": {
														Description:         "The AccessKeySecret is used for authentication",
														MarkdownDescription: "The AccessKeySecret is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"endpoint": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"region_id": {
										Description:         "Alibaba Region to be used for the provider",
										MarkdownDescription: "Alibaba Region to be used for the provider",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"aws": {
								Description:         "AWS configures this store to sync secrets using AWS Secret Manager provider",
								MarkdownDescription: "AWS configures this store to sync secrets using AWS Secret Manager provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
										MarkdownDescription: "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"jwt": {
												Description:         "Authenticate against AWS using service account tokens.",
												MarkdownDescription: "Authenticate against AWS using service account tokens.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"service_account_ref": {
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audiences": {
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
												MarkdownDescription: "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_key_id_secret_ref": {
														Description:         "The AccessKeyID is used for authentication",
														MarkdownDescription: "The AccessKeyID is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"secret_access_key_secret_ref": {
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

									"region": {
										Description:         "AWS Region to be used for the provider",
										MarkdownDescription: "AWS Region to be used for the provider",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"role": {
										Description:         "Role is a Role ARN which the SecretManager provider will assume",
										MarkdownDescription: "Role is a Role ARN which the SecretManager provider will assume",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service": {
										Description:         "Service defines which service should be used to fetch the secrets",
										MarkdownDescription: "Service defines which service should be used to fetch the secrets",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("SecretsManager", "ParameterStore"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"azurekv": {
								Description:         "AzureKV configures this store to sync secrets using Azure Key Vault provider",
								MarkdownDescription: "AzureKV configures this store to sync secrets using Azure Key Vault provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth_secret_ref": {
										Description:         "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
										MarkdownDescription: "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"client_id": {
												Description:         "The Azure clientId of the service principle used for authentication.",
												MarkdownDescription: "The Azure clientId of the service principle used for authentication.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

											"client_secret": {
												Description:         "The Azure ClientSecret of the service principle used for authentication.",
												MarkdownDescription: "The Azure ClientSecret of the service principle used for authentication.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth_type": {
										Description:         "Auth type defines how to authenticate to the keyvault service. Valid values are: - 'ServicePrincipal' (default): Using a service principal (tenantId, clientId, clientSecret) - 'ManagedIdentity': Using Managed Identity assigned to the pod (see aad-pod-identity)",
										MarkdownDescription: "Auth type defines how to authenticate to the keyvault service. Valid values are: - 'ServicePrincipal' (default): Using a service principal (tenantId, clientId, clientSecret) - 'ManagedIdentity': Using Managed Identity assigned to the pod (see aad-pod-identity)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ServicePrincipal", "ManagedIdentity", "WorkloadIdentity"),
										},
									},

									"identity_id": {
										Description:         "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
										MarkdownDescription: "If multiple Managed Identity is assigned to the pod, you can select the one to be used",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_account_ref": {
										Description:         "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
										MarkdownDescription: "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"audiences": {
												Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
												MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "The name of the ServiceAccount resource being referred to.",
												MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

									"tenant_id": {
										Description:         "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
										MarkdownDescription: "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"vault_url": {
										Description:         "Vault Url from which the secrets to be fetched from.",
										MarkdownDescription: "Vault Url from which the secrets to be fetched from.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"fake": {
								Description:         "Fake configures a store with static key/value pairs",
								MarkdownDescription: "Fake configures a store with static key/value pairs",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"data": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_map": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "",
												MarkdownDescription: "",

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

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"gcpsm": {
								Description:         "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
								MarkdownDescription: "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth defines the information necessary to authenticate against GCP",
										MarkdownDescription: "Auth defines the information necessary to authenticate against GCP",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_access_key_secret_ref": {
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"workload_identity": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cluster_location": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"cluster_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"cluster_project_id": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_account_ref": {
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audiences": {
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"project_id": {
										Description:         "ProjectID project where secret is located",
										MarkdownDescription: "ProjectID project where secret is located",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"gitlab": {
								Description:         "Gitlab configures this store to sync secrets using Gitlab Variables provider",
								MarkdownDescription: "Gitlab configures this store to sync secrets using Gitlab Variables provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth configures how secret-manager authenticates with a GitLab instance.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with a GitLab instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_token": {
														Description:         "AccessToken is used for authentication.",
														MarkdownDescription: "AccessToken is used for authentication.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"project_id": {
										Description:         "ProjectID specifies a project where secrets are located.",
										MarkdownDescription: "ProjectID specifies a project where secrets are located.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
										MarkdownDescription: "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"ibm": {
								Description:         "IBM configures this store to sync secrets using IBM Cloud provider",
								MarkdownDescription: "IBM configures this store to sync secrets using IBM Cloud provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth configures how secret-manager authenticates with the IBM secrets manager.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the IBM secrets manager.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_api_key_secret_ref": {
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"service_url": {
										Description:         "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
										MarkdownDescription: "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"kubernetes": {
								Description:         "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
								MarkdownDescription: "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth configures how secret-manager authenticates with a Kubernetes instance.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with a Kubernetes instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cert": {
												Description:         "has both clientCert and clientKey as secretKeySelector",
												MarkdownDescription: "has both clientCert and clientKey as secretKeySelector",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_cert": {
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"client_key": {
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("service_account"), path.MatchRelative().AtParent().AtName("token")),
												},
											},

											"service_account": {
												Description:         "points to a service account that should be used for authentication",
												MarkdownDescription: "points to a service account that should be used for authentication",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"service_account": {
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audiences": {
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("cert"), path.MatchRelative().AtParent().AtName("token")),
												},
											},

											"token": {
												Description:         "use static token to authenticate with",
												MarkdownDescription: "use static token to authenticate with",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"bearer_token": {
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("cert"), path.MatchRelative().AtParent().AtName("service_account")),
												},
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"remote_namespace": {
										Description:         "Remote namespace to fetch the secrets from",
										MarkdownDescription: "Remote namespace to fetch the secrets from",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"server": {
										Description:         "configures the Kubernetes server Address.",
										MarkdownDescription: "configures the Kubernetes server Address.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_bundle": {
												Description:         "CABundle is a base64-encoded CA certificate",
												MarkdownDescription: "CABundle is a base64-encoded CA certificate",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"ca_provider": {
												Description:         "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
												MarkdownDescription: "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
														MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the object located at the provider type.",
														MarkdownDescription: "The name of the object located at the provider type.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "The namespace the Provider type is in.",
														MarkdownDescription: "The namespace the Provider type is in.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
														MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Secret", "ConfigMap"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": {
												Description:         "configures the Kubernetes server Address.",
												MarkdownDescription: "configures the Kubernetes server Address.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"oracle": {
								Description:         "Oracle configures this store to sync secrets using Oracle Vault provider",
								MarkdownDescription: "Oracle configures this store to sync secrets using Oracle Vault provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, use the instance principal, otherwise the user credentials specified in Auth.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, use the instance principal, otherwise the user credentials specified in Auth.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_ref": {
												Description:         "SecretRef to pass through sensitive information.",
												MarkdownDescription: "SecretRef to pass through sensitive information.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fingerprint": {
														Description:         "Fingerprint is the fingerprint of the API private key.",
														MarkdownDescription: "Fingerprint is the fingerprint of the API private key.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"privatekey": {
														Description:         "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
														MarkdownDescription: "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

												Required: true,
												Optional: false,
												Computed: false,
											},

											"tenancy": {
												Description:         "Tenancy is the tenancy OCID where user is located.",
												MarkdownDescription: "Tenancy is the tenancy OCID where user is located.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"user": {
												Description:         "User is an access OCID specific to the account.",
												MarkdownDescription: "User is an access OCID specific to the account.",

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

									"region": {
										Description:         "Region is the region where vault is located.",
										MarkdownDescription: "Region is the region where vault is located.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"vault": {
										Description:         "Vault is the vault's OCID of the specific vault where secret is located.",
										MarkdownDescription: "Vault is the vault's OCID of the specific vault where secret is located.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"vault": {
								Description:         "Vault configures this store to sync secrets using Hashi provider",
								MarkdownDescription: "Vault configures this store to sync secrets using Hashi provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth": {
										Description:         "Auth configures how secret-manager authenticates with the Vault server.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the Vault server.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"app_role": {
												Description:         "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
												MarkdownDescription: "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",
														MarkdownDescription: "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"role_id": {
														Description:         "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
														MarkdownDescription: "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
														MarkdownDescription: "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

											"cert": {
												Description:         "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
												MarkdownDescription: "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_cert": {
														Description:         "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
														MarkdownDescription: "ClientCert is a certificate to authenticate using the Cert Vault authentication method",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"secret_ref": {
														Description:         "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwt": {
												Description:         "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
												MarkdownDescription: "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"kubernetes_service_account_token": {
														Description:         "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the 'TokenRequest' API.",
														MarkdownDescription: "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the 'TokenRequest' API.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audiences": {
																Description:         "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified.",
																MarkdownDescription: "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expiration_seconds": {
																Description:         "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to 10 minutes.",
																MarkdownDescription: "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to 10 minutes.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_ref": {
																Description:         "Service account field containing the name of a kubernetes ServiceAccount.",
																MarkdownDescription: "Service account field containing the name of a kubernetes ServiceAccount.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"audiences": {
																		Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "The name of the ServiceAccount resource being referred to.",
																		MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"path": {
														Description:         "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
														MarkdownDescription: "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"role": {
														Description:         "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
														MarkdownDescription: "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
														MarkdownDescription: "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kubernetes": {
												Description:         "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
												MarkdownDescription: "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
														MarkdownDescription: "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"role": {
														Description:         "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
														MarkdownDescription: "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"service_account_ref": {
														Description:         "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
														MarkdownDescription: "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audiences": {
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ldap": {
												Description:         "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
												MarkdownDescription: "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
														MarkdownDescription: "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_ref": {
														Description:         "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

													"username": {
														Description:         "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
														MarkdownDescription: "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",

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

											"token_secret_ref": {
												Description:         "TokenSecretRef authenticates with Vault by presenting a token.",
												MarkdownDescription: "TokenSecretRef authenticates with Vault by presenting a token.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ca_bundle": {
										Description:         "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.Base64Validator(),
										},
									},

									"ca_provider": {
										Description:         "The provider for the CA bundle to use to validate Vault server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Vault server certificate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
												MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "The namespace the Provider type is in.",
												MarkdownDescription: "The namespace the Provider type is in.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Secret", "ConfigMap"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"forward_inconsistent": {
										Description:         "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
										MarkdownDescription: "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
										MarkdownDescription: "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
										MarkdownDescription: "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_your_writes": {
										Description:         "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
										MarkdownDescription: "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"server": {
										Description:         "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
										MarkdownDescription: "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"version": {
										Description:         "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
										MarkdownDescription: "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("v1", "v2"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("webhook"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"webhook": {
								Description:         "Webhook configures this store to sync secrets using a generic templated webhook",
								MarkdownDescription: "Webhook configures this store to sync secrets using a generic templated webhook",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"body": {
										Description:         "Body",
										MarkdownDescription: "Body",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ca_bundle": {
										Description:         "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.Base64Validator(),
										},
									},

									"ca_provider": {
										Description:         "The provider for the CA bundle to use to validate webhook server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate webhook server certificate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
												MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "The namespace the Provider type is in.",
												MarkdownDescription: "The namespace the Provider type is in.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Secret", "ConfigMap"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers": {
										Description:         "Headers",
										MarkdownDescription: "Headers",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "Webhook Method",
										MarkdownDescription: "Webhook Method",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"result": {
										Description:         "Result formatting",
										MarkdownDescription: "Result formatting",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"json_path": {
												Description:         "Json path of return value",
												MarkdownDescription: "Json path of return value",

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

									"secrets": {
										Description:         "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
										MarkdownDescription: "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of this secret in templates",
												MarkdownDescription: "Name of this secret in templates",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_ref": {
												Description:         "Secret ref to fill in credentials",
												MarkdownDescription: "Secret ref to fill in credentials",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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

									"timeout": {
										Description:         "Timeout",
										MarkdownDescription: "Timeout",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "Webhook url to call",
										MarkdownDescription: "Webhook url to call",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("yandexlockbox")),
								},
							},

							"yandexlockbox": {
								Description:         "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
								MarkdownDescription: "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_endpoint": {
										Description:         "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										MarkdownDescription: "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth": {
										Description:         "Auth defines the information necessary to authenticate against Yandex Lockbox",
										MarkdownDescription: "Auth defines the information necessary to authenticate against Yandex Lockbox",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorized_key_secret_ref": {
												Description:         "The authorized key used for authentication",
												MarkdownDescription: "The authorized key used for authentication",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ca_provider": {
										Description:         "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cert_secret_ref": {
												Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
												MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("akeyless"), path.MatchRelative().AtParent().AtName("alibaba"), path.MatchRelative().AtParent().AtName("aws"), path.MatchRelative().AtParent().AtName("azurekv"), path.MatchRelative().AtParent().AtName("fake"), path.MatchRelative().AtParent().AtName("gcpsm"), path.MatchRelative().AtParent().AtName("gitlab"), path.MatchRelative().AtParent().AtName("ibm"), path.MatchRelative().AtParent().AtName("kubernetes"), path.MatchRelative().AtParent().AtName("oracle"), path.MatchRelative().AtParent().AtName("vault"), path.MatchRelative().AtParent().AtName("webhook")),
								},
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"retry_settings": {
						Description:         "Used to configure http retries if failed",
						MarkdownDescription: "Used to configure http retries if failed",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"retry_interval": {
								Description:         "",
								MarkdownDescription: "",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ExternalSecretsIoClusterSecretStoreV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_external_secrets_io_cluster_secret_store_v1alpha1")

	var state ExternalSecretsIoClusterSecretStoreV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExternalSecretsIoClusterSecretStoreV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("external-secrets.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ClusterSecretStore")

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

func (r *ExternalSecretsIoClusterSecretStoreV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_cluster_secret_store_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ExternalSecretsIoClusterSecretStoreV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_external_secrets_io_cluster_secret_store_v1alpha1")

	var state ExternalSecretsIoClusterSecretStoreV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExternalSecretsIoClusterSecretStoreV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("external-secrets.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ClusterSecretStore")

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

func (r *ExternalSecretsIoClusterSecretStoreV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_external_secrets_io_cluster_secret_store_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
