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
	_ datasource.DataSource = &FluentbitFluentIoClusterInputV1Alpha2Manifest{}
)

func NewFluentbitFluentIoClusterInputV1Alpha2Manifest() datasource.DataSource {
	return &FluentbitFluentIoClusterInputV1Alpha2Manifest{}
}

type FluentbitFluentIoClusterInputV1Alpha2Manifest struct{}

type FluentbitFluentIoClusterInputV1Alpha2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Alias    *string `tfsdk:"alias" json:"alias,omitempty"`
		Collectd *struct {
			Listen  *string `tfsdk:"listen" json:"listen,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			TypesDB *string `tfsdk:"types_db" json:"typesDB,omitempty"`
		} `tfsdk:"collectd" json:"collectd,omitempty"`
		CustomPlugin *struct {
			Config     *string            `tfsdk:"config" json:"config,omitempty"`
			YamlConfig *map[string]string `tfsdk:"yaml_config" json:"yamlConfig,omitempty"`
		} `tfsdk:"custom_plugin" json:"customPlugin,omitempty"`
		Dummy *struct {
			Dummy   *string `tfsdk:"dummy" json:"dummy,omitempty"`
			Rate    *int64  `tfsdk:"rate" json:"rate,omitempty"`
			Samples *int64  `tfsdk:"samples" json:"samples,omitempty"`
			Tag     *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"dummy" json:"dummy,omitempty"`
		FluentBitMetrics *struct {
			ScrapeInterval *string `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
			ScrapeOnStart  *bool   `tfsdk:"scrape_on_start" json:"scrapeOnStart,omitempty"`
			Tag            *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"fluent_bit_metrics" json:"fluentBitMetrics,omitempty"`
		Forward *struct {
			BufferMaxSize   *string `tfsdk:"buffer_max_size" json:"bufferMaxSize,omitempty"`
			BufferchunkSize *string `tfsdk:"bufferchunk_size" json:"bufferchunkSize,omitempty"`
			Listen          *string `tfsdk:"listen" json:"listen,omitempty"`
			Port            *int64  `tfsdk:"port" json:"port,omitempty"`
			Tag             *string `tfsdk:"tag" json:"tag,omitempty"`
			TagPrefix       *string `tfsdk:"tag_prefix" json:"tagPrefix,omitempty"`
			Threaded        *string `tfsdk:"threaded" json:"threaded,omitempty"`
			UnixPath        *string `tfsdk:"unix_path" json:"unixPath,omitempty"`
			UnixPerm        *string `tfsdk:"unix_perm" json:"unixPerm,omitempty"`
		} `tfsdk:"forward" json:"forward,omitempty"`
		Http *struct {
			BufferChunkSize        *string `tfsdk:"buffer_chunk_size" json:"bufferChunkSize,omitempty"`
			BufferMaxSize          *string `tfsdk:"buffer_max_size" json:"bufferMaxSize,omitempty"`
			Listen                 *string `tfsdk:"listen" json:"listen,omitempty"`
			Port                   *int64  `tfsdk:"port" json:"port,omitempty"`
			SuccessfulHeader       *string `tfsdk:"successful_header" json:"successfulHeader,omitempty"`
			SuccessfulResponseCode *int64  `tfsdk:"successful_response_code" json:"successfulResponseCode,omitempty"`
			TagKey                 *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
			Tls                    *struct {
				CaFile      *string `tfsdk:"ca_file" json:"caFile,omitempty"`
				CaPath      *string `tfsdk:"ca_path" json:"caPath,omitempty"`
				CrtFile     *string `tfsdk:"crt_file" json:"crtFile,omitempty"`
				Debug       *int64  `tfsdk:"debug" json:"debug,omitempty"`
				KeyFile     *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				KeyPassword *struct {
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"key_password" json:"keyPassword,omitempty"`
				Verify *bool   `tfsdk:"verify" json:"verify,omitempty"`
				Vhost  *string `tfsdk:"vhost" json:"vhost,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
		KubernetesEvents *struct {
			Db                *string `tfsdk:"db" json:"db,omitempty"`
			DbSync            *string `tfsdk:"db_sync" json:"dbSync,omitempty"`
			IntervalNsec      *int64  `tfsdk:"interval_nsec" json:"intervalNsec,omitempty"`
			IntervalSec       *int64  `tfsdk:"interval_sec" json:"intervalSec,omitempty"`
			KubeCAFile        *string `tfsdk:"kube_ca_file" json:"kubeCAFile,omitempty"`
			KubeCAPath        *string `tfsdk:"kube_ca_path" json:"kubeCAPath,omitempty"`
			KubeNamespace     *string `tfsdk:"kube_namespace" json:"kubeNamespace,omitempty"`
			KubeRequestLimit  *int64  `tfsdk:"kube_request_limit" json:"kubeRequestLimit,omitempty"`
			KubeRetentionTime *string `tfsdk:"kube_retention_time" json:"kubeRetentionTime,omitempty"`
			KubeTokenFile     *string `tfsdk:"kube_token_file" json:"kubeTokenFile,omitempty"`
			KubeTokenTTL      *string `tfsdk:"kube_token_ttl" json:"kubeTokenTTL,omitempty"`
			KubeURL           *string `tfsdk:"kube_url" json:"kubeURL,omitempty"`
			Tag               *string `tfsdk:"tag" json:"tag,omitempty"`
			TlsDebug          *int64  `tfsdk:"tls_debug" json:"tlsDebug,omitempty"`
			TlsVerify         *bool   `tfsdk:"tls_verify" json:"tlsVerify,omitempty"`
			TlsVhost          *string `tfsdk:"tls_vhost" json:"tlsVhost,omitempty"`
		} `tfsdk:"kubernetes_events" json:"kubernetesEvents,omitempty"`
		LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		Mqtt     *struct {
			Listen *string `tfsdk:"listen" json:"listen,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"mqtt" json:"mqtt,omitempty"`
		Nginx *struct {
			Host      *string `tfsdk:"host" json:"host,omitempty"`
			NginxPlus *bool   `tfsdk:"nginx_plus" json:"nginxPlus,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			StatusURL *string `tfsdk:"status_url" json:"statusURL,omitempty"`
		} `tfsdk:"nginx" json:"nginx,omitempty"`
		NodeExporterMetrics *struct {
			Path *struct {
				Procfs *string `tfsdk:"procfs" json:"procfs,omitempty"`
				Sysfs  *string `tfsdk:"sysfs" json:"sysfs,omitempty"`
			} `tfsdk:"path" json:"path,omitempty"`
			ScrapeInterval *string `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
			Tag            *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"node_exporter_metrics" json:"nodeExporterMetrics,omitempty"`
		OpenTelemetry *struct {
			BufferChunkSize        *string `tfsdk:"buffer_chunk_size" json:"bufferChunkSize,omitempty"`
			BufferMaxSize          *string `tfsdk:"buffer_max_size" json:"bufferMaxSize,omitempty"`
			Listen                 *string `tfsdk:"listen" json:"listen,omitempty"`
			Port                   *int64  `tfsdk:"port" json:"port,omitempty"`
			RawTraces              *bool   `tfsdk:"raw_traces" json:"rawTraces,omitempty"`
			SuccessfulResponseCode *int64  `tfsdk:"successful_response_code" json:"successfulResponseCode,omitempty"`
			TagKey                 *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
		} `tfsdk:"open_telemetry" json:"openTelemetry,omitempty"`
		Processors              *map[string]string `tfsdk:"processors" json:"processors,omitempty"`
		PrometheusScrapeMetrics *struct {
			Host           *string `tfsdk:"host" json:"host,omitempty"`
			MetricsPath    *string `tfsdk:"metrics_path" json:"metricsPath,omitempty"`
			Port           *int64  `tfsdk:"port" json:"port,omitempty"`
			ScrapeInterval *string `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
			Tag            *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"prometheus_scrape_metrics" json:"prometheusScrapeMetrics,omitempty"`
		Statsd *struct {
			Listen *string `tfsdk:"listen" json:"listen,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"statsd" json:"statsd,omitempty"`
		Syslog *struct {
			BufferChunkSize   *string `tfsdk:"buffer_chunk_size" json:"bufferChunkSize,omitempty"`
			BufferMaxSize     *string `tfsdk:"buffer_max_size" json:"bufferMaxSize,omitempty"`
			Listen            *string `tfsdk:"listen" json:"listen,omitempty"`
			Mode              *string `tfsdk:"mode" json:"mode,omitempty"`
			Parser            *string `tfsdk:"parser" json:"parser,omitempty"`
			Path              *string `tfsdk:"path" json:"path,omitempty"`
			Port              *int64  `tfsdk:"port" json:"port,omitempty"`
			ReceiveBufferSize *string `tfsdk:"receive_buffer_size" json:"receiveBufferSize,omitempty"`
			SourceAddressKey  *string `tfsdk:"source_address_key" json:"sourceAddressKey,omitempty"`
			UnixPerm          *int64  `tfsdk:"unix_perm" json:"unixPerm,omitempty"`
		} `tfsdk:"syslog" json:"syslog,omitempty"`
		Systemd *struct {
			Db                     *string   `tfsdk:"db" json:"db,omitempty"`
			DbSync                 *string   `tfsdk:"db_sync" json:"dbSync,omitempty"`
			MaxEntries             *int64    `tfsdk:"max_entries" json:"maxEntries,omitempty"`
			MaxFields              *int64    `tfsdk:"max_fields" json:"maxFields,omitempty"`
			Path                   *string   `tfsdk:"path" json:"path,omitempty"`
			PauseOnChunksOverlimit *string   `tfsdk:"pause_on_chunks_overlimit" json:"pauseOnChunksOverlimit,omitempty"`
			ReadFromTail           *string   `tfsdk:"read_from_tail" json:"readFromTail,omitempty"`
			StorageType            *string   `tfsdk:"storage_type" json:"storageType,omitempty"`
			StripUnderscores       *string   `tfsdk:"strip_underscores" json:"stripUnderscores,omitempty"`
			SystemdFilter          *[]string `tfsdk:"systemd_filter" json:"systemdFilter,omitempty"`
			SystemdFilterType      *string   `tfsdk:"systemd_filter_type" json:"systemdFilterType,omitempty"`
			Tag                    *string   `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"systemd" json:"systemd,omitempty"`
		Tail *struct {
			BufferChunkSize        *string   `tfsdk:"buffer_chunk_size" json:"bufferChunkSize,omitempty"`
			BufferMaxSize          *string   `tfsdk:"buffer_max_size" json:"bufferMaxSize,omitempty"`
			Db                     *string   `tfsdk:"db" json:"db,omitempty"`
			DbSync                 *string   `tfsdk:"db_sync" json:"dbSync,omitempty"`
			DisableInotifyWatcher  *bool     `tfsdk:"disable_inotify_watcher" json:"disableInotifyWatcher,omitempty"`
			DockerMode             *bool     `tfsdk:"docker_mode" json:"dockerMode,omitempty"`
			DockerModeFlushSeconds *int64    `tfsdk:"docker_mode_flush_seconds" json:"dockerModeFlushSeconds,omitempty"`
			DockerModeParser       *string   `tfsdk:"docker_mode_parser" json:"dockerModeParser,omitempty"`
			ExcludePath            *string   `tfsdk:"exclude_path" json:"excludePath,omitempty"`
			IgnoredOlder           *string   `tfsdk:"ignored_older" json:"ignoredOlder,omitempty"`
			Key                    *string   `tfsdk:"key" json:"key,omitempty"`
			MemBufLimit            *string   `tfsdk:"mem_buf_limit" json:"memBufLimit,omitempty"`
			Multiline              *bool     `tfsdk:"multiline" json:"multiline,omitempty"`
			MultilineFlushSeconds  *int64    `tfsdk:"multiline_flush_seconds" json:"multilineFlushSeconds,omitempty"`
			MultilineParser        *string   `tfsdk:"multiline_parser" json:"multilineParser,omitempty"`
			Parser                 *string   `tfsdk:"parser" json:"parser,omitempty"`
			ParserFirstline        *string   `tfsdk:"parser_firstline" json:"parserFirstline,omitempty"`
			ParserN                *[]string `tfsdk:"parser_n" json:"parserN,omitempty"`
			Path                   *string   `tfsdk:"path" json:"path,omitempty"`
			PathKey                *string   `tfsdk:"path_key" json:"pathKey,omitempty"`
			PauseOnChunksOverlimit *string   `tfsdk:"pause_on_chunks_overlimit" json:"pauseOnChunksOverlimit,omitempty"`
			ReadFromHead           *bool     `tfsdk:"read_from_head" json:"readFromHead,omitempty"`
			RefreshIntervalSeconds *int64    `tfsdk:"refresh_interval_seconds" json:"refreshIntervalSeconds,omitempty"`
			RotateWaitSeconds      *int64    `tfsdk:"rotate_wait_seconds" json:"rotateWaitSeconds,omitempty"`
			SkipLongLines          *bool     `tfsdk:"skip_long_lines" json:"skipLongLines,omitempty"`
			StorageType            *string   `tfsdk:"storage_type" json:"storageType,omitempty"`
			Tag                    *string   `tfsdk:"tag" json:"tag,omitempty"`
			TagRegex               *string   `tfsdk:"tag_regex" json:"tagRegex,omitempty"`
		} `tfsdk:"tail" json:"tail,omitempty"`
		Tcp *struct {
			BufferSize *string `tfsdk:"buffer_size" json:"bufferSize,omitempty"`
			ChunkSize  *string `tfsdk:"chunk_size" json:"chunkSize,omitempty"`
			Format     *string `tfsdk:"format" json:"format,omitempty"`
			Listen     *string `tfsdk:"listen" json:"listen,omitempty"`
			Port       *int64  `tfsdk:"port" json:"port,omitempty"`
			Separator  *string `tfsdk:"separator" json:"separator,omitempty"`
		} `tfsdk:"tcp" json:"tcp,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentbitFluentIoClusterInputV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentbit_fluent_io_cluster_input_v1alpha2_manifest"
}

func (r *FluentbitFluentIoClusterInputV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterInput is the Schema for the inputs API",
		MarkdownDescription: "ClusterInput is the Schema for the inputs API",
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
				Description:         "InputSpec defines the desired state of ClusterInput",
				MarkdownDescription: "InputSpec defines the desired state of ClusterInput",
				Attributes: map[string]schema.Attribute{
					"alias": schema.StringAttribute{
						Description:         "A user friendly alias name for this input plugin.Used in metrics for distinction of each configured input.",
						MarkdownDescription: "A user friendly alias name for this input plugin.Used in metrics for distinction of each configured input.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"collectd": schema.SingleNestedAttribute{
						Description:         "Collectd defines the Collectd input plugin configuration",
						MarkdownDescription: "Collectd defines the Collectd input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"listen": schema.StringAttribute{
								Description:         "Set the address to listen to, default: 0.0.0.0",
								MarkdownDescription: "Set the address to listen to, default: 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Set the port to listen to, default: 25826",
								MarkdownDescription: "Set the port to listen to, default: 25826",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"types_db": schema.StringAttribute{
								Description:         "Set the data specification file,default: /usr/share/collectd/types.db",
								MarkdownDescription: "Set the data specification file,default: /usr/share/collectd/types.db",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_plugin": schema.SingleNestedAttribute{
						Description:         "CustomPlugin defines Custom Input configuration.",
						MarkdownDescription: "CustomPlugin defines Custom Input configuration.",
						Attributes: map[string]schema.Attribute{
							"config": schema.StringAttribute{
								Description:         "Config holds any unsupported plugins classic configurations,if ConfigFileFormat is set to yaml, this filed will be ignored",
								MarkdownDescription: "Config holds any unsupported plugins classic configurations,if ConfigFileFormat is set to yaml, this filed will be ignored",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"yaml_config": schema.MapAttribute{
								Description:         "YamlConfig holds the unsupported plugins yaml configurations, it only works when the ConfigFileFormat is yaml",
								MarkdownDescription: "YamlConfig holds the unsupported plugins yaml configurations, it only works when the ConfigFileFormat is yaml",
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

					"dummy": schema.SingleNestedAttribute{
						Description:         "Dummy defines Dummy Input configuration.",
						MarkdownDescription: "Dummy defines Dummy Input configuration.",
						Attributes: map[string]schema.Attribute{
							"dummy": schema.StringAttribute{
								Description:         "Dummy JSON record.",
								MarkdownDescription: "Dummy JSON record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate": schema.Int64Attribute{
								Description:         "Events number generated per second.",
								MarkdownDescription: "Events number generated per second.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"samples": schema.Int64Attribute{
								Description:         "Sample events to generate.",
								MarkdownDescription: "Sample events to generate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag name associated to all records comming from this plugin.",
								MarkdownDescription: "Tag name associated to all records comming from this plugin.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"fluent_bit_metrics": schema.SingleNestedAttribute{
						Description:         "FluentBitMetrics defines Fluent Bit Metrics Input configuration.",
						MarkdownDescription: "FluentBitMetrics defines Fluent Bit Metrics Input configuration.",
						Attributes: map[string]schema.Attribute{
							"scrape_interval": schema.StringAttribute{
								Description:         "The rate at which metrics are collected from the host operating system. default is 2 seconds.",
								MarkdownDescription: "The rate at which metrics are collected from the host operating system. default is 2 seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scrape_on_start": schema.BoolAttribute{
								Description:         "Scrape metrics upon start, useful to avoid waiting for 'scrape_interval' for the first round of metrics.",
								MarkdownDescription: "Scrape metrics upon start, useful to avoid waiting for 'scrape_interval' for the first round of metrics.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"forward": schema.SingleNestedAttribute{
						Description:         "Forward defines forward  input plugin configuration",
						MarkdownDescription: "Forward defines forward  input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"buffer_max_size": schema.StringAttribute{
								Description:         "Specify maximum buffer memory size used to recieve a forward message.The value must be according to the Unit Size specification.",
								MarkdownDescription: "Specify maximum buffer memory size used to recieve a forward message.The value must be according to the Unit Size specification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"bufferchunk_size": schema.StringAttribute{
								Description:         "Set the initial buffer size to store incoming data.This value is used too to increase buffer size as required.The value must be according to the Unit Size specification.",
								MarkdownDescription: "Set the initial buffer size to store incoming data.This value is used too to increase buffer size as required.The value must be according to the Unit Size specification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"listen": schema.StringAttribute{
								Description:         "Listener network interface.",
								MarkdownDescription: "Listener network interface.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port for forward plugin instance.",
								MarkdownDescription: "Port for forward plugin instance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"tag": schema.StringAttribute{
								Description:         "in_forward uses the tag value for incoming logs. If not set it uses tag from incoming log.",
								MarkdownDescription: "in_forward uses the tag value for incoming logs. If not set it uses tag from incoming log.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_prefix": schema.StringAttribute{
								Description:         "Adds the prefix to incoming event's tag",
								MarkdownDescription: "Adds the prefix to incoming event's tag",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"threaded": schema.StringAttribute{
								Description:         "Threaded mechanism allows input plugin to run in a separate thread which helps to desaturate the main pipeline.",
								MarkdownDescription: "Threaded mechanism allows input plugin to run in a separate thread which helps to desaturate the main pipeline.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unix_path": schema.StringAttribute{
								Description:         "Specify the path to unix socket to recieve a forward message. If set, Listen and port are ignnored.",
								MarkdownDescription: "Specify the path to unix socket to recieve a forward message. If set, Listen and port are ignnored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unix_perm": schema.StringAttribute{
								Description:         "Set the permission of unix socket file.",
								MarkdownDescription: "Set the permission of unix socket file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http": schema.SingleNestedAttribute{
						Description:         "HTTP defines the HTTP input plugin configuration",
						MarkdownDescription: "HTTP defines the HTTP input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"buffer_chunk_size": schema.StringAttribute{
								Description:         "This sets the chunk size for incoming incoming JSON messages.These chunks are then stored/managed in the space available by buffer_max_size,default 512K.",
								MarkdownDescription: "This sets the chunk size for incoming incoming JSON messages.These chunks are then stored/managed in the space available by buffer_max_size,default 512K.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"buffer_max_size": schema.StringAttribute{
								Description:         "Specify the maximum buffer size in KB to receive a JSON message,default 4M.",
								MarkdownDescription: "Specify the maximum buffer size in KB to receive a JSON message,default 4M.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"listen": schema.StringAttribute{
								Description:         "The address to listen on,default 0.0.0.0",
								MarkdownDescription: "The address to listen on,default 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port for Fluent Bit to listen on,default 9880",
								MarkdownDescription: "The port for Fluent Bit to listen on,default 9880",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"successful_header": schema.StringAttribute{
								Description:         "Add an HTTP header key/value pair on success. Multiple headers can be set. Example: X-Custom custom-answer.",
								MarkdownDescription: "Add an HTTP header key/value pair on success. Multiple headers can be set. Example: X-Custom custom-answer.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"successful_response_code": schema.Int64Attribute{
								Description:         "It allows to set successful response code. 200, 201 and 204 are supported,default 201.",
								MarkdownDescription: "It allows to set successful response code. 200, 201 and 204 are supported,default 201.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_key": schema.StringAttribute{
								Description:         "Specify the key name to overwrite a tag. If set, the tag will be overwritten by a value of the key.",
								MarkdownDescription: "Specify the key name to overwrite a tag. If set, the tag will be overwritten by a value of the key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								MarkdownDescription: "Fluent Bit provides integrated support for Transport Layer Security (TLS) and it predecessor Secure Sockets Layer (SSL) respectively.",
								Attributes: map[string]schema.Attribute{
									"ca_file": schema.StringAttribute{
										Description:         "Absolute path to CA certificate file",
										MarkdownDescription: "Absolute path to CA certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_path": schema.StringAttribute{
										Description:         "Absolute path to scan for certificate files",
										MarkdownDescription: "Absolute path to scan for certificate files",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crt_file": schema.StringAttribute{
										Description:         "Absolute path to Certificate file",
										MarkdownDescription: "Absolute path to Certificate file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.Int64Attribute{
										Description:         "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										MarkdownDescription: "Set TLS debug verbosity level.It accept the following values: 0 (No debug), 1 (Error), 2 (State change), 3 (Informational) and 4 Verbose",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.OneOf(0, 1, 2, 3, 4),
										},
									},

									"key_file": schema.StringAttribute{
										Description:         "Absolute path to private Key file",
										MarkdownDescription: "Absolute path to private Key file",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_password": schema.SingleNestedAttribute{
										Description:         "Optional password for tls.key_file file",
										MarkdownDescription: "Optional password for tls.key_file file",
										Attributes: map[string]schema.Attribute{
											"value_from": schema.SingleNestedAttribute{
												Description:         "ValueSource defines how to find a value's key.",
												MarkdownDescription: "ValueSource defines how to find a value's key.",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "Selects a key of a secret in the pod's namespace",
														MarkdownDescription: "Selects a key of a secret in the pod's namespace",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify": schema.BoolAttribute{
										Description:         "Force certificate validation",
										MarkdownDescription: "Force certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vhost": schema.StringAttribute{
										Description:         "Hostname to be used for TLS SNI extension",
										MarkdownDescription: "Hostname to be used for TLS SNI extension",
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

					"kubernetes_events": schema.SingleNestedAttribute{
						Description:         "KubernetesEvents defines the KubernetesEvents input plugin configuration",
						MarkdownDescription: "KubernetesEvents defines the KubernetesEvents input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"db": schema.StringAttribute{
								Description:         "Set a database file to keep track of recorded Kubernetes events",
								MarkdownDescription: "Set a database file to keep track of recorded Kubernetes events",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_sync": schema.StringAttribute{
								Description:         "Set a database sync method. values: extra, full, normal and off",
								MarkdownDescription: "Set a database sync method. values: extra, full, normal and off",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"interval_nsec": schema.Int64Attribute{
								Description:         "Set the polling interval for each channel (sub seconds: nanoseconds).",
								MarkdownDescription: "Set the polling interval for each channel (sub seconds: nanoseconds).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"interval_sec": schema.Int64Attribute{
								Description:         "Set the polling interval for each channel.",
								MarkdownDescription: "Set the polling interval for each channel.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_ca_file": schema.StringAttribute{
								Description:         "CA certificate file",
								MarkdownDescription: "CA certificate file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_ca_path": schema.StringAttribute{
								Description:         "Absolute path to scan for certificate files",
								MarkdownDescription: "Absolute path to scan for certificate files",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_namespace": schema.StringAttribute{
								Description:         "Kubernetes namespace to query events from. Gets events from all namespaces by default",
								MarkdownDescription: "Kubernetes namespace to query events from. Gets events from all namespaces by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_request_limit": schema.Int64Attribute{
								Description:         "kubernetes limit parameter for events query, no limit applied when set to 0.",
								MarkdownDescription: "kubernetes limit parameter for events query, no limit applied when set to 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_retention_time": schema.StringAttribute{
								Description:         "Kubernetes retention time for events.",
								MarkdownDescription: "Kubernetes retention time for events.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_token_file": schema.StringAttribute{
								Description:         "Token file",
								MarkdownDescription: "Token file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_token_ttl": schema.StringAttribute{
								Description:         "configurable 'time to live' for the K8s token. By default, it is set to 600 seconds.After this time, the token is reloaded from Kube_Token_File or the Kube_Token_Command.",
								MarkdownDescription: "configurable 'time to live' for the K8s token. By default, it is set to 600 seconds.After this time, the token is reloaded from Kube_Token_File or the Kube_Token_Command.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_url": schema.StringAttribute{
								Description:         "API Server end-point",
								MarkdownDescription: "API Server end-point",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag name associated to all records comming from this plugin.",
								MarkdownDescription: "Tag name associated to all records comming from this plugin.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_debug": schema.Int64Attribute{
								Description:         "Debug level between 0 (nothing) and 4 (every detail).",
								MarkdownDescription: "Debug level between 0 (nothing) and 4 (every detail).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_verify": schema.BoolAttribute{
								Description:         "When enabled, turns on certificate validation when connecting to the Kubernetes API server.",
								MarkdownDescription: "When enabled, turns on certificate validation when connecting to the Kubernetes API server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_vhost": schema.StringAttribute{
								Description:         "Set optional TLS virtual host.",
								MarkdownDescription: "Set optional TLS virtual host.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("off", "error", "warning", "info", "debug", "trace"),
						},
					},

					"mqtt": schema.SingleNestedAttribute{
						Description:         "MQTT defines the MQTT input plugin configuration",
						MarkdownDescription: "MQTT defines the MQTT input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"listen": schema.StringAttribute{
								Description:         "Listener network interface, default: 0.0.0.0",
								MarkdownDescription: "Listener network interface, default: 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port where listening for connections, default: 1883",
								MarkdownDescription: "TCP port where listening for connections, default: 1883",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"nginx": schema.SingleNestedAttribute{
						Description:         "Nginx defines the Nginx input plugin configuration",
						MarkdownDescription: "Nginx defines the Nginx input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Name of the target host or IP address to check, default: localhost",
								MarkdownDescription: "Name of the target host or IP address to check, default: localhost",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"nginx_plus": schema.BoolAttribute{
								Description:         "Turn on NGINX plus mode,default: true",
								MarkdownDescription: "Turn on NGINX plus mode,default: true",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port of the target nginx service to connect to, default: 80",
								MarkdownDescription: "Port of the target nginx service to connect to, default: 80",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"status_url": schema.StringAttribute{
								Description:         "The URL of the Stub Status Handler,default: /status",
								MarkdownDescription: "The URL of the Stub Status Handler,default: /status",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_exporter_metrics": schema.SingleNestedAttribute{
						Description:         "NodeExporterMetrics defines Node Exporter Metrics Input configuration.",
						MarkdownDescription: "NodeExporterMetrics defines Node Exporter Metrics Input configuration.",
						Attributes: map[string]schema.Attribute{
							"path": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"procfs": schema.StringAttribute{
										Description:         "The mount point used to collect process information and metrics.",
										MarkdownDescription: "The mount point used to collect process information and metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sysfs": schema.StringAttribute{
										Description:         "The path in the filesystem used to collect system metrics.",
										MarkdownDescription: "The path in the filesystem used to collect system metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"scrape_interval": schema.StringAttribute{
								Description:         "The rate at which metrics are collected from the host operating system, default is 5 seconds.",
								MarkdownDescription: "The rate at which metrics are collected from the host operating system, default is 5 seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag name associated to all records comming from this plugin.",
								MarkdownDescription: "Tag name associated to all records comming from this plugin.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"open_telemetry": schema.SingleNestedAttribute{
						Description:         "OpenTelemetry defines the OpenTelemetry input plugin configuration",
						MarkdownDescription: "OpenTelemetry defines the OpenTelemetry input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"buffer_chunk_size": schema.StringAttribute{
								Description:         "This sets the chunk size for incoming incoming JSON messages. These chunks are then stored/managed in the space available by buffer_max_size(default 512K).",
								MarkdownDescription: "This sets the chunk size for incoming incoming JSON messages. These chunks are then stored/managed in the space available by buffer_max_size(default 512K).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"buffer_max_size": schema.StringAttribute{
								Description:         "Specify the maximum buffer size in KB to receive a JSON message(default 4M).",
								MarkdownDescription: "Specify the maximum buffer size in KB to receive a JSON message(default 4M).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"listen": schema.StringAttribute{
								Description:         "The address to listen on,default 0.0.0.0",
								MarkdownDescription: "The address to listen on,default 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port for Fluent Bit to listen on.default 4318.",
								MarkdownDescription: "The port for Fluent Bit to listen on.default 4318.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"raw_traces": schema.BoolAttribute{
								Description:         "Route trace data as a log message(default false).",
								MarkdownDescription: "Route trace data as a log message(default false).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"successful_response_code": schema.Int64Attribute{
								Description:         "It allows to set successful response code. 200, 201 and 204 are supported(default 201).",
								MarkdownDescription: "It allows to set successful response code. 200, 201 and 204 are supported(default 201).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_key": schema.StringAttribute{
								Description:         "Specify the key name to overwrite a tag. If set, the tag will be overwritten by a value of the key.",
								MarkdownDescription: "Specify the key name to overwrite a tag. If set, the tag will be overwritten by a value of the key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"processors": schema.MapAttribute{
						Description:         "Processors defines the processors configuration",
						MarkdownDescription: "Processors defines the processors configuration",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_scrape_metrics": schema.SingleNestedAttribute{
						Description:         "PrometheusScrapeMetrics  defines Prometheus Scrape Metrics Input configuration.",
						MarkdownDescription: "PrometheusScrapeMetrics  defines Prometheus Scrape Metrics Input configuration.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The host of the prometheus metric endpoint that you want to scrape",
								MarkdownDescription: "The host of the prometheus metric endpoint that you want to scrape",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_path": schema.StringAttribute{
								Description:         "The metrics URI endpoint, that must start with a forward slash, deflaut: /metrics",
								MarkdownDescription: "The metrics URI endpoint, that must start with a forward slash, deflaut: /metrics",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port of the promethes metric endpoint that you want to scrape",
								MarkdownDescription: "The port of the promethes metric endpoint that you want to scrape",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"scrape_interval": schema.StringAttribute{
								Description:         "The interval to scrape metrics, default: 10s",
								MarkdownDescription: "The interval to scrape metrics, default: 10s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag name associated to all records comming from this plugin",
								MarkdownDescription: "Tag name associated to all records comming from this plugin",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"statsd": schema.SingleNestedAttribute{
						Description:         "StatsD defines the StatsD input plugin configuration",
						MarkdownDescription: "StatsD defines the StatsD input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"listen": schema.StringAttribute{
								Description:         "Listener network interface, default: 0.0.0.0",
								MarkdownDescription: "Listener network interface, default: 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "UDP port where listening for connections, default: 8125",
								MarkdownDescription: "UDP port where listening for connections, default: 8125",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"syslog": schema.SingleNestedAttribute{
						Description:         "Syslog defines the Syslog input plugin configuration",
						MarkdownDescription: "Syslog defines the Syslog input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"buffer_chunk_size": schema.StringAttribute{
								Description:         "By default the buffer to store the incoming Syslog messages, do not allocate the maximum memory allowed, instead it allocate memory when is required.The rounds of allocations are set by Buffer_Chunk_Size. If not set, Buffer_Chunk_Size is equal to 32000 bytes (32KB).",
								MarkdownDescription: "By default the buffer to store the incoming Syslog messages, do not allocate the maximum memory allowed, instead it allocate memory when is required.The rounds of allocations are set by Buffer_Chunk_Size. If not set, Buffer_Chunk_Size is equal to 32000 bytes (32KB).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"buffer_max_size": schema.StringAttribute{
								Description:         "Specify the maximum buffer size to receive a Syslog message. If not set, the default size will be the value of Buffer_Chunk_Size.",
								MarkdownDescription: "Specify the maximum buffer size to receive a Syslog message. If not set, the default size will be the value of Buffer_Chunk_Size.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"listen": schema.StringAttribute{
								Description:         "If Mode is set to tcp or udp, specify the network interface to bind, default: 0.0.0.0",
								MarkdownDescription: "If Mode is set to tcp or udp, specify the network interface to bind, default: 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "Defines transport protocol mode: unix_udp (UDP over Unix socket), unix_tcp (TCP over Unix socket), tcp or udp",
								MarkdownDescription: "Defines transport protocol mode: unix_udp (UDP over Unix socket), unix_tcp (TCP over Unix socket), tcp or udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("unix_udp", "unix_tcp", "tcp", "udp"),
								},
							},

							"parser": schema.StringAttribute{
								Description:         "Specify an alternative parser for the message. If Mode is set to tcp or udp then the default parser is syslog-rfc5424 otherwise syslog-rfc3164-local is used.If your syslog messages have fractional seconds set this Parser value to syslog-rfc5424 instead.",
								MarkdownDescription: "Specify an alternative parser for the message. If Mode is set to tcp or udp then the default parser is syslog-rfc5424 otherwise syslog-rfc3164-local is used.If your syslog messages have fractional seconds set this Parser value to syslog-rfc5424 instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "If Mode is set to unix_tcp or unix_udp, set the absolute path to the Unix socket file.",
								MarkdownDescription: "If Mode is set to unix_tcp or unix_udp, set the absolute path to the Unix socket file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "If Mode is set to tcp or udp, specify the TCP port to listen for incoming connections.",
								MarkdownDescription: "If Mode is set to tcp or udp, specify the TCP port to listen for incoming connections.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"receive_buffer_size": schema.StringAttribute{
								Description:         "Specify the maximum socket receive buffer size. If not set, the default value is OS-dependant,but generally too low to accept thousands of syslog messages per second without loss on udp or unix_udp sockets. Note that on Linux the value is capped by sysctl net.core.rmem_max.",
								MarkdownDescription: "Specify the maximum socket receive buffer size. If not set, the default value is OS-dependant,but generally too low to accept thousands of syslog messages per second without loss on udp or unix_udp sockets. Note that on Linux the value is capped by sysctl net.core.rmem_max.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"source_address_key": schema.StringAttribute{
								Description:         "Specify the key where the source address will be injected.",
								MarkdownDescription: "Specify the key where the source address will be injected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unix_perm": schema.Int64Attribute{
								Description:         "If Mode is set to unix_tcp or unix_udp, set the permission of the Unix socket file, default: 0644",
								MarkdownDescription: "If Mode is set to unix_tcp or unix_udp, set the permission of the Unix socket file, default: 0644",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"systemd": schema.SingleNestedAttribute{
						Description:         "Systemd defines Systemd Input configuration.",
						MarkdownDescription: "Systemd defines Systemd Input configuration.",
						Attributes: map[string]schema.Attribute{
							"db": schema.StringAttribute{
								Description:         "Specify the database file to keep track of monitored files and offsets.",
								MarkdownDescription: "Specify the database file to keep track of monitored files and offsets.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_sync": schema.StringAttribute{
								Description:         "Set a default synchronization (I/O) method. values: Extra, Full, Normal, Off.This flag affects how the internal SQLite engine do synchronization to disk,for more details about each option please refer to this section.note: this option was introduced on Fluent Bit v1.4.6.",
								MarkdownDescription: "Set a default synchronization (I/O) method. values: Extra, Full, Normal, Off.This flag affects how the internal SQLite engine do synchronization to disk,for more details about each option please refer to this section.note: this option was introduced on Fluent Bit v1.4.6.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Extra", "Full", "Normal", "Off"),
								},
							},

							"max_entries": schema.Int64Attribute{
								Description:         "When Fluent Bit starts, the Journal might have a high number of logs in the queue.In order to avoid delays and reduce memory usage, this option allows to specify the maximum number of log entries that can be processed per round.Once the limit is reached, Fluent Bit will continue processing the remaining log entries once Journald performs the notification.",
								MarkdownDescription: "When Fluent Bit starts, the Journal might have a high number of logs in the queue.In order to avoid delays and reduce memory usage, this option allows to specify the maximum number of log entries that can be processed per round.Once the limit is reached, Fluent Bit will continue processing the remaining log entries once Journald performs the notification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_fields": schema.Int64Attribute{
								Description:         "Set a maximum number of fields (keys) allowed per record.",
								MarkdownDescription: "Set a maximum number of fields (keys) allowed per record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Optional path to the Systemd journal directory,if not set, the plugin will use default paths to read local-only logs.",
								MarkdownDescription: "Optional path to the Systemd journal directory,if not set, the plugin will use default paths to read local-only logs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pause_on_chunks_overlimit": schema.StringAttribute{
								Description:         "Specifies if the input plugin should be paused (stop ingesting new data) when the storage.max_chunks_up value is reached.",
								MarkdownDescription: "Specifies if the input plugin should be paused (stop ingesting new data) when the storage.max_chunks_up value is reached.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"read_from_tail": schema.StringAttribute{
								Description:         "Start reading new entries. Skip entries already stored in Journald.",
								MarkdownDescription: "Start reading new entries. Skip entries already stored in Journald.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"storage_type": schema.StringAttribute{
								Description:         "Specify the buffering mechanism to use. It can be memory or filesystem",
								MarkdownDescription: "Specify the buffering mechanism to use. It can be memory or filesystem",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("filesystem", "memory"),
								},
							},

							"strip_underscores": schema.StringAttribute{
								Description:         "Remove the leading underscore of the Journald field (key). For example the Journald field _PID becomes the key PID.",
								MarkdownDescription: "Remove the leading underscore of the Journald field (key). For example the Journald field _PID becomes the key PID.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"systemd_filter": schema.ListAttribute{
								Description:         "Allows to perform a query over logs that contains a specific Journald key/value pairs, e.g: _SYSTEMD_UNIT=UNIT.The Systemd_Filter option can be specified multiple times in the input section to apply multiple filters as required.",
								MarkdownDescription: "Allows to perform a query over logs that contains a specific Journald key/value pairs, e.g: _SYSTEMD_UNIT=UNIT.The Systemd_Filter option can be specified multiple times in the input section to apply multiple filters as required.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"systemd_filter_type": schema.StringAttribute{
								Description:         "Define the filter type when Systemd_Filter is specified multiple times. Allowed values are And and Or.With And a record is matched only when all of the Systemd_Filter have a match.With Or a record is matched when any of the Systemd_Filter has a match.",
								MarkdownDescription: "Define the filter type when Systemd_Filter is specified multiple times. Allowed values are And and Or.With And a record is matched only when all of the Systemd_Filter have a match.With Or a record is matched when any of the Systemd_Filter has a match.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("And", "Or"),
								},
							},

							"tag": schema.StringAttribute{
								Description:         "The tag is used to route messages but on Systemd plugin there is an extra functionality:if the tag includes a star/wildcard, it will be expanded with the Systemd Unit file (e.g: host.* => host.UNIT_NAME).",
								MarkdownDescription: "The tag is used to route messages but on Systemd plugin there is an extra functionality:if the tag includes a star/wildcard, it will be expanded with the Systemd Unit file (e.g: host.* => host.UNIT_NAME).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tail": schema.SingleNestedAttribute{
						Description:         "Tail defines Tail Input configuration.",
						MarkdownDescription: "Tail defines Tail Input configuration.",
						Attributes: map[string]schema.Attribute{
							"buffer_chunk_size": schema.StringAttribute{
								Description:         "Set the initial buffer size to read files data.This value is used too to increase buffer size.The value must be according to the Unit Size specification.",
								MarkdownDescription: "Set the initial buffer size to read files data.This value is used too to increase buffer size.The value must be according to the Unit Size specification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"buffer_max_size": schema.StringAttribute{
								Description:         "Set the limit of the buffer size per monitored file.When a buffer needs to be increased (e.g: very long lines),this value is used to restrict how much the memory buffer can grow.If reading a file exceed this limit, the file is removed from the monitored file listThe value must be according to the Unit Size specification.",
								MarkdownDescription: "Set the limit of the buffer size per monitored file.When a buffer needs to be increased (e.g: very long lines),this value is used to restrict how much the memory buffer can grow.If reading a file exceed this limit, the file is removed from the monitored file listThe value must be according to the Unit Size specification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"db": schema.StringAttribute{
								Description:         "Specify the database file to keep track of monitored files and offsets.",
								MarkdownDescription: "Specify the database file to keep track of monitored files and offsets.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_sync": schema.StringAttribute{
								Description:         "Set a default synchronization (I/O) method. Values: Extra, Full, Normal, Off.",
								MarkdownDescription: "Set a default synchronization (I/O) method. Values: Extra, Full, Normal, Off.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Extra", "Full", "Normal", "Off"),
								},
							},

							"disable_inotify_watcher": schema.BoolAttribute{
								Description:         "DisableInotifyWatcher will disable inotify and use the file stat watcher instead.",
								MarkdownDescription: "DisableInotifyWatcher will disable inotify and use the file stat watcher instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"docker_mode": schema.BoolAttribute{
								Description:         "If enabled, the plugin will recombine split Docker log lines before passing them to any parser as configured above.This mode cannot be used at the same time as Multiline.",
								MarkdownDescription: "If enabled, the plugin will recombine split Docker log lines before passing them to any parser as configured above.This mode cannot be used at the same time as Multiline.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"docker_mode_flush_seconds": schema.Int64Attribute{
								Description:         "Wait period time in seconds to flush queued unfinished split lines.",
								MarkdownDescription: "Wait period time in seconds to flush queued unfinished split lines.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"docker_mode_parser": schema.StringAttribute{
								Description:         "Specify an optional parser for the first line of the docker multiline mode. The parser name to be specified must be registered in the parsers.conf file.",
								MarkdownDescription: "Specify an optional parser for the first line of the docker multiline mode. The parser name to be specified must be registered in the parsers.conf file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exclude_path": schema.StringAttribute{
								Description:         "Set one or multiple shell patterns separated by commas to exclude files matching a certain criteria,e.g: exclude_path=*.gz,*.zip",
								MarkdownDescription: "Set one or multiple shell patterns separated by commas to exclude files matching a certain criteria,e.g: exclude_path=*.gz,*.zip",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ignored_older": schema.StringAttribute{
								Description:         "Ignores records which are older than this time in seconds.Supports m,h,d (minutes, hours, days) syntax.Default behavior is to read all records from specified files.Only available when a Parser is specificied and it can parse the time of a record.",
								MarkdownDescription: "Ignores records which are older than this time in seconds.Supports m,h,d (minutes, hours, days) syntax.Default behavior is to read all records from specified files.Only available when a Parser is specificied and it can parse the time of a record.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(m|h|d)?$`), ""),
								},
							},

							"key": schema.StringAttribute{
								Description:         "When a message is unstructured (no parser applied), it's appended as a string under the key name log.This option allows to define an alternative name for that key.",
								MarkdownDescription: "When a message is unstructured (no parser applied), it's appended as a string under the key name log.This option allows to define an alternative name for that key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mem_buf_limit": schema.StringAttribute{
								Description:         "Set a limit of memory that Tail plugin can use when appending data to the Engine.If the limit is reach, it will be paused; when the data is flushed it resumes.",
								MarkdownDescription: "Set a limit of memory that Tail plugin can use when appending data to the Engine.If the limit is reach, it will be paused; when the data is flushed it resumes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"multiline": schema.BoolAttribute{
								Description:         "If enabled, the plugin will try to discover multiline messagesand use the proper parsers to compose the outgoing messages.Note that when this option is enabled the Parser option is not used.",
								MarkdownDescription: "If enabled, the plugin will try to discover multiline messagesand use the proper parsers to compose the outgoing messages.Note that when this option is enabled the Parser option is not used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"multiline_flush_seconds": schema.Int64Attribute{
								Description:         "Wait period time in seconds to process queued multiline messages",
								MarkdownDescription: "Wait period time in seconds to process queued multiline messages",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"multiline_parser": schema.StringAttribute{
								Description:         "This will help to reassembly multiline messages originally split by Docker or CRISpecify one or Multiline Parser definition to apply to the content.",
								MarkdownDescription: "This will help to reassembly multiline messages originally split by Docker or CRISpecify one or Multiline Parser definition to apply to the content.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parser": schema.StringAttribute{
								Description:         "Specify the name of a parser to interpret the entry as a structured message.",
								MarkdownDescription: "Specify the name of a parser to interpret the entry as a structured message.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parser_firstline": schema.StringAttribute{
								Description:         "Name of the parser that matchs the beginning of a multiline message.Note that the regular expression defined in the parser must include a group name (named capture)",
								MarkdownDescription: "Name of the parser that matchs the beginning of a multiline message.Note that the regular expression defined in the parser must include a group name (named capture)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parser_n": schema.ListAttribute{
								Description:         "Optional-extra parser to interpret and structure multiline entries.This option can be used to define multiple parsers.",
								MarkdownDescription: "Optional-extra parser to interpret and structure multiline entries.This option can be used to define multiple parsers.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Pattern specifying a specific log files or multiple ones through the use of common wildcards.",
								MarkdownDescription: "Pattern specifying a specific log files or multiple ones through the use of common wildcards.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path_key": schema.StringAttribute{
								Description:         "If enabled, it appends the name of the monitored file as part of the record.The value assigned becomes the key in the map.",
								MarkdownDescription: "If enabled, it appends the name of the monitored file as part of the record.The value assigned becomes the key in the map.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pause_on_chunks_overlimit": schema.StringAttribute{
								Description:         "Specifies if the input plugin should be paused (stop ingesting new data) when the storage.max_chunks_up value is reached.",
								MarkdownDescription: "Specifies if the input plugin should be paused (stop ingesting new data) when the storage.max_chunks_up value is reached.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off"),
								},
							},

							"read_from_head": schema.BoolAttribute{
								Description:         "For new discovered files on start (without a database offset/position),read the content from the head of the file, not tail.",
								MarkdownDescription: "For new discovered files on start (without a database offset/position),read the content from the head of the file, not tail.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"refresh_interval_seconds": schema.Int64Attribute{
								Description:         "The interval of refreshing the list of watched files in seconds.",
								MarkdownDescription: "The interval of refreshing the list of watched files in seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rotate_wait_seconds": schema.Int64Attribute{
								Description:         "Specify the number of extra time in seconds to monitor a file once is rotated in case some pending data is flushed.",
								MarkdownDescription: "Specify the number of extra time in seconds to monitor a file once is rotated in case some pending data is flushed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"skip_long_lines": schema.BoolAttribute{
								Description:         "When a monitored file reach it buffer capacity due to a very long line (Buffer_Max_Size),the default behavior is to stop monitoring that file.Skip_Long_Lines alter that behavior and instruct Fluent Bit to skip long linesand continue processing other lines that fits into the buffer size.",
								MarkdownDescription: "When a monitored file reach it buffer capacity due to a very long line (Buffer_Max_Size),the default behavior is to stop monitoring that file.Skip_Long_Lines alter that behavior and instruct Fluent Bit to skip long linesand continue processing other lines that fits into the buffer size.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_type": schema.StringAttribute{
								Description:         "Specify the buffering mechanism to use. It can be memory or filesystem",
								MarkdownDescription: "Specify the buffering mechanism to use. It can be memory or filesystem",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("filesystem", "memory"),
								},
							},

							"tag": schema.StringAttribute{
								Description:         "Set a tag (with regex-extract fields) that will be placed on lines read.E.g. kube.<namespace_name>.<pod_name>.<container_name>",
								MarkdownDescription: "Set a tag (with regex-extract fields) that will be placed on lines read.E.g. kube.<namespace_name>.<pod_name>.<container_name>",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag_regex": schema.StringAttribute{
								Description:         "Set a regex to exctract fields from the file",
								MarkdownDescription: "Set a regex to exctract fields from the file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tcp": schema.SingleNestedAttribute{
						Description:         "TCP defines the TCP input plugin configuration",
						MarkdownDescription: "TCP defines the TCP input plugin configuration",
						Attributes: map[string]schema.Attribute{
							"buffer_size": schema.StringAttribute{
								Description:         "Specify the maximum buffer size in KB to receive a JSON message. If not set, the default size will be the value of Chunk_Size.",
								MarkdownDescription: "Specify the maximum buffer size in KB to receive a JSON message. If not set, the default size will be the value of Chunk_Size.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"chunk_size": schema.StringAttribute{
								Description:         "By default the buffer to store the incoming JSON messages, do not allocate the maximum memory allowed, instead it allocate memory when is required.The rounds of allocations are set by Chunk_Size in KB. If not set, Chunk_Size is equal to 32 (32KB).",
								MarkdownDescription: "By default the buffer to store the incoming JSON messages, do not allocate the maximum memory allowed, instead it allocate memory when is required.The rounds of allocations are set by Chunk_Size in KB. If not set, Chunk_Size is equal to 32 (32KB).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(k|K|KB|kb|m|M|MB|mb|g|G|GB|gb)?$`), ""),
								},
							},

							"format": schema.StringAttribute{
								Description:         "Specify the expected payload format. It support the options json and none.When using json, it expects JSON maps, when is set to none, it will split every record using the defined Separator (option below).",
								MarkdownDescription: "Specify the expected payload format. It support the options json and none.When using json, it expects JSON maps, when is set to none, it will split every record using the defined Separator (option below).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"listen": schema.StringAttribute{
								Description:         "Listener network interface,default 0.0.0.0",
								MarkdownDescription: "Listener network interface,default 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "TCP port where listening for connections,default 5170",
								MarkdownDescription: "TCP port where listening for connections,default 5170",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"separator": schema.StringAttribute{
								Description:         "When the expected Format is set to none, Fluent Bit needs a separator string to split the records. By default it uses the breakline character (LF or 0x10).",
								MarkdownDescription: "When the expected Format is set to none, Fluent Bit needs a separator string to split the records. By default it uses the breakline character (LF or 0x10).",
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
	}
}

func (r *FluentbitFluentIoClusterInputV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentbit_fluent_io_cluster_input_v1alpha2_manifest")

	var model FluentbitFluentIoClusterInputV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentbit.fluent.io/v1alpha2")
	model.Kind = pointer.String("ClusterInput")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
