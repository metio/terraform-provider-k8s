/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

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
	_ datasource.DataSource = &CiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest{}
)

func NewCiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest() datasource.DataSource {
	return &CiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest{}
}

type CiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest struct{}

type CiliumIoCiliumBgppeeringPolicyV2Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		VirtualRouters *[]struct {
			ExportPodCIDR *bool  `tfsdk:"export_pod_cidr" json:"exportPodCIDR,omitempty"`
			LocalASN      *int64 `tfsdk:"local_asn" json:"localASN,omitempty"`
			Neighbors     *[]struct {
				AdvertisedPathAttributes *[]struct {
					Communities *struct {
						Large    *[]string `tfsdk:"large" json:"large,omitempty"`
						Standard *[]string `tfsdk:"standard" json:"standard,omitempty"`
					} `tfsdk:"communities" json:"communities,omitempty"`
					LocalPreference *int64 `tfsdk:"local_preference" json:"localPreference,omitempty"`
					Selector        *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					SelectorType *string `tfsdk:"selector_type" json:"selectorType,omitempty"`
				} `tfsdk:"advertised_path_attributes" json:"advertisedPathAttributes,omitempty"`
				ConnectRetryTimeSeconds *int64 `tfsdk:"connect_retry_time_seconds" json:"connectRetryTimeSeconds,omitempty"`
				EBGPMultihopTTL         *int64 `tfsdk:"e_bgp_multihop_ttl" json:"eBGPMultihopTTL,omitempty"`
				Families                *[]struct {
					Afi  *string `tfsdk:"afi" json:"afi,omitempty"`
					Safi *string `tfsdk:"safi" json:"safi,omitempty"`
				} `tfsdk:"families" json:"families,omitempty"`
				GracefulRestart *struct {
					Enabled            *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					RestartTimeSeconds *int64 `tfsdk:"restart_time_seconds" json:"restartTimeSeconds,omitempty"`
				} `tfsdk:"graceful_restart" json:"gracefulRestart,omitempty"`
				HoldTimeSeconds      *int64  `tfsdk:"hold_time_seconds" json:"holdTimeSeconds,omitempty"`
				KeepAliveTimeSeconds *int64  `tfsdk:"keep_alive_time_seconds" json:"keepAliveTimeSeconds,omitempty"`
				PeerASN              *int64  `tfsdk:"peer_asn" json:"peerASN,omitempty"`
				PeerAddress          *string `tfsdk:"peer_address" json:"peerAddress,omitempty"`
				PeerPort             *int64  `tfsdk:"peer_port" json:"peerPort,omitempty"`
			} `tfsdk:"neighbors" json:"neighbors,omitempty"`
			ServiceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"service_selector" json:"serviceSelector,omitempty"`
		} `tfsdk:"virtual_routers" json:"virtualRouters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest"
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumBGPPeeringPolicy is a Kubernetes third-party resource for instructing Cilium's BGP control plane to create virtual BGP routers.",
		MarkdownDescription: "CiliumBGPPeeringPolicy is a Kubernetes third-party resource for instructing Cilium's BGP control plane to create virtual BGP routers.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "Spec is a human readable description of a BGP peering policy",
				MarkdownDescription: "Spec is a human readable description of a BGP peering policy",
				Attributes: map[string]schema.Attribute{
					"node_selector": schema.SingleNestedAttribute{
						Description:         "NodeSelector selects a group of nodes where this BGP Peering Policy applies.  If empty / nil this policy applies to all nodes.",
						MarkdownDescription: "NodeSelector selects a group of nodes where this BGP Peering Policy applies.  If empty / nil this policy applies to all nodes.",
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

					"virtual_routers": schema.ListNestedAttribute{
						Description:         "A list of CiliumBGPVirtualRouter(s) which instructs the BGP control plane how to instantiate virtual BGP routers.",
						MarkdownDescription: "A list of CiliumBGPVirtualRouter(s) which instructs the BGP control plane how to instantiate virtual BGP routers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"export_pod_cidr": schema.BoolAttribute{
									Description:         "ExportPodCIDR determines whether to export the Node's private CIDR block to the configured neighbors.",
									MarkdownDescription: "ExportPodCIDR determines whether to export the Node's private CIDR block to the configured neighbors.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"local_asn": schema.Int64Attribute{
									Description:         "LocalASN is the ASN of this virtual router. Supports extended 32bit ASNs",
									MarkdownDescription: "LocalASN is the ASN of this virtual router. Supports extended 32bit ASNs",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(4.294967295e+09),
									},
								},

								"neighbors": schema.ListNestedAttribute{
									Description:         "Neighbors is a list of neighboring BGP peers for this virtual router",
									MarkdownDescription: "Neighbors is a list of neighboring BGP peers for this virtual router",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"advertised_path_attributes": schema.ListNestedAttribute{
												Description:         "AdvertisedPathAttributes can be used to apply additional path attributes to selected routes when advertising them to the peer. If empty / nil, no additional path attributes are advertised.",
												MarkdownDescription: "AdvertisedPathAttributes can be used to apply additional path attributes to selected routes when advertising them to the peer. If empty / nil, no additional path attributes are advertised.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"communities": schema.SingleNestedAttribute{
															Description:         "Communities defines a set of community values advertised in the supported BGP Communities path attributes. If nil / not set, no BGP Communities path attribute will be advertised.",
															MarkdownDescription: "Communities defines a set of community values advertised in the supported BGP Communities path attributes. If nil / not set, no BGP Communities path attribute will be advertised.",
															Attributes: map[string]schema.Attribute{
																"large": schema.ListAttribute{
																	Description:         "Large holds a list of the BGP Large Communities Attribute (RFC 8092) values.",
																	MarkdownDescription: "Large holds a list of the BGP Large Communities Attribute (RFC 8092) values.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"standard": schema.ListAttribute{
																	Description:         "Standard holds a list of 'standard' 32-bit BGP Communities Attribute (RFC 1997) values.",
																	MarkdownDescription: "Standard holds a list of 'standard' 32-bit BGP Communities Attribute (RFC 1997) values.",
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

														"local_preference": schema.Int64Attribute{
															Description:         "LocalPreference defines the preference value advertised in the BGP Local Preference path attribute. As Local Preference is only valid for iBGP peers, this value will be ignored for eBGP peers (no Local Preference path attribute will be advertised). If nil / not set, the default Local Preference of 100 will be advertised in the Local Preference path attribute for iBGP peers.",
															MarkdownDescription: "LocalPreference defines the preference value advertised in the BGP Local Preference path attribute. As Local Preference is only valid for iBGP peers, this value will be ignored for eBGP peers (no Local Preference path attribute will be advertised). If nil / not set, the default Local Preference of 100 will be advertised in the Local Preference path attribute for iBGP peers.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(4.294967295e+09),
															},
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "Selector selects a group of objects of the SelectorType resulting into routes that will be announced with the configured Attributes. If nil / not set, all objects of the SelectorType are selected.",
															MarkdownDescription: "Selector selects a group of objects of the SelectorType resulting into routes that will be announced with the configured Attributes. If nil / not set, all objects of the SelectorType are selected.",
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

														"selector_type": schema.StringAttribute{
															Description:         "SelectorType defines the object type on which the Selector applies: - For 'PodCIDR' the Selector matches k8s CiliumNode resources (path attributes apply to routes announced for PodCIDRs of selected CiliumNodes. Only affects routes of cluster scope / Kubernetes IPAM CIDRs, not Multi-Pool IPAM CIDRs. - For 'CiliumLoadBalancerIPPool' the Selector matches CiliumLoadBalancerIPPool custom resources (path attributes apply to routes announced for selected CiliumLoadBalancerIPPools).",
															MarkdownDescription: "SelectorType defines the object type on which the Selector applies: - For 'PodCIDR' the Selector matches k8s CiliumNode resources (path attributes apply to routes announced for PodCIDRs of selected CiliumNodes. Only affects routes of cluster scope / Kubernetes IPAM CIDRs, not Multi-Pool IPAM CIDRs. - For 'CiliumLoadBalancerIPPool' the Selector matches CiliumLoadBalancerIPPool custom resources (path attributes apply to routes announced for selected CiliumLoadBalancerIPPools).",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("PodCIDR", "CiliumLoadBalancerIPPool"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"connect_retry_time_seconds": schema.Int64Attribute{
												Description:         "ConnectRetryTimeSeconds defines the initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8).",
												MarkdownDescription: "ConnectRetryTimeSeconds defines the initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8).",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(2.147483647e+09),
												},
											},

											"e_bgp_multihop_ttl": schema.Int64Attribute{
												Description:         "EBGPMultihopTTL controls the multi-hop feature for eBGP peers. Its value defines the Time To Live (TTL) value used in BGP packets sent to the neighbor. The value 1 implies that eBGP multi-hop feature is disabled (only a single hop is allowed). This field is ignored for iBGP peers.",
												MarkdownDescription: "EBGPMultihopTTL controls the multi-hop feature for eBGP peers. Its value defines the Time To Live (TTL) value used in BGP packets sent to the neighbor. The value 1 implies that eBGP multi-hop feature is disabled (only a single hop is allowed). This field is ignored for iBGP peers.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(255),
												},
											},

											"families": schema.ListNestedAttribute{
												Description:         "Families, if provided, defines a set of AFI/SAFIs the speaker will negotiate with it's peer.  If this slice is not provided the default families of IPv6 and IPv4 will be provided.",
												MarkdownDescription: "Families, if provided, defines a set of AFI/SAFIs the speaker will negotiate with it's peer.  If this slice is not provided the default families of IPv6 and IPv4 will be provided.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"afi": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("ipv4", "ipv6", "l2vpn", "ls", "opaque"),
															},
														},

														"safi": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "GracefulRestart defines graceful restart parameters which are negotiated with this neighbor. If empty / nil, the graceful restart capability is disabled.",
												MarkdownDescription: "GracefulRestart defines graceful restart parameters which are negotiated with this neighbor. If empty / nil, the graceful restart capability is disabled.",
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

											"hold_time_seconds": schema.Int64Attribute{
												Description:         "HoldTimeSeconds defines the initial value for the BGP HoldTimer (RFC 4271, Section 4.2). Updating this value will cause a session reset.",
												MarkdownDescription: "HoldTimeSeconds defines the initial value for the BGP HoldTimer (RFC 4271, Section 4.2). Updating this value will cause a session reset.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(3),
													int64validator.AtMost(65535),
												},
											},

											"keep_alive_time_seconds": schema.Int64Attribute{
												Description:         "KeepaliveTimeSeconds defines the initial value for the BGP KeepaliveTimer (RFC 4271, Section 8). It can not be larger than HoldTimeSeconds. Updating this value will cause a session reset.",
												MarkdownDescription: "KeepaliveTimeSeconds defines the initial value for the BGP KeepaliveTimer (RFC 4271, Section 8). It can not be larger than HoldTimeSeconds. Updating this value will cause a session reset.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"peer_asn": schema.Int64Attribute{
												Description:         "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs",
												MarkdownDescription: "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"peer_address": schema.StringAttribute{
												Description:         "PeerAddress is the IP address of the peer. This must be in CIDR notation and use a /32 to express a single host.",
												MarkdownDescription: "PeerAddress is the IP address of the peer. This must be in CIDR notation and use a /32 to express a single host.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"peer_port": schema.Int64Attribute{
												Description:         "PeerPort is the TCP port of the peer. 1-65535 is the range of valid port numbers that can be specified. If unset, defaults to 179.",
												MarkdownDescription: "PeerPort is the TCP port of the peer. 1-65535 is the range of valid port numbers that can be specified. If unset, defaults to 179.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"service_selector": schema.SingleNestedAttribute{
									Description:         "ServiceSelector selects a group of load balancer services which this virtual router will announce. The loadBalancerClass for a service must be nil or specify a class supported by Cilium, e.g. 'io.cilium/bgp-control-plane'. Refer to the following document for additional details regarding load balancer classes:  https://kubernetes.io/docs/concepts/services-networking/service/#load-balancer-class  If empty / nil no services will be announced.",
									MarkdownDescription: "ServiceSelector selects a group of load balancer services which this virtual router will announce. The loadBalancerClass for a service must be nil or specify a class supported by Cilium, e.g. 'io.cilium/bgp-control-plane'. Refer to the following document for additional details regarding load balancer classes:  https://kubernetes.io/docs/concepts/services-networking/service/#load-balancer-class  If empty / nil no services will be announced.",
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
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1_manifest")

	var model CiliumIoCiliumBgppeeringPolicyV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumBGPPeeringPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
