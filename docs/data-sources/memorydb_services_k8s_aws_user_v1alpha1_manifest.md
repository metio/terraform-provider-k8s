---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_memorydb_services_k8s_aws_user_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "memorydb.services.k8s.aws"
description: |-
  User is the Schema for the Users API
---

# k8s_memorydb_services_k8s_aws_user_v1alpha1_manifest (Data Source)

User is the Schema for the Users API

## Example Usage

```terraform
data "k8s_memorydb_services_k8s_aws_user_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) UserSpec defines the desired state of User.  You create users and assign them specific permissions by using an access string. You assign the users to Access Control Lists aligned with a specific role (administrators, human resources) that are then deployed to one or more MemoryDB clusters. (see [below for nested schema](#nestedatt--spec))

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

- `access_string` (String) Access permissions string used for this user.
- `authentication_mode` (Attributes) Denotes the user's authentication properties, such as whether it requires a password to authenticate. (see [below for nested schema](#nestedatt--spec--authentication_mode))
- `name` (String) The name of the user. This value must be unique as it also serves as the user identifier.

Optional:

- `tags` (Attributes List) A list of tags to be added to this resource. A tag is a key-value pair. A tag key must be accompanied by a tag value, although null is accepted. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--authentication_mode"></a>
### Nested Schema for `spec.authentication_mode`

Optional:

- `passwords` (Attributes List) (see [below for nested schema](#nestedatt--spec--authentication_mode--passwords))
- `type_` (String)

<a id="nestedatt--spec--authentication_mode--passwords"></a>
### Nested Schema for `spec.authentication_mode.passwords`

Required:

- `key` (String) Key is the key within the secret

Optional:

- `name` (String) name is unique within a namespace to reference a secret resource.
- `namespace` (String) namespace defines the space within which the secret name must be unique.



<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)