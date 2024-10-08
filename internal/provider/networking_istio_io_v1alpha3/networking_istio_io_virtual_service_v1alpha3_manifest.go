/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

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
	_ datasource.DataSource = &NetworkingIstioIoVirtualServiceV1Alpha3Manifest{}
)

func NewNetworkingIstioIoVirtualServiceV1Alpha3Manifest() datasource.DataSource {
	return &NetworkingIstioIoVirtualServiceV1Alpha3Manifest{}
}

type NetworkingIstioIoVirtualServiceV1Alpha3Manifest struct{}

type NetworkingIstioIoVirtualServiceV1Alpha3ManifestData struct {
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
		ExportTo *[]string `tfsdk:"export_to" json:"exportTo,omitempty"`
		Gateways *[]string `tfsdk:"gateways" json:"gateways,omitempty"`
		Hosts    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
		Http     *[]struct {
			CorsPolicy *struct {
				AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
				AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
				AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
				AllowOrigin      *[]string `tfsdk:"allow_origin" json:"allowOrigin,omitempty"`
				AllowOrigins     *[]struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"allow_origins" json:"allowOrigins,omitempty"`
				ExposeHeaders       *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
				MaxAge              *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
				UnmatchedPreflights *string   `tfsdk:"unmatched_preflights" json:"unmatchedPreflights,omitempty"`
			} `tfsdk:"cors_policy" json:"corsPolicy,omitempty"`
			Delegate *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"delegate" json:"delegate,omitempty"`
			DirectResponse *struct {
				Body *struct {
					Bytes  *string `tfsdk:"bytes" json:"bytes,omitempty"`
					String *string `tfsdk:"string" json:"string,omitempty"`
				} `tfsdk:"body" json:"body,omitempty"`
				Status *int64 `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"direct_response" json:"directResponse,omitempty"`
			Fault *struct {
				Abort *struct {
					GrpcStatus *string `tfsdk:"grpc_status" json:"grpcStatus,omitempty"`
					Http2Error *string `tfsdk:"http2_error" json:"http2Error,omitempty"`
					HttpStatus *int64  `tfsdk:"http_status" json:"httpStatus,omitempty"`
					Percentage *struct {
						Value *float64 `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"percentage" json:"percentage,omitempty"`
				} `tfsdk:"abort" json:"abort,omitempty"`
				Delay *struct {
					ExponentialDelay *string `tfsdk:"exponential_delay" json:"exponentialDelay,omitempty"`
					FixedDelay       *string `tfsdk:"fixed_delay" json:"fixedDelay,omitempty"`
					Percent          *int64  `tfsdk:"percent" json:"percent,omitempty"`
					Percentage       *struct {
						Value *float64 `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"percentage" json:"percentage,omitempty"`
				} `tfsdk:"delay" json:"delay,omitempty"`
			} `tfsdk:"fault" json:"fault,omitempty"`
			Headers *struct {
				Request *struct {
					Add    *map[string]string `tfsdk:"add" json:"add,omitempty"`
					Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
					Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"request" json:"request,omitempty"`
				Response *struct {
					Add    *map[string]string `tfsdk:"add" json:"add,omitempty"`
					Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
					Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"response" json:"response,omitempty"`
			} `tfsdk:"headers" json:"headers,omitempty"`
			Match *[]struct {
				Authority *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"authority" json:"authority,omitempty"`
				Gateways *[]string `tfsdk:"gateways" json:"gateways,omitempty"`
				Headers  *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				IgnoreUriCase *bool `tfsdk:"ignore_uri_case" json:"ignoreUriCase,omitempty"`
				Method        *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"method" json:"method,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				QueryParams *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"query_params" json:"queryParams,omitempty"`
				Scheme *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"scheme" json:"scheme,omitempty"`
				SourceLabels    *map[string]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				SourceNamespace *string            `tfsdk:"source_namespace" json:"sourceNamespace,omitempty"`
				StatPrefix      *string            `tfsdk:"stat_prefix" json:"statPrefix,omitempty"`
				Uri             *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				WithoutHeaders *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"without_headers" json:"withoutHeaders,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Mirror *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *struct {
					Number *int64 `tfsdk:"number" json:"number,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Subset *string `tfsdk:"subset" json:"subset,omitempty"`
			} `tfsdk:"mirror" json:"mirror,omitempty"`
			MirrorPercent    *int64 `tfsdk:"mirror_percent" json:"mirrorPercent,omitempty"`
			MirrorPercentage *struct {
				Value *float64 `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"mirror_percentage" json:"mirrorPercentage,omitempty"`
			Mirrors *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Port *struct {
						Number *int64 `tfsdk:"number" json:"number,omitempty"`
					} `tfsdk:"port" json:"port,omitempty"`
					Subset *string `tfsdk:"subset" json:"subset,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				Percentage *struct {
					Value *float64 `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"percentage" json:"percentage,omitempty"`
			} `tfsdk:"mirrors" json:"mirrors,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Redirect *struct {
				Authority    *string `tfsdk:"authority" json:"authority,omitempty"`
				DerivePort   *string `tfsdk:"derive_port" json:"derivePort,omitempty"`
				Port         *int64  `tfsdk:"port" json:"port,omitempty"`
				RedirectCode *int64  `tfsdk:"redirect_code" json:"redirectCode,omitempty"`
				Scheme       *string `tfsdk:"scheme" json:"scheme,omitempty"`
				Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"redirect" json:"redirect,omitempty"`
			Retries *struct {
				Attempts              *int64  `tfsdk:"attempts" json:"attempts,omitempty"`
				PerTryTimeout         *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
				RetryOn               *string `tfsdk:"retry_on" json:"retryOn,omitempty"`
				RetryRemoteLocalities *bool   `tfsdk:"retry_remote_localities" json:"retryRemoteLocalities,omitempty"`
			} `tfsdk:"retries" json:"retries,omitempty"`
			Rewrite *struct {
				Authority       *string `tfsdk:"authority" json:"authority,omitempty"`
				Uri             *string `tfsdk:"uri" json:"uri,omitempty"`
				UriRegexRewrite *struct {
					Match   *string `tfsdk:"match" json:"match,omitempty"`
					Rewrite *string `tfsdk:"rewrite" json:"rewrite,omitempty"`
				} `tfsdk:"uri_regex_rewrite" json:"uriRegexRewrite,omitempty"`
			} `tfsdk:"rewrite" json:"rewrite,omitempty"`
			Route *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Port *struct {
						Number *int64 `tfsdk:"number" json:"number,omitempty"`
					} `tfsdk:"port" json:"port,omitempty"`
					Subset *string `tfsdk:"subset" json:"subset,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				Headers *struct {
					Request *struct {
						Add    *map[string]string `tfsdk:"add" json:"add,omitempty"`
						Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
						Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
					} `tfsdk:"request" json:"request,omitempty"`
					Response *struct {
						Add    *map[string]string `tfsdk:"add" json:"add,omitempty"`
						Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
						Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
					} `tfsdk:"response" json:"response,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
		Tcp *[]struct {
			Match *[]struct {
				DestinationSubnets *[]string          `tfsdk:"destination_subnets" json:"destinationSubnets,omitempty"`
				Gateways           *[]string          `tfsdk:"gateways" json:"gateways,omitempty"`
				Port               *int64             `tfsdk:"port" json:"port,omitempty"`
				SourceLabels       *map[string]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				SourceNamespace    *string            `tfsdk:"source_namespace" json:"sourceNamespace,omitempty"`
				SourceSubnet       *string            `tfsdk:"source_subnet" json:"sourceSubnet,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Route *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Port *struct {
						Number *int64 `tfsdk:"number" json:"number,omitempty"`
					} `tfsdk:"port" json:"port,omitempty"`
					Subset *string `tfsdk:"subset" json:"subset,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
		} `tfsdk:"tcp" json:"tcp,omitempty"`
		Tls *[]struct {
			Match *[]struct {
				DestinationSubnets *[]string          `tfsdk:"destination_subnets" json:"destinationSubnets,omitempty"`
				Gateways           *[]string          `tfsdk:"gateways" json:"gateways,omitempty"`
				Port               *int64             `tfsdk:"port" json:"port,omitempty"`
				SniHosts           *[]string          `tfsdk:"sni_hosts" json:"sniHosts,omitempty"`
				SourceLabels       *map[string]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				SourceNamespace    *string            `tfsdk:"source_namespace" json:"sourceNamespace,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Route *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Port *struct {
						Number *int64 `tfsdk:"number" json:"number,omitempty"`
					} `tfsdk:"port" json:"port,omitempty"`
					Subset *string `tfsdk:"subset" json:"subset,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_virtual_service_v1alpha3_manifest"
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html",
				MarkdownDescription: "Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html",
				Attributes: map[string]schema.Attribute{
					"export_to": schema.ListAttribute{
						Description:         "A list of namespaces to which this virtual service is exported.",
						MarkdownDescription: "A list of namespaces to which this virtual service is exported.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gateways": schema.ListAttribute{
						Description:         "The names of gateways and sidecars that should apply these routes.",
						MarkdownDescription: "The names of gateways and sidecars that should apply these routes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hosts": schema.ListAttribute{
						Description:         "The destination hosts to which traffic is being sent.",
						MarkdownDescription: "The destination hosts to which traffic is being sent.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"http": schema.ListNestedAttribute{
						Description:         "An ordered list of route rules for HTTP traffic.",
						MarkdownDescription: "An ordered list of route rules for HTTP traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cors_policy": schema.SingleNestedAttribute{
									Description:         "Cross-Origin Resource Sharing policy (CORS).",
									MarkdownDescription: "Cross-Origin Resource Sharing policy (CORS).",
									Attributes: map[string]schema.Attribute{
										"allow_credentials": schema.BoolAttribute{
											Description:         "Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials.",
											MarkdownDescription: "Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allow_headers": schema.ListAttribute{
											Description:         "List of HTTP headers that can be used when requesting the resource.",
											MarkdownDescription: "List of HTTP headers that can be used when requesting the resource.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allow_methods": schema.ListAttribute{
											Description:         "List of HTTP methods allowed to access the resource.",
											MarkdownDescription: "List of HTTP methods allowed to access the resource.",
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

										"allow_origins": schema.ListNestedAttribute{
											Description:         "String patterns that match allowed origins.",
											MarkdownDescription: "String patterns that match allowed origins.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
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

										"expose_headers": schema.ListAttribute{
											Description:         "A list of HTTP headers that the browsers are allowed to access.",
											MarkdownDescription: "A list of HTTP headers that the browsers are allowed to access.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_age": schema.StringAttribute{
											Description:         "Specifies how long the results of a preflight request can be cached.",
											MarkdownDescription: "Specifies how long the results of a preflight request can be cached.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"unmatched_preflights": schema.StringAttribute{
											Description:         "Indicates whether preflight requests not matching the configured allowed origin shouldn't be forwarded to the upstream. Valid Options: FORWARD, IGNORE",
											MarkdownDescription: "Indicates whether preflight requests not matching the configured allowed origin shouldn't be forwarded to the upstream. Valid Options: FORWARD, IGNORE",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("UNSPECIFIED", "FORWARD", "IGNORE"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"delegate": schema.SingleNestedAttribute{
									Description:         "Delegate is used to specify the particular VirtualService which can be used to define delegate HTTPRoute.",
									MarkdownDescription: "Delegate is used to specify the particular VirtualService which can be used to define delegate HTTPRoute.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name specifies the name of the delegate VirtualService.",
											MarkdownDescription: "Name specifies the name of the delegate VirtualService.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace specifies the namespace where the delegate VirtualService resides.",
											MarkdownDescription: "Namespace specifies the namespace where the delegate VirtualService resides.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"direct_response": schema.SingleNestedAttribute{
									Description:         "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
									MarkdownDescription: "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
									Attributes: map[string]schema.Attribute{
										"body": schema.SingleNestedAttribute{
											Description:         "Specifies the content of the response body.",
											MarkdownDescription: "Specifies the content of the response body.",
											Attributes: map[string]schema.Attribute{
												"bytes": schema.StringAttribute{
													Description:         "response body as base64 encoded bytes.",
													MarkdownDescription: "response body as base64 encoded bytes.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"string": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"status": schema.Int64Attribute{
											Description:         "Specifies the HTTP response status to be returned.",
											MarkdownDescription: "Specifies the HTTP response status to be returned.",
											Required:            true,
											Optional:            false,
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

								"fault": schema.SingleNestedAttribute{
									Description:         "Fault injection policy to apply on HTTP traffic at the client side.",
									MarkdownDescription: "Fault injection policy to apply on HTTP traffic at the client side.",
									Attributes: map[string]schema.Attribute{
										"abort": schema.SingleNestedAttribute{
											Description:         "Abort Http request attempts and return error codes back to downstream service, giving the impression that the upstream service is faulty.",
											MarkdownDescription: "Abort Http request attempts and return error codes back to downstream service, giving the impression that the upstream service is faulty.",
											Attributes: map[string]schema.Attribute{
												"grpc_status": schema.StringAttribute{
													Description:         "GRPC status code to use to abort the request.",
													MarkdownDescription: "GRPC status code to use to abort the request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http2_error": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_status": schema.Int64Attribute{
													Description:         "HTTP status code to use to abort the Http request.",
													MarkdownDescription: "HTTP status code to use to abort the Http request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percentage": schema.SingleNestedAttribute{
													Description:         "Percentage of requests to be aborted with the error code provided.",
													MarkdownDescription: "Percentage of requests to be aborted with the error code provided.",
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

										"delay": schema.SingleNestedAttribute{
											Description:         "Delay requests before forwarding, emulating various failures such as network issues, overloaded upstream service, etc.",
											MarkdownDescription: "Delay requests before forwarding, emulating various failures such as network issues, overloaded upstream service, etc.",
											Attributes: map[string]schema.Attribute{
												"exponential_delay": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"fixed_delay": schema.StringAttribute{
													Description:         "Add a fixed delay before forwarding the request.",
													MarkdownDescription: "Add a fixed delay before forwarding the request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percent": schema.Int64Attribute{
													Description:         "Percentage of requests on which the delay will be injected (0-100).",
													MarkdownDescription: "Percentage of requests on which the delay will be injected (0-100).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"percentage": schema.SingleNestedAttribute{
													Description:         "Percentage of requests on which the delay will be injected.",
													MarkdownDescription: "Percentage of requests on which the delay will be injected.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"headers": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"request": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"add": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remove": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"set": schema.MapAttribute{
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

										"response": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"add": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remove": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"set": schema.MapAttribute{
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

								"match": schema.ListNestedAttribute{
									Description:         "Match conditions to be satisfied for the rule to be activated.",
									MarkdownDescription: "Match conditions to be satisfied for the rule to be activated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"authority": schema.SingleNestedAttribute{
												Description:         "HTTP Authority values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "HTTP Authority values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"gateways": schema.ListAttribute{
												Description:         "Names of gateways where the rule should be applied.",
												MarkdownDescription: "Names of gateways where the rule should be applied.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "The header keys must be lowercase and use hyphen as the separator, e.g.",
												MarkdownDescription: "The header keys must be lowercase and use hyphen as the separator, e.g.",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ignore_uri_case": schema.BoolAttribute{
												Description:         "Flag to specify whether the URI matching should be case-insensitive.",
												MarkdownDescription: "Flag to specify whether the URI matching should be case-insensitive.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"method": schema.SingleNestedAttribute{
												Description:         "HTTP Method values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "HTTP Method values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "The name assigned to a match.",
												MarkdownDescription: "The name assigned to a match.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Specifies the ports on the host that is being addressed.",
												MarkdownDescription: "Specifies the ports on the host that is being addressed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"query_params": schema.SingleNestedAttribute{
												Description:         "Query parameters for matching.",
												MarkdownDescription: "Query parameters for matching.",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheme": schema.SingleNestedAttribute{
												Description:         "URI Scheme values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "URI Scheme values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_labels": schema.MapAttribute{
												Description:         "One or more labels that constrain the applicability of a rule to source (client) workloads with the given labels.",
												MarkdownDescription: "One or more labels that constrain the applicability of a rule to source (client) workloads with the given labels.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_namespace": schema.StringAttribute{
												Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
												MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"stat_prefix": schema.StringAttribute{
												Description:         "The human readable prefix to use when emitting statistics for this route.",
												MarkdownDescription: "The human readable prefix to use when emitting statistics for this route.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uri": schema.SingleNestedAttribute{
												Description:         "URI to match values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "URI to match values are case-sensitive and formatted as follows: - 'exact: 'value'' for exact string match - 'prefix: 'value'' for prefix-based match - 'regex: 'value'' for [RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"without_headers": schema.SingleNestedAttribute{
												Description:         "withoutHeader has the same syntax with the header, but has opposite meaning.",
												MarkdownDescription: "withoutHeader has the same syntax with the header, but has opposite meaning.",
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
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

													"regex": schema.StringAttribute{
														Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
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

								"mirror": schema.SingleNestedAttribute{
									Description:         "Mirror HTTP traffic to a another destination in addition to forwarding the requests to the intended destination.",
									MarkdownDescription: "Mirror HTTP traffic to a another destination in addition to forwarding the requests to the intended destination.",
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "The name of a service from the service registry.",
											MarkdownDescription: "The name of a service from the service registry.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.SingleNestedAttribute{
											Description:         "Specifies the port on the host that is being addressed.",
											MarkdownDescription: "Specifies the port on the host that is being addressed.",
											Attributes: map[string]schema.Attribute{
												"number": schema.Int64Attribute{
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

										"subset": schema.StringAttribute{
											Description:         "The name of a subset within the service.",
											MarkdownDescription: "The name of a subset within the service.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"mirror_percent": schema.Int64Attribute{
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

								"mirror_percentage": schema.SingleNestedAttribute{
									Description:         "Percentage of the traffic to be mirrored by the 'mirror' field.",
									MarkdownDescription: "Percentage of the traffic to be mirrored by the 'mirror' field.",
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

								"mirrors": schema.ListNestedAttribute{
									Description:         "Specifies the destinations to mirror HTTP traffic in addition to the original destination.",
									MarkdownDescription: "Specifies the destinations to mirror HTTP traffic in addition to the original destination.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination": schema.SingleNestedAttribute{
												Description:         "Destination specifies the target of the mirror operation.",
												MarkdownDescription: "Destination specifies the target of the mirror operation.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "Specifies the port on the host that is being addressed.",
														MarkdownDescription: "Specifies the port on the host that is being addressed.",
														Attributes: map[string]schema.Attribute{
															"number": schema.Int64Attribute{
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

													"subset": schema.StringAttribute{
														Description:         "The name of a subset within the service.",
														MarkdownDescription: "The name of a subset within the service.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"percentage": schema.SingleNestedAttribute{
												Description:         "Percentage of the traffic to be mirrored by the 'destination' field.",
												MarkdownDescription: "Percentage of the traffic to be mirrored by the 'destination' field.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "The name assigned to the route for debugging purposes.",
									MarkdownDescription: "The name assigned to the route for debugging purposes.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"redirect": schema.SingleNestedAttribute{
									Description:         "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
									MarkdownDescription: "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
									Attributes: map[string]schema.Attribute{
										"authority": schema.StringAttribute{
											Description:         "On a redirect, overwrite the Authority/Host portion of the URL with this value.",
											MarkdownDescription: "On a redirect, overwrite the Authority/Host portion of the URL with this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"derive_port": schema.StringAttribute{
											Description:         "On a redirect, dynamically set the port: * FROM_PROTOCOL_DEFAULT: automatically set to 80 for HTTP and 443 for HTTPS. Valid Options: FROM_PROTOCOL_DEFAULT, FROM_REQUEST_PORT",
											MarkdownDescription: "On a redirect, dynamically set the port: * FROM_PROTOCOL_DEFAULT: automatically set to 80 for HTTP and 443 for HTTPS. Valid Options: FROM_PROTOCOL_DEFAULT, FROM_REQUEST_PORT",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("FROM_PROTOCOL_DEFAULT", "FROM_REQUEST_PORT"),
											},
										},

										"port": schema.Int64Attribute{
											Description:         "On a redirect, overwrite the port portion of the URL with this value.",
											MarkdownDescription: "On a redirect, overwrite the port portion of the URL with this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(4.294967295e+09),
											},
										},

										"redirect_code": schema.Int64Attribute{
											Description:         "On a redirect, Specifies the HTTP status code to use in the redirect response.",
											MarkdownDescription: "On a redirect, Specifies the HTTP status code to use in the redirect response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(4.294967295e+09),
											},
										},

										"scheme": schema.StringAttribute{
											Description:         "On a redirect, overwrite the scheme portion of the URL with this value.",
											MarkdownDescription: "On a redirect, overwrite the scheme portion of the URL with this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uri": schema.StringAttribute{
											Description:         "On a redirect, overwrite the Path portion of the URL with this value.",
											MarkdownDescription: "On a redirect, overwrite the Path portion of the URL with this value.",
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
									Description:         "Retry policy for HTTP requests.",
									MarkdownDescription: "Retry policy for HTTP requests.",
									Attributes: map[string]schema.Attribute{
										"attempts": schema.Int64Attribute{
											Description:         "Number of retries to be allowed for a given request.",
											MarkdownDescription: "Number of retries to be allowed for a given request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"per_try_timeout": schema.StringAttribute{
											Description:         "Timeout per attempt for a given request, including the initial call and any retries.",
											MarkdownDescription: "Timeout per attempt for a given request, including the initial call and any retries.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_on": schema.StringAttribute{
											Description:         "Specifies the conditions under which retry takes place.",
											MarkdownDescription: "Specifies the conditions under which retry takes place.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"retry_remote_localities": schema.BoolAttribute{
											Description:         "Flag to specify whether the retries should retry to other localities.",
											MarkdownDescription: "Flag to specify whether the retries should retry to other localities.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"rewrite": schema.SingleNestedAttribute{
									Description:         "Rewrite HTTP URIs and Authority headers.",
									MarkdownDescription: "Rewrite HTTP URIs and Authority headers.",
									Attributes: map[string]schema.Attribute{
										"authority": schema.StringAttribute{
											Description:         "rewrite the Authority/Host header with this value.",
											MarkdownDescription: "rewrite the Authority/Host header with this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uri": schema.StringAttribute{
											Description:         "rewrite the path (or the prefix) portion of the URI with this value.",
											MarkdownDescription: "rewrite the path (or the prefix) portion of the URI with this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uri_regex_rewrite": schema.SingleNestedAttribute{
											Description:         "rewrite the path portion of the URI with the specified regex.",
											MarkdownDescription: "rewrite the path portion of the URI with the specified regex.",
											Attributes: map[string]schema.Attribute{
												"match": schema.StringAttribute{
													Description:         "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "[RE2 style regex-based match](https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"rewrite": schema.StringAttribute{
													Description:         "The string that should replace into matching portions of original URI.",
													MarkdownDescription: "The string that should replace into matching portions of original URI.",
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

								"route": schema.ListNestedAttribute{
									Description:         "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
									MarkdownDescription: "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination": schema.SingleNestedAttribute{
												Description:         "Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to.",
												MarkdownDescription: "Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "Specifies the port on the host that is being addressed.",
														MarkdownDescription: "Specifies the port on the host that is being addressed.",
														Attributes: map[string]schema.Attribute{
															"number": schema.Int64Attribute{
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

													"subset": schema.StringAttribute{
														Description:         "The name of a subset within the service.",
														MarkdownDescription: "The name of a subset within the service.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"request": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"add": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"remove": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"set": schema.MapAttribute{
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

													"response": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"add": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"remove": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"set": schema.MapAttribute{
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

											"weight": schema.Int64Attribute{
												Description:         "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
												MarkdownDescription: "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
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

								"timeout": schema.StringAttribute{
									Description:         "Timeout for HTTP requests, default is disabled.",
									MarkdownDescription: "Timeout for HTTP requests, default is disabled.",
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

					"tcp": schema.ListNestedAttribute{
						Description:         "An ordered list of route rules for opaque TCP traffic.",
						MarkdownDescription: "An ordered list of route rules for opaque TCP traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match": schema.ListNestedAttribute{
									Description:         "Match conditions to be satisfied for the rule to be activated.",
									MarkdownDescription: "Match conditions to be satisfied for the rule to be activated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination_subnets": schema.ListAttribute{
												Description:         "IPv4 or IPv6 ip addresses of destination with optional subnet.",
												MarkdownDescription: "IPv4 or IPv6 ip addresses of destination with optional subnet.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gateways": schema.ListAttribute{
												Description:         "Names of gateways where the rule should be applied.",
												MarkdownDescription: "Names of gateways where the rule should be applied.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Specifies the port on the host that is being addressed.",
												MarkdownDescription: "Specifies the port on the host that is being addressed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"source_labels": schema.MapAttribute{
												Description:         "One or more labels that constrain the applicability of a rule to workloads with the given labels.",
												MarkdownDescription: "One or more labels that constrain the applicability of a rule to workloads with the given labels.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_namespace": schema.StringAttribute{
												Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
												MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_subnet": schema.StringAttribute{
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

								"route": schema.ListNestedAttribute{
									Description:         "The destination to which the connection should be forwarded to.",
									MarkdownDescription: "The destination to which the connection should be forwarded to.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination": schema.SingleNestedAttribute{
												Description:         "Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to.",
												MarkdownDescription: "Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "Specifies the port on the host that is being addressed.",
														MarkdownDescription: "Specifies the port on the host that is being addressed.",
														Attributes: map[string]schema.Attribute{
															"number": schema.Int64Attribute{
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

													"subset": schema.StringAttribute{
														Description:         "The name of a subset within the service.",
														MarkdownDescription: "The name of a subset within the service.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
												MarkdownDescription: "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
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

					"tls": schema.ListNestedAttribute{
						Description:         "An ordered list of route rule for non-terminated TLS & HTTPS traffic.",
						MarkdownDescription: "An ordered list of route rule for non-terminated TLS & HTTPS traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match": schema.ListNestedAttribute{
									Description:         "Match conditions to be satisfied for the rule to be activated.",
									MarkdownDescription: "Match conditions to be satisfied for the rule to be activated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination_subnets": schema.ListAttribute{
												Description:         "IPv4 or IPv6 ip addresses of destination with optional subnet.",
												MarkdownDescription: "IPv4 or IPv6 ip addresses of destination with optional subnet.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gateways": schema.ListAttribute{
												Description:         "Names of gateways where the rule should be applied.",
												MarkdownDescription: "Names of gateways where the rule should be applied.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Specifies the port on the host that is being addressed.",
												MarkdownDescription: "Specifies the port on the host that is being addressed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"sni_hosts": schema.ListAttribute{
												Description:         "SNI (server name indicator) to match on.",
												MarkdownDescription: "SNI (server name indicator) to match on.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"source_labels": schema.MapAttribute{
												Description:         "One or more labels that constrain the applicability of a rule to workloads with the given labels.",
												MarkdownDescription: "One or more labels that constrain the applicability of a rule to workloads with the given labels.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_namespace": schema.StringAttribute{
												Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
												MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"route": schema.ListNestedAttribute{
									Description:         "The destination to which the connection should be forwarded to.",
									MarkdownDescription: "The destination to which the connection should be forwarded to.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination": schema.SingleNestedAttribute{
												Description:         "Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to.",
												MarkdownDescription: "Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "Specifies the port on the host that is being addressed.",
														MarkdownDescription: "Specifies the port on the host that is being addressed.",
														Attributes: map[string]schema.Attribute{
															"number": schema.Int64Attribute{
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

													"subset": schema.StringAttribute{
														Description:         "The name of a subset within the service.",
														MarkdownDescription: "The name of a subset within the service.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"weight": schema.Int64Attribute{
												Description:         "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
												MarkdownDescription: "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
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
		},
	}
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_virtual_service_v1alpha3_manifest")

	var model NetworkingIstioIoVirtualServiceV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("VirtualService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
