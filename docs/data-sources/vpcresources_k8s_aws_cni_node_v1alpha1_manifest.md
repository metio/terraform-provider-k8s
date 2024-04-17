---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_vpcresources_k8s_aws_cni_node_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "vpcresources.k8s.aws"
description: |-
  
---

# k8s_vpcresources_k8s_aws_cni_node_v1alpha1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_vpcresources_k8s_aws_cni_node_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) Important: Run 'make' to regenerate code after modifying this file CNINodeSpec defines the desired state of CNINode (see [below for nested schema](#nestedatt--spec))

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

- `features` (Attributes List) (see [below for nested schema](#nestedatt--spec--features))

<a id="nestedatt--spec--features"></a>
### Nested Schema for `spec.features`

Optional:

- `name` (String) FeatureName is a type of feature name supported by AWS VPC CNI. It can be Security Group for Pods, custom networking, or others
- `value` (String)