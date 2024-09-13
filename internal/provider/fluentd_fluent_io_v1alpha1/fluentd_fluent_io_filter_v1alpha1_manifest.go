/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluentd_fluent_io_v1alpha1

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
	_ datasource.DataSource = &FluentdFluentIoFilterV1Alpha1Manifest{}
)

func NewFluentdFluentIoFilterV1Alpha1Manifest() datasource.DataSource {
	return &FluentdFluentIoFilterV1Alpha1Manifest{}
}

type FluentdFluentIoFilterV1Alpha1Manifest struct{}

type FluentdFluentIoFilterV1Alpha1ManifestData struct {
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
		Filters *[]struct {
			CustomPlugin *struct {
				Config *string `tfsdk:"config" json:"config,omitempty"`
			} `tfsdk:"custom_plugin" json:"customPlugin,omitempty"`
			Grep *struct {
				And *[]struct {
					Exclude *struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"exclude" json:"exclude,omitempty"`
					Regexp *struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"regexp" json:"regexp,omitempty"`
				} `tfsdk:"and" json:"and,omitempty"`
				Exclude *[]struct {
					Key     *string `tfsdk:"key" json:"key,omitempty"`
					Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
				} `tfsdk:"exclude" json:"exclude,omitempty"`
				Or *[]struct {
					Exclude *struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"exclude" json:"exclude,omitempty"`
					Regexp *struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"regexp" json:"regexp,omitempty"`
				} `tfsdk:"or" json:"or,omitempty"`
				Regexp *[]struct {
					Key     *string `tfsdk:"key" json:"key,omitempty"`
					Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
				} `tfsdk:"regexp" json:"regexp,omitempty"`
			} `tfsdk:"grep" json:"grep,omitempty"`
			LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Parser   *struct {
				EmitInvalidRecordToError *bool   `tfsdk:"emit_invalid_record_to_error" json:"emitInvalidRecordToError,omitempty"`
				HashValueField           *string `tfsdk:"hash_value_field" json:"hashValueField,omitempty"`
				InjectKeyPrefix          *string `tfsdk:"inject_key_prefix" json:"injectKeyPrefix,omitempty"`
				KeyName                  *string `tfsdk:"key_name" json:"keyName,omitempty"`
				Parse                    *struct {
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
				RemoveKeyNameField     *bool `tfsdk:"remove_key_name_field" json:"removeKeyNameField,omitempty"`
				ReplaceInvalidSequence *bool `tfsdk:"replace_invalid_sequence" json:"replaceInvalidSequence,omitempty"`
				ReserveData            *bool `tfsdk:"reserve_data" json:"reserveData,omitempty"`
				ReserveTime            *bool `tfsdk:"reserve_time" json:"reserveTime,omitempty"`
			} `tfsdk:"parser" json:"parser,omitempty"`
			RecordTransformer *struct {
				AutoTypecast *bool   `tfsdk:"auto_typecast" json:"autoTypecast,omitempty"`
				EnableRuby   *bool   `tfsdk:"enable_ruby" json:"enableRuby,omitempty"`
				KeepKeys     *string `tfsdk:"keep_keys" json:"keepKeys,omitempty"`
				Records      *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"records" json:"records,omitempty"`
				RemoveKeys   *string `tfsdk:"remove_keys" json:"removeKeys,omitempty"`
				RenewRecord  *bool   `tfsdk:"renew_record" json:"renewRecord,omitempty"`
				RenewTimeKey *string `tfsdk:"renew_time_key" json:"renewTimeKey,omitempty"`
			} `tfsdk:"record_transformer" json:"recordTransformer,omitempty"`
			Stdout *struct {
				Format *struct {
					Delimiter           *string `tfsdk:"delimiter" json:"delimiter,omitempty"`
					Id                  *string `tfsdk:"id" json:"id,omitempty"`
					Localtime           *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
					LogLevel            *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					Newline             *string `tfsdk:"newline" json:"newline,omitempty"`
					OutputTag           *bool   `tfsdk:"output_tag" json:"outputTag,omitempty"`
					OutputTime          *bool   `tfsdk:"output_time" json:"outputTime,omitempty"`
					TimeFormat          *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
					TimeFormatFallbacks *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
					TimeType            *string `tfsdk:"time_type" json:"timeType,omitempty"`
					Timezone            *string `tfsdk:"timezone" json:"timezone,omitempty"`
					Type                *string `tfsdk:"type" json:"type,omitempty"`
					Utc                 *bool   `tfsdk:"utc" json:"utc,omitempty"`
				} `tfsdk:"format" json:"format,omitempty"`
				Inject *struct {
					Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
					HostnameKey *string `tfsdk:"hostname_key" json:"hostnameKey,omitempty"`
					Inline      *struct {
						Localtime           *bool   `tfsdk:"localtime" json:"localtime,omitempty"`
						TimeFormat          *string `tfsdk:"time_format" json:"timeFormat,omitempty"`
						TimeFormatFallbacks *string `tfsdk:"time_format_fallbacks" json:"timeFormatFallbacks,omitempty"`
						TimeType            *string `tfsdk:"time_type" json:"timeType,omitempty"`
						Timezone            *string `tfsdk:"timezone" json:"timezone,omitempty"`
						Utc                 *bool   `tfsdk:"utc" json:"utc,omitempty"`
					} `tfsdk:"inline" json:"inline,omitempty"`
					TagKey      *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
					TimeKey     *string `tfsdk:"time_key" json:"timeKey,omitempty"`
					WorkerIdKey *string `tfsdk:"worker_id_key" json:"workerIdKey,omitempty"`
				} `tfsdk:"inject" json:"inject,omitempty"`
			} `tfsdk:"stdout" json:"stdout,omitempty"`
			Tag *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"filters" json:"filters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluentdFluentIoFilterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluentd_fluent_io_filter_v1alpha1_manifest"
}

func (r *FluentdFluentIoFilterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Filter is the Schema for the filters API",
		MarkdownDescription: "Filter is the Schema for the filters API",
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
				Description:         "FilterSpec defines the desired state of Filter",
				MarkdownDescription: "FilterSpec defines the desired state of Filter",
				Attributes: map[string]schema.Attribute{
					"filters": schema.ListNestedAttribute{
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

								"grep": schema.SingleNestedAttribute{
									Description:         "The filter_grep filter plugin",
									MarkdownDescription: "The filter_grep filter plugin",
									Attributes: map[string]schema.Attribute{
										"and": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exclude": schema.SingleNestedAttribute{
														Description:         "Exclude defines the parameters for the exclude plugin",
														MarkdownDescription: "Exclude defines the parameters for the exclude plugin",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
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

													"regexp": schema.SingleNestedAttribute{
														Description:         "Regexp defines the parameters for the regexp plugin",
														MarkdownDescription: "Regexp defines the parameters for the regexp plugin",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"exclude": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

										"or": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exclude": schema.SingleNestedAttribute{
														Description:         "Exclude defines the parameters for the exclude plugin",
														MarkdownDescription: "Exclude defines the parameters for the exclude plugin",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
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

													"regexp": schema.SingleNestedAttribute{
														Description:         "Regexp defines the parameters for the regexp plugin",
														MarkdownDescription: "Regexp defines the parameters for the regexp plugin",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"regexp": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

								"log_level": schema.StringAttribute{
									Description:         "The @log_level parameter specifies the plugin-specific logging level",
									MarkdownDescription: "The @log_level parameter specifies the plugin-specific logging level",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"parser": schema.SingleNestedAttribute{
									Description:         "The filter_parser filter plugin",
									MarkdownDescription: "The filter_parser filter plugin",
									Attributes: map[string]schema.Attribute{
										"emit_invalid_record_to_error": schema.BoolAttribute{
											Description:         "Emits invalid record to @ERROR label. Invalid cases are: key does not exist;the format is not matched;an unexpected error. If you want to ignore these errors, set false.",
											MarkdownDescription: "Emits invalid record to @ERROR label. Invalid cases are: key does not exist;the format is not matched;an unexpected error. If you want to ignore these errors, set false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hash_value_field": schema.StringAttribute{
											Description:         "Stores the parsed values as a hash value in a field.",
											MarkdownDescription: "Stores the parsed values as a hash value in a field.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"inject_key_prefix": schema.StringAttribute{
											Description:         "Stores the parsed values with the specified key name prefix.",
											MarkdownDescription: "Stores the parsed values with the specified key name prefix.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_name": schema.StringAttribute{
											Description:         "Specifies the field name in the record to parse. Required parameter. i.e: If set keyName to log, {'key':'value','log':'{'time':1622473200,'user':1}'} => {'user':1}",
											MarkdownDescription: "Specifies the field name in the record to parse. Required parameter. i.e: If set keyName to log, {'key':'value','log':'{'time':1622473200,'user':1}'} => {'user':1}",
											Required:            true,
											Optional:            false,
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

										"remove_key_name_field": schema.BoolAttribute{
											Description:         "Removes key_name field when parsing is succeeded.",
											MarkdownDescription: "Removes key_name field when parsing is succeeded.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replace_invalid_sequence": schema.BoolAttribute{
											Description:         "If true, invalid string is replaced with safe characters and re-parse it.",
											MarkdownDescription: "If true, invalid string is replaced with safe characters and re-parse it.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reserve_data": schema.BoolAttribute{
											Description:         "Keeps the original key-value pair in the parsed result. Default is false. i.e: If set keyName to log, reverseData to true, {'key':'value','log':'{'user':1,'num':2}'} => {'key':'value','log':'{'user':1,'num':2}','user':1,'num':2}",
											MarkdownDescription: "Keeps the original key-value pair in the parsed result. Default is false. i.e: If set keyName to log, reverseData to true, {'key':'value','log':'{'user':1,'num':2}'} => {'key':'value','log':'{'user':1,'num':2}','user':1,'num':2}",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reserve_time": schema.BoolAttribute{
											Description:         "Keeps the original event time in the parsed result. Default is false.",
											MarkdownDescription: "Keeps the original event time in the parsed result. Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"record_transformer": schema.SingleNestedAttribute{
									Description:         "The filter_record_transformer filter plugin",
									MarkdownDescription: "The filter_record_transformer filter plugin",
									Attributes: map[string]schema.Attribute{
										"auto_typecast": schema.BoolAttribute{
											Description:         "Automatically casts the field types. Default is false. This option is effective only for field values comprised of a single placeholder.",
											MarkdownDescription: "Automatically casts the field types. Default is false. This option is effective only for field values comprised of a single placeholder.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_ruby": schema.BoolAttribute{
											Description:         "When set to true, the full Ruby syntax is enabled in the ${...} expression. The default value is false. i.e: jsonized_record ${record.to_json}",
											MarkdownDescription: "When set to true, the full Ruby syntax is enabled in the ${...} expression. The default value is false. i.e: jsonized_record ${record.to_json}",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keep_keys": schema.StringAttribute{
											Description:         "A list of keys to keep. Only relevant if renew_record is set to true.",
											MarkdownDescription: "A list of keys to keep. Only relevant if renew_record is set to true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"records": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "New field can be defined as key",
														MarkdownDescription: "New field can be defined as key",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "The value must from Record properties. See https://docs.fluentd.org/filter/record_transformer#less-than-record-greater-than-directive",
														MarkdownDescription: "The value must from Record properties. See https://docs.fluentd.org/filter/record_transformer#less-than-record-greater-than-directive",
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

										"remove_keys": schema.StringAttribute{
											Description:         "A list of keys to delete. Supports nested field via record_accessor syntax since v1.1.0.",
											MarkdownDescription: "A list of keys to delete. Supports nested field via record_accessor syntax since v1.1.0.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"renew_record": schema.BoolAttribute{
											Description:         "By default, the record transformer filter mutates the incoming data. However, if this parameter is set to true, it modifies a new empty hash instead.",
											MarkdownDescription: "By default, the record transformer filter mutates the incoming data. However, if this parameter is set to true, it modifies a new empty hash instead.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"renew_time_key": schema.StringAttribute{
											Description:         "renew_time_key foo overwrites the time of events with a value of the record field foo if exists. The value of foo must be a Unix timestamp.",
											MarkdownDescription: "renew_time_key foo overwrites the time of events with a value of the record field foo if exists. The value of foo must be a Unix timestamp.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"stdout": schema.SingleNestedAttribute{
									Description:         "The filter_stdout filter plugin",
									MarkdownDescription: "The filter_stdout filter plugin",
									Attributes: map[string]schema.Attribute{
										"format": schema.SingleNestedAttribute{
											Description:         "The format section",
											MarkdownDescription: "The format section",
											Attributes: map[string]schema.Attribute{
												"delimiter": schema.StringAttribute{
													Description:         "Delimiter for each field.",
													MarkdownDescription: "Delimiter for each field.",
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

												"newline": schema.StringAttribute{
													Description:         "Specify newline characters.",
													MarkdownDescription: "Specify newline characters.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("lf", "crlf"),
													},
												},

												"output_tag": schema.BoolAttribute{
													Description:         "Output tag field if true.",
													MarkdownDescription: "Output tag field if true.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_time": schema.BoolAttribute{
													Description:         "Output time field if true.",
													MarkdownDescription: "Output time field if true.",
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
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("out_file", "json", "ltsv", "csv", "msgpack", "hash", "single_value"),
													},
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

										"inject": schema.SingleNestedAttribute{
											Description:         "The inject section",
											MarkdownDescription: "The inject section",
											Attributes: map[string]schema.Attribute{
												"hostname": schema.StringAttribute{
													Description:         "Hostname value",
													MarkdownDescription: "Hostname value",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"hostname_key": schema.StringAttribute{
													Description:         "The field name to inject hostname",
													MarkdownDescription: "The field name to inject hostname",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"inline": schema.SingleNestedAttribute{
													Description:         "Time section",
													MarkdownDescription: "Time section",
													Attributes: map[string]schema.Attribute{
														"localtime": schema.BoolAttribute{
															Description:         "If true, uses local time.",
															MarkdownDescription: "If true, uses local time.",
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

														"timezone": schema.StringAttribute{
															Description:         "Uses the specified timezone.",
															MarkdownDescription: "Uses the specified timezone.",
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

												"tag_key": schema.StringAttribute{
													Description:         "The field name to inject tag",
													MarkdownDescription: "The field name to inject tag",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_key": schema.StringAttribute{
													Description:         "The field name to inject time",
													MarkdownDescription: "The field name to inject time",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"worker_id_key": schema.StringAttribute{
													Description:         "The field name to inject worker_id",
													MarkdownDescription: "The field name to inject worker_id",
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

								"tag": schema.StringAttribute{
									Description:         "Which tag to be matched.",
									MarkdownDescription: "Which tag to be matched.",
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
	}
}

func (r *FluentdFluentIoFilterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluentd_fluent_io_filter_v1alpha1_manifest")

	var model FluentdFluentIoFilterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluentd.fluent.io/v1alpha1")
	model.Kind = pointer.String("Filter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
