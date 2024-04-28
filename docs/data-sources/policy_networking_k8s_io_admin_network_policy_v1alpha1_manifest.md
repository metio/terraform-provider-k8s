---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "policy.networking.k8s.io"
description: |-
  AdminNetworkPolicy is  a cluster level resource that is part of theAdminNetworkPolicy API.
---

# k8s_policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest (Data Source)

AdminNetworkPolicy is  a cluster level resource that is part of theAdminNetworkPolicy API.

## Example Usage

```terraform
data "k8s_policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Specification of the desired behavior of AdminNetworkPolicy. (see [below for nested schema](#nestedatt--spec))

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

- `priority` (Number) Priority is a value from 0 to 1000. Rules with lower priority values havehigher precedence, and are checked before rules with higher priority values.All AdminNetworkPolicy rules have higher precedence than NetworkPolicy orBaselineAdminNetworkPolicy rulesThe behavior is undefined if two ANP objects have same priority.Support: Core
- `subject` (Attributes) Subject defines the pods to which this AdminNetworkPolicy applies.Note that host-networked pods are not included in subject selection.Support: Core (see [below for nested schema](#nestedatt--spec--subject))

Optional:

- `egress` (Attributes List) Egress is the list of Egress rules to be applied to the selected pods.A total of 100 rules will be allowed in each ANP instance.The relative precedence of egress rules within a single ANP object (all ofwhich share the priority) will be determined by the order in which the ruleis written. Thus, a rule that appears at the top of the egress ruleswould take the highest precedence.ANPs with no egress rules do not affect egress traffic.Support: Core (see [below for nested schema](#nestedatt--spec--egress))
- `ingress` (Attributes List) Ingress is the list of Ingress rules to be applied to the selected pods.A total of 100 rules will be allowed in each ANP instance.The relative precedence of ingress rules within a single ANP object (all ofwhich share the priority) will be determined by the order in which the ruleis written. Thus, a rule that appears at the top of the ingress ruleswould take the highest precedence.ANPs with no ingress rules do not affect ingress traffic.Support: Core (see [below for nested schema](#nestedatt--spec--ingress))

<a id="nestedatt--spec--subject"></a>
### Nested Schema for `spec.subject`

Optional:

- `namespaces` (Attributes) Namespaces is used to select pods via namespace selectors. (see [below for nested schema](#nestedatt--spec--subject--namespaces))
- `pods` (Attributes) Pods is used to select pods via namespace AND pod selectors. (see [below for nested schema](#nestedatt--spec--subject--pods))

<a id="nestedatt--spec--subject--namespaces"></a>
### Nested Schema for `spec.subject.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--subject--namespaces--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--subject--namespaces--match_expressions"></a>
### Nested Schema for `spec.subject.namespaces.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--subject--pods"></a>
### Nested Schema for `spec.subject.pods`

Required:

- `namespace_selector` (Attributes) NamespaceSelector follows standard label selector semantics; if empty,it selects all Namespaces. (see [below for nested schema](#nestedatt--spec--subject--pods--namespace_selector))
- `pod_selector` (Attributes) PodSelector is used to explicitly select pods within a namespace; if empty,it selects all Pods. (see [below for nested schema](#nestedatt--spec--subject--pods--pod_selector))

<a id="nestedatt--spec--subject--pods--namespace_selector"></a>
### Nested Schema for `spec.subject.pods.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--subject--pods--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--subject--pods--pod_selector--match_expressions"></a>
### Nested Schema for `spec.subject.pods.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--subject--pods--pod_selector"></a>
### Nested Schema for `spec.subject.pods.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--subject--pods--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--subject--pods--pod_selector--match_expressions"></a>
### Nested Schema for `spec.subject.pods.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.





<a id="nestedatt--spec--egress"></a>
### Nested Schema for `spec.egress`

Required:

- `action` (String) Action specifies the effect this rule will have on matching traffic.Currently the following actions are supported:Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy)Deny: denies the selected trafficPass: instructs the selected traffic to skip any remaining ANP rules, andthen pass execution to any NetworkPolicies that select the pod.If the pod is not selected by any NetworkPolicies then executionis passed to any BaselineAdminNetworkPolicies that select the pod.Support: Core
- `to` (Attributes List) To is the List of destinations whose traffic this rule applies to.If any AdminNetworkPolicyEgressPeer matches the destination of outgoingtraffic then the specified action is applied.This field must be defined and contain at least one item.Support: Core (see [below for nested schema](#nestedatt--spec--egress--to))

Optional:

- `name` (String) Name is an identifier for this rule, that may be no more than 100 charactersin length. This field should be used by the implementation to helpimprove observability, readability and error-reporting for any appliedAdminNetworkPolicies.Support: Core
- `ports` (Attributes List) Ports allows for matching traffic based on port and protocols.This field is a list of destination ports for the outgoing egress traffic.If Ports is not set then the rule does not filter traffic via port.Support: Core (see [below for nested schema](#nestedatt--spec--egress--ports))

<a id="nestedatt--spec--egress--to"></a>
### Nested Schema for `spec.egress.to`

Optional:

- `namespaces` (Attributes) Namespaces defines a way to select all pods within a set of Namespaces.Note that host-networked pods are not included in this type of peer.Support: Core (see [below for nested schema](#nestedatt--spec--egress--to--namespaces))
- `pods` (Attributes) Pods defines a way to select a set of pods ina set of namespaces. Note that host-networked podsare not included in this type of peer.Support: Core (see [below for nested schema](#nestedatt--spec--egress--to--pods))

<a id="nestedatt--spec--egress--to--namespaces"></a>
### Nested Schema for `spec.egress.to.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--egress--to--pods--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--egress--to--pods--match_expressions"></a>
### Nested Schema for `spec.egress.to.pods.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--egress--to--pods"></a>
### Nested Schema for `spec.egress.to.pods`

Required:

- `namespace_selector` (Attributes) NamespaceSelector follows standard label selector semantics; if empty,it selects all Namespaces. (see [below for nested schema](#nestedatt--spec--egress--to--pods--namespace_selector))
- `pod_selector` (Attributes) PodSelector is used to explicitly select pods within a namespace; if empty,it selects all Pods. (see [below for nested schema](#nestedatt--spec--egress--to--pods--pod_selector))

<a id="nestedatt--spec--egress--to--pods--namespace_selector"></a>
### Nested Schema for `spec.egress.to.pods.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--egress--to--pods--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--egress--to--pods--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.egress.to.pods.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--egress--to--pods--pod_selector"></a>
### Nested Schema for `spec.egress.to.pods.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--egress--to--pods--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--egress--to--pods--pod_selector--match_expressions"></a>
### Nested Schema for `spec.egress.to.pods.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.





<a id="nestedatt--spec--egress--ports"></a>
### Nested Schema for `spec.egress.ports`

Optional:

- `port_number` (Attributes) Port selects a port on a pod(s) based on number.Support: Core (see [below for nested schema](#nestedatt--spec--egress--ports--port_number))
- `port_range` (Attributes) PortRange selects a port range on a pod(s) based on provided start and endvalues.Support: Core (see [below for nested schema](#nestedatt--spec--egress--ports--port_range))

<a id="nestedatt--spec--egress--ports--port_number"></a>
### Nested Schema for `spec.egress.ports.port_number`

Required:

- `port` (Number) Number defines a network port value.Support: Core
- `protocol` (String) Protocol is the network protocol (TCP, UDP, or SCTP) which traffic mustmatch. If not specified, this field defaults to TCP.Support: Core


<a id="nestedatt--spec--egress--ports--port_range"></a>
### Nested Schema for `spec.egress.ports.port_range`

Required:

- `end` (Number) End defines a network port that is the end of a port range, the End valuemust be greater than Start.Support: Core
- `start` (Number) Start defines a network port that is the start of a port range, the Startvalue must be less than End.Support: Core

Optional:

- `protocol` (String) Protocol is the network protocol (TCP, UDP, or SCTP) which traffic mustmatch. If not specified, this field defaults to TCP.Support: Core




<a id="nestedatt--spec--ingress"></a>
### Nested Schema for `spec.ingress`

Required:

- `action` (String) Action specifies the effect this rule will have on matching traffic.Currently the following actions are supported:Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy)Deny: denies the selected trafficPass: instructs the selected traffic to skip any remaining ANP rules, andthen pass execution to any NetworkPolicies that select the pod.If the pod is not selected by any NetworkPolicies then executionis passed to any BaselineAdminNetworkPolicies that select the pod.Support: Core
- `from` (Attributes List) From is the list of sources whose traffic this rule applies to.If any AdminNetworkPolicyIngressPeer matches the source of incomingtraffic then the specified action is applied.This field must be defined and contain at least one item.Support: Core (see [below for nested schema](#nestedatt--spec--ingress--from))

Optional:

- `name` (String) Name is an identifier for this rule, that may be no more than 100 charactersin length. This field should be used by the implementation to helpimprove observability, readability and error-reporting for any appliedAdminNetworkPolicies.Support: Core
- `ports` (Attributes List) Ports allows for matching traffic based on port and protocols.This field is a list of ports which should be matched onthe pods selected for this policy i.e the subject of the policy.So it matches on the destination port for the ingress traffic.If Ports is not set then the rule does not filter traffic via port.Support: Core (see [below for nested schema](#nestedatt--spec--ingress--ports))

<a id="nestedatt--spec--ingress--from"></a>
### Nested Schema for `spec.ingress.from`

Optional:

- `namespaces` (Attributes) Namespaces defines a way to select all pods within a set of Namespaces.Note that host-networked pods are not included in this type of peer.Support: Core (see [below for nested schema](#nestedatt--spec--ingress--from--namespaces))
- `pods` (Attributes) Pods defines a way to select a set of pods ina set of namespaces. Note that host-networked podsare not included in this type of peer.Support: Core (see [below for nested schema](#nestedatt--spec--ingress--from--pods))

<a id="nestedatt--spec--ingress--from--namespaces"></a>
### Nested Schema for `spec.ingress.from.namespaces`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--ingress--from--pods--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--ingress--from--pods--match_expressions"></a>
### Nested Schema for `spec.ingress.from.pods.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--ingress--from--pods"></a>
### Nested Schema for `spec.ingress.from.pods`

Required:

- `namespace_selector` (Attributes) NamespaceSelector follows standard label selector semantics; if empty,it selects all Namespaces. (see [below for nested schema](#nestedatt--spec--ingress--from--pods--namespace_selector))
- `pod_selector` (Attributes) PodSelector is used to explicitly select pods within a namespace; if empty,it selects all Pods. (see [below for nested schema](#nestedatt--spec--ingress--from--pods--pod_selector))

<a id="nestedatt--spec--ingress--from--pods--namespace_selector"></a>
### Nested Schema for `spec.ingress.from.pods.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--ingress--from--pods--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--ingress--from--pods--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.ingress.from.pods.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--ingress--from--pods--pod_selector"></a>
### Nested Schema for `spec.ingress.from.pods.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--ingress--from--pods--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--ingress--from--pods--pod_selector--match_expressions"></a>
### Nested Schema for `spec.ingress.from.pods.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.





<a id="nestedatt--spec--ingress--ports"></a>
### Nested Schema for `spec.ingress.ports`

Optional:

- `port_number` (Attributes) Port selects a port on a pod(s) based on number.Support: Core (see [below for nested schema](#nestedatt--spec--ingress--ports--port_number))
- `port_range` (Attributes) PortRange selects a port range on a pod(s) based on provided start and endvalues.Support: Core (see [below for nested schema](#nestedatt--spec--ingress--ports--port_range))

<a id="nestedatt--spec--ingress--ports--port_number"></a>
### Nested Schema for `spec.ingress.ports.port_number`

Required:

- `port` (Number) Number defines a network port value.Support: Core
- `protocol` (String) Protocol is the network protocol (TCP, UDP, or SCTP) which traffic mustmatch. If not specified, this field defaults to TCP.Support: Core


<a id="nestedatt--spec--ingress--ports--port_range"></a>
### Nested Schema for `spec.ingress.ports.port_range`

Required:

- `end` (Number) End defines a network port that is the end of a port range, the End valuemust be greater than Start.Support: Core
- `start` (Number) Start defines a network port that is the start of a port range, the Startvalue must be less than End.Support: Core

Optional:

- `protocol` (String) Protocol is the network protocol (TCP, UDP, or SCTP) which traffic mustmatch. If not specified, this field defaults to TCP.Support: Core