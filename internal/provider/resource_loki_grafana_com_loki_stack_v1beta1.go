/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type LokiGrafanaComLokiStackV1Beta1Resource struct{}

var (
	_ resource.Resource = (*LokiGrafanaComLokiStackV1Beta1Resource)(nil)
)

type LokiGrafanaComLokiStackV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LokiGrafanaComLokiStackV1Beta1GoModel struct {
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
		Limits *struct {
			Global *struct {
				Ingestion *struct {
					IngestionBurstSize *int64 `tfsdk:"ingestion_burst_size" yaml:"ingestionBurstSize,omitempty"`

					IngestionRate *int64 `tfsdk:"ingestion_rate" yaml:"ingestionRate,omitempty"`

					MaxGlobalStreamsPerTenant *int64 `tfsdk:"max_global_streams_per_tenant" yaml:"maxGlobalStreamsPerTenant,omitempty"`

					MaxLabelNameLength *int64 `tfsdk:"max_label_name_length" yaml:"maxLabelNameLength,omitempty"`

					MaxLabelNamesPerSeries *int64 `tfsdk:"max_label_names_per_series" yaml:"maxLabelNamesPerSeries,omitempty"`

					MaxLabelValueLength *int64 `tfsdk:"max_label_value_length" yaml:"maxLabelValueLength,omitempty"`

					MaxLineSize *int64 `tfsdk:"max_line_size" yaml:"maxLineSize,omitempty"`
				} `tfsdk:"ingestion" yaml:"ingestion,omitempty"`

				Queries *struct {
					MaxChunksPerQuery *int64 `tfsdk:"max_chunks_per_query" yaml:"maxChunksPerQuery,omitempty"`

					MaxEntriesLimitPerQuery *int64 `tfsdk:"max_entries_limit_per_query" yaml:"maxEntriesLimitPerQuery,omitempty"`

					MaxQuerySeries *int64 `tfsdk:"max_query_series" yaml:"maxQuerySeries,omitempty"`
				} `tfsdk:"queries" yaml:"queries,omitempty"`
			} `tfsdk:"global" yaml:"global,omitempty"`

			Tenants *struct {
				Ingestion *struct {
					IngestionBurstSize *int64 `tfsdk:"ingestion_burst_size" yaml:"ingestionBurstSize,omitempty"`

					IngestionRate *int64 `tfsdk:"ingestion_rate" yaml:"ingestionRate,omitempty"`

					MaxGlobalStreamsPerTenant *int64 `tfsdk:"max_global_streams_per_tenant" yaml:"maxGlobalStreamsPerTenant,omitempty"`

					MaxLabelNameLength *int64 `tfsdk:"max_label_name_length" yaml:"maxLabelNameLength,omitempty"`

					MaxLabelNamesPerSeries *int64 `tfsdk:"max_label_names_per_series" yaml:"maxLabelNamesPerSeries,omitempty"`

					MaxLabelValueLength *int64 `tfsdk:"max_label_value_length" yaml:"maxLabelValueLength,omitempty"`

					MaxLineSize *int64 `tfsdk:"max_line_size" yaml:"maxLineSize,omitempty"`
				} `tfsdk:"ingestion" yaml:"ingestion,omitempty"`

				Queries *struct {
					MaxChunksPerQuery *int64 `tfsdk:"max_chunks_per_query" yaml:"maxChunksPerQuery,omitempty"`

					MaxEntriesLimitPerQuery *int64 `tfsdk:"max_entries_limit_per_query" yaml:"maxEntriesLimitPerQuery,omitempty"`

					MaxQuerySeries *int64 `tfsdk:"max_query_series" yaml:"maxQuerySeries,omitempty"`
				} `tfsdk:"queries" yaml:"queries,omitempty"`
			} `tfsdk:"tenants" yaml:"tenants,omitempty"`
		} `tfsdk:"limits" yaml:"limits,omitempty"`

		ManagementState *string `tfsdk:"management_state" yaml:"managementState,omitempty"`

		ReplicationFactor *int64 `tfsdk:"replication_factor" yaml:"replicationFactor,omitempty"`

		Rules *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

			Selector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`

		Size *string `tfsdk:"size" yaml:"size,omitempty"`

		Storage *struct {
			Schemas *[]struct {
				EffectiveDate *string `tfsdk:"effective_date" yaml:"effectiveDate,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"schemas" yaml:"schemas,omitempty"`

			Secret *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"secret" yaml:"secret,omitempty"`

			Tls *struct {
				CaName *string `tfsdk:"ca_name" yaml:"caName,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"storage" yaml:"storage,omitempty"`

		StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

		Template *struct {
			Compactor *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"compactor" yaml:"compactor,omitempty"`

			Distributor *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"distributor" yaml:"distributor,omitempty"`

			Gateway *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"gateway" yaml:"gateway,omitempty"`

			IndexGateway *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"index_gateway" yaml:"indexGateway,omitempty"`

			Ingester *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"ingester" yaml:"ingester,omitempty"`

			Querier *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"querier" yaml:"querier,omitempty"`

			QueryFrontend *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"query_frontend" yaml:"queryFrontend,omitempty"`

			Ruler *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
			} `tfsdk:"ruler" yaml:"ruler,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`

		Tenants *struct {
			Authentication *[]struct {
				Oidc *struct {
					GroupClaim *string `tfsdk:"group_claim" yaml:"groupClaim,omitempty"`

					IssuerURL *string `tfsdk:"issuer_url" yaml:"issuerURL,omitempty"`

					RedirectURL *string `tfsdk:"redirect_url" yaml:"redirectURL,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					UsernameClaim *string `tfsdk:"username_claim" yaml:"usernameClaim,omitempty"`
				} `tfsdk:"oidc" yaml:"oidc,omitempty"`

				TenantId *string `tfsdk:"tenant_id" yaml:"tenantId,omitempty"`

				TenantName *string `tfsdk:"tenant_name" yaml:"tenantName,omitempty"`
			} `tfsdk:"authentication" yaml:"authentication,omitempty"`

			Authorization *struct {
				Opa *struct {
					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"opa" yaml:"opa,omitempty"`

				RoleBindings *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

					Subjects *[]struct {
						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"subjects" yaml:"subjects,omitempty"`
				} `tfsdk:"role_bindings" yaml:"roleBindings,omitempty"`

				Roles *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Permissions *[]string `tfsdk:"permissions" yaml:"permissions,omitempty"`

					Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

					Tenants *[]string `tfsdk:"tenants" yaml:"tenants,omitempty"`
				} `tfsdk:"roles" yaml:"roles,omitempty"`
			} `tfsdk:"authorization" yaml:"authorization,omitempty"`

			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
		} `tfsdk:"tenants" yaml:"tenants,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLokiGrafanaComLokiStackV1Beta1Resource() resource.Resource {
	return &LokiGrafanaComLokiStackV1Beta1Resource{}
}

func (r *LokiGrafanaComLokiStackV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loki_grafana_com_loki_stack_v1beta1"
}

func (r *LokiGrafanaComLokiStackV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "LokiStack is the Schema for the lokistacks API",
		MarkdownDescription: "LokiStack is the Schema for the lokistacks API",
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
				Description:         "LokiStackSpec defines the desired state of LokiStack",
				MarkdownDescription: "LokiStackSpec defines the desired state of LokiStack",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"limits": {
						Description:         "Limits defines the limits to be applied to log stream processing.",
						MarkdownDescription: "Limits defines the limits to be applied to log stream processing.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"global": {
								Description:         "Global defines the limits applied globally across the cluster.",
								MarkdownDescription: "Global defines the limits applied globally across the cluster.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ingestion": {
										Description:         "IngestionLimits defines the limits applied on ingested log streams.",
										MarkdownDescription: "IngestionLimits defines the limits applied on ingested log streams.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ingestion_burst_size": {
												Description:         "IngestionBurstSize defines the local rate-limited sample size per distributor replica. It should be set to the set at least to the maximum logs size expected in a single push request.",
												MarkdownDescription: "IngestionBurstSize defines the local rate-limited sample size per distributor replica. It should be set to the set at least to the maximum logs size expected in a single push request.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ingestion_rate": {
												Description:         "IngestionRate defines the sample size per second. Units MB.",
												MarkdownDescription: "IngestionRate defines the sample size per second. Units MB.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_global_streams_per_tenant": {
												Description:         "MaxGlobalStreamsPerTenant defines the maximum number of active streams per tenant, across the cluster.",
												MarkdownDescription: "MaxGlobalStreamsPerTenant defines the maximum number of active streams per tenant, across the cluster.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_label_name_length": {
												Description:         "MaxLabelNameLength defines the maximum number of characters allowed for label keys in log streams.",
												MarkdownDescription: "MaxLabelNameLength defines the maximum number of characters allowed for label keys in log streams.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_label_names_per_series": {
												Description:         "MaxLabelNamesPerSeries defines the maximum number of label names per series in each log stream.",
												MarkdownDescription: "MaxLabelNamesPerSeries defines the maximum number of label names per series in each log stream.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_label_value_length": {
												Description:         "MaxLabelValueLength defines the maximum number of characters allowed for label values in log streams.",
												MarkdownDescription: "MaxLabelValueLength defines the maximum number of characters allowed for label values in log streams.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_line_size": {
												Description:         "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",
												MarkdownDescription: "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"queries": {
										Description:         "QueryLimits defines the limit applied on querying log streams.",
										MarkdownDescription: "QueryLimits defines the limit applied on querying log streams.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_chunks_per_query": {
												Description:         "MaxChunksPerQuery defines the maximum number of chunks that can be fetched by a single query.",
												MarkdownDescription: "MaxChunksPerQuery defines the maximum number of chunks that can be fetched by a single query.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_entries_limit_per_query": {
												Description:         "MaxEntriesLimitsPerQuery defines the maximum number of log entries that will be returned for a query.",
												MarkdownDescription: "MaxEntriesLimitsPerQuery defines the maximum number of log entries that will be returned for a query.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_query_series": {
												Description:         "MaxQuerySeries defines the the maximum of unique series that is returned by a metric query.",
												MarkdownDescription: "MaxQuerySeries defines the the maximum of unique series that is returned by a metric query.",

												Type: types.Int64Type,

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

							"tenants": {
								Description:         "Tenants defines the limits applied per tenant.",
								MarkdownDescription: "Tenants defines the limits applied per tenant.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ingestion": {
										Description:         "IngestionLimits defines the limits applied on ingested log streams.",
										MarkdownDescription: "IngestionLimits defines the limits applied on ingested log streams.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ingestion_burst_size": {
												Description:         "IngestionBurstSize defines the local rate-limited sample size per distributor replica. It should be set to the set at least to the maximum logs size expected in a single push request.",
												MarkdownDescription: "IngestionBurstSize defines the local rate-limited sample size per distributor replica. It should be set to the set at least to the maximum logs size expected in a single push request.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ingestion_rate": {
												Description:         "IngestionRate defines the sample size per second. Units MB.",
												MarkdownDescription: "IngestionRate defines the sample size per second. Units MB.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_global_streams_per_tenant": {
												Description:         "MaxGlobalStreamsPerTenant defines the maximum number of active streams per tenant, across the cluster.",
												MarkdownDescription: "MaxGlobalStreamsPerTenant defines the maximum number of active streams per tenant, across the cluster.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_label_name_length": {
												Description:         "MaxLabelNameLength defines the maximum number of characters allowed for label keys in log streams.",
												MarkdownDescription: "MaxLabelNameLength defines the maximum number of characters allowed for label keys in log streams.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_label_names_per_series": {
												Description:         "MaxLabelNamesPerSeries defines the maximum number of label names per series in each log stream.",
												MarkdownDescription: "MaxLabelNamesPerSeries defines the maximum number of label names per series in each log stream.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_label_value_length": {
												Description:         "MaxLabelValueLength defines the maximum number of characters allowed for label values in log streams.",
												MarkdownDescription: "MaxLabelValueLength defines the maximum number of characters allowed for label values in log streams.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_line_size": {
												Description:         "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",
												MarkdownDescription: "MaxLineSize defines the maximum line size on ingestion path. Units in Bytes.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"queries": {
										Description:         "QueryLimits defines the limit applied on querying log streams.",
										MarkdownDescription: "QueryLimits defines the limit applied on querying log streams.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_chunks_per_query": {
												Description:         "MaxChunksPerQuery defines the maximum number of chunks that can be fetched by a single query.",
												MarkdownDescription: "MaxChunksPerQuery defines the maximum number of chunks that can be fetched by a single query.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_entries_limit_per_query": {
												Description:         "MaxEntriesLimitsPerQuery defines the maximum number of log entries that will be returned for a query.",
												MarkdownDescription: "MaxEntriesLimitsPerQuery defines the maximum number of log entries that will be returned for a query.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_query_series": {
												Description:         "MaxQuerySeries defines the the maximum of unique series that is returned by a metric query.",
												MarkdownDescription: "MaxQuerySeries defines the the maximum of unique series that is returned by a metric query.",

												Type: types.Int64Type,

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

					"management_state": {
						Description:         "ManagementState defines if the CR should be managed by the operator or not. Default is managed.",
						MarkdownDescription: "ManagementState defines if the CR should be managed by the operator or not. Default is managed.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Managed", "Unmanaged"),
						},
					},

					"replication_factor": {
						Description:         "ReplicationFactor defines the policy for log stream replication.",
						MarkdownDescription: "ReplicationFactor defines the policy for log stream replication.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},

					"rules": {
						Description:         "Rules defines the spec for the ruler component",
						MarkdownDescription: "Rules defines the spec for the ruler component",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled defines a flag to enable/disable the ruler component",
								MarkdownDescription: "Enabled defines a flag to enable/disable the ruler component",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace_selector": {
								Description:         "Namespaces to be selected for PrometheusRules discovery. If unspecified, only the same namespace as the LokiStack object is in is used.",
								MarkdownDescription: "Namespaces to be selected for PrometheusRules discovery. If unspecified, only the same namespace as the LokiStack object is in is used.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": {
								Description:         "A selector to select which LokiRules to mount for loading alerting/recording rules from.",
								MarkdownDescription: "A selector to select which LokiRules to mount for loading alerting/recording rules from.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

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

					"size": {
						Description:         "Size defines one of the support Loki deployment scale out sizes.",
						MarkdownDescription: "Size defines one of the support Loki deployment scale out sizes.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("1x.extra-small", "1x.small", "1x.medium"),
						},
					},

					"storage": {
						Description:         "Storage defines the spec for the object storage endpoint to store logs.",
						MarkdownDescription: "Storage defines the spec for the object storage endpoint to store logs.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"schemas": {
								Description:         "Schemas for reading and writing logs.",
								MarkdownDescription: "Schemas for reading and writing logs.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"effective_date": {
										Description:         "EffectiveDate is the date in UTC that the schema will be applied on. To ensure readibility of logs, this date should be before the current date in UTC.",
										MarkdownDescription: "EffectiveDate is the date in UTC that the schema will be applied on. To ensure readibility of logs, this date should be before the current date in UTC.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{4,})([-]([0-9]{2})){2}$`), ""),
										},
									},

									"version": {
										Description:         "Version for writing and reading logs.",
										MarkdownDescription: "Version for writing and reading logs.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("v11", "v12"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": {
								Description:         "Secret for object storage authentication. Name of a secret in the same namespace as the LokiStack custom resource.",
								MarkdownDescription: "Secret for object storage authentication. Name of a secret in the same namespace as the LokiStack custom resource.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of a secret in the namespace configured for object storage secrets.",
										MarkdownDescription: "Name of a secret in the namespace configured for object storage secrets.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"type": {
										Description:         "Type of object storage that should be used",
										MarkdownDescription: "Type of object storage that should be used",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("azure", "gcs", "s3", "swift"),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"tls": {
								Description:         "TLS configuration for reaching the object storage endpoint.",
								MarkdownDescription: "TLS configuration for reaching the object storage endpoint.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_name": {
										Description:         "CA is the name of a ConfigMap containing a CA certificate. It needs to be in the same namespace as the LokiStack custom resource.",
										MarkdownDescription: "CA is the name of a ConfigMap containing a CA certificate. It needs to be in the same namespace as the LokiStack custom resource.",

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

					"storage_class_name": {
						Description:         "Storage class name defines the storage class for ingester/querier PVCs.",
						MarkdownDescription: "Storage class name defines the storage class for ingester/querier PVCs.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"template": {
						Description:         "Template defines the resource/limits/tolerations/nodeselectors per component",
						MarkdownDescription: "Template defines the resource/limits/tolerations/nodeselectors per component",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"compactor": {
								Description:         "Compactor defines the compaction component spec.",
								MarkdownDescription: "Compactor defines the compaction component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"distributor": {
								Description:         "Distributor defines the distributor component spec.",
								MarkdownDescription: "Distributor defines the distributor component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"gateway": {
								Description:         "Gateway defines the lokistack gateway component spec.",
								MarkdownDescription: "Gateway defines the lokistack gateway component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"index_gateway": {
								Description:         "IndexGateway defines the index gateway component spec.",
								MarkdownDescription: "IndexGateway defines the index gateway component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"ingester": {
								Description:         "Ingester defines the ingester component spec.",
								MarkdownDescription: "Ingester defines the ingester component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"querier": {
								Description:         "Querier defines the querier component spec.",
								MarkdownDescription: "Querier defines the querier component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"query_frontend": {
								Description:         "QueryFrontend defines the query frontend component spec.",
								MarkdownDescription: "QueryFrontend defines the query frontend component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"ruler": {
								Description:         "Ruler defines the ruler component spec.",
								MarkdownDescription: "Ruler defines the ruler component spec.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector": {
										Description:         "NodeSelector defines the labels required by a node to schedule the component onto it.",
										MarkdownDescription: "NodeSelector defines the labels required by a node to schedule the component onto it.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replica pods of the component.",
										MarkdownDescription: "Replicas defines the number of replica pods of the component.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Tolerations defines the tolerations required by a node to schedule the component onto it.",
										MarkdownDescription: "Tolerations defines the tolerations required by a node to schedule the component onto it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

					"tenants": {
						Description:         "Tenants defines the per-tenant authentication and authorization spec for the lokistack-gateway component.",
						MarkdownDescription: "Tenants defines the per-tenant authentication and authorization spec for the lokistack-gateway component.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"authentication": {
								Description:         "Authentication defines the lokistack-gateway component authentication configuration spec per tenant.",
								MarkdownDescription: "Authentication defines the lokistack-gateway component authentication configuration spec per tenant.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"oidc": {
										Description:         "OIDC defines the spec for the OIDC tenant's authentication.",
										MarkdownDescription: "OIDC defines the spec for the OIDC tenant's authentication.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"group_claim": {
												Description:         "Group claim field from ID Token",
												MarkdownDescription: "Group claim field from ID Token",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"issuer_url": {
												Description:         "IssuerURL defines the URL for issuer.",
												MarkdownDescription: "IssuerURL defines the URL for issuer.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"redirect_url": {
												Description:         "RedirectURL defines the URL for redirect.",
												MarkdownDescription: "RedirectURL defines the URL for redirect.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret defines the spec for the clientID, clientSecret and issuerCAPath for tenant's authentication.",
												MarkdownDescription: "Secret defines the spec for the clientID, clientSecret and issuerCAPath for tenant's authentication.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of a secret in the namespace configured for tenant secrets.",
														MarkdownDescription: "Name of a secret in the namespace configured for tenant secrets.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"username_claim": {
												Description:         "User claim field from ID Token",
												MarkdownDescription: "User claim field from ID Token",

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

									"tenant_id": {
										Description:         "TenantID defines the id of the tenant.",
										MarkdownDescription: "TenantID defines the id of the tenant.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"tenant_name": {
										Description:         "TenantName defines the name of the tenant.",
										MarkdownDescription: "TenantName defines the name of the tenant.",

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

							"authorization": {
								Description:         "Authorization defines the lokistack-gateway component authorization configuration spec per tenant.",
								MarkdownDescription: "Authorization defines the lokistack-gateway component authorization configuration spec per tenant.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"opa": {
										Description:         "OPA defines the spec for the third-party endpoint for tenant's authorization.",
										MarkdownDescription: "OPA defines the spec for the third-party endpoint for tenant's authorization.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"url": {
												Description:         "URL defines the third-party endpoint for authorization.",
												MarkdownDescription: "URL defines the third-party endpoint for authorization.",

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

									"role_bindings": {
										Description:         "RoleBindings defines configuration to bind a set of roles to a set of subjects.",
										MarkdownDescription: "RoleBindings defines configuration to bind a set of roles to a set of subjects.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"roles": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"subjects": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"kind": {
														Description:         "SubjectKind is a kind of LokiStack Gateway RBAC subject.",
														MarkdownDescription: "SubjectKind is a kind of LokiStack Gateway RBAC subject.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("user", "group"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

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

										Required: false,
										Optional: true,
										Computed: false,
									},

									"roles": {
										Description:         "Roles defines a set of permissions to interact with a tenant.",
										MarkdownDescription: "Roles defines a set of permissions to interact with a tenant.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"permissions": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"resources": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"tenants": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

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

							"mode": {
								Description:         "Mode defines the mode in which lokistack-gateway component will be configured.",
								MarkdownDescription: "Mode defines the mode in which lokistack-gateway component will be configured.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("static", "dynamic", "openshift-logging"),
								},
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

func (r *LokiGrafanaComLokiStackV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_loki_grafana_com_loki_stack_v1beta1")

	var state LokiGrafanaComLokiStackV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LokiGrafanaComLokiStackV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("loki.grafana.com/v1beta1")
	goModel.Kind = utilities.Ptr("LokiStack")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LokiGrafanaComLokiStackV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_loki_stack_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *LokiGrafanaComLokiStackV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_loki_grafana_com_loki_stack_v1beta1")

	var state LokiGrafanaComLokiStackV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LokiGrafanaComLokiStackV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("loki.grafana.com/v1beta1")
	goModel.Kind = utilities.Ptr("LokiStack")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LokiGrafanaComLokiStackV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_loki_grafana_com_loki_stack_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
