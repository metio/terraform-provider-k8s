/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

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
	_ resource.Resource                = &NetworkingIstioIoEnvoyFilterV1Alpha3Resource{}
	_ resource.ResourceWithConfigure   = &NetworkingIstioIoEnvoyFilterV1Alpha3Resource{}
	_ resource.ResourceWithImportState = &NetworkingIstioIoEnvoyFilterV1Alpha3Resource{}
)

func NewNetworkingIstioIoEnvoyFilterV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoEnvoyFilterV1Alpha3Resource{}
}

type NetworkingIstioIoEnvoyFilterV1Alpha3Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData struct {
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
		ConfigPatches *[]struct {
			ApplyTo *string `tfsdk:"apply_to" json:"applyTo,omitempty"`
			Match   *struct {
				Cluster *struct {
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					PortNumber *int64  `tfsdk:"port_number" json:"portNumber,omitempty"`
					Service    *string `tfsdk:"service" json:"service,omitempty"`
					Subset     *string `tfsdk:"subset" json:"subset,omitempty"`
				} `tfsdk:"cluster" json:"cluster,omitempty"`
				Context  *string `tfsdk:"context" json:"context,omitempty"`
				Listener *struct {
					FilterChain *struct {
						ApplicationProtocols *string `tfsdk:"application_protocols" json:"applicationProtocols,omitempty"`
						DestinationPort      *int64  `tfsdk:"destination_port" json:"destinationPort,omitempty"`
						Filter               *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							SubFilter *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"sub_filter" json:"subFilter,omitempty"`
						} `tfsdk:"filter" json:"filter,omitempty"`
						Name              *string `tfsdk:"name" json:"name,omitempty"`
						Sni               *string `tfsdk:"sni" json:"sni,omitempty"`
						TransportProtocol *string `tfsdk:"transport_protocol" json:"transportProtocol,omitempty"`
					} `tfsdk:"filter_chain" json:"filterChain,omitempty"`
					ListenerFilter *string `tfsdk:"listener_filter" json:"listenerFilter,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					PortName       *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortNumber     *int64  `tfsdk:"port_number" json:"portNumber,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				Proxy *struct {
					Metadata     *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
					ProxyVersion *string            `tfsdk:"proxy_version" json:"proxyVersion,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RouteConfiguration *struct {
					Gateway    *string `tfsdk:"gateway" json:"gateway,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					PortName   *string `tfsdk:"port_name" json:"portName,omitempty"`
					PortNumber *int64  `tfsdk:"port_number" json:"portNumber,omitempty"`
					Vhost      *struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Route *struct {
							Action *string `tfsdk:"action" json:"action,omitempty"`
							Name   *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"route" json:"route,omitempty"`
					} `tfsdk:"vhost" json:"vhost,omitempty"`
				} `tfsdk:"route_configuration" json:"routeConfiguration,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Patch *struct {
				FilterClass *string            `tfsdk:"filter_class" json:"filterClass,omitempty"`
				Operation   *string            `tfsdk:"operation" json:"operation,omitempty"`
				Value       *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"patch" json:"patch,omitempty"`
		} `tfsdk:"config_patches" json:"configPatches,omitempty"`
		Priority         *int64 `tfsdk:"priority" json:"priority,omitempty"`
		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_envoy_filter_v1alpha3"
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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
				Description:         "Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html",
				MarkdownDescription: "Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html",
				Attributes: map[string]schema.Attribute{
					"config_patches": schema.ListNestedAttribute{
						Description:         "One or more patches with match conditions.",
						MarkdownDescription: "One or more patches with match conditions.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"apply_to": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("INVALID", "LISTENER", "FILTER_CHAIN", "NETWORK_FILTER", "HTTP_FILTER", "ROUTE_CONFIGURATION", "VIRTUAL_HOST", "HTTP_ROUTE", "CLUSTER", "EXTENSION_CONFIG", "BOOTSTRAP", "LISTENER_FILTER"),
									},
								},

								"match": schema.SingleNestedAttribute{
									Description:         "Match on listener/route configuration/cluster.",
									MarkdownDescription: "Match on listener/route configuration/cluster.",
									Attributes: map[string]schema.Attribute{
										"cluster": schema.SingleNestedAttribute{
											Description:         "Match on envoy cluster attributes.",
											MarkdownDescription: "Match on envoy cluster attributes.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The exact name of the cluster to match.",
													MarkdownDescription: "The exact name of the cluster to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_number": schema.Int64Attribute{
													Description:         "The service port for which this cluster was generated.",
													MarkdownDescription: "The service port for which this cluster was generated.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service": schema.StringAttribute{
													Description:         "The fully qualified service name for this cluster.",
													MarkdownDescription: "The fully qualified service name for this cluster.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"subset": schema.StringAttribute{
													Description:         "The subset associated with the service.",
													MarkdownDescription: "The subset associated with the service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"context": schema.StringAttribute{
											Description:         "The specific config generation context to match on.",
											MarkdownDescription: "The specific config generation context to match on.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("ANY", "SIDECAR_INBOUND", "SIDECAR_OUTBOUND", "GATEWAY"),
											},
										},

										"listener": schema.SingleNestedAttribute{
											Description:         "Match on envoy listener attributes.",
											MarkdownDescription: "Match on envoy listener attributes.",
											Attributes: map[string]schema.Attribute{
												"filter_chain": schema.SingleNestedAttribute{
													Description:         "Match a specific filter chain in a listener.",
													MarkdownDescription: "Match a specific filter chain in a listener.",
													Attributes: map[string]schema.Attribute{
														"application_protocols": schema.StringAttribute{
															Description:         "Applies only to sidecars.",
															MarkdownDescription: "Applies only to sidecars.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"destination_port": schema.Int64Attribute{
															Description:         "The destination_port value used by a filter chain's match condition.",
															MarkdownDescription: "The destination_port value used by a filter chain's match condition.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"filter": schema.SingleNestedAttribute{
															Description:         "The name of a specific filter to apply the patch to.",
															MarkdownDescription: "The name of a specific filter to apply the patch to.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "The filter name to match on.",
																	MarkdownDescription: "The filter name to match on.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_filter": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "The filter name to match on.",
																			MarkdownDescription: "The filter name to match on.",
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

														"name": schema.StringAttribute{
															Description:         "The name assigned to the filter chain.",
															MarkdownDescription: "The name assigned to the filter chain.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sni": schema.StringAttribute{
															Description:         "The SNI value used by a filter chain's match condition.",
															MarkdownDescription: "The SNI value used by a filter chain's match condition.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"transport_protocol": schema.StringAttribute{
															Description:         "Applies only to 'SIDECAR_INBOUND' context.",
															MarkdownDescription: "Applies only to 'SIDECAR_INBOUND' context.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"listener_filter": schema.StringAttribute{
													Description:         "Match a specific listener filter.",
													MarkdownDescription: "Match a specific listener filter.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Match a specific listener by its name.",
													MarkdownDescription: "Match a specific listener by its name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_number": schema.Int64Attribute{
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

										"proxy": schema.SingleNestedAttribute{
											Description:         "Match on properties associated with a proxy.",
											MarkdownDescription: "Match on properties associated with a proxy.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy_version": schema.StringAttribute{
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

										"route_configuration": schema.SingleNestedAttribute{
											Description:         "Match on envoy HTTP route configuration attributes.",
											MarkdownDescription: "Match on envoy HTTP route configuration attributes.",
											Attributes: map[string]schema.Attribute{
												"gateway": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Route configuration name to match on.",
													MarkdownDescription: "Route configuration name to match on.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_name": schema.StringAttribute{
													Description:         "Applicable only for GATEWAY context.",
													MarkdownDescription: "Applicable only for GATEWAY context.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_number": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"vhost": schema.SingleNestedAttribute{
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

														"route": schema.SingleNestedAttribute{
															Description:         "Match a specific route within the virtual host.",
															MarkdownDescription: "Match a specific route within the virtual host.",
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "Match a route with specific action type.",
																	MarkdownDescription: "Match a route with specific action type.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ANY", "ROUTE", "REDIRECT", "DIRECT_RESPONSE"),
																	},
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

								"patch": schema.SingleNestedAttribute{
									Description:         "The patch to apply along with the operation.",
									MarkdownDescription: "The patch to apply along with the operation.",
									Attributes: map[string]schema.Attribute{
										"filter_class": schema.StringAttribute{
											Description:         "Determines the filter insertion order.",
											MarkdownDescription: "Determines the filter insertion order.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("UNSPECIFIED", "AUTHN", "AUTHZ", "STATS"),
											},
										},

										"operation": schema.StringAttribute{
											Description:         "Determines how the patch should be applied.",
											MarkdownDescription: "Determines how the patch should be applied.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("INVALID", "MERGE", "ADD", "REMOVE", "INSERT_BEFORE", "INSERT_AFTER", "INSERT_FIRST", "REPLACE"),
											},
										},

										"value": schema.MapAttribute{
											Description:         "The JSON config of the object being patched.",
											MarkdownDescription: "The JSON config of the object being patched.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority defines the order in which patch sets are applied within a context.",
						MarkdownDescription: "Priority defines the order in which patch sets are applied within a context.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_envoy_filter_v1alpha3")

	var model NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("EnvoyFilter")

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
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "envoyfilters"}).
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

	var readResponse NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_envoy_filter_v1alpha3")

	var data NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "envoyfilters"}).
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

	var readResponse NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_envoy_filter_v1alpha3")

	var model NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("EnvoyFilter")

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
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "envoyfilters"}).
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

	var readResponse NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_envoy_filter_v1alpha3")

	var data NetworkingIstioIoEnvoyFilterV1Alpha3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "envoyfilters"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "envoyfilters"}).
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
