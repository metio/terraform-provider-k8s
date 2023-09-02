/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gloo_solo_io_v1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GlooSoloIoUpstreamGroupV1Manifest{}
)

func NewGlooSoloIoUpstreamGroupV1Manifest() datasource.DataSource {
	return &GlooSoloIoUpstreamGroupV1Manifest{}
}

type GlooSoloIoUpstreamGroupV1Manifest struct{}

type GlooSoloIoUpstreamGroupV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Destinations *[]struct {
			Destination *struct {
				Consul *struct {
					DataCenters *[]string `tfsdk:"data_centers" json:"dataCenters,omitempty"`
					ServiceName *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
					Tags        *[]string `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"consul" json:"consul,omitempty"`
				DestinationSpec *struct {
					Aws *struct {
						InvocationStyle        *string `tfsdk:"invocation_style" json:"invocationStyle,omitempty"`
						LogicalName            *string `tfsdk:"logical_name" json:"logicalName,omitempty"`
						RequestTransformation  *bool   `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
						ResponseTransformation *bool   `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
						UnwrapAsAlb            *bool   `tfsdk:"unwrap_as_alb" json:"unwrapAsAlb,omitempty"`
						UnwrapAsApiGateway     *bool   `tfsdk:"unwrap_as_api_gateway" json:"unwrapAsApiGateway,omitempty"`
						WrapAsApiGateway       *bool   `tfsdk:"wrap_as_api_gateway" json:"wrapAsApiGateway,omitempty"`
					} `tfsdk:"aws" json:"aws,omitempty"`
					Azure *struct {
						FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
					} `tfsdk:"azure" json:"azure,omitempty"`
					Grpc *struct {
						Function   *string `tfsdk:"function" json:"function,omitempty"`
						Package    *string `tfsdk:"package" json:"package,omitempty"`
						Parameters *struct {
							Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
							Path    *string            `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"parameters" json:"parameters,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					Rest *struct {
						FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
						Parameters   *struct {
							Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
							Path    *string            `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"parameters" json:"parameters,omitempty"`
						ResponseTransformation *struct {
							AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
							Body              *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"body" json:"body,omitempty"`
							DynamicMetadataValues *[]struct {
								Key               *string `tfsdk:"key" json:"key,omitempty"`
								MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
								Value             *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
							EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
							Extractors       *struct {
								Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
								Header   *string            `tfsdk:"header" json:"header,omitempty"`
								Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
								Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"rest" json:"rest,omitempty"`
				} `tfsdk:"destination_spec" json:"destinationSpec,omitempty"`
				Kube *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
					Ref  *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ref" json:"ref,omitempty"`
				} `tfsdk:"kube" json:"kube,omitempty"`
				Subset *struct {
					Values *map[string]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"subset" json:"subset,omitempty"`
				Upstream *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"upstream" json:"upstream,omitempty"`
			} `tfsdk:"destination" json:"destination,omitempty"`
			Options *struct {
				BufferPerRoute *struct {
					Buffer *struct {
						MaxRequestBytes *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
					} `tfsdk:"buffer" json:"buffer,omitempty"`
					Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				} `tfsdk:"buffer_per_route" json:"bufferPerRoute,omitempty"`
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
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
					Regular                *struct {
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
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
								Key               *string `tfsdk:"key" json:"key,omitempty"`
								MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
								Value             *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
							EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
							Extractors       *struct {
								Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
								Header   *string            `tfsdk:"header" json:"header,omitempty"`
								Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
								Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
								Key               *string `tfsdk:"key" json:"key,omitempty"`
								MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
								Value             *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
							EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
							Extractors       *struct {
								Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
								Header   *string            `tfsdk:"header" json:"header,omitempty"`
								Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
								Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
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
			} `tfsdk:"options" json:"options,omitempty"`
			Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"destinations" json:"destinations,omitempty"`
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GlooSoloIoUpstreamGroupV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gloo_solo_io_upstream_group_v1_manifest"
}

func (r *GlooSoloIoUpstreamGroupV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
					"destinations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"destination": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"consul": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"data_centers": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tags": schema.ListAttribute{
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

										"destination_spec": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"aws": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"invocation_style": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"logical_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"request_transformation": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"response_transformation": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"unwrap_as_alb": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"unwrap_as_api_gateway": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"wrap_as_api_gateway": schema.BoolAttribute{
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

												"azure": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"function_name": schema.StringAttribute{
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"function": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"package": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"parameters": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"headers": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service": schema.StringAttribute{
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

												"rest": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"function_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"parameters": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"headers": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"response_transformation": schema.SingleNestedAttribute{
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

										"kube": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ref": schema.SingleNestedAttribute{
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

										"subset": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"values": schema.MapAttribute{
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

								"options": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"weight": schema.Int64Attribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *GlooSoloIoUpstreamGroupV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_upstream_group_v1_manifest")

	var model GlooSoloIoUpstreamGroupV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("gloo.solo.io/v1")
	model.Kind = pointer.String("UpstreamGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
