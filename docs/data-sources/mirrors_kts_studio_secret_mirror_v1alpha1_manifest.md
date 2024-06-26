---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_mirrors_kts_studio_secret_mirror_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "mirrors.kts.studio"
description: |-
  SecretMirror is the Schema for the secretmirrors API
---

# k8s_mirrors_kts_studio_secret_mirror_v1alpha1_manifest (Data Source)

SecretMirror is the Schema for the secretmirrors API

## Example Usage

```terraform
data "k8s_mirrors_kts_studio_secret_mirror_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) SecretMirrorSpec defines the desired state of SecretMirror (see [below for nested schema](#nestedatt--spec))

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

- `destination` (Attributes) (see [below for nested schema](#nestedatt--spec--destination))
- `poll_period_seconds` (Number)
- `source` (Attributes) (see [below for nested schema](#nestedatt--spec--source))

<a id="nestedatt--spec--destination"></a>
### Nested Schema for `spec.destination`

Optional:

- `namespace` (String)
- `namespace_regex` (String)


<a id="nestedatt--spec--source"></a>
### Nested Schema for `spec.source`

Optional:

- `name` (String)
