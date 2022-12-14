---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_scheduling_k8s_io_priority_class_v1 Resource - terraform-provider-k8s"
subcategory: "scheduling.k8s.io"
description: |-
  PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
---

# k8s_scheduling_k8s_io_priority_class_v1 (Resource)

PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.

## Example Usage

```terraform
resource "k8s_scheduling_k8s_io_priority_class_v1" "minimal" {
  metadata = {
    name = "test"
  }
  value = -100
}

resource "k8s_scheduling_k8s_io_priority_class_v1" "example" {
  metadata = {
    name = "test"
  }
  value = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `value` (Number) The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.

### Optional

- `description` (String) description is an arbitrary string that usually provides guidelines on when this priority class should be used.
- `global_default` (Boolean) globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as 'globalDefault'. However, if more than one PriorityClasses exists with their 'globalDefault' field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
- `preemption_policy` (String) PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.

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
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


