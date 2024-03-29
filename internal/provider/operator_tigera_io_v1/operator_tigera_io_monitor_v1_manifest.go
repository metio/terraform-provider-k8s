/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoMonitorV1Manifest{}
)

func NewOperatorTigeraIoMonitorV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoMonitorV1Manifest{}
}

type OperatorTigeraIoMonitorV1Manifest struct{}

type OperatorTigeraIoMonitorV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AlertManager *struct {
			Spec *struct {
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"alert_manager" json:"alertManager,omitempty"`
		ExternalPrometheus *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ServiceMonitor *struct {
				Endpoints *[]struct {
					BearerTokenSecret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
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
					Params      *map[string][]string `tfsdk:"params" json:"params,omitempty"`
					Relabelings *[]struct {
						Action       *string   `tfsdk:"action" json:"action,omitempty"`
						Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"relabelings" json:"relabelings,omitempty"`
					ScrapeTimeout *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
				} `tfsdk:"endpoints" json:"endpoints,omitempty"`
				Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
		} `tfsdk:"external_prometheus" json:"externalPrometheus,omitempty"`
		Prometheus *struct {
			Spec *struct {
				CommonPrometheusFields *struct {
					Containers *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Resources *struct {
							Claims *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"claims" json:"claims,omitempty"`
							Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					Resources *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"common_prometheus_fields" json:"commonPrometheusFields,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"prometheus" json:"prometheus,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoMonitorV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_monitor_v1_manifest"
}

func (r *OperatorTigeraIoMonitorV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Monitor is the Schema for the monitor API. At most one instance of this resource is supported. It must be named 'tigera-secure'.",
		MarkdownDescription: "Monitor is the Schema for the monitor API. At most one instance of this resource is supported. It must be named 'tigera-secure'.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "MonitorSpec defines the desired state of Tigera monitor.",
				MarkdownDescription: "MonitorSpec defines the desired state of Tigera monitor.",
				Attributes: map[string]schema.Attribute{
					"alert_manager": schema.SingleNestedAttribute{
						Description:         "AlertManager is the configuration for the AlertManager.",
						MarkdownDescription: "AlertManager is the configuration for the AlertManager.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the Alertmanager.",
								MarkdownDescription: "Spec is the specification of the Alertmanager.",
								Attributes: map[string]schema.Attribute{
									"resources": schema.SingleNestedAttribute{
										Description:         "Define resources requests and limits for single Pods.",
										MarkdownDescription: "Define resources requests and limits for single Pods.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_prometheus": schema.SingleNestedAttribute{
						Description:         "ExternalPrometheus optionally configures integration with an external Prometheus for scraping Calico metrics. When specified, the operator will render resources in the defined namespace. This option can be useful for configuring scraping from git-ops tools without the need of post-installation steps.",
						MarkdownDescription: "ExternalPrometheus optionally configures integration with an external Prometheus for scraping Calico metrics. When specified, the operator will render resources in the defined namespace. This option can be useful for configuring scraping from git-ops tools without the need of post-installation steps.",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace where the operator will create resources for your Prometheus instance. The namespace must be created before the operator will create Prometheus resources.",
								MarkdownDescription: "Namespace is the namespace where the operator will create resources for your Prometheus instance. The namespace must be created before the operator will create Prometheus resources.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"service_monitor": schema.SingleNestedAttribute{
								Description:         "ServiceMonitor when specified, the operator will create a ServiceMonitor object in the namespace. It is recommended that you configure labels if you want your prometheus instance to pick up the configuration automatically. The operator will configure 1 endpoint by default: - Params to scrape all metrics available in Calico Enterprise. - BearerTokenSecret (If not overridden, the operator will also create corresponding RBAC that allows authz to the metrics.) - TLSConfig, containing the caFile and serverName.",
								MarkdownDescription: "ServiceMonitor when specified, the operator will create a ServiceMonitor object in the namespace. It is recommended that you configure labels if you want your prometheus instance to pick up the configuration automatically. The operator will configure 1 endpoint by default: - Params to scrape all metrics available in Calico Enterprise. - BearerTokenSecret (If not overridden, the operator will also create corresponding RBAC that allows authz to the metrics.) - TLSConfig, containing the caFile and serverName.",
								Attributes: map[string]schema.Attribute{
									"endpoints": schema.ListNestedAttribute{
										Description:         "The endpoints to scrape. This struct contains a subset of the Endpoint as defined in the prometheus docs. Fields related to connecting to our Prometheus server are automatically set by the operator.",
										MarkdownDescription: "The endpoints to scrape. This struct contains a subset of the Endpoint as defined in the prometheus docs. Fields related to connecting to our Prometheus server are automatically set by the operator.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"bearer_token_secret": schema.SingleNestedAttribute{
													Description:         "Secret to mount to read bearer token for scraping targets. Recommended: when unset, the operator will create a Secret, a ClusterRole and a ClusterRoleBinding.",
													MarkdownDescription: "Secret to mount to read bearer token for scraping targets. Recommended: when unset, the operator will create a Secret, a ClusterRole and a ClusterRoleBinding.",
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

												"honor_labels": schema.BoolAttribute{
													Description:         "HonorLabels chooses the metric's labels on collisions with target labels.",
													MarkdownDescription: "HonorLabels chooses the metric's labels on collisions with target labels.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"honor_timestamps": schema.BoolAttribute{
													Description:         "HonorTimestamps controls whether Prometheus respects the timestamps present in scraped data.",
													MarkdownDescription: "HonorTimestamps controls whether Prometheus respects the timestamps present in scraped data.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Interval at which metrics should be scraped. If not specified Prometheus' global scrape interval is used.",
													MarkdownDescription: "Interval at which metrics should be scraped. If not specified Prometheus' global scrape interval is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
													},
												},

												"metric_relabelings": schema.ListNestedAttribute{
													Description:         "MetricRelabelConfigs to apply to samples before ingestion.",
													MarkdownDescription: "MetricRelabelConfigs to apply to samples before ingestion.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"action": schema.StringAttribute{
																Description:         "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
																MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase"),
																},
															},

															"modulus": schema.Int64Attribute{
																Description:         "Modulus to take of the hash of the source label values.",
																MarkdownDescription: "Modulus to take of the hash of the source label values.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
																MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"replacement": schema.StringAttribute{
																Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
																MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"separator": schema.StringAttribute{
																Description:         "Separator placed between concatenated source label values. default is ';'.",
																MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"source_labels": schema.ListAttribute{
																Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
																MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_label": schema.StringAttribute{
																Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
																MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
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

												"params": schema.MapAttribute{
													Description:         "Optional HTTP URL parameters Default: scrape all metrics.",
													MarkdownDescription: "Optional HTTP URL parameters Default: scrape all metrics.",
													ElementType:         types.ListType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"relabelings": schema.ListNestedAttribute{
													Description:         "RelabelConfigs to apply to samples before scraping. Prometheus Operator automatically adds relabelings for a few standard Kubernetes fields. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
													MarkdownDescription: "RelabelConfigs to apply to samples before scraping. Prometheus Operator automatically adds relabelings for a few standard Kubernetes fields. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"action": schema.StringAttribute{
																Description:         "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
																MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase"),
																},
															},

															"modulus": schema.Int64Attribute{
																Description:         "Modulus to take of the hash of the source label values.",
																MarkdownDescription: "Modulus to take of the hash of the source label values.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
																MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"replacement": schema.StringAttribute{
																Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
																MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"separator": schema.StringAttribute{
																Description:         "Separator placed between concatenated source label values. default is ';'.",
																MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"source_labels": schema.ListAttribute{
																Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
																MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_label": schema.StringAttribute{
																Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
																MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
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

												"scrape_timeout": schema.StringAttribute{
													Description:         "Timeout after which the scrape is ended. If not specified, the Prometheus global scrape timeout is used unless it is less than 'Interval' in which the latter is used.",
													MarkdownDescription: "Timeout after which the scrape is ended. If not specified, the Prometheus global scrape timeout is used unless it is less than 'Interval' in which the latter is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are the metadata.labels of the ServiceMonitor. When combined with spec.serviceMonitorSelector.matchLabels on your prometheus instance, the service monitor will automatically be picked up. Default: k8s-app=tigera-prometheus",
										MarkdownDescription: "Labels are the metadata.labels of the ServiceMonitor. When combined with spec.serviceMonitorSelector.matchLabels on your prometheus instance, the service monitor will automatically be picked up. Default: k8s-app=tigera-prometheus",
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

					"prometheus": schema.SingleNestedAttribute{
						Description:         "Prometheus is the configuration for the Prometheus.",
						MarkdownDescription: "Prometheus is the configuration for the Prometheus.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the Prometheus.",
								MarkdownDescription: "Spec is the specification of the Prometheus.",
								Attributes: map[string]schema.Attribute{
									"common_prometheus_fields": schema.SingleNestedAttribute{
										Description:         "CommonPrometheusFields are the options available to both the Prometheus server and agent.",
										MarkdownDescription: "CommonPrometheusFields are the options available to both the Prometheus server and agent.",
										Attributes: map[string]schema.Attribute{
											"containers": schema.ListNestedAttribute{
												Description:         "Containers is a list of Prometheus containers. If specified, this overrides the specified Prometheus Deployment containers. If omitted, the Prometheus Deployment will use its default values for its containers.",
												MarkdownDescription: "Containers is a list of Prometheus containers. If specified, this overrides the specified Prometheus Deployment containers. If omitted, the Prometheus Deployment will use its default values for its containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is an enum which identifies the Prometheus Deployment container by name.",
															MarkdownDescription: "Name is an enum which identifies the Prometheus Deployment container by name.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("authn-proxy"),
															},
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Prometheus container's resources. If omitted, the Prometheus will use its default value for this container's resources.",
															MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Prometheus container's resources. If omitted, the Prometheus will use its default value for this container's resources.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Define resources requests and limits for single Pods.",
												MarkdownDescription: "Define resources requests and limits for single Pods.",
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
		},
	}
}

func (r *OperatorTigeraIoMonitorV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_monitor_v1_manifest")

	var model OperatorTigeraIoMonitorV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("Monitor")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
