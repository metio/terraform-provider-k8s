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

type GraphqlGlooSoloIoGraphQLApiV1Beta1Resource struct{}

var (
	_ resource.Resource = (*GraphqlGlooSoloIoGraphQLApiV1Beta1Resource)(nil)
)

type GraphqlGlooSoloIoGraphQLApiV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GraphqlGlooSoloIoGraphQLApiV1Beta1GoModel struct {
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
		AllowedQueryHashes *[]string `tfsdk:"allowed_query_hashes" yaml:"allowedQueryHashes,omitempty"`

		ExecutableSchema *struct {
			Executor *struct {
				Local *struct {
					EnableIntrospection *bool `tfsdk:"enable_introspection" yaml:"enableIntrospection,omitempty"`

					Options *struct {
						MaxDepth *int64 `tfsdk:"max_depth" yaml:"maxDepth,omitempty"`
					} `tfsdk:"options" yaml:"options,omitempty"`

					Resolutions *struct {
						GrpcResolver *struct {
							RequestTransform *struct {
								MethodName *string `tfsdk:"method_name" yaml:"methodName,omitempty"`

								OutgoingMessageJson utilities.Dynamic `tfsdk:"outgoing_message_json" yaml:"outgoingMessageJson,omitempty"`

								RequestMetadata *map[string]string `tfsdk:"request_metadata" yaml:"requestMetadata,omitempty"`

								ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
							} `tfsdk:"request_transform" yaml:"requestTransform,omitempty"`

							SpanName *string `tfsdk:"span_name" yaml:"spanName,omitempty"`

							UpstreamRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`
						} `tfsdk:"grpc_resolver" yaml:"grpcResolver,omitempty"`

						MockResolver *struct {
							AsyncResponse *struct {
								Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

								Response utilities.Dynamic `tfsdk:"response" yaml:"response,omitempty"`
							} `tfsdk:"async_response" yaml:"asyncResponse,omitempty"`

							ErrorResponse *string `tfsdk:"error_response" yaml:"errorResponse,omitempty"`

							SyncResponse utilities.Dynamic `tfsdk:"sync_response" yaml:"syncResponse,omitempty"`
						} `tfsdk:"mock_resolver" yaml:"mockResolver,omitempty"`

						RestResolver *struct {
							Request *struct {
								Body utilities.Dynamic `tfsdk:"body" yaml:"body,omitempty"`

								Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

								QueryParams *map[string]string `tfsdk:"query_params" yaml:"queryParams,omitempty"`
							} `tfsdk:"request" yaml:"request,omitempty"`

							Response *struct {
								ResultRoot *string `tfsdk:"result_root" yaml:"resultRoot,omitempty"`

								Setters *map[string]string `tfsdk:"setters" yaml:"setters,omitempty"`
							} `tfsdk:"response" yaml:"response,omitempty"`

							SpanName *string `tfsdk:"span_name" yaml:"spanName,omitempty"`

							UpstreamRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`
						} `tfsdk:"rest_resolver" yaml:"restResolver,omitempty"`

						StatPrefix *string `tfsdk:"stat_prefix" yaml:"statPrefix,omitempty"`
					} `tfsdk:"resolutions" yaml:"resolutions,omitempty"`
				} `tfsdk:"local" yaml:"local,omitempty"`

				Remote *struct {
					Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

					QueryParams *map[string]string `tfsdk:"query_params" yaml:"queryParams,omitempty"`

					SpanName *string `tfsdk:"span_name" yaml:"spanName,omitempty"`

					UpstreamRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`
				} `tfsdk:"remote" yaml:"remote,omitempty"`
			} `tfsdk:"executor" yaml:"executor,omitempty"`

			GrpcDescriptorRegistry *struct {
				ProtoDescriptor *string `tfsdk:"proto_descriptor" yaml:"protoDescriptor,omitempty"`

				ProtoDescriptorBin *string `tfsdk:"proto_descriptor_bin" yaml:"protoDescriptorBin,omitempty"`

				ProtoRefsList *struct {
					ConfigMapRefs *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"config_map_refs" yaml:"configMapRefs,omitempty"`
				} `tfsdk:"proto_refs_list" yaml:"protoRefsList,omitempty"`
			} `tfsdk:"grpc_descriptor_registry" yaml:"grpcDescriptorRegistry,omitempty"`

			SchemaDefinition *string `tfsdk:"schema_definition" yaml:"schemaDefinition,omitempty"`
		} `tfsdk:"executable_schema" yaml:"executableSchema,omitempty"`

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`

		Options *struct {
			LogSensitiveInfo *bool `tfsdk:"log_sensitive_info" yaml:"logSensitiveInfo,omitempty"`
		} `tfsdk:"options" yaml:"options,omitempty"`

		PersistedQueryCacheConfig *struct {
			CacheSize *int64 `tfsdk:"cache_size" yaml:"cacheSize,omitempty"`
		} `tfsdk:"persisted_query_cache_config" yaml:"persistedQueryCacheConfig,omitempty"`

		StatPrefix *string `tfsdk:"stat_prefix" yaml:"statPrefix,omitempty"`

		StitchedSchema *struct {
			Subschemas *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				TypeMerge *struct {
					Args *map[string]string `tfsdk:"args" yaml:"args,omitempty"`

					QueryName *string `tfsdk:"query_name" yaml:"queryName,omitempty"`

					SelectionSet *string `tfsdk:"selection_set" yaml:"selectionSet,omitempty"`
				} `tfsdk:"type_merge" yaml:"typeMerge,omitempty"`
			} `tfsdk:"subschemas" yaml:"subschemas,omitempty"`
		} `tfsdk:"stitched_schema" yaml:"stitchedSchema,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGraphqlGlooSoloIoGraphQLApiV1Beta1Resource() resource.Resource {
	return &GraphqlGlooSoloIoGraphQLApiV1Beta1Resource{}
}

func (r *GraphqlGlooSoloIoGraphQLApiV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_graphql_gloo_solo_io_graph_ql_api_v1beta1"
}

func (r *GraphqlGlooSoloIoGraphQLApiV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

					"allowed_query_hashes": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"executable_schema": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"executor": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"local": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable_introspection": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"max_depth": {
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

											"resolutions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"grpc_resolver": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"request_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"method_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"outgoing_message_json": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

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
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"span_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"upstream_ref": {
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

													"mock_resolver": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"async_response": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"delay": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"response": {
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

															"error_response": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sync_response": {
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

													"rest_resolver": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"request": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"query_params": {
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

															"response": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"result_root": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"setters": {
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

															"span_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"upstream_ref": {
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

													"stat_prefix": {
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

									"remote": {
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

											"query_params": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"span_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"upstream_ref": {
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

							"grpc_descriptor_registry": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"proto_descriptor": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"proto_descriptor_bin": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.Base64Validator(),
										},
									},

									"proto_refs_list": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_refs": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schema_definition": {
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

					"options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"log_sensitive_info": {
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

					"persisted_query_cache_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cache_size": {
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

					"stat_prefix": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"stitched_schema": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"subschemas": {
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

									"type_merge": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"args": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selection_set": {
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
		},
	}, nil
}

func (r *GraphqlGlooSoloIoGraphQLApiV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1")

	var state GraphqlGlooSoloIoGraphQLApiV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GraphqlGlooSoloIoGraphQLApiV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("graphql.gloo.solo.io/v1beta1")
	goModel.Kind = utilities.Ptr("GraphQLApi")

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

func (r *GraphqlGlooSoloIoGraphQLApiV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *GraphqlGlooSoloIoGraphQLApiV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1")

	var state GraphqlGlooSoloIoGraphQLApiV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GraphqlGlooSoloIoGraphQLApiV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("graphql.gloo.solo.io/v1beta1")
	goModel.Kind = utilities.Ptr("GraphQLApi")

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

func (r *GraphqlGlooSoloIoGraphQLApiV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
