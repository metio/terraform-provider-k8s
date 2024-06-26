---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_multicluster_crd_antrea_io_gateway_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "multicluster.crd.antrea.io"
description: |-
  Gateway includes information of a Multi-cluster Gateway.
---

# k8s_multicluster_crd_antrea_io_gateway_v1alpha1_manifest (Data Source)

Gateway includes information of a Multi-cluster Gateway.

## Example Usage

```terraform
data "k8s_multicluster_crd_antrea_io_gateway_v1alpha1_manifest" "example" {
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

- `gateway_ip` (String) Cross-cluster tunnel IP of the Gateway.
- `internal_ip` (String) In-cluster tunnel IP of the Gateway.
- `service_cidr` (String) Service CIDR of the local member cluster.
- `wire_guard` (Attributes) WireGuardInfo includes information of a WireGuard tunnel. (see [below for nested schema](#nestedatt--wire_guard))

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


<a id="nestedatt--wire_guard"></a>
### Nested Schema for `wire_guard`

Optional:

- `public_key` (String) Public key of the WireGuard tunnel.
