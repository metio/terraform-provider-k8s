/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type ChaosMeshOrgPodNetworkChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgPodNetworkChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgPodNetworkChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgPodNetworkChaosV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Ipsets *[]struct {
			CidrAndPorts *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
			} `tfsdk:"cidr_and_ports" yaml:"cidrAndPorts,omitempty"`

			Cidrs *[]string `tfsdk:"cidrs" yaml:"cidrs,omitempty"`

			IpsetType *string `tfsdk:"ipset_type" yaml:"ipsetType,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			SetNames *[]string `tfsdk:"set_names" yaml:"setNames,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`
		} `tfsdk:"ipsets" yaml:"ipsets,omitempty"`

		Iptables *[]struct {
			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

			Ipsets *[]string `tfsdk:"ipsets" yaml:"ipsets,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`
		} `tfsdk:"iptables" yaml:"iptables,omitempty"`

		Tcs *[]struct {
			Bandwidth *struct {
				Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

				Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

				Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

				Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

				Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
			} `tfsdk:"bandwidth" yaml:"bandwidth,omitempty"`

			Corrupt *struct {
				Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

				Corrupt *string `tfsdk:"corrupt" yaml:"corrupt,omitempty"`
			} `tfsdk:"corrupt" yaml:"corrupt,omitempty"`

			Delay *struct {
				Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

				Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

				Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

				Reorder *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Gap *int64 `tfsdk:"gap" yaml:"gap,omitempty"`

					Reorder *string `tfsdk:"reorder" yaml:"reorder,omitempty"`
				} `tfsdk:"reorder" yaml:"reorder,omitempty"`
			} `tfsdk:"delay" yaml:"delay,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Duplicate *struct {
				Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

				Duplicate *string `tfsdk:"duplicate" yaml:"duplicate,omitempty"`
			} `tfsdk:"duplicate" yaml:"duplicate,omitempty"`

			Ipset *string `tfsdk:"ipset" yaml:"ipset,omitempty"`

			Loss *struct {
				Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

				Loss *string `tfsdk:"loss" yaml:"loss,omitempty"`
			} `tfsdk:"loss" yaml:"loss,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"tcs" yaml:"tcs,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgPodNetworkChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgPodNetworkChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_pod_network_chaos_v1alpha1"
}

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "PodNetworkChaos is the Schema for the PodNetworkChaos API",
		MarkdownDescription: "PodNetworkChaos is the Schema for the PodNetworkChaos API",
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
				Description:         "Spec defines the behavior of a pod chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a pod chaos experiment",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"ipsets": {
						Description:         "The ipset on the pod",
						MarkdownDescription: "The ipset on the pod",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cidr_and_ports": {
								Description:         "The contents of ipset. Only available when IPSetType is NetPortIPSet.",
								MarkdownDescription: "The contents of ipset. Only available when IPSetType is NetPortIPSet.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),

											int64validator.AtMost(65535),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cidrs": {
								Description:         "The contents of ipset. Only available when IPSetType is NetIPSet.",
								MarkdownDescription: "The contents of ipset. Only available when IPSetType is NetIPSet.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipset_type": {
								Description:         "IPSetType represents the type of IP set",
								MarkdownDescription: "IPSetType represents the type of IP set",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "The name of ipset",
								MarkdownDescription: "The name of ipset",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"set_names": {
								Description:         "The contents of ipset. Only available when IPSetType is SetIPSet.",
								MarkdownDescription: "The contents of ipset. Only available when IPSetType is SetIPSet.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables": {
						Description:         "The iptables rules on the pod",
						MarkdownDescription: "The iptables rules on the pod",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"device": {
								Description:         "Device represents the network device to be affected.",
								MarkdownDescription: "Device represents the network device to be affected.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"direction": {
								Description:         "The block direction of this iptables rule",
								MarkdownDescription: "The block direction of this iptables rule",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ipsets": {
								Description:         "The name of related ipset",
								MarkdownDescription: "The name of related ipset",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "The name of iptables chain",
								MarkdownDescription: "The name of iptables chain",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"source": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tcs": {
						Description:         "The tc rules on the pod",
						MarkdownDescription: "The tc rules on the pod",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"bandwidth": {
								Description:         "Bandwidth represents the detail about bandwidth control action",
								MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"buffer": {
										Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
										MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"limit": {
										Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
										MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"minburst": {
										Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
										MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"peakrate": {
										Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
										MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"rate": {
										Description:         "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
										MarkdownDescription: "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"corrupt": {
								Description:         "Corrupt represents the detail about corrupt action",
								MarkdownDescription: "Corrupt represents the detail about corrupt action",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"corrupt": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"delay": {
								Description:         "Delay represents the detail about delay action",
								MarkdownDescription: "Delay represents the detail about delay action",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jitter": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"latency": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"reorder": {
										Description:         "ReorderSpec defines details of packet reorder.",
										MarkdownDescription: "ReorderSpec defines details of packet reorder.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"gap": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"reorder": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
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

							"device": {
								Description:         "Device represents the network device to be affected.",
								MarkdownDescription: "Device represents the network device to be affected.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplicate": {
								Description:         "DuplicateSpec represents the detail about loss action",
								MarkdownDescription: "DuplicateSpec represents the detail about loss action",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duplicate": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipset": {
								Description:         "The name of target ipset",
								MarkdownDescription: "The name of target ipset",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss": {
								Description:         "Loss represents the detail about loss action",
								MarkdownDescription: "Loss represents the detail about loss action",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"loss": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source": {
								Description:         "The name and namespace of the source network chaos",
								MarkdownDescription: "The name and namespace of the source network chaos",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"type": {
								Description:         "The type of traffic control",
								MarkdownDescription: "The type of traffic control",

								Type: types.StringType,

								Required: true,
								Optional: false,
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
		},
	}, nil
}

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_pod_network_chaos_v1alpha1")

	var state ChaosMeshOrgPodNetworkChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPodNetworkChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PodNetworkChaos")

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

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_pod_network_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_pod_network_chaos_v1alpha1")

	var state ChaosMeshOrgPodNetworkChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPodNetworkChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PodNetworkChaos")

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

func (r *ChaosMeshOrgPodNetworkChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_pod_network_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
