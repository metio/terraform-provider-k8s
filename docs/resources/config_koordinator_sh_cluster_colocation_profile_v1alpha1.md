---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "config.koordinator.sh"
description: |-
  ClusterColocationProfile is the Schema for the ClusterColocationProfile API
---

# k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1 (Resource)

ClusterColocationProfile is the Schema for the ClusterColocationProfile API

## Example Usage

```terraform
resource "k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1" "minimal" {
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

- `spec` (Attributes) ClusterColocationProfileSpec is a description of a ClusterColocationProfile. (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `annotations` (Map of String) Annotations describes the k/v pair that needs to inject into Pod.Annotations
- `koordinator_priority` (Number) KoordinatorPriority defines the Pod sub-priority in Koordinator. The priority value will be injected into Pod as label koordinator.sh/priority. Various Koordinator components determine the priority of the Pod in the Koordinator through KoordinatorPriority and the priority value in PriorityClassName. The higher the value, the higher the priority.
- `labels` (Map of String) Labels describes the k/v pair that needs to inject into Pod.Labels
- `namespace_selector` (Attributes) NamespaceSelector decides whether to mutate/validate Pods if the namespace matches the selector. Default to the empty LabelSelector, which matches everything. (see [below for nested schema](#nestedatt--spec--namespace_selector))
- `patch` (Dynamic) Patch indicates patching podTemplate that will be injected to the Pod.
- `priority_class_name` (String) If specified, the priorityClassName and the priority value defined in PriorityClass will be injected into the Pod. The PriorityClassName, priority value in PriorityClassName and KoordinatorPriority will affect the scheduling, preemption and other behaviors of Koordinator system.
- `qos_class` (String) QoSClass describes the type of Koordinator QoS that the Pod is running. The value will be injected into Pod as label koordinator.sh/qosClass. Options are LSE/LSR/LS/BE/SYSTEM.
- `scheduler_name` (String) If specified, the pod will be dispatched by specified scheduler.
- `selector` (Attributes) Selector decides whether to mutate/validate Pods if the Pod matches the selector. Default to the empty LabelSelector, which matches everything. (see [below for nested schema](#nestedatt--spec--selector))

<a id="nestedatt--spec--namespace_selector"></a>
### Nested Schema for `spec.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--selector--match_expressions"></a>
### Nested Schema for `spec.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.

