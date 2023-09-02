/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gloo_solo_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &GlooSoloIoProxyV1DataSource{}
	_ datasource.DataSourceWithConfigure = &GlooSoloIoProxyV1DataSource{}
)

func NewGlooSoloIoProxyV1DataSource() datasource.DataSource {
	return &GlooSoloIoProxyV1DataSource{}
}

type GlooSoloIoProxyV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GlooSoloIoProxyV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CompressedSpec *string `tfsdk:"compressed_spec" json:"compressedSpec,omitempty"`
		Listeners      *[]struct {
			AggregateListener *map[string]string `tfsdk:"aggregate_listener" json:"aggregateListener,omitempty"`
			BindAddress       *string            `tfsdk:"bind_address" json:"bindAddress,omitempty"`
			BindPort          *int64             `tfsdk:"bind_port" json:"bindPort,omitempty"`
			HttpListener      *map[string]string `tfsdk:"http_listener" json:"httpListener,omitempty"`
			HybridListener    *map[string]string `tfsdk:"hybrid_listener" json:"hybridListener,omitempty"`
			Metadata          *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			MetadataStatic    *struct {
				Sources *[]struct {
					ObservedGeneration *int64  `tfsdk:"observed_generation" json:"observedGeneration,omitempty"`
					ResourceKind       *string `tfsdk:"resource_kind" json:"resourceKind,omitempty"`
					ResourceRef        *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"resource_ref" json:"resourceRef,omitempty"`
				} `tfsdk:"sources" json:"sources,omitempty"`
			} `tfsdk:"metadata_static" json:"metadataStatic,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Options *struct {
				AccessLoggingService *struct {
					AccessLog *[]struct {
						FileSink *struct {
							JsonFormat   *map[string]string `tfsdk:"json_format" json:"jsonFormat,omitempty"`
							Path         *string            `tfsdk:"path" json:"path,omitempty"`
							StringFormat *string            `tfsdk:"string_format" json:"stringFormat,omitempty"`
						} `tfsdk:"file_sink" json:"fileSink,omitempty"`
						Filter *struct {
							AndFilter      *map[string]string `tfsdk:"and_filter" json:"andFilter,omitempty"`
							DurationFilter *struct {
								Comparison *struct {
									Op    *string `tfsdk:"op" json:"op,omitempty"`
									Value *struct {
										DefaultValue *int64  `tfsdk:"default_value" json:"defaultValue,omitempty"`
										RuntimeKey   *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"comparison" json:"comparison,omitempty"`
							} `tfsdk:"duration_filter" json:"durationFilter,omitempty"`
							GrpcStatusFilter *struct {
								Exclude  *bool     `tfsdk:"exclude" json:"exclude,omitempty"`
								Statuses *[]string `tfsdk:"statuses" json:"statuses,omitempty"`
							} `tfsdk:"grpc_status_filter" json:"grpcStatusFilter,omitempty"`
							HeaderFilter *struct {
								Header *struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									SafeRegexMatch *struct {
										GoogleRe2 *struct {
											MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
										} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
										Regex *string `tfsdk:"regex" json:"regex,omitempty"`
									} `tfsdk:"safe_regex_match" json:"safeRegexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"header" json:"header,omitempty"`
							} `tfsdk:"header_filter" json:"headerFilter,omitempty"`
							NotHealthCheckFilter *map[string]string `tfsdk:"not_health_check_filter" json:"notHealthCheckFilter,omitempty"`
							OrFilter             *map[string]string `tfsdk:"or_filter" json:"orFilter,omitempty"`
							ResponseFlagFilter   *struct {
								Flags *[]string `tfsdk:"flags" json:"flags,omitempty"`
							} `tfsdk:"response_flag_filter" json:"responseFlagFilter,omitempty"`
							RuntimeFilter *struct {
								PercentSampled *struct {
									Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
									Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
								} `tfsdk:"percent_sampled" json:"percentSampled,omitempty"`
								RuntimeKey               *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
								UseIndependentRandomness *bool   `tfsdk:"use_independent_randomness" json:"useIndependentRandomness,omitempty"`
							} `tfsdk:"runtime_filter" json:"runtimeFilter,omitempty"`
							StatusCodeFilter *struct {
								Comparison *struct {
									Op    *string `tfsdk:"op" json:"op,omitempty"`
									Value *struct {
										DefaultValue *int64  `tfsdk:"default_value" json:"defaultValue,omitempty"`
										RuntimeKey   *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"comparison" json:"comparison,omitempty"`
							} `tfsdk:"status_code_filter" json:"statusCodeFilter,omitempty"`
							TraceableFilter *map[string]string `tfsdk:"traceable_filter" json:"traceableFilter,omitempty"`
						} `tfsdk:"filter" json:"filter,omitempty"`
						GrpcService *struct {
							AdditionalRequestHeadersToLog   *[]string `tfsdk:"additional_request_headers_to_log" json:"additionalRequestHeadersToLog,omitempty"`
							AdditionalResponseHeadersToLog  *[]string `tfsdk:"additional_response_headers_to_log" json:"additionalResponseHeadersToLog,omitempty"`
							AdditionalResponseTrailersToLog *[]string `tfsdk:"additional_response_trailers_to_log" json:"additionalResponseTrailersToLog,omitempty"`
							LogName                         *string   `tfsdk:"log_name" json:"logName,omitempty"`
							StaticClusterName               *string   `tfsdk:"static_cluster_name" json:"staticClusterName,omitempty"`
						} `tfsdk:"grpc_service" json:"grpcService,omitempty"`
					} `tfsdk:"access_log" json:"accessLog,omitempty"`
				} `tfsdk:"access_logging_service" json:"accessLoggingService,omitempty"`
				ConnectionBalanceConfig *struct {
					ExactBalance *map[string]string `tfsdk:"exact_balance" json:"exactBalance,omitempty"`
				} `tfsdk:"connection_balance_config" json:"connectionBalanceConfig,omitempty"`
				Extensions *struct {
					Configs *map[string]string `tfsdk:"configs" json:"configs,omitempty"`
				} `tfsdk:"extensions" json:"extensions,omitempty"`
				PerConnectionBufferLimitBytes *int64 `tfsdk:"per_connection_buffer_limit_bytes" json:"perConnectionBufferLimitBytes,omitempty"`
				ProxyProtocol                 *struct {
					AllowRequestsWithoutProxyProtocol *bool `tfsdk:"allow_requests_without_proxy_protocol" json:"allowRequestsWithoutProxyProtocol,omitempty"`
					Rules                             *[]struct {
						OnTlvPresent *struct {
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
						} `tfsdk:"on_tlv_present" json:"onTlvPresent,omitempty"`
						TlvType *int64 `tfsdk:"tlv_type" json:"tlvType,omitempty"`
					} `tfsdk:"rules" json:"rules,omitempty"`
				} `tfsdk:"proxy_protocol" json:"proxyProtocol,omitempty"`
				SocketOptions *[]struct {
					BufValue    *string `tfsdk:"buf_value" json:"bufValue,omitempty"`
					Description *string `tfsdk:"description" json:"description,omitempty"`
					IntValue    *int64  `tfsdk:"int_value" json:"intValue,omitempty"`
					Level       *int64  `tfsdk:"level" json:"level,omitempty"`
					Name        *int64  `tfsdk:"name" json:"name,omitempty"`
					State       *string `tfsdk:"state" json:"state,omitempty"`
				} `tfsdk:"socket_options" json:"socketOptions,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			RouteOptions *struct {
				MaxDirectResponseBodySizeBytes *int64 `tfsdk:"max_direct_response_body_size_bytes" json:"maxDirectResponseBodySizeBytes,omitempty"`
			} `tfsdk:"route_options" json:"routeOptions,omitempty"`
			SslConfigurations *[]struct {
				AlpnProtocols               *[]string `tfsdk:"alpn_protocols" json:"alpnProtocols,omitempty"`
				DisableTlsSessionResumption *bool     `tfsdk:"disable_tls_session_resumption" json:"disableTlsSessionResumption,omitempty"`
				OcspStaplePolicy            *string   `tfsdk:"ocsp_staple_policy" json:"ocspStaplePolicy,omitempty"`
				OneWayTls                   *bool     `tfsdk:"one_way_tls" json:"oneWayTls,omitempty"`
				Parameters                  *struct {
					CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
					EcdhCurves             *[]string `tfsdk:"ecdh_curves" json:"ecdhCurves,omitempty"`
					MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
					MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
				Sds *struct {
					CallCredentials *struct {
						FileCredentialSource *struct {
							Header        *string `tfsdk:"header" json:"header,omitempty"`
							TokenFileName *string `tfsdk:"token_file_name" json:"tokenFileName,omitempty"`
						} `tfsdk:"file_credential_source" json:"fileCredentialSource,omitempty"`
					} `tfsdk:"call_credentials" json:"callCredentials,omitempty"`
					CertificatesSecretName *string `tfsdk:"certificates_secret_name" json:"certificatesSecretName,omitempty"`
					ClusterName            *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
					TargetUri              *string `tfsdk:"target_uri" json:"targetUri,omitempty"`
					ValidationContextName  *string `tfsdk:"validation_context_name" json:"validationContextName,omitempty"`
				} `tfsdk:"sds" json:"sds,omitempty"`
				SecretRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				SniDomains *[]string `tfsdk:"sni_domains" json:"sniDomains,omitempty"`
				SslFiles   *struct {
					OcspStaple *string `tfsdk:"ocsp_staple" json:"ocspStaple,omitempty"`
					RootCa     *string `tfsdk:"root_ca" json:"rootCa,omitempty"`
					TlsCert    *string `tfsdk:"tls_cert" json:"tlsCert,omitempty"`
					TlsKey     *string `tfsdk:"tls_key" json:"tlsKey,omitempty"`
				} `tfsdk:"ssl_files" json:"sslFiles,omitempty"`
				TransportSocketConnectTimeout *string   `tfsdk:"transport_socket_connect_timeout" json:"transportSocketConnectTimeout,omitempty"`
				VerifySubjectAltName          *[]string `tfsdk:"verify_subject_alt_name" json:"verifySubjectAltName,omitempty"`
			} `tfsdk:"ssl_configurations" json:"sslConfigurations,omitempty"`
			TcpListener   *map[string]string `tfsdk:"tcp_listener" json:"tcpListener,omitempty"`
			UseProxyProto *bool              `tfsdk:"use_proxy_proto" json:"useProxyProto,omitempty"`
		} `tfsdk:"listeners" json:"listeners,omitempty"`
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GlooSoloIoProxyV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gloo_solo_io_proxy_v1"
}

func (r *GlooSoloIoProxyV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"compressed_spec": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"listeners": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aggregate_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"bind_address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"bind_port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"http_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"hybrid_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"metadata": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"metadata_static": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"sources": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"observed_generation": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource_kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"options": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"access_logging_service": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_log": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"file_sink": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"json_format": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"path": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"string_format": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"filter": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"and_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"duration_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"comparison": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"op": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"default_value": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"runtime_key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"grpc_status_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"exclude": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"statuses": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"header": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"exact_match": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"invert_match": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"prefix_match": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"present_match": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"range_match": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"end": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"start": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"safe_regex_match": schema.SingleNestedAttribute{
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
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"regex": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"suffix_match": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"not_health_check_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"or_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"response_flag_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"flags": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"runtime_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"percent_sampled": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"denominator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"numerator": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"runtime_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"use_independent_randomness": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"status_code_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"comparison": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"op": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"default_value": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"runtime_key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"traceable_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"grpc_service": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"additional_request_headers_to_log": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"additional_response_headers_to_log": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"additional_response_trailers_to_log": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"log_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"static_cluster_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"connection_balance_config": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact_balance": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
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
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"per_connection_buffer_limit_bytes": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"proxy_protocol": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_requests_without_proxy_protocol": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"rules": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"on_tlv_present": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"metadata_namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"tlv_type": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"socket_options": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"buf_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"int_value": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"level": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"state": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"route_options": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max_direct_response_body_size_bytes": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"ssl_configurations": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alpn_protocols": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"disable_tls_session_resumption": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ocsp_staple_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"one_way_tls": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"parameters": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"cipher_suites": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"ecdh_curves": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"maximum_protocol_version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"minimum_protocol_version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"sds": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"call_credentials": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"file_credential_source": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"token_file_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"certificates_secret_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"cluster_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"target_uri": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"validation_context_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"sni_domains": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"ssl_files": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"ocsp_staple": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"root_ca": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"tls_cert": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"tls_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"transport_socket_connect_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"verify_subject_alt_name": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"tcp_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"use_proxy_proto": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *GlooSoloIoProxyV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *GlooSoloIoProxyV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_gloo_solo_io_proxy_v1")

	var data GlooSoloIoProxyV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "Proxy"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse GlooSoloIoProxyV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("gloo.solo.io/v1")
	data.Kind = pointer.String("Proxy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
