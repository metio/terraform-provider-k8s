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

type NetworkingIstioIoEnvoyFilterV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*NetworkingIstioIoEnvoyFilterV1Alpha3Resource)(nil)
)

type NetworkingIstioIoEnvoyFilterV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingIstioIoEnvoyFilterV1Alpha3GoModel struct {
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
		ConfigPatches *[]struct {
			ApplyTo *string `tfsdk:"apply_to" yaml:"applyTo,omitempty"`

			Match *struct {
				Cluster *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					PortNumber *int64 `tfsdk:"port_number" yaml:"portNumber,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`

					Subset *string `tfsdk:"subset" yaml:"subset,omitempty"`
				} `tfsdk:"cluster" yaml:"cluster,omitempty"`

				Context *string `tfsdk:"context" yaml:"context,omitempty"`

				Listener *struct {
					FilterChain *struct {
						ApplicationProtocols *string `tfsdk:"application_protocols" yaml:"applicationProtocols,omitempty"`

						DestinationPort *int64 `tfsdk:"destination_port" yaml:"destinationPort,omitempty"`

						Filter *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							SubFilter *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"sub_filter" yaml:"subFilter,omitempty"`
						} `tfsdk:"filter" yaml:"filter,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

						TransportProtocol *string `tfsdk:"transport_protocol" yaml:"transportProtocol,omitempty"`
					} `tfsdk:"filter_chain" yaml:"filterChain,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`

					PortNumber *int64 `tfsdk:"port_number" yaml:"portNumber,omitempty"`
				} `tfsdk:"listener" yaml:"listener,omitempty"`

				Proxy *struct {
					Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

					ProxyVersion *string `tfsdk:"proxy_version" yaml:"proxyVersion,omitempty"`
				} `tfsdk:"proxy" yaml:"proxy,omitempty"`

				RouteConfiguration *struct {
					Gateway *string `tfsdk:"gateway" yaml:"gateway,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`

					PortNumber *int64 `tfsdk:"port_number" yaml:"portNumber,omitempty"`

					Vhost *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Route *struct {
							Action *string `tfsdk:"action" yaml:"action,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"route" yaml:"route,omitempty"`
					} `tfsdk:"vhost" yaml:"vhost,omitempty"`
				} `tfsdk:"route_configuration" yaml:"routeConfiguration,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Patch *struct {
				FilterClass *string `tfsdk:"filter_class" yaml:"filterClass,omitempty"`

				Operation *string `tfsdk:"operation" yaml:"operation,omitempty"`

				Value *map[string]string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"patch" yaml:"patch,omitempty"`
		} `tfsdk:"config_patches" yaml:"configPatches,omitempty"`

		Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
		} `tfsdk:"workload_selector" yaml:"workloadSelector,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingIstioIoEnvoyFilterV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoEnvoyFilterV1Alpha3Resource{}
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_istio_io_envoy_filter_v1alpha3"
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html",
				MarkdownDescription: "Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"config_patches": {
						Description:         "One or more patches with match conditions.",
						MarkdownDescription: "One or more patches with match conditions.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"apply_to": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("INVALID", "LISTENER", "FILTER_CHAIN", "NETWORK_FILTER", "HTTP_FILTER", "ROUTE_CONFIGURATION", "VIRTUAL_HOST", "HTTP_ROUTE", "CLUSTER", "EXTENSION_CONFIG", "BOOTSTRAP"),
								},
							},

							"match": {
								Description:         "Match on listener/route configuration/cluster.",
								MarkdownDescription: "Match on listener/route configuration/cluster.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cluster": {
										Description:         "Match on envoy cluster attributes.",
										MarkdownDescription: "Match on envoy cluster attributes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "The exact name of the cluster to match.",
												MarkdownDescription: "The exact name of the cluster to match.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port_number": {
												Description:         "The service port for which this cluster was generated.",
												MarkdownDescription: "The service port for which this cluster was generated.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service": {
												Description:         "The fully qualified service name for this cluster.",
												MarkdownDescription: "The fully qualified service name for this cluster.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subset": {
												Description:         "The subset associated with the service.",
												MarkdownDescription: "The subset associated with the service.",

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

									"context": {
										Description:         "The specific config generation context to match on.",
										MarkdownDescription: "The specific config generation context to match on.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ANY", "SIDECAR_INBOUND", "SIDECAR_OUTBOUND", "GATEWAY"),
										},
									},

									"listener": {
										Description:         "Match on envoy listener attributes.",
										MarkdownDescription: "Match on envoy listener attributes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"filter_chain": {
												Description:         "Match a specific filter chain in a listener.",
												MarkdownDescription: "Match a specific filter chain in a listener.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"application_protocols": {
														Description:         "Applies only to sidecars.",
														MarkdownDescription: "Applies only to sidecars.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"destination_port": {
														Description:         "The destination_port value used by a filter chain's match condition.",
														MarkdownDescription: "The destination_port value used by a filter chain's match condition.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"filter": {
														Description:         "The name of a specific filter to apply the patch to.",
														MarkdownDescription: "The name of a specific filter to apply the patch to.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "The filter name to match on.",
																MarkdownDescription: "The filter name to match on.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_filter": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The filter name to match on.",
																		MarkdownDescription: "The filter name to match on.",

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

													"name": {
														Description:         "The name assigned to the filter chain.",
														MarkdownDescription: "The name assigned to the filter chain.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sni": {
														Description:         "The SNI value used by a filter chain's match condition.",
														MarkdownDescription: "The SNI value used by a filter chain's match condition.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"transport_protocol": {
														Description:         "Applies only to 'SIDECAR_INBOUND' context.",
														MarkdownDescription: "Applies only to 'SIDECAR_INBOUND' context.",

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
												Description:         "Match a specific listener by its name.",
												MarkdownDescription: "Match a specific listener by its name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port_number": {
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

									"proxy": {
										Description:         "Match on properties associated with a proxy.",
										MarkdownDescription: "Match on properties associated with a proxy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_version": {
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

									"route_configuration": {
										Description:         "Match on envoy HTTP route configuration attributes.",
										MarkdownDescription: "Match on envoy HTTP route configuration attributes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"gateway": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Route configuration name to match on.",
												MarkdownDescription: "Route configuration name to match on.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port_name": {
												Description:         "Applicable only for GATEWAY context.",
												MarkdownDescription: "Applicable only for GATEWAY context.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port_number": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"vhost": {
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

													"route": {
														Description:         "Match a specific route within the virtual host.",
														MarkdownDescription: "Match a specific route within the virtual host.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"action": {
																Description:         "Match a route with specific action type.",
																MarkdownDescription: "Match a route with specific action type.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("ANY", "ROUTE", "REDIRECT", "DIRECT_RESPONSE"),
																},
															},

															"name": {
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

							"patch": {
								Description:         "The patch to apply along with the operation.",
								MarkdownDescription: "The patch to apply along with the operation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"filter_class": {
										Description:         "Determines the filter insertion order.",
										MarkdownDescription: "Determines the filter insertion order.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("UNSPECIFIED", "AUTHN", "AUTHZ", "STATS"),
										},
									},

									"operation": {
										Description:         "Determines how the patch should be applied.",
										MarkdownDescription: "Determines how the patch should be applied.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("INVALID", "MERGE", "ADD", "REMOVE", "INSERT_BEFORE", "INSERT_AFTER", "INSERT_FIRST", "REPLACE"),
										},
									},

									"value": {
										Description:         "The JSON config of the object being patched.",
										MarkdownDescription: "The JSON config of the object being patched.",

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

					"priority": {
						Description:         "Priority defines the order in which patch sets are applied within a context.",
						MarkdownDescription: "Priority defines the order in which patch sets are applied within a context.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"workload_selector": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"labels": {
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
		},
	}, nil
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_envoy_filter_v1alpha3")

	var state NetworkingIstioIoEnvoyFilterV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoEnvoyFilterV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("EnvoyFilter")

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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_envoy_filter_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_envoy_filter_v1alpha3")

	var state NetworkingIstioIoEnvoyFilterV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoEnvoyFilterV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("EnvoyFilter")

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

func (r *NetworkingIstioIoEnvoyFilterV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_envoy_filter_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
