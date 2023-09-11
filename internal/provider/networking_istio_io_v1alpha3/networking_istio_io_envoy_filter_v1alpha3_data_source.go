/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

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
	_ datasource.DataSource              = &NetworkingIstioIoEnvoyFilterV1Alpha3DataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkingIstioIoEnvoyFilterV1Alpha3DataSource{}
)

func NewNetworkingIstioIoEnvoyFilterV1Alpha3DataSource() datasource.DataSource {
	return &NetworkingIstioIoEnvoyFilterV1Alpha3DataSource{}
}

type NetworkingIstioIoEnvoyFilterV1Alpha3DataSource struct {
	kubernetesClient dynamic.Interface
}

type NetworkingIstioIoEnvoyFilterV1Alpha3DataSourceData struct {
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_envoy_filter_v1alpha3"
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
									Optional:            false,
									Computed:            true,
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
													Optional:            false,
													Computed:            true,
												},

												"port_number": schema.Int64Attribute{
													Description:         "The service port for which this cluster was generated.",
													MarkdownDescription: "The service port for which this cluster was generated.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "The fully qualified service name for this cluster.",
													MarkdownDescription: "The fully qualified service name for this cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"subset": schema.StringAttribute{
													Description:         "The subset associated with the service.",
													MarkdownDescription: "The subset associated with the service.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"context": schema.StringAttribute{
											Description:         "The specific config generation context to match on.",
											MarkdownDescription: "The specific config generation context to match on.",
											Required:            false,
											Optional:            false,
											Computed:            true,
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
															Optional:            false,
															Computed:            true,
														},

														"destination_port": schema.Int64Attribute{
															Description:         "The destination_port value used by a filter chain's match condition.",
															MarkdownDescription: "The destination_port value used by a filter chain's match condition.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"filter": schema.SingleNestedAttribute{
															Description:         "The name of a specific filter to apply the patch to.",
															MarkdownDescription: "The name of a specific filter to apply the patch to.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "The filter name to match on.",
																	MarkdownDescription: "The filter name to match on.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"sub_filter": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "The filter name to match on.",
																			MarkdownDescription: "The filter name to match on.",
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

														"name": schema.StringAttribute{
															Description:         "The name assigned to the filter chain.",
															MarkdownDescription: "The name assigned to the filter chain.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"sni": schema.StringAttribute{
															Description:         "The SNI value used by a filter chain's match condition.",
															MarkdownDescription: "The SNI value used by a filter chain's match condition.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"transport_protocol": schema.StringAttribute{
															Description:         "Applies only to 'SIDECAR_INBOUND' context.",
															MarkdownDescription: "Applies only to 'SIDECAR_INBOUND' context.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"listener_filter": schema.StringAttribute{
													Description:         "Match a specific listener filter.",
													MarkdownDescription: "Match a specific listener filter.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Match a specific listener by its name.",
													MarkdownDescription: "Match a specific listener by its name.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port_number": schema.Int64Attribute{
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

										"proxy": schema.SingleNestedAttribute{
											Description:         "Match on properties associated with a proxy.",
											MarkdownDescription: "Match on properties associated with a proxy.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"proxy_version": schema.StringAttribute{
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

										"route_configuration": schema.SingleNestedAttribute{
											Description:         "Match on envoy HTTP route configuration attributes.",
											MarkdownDescription: "Match on envoy HTTP route configuration attributes.",
											Attributes: map[string]schema.Attribute{
												"gateway": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Route configuration name to match on.",
													MarkdownDescription: "Route configuration name to match on.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port_name": schema.StringAttribute{
													Description:         "Applicable only for GATEWAY context.",
													MarkdownDescription: "Applicable only for GATEWAY context.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port_number": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"vhost": schema.SingleNestedAttribute{
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

														"route": schema.SingleNestedAttribute{
															Description:         "Match a specific route within the virtual host.",
															MarkdownDescription: "Match a specific route within the virtual host.",
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "Match a route with specific action type.",
																	MarkdownDescription: "Match a route with specific action type.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"patch": schema.SingleNestedAttribute{
									Description:         "The patch to apply along with the operation.",
									MarkdownDescription: "The patch to apply along with the operation.",
									Attributes: map[string]schema.Attribute{
										"filter_class": schema.StringAttribute{
											Description:         "Determines the filter insertion order.",
											MarkdownDescription: "Determines the filter insertion order.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operation": schema.StringAttribute{
											Description:         "Determines how the patch should be applied.",
											MarkdownDescription: "Determines how the patch should be applied.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.MapAttribute{
											Description:         "The JSON config of the object being patched.",
											MarkdownDescription: "The JSON config of the object being patched.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority defines the order in which patch sets are applied within a context.",
						MarkdownDescription: "Priority defines the order in which patch sets are applied within a context.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_networking_istio_io_envoy_filter_v1alpha3")

	var data NetworkingIstioIoEnvoyFilterV1Alpha3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
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

	var readResponse NetworkingIstioIoEnvoyFilterV1Alpha3DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	data.Kind = pointer.String("EnvoyFilter")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
