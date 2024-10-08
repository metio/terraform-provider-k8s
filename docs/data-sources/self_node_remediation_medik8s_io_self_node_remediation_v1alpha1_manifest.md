---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "self-node-remediation.medik8s.io"
description: |-
  SelfNodeRemediation is the Schema for the selfnoderemediations API
---

# k8s_self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest (Data Source)

SelfNodeRemediation is the Schema for the selfnoderemediations API

## Example Usage

```terraform
data "k8s_self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) SelfNodeRemediationSpec defines the desired state of SelfNodeRemediation (see [below for nested schema](#nestedatt--spec))

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

- `remediation_strategy` (String) RemediationStrategy is the remediation method for unhealthy nodes. Currently, it could be either 'Automatic', 'OutOfServiceTaint' or 'ResourceDeletion'. ResourceDeletion will iterate over all pods and VolumeAttachment related to the unhealthy node and delete them. OutOfServiceTaint will add the out-of-service taint which is a new well-known taint 'node.kubernetes.io/out-of-service' that enables automatic deletion of pv-attached pods on failed nodes, 'out-of-service' taint is only supported on clusters with k8s version 1.26+ or OCP/OKD version 4.13+. Automatic will choose the most appropriate strategy during runtime.
