---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "multicluster.crd.antrea.io"
description: |-
  ClusterSet represents a ClusterSet.
---

# k8s_multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest (Data Source)

ClusterSet represents a ClusterSet.

## Example Usage

```terraform
data "k8s_multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ClusterSetSpec defines the desired state of ClusterSet. (see [below for nested schema](#nestedatt--spec))

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

- `leaders` (Attributes List) Leaders include leader clusters known to the member clusters. (see [below for nested schema](#nestedatt--spec--leaders))

Optional:

- `members` (Attributes List) Members include member clusters known to the leader clusters. Used in leader cluster. (see [below for nested schema](#nestedatt--spec--members))
- `namespace` (String) The leader cluster Namespace in which the ClusterSet is defined. Used in member cluster.

<a id="nestedatt--spec--leaders"></a>
### Nested Schema for `spec.leaders`

Optional:

- `cluster_id` (String) Identify member cluster in ClusterSet.
- `secret` (String) Secret name to access API server of the member from the leader cluster.
- `server` (String) API server of the destination cluster.
- `service_account` (String) ServiceAccount used by the member cluster to access into leader cluster.


<a id="nestedatt--spec--members"></a>
### Nested Schema for `spec.members`

Optional:

- `cluster_id` (String) Identify member cluster in ClusterSet.
- `secret` (String) Secret name to access API server of the member from the leader cluster.
- `server` (String) API server of the destination cluster.
- `service_account` (String) ServiceAccount used by the member cluster to access into leader cluster.