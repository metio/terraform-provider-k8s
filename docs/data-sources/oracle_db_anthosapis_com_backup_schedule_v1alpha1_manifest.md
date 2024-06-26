---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "oracle.db.anthosapis.com"
description: |-
  BackupSchedule is the Schema for the backupschedules API.
---

# k8s_oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest (Data Source)

BackupSchedule is the Schema for the backupschedules API.

## Example Usage

```terraform
data "k8s_oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

- `backup_spec` (Attributes) BackupSpec defines the Backup that will be created on the provided schedule. (see [below for nested schema](#nestedatt--spec--backup_spec))
- `schedule` (String) Schedule is a cron-style expression of the schedule on which Backup will be created. For allowed syntax, see en.wikipedia.org/wiki/Cron and godoc.org/github.com/robfig/cron.

Optional:

- `backup_labels` (Map of String) BackupLabels define the desired labels that scheduled backups will be created with.
- `backup_retention_policy` (Attributes) BackupRetentionPolicy is the policy used to trigger automatic deletion of backups produced from this BackupSchedule. (see [below for nested schema](#nestedatt--spec--backup_retention_policy))
- `starting_deadline_seconds` (Number) StartingDeadlineSeconds is an optional deadline in seconds for starting the backup creation if it misses scheduled time for any reason. The default is 30 seconds.
- `suspend` (Boolean) Suspend tells the controller to suspend operations - both creation of new Backup and retention actions. This will not have any effect on backups currently in progress. Default is false.

<a id="nestedatt--spec--backup_spec"></a>
### Nested Schema for `spec.backup_spec`

Optional:

- `backup_items` (List of String) For a Physical backup this slice can be used to indicate what PDBs, schemas, tablespaces or tables to back up.
- `backupset` (Boolean) For a Physical backup the choices are Backupset and Image Copies. Backupset is the default, but if Image Copies are required, flip this flag to false.
- `check_logical` (Boolean) For a Physical backup, optionally turn on an additional 'check logical' option. The default is off.
- `compressed` (Boolean) For a Physical backup, optionally turn on compression, by flipping this flag to true. The default is false.
- `dop` (Number) For a Physical backup, optionally indicate a degree of parallelism also known as DOP.
- `filesperset` (Number) For a Physical backup, optionally specify filesperset. The default depends on a type of backup, generally 64.
- `gcs_dir` (String) Similar to GcsPath but specify a Gcs directory. The backup sets of physical backup will be transferred to this GcsDir under a folder named .backup.Spec.Name. This field is usually set in .backupSchedule.Spec.backSpec to specify a GcsDir which all scheduled backups will be uploaded to. A user is to ensure proper write access to the bucket from within the Oracle Operator.
- `gcs_path` (String) If set up ahead of time, the backup sets of a physical backup can be optionally transferred to a GCS bucket. A user is to ensure proper write access to the bucket from within the Oracle Operator.
- `instance` (String) Instance is a name of an instance to take a backup for.
- `keep_data_on_deletion` (Boolean) KeepDataOnDeletion defines whether to keep backup data when backup resource is removed. The default value is false.
- `level` (Number) For a Physical backup, optionally specify an incremental level. The default is 0 (the whole database).
- `local_path` (String) For a Physical backup, optionally specify a local backup dir. If omitted, /u03/app/oracle/rman is assumed.
- `mode` (String) Mode specifies how this backup will be managed by the operator. if it is not set, the operator tries to create a backup based on the specifications. if it is set to VerifyExists, the operator verifies the existence of a backup.
- `section_size` (String) For a Physical backup, optionally specify a section size in various units (K M G).
- `sub_type` (String) Backup sub-type, which is only relevant for a Physical backup type (e.g. RMAN). If omitted, the default of Instance(Level) is assumed. Supported options at this point are: Instance or Database level backups.
- `time_limit_minutes` (Number) For a Physical backup, optionally specify the time threshold. If a threshold is reached, the backup request would time out and error out. The threshold is expressed in minutes. Don't include the unit (minutes), just the integer.
- `type` (String) Type describes a type of a backup to take. Immutable. Available options are: - Snapshot: storage level disk snapshot. - Physical: database engine specific backup that relies on a redo stream / continuous archiving (WAL) and may allow a PITR. Examples include pg_backup, pgBackRest, mysqlbackup. A Physical backup may be file based or database block based (e.g. Oracle RMAN). - Logical: database engine specific backup that relies on running SQL statements, e.g. mysqldump, pg_dump, expdp. If not specified, the default of Snapshot is assumed.
- `volume_snapshot_class` (String) VolumeSnapshotClass points to a particular CSI driver and is used for taking a volume snapshot. If requested here at the Backup level, this setting overrides the platform default as well as the default set via the Config (global user preferences).


<a id="nestedatt--spec--backup_retention_policy"></a>
### Nested Schema for `spec.backup_retention_policy`

Optional:

- `backup_retention` (Number) BackupRetention is the number of successful backups to keep around. The default is 7. A value of 0 means 'do not delete backups based on count'. Max of 512 allows for ~21 days of hourly backups or ~1.4 years of daily backups.
