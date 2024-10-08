---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest Data Source - terraform-provider-k8s"
subcategory: "infrastructure.cluster.x-k8s.io"
description: |-
  VSphereClusterIdentity defines the account to be used for reconciling clusters Deprecated: This type will be removed in one of the next releases.
---

# k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest (Data Source)

VSphereClusterIdentity defines the account to be used for reconciling clusters Deprecated: This type will be removed in one of the next releases.

## Example Usage

```terraform
data "k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_identity_v1alpha4_manifest" "example" {
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

Optional:

- `allowed_namespaces` (Attributes) AllowedNamespaces is used to identify which namespaces are allowed to use this account. Namespaces can be selected with a label selector. If this object is nil, no namespaces will be allowed (see [below for nested schema](#nestedatt--spec--allowed_namespaces))
- `secret_name` (String) SecretName references a Secret inside the controller namespace with the credentials to use

<a id="nestedatt--spec--allowed_namespaces"></a>
### Nested Schema for `spec.allowed_namespaces`

Optional:

- `selector` (Attributes) Selector is a standard Kubernetes LabelSelector. A label query over a set of resources. (see [below for nested schema](#nestedatt--spec--allowed_namespaces--selector))

<a id="nestedatt--spec--allowed_namespaces--selector"></a>
### Nested Schema for `spec.allowed_namespaces.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--allowed_namespaces--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--allowed_namespaces--selector--match_expressions"></a>
### Nested Schema for `spec.allowed_namespaces.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
