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
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GlooSoloIoProxyV1Manifest{}
)

func NewGlooSoloIoProxyV1Manifest() datasource.DataSource {
	return &GlooSoloIoProxyV1Manifest{}
}

type GlooSoloIoProxyV1Manifest struct{}

type GlooSoloIoProxyV1ManifestData struct {
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

func (r *GlooSoloIoProxyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gloo_solo_io_proxy_v1_manifest"
}

func (r *GlooSoloIoProxyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
					"compressed_spec": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
									Optional:            true,
									Computed:            false,
								},

								"bind_address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"bind_port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"http_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"hybrid_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metadata": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
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
														Optional:            true,
														Computed:            false,
													},

													"resource_kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_ref": schema.SingleNestedAttribute{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
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

																	"string_format": schema.StringAttribute{
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

															"filter": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"and_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"default_value": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
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
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"grpc_status_filter": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"exclude": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"statuses": schema.ListAttribute{
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

																					"suffix_match": schema.StringAttribute{
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

																	"not_health_check_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"or_filter": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
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

																			"use_independent_randomness": schema.BoolAttribute{
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
																						Optional:            true,
																						Computed:            false,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"default_value": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
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
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"traceable_filter": schema.MapAttribute{
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

															"grpc_service": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"additional_request_headers_to_log": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"additional_response_headers_to_log": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"additional_response_trailers_to_log": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"log_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"static_cluster_name": schema.StringAttribute{
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

										"connection_balance_config": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact_balance": schema.MapAttribute{
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

										"per_connection_buffer_limit_bytes": schema.Int64Attribute{
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

										"proxy_protocol": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"allow_requests_without_proxy_protocol": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"tlv_type": schema.Int64Attribute{
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

										"socket_options": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"buf_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.Base64Validator(),
														},
													},

													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"int_value": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"level": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"state": schema.StringAttribute{
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

								"route_options": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max_direct_response_body_size_bytes": schema.Int64Attribute{
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
												Optional:            true,
												Computed:            false,
											},

											"disable_tls_session_resumption": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ocsp_staple_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"one_way_tls": schema.BoolAttribute{
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
													"cipher_suites": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ecdh_curves": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum_protocol_version": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"minimum_protocol_version": schema.StringAttribute{
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
																		Optional:            true,
																		Computed:            false,
																	},

																	"token_file_name": schema.StringAttribute{
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

													"certificates_secret_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_uri": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"validation_context_name": schema.StringAttribute{
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

											"secret_ref": schema.SingleNestedAttribute{
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

											"sni_domains": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ssl_files": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"ocsp_staple": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"root_ca": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_cert": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_key": schema.StringAttribute{
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

											"transport_socket_connect_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"verify_subject_alt_name": schema.ListAttribute{
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

								"tcp_listener": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"use_proxy_proto": schema.BoolAttribute{
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

func (r *GlooSoloIoProxyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_proxy_v1_manifest")

	var model GlooSoloIoProxyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("gloo.solo.io/v1")
	model.Kind = pointer.String("Proxy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
