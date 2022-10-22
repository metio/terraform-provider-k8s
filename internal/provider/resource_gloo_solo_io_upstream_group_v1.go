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

type GlooSoloIoUpstreamGroupV1Resource struct{}

var (
	_ resource.Resource = (*GlooSoloIoUpstreamGroupV1Resource)(nil)
)

type GlooSoloIoUpstreamGroupV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GlooSoloIoUpstreamGroupV1GoModel struct {
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

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGlooSoloIoUpstreamGroupV1Resource() resource.Resource {
	return &GlooSoloIoUpstreamGroupV1Resource{}
}

func (r *GlooSoloIoUpstreamGroupV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gloo_solo_io_upstream_group_v1"
}

func (r *GlooSoloIoUpstreamGroupV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *GlooSoloIoUpstreamGroupV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gloo_solo_io_upstream_group_v1")

	var state GlooSoloIoUpstreamGroupV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoUpstreamGroupV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("UpstreamGroup")

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

func (r *GlooSoloIoUpstreamGroupV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_upstream_group_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GlooSoloIoUpstreamGroupV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gloo_solo_io_upstream_group_v1")

	var state GlooSoloIoUpstreamGroupV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoUpstreamGroupV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("UpstreamGroup")

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

func (r *GlooSoloIoUpstreamGroupV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gloo_solo_io_upstream_group_v1")
	// NO-OP: Terraform removes the state automatically for us
}
