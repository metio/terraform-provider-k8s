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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CiliumIoCiliumBgpclusterConfigV2Alpha1Manifest{}
)

func NewCiliumIoCiliumBgpclusterConfigV2Alpha1Manifest() datasource.DataSource {
	return &CiliumIoCiliumBgpclusterConfigV2Alpha1Manifest{}
}

type CiliumIoCiliumBgpclusterConfigV2Alpha1Manifest struct{}

type CiliumIoCiliumBgpclusterConfigV2Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BgpInstances *[]struct {
			LocalASN  *int64  `tfsdk:"local_asn" json:"localASN,omitempty"`
			LocalPort *int64  `tfsdk:"local_port" json:"localPort,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Peers     *[]struct {
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				PeerASN       *int64  `tfsdk:"peer_asn" json:"peerASN,omitempty"`
				PeerAddress   *string `tfsdk:"peer_address" json:"peerAddress,omitempty"`
				PeerConfigRef *struct {
					Group *string `tfsdk:"group" json:"group,omitempty"`
					Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"peer_config_ref" json:"peerConfigRef,omitempty"`
			} `tfsdk:"peers" json:"peers,omitempty"`
		} `tfsdk:"bgp_instances" json:"bgpInstances,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumBgpclusterConfigV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_cluster_config_v2alpha1_manifest"
}

func (r *CiliumIoCiliumBgpclusterConfigV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumBGPClusterConfig is the Schema for the CiliumBGPClusterConfig API",
		MarkdownDescription: "CiliumBGPClusterConfig is the Schema for the CiliumBGPClusterConfig API",
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
				Description:         "Spec defines the desired cluster configuration of the BGP control plane.",
				MarkdownDescription: "Spec defines the desired cluster configuration of the BGP control plane.",
				Attributes: map[string]schema.Attribute{
					"bgp_instances": schema.ListNestedAttribute{
						Description:         "A list of CiliumBGPInstance(s) which instructs the BGP control plane how to instantiate virtual BGP routers.",
						MarkdownDescription: "A list of CiliumBGPInstance(s) which instructs the BGP control plane how to instantiate virtual BGP routers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"local_asn": schema.Int64Attribute{
									Description:         "LocalASN is the ASN of this BGP instance. Supports extended 32bit ASNs.",
									MarkdownDescription: "LocalASN is the ASN of this BGP instance. Supports extended 32bit ASNs.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(4.294967295e+09),
									},
								},

								"local_port": schema.Int64Attribute{
									Description:         "LocalPort is the port on which the BGP daemon listens for incoming connections. If not specified, BGP instance will not listen for incoming connections.",
									MarkdownDescription: "LocalPort is the port on which the BGP daemon listens for incoming connections. If not specified, BGP instance will not listen for incoming connections.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(65535),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the BGP instance. It is a unique identifier for the BGP instance within the cluster configuration.",
									MarkdownDescription: "Name is the name of the BGP instance. It is a unique identifier for the BGP instance within the cluster configuration.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(255),
									},
								},

								"peers": schema.ListNestedAttribute{
									Description:         "Peers is a list of neighboring BGP peers for this virtual router",
									MarkdownDescription: "Peers is a list of neighboring BGP peers for this virtual router",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the name of the BGP peer. It is a unique identifier for the peer within the BGP instance.",
												MarkdownDescription: "Name is the name of the BGP peer. It is a unique identifier for the peer within the BGP instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(255),
												},
											},

											"peer_asn": schema.Int64Attribute{
												Description:         "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs. If peerASN is 0, the BGP OPEN message validation of ASN will be disabled and ASN will be determined based on peer's OPEN message.",
												MarkdownDescription: "PeerASN is the ASN of the peer BGP router. Supports extended 32bit ASNs. If peerASN is 0, the BGP OPEN message validation of ASN will be disabled and ASN will be determined based on peer's OPEN message.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(4.294967295e+09),
												},
											},

											"peer_address": schema.StringAttribute{
												Description:         "PeerAddress is the IP address of the neighbor. Supports IPv4 and IPv6 addresses.",
												MarkdownDescription: "PeerAddress is the IP address of the neighbor. Supports IPv4 and IPv6 addresses.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))`), ""),
												},
											},

											"peer_config_ref": schema.SingleNestedAttribute{
												Description:         "PeerConfigRef is a reference to a peer configuration resource. If not specified, the default BGP configuration is used for this peer.",
												MarkdownDescription: "PeerConfigRef is a reference to a peer configuration resource. If not specified, the default BGP configuration is used for this peer.",
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "Group is the group of the peer config resource. If not specified, the default of 'cilium.io' is used.",
														MarkdownDescription: "Group is the group of the peer config resource. If not specified, the default of 'cilium.io' is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the kind of the peer config resource. If not specified, the default of 'CiliumBGPPeerConfig' is used.",
														MarkdownDescription: "Kind is the kind of the peer config resource. If not specified, the default of 'CiliumBGPPeerConfig' is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the peer config resource. Name refers to the name of a Kubernetes object (typically a CiliumBGPPeerConfig).",
														MarkdownDescription: "Name is the name of the peer config resource. Name refers to the name of a Kubernetes object (typically a CiliumBGPPeerConfig).",
														Required:            true,
														Optional:            false,
														Computed:            false,
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
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"node_selector": schema.SingleNestedAttribute{
						Description:         "NodeSelector selects a group of nodes where this BGP Cluster config applies. If empty / nil this config applies to all nodes.",
						MarkdownDescription: "NodeSelector selects a group of nodes where this BGP Cluster config applies. If empty / nil this config applies to all nodes.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumBgpclusterConfigV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_cluster_config_v2alpha1_manifest")

	var model CiliumIoCiliumBgpclusterConfigV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumBGPClusterConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
