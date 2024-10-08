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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &NetworkingIstioIoEnvoyFilterV1Alpha3Manifest{}
)

func NewNetworkingIstioIoEnvoyFilterV1Alpha3Manifest() datasource.DataSource {
	return &NetworkingIstioIoEnvoyFilterV1Alpha3Manifest{}
}

type NetworkingIstioIoEnvoyFilterV1Alpha3Manifest struct{}

type NetworkingIstioIoEnvoyFilterV1Alpha3ManifestData struct {
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
		Priority   *int64 `tfsdk:"priority" json:"priority,omitempty"`
		TargetRefs *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_envoy_filter_v1alpha3_manifest"
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html",
				MarkdownDescription: "Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html",
				Attributes: map[string]schema.Attribute{
					"config_patches": schema.ListNestedAttribute{
						Description:         "One or more patches with match conditions.",
						MarkdownDescription: "One or more patches with match conditions.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"apply_to": schema.StringAttribute{
									Description:         "Specifies where in the Envoy configuration, the patch should be applied. Valid Options: LISTENER, FILTER_CHAIN, NETWORK_FILTER, HTTP_FILTER, ROUTE_CONFIGURATION, VIRTUAL_HOST, HTTP_ROUTE, CLUSTER, EXTENSION_CONFIG, BOOTSTRAP, LISTENER_FILTER",
									MarkdownDescription: "Specifies where in the Envoy configuration, the patch should be applied. Valid Options: LISTENER, FILTER_CHAIN, NETWORK_FILTER, HTTP_FILTER, ROUTE_CONFIGURATION, VIRTUAL_HOST, HTTP_ROUTE, CLUSTER, EXTENSION_CONFIG, BOOTSTRAP, LISTENER_FILTER",
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
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(4.294967295e+09),
													},
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
											Description:         "The specific config generation context to match on. Valid Options: ANY, SIDECAR_INBOUND, SIDECAR_OUTBOUND, GATEWAY",
											MarkdownDescription: "The specific config generation context to match on. Valid Options: ANY, SIDECAR_INBOUND, SIDECAR_OUTBOUND, GATEWAY",
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
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(4.294967295e+09),
															},
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
																	Description:         "The next level filter within this filter to match upon.",
																	MarkdownDescription: "The next level filter within this filter to match upon.",
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
													Description:         "The service port/gateway port to which traffic is being sent/received.",
													MarkdownDescription: "The service port/gateway port to which traffic is being sent/received.",
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

										"proxy": schema.SingleNestedAttribute{
											Description:         "Match on properties associated with a proxy.",
											MarkdownDescription: "Match on properties associated with a proxy.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.MapAttribute{
													Description:         "Match on the node metadata supplied by a proxy when connecting to Istio Pilot.",
													MarkdownDescription: "Match on the node metadata supplied by a proxy when connecting to Istio Pilot.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy_version": schema.StringAttribute{
													Description:         "A regular expression in golang regex format (RE2) that can be used to select proxies using a specific version of istio proxy.",
													MarkdownDescription: "A regular expression in golang regex format (RE2) that can be used to select proxies using a specific version of istio proxy.",
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
													Description:         "The Istio gateway config's namespace/name for which this route configuration was generated.",
													MarkdownDescription: "The Istio gateway config's namespace/name for which this route configuration was generated.",
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
													Description:         "The service port number or gateway server port number for which this route configuration was generated.",
													MarkdownDescription: "The service port number or gateway server port number for which this route configuration was generated.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(4.294967295e+09),
													},
												},

												"vhost": schema.SingleNestedAttribute{
													Description:         "Match a specific virtual host in a route configuration and apply the patch to the virtual host.",
													MarkdownDescription: "Match a specific virtual host in a route configuration and apply the patch to the virtual host.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "The VirtualHosts objects generated by Istio are named as host:port, where the host typically corresponds to the VirtualService's host field or the hostname of a service in the registry.",
															MarkdownDescription: "The VirtualHosts objects generated by Istio are named as host:port, where the host typically corresponds to the VirtualService's host field or the hostname of a service in the registry.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"route": schema.SingleNestedAttribute{
															Description:         "Match a specific route within the virtual host.",
															MarkdownDescription: "Match a specific route within the virtual host.",
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "Match a route with specific action type. Valid Options: ANY, ROUTE, REDIRECT, DIRECT_RESPONSE",
																	MarkdownDescription: "Match a route with specific action type. Valid Options: ANY, ROUTE, REDIRECT, DIRECT_RESPONSE",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ANY", "ROUTE", "REDIRECT", "DIRECT_RESPONSE"),
																	},
																},

																"name": schema.StringAttribute{
																	Description:         "The Route objects generated by default are named as default.",
																	MarkdownDescription: "The Route objects generated by default are named as default.",
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
											Description:         "Determines the filter insertion order. Valid Options: AUTHN, AUTHZ, STATS",
											MarkdownDescription: "Determines the filter insertion order. Valid Options: AUTHN, AUTHZ, STATS",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("UNSPECIFIED", "AUTHN", "AUTHZ", "STATS"),
											},
										},

										"operation": schema.StringAttribute{
											Description:         "Determines how the patch should be applied. Valid Options: MERGE, ADD, REMOVE, INSERT_BEFORE, INSERT_AFTER, INSERT_FIRST, REPLACE",
											MarkdownDescription: "Determines how the patch should be applied. Valid Options: MERGE, ADD, REMOVE, INSERT_BEFORE, INSERT_AFTER, INSERT_FIRST, REPLACE",
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

					"target_refs": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "group is the group of the target resource.",
									MarkdownDescription: "group is the group of the target resource.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "kind is kind of the target resource.",
									MarkdownDescription: "kind is kind of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "name is the name of the target resource.",
									MarkdownDescription: "name is the name of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "namespace is the namespace of the referent.",
									MarkdownDescription: "namespace is the namespace of the referent.",
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

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "Criteria used to select the specific set of pods/VMs on which this patch configuration should be applied.",
						MarkdownDescription: "Criteria used to select the specific set of pods/VMs on which this patch configuration should be applied.",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
								Description:         "One or more labels that indicate a specific set of pods/VMs on which the configuration should be applied.",
								MarkdownDescription: "One or more labels that indicate a specific set of pods/VMs on which the configuration should be applied.",
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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_envoy_filter_v1alpha3_manifest")

	var model NetworkingIstioIoEnvoyFilterV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("EnvoyFilter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
