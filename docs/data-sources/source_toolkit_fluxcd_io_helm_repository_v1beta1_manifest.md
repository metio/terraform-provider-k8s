---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "source.toolkit.fluxcd.io"
description: |-
  HelmRepository is the Schema for the helmrepositories API
---

# k8s_source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest (Data Source)

HelmRepository is the Schema for the helmrepositories API

## Example Usage

```terraform
data "k8s_source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest" "example" {
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

- `spec` (Attributes) HelmRepositorySpec defines the reference to a Helm repository. (see [below for nested schema](#nestedatt--spec))

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

- `interval` (String) The interval at which to check the upstream for updates.
- `url` (String) The Helm repository URL, a valid URL contains at least a protocol and host.

Optional:

- `access_from` (Attributes) AccessFrom defines an Access Control List for allowing cross-namespace references to this object. (see [below for nested schema](#nestedatt--spec--access_from))
- `pass_credentials` (Boolean) PassCredentials allows the credentials from the SecretRef to be passed on to a host that does not match the host as defined in URL. This may be required if the host of the advertised chart URLs in the index differ from the defined URL. Enabling this should be done with caution, as it can potentially result in credentials getting stolen in a MITM-attack.
- `secret_ref` (Attributes) The name of the secret containing authentication credentials for the Helm repository. For HTTP/S basic auth the secret must contain username and password fields. For TLS the secret must contain a certFile and keyFile, and/or caFile fields. (see [below for nested schema](#nestedatt--spec--secret_ref))
- `suspend` (Boolean) This flag tells the controller to suspend the reconciliation of this source.
- `timeout` (String) The timeout of index downloading, defaults to 60s.

<a id="nestedatt--spec--access_from"></a>
### Nested Schema for `spec.access_from`

Required:

- `namespace_selectors` (Attributes List) NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation. (see [below for nested schema](#nestedatt--spec--access_from--namespace_selectors))

<a id="nestedatt--spec--access_from--namespace_selectors"></a>
### Nested Schema for `spec.access_from.namespace_selectors`

Optional:

- `match_labels` (Map of String) MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.



<a id="nestedatt--spec--secret_ref"></a>
### Nested Schema for `spec.secret_ref`

Required:

- `name` (String) Name of the referent.
