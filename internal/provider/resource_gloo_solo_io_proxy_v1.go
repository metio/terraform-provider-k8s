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

type GlooSoloIoProxyV1Resource struct{}

var (
	_ resource.Resource = (*GlooSoloIoProxyV1Resource)(nil)
)

type GlooSoloIoProxyV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GlooSoloIoProxyV1GoModel struct {
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
		CompressedSpec *string `tfsdk:"compressed_spec" yaml:"compressedSpec,omitempty"`

		Listeners *[]struct {
			AggregateListener utilities.Dynamic `tfsdk:"aggregate_listener" yaml:"aggregateListener,omitempty"`

			BindAddress *string `tfsdk:"bind_address" yaml:"bindAddress,omitempty"`

			BindPort *int64 `tfsdk:"bind_port" yaml:"bindPort,omitempty"`

			HttpListener utilities.Dynamic `tfsdk:"http_listener" yaml:"httpListener,omitempty"`

			HybridListener utilities.Dynamic `tfsdk:"hybrid_listener" yaml:"hybridListener,omitempty"`

			Metadata utilities.Dynamic `tfsdk:"metadata" yaml:"metadata,omitempty"`

			MetadataStatic *struct {
				Sources *[]struct {
					ObservedGeneration utilities.IntOrString `tfsdk:"observed_generation" yaml:"observedGeneration,omitempty"`

					ResourceKind *string `tfsdk:"resource_kind" yaml:"resourceKind,omitempty"`

					ResourceRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"resource_ref" yaml:"resourceRef,omitempty"`
				} `tfsdk:"sources" yaml:"sources,omitempty"`
			} `tfsdk:"metadata_static" yaml:"metadataStatic,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Options *struct {
				AccessLoggingService *struct {
					AccessLog *[]struct {
						FileSink *struct {
							JsonFormat utilities.Dynamic `tfsdk:"json_format" yaml:"jsonFormat,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							StringFormat *string `tfsdk:"string_format" yaml:"stringFormat,omitempty"`
						} `tfsdk:"file_sink" yaml:"fileSink,omitempty"`

						GrpcService *struct {
							AdditionalRequestHeadersToLog *[]string `tfsdk:"additional_request_headers_to_log" yaml:"additionalRequestHeadersToLog,omitempty"`

							AdditionalResponseHeadersToLog *[]string `tfsdk:"additional_response_headers_to_log" yaml:"additionalResponseHeadersToLog,omitempty"`

							AdditionalResponseTrailersToLog *[]string `tfsdk:"additional_response_trailers_to_log" yaml:"additionalResponseTrailersToLog,omitempty"`

							LogName *string `tfsdk:"log_name" yaml:"logName,omitempty"`

							StaticClusterName *string `tfsdk:"static_cluster_name" yaml:"staticClusterName,omitempty"`
						} `tfsdk:"grpc_service" yaml:"grpcService,omitempty"`
					} `tfsdk:"access_log" yaml:"accessLog,omitempty"`
				} `tfsdk:"access_logging_service" yaml:"accessLoggingService,omitempty"`

				Extensions *struct {
					Configs utilities.Dynamic `tfsdk:"configs" yaml:"configs,omitempty"`
				} `tfsdk:"extensions" yaml:"extensions,omitempty"`

				PerConnectionBufferLimitBytes *int64 `tfsdk:"per_connection_buffer_limit_bytes" yaml:"perConnectionBufferLimitBytes,omitempty"`

				ProxyProtocol *struct {
					AllowRequestsWithoutProxyProtocol *bool `tfsdk:"allow_requests_without_proxy_protocol" yaml:"allowRequestsWithoutProxyProtocol,omitempty"`

					Rules *[]struct {
						OnTlvPresent *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`
						} `tfsdk:"on_tlv_present" yaml:"onTlvPresent,omitempty"`

						TlvType *int64 `tfsdk:"tlv_type" yaml:"tlvType,omitempty"`
					} `tfsdk:"rules" yaml:"rules,omitempty"`
				} `tfsdk:"proxy_protocol" yaml:"proxyProtocol,omitempty"`

				SocketOptions *[]struct {
					BufValue *string `tfsdk:"buf_value" yaml:"bufValue,omitempty"`

					Description *string `tfsdk:"description" yaml:"description,omitempty"`

					IntValue utilities.IntOrString `tfsdk:"int_value" yaml:"intValue,omitempty"`

					Level utilities.IntOrString `tfsdk:"level" yaml:"level,omitempty"`

					Name utilities.IntOrString `tfsdk:"name" yaml:"name,omitempty"`

					State utilities.IntOrString `tfsdk:"state" yaml:"state,omitempty"`
				} `tfsdk:"socket_options" yaml:"socketOptions,omitempty"`
			} `tfsdk:"options" yaml:"options,omitempty"`

			RouteOptions *struct {
				MaxDirectResponseBodySizeBytes *int64 `tfsdk:"max_direct_response_body_size_bytes" yaml:"maxDirectResponseBodySizeBytes,omitempty"`
			} `tfsdk:"route_options" yaml:"routeOptions,omitempty"`

			SslConfigurations *[]struct {
				AlpnProtocols *[]string `tfsdk:"alpn_protocols" yaml:"alpnProtocols,omitempty"`

				DisableTlsSessionResumption *bool `tfsdk:"disable_tls_session_resumption" yaml:"disableTlsSessionResumption,omitempty"`

				OneWayTls *bool `tfsdk:"one_way_tls" yaml:"oneWayTls,omitempty"`

				Parameters *struct {
					CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

					EcdhCurves *[]string `tfsdk:"ecdh_curves" yaml:"ecdhCurves,omitempty"`

					MaximumProtocolVersion utilities.IntOrString `tfsdk:"maximum_protocol_version" yaml:"maximumProtocolVersion,omitempty"`

					MinimumProtocolVersion utilities.IntOrString `tfsdk:"minimum_protocol_version" yaml:"minimumProtocolVersion,omitempty"`
				} `tfsdk:"parameters" yaml:"parameters,omitempty"`

				Sds *struct {
					CallCredentials *struct {
						FileCredentialSource *struct {
							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							TokenFileName *string `tfsdk:"token_file_name" yaml:"tokenFileName,omitempty"`
						} `tfsdk:"file_credential_source" yaml:"fileCredentialSource,omitempty"`
					} `tfsdk:"call_credentials" yaml:"callCredentials,omitempty"`

					CertificatesSecretName *string `tfsdk:"certificates_secret_name" yaml:"certificatesSecretName,omitempty"`

					ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

					TargetUri *string `tfsdk:"target_uri" yaml:"targetUri,omitempty"`

					ValidationContextName *string `tfsdk:"validation_context_name" yaml:"validationContextName,omitempty"`
				} `tfsdk:"sds" yaml:"sds,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				SniDomains *[]string `tfsdk:"sni_domains" yaml:"sniDomains,omitempty"`

				SslFiles *struct {
					RootCa *string `tfsdk:"root_ca" yaml:"rootCa,omitempty"`

					TlsCert *string `tfsdk:"tls_cert" yaml:"tlsCert,omitempty"`

					TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`
				} `tfsdk:"ssl_files" yaml:"sslFiles,omitempty"`

				TransportSocketConnectTimeout *string `tfsdk:"transport_socket_connect_timeout" yaml:"transportSocketConnectTimeout,omitempty"`

				VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" yaml:"verifySubjectAltName,omitempty"`
			} `tfsdk:"ssl_configurations" yaml:"sslConfigurations,omitempty"`

			TcpListener utilities.Dynamic `tfsdk:"tcp_listener" yaml:"tcpListener,omitempty"`

			UseProxyProto *bool `tfsdk:"use_proxy_proto" yaml:"useProxyProto,omitempty"`
		} `tfsdk:"listeners" yaml:"listeners,omitempty"`

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGlooSoloIoProxyV1Resource() resource.Resource {
	return &GlooSoloIoProxyV1Resource{}
}

func (r *GlooSoloIoProxyV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gloo_solo_io_proxy_v1"
}

func (r *GlooSoloIoProxyV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

					"compressed_spec": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"listeners": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"aggregate_listener": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bind_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bind_port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_listener": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hybrid_listener": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metadata": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metadata_static": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"sources": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"observed_generation": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_kind": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_ref": {
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

									"access_logging_service": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_log": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"file_sink": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"json_format": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.DynamicType{},

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

															"string_format": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc_service": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"additional_request_headers_to_log": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"additional_response_headers_to_log": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"additional_response_trailers_to_log": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"log_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"static_cluster_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

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

									"per_connection_buffer_limit_bytes": {
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

									"proxy_protocol": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_requests_without_proxy_protocol": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rules": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"on_tlv_present": {
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

															"metadata_namespace": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tlv_type": {
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

									"socket_options": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"buf_value": {
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

											"description": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"int_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"level": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"state": {
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

							"route_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_direct_response_body_size_bytes": {
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

							"ssl_configurations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"alpn_protocols": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_tls_session_resumption": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"one_way_tls": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"parameters": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cipher_suites": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ecdh_curves": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maximum_protocol_version": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum_protocol_version": {
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

									"sds": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"call_credentials": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"file_credential_source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"token_file_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

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

											"certificates_secret_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_uri": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"validation_context_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
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

									"sni_domains": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ssl_files": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"root_ca": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_cert": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"transport_socket_connect_timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify_subject_alt_name": {
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

							"tcp_listener": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_proxy_proto": {
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

func (r *GlooSoloIoProxyV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gloo_solo_io_proxy_v1")

	var state GlooSoloIoProxyV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoProxyV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("Proxy")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GlooSoloIoProxyV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_proxy_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GlooSoloIoProxyV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gloo_solo_io_proxy_v1")

	var state GlooSoloIoProxyV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoProxyV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("Proxy")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GlooSoloIoProxyV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gloo_solo_io_proxy_v1")
	// NO-OP: Terraform removes the state automatically for us
}
