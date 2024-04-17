---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest Data Source - terraform-provider-k8s"
subcategory: "azure.microsoft.com"
description: |-
  MySQLAADUser is the Schema for an AAD user for MySQL
---

# k8s_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest (Data Source)

MySQLAADUser is the Schema for an AAD user for MySQL

## Example Usage

```terraform
data "k8s_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest" "example" {
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

- `spec` (Attributes) MySQLAADUserSpec defines the desired state of MySQLAADUser (see [below for nested schema](#nestedatt--spec))

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

- `resource_group` (String)
- `roles` (List of String) The server-level roles assigned to the user.
- `server` (String)

Optional:

- `aad_id` (String) AAD ID is the ID of the user in Azure Active Directory. When creating a user for a managed identity this must be the client id (sometimes called app id) of the managed identity. When creating a user for a 'normal' (non-managed identity) user or group, this is the OID of the user or group.
- `database_roles` (Map of List of String) The database-level roles assigned to the user (keyed by database name).
- `username` (String)