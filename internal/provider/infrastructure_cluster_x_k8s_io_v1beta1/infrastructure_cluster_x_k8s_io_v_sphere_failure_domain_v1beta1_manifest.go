/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

import (
	"context"
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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Region *struct {
			AutoConfigure *bool   `tfsdk:"auto_configure" json:"autoConfigure,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			TagCategory   *string `tfsdk:"tag_category" json:"tagCategory,omitempty"`
			Type          *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"region" json:"region,omitempty"`
		Topology *struct {
			ComputeCluster *string `tfsdk:"compute_cluster" json:"computeCluster,omitempty"`
			Datacenter     *string `tfsdk:"datacenter" json:"datacenter,omitempty"`
			Datastore      *string `tfsdk:"datastore" json:"datastore,omitempty"`
			Hosts          *struct {
				HostGroupName *string `tfsdk:"host_group_name" json:"hostGroupName,omitempty"`
				VmGroupName   *string `tfsdk:"vm_group_name" json:"vmGroupName,omitempty"`
			} `tfsdk:"hosts" json:"hosts,omitempty"`
			NetworkConfigurations *[]struct {
				AddressesFromPools *[]struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"addresses_from_pools" json:"addressesFromPools,omitempty"`
				Dhcp4          *bool `tfsdk:"dhcp4" json:"dhcp4,omitempty"`
				Dhcp4Overrides *struct {
					Hostname     *string `tfsdk:"hostname" json:"hostname,omitempty"`
					RouteMetric  *int64  `tfsdk:"route_metric" json:"routeMetric,omitempty"`
					SendHostname *bool   `tfsdk:"send_hostname" json:"sendHostname,omitempty"`
					UseDNS       *bool   `tfsdk:"use_dns" json:"useDNS,omitempty"`
					UseDomains   *string `tfsdk:"use_domains" json:"useDomains,omitempty"`
					UseHostname  *bool   `tfsdk:"use_hostname" json:"useHostname,omitempty"`
					UseMTU       *bool   `tfsdk:"use_mtu" json:"useMTU,omitempty"`
					UseNTP       *bool   `tfsdk:"use_ntp" json:"useNTP,omitempty"`
					UseRoutes    *string `tfsdk:"use_routes" json:"useRoutes,omitempty"`
				} `tfsdk:"dhcp4_overrides" json:"dhcp4Overrides,omitempty"`
				Dhcp6          *bool `tfsdk:"dhcp6" json:"dhcp6,omitempty"`
				Dhcp6Overrides *struct {
					Hostname     *string `tfsdk:"hostname" json:"hostname,omitempty"`
					RouteMetric  *int64  `tfsdk:"route_metric" json:"routeMetric,omitempty"`
					SendHostname *bool   `tfsdk:"send_hostname" json:"sendHostname,omitempty"`
					UseDNS       *bool   `tfsdk:"use_dns" json:"useDNS,omitempty"`
					UseDomains   *string `tfsdk:"use_domains" json:"useDomains,omitempty"`
					UseHostname  *bool   `tfsdk:"use_hostname" json:"useHostname,omitempty"`
					UseMTU       *bool   `tfsdk:"use_mtu" json:"useMTU,omitempty"`
					UseNTP       *bool   `tfsdk:"use_ntp" json:"useNTP,omitempty"`
					UseRoutes    *string `tfsdk:"use_routes" json:"useRoutes,omitempty"`
				} `tfsdk:"dhcp6_overrides" json:"dhcp6Overrides,omitempty"`
				Nameservers   *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				NetworkName   *string   `tfsdk:"network_name" json:"networkName,omitempty"`
				SearchDomains *[]string `tfsdk:"search_domains" json:"searchDomains,omitempty"`
			} `tfsdk:"network_configurations" json:"networkConfigurations,omitempty"`
			Networks *[]string `tfsdk:"networks" json:"networks,omitempty"`
		} `tfsdk:"topology" json:"topology,omitempty"`
		Zone *struct {
			AutoConfigure *bool   `tfsdk:"auto_configure" json:"autoConfigure,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			TagCategory   *string `tfsdk:"tag_category" json:"tagCategory,omitempty"`
			Type          *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereFailureDomain is the Schema for the vspherefailuredomains API.",
		MarkdownDescription: "VSphereFailureDomain is the Schema for the vspherefailuredomains API.",
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
				Description:         "VSphereFailureDomainSpec defines the desired state of VSphereFailureDomain.",
				MarkdownDescription: "VSphereFailureDomainSpec defines the desired state of VSphereFailureDomain.",
				Attributes: map[string]schema.Attribute{
					"region": schema.SingleNestedAttribute{
						Description:         "Region defines the name and type of a region",
						MarkdownDescription: "Region defines the name and type of a region",
						Attributes: map[string]schema.Attribute{
							"auto_configure": schema.BoolAttribute{
								Description:         "AutoConfigure tags the Type which is specified in the Topology Deprecated: This field is going to be removed in a future release.",
								MarkdownDescription: "AutoConfigure tags the Type which is specified in the Topology Deprecated: This field is going to be removed in a future release.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the tag that represents this failure domain",
								MarkdownDescription: "Name is the name of the tag that represents this failure domain",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tag_category": schema.StringAttribute{
								Description:         "TagCategory is the category used for the tag",
								MarkdownDescription: "TagCategory is the category used for the tag",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								MarkdownDescription: "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Datacenter", "ComputeCluster", "HostGroup"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"topology": schema.SingleNestedAttribute{
						Description:         "Topology describes a given failure domain using vSphere constructs",
						MarkdownDescription: "Topology describes a given failure domain using vSphere constructs",
						Attributes: map[string]schema.Attribute{
							"compute_cluster": schema.StringAttribute{
								Description:         "ComputeCluster as the failure domain",
								MarkdownDescription: "ComputeCluster as the failure domain",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"datacenter": schema.StringAttribute{
								Description:         "Datacenter as the failure domain.",
								MarkdownDescription: "Datacenter as the failure domain.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"datastore": schema.StringAttribute{
								Description:         "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hosts": schema.SingleNestedAttribute{
								Description:         "Hosts has information required for placement of machines on VSphere hosts.",
								MarkdownDescription: "Hosts has information required for placement of machines on VSphere hosts.",
								Attributes: map[string]schema.Attribute{
									"host_group_name": schema.StringAttribute{
										Description:         "HostGroupName is the name of the Host group",
										MarkdownDescription: "HostGroupName is the name of the Host group",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"vm_group_name": schema.StringAttribute{
										Description:         "VMGroupName is the name of the VM group",
										MarkdownDescription: "VMGroupName is the name of the VM group",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"network_configurations": schema.ListNestedAttribute{
								Description:         "NetworkConfigurations is a list of network configurations within this failure domain.",
								MarkdownDescription: "NetworkConfigurations is a list of network configurations within this failure domain.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"addresses_from_pools": schema.ListNestedAttribute{
											Description:         "AddressesFromPools is a list of IPAddressPools that should be assigned to IPAddressClaims. The machine's cloud-init metadata will be populated with IPAddresses fulfilled by an IPAM provider.",
											MarkdownDescription: "AddressesFromPools is a list of IPAddressPools that should be assigned to IPAddressClaims. The machine's cloud-init metadata will be populated with IPAddresses fulfilled by an IPAM provider.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
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

										"dhcp4": schema.BoolAttribute{
											Description:         "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4.",
											MarkdownDescription: "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dhcp4_overrides": schema.SingleNestedAttribute{
											Description:         "DHCP4Overrides allows for the control over several DHCP behaviors. Overrides will only be applied when the corresponding DHCP flag is set. Only configured values will be sent, omitted values will default to distribution defaults. Dependent on support in the network stack for your distribution. For more information see the netplan reference (https://netplan.io/reference#dhcp-overrides)",
											MarkdownDescription: "DHCP4Overrides allows for the control over several DHCP behaviors. Overrides will only be applied when the corresponding DHCP flag is set. Only configured values will be sent, omitted values will default to distribution defaults. Dependent on support in the network stack for your distribution. For more information see the netplan reference (https://netplan.io/reference#dhcp-overrides)",
											Attributes: map[string]schema.Attribute{
												"hostname": schema.StringAttribute{
													Description:         "Hostname is the name which will be sent to the DHCP server instead of the machine's hostname.",
													MarkdownDescription: "Hostname is the name which will be sent to the DHCP server instead of the machine's hostname.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"route_metric": schema.Int64Attribute{
													Description:         "RouteMetric is used to prioritize routes for devices. A lower metric for an interface will have a higher priority.",
													MarkdownDescription: "RouteMetric is used to prioritize routes for devices. A lower metric for an interface will have a higher priority.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"send_hostname": schema.BoolAttribute{
													Description:         "SendHostname when 'true', the hostname of the machine will be sent to the DHCP server.",
													MarkdownDescription: "SendHostname when 'true', the hostname of the machine will be sent to the DHCP server.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_dns": schema.BoolAttribute{
													Description:         "UseDNS when 'true', the DNS servers in the DHCP server will be used and take precedence.",
													MarkdownDescription: "UseDNS when 'true', the DNS servers in the DHCP server will be used and take precedence.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_domains": schema.StringAttribute{
													Description:         "UseDomains can take the values 'true', 'false', or 'route'. When 'true', the domain name from the DHCP server will be used as the DNS search domain for this device. When 'route', the domain name from the DHCP response will be used for routing DNS only, not for searching.",
													MarkdownDescription: "UseDomains can take the values 'true', 'false', or 'route'. When 'true', the domain name from the DHCP server will be used as the DNS search domain for this device. When 'route', the domain name from the DHCP response will be used for routing DNS only, not for searching.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_hostname": schema.BoolAttribute{
													Description:         "UseHostname when 'true', the hostname from the DHCP server will be set as the transient hostname of the machine.",
													MarkdownDescription: "UseHostname when 'true', the hostname from the DHCP server will be set as the transient hostname of the machine.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_mtu": schema.BoolAttribute{
													Description:         "UseMTU when 'true', the MTU from the DHCP server will be set as the MTU of the device.",
													MarkdownDescription: "UseMTU when 'true', the MTU from the DHCP server will be set as the MTU of the device.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_ntp": schema.BoolAttribute{
													Description:         "UseNTP when 'true', the NTP servers from the DHCP server will be used by systemd-timesyncd and take precedence.",
													MarkdownDescription: "UseNTP when 'true', the NTP servers from the DHCP server will be used by systemd-timesyncd and take precedence.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_routes": schema.StringAttribute{
													Description:         "UseRoutes when 'true', the routes from the DHCP server will be installed in the routing table.",
													MarkdownDescription: "UseRoutes when 'true', the routes from the DHCP server will be installed in the routing table.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"dhcp6": schema.BoolAttribute{
											Description:         "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6.",
											MarkdownDescription: "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dhcp6_overrides": schema.SingleNestedAttribute{
											Description:         "DHCP6Overrides allows for the control over several DHCP behaviors. Overrides will only be applied when the corresponding DHCP flag is set. Only configured values will be sent, omitted values will default to distribution defaults. Dependent on support in the network stack for your distribution. For more information see the netplan reference (https://netplan.io/reference#dhcp-overrides)",
											MarkdownDescription: "DHCP6Overrides allows for the control over several DHCP behaviors. Overrides will only be applied when the corresponding DHCP flag is set. Only configured values will be sent, omitted values will default to distribution defaults. Dependent on support in the network stack for your distribution. For more information see the netplan reference (https://netplan.io/reference#dhcp-overrides)",
											Attributes: map[string]schema.Attribute{
												"hostname": schema.StringAttribute{
													Description:         "Hostname is the name which will be sent to the DHCP server instead of the machine's hostname.",
													MarkdownDescription: "Hostname is the name which will be sent to the DHCP server instead of the machine's hostname.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"route_metric": schema.Int64Attribute{
													Description:         "RouteMetric is used to prioritize routes for devices. A lower metric for an interface will have a higher priority.",
													MarkdownDescription: "RouteMetric is used to prioritize routes for devices. A lower metric for an interface will have a higher priority.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"send_hostname": schema.BoolAttribute{
													Description:         "SendHostname when 'true', the hostname of the machine will be sent to the DHCP server.",
													MarkdownDescription: "SendHostname when 'true', the hostname of the machine will be sent to the DHCP server.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_dns": schema.BoolAttribute{
													Description:         "UseDNS when 'true', the DNS servers in the DHCP server will be used and take precedence.",
													MarkdownDescription: "UseDNS when 'true', the DNS servers in the DHCP server will be used and take precedence.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_domains": schema.StringAttribute{
													Description:         "UseDomains can take the values 'true', 'false', or 'route'. When 'true', the domain name from the DHCP server will be used as the DNS search domain for this device. When 'route', the domain name from the DHCP response will be used for routing DNS only, not for searching.",
													MarkdownDescription: "UseDomains can take the values 'true', 'false', or 'route'. When 'true', the domain name from the DHCP server will be used as the DNS search domain for this device. When 'route', the domain name from the DHCP response will be used for routing DNS only, not for searching.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_hostname": schema.BoolAttribute{
													Description:         "UseHostname when 'true', the hostname from the DHCP server will be set as the transient hostname of the machine.",
													MarkdownDescription: "UseHostname when 'true', the hostname from the DHCP server will be set as the transient hostname of the machine.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_mtu": schema.BoolAttribute{
													Description:         "UseMTU when 'true', the MTU from the DHCP server will be set as the MTU of the device.",
													MarkdownDescription: "UseMTU when 'true', the MTU from the DHCP server will be set as the MTU of the device.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_ntp": schema.BoolAttribute{
													Description:         "UseNTP when 'true', the NTP servers from the DHCP server will be used by systemd-timesyncd and take precedence.",
													MarkdownDescription: "UseNTP when 'true', the NTP servers from the DHCP server will be used by systemd-timesyncd and take precedence.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"use_routes": schema.StringAttribute{
													Description:         "UseRoutes when 'true', the routes from the DHCP server will be installed in the routing table.",
													MarkdownDescription: "UseRoutes when 'true', the routes from the DHCP server will be installed in the routing table.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"nameservers": schema.ListAttribute{
											Description:         "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
											MarkdownDescription: "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"network_name": schema.StringAttribute{
											Description:         "NetworkName is the network name for this machine's VM.",
											MarkdownDescription: "NetworkName is the network name for this machine's VM.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"search_domains": schema.ListAttribute{
											Description:         "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
											MarkdownDescription: "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
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

							"networks": schema.ListAttribute{
								Description:         "Networks is the list of networks within this failure domain",
								MarkdownDescription: "Networks is the list of networks within this failure domain",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"zone": schema.SingleNestedAttribute{
						Description:         "Zone defines the name and type of a zone",
						MarkdownDescription: "Zone defines the name and type of a zone",
						Attributes: map[string]schema.Attribute{
							"auto_configure": schema.BoolAttribute{
								Description:         "AutoConfigure tags the Type which is specified in the Topology Deprecated: This field is going to be removed in a future release.",
								MarkdownDescription: "AutoConfigure tags the Type which is specified in the Topology Deprecated: This field is going to be removed in a future release.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the tag that represents this failure domain",
								MarkdownDescription: "Name is the name of the tag that represents this failure domain",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tag_category": schema.StringAttribute{
								Description:         "TagCategory is the category used for the tag",
								MarkdownDescription: "TagCategory is the category used for the tag",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								MarkdownDescription: "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Datacenter", "ComputeCluster", "HostGroup"),
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

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("VSphereFailureDomain")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
