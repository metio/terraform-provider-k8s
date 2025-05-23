/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package frrk8s_metallb_io_v1beta1

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
	_ datasource.DataSource = &Frrk8SMetallbIoFrrconfigurationV1Beta1Manifest{}
)

func NewFrrk8SMetallbIoFrrconfigurationV1Beta1Manifest() datasource.DataSource {
	return &Frrk8SMetallbIoFrrconfigurationV1Beta1Manifest{}
}

type Frrk8SMetallbIoFrrconfigurationV1Beta1Manifest struct{}

type Frrk8SMetallbIoFrrconfigurationV1Beta1ManifestData struct {
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
		Bgp *struct {
			BfdProfiles *[]struct {
				DetectMultiplier *int64  `tfsdk:"detect_multiplier" json:"detectMultiplier,omitempty"`
				EchoInterval     *int64  `tfsdk:"echo_interval" json:"echoInterval,omitempty"`
				EchoMode         *bool   `tfsdk:"echo_mode" json:"echoMode,omitempty"`
				MinimumTtl       *int64  `tfsdk:"minimum_ttl" json:"minimumTtl,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				PassiveMode      *bool   `tfsdk:"passive_mode" json:"passiveMode,omitempty"`
				ReceiveInterval  *int64  `tfsdk:"receive_interval" json:"receiveInterval,omitempty"`
				TransmitInterval *int64  `tfsdk:"transmit_interval" json:"transmitInterval,omitempty"`
			} `tfsdk:"bfd_profiles" json:"bfdProfiles,omitempty"`
			Routers *[]struct {
				Asn     *int64  `tfsdk:"asn" json:"asn,omitempty"`
				Id      *string `tfsdk:"id" json:"id,omitempty"`
				Imports *[]struct {
					Vrf *string `tfsdk:"vrf" json:"vrf,omitempty"`
				} `tfsdk:"imports" json:"imports,omitempty"`
				Neighbors *[]struct {
					Address               *string `tfsdk:"address" json:"address,omitempty"`
					Asn                   *int64  `tfsdk:"asn" json:"asn,omitempty"`
					BfdProfile            *string `tfsdk:"bfd_profile" json:"bfdProfile,omitempty"`
					ConnectTime           *string `tfsdk:"connect_time" json:"connectTime,omitempty"`
					DisableMP             *bool   `tfsdk:"disable_mp" json:"disableMP,omitempty"`
					DynamicASN            *string `tfsdk:"dynamic_asn" json:"dynamicASN,omitempty"`
					EbgpMultiHop          *bool   `tfsdk:"ebgp_multi_hop" json:"ebgpMultiHop,omitempty"`
					EnableGracefulRestart *bool   `tfsdk:"enable_graceful_restart" json:"enableGracefulRestart,omitempty"`
					HoldTime              *string `tfsdk:"hold_time" json:"holdTime,omitempty"`
					Interface             *string `tfsdk:"interface" json:"interface,omitempty"`
					KeepaliveTime         *string `tfsdk:"keepalive_time" json:"keepaliveTime,omitempty"`
					Password              *string `tfsdk:"password" json:"password,omitempty"`
					PasswordSecret        *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"password_secret" json:"passwordSecret,omitempty"`
					Port          *int64  `tfsdk:"port" json:"port,omitempty"`
					Sourceaddress *string `tfsdk:"sourceaddress" json:"sourceaddress,omitempty"`
					ToAdvertise   *struct {
						Allowed *struct {
							Mode     *string   `tfsdk:"mode" json:"mode,omitempty"`
							Prefixes *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
						} `tfsdk:"allowed" json:"allowed,omitempty"`
						WithCommunity *[]struct {
							Community *string   `tfsdk:"community" json:"community,omitempty"`
							Prefixes  *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
						} `tfsdk:"with_community" json:"withCommunity,omitempty"`
						WithLocalPref *[]struct {
							LocalPref *int64    `tfsdk:"local_pref" json:"localPref,omitempty"`
							Prefixes  *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
						} `tfsdk:"with_local_pref" json:"withLocalPref,omitempty"`
					} `tfsdk:"to_advertise" json:"toAdvertise,omitempty"`
					ToReceive *struct {
						Allowed *struct {
							Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
							Prefixes *[]struct {
								Ge     *int64  `tfsdk:"ge" json:"ge,omitempty"`
								Le     *int64  `tfsdk:"le" json:"le,omitempty"`
								Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
							} `tfsdk:"prefixes" json:"prefixes,omitempty"`
						} `tfsdk:"allowed" json:"allowed,omitempty"`
					} `tfsdk:"to_receive" json:"toReceive,omitempty"`
				} `tfsdk:"neighbors" json:"neighbors,omitempty"`
				Prefixes *[]string `tfsdk:"prefixes" json:"prefixes,omitempty"`
				Vrf      *string   `tfsdk:"vrf" json:"vrf,omitempty"`
			} `tfsdk:"routers" json:"routers,omitempty"`
		} `tfsdk:"bgp" json:"bgp,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Raw *struct {
			Priority  *int64  `tfsdk:"priority" json:"priority,omitempty"`
			RawConfig *string `tfsdk:"raw_config" json:"rawConfig,omitempty"`
		} `tfsdk:"raw" json:"raw,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Frrk8SMetallbIoFrrconfigurationV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_frrk8s_metallb_io_frr_configuration_v1beta1_manifest"
}

func (r *Frrk8SMetallbIoFrrconfigurationV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FRRConfiguration is a piece of FRR configuration.",
		MarkdownDescription: "FRRConfiguration is a piece of FRR configuration.",
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
				Description:         "FRRConfigurationSpec defines the desired state of FRRConfiguration.",
				MarkdownDescription: "FRRConfigurationSpec defines the desired state of FRRConfiguration.",
				Attributes: map[string]schema.Attribute{
					"bgp": schema.SingleNestedAttribute{
						Description:         "BGP is the configuration related to the BGP protocol.",
						MarkdownDescription: "BGP is the configuration related to the BGP protocol.",
						Attributes: map[string]schema.Attribute{
							"bfd_profiles": schema.ListNestedAttribute{
								Description:         "BFDProfiles is the list of bfd profiles to be used when configuring the neighbors.",
								MarkdownDescription: "BFDProfiles is the list of bfd profiles to be used when configuring the neighbors.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"detect_multiplier": schema.Int64Attribute{
											Description:         "Configures the detection multiplier to determine packet loss. The remote transmission interval will be multiplied by this value to determine the connection loss detection timer.",
											MarkdownDescription: "Configures the detection multiplier to determine packet loss. The remote transmission interval will be multiplied by this value to determine the connection loss detection timer.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(2),
												int64validator.AtMost(255),
											},
										},

										"echo_interval": schema.Int64Attribute{
											Description:         "Configures the minimal echo receive transmission interval that this system is capable of handling in milliseconds. Defaults to 50ms",
											MarkdownDescription: "Configures the minimal echo receive transmission interval that this system is capable of handling in milliseconds. Defaults to 50ms",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(10),
												int64validator.AtMost(60000),
											},
										},

										"echo_mode": schema.BoolAttribute{
											Description:         "Enables or disables the echo transmission mode. This mode is disabled by default, and not supported on multi hops setups.",
											MarkdownDescription: "Enables or disables the echo transmission mode. This mode is disabled by default, and not supported on multi hops setups.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"minimum_ttl": schema.Int64Attribute{
											Description:         "For multi hop sessions only: configure the minimum expected TTL for an incoming BFD control packet.",
											MarkdownDescription: "For multi hop sessions only: configure the minimum expected TTL for an incoming BFD control packet.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(254),
											},
										},

										"name": schema.StringAttribute{
											Description:         "The name of the BFD Profile to be referenced in other parts of the configuration.",
											MarkdownDescription: "The name of the BFD Profile to be referenced in other parts of the configuration.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"passive_mode": schema.BoolAttribute{
											Description:         "Mark session as passive: a passive session will not attempt to start the connection and will wait for control packets from peer before it begins replying.",
											MarkdownDescription: "Mark session as passive: a passive session will not attempt to start the connection and will wait for control packets from peer before it begins replying.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"receive_interval": schema.Int64Attribute{
											Description:         "The minimum interval that this system is capable of receiving control packets in milliseconds. Defaults to 300ms.",
											MarkdownDescription: "The minimum interval that this system is capable of receiving control packets in milliseconds. Defaults to 300ms.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(10),
												int64validator.AtMost(60000),
											},
										},

										"transmit_interval": schema.Int64Attribute{
											Description:         "The minimum transmission interval (less jitter) that this system wants to use to send BFD control packets in milliseconds. Defaults to 300ms",
											MarkdownDescription: "The minimum transmission interval (less jitter) that this system wants to use to send BFD control packets in milliseconds. Defaults to 300ms",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(10),
												int64validator.AtMost(60000),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"routers": schema.ListNestedAttribute{
								Description:         "Routers is the list of routers we want FRR to configure (one per VRF).",
								MarkdownDescription: "Routers is the list of routers we want FRR to configure (one per VRF).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"asn": schema.Int64Attribute{
											Description:         "ASN is the AS number to use for the local end of the session.",
											MarkdownDescription: "ASN is the AS number to use for the local end of the session.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(4.294967295e+09),
											},
										},

										"id": schema.StringAttribute{
											Description:         "ID is the BGP router ID",
											MarkdownDescription: "ID is the BGP router ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"imports": schema.ListNestedAttribute{
											Description:         "Imports is the list of imported VRFs we want for this router / vrf.",
											MarkdownDescription: "Imports is the list of imported VRFs we want for this router / vrf.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"vrf": schema.StringAttribute{
														Description:         "Vrf is the vrf we want to import from",
														MarkdownDescription: "Vrf is the vrf we want to import from",
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

										"neighbors": schema.ListNestedAttribute{
											Description:         "Neighbors is the list of neighbors we want to establish BGP sessions with.",
											MarkdownDescription: "Neighbors is the list of neighbors we want to establish BGP sessions with.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"address": schema.StringAttribute{
														Description:         "Address is the IP address to establish the session with.",
														MarkdownDescription: "Address is the IP address to establish the session with.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"asn": schema.Int64Attribute{
														Description:         "ASN is the AS number to use for the local end of the session. ASN and DynamicASN are mutually exclusive and one of them must be specified.",
														MarkdownDescription: "ASN is the AS number to use for the local end of the session. ASN and DynamicASN are mutually exclusive and one of them must be specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4.294967295e+09),
														},
													},

													"bfd_profile": schema.StringAttribute{
														Description:         "BFDProfile is the name of the BFD Profile to be used for the BFD session associated to the BGP session. If not set, the BFD session won't be set up.",
														MarkdownDescription: "BFDProfile is the name of the BFD Profile to be used for the BFD session associated to the BGP session. If not set, the BFD session won't be set up.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"connect_time": schema.StringAttribute{
														Description:         "Requested BGP connect time, controls how long BGP waits between connection attempts to a neighbor.",
														MarkdownDescription: "Requested BGP connect time, controls how long BGP waits between connection attempts to a neighbor.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"disable_mp": schema.BoolAttribute{
														Description:         "To set if we want to disable MP BGP that will separate IPv4 and IPv6 route exchanges into distinct BGP sessions.",
														MarkdownDescription: "To set if we want to disable MP BGP that will separate IPv4 and IPv6 route exchanges into distinct BGP sessions.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"dynamic_asn": schema.StringAttribute{
														Description:         "DynamicASN detects the AS number to use for the local end of the session without explicitly setting it via the ASN field. Limited to: internal - if the neighbor's ASN is different than the router's the connection is denied. external - if the neighbor's ASN is the same as the router's the connection is denied. ASN and DynamicASN are mutually exclusive and one of them must be specified.",
														MarkdownDescription: "DynamicASN detects the AS number to use for the local end of the session without explicitly setting it via the ASN field. Limited to: internal - if the neighbor's ASN is different than the router's the connection is denied. external - if the neighbor's ASN is the same as the router's the connection is denied. ASN and DynamicASN are mutually exclusive and one of them must be specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("internal", "external"),
														},
													},

													"ebgp_multi_hop": schema.BoolAttribute{
														Description:         "EBGPMultiHop indicates if the BGPPeer is multi-hops away.",
														MarkdownDescription: "EBGPMultiHop indicates if the BGPPeer is multi-hops away.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_graceful_restart": schema.BoolAttribute{
														Description:         "EnableGracefulRestart allows BGP peer to continue to forward data packets along known routes while the routing protocol information is being restored. If the session is already established, the configuration will have effect after reconnecting to the peer",
														MarkdownDescription: "EnableGracefulRestart allows BGP peer to continue to forward data packets along known routes while the routing protocol information is being restored. If the session is already established, the configuration will have effect after reconnecting to the peer",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hold_time": schema.StringAttribute{
														Description:         "HoldTime is the requested BGP hold time, per RFC4271. Defaults to 180s.",
														MarkdownDescription: "HoldTime is the requested BGP hold time, per RFC4271. Defaults to 180s.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"interface": schema.StringAttribute{
														Description:         "Interface is the node interface over which the unnumbered BGP peering will be established. No API validation takes place as that string value represents an interface name on the host and if user provides an invalid value, only the actual BGP session will not be established. Address and Interface are mutually exclusive and one of them must be specified.",
														MarkdownDescription: "Interface is the node interface over which the unnumbered BGP peering will be established. No API validation takes place as that string value represents an interface name on the host and if user provides an invalid value, only the actual BGP session will not be established. Address and Interface are mutually exclusive and one of them must be specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"keepalive_time": schema.StringAttribute{
														Description:         "KeepaliveTime is the requested BGP keepalive time, per RFC4271. Defaults to 60s.",
														MarkdownDescription: "KeepaliveTime is the requested BGP keepalive time, per RFC4271. Defaults to 60s.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password to be used for establishing the BGP session. Password and PasswordSecret are mutually exclusive.",
														MarkdownDescription: "Password to be used for establishing the BGP session. Password and PasswordSecret are mutually exclusive.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password_secret": schema.SingleNestedAttribute{
														Description:         "PasswordSecret is name of the authentication secret for the neighbor. the secret must be of type 'kubernetes.io/basic-auth', and created in the same namespace as the frr-k8s daemon. The password is stored in the secret as the key 'password'. Password and PasswordSecret are mutually exclusive.",
														MarkdownDescription: "PasswordSecret is name of the authentication secret for the neighbor. the secret must be of type 'kubernetes.io/basic-auth', and created in the same namespace as the frr-k8s daemon. The password is stored in the secret as the key 'password'. Password and PasswordSecret are mutually exclusive.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "name is unique within a namespace to reference a secret resource.",
																MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "namespace defines the space within which the secret name must be unique.",
																MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port is the port to dial when establishing the session. Defaults to 179.",
														MarkdownDescription: "Port is the port to dial when establishing the session. Defaults to 179.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(16384),
														},
													},

													"sourceaddress": schema.StringAttribute{
														Description:         "SourceAddress is the IPv4 or IPv6 source address to use for the BGP session to this neighbour, may be specified as either an IP address directly or as an interface name",
														MarkdownDescription: "SourceAddress is the IPv4 or IPv6 source address to use for the BGP session to this neighbour, may be specified as either an IP address directly or as an interface name",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"to_advertise": schema.SingleNestedAttribute{
														Description:         "ToAdvertise represents the list of prefixes to advertise to the given neighbor and the associated properties.",
														MarkdownDescription: "ToAdvertise represents the list of prefixes to advertise to the given neighbor and the associated properties.",
														Attributes: map[string]schema.Attribute{
															"allowed": schema.SingleNestedAttribute{
																Description:         "Allowed is is the list of prefixes allowed to be propagated to this neighbor. They must match the prefixes defined in the router.",
																MarkdownDescription: "Allowed is is the list of prefixes allowed to be propagated to this neighbor. They must match the prefixes defined in the router.",
																Attributes: map[string]schema.Attribute{
																	"mode": schema.StringAttribute{
																		Description:         "Mode is the mode to use when handling the prefixes. When set to 'filtered', only the prefixes in the given list will be allowed. When set to 'all', all the prefixes configured on the router will be allowed.",
																		MarkdownDescription: "Mode is the mode to use when handling the prefixes. When set to 'filtered', only the prefixes in the given list will be allowed. When set to 'all', all the prefixes configured on the router will be allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("all", "filtered"),
																		},
																	},

																	"prefixes": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

															"with_community": schema.ListNestedAttribute{
																Description:         "PrefixesWithCommunity is a list of prefixes that are associated to a bgp community when being advertised. The prefixes associated to a given local pref must be in the prefixes allowed to be advertised.",
																MarkdownDescription: "PrefixesWithCommunity is a list of prefixes that are associated to a bgp community when being advertised. The prefixes associated to a given local pref must be in the prefixes allowed to be advertised.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"community": schema.StringAttribute{
																			Description:         "Community is the community associated to the prefixes.",
																			MarkdownDescription: "Community is the community associated to the prefixes.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"prefixes": schema.ListAttribute{
																			Description:         "Prefixes is the list of prefixes associated to the community.",
																			MarkdownDescription: "Prefixes is the list of prefixes associated to the community.",
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

															"with_local_pref": schema.ListNestedAttribute{
																Description:         "PrefixesWithLocalPref is a list of prefixes that are associated to a local preference when being advertised. The prefixes associated to a given local pref must be in the prefixes allowed to be advertised.",
																MarkdownDescription: "PrefixesWithLocalPref is a list of prefixes that are associated to a local preference when being advertised. The prefixes associated to a given local pref must be in the prefixes allowed to be advertised.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"local_pref": schema.Int64Attribute{
																			Description:         "LocalPref is the local preference associated to the prefixes.",
																			MarkdownDescription: "LocalPref is the local preference associated to the prefixes.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"prefixes": schema.ListAttribute{
																			Description:         "Prefixes is the list of prefixes associated to the local preference.",
																			MarkdownDescription: "Prefixes is the list of prefixes associated to the local preference.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"to_receive": schema.SingleNestedAttribute{
														Description:         "ToReceive represents the list of prefixes to receive from the given neighbor.",
														MarkdownDescription: "ToReceive represents the list of prefixes to receive from the given neighbor.",
														Attributes: map[string]schema.Attribute{
															"allowed": schema.SingleNestedAttribute{
																Description:         "Allowed is the list of prefixes allowed to be received from this neighbor.",
																MarkdownDescription: "Allowed is the list of prefixes allowed to be received from this neighbor.",
																Attributes: map[string]schema.Attribute{
																	"mode": schema.StringAttribute{
																		Description:         "Mode is the mode to use when handling the prefixes. When set to 'filtered', only the prefixes in the given list will be allowed. When set to 'all', all the prefixes configured on the router will be allowed.",
																		MarkdownDescription: "Mode is the mode to use when handling the prefixes. When set to 'filtered', only the prefixes in the given list will be allowed. When set to 'all', all the prefixes configured on the router will be allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("all", "filtered"),
																		},
																	},

																	"prefixes": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"ge": schema.Int64Attribute{
																					Description:         "The prefix length modifier. This selector accepts any matching prefix with length greater or equal the given value.",
																					MarkdownDescription: "The prefix length modifier. This selector accepts any matching prefix with length greater or equal the given value.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(1),
																						int64validator.AtMost(128),
																					},
																				},

																				"le": schema.Int64Attribute{
																					Description:         "The prefix length modifier. This selector accepts any matching prefix with length less or equal the given value.",
																					MarkdownDescription: "The prefix length modifier. This selector accepts any matching prefix with length less or equal the given value.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.Int64{
																						int64validator.AtLeast(1),
																						int64validator.AtMost(128),
																					},
																				},

																				"prefix": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"prefixes": schema.ListAttribute{
											Description:         "Prefixes is the list of prefixes we want to advertise from this router instance.",
											MarkdownDescription: "Prefixes is the list of prefixes we want to advertise from this router instance.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"vrf": schema.StringAttribute{
											Description:         "VRF is the host vrf used to establish sessions from this router.",
											MarkdownDescription: "VRF is the host vrf used to establish sessions from this router.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": schema.SingleNestedAttribute{
						Description:         "NodeSelector limits the nodes that will attempt to apply this config. When specified, the configuration will be considered only on nodes whose labels match the specified selectors. When it is not specified all nodes will attempt to apply this config.",
						MarkdownDescription: "NodeSelector limits the nodes that will attempt to apply this config. When specified, the configuration will be considered only on nodes whose labels match the specified selectors. When it is not specified all nodes will attempt to apply this config.",
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

					"raw": schema.SingleNestedAttribute{
						Description:         "Raw is a snippet of raw frr configuration that gets appended to the one rendered translating the type safe API.",
						MarkdownDescription: "Raw is a snippet of raw frr configuration that gets appended to the one rendered translating the type safe API.",
						Attributes: map[string]schema.Attribute{
							"priority": schema.Int64Attribute{
								Description:         "Priority is the order with this configuration is appended to the bottom of the rendered configuration. A higher value means the raw config is appended later in the configuration file.",
								MarkdownDescription: "Priority is the order with this configuration is appended to the bottom of the rendered configuration. A higher value means the raw config is appended later in the configuration file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"raw_config": schema.StringAttribute{
								Description:         "Config is a raw FRR configuration to be appended to the configuration rendered via the k8s api.",
								MarkdownDescription: "Config is a raw FRR configuration to be appended to the configuration rendered via the k8s api.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Frrk8SMetallbIoFrrconfigurationV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_frrk8s_metallb_io_frr_configuration_v1beta1_manifest")

	var model Frrk8SMetallbIoFrrconfigurationV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("frrk8s.metallb.io/v1beta1")
	model.Kind = pointer.String("FRRConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
