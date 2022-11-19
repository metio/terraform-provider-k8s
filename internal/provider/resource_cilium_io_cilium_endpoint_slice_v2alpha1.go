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

type CiliumIoCiliumEndpointSliceV2Alpha1Resource struct{}

var (
	_ resource.Resource = (*CiliumIoCiliumEndpointSliceV2Alpha1Resource)(nil)
)

type CiliumIoCiliumEndpointSliceV2Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Endpoints  types.List   `tfsdk:"endpoints"`
	Namespace  types.String `tfsdk:"namespace"`
}

type CiliumIoCiliumEndpointSliceV2Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Endpoints *[]struct {
		Encryption *struct {
			Key *int64 `tfsdk:"key" yaml:"key,omitempty"`
		} `tfsdk:"encryption" yaml:"encryption,omitempty"`

		Id *int64 `tfsdk:"id" yaml:"id,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Named_ports *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
		} `tfsdk:"named_ports" yaml:"named-ports,omitempty"`

		Networking *struct {
			Addressing *[]struct {
				Ipv4 *string `tfsdk:"ipv4" yaml:"ipv4,omitempty"`

				Ipv6 *string `tfsdk:"ipv6" yaml:"ipv6,omitempty"`
			} `tfsdk:"addressing" yaml:"addressing,omitempty"`

			Node *string `tfsdk:"node" yaml:"node,omitempty"`
		} `tfsdk:"networking" yaml:"networking,omitempty"`
	} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

	Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
}

func NewCiliumIoCiliumEndpointSliceV2Alpha1Resource() resource.Resource {
	return &CiliumIoCiliumEndpointSliceV2Alpha1Resource{}
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cilium_io_cilium_endpoint_slice_v2alpha1"
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CiliumEndpointSlice contains a group of CoreCiliumendpoints.",
		MarkdownDescription: "CiliumEndpointSlice contains a group of CoreCiliumendpoints.",
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

			"endpoints": {
				Description:         "Endpoints is a list of coreCEPs packed in a CiliumEndpointSlice",
				MarkdownDescription: "Endpoints is a list of coreCEPs packed in a CiliumEndpointSlice",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"encryption": {
						Description:         "EncryptionSpec defines the encryption relevant configuration of a node.",
						MarkdownDescription: "EncryptionSpec defines the encryption relevant configuration of a node.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
								MarkdownDescription: "Key is the index to the key to use for encryption or 0 if encryption is disabled.",

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

					"id": {
						Description:         "IdentityID is the numeric identity of the endpoint",
						MarkdownDescription: "IdentityID is the numeric identity of the endpoint",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "Name indicate as CiliumEndpoint name.",
						MarkdownDescription: "Name indicate as CiliumEndpoint name.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"named_ports": {
						Description:         "NamedPorts List of named Layer 4 port and protocol pairs which will be used in Network Policy specs.  swagger:model NamedPorts",
						MarkdownDescription: "NamedPorts List of named Layer 4 port and protocol pairs which will be used in Network Policy specs.  swagger:model NamedPorts",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Optional layer 4 port name",
								MarkdownDescription: "Optional layer 4 port name",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "Layer 4 port number",
								MarkdownDescription: "Layer 4 port number",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"protocol": {
								Description:         "Layer 4 protocol Enum: [TCP UDP SCTP ICMP ICMPV6 ANY]",
								MarkdownDescription: "Layer 4 protocol Enum: [TCP UDP SCTP ICMP ICMPV6 ANY]",

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

					"networking": {
						Description:         "EndpointNetworking is the addressing information of an endpoint.",
						MarkdownDescription: "EndpointNetworking is the addressing information of an endpoint.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addressing": {
								Description:         "IP4/6 addresses assigned to this Endpoint",
								MarkdownDescription: "IP4/6 addresses assigned to this Endpoint",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ipv4": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ipv6": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"node": {
								Description:         "NodeIP is the IP of the node the endpoint is running on. The IP must be reachable between nodes.",
								MarkdownDescription: "NodeIP is the IP of the node the endpoint is running on. The IP must be reachable between nodes.",

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

				Required: true,
				Optional: false,
				Computed: false,
			},

			"namespace": {
				Description:         "Namespace indicate as CiliumEndpointSlice namespace. All the CiliumEndpoints within the same namespace are put together in CiliumEndpointSlice.",
				MarkdownDescription: "Namespace indicate as CiliumEndpointSlice namespace. All the CiliumEndpoints within the same namespace are put together in CiliumEndpointSlice.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cilium_io_cilium_endpoint_slice_v2alpha1")

	var state CiliumIoCiliumEndpointSliceV2Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumEndpointSliceV2Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2alpha1")
	goModel.Kind = utilities.Ptr("CiliumEndpointSlice")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_endpoint_slice_v2alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cilium_io_cilium_endpoint_slice_v2alpha1")

	var state CiliumIoCiliumEndpointSliceV2Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumEndpointSliceV2Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2alpha1")
	goModel.Kind = utilities.Ptr("CiliumEndpointSlice")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cilium_io_cilium_endpoint_slice_v2alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
