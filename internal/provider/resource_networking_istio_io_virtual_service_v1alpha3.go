/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type NetworkingIstioIoVirtualServiceV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*NetworkingIstioIoVirtualServiceV1Alpha3Resource)(nil)
)

type NetworkingIstioIoVirtualServiceV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingIstioIoVirtualServiceV1Alpha3GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ExportTo *[]string `tfsdk:"export_to" yaml:"exportTo,omitempty"`

		Gateways *[]string `tfsdk:"gateways" yaml:"gateways,omitempty"`

		Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

		Http *[]struct {
			CorsPolicy *struct {
				AllowCredentials *bool `tfsdk:"allow_credentials" yaml:"allowCredentials,omitempty"`

				AllowHeaders *[]string `tfsdk:"allow_headers" yaml:"allowHeaders,omitempty"`

				AllowMethods *[]string `tfsdk:"allow_methods" yaml:"allowMethods,omitempty"`

				AllowOrigin *[]string `tfsdk:"allow_origin" yaml:"allowOrigin,omitempty"`

				AllowOrigins *[]struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"allow_origins" yaml:"allowOrigins,omitempty"`

				ExposeHeaders *[]string `tfsdk:"expose_headers" yaml:"exposeHeaders,omitempty"`

				MaxAge *string `tfsdk:"max_age" yaml:"maxAge,omitempty"`
			} `tfsdk:"cors_policy" yaml:"corsPolicy,omitempty"`

			Delegate *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"delegate" yaml:"delegate,omitempty"`

			DirectResponse *struct {
				Body *struct {
					Bytes *string `tfsdk:"bytes" yaml:"bytes,omitempty"`

					String *string `tfsdk:"string" yaml:"string,omitempty"`
				} `tfsdk:"body" yaml:"body,omitempty"`

				Status *int64 `tfsdk:"status" yaml:"status,omitempty"`
			} `tfsdk:"direct_response" yaml:"directResponse,omitempty"`

			Fault *struct {
				Abort *struct {
					GrpcStatus *string `tfsdk:"grpc_status" yaml:"grpcStatus,omitempty"`

					Http2Error *string `tfsdk:"http2_error" yaml:"http2Error,omitempty"`

					HttpStatus *int64 `tfsdk:"http_status" yaml:"httpStatus,omitempty"`

					Percentage *struct {
						Value *float64 `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"percentage" yaml:"percentage,omitempty"`
				} `tfsdk:"abort" yaml:"abort,omitempty"`

				Delay *struct {
					ExponentialDelay *string `tfsdk:"exponential_delay" yaml:"exponentialDelay,omitempty"`

					FixedDelay *string `tfsdk:"fixed_delay" yaml:"fixedDelay,omitempty"`

					Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

					Percentage *struct {
						Value *float64 `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"percentage" yaml:"percentage,omitempty"`
				} `tfsdk:"delay" yaml:"delay,omitempty"`
			} `tfsdk:"fault" yaml:"fault,omitempty"`

			Headers *struct {
				Request *struct {
					Add *map[string]string `tfsdk:"add" yaml:"add,omitempty"`

					Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

					Set *map[string]string `tfsdk:"set" yaml:"set,omitempty"`
				} `tfsdk:"request" yaml:"request,omitempty"`

				Response *struct {
					Add *map[string]string `tfsdk:"add" yaml:"add,omitempty"`

					Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

					Set *map[string]string `tfsdk:"set" yaml:"set,omitempty"`
				} `tfsdk:"response" yaml:"response,omitempty"`
			} `tfsdk:"headers" yaml:"headers,omitempty"`

			Match *[]struct {
				Authority *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"authority" yaml:"authority,omitempty"`

				Gateways *[]string `tfsdk:"gateways" yaml:"gateways,omitempty"`

				Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

				IgnoreUriCase *bool `tfsdk:"ignore_uri_case" yaml:"ignoreUriCase,omitempty"`

				Method *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"method" yaml:"method,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				QueryParams *map[string]string `tfsdk:"query_params" yaml:"queryParams,omitempty"`

				Scheme *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"scheme" yaml:"scheme,omitempty"`

				SourceLabels *map[string]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

				SourceNamespace *string `tfsdk:"source_namespace" yaml:"sourceNamespace,omitempty"`

				StatPrefix *string `tfsdk:"stat_prefix" yaml:"statPrefix,omitempty"`

				Uri *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"uri" yaml:"uri,omitempty"`

				WithoutHeaders *map[string]string `tfsdk:"without_headers" yaml:"withoutHeaders,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Mirror *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Port *struct {
					Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`

				Subset *string `tfsdk:"subset" yaml:"subset,omitempty"`
			} `tfsdk:"mirror" yaml:"mirror,omitempty"`

			MirrorPercent *int64 `tfsdk:"mirror_percent" yaml:"mirrorPercent,omitempty"`

			MirrorPercentage *struct {
				Value *float64 `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"mirror_percentage" yaml:"mirrorPercentage,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Redirect *struct {
				Authority *string `tfsdk:"authority" yaml:"authority,omitempty"`

				DerivePort *string `tfsdk:"derive_port" yaml:"derivePort,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				RedirectCode *int64 `tfsdk:"redirect_code" yaml:"redirectCode,omitempty"`

				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

				Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
			} `tfsdk:"redirect" yaml:"redirect,omitempty"`

			Retries *struct {
				Attempts *int64 `tfsdk:"attempts" yaml:"attempts,omitempty"`

				PerTryTimeout *string `tfsdk:"per_try_timeout" yaml:"perTryTimeout,omitempty"`

				RetryOn *string `tfsdk:"retry_on" yaml:"retryOn,omitempty"`

				RetryRemoteLocalities *bool `tfsdk:"retry_remote_localities" yaml:"retryRemoteLocalities,omitempty"`
			} `tfsdk:"retries" yaml:"retries,omitempty"`

			Rewrite *struct {
				Authority *string `tfsdk:"authority" yaml:"authority,omitempty"`

				Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
			} `tfsdk:"rewrite" yaml:"rewrite,omitempty"`

			Route *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *struct {
						Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
					} `tfsdk:"port" yaml:"port,omitempty"`

					Subset *string `tfsdk:"subset" yaml:"subset,omitempty"`
				} `tfsdk:"destination" yaml:"destination,omitempty"`

				Headers *struct {
					Request *struct {
						Add *map[string]string `tfsdk:"add" yaml:"add,omitempty"`

						Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

						Set *map[string]string `tfsdk:"set" yaml:"set,omitempty"`
					} `tfsdk:"request" yaml:"request,omitempty"`

					Response *struct {
						Add *map[string]string `tfsdk:"add" yaml:"add,omitempty"`

						Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

						Set *map[string]string `tfsdk:"set" yaml:"set,omitempty"`
					} `tfsdk:"response" yaml:"response,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"http" yaml:"http,omitempty"`

		Tcp *[]struct {
			Match *[]struct {
				DestinationSubnets *[]string `tfsdk:"destination_subnets" yaml:"destinationSubnets,omitempty"`

				Gateways *[]string `tfsdk:"gateways" yaml:"gateways,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				SourceLabels *map[string]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

				SourceNamespace *string `tfsdk:"source_namespace" yaml:"sourceNamespace,omitempty"`

				SourceSubnet *string `tfsdk:"source_subnet" yaml:"sourceSubnet,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Route *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *struct {
						Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
					} `tfsdk:"port" yaml:"port,omitempty"`

					Subset *string `tfsdk:"subset" yaml:"subset,omitempty"`
				} `tfsdk:"destination" yaml:"destination,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`
		} `tfsdk:"tcp" yaml:"tcp,omitempty"`

		Tls *[]struct {
			Match *[]struct {
				DestinationSubnets *[]string `tfsdk:"destination_subnets" yaml:"destinationSubnets,omitempty"`

				Gateways *[]string `tfsdk:"gateways" yaml:"gateways,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				SniHosts *[]string `tfsdk:"sni_hosts" yaml:"sniHosts,omitempty"`

				SourceLabels *map[string]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

				SourceNamespace *string `tfsdk:"source_namespace" yaml:"sourceNamespace,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Route *[]struct {
				Destination *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *struct {
						Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
					} `tfsdk:"port" yaml:"port,omitempty"`

					Subset *string `tfsdk:"subset" yaml:"subset,omitempty"`
				} `tfsdk:"destination" yaml:"destination,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`
		} `tfsdk:"tls" yaml:"tls,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingIstioIoVirtualServiceV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoVirtualServiceV1Alpha3Resource{}
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_istio_io_virtual_service_v1alpha3"
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html",
				MarkdownDescription: "Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"export_to": {
						Description:         "A list of namespaces to which this virtual service is exported.",
						MarkdownDescription: "A list of namespaces to which this virtual service is exported.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"gateways": {
						Description:         "The names of gateways and sidecars that should apply these routes.",
						MarkdownDescription: "The names of gateways and sidecars that should apply these routes.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hosts": {
						Description:         "The destination hosts to which traffic is being sent.",
						MarkdownDescription: "The destination hosts to which traffic is being sent.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http": {
						Description:         "An ordered list of route rules for HTTP traffic.",
						MarkdownDescription: "An ordered list of route rules for HTTP traffic.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cors_policy": {
								Description:         "Cross-Origin Resource Sharing policy (CORS).",
								MarkdownDescription: "Cross-Origin Resource Sharing policy (CORS).",

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
										Description:         "List of HTTP methods allowed to access the resource.",
										MarkdownDescription: "List of HTTP methods allowed to access the resource.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_origin": {
										Description:         "The list of origins that are allowed to perform CORS requests.",
										MarkdownDescription: "The list of origins that are allowed to perform CORS requests.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_origins": {
										Description:         "String patterns that match allowed origins.",
										MarkdownDescription: "String patterns that match allowed origins.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

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

							"delegate": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name specifies the name of the delegate VirtualService.",
										MarkdownDescription: "Name specifies the name of the delegate VirtualService.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace specifies the namespace where the delegate VirtualService resides.",
										MarkdownDescription: "Namespace specifies the namespace where the delegate VirtualService resides.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"direct_response": {
								Description:         "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
								MarkdownDescription: "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"body": {
										Description:         "Specifies the content of the response body.",
										MarkdownDescription: "Specifies the content of the response body.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"bytes": {
												Description:         "response body as base64 encoded bytes.",
												MarkdownDescription: "response body as base64 encoded bytes.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"string": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"status": {
										Description:         "Specifies the HTTP response status to be returned.",
										MarkdownDescription: "Specifies the HTTP response status to be returned.",

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

							"fault": {
								Description:         "Fault injection policy to apply on HTTP traffic at the client side.",
								MarkdownDescription: "Fault injection policy to apply on HTTP traffic at the client side.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"abort": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"grpc_status": {
												Description:         "GRPC status code to use to abort the request.",
												MarkdownDescription: "GRPC status code to use to abort the request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http2_error": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_status": {
												Description:         "HTTP status code to use to abort the Http request.",
												MarkdownDescription: "HTTP status code to use to abort the Http request.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percentage": {
												Description:         "Percentage of requests to be aborted with the error code provided.",
												MarkdownDescription: "Percentage of requests to be aborted with the error code provided.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"value": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.NumberType,

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

									"delay": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exponential_delay": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fixed_delay": {
												Description:         "Add a fixed delay before forwarding the request.",
												MarkdownDescription: "Add a fixed delay before forwarding the request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percent": {
												Description:         "Percentage of requests on which the delay will be injected (0-100).",
												MarkdownDescription: "Percentage of requests on which the delay will be injected (0-100).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percentage": {
												Description:         "Percentage of requests on which the delay will be injected.",
												MarkdownDescription: "Percentage of requests on which the delay will be injected.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"value": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.NumberType,

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

							"headers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"request": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"add": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"remove": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": {
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

											"add": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"remove": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": {
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

							"match": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"authority": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gateways": {
										Description:         "Names of gateways where the rule should be applied.",
										MarkdownDescription: "Names of gateways where the rule should be applied.",

										Type: types.ListType{ElemType: types.StringType},

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

									"ignore_uri_case": {
										Description:         "Flag to specify whether the URI matching should be case-insensitive.",
										MarkdownDescription: "Flag to specify whether the URI matching should be case-insensitive.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

												Type: types.StringType,

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
										Description:         "The name assigned to a match.",
										MarkdownDescription: "The name assigned to a match.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Specifies the ports on the host that is being addressed.",
										MarkdownDescription: "Specifies the ports on the host that is being addressed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_params": {
										Description:         "Query parameters for matching.",
										MarkdownDescription: "Query parameters for matching.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scheme": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_namespace": {
										Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
										MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stat_prefix": {
										Description:         "The human readable prefix to use when emitting statistics for this route.",
										MarkdownDescription: "The human readable prefix to use when emitting statistics for this route.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"without_headers": {
										Description:         "withoutHeader has the same syntax with the header, but has opposite meaning.",
										MarkdownDescription: "withoutHeader has the same syntax with the header, but has opposite meaning.",

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

							"mirror": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "The name of a service from the service registry.",
										MarkdownDescription: "The name of a service from the service registry.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Specifies the port on the host that is being addressed.",
										MarkdownDescription: "Specifies the port on the host that is being addressed.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"number": {
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

									"subset": {
										Description:         "The name of a subset within the service.",
										MarkdownDescription: "The name of a subset within the service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mirror_percent": {
								Description:         "Percentage of the traffic to be mirrored by the 'mirror' field.",
								MarkdownDescription: "Percentage of the traffic to be mirrored by the 'mirror' field.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mirror_percentage": {
								Description:         "Percentage of the traffic to be mirrored by the 'mirror' field.",
								MarkdownDescription: "Percentage of the traffic to be mirrored by the 'mirror' field.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.NumberType,

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
								Description:         "The name assigned to the route for debugging purposes.",
								MarkdownDescription: "The name assigned to the route for debugging purposes.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"redirect": {
								Description:         "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
								MarkdownDescription: "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"authority": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"derive_port": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("FROM_PROTOCOL_DEFAULT", "FROM_REQUEST_PORT"),
										},
									},

									"port": {
										Description:         "On a redirect, overwrite the port portion of the URL with this value.",
										MarkdownDescription: "On a redirect, overwrite the port portion of the URL with this value.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"redirect_code": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scheme": {
										Description:         "On a redirect, overwrite the scheme portion of the URL with this value.",
										MarkdownDescription: "On a redirect, overwrite the scheme portion of the URL with this value.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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
								Description:         "Retry policy for HTTP requests.",
								MarkdownDescription: "Retry policy for HTTP requests.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"attempts": {
										Description:         "Number of retries to be allowed for a given request.",
										MarkdownDescription: "Number of retries to be allowed for a given request.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"per_try_timeout": {
										Description:         "Timeout per attempt for a given request, including the initial call and any retries.",
										MarkdownDescription: "Timeout per attempt for a given request, including the initial call and any retries.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_on": {
										Description:         "Specifies the conditions under which retry takes place.",
										MarkdownDescription: "Specifies the conditions under which retry takes place.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_remote_localities": {
										Description:         "Flag to specify whether the retries should retry to other localities.",
										MarkdownDescription: "Flag to specify whether the retries should retry to other localities.",

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

							"rewrite": {
								Description:         "Rewrite HTTP URIs and Authority headers.",
								MarkdownDescription: "Rewrite HTTP URIs and Authority headers.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"authority": {
										Description:         "rewrite the Authority/Host header with this value.",
										MarkdownDescription: "rewrite the Authority/Host header with this value.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route": {
								Description:         "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",
								MarkdownDescription: "A HTTP rule can either return a direct_response, redirect or forward (default) traffic.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"destination": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "The name of a service from the service registry.",
												MarkdownDescription: "The name of a service from the service registry.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Specifies the port on the host that is being addressed.",
												MarkdownDescription: "Specifies the port on the host that is being addressed.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"number": {
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

											"subset": {
												Description:         "The name of a subset within the service.",
												MarkdownDescription: "The name of a subset within the service.",

												Type: types.StringType,

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

											"request": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"add": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"set": {
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

													"add": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"set": {
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

									"weight": {
										Description:         "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
										MarkdownDescription: "Weight specifies the relative proportion of traffic to be forwarded to the destination.",

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

							"timeout": {
								Description:         "Timeout for HTTP requests, default is disabled.",
								MarkdownDescription: "Timeout for HTTP requests, default is disabled.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tcp": {
						Description:         "An ordered list of route rules for opaque TCP traffic.",
						MarkdownDescription: "An ordered list of route rules for opaque TCP traffic.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"match": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"destination_subnets": {
										Description:         "IPv4 or IPv6 ip addresses of destination with optional subnet.",
										MarkdownDescription: "IPv4 or IPv6 ip addresses of destination with optional subnet.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gateways": {
										Description:         "Names of gateways where the rule should be applied.",
										MarkdownDescription: "Names of gateways where the rule should be applied.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Specifies the port on the host that is being addressed.",
										MarkdownDescription: "Specifies the port on the host that is being addressed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_namespace": {
										Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
										MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_subnet": {
										Description:         "IPv4 or IPv6 ip address of source with optional subnet.",
										MarkdownDescription: "IPv4 or IPv6 ip address of source with optional subnet.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route": {
								Description:         "The destination to which the connection should be forwarded to.",
								MarkdownDescription: "The destination to which the connection should be forwarded to.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"destination": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "The name of a service from the service registry.",
												MarkdownDescription: "The name of a service from the service registry.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Specifies the port on the host that is being addressed.",
												MarkdownDescription: "Specifies the port on the host that is being addressed.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"number": {
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

											"subset": {
												Description:         "The name of a subset within the service.",
												MarkdownDescription: "The name of a subset within the service.",

												Type: types.StringType,

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
										Description:         "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
										MarkdownDescription: "Weight specifies the relative proportion of traffic to be forwarded to the destination.",

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

					"tls": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"match": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"destination_subnets": {
										Description:         "IPv4 or IPv6 ip addresses of destination with optional subnet.",
										MarkdownDescription: "IPv4 or IPv6 ip addresses of destination with optional subnet.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gateways": {
										Description:         "Names of gateways where the rule should be applied.",
										MarkdownDescription: "Names of gateways where the rule should be applied.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Specifies the port on the host that is being addressed.",
										MarkdownDescription: "Specifies the port on the host that is being addressed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sni_hosts": {
										Description:         "SNI (server name indicator) to match on.",
										MarkdownDescription: "SNI (server name indicator) to match on.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_namespace": {
										Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
										MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route": {
								Description:         "The destination to which the connection should be forwarded to.",
								MarkdownDescription: "The destination to which the connection should be forwarded to.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"destination": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "The name of a service from the service registry.",
												MarkdownDescription: "The name of a service from the service registry.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Specifies the port on the host that is being addressed.",
												MarkdownDescription: "Specifies the port on the host that is being addressed.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"number": {
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

											"subset": {
												Description:         "The name of a subset within the service.",
												MarkdownDescription: "The name of a subset within the service.",

												Type: types.StringType,

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
										Description:         "Weight specifies the relative proportion of traffic to be forwarded to the destination.",
										MarkdownDescription: "Weight specifies the relative proportion of traffic to be forwarded to the destination.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_virtual_service_v1alpha3")

	var state NetworkingIstioIoVirtualServiceV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoVirtualServiceV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("VirtualService")

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

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_virtual_service_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_virtual_service_v1alpha3")

	var state NetworkingIstioIoVirtualServiceV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoVirtualServiceV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("VirtualService")

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

func (r *NetworkingIstioIoVirtualServiceV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_virtual_service_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
