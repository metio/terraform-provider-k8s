---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_notification_toolkit_fluxcd_io_provider_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "notification.toolkit.fluxcd.io"
description: |-
  Provider is the Schema for the providers API
---

# k8s_notification_toolkit_fluxcd_io_provider_v1beta1_manifest (Data Source)

Provider is the Schema for the providers API

## Example Usage

```terraform
data "k8s_notification_toolkit_fluxcd_io_provider_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    type = "matrix"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) ProviderSpec defines the desired state of Provider (see [below for nested schema](#nestedatt--spec))

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

- `type` (String) Type of provider

Optional:

- `address` (String) HTTP/S webhook address of this provider
- `cert_secret_ref` (Attributes) CertSecretRef can be given the name of a secret containing a PEM-encoded CA certificate ('caFile') (see [below for nested schema](#nestedatt--spec--cert_secret_ref))
- `channel` (String) Alert channel for this provider
- `proxy` (String) HTTP/S address of the proxy
- `secret_ref` (Attributes) Secret reference containing the provider webhook URL using 'address' as data key (see [below for nested schema](#nestedatt--spec--secret_ref))
- `suspend` (Boolean) This flag tells the controller to suspend subsequent events handling. Defaults to false.
- `timeout` (String) Timeout for sending alerts to the provider.
- `username` (String) Bot username for this provider

<a id="nestedatt--spec--cert_secret_ref"></a>
### Nested Schema for `spec.cert_secret_ref`

Required:

- `name` (String) Name of the referent.


<a id="nestedatt--spec--secret_ref"></a>
### Nested Schema for `spec.secret_ref`

Required:

- `name` (String) Name of the referent.
