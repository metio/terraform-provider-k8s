/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &LokiGrafanaComLokiStackV1Beta1Manifest{}
)

func NewLokiGrafanaComLokiStackV1Beta1Manifest() datasource.DataSource {
	return &LokiGrafanaComLokiStackV1Beta1Manifest{}
}

type LokiGrafanaComLokiStackV1Beta1Manifest struct{}

type LokiGrafanaComLokiStackV1Beta1ManifestData struct {
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
		Limits *struct {
			Global *struct {
				Ingestion *struct {
					IngestionBurstSize        *int64 `tfsdk:"ingestion_burst_size" json:"ingestionBurstSize,omitempty"`
					IngestionRate             *int64 `tfsdk:"ingestion_rate" json:"ingestionRate,omitempty"`
					MaxGlobalStreamsPerTenant *int64 `tfsdk:"max_global_streams_per_tenant" json:"maxGlobalStreamsPerTenant,omitempty"`
					MaxLabelNameLength        *int64 `tfsdk:"max_label_name_length" json:"maxLabelNameLength,omitempty"`
					MaxLabelNamesPerSeries    *int64 `tfsdk:"max_label_names_per_series" json:"maxLabelNamesPerSeries,omitempty"`
					MaxLabelValueLength       *int64 `tfsdk:"max_label_value_length" json:"maxLabelValueLength,omitempty"`
					MaxLineSize               *int64 `tfsdk:"max_line_size" json:"maxLineSize,omitempty"`
				} `tfsdk:"ingestion" json:"ingestion,omitempty"`
				Queries *struct {
					MaxChunksPerQuery       *int64 `tfsdk:"max_chunks_per_query" json:"maxChunksPerQuery,omitempty"`
					MaxEntriesLimitPerQuery *int64 `tfsdk:"max_entries_limit_per_query" json:"maxEntriesLimitPerQuery,omitempty"`
					MaxQuerySeries          *int64 `tfsdk:"max_query_series" json:"maxQuerySeries,omitempty"`
				} `tfsdk:"queries" json:"queries,omitempty"`
			} `tfsdk:"global" json:"global,omitempty"`
			Tenants *struct {
				Ingestion *struct {
					IngestionBurstSize        *int64 `tfsdk:"ingestion_burst_size" json:"ingestionBurstSize,omitempty"`
					IngestionRate             *int64 `tfsdk:"ingestion_rate" json:"ingestionRate,omitempty"`
					MaxGlobalStreamsPerTenant *int64 `tfsdk:"max_global_streams_per_tenant" json:"maxGlobalStreamsPerTenant,omitempty"`
					MaxLabelNameLength        *int64 `tfsdk:"max_label_name_length" json:"maxLabelNameLength,omitempty"`
					MaxLabelNamesPerSeries    *int64 `tfsdk:"max_label_names_per_series" json:"maxLabelNamesPerSeries,omitempty"`
					MaxLabelValueLength       *int64 `tfsdk:"max_label_value_length" json:"maxLabelValueLength,omitempty"`
					MaxLineSize               *int64 `tfsdk:"max_line_size" json:"maxLineSize,omitempty"`
				} `tfsdk:"ingestion" json:"ingestion,omitempty"`
				Queries *struct {
					MaxChunksPerQuery       *int64 `tfsdk:"max_chunks_per_query" json:"maxChunksPerQuery,omitempty"`
					MaxEntriesLimitPerQuery *int64 `tfsdk:"max_entries_limit_per_query" json:"maxEntriesLimitPerQuery,omitempty"`
					MaxQuerySeries          *int64 `tfsdk:"max_query_series" json:"maxQuerySeries,omitempty"`
				} `tfsdk:"queries" json:"queries,omitempty"`
			} `tfsdk:"tenants" json:"tenants,omitempty"`
		} `tfsdk:"limits" json:"limits,omitempty"`
		ManagementState   *string `tfsdk:"management_state" json:"managementState,omitempty"`
		ReplicationFactor *int64  `tfsdk:"replication_factor" json:"replicationFactor,omitempty"`
		Rules             *struct {
			Enabled           *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		Size    *string `tfsdk:"size" json:"size,omitempty"`
		Storage *struct {
			Schemas *[]struct {
				EffectiveDate *string `tfsdk:"effective_date" json:"effectiveDate,omitempty"`
				Version       *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"schemas" json:"schemas,omitempty"`
			Secret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			Tls *struct {
				CaName *string `tfsdk:"ca_name" json:"caName,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		Template         *struct {
			Compactor *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"compactor" json:"compactor,omitempty"`
			Distributor *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"distributor" json:"distributor,omitempty"`
			Gateway *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"gateway" json:"gateway,omitempty"`
			IndexGateway *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"index_gateway" json:"indexGateway,omitempty"`
			Ingester *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
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
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"querier" json:"querier,omitempty"`
			QueryFrontend *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"query_frontend" json:"queryFrontend,omitempty"`
			Ruler *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"ruler" json:"ruler,omitempty"`
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
				Opa *struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"opa" json:"opa,omitempty"`
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

func (r *LokiGrafanaComLokiStackV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_loki_grafana_com_loki_stack_v1beta1_manifest"
}

func (r *LokiGrafanaComLokiStackV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LokiStack is the Schema for the lokistacks API",
		MarkdownDescription: "LokiStack is the Schema for the lokistacks API",
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
				Description:         "LokiStackSpec defines the desired state of LokiStack",
				MarkdownDescription: "LokiStackSpec defines the desired state of LokiStack",
				Attributes: map[string]schema.Attribute{
					"limits": schema.SingleNestedAttribute{
						Description:         "Limits defines the per-tenant limits to be applied to log stream processing and the per-tenant the config overrides.",
						MarkdownDescription: "Limits defines the per-tenant limits to be applied to log stream processing and the per-tenant the config overrides.",
						Attributes: map[string]schema.Attribute{
							"global": schema.SingleNestedAttribute{
								Description:         "Global defines the limits applied globally across the cluster.",
								MarkdownDescription: "Global defines the limits applied globally across the cluster.",
								Attributes: map[string]schema.Attribute{
									"ingestion": schema.SingleNestedAttribute{
										Description:         "IngestionLimits defines the limits applied on ingested log streams.",
										MarkdownDescription: "IngestionLimits defines the limits applied on ingested log streams.",
										Attributes: map[string]schema.Attribute{
											"ingestion_burst_size": schema.Int64Attribute{
												Description:         "IngestionBurstSize defines the local rate-limited sample size perdistributor replica. It should be set to the set at least to themaximum logs size expected in a single push request.",
												MarkdownDescription: "IngestionBurstSize defines the local rate-limited sample size perdistributor replica. It should be set to the set at least to themaximum logs size expected in a single push request.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingestion_rate": schema.Int64Attribute{
												Description:         "IngestionRate defines the sample size per second. Units MB.",
												MarkdownDescription: "IngestionRate defines the sample size per second. Units MB.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_global_streams_per_tenant": schema.Int64Attribute{
												Description:         "MaxGlobalStreamsPerTenant defines the maximum number of active streamsper tenant, across the cluster.",
												MarkdownDescription: "MaxGlobalStreamsPerTenant defines the maximum number of active streamsper tenant, across the cluster.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_label_name_length": schema.Int64Attribute{
												Description:         "MaxLabelNameLength defines the maximum number of characters allowedfor label keys in log streams.",
												MarkdownDescription: "MaxLabelNameLength defines the maximum number of characters allowedfor label keys in log streams.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_label_names_per_series": schema.Int64Attribute{
												Description:         "MaxLabelNamesPerSeries defines the maximum number of label names per seriesin each log stream.",
												MarkdownDescription: "MaxLabelNamesPerSeries defines the maximum number of label names per seriesin each log stream.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_label_value_length": schema.Int64Attribute{
												Description:         "MaxLabelValueLength defines the maximum number of characters allowedfor label values in log streams.",
												MarkdownDescription: "MaxLabelValueLength defines the maximum number of characters allowedfor label values in log streams.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_line_size": schema.Int64Attribute{
												Description:         "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",
												MarkdownDescription: "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"queries": schema.SingleNestedAttribute{
										Description:         "QueryLimits defines the limit applied on querying log streams.",
										MarkdownDescription: "QueryLimits defines the limit applied on querying log streams.",
										Attributes: map[string]schema.Attribute{
											"max_chunks_per_query": schema.Int64Attribute{
												Description:         "MaxChunksPerQuery defines the maximum number of chunksthat can be fetched by a single query.",
												MarkdownDescription: "MaxChunksPerQuery defines the maximum number of chunksthat can be fetched by a single query.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_entries_limit_per_query": schema.Int64Attribute{
												Description:         "MaxEntriesLimitsPerQuery defines the maximum number of log entriesthat will be returned for a query.",
												MarkdownDescription: "MaxEntriesLimitsPerQuery defines the maximum number of log entriesthat will be returned for a query.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_query_series": schema.Int64Attribute{
												Description:         "MaxQuerySeries defines the maximum of unique seriesthat is returned by a metric query.",
												MarkdownDescription: "MaxQuerySeries defines the maximum of unique seriesthat is returned by a metric query.",
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

							"tenants": schema.SingleNestedAttribute{
								Description:         "Tenants defines the limits and overrides applied per tenant.",
								MarkdownDescription: "Tenants defines the limits and overrides applied per tenant.",
								Attributes: map[string]schema.Attribute{
									"ingestion": schema.SingleNestedAttribute{
										Description:         "IngestionLimits defines the limits applied on ingested log streams.",
										MarkdownDescription: "IngestionLimits defines the limits applied on ingested log streams.",
										Attributes: map[string]schema.Attribute{
											"ingestion_burst_size": schema.Int64Attribute{
												Description:         "IngestionBurstSize defines the local rate-limited sample size perdistributor replica. It should be set to the set at least to themaximum logs size expected in a single push request.",
												MarkdownDescription: "IngestionBurstSize defines the local rate-limited sample size perdistributor replica. It should be set to the set at least to themaximum logs size expected in a single push request.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingestion_rate": schema.Int64Attribute{
												Description:         "IngestionRate defines the sample size per second. Units MB.",
												MarkdownDescription: "IngestionRate defines the sample size per second. Units MB.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_global_streams_per_tenant": schema.Int64Attribute{
												Description:         "MaxGlobalStreamsPerTenant defines the maximum number of active streamsper tenant, across the cluster.",
												MarkdownDescription: "MaxGlobalStreamsPerTenant defines the maximum number of active streamsper tenant, across the cluster.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_label_name_length": schema.Int64Attribute{
												Description:         "MaxLabelNameLength defines the maximum number of characters allowedfor label keys in log streams.",
												MarkdownDescription: "MaxLabelNameLength defines the maximum number of characters allowedfor label keys in log streams.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_label_names_per_series": schema.Int64Attribute{
												Description:         "MaxLabelNamesPerSeries defines the maximum number of label names per seriesin each log stream.",
												MarkdownDescription: "MaxLabelNamesPerSeries defines the maximum number of label names per seriesin each log stream.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_label_value_length": schema.Int64Attribute{
												Description:         "MaxLabelValueLength defines the maximum number of characters allowedfor label values in log streams.",
												MarkdownDescription: "MaxLabelValueLength defines the maximum number of characters allowedfor label values in log streams.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_line_size": schema.Int64Attribute{
												Description:         "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",
												MarkdownDescription: "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"queries": schema.SingleNestedAttribute{
										Description:         "QueryLimits defines the limit applied on querying log streams.",
										MarkdownDescription: "QueryLimits defines the limit applied on querying log streams.",
										Attributes: map[string]schema.Attribute{
											"max_chunks_per_query": schema.Int64Attribute{
												Description:         "MaxChunksPerQuery defines the maximum number of chunksthat can be fetched by a single query.",
												MarkdownDescription: "MaxChunksPerQuery defines the maximum number of chunksthat can be fetched by a single query.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_entries_limit_per_query": schema.Int64Attribute{
												Description:         "MaxEntriesLimitsPerQuery defines the maximum number of log entriesthat will be returned for a query.",
												MarkdownDescription: "MaxEntriesLimitsPerQuery defines the maximum number of log entriesthat will be returned for a query.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_query_series": schema.Int64Attribute{
												Description:         "MaxQuerySeries defines the maximum of unique seriesthat is returned by a metric query.",
												MarkdownDescription: "MaxQuerySeries defines the maximum of unique seriesthat is returned by a metric query.",
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
						Description:         "ManagementState defines if the CR should be managed by the operator or not.Default is managed.",
						MarkdownDescription: "ManagementState defines if the CR should be managed by the operator or not.Default is managed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Managed", "Unmanaged"),
						},
					},

					"replication_factor": schema.Int64Attribute{
						Description:         "ReplicationFactor defines the policy for log stream replication.",
						MarkdownDescription: "ReplicationFactor defines the policy for log stream replication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"rules": schema.SingleNestedAttribute{
						Description:         "Rules defines the spec for the ruler component",
						MarkdownDescription: "Rules defines the spec for the ruler component",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines a flag to enable/disable the ruler component",
								MarkdownDescription: "Enabled defines a flag to enable/disable the ruler component",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace_selector": schema.SingleNestedAttribute{
								Description:         "Namespaces to be selected for PrometheusRules discovery. If unspecified, onlythe same namespace as the LokiStack object is in is used.",
								MarkdownDescription: "Namespaces to be selected for PrometheusRules discovery. If unspecified, onlythe same namespace as the LokiStack object is in is used.",
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
													Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"selector": schema.SingleNestedAttribute{
								Description:         "A selector to select which LokiRules to mount for loading alerting/recordingrules from.",
								MarkdownDescription: "A selector to select which LokiRules to mount for loading alerting/recordingrules from.",
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
													Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"size": schema.StringAttribute{
						Description:         "Size defines one of the support Loki deployment scale out sizes.",
						MarkdownDescription: "Size defines one of the support Loki deployment scale out sizes.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("1x.extra-small", "1x.small", "1x.medium"),
						},
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage defines the spec for the object storage endpoint to store logs.",
						MarkdownDescription: "Storage defines the spec for the object storage endpoint to store logs.",
						Attributes: map[string]schema.Attribute{
							"schemas": schema.ListNestedAttribute{
								Description:         "Schemas for reading and writing logs.",
								MarkdownDescription: "Schemas for reading and writing logs.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effective_date": schema.StringAttribute{
											Description:         "EffectiveDate is the date in UTC that the schema will be applied on.To ensure readibility of logs, this date should be before the currentdate in UTC.",
											MarkdownDescription: "EffectiveDate is the date in UTC that the schema will be applied on.To ensure readibility of logs, this date should be before the currentdate in UTC.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{4,})([-]([0-9]{2})){2}$`), ""),
											},
										},

										"version": schema.StringAttribute{
											Description:         "Version for writing and reading logs.",
											MarkdownDescription: "Version for writing and reading logs.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("v11", "v12"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": schema.SingleNestedAttribute{
								Description:         "Secret for object storage authentication.Name of a secret in the same namespace as the LokiStack custom resource.",
								MarkdownDescription: "Secret for object storage authentication.Name of a secret in the same namespace as the LokiStack custom resource.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of a secret in the namespace configured for object storage secrets.",
										MarkdownDescription: "Name of a secret in the namespace configured for object storage secrets.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type of object storage that should be used",
										MarkdownDescription: "Type of object storage that should be used",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("azure", "gcs", "s3", "swift"),
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
										Description:         "CA is the name of a ConfigMap containing a CA certificate.It needs to be in the same namespace as the LokiStack custom resource.",
										MarkdownDescription: "CA is the name of a ConfigMap containing a CA certificate.It needs to be in the same namespace as the LokiStack custom resource.",
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
						Description:         "Storage class name defines the storage class for ingester/querier PVCs.",
						MarkdownDescription: "Storage class name defines the storage class for ingester/querier PVCs.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template defines the resource/limits/tolerations/nodeselectors per component",
						MarkdownDescription: "Template defines the resource/limits/tolerations/nodeselectors per component",
						Attributes: map[string]schema.Attribute{
							"compactor": schema.SingleNestedAttribute{
								Description:         "Compactor defines the compaction component spec.",
								MarkdownDescription: "Compactor defines the compaction component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"gateway": schema.SingleNestedAttribute{
								Description:         "Gateway defines the lokistack gateway component spec.",
								MarkdownDescription: "Gateway defines the lokistack gateway component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"index_gateway": schema.SingleNestedAttribute{
								Description:         "IndexGateway defines the index gateway component spec.",
								MarkdownDescription: "IndexGateway defines the index gateway component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"ingester": schema.SingleNestedAttribute{
								Description:         "Ingester defines the ingester component spec.",
								MarkdownDescription: "Ingester defines the ingester component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
								Description:         "QueryFrontend defines the query frontend component spec.",
								MarkdownDescription: "QueryFrontend defines the query frontend component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"ruler": schema.SingleNestedAttribute{
								Description:         "Ruler defines the ruler component spec.",
								MarkdownDescription: "Ruler defines the ruler component spec.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedulethe component onto it.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedulethe component onto it.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"tenants": schema.SingleNestedAttribute{
						Description:         "Tenants defines the per-tenant authentication and authorization spec for the lokistack-gateway component.",
						MarkdownDescription: "Tenants defines the per-tenant authentication and authorization spec for the lokistack-gateway component.",
						Attributes: map[string]schema.Attribute{
							"authentication": schema.ListNestedAttribute{
								Description:         "Authentication defines the lokistack-gateway component authentication configuration spec per tenant.",
								MarkdownDescription: "Authentication defines the lokistack-gateway component authentication configuration spec per tenant.",
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
													Required:            true,
													Optional:            false,
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
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"tenant_id": schema.StringAttribute{
											Description:         "TenantID defines the id of the tenant.",
											MarkdownDescription: "TenantID defines the id of the tenant.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"tenant_name": schema.StringAttribute{
											Description:         "TenantName defines the name of the tenant.",
											MarkdownDescription: "TenantName defines the name of the tenant.",
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
								Description:         "Authorization defines the lokistack-gateway component authorization configuration spec per tenant.",
								MarkdownDescription: "Authorization defines the lokistack-gateway component authorization configuration spec per tenant.",
								Attributes: map[string]schema.Attribute{
									"opa": schema.SingleNestedAttribute{
										Description:         "OPA defines the spec for the third-party endpoint for tenant's authorization.",
										MarkdownDescription: "OPA defines the spec for the third-party endpoint for tenant's authorization.",
										Attributes: map[string]schema.Attribute{
											"url": schema.StringAttribute{
												Description:         "URL defines the third-party endpoint for authorization.",
												MarkdownDescription: "URL defines the third-party endpoint for authorization.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

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
																Description:         "SubjectKind is a kind of LokiStack Gateway RBAC subject.",
																MarkdownDescription: "SubjectKind is a kind of LokiStack Gateway RBAC subject.",
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
								Description:         "Mode defines the mode in which lokistack-gateway component will be configured.",
								MarkdownDescription: "Mode defines the mode in which lokistack-gateway component will be configured.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("static", "dynamic", "openshift-logging"),
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

func (r *LokiGrafanaComLokiStackV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_loki_stack_v1beta1_manifest")

	var model LokiGrafanaComLokiStackV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("loki.grafana.com/v1beta1")
	model.Kind = pointer.String("LokiStack")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
