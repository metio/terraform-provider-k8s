/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package logging_banzaicloud_io_v1alpha1

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
	_ datasource.DataSource = &LoggingBanzaicloudIoClusterFlowV1Alpha1Manifest{}
)

func NewLoggingBanzaicloudIoClusterFlowV1Alpha1Manifest() datasource.DataSource {
	return &LoggingBanzaicloudIoClusterFlowV1Alpha1Manifest{}
}

type LoggingBanzaicloudIoClusterFlowV1Alpha1Manifest struct{}

type LoggingBanzaicloudIoClusterFlowV1Alpha1ManifestData struct {
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
			Concat *struct {
				Continuous_line_regexp  *string `tfsdk:"continuous_line_regexp" json:"continuous_line_regexp,omitempty"`
				Flush_interval          *int64  `tfsdk:"flush_interval" json:"flush_interval,omitempty"`
				Keep_partial_key        *bool   `tfsdk:"keep_partial_key" json:"keep_partial_key,omitempty"`
				Keep_partial_metadata   *string `tfsdk:"keep_partial_metadata" json:"keep_partial_metadata,omitempty"`
				Key                     *string `tfsdk:"key" json:"key,omitempty"`
				Multiline_end_regexp    *string `tfsdk:"multiline_end_regexp" json:"multiline_end_regexp,omitempty"`
				Multiline_start_regexp  *string `tfsdk:"multiline_start_regexp" json:"multiline_start_regexp,omitempty"`
				N_lines                 *int64  `tfsdk:"n_lines" json:"n_lines,omitempty"`
				Partial_cri_logtag_key  *string `tfsdk:"partial_cri_logtag_key" json:"partial_cri_logtag_key,omitempty"`
				Partial_cri_stream_key  *string `tfsdk:"partial_cri_stream_key" json:"partial_cri_stream_key,omitempty"`
				Partial_key             *string `tfsdk:"partial_key" json:"partial_key,omitempty"`
				Partial_metadata_format *string `tfsdk:"partial_metadata_format" json:"partial_metadata_format,omitempty"`
				Partial_value           *string `tfsdk:"partial_value" json:"partial_value,omitempty"`
				Separator               *string `tfsdk:"separator" json:"separator,omitempty"`
				Stream_identity_key     *string `tfsdk:"stream_identity_key" json:"stream_identity_key,omitempty"`
				Timeout_label           *string `tfsdk:"timeout_label" json:"timeout_label,omitempty"`
				Use_first_timestamp     *bool   `tfsdk:"use_first_timestamp" json:"use_first_timestamp,omitempty"`
				Use_partial_cri_logtag  *bool   `tfsdk:"use_partial_cri_logtag" json:"use_partial_cri_logtag,omitempty"`
				Use_partial_metadata    *string `tfsdk:"use_partial_metadata" json:"use_partial_metadata,omitempty"`
			} `tfsdk:"concat" json:"concat,omitempty"`
			Dedot *struct {
				De_dot_nested    *bool   `tfsdk:"de_dot_nested" json:"de_dot_nested,omitempty"`
				De_dot_separator *string `tfsdk:"de_dot_separator" json:"de_dot_separator,omitempty"`
			} `tfsdk:"dedot" json:"dedot,omitempty"`
			DetectExceptions *struct {
				Force_line_breaks        *bool     `tfsdk:"force_line_breaks" json:"force_line_breaks,omitempty"`
				Languages                *[]string `tfsdk:"languages" json:"languages,omitempty"`
				Match_tag                *string   `tfsdk:"match_tag" json:"match_tag,omitempty"`
				Max_bytes                *int64    `tfsdk:"max_bytes" json:"max_bytes,omitempty"`
				Max_lines                *int64    `tfsdk:"max_lines" json:"max_lines,omitempty"`
				Message                  *string   `tfsdk:"message" json:"message,omitempty"`
				Multiline_flush_interval *string   `tfsdk:"multiline_flush_interval" json:"multiline_flush_interval,omitempty"`
				Remove_tag_prefix        *string   `tfsdk:"remove_tag_prefix" json:"remove_tag_prefix,omitempty"`
				Stream                   *string   `tfsdk:"stream" json:"stream,omitempty"`
			} `tfsdk:"detect_exceptions" json:"detectExceptions,omitempty"`
			Elasticsearch_genid *struct {
				Hash_id_key          *string `tfsdk:"hash_id_key" json:"hash_id_key,omitempty"`
				Hash_type            *string `tfsdk:"hash_type" json:"hash_type,omitempty"`
				Include_tag_in_seed  *bool   `tfsdk:"include_tag_in_seed" json:"include_tag_in_seed,omitempty"`
				Include_time_in_seed *bool   `tfsdk:"include_time_in_seed" json:"include_time_in_seed,omitempty"`
				Record_keys          *string `tfsdk:"record_keys" json:"record_keys,omitempty"`
				Separator            *string `tfsdk:"separator" json:"separator,omitempty"`
				Use_entire_record    *bool   `tfsdk:"use_entire_record" json:"use_entire_record,omitempty"`
				Use_record_as_seed   *bool   `tfsdk:"use_record_as_seed" json:"use_record_as_seed,omitempty"`
			} `tfsdk:"elasticsearch_genid" json:"elasticsearch_genid,omitempty"`
			Geoip *struct {
				Backend_library         *string              `tfsdk:"backend_library" json:"backend_library,omitempty"`
				Geoip2_database         *string              `tfsdk:"geoip2_database" json:"geoip2_database,omitempty"`
				Geoip_database          *string              `tfsdk:"geoip_database" json:"geoip_database,omitempty"`
				Geoip_lookup_keys       *string              `tfsdk:"geoip_lookup_keys" json:"geoip_lookup_keys,omitempty"`
				Records                 *[]map[string]string `tfsdk:"records" json:"records,omitempty"`
				Skip_adding_null_record *bool                `tfsdk:"skip_adding_null_record" json:"skip_adding_null_record,omitempty"`
			} `tfsdk:"geoip" json:"geoip,omitempty"`
			Grep *struct {
				And *[]struct {
					Exclude *[]struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"exclude" json:"exclude,omitempty"`
					Regexp *[]struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"regexp" json:"regexp,omitempty"`
				} `tfsdk:"and" json:"and,omitempty"`
				Exclude *[]struct {
					Key     *string `tfsdk:"key" json:"key,omitempty"`
					Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
				} `tfsdk:"exclude" json:"exclude,omitempty"`
				Or *[]struct {
					Exclude *[]struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"exclude" json:"exclude,omitempty"`
					Regexp *[]struct {
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"regexp" json:"regexp,omitempty"`
				} `tfsdk:"or" json:"or,omitempty"`
				Regexp *[]struct {
					Key     *string `tfsdk:"key" json:"key,omitempty"`
					Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
				} `tfsdk:"regexp" json:"regexp,omitempty"`
			} `tfsdk:"grep" json:"grep,omitempty"`
			Kube_events_timestamp *struct {
				Mapped_time_key  *string   `tfsdk:"mapped_time_key" json:"mapped_time_key,omitempty"`
				Timestamp_fields *[]string `tfsdk:"timestamp_fields" json:"timestamp_fields,omitempty"`
			} `tfsdk:"kube_events_timestamp" json:"kube_events_timestamp,omitempty"`
			Parser *struct {
				Emit_invalid_record_to_error *bool   `tfsdk:"emit_invalid_record_to_error" json:"emit_invalid_record_to_error,omitempty"`
				Hash_value_field             *string `tfsdk:"hash_value_field" json:"hash_value_field,omitempty"`
				Inject_key_prefix            *string `tfsdk:"inject_key_prefix" json:"inject_key_prefix,omitempty"`
				Key_name                     *string `tfsdk:"key_name" json:"key_name,omitempty"`
				Parse                        *struct {
					Custom_pattern_path *struct {
						MountFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"mount_from" json:"mountFrom,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"custom_pattern_path" json:"custom_pattern_path,omitempty"`
					Delimiter              *string `tfsdk:"delimiter" json:"delimiter,omitempty"`
					Delimiter_pattern      *string `tfsdk:"delimiter_pattern" json:"delimiter_pattern,omitempty"`
					Estimate_current_event *bool   `tfsdk:"estimate_current_event" json:"estimate_current_event,omitempty"`
					Expression             *string `tfsdk:"expression" json:"expression,omitempty"`
					Format                 *string `tfsdk:"format" json:"format,omitempty"`
					Format_firstline       *string `tfsdk:"format_firstline" json:"format_firstline,omitempty"`
					Format_key             *string `tfsdk:"format_key" json:"format_key,omitempty"`
					Grok_failure_key       *string `tfsdk:"grok_failure_key" json:"grok_failure_key,omitempty"`
					Grok_name_key          *string `tfsdk:"grok_name_key" json:"grok_name_key,omitempty"`
					Grok_pattern           *string `tfsdk:"grok_pattern" json:"grok_pattern,omitempty"`
					Grok_patterns          *[]struct {
						Keep_time_key *bool   `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
						Name          *string `tfsdk:"name" json:"name,omitempty"`
						Pattern       *string `tfsdk:"pattern" json:"pattern,omitempty"`
						Time_format   *string `tfsdk:"time_format" json:"time_format,omitempty"`
						Time_key      *string `tfsdk:"time_key" json:"time_key,omitempty"`
						Timezone      *string `tfsdk:"timezone" json:"timezone,omitempty"`
					} `tfsdk:"grok_patterns" json:"grok_patterns,omitempty"`
					Keep_time_key          *bool     `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
					Keys                   *string   `tfsdk:"keys" json:"keys,omitempty"`
					Label_delimiter        *string   `tfsdk:"label_delimiter" json:"label_delimiter,omitempty"`
					Local_time             *bool     `tfsdk:"local_time" json:"local_time,omitempty"`
					Multiline              *[]string `tfsdk:"multiline" json:"multiline,omitempty"`
					Multiline_start_regexp *string   `tfsdk:"multiline_start_regexp" json:"multiline_start_regexp,omitempty"`
					Null_empty_string      *bool     `tfsdk:"null_empty_string" json:"null_empty_string,omitempty"`
					Null_value_pattern     *string   `tfsdk:"null_value_pattern" json:"null_value_pattern,omitempty"`
					Patterns               *[]struct {
						Custom_pattern_path *struct {
							MountFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"mount_from" json:"mountFrom,omitempty"`
							Value     *string `tfsdk:"value" json:"value,omitempty"`
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"custom_pattern_path" json:"custom_pattern_path,omitempty"`
						Estimate_current_event *bool   `tfsdk:"estimate_current_event" json:"estimate_current_event,omitempty"`
						Expression             *string `tfsdk:"expression" json:"expression,omitempty"`
						Format                 *string `tfsdk:"format" json:"format,omitempty"`
						Format_name            *string `tfsdk:"format_name" json:"format_name,omitempty"`
						Grok_failure_key       *string `tfsdk:"grok_failure_key" json:"grok_failure_key,omitempty"`
						Grok_name_key          *string `tfsdk:"grok_name_key" json:"grok_name_key,omitempty"`
						Grok_pattern           *string `tfsdk:"grok_pattern" json:"grok_pattern,omitempty"`
						Grok_patterns          *[]struct {
							Keep_time_key *bool   `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
							Name          *string `tfsdk:"name" json:"name,omitempty"`
							Pattern       *string `tfsdk:"pattern" json:"pattern,omitempty"`
							Time_format   *string `tfsdk:"time_format" json:"time_format,omitempty"`
							Time_key      *string `tfsdk:"time_key" json:"time_key,omitempty"`
							Timezone      *string `tfsdk:"timezone" json:"timezone,omitempty"`
						} `tfsdk:"grok_patterns" json:"grok_patterns,omitempty"`
						Keep_time_key          *bool   `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
						Local_time             *bool   `tfsdk:"local_time" json:"local_time,omitempty"`
						Multiline_start_regexp *string `tfsdk:"multiline_start_regexp" json:"multiline_start_regexp,omitempty"`
						Null_empty_string      *bool   `tfsdk:"null_empty_string" json:"null_empty_string,omitempty"`
						Null_value_pattern     *string `tfsdk:"null_value_pattern" json:"null_value_pattern,omitempty"`
						Time_format            *string `tfsdk:"time_format" json:"time_format,omitempty"`
						Time_key               *string `tfsdk:"time_key" json:"time_key,omitempty"`
						Time_type              *string `tfsdk:"time_type" json:"time_type,omitempty"`
						Timezone               *string `tfsdk:"timezone" json:"timezone,omitempty"`
						Type                   *string `tfsdk:"type" json:"type,omitempty"`
						Types                  *string `tfsdk:"types" json:"types,omitempty"`
						Utc                    *bool   `tfsdk:"utc" json:"utc,omitempty"`
					} `tfsdk:"patterns" json:"patterns,omitempty"`
					Time_format *string `tfsdk:"time_format" json:"time_format,omitempty"`
					Time_key    *string `tfsdk:"time_key" json:"time_key,omitempty"`
					Time_type   *string `tfsdk:"time_type" json:"time_type,omitempty"`
					Timezone    *string `tfsdk:"timezone" json:"timezone,omitempty"`
					Type        *string `tfsdk:"type" json:"type,omitempty"`
					Types       *string `tfsdk:"types" json:"types,omitempty"`
					Utc         *bool   `tfsdk:"utc" json:"utc,omitempty"`
				} `tfsdk:"parse" json:"parse,omitempty"`
				Parsers *[]struct {
					Custom_pattern_path *struct {
						MountFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"mount_from" json:"mountFrom,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"custom_pattern_path" json:"custom_pattern_path,omitempty"`
					Delimiter              *string `tfsdk:"delimiter" json:"delimiter,omitempty"`
					Delimiter_pattern      *string `tfsdk:"delimiter_pattern" json:"delimiter_pattern,omitempty"`
					Estimate_current_event *bool   `tfsdk:"estimate_current_event" json:"estimate_current_event,omitempty"`
					Expression             *string `tfsdk:"expression" json:"expression,omitempty"`
					Format                 *string `tfsdk:"format" json:"format,omitempty"`
					Format_firstline       *string `tfsdk:"format_firstline" json:"format_firstline,omitempty"`
					Format_key             *string `tfsdk:"format_key" json:"format_key,omitempty"`
					Grok_failure_key       *string `tfsdk:"grok_failure_key" json:"grok_failure_key,omitempty"`
					Grok_name_key          *string `tfsdk:"grok_name_key" json:"grok_name_key,omitempty"`
					Grok_pattern           *string `tfsdk:"grok_pattern" json:"grok_pattern,omitempty"`
					Grok_patterns          *[]struct {
						Keep_time_key *bool   `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
						Name          *string `tfsdk:"name" json:"name,omitempty"`
						Pattern       *string `tfsdk:"pattern" json:"pattern,omitempty"`
						Time_format   *string `tfsdk:"time_format" json:"time_format,omitempty"`
						Time_key      *string `tfsdk:"time_key" json:"time_key,omitempty"`
						Timezone      *string `tfsdk:"timezone" json:"timezone,omitempty"`
					} `tfsdk:"grok_patterns" json:"grok_patterns,omitempty"`
					Keep_time_key          *bool     `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
					Keys                   *string   `tfsdk:"keys" json:"keys,omitempty"`
					Label_delimiter        *string   `tfsdk:"label_delimiter" json:"label_delimiter,omitempty"`
					Local_time             *bool     `tfsdk:"local_time" json:"local_time,omitempty"`
					Multiline              *[]string `tfsdk:"multiline" json:"multiline,omitempty"`
					Multiline_start_regexp *string   `tfsdk:"multiline_start_regexp" json:"multiline_start_regexp,omitempty"`
					Null_empty_string      *bool     `tfsdk:"null_empty_string" json:"null_empty_string,omitempty"`
					Null_value_pattern     *string   `tfsdk:"null_value_pattern" json:"null_value_pattern,omitempty"`
					Patterns               *[]struct {
						Custom_pattern_path *struct {
							MountFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"mount_from" json:"mountFrom,omitempty"`
							Value     *string `tfsdk:"value" json:"value,omitempty"`
							ValueFrom *struct {
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"custom_pattern_path" json:"custom_pattern_path,omitempty"`
						Estimate_current_event *bool   `tfsdk:"estimate_current_event" json:"estimate_current_event,omitempty"`
						Expression             *string `tfsdk:"expression" json:"expression,omitempty"`
						Format                 *string `tfsdk:"format" json:"format,omitempty"`
						Format_name            *string `tfsdk:"format_name" json:"format_name,omitempty"`
						Grok_failure_key       *string `tfsdk:"grok_failure_key" json:"grok_failure_key,omitempty"`
						Grok_name_key          *string `tfsdk:"grok_name_key" json:"grok_name_key,omitempty"`
						Grok_pattern           *string `tfsdk:"grok_pattern" json:"grok_pattern,omitempty"`
						Grok_patterns          *[]struct {
							Keep_time_key *bool   `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
							Name          *string `tfsdk:"name" json:"name,omitempty"`
							Pattern       *string `tfsdk:"pattern" json:"pattern,omitempty"`
							Time_format   *string `tfsdk:"time_format" json:"time_format,omitempty"`
							Time_key      *string `tfsdk:"time_key" json:"time_key,omitempty"`
							Timezone      *string `tfsdk:"timezone" json:"timezone,omitempty"`
						} `tfsdk:"grok_patterns" json:"grok_patterns,omitempty"`
						Keep_time_key          *bool   `tfsdk:"keep_time_key" json:"keep_time_key,omitempty"`
						Local_time             *bool   `tfsdk:"local_time" json:"local_time,omitempty"`
						Multiline_start_regexp *string `tfsdk:"multiline_start_regexp" json:"multiline_start_regexp,omitempty"`
						Null_empty_string      *bool   `tfsdk:"null_empty_string" json:"null_empty_string,omitempty"`
						Null_value_pattern     *string `tfsdk:"null_value_pattern" json:"null_value_pattern,omitempty"`
						Time_format            *string `tfsdk:"time_format" json:"time_format,omitempty"`
						Time_key               *string `tfsdk:"time_key" json:"time_key,omitempty"`
						Time_type              *string `tfsdk:"time_type" json:"time_type,omitempty"`
						Timezone               *string `tfsdk:"timezone" json:"timezone,omitempty"`
						Type                   *string `tfsdk:"type" json:"type,omitempty"`
						Types                  *string `tfsdk:"types" json:"types,omitempty"`
						Utc                    *bool   `tfsdk:"utc" json:"utc,omitempty"`
					} `tfsdk:"patterns" json:"patterns,omitempty"`
					Time_format *string `tfsdk:"time_format" json:"time_format,omitempty"`
					Time_key    *string `tfsdk:"time_key" json:"time_key,omitempty"`
					Time_type   *string `tfsdk:"time_type" json:"time_type,omitempty"`
					Timezone    *string `tfsdk:"timezone" json:"timezone,omitempty"`
					Type        *string `tfsdk:"type" json:"type,omitempty"`
					Types       *string `tfsdk:"types" json:"types,omitempty"`
					Utc         *bool   `tfsdk:"utc" json:"utc,omitempty"`
				} `tfsdk:"parsers" json:"parsers,omitempty"`
				Remove_key_name_field    *bool `tfsdk:"remove_key_name_field" json:"remove_key_name_field,omitempty"`
				Replace_invalid_sequence *bool `tfsdk:"replace_invalid_sequence" json:"replace_invalid_sequence,omitempty"`
				Reserve_data             *bool `tfsdk:"reserve_data" json:"reserve_data,omitempty"`
				Reserve_time             *bool `tfsdk:"reserve_time" json:"reserve_time,omitempty"`
			} `tfsdk:"parser" json:"parser,omitempty"`
			Prometheus *struct {
				Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Metrics *[]struct {
					Buckets *string            `tfsdk:"buckets" json:"buckets,omitempty"`
					Desc    *string            `tfsdk:"desc" json:"desc,omitempty"`
					Key     *string            `tfsdk:"key" json:"key,omitempty"`
					Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name    *string            `tfsdk:"name" json:"name,omitempty"`
					Type    *string            `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
			Record_modifier *struct {
				Char_encoding *string              `tfsdk:"char_encoding" json:"char_encoding,omitempty"`
				Prepare_value *string              `tfsdk:"prepare_value" json:"prepare_value,omitempty"`
				Records       *[]map[string]string `tfsdk:"records" json:"records,omitempty"`
				Remove_keys   *string              `tfsdk:"remove_keys" json:"remove_keys,omitempty"`
				Replaces      *[]struct {
					Expression *string `tfsdk:"expression" json:"expression,omitempty"`
					Key        *string `tfsdk:"key" json:"key,omitempty"`
					Replace    *string `tfsdk:"replace" json:"replace,omitempty"`
				} `tfsdk:"replaces" json:"replaces,omitempty"`
				Whitelist_keys *string `tfsdk:"whitelist_keys" json:"whitelist_keys,omitempty"`
			} `tfsdk:"record_modifier" json:"record_modifier,omitempty"`
			Record_transformer *struct {
				Auto_typecast  *bool                `tfsdk:"auto_typecast" json:"auto_typecast,omitempty"`
				Enable_ruby    *bool                `tfsdk:"enable_ruby" json:"enable_ruby,omitempty"`
				Keep_keys      *string              `tfsdk:"keep_keys" json:"keep_keys,omitempty"`
				Records        *[]map[string]string `tfsdk:"records" json:"records,omitempty"`
				Remove_keys    *string              `tfsdk:"remove_keys" json:"remove_keys,omitempty"`
				Renew_record   *bool                `tfsdk:"renew_record" json:"renew_record,omitempty"`
				Renew_time_key *string              `tfsdk:"renew_time_key" json:"renew_time_key,omitempty"`
			} `tfsdk:"record_transformer" json:"record_transformer,omitempty"`
			Stdout *struct {
				Output_type *string `tfsdk:"output_type" json:"output_type,omitempty"`
			} `tfsdk:"stdout" json:"stdout,omitempty"`
			Tag_normaliser *struct {
				Format    *string `tfsdk:"format" json:"format,omitempty"`
				Match_tag *string `tfsdk:"match_tag" json:"match_tag,omitempty"`
			} `tfsdk:"tag_normaliser" json:"tag_normaliser,omitempty"`
			Throttle *struct {
				Group_bucket_limit    *int64  `tfsdk:"group_bucket_limit" json:"group_bucket_limit,omitempty"`
				Group_bucket_period_s *int64  `tfsdk:"group_bucket_period_s" json:"group_bucket_period_s,omitempty"`
				Group_drop_logs       *bool   `tfsdk:"group_drop_logs" json:"group_drop_logs,omitempty"`
				Group_key             *string `tfsdk:"group_key" json:"group_key,omitempty"`
				Group_reset_rate_s    *int64  `tfsdk:"group_reset_rate_s" json:"group_reset_rate_s,omitempty"`
				Group_warning_delay_s *int64  `tfsdk:"group_warning_delay_s" json:"group_warning_delay_s,omitempty"`
			} `tfsdk:"throttle" json:"throttle,omitempty"`
			Useragent *struct {
				Delete_key *bool   `tfsdk:"delete_key" json:"delete_key,omitempty"`
				Flatten    *bool   `tfsdk:"flatten" json:"flatten,omitempty"`
				Key_name   *string `tfsdk:"key_name" json:"key_name,omitempty"`
				Out_key    *string `tfsdk:"out_key" json:"out_key,omitempty"`
			} `tfsdk:"useragent" json:"useragent,omitempty"`
		} `tfsdk:"filters" json:"filters,omitempty"`
		FlowLabel            *string   `tfsdk:"flow_label" json:"flowLabel,omitempty"`
		GlobalOutputRefs     *[]string `tfsdk:"global_output_refs" json:"globalOutputRefs,omitempty"`
		IncludeLabelInRouter *bool     `tfsdk:"include_label_in_router" json:"includeLabelInRouter,omitempty"`
		LoggingRef           *string   `tfsdk:"logging_ref" json:"loggingRef,omitempty"`
		Match                *[]struct {
			Exclude *struct {
				Container_names  *[]string          `tfsdk:"container_names" json:"container_names,omitempty"`
				Hosts            *[]string          `tfsdk:"hosts" json:"hosts,omitempty"`
				Labels           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Namespace_labels *map[string]string `tfsdk:"namespace_labels" json:"namespace_labels,omitempty"`
				Namespaces       *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Namespaces_regex *[]string          `tfsdk:"namespaces_regex" json:"namespaces_regex,omitempty"`
			} `tfsdk:"exclude" json:"exclude,omitempty"`
			Select *struct {
				Container_names  *[]string          `tfsdk:"container_names" json:"container_names,omitempty"`
				Hosts            *[]string          `tfsdk:"hosts" json:"hosts,omitempty"`
				Labels           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Namespace_labels *map[string]string `tfsdk:"namespace_labels" json:"namespace_labels,omitempty"`
				Namespaces       *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Namespaces_regex *[]string          `tfsdk:"namespaces_regex" json:"namespaces_regex,omitempty"`
			} `tfsdk:"select" json:"select,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
		OutputRefs *[]string          `tfsdk:"output_refs" json:"outputRefs,omitempty"`
		Selectors  *map[string]string `tfsdk:"selectors" json:"selectors,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LoggingBanzaicloudIoClusterFlowV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_logging_banzaicloud_io_cluster_flow_v1alpha1_manifest"
}

func (r *LoggingBanzaicloudIoClusterFlowV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"filters": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"concat": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"continuous_line_regexp": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"flush_interval": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keep_partial_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keep_partial_metadata": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"multiline_end_regexp": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"multiline_start_regexp": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"n_lines": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partial_cri_logtag_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partial_cri_stream_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partial_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partial_metadata_format": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partial_value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"separator": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stream_identity_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout_label": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_first_timestamp": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_partial_cri_logtag": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_partial_metadata": schema.StringAttribute{
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

								"dedot": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"de_dot_nested": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"de_dot_separator": schema.StringAttribute{
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

								"detect_exceptions": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"force_line_breaks": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"languages": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"match_tag": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_bytes": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_lines": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"message": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"multiline_flush_interval": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_tag_prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stream": schema.StringAttribute{
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

								"elasticsearch_genid": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"hash_id_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hash_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_tag_in_seed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_time_in_seed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"record_keys": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"separator": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_entire_record": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_record_as_seed": schema.BoolAttribute{
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

								"geoip": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"backend_library": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"geoip2_database": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"geoip_database": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"geoip_lookup_keys": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"records": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.MapType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"skip_adding_null_record": schema.BoolAttribute{
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

								"grep": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"and": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exclude": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"regexp": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

										"or": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exclude": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"regexp": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

								"kube_events_timestamp": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"mapped_time_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timestamp_fields": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"parser": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"emit_invalid_record_to_error": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hash_value_field": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"inject_key_prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parse": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"custom_pattern_path": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"mount_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"secret_key_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
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

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"secret_key_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"delimiter": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"delimiter_pattern": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"estimate_current_event": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expression": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"format": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"format_firstline": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"format_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_failure_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_name_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_pattern": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grok_patterns": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"keep_time_key": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"pattern": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"time_format": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"timezone": schema.StringAttribute{
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

												"keep_time_key": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"keys": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_delimiter": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"local_time": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"multiline": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"multiline_start_regexp": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"null_empty_string": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"null_value_pattern": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"patterns": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"custom_pattern_path": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"mount_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"estimate_current_event": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"expression": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"format": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"format_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"grok_failure_key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"grok_name_key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"grok_pattern": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"grok_patterns": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"keep_time_key": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pattern": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"time_format": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"time_key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"timezone": schema.StringAttribute{
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

															"keep_time_key": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"local_time": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"multiline_start_regexp": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"null_empty_string": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"null_value_pattern": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_format": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"time_type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"timezone": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"types": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"utc": schema.BoolAttribute{
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

												"time_format": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"time_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timezone": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"types": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"utc": schema.BoolAttribute{
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

										"parsers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"custom_pattern_path": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"mount_from": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"secret_key_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
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

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value_from": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"secret_key_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"delimiter": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"delimiter_pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"estimate_current_event": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"expression": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"format": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"format_firstline": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"format_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"grok_failure_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"grok_name_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"grok_pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"grok_patterns": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"keep_time_key": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"time_format": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timezone": schema.StringAttribute{
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

													"keep_time_key": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"keys": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"label_delimiter": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"local_time": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"multiline": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"multiline_start_regexp": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"null_empty_string": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"null_value_pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"patterns": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"custom_pattern_path": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"mount_from": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"secret_key_ref": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"optional": schema.BoolAttribute{
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

																		"value": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"value_from": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"secret_key_ref": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"optional": schema.BoolAttribute{
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"estimate_current_event": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"expression": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"format": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"format_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grok_failure_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grok_name_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grok_pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grok_patterns": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"keep_time_key": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"pattern": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"time_format": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"time_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"timezone": schema.StringAttribute{
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

																"keep_time_key": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"local_time": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"multiline_start_regexp": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"null_empty_string": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"null_value_pattern": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time_format": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time_type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timezone": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"types": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"utc": schema.BoolAttribute{
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

													"time_format": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"time_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"time_type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timezone": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"types": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"utc": schema.BoolAttribute{
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

										"remove_key_name_field": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replace_invalid_sequence": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reserve_data": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reserve_time": schema.BoolAttribute{
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

								"prometheus": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metrics": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"buckets": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"desc": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

								"record_modifier": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"char_encoding": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"prepare_value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"records": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.MapType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_keys": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replaces": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"expression": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"replace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

										"whitelist_keys": schema.StringAttribute{
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

								"record_transformer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"auto_typecast": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_ruby": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keep_keys": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"records": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.MapType{ElemType: types.StringType},
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"remove_keys": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"renew_record": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"renew_time_key": schema.StringAttribute{
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

								"stdout": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"output_type": schema.StringAttribute{
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

								"tag_normaliser": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"format": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"match_tag": schema.StringAttribute{
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

								"throttle": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"group_bucket_limit": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_bucket_period_s": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_drop_logs": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_reset_rate_s": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"group_warning_delay_s": schema.Int64Attribute{
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

								"useragent": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"delete_key": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"flatten": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"out_key": schema.StringAttribute{
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

					"flow_label": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"global_output_refs": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"include_label_in_router": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"logging_ref": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"match": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"exclude": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"container_names": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hosts": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace_labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces_regex": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"select": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"container_names": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hosts": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace_labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespaces_regex": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

					"output_refs": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selectors": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
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
	}
}

func (r *LoggingBanzaicloudIoClusterFlowV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_logging_banzaicloud_io_cluster_flow_v1alpha1_manifest")

	var model LoggingBanzaicloudIoClusterFlowV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("logging.banzaicloud.io/v1alpha1")
	model.Kind = pointer.String("ClusterFlow")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
