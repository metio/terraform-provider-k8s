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

type CrdProjectcalicoOrgBGPPeerV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgBGPPeerV1Resource)(nil)
)

type CrdProjectcalicoOrgBGPPeerV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgBGPPeerV1GoModel struct {
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
		AsNumber *int64 `tfsdk:"as_number" yaml:"asNumber,omitempty"`

		KeepOriginalNextHop *bool `tfsdk:"keep_original_next_hop" yaml:"keepOriginalNextHop,omitempty"`

		MaxRestartTime *string `tfsdk:"max_restart_time" yaml:"maxRestartTime,omitempty"`

		Node *string `tfsdk:"node" yaml:"node,omitempty"`

		NodeSelector *string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		NumAllowedLocalASNumbers *int64 `tfsdk:"num_allowed_local_as_numbers" yaml:"numAllowedLocalASNumbers,omitempty"`

		Password *struct {
			SecretKeyRef *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
		} `tfsdk:"password" yaml:"password,omitempty"`

		PeerIP *string `tfsdk:"peer_ip" yaml:"peerIP,omitempty"`

		PeerSelector *string `tfsdk:"peer_selector" yaml:"peerSelector,omitempty"`

		SourceAddress *string `tfsdk:"source_address" yaml:"sourceAddress,omitempty"`

		TtlSecurity *int64 `tfsdk:"ttl_security" yaml:"ttlSecurity,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgBGPPeerV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgBGPPeerV1Resource{}
}

func (r *CrdProjectcalicoOrgBGPPeerV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_bgp_peer_v1"
}

func (r *CrdProjectcalicoOrgBGPPeerV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "BGPPeerSpec contains the specification for a BGPPeer resource.",
				MarkdownDescription: "BGPPeerSpec contains the specification for a BGPPeer resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"as_number": {
						Description:         "The AS Number of the peer.",
						MarkdownDescription: "The AS Number of the peer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"keep_original_next_hop": {
						Description:         "Option to keep the original nexthop field when routes are sent to a BGP Peer. Setting 'true' configures the selected BGP Peers node to use the 'next hop keep;' instead of 'next hop self;'(default) in the specific branch of the Node on 'bird.cfg'.",
						MarkdownDescription: "Option to keep the original nexthop field when routes are sent to a BGP Peer. Setting 'true' configures the selected BGP Peers node to use the 'next hop keep;' instead of 'next hop self;'(default) in the specific branch of the Node on 'bird.cfg'.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_restart_time": {
						Description:         "Time to allow for software restart.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used.",
						MarkdownDescription: "Time to allow for software restart.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node": {
						Description:         "The node name identifying the Calico node instance that is targeted by this peer. If this is not set, and no nodeSelector is specified, then this BGP peer selects all nodes in the cluster.",
						MarkdownDescription: "The node name identifying the Calico node instance that is targeted by this peer. If this is not set, and no nodeSelector is specified, then this BGP peer selects all nodes in the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "Selector for the nodes that should have this peering.  When this is set, the Node field must be empty.",
						MarkdownDescription: "Selector for the nodes that should have this peering.  When this is set, the Node field must be empty.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"num_allowed_local_as_numbers": {
						Description:         "Maximum number of local AS numbers that are allowed in the AS path for received routes. This removes BGP loop prevention and should only be used if absolutely necesssary.",
						MarkdownDescription: "Maximum number of local AS numbers that are allowed in the AS path for received routes. This removes BGP loop prevention and should only be used if absolutely necesssary.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"password": {
						Description:         "Optional BGP password for the peerings generated by this BGPPeer resource.",
						MarkdownDescription: "Optional BGP password for the peerings generated by this BGPPeer resource.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret_key_ref": {
								Description:         "Selects a key of a secret in the node pod's namespace.",
								MarkdownDescription: "Selects a key of a secret in the node pod's namespace.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

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

					"peer_ip": {
						Description:         "The IP address of the peer followed by an optional port number to peer with. If port number is given, format should be '[<IPv6>]:port' or '<IPv4>:<port>' for IPv4. If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node with ListenPort set in BGPConfiguration, then we use that port to peer.",
						MarkdownDescription: "The IP address of the peer followed by an optional port number to peer with. If port number is given, format should be '[<IPv6>]:port' or '<IPv4>:<port>' for IPv4. If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node with ListenPort set in BGPConfiguration, then we use that port to peer.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"peer_selector": {
						Description:         "Selector for the remote nodes to peer with.  When this is set, the PeerIP and ASNumber fields must be empty.  For each peering between the local node and selected remote nodes, we configure an IPv4 peering if both ends have NodeBGPSpec.IPv4Address specified, and an IPv6 peering if both ends have NodeBGPSpec.IPv6Address specified.  The remote AS number comes from the remote node's NodeBGPSpec.ASNumber, or the global default if that is not set.",
						MarkdownDescription: "Selector for the remote nodes to peer with.  When this is set, the PeerIP and ASNumber fields must be empty.  For each peering between the local node and selected remote nodes, we configure an IPv4 peering if both ends have NodeBGPSpec.IPv4Address specified, and an IPv6 peering if both ends have NodeBGPSpec.IPv6Address specified.  The remote AS number comes from the remote node's NodeBGPSpec.ASNumber, or the global default if that is not set.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_address": {
						Description:         "Specifies whether and how to configure a source address for the peerings generated by this BGPPeer resource.  Default value 'UseNodeIP' means to configure the node IP as the source address.  'None' means not to configure a source address.",
						MarkdownDescription: "Specifies whether and how to configure a source address for the peerings generated by this BGPPeer resource.  Default value 'UseNodeIP' means to configure the node IP as the source address.  'None' means not to configure a source address.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ttl_security": {
						Description:         "TTLSecurity enables the generalized TTL security mechanism (GTSM) which protects against spoofed packets by ignoring received packets with a smaller than expected TTL value. The provided value is the number of hops (edges) between the peers.",
						MarkdownDescription: "TTLSecurity enables the generalized TTL security mechanism (GTSM) which protects against spoofed packets by ignoring received packets with a smaller than expected TTL value. The provided value is the number of hops (edges) between the peers.",

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
		},
	}, nil
}

func (r *CrdProjectcalicoOrgBGPPeerV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_bgp_peer_v1")

	var state CrdProjectcalicoOrgBGPPeerV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgBGPPeerV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("BGPPeer")

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

func (r *CrdProjectcalicoOrgBGPPeerV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_bgp_peer_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgBGPPeerV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_bgp_peer_v1")

	var state CrdProjectcalicoOrgBGPPeerV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgBGPPeerV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("BGPPeer")

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

func (r *CrdProjectcalicoOrgBGPPeerV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_bgp_peer_v1")
	// NO-OP: Terraform removes the state automatically for us
}
