/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tempo_grafana_com_v1alpha1

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
	_ datasource.DataSource = &TempoGrafanaComTempoStackV1Alpha1Manifest{}
)

func NewTempoGrafanaComTempoStackV1Alpha1Manifest() datasource.DataSource {
	return &TempoGrafanaComTempoStackV1Alpha1Manifest{}
}

type TempoGrafanaComTempoStackV1Alpha1Manifest struct{}

type TempoGrafanaComTempoStackV1Alpha1ManifestData struct {
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
		ExtraConfig *struct {
			Tempo *map[string]string `tfsdk:"tempo" json:"tempo,omitempty"`
		} `tfsdk:"extra_config" json:"extraConfig,omitempty"`
		HashRing *struct {
			Memberlist *struct {
				EnableIPv6 *bool `tfsdk:"enable_i_pv6" json:"enableIPv6,omitempty"`
			} `tfsdk:"memberlist" json:"memberlist,omitempty"`
		} `tfsdk:"hash_ring" json:"hashRing,omitempty"`
		Images *struct {
			OauthProxy      *string `tfsdk:"oauth_proxy" json:"oauthProxy,omitempty"`
			Tempo           *string `tfsdk:"tempo" json:"tempo,omitempty"`
			TempoGateway    *string `tfsdk:"tempo_gateway" json:"tempoGateway,omitempty"`
			TempoGatewayOpa *string `tfsdk:"tempo_gateway_opa" json:"tempoGatewayOpa,omitempty"`
			TempoQuery      *string `tfsdk:"tempo_query" json:"tempoQuery,omitempty"`
		} `tfsdk:"images" json:"images,omitempty"`
		Limits *struct {
			Global *struct {
				Ingestion *struct {
					IngestionBurstSizeBytes *int64 `tfsdk:"ingestion_burst_size_bytes" json:"ingestionBurstSizeBytes,omitempty"`
					IngestionRateLimitBytes *int64 `tfsdk:"ingestion_rate_limit_bytes" json:"ingestionRateLimitBytes,omitempty"`
					MaxBytesPerTrace        *int64 `tfsdk:"max_bytes_per_trace" json:"maxBytesPerTrace,omitempty"`
					MaxTracesPerUser        *int64 `tfsdk:"max_traces_per_user" json:"maxTracesPerUser,omitempty"`
				} `tfsdk:"ingestion" json:"ingestion,omitempty"`
				Query *struct {
					MaxBytesPerTagValues   *int64  `tfsdk:"max_bytes_per_tag_values" json:"maxBytesPerTagValues,omitempty"`
					MaxSearchBytesPerTrace *int64  `tfsdk:"max_search_bytes_per_trace" json:"maxSearchBytesPerTrace,omitempty"`
					MaxSearchDuration      *string `tfsdk:"max_search_duration" json:"maxSearchDuration,omitempty"`
				} `tfsdk:"query" json:"query,omitempty"`
			} `tfsdk:"global" json:"global,omitempty"`
			PerTenant *struct {
				Ingestion *struct {
					IngestionBurstSizeBytes *int64 `tfsdk:"ingestion_burst_size_bytes" json:"ingestionBurstSizeBytes,omitempty"`
					IngestionRateLimitBytes *int64 `tfsdk:"ingestion_rate_limit_bytes" json:"ingestionRateLimitBytes,omitempty"`
					MaxBytesPerTrace        *int64 `tfsdk:"max_bytes_per_trace" json:"maxBytesPerTrace,omitempty"`
					MaxTracesPerUser        *int64 `tfsdk:"max_traces_per_user" json:"maxTracesPerUser,omitempty"`
				} `tfsdk:"ingestion" json:"ingestion,omitempty"`
				Query *struct {
					MaxBytesPerTagValues   *int64  `tfsdk:"max_bytes_per_tag_values" json:"maxBytesPerTagValues,omitempty"`
					MaxSearchBytesPerTrace *int64  `tfsdk:"max_search_bytes_per_trace" json:"maxSearchBytesPerTrace,omitempty"`
					MaxSearchDuration      *string `tfsdk:"max_search_duration" json:"maxSearchDuration,omitempty"`
				} `tfsdk:"query" json:"query,omitempty"`
			} `tfsdk:"per_tenant" json:"perTenant,omitempty"`
		} `tfsdk:"limits" json:"limits,omitempty"`
		ManagementState *string `tfsdk:"management_state" json:"managementState,omitempty"`
		Observability   *struct {
			Grafana *struct {
				CreateDatasource *bool `tfsdk:"create_datasource" json:"createDatasource,omitempty"`
				InstanceSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"instance_selector" json:"instanceSelector,omitempty"`
			} `tfsdk:"grafana" json:"grafana,omitempty"`
			Metrics *struct {
				CreatePrometheusRules *bool `tfsdk:"create_prometheus_rules" json:"createPrometheusRules,omitempty"`
				CreateServiceMonitors *bool `tfsdk:"create_service_monitors" json:"createServiceMonitors,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Tracing *struct {
				Jaeger_agent_endpoint *string `tfsdk:"jaeger_agent_endpoint" json:"jaeger_agent_endpoint,omitempty"`
				Sampling_fraction     *string `tfsdk:"sampling_fraction" json:"sampling_fraction,omitempty"`
			} `tfsdk:"tracing" json:"tracing,omitempty"`
		} `tfsdk:"observability" json:"observability,omitempty"`
		ReplicationFactor *int64 `tfsdk:"replication_factor" json:"replicationFactor,omitempty"`
		Resources         *struct {
			Total *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"total" json:"total,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Retention *struct {
			Global *struct {
				Traces *string `tfsdk:"traces" json:"traces,omitempty"`
			} `tfsdk:"global" json:"global,omitempty"`
			PerTenant *struct {
				Traces *string `tfsdk:"traces" json:"traces,omitempty"`
			} `tfsdk:"per_tenant" json:"perTenant,omitempty"`
		} `tfsdk:"retention" json:"retention,omitempty"`
		Search *struct {
			DefaultResultLimit *int64  `tfsdk:"default_result_limit" json:"defaultResultLimit,omitempty"`
			MaxDuration        *string `tfsdk:"max_duration" json:"maxDuration,omitempty"`
			MaxResultLimit     *int64  `tfsdk:"max_result_limit" json:"maxResultLimit,omitempty"`
		} `tfsdk:"search" json:"search,omitempty"`
		ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		Storage        *struct {
			Secret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			Tls *struct {
				CaName     *string `tfsdk:"ca_name" json:"caName,omitempty"`
				CertName   *string `tfsdk:"cert_name" json:"certName,omitempty"`
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		StorageSize      *string `tfsdk:"storage_size" json:"storageSize,omitempty"`
		Template         *struct {
			Compactor *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources    *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Tolerations *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"compactor" json:"compactor,omitempty"`
			Distributor *struct {
				Component *struct {
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
					Resources    *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Tolerations *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Tls *struct {
					CaName     *string `tfsdk:"ca_name" json:"caName,omitempty"`
					CertName   *string `tfsdk:"cert_name" json:"certName,omitempty"`
					Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"distributor" json:"distributor,omitempty"`
			Gateway *struct {
				Component *struct {
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
					Resources    *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Tolerations *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				Ingress *struct {
					Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Host             *string            `tfsdk:"host" json:"host,omitempty"`
					IngressClassName *string            `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
					Route            *struct {
						Termination *string `tfsdk:"termination" json:"termination,omitempty"`
					} `tfsdk:"route" json:"route,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ingress" json:"ingress,omitempty"`
			} `tfsdk:"gateway" json:"gateway,omitempty"`
			Ingester *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources    *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Tolerations *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"ingester" json:"ingester,omitempty"`
			Querier *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources    *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Tolerations *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"querier" json:"querier,omitempty"`
			QueryFrontend *struct {
				Component *struct {
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
					Resources    *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Tolerations *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				JaegerQuery *struct {
					Authentication *struct {
						Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
						Sar     *string `tfsdk:"sar" json:"sar,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					Ingress *struct {
						Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Host             *string            `tfsdk:"host" json:"host,omitempty"`
						IngressClassName *string            `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
						Route            *struct {
							Termination *string `tfsdk:"termination" json:"termination,omitempty"`
						} `tfsdk:"route" json:"route,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ingress" json:"ingress,omitempty"`
					MonitorTab *struct {
						Enabled            *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
						PrometheusEndpoint *string `tfsdk:"prometheus_endpoint" json:"prometheusEndpoint,omitempty"`
					} `tfsdk:"monitor_tab" json:"monitorTab,omitempty"`
					Resources *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					ServicesQueryDuration *string `tfsdk:"services_query_duration" json:"servicesQueryDuration,omitempty"`
				} `tfsdk:"jaeger_query" json:"jaegerQuery,omitempty"`
			} `tfsdk:"query_frontend" json:"queryFrontend,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		Tenants *struct {
			Authentication *[]struct {
				Oidc *struct {
					GroupClaim  *string `tfsdk:"group_claim" json:"groupClaim,omitempty"`
					IssuerURL   *string `tfsdk:"issuer_url" json:"issuerURL,omitempty"`
					RedirectURL *string `tfsdk:"redirect_url" json:"redirectURL,omitempty"`
					Secret      *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					UsernameClaim *string `tfsdk:"username_claim" json:"usernameClaim,omitempty"`
				} `tfsdk:"oidc" json:"oidc,omitempty"`
				TenantId   *string `tfsdk:"tenant_id" json:"tenantId,omitempty"`
				TenantName *string `tfsdk:"tenant_name" json:"tenantName,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			Authorization *struct {
				RoleBindings *[]struct {
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
					Subjects *[]struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"subjects" json:"subjects,omitempty"`
				} `tfsdk:"role_bindings" json:"roleBindings,omitempty"`
				Roles *[]struct {
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Permissions *[]string `tfsdk:"permissions" json:"permissions,omitempty"`
					Resources   *[]string `tfsdk:"resources" json:"resources,omitempty"`
					Tenants     *[]string `tfsdk:"tenants" json:"tenants,omitempty"`
				} `tfsdk:"roles" json:"roles,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"tenants" json:"tenants,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TempoGrafanaComTempoStackV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tempo_grafana_com_tempo_stack_v1alpha1_manifest"
}

func (r *TempoGrafanaComTempoStackV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TempoStack manages a Tempo deployment in microservices mode.",
		MarkdownDescription: "TempoStack manages a Tempo deployment in microservices mode.",
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
				Description:         "TempoStackSpec defines the desired state of TempoStack.",
				MarkdownDescription: "TempoStackSpec defines the desired state of TempoStack.",
				Attributes: map[string]schema.Attribute{
					"extra_config": schema.SingleNestedAttribute{
						Description:         "ExtraConfigSpec defines extra configurations for tempo that will be merged with the operator generated, configurations defined here has precedence and could override generated config.",
						MarkdownDescription: "ExtraConfigSpec defines extra configurations for tempo that will be merged with the operator generated, configurations defined here has precedence and could override generated config.",
						Attributes: map[string]schema.Attribute{
							"tempo": schema.MapAttribute{
								Description:         "Tempo defines any extra Tempo configuration, which will be merged with the operator's generated Tempo configuration",
								MarkdownDescription: "Tempo defines any extra Tempo configuration, which will be merged with the operator's generated Tempo configuration",
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

					"hash_ring": schema.SingleNestedAttribute{
						Description:         "HashRing defines the spec for the distributed hash ring configuration.",
						MarkdownDescription: "HashRing defines the spec for the distributed hash ring configuration.",
						Attributes: map[string]schema.Attribute{
							"memberlist": schema.SingleNestedAttribute{
								Description:         "MemberList configuration spec",
								MarkdownDescription: "MemberList configuration spec",
								Attributes: map[string]schema.Attribute{
									"enable_i_pv6": schema.BoolAttribute{
										Description:         "EnableIPv6 enables IPv6 support for the memberlist based hash ring.",
										MarkdownDescription: "EnableIPv6 enables IPv6 support for the memberlist based hash ring.",
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

					"images": schema.SingleNestedAttribute{
						Description:         "Images defines the image for each container.",
						MarkdownDescription: "Images defines the image for each container.",
						Attributes: map[string]schema.Attribute{
							"oauth_proxy": schema.StringAttribute{
								Description:         "OauthProxy defines the oauth proxy image used to protect the jaegerUI on single tenant.",
								MarkdownDescription: "OauthProxy defines the oauth proxy image used to protect the jaegerUI on single tenant.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tempo": schema.StringAttribute{
								Description:         "Tempo defines the tempo container image.",
								MarkdownDescription: "Tempo defines the tempo container image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tempo_gateway": schema.StringAttribute{
								Description:         "TempoGateway defines the tempo-gateway container image.",
								MarkdownDescription: "TempoGateway defines the tempo-gateway container image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tempo_gateway_opa": schema.StringAttribute{
								Description:         "TempoGatewayOpa defines the OPA sidecar container for TempoGateway.",
								MarkdownDescription: "TempoGatewayOpa defines the OPA sidecar container for TempoGateway.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tempo_query": schema.StringAttribute{
								Description:         "TempoQuery defines the tempo-query container image.",
								MarkdownDescription: "TempoQuery defines the tempo-query container image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"limits": schema.SingleNestedAttribute{
						Description:         "LimitSpec is used to limit ingestion and querying rates.",
						MarkdownDescription: "LimitSpec is used to limit ingestion and querying rates.",
						Attributes: map[string]schema.Attribute{
							"global": schema.SingleNestedAttribute{
								Description:         "Global is used to define global rate limits.",
								MarkdownDescription: "Global is used to define global rate limits.",
								Attributes: map[string]schema.Attribute{
									"ingestion": schema.SingleNestedAttribute{
										Description:         "Ingestion is used to define ingestion rate limits.",
										MarkdownDescription: "Ingestion is used to define ingestion rate limits.",
										Attributes: map[string]schema.Attribute{
											"ingestion_burst_size_bytes": schema.Int64Attribute{
												Description:         "IngestionBurstSizeBytes defines the burst size (bytes) used in ingestion.",
												MarkdownDescription: "IngestionBurstSizeBytes defines the burst size (bytes) used in ingestion.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingestion_rate_limit_bytes": schema.Int64Attribute{
												Description:         "IngestionRateLimitBytes defines the Per-user ingestion rate limit (bytes) used in ingestion.",
												MarkdownDescription: "IngestionRateLimitBytes defines the Per-user ingestion rate limit (bytes) used in ingestion.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_bytes_per_trace": schema.Int64Attribute{
												Description:         "MaxBytesPerTrace defines the maximum number of bytes of an acceptable trace.",
												MarkdownDescription: "MaxBytesPerTrace defines the maximum number of bytes of an acceptable trace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_traces_per_user": schema.Int64Attribute{
												Description:         "MaxTracesPerUser defines the maximum number of traces a user can send.",
												MarkdownDescription: "MaxTracesPerUser defines the maximum number of traces a user can send.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"query": schema.SingleNestedAttribute{
										Description:         "Query is used to define query rate limits.",
										MarkdownDescription: "Query is used to define query rate limits.",
										Attributes: map[string]schema.Attribute{
											"max_bytes_per_tag_values": schema.Int64Attribute{
												Description:         "MaxBytesPerTagValues defines the maximum size in bytes of a tag-values query.",
												MarkdownDescription: "MaxBytesPerTagValues defines the maximum size in bytes of a tag-values query.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_search_bytes_per_trace": schema.Int64Attribute{
												Description:         "DEPRECATED. MaxSearchBytesPerTrace defines the maximum size of search data for a single trace in bytes. default: '0' to disable.",
												MarkdownDescription: "DEPRECATED. MaxSearchBytesPerTrace defines the maximum size of search data for a single trace in bytes. default: '0' to disable.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_search_duration": schema.StringAttribute{
												Description:         "MaxSearchDuration defines the maximum allowed time range for a search. If this value is not set, then spec.search.maxDuration is used.",
												MarkdownDescription: "MaxSearchDuration defines the maximum allowed time range for a search. If this value is not set, then spec.search.maxDuration is used.",
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

							"per_tenant": schema.SingleNestedAttribute{
								Description:         "PerTenant is used to define rate limits per tenant.",
								MarkdownDescription: "PerTenant is used to define rate limits per tenant.",
								Attributes: map[string]schema.Attribute{
									"ingestion": schema.SingleNestedAttribute{
										Description:         "Ingestion is used to define ingestion rate limits.",
										MarkdownDescription: "Ingestion is used to define ingestion rate limits.",
										Attributes: map[string]schema.Attribute{
											"ingestion_burst_size_bytes": schema.Int64Attribute{
												Description:         "IngestionBurstSizeBytes defines the burst size (bytes) used in ingestion.",
												MarkdownDescription: "IngestionBurstSizeBytes defines the burst size (bytes) used in ingestion.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingestion_rate_limit_bytes": schema.Int64Attribute{
												Description:         "IngestionRateLimitBytes defines the Per-user ingestion rate limit (bytes) used in ingestion.",
												MarkdownDescription: "IngestionRateLimitBytes defines the Per-user ingestion rate limit (bytes) used in ingestion.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_bytes_per_trace": schema.Int64Attribute{
												Description:         "MaxBytesPerTrace defines the maximum number of bytes of an acceptable trace.",
												MarkdownDescription: "MaxBytesPerTrace defines the maximum number of bytes of an acceptable trace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_traces_per_user": schema.Int64Attribute{
												Description:         "MaxTracesPerUser defines the maximum number of traces a user can send.",
												MarkdownDescription: "MaxTracesPerUser defines the maximum number of traces a user can send.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"query": schema.SingleNestedAttribute{
										Description:         "Query is used to define query rate limits.",
										MarkdownDescription: "Query is used to define query rate limits.",
										Attributes: map[string]schema.Attribute{
											"max_bytes_per_tag_values": schema.Int64Attribute{
												Description:         "MaxBytesPerTagValues defines the maximum size in bytes of a tag-values query.",
												MarkdownDescription: "MaxBytesPerTagValues defines the maximum size in bytes of a tag-values query.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_search_bytes_per_trace": schema.Int64Attribute{
												Description:         "DEPRECATED. MaxSearchBytesPerTrace defines the maximum size of search data for a single trace in bytes. default: '0' to disable.",
												MarkdownDescription: "DEPRECATED. MaxSearchBytesPerTrace defines the maximum size of search data for a single trace in bytes. default: '0' to disable.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_search_duration": schema.StringAttribute{
												Description:         "MaxSearchDuration defines the maximum allowed time range for a search. If this value is not set, then spec.search.maxDuration is used.",
												MarkdownDescription: "MaxSearchDuration defines the maximum allowed time range for a search. If this value is not set, then spec.search.maxDuration is used.",
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

					"management_state": schema.StringAttribute{
						Description:         "ManagementState defines if the CR should be managed by the operator or not. Default is managed.",
						MarkdownDescription: "ManagementState defines if the CR should be managed by the operator or not. Default is managed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Managed", "Unmanaged"),
						},
					},

					"observability": schema.SingleNestedAttribute{
						Description:         "ObservabilitySpec defines how telemetry data gets handled.",
						MarkdownDescription: "ObservabilitySpec defines how telemetry data gets handled.",
						Attributes: map[string]schema.Attribute{
							"grafana": schema.SingleNestedAttribute{
								Description:         "Grafana defines the Grafana configuration for operands.",
								MarkdownDescription: "Grafana defines the Grafana configuration for operands.",
								Attributes: map[string]schema.Attribute{
									"create_datasource": schema.BoolAttribute{
										Description:         "CreateDatasource specifies if a Grafana Datasource should be created for Tempo.",
										MarkdownDescription: "CreateDatasource specifies if a Grafana Datasource should be created for Tempo.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"instance_selector": schema.SingleNestedAttribute{
										Description:         "InstanceSelector specifies the Grafana instance where the datasource should be created.",
										MarkdownDescription: "InstanceSelector specifies the Grafana instance where the datasource should be created.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "Metrics defines the metrics configuration for operands.",
								MarkdownDescription: "Metrics defines the metrics configuration for operands.",
								Attributes: map[string]schema.Attribute{
									"create_prometheus_rules": schema.BoolAttribute{
										Description:         "CreatePrometheusRules specifies if Prometheus rules for alerts should be created for Tempo components.",
										MarkdownDescription: "CreatePrometheusRules specifies if Prometheus rules for alerts should be created for Tempo components.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"create_service_monitors": schema.BoolAttribute{
										Description:         "CreateServiceMonitors specifies if ServiceMonitors should be created for Tempo components.",
										MarkdownDescription: "CreateServiceMonitors specifies if ServiceMonitors should be created for Tempo components.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tracing": schema.SingleNestedAttribute{
								Description:         "Tracing defines a config for operands.",
								MarkdownDescription: "Tracing defines a config for operands.",
								Attributes: map[string]schema.Attribute{
									"jaeger_agent_endpoint": schema.StringAttribute{
										Description:         "JaegerAgentEndpoint defines the jaeger endpoint data gets send to.",
										MarkdownDescription: "JaegerAgentEndpoint defines the jaeger endpoint data gets send to.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sampling_fraction": schema.StringAttribute{
										Description:         "SamplingFraction defines the sampling ratio. Valid values are 0 to 1.",
										MarkdownDescription: "SamplingFraction defines the sampling ratio. Valid values are 0 to 1.",
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

					"replication_factor": schema.Int64Attribute{
						Description:         "ReplicationFactor is used to define how many component replicas should exist.",
						MarkdownDescription: "ReplicationFactor is used to define how many component replicas should exist.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources defines resources configuration.",
						MarkdownDescription: "Resources defines resources configuration.",
						Attributes: map[string]schema.Attribute{
							"total": schema.SingleNestedAttribute{
								Description:         "The total amount of resources for Tempo instance. The operator autonomously splits resources between deployed Tempo components. Only limits are supported, the operator calculates requests automatically. See http://github.com/grafana/tempo/issues/1540.",
								MarkdownDescription: "The total amount of resources for Tempo instance. The operator autonomously splits resources between deployed Tempo components. Only limits are supported, the operator calculates requests automatically. See http://github.com/grafana/tempo/issues/1540.",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"retention": schema.SingleNestedAttribute{
						Description:         "NOTE: currently this field is not considered. Retention period defined by dataset. User can specify how long data should be stored.",
						MarkdownDescription: "NOTE: currently this field is not considered. Retention period defined by dataset. User can specify how long data should be stored.",
						Attributes: map[string]schema.Attribute{
							"global": schema.SingleNestedAttribute{
								Description:         "Global is used to configure global retention.",
								MarkdownDescription: "Global is used to configure global retention.",
								Attributes: map[string]schema.Attribute{
									"traces": schema.StringAttribute{
										Description:         "Traces defines retention period. Supported parameter suffixes are 's', 'm' and 'h'. example: 336h default: value is 48h.",
										MarkdownDescription: "Traces defines retention period. Supported parameter suffixes are 's', 'm' and 'h'. example: 336h default: value is 48h.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_tenant": schema.SingleNestedAttribute{
								Description:         "PerTenant is used to configure retention per tenant.",
								MarkdownDescription: "PerTenant is used to configure retention per tenant.",
								Attributes: map[string]schema.Attribute{
									"traces": schema.StringAttribute{
										Description:         "Traces defines retention period. Supported parameter suffixes are 's', 'm' and 'h'. example: 336h default: value is 48h.",
										MarkdownDescription: "Traces defines retention period. Supported parameter suffixes are 's', 'm' and 'h'. example: 336h default: value is 48h.",
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

					"search": schema.SingleNestedAttribute{
						Description:         "SearchSpec control the configuration for the search capabilities.",
						MarkdownDescription: "SearchSpec control the configuration for the search capabilities.",
						Attributes: map[string]schema.Attribute{
							"default_result_limit": schema.Int64Attribute{
								Description:         "Limit used for search requests if none is set by the caller (default: 20)",
								MarkdownDescription: "Limit used for search requests if none is set by the caller (default: 20)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_duration": schema.StringAttribute{
								Description:         "The maximum allowed time range for a search, default: 0s which means unlimited.",
								MarkdownDescription: "The maximum allowed time range for a search, default: 0s which means unlimited.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_result_limit": schema.Int64Attribute{
								Description:         "The maximum allowed value of the limit parameter on search requests. If the search request limit parameter exceeds the value configured here it will be set to the value configured here. The default value of 0 disables this limit.",
								MarkdownDescription: "The maximum allowed value of the limit parameter on search requests. If the search request limit parameter exceeds the value configured here it will be set to the value configured here. The default value of 0 disables this limit.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account": schema.StringAttribute{
						Description:         "ServiceAccount defines the service account to use for all tempo components.",
						MarkdownDescription: "ServiceAccount defines the service account to use for all tempo components.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage defines the spec for the object storage endpoint to store traces. User is required to create secret and supply it.",
						MarkdownDescription: "Storage defines the spec for the object storage endpoint to store traces. User is required to create secret and supply it.",
						Attributes: map[string]schema.Attribute{
							"secret": schema.SingleNestedAttribute{
								Description:         "Secret for object storage authentication. Name of a secret in the same namespace as the TempoStack custom resource.",
								MarkdownDescription: "Secret for object storage authentication. Name of a secret in the same namespace as the TempoStack custom resource.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of a secret in the namespace configured for object storage secrets.",
										MarkdownDescription: "Name of a secret in the namespace configured for object storage secrets.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"type": schema.StringAttribute{
										Description:         "Type of object storage that should be used",
										MarkdownDescription: "Type of object storage that should be used",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("azure", "gcs", "s3"),
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS configuration for reaching the object storage endpoint.",
								MarkdownDescription: "TLS configuration for reaching the object storage endpoint.",
								Attributes: map[string]schema.Attribute{
									"ca_name": schema.StringAttribute{
										Description:         "CA is the name of a ConfigMap containing a CA certificate (service-ca.crt). It needs to be in the same namespace as the Tempo custom resource.",
										MarkdownDescription: "CA is the name of a ConfigMap containing a CA certificate (service-ca.crt). It needs to be in the same namespace as the Tempo custom resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cert_name": schema.StringAttribute{
										Description:         "Cert is the name of a Secret containing a certificate (tls.crt) and private key (tls.key). It needs to be in the same namespace as the Tempo custom resource.",
										MarkdownDescription: "Cert is the name of a Secret containing a certificate (tls.crt) and private key (tls.key). It needs to be in the same namespace as the Tempo custom resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if TLS is enabled.",
										MarkdownDescription: "Enabled defines if TLS is enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_version": schema.StringAttribute{
										Description:         "MinVersion defines the minimum acceptable TLS version.",
										MarkdownDescription: "MinVersion defines the minimum acceptable TLS version.",
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

					"storage_class_name": schema.StringAttribute{
						Description:         "StorageClassName for PVCs used by ingester. Defaults to nil (default storage class in the cluster).",
						MarkdownDescription: "StorageClassName for PVCs used by ingester. Defaults to nil (default storage class in the cluster).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_size": schema.StringAttribute{
						Description:         "StorageSize for PVCs used by ingester. Defaults to 10Gi.",
						MarkdownDescription: "StorageSize for PVCs used by ingester. Defaults to 10Gi.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template defines requirements for a set of tempo components.",
						MarkdownDescription: "Template defines requirements for a set of tempo components.",
						Attributes: map[string]schema.Attribute{
							"compactor": schema.SingleNestedAttribute{
								Description:         "Compactor defines the tempo compactor component spec.",
								MarkdownDescription: "Compactor defines the tempo compactor component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the simple form of the node-selection constraint.",
										MarkdownDescription: "NodeSelector defines the simple form of the node-selection constraint.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replicas to be created for this component.",
										MarkdownDescription: "Replicas defines the number of replicas to be created for this component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
										MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines component-specific pod tolerations.",
										MarkdownDescription: "Tolerations defines component-specific pod tolerations.",
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

							"distributor": schema.SingleNestedAttribute{
								Description:         "Distributor defines the distributor component spec.",
								MarkdownDescription: "Distributor defines the distributor component spec.",
								Attributes: map[string]schema.Attribute{
									"component": schema.SingleNestedAttribute{
										Description:         "TempoComponentSpec is embedded to extend this definition with further options.  Currently, there is no way to inline this field. See: https://github.com/golang/go/issues/6213",
										MarkdownDescription: "TempoComponentSpec is embedded to extend this definition with further options.  Currently, there is no way to inline this field. See: https://github.com/golang/go/issues/6213",
										Attributes: map[string]schema.Attribute{
											"node_selector": schema.MapAttribute{
												Description:         "NodeSelector defines the simple form of the node-selection constraint.",
												MarkdownDescription: "NodeSelector defines the simple form of the node-selection constraint.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replicas": schema.Int64Attribute{
												Description:         "Replicas defines the number of replicas to be created for this component.",
												MarkdownDescription: "Replicas defines the number of replicas to be created for this component.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
												MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
												Attributes: map[string]schema.Attribute{
													"claims": schema.ListNestedAttribute{
														Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																	MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

													"limits": schema.MapAttribute{
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

											"tolerations": schema.ListNestedAttribute{
												Description:         "Tolerations defines component-specific pod tolerations.",
												MarkdownDescription: "Tolerations defines component-specific pod tolerations.",
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

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS defines TLS configuration for distributor receivers",
										MarkdownDescription: "TLS defines TLS configuration for distributor receivers",
										Attributes: map[string]schema.Attribute{
											"ca_name": schema.StringAttribute{
												Description:         "CA is the name of a ConfigMap containing a CA certificate (service-ca.crt). It needs to be in the same namespace as the Tempo custom resource.",
												MarkdownDescription: "CA is the name of a ConfigMap containing a CA certificate (service-ca.crt). It needs to be in the same namespace as the Tempo custom resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_name": schema.StringAttribute{
												Description:         "Cert is the name of a Secret containing a certificate (tls.crt) and private key (tls.key). It needs to be in the same namespace as the Tempo custom resource.",
												MarkdownDescription: "Cert is the name of a Secret containing a certificate (tls.crt) and private key (tls.key). It needs to be in the same namespace as the Tempo custom resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if TLS is enabled.",
												MarkdownDescription: "Enabled defines if TLS is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"min_version": schema.StringAttribute{
												Description:         "MinVersion defines the minimum acceptable TLS version.",
												MarkdownDescription: "MinVersion defines the minimum acceptable TLS version.",
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

							"gateway": schema.SingleNestedAttribute{
								Description:         "Gateway defines the tempo gateway spec.",
								MarkdownDescription: "Gateway defines the tempo gateway spec.",
								Attributes: map[string]schema.Attribute{
									"component": schema.SingleNestedAttribute{
										Description:         "TempoComponentSpec is embedded to extend this definition with further options.  Currently there is no way to inline this field. See: https://github.com/golang/go/issues/6213",
										MarkdownDescription: "TempoComponentSpec is embedded to extend this definition with further options.  Currently there is no way to inline this field. See: https://github.com/golang/go/issues/6213",
										Attributes: map[string]schema.Attribute{
											"node_selector": schema.MapAttribute{
												Description:         "NodeSelector defines the simple form of the node-selection constraint.",
												MarkdownDescription: "NodeSelector defines the simple form of the node-selection constraint.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replicas": schema.Int64Attribute{
												Description:         "Replicas defines the number of replicas to be created for this component.",
												MarkdownDescription: "Replicas defines the number of replicas to be created for this component.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
												MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
												Attributes: map[string]schema.Attribute{
													"claims": schema.ListNestedAttribute{
														Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																	MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

													"limits": schema.MapAttribute{
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

											"tolerations": schema.ListNestedAttribute{
												Description:         "Tolerations defines component-specific pod tolerations.",
												MarkdownDescription: "Tolerations defines component-specific pod tolerations.",
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

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"ingress": schema.SingleNestedAttribute{
										Description:         "Ingress defines gateway Ingress options.",
										MarkdownDescription: "Ingress defines gateway Ingress options.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations defines the annotations of the Ingress object.",
												MarkdownDescription: "Annotations defines the annotations of the Ingress object.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host": schema.StringAttribute{
												Description:         "Host defines the hostname of the Ingress object.",
												MarkdownDescription: "Host defines the hostname of the Ingress object.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_class_name": schema.StringAttribute{
												Description:         "IngressClassName defines the name of an IngressClass cluster resource. Defines which ingress controller serves this ingress resource.",
												MarkdownDescription: "IngressClassName defines the name of an IngressClass cluster resource. Defines which ingress controller serves this ingress resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"route": schema.SingleNestedAttribute{
												Description:         "Route defines the options for the OpenShift route.",
												MarkdownDescription: "Route defines the options for the OpenShift route.",
												Attributes: map[string]schema.Attribute{
													"termination": schema.StringAttribute{
														Description:         "Termination defines the termination type. The default is 'edge'.",
														MarkdownDescription: "Termination defines the termination type. The default is 'edge'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("insecure", "edge", "passthrough", "reencrypt"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type defines the type of Ingress for the Jaeger Query UI. Currently ingress, route and none are supported.",
												MarkdownDescription: "Type defines the type of Ingress for the Jaeger Query UI. Currently ingress, route and none are supported.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ingress", "route"),
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

							"ingester": schema.SingleNestedAttribute{
								Description:         "Ingester defines the ingester component spec.",
								MarkdownDescription: "Ingester defines the ingester component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the simple form of the node-selection constraint.",
										MarkdownDescription: "NodeSelector defines the simple form of the node-selection constraint.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replicas to be created for this component.",
										MarkdownDescription: "Replicas defines the number of replicas to be created for this component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
										MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines component-specific pod tolerations.",
										MarkdownDescription: "Tolerations defines component-specific pod tolerations.",
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

							"querier": schema.SingleNestedAttribute{
								Description:         "Querier defines the querier component spec.",
								MarkdownDescription: "Querier defines the querier component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the simple form of the node-selection constraint.",
										MarkdownDescription: "NodeSelector defines the simple form of the node-selection constraint.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replicas to be created for this component.",
										MarkdownDescription: "Replicas defines the number of replicas to be created for this component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
										MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines component-specific pod tolerations.",
										MarkdownDescription: "Tolerations defines component-specific pod tolerations.",
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

							"query_frontend": schema.SingleNestedAttribute{
								Description:         "TempoQueryFrontendSpec defines the query frontend spec.",
								MarkdownDescription: "TempoQueryFrontendSpec defines the query frontend spec.",
								Attributes: map[string]schema.Attribute{
									"component": schema.SingleNestedAttribute{
										Description:         "TempoComponentSpec is embedded to extend this definition with further options.  Currently there is no way to inline this field. See: https://github.com/golang/go/issues/6213",
										MarkdownDescription: "TempoComponentSpec is embedded to extend this definition with further options.  Currently there is no way to inline this field. See: https://github.com/golang/go/issues/6213",
										Attributes: map[string]schema.Attribute{
											"node_selector": schema.MapAttribute{
												Description:         "NodeSelector defines the simple form of the node-selection constraint.",
												MarkdownDescription: "NodeSelector defines the simple form of the node-selection constraint.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replicas": schema.Int64Attribute{
												Description:         "Replicas defines the number of replicas to be created for this component.",
												MarkdownDescription: "Replicas defines the number of replicas to be created for this component.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
												MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
												Attributes: map[string]schema.Attribute{
													"claims": schema.ListNestedAttribute{
														Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																	MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

													"limits": schema.MapAttribute{
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

											"tolerations": schema.ListNestedAttribute{
												Description:         "Tolerations defines component-specific pod tolerations.",
												MarkdownDescription: "Tolerations defines component-specific pod tolerations.",
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

									"jaeger_query": schema.SingleNestedAttribute{
										Description:         "JaegerQuery defines options specific to the Jaeger Query component.",
										MarkdownDescription: "JaegerQuery defines options specific to the Jaeger Query component.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "Oauth defines the options for the oauth proxy used to protect jaeger UI",
												MarkdownDescription: "Oauth defines the options for the oauth proxy used to protect jaeger UI",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "Defines if the authentication will be enabled for jaeger UI.",
														MarkdownDescription: "Defines if the authentication will be enabled for jaeger UI.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sar": schema.StringAttribute{
														Description:         "SAR defines the SAR to be used in the oauth-proxy default is '{'namespace': '<tempo_stack_namespace>', 'resource': 'pods', 'verb': 'get'}",
														MarkdownDescription: "SAR defines the SAR to be used in the oauth-proxy default is '{'namespace': '<tempo_stack_namespace>', 'resource': 'pods', 'verb': 'get'}",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled defines if the Jaeger Query component should be created.",
												MarkdownDescription: "Enabled defines if the Jaeger Query component should be created.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress": schema.SingleNestedAttribute{
												Description:         "Ingress defines the options for the Jaeger Query ingress.",
												MarkdownDescription: "Ingress defines the options for the Jaeger Query ingress.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations defines the annotations of the Ingress object.",
														MarkdownDescription: "Annotations defines the annotations of the Ingress object.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host": schema.StringAttribute{
														Description:         "Host defines the hostname of the Ingress object.",
														MarkdownDescription: "Host defines the hostname of the Ingress object.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ingress_class_name": schema.StringAttribute{
														Description:         "IngressClassName defines the name of an IngressClass cluster resource. Defines which ingress controller serves this ingress resource.",
														MarkdownDescription: "IngressClassName defines the name of an IngressClass cluster resource. Defines which ingress controller serves this ingress resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"route": schema.SingleNestedAttribute{
														Description:         "Route defines the options for the OpenShift route.",
														MarkdownDescription: "Route defines the options for the OpenShift route.",
														Attributes: map[string]schema.Attribute{
															"termination": schema.StringAttribute{
																Description:         "Termination defines the termination type. The default is 'edge'.",
																MarkdownDescription: "Termination defines the termination type. The default is 'edge'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("insecure", "edge", "passthrough", "reencrypt"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.StringAttribute{
														Description:         "Type defines the type of Ingress for the Jaeger Query UI. Currently ingress, route and none are supported.",
														MarkdownDescription: "Type defines the type of Ingress for the Jaeger Query UI. Currently ingress, route and none are supported.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ingress", "route"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"monitor_tab": schema.SingleNestedAttribute{
												Description:         "MonitorTab defines the monitor tab configuration.",
												MarkdownDescription: "MonitorTab defines the monitor tab configuration.",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "Enabled enables the monitor tab in the Jaeger console. The PrometheusEndpoint must be configured to enable this feature.",
														MarkdownDescription: "Enabled enables the monitor tab in the Jaeger console. The PrometheusEndpoint must be configured to enable this feature.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"prometheus_endpoint": schema.StringAttribute{
														Description:         "PrometheusEndpoint defines the endpoint to the Prometheus instance that contains the span rate, error, and duration (RED) metrics. For instance on OpenShift this is set to https://thanos-querier.openshift-monitoring.svc.cluster.local:9091",
														MarkdownDescription: "PrometheusEndpoint defines the endpoint to the Prometheus instance that contains the span rate, error, and duration (RED) metrics. For instance on OpenShift this is set to https://thanos-querier.openshift-monitoring.svc.cluster.local:9091",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resources defines resources for this component, this will override the calculated resources derived from total",
												MarkdownDescription: "Resources defines resources for this component, this will override the calculated resources derived from total",
												Attributes: map[string]schema.Attribute{
													"claims": schema.ListNestedAttribute{
														Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																	MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

													"limits": schema.MapAttribute{
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

											"services_query_duration": schema.StringAttribute{
												Description:         "ServicesQueryDuration defines how long the services will be available in the services list",
												MarkdownDescription: "ServicesQueryDuration defines how long the services will be available in the services list",
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

					"tenants": schema.SingleNestedAttribute{
						Description:         "Tenants defines the per-tenant authentication and authorization spec.",
						MarkdownDescription: "Tenants defines the per-tenant authentication and authorization spec.",
						Attributes: map[string]schema.Attribute{
							"authentication": schema.ListNestedAttribute{
								Description:         "Authentication defines the tempo-gateway component authentication configuration spec per tenant.",
								MarkdownDescription: "Authentication defines the tempo-gateway component authentication configuration spec per tenant.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"oidc": schema.SingleNestedAttribute{
											Description:         "OIDC defines the spec for the OIDC tenant's authentication.",
											MarkdownDescription: "OIDC defines the spec for the OIDC tenant's authentication.",
											Attributes: map[string]schema.Attribute{
												"group_claim": schema.StringAttribute{
													Description:         "Group claim field from ID Token",
													MarkdownDescription: "Group claim field from ID Token",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"issuer_url": schema.StringAttribute{
													Description:         "IssuerURL defines the URL for issuer.",
													MarkdownDescription: "IssuerURL defines the URL for issuer.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"redirect_url": schema.StringAttribute{
													Description:         "RedirectURL defines the URL for redirect.",
													MarkdownDescription: "RedirectURL defines the URL for redirect.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "Secret defines the spec for the clientID, clientSecret and issuerCAPath for tenant's authentication.",
													MarkdownDescription: "Secret defines the spec for the clientID, clientSecret and issuerCAPath for tenant's authentication.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of a secret in the namespace configured for tenant secrets.",
															MarkdownDescription: "Name of a secret in the namespace configured for tenant secrets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"username_claim": schema.StringAttribute{
													Description:         "User claim field from ID Token",
													MarkdownDescription: "User claim field from ID Token",
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
											Description:         "TenantID defines a universally unique identifier of the tenant. Unlike the tenantName, which must be unique at a given time, the tenantId must be unique over the entire lifetime of the Tempo deployment. Tempo uses this ID to prefix objects in the object storage.",
											MarkdownDescription: "TenantID defines a universally unique identifier of the tenant. Unlike the tenantName, which must be unique at a given time, the tenantId must be unique over the entire lifetime of the Tempo deployment. Tempo uses this ID to prefix objects in the object storage.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"tenant_name": schema.StringAttribute{
											Description:         "TenantName defines a human readable, unique name of the tenant. The value of this field must be specified in the X-Scope-OrgID header and in the resources field of a ClusterRole to identify the tenant.",
											MarkdownDescription: "TenantName defines a human readable, unique name of the tenant. The value of this field must be specified in the X-Scope-OrgID header and in the resources field of a ClusterRole to identify the tenant.",
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

							"authorization": schema.SingleNestedAttribute{
								Description:         "Authorization defines the tempo-gateway component authorization configuration spec per tenant.",
								MarkdownDescription: "Authorization defines the tempo-gateway component authorization configuration spec per tenant.",
								Attributes: map[string]schema.Attribute{
									"role_bindings": schema.ListNestedAttribute{
										Description:         "RoleBindings defines configuration to bind a set of roles to a set of subjects.",
										MarkdownDescription: "RoleBindings defines configuration to bind a set of roles to a set of subjects.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"roles": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"subjects": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"kind": schema.StringAttribute{
																Description:         "SubjectKind is a kind of Tempo Gateway RBAC subject.",
																MarkdownDescription: "SubjectKind is a kind of Tempo Gateway RBAC subject.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("user", "group"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
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

									"roles": schema.ListNestedAttribute{
										Description:         "Roles defines a set of permissions to interact with a tenant.",
										MarkdownDescription: "Roles defines a set of permissions to interact with a tenant.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"permissions": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"resources": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"tenants": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
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

							"mode": schema.StringAttribute{
								Description:         "Mode defines the multitenancy mode.",
								MarkdownDescription: "Mode defines the multitenancy mode.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("static", "openshift"),
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
	}
}

func (r *TempoGrafanaComTempoStackV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tempo_grafana_com_tempo_stack_v1alpha1_manifest")

	var model TempoGrafanaComTempoStackV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tempo.grafana.com/v1alpha1")
	model.Kind = pointer.String("TempoStack")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
