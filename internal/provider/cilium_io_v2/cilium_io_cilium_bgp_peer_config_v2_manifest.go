/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

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
	_ datasource.DataSource = &CiliumIoCiliumBgppeerConfigV2Manifest{}
)

func NewCiliumIoCiliumBgppeerConfigV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumBgppeerConfigV2Manifest{}
}

type CiliumIoCiliumBgppeerConfigV2Manifest struct{}

type CiliumIoCiliumBgppeerConfigV2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AuthSecretRef *string `tfsdk:"auth_secret_ref" json:"authSecretRef,omitempty"`
		EbgpMultihop  *int64  `tfsdk:"ebgp_multihop" json:"ebgpMultihop,omitempty"`
		Families      *[]struct {
			Advertisements *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"advertisements" json:"advertisements,omitempty"`
			Afi  *string `tfsdk:"afi" json:"afi,omitempty"`
			Safi *string `tfsdk:"safi" json:"safi,omitempty"`
		} `tfsdk:"families" json:"families,omitempty"`
		GracefulRestart *struct {
			Enabled            *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
			RestartTimeSeconds *int64 `tfsdk:"restart_time_seconds" json:"restartTimeSeconds,omitempty"`
		} `tfsdk:"graceful_restart" json:"gracefulRestart,omitempty"`
		Timers *struct {
			ConnectRetryTimeSeconds *int64 `tfsdk:"connect_retry_time_seconds" json:"connectRetryTimeSeconds,omitempty"`
			HoldTimeSeconds         *int64 `tfsdk:"hold_time_seconds" json:"holdTimeSeconds,omitempty"`
			KeepAliveTimeSeconds    *int64 `tfsdk:"keep_alive_time_seconds" json:"keepAliveTimeSeconds,omitempty"`
		} `tfsdk:"timers" json:"timers,omitempty"`
		Transport *struct {
			PeerPort *int64 `tfsdk:"peer_port" json:"peerPort,omitempty"`
		} `tfsdk:"transport" json:"transport,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumBgppeerConfigV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_peer_config_v2_manifest"
}

func (r *CiliumIoCiliumBgppeerConfigV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec is the specification of the desired behavior of the CiliumBGPPeerConfig.",
				MarkdownDescription: "Spec is the specification of the desired behavior of the CiliumBGPPeerConfig.",
				Attributes: map[string]schema.Attribute{
					"auth_secret_ref": schema.StringAttribute{
						Description:         "AuthSecretRef is the name of the secret to use to fetch a TCP authentication password for this peer. If not specified, no authentication is used.",
						MarkdownDescription: "AuthSecretRef is the name of the secret to use to fetch a TCP authentication password for this peer. If not specified, no authentication is used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ebgp_multihop": schema.Int64Attribute{
						Description:         "EBGPMultihopTTL controls the multi-hop feature for eBGP peers. Its value defines the Time To Live (TTL) value used in BGP packets sent to the peer. If not specified, EBGP multihop is disabled. This field is ignored for iBGP neighbors.",
						MarkdownDescription: "EBGPMultihopTTL controls the multi-hop feature for eBGP peers. Its value defines the Time To Live (TTL) value used in BGP packets sent to the peer. If not specified, EBGP multihop is disabled. This field is ignored for iBGP neighbors.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(255),
						},
					},

					"families": schema.ListNestedAttribute{
						Description:         "Families, if provided, defines a set of AFI/SAFIs the speaker will negotiate with it's peer. If not specified, the default families of IPv6/unicast and IPv4/unicast will be created.",
						MarkdownDescription: "Families, if provided, defines a set of AFI/SAFIs the speaker will negotiate with it's peer. If not specified, the default families of IPv6/unicast and IPv4/unicast will be created.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"advertisements": schema.SingleNestedAttribute{
									Description:         "Advertisements selects group of BGP Advertisement(s) to advertise for this family. If not specified, no advertisements are sent for this family. This field is ignored in CiliumBGPNeighbor which is used in CiliumBGPPeeringPolicy. Use CiliumBGPPeeringPolicy advertisement options instead.",
									MarkdownDescription: "Advertisements selects group of BGP Advertisement(s) to advertise for this family. If not specified, no advertisements are sent for this family. This field is ignored in CiliumBGPNeighbor which is used in CiliumBGPPeeringPolicy. Use CiliumBGPPeeringPolicy advertisement options instead.",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
														},
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"afi": schema.StringAttribute{
									Description:         "Afi is the Address Family Identifier (AFI) of the family.",
									MarkdownDescription: "Afi is the Address Family Identifier (AFI) of the family.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("ipv4", "ipv6", "l2vpn", "ls", "opaque"),
									},
								},

								"safi": schema.StringAttribute{
									Description:         "Safi is the Subsequent Address Family Identifier (SAFI) of the family.",
									MarkdownDescription: "Safi is the Subsequent Address Family Identifier (SAFI) of the family.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("unicast", "multicast", "mpls_label", "encapsulation", "vpls", "evpn", "ls", "sr_policy", "mup", "mpls_vpn", "mpls_vpn_multicast", "route_target_constraints", "flowspec_unicast", "flowspec_vpn", "key_value"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"graceful_restart": schema.SingleNestedAttribute{
						Description:         "GracefulRestart defines graceful restart parameters which are negotiated with this peer. If not specified, the graceful restart capability is disabled.",
						MarkdownDescription: "GracefulRestart defines graceful restart parameters which are negotiated with this peer. If not specified, the graceful restart capability is disabled.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled flag, when set enables graceful restart capability.",
								MarkdownDescription: "Enabled flag, when set enables graceful restart capability.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"restart_time_seconds": schema.Int64Attribute{
								Description:         "RestartTimeSeconds is the estimated time it will take for the BGP session to be re-established with peer after a restart. After this period, peer will remove stale routes. This is described RFC 4724 section 4.2.",
								MarkdownDescription: "RestartTimeSeconds is the estimated time it will take for the BGP session to be re-established with peer after a restart. After this period, peer will remove stale routes. This is described RFC 4724 section 4.2.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(4095),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"timers": schema.SingleNestedAttribute{
						Description:         "Timers defines the BGP timers for the peer. If not specified, the default timers are used.",
						MarkdownDescription: "Timers defines the BGP timers for the peer. If not specified, the default timers are used.",
						Attributes: map[string]schema.Attribute{
							"connect_retry_time_seconds": schema.Int64Attribute{
								Description:         "ConnectRetryTimeSeconds defines the initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8). If not specified, defaults to 120 seconds.",
								MarkdownDescription: "ConnectRetryTimeSeconds defines the initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8). If not specified, defaults to 120 seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(2.147483647e+09),
								},
							},

							"hold_time_seconds": schema.Int64Attribute{
								Description:         "HoldTimeSeconds defines the initial value for the BGP HoldTimer (RFC 4271, Section 4.2). Updating this value will cause a session reset. If not specified, defaults to 90 seconds.",
								MarkdownDescription: "HoldTimeSeconds defines the initial value for the BGP HoldTimer (RFC 4271, Section 4.2). Updating this value will cause a session reset. If not specified, defaults to 90 seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(3),
									int64validator.AtMost(65535),
								},
							},

							"keep_alive_time_seconds": schema.Int64Attribute{
								Description:         "KeepaliveTimeSeconds defines the initial value for the BGP KeepaliveTimer (RFC 4271, Section 8). It can not be larger than HoldTimeSeconds. Updating this value will cause a session reset. If not specified, defaults to 30 seconds.",
								MarkdownDescription: "KeepaliveTimeSeconds defines the initial value for the BGP KeepaliveTimer (RFC 4271, Section 8). It can not be larger than HoldTimeSeconds. Updating this value will cause a session reset. If not specified, defaults to 30 seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"transport": schema.SingleNestedAttribute{
						Description:         "Transport defines the BGP transport parameters for the peer. If not specified, the default transport parameters are used.",
						MarkdownDescription: "Transport defines the BGP transport parameters for the peer. If not specified, the default transport parameters are used.",
						Attributes: map[string]schema.Attribute{
							"peer_port": schema.Int64Attribute{
								Description:         "PeerPort is the peer port to be used for the BGP session. If not specified, defaults to TCP port 179.",
								MarkdownDescription: "PeerPort is the peer port to be used for the BGP session. If not specified, defaults to TCP port 179.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumBgppeerConfigV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_peer_config_v2_manifest")

	var model CiliumIoCiliumBgppeerConfigV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumBGPPeerConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
