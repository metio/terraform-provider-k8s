---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "addons.cluster.x-k8s.io"
description: |-
  ClusterResourceSetBinding lists all matching ClusterResourceSets with the cluster it belongs to.
---

# k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest (Data Source)

ClusterResourceSetBinding lists all matching ClusterResourceSets with the cluster it belongs to.

## Example Usage

```terraform
data "k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1beta1_manifest" "example" {
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

- `spec` (Attributes) ClusterResourceSetBindingSpec defines the desired state of ClusterResourceSetBinding. (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `bindings` (Attributes List) Bindings is a list of ClusterResourceSets and their resources. (see [below for nested schema](#nestedatt--spec--bindings))
- `cluster_name` (String) ClusterName is the name of the Cluster this binding applies to.Note: this field mandatory in v1beta2.

<a id="nestedatt--spec--bindings"></a>
### Nested Schema for `spec.bindings`

Required:

- `cluster_resource_set_name` (String) ClusterResourceSetName is the name of the ClusterResourceSet that is applied to the owner cluster of the binding.

Optional:

- `resources` (Attributes List) Resources is a list of resources that the ClusterResourceSet has. (see [below for nested schema](#nestedatt--spec--bindings--resources))

<a id="nestedatt--spec--bindings--resources"></a>
### Nested Schema for `spec.bindings.resources`

Required:

- `applied` (Boolean) Applied is to track if a resource is applied to the cluster or not.
- `kind` (String) Kind of the resource. Supported kinds are: Secrets and ConfigMaps.
- `name` (String) Name of the resource that is in the same namespace with ClusterResourceSet object.

Optional:

- `hash` (String) Hash is the hash of a resource's data. This can be used to decide if a resource is changed.For 'ApplyOnce' ClusterResourceSet.spec.strategy, this is no-op as that strategy does not act on change.
- `last_applied_time` (String) LastAppliedTime identifies when this resource was last applied to the cluster.