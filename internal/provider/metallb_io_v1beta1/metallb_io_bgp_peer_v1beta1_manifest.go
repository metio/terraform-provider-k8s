/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metallb_io_v1beta1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MetallbIoBgppeerV1Beta1Manifest{}
)

func NewMetallbIoBgppeerV1Beta1Manifest() datasource.DataSource {
	return &MetallbIoBgppeerV1Beta1Manifest{}
}

type MetallbIoBgppeerV1Beta1Manifest struct{}

type MetallbIoBgppeerV1Beta1ManifestData struct {
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
		BfdProfile    *string `tfsdk:"bfd_profile" json:"bfdProfile,omitempty"`
		EbgpMultiHop  *bool   `tfsdk:"ebgp_multi_hop" json:"ebgpMultiHop,omitempty"`
		HoldTime      *string `tfsdk:"hold_time" json:"holdTime,omitempty"`
		KeepaliveTime *string `tfsdk:"keepalive_time" json:"keepaliveTime,omitempty"`
		MyASN         *int64  `tfsdk:"my_asn" json:"myASN,omitempty"`
		NodeSelectors *[]struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
		Password      *string `tfsdk:"password" json:"password,omitempty"`
		PeerASN       *int64  `tfsdk:"peer_asn" json:"peerASN,omitempty"`
		PeerAddress   *string `tfsdk:"peer_address" json:"peerAddress,omitempty"`
		PeerPort      *int64  `tfsdk:"peer_port" json:"peerPort,omitempty"`
		RouterID      *string `tfsdk:"router_id" json:"routerID,omitempty"`
		SourceAddress *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MetallbIoBgppeerV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metallb_io_bgp_peer_v1beta1_manifest"
}

func (r *MetallbIoBgppeerV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BGPPeer is the Schema for the peers API.",
		MarkdownDescription: "BGPPeer is the Schema for the peers API.",
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
				Description:         "BGPPeerSpec defines the desired state of Peer.",
				MarkdownDescription: "BGPPeerSpec defines the desired state of Peer.",
				Attributes: map[string]schema.Attribute{
					"bfd_profile": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ebgp_multi_hop": schema.BoolAttribute{
						Description:         "EBGP peer is multi-hops away",
						MarkdownDescription: "EBGP peer is multi-hops away",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hold_time": schema.StringAttribute{
						Description:         "Requested BGP hold time, per RFC4271.",
						MarkdownDescription: "Requested BGP hold time, per RFC4271.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"keepalive_time": schema.StringAttribute{
						Description:         "Requested BGP keepalive time, per RFC4271.",
						MarkdownDescription: "Requested BGP keepalive time, per RFC4271.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"my_asn": schema.Int64Attribute{
						Description:         "AS number to use for the local end of the session.",
						MarkdownDescription: "AS number to use for the local end of the session.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4.294967295e+09),
						},
					},

					"node_selectors": schema.ListNestedAttribute{
						Description:         "Only connect to this peer on nodes that match one of these selectors.",
						MarkdownDescription: "Only connect to this peer on nodes that match one of these selectors.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match_expressions": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operator": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"values": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"match_labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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

					"password": schema.StringAttribute{
						Description:         "Authentication password for routers enforcing TCP MD5 authenticated sessions",
						MarkdownDescription: "Authentication password for routers enforcing TCP MD5 authenticated sessions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"peer_asn": schema.Int64Attribute{
						Description:         "AS number to expect from the remote end of the session.",
						MarkdownDescription: "AS number to expect from the remote end of the session.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4.294967295e+09),
						},
					},

					"peer_address": schema.StringAttribute{
						Description:         "Address to dial when establishing the session.",
						MarkdownDescription: "Address to dial when establishing the session.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"peer_port": schema.Int64Attribute{
						Description:         "Port to dial when establishing the session.",
						MarkdownDescription: "Port to dial when establishing the session.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(16384),
						},
					},

					"router_id": schema.StringAttribute{
						Description:         "BGP router ID to advertise to the peer",
						MarkdownDescription: "BGP router ID to advertise to the peer",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_address": schema.StringAttribute{
						Description:         "Source address to use when establishing the session.",
						MarkdownDescription: "Source address to use when establishing the session.",
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
	}
}

func (r *MetallbIoBgppeerV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_metallb_io_bgp_peer_v1beta1_manifest")

	var model MetallbIoBgppeerV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("metallb.io/v1beta1")
	model.Kind = pointer.String("BGPPeer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
