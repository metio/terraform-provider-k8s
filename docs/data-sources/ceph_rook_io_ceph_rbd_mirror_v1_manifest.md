---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ceph_rook_io_ceph_rbd_mirror_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "ceph.rook.io"
description: |-
  CephRBDMirror represents a Ceph RBD Mirror
---

# k8s_ceph_rook_io_ceph_rbd_mirror_v1_manifest (Data Source)

CephRBDMirror represents a Ceph RBD Mirror

## Example Usage

```terraform
data "k8s_ceph_rook_io_ceph_rbd_mirror_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    count = 7
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) RBDMirroringSpec represents the specification of an RBD mirror daemon (see [below for nested schema](#nestedatt--spec))

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

- `count` (Number) Count represents the number of rbd mirror instance to run

Optional:

- `annotations` (Map of String) The annotations-related configuration to add/set on each Pod related object.
- `labels` (Map of String) The labels-related configuration to add/set on each Pod related object.
- `peers` (Attributes) Peers represents the peers spec (see [below for nested schema](#nestedatt--spec--peers))
- `placement` (Attributes) (see [below for nested schema](#nestedatt--spec--placement))
- `priority_class_name` (String) PriorityClassName sets priority class on the rbd mirror pods
- `resources` (Attributes) The resource requirements for the rbd mirror pods (see [below for nested schema](#nestedatt--spec--resources))

<a id="nestedatt--spec--peers"></a>
### Nested Schema for `spec.peers`

Optional:

- `secret_names` (List of String) SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers


<a id="nestedatt--spec--placement"></a>
### Nested Schema for `spec.placement`

Optional:

- `node_affinity` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--node_affinity))
- `pod_affinity` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity))
- `pod_anti_affinity` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity))
- `tolerations` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--tolerations))
- `topology_spread_constraints` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--topology_spread_constraints))

<a id="nestedatt--spec--placement--node_affinity"></a>
### Nested Schema for `spec.placement.node_affinity`

Optional:

- `preferred_during_scheduling_ignored_during_execution` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution))
- `required_during_scheduling_ignored_during_execution` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.node_affinity.preferred_during_scheduling_ignored_during_execution`

Required:

- `preference` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution--preference))
- `weight` (Number)

<a id="nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution--preference"></a>
### Nested Schema for `spec.placement.node_affinity.preferred_during_scheduling_ignored_during_execution.preference`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution--preference--match_expressions))
- `match_fields` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution--preference--match_fields))

<a id="nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution--preference--match_expressions"></a>
### Nested Schema for `spec.placement.node_affinity.preferred_during_scheduling_ignored_during_execution.preference.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)


<a id="nestedatt--spec--placement--node_affinity--preferred_during_scheduling_ignored_during_execution--preference--match_fields"></a>
### Nested Schema for `spec.placement.node_affinity.preferred_during_scheduling_ignored_during_execution.preference.match_fields`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)




<a id="nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.node_affinity.required_during_scheduling_ignored_during_execution`

Required:

- `node_selector_terms` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms))

<a id="nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms"></a>
### Nested Schema for `spec.placement.node_affinity.required_during_scheduling_ignored_during_execution.node_selector_terms`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_expressions))
- `match_fields` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_fields))

<a id="nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_expressions"></a>
### Nested Schema for `spec.placement.node_affinity.required_during_scheduling_ignored_during_execution.node_selector_terms.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)


<a id="nestedatt--spec--placement--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_fields"></a>
### Nested Schema for `spec.placement.node_affinity.required_during_scheduling_ignored_during_execution.node_selector_terms.match_fields`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)





<a id="nestedatt--spec--placement--pod_affinity"></a>
### Nested Schema for `spec.placement.pod_affinity`

Optional:

- `preferred_during_scheduling_ignored_during_execution` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution))
- `required_during_scheduling_ignored_during_execution` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.pod_affinity.preferred_during_scheduling_ignored_during_execution`

Required:

- `pod_affinity_term` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term))
- `weight` (Number)

<a id="nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term"></a>
### Nested Schema for `spec.placement.pod_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term`

Required:

- `topology_key` (String)

Optional:

- `label_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector))
- `match_label_keys` (List of String)
- `mismatch_label_keys` (List of String)
- `namespace_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector))
- `namespaces` (List of String)

<a id="nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector"></a>
### Nested Schema for `spec.placement.pod_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.label_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.label_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)



<a id="nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector"></a>
### Nested Schema for `spec.placement.pod_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.namespace_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.namespace_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)





<a id="nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.pod_affinity.required_during_scheduling_ignored_during_execution`

Required:

- `topology_key` (String)

Optional:

- `label_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--label_selector))
- `match_label_keys` (List of String)
- `mismatch_label_keys` (List of String)
- `namespace_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--namespace_selector))
- `namespaces` (List of String)

<a id="nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--label_selector"></a>
### Nested Schema for `spec.placement.pod_affinity.required_during_scheduling_ignored_during_execution.label_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--label_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--label_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_affinity.required_during_scheduling_ignored_during_execution.label_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)



<a id="nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--namespace_selector"></a>
### Nested Schema for `spec.placement.pod_affinity.required_during_scheduling_ignored_during_execution.namespace_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--namespace_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_affinity--required_during_scheduling_ignored_during_execution--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_affinity.required_during_scheduling_ignored_during_execution.namespace_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)





<a id="nestedatt--spec--placement--pod_anti_affinity"></a>
### Nested Schema for `spec.placement.pod_anti_affinity`

Optional:

- `preferred_during_scheduling_ignored_during_execution` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution))
- `required_during_scheduling_ignored_during_execution` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution`

Required:

- `pod_affinity_term` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term))
- `weight` (Number)

<a id="nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term`

Required:

- `topology_key` (String)

Optional:

- `label_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector))
- `match_label_keys` (List of String)
- `mismatch_label_keys` (List of String)
- `namespace_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector))
- `namespaces` (List of String)

<a id="nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.label_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--label_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.label_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)



<a id="nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.namespace_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.pod_affinity_term.namespace_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)





<a id="nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.required_during_scheduling_ignored_during_execution`

Required:

- `topology_key` (String)

Optional:

- `label_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector))
- `match_label_keys` (List of String)
- `mismatch_label_keys` (List of String)
- `namespace_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector))
- `namespaces` (List of String)

<a id="nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.required_during_scheduling_ignored_during_execution.label_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.required_during_scheduling_ignored_during_execution.label_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)



<a id="nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespace_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.placement.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespace_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)





<a id="nestedatt--spec--placement--tolerations"></a>
### Nested Schema for `spec.placement.tolerations`

Optional:

- `effect` (String)
- `key` (String)
- `operator` (String)
- `toleration_seconds` (Number)
- `value` (String)


<a id="nestedatt--spec--placement--topology_spread_constraints"></a>
### Nested Schema for `spec.placement.topology_spread_constraints`

Required:

- `max_skew` (Number)
- `topology_key` (String)
- `when_unsatisfiable` (String)

Optional:

- `label_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--placement--topology_spread_constraints--label_selector))
- `match_label_keys` (List of String)
- `min_domains` (Number)
- `node_affinity_policy` (String)
- `node_taints_policy` (String)

<a id="nestedatt--spec--placement--topology_spread_constraints--label_selector"></a>
### Nested Schema for `spec.placement.topology_spread_constraints.label_selector`

Optional:

- `match_expressions` (Attributes List) (see [below for nested schema](#nestedatt--spec--placement--topology_spread_constraints--label_selector--match_expressions))
- `match_labels` (Map of String)

<a id="nestedatt--spec--placement--topology_spread_constraints--label_selector--match_expressions"></a>
### Nested Schema for `spec.placement.topology_spread_constraints.label_selector.match_expressions`

Required:

- `key` (String)
- `operator` (String)

Optional:

- `values` (List of String)





<a id="nestedatt--spec--resources"></a>
### Nested Schema for `spec.resources`

Optional:

- `claims` (Attributes List) Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers. (see [below for nested schema](#nestedatt--spec--resources--claims))
- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

<a id="nestedatt--spec--resources--claims"></a>
### Nested Schema for `spec.resources.claims`

Required:

- `name` (String) Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.

Optional:

- `request` (String) Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.
