---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_velero_io_volume_snapshot_location_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "velero.io"
description: |-
  VolumeSnapshotLocation is a location where Velero stores volume snapshots.
---

# k8s_velero_io_volume_snapshot_location_v1_manifest (Data Source)

VolumeSnapshotLocation is a location where Velero stores volume snapshots.

## Example Usage

```terraform
data "k8s_velero_io_volume_snapshot_location_v1_manifest" "example" {
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

- `spec` (Attributes) VolumeSnapshotLocationSpec defines the specification for a Velero VolumeSnapshotLocation. (see [below for nested schema](#nestedatt--spec))

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

- `provider` (String) Provider is the provider of the volume storage.

Optional:

- `config` (Map of String) Config is for provider-specific configuration fields.
- `credential` (Attributes) Credential contains the credential information intended to be used with this location (see [below for nested schema](#nestedatt--spec--credential))

<a id="nestedatt--spec--credential"></a>
### Nested Schema for `spec.credential`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined