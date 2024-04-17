---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_azure_microsoft_com_my_sql_server_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "azure.microsoft.com"
description: |-
  MySQLServer is the Schema for the mysqlservers API
---

# k8s_azure_microsoft_com_my_sql_server_v1alpha1_manifest (Data Source)

MySQLServer is the Schema for the mysqlservers API

## Example Usage

```terraform
data "k8s_azure_microsoft_com_my_sql_server_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) MySQLServerSpec defines the desired state of MySQLServer (see [below for nested schema](#nestedatt--spec))

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

- `location` (String)
- `resource_group` (String)

Optional:

- `create_mode` (String)
- `key_vault_to_store_secrets` (String)
- `replica_properties` (Attributes) (see [below for nested schema](#nestedatt--spec--replica_properties))
- `server_version` (String) ServerVersion enumerates the values for server version.
- `sku` (Attributes) (see [below for nested schema](#nestedatt--spec--sku))
- `ssl_enforcement` (String)

<a id="nestedatt--spec--replica_properties"></a>
### Nested Schema for `spec.replica_properties`

Optional:

- `source_server_id` (String)


<a id="nestedatt--spec--sku"></a>
### Nested Schema for `spec.sku`

Optional:

- `capacity` (Number) Capacity - The scale up/out capacity, representing server's compute units.
- `family` (String) Family - The family of hardware.
- `name` (String) Name - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
- `size` (String) Size - The size code, to be interpreted by resource as appropriate.
- `tier` (String) Tier - The tier of the particular SKU, e.g. Basic. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'