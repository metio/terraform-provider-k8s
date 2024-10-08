---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_operator_tigera_io_image_set_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "operator.tigera.io"
description: |-
  ImageSet is used to specify image digests for the images that the operator deploys. The name of the ImageSet is expected to be in the format '-'. The 'variant' used is 'enterprise' if the InstallationSpec Variant is 'TigeraSecureEnterprise' otherwise it is 'calico'. The 'release' must match the version of the variant that the operator is built to deploy, this version can be obtained by passing the '--version' flag to the operator binary.
---

# k8s_operator_tigera_io_image_set_v1_manifest (Data Source)

ImageSet is used to specify image digests for the images that the operator deploys. The name of the ImageSet is expected to be in the format '<variant>-<release>'. The 'variant' used is 'enterprise' if the InstallationSpec Variant is 'TigeraSecureEnterprise' otherwise it is 'calico'. The 'release' must match the version of the variant that the operator is built to deploy, this version can be obtained by passing the '--version' flag to the operator binary.

## Example Usage

```terraform
data "k8s_operator_tigera_io_image_set_v1_manifest" "example" {
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

- `spec` (Attributes) ImageSetSpec defines the desired state of ImageSet. (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `images` (Attributes List) Images is the list of images to use digests. All images that the operator will deploy must be specified. (see [below for nested schema](#nestedatt--spec--images))

<a id="nestedatt--spec--images"></a>
### Nested Schema for `spec.images`

Required:

- `digest` (String) Digest is the image identifier that will be used for the Image. The field should not include a leading '@' and must be prefixed with 'sha256:'.
- `image` (String) Image is an image that the operator deploys and instead of using the built in tag the operator will use the Digest for the image identifier. The value should be the image name without registry or tag or digest. For the image 'docker.io/calico/node:v3.17.1' it should be represented as 'calico/node'
