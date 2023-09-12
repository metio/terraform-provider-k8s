/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha3

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
	_ datasource.DataSource              = &InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource{}
	_ datasource.DataSourceWithConfigure = &InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource{}
)

func NewInfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource() datasource.DataSource {
	return &InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource{}
}

type InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		User *struct {
			AuthorizedKeys *[]string `tfsdk:"authorized_keys" json:"authorizedKeys,omitempty"`
			Name           *string   `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"user" json:"user,omitempty"`
		VirtualMachineConfiguration *struct {
			CloneMode     *string            `tfsdk:"clone_mode" json:"cloneMode,omitempty"`
			CustomVMXKeys *map[string]string `tfsdk:"custom_vmx_keys" json:"customVMXKeys,omitempty"`
			Datacenter    *string            `tfsdk:"datacenter" json:"datacenter,omitempty"`
			Datastore     *string            `tfsdk:"datastore" json:"datastore,omitempty"`
			DiskGiB       *int64             `tfsdk:"disk_gi_b" json:"diskGiB,omitempty"`
			Folder        *string            `tfsdk:"folder" json:"folder,omitempty"`
			MemoryMiB     *int64             `tfsdk:"memory_mi_b" json:"memoryMiB,omitempty"`
			Network       *struct {
				Devices *[]struct {
					DeviceName  *string   `tfsdk:"device_name" json:"deviceName,omitempty"`
					Dhcp4       *bool     `tfsdk:"dhcp4" json:"dhcp4,omitempty"`
					Dhcp6       *bool     `tfsdk:"dhcp6" json:"dhcp6,omitempty"`
					Gateway4    *string   `tfsdk:"gateway4" json:"gateway4,omitempty"`
					Gateway6    *string   `tfsdk:"gateway6" json:"gateway6,omitempty"`
					IpAddrs     *[]string `tfsdk:"ip_addrs" json:"ipAddrs,omitempty"`
					MacAddr     *string   `tfsdk:"mac_addr" json:"macAddr,omitempty"`
					Mtu         *int64    `tfsdk:"mtu" json:"mtu,omitempty"`
					Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
					NetworkName *string   `tfsdk:"network_name" json:"networkName,omitempty"`
					Routes      *[]struct {
						Metric *int64  `tfsdk:"metric" json:"metric,omitempty"`
						To     *string `tfsdk:"to" json:"to,omitempty"`
						Via    *string `tfsdk:"via" json:"via,omitempty"`
					} `tfsdk:"routes" json:"routes,omitempty"`
					SearchDomains *[]string `tfsdk:"search_domains" json:"searchDomains,omitempty"`
				} `tfsdk:"devices" json:"devices,omitempty"`
				PreferredAPIServerCidr *string `tfsdk:"preferred_api_server_cidr" json:"preferredAPIServerCidr,omitempty"`
				Routes                 *[]struct {
					Metric *int64  `tfsdk:"metric" json:"metric,omitempty"`
					To     *string `tfsdk:"to" json:"to,omitempty"`
					Via    *string `tfsdk:"via" json:"via,omitempty"`
				} `tfsdk:"routes" json:"routes,omitempty"`
			} `tfsdk:"network" json:"network,omitempty"`
			NumCPUs           *int64  `tfsdk:"num_cp_us" json:"numCPUs,omitempty"`
			NumCoresPerSocket *int64  `tfsdk:"num_cores_per_socket" json:"numCoresPerSocket,omitempty"`
			ResourcePool      *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
			Server            *string `tfsdk:"server" json:"server,omitempty"`
			Snapshot          *string `tfsdk:"snapshot" json:"snapshot,omitempty"`
			StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
			Template          *string `tfsdk:"template" json:"template,omitempty"`
			Thumbprint        *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
		} `tfsdk:"virtual_machine_configuration" json:"virtualMachineConfiguration,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ha_proxy_load_balancer_v1alpha3"
}

func (r *InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HAProxyLoadBalancer is the Schema for the haproxyloadbalancers API  Deprecated: This type will be removed in v1alpha4.",
		MarkdownDescription: "HAProxyLoadBalancer is the Schema for the haproxyloadbalancers API  Deprecated: This type will be removed in v1alpha4.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "HAProxyLoadBalancerSpec defines the desired state of HAProxyLoadBalancer.",
				MarkdownDescription: "HAProxyLoadBalancerSpec defines the desired state of HAProxyLoadBalancer.",
				Attributes: map[string]schema.Attribute{
					"user": schema.SingleNestedAttribute{
						Description:         "SSHUser specifies the name of a user that is granted remote access to the deployed VM.",
						MarkdownDescription: "SSHUser specifies the name of a user that is granted remote access to the deployed VM.",
						Attributes: map[string]schema.Attribute{
							"authorized_keys": schema.ListAttribute{
								Description:         "AuthorizedKeys is one or more public SSH keys that grant remote access.",
								MarkdownDescription: "AuthorizedKeys is one or more public SSH keys that grant remote access.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the SSH user.",
								MarkdownDescription: "Name is the name of the SSH user.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"virtual_machine_configuration": schema.SingleNestedAttribute{
						Description:         "VirtualMachineConfiguration is information used to deploy a load balancer VM.",
						MarkdownDescription: "VirtualMachineConfiguration is information used to deploy a load balancer VM.",
						Attributes: map[string]schema.Attribute{
							"clone_mode": schema.StringAttribute{
								Description:         "CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.",
								MarkdownDescription: "CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"custom_vmx_keys": schema.MapAttribute{
								Description:         "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map",
								MarkdownDescription: "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"datacenter": schema.StringAttribute{
								Description:         "Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located.",
								MarkdownDescription: "Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"datastore": schema.StringAttribute{
								Description:         "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"disk_gi_b": schema.Int64Attribute{
								Description:         "DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								MarkdownDescription: "DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"folder": schema.StringAttribute{
								Description:         "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
								MarkdownDescription: "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"memory_mi_b": schema.Int64Attribute{
								Description:         "MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								MarkdownDescription: "MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network": schema.SingleNestedAttribute{
								Description:         "Network is the network configuration for this machine's VM.",
								MarkdownDescription: "Network is the network configuration for this machine's VM.",
								Attributes: map[string]schema.Attribute{
									"devices": schema.ListNestedAttribute{
										Description:         "Devices is the list of network devices used by the virtual machine. TODO(akutz) Make sure at least one network matches the ClusterSpec.CloudProviderConfiguration.Network.Name",
										MarkdownDescription: "Devices is the list of network devices used by the virtual machine. TODO(akutz) Make sure at least one network matches the ClusterSpec.CloudProviderConfiguration.Network.Name",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"device_name": schema.StringAttribute{
													Description:         "DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.",
													MarkdownDescription: "DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"dhcp4": schema.BoolAttribute{
													Description:         "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.",
													MarkdownDescription: "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"dhcp6": schema.BoolAttribute{
													Description:         "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.",
													MarkdownDescription: "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"gateway4": schema.StringAttribute{
													Description:         "Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.",
													MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"gateway6": schema.StringAttribute{
													Description:         "Gateway4 is the IPv4 gateway used by this device. Required when DHCP6 is false.",
													MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device. Required when DHCP6 is false.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ip_addrs": schema.ListAttribute{
													Description:         "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4 and DHCP6 are both false.",
													MarkdownDescription: "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4 and DHCP6 are both false.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"mac_addr": schema.StringAttribute{
													Description:         "MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.",
													MarkdownDescription: "MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"mtu": schema.Int64Attribute{
													Description:         "MTU is the device’s Maximum Transmission Unit size in bytes.",
													MarkdownDescription: "MTU is the device’s Maximum Transmission Unit size in bytes.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"nameservers": schema.ListAttribute{
													Description:         "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
													MarkdownDescription: "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"network_name": schema.StringAttribute{
													Description:         "NetworkName is the name of the vSphere network to which the device will be connected.",
													MarkdownDescription: "NetworkName is the name of the vSphere network to which the device will be connected.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"routes": schema.ListNestedAttribute{
													Description:         "Routes is a list of optional, static routes applied to the device.",
													MarkdownDescription: "Routes is a list of optional, static routes applied to the device.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"metric": schema.Int64Attribute{
																Description:         "Metric is the weight/priority of the route.",
																MarkdownDescription: "Metric is the weight/priority of the route.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"to": schema.StringAttribute{
																Description:         "To is an IPv4 or IPv6 address.",
																MarkdownDescription: "To is an IPv4 or IPv6 address.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"via": schema.StringAttribute{
																Description:         "Via is an IPv4 or IPv6 address.",
																MarkdownDescription: "Via is an IPv4 or IPv6 address.",
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

												"search_domains": schema.ListAttribute{
													Description:         "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
													MarkdownDescription: "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
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

									"preferred_api_server_cidr": schema.StringAttribute{
										Description:         "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine",
										MarkdownDescription: "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"routes": schema.ListNestedAttribute{
										Description:         "Routes is a list of optional, static routes applied to the virtual machine.",
										MarkdownDescription: "Routes is a list of optional, static routes applied to the virtual machine.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"metric": schema.Int64Attribute{
													Description:         "Metric is the weight/priority of the route.",
													MarkdownDescription: "Metric is the weight/priority of the route.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"to": schema.StringAttribute{
													Description:         "To is an IPv4 or IPv6 address.",
													MarkdownDescription: "To is an IPv4 or IPv6 address.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"via": schema.StringAttribute{
													Description:         "Via is an IPv4 or IPv6 address.",
													MarkdownDescription: "Via is an IPv4 or IPv6 address.",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"num_cp_us": schema.Int64Attribute{
								Description:         "NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								MarkdownDescription: "NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"num_cores_per_socket": schema.Int64Attribute{
								Description:         "NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								MarkdownDescription: "NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resource_pool": schema.StringAttribute{
								Description:         "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
								MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"server": schema.StringAttribute{
								Description:         "Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.",
								MarkdownDescription: "Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"snapshot": schema.StringAttribute{
								Description:         "Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.",
								MarkdownDescription: "Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"storage_policy_name": schema.StringAttribute{
								Description:         "StoragePolicyName of the storage policy to use with this Virtual Machine",
								MarkdownDescription: "StoragePolicyName of the storage policy to use with this Virtual Machine",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"template": schema.StringAttribute{
								Description:         "Template is the name or inventory path of the template used to clone the virtual machine.",
								MarkdownDescription: "Template is the name or inventory path of the template used to clone the virtual machine.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"thumbprint": schema.StringAttribute{
								Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.",
								MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.",
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
		},
	}
}

func (r *InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infrastructure_cluster_x_k8s_io_ha_proxy_load_balancer_v1alpha3")

	var data InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "haproxyloadbalancers"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoHaproxyLoadBalancerV1Alpha3DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha3")
	data.Kind = pointer.String("HAProxyLoadBalancer")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
