---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_postgresql_cnpg_io_scheduled_backup_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "postgresql.cnpg.io"
description: |-
  ScheduledBackup is the Schema for the scheduledbackups API
---

# k8s_postgresql_cnpg_io_scheduled_backup_v1_manifest (Data Source)

ScheduledBackup is the Schema for the scheduledbackups API

## Example Usage

```terraform
data "k8s_postgresql_cnpg_io_scheduled_backup_v1_manifest" "example" {
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
- `spec` (Attributes) Specification of the desired behavior of the ScheduledBackup.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status (see [below for nested schema](#nestedatt--spec))

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

- `cluster` (Attributes) The cluster to backup (see [below for nested schema](#nestedatt--spec--cluster))
- `schedule` (String) The schedule does not follow the same format used in Kubernetes CronJobsas it includes an additional seconds specifier,see https://pkg.go.dev/github.com/robfig/cron#hdr-CRON_Expression_Format

Optional:

- `backup_owner_reference` (String) Indicates which ownerReference should be put inside the created backup resources.<br />- none: no owner reference for created backup objects (same behavior as before the field was introduced)<br />- self: sets the Scheduled backup object as owner of the backup<br />- cluster: set the cluster as owner of the backup<br />
- `immediate` (Boolean) If the first backup has to be immediately start after creation or not
- `method` (String) The backup method to be used, possible options are 'barmanObjectStore'and 'volumeSnapshot'. Defaults to: 'barmanObjectStore'.
- `online` (Boolean) Whether the default type of backup with volume snapshots isonline/hot ('true', default) or offline/cold ('false')Overrides the default setting specified in the cluster field '.spec.backup.volumeSnapshot.online'
- `online_configuration` (Attributes) Configuration parameters to control the online/hot backup with volume snapshotsOverrides the default settings specified in the cluster '.backup.volumeSnapshot.onlineConfiguration' stanza (see [below for nested schema](#nestedatt--spec--online_configuration))
- `plugin_configuration` (Attributes) Configuration parameters passed to the plugin managing this backup (see [below for nested schema](#nestedatt--spec--plugin_configuration))
- `suspend` (Boolean) If this backup is suspended or not
- `target` (String) The policy to decide which instance should perform this backup. If empty,it defaults to 'cluster.spec.backup.target'.Available options are empty string, 'primary' and 'prefer-standby'.'primary' to have backups run always on primary instances,'prefer-standby' to have backups run preferably on the most updatedstandby, if available.

<a id="nestedatt--spec--cluster"></a>
### Nested Schema for `spec.cluster`

Required:

- `name` (String) Name of the referent.


<a id="nestedatt--spec--online_configuration"></a>
### Nested Schema for `spec.online_configuration`

Optional:

- `immediate_checkpoint` (Boolean) Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.
- `wait_for_archive` (Boolean) If false, the function will return immediately after the backup is completed,without waiting for WAL to be archived.This behavior is only useful with backup software that independently monitors WAL archiving.Otherwise, WAL required to make the backup consistent might be missing and make the backup useless.By default, or when this parameter is true, pg_backup_stop will wait for WAL to be archived when archiving isenabled.On a standby, this means that it will wait only when archive_mode = always.If write activity on the primary is low, it may be useful to run pg_switch_wal on the primary in order to triggeran immediate segment switch.


<a id="nestedatt--spec--plugin_configuration"></a>
### Nested Schema for `spec.plugin_configuration`

Required:

- `name` (String) Name is the name of the plugin managing this backup

Optional:

- `parameters` (Map of String) Parameters are the configuration parameters passed to the backupplugin for this backup