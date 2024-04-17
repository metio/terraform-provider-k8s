---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_image_toolkit_fluxcd_io_image_policy_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "image.toolkit.fluxcd.io"
description: |-
  ImagePolicy is the Schema for the imagepolicies API
---

# k8s_image_toolkit_fluxcd_io_image_policy_v1beta1_manifest (Data Source)

ImagePolicy is the Schema for the imagepolicies API

## Example Usage

```terraform
data "k8s_image_toolkit_fluxcd_io_image_policy_v1beta1_manifest" "example" {
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

- `spec` (Attributes) ImagePolicySpec defines the parameters for calculating theImagePolicy (see [below for nested schema](#nestedatt--spec))

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

- `image_repository_ref` (Attributes) ImageRepositoryRef points at the object specifying the imagebeing scanned (see [below for nested schema](#nestedatt--spec--image_repository_ref))
- `policy` (Attributes) Policy gives the particulars of the policy to be followed inselecting the most recent image (see [below for nested schema](#nestedatt--spec--policy))

Optional:

- `filter_tags` (Attributes) FilterTags enables filtering for only a subset of tags based on a set ofrules. If no rules are provided, all the tags from the repository will beordered and compared. (see [below for nested schema](#nestedatt--spec--filter_tags))

<a id="nestedatt--spec--image_repository_ref"></a>
### Nested Schema for `spec.image_repository_ref`

Required:

- `name` (String) Name of the referent.

Optional:

- `namespace` (String) Namespace of the referent, when not specified it acts as LocalObjectReference.


<a id="nestedatt--spec--policy"></a>
### Nested Schema for `spec.policy`

Optional:

- `alphabetical` (Attributes) Alphabetical set of rules to use for alphabetical ordering of the tags. (see [below for nested schema](#nestedatt--spec--policy--alphabetical))
- `numerical` (Attributes) Numerical set of rules to use for numerical ordering of the tags. (see [below for nested schema](#nestedatt--spec--policy--numerical))
- `semver` (Attributes) SemVer gives a semantic version range to check against the tagsavailable. (see [below for nested schema](#nestedatt--spec--policy--semver))

<a id="nestedatt--spec--policy--alphabetical"></a>
### Nested Schema for `spec.policy.alphabetical`

Optional:

- `order` (String) Order specifies the sorting order of the tags. Given the letters of thealphabet as tags, ascending order would select Z, and descending orderwould select A.


<a id="nestedatt--spec--policy--numerical"></a>
### Nested Schema for `spec.policy.numerical`

Optional:

- `order` (String) Order specifies the sorting order of the tags. Given the integer valuesfrom 0 to 9 as tags, ascending order would select 9, and descending orderwould select 0.


<a id="nestedatt--spec--policy--semver"></a>
### Nested Schema for `spec.policy.semver`

Required:

- `range` (String) Range gives a semver range for the image tag; the highestversion within the range that's a tag yields the latest image.



<a id="nestedatt--spec--filter_tags"></a>
### Nested Schema for `spec.filter_tags`

Optional:

- `extract` (String) Extract allows a capture group to be extracted from the specified regularexpression pattern, useful before tag evaluation.
- `pattern` (String) Pattern specifies a regular expression pattern used to filter for imagetags.