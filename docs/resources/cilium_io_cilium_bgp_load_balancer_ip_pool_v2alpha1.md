---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cilium_io_cilium_bgp_load_balancer_ip_pool_v2alpha1 Resource - terraform-provider-k8s"
subcategory: "cilium.io"
description: |-
  CiliumBGPLoadBalancerIPPool is a Kubernetes third-party resource which instructs the BGP control plane to allocate and advertise IPs for Services of type LoadBalancer.
---

# k8s_cilium_io_cilium_bgp_load_balancer_ip_pool_v2alpha1 (Resource)

CiliumBGPLoadBalancerIPPool is a Kubernetes third-party resource which instructs the BGP control plane to allocate and advertise IPs for Services of type LoadBalancer.

## Example Usage

```terraform
resource "k8s_cilium_io_cilium_bgp_load_balancer_ip_pool_v2alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) Spec is a human readable description for a BGP load balancer ip pool. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
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

- `prefix` (String) The CIDR block of IPs to allocate from.

Optional:

- `default` (Boolean) Default determines if this is the default IP pool for allocating from when LBSelector is nil or empty.
- `lb_selector` (Attributes) LBSelector will determine if a created LoadBalancer is allocated an IP from this pool. (see [below for nested schema](#nestedatt--spec--lb_selector))
- `node_selector` (Attributes) NodeSelector selects a group of nodes which will advertise the presence of any LoadBalancers allocated from this IP pool.  If nil all nodes will advertise the presence of any LoadBalancer allocated an IP from this pool. (see [below for nested schema](#nestedatt--spec--node_selector))

<a id="nestedatt--spec--lb_selector"></a>
### Nested Schema for `spec.lb_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--lb_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--lb_selector--match_expressions"></a>
### Nested Schema for `spec.lb_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--node_selector"></a>
### Nested Schema for `spec.node_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--node_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--node_selector--match_expressions"></a>
### Nested Schema for `spec.node_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.


