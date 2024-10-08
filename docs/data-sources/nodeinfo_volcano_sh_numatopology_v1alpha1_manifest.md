---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_nodeinfo_volcano_sh_numatopology_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "nodeinfo.volcano.sh"
description: |-
  Numatopology is the Schema for the Numatopologies API
---

# k8s_nodeinfo_volcano_sh_numatopology_v1alpha1_manifest (Data Source)

Numatopology is the Schema for the Numatopologies API

## Example Usage

```terraform
data "k8s_nodeinfo_volcano_sh_numatopology_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) Specification of the numa information of the worker node (see [below for nested schema](#nestedatt--spec))

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

- `cpu_detail` (Attributes) Specifies the cpu topology info Key is cpu id (see [below for nested schema](#nestedatt--spec--cpu_detail))
- `numares` (Attributes) Specifies the numa info for the resource Key is resource name (see [below for nested schema](#nestedatt--spec--numares))
- `policies` (Map of String) Specifies the policy of the manager
- `res_reserved` (Map of String) Specifies the reserved resource of the node Key is resource name

<a id="nestedatt--spec--cpu_detail"></a>
### Nested Schema for `spec.cpu_detail`

Optional:

- `core` (Number)
- `numa` (Number)
- `socket` (Number)


<a id="nestedatt--spec--numares"></a>
### Nested Schema for `spec.numares`

Optional:

- `allocatable` (String)
- `capacity` (Number)
