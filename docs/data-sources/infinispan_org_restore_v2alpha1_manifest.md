---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infinispan_org_restore_v2alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "infinispan.org"
description: |-
  Restore is the Schema for the restores API
---

# k8s_infinispan_org_restore_v2alpha1_manifest (Data Source)

Restore is the Schema for the restores API

## Example Usage

```terraform
data "k8s_infinispan_org_restore_v2alpha1_manifest" "example" {
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

- `spec` (Attributes) BackupSpec defines the desired state of Backup (see [below for nested schema](#nestedatt--spec))

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

- `backup` (String) The Infinispan Backup to restore
- `cluster` (String) Infinispan cluster name

Optional:

- `container` (Attributes) InfinispanContainerSpec specify resource requirements per container (see [below for nested schema](#nestedatt--spec--container))
- `resources` (Attributes) (see [below for nested schema](#nestedatt--spec--resources))

<a id="nestedatt--spec--container"></a>
### Nested Schema for `spec.container`

Optional:

- `cli_extra_jvm_opts` (String)
- `cpu` (String)
- `extra_jvm_opts` (String)
- `memory` (String)
- `router_extra_jvm_opts` (String)


<a id="nestedatt--spec--resources"></a>
### Nested Schema for `spec.resources`

Optional:

- `cache_configs` (List of String) Deprecated and to be removed on subsequent release. Use .Templates instead.
- `caches` (List of String)
- `counters` (List of String)
- `proto_schemas` (List of String)
- `scripts` (List of String) Deprecated and to be removed on subsequent release. Use .Tasks instead.
- `tasks` (List of String)
- `templates` (List of String)
