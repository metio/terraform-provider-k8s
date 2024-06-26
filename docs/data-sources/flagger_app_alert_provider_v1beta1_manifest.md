---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_flagger_app_alert_provider_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "flagger.app"
description: |-
  AlertProvider is the Schema for the AlertProvider API.
---

# k8s_flagger_app_alert_provider_v1beta1_manifest (Data Source)

AlertProvider is the Schema for the AlertProvider API.

## Example Usage

```terraform
data "k8s_flagger_app_alert_provider_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    type    = "slack"
    address = "https://example.com"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) AlertProviderSpec defines the desired state of a AlertProvider. (see [below for nested schema](#nestedatt--spec))

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

- `address` (String) Hook URL address of this provider
- `channel` (String) Alert channel for this provider
- `proxy` (String) Http/s proxy of this provider
- `secret_ref` (Attributes) Kubernetes secret reference containing the provider address (see [below for nested schema](#nestedatt--spec--secret_ref))
- `type` (String) Type of this provider
- `username` (String) Bot username for this provider

<a id="nestedatt--spec--secret_ref"></a>
### Nested Schema for `spec.secret_ref`

Required:

- `name` (String) Name of the Kubernetes secret
