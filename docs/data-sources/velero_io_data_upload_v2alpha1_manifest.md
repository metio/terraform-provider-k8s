---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_velero_io_data_upload_v2alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "velero.io"
description: |-
  DataUpload acts as the protocol between data mover plugins and data mover controller for the datamover backup operation
---

# k8s_velero_io_data_upload_v2alpha1_manifest (Data Source)

DataUpload acts as the protocol between data mover plugins and data mover controller for the datamover backup operation

## Example Usage

```terraform
data "k8s_velero_io_data_upload_v2alpha1_manifest" "example" {
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

- `spec` (Attributes) DataUploadSpec is the specification for a DataUpload. (see [below for nested schema](#nestedatt--spec))

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

- `backup_storage_location` (String) BackupStorageLocation is the name of the backup storage location where the backup repository is stored.
- `operation_timeout` (String) OperationTimeout specifies the time used to wait internal operations, before returning error as timeout.
- `snapshot_type` (String) SnapshotType is the type of the snapshot to be backed up.
- `source_namespace` (String) SourceNamespace is the original namespace where the volume is backed up from. It is the same namespace for SourcePVC and CSI namespaced objects.
- `source_pvc` (String) SourcePVC is the name of the PVC which the snapshot is taken for.

Optional:

- `cancel` (Boolean) Cancel indicates request to cancel the ongoing DataUpload. It can be set when the DataUpload is in InProgress phase
- `csi_snapshot` (Attributes) If SnapshotType is CSI, CSISnapshot provides the information of the CSI snapshot. (see [below for nested schema](#nestedatt--spec--csi_snapshot))
- `data_mover_config` (Map of String) DataMoverConfig is for data-mover-specific configuration fields.
- `datamover` (String) DataMover specifies the data mover to be used by the backup. If DataMover is '' or 'velero', the built-in data mover will be used.

<a id="nestedatt--spec--csi_snapshot"></a>
### Nested Schema for `spec.csi_snapshot`

Required:

- `storage_class` (String) StorageClass is the name of the storage class of the PVC that the volume snapshot is created from
- `volume_snapshot` (String) VolumeSnapshot is the name of the volume snapshot to be backed up

Optional:

- `snapshot_class` (String) SnapshotClass is the name of the snapshot class that the volume snapshot is created with