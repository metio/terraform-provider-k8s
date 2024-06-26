---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_networking_istio_io_proxy_config_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "networking.istio.io"
description: |-
  
---

# k8s_networking_istio_io_proxy_config_v1beta1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_networking_istio_io_proxy_config_v1beta1_manifest" "example" {
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

- `spec` (Attributes) Provides configuration for individual workloads. See more details at: https://istio.io/docs/reference/config/networking/proxy-config.html (see [below for nested schema](#nestedatt--spec))

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

- `concurrency` (Number) The number of worker threads to run.
- `environment_variables` (Map of String) Additional environment variables for the proxy.
- `image` (Attributes) Specifies the details of the proxy image. (see [below for nested schema](#nestedatt--spec--image))
- `selector` (Attributes) Optional. (see [below for nested schema](#nestedatt--spec--selector))

<a id="nestedatt--spec--image"></a>
### Nested Schema for `spec.image`

Optional:

- `image_type` (String) The image type of the image.


<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- `match_labels` (Map of String) One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.
