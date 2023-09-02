/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cert_manager_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &CertManagerIoIssuerV1Resource{}
	_ resource.ResourceWithConfigure   = &CertManagerIoIssuerV1Resource{}
	_ resource.ResourceWithImportState = &CertManagerIoIssuerV1Resource{}
)

func NewCertManagerIoIssuerV1Resource() resource.Resource {
	return &CertManagerIoIssuerV1Resource{}
}

type CertManagerIoIssuerV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CertManagerIoIssuerV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Acme *struct {
			CaBundle                    *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			DisableAccountKeyGeneration *bool   `tfsdk:"disable_account_key_generation" json:"disableAccountKeyGeneration,omitempty"`
			Email                       *string `tfsdk:"email" json:"email,omitempty"`
			EnableDurationFeature       *bool   `tfsdk:"enable_duration_feature" json:"enableDurationFeature,omitempty"`
			ExternalAccountBinding      *struct {
				KeyAlgorithm *string `tfsdk:"key_algorithm" json:"keyAlgorithm,omitempty"`
				KeyID        *string `tfsdk:"key_id" json:"keyID,omitempty"`
				KeySecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"key_secret_ref" json:"keySecretRef,omitempty"`
			} `tfsdk:"external_account_binding" json:"externalAccountBinding,omitempty"`
			PreferredChain      *string `tfsdk:"preferred_chain" json:"preferredChain,omitempty"`
			PrivateKeySecretRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"private_key_secret_ref" json:"privateKeySecretRef,omitempty"`
			Server        *string `tfsdk:"server" json:"server,omitempty"`
			SkipTLSVerify *bool   `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			Solvers       *[]struct {
				Dns01 *struct {
					AcmeDNS *struct {
						AccountSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"account_secret_ref" json:"accountSecretRef,omitempty"`
						Host *string `tfsdk:"host" json:"host,omitempty"`
					} `tfsdk:"acme_dns" json:"acmeDNS,omitempty"`
					Akamai *struct {
						AccessTokenSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"access_token_secret_ref" json:"accessTokenSecretRef,omitempty"`
						ClientSecretSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"client_secret_secret_ref" json:"clientSecretSecretRef,omitempty"`
						ClientTokenSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"client_token_secret_ref" json:"clientTokenSecretRef,omitempty"`
						ServiceConsumerDomain *string `tfsdk:"service_consumer_domain" json:"serviceConsumerDomain,omitempty"`
					} `tfsdk:"akamai" json:"akamai,omitempty"`
					AzureDNS *struct {
						ClientID              *string `tfsdk:"client_id" json:"clientID,omitempty"`
						ClientSecretSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"client_secret_secret_ref" json:"clientSecretSecretRef,omitempty"`
						Environment     *string `tfsdk:"environment" json:"environment,omitempty"`
						HostedZoneName  *string `tfsdk:"hosted_zone_name" json:"hostedZoneName,omitempty"`
						ManagedIdentity *struct {
							ClientID   *string `tfsdk:"client_id" json:"clientID,omitempty"`
							ResourceID *string `tfsdk:"resource_id" json:"resourceID,omitempty"`
						} `tfsdk:"managed_identity" json:"managedIdentity,omitempty"`
						ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
						SubscriptionID    *string `tfsdk:"subscription_id" json:"subscriptionID,omitempty"`
						TenantID          *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
					} `tfsdk:"azure_dns" json:"azureDNS,omitempty"`
					CloudDNS *struct {
						HostedZoneName          *string `tfsdk:"hosted_zone_name" json:"hostedZoneName,omitempty"`
						Project                 *string `tfsdk:"project" json:"project,omitempty"`
						ServiceAccountSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"service_account_secret_ref" json:"serviceAccountSecretRef,omitempty"`
					} `tfsdk:"cloud_dns" json:"cloudDNS,omitempty"`
					Cloudflare *struct {
						ApiKeySecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"api_key_secret_ref" json:"apiKeySecretRef,omitempty"`
						ApiTokenSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"api_token_secret_ref" json:"apiTokenSecretRef,omitempty"`
						Email *string `tfsdk:"email" json:"email,omitempty"`
					} `tfsdk:"cloudflare" json:"cloudflare,omitempty"`
					CnameStrategy *string `tfsdk:"cname_strategy" json:"cnameStrategy,omitempty"`
					Digitalocean  *struct {
						TokenSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"token_secret_ref" json:"tokenSecretRef,omitempty"`
					} `tfsdk:"digitalocean" json:"digitalocean,omitempty"`
					Rfc2136 *struct {
						Nameserver          *string `tfsdk:"nameserver" json:"nameserver,omitempty"`
						TsigAlgorithm       *string `tfsdk:"tsig_algorithm" json:"tsigAlgorithm,omitempty"`
						TsigKeyName         *string `tfsdk:"tsig_key_name" json:"tsigKeyName,omitempty"`
						TsigSecretSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"tsig_secret_secret_ref" json:"tsigSecretSecretRef,omitempty"`
					} `tfsdk:"rfc2136" json:"rfc2136,omitempty"`
					Route53 *struct {
						AccessKeyID          *string `tfsdk:"access_key_id" json:"accessKeyID,omitempty"`
						AccessKeyIDSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" json:"accessKeyIDSecretRef,omitempty"`
						HostedZoneID             *string `tfsdk:"hosted_zone_id" json:"hostedZoneID,omitempty"`
						Region                   *string `tfsdk:"region" json:"region,omitempty"`
						Role                     *string `tfsdk:"role" json:"role,omitempty"`
						SecretAccessKeySecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" json:"secretAccessKeySecretRef,omitempty"`
					} `tfsdk:"route53" json:"route53,omitempty"`
					Webhook *struct {
						Config     *map[string]string `tfsdk:"config" json:"config,omitempty"`
						GroupName  *string            `tfsdk:"group_name" json:"groupName,omitempty"`
						SolverName *string            `tfsdk:"solver_name" json:"solverName,omitempty"`
					} `tfsdk:"webhook" json:"webhook,omitempty"`
				} `tfsdk:"dns01" json:"dns01,omitempty"`
				Http01 *struct {
					GatewayHTTPRoute *struct {
						Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						ParentRefs *[]struct {
							Group       *string `tfsdk:"group" json:"group,omitempty"`
							Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Port        *int64  `tfsdk:"port" json:"port,omitempty"`
							SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
						} `tfsdk:"parent_refs" json:"parentRefs,omitempty"`
						ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
					} `tfsdk:"gateway_http_route" json:"gatewayHTTPRoute,omitempty"`
					Ingress *struct {
						Class            *string `tfsdk:"class" json:"class,omitempty"`
						IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
						IngressTemplate  *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
								Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
						} `tfsdk:"ingress_template" json:"ingressTemplate,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						PodTemplate *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
								Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec *struct {
								Affinity *struct {
									NodeAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											Preference *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchFields *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_fields" json:"matchFields,omitempty"`
											} `tfsdk:"preference" json:"preference,omitempty"`
											Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
										RequiredDuringSchedulingIgnoredDuringExecution *struct {
											NodeSelectorTerms *[]struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchFields *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_fields" json:"matchFields,omitempty"`
											} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
									PodAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											PodAffinityTerm *struct {
												LabelSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
												NamespaceSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
												Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
												TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
											} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
											Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
										RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
											LabelSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
											NamespaceSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
											Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
											TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
									PodAntiAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											PodAffinityTerm *struct {
												LabelSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
												NamespaceSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
												Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
												TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
											} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
											Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
										RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
											LabelSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
											NamespaceSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
											Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
											TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
								} `tfsdk:"affinity" json:"affinity,omitempty"`
								ImagePullSecrets *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
								NodeSelector       *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
								PriorityClassName  *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
								ServiceAccountName *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
								Tolerations        *[]struct {
									Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
									TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
									Value             *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"tolerations" json:"tolerations,omitempty"`
							} `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"pod_template" json:"podTemplate,omitempty"`
						ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
					} `tfsdk:"ingress" json:"ingress,omitempty"`
				} `tfsdk:"http01" json:"http01,omitempty"`
				Selector *struct {
					DnsNames    *[]string          `tfsdk:"dns_names" json:"dnsNames,omitempty"`
					DnsZones    *[]string          `tfsdk:"dns_zones" json:"dnsZones,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
			} `tfsdk:"solvers" json:"solvers,omitempty"`
		} `tfsdk:"acme" json:"acme,omitempty"`
		Ca *struct {
			CrlDistributionPoints *[]string `tfsdk:"crl_distribution_points" json:"crlDistributionPoints,omitempty"`
			OcspServers           *[]string `tfsdk:"ocsp_servers" json:"ocspServers,omitempty"`
			SecretName            *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"ca" json:"ca,omitempty"`
		SelfSigned *struct {
			CrlDistributionPoints *[]string `tfsdk:"crl_distribution_points" json:"crlDistributionPoints,omitempty"`
		} `tfsdk:"self_signed" json:"selfSigned,omitempty"`
		Vault *struct {
			Auth *struct {
				AppRole *struct {
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					RoleId    *string `tfsdk:"role_id" json:"roleId,omitempty"`
					SecretRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"app_role" json:"appRole,omitempty"`
				Kubernetes *struct {
					MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					Role      *string `tfsdk:"role" json:"role,omitempty"`
					SecretRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					ServiceAccountRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
				TokenSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"token_secret_ref" json:"tokenSecretRef,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			CaBundle          *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			CaBundleSecretRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"ca_bundle_secret_ref" json:"caBundleSecretRef,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path      *string `tfsdk:"path" json:"path,omitempty"`
			Server    *string `tfsdk:"server" json:"server,omitempty"`
		} `tfsdk:"vault" json:"vault,omitempty"`
		Venafi *struct {
			Cloud *struct {
				ApiTokenSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"api_token_secret_ref" json:"apiTokenSecretRef,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"cloud" json:"cloud,omitempty"`
			Tpp *struct {
				CaBundle       *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
				CredentialsRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"credentials_ref" json:"credentialsRef,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"tpp" json:"tpp,omitempty"`
			Zone *string `tfsdk:"zone" json:"zone,omitempty"`
		} `tfsdk:"venafi" json:"venafi,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CertManagerIoIssuerV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cert_manager_io_issuer_v1"
}

func (r *CertManagerIoIssuerV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "An Issuer represents a certificate issuing authority which can be referenced as part of 'issuerRef' fields. It is scoped to a single namespace and can therefore only be referenced by resources within the same namespace.",
		MarkdownDescription: "An Issuer represents a certificate issuing authority which can be referenced as part of 'issuerRef' fields. It is scoped to a single namespace and can therefore only be referenced by resources within the same namespace.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Desired state of the Issuer resource.",
				MarkdownDescription: "Desired state of the Issuer resource.",
				Attributes: map[string]schema.Attribute{
					"acme": schema.SingleNestedAttribute{
						Description:         "ACME configures this issuer to communicate with a RFC8555 (ACME) server to obtain signed x509 certificates.",
						MarkdownDescription: "ACME configures this issuer to communicate with a RFC8555 (ACME) server to obtain signed x509 certificates.",
						Attributes: map[string]schema.Attribute{
							"ca_bundle": schema.StringAttribute{
								Description:         "Base64-encoded bundle of PEM CAs which can be used to validate the certificate chain presented by the ACME server. Mutually exclusive with SkipTLSVerify; prefer using CABundle to prevent various kinds of security vulnerabilities. If CABundle and SkipTLSVerify are unset, the system certificate bundle inside the container is used to validate the TLS connection.",
								MarkdownDescription: "Base64-encoded bundle of PEM CAs which can be used to validate the certificate chain presented by the ACME server. Mutually exclusive with SkipTLSVerify; prefer using CABundle to prevent various kinds of security vulnerabilities. If CABundle and SkipTLSVerify are unset, the system certificate bundle inside the container is used to validate the TLS connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"disable_account_key_generation": schema.BoolAttribute{
								Description:         "Enables or disables generating a new ACME account key. If true, the Issuer resource will *not* request a new account but will expect the account key to be supplied via an existing secret. If false, the cert-manager system will generate a new ACME account key for the Issuer. Defaults to false.",
								MarkdownDescription: "Enables or disables generating a new ACME account key. If true, the Issuer resource will *not* request a new account but will expect the account key to be supplied via an existing secret. If false, the cert-manager system will generate a new ACME account key for the Issuer. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"email": schema.StringAttribute{
								Description:         "Email is the email address to be associated with the ACME account. This field is optional, but it is strongly recommended to be set. It will be used to contact you in case of issues with your account or certificates, including expiry notification emails. This field may be updated after the account is initially registered.",
								MarkdownDescription: "Email is the email address to be associated with the ACME account. This field is optional, but it is strongly recommended to be set. It will be used to contact you in case of issues with your account or certificates, including expiry notification emails. This field may be updated after the account is initially registered.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_duration_feature": schema.BoolAttribute{
								Description:         "Enables requesting a Not After date on certificates that matches the duration of the certificate. This is not supported by all ACME servers like Let's Encrypt. If set to true when the ACME server does not support it it will create an error on the Order. Defaults to false.",
								MarkdownDescription: "Enables requesting a Not After date on certificates that matches the duration of the certificate. This is not supported by all ACME servers like Let's Encrypt. If set to true when the ACME server does not support it it will create an error on the Order. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_account_binding": schema.SingleNestedAttribute{
								Description:         "ExternalAccountBinding is a reference to a CA external account of the ACME server. If set, upon registration cert-manager will attempt to associate the given external account credentials with the registered ACME account.",
								MarkdownDescription: "ExternalAccountBinding is a reference to a CA external account of the ACME server. If set, upon registration cert-manager will attempt to associate the given external account credentials with the registered ACME account.",
								Attributes: map[string]schema.Attribute{
									"key_algorithm": schema.StringAttribute{
										Description:         "Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used. The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.",
										MarkdownDescription: "Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used. The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("HS256", "HS384", "HS512"),
										},
									},

									"key_id": schema.StringAttribute{
										Description:         "keyID is the ID of the CA key that the External Account is bound to.",
										MarkdownDescription: "keyID is the ID of the CA key that the External Account is bound to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"key_secret_ref": schema.SingleNestedAttribute{
										Description:         "keySecretRef is a Secret Key Selector referencing a data item in a Kubernetes Secret which holds the symmetric MAC key of the External Account Binding. The 'key' is the index string that is paired with the key data in the Secret and should not be confused with the key data itself, or indeed with the External Account Binding keyID above. The secret key stored in the Secret **must** be un-padded, base64 URL encoded data.",
										MarkdownDescription: "keySecretRef is a Secret Key Selector referencing a data item in a Kubernetes Secret which holds the symmetric MAC key of the External Account Binding. The 'key' is the index string that is paired with the key data in the Secret and should not be confused with the key data itself, or indeed with the External Account Binding keyID above. The secret key stored in the Secret **must** be un-padded, base64 URL encoded data.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            true,
												Optional:            false,
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

							"preferred_chain": schema.StringAttribute{
								Description:         "PreferredChain is the chain to use if the ACME server outputs multiple. PreferredChain is no guarantee that this one gets delivered by the ACME endpoint. For example, for Let's Encrypt's DST crosssign you would use: 'DST Root CA X3' or 'ISRG Root X1' for the newer Let's Encrypt root CA. This value picks the first certificate bundle in the ACME alternative chains that has a certificate with this value as its issuer's CN",
								MarkdownDescription: "PreferredChain is the chain to use if the ACME server outputs multiple. PreferredChain is no guarantee that this one gets delivered by the ACME endpoint. For example, for Let's Encrypt's DST crosssign you would use: 'DST Root CA X3' or 'ISRG Root X1' for the newer Let's Encrypt root CA. This value picks the first certificate bundle in the ACME alternative chains that has a certificate with this value as its issuer's CN",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(64),
								},
							},

							"private_key_secret_ref": schema.SingleNestedAttribute{
								Description:         "PrivateKey is the name of a Kubernetes Secret resource that will be used to store the automatically generated ACME account private key. Optionally, a 'key' may be specified to select a specific entry within the named Secret resource. If 'key' is not specified, a default of 'tls.key' will be used.",
								MarkdownDescription: "PrivateKey is the name of a Kubernetes Secret resource that will be used to store the automatically generated ACME account private key. Optionally, a 'key' may be specified to select a specific entry within the named Secret resource. If 'key' is not specified, a default of 'tls.key' will be used.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
										MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"server": schema.StringAttribute{
								Description:         "Server is the URL used to access the ACME server's 'directory' endpoint. For example, for Let's Encrypt's staging endpoint, you would use: 'https://acme-staging-v02.api.letsencrypt.org/directory'. Only ACME v2 endpoints (i.e. RFC 8555) are supported.",
								MarkdownDescription: "Server is the URL used to access the ACME server's 'directory' endpoint. For example, for Let's Encrypt's staging endpoint, you would use: 'https://acme-staging-v02.api.letsencrypt.org/directory'. Only ACME v2 endpoints (i.e. RFC 8555) are supported.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"skip_tls_verify": schema.BoolAttribute{
								Description:         "INSECURE: Enables or disables validation of the ACME server TLS certificate. If true, requests to the ACME server will not have the TLS certificate chain validated. Mutually exclusive with CABundle; prefer using CABundle to prevent various kinds of security vulnerabilities. Only enable this option in development environments. If CABundle and SkipTLSVerify are unset, the system certificate bundle inside the container is used to validate the TLS connection. Defaults to false.",
								MarkdownDescription: "INSECURE: Enables or disables validation of the ACME server TLS certificate. If true, requests to the ACME server will not have the TLS certificate chain validated. Mutually exclusive with CABundle; prefer using CABundle to prevent various kinds of security vulnerabilities. Only enable this option in development environments. If CABundle and SkipTLSVerify are unset, the system certificate bundle inside the container is used to validate the TLS connection. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"solvers": schema.ListNestedAttribute{
								Description:         "Solvers is a list of challenge solvers that will be used to solve ACME challenges for the matching domains. Solver configurations must be provided in order to obtain certificates from an ACME server. For more information, see: https://cert-manager.io/docs/configuration/acme/",
								MarkdownDescription: "Solvers is a list of challenge solvers that will be used to solve ACME challenges for the matching domains. Solver configurations must be provided in order to obtain certificates from an ACME server. For more information, see: https://cert-manager.io/docs/configuration/acme/",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dns01": schema.SingleNestedAttribute{
											Description:         "Configures cert-manager to attempt to complete authorizations by performing the DNS01 challenge flow.",
											MarkdownDescription: "Configures cert-manager to attempt to complete authorizations by performing the DNS01 challenge flow.",
											Attributes: map[string]schema.Attribute{
												"acme_dns": schema.SingleNestedAttribute{
													Description:         "Use the 'ACME DNS' (https://github.com/joohoi/acme-dns) API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the 'ACME DNS' (https://github.com/joohoi/acme-dns) API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"account_secret_ref": schema.SingleNestedAttribute{
															Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"host": schema.StringAttribute{
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

												"akamai": schema.SingleNestedAttribute{
													Description:         "Use the Akamai DNS zone management API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the Akamai DNS zone management API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"access_token_secret_ref": schema.SingleNestedAttribute{
															Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"client_secret_secret_ref": schema.SingleNestedAttribute{
															Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"client_token_secret_ref": schema.SingleNestedAttribute{
															Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"service_consumer_domain": schema.StringAttribute{
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

												"azure_dns": schema.SingleNestedAttribute{
													Description:         "Use the Microsoft Azure DNS API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the Microsoft Azure DNS API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"client_id": schema.StringAttribute{
															Description:         "if both this and ClientSecret are left unset MSI will be used",
															MarkdownDescription: "if both this and ClientSecret are left unset MSI will be used",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_secret_secret_ref": schema.SingleNestedAttribute{
															Description:         "if both this and ClientID are left unset MSI will be used",
															MarkdownDescription: "if both this and ClientID are left unset MSI will be used",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"environment": schema.StringAttribute{
															Description:         "name of the Azure environment (default AzurePublicCloud)",
															MarkdownDescription: "name of the Azure environment (default AzurePublicCloud)",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("AzurePublicCloud", "AzureChinaCloud", "AzureGermanCloud", "AzureUSGovernmentCloud"),
															},
														},

														"hosted_zone_name": schema.StringAttribute{
															Description:         "name of the DNS zone that should be used",
															MarkdownDescription: "name of the DNS zone that should be used",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"managed_identity": schema.SingleNestedAttribute{
															Description:         "managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID",
															MarkdownDescription: "managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID",
															Attributes: map[string]schema.Attribute{
																"client_id": schema.StringAttribute{
																	Description:         "client ID of the managed identity, can not be used at the same time as resourceID",
																	MarkdownDescription: "client ID of the managed identity, can not be used at the same time as resourceID",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource_id": schema.StringAttribute{
																	Description:         "resource ID of the managed identity, can not be used at the same time as clientID",
																	MarkdownDescription: "resource ID of the managed identity, can not be used at the same time as clientID",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"resource_group_name": schema.StringAttribute{
															Description:         "resource group the DNS zone is located in",
															MarkdownDescription: "resource group the DNS zone is located in",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"subscription_id": schema.StringAttribute{
															Description:         "ID of the Azure subscription",
															MarkdownDescription: "ID of the Azure subscription",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"tenant_id": schema.StringAttribute{
															Description:         "when specifying ClientID and ClientSecret then this field is also needed",
															MarkdownDescription: "when specifying ClientID and ClientSecret then this field is also needed",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"cloud_dns": schema.SingleNestedAttribute{
													Description:         "Use the Google Cloud DNS API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the Google Cloud DNS API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"hosted_zone_name": schema.StringAttribute{
															Description:         "HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created. If left empty cert-manager will automatically choose a zone.",
															MarkdownDescription: "HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created. If left empty cert-manager will automatically choose a zone.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"project": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service_account_secret_ref": schema.SingleNestedAttribute{
															Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

												"cloudflare": schema.SingleNestedAttribute{
													Description:         "Use the Cloudflare API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the Cloudflare API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"api_key_secret_ref": schema.SingleNestedAttribute{
															Description:         "API key to use to authenticate with Cloudflare. Note: using an API token to authenticate is now the recommended method as it allows greater control of permissions.",
															MarkdownDescription: "API key to use to authenticate with Cloudflare. Note: using an API token to authenticate is now the recommended method as it allows greater control of permissions.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"api_token_secret_ref": schema.SingleNestedAttribute{
															Description:         "API token used to authenticate with Cloudflare.",
															MarkdownDescription: "API token used to authenticate with Cloudflare.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"email": schema.StringAttribute{
															Description:         "Email of the account, only required when using API key based authentication.",
															MarkdownDescription: "Email of the account, only required when using API key based authentication.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"cname_strategy": schema.StringAttribute{
													Description:         "CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.",
													MarkdownDescription: "CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("None", "Follow"),
													},
												},

												"digitalocean": schema.SingleNestedAttribute{
													Description:         "Use the DigitalOcean DNS API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the DigitalOcean DNS API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"token_secret_ref": schema.SingleNestedAttribute{
															Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
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

												"rfc2136": schema.SingleNestedAttribute{
													Description:         "Use RFC2136 ('Dynamic Updates in the Domain Name System') (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.",
													MarkdownDescription: "Use RFC2136 ('Dynamic Updates in the Domain Name System') (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"nameserver": schema.StringAttribute{
															Description:         "The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port. If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.",
															MarkdownDescription: "The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port. If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"tsig_algorithm": schema.StringAttribute{
															Description:         "The TSIG Algorithm configured in the DNS supporting RFC2136. Used only when ''tsigSecretSecretRef'' and ''tsigKeyName'' are defined. Supported values are (case-insensitive): ''HMACMD5'' (default), ''HMACSHA1'', ''HMACSHA256'' or ''HMACSHA512''.",
															MarkdownDescription: "The TSIG Algorithm configured in the DNS supporting RFC2136. Used only when ''tsigSecretSecretRef'' and ''tsigKeyName'' are defined. Supported values are (case-insensitive): ''HMACMD5'' (default), ''HMACSHA1'', ''HMACSHA256'' or ''HMACSHA512''.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tsig_key_name": schema.StringAttribute{
															Description:         "The TSIG Key name configured in the DNS. If ''tsigSecretSecretRef'' is defined, this field is required.",
															MarkdownDescription: "The TSIG Key name configured in the DNS. If ''tsigSecretSecretRef'' is defined, this field is required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tsig_secret_secret_ref": schema.SingleNestedAttribute{
															Description:         "The name of the secret containing the TSIG value. If ''tsigKeyName'' is defined, this field is required.",
															MarkdownDescription: "The name of the secret containing the TSIG value. If ''tsigKeyName'' is defined, this field is required.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

												"route53": schema.SingleNestedAttribute{
													Description:         "Use the AWS Route53 API to manage DNS01 challenge records.",
													MarkdownDescription: "Use the AWS Route53 API to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"access_key_id": schema.StringAttribute{
															Description:         "The AccessKeyID is used for authentication. Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
															MarkdownDescription: "The AccessKeyID is used for authentication. Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"access_key_id_secret_ref": schema.SingleNestedAttribute{
															Description:         "The SecretAccessKey is used for authentication. If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
															MarkdownDescription: "The SecretAccessKey is used for authentication. If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"hosted_zone_id": schema.StringAttribute{
															Description:         "If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.",
															MarkdownDescription: "If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"region": schema.StringAttribute{
															Description:         "Always set the region when using AccessKeyID and SecretAccessKey",
															MarkdownDescription: "Always set the region when using AccessKeyID and SecretAccessKey",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"role": schema.StringAttribute{
															Description:         "Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata",
															MarkdownDescription: "Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret_access_key_secret_ref": schema.SingleNestedAttribute{
															Description:         "The SecretAccessKey is used for authentication. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
															MarkdownDescription: "The SecretAccessKey is used for authentication. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

												"webhook": schema.SingleNestedAttribute{
													Description:         "Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.",
													MarkdownDescription: "Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.",
													Attributes: map[string]schema.Attribute{
														"config": schema.MapAttribute{
															Description:         "Additional configuration that should be passed to the webhook apiserver when challenges are processed. This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.",
															MarkdownDescription: "Additional configuration that should be passed to the webhook apiserver when challenges are processed. This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"group_name": schema.StringAttribute{
															Description:         "The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver. This should be the same as the GroupName specified in the webhook provider implementation.",
															MarkdownDescription: "The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver. This should be the same as the GroupName specified in the webhook provider implementation.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"solver_name": schema.StringAttribute{
															Description:         "The name of the solver to use, as defined in the webhook provider implementation. This will typically be the name of the provider, e.g. 'cloudflare'.",
															MarkdownDescription: "The name of the solver to use, as defined in the webhook provider implementation. This will typically be the name of the provider, e.g. 'cloudflare'.",
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

										"http01": schema.SingleNestedAttribute{
											Description:         "Configures cert-manager to attempt to complete authorizations by performing the HTTP01 challenge flow. It is not possible to obtain certificates for wildcard domain names (e.g. '*.example.com') using the HTTP01 challenge mechanism.",
											MarkdownDescription: "Configures cert-manager to attempt to complete authorizations by performing the HTTP01 challenge flow. It is not possible to obtain certificates for wildcard domain names (e.g. '*.example.com') using the HTTP01 challenge mechanism.",
											Attributes: map[string]schema.Attribute{
												"gateway_http_route": schema.SingleNestedAttribute{
													Description:         "The Gateway API is a sig-network community API that models service networking in Kubernetes (https://gateway-api.sigs.k8s.io/). The Gateway solver will create HTTPRoutes with the specified labels in the same namespace as the challenge. This solver is experimental, and fields / behaviour may change in the future.",
													MarkdownDescription: "The Gateway API is a sig-network community API that models service networking in Kubernetes (https://gateway-api.sigs.k8s.io/). The Gateway solver will create HTTPRoutes with the specified labels in the same namespace as the challenge. This solver is experimental, and fields / behaviour may change in the future.",
													Attributes: map[string]schema.Attribute{
														"labels": schema.MapAttribute{
															Description:         "Custom labels that will be applied to HTTPRoutes created by cert-manager while solving HTTP-01 challenges.",
															MarkdownDescription: "Custom labels that will be applied to HTTPRoutes created by cert-manager while solving HTTP-01 challenges.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"parent_refs": schema.ListNestedAttribute{
															Description:         "When solving an HTTP-01 challenge, cert-manager creates an HTTPRoute. cert-manager needs to know which parentRefs should be used when creating the HTTPRoute. Usually, the parentRef references a Gateway. See: https://gateway-api.sigs.k8s.io/api-types/httproute/#attaching-to-gateways",
															MarkdownDescription: "When solving an HTTP-01 challenge, cert-manager creates an HTTPRoute. cert-manager needs to know which parentRefs should be used when creating the HTTPRoute. Usually, the parentRef references a Gateway. See: https://gateway-api.sigs.k8s.io/api-types/httproute/#attaching-to-gateways",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"group": schema.StringAttribute{
																		Description:         "Group is the group of the referent. When unspecified, 'gateway.networking.k8s.io' is inferred. To set the core API group (such as for a 'Service' kind referent), Group must be explicitly set to '' (empty string).  Support: Core",
																		MarkdownDescription: "Group is the group of the referent. When unspecified, 'gateway.networking.k8s.io' is inferred. To set the core API group (such as for a 'Service' kind referent), Group must be explicitly set to '' (empty string).  Support: Core",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(253),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																		},
																	},

																	"kind": schema.StringAttribute{
																		Description:         "Kind is kind of the referent.  Support: Core (Gateway)  Support: Implementation-specific (Other Resources)",
																		MarkdownDescription: "Kind is kind of the referent.  Support: Core (Gateway)  Support: Implementation-specific (Other Resources)",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(63),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																		},
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name is the name of the referent.  Support: Core",
																		MarkdownDescription: "Name is the name of the referent.  Support: Core",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(253),
																		},
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of the referent. When unspecified, this refers to the local namespace of the Route.  Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.  Support: Core",
																		MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, this refers to the local namespace of the Route.  Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.  Support: Core",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(63),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"port": schema.Int64Attribute{
																		Description:         "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  <gateway:experimental>",
																		MarkdownDescription: "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  <gateway:experimental>",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.Int64{
																			int64validator.AtLeast(1),
																			int64validator.AtMost(65535),
																		},
																	},

																	"section_name": schema.StringAttribute{
																		Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
																		MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.LengthAtMost(253),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_type": schema.StringAttribute{
															Description:         "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
															MarkdownDescription: "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"ingress": schema.SingleNestedAttribute{
													Description:         "The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.",
													MarkdownDescription: "The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.",
													Attributes: map[string]schema.Attribute{
														"class": schema.StringAttribute{
															Description:         "This field configures the annotation 'kubernetes.io/ingress.class' when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
															MarkdownDescription: "This field configures the annotation 'kubernetes.io/ingress.class' when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ingress_class_name": schema.StringAttribute{
															Description:         "This field configures the field 'ingressClassName' on the created Ingress resources used to solve ACME challenges that use this challenge solver. This is the recommended way of configuring the ingress class. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
															MarkdownDescription: "This field configures the field 'ingressClassName' on the created Ingress resources used to solve ACME challenges that use this challenge solver. This is the recommended way of configuring the ingress class. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ingress_template": schema.SingleNestedAttribute{
															Description:         "Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.",
															MarkdownDescription: "Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.",
															Attributes: map[string]schema.Attribute{
																"metadata": schema.SingleNestedAttribute{
																	Description:         "ObjectMeta overrides for the ingress used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",
																	MarkdownDescription: "ObjectMeta overrides for the ingress used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",
																	Attributes: map[string]schema.Attribute{
																		"annotations": schema.MapAttribute{
																			Description:         "Annotations that should be added to the created ACME HTTP01 solver ingress.",
																			MarkdownDescription: "Annotations that should be added to the created ACME HTTP01 solver ingress.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"labels": schema.MapAttribute{
																			Description:         "Labels that should be added to the created ACME HTTP01 solver ingress.",
																			MarkdownDescription: "Labels that should be added to the created ACME HTTP01 solver ingress.",
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"name": schema.StringAttribute{
															Description:         "The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges. This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
															MarkdownDescription: "The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges. This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_template": schema.SingleNestedAttribute{
															Description:         "Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.",
															MarkdownDescription: "Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.",
															Attributes: map[string]schema.Attribute{
																"metadata": schema.SingleNestedAttribute{
																	Description:         "ObjectMeta overrides for the pod used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",
																	MarkdownDescription: "ObjectMeta overrides for the pod used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",
																	Attributes: map[string]schema.Attribute{
																		"annotations": schema.MapAttribute{
																			Description:         "Annotations that should be added to the create ACME HTTP01 solver pods.",
																			MarkdownDescription: "Annotations that should be added to the create ACME HTTP01 solver pods.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"labels": schema.MapAttribute{
																			Description:         "Labels that should be added to the created ACME HTTP01 solver pods.",
																			MarkdownDescription: "Labels that should be added to the created ACME HTTP01 solver pods.",
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

																"spec": schema.SingleNestedAttribute{
																	Description:         "PodSpec defines overrides for the HTTP01 challenge solver pod. Check ACMEChallengeSolverHTTP01IngressPodSpec to find out currently supported fields. All other fields will be ignored.",
																	MarkdownDescription: "PodSpec defines overrides for the HTTP01 challenge solver pod. Check ACMEChallengeSolverHTTP01IngressPodSpec to find out currently supported fields. All other fields will be ignored.",
																	Attributes: map[string]schema.Attribute{
																		"affinity": schema.SingleNestedAttribute{
																			Description:         "If specified, the pod's scheduling constraints",
																			MarkdownDescription: "If specified, the pod's scheduling constraints",
																			Attributes: map[string]schema.Attribute{
																				"node_affinity": schema.SingleNestedAttribute{
																					Description:         "Describes node affinity scheduling rules for the pod.",
																					MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
																					Attributes: map[string]schema.Attribute{
																						"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																							Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																							MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"preference": schema.SingleNestedAttribute{
																										Description:         "A node selector term, associated with the corresponding weight.",
																										MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																										Attributes: map[string]schema.Attribute{
																											"match_expressions": schema.ListNestedAttribute{
																												Description:         "A list of node selector requirements by node's labels.",
																												MarkdownDescription: "A list of node selector requirements by node's labels.",
																												NestedObject: schema.NestedAttributeObject{
																													Attributes: map[string]schema.Attribute{
																														"key": schema.StringAttribute{
																															Description:         "The label key that the selector applies to.",
																															MarkdownDescription: "The label key that the selector applies to.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"operator": schema.StringAttribute{
																															Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"values": schema.ListAttribute{
																															Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																															MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																											"match_fields": schema.ListNestedAttribute{
																												Description:         "A list of node selector requirements by node's fields.",
																												MarkdownDescription: "A list of node selector requirements by node's fields.",
																												NestedObject: schema.NestedAttributeObject{
																													Attributes: map[string]schema.Attribute{
																														"key": schema.StringAttribute{
																															Description:         "The label key that the selector applies to.",
																															MarkdownDescription: "The label key that the selector applies to.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"operator": schema.StringAttribute{
																															Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"values": schema.ListAttribute{
																															Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																															MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																										},
																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"weight": schema.Int64Attribute{
																										Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																										MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

																						"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
																							Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																							MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																							Attributes: map[string]schema.Attribute{
																								"node_selector_terms": schema.ListNestedAttribute{
																									Description:         "Required. A list of node selector terms. The terms are ORed.",
																									MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"match_expressions": schema.ListNestedAttribute{
																												Description:         "A list of node selector requirements by node's labels.",
																												MarkdownDescription: "A list of node selector requirements by node's labels.",
																												NestedObject: schema.NestedAttributeObject{
																													Attributes: map[string]schema.Attribute{
																														"key": schema.StringAttribute{
																															Description:         "The label key that the selector applies to.",
																															MarkdownDescription: "The label key that the selector applies to.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"operator": schema.StringAttribute{
																															Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"values": schema.ListAttribute{
																															Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																															MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																											"match_fields": schema.ListNestedAttribute{
																												Description:         "A list of node selector requirements by node's fields.",
																												MarkdownDescription: "A list of node selector requirements by node's fields.",
																												NestedObject: schema.NestedAttributeObject{
																													Attributes: map[string]schema.Attribute{
																														"key": schema.StringAttribute{
																															Description:         "The label key that the selector applies to.",
																															MarkdownDescription: "The label key that the selector applies to.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"operator": schema.StringAttribute{
																															Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"values": schema.ListAttribute{
																															Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																															MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																				"pod_affinity": schema.SingleNestedAttribute{
																					Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																					MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																					Attributes: map[string]schema.Attribute{
																						"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																							Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																							MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"pod_affinity_term": schema.SingleNestedAttribute{
																										Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																										MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																										Attributes: map[string]schema.Attribute{
																											"label_selector": schema.SingleNestedAttribute{
																												Description:         "A label query over a set of resources, in this case pods.",
																												MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																											"namespace_selector": schema.SingleNestedAttribute{
																												Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																												MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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

																											"namespaces": schema.ListAttribute{
																												Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																												MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																												ElementType:         types.StringType,
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"topology_key": schema.StringAttribute{
																												Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																												MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																												Required:            true,
																												Optional:            false,
																												Computed:            false,
																											},
																										},
																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"weight": schema.Int64Attribute{
																										Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																										MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

																						"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																							Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																							MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"label_selector": schema.SingleNestedAttribute{
																										Description:         "A label query over a set of resources, in this case pods.",
																										MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																									"namespace_selector": schema.SingleNestedAttribute{
																										Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																										MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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

																									"namespaces": schema.ListAttribute{
																										Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																										MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"topology_key": schema.StringAttribute{
																										Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																										MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"pod_anti_affinity": schema.SingleNestedAttribute{
																					Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																					MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																					Attributes: map[string]schema.Attribute{
																						"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																							Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																							MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"pod_affinity_term": schema.SingleNestedAttribute{
																										Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																										MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																										Attributes: map[string]schema.Attribute{
																											"label_selector": schema.SingleNestedAttribute{
																												Description:         "A label query over a set of resources, in this case pods.",
																												MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																											"namespace_selector": schema.SingleNestedAttribute{
																												Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																												MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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

																											"namespaces": schema.ListAttribute{
																												Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																												MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																												ElementType:         types.StringType,
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"topology_key": schema.StringAttribute{
																												Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																												MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																												Required:            true,
																												Optional:            false,
																												Computed:            false,
																											},
																										},
																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"weight": schema.Int64Attribute{
																										Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																										MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

																						"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																							Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																							MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"label_selector": schema.SingleNestedAttribute{
																										Description:         "A label query over a set of resources, in this case pods.",
																										MarkdownDescription: "A label query over a set of resources, in this case pods.",
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

																									"namespace_selector": schema.SingleNestedAttribute{
																										Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																										MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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

																									"namespaces": schema.ListAttribute{
																										Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																										MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"topology_key": schema.StringAttribute{
																										Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																										MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

																		"image_pull_secrets": schema.ListNestedAttribute{
																			Description:         "If specified, the pod's imagePullSecrets",
																			MarkdownDescription: "If specified, the pod's imagePullSecrets",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

																		"node_selector": schema.MapAttribute{
																			Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																			MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"priority_class_name": schema.StringAttribute{
																			Description:         "If specified, the pod's priorityClassName.",
																			MarkdownDescription: "If specified, the pod's priorityClassName.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"service_account_name": schema.StringAttribute{
																			Description:         "If specified, the pod's service account",
																			MarkdownDescription: "If specified, the pod's service account",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"tolerations": schema.ListNestedAttribute{
																			Description:         "If specified, the pod's tolerations.",
																			MarkdownDescription: "If specified, the pod's tolerations.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"effect": schema.StringAttribute{
																						Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																						MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"key": schema.StringAttribute{
																						Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																						MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																						MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"toleration_seconds": schema.Int64Attribute{
																						Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																						MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
																						MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_type": schema.StringAttribute{
															Description:         "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
															MarkdownDescription: "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
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

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver. If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.",
											MarkdownDescription: "Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver. If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.",
											Attributes: map[string]schema.Attribute{
												"dns_names": schema.ListAttribute{
													Description:         "List of DNSNames that this solver will be used to solve. If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
													MarkdownDescription: "List of DNSNames that this solver will be used to solve. If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dns_zones": schema.ListAttribute{
													Description:         "List of DNSZones that this solver will be used to solve. The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
													MarkdownDescription: "List of DNSZones that this solver will be used to solve. The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "A label selector that is used to refine the set of certificate's that this challenge solver will apply to.",
													MarkdownDescription: "A label selector that is used to refine the set of certificate's that this challenge solver will apply to.",
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

					"ca": schema.SingleNestedAttribute{
						Description:         "CA configures this issuer to sign certificates using a signing CA keypair stored in a Secret resource. This is used to build internal PKIs that are managed by cert-manager.",
						MarkdownDescription: "CA configures this issuer to sign certificates using a signing CA keypair stored in a Secret resource. This is used to build internal PKIs that are managed by cert-manager.",
						Attributes: map[string]schema.Attribute{
							"crl_distribution_points": schema.ListAttribute{
								Description:         "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set, certificates will be issued without distribution points set.",
								MarkdownDescription: "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set, certificates will be issued without distribution points set.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ocsp_servers": schema.ListAttribute{
								Description:         "The OCSP server list is an X.509 v3 extension that defines a list of URLs of OCSP responders. The OCSP responders can be queried for the revocation status of an issued certificate. If not set, the certificate will be issued with no OCSP servers set. For example, an OCSP server URL could be 'http://ocsp.int-x3.letsencrypt.org'.",
								MarkdownDescription: "The OCSP server list is an X.509 v3 extension that defines a list of URLs of OCSP responders. The OCSP responders can be queried for the revocation status of an issued certificate. If not set, the certificate will be issued with no OCSP servers set. For example, an OCSP server URL could be 'http://ocsp.int-x3.letsencrypt.org'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of the secret used to sign Certificates issued by this Issuer.",
								MarkdownDescription: "SecretName is the name of the secret used to sign Certificates issued by this Issuer.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"self_signed": schema.SingleNestedAttribute{
						Description:         "SelfSigned configures this issuer to 'self sign' certificates using the private key used to create the CertificateRequest object.",
						MarkdownDescription: "SelfSigned configures this issuer to 'self sign' certificates using the private key used to create the CertificateRequest object.",
						Attributes: map[string]schema.Attribute{
							"crl_distribution_points": schema.ListAttribute{
								Description:         "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set certificate will be issued without CDP. Values are strings.",
								MarkdownDescription: "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set certificate will be issued without CDP. Values are strings.",
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

					"vault": schema.SingleNestedAttribute{
						Description:         "Vault configures this issuer to sign certificates using a HashiCorp Vault PKI backend.",
						MarkdownDescription: "Vault configures this issuer to sign certificates using a HashiCorp Vault PKI backend.",
						Attributes: map[string]schema.Attribute{
							"auth": schema.SingleNestedAttribute{
								Description:         "Auth configures how cert-manager authenticates with the Vault server.",
								MarkdownDescription: "Auth configures how cert-manager authenticates with the Vault server.",
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
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            true,
														Optional:            false,
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

									"kubernetes": schema.SingleNestedAttribute{
										Description:         "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
										MarkdownDescription: "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
										Attributes: map[string]schema.Attribute{
											"mount_path": schema.StringAttribute{
												Description:         "The Vault mountPath here is the mount path to use when authenticating with Vault. For example, setting a value to '/v1/auth/foo', will use the path '/v1/auth/foo/login' to authenticate with Vault. If unspecified, the default value '/v1/auth/kubernetes' will be used.",
												MarkdownDescription: "The Vault mountPath here is the mount path to use when authenticating with Vault. For example, setting a value to '/v1/auth/foo', will use the path '/v1/auth/foo/login' to authenticate with Vault. If unspecified, the default value '/v1/auth/kubernetes' will be used.",
												Required:            false,
												Optional:            true,
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
												Description:         "The required Secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. Use of 'ambient credentials' is not supported.",
												MarkdownDescription: "The required Secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. Use of 'ambient credentials' is not supported.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account_ref": schema.SingleNestedAttribute{
												Description:         "A reference to a service account that will be used to request a bound token (also known as 'projected token'). Compared to using 'secretRef', using this field means that you don't rely on statically bound tokens. To use this field, you must configure an RBAC rule to let cert-manager request a token.",
												MarkdownDescription: "A reference to a service account that will be used to request a bound token (also known as 'projected token'). Compared to using 'secretRef', using this field means that you don't rely on statically bound tokens. To use this field, you must configure an RBAC rule to let cert-manager request a token.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the ServiceAccount used to request a token.",
														MarkdownDescription: "Name of the ServiceAccount used to request a token.",
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
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"ca_bundle": schema.StringAttribute{
								Description:         "Base64-encoded bundle of PEM CAs which will be used to validate the certificate chain presented by Vault. Only used if using HTTPS to connect to Vault and ignored for HTTP connections. Mutually exclusive with CABundleSecretRef. If neither CABundle nor CABundleSecretRef are defined, the certificate bundle in the cert-manager controller container is used to validate the TLS connection.",
								MarkdownDescription: "Base64-encoded bundle of PEM CAs which will be used to validate the certificate chain presented by Vault. Only used if using HTTPS to connect to Vault and ignored for HTTP connections. Mutually exclusive with CABundleSecretRef. If neither CABundle nor CABundleSecretRef are defined, the certificate bundle in the cert-manager controller container is used to validate the TLS connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"ca_bundle_secret_ref": schema.SingleNestedAttribute{
								Description:         "Reference to a Secret containing a bundle of PEM-encoded CAs to use when verifying the certificate chain presented by Vault when using HTTPS. Mutually exclusive with CABundle. If neither CABundle nor CABundleSecretRef are defined, the certificate bundle in the cert-manager controller container is used to validate the TLS connection. If no key for the Secret is specified, cert-manager will default to 'ca.crt'.",
								MarkdownDescription: "Reference to a Secret containing a bundle of PEM-encoded CAs to use when verifying the certificate chain presented by Vault when using HTTPS. Mutually exclusive with CABundle. If neither CABundle nor CABundleSecretRef are defined, the certificate bundle in the cert-manager controller container is used to validate the TLS connection. If no key for the Secret is specified, cert-manager will default to 'ca.crt'.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
										MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1' More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
								MarkdownDescription: "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1' More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the mount path of the Vault PKI backend's 'sign' endpoint, e.g: 'my_pki_mount/sign/my-role-name'.",
								MarkdownDescription: "Path is the mount path of the Vault PKI backend's 'sign' endpoint, e.g: 'my_pki_mount/sign/my-role-name'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"server": schema.StringAttribute{
								Description:         "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
								MarkdownDescription: "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"venafi": schema.SingleNestedAttribute{
						Description:         "Venafi configures this issuer to sign certificates using a Venafi TPP or Venafi Cloud policy zone.",
						MarkdownDescription: "Venafi configures this issuer to sign certificates using a Venafi TPP or Venafi Cloud policy zone.",
						Attributes: map[string]schema.Attribute{
							"cloud": schema.SingleNestedAttribute{
								Description:         "Cloud specifies the Venafi cloud configuration settings. Only one of TPP or Cloud may be specified.",
								MarkdownDescription: "Cloud specifies the Venafi cloud configuration settings. Only one of TPP or Cloud may be specified.",
								Attributes: map[string]schema.Attribute{
									"api_token_secret_ref": schema.SingleNestedAttribute{
										Description:         "APITokenSecretRef is a secret key selector for the Venafi Cloud API token.",
										MarkdownDescription: "APITokenSecretRef is a secret key selector for the Venafi Cloud API token.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"url": schema.StringAttribute{
										Description:         "URL is the base URL for Venafi Cloud. Defaults to 'https://api.venafi.cloud/v1'.",
										MarkdownDescription: "URL is the base URL for Venafi Cloud. Defaults to 'https://api.venafi.cloud/v1'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tpp": schema.SingleNestedAttribute{
								Description:         "TPP specifies Trust Protection Platform configuration settings. Only one of TPP or Cloud may be specified.",
								MarkdownDescription: "TPP specifies Trust Protection Platform configuration settings. Only one of TPP or Cloud may be specified.",
								Attributes: map[string]schema.Attribute{
									"ca_bundle": schema.StringAttribute{
										Description:         "Base64-encoded bundle of PEM CAs which will be used to validate the certificate chain presented by the TPP server. Only used if using HTTPS; ignored for HTTP. If undefined, the certificate bundle in the cert-manager controller container is used to validate the chain.",
										MarkdownDescription: "Base64-encoded bundle of PEM CAs which will be used to validate the certificate chain presented by the TPP server. Only used if using HTTPS; ignored for HTTP. If undefined, the certificate bundle in the cert-manager controller container is used to validate the chain.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.Base64Validator(),
										},
									},

									"credentials_ref": schema.SingleNestedAttribute{
										Description:         "CredentialsRef is a reference to a Secret containing the username and password for the TPP server. The secret must contain two keys, 'username' and 'password'.",
										MarkdownDescription: "CredentialsRef is a reference to a Secret containing the username and password for the TPP server. The secret must contain two keys, 'username' and 'password'.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"url": schema.StringAttribute{
										Description:         "URL is the base URL for the vedsdk endpoint of the Venafi TPP instance, for example: 'https://tpp.example.com/vedsdk'.",
										MarkdownDescription: "URL is the base URL for the vedsdk endpoint of the Venafi TPP instance, for example: 'https://tpp.example.com/vedsdk'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"zone": schema.StringAttribute{
								Description:         "Zone is the Venafi Policy Zone to use for this issuer. All requests made to the Venafi platform will be restricted by the named zone policy. This field is required.",
								MarkdownDescription: "Zone is the Venafi Policy Zone to use for this issuer. All requests made to the Venafi platform will be restricted by the named zone policy. This field is required.",
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
	}
}

func (r *CertManagerIoIssuerV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CertManagerIoIssuerV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cert_manager_io_issuer_v1")

	var model CertManagerIoIssuerV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("cert-manager.io/v1")
	model.Kind = pointer.String("Issuer")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Issuer"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CertManagerIoIssuerV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CertManagerIoIssuerV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_issuer_v1")

	var data CertManagerIoIssuerV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Issuer"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse CertManagerIoIssuerV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CertManagerIoIssuerV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cert_manager_io_issuer_v1")

	var model CertManagerIoIssuerV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cert-manager.io/v1")
	model.Kind = pointer.String("Issuer")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Issuer"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CertManagerIoIssuerV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CertManagerIoIssuerV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cert_manager_io_issuer_v1")

	var data CertManagerIoIssuerV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "Issuer"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CertManagerIoIssuerV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
