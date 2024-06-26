---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_site_superedge_io_node_group_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "site.superedge.io"
description: |-
  NodeGroup is the Schema for the nodegroups API
---

# k8s_site_superedge_io_node_group_v1alpha1_manifest (Data Source)

NodeGroup is the Schema for the nodegroups API

## Example Usage

```terraform
data "k8s_site_superedge_io_node_group_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) NodeGroupSpec defines the desired state of NodeGroup (see [below for nested schema](#nestedatt--spec))

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

- `autofindnodekeys` (List of String) If specified, create new NodeUnits based on node have same label keys, for different values will create different nodeunites
- `nodeunits` (List of String) If specified, If nodeUnit exists, join NodeGroup directly
- `selector` (Attributes) If specified, Label selector for nodeUnit. (see [below for nested schema](#nestedatt--spec--selector))
- `workload` (Attributes List) If specified, Nodegroup bound workload (see [below for nested schema](#nestedatt--spec--workload))

<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- `annotations` (Map of String) If specified, select node to join nodeUnit according to Annotations
- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs.

<a id="nestedatt--spec--selector--match_expressions"></a>
### Nested Schema for `spec.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--workload"></a>
### Nested Schema for `spec.workload`

Optional:

- `name` (String) workload name
- `selector` (Attributes) If specified, Label selector for workload. (see [below for nested schema](#nestedatt--spec--workload--selector))
- `type` (String) workload type, Value can be pod, deploy, ds, service, job, st

<a id="nestedatt--spec--workload--selector"></a>
### Nested Schema for `spec.workload.selector`

Optional:

- `annotations` (Map of String) If specified, select node to join nodeUnit according to Annotations
- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--workload--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs.

<a id="nestedatt--spec--workload--selector--match_expressions"></a>
### Nested Schema for `spec.workload.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
