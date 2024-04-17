---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_network_openshift_io_net_namespace_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "network.openshift.io"
description: |-
  NetNamespace describes a single isolated network. When using the redhat/openshift-ovs-multitenant plugin, every Namespace will have a corresponding NetNamespace object with the same name. (When using redhat/openshift-ovs-subnet, NetNamespaces are not used.)  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
---

# k8s_network_openshift_io_net_namespace_v1_manifest (Data Source)

NetNamespace describes a single isolated network. When using the redhat/openshift-ovs-multitenant plugin, every Namespace will have a corresponding NetNamespace object with the same name. (When using redhat/openshift-ovs-subnet, NetNamespaces are not used.)  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).

## Example Usage

```terraform
data "k8s_network_openshift_io_net_namespace_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `netid` (Number) NetID is the network identifier of the network namespace assigned to each overlay network packet. This can be manipulated with the 'oc adm pod-network' commands.
- `netname` (String) NetName is the name of the network namespace. (This is the same as the object's name, but both fields must be set.)

### Optional

- `egress_i_ps` (List of String) EgressIPs is a list of reserved IPs that will be used as the source for external traffic coming from pods in this namespace. (If empty, external traffic will be masqueraded to Node IPs.)

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.