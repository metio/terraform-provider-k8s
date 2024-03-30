/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MonitoringCoreosComProbeV1Manifest{}
)

func NewMonitoringCoreosComProbeV1Manifest() datasource.DataSource {
	return &MonitoringCoreosComProbeV1Manifest{}
}

type MonitoringCoreosComProbeV1Manifest struct{}

type MonitoringCoreosComProbeV1ManifestData struct {
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
		Authorization *struct {
			Credentials *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"authorization" json:"authorization,omitempty"`
		BasicAuth *struct {
			Password *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Username *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
		BearerTokenSecret *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
		Interval              *string `tfsdk:"interval" json:"interval,omitempty"`
		JobName               *string `tfsdk:"job_name" json:"jobName,omitempty"`
		KeepDroppedTargets    *int64  `tfsdk:"keep_dropped_targets" json:"keepDroppedTargets,omitempty"`
		LabelLimit            *int64  `tfsdk:"label_limit" json:"labelLimit,omitempty"`
		LabelNameLengthLimit  *int64  `tfsdk:"label_name_length_limit" json:"labelNameLengthLimit,omitempty"`
		LabelValueLengthLimit *int64  `tfsdk:"label_value_length_limit" json:"labelValueLengthLimit,omitempty"`
		MetricRelabelings     *[]struct {
			Action       *string   `tfsdk:"action" json:"action,omitempty"`
			Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"metric_relabelings" json:"metricRelabelings,omitempty"`
		Module *string `tfsdk:"module" json:"module,omitempty"`
		Oauth2 *struct {
			ClientId *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
			} `tfsdk:"client_id" json:"clientId,omitempty"`
			ClientSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			EndpointParams *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
			Scopes         *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
			TokenUrl       *string            `tfsdk:"token_url" json:"tokenUrl,omitempty"`
		} `tfsdk:"oauth2" json:"oauth2,omitempty"`
		Prober *struct {
			Path     *string `tfsdk:"path" json:"path,omitempty"`
			ProxyUrl *string `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
			Scheme   *string `tfsdk:"scheme" json:"scheme,omitempty"`
			Url      *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"prober" json:"prober,omitempty"`
		SampleLimit     *int64    `tfsdk:"sample_limit" json:"sampleLimit,omitempty"`
		ScrapeClass     *string   `tfsdk:"scrape_class" json:"scrapeClass,omitempty"`
		ScrapeProtocols *[]string `tfsdk:"scrape_protocols" json:"scrapeProtocols,omitempty"`
		ScrapeTimeout   *string   `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
		TargetLimit     *int64    `tfsdk:"target_limit" json:"targetLimit,omitempty"`
		Targets         *struct {
			Ingress *struct {
				NamespaceSelector *struct {
					Any        *bool     `tfsdk:"any" json:"any,omitempty"`
					MatchNames *[]string `tfsdk:"match_names" json:"matchNames,omitempty"`
				} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				RelabelingConfigs *[]struct {
					Action       *string   `tfsdk:"action" json:"action,omitempty"`
					Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
					Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
					Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
					Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
					SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
					TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
				} `tfsdk:"relabeling_configs" json:"relabelingConfigs,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			StaticConfig *struct {
				Labels            *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				RelabelingConfigs *[]struct {
					Action       *string   `tfsdk:"action" json:"action,omitempty"`
					Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
					Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
					Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
					Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
					SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
					TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
				} `tfsdk:"relabeling_configs" json:"relabelingConfigs,omitempty"`
				Static *[]string `tfsdk:"static" json:"static,omitempty"`
			} `tfsdk:"static_config" json:"staticConfig,omitempty"`
		} `tfsdk:"targets" json:"targets,omitempty"`
		TlsConfig *struct {
			Ca *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
			} `tfsdk:"ca" json:"ca,omitempty"`
			Cert *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
			} `tfsdk:"cert" json:"cert,omitempty"`
			InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			KeySecret          *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_secret" json:"keySecret,omitempty"`
			ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
		} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringCoreosComProbeV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_coreos_com_probe_v1_manifest"
}

func (r *MonitoringCoreosComProbeV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Probe defines monitoring for a set of static targets or ingresses.",
		MarkdownDescription: "Probe defines monitoring for a set of static targets or ingresses.",
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
				Description:         "Specification of desired Ingress selection for target discovery by Prometheus.",
				MarkdownDescription: "Specification of desired Ingress selection for target discovery by Prometheus.",
				Attributes: map[string]schema.Attribute{
					"authorization": schema.SingleNestedAttribute{
						Description:         "Authorization section for this endpoint",
						MarkdownDescription: "Authorization section for this endpoint",
						Attributes: map[string]schema.Attribute{
							"credentials": schema.SingleNestedAttribute{
								Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
								MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
								MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"basic_auth": schema.SingleNestedAttribute{
						Description:         "BasicAuth allow an endpoint to authenticate over basic authentication. More info: https://prometheus.io/docs/operating/configuration/#endpoint",
						MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication. More info: https://prometheus.io/docs/operating/configuration/#endpoint",
						Attributes: map[string]schema.Attribute{
							"password": schema.SingleNestedAttribute{
								Description:         "'password' specifies a key of a Secret containing the password for authentication.",
								MarkdownDescription: "'password' specifies a key of a Secret containing the password for authentication.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": schema.SingleNestedAttribute{
								Description:         "'username' specifies a key of a Secret containing the username for authentication.",
								MarkdownDescription: "'username' specifies a key of a Secret containing the username for authentication.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

					"bearer_token_secret": schema.SingleNestedAttribute{
						Description:         "Secret to mount to read bearer token for scraping targets. The secret needs to be in the same namespace as the probe and accessible by the Prometheus Operator.",
						MarkdownDescription: "Secret to mount to read bearer token for scraping targets. The secret needs to be in the same namespace as the probe and accessible by the Prometheus Operator.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key of the secret to select from.  Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"optional": schema.BoolAttribute{
								Description:         "Specify whether the Secret or its key must be defined",
								MarkdownDescription: "Specify whether the Secret or its key must be defined",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval at which targets are probed using the configured prober. If not specified Prometheus' global scrape interval is used.",
						MarkdownDescription: "Interval at which targets are probed using the configured prober. If not specified Prometheus' global scrape interval is used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
						},
					},

					"job_name": schema.StringAttribute{
						Description:         "The job name assigned to scraped metrics by default.",
						MarkdownDescription: "The job name assigned to scraped metrics by default.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"keep_dropped_targets": schema.Int64Attribute{
						Description:         "Per-scrape limit on the number of targets dropped by relabeling that will be kept in memory. 0 means no limit.  It requires Prometheus >= v2.47.0.",
						MarkdownDescription: "Per-scrape limit on the number of targets dropped by relabeling that will be kept in memory. 0 means no limit.  It requires Prometheus >= v2.47.0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_name_length_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_value_length_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metric_relabelings": schema.ListNestedAttribute{
						Description:         "MetricRelabelConfigs to apply to samples before ingestion.",
						MarkdownDescription: "MetricRelabelConfigs to apply to samples before ingestion.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
									MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
									},
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.StringAttribute{
									Description:         "Regular expression against which the extracted value is matched.",
									MarkdownDescription: "Regular expression against which the extracted value is matched.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
									MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator is the string between concatenated SourceLabels.",
									MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

					"module": schema.StringAttribute{
						Description:         "The module to use for probing specifying how to probe the target. Example module configuring in the blackbox exporter: https://github.com/prometheus/blackbox_exporter/blob/master/example.yml",
						MarkdownDescription: "The module to use for probing specifying how to probe the target. Example module configuring in the blackbox exporter: https://github.com/prometheus/blackbox_exporter/blob/master/example.yml",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"oauth2": schema.SingleNestedAttribute{
						Description:         "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",
						Attributes: map[string]schema.Attribute{
							"client_id": schema.SingleNestedAttribute{
								Description:         "'clientId' specifies a key of a Secret or ConfigMap containing the OAuth2 client's ID.",
								MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing the OAuth2 client's ID.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

							"client_secret": schema.SingleNestedAttribute{
								Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2 client's secret.",
								MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2 client's secret.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"endpoint_params": schema.MapAttribute{
								Description:         "'endpointParams' configures the HTTP parameters to append to the token URL.",
								MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the token URL.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scopes": schema.ListAttribute{
								Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
								MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_url": schema.StringAttribute{
								Description:         "'tokenURL' configures the URL to fetch the token from.",
								MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"prober": schema.SingleNestedAttribute{
						Description:         "Specification for the prober to use for probing targets. The prober.URL parameter is required. Targets cannot be probed if left empty.",
						MarkdownDescription: "Specification for the prober to use for probing targets. The prober.URL parameter is required. Targets cannot be probed if left empty.",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "Path to collect metrics from. Defaults to '/probe'.",
								MarkdownDescription: "Path to collect metrics from. Defaults to '/probe'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_url": schema.StringAttribute{
								Description:         "Optional ProxyURL.",
								MarkdownDescription: "Optional ProxyURL.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheme": schema.StringAttribute{
								Description:         "HTTP scheme to use for scraping. 'http' and 'https' are the expected values unless you rewrite the '__scheme__' label via relabeling. If empty, Prometheus uses the default value 'http'.",
								MarkdownDescription: "HTTP scheme to use for scraping. 'http' and 'https' are the expected values unless you rewrite the '__scheme__' label via relabeling. If empty, Prometheus uses the default value 'http'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("http", "https"),
								},
							},

							"url": schema.StringAttribute{
								Description:         "Mandatory URL of the prober.",
								MarkdownDescription: "Mandatory URL of the prober.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sample_limit": schema.Int64Attribute{
						Description:         "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						MarkdownDescription: "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scrape_class": schema.StringAttribute{
						Description:         "The scrape class to apply.",
						MarkdownDescription: "The scrape class to apply.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"scrape_protocols": schema.ListAttribute{
						Description:         "'scrapeProtocols' defines the protocols to negotiate during a scrape. It tells clients the protocols supported by Prometheus in order of preference (from most to least preferred).  If unset, Prometheus uses its default value.  It requires Prometheus >= v2.49.0.",
						MarkdownDescription: "'scrapeProtocols' defines the protocols to negotiate during a scrape. It tells clients the protocols supported by Prometheus in order of preference (from most to least preferred).  If unset, Prometheus uses its default value.  It requires Prometheus >= v2.49.0.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scrape_timeout": schema.StringAttribute{
						Description:         "Timeout for scraping metrics from the Prometheus exporter. If not specified, the Prometheus global scrape timeout is used.",
						MarkdownDescription: "Timeout for scraping metrics from the Prometheus exporter. If not specified, the Prometheus global scrape timeout is used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
						},
					},

					"target_limit": schema.Int64Attribute{
						Description:         "TargetLimit defines a limit on the number of scraped targets that will be accepted.",
						MarkdownDescription: "TargetLimit defines a limit on the number of scraped targets that will be accepted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"targets": schema.SingleNestedAttribute{
						Description:         "Targets defines a set of static or dynamically discovered targets to probe.",
						MarkdownDescription: "Targets defines a set of static or dynamically discovered targets to probe.",
						Attributes: map[string]schema.Attribute{
							"ingress": schema.SingleNestedAttribute{
								Description:         "ingress defines the Ingress objects to probe and the relabeling configuration. If 'staticConfig' is also defined, 'staticConfig' takes precedence.",
								MarkdownDescription: "ingress defines the Ingress objects to probe and the relabeling configuration. If 'staticConfig' is also defined, 'staticConfig' takes precedence.",
								Attributes: map[string]schema.Attribute{
									"namespace_selector": schema.SingleNestedAttribute{
										Description:         "From which namespaces to select Ingress objects.",
										MarkdownDescription: "From which namespaces to select Ingress objects.",
										Attributes: map[string]schema.Attribute{
											"any": schema.BoolAttribute{
												Description:         "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
												MarkdownDescription: "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_names": schema.ListAttribute{
												Description:         "List of namespace names to select from.",
												MarkdownDescription: "List of namespace names to select from.",
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

									"relabeling_configs": schema.ListNestedAttribute{
										Description:         "RelabelConfigs to apply to the label set of the target before it gets scraped. The original ingress address is available via the '__tmp_prometheus_ingress_address' label. It can be used to customize the probed URL. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
										MarkdownDescription: "RelabelConfigs to apply to the label set of the target before it gets scraped. The original ingress address is available via the '__tmp_prometheus_ingress_address' label. It can be used to customize the probed URL. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
													MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
													},
												},

												"modulus": schema.Int64Attribute{
													Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
													MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"regex": schema.StringAttribute{
													Description:         "Regular expression against which the extracted value is matched.",
													MarkdownDescription: "Regular expression against which the extracted value is matched.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replacement": schema.StringAttribute{
													Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
													MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"separator": schema.StringAttribute{
													Description:         "Separator is the string between concatenated SourceLabels.",
													MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_labels": schema.ListAttribute{
													Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
													MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_label": schema.StringAttribute{
													Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
													MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

									"selector": schema.SingleNestedAttribute{
										Description:         "Selector to select the Ingress objects.",
										MarkdownDescription: "Selector to select the Ingress objects.",
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

							"static_config": schema.SingleNestedAttribute{
								Description:         "staticConfig defines the static list of targets to probe and the relabeling configuration. If 'ingress' is also defined, 'staticConfig' takes precedence. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#static_config.",
								MarkdownDescription: "staticConfig defines the static list of targets to probe and the relabeling configuration. If 'ingress' is also defined, 'staticConfig' takes precedence. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#static_config.",
								Attributes: map[string]schema.Attribute{
									"labels": schema.MapAttribute{
										Description:         "Labels assigned to all metrics scraped from the targets.",
										MarkdownDescription: "Labels assigned to all metrics scraped from the targets.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"relabeling_configs": schema.ListNestedAttribute{
										Description:         "RelabelConfigs to apply to the label set of the targets before it gets scraped. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
										MarkdownDescription: "RelabelConfigs to apply to the label set of the targets before it gets scraped. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
													MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
													},
												},

												"modulus": schema.Int64Attribute{
													Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
													MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"regex": schema.StringAttribute{
													Description:         "Regular expression against which the extracted value is matched.",
													MarkdownDescription: "Regular expression against which the extracted value is matched.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replacement": schema.StringAttribute{
													Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
													MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"separator": schema.StringAttribute{
													Description:         "Separator is the string between concatenated SourceLabels.",
													MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_labels": schema.ListAttribute{
													Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
													MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_label": schema.StringAttribute{
													Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
													MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

									"static": schema.ListAttribute{
										Description:         "The list of hosts to probe.",
										MarkdownDescription: "The list of hosts to probe.",
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

					"tls_config": schema.SingleNestedAttribute{
						Description:         "TLS configuration to use when scraping the endpoint.",
						MarkdownDescription: "TLS configuration to use when scraping the endpoint.",
						Attributes: map[string]schema.Attribute{
							"ca": schema.SingleNestedAttribute{
								Description:         "Certificate authority used when verifying server certificates.",
								MarkdownDescription: "Certificate authority used when verifying server certificates.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

							"cert": schema.SingleNestedAttribute{
								Description:         "Client certificate to present when doing client-authentication.",
								MarkdownDescription: "Client certificate to present when doing client-authentication.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "Disable target certificate validation.",
								MarkdownDescription: "Disable target certificate validation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_secret": schema.SingleNestedAttribute{
								Description:         "Secret containing the client key file for the targets.",
								MarkdownDescription: "Secret containing the client key file for the targets.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_name": schema.StringAttribute{
								Description:         "Used to verify the hostname for the targets.",
								MarkdownDescription: "Used to verify the hostname for the targets.",
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
	}
}

func (r *MonitoringCoreosComProbeV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_probe_v1_manifest")

	var model MonitoringCoreosComProbeV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("monitoring.coreos.com/v1")
	model.Kind = pointer.String("Probe")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
