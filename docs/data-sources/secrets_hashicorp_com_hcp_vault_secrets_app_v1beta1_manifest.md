---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "secrets.hashicorp.com"
description: |-
  HCPVaultSecretsApp is the Schema for the hcpvaultsecretsapps API
---

# k8s_secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest (Data Source)

HCPVaultSecretsApp is the Schema for the hcpvaultsecretsapps API

## Example Usage

```terraform
data "k8s_secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest" "example" {
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

- `spec` (Attributes) HCPVaultSecretsAppSpec defines the desired state of HCPVaultSecretsApp (see [below for nested schema](#nestedatt--spec))

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

- `app_name` (String) AppName of the Vault Secrets Application that is to be synced.
- `destination` (Attributes) Destination provides configuration necessary for syncing the HCP Vault Application secrets to Kubernetes. (see [below for nested schema](#nestedatt--spec--destination))

Optional:

- `hcp_auth_ref` (String) HCPAuthRef to the HCPAuth resource, can be prefixed with a namespace, eg: 'namespaceA/vaultAuthRefB'. If no namespace prefix is provided it will default to the namespace of the HCPAuth CR. If no value is specified for HCPAuthRef the Operator will default to the 'default' HCPAuth, configured in the operator's namespace.
- `refresh_after` (String) RefreshAfter a period of time, in duration notation e.g. 30s, 1m, 24h
- `rollout_restart_targets` (Attributes List) RolloutRestartTargets should be configured whenever the application(s) consuming the HCP Vault Secrets App does not support dynamically reloading a rotated secret. In that case one, or more RolloutRestartTarget(s) can be configured here. The Operator will trigger a 'rollout-restart' for each target whenever the Vault secret changes between reconciliation events. See RolloutRestartTarget for more details. (see [below for nested schema](#nestedatt--spec--rollout_restart_targets))

<a id="nestedatt--spec--destination"></a>
### Nested Schema for `spec.destination`

Required:

- `name` (String) Name of the Secret

Optional:

- `annotations` (Map of String) Annotations to apply to the Secret. Requires Create to be set to true.
- `create` (Boolean) Create the destination Secret. If the Secret already exists this should be set to false.
- `labels` (Map of String) Labels to apply to the Secret. Requires Create to be set to true.
- `overwrite` (Boolean) Overwrite the destination Secret if it exists and Create is true. This is useful when migrating to VSO from a previous secret deployment strategy.
- `transformation` (Attributes) Transformation provides configuration for transforming the secret data before it is stored in the Destination. (see [below for nested schema](#nestedatt--spec--destination--transformation))
- `type` (String) Type of Kubernetes Secret. Requires Create to be set to true. Defaults to Opaque.

<a id="nestedatt--spec--destination--transformation"></a>
### Nested Schema for `spec.destination.transformation`

Optional:

- `exclude_raw` (Boolean) ExcludeRaw data from the destination Secret. Exclusion policy can be set globally by including 'exclude-raw' in the '--global-transformation-options' command line flag. If set, the command line flag always takes precedence over this configuration.
- `excludes` (List of String) Excludes contains regex patterns used to filter top-level source secret data fields for exclusion from the final K8s Secret data. These pattern filters are never applied to templated fields as defined in Templates. They are always applied before any inclusion patterns. To exclude all source secret data fields, you can configure the single pattern '.*'.
- `includes` (List of String) Includes contains regex patterns used to filter top-level source secret data fields for inclusion in the final K8s Secret data. These pattern filters are never applied to templated fields as defined in Templates. They are always applied last.
- `templates` (Attributes) Templates maps a template name to its Template. Templates are always included in the rendered K8s Secret, and take precedence over templates defined in a SecretTransformation. (see [below for nested schema](#nestedatt--spec--destination--transformation--templates))
- `transformation_refs` (Attributes List) TransformationRefs contain references to template configuration from SecretTransformation. (see [below for nested schema](#nestedatt--spec--destination--transformation--transformation_refs))

<a id="nestedatt--spec--destination--transformation--templates"></a>
### Nested Schema for `spec.destination.transformation.templates`

Required:

- `text` (String) Text contains the Go text template format. The template references attributes from the data structure of the source secret. Refer to https://pkg.go.dev/text/template for more information.

Optional:

- `name` (String) Name of the Template


<a id="nestedatt--spec--destination--transformation--transformation_refs"></a>
### Nested Schema for `spec.destination.transformation.transformation_refs`

Required:

- `name` (String) Name of the SecretTransformation resource.

Optional:

- `ignore_excludes` (Boolean) IgnoreExcludes controls whether to use the SecretTransformation's Excludes data key filters.
- `ignore_includes` (Boolean) IgnoreIncludes controls whether to use the SecretTransformation's Includes data key filters.
- `namespace` (String) Namespace of the SecretTransformation resource.
- `template_refs` (Attributes List) TemplateRefs map to a Template found in this TransformationRef. If empty, then all templates from the SecretTransformation will be rendered to the K8s Secret. (see [below for nested schema](#nestedatt--spec--destination--transformation--transformation_refs--template_refs))

<a id="nestedatt--spec--destination--transformation--transformation_refs--template_refs"></a>
### Nested Schema for `spec.destination.transformation.transformation_refs.template_refs`

Required:

- `name` (String) Name of the Template in SecretTransformationSpec.Templates. the rendered secret data.

Optional:

- `key_override` (String) KeyOverride to the rendered template in the Destination secret. If Key is empty, then the Key from reference spec will be used. Set this to override the Key set from the reference spec.





<a id="nestedatt--spec--rollout_restart_targets"></a>
### Nested Schema for `spec.rollout_restart_targets`

Required:

- `kind` (String)
- `name` (String)