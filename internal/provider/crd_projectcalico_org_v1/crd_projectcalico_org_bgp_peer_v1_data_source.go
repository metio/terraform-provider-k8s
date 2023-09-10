/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource              = &CrdProjectcalicoOrgBgppeerV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CrdProjectcalicoOrgBgppeerV1DataSource{}
)

func NewCrdProjectcalicoOrgBgppeerV1DataSource() datasource.DataSource {
	return &CrdProjectcalicoOrgBgppeerV1DataSource{}
}

type CrdProjectcalicoOrgBgppeerV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CrdProjectcalicoOrgBgppeerV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AsNumber                 *int64    `tfsdk:"as_number" json:"asNumber,omitempty"`
		Filters                  *[]string `tfsdk:"filters" json:"filters,omitempty"`
		KeepOriginalNextHop      *bool     `tfsdk:"keep_original_next_hop" json:"keepOriginalNextHop,omitempty"`
		MaxRestartTime           *string   `tfsdk:"max_restart_time" json:"maxRestartTime,omitempty"`
		Node                     *string   `tfsdk:"node" json:"node,omitempty"`
		NodeSelector             *string   `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		NumAllowedLocalASNumbers *int64    `tfsdk:"num_allowed_local_as_numbers" json:"numAllowedLocalASNumbers,omitempty"`
		Password                 *struct {
			SecretKeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
		} `tfsdk:"password" json:"password,omitempty"`
		PeerIP        *string `tfsdk:"peer_ip" json:"peerIP,omitempty"`
		PeerSelector  *string `tfsdk:"peer_selector" json:"peerSelector,omitempty"`
		ReachableBy   *string `tfsdk:"reachable_by" json:"reachableBy,omitempty"`
		SourceAddress *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
		TtlSecurity   *int64  `tfsdk:"ttl_security" json:"ttlSecurity,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgBgppeerV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_bgp_peer_v1"
}

func (r *CrdProjectcalicoOrgBgppeerV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "BGPPeerSpec contains the specification for a BGPPeer resource.",
				MarkdownDescription: "BGPPeerSpec contains the specification for a BGPPeer resource.",
				Attributes: map[string]schema.Attribute{
					"as_number": schema.Int64Attribute{
						Description:         "The AS Number of the peer.",
						MarkdownDescription: "The AS Number of the peer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"filters": schema.ListAttribute{
						Description:         "The ordered set of BGPFilters applied on this BGP peer.",
						MarkdownDescription: "The ordered set of BGPFilters applied on this BGP peer.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"keep_original_next_hop": schema.BoolAttribute{
						Description:         "Option to keep the original nexthop field when routes are sent to a BGP Peer. Setting 'true' configures the selected BGP Peers node to use the 'next hop keep;' instead of 'next hop self;'(default) in the specific branch of the Node on 'bird.cfg'.",
						MarkdownDescription: "Option to keep the original nexthop field when routes are sent to a BGP Peer. Setting 'true' configures the selected BGP Peers node to use the 'next hop keep;' instead of 'next hop self;'(default) in the specific branch of the Node on 'bird.cfg'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_restart_time": schema.StringAttribute{
						Description:         "Time to allow for software restart.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used.",
						MarkdownDescription: "Time to allow for software restart.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node": schema.StringAttribute{
						Description:         "The node name identifying the Calico node instance that is targeted by this peer. If this is not set, and no nodeSelector is specified, then this BGP peer selects all nodes in the cluster.",
						MarkdownDescription: "The node name identifying the Calico node instance that is targeted by this peer. If this is not set, and no nodeSelector is specified, then this BGP peer selects all nodes in the cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_selector": schema.StringAttribute{
						Description:         "Selector for the nodes that should have this peering.  When this is set, the Node field must be empty.",
						MarkdownDescription: "Selector for the nodes that should have this peering.  When this is set, the Node field must be empty.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"num_allowed_local_as_numbers": schema.Int64Attribute{
						Description:         "Maximum number of local AS numbers that are allowed in the AS path for received routes. This removes BGP loop prevention and should only be used if absolutely necesssary.",
						MarkdownDescription: "Maximum number of local AS numbers that are allowed in the AS path for received routes. This removes BGP loop prevention and should only be used if absolutely necesssary.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"password": schema.SingleNestedAttribute{
						Description:         "Optional BGP password for the peerings generated by this BGPPeer resource.",
						MarkdownDescription: "Optional BGP password for the peerings generated by this BGPPeer resource.",
						Attributes: map[string]schema.Attribute{
							"secret_key_ref": schema.SingleNestedAttribute{
								Description:         "Selects a key of a secret in the node pod's namespace.",
								MarkdownDescription: "Selects a key of a secret in the node pod's namespace.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"optional": schema.BoolAttribute{
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

					"peer_ip": schema.StringAttribute{
						Description:         "The IP address of the peer followed by an optional port number to peer with. If port number is given, format should be '[<IPv6>]:port' or '<IPv4>:<port>' for IPv4. If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node with ListenPort set in BGPConfiguration, then we use that port to peer.",
						MarkdownDescription: "The IP address of the peer followed by an optional port number to peer with. If port number is given, format should be '[<IPv6>]:port' or '<IPv4>:<port>' for IPv4. If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node with ListenPort set in BGPConfiguration, then we use that port to peer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"peer_selector": schema.StringAttribute{
						Description:         "Selector for the remote nodes to peer with.  When this is set, the PeerIP and ASNumber fields must be empty.  For each peering between the local node and selected remote nodes, we configure an IPv4 peering if both ends have NodeBGPSpec.IPv4Address specified, and an IPv6 peering if both ends have NodeBGPSpec.IPv6Address specified.  The remote AS number comes from the remote node's NodeBGPSpec.ASNumber, or the global default if that is not set.",
						MarkdownDescription: "Selector for the remote nodes to peer with.  When this is set, the PeerIP and ASNumber fields must be empty.  For each peering between the local node and selected remote nodes, we configure an IPv4 peering if both ends have NodeBGPSpec.IPv4Address specified, and an IPv6 peering if both ends have NodeBGPSpec.IPv6Address specified.  The remote AS number comes from the remote node's NodeBGPSpec.ASNumber, or the global default if that is not set.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"reachable_by": schema.StringAttribute{
						Description:         "Add an exact, i.e. /32, static route toward peer IP in order to prevent route flapping. ReachableBy contains the address of the gateway which peer can be reached by.",
						MarkdownDescription: "Add an exact, i.e. /32, static route toward peer IP in order to prevent route flapping. ReachableBy contains the address of the gateway which peer can be reached by.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"source_address": schema.StringAttribute{
						Description:         "Specifies whether and how to configure a source address for the peerings generated by this BGPPeer resource.  Default value 'UseNodeIP' means to configure the node IP as the source address.  'None' means not to configure a source address.",
						MarkdownDescription: "Specifies whether and how to configure a source address for the peerings generated by this BGPPeer resource.  Default value 'UseNodeIP' means to configure the node IP as the source address.  'None' means not to configure a source address.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ttl_security": schema.Int64Attribute{
						Description:         "TTLSecurity enables the generalized TTL security mechanism (GTSM) which protects against spoofed packets by ignoring received packets with a smaller than expected TTL value. The provided value is the number of hops (edges) between the peers.",
						MarkdownDescription: "TTLSecurity enables the generalized TTL security mechanism (GTSM) which protects against spoofed packets by ignoring received packets with a smaller than expected TTL value. The provided value is the number of hops (edges) between the peers.",
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
	}
}

func (r *CrdProjectcalicoOrgBgppeerV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CrdProjectcalicoOrgBgppeerV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_crd_projectcalico_org_bgp_peer_v1")

	var data CrdProjectcalicoOrgBgppeerV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "BGPPeer"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CrdProjectcalicoOrgBgppeerV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	data.Kind = pointer.String("BGPPeer")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
