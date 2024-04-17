---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_trust_cert_manager_io_bundle_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "trust.cert-manager.io"
description: |-
  
---

# k8s_trust_cert_manager_io_bundle_v1alpha1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_trust_cert_manager_io_bundle_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Desired state of the Bundle resource. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `sources` (Attributes List) Sources is a set of references to data whose data will sync to the target. (see [below for nested schema](#nestedatt--spec--sources))
- `target` (Attributes) Target is the target location in all namespaces to sync source data to. (see [below for nested schema](#nestedatt--spec--target))

<a id="nestedatt--spec--sources"></a>
### Nested Schema for `spec.sources`

Optional:

- `config_map` (Attributes) ConfigMap is a reference (by name) to a ConfigMap's 'data' key, or to alist of ConfigMap's 'data' key using label selector, in the trust Namespace. (see [below for nested schema](#nestedatt--spec--sources--config_map))
- `in_line` (String) InLine is a simple string to append as the source data.
- `secret` (Attributes) Secret is a reference (by name) to a Secret's 'data' key, or to alist of Secret's 'data' key using label selector, in the trust Namespace. (see [below for nested schema](#nestedatt--spec--sources--secret))
- `use_default_c_as` (Boolean) UseDefaultCAs, when true, requests the default CA bundle to be used as a source.Default CAs are available if trust-manager was installed via Helmor was otherwise set up to include a package-injecting init container by using the'--default-package-location' flag when starting the trust-manager controller.If default CAs were not configured at start-up, any request to use the defaultCAs will fail.The version of the default CA package which is used for a Bundle is stored in thedefaultCAPackageVersion field of the Bundle's status field.

<a id="nestedatt--spec--sources--config_map"></a>
### Nested Schema for `spec.sources.config_map`

Required:

- `key` (String) Key is the key of the entry in the object's 'data' field to be used.

Optional:

- `name` (String) Name is the name of the source object in the trust Namespace.This field must be left empty when 'selector' is set
- `selector` (Attributes) Selector is the label selector to use to fetch a list of objects. Must not be setwhen 'Name' is set. (see [below for nested schema](#nestedatt--spec--sources--config_map--selector))

<a id="nestedatt--spec--sources--config_map--selector"></a>
### Nested Schema for `spec.sources.config_map.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--sources--config_map--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--sources--config_map--selector--match_expressions"></a>
### Nested Schema for `spec.sources.config_map.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.




<a id="nestedatt--spec--sources--secret"></a>
### Nested Schema for `spec.sources.secret`

Required:

- `key` (String) Key is the key of the entry in the object's 'data' field to be used.

Optional:

- `name` (String) Name is the name of the source object in the trust Namespace.This field must be left empty when 'selector' is set
- `selector` (Attributes) Selector is the label selector to use to fetch a list of objects. Must not be setwhen 'Name' is set. (see [below for nested schema](#nestedatt--spec--sources--secret--selector))

<a id="nestedatt--spec--sources--secret--selector"></a>
### Nested Schema for `spec.sources.secret.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--sources--secret--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--sources--secret--selector--match_expressions"></a>
### Nested Schema for `spec.sources.secret.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.





<a id="nestedatt--spec--target"></a>
### Nested Schema for `spec.target`

Optional:

- `additional_formats` (Attributes) AdditionalFormats specifies any additional formats to write to the target (see [below for nested schema](#nestedatt--spec--target--additional_formats))
- `config_map` (Attributes) ConfigMap is the target ConfigMap in Namespaces that all Bundle sourcedata will be synced to. (see [below for nested schema](#nestedatt--spec--target--config_map))
- `namespace_selector` (Attributes) NamespaceSelector will, if set, only sync the target resource inNamespaces which match the selector. (see [below for nested schema](#nestedatt--spec--target--namespace_selector))
- `secret` (Attributes) Secret is the target Secret that all Bundle source data will be synced to.Using Secrets as targets is only supported if enabled at trust-manager startup.By default, trust-manager has no permissions for writing to secrets and can only read secrets in the trust namespace. (see [below for nested schema](#nestedatt--spec--target--secret))

<a id="nestedatt--spec--target--additional_formats"></a>
### Nested Schema for `spec.target.additional_formats`

Optional:

- `jks` (Attributes) JKS requests a JKS-formatted binary trust bundle to be written to the target.The bundle has 'changeit' as the default password.For more information refer to this link https://cert-manager.io/docs/faq/#keystore-passwords (see [below for nested schema](#nestedatt--spec--target--additional_formats--jks))
- `pkcs12` (Attributes) PKCS12 requests a PKCS12-formatted binary trust bundle to be written to the target.The bundle is by default created without a password. (see [below for nested schema](#nestedatt--spec--target--additional_formats--pkcs12))

<a id="nestedatt--spec--target--additional_formats--jks"></a>
### Nested Schema for `spec.target.additional_formats.jks`

Required:

- `key` (String) Key is the key of the entry in the object's 'data' field to be used.

Optional:

- `password` (String) Password for JKS trust store


<a id="nestedatt--spec--target--additional_formats--pkcs12"></a>
### Nested Schema for `spec.target.additional_formats.pkcs12`

Required:

- `key` (String) Key is the key of the entry in the object's 'data' field to be used.

Optional:

- `password` (String) Password for PKCS12 trust store



<a id="nestedatt--spec--target--config_map"></a>
### Nested Schema for `spec.target.config_map`

Required:

- `key` (String) Key is the key of the entry in the object's 'data' field to be used.


<a id="nestedatt--spec--target--namespace_selector"></a>
### Nested Schema for `spec.target.namespace_selector`

Optional:

- `match_labels` (Map of String) MatchLabels matches on the set of labels that must be present on aNamespace for the Bundle target to be synced there.


<a id="nestedatt--spec--target--secret"></a>
### Nested Schema for `spec.target.secret`

Required:

- `key` (String) Key is the key of the entry in the object's 'data' field to be used.