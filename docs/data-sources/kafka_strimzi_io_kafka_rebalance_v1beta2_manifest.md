---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kafka_strimzi_io_kafka_rebalance_v1beta2_manifest Data Source - terraform-provider-k8s"
subcategory: "kafka.strimzi.io"
description: |-
  
---

# k8s_kafka_strimzi_io_kafka_rebalance_v1beta2_manifest (Data Source)



## Example Usage

```terraform
data "k8s_kafka_strimzi_io_kafka_rebalance_v1beta2_manifest" "example" {
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

- `spec` (Attributes) The specification of the Kafka rebalance. (see [below for nested schema](#nestedatt--spec))

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

- `brokers` (List of String) The list of newly added brokers in case of scaling up or the ones to be removed in case of scaling down to use for rebalancing. This list can be used only with rebalancing mode 'add-brokers' and 'removed-brokers'. It is ignored with 'full' mode.
- `concurrent_intra_broker_partition_movements` (Number) The upper bound of ongoing partition replica movements between disks within each broker. Default is 2.
- `concurrent_leader_movements` (Number) The upper bound of ongoing partition leadership movements. Default is 1000.
- `concurrent_partition_movements_per_broker` (Number) The upper bound of ongoing partition replica movements going into/out of each broker. Default is 5.
- `excluded_topics` (String) A regular expression where any matching topics will be excluded from the calculation of optimization proposals. This expression will be parsed by the java.util.regex.Pattern class; for more information on the supported format consult the documentation for that class.
- `goals` (List of String) A list of goals, ordered by decreasing priority, to use for generating and executing the rebalance proposal. The supported goals are available at https://github.com/linkedin/cruise-control#goals. If an empty goals list is provided, the goals declared in the default.goals Cruise Control configuration parameter are used.
- `mode` (String) Mode to run the rebalancing. The supported modes are 'full', 'add-brokers', 'remove-brokers'.If not specified, the 'full' mode is used by default. * 'full' mode runs the rebalancing across all the brokers in the cluster.* 'add-brokers' mode can be used after scaling up the cluster to move some replicas to the newly added brokers.* 'remove-brokers' mode can be used before scaling down the cluster to move replicas out of the brokers to be removed.
- `rebalance_disk` (Boolean) Enables intra-broker disk balancing, which balances disk space utilization between disks on the same broker. Only applies to Kafka deployments that use JBOD storage with multiple disks. When enabled, inter-broker balancing is disabled. Default is false.
- `replica_movement_strategies` (List of String) A list of strategy class names used to determine the execution order for the replica movements in the generated optimization proposal. By default BaseReplicaMovementStrategy is used, which will execute the replica movements in the order that they were generated.
- `replication_throttle` (Number) The upper bound, in bytes per second, on the bandwidth used to move replicas. There is no limit by default.
- `skip_hard_goal_check` (Boolean) Whether to allow the hard goals specified in the Kafka CR to be skipped in optimization proposal generation. This can be useful when some of those hard goals are preventing a balance solution being found. Default is false.