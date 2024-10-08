---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sriovnetwork_openshift_io_sriov_network_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "sriovnetwork.openshift.io"
description: |-
  SriovNetwork is the Schema for the sriovnetworks API
---

# k8s_sriovnetwork_openshift_io_sriov_network_v1_manifest (Data Source)

SriovNetwork is the Schema for the sriovnetworks API

## Example Usage

```terraform
data "k8s_sriovnetwork_openshift_io_sriov_network_v1_manifest" "example" {
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

- `spec` (Attributes) SriovNetworkSpec defines the desired state of SriovNetwork (see [below for nested schema](#nestedatt--spec))

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

- `resource_name` (String) SRIOV Network device plugin endpoint resource name

Optional:

- `capabilities` (String) Capabilities to be configured for this network. Capabilities supported: (mac|ips), e.g. '{'mac': true}'
- `ipam` (String) IPAM configuration to be used for this network.
- `link_state` (String) VF link state (enable|disable|auto)
- `log_file` (String) LogFile sets the log file of the SRIOV CNI plugin logs. If unset (default), this will log to stderr and thus to multus and container runtime logs.
- `log_level` (String) LogLevel sets the log level of the SRIOV CNI plugin - either of panic, error, warning, info, debug. Defaults to info if left blank.
- `max_tx_rate` (Number) Maximum tx rate, in Mbps, for the VF. Defaults to 0 (no rate limiting)
- `meta_plugins` (String) MetaPluginsConfig configuration to be used in order to chain metaplugins to the sriov interface returned by the operator.
- `min_tx_rate` (Number) Minimum tx rate, in Mbps, for the VF. Defaults to 0 (no rate limiting). min_tx_rate should be <= max_tx_rate.
- `network_namespace` (String) Namespace of the NetworkAttachmentDefinition custom resource
- `spoof_chk` (String) VF spoof check, (on|off)
- `trust` (String) VF trust mode (on|off)
- `vlan` (Number) VLAN ID to assign for the VF. Defaults to 0.
- `vlan_proto` (String) VLAN proto to assign for the VF. Defaults to 802.1q.
- `vlan_qo_s` (Number) VLAN QoS ID to assign for the VF. Defaults to 0.
