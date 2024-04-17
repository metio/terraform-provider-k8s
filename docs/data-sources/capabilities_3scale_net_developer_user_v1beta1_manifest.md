---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_capabilities_3scale_net_developer_user_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "capabilities.3scale.net"
description: |-
  DeveloperUser is the Schema for the developerusers API
---

# k8s_capabilities_3scale_net_developer_user_v1beta1_manifest (Data Source)

DeveloperUser is the Schema for the developerusers API

## Example Usage

```terraform
data "k8s_capabilities_3scale_net_developer_user_v1beta1_manifest" "example" {
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

- `spec` (Attributes) DeveloperUserSpec defines the desired state of DeveloperUser (see [below for nested schema](#nestedatt--spec))

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

- `developer_account_ref` (Attributes) DeveloperAccountRef is the reference to the parent developer account (see [below for nested schema](#nestedatt--spec--developer_account_ref))
- `email` (String) Email
- `password_credentials_ref` (Attributes) Password (see [below for nested schema](#nestedatt--spec--password_credentials_ref))
- `username` (String) Username

Optional:

- `provider_account_ref` (Attributes) ProviderAccountRef references account provider credentials (see [below for nested schema](#nestedatt--spec--provider_account_ref))
- `role` (String) Role defines the desired 3scale role. Defaults to 'member'
- `suspended` (Boolean) State defines the desired state. Defaults to 'false', ie, active

<a id="nestedatt--spec--developer_account_ref"></a>
### Nested Schema for `spec.developer_account_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--password_credentials_ref"></a>
### Nested Schema for `spec.password_credentials_ref`

Optional:

- `name` (String) name is unique within a namespace to reference a secret resource.
- `namespace` (String) namespace defines the space within which the secret name must be unique.


<a id="nestedatt--spec--provider_account_ref"></a>
### Nested Schema for `spec.provider_account_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?