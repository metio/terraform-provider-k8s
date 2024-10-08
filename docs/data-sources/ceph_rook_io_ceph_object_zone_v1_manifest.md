---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ceph_rook_io_ceph_object_zone_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "ceph.rook.io"
description: |-
  CephObjectZone represents a Ceph Object Store Gateway Zone
---

# k8s_ceph_rook_io_ceph_object_zone_v1_manifest (Data Source)

CephObjectZone represents a Ceph Object Store Gateway Zone

## Example Usage

```terraform
data "k8s_ceph_rook_io_ceph_object_zone_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    data_pool     = {}
    metadata_pool = {}
    zone_group    = "group-1"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) ObjectZoneSpec represent the spec of an ObjectZone (see [below for nested schema](#nestedatt--spec))

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

- `zone_group` (String) The display name for the ceph users

Optional:

- `custom_endpoints` (List of String) If this zone cannot be accessed from other peer Ceph clusters via the ClusterIP Service endpoint created by Rook, you must set this to the externally reachable endpoint(s). You may include the port in the definition. For example: 'https://my-object-store.my-domain.net:443'. In many cases, you should set this to the endpoint of the ingress resource that makes the CephObjectStore associated with this CephObjectStoreZone reachable to peer clusters. The list can have one or more endpoints pointing to different RGW servers in the zone. If a CephObjectStore endpoint is omitted from this list, that object store's gateways will not receive multisite replication data (see CephObjectStore.spec.gateway.disableMultisiteSyncTraffic).
- `data_pool` (Attributes) The data pool settings (see [below for nested schema](#nestedatt--spec--data_pool))
- `metadata_pool` (Attributes) The metadata pool settings (see [below for nested schema](#nestedatt--spec--metadata_pool))
- `preserve_pools_on_delete` (Boolean) Preserve pools on object zone deletion
- `shared_pools` (Attributes) The pool information when configuring RADOS namespaces in existing pools. (see [below for nested schema](#nestedatt--spec--shared_pools))

<a id="nestedatt--spec--data_pool"></a>
### Nested Schema for `spec.data_pool`

Optional:

- `application` (String) The application name to set on the pool. Only expected to be set for rgw pools.
- `compression_mode` (String) DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force' The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force) Do NOT set a default value for kubebuilder as this will override the Parameters
- `crush_root` (String) The root of the crush hierarchy utilized by the pool
- `device_class` (String) The device class the OSD should set to for use in the pool
- `enable_crush_updates` (Boolean) Allow rook operator to change the pool CRUSH tunables once the pool is created
- `enable_rbd_stats` (Boolean) EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool
- `erasure_coded` (Attributes) The erasure code settings (see [below for nested schema](#nestedatt--spec--data_pool--erasure_coded))
- `failure_domain` (String) The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map
- `mirroring` (Attributes) The mirroring settings (see [below for nested schema](#nestedatt--spec--data_pool--mirroring))
- `parameters` (Map of String) Parameters is a list of properties to enable on a given pool
- `quotas` (Attributes) The quota settings (see [below for nested schema](#nestedatt--spec--data_pool--quotas))
- `replicated` (Attributes) The replication settings (see [below for nested schema](#nestedatt--spec--data_pool--replicated))
- `status_check` (Attributes) The mirroring statusCheck (see [below for nested schema](#nestedatt--spec--data_pool--status_check))

<a id="nestedatt--spec--data_pool--erasure_coded"></a>
### Nested Schema for `spec.data_pool.erasure_coded`

Required:

- `coding_chunks` (Number) Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type). This is the number of OSDs that can be lost simultaneously before data cannot be recovered.
- `data_chunks` (Number) Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type). The number of chunks required to recover an object when any single OSD is lost is the same as dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.

Optional:

- `algorithm` (String) The algorithm for erasure coding


<a id="nestedatt--spec--data_pool--mirroring"></a>
### Nested Schema for `spec.data_pool.mirroring`

Optional:

- `enabled` (Boolean) Enabled whether this pool is mirrored or not
- `mode` (String) Mode is the mirroring mode: either pool or image
- `peers` (Attributes) Peers represents the peers spec (see [below for nested schema](#nestedatt--spec--data_pool--mirroring--peers))
- `snapshot_schedules` (Attributes List) SnapshotSchedules is the scheduling of snapshot for mirrored images/pools (see [below for nested schema](#nestedatt--spec--data_pool--mirroring--snapshot_schedules))

<a id="nestedatt--spec--data_pool--mirroring--peers"></a>
### Nested Schema for `spec.data_pool.mirroring.peers`

Optional:

- `secret_names` (List of String) SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers


<a id="nestedatt--spec--data_pool--mirroring--snapshot_schedules"></a>
### Nested Schema for `spec.data_pool.mirroring.snapshot_schedules`

Optional:

- `interval` (String) Interval represent the periodicity of the snapshot.
- `path` (String) Path is the path to snapshot, only valid for CephFS
- `start_time` (String) StartTime indicates when to start the snapshot



<a id="nestedatt--spec--data_pool--quotas"></a>
### Nested Schema for `spec.data_pool.quotas`

Optional:

- `max_bytes` (Number) MaxBytes represents the quota in bytes Deprecated in favor of MaxSize
- `max_objects` (Number) MaxObjects represents the quota in objects
- `max_size` (String) MaxSize represents the quota in bytes as a string


<a id="nestedatt--spec--data_pool--replicated"></a>
### Nested Schema for `spec.data_pool.replicated`

Required:

- `size` (Number) Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)

Optional:

- `hybrid_storage` (Attributes) HybridStorage represents hybrid storage tier settings (see [below for nested schema](#nestedatt--spec--data_pool--replicated--hybrid_storage))
- `replicas_per_failure_domain` (Number) ReplicasPerFailureDomain the number of replica in the specified failure domain
- `require_safe_replica_size` (Boolean) RequireSafeReplicaSize if false allows you to set replica 1
- `sub_failure_domain` (String) SubFailureDomain the name of the sub-failure domain
- `target_size_ratio` (Number) TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity

<a id="nestedatt--spec--data_pool--replicated--hybrid_storage"></a>
### Nested Schema for `spec.data_pool.replicated.hybrid_storage`

Required:

- `primary_device_class` (String) PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD
- `secondary_device_class` (String) SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs



<a id="nestedatt--spec--data_pool--status_check"></a>
### Nested Schema for `spec.data_pool.status_check`

Optional:

- `mirror` (Attributes) HealthCheckSpec represents the health check of an object store bucket (see [below for nested schema](#nestedatt--spec--data_pool--status_check--mirror))

<a id="nestedatt--spec--data_pool--status_check--mirror"></a>
### Nested Schema for `spec.data_pool.status_check.mirror`

Optional:

- `disabled` (Boolean)
- `interval` (String) Interval is the internal in second or minute for the health check to run like 60s for 60 seconds
- `timeout` (String)




<a id="nestedatt--spec--metadata_pool"></a>
### Nested Schema for `spec.metadata_pool`

Optional:

- `application` (String) The application name to set on the pool. Only expected to be set for rgw pools.
- `compression_mode` (String) DEPRECATED: use Parameters instead, e.g., Parameters['compression_mode'] = 'force' The inline compression mode in Bluestore OSD to set to (options are: none, passive, aggressive, force) Do NOT set a default value for kubebuilder as this will override the Parameters
- `crush_root` (String) The root of the crush hierarchy utilized by the pool
- `device_class` (String) The device class the OSD should set to for use in the pool
- `enable_crush_updates` (Boolean) Allow rook operator to change the pool CRUSH tunables once the pool is created
- `enable_rbd_stats` (Boolean) EnableRBDStats is used to enable gathering of statistics for all RBD images in the pool
- `erasure_coded` (Attributes) The erasure code settings (see [below for nested schema](#nestedatt--spec--metadata_pool--erasure_coded))
- `failure_domain` (String) The failure domain: osd/host/(region or zone if available) - technically also any type in the crush map
- `mirroring` (Attributes) The mirroring settings (see [below for nested schema](#nestedatt--spec--metadata_pool--mirroring))
- `parameters` (Map of String) Parameters is a list of properties to enable on a given pool
- `quotas` (Attributes) The quota settings (see [below for nested schema](#nestedatt--spec--metadata_pool--quotas))
- `replicated` (Attributes) The replication settings (see [below for nested schema](#nestedatt--spec--metadata_pool--replicated))
- `status_check` (Attributes) The mirroring statusCheck (see [below for nested schema](#nestedatt--spec--metadata_pool--status_check))

<a id="nestedatt--spec--metadata_pool--erasure_coded"></a>
### Nested Schema for `spec.metadata_pool.erasure_coded`

Required:

- `coding_chunks` (Number) Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type). This is the number of OSDs that can be lost simultaneously before data cannot be recovered.
- `data_chunks` (Number) Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type). The number of chunks required to recover an object when any single OSD is lost is the same as dataChunks so be aware that the larger the number of data chunks, the higher the cost of recovery.

Optional:

- `algorithm` (String) The algorithm for erasure coding


<a id="nestedatt--spec--metadata_pool--mirroring"></a>
### Nested Schema for `spec.metadata_pool.mirroring`

Optional:

- `enabled` (Boolean) Enabled whether this pool is mirrored or not
- `mode` (String) Mode is the mirroring mode: either pool or image
- `peers` (Attributes) Peers represents the peers spec (see [below for nested schema](#nestedatt--spec--metadata_pool--mirroring--peers))
- `snapshot_schedules` (Attributes List) SnapshotSchedules is the scheduling of snapshot for mirrored images/pools (see [below for nested schema](#nestedatt--spec--metadata_pool--mirroring--snapshot_schedules))

<a id="nestedatt--spec--metadata_pool--mirroring--peers"></a>
### Nested Schema for `spec.metadata_pool.mirroring.peers`

Optional:

- `secret_names` (List of String) SecretNames represents the Kubernetes Secret names to add rbd-mirror or cephfs-mirror peers


<a id="nestedatt--spec--metadata_pool--mirroring--snapshot_schedules"></a>
### Nested Schema for `spec.metadata_pool.mirroring.snapshot_schedules`

Optional:

- `interval` (String) Interval represent the periodicity of the snapshot.
- `path` (String) Path is the path to snapshot, only valid for CephFS
- `start_time` (String) StartTime indicates when to start the snapshot



<a id="nestedatt--spec--metadata_pool--quotas"></a>
### Nested Schema for `spec.metadata_pool.quotas`

Optional:

- `max_bytes` (Number) MaxBytes represents the quota in bytes Deprecated in favor of MaxSize
- `max_objects` (Number) MaxObjects represents the quota in objects
- `max_size` (String) MaxSize represents the quota in bytes as a string


<a id="nestedatt--spec--metadata_pool--replicated"></a>
### Nested Schema for `spec.metadata_pool.replicated`

Required:

- `size` (Number) Size - Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)

Optional:

- `hybrid_storage` (Attributes) HybridStorage represents hybrid storage tier settings (see [below for nested schema](#nestedatt--spec--metadata_pool--replicated--hybrid_storage))
- `replicas_per_failure_domain` (Number) ReplicasPerFailureDomain the number of replica in the specified failure domain
- `require_safe_replica_size` (Boolean) RequireSafeReplicaSize if false allows you to set replica 1
- `sub_failure_domain` (String) SubFailureDomain the name of the sub-failure domain
- `target_size_ratio` (Number) TargetSizeRatio gives a hint (%) to Ceph in terms of expected consumption of the total cluster capacity

<a id="nestedatt--spec--metadata_pool--replicated--hybrid_storage"></a>
### Nested Schema for `spec.metadata_pool.replicated.hybrid_storage`

Required:

- `primary_device_class` (String) PrimaryDeviceClass represents high performance tier (for example SSD or NVME) for Primary OSD
- `secondary_device_class` (String) SecondaryDeviceClass represents low performance tier (for example HDDs) for remaining OSDs



<a id="nestedatt--spec--metadata_pool--status_check"></a>
### Nested Schema for `spec.metadata_pool.status_check`

Optional:

- `mirror` (Attributes) HealthCheckSpec represents the health check of an object store bucket (see [below for nested schema](#nestedatt--spec--metadata_pool--status_check--mirror))

<a id="nestedatt--spec--metadata_pool--status_check--mirror"></a>
### Nested Schema for `spec.metadata_pool.status_check.mirror`

Optional:

- `disabled` (Boolean)
- `interval` (String) Interval is the internal in second or minute for the health check to run like 60s for 60 seconds
- `timeout` (String)




<a id="nestedatt--spec--shared_pools"></a>
### Nested Schema for `spec.shared_pools`

Optional:

- `data_pool_name` (String) The data pool used for creating RADOS namespaces in the object store
- `metadata_pool_name` (String) The metadata pool used for creating RADOS namespaces in the object store
- `pool_placements` (Attributes List) PoolPlacements control which Pools are associated with a particular RGW bucket. Once PoolPlacements are defined, RGW client will be able to associate pool with ObjectStore bucket by providing '<LocationConstraint>' during s3 bucket creation or 'X-Storage-Policy' header during swift container creation. See: https://docs.ceph.com/en/latest/radosgw/placement/#placement-targets PoolPlacement with name: 'default' will be used as a default pool if no option is provided during bucket creation. If default placement is not provided, spec.sharedPools.dataPoolName and spec.sharedPools.MetadataPoolName will be used as default pools. If spec.sharedPools are also empty, then RGW pools (spec.dataPool and spec.metadataPool) will be used as defaults. (see [below for nested schema](#nestedatt--spec--shared_pools--pool_placements))
- `preserve_rados_namespace_data_on_delete` (Boolean) Whether the RADOS namespaces should be preserved on deletion of the object store

<a id="nestedatt--spec--shared_pools--pool_placements"></a>
### Nested Schema for `spec.shared_pools.pool_placements`

Required:

- `data_pool_name` (String) The data pool used to store ObjectStore objects data.
- `metadata_pool_name` (String) The metadata pool used to store ObjectStore bucket index.
- `name` (String) Pool placement name. Name can be arbitrary. Placement with name 'default' will be used as default.

Optional:

- `data_non_ec_pool_name` (String) The data pool used to store ObjectStore data that cannot use erasure coding (ex: multi-part uploads). If dataPoolName is not erasure coded, then there is no need for dataNonECPoolName.
- `storage_classes` (Attributes List) StorageClasses can be selected by user to override dataPoolName during object creation. Each placement has default STANDARD StorageClass pointing to dataPoolName. This list allows defining additional StorageClasses on top of default STANDARD storage class. (see [below for nested schema](#nestedatt--spec--shared_pools--pool_placements--storage_classes))

<a id="nestedatt--spec--shared_pools--pool_placements--storage_classes"></a>
### Nested Schema for `spec.shared_pools.pool_placements.storage_classes`

Required:

- `data_pool_name` (String) DataPoolName is the data pool used to store ObjectStore objects data.
- `name` (String) Name is the StorageClass name. Ceph allows arbitrary name for StorageClasses, however most clients/libs insist on AWS names so it is recommended to use one of the valid x-amz-storage-class values for better compatibility: REDUCED_REDUNDANCY | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | GLACIER | DEEP_ARCHIVE | OUTPOSTS | GLACIER_IR | SNOW | EXPRESS_ONEZONE See AWS docs: https://aws.amazon.com/de/s3/storage-classes/
