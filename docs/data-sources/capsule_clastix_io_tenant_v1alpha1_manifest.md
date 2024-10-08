---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_capsule_clastix_io_tenant_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "capsule.clastix.io"
description: |-
  Tenant is the Schema for the tenants API.
---

# k8s_capsule_clastix_io_tenant_v1alpha1_manifest (Data Source)

Tenant is the Schema for the tenants API.

## Example Usage

```terraform
data "k8s_capsule_clastix_io_tenant_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) TenantSpec defines the desired state of Tenant. (see [below for nested schema](#nestedatt--spec))

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

- `owner` (Attributes) OwnerSpec defines tenant owner name and kind. (see [below for nested schema](#nestedatt--spec--owner))

Optional:

- `additional_role_bindings` (Attributes List) (see [below for nested schema](#nestedatt--spec--additional_role_bindings))
- `container_registries` (Attributes) (see [below for nested schema](#nestedatt--spec--container_registries))
- `external_service_i_ps` (Attributes) (see [below for nested schema](#nestedatt--spec--external_service_i_ps))
- `ingress_classes` (Attributes) (see [below for nested schema](#nestedatt--spec--ingress_classes))
- `ingress_hostnames` (Attributes) (see [below for nested schema](#nestedatt--spec--ingress_hostnames))
- `limit_ranges` (Attributes List) (see [below for nested schema](#nestedatt--spec--limit_ranges))
- `namespace_quota` (Number)
- `namespaces_metadata` (Attributes) (see [below for nested schema](#nestedatt--spec--namespaces_metadata))
- `network_policies` (Attributes List) (see [below for nested schema](#nestedatt--spec--network_policies))
- `node_selector` (Map of String)
- `resource_quotas` (Attributes List) (see [below for nested schema](#nestedatt--spec--resource_quotas))
- `services_metadata` (Attributes) (see [below for nested schema](#nestedatt--spec--services_metadata))
- `storage_classes` (Attributes) (see [below for nested schema](#nestedatt--spec--storage_classes))

<a id="nestedatt--spec--owner"></a>
### Nested Schema for `spec.owner`

Required:

- `kind` (String)
- `name` (String)


<a id="nestedatt--spec--additional_role_bindings"></a>
### Nested Schema for `spec.additional_role_bindings`

Required:

- `cluster_role_name` (String)
- `subjects` (Attributes List) kubebuilder:validation:Minimum=1 (see [below for nested schema](#nestedatt--spec--additional_role_bindings--subjects))

<a id="nestedatt--spec--additional_role_bindings--subjects"></a>
### Nested Schema for `spec.additional_role_bindings.subjects`

Required:

- `kind` (String) Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.
- `name` (String) Name of the object being referenced.

Optional:

- `api_group` (String) APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.
- `namespace` (String) Namespace of the referenced object. If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.



<a id="nestedatt--spec--container_registries"></a>
### Nested Schema for `spec.container_registries`

Optional:

- `allowed` (List of String)
- `allowed_regex` (String)


<a id="nestedatt--spec--external_service_i_ps"></a>
### Nested Schema for `spec.external_service_i_ps`

Required:

- `allowed` (List of String)


<a id="nestedatt--spec--ingress_classes"></a>
### Nested Schema for `spec.ingress_classes`

Optional:

- `allowed` (List of String)
- `allowed_regex` (String)


<a id="nestedatt--spec--ingress_hostnames"></a>
### Nested Schema for `spec.ingress_hostnames`

Optional:

- `allowed` (List of String)
- `allowed_regex` (String)


<a id="nestedatt--spec--limit_ranges"></a>
### Nested Schema for `spec.limit_ranges`

Required:

- `limits` (Attributes List) Limits is the list of LimitRangeItem objects that are enforced. (see [below for nested schema](#nestedatt--spec--limit_ranges--limits))

<a id="nestedatt--spec--limit_ranges--limits"></a>
### Nested Schema for `spec.limit_ranges.limits`

Required:

- `type` (String) Type of resource that this limit applies to.

Optional:

- `default` (Map of String) Default resource requirement limit value by resource name if resource limit is omitted.
- `default_request` (Map of String) DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
- `max` (Map of String) Max usage constraints on this kind by resource name.
- `max_limit_request_ratio` (Map of String) MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
- `min` (Map of String) Min usage constraints on this kind by resource name.



<a id="nestedatt--spec--namespaces_metadata"></a>
### Nested Schema for `spec.namespaces_metadata`

Optional:

- `additional_annotations` (Map of String)
- `additional_labels` (Map of String)


<a id="nestedatt--spec--network_policies"></a>
### Nested Schema for `spec.network_policies`

Required:

- `pod_selector` (Attributes) podSelector selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace. (see [below for nested schema](#nestedatt--spec--network_policies--pod_selector))

Optional:

- `egress` (Attributes List) egress is a list of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8 (see [below for nested schema](#nestedatt--spec--network_policies--egress))
- `ingress` (Attributes List) ingress is a list of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default) (see [below for nested schema](#nestedatt--spec--network_policies--ingress))
- `policy_types` (List of String) policyTypes is a list of rule types that the NetworkPolicy relates to. Valid options are ['Ingress'], ['Egress'], or ['Ingress', 'Egress']. If this field is not specified, it will default based on the existence of ingress or egress rules; policies that contain an egress section are assumed to affect egress, and all policies (whether or not they contain an ingress section) are assumed to affect ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ 'Egress' ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include 'Egress' (since such a policy would not include an egress section and would otherwise default to just [ 'Ingress' ]). This field is beta-level in 1.8

<a id="nestedatt--spec--network_policies--pod_selector"></a>
### Nested Schema for `spec.network_policies.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--network_policies--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--network_policies--pod_selector--match_expressions"></a>
### Nested Schema for `spec.network_policies.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--network_policies--egress"></a>
### Nested Schema for `spec.network_policies.egress`

Optional:

- `ports` (Attributes List) ports is a list of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list. (see [below for nested schema](#nestedatt--spec--network_policies--egress--ports))
- `to` (Attributes List) to is a list of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list. (see [below for nested schema](#nestedatt--spec--network_policies--egress--to))

<a id="nestedatt--spec--network_policies--egress--ports"></a>
### Nested Schema for `spec.network_policies.egress.ports`

Optional:

- `end_port` (Number) endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
- `port` (String) port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
- `protocol` (String) protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.


<a id="nestedatt--spec--network_policies--egress--to"></a>
### Nested Schema for `spec.network_policies.egress.to`

Optional:

- `ip_block` (Attributes) ipBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be. (see [below for nested schema](#nestedatt--spec--network_policies--egress--to--ip_block))
- `namespace_selector` (Attributes) namespaceSelector selects namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces. If podSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the namespaces selected by namespaceSelector. Otherwise it selects all pods in the namespaces selected by namespaceSelector. (see [below for nested schema](#nestedatt--spec--network_policies--egress--to--namespace_selector))
- `pod_selector` (Attributes) podSelector is a label selector which selects pods. This field follows standard label selector semantics; if present but empty, it selects all pods. If namespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the pods matching podSelector in the policy's own namespace. (see [below for nested schema](#nestedatt--spec--network_policies--egress--to--pod_selector))

<a id="nestedatt--spec--network_policies--egress--to--ip_block"></a>
### Nested Schema for `spec.network_policies.egress.to.ip_block`

Required:

- `cidr` (String) cidr is a string representing the IPBlock Valid examples are '192.168.1.0/24' or '2001:db8::/64'

Optional:

- `except` (List of String) except is a slice of CIDRs that should not be included within an IPBlock Valid examples are '192.168.1.0/24' or '2001:db8::/64' Except values will be rejected if they are outside the cidr range


<a id="nestedatt--spec--network_policies--egress--to--namespace_selector"></a>
### Nested Schema for `spec.network_policies.egress.to.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--network_policies--egress--to--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--network_policies--egress--to--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.network_policies.egress.to.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--network_policies--egress--to--pod_selector"></a>
### Nested Schema for `spec.network_policies.egress.to.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--network_policies--egress--to--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--network_policies--egress--to--pod_selector--match_expressions"></a>
### Nested Schema for `spec.network_policies.egress.to.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedatt--spec--network_policies--ingress"></a>
### Nested Schema for `spec.network_policies.ingress`

Optional:

- `from` (Attributes List) from is a list of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--from))
- `ports` (Attributes List) ports is a list of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--ports))

<a id="nestedatt--spec--network_policies--ingress--from"></a>
### Nested Schema for `spec.network_policies.ingress.from`

Optional:

- `ip_block` (Attributes) ipBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--from--ip_block))
- `namespace_selector` (Attributes) namespaceSelector selects namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces. If podSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the namespaces selected by namespaceSelector. Otherwise it selects all pods in the namespaces selected by namespaceSelector. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--from--namespace_selector))
- `pod_selector` (Attributes) podSelector is a label selector which selects pods. This field follows standard label selector semantics; if present but empty, it selects all pods. If namespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the pods matching podSelector in the policy's own namespace. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--from--pod_selector))

<a id="nestedatt--spec--network_policies--ingress--from--ip_block"></a>
### Nested Schema for `spec.network_policies.ingress.from.ip_block`

Required:

- `cidr` (String) cidr is a string representing the IPBlock Valid examples are '192.168.1.0/24' or '2001:db8::/64'

Optional:

- `except` (List of String) except is a slice of CIDRs that should not be included within an IPBlock Valid examples are '192.168.1.0/24' or '2001:db8::/64' Except values will be rejected if they are outside the cidr range


<a id="nestedatt--spec--network_policies--ingress--from--namespace_selector"></a>
### Nested Schema for `spec.network_policies.ingress.from.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--from--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--network_policies--ingress--from--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.network_policies.ingress.from.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--network_policies--ingress--from--pod_selector"></a>
### Nested Schema for `spec.network_policies.ingress.from.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--network_policies--ingress--from--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--network_policies--ingress--from--pod_selector--match_expressions"></a>
### Nested Schema for `spec.network_policies.ingress.from.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.




<a id="nestedatt--spec--network_policies--ingress--ports"></a>
### Nested Schema for `spec.network_policies.ingress.ports`

Optional:

- `end_port` (Number) endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
- `port` (String) port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
- `protocol` (String) protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.




<a id="nestedatt--spec--resource_quotas"></a>
### Nested Schema for `spec.resource_quotas`

Optional:

- `hard` (Map of String) hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
- `scope_selector` (Attributes) scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched. (see [below for nested schema](#nestedatt--spec--resource_quotas--scope_selector))
- `scopes` (List of String) A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.

<a id="nestedatt--spec--resource_quotas--scope_selector"></a>
### Nested Schema for `spec.resource_quotas.scope_selector`

Optional:

- `match_expressions` (Attributes List) A list of scope selector requirements by scope of the resources. (see [below for nested schema](#nestedatt--spec--resource_quotas--scope_selector--match_expressions))

<a id="nestedatt--spec--resource_quotas--scope_selector--match_expressions"></a>
### Nested Schema for `spec.resource_quotas.scope_selector.match_expressions`

Required:

- `operator` (String) Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
- `scope_name` (String) The name of the scope that the selector applies to.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.




<a id="nestedatt--spec--services_metadata"></a>
### Nested Schema for `spec.services_metadata`

Optional:

- `additional_annotations` (Map of String)
- `additional_labels` (Map of String)


<a id="nestedatt--spec--storage_classes"></a>
### Nested Schema for `spec.storage_classes`

Optional:

- `allowed` (List of String)
- `allowed_regex` (String)
