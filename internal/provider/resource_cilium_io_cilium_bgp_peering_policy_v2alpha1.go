/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource struct{}

var (
	_ resource.Resource = (*CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource)(nil)
)

type CiliumIoCiliumBGPPeeringPolicyV2Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CiliumIoCiliumBGPPeeringPolicyV2Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		VirtualRouters *[]struct {
			ExportPodCIDR *bool `tfsdk:"export_pod_cidr" yaml:"exportPodCIDR,omitempty"`

			LocalASN *int64 `tfsdk:"local_asn" yaml:"localASN,omitempty"`

			Neighbors *[]struct {
				PeerASN *int64 `tfsdk:"peer_asn" yaml:"peerASN,omitempty"`

				PeerAddress *string `tfsdk:"peer_address" yaml:"peerAddress,omitempty"`
			} `tfsdk:"neighbors" yaml:"neighbors,omitempty"`
		} `tfsdk:"virtual_routers" yaml:"virtualRouters,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource() resource.Resource {
	return &CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource{}
}

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cilium_io_cilium_bgp_peering_policy_v2alpha1"
}

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CiliumBGPPeeringPolicy is a Kubernetes third-party resource for instructing Cilium's BGP control plane to create virtual BGP routers.",
		MarkdownDescription: "CiliumBGPPeeringPolicy is a Kubernetes third-party resource for instructing Cilium's BGP control plane to create virtual BGP routers.",
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

			"spec": {
				Description:         "Spec is a human readable description of a BGP peering policy",
				MarkdownDescription: "Spec is a human readable description of a BGP peering policy",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"node_selector": {
						Description:         "NodeSelector selects a group of nodes where this BGP Peering Policy applies.  If nil this policy applies to all nodes.",
						MarkdownDescription: "NodeSelector selects a group of nodes where this BGP Peering Policy applies.  If nil this policy applies to all nodes.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
										},
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

					"virtual_routers": {
						Description:         "A list of CiliumBGPVirtualRouter(s) which instructs the BGP control plane how to instantiate virtual BGP routers.",
						MarkdownDescription: "A list of CiliumBGPVirtualRouter(s) which instructs the BGP control plane how to instantiate virtual BGP routers.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"export_pod_cidr": {
								Description:         "ExportPodCIDR determines whether to export the Node's private CIDR block to the configured neighbors.",
								MarkdownDescription: "ExportPodCIDR determines whether to export the Node's private CIDR block to the configured neighbors.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"local_asn": {
								Description:         "LocalASN is the ASN of this virtual router. Supports extended 32bit ASNs",
								MarkdownDescription: "LocalASN is the ASN of this virtual router. Supports extended 32bit ASNs",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"neighbors": {
								Description:         "Neighbors is a list of neighboring BGP peers for this virtual router",
								MarkdownDescription: "Neighbors is a list of neighboring BGP peers for this virtual router",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"peer_asn": {
										Description:         "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs",
										MarkdownDescription: "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),

											int64validator.AtMost(4.294967295e+09),
										},
									},

									"peer_address": {
										Description:         "PeerAddress is the IP address of the peer. This must be in CIDR notation and use a /32 to express a single host.",
										MarkdownDescription: "PeerAddress is the IP address of the peer. This must be in CIDR notation and use a /32 to express a single host.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
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

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var state CiliumIoCiliumBGPPeeringPolicyV2Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumBGPPeeringPolicyV2Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2alpha1")
	goModel.Kind = utilities.Ptr("CiliumBGPPeeringPolicy")

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

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var state CiliumIoCiliumBGPPeeringPolicyV2Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumBGPPeeringPolicyV2Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2alpha1")
	goModel.Kind = utilities.Ptr("CiliumBGPPeeringPolicy")

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

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
