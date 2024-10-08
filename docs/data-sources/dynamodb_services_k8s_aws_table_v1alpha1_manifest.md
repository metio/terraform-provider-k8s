---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_dynamodb_services_k8s_aws_table_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "dynamodb.services.k8s.aws"
description: |-
  Table is the Schema for the Tables API
---

# k8s_dynamodb_services_k8s_aws_table_v1alpha1_manifest (Data Source)

Table is the Schema for the Tables API

## Example Usage

```terraform
data "k8s_dynamodb_services_k8s_aws_table_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) TableSpec defines the desired state of Table. (see [below for nested schema](#nestedatt--spec))

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

- `attribute_definitions` (Attributes List) An array of attributes that describe the key schema for the table and indexes. (see [below for nested schema](#nestedatt--spec--attribute_definitions))
- `key_schema` (Attributes List) Specifies the attributes that make up the primary key for a table or an index. The attributes in KeySchema must also be defined in the AttributeDefinitions array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html) in the Amazon DynamoDB Developer Guide. Each KeySchemaElement in the array is composed of: * AttributeName - The name of this key attribute. * KeyType - The role that the key attribute will assume: HASH - partition key RANGE - sort key The partition key of an item is also known as its hash attribute. The term 'hash attribute' derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values. The sort key of an item is also known as its range attribute. The term 'range attribute' derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value. For a simple primary key (partition key), you must provide exactly one element with a KeyType of HASH. For a composite primary key (partition key and sort key), you must provide exactly two elements, in this order: The first element must have a KeyType of HASH, and the second element must have a KeyType of RANGE. For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key) in the Amazon DynamoDB Developer Guide. (see [below for nested schema](#nestedatt--spec--key_schema))
- `table_name` (String) The name of the table to create.

Optional:

- `billing_mode` (String) Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later. * PROVISIONED - We recommend using PROVISIONED for predictable workloads. PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual). * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
- `continuous_backups` (Attributes) Represents the settings used to enable point in time recovery. (see [below for nested schema](#nestedatt--spec--continuous_backups))
- `deletion_protection_enabled` (Boolean) Indicates whether deletion protection is to be enabled (true) or disabled (false) on the table.
- `global_secondary_indexes` (Attributes List) One or more global secondary indexes (the maximum is 20) to be created on the table. Each global secondary index in the array includes the following: * IndexName - The name of the global secondary index. Must be unique only for this table. * KeySchema - Specifies the key schema for the global secondary index. * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total. * ProvisionedThroughput - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units. (see [below for nested schema](#nestedatt--spec--global_secondary_indexes))
- `local_secondary_indexes` (Attributes List) One or more local secondary indexes (the maximum is 5) to be created on the table. Each index is scoped to a given partition key value. There is a 10 GB size limit per partition key value; otherwise, the size of a local secondary index is unconstrained. Each local secondary index in the array includes the following: * IndexName - The name of the local secondary index. Must be unique only for this table. * KeySchema - Specifies the key schema for the local secondary index. The key schema must begin with the same partition key as the table. * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total. (see [below for nested schema](#nestedatt--spec--local_secondary_indexes))
- `provisioned_throughput` (Attributes) Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation. If you set BillingMode as PROVISIONED, you must specify this property. If you set BillingMode as PAY_PER_REQUEST, you cannot specify this property. For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide. (see [below for nested schema](#nestedatt--spec--provisioned_throughput))
- `sse_specification` (Attributes) Represents the settings used to enable server-side encryption. (see [below for nested schema](#nestedatt--spec--sse_specification))
- `stream_specification` (Attributes) The settings for DynamoDB Streams on the table. These settings consist of: * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled (true) or disabled (false). * StreamViewType - When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified item are written to the stream. NEW_IMAGE - The entire item, as it appears after it was modified, is written to the stream. OLD_IMAGE - The entire item, as it appeared before it was modified, is written to the stream. NEW_AND_OLD_IMAGES - Both the new and the old item images of the item are written to the stream. (see [below for nested schema](#nestedatt--spec--stream_specification))
- `table_class` (String) The table class of the new table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.
- `tags` (Attributes List) A list of key-value pairs to label the table. For more information, see Tagging for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html). (see [below for nested schema](#nestedatt--spec--tags))
- `time_to_live` (Attributes) Represents the settings used to enable or disable Time to Live for the specified table. (see [below for nested schema](#nestedatt--spec--time_to_live))

<a id="nestedatt--spec--attribute_definitions"></a>
### Nested Schema for `spec.attribute_definitions`

Optional:

- `attribute_name` (String)
- `attribute_type` (String)


<a id="nestedatt--spec--key_schema"></a>
### Nested Schema for `spec.key_schema`

Optional:

- `attribute_name` (String)
- `key_type` (String)


<a id="nestedatt--spec--continuous_backups"></a>
### Nested Schema for `spec.continuous_backups`

Optional:

- `point_in_time_recovery_enabled` (Boolean)


<a id="nestedatt--spec--global_secondary_indexes"></a>
### Nested Schema for `spec.global_secondary_indexes`

Optional:

- `index_name` (String)
- `key_schema` (Attributes List) (see [below for nested schema](#nestedatt--spec--global_secondary_indexes--key_schema))
- `projection` (Attributes) Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. (see [below for nested schema](#nestedatt--spec--global_secondary_indexes--projection))
- `provisioned_throughput` (Attributes) Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation. For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide. (see [below for nested schema](#nestedatt--spec--global_secondary_indexes--provisioned_throughput))

<a id="nestedatt--spec--global_secondary_indexes--key_schema"></a>
### Nested Schema for `spec.global_secondary_indexes.key_schema`

Optional:

- `attribute_name` (String)
- `key_type` (String)


<a id="nestedatt--spec--global_secondary_indexes--projection"></a>
### Nested Schema for `spec.global_secondary_indexes.projection`

Optional:

- `non_key_attributes` (List of String)
- `projection_type` (String)


<a id="nestedatt--spec--global_secondary_indexes--provisioned_throughput"></a>
### Nested Schema for `spec.global_secondary_indexes.provisioned_throughput`

Optional:

- `read_capacity_units` (Number)
- `write_capacity_units` (Number)



<a id="nestedatt--spec--local_secondary_indexes"></a>
### Nested Schema for `spec.local_secondary_indexes`

Optional:

- `index_name` (String)
- `key_schema` (Attributes List) (see [below for nested schema](#nestedatt--spec--local_secondary_indexes--key_schema))
- `projection` (Attributes) Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. (see [below for nested schema](#nestedatt--spec--local_secondary_indexes--projection))

<a id="nestedatt--spec--local_secondary_indexes--key_schema"></a>
### Nested Schema for `spec.local_secondary_indexes.key_schema`

Optional:

- `attribute_name` (String)
- `key_type` (String)


<a id="nestedatt--spec--local_secondary_indexes--projection"></a>
### Nested Schema for `spec.local_secondary_indexes.projection`

Optional:

- `non_key_attributes` (List of String)
- `projection_type` (String)



<a id="nestedatt--spec--provisioned_throughput"></a>
### Nested Schema for `spec.provisioned_throughput`

Optional:

- `read_capacity_units` (Number)
- `write_capacity_units` (Number)


<a id="nestedatt--spec--sse_specification"></a>
### Nested Schema for `spec.sse_specification`

Optional:

- `enabled` (Boolean)
- `kms_master_key_id` (String)
- `sse_type` (String)


<a id="nestedatt--spec--stream_specification"></a>
### Nested Schema for `spec.stream_specification`

Optional:

- `stream_enabled` (Boolean)
- `stream_view_type` (String)


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--spec--time_to_live"></a>
### Nested Schema for `spec.time_to_live`

Optional:

- `attribute_name` (String)
- `enabled` (Boolean)
