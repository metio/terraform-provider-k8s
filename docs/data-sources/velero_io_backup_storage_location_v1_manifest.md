---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_velero_io_backup_storage_location_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "velero.io"
description: |-
  BackupStorageLocation is a location where Velero stores backup objects
---

# k8s_velero_io_backup_storage_location_v1_manifest (Data Source)

BackupStorageLocation is a location where Velero stores backup objects

## Example Usage

```terraform
data "k8s_velero_io_backup_storage_location_v1_manifest" "example" {
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

- `spec` (Attributes) BackupStorageLocationSpec defines the desired state of a Velero BackupStorageLocation (see [below for nested schema](#nestedatt--spec))

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

- `object_storage` (Attributes) ObjectStorageLocation specifies the settings necessary to connect to a provider's object storage. (see [below for nested schema](#nestedatt--spec--object_storage))
- `provider` (String) Provider is the provider of the backup storage.

Optional:

- `access_mode` (String) AccessMode defines the permissions for the backup storage location.
- `backup_sync_period` (String) BackupSyncPeriod defines how frequently to sync backup API objects from object storage. A value of 0 disables sync.
- `config` (Map of String) Config is for provider-specific configuration fields.
- `credential` (Attributes) Credential contains the credential information intended to be used with this location (see [below for nested schema](#nestedatt--spec--credential))
- `default` (Boolean) Default indicates this location is the default backup storage location.
- `validation_frequency` (String) ValidationFrequency defines how frequently to validate the corresponding object storage. A value of 0 disables validation.

<a id="nestedatt--spec--object_storage"></a>
### Nested Schema for `spec.object_storage`

Required:

- `bucket` (String) Bucket is the bucket to use for object storage.

Optional:

- `ca_cert` (String) CACert defines a CA bundle to use when verifying TLS connections to the provider.
- `prefix` (String) Prefix is the path inside a bucket to use for Velero storage. Optional.


<a id="nestedatt--spec--credential"></a>
### Nested Schema for `spec.credential`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined