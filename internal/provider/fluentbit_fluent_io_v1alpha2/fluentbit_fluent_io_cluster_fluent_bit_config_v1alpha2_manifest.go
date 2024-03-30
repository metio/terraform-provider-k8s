/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentbit_fluent_io_v1alpha2

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
	_ datasource.DataSource = &FluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest{}
)

func NewFluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest() datasource.DataSource {
	return &FluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest{}
}

type FluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest struct{}

type FluentbitFluentIoClusterFluentBitConfigV1Alpha2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		FilterSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"filter_selector" json:"filterSelector,omitempty"`
		InputSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"input_selector" json:"inputSelector,omitempty"`
		Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
		OutputSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"output_selector" json:"outputSelector,omitempty"`
		ParserSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"parser_selector" json:"parserSelector,omitempty"`
		Service *struct {
			Daemon              *bool   `tfsdk:"daemon" json:"daemon,omitempty"`
			FlushSeconds        *int64  `tfsdk:"flush_seconds" json:"flushSeconds,omitempty"`
			GraceSeconds        *int64  `tfsdk:"grace_seconds" json:"graceSeconds,omitempty"`
			HcErrorsCount       *int64  `tfsdk:"hc_errors_count" json:"hcErrorsCount,omitempty"`
			HcPeriod            *int64  `tfsdk:"hc_period" json:"hcPeriod,omitempty"`
			HcRetryFailureCount *int64  `tfsdk:"hc_retry_failure_count" json:"hcRetryFailureCount,omitempty"`
			HealthCheck         *bool   `tfsdk:"health_check" json:"healthCheck,omitempty"`
			HttpListen          *string `tfsdk:"http_listen" json:"httpListen,omitempty"`
			HttpPort            *int64  `tfsdk:"http_port" json:"httpPort,omitempty"`
			HttpServer          *bool   `tfsdk:"http_server" json:"httpServer,omitempty"`
			LogFile             *string `tfsdk:"log_file" json:"logFile,omitempty"`
			LogLevel            *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			ParsersFile         *string `tfsdk:"parsers_file" json:"parsersFile,omitempty"`
			Storage             *struct {
				BacklogMemLimit           *string `tfsdk:"backlog_mem_limit" json:"backlogMemLimit,omitempty"`
				Checksum                  *string `tfsdk:"checksum" json:"checksum,omitempty"`
				DeleteIrrecoverableChunks *string `tfsdk:"delete_irrecoverable_chunks" json:"deleteIrrecoverableChunks,omitempty"`
				MaxChunksUp               *int64  `tfsdk:"max_chunks_up" json:"maxChunksUp,omitempty"`
				Metrics                   *string `tfsdk:"metrics" json:"metrics,omitempty"`
				Path                      *string `tfsdk:"path" json:"path,omitempty"`
				Sync                      *string `tfsdk:"sync" json:"sync,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest"
}

func (r *FluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterFluentBitConfig is the Schema for the cluster-level fluentbitconfigs API",
		MarkdownDescription: "ClusterFluentBitConfig is the Schema for the cluster-level fluentbitconfigs API",
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
				Description:         "FluentBitConfigSpec defines the desired state of ClusterFluentBitConfig",
				MarkdownDescription: "FluentBitConfigSpec defines the desired state of ClusterFluentBitConfig",
				Attributes: map[string]schema.Attribute{
					"filter_selector": schema.SingleNestedAttribute{
						Description:         "Select filter plugins",
						MarkdownDescription: "Select filter plugins",
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

					"input_selector": schema.SingleNestedAttribute{
						Description:         "Select input plugins",
						MarkdownDescription: "Select input plugins",
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

					"namespace": schema.StringAttribute{
						Description:         "If namespace is defined, then the configmap and secret for fluent-bit is in this namespace. If it is not defined, it is in the namespace of the fluentd-operator",
						MarkdownDescription: "If namespace is defined, then the configmap and secret for fluent-bit is in this namespace. If it is not defined, it is in the namespace of the fluentd-operator",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"output_selector": schema.SingleNestedAttribute{
						Description:         "Select output plugins",
						MarkdownDescription: "Select output plugins",
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

					"parser_selector": schema.SingleNestedAttribute{
						Description:         "Select parser plugins",
						MarkdownDescription: "Select parser plugins",
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

					"service": schema.SingleNestedAttribute{
						Description:         "Service defines the global behaviour of the Fluent Bit engine.",
						MarkdownDescription: "Service defines the global behaviour of the Fluent Bit engine.",
						Attributes: map[string]schema.Attribute{
							"daemon": schema.BoolAttribute{
								Description:         "If true go to background on start",
								MarkdownDescription: "If true go to background on start",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flush_seconds": schema.Int64Attribute{
								Description:         "Interval to flush output",
								MarkdownDescription: "Interval to flush output",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grace_seconds": schema.Int64Attribute{
								Description:         "Wait time on exit",
								MarkdownDescription: "Wait time on exit",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hc_errors_count": schema.Int64Attribute{
								Description:         "the error count to meet the unhealthy requirement, this is a sum for all output plugins in a defined HC_Period, example for output error: [2022/02/16 10:44:10] [ warn] [engine] failed to flush chunk '1-1645008245.491540684.flb', retry in 7 seconds: task_id=0, input=forward.1 > output=cloudwatch_logs.3 (out_id=3)",
								MarkdownDescription: "the error count to meet the unhealthy requirement, this is a sum for all output plugins in a defined HC_Period, example for output error: [2022/02/16 10:44:10] [ warn] [engine] failed to flush chunk '1-1645008245.491540684.flb', retry in 7 seconds: task_id=0, input=forward.1 > output=cloudwatch_logs.3 (out_id=3)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"hc_period": schema.Int64Attribute{
								Description:         "The time period by second to count the error and retry failure data point",
								MarkdownDescription: "The time period by second to count the error and retry failure data point",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"hc_retry_failure_count": schema.Int64Attribute{
								Description:         "the retry failure count to meet the unhealthy requirement, this is a sum for all output plugins in a defined HC_Period, example for retry failure: [2022/02/16 20:11:36] [ warn] [engine] chunk '1-1645042288.260516436.flb' cannot be retried: task_id=0, input=tcp.3 > output=cloudwatch_logs.1",
								MarkdownDescription: "the retry failure count to meet the unhealthy requirement, this is a sum for all output plugins in a defined HC_Period, example for retry failure: [2022/02/16 20:11:36] [ warn] [engine] chunk '1-1645042288.260516436.flb' cannot be retried: task_id=0, input=tcp.3 > output=cloudwatch_logs.1",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"health_check": schema.BoolAttribute{
								Description:         "enable Health check feature at http://127.0.0.1:2020/api/v1/health Note: Enabling this will not automatically configure kubernetes to use fluentbit's healthcheck endpoint",
								MarkdownDescription: "enable Health check feature at http://127.0.0.1:2020/api/v1/health Note: Enabling this will not automatically configure kubernetes to use fluentbit's healthcheck endpoint",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_listen": schema.StringAttribute{
								Description:         "Address to listen",
								MarkdownDescription: "Address to listen",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}$`), ""),
								},
							},

							"http_port": schema.Int64Attribute{
								Description:         "Port to listen",
								MarkdownDescription: "Port to listen",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"http_server": schema.BoolAttribute{
								Description:         "If true enable statistics HTTP server",
								MarkdownDescription: "If true enable statistics HTTP server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_file": schema.StringAttribute{
								Description:         "File to log diagnostic output",
								MarkdownDescription: "File to log diagnostic output",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Diagnostic level (error/warning/info/debug/trace)",
								MarkdownDescription: "Diagnostic level (error/warning/info/debug/trace)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("off", "error", "warning", "info", "debug", "trace"),
								},
							},

							"parsers_file": schema.StringAttribute{
								Description:         "Optional 'parsers' config file (can be multiple)",
								MarkdownDescription: "Optional 'parsers' config file (can be multiple)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage": schema.SingleNestedAttribute{
								Description:         "Configure a global environment for the storage layer in Service. It is recommended to configure the volume and volumeMount separately for this storage. The hostPath type should be used for that Volume in Fluentbit daemon set.",
								MarkdownDescription: "Configure a global environment for the storage layer in Service. It is recommended to configure the volume and volumeMount separately for this storage. The hostPath type should be used for that Volume in Fluentbit daemon set.",
								Attributes: map[string]schema.Attribute{
									"backlog_mem_limit": schema.StringAttribute{
										Description:         "This option configure a hint of maximum value of memory to use when processing these records",
										MarkdownDescription: "This option configure a hint of maximum value of memory to use when processing these records",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"checksum": schema.StringAttribute{
										Description:         "Enable the data integrity check when writing and reading data from the filesystem",
										MarkdownDescription: "Enable the data integrity check when writing and reading data from the filesystem",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"delete_irrecoverable_chunks": schema.StringAttribute{
										Description:         "When enabled, irrecoverable chunks will be deleted during runtime, and any other irrecoverable chunk located in the configured storage path directory will be deleted when Fluent-Bit starts.",
										MarkdownDescription: "When enabled, irrecoverable chunks will be deleted during runtime, and any other irrecoverable chunk located in the configured storage path directory will be deleted when Fluent-Bit starts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"max_chunks_up": schema.Int64Attribute{
										Description:         "If the input plugin has enabled filesystem storage type, this property sets the maximum number of Chunks that can be up in memory",
										MarkdownDescription: "If the input plugin has enabled filesystem storage type, this property sets the maximum number of Chunks that can be up in memory",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.StringAttribute{
										Description:         "If http_server option has been enabled in the Service section, this option registers a new endpoint where internal metrics of the storage layer can be consumed",
										MarkdownDescription: "If http_server option has been enabled in the Service section, this option registers a new endpoint where internal metrics of the storage layer can be consumed",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("on", "off"),
										},
									},

									"path": schema.StringAttribute{
										Description:         "Select an optional location in the file system to store streams and chunks of data/",
										MarkdownDescription: "Select an optional location in the file system to store streams and chunks of data/",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sync": schema.StringAttribute{
										Description:         "Configure the synchronization mode used to store the data into the file system",
										MarkdownDescription: "Configure the synchronization mode used to store the data into the file system",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("normal", "full"),
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
		},
	}
}

func (r *FluentbitFluentIoClusterFluentBitConfigV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentbit_fluent_io_cluster_fluent_bit_config_v1alpha2_manifest")

	var model FluentbitFluentIoClusterFluentBitConfigV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentbit.fluent.io/v1alpha2")
	model.Kind = pointer.String("ClusterFluentBitConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
