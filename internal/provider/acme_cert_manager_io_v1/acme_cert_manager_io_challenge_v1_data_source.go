/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package acme_cert_manager_io_v1

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
	_ datasource.DataSource              = &AcmeCertManagerIoChallengeV1DataSource{}
	_ datasource.DataSourceWithConfigure = &AcmeCertManagerIoChallengeV1DataSource{}
)

func NewAcmeCertManagerIoChallengeV1DataSource() datasource.DataSource {
	return &AcmeCertManagerIoChallengeV1DataSource{}
}

type AcmeCertManagerIoChallengeV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AcmeCertManagerIoChallengeV1DataSourceData struct {
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
		AuthorizationURL *string `tfsdk:"authorization_url" json:"authorizationURL,omitempty"`
		DnsName          *string `tfsdk:"dns_name" json:"dnsName,omitempty"`
		IssuerRef        *struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"issuer_ref" json:"issuerRef,omitempty"`
		Key    *string `tfsdk:"key" json:"key,omitempty"`
		Solver *struct {
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
		} `tfsdk:"solver" json:"solver,omitempty"`
		Token    *string `tfsdk:"token" json:"token,omitempty"`
		Type     *string `tfsdk:"type" json:"type,omitempty"`
		Url      *string `tfsdk:"url" json:"url,omitempty"`
		Wildcard *bool   `tfsdk:"wildcard" json:"wildcard,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AcmeCertManagerIoChallengeV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_acme_cert_manager_io_challenge_v1"
}

func (r *AcmeCertManagerIoChallengeV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Challenge is a type to represent a Challenge request with an ACME server",
		MarkdownDescription: "Challenge is a type to represent a Challenge request with an ACME server",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"authorization_url": schema.StringAttribute{
						Description:         "The URL to the ACME Authorization resource that this challenge is a part of.",
						MarkdownDescription: "The URL to the ACME Authorization resource that this challenge is a part of.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"dns_name": schema.StringAttribute{
						Description:         "dnsName is the identifier that this challenge is for, e.g. example.com. If the requested DNSName is a 'wildcard', this field MUST be set to the non-wildcard domain, e.g. for '*.example.com', it must be 'example.com'.",
						MarkdownDescription: "dnsName is the identifier that this challenge is for, e.g. example.com. If the requested DNSName is a 'wildcard', this field MUST be set to the non-wildcard domain, e.g. for '*.example.com', it must be 'example.com'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"issuer_ref": schema.SingleNestedAttribute{
						Description:         "References a properly configured ACME-type Issuer which should be used to create this Challenge. If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Challenge will be marked as failed.",
						MarkdownDescription: "References a properly configured ACME-type Issuer which should be used to create this Challenge. If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Challenge will be marked as failed.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group of the resource being referred to.",
								MarkdownDescription: "Group of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the resource being referred to.",
								MarkdownDescription: "Kind of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the resource being referred to.",
								MarkdownDescription: "Name of the resource being referred to.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"key": schema.StringAttribute{
						Description:         "The ACME challenge key for this challenge For HTTP01 challenges, this is the value that must be responded with to complete the HTTP01 challenge in the format: '<private key JWK thumbprint>.<key from acme server for challenge>'. For DNS01 challenges, this is the base64 encoded SHA256 sum of the '<private key JWK thumbprint>.<key from acme server for challenge>' text that must be set as the TXT record content.",
						MarkdownDescription: "The ACME challenge key for this challenge For HTTP01 challenges, this is the value that must be responded with to complete the HTTP01 challenge in the format: '<private key JWK thumbprint>.<key from acme server for challenge>'. For DNS01 challenges, this is the base64 encoded SHA256 sum of the '<private key JWK thumbprint>.<key from acme server for challenge>' text that must be set as the TXT record content.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"solver": schema.SingleNestedAttribute{
						Description:         "Contains the domain solving configuration that should be used to solve this challenge resource.",
						MarkdownDescription: "Contains the domain solving configuration that should be used to solve this challenge resource.",
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"host": schema.StringAttribute{
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"client_secret_secret_ref": schema.SingleNestedAttribute{
												Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
												MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"client_token_secret_ref": schema.SingleNestedAttribute{
												Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
												MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"service_consumer_domain": schema.StringAttribute{
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

									"azure_dns": schema.SingleNestedAttribute{
										Description:         "Use the Microsoft Azure DNS API to manage DNS01 challenge records.",
										MarkdownDescription: "Use the Microsoft Azure DNS API to manage DNS01 challenge records.",
										Attributes: map[string]schema.Attribute{
											"client_id": schema.StringAttribute{
												Description:         "if both this and ClientSecret are left unset MSI will be used",
												MarkdownDescription: "if both this and ClientSecret are left unset MSI will be used",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"client_secret_secret_ref": schema.SingleNestedAttribute{
												Description:         "if both this and ClientID are left unset MSI will be used",
												MarkdownDescription: "if both this and ClientID are left unset MSI will be used",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"environment": schema.StringAttribute{
												Description:         "name of the Azure environment (default AzurePublicCloud)",
												MarkdownDescription: "name of the Azure environment (default AzurePublicCloud)",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"hosted_zone_name": schema.StringAttribute{
												Description:         "name of the DNS zone that should be used",
												MarkdownDescription: "name of the DNS zone that should be used",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"managed_identity": schema.SingleNestedAttribute{
												Description:         "managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID",
												MarkdownDescription: "managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID",
												Attributes: map[string]schema.Attribute{
													"client_id": schema.StringAttribute{
														Description:         "client ID of the managed identity, can not be used at the same time as resourceID",
														MarkdownDescription: "client ID of the managed identity, can not be used at the same time as resourceID",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource_id": schema.StringAttribute{
														Description:         "resource ID of the managed identity, can not be used at the same time as clientID",
														MarkdownDescription: "resource ID of the managed identity, can not be used at the same time as clientID",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"resource_group_name": schema.StringAttribute{
												Description:         "resource group the DNS zone is located in",
												MarkdownDescription: "resource group the DNS zone is located in",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"subscription_id": schema.StringAttribute{
												Description:         "ID of the Azure subscription",
												MarkdownDescription: "ID of the Azure subscription",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tenant_id": schema.StringAttribute{
												Description:         "when specifying ClientID and ClientSecret then this field is also needed",
												MarkdownDescription: "when specifying ClientID and ClientSecret then this field is also needed",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"cloud_dns": schema.SingleNestedAttribute{
										Description:         "Use the Google Cloud DNS API to manage DNS01 challenge records.",
										MarkdownDescription: "Use the Google Cloud DNS API to manage DNS01 challenge records.",
										Attributes: map[string]schema.Attribute{
											"hosted_zone_name": schema.StringAttribute{
												Description:         "HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created. If left empty cert-manager will automatically choose a zone.",
												MarkdownDescription: "HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created. If left empty cert-manager will automatically choose a zone.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"project": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"service_account_secret_ref": schema.SingleNestedAttribute{
												Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
												MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"api_token_secret_ref": schema.SingleNestedAttribute{
												Description:         "API token used to authenticate with Cloudflare.",
												MarkdownDescription: "API token used to authenticate with Cloudflare.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"email": schema.StringAttribute{
												Description:         "Email of the account, only required when using API key based authentication.",
												MarkdownDescription: "Email of the account, only required when using API key based authentication.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"cname_strategy": schema.StringAttribute{
										Description:         "CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.",
										MarkdownDescription: "CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"rfc2136": schema.SingleNestedAttribute{
										Description:         "Use RFC2136 ('Dynamic Updates in the Domain Name System') (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.",
										MarkdownDescription: "Use RFC2136 ('Dynamic Updates in the Domain Name System') (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.",
										Attributes: map[string]schema.Attribute{
											"nameserver": schema.StringAttribute{
												Description:         "The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port. If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.",
												MarkdownDescription: "The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port. If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tsig_algorithm": schema.StringAttribute{
												Description:         "The TSIG Algorithm configured in the DNS supporting RFC2136. Used only when ''tsigSecretSecretRef'' and ''tsigKeyName'' are defined. Supported values are (case-insensitive): ''HMACMD5'' (default), ''HMACSHA1'', ''HMACSHA256'' or ''HMACSHA512''.",
												MarkdownDescription: "The TSIG Algorithm configured in the DNS supporting RFC2136. Used only when ''tsigSecretSecretRef'' and ''tsigKeyName'' are defined. Supported values are (case-insensitive): ''HMACMD5'' (default), ''HMACSHA1'', ''HMACSHA256'' or ''HMACSHA512''.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tsig_key_name": schema.StringAttribute{
												Description:         "The TSIG Key name configured in the DNS. If ''tsigSecretSecretRef'' is defined, this field is required.",
												MarkdownDescription: "The TSIG Key name configured in the DNS. If ''tsigSecretSecretRef'' is defined, this field is required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"tsig_secret_secret_ref": schema.SingleNestedAttribute{
												Description:         "The name of the secret containing the TSIG value. If ''tsigKeyName'' is defined, this field is required.",
												MarkdownDescription: "The name of the secret containing the TSIG value. If ''tsigKeyName'' is defined, this field is required.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"route53": schema.SingleNestedAttribute{
										Description:         "Use the AWS Route53 API to manage DNS01 challenge records.",
										MarkdownDescription: "Use the AWS Route53 API to manage DNS01 challenge records.",
										Attributes: map[string]schema.Attribute{
											"access_key_id": schema.StringAttribute{
												Description:         "The AccessKeyID is used for authentication. Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
												MarkdownDescription: "The AccessKeyID is used for authentication. Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"access_key_id_secret_ref": schema.SingleNestedAttribute{
												Description:         "The SecretAccessKey is used for authentication. If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
												MarkdownDescription: "The SecretAccessKey is used for authentication. If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"hosted_zone_id": schema.StringAttribute{
												Description:         "If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.",
												MarkdownDescription: "If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"region": schema.StringAttribute{
												Description:         "Always set the region when using AccessKeyID and SecretAccessKey",
												MarkdownDescription: "Always set the region when using AccessKeyID and SecretAccessKey",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"role": schema.StringAttribute{
												Description:         "Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata",
												MarkdownDescription: "Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"secret_access_key_secret_ref": schema.SingleNestedAttribute{
												Description:         "The SecretAccessKey is used for authentication. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
												MarkdownDescription: "The SecretAccessKey is used for authentication. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"webhook": schema.SingleNestedAttribute{
										Description:         "Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.",
										MarkdownDescription: "Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.",
										Attributes: map[string]schema.Attribute{
											"config": schema.MapAttribute{
												Description:         "Additional configuration that should be passed to the webhook apiserver when challenges are processed. This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.",
												MarkdownDescription: "Additional configuration that should be passed to the webhook apiserver when challenges are processed. This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"group_name": schema.StringAttribute{
												Description:         "The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver. This should be the same as the GroupName specified in the webhook provider implementation.",
												MarkdownDescription: "The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver. This should be the same as the GroupName specified in the webhook provider implementation.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"solver_name": schema.StringAttribute{
												Description:         "The name of the solver to use, as defined in the webhook provider implementation. This will typically be the name of the provider, e.g. 'cloudflare'.",
												MarkdownDescription: "The name of the solver to use, as defined in the webhook provider implementation. This will typically be the name of the provider, e.g. 'cloudflare'.",
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
												Optional:            false,
												Computed:            true,
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
															Optional:            false,
															Computed:            true,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind is kind of the referent.  Support: Core (Gateway)  Support: Implementation-specific (Other Resources)",
															MarkdownDescription: "Kind is kind of the referent.  Support: Core (Gateway)  Support: Implementation-specific (Other Resources)",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of the referent.  Support: Core",
															MarkdownDescription: "Name is the name of the referent.  Support: Core",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace is the namespace of the referent. When unspecified, this refers to the local namespace of the Route.  Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.  Support: Core",
															MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, this refers to the local namespace of the Route.  Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.  Support: Core",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.Int64Attribute{
															Description:         "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  <gateway:experimental>",
															MarkdownDescription: "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  <gateway:experimental>",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"section_name": schema.StringAttribute{
															Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
															MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
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

											"service_type": schema.StringAttribute{
												Description:         "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
												MarkdownDescription: "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ingress": schema.SingleNestedAttribute{
										Description:         "The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.",
										MarkdownDescription: "The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.",
										Attributes: map[string]schema.Attribute{
											"class": schema.StringAttribute{
												Description:         "This field configures the annotation 'kubernetes.io/ingress.class' when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
												MarkdownDescription: "This field configures the annotation 'kubernetes.io/ingress.class' when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ingress_class_name": schema.StringAttribute{
												Description:         "This field configures the field 'ingressClassName' on the created Ingress resources used to solve ACME challenges that use this challenge solver. This is the recommended way of configuring the ingress class. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
												MarkdownDescription: "This field configures the field 'ingressClassName' on the created Ingress resources used to solve ACME challenges that use this challenge solver. This is the recommended way of configuring the ingress class. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Optional:            false,
																Computed:            true,
															},

															"labels": schema.MapAttribute{
																Description:         "Labels that should be added to the created ACME HTTP01 solver ingress.",
																MarkdownDescription: "Labels that should be added to the created ACME HTTP01 solver ingress.",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges. This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
												MarkdownDescription: "The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges. This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources. Only one of 'class', 'name' or 'ingressClassName' may be specified.",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Optional:            false,
																Computed:            true,
															},

															"labels": schema.MapAttribute{
																Description:         "Labels that should be added to the created ACME HTTP01 solver pods.",
																MarkdownDescription: "Labels that should be added to the created ACME HTTP01 solver pods.",
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																								"match_fields": schema.ListNestedAttribute{
																									Description:         "A list of node selector requirements by node's fields.",
																									MarkdownDescription: "A list of node selector requirements by node's fields.",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "The label key that the selector applies to.",
																												MarkdownDescription: "The label key that the selector applies to.",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"weight": schema.Int64Attribute{
																							Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																							MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																								"match_fields": schema.ListNestedAttribute{
																									Description:         "A list of node selector requirements by node's fields.",
																									MarkdownDescription: "A list of node selector requirements by node's fields.",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "The label key that the selector applies to.",
																												MarkdownDescription: "The label key that the selector applies to.",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																									Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"topology_key": schema.StringAttribute{
																									Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																									MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"weight": schema.Int64Attribute{
																							Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																							MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
																							Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"topology_key": schema.StringAttribute{
																							Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																							MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																									Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"topology_key": schema.StringAttribute{
																									Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																									MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"weight": schema.Int64Attribute{
																							Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																							MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
																							Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"topology_key": schema.StringAttribute{
																							Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																							MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																Required: false,
																Optional: false,
																Computed: true,
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
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"node_selector": schema.MapAttribute{
																Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"priority_class_name": schema.StringAttribute{
																Description:         "If specified, the pod's priorityClassName.",
																MarkdownDescription: "If specified, the pod's priorityClassName.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_name": schema.StringAttribute{
																Description:         "If specified, the pod's service account",
																MarkdownDescription: "If specified, the pod's service account",
																Required:            false,
																Optional:            false,
																Computed:            true,
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
																			Optional:            false,
																			Computed:            true,
																		},

																		"key": schema.StringAttribute{
																			Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																			MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																			MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"toleration_seconds": schema.Int64Attribute{
																			Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																			MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
																			MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
												Required: false,
												Optional: false,
												Computed: true,
											},

											"service_type": schema.StringAttribute{
												Description:         "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
												MarkdownDescription: "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
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

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver. If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.",
								MarkdownDescription: "Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver. If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.",
								Attributes: map[string]schema.Attribute{
									"dns_names": schema.ListAttribute{
										Description:         "List of DNSNames that this solver will be used to solve. If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
										MarkdownDescription: "List of DNSNames that this solver will be used to solve. If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"dns_zones": schema.ListAttribute{
										Description:         "List of DNSZones that this solver will be used to solve. The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
										MarkdownDescription: "List of DNSZones that this solver will be used to solve. The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"match_labels": schema.MapAttribute{
										Description:         "A label selector that is used to refine the set of certificate's that this challenge solver will apply to.",
										MarkdownDescription: "A label selector that is used to refine the set of certificate's that this challenge solver will apply to.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"token": schema.StringAttribute{
						Description:         "The ACME challenge token for this challenge. This is the raw value returned from the ACME server.",
						MarkdownDescription: "The ACME challenge token for this challenge. This is the raw value returned from the ACME server.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"type": schema.StringAttribute{
						Description:         "The type of ACME challenge this resource represents. One of 'HTTP-01' or 'DNS-01'.",
						MarkdownDescription: "The type of ACME challenge this resource represents. One of 'HTTP-01' or 'DNS-01'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"url": schema.StringAttribute{
						Description:         "The URL of the ACME Challenge resource for this challenge. This can be used to lookup details about the status of this challenge.",
						MarkdownDescription: "The URL of the ACME Challenge resource for this challenge. This can be used to lookup details about the status of this challenge.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"wildcard": schema.BoolAttribute{
						Description:         "wildcard will be true if this challenge is for a wildcard identifier, for example '*.example.com'.",
						MarkdownDescription: "wildcard will be true if this challenge is for a wildcard identifier, for example '*.example.com'.",
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
	}
}

func (r *AcmeCertManagerIoChallengeV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *AcmeCertManagerIoChallengeV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_acme_cert_manager_io_challenge_v1")

	var data AcmeCertManagerIoChallengeV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "acme.cert-manager.io", Version: "v1", Resource: "challenges"}).
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

	var readResponse AcmeCertManagerIoChallengeV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("acme.cert-manager.io/v1")
	data.Kind = pointer.String("Challenge")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
