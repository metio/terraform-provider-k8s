/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_solo_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GatewaySoloIoRouteOptionV1Manifest{}
)

func NewGatewaySoloIoRouteOptionV1Manifest() datasource.DataSource {
	return &GatewaySoloIoRouteOptionV1Manifest{}
}

type GatewaySoloIoRouteOptionV1Manifest struct{}

type GatewaySoloIoRouteOptionV1ManifestData struct {
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
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
		Options *struct {
			AppendXForwardedHost *bool `tfsdk:"append_x_forwarded_host" json:"appendXForwardedHost,omitempty"`
			AutoHostRewrite      *bool `tfsdk:"auto_host_rewrite" json:"autoHostRewrite,omitempty"`
			BufferPerRoute       *struct {
				Buffer *struct {
					MaxRequestBytes *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
				} `tfsdk:"buffer" json:"buffer,omitempty"`
				Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
			} `tfsdk:"buffer_per_route" json:"bufferPerRoute,omitempty"`
			Cors *struct {
				AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
				AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
				AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
				AllowOrigin      *[]string `tfsdk:"allow_origin" json:"allowOrigin,omitempty"`
				AllowOriginRegex *[]string `tfsdk:"allow_origin_regex" json:"allowOriginRegex,omitempty"`
				DisableForRoute  *bool     `tfsdk:"disable_for_route" json:"disableForRoute,omitempty"`
				ExposeHeaders    *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
				MaxAge           *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
			} `tfsdk:"cors" json:"cors,omitempty"`
			Csrf *struct {
				AdditionalOrigins *[]struct {
					Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
					IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
					Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
					SafeRegex  *struct {
						GoogleRe2 *struct {
							MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
						} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
						Regex *string `tfsdk:"regex" json:"regex,omitempty"`
					} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
					Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"additional_origins" json:"additionalOrigins,omitempty"`
				FilterEnabled *struct {
					DefaultValue *struct {
						Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
						Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
					} `tfsdk:"default_value" json:"defaultValue,omitempty"`
					RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
				} `tfsdk:"filter_enabled" json:"filterEnabled,omitempty"`
				ShadowEnabled *struct {
					DefaultValue *struct {
						Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
						Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
					} `tfsdk:"default_value" json:"defaultValue,omitempty"`
					RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
				} `tfsdk:"shadow_enabled" json:"shadowEnabled,omitempty"`
			} `tfsdk:"csrf" json:"csrf,omitempty"`
			Dlp *struct {
				Actions *[]struct {
					ActionType   *string `tfsdk:"action_type" json:"actionType,omitempty"`
					CustomAction *struct {
						MaskChar *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Percent  *struct {
							Value *float64 `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"percent" json:"percent,omitempty"`
						Regex        *[]string `tfsdk:"regex" json:"regex,omitempty"`
						RegexActions *[]struct {
							Regex    *string `tfsdk:"regex" json:"regex,omitempty"`
							Subgroup *int64  `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"regex_actions" json:"regexActions,omitempty"`
					} `tfsdk:"custom_action" json:"customAction,omitempty"`
					KeyValueAction *struct {
						KeyToMask *string `tfsdk:"key_to_mask" json:"keyToMask,omitempty"`
						MaskChar  *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Percent   *struct {
							Value *float64 `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"percent" json:"percent,omitempty"`
					} `tfsdk:"key_value_action" json:"keyValueAction,omitempty"`
					Shadow *bool `tfsdk:"shadow" json:"shadow,omitempty"`
				} `tfsdk:"actions" json:"actions,omitempty"`
				EnabledFor *string `tfsdk:"enabled_for" json:"enabledFor,omitempty"`
			} `tfsdk:"dlp" json:"dlp,omitempty"`
			EnvoyMetadata *map[string]string `tfsdk:"envoy_metadata" json:"envoyMetadata,omitempty"`
			ExtProc       *struct {
				Disabled  *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				Overrides *struct {
					AsyncMode   *bool `tfsdk:"async_mode" json:"asyncMode,omitempty"`
					GrpcService *struct {
						Authority        *string `tfsdk:"authority" json:"authority,omitempty"`
						ExtProcServerRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"ext_proc_server_ref" json:"extProcServerRef,omitempty"`
						InitialMetadata *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"initial_metadata" json:"initialMetadata,omitempty"`
						RetryPolicy *struct {
							NumRetries   *int64 `tfsdk:"num_retries" json:"numRetries,omitempty"`
							RetryBackOff *struct {
								BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
								MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
							} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
						} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
						Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					} `tfsdk:"grpc_service" json:"grpcService,omitempty"`
					MetadataContextNamespaces *[]string `tfsdk:"metadata_context_namespaces" json:"metadataContextNamespaces,omitempty"`
					ProcessingMode            *struct {
						RequestBodyMode     *string `tfsdk:"request_body_mode" json:"requestBodyMode,omitempty"`
						RequestHeaderMode   *string `tfsdk:"request_header_mode" json:"requestHeaderMode,omitempty"`
						RequestTrailerMode  *string `tfsdk:"request_trailer_mode" json:"requestTrailerMode,omitempty"`
						ResponseBodyMode    *string `tfsdk:"response_body_mode" json:"responseBodyMode,omitempty"`
						ResponseHeaderMode  *string `tfsdk:"response_header_mode" json:"responseHeaderMode,omitempty"`
						ResponseTrailerMode *string `tfsdk:"response_trailer_mode" json:"responseTrailerMode,omitempty"`
					} `tfsdk:"processing_mode" json:"processingMode,omitempty"`
					RequestAttributes              *[]string `tfsdk:"request_attributes" json:"requestAttributes,omitempty"`
					ResponseAttributes             *[]string `tfsdk:"response_attributes" json:"responseAttributes,omitempty"`
					TypedMetadataContextNamespaces *[]string `tfsdk:"typed_metadata_context_namespaces" json:"typedMetadataContextNamespaces,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
			} `tfsdk:"ext_proc" json:"extProc,omitempty"`
			Extauth *struct {
				ConfigRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"config_ref" json:"configRef,omitempty"`
				CustomAuth *struct {
					ContextExtensions *map[string]string `tfsdk:"context_extensions" json:"contextExtensions,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"custom_auth" json:"customAuth,omitempty"`
				Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
			} `tfsdk:"extauth" json:"extauth,omitempty"`
			Extensions *struct {
				Configs *map[string]string `tfsdk:"configs" json:"configs,omitempty"`
			} `tfsdk:"extensions" json:"extensions,omitempty"`
			Faults *struct {
				Abort *struct {
					HttpStatus *int64   `tfsdk:"http_status" json:"httpStatus,omitempty"`
					Percentage *float64 `tfsdk:"percentage" json:"percentage,omitempty"`
				} `tfsdk:"abort" json:"abort,omitempty"`
				Delay *struct {
					FixedDelay *string  `tfsdk:"fixed_delay" json:"fixedDelay,omitempty"`
					Percentage *float64 `tfsdk:"percentage" json:"percentage,omitempty"`
				} `tfsdk:"delay" json:"delay,omitempty"`
			} `tfsdk:"faults" json:"faults,omitempty"`
			HeaderManipulation *struct {
				RequestHeadersToAdd *[]struct {
					Append *bool `tfsdk:"append" json:"append,omitempty"`
					Header *struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"header" json:"header,omitempty"`
					HeaderSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"header_secret_ref" json:"headerSecretRef,omitempty"`
				} `tfsdk:"request_headers_to_add" json:"requestHeadersToAdd,omitempty"`
				RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" json:"requestHeadersToRemove,omitempty"`
				ResponseHeadersToAdd   *[]struct {
					Append *bool `tfsdk:"append" json:"append,omitempty"`
					Header *struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"header" json:"header,omitempty"`
				} `tfsdk:"response_headers_to_add" json:"responseHeadersToAdd,omitempty"`
				ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" json:"responseHeadersToRemove,omitempty"`
			} `tfsdk:"header_manipulation" json:"headerManipulation,omitempty"`
			HostRewrite          *string `tfsdk:"host_rewrite" json:"hostRewrite,omitempty"`
			HostRewritePathRegex *struct {
				Pattern *struct {
					GoogleRe2 *struct {
						MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
					} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
					Regex *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"pattern" json:"pattern,omitempty"`
				Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
			} `tfsdk:"host_rewrite_path_regex" json:"hostRewritePathRegex,omitempty"`
			IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
			Jwt         *struct {
				Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
			} `tfsdk:"jwt" json:"jwt,omitempty"`
			JwtStaged *struct {
				AfterExtAuth *struct {
					Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
				} `tfsdk:"after_ext_auth" json:"afterExtAuth,omitempty"`
				BeforeExtAuth *struct {
					Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
				} `tfsdk:"before_ext_auth" json:"beforeExtAuth,omitempty"`
			} `tfsdk:"jwt_staged" json:"jwtStaged,omitempty"`
			LbHash *struct {
				HashPolicies *[]struct {
					Cookie *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
						Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
					} `tfsdk:"cookie" json:"cookie,omitempty"`
					Header   *string `tfsdk:"header" json:"header,omitempty"`
					SourceIp *bool   `tfsdk:"source_ip" json:"sourceIp,omitempty"`
					Terminal *bool   `tfsdk:"terminal" json:"terminal,omitempty"`
				} `tfsdk:"hash_policies" json:"hashPolicies,omitempty"`
			} `tfsdk:"lb_hash" json:"lbHash,omitempty"`
			MaxStreamDuration *struct {
				GrpcTimeoutHeaderMax    *string `tfsdk:"grpc_timeout_header_max" json:"grpcTimeoutHeaderMax,omitempty"`
				GrpcTimeoutHeaderOffset *string `tfsdk:"grpc_timeout_header_offset" json:"grpcTimeoutHeaderOffset,omitempty"`
				MaxStreamDuration       *string `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
			} `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
			PrefixRewrite    *string `tfsdk:"prefix_rewrite" json:"prefixRewrite,omitempty"`
			RateLimitConfigs *struct {
				Refs *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"refs" json:"refs,omitempty"`
			} `tfsdk:"rate_limit_configs" json:"rateLimitConfigs,omitempty"`
			RateLimitEarlyConfigs *struct {
				Refs *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"refs" json:"refs,omitempty"`
			} `tfsdk:"rate_limit_early_configs" json:"rateLimitEarlyConfigs,omitempty"`
			RateLimitRegularConfigs *struct {
				Refs *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"refs" json:"refs,omitempty"`
			} `tfsdk:"rate_limit_regular_configs" json:"rateLimitRegularConfigs,omitempty"`
			Ratelimit *struct {
				IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" json:"includeVhRateLimits,omitempty"`
				LocalRatelimit      *struct {
					FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
					MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
					TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
				} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
				RateLimits *[]struct {
					Actions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
						GenericKey         *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers         *[]struct {
								ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
								InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name         *string `tfsdk:"name" json:"name,omitempty"`
								PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
								PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
								RangeMatch   *struct {
									End   *int64 `tfsdk:"end" json:"end,omitempty"`
									Start *int64 `tfsdk:"start" json:"start,omitempty"`
								} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
								RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
								SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
						} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
						Metadata *struct {
							DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							MetadataKey   *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
							Source *string `tfsdk:"source" json:"source,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
					} `tfsdk:"actions" json:"actions,omitempty"`
					Limit *struct {
						DynamicMetadata *struct {
							MetadataKey *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
						} `tfsdk:"dynamic_metadata" json:"dynamicMetadata,omitempty"`
					} `tfsdk:"limit" json:"limit,omitempty"`
					SetActions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
						GenericKey         *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers         *[]struct {
								ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
								InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name         *string `tfsdk:"name" json:"name,omitempty"`
								PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
								PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
								RangeMatch   *struct {
									End   *int64 `tfsdk:"end" json:"end,omitempty"`
									Start *int64 `tfsdk:"start" json:"start,omitempty"`
								} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
								RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
								SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
						} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
						Metadata *struct {
							DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							MetadataKey   *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
							Source *string `tfsdk:"source" json:"source,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
					} `tfsdk:"set_actions" json:"setActions,omitempty"`
				} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
			} `tfsdk:"ratelimit" json:"ratelimit,omitempty"`
			RatelimitBasic *struct {
				AnonymousLimits *struct {
					RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
					Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"anonymous_limits" json:"anonymousLimits,omitempty"`
				AuthorizedLimits *struct {
					RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
					Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"authorized_limits" json:"authorizedLimits,omitempty"`
			} `tfsdk:"ratelimit_basic" json:"ratelimitBasic,omitempty"`
			RatelimitEarly *struct {
				IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" json:"includeVhRateLimits,omitempty"`
				LocalRatelimit      *struct {
					FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
					MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
					TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
				} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
				RateLimits *[]struct {
					Actions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
						GenericKey         *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers         *[]struct {
								ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
								InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name         *string `tfsdk:"name" json:"name,omitempty"`
								PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
								PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
								RangeMatch   *struct {
									End   *int64 `tfsdk:"end" json:"end,omitempty"`
									Start *int64 `tfsdk:"start" json:"start,omitempty"`
								} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
								RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
								SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
						} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
						Metadata *struct {
							DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							MetadataKey   *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
							Source *string `tfsdk:"source" json:"source,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
					} `tfsdk:"actions" json:"actions,omitempty"`
					Limit *struct {
						DynamicMetadata *struct {
							MetadataKey *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
						} `tfsdk:"dynamic_metadata" json:"dynamicMetadata,omitempty"`
					} `tfsdk:"limit" json:"limit,omitempty"`
					SetActions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
						GenericKey         *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers         *[]struct {
								ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
								InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name         *string `tfsdk:"name" json:"name,omitempty"`
								PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
								PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
								RangeMatch   *struct {
									End   *int64 `tfsdk:"end" json:"end,omitempty"`
									Start *int64 `tfsdk:"start" json:"start,omitempty"`
								} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
								RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
								SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
						} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
						Metadata *struct {
							DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							MetadataKey   *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
							Source *string `tfsdk:"source" json:"source,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
					} `tfsdk:"set_actions" json:"setActions,omitempty"`
				} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
			} `tfsdk:"ratelimit_early" json:"ratelimitEarly,omitempty"`
			RatelimitRegular *struct {
				IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" json:"includeVhRateLimits,omitempty"`
				LocalRatelimit      *struct {
					FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
					MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
					TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
				} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
				RateLimits *[]struct {
					Actions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
						GenericKey         *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers         *[]struct {
								ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
								InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name         *string `tfsdk:"name" json:"name,omitempty"`
								PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
								PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
								RangeMatch   *struct {
									End   *int64 `tfsdk:"end" json:"end,omitempty"`
									Start *int64 `tfsdk:"start" json:"start,omitempty"`
								} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
								RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
								SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
						} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
						Metadata *struct {
							DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							MetadataKey   *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
							Source *string `tfsdk:"source" json:"source,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
					} `tfsdk:"actions" json:"actions,omitempty"`
					Limit *struct {
						DynamicMetadata *struct {
							MetadataKey *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
						} `tfsdk:"dynamic_metadata" json:"dynamicMetadata,omitempty"`
					} `tfsdk:"limit" json:"limit,omitempty"`
					SetActions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
						GenericKey         *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" json:"genericKey,omitempty"`
						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
							Headers         *[]struct {
								ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
								InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name         *string `tfsdk:"name" json:"name,omitempty"`
								PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
								PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
								RangeMatch   *struct {
									End   *int64 `tfsdk:"end" json:"end,omitempty"`
									Start *int64 `tfsdk:"start" json:"start,omitempty"`
								} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
								RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
								SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
						} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
						Metadata *struct {
							DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							MetadataKey   *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Path *[]struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
							Source *string `tfsdk:"source" json:"source,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
							HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
						} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
						SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
					} `tfsdk:"set_actions" json:"setActions,omitempty"`
				} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
			} `tfsdk:"ratelimit_regular" json:"ratelimitRegular,omitempty"`
			Rbac *struct {
				Disable  *bool `tfsdk:"disable" json:"disable,omitempty"`
				Policies *struct {
					NestedClaimDelimiter *string `tfsdk:"nested_claim_delimiter" json:"nestedClaimDelimiter,omitempty"`
					Permissions          *struct {
						Methods    *[]string `tfsdk:"methods" json:"methods,omitempty"`
						PathPrefix *string   `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
					} `tfsdk:"permissions" json:"permissions,omitempty"`
					Principals *[]struct {
						JwtPrincipal *struct {
							Claims   *map[string]string `tfsdk:"claims" json:"claims,omitempty"`
							Matcher  *string            `tfsdk:"matcher" json:"matcher,omitempty"`
							Provider *string            `tfsdk:"provider" json:"provider,omitempty"`
						} `tfsdk:"jwt_principal" json:"jwtPrincipal,omitempty"`
					} `tfsdk:"principals" json:"principals,omitempty"`
				} `tfsdk:"policies" json:"policies,omitempty"`
			} `tfsdk:"rbac" json:"rbac,omitempty"`
			RegexRewrite *struct {
				Pattern *struct {
					GoogleRe2 *struct {
						MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
					} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
					Regex *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"pattern" json:"pattern,omitempty"`
				Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
			} `tfsdk:"regex_rewrite" json:"regexRewrite,omitempty"`
			Retries *struct {
				NumRetries         *int64  `tfsdk:"num_retries" json:"numRetries,omitempty"`
				PerTryTimeout      *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
				PreviousPriorities *struct {
					UpdateFrequency *int64 `tfsdk:"update_frequency" json:"updateFrequency,omitempty"`
				} `tfsdk:"previous_priorities" json:"previousPriorities,omitempty"`
				RetryBackOff *struct {
					BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
					MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
				} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
				RetryOn *string `tfsdk:"retry_on" json:"retryOn,omitempty"`
			} `tfsdk:"retries" json:"retries,omitempty"`
			Shadowing *struct {
				Percentage *float64 `tfsdk:"percentage" json:"percentage,omitempty"`
				Upstream   *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"upstream" json:"upstream,omitempty"`
			} `tfsdk:"shadowing" json:"shadowing,omitempty"`
			StagedTransformations *struct {
				Early *struct {
					RequestTransforms *[]struct {
						ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
						Matcher         *struct {
							CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
							ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
							Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
							Headers        *[]struct {
								InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value       *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
							Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
							QueryParameters *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"matcher" json:"matcher,omitempty"`
						RequestTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
					ResponseTransforms *[]struct {
						Matchers *[]struct {
							InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
							Value       *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"matchers" json:"matchers,omitempty"`
						ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
				} `tfsdk:"early" json:"early,omitempty"`
				EscapeCharacters       *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
				InheritTransformation  *bool `tfsdk:"inherit_transformation" json:"inheritTransformation,omitempty"`
				LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
				PostRouting            *struct {
					RequestTransforms *[]struct {
						ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
						Matcher         *struct {
							CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
							ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
							Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
							Headers        *[]struct {
								InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value       *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
							Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
							QueryParameters *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"matcher" json:"matcher,omitempty"`
						RequestTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
					ResponseTransforms *[]struct {
						Matchers *[]struct {
							InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
							Value       *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"matchers" json:"matchers,omitempty"`
						ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
				} `tfsdk:"post_routing" json:"postRouting,omitempty"`
				Regular *struct {
					RequestTransforms *[]struct {
						ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
						Matcher         *struct {
							CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
							ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
							Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
							Headers        *[]struct {
								InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value       *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
							Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
							QueryParameters *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"matcher" json:"matcher,omitempty"`
						RequestTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
					ResponseTransforms *[]struct {
						Matchers *[]struct {
							InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
							Value       *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"matchers" json:"matchers,omitempty"`
						ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header          *string            `tfsdk:"header" json:"header,omitempty"`
									Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
									Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
									ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
									Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
				} `tfsdk:"regular" json:"regular,omitempty"`
			} `tfsdk:"staged_transformations" json:"stagedTransformations,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			Tracing *struct {
				Propagate        *bool   `tfsdk:"propagate" json:"propagate,omitempty"`
				RouteDescriptor  *string `tfsdk:"route_descriptor" json:"routeDescriptor,omitempty"`
				TracePercentages *struct {
					ClientSamplePercentage  *float64 `tfsdk:"client_sample_percentage" json:"clientSamplePercentage,omitempty"`
					OverallSamplePercentage *float64 `tfsdk:"overall_sample_percentage" json:"overallSamplePercentage,omitempty"`
					RandomSamplePercentage  *float64 `tfsdk:"random_sample_percentage" json:"randomSamplePercentage,omitempty"`
				} `tfsdk:"trace_percentages" json:"tracePercentages,omitempty"`
			} `tfsdk:"tracing" json:"tracing,omitempty"`
			Transformations *struct {
				ClearRouteCache       *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
				RequestTransformation *struct {
					HeaderBodyTransform *struct {
						AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
					} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
					LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
					TransformationTemplate *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
						Body              *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						DynamicMetadataValues *[]struct {
							JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
							Value             *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
						EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						Extractors       *struct {
							Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header          *string            `tfsdk:"header" json:"header,omitempty"`
							Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
							Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
							ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
							Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"extractors" json:"extractors,omitempty"`
						Headers *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						HeadersToAppend *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
						HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
						IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
						ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
						Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
					} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
					XsltTransformation *struct {
						NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
						SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
						Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
					} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
				} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
				ResponseTransformation *struct {
					HeaderBodyTransform *struct {
						AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
					} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
					LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
					TransformationTemplate *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
						Body              *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						DynamicMetadataValues *[]struct {
							JsonToProto       *bool   `tfsdk:"json_to_proto" json:"jsonToProto,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
							Value             *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
						EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						Extractors       *struct {
							Body            *map[string]string `tfsdk:"body" json:"body,omitempty"`
							Header          *string            `tfsdk:"header" json:"header,omitempty"`
							Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
							Regex           *string            `tfsdk:"regex" json:"regex,omitempty"`
							ReplacementText *string            `tfsdk:"replacement_text" json:"replacementText,omitempty"`
							Subgroup        *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
						} `tfsdk:"extractors" json:"extractors,omitempty"`
						Headers *struct {
							Text *string `tfsdk:"text" json:"text,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						HeadersToAppend *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
						HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
						IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
						ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
						Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
					} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
					XsltTransformation *struct {
						NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
						SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
						Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
					} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
				} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
			} `tfsdk:"transformations" json:"transformations,omitempty"`
			Upgrades *[]struct {
				Connect *struct {
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"connect" json:"connect,omitempty"`
				Websocket *struct {
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"websocket" json:"websocket,omitempty"`
			} `tfsdk:"upgrades" json:"upgrades,omitempty"`
			Waf *struct {
				AuditLogging *struct {
					Action   *string `tfsdk:"action" json:"action,omitempty"`
					Location *string `tfsdk:"location" json:"location,omitempty"`
				} `tfsdk:"audit_logging" json:"auditLogging,omitempty"`
				ConfigMapRuleSets *[]struct {
					ConfigMapRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
					DataMapKeys *[]string `tfsdk:"data_map_keys" json:"dataMapKeys,omitempty"`
				} `tfsdk:"config_map_rule_sets" json:"configMapRuleSets,omitempty"`
				CoreRuleSet *struct {
					CustomSettingsFile   *string `tfsdk:"custom_settings_file" json:"customSettingsFile,omitempty"`
					CustomSettingsString *string `tfsdk:"custom_settings_string" json:"customSettingsString,omitempty"`
				} `tfsdk:"core_rule_set" json:"coreRuleSet,omitempty"`
				CustomInterventionMessage *string `tfsdk:"custom_intervention_message" json:"customInterventionMessage,omitempty"`
				Disabled                  *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
				RequestHeadersOnly        *bool   `tfsdk:"request_headers_only" json:"requestHeadersOnly,omitempty"`
				ResponseHeadersOnly       *bool   `tfsdk:"response_headers_only" json:"responseHeadersOnly,omitempty"`
				RuleSets                  *[]struct {
					Directory *string   `tfsdk:"directory" json:"directory,omitempty"`
					Files     *[]string `tfsdk:"files" json:"files,omitempty"`
					RuleStr   *string   `tfsdk:"rule_str" json:"ruleStr,omitempty"`
				} `tfsdk:"rule_sets" json:"ruleSets,omitempty"`
			} `tfsdk:"waf" json:"waf,omitempty"`
		} `tfsdk:"options" json:"options,omitempty"`
		TargetRefs *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewaySoloIoRouteOptionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_solo_io_route_option_v1_manifest"
}

func (r *GatewaySoloIoRouteOptionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
					"namespaced_statuses": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"statuses": schema.MapAttribute{
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

					"options": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"append_x_forwarded_host": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auto_host_rewrite": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"buffer_per_route": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"buffer": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"max_request_bytes": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"disabled": schema.BoolAttribute{
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

							"cors": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allow_credentials": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_headers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_methods": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_origin": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allow_origin_regex": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_for_route": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"expose_headers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_age": schema.StringAttribute{
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

							"csrf": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"additional_origins": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ignore_case": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"safe_regex": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"google_re2": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"max_program_size": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(4.294967295e+09),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"regex": schema.StringAttribute{
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

												"suffix": schema.StringAttribute{
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

									"filter_enabled": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"default_value": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"denominator": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"numerator": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4.294967295e+09),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"runtime_key": schema.StringAttribute{
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

									"shadow_enabled": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"default_value": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"denominator": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"numerator": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4.294967295e+09),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"runtime_key": schema.StringAttribute{
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

							"dlp": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"actions": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"custom_action": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"mask_char": schema.StringAttribute{
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

														"percent": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"value": schema.Float64Attribute{
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

														"regex": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"regex_actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"regex": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"subgroup": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.Int64{
																			int64validator.AtLeast(0),
																			int64validator.AtMost(4.294967295e+09),
																		},
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

												"key_value_action": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key_to_mask": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mask_char": schema.StringAttribute{
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

														"percent": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"value": schema.Float64Attribute{
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

												"shadow": schema.BoolAttribute{
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

									"enabled_for": schema.StringAttribute{
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

							"envoy_metadata": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ext_proc": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"async_mode": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"grpc_service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"authority": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ext_proc_server_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"initial_metadata": schema.ListNestedAttribute{
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

																"value": schema.StringAttribute{
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

													"retry_policy": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"num_retries": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(4.294967295e+09),
																},
															},

															"retry_back_off": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"base_interval": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_interval": schema.StringAttribute{
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

													"timeout": schema.StringAttribute{
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

											"metadata_context_namespaces": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"processing_mode": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"request_body_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"request_header_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"request_trailer_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"response_body_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"response_header_mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"response_trailer_mode": schema.StringAttribute{
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

											"request_attributes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"response_attributes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"typed_metadata_context_namespaces": schema.ListAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"extauth": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"config_ref": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
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

									"custom_auth": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"context_extensions": schema.MapAttribute{
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
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable": schema.BoolAttribute{
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

							"extensions": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"configs": schema.MapAttribute{
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

							"faults": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"abort": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_status": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"percentage": schema.Float64Attribute{
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

									"delay": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"fixed_delay": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"percentage": schema.Float64Attribute{
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

							"header_manipulation": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"request_headers_to_add": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"append": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"header": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
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

												"header_secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
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

									"request_headers_to_remove": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_headers_to_add": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"append": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"header": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
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

									"response_headers_to_remove": schema.ListAttribute{
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

							"host_rewrite": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_rewrite_path_regex": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"pattern": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"google_re2": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max_program_size": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4.294967295e+09),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": schema.StringAttribute{
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

									"substitution": schema.StringAttribute{
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

							"idle_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jwt": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"disable": schema.BoolAttribute{
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

							"jwt_staged": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"after_ext_auth": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"disable": schema.BoolAttribute{
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

									"before_ext_auth": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"disable": schema.BoolAttribute{
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

							"lb_hash": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"hash_policies": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cookie": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ttl": schema.StringAttribute{
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

												"header": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_ip": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"terminal": schema.BoolAttribute{
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

							"max_stream_duration": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"grpc_timeout_header_max": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"grpc_timeout_header_offset": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_stream_duration": schema.StringAttribute{
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

							"prefix_rewrite": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate_limit_configs": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"refs": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
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

							"rate_limit_early_configs": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"refs": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
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

							"rate_limit_regular_configs": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"refs": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
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

							"ratelimit": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"include_vh_rate_limits": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"local_ratelimit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"fill_interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_tokens": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"tokens_per_fill": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"rate_limits": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"actions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"destination_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"generic_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
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

															"header_value_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expect_match": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"exact_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"invert_match": schema.BoolAttribute{
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

																				"prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present_match": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"range_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"end": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"start": schema.Int64Attribute{
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

																				"regex_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"suffix_match": schema.StringAttribute{
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

															"metadata": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"metadata_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.ListNestedAttribute{
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

																	"source": schema.StringAttribute{
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

															"remote_address": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_name": schema.StringAttribute{
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

															"source_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"limit": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dynamic_metadata": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"metadata_key": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.ListNestedAttribute{
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

												"set_actions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"destination_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"generic_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
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

															"header_value_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expect_match": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"exact_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"invert_match": schema.BoolAttribute{
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

																				"prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present_match": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"range_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"end": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"start": schema.Int64Attribute{
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

																				"regex_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"suffix_match": schema.StringAttribute{
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

															"metadata": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"metadata_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.ListNestedAttribute{
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

																	"source": schema.StringAttribute{
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

															"remote_address": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_name": schema.StringAttribute{
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

															"source_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
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

							"ratelimit_basic": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"anonymous_limits": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"requests_per_unit": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"unit": schema.StringAttribute{
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

									"authorized_limits": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"requests_per_unit": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"unit": schema.StringAttribute{
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

							"ratelimit_early": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"include_vh_rate_limits": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"local_ratelimit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"fill_interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_tokens": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"tokens_per_fill": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"rate_limits": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"actions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"destination_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"generic_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
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

															"header_value_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expect_match": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"exact_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"invert_match": schema.BoolAttribute{
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

																				"prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present_match": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"range_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"end": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"start": schema.Int64Attribute{
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

																				"regex_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"suffix_match": schema.StringAttribute{
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

															"metadata": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"metadata_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.ListNestedAttribute{
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

																	"source": schema.StringAttribute{
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

															"remote_address": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_name": schema.StringAttribute{
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

															"source_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"limit": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dynamic_metadata": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"metadata_key": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.ListNestedAttribute{
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

												"set_actions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"destination_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"generic_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
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

															"header_value_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expect_match": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"exact_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"invert_match": schema.BoolAttribute{
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

																				"prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present_match": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"range_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"end": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"start": schema.Int64Attribute{
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

																				"regex_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"suffix_match": schema.StringAttribute{
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

															"metadata": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"metadata_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.ListNestedAttribute{
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

																	"source": schema.StringAttribute{
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

															"remote_address": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_name": schema.StringAttribute{
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

															"source_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
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

							"ratelimit_regular": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"include_vh_rate_limits": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"local_ratelimit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"fill_interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_tokens": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"tokens_per_fill": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"rate_limits": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"actions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"destination_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"generic_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
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

															"header_value_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expect_match": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"exact_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"invert_match": schema.BoolAttribute{
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

																				"prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present_match": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"range_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"end": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"start": schema.Int64Attribute{
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

																				"regex_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"suffix_match": schema.StringAttribute{
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

															"metadata": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"metadata_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.ListNestedAttribute{
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

																	"source": schema.StringAttribute{
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

															"remote_address": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_name": schema.StringAttribute{
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

															"source_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"limit": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dynamic_metadata": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"metadata_key": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.ListNestedAttribute{
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

												"set_actions": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"destination_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"generic_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
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

															"header_value_match": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expect_match": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"exact_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"invert_match": schema.BoolAttribute{
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

																				"prefix_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"present_match": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"range_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"end": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"start": schema.Int64Attribute{
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

																				"regex_match": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"suffix_match": schema.StringAttribute{
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

															"metadata": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"metadata_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.ListNestedAttribute{
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

																	"source": schema.StringAttribute{
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

															"remote_address": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"request_headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"descriptor_key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_name": schema.StringAttribute{
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

															"source_cluster": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
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

							"rbac": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"disable": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"policies": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"nested_claim_delimiter": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"permissions": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"methods": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path_prefix": schema.StringAttribute{
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

											"principals": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"jwt_principal": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"claims": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"matcher": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"provider": schema.StringAttribute{
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

							"regex_rewrite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"pattern": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"google_re2": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max_program_size": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4.294967295e+09),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": schema.StringAttribute{
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

									"substitution": schema.StringAttribute{
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

							"retries": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"num_retries": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(4.294967295e+09),
										},
									},

									"per_try_timeout": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"previous_priorities": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"update_frequency": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_back_off": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"base_interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_interval": schema.StringAttribute{
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

									"retry_on": schema.StringAttribute{
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

							"shadowing": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"percentage": schema.Float64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"upstream": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
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

							"staged_transformations": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"early": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"request_transforms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"clear_route_cache": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"matcher": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"case_sensitive": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"connect_matcher": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"exact": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"invert_match": schema.BoolAttribute{
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

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

																"methods": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"query_parameters": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

																"regex": schema.StringAttribute{
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

														"request_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_transforms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"matchers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"invert_match": schema.BoolAttribute{
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

																	"regex": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

														"response_code_details": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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

									"escape_characters": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"inherit_transformation": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_request_response_info": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"post_routing": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"request_transforms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"clear_route_cache": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"matcher": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"case_sensitive": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"connect_matcher": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"exact": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"invert_match": schema.BoolAttribute{
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

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

																"methods": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"query_parameters": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

																"regex": schema.StringAttribute{
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

														"request_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_transforms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"matchers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"invert_match": schema.BoolAttribute{
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

																	"regex": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

														"response_code_details": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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

									"regular": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"request_transforms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"clear_route_cache": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"matcher": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"case_sensitive": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"connect_matcher": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"exact": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"invert_match": schema.BoolAttribute{
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

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

																"methods": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"query_parameters": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

																"regex": schema.StringAttribute{
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

														"request_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_transforms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"matchers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"invert_match": schema.BoolAttribute{
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

																	"regex": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

														"response_code_details": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
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

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"json_to_proto": schema.BoolAttribute{
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

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mode": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replacement_text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(0),
																						int64validator.AtMost(4.294967295e+09),
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
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

																		"headers_to_append": schema.ListNestedAttribute{
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

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"passthrough": schema.MapAttribute{
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

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"xslt": schema.StringAttribute{
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

							"timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tracing": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"propagate": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"route_descriptor": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"trace_percentages": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"client_sample_percentage": schema.Float64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"overall_sample_percentage": schema.Float64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"random_sample_percentage": schema.Float64Attribute{
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

							"transformations": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"clear_route_cache": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_transformation": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"header_body_transform": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_request_metadata": schema.BoolAttribute{
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

											"log_request_response_info": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"transformation_template": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"advanced_templates": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"body": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"dynamic_metadata_values": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"json_to_proto": schema.BoolAttribute{
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

																"metadata_namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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

													"escape_characters": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"extractors": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"header": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"replacement_text": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subgroup": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(4.294967295e+09),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"headers_to_append": schema.ListNestedAttribute{
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

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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

													"headers_to_remove": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_error_on_parse": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"merge_extractors_to_body": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"parse_body_behavior": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"passthrough": schema.MapAttribute{
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

											"xslt_transformation": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"non_xml_transform": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set_content_type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"xslt": schema.StringAttribute{
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

									"response_transformation": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"header_body_transform": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"add_request_metadata": schema.BoolAttribute{
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

											"log_request_response_info": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"transformation_template": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"advanced_templates": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"body": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"dynamic_metadata_values": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"json_to_proto": schema.BoolAttribute{
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

																"metadata_namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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

													"escape_characters": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"extractors": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"header": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regex": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"replacement_text": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subgroup": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(4.294967295e+09),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"text": schema.StringAttribute{
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

													"headers_to_append": schema.ListNestedAttribute{
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

																"value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"text": schema.StringAttribute{
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

													"headers_to_remove": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ignore_error_on_parse": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"merge_extractors_to_body": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"parse_body_behavior": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"passthrough": schema.MapAttribute{
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

											"xslt_transformation": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"non_xml_transform": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set_content_type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"xslt": schema.StringAttribute{
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

							"upgrades": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"connect": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
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

										"websocket": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
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

							"waf": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"audit_logging": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"location": schema.StringAttribute{
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

									"config_map_rule_sets": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"config_map_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
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

												"data_map_keys": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"core_rule_set": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"custom_settings_file": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"custom_settings_string": schema.StringAttribute{
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

									"custom_intervention_message": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_headers_only": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_headers_only": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"rule_sets": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"directory": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"files": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rule_str": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
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

								"namespace": schema.StringAttribute{
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
		},
	}
}

func (r *GatewaySoloIoRouteOptionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_solo_io_route_option_v1_manifest")

	var model GatewaySoloIoRouteOptionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.solo.io/v1")
	model.Kind = pointer.String("RouteOption")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
