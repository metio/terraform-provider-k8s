---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_pkg_crossplane_io_lock_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "pkg.crossplane.io"
description: |-
  Lock is the CRD type that tracks package dependencies.
---

# k8s_pkg_crossplane_io_lock_v1beta1_manifest (Data Source)

Lock is the CRD type that tracks package dependencies.

## Example Usage

```terraform
data "k8s_pkg_crossplane_io_lock_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `packages` (Attributes List) (see [below for nested schema](#nestedatt--packages))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--packages"></a>
### Nested Schema for `packages`

Required:

- `dependencies` (Attributes List) Dependencies are the list of dependencies of this package. The order ofthe dependencies will dictate the order in which they are resolved. (see [below for nested schema](#nestedatt--packages--dependencies))
- `name` (String) Name corresponds to the name of the package revision for this package.
- `source` (String) Source is the OCI image name without a tag or digest.
- `type` (String) Type is the type of package. Can be either Configuration or Provider.
- `version` (String) Version is the tag or digest of the OCI image.

<a id="nestedatt--packages--dependencies"></a>
### Nested Schema for `packages.dependencies`

Required:

- `constraints` (String) Constraints is a valid semver range, which will be used to select a validdependency version.
- `package` (String) Package is the OCI image name without a tag or digest.
- `type` (String) Type is the type of package. Can be either Configuration or Provider.