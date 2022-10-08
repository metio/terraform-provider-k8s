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

type NetworkingIstioIoServiceEntryV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*NetworkingIstioIoServiceEntryV1Alpha3Resource)(nil)
)

type NetworkingIstioIoServiceEntryV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingIstioIoServiceEntryV1Alpha3GoModel struct {
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
		Addresses *[]string `tfsdk:"addresses" yaml:"addresses,omitempty"`

		ExportTo *[]string `tfsdk:"export_to" yaml:"exportTo,omitempty"`

		Ports *[]struct {
			TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Number *int64 `tfsdk:"number" yaml:"number,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
		} `tfsdk:"ports" yaml:"ports,omitempty"`

		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
		} `tfsdk:"workload_selector" yaml:"workloadSelector,omitempty"`

		Endpoints *[]struct {
			Address *string `tfsdk:"address" yaml:"address,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			Locality *string `tfsdk:"locality" yaml:"locality,omitempty"`

			Network *string `tfsdk:"network" yaml:"network,omitempty"`

			Ports *map[string]string `tfsdk:"ports" yaml:"ports,omitempty"`

			ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

			Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
		} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

		Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

		Location *string `tfsdk:"location" yaml:"location,omitempty"`

		Resolution *string `tfsdk:"resolution" yaml:"resolution,omitempty"`

		SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingIstioIoServiceEntryV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoServiceEntryV1Alpha3Resource{}
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_istio_io_service_entry_v1alpha3"
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "Configuration affecting service registry. See more details at: https://istio.io/docs/reference/config/networking/service-entry.html",
				MarkdownDescription: "Configuration affecting service registry. See more details at: https://istio.io/docs/reference/config/networking/service-entry.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"addresses": {
						Description:         "The virtual IP addresses associated with the service.",
						MarkdownDescription: "The virtual IP addresses associated with the service.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"export_to": {
						Description:         "A list of namespaces to which this service is exported.",
						MarkdownDescription: "A list of namespaces to which this service is exported.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ports": {
						Description:         "The ports associated with the external service.",
						MarkdownDescription: "The ports associated with the external service.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"target_port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Label assigned to the port.",
								MarkdownDescription: "Label assigned to the port.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"number": {
								Description:         "A valid non-negative integer port number.",
								MarkdownDescription: "A valid non-negative integer port number.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"protocol": {
								Description:         "The protocol exposed on the port.",
								MarkdownDescription: "The protocol exposed on the port.",

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

					"workload_selector": {
						Description:         "Applicable only for MESH_INTERNAL services.",
						MarkdownDescription: "Applicable only for MESH_INTERNAL services.",

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

					"endpoints": {
						Description:         "One or more endpoints associated with the service.",
						MarkdownDescription: "One or more endpoints associated with the service.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

					"hosts": {
						Description:         "The hosts associated with the ServiceEntry.",
						MarkdownDescription: "The hosts associated with the ServiceEntry.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"location": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resolution": {
						Description:         "Service discovery mode for the hosts.",
						MarkdownDescription: "Service discovery mode for the hosts.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subject_alt_names": {
						Description:         "",
						MarkdownDescription: "",

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
		},
	}, nil
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_service_entry_v1alpha3")

	var state NetworkingIstioIoServiceEntryV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoServiceEntryV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("ServiceEntry")

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

func (r *NetworkingIstioIoServiceEntryV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_service_entry_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingIstioIoServiceEntryV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_service_entry_v1alpha3")

	var state NetworkingIstioIoServiceEntryV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoServiceEntryV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("ServiceEntry")

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

func (r *NetworkingIstioIoServiceEntryV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_service_entry_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
