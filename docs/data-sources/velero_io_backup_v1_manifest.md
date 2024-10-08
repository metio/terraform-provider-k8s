---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_velero_io_backup_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "velero.io"
description: |-
  Backup is a Velero resource that represents the capture of Kubernetes cluster state at a point in time (API objects and associated volume state).
---

# k8s_velero_io_backup_v1_manifest (Data Source)

Backup is a Velero resource that represents the capture of Kubernetes cluster state at a point in time (API objects and associated volume state).

## Example Usage

```terraform
data "k8s_velero_io_backup_v1_manifest" "example" {
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

- `spec` (Attributes) BackupSpec defines the specification for a Velero backup. (see [below for nested schema](#nestedatt--spec))

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

- `csi_snapshot_timeout` (String) CSISnapshotTimeout specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse during creation, before returning error as timeout. The default value is 10 minute.
- `datamover` (String) DataMover specifies the data mover to be used by the backup. If DataMover is '' or 'velero', the built-in data mover will be used.
- `default_volumes_to_fs_backup` (Boolean) DefaultVolumesToFsBackup specifies whether pod volume file system backup should be used for all volumes by default.
- `default_volumes_to_restic` (Boolean) DefaultVolumesToRestic specifies whether restic should be used to take a backup of all pod volumes by default. Deprecated: this field is no longer used and will be removed entirely in future. Use DefaultVolumesToFsBackup instead.
- `excluded_cluster_scoped_resources` (List of String) ExcludedClusterScopedResources is a slice of cluster-scoped resource type names to exclude from the backup. If set to '*', all cluster-scoped resource types are excluded. The default value is empty.
- `excluded_namespace_scoped_resources` (List of String) ExcludedNamespaceScopedResources is a slice of namespace-scoped resource type names to exclude from the backup. If set to '*', all namespace-scoped resource types are excluded. The default value is empty.
- `excluded_namespaces` (List of String) ExcludedNamespaces contains a list of namespaces that are not included in the backup.
- `excluded_resources` (List of String) ExcludedResources is a slice of resource names that are not included in the backup.
- `hooks` (Attributes) Hooks represent custom behaviors that should be executed at different phases of the backup. (see [below for nested schema](#nestedatt--spec--hooks))
- `include_cluster_resources` (Boolean) IncludeClusterResources specifies whether cluster-scoped resources should be included for consideration in the backup.
- `included_cluster_scoped_resources` (List of String) IncludedClusterScopedResources is a slice of cluster-scoped resource type names to include in the backup. If set to '*', all cluster-scoped resource types are included. The default value is empty, which means only related cluster-scoped resources are included.
- `included_namespace_scoped_resources` (List of String) IncludedNamespaceScopedResources is a slice of namespace-scoped resource type names to include in the backup. The default value is '*'.
- `included_namespaces` (List of String) IncludedNamespaces is a slice of namespace names to include objects from. If empty, all namespaces are included.
- `included_resources` (List of String) IncludedResources is a slice of resource names to include in the backup. If empty, all resources are included.
- `item_operation_timeout` (String) ItemOperationTimeout specifies the time used to wait for asynchronous BackupItemAction operations The default value is 4 hour.
- `label_selector` (Attributes) LabelSelector is a metav1.LabelSelector to filter with when adding individual objects to the backup. If empty or nil, all objects are included. Optional. (see [below for nested schema](#nestedatt--spec--label_selector))
- `metadata` (Attributes) (see [below for nested schema](#nestedatt--spec--metadata))
- `or_label_selectors` (Attributes List) OrLabelSelectors is list of metav1.LabelSelector to filter with when adding individual objects to the backup. If multiple provided they will be joined by the OR operator. LabelSelector as well as OrLabelSelectors cannot co-exist in backup request, only one of them can be used. (see [below for nested schema](#nestedatt--spec--or_label_selectors))
- `ordered_resources` (Map of String) OrderedResources specifies the backup order of resources of specific Kind. The map key is the resource name and value is a list of object names separated by commas. Each resource name has format 'namespace/objectname'. For cluster resources, simply use 'objectname'.
- `resource_policy` (Attributes) ResourcePolicy specifies the referenced resource policies that backup should follow (see [below for nested schema](#nestedatt--spec--resource_policy))
- `snapshot_move_data` (Boolean) SnapshotMoveData specifies whether snapshot data should be moved
- `snapshot_volumes` (Boolean) SnapshotVolumes specifies whether to take snapshots of any PV's referenced in the set of objects included in the Backup.
- `storage_location` (String) StorageLocation is a string containing the name of a BackupStorageLocation where the backup should be stored.
- `ttl` (String) TTL is a time.Duration-parseable string describing how long the Backup should be retained for.
- `uploader_config` (Attributes) UploaderConfig specifies the configuration for the uploader. (see [below for nested schema](#nestedatt--spec--uploader_config))
- `volume_snapshot_locations` (List of String) VolumeSnapshotLocations is a list containing names of VolumeSnapshotLocations associated with this backup.

<a id="nestedatt--spec--hooks"></a>
### Nested Schema for `spec.hooks`

Optional:

- `resources` (Attributes List) Resources are hooks that should be executed when backing up individual instances of a resource. (see [below for nested schema](#nestedatt--spec--hooks--resources))

<a id="nestedatt--spec--hooks--resources"></a>
### Nested Schema for `spec.hooks.resources`

Required:

- `name` (String) Name is the name of this hook.

Optional:

- `excluded_namespaces` (List of String) ExcludedNamespaces specifies the namespaces to which this hook spec does not apply.
- `excluded_resources` (List of String) ExcludedResources specifies the resources to which this hook spec does not apply.
- `included_namespaces` (List of String) IncludedNamespaces specifies the namespaces to which this hook spec applies. If empty, it applies to all namespaces.
- `included_resources` (List of String) IncludedResources specifies the resources to which this hook spec applies. If empty, it applies to all resources.
- `label_selector` (Attributes) LabelSelector, if specified, filters the resources to which this hook spec applies. (see [below for nested schema](#nestedatt--spec--hooks--resources--label_selector))
- `post` (Attributes List) PostHooks is a list of BackupResourceHooks to execute after storing the item in the backup. These are executed after all 'additional items' from item actions are processed. (see [below for nested schema](#nestedatt--spec--hooks--resources--post))
- `pre` (Attributes List) PreHooks is a list of BackupResourceHooks to execute prior to storing the item in the backup. These are executed before any 'additional items' from item actions are processed. (see [below for nested schema](#nestedatt--spec--hooks--resources--pre))

<a id="nestedatt--spec--hooks--resources--label_selector"></a>
### Nested Schema for `spec.hooks.resources.label_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--hooks--resources--label_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--hooks--resources--label_selector--match_expressions"></a>
### Nested Schema for `spec.hooks.resources.label_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--hooks--resources--post"></a>
### Nested Schema for `spec.hooks.resources.post`

Required:

- `exec` (Attributes) Exec defines an exec hook. (see [below for nested schema](#nestedatt--spec--hooks--resources--post--exec))

<a id="nestedatt--spec--hooks--resources--post--exec"></a>
### Nested Schema for `spec.hooks.resources.post.exec`

Required:

- `command` (List of String) Command is the command and arguments to execute.

Optional:

- `container` (String) Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.
- `on_error` (String) OnError specifies how Velero should behave if it encounters an error executing this hook.
- `timeout` (String) Timeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.



<a id="nestedatt--spec--hooks--resources--pre"></a>
### Nested Schema for `spec.hooks.resources.pre`

Required:

- `exec` (Attributes) Exec defines an exec hook. (see [below for nested schema](#nestedatt--spec--hooks--resources--pre--exec))

<a id="nestedatt--spec--hooks--resources--pre--exec"></a>
### Nested Schema for `spec.hooks.resources.pre.exec`

Required:

- `command` (List of String) Command is the command and arguments to execute.

Optional:

- `container` (String) Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.
- `on_error` (String) OnError specifies how Velero should behave if it encounters an error executing this hook.
- `timeout` (String) Timeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.





<a id="nestedatt--spec--label_selector"></a>
### Nested Schema for `spec.label_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--label_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--label_selector--match_expressions"></a>
### Nested Schema for `spec.label_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--metadata"></a>
### Nested Schema for `spec.metadata`

Optional:

- `labels` (Map of String)


<a id="nestedatt--spec--or_label_selectors"></a>
### Nested Schema for `spec.or_label_selectors`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--or_label_selectors--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--or_label_selectors--match_expressions"></a>
### Nested Schema for `spec.or_label_selectors.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--resource_policy"></a>
### Nested Schema for `spec.resource_policy`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.


<a id="nestedatt--spec--uploader_config"></a>
### Nested Schema for `spec.uploader_config`

Optional:

- `parallel_files_upload` (Number) ParallelFilesUpload is the number of files parallel uploads to perform when using the uploader.
