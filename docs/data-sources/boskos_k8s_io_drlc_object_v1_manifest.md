---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_boskos_k8s_io_drlc_object_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "boskos.k8s.io"
description: |-
  Defines the lifecycle of a dynamic resource. All Resource of a given type will be constructed using the same configuration
---

# k8s_boskos_k8s_io_drlc_object_v1_manifest (Data Source)

Defines the lifecycle of a dynamic resource. All Resource of a given type will be constructed using the same configuration

## Example Usage

```terraform
data "k8s_boskos_k8s_io_drlc_object_v1_manifest" "example" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

- `config` (Attributes) Config information about how to create the object (see [below for nested schema](#nestedatt--spec--config))
- `lifespan` (Number) Lifespan of a resource, time after which the resource should be reset
- `max_count` (Number) Maxiumum number of resources expected. This maximum may be temporarily exceeded while resources are in the process of being deleted, though this is only expected when MaxCount is lowered.
- `min_count` (Number) Minimum number of resources to be used as a buffer. Resources in the process of being deleted and cleaned up are included in this count.
- `needs` (Map of String) Define the resource needs to create the object
- `state` (String)

<a id="nestedatt--spec--config"></a>
### Nested Schema for `spec.config`

Optional:

- `content` (String)
- `type` (String) The dynamic resource type
