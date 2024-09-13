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
	_ datasource.DataSource = &MonitoringCoreosComPodMonitorV1Manifest{}
)

func NewMonitoringCoreosComPodMonitorV1Manifest() datasource.DataSource {
	return &MonitoringCoreosComPodMonitorV1Manifest{}
}

type MonitoringCoreosComPodMonitorV1Manifest struct{}

type MonitoringCoreosComPodMonitorV1ManifestData struct {
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
		AttachMetadata *struct {
			Node *bool `tfsdk:"node" json:"node,omitempty"`
		} `tfsdk:"attach_metadata" json:"attachMetadata,omitempty"`
		BodySizeLimit         *string `tfsdk:"body_size_limit" json:"bodySizeLimit,omitempty"`
		JobLabel              *string `tfsdk:"job_label" json:"jobLabel,omitempty"`
		KeepDroppedTargets    *int64  `tfsdk:"keep_dropped_targets" json:"keepDroppedTargets,omitempty"`
		LabelLimit            *int64  `tfsdk:"label_limit" json:"labelLimit,omitempty"`
		LabelNameLengthLimit  *int64  `tfsdk:"label_name_length_limit" json:"labelNameLengthLimit,omitempty"`
		LabelValueLengthLimit *int64  `tfsdk:"label_value_length_limit" json:"labelValueLengthLimit,omitempty"`
		NamespaceSelector     *struct {
			Any        *bool     `tfsdk:"any" json:"any,omitempty"`
			MatchNames *[]string `tfsdk:"match_names" json:"matchNames,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		PodMetricsEndpoints *[]struct {
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
			EnableHttp2       *bool   `tfsdk:"enable_http2" json:"enableHttp2,omitempty"`
			FilterRunning     *bool   `tfsdk:"filter_running" json:"filterRunning,omitempty"`
			FollowRedirects   *bool   `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
			HonorLabels       *bool   `tfsdk:"honor_labels" json:"honorLabels,omitempty"`
			HonorTimestamps   *bool   `tfsdk:"honor_timestamps" json:"honorTimestamps,omitempty"`
			Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
			MetricRelabelings *[]struct {
				Action       *string   `tfsdk:"action" json:"action,omitempty"`
				Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"metric_relabelings" json:"metricRelabelings,omitempty"`
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
				EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
				NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
				ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
				ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
				ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
				Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
				TlsConfig            *struct {
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
					MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
					MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
			} `tfsdk:"oauth2" json:"oauth2,omitempty"`
			Params      *map[string][]string `tfsdk:"params" json:"params,omitempty"`
			Path        *string              `tfsdk:"path" json:"path,omitempty"`
			Port        *string              `tfsdk:"port" json:"port,omitempty"`
			ProxyUrl    *string              `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
			Relabelings *[]struct {
				Action       *string   `tfsdk:"action" json:"action,omitempty"`
				Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"relabelings" json:"relabelings,omitempty"`
			Scheme        *string `tfsdk:"scheme" json:"scheme,omitempty"`
			ScrapeTimeout *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
			TargetPort    *string `tfsdk:"target_port" json:"targetPort,omitempty"`
			TlsConfig     *struct {
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
				MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
				MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
				ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			TrackTimestampsStaleness *bool `tfsdk:"track_timestamps_staleness" json:"trackTimestampsStaleness,omitempty"`
		} `tfsdk:"pod_metrics_endpoints" json:"podMetricsEndpoints,omitempty"`
		PodTargetLabels *[]string `tfsdk:"pod_target_labels" json:"podTargetLabels,omitempty"`
		SampleLimit     *int64    `tfsdk:"sample_limit" json:"sampleLimit,omitempty"`
		ScrapeClass     *string   `tfsdk:"scrape_class" json:"scrapeClass,omitempty"`
		ScrapeProtocols *[]string `tfsdk:"scrape_protocols" json:"scrapeProtocols,omitempty"`
		Selector        *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		TargetLimit *int64 `tfsdk:"target_limit" json:"targetLimit,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringCoreosComPodMonitorV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_coreos_com_pod_monitor_v1_manifest"
}

func (r *MonitoringCoreosComPodMonitorV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The 'PodMonitor' custom resource definition (CRD) defines how 'Prometheus' and 'PrometheusAgent' can scrape metrics from a group of pods.Among other things, it allows to specify:* The pods to scrape via label selectors.* The container ports to scrape.* Authentication credentials to use.* Target and metric relabeling.'Prometheus' and 'PrometheusAgent' objects select 'PodMonitor' objects using label and namespace selectors.",
		MarkdownDescription: "The 'PodMonitor' custom resource definition (CRD) defines how 'Prometheus' and 'PrometheusAgent' can scrape metrics from a group of pods.Among other things, it allows to specify:* The pods to scrape via label selectors.* The container ports to scrape.* Authentication credentials to use.* Target and metric relabeling.'Prometheus' and 'PrometheusAgent' objects select 'PodMonitor' objects using label and namespace selectors.",
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
				Description:         "Specification of desired Pod selection for target discovery by Prometheus.",
				MarkdownDescription: "Specification of desired Pod selection for target discovery by Prometheus.",
				Attributes: map[string]schema.Attribute{
					"attach_metadata": schema.SingleNestedAttribute{
						Description:         "'attachMetadata' defines additional metadata which is added to thediscovered targets.It requires Prometheus >= v2.35.0.",
						MarkdownDescription: "'attachMetadata' defines additional metadata which is added to thediscovered targets.It requires Prometheus >= v2.35.0.",
						Attributes: map[string]schema.Attribute{
							"node": schema.BoolAttribute{
								Description:         "When set to true, Prometheus attaches node metadata to the discoveredtargets.The Prometheus service account must have the 'list' and 'watch'permissions on the 'Nodes' objects.",
								MarkdownDescription: "When set to true, Prometheus attaches node metadata to the discoveredtargets.The Prometheus service account must have the 'list' and 'watch'permissions on the 'Nodes' objects.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"body_size_limit": schema.StringAttribute{
						Description:         "When defined, bodySizeLimit specifies a job level limit on the sizeof uncompressed response body that will be accepted by Prometheus.It requires Prometheus >= v2.28.0.",
						MarkdownDescription: "When defined, bodySizeLimit specifies a job level limit on the sizeof uncompressed response body that will be accepted by Prometheus.It requires Prometheus >= v2.28.0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`(^0|([0-9]*[.])?[0-9]+((K|M|G|T|E|P)i?)?B)$`), ""),
						},
					},

					"job_label": schema.StringAttribute{
						Description:         "The label to use to retrieve the job name from.'jobLabel' selects the label from the associated Kubernetes 'Pod'object which will be used as the 'job' label for all metrics.For example if 'jobLabel' is set to 'foo' and the Kubernetes 'Pod'object is labeled with 'foo: bar', then Prometheus adds the 'job='bar''label to all ingested metrics.If the value of this field is empty, the 'job' label of the metricsdefaults to the namespace and name of the PodMonitor object (e.g. '<namespace>/<name>').",
						MarkdownDescription: "The label to use to retrieve the job name from.'jobLabel' selects the label from the associated Kubernetes 'Pod'object which will be used as the 'job' label for all metrics.For example if 'jobLabel' is set to 'foo' and the Kubernetes 'Pod'object is labeled with 'foo: bar', then Prometheus adds the 'job='bar''label to all ingested metrics.If the value of this field is empty, the 'job' label of the metricsdefaults to the namespace and name of the PodMonitor object (e.g. '<namespace>/<name>').",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"keep_dropped_targets": schema.Int64Attribute{
						Description:         "Per-scrape limit on the number of targets dropped by relabelingthat will be kept in memory. 0 means no limit.It requires Prometheus >= v2.47.0.",
						MarkdownDescription: "Per-scrape limit on the number of targets dropped by relabelingthat will be kept in memory. 0 means no limit.It requires Prometheus >= v2.47.0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on number of labels that will be accepted for a sample.It requires Prometheus >= v2.27.0.",
						MarkdownDescription: "Per-scrape limit on number of labels that will be accepted for a sample.It requires Prometheus >= v2.27.0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_name_length_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on length of labels name that will be accepted for a sample.It requires Prometheus >= v2.27.0.",
						MarkdownDescription: "Per-scrape limit on length of labels name that will be accepted for a sample.It requires Prometheus >= v2.27.0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_value_length_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on length of labels value that will be accepted for a sample.It requires Prometheus >= v2.27.0.",
						MarkdownDescription: "Per-scrape limit on length of labels value that will be accepted for a sample.It requires Prometheus >= v2.27.0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "'namespaceSelector' defines in which namespace(s) Prometheus should discover the pods.By default, the pods are discovered in the same namespace as the 'PodMonitor' object but it is possible to select pods across different/all namespaces.",
						MarkdownDescription: "'namespaceSelector' defines in which namespace(s) Prometheus should discover the pods.By default, the pods are discovered in the same namespace as the 'PodMonitor' object but it is possible to select pods across different/all namespaces.",
						Attributes: map[string]schema.Attribute{
							"any": schema.BoolAttribute{
								Description:         "Boolean describing whether all namespaces are selected in contrast to alist restricting them.",
								MarkdownDescription: "Boolean describing whether all namespaces are selected in contrast to alist restricting them.",
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

					"pod_metrics_endpoints": schema.ListNestedAttribute{
						Description:         "Defines how to scrape metrics from the selected pods.",
						MarkdownDescription: "Defines how to scrape metrics from the selected pods.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authorization": schema.SingleNestedAttribute{
									Description:         "'authorization' configures the Authorization header credentials to use whenscraping the target.Cannot be set at the same time as 'basicAuth', or 'oauth2'.",
									MarkdownDescription: "'authorization' configures the Authorization header credentials to use whenscraping the target.Cannot be set at the same time as 'basicAuth', or 'oauth2'.",
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
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
											Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
											MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
									Description:         "'basicAuth' configures the Basic Authentication credentials to use whenscraping the target.Cannot be set at the same time as 'authorization', or 'oauth2'.",
									MarkdownDescription: "'basicAuth' configures the Basic Authentication credentials to use whenscraping the target.Cannot be set at the same time as 'authorization', or 'oauth2'.",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
											MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
											Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
											MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
									Description:         "'bearerTokenSecret' specifies a key of a Secret containing the bearertoken for scraping targets. The secret needs to be in the same namespaceas the PodMonitor object and readable by the Prometheus Operator.Deprecated: use 'authorization' instead.",
									MarkdownDescription: "'bearerTokenSecret' specifies a key of a Secret containing the bearertoken for scraping targets. The secret needs to be in the same namespaceas the PodMonitor object and readable by the Prometheus Operator.Deprecated: use 'authorization' instead.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

								"enable_http2": schema.BoolAttribute{
									Description:         "'enableHttp2' can be used to disable HTTP2 when scraping the target.",
									MarkdownDescription: "'enableHttp2' can be used to disable HTTP2 when scraping the target.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"filter_running": schema.BoolAttribute{
									Description:         "When true, the pods which are not running (e.g. either in Failed orSucceeded state) are dropped during the target discovery.If unset, the filtering is enabled.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase",
									MarkdownDescription: "When true, the pods which are not running (e.g. either in Failed orSucceeded state) are dropped during the target discovery.If unset, the filtering is enabled.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"follow_redirects": schema.BoolAttribute{
									Description:         "'followRedirects' defines whether the scrape requests should follow HTTP3xx redirects.",
									MarkdownDescription: "'followRedirects' defines whether the scrape requests should follow HTTP3xx redirects.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"honor_labels": schema.BoolAttribute{
									Description:         "When true, 'honorLabels' preserves the metric's labels when they collidewith the target's labels.",
									MarkdownDescription: "When true, 'honorLabels' preserves the metric's labels when they collidewith the target's labels.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"honor_timestamps": schema.BoolAttribute{
									Description:         "'honorTimestamps' controls whether Prometheus preserves the timestampswhen exposed by the target.",
									MarkdownDescription: "'honorTimestamps' controls whether Prometheus preserves the timestampswhen exposed by the target.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interval": schema.StringAttribute{
									Description:         "Interval at which Prometheus scrapes the metrics from the target.If empty, Prometheus uses the global scrape interval.",
									MarkdownDescription: "Interval at which Prometheus scrapes the metrics from the target.If empty, Prometheus uses the global scrape interval.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
									},
								},

								"metric_relabelings": schema.ListNestedAttribute{
									Description:         "'metricRelabelings' configures the relabeling rules to apply to thesamples before ingestion.",
									MarkdownDescription: "'metricRelabelings' configures the relabeling rules to apply to thesamples before ingestion.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
												MarkdownDescription: "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
												},
											},

											"modulus": schema.Int64Attribute{
												Description:         "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
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
												Description:         "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
												MarkdownDescription: "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
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
												Description:         "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
												MarkdownDescription: "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_label": schema.StringAttribute{
												Description:         "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
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

								"oauth2": schema.SingleNestedAttribute{
									Description:         "'oauth2' configures the OAuth2 settings to use when scraping the target.It requires Prometheus >= 2.27.0.Cannot be set at the same time as 'authorization', or 'basicAuth'.",
									MarkdownDescription: "'oauth2' configures the OAuth2 settings to use when scraping the target.It requires Prometheus >= 2.27.0.Cannot be set at the same time as 'authorization', or 'basicAuth'.",
									Attributes: map[string]schema.Attribute{
										"client_id": schema.SingleNestedAttribute{
											Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
											MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
											Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
											MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
											Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
											MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"no_proxy": schema.StringAttribute{
											Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0 or Alertmanager >= 0.25.0.",
											MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0 or Alertmanager >= 0.25.0.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_connect_header": schema.MapAttribute{
											Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0 or Alertmanager >= 0.25.0.",
											MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0 or Alertmanager >= 0.25.0.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_from_environment": schema.BoolAttribute{
											Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).It requires Prometheus >= v2.43.0 or Alertmanager >= 0.25.0.",
											MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).It requires Prometheus >= v2.43.0 or Alertmanager >= 0.25.0.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"proxy_url": schema.StringAttribute{
											Description:         "'proxyURL' defines the HTTP proxy server to use.",
											MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
											},
										},

										"scopes": schema.ListAttribute{
											Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
											MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tls_config": schema.SingleNestedAttribute{
											Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
											MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																	Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
																	Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
																	Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
																	Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

												"max_version": schema.StringAttribute{
													Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
													MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
													},
												},

												"min_version": schema.StringAttribute{
													Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
													MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
													},
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

								"params": schema.MapAttribute{
									Description:         "'params' define optional HTTP URL parameters.",
									MarkdownDescription: "'params' define optional HTTP URL parameters.",
									ElementType:         types.ListType{ElemType: types.StringType},
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "HTTP path from which to scrape for metrics.If empty, Prometheus uses the default value (e.g. '/metrics').",
									MarkdownDescription: "HTTP path from which to scrape for metrics.If empty, Prometheus uses the default value (e.g. '/metrics').",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.StringAttribute{
									Description:         "Name of the Pod port which this endpoint refers to.It takes precedence over 'targetPort'.",
									MarkdownDescription: "Name of the Pod port which this endpoint refers to.It takes precedence over 'targetPort'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"proxy_url": schema.StringAttribute{
									Description:         "'proxyURL' configures the HTTP Proxy URL (e.g.'http://proxyserver:2195') to go through when scraping the target.",
									MarkdownDescription: "'proxyURL' configures the HTTP Proxy URL (e.g.'http://proxyserver:2195') to go through when scraping the target.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"relabelings": schema.ListNestedAttribute{
									Description:         "'relabelings' configures the relabeling rules to apply the target'smetadata labels.The Operator automatically adds relabelings for a few standard Kubernetes fields.The original scrape job's name is available via the '__tmp_prometheus_job_name' label.More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
									MarkdownDescription: "'relabelings' configures the relabeling rules to apply the target'smetadata labels.The Operator automatically adds relabelings for a few standard Kubernetes fields.The original scrape job's name is available via the '__tmp_prometheus_job_name' label.More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
												MarkdownDescription: "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
												},
											},

											"modulus": schema.Int64Attribute{
												Description:         "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
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
												Description:         "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
												MarkdownDescription: "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
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
												Description:         "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
												MarkdownDescription: "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_label": schema.StringAttribute{
												Description:         "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
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

								"scheme": schema.StringAttribute{
									Description:         "HTTP scheme to use for scraping.'http' and 'https' are the expected values unless you rewrite the'__scheme__' label via relabeling.If empty, Prometheus uses the default value 'http'.",
									MarkdownDescription: "HTTP scheme to use for scraping.'http' and 'https' are the expected values unless you rewrite the'__scheme__' label via relabeling.If empty, Prometheus uses the default value 'http'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("http", "https"),
									},
								},

								"scrape_timeout": schema.StringAttribute{
									Description:         "Timeout after which Prometheus considers the scrape to be failed.If empty, Prometheus uses the global scrape timeout unless it is lessthan the target's scrape interval value in which the latter is used.",
									MarkdownDescription: "Timeout after which Prometheus considers the scrape to be failed.If empty, Prometheus uses the global scrape timeout unless it is lessthan the target's scrape interval value in which the latter is used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
									},
								},

								"target_port": schema.StringAttribute{
									Description:         "Name or number of the target port of the 'Pod' object behind the Service, theport must be specified with container port property.Deprecated: use 'port' instead.",
									MarkdownDescription: "Name or number of the target port of the 'Pod' object behind the Service, theport must be specified with container port property.Deprecated: use 'port' instead.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tls_config": schema.SingleNestedAttribute{
									Description:         "TLS configuration to use when scraping the target.",
									MarkdownDescription: "TLS configuration to use when scraping the target.",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

										"max_version": schema.StringAttribute{
											Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
											MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
											},
										},

										"min_version": schema.StringAttribute{
											Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
											MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
											},
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

								"track_timestamps_staleness": schema.BoolAttribute{
									Description:         "'trackTimestampsStaleness' defines whether Prometheus tracks staleness ofthe metrics that have an explicit timestamp present in scraped data.Has no effect if 'honorTimestamps' is false.It requires Prometheus >= v2.48.0.",
									MarkdownDescription: "'trackTimestampsStaleness' defines whether Prometheus tracks staleness ofthe metrics that have an explicit timestamp present in scraped data.Has no effect if 'honorTimestamps' is false.It requires Prometheus >= v2.48.0.",
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

					"pod_target_labels": schema.ListAttribute{
						Description:         "'podTargetLabels' defines the labels which are transferred from theassociated Kubernetes 'Pod' object onto the ingested metrics.",
						MarkdownDescription: "'podTargetLabels' defines the labels which are transferred from theassociated Kubernetes 'Pod' object onto the ingested metrics.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sample_limit": schema.Int64Attribute{
						Description:         "'sampleLimit' defines a per-scrape limit on the number of scraped samplesthat will be accepted.",
						MarkdownDescription: "'sampleLimit' defines a per-scrape limit on the number of scraped samplesthat will be accepted.",
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
						Description:         "'scrapeProtocols' defines the protocols to negotiate during a scrape. It tells clients theprotocols supported by Prometheus in order of preference (from most to least preferred).If unset, Prometheus uses its default value.It requires Prometheus >= v2.49.0.",
						MarkdownDescription: "'scrapeProtocols' defines the protocols to negotiate during a scrape. It tells clients theprotocols supported by Prometheus in order of preference (from most to least preferred).If unset, Prometheus uses its default value.It requires Prometheus >= v2.49.0.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Label selector to select the Kubernetes 'Pod' objects to scrape metrics from.",
						MarkdownDescription: "Label selector to select the Kubernetes 'Pod' objects to scrape metrics from.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_limit": schema.Int64Attribute{
						Description:         "'targetLimit' defines a limit on the number of scraped targets that willbe accepted.",
						MarkdownDescription: "'targetLimit' defines a limit on the number of scraped targets that willbe accepted.",
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
	}
}

func (r *MonitoringCoreosComPodMonitorV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_pod_monitor_v1_manifest")

	var model MonitoringCoreosComPodMonitorV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("monitoring.coreos.com/v1")
	model.Kind = pointer.String("PodMonitor")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
