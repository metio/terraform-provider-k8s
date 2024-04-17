---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_config_openshift_io_feature_gate_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "config.openshift.io"
description: |-
  Feature holds cluster-wide information about feature gates.  The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
---

# k8s_config_openshift_io_feature_gate_v1_manifest (Data Source)

Feature holds cluster-wide information about feature gates.  The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).

## Example Usage

```terraform
data "k8s_config_openshift_io_feature_gate_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) spec holds user settable values for configuration (see [below for nested schema](#nestedatt--spec))

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

- `custom_no_upgrade` (Attributes) customNoUpgrade allows the enabling or disabling of any feature. Turning this feature set on IS NOT SUPPORTED, CANNOT BE UNDONE, and PREVENTS UPGRADES. Because of its nature, this setting cannot be validated.  If you have any typos or accidentally apply invalid combinations your cluster may fail in an unrecoverable way.  featureSet must equal 'CustomNoUpgrade' must be set to use this field. (see [below for nested schema](#nestedatt--spec--custom_no_upgrade))
- `feature_set` (String) featureSet changes the list of features in the cluster.  The default is empty.  Be very careful adjusting this setting. Turning on or off features may cause irreversible changes in your cluster which cannot be undone.

<a id="nestedatt--spec--custom_no_upgrade"></a>
### Nested Schema for `spec.custom_no_upgrade`

Optional:

- `disabled` (List of String) disabled is a list of all feature gates that you want to force off
- `enabled` (List of String) enabled is a list of all feature gates that you want to force on