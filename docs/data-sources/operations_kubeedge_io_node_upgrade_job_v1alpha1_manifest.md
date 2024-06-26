---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "operations.kubeedge.io"
description: |-
  NodeUpgradeJob is used to upgrade edge node from cloud side.
---

# k8s_operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest (Data Source)

NodeUpgradeJob is used to upgrade edge node from cloud side.

## Example Usage

```terraform
data "k8s_operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) Specification of the desired behavior of NodeUpgradeJob. (see [below for nested schema](#nestedatt--spec))

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

- `check_items` (List of String) CheckItems specifies the items need to be checked before the task is executed. The default CheckItems value is nil.
- `concurrency` (Number) Concurrency specifies the max number of edge nodes that can be upgraded at the same time. The default Concurrency value is 1.
- `failure_tolerate` (String) FailureTolerate specifies the task tolerance failure ratio. The default FailureTolerate value is 0.1.
- `image` (String) Image specifies a container image name, the image contains: keadm and edgecore. keadm is used as upgradetool, to install the new version of edgecore. The image name consists of registry hostname and repository name, if it includes the tag or digest, the tag or digest will be overwritten by Version field above. If the registry hostname is empty, docker.io will be used as default. The default image name is: kubeedge/installation-package.
- `label_selector` (Attributes) LabelSelector is a filter to select member clusters by labels. It must match a node's labels for the NodeUpgradeJob to be operated on that node. Please note that sets of NodeNames and LabelSelector are ORed. Users must set one and can only set one. (see [below for nested schema](#nestedatt--spec--label_selector))
- `node_names` (List of String) NodeNames is a request to select some specific nodes. If it is non-empty, the upgrade job simply select these edge nodes to do upgrade operation. Please note that sets of NodeNames and LabelSelector are ORed. Users must set one and can only set one.
- `timeout_seconds` (Number) TimeoutSeconds limits the duration of the node upgrade job. Default to 300. If set to 0, we'll use the default value 300.
- `version` (String)

<a id="nestedatt--spec--label_selector"></a>
### Nested Schema for `spec.label_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--label_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--label_selector--match_expressions"></a>
### Nested Schema for `spec.label_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
