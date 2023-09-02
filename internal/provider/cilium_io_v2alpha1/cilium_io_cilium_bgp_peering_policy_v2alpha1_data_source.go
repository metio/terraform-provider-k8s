/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

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
	_ datasource.DataSource              = &CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource{}
)

func NewCiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource() datasource.DataSource {
	return &CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource{}
}

type CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_peering_policy_v2alpha1"
}

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
									Optional:            false,
									Computed:            true,
								},

								"local_asn": schema.Int64Attribute{
									Description:         "LocalASN is the ASN of this virtual router. Supports extended 32bit ASNs",
									MarkdownDescription: "LocalASN is the ASN of this virtual router. Supports extended 32bit ASNs",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"neighbors": schema.ListNestedAttribute{
									Description:         "Neighbors is a list of neighboring BGP peers for this virtual router",
									MarkdownDescription: "Neighbors is a list of neighboring BGP peers for this virtual router",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"connect_retry_time_seconds": schema.Int64Attribute{
												Description:         "ConnectRetryTimeSeconds defines the initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8).",
												MarkdownDescription: "ConnectRetryTimeSeconds defines the initial value for the BGP ConnectRetryTimer (RFC 4271, Section 8).",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"e_bgp_multihop_ttl": schema.Int64Attribute{
												Description:         "EBGPMultihopTTL controls the multi-hop feature for eBGP peers. Its value defines the Time To Live (TTL) value used in BGP packets sent to the neighbor. The value 1 implies that eBGP multi-hop feature is disabled (only a single hop is allowed). This field is ignored for iBGP peers.",
												MarkdownDescription: "EBGPMultihopTTL controls the multi-hop feature for eBGP peers. Its value defines the Time To Live (TTL) value used in BGP packets sent to the neighbor. The value 1 implies that eBGP multi-hop feature is disabled (only a single hop is allowed). This field is ignored for iBGP peers.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"families": schema.ListNestedAttribute{
												Description:         "Families, if provided, defines a set of AFI/SAFIs the speaker will negotiate with it's peer.  If this slice is not provided the default families of IPv6 and IPv4 will be provided.",
												MarkdownDescription: "Families, if provided, defines a set of AFI/SAFIs the speaker will negotiate with it's peer.  If this slice is not provided the default families of IPv6 and IPv4 will be provided.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"afi": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"safi": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"graceful_restart": schema.SingleNestedAttribute{
												Description:         "GracefulRestart defines graceful restart parameters which are negotiated with this neighbor. If empty / nil, the graceful restart capability is disabled.",
												MarkdownDescription: "GracefulRestart defines graceful restart parameters which are negotiated with this neighbor. If empty / nil, the graceful restart capability is disabled.",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "Enabled flag, when set enables graceful restart capability.",
														MarkdownDescription: "Enabled flag, when set enables graceful restart capability.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"restart_time_seconds": schema.Int64Attribute{
														Description:         "RestartTimeSeconds is the estimated time it will take for the BGP session to be re-established with peer after a restart. After this period, peer will remove stale routes. This is described RFC 4724 section 4.2.",
														MarkdownDescription: "RestartTimeSeconds is the estimated time it will take for the BGP session to be re-established with peer after a restart. After this period, peer will remove stale routes. This is described RFC 4724 section 4.2.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"hold_time_seconds": schema.Int64Attribute{
												Description:         "HoldTimeSeconds defines the initial value for the BGP HoldTimer (RFC 4271, Section 4.2). Updating this value will cause a session reset.",
												MarkdownDescription: "HoldTimeSeconds defines the initial value for the BGP HoldTimer (RFC 4271, Section 4.2). Updating this value will cause a session reset.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"keep_alive_time_seconds": schema.Int64Attribute{
												Description:         "KeepaliveTimeSeconds defines the initial value for the BGP KeepaliveTimer (RFC 4271, Section 8). It can not be larger than HoldTimeSeconds. Updating this value will cause a session reset.",
												MarkdownDescription: "KeepaliveTimeSeconds defines the initial value for the BGP KeepaliveTimer (RFC 4271, Section 8). It can not be larger than HoldTimeSeconds. Updating this value will cause a session reset.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"peer_asn": schema.Int64Attribute{
												Description:         "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs",
												MarkdownDescription: "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"peer_address": schema.StringAttribute{
												Description:         "PeerAddress is the IP address of the peer. This must be in CIDR notation and use a /32 to express a single host.",
												MarkdownDescription: "PeerAddress is the IP address of the peer. This must be in CIDR notation and use a /32 to express a single host.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"peer_port": schema.Int64Attribute{
												Description:         "PeerPort is the TCP port of the peer. 1-65535 is the range of valid port numbers that can be specified. If unset, defaults to 179.",
												MarkdownDescription: "PeerPort is the TCP port of the peer. 1-65535 is the range of valid port numbers that can be specified. If unset, defaults to 179.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
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
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
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
		},
	}
}

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var data CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "CiliumBGPPeeringPolicy"}).
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

	var readResponse CiliumIoCiliumBGPPeeringPolicyV2Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("cilium.io/v2alpha1")
	data.Kind = pointer.String("CiliumBGPPeeringPolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
