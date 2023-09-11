/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1beta1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &ExternalSecretsIoSecretStoreV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExternalSecretsIoSecretStoreV1Beta1DataSource{}
)

func NewExternalSecretsIoSecretStoreV1Beta1DataSource() datasource.DataSource {
	return &ExternalSecretsIoSecretStoreV1Beta1DataSource{}
}

type ExternalSecretsIoSecretStoreV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExternalSecretsIoSecretStoreV1Beta1DataSourceData struct {
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
		Conditions *[]struct {
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
		} `tfsdk:"conditions" json:"conditions,omitempty"`
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
				AdditionalRoles *[]string `tfsdk:"additional_roles" json:"additionalRoles,omitempty"`
				Auth            *struct {
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
						SessionTokenSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"session_token_secret_ref" json:"sessionTokenSecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				ExternalID  *string `tfsdk:"external_id" json:"externalID,omitempty"`
				Region      *string `tfsdk:"region" json:"region,omitempty"`
				Role        *string `tfsdk:"role" json:"role,omitempty"`
				Service     *string `tfsdk:"service" json:"service,omitempty"`
				SessionTags *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"session_tags" json:"sessionTags,omitempty"`
				TransitiveTagKeys *[]string `tfsdk:"transitive_tag_keys" json:"transitiveTagKeys,omitempty"`
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
				EnvironmentType   *string `tfsdk:"environment_type" json:"environmentType,omitempty"`
				IdentityId        *string `tfsdk:"identity_id" json:"identityId,omitempty"`
				ServiceAccountRef *struct {
					Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
					Name      *string   `tfsdk:"name" json:"name,omitempty"`
					Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
				TenantId *string `tfsdk:"tenant_id" json:"tenantId,omitempty"`
				VaultUrl *string `tfsdk:"vault_url" json:"vaultUrl,omitempty"`
			} `tfsdk:"azurekv" json:"azurekv,omitempty"`
			Conjur *struct {
				Auth *struct {
					Apikey *struct {
						Account   *string `tfsdk:"account" json:"account,omitempty"`
						ApiKeyRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"api_key_ref" json:"apiKeyRef,omitempty"`
						UserRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"user_ref" json:"userRef,omitempty"`
					} `tfsdk:"apikey" json:"apikey,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"conjur" json:"conjur,omitempty"`
			Delinea *struct {
				ClientId *struct {
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientSecret *struct {
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				Tenant      *string `tfsdk:"tenant" json:"tenant,omitempty"`
				Tld         *string `tfsdk:"tld" json:"tld,omitempty"`
				UrlTemplate *string `tfsdk:"url_template" json:"urlTemplate,omitempty"`
			} `tfsdk:"delinea" json:"delinea,omitempty"`
			Doppler *struct {
				Auth *struct {
					SecretRef *struct {
						DopplerToken *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"doppler_token" json:"dopplerToken,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Config          *string `tfsdk:"config" json:"config,omitempty"`
				Format          *string `tfsdk:"format" json:"format,omitempty"`
				NameTransformer *string `tfsdk:"name_transformer" json:"nameTransformer,omitempty"`
				Project         *string `tfsdk:"project" json:"project,omitempty"`
			} `tfsdk:"doppler" json:"doppler,omitempty"`
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
				Environment       *string   `tfsdk:"environment" json:"environment,omitempty"`
				GroupIDs          *[]string `tfsdk:"group_i_ds" json:"groupIDs,omitempty"`
				InheritFromGroups *bool     `tfsdk:"inherit_from_groups" json:"inheritFromGroups,omitempty"`
				ProjectID         *string   `tfsdk:"project_id" json:"projectID,omitempty"`
				Url               *string   `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"gitlab" json:"gitlab,omitempty"`
			Ibm *struct {
				Auth *struct {
					ContainerAuth *struct {
						IamEndpoint   *string `tfsdk:"iam_endpoint" json:"iamEndpoint,omitempty"`
						Profile       *string `tfsdk:"profile" json:"profile,omitempty"`
						TokenLocation *string `tfsdk:"token_location" json:"tokenLocation,omitempty"`
					} `tfsdk:"container_auth" json:"containerAuth,omitempty"`
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
			Keepersecurity *struct {
				AuthRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"auth_ref" json:"authRef,omitempty"`
				FolderID *string `tfsdk:"folder_id" json:"folderID,omitempty"`
			} `tfsdk:"keepersecurity" json:"keepersecurity,omitempty"`
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
						Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
						Name      *string   `tfsdk:"name" json:"name,omitempty"`
						Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
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
			Onepassword *struct {
				Auth *struct {
					SecretRef *struct {
						ConnectTokenSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"connect_token_secret_ref" json:"connectTokenSecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				ConnectHost *string            `tfsdk:"connect_host" json:"connectHost,omitempty"`
				Vaults      *map[string]string `tfsdk:"vaults" json:"vaults,omitempty"`
			} `tfsdk:"onepassword" json:"onepassword,omitempty"`
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
			Scaleway *struct {
				AccessKey *struct {
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"access_key" json:"accessKey,omitempty"`
				ApiUrl    *string `tfsdk:"api_url" json:"apiUrl,omitempty"`
				ProjectId *string `tfsdk:"project_id" json:"projectId,omitempty"`
				Region    *string `tfsdk:"region" json:"region,omitempty"`
				SecretKey *struct {
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			} `tfsdk:"scaleway" json:"scaleway,omitempty"`
			Senhasegura *struct {
				Auth *struct {
					ClientId              *string `tfsdk:"client_id" json:"clientId,omitempty"`
					ClientSecretSecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"client_secret_secret_ref" json:"clientSecretSecretRef,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				IgnoreSslCertificate *bool   `tfsdk:"ignore_ssl_certificate" json:"ignoreSslCertificate,omitempty"`
				Module               *string `tfsdk:"module" json:"module,omitempty"`
				Url                  *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"senhasegura" json:"senhasegura,omitempty"`
			Vault *struct {
				Auth *struct {
					AppRole *struct {
						Path    *string `tfsdk:"path" json:"path,omitempty"`
						RoleId  *string `tfsdk:"role_id" json:"roleId,omitempty"`
						RoleRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"role_ref" json:"roleRef,omitempty"`
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
					Iam *struct {
						ExternalID *string `tfsdk:"external_id" json:"externalID,omitempty"`
						Jwt        *struct {
							ServiceAccountRef *struct {
								Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
								Name      *string   `tfsdk:"name" json:"name,omitempty"`
								Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
						} `tfsdk:"jwt" json:"jwt,omitempty"`
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						Region    *string `tfsdk:"region" json:"region,omitempty"`
						Role      *string `tfsdk:"role" json:"role,omitempty"`
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
							SessionTokenSecretRef *struct {
								Key       *string `tfsdk:"key" json:"key,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"session_token_secret_ref" json:"sessionTokenSecretRef,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						VaultAwsIamServerID *string `tfsdk:"vault_aws_iam_server_id" json:"vaultAwsIamServerID,omitempty"`
						VaultRole           *string `tfsdk:"vault_role" json:"vaultRole,omitempty"`
					} `tfsdk:"iam" json:"iam,omitempty"`
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
					UserPass *struct {
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						SecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						Username *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"user_pass" json:"userPass,omitempty"`
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
			Yandexcertificatemanager *struct {
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
			} `tfsdk:"yandexcertificatemanager" json:"yandexcertificatemanager,omitempty"`
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
		RefreshInterval *int64 `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
		RetrySettings   *struct {
			MaxRetries    *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
		} `tfsdk:"retry_settings" json:"retrySettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternalSecretsIoSecretStoreV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_secret_store_v1beta1"
}

func (r *ExternalSecretsIoSecretStoreV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "SecretStoreSpec defines the desired state of SecretStore.",
				MarkdownDescription: "SecretStoreSpec defines the desired state of SecretStore.",
				Attributes: map[string]schema.Attribute{
					"conditions": schema.ListNestedAttribute{
						Description:         "Used to constraint a ClusterSecretStore to specific namespaces. Relevant only to ClusterSecretStore",
						MarkdownDescription: "Used to constraint a ClusterSecretStore to specific namespaces. Relevant only to ClusterSecretStore",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"namespace_selector": schema.SingleNestedAttribute{
									Description:         "Choose namespace using a labelSelector",
									MarkdownDescription: "Choose namespace using a labelSelector",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
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

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"namespaces": schema.ListAttribute{
									Description:         "Choose namespaces by name",
									MarkdownDescription: "Choose namespaces by name",
									ElementType:         types.StringType,
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

					"controller": schema.StringAttribute{
						Description:         "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
						MarkdownDescription: "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
										Required:            false,
										Optional:            false,
										Computed:            true,
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
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"k8s_conf_name": schema.StringAttribute{
														Description:         "Kubernetes-auth configuration name in Akeyless-Gateway",
														MarkdownDescription: "Kubernetes-auth configuration name in Akeyless-Gateway",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"access_type": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"access_type_param": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ca_bundle": schema.StringAttribute{
										Description:         "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key where the CA certificate can be found in the Secret or ConfigMap.",
												MarkdownDescription: "The key where the CA certificate can be found in the Secret or ConfigMap.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
												MarkdownDescription: "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
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
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"oidc_token_file_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role_arn": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"session_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"access_key_secret_secret_ref": schema.SingleNestedAttribute{
														Description:         "The AccessKeySecret is used for authentication",
														MarkdownDescription: "The AccessKeySecret is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"region_id": schema.StringAttribute{
										Description:         "Alibaba Region to be used for the provider",
										MarkdownDescription: "Alibaba Region to be used for the provider",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"aws": schema.SingleNestedAttribute{
								Description:         "AWS configures this store to sync secrets using AWS Secret Manager provider",
								MarkdownDescription: "AWS configures this store to sync secrets using AWS Secret Manager provider",
								Attributes: map[string]schema.Attribute{
									"additional_roles": schema.ListAttribute{
										Description:         "AdditionalRoles is a chained list of Role ARNs which the SecretManager provider will sequentially assume before assuming Role",
										MarkdownDescription: "AdditionalRoles is a chained list of Role ARNs which the SecretManager provider will sequentially assume before assuming Role",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"secret_access_key_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"session_token_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
														MarkdownDescription: "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"external_id": schema.StringAttribute{
										Description:         "AWS External ID set on assumed IAM roles",
										MarkdownDescription: "AWS External ID set on assumed IAM roles",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"region": schema.StringAttribute{
										Description:         "AWS Region to be used for the provider",
										MarkdownDescription: "AWS Region to be used for the provider",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"role": schema.StringAttribute{
										Description:         "Role is a Role ARN which the SecretManager provider will assume",
										MarkdownDescription: "Role is a Role ARN which the SecretManager provider will assume",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service": schema.StringAttribute{
										Description:         "Service defines which service should be used to fetch the secrets",
										MarkdownDescription: "Service defines which service should be used to fetch the secrets",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"session_tags": schema.ListNestedAttribute{
										Description:         "AWS STS assume role session tags",
										MarkdownDescription: "AWS STS assume role session tags",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"transitive_tag_keys": schema.ListAttribute{
										Description:         "AWS STS assume role transitive session tags. Required when multiple rules are used with SecretStore",
										MarkdownDescription: "AWS STS assume role transitive session tags. Required when multiple rules are used with SecretStore",
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"client_secret": schema.SingleNestedAttribute{
												Description:         "The Azure ClientSecret of the service principle used for authentication.",
												MarkdownDescription: "The Azure ClientSecret of the service principle used for authentication.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

									"auth_type": schema.StringAttribute{
										Description:         "Auth type defines how to authenticate to the keyvault service. Valid values are: - 'ServicePrincipal' (default): Using a service principal (tenantId, clientId, clientSecret) - 'ManagedIdentity': Using Managed Identity assigned to the pod (see aad-pod-identity)",
										MarkdownDescription: "Auth type defines how to authenticate to the keyvault service. Valid values are: - 'ServicePrincipal' (default): Using a service principal (tenantId, clientId, clientSecret) - 'ManagedIdentity': Using Managed Identity assigned to the pod (see aad-pod-identity)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"environment_type": schema.StringAttribute{
										Description:         "EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure. By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud",
										MarkdownDescription: "EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure. By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"identity_id": schema.StringAttribute{
										Description:         "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
										MarkdownDescription: "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the ServiceAccount resource being referred to.",
												MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
										MarkdownDescription: "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"vault_url": schema.StringAttribute{
										Description:         "Vault Url from which the secrets to be fetched from.",
										MarkdownDescription: "Vault Url from which the secrets to be fetched from.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"conjur": schema.SingleNestedAttribute{
								Description:         "Conjur configures this store to sync secrets using conjur provider",
								MarkdownDescription: "Conjur configures this store to sync secrets using conjur provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"apikey": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"account": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"api_key_ref": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"user_ref": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ca_bundle": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"delinea": schema.SingleNestedAttribute{
								Description:         "Delinea DevOps Secrets Vault https://docs.delinea.com/online-help/products/devops-secrets-vault/current",
								MarkdownDescription: "Delinea DevOps Secrets Vault https://docs.delinea.com/online-help/products/devops-secrets-vault/current",
								Attributes: map[string]schema.Attribute{
									"client_id": schema.SingleNestedAttribute{
										Description:         "ClientID is the non-secret part of the credential.",
										MarkdownDescription: "ClientID is the non-secret part of the credential.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef references a key in a secret that will be used as value.",
												MarkdownDescription: "SecretRef references a key in a secret that will be used as value.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"value": schema.StringAttribute{
												Description:         "Value can be specified directly to set a value without using a secret.",
												MarkdownDescription: "Value can be specified directly to set a value without using a secret.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"client_secret": schema.SingleNestedAttribute{
										Description:         "ClientSecret is the secret part of the credential.",
										MarkdownDescription: "ClientSecret is the secret part of the credential.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef references a key in a secret that will be used as value.",
												MarkdownDescription: "SecretRef references a key in a secret that will be used as value.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"value": schema.StringAttribute{
												Description:         "Value can be specified directly to set a value without using a secret.",
												MarkdownDescription: "Value can be specified directly to set a value without using a secret.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"tenant": schema.StringAttribute{
										Description:         "Tenant is the chosen hostname / site name.",
										MarkdownDescription: "Tenant is the chosen hostname / site name.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tld": schema.StringAttribute{
										Description:         "TLD is based on the server location that was chosen during provisioning. If unset, defaults to 'com'.",
										MarkdownDescription: "TLD is based on the server location that was chosen during provisioning. If unset, defaults to 'com'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"url_template": schema.StringAttribute{
										Description:         "URLTemplate If unset, defaults to 'https://%s.secretsvaultcloud.%s/v1/%s%s'.",
										MarkdownDescription: "URLTemplate If unset, defaults to 'https://%s.secretsvaultcloud.%s/v1/%s%s'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"doppler": schema.SingleNestedAttribute{
								Description:         "Doppler configures this store to sync secrets using the Doppler provider",
								MarkdownDescription: "Doppler configures this store to sync secrets using the Doppler provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how the Operator authenticates with the Doppler API",
										MarkdownDescription: "Auth configures how the Operator authenticates with the Doppler API",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"doppler_token": schema.SingleNestedAttribute{
														Description:         "The DopplerToken is used for authentication. See https://docs.doppler.com/reference/api#authentication for auth token types. The Key attribute defaults to dopplerToken if not specified.",
														MarkdownDescription: "The DopplerToken is used for authentication. See https://docs.doppler.com/reference/api#authentication for auth token types. The Key attribute defaults to dopplerToken if not specified.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"config": schema.StringAttribute{
										Description:         "Doppler config (required if not using a Service Token)",
										MarkdownDescription: "Doppler config (required if not using a Service Token)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"format": schema.StringAttribute{
										Description:         "Format enables the downloading of secrets as a file (string)",
										MarkdownDescription: "Format enables the downloading of secrets as a file (string)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name_transformer": schema.StringAttribute{
										Description:         "Environment variable compatible name transforms that change secret names to a different format",
										MarkdownDescription: "Environment variable compatible name transforms that change secret names to a different format",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"project": schema.StringAttribute{
										Description:         "Doppler project (required if not using a Service Token)",
										MarkdownDescription: "Doppler project (required if not using a Service Token)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value_map": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"workload_identity": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cluster_location": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cluster_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cluster_project_id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"project_id": schema.StringAttribute{
										Description:         "ProjectID project where secret is located",
										MarkdownDescription: "ProjectID project where secret is located",
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"environment": schema.StringAttribute{
										Description:         "Environment environment_scope of gitlab CI/CD variables (Please see https://docs.gitlab.com/ee/ci/environments/#create-a-static-environment on how to create environments)",
										MarkdownDescription: "Environment environment_scope of gitlab CI/CD variables (Please see https://docs.gitlab.com/ee/ci/environments/#create-a-static-environment on how to create environments)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"group_i_ds": schema.ListAttribute{
										Description:         "GroupIDs specify, which gitlab groups to pull secrets from. Group secrets are read from left to right followed by the project variables.",
										MarkdownDescription: "GroupIDs specify, which gitlab groups to pull secrets from. Group secrets are read from left to right followed by the project variables.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"inherit_from_groups": schema.BoolAttribute{
										Description:         "InheritFromGroups specifies whether parent groups should be discovered and checked for secrets.",
										MarkdownDescription: "InheritFromGroups specifies whether parent groups should be discovered and checked for secrets.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"project_id": schema.StringAttribute{
										Description:         "ProjectID specifies a project where secrets are located.",
										MarkdownDescription: "ProjectID specifies a project where secrets are located.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"url": schema.StringAttribute{
										Description:         "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
										MarkdownDescription: "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ibm": schema.SingleNestedAttribute{
								Description:         "IBM configures this store to sync secrets using IBM Cloud provider",
								MarkdownDescription: "IBM configures this store to sync secrets using IBM Cloud provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth configures how secret-manager authenticates with the IBM secrets manager.",
										MarkdownDescription: "Auth configures how secret-manager authenticates with the IBM secrets manager.",
										Attributes: map[string]schema.Attribute{
											"container_auth": schema.SingleNestedAttribute{
												Description:         "IBM Container-based auth with IAM Trusted Profile.",
												MarkdownDescription: "IBM Container-based auth with IAM Trusted Profile.",
												Attributes: map[string]schema.Attribute{
													"iam_endpoint": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"profile": schema.StringAttribute{
														Description:         "the IBM Trusted Profile",
														MarkdownDescription: "the IBM Trusted Profile",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"token_location": schema.StringAttribute{
														Description:         "Location the token is mounted on the pod",
														MarkdownDescription: "Location the token is mounted on the pod",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"service_url": schema.StringAttribute{
										Description:         "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
										MarkdownDescription: "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"keepersecurity": schema.SingleNestedAttribute{
								Description:         "KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider",
								MarkdownDescription: "KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider",
								Attributes: map[string]schema.Attribute{
									"auth_ref": schema.SingleNestedAttribute{
										Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
										MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the Secret resource being referred to.",
												MarkdownDescription: "The name of the Secret resource being referred to.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"folder_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"client_key": schema.SingleNestedAttribute{
														Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"service_account": schema.SingleNestedAttribute{
												Description:         "points to a service account that should be used for authentication",
												MarkdownDescription: "points to a service account that should be used for authentication",
												Attributes: map[string]schema.Attribute{
													"audiences": schema.ListAttribute{
														Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
														MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the ServiceAccount resource being referred to.",
														MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"remote_namespace": schema.StringAttribute{
										Description:         "Remote namespace to fetch the secrets from",
										MarkdownDescription: "Remote namespace to fetch the secrets from",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"server": schema.SingleNestedAttribute{
										Description:         "configures the Kubernetes server Address.",
										MarkdownDescription: "configures the Kubernetes server Address.",
										Attributes: map[string]schema.Attribute{
											"ca_bundle": schema.StringAttribute{
												Description:         "CABundle is a base64-encoded CA certificate",
												MarkdownDescription: "CABundle is a base64-encoded CA certificate",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ca_provider": schema.SingleNestedAttribute{
												Description:         "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
												MarkdownDescription: "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key where the CA certificate can be found in the Secret or ConfigMap.",
														MarkdownDescription: "The key where the CA certificate can be found in the Secret or ConfigMap.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the object located at the provider type.",
														MarkdownDescription: "The name of the object located at the provider type.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
														MarkdownDescription: "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
														MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
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
												Description:         "configures the Kubernetes server Address.",
												MarkdownDescription: "configures the Kubernetes server Address.",
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

							"onepassword": schema.SingleNestedAttribute{
								Description:         "OnePassword configures this store to sync secrets using the 1Password Cloud provider",
								MarkdownDescription: "OnePassword configures this store to sync secrets using the 1Password Cloud provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth defines the information necessary to authenticate against OnePassword Connect Server",
										MarkdownDescription: "Auth defines the information necessary to authenticate against OnePassword Connect Server",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "OnePasswordAuthSecretRef holds secret references for 1Password credentials.",
												MarkdownDescription: "OnePasswordAuthSecretRef holds secret references for 1Password credentials.",
												Attributes: map[string]schema.Attribute{
													"connect_token_secret_ref": schema.SingleNestedAttribute{
														Description:         "The ConnectToken is used for authentication to a 1Password Connect Server.",
														MarkdownDescription: "The ConnectToken is used for authentication to a 1Password Connect Server.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"connect_host": schema.StringAttribute{
										Description:         "ConnectHost defines the OnePassword Connect Server to connect to",
										MarkdownDescription: "ConnectHost defines the OnePassword Connect Server to connect to",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"vaults": schema.MapAttribute{
										Description:         "Vaults defines which OnePassword vaults to search in which order",
										MarkdownDescription: "Vaults defines which OnePassword vaults to search in which order",
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"privatekey": schema.SingleNestedAttribute{
														Description:         "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
														MarkdownDescription: "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"tenancy": schema.StringAttribute{
												Description:         "Tenancy is the tenancy OCID where user is located.",
												MarkdownDescription: "Tenancy is the tenancy OCID where user is located.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"user": schema.StringAttribute{
												Description:         "User is an access OCID specific to the account.",
												MarkdownDescription: "User is an access OCID specific to the account.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the region where vault is located.",
										MarkdownDescription: "Region is the region where vault is located.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"vault": schema.StringAttribute{
										Description:         "Vault is the vault's OCID of the specific vault where secret is located.",
										MarkdownDescription: "Vault is the vault's OCID of the specific vault where secret is located.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"scaleway": schema.SingleNestedAttribute{
								Description:         "Scaleway",
								MarkdownDescription: "Scaleway",
								Attributes: map[string]schema.Attribute{
									"access_key": schema.SingleNestedAttribute{
										Description:         "AccessKey is the non-secret part of the api key.",
										MarkdownDescription: "AccessKey is the non-secret part of the api key.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef references a key in a secret that will be used as value.",
												MarkdownDescription: "SecretRef references a key in a secret that will be used as value.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"value": schema.StringAttribute{
												Description:         "Value can be specified directly to set a value without using a secret.",
												MarkdownDescription: "Value can be specified directly to set a value without using a secret.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"api_url": schema.StringAttribute{
										Description:         "APIURL is the url of the api to use. Defaults to https://api.scaleway.com",
										MarkdownDescription: "APIURL is the url of the api to use. Defaults to https://api.scaleway.com",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"project_id": schema.StringAttribute{
										Description:         "ProjectID is the id of your project, which you can find in the console: https://console.scaleway.com/project/settings",
										MarkdownDescription: "ProjectID is the id of your project, which you can find in the console: https://console.scaleway.com/project/settings",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"region": schema.StringAttribute{
										Description:         "Region where your secrets are located: https://developers.scaleway.com/en/quickstart/#region-and-zone",
										MarkdownDescription: "Region where your secrets are located: https://developers.scaleway.com/en/quickstart/#region-and-zone",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_key": schema.SingleNestedAttribute{
										Description:         "SecretKey is the non-secret part of the api key.",
										MarkdownDescription: "SecretKey is the non-secret part of the api key.",
										Attributes: map[string]schema.Attribute{
											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef references a key in a secret that will be used as value.",
												MarkdownDescription: "SecretRef references a key in a secret that will be used as value.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"value": schema.StringAttribute{
												Description:         "Value can be specified directly to set a value without using a secret.",
												MarkdownDescription: "Value can be specified directly to set a value without using a secret.",
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

							"senhasegura": schema.SingleNestedAttribute{
								Description:         "Senhasegura configures this store to sync secrets using senhasegura provider",
								MarkdownDescription: "Senhasegura configures this store to sync secrets using senhasegura provider",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Auth defines parameters to authenticate in senhasegura",
										MarkdownDescription: "Auth defines parameters to authenticate in senhasegura",
										Attributes: map[string]schema.Attribute{
											"client_id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"client_secret_secret_ref": schema.SingleNestedAttribute{
												Description:         "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
												MarkdownDescription: "A reference to a specific 'key' within a Secret resource, In some instances, 'key' is a required field.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

									"ignore_ssl_certificate": schema.BoolAttribute{
										Description:         "IgnoreSslCertificate defines if SSL certificate must be ignored",
										MarkdownDescription: "IgnoreSslCertificate defines if SSL certificate must be ignored",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"module": schema.StringAttribute{
										Description:         "Module defines which senhasegura module should be used to get secrets",
										MarkdownDescription: "Module defines which senhasegura module should be used to get secrets",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"url": schema.StringAttribute{
										Description:         "URL of senhasegura",
										MarkdownDescription: "URL of senhasegura",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role_id": schema.StringAttribute{
														Description:         "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
														MarkdownDescription: "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role id.",
														MarkdownDescription: "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role id.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
														MarkdownDescription: "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"iam": schema.SingleNestedAttribute{
												Description:         "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
												MarkdownDescription: "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
												Attributes: map[string]schema.Attribute{
													"external_id": schema.StringAttribute{
														Description:         "AWS External ID set on assumed IAM roles",
														MarkdownDescription: "AWS External ID set on assumed IAM roles",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"jwt": schema.SingleNestedAttribute{
														Description:         "Specify a service account with IRSA enabled",
														MarkdownDescription: "Specify a service account with IRSA enabled",
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
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the ServiceAccount resource being referred to.",
																		MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

													"path": schema.StringAttribute{
														Description:         "Path where the AWS auth method is enabled in Vault, e.g: 'aws'",
														MarkdownDescription: "Path where the AWS auth method is enabled in Vault, e.g: 'aws'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"region": schema.StringAttribute{
														Description:         "AWS region",
														MarkdownDescription: "AWS region",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role": schema.StringAttribute{
														Description:         "This is the AWS role to be assumed before talking to vault",
														MarkdownDescription: "This is the AWS role to be assumed before talking to vault",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Specify credentials in a Secret object",
														MarkdownDescription: "Specify credentials in a Secret object",
														Attributes: map[string]schema.Attribute{
															"access_key_id_secret_ref": schema.SingleNestedAttribute{
																Description:         "The AccessKeyID is used for authentication",
																MarkdownDescription: "The AccessKeyID is used for authentication",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the Secret resource being referred to.",
																		MarkdownDescription: "The name of the Secret resource being referred to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"secret_access_key_secret_ref": schema.SingleNestedAttribute{
																Description:         "The SecretAccessKey is used for authentication",
																MarkdownDescription: "The SecretAccessKey is used for authentication",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the Secret resource being referred to.",
																		MarkdownDescription: "The name of the Secret resource being referred to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"session_token_secret_ref": schema.SingleNestedAttribute{
																Description:         "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																MarkdownDescription: "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the Secret resource being referred to.",
																		MarkdownDescription: "The name of the Secret resource being referred to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

													"vault_aws_iam_server_id": schema.StringAttribute{
														Description:         "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
														MarkdownDescription: "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"vault_role": schema.StringAttribute{
														Description:         "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
														MarkdownDescription: "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
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
																Description:         "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified. Deprecated: use serviceAccountRef.Audiences instead",
																MarkdownDescription: "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified. Deprecated: use serviceAccountRef.Audiences instead",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"expiration_seconds": schema.Int64Attribute{
																Description:         "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
																MarkdownDescription: "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
																Required:            false,
																Optional:            false,
																Computed:            true,
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
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the ServiceAccount resource being referred to.",
																		MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

													"path": schema.StringAttribute{
														Description:         "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
														MarkdownDescription: "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role": schema.StringAttribute{
														Description:         "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
														MarkdownDescription: "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
														MarkdownDescription: "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"kubernetes": schema.SingleNestedAttribute{
												Description:         "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
												MarkdownDescription: "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
														MarkdownDescription: "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role": schema.StringAttribute{
														Description:         "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
														MarkdownDescription: "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

											"ldap": schema.SingleNestedAttribute{
												Description:         "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
												MarkdownDescription: "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
														MarkdownDescription: "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"username": schema.StringAttribute{
														Description:         "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
														MarkdownDescription: "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"token_secret_ref": schema.SingleNestedAttribute{
												Description:         "TokenSecretRef authenticates with Vault by presenting a token.",
												MarkdownDescription: "TokenSecretRef authenticates with Vault by presenting a token.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"user_pass": schema.SingleNestedAttribute{
												Description:         "UserPass authenticates with Vault by passing username/password pair",
												MarkdownDescription: "UserPass authenticates with Vault by passing username/password pair",
												Attributes: map[string]schema.Attribute{
													"path": schema.StringAttribute{
														Description:         "Path where the UserPassword authentication backend is mounted in Vault, e.g: 'user'",
														MarkdownDescription: "Path where the UserPassword authentication backend is mounted in Vault, e.g: 'user'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
														MarkdownDescription: "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"username": schema.StringAttribute{
														Description:         "Username is a user name used to authenticate using the UserPass Vault authentication method",
														MarkdownDescription: "Username is a user name used to authenticate using the UserPass Vault authentication method",
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

									"ca_bundle": schema.StringAttribute{
										Description:         "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate Vault server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate Vault server certificate.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key where the CA certificate can be found in the Secret or ConfigMap.",
												MarkdownDescription: "The key where the CA certificate can be found in the Secret or ConfigMap.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
												MarkdownDescription: "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"forward_inconsistent": schema.BoolAttribute{
										Description:         "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
										MarkdownDescription: "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
										MarkdownDescription: "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"path": schema.StringAttribute{
										Description:         "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
										MarkdownDescription: "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"read_your_writes": schema.BoolAttribute{
										Description:         "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
										MarkdownDescription: "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"server": schema.StringAttribute{
										Description:         "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
										MarkdownDescription: "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"version": schema.StringAttribute{
										Description:         "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
										MarkdownDescription: "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"webhook": schema.SingleNestedAttribute{
								Description:         "Webhook configures this store to sync secrets using a generic templated webhook",
								MarkdownDescription: "Webhook configures this store to sync secrets using a generic templated webhook",
								Attributes: map[string]schema.Attribute{
									"body": schema.StringAttribute{
										Description:         "Body",
										MarkdownDescription: "Body",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ca_bundle": schema.StringAttribute{
										Description:         "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										MarkdownDescription: "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ca_provider": schema.SingleNestedAttribute{
										Description:         "The provider for the CA bundle to use to validate webhook server certificate.",
										MarkdownDescription: "The provider for the CA bundle to use to validate webhook server certificate.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key the value inside of the provider type to use, only used with 'Secret' type",
												MarkdownDescription: "The key the value inside of the provider type to use, only used with 'Secret' type",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the object located at the provider type.",
												MarkdownDescription: "The name of the object located at the provider type.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace the Provider type is in.",
												MarkdownDescription: "The namespace the Provider type is in.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"type": schema.StringAttribute{
												Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"headers": schema.MapAttribute{
										Description:         "Headers",
										MarkdownDescription: "Headers",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"method": schema.StringAttribute{
										Description:         "Webhook Method",
										MarkdownDescription: "Webhook Method",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"result": schema.SingleNestedAttribute{
										Description:         "Result formatting",
										MarkdownDescription: "Result formatting",
										Attributes: map[string]schema.Attribute{
											"json_path": schema.StringAttribute{
												Description:         "Json path of return value",
												MarkdownDescription: "Json path of return value",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secrets": schema.ListNestedAttribute{
										Description:         "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
										MarkdownDescription: "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of this secret in templates",
													MarkdownDescription: "Name of this secret in templates",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "Secret ref to fill in credentials",
													MarkdownDescription: "Secret ref to fill in credentials",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
															MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "The name of the Secret resource being referred to.",
															MarkdownDescription: "The name of the Secret resource being referred to.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
															MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"timeout": schema.StringAttribute{
										Description:         "Timeout",
										MarkdownDescription: "Timeout",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"url": schema.StringAttribute{
										Description:         "Webhook url to call",
										MarkdownDescription: "Webhook url to call",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"yandexcertificatemanager": schema.SingleNestedAttribute{
								Description:         "YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider",
								MarkdownDescription: "YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider",
								Attributes: map[string]schema.Attribute{
									"api_endpoint": schema.StringAttribute{
										Description:         "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										MarkdownDescription: "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"auth": schema.SingleNestedAttribute{
										Description:         "Auth defines the information necessary to authenticate against Yandex Certificate Manager",
										MarkdownDescription: "Auth defines the information necessary to authenticate against Yandex Certificate Manager",
										Attributes: map[string]schema.Attribute{
											"authorized_key_secret_ref": schema.SingleNestedAttribute{
												Description:         "The authorized key used for authentication",
												MarkdownDescription: "The authorized key used for authentication",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"yandexlockbox": schema.SingleNestedAttribute{
								Description:         "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
								MarkdownDescription: "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
								Attributes: map[string]schema.Attribute{
									"api_endpoint": schema.StringAttribute{
										Description:         "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										MarkdownDescription: "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
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

					"refresh_interval": schema.Int64Attribute{
						Description:         "Used to configure store refresh interval in seconds. Empty or 0 will default to the controller config.",
						MarkdownDescription: "Used to configure store refresh interval in seconds. Empty or 0 will default to the controller config.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"retry_settings": schema.SingleNestedAttribute{
						Description:         "Used to configure http retries if failed",
						MarkdownDescription: "Used to configure http retries if failed",
						Attributes: map[string]schema.Attribute{
							"max_retries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"retry_interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
		},
	}
}

func (r *ExternalSecretsIoSecretStoreV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ExternalSecretsIoSecretStoreV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_external_secrets_io_secret_store_v1beta1")

	var data ExternalSecretsIoSecretStoreV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "external-secrets.io", Version: "v1beta1", Resource: "secretstores"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ExternalSecretsIoSecretStoreV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("external-secrets.io/v1beta1")
	data.Kind = pointer.String("SecretStore")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
