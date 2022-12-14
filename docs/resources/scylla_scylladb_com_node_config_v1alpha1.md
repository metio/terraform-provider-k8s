---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_scylla_scylladb_com_node_config_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "scylla.scylladb.com"
description: |-
  
---

# k8s_scylla_scylladb_com_node_config_v1alpha1 (Resource)



## Example Usage

```terraform
resource "k8s_scylla_scylladb_com_node_config_v1alpha1" "minimal" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

- `placement` (Attributes) placement contains scheduling rules for NodeConfig Pods. (see [below for nested schema](#nestedatt--spec--placement))

Optional:

- `disable_optimizations` (Boolean) disableOptimizations controls if nodes matching placement requirements are going to be optimized. Turning off optimizations on already optimized Nodes does not revert changes.

<a id="nestedatt--spec--placement"></a>
### Nested Schema for `spec.placement`

Required:

- `node_selector` (Map of String) nodeSelector is a selector which must be true for the NodeConfig Pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node.

Optional:

- `affinity` (Attributes) affinity is a group of affinity scheduling rules for NodeConfig Pods. (see [below for nested schema](#nestedatt--spec--placement--affinity))
- `tolerations` (Attributes List) tolerations is a group of tolerations NodeConfig Pods are going to have. (see [below for nested schema](#nestedatt--spec--placement--tolerations))

<a id="nestedatt--spec--placement--affinity"></a>
### Nested Schema for `spec.placement.affinity`

Optional:

- `node_affinity` (Attributes) Describes node affinity scheduling rules for the pod. (see [below for nested schema](#nestedatt--spec--placement--affinity--node_affinity))
- `pod_affinity` (Attributes) Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_affinity))
- `pod_anti_affinity` (Attributes) Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity))

<a id="nestedatt--spec--placement--affinity--node_affinity"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity`

Optional:

- `preferred_during_scheduling_ignored_during_execution` (Attributes List) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution))
- `required_during_scheduling_ignored_during_execution` (Attributes) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution`

Required:

- `preference` (Attributes) A node selector term, associated with the corresponding weight. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--preference))
- `weight` (Number) Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--preference"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight`

Optional:

- `match_expressions` (Attributes List) A list of node selector requirements by node's labels. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--match_expressions))
- `match_fields` (Attributes List) A list of node selector requirements by node's fields. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--match_fields))

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.match_fields`

Required:

- `key` (String) The label key that the selector applies to.
- `operator` (String) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.


<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--match_fields"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.match_fields`

Required:

- `key` (String) The label key that the selector applies to.
- `operator` (String) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.




<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution`

Required:

- `node_selector_terms` (Attributes List) Required. A list of node selector terms. The terms are ORed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms))

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.node_selector_terms`

Optional:

- `match_expressions` (Attributes List) A list of node selector requirements by node's labels. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_expressions))
- `match_fields` (Attributes List) A list of node selector requirements by node's fields. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_fields))

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.node_selector_terms.match_fields`

Required:

- `key` (String) The label key that the selector applies to.
- `operator` (String) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.


<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--node_selector_terms--match_fields"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.node_selector_terms.match_fields`

Required:

- `key` (String) The label key that the selector applies to.
- `operator` (String) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.





<a id="nestedatt--spec--placement--affinity--pod_affinity"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity`

Optional:

- `preferred_during_scheduling_ignored_during_execution` (Attributes List) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution))
- `required_during_scheduling_ignored_during_execution` (Attributes List) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution`

Required:

- `pod_affinity_term` (Attributes) Required. A pod affinity term, associated with the corresponding weight. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term))
- `weight` (Number) weight associated with matching the corresponding podAffinityTerm, in the range 1-100.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight`

Required:

- `topology_key` (String) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.

Optional:

- `label_selector` (Attributes) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--label_selector))
- `namespace_selector` (Attributes) A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespace_selector))
- `namespaces` (List of String) namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--label_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespace_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution`

Required:

- `topology_key` (String) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.

Optional:

- `label_selector` (Attributes) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector))
- `namespace_selector` (Attributes) A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector))
- `namespaces` (List of String) namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedatt--spec--placement--affinity--pod_anti_affinity"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity`

Optional:

- `preferred_during_scheduling_ignored_during_execution` (Attributes List) The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution))
- `required_during_scheduling_ignored_during_execution` (Attributes List) If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution`

Required:

- `pod_affinity_term` (Attributes) Required. A pod affinity term, associated with the corresponding weight. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term))
- `weight` (Number) weight associated with matching the corresponding podAffinityTerm, in the range 1-100.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--pod_affinity_term"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight`

Required:

- `topology_key` (String) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.

Optional:

- `label_selector` (Attributes) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--label_selector))
- `namespace_selector` (Attributes) A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespace_selector))
- `namespaces` (List of String) namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--label_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespace_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution--weight--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.preferred_during_scheduling_ignored_during_execution.weight.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution`

Required:

- `topology_key` (String) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.

Optional:

- `label_selector` (Attributes) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector))
- `namespace_selector` (Attributes) A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector))
- `namespaces` (List of String) namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespace_selector"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--placement--affinity--pod_anti_affinity--required_during_scheduling_ignored_during_execution--namespaces--match_expressions"></a>
### Nested Schema for `spec.placement.affinity.pod_anti_affinity.required_during_scheduling_ignored_during_execution.namespaces.match_labels`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.






<a id="nestedatt--spec--placement--tolerations"></a>
### Nested Schema for `spec.placement.tolerations`

Optional:

- `effect` (String) Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
- `key` (String) Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
- `operator` (String) Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
- `toleration_seconds` (Number) TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
- `value` (String) Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.


