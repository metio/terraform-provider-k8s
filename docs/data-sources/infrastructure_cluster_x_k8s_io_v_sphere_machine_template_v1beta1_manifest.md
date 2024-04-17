---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "infrastructure.cluster.x-k8s.io"
description: |-
  VSphereMachineTemplate is the Schema for the vspheremachinetemplates API.
---

# k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest (Data Source)

VSphereMachineTemplate is the Schema for the vspheremachinetemplates API.

## Example Usage

```terraform
data "k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_template_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) VSphereMachineTemplateSpec defines the desired state of VSphereMachineTemplate. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `template` (Attributes) VSphereMachineTemplateResource describes the data needed to create a VSphereMachine from a template. (see [below for nested schema](#nestedatt--spec--template))

<a id="nestedatt--spec--template"></a>
### Nested Schema for `spec.template`

Required:

- `spec` (Attributes) Spec is the specification of the desired behavior of the machine. (see [below for nested schema](#nestedatt--spec--template--spec))

Optional:

- `metadata` (Attributes) Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata (see [below for nested schema](#nestedatt--spec--template--metadata))

<a id="nestedatt--spec--template--spec"></a>
### Nested Schema for `spec.template.spec`

Required:

- `network` (Attributes) Network is the network configuration for this machine's VM. (see [below for nested schema](#nestedatt--spec--template--spec--network))
- `template` (String) Template is the name or inventory path of the template used to clone the virtual machine.

Optional:

- `additional_disks_gi_b` (List of String) AdditionalDisksGiB holds the sizes of additional disks of the virtual machine, in GiB Defaults to the eponymous property value in the template from which the virtual machine is cloned.
- `clone_mode` (String) CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.
- `custom_vmx_keys` (Map of String) CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map
- `datacenter` (String) Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located. Defaults to * which selects the default datacenter.
- `datastore` (String) Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.
- `disk_gi_b` (Number) DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.
- `failure_domain` (String) FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. For this infrastructure provider, the name is equivalent to the name of the VSphereDeploymentZone.
- `folder` (String) Folder is the name or inventory path of the folder in which the virtual machine is created/located.
- `guest_soft_power_off_timeout` (String) GuestSoftPowerOffTimeout sets the wait timeout for shutdown in the VM guest. The VM will be powered off forcibly after the timeout if the VM is still up and running when the PowerOffMode is set to trySoft.  This parameter only applies when the PowerOffMode is set to trySoft.  If omitted, the timeout defaults to 5 minutes.
- `hardware_version` (String) HardwareVersion is the hardware version of the virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned. Check the compatibility with the ESXi version before setting the value.
- `memory_mi_b` (Number) MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.
- `num_cores_per_socket` (Number) NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.
- `num_cp_us` (Number) NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.
- `os` (String) OS is the Operating System of the virtual machine Defaults to Linux
- `pci_devices` (Attributes List) PciDevices is the list of pci devices used by the virtual machine. (see [below for nested schema](#nestedatt--spec--template--spec--pci_devices))
- `power_off_mode` (String) PowerOffMode describes the desired behavior when powering off a VM.  There are three, supported power off modes: hard, soft, and trySoft. The first mode, hard, is the equivalent of a physical system's power cord being ripped from the wall. The soft mode requires the VM's guest to have VM Tools installed and attempts to gracefully shut down the VM. Its variant, trySoft, first attempts a graceful shutdown, and if that fails or the VM is not in a powered off state after reaching the GuestSoftPowerOffTimeout, the VM is halted.  If omitted, the mode defaults to hard.
- `provider_id` (String) ProviderID is the virtual machine's BIOS UUID formated as vsphere://12345678-1234-1234-1234-123456789abc
- `resource_pool` (String) ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.
- `server` (String) Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.
- `snapshot` (String) Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.
- `storage_policy_name` (String) StoragePolicyName of the storage policy to use with this Virtual Machine
- `tag_i_ds` (List of String) TagIDs is an optional set of tags to add to an instance. Specified tagIDs must use URN-notation instead of display names.
- `thumbprint` (String) Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.

<a id="nestedatt--spec--template--spec--network"></a>
### Nested Schema for `spec.template.spec.network`

Required:

- `devices` (Attributes List) Devices is the list of network devices used by the virtual machine. TODO(akutz) Make sure at least one network matches the ClusterSpec.CloudProviderConfiguration.Network.Name (see [below for nested schema](#nestedatt--spec--template--spec--thumbprint--devices))

Optional:

- `preferred_api_server_cidr` (String) PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine  Deprecated: This field is going to be removed in a future release.
- `routes` (Attributes List) Routes is a list of optional, static routes applied to the virtual machine. (see [below for nested schema](#nestedatt--spec--template--spec--thumbprint--routes))

<a id="nestedatt--spec--template--spec--thumbprint--devices"></a>
### Nested Schema for `spec.template.spec.thumbprint.devices`

Required:

- `network_name` (String) NetworkName is the name of the vSphere network to which the device will be connected.

Optional:

- `addresses_from_pools` (Attributes List) AddressesFromPools is a list of IPAddressPools that should be assigned to IPAddressClaims. The machine's cloud-init metadata will be populated with IPAddresses fulfilled by an IPAM provider. (see [below for nested schema](#nestedatt--spec--template--spec--thumbprint--devices--addresses_from_pools))
- `device_name` (String) DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.
- `dhcp4` (Boolean) DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.
- `dhcp4_overrides` (Attributes) DHCP4Overrides allows for the control over several DHCP behaviors. Overrides will only be applied when the corresponding DHCP flag is set. Only configured values will be sent, omitted values will default to distribution defaults. Dependent on support in the network stack for your distribution. For more information see the netplan reference (https://netplan.io/reference#dhcp-overrides) (see [below for nested schema](#nestedatt--spec--template--spec--thumbprint--devices--dhcp4_overrides))
- `dhcp6` (Boolean) DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.
- `dhcp6_overrides` (Attributes) DHCP6Overrides allows for the control over several DHCP behaviors. Overrides will only be applied when the corresponding DHCP flag is set. Only configured values will be sent, omitted values will default to distribution defaults. Dependent on support in the network stack for your distribution. For more information see the netplan reference (https://netplan.io/reference#dhcp-overrides) (see [below for nested schema](#nestedatt--spec--template--spec--thumbprint--devices--dhcp6_overrides))
- `gateway4` (String) Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.
- `gateway6` (String) Gateway4 is the IPv4 gateway used by this device.
- `ip_addrs` (List of String) IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4, DHCP6 and SkipIPAllocation are false.
- `mac_addr` (String) MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.
- `mtu` (Number) MTU is the device’s Maximum Transmission Unit size in bytes.
- `nameservers` (List of String) Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).
- `routes` (Attributes List) Routes is a list of optional, static routes applied to the device. (see [below for nested schema](#nestedatt--spec--template--spec--thumbprint--devices--routes))
- `search_domains` (List of String) SearchDomains is a list of search domains used when resolving IP addresses with DNS.
- `skip_ip_allocation` (Boolean) SkipIPAllocation allows the device to not have IP address or DHCP configured. This is suitable for devices for which IP allocation is handled externally, eg. using Multus CNI. If true, CAPV will not verify IP address allocation.

<a id="nestedatt--spec--template--spec--thumbprint--devices--addresses_from_pools"></a>
### Nested Schema for `spec.template.spec.thumbprint.devices.addresses_from_pools`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.


<a id="nestedatt--spec--template--spec--thumbprint--devices--dhcp4_overrides"></a>
### Nested Schema for `spec.template.spec.thumbprint.devices.dhcp4_overrides`

Optional:

- `hostname` (String) Hostname is the name which will be sent to the DHCP server instead of the machine's hostname.
- `route_metric` (Number) RouteMetric is used to prioritize routes for devices. A lower metric for an interface will have a higher priority.
- `send_hostname` (Boolean) SendHostname when 'true', the hostname of the machine will be sent to the DHCP server.
- `use_dns` (Boolean) UseDNS when 'true', the DNS servers in the DHCP server will be used and take precedence.
- `use_domains` (String) UseDomains can take the values 'true', 'false', or 'route'. When 'true', the domain name from the DHCP server will be used as the DNS search domain for this device. When 'route', the domain name from the DHCP response will be used for routing DNS only, not for searching.
- `use_hostname` (Boolean) UseHostname when 'true', the hostname from the DHCP server will be set as the transient hostname of the machine.
- `use_mtu` (Boolean) UseMTU when 'true', the MTU from the DHCP server will be set as the MTU of the device.
- `use_ntp` (Boolean) UseNTP when 'true', the NTP servers from the DHCP server will be used by systemd-timesyncd and take precedence.
- `use_routes` (String) UseRoutes when 'true', the routes from the DHCP server will be installed in the routing table.


<a id="nestedatt--spec--template--spec--thumbprint--devices--dhcp6_overrides"></a>
### Nested Schema for `spec.template.spec.thumbprint.devices.dhcp6_overrides`

Optional:

- `hostname` (String) Hostname is the name which will be sent to the DHCP server instead of the machine's hostname.
- `route_metric` (Number) RouteMetric is used to prioritize routes for devices. A lower metric for an interface will have a higher priority.
- `send_hostname` (Boolean) SendHostname when 'true', the hostname of the machine will be sent to the DHCP server.
- `use_dns` (Boolean) UseDNS when 'true', the DNS servers in the DHCP server will be used and take precedence.
- `use_domains` (String) UseDomains can take the values 'true', 'false', or 'route'. When 'true', the domain name from the DHCP server will be used as the DNS search domain for this device. When 'route', the domain name from the DHCP response will be used for routing DNS only, not for searching.
- `use_hostname` (Boolean) UseHostname when 'true', the hostname from the DHCP server will be set as the transient hostname of the machine.
- `use_mtu` (Boolean) UseMTU when 'true', the MTU from the DHCP server will be set as the MTU of the device.
- `use_ntp` (Boolean) UseNTP when 'true', the NTP servers from the DHCP server will be used by systemd-timesyncd and take precedence.
- `use_routes` (String) UseRoutes when 'true', the routes from the DHCP server will be installed in the routing table.


<a id="nestedatt--spec--template--spec--thumbprint--devices--routes"></a>
### Nested Schema for `spec.template.spec.thumbprint.devices.routes`

Required:

- `metric` (Number) Metric is the weight/priority of the route.
- `to` (String) To is an IPv4 or IPv6 address.
- `via` (String) Via is an IPv4 or IPv6 address.



<a id="nestedatt--spec--template--spec--thumbprint--routes"></a>
### Nested Schema for `spec.template.spec.thumbprint.routes`

Required:

- `metric` (Number) Metric is the weight/priority of the route.
- `to` (String) To is an IPv4 or IPv6 address.
- `via` (String) Via is an IPv4 or IPv6 address.



<a id="nestedatt--spec--template--spec--pci_devices"></a>
### Nested Schema for `spec.template.spec.pci_devices`

Optional:

- `device_id` (Number) DeviceID is the device ID of a virtual machine's PCI, in integer. Defaults to the eponymous property value in the template from which the virtual machine is cloned.
- `vendor_id` (Number) VendorId is the vendor ID of a virtual machine's PCI, in integer. Defaults to the eponymous property value in the template from which the virtual machine is cloned.



<a id="nestedatt--spec--template--metadata"></a>
### Nested Schema for `spec.template.metadata`

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
- `labels` (Map of String) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels