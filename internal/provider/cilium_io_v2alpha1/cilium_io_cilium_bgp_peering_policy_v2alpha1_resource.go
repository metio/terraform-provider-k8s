/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"time"
)

var (
	_ resource.Resource                = &CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource{}
	_ resource.ResourceWithImportState = &CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource{}
)

func NewCiliumIoCiliumBgppeeringPolicyV2Alpha1Resource() resource.Resource {
	return &CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource{}
}

type CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_peering_policy_v2alpha1"
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var model CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumBGPPeeringPolicy")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumbgppeeringpolicies"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var data CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumbgppeeringpolicies"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var model CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumBGPPeeringPolicy")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumbgppeeringpolicies"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cilium_io_cilium_bgp_peering_policy_v2alpha1")

	var data CiliumIoCiliumBgppeeringPolicyV2Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumbgppeeringpolicies"}).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumbgppeeringpolicies"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *CiliumIoCiliumBgppeeringPolicyV2Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
