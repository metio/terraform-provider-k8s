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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1ManifestData struct {
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
		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				AdditionalDisksGiB       *[]string          `tfsdk:"additional_disks_gi_b" json:"additionalDisksGiB,omitempty"`
				CloneMode                *string            `tfsdk:"clone_mode" json:"cloneMode,omitempty"`
				CustomVMXKeys            *map[string]string `tfsdk:"custom_vmx_keys" json:"customVMXKeys,omitempty"`
				Datacenter               *string            `tfsdk:"datacenter" json:"datacenter,omitempty"`
				Datastore                *string            `tfsdk:"datastore" json:"datastore,omitempty"`
				DiskGiB                  *int64             `tfsdk:"disk_gi_b" json:"diskGiB,omitempty"`
				FailureDomain            *string            `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
				Folder                   *string            `tfsdk:"folder" json:"folder,omitempty"`
				GuestSoftPowerOffTimeout *string            `tfsdk:"guest_soft_power_off_timeout" json:"guestSoftPowerOffTimeout,omitempty"`
				HardwareVersion          *string            `tfsdk:"hardware_version" json:"hardwareVersion,omitempty"`
				MemoryMiB                *int64             `tfsdk:"memory_mi_b" json:"memoryMiB,omitempty"`
				Network                  *struct {
					Devices *[]struct {
						AddressesFromPools *[]struct {
							ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
							Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"addresses_from_pools" json:"addressesFromPools,omitempty"`
						DeviceName     *string `tfsdk:"device_name" json:"deviceName,omitempty"`
						Dhcp4          *bool   `tfsdk:"dhcp4" json:"dhcp4,omitempty"`
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
						SearchDomains    *[]string `tfsdk:"search_domains" json:"searchDomains,omitempty"`
						SkipIPAllocation *bool     `tfsdk:"skip_ip_allocation" json:"skipIPAllocation,omitempty"`
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
				Os                *string `tfsdk:"os" json:"os,omitempty"`
				PciDevices        *[]struct {
					CustomLabel *string `tfsdk:"custom_label" json:"customLabel,omitempty"`
					DeviceId    *int64  `tfsdk:"device_id" json:"deviceId,omitempty"`
					VGPUProfile *string `tfsdk:"v_gpu_profile" json:"vGPUProfile,omitempty"`
					VendorId    *int64  `tfsdk:"vendor_id" json:"vendorId,omitempty"`
				} `tfsdk:"pci_devices" json:"pciDevices,omitempty"`
				PowerOffMode      *string   `tfsdk:"power_off_mode" json:"powerOffMode,omitempty"`
				ProviderID        *string   `tfsdk:"provider_id" json:"providerID,omitempty"`
				ResourcePool      *string   `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
				Server            *string   `tfsdk:"server" json:"server,omitempty"`
				Snapshot          *string   `tfsdk:"snapshot" json:"snapshot,omitempty"`
				StoragePolicyName *string   `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
				TagIDs            *[]string `tfsdk:"tag_i_ds" json:"tagIDs,omitempty"`
				Template          *string   `tfsdk:"template" json:"template,omitempty"`
				Thumbprint        *string   `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereMachineTemplate is the Schema for the vspheremachinetemplates API.",
		MarkdownDescription: "VSphereMachineTemplate is the Schema for the vspheremachinetemplates API.",
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
				Description:         "VSphereMachineTemplateSpec defines the desired state of VSphereMachineTemplate.",
				MarkdownDescription: "VSphereMachineTemplateSpec defines the desired state of VSphereMachineTemplate.",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "VSphereMachineTemplateResource describes the data needed to create a VSphereMachine from a template.",
						MarkdownDescription: "VSphereMachineTemplateResource describes the data needed to create a VSphereMachine from a template.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
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

							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the desired behavior of the machine.",
								MarkdownDescription: "Spec is the specification of the desired behavior of the machine.",
								Attributes: map[string]schema.Attribute{
									"additional_disks_gi_b": schema.ListAttribute{
										Description:         "AdditionalDisksGiB holds the sizes of additional disks of the virtual machine, in GiB Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										MarkdownDescription: "AdditionalDisksGiB holds the sizes of additional disks of the virtual machine, in GiB Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"clone_mode": schema.StringAttribute{
										Description:         "CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.",
										MarkdownDescription: "CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"custom_vmx_keys": schema.MapAttribute{
										Description:         "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map",
										MarkdownDescription: "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"datacenter": schema.StringAttribute{
										Description:         "Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located. Defaults to * which selects the default datacenter.",
										MarkdownDescription: "Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located. Defaults to * which selects the default datacenter.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"datastore": schema.StringAttribute{
										Description:         "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
										MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disk_gi_b": schema.Int64Attribute{
										Description:         "DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										MarkdownDescription: "DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"failure_domain": schema.StringAttribute{
										Description:         "FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. For this infrastructure provider, the name is equivalent to the name of the VSphereDeploymentZone.",
										MarkdownDescription: "FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. For this infrastructure provider, the name is equivalent to the name of the VSphereDeploymentZone.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"folder": schema.StringAttribute{
										Description:         "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
										MarkdownDescription: "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"guest_soft_power_off_timeout": schema.StringAttribute{
										Description:         "GuestSoftPowerOffTimeout sets the wait timeout for shutdown in the VM guest. The VM will be powered off forcibly after the timeout if the VM is still up and running when the PowerOffMode is set to trySoft. This parameter only applies when the PowerOffMode is set to trySoft. If omitted, the timeout defaults to 5 minutes.",
										MarkdownDescription: "GuestSoftPowerOffTimeout sets the wait timeout for shutdown in the VM guest. The VM will be powered off forcibly after the timeout if the VM is still up and running when the PowerOffMode is set to trySoft. This parameter only applies when the PowerOffMode is set to trySoft. If omitted, the timeout defaults to 5 minutes.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hardware_version": schema.StringAttribute{
										Description:         "HardwareVersion is the hardware version of the virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Check the compatibility with the ESXi version before setting the value.",
										MarkdownDescription: "HardwareVersion is the hardware version of the virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Check the compatibility with the ESXi version before setting the value.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory_mi_b": schema.Int64Attribute{
										Description:         "MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										MarkdownDescription: "MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										Required:            false,
										Optional:            true,
										Computed:            false,
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

														"device_name": schema.StringAttribute{
															Description:         "DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.",
															MarkdownDescription: "DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dhcp4": schema.BoolAttribute{
															Description:         "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.",
															MarkdownDescription: "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.",
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
															Description:         "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.",
															MarkdownDescription: "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.",
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

														"gateway4": schema.StringAttribute{
															Description:         "Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.",
															MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"gateway6": schema.StringAttribute{
															Description:         "Gateway4 is the IPv4 gateway used by this device.",
															MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip_addrs": schema.ListAttribute{
															Description:         "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4, DHCP6 and SkipIPAllocation are false.",
															MarkdownDescription: "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4, DHCP6 and SkipIPAllocation are false.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mac_addr": schema.StringAttribute{
															Description:         "MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.",
															MarkdownDescription: "MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mtu": schema.Int64Attribute{
															Description:         "MTU is the device’s Maximum Transmission Unit size in bytes.",
															MarkdownDescription: "MTU is the device’s Maximum Transmission Unit size in bytes.",
															Required:            false,
															Optional:            true,
															Computed:            false,
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
															Description:         "NetworkName is the name of the vSphere network to which the device will be connected.",
															MarkdownDescription: "NetworkName is the name of the vSphere network to which the device will be connected.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"routes": schema.ListNestedAttribute{
															Description:         "Routes is a list of optional, static routes applied to the device.",
															MarkdownDescription: "Routes is a list of optional, static routes applied to the device.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"metric": schema.Int64Attribute{
																		Description:         "Metric is the weight/priority of the route.",
																		MarkdownDescription: "Metric is the weight/priority of the route.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"to": schema.StringAttribute{
																		Description:         "To is an IPv4 or IPv6 address.",
																		MarkdownDescription: "To is an IPv4 or IPv6 address.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"via": schema.StringAttribute{
																		Description:         "Via is an IPv4 or IPv6 address.",
																		MarkdownDescription: "Via is an IPv4 or IPv6 address.",
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

														"search_domains": schema.ListAttribute{
															Description:         "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
															MarkdownDescription: "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"skip_ip_allocation": schema.BoolAttribute{
															Description:         "SkipIPAllocation allows the device to not have IP address or DHCP configured. This is suitable for devices for which IP allocation is handled externally, eg. using Multus CNI. If true, CAPV will not verify IP address allocation.",
															MarkdownDescription: "SkipIPAllocation allows the device to not have IP address or DHCP configured. This is suitable for devices for which IP allocation is handled externally, eg. using Multus CNI. If true, CAPV will not verify IP address allocation.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"preferred_api_server_cidr": schema.StringAttribute{
												Description:         "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine Deprecated: This field is going to be removed in a future release.",
												MarkdownDescription: "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine Deprecated: This field is going to be removed in a future release.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"routes": schema.ListNestedAttribute{
												Description:         "Routes is a list of optional, static routes applied to the virtual machine.",
												MarkdownDescription: "Routes is a list of optional, static routes applied to the virtual machine.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"metric": schema.Int64Attribute{
															Description:         "Metric is the weight/priority of the route.",
															MarkdownDescription: "Metric is the weight/priority of the route.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"to": schema.StringAttribute{
															Description:         "To is an IPv4 or IPv6 address.",
															MarkdownDescription: "To is an IPv4 or IPv6 address.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"via": schema.StringAttribute{
															Description:         "Via is an IPv4 or IPv6 address.",
															MarkdownDescription: "Via is an IPv4 or IPv6 address.",
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
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"num_cp_us": schema.Int64Attribute{
										Description:         "NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										MarkdownDescription: "NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"num_cores_per_socket": schema.Int64Attribute{
										Description:         "NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										MarkdownDescription: "NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"os": schema.StringAttribute{
										Description:         "OS is the Operating System of the virtual machine Defaults to Linux",
										MarkdownDescription: "OS is the Operating System of the virtual machine Defaults to Linux",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pci_devices": schema.ListNestedAttribute{
										Description:         "PciDevices is the list of pci devices used by the virtual machine.",
										MarkdownDescription: "PciDevices is the list of pci devices used by the virtual machine.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"custom_label": schema.StringAttribute{
													Description:         "CustomLabel is the hardware label of a virtual machine's PCI device. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
													MarkdownDescription: "CustomLabel is the hardware label of a virtual machine's PCI device. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"device_id": schema.Int64Attribute{
													Description:         "DeviceID is the device ID of a virtual machine's PCI, in integer. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Mutually exclusive with VGPUProfile as VGPUProfile and DeviceID + VendorID are two independent ways to define PCI devices.",
													MarkdownDescription: "DeviceID is the device ID of a virtual machine's PCI, in integer. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Mutually exclusive with VGPUProfile as VGPUProfile and DeviceID + VendorID are two independent ways to define PCI devices.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"v_gpu_profile": schema.StringAttribute{
													Description:         "VGPUProfile is the profile name of a virtual machine's vGPU, in string. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Mutually exclusive with DeviceID and VendorID as VGPUProfile and DeviceID + VendorID are two independent ways to define PCI devices.",
													MarkdownDescription: "VGPUProfile is the profile name of a virtual machine's vGPU, in string. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Mutually exclusive with DeviceID and VendorID as VGPUProfile and DeviceID + VendorID are two independent ways to define PCI devices.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"vendor_id": schema.Int64Attribute{
													Description:         "VendorId is the vendor ID of a virtual machine's PCI, in integer. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Mutually exclusive with VGPUProfile as VGPUProfile and DeviceID + VendorID are two independent ways to define PCI devices.",
													MarkdownDescription: "VendorId is the vendor ID of a virtual machine's PCI, in integer. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Mutually exclusive with VGPUProfile as VGPUProfile and DeviceID + VendorID are two independent ways to define PCI devices.",
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

									"power_off_mode": schema.StringAttribute{
										Description:         "PowerOffMode describes the desired behavior when powering off a VM. There are three, supported power off modes: hard, soft, and trySoft. The first mode, hard, is the equivalent of a physical system's power cord being ripped from the wall. The soft mode requires the VM's guest to have VM Tools installed and attempts to gracefully shut down the VM. Its variant, trySoft, first attempts a graceful shutdown, and if that fails or the VM is not in a powered off state after reaching the GuestSoftPowerOffTimeout, the VM is halted. If omitted, the mode defaults to hard.",
										MarkdownDescription: "PowerOffMode describes the desired behavior when powering off a VM. There are three, supported power off modes: hard, soft, and trySoft. The first mode, hard, is the equivalent of a physical system's power cord being ripped from the wall. The soft mode requires the VM's guest to have VM Tools installed and attempts to gracefully shut down the VM. Its variant, trySoft, first attempts a graceful shutdown, and if that fails or the VM is not in a powered off state after reaching the GuestSoftPowerOffTimeout, the VM is halted. If omitted, the mode defaults to hard.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("hard", "soft", "trySoft"),
										},
									},

									"provider_id": schema.StringAttribute{
										Description:         "ProviderID is the virtual machine's BIOS UUID formated as vsphere://12345678-1234-1234-1234-123456789abc",
										MarkdownDescription: "ProviderID is the virtual machine's BIOS UUID formated as vsphere://12345678-1234-1234-1234-123456789abc",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource_pool": schema.StringAttribute{
										Description:         "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
										MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.StringAttribute{
										Description:         "Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.",
										MarkdownDescription: "Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"snapshot": schema.StringAttribute{
										Description:         "Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.",
										MarkdownDescription: "Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_policy_name": schema.StringAttribute{
										Description:         "StoragePolicyName of the storage policy to use with this Virtual Machine",
										MarkdownDescription: "StoragePolicyName of the storage policy to use with this Virtual Machine",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag_i_ds": schema.ListAttribute{
										Description:         "TagIDs is an optional set of tags to add to an instance. Specified tagIDs must use URN-notation instead of display names.",
										MarkdownDescription: "TagIDs is an optional set of tags to add to an instance. Specified tagIDs must use URN-notation instead of display names.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"template": schema.StringAttribute{
										Description:         "Template is the name or inventory path of the template used to clone the virtual machine.",
										MarkdownDescription: "Template is the name or inventory path of the template used to clone the virtual machine.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"thumbprint": schema.StringAttribute{
										Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.",
										MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
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

func (r *InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoVsphereMachineTemplateV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("VSphereMachineTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
