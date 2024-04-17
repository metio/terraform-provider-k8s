---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_machineconfiguration_openshift_io_machine_config_node_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "machineconfiguration.openshift.io"
description: |-
  MachineConfigNode describes the health of the Machines on the system Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
---

# k8s_machineconfiguration_openshift_io_machine_config_node_v1alpha1_manifest (Data Source)

MachineConfigNode describes the health of the Machines on the system Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.

## Example Usage

```terraform
data "k8s_machineconfiguration_openshift_io_machine_config_node_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) spec describes the configuration of the machine config node. (see [below for nested schema](#nestedatt--spec))

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

Required:

- `config_version` (Attributes) configVersion holds the desired config version for the node targeted by this machine config node resource. The desired version represents the machine config the node will attempt to update to. This gets set before the machine config operator validates the new machine config against the current machine config. (see [below for nested schema](#nestedatt--spec--config_version))
- `node` (Attributes) node contains a reference to the node for this machine config node. (see [below for nested schema](#nestedatt--spec--node))
- `pool` (Attributes) pool contains a reference to the machine config pool that this machine config node's referenced node belongs to. (see [below for nested schema](#nestedatt--spec--pool))

<a id="nestedatt--spec--config_version"></a>
### Nested Schema for `spec.config_version`

Required:

- `desired` (String) desired is the name of the machine config that the the node should be upgraded to. This value is set when the machine config pool generates a new version of its rendered configuration. When this value is changed, the machine config daemon starts the node upgrade process. This value gets set in the machine config node spec once the machine config has been targeted for upgrade and before it is validated. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.


<a id="nestedatt--spec--node"></a>
### Nested Schema for `spec.node`

Required:

- `name` (String) name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.


<a id="nestedatt--spec--pool"></a>
### Nested Schema for `spec.pool`

Required:

- `name` (String) name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.