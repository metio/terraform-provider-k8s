---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_longhorn_io_backup_target_v1beta2_manifest Data Source - terraform-provider-k8s"
subcategory: "longhorn.io"
description: |-
  BackupTarget is where Longhorn stores backup target object.
---

# k8s_longhorn_io_backup_target_v1beta2_manifest (Data Source)

BackupTarget is where Longhorn stores backup target object.

## Example Usage

```terraform
data "k8s_longhorn_io_backup_target_v1beta2_manifest" "example" {
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

- `spec` (Attributes) BackupTargetSpec defines the desired state of the Longhorn backup target (see [below for nested schema](#nestedatt--spec))

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

- `backup_target_url` (String) The backup target URL.
- `credential_secret` (String) The backup target credential secret.
- `poll_interval` (String) The interval that the cluster needs to run sync with the backup target.
- `sync_requested_at` (String) The time to request run sync the remote backup target.
