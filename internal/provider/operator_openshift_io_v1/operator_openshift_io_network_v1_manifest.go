/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_openshift_io_v1

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
	_ datasource.DataSource = &OperatorOpenshiftIoNetworkV1Manifest{}
)

func NewOperatorOpenshiftIoNetworkV1Manifest() datasource.DataSource {
	return &OperatorOpenshiftIoNetworkV1Manifest{}
}

type OperatorOpenshiftIoNetworkV1Manifest struct{}

type OperatorOpenshiftIoNetworkV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdditionalNetworks *[]struct {
			Name                *string `tfsdk:"name" json:"name,omitempty"`
			Namespace           *string `tfsdk:"namespace" json:"namespace,omitempty"`
			RawCNIConfig        *string `tfsdk:"raw_cni_config" json:"rawCNIConfig,omitempty"`
			SimpleMacvlanConfig *struct {
				IpamConfig *struct {
					StaticIPAMConfig *struct {
						Addresses *[]struct {
							Address *string `tfsdk:"address" json:"address,omitempty"`
							Gateway *string `tfsdk:"gateway" json:"gateway,omitempty"`
						} `tfsdk:"addresses" json:"addresses,omitempty"`
						Dns *struct {
							Domain      *string   `tfsdk:"domain" json:"domain,omitempty"`
							Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
							Search      *[]string `tfsdk:"search" json:"search,omitempty"`
						} `tfsdk:"dns" json:"dns,omitempty"`
						Routes *[]struct {
							Destination *string `tfsdk:"destination" json:"destination,omitempty"`
							Gateway     *string `tfsdk:"gateway" json:"gateway,omitempty"`
						} `tfsdk:"routes" json:"routes,omitempty"`
					} `tfsdk:"static_ipam_config" json:"staticIPAMConfig,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ipam_config" json:"ipamConfig,omitempty"`
				Master *string `tfsdk:"master" json:"master,omitempty"`
				Mode   *string `tfsdk:"mode" json:"mode,omitempty"`
				Mtu    *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
			} `tfsdk:"simple_macvlan_config" json:"simpleMacvlanConfig,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"additional_networks" json:"additionalNetworks,omitempty"`
		ClusterNetwork *[]struct {
			Cidr       *string `tfsdk:"cidr" json:"cidr,omitempty"`
			HostPrefix *int64  `tfsdk:"host_prefix" json:"hostPrefix,omitempty"`
		} `tfsdk:"cluster_network" json:"clusterNetwork,omitempty"`
		DefaultNetwork *struct {
			OpenshiftSDNConfig *struct {
				EnableUnidling         *bool   `tfsdk:"enable_unidling" json:"enableUnidling,omitempty"`
				Mode                   *string `tfsdk:"mode" json:"mode,omitempty"`
				Mtu                    *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
				UseExternalOpenvswitch *bool   `tfsdk:"use_external_openvswitch" json:"useExternalOpenvswitch,omitempty"`
				VxlanPort              *int64  `tfsdk:"vxlan_port" json:"vxlanPort,omitempty"`
			} `tfsdk:"openshift_sdn_config" json:"openshiftSDNConfig,omitempty"`
			OvnKubernetesConfig *struct {
				EgressIPConfig *struct {
					ReachabilityTotalTimeoutSeconds *int64 `tfsdk:"reachability_total_timeout_seconds" json:"reachabilityTotalTimeoutSeconds,omitempty"`
				} `tfsdk:"egress_ip_config" json:"egressIPConfig,omitempty"`
				GatewayConfig *struct {
					IpForwarding *string `tfsdk:"ip_forwarding" json:"ipForwarding,omitempty"`
					Ipv4         *struct {
						InternalMasqueradeSubnet *string `tfsdk:"internal_masquerade_subnet" json:"internalMasqueradeSubnet,omitempty"`
					} `tfsdk:"ipv4" json:"ipv4,omitempty"`
					Ipv6 *struct {
						InternalMasqueradeSubnet *string `tfsdk:"internal_masquerade_subnet" json:"internalMasqueradeSubnet,omitempty"`
					} `tfsdk:"ipv6" json:"ipv6,omitempty"`
					RoutingViaHost *bool `tfsdk:"routing_via_host" json:"routingViaHost,omitempty"`
				} `tfsdk:"gateway_config" json:"gatewayConfig,omitempty"`
				GenevePort          *int64 `tfsdk:"geneve_port" json:"genevePort,omitempty"`
				HybridOverlayConfig *struct {
					HybridClusterNetwork *[]struct {
						Cidr       *string `tfsdk:"cidr" json:"cidr,omitempty"`
						HostPrefix *int64  `tfsdk:"host_prefix" json:"hostPrefix,omitempty"`
					} `tfsdk:"hybrid_cluster_network" json:"hybridClusterNetwork,omitempty"`
					HybridOverlayVXLANPort *int64 `tfsdk:"hybrid_overlay_vxlan_port" json:"hybridOverlayVXLANPort,omitempty"`
				} `tfsdk:"hybrid_overlay_config" json:"hybridOverlayConfig,omitempty"`
				IpsecConfig *struct {
					Mode *string `tfsdk:"mode" json:"mode,omitempty"`
				} `tfsdk:"ipsec_config" json:"ipsecConfig,omitempty"`
				Mtu               *int64 `tfsdk:"mtu" json:"mtu,omitempty"`
				PolicyAuditConfig *struct {
					Destination    *string `tfsdk:"destination" json:"destination,omitempty"`
					MaxFileSize    *int64  `tfsdk:"max_file_size" json:"maxFileSize,omitempty"`
					MaxLogFiles    *int64  `tfsdk:"max_log_files" json:"maxLogFiles,omitempty"`
					RateLimit      *int64  `tfsdk:"rate_limit" json:"rateLimit,omitempty"`
					SyslogFacility *string `tfsdk:"syslog_facility" json:"syslogFacility,omitempty"`
				} `tfsdk:"policy_audit_config" json:"policyAuditConfig,omitempty"`
				V4InternalSubnet *string `tfsdk:"v4_internal_subnet" json:"v4InternalSubnet,omitempty"`
				V6InternalSubnet *string `tfsdk:"v6_internal_subnet" json:"v6InternalSubnet,omitempty"`
			} `tfsdk:"ovn_kubernetes_config" json:"ovnKubernetesConfig,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"default_network" json:"defaultNetwork,omitempty"`
		DeployKubeProxy           *bool `tfsdk:"deploy_kube_proxy" json:"deployKubeProxy,omitempty"`
		DisableMultiNetwork       *bool `tfsdk:"disable_multi_network" json:"disableMultiNetwork,omitempty"`
		DisableNetworkDiagnostics *bool `tfsdk:"disable_network_diagnostics" json:"disableNetworkDiagnostics,omitempty"`
		ExportNetworkFlows        *struct {
			Ipfix *struct {
				Collectors *[]string `tfsdk:"collectors" json:"collectors,omitempty"`
			} `tfsdk:"ipfix" json:"ipfix,omitempty"`
			NetFlow *struct {
				Collectors *[]string `tfsdk:"collectors" json:"collectors,omitempty"`
			} `tfsdk:"net_flow" json:"netFlow,omitempty"`
			SFlow *struct {
				Collectors *[]string `tfsdk:"collectors" json:"collectors,omitempty"`
			} `tfsdk:"s_flow" json:"sFlow,omitempty"`
		} `tfsdk:"export_network_flows" json:"exportNetworkFlows,omitempty"`
		KubeProxyConfig *struct {
			BindAddress        *string              `tfsdk:"bind_address" json:"bindAddress,omitempty"`
			IptablesSyncPeriod *string              `tfsdk:"iptables_sync_period" json:"iptablesSyncPeriod,omitempty"`
			ProxyArguments     *map[string][]string `tfsdk:"proxy_arguments" json:"proxyArguments,omitempty"`
		} `tfsdk:"kube_proxy_config" json:"kubeProxyConfig,omitempty"`
		LogLevel        *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		ManagementState *string `tfsdk:"management_state" json:"managementState,omitempty"`
		Migration       *struct {
			Features *struct {
				EgressFirewall *bool `tfsdk:"egress_firewall" json:"egressFirewall,omitempty"`
				EgressIP       *bool `tfsdk:"egress_ip" json:"egressIP,omitempty"`
				Multicast      *bool `tfsdk:"multicast" json:"multicast,omitempty"`
			} `tfsdk:"features" json:"features,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			Mtu  *struct {
				Machine *struct {
					From *int64 `tfsdk:"from" json:"from,omitempty"`
					To   *int64 `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"machine" json:"machine,omitempty"`
				Network *struct {
					From *int64 `tfsdk:"from" json:"from,omitempty"`
					To   *int64 `tfsdk:"to" json:"to,omitempty"`
				} `tfsdk:"network" json:"network,omitempty"`
			} `tfsdk:"mtu" json:"mtu,omitempty"`
			NetworkType *string `tfsdk:"network_type" json:"networkType,omitempty"`
		} `tfsdk:"migration" json:"migration,omitempty"`
		ObservedConfig             *map[string]string `tfsdk:"observed_config" json:"observedConfig,omitempty"`
		OperatorLogLevel           *string            `tfsdk:"operator_log_level" json:"operatorLogLevel,omitempty"`
		ServiceNetwork             *[]string          `tfsdk:"service_network" json:"serviceNetwork,omitempty"`
		UnsupportedConfigOverrides *map[string]string `tfsdk:"unsupported_config_overrides" json:"unsupportedConfigOverrides,omitempty"`
		UseMultiNetworkPolicy      *bool              `tfsdk:"use_multi_network_policy" json:"useMultiNetworkPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenshiftIoNetworkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_openshift_io_network_v1_manifest"
}

func (r *OperatorOpenshiftIoNetworkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Network describes the cluster's desired network configuration. It is consumed by the cluster-network-operator.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Network describes the cluster's desired network configuration. It is consumed by the cluster-network-operator.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "NetworkSpec is the top-level network configuration object.",
				MarkdownDescription: "NetworkSpec is the top-level network configuration object.",
				Attributes: map[string]schema.Attribute{
					"additional_networks": schema.ListNestedAttribute{
						Description:         "additionalNetworks is a list of extra networks to make available to pods when multiple networks are enabled.",
						MarkdownDescription: "additionalNetworks is a list of extra networks to make available to pods when multiple networks are enabled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "name is the name of the network. This will be populated in the resulting CRD This must be unique.",
									MarkdownDescription: "name is the name of the network. This will be populated in the resulting CRD This must be unique.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "namespace is the namespace of the network. This will be populated in the resulting CRD If not given the network will be created in the default namespace.",
									MarkdownDescription: "namespace is the namespace of the network. This will be populated in the resulting CRD If not given the network will be created in the default namespace.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"raw_cni_config": schema.StringAttribute{
									Description:         "rawCNIConfig is the raw CNI configuration json to create in the NetworkAttachmentDefinition CRD",
									MarkdownDescription: "rawCNIConfig is the raw CNI configuration json to create in the NetworkAttachmentDefinition CRD",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"simple_macvlan_config": schema.SingleNestedAttribute{
									Description:         "SimpleMacvlanConfig configures the macvlan interface in case of type:NetworkTypeSimpleMacvlan",
									MarkdownDescription: "SimpleMacvlanConfig configures the macvlan interface in case of type:NetworkTypeSimpleMacvlan",
									Attributes: map[string]schema.Attribute{
										"ipam_config": schema.SingleNestedAttribute{
											Description:         "IPAMConfig configures IPAM module will be used for IP Address Management (IPAM).",
											MarkdownDescription: "IPAMConfig configures IPAM module will be used for IP Address Management (IPAM).",
											Attributes: map[string]schema.Attribute{
												"static_ipam_config": schema.SingleNestedAttribute{
													Description:         "StaticIPAMConfig configures the static IP address in case of type:IPAMTypeStatic",
													MarkdownDescription: "StaticIPAMConfig configures the static IP address in case of type:IPAMTypeStatic",
													Attributes: map[string]schema.Attribute{
														"addresses": schema.ListNestedAttribute{
															Description:         "Addresses configures IP address for the interface",
															MarkdownDescription: "Addresses configures IP address for the interface",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"address": schema.StringAttribute{
																		Description:         "Address is the IP address in CIDR format",
																		MarkdownDescription: "Address is the IP address in CIDR format",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"gateway": schema.StringAttribute{
																		Description:         "Gateway is IP inside of subnet to designate as the gateway",
																		MarkdownDescription: "Gateway is IP inside of subnet to designate as the gateway",
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

														"dns": schema.SingleNestedAttribute{
															Description:         "DNS configures DNS for the interface",
															MarkdownDescription: "DNS configures DNS for the interface",
															Attributes: map[string]schema.Attribute{
																"domain": schema.StringAttribute{
																	Description:         "Domain configures the domainname the local domain used for short hostname lookups",
																	MarkdownDescription: "Domain configures the domainname the local domain used for short hostname lookups",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"nameservers": schema.ListAttribute{
																	Description:         "Nameservers points DNS servers for IP lookup",
																	MarkdownDescription: "Nameservers points DNS servers for IP lookup",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"search": schema.ListAttribute{
																	Description:         "Search configures priority ordered search domains for short hostname lookups",
																	MarkdownDescription: "Search configures priority ordered search domains for short hostname lookups",
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

														"routes": schema.ListNestedAttribute{
															Description:         "Routes configures IP routes for the interface",
															MarkdownDescription: "Routes configures IP routes for the interface",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination": schema.StringAttribute{
																		Description:         "Destination points the IP route destination",
																		MarkdownDescription: "Destination points the IP route destination",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"gateway": schema.StringAttribute{
																		Description:         "Gateway is the route's next-hop IP address If unset, a default gateway is assumed (as determined by the CNI plugin).",
																		MarkdownDescription: "Gateway is the route's next-hop IP address If unset, a default gateway is assumed (as determined by the CNI plugin).",
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

												"type": schema.StringAttribute{
													Description:         "Type is the type of IPAM module will be used for IP Address Management(IPAM). The supported values are IPAMTypeDHCP, IPAMTypeStatic",
													MarkdownDescription: "Type is the type of IPAM module will be used for IP Address Management(IPAM). The supported values are IPAMTypeDHCP, IPAMTypeStatic",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"master": schema.StringAttribute{
											Description:         "master is the host interface to create the macvlan interface from. If not specified, it will be default route interface",
											MarkdownDescription: "master is the host interface to create the macvlan interface from. If not specified, it will be default route interface",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "mode is the macvlan mode: bridge, private, vepa, passthru. The default is bridge",
											MarkdownDescription: "mode is the macvlan mode: bridge, private, vepa, passthru. The default is bridge",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mtu": schema.Int64Attribute{
											Description:         "mtu is the mtu to use for the macvlan interface. if unset, host's kernel will select the value.",
											MarkdownDescription: "mtu is the mtu to use for the macvlan interface. if unset, host's kernel will select the value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "type is the type of network The supported values are NetworkTypeRaw, NetworkTypeSimpleMacvlan",
									MarkdownDescription: "type is the type of network The supported values are NetworkTypeRaw, NetworkTypeSimpleMacvlan",
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

					"cluster_network": schema.ListNestedAttribute{
						Description:         "clusterNetwork is the IP address pool to use for pod IPs. Some network providers, e.g. OpenShift SDN, support multiple ClusterNetworks. Others only support one. This is equivalent to the cluster-cidr.",
						MarkdownDescription: "clusterNetwork is the IP address pool to use for pod IPs. Some network providers, e.g. OpenShift SDN, support multiple ClusterNetworks. Others only support one. This is equivalent to the cluster-cidr.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"host_prefix": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_network": schema.SingleNestedAttribute{
						Description:         "defaultNetwork is the 'default' network that all pods will receive",
						MarkdownDescription: "defaultNetwork is the 'default' network that all pods will receive",
						Attributes: map[string]schema.Attribute{
							"openshift_sdn_config": schema.SingleNestedAttribute{
								Description:         "openShiftSDNConfig configures the openshift-sdn plugin",
								MarkdownDescription: "openShiftSDNConfig configures the openshift-sdn plugin",
								Attributes: map[string]schema.Attribute{
									"enable_unidling": schema.BoolAttribute{
										Description:         "enableUnidling controls whether or not the service proxy will support idling and unidling of services. By default, unidling is enabled.",
										MarkdownDescription: "enableUnidling controls whether or not the service proxy will support idling and unidling of services. By default, unidling is enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mode": schema.StringAttribute{
										Description:         "mode is one of 'Multitenant', 'Subnet', or 'NetworkPolicy'",
										MarkdownDescription: "mode is one of 'Multitenant', 'Subnet', or 'NetworkPolicy'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mtu": schema.Int64Attribute{
										Description:         "mtu is the mtu to use for the tunnel interface. Defaults to 1450 if unset. This must be 50 bytes smaller than the machine's uplink.",
										MarkdownDescription: "mtu is the mtu to use for the tunnel interface. Defaults to 1450 if unset. This must be 50 bytes smaller than the machine's uplink.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"use_external_openvswitch": schema.BoolAttribute{
										Description:         "useExternalOpenvswitch used to control whether the operator would deploy an OVS DaemonSet itself or expect someone else to start OVS. As of 4.6, OVS is always run as a system service, and this flag is ignored. DEPRECATED: non-functional as of 4.6",
										MarkdownDescription: "useExternalOpenvswitch used to control whether the operator would deploy an OVS DaemonSet itself or expect someone else to start OVS. As of 4.6, OVS is always run as a system service, and this flag is ignored. DEPRECATED: non-functional as of 4.6",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vxlan_port": schema.Int64Attribute{
										Description:         "vxlanPort is the port to use for all vxlan packets. The default is 4789.",
										MarkdownDescription: "vxlanPort is the port to use for all vxlan packets. The default is 4789.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ovn_kubernetes_config": schema.SingleNestedAttribute{
								Description:         "ovnKubernetesConfig configures the ovn-kubernetes plugin.",
								MarkdownDescription: "ovnKubernetesConfig configures the ovn-kubernetes plugin.",
								Attributes: map[string]schema.Attribute{
									"egress_ip_config": schema.SingleNestedAttribute{
										Description:         "egressIPConfig holds the configuration for EgressIP options.",
										MarkdownDescription: "egressIPConfig holds the configuration for EgressIP options.",
										Attributes: map[string]schema.Attribute{
											"reachability_total_timeout_seconds": schema.Int64Attribute{
												Description:         "reachabilityTotalTimeout configures the EgressIP node reachability check total timeout in seconds. If the EgressIP node cannot be reached within this timeout, the node is declared down. Setting a large value may cause the EgressIP feature to react slowly to node changes. In particular, it may react slowly for EgressIP nodes that really have a genuine problem and are unreachable. When omitted, this means the user has no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is 1 second. A value of 0 disables the EgressIP node's reachability check.",
												MarkdownDescription: "reachabilityTotalTimeout configures the EgressIP node reachability check total timeout in seconds. If the EgressIP node cannot be reached within this timeout, the node is declared down. Setting a large value may cause the EgressIP feature to react slowly to node changes. In particular, it may react slowly for EgressIP nodes that really have a genuine problem and are unreachable. When omitted, this means the user has no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is 1 second. A value of 0 disables the EgressIP node's reachability check.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(60),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"gateway_config": schema.SingleNestedAttribute{
										Description:         "gatewayConfig holds the configuration for node gateway options.",
										MarkdownDescription: "gatewayConfig holds the configuration for node gateway options.",
										Attributes: map[string]schema.Attribute{
											"ip_forwarding": schema.StringAttribute{
												Description:         "IPForwarding controls IP forwarding for all traffic on OVN-Kubernetes managed interfaces (such as br-ex). By default this is set to Restricted, and Kubernetes related traffic is still forwarded appropriately, but other IP traffic will not be routed by the OCP node. If there is a desire to allow the host to forward traffic across OVN-Kubernetes managed interfaces, then set this field to 'Global'. The supported values are 'Restricted' and 'Global'.",
												MarkdownDescription: "IPForwarding controls IP forwarding for all traffic on OVN-Kubernetes managed interfaces (such as br-ex). By default this is set to Restricted, and Kubernetes related traffic is still forwarded appropriately, but other IP traffic will not be routed by the OCP node. If there is a desire to allow the host to forward traffic across OVN-Kubernetes managed interfaces, then set this field to 'Global'. The supported values are 'Restricted' and 'Global'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ipv4": schema.SingleNestedAttribute{
												Description:         "ipv4 allows users to configure IP settings for IPv4 connections. When omitted, this means no opinion and the default configuration is used. Check individual members fields within ipv4 for details of default values.",
												MarkdownDescription: "ipv4 allows users to configure IP settings for IPv4 connections. When omitted, this means no opinion and the default configuration is used. Check individual members fields within ipv4 for details of default values.",
												Attributes: map[string]schema.Attribute{
													"internal_masquerade_subnet": schema.StringAttribute{
														Description:         "internalMasqueradeSubnet contains the masquerade addresses in IPV4 CIDR format used internally by ovn-kubernetes to enable host to service traffic. Each host in the cluster is configured with these addresses, as well as the shared gateway bridge interface. The values can be changed after installation. The subnet chosen should not overlap with other networks specified for OVN-Kubernetes as well as other networks used on the host. Additionally the subnet must be large enough to accommodate 6 IPs (maximum prefix length /29). When omitted, this means no opinion and the platform is left to choose a reasonable default which is subject to change over time. The current default subnet is 169.254.169.0/29 The value must be in proper IPV4 CIDR format",
														MarkdownDescription: "internalMasqueradeSubnet contains the masquerade addresses in IPV4 CIDR format used internally by ovn-kubernetes to enable host to service traffic. Each host in the cluster is configured with these addresses, as well as the shared gateway bridge interface. The values can be changed after installation. The subnet chosen should not overlap with other networks specified for OVN-Kubernetes as well as other networks used on the host. Additionally the subnet must be large enough to accommodate 6 IPs (maximum prefix length /29). When omitted, this means no opinion and the platform is left to choose a reasonable default which is subject to change over time. The current default subnet is 169.254.169.0/29 The value must be in proper IPV4 CIDR format",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(18),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ipv6": schema.SingleNestedAttribute{
												Description:         "ipv6 allows users to configure IP settings for IPv6 connections. When omitted, this means no opinion and the default configuration is used. Check individual members fields within ipv6 for details of default values.",
												MarkdownDescription: "ipv6 allows users to configure IP settings for IPv6 connections. When omitted, this means no opinion and the default configuration is used. Check individual members fields within ipv6 for details of default values.",
												Attributes: map[string]schema.Attribute{
													"internal_masquerade_subnet": schema.StringAttribute{
														Description:         "internalMasqueradeSubnet contains the masquerade addresses in IPV6 CIDR format used internally by ovn-kubernetes to enable host to service traffic. Each host in the cluster is configured with these addresses, as well as the shared gateway bridge interface. The values can be changed after installation. The subnet chosen should not overlap with other networks specified for OVN-Kubernetes as well as other networks used on the host. Additionally the subnet must be large enough to accommodate 6 IPs (maximum prefix length /125). When omitted, this means no opinion and the platform is left to choose a reasonable default which is subject to change over time. The current default subnet is fd69::/125 Note that IPV6 dual addresses are not permitted",
														MarkdownDescription: "internalMasqueradeSubnet contains the masquerade addresses in IPV6 CIDR format used internally by ovn-kubernetes to enable host to service traffic. Each host in the cluster is configured with these addresses, as well as the shared gateway bridge interface. The values can be changed after installation. The subnet chosen should not overlap with other networks specified for OVN-Kubernetes as well as other networks used on the host. Additionally the subnet must be large enough to accommodate 6 IPs (maximum prefix length /125). When omitted, this means no opinion and the platform is left to choose a reasonable default which is subject to change over time. The current default subnet is fd69::/125 Note that IPV6 dual addresses are not permitted",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"routing_via_host": schema.BoolAttribute{
												Description:         "RoutingViaHost allows pod egress traffic to exit via the ovn-k8s-mp0 management port into the host before sending it out. If this is not set, traffic will always egress directly from OVN to outside without touching the host stack. Setting this to true means hardware offload will not be supported. Default is false if GatewayConfig is specified.",
												MarkdownDescription: "RoutingViaHost allows pod egress traffic to exit via the ovn-k8s-mp0 management port into the host before sending it out. If this is not set, traffic will always egress directly from OVN to outside without touching the host stack. Setting this to true means hardware offload will not be supported. Default is false if GatewayConfig is specified.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"geneve_port": schema.Int64Attribute{
										Description:         "geneve port is the UDP port to be used by geneve encapulation. Default is 6081",
										MarkdownDescription: "geneve port is the UDP port to be used by geneve encapulation. Default is 6081",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"hybrid_overlay_config": schema.SingleNestedAttribute{
										Description:         "HybridOverlayConfig configures an additional overlay network for peers that are not using OVN.",
										MarkdownDescription: "HybridOverlayConfig configures an additional overlay network for peers that are not using OVN.",
										Attributes: map[string]schema.Attribute{
											"hybrid_cluster_network": schema.ListNestedAttribute{
												Description:         "HybridClusterNetwork defines a network space given to nodes on an additional overlay network.",
												MarkdownDescription: "HybridClusterNetwork defines a network space given to nodes on an additional overlay network.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"cidr": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"host_prefix": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"hybrid_overlay_vxlan_port": schema.Int64Attribute{
												Description:         "HybridOverlayVXLANPort defines the VXLAN port number to be used by the additional overlay network. Default is 4789",
												MarkdownDescription: "HybridOverlayVXLANPort defines the VXLAN port number to be used by the additional overlay network. Default is 4789",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"ipsec_config": schema.SingleNestedAttribute{
										Description:         "ipsecConfig enables and configures IPsec for pods on the pod network within the cluster.",
										MarkdownDescription: "ipsecConfig enables and configures IPsec for pods on the pod network within the cluster.",
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Description:         "mode defines the behaviour of the ipsec configuration within the platform. Valid values are 'Disabled', 'External' and 'Full'. When 'Disabled', ipsec will not be enabled at the node level. When 'External', ipsec is enabled on the node level but requires the user to configure the secure communication parameters. This mode is for external secure communications and the configuration can be done using the k8s-nmstate operator. When 'Full', ipsec is configured on the node level and inter-pod secure communication within the cluster is configured. Note with 'Full', if ipsec is desired for communication with external (to the cluster) entities (such as storage arrays), this is left to the user to configure.",
												MarkdownDescription: "mode defines the behaviour of the ipsec configuration within the platform. Valid values are 'Disabled', 'External' and 'Full'. When 'Disabled', ipsec will not be enabled at the node level. When 'External', ipsec is enabled on the node level but requires the user to configure the secure communication parameters. This mode is for external secure communications and the configuration can be done using the k8s-nmstate operator. When 'Full', ipsec is configured on the node level and inter-pod secure communication within the cluster is configured. Note with 'Full', if ipsec is desired for communication with external (to the cluster) entities (such as storage arrays), this is left to the user to configure.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Disabled", "External", "Full"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"mtu": schema.Int64Attribute{
										Description:         "mtu is the MTU to use for the tunnel interface. This must be 100 bytes smaller than the uplink mtu. Default is 1400",
										MarkdownDescription: "mtu is the MTU to use for the tunnel interface. This must be 100 bytes smaller than the uplink mtu. Default is 1400",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"policy_audit_config": schema.SingleNestedAttribute{
										Description:         "policyAuditConfig is the configuration for network policy audit events. If unset, reported defaults are used.",
										MarkdownDescription: "policyAuditConfig is the configuration for network policy audit events. If unset, reported defaults are used.",
										Attributes: map[string]schema.Attribute{
											"destination": schema.StringAttribute{
												Description:         "destination is the location for policy log messages. Regardless of this config, persistent logs will always be dumped to the host at /var/log/ovn/ however Additionally syslog output may be configured as follows. Valid values are: - 'libc' -> to use the libc syslog() function of the host node's journdald process - 'udp:host:port' -> for sending syslog over UDP - 'unix:file' -> for using the UNIX domain socket directly - 'null' -> to discard all messages logged to syslog The default is 'null'",
												MarkdownDescription: "destination is the location for policy log messages. Regardless of this config, persistent logs will always be dumped to the host at /var/log/ovn/ however Additionally syslog output may be configured as follows. Valid values are: - 'libc' -> to use the libc syslog() function of the host node's journdald process - 'udp:host:port' -> for sending syslog over UDP - 'unix:file' -> for using the UNIX domain socket directly - 'null' -> to discard all messages logged to syslog The default is 'null'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_file_size": schema.Int64Attribute{
												Description:         "maxFilesSize is the max size an ACL_audit log file is allowed to reach before rotation occurs Units are in MB and the Default is 50MB",
												MarkdownDescription: "maxFilesSize is the max size an ACL_audit log file is allowed to reach before rotation occurs Units are in MB and the Default is 50MB",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"max_log_files": schema.Int64Attribute{
												Description:         "maxLogFiles specifies the maximum number of ACL_audit log files that can be present.",
												MarkdownDescription: "maxLogFiles specifies the maximum number of ACL_audit log files that can be present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"rate_limit": schema.Int64Attribute{
												Description:         "rateLimit is the approximate maximum number of messages to generate per-second per-node. If unset the default of 20 msg/sec is used.",
												MarkdownDescription: "rateLimit is the approximate maximum number of messages to generate per-second per-node. If unset the default of 20 msg/sec is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"syslog_facility": schema.StringAttribute{
												Description:         "syslogFacility the RFC5424 facility for generated messages, e.g. 'kern'. Default is 'local0'",
												MarkdownDescription: "syslogFacility the RFC5424 facility for generated messages, e.g. 'kern'. Default is 'local0'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"v4_internal_subnet": schema.StringAttribute{
										Description:         "v4InternalSubnet is a v4 subnet used internally by ovn-kubernetes in case the default one is being already used by something else. It must not overlap with any other subnet being used by OpenShift or by the node network. The size of the subnet must be larger than the number of nodes. The value cannot be changed after installation. Default is 100.64.0.0/16",
										MarkdownDescription: "v4InternalSubnet is a v4 subnet used internally by ovn-kubernetes in case the default one is being already used by something else. It must not overlap with any other subnet being used by OpenShift or by the node network. The size of the subnet must be larger than the number of nodes. The value cannot be changed after installation. Default is 100.64.0.0/16",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"v6_internal_subnet": schema.StringAttribute{
										Description:         "v6InternalSubnet is a v6 subnet used internally by ovn-kubernetes in case the default one is being already used by something else. It must not overlap with any other subnet being used by OpenShift or by the node network. The size of the subnet must be larger than the number of nodes. The value cannot be changed after installation. Default is fd98::/48",
										MarkdownDescription: "v6InternalSubnet is a v6 subnet used internally by ovn-kubernetes in case the default one is being already used by something else. It must not overlap with any other subnet being used by OpenShift or by the node network. The size of the subnet must be larger than the number of nodes. The value cannot be changed after installation. Default is fd98::/48",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "type is the type of network All NetworkTypes are supported except for NetworkTypeRaw",
								MarkdownDescription: "type is the type of network All NetworkTypes are supported except for NetworkTypeRaw",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"deploy_kube_proxy": schema.BoolAttribute{
						Description:         "deployKubeProxy specifies whether or not a standalone kube-proxy should be deployed by the operator. Some network providers include kube-proxy or similar functionality. If unset, the plugin will attempt to select the correct value, which is false when OpenShift SDN and ovn-kubernetes are used and true otherwise.",
						MarkdownDescription: "deployKubeProxy specifies whether or not a standalone kube-proxy should be deployed by the operator. Some network providers include kube-proxy or similar functionality. If unset, the plugin will attempt to select the correct value, which is false when OpenShift SDN and ovn-kubernetes are used and true otherwise.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_multi_network": schema.BoolAttribute{
						Description:         "disableMultiNetwork specifies whether or not multiple pod network support should be disabled. If unset, this property defaults to 'false' and multiple network support is enabled.",
						MarkdownDescription: "disableMultiNetwork specifies whether or not multiple pod network support should be disabled. If unset, this property defaults to 'false' and multiple network support is enabled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_network_diagnostics": schema.BoolAttribute{
						Description:         "disableNetworkDiagnostics specifies whether or not PodNetworkConnectivityCheck CRs from a test pod to every node, apiserver and LB should be disabled or not. If unset, this property defaults to 'false' and network diagnostics is enabled. Setting this to 'true' would reduce the additional load of the pods performing the checks.",
						MarkdownDescription: "disableNetworkDiagnostics specifies whether or not PodNetworkConnectivityCheck CRs from a test pod to every node, apiserver and LB should be disabled or not. If unset, this property defaults to 'false' and network diagnostics is enabled. Setting this to 'true' would reduce the additional load of the pods performing the checks.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"export_network_flows": schema.SingleNestedAttribute{
						Description:         "exportNetworkFlows enables and configures the export of network flow metadata from the pod network by using protocols NetFlow, SFlow or IPFIX. Currently only supported on OVN-Kubernetes plugin. If unset, flows will not be exported to any collector.",
						MarkdownDescription: "exportNetworkFlows enables and configures the export of network flow metadata from the pod network by using protocols NetFlow, SFlow or IPFIX. Currently only supported on OVN-Kubernetes plugin. If unset, flows will not be exported to any collector.",
						Attributes: map[string]schema.Attribute{
							"ipfix": schema.SingleNestedAttribute{
								Description:         "ipfix defines IPFIX configuration.",
								MarkdownDescription: "ipfix defines IPFIX configuration.",
								Attributes: map[string]schema.Attribute{
									"collectors": schema.ListAttribute{
										Description:         "ipfixCollectors is list of strings formatted as ip:port with a maximum of ten items",
										MarkdownDescription: "ipfixCollectors is list of strings formatted as ip:port with a maximum of ten items",
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

							"net_flow": schema.SingleNestedAttribute{
								Description:         "netFlow defines the NetFlow configuration.",
								MarkdownDescription: "netFlow defines the NetFlow configuration.",
								Attributes: map[string]schema.Attribute{
									"collectors": schema.ListAttribute{
										Description:         "netFlow defines the NetFlow collectors that will consume the flow data exported from OVS. It is a list of strings formatted as ip:port with a maximum of ten items",
										MarkdownDescription: "netFlow defines the NetFlow collectors that will consume the flow data exported from OVS. It is a list of strings formatted as ip:port with a maximum of ten items",
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

							"s_flow": schema.SingleNestedAttribute{
								Description:         "sFlow defines the SFlow configuration.",
								MarkdownDescription: "sFlow defines the SFlow configuration.",
								Attributes: map[string]schema.Attribute{
									"collectors": schema.ListAttribute{
										Description:         "sFlowCollectors is list of strings formatted as ip:port with a maximum of ten items",
										MarkdownDescription: "sFlowCollectors is list of strings formatted as ip:port with a maximum of ten items",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kube_proxy_config": schema.SingleNestedAttribute{
						Description:         "kubeProxyConfig lets us configure desired proxy configuration. If not specified, sensible defaults will be chosen by OpenShift directly. Not consumed by all network providers - currently only openshift-sdn.",
						MarkdownDescription: "kubeProxyConfig lets us configure desired proxy configuration. If not specified, sensible defaults will be chosen by OpenShift directly. Not consumed by all network providers - currently only openshift-sdn.",
						Attributes: map[string]schema.Attribute{
							"bind_address": schema.StringAttribute{
								Description:         "The address to 'bind' on Defaults to 0.0.0.0",
								MarkdownDescription: "The address to 'bind' on Defaults to 0.0.0.0",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"iptables_sync_period": schema.StringAttribute{
								Description:         "An internal kube-proxy parameter. In older releases of OCP, this sometimes needed to be adjusted in large clusters for performance reasons, but this is no longer necessary, and there is no reason to change this from the default value. Default: 30s",
								MarkdownDescription: "An internal kube-proxy parameter. In older releases of OCP, this sometimes needed to be adjusted in large clusters for performance reasons, but this is no longer necessary, and there is no reason to change this from the default value. Default: 30s",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_arguments": schema.MapAttribute{
								Description:         "Any additional arguments to pass to the kubeproxy process",
								MarkdownDescription: "Any additional arguments to pass to the kubeproxy process",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": schema.StringAttribute{
						Description:         "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"management_state": schema.StringAttribute{
						Description:         "managementState indicates whether and how the operator should manage the component",
						MarkdownDescription: "managementState indicates whether and how the operator should manage the component",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(Managed|Unmanaged|Force|Removed)$`), ""),
						},
					},

					"migration": schema.SingleNestedAttribute{
						Description:         "migration enables and configures the cluster network migration. The migration procedure allows to change the network type and the MTU.",
						MarkdownDescription: "migration enables and configures the cluster network migration. The migration procedure allows to change the network type and the MTU.",
						Attributes: map[string]schema.Attribute{
							"features": schema.SingleNestedAttribute{
								Description:         "features contains the features migration configuration. Set this to migrate feature configuration when changing the cluster default network provider. if unset, the default operation is to migrate all the configuration of supported features.",
								MarkdownDescription: "features contains the features migration configuration. Set this to migrate feature configuration when changing the cluster default network provider. if unset, the default operation is to migrate all the configuration of supported features.",
								Attributes: map[string]schema.Attribute{
									"egress_firewall": schema.BoolAttribute{
										Description:         "egressFirewall specifies whether or not the Egress Firewall configuration is migrated automatically when changing the cluster default network provider. If unset, this property defaults to 'true' and Egress Firewall configure is migrated.",
										MarkdownDescription: "egressFirewall specifies whether or not the Egress Firewall configuration is migrated automatically when changing the cluster default network provider. If unset, this property defaults to 'true' and Egress Firewall configure is migrated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"egress_ip": schema.BoolAttribute{
										Description:         "egressIP specifies whether or not the Egress IP configuration is migrated automatically when changing the cluster default network provider. If unset, this property defaults to 'true' and Egress IP configure is migrated.",
										MarkdownDescription: "egressIP specifies whether or not the Egress IP configuration is migrated automatically when changing the cluster default network provider. If unset, this property defaults to 'true' and Egress IP configure is migrated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"multicast": schema.BoolAttribute{
										Description:         "multicast specifies whether or not the multicast configuration is migrated automatically when changing the cluster default network provider. If unset, this property defaults to 'true' and multicast configure is migrated.",
										MarkdownDescription: "multicast specifies whether or not the multicast configuration is migrated automatically when changing the cluster default network provider. If unset, this property defaults to 'true' and multicast configure is migrated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mode": schema.StringAttribute{
								Description:         "mode indicates the mode of network migration. The supported values are 'Live', 'Offline' and omitted. A 'Live' migration operation will not cause service interruption by migrating the CNI of each node one by one. The cluster network will work as normal during the network migration. An 'Offline' migration operation will cause service interruption. During an 'Offline' migration, two rounds of node reboots are required. The cluster network will be malfunctioning during the network migration. When omitted, this means no opinion and the platform is left to choose a reasonable default which is subject to change over time. The current default value is 'Offline'.",
								MarkdownDescription: "mode indicates the mode of network migration. The supported values are 'Live', 'Offline' and omitted. A 'Live' migration operation will not cause service interruption by migrating the CNI of each node one by one. The cluster network will work as normal during the network migration. An 'Offline' migration operation will cause service interruption. During an 'Offline' migration, two rounds of node reboots are required. The cluster network will be malfunctioning during the network migration. When omitted, this means no opinion and the platform is left to choose a reasonable default which is subject to change over time. The current default value is 'Offline'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Live", "Offline", ""),
								},
							},

							"mtu": schema.SingleNestedAttribute{
								Description:         "mtu contains the MTU migration configuration. Set this to allow changing the MTU values for the default network. If unset, the operation of changing the MTU for the default network will be rejected.",
								MarkdownDescription: "mtu contains the MTU migration configuration. Set this to allow changing the MTU values for the default network. If unset, the operation of changing the MTU for the default network will be rejected.",
								Attributes: map[string]schema.Attribute{
									"machine": schema.SingleNestedAttribute{
										Description:         "machine contains MTU migration configuration for the machine's uplink. Needs to be migrated along with the default network MTU unless the current uplink MTU already accommodates the default network MTU.",
										MarkdownDescription: "machine contains MTU migration configuration for the machine's uplink. Needs to be migrated along with the default network MTU unless the current uplink MTU already accommodates the default network MTU.",
										Attributes: map[string]schema.Attribute{
											"from": schema.Int64Attribute{
												Description:         "from is the MTU to migrate from.",
												MarkdownDescription: "from is the MTU to migrate from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"to": schema.Int64Attribute{
												Description:         "to is the MTU to migrate to.",
												MarkdownDescription: "to is the MTU to migrate to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network": schema.SingleNestedAttribute{
										Description:         "network contains information about MTU migration for the default network. Migrations are only allowed to MTU values lower than the machine's uplink MTU by the minimum appropriate offset.",
										MarkdownDescription: "network contains information about MTU migration for the default network. Migrations are only allowed to MTU values lower than the machine's uplink MTU by the minimum appropriate offset.",
										Attributes: map[string]schema.Attribute{
											"from": schema.Int64Attribute{
												Description:         "from is the MTU to migrate from.",
												MarkdownDescription: "from is the MTU to migrate from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"to": schema.Int64Attribute{
												Description:         "to is the MTU to migrate to.",
												MarkdownDescription: "to is the MTU to migrate to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
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

							"network_type": schema.StringAttribute{
								Description:         "networkType is the target type of network migration. Set this to the target network type to allow changing the default network. If unset, the operation of changing cluster default network plugin will be rejected. The supported values are OpenShiftSDN, OVNKubernetes",
								MarkdownDescription: "networkType is the target type of network migration. Set this to the target network type to allow changing the default network. If unset, the operation of changing cluster default network plugin will be rejected. The supported values are OpenShiftSDN, OVNKubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"observed_config": schema.MapAttribute{
						Description:         "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						MarkdownDescription: "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operator_log_level": schema.StringAttribute{
						Description:         "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"service_network": schema.ListAttribute{
						Description:         "serviceNetwork is the ip address pool to use for Service IPs Currently, all existing network providers only support a single value here, but this is an array to allow for growth.",
						MarkdownDescription: "serviceNetwork is the ip address pool to use for Service IPs Currently, all existing network providers only support a single value here, but this is an array to allow for growth.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unsupported_config_overrides": schema.MapAttribute{
						Description:         "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						MarkdownDescription: "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_multi_network_policy": schema.BoolAttribute{
						Description:         "useMultiNetworkPolicy enables a controller which allows for MultiNetworkPolicy objects to be used on additional networks as created by Multus CNI. MultiNetworkPolicy are similar to NetworkPolicy objects, but NetworkPolicy objects only apply to the primary interface. With MultiNetworkPolicy, you can control the traffic that a pod can receive over the secondary interfaces. If unset, this property defaults to 'false' and MultiNetworkPolicy objects are ignored. If 'disableMultiNetwork' is 'true' then the value of this field is ignored.",
						MarkdownDescription: "useMultiNetworkPolicy enables a controller which allows for MultiNetworkPolicy objects to be used on additional networks as created by Multus CNI. MultiNetworkPolicy are similar to NetworkPolicy objects, but NetworkPolicy objects only apply to the primary interface. With MultiNetworkPolicy, you can control the traffic that a pod can receive over the secondary interfaces. If unset, this property defaults to 'false' and MultiNetworkPolicy objects are ignored. If 'disableMultiNetwork' is 'true' then the value of this field is ignored.",
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

func (r *OperatorOpenshiftIoNetworkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_openshift_io_network_v1_manifest")

	var model OperatorOpenshiftIoNetworkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.openshift.io/v1")
	model.Kind = pointer.String("Network")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
