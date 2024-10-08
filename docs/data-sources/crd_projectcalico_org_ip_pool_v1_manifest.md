---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_crd_projectcalico_org_ip_pool_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "crd.projectcalico.org"
description: |-
  
---

# k8s_crd_projectcalico_org_ip_pool_v1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_crd_projectcalico_org_ip_pool_v1_manifest" "example" {
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

- `spec` (Attributes) IPPoolSpec contains the specification for an IPPool resource. (see [below for nested schema](#nestedatt--spec))

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

Required:

- `cidr` (String) The pool CIDR.

Optional:

- `allowed_uses` (List of String) AllowedUse controls what the IP pool will be used for. If not specified or empty, defaults to ['Tunnel', 'Workload'] for back-compatibility
- `block_size` (Number) The block size to use for IP address assignments from this pool. Defaults to 26 for IPv4 and 122 for IPv6.
- `disable_bgp_export` (Boolean) Disable exporting routes from this IP Pool's CIDR over BGP. [Default: false]
- `disabled` (Boolean) When disabled is true, Calico IPAM will not assign addresses from this pool.
- `ipip` (Attributes) Deprecated: this field is only used for APIv1 backwards compatibility. Setting this field is not allowed, this field is for internal use only. (see [below for nested schema](#nestedatt--spec--ipip))
- `ipip_mode` (String) Contains configuration for IPIP tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. IPIP tunneling is disabled).
- `nat_outgoing` (Boolean) When natOutgoing is true, packets sent from Calico networked containers in this pool to destinations outside of this pool will be masqueraded.
- `node_selector` (String) Allows IPPool to allocate for a specific node by label selector.
- `vxlan_mode` (String) Contains configuration for VXLAN tunneling for this pool. If not specified, then this is defaulted to 'Never' (i.e. VXLAN tunneling is disabled).

<a id="nestedatt--spec--ipip"></a>
### Nested Schema for `spec.ipip`

Optional:

- `enabled` (Boolean) When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.
- `mode` (String) The IPIP mode. This can be one of 'always' or 'cross-subnet'. A mode of 'always' will also use IPIP tunneling for routing to destination IP addresses within this pool. A mode of 'cross-subnet' will only use IPIP tunneling when the destination node is on a different subnet to the originating node. The default value (if not specified) is 'always'.
