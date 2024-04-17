---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "cilium.io"
description: |-
  CiliumLoadBalancerIPPool is a Kubernetes third-party resource which is used to defined pools of IPs which the operator can use to to allocate and advertise IPs for Services of type LoadBalancer.
---

# k8s_cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest (Data Source)

CiliumLoadBalancerIPPool is a Kubernetes third-party resource which is used to defined pools of IPs which the operator can use to to allocate and advertise IPs for Services of type LoadBalancer.

## Example Usage

```terraform
data "k8s_cilium_io_cilium_load_balancer_ip_pool_v2alpha1_manifest" "example" {
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

- `spec` (Attributes) Spec is a human readable description for a BGP load balancer ip pool. (see [below for nested schema](#nestedatt--spec))

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

- `allow_first_last_i_ps` (String) AllowFirstLastIPs, if set to 'yes' means that the first and last IPs of each CIDR will be allocatable. If 'no' or undefined, these IPs will be reserved. This field is ignored for /{31,32} and /{127,128} CIDRs since reserving the first and last IPs would make the CIDRs unusable.
- `blocks` (Attributes List) Blocks is a list of CIDRs comprising this IP Pool (see [below for nested schema](#nestedatt--spec--blocks))
- `cidrs` (Attributes List) Cidrs is a list of CIDRs comprising this IP Pool Deprecated: please use the 'blocks' field instead. This field will be removed in a future release. https://github.com/cilium/cilium/issues/28590 (see [below for nested schema](#nestedatt--spec--cidrs))
- `disabled` (Boolean) Disabled, if set to true means that no new IPs will be allocated from this pool. Existing allocations will not be removed from services.
- `service_selector` (Attributes) ServiceSelector selects a set of services which are eligible to receive IPs from this (see [below for nested schema](#nestedatt--spec--service_selector))

<a id="nestedatt--spec--blocks"></a>
### Nested Schema for `spec.blocks`

Optional:

- `cidr` (String)
- `start` (String)
- `stop` (String)


<a id="nestedatt--spec--cidrs"></a>
### Nested Schema for `spec.cidrs`

Optional:

- `cidr` (String)
- `start` (String)
- `stop` (String)


<a id="nestedatt--spec--service_selector"></a>
### Nested Schema for `spec.service_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--service_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--service_selector--match_expressions"></a>
### Nested Schema for `spec.service_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.