/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type GatewaySoloIoRouteTableV1Resource struct{}

var (
	_ resource.Resource = (*GatewaySoloIoRouteTableV1Resource)(nil)
)

type GatewaySoloIoRouteTableV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GatewaySoloIoRouteTableV1GoModel struct {
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
		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`

		Routes *[]struct {
			DelegateAction *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Ref *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"ref" yaml:"ref,omitempty"`

				Selector *struct {
					Expressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator utilities.IntOrString `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expressions" yaml:"expressions,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`
			} `tfsdk:"delegate_action" yaml:"delegateAction,omitempty"`

			DirectResponseAction *struct {
				Body *string `tfsdk:"body" yaml:"body,omitempty"`

				Status *int64 `tfsdk:"status" yaml:"status,omitempty"`
			} `tfsdk:"direct_response_action" yaml:"directResponseAction,omitempty"`

			GraphqlApiRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"graphql_api_ref" yaml:"graphqlApiRef,omitempty"`

			InheritableMatchers *bool `tfsdk:"inheritable_matchers" yaml:"inheritableMatchers,omitempty"`

			InheritablePathMatchers *bool `tfsdk:"inheritable_path_matchers" yaml:"inheritablePathMatchers,omitempty"`

			Matchers *[]struct {
				CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

				Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

				Headers *[]struct {
					InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				QueryParameters *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

				Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
			} `tfsdk:"matchers" yaml:"matchers,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Options *struct {
				AutoHostRewrite *bool `tfsdk:"auto_host_rewrite" yaml:"autoHostRewrite,omitempty"`

				BufferPerRoute *struct {
					Buffer *struct {
						MaxRequestBytes *int64 `tfsdk:"max_request_bytes" yaml:"maxRequestBytes,omitempty"`
					} `tfsdk:"buffer" yaml:"buffer,omitempty"`

					Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`
				} `tfsdk:"buffer_per_route" yaml:"bufferPerRoute,omitempty"`

				Cors *struct {
					AllowCredentials *bool `tfsdk:"allow_credentials" yaml:"allowCredentials,omitempty"`

					AllowHeaders *[]string `tfsdk:"allow_headers" yaml:"allowHeaders,omitempty"`

					AllowMethods *[]string `tfsdk:"allow_methods" yaml:"allowMethods,omitempty"`

					AllowOrigin *[]string `tfsdk:"allow_origin" yaml:"allowOrigin,omitempty"`

					AllowOriginRegex *[]string `tfsdk:"allow_origin_regex" yaml:"allowOriginRegex,omitempty"`

					DisableForRoute *bool `tfsdk:"disable_for_route" yaml:"disableForRoute,omitempty"`

					ExposeHeaders *[]string `tfsdk:"expose_headers" yaml:"exposeHeaders,omitempty"`

					MaxAge *string `tfsdk:"max_age" yaml:"maxAge,omitempty"`
				} `tfsdk:"cors" yaml:"cors,omitempty"`

				Csrf *struct {
					AdditionalOrigins *[]struct {
						Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

						IgnoreCase *bool `tfsdk:"ignore_case" yaml:"ignoreCase,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						SafeRegex *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"safe_regex" yaml:"safeRegex,omitempty"`

						Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
					} `tfsdk:"additional_origins" yaml:"additionalOrigins,omitempty"`

					FilterEnabled *struct {
						DefaultValue *struct {
							Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

							Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
						} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

						RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
					} `tfsdk:"filter_enabled" yaml:"filterEnabled,omitempty"`

					ShadowEnabled *struct {
						DefaultValue *struct {
							Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

							Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
						} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

						RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
					} `tfsdk:"shadow_enabled" yaml:"shadowEnabled,omitempty"`
				} `tfsdk:"csrf" yaml:"csrf,omitempty"`

				Dlp *struct {
					Actions *[]struct {
						ActionType utilities.IntOrString `tfsdk:"action_type" yaml:"actionType,omitempty"`

						CustomAction *struct {
							MaskChar *string `tfsdk:"mask_char" yaml:"maskChar,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Percent *struct {
								Value utilities.DynamicNumber `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"percent" yaml:"percent,omitempty"`

							Regex *[]string `tfsdk:"regex" yaml:"regex,omitempty"`

							RegexActions *[]struct {
								Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

								Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
							} `tfsdk:"regex_actions" yaml:"regexActions,omitempty"`
						} `tfsdk:"custom_action" yaml:"customAction,omitempty"`

						KeyValueAction *struct {
							KeyToMask *string `tfsdk:"key_to_mask" yaml:"keyToMask,omitempty"`

							MaskChar *string `tfsdk:"mask_char" yaml:"maskChar,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Percent *struct {
								Value utilities.DynamicNumber `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"percent" yaml:"percent,omitempty"`
						} `tfsdk:"key_value_action" yaml:"keyValueAction,omitempty"`

						Shadow *bool `tfsdk:"shadow" yaml:"shadow,omitempty"`
					} `tfsdk:"actions" yaml:"actions,omitempty"`

					EnabledFor utilities.IntOrString `tfsdk:"enabled_for" yaml:"enabledFor,omitempty"`
				} `tfsdk:"dlp" yaml:"dlp,omitempty"`

				EnvoyMetadata utilities.Dynamic `tfsdk:"envoy_metadata" yaml:"envoyMetadata,omitempty"`

				Extauth *struct {
					ConfigRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"config_ref" yaml:"configRef,omitempty"`

					CustomAuth *struct {
						ContextExtensions *map[string]string `tfsdk:"context_extensions" yaml:"contextExtensions,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"custom_auth" yaml:"customAuth,omitempty"`

					Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
				} `tfsdk:"extauth" yaml:"extauth,omitempty"`

				Extensions *struct {
					Configs utilities.Dynamic `tfsdk:"configs" yaml:"configs,omitempty"`
				} `tfsdk:"extensions" yaml:"extensions,omitempty"`

				Faults *struct {
					Abort *struct {
						HttpStatus *int64 `tfsdk:"http_status" yaml:"httpStatus,omitempty"`

						Percentage utilities.DynamicNumber `tfsdk:"percentage" yaml:"percentage,omitempty"`
					} `tfsdk:"abort" yaml:"abort,omitempty"`

					Delay *struct {
						FixedDelay *string `tfsdk:"fixed_delay" yaml:"fixedDelay,omitempty"`

						Percentage utilities.DynamicNumber `tfsdk:"percentage" yaml:"percentage,omitempty"`
					} `tfsdk:"delay" yaml:"delay,omitempty"`
				} `tfsdk:"faults" yaml:"faults,omitempty"`

				HeaderManipulation *struct {
					RequestHeadersToAdd *[]struct {
						Append *bool `tfsdk:"append" yaml:"append,omitempty"`

						Header *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"header" yaml:"header,omitempty"`

						HeaderSecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"header_secret_ref" yaml:"headerSecretRef,omitempty"`
					} `tfsdk:"request_headers_to_add" yaml:"requestHeadersToAdd,omitempty"`

					RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" yaml:"requestHeadersToRemove,omitempty"`

					ResponseHeadersToAdd *[]struct {
						Append *bool `tfsdk:"append" yaml:"append,omitempty"`

						Header *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"header" yaml:"header,omitempty"`
					} `tfsdk:"response_headers_to_add" yaml:"responseHeadersToAdd,omitempty"`

					ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" yaml:"responseHeadersToRemove,omitempty"`
				} `tfsdk:"header_manipulation" yaml:"headerManipulation,omitempty"`

				HostRewrite *string `tfsdk:"host_rewrite" yaml:"hostRewrite,omitempty"`

				Jwt *struct {
					Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
				} `tfsdk:"jwt" yaml:"jwt,omitempty"`

				JwtStaged *struct {
					AfterExtAuth *struct {
						Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
					} `tfsdk:"after_ext_auth" yaml:"afterExtAuth,omitempty"`

					BeforeExtAuth *struct {
						Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
					} `tfsdk:"before_ext_auth" yaml:"beforeExtAuth,omitempty"`
				} `tfsdk:"jwt_staged" yaml:"jwtStaged,omitempty"`

				LbHash *struct {
					HashPolicies *[]struct {
						Cookie *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
						} `tfsdk:"cookie" yaml:"cookie,omitempty"`

						Header *string `tfsdk:"header" yaml:"header,omitempty"`

						SourceIp *bool `tfsdk:"source_ip" yaml:"sourceIp,omitempty"`

						Terminal *bool `tfsdk:"terminal" yaml:"terminal,omitempty"`
					} `tfsdk:"hash_policies" yaml:"hashPolicies,omitempty"`
				} `tfsdk:"lb_hash" yaml:"lbHash,omitempty"`

				MaxStreamDuration *struct {
					GrpcTimeoutHeaderMax *string `tfsdk:"grpc_timeout_header_max" yaml:"grpcTimeoutHeaderMax,omitempty"`

					GrpcTimeoutHeaderOffset *string `tfsdk:"grpc_timeout_header_offset" yaml:"grpcTimeoutHeaderOffset,omitempty"`

					MaxStreamDuration *string `tfsdk:"max_stream_duration" yaml:"maxStreamDuration,omitempty"`
				} `tfsdk:"max_stream_duration" yaml:"maxStreamDuration,omitempty"`

				PrefixRewrite *string `tfsdk:"prefix_rewrite" yaml:"prefixRewrite,omitempty"`

				RateLimitConfigs *struct {
					Refs *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"refs" yaml:"refs,omitempty"`
				} `tfsdk:"rate_limit_configs" yaml:"rateLimitConfigs,omitempty"`

				RateLimitEarlyConfigs *struct {
					Refs *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"refs" yaml:"refs,omitempty"`
				} `tfsdk:"rate_limit_early_configs" yaml:"rateLimitEarlyConfigs,omitempty"`

				RateLimitRegularConfigs *struct {
					Refs *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"refs" yaml:"refs,omitempty"`
				} `tfsdk:"rate_limit_regular_configs" yaml:"rateLimitRegularConfigs,omitempty"`

				Ratelimit *struct {
					IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" yaml:"includeVhRateLimits,omitempty"`

					RateLimits *[]struct {
						Actions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

							GenericKey *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

								ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

								Headers *[]struct {
									ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

									PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

									RangeMatch *struct {
										End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

										Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
									} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

									RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

									SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`
							} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

							Metadata *struct {
								DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								MetadataKey *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Path *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

								Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
							} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

							SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
						} `tfsdk:"actions" yaml:"actions,omitempty"`

						SetActions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

							GenericKey *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

								ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

								Headers *[]struct {
									ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

									PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

									RangeMatch *struct {
										End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

										Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
									} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

									RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

									SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`
							} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

							Metadata *struct {
								DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								MetadataKey *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Path *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

								Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
							} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

							SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
						} `tfsdk:"set_actions" yaml:"setActions,omitempty"`
					} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
				} `tfsdk:"ratelimit" yaml:"ratelimit,omitempty"`

				RatelimitBasic *struct {
					AnonymousLimits *struct {
						RequestsPerUnit *int64 `tfsdk:"requests_per_unit" yaml:"requestsPerUnit,omitempty"`

						Unit utilities.IntOrString `tfsdk:"unit" yaml:"unit,omitempty"`
					} `tfsdk:"anonymous_limits" yaml:"anonymousLimits,omitempty"`

					AuthorizedLimits *struct {
						RequestsPerUnit *int64 `tfsdk:"requests_per_unit" yaml:"requestsPerUnit,omitempty"`

						Unit utilities.IntOrString `tfsdk:"unit" yaml:"unit,omitempty"`
					} `tfsdk:"authorized_limits" yaml:"authorizedLimits,omitempty"`
				} `tfsdk:"ratelimit_basic" yaml:"ratelimitBasic,omitempty"`

				RatelimitEarly *struct {
					IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" yaml:"includeVhRateLimits,omitempty"`

					RateLimits *[]struct {
						Actions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

							GenericKey *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

								ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

								Headers *[]struct {
									ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

									PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

									RangeMatch *struct {
										End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

										Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
									} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

									RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

									SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`
							} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

							Metadata *struct {
								DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								MetadataKey *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Path *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

								Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
							} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

							SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
						} `tfsdk:"actions" yaml:"actions,omitempty"`

						SetActions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

							GenericKey *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

								ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

								Headers *[]struct {
									ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

									PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

									RangeMatch *struct {
										End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

										Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
									} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

									RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

									SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`
							} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

							Metadata *struct {
								DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								MetadataKey *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Path *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

								Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
							} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

							SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
						} `tfsdk:"set_actions" yaml:"setActions,omitempty"`
					} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
				} `tfsdk:"ratelimit_early" yaml:"ratelimitEarly,omitempty"`

				RatelimitRegular *struct {
					IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" yaml:"includeVhRateLimits,omitempty"`

					RateLimits *[]struct {
						Actions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

							GenericKey *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

								ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

								Headers *[]struct {
									ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

									PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

									RangeMatch *struct {
										End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

										Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
									} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

									RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

									SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`
							} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

							Metadata *struct {
								DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								MetadataKey *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Path *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

								Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
							} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

							SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
						} `tfsdk:"actions" yaml:"actions,omitempty"`

						SetActions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

							GenericKey *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

								ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

								Headers *[]struct {
									ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

									PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

									RangeMatch *struct {
										End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

										Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
									} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

									RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

									SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`
							} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

							Metadata *struct {
								DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								MetadataKey *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Path *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

								Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

								HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
							} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

							SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
						} `tfsdk:"set_actions" yaml:"setActions,omitempty"`
					} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
				} `tfsdk:"ratelimit_regular" yaml:"ratelimitRegular,omitempty"`

				Rbac *struct {
					Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`

					Policies *struct {
						NestedClaimDelimiter *string `tfsdk:"nested_claim_delimiter" yaml:"nestedClaimDelimiter,omitempty"`

						Permissions *struct {
							Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

							PathPrefix *string `tfsdk:"path_prefix" yaml:"pathPrefix,omitempty"`
						} `tfsdk:"permissions" yaml:"permissions,omitempty"`

						Principals *[]struct {
							JwtPrincipal *struct {
								Claims *map[string]string `tfsdk:"claims" yaml:"claims,omitempty"`

								Matcher utilities.IntOrString `tfsdk:"matcher" yaml:"matcher,omitempty"`

								Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`
							} `tfsdk:"jwt_principal" yaml:"jwtPrincipal,omitempty"`
						} `tfsdk:"principals" yaml:"principals,omitempty"`
					} `tfsdk:"policies" yaml:"policies,omitempty"`
				} `tfsdk:"rbac" yaml:"rbac,omitempty"`

				RegexRewrite *struct {
					Pattern *struct {
						GoogleRe2 *struct {
							MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
						} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

						Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
					} `tfsdk:"pattern" yaml:"pattern,omitempty"`

					Substitution *string `tfsdk:"substitution" yaml:"substitution,omitempty"`
				} `tfsdk:"regex_rewrite" yaml:"regexRewrite,omitempty"`

				Retries *struct {
					NumRetries *int64 `tfsdk:"num_retries" yaml:"numRetries,omitempty"`

					PerTryTimeout *string `tfsdk:"per_try_timeout" yaml:"perTryTimeout,omitempty"`

					RetryOn *string `tfsdk:"retry_on" yaml:"retryOn,omitempty"`
				} `tfsdk:"retries" yaml:"retries,omitempty"`

				Shadowing *struct {
					Percentage utilities.DynamicNumber `tfsdk:"percentage" yaml:"percentage,omitempty"`

					Upstream *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"upstream" yaml:"upstream,omitempty"`
				} `tfsdk:"shadowing" yaml:"shadowing,omitempty"`

				StagedTransformations *struct {
					Early *struct {
						RequestTransforms *[]struct {
							ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

							Matcher *struct {
								CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

								Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

								Headers *[]struct {
									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

								Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

								QueryParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

								Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
							} `tfsdk:"matcher" yaml:"matcher,omitempty"`

							RequestTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

									Body *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"body" yaml:"body,omitempty"`

									DynamicMetadataValues *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

									Extractors *struct {
										Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

										Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
									} `tfsdk:"extractors" yaml:"extractors,omitempty"`

									Headers *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"headers" yaml:"headers,omitempty"`

									HeadersToAppend *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

									HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

									IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

									ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

									Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

								XsltTransformation *struct {
									NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

									SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

									Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
							} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

									Body *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"body" yaml:"body,omitempty"`

									DynamicMetadataValues *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

									Extractors *struct {
										Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

										Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
									} `tfsdk:"extractors" yaml:"extractors,omitempty"`

									Headers *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"headers" yaml:"headers,omitempty"`

									HeadersToAppend *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

									HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

									IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

									ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

									Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

								XsltTransformation *struct {
									NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

									SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

									Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
						} `tfsdk:"request_transforms" yaml:"requestTransforms,omitempty"`

						ResponseTransforms *[]struct {
							Matchers *[]struct {
								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"matchers" yaml:"matchers,omitempty"`

							ResponseCodeDetails *string `tfsdk:"response_code_details" yaml:"responseCodeDetails,omitempty"`

							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

									Body *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"body" yaml:"body,omitempty"`

									DynamicMetadataValues *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

									Extractors *struct {
										Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

										Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
									} `tfsdk:"extractors" yaml:"extractors,omitempty"`

									Headers *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"headers" yaml:"headers,omitempty"`

									HeadersToAppend *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

									HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

									IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

									ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

									Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

								XsltTransformation *struct {
									NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

									SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

									Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
						} `tfsdk:"response_transforms" yaml:"responseTransforms,omitempty"`
					} `tfsdk:"early" yaml:"early,omitempty"`

					InheritTransformation *bool `tfsdk:"inherit_transformation" yaml:"inheritTransformation,omitempty"`

					Regular *struct {
						RequestTransforms *[]struct {
							ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

							Matcher *struct {
								CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

								Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

								Headers *[]struct {
									InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

								Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

								QueryParameters *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

								Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
							} `tfsdk:"matcher" yaml:"matcher,omitempty"`

							RequestTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

									Body *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"body" yaml:"body,omitempty"`

									DynamicMetadataValues *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

									Extractors *struct {
										Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

										Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
									} `tfsdk:"extractors" yaml:"extractors,omitempty"`

									Headers *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"headers" yaml:"headers,omitempty"`

									HeadersToAppend *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

									HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

									IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

									ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

									Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

								XsltTransformation *struct {
									NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

									SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

									Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
							} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

									Body *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"body" yaml:"body,omitempty"`

									DynamicMetadataValues *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

									Extractors *struct {
										Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

										Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
									} `tfsdk:"extractors" yaml:"extractors,omitempty"`

									Headers *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"headers" yaml:"headers,omitempty"`

									HeadersToAppend *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

									HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

									IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

									ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

									Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

								XsltTransformation *struct {
									NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

									SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

									Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
						} `tfsdk:"request_transforms" yaml:"requestTransforms,omitempty"`

						ResponseTransforms *[]struct {
							Matchers *[]struct {
								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"matchers" yaml:"matchers,omitempty"`

							ResponseCodeDetails *string `tfsdk:"response_code_details" yaml:"responseCodeDetails,omitempty"`

							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

									Body *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"body" yaml:"body,omitempty"`

									DynamicMetadataValues *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

									Extractors *struct {
										Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

										Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
									} `tfsdk:"extractors" yaml:"extractors,omitempty"`

									Headers *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"headers" yaml:"headers,omitempty"`

									HeadersToAppend *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

									HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

									IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

									ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

									Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

								XsltTransformation *struct {
									NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

									SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

									Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
						} `tfsdk:"response_transforms" yaml:"responseTransforms,omitempty"`
					} `tfsdk:"regular" yaml:"regular,omitempty"`
				} `tfsdk:"staged_transformations" yaml:"stagedTransformations,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

				Tracing *struct {
					Propagate *bool `tfsdk:"propagate" yaml:"propagate,omitempty"`

					RouteDescriptor *string `tfsdk:"route_descriptor" yaml:"routeDescriptor,omitempty"`

					TracePercentages *struct {
						ClientSamplePercentage utilities.DynamicNumber `tfsdk:"client_sample_percentage" yaml:"clientSamplePercentage,omitempty"`

						OverallSamplePercentage utilities.DynamicNumber `tfsdk:"overall_sample_percentage" yaml:"overallSamplePercentage,omitempty"`

						RandomSamplePercentage utilities.DynamicNumber `tfsdk:"random_sample_percentage" yaml:"randomSamplePercentage,omitempty"`
					} `tfsdk:"trace_percentages" yaml:"tracePercentages,omitempty"`
				} `tfsdk:"tracing" yaml:"tracing,omitempty"`

				Transformations *struct {
					ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

					RequestTransformation *struct {
						HeaderBodyTransform *struct {
							AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
						} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

						TransformationTemplate *struct {
							AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

							Body *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"body" yaml:"body,omitempty"`

							DynamicMetadataValues *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

								Value *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

							Extractors *struct {
								Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

								Header *string `tfsdk:"header" yaml:"header,omitempty"`

								Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

								Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
							} `tfsdk:"extractors" yaml:"extractors,omitempty"`

							Headers *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							HeadersToAppend *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Value *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

							HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

							IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

							MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

							ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

							Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
						} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

						XsltTransformation *struct {
							NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

							SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

							Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
						} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
					} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

					ResponseTransformation *struct {
						HeaderBodyTransform *struct {
							AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
						} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

						TransformationTemplate *struct {
							AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

							Body *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"body" yaml:"body,omitempty"`

							DynamicMetadataValues *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

								Value *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

							Extractors *struct {
								Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

								Header *string `tfsdk:"header" yaml:"header,omitempty"`

								Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

								Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
							} `tfsdk:"extractors" yaml:"extractors,omitempty"`

							Headers *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							HeadersToAppend *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Value *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

							HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

							IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

							MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

							ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

							Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
						} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

						XsltTransformation *struct {
							NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

							SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

							Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
						} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
					} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
				} `tfsdk:"transformations" yaml:"transformations,omitempty"`

				Upgrades *[]struct {
					Websocket *struct {
						Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
					} `tfsdk:"websocket" yaml:"websocket,omitempty"`
				} `tfsdk:"upgrades" yaml:"upgrades,omitempty"`

				Waf *struct {
					AuditLogging *struct {
						Action utilities.IntOrString `tfsdk:"action" yaml:"action,omitempty"`

						Location utilities.IntOrString `tfsdk:"location" yaml:"location,omitempty"`
					} `tfsdk:"audit_logging" yaml:"auditLogging,omitempty"`

					ConfigMapRuleSets *[]struct {
						ConfigMapRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						DataMapKeys *[]string `tfsdk:"data_map_keys" yaml:"dataMapKeys,omitempty"`
					} `tfsdk:"config_map_rule_sets" yaml:"configMapRuleSets,omitempty"`

					CoreRuleSet *struct {
						CustomSettingsFile *string `tfsdk:"custom_settings_file" yaml:"customSettingsFile,omitempty"`

						CustomSettingsString *string `tfsdk:"custom_settings_string" yaml:"customSettingsString,omitempty"`
					} `tfsdk:"core_rule_set" yaml:"coreRuleSet,omitempty"`

					CustomInterventionMessage *string `tfsdk:"custom_intervention_message" yaml:"customInterventionMessage,omitempty"`

					Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

					RequestHeadersOnly *bool `tfsdk:"request_headers_only" yaml:"requestHeadersOnly,omitempty"`

					ResponseHeadersOnly *bool `tfsdk:"response_headers_only" yaml:"responseHeadersOnly,omitempty"`

					RuleSets *[]struct {
						Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

						Files *[]string `tfsdk:"files" yaml:"files,omitempty"`

						RuleStr *string `tfsdk:"rule_str" yaml:"ruleStr,omitempty"`
					} `tfsdk:"rule_sets" yaml:"ruleSets,omitempty"`
				} `tfsdk:"waf" yaml:"waf,omitempty"`
			} `tfsdk:"options" yaml:"options,omitempty"`

			OptionsConfigRefs *struct {
				DelegateOptions *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"delegate_options" yaml:"delegateOptions,omitempty"`
			} `tfsdk:"options_config_refs" yaml:"optionsConfigRefs,omitempty"`

			RedirectAction *struct {
				HostRedirect *string `tfsdk:"host_redirect" yaml:"hostRedirect,omitempty"`

				HttpsRedirect *bool `tfsdk:"https_redirect" yaml:"httpsRedirect,omitempty"`

				PathRedirect *string `tfsdk:"path_redirect" yaml:"pathRedirect,omitempty"`

				PrefixRewrite *string `tfsdk:"prefix_rewrite" yaml:"prefixRewrite,omitempty"`

				RegexRewrite *struct {
					Pattern *struct {
						GoogleRe2 *struct {
							MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
						} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

						Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
					} `tfsdk:"pattern" yaml:"pattern,omitempty"`

					Substitution *string `tfsdk:"substitution" yaml:"substitution,omitempty"`
				} `tfsdk:"regex_rewrite" yaml:"regexRewrite,omitempty"`

				ResponseCode utilities.IntOrString `tfsdk:"response_code" yaml:"responseCode,omitempty"`

				StripQuery *bool `tfsdk:"strip_query" yaml:"stripQuery,omitempty"`
			} `tfsdk:"redirect_action" yaml:"redirectAction,omitempty"`

			RouteAction *struct {
				ClusterHeader *string `tfsdk:"cluster_header" yaml:"clusterHeader,omitempty"`

				DynamicForwardProxy *struct {
					AutoHostRewriteHeader *string `tfsdk:"auto_host_rewrite_header" yaml:"autoHostRewriteHeader,omitempty"`

					HostRewrite *string `tfsdk:"host_rewrite" yaml:"hostRewrite,omitempty"`
				} `tfsdk:"dynamic_forward_proxy" yaml:"dynamicForwardProxy,omitempty"`

				Multi *struct {
					Destinations *[]struct {
						Destination *struct {
							Consul *struct {
								DataCenters *[]string `tfsdk:"data_centers" yaml:"dataCenters,omitempty"`

								ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

								Tags *[]string `tfsdk:"tags" yaml:"tags,omitempty"`
							} `tfsdk:"consul" yaml:"consul,omitempty"`

							DestinationSpec *struct {
								Aws *struct {
									InvocationStyle utilities.IntOrString `tfsdk:"invocation_style" yaml:"invocationStyle,omitempty"`

									LogicalName *string `tfsdk:"logical_name" yaml:"logicalName,omitempty"`

									RequestTransformation *bool `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

									ResponseTransformation *bool `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`

									UnwrapAsAlb *bool `tfsdk:"unwrap_as_alb" yaml:"unwrapAsAlb,omitempty"`

									UnwrapAsApiGateway *bool `tfsdk:"unwrap_as_api_gateway" yaml:"unwrapAsApiGateway,omitempty"`

									WrapAsApiGateway *bool `tfsdk:"wrap_as_api_gateway" yaml:"wrapAsApiGateway,omitempty"`
								} `tfsdk:"aws" yaml:"aws,omitempty"`

								Azure *struct {
									FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`
								} `tfsdk:"azure" yaml:"azure,omitempty"`

								Grpc *struct {
									Function *string `tfsdk:"function" yaml:"function,omitempty"`

									Package *string `tfsdk:"package" yaml:"package,omitempty"`

									Parameters *struct {
										Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

										Path *string `tfsdk:"path" yaml:"path,omitempty"`
									} `tfsdk:"parameters" yaml:"parameters,omitempty"`

									Service *string `tfsdk:"service" yaml:"service,omitempty"`
								} `tfsdk:"grpc" yaml:"grpc,omitempty"`

								Rest *struct {
									FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`

									Parameters *struct {
										Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

										Path *string `tfsdk:"path" yaml:"path,omitempty"`
									} `tfsdk:"parameters" yaml:"parameters,omitempty"`

									ResponseTransformation *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

										Body *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"body" yaml:"body,omitempty"`

										DynamicMetadataValues *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

											Value *struct {
												Text *string `tfsdk:"text" yaml:"text,omitempty"`
											} `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

										Extractors *struct {
											Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

											Header *string `tfsdk:"header" yaml:"header,omitempty"`

											Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

											Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
										} `tfsdk:"extractors" yaml:"extractors,omitempty"`

										Headers *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"headers" yaml:"headers,omitempty"`

										HeadersToAppend *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Value *struct {
												Text *string `tfsdk:"text" yaml:"text,omitempty"`
											} `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

										HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

										IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

										ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

										Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
									} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
								} `tfsdk:"rest" yaml:"rest,omitempty"`
							} `tfsdk:"destination_spec" yaml:"destinationSpec,omitempty"`

							Kube *struct {
								Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

								Ref *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"ref" yaml:"ref,omitempty"`
							} `tfsdk:"kube" yaml:"kube,omitempty"`

							Subset *struct {
								Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"subset" yaml:"subset,omitempty"`

							Upstream *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"upstream" yaml:"upstream,omitempty"`
						} `tfsdk:"destination" yaml:"destination,omitempty"`

						Options *struct {
							BufferPerRoute *struct {
								Buffer *struct {
									MaxRequestBytes *int64 `tfsdk:"max_request_bytes" yaml:"maxRequestBytes,omitempty"`
								} `tfsdk:"buffer" yaml:"buffer,omitempty"`

								Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`
							} `tfsdk:"buffer_per_route" yaml:"bufferPerRoute,omitempty"`

							Csrf *struct {
								AdditionalOrigins *[]struct {
									Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

									IgnoreCase *bool `tfsdk:"ignore_case" yaml:"ignoreCase,omitempty"`

									Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

									SafeRegex *struct {
										GoogleRe2 *struct {
											MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
										} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

										Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
									} `tfsdk:"safe_regex" yaml:"safeRegex,omitempty"`

									Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
								} `tfsdk:"additional_origins" yaml:"additionalOrigins,omitempty"`

								FilterEnabled *struct {
									DefaultValue *struct {
										Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

										Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
									} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

									RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
								} `tfsdk:"filter_enabled" yaml:"filterEnabled,omitempty"`

								ShadowEnabled *struct {
									DefaultValue *struct {
										Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

										Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
									} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

									RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
								} `tfsdk:"shadow_enabled" yaml:"shadowEnabled,omitempty"`
							} `tfsdk:"csrf" yaml:"csrf,omitempty"`

							Extauth *struct {
								ConfigRef *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"config_ref" yaml:"configRef,omitempty"`

								CustomAuth *struct {
									ContextExtensions *map[string]string `tfsdk:"context_extensions" yaml:"contextExtensions,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`
								} `tfsdk:"custom_auth" yaml:"customAuth,omitempty"`

								Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
							} `tfsdk:"extauth" yaml:"extauth,omitempty"`

							Extensions *struct {
								Configs utilities.Dynamic `tfsdk:"configs" yaml:"configs,omitempty"`
							} `tfsdk:"extensions" yaml:"extensions,omitempty"`

							HeaderManipulation *struct {
								RequestHeadersToAdd *[]struct {
									Append *bool `tfsdk:"append" yaml:"append,omitempty"`

									Header *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"header" yaml:"header,omitempty"`

									HeaderSecretRef *struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
									} `tfsdk:"header_secret_ref" yaml:"headerSecretRef,omitempty"`
								} `tfsdk:"request_headers_to_add" yaml:"requestHeadersToAdd,omitempty"`

								RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" yaml:"requestHeadersToRemove,omitempty"`

								ResponseHeadersToAdd *[]struct {
									Append *bool `tfsdk:"append" yaml:"append,omitempty"`

									Header *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"header" yaml:"header,omitempty"`
								} `tfsdk:"response_headers_to_add" yaml:"responseHeadersToAdd,omitempty"`

								ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" yaml:"responseHeadersToRemove,omitempty"`
							} `tfsdk:"header_manipulation" yaml:"headerManipulation,omitempty"`

							StagedTransformations *struct {
								Early *struct {
									RequestTransforms *[]struct {
										ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

										Matcher *struct {
											CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

											Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

											Headers *[]struct {
												InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"headers" yaml:"headers,omitempty"`

											Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

											Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

											QueryParameters *[]struct {
												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

											Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
										} `tfsdk:"matcher" yaml:"matcher,omitempty"`

										RequestTransformation *struct {
											HeaderBodyTransform *struct {
												AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
											} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

											TransformationTemplate *struct {
												AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

												Body *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"body" yaml:"body,omitempty"`

												DynamicMetadataValues *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

												Extractors *struct {
													Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

													Header *string `tfsdk:"header" yaml:"header,omitempty"`

													Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

													Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
												} `tfsdk:"extractors" yaml:"extractors,omitempty"`

												Headers *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"headers" yaml:"headers,omitempty"`

												HeadersToAppend *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

												HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

												IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

												MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

												ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

												Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
											} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

											XsltTransformation *struct {
												NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

												SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

												Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
											} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
										} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

										ResponseTransformation *struct {
											HeaderBodyTransform *struct {
												AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
											} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

											TransformationTemplate *struct {
												AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

												Body *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"body" yaml:"body,omitempty"`

												DynamicMetadataValues *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

												Extractors *struct {
													Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

													Header *string `tfsdk:"header" yaml:"header,omitempty"`

													Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

													Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
												} `tfsdk:"extractors" yaml:"extractors,omitempty"`

												Headers *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"headers" yaml:"headers,omitempty"`

												HeadersToAppend *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

												HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

												IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

												MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

												ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

												Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
											} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

											XsltTransformation *struct {
												NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

												SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

												Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
											} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
										} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
									} `tfsdk:"request_transforms" yaml:"requestTransforms,omitempty"`

									ResponseTransforms *[]struct {
										Matchers *[]struct {
											InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"matchers" yaml:"matchers,omitempty"`

										ResponseCodeDetails *string `tfsdk:"response_code_details" yaml:"responseCodeDetails,omitempty"`

										ResponseTransformation *struct {
											HeaderBodyTransform *struct {
												AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
											} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

											TransformationTemplate *struct {
												AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

												Body *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"body" yaml:"body,omitempty"`

												DynamicMetadataValues *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

												Extractors *struct {
													Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

													Header *string `tfsdk:"header" yaml:"header,omitempty"`

													Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

													Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
												} `tfsdk:"extractors" yaml:"extractors,omitempty"`

												Headers *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"headers" yaml:"headers,omitempty"`

												HeadersToAppend *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

												HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

												IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

												MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

												ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

												Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
											} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

											XsltTransformation *struct {
												NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

												SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

												Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
											} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
										} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
									} `tfsdk:"response_transforms" yaml:"responseTransforms,omitempty"`
								} `tfsdk:"early" yaml:"early,omitempty"`

								InheritTransformation *bool `tfsdk:"inherit_transformation" yaml:"inheritTransformation,omitempty"`

								Regular *struct {
									RequestTransforms *[]struct {
										ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

										Matcher *struct {
											CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

											Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

											Headers *[]struct {
												InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"headers" yaml:"headers,omitempty"`

											Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

											Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

											QueryParameters *[]struct {
												Name *string `tfsdk:"name" yaml:"name,omitempty"`

												Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

												Value *string `tfsdk:"value" yaml:"value,omitempty"`
											} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

											Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
										} `tfsdk:"matcher" yaml:"matcher,omitempty"`

										RequestTransformation *struct {
											HeaderBodyTransform *struct {
												AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
											} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

											TransformationTemplate *struct {
												AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

												Body *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"body" yaml:"body,omitempty"`

												DynamicMetadataValues *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

												Extractors *struct {
													Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

													Header *string `tfsdk:"header" yaml:"header,omitempty"`

													Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

													Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
												} `tfsdk:"extractors" yaml:"extractors,omitempty"`

												Headers *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"headers" yaml:"headers,omitempty"`

												HeadersToAppend *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

												HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

												IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

												MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

												ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

												Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
											} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

											XsltTransformation *struct {
												NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

												SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

												Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
											} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
										} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

										ResponseTransformation *struct {
											HeaderBodyTransform *struct {
												AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
											} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

											TransformationTemplate *struct {
												AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

												Body *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"body" yaml:"body,omitempty"`

												DynamicMetadataValues *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

												Extractors *struct {
													Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

													Header *string `tfsdk:"header" yaml:"header,omitempty"`

													Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

													Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
												} `tfsdk:"extractors" yaml:"extractors,omitempty"`

												Headers *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"headers" yaml:"headers,omitempty"`

												HeadersToAppend *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

												HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

												IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

												MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

												ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

												Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
											} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

											XsltTransformation *struct {
												NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

												SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

												Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
											} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
										} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
									} `tfsdk:"request_transforms" yaml:"requestTransforms,omitempty"`

									ResponseTransforms *[]struct {
										Matchers *[]struct {
											InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

											Name *string `tfsdk:"name" yaml:"name,omitempty"`

											Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

											Value *string `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"matchers" yaml:"matchers,omitempty"`

										ResponseCodeDetails *string `tfsdk:"response_code_details" yaml:"responseCodeDetails,omitempty"`

										ResponseTransformation *struct {
											HeaderBodyTransform *struct {
												AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
											} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

											TransformationTemplate *struct {
												AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

												Body *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"body" yaml:"body,omitempty"`

												DynamicMetadataValues *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

												Extractors *struct {
													Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

													Header *string `tfsdk:"header" yaml:"header,omitempty"`

													Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

													Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
												} `tfsdk:"extractors" yaml:"extractors,omitempty"`

												Headers *struct {
													Text *string `tfsdk:"text" yaml:"text,omitempty"`
												} `tfsdk:"headers" yaml:"headers,omitempty"`

												HeadersToAppend *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Value *struct {
														Text *string `tfsdk:"text" yaml:"text,omitempty"`
													} `tfsdk:"value" yaml:"value,omitempty"`
												} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

												HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

												IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

												MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

												ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

												Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
											} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

											XsltTransformation *struct {
												NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

												SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

												Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
											} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
										} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
									} `tfsdk:"response_transforms" yaml:"responseTransforms,omitempty"`
								} `tfsdk:"regular" yaml:"regular,omitempty"`
							} `tfsdk:"staged_transformations" yaml:"stagedTransformations,omitempty"`

							Transformations *struct {
								ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

								RequestTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

										Body *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"body" yaml:"body,omitempty"`

										DynamicMetadataValues *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

											Value *struct {
												Text *string `tfsdk:"text" yaml:"text,omitempty"`
											} `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

										Extractors *struct {
											Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

											Header *string `tfsdk:"header" yaml:"header,omitempty"`

											Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

											Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
										} `tfsdk:"extractors" yaml:"extractors,omitempty"`

										Headers *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"headers" yaml:"headers,omitempty"`

										HeadersToAppend *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Value *struct {
												Text *string `tfsdk:"text" yaml:"text,omitempty"`
											} `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

										HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

										IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

										ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

										Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

									XsltTransformation *struct {
										NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

										SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

										Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
								} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

								ResponseTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

										Body *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"body" yaml:"body,omitempty"`

										DynamicMetadataValues *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

											Value *struct {
												Text *string `tfsdk:"text" yaml:"text,omitempty"`
											} `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

										Extractors *struct {
											Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

											Header *string `tfsdk:"header" yaml:"header,omitempty"`

											Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

											Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
										} `tfsdk:"extractors" yaml:"extractors,omitempty"`

										Headers *struct {
											Text *string `tfsdk:"text" yaml:"text,omitempty"`
										} `tfsdk:"headers" yaml:"headers,omitempty"`

										HeadersToAppend *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Value *struct {
												Text *string `tfsdk:"text" yaml:"text,omitempty"`
											} `tfsdk:"value" yaml:"value,omitempty"`
										} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

										HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

										IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

										ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

										Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

									XsltTransformation *struct {
										NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

										SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

										Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
								} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
							} `tfsdk:"transformations" yaml:"transformations,omitempty"`
						} `tfsdk:"options" yaml:"options,omitempty"`

						Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
					} `tfsdk:"destinations" yaml:"destinations,omitempty"`
				} `tfsdk:"multi" yaml:"multi,omitempty"`

				Single *struct {
					Consul *struct {
						DataCenters *[]string `tfsdk:"data_centers" yaml:"dataCenters,omitempty"`

						ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

						Tags *[]string `tfsdk:"tags" yaml:"tags,omitempty"`
					} `tfsdk:"consul" yaml:"consul,omitempty"`

					DestinationSpec *struct {
						Aws *struct {
							InvocationStyle utilities.IntOrString `tfsdk:"invocation_style" yaml:"invocationStyle,omitempty"`

							LogicalName *string `tfsdk:"logical_name" yaml:"logicalName,omitempty"`

							RequestTransformation *bool `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

							ResponseTransformation *bool `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`

							UnwrapAsAlb *bool `tfsdk:"unwrap_as_alb" yaml:"unwrapAsAlb,omitempty"`

							UnwrapAsApiGateway *bool `tfsdk:"unwrap_as_api_gateway" yaml:"unwrapAsApiGateway,omitempty"`

							WrapAsApiGateway *bool `tfsdk:"wrap_as_api_gateway" yaml:"wrapAsApiGateway,omitempty"`
						} `tfsdk:"aws" yaml:"aws,omitempty"`

						Azure *struct {
							FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`
						} `tfsdk:"azure" yaml:"azure,omitempty"`

						Grpc *struct {
							Function *string `tfsdk:"function" yaml:"function,omitempty"`

							Package *string `tfsdk:"package" yaml:"package,omitempty"`

							Parameters *struct {
								Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"parameters" yaml:"parameters,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

						Rest *struct {
							FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`

							Parameters *struct {
								Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"parameters" yaml:"parameters,omitempty"`

							ResponseTransformation *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
						} `tfsdk:"rest" yaml:"rest,omitempty"`
					} `tfsdk:"destination_spec" yaml:"destinationSpec,omitempty"`

					Kube *struct {
						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						Ref *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
						} `tfsdk:"ref" yaml:"ref,omitempty"`
					} `tfsdk:"kube" yaml:"kube,omitempty"`

					Subset *struct {
						Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"subset" yaml:"subset,omitempty"`

					Upstream *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"upstream" yaml:"upstream,omitempty"`
				} `tfsdk:"single" yaml:"single,omitempty"`

				UpstreamGroup *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"upstream_group" yaml:"upstreamGroup,omitempty"`
			} `tfsdk:"route_action" yaml:"routeAction,omitempty"`
		} `tfsdk:"routes" yaml:"routes,omitempty"`

		Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGatewaySoloIoRouteTableV1Resource() resource.Resource {
	return &GatewaySoloIoRouteTableV1Resource{}
}

func (r *GatewaySoloIoRouteTableV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_solo_io_route_table_v1"
}

func (r *GatewaySoloIoRouteTableV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"namespaced_statuses": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"statuses": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"routes": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"delegate_action": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ref": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

									"selector": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"expressions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operator": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"values": {
														Description:         "",
														MarkdownDescription: "",

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

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"direct_response_action": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"body": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"status": {
										Description:         "",
										MarkdownDescription: "",

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

							"graphql_api_ref": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "",
										MarkdownDescription: "",

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

							"inheritable_matchers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"inheritable_path_matchers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"matchers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"case_sensitive": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"exact": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"invert_match": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

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

									"methods": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_parameters": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

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

									"regex": {
										Description:         "",
										MarkdownDescription: "",

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

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto_host_rewrite": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"buffer_per_route": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"buffer": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"max_request_bytes": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disabled": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cors": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_credentials": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allow_headers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allow_methods": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allow_origin": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allow_origin_regex": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disable_for_route": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expose_headers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_age": {
												Description:         "",
												MarkdownDescription: "",

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

									"csrf": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"additional_origins": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"exact": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_case": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"safe_regex": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"google_re2": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"max_program_size": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			int64validator.AtLeast(0),

																			int64validator.AtMost(4.294967295e+09),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

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

													"suffix": {
														Description:         "",
														MarkdownDescription: "",

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

											"filter_enabled": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_value": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"denominator": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"numerator": {
																Description:         "",
																MarkdownDescription: "",

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

													"runtime_key": {
														Description:         "",
														MarkdownDescription: "",

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

											"shadow_enabled": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_value": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"denominator": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"numerator": {
																Description:         "",
																MarkdownDescription: "",

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

													"runtime_key": {
														Description:         "",
														MarkdownDescription: "",

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

									"dlp": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"action_type": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom_action": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"mask_char": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"percent": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.DynamicNumberType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex_actions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"subgroup": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"key_value_action": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key_to_mask": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"mask_char": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"percent": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.DynamicNumberType{},

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

													"shadow": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled_for": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"envoy_metadata": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"extauth": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

											"custom_auth": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"context_extensions": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

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

											"disable": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"extensions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"configs": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"faults": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"abort": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http_status": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percentage": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicNumberType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delay": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fixed_delay": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percentage": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicNumberType{},

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

									"header_manipulation": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"request_headers_to_add": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"append": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"header": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_secret_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "",
																MarkdownDescription: "",

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

											"request_headers_to_remove": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_headers_to_add": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"append": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"header": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

											"response_headers_to_remove": {
												Description:         "",
												MarkdownDescription: "",

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

									"host_rewrite": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jwt": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"disable": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jwt_staged": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"after_ext_auth": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"disable": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"before_ext_auth": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"disable": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

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

									"lb_hash": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"hash_policies": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"cookie": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ttl": {
																Description:         "",
																MarkdownDescription: "",

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

													"header": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source_ip": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"terminal": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

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

									"max_stream_duration": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"grpc_timeout_header_max": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"grpc_timeout_header_offset": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_stream_duration": {
												Description:         "",
												MarkdownDescription: "",

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

									"prefix_rewrite": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rate_limit_configs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"refs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

									"rate_limit_early_configs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"refs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

									"rate_limit_regular_configs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"refs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

									"ratelimit": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"include_vh_rate_limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rate_limits": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"destination_cluster": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"generic_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"header_value_match": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expect_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"present_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"range_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"end": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"start": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"suffix_match": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remote_address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"source_cluster": {
																Description:         "",
																MarkdownDescription: "",

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

													"set_actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"destination_cluster": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"generic_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"header_value_match": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expect_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"present_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"range_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"end": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"start": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"suffix_match": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remote_address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"source_cluster": {
																Description:         "",
																MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ratelimit_basic": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"anonymous_limits": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"requests_per_unit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"unit": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"authorized_limits": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"requests_per_unit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"unit": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

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

									"ratelimit_early": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"include_vh_rate_limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rate_limits": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"destination_cluster": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"generic_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"header_value_match": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expect_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"present_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"range_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"end": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"start": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"suffix_match": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remote_address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"source_cluster": {
																Description:         "",
																MarkdownDescription: "",

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

													"set_actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"destination_cluster": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"generic_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"header_value_match": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expect_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"present_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"range_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"end": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"start": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"suffix_match": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remote_address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"source_cluster": {
																Description:         "",
																MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ratelimit_regular": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"include_vh_rate_limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rate_limits": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"destination_cluster": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"generic_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"header_value_match": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expect_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"present_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"range_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"end": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"start": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"suffix_match": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remote_address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"source_cluster": {
																Description:         "",
																MarkdownDescription: "",

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

													"set_actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"destination_cluster": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"generic_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"header_value_match": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expect_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"present_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"range_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"end": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"start": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"suffix_match": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"metadata": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"source": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remote_address": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"descriptor_key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header_name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"source_cluster": {
																Description:         "",
																MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rbac": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"disable": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"policies": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"nested_claim_delimiter": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"permissions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"methods": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path_prefix": {
																Description:         "",
																MarkdownDescription: "",

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

													"principals": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"jwt_principal": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"claims": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"matcher": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"provider": {
																		Description:         "",
																		MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"regex_rewrite": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"pattern": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"google_re2": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"max_program_size": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(0),

																	int64validator.AtMost(4.294967295e+09),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regex": {
														Description:         "",
														MarkdownDescription: "",

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

											"substitution": {
												Description:         "",
												MarkdownDescription: "",

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

									"retries": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"num_retries": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"per_try_timeout": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"retry_on": {
												Description:         "",
												MarkdownDescription: "",

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

									"shadowing": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"percentage": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"upstream": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

									"staged_transformations": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"early": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"request_transforms": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"clear_route_cache": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"matcher": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"case_sensitive": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exact": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"methods": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"query_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"request_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"header_body_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"add_request_metadata": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"transformation_template": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"advanced_templates": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"dynamic_metadata_values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"metadata_namespace": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"extractors": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"subgroup": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers_to_append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"headers_to_remove": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_error_on_parse": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"merge_extractors_to_body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parse_body_behavior": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"passthrough": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"xslt_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"non_xml_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"set_content_type": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"xslt": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"response_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"header_body_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"add_request_metadata": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"transformation_template": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"advanced_templates": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"dynamic_metadata_values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"metadata_namespace": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"extractors": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"subgroup": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers_to_append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"headers_to_remove": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_error_on_parse": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"merge_extractors_to_body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parse_body_behavior": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"passthrough": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"xslt_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"non_xml_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"set_content_type": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"xslt": {
																				Description:         "",
																				MarkdownDescription: "",

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

													"response_transforms": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"matchers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"response_code_details": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"response_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"header_body_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"add_request_metadata": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"transformation_template": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"advanced_templates": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"dynamic_metadata_values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"metadata_namespace": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"extractors": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"subgroup": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers_to_append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"headers_to_remove": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_error_on_parse": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"merge_extractors_to_body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parse_body_behavior": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"passthrough": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"xslt_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"non_xml_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"set_content_type": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"xslt": {
																				Description:         "",
																				MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"inherit_transformation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regular": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"request_transforms": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"clear_route_cache": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"matcher": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"case_sensitive": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exact": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"invert_match": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"methods": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"query_parameters": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"request_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"header_body_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"add_request_metadata": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"transformation_template": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"advanced_templates": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"dynamic_metadata_values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"metadata_namespace": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"extractors": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"subgroup": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers_to_append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"headers_to_remove": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_error_on_parse": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"merge_extractors_to_body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parse_body_behavior": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"passthrough": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"xslt_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"non_xml_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"set_content_type": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"xslt": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"response_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"header_body_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"add_request_metadata": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"transformation_template": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"advanced_templates": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"dynamic_metadata_values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"metadata_namespace": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"extractors": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"subgroup": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers_to_append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"headers_to_remove": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_error_on_parse": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"merge_extractors_to_body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parse_body_behavior": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"passthrough": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"xslt_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"non_xml_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"set_content_type": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"xslt": {
																				Description:         "",
																				MarkdownDescription: "",

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

													"response_transforms": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"matchers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"response_code_details": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"response_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"header_body_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"add_request_metadata": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"transformation_template": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"advanced_templates": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"dynamic_metadata_values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"metadata_namespace": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"extractors": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"subgroup": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"headers_to_append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																			"headers_to_remove": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_error_on_parse": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"merge_extractors_to_body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parse_body_behavior": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"passthrough": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"xslt_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"non_xml_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"set_content_type": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"xslt": {
																				Description:         "",
																				MarkdownDescription: "",

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

									"timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tracing": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"propagate": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"route_descriptor": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"trace_percentages": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_sample_percentage": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicNumberType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"overall_sample_percentage": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicNumberType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"random_sample_percentage": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicNumberType{},

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

									"transformations": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"clear_route_cache": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_transformation": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"header_body_transform": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"add_request_metadata": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"transformation_template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"advanced_templates": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"dynamic_metadata_values": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_namespace": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"extractors": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"subgroup": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"headers_to_append": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"headers_to_remove": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ignore_error_on_parse": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"merge_extractors_to_body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parse_body_behavior": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"passthrough": {
																Description:         "",
																MarkdownDescription: "",

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

													"xslt_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"non_xml_transform": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"set_content_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"xslt": {
																Description:         "",
																MarkdownDescription: "",

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

											"response_transformation": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"header_body_transform": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"add_request_metadata": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"transformation_template": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"advanced_templates": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"dynamic_metadata_values": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata_namespace": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"extractors": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"header": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"subgroup": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"headers_to_append": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"headers_to_remove": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ignore_error_on_parse": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"merge_extractors_to_body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parse_body_behavior": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"passthrough": {
																Description:         "",
																MarkdownDescription: "",

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

													"xslt_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"non_xml_transform": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"set_content_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"xslt": {
																Description:         "",
																MarkdownDescription: "",

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

									"upgrades": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"websocket": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"enabled": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

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

									"waf": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"audit_logging": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"action": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"location": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"config_map_rule_sets": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "",
																MarkdownDescription: "",

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

													"data_map_keys": {
														Description:         "",
														MarkdownDescription: "",

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

											"core_rule_set": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"custom_settings_file": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom_settings_string": {
														Description:         "",
														MarkdownDescription: "",

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

											"custom_intervention_message": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disabled": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_headers_only": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_headers_only": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rule_sets": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"directory": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"files": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"rule_str": {
														Description:         "",
														MarkdownDescription: "",

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

							"options_config_refs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"delegate_options": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

							"redirect_action": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host_redirect": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"https_redirect": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path_redirect": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix_rewrite": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"regex_rewrite": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"pattern": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"google_re2": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"max_program_size": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(0),

																	int64validator.AtMost(4.294967295e+09),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regex": {
														Description:         "",
														MarkdownDescription: "",

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

											"substitution": {
												Description:         "",
												MarkdownDescription: "",

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

									"response_code": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strip_query": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route_action": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cluster_header": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dynamic_forward_proxy": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"auto_host_rewrite_header": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_rewrite": {
												Description:         "",
												MarkdownDescription: "",

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

									"multi": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"destinations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"consul": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"data_centers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"service_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"tags": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"destination_spec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"aws": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"invocation_style": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"logical_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"request_transformation": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"response_transformation": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"unwrap_as_alb": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"unwrap_as_api_gateway": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"wrap_as_api_gateway": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"azure": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"function_name": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"grpc": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"function": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"package": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parameters": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"headers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"service": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"rest": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"function_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"parameters": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"headers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"response_transformation": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"advanced_templates": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"dynamic_metadata_values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"metadata_namespace": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"text": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"extractors": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"body": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.MapType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"header": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"subgroup": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"headers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"headers_to_append": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"text": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"headers_to_remove": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"ignore_error_on_parse": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"merge_extractors_to_body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parse_body_behavior": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"passthrough": {
																						Description:         "",
																						MarkdownDescription: "",

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
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kube": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"port": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"subset": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"values": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"upstream": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"options": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"buffer_per_route": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"buffer": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"max_request_bytes": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,

																				Validators: []tfsdk.AttributeValidator{

																					int64validator.AtLeast(0),

																					int64validator.AtMost(4.294967295e+09),
																				},
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"disabled": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"csrf": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"additional_origins": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"exact": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"ignore_case": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"prefix": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"safe_regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"google_re2": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"max_program_size": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,

																								Validators: []tfsdk.AttributeValidator{

																									int64validator.AtLeast(0),

																									int64validator.AtMost(4.294967295e+09),
																								},
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"regex": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"suffix": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"filter_enabled": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"default_value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"denominator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"numerator": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"runtime_key": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"shadow_enabled": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"default_value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"denominator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"numerator": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"runtime_key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"extauth": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"custom_auth": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"context_extensions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"disable": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"extensions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"configs": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_manipulation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"request_headers_to_add": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"header_secret_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"namespace": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"request_headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"response_headers_to_add": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"append": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"value": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"response_headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"staged_transformations": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"early": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"request_transforms": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"clear_route_cache": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"matcher": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"case_sensitive": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"exact": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"headers": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"invert_match": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"regex": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"methods": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"prefix": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"query_parameters": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"regex": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"request_transformation": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"header_body_transform": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"add_request_metadata": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"transformation_template": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"advanced_templates": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"dynamic_metadata_values": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"metadata_namespace": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"extractors": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"body": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"header": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"regex": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"subgroup": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers_to_append": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"headers_to_remove": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"ignore_error_on_parse": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"merge_extractors_to_body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"parse_body_behavior": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: utilities.IntOrStringType{},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"passthrough": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"xslt_transformation": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"non_xml_transform": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"set_content_type": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"xslt": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"response_transformation": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"header_body_transform": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"add_request_metadata": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"transformation_template": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"advanced_templates": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"dynamic_metadata_values": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"metadata_namespace": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"extractors": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"body": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"header": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"regex": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"subgroup": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers_to_append": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"headers_to_remove": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"ignore_error_on_parse": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"merge_extractors_to_body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"parse_body_behavior": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: utilities.IntOrStringType{},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"passthrough": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"xslt_transformation": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"non_xml_transform": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"set_content_type": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"xslt": {
																										Description:         "",
																										MarkdownDescription: "",

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

																			"response_transforms": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"matchers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"invert_match": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"response_code_details": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"response_transformation": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"header_body_transform": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"add_request_metadata": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"transformation_template": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"advanced_templates": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"dynamic_metadata_values": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"metadata_namespace": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"extractors": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"body": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"header": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"regex": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"subgroup": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers_to_append": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"headers_to_remove": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"ignore_error_on_parse": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"merge_extractors_to_body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"parse_body_behavior": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: utilities.IntOrStringType{},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"passthrough": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"xslt_transformation": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"non_xml_transform": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"set_content_type": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"xslt": {
																										Description:         "",
																										MarkdownDescription: "",

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
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"inherit_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regular": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"request_transforms": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"clear_route_cache": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"matcher": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"case_sensitive": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"exact": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"headers": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"invert_match": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"regex": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"methods": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"prefix": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"query_parameters": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"name": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"regex": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"value": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"request_transformation": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"header_body_transform": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"add_request_metadata": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"transformation_template": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"advanced_templates": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"dynamic_metadata_values": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"metadata_namespace": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"extractors": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"body": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"header": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"regex": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"subgroup": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers_to_append": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"headers_to_remove": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"ignore_error_on_parse": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"merge_extractors_to_body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"parse_body_behavior": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: utilities.IntOrStringType{},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"passthrough": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"xslt_transformation": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"non_xml_transform": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"set_content_type": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"xslt": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"response_transformation": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"header_body_transform": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"add_request_metadata": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"transformation_template": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"advanced_templates": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"dynamic_metadata_values": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"metadata_namespace": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"extractors": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"body": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"header": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"regex": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"subgroup": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers_to_append": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"headers_to_remove": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"ignore_error_on_parse": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"merge_extractors_to_body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"parse_body_behavior": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: utilities.IntOrStringType{},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"passthrough": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"xslt_transformation": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"non_xml_transform": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"set_content_type": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"xslt": {
																										Description:         "",
																										MarkdownDescription: "",

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

																			"response_transforms": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"matchers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"invert_match": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.BoolType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"response_code_details": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"response_transformation": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"header_body_transform": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"add_request_metadata": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"transformation_template": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"advanced_templates": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"dynamic_metadata_values": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"metadata_namespace": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"extractors": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"body": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"header": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"regex": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"subgroup": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"text": {
																												Description:         "",
																												MarkdownDescription: "",

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

																									"headers_to_append": {
																										Description:         "",
																										MarkdownDescription: "",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "",
																												MarkdownDescription: "",

																												Type: types.StringType,

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"value": {
																												Description:         "",
																												MarkdownDescription: "",

																												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																													"text": {
																														Description:         "",
																														MarkdownDescription: "",

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

																									"headers_to_remove": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"ignore_error_on_parse": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"merge_extractors_to_body": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"parse_body_behavior": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: utilities.IntOrStringType{},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"passthrough": {
																										Description:         "",
																										MarkdownDescription: "",

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

																							"xslt_transformation": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"non_xml_transform": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.BoolType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"set_content_type": {
																										Description:         "",
																										MarkdownDescription: "",

																										Type: types.StringType,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"xslt": {
																										Description:         "",
																										MarkdownDescription: "",

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

															"transformations": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"clear_route_cache": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"request_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"header_body_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"add_request_metadata": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"transformation_template": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"advanced_templates": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"dynamic_metadata_values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"metadata_namespace": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"text": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"extractors": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"body": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.MapType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"header": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"subgroup": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"headers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"headers_to_append": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"text": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"headers_to_remove": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"ignore_error_on_parse": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"merge_extractors_to_body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parse_body_behavior": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"passthrough": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"xslt_transformation": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"non_xml_transform": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"set_content_type": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"xslt": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"response_transformation": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"header_body_transform": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"add_request_metadata": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"transformation_template": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"advanced_templates": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"dynamic_metadata_values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"metadata_namespace": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"text": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"extractors": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"body": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.MapType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"header": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"regex": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"subgroup": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"headers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"text": {
																								Description:         "",
																								MarkdownDescription: "",

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

																					"headers_to_append": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"value": {
																								Description:         "",
																								MarkdownDescription: "",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"text": {
																										Description:         "",
																										MarkdownDescription: "",

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

																					"headers_to_remove": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"ignore_error_on_parse": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"merge_extractors_to_body": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"parse_body_behavior": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"passthrough": {
																						Description:         "",
																						MarkdownDescription: "",

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

																			"xslt_transformation": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"non_xml_transform": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"set_content_type": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"xslt": {
																						Description:         "",
																						MarkdownDescription: "",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"weight": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
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

									"single": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"consul": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"data_centers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tags": {
														Description:         "",
														MarkdownDescription: "",

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

											"destination_spec": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"aws": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"invocation_style": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"logical_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"response_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"unwrap_as_alb": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"unwrap_as_api_gateway": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"wrap_as_api_gateway": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"azure": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"function_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"grpc": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"function": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"package": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parameters": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"service": {
																Description:         "",
																MarkdownDescription: "",

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

													"rest": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"function_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parameters": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"response_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kube": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace": {
																Description:         "",
																MarkdownDescription: "",

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

											"subset": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"values": {
														Description:         "",
														MarkdownDescription: "",

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

											"upstream": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

									"upstream_group": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

					"weight": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(-2.147483648e+09),

							int64validator.AtMost(2.147483647e+09),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *GatewaySoloIoRouteTableV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gateway_solo_io_route_table_v1")

	var state GatewaySoloIoRouteTableV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewaySoloIoRouteTableV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.solo.io/v1")
	goModel.Kind = utilities.Ptr("RouteTable")

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

func (r *GatewaySoloIoRouteTableV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_solo_io_route_table_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GatewaySoloIoRouteTableV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gateway_solo_io_route_table_v1")

	var state GatewaySoloIoRouteTableV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewaySoloIoRouteTableV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.solo.io/v1")
	goModel.Kind = utilities.Ptr("RouteTable")

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

func (r *GatewaySoloIoRouteTableV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gateway_solo_io_route_table_v1")
	// NO-OP: Terraform removes the state automatically for us
}
