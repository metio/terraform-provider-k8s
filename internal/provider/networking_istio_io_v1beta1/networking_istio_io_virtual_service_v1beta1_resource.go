/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &NetworkingIstioIoVirtualServiceV1Beta1Resource{}
	_ resource.ResourceWithConfigure   = &NetworkingIstioIoVirtualServiceV1Beta1Resource{}
	_ resource.ResourceWithImportState = &NetworkingIstioIoVirtualServiceV1Beta1Resource{}
)

func NewNetworkingIstioIoVirtualServiceV1Beta1Resource() resource.Resource {
	return &NetworkingIstioIoVirtualServiceV1Beta1Resource{}
}

type NetworkingIstioIoVirtualServiceV1Beta1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type NetworkingIstioIoVirtualServiceV1Beta1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

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
				ExposeHeaders *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
				MaxAge        *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
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

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_virtual_service_v1beta1"
}

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
											Description:         "List of HTTP methods allowed to access the resource.",
											MarkdownDescription: "List of HTTP methods allowed to access the resource.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allow_origin": schema.ListAttribute{
											Description:         "The list of origins that are allowed to perform CORS requests.",
											MarkdownDescription: "The list of origins that are allowed to perform CORS requests.",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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

								"delegate": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
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
											Required:            false,
											Optional:            true,
											Computed:            false,
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
											Description:         "",
											MarkdownDescription: "",
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
											Description:         "",
											MarkdownDescription: "",
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
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"authority": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
												Description:         "",
												MarkdownDescription: "",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
												Description:         "",
												MarkdownDescription: "",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
												Description:         "",
												MarkdownDescription: "",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
												Description:         "",
												MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
														Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
														MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "The name of a service from the service registry.",
											MarkdownDescription: "The name of a service from the service registry.",
											Required:            false,
											Optional:            true,
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
									Description:         "Percentage of the traffic to be mirrored by the 'mirror' field.",
									MarkdownDescription: "Percentage of the traffic to be mirrored by the 'mirror' field.",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"destination": schema.SingleNestedAttribute{
												Description:         "Destination specifies the target of the mirror operation.",
												MarkdownDescription: "Destination specifies the target of the mirror operation.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            false,
														Optional:            true,
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
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"derive_port": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
										},

										"redirect_code": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scheme": schema.StringAttribute{
											Description:         "On a redirect, overwrite the scheme portion of the URL with this value.",
											MarkdownDescription: "On a redirect, overwrite the scheme portion of the URL with this value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uri": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uri_regex_rewrite": schema.SingleNestedAttribute{
											Description:         "rewrite the path portion of the URI with the specified regex.",
											MarkdownDescription: "rewrite the path portion of the URI with the specified regex.",
											Attributes: map[string]schema.Attribute{
												"match": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            false,
														Optional:            true,
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
									Description:         "",
									MarkdownDescription: "",
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
											},

											"source_labels": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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
												Description:         "IPv4 or IPv6 ip address of source with optional subnet.",
												MarkdownDescription: "IPv4 or IPv6 ip address of source with optional subnet.",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            false,
														Optional:            true,
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
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
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
											},

											"sni_hosts": schema.ListAttribute{
												Description:         "SNI (server name indicator) to match on.",
												MarkdownDescription: "SNI (server name indicator) to match on.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_labels": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "The name of a service from the service registry.",
														MarkdownDescription: "The name of a service from the service registry.",
														Required:            false,
														Optional:            true,
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

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_virtual_service_v1beta1")

	var model NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("networking.istio.io/v1beta1")
	model.Kind = pointer.String("VirtualService")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "virtualservices"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_virtual_service_v1beta1")

	var data NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "virtualservices"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_virtual_service_v1beta1")

	var model NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1beta1")
	model.Kind = pointer.String("VirtualService")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "virtualservices"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_virtual_service_v1beta1")

	var data NetworkingIstioIoVirtualServiceV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "virtualservices"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "virtualservices"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *NetworkingIstioIoVirtualServiceV1Beta1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
