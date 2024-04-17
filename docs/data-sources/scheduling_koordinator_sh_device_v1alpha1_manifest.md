---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_scheduling_koordinator_sh_device_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "scheduling.koordinator.sh"
description: |-
  
---

# k8s_scheduling_koordinator_sh_device_v1alpha1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_scheduling_koordinator_sh_device_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `devices` (Attributes List) (see [below for nested schema](#nestedatt--spec--devices))

<a id="nestedatt--spec--devices"></a>
### Nested Schema for `spec.devices`

Required:

- `health` (Boolean) Health indicates whether the device is normal

Optional:

- `id` (String) UUID represents the UUID of device
- `labels` (Map of String) Labels represents the device properties that can be used to organize and categorize (scope and select) objects
- `minor` (Number) Minor represents the Minor number of Device, starting from 0
- `module_id` (Number) ModuleID represents the physical id of Device
- `resources` (Map of String) Resources is a set of (resource name, quantity) pairs
- `topology` (Attributes) Topology represents the topology information about the device (see [below for nested schema](#nestedatt--spec--devices--topology))
- `type` (String) Type represents the type of device
- `vf_groups` (Attributes List) VFGroups represents the virtual function devices (see [below for nested schema](#nestedatt--spec--devices--vf_groups))

<a id="nestedatt--spec--devices--topology"></a>
### Nested Schema for `spec.devices.topology`

Required:

- `node_id` (Number) NodeID is the ID of NUMA Node to which the device belongs, it should be unique across different CPU Sockets
- `pcie_id` (String) PCIEID is the ID of PCIE Switch to which the device is connected, it should be unique across difference NUMANodes
- `socket_id` (Number) SocketID is the ID of CPU Socket to which the device belongs

Optional:

- `bus_id` (String) BusID is the domain:bus:device.function formatted identifier of PCI/PCIE device


<a id="nestedatt--spec--devices--vf_groups"></a>
### Nested Schema for `spec.devices.vf_groups`

Optional:

- `labels` (Map of String) Labels represents the Virtual Function properties that can be used to organize and categorize (scope and select) objects
- `vfs` (Attributes List) VFs are the virtual function devices which belong to the group (see [below for nested schema](#nestedatt--spec--devices--vf_groups--vfs))

<a id="nestedatt--spec--devices--vf_groups--vfs"></a>
### Nested Schema for `spec.devices.vf_groups.vfs`

Required:

- `minor` (Number) Minor represents the Minor number of VirtualFunction, starting from 0, used to identify virtual function.

Optional:

- `bus_id` (String) BusID is the domain:bus:device.function formatted identifier of PCI/PCIE virtual function device