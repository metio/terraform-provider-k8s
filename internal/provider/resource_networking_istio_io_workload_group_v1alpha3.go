/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type NetworkingIstioIoWorkloadGroupV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*NetworkingIstioIoWorkloadGroupV1Alpha3Resource)(nil)
)

type NetworkingIstioIoWorkloadGroupV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingIstioIoWorkloadGroupV1Alpha3GoModel struct {
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
		Metadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
		} `tfsdk:"metadata" yaml:"metadata,omitempty"`

		Probe *struct {
			HttpGet *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				HttpHeaders *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
			} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

			PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

			SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

			TcpSocket *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
			} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

			Exec *struct {
				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
			} `tfsdk:"exec" yaml:"exec,omitempty"`

			FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`
		} `tfsdk:"probe" yaml:"probe,omitempty"`

		Template *struct {
			Address *string `tfsdk:"address" yaml:"address,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			Locality *string `tfsdk:"locality" yaml:"locality,omitempty"`

			Network *string `tfsdk:"network" yaml:"network,omitempty"`

			Ports *map[string]string `tfsdk:"ports" yaml:"ports,omitempty"`

			ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

			Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingIstioIoWorkloadGroupV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoWorkloadGroupV1Alpha3Resource{}
}

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_istio_io_workload_group_v1alpha3"
}

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html",
				MarkdownDescription: "Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"metadata": {
						Description:         "Metadata that will be used for all corresponding 'WorkloadEntries'.",
						MarkdownDescription: "Metadata that will be used for all corresponding 'WorkloadEntries'.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

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

					"probe": {
						Description:         "'ReadinessProbe' describes the configuration the user must provide for healthchecking on their workload.",
						MarkdownDescription: "'ReadinessProbe' describes the configuration the user must provide for healthchecking on their workload.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"http_get": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host name to connect to, defaults to the pod IP.",
										MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_headers": {
										Description:         "Headers the proxy will pass on to make the request.",
										MarkdownDescription: "Headers the proxy will pass on to make the request.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
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

									"path": {
										Description:         "Path to access on the HTTP server.",
										MarkdownDescription: "Path to access on the HTTP server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port on which the endpoint lives.",
										MarkdownDescription: "Port on which the endpoint lives.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scheme": {
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

							"initial_delay_seconds": {
								Description:         "Number of seconds after the container has started before readiness probes are initiated.",
								MarkdownDescription: "Number of seconds after the container has started before readiness probes are initiated.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"period_seconds": {
								Description:         "How often (in seconds) to perform the probe.",
								MarkdownDescription: "How often (in seconds) to perform the probe.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"success_threshold": {
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tcp_socket": {
								Description:         "Health is determined by if the proxy is able to connect.",
								MarkdownDescription: "Health is determined by if the proxy is able to connect.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
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

							"timeout_seconds": {
								Description:         "Number of seconds after which the probe times out.",
								MarkdownDescription: "Number of seconds after which the probe times out.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exec": {
								Description:         "Health is determined by how the command that is executed exited.",
								MarkdownDescription: "Health is determined by how the command that is executed exited.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"command": {
										Description:         "Command to run.",
										MarkdownDescription: "Command to run.",

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

							"failure_threshold": {
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.",

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

					"template": {
						Description:         "Template to be used for the generation of 'WorkloadEntry' resources that belong to this 'WorkloadGroup'.",
						MarkdownDescription: "Template to be used for the generation of 'WorkloadEntry' resources that belong to this 'WorkloadGroup'.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": {
								Description:         "One or more labels associated with the endpoint.",
								MarkdownDescription: "One or more labels associated with the endpoint.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"locality": {
								Description:         "The locality associated with the endpoint.",
								MarkdownDescription: "The locality associated with the endpoint.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"network": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ports": {
								Description:         "Set of ports associated with the endpoint.",
								MarkdownDescription: "Set of ports associated with the endpoint.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"weight": {
								Description:         "The load balancing weight associated with the endpoint.",
								MarkdownDescription: "The load balancing weight associated with the endpoint.",

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
		},
	}, nil
}

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_workload_group_v1alpha3")

	var state NetworkingIstioIoWorkloadGroupV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoWorkloadGroupV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("WorkloadGroup")

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

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_workload_group_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_workload_group_v1alpha3")

	var state NetworkingIstioIoWorkloadGroupV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoWorkloadGroupV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("WorkloadGroup")

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

func (r *NetworkingIstioIoWorkloadGroupV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_workload_group_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
