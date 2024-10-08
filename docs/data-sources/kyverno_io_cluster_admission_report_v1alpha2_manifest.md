---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kyverno_io_cluster_admission_report_v1alpha2_manifest Data Source - terraform-provider-k8s"
subcategory: "kyverno.io"
description: |-
  ClusterAdmissionReport is the Schema for the ClusterAdmissionReports API
---

# k8s_kyverno_io_cluster_admission_report_v1alpha2_manifest (Data Source)

ClusterAdmissionReport is the Schema for the ClusterAdmissionReports API

## Example Usage

```terraform
data "k8s_kyverno_io_cluster_admission_report_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    owner = {
      api_version = "some-version"
      kind        = "some-kind"
      name        = "some-name"
      uid         = "some-uid"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

- `owner` (Attributes) Owner is a reference to the report owner (e.g. a Deployment, Namespace, or Node) (see [below for nested schema](#nestedatt--spec--owner))

Optional:

- `results` (Attributes List) PolicyReportResult provides result details (see [below for nested schema](#nestedatt--spec--results))
- `summary` (Attributes) PolicyReportSummary provides a summary of results (see [below for nested schema](#nestedatt--spec--summary))

<a id="nestedatt--spec--owner"></a>
### Nested Schema for `spec.owner`

Required:

- `api_version` (String) API version of the referent.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids

Optional:

- `block_owner_deletion` (Boolean) If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
- `controller` (Boolean) If true, this reference points to the managing controller.


<a id="nestedatt--spec--results"></a>
### Nested Schema for `spec.results`

Required:

- `policy` (String) Policy is the name or identifier of the policy

Optional:

- `category` (String) Category indicates policy category
- `message` (String) Description is a short user friendly message for the policy rule
- `properties` (Map of String) Properties provides additional information for the policy rule
- `resource_selector` (Attributes) SubjectSelector is an optional label selector for checked Kubernetes resources. For example, a policy result may apply to all pods that match a label. Either a Subject or a SubjectSelector can be specified. If neither are provided, the result is assumed to be for the policy report scope. (see [below for nested schema](#nestedatt--spec--results--resource_selector))
- `resources` (Attributes List) Subjects is an optional reference to the checked Kubernetes resources (see [below for nested schema](#nestedatt--spec--results--resources))
- `result` (String) Result indicates the outcome of the policy rule execution
- `rule` (String) Rule is the name or identifier of the rule within the policy
- `scored` (Boolean) Scored indicates if this result is scored
- `severity` (String) Severity indicates policy check result criticality
- `source` (String) Source is an identifier for the policy engine that manages this report
- `timestamp` (Attributes) Timestamp indicates the time the result was found (see [below for nested schema](#nestedatt--spec--results--timestamp))

<a id="nestedatt--spec--results--resource_selector"></a>
### Nested Schema for `spec.results.resource_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--results--resource_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--results--resource_selector--match_expressions"></a>
### Nested Schema for `spec.results.resource_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--results--resources"></a>
### Nested Schema for `spec.results.resources`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--results--timestamp"></a>
### Nested Schema for `spec.results.timestamp`

Required:

- `nanos` (Number) Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.
- `seconds` (Number) Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.



<a id="nestedatt--spec--summary"></a>
### Nested Schema for `spec.summary`

Optional:

- `error` (Number) Error provides the count of policies that could not be evaluated
- `fail` (Number) Fail provides the count of policies whose requirements were not met
- `pass` (Number) Pass provides the count of policies whose requirements were met
- `skip` (Number) Skip indicates the count of policies that were not selected for evaluation
- `warn` (Number) Warn provides the count of non-scored policies whose requirements were not met
