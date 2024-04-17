---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_schemas_schemahero_io_data_type_v1alpha4_manifest Data Source - terraform-provider-k8s"
subcategory: "schemas.schemahero.io"
description: |-
  DataType is the Schema for the datatypes API
---

# k8s_schemas_schemahero_io_data_type_v1alpha4_manifest (Data Source)

DataType is the Schema for the datatypes API

## Example Usage

```terraform
data "k8s_schemas_schemahero_io_data_type_v1alpha4_manifest" "example" {
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

- `spec` (Attributes) DataTypeSpec defines the desired state of Type (see [below for nested schema](#nestedatt--spec))

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

- `database` (String)
- `name` (String)

Optional:

- `schema` (Attributes) (see [below for nested schema](#nestedatt--spec--schema))

<a id="nestedatt--spec--schema"></a>
### Nested Schema for `spec.schema`

Optional:

- `cassandra` (Attributes) (see [below for nested schema](#nestedatt--spec--schema--cassandra))

<a id="nestedatt--spec--schema--cassandra"></a>
### Nested Schema for `spec.schema.cassandra`

Optional:

- `fields` (Attributes List) (see [below for nested schema](#nestedatt--spec--schema--cassandra--fields))
- `is_deleted` (Boolean)

<a id="nestedatt--spec--schema--cassandra--fields"></a>
### Nested Schema for `spec.schema.cassandra.fields`

Required:

- `name` (String)
- `type` (String)