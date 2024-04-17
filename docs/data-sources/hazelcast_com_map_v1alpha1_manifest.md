---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_hazelcast_com_map_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "hazelcast.com"
description: |-
  Map is the Schema for the maps API
---

# k8s_hazelcast_com_map_v1alpha1_manifest (Data Source)

Map is the Schema for the maps API

## Example Usage

```terraform
data "k8s_hazelcast_com_map_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    hazelcast_resource_name = "some-name"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) MapSpec defines the desired state of Hazelcast Map Config (see [below for nested schema](#nestedatt--spec))

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

- `hazelcast_resource_name` (String) HazelcastResourceName defines the name of the Hazelcast resource that this resource is created for.

Optional:

- `async_backup_count` (Number) Number of asynchronous backups.
- `attributes` (Attributes List) Attributes to be used with Predicates API. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/predicate-overview#creating-custom-query-attributes (see [below for nested schema](#nestedatt--spec--attributes))
- `backup_count` (Number) Number of synchronous backups.
- `entry_listeners` (Attributes List) EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events. (see [below for nested schema](#nestedatt--spec--entry_listeners))
- `event_journal` (Attributes) EventJournal specifies event journal configuration of the Map (see [below for nested schema](#nestedatt--spec--event_journal))
- `eviction` (Attributes) Configuration for removing data from the map when it reaches its max size. It can be updated. (see [below for nested schema](#nestedatt--spec--eviction))
- `in_memory_format` (String) InMemoryFormat specifies in which format data will be stored in your map
- `indexes` (Attributes List) Indexes to be created for the map data. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps. It cannot be updated after map config is created successfully. (see [below for nested schema](#nestedatt--spec--indexes))
- `map_store` (Attributes) Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data (see [below for nested schema](#nestedatt--spec--map_store))
- `max_idle_seconds` (Number) Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.
- `merkle_tree` (Attributes) MerkleTree defines the configuration for the Merkle tree data structure. (see [below for nested schema](#nestedatt--spec--merkle_tree))
- `name` (String) Name of the data structure config to be created. If empty, CR name will be used. It cannot be updated after the config is created successfully.
- `near_cache` (Attributes) InMemoryFormat specifies near cache configuration for map (see [below for nested schema](#nestedatt--spec--near_cache))
- `persistence_enabled` (Boolean) When enabled, map data will be persisted. It cannot be updated after map config is created successfully.
- `tiered_store` (Attributes) TieredStore enables the Hazelcast's Tiered-Store feature for the Map (see [below for nested schema](#nestedatt--spec--tiered_store))
- `time_to_live_seconds` (Number) Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.

<a id="nestedatt--spec--attributes"></a>
### Nested Schema for `spec.attributes`

Required:

- `extractor_class_name` (String) Name of the extractor class https://docs.hazelcast.com/hazelcast/latest/query/predicate-overview#implementing-a-valueextractor
- `name` (String) Name of the attribute https://docs.hazelcast.com/hazelcast/latest/query/predicate-overview#creating-custom-query-attributes


<a id="nestedatt--spec--entry_listeners"></a>
### Nested Schema for `spec.entry_listeners`

Required:

- `class_name` (String) ClassName is the fully qualified name of the class that implements any of the Listener interface.

Optional:

- `include_values` (Boolean) IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.
- `local` (Boolean) Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.


<a id="nestedatt--spec--event_journal"></a>
### Nested Schema for `spec.event_journal`

Optional:

- `capacity` (Number) Capacity sets the capacity of the ringbuffer underlying the event journal.
- `time_to_live_seconds` (Number) TimeToLiveSeconds indicates how long the items remain in the event journal before they are expired.


<a id="nestedatt--spec--eviction"></a>
### Nested Schema for `spec.eviction`

Optional:

- `eviction_policy` (String) Eviction policy to be applied when map reaches its max size according to the max size policy.
- `max_size` (Number) Max size of the map.
- `max_size_policy` (String) Policy for deciding if the maxSize is reached.


<a id="nestedatt--spec--indexes"></a>
### Nested Schema for `spec.indexes`

Required:

- `type` (String) Type of the index. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#index-types

Optional:

- `attributes` (List of String) Attributes of the index.
- `bit_map_index_options` (Attributes) Options for 'BITMAP' index type. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#configuring-bitmap-indexes (see [below for nested schema](#nestedatt--spec--indexes--bit_map_index_options))
- `name` (String) Name of the index config.

<a id="nestedatt--spec--indexes--bit_map_index_options"></a>
### Nested Schema for `spec.indexes.bit_map_index_options`

Required:

- `unique_key` (String)
- `unique_key_transition` (String)



<a id="nestedatt--spec--map_store"></a>
### Nested Schema for `spec.map_store`

Required:

- `class_name` (String) Name of your class implementing MapLoader and/or MapStore interface.

Optional:

- `initial_mode` (String) Sets the initial entry loading mode.
- `properties_secret_name` (String) Properties can be used for giving information to the MapStore implementation
- `write_batch_size` (Number) Used to create batches when writing to map store.
- `write_coalescing` (Boolean) It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.
- `write_delay_seconds` (Number) Number of seconds to delay the storing of entries.


<a id="nestedatt--spec--merkle_tree"></a>
### Nested Schema for `spec.merkle_tree`

Optional:

- `depth` (Number) Depth of the merkle tree.


<a id="nestedatt--spec--near_cache"></a>
### Nested Schema for `spec.near_cache`

Required:

- `eviction` (Attributes) NearCacheEviction specifies the eviction behavior in Near Cache (see [below for nested schema](#nestedatt--spec--near_cache--eviction))

Optional:

- `cache_local_entries` (Boolean) CacheLocalEntries specifies whether the local entries are cached
- `in_memory_format` (String) InMemoryFormat specifies in which format data will be stored in your near cache
- `invalidate_on_change` (Boolean) InvalidateOnChange specifies whether the cached entries are evicted when the entries are updated or removed
- `max_idle_seconds` (Number) MaxIdleSeconds Maximum number of seconds each entry can stay in the Near Cache as untouched (not read)
- `name` (String) Name is name of the near cache
- `time_to_live_seconds` (Number) TimeToLiveSeconds maximum number of seconds for each entry to stay in the Near Cache

<a id="nestedatt--spec--near_cache--eviction"></a>
### Nested Schema for `spec.near_cache.eviction`

Optional:

- `eviction_policy` (String) EvictionPolicy to be applied when near cache reaches its max size according to the max size policy.
- `max_size_policy` (String) MaxSizePolicy for deciding if the maxSize is reached.
- `size` (Number) Size is maximum size of the Near Cache used for max-size-policy



<a id="nestedatt--spec--tiered_store"></a>
### Nested Schema for `spec.tiered_store`

Optional:

- `disk_device_name` (String) diskDeviceName defines the name of the device for a given disk tier.
- `memory_capacity` (String) MemoryCapacity sets Memory tier capacity, i.e., how much main memory should this tier consume at most.