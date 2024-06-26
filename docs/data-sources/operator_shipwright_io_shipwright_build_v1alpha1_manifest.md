---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_operator_shipwright_io_shipwright_build_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "operator.shipwright.io"
description: |-
  ShipwrightBuild represents the deployment of Shipwright's build controller on a Kubernetes cluster.
---

# k8s_operator_shipwright_io_shipwright_build_v1alpha1_manifest (Data Source)

ShipwrightBuild represents the deployment of Shipwright's build controller on a Kubernetes cluster.

## Example Usage

```terraform
data "k8s_operator_shipwright_io_shipwright_build_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ShipwrightBuildSpec defines the configuration of a Shipwright Build deployment. (see [below for nested schema](#nestedatt--spec))

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

- `target_namespace` (String) TargetNamespace is the target namespace where Shipwright's build controller will be deployed.
