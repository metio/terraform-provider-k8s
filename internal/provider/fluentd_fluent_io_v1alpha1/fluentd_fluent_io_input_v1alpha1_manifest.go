/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentd_fluent_io_v1alpha1

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
	_ datasource.DataSource = &FluentdFluentIoInputV1Alpha1Manifest{}
)

func NewFluentdFluentIoInputV1Alpha1Manifest() datasource.DataSource {
	return &FluentdFluentIoInputV1Alpha1Manifest{}
}

type FluentdFluentIoInputV1Alpha1Manifest struct{}

type FluentdFluentIoInputV1Alpha1ManifestData struct {
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
		Inputs *[]struct {
			CustomPlugin *struct {
				Config *string `tfsdk:"config" json:"config,omitempty"`
			} `tfsdk:"custom_plugin" json:"customPlugin,omitempty"`
			Forward *struct {
				AddTagPrefix       *string `tfsdk:"add_tag_prefix" json:"addTagPrefix,omitempty"`
				Bind               *string `tfsdk:"bind" json:"bind,omitempty"`
				ChunkSizeLimit     *string `tfsdk:"chunk_size_limit" json:"chunkSizeLimit,omitempty"`
				ChunkSizeWarnLimit *string `tfsdk:"chunk_size_warn_limit" json:"chunkSizeWarnLimit,omitempty"`
				Client             *struct {
					Host      *string `tfsdk:"host" json:"host,omitempty"`
					Network   *string `tfsdk:"network" json:"network,omitempty"`
					SharedKey *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
					Users     *string `tfsdk:"users" json:"users,omitempty"`
				} `tfsdk:"client" json:"client,omitempty"`
				DenyKeepalive   *bool  `tfsdk:"deny_keepalive" json:"denyKeepalive,omitempty"`
				LingerTimeout   *int64 `tfsdk:"linger_timeout" json:"lingerTimeout,omitempty"`
				Port            *int64 `tfsdk:"port" json:"port,omitempty"`
				ResolveHostname *bool  `tfsdk:"resolve_hostname" json:"resolveHostname,omitempty"`
				Security        *struct {
					AllowAnonymousSource *string `tfsdk:"allow_anonymous_source" json:"allowAnonymousSource,omitempty"`
					SelfHostname         *string `tfsdk:"self_hostname" json:"selfHostname,omitempty"`
					SharedKey            *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
					User                 *struct {
						Password *struct {
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Username *struct {
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"user" json:"user,omitempty"`
					UserAuth *string `tfsdk:"user_auth" json:"userAuth,omitempty"`
				} `tfsdk:"security" json:"security,omitempty"`
				SendKeepalivePacket *bool   `tfsdk:"send_keepalive_packet" json:"sendKeepalivePacket,omitempty"`
				SkipInvalidEvent    *bool   `tfsdk:"skip_invalid_event" json:"skipInvalidEvent,omitempty"`
				SourceAddressKey    *string `tfsdk:"source_address_key" json:"sourceAddressKey,omitempty"`
				SourceHostnameKey   *string `tfsdk:"source_hostname_key" json:"sourceHostnameKey,omitempty"`
				Tag                 *string `tfsdk:"tag" json:"tag,omitempty"`
				Transport           *struct {
					CaCertPath             *string `tfsdk:"ca_cert_path" json:"caCertPath,omitempty"`
					CaPath                 *string `tfsdk:"ca_path" json:"caPath,omitempty"`
					CaPrivateKeyPassphrase *string `tfsdk:"ca_private_key_passphrase" json:"caPrivateKeyPassphrase,omitempty"`
					CaPrivateKeyPath       *string `tfsdk:"ca_private_key_path" json:"caPrivateKeyPath,omitempty"`
					CertPath               *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					CertVerifier           *string `tfsdk:"cert_verifier" json:"certVerifier,omitempty"`
					Ciphers                *string `tfsdk:"ciphers" json:"ciphers,omitempty"`
					ClientCertAuth         *bool   `tfsdk:"client_cert_auth" json:"clientCertAuth,omitempty"`
					Insecure               *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					PrivateKeyPassphrase   *string `tfsdk:"private_key_passphrase" json:"privateKeyPassphrase,omitempty"`
					PrivateKeyPath         *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
					Protocol               *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Version                *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"transport" json:"transport,omitempty"`
				User *struct {
					Password *struct {
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"password" json:"password,omitempty"`
					Username *struct {
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"forward" json:"forward,omitempty"`
			Http *struct {
				AddHttpHeaders       *bool   `tfsdk:"add_http_headers" json:"addHttpHeaders,omitempty"`
				AddRemoteAddr        *string `tfsdk:"add_remote_addr" json:"addRemoteAddr,omitempty"`
				Bind                 *string `tfsdk:"bind" json:"bind,omitempty"`
				BodySizeLimit        *string `tfsdk:"body_size_limit" json:"bodySizeLimit,omitempty"`
				CorsAllOrigins       *string `tfsdk:"cors_all_origins" json:"corsAllOrigins,omitempty"`
				CorsAllowCredentials *string `tfsdk:"cors_allow_credentials" json:"corsAllowCredentials,omitempty"`
				KeepaliveTimeout     *string `tfsdk:"keepalive_timeout" json:"keepaliveTimeout,omitempty"`
				Parse                *struct {
					CustomPatternPath    *string `tfsdk:"custom_pattern_path" json:"customPatternPath,omitempty"`
					EstimateCurrentEvent *bool   `tfsdk:"estimate_current_event" json:"estimateCurrentEvent,omitempty"`
					Expression           *string `tfsdk:"expression" json:"expression,omitempty"`
					Grok                 *[]struct {
						KeepTimeKey *bool   `tfsdk:"keep_time_key" json:"keepTimeKey,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						Pattern     *string `tfsdk:"pattern" json:"pattern,omitempty"`
						TimeFormat  *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
						TimeKey     *string `tfsdk:"time_key" json:"timeKey,omitempty"`
						TimeZone    *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
					} `tfsdk:"grok" json:"grok,omitempty"`
					GrokFailureKey       *string `tfsdk:"grok_failure_key" json:"grokFailureKey,omitempty"`
					GrokPattern          *string `tfsdk:"grok_pattern" json:"grokPattern,omitempty"`
					GrokPatternSeries    *string `tfsdk:"grok_pattern_series" json:"grokPatternSeries,omitempty"`
					Id                   *string `tfsdk:"id" json:"id,omitempty"`
					KeepTimeKey          *bool   `tfsdk:"keep_time_key" json:"keepTimeKey,omitempty"`
					Localtime            *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
					LogLevel             *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					MultiLineStartRegexp *string `tfsdk:"multi_line_start_regexp" json:"multiLineStartRegexp,omitempty"`
					TimeFormat           *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
					TimeFormatFallbacks  *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
					TimeKey              *string `tfsdk:"time_key" json:"timeKey,omitempty"`
					TimeType             *string `tfsdk:"time_type" json:"timeType,omitempty"`
					Timeout              *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Timezone             *string `tfsdk:"timezone" json:"timezone,omitempty"`
					Type                 *string `tfsdk:"type" json:"type,omitempty"`
					Types                *string `tfsdk:"types" json:"types,omitempty"`
					Utc                  *bool   `tfsdk:"utc" json:"utc,omitempty"`
				} `tfsdk:"parse" json:"parse,omitempty"`
				Port                 *int64 `tfsdk:"port" json:"port,omitempty"`
				RespondsWithEmptyImg *bool  `tfsdk:"responds_with_empty_img" json:"respondsWithEmptyImg,omitempty"`
				Transport            *struct {
					CaCertPath             *string `tfsdk:"ca_cert_path" json:"caCertPath,omitempty"`
					CaPath                 *string `tfsdk:"ca_path" json:"caPath,omitempty"`
					CaPrivateKeyPassphrase *string `tfsdk:"ca_private_key_passphrase" json:"caPrivateKeyPassphrase,omitempty"`
					CaPrivateKeyPath       *string `tfsdk:"ca_private_key_path" json:"caPrivateKeyPath,omitempty"`
					CertPath               *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					CertVerifier           *string `tfsdk:"cert_verifier" json:"certVerifier,omitempty"`
					Ciphers                *string `tfsdk:"ciphers" json:"ciphers,omitempty"`
					ClientCertAuth         *bool   `tfsdk:"client_cert_auth" json:"clientCertAuth,omitempty"`
					Insecure               *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					PrivateKeyPassphrase   *string `tfsdk:"private_key_passphrase" json:"privateKeyPassphrase,omitempty"`
					PrivateKeyPath         *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
					Protocol               *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Version                *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"transport" json:"transport,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Id           *string `tfsdk:"id" json:"id,omitempty"`
			Label        *string `tfsdk:"label" json:"label,omitempty"`
			LogLevel     *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			MonitorAgent *struct {
				Bind          *string `tfsdk:"bind" json:"bind,omitempty"`
				EmitInterval  *int64  `tfsdk:"emit_interval" json:"emitInterval,omitempty"`
				IncludeConfig *bool   `tfsdk:"include_config" json:"includeConfig,omitempty"`
				IncludeRetry  *bool   `tfsdk:"include_retry" json:"includeRetry,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Tag           *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"monitor_agent" json:"monitorAgent,omitempty"`
			Sample *struct {
				AutoIncrementKey *string `tfsdk:"auto_increment_key" json:"autoIncrementKey,omitempty"`
				Rate             *int64  `tfsdk:"rate" json:"rate,omitempty"`
				Sample           *string `tfsdk:"sample" json:"sample,omitempty"`
				Size             *int64  `tfsdk:"size" json:"size,omitempty"`
				Tag              *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"sample" json:"sample,omitempty"`
			Tail *struct {
				EmitUnmatchedLines *bool     `tfsdk:"emit_unmatched_lines" json:"emitUnmatchedLines,omitempty"`
				EnableStatWatcher  *bool     `tfsdk:"enable_stat_watcher" json:"enableStatWatcher,omitempty"`
				EnableWatchTimer   *bool     `tfsdk:"enable_watch_timer" json:"enableWatchTimer,omitempty"`
				Encoding           *string   `tfsdk:"encoding" json:"encoding,omitempty"`
				ExcludePath        *[]string `tfsdk:"exclude_path" json:"excludePath,omitempty"`
				FollowInodes       *bool     `tfsdk:"follow_inodes" json:"followInodes,omitempty"`
				FromEncoding       *string   `tfsdk:"from_encoding" json:"fromEncoding,omitempty"`
				Group              *struct {
					Pattern    *string `tfsdk:"pattern" json:"pattern,omitempty"`
					RatePeriod *int64  `tfsdk:"rate_period" json:"ratePeriod,omitempty"`
					Rule       *struct {
						Limit *int64             `tfsdk:"limit" json:"limit,omitempty"`
						Match *map[string]string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"rule" json:"rule,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				IgnoreRepeatedPermissionError *bool  `tfsdk:"ignore_repeated_permission_error" json:"ignoreRepeatedPermissionError,omitempty"`
				LimitRecentlyModified         *int64 `tfsdk:"limit_recently_modified" json:"limitRecentlyModified,omitempty"`
				MaxLineSize                   *int64 `tfsdk:"max_line_size" json:"maxLineSize,omitempty"`
				MultilineFlushInterval        *int64 `tfsdk:"multiline_flush_interval" json:"multilineFlushInterval,omitempty"`
				OpenOnEveryUpdate             *bool  `tfsdk:"open_on_every_update" json:"openOnEveryUpdate,omitempty"`
				Parse                         *struct {
					CustomPatternPath    *string `tfsdk:"custom_pattern_path" json:"customPatternPath,omitempty"`
					EstimateCurrentEvent *bool   `tfsdk:"estimate_current_event" json:"estimateCurrentEvent,omitempty"`
					Expression           *string `tfsdk:"expression" json:"expression,omitempty"`
					Grok                 *[]struct {
						KeepTimeKey *bool   `tfsdk:"keep_time_key" json:"keepTimeKey,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						Pattern     *string `tfsdk:"pattern" json:"pattern,omitempty"`
						TimeFormat  *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
						TimeKey     *string `tfsdk:"time_key" json:"timeKey,omitempty"`
						TimeZone    *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
					} `tfsdk:"grok" json:"grok,omitempty"`
					GrokFailureKey       *string `tfsdk:"grok_failure_key" json:"grokFailureKey,omitempty"`
					GrokPattern          *string `tfsdk:"grok_pattern" json:"grokPattern,omitempty"`
					GrokPatternSeries    *string `tfsdk:"grok_pattern_series" json:"grokPatternSeries,omitempty"`
					Id                   *string `tfsdk:"id" json:"id,omitempty"`
					KeepTimeKey          *bool   `tfsdk:"keep_time_key" json:"keepTimeKey,omitempty"`
					Localtime            *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
					LogLevel             *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					MultiLineStartRegexp *string `tfsdk:"multi_line_start_regexp" json:"multiLineStartRegexp,omitempty"`
					TimeFormat           *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
					TimeFormatFallbacks  *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
					TimeKey              *string `tfsdk:"time_key" json:"timeKey,omitempty"`
					TimeType             *string `tfsdk:"time_type" json:"timeType,omitempty"`
					Timeout              *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Timezone             *string `tfsdk:"timezone" json:"timezone,omitempty"`
					Type                 *string `tfsdk:"type" json:"type,omitempty"`
					Types                *string `tfsdk:"types" json:"types,omitempty"`
					Utc                  *bool   `tfsdk:"utc" json:"utc,omitempty"`
				} `tfsdk:"parse" json:"parse,omitempty"`
				Path                      *string `tfsdk:"path" json:"path,omitempty"`
				PathKey                   *string `tfsdk:"path_key" json:"pathKey,omitempty"`
				PathTimezone              *string `tfsdk:"path_timezone" json:"pathTimezone,omitempty"`
				PosFile                   *string `tfsdk:"pos_file" json:"posFile,omitempty"`
				PosFileCompactionInterval *int64  `tfsdk:"pos_file_compaction_interval" json:"posFileCompactionInterval,omitempty"`
				ReadBytesLimitPerSecond   *int64  `tfsdk:"read_bytes_limit_per_second" json:"readBytesLimitPerSecond,omitempty"`
				ReadFromHead              *bool   `tfsdk:"read_from_head" json:"readFromHead,omitempty"`
				ReadLinesLimit            *int64  `tfsdk:"read_lines_limit" json:"readLinesLimit,omitempty"`
				RefreshInterval           *int64  `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
				RotateWait                *int64  `tfsdk:"rotate_wait" json:"rotateWait,omitempty"`
				SkipRefreshOnStartup      *bool   `tfsdk:"skip_refresh_on_startup" json:"skipRefreshOnStartup,omitempty"`
				Tag                       *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"tail" json:"tail,omitempty"`
		} `tfsdk:"inputs" json:"inputs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentdFluentIoInputV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentd_fluent_io_input_v1alpha1_manifest"
}

func (r *FluentdFluentIoInputV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Input is the Schema for the inputs API",
		MarkdownDescription: "Input is the Schema for the inputs API",
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
				Description:         "InputSpec defines the desired state of Input",
				MarkdownDescription: "InputSpec defines the desired state of Input",
				Attributes: map[string]schema.Attribute{
					"inputs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"custom_plugin": schema.SingleNestedAttribute{
									Description:         "Custom plugin type",
									MarkdownDescription: "Custom plugin type",
									Attributes: map[string]schema.Attribute{
										"config": schema.StringAttribute{
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

								"forward": schema.SingleNestedAttribute{
									Description:         "in_forward plugin",
									MarkdownDescription: "in_forward plugin",
									Attributes: map[string]schema.Attribute{
										"add_tag_prefix": schema.StringAttribute{
											Description:         "Adds the prefix to the incoming event's tag.",
											MarkdownDescription: "Adds the prefix to the incoming event's tag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bind": schema.StringAttribute{
											Description:         "The port to listen to, default is '0.0.0.0'",
											MarkdownDescription: "The port to listen to, default is '0.0.0.0'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"chunk_size_limit": schema.StringAttribute{
											Description:         "The size limit of the received chunk. If the chunk size is larger than this value, the received chunk is dropped.",
											MarkdownDescription: "The size limit of the received chunk. If the chunk size is larger than this value, the received chunk is dropped.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(KB|MB|GB|TB)$`), ""),
											},
										},

										"chunk_size_warn_limit": schema.StringAttribute{
											Description:         "The warning size limit of the received chunk. If the chunk size is larger than this value, a warning message will be sent.",
											MarkdownDescription: "The warning size limit of the received chunk. If the chunk size is larger than this value, a warning message will be sent.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(KB|MB|GB|TB)$`), ""),
											},
										},

										"client": schema.SingleNestedAttribute{
											Description:         "The security section of client plugin",
											MarkdownDescription: "The security section of client plugin",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "The IP address or hostname of the client. This is exclusive with Network.",
													MarkdownDescription: "The IP address or hostname of the client. This is exclusive with Network.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"network": schema.StringAttribute{
													Description:         "The network address specification. This is exclusive with Host.",
													MarkdownDescription: "The network address specification. This is exclusive with Host.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"shared_key": schema.StringAttribute{
													Description:         "The shared key per client.",
													MarkdownDescription: "The shared key per client.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"users": schema.StringAttribute{
													Description:         "The array of usernames.",
													MarkdownDescription: "The array of usernames.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"deny_keepalive": schema.BoolAttribute{
											Description:         "The connections will be disconnected right after receiving a message, if true.",
											MarkdownDescription: "The connections will be disconnected right after receiving a message, if true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"linger_timeout": schema.Int64Attribute{
											Description:         "The timeout used to set the linger option.",
											MarkdownDescription: "The timeout used to set the linger option.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port to listen to, default is 24224.",
											MarkdownDescription: "The port to listen to, default is 24224.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"resolve_hostname": schema.BoolAttribute{
											Description:         "Tries to resolve hostname from IP addresses or not.",
											MarkdownDescription: "Tries to resolve hostname from IP addresses or not.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"security": schema.SingleNestedAttribute{
											Description:         "The security section of forward plugin",
											MarkdownDescription: "The security section of forward plugin",
											Attributes: map[string]schema.Attribute{
												"allow_anonymous_source": schema.StringAttribute{
													Description:         "Allows the anonymous source. <client> sections are required, if disabled.",
													MarkdownDescription: "Allows the anonymous source. <client> sections are required, if disabled.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"self_hostname": schema.StringAttribute{
													Description:         "The hostname.",
													MarkdownDescription: "The hostname.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"shared_key": schema.StringAttribute{
													Description:         "The shared key for authentication.",
													MarkdownDescription: "The shared key for authentication.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user": schema.SingleNestedAttribute{
													Description:         "Defines user section directly.",
													MarkdownDescription: "Defines user section directly.",
													Attributes: map[string]schema.Attribute{
														"password": schema.SingleNestedAttribute{
															Description:         "Secret defines the key of a value.",
															MarkdownDescription: "Secret defines the key of a value.",
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

														"username": schema.SingleNestedAttribute{
															Description:         "Secret defines the key of a value.",
															MarkdownDescription: "Secret defines the key of a value.",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"user_auth": schema.StringAttribute{
													Description:         "If true, user-based authentication is used.",
													MarkdownDescription: "If true, user-based authentication is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"send_keepalive_packet": schema.BoolAttribute{
											Description:         "Enables the TCP keepalive for sockets.",
											MarkdownDescription: "Enables the TCP keepalive for sockets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"skip_invalid_event": schema.BoolAttribute{
											Description:         "Skips the invalid incoming event.",
											MarkdownDescription: "Skips the invalid incoming event.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_address_key": schema.StringAttribute{
											Description:         "The field name of the client's source address. If set, the client's address will be set to its key.",
											MarkdownDescription: "The field name of the client's source address. If set, the client's address will be set to its key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_hostname_key": schema.StringAttribute{
											Description:         "The field name of the client's hostname. If set, the client's hostname will be set to its key.",
											MarkdownDescription: "The field name of the client's hostname. If set, the client's hostname will be set to its key.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.StringAttribute{
											Description:         "in_forward uses incoming event's tag by default (See Protocol Section).If the tag parameter is set, its value is used instead.",
											MarkdownDescription: "in_forward uses incoming event's tag by default (See Protocol Section).If the tag parameter is set, its value is used instead.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"transport": schema.SingleNestedAttribute{
											Description:         "The transport section of forward plugin",
											MarkdownDescription: "The transport section of forward plugin",
											Attributes: map[string]schema.Attribute{
												"ca_cert_path": schema.StringAttribute{
													Description:         "for Cert generated",
													MarkdownDescription: "for Cert generated",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_path": schema.StringAttribute{
													Description:         "for Cert signed by public CA",
													MarkdownDescription: "for Cert signed by public CA",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_private_key_passphrase": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_private_key_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cert_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cert_verifier": schema.StringAttribute{
													Description:         "other parameters",
													MarkdownDescription: "other parameters",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ciphers": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_cert_auth": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"private_key_passphrase": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"private_key_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "The protocal name of this plugin, i.e: tls",
													MarkdownDescription: "The protocal name of this plugin, i.e: tls",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
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

										"user": schema.SingleNestedAttribute{
											Description:         "The security section of user plugin",
											MarkdownDescription: "The security section of user plugin",
											Attributes: map[string]schema.Attribute{
												"password": schema.SingleNestedAttribute{
													Description:         "Secret defines the key of a value.",
													MarkdownDescription: "Secret defines the key of a value.",
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

												"username": schema.SingleNestedAttribute{
													Description:         "Secret defines the key of a value.",
													MarkdownDescription: "Secret defines the key of a value.",
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

								"http": schema.SingleNestedAttribute{
									Description:         "in_http plugin",
									MarkdownDescription: "in_http plugin",
									Attributes: map[string]schema.Attribute{
										"add_http_headers": schema.BoolAttribute{
											Description:         "Adds HTTP_ prefix headers to the record.",
											MarkdownDescription: "Adds HTTP_ prefix headers to the record.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"add_remote_addr": schema.StringAttribute{
											Description:         "Adds REMOTE_ADDR field to the record. The value of REMOTE_ADDR is the client's address.i.e: X-Forwarded-For: host1, host2",
											MarkdownDescription: "Adds REMOTE_ADDR field to the record. The value of REMOTE_ADDR is the client's address.i.e: X-Forwarded-For: host1, host2",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"bind": schema.StringAttribute{
											Description:         "The port to listen to, default is '0.0.0.0'",
											MarkdownDescription: "The port to listen to, default is '0.0.0.0'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"body_size_limit": schema.StringAttribute{
											Description:         "The size limit of the POSTed element.",
											MarkdownDescription: "The size limit of the POSTed element.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(KB|MB|GB|TB)$`), ""),
											},
										},

										"cors_all_origins": schema.StringAttribute{
											Description:         "Whitelist domains for CORS.",
											MarkdownDescription: "Whitelist domains for CORS.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cors_allow_credentials": schema.StringAttribute{
											Description:         "Add Access-Control-Allow-Credentials header. It's needed when a request's credentials mode is include",
											MarkdownDescription: "Add Access-Control-Allow-Credentials header. It's needed when a request's credentials mode is include",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keepalive_timeout": schema.StringAttribute{
											Description:         "The timeout limit for keeping the connection alive.",
											MarkdownDescription: "The timeout limit for keeping the connection alive.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
											},
										},

										"parse": schema.SingleNestedAttribute{
											Description:         "The parse section of http plugin",
											MarkdownDescription: "The parse section of http plugin",
											Attributes: map[string]schema.Attribute{
												"custom_pattern_path": schema.StringAttribute{
													Description:         "Path to the file that includes custom grok patterns.",
													MarkdownDescription: "Path to the file that includes custom grok patterns.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"estimate_current_event": schema.BoolAttribute{
													Description:         "If true, use Fluent::Eventnow(current time) as a timestamp when time_key is specified.",
													MarkdownDescription: "If true, use Fluent::Eventnow(current time) as a timestamp when time_key is specified.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression": schema.StringAttribute{
													Description:         "Specifies the regular expression for matching logs. Regular expression also supports i and m suffix.",
													MarkdownDescription: "Specifies the regular expression for matching logs. Regular expression also supports i and m suffix.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok": schema.ListNestedAttribute{
													Description:         "Grok Sections",
													MarkdownDescription: "Grok Sections",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"keep_time_key": schema.BoolAttribute{
																Description:         "If true, keep time field in the record.",
																MarkdownDescription: "If true, keep time field in the record.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of this grok section.",
																MarkdownDescription: "The name of this grok section.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
																Description:         "The pattern of grok. Required parameter.",
																MarkdownDescription: "The pattern of grok. Required parameter.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_format": schema.StringAttribute{
																Description:         "Process value using specified format. This is available only when time_type is string",
																MarkdownDescription: "Process value using specified format. This is available only when time_type is string",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_key": schema.StringAttribute{
																Description:         "Specify time field for event time. If the event doesn't have this field, current time is used.",
																MarkdownDescription: "Specify time field for event time. If the event doesn't have this field, current time is used.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_zone": schema.StringAttribute{
																Description:         "Use specified timezone. one can parse/format the time value in the specified timezone.",
																MarkdownDescription: "Use specified timezone. one can parse/format the time value in the specified timezone.",
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

												"grok_failure_key": schema.StringAttribute{
													Description:         "The key has grok failure reason.",
													MarkdownDescription: "The key has grok failure reason.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_pattern": schema.StringAttribute{
													Description:         "The pattern of grok.",
													MarkdownDescription: "The pattern of grok.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_pattern_series": schema.StringAttribute{
													Description:         "Specify grok pattern series set.",
													MarkdownDescription: "Specify grok pattern series set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
													Description:         "The @id parameter specifies a unique name for the configuration.",
													MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"keep_time_key": schema.BoolAttribute{
													Description:         "If true, keep time field in th record.",
													MarkdownDescription: "If true, keep time field in th record.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"localtime": schema.BoolAttribute{
													Description:         "If true, uses local time.",
													MarkdownDescription: "If true, uses local time.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"log_level": schema.StringAttribute{
													Description:         "The @log_level parameter specifies the plugin-specific logging level",
													MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"multi_line_start_regexp": schema.StringAttribute{
													Description:         "The regexp to match beginning of multiline. This is only for 'multiline_grok'.",
													MarkdownDescription: "The regexp to match beginning of multiline. This is only for 'multiline_grok'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_format": schema.StringAttribute{
													Description:         "Process value according to the specified format. This is available only when time_type is string",
													MarkdownDescription: "Process value according to the specified format. This is available only when time_type is string",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_format_fallbacks": schema.StringAttribute{
													Description:         "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
													MarkdownDescription: "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_key": schema.StringAttribute{
													Description:         "Specify time field for event time. If the event doesn't have this field, current time is used.",
													MarkdownDescription: "Specify time field for event time. If the event doesn't have this field, current time is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_type": schema.StringAttribute{
													Description:         "parses/formats value according to this type, default is string",
													MarkdownDescription: "parses/formats value according to this type, default is string",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("float", "unixtime", "string", "mixed"),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "Specify timeout for parse processing.",
													MarkdownDescription: "Specify timeout for parse processing.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
													},
												},

												"timezone": schema.StringAttribute{
													Description:         "Uses the specified timezone.",
													MarkdownDescription: "Uses the specified timezone.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "The @type parameter specifies the type of the plugin.",
													MarkdownDescription: "The @type parameter specifies the type of the plugin.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("regexp", "apache2", "apache_error", "nginx", "syslog", "csv", "tsv", "ltsv", "json", "multiline", "none", "grok", "multiline_grok"),
													},
												},

												"types": schema.StringAttribute{
													Description:         "Specify types for converting field into another, i.e: types user_id:integer,paid:bool,paid_usd_amount:float",
													MarkdownDescription: "Specify types for converting field into another, i.e: types user_id:integer,paid:bool,paid_usd_amount:float",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"utc": schema.BoolAttribute{
													Description:         "If true, uses UTC.",
													MarkdownDescription: "If true, uses UTC.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port to listen to, default is 9880.",
											MarkdownDescription: "The port to listen to, default is 9880.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},

										"responds_with_empty_img": schema.BoolAttribute{
											Description:         "Responds with an empty GIF image of 1x1 pixel (rather than an empty string).",
											MarkdownDescription: "Responds with an empty GIF image of 1x1 pixel (rather than an empty string).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"transport": schema.SingleNestedAttribute{
											Description:         "The transport section of http plugin",
											MarkdownDescription: "The transport section of http plugin",
											Attributes: map[string]schema.Attribute{
												"ca_cert_path": schema.StringAttribute{
													Description:         "for Cert generated",
													MarkdownDescription: "for Cert generated",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_path": schema.StringAttribute{
													Description:         "for Cert signed by public CA",
													MarkdownDescription: "for Cert signed by public CA",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_private_key_passphrase": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_private_key_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cert_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cert_verifier": schema.StringAttribute{
													Description:         "other parameters",
													MarkdownDescription: "other parameters",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ciphers": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_cert_auth": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"private_key_passphrase": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"private_key_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "The protocal name of this plugin, i.e: tls",
													MarkdownDescription: "The protocal name of this plugin, i.e: tls",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"id": schema.StringAttribute{
									Description:         "The @id parameter specifies a unique name for the configuration.",
									MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"label": schema.StringAttribute{
									Description:         "The @label parameter is to route the input events to <label> sections.",
									MarkdownDescription: "The @label parameter is to route the input events to <label> sections.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"log_level": schema.StringAttribute{
									Description:         "The @log_level parameter specifies the plugin-specific logging level",
									MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"monitor_agent": schema.SingleNestedAttribute{
									Description:         "monitor_agent plugin",
									MarkdownDescription: "monitor_agent plugin",
									Attributes: map[string]schema.Attribute{
										"bind": schema.StringAttribute{
											Description:         "The bind address to listen to.",
											MarkdownDescription: "The bind address to listen to.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"emit_interval": schema.Int64Attribute{
											Description:         "The interval time between event emits. This will be used when 'tag' is configured.",
											MarkdownDescription: "The interval time between event emits. This will be used when 'tag' is configured.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_config": schema.BoolAttribute{
											Description:         "You can set this option to false to remove the config field from the response.",
											MarkdownDescription: "You can set this option to false to remove the config field from the response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_retry": schema.BoolAttribute{
											Description:         "You can set this option to false to remove the retry field from the response.",
											MarkdownDescription: "You can set this option to false to remove the retry field from the response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port to listen to.",
											MarkdownDescription: "The port to listen to.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.StringAttribute{
											Description:         "If you set this parameter, this plugin emits metrics as records.",
											MarkdownDescription: "If you set this parameter, this plugin emits metrics as records.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"sample": schema.SingleNestedAttribute{
									Description:         "in_sample plugin",
									MarkdownDescription: "in_sample plugin",
									Attributes: map[string]schema.Attribute{
										"auto_increment_key": schema.StringAttribute{
											Description:         "If specified, each generated event has an auto-incremented key field.",
											MarkdownDescription: "If specified, each generated event has an auto-incremented key field.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rate": schema.Int64Attribute{
											Description:         "It configures how many events to generate per second.",
											MarkdownDescription: "It configures how many events to generate per second.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sample": schema.StringAttribute{
											Description:         "The sample data to be generated. It should be either an array of JSON hashes or a single JSON hash. If it is an array of JSON hashes, the hashes in the array are cycled through in order.",
											MarkdownDescription: "The sample data to be generated. It should be either an array of JSON hashes or a single JSON hash. If it is an array of JSON hashes, the hashes in the array are cycled through in order.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size": schema.Int64Attribute{
											Description:         "The number of events in the event stream of each emit.",
											MarkdownDescription: "The number of events in the event stream of each emit.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.StringAttribute{
											Description:         "The tag of the event. The value is the tag assigned to the generated events.",
											MarkdownDescription: "The tag of the event. The value is the tag assigned to the generated events.",
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
									Description:         "in_tail plugin",
									MarkdownDescription: "in_tail plugin",
									Attributes: map[string]schema.Attribute{
										"emit_unmatched_lines": schema.BoolAttribute{
											Description:         "Emits unmatched lines when <parse> format is not matched for incoming logs.",
											MarkdownDescription: "Emits unmatched lines when <parse> format is not matched for incoming logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_stat_watcher": schema.BoolAttribute{
											Description:         "Enables the additional inotify-based watcher. Setting this parameter to false will disable the inotify events and use only timer watcher for file tailing.This option is mainly for avoiding the stuck issue with inotify.",
											MarkdownDescription: "Enables the additional inotify-based watcher. Setting this parameter to false will disable the inotify events and use only timer watcher for file tailing.This option is mainly for avoiding the stuck issue with inotify.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_watch_timer": schema.BoolAttribute{
											Description:         "Enables the additional watch timer. Setting this parameter to false will significantly reduce CPU and I/O consumption when tailing a large number of files on systems with inotify support.The default is true which results in an additional 1 second timer being used.",
											MarkdownDescription: "Enables the additional watch timer. Setting this parameter to false will significantly reduce CPU and I/O consumption when tailing a large number of files on systems with inotify support.The default is true which results in an additional 1 second timer being used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"encoding": schema.StringAttribute{
											Description:         "Specifies the encoding of reading lines. By default, in_tail emits string value as ASCII-8BIT encoding.If encoding is specified, in_tail changes string to encoding.If encoding and fromEncoding both are specified, in_tail tries to encode string from fromEncoding to encoding.",
											MarkdownDescription: "Specifies the encoding of reading lines. By default, in_tail emits string value as ASCII-8BIT encoding.If encoding is specified, in_tail changes string to encoding.If encoding and fromEncoding both are specified, in_tail tries to encode string from fromEncoding to encoding.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude_path": schema.ListAttribute{
											Description:         "The paths excluded from the watcher list.",
											MarkdownDescription: "The paths excluded from the watcher list.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"follow_inodes": schema.BoolAttribute{
											Description:         "Avoid to read rotated files duplicately. You should set true when you use * or strftime format in path.",
											MarkdownDescription: "Avoid to read rotated files duplicately. You should set true when you use * or strftime format in path.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"from_encoding": schema.StringAttribute{
											Description:         "Specifies the encoding of reading lines. By default, in_tail emits string value as ASCII-8BIT encoding.If encoding is specified, in_tail changes string to encoding.If encoding and fromEncoding both are specified, in_tail tries to encode string from fromEncoding to encoding.",
											MarkdownDescription: "Specifies the encoding of reading lines. By default, in_tail emits string value as ASCII-8BIT encoding.If encoding is specified, in_tail changes string to encoding.If encoding and fromEncoding both are specified, in_tail tries to encode string from fromEncoding to encoding.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group": schema.SingleNestedAttribute{
											Description:         "The in_tail plugin can assign each log file to a group, based on user defined rules.The limit parameter controls the total number of lines collected for a group within a rate_period time interval.",
											MarkdownDescription: "The in_tail plugin can assign each log file to a group, based on user defined rules.The limit parameter controls the total number of lines collected for a group within a rate_period time interval.",
											Attributes: map[string]schema.Attribute{
												"pattern": schema.StringAttribute{
													Description:         "Specifies the regular expression for extracting metadata (namespace, podname) from log file path.Default value of the pattern regexp extracts information about namespace, podname, docker_id, container of the log (K8s specific).",
													MarkdownDescription: "Specifies the regular expression for extracting metadata (namespace, podname) from log file path.Default value of the pattern regexp extracts information about namespace, podname, docker_id, container of the log (K8s specific).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rate_period": schema.Int64Attribute{
													Description:         "Time period in which the group line limit is applied. in_tail resets the counter after every rate_period interval.",
													MarkdownDescription: "Time period in which the group line limit is applied. in_tail resets the counter after every rate_period interval.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rule": schema.SingleNestedAttribute{
													Description:         "Grouping rules for log files.",
													MarkdownDescription: "Grouping rules for log files.",
													Attributes: map[string]schema.Attribute{
														"limit": schema.Int64Attribute{
															Description:         "Maximum number of lines allowed from a group in rate_period time interval. The default value of -1 doesn't throttle log files of that group.",
															MarkdownDescription: "Maximum number of lines allowed from a group in rate_period time interval. The default value of -1 doesn't throttle log files of that group.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"match": schema.MapAttribute{
															Description:         "match parameter is used to check if a file belongs to a particular group based on hash keys (named captures from pattern) and hash values (regexp in string)",
															MarkdownDescription: "match parameter is used to check if a file belongs to a particular group based on hash keys (named captures from pattern) and hash values (regexp in string)",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"ignore_repeated_permission_error": schema.BoolAttribute{
											Description:         "If you have to exclude the non-permission files from the watch list, set this parameter to true. It suppresses the repeated permission error logs.",
											MarkdownDescription: "If you have to exclude the non-permission files from the watch list, set this parameter to true. It suppresses the repeated permission error logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"limit_recently_modified": schema.Int64Attribute{
											Description:         "Limits the watching files that the modification time is within the specified time range when using * in path.",
											MarkdownDescription: "Limits the watching files that the modification time is within the specified time range when using * in path.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_line_size": schema.Int64Attribute{
											Description:         "The maximum length of a line. Longer lines than it will be just skipped.",
											MarkdownDescription: "The maximum length of a line. Longer lines than it will be just skipped.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"multiline_flush_interval": schema.Int64Attribute{
											Description:         "The interval of flushing the buffer for multiline format.",
											MarkdownDescription: "The interval of flushing the buffer for multiline format.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"open_on_every_update": schema.BoolAttribute{
											Description:         "Opens and closes the file on every update instead of leaving it open until it gets rotated.",
											MarkdownDescription: "Opens and closes the file on every update instead of leaving it open until it gets rotated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parse": schema.SingleNestedAttribute{
											Description:         "Parse defines various parameters for the parse plugin",
											MarkdownDescription: "Parse defines various parameters for the parse plugin",
											Attributes: map[string]schema.Attribute{
												"custom_pattern_path": schema.StringAttribute{
													Description:         "Path to the file that includes custom grok patterns.",
													MarkdownDescription: "Path to the file that includes custom grok patterns.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"estimate_current_event": schema.BoolAttribute{
													Description:         "If true, use Fluent::Eventnow(current time) as a timestamp when time_key is specified.",
													MarkdownDescription: "If true, use Fluent::Eventnow(current time) as a timestamp when time_key is specified.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression": schema.StringAttribute{
													Description:         "Specifies the regular expression for matching logs. Regular expression also supports i and m suffix.",
													MarkdownDescription: "Specifies the regular expression for matching logs. Regular expression also supports i and m suffix.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok": schema.ListNestedAttribute{
													Description:         "Grok Sections",
													MarkdownDescription: "Grok Sections",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"keep_time_key": schema.BoolAttribute{
																Description:         "If true, keep time field in the record.",
																MarkdownDescription: "If true, keep time field in the record.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of this grok section.",
																MarkdownDescription: "The name of this grok section.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
																Description:         "The pattern of grok. Required parameter.",
																MarkdownDescription: "The pattern of grok. Required parameter.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_format": schema.StringAttribute{
																Description:         "Process value using specified format. This is available only when time_type is string",
																MarkdownDescription: "Process value using specified format. This is available only when time_type is string",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_key": schema.StringAttribute{
																Description:         "Specify time field for event time. If the event doesn't have this field, current time is used.",
																MarkdownDescription: "Specify time field for event time. If the event doesn't have this field, current time is used.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_zone": schema.StringAttribute{
																Description:         "Use specified timezone. one can parse/format the time value in the specified timezone.",
																MarkdownDescription: "Use specified timezone. one can parse/format the time value in the specified timezone.",
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

												"grok_failure_key": schema.StringAttribute{
													Description:         "The key has grok failure reason.",
													MarkdownDescription: "The key has grok failure reason.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_pattern": schema.StringAttribute{
													Description:         "The pattern of grok.",
													MarkdownDescription: "The pattern of grok.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_pattern_series": schema.StringAttribute{
													Description:         "Specify grok pattern series set.",
													MarkdownDescription: "Specify grok pattern series set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
													Description:         "The @id parameter specifies a unique name for the configuration.",
													MarkdownDescription: "The @id parameter specifies a unique name for the configuration.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"keep_time_key": schema.BoolAttribute{
													Description:         "If true, keep time field in th record.",
													MarkdownDescription: "If true, keep time field in th record.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"localtime": schema.BoolAttribute{
													Description:         "If true, uses local time.",
													MarkdownDescription: "If true, uses local time.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"log_level": schema.StringAttribute{
													Description:         "The @log_level parameter specifies the plugin-specific logging level",
													MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"multi_line_start_regexp": schema.StringAttribute{
													Description:         "The regexp to match beginning of multiline. This is only for 'multiline_grok'.",
													MarkdownDescription: "The regexp to match beginning of multiline. This is only for 'multiline_grok'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_format": schema.StringAttribute{
													Description:         "Process value according to the specified format. This is available only when time_type is string",
													MarkdownDescription: "Process value according to the specified format. This is available only when time_type is string",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_format_fallbacks": schema.StringAttribute{
													Description:         "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
													MarkdownDescription: "Uses the specified time format as a fallback in the specified order. You can parse undetermined time format by using time_format_fallbacks. This options is enabled when time_type is mixed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_key": schema.StringAttribute{
													Description:         "Specify time field for event time. If the event doesn't have this field, current time is used.",
													MarkdownDescription: "Specify time field for event time. If the event doesn't have this field, current time is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_type": schema.StringAttribute{
													Description:         "parses/formats value according to this type, default is string",
													MarkdownDescription: "parses/formats value according to this type, default is string",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("float", "unixtime", "string", "mixed"),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "Specify timeout for parse processing.",
													MarkdownDescription: "Specify timeout for parse processing.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(\.[0-9]{0,2})?(s|m|h|d)?$`), ""),
													},
												},

												"timezone": schema.StringAttribute{
													Description:         "Uses the specified timezone.",
													MarkdownDescription: "Uses the specified timezone.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "The @type parameter specifies the type of the plugin.",
													MarkdownDescription: "The @type parameter specifies the type of the plugin.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("regexp", "apache2", "apache_error", "nginx", "syslog", "csv", "tsv", "ltsv", "json", "multiline", "none", "grok", "multiline_grok"),
													},
												},

												"types": schema.StringAttribute{
													Description:         "Specify types for converting field into another, i.e: types user_id:integer,paid:bool,paid_usd_amount:float",
													MarkdownDescription: "Specify types for converting field into another, i.e: types user_id:integer,paid:bool,paid_usd_amount:float",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"utc": schema.BoolAttribute{
													Description:         "If true, uses UTC.",
													MarkdownDescription: "If true, uses UTC.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "The path(s) to read. Multiple paths can be specified, separated by comma ','.",
											MarkdownDescription: "The path(s) to read. Multiple paths can be specified, separated by comma ','.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path_key": schema.StringAttribute{
											Description:         "Adds the watching file path to the path_key field.",
											MarkdownDescription: "Adds the watching file path to the path_key field.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path_timezone": schema.StringAttribute{
											Description:         "This parameter is for strftime formatted path like /path/to/%Y/%m/%d/.",
											MarkdownDescription: "This parameter is for strftime formatted path like /path/to/%Y/%m/%d/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pos_file": schema.StringAttribute{
											Description:         "(recommended) Fluentd will record the position it last read from this file.pos_file handles multiple positions in one file so no need to have multiple pos_file parameters per source.Don't share pos_file between in_tail configurations. It causes unexpected behavior e.g. corrupt pos_file content.",
											MarkdownDescription: "(recommended) Fluentd will record the position it last read from this file.pos_file handles multiple positions in one file so no need to have multiple pos_file parameters per source.Don't share pos_file between in_tail configurations. It causes unexpected behavior e.g. corrupt pos_file content.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pos_file_compaction_interval": schema.Int64Attribute{
											Description:         "The interval of doing compaction of pos file.",
											MarkdownDescription: "The interval of doing compaction of pos file.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_bytes_limit_per_second": schema.Int64Attribute{
											Description:         "The number of reading bytes per second to read with I/O operation. This value should be equal or greater than 8192.",
											MarkdownDescription: "The number of reading bytes per second to read with I/O operation. This value should be equal or greater than 8192.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_from_head": schema.BoolAttribute{
											Description:         "Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.",
											MarkdownDescription: "Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read_lines_limit": schema.Int64Attribute{
											Description:         "The number of lines to read with each I/O operation.",
											MarkdownDescription: "The number of lines to read with each I/O operation.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"refresh_interval": schema.Int64Attribute{
											Description:         "The interval to refresh the list of watch files. This is used when the path includes *.",
											MarkdownDescription: "The interval to refresh the list of watch files. This is used when the path includes *.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rotate_wait": schema.Int64Attribute{
											Description:         "in_tail actually does a bit more than tail -F itself. When rotating a file, some data may still need to be written to the old file as opposed to the new one.in_tail takes care of this by keeping a reference to the old file (even after it has been rotated) for some time before transitioning completely to the new file.This helps prevent data designated for the old file from getting lost. By default, this time interval is 5 seconds.The rotate_wait parameter accepts a single integer representing the number of seconds you want this time interval to be.",
											MarkdownDescription: "in_tail actually does a bit more than tail -F itself. When rotating a file, some data may still need to be written to the old file as opposed to the new one.in_tail takes care of this by keeping a reference to the old file (even after it has been rotated) for some time before transitioning completely to the new file.This helps prevent data designated for the old file from getting lost. By default, this time interval is 5 seconds.The rotate_wait parameter accepts a single integer representing the number of seconds you want this time interval to be.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"skip_refresh_on_startup": schema.BoolAttribute{
											Description:         "Skips the refresh of the watch list on startup. This reduces the startup time when * is used in path.",
											MarkdownDescription: "Skips the refresh of the watch list on startup. This reduces the startup time when * is used in path.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.StringAttribute{
											Description:         "The tag of the event.",
											MarkdownDescription: "The tag of the event.",
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

func (r *FluentdFluentIoInputV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentd_fluent_io_input_v1alpha1_manifest")

	var model FluentdFluentIoInputV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentd.fluent.io/v1alpha1")
	model.Kind = pointer.String("Input")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
