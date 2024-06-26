---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_spv_no_azure_key_vault_secret_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "spv.no"
description: |-
  AzureKeyVaultSecret is a specification for a AzureKeyVaultSecret resource
---

# k8s_spv_no_azure_key_vault_secret_v1alpha1_manifest (Data Source)

AzureKeyVaultSecret is a specification for a AzureKeyVaultSecret resource

## Example Usage

```terraform
data "k8s_spv_no_azure_key_vault_secret_v1alpha1_manifest" "example" {
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
- `spec` (Attributes) AzureKeyVaultSecretSpec is the spec for a AzureKeyVaultSecret resource (see [below for nested schema](#nestedatt--spec))

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

- `vault` (Attributes) AzureKeyVault contains information needed to get the Azure Key Vault secret from Azure Key Vault (see [below for nested schema](#nestedatt--spec--vault))

Optional:

- `output` (Attributes) AzureKeyVaultOutput defines output sources, currently only support Secret (see [below for nested schema](#nestedatt--spec--output))

<a id="nestedatt--spec--vault"></a>
### Nested Schema for `spec.vault`

Required:

- `name` (String) Name of the Azure Key Vault
- `object` (Attributes) AzureKeyVaultObject has information about the Azure Key Vault object to get from Azure Key Vault (see [below for nested schema](#nestedatt--spec--vault--object))

<a id="nestedatt--spec--vault--object"></a>
### Nested Schema for `spec.vault.object`

Required:

- `name` (String) The object name in Azure Key Vault
- `type` (String) AzureKeyVaultObjectType defines which Object type to get from Azure Key Vault

Optional:

- `content_type` (String) AzureKeyVaultObjectContentType defines what content type a secret contains, only used when type is multi-key-value-secret
- `version` (String) The object version in Azure Key Vault



<a id="nestedatt--spec--output"></a>
### Nested Schema for `spec.output`

Optional:

- `secret` (Attributes) AzureKeyVaultOutputSecret has information needed to output a secret from Azure Key Vault to Kubernetes as a Secret resource (see [below for nested schema](#nestedatt--spec--output--secret))
- `transforms` (List of String)

<a id="nestedatt--spec--output--secret"></a>
### Nested Schema for `spec.output.secret`

Required:

- `name` (String) Name for Kubernetes secret

Optional:

- `data_key` (String) The key to use in Kubernetes secret when setting the value from Azure Key Vault object data
- `type` (String) Type of Secret in Kubernetes
