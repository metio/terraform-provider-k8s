---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "ipam.cluster.x-k8s.io"
description: |-
  IPAddress is the Schema for the ipaddress API.
---

# k8s_ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest (Data Source)

IPAddress is the Schema for the ipaddress API.

## Example Usage

```terraform
data "k8s_ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest" "example" {
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

- `spec` (Attributes) IPAddressSpec is the desired state of an IPAddress. (see [below for nested schema](#nestedatt--spec))

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

- `address` (String) Address is the IP address.
- `claim_ref` (Attributes) ClaimRef is a reference to the claim this IPAddress was created for. (see [below for nested schema](#nestedatt--spec--claim_ref))
- `pool_ref` (Attributes) PoolRef is a reference to the pool that this IPAddress was created from. (see [below for nested schema](#nestedatt--spec--pool_ref))
- `prefix` (Number) Prefix is the prefix of the address.

Optional:

- `gateway` (String) Gateway is the network gateway of the network the address is from.

<a id="nestedatt--spec--claim_ref"></a>
### Nested Schema for `spec.claim_ref`

Optional:

- `name` (String) Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names


<a id="nestedatt--spec--pool_ref"></a>
### Nested Schema for `spec.pool_ref`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
