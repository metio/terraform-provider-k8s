---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cluster_x_k8s_io_machine_health_check_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "cluster.x-k8s.io"
description: |-
  MachineHealthCheck is the Schema for the machinehealthchecks API.
---

# k8s_cluster_x_k8s_io_machine_health_check_v1beta1_manifest (Data Source)

MachineHealthCheck is the Schema for the machinehealthchecks API.

## Example Usage

```terraform
data "k8s_cluster_x_k8s_io_machine_health_check_v1beta1_manifest" "example" {
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

- `spec` (Attributes) Specification of machine health check policy (see [below for nested schema](#nestedatt--spec))

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

- `cluster_name` (String) ClusterName is the name of the Cluster this object belongs to.
- `selector` (Attributes) Label selector to match machines whose health will be exercised (see [below for nested schema](#nestedatt--spec--selector))
- `unhealthy_conditions` (Attributes List) UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy.  The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy. (see [below for nested schema](#nestedatt--spec--unhealthy_conditions))

Optional:

- `max_unhealthy` (String) Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by'selector' are not healthy.
- `node_startup_timeout` (String) Machines older than this duration without a node will be considered to havefailed and will be remediated.If not set, this value is defaulted to 10 minutes.If you wish to disable this feature, set the value explicitly to 0.
- `remediation_template` (Attributes) RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.This field is completely optional, when filled, the MachineHealthCheck controllercreates a new object from the template referenced and hands off remediation of the machine toa controller that lives outside of Cluster API. (see [below for nested schema](#nestedatt--spec--remediation_template))
- `unhealthy_range` (String) Any further remediation is only allowed if the number of machines selected by 'selector' as not healthyis within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy.Eg. '[3-5]' - This means that remediation will be allowed only when:(a) there are at least 3 unhealthy machines (and)(b) there are at most 5 unhealthy machines

<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--selector--match_expressions"></a>
### Nested Schema for `spec.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--unhealthy_conditions"></a>
### Nested Schema for `spec.unhealthy_conditions`

Required:

- `status` (String)
- `timeout` (String)
- `type` (String)


<a id="nestedatt--spec--remediation_template"></a>
### Nested Schema for `spec.remediation_template`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids