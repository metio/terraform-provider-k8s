---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apps_clusternet_io_description_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "apps.clusternet.io"
description: |-
  Description is the Schema for the resources to be installed
---

# k8s_apps_clusternet_io_description_v1alpha1_manifest (Data Source)

Description is the Schema for the resources to be installed

## Example Usage

```terraform
data "k8s_apps_clusternet_io_description_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    deployer = "Helm"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) DescriptionSpec defines the spec of Description (see [below for nested schema](#nestedatt--spec))

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

- `deployer` (String) Deployer indicates the deployer for this Description

Optional:

- `chart_raw` (List of String) ChartRaw is the underlying serialization of all helm chart objects.
- `charts` (Attributes List) Charts describe all the helm charts to be installed (see [below for nested schema](#nestedatt--spec--charts))
- `raw` (List of String) Raw is the underlying serialization of all objects.

<a id="nestedatt--spec--charts"></a>
### Nested Schema for `spec.charts`

Required:

- `name` (String) Name of the HelmChart.
- `namespace` (String) Namespace of the HelmChart.